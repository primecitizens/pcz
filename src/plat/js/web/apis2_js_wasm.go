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

func NewHTMLSlotElement() HTMLSlotElement {
	return HTMLSlotElement{}.FromRef(
		bindings.NewHTMLSlotElementByHTMLSlotElement(),
	)
}

type HTMLSlotElement struct {
	HTMLElement
}

func (this HTMLSlotElement) Once() HTMLSlotElement {
	this.Ref().Once()
	return this
}

func (this HTMLSlotElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLSlotElement) FromRef(ref js.Ref) HTMLSlotElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLSlotElement) Free() {
	this.Ref().Free()
}

// Name returns the value of property "HTMLSlotElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLSlotElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSlotElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLSlotElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSlotElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLSlotElementName(
		this.Ref(),
		val.Ref(),
	)
}

// AssignedNodes calls the method "HTMLSlotElement.assignedNodes".
//
// The returned bool will be false if there is no such method.
func (this HTMLSlotElement) AssignedNodes(options AssignedNodesOptions) (js.Array[Node], bool) {
	var _ok bool
	_ret := bindings.CallHTMLSlotElementAssignedNodes(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[Node]{}.FromRef(_ret), _ok
}

// AssignedNodesFunc returns the method "HTMLSlotElement.assignedNodes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSlotElement) AssignedNodesFunc() (fn js.Func[func(options AssignedNodesOptions) js.Array[Node]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedNodesFunc(
			this.Ref(),
		),
	)
}

// AssignedNodes1 calls the method "HTMLSlotElement.assignedNodes".
//
// The returned bool will be false if there is no such method.
func (this HTMLSlotElement) AssignedNodes1() (js.Array[Node], bool) {
	var _ok bool
	_ret := bindings.CallHTMLSlotElementAssignedNodes1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Node]{}.FromRef(_ret), _ok
}

// AssignedNodes1Func returns the method "HTMLSlotElement.assignedNodes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSlotElement) AssignedNodes1Func() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedNodes1Func(
			this.Ref(),
		),
	)
}

// AssignedElements calls the method "HTMLSlotElement.assignedElements".
//
// The returned bool will be false if there is no such method.
func (this HTMLSlotElement) AssignedElements(options AssignedNodesOptions) (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallHTMLSlotElementAssignedElements(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// AssignedElementsFunc returns the method "HTMLSlotElement.assignedElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSlotElement) AssignedElementsFunc() (fn js.Func[func(options AssignedNodesOptions) js.Array[Element]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedElementsFunc(
			this.Ref(),
		),
	)
}

// AssignedElements1 calls the method "HTMLSlotElement.assignedElements".
//
// The returned bool will be false if there is no such method.
func (this HTMLSlotElement) AssignedElements1() (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallHTMLSlotElementAssignedElements1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// AssignedElements1Func returns the method "HTMLSlotElement.assignedElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSlotElement) AssignedElements1Func() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedElements1Func(
			this.Ref(),
		),
	)
}

// Assign calls the method "HTMLSlotElement.assign".
//
// The returned bool will be false if there is no such method.
func (this HTMLSlotElement) Assign(nodes ...OneOf_Element_Text) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSlotElementAssign(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AssignFunc returns the method "HTMLSlotElement.assign".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSlotElement) AssignFunc() (fn js.Func[func(nodes ...OneOf_Element_Text)]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignFunc(
			this.Ref(),
		),
	)
}

func NewText(data js.String) Text {
	return Text{}.FromRef(
		bindings.NewTextByText(
			data.Ref()),
	)
}

func NewTextByText1() Text {
	return Text{}.FromRef(
		bindings.NewTextByText1(),
	)
}

type Text struct {
	CharacterData
}

func (this Text) Once() Text {
	this.Ref().Once()
	return this
}

func (this Text) Ref() js.Ref {
	return this.CharacterData.Ref()
}

func (this Text) FromRef(ref js.Ref) Text {
	this.CharacterData = this.CharacterData.FromRef(ref)
	return this
}

func (this Text) Free() {
	this.Ref().Free()
}

// WholeText returns the value of property "Text.wholeText".
//
// The returned bool will be false if there is no such property.
func (this Text) WholeText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextWholeText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AssignedSlot returns the value of property "Text.assignedSlot".
//
// The returned bool will be false if there is no such property.
func (this Text) AssignedSlot() (HTMLSlotElement, bool) {
	var _ok bool
	_ret := bindings.GetTextAssignedSlot(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLSlotElement{}.FromRef(_ret), _ok
}

// SplitText calls the method "Text.splitText".
//
// The returned bool will be false if there is no such method.
func (this Text) SplitText(offset uint32) (Text, bool) {
	var _ok bool
	_ret := bindings.CallTextSplitText(
		this.Ref(), js.Pointer(&_ok),
		uint32(offset),
	)

	return Text{}.FromRef(_ret), _ok
}

// SplitTextFunc returns the method "Text.splitText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) SplitTextFunc() (fn js.Func[func(offset uint32) Text]) {
	return fn.FromRef(
		bindings.TextSplitTextFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Text.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Text) GetBoxQuads(options BoxQuadOptions) (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallTextGetBoxQuads(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuadsFunc returns the method "Text.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.TextGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Text.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Text) GetBoxQuads1() (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallTextGetBoxQuads1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuads1Func returns the method "Text.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.TextGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Text.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertQuadFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNodeFunc returns the method "Text.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Text.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertQuadFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNode1Func returns the method "Text.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Text.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertRectFromNode(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNodeFunc returns the method "Text.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Text.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertRectFromNode1(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNode1Func returns the method "Text.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Text.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertPointFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNodeFunc returns the method "Text.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.TextConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Text.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Text) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallTextConvertPointFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNode1Func returns the method "Text.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Text) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.TextConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

type OneOf_Text_Element_CSSPseudoElement_Document struct {
	ref js.Ref
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) Free() {
	x.ref.Free()
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) FromRef(ref js.Ref) OneOf_Text_Element_CSSPseudoElement_Document {
	return OneOf_Text_Element_CSSPseudoElement_Document{
		ref: ref,
	}
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) Text() Text {
	return Text{}.FromRef(x.ref)
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) Element() Element {
	return Element{}.FromRef(x.ref)
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) CSSPseudoElement() CSSPseudoElement {
	return CSSPseudoElement{}.FromRef(x.ref)
}

func (x OneOf_Text_Element_CSSPseudoElement_Document) Document() Document {
	return Document{}.FromRef(x.ref)
}

type GeometryNode = OneOf_Text_Element_CSSPseudoElement_Document

type BoxQuadOptions struct {
	// Box is "BoxQuadOptions.box"
	//
	// Optional, defaults to "border".
	Box CSSBoxType
	// RelativeTo is "BoxQuadOptions.relativeTo"
	//
	// Optional
	RelativeTo GeometryNode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BoxQuadOptions with all fields set.
func (p BoxQuadOptions) FromRef(ref js.Ref) BoxQuadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BoxQuadOptions in the application heap.
func (p BoxQuadOptions) New() js.Ref {
	return bindings.BoxQuadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BoxQuadOptions) UpdateFrom(ref js.Ref) {
	bindings.BoxQuadOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BoxQuadOptions) Update(ref js.Ref) {
	bindings.BoxQuadOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Element_CSSPseudoElement struct {
	ref js.Ref
}

func (x OneOf_Element_CSSPseudoElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Element_CSSPseudoElement) Free() {
	x.ref.Free()
}

func (x OneOf_Element_CSSPseudoElement) FromRef(ref js.Ref) OneOf_Element_CSSPseudoElement {
	return OneOf_Element_CSSPseudoElement{
		ref: ref,
	}
}

func (x OneOf_Element_CSSPseudoElement) Element() Element {
	return Element{}.FromRef(x.ref)
}

func (x OneOf_Element_CSSPseudoElement) CSSPseudoElement() CSSPseudoElement {
	return CSSPseudoElement{}.FromRef(x.ref)
}

type CSSPseudoElement struct {
	EventTarget
}

func (this CSSPseudoElement) Once() CSSPseudoElement {
	this.Ref().Once()
	return this
}

func (this CSSPseudoElement) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this CSSPseudoElement) FromRef(ref js.Ref) CSSPseudoElement {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this CSSPseudoElement) Free() {
	this.Ref().Free()
}

// Type returns the value of property "CSSPseudoElement.type".
//
// The returned bool will be false if there is no such property.
func (this CSSPseudoElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPseudoElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Element returns the value of property "CSSPseudoElement.element".
//
// The returned bool will be false if there is no such property.
func (this CSSPseudoElement) Element() (Element, bool) {
	var _ok bool
	_ret := bindings.GetCSSPseudoElementElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Parent returns the value of property "CSSPseudoElement.parent".
//
// The returned bool will be false if there is no such property.
func (this CSSPseudoElement) Parent() (OneOf_Element_CSSPseudoElement, bool) {
	var _ok bool
	_ret := bindings.GetCSSPseudoElementParent(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Element_CSSPseudoElement{}.FromRef(_ret), _ok
}

// Pseudo calls the method "CSSPseudoElement.pseudo".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) Pseudo(typ js.String) (CSSPseudoElement, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementPseudo(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return CSSPseudoElement{}.FromRef(_ret), _ok
}

// PseudoFunc returns the method "CSSPseudoElement.pseudo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) PseudoFunc() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	return fn.FromRef(
		bindings.CSSPseudoElementPseudoFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "CSSPseudoElement.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) GetBoxQuads(options BoxQuadOptions) (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementGetBoxQuads(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuadsFunc returns the method "CSSPseudoElement.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.CSSPseudoElementGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "CSSPseudoElement.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) GetBoxQuads1() (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementGetBoxQuads1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuads1Func returns the method "CSSPseudoElement.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.CSSPseudoElementGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "CSSPseudoElement.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNodeFunc returns the method "CSSPseudoElement.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "CSSPseudoElement.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNode1Func returns the method "CSSPseudoElement.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "CSSPseudoElement.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertRectFromNode(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNodeFunc returns the method "CSSPseudoElement.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "CSSPseudoElement.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNode1Func returns the method "CSSPseudoElement.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "CSSPseudoElement.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertPointFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNodeFunc returns the method "CSSPseudoElement.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "CSSPseudoElement.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this CSSPseudoElement) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallCSSPseudoElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNode1Func returns the method "CSSPseudoElement.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSPseudoElement) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

type AttributeMatchList = js.Record[js.Array[js.String]]

type SanitizerConfig struct {
	// AllowElements is "SanitizerConfig.allowElements"
	//
	// Optional
	AllowElements js.Array[js.String]
	// BlockElements is "SanitizerConfig.blockElements"
	//
	// Optional
	BlockElements js.Array[js.String]
	// DropElements is "SanitizerConfig.dropElements"
	//
	// Optional
	DropElements js.Array[js.String]
	// AllowAttributes is "SanitizerConfig.allowAttributes"
	//
	// Optional
	AllowAttributes AttributeMatchList
	// DropAttributes is "SanitizerConfig.dropAttributes"
	//
	// Optional
	DropAttributes AttributeMatchList
	// AllowCustomElements is "SanitizerConfig.allowCustomElements"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowCustomElements MUST be set to true to make this field effective.
	AllowCustomElements bool
	// AllowUnknownMarkup is "SanitizerConfig.allowUnknownMarkup"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowUnknownMarkup MUST be set to true to make this field effective.
	AllowUnknownMarkup bool
	// AllowComments is "SanitizerConfig.allowComments"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowComments MUST be set to true to make this field effective.
	AllowComments bool

	FFI_USE_AllowCustomElements bool // for AllowCustomElements.
	FFI_USE_AllowUnknownMarkup  bool // for AllowUnknownMarkup.
	FFI_USE_AllowComments       bool // for AllowComments.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SanitizerConfig with all fields set.
func (p SanitizerConfig) FromRef(ref js.Ref) SanitizerConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SanitizerConfig in the application heap.
func (p SanitizerConfig) New() js.Ref {
	return bindings.SanitizerConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SanitizerConfig) UpdateFrom(ref js.Ref) {
	bindings.SanitizerConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SanitizerConfig) Update(ref js.Ref) {
	bindings.SanitizerConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Node_String struct {
	ref js.Ref
}

func (x OneOf_Node_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Node_String) Free() {
	x.ref.Free()
}

func (x OneOf_Node_String) FromRef(ref js.Ref) OneOf_Node_String {
	return OneOf_Node_String{
		ref: ref,
	}
}

func (x OneOf_Node_String) Node() Node {
	return Node{}.FromRef(x.ref)
}

func (x OneOf_Node_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type NodeList struct {
	ref js.Ref
}

func (this NodeList) Once() NodeList {
	this.Ref().Once()
	return this
}

func (this NodeList) Ref() js.Ref {
	return this.ref
}

func (this NodeList) FromRef(ref js.Ref) NodeList {
	this.ref = ref
	return this
}

func (this NodeList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "NodeList.length".
//
// The returned bool will be false if there is no such property.
func (this NodeList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetNodeListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "NodeList.item".
//
// The returned bool will be false if there is no such method.
func (this NodeList) Item(index uint32) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Node{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "NodeList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NodeList) ItemFunc() (fn js.Func[func(index uint32) Node]) {
	return fn.FromRef(
		bindings.NodeListItemFunc(
			this.Ref(),
		),
	)
}

type DocumentFragment struct {
	Node
}

func (this DocumentFragment) Once() DocumentFragment {
	this.Ref().Once()
	return this
}

func (this DocumentFragment) Ref() js.Ref {
	return this.Node.Ref()
}

func (this DocumentFragment) FromRef(ref js.Ref) DocumentFragment {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this DocumentFragment) Free() {
	this.Ref().Free()
}

// Children returns the value of property "DocumentFragment.children".
//
// The returned bool will be false if there is no such property.
func (this DocumentFragment) Children() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFragmentChildren(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// FirstElementChild returns the value of property "DocumentFragment.firstElementChild".
//
// The returned bool will be false if there is no such property.
func (this DocumentFragment) FirstElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFragmentFirstElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// LastElementChild returns the value of property "DocumentFragment.lastElementChild".
//
// The returned bool will be false if there is no such property.
func (this DocumentFragment) LastElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFragmentLastElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ChildElementCount returns the value of property "DocumentFragment.childElementCount".
//
// The returned bool will be false if there is no such property.
func (this DocumentFragment) ChildElementCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFragmentChildElementCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// GetElementById calls the method "DocumentFragment.getElementById".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) GetElementById(elementId js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentGetElementById(
		this.Ref(), js.Pointer(&_ok),
		elementId.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// GetElementByIdFunc returns the method "DocumentFragment.getElementById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentFragmentGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "DocumentFragment.prepend".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) Prepend(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentPrepend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PrependFunc returns the method "DocumentFragment.prepend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentPrependFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "DocumentFragment.append".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) Append(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentAppend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "DocumentFragment.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentAppendFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "DocumentFragment.replaceChildren".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) ReplaceChildren(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentReplaceChildren(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceChildrenFunc returns the method "DocumentFragment.replaceChildren".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "DocumentFragment.querySelector".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) QuerySelector(selectors js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentQuerySelector(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// QuerySelectorFunc returns the method "DocumentFragment.querySelector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentFragmentQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "DocumentFragment.querySelectorAll".
//
// The returned bool will be false if there is no such method.
func (this DocumentFragment) QuerySelectorAll(selectors js.String) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallDocumentFragmentQuerySelectorAll(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// QuerySelectorAllFunc returns the method "DocumentFragment.querySelectorAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentFragment) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentFragmentQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

type OneOf_Document_DocumentFragment struct {
	ref js.Ref
}

func (x OneOf_Document_DocumentFragment) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Document_DocumentFragment) Free() {
	x.ref.Free()
}

func (x OneOf_Document_DocumentFragment) FromRef(ref js.Ref) OneOf_Document_DocumentFragment {
	return OneOf_Document_DocumentFragment{
		ref: ref,
	}
}

func (x OneOf_Document_DocumentFragment) Document() Document {
	return Document{}.FromRef(x.ref)
}

func (x OneOf_Document_DocumentFragment) DocumentFragment() DocumentFragment {
	return DocumentFragment{}.FromRef(x.ref)
}

func NewSanitizer(config SanitizerConfig) Sanitizer {
	return Sanitizer{}.FromRef(
		bindings.NewSanitizerBySanitizer(
			js.Pointer(&config)),
	)
}

func NewSanitizerBySanitizer1() Sanitizer {
	return Sanitizer{}.FromRef(
		bindings.NewSanitizerBySanitizer1(),
	)
}

type Sanitizer struct {
	ref js.Ref
}

func (this Sanitizer) Once() Sanitizer {
	this.Ref().Once()
	return this
}

func (this Sanitizer) Ref() js.Ref {
	return this.ref
}

func (this Sanitizer) FromRef(ref js.Ref) Sanitizer {
	this.ref = ref
	return this
}

func (this Sanitizer) Free() {
	this.Ref().Free()
}

// Sanitize calls the method "Sanitizer.sanitize".
//
// The returned bool will be false if there is no such method.
func (this Sanitizer) Sanitize(input OneOf_Document_DocumentFragment) (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallSanitizerSanitize(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// SanitizeFunc returns the method "Sanitizer.sanitize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sanitizer) SanitizeFunc() (fn js.Func[func(input OneOf_Document_DocumentFragment) DocumentFragment]) {
	return fn.FromRef(
		bindings.SanitizerSanitizeFunc(
			this.Ref(),
		),
	)
}

// SanitizeFor calls the method "Sanitizer.sanitizeFor".
//
// The returned bool will be false if there is no such method.
func (this Sanitizer) SanitizeFor(element js.String, input js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallSanitizerSanitizeFor(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
		input.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// SanitizeForFunc returns the method "Sanitizer.sanitizeFor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sanitizer) SanitizeForFunc() (fn js.Func[func(element js.String, input js.String) Element]) {
	return fn.FromRef(
		bindings.SanitizerSanitizeForFunc(
			this.Ref(),
		),
	)
}

// GetConfiguration calls the method "Sanitizer.getConfiguration".
//
// The returned bool will be false if there is no such method.
func (this Sanitizer) GetConfiguration() (SanitizerConfig, bool) {
	var _ret SanitizerConfig
	_ok := js.True == bindings.CallSanitizerGetConfiguration(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetConfigurationFunc returns the method "Sanitizer.getConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sanitizer) GetConfigurationFunc() (fn js.Func[func() SanitizerConfig]) {
	return fn.FromRef(
		bindings.SanitizerGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// GetDefaultConfiguration calls the staticmethod "Sanitizer.getDefaultConfiguration".
//
// The returned bool will be false if there is no such method.
func (this Sanitizer) GetDefaultConfiguration() (SanitizerConfig, bool) {
	var _ret SanitizerConfig
	_ok := js.True == bindings.CallSanitizerGetDefaultConfiguration(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetDefaultConfigurationFunc returns the staticmethod "Sanitizer.getDefaultConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sanitizer) GetDefaultConfigurationFunc() (fn js.Func[func() SanitizerConfig]) {
	return fn.FromRef(
		bindings.SanitizerGetDefaultConfigurationFunc(
			this.Ref(),
		),
	)
}

type SetHTMLOptions struct {
	// Sanitizer is "SetHTMLOptions.sanitizer"
	//
	// Optional
	Sanitizer Sanitizer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetHTMLOptions with all fields set.
func (p SetHTMLOptions) FromRef(ref js.Ref) SetHTMLOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetHTMLOptions in the application heap.
func (p SetHTMLOptions) New() js.Ref {
	return bindings.SetHTMLOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SetHTMLOptions) UpdateFrom(ref js.Ref) {
	bindings.SetHTMLOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SetHTMLOptions) Update(ref js.Ref) {
	bindings.SetHTMLOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DOMRectList struct {
	ref js.Ref
}

func (this DOMRectList) Once() DOMRectList {
	this.Ref().Once()
	return this
}

func (this DOMRectList) Ref() js.Ref {
	return this.ref
}

func (this DOMRectList) FromRef(ref js.Ref) DOMRectList {
	this.ref = ref
	return this
}

func (this DOMRectList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "DOMRectList.length".
//
// The returned bool will be false if there is no such property.
func (this DOMRectList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "DOMRectList.item".
//
// The returned bool will be false if there is no such method.
func (this DOMRectList) Item(index uint32) (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "DOMRectList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRectList) ItemFunc() (fn js.Func[func(index uint32) DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectListItemFunc(
			this.Ref(),
		),
	)
}

type CheckVisibilityOptions struct {
	// CheckOpacity is "CheckVisibilityOptions.checkOpacity"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CheckOpacity MUST be set to true to make this field effective.
	CheckOpacity bool
	// CheckVisibilityCSS is "CheckVisibilityOptions.checkVisibilityCSS"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CheckVisibilityCSS MUST be set to true to make this field effective.
	CheckVisibilityCSS bool

	FFI_USE_CheckOpacity       bool // for CheckOpacity.
	FFI_USE_CheckVisibilityCSS bool // for CheckVisibilityCSS.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CheckVisibilityOptions with all fields set.
func (p CheckVisibilityOptions) FromRef(ref js.Ref) CheckVisibilityOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CheckVisibilityOptions in the application heap.
func (p CheckVisibilityOptions) New() js.Ref {
	return bindings.CheckVisibilityOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CheckVisibilityOptions) UpdateFrom(ref js.Ref) {
	bindings.CheckVisibilityOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CheckVisibilityOptions) Update(ref js.Ref) {
	bindings.CheckVisibilityOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScrollLogicalPosition uint32

const (
	_ ScrollLogicalPosition = iota

	ScrollLogicalPosition_START
	ScrollLogicalPosition_CENTER
	ScrollLogicalPosition_END
	ScrollLogicalPosition_NEAREST
)

func (ScrollLogicalPosition) FromRef(str js.Ref) ScrollLogicalPosition {
	return ScrollLogicalPosition(bindings.ConstOfScrollLogicalPosition(str))
}

func (x ScrollLogicalPosition) String() (string, bool) {
	switch x {
	case ScrollLogicalPosition_START:
		return "start", true
	case ScrollLogicalPosition_CENTER:
		return "center", true
	case ScrollLogicalPosition_END:
		return "end", true
	case ScrollLogicalPosition_NEAREST:
		return "nearest", true
	default:
		return "", false
	}
}

type ScrollBehavior uint32

const (
	_ ScrollBehavior = iota

	ScrollBehavior_AUTO
	ScrollBehavior_INSTANT
	ScrollBehavior_SMOOTH
)

func (ScrollBehavior) FromRef(str js.Ref) ScrollBehavior {
	return ScrollBehavior(bindings.ConstOfScrollBehavior(str))
}

func (x ScrollBehavior) String() (string, bool) {
	switch x {
	case ScrollBehavior_AUTO:
		return "auto", true
	case ScrollBehavior_INSTANT:
		return "instant", true
	case ScrollBehavior_SMOOTH:
		return "smooth", true
	default:
		return "", false
	}
}

type ScrollIntoViewOptions struct {
	// Block is "ScrollIntoViewOptions.block"
	//
	// Optional, defaults to "start".
	Block ScrollLogicalPosition
	// Inline is "ScrollIntoViewOptions.inline"
	//
	// Optional, defaults to "nearest".
	Inline ScrollLogicalPosition
	// Behavior is "ScrollIntoViewOptions.behavior"
	//
	// Optional, defaults to "auto".
	Behavior ScrollBehavior

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollIntoViewOptions with all fields set.
func (p ScrollIntoViewOptions) FromRef(ref js.Ref) ScrollIntoViewOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollIntoViewOptions in the application heap.
func (p ScrollIntoViewOptions) New() js.Ref {
	return bindings.ScrollIntoViewOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScrollIntoViewOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollIntoViewOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScrollIntoViewOptions) Update(ref js.Ref) {
	bindings.ScrollIntoViewOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Bool_ScrollIntoViewOptions struct {
	ref js.Ref
}

func (x OneOf_Bool_ScrollIntoViewOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_ScrollIntoViewOptions) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_ScrollIntoViewOptions) FromRef(ref js.Ref) OneOf_Bool_ScrollIntoViewOptions {
	return OneOf_Bool_ScrollIntoViewOptions{
		ref: ref,
	}
}

func (x OneOf_Bool_ScrollIntoViewOptions) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_ScrollIntoViewOptions) ScrollIntoViewOptions() ScrollIntoViewOptions {
	var ret ScrollIntoViewOptions
	ret.UpdateFrom(x.ref)
	return ret
}

type ScrollToOptions struct {
	// Left is "ScrollToOptions.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left float64
	// Top is "ScrollToOptions.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top float64
	// Behavior is "ScrollToOptions.behavior"
	//
	// Optional, defaults to "auto".
	Behavior ScrollBehavior

	FFI_USE_Left bool // for Left.
	FFI_USE_Top  bool // for Top.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollToOptions with all fields set.
func (p ScrollToOptions) FromRef(ref js.Ref) ScrollToOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollToOptions in the application heap.
func (p ScrollToOptions) New() js.Ref {
	return bindings.ScrollToOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScrollToOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollToOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScrollToOptions) Update(ref js.Ref) {
	bindings.ScrollToOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IterationCompositeOperation uint32

const (
	_ IterationCompositeOperation = iota

	IterationCompositeOperation_REPLACE
	IterationCompositeOperation_ACCUMULATE
)

func (IterationCompositeOperation) FromRef(str js.Ref) IterationCompositeOperation {
	return IterationCompositeOperation(bindings.ConstOfIterationCompositeOperation(str))
}

func (x IterationCompositeOperation) String() (string, bool) {
	switch x {
	case IterationCompositeOperation_REPLACE:
		return "replace", true
	case IterationCompositeOperation_ACCUMULATE:
		return "accumulate", true
	default:
		return "", false
	}
}

type TimelineRangeOffset struct {
	// RangeName is "TimelineRangeOffset.rangeName"
	//
	// Optional
	RangeName js.String
	// Offset is "TimelineRangeOffset.offset"
	//
	// Optional
	Offset CSSNumericValue

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TimelineRangeOffset with all fields set.
func (p TimelineRangeOffset) FromRef(ref js.Ref) TimelineRangeOffset {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TimelineRangeOffset in the application heap.
func (p TimelineRangeOffset) New() js.Ref {
	return bindings.TimelineRangeOffsetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TimelineRangeOffset) UpdateFrom(ref js.Ref) {
	bindings.TimelineRangeOffsetJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TimelineRangeOffset) Update(ref js.Ref) {
	bindings.TimelineRangeOffsetJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCSSKeywordValue(value js.String) CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(
		bindings.NewCSSKeywordValueByCSSKeywordValue(
			value.Ref()),
	)
}

type CSSKeywordValue struct {
	CSSStyleValue
}

func (this CSSKeywordValue) Once() CSSKeywordValue {
	this.Ref().Once()
	return this
}

func (this CSSKeywordValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSKeywordValue) FromRef(ref js.Ref) CSSKeywordValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSKeywordValue) Free() {
	this.Ref().Free()
}

// Value returns the value of property "CSSKeywordValue.value".
//
// The returned bool will be false if there is no such property.
func (this CSSKeywordValue) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeywordValueValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "CSSKeywordValue.value" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeywordValue) SetValue(val js.String) bool {
	return js.True == bindings.SetCSSKeywordValueValue(
		this.Ref(),
		val.Ref(),
	)
}

type OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String struct {
	ref js.Ref
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) Free() {
	x.ref.Free()
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) FromRef(ref js.Ref) OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String {
	return OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String{
		ref: ref,
	}
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) TimelineRangeOffset() TimelineRangeOffset {
	var ret TimelineRangeOffset
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) CSSKeywordValue() CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(x.ref)
}

func (x OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type KeyframeAnimationOptions struct {
	// Id is "KeyframeAnimationOptions.id"
	//
	// Optional, defaults to "".
	Id js.String
	// Timeline is "KeyframeAnimationOptions.timeline"
	//
	// Optional
	Timeline AnimationTimeline
	// IterationComposite is "KeyframeAnimationOptions.iterationComposite"
	//
	// Optional, defaults to "replace".
	IterationComposite IterationCompositeOperation
	// RangeStart is "KeyframeAnimationOptions.rangeStart"
	//
	// Optional, defaults to "normal".
	RangeStart OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String
	// RangeEnd is "KeyframeAnimationOptions.rangeEnd"
	//
	// Optional, defaults to "normal".
	RangeEnd OneOf_TimelineRangeOffset_CSSNumericValue_CSSKeywordValue_String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyframeAnimationOptions with all fields set.
func (p KeyframeAnimationOptions) FromRef(ref js.Ref) KeyframeAnimationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyframeAnimationOptions in the application heap.
func (p KeyframeAnimationOptions) New() js.Ref {
	return bindings.KeyframeAnimationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p KeyframeAnimationOptions) UpdateFrom(ref js.Ref) {
	bindings.KeyframeAnimationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p KeyframeAnimationOptions) Update(ref js.Ref) {
	bindings.KeyframeAnimationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Float64_KeyframeAnimationOptions struct {
	ref js.Ref
}

func (x OneOf_Float64_KeyframeAnimationOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_KeyframeAnimationOptions) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_KeyframeAnimationOptions) FromRef(ref js.Ref) OneOf_Float64_KeyframeAnimationOptions {
	return OneOf_Float64_KeyframeAnimationOptions{
		ref: ref,
	}
}

func (x OneOf_Float64_KeyframeAnimationOptions) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_KeyframeAnimationOptions) KeyframeAnimationOptions() KeyframeAnimationOptions {
	var ret KeyframeAnimationOptions
	ret.UpdateFrom(x.ref)
	return ret
}

type GetAnimationsOptions struct {
	// Subtree is "GetAnimationsOptions.subtree"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Subtree MUST be set to true to make this field effective.
	Subtree bool

	FFI_USE_Subtree bool // for Subtree.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAnimationsOptions with all fields set.
func (p GetAnimationsOptions) FromRef(ref js.Ref) GetAnimationsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAnimationsOptions in the application heap.
func (p GetAnimationsOptions) New() js.Ref {
	return bindings.GetAnimationsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GetAnimationsOptions) UpdateFrom(ref js.Ref) {
	bindings.GetAnimationsOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GetAnimationsOptions) Update(ref js.Ref) {
	bindings.GetAnimationsOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

const (
	Range_START_TO_START uint16 = 0
	Range_START_TO_END   uint16 = 1
	Range_END_TO_END     uint16 = 2
	Range_END_TO_START   uint16 = 3
)

type Range struct {
	AbstractRange
}

func (this Range) Once() Range {
	this.Ref().Once()
	return this
}

func (this Range) Ref() js.Ref {
	return this.AbstractRange.Ref()
}

func (this Range) FromRef(ref js.Ref) Range {
	this.AbstractRange = this.AbstractRange.FromRef(ref)
	return this
}

func (this Range) Free() {
	this.Ref().Free()
}

// CommonAncestorContainer returns the value of property "Range.commonAncestorContainer".
//
// The returned bool will be false if there is no such property.
func (this Range) CommonAncestorContainer() (Node, bool) {
	var _ok bool
	_ret := bindings.GetRangeCommonAncestorContainer(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// SetStart calls the method "Range.setStart".
//
// The returned bool will be false if there is no such method.
func (this Range) SetStart(node Node, offset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetStart(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStartFunc returns the method "Range.setStart".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetStartFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.RangeSetStartFunc(
			this.Ref(),
		),
	)
}

// SetEnd calls the method "Range.setEnd".
//
// The returned bool will be false if there is no such method.
func (this Range) SetEnd(node Node, offset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetEnd(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetEndFunc returns the method "Range.setEnd".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetEndFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.RangeSetEndFunc(
			this.Ref(),
		),
	)
}

// SetStartBefore calls the method "Range.setStartBefore".
//
// The returned bool will be false if there is no such method.
func (this Range) SetStartBefore(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetStartBefore(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStartBeforeFunc returns the method "Range.setStartBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetStartBeforeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetStartBeforeFunc(
			this.Ref(),
		),
	)
}

// SetStartAfter calls the method "Range.setStartAfter".
//
// The returned bool will be false if there is no such method.
func (this Range) SetStartAfter(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetStartAfter(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStartAfterFunc returns the method "Range.setStartAfter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetStartAfterFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetStartAfterFunc(
			this.Ref(),
		),
	)
}

// SetEndBefore calls the method "Range.setEndBefore".
//
// The returned bool will be false if there is no such method.
func (this Range) SetEndBefore(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetEndBefore(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetEndBeforeFunc returns the method "Range.setEndBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetEndBeforeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetEndBeforeFunc(
			this.Ref(),
		),
	)
}

// SetEndAfter calls the method "Range.setEndAfter".
//
// The returned bool will be false if there is no such method.
func (this Range) SetEndAfter(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSetEndAfter(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetEndAfterFunc returns the method "Range.setEndAfter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SetEndAfterFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetEndAfterFunc(
			this.Ref(),
		),
	)
}

// Collapse calls the method "Range.collapse".
//
// The returned bool will be false if there is no such method.
func (this Range) Collapse(toStart bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeCollapse(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(toStart)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CollapseFunc returns the method "Range.collapse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) CollapseFunc() (fn js.Func[func(toStart bool)]) {
	return fn.FromRef(
		bindings.RangeCollapseFunc(
			this.Ref(),
		),
	)
}

// Collapse1 calls the method "Range.collapse".
//
// The returned bool will be false if there is no such method.
func (this Range) Collapse1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeCollapse1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Collapse1Func returns the method "Range.collapse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) Collapse1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeCollapse1Func(
			this.Ref(),
		),
	)
}

// SelectNode calls the method "Range.selectNode".
//
// The returned bool will be false if there is no such method.
func (this Range) SelectNode(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSelectNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectNodeFunc returns the method "Range.selectNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SelectNodeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSelectNodeFunc(
			this.Ref(),
		),
	)
}

// SelectNodeContents calls the method "Range.selectNodeContents".
//
// The returned bool will be false if there is no such method.
func (this Range) SelectNodeContents(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSelectNodeContents(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectNodeContentsFunc returns the method "Range.selectNodeContents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SelectNodeContentsFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSelectNodeContentsFunc(
			this.Ref(),
		),
	)
}

// CompareBoundaryPoints calls the method "Range.compareBoundaryPoints".
//
// The returned bool will be false if there is no such method.
func (this Range) CompareBoundaryPoints(how uint16, sourceRange Range) (int16, bool) {
	var _ok bool
	_ret := bindings.CallRangeCompareBoundaryPoints(
		this.Ref(), js.Pointer(&_ok),
		uint32(how),
		sourceRange.Ref(),
	)

	return int16(_ret), _ok
}

// CompareBoundaryPointsFunc returns the method "Range.compareBoundaryPoints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) CompareBoundaryPointsFunc() (fn js.Func[func(how uint16, sourceRange Range) int16]) {
	return fn.FromRef(
		bindings.RangeCompareBoundaryPointsFunc(
			this.Ref(),
		),
	)
}

// DeleteContents calls the method "Range.deleteContents".
//
// The returned bool will be false if there is no such method.
func (this Range) DeleteContents() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeDeleteContents(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteContentsFunc returns the method "Range.deleteContents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) DeleteContentsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeDeleteContentsFunc(
			this.Ref(),
		),
	)
}

// ExtractContents calls the method "Range.extractContents".
//
// The returned bool will be false if there is no such method.
func (this Range) ExtractContents() (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallRangeExtractContents(
		this.Ref(), js.Pointer(&_ok),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// ExtractContentsFunc returns the method "Range.extractContents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) ExtractContentsFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeExtractContentsFunc(
			this.Ref(),
		),
	)
}

// CloneContents calls the method "Range.cloneContents".
//
// The returned bool will be false if there is no such method.
func (this Range) CloneContents() (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallRangeCloneContents(
		this.Ref(), js.Pointer(&_ok),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// CloneContentsFunc returns the method "Range.cloneContents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) CloneContentsFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeCloneContentsFunc(
			this.Ref(),
		),
	)
}

// InsertNode calls the method "Range.insertNode".
//
// The returned bool will be false if there is no such method.
func (this Range) InsertNode(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeInsertNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertNodeFunc returns the method "Range.insertNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) InsertNodeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeInsertNodeFunc(
			this.Ref(),
		),
	)
}

// SurroundContents calls the method "Range.surroundContents".
//
// The returned bool will be false if there is no such method.
func (this Range) SurroundContents(newParent Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeSurroundContents(
		this.Ref(), js.Pointer(&_ok),
		newParent.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SurroundContentsFunc returns the method "Range.surroundContents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) SurroundContentsFunc() (fn js.Func[func(newParent Node)]) {
	return fn.FromRef(
		bindings.RangeSurroundContentsFunc(
			this.Ref(),
		),
	)
}

// CloneRange calls the method "Range.cloneRange".
//
// The returned bool will be false if there is no such method.
func (this Range) CloneRange() (Range, bool) {
	var _ok bool
	_ret := bindings.CallRangeCloneRange(
		this.Ref(), js.Pointer(&_ok),
	)

	return Range{}.FromRef(_ret), _ok
}

// CloneRangeFunc returns the method "Range.cloneRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) CloneRangeFunc() (fn js.Func[func() Range]) {
	return fn.FromRef(
		bindings.RangeCloneRangeFunc(
			this.Ref(),
		),
	)
}

// Detach calls the method "Range.detach".
//
// The returned bool will be false if there is no such method.
func (this Range) Detach() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRangeDetach(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DetachFunc returns the method "Range.detach".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) DetachFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeDetachFunc(
			this.Ref(),
		),
	)
}

// IsPointInRange calls the method "Range.isPointInRange".
//
// The returned bool will be false if there is no such method.
func (this Range) IsPointInRange(node Node, offset uint32) (bool, bool) {
	var _ok bool
	_ret := bindings.CallRangeIsPointInRange(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	return _ret == js.True, _ok
}

// IsPointInRangeFunc returns the method "Range.isPointInRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) IsPointInRangeFunc() (fn js.Func[func(node Node, offset uint32) bool]) {
	return fn.FromRef(
		bindings.RangeIsPointInRangeFunc(
			this.Ref(),
		),
	)
}

// ComparePoint calls the method "Range.comparePoint".
//
// The returned bool will be false if there is no such method.
func (this Range) ComparePoint(node Node, offset uint32) (int16, bool) {
	var _ok bool
	_ret := bindings.CallRangeComparePoint(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	return int16(_ret), _ok
}

// ComparePointFunc returns the method "Range.comparePoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) ComparePointFunc() (fn js.Func[func(node Node, offset uint32) int16]) {
	return fn.FromRef(
		bindings.RangeComparePointFunc(
			this.Ref(),
		),
	)
}

// IntersectsNode calls the method "Range.intersectsNode".
//
// The returned bool will be false if there is no such method.
func (this Range) IntersectsNode(node Node) (bool, bool) {
	var _ok bool
	_ret := bindings.CallRangeIntersectsNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return _ret == js.True, _ok
}

// IntersectsNodeFunc returns the method "Range.intersectsNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) IntersectsNodeFunc() (fn js.Func[func(node Node) bool]) {
	return fn.FromRef(
		bindings.RangeIntersectsNodeFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "Range.toString".
//
// The returned bool will be false if there is no such method.
func (this Range) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallRangeToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "Range.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.RangeToStringFunc(
			this.Ref(),
		),
	)
}

// GetClientRects calls the method "Range.getClientRects".
//
// The returned bool will be false if there is no such method.
func (this Range) GetClientRects() (DOMRectList, bool) {
	var _ok bool
	_ret := bindings.CallRangeGetClientRects(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRectList{}.FromRef(_ret), _ok
}

// GetClientRectsFunc returns the method "Range.getClientRects".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) GetClientRectsFunc() (fn js.Func[func() DOMRectList]) {
	return fn.FromRef(
		bindings.RangeGetClientRectsFunc(
			this.Ref(),
		),
	)
}

// GetBoundingClientRect calls the method "Range.getBoundingClientRect".
//
// The returned bool will be false if there is no such method.
func (this Range) GetBoundingClientRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallRangeGetBoundingClientRect(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetBoundingClientRectFunc returns the method "Range.getBoundingClientRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) GetBoundingClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.RangeGetBoundingClientRectFunc(
			this.Ref(),
		),
	)
}

// CreateContextualFragment calls the method "Range.createContextualFragment".
//
// The returned bool will be false if there is no such method.
func (this Range) CreateContextualFragment(fragment js.String) (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallRangeCreateContextualFragment(
		this.Ref(), js.Pointer(&_ok),
		fragment.Ref(),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// CreateContextualFragmentFunc returns the method "Range.createContextualFragment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Range) CreateContextualFragmentFunc() (fn js.Func[func(fragment js.String) DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeCreateContextualFragmentFunc(
			this.Ref(),
		),
	)
}

type DOMTokenList struct {
	ref js.Ref
}

func (this DOMTokenList) Once() DOMTokenList {
	this.Ref().Once()
	return this
}

func (this DOMTokenList) Ref() js.Ref {
	return this.ref
}

func (this DOMTokenList) FromRef(ref js.Ref) DOMTokenList {
	this.ref = ref
	return this
}

func (this DOMTokenList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "DOMTokenList.length".
//
// The returned bool will be false if there is no such property.
func (this DOMTokenList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDOMTokenListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Value returns the value of property "DOMTokenList.value".
//
// The returned bool will be false if there is no such property.
func (this DOMTokenList) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDOMTokenListValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "DOMTokenList.value" to val.
//
// It returns false if the property cannot be set.
func (this DOMTokenList) SetValue(val js.String) bool {
	return js.True == bindings.SetDOMTokenListValue(
		this.Ref(),
		val.Ref(),
	)
}

// Item calls the method "DOMTokenList.item".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Item(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "DOMTokenList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.DOMTokenListItemFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "DOMTokenList.contains".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Contains(token js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListContains(
		this.Ref(), js.Pointer(&_ok),
		token.Ref(),
	)

	return _ret == js.True, _ok
}

// ContainsFunc returns the method "DOMTokenList.contains".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) ContainsFunc() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListContainsFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "DOMTokenList.add".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Add(tokens ...js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListAdd(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "DOMTokenList.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) AddFunc() (fn js.Func[func(tokens ...js.String)]) {
	return fn.FromRef(
		bindings.DOMTokenListAddFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "DOMTokenList.remove".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Remove(tokens ...js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListRemove(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "DOMTokenList.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) RemoveFunc() (fn js.Func[func(tokens ...js.String)]) {
	return fn.FromRef(
		bindings.DOMTokenListRemoveFunc(
			this.Ref(),
		),
	)
}

// Toggle calls the method "DOMTokenList.toggle".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Toggle(token js.String, force bool) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListToggle(
		this.Ref(), js.Pointer(&_ok),
		token.Ref(),
		js.Bool(bool(force)),
	)

	return _ret == js.True, _ok
}

// ToggleFunc returns the method "DOMTokenList.toggle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) ToggleFunc() (fn js.Func[func(token js.String, force bool) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListToggleFunc(
			this.Ref(),
		),
	)
}

// Toggle1 calls the method "DOMTokenList.toggle".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Toggle1(token js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListToggle1(
		this.Ref(), js.Pointer(&_ok),
		token.Ref(),
	)

	return _ret == js.True, _ok
}

// Toggle1Func returns the method "DOMTokenList.toggle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) Toggle1Func() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListToggle1Func(
			this.Ref(),
		),
	)
}

// Replace calls the method "DOMTokenList.replace".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Replace(token js.String, newToken js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListReplace(
		this.Ref(), js.Pointer(&_ok),
		token.Ref(),
		newToken.Ref(),
	)

	return _ret == js.True, _ok
}

// ReplaceFunc returns the method "DOMTokenList.replace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) ReplaceFunc() (fn js.Func[func(token js.String, newToken js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListReplaceFunc(
			this.Ref(),
		),
	)
}

// Supports calls the method "DOMTokenList.supports".
//
// The returned bool will be false if there is no such method.
func (this DOMTokenList) Supports(token js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMTokenListSupports(
		this.Ref(), js.Pointer(&_ok),
		token.Ref(),
	)

	return _ret == js.True, _ok
}

// SupportsFunc returns the method "DOMTokenList.supports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMTokenList) SupportsFunc() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListSupportsFunc(
			this.Ref(),
		),
	)
}

type NamedNodeMap struct {
	ref js.Ref
}

func (this NamedNodeMap) Once() NamedNodeMap {
	this.Ref().Once()
	return this
}

func (this NamedNodeMap) Ref() js.Ref {
	return this.ref
}

func (this NamedNodeMap) FromRef(ref js.Ref) NamedNodeMap {
	this.ref = ref
	return this
}

func (this NamedNodeMap) Free() {
	this.Ref().Free()
}

// Length returns the value of property "NamedNodeMap.length".
//
// The returned bool will be false if there is no such property.
func (this NamedNodeMap) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetNamedNodeMapLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "NamedNodeMap.item".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) Item(index uint32) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Attr{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "NamedNodeMap.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) ItemFunc() (fn js.Func[func(index uint32) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapItemFunc(
			this.Ref(),
		),
	)
}

// GetNamedItem calls the method "NamedNodeMap.getNamedItem".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) GetNamedItem(qualifiedName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapGetNamedItem(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// GetNamedItemFunc returns the method "NamedNodeMap.getNamedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) GetNamedItemFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapGetNamedItemFunc(
			this.Ref(),
		),
	)
}

// GetNamedItemNS calls the method "NamedNodeMap.getNamedItemNS".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) GetNamedItemNS(namespace js.String, localName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapGetNamedItemNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// GetNamedItemNSFunc returns the method "NamedNodeMap.getNamedItemNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) GetNamedItemNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapGetNamedItemNSFunc(
			this.Ref(),
		),
	)
}

// SetNamedItem calls the method "NamedNodeMap.setNamedItem".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) SetNamedItem(attr Attr) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapSetNamedItem(
		this.Ref(), js.Pointer(&_ok),
		attr.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// SetNamedItemFunc returns the method "NamedNodeMap.setNamedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) SetNamedItemFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapSetNamedItemFunc(
			this.Ref(),
		),
	)
}

// SetNamedItemNS calls the method "NamedNodeMap.setNamedItemNS".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) SetNamedItemNS(attr Attr) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapSetNamedItemNS(
		this.Ref(), js.Pointer(&_ok),
		attr.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// SetNamedItemNSFunc returns the method "NamedNodeMap.setNamedItemNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) SetNamedItemNSFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapSetNamedItemNSFunc(
			this.Ref(),
		),
	)
}

// RemoveNamedItem calls the method "NamedNodeMap.removeNamedItem".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) RemoveNamedItem(qualifiedName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapRemoveNamedItem(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// RemoveNamedItemFunc returns the method "NamedNodeMap.removeNamedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) RemoveNamedItemFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapRemoveNamedItemFunc(
			this.Ref(),
		),
	)
}

// RemoveNamedItemNS calls the method "NamedNodeMap.removeNamedItemNS".
//
// The returned bool will be false if there is no such method.
func (this NamedNodeMap) RemoveNamedItemNS(namespace js.String, localName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallNamedNodeMapRemoveNamedItemNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// RemoveNamedItemNSFunc returns the method "NamedNodeMap.removeNamedItemNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedNodeMap) RemoveNamedItemNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapRemoveNamedItemNSFunc(
			this.Ref(),
		),
	)
}

type Element struct {
	Node
}

func (this Element) Once() Element {
	this.Ref().Once()
	return this
}

func (this Element) Ref() js.Ref {
	return this.Node.Ref()
}

func (this Element) FromRef(ref js.Ref) Element {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this Element) Free() {
	this.Ref().Free()
}

// NamespaceURI returns the value of property "Element.namespaceURI".
//
// The returned bool will be false if there is no such property.
func (this Element) NamespaceURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementNamespaceURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Prefix returns the value of property "Element.prefix".
//
// The returned bool will be false if there is no such property.
func (this Element) Prefix() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementPrefix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LocalName returns the value of property "Element.localName".
//
// The returned bool will be false if there is no such property.
func (this Element) LocalName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementLocalName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// TagName returns the value of property "Element.tagName".
//
// The returned bool will be false if there is no such property.
func (this Element) TagName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementTagName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "Element.id".
//
// The returned bool will be false if there is no such property.
func (this Element) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id sets the value of property "Element.id" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetId(val js.String) bool {
	return js.True == bindings.SetElementId(
		this.Ref(),
		val.Ref(),
	)
}

// ClassName returns the value of property "Element.className".
//
// The returned bool will be false if there is no such property.
func (this Element) ClassName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementClassName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ClassName sets the value of property "Element.className" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetClassName(val js.String) bool {
	return js.True == bindings.SetElementClassName(
		this.Ref(),
		val.Ref(),
	)
}

// ClassList returns the value of property "Element.classList".
//
// The returned bool will be false if there is no such property.
func (this Element) ClassList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetElementClassList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Slot returns the value of property "Element.slot".
//
// The returned bool will be false if there is no such property.
func (this Element) Slot() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementSlot(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Slot sets the value of property "Element.slot" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetSlot(val js.String) bool {
	return js.True == bindings.SetElementSlot(
		this.Ref(),
		val.Ref(),
	)
}

// Attributes returns the value of property "Element.attributes".
//
// The returned bool will be false if there is no such property.
func (this Element) Attributes() (NamedNodeMap, bool) {
	var _ok bool
	_ret := bindings.GetElementAttributes(
		this.Ref(), js.Pointer(&_ok),
	)
	return NamedNodeMap{}.FromRef(_ret), _ok
}

// ShadowRoot returns the value of property "Element.shadowRoot".
//
// The returned bool will be false if there is no such property.
func (this Element) ShadowRoot() (ShadowRoot, bool) {
	var _ok bool
	_ret := bindings.GetElementShadowRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return ShadowRoot{}.FromRef(_ret), _ok
}

// ElementTiming returns the value of property "Element.elementTiming".
//
// The returned bool will be false if there is no such property.
func (this Element) ElementTiming() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementElementTiming(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ElementTiming sets the value of property "Element.elementTiming" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetElementTiming(val js.String) bool {
	return js.True == bindings.SetElementElementTiming(
		this.Ref(),
		val.Ref(),
	)
}

// Part returns the value of property "Element.part".
//
// The returned bool will be false if there is no such property.
func (this Element) Part() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetElementPart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// OuterHTML returns the value of property "Element.outerHTML".
//
// The returned bool will be false if there is no such property.
func (this Element) OuterHTML() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementOuterHTML(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OuterHTML sets the value of property "Element.outerHTML" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetOuterHTML(val js.String) bool {
	return js.True == bindings.SetElementOuterHTML(
		this.Ref(),
		val.Ref(),
	)
}

// ScrollTop returns the value of property "Element.scrollTop".
//
// The returned bool will be false if there is no such property.
func (this Element) ScrollTop() (float64, bool) {
	var _ok bool
	_ret := bindings.GetElementScrollTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ScrollTop sets the value of property "Element.scrollTop" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetScrollTop(val float64) bool {
	return js.True == bindings.SetElementScrollTop(
		this.Ref(),
		float64(val),
	)
}

// ScrollLeft returns the value of property "Element.scrollLeft".
//
// The returned bool will be false if there is no such property.
func (this Element) ScrollLeft() (float64, bool) {
	var _ok bool
	_ret := bindings.GetElementScrollLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ScrollLeft sets the value of property "Element.scrollLeft" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetScrollLeft(val float64) bool {
	return js.True == bindings.SetElementScrollLeft(
		this.Ref(),
		float64(val),
	)
}

// ScrollWidth returns the value of property "Element.scrollWidth".
//
// The returned bool will be false if there is no such property.
func (this Element) ScrollWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementScrollWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScrollHeight returns the value of property "Element.scrollHeight".
//
// The returned bool will be false if there is no such property.
func (this Element) ScrollHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementScrollHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientTop returns the value of property "Element.clientTop".
//
// The returned bool will be false if there is no such property.
func (this Element) ClientTop() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementClientTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientLeft returns the value of property "Element.clientLeft".
//
// The returned bool will be false if there is no such property.
func (this Element) ClientLeft() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementClientLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientWidth returns the value of property "Element.clientWidth".
//
// The returned bool will be false if there is no such property.
func (this Element) ClientWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementClientWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientHeight returns the value of property "Element.clientHeight".
//
// The returned bool will be false if there is no such property.
func (this Element) ClientHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetElementClientHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Role returns the value of property "Element.role".
//
// The returned bool will be false if there is no such property.
func (this Element) Role() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementRole(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Role sets the value of property "Element.role" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetRole(val js.String) bool {
	return js.True == bindings.SetElementRole(
		this.Ref(),
		val.Ref(),
	)
}

// AriaActiveDescendantElement returns the value of property "Element.ariaActiveDescendantElement".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaActiveDescendantElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaActiveDescendantElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// AriaActiveDescendantElement sets the value of property "Element.ariaActiveDescendantElement" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaActiveDescendantElement(val Element) bool {
	return js.True == bindings.SetElementAriaActiveDescendantElement(
		this.Ref(),
		val.Ref(),
	)
}

// AriaAtomic returns the value of property "Element.ariaAtomic".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaAtomic() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaAtomic(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaAtomic sets the value of property "Element.ariaAtomic" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaAtomic(val js.String) bool {
	return js.True == bindings.SetElementAriaAtomic(
		this.Ref(),
		val.Ref(),
	)
}

// AriaAutoComplete returns the value of property "Element.ariaAutoComplete".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaAutoComplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaAutoComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaAutoComplete sets the value of property "Element.ariaAutoComplete" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaAutoComplete(val js.String) bool {
	return js.True == bindings.SetElementAriaAutoComplete(
		this.Ref(),
		val.Ref(),
	)
}

// AriaBusy returns the value of property "Element.ariaBusy".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaBusy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaBusy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaBusy sets the value of property "Element.ariaBusy" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaBusy(val js.String) bool {
	return js.True == bindings.SetElementAriaBusy(
		this.Ref(),
		val.Ref(),
	)
}

// AriaChecked returns the value of property "Element.ariaChecked".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaChecked() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaChecked(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaChecked sets the value of property "Element.ariaChecked" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaChecked(val js.String) bool {
	return js.True == bindings.SetElementAriaChecked(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColCount returns the value of property "Element.ariaColCount".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaColCount() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaColCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColCount sets the value of property "Element.ariaColCount" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColCount(val js.String) bool {
	return js.True == bindings.SetElementAriaColCount(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColIndex returns the value of property "Element.ariaColIndex".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaColIndex() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaColIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColIndex sets the value of property "Element.ariaColIndex" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColIndex(val js.String) bool {
	return js.True == bindings.SetElementAriaColIndex(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColIndexText returns the value of property "Element.ariaColIndexText".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaColIndexText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaColIndexText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColIndexText sets the value of property "Element.ariaColIndexText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColIndexText(val js.String) bool {
	return js.True == bindings.SetElementAriaColIndexText(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColSpan returns the value of property "Element.ariaColSpan".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaColSpan() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaColSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColSpan sets the value of property "Element.ariaColSpan" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColSpan(val js.String) bool {
	return js.True == bindings.SetElementAriaColSpan(
		this.Ref(),
		val.Ref(),
	)
}

// AriaControlsElements returns the value of property "Element.ariaControlsElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaControlsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaControlsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaControlsElements sets the value of property "Element.ariaControlsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaControlsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaControlsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaCurrent returns the value of property "Element.ariaCurrent".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaCurrent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaCurrent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaCurrent sets the value of property "Element.ariaCurrent" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaCurrent(val js.String) bool {
	return js.True == bindings.SetElementAriaCurrent(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDescribedByElements returns the value of property "Element.ariaDescribedByElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaDescribedByElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaDescribedByElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaDescribedByElements sets the value of property "Element.ariaDescribedByElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDescribedByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaDescribedByElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDescription returns the value of property "Element.ariaDescription".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaDescription() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaDescription sets the value of property "Element.ariaDescription" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDescription(val js.String) bool {
	return js.True == bindings.SetElementAriaDescription(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDetailsElements returns the value of property "Element.ariaDetailsElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaDetailsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaDetailsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaDetailsElements sets the value of property "Element.ariaDetailsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDetailsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaDetailsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDisabled returns the value of property "Element.ariaDisabled".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaDisabled() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaDisabled sets the value of property "Element.ariaDisabled" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDisabled(val js.String) bool {
	return js.True == bindings.SetElementAriaDisabled(
		this.Ref(),
		val.Ref(),
	)
}

// AriaErrorMessageElements returns the value of property "Element.ariaErrorMessageElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaErrorMessageElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaErrorMessageElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaErrorMessageElements sets the value of property "Element.ariaErrorMessageElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaErrorMessageElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaErrorMessageElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaExpanded returns the value of property "Element.ariaExpanded".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaExpanded() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaExpanded(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaExpanded sets the value of property "Element.ariaExpanded" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaExpanded(val js.String) bool {
	return js.True == bindings.SetElementAriaExpanded(
		this.Ref(),
		val.Ref(),
	)
}

// AriaFlowToElements returns the value of property "Element.ariaFlowToElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaFlowToElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaFlowToElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaFlowToElements sets the value of property "Element.ariaFlowToElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaFlowToElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaFlowToElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaHasPopup returns the value of property "Element.ariaHasPopup".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaHasPopup() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaHasPopup(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaHasPopup sets the value of property "Element.ariaHasPopup" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaHasPopup(val js.String) bool {
	return js.True == bindings.SetElementAriaHasPopup(
		this.Ref(),
		val.Ref(),
	)
}

// AriaHidden returns the value of property "Element.ariaHidden".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaHidden() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaHidden(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaHidden sets the value of property "Element.ariaHidden" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaHidden(val js.String) bool {
	return js.True == bindings.SetElementAriaHidden(
		this.Ref(),
		val.Ref(),
	)
}

// AriaInvalid returns the value of property "Element.ariaInvalid".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaInvalid() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaInvalid(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaInvalid sets the value of property "Element.ariaInvalid" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaInvalid(val js.String) bool {
	return js.True == bindings.SetElementAriaInvalid(
		this.Ref(),
		val.Ref(),
	)
}

// AriaKeyShortcuts returns the value of property "Element.ariaKeyShortcuts".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaKeyShortcuts() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaKeyShortcuts(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaKeyShortcuts sets the value of property "Element.ariaKeyShortcuts" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaKeyShortcuts(val js.String) bool {
	return js.True == bindings.SetElementAriaKeyShortcuts(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLabel returns the value of property "Element.ariaLabel".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaLabel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLabel sets the value of property "Element.ariaLabel" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLabel(val js.String) bool {
	return js.True == bindings.SetElementAriaLabel(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLabelledByElements returns the value of property "Element.ariaLabelledByElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaLabelledByElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaLabelledByElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaLabelledByElements sets the value of property "Element.ariaLabelledByElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLabelledByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaLabelledByElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLevel returns the value of property "Element.ariaLevel".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaLevel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaLevel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLevel sets the value of property "Element.ariaLevel" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLevel(val js.String) bool {
	return js.True == bindings.SetElementAriaLevel(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLive returns the value of property "Element.ariaLive".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaLive() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaLive(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLive sets the value of property "Element.ariaLive" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLive(val js.String) bool {
	return js.True == bindings.SetElementAriaLive(
		this.Ref(),
		val.Ref(),
	)
}

// AriaModal returns the value of property "Element.ariaModal".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaModal() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaModal(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaModal sets the value of property "Element.ariaModal" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaModal(val js.String) bool {
	return js.True == bindings.SetElementAriaModal(
		this.Ref(),
		val.Ref(),
	)
}

// AriaMultiLine returns the value of property "Element.ariaMultiLine".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaMultiLine() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaMultiLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaMultiLine sets the value of property "Element.ariaMultiLine" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaMultiLine(val js.String) bool {
	return js.True == bindings.SetElementAriaMultiLine(
		this.Ref(),
		val.Ref(),
	)
}

// AriaMultiSelectable returns the value of property "Element.ariaMultiSelectable".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaMultiSelectable() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaMultiSelectable(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaMultiSelectable sets the value of property "Element.ariaMultiSelectable" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaMultiSelectable(val js.String) bool {
	return js.True == bindings.SetElementAriaMultiSelectable(
		this.Ref(),
		val.Ref(),
	)
}

// AriaOrientation returns the value of property "Element.ariaOrientation".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaOrientation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaOrientation sets the value of property "Element.ariaOrientation" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaOrientation(val js.String) bool {
	return js.True == bindings.SetElementAriaOrientation(
		this.Ref(),
		val.Ref(),
	)
}

// AriaOwnsElements returns the value of property "Element.ariaOwnsElements".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaOwnsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementAriaOwnsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaOwnsElements sets the value of property "Element.ariaOwnsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaOwnsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaOwnsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPlaceholder returns the value of property "Element.ariaPlaceholder".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaPlaceholder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaPlaceholder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPlaceholder sets the value of property "Element.ariaPlaceholder" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPlaceholder(val js.String) bool {
	return js.True == bindings.SetElementAriaPlaceholder(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPosInSet returns the value of property "Element.ariaPosInSet".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaPosInSet() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaPosInSet(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPosInSet sets the value of property "Element.ariaPosInSet" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPosInSet(val js.String) bool {
	return js.True == bindings.SetElementAriaPosInSet(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPressed returns the value of property "Element.ariaPressed".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaPressed() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaPressed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPressed sets the value of property "Element.ariaPressed" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPressed(val js.String) bool {
	return js.True == bindings.SetElementAriaPressed(
		this.Ref(),
		val.Ref(),
	)
}

// AriaReadOnly returns the value of property "Element.ariaReadOnly".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaReadOnly() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaReadOnly(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaReadOnly sets the value of property "Element.ariaReadOnly" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaReadOnly(val js.String) bool {
	return js.True == bindings.SetElementAriaReadOnly(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRequired returns the value of property "Element.ariaRequired".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRequired() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRequired(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRequired sets the value of property "Element.ariaRequired" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRequired(val js.String) bool {
	return js.True == bindings.SetElementAriaRequired(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRoleDescription returns the value of property "Element.ariaRoleDescription".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRoleDescription() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRoleDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRoleDescription sets the value of property "Element.ariaRoleDescription" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRoleDescription(val js.String) bool {
	return js.True == bindings.SetElementAriaRoleDescription(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowCount returns the value of property "Element.ariaRowCount".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRowCount() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRowCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowCount sets the value of property "Element.ariaRowCount" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowCount(val js.String) bool {
	return js.True == bindings.SetElementAriaRowCount(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowIndex returns the value of property "Element.ariaRowIndex".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRowIndex() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRowIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowIndex sets the value of property "Element.ariaRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowIndex(val js.String) bool {
	return js.True == bindings.SetElementAriaRowIndex(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowIndexText returns the value of property "Element.ariaRowIndexText".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRowIndexText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRowIndexText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowIndexText sets the value of property "Element.ariaRowIndexText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowIndexText(val js.String) bool {
	return js.True == bindings.SetElementAriaRowIndexText(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowSpan returns the value of property "Element.ariaRowSpan".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaRowSpan() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaRowSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowSpan sets the value of property "Element.ariaRowSpan" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowSpan(val js.String) bool {
	return js.True == bindings.SetElementAriaRowSpan(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSelected returns the value of property "Element.ariaSelected".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaSelected() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSelected sets the value of property "Element.ariaSelected" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSelected(val js.String) bool {
	return js.True == bindings.SetElementAriaSelected(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSetSize returns the value of property "Element.ariaSetSize".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaSetSize() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaSetSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSetSize sets the value of property "Element.ariaSetSize" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSetSize(val js.String) bool {
	return js.True == bindings.SetElementAriaSetSize(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSort returns the value of property "Element.ariaSort".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaSort() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaSort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSort sets the value of property "Element.ariaSort" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSort(val js.String) bool {
	return js.True == bindings.SetElementAriaSort(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueMax returns the value of property "Element.ariaValueMax".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaValueMax() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaValueMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueMax sets the value of property "Element.ariaValueMax" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueMax(val js.String) bool {
	return js.True == bindings.SetElementAriaValueMax(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueMin returns the value of property "Element.ariaValueMin".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaValueMin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaValueMin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueMin sets the value of property "Element.ariaValueMin" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueMin(val js.String) bool {
	return js.True == bindings.SetElementAriaValueMin(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueNow returns the value of property "Element.ariaValueNow".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaValueNow() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaValueNow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueNow sets the value of property "Element.ariaValueNow" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueNow(val js.String) bool {
	return js.True == bindings.SetElementAriaValueNow(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueText returns the value of property "Element.ariaValueText".
//
// The returned bool will be false if there is no such property.
func (this Element) AriaValueText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementAriaValueText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueText sets the value of property "Element.ariaValueText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueText(val js.String) bool {
	return js.True == bindings.SetElementAriaValueText(
		this.Ref(),
		val.Ref(),
	)
}

// InnerHTML returns the value of property "Element.innerHTML".
//
// The returned bool will be false if there is no such property.
func (this Element) InnerHTML() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInnerHTML(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InnerHTML sets the value of property "Element.innerHTML" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetInnerHTML(val js.String) bool {
	return js.True == bindings.SetElementInnerHTML(
		this.Ref(),
		val.Ref(),
	)
}

// Children returns the value of property "Element.children".
//
// The returned bool will be false if there is no such property.
func (this Element) Children() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetElementChildren(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// FirstElementChild returns the value of property "Element.firstElementChild".
//
// The returned bool will be false if there is no such property.
func (this Element) FirstElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementFirstElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// LastElementChild returns the value of property "Element.lastElementChild".
//
// The returned bool will be false if there is no such property.
func (this Element) LastElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementLastElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ChildElementCount returns the value of property "Element.childElementCount".
//
// The returned bool will be false if there is no such property.
func (this Element) ChildElementCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetElementChildElementCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// PreviousElementSibling returns the value of property "Element.previousElementSibling".
//
// The returned bool will be false if there is no such property.
func (this Element) PreviousElementSibling() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementPreviousElementSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// NextElementSibling returns the value of property "Element.nextElementSibling".
//
// The returned bool will be false if there is no such property.
func (this Element) NextElementSibling() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementNextElementSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// AssignedSlot returns the value of property "Element.assignedSlot".
//
// The returned bool will be false if there is no such property.
func (this Element) AssignedSlot() (HTMLSlotElement, bool) {
	var _ok bool
	_ret := bindings.GetElementAssignedSlot(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLSlotElement{}.FromRef(_ret), _ok
}

// RegionOverset returns the value of property "Element.regionOverset".
//
// The returned bool will be false if there is no such property.
func (this Element) RegionOverset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementRegionOverset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// HasAttributes calls the method "Element.hasAttributes".
//
// The returned bool will be false if there is no such method.
func (this Element) HasAttributes() (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementHasAttributes(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// HasAttributesFunc returns the method "Element.hasAttributes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) HasAttributesFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributesFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNames calls the method "Element.getAttributeNames".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAttributeNames() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallElementGetAttributeNames(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// GetAttributeNamesFunc returns the method "Element.getAttributeNames".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAttributeNamesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNamesFunc(
			this.Ref(),
		),
	)
}

// GetAttribute calls the method "Element.getAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAttribute(qualifiedName js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallElementGetAttribute(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAttributeFunc returns the method "Element.getAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAttributeFunc() (fn js.Func[func(qualifiedName js.String) js.String]) {
	return fn.FromRef(
		bindings.ElementGetAttributeFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNS calls the method "Element.getAttributeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAttributeNS(namespace js.String, localName js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallElementGetAttributeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAttributeNSFunc returns the method "Element.getAttributeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) js.String]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNSFunc(
			this.Ref(),
		),
	)
}

// SetAttribute calls the method "Element.setAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) SetAttribute(qualifiedName js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementSetAttribute(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetAttributeFunc returns the method "Element.setAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetAttributeFunc() (fn js.Func[func(qualifiedName js.String, value js.String)]) {
	return fn.FromRef(
		bindings.ElementSetAttributeFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNS calls the method "Element.setAttributeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) SetAttributeNS(namespace js.String, qualifiedName js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementSetAttributeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetAttributeNSFunc returns the method "Element.setAttributeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetAttributeNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, value js.String)]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNSFunc(
			this.Ref(),
		),
	)
}

// RemoveAttribute calls the method "Element.removeAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) RemoveAttribute(qualifiedName js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementRemoveAttribute(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveAttributeFunc returns the method "Element.removeAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RemoveAttributeFunc() (fn js.Func[func(qualifiedName js.String)]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeFunc(
			this.Ref(),
		),
	)
}

// RemoveAttributeNS calls the method "Element.removeAttributeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) RemoveAttributeNS(namespace js.String, localName js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementRemoveAttributeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveAttributeNSFunc returns the method "Element.removeAttributeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RemoveAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String)]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeNSFunc(
			this.Ref(),
		),
	)
}

// ToggleAttribute calls the method "Element.toggleAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) ToggleAttribute(qualifiedName js.String, force bool) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementToggleAttribute(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
		js.Bool(bool(force)),
	)

	return _ret == js.True, _ok
}

// ToggleAttributeFunc returns the method "Element.toggleAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ToggleAttributeFunc() (fn js.Func[func(qualifiedName js.String, force bool) bool]) {
	return fn.FromRef(
		bindings.ElementToggleAttributeFunc(
			this.Ref(),
		),
	)
}

// ToggleAttribute1 calls the method "Element.toggleAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) ToggleAttribute1(qualifiedName js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementToggleAttribute1(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return _ret == js.True, _ok
}

// ToggleAttribute1Func returns the method "Element.toggleAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ToggleAttribute1Func() (fn js.Func[func(qualifiedName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementToggleAttribute1Func(
			this.Ref(),
		),
	)
}

// HasAttribute calls the method "Element.hasAttribute".
//
// The returned bool will be false if there is no such method.
func (this Element) HasAttribute(qualifiedName js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementHasAttribute(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return _ret == js.True, _ok
}

// HasAttributeFunc returns the method "Element.hasAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) HasAttributeFunc() (fn js.Func[func(qualifiedName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributeFunc(
			this.Ref(),
		),
	)
}

// HasAttributeNS calls the method "Element.hasAttributeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) HasAttributeNS(namespace js.String, localName js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementHasAttributeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return _ret == js.True, _ok
}

// HasAttributeNSFunc returns the method "Element.hasAttributeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) HasAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributeNSFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNode calls the method "Element.getAttributeNode".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAttributeNode(qualifiedName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallElementGetAttributeNode(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// GetAttributeNodeFunc returns the method "Element.getAttributeNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAttributeNodeFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNodeNS calls the method "Element.getAttributeNodeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAttributeNodeNS(namespace js.String, localName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallElementGetAttributeNodeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// GetAttributeNodeNSFunc returns the method "Element.getAttributeNodeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAttributeNodeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNodeNSFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNode calls the method "Element.setAttributeNode".
//
// The returned bool will be false if there is no such method.
func (this Element) SetAttributeNode(attr Attr) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallElementSetAttributeNode(
		this.Ref(), js.Pointer(&_ok),
		attr.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// SetAttributeNodeFunc returns the method "Element.setAttributeNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetAttributeNodeFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNodeNS calls the method "Element.setAttributeNodeNS".
//
// The returned bool will be false if there is no such method.
func (this Element) SetAttributeNodeNS(attr Attr) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallElementSetAttributeNodeNS(
		this.Ref(), js.Pointer(&_ok),
		attr.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// SetAttributeNodeNSFunc returns the method "Element.setAttributeNodeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetAttributeNodeNSFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNodeNSFunc(
			this.Ref(),
		),
	)
}

// RemoveAttributeNode calls the method "Element.removeAttributeNode".
//
// The returned bool will be false if there is no such method.
func (this Element) RemoveAttributeNode(attr Attr) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallElementRemoveAttributeNode(
		this.Ref(), js.Pointer(&_ok),
		attr.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// RemoveAttributeNodeFunc returns the method "Element.removeAttributeNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RemoveAttributeNodeFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// AttachShadow calls the method "Element.attachShadow".
//
// The returned bool will be false if there is no such method.
func (this Element) AttachShadow(init ShadowRootInit) (ShadowRoot, bool) {
	var _ok bool
	_ret := bindings.CallElementAttachShadow(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return ShadowRoot{}.FromRef(_ret), _ok
}

// AttachShadowFunc returns the method "Element.attachShadow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) AttachShadowFunc() (fn js.Func[func(init ShadowRootInit) ShadowRoot]) {
	return fn.FromRef(
		bindings.ElementAttachShadowFunc(
			this.Ref(),
		),
	)
}

// Closest calls the method "Element.closest".
//
// The returned bool will be false if there is no such method.
func (this Element) Closest(selectors js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallElementClosest(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// ClosestFunc returns the method "Element.closest".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ClosestFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.ElementClosestFunc(
			this.Ref(),
		),
	)
}

// Matches calls the method "Element.matches".
//
// The returned bool will be false if there is no such method.
func (this Element) Matches(selectors js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementMatches(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return _ret == js.True, _ok
}

// MatchesFunc returns the method "Element.matches".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) MatchesFunc() (fn js.Func[func(selectors js.String) bool]) {
	return fn.FromRef(
		bindings.ElementMatchesFunc(
			this.Ref(),
		),
	)
}

// WebkitMatchesSelector calls the method "Element.webkitMatchesSelector".
//
// The returned bool will be false if there is no such method.
func (this Element) WebkitMatchesSelector(selectors js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementWebkitMatchesSelector(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return _ret == js.True, _ok
}

// WebkitMatchesSelectorFunc returns the method "Element.webkitMatchesSelector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) WebkitMatchesSelectorFunc() (fn js.Func[func(selectors js.String) bool]) {
	return fn.FromRef(
		bindings.ElementWebkitMatchesSelectorFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagName calls the method "Element.getElementsByTagName".
//
// The returned bool will be false if there is no such method.
func (this Element) GetElementsByTagName(qualifiedName js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallElementGetElementsByTagName(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByTagNameFunc returns the method "Element.getElementsByTagName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetElementsByTagNameFunc() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByTagNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagNameNS calls the method "Element.getElementsByTagNameNS".
//
// The returned bool will be false if there is no such method.
func (this Element) GetElementsByTagNameNS(namespace js.String, localName js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallElementGetElementsByTagNameNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByTagNameNSFunc returns the method "Element.getElementsByTagNameNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetElementsByTagNameNSFunc() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByTagNameNSFunc(
			this.Ref(),
		),
	)
}

// GetElementsByClassName calls the method "Element.getElementsByClassName".
//
// The returned bool will be false if there is no such method.
func (this Element) GetElementsByClassName(classNames js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallElementGetElementsByClassName(
		this.Ref(), js.Pointer(&_ok),
		classNames.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByClassNameFunc returns the method "Element.getElementsByClassName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetElementsByClassNameFunc() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByClassNameFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentElement calls the method "Element.insertAdjacentElement".
//
// The returned bool will be false if there is no such method.
func (this Element) InsertAdjacentElement(where js.String, element Element) (Element, bool) {
	var _ok bool
	_ret := bindings.CallElementInsertAdjacentElement(
		this.Ref(), js.Pointer(&_ok),
		where.Ref(),
		element.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// InsertAdjacentElementFunc returns the method "Element.insertAdjacentElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) InsertAdjacentElementFunc() (fn js.Func[func(where js.String, element Element) Element]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentElementFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentText calls the method "Element.insertAdjacentText".
//
// The returned bool will be false if there is no such method.
func (this Element) InsertAdjacentText(where js.String, data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInsertAdjacentText(
		this.Ref(), js.Pointer(&_ok),
		where.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertAdjacentTextFunc returns the method "Element.insertAdjacentText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) InsertAdjacentTextFunc() (fn js.Func[func(where js.String, data js.String)]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentTextFunc(
			this.Ref(),
		),
	)
}

// RequestFullscreen calls the method "Element.requestFullscreen".
//
// The returned bool will be false if there is no such method.
func (this Element) RequestFullscreen(options FullscreenOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallElementRequestFullscreen(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RequestFullscreenFunc returns the method "Element.requestFullscreen".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RequestFullscreenFunc() (fn js.Func[func(options FullscreenOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ElementRequestFullscreenFunc(
			this.Ref(),
		),
	)
}

// RequestFullscreen1 calls the method "Element.requestFullscreen".
//
// The returned bool will be false if there is no such method.
func (this Element) RequestFullscreen1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallElementRequestFullscreen1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RequestFullscreen1Func returns the method "Element.requestFullscreen".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RequestFullscreen1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ElementRequestFullscreen1Func(
			this.Ref(),
		),
	)
}

// RequestPointerLock calls the method "Element.requestPointerLock".
//
// The returned bool will be false if there is no such method.
func (this Element) RequestPointerLock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementRequestPointerLock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestPointerLockFunc returns the method "Element.requestPointerLock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RequestPointerLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementRequestPointerLockFunc(
			this.Ref(),
		),
	)
}

// ComputedStyleMap calls the method "Element.computedStyleMap".
//
// The returned bool will be false if there is no such method.
func (this Element) ComputedStyleMap() (StylePropertyMapReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallElementComputedStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)

	return StylePropertyMapReadOnly{}.FromRef(_ret), _ok
}

// ComputedStyleMapFunc returns the method "Element.computedStyleMap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ComputedStyleMapFunc() (fn js.Func[func() StylePropertyMapReadOnly]) {
	return fn.FromRef(
		bindings.ElementComputedStyleMapFunc(
			this.Ref(),
		),
	)
}

// GetSpatialNavigationContainer calls the method "Element.getSpatialNavigationContainer".
//
// The returned bool will be false if there is no such method.
func (this Element) GetSpatialNavigationContainer() (Node, bool) {
	var _ok bool
	_ret := bindings.CallElementGetSpatialNavigationContainer(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// GetSpatialNavigationContainerFunc returns the method "Element.getSpatialNavigationContainer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetSpatialNavigationContainerFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.ElementGetSpatialNavigationContainerFunc(
			this.Ref(),
		),
	)
}

// FocusableAreas calls the method "Element.focusableAreas".
//
// The returned bool will be false if there is no such method.
func (this Element) FocusableAreas(option FocusableAreasOption) (js.Array[Node], bool) {
	var _ok bool
	_ret := bindings.CallElementFocusableAreas(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&option),
	)

	return js.Array[Node]{}.FromRef(_ret), _ok
}

// FocusableAreasFunc returns the method "Element.focusableAreas".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) FocusableAreasFunc() (fn js.Func[func(option FocusableAreasOption) js.Array[Node]]) {
	return fn.FromRef(
		bindings.ElementFocusableAreasFunc(
			this.Ref(),
		),
	)
}

// FocusableAreas1 calls the method "Element.focusableAreas".
//
// The returned bool will be false if there is no such method.
func (this Element) FocusableAreas1() (js.Array[Node], bool) {
	var _ok bool
	_ret := bindings.CallElementFocusableAreas1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Node]{}.FromRef(_ret), _ok
}

// FocusableAreas1Func returns the method "Element.focusableAreas".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) FocusableAreas1Func() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.ElementFocusableAreas1Func(
			this.Ref(),
		),
	)
}

// SpatialNavigationSearch calls the method "Element.spatialNavigationSearch".
//
// The returned bool will be false if there is no such method.
func (this Element) SpatialNavigationSearch(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) (Node, bool) {
	var _ok bool
	_ret := bindings.CallElementSpatialNavigationSearch(
		this.Ref(), js.Pointer(&_ok),
		uint32(dir),
		js.Pointer(&options),
	)

	return Node{}.FromRef(_ret), _ok
}

// SpatialNavigationSearchFunc returns the method "Element.spatialNavigationSearch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SpatialNavigationSearchFunc() (fn js.Func[func(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) Node]) {
	return fn.FromRef(
		bindings.ElementSpatialNavigationSearchFunc(
			this.Ref(),
		),
	)
}

// SpatialNavigationSearch1 calls the method "Element.spatialNavigationSearch".
//
// The returned bool will be false if there is no such method.
func (this Element) SpatialNavigationSearch1(dir SpatialNavigationDirection) (Node, bool) {
	var _ok bool
	_ret := bindings.CallElementSpatialNavigationSearch1(
		this.Ref(), js.Pointer(&_ok),
		uint32(dir),
	)

	return Node{}.FromRef(_ret), _ok
}

// SpatialNavigationSearch1Func returns the method "Element.spatialNavigationSearch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SpatialNavigationSearch1Func() (fn js.Func[func(dir SpatialNavigationDirection) Node]) {
	return fn.FromRef(
		bindings.ElementSpatialNavigationSearch1Func(
			this.Ref(),
		),
	)
}

// SetPointerCapture calls the method "Element.setPointerCapture".
//
// The returned bool will be false if there is no such method.
func (this Element) SetPointerCapture(pointerId int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementSetPointerCapture(
		this.Ref(), js.Pointer(&_ok),
		int32(pointerId),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPointerCaptureFunc returns the method "Element.setPointerCapture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetPointerCaptureFunc() (fn js.Func[func(pointerId int32)]) {
	return fn.FromRef(
		bindings.ElementSetPointerCaptureFunc(
			this.Ref(),
		),
	)
}

// ReleasePointerCapture calls the method "Element.releasePointerCapture".
//
// The returned bool will be false if there is no such method.
func (this Element) ReleasePointerCapture(pointerId int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementReleasePointerCapture(
		this.Ref(), js.Pointer(&_ok),
		int32(pointerId),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleasePointerCaptureFunc returns the method "Element.releasePointerCapture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ReleasePointerCaptureFunc() (fn js.Func[func(pointerId int32)]) {
	return fn.FromRef(
		bindings.ElementReleasePointerCaptureFunc(
			this.Ref(),
		),
	)
}

// HasPointerCapture calls the method "Element.hasPointerCapture".
//
// The returned bool will be false if there is no such method.
func (this Element) HasPointerCapture(pointerId int32) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementHasPointerCapture(
		this.Ref(), js.Pointer(&_ok),
		int32(pointerId),
	)

	return _ret == js.True, _ok
}

// HasPointerCaptureFunc returns the method "Element.hasPointerCapture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) HasPointerCaptureFunc() (fn js.Func[func(pointerId int32) bool]) {
	return fn.FromRef(
		bindings.ElementHasPointerCaptureFunc(
			this.Ref(),
		),
	)
}

// Pseudo calls the method "Element.pseudo".
//
// The returned bool will be false if there is no such method.
func (this Element) Pseudo(typ js.String) (CSSPseudoElement, bool) {
	var _ok bool
	_ret := bindings.CallElementPseudo(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return CSSPseudoElement{}.FromRef(_ret), _ok
}

// PseudoFunc returns the method "Element.pseudo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) PseudoFunc() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	return fn.FromRef(
		bindings.ElementPseudoFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentHTML calls the method "Element.insertAdjacentHTML".
//
// The returned bool will be false if there is no such method.
func (this Element) InsertAdjacentHTML(position js.String, text js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInsertAdjacentHTML(
		this.Ref(), js.Pointer(&_ok),
		position.Ref(),
		text.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertAdjacentHTMLFunc returns the method "Element.insertAdjacentHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) InsertAdjacentHTMLFunc() (fn js.Func[func(position js.String, text js.String)]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentHTMLFunc(
			this.Ref(),
		),
	)
}

// SetHTML calls the method "Element.setHTML".
//
// The returned bool will be false if there is no such method.
func (this Element) SetHTML(input js.String, options SetHTMLOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementSetHTML(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetHTMLFunc returns the method "Element.setHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetHTMLFunc() (fn js.Func[func(input js.String, options SetHTMLOptions)]) {
	return fn.FromRef(
		bindings.ElementSetHTMLFunc(
			this.Ref(),
		),
	)
}

// SetHTML1 calls the method "Element.setHTML".
//
// The returned bool will be false if there is no such method.
func (this Element) SetHTML1(input js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementSetHTML1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetHTML1Func returns the method "Element.setHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) SetHTML1Func() (fn js.Func[func(input js.String)]) {
	return fn.FromRef(
		bindings.ElementSetHTML1Func(
			this.Ref(),
		),
	)
}

// GetClientRects calls the method "Element.getClientRects".
//
// The returned bool will be false if there is no such method.
func (this Element) GetClientRects() (DOMRectList, bool) {
	var _ok bool
	_ret := bindings.CallElementGetClientRects(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRectList{}.FromRef(_ret), _ok
}

// GetClientRectsFunc returns the method "Element.getClientRects".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetClientRectsFunc() (fn js.Func[func() DOMRectList]) {
	return fn.FromRef(
		bindings.ElementGetClientRectsFunc(
			this.Ref(),
		),
	)
}

// GetBoundingClientRect calls the method "Element.getBoundingClientRect".
//
// The returned bool will be false if there is no such method.
func (this Element) GetBoundingClientRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallElementGetBoundingClientRect(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetBoundingClientRectFunc returns the method "Element.getBoundingClientRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetBoundingClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.ElementGetBoundingClientRectFunc(
			this.Ref(),
		),
	)
}

// CheckVisibility calls the method "Element.checkVisibility".
//
// The returned bool will be false if there is no such method.
func (this Element) CheckVisibility(options CheckVisibilityOptions) (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementCheckVisibility(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return _ret == js.True, _ok
}

// CheckVisibilityFunc returns the method "Element.checkVisibility".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) CheckVisibilityFunc() (fn js.Func[func(options CheckVisibilityOptions) bool]) {
	return fn.FromRef(
		bindings.ElementCheckVisibilityFunc(
			this.Ref(),
		),
	)
}

// CheckVisibility1 calls the method "Element.checkVisibility".
//
// The returned bool will be false if there is no such method.
func (this Element) CheckVisibility1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementCheckVisibility1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckVisibility1Func returns the method "Element.checkVisibility".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) CheckVisibility1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementCheckVisibility1Func(
			this.Ref(),
		),
	)
}

// ScrollIntoView calls the method "Element.scrollIntoView".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollIntoView(arg OneOf_Bool_ScrollIntoViewOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollIntoView(
		this.Ref(), js.Pointer(&_ok),
		arg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollIntoViewFunc returns the method "Element.scrollIntoView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollIntoViewFunc() (fn js.Func[func(arg OneOf_Bool_ScrollIntoViewOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollIntoViewFunc(
			this.Ref(),
		),
	)
}

// ScrollIntoView1 calls the method "Element.scrollIntoView".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollIntoView1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollIntoView1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollIntoView1Func returns the method "Element.scrollIntoView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollIntoView1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollIntoView1Func(
			this.Ref(),
		),
	)
}

// Scroll calls the method "Element.scroll".
//
// The returned bool will be false if there is no such method.
func (this Element) Scroll(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScroll(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollFunc returns the method "Element.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollFunc(
			this.Ref(),
		),
	)
}

// Scroll1 calls the method "Element.scroll".
//
// The returned bool will be false if there is no such method.
func (this Element) Scroll1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScroll1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Scroll1Func returns the method "Element.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) Scroll1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScroll1Func(
			this.Ref(),
		),
	)
}

// Scroll2 calls the method "Element.scroll".
//
// The returned bool will be false if there is no such method.
func (this Element) Scroll2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScroll2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Scroll2Func returns the method "Element.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) Scroll2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScroll2Func(
			this.Ref(),
		),
	)
}

// ScrollTo calls the method "Element.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollTo(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollTo(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollToFunc returns the method "Element.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollToFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollToFunc(
			this.Ref(),
		),
	)
}

// ScrollTo1 calls the method "Element.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollTo1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollTo1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollTo1Func returns the method "Element.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollTo1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollTo1Func(
			this.Ref(),
		),
	)
}

// ScrollTo2 calls the method "Element.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollTo2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollTo2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollTo2Func returns the method "Element.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollTo2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScrollTo2Func(
			this.Ref(),
		),
	)
}

// ScrollBy calls the method "Element.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollBy(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollBy(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollByFunc returns the method "Element.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollByFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollByFunc(
			this.Ref(),
		),
	)
}

// ScrollBy1 calls the method "Element.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollBy1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollBy1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollBy1Func returns the method "Element.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollBy1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollBy1Func(
			this.Ref(),
		),
	)
}

// ScrollBy2 calls the method "Element.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Element) ScrollBy2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementScrollBy2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollBy2Func returns the method "Element.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ScrollBy2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScrollBy2Func(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Element.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Element) GetBoxQuads(options BoxQuadOptions) (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallElementGetBoxQuads(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuadsFunc returns the method "Element.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.ElementGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Element.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Element) GetBoxQuads1() (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallElementGetBoxQuads1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuads1Func returns the method "Element.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.ElementGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Element.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNodeFunc returns the method "Element.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Element.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNode1Func returns the method "Element.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Element.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertRectFromNode(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNodeFunc returns the method "Element.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Element.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNode1Func returns the method "Element.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Element.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertPointFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNodeFunc returns the method "Element.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.ElementConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Element.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Element) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNode1Func returns the method "Element.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.ElementConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// Prepend calls the method "Element.prepend".
//
// The returned bool will be false if there is no such method.
func (this Element) Prepend(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementPrepend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PrependFunc returns the method "Element.prepend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementPrependFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "Element.append".
//
// The returned bool will be false if there is no such method.
func (this Element) Append(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementAppend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "Element.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementAppendFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "Element.replaceChildren".
//
// The returned bool will be false if there is no such method.
func (this Element) ReplaceChildren(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementReplaceChildren(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceChildrenFunc returns the method "Element.replaceChildren".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "Element.querySelector".
//
// The returned bool will be false if there is no such method.
func (this Element) QuerySelector(selectors js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallElementQuerySelector(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// QuerySelectorFunc returns the method "Element.querySelector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.ElementQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "Element.querySelectorAll".
//
// The returned bool will be false if there is no such method.
func (this Element) QuerySelectorAll(selectors js.String) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallElementQuerySelectorAll(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// QuerySelectorAllFunc returns the method "Element.querySelectorAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.ElementQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "Element.before".
//
// The returned bool will be false if there is no such method.
func (this Element) Before(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementBefore(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeforeFunc returns the method "Element.before".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementBeforeFunc(
			this.Ref(),
		),
	)
}

// After calls the method "Element.after".
//
// The returned bool will be false if there is no such method.
func (this Element) After(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementAfter(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AfterFunc returns the method "Element.after".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementAfterFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "Element.replaceWith".
//
// The returned bool will be false if there is no such method.
func (this Element) ReplaceWith(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementReplaceWith(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceWithFunc returns the method "Element.replaceWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementReplaceWithFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "Element.remove".
//
// The returned bool will be false if there is no such method.
func (this Element) Remove() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "Element.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementRemoveFunc(
			this.Ref(),
		),
	)
}

// Animate calls the method "Element.animate".
//
// The returned bool will be false if there is no such method.
func (this Element) Animate(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) (Animation, bool) {
	var _ok bool
	_ret := bindings.CallElementAnimate(
		this.Ref(), js.Pointer(&_ok),
		keyframes.Ref(),
		options.Ref(),
	)

	return Animation{}.FromRef(_ret), _ok
}

// AnimateFunc returns the method "Element.animate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) AnimateFunc() (fn js.Func[func(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) Animation]) {
	return fn.FromRef(
		bindings.ElementAnimateFunc(
			this.Ref(),
		),
	)
}

// Animate1 calls the method "Element.animate".
//
// The returned bool will be false if there is no such method.
func (this Element) Animate1(keyframes js.Object) (Animation, bool) {
	var _ok bool
	_ret := bindings.CallElementAnimate1(
		this.Ref(), js.Pointer(&_ok),
		keyframes.Ref(),
	)

	return Animation{}.FromRef(_ret), _ok
}

// Animate1Func returns the method "Element.animate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) Animate1Func() (fn js.Func[func(keyframes js.Object) Animation]) {
	return fn.FromRef(
		bindings.ElementAnimate1Func(
			this.Ref(),
		),
	)
}

// GetAnimations calls the method "Element.getAnimations".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAnimations(options GetAnimationsOptions) (js.Array[Animation], bool) {
	var _ok bool
	_ret := bindings.CallElementGetAnimations(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[Animation]{}.FromRef(_ret), _ok
}

// GetAnimationsFunc returns the method "Element.getAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAnimationsFunc() (fn js.Func[func(options GetAnimationsOptions) js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ElementGetAnimationsFunc(
			this.Ref(),
		),
	)
}

// GetAnimations1 calls the method "Element.getAnimations".
//
// The returned bool will be false if there is no such method.
func (this Element) GetAnimations1() (js.Array[Animation], bool) {
	var _ok bool
	_ret := bindings.CallElementGetAnimations1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Animation]{}.FromRef(_ret), _ok
}

// GetAnimations1Func returns the method "Element.getAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetAnimations1Func() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ElementGetAnimations1Func(
			this.Ref(),
		),
	)
}

// GetRegionFlowRanges calls the method "Element.getRegionFlowRanges".
//
// The returned bool will be false if there is no such method.
func (this Element) GetRegionFlowRanges() (js.Array[Range], bool) {
	var _ok bool
	_ret := bindings.CallElementGetRegionFlowRanges(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Range]{}.FromRef(_ret), _ok
}

// GetRegionFlowRangesFunc returns the method "Element.getRegionFlowRanges".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Element) GetRegionFlowRangesFunc() (fn js.Func[func() js.Array[Range]]) {
	return fn.FromRef(
		bindings.ElementGetRegionFlowRangesFunc(
			this.Ref(),
		),
	)
}

type HTMLCollection struct {
	ref js.Ref
}

func (this HTMLCollection) Once() HTMLCollection {
	this.Ref().Once()
	return this
}

func (this HTMLCollection) Ref() js.Ref {
	return this.ref
}

func (this HTMLCollection) FromRef(ref js.Ref) HTMLCollection {
	this.ref = ref
	return this
}

func (this HTMLCollection) Free() {
	this.Ref().Free()
}

// Length returns the value of property "HTMLCollection.length".
//
// The returned bool will be false if there is no such property.
func (this HTMLCollection) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLCollectionLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "HTMLCollection.item".
//
// The returned bool will be false if there is no such method.
func (this HTMLCollection) Item(index uint32) (Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCollectionItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Element{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "HTMLCollection.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCollection) ItemFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLCollectionItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLCollection.namedItem".
//
// The returned bool will be false if there is no such method.
func (this HTMLCollection) NamedItem(name js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCollectionNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "HTMLCollection.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCollection) NamedItemFunc() (fn js.Func[func(name js.String) Element]) {
	return fn.FromRef(
		bindings.HTMLCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

type ElementCreationOptions struct {
	// Is is "ElementCreationOptions.is"
	//
	// Optional
	Is js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ElementCreationOptions with all fields set.
func (p ElementCreationOptions) FromRef(ref js.Ref) ElementCreationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ElementCreationOptions in the application heap.
func (p ElementCreationOptions) New() js.Ref {
	return bindings.ElementCreationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ElementCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.ElementCreationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ElementCreationOptions) Update(ref js.Ref) {
	bindings.ElementCreationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_String_ElementCreationOptions struct {
	ref js.Ref
}

func (x OneOf_String_ElementCreationOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ElementCreationOptions) Free() {
	x.ref.Free()
}

func (x OneOf_String_ElementCreationOptions) FromRef(ref js.Ref) OneOf_String_ElementCreationOptions {
	return OneOf_String_ElementCreationOptions{
		ref: ref,
	}
}

func (x OneOf_String_ElementCreationOptions) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ElementCreationOptions) ElementCreationOptions() ElementCreationOptions {
	var ret ElementCreationOptions
	ret.UpdateFrom(x.ref)
	return ret
}

func NewCDATASection(data js.String) CDATASection {
	return CDATASection{}.FromRef(
		bindings.NewCDATASectionByCDATASection(
			data.Ref()),
	)
}

func NewCDATASectionByCDATASection1() CDATASection {
	return CDATASection{}.FromRef(
		bindings.NewCDATASectionByCDATASection1(),
	)
}

type CDATASection struct {
	Text
}

func (this CDATASection) Once() CDATASection {
	this.Ref().Once()
	return this
}

func (this CDATASection) Ref() js.Ref {
	return this.Text.Ref()
}

func (this CDATASection) FromRef(ref js.Ref) CDATASection {
	this.Text = this.Text.FromRef(ref)
	return this
}

func (this CDATASection) Free() {
	this.Ref().Free()
}

func NewComment(data js.String) Comment {
	return Comment{}.FromRef(
		bindings.NewCommentByComment(
			data.Ref()),
	)
}

func NewCommentByComment1() Comment {
	return Comment{}.FromRef(
		bindings.NewCommentByComment1(),
	)
}

type Comment struct {
	CharacterData
}

func (this Comment) Once() Comment {
	this.Ref().Once()
	return this
}

func (this Comment) Ref() js.Ref {
	return this.CharacterData.Ref()
}

func (this Comment) FromRef(ref js.Ref) Comment {
	this.CharacterData = this.CharacterData.FromRef(ref)
	return this
}

func (this Comment) Free() {
	this.Ref().Free()
}

type ProcessingInstruction struct {
	CharacterData
}

func (this ProcessingInstruction) Once() ProcessingInstruction {
	this.Ref().Once()
	return this
}

func (this ProcessingInstruction) Ref() js.Ref {
	return this.CharacterData.Ref()
}

func (this ProcessingInstruction) FromRef(ref js.Ref) ProcessingInstruction {
	this.CharacterData = this.CharacterData.FromRef(ref)
	return this
}

func (this ProcessingInstruction) Free() {
	this.Ref().Free()
}

// Target returns the value of property "ProcessingInstruction.target".
//
// The returned bool will be false if there is no such property.
func (this ProcessingInstruction) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetProcessingInstructionTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sheet returns the value of property "ProcessingInstruction.sheet".
//
// The returned bool will be false if there is no such property.
func (this ProcessingInstruction) Sheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetProcessingInstructionSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

type NodeFilterFunc func(this js.Ref, node Node) js.Ref

func (fn NodeFilterFunc) Register() js.Func[func(node Node) uint16] {
	return js.RegisterCallback[func(node Node) uint16](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NodeFilterFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Node{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NodeFilter[T any] struct {
	Fn  func(arg T, this js.Ref, node Node) js.Ref
	Arg T
}

func (cb *NodeFilter[T]) Register() js.Func[func(node Node) uint16] {
	return js.RegisterCallback[func(node Node) uint16](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NodeFilter[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		Node{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

const (
	NodeFilter_FILTER_ACCEPT uint16 = 1
	NodeFilter_FILTER_REJECT uint16 = 2
	NodeFilter_FILTER_SKIP   uint16 = 3
)

const (
	NodeFilter_SHOW_ALL                    uint32 = 0xFFFFFFFF
	NodeFilter_SHOW_ELEMENT                uint32 = 0x1
	NodeFilter_SHOW_ATTRIBUTE              uint32 = 0x2
	NodeFilter_SHOW_TEXT                   uint32 = 0x4
	NodeFilter_SHOW_CDATA_SECTION          uint32 = 0x8
	NodeFilter_SHOW_ENTITY_REFERENCE       uint32 = 0x10
	NodeFilter_SHOW_ENTITY                 uint32 = 0x20
	NodeFilter_SHOW_PROCESSING_INSTRUCTION uint32 = 0x40
	NodeFilter_SHOW_COMMENT                uint32 = 0x80
	NodeFilter_SHOW_DOCUMENT               uint32 = 0x100
	NodeFilter_SHOW_DOCUMENT_TYPE          uint32 = 0x200
	NodeFilter_SHOW_DOCUMENT_FRAGMENT      uint32 = 0x400
	NodeFilter_SHOW_NOTATION               uint32 = 0x800
)

type NodeIterator struct {
	ref js.Ref
}

func (this NodeIterator) Once() NodeIterator {
	this.Ref().Once()
	return this
}

func (this NodeIterator) Ref() js.Ref {
	return this.ref
}

func (this NodeIterator) FromRef(ref js.Ref) NodeIterator {
	this.ref = ref
	return this
}

func (this NodeIterator) Free() {
	this.Ref().Free()
}

// Root returns the value of property "NodeIterator.root".
//
// The returned bool will be false if there is no such property.
func (this NodeIterator) Root() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeIteratorRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// ReferenceNode returns the value of property "NodeIterator.referenceNode".
//
// The returned bool will be false if there is no such property.
func (this NodeIterator) ReferenceNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeIteratorReferenceNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// PointerBeforeReferenceNode returns the value of property "NodeIterator.pointerBeforeReferenceNode".
//
// The returned bool will be false if there is no such property.
func (this NodeIterator) PointerBeforeReferenceNode() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNodeIteratorPointerBeforeReferenceNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// WhatToShow returns the value of property "NodeIterator.whatToShow".
//
// The returned bool will be false if there is no such property.
func (this NodeIterator) WhatToShow() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetNodeIteratorWhatToShow(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Filter returns the value of property "NodeIterator.filter".
//
// The returned bool will be false if there is no such property.
func (this NodeIterator) Filter() (js.Func[func(node Node) uint16], bool) {
	var _ok bool
	_ret := bindings.GetNodeIteratorFilter(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Func[func(node Node) uint16]{}.FromRef(_ret), _ok
}

// NextNode calls the method "NodeIterator.nextNode".
//
// The returned bool will be false if there is no such method.
func (this NodeIterator) NextNode() (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeIteratorNextNode(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// NextNodeFunc returns the method "NodeIterator.nextNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NodeIterator) NextNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeIteratorNextNodeFunc(
			this.Ref(),
		),
	)
}

// PreviousNode calls the method "NodeIterator.previousNode".
//
// The returned bool will be false if there is no such method.
func (this NodeIterator) PreviousNode() (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeIteratorPreviousNode(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// PreviousNodeFunc returns the method "NodeIterator.previousNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NodeIterator) PreviousNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeIteratorPreviousNodeFunc(
			this.Ref(),
		),
	)
}

// Detach calls the method "NodeIterator.detach".
//
// The returned bool will be false if there is no such method.
func (this NodeIterator) Detach() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNodeIteratorDetach(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DetachFunc returns the method "NodeIterator.detach".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NodeIterator) DetachFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NodeIteratorDetachFunc(
			this.Ref(),
		),
	)
}

type TreeWalker struct {
	ref js.Ref
}

func (this TreeWalker) Once() TreeWalker {
	this.Ref().Once()
	return this
}

func (this TreeWalker) Ref() js.Ref {
	return this.ref
}

func (this TreeWalker) FromRef(ref js.Ref) TreeWalker {
	this.ref = ref
	return this
}

func (this TreeWalker) Free() {
	this.Ref().Free()
}

// Root returns the value of property "TreeWalker.root".
//
// The returned bool will be false if there is no such property.
func (this TreeWalker) Root() (Node, bool) {
	var _ok bool
	_ret := bindings.GetTreeWalkerRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// WhatToShow returns the value of property "TreeWalker.whatToShow".
//
// The returned bool will be false if there is no such property.
func (this TreeWalker) WhatToShow() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTreeWalkerWhatToShow(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Filter returns the value of property "TreeWalker.filter".
//
// The returned bool will be false if there is no such property.
func (this TreeWalker) Filter() (js.Func[func(node Node) uint16], bool) {
	var _ok bool
	_ret := bindings.GetTreeWalkerFilter(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Func[func(node Node) uint16]{}.FromRef(_ret), _ok
}

// CurrentNode returns the value of property "TreeWalker.currentNode".
//
// The returned bool will be false if there is no such property.
func (this TreeWalker) CurrentNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetTreeWalkerCurrentNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// CurrentNode sets the value of property "TreeWalker.currentNode" to val.
//
// It returns false if the property cannot be set.
func (this TreeWalker) SetCurrentNode(val Node) bool {
	return js.True == bindings.SetTreeWalkerCurrentNode(
		this.Ref(),
		val.Ref(),
	)
}

// ParentNode calls the method "TreeWalker.parentNode".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) ParentNode() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerParentNode(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// ParentNodeFunc returns the method "TreeWalker.parentNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) ParentNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerParentNodeFunc(
			this.Ref(),
		),
	)
}

// FirstChild calls the method "TreeWalker.firstChild".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) FirstChild() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerFirstChild(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// FirstChildFunc returns the method "TreeWalker.firstChild".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) FirstChildFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerFirstChildFunc(
			this.Ref(),
		),
	)
}

// LastChild calls the method "TreeWalker.lastChild".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) LastChild() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerLastChild(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// LastChildFunc returns the method "TreeWalker.lastChild".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) LastChildFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerLastChildFunc(
			this.Ref(),
		),
	)
}

// PreviousSibling calls the method "TreeWalker.previousSibling".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) PreviousSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerPreviousSibling(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// PreviousSiblingFunc returns the method "TreeWalker.previousSibling".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) PreviousSiblingFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerPreviousSiblingFunc(
			this.Ref(),
		),
	)
}

// NextSibling calls the method "TreeWalker.nextSibling".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) NextSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerNextSibling(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// NextSiblingFunc returns the method "TreeWalker.nextSibling".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) NextSiblingFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerNextSiblingFunc(
			this.Ref(),
		),
	)
}

// PreviousNode calls the method "TreeWalker.previousNode".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) PreviousNode() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerPreviousNode(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// PreviousNodeFunc returns the method "TreeWalker.previousNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) PreviousNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerPreviousNodeFunc(
			this.Ref(),
		),
	)
}

// NextNode calls the method "TreeWalker.nextNode".
//
// The returned bool will be false if there is no such method.
func (this TreeWalker) NextNode() (Node, bool) {
	var _ok bool
	_ret := bindings.CallTreeWalkerNextNode(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// NextNodeFunc returns the method "TreeWalker.nextNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TreeWalker) NextNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerNextNodeFunc(
			this.Ref(),
		),
	)
}

type ViewTransition struct {
	ref js.Ref
}

func (this ViewTransition) Once() ViewTransition {
	this.Ref().Once()
	return this
}

func (this ViewTransition) Ref() js.Ref {
	return this.ref
}

func (this ViewTransition) FromRef(ref js.Ref) ViewTransition {
	this.ref = ref
	return this
}

func (this ViewTransition) Free() {
	this.Ref().Free()
}

// UpdateCallbackDone returns the value of property "ViewTransition.updateCallbackDone".
//
// The returned bool will be false if there is no such property.
func (this ViewTransition) UpdateCallbackDone() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetViewTransitionUpdateCallbackDone(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Ready returns the value of property "ViewTransition.ready".
//
// The returned bool will be false if there is no such property.
func (this ViewTransition) Ready() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetViewTransitionReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Finished returns the value of property "ViewTransition.finished".
//
// The returned bool will be false if there is no such property.
func (this ViewTransition) Finished() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetViewTransitionFinished(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SkipTransition calls the method "ViewTransition.skipTransition".
//
// The returned bool will be false if there is no such method.
func (this ViewTransition) SkipTransition() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallViewTransitionSkipTransition(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SkipTransitionFunc returns the method "ViewTransition.skipTransition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ViewTransition) SkipTransitionFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ViewTransitionSkipTransitionFunc(
			this.Ref(),
		),
	)
}

type UpdateCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateCallbackFunc) Register() js.Func[func() js.Promise[js.Any]] {
	return js.RegisterCallback[func() js.Promise[js.Any]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateCallback[T]) Register() js.Func[func() js.Promise[js.Any]] {
	return js.RegisterCallback[func() js.Promise[js.Any]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

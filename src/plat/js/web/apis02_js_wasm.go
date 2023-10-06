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

func NewHTMLSlotElement() (ret HTMLSlotElement) {
	ret.ref = bindings.NewHTMLSlotElementByHTMLSlotElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLSlotElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSlotElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLSlotElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSlotElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLSlotElementName(
		this.Ref(),
		val.Ref(),
	)
}

// HasAssignedNodes returns true if the method "HTMLSlotElement.assignedNodes" exists.
func (this HTMLSlotElement) HasAssignedNodes() bool {
	return js.True == bindings.HasHTMLSlotElementAssignedNodes(
		this.Ref(),
	)
}

// AssignedNodesFunc returns the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodesFunc() (fn js.Func[func(options AssignedNodesOptions) js.Array[Node]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedNodesFunc(
			this.Ref(),
		),
	)
}

// AssignedNodes calls the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodes(options AssignedNodesOptions) (ret js.Array[Node]) {
	bindings.CallHTMLSlotElementAssignedNodes(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAssignedNodes calls the method "HTMLSlotElement.assignedNodes"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSlotElement) TryAssignedNodes(options AssignedNodesOptions) (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedNodes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasAssignedNodes1 returns true if the method "HTMLSlotElement.assignedNodes" exists.
func (this HTMLSlotElement) HasAssignedNodes1() bool {
	return js.True == bindings.HasHTMLSlotElementAssignedNodes1(
		this.Ref(),
	)
}

// AssignedNodes1Func returns the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodes1Func() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedNodes1Func(
			this.Ref(),
		),
	)
}

// AssignedNodes1 calls the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodes1() (ret js.Array[Node]) {
	bindings.CallHTMLSlotElementAssignedNodes1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAssignedNodes1 calls the method "HTMLSlotElement.assignedNodes"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSlotElement) TryAssignedNodes1() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedNodes1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAssignedElements returns true if the method "HTMLSlotElement.assignedElements" exists.
func (this HTMLSlotElement) HasAssignedElements() bool {
	return js.True == bindings.HasHTMLSlotElementAssignedElements(
		this.Ref(),
	)
}

// AssignedElementsFunc returns the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElementsFunc() (fn js.Func[func(options AssignedNodesOptions) js.Array[Element]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedElementsFunc(
			this.Ref(),
		),
	)
}

// AssignedElements calls the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElements(options AssignedNodesOptions) (ret js.Array[Element]) {
	bindings.CallHTMLSlotElementAssignedElements(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAssignedElements calls the method "HTMLSlotElement.assignedElements"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSlotElement) TryAssignedElements(options AssignedNodesOptions) (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedElements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasAssignedElements1 returns true if the method "HTMLSlotElement.assignedElements" exists.
func (this HTMLSlotElement) HasAssignedElements1() bool {
	return js.True == bindings.HasHTMLSlotElementAssignedElements1(
		this.Ref(),
	)
}

// AssignedElements1Func returns the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElements1Func() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignedElements1Func(
			this.Ref(),
		),
	)
}

// AssignedElements1 calls the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElements1() (ret js.Array[Element]) {
	bindings.CallHTMLSlotElementAssignedElements1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAssignedElements1 calls the method "HTMLSlotElement.assignedElements"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSlotElement) TryAssignedElements1() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedElements1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAssign returns true if the method "HTMLSlotElement.assign" exists.
func (this HTMLSlotElement) HasAssign() bool {
	return js.True == bindings.HasHTMLSlotElementAssign(
		this.Ref(),
	)
}

// AssignFunc returns the method "HTMLSlotElement.assign".
func (this HTMLSlotElement) AssignFunc() (fn js.Func[func(nodes ...OneOf_Element_Text)]) {
	return fn.FromRef(
		bindings.HTMLSlotElementAssignFunc(
			this.Ref(),
		),
	)
}

// Assign calls the method "HTMLSlotElement.assign".
func (this HTMLSlotElement) Assign(nodes ...OneOf_Element_Text) (ret js.Void) {
	bindings.CallHTMLSlotElementAssign(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAssign calls the method "HTMLSlotElement.assign"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSlotElement) TryAssign(nodes ...OneOf_Element_Text) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssign(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

func NewText(data js.String) (ret Text) {
	ret.ref = bindings.NewTextByText(
		data.Ref())
	return
}

func NewTextByText1() (ret Text) {
	ret.ref = bindings.NewTextByText1()
	return
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
// It returns ok=false if there is no such property.
func (this Text) WholeText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextWholeText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AssignedSlot returns the value of property "Text.assignedSlot".
//
// It returns ok=false if there is no such property.
func (this Text) AssignedSlot() (ret HTMLSlotElement, ok bool) {
	ok = js.True == bindings.GetTextAssignedSlot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSplitText returns true if the method "Text.splitText" exists.
func (this Text) HasSplitText() bool {
	return js.True == bindings.HasTextSplitText(
		this.Ref(),
	)
}

// SplitTextFunc returns the method "Text.splitText".
func (this Text) SplitTextFunc() (fn js.Func[func(offset uint32) Text]) {
	return fn.FromRef(
		bindings.TextSplitTextFunc(
			this.Ref(),
		),
	)
}

// SplitText calls the method "Text.splitText".
func (this Text) SplitText(offset uint32) (ret Text) {
	bindings.CallTextSplitText(
		this.Ref(), js.Pointer(&ret),
		uint32(offset),
	)

	return
}

// TrySplitText calls the method "Text.splitText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TrySplitText(offset uint32) (ret Text, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextSplitText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
	)

	return
}

// HasGetBoxQuads returns true if the method "Text.getBoxQuads" exists.
func (this Text) HasGetBoxQuads() bool {
	return js.True == bindings.HasTextGetBoxQuads(
		this.Ref(),
	)
}

// GetBoxQuadsFunc returns the method "Text.getBoxQuads".
func (this Text) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.TextGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Text.getBoxQuads".
func (this Text) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallTextGetBoxQuads(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Text.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextGetBoxQuads(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetBoxQuads1 returns true if the method "Text.getBoxQuads" exists.
func (this Text) HasGetBoxQuads1() bool {
	return js.True == bindings.HasTextGetBoxQuads1(
		this.Ref(),
	)
}

// GetBoxQuads1Func returns the method "Text.getBoxQuads".
func (this Text) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.TextGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Text.getBoxQuads".
func (this Text) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallTextGetBoxQuads1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Text.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextGetBoxQuads1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvertQuadFromNode returns true if the method "Text.convertQuadFromNode" exists.
func (this Text) HasConvertQuadFromNode() bool {
	return js.True == bindings.HasTextConvertQuadFromNode(
		this.Ref(),
	)
}

// ConvertQuadFromNodeFunc returns the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallTextConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "Text.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertQuadFromNode1 returns true if the method "Text.convertQuadFromNode" exists.
func (this Text) HasConvertQuadFromNode1() bool {
	return js.True == bindings.HasTextConvertQuadFromNode1(
		this.Ref(),
	)
}

// ConvertQuadFromNode1Func returns the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallTextConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "Text.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasConvertRectFromNode returns true if the method "Text.convertRectFromNode" exists.
func (this Text) HasConvertRectFromNode() bool {
	return js.True == bindings.HasTextConvertRectFromNode(
		this.Ref(),
	)
}

// ConvertRectFromNodeFunc returns the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallTextConvertRectFromNode(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "Text.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertRectFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertRectFromNode1 returns true if the method "Text.convertRectFromNode" exists.
func (this Text) HasConvertRectFromNode1() bool {
	return js.True == bindings.HasTextConvertRectFromNode1(
		this.Ref(),
	)
}

// ConvertRectFromNode1Func returns the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.TextConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallTextConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "Text.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasConvertPointFromNode returns true if the method "Text.convertPointFromNode" exists.
func (this Text) HasConvertPointFromNode() bool {
	return js.True == bindings.HasTextConvertPointFromNode(
		this.Ref(),
	)
}

// ConvertPointFromNodeFunc returns the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.TextConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallTextConvertPointFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "Text.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertPointFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertPointFromNode1 returns true if the method "Text.convertPointFromNode" exists.
func (this Text) HasConvertPointFromNode1() bool {
	return js.True == bindings.HasTextConvertPointFromNode1(
		this.Ref(),
	)
}

// ConvertPointFromNode1Func returns the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.TextConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallTextConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "Text.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Text) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Element returns the value of property "CSSPseudoElement.element".
//
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Element() (ret Element, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Parent returns the value of property "CSSPseudoElement.parent".
//
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Parent() (ret OneOf_Element_CSSPseudoElement, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementParent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPseudo returns true if the method "CSSPseudoElement.pseudo" exists.
func (this CSSPseudoElement) HasPseudo() bool {
	return js.True == bindings.HasCSSPseudoElementPseudo(
		this.Ref(),
	)
}

// PseudoFunc returns the method "CSSPseudoElement.pseudo".
func (this CSSPseudoElement) PseudoFunc() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	return fn.FromRef(
		bindings.CSSPseudoElementPseudoFunc(
			this.Ref(),
		),
	)
}

// Pseudo calls the method "CSSPseudoElement.pseudo".
func (this CSSPseudoElement) Pseudo(typ js.String) (ret CSSPseudoElement) {
	bindings.CallCSSPseudoElementPseudo(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryPseudo calls the method "CSSPseudoElement.pseudo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryPseudo(typ js.String) (ret CSSPseudoElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementPseudo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasGetBoxQuads returns true if the method "CSSPseudoElement.getBoxQuads" exists.
func (this CSSPseudoElement) HasGetBoxQuads() bool {
	return js.True == bindings.HasCSSPseudoElementGetBoxQuads(
		this.Ref(),
	)
}

// GetBoxQuadsFunc returns the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.CSSPseudoElementGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallCSSPseudoElementGetBoxQuads(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "CSSPseudoElement.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementGetBoxQuads(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetBoxQuads1 returns true if the method "CSSPseudoElement.getBoxQuads" exists.
func (this CSSPseudoElement) HasGetBoxQuads1() bool {
	return js.True == bindings.HasCSSPseudoElementGetBoxQuads1(
		this.Ref(),
	)
}

// GetBoxQuads1Func returns the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.CSSPseudoElementGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallCSSPseudoElementGetBoxQuads1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "CSSPseudoElement.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementGetBoxQuads1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvertQuadFromNode returns true if the method "CSSPseudoElement.convertQuadFromNode" exists.
func (this CSSPseudoElement) HasConvertQuadFromNode() bool {
	return js.True == bindings.HasCSSPseudoElementConvertQuadFromNode(
		this.Ref(),
	)
}

// ConvertQuadFromNodeFunc returns the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "CSSPseudoElement.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertQuadFromNode1 returns true if the method "CSSPseudoElement.convertQuadFromNode" exists.
func (this CSSPseudoElement) HasConvertQuadFromNode1() bool {
	return js.True == bindings.HasCSSPseudoElementConvertQuadFromNode1(
		this.Ref(),
	)
}

// ConvertQuadFromNode1Func returns the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "CSSPseudoElement.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasConvertRectFromNode returns true if the method "CSSPseudoElement.convertRectFromNode" exists.
func (this CSSPseudoElement) HasConvertRectFromNode() bool {
	return js.True == bindings.HasCSSPseudoElementConvertRectFromNode(
		this.Ref(),
	)
}

// ConvertRectFromNodeFunc returns the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertRectFromNode(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "CSSPseudoElement.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertRectFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertRectFromNode1 returns true if the method "CSSPseudoElement.convertRectFromNode" exists.
func (this CSSPseudoElement) HasConvertRectFromNode1() bool {
	return js.True == bindings.HasCSSPseudoElementConvertRectFromNode1(
		this.Ref(),
	)
}

// ConvertRectFromNode1Func returns the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "CSSPseudoElement.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasConvertPointFromNode returns true if the method "CSSPseudoElement.convertPointFromNode" exists.
func (this CSSPseudoElement) HasConvertPointFromNode() bool {
	return js.True == bindings.HasCSSPseudoElementConvertPointFromNode(
		this.Ref(),
	)
}

// ConvertPointFromNodeFunc returns the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallCSSPseudoElementConvertPointFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "CSSPseudoElement.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertPointFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertPointFromNode1 returns true if the method "CSSPseudoElement.convertPointFromNode" exists.
func (this CSSPseudoElement) HasConvertPointFromNode1() bool {
	return js.True == bindings.HasCSSPseudoElementConvertPointFromNode1(
		this.Ref(),
	)
}

// ConvertPointFromNode1Func returns the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.CSSPseudoElementConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallCSSPseudoElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "CSSPseudoElement.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSPseudoElement) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this NodeList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNodeListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "NodeList.item" exists.
func (this NodeList) HasItem() bool {
	return js.True == bindings.HasNodeListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "NodeList.item".
func (this NodeList) ItemFunc() (fn js.Func[func(index uint32) Node]) {
	return fn.FromRef(
		bindings.NodeListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "NodeList.item".
func (this NodeList) Item(index uint32) (ret Node) {
	bindings.CallNodeListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "NodeList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NodeList) TryItem(index uint32) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DocumentFragment) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentChildren(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "DocumentFragment.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentFirstElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "DocumentFragment.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentLastElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "DocumentFragment.childElementCount".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentChildElementCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetElementById returns true if the method "DocumentFragment.getElementById" exists.
func (this DocumentFragment) HasGetElementById() bool {
	return js.True == bindings.HasDocumentFragmentGetElementById(
		this.Ref(),
	)
}

// GetElementByIdFunc returns the method "DocumentFragment.getElementById".
func (this DocumentFragment) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentFragmentGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// GetElementById calls the method "DocumentFragment.getElementById".
func (this DocumentFragment) GetElementById(elementId js.String) (ret Element) {
	bindings.CallDocumentFragmentGetElementById(
		this.Ref(), js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "DocumentFragment.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentGetElementById(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasPrepend returns true if the method "DocumentFragment.prepend" exists.
func (this DocumentFragment) HasPrepend() bool {
	return js.True == bindings.HasDocumentFragmentPrepend(
		this.Ref(),
	)
}

// PrependFunc returns the method "DocumentFragment.prepend".
func (this DocumentFragment) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentPrependFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "DocumentFragment.prepend".
func (this DocumentFragment) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentPrepend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryPrepend calls the method "DocumentFragment.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryPrepend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentPrepend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAppend returns true if the method "DocumentFragment.append" exists.
func (this DocumentFragment) HasAppend() bool {
	return js.True == bindings.HasDocumentFragmentAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "DocumentFragment.append".
func (this DocumentFragment) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "DocumentFragment.append".
func (this DocumentFragment) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentAppend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAppend calls the method "DocumentFragment.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryAppend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceChildren returns true if the method "DocumentFragment.replaceChildren" exists.
func (this DocumentFragment) HasReplaceChildren() bool {
	return js.True == bindings.HasDocumentFragmentReplaceChildren(
		this.Ref(),
	)
}

// ReplaceChildrenFunc returns the method "DocumentFragment.replaceChildren".
func (this DocumentFragment) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentFragmentReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "DocumentFragment.replaceChildren".
func (this DocumentFragment) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentReplaceChildren(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceChildren calls the method "DocumentFragment.replaceChildren"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentReplaceChildren(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasQuerySelector returns true if the method "DocumentFragment.querySelector" exists.
func (this DocumentFragment) HasQuerySelector() bool {
	return js.True == bindings.HasDocumentFragmentQuerySelector(
		this.Ref(),
	)
}

// QuerySelectorFunc returns the method "DocumentFragment.querySelector".
func (this DocumentFragment) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentFragmentQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "DocumentFragment.querySelector".
func (this DocumentFragment) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallDocumentFragmentQuerySelector(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "DocumentFragment.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentQuerySelector(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasQuerySelectorAll returns true if the method "DocumentFragment.querySelectorAll" exists.
func (this DocumentFragment) HasQuerySelectorAll() bool {
	return js.True == bindings.HasDocumentFragmentQuerySelectorAll(
		this.Ref(),
	)
}

// QuerySelectorAllFunc returns the method "DocumentFragment.querySelectorAll".
func (this DocumentFragment) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentFragmentQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "DocumentFragment.querySelectorAll".
func (this DocumentFragment) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallDocumentFragmentQuerySelectorAll(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "DocumentFragment.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentFragment) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentQuerySelectorAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
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

func NewSanitizer(config SanitizerConfig) (ret Sanitizer) {
	ret.ref = bindings.NewSanitizerBySanitizer(
		js.Pointer(&config))
	return
}

func NewSanitizerBySanitizer1() (ret Sanitizer) {
	ret.ref = bindings.NewSanitizerBySanitizer1()
	return
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

// HasSanitize returns true if the method "Sanitizer.sanitize" exists.
func (this Sanitizer) HasSanitize() bool {
	return js.True == bindings.HasSanitizerSanitize(
		this.Ref(),
	)
}

// SanitizeFunc returns the method "Sanitizer.sanitize".
func (this Sanitizer) SanitizeFunc() (fn js.Func[func(input OneOf_Document_DocumentFragment) DocumentFragment]) {
	return fn.FromRef(
		bindings.SanitizerSanitizeFunc(
			this.Ref(),
		),
	)
}

// Sanitize calls the method "Sanitizer.sanitize".
func (this Sanitizer) Sanitize(input OneOf_Document_DocumentFragment) (ret DocumentFragment) {
	bindings.CallSanitizerSanitize(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySanitize calls the method "Sanitizer.sanitize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Sanitizer) TrySanitize(input OneOf_Document_DocumentFragment) (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerSanitize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSanitizeFor returns true if the method "Sanitizer.sanitizeFor" exists.
func (this Sanitizer) HasSanitizeFor() bool {
	return js.True == bindings.HasSanitizerSanitizeFor(
		this.Ref(),
	)
}

// SanitizeForFunc returns the method "Sanitizer.sanitizeFor".
func (this Sanitizer) SanitizeForFunc() (fn js.Func[func(element js.String, input js.String) Element]) {
	return fn.FromRef(
		bindings.SanitizerSanitizeForFunc(
			this.Ref(),
		),
	)
}

// SanitizeFor calls the method "Sanitizer.sanitizeFor".
func (this Sanitizer) SanitizeFor(element js.String, input js.String) (ret Element) {
	bindings.CallSanitizerSanitizeFor(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
		input.Ref(),
	)

	return
}

// TrySanitizeFor calls the method "Sanitizer.sanitizeFor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Sanitizer) TrySanitizeFor(element js.String, input js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerSanitizeFor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		input.Ref(),
	)

	return
}

// HasGetConfiguration returns true if the method "Sanitizer.getConfiguration" exists.
func (this Sanitizer) HasGetConfiguration() bool {
	return js.True == bindings.HasSanitizerGetConfiguration(
		this.Ref(),
	)
}

// GetConfigurationFunc returns the method "Sanitizer.getConfiguration".
func (this Sanitizer) GetConfigurationFunc() (fn js.Func[func() SanitizerConfig]) {
	return fn.FromRef(
		bindings.SanitizerGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// GetConfiguration calls the method "Sanitizer.getConfiguration".
func (this Sanitizer) GetConfiguration() (ret SanitizerConfig) {
	bindings.CallSanitizerGetConfiguration(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetConfiguration calls the method "Sanitizer.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Sanitizer) TryGetConfiguration() (ret SanitizerConfig, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerGetConfiguration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDefaultConfiguration returns true if the staticmethod "Sanitizer.getDefaultConfiguration" exists.
func (this Sanitizer) HasGetDefaultConfiguration() bool {
	return js.True == bindings.HasSanitizerGetDefaultConfiguration(
		this.Ref(),
	)
}

// GetDefaultConfigurationFunc returns the staticmethod "Sanitizer.getDefaultConfiguration".
func (this Sanitizer) GetDefaultConfigurationFunc() (fn js.Func[func() SanitizerConfig]) {
	return fn.FromRef(
		bindings.SanitizerGetDefaultConfigurationFunc(
			this.Ref(),
		),
	)
}

// GetDefaultConfiguration calls the staticmethod "Sanitizer.getDefaultConfiguration".
func (this Sanitizer) GetDefaultConfiguration() (ret SanitizerConfig) {
	bindings.CallSanitizerGetDefaultConfiguration(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDefaultConfiguration calls the staticmethod "Sanitizer.getDefaultConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Sanitizer) TryGetDefaultConfiguration() (ret SanitizerConfig, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerGetDefaultConfiguration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DOMRectList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMRectListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "DOMRectList.item" exists.
func (this DOMRectList) HasItem() bool {
	return js.True == bindings.HasDOMRectListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "DOMRectList.item".
func (this DOMRectList) ItemFunc() (fn js.Func[func(index uint32) DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "DOMRectList.item".
func (this DOMRectList) Item(index uint32) (ret DOMRect) {
	bindings.CallDOMRectListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMRectList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRectList) TryItem(index uint32) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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

func NewCSSKeywordValue(value js.String) (ret CSSKeywordValue) {
	ret.ref = bindings.NewCSSKeywordValueByCSSKeywordValue(
		value.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSKeywordValue) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeywordValueValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "CSSKeywordValue.value" to val.
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
// It returns ok=false if there is no such property.
func (this Range) CommonAncestorContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetRangeCommonAncestorContainer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetStart returns true if the method "Range.setStart" exists.
func (this Range) HasSetStart() bool {
	return js.True == bindings.HasRangeSetStart(
		this.Ref(),
	)
}

// SetStartFunc returns the method "Range.setStart".
func (this Range) SetStartFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.RangeSetStartFunc(
			this.Ref(),
		),
	)
}

// SetStart calls the method "Range.setStart".
func (this Range) SetStart(node Node, offset uint32) (ret js.Void) {
	bindings.CallRangeSetStart(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetStart calls the method "Range.setStart"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetStart(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasSetEnd returns true if the method "Range.setEnd" exists.
func (this Range) HasSetEnd() bool {
	return js.True == bindings.HasRangeSetEnd(
		this.Ref(),
	)
}

// SetEndFunc returns the method "Range.setEnd".
func (this Range) SetEndFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.RangeSetEndFunc(
			this.Ref(),
		),
	)
}

// SetEnd calls the method "Range.setEnd".
func (this Range) SetEnd(node Node, offset uint32) (ret js.Void) {
	bindings.CallRangeSetEnd(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetEnd calls the method "Range.setEnd"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetEnd(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEnd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasSetStartBefore returns true if the method "Range.setStartBefore" exists.
func (this Range) HasSetStartBefore() bool {
	return js.True == bindings.HasRangeSetStartBefore(
		this.Ref(),
	)
}

// SetStartBeforeFunc returns the method "Range.setStartBefore".
func (this Range) SetStartBeforeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetStartBeforeFunc(
			this.Ref(),
		),
	)
}

// SetStartBefore calls the method "Range.setStartBefore".
func (this Range) SetStartBefore(node Node) (ret js.Void) {
	bindings.CallRangeSetStartBefore(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetStartBefore calls the method "Range.setStartBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetStartBefore(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStartBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSetStartAfter returns true if the method "Range.setStartAfter" exists.
func (this Range) HasSetStartAfter() bool {
	return js.True == bindings.HasRangeSetStartAfter(
		this.Ref(),
	)
}

// SetStartAfterFunc returns the method "Range.setStartAfter".
func (this Range) SetStartAfterFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetStartAfterFunc(
			this.Ref(),
		),
	)
}

// SetStartAfter calls the method "Range.setStartAfter".
func (this Range) SetStartAfter(node Node) (ret js.Void) {
	bindings.CallRangeSetStartAfter(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetStartAfter calls the method "Range.setStartAfter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetStartAfter(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStartAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSetEndBefore returns true if the method "Range.setEndBefore" exists.
func (this Range) HasSetEndBefore() bool {
	return js.True == bindings.HasRangeSetEndBefore(
		this.Ref(),
	)
}

// SetEndBeforeFunc returns the method "Range.setEndBefore".
func (this Range) SetEndBeforeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetEndBeforeFunc(
			this.Ref(),
		),
	)
}

// SetEndBefore calls the method "Range.setEndBefore".
func (this Range) SetEndBefore(node Node) (ret js.Void) {
	bindings.CallRangeSetEndBefore(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetEndBefore calls the method "Range.setEndBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetEndBefore(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEndBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSetEndAfter returns true if the method "Range.setEndAfter" exists.
func (this Range) HasSetEndAfter() bool {
	return js.True == bindings.HasRangeSetEndAfter(
		this.Ref(),
	)
}

// SetEndAfterFunc returns the method "Range.setEndAfter".
func (this Range) SetEndAfterFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSetEndAfterFunc(
			this.Ref(),
		),
	)
}

// SetEndAfter calls the method "Range.setEndAfter".
func (this Range) SetEndAfter(node Node) (ret js.Void) {
	bindings.CallRangeSetEndAfter(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetEndAfter calls the method "Range.setEndAfter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySetEndAfter(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEndAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasCollapse returns true if the method "Range.collapse" exists.
func (this Range) HasCollapse() bool {
	return js.True == bindings.HasRangeCollapse(
		this.Ref(),
	)
}

// CollapseFunc returns the method "Range.collapse".
func (this Range) CollapseFunc() (fn js.Func[func(toStart bool)]) {
	return fn.FromRef(
		bindings.RangeCollapseFunc(
			this.Ref(),
		),
	)
}

// Collapse calls the method "Range.collapse".
func (this Range) Collapse(toStart bool) (ret js.Void) {
	bindings.CallRangeCollapse(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(toStart)),
	)

	return
}

// TryCollapse calls the method "Range.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCollapse(toStart bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCollapse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(toStart)),
	)

	return
}

// HasCollapse1 returns true if the method "Range.collapse" exists.
func (this Range) HasCollapse1() bool {
	return js.True == bindings.HasRangeCollapse1(
		this.Ref(),
	)
}

// Collapse1Func returns the method "Range.collapse".
func (this Range) Collapse1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeCollapse1Func(
			this.Ref(),
		),
	)
}

// Collapse1 calls the method "Range.collapse".
func (this Range) Collapse1() (ret js.Void) {
	bindings.CallRangeCollapse1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCollapse1 calls the method "Range.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCollapse1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCollapse1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSelectNode returns true if the method "Range.selectNode" exists.
func (this Range) HasSelectNode() bool {
	return js.True == bindings.HasRangeSelectNode(
		this.Ref(),
	)
}

// SelectNodeFunc returns the method "Range.selectNode".
func (this Range) SelectNodeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSelectNodeFunc(
			this.Ref(),
		),
	)
}

// SelectNode calls the method "Range.selectNode".
func (this Range) SelectNode(node Node) (ret js.Void) {
	bindings.CallRangeSelectNode(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectNode calls the method "Range.selectNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySelectNode(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSelectNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSelectNodeContents returns true if the method "Range.selectNodeContents" exists.
func (this Range) HasSelectNodeContents() bool {
	return js.True == bindings.HasRangeSelectNodeContents(
		this.Ref(),
	)
}

// SelectNodeContentsFunc returns the method "Range.selectNodeContents".
func (this Range) SelectNodeContentsFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeSelectNodeContentsFunc(
			this.Ref(),
		),
	)
}

// SelectNodeContents calls the method "Range.selectNodeContents".
func (this Range) SelectNodeContents(node Node) (ret js.Void) {
	bindings.CallRangeSelectNodeContents(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectNodeContents calls the method "Range.selectNodeContents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySelectNodeContents(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSelectNodeContents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasCompareBoundaryPoints returns true if the method "Range.compareBoundaryPoints" exists.
func (this Range) HasCompareBoundaryPoints() bool {
	return js.True == bindings.HasRangeCompareBoundaryPoints(
		this.Ref(),
	)
}

// CompareBoundaryPointsFunc returns the method "Range.compareBoundaryPoints".
func (this Range) CompareBoundaryPointsFunc() (fn js.Func[func(how uint16, sourceRange Range) int16]) {
	return fn.FromRef(
		bindings.RangeCompareBoundaryPointsFunc(
			this.Ref(),
		),
	)
}

// CompareBoundaryPoints calls the method "Range.compareBoundaryPoints".
func (this Range) CompareBoundaryPoints(how uint16, sourceRange Range) (ret int16) {
	bindings.CallRangeCompareBoundaryPoints(
		this.Ref(), js.Pointer(&ret),
		uint32(how),
		sourceRange.Ref(),
	)

	return
}

// TryCompareBoundaryPoints calls the method "Range.compareBoundaryPoints"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCompareBoundaryPoints(how uint16, sourceRange Range) (ret int16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCompareBoundaryPoints(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(how),
		sourceRange.Ref(),
	)

	return
}

// HasDeleteContents returns true if the method "Range.deleteContents" exists.
func (this Range) HasDeleteContents() bool {
	return js.True == bindings.HasRangeDeleteContents(
		this.Ref(),
	)
}

// DeleteContentsFunc returns the method "Range.deleteContents".
func (this Range) DeleteContentsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeDeleteContentsFunc(
			this.Ref(),
		),
	)
}

// DeleteContents calls the method "Range.deleteContents".
func (this Range) DeleteContents() (ret js.Void) {
	bindings.CallRangeDeleteContents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeleteContents calls the method "Range.deleteContents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryDeleteContents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeDeleteContents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExtractContents returns true if the method "Range.extractContents" exists.
func (this Range) HasExtractContents() bool {
	return js.True == bindings.HasRangeExtractContents(
		this.Ref(),
	)
}

// ExtractContentsFunc returns the method "Range.extractContents".
func (this Range) ExtractContentsFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeExtractContentsFunc(
			this.Ref(),
		),
	)
}

// ExtractContents calls the method "Range.extractContents".
func (this Range) ExtractContents() (ret DocumentFragment) {
	bindings.CallRangeExtractContents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryExtractContents calls the method "Range.extractContents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryExtractContents() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeExtractContents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCloneContents returns true if the method "Range.cloneContents" exists.
func (this Range) HasCloneContents() bool {
	return js.True == bindings.HasRangeCloneContents(
		this.Ref(),
	)
}

// CloneContentsFunc returns the method "Range.cloneContents".
func (this Range) CloneContentsFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeCloneContentsFunc(
			this.Ref(),
		),
	)
}

// CloneContents calls the method "Range.cloneContents".
func (this Range) CloneContents() (ret DocumentFragment) {
	bindings.CallRangeCloneContents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCloneContents calls the method "Range.cloneContents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCloneContents() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCloneContents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertNode returns true if the method "Range.insertNode" exists.
func (this Range) HasInsertNode() bool {
	return js.True == bindings.HasRangeInsertNode(
		this.Ref(),
	)
}

// InsertNodeFunc returns the method "Range.insertNode".
func (this Range) InsertNodeFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.RangeInsertNodeFunc(
			this.Ref(),
		),
	)
}

// InsertNode calls the method "Range.insertNode".
func (this Range) InsertNode(node Node) (ret js.Void) {
	bindings.CallRangeInsertNode(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryInsertNode calls the method "Range.insertNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryInsertNode(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeInsertNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSurroundContents returns true if the method "Range.surroundContents" exists.
func (this Range) HasSurroundContents() bool {
	return js.True == bindings.HasRangeSurroundContents(
		this.Ref(),
	)
}

// SurroundContentsFunc returns the method "Range.surroundContents".
func (this Range) SurroundContentsFunc() (fn js.Func[func(newParent Node)]) {
	return fn.FromRef(
		bindings.RangeSurroundContentsFunc(
			this.Ref(),
		),
	)
}

// SurroundContents calls the method "Range.surroundContents".
func (this Range) SurroundContents(newParent Node) (ret js.Void) {
	bindings.CallRangeSurroundContents(
		this.Ref(), js.Pointer(&ret),
		newParent.Ref(),
	)

	return
}

// TrySurroundContents calls the method "Range.surroundContents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TrySurroundContents(newParent Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSurroundContents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newParent.Ref(),
	)

	return
}

// HasCloneRange returns true if the method "Range.cloneRange" exists.
func (this Range) HasCloneRange() bool {
	return js.True == bindings.HasRangeCloneRange(
		this.Ref(),
	)
}

// CloneRangeFunc returns the method "Range.cloneRange".
func (this Range) CloneRangeFunc() (fn js.Func[func() Range]) {
	return fn.FromRef(
		bindings.RangeCloneRangeFunc(
			this.Ref(),
		),
	)
}

// CloneRange calls the method "Range.cloneRange".
func (this Range) CloneRange() (ret Range) {
	bindings.CallRangeCloneRange(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCloneRange calls the method "Range.cloneRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCloneRange() (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCloneRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDetach returns true if the method "Range.detach" exists.
func (this Range) HasDetach() bool {
	return js.True == bindings.HasRangeDetach(
		this.Ref(),
	)
}

// DetachFunc returns the method "Range.detach".
func (this Range) DetachFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.RangeDetachFunc(
			this.Ref(),
		),
	)
}

// Detach calls the method "Range.detach".
func (this Range) Detach() (ret js.Void) {
	bindings.CallRangeDetach(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDetach calls the method "Range.detach"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryDetach() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeDetach(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsPointInRange returns true if the method "Range.isPointInRange" exists.
func (this Range) HasIsPointInRange() bool {
	return js.True == bindings.HasRangeIsPointInRange(
		this.Ref(),
	)
}

// IsPointInRangeFunc returns the method "Range.isPointInRange".
func (this Range) IsPointInRangeFunc() (fn js.Func[func(node Node, offset uint32) bool]) {
	return fn.FromRef(
		bindings.RangeIsPointInRangeFunc(
			this.Ref(),
		),
	)
}

// IsPointInRange calls the method "Range.isPointInRange".
func (this Range) IsPointInRange(node Node, offset uint32) (ret bool) {
	bindings.CallRangeIsPointInRange(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryIsPointInRange calls the method "Range.isPointInRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryIsPointInRange(node Node, offset uint32) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeIsPointInRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasComparePoint returns true if the method "Range.comparePoint" exists.
func (this Range) HasComparePoint() bool {
	return js.True == bindings.HasRangeComparePoint(
		this.Ref(),
	)
}

// ComparePointFunc returns the method "Range.comparePoint".
func (this Range) ComparePointFunc() (fn js.Func[func(node Node, offset uint32) int16]) {
	return fn.FromRef(
		bindings.RangeComparePointFunc(
			this.Ref(),
		),
	)
}

// ComparePoint calls the method "Range.comparePoint".
func (this Range) ComparePoint(node Node, offset uint32) (ret int16) {
	bindings.CallRangeComparePoint(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryComparePoint calls the method "Range.comparePoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryComparePoint(node Node, offset uint32) (ret int16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeComparePoint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasIntersectsNode returns true if the method "Range.intersectsNode" exists.
func (this Range) HasIntersectsNode() bool {
	return js.True == bindings.HasRangeIntersectsNode(
		this.Ref(),
	)
}

// IntersectsNodeFunc returns the method "Range.intersectsNode".
func (this Range) IntersectsNodeFunc() (fn js.Func[func(node Node) bool]) {
	return fn.FromRef(
		bindings.RangeIntersectsNodeFunc(
			this.Ref(),
		),
	)
}

// IntersectsNode calls the method "Range.intersectsNode".
func (this Range) IntersectsNode(node Node) (ret bool) {
	bindings.CallRangeIntersectsNode(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryIntersectsNode calls the method "Range.intersectsNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryIntersectsNode(node Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeIntersectsNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasToString returns true if the method "Range.toString" exists.
func (this Range) HasToString() bool {
	return js.True == bindings.HasRangeToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "Range.toString".
func (this Range) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.RangeToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "Range.toString".
func (this Range) ToString() (ret js.String) {
	bindings.CallRangeToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "Range.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetClientRects returns true if the method "Range.getClientRects" exists.
func (this Range) HasGetClientRects() bool {
	return js.True == bindings.HasRangeGetClientRects(
		this.Ref(),
	)
}

// GetClientRectsFunc returns the method "Range.getClientRects".
func (this Range) GetClientRectsFunc() (fn js.Func[func() DOMRectList]) {
	return fn.FromRef(
		bindings.RangeGetClientRectsFunc(
			this.Ref(),
		),
	)
}

// GetClientRects calls the method "Range.getClientRects".
func (this Range) GetClientRects() (ret DOMRectList) {
	bindings.CallRangeGetClientRects(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetClientRects calls the method "Range.getClientRects"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryGetClientRects() (ret DOMRectList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeGetClientRects(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetBoundingClientRect returns true if the method "Range.getBoundingClientRect" exists.
func (this Range) HasGetBoundingClientRect() bool {
	return js.True == bindings.HasRangeGetBoundingClientRect(
		this.Ref(),
	)
}

// GetBoundingClientRectFunc returns the method "Range.getBoundingClientRect".
func (this Range) GetBoundingClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.RangeGetBoundingClientRectFunc(
			this.Ref(),
		),
	)
}

// GetBoundingClientRect calls the method "Range.getBoundingClientRect".
func (this Range) GetBoundingClientRect() (ret DOMRect) {
	bindings.CallRangeGetBoundingClientRect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoundingClientRect calls the method "Range.getBoundingClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryGetBoundingClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeGetBoundingClientRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateContextualFragment returns true if the method "Range.createContextualFragment" exists.
func (this Range) HasCreateContextualFragment() bool {
	return js.True == bindings.HasRangeCreateContextualFragment(
		this.Ref(),
	)
}

// CreateContextualFragmentFunc returns the method "Range.createContextualFragment".
func (this Range) CreateContextualFragmentFunc() (fn js.Func[func(fragment js.String) DocumentFragment]) {
	return fn.FromRef(
		bindings.RangeCreateContextualFragmentFunc(
			this.Ref(),
		),
	)
}

// CreateContextualFragment calls the method "Range.createContextualFragment".
func (this Range) CreateContextualFragment(fragment js.String) (ret DocumentFragment) {
	bindings.CallRangeCreateContextualFragment(
		this.Ref(), js.Pointer(&ret),
		fragment.Ref(),
	)

	return
}

// TryCreateContextualFragment calls the method "Range.createContextualFragment"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Range) TryCreateContextualFragment(fragment js.String) (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCreateContextualFragment(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		fragment.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DOMTokenList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMTokenListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "DOMTokenList.value".
//
// It returns ok=false if there is no such property.
func (this DOMTokenList) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMTokenListValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "DOMTokenList.value" to val.
//
// It returns false if the property cannot be set.
func (this DOMTokenList) SetValue(val js.String) bool {
	return js.True == bindings.SetDOMTokenListValue(
		this.Ref(),
		val.Ref(),
	)
}

// HasItem returns true if the method "DOMTokenList.item" exists.
func (this DOMTokenList) HasItem() bool {
	return js.True == bindings.HasDOMTokenListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "DOMTokenList.item".
func (this DOMTokenList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.DOMTokenListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "DOMTokenList.item".
func (this DOMTokenList) Item(index uint32) (ret js.String) {
	bindings.CallDOMTokenListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMTokenList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasContains returns true if the method "DOMTokenList.contains" exists.
func (this DOMTokenList) HasContains() bool {
	return js.True == bindings.HasDOMTokenListContains(
		this.Ref(),
	)
}

// ContainsFunc returns the method "DOMTokenList.contains".
func (this DOMTokenList) ContainsFunc() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListContainsFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "DOMTokenList.contains".
func (this DOMTokenList) Contains(token js.String) (ret bool) {
	bindings.CallDOMTokenListContains(
		this.Ref(), js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TryContains calls the method "DOMTokenList.contains"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryContains(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListContains(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
}

// HasAdd returns true if the method "DOMTokenList.add" exists.
func (this DOMTokenList) HasAdd() bool {
	return js.True == bindings.HasDOMTokenListAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "DOMTokenList.add".
func (this DOMTokenList) AddFunc() (fn js.Func[func(tokens ...js.String)]) {
	return fn.FromRef(
		bindings.DOMTokenListAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "DOMTokenList.add".
func (this DOMTokenList) Add(tokens ...js.String) (ret js.Void) {
	bindings.CallDOMTokenListAdd(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// TryAdd calls the method "DOMTokenList.add"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryAdd(tokens ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// HasRemove returns true if the method "DOMTokenList.remove" exists.
func (this DOMTokenList) HasRemove() bool {
	return js.True == bindings.HasDOMTokenListRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "DOMTokenList.remove".
func (this DOMTokenList) RemoveFunc() (fn js.Func[func(tokens ...js.String)]) {
	return fn.FromRef(
		bindings.DOMTokenListRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "DOMTokenList.remove".
func (this DOMTokenList) Remove(tokens ...js.String) (ret js.Void) {
	bindings.CallDOMTokenListRemove(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// TryRemove calls the method "DOMTokenList.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryRemove(tokens ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// HasToggle returns true if the method "DOMTokenList.toggle" exists.
func (this DOMTokenList) HasToggle() bool {
	return js.True == bindings.HasDOMTokenListToggle(
		this.Ref(),
	)
}

// ToggleFunc returns the method "DOMTokenList.toggle".
func (this DOMTokenList) ToggleFunc() (fn js.Func[func(token js.String, force bool) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListToggleFunc(
			this.Ref(),
		),
	)
}

// Toggle calls the method "DOMTokenList.toggle".
func (this DOMTokenList) Toggle(token js.String, force bool) (ret bool) {
	bindings.CallDOMTokenListToggle(
		this.Ref(), js.Pointer(&ret),
		token.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// TryToggle calls the method "DOMTokenList.toggle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryToggle(token js.String, force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListToggle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// HasToggle1 returns true if the method "DOMTokenList.toggle" exists.
func (this DOMTokenList) HasToggle1() bool {
	return js.True == bindings.HasDOMTokenListToggle1(
		this.Ref(),
	)
}

// Toggle1Func returns the method "DOMTokenList.toggle".
func (this DOMTokenList) Toggle1Func() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListToggle1Func(
			this.Ref(),
		),
	)
}

// Toggle1 calls the method "DOMTokenList.toggle".
func (this DOMTokenList) Toggle1(token js.String) (ret bool) {
	bindings.CallDOMTokenListToggle1(
		this.Ref(), js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TryToggle1 calls the method "DOMTokenList.toggle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryToggle1(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListToggle1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
}

// HasReplace returns true if the method "DOMTokenList.replace" exists.
func (this DOMTokenList) HasReplace() bool {
	return js.True == bindings.HasDOMTokenListReplace(
		this.Ref(),
	)
}

// ReplaceFunc returns the method "DOMTokenList.replace".
func (this DOMTokenList) ReplaceFunc() (fn js.Func[func(token js.String, newToken js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListReplaceFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "DOMTokenList.replace".
func (this DOMTokenList) Replace(token js.String, newToken js.String) (ret bool) {
	bindings.CallDOMTokenListReplace(
		this.Ref(), js.Pointer(&ret),
		token.Ref(),
		newToken.Ref(),
	)

	return
}

// TryReplace calls the method "DOMTokenList.replace"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TryReplace(token js.String, newToken js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListReplace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		newToken.Ref(),
	)

	return
}

// HasSupports returns true if the method "DOMTokenList.supports" exists.
func (this DOMTokenList) HasSupports() bool {
	return js.True == bindings.HasDOMTokenListSupports(
		this.Ref(),
	)
}

// SupportsFunc returns the method "DOMTokenList.supports".
func (this DOMTokenList) SupportsFunc() (fn js.Func[func(token js.String) bool]) {
	return fn.FromRef(
		bindings.DOMTokenListSupportsFunc(
			this.Ref(),
		),
	)
}

// Supports calls the method "DOMTokenList.supports".
func (this DOMTokenList) Supports(token js.String) (ret bool) {
	bindings.CallDOMTokenListSupports(
		this.Ref(), js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TrySupports calls the method "DOMTokenList.supports"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMTokenList) TrySupports(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListSupports(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this NamedNodeMap) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNamedNodeMapLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "NamedNodeMap.item" exists.
func (this NamedNodeMap) HasItem() bool {
	return js.True == bindings.HasNamedNodeMapItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "NamedNodeMap.item".
func (this NamedNodeMap) ItemFunc() (fn js.Func[func(index uint32) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "NamedNodeMap.item".
func (this NamedNodeMap) Item(index uint32) (ret Attr) {
	bindings.CallNamedNodeMapItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "NamedNodeMap.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TryItem(index uint32) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasGetNamedItem returns true if the method "NamedNodeMap.getNamedItem" exists.
func (this NamedNodeMap) HasGetNamedItem() bool {
	return js.True == bindings.HasNamedNodeMapGetNamedItem(
		this.Ref(),
	)
}

// GetNamedItemFunc returns the method "NamedNodeMap.getNamedItem".
func (this NamedNodeMap) GetNamedItemFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapGetNamedItemFunc(
			this.Ref(),
		),
	)
}

// GetNamedItem calls the method "NamedNodeMap.getNamedItem".
func (this NamedNodeMap) GetNamedItem(qualifiedName js.String) (ret Attr) {
	bindings.CallNamedNodeMapGetNamedItem(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetNamedItem calls the method "NamedNodeMap.getNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TryGetNamedItem(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapGetNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasGetNamedItemNS returns true if the method "NamedNodeMap.getNamedItemNS" exists.
func (this NamedNodeMap) HasGetNamedItemNS() bool {
	return js.True == bindings.HasNamedNodeMapGetNamedItemNS(
		this.Ref(),
	)
}

// GetNamedItemNSFunc returns the method "NamedNodeMap.getNamedItemNS".
func (this NamedNodeMap) GetNamedItemNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapGetNamedItemNSFunc(
			this.Ref(),
		),
	)
}

// GetNamedItemNS calls the method "NamedNodeMap.getNamedItemNS".
func (this NamedNodeMap) GetNamedItemNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallNamedNodeMapGetNamedItemNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetNamedItemNS calls the method "NamedNodeMap.getNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TryGetNamedItemNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapGetNamedItemNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasSetNamedItem returns true if the method "NamedNodeMap.setNamedItem" exists.
func (this NamedNodeMap) HasSetNamedItem() bool {
	return js.True == bindings.HasNamedNodeMapSetNamedItem(
		this.Ref(),
	)
}

// SetNamedItemFunc returns the method "NamedNodeMap.setNamedItem".
func (this NamedNodeMap) SetNamedItemFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapSetNamedItemFunc(
			this.Ref(),
		),
	)
}

// SetNamedItem calls the method "NamedNodeMap.setNamedItem".
func (this NamedNodeMap) SetNamedItem(attr Attr) (ret Attr) {
	bindings.CallNamedNodeMapSetNamedItem(
		this.Ref(), js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetNamedItem calls the method "NamedNodeMap.setNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TrySetNamedItem(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapSetNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasSetNamedItemNS returns true if the method "NamedNodeMap.setNamedItemNS" exists.
func (this NamedNodeMap) HasSetNamedItemNS() bool {
	return js.True == bindings.HasNamedNodeMapSetNamedItemNS(
		this.Ref(),
	)
}

// SetNamedItemNSFunc returns the method "NamedNodeMap.setNamedItemNS".
func (this NamedNodeMap) SetNamedItemNSFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapSetNamedItemNSFunc(
			this.Ref(),
		),
	)
}

// SetNamedItemNS calls the method "NamedNodeMap.setNamedItemNS".
func (this NamedNodeMap) SetNamedItemNS(attr Attr) (ret Attr) {
	bindings.CallNamedNodeMapSetNamedItemNS(
		this.Ref(), js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetNamedItemNS calls the method "NamedNodeMap.setNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TrySetNamedItemNS(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapSetNamedItemNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasRemoveNamedItem returns true if the method "NamedNodeMap.removeNamedItem" exists.
func (this NamedNodeMap) HasRemoveNamedItem() bool {
	return js.True == bindings.HasNamedNodeMapRemoveNamedItem(
		this.Ref(),
	)
}

// RemoveNamedItemFunc returns the method "NamedNodeMap.removeNamedItem".
func (this NamedNodeMap) RemoveNamedItemFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapRemoveNamedItemFunc(
			this.Ref(),
		),
	)
}

// RemoveNamedItem calls the method "NamedNodeMap.removeNamedItem".
func (this NamedNodeMap) RemoveNamedItem(qualifiedName js.String) (ret Attr) {
	bindings.CallNamedNodeMapRemoveNamedItem(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryRemoveNamedItem calls the method "NamedNodeMap.removeNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TryRemoveNamedItem(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapRemoveNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasRemoveNamedItemNS returns true if the method "NamedNodeMap.removeNamedItemNS" exists.
func (this NamedNodeMap) HasRemoveNamedItemNS() bool {
	return js.True == bindings.HasNamedNodeMapRemoveNamedItemNS(
		this.Ref(),
	)
}

// RemoveNamedItemNSFunc returns the method "NamedNodeMap.removeNamedItemNS".
func (this NamedNodeMap) RemoveNamedItemNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.NamedNodeMapRemoveNamedItemNSFunc(
			this.Ref(),
		),
	)
}

// RemoveNamedItemNS calls the method "NamedNodeMap.removeNamedItemNS".
func (this NamedNodeMap) RemoveNamedItemNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallNamedNodeMapRemoveNamedItemNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryRemoveNamedItemNS calls the method "NamedNodeMap.removeNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedNodeMap) TryRemoveNamedItemNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapRemoveNamedItemNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Element) NamespaceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementNamespaceURI(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Prefix returns the value of property "Element.prefix".
//
// It returns ok=false if there is no such property.
func (this Element) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementPrefix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LocalName returns the value of property "Element.localName".
//
// It returns ok=false if there is no such property.
func (this Element) LocalName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementLocalName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TagName returns the value of property "Element.tagName".
//
// It returns ok=false if there is no such property.
func (this Element) TagName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementTagName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "Element.id".
//
// It returns ok=false if there is no such property.
func (this Element) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "Element.id" to val.
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
// It returns ok=false if there is no such property.
func (this Element) ClassName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementClassName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetClassName sets the value of property "Element.className" to val.
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
// It returns ok=false if there is no such property.
func (this Element) ClassList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetElementClassList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Slot returns the value of property "Element.slot".
//
// It returns ok=false if there is no such property.
func (this Element) Slot() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementSlot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSlot sets the value of property "Element.slot" to val.
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
// It returns ok=false if there is no such property.
func (this Element) Attributes() (ret NamedNodeMap, ok bool) {
	ok = js.True == bindings.GetElementAttributes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ShadowRoot returns the value of property "Element.shadowRoot".
//
// It returns ok=false if there is no such property.
func (this Element) ShadowRoot() (ret ShadowRoot, ok bool) {
	ok = js.True == bindings.GetElementShadowRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ElementTiming returns the value of property "Element.elementTiming".
//
// It returns ok=false if there is no such property.
func (this Element) ElementTiming() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementElementTiming(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetElementTiming sets the value of property "Element.elementTiming" to val.
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
// It returns ok=false if there is no such property.
func (this Element) Part() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetElementPart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OuterHTML returns the value of property "Element.outerHTML".
//
// It returns ok=false if there is no such property.
func (this Element) OuterHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementOuterHTML(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOuterHTML sets the value of property "Element.outerHTML" to val.
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
// It returns ok=false if there is no such property.
func (this Element) ScrollTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetElementScrollTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrollTop sets the value of property "Element.scrollTop" to val.
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
// It returns ok=false if there is no such property.
func (this Element) ScrollLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetElementScrollLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrollLeft sets the value of property "Element.scrollLeft" to val.
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
// It returns ok=false if there is no such property.
func (this Element) ScrollWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementScrollWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScrollHeight returns the value of property "Element.scrollHeight".
//
// It returns ok=false if there is no such property.
func (this Element) ScrollHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementScrollHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientTop returns the value of property "Element.clientTop".
//
// It returns ok=false if there is no such property.
func (this Element) ClientTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientLeft returns the value of property "Element.clientLeft".
//
// It returns ok=false if there is no such property.
func (this Element) ClientLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientWidth returns the value of property "Element.clientWidth".
//
// It returns ok=false if there is no such property.
func (this Element) ClientWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientHeight returns the value of property "Element.clientHeight".
//
// It returns ok=false if there is no such property.
func (this Element) ClientHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Role returns the value of property "Element.role".
//
// It returns ok=false if there is no such property.
func (this Element) Role() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementRole(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRole sets the value of property "Element.role" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaActiveDescendantElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementAriaActiveDescendantElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaActiveDescendantElement sets the value of property "Element.ariaActiveDescendantElement" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaAtomic() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaAtomic(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaAtomic sets the value of property "Element.ariaAtomic" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaAutoComplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaAutoComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaAutoComplete sets the value of property "Element.ariaAutoComplete" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaBusy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaBusy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaBusy sets the value of property "Element.ariaBusy" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaChecked() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaChecked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaChecked sets the value of property "Element.ariaChecked" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaColCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColCount sets the value of property "Element.ariaColCount" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaColIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColIndex sets the value of property "Element.ariaColIndex" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaColIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColIndexText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColIndexText sets the value of property "Element.ariaColIndexText" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaColSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColSpan sets the value of property "Element.ariaColSpan" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaControlsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaControlsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaControlsElements sets the value of property "Element.ariaControlsElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaCurrent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaCurrent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaCurrent sets the value of property "Element.ariaCurrent" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaDescribedByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaDescribedByElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDescribedByElements sets the value of property "Element.ariaDescribedByElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDescription sets the value of property "Element.ariaDescription" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaDetailsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaDetailsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDetailsElements sets the value of property "Element.ariaDetailsElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaDisabled() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDisabled sets the value of property "Element.ariaDisabled" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaErrorMessageElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaErrorMessageElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaErrorMessageElements sets the value of property "Element.ariaErrorMessageElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaExpanded() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaExpanded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaExpanded sets the value of property "Element.ariaExpanded" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaFlowToElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaFlowToElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaFlowToElements sets the value of property "Element.ariaFlowToElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaHasPopup() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaHasPopup(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaHasPopup sets the value of property "Element.ariaHasPopup" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaHidden() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaHidden(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaHidden sets the value of property "Element.ariaHidden" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaInvalid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaInvalid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaInvalid sets the value of property "Element.ariaInvalid" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaKeyShortcuts() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaKeyShortcuts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaKeyShortcuts sets the value of property "Element.ariaKeyShortcuts" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLabel sets the value of property "Element.ariaLabel" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaLabelledByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaLabelledByElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLabelledByElements sets the value of property "Element.ariaLabelledByElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaLevel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLevel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLevel sets the value of property "Element.ariaLevel" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaLive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLive sets the value of property "Element.ariaLive" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaModal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaModal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaModal sets the value of property "Element.ariaModal" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaMultiLine() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaMultiLine(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaMultiLine sets the value of property "Element.ariaMultiLine" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaMultiSelectable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaMultiSelectable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaMultiSelectable sets the value of property "Element.ariaMultiSelectable" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaOrientation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaOrientation sets the value of property "Element.ariaOrientation" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaOwnsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaOwnsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaOwnsElements sets the value of property "Element.ariaOwnsElements" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaPlaceholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPlaceholder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPlaceholder sets the value of property "Element.ariaPlaceholder" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaPosInSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPosInSet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPosInSet sets the value of property "Element.ariaPosInSet" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaPressed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPressed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPressed sets the value of property "Element.ariaPressed" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaReadOnly() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaReadOnly(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaReadOnly sets the value of property "Element.ariaReadOnly" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRequired() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRequired(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRequired sets the value of property "Element.ariaRequired" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRoleDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRoleDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRoleDescription sets the value of property "Element.ariaRoleDescription" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRowCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowCount sets the value of property "Element.ariaRowCount" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRowIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndex sets the value of property "Element.ariaRowIndex" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRowIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowIndexText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndexText sets the value of property "Element.ariaRowIndexText" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaRowSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowSpan sets the value of property "Element.ariaRowSpan" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaSelected() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSelected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSelected sets the value of property "Element.ariaSelected" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaSetSize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSetSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSetSize sets the value of property "Element.ariaSetSize" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaSort() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSort sets the value of property "Element.ariaSort" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaValueMax() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueMax sets the value of property "Element.ariaValueMax" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaValueMin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueMin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueMin sets the value of property "Element.ariaValueMin" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaValueNow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueNow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueNow sets the value of property "Element.ariaValueNow" to val.
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
// It returns ok=false if there is no such property.
func (this Element) AriaValueText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueText sets the value of property "Element.ariaValueText" to val.
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
// It returns ok=false if there is no such property.
func (this Element) InnerHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInnerHTML(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInnerHTML sets the value of property "Element.innerHTML" to val.
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
// It returns ok=false if there is no such property.
func (this Element) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetElementChildren(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "Element.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this Element) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementFirstElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "Element.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this Element) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementLastElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "Element.childElementCount".
//
// It returns ok=false if there is no such property.
func (this Element) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetElementChildElementCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousElementSibling returns the value of property "Element.previousElementSibling".
//
// It returns ok=false if there is no such property.
func (this Element) PreviousElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementPreviousElementSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextElementSibling returns the value of property "Element.nextElementSibling".
//
// It returns ok=false if there is no such property.
func (this Element) NextElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementNextElementSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AssignedSlot returns the value of property "Element.assignedSlot".
//
// It returns ok=false if there is no such property.
func (this Element) AssignedSlot() (ret HTMLSlotElement, ok bool) {
	ok = js.True == bindings.GetElementAssignedSlot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RegionOverset returns the value of property "Element.regionOverset".
//
// It returns ok=false if there is no such property.
func (this Element) RegionOverset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementRegionOverset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasHasAttributes returns true if the method "Element.hasAttributes" exists.
func (this Element) HasHasAttributes() bool {
	return js.True == bindings.HasElementHasAttributes(
		this.Ref(),
	)
}

// HasAttributesFunc returns the method "Element.hasAttributes".
func (this Element) HasAttributesFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributesFunc(
			this.Ref(),
		),
	)
}

// HasAttributes calls the method "Element.hasAttributes".
func (this Element) HasAttributes() (ret bool) {
	bindings.CallElementHasAttributes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHasAttributes calls the method "Element.hasAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryHasAttributes() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttributes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAttributeNames returns true if the method "Element.getAttributeNames" exists.
func (this Element) HasGetAttributeNames() bool {
	return js.True == bindings.HasElementGetAttributeNames(
		this.Ref(),
	)
}

// GetAttributeNamesFunc returns the method "Element.getAttributeNames".
func (this Element) GetAttributeNamesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNamesFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNames calls the method "Element.getAttributeNames".
func (this Element) GetAttributeNames() (ret js.Array[js.String]) {
	bindings.CallElementGetAttributeNames(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAttributeNames calls the method "Element.getAttributeNames"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAttributeNames() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNames(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAttribute returns true if the method "Element.getAttribute" exists.
func (this Element) HasGetAttribute() bool {
	return js.True == bindings.HasElementGetAttribute(
		this.Ref(),
	)
}

// GetAttributeFunc returns the method "Element.getAttribute".
func (this Element) GetAttributeFunc() (fn js.Func[func(qualifiedName js.String) js.String]) {
	return fn.FromRef(
		bindings.ElementGetAttributeFunc(
			this.Ref(),
		),
	)
}

// GetAttribute calls the method "Element.getAttribute".
func (this Element) GetAttribute(qualifiedName js.String) (ret js.String) {
	bindings.CallElementGetAttribute(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetAttribute calls the method "Element.getAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAttribute(qualifiedName js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasGetAttributeNS returns true if the method "Element.getAttributeNS" exists.
func (this Element) HasGetAttributeNS() bool {
	return js.True == bindings.HasElementGetAttributeNS(
		this.Ref(),
	)
}

// GetAttributeNSFunc returns the method "Element.getAttributeNS".
func (this Element) GetAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) js.String]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNSFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNS calls the method "Element.getAttributeNS".
func (this Element) GetAttributeNS(namespace js.String, localName js.String) (ret js.String) {
	bindings.CallElementGetAttributeNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetAttributeNS calls the method "Element.getAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAttributeNS(namespace js.String, localName js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasSetAttribute returns true if the method "Element.setAttribute" exists.
func (this Element) HasSetAttribute() bool {
	return js.True == bindings.HasElementSetAttribute(
		this.Ref(),
	)
}

// SetAttributeFunc returns the method "Element.setAttribute".
func (this Element) SetAttributeFunc() (fn js.Func[func(qualifiedName js.String, value js.String)]) {
	return fn.FromRef(
		bindings.ElementSetAttributeFunc(
			this.Ref(),
		),
	)
}

// SetAttribute calls the method "Element.setAttribute".
func (this Element) SetAttribute(qualifiedName js.String, value js.String) (ret js.Void) {
	bindings.CallElementSetAttribute(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// TrySetAttribute calls the method "Element.setAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetAttribute(qualifiedName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// HasSetAttributeNS returns true if the method "Element.setAttributeNS" exists.
func (this Element) HasSetAttributeNS() bool {
	return js.True == bindings.HasElementSetAttributeNS(
		this.Ref(),
	)
}

// SetAttributeNSFunc returns the method "Element.setAttributeNS".
func (this Element) SetAttributeNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, value js.String)]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNSFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNS calls the method "Element.setAttributeNS".
func (this Element) SetAttributeNS(namespace js.String, qualifiedName js.String, value js.String) (ret js.Void) {
	bindings.CallElementSetAttributeNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// TrySetAttributeNS calls the method "Element.setAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetAttributeNS(namespace js.String, qualifiedName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// HasRemoveAttribute returns true if the method "Element.removeAttribute" exists.
func (this Element) HasRemoveAttribute() bool {
	return js.True == bindings.HasElementRemoveAttribute(
		this.Ref(),
	)
}

// RemoveAttributeFunc returns the method "Element.removeAttribute".
func (this Element) RemoveAttributeFunc() (fn js.Func[func(qualifiedName js.String)]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeFunc(
			this.Ref(),
		),
	)
}

// RemoveAttribute calls the method "Element.removeAttribute".
func (this Element) RemoveAttribute(qualifiedName js.String) (ret js.Void) {
	bindings.CallElementRemoveAttribute(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryRemoveAttribute calls the method "Element.removeAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRemoveAttribute(qualifiedName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasRemoveAttributeNS returns true if the method "Element.removeAttributeNS" exists.
func (this Element) HasRemoveAttributeNS() bool {
	return js.True == bindings.HasElementRemoveAttributeNS(
		this.Ref(),
	)
}

// RemoveAttributeNSFunc returns the method "Element.removeAttributeNS".
func (this Element) RemoveAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String)]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeNSFunc(
			this.Ref(),
		),
	)
}

// RemoveAttributeNS calls the method "Element.removeAttributeNS".
func (this Element) RemoveAttributeNS(namespace js.String, localName js.String) (ret js.Void) {
	bindings.CallElementRemoveAttributeNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryRemoveAttributeNS calls the method "Element.removeAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRemoveAttributeNS(namespace js.String, localName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttributeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasToggleAttribute returns true if the method "Element.toggleAttribute" exists.
func (this Element) HasToggleAttribute() bool {
	return js.True == bindings.HasElementToggleAttribute(
		this.Ref(),
	)
}

// ToggleAttributeFunc returns the method "Element.toggleAttribute".
func (this Element) ToggleAttributeFunc() (fn js.Func[func(qualifiedName js.String, force bool) bool]) {
	return fn.FromRef(
		bindings.ElementToggleAttributeFunc(
			this.Ref(),
		),
	)
}

// ToggleAttribute calls the method "Element.toggleAttribute".
func (this Element) ToggleAttribute(qualifiedName js.String, force bool) (ret bool) {
	bindings.CallElementToggleAttribute(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// TryToggleAttribute calls the method "Element.toggleAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryToggleAttribute(qualifiedName js.String, force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementToggleAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// HasToggleAttribute1 returns true if the method "Element.toggleAttribute" exists.
func (this Element) HasToggleAttribute1() bool {
	return js.True == bindings.HasElementToggleAttribute1(
		this.Ref(),
	)
}

// ToggleAttribute1Func returns the method "Element.toggleAttribute".
func (this Element) ToggleAttribute1Func() (fn js.Func[func(qualifiedName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementToggleAttribute1Func(
			this.Ref(),
		),
	)
}

// ToggleAttribute1 calls the method "Element.toggleAttribute".
func (this Element) ToggleAttribute1(qualifiedName js.String) (ret bool) {
	bindings.CallElementToggleAttribute1(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryToggleAttribute1 calls the method "Element.toggleAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryToggleAttribute1(qualifiedName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementToggleAttribute1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasHasAttribute returns true if the method "Element.hasAttribute" exists.
func (this Element) HasHasAttribute() bool {
	return js.True == bindings.HasElementHasAttribute(
		this.Ref(),
	)
}

// HasAttributeFunc returns the method "Element.hasAttribute".
func (this Element) HasAttributeFunc() (fn js.Func[func(qualifiedName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributeFunc(
			this.Ref(),
		),
	)
}

// HasAttribute calls the method "Element.hasAttribute".
func (this Element) HasAttribute(qualifiedName js.String) (ret bool) {
	bindings.CallElementHasAttribute(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryHasAttribute calls the method "Element.hasAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryHasAttribute(qualifiedName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasHasAttributeNS returns true if the method "Element.hasAttributeNS" exists.
func (this Element) HasHasAttributeNS() bool {
	return js.True == bindings.HasElementHasAttributeNS(
		this.Ref(),
	)
}

// HasAttributeNSFunc returns the method "Element.hasAttributeNS".
func (this Element) HasAttributeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) bool]) {
	return fn.FromRef(
		bindings.ElementHasAttributeNSFunc(
			this.Ref(),
		),
	)
}

// HasAttributeNS calls the method "Element.hasAttributeNS".
func (this Element) HasAttributeNS(namespace js.String, localName js.String) (ret bool) {
	bindings.CallElementHasAttributeNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryHasAttributeNS calls the method "Element.hasAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryHasAttributeNS(namespace js.String, localName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttributeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasGetAttributeNode returns true if the method "Element.getAttributeNode" exists.
func (this Element) HasGetAttributeNode() bool {
	return js.True == bindings.HasElementGetAttributeNode(
		this.Ref(),
	)
}

// GetAttributeNodeFunc returns the method "Element.getAttributeNode".
func (this Element) GetAttributeNodeFunc() (fn js.Func[func(qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNode calls the method "Element.getAttributeNode".
func (this Element) GetAttributeNode(qualifiedName js.String) (ret Attr) {
	bindings.CallElementGetAttributeNode(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetAttributeNode calls the method "Element.getAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAttributeNode(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasGetAttributeNodeNS returns true if the method "Element.getAttributeNodeNS" exists.
func (this Element) HasGetAttributeNodeNS() bool {
	return js.True == bindings.HasElementGetAttributeNodeNS(
		this.Ref(),
	)
}

// GetAttributeNodeNSFunc returns the method "Element.getAttributeNodeNS".
func (this Element) GetAttributeNodeNSFunc() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	return fn.FromRef(
		bindings.ElementGetAttributeNodeNSFunc(
			this.Ref(),
		),
	)
}

// GetAttributeNodeNS calls the method "Element.getAttributeNodeNS".
func (this Element) GetAttributeNodeNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallElementGetAttributeNodeNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetAttributeNodeNS calls the method "Element.getAttributeNodeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAttributeNodeNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNodeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasSetAttributeNode returns true if the method "Element.setAttributeNode" exists.
func (this Element) HasSetAttributeNode() bool {
	return js.True == bindings.HasElementSetAttributeNode(
		this.Ref(),
	)
}

// SetAttributeNodeFunc returns the method "Element.setAttributeNode".
func (this Element) SetAttributeNodeFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNode calls the method "Element.setAttributeNode".
func (this Element) SetAttributeNode(attr Attr) (ret Attr) {
	bindings.CallElementSetAttributeNode(
		this.Ref(), js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetAttributeNode calls the method "Element.setAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetAttributeNode(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasSetAttributeNodeNS returns true if the method "Element.setAttributeNodeNS" exists.
func (this Element) HasSetAttributeNodeNS() bool {
	return js.True == bindings.HasElementSetAttributeNodeNS(
		this.Ref(),
	)
}

// SetAttributeNodeNSFunc returns the method "Element.setAttributeNodeNS".
func (this Element) SetAttributeNodeNSFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementSetAttributeNodeNSFunc(
			this.Ref(),
		),
	)
}

// SetAttributeNodeNS calls the method "Element.setAttributeNodeNS".
func (this Element) SetAttributeNodeNS(attr Attr) (ret Attr) {
	bindings.CallElementSetAttributeNodeNS(
		this.Ref(), js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetAttributeNodeNS calls the method "Element.setAttributeNodeNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetAttributeNodeNS(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNodeNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasRemoveAttributeNode returns true if the method "Element.removeAttributeNode" exists.
func (this Element) HasRemoveAttributeNode() bool {
	return js.True == bindings.HasElementRemoveAttributeNode(
		this.Ref(),
	)
}

// RemoveAttributeNodeFunc returns the method "Element.removeAttributeNode".
func (this Element) RemoveAttributeNodeFunc() (fn js.Func[func(attr Attr) Attr]) {
	return fn.FromRef(
		bindings.ElementRemoveAttributeNodeFunc(
			this.Ref(),
		),
	)
}

// RemoveAttributeNode calls the method "Element.removeAttributeNode".
func (this Element) RemoveAttributeNode(attr Attr) (ret Attr) {
	bindings.CallElementRemoveAttributeNode(
		this.Ref(), js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TryRemoveAttributeNode calls the method "Element.removeAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRemoveAttributeNode(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttributeNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasAttachShadow returns true if the method "Element.attachShadow" exists.
func (this Element) HasAttachShadow() bool {
	return js.True == bindings.HasElementAttachShadow(
		this.Ref(),
	)
}

// AttachShadowFunc returns the method "Element.attachShadow".
func (this Element) AttachShadowFunc() (fn js.Func[func(init ShadowRootInit) ShadowRoot]) {
	return fn.FromRef(
		bindings.ElementAttachShadowFunc(
			this.Ref(),
		),
	)
}

// AttachShadow calls the method "Element.attachShadow".
func (this Element) AttachShadow(init ShadowRootInit) (ret ShadowRoot) {
	bindings.CallElementAttachShadow(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryAttachShadow calls the method "Element.attachShadow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryAttachShadow(init ShadowRootInit) (ret ShadowRoot, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAttachShadow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasClosest returns true if the method "Element.closest" exists.
func (this Element) HasClosest() bool {
	return js.True == bindings.HasElementClosest(
		this.Ref(),
	)
}

// ClosestFunc returns the method "Element.closest".
func (this Element) ClosestFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.ElementClosestFunc(
			this.Ref(),
		),
	)
}

// Closest calls the method "Element.closest".
func (this Element) Closest(selectors js.String) (ret Element) {
	bindings.CallElementClosest(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryClosest calls the method "Element.closest"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryClosest(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementClosest(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasMatches returns true if the method "Element.matches" exists.
func (this Element) HasMatches() bool {
	return js.True == bindings.HasElementMatches(
		this.Ref(),
	)
}

// MatchesFunc returns the method "Element.matches".
func (this Element) MatchesFunc() (fn js.Func[func(selectors js.String) bool]) {
	return fn.FromRef(
		bindings.ElementMatchesFunc(
			this.Ref(),
		),
	)
}

// Matches calls the method "Element.matches".
func (this Element) Matches(selectors js.String) (ret bool) {
	bindings.CallElementMatches(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryMatches calls the method "Element.matches"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryMatches(selectors js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementMatches(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasWebkitMatchesSelector returns true if the method "Element.webkitMatchesSelector" exists.
func (this Element) HasWebkitMatchesSelector() bool {
	return js.True == bindings.HasElementWebkitMatchesSelector(
		this.Ref(),
	)
}

// WebkitMatchesSelectorFunc returns the method "Element.webkitMatchesSelector".
func (this Element) WebkitMatchesSelectorFunc() (fn js.Func[func(selectors js.String) bool]) {
	return fn.FromRef(
		bindings.ElementWebkitMatchesSelectorFunc(
			this.Ref(),
		),
	)
}

// WebkitMatchesSelector calls the method "Element.webkitMatchesSelector".
func (this Element) WebkitMatchesSelector(selectors js.String) (ret bool) {
	bindings.CallElementWebkitMatchesSelector(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryWebkitMatchesSelector calls the method "Element.webkitMatchesSelector"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryWebkitMatchesSelector(selectors js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementWebkitMatchesSelector(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasGetElementsByTagName returns true if the method "Element.getElementsByTagName" exists.
func (this Element) HasGetElementsByTagName() bool {
	return js.True == bindings.HasElementGetElementsByTagName(
		this.Ref(),
	)
}

// GetElementsByTagNameFunc returns the method "Element.getElementsByTagName".
func (this Element) GetElementsByTagNameFunc() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByTagNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagName calls the method "Element.getElementsByTagName".
func (this Element) GetElementsByTagName(qualifiedName js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByTagName(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetElementsByTagName calls the method "Element.getElementsByTagName"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetElementsByTagName(qualifiedName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByTagName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasGetElementsByTagNameNS returns true if the method "Element.getElementsByTagNameNS" exists.
func (this Element) HasGetElementsByTagNameNS() bool {
	return js.True == bindings.HasElementGetElementsByTagNameNS(
		this.Ref(),
	)
}

// GetElementsByTagNameNSFunc returns the method "Element.getElementsByTagNameNS".
func (this Element) GetElementsByTagNameNSFunc() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByTagNameNSFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagNameNS calls the method "Element.getElementsByTagNameNS".
func (this Element) GetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByTagNameNS(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetElementsByTagNameNS calls the method "Element.getElementsByTagNameNS"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByTagNameNS(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasGetElementsByClassName returns true if the method "Element.getElementsByClassName" exists.
func (this Element) HasGetElementsByClassName() bool {
	return js.True == bindings.HasElementGetElementsByClassName(
		this.Ref(),
	)
}

// GetElementsByClassNameFunc returns the method "Element.getElementsByClassName".
func (this Element) GetElementsByClassNameFunc() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.ElementGetElementsByClassNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByClassName calls the method "Element.getElementsByClassName".
func (this Element) GetElementsByClassName(classNames js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByClassName(
		this.Ref(), js.Pointer(&ret),
		classNames.Ref(),
	)

	return
}

// TryGetElementsByClassName calls the method "Element.getElementsByClassName"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetElementsByClassName(classNames js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByClassName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		classNames.Ref(),
	)

	return
}

// HasInsertAdjacentElement returns true if the method "Element.insertAdjacentElement" exists.
func (this Element) HasInsertAdjacentElement() bool {
	return js.True == bindings.HasElementInsertAdjacentElement(
		this.Ref(),
	)
}

// InsertAdjacentElementFunc returns the method "Element.insertAdjacentElement".
func (this Element) InsertAdjacentElementFunc() (fn js.Func[func(where js.String, element Element) Element]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentElementFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentElement calls the method "Element.insertAdjacentElement".
func (this Element) InsertAdjacentElement(where js.String, element Element) (ret Element) {
	bindings.CallElementInsertAdjacentElement(
		this.Ref(), js.Pointer(&ret),
		where.Ref(),
		element.Ref(),
	)

	return
}

// TryInsertAdjacentElement calls the method "Element.insertAdjacentElement"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryInsertAdjacentElement(where js.String, element Element) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		where.Ref(),
		element.Ref(),
	)

	return
}

// HasInsertAdjacentText returns true if the method "Element.insertAdjacentText" exists.
func (this Element) HasInsertAdjacentText() bool {
	return js.True == bindings.HasElementInsertAdjacentText(
		this.Ref(),
	)
}

// InsertAdjacentTextFunc returns the method "Element.insertAdjacentText".
func (this Element) InsertAdjacentTextFunc() (fn js.Func[func(where js.String, data js.String)]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentTextFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentText calls the method "Element.insertAdjacentText".
func (this Element) InsertAdjacentText(where js.String, data js.String) (ret js.Void) {
	bindings.CallElementInsertAdjacentText(
		this.Ref(), js.Pointer(&ret),
		where.Ref(),
		data.Ref(),
	)

	return
}

// TryInsertAdjacentText calls the method "Element.insertAdjacentText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryInsertAdjacentText(where js.String, data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		where.Ref(),
		data.Ref(),
	)

	return
}

// HasRequestFullscreen returns true if the method "Element.requestFullscreen" exists.
func (this Element) HasRequestFullscreen() bool {
	return js.True == bindings.HasElementRequestFullscreen(
		this.Ref(),
	)
}

// RequestFullscreenFunc returns the method "Element.requestFullscreen".
func (this Element) RequestFullscreenFunc() (fn js.Func[func(options FullscreenOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ElementRequestFullscreenFunc(
			this.Ref(),
		),
	)
}

// RequestFullscreen calls the method "Element.requestFullscreen".
func (this Element) RequestFullscreen(options FullscreenOptions) (ret js.Promise[js.Void]) {
	bindings.CallElementRequestFullscreen(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestFullscreen calls the method "Element.requestFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRequestFullscreen(options FullscreenOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestFullscreen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestFullscreen1 returns true if the method "Element.requestFullscreen" exists.
func (this Element) HasRequestFullscreen1() bool {
	return js.True == bindings.HasElementRequestFullscreen1(
		this.Ref(),
	)
}

// RequestFullscreen1Func returns the method "Element.requestFullscreen".
func (this Element) RequestFullscreen1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ElementRequestFullscreen1Func(
			this.Ref(),
		),
	)
}

// RequestFullscreen1 calls the method "Element.requestFullscreen".
func (this Element) RequestFullscreen1() (ret js.Promise[js.Void]) {
	bindings.CallElementRequestFullscreen1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestFullscreen1 calls the method "Element.requestFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRequestFullscreen1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestFullscreen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestPointerLock returns true if the method "Element.requestPointerLock" exists.
func (this Element) HasRequestPointerLock() bool {
	return js.True == bindings.HasElementRequestPointerLock(
		this.Ref(),
	)
}

// RequestPointerLockFunc returns the method "Element.requestPointerLock".
func (this Element) RequestPointerLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementRequestPointerLockFunc(
			this.Ref(),
		),
	)
}

// RequestPointerLock calls the method "Element.requestPointerLock".
func (this Element) RequestPointerLock() (ret js.Void) {
	bindings.CallElementRequestPointerLock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPointerLock calls the method "Element.requestPointerLock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRequestPointerLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestPointerLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasComputedStyleMap returns true if the method "Element.computedStyleMap" exists.
func (this Element) HasComputedStyleMap() bool {
	return js.True == bindings.HasElementComputedStyleMap(
		this.Ref(),
	)
}

// ComputedStyleMapFunc returns the method "Element.computedStyleMap".
func (this Element) ComputedStyleMapFunc() (fn js.Func[func() StylePropertyMapReadOnly]) {
	return fn.FromRef(
		bindings.ElementComputedStyleMapFunc(
			this.Ref(),
		),
	)
}

// ComputedStyleMap calls the method "Element.computedStyleMap".
func (this Element) ComputedStyleMap() (ret StylePropertyMapReadOnly) {
	bindings.CallElementComputedStyleMap(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryComputedStyleMap calls the method "Element.computedStyleMap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryComputedStyleMap() (ret StylePropertyMapReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementComputedStyleMap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSpatialNavigationContainer returns true if the method "Element.getSpatialNavigationContainer" exists.
func (this Element) HasGetSpatialNavigationContainer() bool {
	return js.True == bindings.HasElementGetSpatialNavigationContainer(
		this.Ref(),
	)
}

// GetSpatialNavigationContainerFunc returns the method "Element.getSpatialNavigationContainer".
func (this Element) GetSpatialNavigationContainerFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.ElementGetSpatialNavigationContainerFunc(
			this.Ref(),
		),
	)
}

// GetSpatialNavigationContainer calls the method "Element.getSpatialNavigationContainer".
func (this Element) GetSpatialNavigationContainer() (ret Node) {
	bindings.CallElementGetSpatialNavigationContainer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSpatialNavigationContainer calls the method "Element.getSpatialNavigationContainer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetSpatialNavigationContainer() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetSpatialNavigationContainer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFocusableAreas returns true if the method "Element.focusableAreas" exists.
func (this Element) HasFocusableAreas() bool {
	return js.True == bindings.HasElementFocusableAreas(
		this.Ref(),
	)
}

// FocusableAreasFunc returns the method "Element.focusableAreas".
func (this Element) FocusableAreasFunc() (fn js.Func[func(option FocusableAreasOption) js.Array[Node]]) {
	return fn.FromRef(
		bindings.ElementFocusableAreasFunc(
			this.Ref(),
		),
	)
}

// FocusableAreas calls the method "Element.focusableAreas".
func (this Element) FocusableAreas(option FocusableAreasOption) (ret js.Array[Node]) {
	bindings.CallElementFocusableAreas(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&option),
	)

	return
}

// TryFocusableAreas calls the method "Element.focusableAreas"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryFocusableAreas(option FocusableAreasOption) (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementFocusableAreas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&option),
	)

	return
}

// HasFocusableAreas1 returns true if the method "Element.focusableAreas" exists.
func (this Element) HasFocusableAreas1() bool {
	return js.True == bindings.HasElementFocusableAreas1(
		this.Ref(),
	)
}

// FocusableAreas1Func returns the method "Element.focusableAreas".
func (this Element) FocusableAreas1Func() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.ElementFocusableAreas1Func(
			this.Ref(),
		),
	)
}

// FocusableAreas1 calls the method "Element.focusableAreas".
func (this Element) FocusableAreas1() (ret js.Array[Node]) {
	bindings.CallElementFocusableAreas1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocusableAreas1 calls the method "Element.focusableAreas"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryFocusableAreas1() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementFocusableAreas1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSpatialNavigationSearch returns true if the method "Element.spatialNavigationSearch" exists.
func (this Element) HasSpatialNavigationSearch() bool {
	return js.True == bindings.HasElementSpatialNavigationSearch(
		this.Ref(),
	)
}

// SpatialNavigationSearchFunc returns the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearchFunc() (fn js.Func[func(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) Node]) {
	return fn.FromRef(
		bindings.ElementSpatialNavigationSearchFunc(
			this.Ref(),
		),
	)
}

// SpatialNavigationSearch calls the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearch(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) (ret Node) {
	bindings.CallElementSpatialNavigationSearch(
		this.Ref(), js.Pointer(&ret),
		uint32(dir),
		js.Pointer(&options),
	)

	return
}

// TrySpatialNavigationSearch calls the method "Element.spatialNavigationSearch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySpatialNavigationSearch(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSpatialNavigationSearch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
		js.Pointer(&options),
	)

	return
}

// HasSpatialNavigationSearch1 returns true if the method "Element.spatialNavigationSearch" exists.
func (this Element) HasSpatialNavigationSearch1() bool {
	return js.True == bindings.HasElementSpatialNavigationSearch1(
		this.Ref(),
	)
}

// SpatialNavigationSearch1Func returns the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearch1Func() (fn js.Func[func(dir SpatialNavigationDirection) Node]) {
	return fn.FromRef(
		bindings.ElementSpatialNavigationSearch1Func(
			this.Ref(),
		),
	)
}

// SpatialNavigationSearch1 calls the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearch1(dir SpatialNavigationDirection) (ret Node) {
	bindings.CallElementSpatialNavigationSearch1(
		this.Ref(), js.Pointer(&ret),
		uint32(dir),
	)

	return
}

// TrySpatialNavigationSearch1 calls the method "Element.spatialNavigationSearch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySpatialNavigationSearch1(dir SpatialNavigationDirection) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSpatialNavigationSearch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
	)

	return
}

// HasSetPointerCapture returns true if the method "Element.setPointerCapture" exists.
func (this Element) HasSetPointerCapture() bool {
	return js.True == bindings.HasElementSetPointerCapture(
		this.Ref(),
	)
}

// SetPointerCaptureFunc returns the method "Element.setPointerCapture".
func (this Element) SetPointerCaptureFunc() (fn js.Func[func(pointerId int32)]) {
	return fn.FromRef(
		bindings.ElementSetPointerCaptureFunc(
			this.Ref(),
		),
	)
}

// SetPointerCapture calls the method "Element.setPointerCapture".
func (this Element) SetPointerCapture(pointerId int32) (ret js.Void) {
	bindings.CallElementSetPointerCapture(
		this.Ref(), js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TrySetPointerCapture calls the method "Element.setPointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetPointerCapture(pointerId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetPointerCapture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasReleasePointerCapture returns true if the method "Element.releasePointerCapture" exists.
func (this Element) HasReleasePointerCapture() bool {
	return js.True == bindings.HasElementReleasePointerCapture(
		this.Ref(),
	)
}

// ReleasePointerCaptureFunc returns the method "Element.releasePointerCapture".
func (this Element) ReleasePointerCaptureFunc() (fn js.Func[func(pointerId int32)]) {
	return fn.FromRef(
		bindings.ElementReleasePointerCaptureFunc(
			this.Ref(),
		),
	)
}

// ReleasePointerCapture calls the method "Element.releasePointerCapture".
func (this Element) ReleasePointerCapture(pointerId int32) (ret js.Void) {
	bindings.CallElementReleasePointerCapture(
		this.Ref(), js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TryReleasePointerCapture calls the method "Element.releasePointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryReleasePointerCapture(pointerId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReleasePointerCapture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasHasPointerCapture returns true if the method "Element.hasPointerCapture" exists.
func (this Element) HasHasPointerCapture() bool {
	return js.True == bindings.HasElementHasPointerCapture(
		this.Ref(),
	)
}

// HasPointerCaptureFunc returns the method "Element.hasPointerCapture".
func (this Element) HasPointerCaptureFunc() (fn js.Func[func(pointerId int32) bool]) {
	return fn.FromRef(
		bindings.ElementHasPointerCaptureFunc(
			this.Ref(),
		),
	)
}

// HasPointerCapture calls the method "Element.hasPointerCapture".
func (this Element) HasPointerCapture(pointerId int32) (ret bool) {
	bindings.CallElementHasPointerCapture(
		this.Ref(), js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TryHasPointerCapture calls the method "Element.hasPointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryHasPointerCapture(pointerId int32) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasPointerCapture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasPseudo returns true if the method "Element.pseudo" exists.
func (this Element) HasPseudo() bool {
	return js.True == bindings.HasElementPseudo(
		this.Ref(),
	)
}

// PseudoFunc returns the method "Element.pseudo".
func (this Element) PseudoFunc() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	return fn.FromRef(
		bindings.ElementPseudoFunc(
			this.Ref(),
		),
	)
}

// Pseudo calls the method "Element.pseudo".
func (this Element) Pseudo(typ js.String) (ret CSSPseudoElement) {
	bindings.CallElementPseudo(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryPseudo calls the method "Element.pseudo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryPseudo(typ js.String) (ret CSSPseudoElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementPseudo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasInsertAdjacentHTML returns true if the method "Element.insertAdjacentHTML" exists.
func (this Element) HasInsertAdjacentHTML() bool {
	return js.True == bindings.HasElementInsertAdjacentHTML(
		this.Ref(),
	)
}

// InsertAdjacentHTMLFunc returns the method "Element.insertAdjacentHTML".
func (this Element) InsertAdjacentHTMLFunc() (fn js.Func[func(position js.String, text js.String)]) {
	return fn.FromRef(
		bindings.ElementInsertAdjacentHTMLFunc(
			this.Ref(),
		),
	)
}

// InsertAdjacentHTML calls the method "Element.insertAdjacentHTML".
func (this Element) InsertAdjacentHTML(position js.String, text js.String) (ret js.Void) {
	bindings.CallElementInsertAdjacentHTML(
		this.Ref(), js.Pointer(&ret),
		position.Ref(),
		text.Ref(),
	)

	return
}

// TryInsertAdjacentHTML calls the method "Element.insertAdjacentHTML"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryInsertAdjacentHTML(position js.String, text js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentHTML(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		position.Ref(),
		text.Ref(),
	)

	return
}

// HasSetHTML returns true if the method "Element.setHTML" exists.
func (this Element) HasSetHTML() bool {
	return js.True == bindings.HasElementSetHTML(
		this.Ref(),
	)
}

// SetHTMLFunc returns the method "Element.setHTML".
func (this Element) SetHTMLFunc() (fn js.Func[func(input js.String, options SetHTMLOptions)]) {
	return fn.FromRef(
		bindings.ElementSetHTMLFunc(
			this.Ref(),
		),
	)
}

// SetHTML calls the method "Element.setHTML".
func (this Element) SetHTML(input js.String, options SetHTMLOptions) (ret js.Void) {
	bindings.CallElementSetHTML(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySetHTML calls the method "Element.setHTML"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetHTML(input js.String, options SetHTMLOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetHTML(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSetHTML1 returns true if the method "Element.setHTML" exists.
func (this Element) HasSetHTML1() bool {
	return js.True == bindings.HasElementSetHTML1(
		this.Ref(),
	)
}

// SetHTML1Func returns the method "Element.setHTML".
func (this Element) SetHTML1Func() (fn js.Func[func(input js.String)]) {
	return fn.FromRef(
		bindings.ElementSetHTML1Func(
			this.Ref(),
		),
	)
}

// SetHTML1 calls the method "Element.setHTML".
func (this Element) SetHTML1(input js.String) (ret js.Void) {
	bindings.CallElementSetHTML1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySetHTML1 calls the method "Element.setHTML"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TrySetHTML1(input js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetHTML1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasGetClientRects returns true if the method "Element.getClientRects" exists.
func (this Element) HasGetClientRects() bool {
	return js.True == bindings.HasElementGetClientRects(
		this.Ref(),
	)
}

// GetClientRectsFunc returns the method "Element.getClientRects".
func (this Element) GetClientRectsFunc() (fn js.Func[func() DOMRectList]) {
	return fn.FromRef(
		bindings.ElementGetClientRectsFunc(
			this.Ref(),
		),
	)
}

// GetClientRects calls the method "Element.getClientRects".
func (this Element) GetClientRects() (ret DOMRectList) {
	bindings.CallElementGetClientRects(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetClientRects calls the method "Element.getClientRects"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetClientRects() (ret DOMRectList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetClientRects(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetBoundingClientRect returns true if the method "Element.getBoundingClientRect" exists.
func (this Element) HasGetBoundingClientRect() bool {
	return js.True == bindings.HasElementGetBoundingClientRect(
		this.Ref(),
	)
}

// GetBoundingClientRectFunc returns the method "Element.getBoundingClientRect".
func (this Element) GetBoundingClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.ElementGetBoundingClientRectFunc(
			this.Ref(),
		),
	)
}

// GetBoundingClientRect calls the method "Element.getBoundingClientRect".
func (this Element) GetBoundingClientRect() (ret DOMRect) {
	bindings.CallElementGetBoundingClientRect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoundingClientRect calls the method "Element.getBoundingClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetBoundingClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoundingClientRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCheckVisibility returns true if the method "Element.checkVisibility" exists.
func (this Element) HasCheckVisibility() bool {
	return js.True == bindings.HasElementCheckVisibility(
		this.Ref(),
	)
}

// CheckVisibilityFunc returns the method "Element.checkVisibility".
func (this Element) CheckVisibilityFunc() (fn js.Func[func(options CheckVisibilityOptions) bool]) {
	return fn.FromRef(
		bindings.ElementCheckVisibilityFunc(
			this.Ref(),
		),
	)
}

// CheckVisibility calls the method "Element.checkVisibility".
func (this Element) CheckVisibility(options CheckVisibilityOptions) (ret bool) {
	bindings.CallElementCheckVisibility(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCheckVisibility calls the method "Element.checkVisibility"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryCheckVisibility(options CheckVisibilityOptions) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementCheckVisibility(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCheckVisibility1 returns true if the method "Element.checkVisibility" exists.
func (this Element) HasCheckVisibility1() bool {
	return js.True == bindings.HasElementCheckVisibility1(
		this.Ref(),
	)
}

// CheckVisibility1Func returns the method "Element.checkVisibility".
func (this Element) CheckVisibility1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementCheckVisibility1Func(
			this.Ref(),
		),
	)
}

// CheckVisibility1 calls the method "Element.checkVisibility".
func (this Element) CheckVisibility1() (ret bool) {
	bindings.CallElementCheckVisibility1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckVisibility1 calls the method "Element.checkVisibility"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryCheckVisibility1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementCheckVisibility1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollIntoView returns true if the method "Element.scrollIntoView" exists.
func (this Element) HasScrollIntoView() bool {
	return js.True == bindings.HasElementScrollIntoView(
		this.Ref(),
	)
}

// ScrollIntoViewFunc returns the method "Element.scrollIntoView".
func (this Element) ScrollIntoViewFunc() (fn js.Func[func(arg OneOf_Bool_ScrollIntoViewOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollIntoViewFunc(
			this.Ref(),
		),
	)
}

// ScrollIntoView calls the method "Element.scrollIntoView".
func (this Element) ScrollIntoView(arg OneOf_Bool_ScrollIntoViewOptions) (ret js.Void) {
	bindings.CallElementScrollIntoView(
		this.Ref(), js.Pointer(&ret),
		arg.Ref(),
	)

	return
}

// TryScrollIntoView calls the method "Element.scrollIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollIntoView(arg OneOf_Bool_ScrollIntoViewOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollIntoView(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		arg.Ref(),
	)

	return
}

// HasScrollIntoView1 returns true if the method "Element.scrollIntoView" exists.
func (this Element) HasScrollIntoView1() bool {
	return js.True == bindings.HasElementScrollIntoView1(
		this.Ref(),
	)
}

// ScrollIntoView1Func returns the method "Element.scrollIntoView".
func (this Element) ScrollIntoView1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollIntoView1Func(
			this.Ref(),
		),
	)
}

// ScrollIntoView1 calls the method "Element.scrollIntoView".
func (this Element) ScrollIntoView1() (ret js.Void) {
	bindings.CallElementScrollIntoView1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollIntoView1 calls the method "Element.scrollIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollIntoView1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollIntoView1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScroll returns true if the method "Element.scroll" exists.
func (this Element) HasScroll() bool {
	return js.True == bindings.HasElementScroll(
		this.Ref(),
	)
}

// ScrollFunc returns the method "Element.scroll".
func (this Element) ScrollFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollFunc(
			this.Ref(),
		),
	)
}

// Scroll calls the method "Element.scroll".
func (this Element) Scroll(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScroll(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScroll calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScroll(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScroll1 returns true if the method "Element.scroll" exists.
func (this Element) HasScroll1() bool {
	return js.True == bindings.HasElementScroll1(
		this.Ref(),
	)
}

// Scroll1Func returns the method "Element.scroll".
func (this Element) Scroll1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScroll1Func(
			this.Ref(),
		),
	)
}

// Scroll1 calls the method "Element.scroll".
func (this Element) Scroll1() (ret js.Void) {
	bindings.CallElementScroll1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScroll1 calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScroll1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScroll2 returns true if the method "Element.scroll" exists.
func (this Element) HasScroll2() bool {
	return js.True == bindings.HasElementScroll2(
		this.Ref(),
	)
}

// Scroll2Func returns the method "Element.scroll".
func (this Element) Scroll2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScroll2Func(
			this.Ref(),
		),
	)
}

// Scroll2 calls the method "Element.scroll".
func (this Element) Scroll2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScroll2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScroll2 calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScroll2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasScrollTo returns true if the method "Element.scrollTo" exists.
func (this Element) HasScrollTo() bool {
	return js.True == bindings.HasElementScrollTo(
		this.Ref(),
	)
}

// ScrollToFunc returns the method "Element.scrollTo".
func (this Element) ScrollToFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollToFunc(
			this.Ref(),
		),
	)
}

// ScrollTo calls the method "Element.scrollTo".
func (this Element) ScrollTo(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScrollTo(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollTo calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollTo(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScrollTo1 returns true if the method "Element.scrollTo" exists.
func (this Element) HasScrollTo1() bool {
	return js.True == bindings.HasElementScrollTo1(
		this.Ref(),
	)
}

// ScrollTo1Func returns the method "Element.scrollTo".
func (this Element) ScrollTo1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollTo1Func(
			this.Ref(),
		),
	)
}

// ScrollTo1 calls the method "Element.scrollTo".
func (this Element) ScrollTo1() (ret js.Void) {
	bindings.CallElementScrollTo1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollTo1 calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollTo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollTo2 returns true if the method "Element.scrollTo" exists.
func (this Element) HasScrollTo2() bool {
	return js.True == bindings.HasElementScrollTo2(
		this.Ref(),
	)
}

// ScrollTo2Func returns the method "Element.scrollTo".
func (this Element) ScrollTo2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScrollTo2Func(
			this.Ref(),
		),
	)
}

// ScrollTo2 calls the method "Element.scrollTo".
func (this Element) ScrollTo2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScrollTo2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollTo2 calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollTo2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasScrollBy returns true if the method "Element.scrollBy" exists.
func (this Element) HasScrollBy() bool {
	return js.True == bindings.HasElementScrollBy(
		this.Ref(),
	)
}

// ScrollByFunc returns the method "Element.scrollBy".
func (this Element) ScrollByFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.ElementScrollByFunc(
			this.Ref(),
		),
	)
}

// ScrollBy calls the method "Element.scrollBy".
func (this Element) ScrollBy(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScrollBy(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollBy calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollBy(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScrollBy1 returns true if the method "Element.scrollBy" exists.
func (this Element) HasScrollBy1() bool {
	return js.True == bindings.HasElementScrollBy1(
		this.Ref(),
	)
}

// ScrollBy1Func returns the method "Element.scrollBy".
func (this Element) ScrollBy1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementScrollBy1Func(
			this.Ref(),
		),
	)
}

// ScrollBy1 calls the method "Element.scrollBy".
func (this Element) ScrollBy1() (ret js.Void) {
	bindings.CallElementScrollBy1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollBy1 calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollBy1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollBy2 returns true if the method "Element.scrollBy" exists.
func (this Element) HasScrollBy2() bool {
	return js.True == bindings.HasElementScrollBy2(
		this.Ref(),
	)
}

// ScrollBy2Func returns the method "Element.scrollBy".
func (this Element) ScrollBy2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.ElementScrollBy2Func(
			this.Ref(),
		),
	)
}

// ScrollBy2 calls the method "Element.scrollBy".
func (this Element) ScrollBy2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScrollBy2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollBy2 calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryScrollBy2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasGetBoxQuads returns true if the method "Element.getBoxQuads" exists.
func (this Element) HasGetBoxQuads() bool {
	return js.True == bindings.HasElementGetBoxQuads(
		this.Ref(),
	)
}

// GetBoxQuadsFunc returns the method "Element.getBoxQuads".
func (this Element) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.ElementGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Element.getBoxQuads".
func (this Element) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallElementGetBoxQuads(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Element.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoxQuads(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetBoxQuads1 returns true if the method "Element.getBoxQuads" exists.
func (this Element) HasGetBoxQuads1() bool {
	return js.True == bindings.HasElementGetBoxQuads1(
		this.Ref(),
	)
}

// GetBoxQuads1Func returns the method "Element.getBoxQuads".
func (this Element) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.ElementGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Element.getBoxQuads".
func (this Element) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallElementGetBoxQuads1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Element.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoxQuads1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvertQuadFromNode returns true if the method "Element.convertQuadFromNode" exists.
func (this Element) HasConvertQuadFromNode() bool {
	return js.True == bindings.HasElementConvertQuadFromNode(
		this.Ref(),
	)
}

// ConvertQuadFromNodeFunc returns the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "Element.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertQuadFromNode1 returns true if the method "Element.convertQuadFromNode" exists.
func (this Element) HasConvertQuadFromNode1() bool {
	return js.True == bindings.HasElementConvertQuadFromNode1(
		this.Ref(),
	)
}

// ConvertQuadFromNode1Func returns the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "Element.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasConvertRectFromNode returns true if the method "Element.convertRectFromNode" exists.
func (this Element) HasConvertRectFromNode() bool {
	return js.True == bindings.HasElementConvertRectFromNode(
		this.Ref(),
	)
}

// ConvertRectFromNodeFunc returns the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallElementConvertRectFromNode(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "Element.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertRectFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertRectFromNode1 returns true if the method "Element.convertRectFromNode" exists.
func (this Element) HasConvertRectFromNode1() bool {
	return js.True == bindings.HasElementConvertRectFromNode1(
		this.Ref(),
	)
}

// ConvertRectFromNode1Func returns the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.ElementConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "Element.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasConvertPointFromNode returns true if the method "Element.convertPointFromNode" exists.
func (this Element) HasConvertPointFromNode() bool {
	return js.True == bindings.HasElementConvertPointFromNode(
		this.Ref(),
	)
}

// ConvertPointFromNodeFunc returns the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.ElementConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallElementConvertPointFromNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "Element.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertPointFromNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertPointFromNode1 returns true if the method "Element.convertPointFromNode" exists.
func (this Element) HasConvertPointFromNode1() bool {
	return js.True == bindings.HasElementConvertPointFromNode1(
		this.Ref(),
	)
}

// ConvertPointFromNode1Func returns the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.ElementConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "Element.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// HasPrepend returns true if the method "Element.prepend" exists.
func (this Element) HasPrepend() bool {
	return js.True == bindings.HasElementPrepend(
		this.Ref(),
	)
}

// PrependFunc returns the method "Element.prepend".
func (this Element) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementPrependFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "Element.prepend".
func (this Element) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementPrepend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryPrepend calls the method "Element.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryPrepend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementPrepend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAppend returns true if the method "Element.append" exists.
func (this Element) HasAppend() bool {
	return js.True == bindings.HasElementAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "Element.append".
func (this Element) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "Element.append".
func (this Element) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementAppend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAppend calls the method "Element.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryAppend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceChildren returns true if the method "Element.replaceChildren" exists.
func (this Element) HasReplaceChildren() bool {
	return js.True == bindings.HasElementReplaceChildren(
		this.Ref(),
	)
}

// ReplaceChildrenFunc returns the method "Element.replaceChildren".
func (this Element) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "Element.replaceChildren".
func (this Element) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementReplaceChildren(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceChildren calls the method "Element.replaceChildren"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReplaceChildren(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasQuerySelector returns true if the method "Element.querySelector" exists.
func (this Element) HasQuerySelector() bool {
	return js.True == bindings.HasElementQuerySelector(
		this.Ref(),
	)
}

// QuerySelectorFunc returns the method "Element.querySelector".
func (this Element) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.ElementQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "Element.querySelector".
func (this Element) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallElementQuerySelector(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "Element.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementQuerySelector(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasQuerySelectorAll returns true if the method "Element.querySelectorAll" exists.
func (this Element) HasQuerySelectorAll() bool {
	return js.True == bindings.HasElementQuerySelectorAll(
		this.Ref(),
	)
}

// QuerySelectorAllFunc returns the method "Element.querySelectorAll".
func (this Element) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.ElementQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "Element.querySelectorAll".
func (this Element) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallElementQuerySelectorAll(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "Element.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementQuerySelectorAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasBefore returns true if the method "Element.before" exists.
func (this Element) HasBefore() bool {
	return js.True == bindings.HasElementBefore(
		this.Ref(),
	)
}

// BeforeFunc returns the method "Element.before".
func (this Element) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementBeforeFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "Element.before".
func (this Element) Before(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementBefore(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryBefore calls the method "Element.before"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryBefore(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAfter returns true if the method "Element.after" exists.
func (this Element) HasAfter() bool {
	return js.True == bindings.HasElementAfter(
		this.Ref(),
	)
}

// AfterFunc returns the method "Element.after".
func (this Element) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementAfterFunc(
			this.Ref(),
		),
	)
}

// After calls the method "Element.after".
func (this Element) After(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementAfter(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAfter calls the method "Element.after"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryAfter(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceWith returns true if the method "Element.replaceWith" exists.
func (this Element) HasReplaceWith() bool {
	return js.True == bindings.HasElementReplaceWith(
		this.Ref(),
	)
}

// ReplaceWithFunc returns the method "Element.replaceWith".
func (this Element) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.ElementReplaceWithFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "Element.replaceWith".
func (this Element) ReplaceWith(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementReplaceWith(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceWith calls the method "Element.replaceWith"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryReplaceWith(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReplaceWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasRemove returns true if the method "Element.remove" exists.
func (this Element) HasRemove() bool {
	return js.True == bindings.HasElementRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "Element.remove".
func (this Element) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "Element.remove".
func (this Element) Remove() (ret js.Void) {
	bindings.CallElementRemove(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "Element.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAnimate returns true if the method "Element.animate" exists.
func (this Element) HasAnimate() bool {
	return js.True == bindings.HasElementAnimate(
		this.Ref(),
	)
}

// AnimateFunc returns the method "Element.animate".
func (this Element) AnimateFunc() (fn js.Func[func(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) Animation]) {
	return fn.FromRef(
		bindings.ElementAnimateFunc(
			this.Ref(),
		),
	)
}

// Animate calls the method "Element.animate".
func (this Element) Animate(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) (ret Animation) {
	bindings.CallElementAnimate(
		this.Ref(), js.Pointer(&ret),
		keyframes.Ref(),
		options.Ref(),
	)

	return
}

// TryAnimate calls the method "Element.animate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryAnimate(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAnimate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyframes.Ref(),
		options.Ref(),
	)

	return
}

// HasAnimate1 returns true if the method "Element.animate" exists.
func (this Element) HasAnimate1() bool {
	return js.True == bindings.HasElementAnimate1(
		this.Ref(),
	)
}

// Animate1Func returns the method "Element.animate".
func (this Element) Animate1Func() (fn js.Func[func(keyframes js.Object) Animation]) {
	return fn.FromRef(
		bindings.ElementAnimate1Func(
			this.Ref(),
		),
	)
}

// Animate1 calls the method "Element.animate".
func (this Element) Animate1(keyframes js.Object) (ret Animation) {
	bindings.CallElementAnimate1(
		this.Ref(), js.Pointer(&ret),
		keyframes.Ref(),
	)

	return
}

// TryAnimate1 calls the method "Element.animate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryAnimate1(keyframes js.Object) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAnimate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyframes.Ref(),
	)

	return
}

// HasGetAnimations returns true if the method "Element.getAnimations" exists.
func (this Element) HasGetAnimations() bool {
	return js.True == bindings.HasElementGetAnimations(
		this.Ref(),
	)
}

// GetAnimationsFunc returns the method "Element.getAnimations".
func (this Element) GetAnimationsFunc() (fn js.Func[func(options GetAnimationsOptions) js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ElementGetAnimationsFunc(
			this.Ref(),
		),
	)
}

// GetAnimations calls the method "Element.getAnimations".
func (this Element) GetAnimations(options GetAnimationsOptions) (ret js.Array[Animation]) {
	bindings.CallElementGetAnimations(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetAnimations calls the method "Element.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAnimations(options GetAnimationsOptions) (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAnimations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetAnimations1 returns true if the method "Element.getAnimations" exists.
func (this Element) HasGetAnimations1() bool {
	return js.True == bindings.HasElementGetAnimations1(
		this.Ref(),
	)
}

// GetAnimations1Func returns the method "Element.getAnimations".
func (this Element) GetAnimations1Func() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ElementGetAnimations1Func(
			this.Ref(),
		),
	)
}

// GetAnimations1 calls the method "Element.getAnimations".
func (this Element) GetAnimations1() (ret js.Array[Animation]) {
	bindings.CallElementGetAnimations1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAnimations1 calls the method "Element.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetAnimations1() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAnimations1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetRegionFlowRanges returns true if the method "Element.getRegionFlowRanges" exists.
func (this Element) HasGetRegionFlowRanges() bool {
	return js.True == bindings.HasElementGetRegionFlowRanges(
		this.Ref(),
	)
}

// GetRegionFlowRangesFunc returns the method "Element.getRegionFlowRanges".
func (this Element) GetRegionFlowRangesFunc() (fn js.Func[func() js.Array[Range]]) {
	return fn.FromRef(
		bindings.ElementGetRegionFlowRangesFunc(
			this.Ref(),
		),
	)
}

// GetRegionFlowRanges calls the method "Element.getRegionFlowRanges".
func (this Element) GetRegionFlowRanges() (ret js.Array[Range]) {
	bindings.CallElementGetRegionFlowRanges(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetRegionFlowRanges calls the method "Element.getRegionFlowRanges"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Element) TryGetRegionFlowRanges() (ret js.Array[Range], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetRegionFlowRanges(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this HTMLCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCollectionLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "HTMLCollection.item" exists.
func (this HTMLCollection) HasItem() bool {
	return js.True == bindings.HasHTMLCollectionItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "HTMLCollection.item".
func (this HTMLCollection) ItemFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLCollectionItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "HTMLCollection.item".
func (this HTMLCollection) Item(index uint32) (ret Element) {
	bindings.CallHTMLCollectionItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "HTMLCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCollection) TryItem(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCollectionItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "HTMLCollection.namedItem" exists.
func (this HTMLCollection) HasNamedItem() bool {
	return js.True == bindings.HasHTMLCollectionNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "HTMLCollection.namedItem".
func (this HTMLCollection) NamedItemFunc() (fn js.Func[func(name js.String) Element]) {
	return fn.FromRef(
		bindings.HTMLCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLCollection.namedItem".
func (this HTMLCollection) NamedItem(name js.String) (ret Element) {
	bindings.CallHTMLCollectionNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCollection) TryNamedItem(name js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCollectionNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
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

func NewCDATASection(data js.String) (ret CDATASection) {
	ret.ref = bindings.NewCDATASectionByCDATASection(
		data.Ref())
	return
}

func NewCDATASectionByCDATASection1() (ret CDATASection) {
	ret.ref = bindings.NewCDATASectionByCDATASection1()
	return
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

func NewComment(data js.String) (ret Comment) {
	ret.ref = bindings.NewCommentByComment(
		data.Ref())
	return
}

func NewCommentByComment1() (ret Comment) {
	ret.ref = bindings.NewCommentByComment1()
	return
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
// It returns ok=false if there is no such property.
func (this ProcessingInstruction) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetProcessingInstructionTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sheet returns the value of property "ProcessingInstruction.sheet".
//
// It returns ok=false if there is no such property.
func (this ProcessingInstruction) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetProcessingInstructionSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this NodeIterator) Root() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeIteratorRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReferenceNode returns the value of property "NodeIterator.referenceNode".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) ReferenceNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeIteratorReferenceNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointerBeforeReferenceNode returns the value of property "NodeIterator.pointerBeforeReferenceNode".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) PointerBeforeReferenceNode() (ret bool, ok bool) {
	ok = js.True == bindings.GetNodeIteratorPointerBeforeReferenceNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WhatToShow returns the value of property "NodeIterator.whatToShow".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) WhatToShow() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNodeIteratorWhatToShow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Filter returns the value of property "NodeIterator.filter".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) Filter() (ret js.Func[func(node Node) uint16], ok bool) {
	ok = js.True == bindings.GetNodeIteratorFilter(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasNextNode returns true if the method "NodeIterator.nextNode" exists.
func (this NodeIterator) HasNextNode() bool {
	return js.True == bindings.HasNodeIteratorNextNode(
		this.Ref(),
	)
}

// NextNodeFunc returns the method "NodeIterator.nextNode".
func (this NodeIterator) NextNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeIteratorNextNodeFunc(
			this.Ref(),
		),
	)
}

// NextNode calls the method "NodeIterator.nextNode".
func (this NodeIterator) NextNode() (ret Node) {
	bindings.CallNodeIteratorNextNode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryNextNode calls the method "NodeIterator.nextNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NodeIterator) TryNextNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorNextNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPreviousNode returns true if the method "NodeIterator.previousNode" exists.
func (this NodeIterator) HasPreviousNode() bool {
	return js.True == bindings.HasNodeIteratorPreviousNode(
		this.Ref(),
	)
}

// PreviousNodeFunc returns the method "NodeIterator.previousNode".
func (this NodeIterator) PreviousNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeIteratorPreviousNodeFunc(
			this.Ref(),
		),
	)
}

// PreviousNode calls the method "NodeIterator.previousNode".
func (this NodeIterator) PreviousNode() (ret Node) {
	bindings.CallNodeIteratorPreviousNode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPreviousNode calls the method "NodeIterator.previousNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NodeIterator) TryPreviousNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorPreviousNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDetach returns true if the method "NodeIterator.detach" exists.
func (this NodeIterator) HasDetach() bool {
	return js.True == bindings.HasNodeIteratorDetach(
		this.Ref(),
	)
}

// DetachFunc returns the method "NodeIterator.detach".
func (this NodeIterator) DetachFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NodeIteratorDetachFunc(
			this.Ref(),
		),
	)
}

// Detach calls the method "NodeIterator.detach".
func (this NodeIterator) Detach() (ret js.Void) {
	bindings.CallNodeIteratorDetach(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDetach calls the method "NodeIterator.detach"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NodeIterator) TryDetach() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorDetach(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this TreeWalker) Root() (ret Node, ok bool) {
	ok = js.True == bindings.GetTreeWalkerRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WhatToShow returns the value of property "TreeWalker.whatToShow".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) WhatToShow() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTreeWalkerWhatToShow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Filter returns the value of property "TreeWalker.filter".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) Filter() (ret js.Func[func(node Node) uint16], ok bool) {
	ok = js.True == bindings.GetTreeWalkerFilter(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentNode returns the value of property "TreeWalker.currentNode".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) CurrentNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetTreeWalkerCurrentNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCurrentNode sets the value of property "TreeWalker.currentNode" to val.
//
// It returns false if the property cannot be set.
func (this TreeWalker) SetCurrentNode(val Node) bool {
	return js.True == bindings.SetTreeWalkerCurrentNode(
		this.Ref(),
		val.Ref(),
	)
}

// HasParentNode returns true if the method "TreeWalker.parentNode" exists.
func (this TreeWalker) HasParentNode() bool {
	return js.True == bindings.HasTreeWalkerParentNode(
		this.Ref(),
	)
}

// ParentNodeFunc returns the method "TreeWalker.parentNode".
func (this TreeWalker) ParentNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerParentNodeFunc(
			this.Ref(),
		),
	)
}

// ParentNode calls the method "TreeWalker.parentNode".
func (this TreeWalker) ParentNode() (ret Node) {
	bindings.CallTreeWalkerParentNode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryParentNode calls the method "TreeWalker.parentNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryParentNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerParentNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFirstChild returns true if the method "TreeWalker.firstChild" exists.
func (this TreeWalker) HasFirstChild() bool {
	return js.True == bindings.HasTreeWalkerFirstChild(
		this.Ref(),
	)
}

// FirstChildFunc returns the method "TreeWalker.firstChild".
func (this TreeWalker) FirstChildFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerFirstChildFunc(
			this.Ref(),
		),
	)
}

// FirstChild calls the method "TreeWalker.firstChild".
func (this TreeWalker) FirstChild() (ret Node) {
	bindings.CallTreeWalkerFirstChild(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFirstChild calls the method "TreeWalker.firstChild"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryFirstChild() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerFirstChild(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasLastChild returns true if the method "TreeWalker.lastChild" exists.
func (this TreeWalker) HasLastChild() bool {
	return js.True == bindings.HasTreeWalkerLastChild(
		this.Ref(),
	)
}

// LastChildFunc returns the method "TreeWalker.lastChild".
func (this TreeWalker) LastChildFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerLastChildFunc(
			this.Ref(),
		),
	)
}

// LastChild calls the method "TreeWalker.lastChild".
func (this TreeWalker) LastChild() (ret Node) {
	bindings.CallTreeWalkerLastChild(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLastChild calls the method "TreeWalker.lastChild"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryLastChild() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerLastChild(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPreviousSibling returns true if the method "TreeWalker.previousSibling" exists.
func (this TreeWalker) HasPreviousSibling() bool {
	return js.True == bindings.HasTreeWalkerPreviousSibling(
		this.Ref(),
	)
}

// PreviousSiblingFunc returns the method "TreeWalker.previousSibling".
func (this TreeWalker) PreviousSiblingFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerPreviousSiblingFunc(
			this.Ref(),
		),
	)
}

// PreviousSibling calls the method "TreeWalker.previousSibling".
func (this TreeWalker) PreviousSibling() (ret Node) {
	bindings.CallTreeWalkerPreviousSibling(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPreviousSibling calls the method "TreeWalker.previousSibling"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryPreviousSibling() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerPreviousSibling(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasNextSibling returns true if the method "TreeWalker.nextSibling" exists.
func (this TreeWalker) HasNextSibling() bool {
	return js.True == bindings.HasTreeWalkerNextSibling(
		this.Ref(),
	)
}

// NextSiblingFunc returns the method "TreeWalker.nextSibling".
func (this TreeWalker) NextSiblingFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerNextSiblingFunc(
			this.Ref(),
		),
	)
}

// NextSibling calls the method "TreeWalker.nextSibling".
func (this TreeWalker) NextSibling() (ret Node) {
	bindings.CallTreeWalkerNextSibling(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryNextSibling calls the method "TreeWalker.nextSibling"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryNextSibling() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerNextSibling(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPreviousNode returns true if the method "TreeWalker.previousNode" exists.
func (this TreeWalker) HasPreviousNode() bool {
	return js.True == bindings.HasTreeWalkerPreviousNode(
		this.Ref(),
	)
}

// PreviousNodeFunc returns the method "TreeWalker.previousNode".
func (this TreeWalker) PreviousNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerPreviousNodeFunc(
			this.Ref(),
		),
	)
}

// PreviousNode calls the method "TreeWalker.previousNode".
func (this TreeWalker) PreviousNode() (ret Node) {
	bindings.CallTreeWalkerPreviousNode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPreviousNode calls the method "TreeWalker.previousNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryPreviousNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerPreviousNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasNextNode returns true if the method "TreeWalker.nextNode" exists.
func (this TreeWalker) HasNextNode() bool {
	return js.True == bindings.HasTreeWalkerNextNode(
		this.Ref(),
	)
}

// NextNodeFunc returns the method "TreeWalker.nextNode".
func (this TreeWalker) NextNodeFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.TreeWalkerNextNodeFunc(
			this.Ref(),
		),
	)
}

// NextNode calls the method "TreeWalker.nextNode".
func (this TreeWalker) NextNode() (ret Node) {
	bindings.CallTreeWalkerNextNode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryNextNode calls the method "TreeWalker.nextNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TreeWalker) TryNextNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerNextNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this ViewTransition) UpdateCallbackDone() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionUpdateCallbackDone(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "ViewTransition.ready".
//
// It returns ok=false if there is no such property.
func (this ViewTransition) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "ViewTransition.finished".
//
// It returns ok=false if there is no such property.
func (this ViewTransition) Finished() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionFinished(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSkipTransition returns true if the method "ViewTransition.skipTransition" exists.
func (this ViewTransition) HasSkipTransition() bool {
	return js.True == bindings.HasViewTransitionSkipTransition(
		this.Ref(),
	)
}

// SkipTransitionFunc returns the method "ViewTransition.skipTransition".
func (this ViewTransition) SkipTransitionFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ViewTransitionSkipTransitionFunc(
			this.Ref(),
		),
	)
}

// SkipTransition calls the method "ViewTransition.skipTransition".
func (this ViewTransition) SkipTransition() (ret js.Void) {
	bindings.CallViewTransitionSkipTransition(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySkipTransition calls the method "ViewTransition.skipTransition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ViewTransition) TrySkipTransition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryViewTransitionSkipTransition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

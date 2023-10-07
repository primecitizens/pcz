// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type HTMLSlotElement struct {
	HTMLElement
}

func (this HTMLSlotElement) Once() HTMLSlotElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "HTMLSlotElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLSlotElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSlotElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLSlotElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSlotElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLSlotElementName(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAssignedNodes returns true if the method "HTMLSlotElement.assignedNodes" exists.
func (this HTMLSlotElement) HasFuncAssignedNodes() bool {
	return js.True == bindings.HasFuncHTMLSlotElementAssignedNodes(
		this.ref,
	)
}

// FuncAssignedNodes returns the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) FuncAssignedNodes() (fn js.Func[func(options AssignedNodesOptions) js.Array[Node]]) {
	bindings.FuncHTMLSlotElementAssignedNodes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AssignedNodes calls the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodes(options AssignedNodesOptions) (ret js.Array[Node]) {
	bindings.CallHTMLSlotElementAssignedNodes(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAssignedNodes calls the method "HTMLSlotElement.assignedNodes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSlotElement) TryAssignedNodes(options AssignedNodesOptions) (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedNodes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncAssignedNodes1 returns true if the method "HTMLSlotElement.assignedNodes" exists.
func (this HTMLSlotElement) HasFuncAssignedNodes1() bool {
	return js.True == bindings.HasFuncHTMLSlotElementAssignedNodes1(
		this.ref,
	)
}

// FuncAssignedNodes1 returns the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) FuncAssignedNodes1() (fn js.Func[func() js.Array[Node]]) {
	bindings.FuncHTMLSlotElementAssignedNodes1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AssignedNodes1 calls the method "HTMLSlotElement.assignedNodes".
func (this HTMLSlotElement) AssignedNodes1() (ret js.Array[Node]) {
	bindings.CallHTMLSlotElementAssignedNodes1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAssignedNodes1 calls the method "HTMLSlotElement.assignedNodes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSlotElement) TryAssignedNodes1() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedNodes1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAssignedElements returns true if the method "HTMLSlotElement.assignedElements" exists.
func (this HTMLSlotElement) HasFuncAssignedElements() bool {
	return js.True == bindings.HasFuncHTMLSlotElementAssignedElements(
		this.ref,
	)
}

// FuncAssignedElements returns the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) FuncAssignedElements() (fn js.Func[func(options AssignedNodesOptions) js.Array[Element]]) {
	bindings.FuncHTMLSlotElementAssignedElements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AssignedElements calls the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElements(options AssignedNodesOptions) (ret js.Array[Element]) {
	bindings.CallHTMLSlotElementAssignedElements(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAssignedElements calls the method "HTMLSlotElement.assignedElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSlotElement) TryAssignedElements(options AssignedNodesOptions) (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedElements(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncAssignedElements1 returns true if the method "HTMLSlotElement.assignedElements" exists.
func (this HTMLSlotElement) HasFuncAssignedElements1() bool {
	return js.True == bindings.HasFuncHTMLSlotElementAssignedElements1(
		this.ref,
	)
}

// FuncAssignedElements1 returns the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) FuncAssignedElements1() (fn js.Func[func() js.Array[Element]]) {
	bindings.FuncHTMLSlotElementAssignedElements1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AssignedElements1 calls the method "HTMLSlotElement.assignedElements".
func (this HTMLSlotElement) AssignedElements1() (ret js.Array[Element]) {
	bindings.CallHTMLSlotElementAssignedElements1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAssignedElements1 calls the method "HTMLSlotElement.assignedElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSlotElement) TryAssignedElements1() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssignedElements1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAssign returns true if the method "HTMLSlotElement.assign" exists.
func (this HTMLSlotElement) HasFuncAssign() bool {
	return js.True == bindings.HasFuncHTMLSlotElementAssign(
		this.ref,
	)
}

// FuncAssign returns the method "HTMLSlotElement.assign".
func (this HTMLSlotElement) FuncAssign() (fn js.Func[func(nodes ...OneOf_Element_Text)]) {
	bindings.FuncHTMLSlotElementAssign(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Assign calls the method "HTMLSlotElement.assign".
func (this HTMLSlotElement) Assign(nodes ...OneOf_Element_Text) (ret js.Void) {
	bindings.CallHTMLSlotElementAssign(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAssign calls the method "HTMLSlotElement.assign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSlotElement) TryAssign(nodes ...OneOf_Element_Text) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSlotElementAssign(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// WholeText returns the value of property "Text.wholeText".
//
// It returns ok=false if there is no such property.
func (this Text) WholeText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextWholeText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AssignedSlot returns the value of property "Text.assignedSlot".
//
// It returns ok=false if there is no such property.
func (this Text) AssignedSlot() (ret HTMLSlotElement, ok bool) {
	ok = js.True == bindings.GetTextAssignedSlot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSplitText returns true if the method "Text.splitText" exists.
func (this Text) HasFuncSplitText() bool {
	return js.True == bindings.HasFuncTextSplitText(
		this.ref,
	)
}

// FuncSplitText returns the method "Text.splitText".
func (this Text) FuncSplitText() (fn js.Func[func(offset uint32) Text]) {
	bindings.FuncTextSplitText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SplitText calls the method "Text.splitText".
func (this Text) SplitText(offset uint32) (ret Text) {
	bindings.CallTextSplitText(
		this.ref, js.Pointer(&ret),
		uint32(offset),
	)

	return
}

// TrySplitText calls the method "Text.splitText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TrySplitText(offset uint32) (ret Text, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextSplitText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
	)

	return
}

// HasFuncGetBoxQuads returns true if the method "Text.getBoxQuads" exists.
func (this Text) HasFuncGetBoxQuads() bool {
	return js.True == bindings.HasFuncTextGetBoxQuads(
		this.ref,
	)
}

// FuncGetBoxQuads returns the method "Text.getBoxQuads".
func (this Text) FuncGetBoxQuads() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	bindings.FuncTextGetBoxQuads(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads calls the method "Text.getBoxQuads".
func (this Text) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallTextGetBoxQuads(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Text.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextGetBoxQuads(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetBoxQuads1 returns true if the method "Text.getBoxQuads" exists.
func (this Text) HasFuncGetBoxQuads1() bool {
	return js.True == bindings.HasFuncTextGetBoxQuads1(
		this.ref,
	)
}

// FuncGetBoxQuads1 returns the method "Text.getBoxQuads".
func (this Text) FuncGetBoxQuads1() (fn js.Func[func() js.Array[DOMQuad]]) {
	bindings.FuncTextGetBoxQuads1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads1 calls the method "Text.getBoxQuads".
func (this Text) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallTextGetBoxQuads1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Text.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextGetBoxQuads1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvertQuadFromNode returns true if the method "Text.convertQuadFromNode" exists.
func (this Text) HasFuncConvertQuadFromNode() bool {
	return js.True == bindings.HasFuncTextConvertQuadFromNode(
		this.ref,
	)
}

// FuncConvertQuadFromNode returns the method "Text.convertQuadFromNode".
func (this Text) FuncConvertQuadFromNode() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncTextConvertQuadFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode calls the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallTextConvertQuadFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "Text.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertQuadFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertQuadFromNode1 returns true if the method "Text.convertQuadFromNode" exists.
func (this Text) HasFuncConvertQuadFromNode1() bool {
	return js.True == bindings.HasFuncTextConvertQuadFromNode1(
		this.ref,
	)
}

// FuncConvertQuadFromNode1 returns the method "Text.convertQuadFromNode".
func (this Text) FuncConvertQuadFromNode1() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	bindings.FuncTextConvertQuadFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode1 calls the method "Text.convertQuadFromNode".
func (this Text) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallTextConvertQuadFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "Text.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertQuadFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasFuncConvertRectFromNode returns true if the method "Text.convertRectFromNode" exists.
func (this Text) HasFuncConvertRectFromNode() bool {
	return js.True == bindings.HasFuncTextConvertRectFromNode(
		this.ref,
	)
}

// FuncConvertRectFromNode returns the method "Text.convertRectFromNode".
func (this Text) FuncConvertRectFromNode() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncTextConvertRectFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode calls the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallTextConvertRectFromNode(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "Text.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertRectFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertRectFromNode1 returns true if the method "Text.convertRectFromNode" exists.
func (this Text) HasFuncConvertRectFromNode1() bool {
	return js.True == bindings.HasFuncTextConvertRectFromNode1(
		this.ref,
	)
}

// FuncConvertRectFromNode1 returns the method "Text.convertRectFromNode".
func (this Text) FuncConvertRectFromNode1() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	bindings.FuncTextConvertRectFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode1 calls the method "Text.convertRectFromNode".
func (this Text) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallTextConvertRectFromNode1(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "Text.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertRectFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasFuncConvertPointFromNode returns true if the method "Text.convertPointFromNode" exists.
func (this Text) HasFuncConvertPointFromNode() bool {
	return js.True == bindings.HasFuncTextConvertPointFromNode(
		this.ref,
	)
}

// FuncConvertPointFromNode returns the method "Text.convertPointFromNode".
func (this Text) FuncConvertPointFromNode() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	bindings.FuncTextConvertPointFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode calls the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallTextConvertPointFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "Text.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertPointFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertPointFromNode1 returns true if the method "Text.convertPointFromNode" exists.
func (this Text) HasFuncConvertPointFromNode1() bool {
	return js.True == bindings.HasFuncTextConvertPointFromNode1(
		this.ref,
	)
}

// FuncConvertPointFromNode1 returns the method "Text.convertPointFromNode".
func (this Text) FuncConvertPointFromNode1() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	bindings.FuncTextConvertPointFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode1 calls the method "Text.convertPointFromNode".
func (this Text) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallTextConvertPointFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "Text.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Text) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextConvertPointFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BoxQuadOptions) UpdateFrom(ref js.Ref) {
	bindings.BoxQuadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BoxQuadOptions) Update(ref js.Ref) {
	bindings.BoxQuadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BoxQuadOptions) FreeMembers(recursive bool) {
	js.Free(
		p.RelativeTo.Ref(),
	)
	p.RelativeTo = p.RelativeTo.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "CSSPseudoElement.type".
//
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Element returns the value of property "CSSPseudoElement.element".
//
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Element() (ret Element, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Parent returns the value of property "CSSPseudoElement.parent".
//
// It returns ok=false if there is no such property.
func (this CSSPseudoElement) Parent() (ret OneOf_Element_CSSPseudoElement, ok bool) {
	ok = js.True == bindings.GetCSSPseudoElementParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPseudo returns true if the method "CSSPseudoElement.pseudo" exists.
func (this CSSPseudoElement) HasFuncPseudo() bool {
	return js.True == bindings.HasFuncCSSPseudoElementPseudo(
		this.ref,
	)
}

// FuncPseudo returns the method "CSSPseudoElement.pseudo".
func (this CSSPseudoElement) FuncPseudo() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	bindings.FuncCSSPseudoElementPseudo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pseudo calls the method "CSSPseudoElement.pseudo".
func (this CSSPseudoElement) Pseudo(typ js.String) (ret CSSPseudoElement) {
	bindings.CallCSSPseudoElementPseudo(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryPseudo calls the method "CSSPseudoElement.pseudo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryPseudo(typ js.String) (ret CSSPseudoElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementPseudo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncGetBoxQuads returns true if the method "CSSPseudoElement.getBoxQuads" exists.
func (this CSSPseudoElement) HasFuncGetBoxQuads() bool {
	return js.True == bindings.HasFuncCSSPseudoElementGetBoxQuads(
		this.ref,
	)
}

// FuncGetBoxQuads returns the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) FuncGetBoxQuads() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	bindings.FuncCSSPseudoElementGetBoxQuads(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads calls the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallCSSPseudoElementGetBoxQuads(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "CSSPseudoElement.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementGetBoxQuads(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetBoxQuads1 returns true if the method "CSSPseudoElement.getBoxQuads" exists.
func (this CSSPseudoElement) HasFuncGetBoxQuads1() bool {
	return js.True == bindings.HasFuncCSSPseudoElementGetBoxQuads1(
		this.ref,
	)
}

// FuncGetBoxQuads1 returns the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) FuncGetBoxQuads1() (fn js.Func[func() js.Array[DOMQuad]]) {
	bindings.FuncCSSPseudoElementGetBoxQuads1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads1 calls the method "CSSPseudoElement.getBoxQuads".
func (this CSSPseudoElement) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallCSSPseudoElementGetBoxQuads1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "CSSPseudoElement.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementGetBoxQuads1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvertQuadFromNode returns true if the method "CSSPseudoElement.convertQuadFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertQuadFromNode() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertQuadFromNode(
		this.ref,
	)
}

// FuncConvertQuadFromNode returns the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) FuncConvertQuadFromNode() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncCSSPseudoElementConvertQuadFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode calls the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertQuadFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "CSSPseudoElement.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertQuadFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertQuadFromNode1 returns true if the method "CSSPseudoElement.convertQuadFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertQuadFromNode1() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertQuadFromNode1(
		this.ref,
	)
}

// FuncConvertQuadFromNode1 returns the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) FuncConvertQuadFromNode1() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	bindings.FuncCSSPseudoElementConvertQuadFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode1 calls the method "CSSPseudoElement.convertQuadFromNode".
func (this CSSPseudoElement) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertQuadFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "CSSPseudoElement.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertQuadFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasFuncConvertRectFromNode returns true if the method "CSSPseudoElement.convertRectFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertRectFromNode() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertRectFromNode(
		this.ref,
	)
}

// FuncConvertRectFromNode returns the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) FuncConvertRectFromNode() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncCSSPseudoElementConvertRectFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode calls the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertRectFromNode(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "CSSPseudoElement.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertRectFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertRectFromNode1 returns true if the method "CSSPseudoElement.convertRectFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertRectFromNode1() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertRectFromNode1(
		this.ref,
	)
}

// FuncConvertRectFromNode1 returns the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) FuncConvertRectFromNode1() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	bindings.FuncCSSPseudoElementConvertRectFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode1 calls the method "CSSPseudoElement.convertRectFromNode".
func (this CSSPseudoElement) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallCSSPseudoElementConvertRectFromNode1(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "CSSPseudoElement.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertRectFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasFuncConvertPointFromNode returns true if the method "CSSPseudoElement.convertPointFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertPointFromNode() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertPointFromNode(
		this.ref,
	)
}

// FuncConvertPointFromNode returns the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) FuncConvertPointFromNode() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	bindings.FuncCSSPseudoElementConvertPointFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode calls the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallCSSPseudoElementConvertPointFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "CSSPseudoElement.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertPointFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertPointFromNode1 returns true if the method "CSSPseudoElement.convertPointFromNode" exists.
func (this CSSPseudoElement) HasFuncConvertPointFromNode1() bool {
	return js.True == bindings.HasFuncCSSPseudoElementConvertPointFromNode1(
		this.ref,
	)
}

// FuncConvertPointFromNode1 returns the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) FuncConvertPointFromNode1() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	bindings.FuncCSSPseudoElementConvertPointFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode1 calls the method "CSSPseudoElement.convertPointFromNode".
func (this CSSPseudoElement) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallCSSPseudoElementConvertPointFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "CSSPseudoElement.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSPseudoElement) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSPseudoElementConvertPointFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *SanitizerConfig) UpdateFrom(ref js.Ref) {
	bindings.SanitizerConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SanitizerConfig) Update(ref js.Ref) {
	bindings.SanitizerConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SanitizerConfig) FreeMembers(recursive bool) {
	js.Free(
		p.AllowElements.Ref(),
		p.BlockElements.Ref(),
		p.DropElements.Ref(),
		p.AllowAttributes.Ref(),
		p.DropAttributes.Ref(),
	)
	p.AllowElements = p.AllowElements.FromRef(js.Undefined)
	p.BlockElements = p.BlockElements.FromRef(js.Undefined)
	p.DropElements = p.DropElements.FromRef(js.Undefined)
	p.AllowAttributes = p.AllowAttributes.FromRef(js.Undefined)
	p.DropAttributes = p.DropAttributes.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "NodeList.length".
//
// It returns ok=false if there is no such property.
func (this NodeList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNodeListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "NodeList.item" exists.
func (this NodeList) HasFuncItem() bool {
	return js.True == bindings.HasFuncNodeListItem(
		this.ref,
	)
}

// FuncItem returns the method "NodeList.item".
func (this NodeList) FuncItem() (fn js.Func[func(index uint32) Node]) {
	bindings.FuncNodeListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "NodeList.item".
func (this NodeList) Item(index uint32) (ret Node) {
	bindings.CallNodeListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "NodeList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NodeList) TryItem(index uint32) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type DocumentFragment struct {
	Node
}

func (this DocumentFragment) Once() DocumentFragment {
	this.ref.Once()
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
	this.ref.Free()
}

// Children returns the value of property "DocumentFragment.children".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentChildren(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "DocumentFragment.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentFirstElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "DocumentFragment.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentLastElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "DocumentFragment.childElementCount".
//
// It returns ok=false if there is no such property.
func (this DocumentFragment) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentChildElementCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetElementById returns true if the method "DocumentFragment.getElementById" exists.
func (this DocumentFragment) HasFuncGetElementById() bool {
	return js.True == bindings.HasFuncDocumentFragmentGetElementById(
		this.ref,
	)
}

// FuncGetElementById returns the method "DocumentFragment.getElementById".
func (this DocumentFragment) FuncGetElementById() (fn js.Func[func(elementId js.String) Element]) {
	bindings.FuncDocumentFragmentGetElementById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementById calls the method "DocumentFragment.getElementById".
func (this DocumentFragment) GetElementById(elementId js.String) (ret Element) {
	bindings.CallDocumentFragmentGetElementById(
		this.ref, js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "DocumentFragment.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentGetElementById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasFuncPrepend returns true if the method "DocumentFragment.prepend" exists.
func (this DocumentFragment) HasFuncPrepend() bool {
	return js.True == bindings.HasFuncDocumentFragmentPrepend(
		this.ref,
	)
}

// FuncPrepend returns the method "DocumentFragment.prepend".
func (this DocumentFragment) FuncPrepend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentFragmentPrepend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prepend calls the method "DocumentFragment.prepend".
func (this DocumentFragment) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentPrepend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryPrepend calls the method "DocumentFragment.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryPrepend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentPrepend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncAppend returns true if the method "DocumentFragment.append" exists.
func (this DocumentFragment) HasFuncAppend() bool {
	return js.True == bindings.HasFuncDocumentFragmentAppend(
		this.ref,
	)
}

// FuncAppend returns the method "DocumentFragment.append".
func (this DocumentFragment) FuncAppend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentFragmentAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "DocumentFragment.append".
func (this DocumentFragment) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentAppend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAppend calls the method "DocumentFragment.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryAppend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncReplaceChildren returns true if the method "DocumentFragment.replaceChildren" exists.
func (this DocumentFragment) HasFuncReplaceChildren() bool {
	return js.True == bindings.HasFuncDocumentFragmentReplaceChildren(
		this.ref,
	)
}

// FuncReplaceChildren returns the method "DocumentFragment.replaceChildren".
func (this DocumentFragment) FuncReplaceChildren() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentFragmentReplaceChildren(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceChildren calls the method "DocumentFragment.replaceChildren".
func (this DocumentFragment) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentFragmentReplaceChildren(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceChildren calls the method "DocumentFragment.replaceChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentReplaceChildren(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncQuerySelector returns true if the method "DocumentFragment.querySelector" exists.
func (this DocumentFragment) HasFuncQuerySelector() bool {
	return js.True == bindings.HasFuncDocumentFragmentQuerySelector(
		this.ref,
	)
}

// FuncQuerySelector returns the method "DocumentFragment.querySelector".
func (this DocumentFragment) FuncQuerySelector() (fn js.Func[func(selectors js.String) Element]) {
	bindings.FuncDocumentFragmentQuerySelector(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelector calls the method "DocumentFragment.querySelector".
func (this DocumentFragment) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallDocumentFragmentQuerySelector(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "DocumentFragment.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentQuerySelector(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncQuerySelectorAll returns true if the method "DocumentFragment.querySelectorAll" exists.
func (this DocumentFragment) HasFuncQuerySelectorAll() bool {
	return js.True == bindings.HasFuncDocumentFragmentQuerySelectorAll(
		this.ref,
	)
}

// FuncQuerySelectorAll returns the method "DocumentFragment.querySelectorAll".
func (this DocumentFragment) FuncQuerySelectorAll() (fn js.Func[func(selectors js.String) NodeList]) {
	bindings.FuncDocumentFragmentQuerySelectorAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelectorAll calls the method "DocumentFragment.querySelectorAll".
func (this DocumentFragment) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallDocumentFragmentQuerySelectorAll(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "DocumentFragment.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentFragment) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentFragmentQuerySelectorAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSanitize returns true if the method "Sanitizer.sanitize" exists.
func (this Sanitizer) HasFuncSanitize() bool {
	return js.True == bindings.HasFuncSanitizerSanitize(
		this.ref,
	)
}

// FuncSanitize returns the method "Sanitizer.sanitize".
func (this Sanitizer) FuncSanitize() (fn js.Func[func(input OneOf_Document_DocumentFragment) DocumentFragment]) {
	bindings.FuncSanitizerSanitize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sanitize calls the method "Sanitizer.sanitize".
func (this Sanitizer) Sanitize(input OneOf_Document_DocumentFragment) (ret DocumentFragment) {
	bindings.CallSanitizerSanitize(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySanitize calls the method "Sanitizer.sanitize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sanitizer) TrySanitize(input OneOf_Document_DocumentFragment) (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerSanitize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSanitizeFor returns true if the method "Sanitizer.sanitizeFor" exists.
func (this Sanitizer) HasFuncSanitizeFor() bool {
	return js.True == bindings.HasFuncSanitizerSanitizeFor(
		this.ref,
	)
}

// FuncSanitizeFor returns the method "Sanitizer.sanitizeFor".
func (this Sanitizer) FuncSanitizeFor() (fn js.Func[func(element js.String, input js.String) Element]) {
	bindings.FuncSanitizerSanitizeFor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SanitizeFor calls the method "Sanitizer.sanitizeFor".
func (this Sanitizer) SanitizeFor(element js.String, input js.String) (ret Element) {
	bindings.CallSanitizerSanitizeFor(
		this.ref, js.Pointer(&ret),
		element.Ref(),
		input.Ref(),
	)

	return
}

// TrySanitizeFor calls the method "Sanitizer.sanitizeFor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sanitizer) TrySanitizeFor(element js.String, input js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerSanitizeFor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		input.Ref(),
	)

	return
}

// HasFuncGetConfiguration returns true if the method "Sanitizer.getConfiguration" exists.
func (this Sanitizer) HasFuncGetConfiguration() bool {
	return js.True == bindings.HasFuncSanitizerGetConfiguration(
		this.ref,
	)
}

// FuncGetConfiguration returns the method "Sanitizer.getConfiguration".
func (this Sanitizer) FuncGetConfiguration() (fn js.Func[func() SanitizerConfig]) {
	bindings.FuncSanitizerGetConfiguration(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetConfiguration calls the method "Sanitizer.getConfiguration".
func (this Sanitizer) GetConfiguration() (ret SanitizerConfig) {
	bindings.CallSanitizerGetConfiguration(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetConfiguration calls the method "Sanitizer.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sanitizer) TryGetConfiguration() (ret SanitizerConfig, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerGetConfiguration(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDefaultConfiguration returns true if the static method "Sanitizer.getDefaultConfiguration" exists.
func (this Sanitizer) HasFuncGetDefaultConfiguration() bool {
	return js.True == bindings.HasFuncSanitizerGetDefaultConfiguration(
		this.ref,
	)
}

// FuncGetDefaultConfiguration returns the static method "Sanitizer.getDefaultConfiguration".
func (this Sanitizer) FuncGetDefaultConfiguration() (fn js.Func[func() SanitizerConfig]) {
	bindings.FuncSanitizerGetDefaultConfiguration(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDefaultConfiguration calls the static method "Sanitizer.getDefaultConfiguration".
func (this Sanitizer) GetDefaultConfiguration() (ret SanitizerConfig) {
	bindings.CallSanitizerGetDefaultConfiguration(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDefaultConfiguration calls the static method "Sanitizer.getDefaultConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sanitizer) TryGetDefaultConfiguration() (ret SanitizerConfig, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySanitizerGetDefaultConfiguration(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *SetHTMLOptions) UpdateFrom(ref js.Ref) {
	bindings.SetHTMLOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetHTMLOptions) Update(ref js.Ref) {
	bindings.SetHTMLOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetHTMLOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Sanitizer.Ref(),
	)
	p.Sanitizer = p.Sanitizer.FromRef(js.Undefined)
}

type DOMRectList struct {
	ref js.Ref
}

func (this DOMRectList) Once() DOMRectList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "DOMRectList.length".
//
// It returns ok=false if there is no such property.
func (this DOMRectList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMRectListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "DOMRectList.item" exists.
func (this DOMRectList) HasFuncItem() bool {
	return js.True == bindings.HasFuncDOMRectListItem(
		this.ref,
	)
}

// FuncItem returns the method "DOMRectList.item".
func (this DOMRectList) FuncItem() (fn js.Func[func(index uint32) DOMRect]) {
	bindings.FuncDOMRectListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "DOMRectList.item".
func (this DOMRectList) Item(index uint32) (ret DOMRect) {
	bindings.CallDOMRectListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMRectList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRectList) TryItem(index uint32) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CheckVisibilityOptions) UpdateFrom(ref js.Ref) {
	bindings.CheckVisibilityOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CheckVisibilityOptions) Update(ref js.Ref) {
	bindings.CheckVisibilityOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CheckVisibilityOptions) FreeMembers(recursive bool) {
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
func (p *ScrollIntoViewOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollIntoViewOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollIntoViewOptions) Update(ref js.Ref) {
	bindings.ScrollIntoViewOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollIntoViewOptions) FreeMembers(recursive bool) {
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
func (p *ScrollToOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollToOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollToOptions) Update(ref js.Ref) {
	bindings.ScrollToOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollToOptions) FreeMembers(recursive bool) {
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
func (p *TimelineRangeOffset) UpdateFrom(ref js.Ref) {
	bindings.TimelineRangeOffsetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TimelineRangeOffset) Update(ref js.Ref) {
	bindings.TimelineRangeOffsetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TimelineRangeOffset) FreeMembers(recursive bool) {
	js.Free(
		p.RangeName.Ref(),
		p.Offset.Ref(),
	)
	p.RangeName = p.RangeName.FromRef(js.Undefined)
	p.Offset = p.Offset.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "CSSKeywordValue.value".
//
// It returns ok=false if there is no such property.
func (this CSSKeywordValue) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeywordValueValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "CSSKeywordValue.value" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeywordValue) SetValue(val js.String) bool {
	return js.True == bindings.SetCSSKeywordValueValue(
		this.ref,
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
func (p *KeyframeAnimationOptions) UpdateFrom(ref js.Ref) {
	bindings.KeyframeAnimationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyframeAnimationOptions) Update(ref js.Ref) {
	bindings.KeyframeAnimationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyframeAnimationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Timeline.Ref(),
		p.RangeStart.Ref(),
		p.RangeEnd.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Timeline = p.Timeline.FromRef(js.Undefined)
	p.RangeStart = p.RangeStart.FromRef(js.Undefined)
	p.RangeEnd = p.RangeEnd.FromRef(js.Undefined)
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
func (p *GetAnimationsOptions) UpdateFrom(ref js.Ref) {
	bindings.GetAnimationsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAnimationsOptions) Update(ref js.Ref) {
	bindings.GetAnimationsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAnimationsOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// CommonAncestorContainer returns the value of property "Range.commonAncestorContainer".
//
// It returns ok=false if there is no such property.
func (this Range) CommonAncestorContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetRangeCommonAncestorContainer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetStart returns true if the method "Range.setStart" exists.
func (this Range) HasFuncSetStart() bool {
	return js.True == bindings.HasFuncRangeSetStart(
		this.ref,
	)
}

// FuncSetStart returns the method "Range.setStart".
func (this Range) FuncSetStart() (fn js.Func[func(node Node, offset uint32)]) {
	bindings.FuncRangeSetStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStart calls the method "Range.setStart".
func (this Range) SetStart(node Node, offset uint32) (ret js.Void) {
	bindings.CallRangeSetStart(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetStart calls the method "Range.setStart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetStart(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncSetEnd returns true if the method "Range.setEnd" exists.
func (this Range) HasFuncSetEnd() bool {
	return js.True == bindings.HasFuncRangeSetEnd(
		this.ref,
	)
}

// FuncSetEnd returns the method "Range.setEnd".
func (this Range) FuncSetEnd() (fn js.Func[func(node Node, offset uint32)]) {
	bindings.FuncRangeSetEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetEnd calls the method "Range.setEnd".
func (this Range) SetEnd(node Node, offset uint32) (ret js.Void) {
	bindings.CallRangeSetEnd(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetEnd calls the method "Range.setEnd"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetEnd(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncSetStartBefore returns true if the method "Range.setStartBefore" exists.
func (this Range) HasFuncSetStartBefore() bool {
	return js.True == bindings.HasFuncRangeSetStartBefore(
		this.ref,
	)
}

// FuncSetStartBefore returns the method "Range.setStartBefore".
func (this Range) FuncSetStartBefore() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSetStartBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStartBefore calls the method "Range.setStartBefore".
func (this Range) SetStartBefore(node Node) (ret js.Void) {
	bindings.CallRangeSetStartBefore(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetStartBefore calls the method "Range.setStartBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetStartBefore(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStartBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSetStartAfter returns true if the method "Range.setStartAfter" exists.
func (this Range) HasFuncSetStartAfter() bool {
	return js.True == bindings.HasFuncRangeSetStartAfter(
		this.ref,
	)
}

// FuncSetStartAfter returns the method "Range.setStartAfter".
func (this Range) FuncSetStartAfter() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSetStartAfter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStartAfter calls the method "Range.setStartAfter".
func (this Range) SetStartAfter(node Node) (ret js.Void) {
	bindings.CallRangeSetStartAfter(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetStartAfter calls the method "Range.setStartAfter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetStartAfter(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetStartAfter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSetEndBefore returns true if the method "Range.setEndBefore" exists.
func (this Range) HasFuncSetEndBefore() bool {
	return js.True == bindings.HasFuncRangeSetEndBefore(
		this.ref,
	)
}

// FuncSetEndBefore returns the method "Range.setEndBefore".
func (this Range) FuncSetEndBefore() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSetEndBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetEndBefore calls the method "Range.setEndBefore".
func (this Range) SetEndBefore(node Node) (ret js.Void) {
	bindings.CallRangeSetEndBefore(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetEndBefore calls the method "Range.setEndBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetEndBefore(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEndBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSetEndAfter returns true if the method "Range.setEndAfter" exists.
func (this Range) HasFuncSetEndAfter() bool {
	return js.True == bindings.HasFuncRangeSetEndAfter(
		this.ref,
	)
}

// FuncSetEndAfter returns the method "Range.setEndAfter".
func (this Range) FuncSetEndAfter() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSetEndAfter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetEndAfter calls the method "Range.setEndAfter".
func (this Range) SetEndAfter(node Node) (ret js.Void) {
	bindings.CallRangeSetEndAfter(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetEndAfter calls the method "Range.setEndAfter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySetEndAfter(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSetEndAfter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncCollapse returns true if the method "Range.collapse" exists.
func (this Range) HasFuncCollapse() bool {
	return js.True == bindings.HasFuncRangeCollapse(
		this.ref,
	)
}

// FuncCollapse returns the method "Range.collapse".
func (this Range) FuncCollapse() (fn js.Func[func(toStart bool)]) {
	bindings.FuncRangeCollapse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Collapse calls the method "Range.collapse".
func (this Range) Collapse(toStart bool) (ret js.Void) {
	bindings.CallRangeCollapse(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(toStart)),
	)

	return
}

// TryCollapse calls the method "Range.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCollapse(toStart bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCollapse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(toStart)),
	)

	return
}

// HasFuncCollapse1 returns true if the method "Range.collapse" exists.
func (this Range) HasFuncCollapse1() bool {
	return js.True == bindings.HasFuncRangeCollapse1(
		this.ref,
	)
}

// FuncCollapse1 returns the method "Range.collapse".
func (this Range) FuncCollapse1() (fn js.Func[func()]) {
	bindings.FuncRangeCollapse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Collapse1 calls the method "Range.collapse".
func (this Range) Collapse1() (ret js.Void) {
	bindings.CallRangeCollapse1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCollapse1 calls the method "Range.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCollapse1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCollapse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSelectNode returns true if the method "Range.selectNode" exists.
func (this Range) HasFuncSelectNode() bool {
	return js.True == bindings.HasFuncRangeSelectNode(
		this.ref,
	)
}

// FuncSelectNode returns the method "Range.selectNode".
func (this Range) FuncSelectNode() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSelectNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SelectNode calls the method "Range.selectNode".
func (this Range) SelectNode(node Node) (ret js.Void) {
	bindings.CallRangeSelectNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectNode calls the method "Range.selectNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySelectNode(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSelectNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSelectNodeContents returns true if the method "Range.selectNodeContents" exists.
func (this Range) HasFuncSelectNodeContents() bool {
	return js.True == bindings.HasFuncRangeSelectNodeContents(
		this.ref,
	)
}

// FuncSelectNodeContents returns the method "Range.selectNodeContents".
func (this Range) FuncSelectNodeContents() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeSelectNodeContents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SelectNodeContents calls the method "Range.selectNodeContents".
func (this Range) SelectNodeContents(node Node) (ret js.Void) {
	bindings.CallRangeSelectNodeContents(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectNodeContents calls the method "Range.selectNodeContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySelectNodeContents(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSelectNodeContents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncCompareBoundaryPoints returns true if the method "Range.compareBoundaryPoints" exists.
func (this Range) HasFuncCompareBoundaryPoints() bool {
	return js.True == bindings.HasFuncRangeCompareBoundaryPoints(
		this.ref,
	)
}

// FuncCompareBoundaryPoints returns the method "Range.compareBoundaryPoints".
func (this Range) FuncCompareBoundaryPoints() (fn js.Func[func(how uint16, sourceRange Range) int16]) {
	bindings.FuncRangeCompareBoundaryPoints(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompareBoundaryPoints calls the method "Range.compareBoundaryPoints".
func (this Range) CompareBoundaryPoints(how uint16, sourceRange Range) (ret int16) {
	bindings.CallRangeCompareBoundaryPoints(
		this.ref, js.Pointer(&ret),
		uint32(how),
		sourceRange.Ref(),
	)

	return
}

// TryCompareBoundaryPoints calls the method "Range.compareBoundaryPoints"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCompareBoundaryPoints(how uint16, sourceRange Range) (ret int16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCompareBoundaryPoints(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(how),
		sourceRange.Ref(),
	)

	return
}

// HasFuncDeleteContents returns true if the method "Range.deleteContents" exists.
func (this Range) HasFuncDeleteContents() bool {
	return js.True == bindings.HasFuncRangeDeleteContents(
		this.ref,
	)
}

// FuncDeleteContents returns the method "Range.deleteContents".
func (this Range) FuncDeleteContents() (fn js.Func[func()]) {
	bindings.FuncRangeDeleteContents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteContents calls the method "Range.deleteContents".
func (this Range) DeleteContents() (ret js.Void) {
	bindings.CallRangeDeleteContents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeleteContents calls the method "Range.deleteContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryDeleteContents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeDeleteContents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExtractContents returns true if the method "Range.extractContents" exists.
func (this Range) HasFuncExtractContents() bool {
	return js.True == bindings.HasFuncRangeExtractContents(
		this.ref,
	)
}

// FuncExtractContents returns the method "Range.extractContents".
func (this Range) FuncExtractContents() (fn js.Func[func() DocumentFragment]) {
	bindings.FuncRangeExtractContents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExtractContents calls the method "Range.extractContents".
func (this Range) ExtractContents() (ret DocumentFragment) {
	bindings.CallRangeExtractContents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExtractContents calls the method "Range.extractContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryExtractContents() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeExtractContents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCloneContents returns true if the method "Range.cloneContents" exists.
func (this Range) HasFuncCloneContents() bool {
	return js.True == bindings.HasFuncRangeCloneContents(
		this.ref,
	)
}

// FuncCloneContents returns the method "Range.cloneContents".
func (this Range) FuncCloneContents() (fn js.Func[func() DocumentFragment]) {
	bindings.FuncRangeCloneContents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CloneContents calls the method "Range.cloneContents".
func (this Range) CloneContents() (ret DocumentFragment) {
	bindings.CallRangeCloneContents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCloneContents calls the method "Range.cloneContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCloneContents() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCloneContents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertNode returns true if the method "Range.insertNode" exists.
func (this Range) HasFuncInsertNode() bool {
	return js.True == bindings.HasFuncRangeInsertNode(
		this.ref,
	)
}

// FuncInsertNode returns the method "Range.insertNode".
func (this Range) FuncInsertNode() (fn js.Func[func(node Node)]) {
	bindings.FuncRangeInsertNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertNode calls the method "Range.insertNode".
func (this Range) InsertNode(node Node) (ret js.Void) {
	bindings.CallRangeInsertNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryInsertNode calls the method "Range.insertNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryInsertNode(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeInsertNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSurroundContents returns true if the method "Range.surroundContents" exists.
func (this Range) HasFuncSurroundContents() bool {
	return js.True == bindings.HasFuncRangeSurroundContents(
		this.ref,
	)
}

// FuncSurroundContents returns the method "Range.surroundContents".
func (this Range) FuncSurroundContents() (fn js.Func[func(newParent Node)]) {
	bindings.FuncRangeSurroundContents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SurroundContents calls the method "Range.surroundContents".
func (this Range) SurroundContents(newParent Node) (ret js.Void) {
	bindings.CallRangeSurroundContents(
		this.ref, js.Pointer(&ret),
		newParent.Ref(),
	)

	return
}

// TrySurroundContents calls the method "Range.surroundContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TrySurroundContents(newParent Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeSurroundContents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newParent.Ref(),
	)

	return
}

// HasFuncCloneRange returns true if the method "Range.cloneRange" exists.
func (this Range) HasFuncCloneRange() bool {
	return js.True == bindings.HasFuncRangeCloneRange(
		this.ref,
	)
}

// FuncCloneRange returns the method "Range.cloneRange".
func (this Range) FuncCloneRange() (fn js.Func[func() Range]) {
	bindings.FuncRangeCloneRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CloneRange calls the method "Range.cloneRange".
func (this Range) CloneRange() (ret Range) {
	bindings.CallRangeCloneRange(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCloneRange calls the method "Range.cloneRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCloneRange() (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCloneRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDetach returns true if the method "Range.detach" exists.
func (this Range) HasFuncDetach() bool {
	return js.True == bindings.HasFuncRangeDetach(
		this.ref,
	)
}

// FuncDetach returns the method "Range.detach".
func (this Range) FuncDetach() (fn js.Func[func()]) {
	bindings.FuncRangeDetach(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Detach calls the method "Range.detach".
func (this Range) Detach() (ret js.Void) {
	bindings.CallRangeDetach(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDetach calls the method "Range.detach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryDetach() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeDetach(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsPointInRange returns true if the method "Range.isPointInRange" exists.
func (this Range) HasFuncIsPointInRange() bool {
	return js.True == bindings.HasFuncRangeIsPointInRange(
		this.ref,
	)
}

// FuncIsPointInRange returns the method "Range.isPointInRange".
func (this Range) FuncIsPointInRange() (fn js.Func[func(node Node, offset uint32) bool]) {
	bindings.FuncRangeIsPointInRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInRange calls the method "Range.isPointInRange".
func (this Range) IsPointInRange(node Node, offset uint32) (ret bool) {
	bindings.CallRangeIsPointInRange(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryIsPointInRange calls the method "Range.isPointInRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryIsPointInRange(node Node, offset uint32) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeIsPointInRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncComparePoint returns true if the method "Range.comparePoint" exists.
func (this Range) HasFuncComparePoint() bool {
	return js.True == bindings.HasFuncRangeComparePoint(
		this.ref,
	)
}

// FuncComparePoint returns the method "Range.comparePoint".
func (this Range) FuncComparePoint() (fn js.Func[func(node Node, offset uint32) int16]) {
	bindings.FuncRangeComparePoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ComparePoint calls the method "Range.comparePoint".
func (this Range) ComparePoint(node Node, offset uint32) (ret int16) {
	bindings.CallRangeComparePoint(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryComparePoint calls the method "Range.comparePoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryComparePoint(node Node, offset uint32) (ret int16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeComparePoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncIntersectsNode returns true if the method "Range.intersectsNode" exists.
func (this Range) HasFuncIntersectsNode() bool {
	return js.True == bindings.HasFuncRangeIntersectsNode(
		this.ref,
	)
}

// FuncIntersectsNode returns the method "Range.intersectsNode".
func (this Range) FuncIntersectsNode() (fn js.Func[func(node Node) bool]) {
	bindings.FuncRangeIntersectsNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IntersectsNode calls the method "Range.intersectsNode".
func (this Range) IntersectsNode(node Node) (ret bool) {
	bindings.CallRangeIntersectsNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryIntersectsNode calls the method "Range.intersectsNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryIntersectsNode(node Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeIntersectsNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncToString returns true if the method "Range.toString" exists.
func (this Range) HasFuncToString() bool {
	return js.True == bindings.HasFuncRangeToString(
		this.ref,
	)
}

// FuncToString returns the method "Range.toString".
func (this Range) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncRangeToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "Range.toString".
func (this Range) ToString() (ret js.String) {
	bindings.CallRangeToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "Range.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetClientRects returns true if the method "Range.getClientRects" exists.
func (this Range) HasFuncGetClientRects() bool {
	return js.True == bindings.HasFuncRangeGetClientRects(
		this.ref,
	)
}

// FuncGetClientRects returns the method "Range.getClientRects".
func (this Range) FuncGetClientRects() (fn js.Func[func() DOMRectList]) {
	bindings.FuncRangeGetClientRects(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetClientRects calls the method "Range.getClientRects".
func (this Range) GetClientRects() (ret DOMRectList) {
	bindings.CallRangeGetClientRects(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetClientRects calls the method "Range.getClientRects"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryGetClientRects() (ret DOMRectList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeGetClientRects(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBoundingClientRect returns true if the method "Range.getBoundingClientRect" exists.
func (this Range) HasFuncGetBoundingClientRect() bool {
	return js.True == bindings.HasFuncRangeGetBoundingClientRect(
		this.ref,
	)
}

// FuncGetBoundingClientRect returns the method "Range.getBoundingClientRect".
func (this Range) FuncGetBoundingClientRect() (fn js.Func[func() DOMRect]) {
	bindings.FuncRangeGetBoundingClientRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoundingClientRect calls the method "Range.getBoundingClientRect".
func (this Range) GetBoundingClientRect() (ret DOMRect) {
	bindings.CallRangeGetBoundingClientRect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoundingClientRect calls the method "Range.getBoundingClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryGetBoundingClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeGetBoundingClientRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateContextualFragment returns true if the method "Range.createContextualFragment" exists.
func (this Range) HasFuncCreateContextualFragment() bool {
	return js.True == bindings.HasFuncRangeCreateContextualFragment(
		this.ref,
	)
}

// FuncCreateContextualFragment returns the method "Range.createContextualFragment".
func (this Range) FuncCreateContextualFragment() (fn js.Func[func(fragment js.String) DocumentFragment]) {
	bindings.FuncRangeCreateContextualFragment(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContextualFragment calls the method "Range.createContextualFragment".
func (this Range) CreateContextualFragment(fragment js.String) (ret DocumentFragment) {
	bindings.CallRangeCreateContextualFragment(
		this.ref, js.Pointer(&ret),
		fragment.Ref(),
	)

	return
}

// TryCreateContextualFragment calls the method "Range.createContextualFragment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Range) TryCreateContextualFragment(fragment js.String) (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRangeCreateContextualFragment(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		fragment.Ref(),
	)

	return
}

type DOMTokenList struct {
	ref js.Ref
}

func (this DOMTokenList) Once() DOMTokenList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "DOMTokenList.length".
//
// It returns ok=false if there is no such property.
func (this DOMTokenList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMTokenListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "DOMTokenList.value".
//
// It returns ok=false if there is no such property.
func (this DOMTokenList) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMTokenListValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "DOMTokenList.value" to val.
//
// It returns false if the property cannot be set.
func (this DOMTokenList) SetValue(val js.String) bool {
	return js.True == bindings.SetDOMTokenListValue(
		this.ref,
		val.Ref(),
	)
}

// HasFuncItem returns true if the method "DOMTokenList.item" exists.
func (this DOMTokenList) HasFuncItem() bool {
	return js.True == bindings.HasFuncDOMTokenListItem(
		this.ref,
	)
}

// FuncItem returns the method "DOMTokenList.item".
func (this DOMTokenList) FuncItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncDOMTokenListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "DOMTokenList.item".
func (this DOMTokenList) Item(index uint32) (ret js.String) {
	bindings.CallDOMTokenListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMTokenList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncContains returns true if the method "DOMTokenList.contains" exists.
func (this DOMTokenList) HasFuncContains() bool {
	return js.True == bindings.HasFuncDOMTokenListContains(
		this.ref,
	)
}

// FuncContains returns the method "DOMTokenList.contains".
func (this DOMTokenList) FuncContains() (fn js.Func[func(token js.String) bool]) {
	bindings.FuncDOMTokenListContains(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Contains calls the method "DOMTokenList.contains".
func (this DOMTokenList) Contains(token js.String) (ret bool) {
	bindings.CallDOMTokenListContains(
		this.ref, js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TryContains calls the method "DOMTokenList.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryContains(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListContains(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
}

// HasFuncAdd returns true if the method "DOMTokenList.add" exists.
func (this DOMTokenList) HasFuncAdd() bool {
	return js.True == bindings.HasFuncDOMTokenListAdd(
		this.ref,
	)
}

// FuncAdd returns the method "DOMTokenList.add".
func (this DOMTokenList) FuncAdd() (fn js.Func[func(tokens ...js.String)]) {
	bindings.FuncDOMTokenListAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "DOMTokenList.add".
func (this DOMTokenList) Add(tokens ...js.String) (ret js.Void) {
	bindings.CallDOMTokenListAdd(
		this.ref, js.Pointer(&ret),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// TryAdd calls the method "DOMTokenList.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryAdd(tokens ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// HasFuncRemove returns true if the method "DOMTokenList.remove" exists.
func (this DOMTokenList) HasFuncRemove() bool {
	return js.True == bindings.HasFuncDOMTokenListRemove(
		this.ref,
	)
}

// FuncRemove returns the method "DOMTokenList.remove".
func (this DOMTokenList) FuncRemove() (fn js.Func[func(tokens ...js.String)]) {
	bindings.FuncDOMTokenListRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "DOMTokenList.remove".
func (this DOMTokenList) Remove(tokens ...js.String) (ret js.Void) {
	bindings.CallDOMTokenListRemove(
		this.ref, js.Pointer(&ret),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// TryRemove calls the method "DOMTokenList.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryRemove(tokens ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(tokens),
		js.SizeU(len(tokens)),
	)

	return
}

// HasFuncToggle returns true if the method "DOMTokenList.toggle" exists.
func (this DOMTokenList) HasFuncToggle() bool {
	return js.True == bindings.HasFuncDOMTokenListToggle(
		this.ref,
	)
}

// FuncToggle returns the method "DOMTokenList.toggle".
func (this DOMTokenList) FuncToggle() (fn js.Func[func(token js.String, force bool) bool]) {
	bindings.FuncDOMTokenListToggle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Toggle calls the method "DOMTokenList.toggle".
func (this DOMTokenList) Toggle(token js.String, force bool) (ret bool) {
	bindings.CallDOMTokenListToggle(
		this.ref, js.Pointer(&ret),
		token.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// TryToggle calls the method "DOMTokenList.toggle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryToggle(token js.String, force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListToggle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// HasFuncToggle1 returns true if the method "DOMTokenList.toggle" exists.
func (this DOMTokenList) HasFuncToggle1() bool {
	return js.True == bindings.HasFuncDOMTokenListToggle1(
		this.ref,
	)
}

// FuncToggle1 returns the method "DOMTokenList.toggle".
func (this DOMTokenList) FuncToggle1() (fn js.Func[func(token js.String) bool]) {
	bindings.FuncDOMTokenListToggle1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Toggle1 calls the method "DOMTokenList.toggle".
func (this DOMTokenList) Toggle1(token js.String) (ret bool) {
	bindings.CallDOMTokenListToggle1(
		this.ref, js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TryToggle1 calls the method "DOMTokenList.toggle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryToggle1(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListToggle1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
}

// HasFuncReplace returns true if the method "DOMTokenList.replace" exists.
func (this DOMTokenList) HasFuncReplace() bool {
	return js.True == bindings.HasFuncDOMTokenListReplace(
		this.ref,
	)
}

// FuncReplace returns the method "DOMTokenList.replace".
func (this DOMTokenList) FuncReplace() (fn js.Func[func(token js.String, newToken js.String) bool]) {
	bindings.FuncDOMTokenListReplace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Replace calls the method "DOMTokenList.replace".
func (this DOMTokenList) Replace(token js.String, newToken js.String) (ret bool) {
	bindings.CallDOMTokenListReplace(
		this.ref, js.Pointer(&ret),
		token.Ref(),
		newToken.Ref(),
	)

	return
}

// TryReplace calls the method "DOMTokenList.replace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TryReplace(token js.String, newToken js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListReplace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		newToken.Ref(),
	)

	return
}

// HasFuncSupports returns true if the method "DOMTokenList.supports" exists.
func (this DOMTokenList) HasFuncSupports() bool {
	return js.True == bindings.HasFuncDOMTokenListSupports(
		this.ref,
	)
}

// FuncSupports returns the method "DOMTokenList.supports".
func (this DOMTokenList) FuncSupports() (fn js.Func[func(token js.String) bool]) {
	bindings.FuncDOMTokenListSupports(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Supports calls the method "DOMTokenList.supports".
func (this DOMTokenList) Supports(token js.String) (ret bool) {
	bindings.CallDOMTokenListSupports(
		this.ref, js.Pointer(&ret),
		token.Ref(),
	)

	return
}

// TrySupports calls the method "DOMTokenList.supports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMTokenList) TrySupports(token js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMTokenListSupports(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
	)

	return
}

type NamedNodeMap struct {
	ref js.Ref
}

func (this NamedNodeMap) Once() NamedNodeMap {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "NamedNodeMap.length".
//
// It returns ok=false if there is no such property.
func (this NamedNodeMap) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNamedNodeMapLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "NamedNodeMap.item" exists.
func (this NamedNodeMap) HasFuncItem() bool {
	return js.True == bindings.HasFuncNamedNodeMapItem(
		this.ref,
	)
}

// FuncItem returns the method "NamedNodeMap.item".
func (this NamedNodeMap) FuncItem() (fn js.Func[func(index uint32) Attr]) {
	bindings.FuncNamedNodeMapItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "NamedNodeMap.item".
func (this NamedNodeMap) Item(index uint32) (ret Attr) {
	bindings.CallNamedNodeMapItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "NamedNodeMap.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TryItem(index uint32) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetNamedItem returns true if the method "NamedNodeMap.getNamedItem" exists.
func (this NamedNodeMap) HasFuncGetNamedItem() bool {
	return js.True == bindings.HasFuncNamedNodeMapGetNamedItem(
		this.ref,
	)
}

// FuncGetNamedItem returns the method "NamedNodeMap.getNamedItem".
func (this NamedNodeMap) FuncGetNamedItem() (fn js.Func[func(qualifiedName js.String) Attr]) {
	bindings.FuncNamedNodeMapGetNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetNamedItem calls the method "NamedNodeMap.getNamedItem".
func (this NamedNodeMap) GetNamedItem(qualifiedName js.String) (ret Attr) {
	bindings.CallNamedNodeMapGetNamedItem(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetNamedItem calls the method "NamedNodeMap.getNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TryGetNamedItem(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapGetNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncGetNamedItemNS returns true if the method "NamedNodeMap.getNamedItemNS" exists.
func (this NamedNodeMap) HasFuncGetNamedItemNS() bool {
	return js.True == bindings.HasFuncNamedNodeMapGetNamedItemNS(
		this.ref,
	)
}

// FuncGetNamedItemNS returns the method "NamedNodeMap.getNamedItemNS".
func (this NamedNodeMap) FuncGetNamedItemNS() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	bindings.FuncNamedNodeMapGetNamedItemNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetNamedItemNS calls the method "NamedNodeMap.getNamedItemNS".
func (this NamedNodeMap) GetNamedItemNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallNamedNodeMapGetNamedItemNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetNamedItemNS calls the method "NamedNodeMap.getNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TryGetNamedItemNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapGetNamedItemNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncSetNamedItem returns true if the method "NamedNodeMap.setNamedItem" exists.
func (this NamedNodeMap) HasFuncSetNamedItem() bool {
	return js.True == bindings.HasFuncNamedNodeMapSetNamedItem(
		this.ref,
	)
}

// FuncSetNamedItem returns the method "NamedNodeMap.setNamedItem".
func (this NamedNodeMap) FuncSetNamedItem() (fn js.Func[func(attr Attr) Attr]) {
	bindings.FuncNamedNodeMapSetNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetNamedItem calls the method "NamedNodeMap.setNamedItem".
func (this NamedNodeMap) SetNamedItem(attr Attr) (ret Attr) {
	bindings.CallNamedNodeMapSetNamedItem(
		this.ref, js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetNamedItem calls the method "NamedNodeMap.setNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TrySetNamedItem(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapSetNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasFuncSetNamedItemNS returns true if the method "NamedNodeMap.setNamedItemNS" exists.
func (this NamedNodeMap) HasFuncSetNamedItemNS() bool {
	return js.True == bindings.HasFuncNamedNodeMapSetNamedItemNS(
		this.ref,
	)
}

// FuncSetNamedItemNS returns the method "NamedNodeMap.setNamedItemNS".
func (this NamedNodeMap) FuncSetNamedItemNS() (fn js.Func[func(attr Attr) Attr]) {
	bindings.FuncNamedNodeMapSetNamedItemNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetNamedItemNS calls the method "NamedNodeMap.setNamedItemNS".
func (this NamedNodeMap) SetNamedItemNS(attr Attr) (ret Attr) {
	bindings.CallNamedNodeMapSetNamedItemNS(
		this.ref, js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetNamedItemNS calls the method "NamedNodeMap.setNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TrySetNamedItemNS(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapSetNamedItemNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasFuncRemoveNamedItem returns true if the method "NamedNodeMap.removeNamedItem" exists.
func (this NamedNodeMap) HasFuncRemoveNamedItem() bool {
	return js.True == bindings.HasFuncNamedNodeMapRemoveNamedItem(
		this.ref,
	)
}

// FuncRemoveNamedItem returns the method "NamedNodeMap.removeNamedItem".
func (this NamedNodeMap) FuncRemoveNamedItem() (fn js.Func[func(qualifiedName js.String) Attr]) {
	bindings.FuncNamedNodeMapRemoveNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveNamedItem calls the method "NamedNodeMap.removeNamedItem".
func (this NamedNodeMap) RemoveNamedItem(qualifiedName js.String) (ret Attr) {
	bindings.CallNamedNodeMapRemoveNamedItem(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryRemoveNamedItem calls the method "NamedNodeMap.removeNamedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TryRemoveNamedItem(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapRemoveNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncRemoveNamedItemNS returns true if the method "NamedNodeMap.removeNamedItemNS" exists.
func (this NamedNodeMap) HasFuncRemoveNamedItemNS() bool {
	return js.True == bindings.HasFuncNamedNodeMapRemoveNamedItemNS(
		this.ref,
	)
}

// FuncRemoveNamedItemNS returns the method "NamedNodeMap.removeNamedItemNS".
func (this NamedNodeMap) FuncRemoveNamedItemNS() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	bindings.FuncNamedNodeMapRemoveNamedItemNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveNamedItemNS calls the method "NamedNodeMap.removeNamedItemNS".
func (this NamedNodeMap) RemoveNamedItemNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallNamedNodeMapRemoveNamedItemNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryRemoveNamedItemNS calls the method "NamedNodeMap.removeNamedItemNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedNodeMap) TryRemoveNamedItemNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedNodeMapRemoveNamedItemNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

type Element struct {
	Node
}

func (this Element) Once() Element {
	this.ref.Once()
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
	this.ref.Free()
}

// NamespaceURI returns the value of property "Element.namespaceURI".
//
// It returns ok=false if there is no such property.
func (this Element) NamespaceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementNamespaceURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Prefix returns the value of property "Element.prefix".
//
// It returns ok=false if there is no such property.
func (this Element) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementPrefix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LocalName returns the value of property "Element.localName".
//
// It returns ok=false if there is no such property.
func (this Element) LocalName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementLocalName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TagName returns the value of property "Element.tagName".
//
// It returns ok=false if there is no such property.
func (this Element) TagName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementTagName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "Element.id".
//
// It returns ok=false if there is no such property.
func (this Element) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "Element.id" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetId(val js.String) bool {
	return js.True == bindings.SetElementId(
		this.ref,
		val.Ref(),
	)
}

// ClassName returns the value of property "Element.className".
//
// It returns ok=false if there is no such property.
func (this Element) ClassName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementClassName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetClassName sets the value of property "Element.className" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetClassName(val js.String) bool {
	return js.True == bindings.SetElementClassName(
		this.ref,
		val.Ref(),
	)
}

// ClassList returns the value of property "Element.classList".
//
// It returns ok=false if there is no such property.
func (this Element) ClassList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetElementClassList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Slot returns the value of property "Element.slot".
//
// It returns ok=false if there is no such property.
func (this Element) Slot() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementSlot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSlot sets the value of property "Element.slot" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetSlot(val js.String) bool {
	return js.True == bindings.SetElementSlot(
		this.ref,
		val.Ref(),
	)
}

// Attributes returns the value of property "Element.attributes".
//
// It returns ok=false if there is no such property.
func (this Element) Attributes() (ret NamedNodeMap, ok bool) {
	ok = js.True == bindings.GetElementAttributes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ShadowRoot returns the value of property "Element.shadowRoot".
//
// It returns ok=false if there is no such property.
func (this Element) ShadowRoot() (ret ShadowRoot, ok bool) {
	ok = js.True == bindings.GetElementShadowRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ElementTiming returns the value of property "Element.elementTiming".
//
// It returns ok=false if there is no such property.
func (this Element) ElementTiming() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementElementTiming(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetElementTiming sets the value of property "Element.elementTiming" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetElementTiming(val js.String) bool {
	return js.True == bindings.SetElementElementTiming(
		this.ref,
		val.Ref(),
	)
}

// Part returns the value of property "Element.part".
//
// It returns ok=false if there is no such property.
func (this Element) Part() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetElementPart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OuterHTML returns the value of property "Element.outerHTML".
//
// It returns ok=false if there is no such property.
func (this Element) OuterHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementOuterHTML(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOuterHTML sets the value of property "Element.outerHTML" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetOuterHTML(val js.String) bool {
	return js.True == bindings.SetElementOuterHTML(
		this.ref,
		val.Ref(),
	)
}

// ScrollTop returns the value of property "Element.scrollTop".
//
// It returns ok=false if there is no such property.
func (this Element) ScrollTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetElementScrollTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollTop sets the value of property "Element.scrollTop" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetScrollTop(val float64) bool {
	return js.True == bindings.SetElementScrollTop(
		this.ref,
		float64(val),
	)
}

// ScrollLeft returns the value of property "Element.scrollLeft".
//
// It returns ok=false if there is no such property.
func (this Element) ScrollLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetElementScrollLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollLeft sets the value of property "Element.scrollLeft" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetScrollLeft(val float64) bool {
	return js.True == bindings.SetElementScrollLeft(
		this.ref,
		float64(val),
	)
}

// ScrollWidth returns the value of property "Element.scrollWidth".
//
// It returns ok=false if there is no such property.
func (this Element) ScrollWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementScrollWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScrollHeight returns the value of property "Element.scrollHeight".
//
// It returns ok=false if there is no such property.
func (this Element) ScrollHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementScrollHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientTop returns the value of property "Element.clientTop".
//
// It returns ok=false if there is no such property.
func (this Element) ClientTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientLeft returns the value of property "Element.clientLeft".
//
// It returns ok=false if there is no such property.
func (this Element) ClientLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientWidth returns the value of property "Element.clientWidth".
//
// It returns ok=false if there is no such property.
func (this Element) ClientWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientHeight returns the value of property "Element.clientHeight".
//
// It returns ok=false if there is no such property.
func (this Element) ClientHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetElementClientHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Role returns the value of property "Element.role".
//
// It returns ok=false if there is no such property.
func (this Element) Role() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementRole(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRole sets the value of property "Element.role" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetRole(val js.String) bool {
	return js.True == bindings.SetElementRole(
		this.ref,
		val.Ref(),
	)
}

// AriaActiveDescendantElement returns the value of property "Element.ariaActiveDescendantElement".
//
// It returns ok=false if there is no such property.
func (this Element) AriaActiveDescendantElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementAriaActiveDescendantElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaActiveDescendantElement sets the value of property "Element.ariaActiveDescendantElement" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaActiveDescendantElement(val Element) bool {
	return js.True == bindings.SetElementAriaActiveDescendantElement(
		this.ref,
		val.Ref(),
	)
}

// AriaAtomic returns the value of property "Element.ariaAtomic".
//
// It returns ok=false if there is no such property.
func (this Element) AriaAtomic() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaAtomic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaAtomic sets the value of property "Element.ariaAtomic" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaAtomic(val js.String) bool {
	return js.True == bindings.SetElementAriaAtomic(
		this.ref,
		val.Ref(),
	)
}

// AriaAutoComplete returns the value of property "Element.ariaAutoComplete".
//
// It returns ok=false if there is no such property.
func (this Element) AriaAutoComplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaAutoComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaAutoComplete sets the value of property "Element.ariaAutoComplete" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaAutoComplete(val js.String) bool {
	return js.True == bindings.SetElementAriaAutoComplete(
		this.ref,
		val.Ref(),
	)
}

// AriaBusy returns the value of property "Element.ariaBusy".
//
// It returns ok=false if there is no such property.
func (this Element) AriaBusy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaBusy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaBusy sets the value of property "Element.ariaBusy" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaBusy(val js.String) bool {
	return js.True == bindings.SetElementAriaBusy(
		this.ref,
		val.Ref(),
	)
}

// AriaChecked returns the value of property "Element.ariaChecked".
//
// It returns ok=false if there is no such property.
func (this Element) AriaChecked() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaChecked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaChecked sets the value of property "Element.ariaChecked" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaChecked(val js.String) bool {
	return js.True == bindings.SetElementAriaChecked(
		this.ref,
		val.Ref(),
	)
}

// AriaColCount returns the value of property "Element.ariaColCount".
//
// It returns ok=false if there is no such property.
func (this Element) AriaColCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColCount sets the value of property "Element.ariaColCount" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColCount(val js.String) bool {
	return js.True == bindings.SetElementAriaColCount(
		this.ref,
		val.Ref(),
	)
}

// AriaColIndex returns the value of property "Element.ariaColIndex".
//
// It returns ok=false if there is no such property.
func (this Element) AriaColIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColIndex sets the value of property "Element.ariaColIndex" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColIndex(val js.String) bool {
	return js.True == bindings.SetElementAriaColIndex(
		this.ref,
		val.Ref(),
	)
}

// AriaColIndexText returns the value of property "Element.ariaColIndexText".
//
// It returns ok=false if there is no such property.
func (this Element) AriaColIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColIndexText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColIndexText sets the value of property "Element.ariaColIndexText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColIndexText(val js.String) bool {
	return js.True == bindings.SetElementAriaColIndexText(
		this.ref,
		val.Ref(),
	)
}

// AriaColSpan returns the value of property "Element.ariaColSpan".
//
// It returns ok=false if there is no such property.
func (this Element) AriaColSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaColSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColSpan sets the value of property "Element.ariaColSpan" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaColSpan(val js.String) bool {
	return js.True == bindings.SetElementAriaColSpan(
		this.ref,
		val.Ref(),
	)
}

// AriaControlsElements returns the value of property "Element.ariaControlsElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaControlsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaControlsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaControlsElements sets the value of property "Element.ariaControlsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaControlsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaControlsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaCurrent returns the value of property "Element.ariaCurrent".
//
// It returns ok=false if there is no such property.
func (this Element) AriaCurrent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaCurrent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaCurrent sets the value of property "Element.ariaCurrent" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaCurrent(val js.String) bool {
	return js.True == bindings.SetElementAriaCurrent(
		this.ref,
		val.Ref(),
	)
}

// AriaDescribedByElements returns the value of property "Element.ariaDescribedByElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaDescribedByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaDescribedByElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDescribedByElements sets the value of property "Element.ariaDescribedByElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDescribedByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaDescribedByElements(
		this.ref,
		val.Ref(),
	)
}

// AriaDescription returns the value of property "Element.ariaDescription".
//
// It returns ok=false if there is no such property.
func (this Element) AriaDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDescription sets the value of property "Element.ariaDescription" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDescription(val js.String) bool {
	return js.True == bindings.SetElementAriaDescription(
		this.ref,
		val.Ref(),
	)
}

// AriaDetailsElements returns the value of property "Element.ariaDetailsElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaDetailsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaDetailsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDetailsElements sets the value of property "Element.ariaDetailsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDetailsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaDetailsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaDisabled returns the value of property "Element.ariaDisabled".
//
// It returns ok=false if there is no such property.
func (this Element) AriaDisabled() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDisabled sets the value of property "Element.ariaDisabled" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaDisabled(val js.String) bool {
	return js.True == bindings.SetElementAriaDisabled(
		this.ref,
		val.Ref(),
	)
}

// AriaErrorMessageElements returns the value of property "Element.ariaErrorMessageElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaErrorMessageElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaErrorMessageElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaErrorMessageElements sets the value of property "Element.ariaErrorMessageElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaErrorMessageElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaErrorMessageElements(
		this.ref,
		val.Ref(),
	)
}

// AriaExpanded returns the value of property "Element.ariaExpanded".
//
// It returns ok=false if there is no such property.
func (this Element) AriaExpanded() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaExpanded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaExpanded sets the value of property "Element.ariaExpanded" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaExpanded(val js.String) bool {
	return js.True == bindings.SetElementAriaExpanded(
		this.ref,
		val.Ref(),
	)
}

// AriaFlowToElements returns the value of property "Element.ariaFlowToElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaFlowToElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaFlowToElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaFlowToElements sets the value of property "Element.ariaFlowToElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaFlowToElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaFlowToElements(
		this.ref,
		val.Ref(),
	)
}

// AriaHasPopup returns the value of property "Element.ariaHasPopup".
//
// It returns ok=false if there is no such property.
func (this Element) AriaHasPopup() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaHasPopup(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaHasPopup sets the value of property "Element.ariaHasPopup" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaHasPopup(val js.String) bool {
	return js.True == bindings.SetElementAriaHasPopup(
		this.ref,
		val.Ref(),
	)
}

// AriaHidden returns the value of property "Element.ariaHidden".
//
// It returns ok=false if there is no such property.
func (this Element) AriaHidden() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaHidden(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaHidden sets the value of property "Element.ariaHidden" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaHidden(val js.String) bool {
	return js.True == bindings.SetElementAriaHidden(
		this.ref,
		val.Ref(),
	)
}

// AriaInvalid returns the value of property "Element.ariaInvalid".
//
// It returns ok=false if there is no such property.
func (this Element) AriaInvalid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaInvalid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaInvalid sets the value of property "Element.ariaInvalid" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaInvalid(val js.String) bool {
	return js.True == bindings.SetElementAriaInvalid(
		this.ref,
		val.Ref(),
	)
}

// AriaKeyShortcuts returns the value of property "Element.ariaKeyShortcuts".
//
// It returns ok=false if there is no such property.
func (this Element) AriaKeyShortcuts() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaKeyShortcuts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaKeyShortcuts sets the value of property "Element.ariaKeyShortcuts" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaKeyShortcuts(val js.String) bool {
	return js.True == bindings.SetElementAriaKeyShortcuts(
		this.ref,
		val.Ref(),
	)
}

// AriaLabel returns the value of property "Element.ariaLabel".
//
// It returns ok=false if there is no such property.
func (this Element) AriaLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLabel sets the value of property "Element.ariaLabel" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLabel(val js.String) bool {
	return js.True == bindings.SetElementAriaLabel(
		this.ref,
		val.Ref(),
	)
}

// AriaLabelledByElements returns the value of property "Element.ariaLabelledByElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaLabelledByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaLabelledByElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLabelledByElements sets the value of property "Element.ariaLabelledByElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLabelledByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaLabelledByElements(
		this.ref,
		val.Ref(),
	)
}

// AriaLevel returns the value of property "Element.ariaLevel".
//
// It returns ok=false if there is no such property.
func (this Element) AriaLevel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLevel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLevel sets the value of property "Element.ariaLevel" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLevel(val js.String) bool {
	return js.True == bindings.SetElementAriaLevel(
		this.ref,
		val.Ref(),
	)
}

// AriaLive returns the value of property "Element.ariaLive".
//
// It returns ok=false if there is no such property.
func (this Element) AriaLive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaLive(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLive sets the value of property "Element.ariaLive" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaLive(val js.String) bool {
	return js.True == bindings.SetElementAriaLive(
		this.ref,
		val.Ref(),
	)
}

// AriaModal returns the value of property "Element.ariaModal".
//
// It returns ok=false if there is no such property.
func (this Element) AriaModal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaModal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaModal sets the value of property "Element.ariaModal" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaModal(val js.String) bool {
	return js.True == bindings.SetElementAriaModal(
		this.ref,
		val.Ref(),
	)
}

// AriaMultiLine returns the value of property "Element.ariaMultiLine".
//
// It returns ok=false if there is no such property.
func (this Element) AriaMultiLine() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaMultiLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaMultiLine sets the value of property "Element.ariaMultiLine" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaMultiLine(val js.String) bool {
	return js.True == bindings.SetElementAriaMultiLine(
		this.ref,
		val.Ref(),
	)
}

// AriaMultiSelectable returns the value of property "Element.ariaMultiSelectable".
//
// It returns ok=false if there is no such property.
func (this Element) AriaMultiSelectable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaMultiSelectable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaMultiSelectable sets the value of property "Element.ariaMultiSelectable" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaMultiSelectable(val js.String) bool {
	return js.True == bindings.SetElementAriaMultiSelectable(
		this.ref,
		val.Ref(),
	)
}

// AriaOrientation returns the value of property "Element.ariaOrientation".
//
// It returns ok=false if there is no such property.
func (this Element) AriaOrientation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaOrientation sets the value of property "Element.ariaOrientation" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaOrientation(val js.String) bool {
	return js.True == bindings.SetElementAriaOrientation(
		this.ref,
		val.Ref(),
	)
}

// AriaOwnsElements returns the value of property "Element.ariaOwnsElements".
//
// It returns ok=false if there is no such property.
func (this Element) AriaOwnsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementAriaOwnsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaOwnsElements sets the value of property "Element.ariaOwnsElements" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaOwnsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementAriaOwnsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaPlaceholder returns the value of property "Element.ariaPlaceholder".
//
// It returns ok=false if there is no such property.
func (this Element) AriaPlaceholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPlaceholder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPlaceholder sets the value of property "Element.ariaPlaceholder" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPlaceholder(val js.String) bool {
	return js.True == bindings.SetElementAriaPlaceholder(
		this.ref,
		val.Ref(),
	)
}

// AriaPosInSet returns the value of property "Element.ariaPosInSet".
//
// It returns ok=false if there is no such property.
func (this Element) AriaPosInSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPosInSet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPosInSet sets the value of property "Element.ariaPosInSet" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPosInSet(val js.String) bool {
	return js.True == bindings.SetElementAriaPosInSet(
		this.ref,
		val.Ref(),
	)
}

// AriaPressed returns the value of property "Element.ariaPressed".
//
// It returns ok=false if there is no such property.
func (this Element) AriaPressed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaPressed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPressed sets the value of property "Element.ariaPressed" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaPressed(val js.String) bool {
	return js.True == bindings.SetElementAriaPressed(
		this.ref,
		val.Ref(),
	)
}

// AriaReadOnly returns the value of property "Element.ariaReadOnly".
//
// It returns ok=false if there is no such property.
func (this Element) AriaReadOnly() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaReadOnly(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaReadOnly sets the value of property "Element.ariaReadOnly" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaReadOnly(val js.String) bool {
	return js.True == bindings.SetElementAriaReadOnly(
		this.ref,
		val.Ref(),
	)
}

// AriaRequired returns the value of property "Element.ariaRequired".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRequired() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRequired(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRequired sets the value of property "Element.ariaRequired" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRequired(val js.String) bool {
	return js.True == bindings.SetElementAriaRequired(
		this.ref,
		val.Ref(),
	)
}

// AriaRoleDescription returns the value of property "Element.ariaRoleDescription".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRoleDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRoleDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRoleDescription sets the value of property "Element.ariaRoleDescription" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRoleDescription(val js.String) bool {
	return js.True == bindings.SetElementAriaRoleDescription(
		this.ref,
		val.Ref(),
	)
}

// AriaRowCount returns the value of property "Element.ariaRowCount".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRowCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowCount sets the value of property "Element.ariaRowCount" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowCount(val js.String) bool {
	return js.True == bindings.SetElementAriaRowCount(
		this.ref,
		val.Ref(),
	)
}

// AriaRowIndex returns the value of property "Element.ariaRowIndex".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRowIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndex sets the value of property "Element.ariaRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowIndex(val js.String) bool {
	return js.True == bindings.SetElementAriaRowIndex(
		this.ref,
		val.Ref(),
	)
}

// AriaRowIndexText returns the value of property "Element.ariaRowIndexText".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRowIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowIndexText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndexText sets the value of property "Element.ariaRowIndexText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowIndexText(val js.String) bool {
	return js.True == bindings.SetElementAriaRowIndexText(
		this.ref,
		val.Ref(),
	)
}

// AriaRowSpan returns the value of property "Element.ariaRowSpan".
//
// It returns ok=false if there is no such property.
func (this Element) AriaRowSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaRowSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowSpan sets the value of property "Element.ariaRowSpan" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaRowSpan(val js.String) bool {
	return js.True == bindings.SetElementAriaRowSpan(
		this.ref,
		val.Ref(),
	)
}

// AriaSelected returns the value of property "Element.ariaSelected".
//
// It returns ok=false if there is no such property.
func (this Element) AriaSelected() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSelected sets the value of property "Element.ariaSelected" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSelected(val js.String) bool {
	return js.True == bindings.SetElementAriaSelected(
		this.ref,
		val.Ref(),
	)
}

// AriaSetSize returns the value of property "Element.ariaSetSize".
//
// It returns ok=false if there is no such property.
func (this Element) AriaSetSize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSetSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSetSize sets the value of property "Element.ariaSetSize" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSetSize(val js.String) bool {
	return js.True == bindings.SetElementAriaSetSize(
		this.ref,
		val.Ref(),
	)
}

// AriaSort returns the value of property "Element.ariaSort".
//
// It returns ok=false if there is no such property.
func (this Element) AriaSort() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaSort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSort sets the value of property "Element.ariaSort" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaSort(val js.String) bool {
	return js.True == bindings.SetElementAriaSort(
		this.ref,
		val.Ref(),
	)
}

// AriaValueMax returns the value of property "Element.ariaValueMax".
//
// It returns ok=false if there is no such property.
func (this Element) AriaValueMax() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueMax sets the value of property "Element.ariaValueMax" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueMax(val js.String) bool {
	return js.True == bindings.SetElementAriaValueMax(
		this.ref,
		val.Ref(),
	)
}

// AriaValueMin returns the value of property "Element.ariaValueMin".
//
// It returns ok=false if there is no such property.
func (this Element) AriaValueMin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueMin sets the value of property "Element.ariaValueMin" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueMin(val js.String) bool {
	return js.True == bindings.SetElementAriaValueMin(
		this.ref,
		val.Ref(),
	)
}

// AriaValueNow returns the value of property "Element.ariaValueNow".
//
// It returns ok=false if there is no such property.
func (this Element) AriaValueNow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueNow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueNow sets the value of property "Element.ariaValueNow" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueNow(val js.String) bool {
	return js.True == bindings.SetElementAriaValueNow(
		this.ref,
		val.Ref(),
	)
}

// AriaValueText returns the value of property "Element.ariaValueText".
//
// It returns ok=false if there is no such property.
func (this Element) AriaValueText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementAriaValueText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueText sets the value of property "Element.ariaValueText" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetAriaValueText(val js.String) bool {
	return js.True == bindings.SetElementAriaValueText(
		this.ref,
		val.Ref(),
	)
}

// InnerHTML returns the value of property "Element.innerHTML".
//
// It returns ok=false if there is no such property.
func (this Element) InnerHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInnerHTML(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInnerHTML sets the value of property "Element.innerHTML" to val.
//
// It returns false if the property cannot be set.
func (this Element) SetInnerHTML(val js.String) bool {
	return js.True == bindings.SetElementInnerHTML(
		this.ref,
		val.Ref(),
	)
}

// Children returns the value of property "Element.children".
//
// It returns ok=false if there is no such property.
func (this Element) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetElementChildren(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "Element.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this Element) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementFirstElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "Element.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this Element) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementLastElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "Element.childElementCount".
//
// It returns ok=false if there is no such property.
func (this Element) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetElementChildElementCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreviousElementSibling returns the value of property "Element.previousElementSibling".
//
// It returns ok=false if there is no such property.
func (this Element) PreviousElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementPreviousElementSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NextElementSibling returns the value of property "Element.nextElementSibling".
//
// It returns ok=false if there is no such property.
func (this Element) NextElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementNextElementSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AssignedSlot returns the value of property "Element.assignedSlot".
//
// It returns ok=false if there is no such property.
func (this Element) AssignedSlot() (ret HTMLSlotElement, ok bool) {
	ok = js.True == bindings.GetElementAssignedSlot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RegionOverset returns the value of property "Element.regionOverset".
//
// It returns ok=false if there is no such property.
func (this Element) RegionOverset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementRegionOverset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncHasAttributes returns true if the method "Element.hasAttributes" exists.
func (this Element) HasFuncHasAttributes() bool {
	return js.True == bindings.HasFuncElementHasAttributes(
		this.ref,
	)
}

// FuncHasAttributes returns the method "Element.hasAttributes".
func (this Element) FuncHasAttributes() (fn js.Func[func() bool]) {
	bindings.FuncElementHasAttributes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasAttributes calls the method "Element.hasAttributes".
func (this Element) HasAttributes() (ret bool) {
	bindings.CallElementHasAttributes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasAttributes calls the method "Element.hasAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryHasAttributes() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttributes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAttributeNames returns true if the method "Element.getAttributeNames" exists.
func (this Element) HasFuncGetAttributeNames() bool {
	return js.True == bindings.HasFuncElementGetAttributeNames(
		this.ref,
	)
}

// FuncGetAttributeNames returns the method "Element.getAttributeNames".
func (this Element) FuncGetAttributeNames() (fn js.Func[func() js.Array[js.String]]) {
	bindings.FuncElementGetAttributeNames(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttributeNames calls the method "Element.getAttributeNames".
func (this Element) GetAttributeNames() (ret js.Array[js.String]) {
	bindings.CallElementGetAttributeNames(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAttributeNames calls the method "Element.getAttributeNames"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAttributeNames() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNames(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAttribute returns true if the method "Element.getAttribute" exists.
func (this Element) HasFuncGetAttribute() bool {
	return js.True == bindings.HasFuncElementGetAttribute(
		this.ref,
	)
}

// FuncGetAttribute returns the method "Element.getAttribute".
func (this Element) FuncGetAttribute() (fn js.Func[func(qualifiedName js.String) js.String]) {
	bindings.FuncElementGetAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttribute calls the method "Element.getAttribute".
func (this Element) GetAttribute(qualifiedName js.String) (ret js.String) {
	bindings.CallElementGetAttribute(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetAttribute calls the method "Element.getAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAttribute(qualifiedName js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncGetAttributeNS returns true if the method "Element.getAttributeNS" exists.
func (this Element) HasFuncGetAttributeNS() bool {
	return js.True == bindings.HasFuncElementGetAttributeNS(
		this.ref,
	)
}

// FuncGetAttributeNS returns the method "Element.getAttributeNS".
func (this Element) FuncGetAttributeNS() (fn js.Func[func(namespace js.String, localName js.String) js.String]) {
	bindings.FuncElementGetAttributeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttributeNS calls the method "Element.getAttributeNS".
func (this Element) GetAttributeNS(namespace js.String, localName js.String) (ret js.String) {
	bindings.CallElementGetAttributeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetAttributeNS calls the method "Element.getAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAttributeNS(namespace js.String, localName js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncSetAttribute returns true if the method "Element.setAttribute" exists.
func (this Element) HasFuncSetAttribute() bool {
	return js.True == bindings.HasFuncElementSetAttribute(
		this.ref,
	)
}

// FuncSetAttribute returns the method "Element.setAttribute".
func (this Element) FuncSetAttribute() (fn js.Func[func(qualifiedName js.String, value js.String)]) {
	bindings.FuncElementSetAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAttribute calls the method "Element.setAttribute".
func (this Element) SetAttribute(qualifiedName js.String, value js.String) (ret js.Void) {
	bindings.CallElementSetAttribute(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// TrySetAttribute calls the method "Element.setAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetAttribute(qualifiedName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncSetAttributeNS returns true if the method "Element.setAttributeNS" exists.
func (this Element) HasFuncSetAttributeNS() bool {
	return js.True == bindings.HasFuncElementSetAttributeNS(
		this.ref,
	)
}

// FuncSetAttributeNS returns the method "Element.setAttributeNS".
func (this Element) FuncSetAttributeNS() (fn js.Func[func(namespace js.String, qualifiedName js.String, value js.String)]) {
	bindings.FuncElementSetAttributeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAttributeNS calls the method "Element.setAttributeNS".
func (this Element) SetAttributeNS(namespace js.String, qualifiedName js.String, value js.String) (ret js.Void) {
	bindings.CallElementSetAttributeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// TrySetAttributeNS calls the method "Element.setAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetAttributeNS(namespace js.String, qualifiedName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncRemoveAttribute returns true if the method "Element.removeAttribute" exists.
func (this Element) HasFuncRemoveAttribute() bool {
	return js.True == bindings.HasFuncElementRemoveAttribute(
		this.ref,
	)
}

// FuncRemoveAttribute returns the method "Element.removeAttribute".
func (this Element) FuncRemoveAttribute() (fn js.Func[func(qualifiedName js.String)]) {
	bindings.FuncElementRemoveAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAttribute calls the method "Element.removeAttribute".
func (this Element) RemoveAttribute(qualifiedName js.String) (ret js.Void) {
	bindings.CallElementRemoveAttribute(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryRemoveAttribute calls the method "Element.removeAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRemoveAttribute(qualifiedName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncRemoveAttributeNS returns true if the method "Element.removeAttributeNS" exists.
func (this Element) HasFuncRemoveAttributeNS() bool {
	return js.True == bindings.HasFuncElementRemoveAttributeNS(
		this.ref,
	)
}

// FuncRemoveAttributeNS returns the method "Element.removeAttributeNS".
func (this Element) FuncRemoveAttributeNS() (fn js.Func[func(namespace js.String, localName js.String)]) {
	bindings.FuncElementRemoveAttributeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAttributeNS calls the method "Element.removeAttributeNS".
func (this Element) RemoveAttributeNS(namespace js.String, localName js.String) (ret js.Void) {
	bindings.CallElementRemoveAttributeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryRemoveAttributeNS calls the method "Element.removeAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRemoveAttributeNS(namespace js.String, localName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttributeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncToggleAttribute returns true if the method "Element.toggleAttribute" exists.
func (this Element) HasFuncToggleAttribute() bool {
	return js.True == bindings.HasFuncElementToggleAttribute(
		this.ref,
	)
}

// FuncToggleAttribute returns the method "Element.toggleAttribute".
func (this Element) FuncToggleAttribute() (fn js.Func[func(qualifiedName js.String, force bool) bool]) {
	bindings.FuncElementToggleAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToggleAttribute calls the method "Element.toggleAttribute".
func (this Element) ToggleAttribute(qualifiedName js.String, force bool) (ret bool) {
	bindings.CallElementToggleAttribute(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// TryToggleAttribute calls the method "Element.toggleAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryToggleAttribute(qualifiedName js.String, force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementToggleAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		js.Bool(bool(force)),
	)

	return
}

// HasFuncToggleAttribute1 returns true if the method "Element.toggleAttribute" exists.
func (this Element) HasFuncToggleAttribute1() bool {
	return js.True == bindings.HasFuncElementToggleAttribute1(
		this.ref,
	)
}

// FuncToggleAttribute1 returns the method "Element.toggleAttribute".
func (this Element) FuncToggleAttribute1() (fn js.Func[func(qualifiedName js.String) bool]) {
	bindings.FuncElementToggleAttribute1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToggleAttribute1 calls the method "Element.toggleAttribute".
func (this Element) ToggleAttribute1(qualifiedName js.String) (ret bool) {
	bindings.CallElementToggleAttribute1(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryToggleAttribute1 calls the method "Element.toggleAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryToggleAttribute1(qualifiedName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementToggleAttribute1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncHasAttribute returns true if the method "Element.hasAttribute" exists.
func (this Element) HasFuncHasAttribute() bool {
	return js.True == bindings.HasFuncElementHasAttribute(
		this.ref,
	)
}

// FuncHasAttribute returns the method "Element.hasAttribute".
func (this Element) FuncHasAttribute() (fn js.Func[func(qualifiedName js.String) bool]) {
	bindings.FuncElementHasAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasAttribute calls the method "Element.hasAttribute".
func (this Element) HasAttribute(qualifiedName js.String) (ret bool) {
	bindings.CallElementHasAttribute(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryHasAttribute calls the method "Element.hasAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryHasAttribute(qualifiedName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncHasAttributeNS returns true if the method "Element.hasAttributeNS" exists.
func (this Element) HasFuncHasAttributeNS() bool {
	return js.True == bindings.HasFuncElementHasAttributeNS(
		this.ref,
	)
}

// FuncHasAttributeNS returns the method "Element.hasAttributeNS".
func (this Element) FuncHasAttributeNS() (fn js.Func[func(namespace js.String, localName js.String) bool]) {
	bindings.FuncElementHasAttributeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasAttributeNS calls the method "Element.hasAttributeNS".
func (this Element) HasAttributeNS(namespace js.String, localName js.String) (ret bool) {
	bindings.CallElementHasAttributeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryHasAttributeNS calls the method "Element.hasAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryHasAttributeNS(namespace js.String, localName js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasAttributeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncGetAttributeNode returns true if the method "Element.getAttributeNode" exists.
func (this Element) HasFuncGetAttributeNode() bool {
	return js.True == bindings.HasFuncElementGetAttributeNode(
		this.ref,
	)
}

// FuncGetAttributeNode returns the method "Element.getAttributeNode".
func (this Element) FuncGetAttributeNode() (fn js.Func[func(qualifiedName js.String) Attr]) {
	bindings.FuncElementGetAttributeNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttributeNode calls the method "Element.getAttributeNode".
func (this Element) GetAttributeNode(qualifiedName js.String) (ret Attr) {
	bindings.CallElementGetAttributeNode(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetAttributeNode calls the method "Element.getAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAttributeNode(qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncGetAttributeNodeNS returns true if the method "Element.getAttributeNodeNS" exists.
func (this Element) HasFuncGetAttributeNodeNS() bool {
	return js.True == bindings.HasFuncElementGetAttributeNodeNS(
		this.ref,
	)
}

// FuncGetAttributeNodeNS returns the method "Element.getAttributeNodeNS".
func (this Element) FuncGetAttributeNodeNS() (fn js.Func[func(namespace js.String, localName js.String) Attr]) {
	bindings.FuncElementGetAttributeNodeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttributeNodeNS calls the method "Element.getAttributeNodeNS".
func (this Element) GetAttributeNodeNS(namespace js.String, localName js.String) (ret Attr) {
	bindings.CallElementGetAttributeNodeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetAttributeNodeNS calls the method "Element.getAttributeNodeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAttributeNodeNS(namespace js.String, localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAttributeNodeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncSetAttributeNode returns true if the method "Element.setAttributeNode" exists.
func (this Element) HasFuncSetAttributeNode() bool {
	return js.True == bindings.HasFuncElementSetAttributeNode(
		this.ref,
	)
}

// FuncSetAttributeNode returns the method "Element.setAttributeNode".
func (this Element) FuncSetAttributeNode() (fn js.Func[func(attr Attr) Attr]) {
	bindings.FuncElementSetAttributeNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAttributeNode calls the method "Element.setAttributeNode".
func (this Element) SetAttributeNode(attr Attr) (ret Attr) {
	bindings.CallElementSetAttributeNode(
		this.ref, js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetAttributeNode calls the method "Element.setAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetAttributeNode(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasFuncSetAttributeNodeNS returns true if the method "Element.setAttributeNodeNS" exists.
func (this Element) HasFuncSetAttributeNodeNS() bool {
	return js.True == bindings.HasFuncElementSetAttributeNodeNS(
		this.ref,
	)
}

// FuncSetAttributeNodeNS returns the method "Element.setAttributeNodeNS".
func (this Element) FuncSetAttributeNodeNS() (fn js.Func[func(attr Attr) Attr]) {
	bindings.FuncElementSetAttributeNodeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAttributeNodeNS calls the method "Element.setAttributeNodeNS".
func (this Element) SetAttributeNodeNS(attr Attr) (ret Attr) {
	bindings.CallElementSetAttributeNodeNS(
		this.ref, js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TrySetAttributeNodeNS calls the method "Element.setAttributeNodeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetAttributeNodeNS(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetAttributeNodeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasFuncRemoveAttributeNode returns true if the method "Element.removeAttributeNode" exists.
func (this Element) HasFuncRemoveAttributeNode() bool {
	return js.True == bindings.HasFuncElementRemoveAttributeNode(
		this.ref,
	)
}

// FuncRemoveAttributeNode returns the method "Element.removeAttributeNode".
func (this Element) FuncRemoveAttributeNode() (fn js.Func[func(attr Attr) Attr]) {
	bindings.FuncElementRemoveAttributeNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAttributeNode calls the method "Element.removeAttributeNode".
func (this Element) RemoveAttributeNode(attr Attr) (ret Attr) {
	bindings.CallElementRemoveAttributeNode(
		this.ref, js.Pointer(&ret),
		attr.Ref(),
	)

	return
}

// TryRemoveAttributeNode calls the method "Element.removeAttributeNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRemoveAttributeNode(attr Attr) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemoveAttributeNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		attr.Ref(),
	)

	return
}

// HasFuncAttachShadow returns true if the method "Element.attachShadow" exists.
func (this Element) HasFuncAttachShadow() bool {
	return js.True == bindings.HasFuncElementAttachShadow(
		this.ref,
	)
}

// FuncAttachShadow returns the method "Element.attachShadow".
func (this Element) FuncAttachShadow() (fn js.Func[func(init ShadowRootInit) ShadowRoot]) {
	bindings.FuncElementAttachShadow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AttachShadow calls the method "Element.attachShadow".
func (this Element) AttachShadow(init ShadowRootInit) (ret ShadowRoot) {
	bindings.CallElementAttachShadow(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryAttachShadow calls the method "Element.attachShadow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryAttachShadow(init ShadowRootInit) (ret ShadowRoot, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAttachShadow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncClosest returns true if the method "Element.closest" exists.
func (this Element) HasFuncClosest() bool {
	return js.True == bindings.HasFuncElementClosest(
		this.ref,
	)
}

// FuncClosest returns the method "Element.closest".
func (this Element) FuncClosest() (fn js.Func[func(selectors js.String) Element]) {
	bindings.FuncElementClosest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Closest calls the method "Element.closest".
func (this Element) Closest(selectors js.String) (ret Element) {
	bindings.CallElementClosest(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryClosest calls the method "Element.closest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryClosest(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementClosest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncMatches returns true if the method "Element.matches" exists.
func (this Element) HasFuncMatches() bool {
	return js.True == bindings.HasFuncElementMatches(
		this.ref,
	)
}

// FuncMatches returns the method "Element.matches".
func (this Element) FuncMatches() (fn js.Func[func(selectors js.String) bool]) {
	bindings.FuncElementMatches(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Matches calls the method "Element.matches".
func (this Element) Matches(selectors js.String) (ret bool) {
	bindings.CallElementMatches(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryMatches calls the method "Element.matches"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryMatches(selectors js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementMatches(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncWebkitMatchesSelector returns true if the method "Element.webkitMatchesSelector" exists.
func (this Element) HasFuncWebkitMatchesSelector() bool {
	return js.True == bindings.HasFuncElementWebkitMatchesSelector(
		this.ref,
	)
}

// FuncWebkitMatchesSelector returns the method "Element.webkitMatchesSelector".
func (this Element) FuncWebkitMatchesSelector() (fn js.Func[func(selectors js.String) bool]) {
	bindings.FuncElementWebkitMatchesSelector(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WebkitMatchesSelector calls the method "Element.webkitMatchesSelector".
func (this Element) WebkitMatchesSelector(selectors js.String) (ret bool) {
	bindings.CallElementWebkitMatchesSelector(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryWebkitMatchesSelector calls the method "Element.webkitMatchesSelector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryWebkitMatchesSelector(selectors js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementWebkitMatchesSelector(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncGetElementsByTagName returns true if the method "Element.getElementsByTagName" exists.
func (this Element) HasFuncGetElementsByTagName() bool {
	return js.True == bindings.HasFuncElementGetElementsByTagName(
		this.ref,
	)
}

// FuncGetElementsByTagName returns the method "Element.getElementsByTagName".
func (this Element) FuncGetElementsByTagName() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	bindings.FuncElementGetElementsByTagName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByTagName calls the method "Element.getElementsByTagName".
func (this Element) GetElementsByTagName(qualifiedName js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByTagName(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetElementsByTagName calls the method "Element.getElementsByTagName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetElementsByTagName(qualifiedName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByTagName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncGetElementsByTagNameNS returns true if the method "Element.getElementsByTagNameNS" exists.
func (this Element) HasFuncGetElementsByTagNameNS() bool {
	return js.True == bindings.HasFuncElementGetElementsByTagNameNS(
		this.ref,
	)
}

// FuncGetElementsByTagNameNS returns the method "Element.getElementsByTagNameNS".
func (this Element) FuncGetElementsByTagNameNS() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	bindings.FuncElementGetElementsByTagNameNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByTagNameNS calls the method "Element.getElementsByTagNameNS".
func (this Element) GetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByTagNameNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetElementsByTagNameNS calls the method "Element.getElementsByTagNameNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByTagNameNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncGetElementsByClassName returns true if the method "Element.getElementsByClassName" exists.
func (this Element) HasFuncGetElementsByClassName() bool {
	return js.True == bindings.HasFuncElementGetElementsByClassName(
		this.ref,
	)
}

// FuncGetElementsByClassName returns the method "Element.getElementsByClassName".
func (this Element) FuncGetElementsByClassName() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	bindings.FuncElementGetElementsByClassName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByClassName calls the method "Element.getElementsByClassName".
func (this Element) GetElementsByClassName(classNames js.String) (ret HTMLCollection) {
	bindings.CallElementGetElementsByClassName(
		this.ref, js.Pointer(&ret),
		classNames.Ref(),
	)

	return
}

// TryGetElementsByClassName calls the method "Element.getElementsByClassName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetElementsByClassName(classNames js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetElementsByClassName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		classNames.Ref(),
	)

	return
}

// HasFuncInsertAdjacentElement returns true if the method "Element.insertAdjacentElement" exists.
func (this Element) HasFuncInsertAdjacentElement() bool {
	return js.True == bindings.HasFuncElementInsertAdjacentElement(
		this.ref,
	)
}

// FuncInsertAdjacentElement returns the method "Element.insertAdjacentElement".
func (this Element) FuncInsertAdjacentElement() (fn js.Func[func(where js.String, element Element) Element]) {
	bindings.FuncElementInsertAdjacentElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertAdjacentElement calls the method "Element.insertAdjacentElement".
func (this Element) InsertAdjacentElement(where js.String, element Element) (ret Element) {
	bindings.CallElementInsertAdjacentElement(
		this.ref, js.Pointer(&ret),
		where.Ref(),
		element.Ref(),
	)

	return
}

// TryInsertAdjacentElement calls the method "Element.insertAdjacentElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryInsertAdjacentElement(where js.String, element Element) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		where.Ref(),
		element.Ref(),
	)

	return
}

// HasFuncInsertAdjacentText returns true if the method "Element.insertAdjacentText" exists.
func (this Element) HasFuncInsertAdjacentText() bool {
	return js.True == bindings.HasFuncElementInsertAdjacentText(
		this.ref,
	)
}

// FuncInsertAdjacentText returns the method "Element.insertAdjacentText".
func (this Element) FuncInsertAdjacentText() (fn js.Func[func(where js.String, data js.String)]) {
	bindings.FuncElementInsertAdjacentText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertAdjacentText calls the method "Element.insertAdjacentText".
func (this Element) InsertAdjacentText(where js.String, data js.String) (ret js.Void) {
	bindings.CallElementInsertAdjacentText(
		this.ref, js.Pointer(&ret),
		where.Ref(),
		data.Ref(),
	)

	return
}

// TryInsertAdjacentText calls the method "Element.insertAdjacentText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryInsertAdjacentText(where js.String, data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		where.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncRequestFullscreen returns true if the method "Element.requestFullscreen" exists.
func (this Element) HasFuncRequestFullscreen() bool {
	return js.True == bindings.HasFuncElementRequestFullscreen(
		this.ref,
	)
}

// FuncRequestFullscreen returns the method "Element.requestFullscreen".
func (this Element) FuncRequestFullscreen() (fn js.Func[func(options FullscreenOptions) js.Promise[js.Void]]) {
	bindings.FuncElementRequestFullscreen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestFullscreen calls the method "Element.requestFullscreen".
func (this Element) RequestFullscreen(options FullscreenOptions) (ret js.Promise[js.Void]) {
	bindings.CallElementRequestFullscreen(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestFullscreen calls the method "Element.requestFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRequestFullscreen(options FullscreenOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestFullscreen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestFullscreen1 returns true if the method "Element.requestFullscreen" exists.
func (this Element) HasFuncRequestFullscreen1() bool {
	return js.True == bindings.HasFuncElementRequestFullscreen1(
		this.ref,
	)
}

// FuncRequestFullscreen1 returns the method "Element.requestFullscreen".
func (this Element) FuncRequestFullscreen1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncElementRequestFullscreen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestFullscreen1 calls the method "Element.requestFullscreen".
func (this Element) RequestFullscreen1() (ret js.Promise[js.Void]) {
	bindings.CallElementRequestFullscreen1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestFullscreen1 calls the method "Element.requestFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRequestFullscreen1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestFullscreen1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestPointerLock returns true if the method "Element.requestPointerLock" exists.
func (this Element) HasFuncRequestPointerLock() bool {
	return js.True == bindings.HasFuncElementRequestPointerLock(
		this.ref,
	)
}

// FuncRequestPointerLock returns the method "Element.requestPointerLock".
func (this Element) FuncRequestPointerLock() (fn js.Func[func()]) {
	bindings.FuncElementRequestPointerLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPointerLock calls the method "Element.requestPointerLock".
func (this Element) RequestPointerLock() (ret js.Void) {
	bindings.CallElementRequestPointerLock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPointerLock calls the method "Element.requestPointerLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRequestPointerLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRequestPointerLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncComputedStyleMap returns true if the method "Element.computedStyleMap" exists.
func (this Element) HasFuncComputedStyleMap() bool {
	return js.True == bindings.HasFuncElementComputedStyleMap(
		this.ref,
	)
}

// FuncComputedStyleMap returns the method "Element.computedStyleMap".
func (this Element) FuncComputedStyleMap() (fn js.Func[func() StylePropertyMapReadOnly]) {
	bindings.FuncElementComputedStyleMap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ComputedStyleMap calls the method "Element.computedStyleMap".
func (this Element) ComputedStyleMap() (ret StylePropertyMapReadOnly) {
	bindings.CallElementComputedStyleMap(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryComputedStyleMap calls the method "Element.computedStyleMap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryComputedStyleMap() (ret StylePropertyMapReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementComputedStyleMap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSpatialNavigationContainer returns true if the method "Element.getSpatialNavigationContainer" exists.
func (this Element) HasFuncGetSpatialNavigationContainer() bool {
	return js.True == bindings.HasFuncElementGetSpatialNavigationContainer(
		this.ref,
	)
}

// FuncGetSpatialNavigationContainer returns the method "Element.getSpatialNavigationContainer".
func (this Element) FuncGetSpatialNavigationContainer() (fn js.Func[func() Node]) {
	bindings.FuncElementGetSpatialNavigationContainer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSpatialNavigationContainer calls the method "Element.getSpatialNavigationContainer".
func (this Element) GetSpatialNavigationContainer() (ret Node) {
	bindings.CallElementGetSpatialNavigationContainer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSpatialNavigationContainer calls the method "Element.getSpatialNavigationContainer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetSpatialNavigationContainer() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetSpatialNavigationContainer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFocusableAreas returns true if the method "Element.focusableAreas" exists.
func (this Element) HasFuncFocusableAreas() bool {
	return js.True == bindings.HasFuncElementFocusableAreas(
		this.ref,
	)
}

// FuncFocusableAreas returns the method "Element.focusableAreas".
func (this Element) FuncFocusableAreas() (fn js.Func[func(option FocusableAreasOption) js.Array[Node]]) {
	bindings.FuncElementFocusableAreas(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FocusableAreas calls the method "Element.focusableAreas".
func (this Element) FocusableAreas(option FocusableAreasOption) (ret js.Array[Node]) {
	bindings.CallElementFocusableAreas(
		this.ref, js.Pointer(&ret),
		js.Pointer(&option),
	)

	return
}

// TryFocusableAreas calls the method "Element.focusableAreas"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryFocusableAreas(option FocusableAreasOption) (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementFocusableAreas(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&option),
	)

	return
}

// HasFuncFocusableAreas1 returns true if the method "Element.focusableAreas" exists.
func (this Element) HasFuncFocusableAreas1() bool {
	return js.True == bindings.HasFuncElementFocusableAreas1(
		this.ref,
	)
}

// FuncFocusableAreas1 returns the method "Element.focusableAreas".
func (this Element) FuncFocusableAreas1() (fn js.Func[func() js.Array[Node]]) {
	bindings.FuncElementFocusableAreas1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FocusableAreas1 calls the method "Element.focusableAreas".
func (this Element) FocusableAreas1() (ret js.Array[Node]) {
	bindings.CallElementFocusableAreas1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocusableAreas1 calls the method "Element.focusableAreas"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryFocusableAreas1() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementFocusableAreas1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSpatialNavigationSearch returns true if the method "Element.spatialNavigationSearch" exists.
func (this Element) HasFuncSpatialNavigationSearch() bool {
	return js.True == bindings.HasFuncElementSpatialNavigationSearch(
		this.ref,
	)
}

// FuncSpatialNavigationSearch returns the method "Element.spatialNavigationSearch".
func (this Element) FuncSpatialNavigationSearch() (fn js.Func[func(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) Node]) {
	bindings.FuncElementSpatialNavigationSearch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SpatialNavigationSearch calls the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearch(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) (ret Node) {
	bindings.CallElementSpatialNavigationSearch(
		this.ref, js.Pointer(&ret),
		uint32(dir),
		js.Pointer(&options),
	)

	return
}

// TrySpatialNavigationSearch calls the method "Element.spatialNavigationSearch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySpatialNavigationSearch(dir SpatialNavigationDirection, options SpatialNavigationSearchOptions) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSpatialNavigationSearch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
		js.Pointer(&options),
	)

	return
}

// HasFuncSpatialNavigationSearch1 returns true if the method "Element.spatialNavigationSearch" exists.
func (this Element) HasFuncSpatialNavigationSearch1() bool {
	return js.True == bindings.HasFuncElementSpatialNavigationSearch1(
		this.ref,
	)
}

// FuncSpatialNavigationSearch1 returns the method "Element.spatialNavigationSearch".
func (this Element) FuncSpatialNavigationSearch1() (fn js.Func[func(dir SpatialNavigationDirection) Node]) {
	bindings.FuncElementSpatialNavigationSearch1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SpatialNavigationSearch1 calls the method "Element.spatialNavigationSearch".
func (this Element) SpatialNavigationSearch1(dir SpatialNavigationDirection) (ret Node) {
	bindings.CallElementSpatialNavigationSearch1(
		this.ref, js.Pointer(&ret),
		uint32(dir),
	)

	return
}

// TrySpatialNavigationSearch1 calls the method "Element.spatialNavigationSearch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySpatialNavigationSearch1(dir SpatialNavigationDirection) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSpatialNavigationSearch1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
	)

	return
}

// HasFuncSetPointerCapture returns true if the method "Element.setPointerCapture" exists.
func (this Element) HasFuncSetPointerCapture() bool {
	return js.True == bindings.HasFuncElementSetPointerCapture(
		this.ref,
	)
}

// FuncSetPointerCapture returns the method "Element.setPointerCapture".
func (this Element) FuncSetPointerCapture() (fn js.Func[func(pointerId int32)]) {
	bindings.FuncElementSetPointerCapture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPointerCapture calls the method "Element.setPointerCapture".
func (this Element) SetPointerCapture(pointerId int32) (ret js.Void) {
	bindings.CallElementSetPointerCapture(
		this.ref, js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TrySetPointerCapture calls the method "Element.setPointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetPointerCapture(pointerId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetPointerCapture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasFuncReleasePointerCapture returns true if the method "Element.releasePointerCapture" exists.
func (this Element) HasFuncReleasePointerCapture() bool {
	return js.True == bindings.HasFuncElementReleasePointerCapture(
		this.ref,
	)
}

// FuncReleasePointerCapture returns the method "Element.releasePointerCapture".
func (this Element) FuncReleasePointerCapture() (fn js.Func[func(pointerId int32)]) {
	bindings.FuncElementReleasePointerCapture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleasePointerCapture calls the method "Element.releasePointerCapture".
func (this Element) ReleasePointerCapture(pointerId int32) (ret js.Void) {
	bindings.CallElementReleasePointerCapture(
		this.ref, js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TryReleasePointerCapture calls the method "Element.releasePointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryReleasePointerCapture(pointerId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReleasePointerCapture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasFuncHasPointerCapture returns true if the method "Element.hasPointerCapture" exists.
func (this Element) HasFuncHasPointerCapture() bool {
	return js.True == bindings.HasFuncElementHasPointerCapture(
		this.ref,
	)
}

// FuncHasPointerCapture returns the method "Element.hasPointerCapture".
func (this Element) FuncHasPointerCapture() (fn js.Func[func(pointerId int32) bool]) {
	bindings.FuncElementHasPointerCapture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasPointerCapture calls the method "Element.hasPointerCapture".
func (this Element) HasPointerCapture(pointerId int32) (ret bool) {
	bindings.CallElementHasPointerCapture(
		this.ref, js.Pointer(&ret),
		int32(pointerId),
	)

	return
}

// TryHasPointerCapture calls the method "Element.hasPointerCapture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryHasPointerCapture(pointerId int32) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementHasPointerCapture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(pointerId),
	)

	return
}

// HasFuncPseudo returns true if the method "Element.pseudo" exists.
func (this Element) HasFuncPseudo() bool {
	return js.True == bindings.HasFuncElementPseudo(
		this.ref,
	)
}

// FuncPseudo returns the method "Element.pseudo".
func (this Element) FuncPseudo() (fn js.Func[func(typ js.String) CSSPseudoElement]) {
	bindings.FuncElementPseudo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pseudo calls the method "Element.pseudo".
func (this Element) Pseudo(typ js.String) (ret CSSPseudoElement) {
	bindings.CallElementPseudo(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryPseudo calls the method "Element.pseudo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryPseudo(typ js.String) (ret CSSPseudoElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementPseudo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncInsertAdjacentHTML returns true if the method "Element.insertAdjacentHTML" exists.
func (this Element) HasFuncInsertAdjacentHTML() bool {
	return js.True == bindings.HasFuncElementInsertAdjacentHTML(
		this.ref,
	)
}

// FuncInsertAdjacentHTML returns the method "Element.insertAdjacentHTML".
func (this Element) FuncInsertAdjacentHTML() (fn js.Func[func(position js.String, text js.String)]) {
	bindings.FuncElementInsertAdjacentHTML(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertAdjacentHTML calls the method "Element.insertAdjacentHTML".
func (this Element) InsertAdjacentHTML(position js.String, text js.String) (ret js.Void) {
	bindings.CallElementInsertAdjacentHTML(
		this.ref, js.Pointer(&ret),
		position.Ref(),
		text.Ref(),
	)

	return
}

// TryInsertAdjacentHTML calls the method "Element.insertAdjacentHTML"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryInsertAdjacentHTML(position js.String, text js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInsertAdjacentHTML(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		position.Ref(),
		text.Ref(),
	)

	return
}

// HasFuncSetHTML returns true if the method "Element.setHTML" exists.
func (this Element) HasFuncSetHTML() bool {
	return js.True == bindings.HasFuncElementSetHTML(
		this.ref,
	)
}

// FuncSetHTML returns the method "Element.setHTML".
func (this Element) FuncSetHTML() (fn js.Func[func(input js.String, options SetHTMLOptions)]) {
	bindings.FuncElementSetHTML(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetHTML calls the method "Element.setHTML".
func (this Element) SetHTML(input js.String, options SetHTMLOptions) (ret js.Void) {
	bindings.CallElementSetHTML(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySetHTML calls the method "Element.setHTML"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetHTML(input js.String, options SetHTMLOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetHTML(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSetHTML1 returns true if the method "Element.setHTML" exists.
func (this Element) HasFuncSetHTML1() bool {
	return js.True == bindings.HasFuncElementSetHTML1(
		this.ref,
	)
}

// FuncSetHTML1 returns the method "Element.setHTML".
func (this Element) FuncSetHTML1() (fn js.Func[func(input js.String)]) {
	bindings.FuncElementSetHTML1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetHTML1 calls the method "Element.setHTML".
func (this Element) SetHTML1(input js.String) (ret js.Void) {
	bindings.CallElementSetHTML1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySetHTML1 calls the method "Element.setHTML"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TrySetHTML1(input js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementSetHTML1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncGetClientRects returns true if the method "Element.getClientRects" exists.
func (this Element) HasFuncGetClientRects() bool {
	return js.True == bindings.HasFuncElementGetClientRects(
		this.ref,
	)
}

// FuncGetClientRects returns the method "Element.getClientRects".
func (this Element) FuncGetClientRects() (fn js.Func[func() DOMRectList]) {
	bindings.FuncElementGetClientRects(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetClientRects calls the method "Element.getClientRects".
func (this Element) GetClientRects() (ret DOMRectList) {
	bindings.CallElementGetClientRects(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetClientRects calls the method "Element.getClientRects"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetClientRects() (ret DOMRectList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetClientRects(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBoundingClientRect returns true if the method "Element.getBoundingClientRect" exists.
func (this Element) HasFuncGetBoundingClientRect() bool {
	return js.True == bindings.HasFuncElementGetBoundingClientRect(
		this.ref,
	)
}

// FuncGetBoundingClientRect returns the method "Element.getBoundingClientRect".
func (this Element) FuncGetBoundingClientRect() (fn js.Func[func() DOMRect]) {
	bindings.FuncElementGetBoundingClientRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoundingClientRect calls the method "Element.getBoundingClientRect".
func (this Element) GetBoundingClientRect() (ret DOMRect) {
	bindings.CallElementGetBoundingClientRect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoundingClientRect calls the method "Element.getBoundingClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetBoundingClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoundingClientRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckVisibility returns true if the method "Element.checkVisibility" exists.
func (this Element) HasFuncCheckVisibility() bool {
	return js.True == bindings.HasFuncElementCheckVisibility(
		this.ref,
	)
}

// FuncCheckVisibility returns the method "Element.checkVisibility".
func (this Element) FuncCheckVisibility() (fn js.Func[func(options CheckVisibilityOptions) bool]) {
	bindings.FuncElementCheckVisibility(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckVisibility calls the method "Element.checkVisibility".
func (this Element) CheckVisibility(options CheckVisibilityOptions) (ret bool) {
	bindings.CallElementCheckVisibility(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCheckVisibility calls the method "Element.checkVisibility"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryCheckVisibility(options CheckVisibilityOptions) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementCheckVisibility(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCheckVisibility1 returns true if the method "Element.checkVisibility" exists.
func (this Element) HasFuncCheckVisibility1() bool {
	return js.True == bindings.HasFuncElementCheckVisibility1(
		this.ref,
	)
}

// FuncCheckVisibility1 returns the method "Element.checkVisibility".
func (this Element) FuncCheckVisibility1() (fn js.Func[func() bool]) {
	bindings.FuncElementCheckVisibility1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckVisibility1 calls the method "Element.checkVisibility".
func (this Element) CheckVisibility1() (ret bool) {
	bindings.CallElementCheckVisibility1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckVisibility1 calls the method "Element.checkVisibility"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryCheckVisibility1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementCheckVisibility1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollIntoView returns true if the method "Element.scrollIntoView" exists.
func (this Element) HasFuncScrollIntoView() bool {
	return js.True == bindings.HasFuncElementScrollIntoView(
		this.ref,
	)
}

// FuncScrollIntoView returns the method "Element.scrollIntoView".
func (this Element) FuncScrollIntoView() (fn js.Func[func(arg OneOf_Bool_ScrollIntoViewOptions)]) {
	bindings.FuncElementScrollIntoView(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollIntoView calls the method "Element.scrollIntoView".
func (this Element) ScrollIntoView(arg OneOf_Bool_ScrollIntoViewOptions) (ret js.Void) {
	bindings.CallElementScrollIntoView(
		this.ref, js.Pointer(&ret),
		arg.Ref(),
	)

	return
}

// TryScrollIntoView calls the method "Element.scrollIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollIntoView(arg OneOf_Bool_ScrollIntoViewOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollIntoView(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		arg.Ref(),
	)

	return
}

// HasFuncScrollIntoView1 returns true if the method "Element.scrollIntoView" exists.
func (this Element) HasFuncScrollIntoView1() bool {
	return js.True == bindings.HasFuncElementScrollIntoView1(
		this.ref,
	)
}

// FuncScrollIntoView1 returns the method "Element.scrollIntoView".
func (this Element) FuncScrollIntoView1() (fn js.Func[func()]) {
	bindings.FuncElementScrollIntoView1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollIntoView1 calls the method "Element.scrollIntoView".
func (this Element) ScrollIntoView1() (ret js.Void) {
	bindings.CallElementScrollIntoView1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollIntoView1 calls the method "Element.scrollIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollIntoView1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollIntoView1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScroll returns true if the method "Element.scroll" exists.
func (this Element) HasFuncScroll() bool {
	return js.True == bindings.HasFuncElementScroll(
		this.ref,
	)
}

// FuncScroll returns the method "Element.scroll".
func (this Element) FuncScroll() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncElementScroll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll calls the method "Element.scroll".
func (this Element) Scroll(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScroll(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScroll calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScroll(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScroll1 returns true if the method "Element.scroll" exists.
func (this Element) HasFuncScroll1() bool {
	return js.True == bindings.HasFuncElementScroll1(
		this.ref,
	)
}

// FuncScroll1 returns the method "Element.scroll".
func (this Element) FuncScroll1() (fn js.Func[func()]) {
	bindings.FuncElementScroll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll1 calls the method "Element.scroll".
func (this Element) Scroll1() (ret js.Void) {
	bindings.CallElementScroll1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScroll1 calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScroll1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScroll2 returns true if the method "Element.scroll" exists.
func (this Element) HasFuncScroll2() bool {
	return js.True == bindings.HasFuncElementScroll2(
		this.ref,
	)
}

// FuncScroll2 returns the method "Element.scroll".
func (this Element) FuncScroll2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncElementScroll2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll2 calls the method "Element.scroll".
func (this Element) Scroll2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScroll2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScroll2 calls the method "Element.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScroll2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScroll2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncScrollTo returns true if the method "Element.scrollTo" exists.
func (this Element) HasFuncScrollTo() bool {
	return js.True == bindings.HasFuncElementScrollTo(
		this.ref,
	)
}

// FuncScrollTo returns the method "Element.scrollTo".
func (this Element) FuncScrollTo() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncElementScrollTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo calls the method "Element.scrollTo".
func (this Element) ScrollTo(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScrollTo(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollTo calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollTo(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScrollTo1 returns true if the method "Element.scrollTo" exists.
func (this Element) HasFuncScrollTo1() bool {
	return js.True == bindings.HasFuncElementScrollTo1(
		this.ref,
	)
}

// FuncScrollTo1 returns the method "Element.scrollTo".
func (this Element) FuncScrollTo1() (fn js.Func[func()]) {
	bindings.FuncElementScrollTo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo1 calls the method "Element.scrollTo".
func (this Element) ScrollTo1() (ret js.Void) {
	bindings.CallElementScrollTo1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollTo1 calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollTo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollTo2 returns true if the method "Element.scrollTo" exists.
func (this Element) HasFuncScrollTo2() bool {
	return js.True == bindings.HasFuncElementScrollTo2(
		this.ref,
	)
}

// FuncScrollTo2 returns the method "Element.scrollTo".
func (this Element) FuncScrollTo2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncElementScrollTo2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo2 calls the method "Element.scrollTo".
func (this Element) ScrollTo2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScrollTo2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollTo2 calls the method "Element.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollTo2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollTo2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncScrollBy returns true if the method "Element.scrollBy" exists.
func (this Element) HasFuncScrollBy() bool {
	return js.True == bindings.HasFuncElementScrollBy(
		this.ref,
	)
}

// FuncScrollBy returns the method "Element.scrollBy".
func (this Element) FuncScrollBy() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncElementScrollBy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy calls the method "Element.scrollBy".
func (this Element) ScrollBy(options ScrollToOptions) (ret js.Void) {
	bindings.CallElementScrollBy(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollBy calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollBy(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScrollBy1 returns true if the method "Element.scrollBy" exists.
func (this Element) HasFuncScrollBy1() bool {
	return js.True == bindings.HasFuncElementScrollBy1(
		this.ref,
	)
}

// FuncScrollBy1 returns the method "Element.scrollBy".
func (this Element) FuncScrollBy1() (fn js.Func[func()]) {
	bindings.FuncElementScrollBy1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy1 calls the method "Element.scrollBy".
func (this Element) ScrollBy1() (ret js.Void) {
	bindings.CallElementScrollBy1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollBy1 calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollBy1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollBy2 returns true if the method "Element.scrollBy" exists.
func (this Element) HasFuncScrollBy2() bool {
	return js.True == bindings.HasFuncElementScrollBy2(
		this.ref,
	)
}

// FuncScrollBy2 returns the method "Element.scrollBy".
func (this Element) FuncScrollBy2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncElementScrollBy2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy2 calls the method "Element.scrollBy".
func (this Element) ScrollBy2(x float64, y float64) (ret js.Void) {
	bindings.CallElementScrollBy2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollBy2 calls the method "Element.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryScrollBy2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementScrollBy2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncGetBoxQuads returns true if the method "Element.getBoxQuads" exists.
func (this Element) HasFuncGetBoxQuads() bool {
	return js.True == bindings.HasFuncElementGetBoxQuads(
		this.ref,
	)
}

// FuncGetBoxQuads returns the method "Element.getBoxQuads".
func (this Element) FuncGetBoxQuads() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	bindings.FuncElementGetBoxQuads(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads calls the method "Element.getBoxQuads".
func (this Element) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallElementGetBoxQuads(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Element.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoxQuads(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetBoxQuads1 returns true if the method "Element.getBoxQuads" exists.
func (this Element) HasFuncGetBoxQuads1() bool {
	return js.True == bindings.HasFuncElementGetBoxQuads1(
		this.ref,
	)
}

// FuncGetBoxQuads1 returns the method "Element.getBoxQuads".
func (this Element) FuncGetBoxQuads1() (fn js.Func[func() js.Array[DOMQuad]]) {
	bindings.FuncElementGetBoxQuads1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads1 calls the method "Element.getBoxQuads".
func (this Element) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallElementGetBoxQuads1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Element.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetBoxQuads1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvertQuadFromNode returns true if the method "Element.convertQuadFromNode" exists.
func (this Element) HasFuncConvertQuadFromNode() bool {
	return js.True == bindings.HasFuncElementConvertQuadFromNode(
		this.ref,
	)
}

// FuncConvertQuadFromNode returns the method "Element.convertQuadFromNode".
func (this Element) FuncConvertQuadFromNode() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncElementConvertQuadFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode calls the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallElementConvertQuadFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "Element.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertQuadFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertQuadFromNode1 returns true if the method "Element.convertQuadFromNode" exists.
func (this Element) HasFuncConvertQuadFromNode1() bool {
	return js.True == bindings.HasFuncElementConvertQuadFromNode1(
		this.ref,
	)
}

// FuncConvertQuadFromNode1 returns the method "Element.convertQuadFromNode".
func (this Element) FuncConvertQuadFromNode1() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	bindings.FuncElementConvertQuadFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode1 calls the method "Element.convertQuadFromNode".
func (this Element) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallElementConvertQuadFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "Element.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertQuadFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasFuncConvertRectFromNode returns true if the method "Element.convertRectFromNode" exists.
func (this Element) HasFuncConvertRectFromNode() bool {
	return js.True == bindings.HasFuncElementConvertRectFromNode(
		this.ref,
	)
}

// FuncConvertRectFromNode returns the method "Element.convertRectFromNode".
func (this Element) FuncConvertRectFromNode() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncElementConvertRectFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode calls the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallElementConvertRectFromNode(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "Element.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertRectFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertRectFromNode1 returns true if the method "Element.convertRectFromNode" exists.
func (this Element) HasFuncConvertRectFromNode1() bool {
	return js.True == bindings.HasFuncElementConvertRectFromNode1(
		this.ref,
	)
}

// FuncConvertRectFromNode1 returns the method "Element.convertRectFromNode".
func (this Element) FuncConvertRectFromNode1() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	bindings.FuncElementConvertRectFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode1 calls the method "Element.convertRectFromNode".
func (this Element) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallElementConvertRectFromNode1(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "Element.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertRectFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasFuncConvertPointFromNode returns true if the method "Element.convertPointFromNode" exists.
func (this Element) HasFuncConvertPointFromNode() bool {
	return js.True == bindings.HasFuncElementConvertPointFromNode(
		this.ref,
	)
}

// FuncConvertPointFromNode returns the method "Element.convertPointFromNode".
func (this Element) FuncConvertPointFromNode() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	bindings.FuncElementConvertPointFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode calls the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallElementConvertPointFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "Element.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertPointFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertPointFromNode1 returns true if the method "Element.convertPointFromNode" exists.
func (this Element) HasFuncConvertPointFromNode1() bool {
	return js.True == bindings.HasFuncElementConvertPointFromNode1(
		this.ref,
	)
}

// FuncConvertPointFromNode1 returns the method "Element.convertPointFromNode".
func (this Element) FuncConvertPointFromNode1() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	bindings.FuncElementConvertPointFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode1 calls the method "Element.convertPointFromNode".
func (this Element) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallElementConvertPointFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "Element.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementConvertPointFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// HasFuncPrepend returns true if the method "Element.prepend" exists.
func (this Element) HasFuncPrepend() bool {
	return js.True == bindings.HasFuncElementPrepend(
		this.ref,
	)
}

// FuncPrepend returns the method "Element.prepend".
func (this Element) FuncPrepend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementPrepend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prepend calls the method "Element.prepend".
func (this Element) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementPrepend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryPrepend calls the method "Element.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryPrepend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementPrepend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncAppend returns true if the method "Element.append" exists.
func (this Element) HasFuncAppend() bool {
	return js.True == bindings.HasFuncElementAppend(
		this.ref,
	)
}

// FuncAppend returns the method "Element.append".
func (this Element) FuncAppend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "Element.append".
func (this Element) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementAppend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAppend calls the method "Element.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryAppend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncReplaceChildren returns true if the method "Element.replaceChildren" exists.
func (this Element) HasFuncReplaceChildren() bool {
	return js.True == bindings.HasFuncElementReplaceChildren(
		this.ref,
	)
}

// FuncReplaceChildren returns the method "Element.replaceChildren".
func (this Element) FuncReplaceChildren() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementReplaceChildren(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceChildren calls the method "Element.replaceChildren".
func (this Element) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementReplaceChildren(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceChildren calls the method "Element.replaceChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReplaceChildren(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncQuerySelector returns true if the method "Element.querySelector" exists.
func (this Element) HasFuncQuerySelector() bool {
	return js.True == bindings.HasFuncElementQuerySelector(
		this.ref,
	)
}

// FuncQuerySelector returns the method "Element.querySelector".
func (this Element) FuncQuerySelector() (fn js.Func[func(selectors js.String) Element]) {
	bindings.FuncElementQuerySelector(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelector calls the method "Element.querySelector".
func (this Element) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallElementQuerySelector(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "Element.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementQuerySelector(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncQuerySelectorAll returns true if the method "Element.querySelectorAll" exists.
func (this Element) HasFuncQuerySelectorAll() bool {
	return js.True == bindings.HasFuncElementQuerySelectorAll(
		this.ref,
	)
}

// FuncQuerySelectorAll returns the method "Element.querySelectorAll".
func (this Element) FuncQuerySelectorAll() (fn js.Func[func(selectors js.String) NodeList]) {
	bindings.FuncElementQuerySelectorAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelectorAll calls the method "Element.querySelectorAll".
func (this Element) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallElementQuerySelectorAll(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "Element.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementQuerySelectorAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncBefore returns true if the method "Element.before" exists.
func (this Element) HasFuncBefore() bool {
	return js.True == bindings.HasFuncElementBefore(
		this.ref,
	)
}

// FuncBefore returns the method "Element.before".
func (this Element) FuncBefore() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Before calls the method "Element.before".
func (this Element) Before(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementBefore(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryBefore calls the method "Element.before"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryBefore(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncAfter returns true if the method "Element.after" exists.
func (this Element) HasFuncAfter() bool {
	return js.True == bindings.HasFuncElementAfter(
		this.ref,
	)
}

// FuncAfter returns the method "Element.after".
func (this Element) FuncAfter() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementAfter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// After calls the method "Element.after".
func (this Element) After(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementAfter(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAfter calls the method "Element.after"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryAfter(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAfter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncReplaceWith returns true if the method "Element.replaceWith" exists.
func (this Element) HasFuncReplaceWith() bool {
	return js.True == bindings.HasFuncElementReplaceWith(
		this.ref,
	)
}

// FuncReplaceWith returns the method "Element.replaceWith".
func (this Element) FuncReplaceWith() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncElementReplaceWith(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceWith calls the method "Element.replaceWith".
func (this Element) ReplaceWith(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallElementReplaceWith(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceWith calls the method "Element.replaceWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryReplaceWith(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementReplaceWith(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncRemove returns true if the method "Element.remove" exists.
func (this Element) HasFuncRemove() bool {
	return js.True == bindings.HasFuncElementRemove(
		this.ref,
	)
}

// FuncRemove returns the method "Element.remove".
func (this Element) FuncRemove() (fn js.Func[func()]) {
	bindings.FuncElementRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "Element.remove".
func (this Element) Remove() (ret js.Void) {
	bindings.CallElementRemove(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "Element.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAnimate returns true if the method "Element.animate" exists.
func (this Element) HasFuncAnimate() bool {
	return js.True == bindings.HasFuncElementAnimate(
		this.ref,
	)
}

// FuncAnimate returns the method "Element.animate".
func (this Element) FuncAnimate() (fn js.Func[func(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) Animation]) {
	bindings.FuncElementAnimate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Animate calls the method "Element.animate".
func (this Element) Animate(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) (ret Animation) {
	bindings.CallElementAnimate(
		this.ref, js.Pointer(&ret),
		keyframes.Ref(),
		options.Ref(),
	)

	return
}

// TryAnimate calls the method "Element.animate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryAnimate(keyframes js.Object, options OneOf_Float64_KeyframeAnimationOptions) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAnimate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keyframes.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncAnimate1 returns true if the method "Element.animate" exists.
func (this Element) HasFuncAnimate1() bool {
	return js.True == bindings.HasFuncElementAnimate1(
		this.ref,
	)
}

// FuncAnimate1 returns the method "Element.animate".
func (this Element) FuncAnimate1() (fn js.Func[func(keyframes js.Object) Animation]) {
	bindings.FuncElementAnimate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Animate1 calls the method "Element.animate".
func (this Element) Animate1(keyframes js.Object) (ret Animation) {
	bindings.CallElementAnimate1(
		this.ref, js.Pointer(&ret),
		keyframes.Ref(),
	)

	return
}

// TryAnimate1 calls the method "Element.animate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryAnimate1(keyframes js.Object) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementAnimate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keyframes.Ref(),
	)

	return
}

// HasFuncGetAnimations returns true if the method "Element.getAnimations" exists.
func (this Element) HasFuncGetAnimations() bool {
	return js.True == bindings.HasFuncElementGetAnimations(
		this.ref,
	)
}

// FuncGetAnimations returns the method "Element.getAnimations".
func (this Element) FuncGetAnimations() (fn js.Func[func(options GetAnimationsOptions) js.Array[Animation]]) {
	bindings.FuncElementGetAnimations(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAnimations calls the method "Element.getAnimations".
func (this Element) GetAnimations(options GetAnimationsOptions) (ret js.Array[Animation]) {
	bindings.CallElementGetAnimations(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetAnimations calls the method "Element.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAnimations(options GetAnimationsOptions) (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAnimations(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetAnimations1 returns true if the method "Element.getAnimations" exists.
func (this Element) HasFuncGetAnimations1() bool {
	return js.True == bindings.HasFuncElementGetAnimations1(
		this.ref,
	)
}

// FuncGetAnimations1 returns the method "Element.getAnimations".
func (this Element) FuncGetAnimations1() (fn js.Func[func() js.Array[Animation]]) {
	bindings.FuncElementGetAnimations1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAnimations1 calls the method "Element.getAnimations".
func (this Element) GetAnimations1() (ret js.Array[Animation]) {
	bindings.CallElementGetAnimations1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAnimations1 calls the method "Element.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetAnimations1() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetAnimations1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRegionFlowRanges returns true if the method "Element.getRegionFlowRanges" exists.
func (this Element) HasFuncGetRegionFlowRanges() bool {
	return js.True == bindings.HasFuncElementGetRegionFlowRanges(
		this.ref,
	)
}

// FuncGetRegionFlowRanges returns the method "Element.getRegionFlowRanges".
func (this Element) FuncGetRegionFlowRanges() (fn js.Func[func() js.Array[Range]]) {
	bindings.FuncElementGetRegionFlowRanges(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRegionFlowRanges calls the method "Element.getRegionFlowRanges".
func (this Element) GetRegionFlowRanges() (ret js.Array[Range]) {
	bindings.CallElementGetRegionFlowRanges(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRegionFlowRanges calls the method "Element.getRegionFlowRanges"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Element) TryGetRegionFlowRanges() (ret js.Array[Range], exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementGetRegionFlowRanges(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLCollection struct {
	ref js.Ref
}

func (this HTMLCollection) Once() HTMLCollection {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "HTMLCollection.length".
//
// It returns ok=false if there is no such property.
func (this HTMLCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCollectionLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "HTMLCollection.item" exists.
func (this HTMLCollection) HasFuncItem() bool {
	return js.True == bindings.HasFuncHTMLCollectionItem(
		this.ref,
	)
}

// FuncItem returns the method "HTMLCollection.item".
func (this HTMLCollection) FuncItem() (fn js.Func[func(index uint32) Element]) {
	bindings.FuncHTMLCollectionItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "HTMLCollection.item".
func (this HTMLCollection) Item(index uint32) (ret Element) {
	bindings.CallHTMLCollectionItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "HTMLCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCollection) TryItem(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCollectionItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "HTMLCollection.namedItem" exists.
func (this HTMLCollection) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncHTMLCollectionNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "HTMLCollection.namedItem".
func (this HTMLCollection) FuncNamedItem() (fn js.Func[func(name js.String) Element]) {
	bindings.FuncHTMLCollectionNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "HTMLCollection.namedItem".
func (this HTMLCollection) NamedItem(name js.String) (ret Element) {
	bindings.CallHTMLCollectionNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCollection) TryNamedItem(name js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCollectionNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ElementCreationOptions) UpdateFrom(ref js.Ref) {
	bindings.ElementCreationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ElementCreationOptions) Update(ref js.Ref) {
	bindings.ElementCreationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ElementCreationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Is.Ref(),
	)
	p.Is = p.Is.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
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
	this.ref.Once()
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
	this.ref.Free()
}

type ProcessingInstruction struct {
	CharacterData
}

func (this ProcessingInstruction) Once() ProcessingInstruction {
	this.ref.Once()
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
	this.ref.Free()
}

// Target returns the value of property "ProcessingInstruction.target".
//
// It returns ok=false if there is no such property.
func (this ProcessingInstruction) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetProcessingInstructionTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Sheet returns the value of property "ProcessingInstruction.sheet".
//
// It returns ok=false if there is no such property.
func (this ProcessingInstruction) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetProcessingInstructionSheet(
		this.ref, js.Pointer(&ret),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Root returns the value of property "NodeIterator.root".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) Root() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeIteratorRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReferenceNode returns the value of property "NodeIterator.referenceNode".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) ReferenceNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeIteratorReferenceNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointerBeforeReferenceNode returns the value of property "NodeIterator.pointerBeforeReferenceNode".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) PointerBeforeReferenceNode() (ret bool, ok bool) {
	ok = js.True == bindings.GetNodeIteratorPointerBeforeReferenceNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WhatToShow returns the value of property "NodeIterator.whatToShow".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) WhatToShow() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNodeIteratorWhatToShow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Filter returns the value of property "NodeIterator.filter".
//
// It returns ok=false if there is no such property.
func (this NodeIterator) Filter() (ret js.Func[func(node Node) uint16], ok bool) {
	ok = js.True == bindings.GetNodeIteratorFilter(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncNextNode returns true if the method "NodeIterator.nextNode" exists.
func (this NodeIterator) HasFuncNextNode() bool {
	return js.True == bindings.HasFuncNodeIteratorNextNode(
		this.ref,
	)
}

// FuncNextNode returns the method "NodeIterator.nextNode".
func (this NodeIterator) FuncNextNode() (fn js.Func[func() Node]) {
	bindings.FuncNodeIteratorNextNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NextNode calls the method "NodeIterator.nextNode".
func (this NodeIterator) NextNode() (ret Node) {
	bindings.CallNodeIteratorNextNode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryNextNode calls the method "NodeIterator.nextNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NodeIterator) TryNextNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorNextNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreviousNode returns true if the method "NodeIterator.previousNode" exists.
func (this NodeIterator) HasFuncPreviousNode() bool {
	return js.True == bindings.HasFuncNodeIteratorPreviousNode(
		this.ref,
	)
}

// FuncPreviousNode returns the method "NodeIterator.previousNode".
func (this NodeIterator) FuncPreviousNode() (fn js.Func[func() Node]) {
	bindings.FuncNodeIteratorPreviousNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreviousNode calls the method "NodeIterator.previousNode".
func (this NodeIterator) PreviousNode() (ret Node) {
	bindings.CallNodeIteratorPreviousNode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreviousNode calls the method "NodeIterator.previousNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NodeIterator) TryPreviousNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorPreviousNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDetach returns true if the method "NodeIterator.detach" exists.
func (this NodeIterator) HasFuncDetach() bool {
	return js.True == bindings.HasFuncNodeIteratorDetach(
		this.ref,
	)
}

// FuncDetach returns the method "NodeIterator.detach".
func (this NodeIterator) FuncDetach() (fn js.Func[func()]) {
	bindings.FuncNodeIteratorDetach(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Detach calls the method "NodeIterator.detach".
func (this NodeIterator) Detach() (ret js.Void) {
	bindings.CallNodeIteratorDetach(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDetach calls the method "NodeIterator.detach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NodeIterator) TryDetach() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIteratorDetach(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type TreeWalker struct {
	ref js.Ref
}

func (this TreeWalker) Once() TreeWalker {
	this.ref.Once()
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
	this.ref.Free()
}

// Root returns the value of property "TreeWalker.root".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) Root() (ret Node, ok bool) {
	ok = js.True == bindings.GetTreeWalkerRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WhatToShow returns the value of property "TreeWalker.whatToShow".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) WhatToShow() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTreeWalkerWhatToShow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Filter returns the value of property "TreeWalker.filter".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) Filter() (ret js.Func[func(node Node) uint16], ok bool) {
	ok = js.True == bindings.GetTreeWalkerFilter(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentNode returns the value of property "TreeWalker.currentNode".
//
// It returns ok=false if there is no such property.
func (this TreeWalker) CurrentNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetTreeWalkerCurrentNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCurrentNode sets the value of property "TreeWalker.currentNode" to val.
//
// It returns false if the property cannot be set.
func (this TreeWalker) SetCurrentNode(val Node) bool {
	return js.True == bindings.SetTreeWalkerCurrentNode(
		this.ref,
		val.Ref(),
	)
}

// HasFuncParentNode returns true if the method "TreeWalker.parentNode" exists.
func (this TreeWalker) HasFuncParentNode() bool {
	return js.True == bindings.HasFuncTreeWalkerParentNode(
		this.ref,
	)
}

// FuncParentNode returns the method "TreeWalker.parentNode".
func (this TreeWalker) FuncParentNode() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerParentNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ParentNode calls the method "TreeWalker.parentNode".
func (this TreeWalker) ParentNode() (ret Node) {
	bindings.CallTreeWalkerParentNode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryParentNode calls the method "TreeWalker.parentNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryParentNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerParentNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFirstChild returns true if the method "TreeWalker.firstChild" exists.
func (this TreeWalker) HasFuncFirstChild() bool {
	return js.True == bindings.HasFuncTreeWalkerFirstChild(
		this.ref,
	)
}

// FuncFirstChild returns the method "TreeWalker.firstChild".
func (this TreeWalker) FuncFirstChild() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerFirstChild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FirstChild calls the method "TreeWalker.firstChild".
func (this TreeWalker) FirstChild() (ret Node) {
	bindings.CallTreeWalkerFirstChild(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFirstChild calls the method "TreeWalker.firstChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryFirstChild() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerFirstChild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLastChild returns true if the method "TreeWalker.lastChild" exists.
func (this TreeWalker) HasFuncLastChild() bool {
	return js.True == bindings.HasFuncTreeWalkerLastChild(
		this.ref,
	)
}

// FuncLastChild returns the method "TreeWalker.lastChild".
func (this TreeWalker) FuncLastChild() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerLastChild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LastChild calls the method "TreeWalker.lastChild".
func (this TreeWalker) LastChild() (ret Node) {
	bindings.CallTreeWalkerLastChild(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLastChild calls the method "TreeWalker.lastChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryLastChild() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerLastChild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreviousSibling returns true if the method "TreeWalker.previousSibling" exists.
func (this TreeWalker) HasFuncPreviousSibling() bool {
	return js.True == bindings.HasFuncTreeWalkerPreviousSibling(
		this.ref,
	)
}

// FuncPreviousSibling returns the method "TreeWalker.previousSibling".
func (this TreeWalker) FuncPreviousSibling() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerPreviousSibling(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreviousSibling calls the method "TreeWalker.previousSibling".
func (this TreeWalker) PreviousSibling() (ret Node) {
	bindings.CallTreeWalkerPreviousSibling(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreviousSibling calls the method "TreeWalker.previousSibling"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryPreviousSibling() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerPreviousSibling(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNextSibling returns true if the method "TreeWalker.nextSibling" exists.
func (this TreeWalker) HasFuncNextSibling() bool {
	return js.True == bindings.HasFuncTreeWalkerNextSibling(
		this.ref,
	)
}

// FuncNextSibling returns the method "TreeWalker.nextSibling".
func (this TreeWalker) FuncNextSibling() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerNextSibling(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NextSibling calls the method "TreeWalker.nextSibling".
func (this TreeWalker) NextSibling() (ret Node) {
	bindings.CallTreeWalkerNextSibling(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryNextSibling calls the method "TreeWalker.nextSibling"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryNextSibling() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerNextSibling(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreviousNode returns true if the method "TreeWalker.previousNode" exists.
func (this TreeWalker) HasFuncPreviousNode() bool {
	return js.True == bindings.HasFuncTreeWalkerPreviousNode(
		this.ref,
	)
}

// FuncPreviousNode returns the method "TreeWalker.previousNode".
func (this TreeWalker) FuncPreviousNode() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerPreviousNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreviousNode calls the method "TreeWalker.previousNode".
func (this TreeWalker) PreviousNode() (ret Node) {
	bindings.CallTreeWalkerPreviousNode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreviousNode calls the method "TreeWalker.previousNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryPreviousNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerPreviousNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNextNode returns true if the method "TreeWalker.nextNode" exists.
func (this TreeWalker) HasFuncNextNode() bool {
	return js.True == bindings.HasFuncTreeWalkerNextNode(
		this.ref,
	)
}

// FuncNextNode returns the method "TreeWalker.nextNode".
func (this TreeWalker) FuncNextNode() (fn js.Func[func() Node]) {
	bindings.FuncTreeWalkerNextNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NextNode calls the method "TreeWalker.nextNode".
func (this TreeWalker) NextNode() (ret Node) {
	bindings.CallTreeWalkerNextNode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryNextNode calls the method "TreeWalker.nextNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TreeWalker) TryNextNode() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTreeWalkerNextNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ViewTransition struct {
	ref js.Ref
}

func (this ViewTransition) Once() ViewTransition {
	this.ref.Once()
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
	this.ref.Free()
}

// UpdateCallbackDone returns the value of property "ViewTransition.updateCallbackDone".
//
// It returns ok=false if there is no such property.
func (this ViewTransition) UpdateCallbackDone() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionUpdateCallbackDone(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "ViewTransition.ready".
//
// It returns ok=false if there is no such property.
func (this ViewTransition) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "ViewTransition.finished".
//
// It returns ok=false if there is no such property.
func (this ViewTransition) Finished() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetViewTransitionFinished(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSkipTransition returns true if the method "ViewTransition.skipTransition" exists.
func (this ViewTransition) HasFuncSkipTransition() bool {
	return js.True == bindings.HasFuncViewTransitionSkipTransition(
		this.ref,
	)
}

// FuncSkipTransition returns the method "ViewTransition.skipTransition".
func (this ViewTransition) FuncSkipTransition() (fn js.Func[func()]) {
	bindings.FuncViewTransitionSkipTransition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkipTransition calls the method "ViewTransition.skipTransition".
func (this ViewTransition) SkipTransition() (ret js.Void) {
	bindings.CallViewTransitionSkipTransition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkipTransition calls the method "ViewTransition.skipTransition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ViewTransition) TrySkipTransition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryViewTransitionSkipTransition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

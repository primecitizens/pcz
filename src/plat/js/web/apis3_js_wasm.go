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

type Baseline struct {
	ref js.Ref
}

func (this Baseline) Once() Baseline {
	this.Ref().Once()
	return this
}

func (this Baseline) Ref() js.Ref {
	return this.ref
}

func (this Baseline) FromRef(ref js.Ref) Baseline {
	this.ref = ref
	return this
}

func (this Baseline) Free() {
	this.Ref().Free()
}

// Name returns the value of property "Baseline.name".
//
// The returned bool will be false if there is no such property.
func (this Baseline) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBaselineName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value returns the value of property "Baseline.value".
//
// The returned bool will be false if there is no such property.
func (this Baseline) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetBaselineValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type Font struct {
	ref js.Ref
}

func (this Font) Once() Font {
	this.Ref().Once()
	return this
}

func (this Font) Ref() js.Ref {
	return this.ref
}

func (this Font) FromRef(ref js.Ref) Font {
	this.ref = ref
	return this
}

func (this Font) Free() {
	this.Ref().Free()
}

// Name returns the value of property "Font.name".
//
// The returned bool will be false if there is no such property.
func (this Font) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GlyphsRendered returns the value of property "Font.glyphsRendered".
//
// The returned bool will be false if there is no such property.
func (this Font) GlyphsRendered() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetFontGlyphsRendered(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

type FontMetrics struct {
	ref js.Ref
}

func (this FontMetrics) Once() FontMetrics {
	this.Ref().Once()
	return this
}

func (this FontMetrics) Ref() js.Ref {
	return this.ref
}

func (this FontMetrics) FromRef(ref js.Ref) FontMetrics {
	this.ref = ref
	return this
}

func (this FontMetrics) Free() {
	this.Ref().Free()
}

// Width returns the value of property "FontMetrics.width".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Advances returns the value of property "FontMetrics.advances".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) Advances() (js.FrozenArray[float64], bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsAdvances(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[float64]{}.FromRef(_ret), _ok
}

// BoundingBoxLeft returns the value of property "FontMetrics.boundingBoxLeft".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) BoundingBoxLeft() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsBoundingBoxLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BoundingBoxRight returns the value of property "FontMetrics.boundingBoxRight".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) BoundingBoxRight() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsBoundingBoxRight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height returns the value of property "FontMetrics.height".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// EmHeightAscent returns the value of property "FontMetrics.emHeightAscent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) EmHeightAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsEmHeightAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// EmHeightDescent returns the value of property "FontMetrics.emHeightDescent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) EmHeightDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsEmHeightDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BoundingBoxAscent returns the value of property "FontMetrics.boundingBoxAscent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) BoundingBoxAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsBoundingBoxAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BoundingBoxDescent returns the value of property "FontMetrics.boundingBoxDescent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) BoundingBoxDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsBoundingBoxDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FontBoundingBoxAscent returns the value of property "FontMetrics.fontBoundingBoxAscent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) FontBoundingBoxAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsFontBoundingBoxAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FontBoundingBoxDescent returns the value of property "FontMetrics.fontBoundingBoxDescent".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) FontBoundingBoxDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsFontBoundingBoxDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DominantBaseline returns the value of property "FontMetrics.dominantBaseline".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) DominantBaseline() (Baseline, bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsDominantBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return Baseline{}.FromRef(_ret), _ok
}

// Baselines returns the value of property "FontMetrics.baselines".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) Baselines() (js.FrozenArray[Baseline], bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsBaselines(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Baseline]{}.FromRef(_ret), _ok
}

// Fonts returns the value of property "FontMetrics.fonts".
//
// The returned bool will be false if there is no such property.
func (this FontMetrics) Fonts() (js.FrozenArray[Font], bool) {
	var _ok bool
	_ret := bindings.GetFontMetricsFonts(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Font]{}.FromRef(_ret), _ok
}

type StaticRangeInit struct {
	// StartContainer is "StaticRangeInit.startContainer"
	//
	// Required
	StartContainer Node
	// StartOffset is "StaticRangeInit.startOffset"
	//
	// Required
	StartOffset uint32
	// EndContainer is "StaticRangeInit.endContainer"
	//
	// Required
	EndContainer Node
	// EndOffset is "StaticRangeInit.endOffset"
	//
	// Required
	EndOffset uint32

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StaticRangeInit with all fields set.
func (p StaticRangeInit) FromRef(ref js.Ref) StaticRangeInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StaticRangeInit in the application heap.
func (p StaticRangeInit) New() js.Ref {
	return bindings.StaticRangeInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p StaticRangeInit) UpdateFrom(ref js.Ref) {
	bindings.StaticRangeInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StaticRangeInit) Update(ref js.Ref) {
	bindings.StaticRangeInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewStaticRange(init StaticRangeInit) StaticRange {
	return StaticRange{}.FromRef(
		bindings.NewStaticRangeByStaticRange(
			js.Pointer(&init)),
	)
}

type StaticRange struct {
	AbstractRange
}

func (this StaticRange) Once() StaticRange {
	this.Ref().Once()
	return this
}

func (this StaticRange) Ref() js.Ref {
	return this.AbstractRange.Ref()
}

func (this StaticRange) FromRef(ref js.Ref) StaticRange {
	this.AbstractRange = this.AbstractRange.FromRef(ref)
	return this
}

func (this StaticRange) Free() {
	this.Ref().Free()
}

type Selection struct {
	ref js.Ref
}

func (this Selection) Once() Selection {
	this.Ref().Once()
	return this
}

func (this Selection) Ref() js.Ref {
	return this.ref
}

func (this Selection) FromRef(ref js.Ref) Selection {
	this.ref = ref
	return this
}

func (this Selection) Free() {
	this.Ref().Free()
}

// AnchorNode returns the value of property "Selection.anchorNode".
//
// The returned bool will be false if there is no such property.
func (this Selection) AnchorNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetSelectionAnchorNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// AnchorOffset returns the value of property "Selection.anchorOffset".
//
// The returned bool will be false if there is no such property.
func (this Selection) AnchorOffset() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSelectionAnchorOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// FocusNode returns the value of property "Selection.focusNode".
//
// The returned bool will be false if there is no such property.
func (this Selection) FocusNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetSelectionFocusNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// FocusOffset returns the value of property "Selection.focusOffset".
//
// The returned bool will be false if there is no such property.
func (this Selection) FocusOffset() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSelectionFocusOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IsCollapsed returns the value of property "Selection.isCollapsed".
//
// The returned bool will be false if there is no such property.
func (this Selection) IsCollapsed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSelectionIsCollapsed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// RangeCount returns the value of property "Selection.rangeCount".
//
// The returned bool will be false if there is no such property.
func (this Selection) RangeCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSelectionRangeCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Type returns the value of property "Selection.type".
//
// The returned bool will be false if there is no such property.
func (this Selection) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSelectionType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Direction returns the value of property "Selection.direction".
//
// The returned bool will be false if there is no such property.
func (this Selection) Direction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSelectionDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GetRangeAt calls the method "Selection.getRangeAt".
//
// The returned bool will be false if there is no such method.
func (this Selection) GetRangeAt(index uint32) (Range, bool) {
	var _ok bool
	_ret := bindings.CallSelectionGetRangeAt(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Range{}.FromRef(_ret), _ok
}

// GetRangeAtFunc returns the method "Selection.getRangeAt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) GetRangeAtFunc() (fn js.Func[func(index uint32) Range]) {
	return fn.FromRef(
		bindings.SelectionGetRangeAtFunc(
			this.Ref(),
		),
	)
}

// AddRange calls the method "Selection.addRange".
//
// The returned bool will be false if there is no such method.
func (this Selection) AddRange(rng Range) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionAddRange(
		this.Ref(), js.Pointer(&_ok),
		rng.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddRangeFunc returns the method "Selection.addRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) AddRangeFunc() (fn js.Func[func(rng Range)]) {
	return fn.FromRef(
		bindings.SelectionAddRangeFunc(
			this.Ref(),
		),
	)
}

// RemoveRange calls the method "Selection.removeRange".
//
// The returned bool will be false if there is no such method.
func (this Selection) RemoveRange(rng Range) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionRemoveRange(
		this.Ref(), js.Pointer(&_ok),
		rng.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveRangeFunc returns the method "Selection.removeRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) RemoveRangeFunc() (fn js.Func[func(rng Range)]) {
	return fn.FromRef(
		bindings.SelectionRemoveRangeFunc(
			this.Ref(),
		),
	)
}

// RemoveAllRanges calls the method "Selection.removeAllRanges".
//
// The returned bool will be false if there is no such method.
func (this Selection) RemoveAllRanges() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionRemoveAllRanges(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveAllRangesFunc returns the method "Selection.removeAllRanges".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) RemoveAllRangesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionRemoveAllRangesFunc(
			this.Ref(),
		),
	)
}

// Empty calls the method "Selection.empty".
//
// The returned bool will be false if there is no such method.
func (this Selection) Empty() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionEmpty(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EmptyFunc returns the method "Selection.empty".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) EmptyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionEmptyFunc(
			this.Ref(),
		),
	)
}

// GetComposedRanges calls the method "Selection.getComposedRanges".
//
// The returned bool will be false if there is no such method.
func (this Selection) GetComposedRanges(shadowRoots ...ShadowRoot) (js.Array[StaticRange], bool) {
	var _ok bool
	_ret := bindings.CallSelectionGetComposedRanges(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(shadowRoots),
		js.SizeU(len(shadowRoots)),
	)

	return js.Array[StaticRange]{}.FromRef(_ret), _ok
}

// GetComposedRangesFunc returns the method "Selection.getComposedRanges".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) GetComposedRangesFunc() (fn js.Func[func(shadowRoots ...ShadowRoot) js.Array[StaticRange]]) {
	return fn.FromRef(
		bindings.SelectionGetComposedRangesFunc(
			this.Ref(),
		),
	)
}

// Collapse calls the method "Selection.collapse".
//
// The returned bool will be false if there is no such method.
func (this Selection) Collapse(node Node, offset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionCollapse(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CollapseFunc returns the method "Selection.collapse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) CollapseFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionCollapseFunc(
			this.Ref(),
		),
	)
}

// Collapse1 calls the method "Selection.collapse".
//
// The returned bool will be false if there is no such method.
func (this Selection) Collapse1(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionCollapse1(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Collapse1Func returns the method "Selection.collapse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) Collapse1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionCollapse1Func(
			this.Ref(),
		),
	)
}

// SetPosition calls the method "Selection.setPosition".
//
// The returned bool will be false if there is no such method.
func (this Selection) SetPosition(node Node, offset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionSetPosition(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPositionFunc returns the method "Selection.setPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) SetPositionFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionSetPositionFunc(
			this.Ref(),
		),
	)
}

// SetPosition1 calls the method "Selection.setPosition".
//
// The returned bool will be false if there is no such method.
func (this Selection) SetPosition1(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionSetPosition1(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPosition1Func returns the method "Selection.setPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) SetPosition1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionSetPosition1Func(
			this.Ref(),
		),
	)
}

// CollapseToStart calls the method "Selection.collapseToStart".
//
// The returned bool will be false if there is no such method.
func (this Selection) CollapseToStart() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionCollapseToStart(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CollapseToStartFunc returns the method "Selection.collapseToStart".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) CollapseToStartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionCollapseToStartFunc(
			this.Ref(),
		),
	)
}

// CollapseToEnd calls the method "Selection.collapseToEnd".
//
// The returned bool will be false if there is no such method.
func (this Selection) CollapseToEnd() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionCollapseToEnd(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CollapseToEndFunc returns the method "Selection.collapseToEnd".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) CollapseToEndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionCollapseToEndFunc(
			this.Ref(),
		),
	)
}

// Extend calls the method "Selection.extend".
//
// The returned bool will be false if there is no such method.
func (this Selection) Extend(node Node, offset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionExtend(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		uint32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ExtendFunc returns the method "Selection.extend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) ExtendFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionExtendFunc(
			this.Ref(),
		),
	)
}

// Extend1 calls the method "Selection.extend".
//
// The returned bool will be false if there is no such method.
func (this Selection) Extend1(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionExtend1(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Extend1Func returns the method "Selection.extend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) Extend1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionExtend1Func(
			this.Ref(),
		),
	)
}

// SetBaseAndExtent calls the method "Selection.setBaseAndExtent".
//
// The returned bool will be false if there is no such method.
func (this Selection) SetBaseAndExtent(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionSetBaseAndExtent(
		this.Ref(), js.Pointer(&_ok),
		anchorNode.Ref(),
		uint32(anchorOffset),
		focusNode.Ref(),
		uint32(focusOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBaseAndExtentFunc returns the method "Selection.setBaseAndExtent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) SetBaseAndExtentFunc() (fn js.Func[func(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32)]) {
	return fn.FromRef(
		bindings.SelectionSetBaseAndExtentFunc(
			this.Ref(),
		),
	)
}

// SelectAllChildren calls the method "Selection.selectAllChildren".
//
// The returned bool will be false if there is no such method.
func (this Selection) SelectAllChildren(node Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionSelectAllChildren(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectAllChildrenFunc returns the method "Selection.selectAllChildren".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) SelectAllChildrenFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionSelectAllChildrenFunc(
			this.Ref(),
		),
	)
}

// Modify calls the method "Selection.modify".
//
// The returned bool will be false if there is no such method.
func (this Selection) Modify(alter js.String, direction js.String, granularity js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionModify(
		this.Ref(), js.Pointer(&_ok),
		alter.Ref(),
		direction.Ref(),
		granularity.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ModifyFunc returns the method "Selection.modify".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) ModifyFunc() (fn js.Func[func(alter js.String, direction js.String, granularity js.String)]) {
	return fn.FromRef(
		bindings.SelectionModifyFunc(
			this.Ref(),
		),
	)
}

// Modify1 calls the method "Selection.modify".
//
// The returned bool will be false if there is no such method.
func (this Selection) Modify1(alter js.String, direction js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionModify1(
		this.Ref(), js.Pointer(&_ok),
		alter.Ref(),
		direction.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Modify1Func returns the method "Selection.modify".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) Modify1Func() (fn js.Func[func(alter js.String, direction js.String)]) {
	return fn.FromRef(
		bindings.SelectionModify1Func(
			this.Ref(),
		),
	)
}

// Modify2 calls the method "Selection.modify".
//
// The returned bool will be false if there is no such method.
func (this Selection) Modify2(alter js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionModify2(
		this.Ref(), js.Pointer(&_ok),
		alter.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Modify2Func returns the method "Selection.modify".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) Modify2Func() (fn js.Func[func(alter js.String)]) {
	return fn.FromRef(
		bindings.SelectionModify2Func(
			this.Ref(),
		),
	)
}

// Modify3 calls the method "Selection.modify".
//
// The returned bool will be false if there is no such method.
func (this Selection) Modify3() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionModify3(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Modify3Func returns the method "Selection.modify".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) Modify3Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionModify3Func(
			this.Ref(),
		),
	)
}

// DeleteFromDocument calls the method "Selection.deleteFromDocument".
//
// The returned bool will be false if there is no such method.
func (this Selection) DeleteFromDocument() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSelectionDeleteFromDocument(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFromDocumentFunc returns the method "Selection.deleteFromDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) DeleteFromDocumentFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionDeleteFromDocumentFunc(
			this.Ref(),
		),
	)
}

// ContainsNode calls the method "Selection.containsNode".
//
// The returned bool will be false if there is no such method.
func (this Selection) ContainsNode(node Node, allowPartialContainment bool) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSelectionContainsNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		js.Bool(bool(allowPartialContainment)),
	)

	return _ret == js.True, _ok
}

// ContainsNodeFunc returns the method "Selection.containsNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) ContainsNodeFunc() (fn js.Func[func(node Node, allowPartialContainment bool) bool]) {
	return fn.FromRef(
		bindings.SelectionContainsNodeFunc(
			this.Ref(),
		),
	)
}

// ContainsNode1 calls the method "Selection.containsNode".
//
// The returned bool will be false if there is no such method.
func (this Selection) ContainsNode1(node Node) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSelectionContainsNode1(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return _ret == js.True, _ok
}

// ContainsNode1Func returns the method "Selection.containsNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) ContainsNode1Func() (fn js.Func[func(node Node) bool]) {
	return fn.FromRef(
		bindings.SelectionContainsNode1Func(
			this.Ref(),
		),
	)
}

// ToString calls the method "Selection.toString".
//
// The returned bool will be false if there is no such method.
func (this Selection) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSelectionToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "Selection.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Selection) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.SelectionToStringFunc(
			this.Ref(),
		),
	)
}

type CaretPosition struct {
	ref js.Ref
}

func (this CaretPosition) Once() CaretPosition {
	this.Ref().Once()
	return this
}

func (this CaretPosition) Ref() js.Ref {
	return this.ref
}

func (this CaretPosition) FromRef(ref js.Ref) CaretPosition {
	this.ref = ref
	return this
}

func (this CaretPosition) Free() {
	this.Ref().Free()
}

// OffsetNode returns the value of property "CaretPosition.offsetNode".
//
// The returned bool will be false if there is no such property.
func (this CaretPosition) OffsetNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetCaretPositionOffsetNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// Offset returns the value of property "CaretPosition.offset".
//
// The returned bool will be false if there is no such property.
func (this CaretPosition) Offset() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCaretPositionOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// GetClientRect calls the method "CaretPosition.getClientRect".
//
// The returned bool will be false if there is no such method.
func (this CaretPosition) GetClientRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallCaretPositionGetClientRect(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetClientRectFunc returns the method "CaretPosition.getClientRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CaretPosition) GetClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.CaretPositionGetClientRectFunc(
			this.Ref(),
		),
	)
}

const (
	XPathResult_ANY_TYPE                     uint16 = 0
	XPathResult_NUMBER_TYPE                  uint16 = 1
	XPathResult_STRING_TYPE                  uint16 = 2
	XPathResult_BOOLEAN_TYPE                 uint16 = 3
	XPathResult_UNORDERED_NODE_ITERATOR_TYPE uint16 = 4
	XPathResult_ORDERED_NODE_ITERATOR_TYPE   uint16 = 5
	XPathResult_UNORDERED_NODE_SNAPSHOT_TYPE uint16 = 6
	XPathResult_ORDERED_NODE_SNAPSHOT_TYPE   uint16 = 7
	XPathResult_ANY_UNORDERED_NODE_TYPE      uint16 = 8
	XPathResult_FIRST_ORDERED_NODE_TYPE      uint16 = 9
)

type XPathResult struct {
	ref js.Ref
}

func (this XPathResult) Once() XPathResult {
	this.Ref().Once()
	return this
}

func (this XPathResult) Ref() js.Ref {
	return this.ref
}

func (this XPathResult) FromRef(ref js.Ref) XPathResult {
	this.ref = ref
	return this
}

func (this XPathResult) Free() {
	this.Ref().Free()
}

// ResultType returns the value of property "XPathResult.resultType".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) ResultType() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultResultType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// NumberValue returns the value of property "XPathResult.numberValue".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) NumberValue() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultNumberValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// StringValue returns the value of property "XPathResult.stringValue".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) StringValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultStringValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BooleanValue returns the value of property "XPathResult.booleanValue".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) BooleanValue() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultBooleanValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// SingleNodeValue returns the value of property "XPathResult.singleNodeValue".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) SingleNodeValue() (Node, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultSingleNodeValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// InvalidIteratorState returns the value of property "XPathResult.invalidIteratorState".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) InvalidIteratorState() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultInvalidIteratorState(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// SnapshotLength returns the value of property "XPathResult.snapshotLength".
//
// The returned bool will be false if there is no such property.
func (this XPathResult) SnapshotLength() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXPathResultSnapshotLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IterateNext calls the method "XPathResult.iterateNext".
//
// The returned bool will be false if there is no such method.
func (this XPathResult) IterateNext() (Node, bool) {
	var _ok bool
	_ret := bindings.CallXPathResultIterateNext(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// IterateNextFunc returns the method "XPathResult.iterateNext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathResult) IterateNextFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.XPathResultIterateNextFunc(
			this.Ref(),
		),
	)
}

// SnapshotItem calls the method "XPathResult.snapshotItem".
//
// The returned bool will be false if there is no such method.
func (this XPathResult) SnapshotItem(index uint32) (Node, bool) {
	var _ok bool
	_ret := bindings.CallXPathResultSnapshotItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Node{}.FromRef(_ret), _ok
}

// SnapshotItemFunc returns the method "XPathResult.snapshotItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathResult) SnapshotItemFunc() (fn js.Func[func(index uint32) Node]) {
	return fn.FromRef(
		bindings.XPathResultSnapshotItemFunc(
			this.Ref(),
		),
	)
}

type XPathExpression struct {
	ref js.Ref
}

func (this XPathExpression) Once() XPathExpression {
	this.Ref().Once()
	return this
}

func (this XPathExpression) Ref() js.Ref {
	return this.ref
}

func (this XPathExpression) FromRef(ref js.Ref) XPathExpression {
	this.ref = ref
	return this
}

func (this XPathExpression) Free() {
	this.Ref().Free()
}

// Evaluate calls the method "XPathExpression.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathExpression) Evaluate(contextNode Node, typ uint16, result XPathResult) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathExpressionEvaluate(
		this.Ref(), js.Pointer(&_ok),
		contextNode.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// EvaluateFunc returns the method "XPathExpression.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathExpression) EvaluateFunc() (fn js.Func[func(contextNode Node, typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "XPathExpression.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathExpression) Evaluate1(contextNode Node, typ uint16) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathExpressionEvaluate1(
		this.Ref(), js.Pointer(&_ok),
		contextNode.Ref(),
		uint32(typ),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate1Func returns the method "XPathExpression.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathExpression) Evaluate1Func() (fn js.Func[func(contextNode Node, typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "XPathExpression.evaluate".
//
// The returned bool will be false if there is no such method.
func (this XPathExpression) Evaluate2(contextNode Node) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallXPathExpressionEvaluate2(
		this.Ref(), js.Pointer(&_ok),
		contextNode.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate2Func returns the method "XPathExpression.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XPathExpression) Evaluate2Func() (fn js.Func[func(contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluate2Func(
			this.Ref(),
		),
	)
}

type XPathNSResolverFunc func(this js.Ref, prefix js.String) js.Ref

func (fn XPathNSResolverFunc) Register() js.Func[func(prefix js.String) js.String] {
	return js.RegisterCallback[func(prefix js.String) js.String](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn XPathNSResolverFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type XPathNSResolver[T any] struct {
	Fn  func(arg T, this js.Ref, prefix js.String) js.Ref
	Arg T
}

func (cb *XPathNSResolver[T]) Register() js.Func[func(prefix js.String) js.String] {
	return js.RegisterCallback[func(prefix js.String) js.String](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *XPathNSResolver[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DocumentType struct {
	Node
}

func (this DocumentType) Once() DocumentType {
	this.Ref().Once()
	return this
}

func (this DocumentType) Ref() js.Ref {
	return this.Node.Ref()
}

func (this DocumentType) FromRef(ref js.Ref) DocumentType {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this DocumentType) Free() {
	this.Ref().Free()
}

// Name returns the value of property "DocumentType.name".
//
// The returned bool will be false if there is no such property.
func (this DocumentType) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentTypeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PublicId returns the value of property "DocumentType.publicId".
//
// The returned bool will be false if there is no such property.
func (this DocumentType) PublicId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentTypePublicId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SystemId returns the value of property "DocumentType.systemId".
//
// The returned bool will be false if there is no such property.
func (this DocumentType) SystemId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentTypeSystemId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Before calls the method "DocumentType.before".
//
// The returned bool will be false if there is no such method.
func (this DocumentType) Before(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentTypeBefore(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeforeFunc returns the method "DocumentType.before".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentType) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeBeforeFunc(
			this.Ref(),
		),
	)
}

// After calls the method "DocumentType.after".
//
// The returned bool will be false if there is no such method.
func (this DocumentType) After(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentTypeAfter(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AfterFunc returns the method "DocumentType.after".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentType) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeAfterFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "DocumentType.replaceWith".
//
// The returned bool will be false if there is no such method.
func (this DocumentType) ReplaceWith(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentTypeReplaceWith(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceWithFunc returns the method "DocumentType.replaceWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentType) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeReplaceWithFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "DocumentType.remove".
//
// The returned bool will be false if there is no such method.
func (this DocumentType) Remove() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentTypeRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "DocumentType.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentType) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentTypeRemoveFunc(
			this.Ref(),
		),
	)
}

type XMLDocument struct {
	Document
}

func (this XMLDocument) Once() XMLDocument {
	this.Ref().Once()
	return this
}

func (this XMLDocument) Ref() js.Ref {
	return this.Document.Ref()
}

func (this XMLDocument) FromRef(ref js.Ref) XMLDocument {
	this.Document = this.Document.FromRef(ref)
	return this
}

func (this XMLDocument) Free() {
	this.Ref().Free()
}

type DOMImplementation struct {
	ref js.Ref
}

func (this DOMImplementation) Once() DOMImplementation {
	this.Ref().Once()
	return this
}

func (this DOMImplementation) Ref() js.Ref {
	return this.ref
}

func (this DOMImplementation) FromRef(ref js.Ref) DOMImplementation {
	this.ref = ref
	return this
}

func (this DOMImplementation) Free() {
	this.Ref().Free()
}

// CreateDocumentType calls the method "DOMImplementation.createDocumentType".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) CreateDocumentType(qualifiedName js.String, publicId js.String, systemId js.String) (DocumentType, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationCreateDocumentType(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
		publicId.Ref(),
		systemId.Ref(),
	)

	return DocumentType{}.FromRef(_ret), _ok
}

// CreateDocumentTypeFunc returns the method "DOMImplementation.createDocumentType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) CreateDocumentTypeFunc() (fn js.Func[func(qualifiedName js.String, publicId js.String, systemId js.String) DocumentType]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocumentTypeFunc(
			this.Ref(),
		),
	)
}

// CreateDocument calls the method "DOMImplementation.createDocument".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) CreateDocument(namespace js.String, qualifiedName js.String, doctype DocumentType) (XMLDocument, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationCreateDocument(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
		doctype.Ref(),
	)

	return XMLDocument{}.FromRef(_ret), _ok
}

// CreateDocumentFunc returns the method "DOMImplementation.createDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) CreateDocumentFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, doctype DocumentType) XMLDocument]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocumentFunc(
			this.Ref(),
		),
	)
}

// CreateDocument1 calls the method "DOMImplementation.createDocument".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) CreateDocument1(namespace js.String, qualifiedName js.String) (XMLDocument, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationCreateDocument1(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return XMLDocument{}.FromRef(_ret), _ok
}

// CreateDocument1Func returns the method "DOMImplementation.createDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) CreateDocument1Func() (fn js.Func[func(namespace js.String, qualifiedName js.String) XMLDocument]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocument1Func(
			this.Ref(),
		),
	)
}

// CreateHTMLDocument calls the method "DOMImplementation.createHTMLDocument".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) CreateHTMLDocument(title js.String) (Document, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationCreateHTMLDocument(
		this.Ref(), js.Pointer(&_ok),
		title.Ref(),
	)

	return Document{}.FromRef(_ret), _ok
}

// CreateHTMLDocumentFunc returns the method "DOMImplementation.createHTMLDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) CreateHTMLDocumentFunc() (fn js.Func[func(title js.String) Document]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateHTMLDocumentFunc(
			this.Ref(),
		),
	)
}

// CreateHTMLDocument1 calls the method "DOMImplementation.createHTMLDocument".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) CreateHTMLDocument1() (Document, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationCreateHTMLDocument1(
		this.Ref(), js.Pointer(&_ok),
	)

	return Document{}.FromRef(_ret), _ok
}

// CreateHTMLDocument1Func returns the method "DOMImplementation.createHTMLDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) CreateHTMLDocument1Func() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateHTMLDocument1Func(
			this.Ref(),
		),
	)
}

// HasFeature calls the method "DOMImplementation.hasFeature".
//
// The returned bool will be false if there is no such method.
func (this DOMImplementation) HasFeature() (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMImplementationHasFeature(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// HasFeatureFunc returns the method "DOMImplementation.hasFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMImplementation) HasFeatureFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.DOMImplementationHasFeatureFunc(
			this.Ref(),
		),
	)
}

type FragmentDirective struct {
	ref js.Ref
}

func (this FragmentDirective) Once() FragmentDirective {
	this.Ref().Once()
	return this
}

func (this FragmentDirective) Ref() js.Ref {
	return this.ref
}

func (this FragmentDirective) FromRef(ref js.Ref) FragmentDirective {
	this.ref = ref
	return this
}

func (this FragmentDirective) Free() {
	this.Ref().Free()
}

type PermissionsPolicy struct {
	ref js.Ref
}

func (this PermissionsPolicy) Once() PermissionsPolicy {
	this.Ref().Once()
	return this
}

func (this PermissionsPolicy) Ref() js.Ref {
	return this.ref
}

func (this PermissionsPolicy) FromRef(ref js.Ref) PermissionsPolicy {
	this.ref = ref
	return this
}

func (this PermissionsPolicy) Free() {
	this.Ref().Free()
}

// AllowsFeature calls the method "PermissionsPolicy.allowsFeature".
//
// The returned bool will be false if there is no such method.
func (this PermissionsPolicy) AllowsFeature(feature js.String, origin js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPermissionsPolicyAllowsFeature(
		this.Ref(), js.Pointer(&_ok),
		feature.Ref(),
		origin.Ref(),
	)

	return _ret == js.True, _ok
}

// AllowsFeatureFunc returns the method "PermissionsPolicy.allowsFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PermissionsPolicy) AllowsFeatureFunc() (fn js.Func[func(feature js.String, origin js.String) bool]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowsFeatureFunc(
			this.Ref(),
		),
	)
}

// AllowsFeature1 calls the method "PermissionsPolicy.allowsFeature".
//
// The returned bool will be false if there is no such method.
func (this PermissionsPolicy) AllowsFeature1(feature js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPermissionsPolicyAllowsFeature1(
		this.Ref(), js.Pointer(&_ok),
		feature.Ref(),
	)

	return _ret == js.True, _ok
}

// AllowsFeature1Func returns the method "PermissionsPolicy.allowsFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PermissionsPolicy) AllowsFeature1Func() (fn js.Func[func(feature js.String) bool]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowsFeature1Func(
			this.Ref(),
		),
	)
}

// Features calls the method "PermissionsPolicy.features".
//
// The returned bool will be false if there is no such method.
func (this PermissionsPolicy) Features() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsPolicyFeatures(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// FeaturesFunc returns the method "PermissionsPolicy.features".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PermissionsPolicy) FeaturesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyFeaturesFunc(
			this.Ref(),
		),
	)
}

// AllowedFeatures calls the method "PermissionsPolicy.allowedFeatures".
//
// The returned bool will be false if there is no such method.
func (this PermissionsPolicy) AllowedFeatures() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsPolicyAllowedFeatures(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// AllowedFeaturesFunc returns the method "PermissionsPolicy.allowedFeatures".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PermissionsPolicy) AllowedFeaturesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowedFeaturesFunc(
			this.Ref(),
		),
	)
}

// GetAllowlistForFeature calls the method "PermissionsPolicy.getAllowlistForFeature".
//
// The returned bool will be false if there is no such method.
func (this PermissionsPolicy) GetAllowlistForFeature(feature js.String) (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallPermissionsPolicyGetAllowlistForFeature(
		this.Ref(), js.Pointer(&_ok),
		feature.Ref(),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// GetAllowlistForFeatureFunc returns the method "PermissionsPolicy.getAllowlistForFeature".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PermissionsPolicy) GetAllowlistForFeatureFunc() (fn js.Func[func(feature js.String) js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyGetAllowlistForFeatureFunc(
			this.Ref(),
		),
	)
}

type DocumentTimelineOptions struct {
	// OriginTime is "DocumentTimelineOptions.originTime"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_OriginTime MUST be set to true to make this field effective.
	OriginTime DOMHighResTimeStamp

	FFI_USE_OriginTime bool // for OriginTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DocumentTimelineOptions with all fields set.
func (p DocumentTimelineOptions) FromRef(ref js.Ref) DocumentTimelineOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DocumentTimelineOptions in the application heap.
func (p DocumentTimelineOptions) New() js.Ref {
	return bindings.DocumentTimelineOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DocumentTimelineOptions) UpdateFrom(ref js.Ref) {
	bindings.DocumentTimelineOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DocumentTimelineOptions) Update(ref js.Ref) {
	bindings.DocumentTimelineOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDocumentTimeline(options DocumentTimelineOptions) DocumentTimeline {
	return DocumentTimeline{}.FromRef(
		bindings.NewDocumentTimelineByDocumentTimeline(
			js.Pointer(&options)),
	)
}

func NewDocumentTimelineByDocumentTimeline1() DocumentTimeline {
	return DocumentTimeline{}.FromRef(
		bindings.NewDocumentTimelineByDocumentTimeline1(),
	)
}

type DocumentTimeline struct {
	AnimationTimeline
}

func (this DocumentTimeline) Once() DocumentTimeline {
	this.Ref().Once()
	return this
}

func (this DocumentTimeline) Ref() js.Ref {
	return this.AnimationTimeline.Ref()
}

func (this DocumentTimeline) FromRef(ref js.Ref) DocumentTimeline {
	this.AnimationTimeline = this.AnimationTimeline.FromRef(ref)
	return this
}

func (this DocumentTimeline) Free() {
	this.Ref().Free()
}

type NamedFlowMap struct {
	ref js.Ref
}

func (this NamedFlowMap) Once() NamedFlowMap {
	this.Ref().Once()
	return this
}

func (this NamedFlowMap) Ref() js.Ref {
	return this.ref
}

func (this NamedFlowMap) FromRef(ref js.Ref) NamedFlowMap {
	this.ref = ref
	return this
}

func (this NamedFlowMap) Free() {
	this.Ref().Free()
}

type FocusOptions struct {
	// PreventScroll is "FocusOptions.preventScroll"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreventScroll MUST be set to true to make this field effective.
	PreventScroll bool
	// FocusVisible is "FocusOptions.focusVisible"
	//
	// Optional
	//
	// NOTE: FFI_USE_FocusVisible MUST be set to true to make this field effective.
	FocusVisible bool

	FFI_USE_PreventScroll bool // for PreventScroll.
	FFI_USE_FocusVisible  bool // for FocusVisible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FocusOptions with all fields set.
func (p FocusOptions) FromRef(ref js.Ref) FocusOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FocusOptions in the application heap.
func (p FocusOptions) New() js.Ref {
	return bindings.FocusOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FocusOptions) UpdateFrom(ref js.Ref) {
	bindings.FocusOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FocusOptions) Update(ref js.Ref) {
	bindings.FocusOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SVGAnimatedString struct {
	ref js.Ref
}

func (this SVGAnimatedString) Once() SVGAnimatedString {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedString) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedString) FromRef(ref js.Ref) SVGAnimatedString {
	this.ref = ref
	return this
}

func (this SVGAnimatedString) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedString.baseVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedString) BaseVal() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedStringBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BaseVal sets the value of property "SVGAnimatedString.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedString) SetBaseVal(val js.String) bool {
	return js.True == bindings.SetSVGAnimatedStringBaseVal(
		this.Ref(),
		val.Ref(),
	)
}

// AnimVal returns the value of property "SVGAnimatedString.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedString) AnimVal() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedStringAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSStyleDeclaration struct {
	ref js.Ref
}

func (this CSSStyleDeclaration) Once() CSSStyleDeclaration {
	this.Ref().Once()
	return this
}

func (this CSSStyleDeclaration) Ref() js.Ref {
	return this.ref
}

func (this CSSStyleDeclaration) FromRef(ref js.Ref) CSSStyleDeclaration {
	this.ref = ref
	return this
}

func (this CSSStyleDeclaration) Free() {
	this.Ref().Free()
}

// CssText returns the value of property "CSSStyleDeclaration.cssText".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleDeclaration) CssText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleDeclarationCssText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CssText sets the value of property "CSSStyleDeclaration.cssText" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleDeclaration) SetCssText(val js.String) bool {
	return js.True == bindings.SetCSSStyleDeclarationCssText(
		this.Ref(),
		val.Ref(),
	)
}

// Length returns the value of property "CSSStyleDeclaration.length".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleDeclaration) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleDeclarationLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ParentRule returns the value of property "CSSStyleDeclaration.parentRule".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleDeclaration) ParentRule() (CSSRule, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleDeclarationParentRule(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRule{}.FromRef(_ret), _ok
}

// CssFloat returns the value of property "CSSStyleDeclaration.cssFloat".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleDeclaration) CssFloat() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleDeclarationCssFloat(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CssFloat sets the value of property "CSSStyleDeclaration.cssFloat" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleDeclaration) SetCssFloat(val js.String) bool {
	return js.True == bindings.SetCSSStyleDeclarationCssFloat(
		this.Ref(),
		val.Ref(),
	)
}

// Item calls the method "CSSStyleDeclaration.item".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) Item(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "CSSStyleDeclaration.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationItemFunc(
			this.Ref(),
		),
	)
}

// GetPropertyValue calls the method "CSSStyleDeclaration.getPropertyValue".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) GetPropertyValue(property js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationGetPropertyValue(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetPropertyValueFunc returns the method "CSSStyleDeclaration.getPropertyValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) GetPropertyValueFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationGetPropertyValueFunc(
			this.Ref(),
		),
	)
}

// GetPropertyPriority calls the method "CSSStyleDeclaration.getPropertyPriority".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) GetPropertyPriority(property js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationGetPropertyPriority(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetPropertyPriorityFunc returns the method "CSSStyleDeclaration.getPropertyPriority".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) GetPropertyPriorityFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationGetPropertyPriorityFunc(
			this.Ref(),
		),
	)
}

// SetProperty calls the method "CSSStyleDeclaration.setProperty".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) SetProperty(property js.String, value js.String, priority js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationSetProperty(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		value.Ref(),
		priority.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPropertyFunc returns the method "CSSStyleDeclaration.setProperty".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) SetPropertyFunc() (fn js.Func[func(property js.String, value js.String, priority js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationSetPropertyFunc(
			this.Ref(),
		),
	)
}

// SetProperty1 calls the method "CSSStyleDeclaration.setProperty".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) SetProperty1(property js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationSetProperty1(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetProperty1Func returns the method "CSSStyleDeclaration.setProperty".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) SetProperty1Func() (fn js.Func[func(property js.String, value js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationSetProperty1Func(
			this.Ref(),
		),
	)
}

// RemoveProperty calls the method "CSSStyleDeclaration.removeProperty".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleDeclaration) RemoveProperty(property js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleDeclarationRemoveProperty(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// RemovePropertyFunc returns the method "CSSStyleDeclaration.removeProperty".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleDeclaration) RemovePropertyFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationRemovePropertyFunc(
			this.Ref(),
		),
	)
}

type OneOf_CSSStyleValue_String struct {
	ref js.Ref
}

func (x OneOf_CSSStyleValue_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CSSStyleValue_String) Free() {
	x.ref.Free()
}

func (x OneOf_CSSStyleValue_String) FromRef(ref js.Ref) OneOf_CSSStyleValue_String {
	return OneOf_CSSStyleValue_String{
		ref: ref,
	}
}

func (x OneOf_CSSStyleValue_String) CSSStyleValue() CSSStyleValue {
	return CSSStyleValue{}.FromRef(x.ref)
}

func (x OneOf_CSSStyleValue_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type StylePropertyMap struct {
	StylePropertyMapReadOnly
}

func (this StylePropertyMap) Once() StylePropertyMap {
	this.Ref().Once()
	return this
}

func (this StylePropertyMap) Ref() js.Ref {
	return this.StylePropertyMapReadOnly.Ref()
}

func (this StylePropertyMap) FromRef(ref js.Ref) StylePropertyMap {
	this.StylePropertyMapReadOnly = this.StylePropertyMapReadOnly.FromRef(ref)
	return this
}

func (this StylePropertyMap) Free() {
	this.Ref().Free()
}

// Set calls the method "StylePropertyMap.set".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMap) Set(property js.String, values ...OneOf_CSSStyleValue_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapSet(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "StylePropertyMap.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMap) SetFunc() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapSetFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "StylePropertyMap.append".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMap) Append(property js.String, values ...OneOf_CSSStyleValue_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapAppend(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "StylePropertyMap.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMap) AppendFunc() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapAppendFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "StylePropertyMap.delete".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMap) Delete(property js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapDelete(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFunc returns the method "StylePropertyMap.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMap) DeleteFunc() (fn js.Func[func(property js.String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapDeleteFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "StylePropertyMap.clear".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMap) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "StylePropertyMap.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMap) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.StylePropertyMapClearFunc(
			this.Ref(),
		),
	)
}

type OnErrorEventHandlerNonNullFunc func(this js.Ref, event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Ref

func (fn OnErrorEventHandlerNonNullFunc) Register() js.Func[func(event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Any] {
	return js.RegisterCallback[func(event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnErrorEventHandlerNonNullFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		OneOf_Event_String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Number[uint32]{}.FromRef(args[2+1]).Get(),
		js.Number[uint32]{}.FromRef(args[3+1]).Get(),
		js.Any{}.FromRef(args[4+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnErrorEventHandlerNonNull[T any] struct {
	Fn  func(arg T, this js.Ref, event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Ref
	Arg T
}

func (cb *OnErrorEventHandlerNonNull[T]) Register() js.Func[func(event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Any] {
	return js.RegisterCallback[func(event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnErrorEventHandlerNonNull[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		OneOf_Event_String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Number[uint32]{}.FromRef(args[2+1]).Get(),
		js.Number[uint32]{}.FromRef(args[3+1]).Get(),
		js.Any{}.FromRef(args[4+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OneOf_Event_String struct {
	ref js.Ref
}

func (x OneOf_Event_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Event_String) Free() {
	x.ref.Free()
}

func (x OneOf_Event_String) FromRef(ref js.Ref) OneOf_Event_String {
	return OneOf_Event_String{
		ref: ref,
	}
}

func (x OneOf_Event_String) Event() Event {
	return Event{}.FromRef(x.ref)
}

func (x OneOf_Event_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type OnErrorEventHandler = js.Func[func(event OneOf_Event_String, source js.String, lineno uint32, colno uint32, err js.Any) js.Any]

const (
	SVGLength_SVG_LENGTHTYPE_UNKNOWN    uint16 = 0
	SVGLength_SVG_LENGTHTYPE_NUMBER     uint16 = 1
	SVGLength_SVG_LENGTHTYPE_PERCENTAGE uint16 = 2
	SVGLength_SVG_LENGTHTYPE_EMS        uint16 = 3
	SVGLength_SVG_LENGTHTYPE_EXS        uint16 = 4
	SVGLength_SVG_LENGTHTYPE_PX         uint16 = 5
	SVGLength_SVG_LENGTHTYPE_CM         uint16 = 6
	SVGLength_SVG_LENGTHTYPE_MM         uint16 = 7
	SVGLength_SVG_LENGTHTYPE_IN         uint16 = 8
	SVGLength_SVG_LENGTHTYPE_PT         uint16 = 9
	SVGLength_SVG_LENGTHTYPE_PC         uint16 = 10
)

type SVGLength struct {
	ref js.Ref
}

func (this SVGLength) Once() SVGLength {
	this.Ref().Once()
	return this
}

func (this SVGLength) Ref() js.Ref {
	return this.ref
}

func (this SVGLength) FromRef(ref js.Ref) SVGLength {
	this.ref = ref
	return this
}

func (this SVGLength) Free() {
	this.Ref().Free()
}

// UnitType returns the value of property "SVGLength.unitType".
//
// The returned bool will be false if there is no such property.
func (this SVGLength) UnitType() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthUnitType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Value returns the value of property "SVGLength.value".
//
// The returned bool will be false if there is no such property.
func (this SVGLength) Value() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Value sets the value of property "SVGLength.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValue(val float32) bool {
	return js.True == bindings.SetSVGLengthValue(
		this.Ref(),
		float32(val),
	)
}

// ValueInSpecifiedUnits returns the value of property "SVGLength.valueInSpecifiedUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGLength) ValueInSpecifiedUnits() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthValueInSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// ValueInSpecifiedUnits sets the value of property "SVGLength.valueInSpecifiedUnits" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValueInSpecifiedUnits(val float32) bool {
	return js.True == bindings.SetSVGLengthValueInSpecifiedUnits(
		this.Ref(),
		float32(val),
	)
}

// ValueAsString returns the value of property "SVGLength.valueAsString".
//
// The returned bool will be false if there is no such property.
func (this SVGLength) ValueAsString() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthValueAsString(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ValueAsString sets the value of property "SVGLength.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGLengthValueAsString(
		this.Ref(),
		val.Ref(),
	)
}

// NewValueSpecifiedUnits calls the method "SVGLength.newValueSpecifiedUnits".
//
// The returned bool will be false if there is no such method.
func (this SVGLength) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	_ = _ret
	return js.Void{}, _ok
}

// NewValueSpecifiedUnitsFunc returns the method "SVGLength.newValueSpecifiedUnits".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLength) NewValueSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	return fn.FromRef(
		bindings.SVGLengthNewValueSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// ConvertToSpecifiedUnits calls the method "SVGLength.convertToSpecifiedUnits".
//
// The returned bool will be false if there is no such method.
func (this SVGLength) ConvertToSpecifiedUnits(unitType uint16) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
		uint32(unitType),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ConvertToSpecifiedUnitsFunc returns the method "SVGLength.convertToSpecifiedUnits".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLength) ConvertToSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16)]) {
	return fn.FromRef(
		bindings.SVGLengthConvertToSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

type SVGAnimatedLength struct {
	ref js.Ref
}

func (this SVGAnimatedLength) Once() SVGAnimatedLength {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedLength) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedLength) FromRef(ref js.Ref) SVGAnimatedLength {
	this.ref = ref
	return this
}

func (this SVGAnimatedLength) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedLength.baseVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedLength) BaseVal() (SVGLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedLengthBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGLength{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedLength.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedLength) AnimVal() (SVGLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedLengthAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGLength{}.FromRef(_ret), _ok
}

type SVGUseElement struct {
	SVGGraphicsElement
}

func (this SVGUseElement) Once() SVGUseElement {
	this.Ref().Once()
	return this
}

func (this SVGUseElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGUseElement) FromRef(ref js.Ref) SVGUseElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGUseElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGUseElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGUseElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGUseElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGUseElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// InstanceRoot returns the value of property "SVGUseElement.instanceRoot".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) InstanceRoot() (SVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementInstanceRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGElement{}.FromRef(_ret), _ok
}

// AnimatedInstanceRoot returns the value of property "SVGUseElement.animatedInstanceRoot".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) AnimatedInstanceRoot() (SVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementAnimatedInstanceRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGElement{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGUseElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGUseElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGUseElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type DOMStringMap struct {
	ref js.Ref
}

func (this DOMStringMap) Once() DOMStringMap {
	this.Ref().Once()
	return this
}

func (this DOMStringMap) Ref() js.Ref {
	return this.ref
}

func (this DOMStringMap) FromRef(ref js.Ref) DOMStringMap {
	this.ref = ref
	return this
}

func (this DOMStringMap) Free() {
	this.Ref().Free()
}

// Get calls the method "DOMStringMap.".
//
// The returned bool will be false if there is no such method.
func (this DOMStringMap) Get(name js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDOMStringMapGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetFunc returns the method "DOMStringMap.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMStringMap) GetFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.DOMStringMapGetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "DOMStringMap.".
//
// The returned bool will be false if there is no such method.
func (this DOMStringMap) Set(name js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDOMStringMapSet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "DOMStringMap.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMStringMap) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.DOMStringMapSetFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "DOMStringMap.".
//
// The returned bool will be false if there is no such method.
func (this DOMStringMap) Delete(name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDOMStringMapDelete(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFunc returns the method "DOMStringMap.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMStringMap) DeleteFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.DOMStringMapDeleteFunc(
			this.Ref(),
		),
	)
}

type SVGElement struct {
	Element
}

func (this SVGElement) Once() SVGElement {
	this.Ref().Once()
	return this
}

func (this SVGElement) Ref() js.Ref {
	return this.Element.Ref()
}

func (this SVGElement) FromRef(ref js.Ref) SVGElement {
	this.Element = this.Element.FromRef(ref)
	return this
}

func (this SVGElement) Free() {
	this.Ref().Free()
}

// ClassName returns the value of property "SVGElement.className".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) ClassName() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementClassName(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// OwnerSVGElement returns the value of property "SVGElement.ownerSVGElement".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) OwnerSVGElement() (SVGSVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementOwnerSVGElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGSVGElement{}.FromRef(_ret), _ok
}

// ViewportElement returns the value of property "SVGElement.viewportElement".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) ViewportElement() (SVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementViewportElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGElement{}.FromRef(_ret), _ok
}

// Style returns the value of property "SVGElement.style".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// AttributeStyleMap returns the value of property "SVGElement.attributeStyleMap".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) AttributeStyleMap() (StylePropertyMap, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementAttributeStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return StylePropertyMap{}.FromRef(_ret), _ok
}

// CorrespondingElement returns the value of property "SVGElement.correspondingElement".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) CorrespondingElement() (SVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementCorrespondingElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGElement{}.FromRef(_ret), _ok
}

// CorrespondingUseElement returns the value of property "SVGElement.correspondingUseElement".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) CorrespondingUseElement() (SVGUseElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementCorrespondingUseElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGUseElement{}.FromRef(_ret), _ok
}

// Dataset returns the value of property "SVGElement.dataset".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) Dataset() (DOMStringMap, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementDataset(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringMap{}.FromRef(_ret), _ok
}

// Nonce returns the value of property "SVGElement.nonce".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) Nonce() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementNonce(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Nonce sets the value of property "SVGElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetSVGElementNonce(
		this.Ref(),
		val.Ref(),
	)
}

// Autofocus returns the value of property "SVGElement.autofocus".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) Autofocus() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementAutofocus(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Autofocus sets the value of property "SVGElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetSVGElementAutofocus(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "SVGElement.tabIndex".
//
// The returned bool will be false if there is no such property.
func (this SVGElement) TabIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetSVGElementTabIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// TabIndex sets the value of property "SVGElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetSVGElementTabIndex(
		this.Ref(),
		int32(val),
	)
}

// Focus calls the method "SVGElement.focus".
//
// The returned bool will be false if there is no such method.
func (this SVGElement) Focus(options FocusOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGElementFocus(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FocusFunc returns the method "SVGElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.SVGElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "SVGElement.focus".
//
// The returned bool will be false if there is no such method.
func (this SVGElement) Focus1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGElementFocus1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Focus1Func returns the method "SVGElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGElementFocus1Func(
			this.Ref(),
		),
	)
}

// Blur calls the method "SVGElement.blur".
//
// The returned bool will be false if there is no such method.
func (this SVGElement) Blur() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGElementBlur(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlurFunc returns the method "SVGElement.blur".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGElementBlurFunc(
			this.Ref(),
		),
	)
}

type SVGNumber struct {
	ref js.Ref
}

func (this SVGNumber) Once() SVGNumber {
	this.Ref().Once()
	return this
}

func (this SVGNumber) Ref() js.Ref {
	return this.ref
}

func (this SVGNumber) FromRef(ref js.Ref) SVGNumber {
	this.ref = ref
	return this
}

func (this SVGNumber) Free() {
	this.Ref().Free()
}

// Value returns the value of property "SVGNumber.value".
//
// The returned bool will be false if there is no such property.
func (this SVGNumber) Value() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGNumberValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Value sets the value of property "SVGNumber.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGNumber) SetValue(val float32) bool {
	return js.True == bindings.SetSVGNumberValue(
		this.Ref(),
		float32(val),
	)
}

const (
	SVGAngle_SVG_ANGLETYPE_UNKNOWN     uint16 = 0
	SVGAngle_SVG_ANGLETYPE_UNSPECIFIED uint16 = 1
	SVGAngle_SVG_ANGLETYPE_DEG         uint16 = 2
	SVGAngle_SVG_ANGLETYPE_RAD         uint16 = 3
	SVGAngle_SVG_ANGLETYPE_GRAD        uint16 = 4
)

type SVGAngle struct {
	ref js.Ref
}

func (this SVGAngle) Once() SVGAngle {
	this.Ref().Once()
	return this
}

func (this SVGAngle) Ref() js.Ref {
	return this.ref
}

func (this SVGAngle) FromRef(ref js.Ref) SVGAngle {
	this.ref = ref
	return this
}

func (this SVGAngle) Free() {
	this.Ref().Free()
}

// UnitType returns the value of property "SVGAngle.unitType".
//
// The returned bool will be false if there is no such property.
func (this SVGAngle) UnitType() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGAngleUnitType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Value returns the value of property "SVGAngle.value".
//
// The returned bool will be false if there is no such property.
func (this SVGAngle) Value() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAngleValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Value sets the value of property "SVGAngle.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValue(val float32) bool {
	return js.True == bindings.SetSVGAngleValue(
		this.Ref(),
		float32(val),
	)
}

// ValueInSpecifiedUnits returns the value of property "SVGAngle.valueInSpecifiedUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGAngle) ValueInSpecifiedUnits() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAngleValueInSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// ValueInSpecifiedUnits sets the value of property "SVGAngle.valueInSpecifiedUnits" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValueInSpecifiedUnits(val float32) bool {
	return js.True == bindings.SetSVGAngleValueInSpecifiedUnits(
		this.Ref(),
		float32(val),
	)
}

// ValueAsString returns the value of property "SVGAngle.valueAsString".
//
// The returned bool will be false if there is no such property.
func (this SVGAngle) ValueAsString() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGAngleValueAsString(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ValueAsString sets the value of property "SVGAngle.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGAngleValueAsString(
		this.Ref(),
		val.Ref(),
	)
}

// NewValueSpecifiedUnits calls the method "SVGAngle.newValueSpecifiedUnits".
//
// The returned bool will be false if there is no such method.
func (this SVGAngle) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAngleNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	_ = _ret
	return js.Void{}, _ok
}

// NewValueSpecifiedUnitsFunc returns the method "SVGAngle.newValueSpecifiedUnits".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAngle) NewValueSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	return fn.FromRef(
		bindings.SVGAngleNewValueSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// ConvertToSpecifiedUnits calls the method "SVGAngle.convertToSpecifiedUnits".
//
// The returned bool will be false if there is no such method.
func (this SVGAngle) ConvertToSpecifiedUnits(unitType uint16) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAngleConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&_ok),
		uint32(unitType),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ConvertToSpecifiedUnitsFunc returns the method "SVGAngle.convertToSpecifiedUnits".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAngle) ConvertToSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16)]) {
	return fn.FromRef(
		bindings.SVGAngleConvertToSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

type OneOf_String_ArrayFloat64 struct {
	ref js.Ref
}

func (x OneOf_String_ArrayFloat64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayFloat64) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayFloat64) FromRef(ref js.Ref) OneOf_String_ArrayFloat64 {
	return OneOf_String_ArrayFloat64{
		ref: ref,
	}
}

func (x OneOf_String_ArrayFloat64) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayFloat64) ArrayFloat64() js.Array[float64] {
	return js.Array[float64]{}.FromRef(x.ref)
}

type DOMMatrixInit struct {
	// M13 is "DOMMatrixInit.m13"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M13 MUST be set to true to make this field effective.
	M13 float64
	// M14 is "DOMMatrixInit.m14"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M14 MUST be set to true to make this field effective.
	M14 float64
	// M23 is "DOMMatrixInit.m23"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M23 MUST be set to true to make this field effective.
	M23 float64
	// M24 is "DOMMatrixInit.m24"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M24 MUST be set to true to make this field effective.
	M24 float64
	// M31 is "DOMMatrixInit.m31"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M31 MUST be set to true to make this field effective.
	M31 float64
	// M32 is "DOMMatrixInit.m32"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M32 MUST be set to true to make this field effective.
	M32 float64
	// M33 is "DOMMatrixInit.m33"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_M33 MUST be set to true to make this field effective.
	M33 float64
	// M34 is "DOMMatrixInit.m34"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M34 MUST be set to true to make this field effective.
	M34 float64
	// M43 is "DOMMatrixInit.m43"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_M43 MUST be set to true to make this field effective.
	M43 float64
	// M44 is "DOMMatrixInit.m44"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_M44 MUST be set to true to make this field effective.
	M44 float64
	// Is2D is "DOMMatrixInit.is2D"
	//
	// Optional
	//
	// NOTE: FFI_USE_Is2D MUST be set to true to make this field effective.
	Is2D bool
	// A is "DOMMatrixInit.a"
	//
	// Optional
	//
	// NOTE: FFI_USE_A MUST be set to true to make this field effective.
	A float64
	// B is "DOMMatrixInit.b"
	//
	// Optional
	//
	// NOTE: FFI_USE_B MUST be set to true to make this field effective.
	B float64
	// C is "DOMMatrixInit.c"
	//
	// Optional
	//
	// NOTE: FFI_USE_C MUST be set to true to make this field effective.
	C float64
	// D is "DOMMatrixInit.d"
	//
	// Optional
	//
	// NOTE: FFI_USE_D MUST be set to true to make this field effective.
	D float64
	// E is "DOMMatrixInit.e"
	//
	// Optional
	//
	// NOTE: FFI_USE_E MUST be set to true to make this field effective.
	E float64
	// F is "DOMMatrixInit.f"
	//
	// Optional
	//
	// NOTE: FFI_USE_F MUST be set to true to make this field effective.
	F float64
	// M11 is "DOMMatrixInit.m11"
	//
	// Optional
	//
	// NOTE: FFI_USE_M11 MUST be set to true to make this field effective.
	M11 float64
	// M12 is "DOMMatrixInit.m12"
	//
	// Optional
	//
	// NOTE: FFI_USE_M12 MUST be set to true to make this field effective.
	M12 float64
	// M21 is "DOMMatrixInit.m21"
	//
	// Optional
	//
	// NOTE: FFI_USE_M21 MUST be set to true to make this field effective.
	M21 float64
	// M22 is "DOMMatrixInit.m22"
	//
	// Optional
	//
	// NOTE: FFI_USE_M22 MUST be set to true to make this field effective.
	M22 float64
	// M41 is "DOMMatrixInit.m41"
	//
	// Optional
	//
	// NOTE: FFI_USE_M41 MUST be set to true to make this field effective.
	M41 float64
	// M42 is "DOMMatrixInit.m42"
	//
	// Optional
	//
	// NOTE: FFI_USE_M42 MUST be set to true to make this field effective.
	M42 float64

	FFI_USE_M13  bool // for M13.
	FFI_USE_M14  bool // for M14.
	FFI_USE_M23  bool // for M23.
	FFI_USE_M24  bool // for M24.
	FFI_USE_M31  bool // for M31.
	FFI_USE_M32  bool // for M32.
	FFI_USE_M33  bool // for M33.
	FFI_USE_M34  bool // for M34.
	FFI_USE_M43  bool // for M43.
	FFI_USE_M44  bool // for M44.
	FFI_USE_Is2D bool // for Is2D.
	FFI_USE_A    bool // for A.
	FFI_USE_B    bool // for B.
	FFI_USE_C    bool // for C.
	FFI_USE_D    bool // for D.
	FFI_USE_E    bool // for E.
	FFI_USE_F    bool // for F.
	FFI_USE_M11  bool // for M11.
	FFI_USE_M12  bool // for M12.
	FFI_USE_M21  bool // for M21.
	FFI_USE_M22  bool // for M22.
	FFI_USE_M41  bool // for M41.
	FFI_USE_M42  bool // for M42.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMMatrixInit with all fields set.
func (p DOMMatrixInit) FromRef(ref js.Ref) DOMMatrixInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMMatrixInit in the application heap.
func (p DOMMatrixInit) New() js.Ref {
	return bindings.DOMMatrixInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DOMMatrixInit) UpdateFrom(ref js.Ref) {
	bindings.DOMMatrixInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DOMMatrixInit) Update(ref js.Ref) {
	bindings.DOMMatrixInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDOMMatrix(init OneOf_String_ArrayFloat64) DOMMatrix {
	return DOMMatrix{}.FromRef(
		bindings.NewDOMMatrixByDOMMatrix(
			init.Ref()),
	)
}

func NewDOMMatrixByDOMMatrix1() DOMMatrix {
	return DOMMatrix{}.FromRef(
		bindings.NewDOMMatrixByDOMMatrix1(),
	)
}

type DOMMatrix struct {
	DOMMatrixReadOnly
}

func (this DOMMatrix) Once() DOMMatrix {
	this.Ref().Once()
	return this
}

func (this DOMMatrix) Ref() js.Ref {
	return this.DOMMatrixReadOnly.Ref()
}

func (this DOMMatrix) FromRef(ref js.Ref) DOMMatrix {
	this.DOMMatrixReadOnly = this.DOMMatrixReadOnly.FromRef(ref)
	return this
}

func (this DOMMatrix) Free() {
	this.Ref().Free()
}

// A returns the value of property "DOMMatrix.a".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) A() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixA(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// A sets the value of property "DOMMatrix.a" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetA(val float64) bool {
	return js.True == bindings.SetDOMMatrixA(
		this.Ref(),
		float64(val),
	)
}

// B returns the value of property "DOMMatrix.b".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) B() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixB(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// B sets the value of property "DOMMatrix.b" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetB(val float64) bool {
	return js.True == bindings.SetDOMMatrixB(
		this.Ref(),
		float64(val),
	)
}

// C returns the value of property "DOMMatrix.c".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) C() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixC(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// C sets the value of property "DOMMatrix.c" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetC(val float64) bool {
	return js.True == bindings.SetDOMMatrixC(
		this.Ref(),
		float64(val),
	)
}

// D returns the value of property "DOMMatrix.d".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) D() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixD(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// D sets the value of property "DOMMatrix.d" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetD(val float64) bool {
	return js.True == bindings.SetDOMMatrixD(
		this.Ref(),
		float64(val),
	)
}

// E returns the value of property "DOMMatrix.e".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) E() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixE(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// E sets the value of property "DOMMatrix.e" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetE(val float64) bool {
	return js.True == bindings.SetDOMMatrixE(
		this.Ref(),
		float64(val),
	)
}

// F returns the value of property "DOMMatrix.f".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) F() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixF(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// F sets the value of property "DOMMatrix.f" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetF(val float64) bool {
	return js.True == bindings.SetDOMMatrixF(
		this.Ref(),
		float64(val),
	)
}

// M11 returns the value of property "DOMMatrix.m11".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M11() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM11(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M11 sets the value of property "DOMMatrix.m11" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM11(val float64) bool {
	return js.True == bindings.SetDOMMatrixM11(
		this.Ref(),
		float64(val),
	)
}

// M12 returns the value of property "DOMMatrix.m12".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M12() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM12(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M12 sets the value of property "DOMMatrix.m12" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM12(val float64) bool {
	return js.True == bindings.SetDOMMatrixM12(
		this.Ref(),
		float64(val),
	)
}

// M13 returns the value of property "DOMMatrix.m13".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M13() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM13(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M13 sets the value of property "DOMMatrix.m13" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM13(val float64) bool {
	return js.True == bindings.SetDOMMatrixM13(
		this.Ref(),
		float64(val),
	)
}

// M14 returns the value of property "DOMMatrix.m14".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M14() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM14(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M14 sets the value of property "DOMMatrix.m14" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM14(val float64) bool {
	return js.True == bindings.SetDOMMatrixM14(
		this.Ref(),
		float64(val),
	)
}

// M21 returns the value of property "DOMMatrix.m21".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M21() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM21(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M21 sets the value of property "DOMMatrix.m21" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM21(val float64) bool {
	return js.True == bindings.SetDOMMatrixM21(
		this.Ref(),
		float64(val),
	)
}

// M22 returns the value of property "DOMMatrix.m22".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M22() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM22(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M22 sets the value of property "DOMMatrix.m22" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM22(val float64) bool {
	return js.True == bindings.SetDOMMatrixM22(
		this.Ref(),
		float64(val),
	)
}

// M23 returns the value of property "DOMMatrix.m23".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M23() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM23(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M23 sets the value of property "DOMMatrix.m23" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM23(val float64) bool {
	return js.True == bindings.SetDOMMatrixM23(
		this.Ref(),
		float64(val),
	)
}

// M24 returns the value of property "DOMMatrix.m24".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M24() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM24(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M24 sets the value of property "DOMMatrix.m24" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM24(val float64) bool {
	return js.True == bindings.SetDOMMatrixM24(
		this.Ref(),
		float64(val),
	)
}

// M31 returns the value of property "DOMMatrix.m31".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M31() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM31(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M31 sets the value of property "DOMMatrix.m31" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM31(val float64) bool {
	return js.True == bindings.SetDOMMatrixM31(
		this.Ref(),
		float64(val),
	)
}

// M32 returns the value of property "DOMMatrix.m32".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M32() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM32(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M32 sets the value of property "DOMMatrix.m32" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM32(val float64) bool {
	return js.True == bindings.SetDOMMatrixM32(
		this.Ref(),
		float64(val),
	)
}

// M33 returns the value of property "DOMMatrix.m33".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M33() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM33(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M33 sets the value of property "DOMMatrix.m33" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM33(val float64) bool {
	return js.True == bindings.SetDOMMatrixM33(
		this.Ref(),
		float64(val),
	)
}

// M34 returns the value of property "DOMMatrix.m34".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M34() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM34(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M34 sets the value of property "DOMMatrix.m34" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM34(val float64) bool {
	return js.True == bindings.SetDOMMatrixM34(
		this.Ref(),
		float64(val),
	)
}

// M41 returns the value of property "DOMMatrix.m41".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M41() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM41(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M41 sets the value of property "DOMMatrix.m41" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM41(val float64) bool {
	return js.True == bindings.SetDOMMatrixM41(
		this.Ref(),
		float64(val),
	)
}

// M42 returns the value of property "DOMMatrix.m42".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M42() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM42(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M42 sets the value of property "DOMMatrix.m42" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM42(val float64) bool {
	return js.True == bindings.SetDOMMatrixM42(
		this.Ref(),
		float64(val),
	)
}

// M43 returns the value of property "DOMMatrix.m43".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M43() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM43(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M43 sets the value of property "DOMMatrix.m43" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM43(val float64) bool {
	return js.True == bindings.SetDOMMatrixM43(
		this.Ref(),
		float64(val),
	)
}

// M44 returns the value of property "DOMMatrix.m44".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrix) M44() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixM44(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M44 sets the value of property "DOMMatrix.m44" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM44(val float64) bool {
	return js.True == bindings.SetDOMMatrixM44(
		this.Ref(),
		float64(val),
	)
}

// FromMatrix calls the staticmethod "DOMMatrix.fromMatrix".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) FromMatrix(other DOMMatrixInit) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixFromMatrix(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FromMatrixFunc returns the staticmethod "DOMMatrix.fromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) FromMatrixFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromMatrixFunc(
			this.Ref(),
		),
	)
}

// FromMatrix1 calls the staticmethod "DOMMatrix.fromMatrix".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) FromMatrix1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixFromMatrix1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FromMatrix1Func returns the staticmethod "DOMMatrix.fromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) FromMatrix1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromMatrix1Func(
			this.Ref(),
		),
	)
}

// FromFloat32Array calls the staticmethod "DOMMatrix.fromFloat32Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) FromFloat32Array(array32 js.TypedArray[float32]) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixFromFloat32Array(
		this.Ref(), js.Pointer(&_ok),
		array32.Ref(),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FromFloat32ArrayFunc returns the staticmethod "DOMMatrix.fromFloat32Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) FromFloat32ArrayFunc() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat64Array calls the staticmethod "DOMMatrix.fromFloat64Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) FromFloat64Array(array64 js.TypedArray[float64]) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixFromFloat64Array(
		this.Ref(), js.Pointer(&_ok),
		array64.Ref(),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FromFloat64ArrayFunc returns the staticmethod "DOMMatrix.fromFloat64Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) FromFloat64ArrayFunc() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// MultiplySelf calls the method "DOMMatrix.multiplySelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) MultiplySelf(other DOMMatrixInit) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixMultiplySelf(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// MultiplySelfFunc returns the method "DOMMatrix.multiplySelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) MultiplySelfFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixMultiplySelfFunc(
			this.Ref(),
		),
	)
}

// MultiplySelf1 calls the method "DOMMatrix.multiplySelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) MultiplySelf1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixMultiplySelf1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// MultiplySelf1Func returns the method "DOMMatrix.multiplySelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) MultiplySelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixMultiplySelf1Func(
			this.Ref(),
		),
	)
}

// PreMultiplySelf calls the method "DOMMatrix.preMultiplySelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) PreMultiplySelf(other DOMMatrixInit) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixPreMultiplySelf(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// PreMultiplySelfFunc returns the method "DOMMatrix.preMultiplySelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) PreMultiplySelfFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixPreMultiplySelfFunc(
			this.Ref(),
		),
	)
}

// PreMultiplySelf1 calls the method "DOMMatrix.preMultiplySelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) PreMultiplySelf1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixPreMultiplySelf1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// PreMultiplySelf1Func returns the method "DOMMatrix.preMultiplySelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) PreMultiplySelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixPreMultiplySelf1Func(
			this.Ref(),
		),
	)
}

// TranslateSelf calls the method "DOMMatrix.translateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) TranslateSelf(tx float64, ty float64, tz float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixTranslateSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// TranslateSelfFunc returns the method "DOMMatrix.translateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) TranslateSelfFunc() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelfFunc(
			this.Ref(),
		),
	)
}

// TranslateSelf1 calls the method "DOMMatrix.translateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) TranslateSelf1(tx float64, ty float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixTranslateSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
		float64(ty),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// TranslateSelf1Func returns the method "DOMMatrix.translateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) TranslateSelf1Func() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf1Func(
			this.Ref(),
		),
	)
}

// TranslateSelf2 calls the method "DOMMatrix.translateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) TranslateSelf2(tx float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixTranslateSelf2(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// TranslateSelf2Func returns the method "DOMMatrix.translateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) TranslateSelf2Func() (fn js.Func[func(tx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf2Func(
			this.Ref(),
		),
	)
}

// TranslateSelf3 calls the method "DOMMatrix.translateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) TranslateSelf3() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixTranslateSelf3(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// TranslateSelf3Func returns the method "DOMMatrix.translateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) TranslateSelf3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf3Func(
			this.Ref(),
		),
	)
}

// ScaleSelf calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelfFunc returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelfFunc() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelfFunc(
			this.Ref(),
		),
	)
}

// ScaleSelf1 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf1Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf1Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf1Func(
			this.Ref(),
		),
	)
}

// ScaleSelf2 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf2(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf2Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf2Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf2Func(
			this.Ref(),
		),
	)
}

// ScaleSelf3 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf3(scaleX float64, scaleY float64, scaleZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf3(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf3Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf3Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf3Func(
			this.Ref(),
		),
	)
}

// ScaleSelf4 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf4(scaleX float64, scaleY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf4(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf4Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf4Func() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf4Func(
			this.Ref(),
		),
	)
}

// ScaleSelf5 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf5(scaleX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf5(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf5Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf5Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf5Func(
			this.Ref(),
		),
	)
}

// ScaleSelf6 calls the method "DOMMatrix.scaleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) ScaleSelf6() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScaleSelf6(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleSelf6Func returns the method "DOMMatrix.scaleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) ScaleSelf6Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf6Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf calls the method "DOMMatrix.scale3dSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) Scale3dSelf(scale float64, originX float64, originY float64, originZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScale3dSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dSelfFunc returns the method "DOMMatrix.scale3dSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) Scale3dSelfFunc() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelfFunc(
			this.Ref(),
		),
	)
}

// Scale3dSelf1 calls the method "DOMMatrix.scale3dSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) Scale3dSelf1(scale float64, originX float64, originY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScale3dSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dSelf1Func returns the method "DOMMatrix.scale3dSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) Scale3dSelf1Func() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf1Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf2 calls the method "DOMMatrix.scale3dSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) Scale3dSelf2(scale float64, originX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScale3dSelf2(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dSelf2Func returns the method "DOMMatrix.scale3dSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) Scale3dSelf2Func() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf2Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf3 calls the method "DOMMatrix.scale3dSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) Scale3dSelf3(scale float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScale3dSelf3(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dSelf3Func returns the method "DOMMatrix.scale3dSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) Scale3dSelf3Func() (fn js.Func[func(scale float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf3Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf4 calls the method "DOMMatrix.scale3dSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) Scale3dSelf4() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixScale3dSelf4(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dSelf4Func returns the method "DOMMatrix.scale3dSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) Scale3dSelf4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf4Func(
			this.Ref(),
		),
	)
}

// RotateSelf calls the method "DOMMatrix.rotateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateSelf(rotX float64, rotY float64, rotZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateSelfFunc returns the method "DOMMatrix.rotateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateSelfFunc() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelfFunc(
			this.Ref(),
		),
	)
}

// RotateSelf1 calls the method "DOMMatrix.rotateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateSelf1(rotX float64, rotY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
		float64(rotY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateSelf1Func returns the method "DOMMatrix.rotateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateSelf1Func() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf1Func(
			this.Ref(),
		),
	)
}

// RotateSelf2 calls the method "DOMMatrix.rotateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateSelf2(rotX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateSelf2(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateSelf2Func returns the method "DOMMatrix.rotateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateSelf2Func() (fn js.Func[func(rotX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf2Func(
			this.Ref(),
		),
	)
}

// RotateSelf3 calls the method "DOMMatrix.rotateSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateSelf3() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateSelf3(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateSelf3Func returns the method "DOMMatrix.rotateSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateSelf3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf3Func(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf calls the method "DOMMatrix.rotateFromVectorSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateFromVectorSelf(x float64, y float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateFromVectorSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVectorSelfFunc returns the method "DOMMatrix.rotateFromVectorSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateFromVectorSelfFunc() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelfFunc(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf1 calls the method "DOMMatrix.rotateFromVectorSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateFromVectorSelf1(x float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateFromVectorSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVectorSelf1Func returns the method "DOMMatrix.rotateFromVectorSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateFromVectorSelf1Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelf1Func(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf2 calls the method "DOMMatrix.rotateFromVectorSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateFromVectorSelf2() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateFromVectorSelf2(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVectorSelf2Func returns the method "DOMMatrix.rotateFromVectorSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateFromVectorSelf2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelf2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf calls the method "DOMMatrix.rotateAxisAngleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateAxisAngleSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleSelfFunc returns the method "DOMMatrix.rotateAxisAngleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelfFunc() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelfFunc(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf1 calls the method "DOMMatrix.rotateAxisAngleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf1(x float64, y float64, z float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateAxisAngleSelf1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(z),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleSelf1Func returns the method "DOMMatrix.rotateAxisAngleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf1Func() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf1Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf2 calls the method "DOMMatrix.rotateAxisAngleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf2(x float64, y float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateAxisAngleSelf2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleSelf2Func returns the method "DOMMatrix.rotateAxisAngleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf2Func() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf3 calls the method "DOMMatrix.rotateAxisAngleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf3(x float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateAxisAngleSelf3(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleSelf3Func returns the method "DOMMatrix.rotateAxisAngleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf3Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf3Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf4 calls the method "DOMMatrix.rotateAxisAngleSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf4() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixRotateAxisAngleSelf4(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleSelf4Func returns the method "DOMMatrix.rotateAxisAngleSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) RotateAxisAngleSelf4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf4Func(
			this.Ref(),
		),
	)
}

// SkewXSelf calls the method "DOMMatrix.skewXSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) SkewXSelf(sx float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixSkewXSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(sx),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewXSelfFunc returns the method "DOMMatrix.skewXSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) SkewXSelfFunc() (fn js.Func[func(sx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewXSelfFunc(
			this.Ref(),
		),
	)
}

// SkewXSelf1 calls the method "DOMMatrix.skewXSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) SkewXSelf1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixSkewXSelf1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewXSelf1Func returns the method "DOMMatrix.skewXSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) SkewXSelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewXSelf1Func(
			this.Ref(),
		),
	)
}

// SkewYSelf calls the method "DOMMatrix.skewYSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) SkewYSelf(sy float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixSkewYSelf(
		this.Ref(), js.Pointer(&_ok),
		float64(sy),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewYSelfFunc returns the method "DOMMatrix.skewYSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) SkewYSelfFunc() (fn js.Func[func(sy float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewYSelfFunc(
			this.Ref(),
		),
	)
}

// SkewYSelf1 calls the method "DOMMatrix.skewYSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) SkewYSelf1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixSkewYSelf1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewYSelf1Func returns the method "DOMMatrix.skewYSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) SkewYSelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewYSelf1Func(
			this.Ref(),
		),
	)
}

// InvertSelf calls the method "DOMMatrix.invertSelf".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) InvertSelf() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixInvertSelf(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// InvertSelfFunc returns the method "DOMMatrix.invertSelf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) InvertSelfFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixInvertSelfFunc(
			this.Ref(),
		),
	)
}

// SetMatrixValue calls the method "DOMMatrix.setMatrixValue".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrix) SetMatrixValue(transformList js.String) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixSetMatrixValue(
		this.Ref(), js.Pointer(&_ok),
		transformList.Ref(),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SetMatrixValueFunc returns the method "DOMMatrix.setMatrixValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrix) SetMatrixValueFunc() (fn js.Func[func(transformList js.String) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSetMatrixValueFunc(
			this.Ref(),
		),
	)
}

const (
	SVGTransform_SVG_TRANSFORM_UNKNOWN   uint16 = 0
	SVGTransform_SVG_TRANSFORM_MATRIX    uint16 = 1
	SVGTransform_SVG_TRANSFORM_TRANSLATE uint16 = 2
	SVGTransform_SVG_TRANSFORM_SCALE     uint16 = 3
	SVGTransform_SVG_TRANSFORM_ROTATE    uint16 = 4
	SVGTransform_SVG_TRANSFORM_SKEWX     uint16 = 5
	SVGTransform_SVG_TRANSFORM_SKEWY     uint16 = 6
)

type DOMMatrix2DInit struct {
	// A is "DOMMatrix2DInit.a"
	//
	// Optional
	//
	// NOTE: FFI_USE_A MUST be set to true to make this field effective.
	A float64
	// B is "DOMMatrix2DInit.b"
	//
	// Optional
	//
	// NOTE: FFI_USE_B MUST be set to true to make this field effective.
	B float64
	// C is "DOMMatrix2DInit.c"
	//
	// Optional
	//
	// NOTE: FFI_USE_C MUST be set to true to make this field effective.
	C float64
	// D is "DOMMatrix2DInit.d"
	//
	// Optional
	//
	// NOTE: FFI_USE_D MUST be set to true to make this field effective.
	D float64
	// E is "DOMMatrix2DInit.e"
	//
	// Optional
	//
	// NOTE: FFI_USE_E MUST be set to true to make this field effective.
	E float64
	// F is "DOMMatrix2DInit.f"
	//
	// Optional
	//
	// NOTE: FFI_USE_F MUST be set to true to make this field effective.
	F float64
	// M11 is "DOMMatrix2DInit.m11"
	//
	// Optional
	//
	// NOTE: FFI_USE_M11 MUST be set to true to make this field effective.
	M11 float64
	// M12 is "DOMMatrix2DInit.m12"
	//
	// Optional
	//
	// NOTE: FFI_USE_M12 MUST be set to true to make this field effective.
	M12 float64
	// M21 is "DOMMatrix2DInit.m21"
	//
	// Optional
	//
	// NOTE: FFI_USE_M21 MUST be set to true to make this field effective.
	M21 float64
	// M22 is "DOMMatrix2DInit.m22"
	//
	// Optional
	//
	// NOTE: FFI_USE_M22 MUST be set to true to make this field effective.
	M22 float64
	// M41 is "DOMMatrix2DInit.m41"
	//
	// Optional
	//
	// NOTE: FFI_USE_M41 MUST be set to true to make this field effective.
	M41 float64
	// M42 is "DOMMatrix2DInit.m42"
	//
	// Optional
	//
	// NOTE: FFI_USE_M42 MUST be set to true to make this field effective.
	M42 float64

	FFI_USE_A   bool // for A.
	FFI_USE_B   bool // for B.
	FFI_USE_C   bool // for C.
	FFI_USE_D   bool // for D.
	FFI_USE_E   bool // for E.
	FFI_USE_F   bool // for F.
	FFI_USE_M11 bool // for M11.
	FFI_USE_M12 bool // for M12.
	FFI_USE_M21 bool // for M21.
	FFI_USE_M22 bool // for M22.
	FFI_USE_M41 bool // for M41.
	FFI_USE_M42 bool // for M42.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMMatrix2DInit with all fields set.
func (p DOMMatrix2DInit) FromRef(ref js.Ref) DOMMatrix2DInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMMatrix2DInit in the application heap.
func (p DOMMatrix2DInit) New() js.Ref {
	return bindings.DOMMatrix2DInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DOMMatrix2DInit) UpdateFrom(ref js.Ref) {
	bindings.DOMMatrix2DInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DOMMatrix2DInit) Update(ref js.Ref) {
	bindings.DOMMatrix2DInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SVGTransform struct {
	ref js.Ref
}

func (this SVGTransform) Once() SVGTransform {
	this.Ref().Once()
	return this
}

func (this SVGTransform) Ref() js.Ref {
	return this.ref
}

func (this SVGTransform) FromRef(ref js.Ref) SVGTransform {
	this.ref = ref
	return this
}

func (this SVGTransform) Free() {
	this.Ref().Free()
}

// Type returns the value of property "SVGTransform.type".
//
// The returned bool will be false if there is no such property.
func (this SVGTransform) Type() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGTransformType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Matrix returns the value of property "SVGTransform.matrix".
//
// The returned bool will be false if there is no such property.
func (this SVGTransform) Matrix() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.GetSVGTransformMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMMatrix{}.FromRef(_ret), _ok
}

// Angle returns the value of property "SVGTransform.angle".
//
// The returned bool will be false if there is no such property.
func (this SVGTransform) Angle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGTransformAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// SetMatrix calls the method "SVGTransform.setMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetMatrix(matrix DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetMatrix(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&matrix),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetMatrixFunc returns the method "SVGTransform.setMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.SVGTransformSetMatrixFunc(
			this.Ref(),
		),
	)
}

// SetMatrix1 calls the method "SVGTransform.setMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetMatrix1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetMatrix1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetMatrix1Func returns the method "SVGTransform.setMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetMatrix1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGTransformSetMatrix1Func(
			this.Ref(),
		),
	)
}

// SetTranslate calls the method "SVGTransform.setTranslate".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetTranslate(tx float32, ty float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetTranslate(
		this.Ref(), js.Pointer(&_ok),
		float32(tx),
		float32(ty),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTranslateFunc returns the method "SVGTransform.setTranslate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetTranslateFunc() (fn js.Func[func(tx float32, ty float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetTranslateFunc(
			this.Ref(),
		),
	)
}

// SetScale calls the method "SVGTransform.setScale".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetScale(sx float32, sy float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetScale(
		this.Ref(), js.Pointer(&_ok),
		float32(sx),
		float32(sy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetScaleFunc returns the method "SVGTransform.setScale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetScaleFunc() (fn js.Func[func(sx float32, sy float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetScaleFunc(
			this.Ref(),
		),
	)
}

// SetRotate calls the method "SVGTransform.setRotate".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetRotate(angle float32, cx float32, cy float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetRotate(
		this.Ref(), js.Pointer(&_ok),
		float32(angle),
		float32(cx),
		float32(cy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRotateFunc returns the method "SVGTransform.setRotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetRotateFunc() (fn js.Func[func(angle float32, cx float32, cy float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetRotateFunc(
			this.Ref(),
		),
	)
}

// SetSkewX calls the method "SVGTransform.setSkewX".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetSkewX(angle float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetSkewX(
		this.Ref(), js.Pointer(&_ok),
		float32(angle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSkewXFunc returns the method "SVGTransform.setSkewX".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetSkewXFunc() (fn js.Func[func(angle float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetSkewXFunc(
			this.Ref(),
		),
	)
}

// SetSkewY calls the method "SVGTransform.setSkewY".
//
// The returned bool will be false if there is no such method.
func (this SVGTransform) SetSkewY(angle float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformSetSkewY(
		this.Ref(), js.Pointer(&_ok),
		float32(angle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSkewYFunc returns the method "SVGTransform.setSkewY".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransform) SetSkewYFunc() (fn js.Func[func(angle float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetSkewYFunc(
			this.Ref(),
		),
	)
}

func NewDOMPointReadOnly(x float64, y float64, z float64, w float64) DOMPointReadOnly {
	return DOMPointReadOnly{}.FromRef(
		bindings.NewDOMPointReadOnlyByDOMPointReadOnly(
			float64(x),
			float64(y),
			float64(z),
			float64(w)),
	)
}

func NewDOMPointReadOnlyByDOMPointReadOnly1(x float64, y float64, z float64) DOMPointReadOnly {
	return DOMPointReadOnly{}.FromRef(
		bindings.NewDOMPointReadOnlyByDOMPointReadOnly1(
			float64(x),
			float64(y),
			float64(z)),
	)
}

func NewDOMPointReadOnlyByDOMPointReadOnly2(x float64, y float64) DOMPointReadOnly {
	return DOMPointReadOnly{}.FromRef(
		bindings.NewDOMPointReadOnlyByDOMPointReadOnly2(
			float64(x),
			float64(y)),
	)
}

func NewDOMPointReadOnlyByDOMPointReadOnly3(x float64) DOMPointReadOnly {
	return DOMPointReadOnly{}.FromRef(
		bindings.NewDOMPointReadOnlyByDOMPointReadOnly3(
			float64(x)),
	)
}

func NewDOMPointReadOnlyByDOMPointReadOnly4() DOMPointReadOnly {
	return DOMPointReadOnly{}.FromRef(
		bindings.NewDOMPointReadOnlyByDOMPointReadOnly4(),
	)
}

type DOMPointReadOnly struct {
	ref js.Ref
}

func (this DOMPointReadOnly) Once() DOMPointReadOnly {
	this.Ref().Once()
	return this
}

func (this DOMPointReadOnly) Ref() js.Ref {
	return this.ref
}

func (this DOMPointReadOnly) FromRef(ref js.Ref) DOMPointReadOnly {
	this.ref = ref
	return this
}

func (this DOMPointReadOnly) Free() {
	this.Ref().Free()
}

// X returns the value of property "DOMPointReadOnly.x".
//
// The returned bool will be false if there is no such property.
func (this DOMPointReadOnly) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointReadOnlyX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "DOMPointReadOnly.y".
//
// The returned bool will be false if there is no such property.
func (this DOMPointReadOnly) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointReadOnlyY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "DOMPointReadOnly.z".
//
// The returned bool will be false if there is no such property.
func (this DOMPointReadOnly) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointReadOnlyZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// W returns the value of property "DOMPointReadOnly.w".
//
// The returned bool will be false if there is no such property.
func (this DOMPointReadOnly) W() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointReadOnlyW(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FromPoint calls the staticmethod "DOMPointReadOnly.fromPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMPointReadOnly) FromPoint(other DOMPointInit) (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointReadOnlyFromPoint(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// FromPointFunc returns the staticmethod "DOMPointReadOnly.fromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPointReadOnly) FromPointFunc() (fn js.Func[func(other DOMPointInit) DOMPointReadOnly]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyFromPointFunc(
			this.Ref(),
		),
	)
}

// FromPoint1 calls the staticmethod "DOMPointReadOnly.fromPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMPointReadOnly) FromPoint1() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointReadOnlyFromPoint1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// FromPoint1Func returns the staticmethod "DOMPointReadOnly.fromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPointReadOnly) FromPoint1Func() (fn js.Func[func() DOMPointReadOnly]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyFromPoint1Func(
			this.Ref(),
		),
	)
}

// MatrixTransform calls the method "DOMPointReadOnly.matrixTransform".
//
// The returned bool will be false if there is no such method.
func (this DOMPointReadOnly) MatrixTransform(matrix DOMMatrixInit) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointReadOnlyMatrixTransform(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&matrix),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// MatrixTransformFunc returns the method "DOMPointReadOnly.matrixTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPointReadOnly) MatrixTransformFunc() (fn js.Func[func(matrix DOMMatrixInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyMatrixTransformFunc(
			this.Ref(),
		),
	)
}

// MatrixTransform1 calls the method "DOMPointReadOnly.matrixTransform".
//
// The returned bool will be false if there is no such method.
func (this DOMPointReadOnly) MatrixTransform1() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointReadOnlyMatrixTransform1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// MatrixTransform1Func returns the method "DOMPointReadOnly.matrixTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPointReadOnly) MatrixTransform1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyMatrixTransform1Func(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMPointReadOnly.toJSON".
//
// The returned bool will be false if there is no such method.
func (this DOMPointReadOnly) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointReadOnlyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "DOMPointReadOnly.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPointReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
}

type SVGAnimatedRect struct {
	ref js.Ref
}

func (this SVGAnimatedRect) Once() SVGAnimatedRect {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedRect) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedRect) FromRef(ref js.Ref) SVGAnimatedRect {
	this.ref = ref
	return this
}

func (this SVGAnimatedRect) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedRect.baseVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedRect) BaseVal() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedRectBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRect{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedRect.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedRect) AnimVal() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedRectAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

const (
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_UNKNOWN  uint16 = 0
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_NONE     uint16 = 1
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMINYMIN uint16 = 2
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMIDYMIN uint16 = 3
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMAXYMIN uint16 = 4
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMINYMID uint16 = 5
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMIDYMID uint16 = 6
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMAXYMID uint16 = 7
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMINYMAX uint16 = 8
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMIDYMAX uint16 = 9
	SVGPreserveAspectRatio_SVG_PRESERVEASPECTRATIO_XMAXYMAX uint16 = 10
	SVGPreserveAspectRatio_SVG_MEETORSLICE_UNKNOWN          uint16 = 0
	SVGPreserveAspectRatio_SVG_MEETORSLICE_MEET             uint16 = 1
	SVGPreserveAspectRatio_SVG_MEETORSLICE_SLICE            uint16 = 2
)

type SVGPreserveAspectRatio struct {
	ref js.Ref
}

func (this SVGPreserveAspectRatio) Once() SVGPreserveAspectRatio {
	this.Ref().Once()
	return this
}

func (this SVGPreserveAspectRatio) Ref() js.Ref {
	return this.ref
}

func (this SVGPreserveAspectRatio) FromRef(ref js.Ref) SVGPreserveAspectRatio {
	this.ref = ref
	return this
}

func (this SVGPreserveAspectRatio) Free() {
	this.Ref().Free()
}

// Align returns the value of property "SVGPreserveAspectRatio.align".
//
// The returned bool will be false if there is no such property.
func (this SVGPreserveAspectRatio) Align() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGPreserveAspectRatioAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Align sets the value of property "SVGPreserveAspectRatio.align" to val.
//
// It returns false if the property cannot be set.
func (this SVGPreserveAspectRatio) SetAlign(val uint16) bool {
	return js.True == bindings.SetSVGPreserveAspectRatioAlign(
		this.Ref(),
		uint32(val),
	)
}

// MeetOrSlice returns the value of property "SVGPreserveAspectRatio.meetOrSlice".
//
// The returned bool will be false if there is no such property.
func (this SVGPreserveAspectRatio) MeetOrSlice() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGPreserveAspectRatioMeetOrSlice(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// MeetOrSlice sets the value of property "SVGPreserveAspectRatio.meetOrSlice" to val.
//
// It returns false if the property cannot be set.
func (this SVGPreserveAspectRatio) SetMeetOrSlice(val uint16) bool {
	return js.True == bindings.SetSVGPreserveAspectRatioMeetOrSlice(
		this.Ref(),
		uint32(val),
	)
}

type SVGAnimatedPreserveAspectRatio struct {
	ref js.Ref
}

func (this SVGAnimatedPreserveAspectRatio) Once() SVGAnimatedPreserveAspectRatio {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedPreserveAspectRatio) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedPreserveAspectRatio) FromRef(ref js.Ref) SVGAnimatedPreserveAspectRatio {
	this.ref = ref
	return this
}

func (this SVGAnimatedPreserveAspectRatio) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedPreserveAspectRatio.baseVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) BaseVal() (SVGPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedPreserveAspectRatioBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPreserveAspectRatio{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedPreserveAspectRatio.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) AnimVal() (SVGPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedPreserveAspectRatioAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPreserveAspectRatio{}.FromRef(_ret), _ok
}

type OnBeforeUnloadEventHandlerNonNullFunc func(this js.Ref, event Event) js.Ref

func (fn OnBeforeUnloadEventHandlerNonNullFunc) Register() js.Func[func(event Event) js.String] {
	return js.RegisterCallback[func(event Event) js.String](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeforeUnloadEventHandlerNonNullFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnBeforeUnloadEventHandlerNonNull[T any] struct {
	Fn  func(arg T, this js.Ref, event Event) js.Ref
	Arg T
}

func (cb *OnBeforeUnloadEventHandlerNonNull[T]) Register() js.Func[func(event Event) js.String] {
	return js.RegisterCallback[func(event Event) js.String](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeforeUnloadEventHandlerNonNull[T]) DispatchCallback(
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

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnBeforeUnloadEventHandler = js.Func[func(event Event) js.String]

type SVGSVGElement struct {
	SVGGraphicsElement
}

func (this SVGSVGElement) Once() SVGSVGElement {
	this.Ref().Once()
	return this
}

func (this SVGSVGElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGSVGElement) FromRef(ref js.Ref) SVGSVGElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGSVGElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGSVGElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGSVGElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGSVGElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGSVGElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// CurrentScale returns the value of property "SVGSVGElement.currentScale".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) CurrentScale() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementCurrentScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// CurrentScale sets the value of property "SVGSVGElement.currentScale" to val.
//
// It returns false if the property cannot be set.
func (this SVGSVGElement) SetCurrentScale(val float32) bool {
	return js.True == bindings.SetSVGSVGElementCurrentScale(
		this.Ref(),
		float32(val),
	)
}

// CurrentTranslate returns the value of property "SVGSVGElement.currentTranslate".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) CurrentTranslate() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementCurrentTranslate(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// ViewBox returns the value of property "SVGSVGElement.viewBox".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) ViewBox() (SVGAnimatedRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementViewBox(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedRect{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGSVGElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGSVGElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGSVGElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

// GetIntersectionList calls the method "SVGSVGElement.getIntersectionList".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) GetIntersectionList(rect DOMRectReadOnly, referenceElement SVGElement) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementGetIntersectionList(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// GetIntersectionListFunc returns the method "SVGSVGElement.getIntersectionList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) GetIntersectionListFunc() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetIntersectionListFunc(
			this.Ref(),
		),
	)
}

// GetEnclosureList calls the method "SVGSVGElement.getEnclosureList".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) GetEnclosureList(rect DOMRectReadOnly, referenceElement SVGElement) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementGetEnclosureList(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// GetEnclosureListFunc returns the method "SVGSVGElement.getEnclosureList".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) GetEnclosureListFunc() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetEnclosureListFunc(
			this.Ref(),
		),
	)
}

// CheckIntersection calls the method "SVGSVGElement.checkIntersection".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CheckIntersection(element SVGElement, rect DOMRectReadOnly) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCheckIntersection(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
		rect.Ref(),
	)

	return _ret == js.True, _ok
}

// CheckIntersectionFunc returns the method "SVGSVGElement.checkIntersection".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CheckIntersectionFunc() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementCheckIntersectionFunc(
			this.Ref(),
		),
	)
}

// CheckEnclosure calls the method "SVGSVGElement.checkEnclosure".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CheckEnclosure(element SVGElement, rect DOMRectReadOnly) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCheckEnclosure(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
		rect.Ref(),
	)

	return _ret == js.True, _ok
}

// CheckEnclosureFunc returns the method "SVGSVGElement.checkEnclosure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CheckEnclosureFunc() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementCheckEnclosureFunc(
			this.Ref(),
		),
	)
}

// DeselectAll calls the method "SVGSVGElement.deselectAll".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) DeselectAll() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementDeselectAll(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeselectAllFunc returns the method "SVGSVGElement.deselectAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) DeselectAllFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementDeselectAllFunc(
			this.Ref(),
		),
	)
}

// CreateSVGNumber calls the method "SVGSVGElement.createSVGNumber".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGNumber() (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGNumber(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// CreateSVGNumberFunc returns the method "SVGSVGElement.createSVGNumber".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGNumberFunc() (fn js.Func[func() SVGNumber]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGNumberFunc(
			this.Ref(),
		),
	)
}

// CreateSVGLength calls the method "SVGSVGElement.createSVGLength".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGLength() (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGLength(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// CreateSVGLengthFunc returns the method "SVGSVGElement.createSVGLength".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGLengthFunc() (fn js.Func[func() SVGLength]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGLengthFunc(
			this.Ref(),
		),
	)
}

// CreateSVGAngle calls the method "SVGSVGElement.createSVGAngle".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGAngle() (SVGAngle, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGAngle(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGAngle{}.FromRef(_ret), _ok
}

// CreateSVGAngleFunc returns the method "SVGSVGElement.createSVGAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGAngleFunc() (fn js.Func[func() SVGAngle]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGAngleFunc(
			this.Ref(),
		),
	)
}

// CreateSVGPoint calls the method "SVGSVGElement.createSVGPoint".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGPoint() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGPoint(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// CreateSVGPointFunc returns the method "SVGSVGElement.createSVGPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGPointFunc() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGPointFunc(
			this.Ref(),
		),
	)
}

// CreateSVGMatrix calls the method "SVGSVGElement.createSVGMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGMatrix() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGMatrix(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// CreateSVGMatrixFunc returns the method "SVGSVGElement.createSVGMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGRect calls the method "SVGSVGElement.createSVGRect".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGRect() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGRect(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// CreateSVGRectFunc returns the method "SVGSVGElement.createSVGRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGRectFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransform calls the method "SVGSVGElement.createSVGTransform".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGTransform() (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// CreateSVGTransformFunc returns the method "SVGSVGElement.createSVGTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGTransformFunc() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix calls the method "SVGSVGElement.createSVGTransformFromMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&matrix),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// CreateSVGTransformFromMatrixFunc returns the method "SVGSVGElement.createSVGTransformFromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGTransformFromMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFromMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix1 calls the method "SVGSVGElement.createSVGTransformFromMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) CreateSVGTransformFromMatrix1() (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// CreateSVGTransformFromMatrix1Func returns the method "SVGSVGElement.createSVGTransformFromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) CreateSVGTransformFromMatrix1Func() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFromMatrix1Func(
			this.Ref(),
		),
	)
}

// GetElementById calls the method "SVGSVGElement.getElementById".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) GetElementById(elementId js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementGetElementById(
		this.Ref(), js.Pointer(&_ok),
		elementId.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// GetElementByIdFunc returns the method "SVGSVGElement.getElementById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// SuspendRedraw calls the method "SVGSVGElement.suspendRedraw".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint32) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementSuspendRedraw(
		this.Ref(), js.Pointer(&_ok),
		uint32(maxWaitMilliseconds),
	)

	return uint32(_ret), _ok
}

// SuspendRedrawFunc returns the method "SVGSVGElement.suspendRedraw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) SuspendRedrawFunc() (fn js.Func[func(maxWaitMilliseconds uint32) uint32]) {
	return fn.FromRef(
		bindings.SVGSVGElementSuspendRedrawFunc(
			this.Ref(),
		),
	)
}

// UnsuspendRedraw calls the method "SVGSVGElement.unsuspendRedraw".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) UnsuspendRedraw(suspendHandleID uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementUnsuspendRedraw(
		this.Ref(), js.Pointer(&_ok),
		uint32(suspendHandleID),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnsuspendRedrawFunc returns the method "SVGSVGElement.unsuspendRedraw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) UnsuspendRedrawFunc() (fn js.Func[func(suspendHandleID uint32)]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnsuspendRedrawFunc(
			this.Ref(),
		),
	)
}

// UnsuspendRedrawAll calls the method "SVGSVGElement.unsuspendRedrawAll".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) UnsuspendRedrawAll() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementUnsuspendRedrawAll(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnsuspendRedrawAllFunc returns the method "SVGSVGElement.unsuspendRedrawAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) UnsuspendRedrawAllFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnsuspendRedrawAllFunc(
			this.Ref(),
		),
	)
}

// ForceRedraw calls the method "SVGSVGElement.forceRedraw".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) ForceRedraw() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementForceRedraw(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ForceRedrawFunc returns the method "SVGSVGElement.forceRedraw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) ForceRedrawFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementForceRedrawFunc(
			this.Ref(),
		),
	)
}

// PauseAnimations calls the method "SVGSVGElement.pauseAnimations".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) PauseAnimations() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementPauseAnimations(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseAnimationsFunc returns the method "SVGSVGElement.pauseAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) PauseAnimationsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementPauseAnimationsFunc(
			this.Ref(),
		),
	)
}

// UnpauseAnimations calls the method "SVGSVGElement.unpauseAnimations".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) UnpauseAnimations() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementUnpauseAnimations(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnpauseAnimationsFunc returns the method "SVGSVGElement.unpauseAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) UnpauseAnimationsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnpauseAnimationsFunc(
			this.Ref(),
		),
	)
}

// AnimationsPaused calls the method "SVGSVGElement.animationsPaused".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) AnimationsPaused() (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementAnimationsPaused(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// AnimationsPausedFunc returns the method "SVGSVGElement.animationsPaused".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) AnimationsPausedFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementAnimationsPausedFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTime calls the method "SVGSVGElement.getCurrentTime".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) GetCurrentTime() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementGetCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetCurrentTimeFunc returns the method "SVGSVGElement.getCurrentTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) GetCurrentTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

// SetCurrentTime calls the method "SVGSVGElement.setCurrentTime".
//
// The returned bool will be false if there is no such method.
func (this SVGSVGElement) SetCurrentTime(seconds float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGSVGElementSetCurrentTime(
		this.Ref(), js.Pointer(&_ok),
		float32(seconds),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCurrentTimeFunc returns the method "SVGSVGElement.setCurrentTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGSVGElement) SetCurrentTimeFunc() (fn js.Func[func(seconds float32)]) {
	return fn.FromRef(
		bindings.SVGSVGElementSetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

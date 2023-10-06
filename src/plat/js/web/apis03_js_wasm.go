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
// It returns ok=false if there is no such property.
func (this Baseline) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBaselineName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "Baseline.value".
//
// It returns ok=false if there is no such property.
func (this Baseline) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetBaselineValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this Font) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// GlyphsRendered returns the value of property "Font.glyphsRendered".
//
// It returns ok=false if there is no such property.
func (this Font) GlyphsRendered() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontGlyphsRendered(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this FontMetrics) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Advances returns the value of property "FontMetrics.advances".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Advances() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetFontMetricsAdvances(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BoundingBoxLeft returns the value of property "FontMetrics.boundingBoxLeft".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BoundingBoxRight returns the value of property "FontMetrics.boundingBoxRight".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxRight() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxRight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "FontMetrics.height".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EmHeightAscent returns the value of property "FontMetrics.emHeightAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) EmHeightAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsEmHeightAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EmHeightDescent returns the value of property "FontMetrics.emHeightDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) EmHeightDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsEmHeightDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BoundingBoxAscent returns the value of property "FontMetrics.boundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BoundingBoxDescent returns the value of property "FontMetrics.boundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxAscent returns the value of property "FontMetrics.fontBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) FontBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsFontBoundingBoxAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxDescent returns the value of property "FontMetrics.fontBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) FontBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsFontBoundingBoxDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DominantBaseline returns the value of property "FontMetrics.dominantBaseline".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) DominantBaseline() (ret Baseline, ok bool) {
	ok = js.True == bindings.GetFontMetricsDominantBaseline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Baselines returns the value of property "FontMetrics.baselines".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Baselines() (ret js.FrozenArray[Baseline], ok bool) {
	ok = js.True == bindings.GetFontMetricsBaselines(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "FontMetrics.fonts".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Fonts() (ret js.FrozenArray[Font], ok bool) {
	ok = js.True == bindings.GetFontMetricsFonts(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewStaticRange(init StaticRangeInit) (ret StaticRange) {
	ret.ref = bindings.NewStaticRangeByStaticRange(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this Selection) AnchorNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetSelectionAnchorNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnchorOffset returns the value of property "Selection.anchorOffset".
//
// It returns ok=false if there is no such property.
func (this Selection) AnchorOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionAnchorOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FocusNode returns the value of property "Selection.focusNode".
//
// It returns ok=false if there is no such property.
func (this Selection) FocusNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetSelectionFocusNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FocusOffset returns the value of property "Selection.focusOffset".
//
// It returns ok=false if there is no such property.
func (this Selection) FocusOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionFocusOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsCollapsed returns the value of property "Selection.isCollapsed".
//
// It returns ok=false if there is no such property.
func (this Selection) IsCollapsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetSelectionIsCollapsed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RangeCount returns the value of property "Selection.rangeCount".
//
// It returns ok=false if there is no such property.
func (this Selection) RangeCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionRangeCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Selection.type".
//
// It returns ok=false if there is no such property.
func (this Selection) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSelectionType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "Selection.direction".
//
// It returns ok=false if there is no such property.
func (this Selection) Direction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSelectionDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetRangeAt returns true if the method "Selection.getRangeAt" exists.
func (this Selection) HasGetRangeAt() bool {
	return js.True == bindings.HasSelectionGetRangeAt(
		this.Ref(),
	)
}

// GetRangeAtFunc returns the method "Selection.getRangeAt".
func (this Selection) GetRangeAtFunc() (fn js.Func[func(index uint32) Range]) {
	return fn.FromRef(
		bindings.SelectionGetRangeAtFunc(
			this.Ref(),
		),
	)
}

// GetRangeAt calls the method "Selection.getRangeAt".
func (this Selection) GetRangeAt(index uint32) (ret Range) {
	bindings.CallSelectionGetRangeAt(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetRangeAt calls the method "Selection.getRangeAt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryGetRangeAt(index uint32) (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionGetRangeAt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAddRange returns true if the method "Selection.addRange" exists.
func (this Selection) HasAddRange() bool {
	return js.True == bindings.HasSelectionAddRange(
		this.Ref(),
	)
}

// AddRangeFunc returns the method "Selection.addRange".
func (this Selection) AddRangeFunc() (fn js.Func[func(rng Range)]) {
	return fn.FromRef(
		bindings.SelectionAddRangeFunc(
			this.Ref(),
		),
	)
}

// AddRange calls the method "Selection.addRange".
func (this Selection) AddRange(rng Range) (ret js.Void) {
	bindings.CallSelectionAddRange(
		this.Ref(), js.Pointer(&ret),
		rng.Ref(),
	)

	return
}

// TryAddRange calls the method "Selection.addRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryAddRange(rng Range) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionAddRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rng.Ref(),
	)

	return
}

// HasRemoveRange returns true if the method "Selection.removeRange" exists.
func (this Selection) HasRemoveRange() bool {
	return js.True == bindings.HasSelectionRemoveRange(
		this.Ref(),
	)
}

// RemoveRangeFunc returns the method "Selection.removeRange".
func (this Selection) RemoveRangeFunc() (fn js.Func[func(rng Range)]) {
	return fn.FromRef(
		bindings.SelectionRemoveRangeFunc(
			this.Ref(),
		),
	)
}

// RemoveRange calls the method "Selection.removeRange".
func (this Selection) RemoveRange(rng Range) (ret js.Void) {
	bindings.CallSelectionRemoveRange(
		this.Ref(), js.Pointer(&ret),
		rng.Ref(),
	)

	return
}

// TryRemoveRange calls the method "Selection.removeRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryRemoveRange(rng Range) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionRemoveRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rng.Ref(),
	)

	return
}

// HasRemoveAllRanges returns true if the method "Selection.removeAllRanges" exists.
func (this Selection) HasRemoveAllRanges() bool {
	return js.True == bindings.HasSelectionRemoveAllRanges(
		this.Ref(),
	)
}

// RemoveAllRangesFunc returns the method "Selection.removeAllRanges".
func (this Selection) RemoveAllRangesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionRemoveAllRangesFunc(
			this.Ref(),
		),
	)
}

// RemoveAllRanges calls the method "Selection.removeAllRanges".
func (this Selection) RemoveAllRanges() (ret js.Void) {
	bindings.CallSelectionRemoveAllRanges(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemoveAllRanges calls the method "Selection.removeAllRanges"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryRemoveAllRanges() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionRemoveAllRanges(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEmpty returns true if the method "Selection.empty" exists.
func (this Selection) HasEmpty() bool {
	return js.True == bindings.HasSelectionEmpty(
		this.Ref(),
	)
}

// EmptyFunc returns the method "Selection.empty".
func (this Selection) EmptyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionEmptyFunc(
			this.Ref(),
		),
	)
}

// Empty calls the method "Selection.empty".
func (this Selection) Empty() (ret js.Void) {
	bindings.CallSelectionEmpty(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEmpty calls the method "Selection.empty"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryEmpty() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionEmpty(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetComposedRanges returns true if the method "Selection.getComposedRanges" exists.
func (this Selection) HasGetComposedRanges() bool {
	return js.True == bindings.HasSelectionGetComposedRanges(
		this.Ref(),
	)
}

// GetComposedRangesFunc returns the method "Selection.getComposedRanges".
func (this Selection) GetComposedRangesFunc() (fn js.Func[func(shadowRoots ...ShadowRoot) js.Array[StaticRange]]) {
	return fn.FromRef(
		bindings.SelectionGetComposedRangesFunc(
			this.Ref(),
		),
	)
}

// GetComposedRanges calls the method "Selection.getComposedRanges".
func (this Selection) GetComposedRanges(shadowRoots ...ShadowRoot) (ret js.Array[StaticRange]) {
	bindings.CallSelectionGetComposedRanges(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(shadowRoots),
		js.SizeU(len(shadowRoots)),
	)

	return
}

// TryGetComposedRanges calls the method "Selection.getComposedRanges"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryGetComposedRanges(shadowRoots ...ShadowRoot) (ret js.Array[StaticRange], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionGetComposedRanges(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(shadowRoots),
		js.SizeU(len(shadowRoots)),
	)

	return
}

// HasCollapse returns true if the method "Selection.collapse" exists.
func (this Selection) HasCollapse() bool {
	return js.True == bindings.HasSelectionCollapse(
		this.Ref(),
	)
}

// CollapseFunc returns the method "Selection.collapse".
func (this Selection) CollapseFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionCollapseFunc(
			this.Ref(),
		),
	)
}

// Collapse calls the method "Selection.collapse".
func (this Selection) Collapse(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionCollapse(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryCollapse calls the method "Selection.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryCollapse(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasCollapse1 returns true if the method "Selection.collapse" exists.
func (this Selection) HasCollapse1() bool {
	return js.True == bindings.HasSelectionCollapse1(
		this.Ref(),
	)
}

// Collapse1Func returns the method "Selection.collapse".
func (this Selection) Collapse1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionCollapse1Func(
			this.Ref(),
		),
	)
}

// Collapse1 calls the method "Selection.collapse".
func (this Selection) Collapse1(node Node) (ret js.Void) {
	bindings.CallSelectionCollapse1(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryCollapse1 calls the method "Selection.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryCollapse1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapse1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSetPosition returns true if the method "Selection.setPosition" exists.
func (this Selection) HasSetPosition() bool {
	return js.True == bindings.HasSelectionSetPosition(
		this.Ref(),
	)
}

// SetPositionFunc returns the method "Selection.setPosition".
func (this Selection) SetPositionFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionSetPositionFunc(
			this.Ref(),
		),
	)
}

// SetPosition calls the method "Selection.setPosition".
func (this Selection) SetPosition(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionSetPosition(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetPosition calls the method "Selection.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TrySetPosition(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetPosition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasSetPosition1 returns true if the method "Selection.setPosition" exists.
func (this Selection) HasSetPosition1() bool {
	return js.True == bindings.HasSelectionSetPosition1(
		this.Ref(),
	)
}

// SetPosition1Func returns the method "Selection.setPosition".
func (this Selection) SetPosition1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionSetPosition1Func(
			this.Ref(),
		),
	)
}

// SetPosition1 calls the method "Selection.setPosition".
func (this Selection) SetPosition1(node Node) (ret js.Void) {
	bindings.CallSelectionSetPosition1(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetPosition1 calls the method "Selection.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TrySetPosition1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetPosition1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasCollapseToStart returns true if the method "Selection.collapseToStart" exists.
func (this Selection) HasCollapseToStart() bool {
	return js.True == bindings.HasSelectionCollapseToStart(
		this.Ref(),
	)
}

// CollapseToStartFunc returns the method "Selection.collapseToStart".
func (this Selection) CollapseToStartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionCollapseToStartFunc(
			this.Ref(),
		),
	)
}

// CollapseToStart calls the method "Selection.collapseToStart".
func (this Selection) CollapseToStart() (ret js.Void) {
	bindings.CallSelectionCollapseToStart(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCollapseToStart calls the method "Selection.collapseToStart"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryCollapseToStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapseToStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCollapseToEnd returns true if the method "Selection.collapseToEnd" exists.
func (this Selection) HasCollapseToEnd() bool {
	return js.True == bindings.HasSelectionCollapseToEnd(
		this.Ref(),
	)
}

// CollapseToEndFunc returns the method "Selection.collapseToEnd".
func (this Selection) CollapseToEndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionCollapseToEndFunc(
			this.Ref(),
		),
	)
}

// CollapseToEnd calls the method "Selection.collapseToEnd".
func (this Selection) CollapseToEnd() (ret js.Void) {
	bindings.CallSelectionCollapseToEnd(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCollapseToEnd calls the method "Selection.collapseToEnd"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryCollapseToEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapseToEnd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExtend returns true if the method "Selection.extend" exists.
func (this Selection) HasExtend() bool {
	return js.True == bindings.HasSelectionExtend(
		this.Ref(),
	)
}

// ExtendFunc returns the method "Selection.extend".
func (this Selection) ExtendFunc() (fn js.Func[func(node Node, offset uint32)]) {
	return fn.FromRef(
		bindings.SelectionExtendFunc(
			this.Ref(),
		),
	)
}

// Extend calls the method "Selection.extend".
func (this Selection) Extend(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionExtend(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryExtend calls the method "Selection.extend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryExtend(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionExtend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasExtend1 returns true if the method "Selection.extend" exists.
func (this Selection) HasExtend1() bool {
	return js.True == bindings.HasSelectionExtend1(
		this.Ref(),
	)
}

// Extend1Func returns the method "Selection.extend".
func (this Selection) Extend1Func() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionExtend1Func(
			this.Ref(),
		),
	)
}

// Extend1 calls the method "Selection.extend".
func (this Selection) Extend1(node Node) (ret js.Void) {
	bindings.CallSelectionExtend1(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryExtend1 calls the method "Selection.extend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryExtend1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionExtend1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasSetBaseAndExtent returns true if the method "Selection.setBaseAndExtent" exists.
func (this Selection) HasSetBaseAndExtent() bool {
	return js.True == bindings.HasSelectionSetBaseAndExtent(
		this.Ref(),
	)
}

// SetBaseAndExtentFunc returns the method "Selection.setBaseAndExtent".
func (this Selection) SetBaseAndExtentFunc() (fn js.Func[func(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32)]) {
	return fn.FromRef(
		bindings.SelectionSetBaseAndExtentFunc(
			this.Ref(),
		),
	)
}

// SetBaseAndExtent calls the method "Selection.setBaseAndExtent".
func (this Selection) SetBaseAndExtent(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32) (ret js.Void) {
	bindings.CallSelectionSetBaseAndExtent(
		this.Ref(), js.Pointer(&ret),
		anchorNode.Ref(),
		uint32(anchorOffset),
		focusNode.Ref(),
		uint32(focusOffset),
	)

	return
}

// TrySetBaseAndExtent calls the method "Selection.setBaseAndExtent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TrySetBaseAndExtent(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetBaseAndExtent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		anchorNode.Ref(),
		uint32(anchorOffset),
		focusNode.Ref(),
		uint32(focusOffset),
	)

	return
}

// HasSelectAllChildren returns true if the method "Selection.selectAllChildren" exists.
func (this Selection) HasSelectAllChildren() bool {
	return js.True == bindings.HasSelectionSelectAllChildren(
		this.Ref(),
	)
}

// SelectAllChildrenFunc returns the method "Selection.selectAllChildren".
func (this Selection) SelectAllChildrenFunc() (fn js.Func[func(node Node)]) {
	return fn.FromRef(
		bindings.SelectionSelectAllChildrenFunc(
			this.Ref(),
		),
	)
}

// SelectAllChildren calls the method "Selection.selectAllChildren".
func (this Selection) SelectAllChildren(node Node) (ret js.Void) {
	bindings.CallSelectionSelectAllChildren(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectAllChildren calls the method "Selection.selectAllChildren"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TrySelectAllChildren(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSelectAllChildren(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasModify returns true if the method "Selection.modify" exists.
func (this Selection) HasModify() bool {
	return js.True == bindings.HasSelectionModify(
		this.Ref(),
	)
}

// ModifyFunc returns the method "Selection.modify".
func (this Selection) ModifyFunc() (fn js.Func[func(alter js.String, direction js.String, granularity js.String)]) {
	return fn.FromRef(
		bindings.SelectionModifyFunc(
			this.Ref(),
		),
	)
}

// Modify calls the method "Selection.modify".
func (this Selection) Modify(alter js.String, direction js.String, granularity js.String) (ret js.Void) {
	bindings.CallSelectionModify(
		this.Ref(), js.Pointer(&ret),
		alter.Ref(),
		direction.Ref(),
		granularity.Ref(),
	)

	return
}

// TryModify calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryModify(alter js.String, direction js.String, granularity js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
		direction.Ref(),
		granularity.Ref(),
	)

	return
}

// HasModify1 returns true if the method "Selection.modify" exists.
func (this Selection) HasModify1() bool {
	return js.True == bindings.HasSelectionModify1(
		this.Ref(),
	)
}

// Modify1Func returns the method "Selection.modify".
func (this Selection) Modify1Func() (fn js.Func[func(alter js.String, direction js.String)]) {
	return fn.FromRef(
		bindings.SelectionModify1Func(
			this.Ref(),
		),
	)
}

// Modify1 calls the method "Selection.modify".
func (this Selection) Modify1(alter js.String, direction js.String) (ret js.Void) {
	bindings.CallSelectionModify1(
		this.Ref(), js.Pointer(&ret),
		alter.Ref(),
		direction.Ref(),
	)

	return
}

// TryModify1 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryModify1(alter js.String, direction js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
		direction.Ref(),
	)

	return
}

// HasModify2 returns true if the method "Selection.modify" exists.
func (this Selection) HasModify2() bool {
	return js.True == bindings.HasSelectionModify2(
		this.Ref(),
	)
}

// Modify2Func returns the method "Selection.modify".
func (this Selection) Modify2Func() (fn js.Func[func(alter js.String)]) {
	return fn.FromRef(
		bindings.SelectionModify2Func(
			this.Ref(),
		),
	)
}

// Modify2 calls the method "Selection.modify".
func (this Selection) Modify2(alter js.String) (ret js.Void) {
	bindings.CallSelectionModify2(
		this.Ref(), js.Pointer(&ret),
		alter.Ref(),
	)

	return
}

// TryModify2 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryModify2(alter js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
	)

	return
}

// HasModify3 returns true if the method "Selection.modify" exists.
func (this Selection) HasModify3() bool {
	return js.True == bindings.HasSelectionModify3(
		this.Ref(),
	)
}

// Modify3Func returns the method "Selection.modify".
func (this Selection) Modify3Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionModify3Func(
			this.Ref(),
		),
	)
}

// Modify3 calls the method "Selection.modify".
func (this Selection) Modify3() (ret js.Void) {
	bindings.CallSelectionModify3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryModify3 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryModify3() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteFromDocument returns true if the method "Selection.deleteFromDocument" exists.
func (this Selection) HasDeleteFromDocument() bool {
	return js.True == bindings.HasSelectionDeleteFromDocument(
		this.Ref(),
	)
}

// DeleteFromDocumentFunc returns the method "Selection.deleteFromDocument".
func (this Selection) DeleteFromDocumentFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SelectionDeleteFromDocumentFunc(
			this.Ref(),
		),
	)
}

// DeleteFromDocument calls the method "Selection.deleteFromDocument".
func (this Selection) DeleteFromDocument() (ret js.Void) {
	bindings.CallSelectionDeleteFromDocument(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeleteFromDocument calls the method "Selection.deleteFromDocument"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryDeleteFromDocument() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionDeleteFromDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasContainsNode returns true if the method "Selection.containsNode" exists.
func (this Selection) HasContainsNode() bool {
	return js.True == bindings.HasSelectionContainsNode(
		this.Ref(),
	)
}

// ContainsNodeFunc returns the method "Selection.containsNode".
func (this Selection) ContainsNodeFunc() (fn js.Func[func(node Node, allowPartialContainment bool) bool]) {
	return fn.FromRef(
		bindings.SelectionContainsNodeFunc(
			this.Ref(),
		),
	)
}

// ContainsNode calls the method "Selection.containsNode".
func (this Selection) ContainsNode(node Node, allowPartialContainment bool) (ret bool) {
	bindings.CallSelectionContainsNode(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
		js.Bool(bool(allowPartialContainment)),
	)

	return
}

// TryContainsNode calls the method "Selection.containsNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryContainsNode(node Node, allowPartialContainment bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionContainsNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		js.Bool(bool(allowPartialContainment)),
	)

	return
}

// HasContainsNode1 returns true if the method "Selection.containsNode" exists.
func (this Selection) HasContainsNode1() bool {
	return js.True == bindings.HasSelectionContainsNode1(
		this.Ref(),
	)
}

// ContainsNode1Func returns the method "Selection.containsNode".
func (this Selection) ContainsNode1Func() (fn js.Func[func(node Node) bool]) {
	return fn.FromRef(
		bindings.SelectionContainsNode1Func(
			this.Ref(),
		),
	)
}

// ContainsNode1 calls the method "Selection.containsNode".
func (this Selection) ContainsNode1(node Node) (ret bool) {
	bindings.CallSelectionContainsNode1(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryContainsNode1 calls the method "Selection.containsNode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryContainsNode1(node Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionContainsNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasToString returns true if the method "Selection.toString" exists.
func (this Selection) HasToString() bool {
	return js.True == bindings.HasSelectionToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "Selection.toString".
func (this Selection) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.SelectionToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "Selection.toString".
func (this Selection) ToString() (ret js.String) {
	bindings.CallSelectionToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "Selection.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Selection) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this CaretPosition) OffsetNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetCaretPositionOffsetNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Offset returns the value of property "CaretPosition.offset".
//
// It returns ok=false if there is no such property.
func (this CaretPosition) Offset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCaretPositionOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetClientRect returns true if the method "CaretPosition.getClientRect" exists.
func (this CaretPosition) HasGetClientRect() bool {
	return js.True == bindings.HasCaretPositionGetClientRect(
		this.Ref(),
	)
}

// GetClientRectFunc returns the method "CaretPosition.getClientRect".
func (this CaretPosition) GetClientRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.CaretPositionGetClientRectFunc(
			this.Ref(),
		),
	)
}

// GetClientRect calls the method "CaretPosition.getClientRect".
func (this CaretPosition) GetClientRect() (ret DOMRect) {
	bindings.CallCaretPositionGetClientRect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetClientRect calls the method "CaretPosition.getClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CaretPosition) TryGetClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaretPositionGetClientRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this XPathResult) ResultType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXPathResultResultType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberValue returns the value of property "XPathResult.numberValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) NumberValue() (ret float64, ok bool) {
	ok = js.True == bindings.GetXPathResultNumberValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StringValue returns the value of property "XPathResult.stringValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) StringValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXPathResultStringValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BooleanValue returns the value of property "XPathResult.booleanValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) BooleanValue() (ret bool, ok bool) {
	ok = js.True == bindings.GetXPathResultBooleanValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SingleNodeValue returns the value of property "XPathResult.singleNodeValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) SingleNodeValue() (ret Node, ok bool) {
	ok = js.True == bindings.GetXPathResultSingleNodeValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InvalidIteratorState returns the value of property "XPathResult.invalidIteratorState".
//
// It returns ok=false if there is no such property.
func (this XPathResult) InvalidIteratorState() (ret bool, ok bool) {
	ok = js.True == bindings.GetXPathResultInvalidIteratorState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SnapshotLength returns the value of property "XPathResult.snapshotLength".
//
// It returns ok=false if there is no such property.
func (this XPathResult) SnapshotLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXPathResultSnapshotLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasIterateNext returns true if the method "XPathResult.iterateNext" exists.
func (this XPathResult) HasIterateNext() bool {
	return js.True == bindings.HasXPathResultIterateNext(
		this.Ref(),
	)
}

// IterateNextFunc returns the method "XPathResult.iterateNext".
func (this XPathResult) IterateNextFunc() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.XPathResultIterateNextFunc(
			this.Ref(),
		),
	)
}

// IterateNext calls the method "XPathResult.iterateNext".
func (this XPathResult) IterateNext() (ret Node) {
	bindings.CallXPathResultIterateNext(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIterateNext calls the method "XPathResult.iterateNext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this XPathResult) TryIterateNext() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathResultIterateNext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSnapshotItem returns true if the method "XPathResult.snapshotItem" exists.
func (this XPathResult) HasSnapshotItem() bool {
	return js.True == bindings.HasXPathResultSnapshotItem(
		this.Ref(),
	)
}

// SnapshotItemFunc returns the method "XPathResult.snapshotItem".
func (this XPathResult) SnapshotItemFunc() (fn js.Func[func(index uint32) Node]) {
	return fn.FromRef(
		bindings.XPathResultSnapshotItemFunc(
			this.Ref(),
		),
	)
}

// SnapshotItem calls the method "XPathResult.snapshotItem".
func (this XPathResult) SnapshotItem(index uint32) (ret Node) {
	bindings.CallXPathResultSnapshotItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TrySnapshotItem calls the method "XPathResult.snapshotItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this XPathResult) TrySnapshotItem(index uint32) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathResultSnapshotItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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

// HasEvaluate returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasEvaluate() bool {
	return js.True == bindings.HasXPathExpressionEvaluate(
		this.Ref(),
	)
}

// EvaluateFunc returns the method "XPathExpression.evaluate".
func (this XPathExpression) EvaluateFunc() (fn js.Func[func(contextNode Node, typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate(contextNode Node, typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate(
		this.Ref(), js.Pointer(&ret),
		contextNode.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// TryEvaluate calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this XPathExpression) TryEvaluate(contextNode Node, typ uint16, result XPathResult) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextNode.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasEvaluate1 returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasEvaluate1() bool {
	return js.True == bindings.HasXPathExpressionEvaluate1(
		this.Ref(),
	)
}

// Evaluate1Func returns the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate1Func() (fn js.Func[func(contextNode Node, typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate1(contextNode Node, typ uint16) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate1(
		this.Ref(), js.Pointer(&ret),
		contextNode.Ref(),
		uint32(typ),
	)

	return
}

// TryEvaluate1 calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this XPathExpression) TryEvaluate1(contextNode Node, typ uint16) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextNode.Ref(),
		uint32(typ),
	)

	return
}

// HasEvaluate2 returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasEvaluate2() bool {
	return js.True == bindings.HasXPathExpressionEvaluate2(
		this.Ref(),
	)
}

// Evaluate2Func returns the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate2Func() (fn js.Func[func(contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.XPathExpressionEvaluate2Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate2(contextNode Node) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate2(
		this.Ref(), js.Pointer(&ret),
		contextNode.Ref(),
	)

	return
}

// TryEvaluate2 calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this XPathExpression) TryEvaluate2(contextNode Node) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextNode.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DocumentType) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PublicId returns the value of property "DocumentType.publicId".
//
// It returns ok=false if there is no such property.
func (this DocumentType) PublicId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypePublicId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SystemId returns the value of property "DocumentType.systemId".
//
// It returns ok=false if there is no such property.
func (this DocumentType) SystemId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypeSystemId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasBefore returns true if the method "DocumentType.before" exists.
func (this DocumentType) HasBefore() bool {
	return js.True == bindings.HasDocumentTypeBefore(
		this.Ref(),
	)
}

// BeforeFunc returns the method "DocumentType.before".
func (this DocumentType) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeBeforeFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "DocumentType.before".
func (this DocumentType) Before(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeBefore(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryBefore calls the method "DocumentType.before"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentType) TryBefore(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAfter returns true if the method "DocumentType.after" exists.
func (this DocumentType) HasAfter() bool {
	return js.True == bindings.HasDocumentTypeAfter(
		this.Ref(),
	)
}

// AfterFunc returns the method "DocumentType.after".
func (this DocumentType) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeAfterFunc(
			this.Ref(),
		),
	)
}

// After calls the method "DocumentType.after".
func (this DocumentType) After(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeAfter(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAfter calls the method "DocumentType.after"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentType) TryAfter(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceWith returns true if the method "DocumentType.replaceWith" exists.
func (this DocumentType) HasReplaceWith() bool {
	return js.True == bindings.HasDocumentTypeReplaceWith(
		this.Ref(),
	)
}

// ReplaceWithFunc returns the method "DocumentType.replaceWith".
func (this DocumentType) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentTypeReplaceWithFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "DocumentType.replaceWith".
func (this DocumentType) ReplaceWith(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeReplaceWith(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceWith calls the method "DocumentType.replaceWith"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentType) TryReplaceWith(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeReplaceWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasRemove returns true if the method "DocumentType.remove" exists.
func (this DocumentType) HasRemove() bool {
	return js.True == bindings.HasDocumentTypeRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "DocumentType.remove".
func (this DocumentType) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentTypeRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "DocumentType.remove".
func (this DocumentType) Remove() (ret js.Void) {
	bindings.CallDocumentTypeRemove(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "DocumentType.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentType) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasCreateDocumentType returns true if the method "DOMImplementation.createDocumentType" exists.
func (this DOMImplementation) HasCreateDocumentType() bool {
	return js.True == bindings.HasDOMImplementationCreateDocumentType(
		this.Ref(),
	)
}

// CreateDocumentTypeFunc returns the method "DOMImplementation.createDocumentType".
func (this DOMImplementation) CreateDocumentTypeFunc() (fn js.Func[func(qualifiedName js.String, publicId js.String, systemId js.String) DocumentType]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocumentTypeFunc(
			this.Ref(),
		),
	)
}

// CreateDocumentType calls the method "DOMImplementation.createDocumentType".
func (this DOMImplementation) CreateDocumentType(qualifiedName js.String, publicId js.String, systemId js.String) (ret DocumentType) {
	bindings.CallDOMImplementationCreateDocumentType(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
		publicId.Ref(),
		systemId.Ref(),
	)

	return
}

// TryCreateDocumentType calls the method "DOMImplementation.createDocumentType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryCreateDocumentType(qualifiedName js.String, publicId js.String, systemId js.String) (ret DocumentType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocumentType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		publicId.Ref(),
		systemId.Ref(),
	)

	return
}

// HasCreateDocument returns true if the method "DOMImplementation.createDocument" exists.
func (this DOMImplementation) HasCreateDocument() bool {
	return js.True == bindings.HasDOMImplementationCreateDocument(
		this.Ref(),
	)
}

// CreateDocumentFunc returns the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocumentFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, doctype DocumentType) XMLDocument]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocumentFunc(
			this.Ref(),
		),
	)
}

// CreateDocument calls the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocument(namespace js.String, qualifiedName js.String, doctype DocumentType) (ret XMLDocument) {
	bindings.CallDOMImplementationCreateDocument(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
		doctype.Ref(),
	)

	return
}

// TryCreateDocument calls the method "DOMImplementation.createDocument"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryCreateDocument(namespace js.String, qualifiedName js.String, doctype DocumentType) (ret XMLDocument, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		doctype.Ref(),
	)

	return
}

// HasCreateDocument1 returns true if the method "DOMImplementation.createDocument" exists.
func (this DOMImplementation) HasCreateDocument1() bool {
	return js.True == bindings.HasDOMImplementationCreateDocument1(
		this.Ref(),
	)
}

// CreateDocument1Func returns the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocument1Func() (fn js.Func[func(namespace js.String, qualifiedName js.String) XMLDocument]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateDocument1Func(
			this.Ref(),
		),
	)
}

// CreateDocument1 calls the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocument1(namespace js.String, qualifiedName js.String) (ret XMLDocument) {
	bindings.CallDOMImplementationCreateDocument1(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// TryCreateDocument1 calls the method "DOMImplementation.createDocument"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryCreateDocument1(namespace js.String, qualifiedName js.String) (ret XMLDocument, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocument1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasCreateHTMLDocument returns true if the method "DOMImplementation.createHTMLDocument" exists.
func (this DOMImplementation) HasCreateHTMLDocument() bool {
	return js.True == bindings.HasDOMImplementationCreateHTMLDocument(
		this.Ref(),
	)
}

// CreateHTMLDocumentFunc returns the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocumentFunc() (fn js.Func[func(title js.String) Document]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateHTMLDocumentFunc(
			this.Ref(),
		),
	)
}

// CreateHTMLDocument calls the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocument(title js.String) (ret Document) {
	bindings.CallDOMImplementationCreateHTMLDocument(
		this.Ref(), js.Pointer(&ret),
		title.Ref(),
	)

	return
}

// TryCreateHTMLDocument calls the method "DOMImplementation.createHTMLDocument"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryCreateHTMLDocument(title js.String) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateHTMLDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		title.Ref(),
	)

	return
}

// HasCreateHTMLDocument1 returns true if the method "DOMImplementation.createHTMLDocument" exists.
func (this DOMImplementation) HasCreateHTMLDocument1() bool {
	return js.True == bindings.HasDOMImplementationCreateHTMLDocument1(
		this.Ref(),
	)
}

// CreateHTMLDocument1Func returns the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocument1Func() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.DOMImplementationCreateHTMLDocument1Func(
			this.Ref(),
		),
	)
}

// CreateHTMLDocument1 calls the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocument1() (ret Document) {
	bindings.CallDOMImplementationCreateHTMLDocument1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateHTMLDocument1 calls the method "DOMImplementation.createHTMLDocument"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryCreateHTMLDocument1() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateHTMLDocument1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHasFeature returns true if the method "DOMImplementation.hasFeature" exists.
func (this DOMImplementation) HasHasFeature() bool {
	return js.True == bindings.HasDOMImplementationHasFeature(
		this.Ref(),
	)
}

// HasFeatureFunc returns the method "DOMImplementation.hasFeature".
func (this DOMImplementation) HasFeatureFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.DOMImplementationHasFeatureFunc(
			this.Ref(),
		),
	)
}

// HasFeature calls the method "DOMImplementation.hasFeature".
func (this DOMImplementation) HasFeature() (ret bool) {
	bindings.CallDOMImplementationHasFeature(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHasFeature calls the method "DOMImplementation.hasFeature"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMImplementation) TryHasFeature() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationHasFeature(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasAllowsFeature returns true if the method "PermissionsPolicy.allowsFeature" exists.
func (this PermissionsPolicy) HasAllowsFeature() bool {
	return js.True == bindings.HasPermissionsPolicyAllowsFeature(
		this.Ref(),
	)
}

// AllowsFeatureFunc returns the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeatureFunc() (fn js.Func[func(feature js.String, origin js.String) bool]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowsFeatureFunc(
			this.Ref(),
		),
	)
}

// AllowsFeature calls the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeature(feature js.String, origin js.String) (ret bool) {
	bindings.CallPermissionsPolicyAllowsFeature(
		this.Ref(), js.Pointer(&ret),
		feature.Ref(),
		origin.Ref(),
	)

	return
}

// TryAllowsFeature calls the method "PermissionsPolicy.allowsFeature"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PermissionsPolicy) TryAllowsFeature(feature js.String, origin js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowsFeature(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
		origin.Ref(),
	)

	return
}

// HasAllowsFeature1 returns true if the method "PermissionsPolicy.allowsFeature" exists.
func (this PermissionsPolicy) HasAllowsFeature1() bool {
	return js.True == bindings.HasPermissionsPolicyAllowsFeature1(
		this.Ref(),
	)
}

// AllowsFeature1Func returns the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeature1Func() (fn js.Func[func(feature js.String) bool]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowsFeature1Func(
			this.Ref(),
		),
	)
}

// AllowsFeature1 calls the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeature1(feature js.String) (ret bool) {
	bindings.CallPermissionsPolicyAllowsFeature1(
		this.Ref(), js.Pointer(&ret),
		feature.Ref(),
	)

	return
}

// TryAllowsFeature1 calls the method "PermissionsPolicy.allowsFeature"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PermissionsPolicy) TryAllowsFeature1(feature js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowsFeature1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
	)

	return
}

// HasFeatures returns true if the method "PermissionsPolicy.features" exists.
func (this PermissionsPolicy) HasFeatures() bool {
	return js.True == bindings.HasPermissionsPolicyFeatures(
		this.Ref(),
	)
}

// FeaturesFunc returns the method "PermissionsPolicy.features".
func (this PermissionsPolicy) FeaturesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyFeaturesFunc(
			this.Ref(),
		),
	)
}

// Features calls the method "PermissionsPolicy.features".
func (this PermissionsPolicy) Features() (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyFeatures(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFeatures calls the method "PermissionsPolicy.features"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PermissionsPolicy) TryFeatures() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyFeatures(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAllowedFeatures returns true if the method "PermissionsPolicy.allowedFeatures" exists.
func (this PermissionsPolicy) HasAllowedFeatures() bool {
	return js.True == bindings.HasPermissionsPolicyAllowedFeatures(
		this.Ref(),
	)
}

// AllowedFeaturesFunc returns the method "PermissionsPolicy.allowedFeatures".
func (this PermissionsPolicy) AllowedFeaturesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyAllowedFeaturesFunc(
			this.Ref(),
		),
	)
}

// AllowedFeatures calls the method "PermissionsPolicy.allowedFeatures".
func (this PermissionsPolicy) AllowedFeatures() (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyAllowedFeatures(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAllowedFeatures calls the method "PermissionsPolicy.allowedFeatures"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PermissionsPolicy) TryAllowedFeatures() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowedFeatures(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAllowlistForFeature returns true if the method "PermissionsPolicy.getAllowlistForFeature" exists.
func (this PermissionsPolicy) HasGetAllowlistForFeature() bool {
	return js.True == bindings.HasPermissionsPolicyGetAllowlistForFeature(
		this.Ref(),
	)
}

// GetAllowlistForFeatureFunc returns the method "PermissionsPolicy.getAllowlistForFeature".
func (this PermissionsPolicy) GetAllowlistForFeatureFunc() (fn js.Func[func(feature js.String) js.Array[js.String]]) {
	return fn.FromRef(
		bindings.PermissionsPolicyGetAllowlistForFeatureFunc(
			this.Ref(),
		),
	)
}

// GetAllowlistForFeature calls the method "PermissionsPolicy.getAllowlistForFeature".
func (this PermissionsPolicy) GetAllowlistForFeature(feature js.String) (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyGetAllowlistForFeature(
		this.Ref(), js.Pointer(&ret),
		feature.Ref(),
	)

	return
}

// TryGetAllowlistForFeature calls the method "PermissionsPolicy.getAllowlistForFeature"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PermissionsPolicy) TryGetAllowlistForFeature(feature js.String) (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyGetAllowlistForFeature(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
	)

	return
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

func NewDocumentTimeline(options DocumentTimelineOptions) (ret DocumentTimeline) {
	ret.ref = bindings.NewDocumentTimelineByDocumentTimeline(
		js.Pointer(&options))
	return
}

func NewDocumentTimelineByDocumentTimeline1() (ret DocumentTimeline) {
	ret.ref = bindings.NewDocumentTimelineByDocumentTimeline1()
	return
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
// It returns ok=false if there is no such property.
func (this SVGAnimatedString) BaseVal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedStringBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedString.baseVal" to val.
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
// It returns ok=false if there is no such property.
func (this SVGAnimatedString) AnimVal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedStringAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) CssText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationCssText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCssText sets the value of property "CSSStyleDeclaration.cssText" to val.
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
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ParentRule returns the value of property "CSSStyleDeclaration.parentRule".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) ParentRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationParentRule(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CssFloat returns the value of property "CSSStyleDeclaration.cssFloat".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) CssFloat() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationCssFloat(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCssFloat sets the value of property "CSSStyleDeclaration.cssFloat" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleDeclaration) SetCssFloat(val js.String) bool {
	return js.True == bindings.SetCSSStyleDeclarationCssFloat(
		this.Ref(),
		val.Ref(),
	)
}

// HasItem returns true if the method "CSSStyleDeclaration.item" exists.
func (this CSSStyleDeclaration) HasItem() bool {
	return js.True == bindings.HasCSSStyleDeclarationItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "CSSStyleDeclaration.item".
func (this CSSStyleDeclaration) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "CSSStyleDeclaration.item".
func (this CSSStyleDeclaration) Item(index uint32) (ret js.String) {
	bindings.CallCSSStyleDeclarationItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "CSSStyleDeclaration.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasGetPropertyValue returns true if the method "CSSStyleDeclaration.getPropertyValue" exists.
func (this CSSStyleDeclaration) HasGetPropertyValue() bool {
	return js.True == bindings.HasCSSStyleDeclarationGetPropertyValue(
		this.Ref(),
	)
}

// GetPropertyValueFunc returns the method "CSSStyleDeclaration.getPropertyValue".
func (this CSSStyleDeclaration) GetPropertyValueFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationGetPropertyValueFunc(
			this.Ref(),
		),
	)
}

// GetPropertyValue calls the method "CSSStyleDeclaration.getPropertyValue".
func (this CSSStyleDeclaration) GetPropertyValue(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationGetPropertyValue(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetPropertyValue calls the method "CSSStyleDeclaration.getPropertyValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TryGetPropertyValue(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationGetPropertyValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasGetPropertyPriority returns true if the method "CSSStyleDeclaration.getPropertyPriority" exists.
func (this CSSStyleDeclaration) HasGetPropertyPriority() bool {
	return js.True == bindings.HasCSSStyleDeclarationGetPropertyPriority(
		this.Ref(),
	)
}

// GetPropertyPriorityFunc returns the method "CSSStyleDeclaration.getPropertyPriority".
func (this CSSStyleDeclaration) GetPropertyPriorityFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationGetPropertyPriorityFunc(
			this.Ref(),
		),
	)
}

// GetPropertyPriority calls the method "CSSStyleDeclaration.getPropertyPriority".
func (this CSSStyleDeclaration) GetPropertyPriority(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationGetPropertyPriority(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetPropertyPriority calls the method "CSSStyleDeclaration.getPropertyPriority"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TryGetPropertyPriority(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationGetPropertyPriority(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasSetProperty returns true if the method "CSSStyleDeclaration.setProperty" exists.
func (this CSSStyleDeclaration) HasSetProperty() bool {
	return js.True == bindings.HasCSSStyleDeclarationSetProperty(
		this.Ref(),
	)
}

// SetPropertyFunc returns the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetPropertyFunc() (fn js.Func[func(property js.String, value js.String, priority js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationSetPropertyFunc(
			this.Ref(),
		),
	)
}

// SetProperty calls the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetProperty(property js.String, value js.String, priority js.String) (ret js.Void) {
	bindings.CallCSSStyleDeclarationSetProperty(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		value.Ref(),
		priority.Ref(),
	)

	return
}

// TrySetProperty calls the method "CSSStyleDeclaration.setProperty"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TrySetProperty(property js.String, value js.String, priority js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationSetProperty(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
		priority.Ref(),
	)

	return
}

// HasSetProperty1 returns true if the method "CSSStyleDeclaration.setProperty" exists.
func (this CSSStyleDeclaration) HasSetProperty1() bool {
	return js.True == bindings.HasCSSStyleDeclarationSetProperty1(
		this.Ref(),
	)
}

// SetProperty1Func returns the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetProperty1Func() (fn js.Func[func(property js.String, value js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationSetProperty1Func(
			this.Ref(),
		),
	)
}

// SetProperty1 calls the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetProperty1(property js.String, value js.String) (ret js.Void) {
	bindings.CallCSSStyleDeclarationSetProperty1(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		value.Ref(),
	)

	return
}

// TrySetProperty1 calls the method "CSSStyleDeclaration.setProperty"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TrySetProperty1(property js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationSetProperty1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
	)

	return
}

// HasRemoveProperty returns true if the method "CSSStyleDeclaration.removeProperty" exists.
func (this CSSStyleDeclaration) HasRemoveProperty() bool {
	return js.True == bindings.HasCSSStyleDeclarationRemoveProperty(
		this.Ref(),
	)
}

// RemovePropertyFunc returns the method "CSSStyleDeclaration.removeProperty".
func (this CSSStyleDeclaration) RemovePropertyFunc() (fn js.Func[func(property js.String) js.String]) {
	return fn.FromRef(
		bindings.CSSStyleDeclarationRemovePropertyFunc(
			this.Ref(),
		),
	)
}

// RemoveProperty calls the method "CSSStyleDeclaration.removeProperty".
func (this CSSStyleDeclaration) RemoveProperty(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationRemoveProperty(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryRemoveProperty calls the method "CSSStyleDeclaration.removeProperty"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleDeclaration) TryRemoveProperty(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationRemoveProperty(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
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

// HasSet returns true if the method "StylePropertyMap.set" exists.
func (this StylePropertyMap) HasSet() bool {
	return js.True == bindings.HasStylePropertyMapSet(
		this.Ref(),
	)
}

// SetFunc returns the method "StylePropertyMap.set".
func (this StylePropertyMap) SetFunc() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "StylePropertyMap.set".
func (this StylePropertyMap) Set(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void) {
	bindings.CallStylePropertyMapSet(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TrySet calls the method "StylePropertyMap.set"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMap) TrySet(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasAppend returns true if the method "StylePropertyMap.append" exists.
func (this StylePropertyMap) HasAppend() bool {
	return js.True == bindings.HasStylePropertyMapAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "StylePropertyMap.append".
func (this StylePropertyMap) AppendFunc() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "StylePropertyMap.append".
func (this StylePropertyMap) Append(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void) {
	bindings.CallStylePropertyMapAppend(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryAppend calls the method "StylePropertyMap.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMap) TryAppend(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasDelete returns true if the method "StylePropertyMap.delete" exists.
func (this StylePropertyMap) HasDelete() bool {
	return js.True == bindings.HasStylePropertyMapDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "StylePropertyMap.delete".
func (this StylePropertyMap) DeleteFunc() (fn js.Func[func(property js.String)]) {
	return fn.FromRef(
		bindings.StylePropertyMapDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "StylePropertyMap.delete".
func (this StylePropertyMap) Delete(property js.String) (ret js.Void) {
	bindings.CallStylePropertyMapDelete(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryDelete calls the method "StylePropertyMap.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMap) TryDelete(property js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasClear returns true if the method "StylePropertyMap.clear" exists.
func (this StylePropertyMap) HasClear() bool {
	return js.True == bindings.HasStylePropertyMapClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "StylePropertyMap.clear".
func (this StylePropertyMap) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.StylePropertyMapClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "StylePropertyMap.clear".
func (this StylePropertyMap) Clear() (ret js.Void) {
	bindings.CallStylePropertyMapClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "StylePropertyMap.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMap) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGLength) UnitType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGLengthUnitType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "SVGLength.value".
//
// It returns ok=false if there is no such property.
func (this SVGLength) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGLengthValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGLength.value" to val.
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
// It returns ok=false if there is no such property.
func (this SVGLength) ValueInSpecifiedUnits() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGLengthValueInSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueInSpecifiedUnits sets the value of property "SVGLength.valueInSpecifiedUnits" to val.
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
// It returns ok=false if there is no such property.
func (this SVGLength) ValueAsString() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGLengthValueAsString(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueAsString sets the value of property "SVGLength.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGLengthValueAsString(
		this.Ref(),
		val.Ref(),
	)
}

// HasNewValueSpecifiedUnits returns true if the method "SVGLength.newValueSpecifiedUnits" exists.
func (this SVGLength) HasNewValueSpecifiedUnits() bool {
	return js.True == bindings.HasSVGLengthNewValueSpecifiedUnits(
		this.Ref(),
	)
}

// NewValueSpecifiedUnitsFunc returns the method "SVGLength.newValueSpecifiedUnits".
func (this SVGLength) NewValueSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	return fn.FromRef(
		bindings.SVGLengthNewValueSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// NewValueSpecifiedUnits calls the method "SVGLength.newValueSpecifiedUnits".
func (this SVGLength) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void) {
	bindings.CallSVGLengthNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// TryNewValueSpecifiedUnits calls the method "SVGLength.newValueSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLength) TryNewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// HasConvertToSpecifiedUnits returns true if the method "SVGLength.convertToSpecifiedUnits" exists.
func (this SVGLength) HasConvertToSpecifiedUnits() bool {
	return js.True == bindings.HasSVGLengthConvertToSpecifiedUnits(
		this.Ref(),
	)
}

// ConvertToSpecifiedUnitsFunc returns the method "SVGLength.convertToSpecifiedUnits".
func (this SVGLength) ConvertToSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16)]) {
	return fn.FromRef(
		bindings.SVGLengthConvertToSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// ConvertToSpecifiedUnits calls the method "SVGLength.convertToSpecifiedUnits".
func (this SVGLength) ConvertToSpecifiedUnits(unitType uint16) (ret js.Void) {
	bindings.CallSVGLengthConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
		uint32(unitType),
	)

	return
}

// TryConvertToSpecifiedUnits calls the method "SVGLength.convertToSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLength) TryConvertToSpecifiedUnits(unitType uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGAnimatedLength) BaseVal() (ret SVGLength, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedLength.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLength) AnimVal() (ret SVGLength, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGUseElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGUseElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGUseElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGUseElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InstanceRoot returns the value of property "SVGUseElement.instanceRoot".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) InstanceRoot() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGUseElementInstanceRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimatedInstanceRoot returns the value of property "SVGUseElement.animatedInstanceRoot".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) AnimatedInstanceRoot() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGUseElementAnimatedInstanceRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGUseElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGUseElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasGet returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasGet() bool {
	return js.True == bindings.HasDOMStringMapGet(
		this.Ref(),
	)
}

// GetFunc returns the method "DOMStringMap.".
func (this DOMStringMap) GetFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.DOMStringMapGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "DOMStringMap.".
func (this DOMStringMap) Get(name js.String) (ret js.String) {
	bindings.CallDOMStringMapGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMStringMap) TryGet(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSet returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasSet() bool {
	return js.True == bindings.HasDOMStringMapSet(
		this.Ref(),
	)
}

// SetFunc returns the method "DOMStringMap.".
func (this DOMStringMap) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.DOMStringMapSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "DOMStringMap.".
func (this DOMStringMap) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallDOMStringMapSet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMStringMap) TrySet(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasDelete returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasDelete() bool {
	return js.True == bindings.HasDOMStringMapDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "DOMStringMap.".
func (this DOMStringMap) DeleteFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.DOMStringMapDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "DOMStringMap.".
func (this DOMStringMap) Delete(name js.String) (ret js.Void) {
	bindings.CallDOMStringMapDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMStringMap) TryDelete(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGElement) ClassName() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGElementClassName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OwnerSVGElement returns the value of property "SVGElement.ownerSVGElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) OwnerSVGElement() (ret SVGSVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementOwnerSVGElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ViewportElement returns the value of property "SVGElement.viewportElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) ViewportElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementViewportElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "SVGElement.style".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetSVGElementStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "SVGElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this SVGElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetSVGElementAttributeStyleMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CorrespondingElement returns the value of property "SVGElement.correspondingElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) CorrespondingElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementCorrespondingElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CorrespondingUseElement returns the value of property "SVGElement.correspondingUseElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) CorrespondingUseElement() (ret SVGUseElement, ok bool) {
	ok = js.True == bindings.GetSVGElementCorrespondingUseElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "SVGElement.dataset".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetSVGElementDataset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "SVGElement.nonce".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGElementNonce(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "SVGElement.nonce" to val.
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
// It returns ok=false if there is no such property.
func (this SVGElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetSVGElementAutofocus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "SVGElement.autofocus" to val.
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
// It returns ok=false if there is no such property.
func (this SVGElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGElementTabIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "SVGElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetSVGElementTabIndex(
		this.Ref(),
		int32(val),
	)
}

// HasFocus returns true if the method "SVGElement.focus" exists.
func (this SVGElement) HasFocus() bool {
	return js.True == bindings.HasSVGElementFocus(
		this.Ref(),
	)
}

// FocusFunc returns the method "SVGElement.focus".
func (this SVGElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.SVGElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "SVGElement.focus".
func (this SVGElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallSVGElementFocus(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "SVGElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFocus1 returns true if the method "SVGElement.focus" exists.
func (this SVGElement) HasFocus1() bool {
	return js.True == bindings.HasSVGElementFocus1(
		this.Ref(),
	)
}

// Focus1Func returns the method "SVGElement.focus".
func (this SVGElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGElementFocus1Func(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "SVGElement.focus".
func (this SVGElement) Focus1() (ret js.Void) {
	bindings.CallSVGElementFocus1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "SVGElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementFocus1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlur returns true if the method "SVGElement.blur" exists.
func (this SVGElement) HasBlur() bool {
	return js.True == bindings.HasSVGElementBlur(
		this.Ref(),
	)
}

// BlurFunc returns the method "SVGElement.blur".
func (this SVGElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGElementBlurFunc(
			this.Ref(),
		),
	)
}

// Blur calls the method "SVGElement.blur".
func (this SVGElement) Blur() (ret js.Void) {
	bindings.CallSVGElementBlur(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "SVGElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementBlur(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGNumber) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGNumberValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGNumber.value" to val.
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
// It returns ok=false if there is no such property.
func (this SVGAngle) UnitType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAngleUnitType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "SVGAngle.value".
//
// It returns ok=false if there is no such property.
func (this SVGAngle) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAngleValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGAngle.value" to val.
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
// It returns ok=false if there is no such property.
func (this SVGAngle) ValueInSpecifiedUnits() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAngleValueInSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueInSpecifiedUnits sets the value of property "SVGAngle.valueInSpecifiedUnits" to val.
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
// It returns ok=false if there is no such property.
func (this SVGAngle) ValueAsString() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAngleValueAsString(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueAsString sets the value of property "SVGAngle.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGAngleValueAsString(
		this.Ref(),
		val.Ref(),
	)
}

// HasNewValueSpecifiedUnits returns true if the method "SVGAngle.newValueSpecifiedUnits" exists.
func (this SVGAngle) HasNewValueSpecifiedUnits() bool {
	return js.True == bindings.HasSVGAngleNewValueSpecifiedUnits(
		this.Ref(),
	)
}

// NewValueSpecifiedUnitsFunc returns the method "SVGAngle.newValueSpecifiedUnits".
func (this SVGAngle) NewValueSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	return fn.FromRef(
		bindings.SVGAngleNewValueSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// NewValueSpecifiedUnits calls the method "SVGAngle.newValueSpecifiedUnits".
func (this SVGAngle) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void) {
	bindings.CallSVGAngleNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// TryNewValueSpecifiedUnits calls the method "SVGAngle.newValueSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAngle) TryNewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAngleNewValueSpecifiedUnits(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// HasConvertToSpecifiedUnits returns true if the method "SVGAngle.convertToSpecifiedUnits" exists.
func (this SVGAngle) HasConvertToSpecifiedUnits() bool {
	return js.True == bindings.HasSVGAngleConvertToSpecifiedUnits(
		this.Ref(),
	)
}

// ConvertToSpecifiedUnitsFunc returns the method "SVGAngle.convertToSpecifiedUnits".
func (this SVGAngle) ConvertToSpecifiedUnitsFunc() (fn js.Func[func(unitType uint16)]) {
	return fn.FromRef(
		bindings.SVGAngleConvertToSpecifiedUnitsFunc(
			this.Ref(),
		),
	)
}

// ConvertToSpecifiedUnits calls the method "SVGAngle.convertToSpecifiedUnits".
func (this SVGAngle) ConvertToSpecifiedUnits(unitType uint16) (ret js.Void) {
	bindings.CallSVGAngleConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&ret),
		uint32(unitType),
	)

	return
}

// TryConvertToSpecifiedUnits calls the method "SVGAngle.convertToSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAngle) TryConvertToSpecifiedUnits(unitType uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAngleConvertToSpecifiedUnits(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
	)

	return
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

func NewDOMMatrix(init OneOf_String_ArrayFloat64) (ret DOMMatrix) {
	ret.ref = bindings.NewDOMMatrixByDOMMatrix(
		init.Ref())
	return
}

func NewDOMMatrixByDOMMatrix1() (ret DOMMatrix) {
	ret.ref = bindings.NewDOMMatrixByDOMMatrix1()
	return
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) A() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixA(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "DOMMatrix.a" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) B() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "DOMMatrix.b" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) C() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixC(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "DOMMatrix.c" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) D() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixD(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetD sets the value of property "DOMMatrix.d" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) E() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixE(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetE sets the value of property "DOMMatrix.e" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) F() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixF(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetF sets the value of property "DOMMatrix.f" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M11() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM11(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM11 sets the value of property "DOMMatrix.m11" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M12() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM12(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM12 sets the value of property "DOMMatrix.m12" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M13() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM13(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM13 sets the value of property "DOMMatrix.m13" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M14() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM14(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM14 sets the value of property "DOMMatrix.m14" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M21() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM21(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM21 sets the value of property "DOMMatrix.m21" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M22() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM22(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM22 sets the value of property "DOMMatrix.m22" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M23() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM23(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM23 sets the value of property "DOMMatrix.m23" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M24() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM24(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM24 sets the value of property "DOMMatrix.m24" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M31() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM31(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM31 sets the value of property "DOMMatrix.m31" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M32() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM32(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM32 sets the value of property "DOMMatrix.m32" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M33() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM33(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM33 sets the value of property "DOMMatrix.m33" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M34() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM34(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM34 sets the value of property "DOMMatrix.m34" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M41() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM41(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM41 sets the value of property "DOMMatrix.m41" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M42() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM42(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM42 sets the value of property "DOMMatrix.m42" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M43() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM43(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM43 sets the value of property "DOMMatrix.m43" to val.
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
// It returns ok=false if there is no such property.
func (this DOMMatrix) M44() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM44(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetM44 sets the value of property "DOMMatrix.m44" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM44(val float64) bool {
	return js.True == bindings.SetDOMMatrixM44(
		this.Ref(),
		float64(val),
	)
}

// HasFromMatrix returns true if the staticmethod "DOMMatrix.fromMatrix" exists.
func (this DOMMatrix) HasFromMatrix() bool {
	return js.True == bindings.HasDOMMatrixFromMatrix(
		this.Ref(),
	)
}

// FromMatrixFunc returns the staticmethod "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrixFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromMatrixFunc(
			this.Ref(),
		),
	)
}

// FromMatrix calls the staticmethod "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrix(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromMatrix(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromMatrix calls the staticmethod "DOMMatrix.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryFromMatrix(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromMatrix1 returns true if the staticmethod "DOMMatrix.fromMatrix" exists.
func (this DOMMatrix) HasFromMatrix1() bool {
	return js.True == bindings.HasDOMMatrixFromMatrix1(
		this.Ref(),
	)
}

// FromMatrix1Func returns the staticmethod "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrix1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromMatrix1Func(
			this.Ref(),
		),
	)
}

// FromMatrix1 calls the staticmethod "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrix1() (ret DOMMatrix) {
	bindings.CallDOMMatrixFromMatrix1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromMatrix1 calls the staticmethod "DOMMatrix.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryFromMatrix1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromMatrix1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromFloat32Array returns true if the staticmethod "DOMMatrix.fromFloat32Array" exists.
func (this DOMMatrix) HasFromFloat32Array() bool {
	return js.True == bindings.HasDOMMatrixFromFloat32Array(
		this.Ref(),
	)
}

// FromFloat32ArrayFunc returns the staticmethod "DOMMatrix.fromFloat32Array".
func (this DOMMatrix) FromFloat32ArrayFunc() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat32Array calls the staticmethod "DOMMatrix.fromFloat32Array".
func (this DOMMatrix) FromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromFloat32Array(
		this.Ref(), js.Pointer(&ret),
		array32.Ref(),
	)

	return
}

// TryFromFloat32Array calls the staticmethod "DOMMatrix.fromFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryFromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromFloat32Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array32.Ref(),
	)

	return
}

// HasFromFloat64Array returns true if the staticmethod "DOMMatrix.fromFloat64Array" exists.
func (this DOMMatrix) HasFromFloat64Array() bool {
	return js.True == bindings.HasDOMMatrixFromFloat64Array(
		this.Ref(),
	)
}

// FromFloat64ArrayFunc returns the staticmethod "DOMMatrix.fromFloat64Array".
func (this DOMMatrix) FromFloat64ArrayFunc() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixFromFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat64Array calls the staticmethod "DOMMatrix.fromFloat64Array".
func (this DOMMatrix) FromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromFloat64Array(
		this.Ref(), js.Pointer(&ret),
		array64.Ref(),
	)

	return
}

// TryFromFloat64Array calls the staticmethod "DOMMatrix.fromFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryFromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromFloat64Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array64.Ref(),
	)

	return
}

// HasMultiplySelf returns true if the method "DOMMatrix.multiplySelf" exists.
func (this DOMMatrix) HasMultiplySelf() bool {
	return js.True == bindings.HasDOMMatrixMultiplySelf(
		this.Ref(),
	)
}

// MultiplySelfFunc returns the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelfFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixMultiplySelfFunc(
			this.Ref(),
		),
	)
}

// MultiplySelf calls the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelf(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixMultiplySelf(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryMultiplySelf calls the method "DOMMatrix.multiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryMultiplySelf(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixMultiplySelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasMultiplySelf1 returns true if the method "DOMMatrix.multiplySelf" exists.
func (this DOMMatrix) HasMultiplySelf1() bool {
	return js.True == bindings.HasDOMMatrixMultiplySelf1(
		this.Ref(),
	)
}

// MultiplySelf1Func returns the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixMultiplySelf1Func(
			this.Ref(),
		),
	)
}

// MultiplySelf1 calls the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixMultiplySelf1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMultiplySelf1 calls the method "DOMMatrix.multiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryMultiplySelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixMultiplySelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPreMultiplySelf returns true if the method "DOMMatrix.preMultiplySelf" exists.
func (this DOMMatrix) HasPreMultiplySelf() bool {
	return js.True == bindings.HasDOMMatrixPreMultiplySelf(
		this.Ref(),
	)
}

// PreMultiplySelfFunc returns the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelfFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixPreMultiplySelfFunc(
			this.Ref(),
		),
	)
}

// PreMultiplySelf calls the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelf(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixPreMultiplySelf(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryPreMultiplySelf calls the method "DOMMatrix.preMultiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryPreMultiplySelf(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixPreMultiplySelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasPreMultiplySelf1 returns true if the method "DOMMatrix.preMultiplySelf" exists.
func (this DOMMatrix) HasPreMultiplySelf1() bool {
	return js.True == bindings.HasDOMMatrixPreMultiplySelf1(
		this.Ref(),
	)
}

// PreMultiplySelf1Func returns the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixPreMultiplySelf1Func(
			this.Ref(),
		),
	)
}

// PreMultiplySelf1 calls the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixPreMultiplySelf1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPreMultiplySelf1 calls the method "DOMMatrix.preMultiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryPreMultiplySelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixPreMultiplySelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTranslateSelf returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasTranslateSelf() bool {
	return js.True == bindings.HasDOMMatrixTranslateSelf(
		this.Ref(),
	)
}

// TranslateSelfFunc returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelfFunc() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelfFunc(
			this.Ref(),
		),
	)
}

// TranslateSelf calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf(tx float64, ty float64, tz float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// TryTranslateSelf calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryTranslateSelf(tx float64, ty float64, tz float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// HasTranslateSelf1 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasTranslateSelf1() bool {
	return js.True == bindings.HasDOMMatrixTranslateSelf1(
		this.Ref(),
	)
}

// TranslateSelf1Func returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf1Func() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf1Func(
			this.Ref(),
		),
	)
}

// TranslateSelf1 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf1(tx float64, ty float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
		float64(ty),
	)

	return
}

// TryTranslateSelf1 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryTranslateSelf1(tx float64, ty float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
	)

	return
}

// HasTranslateSelf2 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasTranslateSelf2() bool {
	return js.True == bindings.HasDOMMatrixTranslateSelf2(
		this.Ref(),
	)
}

// TranslateSelf2Func returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf2Func() (fn js.Func[func(tx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf2Func(
			this.Ref(),
		),
	)
}

// TranslateSelf2 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf2(tx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf2(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
	)

	return
}

// TryTranslateSelf2 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryTranslateSelf2(tx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
	)

	return
}

// HasTranslateSelf3 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasTranslateSelf3() bool {
	return js.True == bindings.HasDOMMatrixTranslateSelf3(
		this.Ref(),
	)
}

// TranslateSelf3Func returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixTranslateSelf3Func(
			this.Ref(),
		),
	)
}

// TranslateSelf3 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf3() (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTranslateSelf3 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryTranslateSelf3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScaleSelf returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf(
		this.Ref(),
	)
}

// ScaleSelfFunc returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelfFunc() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelfFunc(
			this.Ref(),
		),
	)
}

// ScaleSelf calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// TryScaleSelf calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasScaleSelf1 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf1() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf1(
		this.Ref(),
	)
}

// ScaleSelf1Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf1Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf1Func(
			this.Ref(),
		),
	)
}

// ScaleSelf1 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScaleSelf1 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// HasScaleSelf2 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf2() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf2(
		this.Ref(),
	)
}

// ScaleSelf2Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf2Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf2Func(
			this.Ref(),
		),
	)
}

// ScaleSelf2 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf2(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// TryScaleSelf2 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// HasScaleSelf3 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf3() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf3(
		this.Ref(),
	)
}

// ScaleSelf3Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf3Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf3Func(
			this.Ref(),
		),
	)
}

// ScaleSelf3 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf3(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// TryScaleSelf3 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// HasScaleSelf4 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf4() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf4(
		this.Ref(),
	)
}

// ScaleSelf4Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf4Func() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf4Func(
			this.Ref(),
		),
	)
}

// ScaleSelf4 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf4(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf4(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// TryScaleSelf4 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf4(scaleX float64, scaleY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasScaleSelf5 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf5() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf5(
		this.Ref(),
	)
}

// ScaleSelf5Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf5Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf5Func(
			this.Ref(),
		),
	)
}

// ScaleSelf5 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf5(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf5(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScaleSelf5 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf5(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasScaleSelf6 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasScaleSelf6() bool {
	return js.True == bindings.HasDOMMatrixScaleSelf6(
		this.Ref(),
	)
}

// ScaleSelf6Func returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf6Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScaleSelf6Func(
			this.Ref(),
		),
	)
}

// ScaleSelf6 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf6() (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf6(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScaleSelf6 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScaleSelf6() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScale3dSelf returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasScale3dSelf() bool {
	return js.True == bindings.HasDOMMatrixScale3dSelf(
		this.Ref(),
	)
}

// Scale3dSelfFunc returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelfFunc() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelfFunc(
			this.Ref(),
		),
	)
}

// Scale3dSelf calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// TryScale3dSelf calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScale3dSelf(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasScale3dSelf1 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasScale3dSelf1() bool {
	return js.True == bindings.HasDOMMatrixScale3dSelf1(
		this.Ref(),
	)
}

// Scale3dSelf1Func returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf1Func() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf1Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf1 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf1(scale float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScale3dSelf1 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScale3dSelf1(scale float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// HasScale3dSelf2 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasScale3dSelf2() bool {
	return js.True == bindings.HasDOMMatrixScale3dSelf2(
		this.Ref(),
	)
}

// Scale3dSelf2Func returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf2Func() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf2Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf2 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf2(scale float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf2(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
	)

	return
}

// TryScale3dSelf2 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScale3dSelf2(scale float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
	)

	return
}

// HasScale3dSelf3 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasScale3dSelf3() bool {
	return js.True == bindings.HasDOMMatrixScale3dSelf3(
		this.Ref(),
	)
}

// Scale3dSelf3Func returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf3Func() (fn js.Func[func(scale float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf3Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf3 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf3(scale float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf3(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
	)

	return
}

// TryScale3dSelf3 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScale3dSelf3(scale float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
	)

	return
}

// HasScale3dSelf4 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasScale3dSelf4() bool {
	return js.True == bindings.HasDOMMatrixScale3dSelf4(
		this.Ref(),
	)
}

// Scale3dSelf4Func returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixScale3dSelf4Func(
			this.Ref(),
		),
	)
}

// Scale3dSelf4 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf4() (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScale3dSelf4 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryScale3dSelf4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotateSelf returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasRotateSelf() bool {
	return js.True == bindings.HasDOMMatrixRotateSelf(
		this.Ref(),
	)
}

// RotateSelfFunc returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelfFunc() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelfFunc(
			this.Ref(),
		),
	)
}

// RotateSelf calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// TryRotateSelf calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateSelf(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// HasRotateSelf1 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasRotateSelf1() bool {
	return js.True == bindings.HasDOMMatrixRotateSelf1(
		this.Ref(),
	)
}

// RotateSelf1Func returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf1Func() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf1Func(
			this.Ref(),
		),
	)
}

// RotateSelf1 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf1(rotX float64, rotY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
	)

	return
}

// TryRotateSelf1 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateSelf1(rotX float64, rotY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
	)

	return
}

// HasRotateSelf2 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasRotateSelf2() bool {
	return js.True == bindings.HasDOMMatrixRotateSelf2(
		this.Ref(),
	)
}

// RotateSelf2Func returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf2Func() (fn js.Func[func(rotX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf2Func(
			this.Ref(),
		),
	)
}

// RotateSelf2 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf2(rotX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf2(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
	)

	return
}

// TryRotateSelf2 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateSelf2(rotX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
	)

	return
}

// HasRotateSelf3 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasRotateSelf3() bool {
	return js.True == bindings.HasDOMMatrixRotateSelf3(
		this.Ref(),
	)
}

// RotateSelf3Func returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateSelf3Func(
			this.Ref(),
		),
	)
}

// RotateSelf3 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf3() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotateSelf3 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateSelf3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotateFromVectorSelf returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasRotateFromVectorSelf() bool {
	return js.True == bindings.HasDOMMatrixRotateFromVectorSelf(
		this.Ref(),
	)
}

// RotateFromVectorSelfFunc returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelfFunc() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelfFunc(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateFromVectorSelf calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotateFromVectorSelf1 returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasRotateFromVectorSelf1() bool {
	return js.True == bindings.HasDOMMatrixRotateFromVectorSelf1(
		this.Ref(),
	)
}

// RotateFromVectorSelf1Func returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf1Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelf1Func(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf1 calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf1(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateFromVectorSelf1 calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf1(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasRotateFromVectorSelf2 returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasRotateFromVectorSelf2() bool {
	return js.True == bindings.HasDOMMatrixRotateFromVectorSelf2(
		this.Ref(),
	)
}

// RotateFromVectorSelf2Func returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateFromVectorSelf2Func(
			this.Ref(),
		),
	)
}

// RotateFromVectorSelf2 calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf2() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotateFromVectorSelf2 calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotateAxisAngleSelf returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasRotateAxisAngleSelf() bool {
	return js.True == bindings.HasDOMMatrixRotateAxisAngleSelf(
		this.Ref(),
	)
}

// RotateAxisAngleSelfFunc returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelfFunc() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelfFunc(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// TryRotateAxisAngleSelf calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf(x float64, y float64, z float64, angle float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// HasRotateAxisAngleSelf1 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasRotateAxisAngleSelf1() bool {
	return js.True == bindings.HasDOMMatrixRotateAxisAngleSelf1(
		this.Ref(),
	)
}

// RotateAxisAngleSelf1Func returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf1Func() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf1Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf1 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf1(x float64, y float64, z float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// TryRotateAxisAngleSelf1 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf1(x float64, y float64, z float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// HasRotateAxisAngleSelf2 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasRotateAxisAngleSelf2() bool {
	return js.True == bindings.HasDOMMatrixRotateAxisAngleSelf2(
		this.Ref(),
	)
}

// RotateAxisAngleSelf2Func returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf2Func() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf2 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf2(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateAxisAngleSelf2 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf2(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotateAxisAngleSelf3 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasRotateAxisAngleSelf3() bool {
	return js.True == bindings.HasDOMMatrixRotateAxisAngleSelf3(
		this.Ref(),
	)
}

// RotateAxisAngleSelf3Func returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf3Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf3Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf3 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf3(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf3(
		this.Ref(), js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateAxisAngleSelf3 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf3(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasRotateAxisAngleSelf4 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasRotateAxisAngleSelf4() bool {
	return js.True == bindings.HasDOMMatrixRotateAxisAngleSelf4(
		this.Ref(),
	)
}

// RotateAxisAngleSelf4Func returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixRotateAxisAngleSelf4Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngleSelf4 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf4() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotateAxisAngleSelf4 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSkewXSelf returns true if the method "DOMMatrix.skewXSelf" exists.
func (this DOMMatrix) HasSkewXSelf() bool {
	return js.True == bindings.HasDOMMatrixSkewXSelf(
		this.Ref(),
	)
}

// SkewXSelfFunc returns the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelfFunc() (fn js.Func[func(sx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewXSelfFunc(
			this.Ref(),
		),
	)
}

// SkewXSelf calls the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelf(sx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewXSelf(
		this.Ref(), js.Pointer(&ret),
		float64(sx),
	)

	return
}

// TrySkewXSelf calls the method "DOMMatrix.skewXSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TrySkewXSelf(sx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewXSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(sx),
	)

	return
}

// HasSkewXSelf1 returns true if the method "DOMMatrix.skewXSelf" exists.
func (this DOMMatrix) HasSkewXSelf1() bool {
	return js.True == bindings.HasDOMMatrixSkewXSelf1(
		this.Ref(),
	)
}

// SkewXSelf1Func returns the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewXSelf1Func(
			this.Ref(),
		),
	)
}

// SkewXSelf1 calls the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewXSelf1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySkewXSelf1 calls the method "DOMMatrix.skewXSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TrySkewXSelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewXSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSkewYSelf returns true if the method "DOMMatrix.skewYSelf" exists.
func (this DOMMatrix) HasSkewYSelf() bool {
	return js.True == bindings.HasDOMMatrixSkewYSelf(
		this.Ref(),
	)
}

// SkewYSelfFunc returns the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelfFunc() (fn js.Func[func(sy float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewYSelfFunc(
			this.Ref(),
		),
	)
}

// SkewYSelf calls the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelf(sy float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewYSelf(
		this.Ref(), js.Pointer(&ret),
		float64(sy),
	)

	return
}

// TrySkewYSelf calls the method "DOMMatrix.skewYSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TrySkewYSelf(sy float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewYSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(sy),
	)

	return
}

// HasSkewYSelf1 returns true if the method "DOMMatrix.skewYSelf" exists.
func (this DOMMatrix) HasSkewYSelf1() bool {
	return js.True == bindings.HasDOMMatrixSkewYSelf1(
		this.Ref(),
	)
}

// SkewYSelf1Func returns the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelf1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSkewYSelf1Func(
			this.Ref(),
		),
	)
}

// SkewYSelf1 calls the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewYSelf1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySkewYSelf1 calls the method "DOMMatrix.skewYSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TrySkewYSelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewYSelf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInvertSelf returns true if the method "DOMMatrix.invertSelf" exists.
func (this DOMMatrix) HasInvertSelf() bool {
	return js.True == bindings.HasDOMMatrixInvertSelf(
		this.Ref(),
	)
}

// InvertSelfFunc returns the method "DOMMatrix.invertSelf".
func (this DOMMatrix) InvertSelfFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixInvertSelfFunc(
			this.Ref(),
		),
	)
}

// InvertSelf calls the method "DOMMatrix.invertSelf".
func (this DOMMatrix) InvertSelf() (ret DOMMatrix) {
	bindings.CallDOMMatrixInvertSelf(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryInvertSelf calls the method "DOMMatrix.invertSelf"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TryInvertSelf() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixInvertSelf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetMatrixValue returns true if the method "DOMMatrix.setMatrixValue" exists.
func (this DOMMatrix) HasSetMatrixValue() bool {
	return js.True == bindings.HasDOMMatrixSetMatrixValue(
		this.Ref(),
	)
}

// SetMatrixValueFunc returns the method "DOMMatrix.setMatrixValue".
func (this DOMMatrix) SetMatrixValueFunc() (fn js.Func[func(transformList js.String) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixSetMatrixValueFunc(
			this.Ref(),
		),
	)
}

// SetMatrixValue calls the method "DOMMatrix.setMatrixValue".
func (this DOMMatrix) SetMatrixValue(transformList js.String) (ret DOMMatrix) {
	bindings.CallDOMMatrixSetMatrixValue(
		this.Ref(), js.Pointer(&ret),
		transformList.Ref(),
	)

	return
}

// TrySetMatrixValue calls the method "DOMMatrix.setMatrixValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMMatrix) TrySetMatrixValue(transformList js.String) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSetMatrixValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		transformList.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGTransform) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGTransformType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "SVGTransform.matrix".
//
// It returns ok=false if there is no such property.
func (this SVGTransform) Matrix() (ret DOMMatrix, ok bool) {
	ok = js.True == bindings.GetSVGTransformMatrix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Angle returns the value of property "SVGTransform.angle".
//
// It returns ok=false if there is no such property.
func (this SVGTransform) Angle() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGTransformAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetMatrix returns true if the method "SVGTransform.setMatrix" exists.
func (this SVGTransform) HasSetMatrix() bool {
	return js.True == bindings.HasSVGTransformSetMatrix(
		this.Ref(),
	)
}

// SetMatrixFunc returns the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.SVGTransformSetMatrixFunc(
			this.Ref(),
		),
	)
}

// SetMatrix calls the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrix(matrix DOMMatrix2DInit) (ret js.Void) {
	bindings.CallSVGTransformSetMatrix(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TrySetMatrix calls the method "SVGTransform.setMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetMatrix(matrix DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasSetMatrix1 returns true if the method "SVGTransform.setMatrix" exists.
func (this SVGTransform) HasSetMatrix1() bool {
	return js.True == bindings.HasSVGTransformSetMatrix1(
		this.Ref(),
	)
}

// SetMatrix1Func returns the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrix1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGTransformSetMatrix1Func(
			this.Ref(),
		),
	)
}

// SetMatrix1 calls the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrix1() (ret js.Void) {
	bindings.CallSVGTransformSetMatrix1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetMatrix1 calls the method "SVGTransform.setMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetMatrix1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetMatrix1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetTranslate returns true if the method "SVGTransform.setTranslate" exists.
func (this SVGTransform) HasSetTranslate() bool {
	return js.True == bindings.HasSVGTransformSetTranslate(
		this.Ref(),
	)
}

// SetTranslateFunc returns the method "SVGTransform.setTranslate".
func (this SVGTransform) SetTranslateFunc() (fn js.Func[func(tx float32, ty float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetTranslateFunc(
			this.Ref(),
		),
	)
}

// SetTranslate calls the method "SVGTransform.setTranslate".
func (this SVGTransform) SetTranslate(tx float32, ty float32) (ret js.Void) {
	bindings.CallSVGTransformSetTranslate(
		this.Ref(), js.Pointer(&ret),
		float32(tx),
		float32(ty),
	)

	return
}

// TrySetTranslate calls the method "SVGTransform.setTranslate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetTranslate(tx float32, ty float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetTranslate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(tx),
		float32(ty),
	)

	return
}

// HasSetScale returns true if the method "SVGTransform.setScale" exists.
func (this SVGTransform) HasSetScale() bool {
	return js.True == bindings.HasSVGTransformSetScale(
		this.Ref(),
	)
}

// SetScaleFunc returns the method "SVGTransform.setScale".
func (this SVGTransform) SetScaleFunc() (fn js.Func[func(sx float32, sy float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetScaleFunc(
			this.Ref(),
		),
	)
}

// SetScale calls the method "SVGTransform.setScale".
func (this SVGTransform) SetScale(sx float32, sy float32) (ret js.Void) {
	bindings.CallSVGTransformSetScale(
		this.Ref(), js.Pointer(&ret),
		float32(sx),
		float32(sy),
	)

	return
}

// TrySetScale calls the method "SVGTransform.setScale"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetScale(sx float32, sy float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetScale(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(sx),
		float32(sy),
	)

	return
}

// HasSetRotate returns true if the method "SVGTransform.setRotate" exists.
func (this SVGTransform) HasSetRotate() bool {
	return js.True == bindings.HasSVGTransformSetRotate(
		this.Ref(),
	)
}

// SetRotateFunc returns the method "SVGTransform.setRotate".
func (this SVGTransform) SetRotateFunc() (fn js.Func[func(angle float32, cx float32, cy float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetRotateFunc(
			this.Ref(),
		),
	)
}

// SetRotate calls the method "SVGTransform.setRotate".
func (this SVGTransform) SetRotate(angle float32, cx float32, cy float32) (ret js.Void) {
	bindings.CallSVGTransformSetRotate(
		this.Ref(), js.Pointer(&ret),
		float32(angle),
		float32(cx),
		float32(cy),
	)

	return
}

// TrySetRotate calls the method "SVGTransform.setRotate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetRotate(angle float32, cx float32, cy float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetRotate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(angle),
		float32(cx),
		float32(cy),
	)

	return
}

// HasSetSkewX returns true if the method "SVGTransform.setSkewX" exists.
func (this SVGTransform) HasSetSkewX() bool {
	return js.True == bindings.HasSVGTransformSetSkewX(
		this.Ref(),
	)
}

// SetSkewXFunc returns the method "SVGTransform.setSkewX".
func (this SVGTransform) SetSkewXFunc() (fn js.Func[func(angle float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetSkewXFunc(
			this.Ref(),
		),
	)
}

// SetSkewX calls the method "SVGTransform.setSkewX".
func (this SVGTransform) SetSkewX(angle float32) (ret js.Void) {
	bindings.CallSVGTransformSetSkewX(
		this.Ref(), js.Pointer(&ret),
		float32(angle),
	)

	return
}

// TrySetSkewX calls the method "SVGTransform.setSkewX"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetSkewX(angle float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetSkewX(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(angle),
	)

	return
}

// HasSetSkewY returns true if the method "SVGTransform.setSkewY" exists.
func (this SVGTransform) HasSetSkewY() bool {
	return js.True == bindings.HasSVGTransformSetSkewY(
		this.Ref(),
	)
}

// SetSkewYFunc returns the method "SVGTransform.setSkewY".
func (this SVGTransform) SetSkewYFunc() (fn js.Func[func(angle float32)]) {
	return fn.FromRef(
		bindings.SVGTransformSetSkewYFunc(
			this.Ref(),
		),
	)
}

// SetSkewY calls the method "SVGTransform.setSkewY".
func (this SVGTransform) SetSkewY(angle float32) (ret js.Void) {
	bindings.CallSVGTransformSetSkewY(
		this.Ref(), js.Pointer(&ret),
		float32(angle),
	)

	return
}

// TrySetSkewY calls the method "SVGTransform.setSkewY"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransform) TrySetSkewY(angle float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetSkewY(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(angle),
	)

	return
}

func NewDOMPointReadOnly(x float64, y float64, z float64, w float64) (ret DOMPointReadOnly) {
	ret.ref = bindings.NewDOMPointReadOnlyByDOMPointReadOnly(
		float64(x),
		float64(y),
		float64(z),
		float64(w))
	return
}

func NewDOMPointReadOnlyByDOMPointReadOnly1(x float64, y float64, z float64) (ret DOMPointReadOnly) {
	ret.ref = bindings.NewDOMPointReadOnlyByDOMPointReadOnly1(
		float64(x),
		float64(y),
		float64(z))
	return
}

func NewDOMPointReadOnlyByDOMPointReadOnly2(x float64, y float64) (ret DOMPointReadOnly) {
	ret.ref = bindings.NewDOMPointReadOnlyByDOMPointReadOnly2(
		float64(x),
		float64(y))
	return
}

func NewDOMPointReadOnlyByDOMPointReadOnly3(x float64) (ret DOMPointReadOnly) {
	ret.ref = bindings.NewDOMPointReadOnlyByDOMPointReadOnly3(
		float64(x))
	return
}

func NewDOMPointReadOnlyByDOMPointReadOnly4() (ret DOMPointReadOnly) {
	ret.ref = bindings.NewDOMPointReadOnlyByDOMPointReadOnly4()
	return
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
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DOMPointReadOnly.y".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "DOMPointReadOnly.z".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// W returns the value of property "DOMPointReadOnly.w".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) W() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyW(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFromPoint returns true if the staticmethod "DOMPointReadOnly.fromPoint" exists.
func (this DOMPointReadOnly) HasFromPoint() bool {
	return js.True == bindings.HasDOMPointReadOnlyFromPoint(
		this.Ref(),
	)
}

// FromPointFunc returns the staticmethod "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPointFunc() (fn js.Func[func(other DOMPointInit) DOMPointReadOnly]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyFromPointFunc(
			this.Ref(),
		),
	)
}

// FromPoint calls the staticmethod "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPoint(other DOMPointInit) (ret DOMPointReadOnly) {
	bindings.CallDOMPointReadOnlyFromPoint(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromPoint calls the staticmethod "DOMPointReadOnly.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPointReadOnly) TryFromPoint(other DOMPointInit) (ret DOMPointReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyFromPoint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromPoint1 returns true if the staticmethod "DOMPointReadOnly.fromPoint" exists.
func (this DOMPointReadOnly) HasFromPoint1() bool {
	return js.True == bindings.HasDOMPointReadOnlyFromPoint1(
		this.Ref(),
	)
}

// FromPoint1Func returns the staticmethod "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPoint1Func() (fn js.Func[func() DOMPointReadOnly]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyFromPoint1Func(
			this.Ref(),
		),
	)
}

// FromPoint1 calls the staticmethod "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPoint1() (ret DOMPointReadOnly) {
	bindings.CallDOMPointReadOnlyFromPoint1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromPoint1 calls the staticmethod "DOMPointReadOnly.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPointReadOnly) TryFromPoint1() (ret DOMPointReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyFromPoint1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMatrixTransform returns true if the method "DOMPointReadOnly.matrixTransform" exists.
func (this DOMPointReadOnly) HasMatrixTransform() bool {
	return js.True == bindings.HasDOMPointReadOnlyMatrixTransform(
		this.Ref(),
	)
}

// MatrixTransformFunc returns the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransformFunc() (fn js.Func[func(matrix DOMMatrixInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyMatrixTransformFunc(
			this.Ref(),
		),
	)
}

// MatrixTransform calls the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransform(matrix DOMMatrixInit) (ret DOMPoint) {
	bindings.CallDOMPointReadOnlyMatrixTransform(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryMatrixTransform calls the method "DOMPointReadOnly.matrixTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPointReadOnly) TryMatrixTransform(matrix DOMMatrixInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyMatrixTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasMatrixTransform1 returns true if the method "DOMPointReadOnly.matrixTransform" exists.
func (this DOMPointReadOnly) HasMatrixTransform1() bool {
	return js.True == bindings.HasDOMPointReadOnlyMatrixTransform1(
		this.Ref(),
	)
}

// MatrixTransform1Func returns the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransform1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyMatrixTransform1Func(
			this.Ref(),
		),
	)
}

// MatrixTransform1 calls the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransform1() (ret DOMPoint) {
	bindings.CallDOMPointReadOnlyMatrixTransform1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMatrixTransform1 calls the method "DOMPointReadOnly.matrixTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPointReadOnly) TryMatrixTransform1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyMatrixTransform1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "DOMPointReadOnly.toJSON" exists.
func (this DOMPointReadOnly) HasToJSON() bool {
	return js.True == bindings.HasDOMPointReadOnlyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "DOMPointReadOnly.toJSON".
func (this DOMPointReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMPointReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMPointReadOnly.toJSON".
func (this DOMPointReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMPointReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMPointReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPointReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGAnimatedRect) BaseVal() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedRectBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedRect.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedRect) AnimVal() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedRectAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGPreserveAspectRatio) Align() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGPreserveAspectRatioAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "SVGPreserveAspectRatio.align" to val.
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
// It returns ok=false if there is no such property.
func (this SVGPreserveAspectRatio) MeetOrSlice() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGPreserveAspectRatioMeetOrSlice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMeetOrSlice sets the value of property "SVGPreserveAspectRatio.meetOrSlice" to val.
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
// It returns ok=false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) BaseVal() (ret SVGPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedPreserveAspectRatioBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedPreserveAspectRatio.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) AnimVal() (ret SVGPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedPreserveAspectRatioAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGSVGElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGSVGElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGSVGElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGSVGElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentScale returns the value of property "SVGSVGElement.currentScale".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) CurrentScale() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementCurrentScale(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCurrentScale sets the value of property "SVGSVGElement.currentScale" to val.
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
// It returns ok=false if there is no such property.
func (this SVGSVGElement) CurrentTranslate() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementCurrentTranslate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ViewBox returns the value of property "SVGSVGElement.viewBox".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementViewBox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGSVGElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetIntersectionList returns true if the method "SVGSVGElement.getIntersectionList" exists.
func (this SVGSVGElement) HasGetIntersectionList() bool {
	return js.True == bindings.HasSVGSVGElementGetIntersectionList(
		this.Ref(),
	)
}

// GetIntersectionListFunc returns the method "SVGSVGElement.getIntersectionList".
func (this SVGSVGElement) GetIntersectionListFunc() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetIntersectionListFunc(
			this.Ref(),
		),
	)
}

// GetIntersectionList calls the method "SVGSVGElement.getIntersectionList".
func (this SVGSVGElement) GetIntersectionList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList) {
	bindings.CallSVGSVGElementGetIntersectionList(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// TryGetIntersectionList calls the method "SVGSVGElement.getIntersectionList"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryGetIntersectionList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetIntersectionList(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// HasGetEnclosureList returns true if the method "SVGSVGElement.getEnclosureList" exists.
func (this SVGSVGElement) HasGetEnclosureList() bool {
	return js.True == bindings.HasSVGSVGElementGetEnclosureList(
		this.Ref(),
	)
}

// GetEnclosureListFunc returns the method "SVGSVGElement.getEnclosureList".
func (this SVGSVGElement) GetEnclosureListFunc() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetEnclosureListFunc(
			this.Ref(),
		),
	)
}

// GetEnclosureList calls the method "SVGSVGElement.getEnclosureList".
func (this SVGSVGElement) GetEnclosureList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList) {
	bindings.CallSVGSVGElementGetEnclosureList(
		this.Ref(), js.Pointer(&ret),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// TryGetEnclosureList calls the method "SVGSVGElement.getEnclosureList"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryGetEnclosureList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetEnclosureList(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// HasCheckIntersection returns true if the method "SVGSVGElement.checkIntersection" exists.
func (this SVGSVGElement) HasCheckIntersection() bool {
	return js.True == bindings.HasSVGSVGElementCheckIntersection(
		this.Ref(),
	)
}

// CheckIntersectionFunc returns the method "SVGSVGElement.checkIntersection".
func (this SVGSVGElement) CheckIntersectionFunc() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementCheckIntersectionFunc(
			this.Ref(),
		),
	)
}

// CheckIntersection calls the method "SVGSVGElement.checkIntersection".
func (this SVGSVGElement) CheckIntersection(element SVGElement, rect DOMRectReadOnly) (ret bool) {
	bindings.CallSVGSVGElementCheckIntersection(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// TryCheckIntersection calls the method "SVGSVGElement.checkIntersection"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCheckIntersection(element SVGElement, rect DOMRectReadOnly) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCheckIntersection(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// HasCheckEnclosure returns true if the method "SVGSVGElement.checkEnclosure" exists.
func (this SVGSVGElement) HasCheckEnclosure() bool {
	return js.True == bindings.HasSVGSVGElementCheckEnclosure(
		this.Ref(),
	)
}

// CheckEnclosureFunc returns the method "SVGSVGElement.checkEnclosure".
func (this SVGSVGElement) CheckEnclosureFunc() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementCheckEnclosureFunc(
			this.Ref(),
		),
	)
}

// CheckEnclosure calls the method "SVGSVGElement.checkEnclosure".
func (this SVGSVGElement) CheckEnclosure(element SVGElement, rect DOMRectReadOnly) (ret bool) {
	bindings.CallSVGSVGElementCheckEnclosure(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// TryCheckEnclosure calls the method "SVGSVGElement.checkEnclosure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCheckEnclosure(element SVGElement, rect DOMRectReadOnly) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCheckEnclosure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// HasDeselectAll returns true if the method "SVGSVGElement.deselectAll" exists.
func (this SVGSVGElement) HasDeselectAll() bool {
	return js.True == bindings.HasSVGSVGElementDeselectAll(
		this.Ref(),
	)
}

// DeselectAllFunc returns the method "SVGSVGElement.deselectAll".
func (this SVGSVGElement) DeselectAllFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementDeselectAllFunc(
			this.Ref(),
		),
	)
}

// DeselectAll calls the method "SVGSVGElement.deselectAll".
func (this SVGSVGElement) DeselectAll() (ret js.Void) {
	bindings.CallSVGSVGElementDeselectAll(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeselectAll calls the method "SVGSVGElement.deselectAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryDeselectAll() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementDeselectAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGNumber returns true if the method "SVGSVGElement.createSVGNumber" exists.
func (this SVGSVGElement) HasCreateSVGNumber() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGNumber(
		this.Ref(),
	)
}

// CreateSVGNumberFunc returns the method "SVGSVGElement.createSVGNumber".
func (this SVGSVGElement) CreateSVGNumberFunc() (fn js.Func[func() SVGNumber]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGNumberFunc(
			this.Ref(),
		),
	)
}

// CreateSVGNumber calls the method "SVGSVGElement.createSVGNumber".
func (this SVGSVGElement) CreateSVGNumber() (ret SVGNumber) {
	bindings.CallSVGSVGElementCreateSVGNumber(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGNumber calls the method "SVGSVGElement.createSVGNumber"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGNumber() (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGNumber(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGLength returns true if the method "SVGSVGElement.createSVGLength" exists.
func (this SVGSVGElement) HasCreateSVGLength() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGLength(
		this.Ref(),
	)
}

// CreateSVGLengthFunc returns the method "SVGSVGElement.createSVGLength".
func (this SVGSVGElement) CreateSVGLengthFunc() (fn js.Func[func() SVGLength]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGLengthFunc(
			this.Ref(),
		),
	)
}

// CreateSVGLength calls the method "SVGSVGElement.createSVGLength".
func (this SVGSVGElement) CreateSVGLength() (ret SVGLength) {
	bindings.CallSVGSVGElementCreateSVGLength(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGLength calls the method "SVGSVGElement.createSVGLength"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGLength() (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGLength(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGAngle returns true if the method "SVGSVGElement.createSVGAngle" exists.
func (this SVGSVGElement) HasCreateSVGAngle() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGAngle(
		this.Ref(),
	)
}

// CreateSVGAngleFunc returns the method "SVGSVGElement.createSVGAngle".
func (this SVGSVGElement) CreateSVGAngleFunc() (fn js.Func[func() SVGAngle]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGAngleFunc(
			this.Ref(),
		),
	)
}

// CreateSVGAngle calls the method "SVGSVGElement.createSVGAngle".
func (this SVGSVGElement) CreateSVGAngle() (ret SVGAngle) {
	bindings.CallSVGSVGElementCreateSVGAngle(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGAngle calls the method "SVGSVGElement.createSVGAngle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGAngle() (ret SVGAngle, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGAngle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGPoint returns true if the method "SVGSVGElement.createSVGPoint" exists.
func (this SVGSVGElement) HasCreateSVGPoint() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGPoint(
		this.Ref(),
	)
}

// CreateSVGPointFunc returns the method "SVGSVGElement.createSVGPoint".
func (this SVGSVGElement) CreateSVGPointFunc() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGPointFunc(
			this.Ref(),
		),
	)
}

// CreateSVGPoint calls the method "SVGSVGElement.createSVGPoint".
func (this SVGSVGElement) CreateSVGPoint() (ret DOMPoint) {
	bindings.CallSVGSVGElementCreateSVGPoint(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGPoint calls the method "SVGSVGElement.createSVGPoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGPoint() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGPoint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGMatrix returns true if the method "SVGSVGElement.createSVGMatrix" exists.
func (this SVGSVGElement) HasCreateSVGMatrix() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGMatrix(
		this.Ref(),
	)
}

// CreateSVGMatrixFunc returns the method "SVGSVGElement.createSVGMatrix".
func (this SVGSVGElement) CreateSVGMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGMatrix calls the method "SVGSVGElement.createSVGMatrix".
func (this SVGSVGElement) CreateSVGMatrix() (ret DOMMatrix) {
	bindings.CallSVGSVGElementCreateSVGMatrix(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGMatrix calls the method "SVGSVGElement.createSVGMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGMatrix() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGRect returns true if the method "SVGSVGElement.createSVGRect" exists.
func (this SVGSVGElement) HasCreateSVGRect() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGRect(
		this.Ref(),
	)
}

// CreateSVGRectFunc returns the method "SVGSVGElement.createSVGRect".
func (this SVGSVGElement) CreateSVGRectFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGRectFunc(
			this.Ref(),
		),
	)
}

// CreateSVGRect calls the method "SVGSVGElement.createSVGRect".
func (this SVGSVGElement) CreateSVGRect() (ret DOMRect) {
	bindings.CallSVGSVGElementCreateSVGRect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGRect calls the method "SVGSVGElement.createSVGRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGTransform returns true if the method "SVGSVGElement.createSVGTransform" exists.
func (this SVGSVGElement) HasCreateSVGTransform() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGTransform(
		this.Ref(),
	)
}

// CreateSVGTransformFunc returns the method "SVGSVGElement.createSVGTransform".
func (this SVGSVGElement) CreateSVGTransformFunc() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransform calls the method "SVGSVGElement.createSVGTransform".
func (this SVGSVGElement) CreateSVGTransform() (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransform(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransform calls the method "SVGSVGElement.createSVGTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransform() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSVGTransformFromMatrix returns true if the method "SVGSVGElement.createSVGTransformFromMatrix" exists.
func (this SVGSVGElement) HasCreateSVGTransformFromMatrix() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGTransformFromMatrix(
		this.Ref(),
	)
}

// CreateSVGTransformFromMatrixFunc returns the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFromMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix calls the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryCreateSVGTransformFromMatrix calls the method "SVGSVGElement.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasCreateSVGTransformFromMatrix1 returns true if the method "SVGSVGElement.createSVGTransformFromMatrix" exists.
func (this SVGSVGElement) HasCreateSVGTransformFromMatrix1() bool {
	return js.True == bindings.HasSVGSVGElementCreateSVGTransformFromMatrix1(
		this.Ref(),
	)
}

// CreateSVGTransformFromMatrix1Func returns the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrix1Func() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGSVGElementCreateSVGTransformFromMatrix1Func(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix1 calls the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrix1() (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransformFromMatrix1 calls the method "SVGSVGElement.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransformFromMatrix1() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetElementById returns true if the method "SVGSVGElement.getElementById" exists.
func (this SVGSVGElement) HasGetElementById() bool {
	return js.True == bindings.HasSVGSVGElementGetElementById(
		this.Ref(),
	)
}

// GetElementByIdFunc returns the method "SVGSVGElement.getElementById".
func (this SVGSVGElement) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// GetElementById calls the method "SVGSVGElement.getElementById".
func (this SVGSVGElement) GetElementById(elementId js.String) (ret Element) {
	bindings.CallSVGSVGElementGetElementById(
		this.Ref(), js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "SVGSVGElement.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetElementById(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasSuspendRedraw returns true if the method "SVGSVGElement.suspendRedraw" exists.
func (this SVGSVGElement) HasSuspendRedraw() bool {
	return js.True == bindings.HasSVGSVGElementSuspendRedraw(
		this.Ref(),
	)
}

// SuspendRedrawFunc returns the method "SVGSVGElement.suspendRedraw".
func (this SVGSVGElement) SuspendRedrawFunc() (fn js.Func[func(maxWaitMilliseconds uint32) uint32]) {
	return fn.FromRef(
		bindings.SVGSVGElementSuspendRedrawFunc(
			this.Ref(),
		),
	)
}

// SuspendRedraw calls the method "SVGSVGElement.suspendRedraw".
func (this SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint32) (ret uint32) {
	bindings.CallSVGSVGElementSuspendRedraw(
		this.Ref(), js.Pointer(&ret),
		uint32(maxWaitMilliseconds),
	)

	return
}

// TrySuspendRedraw calls the method "SVGSVGElement.suspendRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TrySuspendRedraw(maxWaitMilliseconds uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementSuspendRedraw(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(maxWaitMilliseconds),
	)

	return
}

// HasUnsuspendRedraw returns true if the method "SVGSVGElement.unsuspendRedraw" exists.
func (this SVGSVGElement) HasUnsuspendRedraw() bool {
	return js.True == bindings.HasSVGSVGElementUnsuspendRedraw(
		this.Ref(),
	)
}

// UnsuspendRedrawFunc returns the method "SVGSVGElement.unsuspendRedraw".
func (this SVGSVGElement) UnsuspendRedrawFunc() (fn js.Func[func(suspendHandleID uint32)]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnsuspendRedrawFunc(
			this.Ref(),
		),
	)
}

// UnsuspendRedraw calls the method "SVGSVGElement.unsuspendRedraw".
func (this SVGSVGElement) UnsuspendRedraw(suspendHandleID uint32) (ret js.Void) {
	bindings.CallSVGSVGElementUnsuspendRedraw(
		this.Ref(), js.Pointer(&ret),
		uint32(suspendHandleID),
	)

	return
}

// TryUnsuspendRedraw calls the method "SVGSVGElement.unsuspendRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryUnsuspendRedraw(suspendHandleID uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnsuspendRedraw(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(suspendHandleID),
	)

	return
}

// HasUnsuspendRedrawAll returns true if the method "SVGSVGElement.unsuspendRedrawAll" exists.
func (this SVGSVGElement) HasUnsuspendRedrawAll() bool {
	return js.True == bindings.HasSVGSVGElementUnsuspendRedrawAll(
		this.Ref(),
	)
}

// UnsuspendRedrawAllFunc returns the method "SVGSVGElement.unsuspendRedrawAll".
func (this SVGSVGElement) UnsuspendRedrawAllFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnsuspendRedrawAllFunc(
			this.Ref(),
		),
	)
}

// UnsuspendRedrawAll calls the method "SVGSVGElement.unsuspendRedrawAll".
func (this SVGSVGElement) UnsuspendRedrawAll() (ret js.Void) {
	bindings.CallSVGSVGElementUnsuspendRedrawAll(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnsuspendRedrawAll calls the method "SVGSVGElement.unsuspendRedrawAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryUnsuspendRedrawAll() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnsuspendRedrawAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasForceRedraw returns true if the method "SVGSVGElement.forceRedraw" exists.
func (this SVGSVGElement) HasForceRedraw() bool {
	return js.True == bindings.HasSVGSVGElementForceRedraw(
		this.Ref(),
	)
}

// ForceRedrawFunc returns the method "SVGSVGElement.forceRedraw".
func (this SVGSVGElement) ForceRedrawFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementForceRedrawFunc(
			this.Ref(),
		),
	)
}

// ForceRedraw calls the method "SVGSVGElement.forceRedraw".
func (this SVGSVGElement) ForceRedraw() (ret js.Void) {
	bindings.CallSVGSVGElementForceRedraw(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryForceRedraw calls the method "SVGSVGElement.forceRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryForceRedraw() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementForceRedraw(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPauseAnimations returns true if the method "SVGSVGElement.pauseAnimations" exists.
func (this SVGSVGElement) HasPauseAnimations() bool {
	return js.True == bindings.HasSVGSVGElementPauseAnimations(
		this.Ref(),
	)
}

// PauseAnimationsFunc returns the method "SVGSVGElement.pauseAnimations".
func (this SVGSVGElement) PauseAnimationsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementPauseAnimationsFunc(
			this.Ref(),
		),
	)
}

// PauseAnimations calls the method "SVGSVGElement.pauseAnimations".
func (this SVGSVGElement) PauseAnimations() (ret js.Void) {
	bindings.CallSVGSVGElementPauseAnimations(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPauseAnimations calls the method "SVGSVGElement.pauseAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryPauseAnimations() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementPauseAnimations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnpauseAnimations returns true if the method "SVGSVGElement.unpauseAnimations" exists.
func (this SVGSVGElement) HasUnpauseAnimations() bool {
	return js.True == bindings.HasSVGSVGElementUnpauseAnimations(
		this.Ref(),
	)
}

// UnpauseAnimationsFunc returns the method "SVGSVGElement.unpauseAnimations".
func (this SVGSVGElement) UnpauseAnimationsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGSVGElementUnpauseAnimationsFunc(
			this.Ref(),
		),
	)
}

// UnpauseAnimations calls the method "SVGSVGElement.unpauseAnimations".
func (this SVGSVGElement) UnpauseAnimations() (ret js.Void) {
	bindings.CallSVGSVGElementUnpauseAnimations(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnpauseAnimations calls the method "SVGSVGElement.unpauseAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryUnpauseAnimations() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnpauseAnimations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAnimationsPaused returns true if the method "SVGSVGElement.animationsPaused" exists.
func (this SVGSVGElement) HasAnimationsPaused() bool {
	return js.True == bindings.HasSVGSVGElementAnimationsPaused(
		this.Ref(),
	)
}

// AnimationsPausedFunc returns the method "SVGSVGElement.animationsPaused".
func (this SVGSVGElement) AnimationsPausedFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGSVGElementAnimationsPausedFunc(
			this.Ref(),
		),
	)
}

// AnimationsPaused calls the method "SVGSVGElement.animationsPaused".
func (this SVGSVGElement) AnimationsPaused() (ret bool) {
	bindings.CallSVGSVGElementAnimationsPaused(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAnimationsPaused calls the method "SVGSVGElement.animationsPaused"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryAnimationsPaused() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementAnimationsPaused(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetCurrentTime returns true if the method "SVGSVGElement.getCurrentTime" exists.
func (this SVGSVGElement) HasGetCurrentTime() bool {
	return js.True == bindings.HasSVGSVGElementGetCurrentTime(
		this.Ref(),
	)
}

// GetCurrentTimeFunc returns the method "SVGSVGElement.getCurrentTime".
func (this SVGSVGElement) GetCurrentTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGSVGElementGetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTime calls the method "SVGSVGElement.getCurrentTime".
func (this SVGSVGElement) GetCurrentTime() (ret float32) {
	bindings.CallSVGSVGElementGetCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTime calls the method "SVGSVGElement.getCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TryGetCurrentTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetCurrentTime(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCurrentTime returns true if the method "SVGSVGElement.setCurrentTime" exists.
func (this SVGSVGElement) HasSetCurrentTime() bool {
	return js.True == bindings.HasSVGSVGElementSetCurrentTime(
		this.Ref(),
	)
}

// SetCurrentTimeFunc returns the method "SVGSVGElement.setCurrentTime".
func (this SVGSVGElement) SetCurrentTimeFunc() (fn js.Func[func(seconds float32)]) {
	return fn.FromRef(
		bindings.SVGSVGElementSetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

// SetCurrentTime calls the method "SVGSVGElement.setCurrentTime".
func (this SVGSVGElement) SetCurrentTime(seconds float32) (ret js.Void) {
	bindings.CallSVGSVGElementSetCurrentTime(
		this.Ref(), js.Pointer(&ret),
		float32(seconds),
	)

	return
}

// TrySetCurrentTime calls the method "SVGSVGElement.setCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGSVGElement) TrySetCurrentTime(seconds float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementSetCurrentTime(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(seconds),
	)

	return
}

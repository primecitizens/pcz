// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type Baseline struct {
	ref js.Ref
}

func (this Baseline) Once() Baseline {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "Baseline.name".
//
// It returns ok=false if there is no such property.
func (this Baseline) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBaselineName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "Baseline.value".
//
// It returns ok=false if there is no such property.
func (this Baseline) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetBaselineValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Font struct {
	ref js.Ref
}

func (this Font) Once() Font {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "Font.name".
//
// It returns ok=false if there is no such property.
func (this Font) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// GlyphsRendered returns the value of property "Font.glyphsRendered".
//
// It returns ok=false if there is no such property.
func (this Font) GlyphsRendered() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontGlyphsRendered(
		this.ref, js.Pointer(&ret),
	)
	return
}

type FontMetrics struct {
	ref js.Ref
}

func (this FontMetrics) Once() FontMetrics {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "FontMetrics.width".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Advances returns the value of property "FontMetrics.advances".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Advances() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetFontMetricsAdvances(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BoundingBoxLeft returns the value of property "FontMetrics.boundingBoxLeft".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BoundingBoxRight returns the value of property "FontMetrics.boundingBoxRight".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxRight() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxRight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "FontMetrics.height".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EmHeightAscent returns the value of property "FontMetrics.emHeightAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) EmHeightAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsEmHeightAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EmHeightDescent returns the value of property "FontMetrics.emHeightDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) EmHeightDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsEmHeightDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BoundingBoxAscent returns the value of property "FontMetrics.boundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BoundingBoxDescent returns the value of property "FontMetrics.boundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) BoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsBoundingBoxDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxAscent returns the value of property "FontMetrics.fontBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) FontBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsFontBoundingBoxAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxDescent returns the value of property "FontMetrics.fontBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) FontBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontMetricsFontBoundingBoxDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DominantBaseline returns the value of property "FontMetrics.dominantBaseline".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) DominantBaseline() (ret Baseline, ok bool) {
	ok = js.True == bindings.GetFontMetricsDominantBaseline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Baselines returns the value of property "FontMetrics.baselines".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Baselines() (ret js.FrozenArray[Baseline], ok bool) {
	ok = js.True == bindings.GetFontMetricsBaselines(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "FontMetrics.fonts".
//
// It returns ok=false if there is no such property.
func (this FontMetrics) Fonts() (ret js.FrozenArray[Font], ok bool) {
	ok = js.True == bindings.GetFontMetricsFonts(
		this.ref, js.Pointer(&ret),
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
func (p *StaticRangeInit) UpdateFrom(ref js.Ref) {
	bindings.StaticRangeInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StaticRangeInit) Update(ref js.Ref) {
	bindings.StaticRangeInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StaticRangeInit) FreeMembers(recursive bool) {
	js.Free(
		p.StartContainer.Ref(),
		p.EndContainer.Ref(),
	)
	p.StartContainer = p.StartContainer.FromRef(js.Undefined)
	p.EndContainer = p.EndContainer.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

type Selection struct {
	ref js.Ref
}

func (this Selection) Once() Selection {
	this.ref.Once()
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
	this.ref.Free()
}

// AnchorNode returns the value of property "Selection.anchorNode".
//
// It returns ok=false if there is no such property.
func (this Selection) AnchorNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetSelectionAnchorNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnchorOffset returns the value of property "Selection.anchorOffset".
//
// It returns ok=false if there is no such property.
func (this Selection) AnchorOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionAnchorOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FocusNode returns the value of property "Selection.focusNode".
//
// It returns ok=false if there is no such property.
func (this Selection) FocusNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetSelectionFocusNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FocusOffset returns the value of property "Selection.focusOffset".
//
// It returns ok=false if there is no such property.
func (this Selection) FocusOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionFocusOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsCollapsed returns the value of property "Selection.isCollapsed".
//
// It returns ok=false if there is no such property.
func (this Selection) IsCollapsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetSelectionIsCollapsed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RangeCount returns the value of property "Selection.rangeCount".
//
// It returns ok=false if there is no such property.
func (this Selection) RangeCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSelectionRangeCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Selection.type".
//
// It returns ok=false if there is no such property.
func (this Selection) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSelectionType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "Selection.direction".
//
// It returns ok=false if there is no such property.
func (this Selection) Direction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSelectionDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetRangeAt returns true if the method "Selection.getRangeAt" exists.
func (this Selection) HasFuncGetRangeAt() bool {
	return js.True == bindings.HasFuncSelectionGetRangeAt(
		this.ref,
	)
}

// FuncGetRangeAt returns the method "Selection.getRangeAt".
func (this Selection) FuncGetRangeAt() (fn js.Func[func(index uint32) Range]) {
	bindings.FuncSelectionGetRangeAt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRangeAt calls the method "Selection.getRangeAt".
func (this Selection) GetRangeAt(index uint32) (ret Range) {
	bindings.CallSelectionGetRangeAt(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetRangeAt calls the method "Selection.getRangeAt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryGetRangeAt(index uint32) (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionGetRangeAt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAddRange returns true if the method "Selection.addRange" exists.
func (this Selection) HasFuncAddRange() bool {
	return js.True == bindings.HasFuncSelectionAddRange(
		this.ref,
	)
}

// FuncAddRange returns the method "Selection.addRange".
func (this Selection) FuncAddRange() (fn js.Func[func(rng Range)]) {
	bindings.FuncSelectionAddRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRange calls the method "Selection.addRange".
func (this Selection) AddRange(rng Range) (ret js.Void) {
	bindings.CallSelectionAddRange(
		this.ref, js.Pointer(&ret),
		rng.Ref(),
	)

	return
}

// TryAddRange calls the method "Selection.addRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryAddRange(rng Range) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionAddRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rng.Ref(),
	)

	return
}

// HasFuncRemoveRange returns true if the method "Selection.removeRange" exists.
func (this Selection) HasFuncRemoveRange() bool {
	return js.True == bindings.HasFuncSelectionRemoveRange(
		this.ref,
	)
}

// FuncRemoveRange returns the method "Selection.removeRange".
func (this Selection) FuncRemoveRange() (fn js.Func[func(rng Range)]) {
	bindings.FuncSelectionRemoveRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRange calls the method "Selection.removeRange".
func (this Selection) RemoveRange(rng Range) (ret js.Void) {
	bindings.CallSelectionRemoveRange(
		this.ref, js.Pointer(&ret),
		rng.Ref(),
	)

	return
}

// TryRemoveRange calls the method "Selection.removeRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryRemoveRange(rng Range) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionRemoveRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rng.Ref(),
	)

	return
}

// HasFuncRemoveAllRanges returns true if the method "Selection.removeAllRanges" exists.
func (this Selection) HasFuncRemoveAllRanges() bool {
	return js.True == bindings.HasFuncSelectionRemoveAllRanges(
		this.ref,
	)
}

// FuncRemoveAllRanges returns the method "Selection.removeAllRanges".
func (this Selection) FuncRemoveAllRanges() (fn js.Func[func()]) {
	bindings.FuncSelectionRemoveAllRanges(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAllRanges calls the method "Selection.removeAllRanges".
func (this Selection) RemoveAllRanges() (ret js.Void) {
	bindings.CallSelectionRemoveAllRanges(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemoveAllRanges calls the method "Selection.removeAllRanges"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryRemoveAllRanges() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionRemoveAllRanges(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEmpty returns true if the method "Selection.empty" exists.
func (this Selection) HasFuncEmpty() bool {
	return js.True == bindings.HasFuncSelectionEmpty(
		this.ref,
	)
}

// FuncEmpty returns the method "Selection.empty".
func (this Selection) FuncEmpty() (fn js.Func[func()]) {
	bindings.FuncSelectionEmpty(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Empty calls the method "Selection.empty".
func (this Selection) Empty() (ret js.Void) {
	bindings.CallSelectionEmpty(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEmpty calls the method "Selection.empty"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryEmpty() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionEmpty(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetComposedRanges returns true if the method "Selection.getComposedRanges" exists.
func (this Selection) HasFuncGetComposedRanges() bool {
	return js.True == bindings.HasFuncSelectionGetComposedRanges(
		this.ref,
	)
}

// FuncGetComposedRanges returns the method "Selection.getComposedRanges".
func (this Selection) FuncGetComposedRanges() (fn js.Func[func(shadowRoots ...ShadowRoot) js.Array[StaticRange]]) {
	bindings.FuncSelectionGetComposedRanges(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetComposedRanges calls the method "Selection.getComposedRanges".
func (this Selection) GetComposedRanges(shadowRoots ...ShadowRoot) (ret js.Array[StaticRange]) {
	bindings.CallSelectionGetComposedRanges(
		this.ref, js.Pointer(&ret),
		js.SliceData(shadowRoots),
		js.SizeU(len(shadowRoots)),
	)

	return
}

// TryGetComposedRanges calls the method "Selection.getComposedRanges"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryGetComposedRanges(shadowRoots ...ShadowRoot) (ret js.Array[StaticRange], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionGetComposedRanges(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(shadowRoots),
		js.SizeU(len(shadowRoots)),
	)

	return
}

// HasFuncCollapse returns true if the method "Selection.collapse" exists.
func (this Selection) HasFuncCollapse() bool {
	return js.True == bindings.HasFuncSelectionCollapse(
		this.ref,
	)
}

// FuncCollapse returns the method "Selection.collapse".
func (this Selection) FuncCollapse() (fn js.Func[func(node Node, offset uint32)]) {
	bindings.FuncSelectionCollapse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Collapse calls the method "Selection.collapse".
func (this Selection) Collapse(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionCollapse(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryCollapse calls the method "Selection.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryCollapse(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncCollapse1 returns true if the method "Selection.collapse" exists.
func (this Selection) HasFuncCollapse1() bool {
	return js.True == bindings.HasFuncSelectionCollapse1(
		this.ref,
	)
}

// FuncCollapse1 returns the method "Selection.collapse".
func (this Selection) FuncCollapse1() (fn js.Func[func(node Node)]) {
	bindings.FuncSelectionCollapse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Collapse1 calls the method "Selection.collapse".
func (this Selection) Collapse1(node Node) (ret js.Void) {
	bindings.CallSelectionCollapse1(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryCollapse1 calls the method "Selection.collapse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryCollapse1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSetPosition returns true if the method "Selection.setPosition" exists.
func (this Selection) HasFuncSetPosition() bool {
	return js.True == bindings.HasFuncSelectionSetPosition(
		this.ref,
	)
}

// FuncSetPosition returns the method "Selection.setPosition".
func (this Selection) FuncSetPosition() (fn js.Func[func(node Node, offset uint32)]) {
	bindings.FuncSelectionSetPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPosition calls the method "Selection.setPosition".
func (this Selection) SetPosition(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionSetPosition(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TrySetPosition calls the method "Selection.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TrySetPosition(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncSetPosition1 returns true if the method "Selection.setPosition" exists.
func (this Selection) HasFuncSetPosition1() bool {
	return js.True == bindings.HasFuncSelectionSetPosition1(
		this.ref,
	)
}

// FuncSetPosition1 returns the method "Selection.setPosition".
func (this Selection) FuncSetPosition1() (fn js.Func[func(node Node)]) {
	bindings.FuncSelectionSetPosition1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPosition1 calls the method "Selection.setPosition".
func (this Selection) SetPosition1(node Node) (ret js.Void) {
	bindings.CallSelectionSetPosition1(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySetPosition1 calls the method "Selection.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TrySetPosition1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetPosition1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncCollapseToStart returns true if the method "Selection.collapseToStart" exists.
func (this Selection) HasFuncCollapseToStart() bool {
	return js.True == bindings.HasFuncSelectionCollapseToStart(
		this.ref,
	)
}

// FuncCollapseToStart returns the method "Selection.collapseToStart".
func (this Selection) FuncCollapseToStart() (fn js.Func[func()]) {
	bindings.FuncSelectionCollapseToStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CollapseToStart calls the method "Selection.collapseToStart".
func (this Selection) CollapseToStart() (ret js.Void) {
	bindings.CallSelectionCollapseToStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCollapseToStart calls the method "Selection.collapseToStart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryCollapseToStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapseToStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCollapseToEnd returns true if the method "Selection.collapseToEnd" exists.
func (this Selection) HasFuncCollapseToEnd() bool {
	return js.True == bindings.HasFuncSelectionCollapseToEnd(
		this.ref,
	)
}

// FuncCollapseToEnd returns the method "Selection.collapseToEnd".
func (this Selection) FuncCollapseToEnd() (fn js.Func[func()]) {
	bindings.FuncSelectionCollapseToEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CollapseToEnd calls the method "Selection.collapseToEnd".
func (this Selection) CollapseToEnd() (ret js.Void) {
	bindings.CallSelectionCollapseToEnd(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCollapseToEnd calls the method "Selection.collapseToEnd"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryCollapseToEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionCollapseToEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExtend returns true if the method "Selection.extend" exists.
func (this Selection) HasFuncExtend() bool {
	return js.True == bindings.HasFuncSelectionExtend(
		this.ref,
	)
}

// FuncExtend returns the method "Selection.extend".
func (this Selection) FuncExtend() (fn js.Func[func(node Node, offset uint32)]) {
	bindings.FuncSelectionExtend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Extend calls the method "Selection.extend".
func (this Selection) Extend(node Node, offset uint32) (ret js.Void) {
	bindings.CallSelectionExtend(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		uint32(offset),
	)

	return
}

// TryExtend calls the method "Selection.extend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryExtend(node Node, offset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionExtend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		uint32(offset),
	)

	return
}

// HasFuncExtend1 returns true if the method "Selection.extend" exists.
func (this Selection) HasFuncExtend1() bool {
	return js.True == bindings.HasFuncSelectionExtend1(
		this.ref,
	)
}

// FuncExtend1 returns the method "Selection.extend".
func (this Selection) FuncExtend1() (fn js.Func[func(node Node)]) {
	bindings.FuncSelectionExtend1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Extend1 calls the method "Selection.extend".
func (this Selection) Extend1(node Node) (ret js.Void) {
	bindings.CallSelectionExtend1(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryExtend1 calls the method "Selection.extend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryExtend1(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionExtend1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncSetBaseAndExtent returns true if the method "Selection.setBaseAndExtent" exists.
func (this Selection) HasFuncSetBaseAndExtent() bool {
	return js.True == bindings.HasFuncSelectionSetBaseAndExtent(
		this.ref,
	)
}

// FuncSetBaseAndExtent returns the method "Selection.setBaseAndExtent".
func (this Selection) FuncSetBaseAndExtent() (fn js.Func[func(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32)]) {
	bindings.FuncSelectionSetBaseAndExtent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBaseAndExtent calls the method "Selection.setBaseAndExtent".
func (this Selection) SetBaseAndExtent(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32) (ret js.Void) {
	bindings.CallSelectionSetBaseAndExtent(
		this.ref, js.Pointer(&ret),
		anchorNode.Ref(),
		uint32(anchorOffset),
		focusNode.Ref(),
		uint32(focusOffset),
	)

	return
}

// TrySetBaseAndExtent calls the method "Selection.setBaseAndExtent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TrySetBaseAndExtent(anchorNode Node, anchorOffset uint32, focusNode Node, focusOffset uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSetBaseAndExtent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		anchorNode.Ref(),
		uint32(anchorOffset),
		focusNode.Ref(),
		uint32(focusOffset),
	)

	return
}

// HasFuncSelectAllChildren returns true if the method "Selection.selectAllChildren" exists.
func (this Selection) HasFuncSelectAllChildren() bool {
	return js.True == bindings.HasFuncSelectionSelectAllChildren(
		this.ref,
	)
}

// FuncSelectAllChildren returns the method "Selection.selectAllChildren".
func (this Selection) FuncSelectAllChildren() (fn js.Func[func(node Node)]) {
	bindings.FuncSelectionSelectAllChildren(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SelectAllChildren calls the method "Selection.selectAllChildren".
func (this Selection) SelectAllChildren(node Node) (ret js.Void) {
	bindings.CallSelectionSelectAllChildren(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TrySelectAllChildren calls the method "Selection.selectAllChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TrySelectAllChildren(node Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionSelectAllChildren(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncModify returns true if the method "Selection.modify" exists.
func (this Selection) HasFuncModify() bool {
	return js.True == bindings.HasFuncSelectionModify(
		this.ref,
	)
}

// FuncModify returns the method "Selection.modify".
func (this Selection) FuncModify() (fn js.Func[func(alter js.String, direction js.String, granularity js.String)]) {
	bindings.FuncSelectionModify(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Modify calls the method "Selection.modify".
func (this Selection) Modify(alter js.String, direction js.String, granularity js.String) (ret js.Void) {
	bindings.CallSelectionModify(
		this.ref, js.Pointer(&ret),
		alter.Ref(),
		direction.Ref(),
		granularity.Ref(),
	)

	return
}

// TryModify calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryModify(alter js.String, direction js.String, granularity js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
		direction.Ref(),
		granularity.Ref(),
	)

	return
}

// HasFuncModify1 returns true if the method "Selection.modify" exists.
func (this Selection) HasFuncModify1() bool {
	return js.True == bindings.HasFuncSelectionModify1(
		this.ref,
	)
}

// FuncModify1 returns the method "Selection.modify".
func (this Selection) FuncModify1() (fn js.Func[func(alter js.String, direction js.String)]) {
	bindings.FuncSelectionModify1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Modify1 calls the method "Selection.modify".
func (this Selection) Modify1(alter js.String, direction js.String) (ret js.Void) {
	bindings.CallSelectionModify1(
		this.ref, js.Pointer(&ret),
		alter.Ref(),
		direction.Ref(),
	)

	return
}

// TryModify1 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryModify1(alter js.String, direction js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
		direction.Ref(),
	)

	return
}

// HasFuncModify2 returns true if the method "Selection.modify" exists.
func (this Selection) HasFuncModify2() bool {
	return js.True == bindings.HasFuncSelectionModify2(
		this.ref,
	)
}

// FuncModify2 returns the method "Selection.modify".
func (this Selection) FuncModify2() (fn js.Func[func(alter js.String)]) {
	bindings.FuncSelectionModify2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Modify2 calls the method "Selection.modify".
func (this Selection) Modify2(alter js.String) (ret js.Void) {
	bindings.CallSelectionModify2(
		this.ref, js.Pointer(&ret),
		alter.Ref(),
	)

	return
}

// TryModify2 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryModify2(alter js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		alter.Ref(),
	)

	return
}

// HasFuncModify3 returns true if the method "Selection.modify" exists.
func (this Selection) HasFuncModify3() bool {
	return js.True == bindings.HasFuncSelectionModify3(
		this.ref,
	)
}

// FuncModify3 returns the method "Selection.modify".
func (this Selection) FuncModify3() (fn js.Func[func()]) {
	bindings.FuncSelectionModify3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Modify3 calls the method "Selection.modify".
func (this Selection) Modify3() (ret js.Void) {
	bindings.CallSelectionModify3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryModify3 calls the method "Selection.modify"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryModify3() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionModify3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteFromDocument returns true if the method "Selection.deleteFromDocument" exists.
func (this Selection) HasFuncDeleteFromDocument() bool {
	return js.True == bindings.HasFuncSelectionDeleteFromDocument(
		this.ref,
	)
}

// FuncDeleteFromDocument returns the method "Selection.deleteFromDocument".
func (this Selection) FuncDeleteFromDocument() (fn js.Func[func()]) {
	bindings.FuncSelectionDeleteFromDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteFromDocument calls the method "Selection.deleteFromDocument".
func (this Selection) DeleteFromDocument() (ret js.Void) {
	bindings.CallSelectionDeleteFromDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeleteFromDocument calls the method "Selection.deleteFromDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryDeleteFromDocument() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionDeleteFromDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncContainsNode returns true if the method "Selection.containsNode" exists.
func (this Selection) HasFuncContainsNode() bool {
	return js.True == bindings.HasFuncSelectionContainsNode(
		this.ref,
	)
}

// FuncContainsNode returns the method "Selection.containsNode".
func (this Selection) FuncContainsNode() (fn js.Func[func(node Node, allowPartialContainment bool) bool]) {
	bindings.FuncSelectionContainsNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ContainsNode calls the method "Selection.containsNode".
func (this Selection) ContainsNode(node Node, allowPartialContainment bool) (ret bool) {
	bindings.CallSelectionContainsNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		js.Bool(bool(allowPartialContainment)),
	)

	return
}

// TryContainsNode calls the method "Selection.containsNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryContainsNode(node Node, allowPartialContainment bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionContainsNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		js.Bool(bool(allowPartialContainment)),
	)

	return
}

// HasFuncContainsNode1 returns true if the method "Selection.containsNode" exists.
func (this Selection) HasFuncContainsNode1() bool {
	return js.True == bindings.HasFuncSelectionContainsNode1(
		this.ref,
	)
}

// FuncContainsNode1 returns the method "Selection.containsNode".
func (this Selection) FuncContainsNode1() (fn js.Func[func(node Node) bool]) {
	bindings.FuncSelectionContainsNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ContainsNode1 calls the method "Selection.containsNode".
func (this Selection) ContainsNode1(node Node) (ret bool) {
	bindings.CallSelectionContainsNode1(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryContainsNode1 calls the method "Selection.containsNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryContainsNode1(node Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionContainsNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncToString returns true if the method "Selection.toString" exists.
func (this Selection) HasFuncToString() bool {
	return js.True == bindings.HasFuncSelectionToString(
		this.ref,
	)
}

// FuncToString returns the method "Selection.toString".
func (this Selection) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncSelectionToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "Selection.toString".
func (this Selection) ToString() (ret js.String) {
	bindings.CallSelectionToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "Selection.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Selection) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectionToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CaretPosition struct {
	ref js.Ref
}

func (this CaretPosition) Once() CaretPosition {
	this.ref.Once()
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
	this.ref.Free()
}

// OffsetNode returns the value of property "CaretPosition.offsetNode".
//
// It returns ok=false if there is no such property.
func (this CaretPosition) OffsetNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetCaretPositionOffsetNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Offset returns the value of property "CaretPosition.offset".
//
// It returns ok=false if there is no such property.
func (this CaretPosition) Offset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCaretPositionOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetClientRect returns true if the method "CaretPosition.getClientRect" exists.
func (this CaretPosition) HasFuncGetClientRect() bool {
	return js.True == bindings.HasFuncCaretPositionGetClientRect(
		this.ref,
	)
}

// FuncGetClientRect returns the method "CaretPosition.getClientRect".
func (this CaretPosition) FuncGetClientRect() (fn js.Func[func() DOMRect]) {
	bindings.FuncCaretPositionGetClientRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetClientRect calls the method "CaretPosition.getClientRect".
func (this CaretPosition) GetClientRect() (ret DOMRect) {
	bindings.CallCaretPositionGetClientRect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetClientRect calls the method "CaretPosition.getClientRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CaretPosition) TryGetClientRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaretPositionGetClientRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// ResultType returns the value of property "XPathResult.resultType".
//
// It returns ok=false if there is no such property.
func (this XPathResult) ResultType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetXPathResultResultType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberValue returns the value of property "XPathResult.numberValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) NumberValue() (ret float64, ok bool) {
	ok = js.True == bindings.GetXPathResultNumberValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StringValue returns the value of property "XPathResult.stringValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) StringValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXPathResultStringValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BooleanValue returns the value of property "XPathResult.booleanValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) BooleanValue() (ret bool, ok bool) {
	ok = js.True == bindings.GetXPathResultBooleanValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SingleNodeValue returns the value of property "XPathResult.singleNodeValue".
//
// It returns ok=false if there is no such property.
func (this XPathResult) SingleNodeValue() (ret Node, ok bool) {
	ok = js.True == bindings.GetXPathResultSingleNodeValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InvalidIteratorState returns the value of property "XPathResult.invalidIteratorState".
//
// It returns ok=false if there is no such property.
func (this XPathResult) InvalidIteratorState() (ret bool, ok bool) {
	ok = js.True == bindings.GetXPathResultInvalidIteratorState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SnapshotLength returns the value of property "XPathResult.snapshotLength".
//
// It returns ok=false if there is no such property.
func (this XPathResult) SnapshotLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXPathResultSnapshotLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncIterateNext returns true if the method "XPathResult.iterateNext" exists.
func (this XPathResult) HasFuncIterateNext() bool {
	return js.True == bindings.HasFuncXPathResultIterateNext(
		this.ref,
	)
}

// FuncIterateNext returns the method "XPathResult.iterateNext".
func (this XPathResult) FuncIterateNext() (fn js.Func[func() Node]) {
	bindings.FuncXPathResultIterateNext(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IterateNext calls the method "XPathResult.iterateNext".
func (this XPathResult) IterateNext() (ret Node) {
	bindings.CallXPathResultIterateNext(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIterateNext calls the method "XPathResult.iterateNext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathResult) TryIterateNext() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathResultIterateNext(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSnapshotItem returns true if the method "XPathResult.snapshotItem" exists.
func (this XPathResult) HasFuncSnapshotItem() bool {
	return js.True == bindings.HasFuncXPathResultSnapshotItem(
		this.ref,
	)
}

// FuncSnapshotItem returns the method "XPathResult.snapshotItem".
func (this XPathResult) FuncSnapshotItem() (fn js.Func[func(index uint32) Node]) {
	bindings.FuncXPathResultSnapshotItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SnapshotItem calls the method "XPathResult.snapshotItem".
func (this XPathResult) SnapshotItem(index uint32) (ret Node) {
	bindings.CallXPathResultSnapshotItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TrySnapshotItem calls the method "XPathResult.snapshotItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathResult) TrySnapshotItem(index uint32) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathResultSnapshotItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type XPathExpression struct {
	ref js.Ref
}

func (this XPathExpression) Once() XPathExpression {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncEvaluate returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasFuncEvaluate() bool {
	return js.True == bindings.HasFuncXPathExpressionEvaluate(
		this.ref,
	)
}

// FuncEvaluate returns the method "XPathExpression.evaluate".
func (this XPathExpression) FuncEvaluate() (fn js.Func[func(contextNode Node, typ uint16, result XPathResult) XPathResult]) {
	bindings.FuncXPathExpressionEvaluate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate(contextNode Node, typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate(
		this.ref, js.Pointer(&ret),
		contextNode.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// TryEvaluate calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathExpression) TryEvaluate(contextNode Node, typ uint16, result XPathResult) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		contextNode.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasFuncEvaluate1 returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasFuncEvaluate1() bool {
	return js.True == bindings.HasFuncXPathExpressionEvaluate1(
		this.ref,
	)
}

// FuncEvaluate1 returns the method "XPathExpression.evaluate".
func (this XPathExpression) FuncEvaluate1() (fn js.Func[func(contextNode Node, typ uint16) XPathResult]) {
	bindings.FuncXPathExpressionEvaluate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate1 calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate1(contextNode Node, typ uint16) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate1(
		this.ref, js.Pointer(&ret),
		contextNode.Ref(),
		uint32(typ),
	)

	return
}

// TryEvaluate1 calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathExpression) TryEvaluate1(contextNode Node, typ uint16) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		contextNode.Ref(),
		uint32(typ),
	)

	return
}

// HasFuncEvaluate2 returns true if the method "XPathExpression.evaluate" exists.
func (this XPathExpression) HasFuncEvaluate2() bool {
	return js.True == bindings.HasFuncXPathExpressionEvaluate2(
		this.ref,
	)
}

// FuncEvaluate2 returns the method "XPathExpression.evaluate".
func (this XPathExpression) FuncEvaluate2() (fn js.Func[func(contextNode Node) XPathResult]) {
	bindings.FuncXPathExpressionEvaluate2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate2 calls the method "XPathExpression.evaluate".
func (this XPathExpression) Evaluate2(contextNode Node) (ret XPathResult) {
	bindings.CallXPathExpressionEvaluate2(
		this.ref, js.Pointer(&ret),
		contextNode.Ref(),
	)

	return
}

// TryEvaluate2 calls the method "XPathExpression.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XPathExpression) TryEvaluate2(contextNode Node) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXPathExpressionEvaluate2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "DocumentType.name".
//
// It returns ok=false if there is no such property.
func (this DocumentType) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PublicId returns the value of property "DocumentType.publicId".
//
// It returns ok=false if there is no such property.
func (this DocumentType) PublicId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypePublicId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SystemId returns the value of property "DocumentType.systemId".
//
// It returns ok=false if there is no such property.
func (this DocumentType) SystemId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTypeSystemId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncBefore returns true if the method "DocumentType.before" exists.
func (this DocumentType) HasFuncBefore() bool {
	return js.True == bindings.HasFuncDocumentTypeBefore(
		this.ref,
	)
}

// FuncBefore returns the method "DocumentType.before".
func (this DocumentType) FuncBefore() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentTypeBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Before calls the method "DocumentType.before".
func (this DocumentType) Before(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeBefore(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryBefore calls the method "DocumentType.before"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentType) TryBefore(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncAfter returns true if the method "DocumentType.after" exists.
func (this DocumentType) HasFuncAfter() bool {
	return js.True == bindings.HasFuncDocumentTypeAfter(
		this.ref,
	)
}

// FuncAfter returns the method "DocumentType.after".
func (this DocumentType) FuncAfter() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentTypeAfter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// After calls the method "DocumentType.after".
func (this DocumentType) After(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeAfter(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAfter calls the method "DocumentType.after"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentType) TryAfter(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeAfter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncReplaceWith returns true if the method "DocumentType.replaceWith" exists.
func (this DocumentType) HasFuncReplaceWith() bool {
	return js.True == bindings.HasFuncDocumentTypeReplaceWith(
		this.ref,
	)
}

// FuncReplaceWith returns the method "DocumentType.replaceWith".
func (this DocumentType) FuncReplaceWith() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentTypeReplaceWith(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceWith calls the method "DocumentType.replaceWith".
func (this DocumentType) ReplaceWith(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentTypeReplaceWith(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceWith calls the method "DocumentType.replaceWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentType) TryReplaceWith(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeReplaceWith(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncRemove returns true if the method "DocumentType.remove" exists.
func (this DocumentType) HasFuncRemove() bool {
	return js.True == bindings.HasFuncDocumentTypeRemove(
		this.ref,
	)
}

// FuncRemove returns the method "DocumentType.remove".
func (this DocumentType) FuncRemove() (fn js.Func[func()]) {
	bindings.FuncDocumentTypeRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "DocumentType.remove".
func (this DocumentType) Remove() (ret js.Void) {
	bindings.CallDocumentTypeRemove(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "DocumentType.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DocumentType) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentTypeRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type XMLDocument struct {
	Document
}

func (this XMLDocument) Once() XMLDocument {
	this.ref.Once()
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
	this.ref.Free()
}

type DOMImplementation struct {
	ref js.Ref
}

func (this DOMImplementation) Once() DOMImplementation {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncCreateDocumentType returns true if the method "DOMImplementation.createDocumentType" exists.
func (this DOMImplementation) HasFuncCreateDocumentType() bool {
	return js.True == bindings.HasFuncDOMImplementationCreateDocumentType(
		this.ref,
	)
}

// FuncCreateDocumentType returns the method "DOMImplementation.createDocumentType".
func (this DOMImplementation) FuncCreateDocumentType() (fn js.Func[func(qualifiedName js.String, publicId js.String, systemId js.String) DocumentType]) {
	bindings.FuncDOMImplementationCreateDocumentType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDocumentType calls the method "DOMImplementation.createDocumentType".
func (this DOMImplementation) CreateDocumentType(qualifiedName js.String, publicId js.String, systemId js.String) (ret DocumentType) {
	bindings.CallDOMImplementationCreateDocumentType(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
		publicId.Ref(),
		systemId.Ref(),
	)

	return
}

// TryCreateDocumentType calls the method "DOMImplementation.createDocumentType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryCreateDocumentType(qualifiedName js.String, publicId js.String, systemId js.String) (ret DocumentType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocumentType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
		publicId.Ref(),
		systemId.Ref(),
	)

	return
}

// HasFuncCreateDocument returns true if the method "DOMImplementation.createDocument" exists.
func (this DOMImplementation) HasFuncCreateDocument() bool {
	return js.True == bindings.HasFuncDOMImplementationCreateDocument(
		this.ref,
	)
}

// FuncCreateDocument returns the method "DOMImplementation.createDocument".
func (this DOMImplementation) FuncCreateDocument() (fn js.Func[func(namespace js.String, qualifiedName js.String, doctype DocumentType) XMLDocument]) {
	bindings.FuncDOMImplementationCreateDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDocument calls the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocument(namespace js.String, qualifiedName js.String, doctype DocumentType) (ret XMLDocument) {
	bindings.CallDOMImplementationCreateDocument(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
		doctype.Ref(),
	)

	return
}

// TryCreateDocument calls the method "DOMImplementation.createDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryCreateDocument(namespace js.String, qualifiedName js.String, doctype DocumentType) (ret XMLDocument, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		doctype.Ref(),
	)

	return
}

// HasFuncCreateDocument1 returns true if the method "DOMImplementation.createDocument" exists.
func (this DOMImplementation) HasFuncCreateDocument1() bool {
	return js.True == bindings.HasFuncDOMImplementationCreateDocument1(
		this.ref,
	)
}

// FuncCreateDocument1 returns the method "DOMImplementation.createDocument".
func (this DOMImplementation) FuncCreateDocument1() (fn js.Func[func(namespace js.String, qualifiedName js.String) XMLDocument]) {
	bindings.FuncDOMImplementationCreateDocument1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDocument1 calls the method "DOMImplementation.createDocument".
func (this DOMImplementation) CreateDocument1(namespace js.String, qualifiedName js.String) (ret XMLDocument) {
	bindings.CallDOMImplementationCreateDocument1(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// TryCreateDocument1 calls the method "DOMImplementation.createDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryCreateDocument1(namespace js.String, qualifiedName js.String) (ret XMLDocument, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateDocument1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncCreateHTMLDocument returns true if the method "DOMImplementation.createHTMLDocument" exists.
func (this DOMImplementation) HasFuncCreateHTMLDocument() bool {
	return js.True == bindings.HasFuncDOMImplementationCreateHTMLDocument(
		this.ref,
	)
}

// FuncCreateHTMLDocument returns the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) FuncCreateHTMLDocument() (fn js.Func[func(title js.String) Document]) {
	bindings.FuncDOMImplementationCreateHTMLDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateHTMLDocument calls the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocument(title js.String) (ret Document) {
	bindings.CallDOMImplementationCreateHTMLDocument(
		this.ref, js.Pointer(&ret),
		title.Ref(),
	)

	return
}

// TryCreateHTMLDocument calls the method "DOMImplementation.createHTMLDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryCreateHTMLDocument(title js.String) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateHTMLDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		title.Ref(),
	)

	return
}

// HasFuncCreateHTMLDocument1 returns true if the method "DOMImplementation.createHTMLDocument" exists.
func (this DOMImplementation) HasFuncCreateHTMLDocument1() bool {
	return js.True == bindings.HasFuncDOMImplementationCreateHTMLDocument1(
		this.ref,
	)
}

// FuncCreateHTMLDocument1 returns the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) FuncCreateHTMLDocument1() (fn js.Func[func() Document]) {
	bindings.FuncDOMImplementationCreateHTMLDocument1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateHTMLDocument1 calls the method "DOMImplementation.createHTMLDocument".
func (this DOMImplementation) CreateHTMLDocument1() (ret Document) {
	bindings.CallDOMImplementationCreateHTMLDocument1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateHTMLDocument1 calls the method "DOMImplementation.createHTMLDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryCreateHTMLDocument1() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationCreateHTMLDocument1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHasFeature returns true if the method "DOMImplementation.hasFeature" exists.
func (this DOMImplementation) HasFuncHasFeature() bool {
	return js.True == bindings.HasFuncDOMImplementationHasFeature(
		this.ref,
	)
}

// FuncHasFeature returns the method "DOMImplementation.hasFeature".
func (this DOMImplementation) FuncHasFeature() (fn js.Func[func() bool]) {
	bindings.FuncDOMImplementationHasFeature(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasFeature calls the method "DOMImplementation.hasFeature".
func (this DOMImplementation) HasFeature() (ret bool) {
	bindings.CallDOMImplementationHasFeature(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasFeature calls the method "DOMImplementation.hasFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMImplementation) TryHasFeature() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMImplementationHasFeature(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FragmentDirective struct {
	ref js.Ref
}

func (this FragmentDirective) Once() FragmentDirective {
	this.ref.Once()
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
	this.ref.Free()
}

type PermissionsPolicy struct {
	ref js.Ref
}

func (this PermissionsPolicy) Once() PermissionsPolicy {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAllowsFeature returns true if the method "PermissionsPolicy.allowsFeature" exists.
func (this PermissionsPolicy) HasFuncAllowsFeature() bool {
	return js.True == bindings.HasFuncPermissionsPolicyAllowsFeature(
		this.ref,
	)
}

// FuncAllowsFeature returns the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) FuncAllowsFeature() (fn js.Func[func(feature js.String, origin js.String) bool]) {
	bindings.FuncPermissionsPolicyAllowsFeature(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AllowsFeature calls the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeature(feature js.String, origin js.String) (ret bool) {
	bindings.CallPermissionsPolicyAllowsFeature(
		this.ref, js.Pointer(&ret),
		feature.Ref(),
		origin.Ref(),
	)

	return
}

// TryAllowsFeature calls the method "PermissionsPolicy.allowsFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PermissionsPolicy) TryAllowsFeature(feature js.String, origin js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowsFeature(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
		origin.Ref(),
	)

	return
}

// HasFuncAllowsFeature1 returns true if the method "PermissionsPolicy.allowsFeature" exists.
func (this PermissionsPolicy) HasFuncAllowsFeature1() bool {
	return js.True == bindings.HasFuncPermissionsPolicyAllowsFeature1(
		this.ref,
	)
}

// FuncAllowsFeature1 returns the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) FuncAllowsFeature1() (fn js.Func[func(feature js.String) bool]) {
	bindings.FuncPermissionsPolicyAllowsFeature1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AllowsFeature1 calls the method "PermissionsPolicy.allowsFeature".
func (this PermissionsPolicy) AllowsFeature1(feature js.String) (ret bool) {
	bindings.CallPermissionsPolicyAllowsFeature1(
		this.ref, js.Pointer(&ret),
		feature.Ref(),
	)

	return
}

// TryAllowsFeature1 calls the method "PermissionsPolicy.allowsFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PermissionsPolicy) TryAllowsFeature1(feature js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowsFeature1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		feature.Ref(),
	)

	return
}

// HasFuncFeatures returns true if the method "PermissionsPolicy.features" exists.
func (this PermissionsPolicy) HasFuncFeatures() bool {
	return js.True == bindings.HasFuncPermissionsPolicyFeatures(
		this.ref,
	)
}

// FuncFeatures returns the method "PermissionsPolicy.features".
func (this PermissionsPolicy) FuncFeatures() (fn js.Func[func() js.Array[js.String]]) {
	bindings.FuncPermissionsPolicyFeatures(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Features calls the method "PermissionsPolicy.features".
func (this PermissionsPolicy) Features() (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyFeatures(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFeatures calls the method "PermissionsPolicy.features"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PermissionsPolicy) TryFeatures() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyFeatures(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAllowedFeatures returns true if the method "PermissionsPolicy.allowedFeatures" exists.
func (this PermissionsPolicy) HasFuncAllowedFeatures() bool {
	return js.True == bindings.HasFuncPermissionsPolicyAllowedFeatures(
		this.ref,
	)
}

// FuncAllowedFeatures returns the method "PermissionsPolicy.allowedFeatures".
func (this PermissionsPolicy) FuncAllowedFeatures() (fn js.Func[func() js.Array[js.String]]) {
	bindings.FuncPermissionsPolicyAllowedFeatures(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AllowedFeatures calls the method "PermissionsPolicy.allowedFeatures".
func (this PermissionsPolicy) AllowedFeatures() (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyAllowedFeatures(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllowedFeatures calls the method "PermissionsPolicy.allowedFeatures"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PermissionsPolicy) TryAllowedFeatures() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyAllowedFeatures(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAllowlistForFeature returns true if the method "PermissionsPolicy.getAllowlistForFeature" exists.
func (this PermissionsPolicy) HasFuncGetAllowlistForFeature() bool {
	return js.True == bindings.HasFuncPermissionsPolicyGetAllowlistForFeature(
		this.ref,
	)
}

// FuncGetAllowlistForFeature returns the method "PermissionsPolicy.getAllowlistForFeature".
func (this PermissionsPolicy) FuncGetAllowlistForFeature() (fn js.Func[func(feature js.String) js.Array[js.String]]) {
	bindings.FuncPermissionsPolicyGetAllowlistForFeature(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllowlistForFeature calls the method "PermissionsPolicy.getAllowlistForFeature".
func (this PermissionsPolicy) GetAllowlistForFeature(feature js.String) (ret js.Array[js.String]) {
	bindings.CallPermissionsPolicyGetAllowlistForFeature(
		this.ref, js.Pointer(&ret),
		feature.Ref(),
	)

	return
}

// TryGetAllowlistForFeature calls the method "PermissionsPolicy.getAllowlistForFeature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PermissionsPolicy) TryGetAllowlistForFeature(feature js.String) (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPermissionsPolicyGetAllowlistForFeature(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *DocumentTimelineOptions) UpdateFrom(ref js.Ref) {
	bindings.DocumentTimelineOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DocumentTimelineOptions) Update(ref js.Ref) {
	bindings.DocumentTimelineOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DocumentTimelineOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

type NamedFlowMap struct {
	ref js.Ref
}

func (this NamedFlowMap) Once() NamedFlowMap {
	this.ref.Once()
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
	this.ref.Free()
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
func (p *FocusOptions) UpdateFrom(ref js.Ref) {
	bindings.FocusOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FocusOptions) Update(ref js.Ref) {
	bindings.FocusOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FocusOptions) FreeMembers(recursive bool) {
}

type SVGAnimatedString struct {
	ref js.Ref
}

func (this SVGAnimatedString) Once() SVGAnimatedString {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedString.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedString) BaseVal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedStringBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedString.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedString) SetBaseVal(val js.String) bool {
	return js.True == bindings.SetSVGAnimatedStringBaseVal(
		this.ref,
		val.Ref(),
	)
}

// AnimVal returns the value of property "SVGAnimatedString.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedString) AnimVal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedStringAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSStyleDeclaration struct {
	ref js.Ref
}

func (this CSSStyleDeclaration) Once() CSSStyleDeclaration {
	this.ref.Once()
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
	this.ref.Free()
}

// CssText returns the value of property "CSSStyleDeclaration.cssText".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) CssText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationCssText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCssText sets the value of property "CSSStyleDeclaration.cssText" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleDeclaration) SetCssText(val js.String) bool {
	return js.True == bindings.SetCSSStyleDeclarationCssText(
		this.ref,
		val.Ref(),
	)
}

// Length returns the value of property "CSSStyleDeclaration.length".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ParentRule returns the value of property "CSSStyleDeclaration.parentRule".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) ParentRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationParentRule(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CssFloat returns the value of property "CSSStyleDeclaration.cssFloat".
//
// It returns ok=false if there is no such property.
func (this CSSStyleDeclaration) CssFloat() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSStyleDeclarationCssFloat(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCssFloat sets the value of property "CSSStyleDeclaration.cssFloat" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleDeclaration) SetCssFloat(val js.String) bool {
	return js.True == bindings.SetCSSStyleDeclarationCssFloat(
		this.ref,
		val.Ref(),
	)
}

// HasFuncItem returns true if the method "CSSStyleDeclaration.item" exists.
func (this CSSStyleDeclaration) HasFuncItem() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationItem(
		this.ref,
	)
}

// FuncItem returns the method "CSSStyleDeclaration.item".
func (this CSSStyleDeclaration) FuncItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncCSSStyleDeclarationItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "CSSStyleDeclaration.item".
func (this CSSStyleDeclaration) Item(index uint32) (ret js.String) {
	bindings.CallCSSStyleDeclarationItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "CSSStyleDeclaration.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetPropertyValue returns true if the method "CSSStyleDeclaration.getPropertyValue" exists.
func (this CSSStyleDeclaration) HasFuncGetPropertyValue() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationGetPropertyValue(
		this.ref,
	)
}

// FuncGetPropertyValue returns the method "CSSStyleDeclaration.getPropertyValue".
func (this CSSStyleDeclaration) FuncGetPropertyValue() (fn js.Func[func(property js.String) js.String]) {
	bindings.FuncCSSStyleDeclarationGetPropertyValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPropertyValue calls the method "CSSStyleDeclaration.getPropertyValue".
func (this CSSStyleDeclaration) GetPropertyValue(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationGetPropertyValue(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetPropertyValue calls the method "CSSStyleDeclaration.getPropertyValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TryGetPropertyValue(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationGetPropertyValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasFuncGetPropertyPriority returns true if the method "CSSStyleDeclaration.getPropertyPriority" exists.
func (this CSSStyleDeclaration) HasFuncGetPropertyPriority() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationGetPropertyPriority(
		this.ref,
	)
}

// FuncGetPropertyPriority returns the method "CSSStyleDeclaration.getPropertyPriority".
func (this CSSStyleDeclaration) FuncGetPropertyPriority() (fn js.Func[func(property js.String) js.String]) {
	bindings.FuncCSSStyleDeclarationGetPropertyPriority(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPropertyPriority calls the method "CSSStyleDeclaration.getPropertyPriority".
func (this CSSStyleDeclaration) GetPropertyPriority(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationGetPropertyPriority(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetPropertyPriority calls the method "CSSStyleDeclaration.getPropertyPriority"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TryGetPropertyPriority(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationGetPropertyPriority(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasFuncSetProperty returns true if the method "CSSStyleDeclaration.setProperty" exists.
func (this CSSStyleDeclaration) HasFuncSetProperty() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationSetProperty(
		this.ref,
	)
}

// FuncSetProperty returns the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) FuncSetProperty() (fn js.Func[func(property js.String, value js.String, priority js.String)]) {
	bindings.FuncCSSStyleDeclarationSetProperty(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetProperty calls the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetProperty(property js.String, value js.String, priority js.String) (ret js.Void) {
	bindings.CallCSSStyleDeclarationSetProperty(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		value.Ref(),
		priority.Ref(),
	)

	return
}

// TrySetProperty calls the method "CSSStyleDeclaration.setProperty"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TrySetProperty(property js.String, value js.String, priority js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationSetProperty(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
		priority.Ref(),
	)

	return
}

// HasFuncSetProperty1 returns true if the method "CSSStyleDeclaration.setProperty" exists.
func (this CSSStyleDeclaration) HasFuncSetProperty1() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationSetProperty1(
		this.ref,
	)
}

// FuncSetProperty1 returns the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) FuncSetProperty1() (fn js.Func[func(property js.String, value js.String)]) {
	bindings.FuncCSSStyleDeclarationSetProperty1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetProperty1 calls the method "CSSStyleDeclaration.setProperty".
func (this CSSStyleDeclaration) SetProperty1(property js.String, value js.String) (ret js.Void) {
	bindings.CallCSSStyleDeclarationSetProperty1(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		value.Ref(),
	)

	return
}

// TrySetProperty1 calls the method "CSSStyleDeclaration.setProperty"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TrySetProperty1(property js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationSetProperty1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncRemoveProperty returns true if the method "CSSStyleDeclaration.removeProperty" exists.
func (this CSSStyleDeclaration) HasFuncRemoveProperty() bool {
	return js.True == bindings.HasFuncCSSStyleDeclarationRemoveProperty(
		this.ref,
	)
}

// FuncRemoveProperty returns the method "CSSStyleDeclaration.removeProperty".
func (this CSSStyleDeclaration) FuncRemoveProperty() (fn js.Func[func(property js.String) js.String]) {
	bindings.FuncCSSStyleDeclarationRemoveProperty(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveProperty calls the method "CSSStyleDeclaration.removeProperty".
func (this CSSStyleDeclaration) RemoveProperty(property js.String) (ret js.String) {
	bindings.CallCSSStyleDeclarationRemoveProperty(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryRemoveProperty calls the method "CSSStyleDeclaration.removeProperty"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleDeclaration) TryRemoveProperty(property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleDeclarationRemoveProperty(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSet returns true if the method "StylePropertyMap.set" exists.
func (this StylePropertyMap) HasFuncSet() bool {
	return js.True == bindings.HasFuncStylePropertyMapSet(
		this.ref,
	)
}

// FuncSet returns the method "StylePropertyMap.set".
func (this StylePropertyMap) FuncSet() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	bindings.FuncStylePropertyMapSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "StylePropertyMap.set".
func (this StylePropertyMap) Set(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void) {
	bindings.CallStylePropertyMapSet(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TrySet calls the method "StylePropertyMap.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMap) TrySet(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncAppend returns true if the method "StylePropertyMap.append" exists.
func (this StylePropertyMap) HasFuncAppend() bool {
	return js.True == bindings.HasFuncStylePropertyMapAppend(
		this.ref,
	)
}

// FuncAppend returns the method "StylePropertyMap.append".
func (this StylePropertyMap) FuncAppend() (fn js.Func[func(property js.String, values ...OneOf_CSSStyleValue_String)]) {
	bindings.FuncStylePropertyMapAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "StylePropertyMap.append".
func (this StylePropertyMap) Append(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void) {
	bindings.CallStylePropertyMapAppend(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryAppend calls the method "StylePropertyMap.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMap) TryAppend(property js.String, values ...OneOf_CSSStyleValue_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncDelete returns true if the method "StylePropertyMap.delete" exists.
func (this StylePropertyMap) HasFuncDelete() bool {
	return js.True == bindings.HasFuncStylePropertyMapDelete(
		this.ref,
	)
}

// FuncDelete returns the method "StylePropertyMap.delete".
func (this StylePropertyMap) FuncDelete() (fn js.Func[func(property js.String)]) {
	bindings.FuncStylePropertyMapDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "StylePropertyMap.delete".
func (this StylePropertyMap) Delete(property js.String) (ret js.Void) {
	bindings.CallStylePropertyMapDelete(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryDelete calls the method "StylePropertyMap.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMap) TryDelete(property js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "StylePropertyMap.clear" exists.
func (this StylePropertyMap) HasFuncClear() bool {
	return js.True == bindings.HasFuncStylePropertyMapClear(
		this.ref,
	)
}

// FuncClear returns the method "StylePropertyMap.clear".
func (this StylePropertyMap) FuncClear() (fn js.Func[func()]) {
	bindings.FuncStylePropertyMapClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "StylePropertyMap.clear".
func (this StylePropertyMap) Clear() (ret js.Void) {
	bindings.CallStylePropertyMapClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "StylePropertyMap.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMap) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// UnitType returns the value of property "SVGLength.unitType".
//
// It returns ok=false if there is no such property.
func (this SVGLength) UnitType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGLengthUnitType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "SVGLength.value".
//
// It returns ok=false if there is no such property.
func (this SVGLength) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGLengthValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGLength.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValue(val float32) bool {
	return js.True == bindings.SetSVGLengthValue(
		this.ref,
		float32(val),
	)
}

// ValueInSpecifiedUnits returns the value of property "SVGLength.valueInSpecifiedUnits".
//
// It returns ok=false if there is no such property.
func (this SVGLength) ValueInSpecifiedUnits() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGLengthValueInSpecifiedUnits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueInSpecifiedUnits sets the value of property "SVGLength.valueInSpecifiedUnits" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValueInSpecifiedUnits(val float32) bool {
	return js.True == bindings.SetSVGLengthValueInSpecifiedUnits(
		this.ref,
		float32(val),
	)
}

// ValueAsString returns the value of property "SVGLength.valueAsString".
//
// It returns ok=false if there is no such property.
func (this SVGLength) ValueAsString() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGLengthValueAsString(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueAsString sets the value of property "SVGLength.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGLength) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGLengthValueAsString(
		this.ref,
		val.Ref(),
	)
}

// HasFuncNewValueSpecifiedUnits returns true if the method "SVGLength.newValueSpecifiedUnits" exists.
func (this SVGLength) HasFuncNewValueSpecifiedUnits() bool {
	return js.True == bindings.HasFuncSVGLengthNewValueSpecifiedUnits(
		this.ref,
	)
}

// FuncNewValueSpecifiedUnits returns the method "SVGLength.newValueSpecifiedUnits".
func (this SVGLength) FuncNewValueSpecifiedUnits() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	bindings.FuncSVGLengthNewValueSpecifiedUnits(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NewValueSpecifiedUnits calls the method "SVGLength.newValueSpecifiedUnits".
func (this SVGLength) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void) {
	bindings.CallSVGLengthNewValueSpecifiedUnits(
		this.ref, js.Pointer(&ret),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// TryNewValueSpecifiedUnits calls the method "SVGLength.newValueSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLength) TryNewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthNewValueSpecifiedUnits(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// HasFuncConvertToSpecifiedUnits returns true if the method "SVGLength.convertToSpecifiedUnits" exists.
func (this SVGLength) HasFuncConvertToSpecifiedUnits() bool {
	return js.True == bindings.HasFuncSVGLengthConvertToSpecifiedUnits(
		this.ref,
	)
}

// FuncConvertToSpecifiedUnits returns the method "SVGLength.convertToSpecifiedUnits".
func (this SVGLength) FuncConvertToSpecifiedUnits() (fn js.Func[func(unitType uint16)]) {
	bindings.FuncSVGLengthConvertToSpecifiedUnits(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertToSpecifiedUnits calls the method "SVGLength.convertToSpecifiedUnits".
func (this SVGLength) ConvertToSpecifiedUnits(unitType uint16) (ret js.Void) {
	bindings.CallSVGLengthConvertToSpecifiedUnits(
		this.ref, js.Pointer(&ret),
		uint32(unitType),
	)

	return
}

// TryConvertToSpecifiedUnits calls the method "SVGLength.convertToSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLength) TryConvertToSpecifiedUnits(unitType uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthConvertToSpecifiedUnits(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
	)

	return
}

type SVGAnimatedLength struct {
	ref js.Ref
}

func (this SVGAnimatedLength) Once() SVGAnimatedLength {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedLength.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLength) BaseVal() (ret SVGLength, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedLength.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLength) AnimVal() (ret SVGLength, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGUseElement struct {
	SVGGraphicsElement
}

func (this SVGUseElement) Once() SVGUseElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGUseElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGUseElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGUseElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGUseElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGUseElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InstanceRoot returns the value of property "SVGUseElement.instanceRoot".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) InstanceRoot() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGUseElementInstanceRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimatedInstanceRoot returns the value of property "SVGUseElement.animatedInstanceRoot".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) AnimatedInstanceRoot() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGUseElementAnimatedInstanceRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGUseElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGUseElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGUseElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

type DOMStringMap struct {
	ref js.Ref
}

func (this DOMStringMap) Once() DOMStringMap {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGet returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasFuncGet() bool {
	return js.True == bindings.HasFuncDOMStringMapGet(
		this.ref,
	)
}

// FuncGet returns the method "DOMStringMap.".
func (this DOMStringMap) FuncGet() (fn js.Func[func(name js.String) js.String]) {
	bindings.FuncDOMStringMapGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "DOMStringMap.".
func (this DOMStringMap) Get(name js.String) (ret js.String) {
	bindings.CallDOMStringMapGet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringMap) TryGet(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasFuncSet() bool {
	return js.True == bindings.HasFuncDOMStringMapSet(
		this.ref,
	)
}

// FuncSet returns the method "DOMStringMap.".
func (this DOMStringMap) FuncSet() (fn js.Func[func(name js.String, value js.String)]) {
	bindings.FuncDOMStringMapSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "DOMStringMap.".
func (this DOMStringMap) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallDOMStringMapSet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringMap) TrySet(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "DOMStringMap." exists.
func (this DOMStringMap) HasFuncDelete() bool {
	return js.True == bindings.HasFuncDOMStringMapDelete(
		this.ref,
	)
}

// FuncDelete returns the method "DOMStringMap.".
func (this DOMStringMap) FuncDelete() (fn js.Func[func(name js.String)]) {
	bindings.FuncDOMStringMapDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "DOMStringMap.".
func (this DOMStringMap) Delete(name js.String) (ret js.Void) {
	bindings.CallDOMStringMapDelete(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "DOMStringMap."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringMap) TryDelete(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringMapDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type SVGElement struct {
	Element
}

func (this SVGElement) Once() SVGElement {
	this.ref.Once()
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
	this.ref.Free()
}

// ClassName returns the value of property "SVGElement.className".
//
// It returns ok=false if there is no such property.
func (this SVGElement) ClassName() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGElementClassName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OwnerSVGElement returns the value of property "SVGElement.ownerSVGElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) OwnerSVGElement() (ret SVGSVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementOwnerSVGElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ViewportElement returns the value of property "SVGElement.viewportElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) ViewportElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementViewportElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "SVGElement.style".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetSVGElementStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "SVGElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this SVGElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetSVGElementAttributeStyleMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CorrespondingElement returns the value of property "SVGElement.correspondingElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) CorrespondingElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGElementCorrespondingElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CorrespondingUseElement returns the value of property "SVGElement.correspondingUseElement".
//
// It returns ok=false if there is no such property.
func (this SVGElement) CorrespondingUseElement() (ret SVGUseElement, ok bool) {
	ok = js.True == bindings.GetSVGElementCorrespondingUseElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "SVGElement.dataset".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetSVGElementDataset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "SVGElement.nonce".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGElementNonce(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "SVGElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetSVGElementNonce(
		this.ref,
		val.Ref(),
	)
}

// Autofocus returns the value of property "SVGElement.autofocus".
//
// It returns ok=false if there is no such property.
func (this SVGElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetSVGElementAutofocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "SVGElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetSVGElementAutofocus(
		this.ref,
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "SVGElement.tabIndex".
//
// It returns ok=false if there is no such property.
func (this SVGElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGElementTabIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "SVGElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this SVGElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetSVGElementTabIndex(
		this.ref,
		int32(val),
	)
}

// HasFuncFocus returns true if the method "SVGElement.focus" exists.
func (this SVGElement) HasFuncFocus() bool {
	return js.True == bindings.HasFuncSVGElementFocus(
		this.ref,
	)
}

// FuncFocus returns the method "SVGElement.focus".
func (this SVGElement) FuncFocus() (fn js.Func[func(options FocusOptions)]) {
	bindings.FuncSVGElementFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "SVGElement.focus".
func (this SVGElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallSVGElementFocus(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "SVGElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncFocus1 returns true if the method "SVGElement.focus" exists.
func (this SVGElement) HasFuncFocus1() bool {
	return js.True == bindings.HasFuncSVGElementFocus1(
		this.ref,
	)
}

// FuncFocus1 returns the method "SVGElement.focus".
func (this SVGElement) FuncFocus1() (fn js.Func[func()]) {
	bindings.FuncSVGElementFocus1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus1 calls the method "SVGElement.focus".
func (this SVGElement) Focus1() (ret js.Void) {
	bindings.CallSVGElementFocus1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "SVGElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementFocus1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlur returns true if the method "SVGElement.blur" exists.
func (this SVGElement) HasFuncBlur() bool {
	return js.True == bindings.HasFuncSVGElementBlur(
		this.ref,
	)
}

// FuncBlur returns the method "SVGElement.blur".
func (this SVGElement) FuncBlur() (fn js.Func[func()]) {
	bindings.FuncSVGElementBlur(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blur calls the method "SVGElement.blur".
func (this SVGElement) Blur() (ret js.Void) {
	bindings.CallSVGElementBlur(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "SVGElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGElementBlur(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGNumber struct {
	ref js.Ref
}

func (this SVGNumber) Once() SVGNumber {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "SVGNumber.value".
//
// It returns ok=false if there is no such property.
func (this SVGNumber) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGNumberValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGNumber.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGNumber) SetValue(val float32) bool {
	return js.True == bindings.SetSVGNumberValue(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// UnitType returns the value of property "SVGAngle.unitType".
//
// It returns ok=false if there is no such property.
func (this SVGAngle) UnitType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAngleUnitType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "SVGAngle.value".
//
// It returns ok=false if there is no such property.
func (this SVGAngle) Value() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAngleValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "SVGAngle.value" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValue(val float32) bool {
	return js.True == bindings.SetSVGAngleValue(
		this.ref,
		float32(val),
	)
}

// ValueInSpecifiedUnits returns the value of property "SVGAngle.valueInSpecifiedUnits".
//
// It returns ok=false if there is no such property.
func (this SVGAngle) ValueInSpecifiedUnits() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAngleValueInSpecifiedUnits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueInSpecifiedUnits sets the value of property "SVGAngle.valueInSpecifiedUnits" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValueInSpecifiedUnits(val float32) bool {
	return js.True == bindings.SetSVGAngleValueInSpecifiedUnits(
		this.ref,
		float32(val),
	)
}

// ValueAsString returns the value of property "SVGAngle.valueAsString".
//
// It returns ok=false if there is no such property.
func (this SVGAngle) ValueAsString() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGAngleValueAsString(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueAsString sets the value of property "SVGAngle.valueAsString" to val.
//
// It returns false if the property cannot be set.
func (this SVGAngle) SetValueAsString(val js.String) bool {
	return js.True == bindings.SetSVGAngleValueAsString(
		this.ref,
		val.Ref(),
	)
}

// HasFuncNewValueSpecifiedUnits returns true if the method "SVGAngle.newValueSpecifiedUnits" exists.
func (this SVGAngle) HasFuncNewValueSpecifiedUnits() bool {
	return js.True == bindings.HasFuncSVGAngleNewValueSpecifiedUnits(
		this.ref,
	)
}

// FuncNewValueSpecifiedUnits returns the method "SVGAngle.newValueSpecifiedUnits".
func (this SVGAngle) FuncNewValueSpecifiedUnits() (fn js.Func[func(unitType uint16, valueInSpecifiedUnits float32)]) {
	bindings.FuncSVGAngleNewValueSpecifiedUnits(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NewValueSpecifiedUnits calls the method "SVGAngle.newValueSpecifiedUnits".
func (this SVGAngle) NewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void) {
	bindings.CallSVGAngleNewValueSpecifiedUnits(
		this.ref, js.Pointer(&ret),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// TryNewValueSpecifiedUnits calls the method "SVGAngle.newValueSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAngle) TryNewValueSpecifiedUnits(unitType uint16, valueInSpecifiedUnits float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAngleNewValueSpecifiedUnits(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(unitType),
		float32(valueInSpecifiedUnits),
	)

	return
}

// HasFuncConvertToSpecifiedUnits returns true if the method "SVGAngle.convertToSpecifiedUnits" exists.
func (this SVGAngle) HasFuncConvertToSpecifiedUnits() bool {
	return js.True == bindings.HasFuncSVGAngleConvertToSpecifiedUnits(
		this.ref,
	)
}

// FuncConvertToSpecifiedUnits returns the method "SVGAngle.convertToSpecifiedUnits".
func (this SVGAngle) FuncConvertToSpecifiedUnits() (fn js.Func[func(unitType uint16)]) {
	bindings.FuncSVGAngleConvertToSpecifiedUnits(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertToSpecifiedUnits calls the method "SVGAngle.convertToSpecifiedUnits".
func (this SVGAngle) ConvertToSpecifiedUnits(unitType uint16) (ret js.Void) {
	bindings.CallSVGAngleConvertToSpecifiedUnits(
		this.ref, js.Pointer(&ret),
		uint32(unitType),
	)

	return
}

// TryConvertToSpecifiedUnits calls the method "SVGAngle.convertToSpecifiedUnits"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAngle) TryConvertToSpecifiedUnits(unitType uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAngleConvertToSpecifiedUnits(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *DOMMatrixInit) UpdateFrom(ref js.Ref) {
	bindings.DOMMatrixInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMMatrixInit) Update(ref js.Ref) {
	bindings.DOMMatrixInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMMatrixInit) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// A returns the value of property "DOMMatrix.a".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) A() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixA(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "DOMMatrix.a" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetA(val float64) bool {
	return js.True == bindings.SetDOMMatrixA(
		this.ref,
		float64(val),
	)
}

// B returns the value of property "DOMMatrix.b".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) B() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "DOMMatrix.b" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetB(val float64) bool {
	return js.True == bindings.SetDOMMatrixB(
		this.ref,
		float64(val),
	)
}

// C returns the value of property "DOMMatrix.c".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) C() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixC(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "DOMMatrix.c" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetC(val float64) bool {
	return js.True == bindings.SetDOMMatrixC(
		this.ref,
		float64(val),
	)
}

// D returns the value of property "DOMMatrix.d".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) D() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixD(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetD sets the value of property "DOMMatrix.d" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetD(val float64) bool {
	return js.True == bindings.SetDOMMatrixD(
		this.ref,
		float64(val),
	)
}

// E returns the value of property "DOMMatrix.e".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) E() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixE(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetE sets the value of property "DOMMatrix.e" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetE(val float64) bool {
	return js.True == bindings.SetDOMMatrixE(
		this.ref,
		float64(val),
	)
}

// F returns the value of property "DOMMatrix.f".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) F() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixF(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetF sets the value of property "DOMMatrix.f" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetF(val float64) bool {
	return js.True == bindings.SetDOMMatrixF(
		this.ref,
		float64(val),
	)
}

// M11 returns the value of property "DOMMatrix.m11".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M11() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM11(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM11 sets the value of property "DOMMatrix.m11" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM11(val float64) bool {
	return js.True == bindings.SetDOMMatrixM11(
		this.ref,
		float64(val),
	)
}

// M12 returns the value of property "DOMMatrix.m12".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M12() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM12(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM12 sets the value of property "DOMMatrix.m12" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM12(val float64) bool {
	return js.True == bindings.SetDOMMatrixM12(
		this.ref,
		float64(val),
	)
}

// M13 returns the value of property "DOMMatrix.m13".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M13() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM13(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM13 sets the value of property "DOMMatrix.m13" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM13(val float64) bool {
	return js.True == bindings.SetDOMMatrixM13(
		this.ref,
		float64(val),
	)
}

// M14 returns the value of property "DOMMatrix.m14".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M14() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM14(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM14 sets the value of property "DOMMatrix.m14" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM14(val float64) bool {
	return js.True == bindings.SetDOMMatrixM14(
		this.ref,
		float64(val),
	)
}

// M21 returns the value of property "DOMMatrix.m21".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M21() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM21(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM21 sets the value of property "DOMMatrix.m21" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM21(val float64) bool {
	return js.True == bindings.SetDOMMatrixM21(
		this.ref,
		float64(val),
	)
}

// M22 returns the value of property "DOMMatrix.m22".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M22() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM22(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM22 sets the value of property "DOMMatrix.m22" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM22(val float64) bool {
	return js.True == bindings.SetDOMMatrixM22(
		this.ref,
		float64(val),
	)
}

// M23 returns the value of property "DOMMatrix.m23".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M23() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM23(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM23 sets the value of property "DOMMatrix.m23" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM23(val float64) bool {
	return js.True == bindings.SetDOMMatrixM23(
		this.ref,
		float64(val),
	)
}

// M24 returns the value of property "DOMMatrix.m24".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M24() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM24(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM24 sets the value of property "DOMMatrix.m24" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM24(val float64) bool {
	return js.True == bindings.SetDOMMatrixM24(
		this.ref,
		float64(val),
	)
}

// M31 returns the value of property "DOMMatrix.m31".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M31() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM31(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM31 sets the value of property "DOMMatrix.m31" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM31(val float64) bool {
	return js.True == bindings.SetDOMMatrixM31(
		this.ref,
		float64(val),
	)
}

// M32 returns the value of property "DOMMatrix.m32".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M32() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM32(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM32 sets the value of property "DOMMatrix.m32" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM32(val float64) bool {
	return js.True == bindings.SetDOMMatrixM32(
		this.ref,
		float64(val),
	)
}

// M33 returns the value of property "DOMMatrix.m33".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M33() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM33(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM33 sets the value of property "DOMMatrix.m33" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM33(val float64) bool {
	return js.True == bindings.SetDOMMatrixM33(
		this.ref,
		float64(val),
	)
}

// M34 returns the value of property "DOMMatrix.m34".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M34() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM34(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM34 sets the value of property "DOMMatrix.m34" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM34(val float64) bool {
	return js.True == bindings.SetDOMMatrixM34(
		this.ref,
		float64(val),
	)
}

// M41 returns the value of property "DOMMatrix.m41".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M41() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM41(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM41 sets the value of property "DOMMatrix.m41" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM41(val float64) bool {
	return js.True == bindings.SetDOMMatrixM41(
		this.ref,
		float64(val),
	)
}

// M42 returns the value of property "DOMMatrix.m42".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M42() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM42(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM42 sets the value of property "DOMMatrix.m42" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM42(val float64) bool {
	return js.True == bindings.SetDOMMatrixM42(
		this.ref,
		float64(val),
	)
}

// M43 returns the value of property "DOMMatrix.m43".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M43() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM43(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM43 sets the value of property "DOMMatrix.m43" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM43(val float64) bool {
	return js.True == bindings.SetDOMMatrixM43(
		this.ref,
		float64(val),
	)
}

// M44 returns the value of property "DOMMatrix.m44".
//
// It returns ok=false if there is no such property.
func (this DOMMatrix) M44() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixM44(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetM44 sets the value of property "DOMMatrix.m44" to val.
//
// It returns false if the property cannot be set.
func (this DOMMatrix) SetM44(val float64) bool {
	return js.True == bindings.SetDOMMatrixM44(
		this.ref,
		float64(val),
	)
}

// HasFuncFromMatrix returns true if the static method "DOMMatrix.fromMatrix" exists.
func (this DOMMatrix) HasFuncFromMatrix() bool {
	return js.True == bindings.HasFuncDOMMatrixFromMatrix(
		this.ref,
	)
}

// FuncFromMatrix returns the static method "DOMMatrix.fromMatrix".
func (this DOMMatrix) FuncFromMatrix() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	bindings.FuncDOMMatrixFromMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromMatrix calls the static method "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrix(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromMatrix(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromMatrix calls the static method "DOMMatrix.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryFromMatrix(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromMatrix1 returns true if the static method "DOMMatrix.fromMatrix" exists.
func (this DOMMatrix) HasFuncFromMatrix1() bool {
	return js.True == bindings.HasFuncDOMMatrixFromMatrix1(
		this.ref,
	)
}

// FuncFromMatrix1 returns the static method "DOMMatrix.fromMatrix".
func (this DOMMatrix) FuncFromMatrix1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixFromMatrix1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromMatrix1 calls the static method "DOMMatrix.fromMatrix".
func (this DOMMatrix) FromMatrix1() (ret DOMMatrix) {
	bindings.CallDOMMatrixFromMatrix1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromMatrix1 calls the static method "DOMMatrix.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryFromMatrix1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromMatrix1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFromFloat32Array returns true if the static method "DOMMatrix.fromFloat32Array" exists.
func (this DOMMatrix) HasFuncFromFloat32Array() bool {
	return js.True == bindings.HasFuncDOMMatrixFromFloat32Array(
		this.ref,
	)
}

// FuncFromFloat32Array returns the static method "DOMMatrix.fromFloat32Array".
func (this DOMMatrix) FuncFromFloat32Array() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrix]) {
	bindings.FuncDOMMatrixFromFloat32Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromFloat32Array calls the static method "DOMMatrix.fromFloat32Array".
func (this DOMMatrix) FromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromFloat32Array(
		this.ref, js.Pointer(&ret),
		array32.Ref(),
	)

	return
}

// TryFromFloat32Array calls the static method "DOMMatrix.fromFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryFromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromFloat32Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array32.Ref(),
	)

	return
}

// HasFuncFromFloat64Array returns true if the static method "DOMMatrix.fromFloat64Array" exists.
func (this DOMMatrix) HasFuncFromFloat64Array() bool {
	return js.True == bindings.HasFuncDOMMatrixFromFloat64Array(
		this.ref,
	)
}

// FuncFromFloat64Array returns the static method "DOMMatrix.fromFloat64Array".
func (this DOMMatrix) FuncFromFloat64Array() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrix]) {
	bindings.FuncDOMMatrixFromFloat64Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromFloat64Array calls the static method "DOMMatrix.fromFloat64Array".
func (this DOMMatrix) FromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrix) {
	bindings.CallDOMMatrixFromFloat64Array(
		this.ref, js.Pointer(&ret),
		array64.Ref(),
	)

	return
}

// TryFromFloat64Array calls the static method "DOMMatrix.fromFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryFromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixFromFloat64Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array64.Ref(),
	)

	return
}

// HasFuncMultiplySelf returns true if the method "DOMMatrix.multiplySelf" exists.
func (this DOMMatrix) HasFuncMultiplySelf() bool {
	return js.True == bindings.HasFuncDOMMatrixMultiplySelf(
		this.ref,
	)
}

// FuncMultiplySelf returns the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) FuncMultiplySelf() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	bindings.FuncDOMMatrixMultiplySelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MultiplySelf calls the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelf(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixMultiplySelf(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryMultiplySelf calls the method "DOMMatrix.multiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryMultiplySelf(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixMultiplySelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncMultiplySelf1 returns true if the method "DOMMatrix.multiplySelf" exists.
func (this DOMMatrix) HasFuncMultiplySelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixMultiplySelf1(
		this.ref,
	)
}

// FuncMultiplySelf1 returns the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) FuncMultiplySelf1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixMultiplySelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MultiplySelf1 calls the method "DOMMatrix.multiplySelf".
func (this DOMMatrix) MultiplySelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixMultiplySelf1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMultiplySelf1 calls the method "DOMMatrix.multiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryMultiplySelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixMultiplySelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreMultiplySelf returns true if the method "DOMMatrix.preMultiplySelf" exists.
func (this DOMMatrix) HasFuncPreMultiplySelf() bool {
	return js.True == bindings.HasFuncDOMMatrixPreMultiplySelf(
		this.ref,
	)
}

// FuncPreMultiplySelf returns the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) FuncPreMultiplySelf() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	bindings.FuncDOMMatrixPreMultiplySelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreMultiplySelf calls the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelf(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixPreMultiplySelf(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryPreMultiplySelf calls the method "DOMMatrix.preMultiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryPreMultiplySelf(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixPreMultiplySelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncPreMultiplySelf1 returns true if the method "DOMMatrix.preMultiplySelf" exists.
func (this DOMMatrix) HasFuncPreMultiplySelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixPreMultiplySelf1(
		this.ref,
	)
}

// FuncPreMultiplySelf1 returns the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) FuncPreMultiplySelf1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixPreMultiplySelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreMultiplySelf1 calls the method "DOMMatrix.preMultiplySelf".
func (this DOMMatrix) PreMultiplySelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixPreMultiplySelf1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreMultiplySelf1 calls the method "DOMMatrix.preMultiplySelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryPreMultiplySelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixPreMultiplySelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTranslateSelf returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasFuncTranslateSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixTranslateSelf(
		this.ref,
	)
}

// FuncTranslateSelf returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) FuncTranslateSelf() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	bindings.FuncDOMMatrixTranslateSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TranslateSelf calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf(tx float64, ty float64, tz float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf(
		this.ref, js.Pointer(&ret),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// TryTranslateSelf calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryTranslateSelf(tx float64, ty float64, tz float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// HasFuncTranslateSelf1 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasFuncTranslateSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixTranslateSelf1(
		this.ref,
	)
}

// FuncTranslateSelf1 returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) FuncTranslateSelf1() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	bindings.FuncDOMMatrixTranslateSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TranslateSelf1 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf1(tx float64, ty float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf1(
		this.ref, js.Pointer(&ret),
		float64(tx),
		float64(ty),
	)

	return
}

// TryTranslateSelf1 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryTranslateSelf1(tx float64, ty float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
	)

	return
}

// HasFuncTranslateSelf2 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasFuncTranslateSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixTranslateSelf2(
		this.ref,
	)
}

// FuncTranslateSelf2 returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) FuncTranslateSelf2() (fn js.Func[func(tx float64) DOMMatrix]) {
	bindings.FuncDOMMatrixTranslateSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TranslateSelf2 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf2(tx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf2(
		this.ref, js.Pointer(&ret),
		float64(tx),
	)

	return
}

// TryTranslateSelf2 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryTranslateSelf2(tx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
	)

	return
}

// HasFuncTranslateSelf3 returns true if the method "DOMMatrix.translateSelf" exists.
func (this DOMMatrix) HasFuncTranslateSelf3() bool {
	return js.True == bindings.HasFuncDOMMatrixTranslateSelf3(
		this.ref,
	)
}

// FuncTranslateSelf3 returns the method "DOMMatrix.translateSelf".
func (this DOMMatrix) FuncTranslateSelf3() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixTranslateSelf3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TranslateSelf3 calls the method "DOMMatrix.translateSelf".
func (this DOMMatrix) TranslateSelf3() (ret DOMMatrix) {
	bindings.CallDOMMatrixTranslateSelf3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTranslateSelf3 calls the method "DOMMatrix.translateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryTranslateSelf3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixTranslateSelf3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScaleSelf returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf(
		this.ref,
	)
}

// FuncScaleSelf returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasFuncScaleSelf1 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf1(
		this.ref,
	)
}

// FuncScaleSelf1 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf1() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf1 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf1(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScaleSelf1 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// HasFuncScaleSelf2 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf2(
		this.ref,
	)
}

// FuncScaleSelf2 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf2() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf2 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf2(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// TryScaleSelf2 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// HasFuncScaleSelf3 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf3() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf3(
		this.ref,
	)
}

// FuncScaleSelf3 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf3() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf3 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf3(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// TryScaleSelf3 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// HasFuncScaleSelf4 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf4() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf4(
		this.ref,
	)
}

// FuncScaleSelf4 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf4() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf4 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf4(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf4(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// TryScaleSelf4 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf4(scaleX float64, scaleY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasFuncScaleSelf5 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf5() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf5(
		this.ref,
	)
}

// FuncScaleSelf5 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf5() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf5 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf5(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf5(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScaleSelf5 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf5(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasFuncScaleSelf6 returns true if the method "DOMMatrix.scaleSelf" exists.
func (this DOMMatrix) HasFuncScaleSelf6() bool {
	return js.True == bindings.HasFuncDOMMatrixScaleSelf6(
		this.ref,
	)
}

// FuncScaleSelf6 returns the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) FuncScaleSelf6() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixScaleSelf6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleSelf6 calls the method "DOMMatrix.scaleSelf".
func (this DOMMatrix) ScaleSelf6() (ret DOMMatrix) {
	bindings.CallDOMMatrixScaleSelf6(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScaleSelf6 calls the method "DOMMatrix.scaleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScaleSelf6() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScaleSelf6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScale3dSelf returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasFuncScale3dSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixScale3dSelf(
		this.ref,
	)
}

// FuncScale3dSelf returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) FuncScale3dSelf() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScale3dSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3dSelf calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf(
		this.ref, js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// TryScale3dSelf calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScale3dSelf(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasFuncScale3dSelf1 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasFuncScale3dSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixScale3dSelf1(
		this.ref,
	)
}

// FuncScale3dSelf1 returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) FuncScale3dSelf1() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScale3dSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3dSelf1 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf1(scale float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf1(
		this.ref, js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScale3dSelf1 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScale3dSelf1(scale float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// HasFuncScale3dSelf2 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasFuncScale3dSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixScale3dSelf2(
		this.ref,
	)
}

// FuncScale3dSelf2 returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) FuncScale3dSelf2() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScale3dSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3dSelf2 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf2(scale float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf2(
		this.ref, js.Pointer(&ret),
		float64(scale),
		float64(originX),
	)

	return
}

// TryScale3dSelf2 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScale3dSelf2(scale float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
	)

	return
}

// HasFuncScale3dSelf3 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasFuncScale3dSelf3() bool {
	return js.True == bindings.HasFuncDOMMatrixScale3dSelf3(
		this.ref,
	)
}

// FuncScale3dSelf3 returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) FuncScale3dSelf3() (fn js.Func[func(scale float64) DOMMatrix]) {
	bindings.FuncDOMMatrixScale3dSelf3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3dSelf3 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf3(scale float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf3(
		this.ref, js.Pointer(&ret),
		float64(scale),
	)

	return
}

// TryScale3dSelf3 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScale3dSelf3(scale float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
	)

	return
}

// HasFuncScale3dSelf4 returns true if the method "DOMMatrix.scale3dSelf" exists.
func (this DOMMatrix) HasFuncScale3dSelf4() bool {
	return js.True == bindings.HasFuncDOMMatrixScale3dSelf4(
		this.ref,
	)
}

// FuncScale3dSelf4 returns the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) FuncScale3dSelf4() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixScale3dSelf4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3dSelf4 calls the method "DOMMatrix.scale3dSelf".
func (this DOMMatrix) Scale3dSelf4() (ret DOMMatrix) {
	bindings.CallDOMMatrixScale3dSelf4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScale3dSelf4 calls the method "DOMMatrix.scale3dSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryScale3dSelf4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixScale3dSelf4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotateSelf returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasFuncRotateSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateSelf(
		this.ref,
	)
}

// FuncRotateSelf returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) FuncRotateSelf() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateSelf calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf(
		this.ref, js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// TryRotateSelf calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateSelf(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// HasFuncRotateSelf1 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasFuncRotateSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateSelf1(
		this.ref,
	)
}

// FuncRotateSelf1 returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) FuncRotateSelf1() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateSelf1 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf1(rotX float64, rotY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf1(
		this.ref, js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
	)

	return
}

// TryRotateSelf1 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateSelf1(rotX float64, rotY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
	)

	return
}

// HasFuncRotateSelf2 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasFuncRotateSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateSelf2(
		this.ref,
	)
}

// FuncRotateSelf2 returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) FuncRotateSelf2() (fn js.Func[func(rotX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateSelf2 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf2(rotX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf2(
		this.ref, js.Pointer(&ret),
		float64(rotX),
	)

	return
}

// TryRotateSelf2 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateSelf2(rotX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
	)

	return
}

// HasFuncRotateSelf3 returns true if the method "DOMMatrix.rotateSelf" exists.
func (this DOMMatrix) HasFuncRotateSelf3() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateSelf3(
		this.ref,
	)
}

// FuncRotateSelf3 returns the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) FuncRotateSelf3() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixRotateSelf3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateSelf3 calls the method "DOMMatrix.rotateSelf".
func (this DOMMatrix) RotateSelf3() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateSelf3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotateSelf3 calls the method "DOMMatrix.rotateSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateSelf3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateSelf3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotateFromVectorSelf returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasFuncRotateFromVectorSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateFromVectorSelf(
		this.ref,
	)
}

// FuncRotateFromVectorSelf returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) FuncRotateFromVectorSelf() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateFromVectorSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVectorSelf calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateFromVectorSelf calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRotateFromVectorSelf1 returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasFuncRotateFromVectorSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateFromVectorSelf1(
		this.ref,
	)
}

// FuncRotateFromVectorSelf1 returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) FuncRotateFromVectorSelf1() (fn js.Func[func(x float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateFromVectorSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVectorSelf1 calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf1(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf1(
		this.ref, js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateFromVectorSelf1 calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf1(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasFuncRotateFromVectorSelf2 returns true if the method "DOMMatrix.rotateFromVectorSelf" exists.
func (this DOMMatrix) HasFuncRotateFromVectorSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateFromVectorSelf2(
		this.ref,
	)
}

// FuncRotateFromVectorSelf2 returns the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) FuncRotateFromVectorSelf2() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixRotateFromVectorSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVectorSelf2 calls the method "DOMMatrix.rotateFromVectorSelf".
func (this DOMMatrix) RotateFromVectorSelf2() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateFromVectorSelf2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotateFromVectorSelf2 calls the method "DOMMatrix.rotateFromVectorSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateFromVectorSelf2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateFromVectorSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotateAxisAngleSelf returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasFuncRotateAxisAngleSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateAxisAngleSelf(
		this.ref,
	)
}

// FuncRotateAxisAngleSelf returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) FuncRotateAxisAngleSelf() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateAxisAngleSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngleSelf calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// TryRotateAxisAngleSelf calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf(x float64, y float64, z float64, angle float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// HasFuncRotateAxisAngleSelf1 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasFuncRotateAxisAngleSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateAxisAngleSelf1(
		this.ref,
	)
}

// FuncRotateAxisAngleSelf1 returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) FuncRotateAxisAngleSelf1() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateAxisAngleSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngleSelf1 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf1(x float64, y float64, z float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// TryRotateAxisAngleSelf1 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf1(x float64, y float64, z float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// HasFuncRotateAxisAngleSelf2 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasFuncRotateAxisAngleSelf2() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateAxisAngleSelf2(
		this.ref,
	)
}

// FuncRotateAxisAngleSelf2 returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) FuncRotateAxisAngleSelf2() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateAxisAngleSelf2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngleSelf2 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf2(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateAxisAngleSelf2 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf2(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRotateAxisAngleSelf3 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasFuncRotateAxisAngleSelf3() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateAxisAngleSelf3(
		this.ref,
	)
}

// FuncRotateAxisAngleSelf3 returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) FuncRotateAxisAngleSelf3() (fn js.Func[func(x float64) DOMMatrix]) {
	bindings.FuncDOMMatrixRotateAxisAngleSelf3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngleSelf3 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf3(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf3(
		this.ref, js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateAxisAngleSelf3 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf3(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasFuncRotateAxisAngleSelf4 returns true if the method "DOMMatrix.rotateAxisAngleSelf" exists.
func (this DOMMatrix) HasFuncRotateAxisAngleSelf4() bool {
	return js.True == bindings.HasFuncDOMMatrixRotateAxisAngleSelf4(
		this.ref,
	)
}

// FuncRotateAxisAngleSelf4 returns the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) FuncRotateAxisAngleSelf4() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixRotateAxisAngleSelf4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngleSelf4 calls the method "DOMMatrix.rotateAxisAngleSelf".
func (this DOMMatrix) RotateAxisAngleSelf4() (ret DOMMatrix) {
	bindings.CallDOMMatrixRotateAxisAngleSelf4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotateAxisAngleSelf4 calls the method "DOMMatrix.rotateAxisAngleSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryRotateAxisAngleSelf4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixRotateAxisAngleSelf4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSkewXSelf returns true if the method "DOMMatrix.skewXSelf" exists.
func (this DOMMatrix) HasFuncSkewXSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixSkewXSelf(
		this.ref,
	)
}

// FuncSkewXSelf returns the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) FuncSkewXSelf() (fn js.Func[func(sx float64) DOMMatrix]) {
	bindings.FuncDOMMatrixSkewXSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewXSelf calls the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelf(sx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewXSelf(
		this.ref, js.Pointer(&ret),
		float64(sx),
	)

	return
}

// TrySkewXSelf calls the method "DOMMatrix.skewXSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TrySkewXSelf(sx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewXSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(sx),
	)

	return
}

// HasFuncSkewXSelf1 returns true if the method "DOMMatrix.skewXSelf" exists.
func (this DOMMatrix) HasFuncSkewXSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixSkewXSelf1(
		this.ref,
	)
}

// FuncSkewXSelf1 returns the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) FuncSkewXSelf1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixSkewXSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewXSelf1 calls the method "DOMMatrix.skewXSelf".
func (this DOMMatrix) SkewXSelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewXSelf1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkewXSelf1 calls the method "DOMMatrix.skewXSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TrySkewXSelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewXSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSkewYSelf returns true if the method "DOMMatrix.skewYSelf" exists.
func (this DOMMatrix) HasFuncSkewYSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixSkewYSelf(
		this.ref,
	)
}

// FuncSkewYSelf returns the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) FuncSkewYSelf() (fn js.Func[func(sy float64) DOMMatrix]) {
	bindings.FuncDOMMatrixSkewYSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewYSelf calls the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelf(sy float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewYSelf(
		this.ref, js.Pointer(&ret),
		float64(sy),
	)

	return
}

// TrySkewYSelf calls the method "DOMMatrix.skewYSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TrySkewYSelf(sy float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewYSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(sy),
	)

	return
}

// HasFuncSkewYSelf1 returns true if the method "DOMMatrix.skewYSelf" exists.
func (this DOMMatrix) HasFuncSkewYSelf1() bool {
	return js.True == bindings.HasFuncDOMMatrixSkewYSelf1(
		this.ref,
	)
}

// FuncSkewYSelf1 returns the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) FuncSkewYSelf1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixSkewYSelf1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewYSelf1 calls the method "DOMMatrix.skewYSelf".
func (this DOMMatrix) SkewYSelf1() (ret DOMMatrix) {
	bindings.CallDOMMatrixSkewYSelf1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkewYSelf1 calls the method "DOMMatrix.skewYSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TrySkewYSelf1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSkewYSelf1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInvertSelf returns true if the method "DOMMatrix.invertSelf" exists.
func (this DOMMatrix) HasFuncInvertSelf() bool {
	return js.True == bindings.HasFuncDOMMatrixInvertSelf(
		this.ref,
	)
}

// FuncInvertSelf returns the method "DOMMatrix.invertSelf".
func (this DOMMatrix) FuncInvertSelf() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixInvertSelf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InvertSelf calls the method "DOMMatrix.invertSelf".
func (this DOMMatrix) InvertSelf() (ret DOMMatrix) {
	bindings.CallDOMMatrixInvertSelf(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryInvertSelf calls the method "DOMMatrix.invertSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TryInvertSelf() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixInvertSelf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetMatrixValue returns true if the method "DOMMatrix.setMatrixValue" exists.
func (this DOMMatrix) HasFuncSetMatrixValue() bool {
	return js.True == bindings.HasFuncDOMMatrixSetMatrixValue(
		this.ref,
	)
}

// FuncSetMatrixValue returns the method "DOMMatrix.setMatrixValue".
func (this DOMMatrix) FuncSetMatrixValue() (fn js.Func[func(transformList js.String) DOMMatrix]) {
	bindings.FuncDOMMatrixSetMatrixValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetMatrixValue calls the method "DOMMatrix.setMatrixValue".
func (this DOMMatrix) SetMatrixValue(transformList js.String) (ret DOMMatrix) {
	bindings.CallDOMMatrixSetMatrixValue(
		this.ref, js.Pointer(&ret),
		transformList.Ref(),
	)

	return
}

// TrySetMatrixValue calls the method "DOMMatrix.setMatrixValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrix) TrySetMatrixValue(transformList js.String) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixSetMatrixValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *DOMMatrix2DInit) UpdateFrom(ref js.Ref) {
	bindings.DOMMatrix2DInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMMatrix2DInit) Update(ref js.Ref) {
	bindings.DOMMatrix2DInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMMatrix2DInit) FreeMembers(recursive bool) {
}

type SVGTransform struct {
	ref js.Ref
}

func (this SVGTransform) Once() SVGTransform {
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "SVGTransform.type".
//
// It returns ok=false if there is no such property.
func (this SVGTransform) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGTransformType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "SVGTransform.matrix".
//
// It returns ok=false if there is no such property.
func (this SVGTransform) Matrix() (ret DOMMatrix, ok bool) {
	ok = js.True == bindings.GetSVGTransformMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Angle returns the value of property "SVGTransform.angle".
//
// It returns ok=false if there is no such property.
func (this SVGTransform) Angle() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGTransformAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetMatrix returns true if the method "SVGTransform.setMatrix" exists.
func (this SVGTransform) HasFuncSetMatrix() bool {
	return js.True == bindings.HasFuncSVGTransformSetMatrix(
		this.ref,
	)
}

// FuncSetMatrix returns the method "SVGTransform.setMatrix".
func (this SVGTransform) FuncSetMatrix() (fn js.Func[func(matrix DOMMatrix2DInit)]) {
	bindings.FuncSVGTransformSetMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetMatrix calls the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrix(matrix DOMMatrix2DInit) (ret js.Void) {
	bindings.CallSVGTransformSetMatrix(
		this.ref, js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TrySetMatrix calls the method "SVGTransform.setMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetMatrix(matrix DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasFuncSetMatrix1 returns true if the method "SVGTransform.setMatrix" exists.
func (this SVGTransform) HasFuncSetMatrix1() bool {
	return js.True == bindings.HasFuncSVGTransformSetMatrix1(
		this.ref,
	)
}

// FuncSetMatrix1 returns the method "SVGTransform.setMatrix".
func (this SVGTransform) FuncSetMatrix1() (fn js.Func[func()]) {
	bindings.FuncSVGTransformSetMatrix1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetMatrix1 calls the method "SVGTransform.setMatrix".
func (this SVGTransform) SetMatrix1() (ret js.Void) {
	bindings.CallSVGTransformSetMatrix1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetMatrix1 calls the method "SVGTransform.setMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetMatrix1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetMatrix1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetTranslate returns true if the method "SVGTransform.setTranslate" exists.
func (this SVGTransform) HasFuncSetTranslate() bool {
	return js.True == bindings.HasFuncSVGTransformSetTranslate(
		this.ref,
	)
}

// FuncSetTranslate returns the method "SVGTransform.setTranslate".
func (this SVGTransform) FuncSetTranslate() (fn js.Func[func(tx float32, ty float32)]) {
	bindings.FuncSVGTransformSetTranslate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTranslate calls the method "SVGTransform.setTranslate".
func (this SVGTransform) SetTranslate(tx float32, ty float32) (ret js.Void) {
	bindings.CallSVGTransformSetTranslate(
		this.ref, js.Pointer(&ret),
		float32(tx),
		float32(ty),
	)

	return
}

// TrySetTranslate calls the method "SVGTransform.setTranslate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetTranslate(tx float32, ty float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetTranslate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(tx),
		float32(ty),
	)

	return
}

// HasFuncSetScale returns true if the method "SVGTransform.setScale" exists.
func (this SVGTransform) HasFuncSetScale() bool {
	return js.True == bindings.HasFuncSVGTransformSetScale(
		this.ref,
	)
}

// FuncSetScale returns the method "SVGTransform.setScale".
func (this SVGTransform) FuncSetScale() (fn js.Func[func(sx float32, sy float32)]) {
	bindings.FuncSVGTransformSetScale(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetScale calls the method "SVGTransform.setScale".
func (this SVGTransform) SetScale(sx float32, sy float32) (ret js.Void) {
	bindings.CallSVGTransformSetScale(
		this.ref, js.Pointer(&ret),
		float32(sx),
		float32(sy),
	)

	return
}

// TrySetScale calls the method "SVGTransform.setScale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetScale(sx float32, sy float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetScale(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(sx),
		float32(sy),
	)

	return
}

// HasFuncSetRotate returns true if the method "SVGTransform.setRotate" exists.
func (this SVGTransform) HasFuncSetRotate() bool {
	return js.True == bindings.HasFuncSVGTransformSetRotate(
		this.ref,
	)
}

// FuncSetRotate returns the method "SVGTransform.setRotate".
func (this SVGTransform) FuncSetRotate() (fn js.Func[func(angle float32, cx float32, cy float32)]) {
	bindings.FuncSVGTransformSetRotate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRotate calls the method "SVGTransform.setRotate".
func (this SVGTransform) SetRotate(angle float32, cx float32, cy float32) (ret js.Void) {
	bindings.CallSVGTransformSetRotate(
		this.ref, js.Pointer(&ret),
		float32(angle),
		float32(cx),
		float32(cy),
	)

	return
}

// TrySetRotate calls the method "SVGTransform.setRotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetRotate(angle float32, cx float32, cy float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetRotate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(angle),
		float32(cx),
		float32(cy),
	)

	return
}

// HasFuncSetSkewX returns true if the method "SVGTransform.setSkewX" exists.
func (this SVGTransform) HasFuncSetSkewX() bool {
	return js.True == bindings.HasFuncSVGTransformSetSkewX(
		this.ref,
	)
}

// FuncSetSkewX returns the method "SVGTransform.setSkewX".
func (this SVGTransform) FuncSetSkewX() (fn js.Func[func(angle float32)]) {
	bindings.FuncSVGTransformSetSkewX(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSkewX calls the method "SVGTransform.setSkewX".
func (this SVGTransform) SetSkewX(angle float32) (ret js.Void) {
	bindings.CallSVGTransformSetSkewX(
		this.ref, js.Pointer(&ret),
		float32(angle),
	)

	return
}

// TrySetSkewX calls the method "SVGTransform.setSkewX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetSkewX(angle float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetSkewX(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(angle),
	)

	return
}

// HasFuncSetSkewY returns true if the method "SVGTransform.setSkewY" exists.
func (this SVGTransform) HasFuncSetSkewY() bool {
	return js.True == bindings.HasFuncSVGTransformSetSkewY(
		this.ref,
	)
}

// FuncSetSkewY returns the method "SVGTransform.setSkewY".
func (this SVGTransform) FuncSetSkewY() (fn js.Func[func(angle float32)]) {
	bindings.FuncSVGTransformSetSkewY(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSkewY calls the method "SVGTransform.setSkewY".
func (this SVGTransform) SetSkewY(angle float32) (ret js.Void) {
	bindings.CallSVGTransformSetSkewY(
		this.ref, js.Pointer(&ret),
		float32(angle),
	)

	return
}

// TrySetSkewY calls the method "SVGTransform.setSkewY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransform) TrySetSkewY(angle float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformSetSkewY(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "DOMPointReadOnly.x".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DOMPointReadOnly.y".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "DOMPointReadOnly.z".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// W returns the value of property "DOMPointReadOnly.w".
//
// It returns ok=false if there is no such property.
func (this DOMPointReadOnly) W() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointReadOnlyW(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFromPoint returns true if the static method "DOMPointReadOnly.fromPoint" exists.
func (this DOMPointReadOnly) HasFuncFromPoint() bool {
	return js.True == bindings.HasFuncDOMPointReadOnlyFromPoint(
		this.ref,
	)
}

// FuncFromPoint returns the static method "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FuncFromPoint() (fn js.Func[func(other DOMPointInit) DOMPointReadOnly]) {
	bindings.FuncDOMPointReadOnlyFromPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromPoint calls the static method "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPoint(other DOMPointInit) (ret DOMPointReadOnly) {
	bindings.CallDOMPointReadOnlyFromPoint(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromPoint calls the static method "DOMPointReadOnly.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPointReadOnly) TryFromPoint(other DOMPointInit) (ret DOMPointReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyFromPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromPoint1 returns true if the static method "DOMPointReadOnly.fromPoint" exists.
func (this DOMPointReadOnly) HasFuncFromPoint1() bool {
	return js.True == bindings.HasFuncDOMPointReadOnlyFromPoint1(
		this.ref,
	)
}

// FuncFromPoint1 returns the static method "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FuncFromPoint1() (fn js.Func[func() DOMPointReadOnly]) {
	bindings.FuncDOMPointReadOnlyFromPoint1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromPoint1 calls the static method "DOMPointReadOnly.fromPoint".
func (this DOMPointReadOnly) FromPoint1() (ret DOMPointReadOnly) {
	bindings.CallDOMPointReadOnlyFromPoint1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromPoint1 calls the static method "DOMPointReadOnly.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPointReadOnly) TryFromPoint1() (ret DOMPointReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyFromPoint1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMatrixTransform returns true if the method "DOMPointReadOnly.matrixTransform" exists.
func (this DOMPointReadOnly) HasFuncMatrixTransform() bool {
	return js.True == bindings.HasFuncDOMPointReadOnlyMatrixTransform(
		this.ref,
	)
}

// FuncMatrixTransform returns the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) FuncMatrixTransform() (fn js.Func[func(matrix DOMMatrixInit) DOMPoint]) {
	bindings.FuncDOMPointReadOnlyMatrixTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatrixTransform calls the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransform(matrix DOMMatrixInit) (ret DOMPoint) {
	bindings.CallDOMPointReadOnlyMatrixTransform(
		this.ref, js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryMatrixTransform calls the method "DOMPointReadOnly.matrixTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPointReadOnly) TryMatrixTransform(matrix DOMMatrixInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyMatrixTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasFuncMatrixTransform1 returns true if the method "DOMPointReadOnly.matrixTransform" exists.
func (this DOMPointReadOnly) HasFuncMatrixTransform1() bool {
	return js.True == bindings.HasFuncDOMPointReadOnlyMatrixTransform1(
		this.ref,
	)
}

// FuncMatrixTransform1 returns the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) FuncMatrixTransform1() (fn js.Func[func() DOMPoint]) {
	bindings.FuncDOMPointReadOnlyMatrixTransform1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatrixTransform1 calls the method "DOMPointReadOnly.matrixTransform".
func (this DOMPointReadOnly) MatrixTransform1() (ret DOMPoint) {
	bindings.CallDOMPointReadOnlyMatrixTransform1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMatrixTransform1 calls the method "DOMPointReadOnly.matrixTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPointReadOnly) TryMatrixTransform1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyMatrixTransform1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "DOMPointReadOnly.toJSON" exists.
func (this DOMPointReadOnly) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncDOMPointReadOnlyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "DOMPointReadOnly.toJSON".
func (this DOMPointReadOnly) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncDOMPointReadOnlyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "DOMPointReadOnly.toJSON".
func (this DOMPointReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMPointReadOnlyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMPointReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPointReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointReadOnlyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGAnimatedRect struct {
	ref js.Ref
}

func (this SVGAnimatedRect) Once() SVGAnimatedRect {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedRect.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedRect) BaseVal() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedRectBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedRect.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedRect) AnimVal() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedRectAnimVal(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "SVGPreserveAspectRatio.align".
//
// It returns ok=false if there is no such property.
func (this SVGPreserveAspectRatio) Align() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGPreserveAspectRatioAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "SVGPreserveAspectRatio.align" to val.
//
// It returns false if the property cannot be set.
func (this SVGPreserveAspectRatio) SetAlign(val uint16) bool {
	return js.True == bindings.SetSVGPreserveAspectRatioAlign(
		this.ref,
		uint32(val),
	)
}

// MeetOrSlice returns the value of property "SVGPreserveAspectRatio.meetOrSlice".
//
// It returns ok=false if there is no such property.
func (this SVGPreserveAspectRatio) MeetOrSlice() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGPreserveAspectRatioMeetOrSlice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMeetOrSlice sets the value of property "SVGPreserveAspectRatio.meetOrSlice" to val.
//
// It returns false if the property cannot be set.
func (this SVGPreserveAspectRatio) SetMeetOrSlice(val uint16) bool {
	return js.True == bindings.SetSVGPreserveAspectRatioMeetOrSlice(
		this.ref,
		uint32(val),
	)
}

type SVGAnimatedPreserveAspectRatio struct {
	ref js.Ref
}

func (this SVGAnimatedPreserveAspectRatio) Once() SVGAnimatedPreserveAspectRatio {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedPreserveAspectRatio.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) BaseVal() (ret SVGPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedPreserveAspectRatioBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedPreserveAspectRatio.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedPreserveAspectRatio) AnimVal() (ret SVGPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedPreserveAspectRatioAnimVal(
		this.ref, js.Pointer(&ret),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGSVGElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGSVGElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGSVGElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGSVGElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentScale returns the value of property "SVGSVGElement.currentScale".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) CurrentScale() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementCurrentScale(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCurrentScale sets the value of property "SVGSVGElement.currentScale" to val.
//
// It returns false if the property cannot be set.
func (this SVGSVGElement) SetCurrentScale(val float32) bool {
	return js.True == bindings.SetSVGSVGElementCurrentScale(
		this.ref,
		float32(val),
	)
}

// CurrentTranslate returns the value of property "SVGSVGElement.currentTranslate".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) CurrentTranslate() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementCurrentTranslate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ViewBox returns the value of property "SVGSVGElement.viewBox".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementViewBox(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGSVGElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGSVGElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGSVGElementPreserveAspectRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetIntersectionList returns true if the method "SVGSVGElement.getIntersectionList" exists.
func (this SVGSVGElement) HasFuncGetIntersectionList() bool {
	return js.True == bindings.HasFuncSVGSVGElementGetIntersectionList(
		this.ref,
	)
}

// FuncGetIntersectionList returns the method "SVGSVGElement.getIntersectionList".
func (this SVGSVGElement) FuncGetIntersectionList() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	bindings.FuncSVGSVGElementGetIntersectionList(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIntersectionList calls the method "SVGSVGElement.getIntersectionList".
func (this SVGSVGElement) GetIntersectionList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList) {
	bindings.CallSVGSVGElementGetIntersectionList(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// TryGetIntersectionList calls the method "SVGSVGElement.getIntersectionList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryGetIntersectionList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetIntersectionList(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// HasFuncGetEnclosureList returns true if the method "SVGSVGElement.getEnclosureList" exists.
func (this SVGSVGElement) HasFuncGetEnclosureList() bool {
	return js.True == bindings.HasFuncSVGSVGElementGetEnclosureList(
		this.ref,
	)
}

// FuncGetEnclosureList returns the method "SVGSVGElement.getEnclosureList".
func (this SVGSVGElement) FuncGetEnclosureList() (fn js.Func[func(rect DOMRectReadOnly, referenceElement SVGElement) NodeList]) {
	bindings.FuncSVGSVGElementGetEnclosureList(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetEnclosureList calls the method "SVGSVGElement.getEnclosureList".
func (this SVGSVGElement) GetEnclosureList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList) {
	bindings.CallSVGSVGElementGetEnclosureList(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// TryGetEnclosureList calls the method "SVGSVGElement.getEnclosureList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryGetEnclosureList(rect DOMRectReadOnly, referenceElement SVGElement) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetEnclosureList(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		referenceElement.Ref(),
	)

	return
}

// HasFuncCheckIntersection returns true if the method "SVGSVGElement.checkIntersection" exists.
func (this SVGSVGElement) HasFuncCheckIntersection() bool {
	return js.True == bindings.HasFuncSVGSVGElementCheckIntersection(
		this.ref,
	)
}

// FuncCheckIntersection returns the method "SVGSVGElement.checkIntersection".
func (this SVGSVGElement) FuncCheckIntersection() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	bindings.FuncSVGSVGElementCheckIntersection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckIntersection calls the method "SVGSVGElement.checkIntersection".
func (this SVGSVGElement) CheckIntersection(element SVGElement, rect DOMRectReadOnly) (ret bool) {
	bindings.CallSVGSVGElementCheckIntersection(
		this.ref, js.Pointer(&ret),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// TryCheckIntersection calls the method "SVGSVGElement.checkIntersection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCheckIntersection(element SVGElement, rect DOMRectReadOnly) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCheckIntersection(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// HasFuncCheckEnclosure returns true if the method "SVGSVGElement.checkEnclosure" exists.
func (this SVGSVGElement) HasFuncCheckEnclosure() bool {
	return js.True == bindings.HasFuncSVGSVGElementCheckEnclosure(
		this.ref,
	)
}

// FuncCheckEnclosure returns the method "SVGSVGElement.checkEnclosure".
func (this SVGSVGElement) FuncCheckEnclosure() (fn js.Func[func(element SVGElement, rect DOMRectReadOnly) bool]) {
	bindings.FuncSVGSVGElementCheckEnclosure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckEnclosure calls the method "SVGSVGElement.checkEnclosure".
func (this SVGSVGElement) CheckEnclosure(element SVGElement, rect DOMRectReadOnly) (ret bool) {
	bindings.CallSVGSVGElementCheckEnclosure(
		this.ref, js.Pointer(&ret),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// TryCheckEnclosure calls the method "SVGSVGElement.checkEnclosure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCheckEnclosure(element SVGElement, rect DOMRectReadOnly) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCheckEnclosure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		rect.Ref(),
	)

	return
}

// HasFuncDeselectAll returns true if the method "SVGSVGElement.deselectAll" exists.
func (this SVGSVGElement) HasFuncDeselectAll() bool {
	return js.True == bindings.HasFuncSVGSVGElementDeselectAll(
		this.ref,
	)
}

// FuncDeselectAll returns the method "SVGSVGElement.deselectAll".
func (this SVGSVGElement) FuncDeselectAll() (fn js.Func[func()]) {
	bindings.FuncSVGSVGElementDeselectAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeselectAll calls the method "SVGSVGElement.deselectAll".
func (this SVGSVGElement) DeselectAll() (ret js.Void) {
	bindings.CallSVGSVGElementDeselectAll(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeselectAll calls the method "SVGSVGElement.deselectAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryDeselectAll() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementDeselectAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGNumber returns true if the method "SVGSVGElement.createSVGNumber" exists.
func (this SVGSVGElement) HasFuncCreateSVGNumber() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGNumber(
		this.ref,
	)
}

// FuncCreateSVGNumber returns the method "SVGSVGElement.createSVGNumber".
func (this SVGSVGElement) FuncCreateSVGNumber() (fn js.Func[func() SVGNumber]) {
	bindings.FuncSVGSVGElementCreateSVGNumber(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGNumber calls the method "SVGSVGElement.createSVGNumber".
func (this SVGSVGElement) CreateSVGNumber() (ret SVGNumber) {
	bindings.CallSVGSVGElementCreateSVGNumber(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGNumber calls the method "SVGSVGElement.createSVGNumber"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGNumber() (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGNumber(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGLength returns true if the method "SVGSVGElement.createSVGLength" exists.
func (this SVGSVGElement) HasFuncCreateSVGLength() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGLength(
		this.ref,
	)
}

// FuncCreateSVGLength returns the method "SVGSVGElement.createSVGLength".
func (this SVGSVGElement) FuncCreateSVGLength() (fn js.Func[func() SVGLength]) {
	bindings.FuncSVGSVGElementCreateSVGLength(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGLength calls the method "SVGSVGElement.createSVGLength".
func (this SVGSVGElement) CreateSVGLength() (ret SVGLength) {
	bindings.CallSVGSVGElementCreateSVGLength(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGLength calls the method "SVGSVGElement.createSVGLength"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGLength() (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGLength(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGAngle returns true if the method "SVGSVGElement.createSVGAngle" exists.
func (this SVGSVGElement) HasFuncCreateSVGAngle() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGAngle(
		this.ref,
	)
}

// FuncCreateSVGAngle returns the method "SVGSVGElement.createSVGAngle".
func (this SVGSVGElement) FuncCreateSVGAngle() (fn js.Func[func() SVGAngle]) {
	bindings.FuncSVGSVGElementCreateSVGAngle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGAngle calls the method "SVGSVGElement.createSVGAngle".
func (this SVGSVGElement) CreateSVGAngle() (ret SVGAngle) {
	bindings.CallSVGSVGElementCreateSVGAngle(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGAngle calls the method "SVGSVGElement.createSVGAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGAngle() (ret SVGAngle, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGAngle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGPoint returns true if the method "SVGSVGElement.createSVGPoint" exists.
func (this SVGSVGElement) HasFuncCreateSVGPoint() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGPoint(
		this.ref,
	)
}

// FuncCreateSVGPoint returns the method "SVGSVGElement.createSVGPoint".
func (this SVGSVGElement) FuncCreateSVGPoint() (fn js.Func[func() DOMPoint]) {
	bindings.FuncSVGSVGElementCreateSVGPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGPoint calls the method "SVGSVGElement.createSVGPoint".
func (this SVGSVGElement) CreateSVGPoint() (ret DOMPoint) {
	bindings.CallSVGSVGElementCreateSVGPoint(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGPoint calls the method "SVGSVGElement.createSVGPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGPoint() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGMatrix returns true if the method "SVGSVGElement.createSVGMatrix" exists.
func (this SVGSVGElement) HasFuncCreateSVGMatrix() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGMatrix(
		this.ref,
	)
}

// FuncCreateSVGMatrix returns the method "SVGSVGElement.createSVGMatrix".
func (this SVGSVGElement) FuncCreateSVGMatrix() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncSVGSVGElementCreateSVGMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGMatrix calls the method "SVGSVGElement.createSVGMatrix".
func (this SVGSVGElement) CreateSVGMatrix() (ret DOMMatrix) {
	bindings.CallSVGSVGElementCreateSVGMatrix(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGMatrix calls the method "SVGSVGElement.createSVGMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGMatrix() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGRect returns true if the method "SVGSVGElement.createSVGRect" exists.
func (this SVGSVGElement) HasFuncCreateSVGRect() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGRect(
		this.ref,
	)
}

// FuncCreateSVGRect returns the method "SVGSVGElement.createSVGRect".
func (this SVGSVGElement) FuncCreateSVGRect() (fn js.Func[func() DOMRect]) {
	bindings.FuncSVGSVGElementCreateSVGRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGRect calls the method "SVGSVGElement.createSVGRect".
func (this SVGSVGElement) CreateSVGRect() (ret DOMRect) {
	bindings.CallSVGSVGElementCreateSVGRect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGRect calls the method "SVGSVGElement.createSVGRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGRect() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGTransform returns true if the method "SVGSVGElement.createSVGTransform" exists.
func (this SVGSVGElement) HasFuncCreateSVGTransform() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGTransform(
		this.ref,
	)
}

// FuncCreateSVGTransform returns the method "SVGSVGElement.createSVGTransform".
func (this SVGSVGElement) FuncCreateSVGTransform() (fn js.Func[func() SVGTransform]) {
	bindings.FuncSVGSVGElementCreateSVGTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGTransform calls the method "SVGSVGElement.createSVGTransform".
func (this SVGSVGElement) CreateSVGTransform() (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransform(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransform calls the method "SVGSVGElement.createSVGTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransform() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSVGTransformFromMatrix returns true if the method "SVGSVGElement.createSVGTransformFromMatrix" exists.
func (this SVGSVGElement) HasFuncCreateSVGTransformFromMatrix() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGTransformFromMatrix(
		this.ref,
	)
}

// FuncCreateSVGTransformFromMatrix returns the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) FuncCreateSVGTransformFromMatrix() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	bindings.FuncSVGSVGElementCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGTransformFromMatrix calls the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryCreateSVGTransformFromMatrix calls the method "SVGSVGElement.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasFuncCreateSVGTransformFromMatrix1 returns true if the method "SVGSVGElement.createSVGTransformFromMatrix" exists.
func (this SVGSVGElement) HasFuncCreateSVGTransformFromMatrix1() bool {
	return js.True == bindings.HasFuncSVGSVGElementCreateSVGTransformFromMatrix1(
		this.ref,
	)
}

// FuncCreateSVGTransformFromMatrix1 returns the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) FuncCreateSVGTransformFromMatrix1() (fn js.Func[func() SVGTransform]) {
	bindings.FuncSVGSVGElementCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGTransformFromMatrix1 calls the method "SVGSVGElement.createSVGTransformFromMatrix".
func (this SVGSVGElement) CreateSVGTransformFromMatrix1() (ret SVGTransform) {
	bindings.CallSVGSVGElementCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransformFromMatrix1 calls the method "SVGSVGElement.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryCreateSVGTransformFromMatrix1() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetElementById returns true if the method "SVGSVGElement.getElementById" exists.
func (this SVGSVGElement) HasFuncGetElementById() bool {
	return js.True == bindings.HasFuncSVGSVGElementGetElementById(
		this.ref,
	)
}

// FuncGetElementById returns the method "SVGSVGElement.getElementById".
func (this SVGSVGElement) FuncGetElementById() (fn js.Func[func(elementId js.String) Element]) {
	bindings.FuncSVGSVGElementGetElementById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementById calls the method "SVGSVGElement.getElementById".
func (this SVGSVGElement) GetElementById(elementId js.String) (ret Element) {
	bindings.CallSVGSVGElementGetElementById(
		this.ref, js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "SVGSVGElement.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetElementById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasFuncSuspendRedraw returns true if the method "SVGSVGElement.suspendRedraw" exists.
func (this SVGSVGElement) HasFuncSuspendRedraw() bool {
	return js.True == bindings.HasFuncSVGSVGElementSuspendRedraw(
		this.ref,
	)
}

// FuncSuspendRedraw returns the method "SVGSVGElement.suspendRedraw".
func (this SVGSVGElement) FuncSuspendRedraw() (fn js.Func[func(maxWaitMilliseconds uint32) uint32]) {
	bindings.FuncSVGSVGElementSuspendRedraw(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SuspendRedraw calls the method "SVGSVGElement.suspendRedraw".
func (this SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint32) (ret uint32) {
	bindings.CallSVGSVGElementSuspendRedraw(
		this.ref, js.Pointer(&ret),
		uint32(maxWaitMilliseconds),
	)

	return
}

// TrySuspendRedraw calls the method "SVGSVGElement.suspendRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TrySuspendRedraw(maxWaitMilliseconds uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementSuspendRedraw(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(maxWaitMilliseconds),
	)

	return
}

// HasFuncUnsuspendRedraw returns true if the method "SVGSVGElement.unsuspendRedraw" exists.
func (this SVGSVGElement) HasFuncUnsuspendRedraw() bool {
	return js.True == bindings.HasFuncSVGSVGElementUnsuspendRedraw(
		this.ref,
	)
}

// FuncUnsuspendRedraw returns the method "SVGSVGElement.unsuspendRedraw".
func (this SVGSVGElement) FuncUnsuspendRedraw() (fn js.Func[func(suspendHandleID uint32)]) {
	bindings.FuncSVGSVGElementUnsuspendRedraw(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UnsuspendRedraw calls the method "SVGSVGElement.unsuspendRedraw".
func (this SVGSVGElement) UnsuspendRedraw(suspendHandleID uint32) (ret js.Void) {
	bindings.CallSVGSVGElementUnsuspendRedraw(
		this.ref, js.Pointer(&ret),
		uint32(suspendHandleID),
	)

	return
}

// TryUnsuspendRedraw calls the method "SVGSVGElement.unsuspendRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryUnsuspendRedraw(suspendHandleID uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnsuspendRedraw(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(suspendHandleID),
	)

	return
}

// HasFuncUnsuspendRedrawAll returns true if the method "SVGSVGElement.unsuspendRedrawAll" exists.
func (this SVGSVGElement) HasFuncUnsuspendRedrawAll() bool {
	return js.True == bindings.HasFuncSVGSVGElementUnsuspendRedrawAll(
		this.ref,
	)
}

// FuncUnsuspendRedrawAll returns the method "SVGSVGElement.unsuspendRedrawAll".
func (this SVGSVGElement) FuncUnsuspendRedrawAll() (fn js.Func[func()]) {
	bindings.FuncSVGSVGElementUnsuspendRedrawAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UnsuspendRedrawAll calls the method "SVGSVGElement.unsuspendRedrawAll".
func (this SVGSVGElement) UnsuspendRedrawAll() (ret js.Void) {
	bindings.CallSVGSVGElementUnsuspendRedrawAll(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUnsuspendRedrawAll calls the method "SVGSVGElement.unsuspendRedrawAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryUnsuspendRedrawAll() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnsuspendRedrawAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncForceRedraw returns true if the method "SVGSVGElement.forceRedraw" exists.
func (this SVGSVGElement) HasFuncForceRedraw() bool {
	return js.True == bindings.HasFuncSVGSVGElementForceRedraw(
		this.ref,
	)
}

// FuncForceRedraw returns the method "SVGSVGElement.forceRedraw".
func (this SVGSVGElement) FuncForceRedraw() (fn js.Func[func()]) {
	bindings.FuncSVGSVGElementForceRedraw(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ForceRedraw calls the method "SVGSVGElement.forceRedraw".
func (this SVGSVGElement) ForceRedraw() (ret js.Void) {
	bindings.CallSVGSVGElementForceRedraw(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForceRedraw calls the method "SVGSVGElement.forceRedraw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryForceRedraw() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementForceRedraw(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPauseAnimations returns true if the method "SVGSVGElement.pauseAnimations" exists.
func (this SVGSVGElement) HasFuncPauseAnimations() bool {
	return js.True == bindings.HasFuncSVGSVGElementPauseAnimations(
		this.ref,
	)
}

// FuncPauseAnimations returns the method "SVGSVGElement.pauseAnimations".
func (this SVGSVGElement) FuncPauseAnimations() (fn js.Func[func()]) {
	bindings.FuncSVGSVGElementPauseAnimations(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PauseAnimations calls the method "SVGSVGElement.pauseAnimations".
func (this SVGSVGElement) PauseAnimations() (ret js.Void) {
	bindings.CallSVGSVGElementPauseAnimations(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPauseAnimations calls the method "SVGSVGElement.pauseAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryPauseAnimations() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementPauseAnimations(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUnpauseAnimations returns true if the method "SVGSVGElement.unpauseAnimations" exists.
func (this SVGSVGElement) HasFuncUnpauseAnimations() bool {
	return js.True == bindings.HasFuncSVGSVGElementUnpauseAnimations(
		this.ref,
	)
}

// FuncUnpauseAnimations returns the method "SVGSVGElement.unpauseAnimations".
func (this SVGSVGElement) FuncUnpauseAnimations() (fn js.Func[func()]) {
	bindings.FuncSVGSVGElementUnpauseAnimations(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UnpauseAnimations calls the method "SVGSVGElement.unpauseAnimations".
func (this SVGSVGElement) UnpauseAnimations() (ret js.Void) {
	bindings.CallSVGSVGElementUnpauseAnimations(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUnpauseAnimations calls the method "SVGSVGElement.unpauseAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryUnpauseAnimations() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementUnpauseAnimations(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAnimationsPaused returns true if the method "SVGSVGElement.animationsPaused" exists.
func (this SVGSVGElement) HasFuncAnimationsPaused() bool {
	return js.True == bindings.HasFuncSVGSVGElementAnimationsPaused(
		this.ref,
	)
}

// FuncAnimationsPaused returns the method "SVGSVGElement.animationsPaused".
func (this SVGSVGElement) FuncAnimationsPaused() (fn js.Func[func() bool]) {
	bindings.FuncSVGSVGElementAnimationsPaused(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AnimationsPaused calls the method "SVGSVGElement.animationsPaused".
func (this SVGSVGElement) AnimationsPaused() (ret bool) {
	bindings.CallSVGSVGElementAnimationsPaused(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAnimationsPaused calls the method "SVGSVGElement.animationsPaused"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryAnimationsPaused() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementAnimationsPaused(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCurrentTime returns true if the method "SVGSVGElement.getCurrentTime" exists.
func (this SVGSVGElement) HasFuncGetCurrentTime() bool {
	return js.True == bindings.HasFuncSVGSVGElementGetCurrentTime(
		this.ref,
	)
}

// FuncGetCurrentTime returns the method "SVGSVGElement.getCurrentTime".
func (this SVGSVGElement) FuncGetCurrentTime() (fn js.Func[func() float32]) {
	bindings.FuncSVGSVGElementGetCurrentTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentTime calls the method "SVGSVGElement.getCurrentTime".
func (this SVGSVGElement) GetCurrentTime() (ret float32) {
	bindings.CallSVGSVGElementGetCurrentTime(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTime calls the method "SVGSVGElement.getCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TryGetCurrentTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementGetCurrentTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCurrentTime returns true if the method "SVGSVGElement.setCurrentTime" exists.
func (this SVGSVGElement) HasFuncSetCurrentTime() bool {
	return js.True == bindings.HasFuncSVGSVGElementSetCurrentTime(
		this.ref,
	)
}

// FuncSetCurrentTime returns the method "SVGSVGElement.setCurrentTime".
func (this SVGSVGElement) FuncSetCurrentTime() (fn js.Func[func(seconds float32)]) {
	bindings.FuncSVGSVGElementSetCurrentTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCurrentTime calls the method "SVGSVGElement.setCurrentTime".
func (this SVGSVGElement) SetCurrentTime(seconds float32) (ret js.Void) {
	bindings.CallSVGSVGElementSetCurrentTime(
		this.ref, js.Pointer(&ret),
		float32(seconds),
	)

	return
}

// TrySetCurrentTime calls the method "SVGSVGElement.setCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGSVGElement) TrySetCurrentTime(seconds float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGSVGElementSetCurrentTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(seconds),
	)

	return
}

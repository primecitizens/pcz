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

type OptionalEffectTiming struct {
	// Delay is "OptionalEffectTiming.delay"
	//
	// Optional
	//
	// NOTE: FFI_USE_Delay MUST be set to true to make this field effective.
	Delay float64
	// EndDelay is "OptionalEffectTiming.endDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndDelay MUST be set to true to make this field effective.
	EndDelay float64
	// Fill is "OptionalEffectTiming.fill"
	//
	// Optional
	Fill FillMode
	// IterationStart is "OptionalEffectTiming.iterationStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_IterationStart MUST be set to true to make this field effective.
	IterationStart float64
	// Iterations is "OptionalEffectTiming.iterations"
	//
	// Optional
	//
	// NOTE: FFI_USE_Iterations MUST be set to true to make this field effective.
	Iterations float64
	// Duration is "OptionalEffectTiming.duration"
	//
	// Optional
	Duration OneOf_Float64_String
	// Direction is "OptionalEffectTiming.direction"
	//
	// Optional
	Direction PlaybackDirection
	// Easing is "OptionalEffectTiming.easing"
	//
	// Optional
	Easing js.String
	// PlaybackRate is "OptionalEffectTiming.playbackRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_PlaybackRate MUST be set to true to make this field effective.
	PlaybackRate float64

	FFI_USE_Delay          bool // for Delay.
	FFI_USE_EndDelay       bool // for EndDelay.
	FFI_USE_IterationStart bool // for IterationStart.
	FFI_USE_Iterations     bool // for Iterations.
	FFI_USE_PlaybackRate   bool // for PlaybackRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OptionalEffectTiming with all fields set.
func (p OptionalEffectTiming) FromRef(ref js.Ref) OptionalEffectTiming {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OptionalEffectTiming in the application heap.
func (p OptionalEffectTiming) New() js.Ref {
	return bindings.OptionalEffectTimingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OptionalEffectTiming) UpdateFrom(ref js.Ref) {
	bindings.OptionalEffectTimingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OptionalEffectTiming) Update(ref js.Ref) {
	bindings.OptionalEffectTimingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Float64_EffectTiming struct {
	ref js.Ref
}

func (x OneOf_Float64_EffectTiming) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_EffectTiming) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_EffectTiming) FromRef(ref js.Ref) OneOf_Float64_EffectTiming {
	return OneOf_Float64_EffectTiming{
		ref: ref,
	}
}

func (x OneOf_Float64_EffectTiming) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_EffectTiming) EffectTiming() EffectTiming {
	var ret EffectTiming
	ret.UpdateFrom(x.ref)
	return ret
}

type AnimationNodeList struct {
	ref js.Ref
}

func (this AnimationNodeList) Once() AnimationNodeList {
	this.Ref().Once()
	return this
}

func (this AnimationNodeList) Ref() js.Ref {
	return this.ref
}

func (this AnimationNodeList) FromRef(ref js.Ref) AnimationNodeList {
	this.ref = ref
	return this
}

func (this AnimationNodeList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "AnimationNodeList.length".
//
// It returns ok=false if there is no such property.
func (this AnimationNodeList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnimationNodeListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "AnimationNodeList.item" exists.
func (this AnimationNodeList) HasItem() bool {
	return js.True == bindings.HasAnimationNodeListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "AnimationNodeList.item".
func (this AnimationNodeList) ItemFunc() (fn js.Func[func(index uint32) AnimationEffect]) {
	return fn.FromRef(
		bindings.AnimationNodeListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "AnimationNodeList.item".
func (this AnimationNodeList) Item(index uint32) (ret AnimationEffect) {
	bindings.CallAnimationNodeListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "AnimationNodeList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationNodeList) TryItem(index uint32) (ret AnimationEffect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationNodeListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewGroupEffect(children js.Array[AnimationEffect], timing OneOf_Float64_EffectTiming) (ret GroupEffect) {
	ret.ref = bindings.NewGroupEffectByGroupEffect(
		children.Ref(),
		timing.Ref())
	return
}

func NewGroupEffectByGroupEffect1(children js.Array[AnimationEffect]) (ret GroupEffect) {
	ret.ref = bindings.NewGroupEffectByGroupEffect1(
		children.Ref())
	return
}

type GroupEffect struct {
	ref js.Ref
}

func (this GroupEffect) Once() GroupEffect {
	this.Ref().Once()
	return this
}

func (this GroupEffect) Ref() js.Ref {
	return this.ref
}

func (this GroupEffect) FromRef(ref js.Ref) GroupEffect {
	this.ref = ref
	return this
}

func (this GroupEffect) Free() {
	this.Ref().Free()
}

// Children returns the value of property "GroupEffect.children".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) Children() (ret AnimationNodeList, ok bool) {
	ok = js.True == bindings.GetGroupEffectChildren(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstChild returns the value of property "GroupEffect.firstChild".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) FirstChild() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetGroupEffectFirstChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastChild returns the value of property "GroupEffect.lastChild".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) LastChild() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetGroupEffectLastChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClone returns true if the method "GroupEffect.clone" exists.
func (this GroupEffect) HasClone() bool {
	return js.True == bindings.HasGroupEffectClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "GroupEffect.clone".
func (this GroupEffect) CloneFunc() (fn js.Func[func() GroupEffect]) {
	return fn.FromRef(
		bindings.GroupEffectCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "GroupEffect.clone".
func (this GroupEffect) Clone() (ret GroupEffect) {
	bindings.CallGroupEffectClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "GroupEffect.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GroupEffect) TryClone() (ret GroupEffect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPrepend returns true if the method "GroupEffect.prepend" exists.
func (this GroupEffect) HasPrepend() bool {
	return js.True == bindings.HasGroupEffectPrepend(
		this.Ref(),
	)
}

// PrependFunc returns the method "GroupEffect.prepend".
func (this GroupEffect) PrependFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.GroupEffectPrependFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "GroupEffect.prepend".
func (this GroupEffect) Prepend(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallGroupEffectPrepend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryPrepend calls the method "GroupEffect.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GroupEffect) TryPrepend(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectPrepend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasAppend returns true if the method "GroupEffect.append" exists.
func (this GroupEffect) HasAppend() bool {
	return js.True == bindings.HasGroupEffectAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "GroupEffect.append".
func (this GroupEffect) AppendFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.GroupEffectAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "GroupEffect.append".
func (this GroupEffect) Append(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallGroupEffectAppend(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryAppend calls the method "GroupEffect.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GroupEffect) TryAppend(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

type AnimationEffect struct {
	ref js.Ref
}

func (this AnimationEffect) Once() AnimationEffect {
	this.Ref().Once()
	return this
}

func (this AnimationEffect) Ref() js.Ref {
	return this.ref
}

func (this AnimationEffect) FromRef(ref js.Ref) AnimationEffect {
	this.ref = ref
	return this
}

func (this AnimationEffect) Free() {
	this.Ref().Free()
}

// Parent returns the value of property "AnimationEffect.parent".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) Parent() (ret GroupEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectParent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "AnimationEffect.previousSibling".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) PreviousSibling() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectPreviousSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "AnimationEffect.nextSibling".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) NextSibling() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectNextSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetTiming returns true if the method "AnimationEffect.getTiming" exists.
func (this AnimationEffect) HasGetTiming() bool {
	return js.True == bindings.HasAnimationEffectGetTiming(
		this.Ref(),
	)
}

// GetTimingFunc returns the method "AnimationEffect.getTiming".
func (this AnimationEffect) GetTimingFunc() (fn js.Func[func() EffectTiming]) {
	return fn.FromRef(
		bindings.AnimationEffectGetTimingFunc(
			this.Ref(),
		),
	)
}

// GetTiming calls the method "AnimationEffect.getTiming".
func (this AnimationEffect) GetTiming() (ret EffectTiming) {
	bindings.CallAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTiming calls the method "AnimationEffect.getTiming"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryGetTiming() (ret EffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetComputedTiming returns true if the method "AnimationEffect.getComputedTiming" exists.
func (this AnimationEffect) HasGetComputedTiming() bool {
	return js.True == bindings.HasAnimationEffectGetComputedTiming(
		this.Ref(),
	)
}

// GetComputedTimingFunc returns the method "AnimationEffect.getComputedTiming".
func (this AnimationEffect) GetComputedTimingFunc() (fn js.Func[func() ComputedEffectTiming]) {
	return fn.FromRef(
		bindings.AnimationEffectGetComputedTimingFunc(
			this.Ref(),
		),
	)
}

// GetComputedTiming calls the method "AnimationEffect.getComputedTiming".
func (this AnimationEffect) GetComputedTiming() (ret ComputedEffectTiming) {
	bindings.CallAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetComputedTiming calls the method "AnimationEffect.getComputedTiming"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryGetComputedTiming() (ret ComputedEffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUpdateTiming returns true if the method "AnimationEffect.updateTiming" exists.
func (this AnimationEffect) HasUpdateTiming() bool {
	return js.True == bindings.HasAnimationEffectUpdateTiming(
		this.Ref(),
	)
}

// UpdateTimingFunc returns the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTimingFunc() (fn js.Func[func(timing OptionalEffectTiming)]) {
	return fn.FromRef(
		bindings.AnimationEffectUpdateTimingFunc(
			this.Ref(),
		),
	)
}

// UpdateTiming calls the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTiming(timing OptionalEffectTiming) (ret js.Void) {
	bindings.CallAnimationEffectUpdateTiming(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&timing),
	)

	return
}

// TryUpdateTiming calls the method "AnimationEffect.updateTiming"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryUpdateTiming(timing OptionalEffectTiming) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectUpdateTiming(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&timing),
	)

	return
}

// HasUpdateTiming1 returns true if the method "AnimationEffect.updateTiming" exists.
func (this AnimationEffect) HasUpdateTiming1() bool {
	return js.True == bindings.HasAnimationEffectUpdateTiming1(
		this.Ref(),
	)
}

// UpdateTiming1Func returns the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTiming1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationEffectUpdateTiming1Func(
			this.Ref(),
		),
	)
}

// UpdateTiming1 calls the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTiming1() (ret js.Void) {
	bindings.CallAnimationEffectUpdateTiming1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUpdateTiming1 calls the method "AnimationEffect.updateTiming"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryUpdateTiming1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectUpdateTiming1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBefore returns true if the method "AnimationEffect.before" exists.
func (this AnimationEffect) HasBefore() bool {
	return js.True == bindings.HasAnimationEffectBefore(
		this.Ref(),
	)
}

// BeforeFunc returns the method "AnimationEffect.before".
func (this AnimationEffect) BeforeFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectBeforeFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "AnimationEffect.before".
func (this AnimationEffect) Before(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectBefore(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryBefore calls the method "AnimationEffect.before"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryBefore(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasAfter returns true if the method "AnimationEffect.after" exists.
func (this AnimationEffect) HasAfter() bool {
	return js.True == bindings.HasAnimationEffectAfter(
		this.Ref(),
	)
}

// AfterFunc returns the method "AnimationEffect.after".
func (this AnimationEffect) AfterFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectAfterFunc(
			this.Ref(),
		),
	)
}

// After calls the method "AnimationEffect.after".
func (this AnimationEffect) After(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectAfter(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryAfter calls the method "AnimationEffect.after"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryAfter(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasReplace returns true if the method "AnimationEffect.replace" exists.
func (this AnimationEffect) HasReplace() bool {
	return js.True == bindings.HasAnimationEffectReplace(
		this.Ref(),
	)
}

// ReplaceFunc returns the method "AnimationEffect.replace".
func (this AnimationEffect) ReplaceFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectReplaceFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "AnimationEffect.replace".
func (this AnimationEffect) Replace(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectReplace(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryReplace calls the method "AnimationEffect.replace"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryReplace(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectReplace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasRemove returns true if the method "AnimationEffect.remove" exists.
func (this AnimationEffect) HasRemove() bool {
	return js.True == bindings.HasAnimationEffectRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "AnimationEffect.remove".
func (this AnimationEffect) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationEffectRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "AnimationEffect.remove".
func (this AnimationEffect) Remove() (ret js.Void) {
	bindings.CallAnimationEffectRemove(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "AnimationEffect.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationEffect) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AnimationTimeline struct {
	ref js.Ref
}

func (this AnimationTimeline) Once() AnimationTimeline {
	this.Ref().Once()
	return this
}

func (this AnimationTimeline) Ref() js.Ref {
	return this.ref
}

func (this AnimationTimeline) FromRef(ref js.Ref) AnimationTimeline {
	this.ref = ref
	return this
}

func (this AnimationTimeline) Free() {
	this.Ref().Free()
}

// CurrentTime returns the value of property "AnimationTimeline.currentTime".
//
// It returns ok=false if there is no such property.
func (this AnimationTimeline) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationTimelineCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "AnimationTimeline.duration".
//
// It returns ok=false if there is no such property.
func (this AnimationTimeline) Duration() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationTimelineDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPlay returns true if the method "AnimationTimeline.play" exists.
func (this AnimationTimeline) HasPlay() bool {
	return js.True == bindings.HasAnimationTimelinePlay(
		this.Ref(),
	)
}

// PlayFunc returns the method "AnimationTimeline.play".
func (this AnimationTimeline) PlayFunc() (fn js.Func[func(effect AnimationEffect) Animation]) {
	return fn.FromRef(
		bindings.AnimationTimelinePlayFunc(
			this.Ref(),
		),
	)
}

// Play calls the method "AnimationTimeline.play".
func (this AnimationTimeline) Play(effect AnimationEffect) (ret Animation) {
	bindings.CallAnimationTimelinePlay(
		this.Ref(), js.Pointer(&ret),
		effect.Ref(),
	)

	return
}

// TryPlay calls the method "AnimationTimeline.play"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationTimeline) TryPlay(effect AnimationEffect) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationTimelinePlay(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		effect.Ref(),
	)

	return
}

// HasPlay1 returns true if the method "AnimationTimeline.play" exists.
func (this AnimationTimeline) HasPlay1() bool {
	return js.True == bindings.HasAnimationTimelinePlay1(
		this.Ref(),
	)
}

// Play1Func returns the method "AnimationTimeline.play".
func (this AnimationTimeline) Play1Func() (fn js.Func[func() Animation]) {
	return fn.FromRef(
		bindings.AnimationTimelinePlay1Func(
			this.Ref(),
		),
	)
}

// Play1 calls the method "AnimationTimeline.play".
func (this AnimationTimeline) Play1() (ret Animation) {
	bindings.CallAnimationTimelinePlay1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPlay1 calls the method "AnimationTimeline.play"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationTimeline) TryPlay1() (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationTimelinePlay1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AnimationPlayState uint32

const (
	_ AnimationPlayState = iota

	AnimationPlayState_IDLE
	AnimationPlayState_RUNNING
	AnimationPlayState_PAUSED
	AnimationPlayState_FINISHED
)

func (AnimationPlayState) FromRef(str js.Ref) AnimationPlayState {
	return AnimationPlayState(bindings.ConstOfAnimationPlayState(str))
}

func (x AnimationPlayState) String() (string, bool) {
	switch x {
	case AnimationPlayState_IDLE:
		return "idle", true
	case AnimationPlayState_RUNNING:
		return "running", true
	case AnimationPlayState_PAUSED:
		return "paused", true
	case AnimationPlayState_FINISHED:
		return "finished", true
	default:
		return "", false
	}
}

type AnimationReplaceState uint32

const (
	_ AnimationReplaceState = iota

	AnimationReplaceState_ACTIVE
	AnimationReplaceState_REMOVED
	AnimationReplaceState_PERSISTED
)

func (AnimationReplaceState) FromRef(str js.Ref) AnimationReplaceState {
	return AnimationReplaceState(bindings.ConstOfAnimationReplaceState(str))
}

func (x AnimationReplaceState) String() (string, bool) {
	switch x {
	case AnimationReplaceState_ACTIVE:
		return "active", true
	case AnimationReplaceState_REMOVED:
		return "removed", true
	case AnimationReplaceState_PERSISTED:
		return "persisted", true
	default:
		return "", false
	}
}

func NewAnimation(effect AnimationEffect, timeline AnimationTimeline) (ret Animation) {
	ret.ref = bindings.NewAnimationByAnimation(
		effect.Ref(),
		timeline.Ref())
	return
}

func NewAnimationByAnimation1(effect AnimationEffect) (ret Animation) {
	ret.ref = bindings.NewAnimationByAnimation1(
		effect.Ref())
	return
}

func NewAnimationByAnimation2() (ret Animation) {
	ret.ref = bindings.NewAnimationByAnimation2()
	return
}

type Animation struct {
	EventTarget
}

func (this Animation) Once() Animation {
	this.Ref().Once()
	return this
}

func (this Animation) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Animation) FromRef(ref js.Ref) Animation {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Animation) Free() {
	this.Ref().Free()
}

// Id returns the value of property "Animation.id".
//
// It returns ok=false if there is no such property.
func (this Animation) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "Animation.id" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetId(val js.String) bool {
	return js.True == bindings.SetAnimationId(
		this.Ref(),
		val.Ref(),
	)
}

// Effect returns the value of property "Animation.effect".
//
// It returns ok=false if there is no such property.
func (this Animation) Effect() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEffect sets the value of property "Animation.effect" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetEffect(val AnimationEffect) bool {
	return js.True == bindings.SetAnimationEffect(
		this.Ref(),
		val.Ref(),
	)
}

// Timeline returns the value of property "Animation.timeline".
//
// It returns ok=false if there is no such property.
func (this Animation) Timeline() (ret AnimationTimeline, ok bool) {
	ok = js.True == bindings.GetAnimationTimeline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTimeline sets the value of property "Animation.timeline" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetTimeline(val AnimationTimeline) bool {
	return js.True == bindings.SetAnimationTimeline(
		this.Ref(),
		val.Ref(),
	)
}

// PlaybackRate returns the value of property "Animation.playbackRate".
//
// It returns ok=false if there is no such property.
func (this Animation) PlaybackRate() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaybackRate sets the value of property "Animation.playbackRate" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetPlaybackRate(val float64) bool {
	return js.True == bindings.SetAnimationPlaybackRate(
		this.Ref(),
		float64(val),
	)
}

// PlayState returns the value of property "Animation.playState".
//
// It returns ok=false if there is no such property.
func (this Animation) PlayState() (ret AnimationPlayState, ok bool) {
	ok = js.True == bindings.GetAnimationPlayState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReplaceState returns the value of property "Animation.replaceState".
//
// It returns ok=false if there is no such property.
func (this Animation) ReplaceState() (ret AnimationReplaceState, ok bool) {
	ok = js.True == bindings.GetAnimationReplaceState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Pending returns the value of property "Animation.pending".
//
// It returns ok=false if there is no such property.
func (this Animation) Pending() (ret bool, ok bool) {
	ok = js.True == bindings.GetAnimationPending(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "Animation.ready".
//
// It returns ok=false if there is no such property.
func (this Animation) Ready() (ret js.Promise[Animation], ok bool) {
	ok = js.True == bindings.GetAnimationReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "Animation.finished".
//
// It returns ok=false if there is no such property.
func (this Animation) Finished() (ret js.Promise[Animation], ok bool) {
	ok = js.True == bindings.GetAnimationFinished(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StartTime returns the value of property "Animation.startTime".
//
// It returns ok=false if there is no such property.
func (this Animation) StartTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationStartTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStartTime sets the value of property "Animation.startTime" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetStartTime(val CSSNumberish) bool {
	return js.True == bindings.SetAnimationStartTime(
		this.Ref(),
		val.Ref(),
	)
}

// CurrentTime returns the value of property "Animation.currentTime".
//
// It returns ok=false if there is no such property.
func (this Animation) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCurrentTime sets the value of property "Animation.currentTime" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetCurrentTime(val CSSNumberish) bool {
	return js.True == bindings.SetAnimationCurrentTime(
		this.Ref(),
		val.Ref(),
	)
}

// HasCancel returns true if the method "Animation.cancel" exists.
func (this Animation) HasCancel() bool {
	return js.True == bindings.HasAnimationCancel(
		this.Ref(),
	)
}

// CancelFunc returns the method "Animation.cancel".
func (this Animation) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "Animation.cancel".
func (this Animation) Cancel() (ret js.Void) {
	bindings.CallAnimationCancel(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "Animation.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationCancel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFinish returns true if the method "Animation.finish" exists.
func (this Animation) HasFinish() bool {
	return js.True == bindings.HasAnimationFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "Animation.finish".
func (this Animation) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "Animation.finish".
func (this Animation) Finish() (ret js.Void) {
	bindings.CallAnimationFinish(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "Animation.finish"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPlay returns true if the method "Animation.play" exists.
func (this Animation) HasPlay() bool {
	return js.True == bindings.HasAnimationPlay(
		this.Ref(),
	)
}

// PlayFunc returns the method "Animation.play".
func (this Animation) PlayFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPlayFunc(
			this.Ref(),
		),
	)
}

// Play calls the method "Animation.play".
func (this Animation) Play() (ret js.Void) {
	bindings.CallAnimationPlay(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPlay calls the method "Animation.play"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryPlay() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPlay(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPause returns true if the method "Animation.pause" exists.
func (this Animation) HasPause() bool {
	return js.True == bindings.HasAnimationPause(
		this.Ref(),
	)
}

// PauseFunc returns the method "Animation.pause".
func (this Animation) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPauseFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "Animation.pause".
func (this Animation) Pause() (ret js.Void) {
	bindings.CallAnimationPause(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "Animation.pause"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPause(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUpdatePlaybackRate returns true if the method "Animation.updatePlaybackRate" exists.
func (this Animation) HasUpdatePlaybackRate() bool {
	return js.True == bindings.HasAnimationUpdatePlaybackRate(
		this.Ref(),
	)
}

// UpdatePlaybackRateFunc returns the method "Animation.updatePlaybackRate".
func (this Animation) UpdatePlaybackRateFunc() (fn js.Func[func(playbackRate float64)]) {
	return fn.FromRef(
		bindings.AnimationUpdatePlaybackRateFunc(
			this.Ref(),
		),
	)
}

// UpdatePlaybackRate calls the method "Animation.updatePlaybackRate".
func (this Animation) UpdatePlaybackRate(playbackRate float64) (ret js.Void) {
	bindings.CallAnimationUpdatePlaybackRate(
		this.Ref(), js.Pointer(&ret),
		float64(playbackRate),
	)

	return
}

// TryUpdatePlaybackRate calls the method "Animation.updatePlaybackRate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryUpdatePlaybackRate(playbackRate float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationUpdatePlaybackRate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(playbackRate),
	)

	return
}

// HasReverse returns true if the method "Animation.reverse" exists.
func (this Animation) HasReverse() bool {
	return js.True == bindings.HasAnimationReverse(
		this.Ref(),
	)
}

// ReverseFunc returns the method "Animation.reverse".
func (this Animation) ReverseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationReverseFunc(
			this.Ref(),
		),
	)
}

// Reverse calls the method "Animation.reverse".
func (this Animation) Reverse() (ret js.Void) {
	bindings.CallAnimationReverse(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReverse calls the method "Animation.reverse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryReverse() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationReverse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPersist returns true if the method "Animation.persist" exists.
func (this Animation) HasPersist() bool {
	return js.True == bindings.HasAnimationPersist(
		this.Ref(),
	)
}

// PersistFunc returns the method "Animation.persist".
func (this Animation) PersistFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPersistFunc(
			this.Ref(),
		),
	)
}

// Persist calls the method "Animation.persist".
func (this Animation) Persist() (ret js.Void) {
	bindings.CallAnimationPersist(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "Animation.persist"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryPersist() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPersist(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCommitStyles returns true if the method "Animation.commitStyles" exists.
func (this Animation) HasCommitStyles() bool {
	return js.True == bindings.HasAnimationCommitStyles(
		this.Ref(),
	)
}

// CommitStylesFunc returns the method "Animation.commitStyles".
func (this Animation) CommitStylesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationCommitStylesFunc(
			this.Ref(),
		),
	)
}

// CommitStyles calls the method "Animation.commitStyles".
func (this Animation) CommitStyles() (ret js.Void) {
	bindings.CallAnimationCommitStyles(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCommitStyles calls the method "Animation.commitStyles"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Animation) TryCommitStyles() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationCommitStyles(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ShadowRootMode uint32

const (
	_ ShadowRootMode = iota

	ShadowRootMode_OPEN
	ShadowRootMode_CLOSED
)

func (ShadowRootMode) FromRef(str js.Ref) ShadowRootMode {
	return ShadowRootMode(bindings.ConstOfShadowRootMode(str))
}

func (x ShadowRootMode) String() (string, bool) {
	switch x {
	case ShadowRootMode_OPEN:
		return "open", true
	case ShadowRootMode_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

type SlotAssignmentMode uint32

const (
	_ SlotAssignmentMode = iota

	SlotAssignmentMode_MANUAL
	SlotAssignmentMode_NAMED
)

func (SlotAssignmentMode) FromRef(str js.Ref) SlotAssignmentMode {
	return SlotAssignmentMode(bindings.ConstOfSlotAssignmentMode(str))
}

func (x SlotAssignmentMode) String() (string, bool) {
	switch x {
	case SlotAssignmentMode_MANUAL:
		return "manual", true
	case SlotAssignmentMode_NAMED:
		return "named", true
	default:
		return "", false
	}
}

type MediaList struct {
	ref js.Ref
}

func (this MediaList) Once() MediaList {
	this.Ref().Once()
	return this
}

func (this MediaList) Ref() js.Ref {
	return this.ref
}

func (this MediaList) FromRef(ref js.Ref) MediaList {
	this.ref = ref
	return this
}

func (this MediaList) Free() {
	this.Ref().Free()
}

// MediaText returns the value of property "MediaList.mediaText".
//
// It returns ok=false if there is no such property.
func (this MediaList) MediaText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaListMediaText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMediaText sets the value of property "MediaList.mediaText" to val.
//
// It returns false if the property cannot be set.
func (this MediaList) SetMediaText(val js.String) bool {
	return js.True == bindings.SetMediaListMediaText(
		this.Ref(),
		val.Ref(),
	)
}

// Length returns the value of property "MediaList.length".
//
// It returns ok=false if there is no such property.
func (this MediaList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "MediaList.item" exists.
func (this MediaList) HasItem() bool {
	return js.True == bindings.HasMediaListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "MediaList.item".
func (this MediaList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.MediaListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "MediaList.item".
func (this MediaList) Item(index uint32) (ret js.String) {
	bindings.CallMediaListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "MediaList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendMedium returns true if the method "MediaList.appendMedium" exists.
func (this MediaList) HasAppendMedium() bool {
	return js.True == bindings.HasMediaListAppendMedium(
		this.Ref(),
	)
}

// AppendMediumFunc returns the method "MediaList.appendMedium".
func (this MediaList) AppendMediumFunc() (fn js.Func[func(medium js.String)]) {
	return fn.FromRef(
		bindings.MediaListAppendMediumFunc(
			this.Ref(),
		),
	)
}

// AppendMedium calls the method "MediaList.appendMedium".
func (this MediaList) AppendMedium(medium js.String) (ret js.Void) {
	bindings.CallMediaListAppendMedium(
		this.Ref(), js.Pointer(&ret),
		medium.Ref(),
	)

	return
}

// TryAppendMedium calls the method "MediaList.appendMedium"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaList) TryAppendMedium(medium js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListAppendMedium(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		medium.Ref(),
	)

	return
}

// HasDeleteMedium returns true if the method "MediaList.deleteMedium" exists.
func (this MediaList) HasDeleteMedium() bool {
	return js.True == bindings.HasMediaListDeleteMedium(
		this.Ref(),
	)
}

// DeleteMediumFunc returns the method "MediaList.deleteMedium".
func (this MediaList) DeleteMediumFunc() (fn js.Func[func(medium js.String)]) {
	return fn.FromRef(
		bindings.MediaListDeleteMediumFunc(
			this.Ref(),
		),
	)
}

// DeleteMedium calls the method "MediaList.deleteMedium".
func (this MediaList) DeleteMedium(medium js.String) (ret js.Void) {
	bindings.CallMediaListDeleteMedium(
		this.Ref(), js.Pointer(&ret),
		medium.Ref(),
	)

	return
}

// TryDeleteMedium calls the method "MediaList.deleteMedium"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaList) TryDeleteMedium(medium js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListDeleteMedium(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		medium.Ref(),
	)

	return
}

type OneOf_MediaList_String struct {
	ref js.Ref
}

func (x OneOf_MediaList_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_MediaList_String) Free() {
	x.ref.Free()
}

func (x OneOf_MediaList_String) FromRef(ref js.Ref) OneOf_MediaList_String {
	return OneOf_MediaList_String{
		ref: ref,
	}
}

func (x OneOf_MediaList_String) MediaList() MediaList {
	return MediaList{}.FromRef(x.ref)
}

func (x OneOf_MediaList_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type CSSStyleSheetInit struct {
	// BaseURL is "CSSStyleSheetInit.baseURL"
	//
	// Optional, defaults to null.
	BaseURL js.String
	// Media is "CSSStyleSheetInit.media"
	//
	// Optional, defaults to "".
	Media OneOf_MediaList_String
	// Disabled is "CSSStyleSheetInit.disabled"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Disabled MUST be set to true to make this field effective.
	Disabled bool

	FFI_USE_Disabled bool // for Disabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CSSStyleSheetInit with all fields set.
func (p CSSStyleSheetInit) FromRef(ref js.Ref) CSSStyleSheetInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CSSStyleSheetInit in the application heap.
func (p CSSStyleSheetInit) New() js.Ref {
	return bindings.CSSStyleSheetInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CSSStyleSheetInit) UpdateFrom(ref js.Ref) {
	bindings.CSSStyleSheetInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CSSStyleSheetInit) Update(ref js.Ref) {
	bindings.CSSStyleSheetInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

const (
	CSSRule_STYLE_RULE               uint16 = 1
	CSSRule_CHARSET_RULE             uint16 = 2
	CSSRule_IMPORT_RULE              uint16 = 3
	CSSRule_MEDIA_RULE               uint16 = 4
	CSSRule_FONT_FACE_RULE           uint16 = 5
	CSSRule_PAGE_RULE                uint16 = 6
	CSSRule_MARGIN_RULE              uint16 = 9
	CSSRule_NAMESPACE_RULE           uint16 = 10
	CSSRule_COUNTER_STYLE_RULE       uint16 = 11
	CSSRule_SUPPORTS_RULE            uint16 = 12
	CSSRule_FONT_FEATURE_VALUES_RULE uint16 = 14
	CSSRule_KEYFRAMES_RULE           uint16 = 7
	CSSRule_KEYFRAME_RULE            uint16 = 8
)

type CSSRule struct {
	ref js.Ref
}

func (this CSSRule) Once() CSSRule {
	this.Ref().Once()
	return this
}

func (this CSSRule) Ref() js.Ref {
	return this.ref
}

func (this CSSRule) FromRef(ref js.Ref) CSSRule {
	this.ref = ref
	return this
}

func (this CSSRule) Free() {
	this.Ref().Free()
}

// CssText returns the value of property "CSSRule.cssText".
//
// It returns ok=false if there is no such property.
func (this CSSRule) CssText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSRuleCssText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCssText sets the value of property "CSSRule.cssText" to val.
//
// It returns false if the property cannot be set.
func (this CSSRule) SetCssText(val js.String) bool {
	return js.True == bindings.SetCSSRuleCssText(
		this.Ref(),
		val.Ref(),
	)
}

// ParentRule returns the value of property "CSSRule.parentRule".
//
// It returns ok=false if there is no such property.
func (this CSSRule) ParentRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSRuleParentRule(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ParentStyleSheet returns the value of property "CSSRule.parentStyleSheet".
//
// It returns ok=false if there is no such property.
func (this CSSRule) ParentStyleSheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetCSSRuleParentStyleSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "CSSRule.type".
//
// It returns ok=false if there is no such property.
func (this CSSRule) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCSSRuleType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CSSRuleList struct {
	ref js.Ref
}

func (this CSSRuleList) Once() CSSRuleList {
	this.Ref().Once()
	return this
}

func (this CSSRuleList) Ref() js.Ref {
	return this.ref
}

func (this CSSRuleList) FromRef(ref js.Ref) CSSRuleList {
	this.ref = ref
	return this
}

func (this CSSRuleList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "CSSRuleList.length".
//
// It returns ok=false if there is no such property.
func (this CSSRuleList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSRuleListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "CSSRuleList.item" exists.
func (this CSSRuleList) HasItem() bool {
	return js.True == bindings.HasCSSRuleListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "CSSRuleList.item".
func (this CSSRuleList) ItemFunc() (fn js.Func[func(index uint32) CSSRule]) {
	return fn.FromRef(
		bindings.CSSRuleListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "CSSRuleList.item".
func (this CSSRuleList) Item(index uint32) (ret CSSRule) {
	bindings.CallCSSRuleListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "CSSRuleList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSRuleList) TryItem(index uint32) (ret CSSRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRuleListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewCSSStyleSheet(options CSSStyleSheetInit) (ret CSSStyleSheet) {
	ret.ref = bindings.NewCSSStyleSheetByCSSStyleSheet(
		js.Pointer(&options))
	return
}

func NewCSSStyleSheetByCSSStyleSheet1() (ret CSSStyleSheet) {
	ret.ref = bindings.NewCSSStyleSheetByCSSStyleSheet1()
	return
}

type CSSStyleSheet struct {
	StyleSheet
}

func (this CSSStyleSheet) Once() CSSStyleSheet {
	this.Ref().Once()
	return this
}

func (this CSSStyleSheet) Ref() js.Ref {
	return this.StyleSheet.Ref()
}

func (this CSSStyleSheet) FromRef(ref js.Ref) CSSStyleSheet {
	this.StyleSheet = this.StyleSheet.FromRef(ref)
	return this
}

func (this CSSStyleSheet) Free() {
	this.Ref().Free()
}

// OwnerRule returns the value of property "CSSStyleSheet.ownerRule".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) OwnerRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetOwnerRule(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CssRules returns the value of property "CSSStyleSheet.cssRules".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetCssRules(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rules returns the value of property "CSSStyleSheet.rules".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) Rules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetRules(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInsertRule returns true if the method "CSSStyleSheet.insertRule" exists.
func (this CSSStyleSheet) HasInsertRule() bool {
	return js.True == bindings.HasCSSStyleSheetInsertRule(
		this.Ref(),
	)
}

// InsertRuleFunc returns the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRuleFunc() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetInsertRuleFunc(
			this.Ref(),
		),
	)
}

// InsertRule calls the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRule(rule js.String, index uint32) (ret uint32) {
	bindings.CallCSSStyleSheetInsertRule(
		this.Ref(), js.Pointer(&ret),
		rule.Ref(),
		uint32(index),
	)

	return
}

// TryInsertRule calls the method "CSSStyleSheet.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryInsertRule(rule js.String, index uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetInsertRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
		uint32(index),
	)

	return
}

// HasInsertRule1 returns true if the method "CSSStyleSheet.insertRule" exists.
func (this CSSStyleSheet) HasInsertRule1() bool {
	return js.True == bindings.HasCSSStyleSheetInsertRule1(
		this.Ref(),
	)
}

// InsertRule1Func returns the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRule1Func() (fn js.Func[func(rule js.String) uint32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetInsertRule1Func(
			this.Ref(),
		),
	)
}

// InsertRule1 calls the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRule1(rule js.String) (ret uint32) {
	bindings.CallCSSStyleSheetInsertRule1(
		this.Ref(), js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryInsertRule1 calls the method "CSSStyleSheet.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryInsertRule1(rule js.String) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetInsertRule1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasDeleteRule returns true if the method "CSSStyleSheet.deleteRule" exists.
func (this CSSStyleSheet) HasDeleteRule() bool {
	return js.True == bindings.HasCSSStyleSheetDeleteRule(
		this.Ref(),
	)
}

// DeleteRuleFunc returns the method "CSSStyleSheet.deleteRule".
func (this CSSStyleSheet) DeleteRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetDeleteRuleFunc(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSStyleSheet.deleteRule".
func (this CSSStyleSheet) DeleteRule(index uint32) (ret js.Void) {
	bindings.CallCSSStyleSheetDeleteRule(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDeleteRule calls the method "CSSStyleSheet.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryDeleteRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetDeleteRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasReplace returns true if the method "CSSStyleSheet.replace" exists.
func (this CSSStyleSheet) HasReplace() bool {
	return js.True == bindings.HasCSSStyleSheetReplace(
		this.Ref(),
	)
}

// ReplaceFunc returns the method "CSSStyleSheet.replace".
func (this CSSStyleSheet) ReplaceFunc() (fn js.Func[func(text js.String) js.Promise[CSSStyleSheet]]) {
	return fn.FromRef(
		bindings.CSSStyleSheetReplaceFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "CSSStyleSheet.replace".
func (this CSSStyleSheet) Replace(text js.String) (ret js.Promise[CSSStyleSheet]) {
	bindings.CallCSSStyleSheetReplace(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryReplace calls the method "CSSStyleSheet.replace"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryReplace(text js.String) (ret js.Promise[CSSStyleSheet], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetReplace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasReplaceSync returns true if the method "CSSStyleSheet.replaceSync" exists.
func (this CSSStyleSheet) HasReplaceSync() bool {
	return js.True == bindings.HasCSSStyleSheetReplaceSync(
		this.Ref(),
	)
}

// ReplaceSyncFunc returns the method "CSSStyleSheet.replaceSync".
func (this CSSStyleSheet) ReplaceSyncFunc() (fn js.Func[func(text js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetReplaceSyncFunc(
			this.Ref(),
		),
	)
}

// ReplaceSync calls the method "CSSStyleSheet.replaceSync".
func (this CSSStyleSheet) ReplaceSync(text js.String) (ret js.Void) {
	bindings.CallCSSStyleSheetReplaceSync(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryReplaceSync calls the method "CSSStyleSheet.replaceSync"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryReplaceSync(text js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetReplaceSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasAddRule returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasAddRule() bool {
	return js.True == bindings.HasCSSStyleSheetAddRule(
		this.Ref(),
	)
}

// AddRuleFunc returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRuleFunc() (fn js.Func[func(selector js.String, style js.String, index uint32) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRuleFunc(
			this.Ref(),
		),
	)
}

// AddRule calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule(selector js.String, style js.String, index uint32) (ret int32) {
	bindings.CallCSSStyleSheetAddRule(
		this.Ref(), js.Pointer(&ret),
		selector.Ref(),
		style.Ref(),
		uint32(index),
	)

	return
}

// TryAddRule calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryAddRule(selector js.String, style js.String, index uint32) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
		style.Ref(),
		uint32(index),
	)

	return
}

// HasAddRule1 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasAddRule1() bool {
	return js.True == bindings.HasCSSStyleSheetAddRule1(
		this.Ref(),
	)
}

// AddRule1Func returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule1Func() (fn js.Func[func(selector js.String, style js.String) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule1Func(
			this.Ref(),
		),
	)
}

// AddRule1 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule1(selector js.String, style js.String) (ret int32) {
	bindings.CallCSSStyleSheetAddRule1(
		this.Ref(), js.Pointer(&ret),
		selector.Ref(),
		style.Ref(),
	)

	return
}

// TryAddRule1 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryAddRule1(selector js.String, style js.String) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
		style.Ref(),
	)

	return
}

// HasAddRule2 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasAddRule2() bool {
	return js.True == bindings.HasCSSStyleSheetAddRule2(
		this.Ref(),
	)
}

// AddRule2Func returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule2Func() (fn js.Func[func(selector js.String) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule2Func(
			this.Ref(),
		),
	)
}

// AddRule2 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule2(selector js.String) (ret int32) {
	bindings.CallCSSStyleSheetAddRule2(
		this.Ref(), js.Pointer(&ret),
		selector.Ref(),
	)

	return
}

// TryAddRule2 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryAddRule2(selector js.String) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
	)

	return
}

// HasAddRule3 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasAddRule3() bool {
	return js.True == bindings.HasCSSStyleSheetAddRule3(
		this.Ref(),
	)
}

// AddRule3Func returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule3Func() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule3Func(
			this.Ref(),
		),
	)
}

// AddRule3 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule3() (ret int32) {
	bindings.CallCSSStyleSheetAddRule3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAddRule3 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryAddRule3() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRemoveRule returns true if the method "CSSStyleSheet.removeRule" exists.
func (this CSSStyleSheet) HasRemoveRule() bool {
	return js.True == bindings.HasCSSStyleSheetRemoveRule(
		this.Ref(),
	)
}

// RemoveRuleFunc returns the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetRemoveRuleFunc(
			this.Ref(),
		),
	)
}

// RemoveRule calls the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRule(index uint32) (ret js.Void) {
	bindings.CallCSSStyleSheetRemoveRule(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveRule calls the method "CSSStyleSheet.removeRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryRemoveRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetRemoveRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasRemoveRule1 returns true if the method "CSSStyleSheet.removeRule" exists.
func (this CSSStyleSheet) HasRemoveRule1() bool {
	return js.True == bindings.HasCSSStyleSheetRemoveRule1(
		this.Ref(),
	)
}

// RemoveRule1Func returns the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRule1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CSSStyleSheetRemoveRule1Func(
			this.Ref(),
		),
	)
}

// RemoveRule1 calls the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRule1() (ret js.Void) {
	bindings.CallCSSStyleSheetRemoveRule1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemoveRule1 calls the method "CSSStyleSheet.removeRule"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleSheet) TryRemoveRule1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetRemoveRule1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type StyleSheetList struct {
	ref js.Ref
}

func (this StyleSheetList) Once() StyleSheetList {
	this.Ref().Once()
	return this
}

func (this StyleSheetList) Ref() js.Ref {
	return this.ref
}

func (this StyleSheetList) FromRef(ref js.Ref) StyleSheetList {
	this.ref = ref
	return this
}

func (this StyleSheetList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "StyleSheetList.length".
//
// It returns ok=false if there is no such property.
func (this StyleSheetList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStyleSheetListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "StyleSheetList.item" exists.
func (this StyleSheetList) HasItem() bool {
	return js.True == bindings.HasStyleSheetListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "StyleSheetList.item".
func (this StyleSheetList) ItemFunc() (fn js.Func[func(index uint32) CSSStyleSheet]) {
	return fn.FromRef(
		bindings.StyleSheetListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "StyleSheetList.item".
func (this StyleSheetList) Item(index uint32) (ret CSSStyleSheet) {
	bindings.CallStyleSheetListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "StyleSheetList.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StyleSheetList) TryItem(index uint32) (ret CSSStyleSheet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStyleSheetListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type ShadowRoot struct {
	DocumentFragment
}

func (this ShadowRoot) Once() ShadowRoot {
	this.Ref().Once()
	return this
}

func (this ShadowRoot) Ref() js.Ref {
	return this.DocumentFragment.Ref()
}

func (this ShadowRoot) FromRef(ref js.Ref) ShadowRoot {
	this.DocumentFragment = this.DocumentFragment.FromRef(ref)
	return this
}

func (this ShadowRoot) Free() {
	this.Ref().Free()
}

// Mode returns the value of property "ShadowRoot.mode".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) Mode() (ret ShadowRootMode, ok bool) {
	ok = js.True == bindings.GetShadowRootMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DelegatesFocus returns the value of property "ShadowRoot.delegatesFocus".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) DelegatesFocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetShadowRootDelegatesFocus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SlotAssignment returns the value of property "ShadowRoot.slotAssignment".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) SlotAssignment() (ret SlotAssignmentMode, ok bool) {
	ok = js.True == bindings.GetShadowRootSlotAssignment(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Host returns the value of property "ShadowRoot.host".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) Host() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootHost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InnerHTML returns the value of property "ShadowRoot.innerHTML".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) InnerHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetShadowRootInnerHTML(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInnerHTML sets the value of property "ShadowRoot.innerHTML" to val.
//
// It returns false if the property cannot be set.
func (this ShadowRoot) SetInnerHTML(val js.String) bool {
	return js.True == bindings.SetShadowRootInnerHTML(
		this.Ref(),
		val.Ref(),
	)
}

// FullscreenElement returns the value of property "ShadowRoot.fullscreenElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) FullscreenElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootFullscreenElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StyleSheets returns the value of property "ShadowRoot.styleSheets".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) StyleSheets() (ret StyleSheetList, ok bool) {
	ok = js.True == bindings.GetShadowRootStyleSheets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AdoptedStyleSheets returns the value of property "ShadowRoot.adoptedStyleSheets".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) AdoptedStyleSheets() (ret js.ObservableArray[CSSStyleSheet], ok bool) {
	ok = js.True == bindings.GetShadowRootAdoptedStyleSheets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAdoptedStyleSheets sets the value of property "ShadowRoot.adoptedStyleSheets" to val.
//
// It returns false if the property cannot be set.
func (this ShadowRoot) SetAdoptedStyleSheets(val js.ObservableArray[CSSStyleSheet]) bool {
	return js.True == bindings.SetShadowRootAdoptedStyleSheets(
		this.Ref(),
		val.Ref(),
	)
}

// PictureInPictureElement returns the value of property "ShadowRoot.pictureInPictureElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) PictureInPictureElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootPictureInPictureElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActiveElement returns the value of property "ShadowRoot.activeElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) ActiveElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootActiveElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointerLockElement returns the value of property "ShadowRoot.pointerLockElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) PointerLockElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootPointerLockElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetAnimations returns true if the method "ShadowRoot.getAnimations" exists.
func (this ShadowRoot) HasGetAnimations() bool {
	return js.True == bindings.HasShadowRootGetAnimations(
		this.Ref(),
	)
}

// GetAnimationsFunc returns the method "ShadowRoot.getAnimations".
func (this ShadowRoot) GetAnimationsFunc() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ShadowRootGetAnimationsFunc(
			this.Ref(),
		),
	)
}

// GetAnimations calls the method "ShadowRoot.getAnimations".
func (this ShadowRoot) GetAnimations() (ret js.Array[Animation]) {
	bindings.CallShadowRootGetAnimations(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAnimations calls the method "ShadowRoot.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ShadowRoot) TryGetAnimations() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShadowRootGetAnimations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ShadowRootInit struct {
	// Mode is "ShadowRootInit.mode"
	//
	// Required
	Mode ShadowRootMode
	// DelegatesFocus is "ShadowRootInit.delegatesFocus"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_DelegatesFocus MUST be set to true to make this field effective.
	DelegatesFocus bool
	// SlotAssignment is "ShadowRootInit.slotAssignment"
	//
	// Optional, defaults to "named".
	SlotAssignment SlotAssignmentMode

	FFI_USE_DelegatesFocus bool // for DelegatesFocus.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShadowRootInit with all fields set.
func (p ShadowRootInit) FromRef(ref js.Ref) ShadowRootInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShadowRootInit in the application heap.
func (p ShadowRootInit) New() js.Ref {
	return bindings.ShadowRootInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ShadowRootInit) UpdateFrom(ref js.Ref) {
	bindings.ShadowRootInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ShadowRootInit) Update(ref js.Ref) {
	bindings.ShadowRootInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FullscreenNavigationUI uint32

const (
	_ FullscreenNavigationUI = iota

	FullscreenNavigationUI_AUTO
	FullscreenNavigationUI_SHOW
	FullscreenNavigationUI_HIDE
)

func (FullscreenNavigationUI) FromRef(str js.Ref) FullscreenNavigationUI {
	return FullscreenNavigationUI(bindings.ConstOfFullscreenNavigationUI(str))
}

func (x FullscreenNavigationUI) String() (string, bool) {
	switch x {
	case FullscreenNavigationUI_AUTO:
		return "auto", true
	case FullscreenNavigationUI_SHOW:
		return "show", true
	case FullscreenNavigationUI_HIDE:
		return "hide", true
	default:
		return "", false
	}
}

type ScreenDetailed struct {
	Screen
}

func (this ScreenDetailed) Once() ScreenDetailed {
	this.Ref().Once()
	return this
}

func (this ScreenDetailed) Ref() js.Ref {
	return this.Screen.Ref()
}

func (this ScreenDetailed) FromRef(ref js.Ref) ScreenDetailed {
	this.Screen = this.Screen.FromRef(ref)
	return this
}

func (this ScreenDetailed) Free() {
	this.Ref().Free()
}

// AvailLeft returns the value of property "ScreenDetailed.availLeft".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) AvailLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedAvailLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AvailTop returns the value of property "ScreenDetailed.availTop".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) AvailTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedAvailTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Left returns the value of property "ScreenDetailed.left".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Left() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "ScreenDetailed.top".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Top() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "ScreenDetailed.isPrimary".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenDetailedIsPrimary(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsInternal returns the value of property "ScreenDetailed.isInternal".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) IsInternal() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenDetailedIsInternal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DevicePixelRatio returns the value of property "ScreenDetailed.devicePixelRatio".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) DevicePixelRatio() (ret float32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedDevicePixelRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "ScreenDetailed.label".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetScreenDetailedLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type FullscreenOptions struct {
	// NavigationUI is "FullscreenOptions.navigationUI"
	//
	// Optional, defaults to "auto".
	NavigationUI FullscreenNavigationUI
	// Screen is "FullscreenOptions.screen"
	//
	// Optional
	Screen ScreenDetailed

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FullscreenOptions with all fields set.
func (p FullscreenOptions) FromRef(ref js.Ref) FullscreenOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FullscreenOptions in the application heap.
func (p FullscreenOptions) New() js.Ref {
	return bindings.FullscreenOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FullscreenOptions) UpdateFrom(ref js.Ref) {
	bindings.FullscreenOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FullscreenOptions) Update(ref js.Ref) {
	bindings.FullscreenOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CSSStyleValue struct {
	ref js.Ref
}

func (this CSSStyleValue) Once() CSSStyleValue {
	this.Ref().Once()
	return this
}

func (this CSSStyleValue) Ref() js.Ref {
	return this.ref
}

func (this CSSStyleValue) FromRef(ref js.Ref) CSSStyleValue {
	this.ref = ref
	return this
}

func (this CSSStyleValue) Free() {
	this.Ref().Free()
}

// HasToString returns true if the method "CSSStyleValue.toString" exists.
func (this CSSStyleValue) HasToString() bool {
	return js.True == bindings.HasCSSStyleValueToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSStyleValue.toString".
func (this CSSStyleValue) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSStyleValueToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSStyleValue.toString".
func (this CSSStyleValue) ToString() (ret js.String) {
	bindings.CallCSSStyleValueToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSStyleValue.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleValue) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasParse returns true if the staticmethod "CSSStyleValue.parse" exists.
func (this CSSStyleValue) HasParse() bool {
	return js.True == bindings.HasCSSStyleValueParse(
		this.Ref(),
	)
}

// ParseFunc returns the staticmethod "CSSStyleValue.parse".
func (this CSSStyleValue) ParseFunc() (fn js.Func[func(property js.String, cssText js.String) CSSStyleValue]) {
	return fn.FromRef(
		bindings.CSSStyleValueParseFunc(
			this.Ref(),
		),
	)
}

// Parse calls the staticmethod "CSSStyleValue.parse".
func (this CSSStyleValue) Parse(property js.String, cssText js.String) (ret CSSStyleValue) {
	bindings.CallCSSStyleValueParse(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// TryParse calls the staticmethod "CSSStyleValue.parse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleValue) TryParse(property js.String, cssText js.String) (ret CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueParse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// HasParseAll returns true if the staticmethod "CSSStyleValue.parseAll" exists.
func (this CSSStyleValue) HasParseAll() bool {
	return js.True == bindings.HasCSSStyleValueParseAll(
		this.Ref(),
	)
}

// ParseAllFunc returns the staticmethod "CSSStyleValue.parseAll".
func (this CSSStyleValue) ParseAllFunc() (fn js.Func[func(property js.String, cssText js.String) js.Array[CSSStyleValue]]) {
	return fn.FromRef(
		bindings.CSSStyleValueParseAllFunc(
			this.Ref(),
		),
	)
}

// ParseAll calls the staticmethod "CSSStyleValue.parseAll".
func (this CSSStyleValue) ParseAll(property js.String, cssText js.String) (ret js.Array[CSSStyleValue]) {
	bindings.CallCSSStyleValueParseAll(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// TryParseAll calls the staticmethod "CSSStyleValue.parseAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CSSStyleValue) TryParseAll(property js.String, cssText js.String) (ret js.Array[CSSStyleValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueParseAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

type OneOf_undefined_CSSStyleValue struct {
	ref js.Ref
}

func (x OneOf_undefined_CSSStyleValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_undefined_CSSStyleValue) Free() {
	x.ref.Free()
}

func (x OneOf_undefined_CSSStyleValue) FromRef(ref js.Ref) OneOf_undefined_CSSStyleValue {
	return OneOf_undefined_CSSStyleValue{
		ref: ref,
	}
}

func (x OneOf_undefined_CSSStyleValue) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_undefined_CSSStyleValue) CSSStyleValue() CSSStyleValue {
	return CSSStyleValue{}.FromRef(x.ref)
}

type StylePropertyMapReadOnly struct {
	ref js.Ref
}

func (this StylePropertyMapReadOnly) Once() StylePropertyMapReadOnly {
	this.Ref().Once()
	return this
}

func (this StylePropertyMapReadOnly) Ref() js.Ref {
	return this.ref
}

func (this StylePropertyMapReadOnly) FromRef(ref js.Ref) StylePropertyMapReadOnly {
	this.ref = ref
	return this
}

func (this StylePropertyMapReadOnly) Free() {
	this.Ref().Free()
}

// Size returns the value of property "StylePropertyMapReadOnly.size".
//
// It returns ok=false if there is no such property.
func (this StylePropertyMapReadOnly) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStylePropertyMapReadOnlySize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "StylePropertyMapReadOnly.get" exists.
func (this StylePropertyMapReadOnly) HasGet() bool {
	return js.True == bindings.HasStylePropertyMapReadOnlyGet(
		this.Ref(),
	)
}

// GetFunc returns the method "StylePropertyMapReadOnly.get".
func (this StylePropertyMapReadOnly) GetFunc() (fn js.Func[func(property js.String) OneOf_undefined_CSSStyleValue]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "StylePropertyMapReadOnly.get".
func (this StylePropertyMapReadOnly) Get(property js.String) (ret OneOf_undefined_CSSStyleValue) {
	bindings.CallStylePropertyMapReadOnlyGet(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGet calls the method "StylePropertyMapReadOnly.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMapReadOnly) TryGet(property js.String) (ret OneOf_undefined_CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasGetAll returns true if the method "StylePropertyMapReadOnly.getAll" exists.
func (this StylePropertyMapReadOnly) HasGetAll() bool {
	return js.True == bindings.HasStylePropertyMapReadOnlyGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "StylePropertyMapReadOnly.getAll".
func (this StylePropertyMapReadOnly) GetAllFunc() (fn js.Func[func(property js.String) js.Array[CSSStyleValue]]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "StylePropertyMapReadOnly.getAll".
func (this StylePropertyMapReadOnly) GetAll(property js.String) (ret js.Array[CSSStyleValue]) {
	bindings.CallStylePropertyMapReadOnlyGetAll(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetAll calls the method "StylePropertyMapReadOnly.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMapReadOnly) TryGetAll(property js.String) (ret js.Array[CSSStyleValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasHas returns true if the method "StylePropertyMapReadOnly.has" exists.
func (this StylePropertyMapReadOnly) HasHas() bool {
	return js.True == bindings.HasStylePropertyMapReadOnlyHas(
		this.Ref(),
	)
}

// HasFunc returns the method "StylePropertyMapReadOnly.has".
func (this StylePropertyMapReadOnly) HasFunc() (fn js.Func[func(property js.String) bool]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyHasFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "StylePropertyMapReadOnly.has".
func (this StylePropertyMapReadOnly) Has(property js.String) (ret bool) {
	bindings.CallStylePropertyMapReadOnlyHas(
		this.Ref(), js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryHas calls the method "StylePropertyMapReadOnly.has"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this StylePropertyMapReadOnly) TryHas(property js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyHas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

type FocusableAreaSearchMode uint32

const (
	_ FocusableAreaSearchMode = iota

	FocusableAreaSearchMode_VISIBLE
	FocusableAreaSearchMode_ALL
)

func (FocusableAreaSearchMode) FromRef(str js.Ref) FocusableAreaSearchMode {
	return FocusableAreaSearchMode(bindings.ConstOfFocusableAreaSearchMode(str))
}

func (x FocusableAreaSearchMode) String() (string, bool) {
	switch x {
	case FocusableAreaSearchMode_VISIBLE:
		return "visible", true
	case FocusableAreaSearchMode_ALL:
		return "all", true
	default:
		return "", false
	}
}

type FocusableAreasOption struct {
	// Mode is "FocusableAreasOption.mode"
	//
	// Optional
	Mode FocusableAreaSearchMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FocusableAreasOption with all fields set.
func (p FocusableAreasOption) FromRef(ref js.Ref) FocusableAreasOption {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FocusableAreasOption in the application heap.
func (p FocusableAreasOption) New() js.Ref {
	return bindings.FocusableAreasOptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FocusableAreasOption) UpdateFrom(ref js.Ref) {
	bindings.FocusableAreasOptionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FocusableAreasOption) Update(ref js.Ref) {
	bindings.FocusableAreasOptionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SpatialNavigationDirection uint32

const (
	_ SpatialNavigationDirection = iota

	SpatialNavigationDirection_UP
	SpatialNavigationDirection_DOWN
	SpatialNavigationDirection_LEFT
	SpatialNavigationDirection_RIGHT
)

func (SpatialNavigationDirection) FromRef(str js.Ref) SpatialNavigationDirection {
	return SpatialNavigationDirection(bindings.ConstOfSpatialNavigationDirection(str))
}

func (x SpatialNavigationDirection) String() (string, bool) {
	switch x {
	case SpatialNavigationDirection_UP:
		return "up", true
	case SpatialNavigationDirection_DOWN:
		return "down", true
	case SpatialNavigationDirection_LEFT:
		return "left", true
	case SpatialNavigationDirection_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type SpatialNavigationSearchOptions struct {
	// Candidates is "SpatialNavigationSearchOptions.candidates"
	//
	// Optional
	Candidates js.Array[Node]
	// Container is "SpatialNavigationSearchOptions.container"
	//
	// Optional
	Container Node

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpatialNavigationSearchOptions with all fields set.
func (p SpatialNavigationSearchOptions) FromRef(ref js.Ref) SpatialNavigationSearchOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpatialNavigationSearchOptions in the application heap.
func (p SpatialNavigationSearchOptions) New() js.Ref {
	return bindings.SpatialNavigationSearchOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SpatialNavigationSearchOptions) UpdateFrom(ref js.Ref) {
	bindings.SpatialNavigationSearchOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SpatialNavigationSearchOptions) Update(ref js.Ref) {
	bindings.SpatialNavigationSearchOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DOMPointInit struct {
	// X is "DOMPointInit.x"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "DOMPointInit.y"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64
	// Z is "DOMPointInit.z"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Z MUST be set to true to make this field effective.
	Z float64
	// W is "DOMPointInit.w"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_W MUST be set to true to make this field effective.
	W float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.
	FFI_USE_Z bool // for Z.
	FFI_USE_W bool // for W.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMPointInit with all fields set.
func (p DOMPointInit) FromRef(ref js.Ref) DOMPointInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMPointInit in the application heap.
func (p DOMPointInit) New() js.Ref {
	return bindings.DOMPointInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DOMPointInit) UpdateFrom(ref js.Ref) {
	bindings.DOMPointInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DOMPointInit) Update(ref js.Ref) {
	bindings.DOMPointInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DOMRectInit struct {
	// X is "DOMRectInit.x"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "DOMRectInit.y"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64
	// Width is "DOMRectInit.width"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float64
	// Height is "DOMRectInit.height"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float64

	FFI_USE_X      bool // for X.
	FFI_USE_Y      bool // for Y.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMRectInit with all fields set.
func (p DOMRectInit) FromRef(ref js.Ref) DOMRectInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMRectInit in the application heap.
func (p DOMRectInit) New() js.Ref {
	return bindings.DOMRectInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DOMRectInit) UpdateFrom(ref js.Ref) {
	bindings.DOMRectInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DOMRectInit) Update(ref js.Ref) {
	bindings.DOMRectInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DOMQuadInit struct {
	// P1 is "DOMQuadInit.p1"
	//
	// Optional
	//
	// NOTE: P1.FFI_USE MUST be set to true to get P1 used.
	P1 DOMPointInit
	// P2 is "DOMQuadInit.p2"
	//
	// Optional
	//
	// NOTE: P2.FFI_USE MUST be set to true to get P2 used.
	P2 DOMPointInit
	// P3 is "DOMQuadInit.p3"
	//
	// Optional
	//
	// NOTE: P3.FFI_USE MUST be set to true to get P3 used.
	P3 DOMPointInit
	// P4 is "DOMQuadInit.p4"
	//
	// Optional
	//
	// NOTE: P4.FFI_USE MUST be set to true to get P4 used.
	P4 DOMPointInit

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMQuadInit with all fields set.
func (p DOMQuadInit) FromRef(ref js.Ref) DOMQuadInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMQuadInit in the application heap.
func (p DOMQuadInit) New() js.Ref {
	return bindings.DOMQuadInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DOMQuadInit) UpdateFrom(ref js.Ref) {
	bindings.DOMQuadInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DOMQuadInit) Update(ref js.Ref) {
	bindings.DOMQuadInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDOMRect(x float64, y float64, width float64, height float64) (ret DOMRect) {
	ret.ref = bindings.NewDOMRectByDOMRect(
		float64(x),
		float64(y),
		float64(width),
		float64(height))
	return
}

func NewDOMRectByDOMRect1(x float64, y float64, width float64) (ret DOMRect) {
	ret.ref = bindings.NewDOMRectByDOMRect1(
		float64(x),
		float64(y),
		float64(width))
	return
}

func NewDOMRectByDOMRect2(x float64, y float64) (ret DOMRect) {
	ret.ref = bindings.NewDOMRectByDOMRect2(
		float64(x),
		float64(y))
	return
}

func NewDOMRectByDOMRect3(x float64) (ret DOMRect) {
	ret.ref = bindings.NewDOMRectByDOMRect3(
		float64(x))
	return
}

func NewDOMRectByDOMRect4() (ret DOMRect) {
	ret.ref = bindings.NewDOMRectByDOMRect4()
	return
}

type DOMRect struct {
	DOMRectReadOnly
}

func (this DOMRect) Once() DOMRect {
	this.Ref().Once()
	return this
}

func (this DOMRect) Ref() js.Ref {
	return this.DOMRectReadOnly.Ref()
}

func (this DOMRect) FromRef(ref js.Ref) DOMRect {
	this.DOMRectReadOnly = this.DOMRectReadOnly.FromRef(ref)
	return this
}

func (this DOMRect) Free() {
	this.Ref().Free()
}

// X returns the value of property "DOMRect.x".
//
// It returns ok=false if there is no such property.
func (this DOMRect) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "DOMRect.x" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetX(val float64) bool {
	return js.True == bindings.SetDOMRectX(
		this.Ref(),
		float64(val),
	)
}

// Y returns the value of property "DOMRect.y".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "DOMRect.y" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetY(val float64) bool {
	return js.True == bindings.SetDOMRectY(
		this.Ref(),
		float64(val),
	)
}

// Width returns the value of property "DOMRect.width".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "DOMRect.width" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetWidth(val float64) bool {
	return js.True == bindings.SetDOMRectWidth(
		this.Ref(),
		float64(val),
	)
}

// Height returns the value of property "DOMRect.height".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "DOMRect.height" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetHeight(val float64) bool {
	return js.True == bindings.SetDOMRectHeight(
		this.Ref(),
		float64(val),
	)
}

// HasFromRect returns true if the staticmethod "DOMRect.fromRect" exists.
func (this DOMRect) HasFromRect() bool {
	return js.True == bindings.HasDOMRectFromRect(
		this.Ref(),
	)
}

// FromRectFunc returns the staticmethod "DOMRect.fromRect".
func (this DOMRect) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect calls the staticmethod "DOMRect.fromRect".
func (this DOMRect) FromRect(other DOMRectInit) (ret DOMRect) {
	bindings.CallDOMRectFromRect(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the staticmethod "DOMRect.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRect) TryFromRect(other DOMRectInit) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectFromRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromRect1 returns true if the staticmethod "DOMRect.fromRect" exists.
func (this DOMRect) HasFromRect1() bool {
	return js.True == bindings.HasDOMRectFromRect1(
		this.Ref(),
	)
}

// FromRect1Func returns the staticmethod "DOMRect.fromRect".
func (this DOMRect) FromRect1Func() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectFromRect1Func(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMRect.fromRect".
func (this DOMRect) FromRect1() (ret DOMRect) {
	bindings.CallDOMRectFromRect1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the staticmethod "DOMRect.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRect) TryFromRect1() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectFromRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewDOMPoint(x float64, y float64, z float64, w float64) (ret DOMPoint) {
	ret.ref = bindings.NewDOMPointByDOMPoint(
		float64(x),
		float64(y),
		float64(z),
		float64(w))
	return
}

func NewDOMPointByDOMPoint1(x float64, y float64, z float64) (ret DOMPoint) {
	ret.ref = bindings.NewDOMPointByDOMPoint1(
		float64(x),
		float64(y),
		float64(z))
	return
}

func NewDOMPointByDOMPoint2(x float64, y float64) (ret DOMPoint) {
	ret.ref = bindings.NewDOMPointByDOMPoint2(
		float64(x),
		float64(y))
	return
}

func NewDOMPointByDOMPoint3(x float64) (ret DOMPoint) {
	ret.ref = bindings.NewDOMPointByDOMPoint3(
		float64(x))
	return
}

func NewDOMPointByDOMPoint4() (ret DOMPoint) {
	ret.ref = bindings.NewDOMPointByDOMPoint4()
	return
}

type DOMPoint struct {
	DOMPointReadOnly
}

func (this DOMPoint) Once() DOMPoint {
	this.Ref().Once()
	return this
}

func (this DOMPoint) Ref() js.Ref {
	return this.DOMPointReadOnly.Ref()
}

func (this DOMPoint) FromRef(ref js.Ref) DOMPoint {
	this.DOMPointReadOnly = this.DOMPointReadOnly.FromRef(ref)
	return this
}

func (this DOMPoint) Free() {
	this.Ref().Free()
}

// X returns the value of property "DOMPoint.x".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "DOMPoint.x" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetX(val float64) bool {
	return js.True == bindings.SetDOMPointX(
		this.Ref(),
		float64(val),
	)
}

// Y returns the value of property "DOMPoint.y".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "DOMPoint.y" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetY(val float64) bool {
	return js.True == bindings.SetDOMPointY(
		this.Ref(),
		float64(val),
	)
}

// Z returns the value of property "DOMPoint.z".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetZ sets the value of property "DOMPoint.z" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetZ(val float64) bool {
	return js.True == bindings.SetDOMPointZ(
		this.Ref(),
		float64(val),
	)
}

// W returns the value of property "DOMPoint.w".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) W() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointW(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetW sets the value of property "DOMPoint.w" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetW(val float64) bool {
	return js.True == bindings.SetDOMPointW(
		this.Ref(),
		float64(val),
	)
}

// HasFromPoint returns true if the staticmethod "DOMPoint.fromPoint" exists.
func (this DOMPoint) HasFromPoint() bool {
	return js.True == bindings.HasDOMPointFromPoint(
		this.Ref(),
	)
}

// FromPointFunc returns the staticmethod "DOMPoint.fromPoint".
func (this DOMPoint) FromPointFunc() (fn js.Func[func(other DOMPointInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointFromPointFunc(
			this.Ref(),
		),
	)
}

// FromPoint calls the staticmethod "DOMPoint.fromPoint".
func (this DOMPoint) FromPoint(other DOMPointInit) (ret DOMPoint) {
	bindings.CallDOMPointFromPoint(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromPoint calls the staticmethod "DOMPoint.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPoint) TryFromPoint(other DOMPointInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointFromPoint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromPoint1 returns true if the staticmethod "DOMPoint.fromPoint" exists.
func (this DOMPoint) HasFromPoint1() bool {
	return js.True == bindings.HasDOMPointFromPoint1(
		this.Ref(),
	)
}

// FromPoint1Func returns the staticmethod "DOMPoint.fromPoint".
func (this DOMPoint) FromPoint1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointFromPoint1Func(
			this.Ref(),
		),
	)
}

// FromPoint1 calls the staticmethod "DOMPoint.fromPoint".
func (this DOMPoint) FromPoint1() (ret DOMPoint) {
	bindings.CallDOMPointFromPoint1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromPoint1 calls the staticmethod "DOMPoint.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMPoint) TryFromPoint1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointFromPoint1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewDOMQuad(p1 DOMPointInit, p2 DOMPointInit, p3 DOMPointInit, p4 DOMPointInit) (ret DOMQuad) {
	ret.ref = bindings.NewDOMQuadByDOMQuad(
		js.Pointer(&p1),
		js.Pointer(&p2),
		js.Pointer(&p3),
		js.Pointer(&p4))
	return
}

func NewDOMQuadByDOMQuad1(p1 DOMPointInit, p2 DOMPointInit, p3 DOMPointInit) (ret DOMQuad) {
	ret.ref = bindings.NewDOMQuadByDOMQuad1(
		js.Pointer(&p1),
		js.Pointer(&p2),
		js.Pointer(&p3))
	return
}

func NewDOMQuadByDOMQuad2(p1 DOMPointInit, p2 DOMPointInit) (ret DOMQuad) {
	ret.ref = bindings.NewDOMQuadByDOMQuad2(
		js.Pointer(&p1),
		js.Pointer(&p2))
	return
}

func NewDOMQuadByDOMQuad3(p1 DOMPointInit) (ret DOMQuad) {
	ret.ref = bindings.NewDOMQuadByDOMQuad3(
		js.Pointer(&p1))
	return
}

func NewDOMQuadByDOMQuad4() (ret DOMQuad) {
	ret.ref = bindings.NewDOMQuadByDOMQuad4()
	return
}

type DOMQuad struct {
	ref js.Ref
}

func (this DOMQuad) Once() DOMQuad {
	this.Ref().Once()
	return this
}

func (this DOMQuad) Ref() js.Ref {
	return this.ref
}

func (this DOMQuad) FromRef(ref js.Ref) DOMQuad {
	this.ref = ref
	return this
}

func (this DOMQuad) Free() {
	this.Ref().Free()
}

// P1 returns the value of property "DOMQuad.p1".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P1() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// P2 returns the value of property "DOMQuad.p2".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P2() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// P3 returns the value of property "DOMQuad.p3".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P3() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP3(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// P4 returns the value of property "DOMQuad.p4".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P4() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP4(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFromRect returns true if the staticmethod "DOMQuad.fromRect" exists.
func (this DOMQuad) HasFromRect() bool {
	return js.True == bindings.HasDOMQuadFromRect(
		this.Ref(),
	)
}

// FromRectFunc returns the staticmethod "DOMQuad.fromRect".
func (this DOMQuad) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect calls the staticmethod "DOMQuad.fromRect".
func (this DOMQuad) FromRect(other DOMRectInit) (ret DOMQuad) {
	bindings.CallDOMQuadFromRect(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the staticmethod "DOMQuad.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryFromRect(other DOMRectInit) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromRect1 returns true if the staticmethod "DOMQuad.fromRect" exists.
func (this DOMQuad) HasFromRect1() bool {
	return js.True == bindings.HasDOMQuadFromRect1(
		this.Ref(),
	)
}

// FromRect1Func returns the staticmethod "DOMQuad.fromRect".
func (this DOMQuad) FromRect1Func() (fn js.Func[func() DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromRect1Func(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMQuad.fromRect".
func (this DOMQuad) FromRect1() (ret DOMQuad) {
	bindings.CallDOMQuadFromRect1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the staticmethod "DOMQuad.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryFromRect1() (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromQuad returns true if the staticmethod "DOMQuad.fromQuad" exists.
func (this DOMQuad) HasFromQuad() bool {
	return js.True == bindings.HasDOMQuadFromQuad(
		this.Ref(),
	)
}

// FromQuadFunc returns the staticmethod "DOMQuad.fromQuad".
func (this DOMQuad) FromQuadFunc() (fn js.Func[func(other DOMQuadInit) DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromQuadFunc(
			this.Ref(),
		),
	)
}

// FromQuad calls the staticmethod "DOMQuad.fromQuad".
func (this DOMQuad) FromQuad(other DOMQuadInit) (ret DOMQuad) {
	bindings.CallDOMQuadFromQuad(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromQuad calls the staticmethod "DOMQuad.fromQuad"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryFromQuad(other DOMQuadInit) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromQuad(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromQuad1 returns true if the staticmethod "DOMQuad.fromQuad" exists.
func (this DOMQuad) HasFromQuad1() bool {
	return js.True == bindings.HasDOMQuadFromQuad1(
		this.Ref(),
	)
}

// FromQuad1Func returns the staticmethod "DOMQuad.fromQuad".
func (this DOMQuad) FromQuad1Func() (fn js.Func[func() DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromQuad1Func(
			this.Ref(),
		),
	)
}

// FromQuad1 calls the staticmethod "DOMQuad.fromQuad".
func (this DOMQuad) FromQuad1() (ret DOMQuad) {
	bindings.CallDOMQuadFromQuad1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromQuad1 calls the staticmethod "DOMQuad.fromQuad"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryFromQuad1() (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromQuad1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetBounds returns true if the method "DOMQuad.getBounds" exists.
func (this DOMQuad) HasGetBounds() bool {
	return js.True == bindings.HasDOMQuadGetBounds(
		this.Ref(),
	)
}

// GetBoundsFunc returns the method "DOMQuad.getBounds".
func (this DOMQuad) GetBoundsFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.DOMQuadGetBoundsFunc(
			this.Ref(),
		),
	)
}

// GetBounds calls the method "DOMQuad.getBounds".
func (this DOMQuad) GetBounds() (ret DOMRect) {
	bindings.CallDOMQuadGetBounds(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBounds calls the method "DOMQuad.getBounds"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryGetBounds() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadGetBounds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "DOMQuad.toJSON" exists.
func (this DOMQuad) HasToJSON() bool {
	return js.True == bindings.HasDOMQuadToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "DOMQuad.toJSON".
func (this DOMQuad) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMQuadToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMQuad.toJSON".
func (this DOMQuad) ToJSON() (ret js.Object) {
	bindings.CallDOMQuadToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMQuad.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMQuad) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CSSBoxType uint32

const (
	_ CSSBoxType = iota

	CSSBoxType_MARGIN
	CSSBoxType_BORDER
	CSSBoxType_PADDING
	CSSBoxType_CONTENT
)

func (CSSBoxType) FromRef(str js.Ref) CSSBoxType {
	return CSSBoxType(bindings.ConstOfCSSBoxType(str))
}

func (x CSSBoxType) String() (string, bool) {
	switch x {
	case CSSBoxType_MARGIN:
		return "margin", true
	case CSSBoxType_BORDER:
		return "border", true
	case CSSBoxType_PADDING:
		return "padding", true
	case CSSBoxType_CONTENT:
		return "content", true
	default:
		return "", false
	}
}

type ConvertCoordinateOptions struct {
	// FromBox is "ConvertCoordinateOptions.fromBox"
	//
	// Optional, defaults to "border".
	FromBox CSSBoxType
	// ToBox is "ConvertCoordinateOptions.toBox"
	//
	// Optional, defaults to "border".
	ToBox CSSBoxType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConvertCoordinateOptions with all fields set.
func (p ConvertCoordinateOptions) FromRef(ref js.Ref) ConvertCoordinateOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConvertCoordinateOptions in the application heap.
func (p ConvertCoordinateOptions) New() js.Ref {
	return bindings.ConvertCoordinateOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ConvertCoordinateOptions) UpdateFrom(ref js.Ref) {
	bindings.ConvertCoordinateOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConvertCoordinateOptions) Update(ref js.Ref) {
	bindings.ConvertCoordinateOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewDOMRectReadOnly(x float64, y float64, width float64, height float64) (ret DOMRectReadOnly) {
	ret.ref = bindings.NewDOMRectReadOnlyByDOMRectReadOnly(
		float64(x),
		float64(y),
		float64(width),
		float64(height))
	return
}

func NewDOMRectReadOnlyByDOMRectReadOnly1(x float64, y float64, width float64) (ret DOMRectReadOnly) {
	ret.ref = bindings.NewDOMRectReadOnlyByDOMRectReadOnly1(
		float64(x),
		float64(y),
		float64(width))
	return
}

func NewDOMRectReadOnlyByDOMRectReadOnly2(x float64, y float64) (ret DOMRectReadOnly) {
	ret.ref = bindings.NewDOMRectReadOnlyByDOMRectReadOnly2(
		float64(x),
		float64(y))
	return
}

func NewDOMRectReadOnlyByDOMRectReadOnly3(x float64) (ret DOMRectReadOnly) {
	ret.ref = bindings.NewDOMRectReadOnlyByDOMRectReadOnly3(
		float64(x))
	return
}

func NewDOMRectReadOnlyByDOMRectReadOnly4() (ret DOMRectReadOnly) {
	ret.ref = bindings.NewDOMRectReadOnlyByDOMRectReadOnly4()
	return
}

type DOMRectReadOnly struct {
	ref js.Ref
}

func (this DOMRectReadOnly) Once() DOMRectReadOnly {
	this.Ref().Once()
	return this
}

func (this DOMRectReadOnly) Ref() js.Ref {
	return this.ref
}

func (this DOMRectReadOnly) FromRef(ref js.Ref) DOMRectReadOnly {
	this.ref = ref
	return this
}

func (this DOMRectReadOnly) Free() {
	this.Ref().Free()
}

// X returns the value of property "DOMRectReadOnly.x".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DOMRectReadOnly.y".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "DOMRectReadOnly.width".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "DOMRectReadOnly.height".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "DOMRectReadOnly.top".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Top() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Right returns the value of property "DOMRectReadOnly.right".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Right() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyRight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Bottom returns the value of property "DOMRectReadOnly.bottom".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Bottom() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyBottom(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Left returns the value of property "DOMRectReadOnly.left".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Left() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFromRect returns true if the staticmethod "DOMRectReadOnly.fromRect" exists.
func (this DOMRectReadOnly) HasFromRect() bool {
	return js.True == bindings.HasDOMRectReadOnlyFromRect(
		this.Ref(),
	)
}

// FromRectFunc returns the staticmethod "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMRectReadOnly]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect calls the staticmethod "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRect(other DOMRectInit) (ret DOMRectReadOnly) {
	bindings.CallDOMRectReadOnlyFromRect(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the staticmethod "DOMRectReadOnly.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRectReadOnly) TryFromRect(other DOMRectInit) (ret DOMRectReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyFromRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromRect1 returns true if the staticmethod "DOMRectReadOnly.fromRect" exists.
func (this DOMRectReadOnly) HasFromRect1() bool {
	return js.True == bindings.HasDOMRectReadOnlyFromRect1(
		this.Ref(),
	)
}

// FromRect1Func returns the staticmethod "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRect1Func() (fn js.Func[func() DOMRectReadOnly]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyFromRect1Func(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRect1() (ret DOMRectReadOnly) {
	bindings.CallDOMRectReadOnlyFromRect1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the staticmethod "DOMRectReadOnly.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRectReadOnly) TryFromRect1() (ret DOMRectReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyFromRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "DOMRectReadOnly.toJSON" exists.
func (this DOMRectReadOnly) HasToJSON() bool {
	return js.True == bindings.HasDOMRectReadOnlyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "DOMRectReadOnly.toJSON".
func (this DOMRectReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMRectReadOnly.toJSON".
func (this DOMRectReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMRectReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMRectReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMRectReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AssignedNodesOptions struct {
	// Flatten is "AssignedNodesOptions.flatten"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Flatten MUST be set to true to make this field effective.
	Flatten bool

	FFI_USE_Flatten bool // for Flatten.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssignedNodesOptions with all fields set.
func (p AssignedNodesOptions) FromRef(ref js.Ref) AssignedNodesOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssignedNodesOptions in the application heap.
func (p AssignedNodesOptions) New() js.Ref {
	return bindings.AssignedNodesOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AssignedNodesOptions) UpdateFrom(ref js.Ref) {
	bindings.AssignedNodesOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AssignedNodesOptions) Update(ref js.Ref) {
	bindings.AssignedNodesOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Element_Text struct {
	ref js.Ref
}

func (x OneOf_Element_Text) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Element_Text) Free() {
	x.ref.Free()
}

func (x OneOf_Element_Text) FromRef(ref js.Ref) OneOf_Element_Text {
	return OneOf_Element_Text{
		ref: ref,
	}
}

func (x OneOf_Element_Text) Element() Element {
	return Element{}.FromRef(x.ref)
}

func (x OneOf_Element_Text) Text() Text {
	return Text{}.FromRef(x.ref)
}

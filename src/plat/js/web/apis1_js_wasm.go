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
// The returned bool will be false if there is no such property.
func (this AnimationNodeList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAnimationNodeListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "AnimationNodeList.item".
//
// The returned bool will be false if there is no such method.
func (this AnimationNodeList) Item(index uint32) (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.CallAnimationNodeListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return AnimationEffect{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "AnimationNodeList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationNodeList) ItemFunc() (fn js.Func[func(index uint32) AnimationEffect]) {
	return fn.FromRef(
		bindings.AnimationNodeListItemFunc(
			this.Ref(),
		),
	)
}

func NewGroupEffect(children js.Array[AnimationEffect], timing OneOf_Float64_EffectTiming) GroupEffect {
	return GroupEffect{}.FromRef(
		bindings.NewGroupEffectByGroupEffect(
			children.Ref(),
			timing.Ref()),
	)
}

func NewGroupEffectByGroupEffect1(children js.Array[AnimationEffect]) GroupEffect {
	return GroupEffect{}.FromRef(
		bindings.NewGroupEffectByGroupEffect1(
			children.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this GroupEffect) Children() (AnimationNodeList, bool) {
	var _ok bool
	_ret := bindings.GetGroupEffectChildren(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationNodeList{}.FromRef(_ret), _ok
}

// FirstChild returns the value of property "GroupEffect.firstChild".
//
// The returned bool will be false if there is no such property.
func (this GroupEffect) FirstChild() (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.GetGroupEffectFirstChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationEffect{}.FromRef(_ret), _ok
}

// LastChild returns the value of property "GroupEffect.lastChild".
//
// The returned bool will be false if there is no such property.
func (this GroupEffect) LastChild() (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.GetGroupEffectLastChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationEffect{}.FromRef(_ret), _ok
}

// Clone calls the method "GroupEffect.clone".
//
// The returned bool will be false if there is no such method.
func (this GroupEffect) Clone() (GroupEffect, bool) {
	var _ok bool
	_ret := bindings.CallGroupEffectClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return GroupEffect{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "GroupEffect.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GroupEffect) CloneFunc() (fn js.Func[func() GroupEffect]) {
	return fn.FromRef(
		bindings.GroupEffectCloneFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "GroupEffect.prepend".
//
// The returned bool will be false if there is no such method.
func (this GroupEffect) Prepend(effects ...AnimationEffect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGroupEffectPrepend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PrependFunc returns the method "GroupEffect.prepend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GroupEffect) PrependFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.GroupEffectPrependFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "GroupEffect.append".
//
// The returned bool will be false if there is no such method.
func (this GroupEffect) Append(effects ...AnimationEffect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGroupEffectAppend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "GroupEffect.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GroupEffect) AppendFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.GroupEffectAppendFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this AnimationEffect) Parent() (GroupEffect, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEffectParent(
		this.Ref(), js.Pointer(&_ok),
	)
	return GroupEffect{}.FromRef(_ret), _ok
}

// PreviousSibling returns the value of property "AnimationEffect.previousSibling".
//
// The returned bool will be false if there is no such property.
func (this AnimationEffect) PreviousSibling() (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEffectPreviousSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationEffect{}.FromRef(_ret), _ok
}

// NextSibling returns the value of property "AnimationEffect.nextSibling".
//
// The returned bool will be false if there is no such property.
func (this AnimationEffect) NextSibling() (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEffectNextSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationEffect{}.FromRef(_ret), _ok
}

// GetTiming calls the method "AnimationEffect.getTiming".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) GetTiming() (EffectTiming, bool) {
	var _ret EffectTiming
	_ok := js.True == bindings.CallAnimationEffectGetTiming(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetTimingFunc returns the method "AnimationEffect.getTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) GetTimingFunc() (fn js.Func[func() EffectTiming]) {
	return fn.FromRef(
		bindings.AnimationEffectGetTimingFunc(
			this.Ref(),
		),
	)
}

// GetComputedTiming calls the method "AnimationEffect.getComputedTiming".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) GetComputedTiming() (ComputedEffectTiming, bool) {
	var _ret ComputedEffectTiming
	_ok := js.True == bindings.CallAnimationEffectGetComputedTiming(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetComputedTimingFunc returns the method "AnimationEffect.getComputedTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) GetComputedTimingFunc() (fn js.Func[func() ComputedEffectTiming]) {
	return fn.FromRef(
		bindings.AnimationEffectGetComputedTimingFunc(
			this.Ref(),
		),
	)
}

// UpdateTiming calls the method "AnimationEffect.updateTiming".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) UpdateTiming(timing OptionalEffectTiming) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectUpdateTiming(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&timing),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateTimingFunc returns the method "AnimationEffect.updateTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) UpdateTimingFunc() (fn js.Func[func(timing OptionalEffectTiming)]) {
	return fn.FromRef(
		bindings.AnimationEffectUpdateTimingFunc(
			this.Ref(),
		),
	)
}

// UpdateTiming1 calls the method "AnimationEffect.updateTiming".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) UpdateTiming1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectUpdateTiming1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateTiming1Func returns the method "AnimationEffect.updateTiming".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) UpdateTiming1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationEffectUpdateTiming1Func(
			this.Ref(),
		),
	)
}

// Before calls the method "AnimationEffect.before".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) Before(effects ...AnimationEffect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectBefore(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeforeFunc returns the method "AnimationEffect.before".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) BeforeFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectBeforeFunc(
			this.Ref(),
		),
	)
}

// After calls the method "AnimationEffect.after".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) After(effects ...AnimationEffect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectAfter(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AfterFunc returns the method "AnimationEffect.after".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) AfterFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectAfterFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "AnimationEffect.replace".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) Replace(effects ...AnimationEffect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectReplace(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceFunc returns the method "AnimationEffect.replace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) ReplaceFunc() (fn js.Func[func(effects ...AnimationEffect)]) {
	return fn.FromRef(
		bindings.AnimationEffectReplaceFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "AnimationEffect.remove".
//
// The returned bool will be false if there is no such method.
func (this AnimationEffect) Remove() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationEffectRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "AnimationEffect.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationEffect) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationEffectRemoveFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this AnimationTimeline) CurrentTime() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationTimelineCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Duration returns the value of property "AnimationTimeline.duration".
//
// The returned bool will be false if there is no such property.
func (this AnimationTimeline) Duration() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationTimelineDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Play calls the method "AnimationTimeline.play".
//
// The returned bool will be false if there is no such method.
func (this AnimationTimeline) Play(effect AnimationEffect) (Animation, bool) {
	var _ok bool
	_ret := bindings.CallAnimationTimelinePlay(
		this.Ref(), js.Pointer(&_ok),
		effect.Ref(),
	)

	return Animation{}.FromRef(_ret), _ok
}

// PlayFunc returns the method "AnimationTimeline.play".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationTimeline) PlayFunc() (fn js.Func[func(effect AnimationEffect) Animation]) {
	return fn.FromRef(
		bindings.AnimationTimelinePlayFunc(
			this.Ref(),
		),
	)
}

// Play1 calls the method "AnimationTimeline.play".
//
// The returned bool will be false if there is no such method.
func (this AnimationTimeline) Play1() (Animation, bool) {
	var _ok bool
	_ret := bindings.CallAnimationTimelinePlay1(
		this.Ref(), js.Pointer(&_ok),
	)

	return Animation{}.FromRef(_ret), _ok
}

// Play1Func returns the method "AnimationTimeline.play".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationTimeline) Play1Func() (fn js.Func[func() Animation]) {
	return fn.FromRef(
		bindings.AnimationTimelinePlay1Func(
			this.Ref(),
		),
	)
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

func NewAnimation(effect AnimationEffect, timeline AnimationTimeline) Animation {
	return Animation{}.FromRef(
		bindings.NewAnimationByAnimation(
			effect.Ref(),
			timeline.Ref()),
	)
}

func NewAnimationByAnimation1(effect AnimationEffect) Animation {
	return Animation{}.FromRef(
		bindings.NewAnimationByAnimation1(
			effect.Ref()),
	)
}

func NewAnimationByAnimation2() Animation {
	return Animation{}.FromRef(
		bindings.NewAnimationByAnimation2(),
	)
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
// The returned bool will be false if there is no such property.
func (this Animation) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAnimationId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id sets the value of property "Animation.id" to val.
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
// The returned bool will be false if there is no such property.
func (this Animation) Effect() (AnimationEffect, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEffect(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationEffect{}.FromRef(_ret), _ok
}

// Effect sets the value of property "Animation.effect" to val.
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
// The returned bool will be false if there is no such property.
func (this Animation) Timeline() (AnimationTimeline, bool) {
	var _ok bool
	_ret := bindings.GetAnimationTimeline(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationTimeline{}.FromRef(_ret), _ok
}

// Timeline sets the value of property "Animation.timeline" to val.
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
// The returned bool will be false if there is no such property.
func (this Animation) PlaybackRate() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAnimationPlaybackRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PlaybackRate sets the value of property "Animation.playbackRate" to val.
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
// The returned bool will be false if there is no such property.
func (this Animation) PlayState() (AnimationPlayState, bool) {
	var _ok bool
	_ret := bindings.GetAnimationPlayState(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationPlayState(_ret), _ok
}

// ReplaceState returns the value of property "Animation.replaceState".
//
// The returned bool will be false if there is no such property.
func (this Animation) ReplaceState() (AnimationReplaceState, bool) {
	var _ok bool
	_ret := bindings.GetAnimationReplaceState(
		this.Ref(), js.Pointer(&_ok),
	)
	return AnimationReplaceState(_ret), _ok
}

// Pending returns the value of property "Animation.pending".
//
// The returned bool will be false if there is no such property.
func (this Animation) Pending() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAnimationPending(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Ready returns the value of property "Animation.ready".
//
// The returned bool will be false if there is no such property.
func (this Animation) Ready() (js.Promise[Animation], bool) {
	var _ok bool
	_ret := bindings.GetAnimationReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[Animation]{}.FromRef(_ret), _ok
}

// Finished returns the value of property "Animation.finished".
//
// The returned bool will be false if there is no such property.
func (this Animation) Finished() (js.Promise[Animation], bool) {
	var _ok bool
	_ret := bindings.GetAnimationFinished(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[Animation]{}.FromRef(_ret), _ok
}

// StartTime returns the value of property "Animation.startTime".
//
// The returned bool will be false if there is no such property.
func (this Animation) StartTime() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationStartTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// StartTime sets the value of property "Animation.startTime" to val.
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
// The returned bool will be false if there is no such property.
func (this Animation) CurrentTime() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// CurrentTime sets the value of property "Animation.currentTime" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetCurrentTime(val CSSNumberish) bool {
	return js.True == bindings.SetAnimationCurrentTime(
		this.Ref(),
		val.Ref(),
	)
}

// Cancel calls the method "Animation.cancel".
//
// The returned bool will be false if there is no such method.
func (this Animation) Cancel() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationCancel(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelFunc returns the method "Animation.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationCancelFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "Animation.finish".
//
// The returned bool will be false if there is no such method.
func (this Animation) Finish() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationFinish(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FinishFunc returns the method "Animation.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationFinishFunc(
			this.Ref(),
		),
	)
}

// Play calls the method "Animation.play".
//
// The returned bool will be false if there is no such method.
func (this Animation) Play() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationPlay(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PlayFunc returns the method "Animation.play".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) PlayFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPlayFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "Animation.pause".
//
// The returned bool will be false if there is no such method.
func (this Animation) Pause() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationPause(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseFunc returns the method "Animation.pause".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPauseFunc(
			this.Ref(),
		),
	)
}

// UpdatePlaybackRate calls the method "Animation.updatePlaybackRate".
//
// The returned bool will be false if there is no such method.
func (this Animation) UpdatePlaybackRate(playbackRate float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationUpdatePlaybackRate(
		this.Ref(), js.Pointer(&_ok),
		float64(playbackRate),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdatePlaybackRateFunc returns the method "Animation.updatePlaybackRate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) UpdatePlaybackRateFunc() (fn js.Func[func(playbackRate float64)]) {
	return fn.FromRef(
		bindings.AnimationUpdatePlaybackRateFunc(
			this.Ref(),
		),
	)
}

// Reverse calls the method "Animation.reverse".
//
// The returned bool will be false if there is no such method.
func (this Animation) Reverse() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationReverse(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReverseFunc returns the method "Animation.reverse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) ReverseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationReverseFunc(
			this.Ref(),
		),
	)
}

// Persist calls the method "Animation.persist".
//
// The returned bool will be false if there is no such method.
func (this Animation) Persist() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationPersist(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PersistFunc returns the method "Animation.persist".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) PersistFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationPersistFunc(
			this.Ref(),
		),
	)
}

// CommitStyles calls the method "Animation.commitStyles".
//
// The returned bool will be false if there is no such method.
func (this Animation) CommitStyles() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationCommitStyles(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CommitStylesFunc returns the method "Animation.commitStyles".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Animation) CommitStylesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AnimationCommitStylesFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this MediaList) MediaText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaListMediaText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MediaText sets the value of property "MediaList.mediaText" to val.
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
// The returned bool will be false if there is no such property.
func (this MediaList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetMediaListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "MediaList.item".
//
// The returned bool will be false if there is no such method.
func (this MediaList) Item(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallMediaListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "MediaList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.MediaListItemFunc(
			this.Ref(),
		),
	)
}

// AppendMedium calls the method "MediaList.appendMedium".
//
// The returned bool will be false if there is no such method.
func (this MediaList) AppendMedium(medium js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaListAppendMedium(
		this.Ref(), js.Pointer(&_ok),
		medium.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendMediumFunc returns the method "MediaList.appendMedium".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaList) AppendMediumFunc() (fn js.Func[func(medium js.String)]) {
	return fn.FromRef(
		bindings.MediaListAppendMediumFunc(
			this.Ref(),
		),
	)
}

// DeleteMedium calls the method "MediaList.deleteMedium".
//
// The returned bool will be false if there is no such method.
func (this MediaList) DeleteMedium(medium js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaListDeleteMedium(
		this.Ref(), js.Pointer(&_ok),
		medium.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteMediumFunc returns the method "MediaList.deleteMedium".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaList) DeleteMediumFunc() (fn js.Func[func(medium js.String)]) {
	return fn.FromRef(
		bindings.MediaListDeleteMediumFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSRule) CssText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSRuleCssText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CssText sets the value of property "CSSRule.cssText" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRule) ParentRule() (CSSRule, bool) {
	var _ok bool
	_ret := bindings.GetCSSRuleParentRule(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRule{}.FromRef(_ret), _ok
}

// ParentStyleSheet returns the value of property "CSSRule.parentStyleSheet".
//
// The returned bool will be false if there is no such property.
func (this CSSRule) ParentStyleSheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetCSSRuleParentStyleSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

// Type returns the value of property "CSSRule.type".
//
// The returned bool will be false if there is no such property.
func (this CSSRule) Type() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetCSSRuleType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this CSSRuleList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSRuleListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "CSSRuleList.item".
//
// The returned bool will be false if there is no such method.
func (this CSSRuleList) Item(index uint32) (CSSRule, bool) {
	var _ok bool
	_ret := bindings.CallCSSRuleListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSRule{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "CSSRuleList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSRuleList) ItemFunc() (fn js.Func[func(index uint32) CSSRule]) {
	return fn.FromRef(
		bindings.CSSRuleListItemFunc(
			this.Ref(),
		),
	)
}

func NewCSSStyleSheet(options CSSStyleSheetInit) CSSStyleSheet {
	return CSSStyleSheet{}.FromRef(
		bindings.NewCSSStyleSheetByCSSStyleSheet(
			js.Pointer(&options)),
	)
}

func NewCSSStyleSheetByCSSStyleSheet1() CSSStyleSheet {
	return CSSStyleSheet{}.FromRef(
		bindings.NewCSSStyleSheetByCSSStyleSheet1(),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSStyleSheet) OwnerRule() (CSSRule, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleSheetOwnerRule(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRule{}.FromRef(_ret), _ok
}

// CssRules returns the value of property "CSSStyleSheet.cssRules".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleSheet) CssRules() (CSSRuleList, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleSheetCssRules(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRuleList{}.FromRef(_ret), _ok
}

// Rules returns the value of property "CSSStyleSheet.rules".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleSheet) Rules() (CSSRuleList, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleSheetRules(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRuleList{}.FromRef(_ret), _ok
}

// InsertRule calls the method "CSSStyleSheet.insertRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) InsertRule(rule js.String, index uint32) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetInsertRule(
		this.Ref(), js.Pointer(&_ok),
		rule.Ref(),
		uint32(index),
	)

	return uint32(_ret), _ok
}

// InsertRuleFunc returns the method "CSSStyleSheet.insertRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) InsertRuleFunc() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetInsertRuleFunc(
			this.Ref(),
		),
	)
}

// InsertRule1 calls the method "CSSStyleSheet.insertRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) InsertRule1(rule js.String) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetInsertRule1(
		this.Ref(), js.Pointer(&_ok),
		rule.Ref(),
	)

	return uint32(_ret), _ok
}

// InsertRule1Func returns the method "CSSStyleSheet.insertRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) InsertRule1Func() (fn js.Func[func(rule js.String) uint32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetInsertRule1Func(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSStyleSheet.deleteRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) DeleteRule(index uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetDeleteRule(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRuleFunc returns the method "CSSStyleSheet.deleteRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) DeleteRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetDeleteRuleFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "CSSStyleSheet.replace".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) Replace(text js.String) (js.Promise[CSSStyleSheet], bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetReplace(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
	)

	return js.Promise[CSSStyleSheet]{}.FromRef(_ret), _ok
}

// ReplaceFunc returns the method "CSSStyleSheet.replace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) ReplaceFunc() (fn js.Func[func(text js.String) js.Promise[CSSStyleSheet]]) {
	return fn.FromRef(
		bindings.CSSStyleSheetReplaceFunc(
			this.Ref(),
		),
	)
}

// ReplaceSync calls the method "CSSStyleSheet.replaceSync".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) ReplaceSync(text js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetReplaceSync(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceSyncFunc returns the method "CSSStyleSheet.replaceSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) ReplaceSyncFunc() (fn js.Func[func(text js.String)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetReplaceSyncFunc(
			this.Ref(),
		),
	)
}

// AddRule calls the method "CSSStyleSheet.addRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) AddRule(selector js.String, style js.String, index uint32) (int32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetAddRule(
		this.Ref(), js.Pointer(&_ok),
		selector.Ref(),
		style.Ref(),
		uint32(index),
	)

	return int32(_ret), _ok
}

// AddRuleFunc returns the method "CSSStyleSheet.addRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) AddRuleFunc() (fn js.Func[func(selector js.String, style js.String, index uint32) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRuleFunc(
			this.Ref(),
		),
	)
}

// AddRule1 calls the method "CSSStyleSheet.addRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) AddRule1(selector js.String, style js.String) (int32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetAddRule1(
		this.Ref(), js.Pointer(&_ok),
		selector.Ref(),
		style.Ref(),
	)

	return int32(_ret), _ok
}

// AddRule1Func returns the method "CSSStyleSheet.addRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) AddRule1Func() (fn js.Func[func(selector js.String, style js.String) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule1Func(
			this.Ref(),
		),
	)
}

// AddRule2 calls the method "CSSStyleSheet.addRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) AddRule2(selector js.String) (int32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetAddRule2(
		this.Ref(), js.Pointer(&_ok),
		selector.Ref(),
	)

	return int32(_ret), _ok
}

// AddRule2Func returns the method "CSSStyleSheet.addRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) AddRule2Func() (fn js.Func[func(selector js.String) int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule2Func(
			this.Ref(),
		),
	)
}

// AddRule3 calls the method "CSSStyleSheet.addRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) AddRule3() (int32, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetAddRule3(
		this.Ref(), js.Pointer(&_ok),
	)

	return int32(_ret), _ok
}

// AddRule3Func returns the method "CSSStyleSheet.addRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) AddRule3Func() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.CSSStyleSheetAddRule3Func(
			this.Ref(),
		),
	)
}

// RemoveRule calls the method "CSSStyleSheet.removeRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) RemoveRule(index uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetRemoveRule(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveRuleFunc returns the method "CSSStyleSheet.removeRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) RemoveRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSStyleSheetRemoveRuleFunc(
			this.Ref(),
		),
	)
}

// RemoveRule1 calls the method "CSSStyleSheet.removeRule".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleSheet) RemoveRule1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleSheetRemoveRule1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveRule1Func returns the method "CSSStyleSheet.removeRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleSheet) RemoveRule1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CSSStyleSheetRemoveRule1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this StyleSheetList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "StyleSheetList.item".
//
// The returned bool will be false if there is no such method.
func (this StyleSheetList) Item(index uint32) (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.CallStyleSheetListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSStyleSheet{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "StyleSheetList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StyleSheetList) ItemFunc() (fn js.Func[func(index uint32) CSSStyleSheet]) {
	return fn.FromRef(
		bindings.StyleSheetListItemFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ShadowRoot) Mode() (ShadowRootMode, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return ShadowRootMode(_ret), _ok
}

// DelegatesFocus returns the value of property "ShadowRoot.delegatesFocus".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) DelegatesFocus() (bool, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootDelegatesFocus(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// SlotAssignment returns the value of property "ShadowRoot.slotAssignment".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) SlotAssignment() (SlotAssignmentMode, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootSlotAssignment(
		this.Ref(), js.Pointer(&_ok),
	)
	return SlotAssignmentMode(_ret), _ok
}

// Host returns the value of property "ShadowRoot.host".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) Host() (Element, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// InnerHTML returns the value of property "ShadowRoot.innerHTML".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) InnerHTML() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootInnerHTML(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InnerHTML sets the value of property "ShadowRoot.innerHTML" to val.
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
// The returned bool will be false if there is no such property.
func (this ShadowRoot) FullscreenElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootFullscreenElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// StyleSheets returns the value of property "ShadowRoot.styleSheets".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) StyleSheets() (StyleSheetList, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootStyleSheets(
		this.Ref(), js.Pointer(&_ok),
	)
	return StyleSheetList{}.FromRef(_ret), _ok
}

// AdoptedStyleSheets returns the value of property "ShadowRoot.adoptedStyleSheets".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) AdoptedStyleSheets() (js.ObservableArray[CSSStyleSheet], bool) {
	var _ok bool
	_ret := bindings.GetShadowRootAdoptedStyleSheets(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ObservableArray[CSSStyleSheet]{}.FromRef(_ret), _ok
}

// AdoptedStyleSheets sets the value of property "ShadowRoot.adoptedStyleSheets" to val.
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
// The returned bool will be false if there is no such property.
func (this ShadowRoot) PictureInPictureElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootPictureInPictureElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ActiveElement returns the value of property "ShadowRoot.activeElement".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) ActiveElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootActiveElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// PointerLockElement returns the value of property "ShadowRoot.pointerLockElement".
//
// The returned bool will be false if there is no such property.
func (this ShadowRoot) PointerLockElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetShadowRootPointerLockElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// GetAnimations calls the method "ShadowRoot.getAnimations".
//
// The returned bool will be false if there is no such method.
func (this ShadowRoot) GetAnimations() (js.Array[Animation], bool) {
	var _ok bool
	_ret := bindings.CallShadowRootGetAnimations(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Animation]{}.FromRef(_ret), _ok
}

// GetAnimationsFunc returns the method "ShadowRoot.getAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ShadowRoot) GetAnimationsFunc() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.ShadowRootGetAnimationsFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) AvailLeft() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedAvailLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// AvailTop returns the value of property "ScreenDetailed.availTop".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) AvailTop() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedAvailTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Left returns the value of property "ScreenDetailed.left".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) Left() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Top returns the value of property "ScreenDetailed.top".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) Top() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// IsPrimary returns the value of property "ScreenDetailed.isPrimary".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) IsPrimary() (bool, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedIsPrimary(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsInternal returns the value of property "ScreenDetailed.isInternal".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) IsInternal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedIsInternal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DevicePixelRatio returns the value of property "ScreenDetailed.devicePixelRatio".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) DevicePixelRatio() (float32, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedDevicePixelRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Label returns the value of property "ScreenDetailed.label".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetailed) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailedLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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

// ToString calls the method "CSSStyleValue.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleValue) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleValueToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSStyleValue.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleValue) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSStyleValueToStringFunc(
			this.Ref(),
		),
	)
}

// Parse calls the staticmethod "CSSStyleValue.parse".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleValue) Parse(property js.String, cssText js.String) (CSSStyleValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleValueParse(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		cssText.Ref(),
	)

	return CSSStyleValue{}.FromRef(_ret), _ok
}

// ParseFunc returns the staticmethod "CSSStyleValue.parse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleValue) ParseFunc() (fn js.Func[func(property js.String, cssText js.String) CSSStyleValue]) {
	return fn.FromRef(
		bindings.CSSStyleValueParseFunc(
			this.Ref(),
		),
	)
}

// ParseAll calls the staticmethod "CSSStyleValue.parseAll".
//
// The returned bool will be false if there is no such method.
func (this CSSStyleValue) ParseAll(property js.String, cssText js.String) (js.Array[CSSStyleValue], bool) {
	var _ok bool
	_ret := bindings.CallCSSStyleValueParseAll(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
		cssText.Ref(),
	)

	return js.Array[CSSStyleValue]{}.FromRef(_ret), _ok
}

// ParseAllFunc returns the staticmethod "CSSStyleValue.parseAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSStyleValue) ParseAllFunc() (fn js.Func[func(property js.String, cssText js.String) js.Array[CSSStyleValue]]) {
	return fn.FromRef(
		bindings.CSSStyleValueParseAllFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this StylePropertyMapReadOnly) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetStylePropertyMapReadOnlySize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "StylePropertyMapReadOnly.get".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMapReadOnly) Get(property js.String) (OneOf_undefined_CSSStyleValue, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapReadOnlyGet(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return OneOf_undefined_CSSStyleValue{}.FromRef(_ret), _ok
}

// GetFunc returns the method "StylePropertyMapReadOnly.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMapReadOnly) GetFunc() (fn js.Func[func(property js.String) OneOf_undefined_CSSStyleValue]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyGetFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "StylePropertyMapReadOnly.getAll".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMapReadOnly) GetAll(property js.String) (js.Array[CSSStyleValue], bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapReadOnlyGetAll(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return js.Array[CSSStyleValue]{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "StylePropertyMapReadOnly.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMapReadOnly) GetAllFunc() (fn js.Func[func(property js.String) js.Array[CSSStyleValue]]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyGetAllFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "StylePropertyMapReadOnly.has".
//
// The returned bool will be false if there is no such method.
func (this StylePropertyMapReadOnly) Has(property js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallStylePropertyMapReadOnlyHas(
		this.Ref(), js.Pointer(&_ok),
		property.Ref(),
	)

	return _ret == js.True, _ok
}

// HasFunc returns the method "StylePropertyMapReadOnly.has".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StylePropertyMapReadOnly) HasFunc() (fn js.Func[func(property js.String) bool]) {
	return fn.FromRef(
		bindings.StylePropertyMapReadOnlyHasFunc(
			this.Ref(),
		),
	)
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
	P1 DOMPointInit
	// P2 is "DOMQuadInit.p2"
	//
	// Optional
	P2 DOMPointInit
	// P3 is "DOMQuadInit.p3"
	//
	// Optional
	P3 DOMPointInit
	// P4 is "DOMQuadInit.p4"
	//
	// Optional
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

func NewDOMRect(x float64, y float64, width float64, height float64) DOMRect {
	return DOMRect{}.FromRef(
		bindings.NewDOMRectByDOMRect(
			float64(x),
			float64(y),
			float64(width),
			float64(height)),
	)
}

func NewDOMRectByDOMRect1(x float64, y float64, width float64) DOMRect {
	return DOMRect{}.FromRef(
		bindings.NewDOMRectByDOMRect1(
			float64(x),
			float64(y),
			float64(width)),
	)
}

func NewDOMRectByDOMRect2(x float64, y float64) DOMRect {
	return DOMRect{}.FromRef(
		bindings.NewDOMRectByDOMRect2(
			float64(x),
			float64(y)),
	)
}

func NewDOMRectByDOMRect3(x float64) DOMRect {
	return DOMRect{}.FromRef(
		bindings.NewDOMRectByDOMRect3(
			float64(x)),
	)
}

func NewDOMRectByDOMRect4() DOMRect {
	return DOMRect{}.FromRef(
		bindings.NewDOMRectByDOMRect4(),
	)
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
// The returned bool will be false if there is no such property.
func (this DOMRect) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// X sets the value of property "DOMRect.x" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMRect) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y sets the value of property "DOMRect.y" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMRect) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Width sets the value of property "DOMRect.width" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMRect) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height sets the value of property "DOMRect.height" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetHeight(val float64) bool {
	return js.True == bindings.SetDOMRectHeight(
		this.Ref(),
		float64(val),
	)
}

// FromRect calls the staticmethod "DOMRect.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMRect) FromRect(other DOMRectInit) (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectFromRect(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// FromRectFunc returns the staticmethod "DOMRect.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRect) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMRect.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMRect) FromRect1() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectFromRect1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// FromRect1Func returns the staticmethod "DOMRect.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRect) FromRect1Func() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.DOMRectFromRect1Func(
			this.Ref(),
		),
	)
}

func NewDOMPoint(x float64, y float64, z float64, w float64) DOMPoint {
	return DOMPoint{}.FromRef(
		bindings.NewDOMPointByDOMPoint(
			float64(x),
			float64(y),
			float64(z),
			float64(w)),
	)
}

func NewDOMPointByDOMPoint1(x float64, y float64, z float64) DOMPoint {
	return DOMPoint{}.FromRef(
		bindings.NewDOMPointByDOMPoint1(
			float64(x),
			float64(y),
			float64(z)),
	)
}

func NewDOMPointByDOMPoint2(x float64, y float64) DOMPoint {
	return DOMPoint{}.FromRef(
		bindings.NewDOMPointByDOMPoint2(
			float64(x),
			float64(y)),
	)
}

func NewDOMPointByDOMPoint3(x float64) DOMPoint {
	return DOMPoint{}.FromRef(
		bindings.NewDOMPointByDOMPoint3(
			float64(x)),
	)
}

func NewDOMPointByDOMPoint4() DOMPoint {
	return DOMPoint{}.FromRef(
		bindings.NewDOMPointByDOMPoint4(),
	)
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
// The returned bool will be false if there is no such property.
func (this DOMPoint) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// X sets the value of property "DOMPoint.x" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMPoint) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y sets the value of property "DOMPoint.y" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMPoint) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z sets the value of property "DOMPoint.z" to val.
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
// The returned bool will be false if there is no such property.
func (this DOMPoint) W() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMPointW(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// W sets the value of property "DOMPoint.w" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetW(val float64) bool {
	return js.True == bindings.SetDOMPointW(
		this.Ref(),
		float64(val),
	)
}

// FromPoint calls the staticmethod "DOMPoint.fromPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMPoint) FromPoint(other DOMPointInit) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointFromPoint(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// FromPointFunc returns the staticmethod "DOMPoint.fromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPoint) FromPointFunc() (fn js.Func[func(other DOMPointInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointFromPointFunc(
			this.Ref(),
		),
	)
}

// FromPoint1 calls the staticmethod "DOMPoint.fromPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMPoint) FromPoint1() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMPointFromPoint1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// FromPoint1Func returns the staticmethod "DOMPoint.fromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMPoint) FromPoint1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMPointFromPoint1Func(
			this.Ref(),
		),
	)
}

func NewDOMQuad(p1 DOMPointInit, p2 DOMPointInit, p3 DOMPointInit, p4 DOMPointInit) DOMQuad {
	return DOMQuad{}.FromRef(
		bindings.NewDOMQuadByDOMQuad(
			js.Pointer(&p1),
			js.Pointer(&p2),
			js.Pointer(&p3),
			js.Pointer(&p4)),
	)
}

func NewDOMQuadByDOMQuad1(p1 DOMPointInit, p2 DOMPointInit, p3 DOMPointInit) DOMQuad {
	return DOMQuad{}.FromRef(
		bindings.NewDOMQuadByDOMQuad1(
			js.Pointer(&p1),
			js.Pointer(&p2),
			js.Pointer(&p3)),
	)
}

func NewDOMQuadByDOMQuad2(p1 DOMPointInit, p2 DOMPointInit) DOMQuad {
	return DOMQuad{}.FromRef(
		bindings.NewDOMQuadByDOMQuad2(
			js.Pointer(&p1),
			js.Pointer(&p2)),
	)
}

func NewDOMQuadByDOMQuad3(p1 DOMPointInit) DOMQuad {
	return DOMQuad{}.FromRef(
		bindings.NewDOMQuadByDOMQuad3(
			js.Pointer(&p1)),
	)
}

func NewDOMQuadByDOMQuad4() DOMQuad {
	return DOMQuad{}.FromRef(
		bindings.NewDOMQuadByDOMQuad4(),
	)
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
// The returned bool will be false if there is no such property.
func (this DOMQuad) P1() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.GetDOMQuadP1(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPoint{}.FromRef(_ret), _ok
}

// P2 returns the value of property "DOMQuad.p2".
//
// The returned bool will be false if there is no such property.
func (this DOMQuad) P2() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.GetDOMQuadP2(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPoint{}.FromRef(_ret), _ok
}

// P3 returns the value of property "DOMQuad.p3".
//
// The returned bool will be false if there is no such property.
func (this DOMQuad) P3() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.GetDOMQuadP3(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPoint{}.FromRef(_ret), _ok
}

// P4 returns the value of property "DOMQuad.p4".
//
// The returned bool will be false if there is no such property.
func (this DOMQuad) P4() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.GetDOMQuadP4(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPoint{}.FromRef(_ret), _ok
}

// FromRect calls the staticmethod "DOMQuad.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) FromRect(other DOMRectInit) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadFromRect(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// FromRectFunc returns the staticmethod "DOMQuad.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMQuad.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) FromRect1() (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadFromRect1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// FromRect1Func returns the staticmethod "DOMQuad.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) FromRect1Func() (fn js.Func[func() DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromRect1Func(
			this.Ref(),
		),
	)
}

// FromQuad calls the staticmethod "DOMQuad.fromQuad".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) FromQuad(other DOMQuadInit) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadFromQuad(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// FromQuadFunc returns the staticmethod "DOMQuad.fromQuad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) FromQuadFunc() (fn js.Func[func(other DOMQuadInit) DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromQuadFunc(
			this.Ref(),
		),
	)
}

// FromQuad1 calls the staticmethod "DOMQuad.fromQuad".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) FromQuad1() (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadFromQuad1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// FromQuad1Func returns the staticmethod "DOMQuad.fromQuad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) FromQuad1Func() (fn js.Func[func() DOMQuad]) {
	return fn.FromRef(
		bindings.DOMQuadFromQuad1Func(
			this.Ref(),
		),
	)
}

// GetBounds calls the method "DOMQuad.getBounds".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) GetBounds() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadGetBounds(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetBoundsFunc returns the method "DOMQuad.getBounds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) GetBoundsFunc() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.DOMQuadGetBoundsFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMQuad.toJSON".
//
// The returned bool will be false if there is no such method.
func (this DOMQuad) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDOMQuadToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "DOMQuad.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMQuad) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMQuadToJSONFunc(
			this.Ref(),
		),
	)
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

func NewDOMRectReadOnly(x float64, y float64, width float64, height float64) DOMRectReadOnly {
	return DOMRectReadOnly{}.FromRef(
		bindings.NewDOMRectReadOnlyByDOMRectReadOnly(
			float64(x),
			float64(y),
			float64(width),
			float64(height)),
	)
}

func NewDOMRectReadOnlyByDOMRectReadOnly1(x float64, y float64, width float64) DOMRectReadOnly {
	return DOMRectReadOnly{}.FromRef(
		bindings.NewDOMRectReadOnlyByDOMRectReadOnly1(
			float64(x),
			float64(y),
			float64(width)),
	)
}

func NewDOMRectReadOnlyByDOMRectReadOnly2(x float64, y float64) DOMRectReadOnly {
	return DOMRectReadOnly{}.FromRef(
		bindings.NewDOMRectReadOnlyByDOMRectReadOnly2(
			float64(x),
			float64(y)),
	)
}

func NewDOMRectReadOnlyByDOMRectReadOnly3(x float64) DOMRectReadOnly {
	return DOMRectReadOnly{}.FromRef(
		bindings.NewDOMRectReadOnlyByDOMRectReadOnly3(
			float64(x)),
	)
}

func NewDOMRectReadOnlyByDOMRectReadOnly4() DOMRectReadOnly {
	return DOMRectReadOnly{}.FromRef(
		bindings.NewDOMRectReadOnlyByDOMRectReadOnly4(),
	)
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
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "DOMRectReadOnly.y".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Width returns the value of property "DOMRectReadOnly.width".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height returns the value of property "DOMRectReadOnly.height".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Top returns the value of property "DOMRectReadOnly.top".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Top() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Right returns the value of property "DOMRectReadOnly.right".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Right() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyRight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Bottom returns the value of property "DOMRectReadOnly.bottom".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Bottom() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyBottom(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Left returns the value of property "DOMRectReadOnly.left".
//
// The returned bool will be false if there is no such property.
func (this DOMRectReadOnly) Left() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMRectReadOnlyLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FromRect calls the staticmethod "DOMRectReadOnly.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMRectReadOnly) FromRect(other DOMRectInit) (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectReadOnlyFromRect(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// FromRectFunc returns the staticmethod "DOMRectReadOnly.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRectReadOnly) FromRectFunc() (fn js.Func[func(other DOMRectInit) DOMRectReadOnly]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyFromRectFunc(
			this.Ref(),
		),
	)
}

// FromRect1 calls the staticmethod "DOMRectReadOnly.fromRect".
//
// The returned bool will be false if there is no such method.
func (this DOMRectReadOnly) FromRect1() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectReadOnlyFromRect1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// FromRect1Func returns the staticmethod "DOMRectReadOnly.fromRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRectReadOnly) FromRect1Func() (fn js.Func[func() DOMRectReadOnly]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyFromRect1Func(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMRectReadOnly.toJSON".
//
// The returned bool will be false if there is no such method.
func (this DOMRectReadOnly) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDOMRectReadOnlyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "DOMRectReadOnly.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMRectReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMRectReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
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

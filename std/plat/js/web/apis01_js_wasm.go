// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *OptionalEffectTiming) UpdateFrom(ref js.Ref) {
	bindings.OptionalEffectTimingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OptionalEffectTiming) Update(ref js.Ref) {
	bindings.OptionalEffectTimingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OptionalEffectTiming) FreeMembers(recursive bool) {
	js.Free(
		p.Duration.Ref(),
		p.Easing.Ref(),
	)
	p.Duration = p.Duration.FromRef(js.Undefined)
	p.Easing = p.Easing.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "AnimationNodeList.length".
//
// It returns ok=false if there is no such property.
func (this AnimationNodeList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnimationNodeListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "AnimationNodeList.item" exists.
func (this AnimationNodeList) HasFuncItem() bool {
	return js.True == bindings.HasFuncAnimationNodeListItem(
		this.ref,
	)
}

// FuncItem returns the method "AnimationNodeList.item".
func (this AnimationNodeList) FuncItem() (fn js.Func[func(index uint32) AnimationEffect]) {
	bindings.FuncAnimationNodeListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "AnimationNodeList.item".
func (this AnimationNodeList) Item(index uint32) (ret AnimationEffect) {
	bindings.CallAnimationNodeListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "AnimationNodeList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationNodeList) TryItem(index uint32) (ret AnimationEffect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationNodeListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Children returns the value of property "GroupEffect.children".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) Children() (ret AnimationNodeList, ok bool) {
	ok = js.True == bindings.GetGroupEffectChildren(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstChild returns the value of property "GroupEffect.firstChild".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) FirstChild() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetGroupEffectFirstChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastChild returns the value of property "GroupEffect.lastChild".
//
// It returns ok=false if there is no such property.
func (this GroupEffect) LastChild() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetGroupEffectLastChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClone returns true if the method "GroupEffect.clone" exists.
func (this GroupEffect) HasFuncClone() bool {
	return js.True == bindings.HasFuncGroupEffectClone(
		this.ref,
	)
}

// FuncClone returns the method "GroupEffect.clone".
func (this GroupEffect) FuncClone() (fn js.Func[func() GroupEffect]) {
	bindings.FuncGroupEffectClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "GroupEffect.clone".
func (this GroupEffect) Clone() (ret GroupEffect) {
	bindings.CallGroupEffectClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "GroupEffect.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GroupEffect) TryClone() (ret GroupEffect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPrepend returns true if the method "GroupEffect.prepend" exists.
func (this GroupEffect) HasFuncPrepend() bool {
	return js.True == bindings.HasFuncGroupEffectPrepend(
		this.ref,
	)
}

// FuncPrepend returns the method "GroupEffect.prepend".
func (this GroupEffect) FuncPrepend() (fn js.Func[func(effects ...AnimationEffect)]) {
	bindings.FuncGroupEffectPrepend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prepend calls the method "GroupEffect.prepend".
func (this GroupEffect) Prepend(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallGroupEffectPrepend(
		this.ref, js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryPrepend calls the method "GroupEffect.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GroupEffect) TryPrepend(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectPrepend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasFuncAppend returns true if the method "GroupEffect.append" exists.
func (this GroupEffect) HasFuncAppend() bool {
	return js.True == bindings.HasFuncGroupEffectAppend(
		this.ref,
	)
}

// FuncAppend returns the method "GroupEffect.append".
func (this GroupEffect) FuncAppend() (fn js.Func[func(effects ...AnimationEffect)]) {
	bindings.FuncGroupEffectAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "GroupEffect.append".
func (this GroupEffect) Append(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallGroupEffectAppend(
		this.ref, js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryAppend calls the method "GroupEffect.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GroupEffect) TryAppend(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGroupEffectAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

type AnimationEffect struct {
	ref js.Ref
}

func (this AnimationEffect) Once() AnimationEffect {
	this.ref.Once()
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
	this.ref.Free()
}

// Parent returns the value of property "AnimationEffect.parent".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) Parent() (ret GroupEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "AnimationEffect.previousSibling".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) PreviousSibling() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectPreviousSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "AnimationEffect.nextSibling".
//
// It returns ok=false if there is no such property.
func (this AnimationEffect) NextSibling() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffectNextSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetTiming returns true if the method "AnimationEffect.getTiming" exists.
func (this AnimationEffect) HasFuncGetTiming() bool {
	return js.True == bindings.HasFuncAnimationEffectGetTiming(
		this.ref,
	)
}

// FuncGetTiming returns the method "AnimationEffect.getTiming".
func (this AnimationEffect) FuncGetTiming() (fn js.Func[func() EffectTiming]) {
	bindings.FuncAnimationEffectGetTiming(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTiming calls the method "AnimationEffect.getTiming".
func (this AnimationEffect) GetTiming() (ret EffectTiming) {
	bindings.CallAnimationEffectGetTiming(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetTiming calls the method "AnimationEffect.getTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryGetTiming() (ret EffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectGetTiming(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetComputedTiming returns true if the method "AnimationEffect.getComputedTiming" exists.
func (this AnimationEffect) HasFuncGetComputedTiming() bool {
	return js.True == bindings.HasFuncAnimationEffectGetComputedTiming(
		this.ref,
	)
}

// FuncGetComputedTiming returns the method "AnimationEffect.getComputedTiming".
func (this AnimationEffect) FuncGetComputedTiming() (fn js.Func[func() ComputedEffectTiming]) {
	bindings.FuncAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetComputedTiming calls the method "AnimationEffect.getComputedTiming".
func (this AnimationEffect) GetComputedTiming() (ret ComputedEffectTiming) {
	bindings.CallAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetComputedTiming calls the method "AnimationEffect.getComputedTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryGetComputedTiming() (ret ComputedEffectTiming, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectGetComputedTiming(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUpdateTiming returns true if the method "AnimationEffect.updateTiming" exists.
func (this AnimationEffect) HasFuncUpdateTiming() bool {
	return js.True == bindings.HasFuncAnimationEffectUpdateTiming(
		this.ref,
	)
}

// FuncUpdateTiming returns the method "AnimationEffect.updateTiming".
func (this AnimationEffect) FuncUpdateTiming() (fn js.Func[func(timing OptionalEffectTiming)]) {
	bindings.FuncAnimationEffectUpdateTiming(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateTiming calls the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTiming(timing OptionalEffectTiming) (ret js.Void) {
	bindings.CallAnimationEffectUpdateTiming(
		this.ref, js.Pointer(&ret),
		js.Pointer(&timing),
	)

	return
}

// TryUpdateTiming calls the method "AnimationEffect.updateTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryUpdateTiming(timing OptionalEffectTiming) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectUpdateTiming(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&timing),
	)

	return
}

// HasFuncUpdateTiming1 returns true if the method "AnimationEffect.updateTiming" exists.
func (this AnimationEffect) HasFuncUpdateTiming1() bool {
	return js.True == bindings.HasFuncAnimationEffectUpdateTiming1(
		this.ref,
	)
}

// FuncUpdateTiming1 returns the method "AnimationEffect.updateTiming".
func (this AnimationEffect) FuncUpdateTiming1() (fn js.Func[func()]) {
	bindings.FuncAnimationEffectUpdateTiming1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateTiming1 calls the method "AnimationEffect.updateTiming".
func (this AnimationEffect) UpdateTiming1() (ret js.Void) {
	bindings.CallAnimationEffectUpdateTiming1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUpdateTiming1 calls the method "AnimationEffect.updateTiming"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryUpdateTiming1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectUpdateTiming1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBefore returns true if the method "AnimationEffect.before" exists.
func (this AnimationEffect) HasFuncBefore() bool {
	return js.True == bindings.HasFuncAnimationEffectBefore(
		this.ref,
	)
}

// FuncBefore returns the method "AnimationEffect.before".
func (this AnimationEffect) FuncBefore() (fn js.Func[func(effects ...AnimationEffect)]) {
	bindings.FuncAnimationEffectBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Before calls the method "AnimationEffect.before".
func (this AnimationEffect) Before(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectBefore(
		this.ref, js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryBefore calls the method "AnimationEffect.before"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryBefore(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasFuncAfter returns true if the method "AnimationEffect.after" exists.
func (this AnimationEffect) HasFuncAfter() bool {
	return js.True == bindings.HasFuncAnimationEffectAfter(
		this.ref,
	)
}

// FuncAfter returns the method "AnimationEffect.after".
func (this AnimationEffect) FuncAfter() (fn js.Func[func(effects ...AnimationEffect)]) {
	bindings.FuncAnimationEffectAfter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// After calls the method "AnimationEffect.after".
func (this AnimationEffect) After(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectAfter(
		this.ref, js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryAfter calls the method "AnimationEffect.after"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryAfter(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectAfter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasFuncReplace returns true if the method "AnimationEffect.replace" exists.
func (this AnimationEffect) HasFuncReplace() bool {
	return js.True == bindings.HasFuncAnimationEffectReplace(
		this.ref,
	)
}

// FuncReplace returns the method "AnimationEffect.replace".
func (this AnimationEffect) FuncReplace() (fn js.Func[func(effects ...AnimationEffect)]) {
	bindings.FuncAnimationEffectReplace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Replace calls the method "AnimationEffect.replace".
func (this AnimationEffect) Replace(effects ...AnimationEffect) (ret js.Void) {
	bindings.CallAnimationEffectReplace(
		this.ref, js.Pointer(&ret),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// TryReplace calls the method "AnimationEffect.replace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryReplace(effects ...AnimationEffect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectReplace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(effects),
		js.SizeU(len(effects)),
	)

	return
}

// HasFuncRemove returns true if the method "AnimationEffect.remove" exists.
func (this AnimationEffect) HasFuncRemove() bool {
	return js.True == bindings.HasFuncAnimationEffectRemove(
		this.ref,
	)
}

// FuncRemove returns the method "AnimationEffect.remove".
func (this AnimationEffect) FuncRemove() (fn js.Func[func()]) {
	bindings.FuncAnimationEffectRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "AnimationEffect.remove".
func (this AnimationEffect) Remove() (ret js.Void) {
	bindings.CallAnimationEffectRemove(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "AnimationEffect.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationEffect) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationEffectRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AnimationTimeline struct {
	ref js.Ref
}

func (this AnimationTimeline) Once() AnimationTimeline {
	this.ref.Once()
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
	this.ref.Free()
}

// CurrentTime returns the value of property "AnimationTimeline.currentTime".
//
// It returns ok=false if there is no such property.
func (this AnimationTimeline) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationTimelineCurrentTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "AnimationTimeline.duration".
//
// It returns ok=false if there is no such property.
func (this AnimationTimeline) Duration() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationTimelineDuration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPlay returns true if the method "AnimationTimeline.play" exists.
func (this AnimationTimeline) HasFuncPlay() bool {
	return js.True == bindings.HasFuncAnimationTimelinePlay(
		this.ref,
	)
}

// FuncPlay returns the method "AnimationTimeline.play".
func (this AnimationTimeline) FuncPlay() (fn js.Func[func(effect AnimationEffect) Animation]) {
	bindings.FuncAnimationTimelinePlay(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Play calls the method "AnimationTimeline.play".
func (this AnimationTimeline) Play(effect AnimationEffect) (ret Animation) {
	bindings.CallAnimationTimelinePlay(
		this.ref, js.Pointer(&ret),
		effect.Ref(),
	)

	return
}

// TryPlay calls the method "AnimationTimeline.play"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationTimeline) TryPlay(effect AnimationEffect) (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationTimelinePlay(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		effect.Ref(),
	)

	return
}

// HasFuncPlay1 returns true if the method "AnimationTimeline.play" exists.
func (this AnimationTimeline) HasFuncPlay1() bool {
	return js.True == bindings.HasFuncAnimationTimelinePlay1(
		this.ref,
	)
}

// FuncPlay1 returns the method "AnimationTimeline.play".
func (this AnimationTimeline) FuncPlay1() (fn js.Func[func() Animation]) {
	bindings.FuncAnimationTimelinePlay1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Play1 calls the method "AnimationTimeline.play".
func (this AnimationTimeline) Play1() (ret Animation) {
	bindings.CallAnimationTimelinePlay1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPlay1 calls the method "AnimationTimeline.play"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationTimeline) TryPlay1() (ret Animation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationTimelinePlay1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "Animation.id".
//
// It returns ok=false if there is no such property.
func (this Animation) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "Animation.id" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetId(val js.String) bool {
	return js.True == bindings.SetAnimationId(
		this.ref,
		val.Ref(),
	)
}

// Effect returns the value of property "Animation.effect".
//
// It returns ok=false if there is no such property.
func (this Animation) Effect() (ret AnimationEffect, ok bool) {
	ok = js.True == bindings.GetAnimationEffect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEffect sets the value of property "Animation.effect" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetEffect(val AnimationEffect) bool {
	return js.True == bindings.SetAnimationEffect(
		this.ref,
		val.Ref(),
	)
}

// Timeline returns the value of property "Animation.timeline".
//
// It returns ok=false if there is no such property.
func (this Animation) Timeline() (ret AnimationTimeline, ok bool) {
	ok = js.True == bindings.GetAnimationTimeline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTimeline sets the value of property "Animation.timeline" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetTimeline(val AnimationTimeline) bool {
	return js.True == bindings.SetAnimationTimeline(
		this.ref,
		val.Ref(),
	)
}

// PlaybackRate returns the value of property "Animation.playbackRate".
//
// It returns ok=false if there is no such property.
func (this Animation) PlaybackRate() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPlaybackRate sets the value of property "Animation.playbackRate" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetPlaybackRate(val float64) bool {
	return js.True == bindings.SetAnimationPlaybackRate(
		this.ref,
		float64(val),
	)
}

// PlayState returns the value of property "Animation.playState".
//
// It returns ok=false if there is no such property.
func (this Animation) PlayState() (ret AnimationPlayState, ok bool) {
	ok = js.True == bindings.GetAnimationPlayState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReplaceState returns the value of property "Animation.replaceState".
//
// It returns ok=false if there is no such property.
func (this Animation) ReplaceState() (ret AnimationReplaceState, ok bool) {
	ok = js.True == bindings.GetAnimationReplaceState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Pending returns the value of property "Animation.pending".
//
// It returns ok=false if there is no such property.
func (this Animation) Pending() (ret bool, ok bool) {
	ok = js.True == bindings.GetAnimationPending(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "Animation.ready".
//
// It returns ok=false if there is no such property.
func (this Animation) Ready() (ret js.Promise[Animation], ok bool) {
	ok = js.True == bindings.GetAnimationReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "Animation.finished".
//
// It returns ok=false if there is no such property.
func (this Animation) Finished() (ret js.Promise[Animation], ok bool) {
	ok = js.True == bindings.GetAnimationFinished(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StartTime returns the value of property "Animation.startTime".
//
// It returns ok=false if there is no such property.
func (this Animation) StartTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationStartTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStartTime sets the value of property "Animation.startTime" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetStartTime(val CSSNumberish) bool {
	return js.True == bindings.SetAnimationStartTime(
		this.ref,
		val.Ref(),
	)
}

// CurrentTime returns the value of property "Animation.currentTime".
//
// It returns ok=false if there is no such property.
func (this Animation) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationCurrentTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCurrentTime sets the value of property "Animation.currentTime" to val.
//
// It returns false if the property cannot be set.
func (this Animation) SetCurrentTime(val CSSNumberish) bool {
	return js.True == bindings.SetAnimationCurrentTime(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCancel returns true if the method "Animation.cancel" exists.
func (this Animation) HasFuncCancel() bool {
	return js.True == bindings.HasFuncAnimationCancel(
		this.ref,
	)
}

// FuncCancel returns the method "Animation.cancel".
func (this Animation) FuncCancel() (fn js.Func[func()]) {
	bindings.FuncAnimationCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "Animation.cancel".
func (this Animation) Cancel() (ret js.Void) {
	bindings.CallAnimationCancel(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "Animation.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFinish returns true if the method "Animation.finish" exists.
func (this Animation) HasFuncFinish() bool {
	return js.True == bindings.HasFuncAnimationFinish(
		this.ref,
	)
}

// FuncFinish returns the method "Animation.finish".
func (this Animation) FuncFinish() (fn js.Func[func()]) {
	bindings.FuncAnimationFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "Animation.finish".
func (this Animation) Finish() (ret js.Void) {
	bindings.CallAnimationFinish(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "Animation.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPlay returns true if the method "Animation.play" exists.
func (this Animation) HasFuncPlay() bool {
	return js.True == bindings.HasFuncAnimationPlay(
		this.ref,
	)
}

// FuncPlay returns the method "Animation.play".
func (this Animation) FuncPlay() (fn js.Func[func()]) {
	bindings.FuncAnimationPlay(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Play calls the method "Animation.play".
func (this Animation) Play() (ret js.Void) {
	bindings.CallAnimationPlay(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPlay calls the method "Animation.play"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryPlay() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPlay(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPause returns true if the method "Animation.pause" exists.
func (this Animation) HasFuncPause() bool {
	return js.True == bindings.HasFuncAnimationPause(
		this.ref,
	)
}

// FuncPause returns the method "Animation.pause".
func (this Animation) FuncPause() (fn js.Func[func()]) {
	bindings.FuncAnimationPause(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pause calls the method "Animation.pause".
func (this Animation) Pause() (ret js.Void) {
	bindings.CallAnimationPause(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "Animation.pause"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPause(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUpdatePlaybackRate returns true if the method "Animation.updatePlaybackRate" exists.
func (this Animation) HasFuncUpdatePlaybackRate() bool {
	return js.True == bindings.HasFuncAnimationUpdatePlaybackRate(
		this.ref,
	)
}

// FuncUpdatePlaybackRate returns the method "Animation.updatePlaybackRate".
func (this Animation) FuncUpdatePlaybackRate() (fn js.Func[func(playbackRate float64)]) {
	bindings.FuncAnimationUpdatePlaybackRate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdatePlaybackRate calls the method "Animation.updatePlaybackRate".
func (this Animation) UpdatePlaybackRate(playbackRate float64) (ret js.Void) {
	bindings.CallAnimationUpdatePlaybackRate(
		this.ref, js.Pointer(&ret),
		float64(playbackRate),
	)

	return
}

// TryUpdatePlaybackRate calls the method "Animation.updatePlaybackRate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryUpdatePlaybackRate(playbackRate float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationUpdatePlaybackRate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(playbackRate),
	)

	return
}

// HasFuncReverse returns true if the method "Animation.reverse" exists.
func (this Animation) HasFuncReverse() bool {
	return js.True == bindings.HasFuncAnimationReverse(
		this.ref,
	)
}

// FuncReverse returns the method "Animation.reverse".
func (this Animation) FuncReverse() (fn js.Func[func()]) {
	bindings.FuncAnimationReverse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reverse calls the method "Animation.reverse".
func (this Animation) Reverse() (ret js.Void) {
	bindings.CallAnimationReverse(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReverse calls the method "Animation.reverse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryReverse() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationReverse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPersist returns true if the method "Animation.persist" exists.
func (this Animation) HasFuncPersist() bool {
	return js.True == bindings.HasFuncAnimationPersist(
		this.ref,
	)
}

// FuncPersist returns the method "Animation.persist".
func (this Animation) FuncPersist() (fn js.Func[func()]) {
	bindings.FuncAnimationPersist(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Persist calls the method "Animation.persist".
func (this Animation) Persist() (ret js.Void) {
	bindings.CallAnimationPersist(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPersist calls the method "Animation.persist"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryPersist() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationPersist(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCommitStyles returns true if the method "Animation.commitStyles" exists.
func (this Animation) HasFuncCommitStyles() bool {
	return js.True == bindings.HasFuncAnimationCommitStyles(
		this.ref,
	)
}

// FuncCommitStyles returns the method "Animation.commitStyles".
func (this Animation) FuncCommitStyles() (fn js.Func[func()]) {
	bindings.FuncAnimationCommitStyles(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CommitStyles calls the method "Animation.commitStyles".
func (this Animation) CommitStyles() (ret js.Void) {
	bindings.CallAnimationCommitStyles(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCommitStyles calls the method "Animation.commitStyles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Animation) TryCommitStyles() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationCommitStyles(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// MediaText returns the value of property "MediaList.mediaText".
//
// It returns ok=false if there is no such property.
func (this MediaList) MediaText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaListMediaText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMediaText sets the value of property "MediaList.mediaText" to val.
//
// It returns false if the property cannot be set.
func (this MediaList) SetMediaText(val js.String) bool {
	return js.True == bindings.SetMediaListMediaText(
		this.ref,
		val.Ref(),
	)
}

// Length returns the value of property "MediaList.length".
//
// It returns ok=false if there is no such property.
func (this MediaList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "MediaList.item" exists.
func (this MediaList) HasFuncItem() bool {
	return js.True == bindings.HasFuncMediaListItem(
		this.ref,
	)
}

// FuncItem returns the method "MediaList.item".
func (this MediaList) FuncItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncMediaListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "MediaList.item".
func (this MediaList) Item(index uint32) (ret js.String) {
	bindings.CallMediaListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "MediaList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendMedium returns true if the method "MediaList.appendMedium" exists.
func (this MediaList) HasFuncAppendMedium() bool {
	return js.True == bindings.HasFuncMediaListAppendMedium(
		this.ref,
	)
}

// FuncAppendMedium returns the method "MediaList.appendMedium".
func (this MediaList) FuncAppendMedium() (fn js.Func[func(medium js.String)]) {
	bindings.FuncMediaListAppendMedium(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendMedium calls the method "MediaList.appendMedium".
func (this MediaList) AppendMedium(medium js.String) (ret js.Void) {
	bindings.CallMediaListAppendMedium(
		this.ref, js.Pointer(&ret),
		medium.Ref(),
	)

	return
}

// TryAppendMedium calls the method "MediaList.appendMedium"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaList) TryAppendMedium(medium js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListAppendMedium(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		medium.Ref(),
	)

	return
}

// HasFuncDeleteMedium returns true if the method "MediaList.deleteMedium" exists.
func (this MediaList) HasFuncDeleteMedium() bool {
	return js.True == bindings.HasFuncMediaListDeleteMedium(
		this.ref,
	)
}

// FuncDeleteMedium returns the method "MediaList.deleteMedium".
func (this MediaList) FuncDeleteMedium() (fn js.Func[func(medium js.String)]) {
	bindings.FuncMediaListDeleteMedium(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteMedium calls the method "MediaList.deleteMedium".
func (this MediaList) DeleteMedium(medium js.String) (ret js.Void) {
	bindings.CallMediaListDeleteMedium(
		this.ref, js.Pointer(&ret),
		medium.Ref(),
	)

	return
}

// TryDeleteMedium calls the method "MediaList.deleteMedium"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaList) TryDeleteMedium(medium js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaListDeleteMedium(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CSSStyleSheetInit) UpdateFrom(ref js.Ref) {
	bindings.CSSStyleSheetInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CSSStyleSheetInit) Update(ref js.Ref) {
	bindings.CSSStyleSheetInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CSSStyleSheetInit) FreeMembers(recursive bool) {
	js.Free(
		p.BaseURL.Ref(),
		p.Media.Ref(),
	)
	p.BaseURL = p.BaseURL.FromRef(js.Undefined)
	p.Media = p.Media.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// CssText returns the value of property "CSSRule.cssText".
//
// It returns ok=false if there is no such property.
func (this CSSRule) CssText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSRuleCssText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCssText sets the value of property "CSSRule.cssText" to val.
//
// It returns false if the property cannot be set.
func (this CSSRule) SetCssText(val js.String) bool {
	return js.True == bindings.SetCSSRuleCssText(
		this.ref,
		val.Ref(),
	)
}

// ParentRule returns the value of property "CSSRule.parentRule".
//
// It returns ok=false if there is no such property.
func (this CSSRule) ParentRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSRuleParentRule(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ParentStyleSheet returns the value of property "CSSRule.parentStyleSheet".
//
// It returns ok=false if there is no such property.
func (this CSSRule) ParentStyleSheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetCSSRuleParentStyleSheet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "CSSRule.type".
//
// It returns ok=false if there is no such property.
func (this CSSRule) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCSSRuleType(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSRuleList struct {
	ref js.Ref
}

func (this CSSRuleList) Once() CSSRuleList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "CSSRuleList.length".
//
// It returns ok=false if there is no such property.
func (this CSSRuleList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSRuleListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "CSSRuleList.item" exists.
func (this CSSRuleList) HasFuncItem() bool {
	return js.True == bindings.HasFuncCSSRuleListItem(
		this.ref,
	)
}

// FuncItem returns the method "CSSRuleList.item".
func (this CSSRuleList) FuncItem() (fn js.Func[func(index uint32) CSSRule]) {
	bindings.FuncCSSRuleListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "CSSRuleList.item".
func (this CSSRuleList) Item(index uint32) (ret CSSRule) {
	bindings.CallCSSRuleListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "CSSRuleList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSRuleList) TryItem(index uint32) (ret CSSRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSRuleListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// OwnerRule returns the value of property "CSSStyleSheet.ownerRule".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) OwnerRule() (ret CSSRule, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetOwnerRule(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CssRules returns the value of property "CSSStyleSheet.cssRules".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetCssRules(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Rules returns the value of property "CSSStyleSheet.rules".
//
// It returns ok=false if there is no such property.
func (this CSSStyleSheet) Rules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSStyleSheetRules(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInsertRule returns true if the method "CSSStyleSheet.insertRule" exists.
func (this CSSStyleSheet) HasFuncInsertRule() bool {
	return js.True == bindings.HasFuncCSSStyleSheetInsertRule(
		this.ref,
	)
}

// FuncInsertRule returns the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) FuncInsertRule() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	bindings.FuncCSSStyleSheetInsertRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRule calls the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRule(rule js.String, index uint32) (ret uint32) {
	bindings.CallCSSStyleSheetInsertRule(
		this.ref, js.Pointer(&ret),
		rule.Ref(),
		uint32(index),
	)

	return
}

// TryInsertRule calls the method "CSSStyleSheet.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryInsertRule(rule js.String, index uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetInsertRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
		uint32(index),
	)

	return
}

// HasFuncInsertRule1 returns true if the method "CSSStyleSheet.insertRule" exists.
func (this CSSStyleSheet) HasFuncInsertRule1() bool {
	return js.True == bindings.HasFuncCSSStyleSheetInsertRule1(
		this.ref,
	)
}

// FuncInsertRule1 returns the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) FuncInsertRule1() (fn js.Func[func(rule js.String) uint32]) {
	bindings.FuncCSSStyleSheetInsertRule1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRule1 calls the method "CSSStyleSheet.insertRule".
func (this CSSStyleSheet) InsertRule1(rule js.String) (ret uint32) {
	bindings.CallCSSStyleSheetInsertRule1(
		this.ref, js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryInsertRule1 calls the method "CSSStyleSheet.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryInsertRule1(rule js.String) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetInsertRule1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasFuncDeleteRule returns true if the method "CSSStyleSheet.deleteRule" exists.
func (this CSSStyleSheet) HasFuncDeleteRule() bool {
	return js.True == bindings.HasFuncCSSStyleSheetDeleteRule(
		this.ref,
	)
}

// FuncDeleteRule returns the method "CSSStyleSheet.deleteRule".
func (this CSSStyleSheet) FuncDeleteRule() (fn js.Func[func(index uint32)]) {
	bindings.FuncCSSStyleSheetDeleteRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRule calls the method "CSSStyleSheet.deleteRule".
func (this CSSStyleSheet) DeleteRule(index uint32) (ret js.Void) {
	bindings.CallCSSStyleSheetDeleteRule(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDeleteRule calls the method "CSSStyleSheet.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryDeleteRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetDeleteRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncReplace returns true if the method "CSSStyleSheet.replace" exists.
func (this CSSStyleSheet) HasFuncReplace() bool {
	return js.True == bindings.HasFuncCSSStyleSheetReplace(
		this.ref,
	)
}

// FuncReplace returns the method "CSSStyleSheet.replace".
func (this CSSStyleSheet) FuncReplace() (fn js.Func[func(text js.String) js.Promise[CSSStyleSheet]]) {
	bindings.FuncCSSStyleSheetReplace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Replace calls the method "CSSStyleSheet.replace".
func (this CSSStyleSheet) Replace(text js.String) (ret js.Promise[CSSStyleSheet]) {
	bindings.CallCSSStyleSheetReplace(
		this.ref, js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryReplace calls the method "CSSStyleSheet.replace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryReplace(text js.String) (ret js.Promise[CSSStyleSheet], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetReplace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasFuncReplaceSync returns true if the method "CSSStyleSheet.replaceSync" exists.
func (this CSSStyleSheet) HasFuncReplaceSync() bool {
	return js.True == bindings.HasFuncCSSStyleSheetReplaceSync(
		this.ref,
	)
}

// FuncReplaceSync returns the method "CSSStyleSheet.replaceSync".
func (this CSSStyleSheet) FuncReplaceSync() (fn js.Func[func(text js.String)]) {
	bindings.FuncCSSStyleSheetReplaceSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceSync calls the method "CSSStyleSheet.replaceSync".
func (this CSSStyleSheet) ReplaceSync(text js.String) (ret js.Void) {
	bindings.CallCSSStyleSheetReplaceSync(
		this.ref, js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryReplaceSync calls the method "CSSStyleSheet.replaceSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryReplaceSync(text js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetReplaceSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasFuncAddRule returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasFuncAddRule() bool {
	return js.True == bindings.HasFuncCSSStyleSheetAddRule(
		this.ref,
	)
}

// FuncAddRule returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) FuncAddRule() (fn js.Func[func(selector js.String, style js.String, index uint32) int32]) {
	bindings.FuncCSSStyleSheetAddRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRule calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule(selector js.String, style js.String, index uint32) (ret int32) {
	bindings.CallCSSStyleSheetAddRule(
		this.ref, js.Pointer(&ret),
		selector.Ref(),
		style.Ref(),
		uint32(index),
	)

	return
}

// TryAddRule calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryAddRule(selector js.String, style js.String, index uint32) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
		style.Ref(),
		uint32(index),
	)

	return
}

// HasFuncAddRule1 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasFuncAddRule1() bool {
	return js.True == bindings.HasFuncCSSStyleSheetAddRule1(
		this.ref,
	)
}

// FuncAddRule1 returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) FuncAddRule1() (fn js.Func[func(selector js.String, style js.String) int32]) {
	bindings.FuncCSSStyleSheetAddRule1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRule1 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule1(selector js.String, style js.String) (ret int32) {
	bindings.CallCSSStyleSheetAddRule1(
		this.ref, js.Pointer(&ret),
		selector.Ref(),
		style.Ref(),
	)

	return
}

// TryAddRule1 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryAddRule1(selector js.String, style js.String) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
		style.Ref(),
	)

	return
}

// HasFuncAddRule2 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasFuncAddRule2() bool {
	return js.True == bindings.HasFuncCSSStyleSheetAddRule2(
		this.ref,
	)
}

// FuncAddRule2 returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) FuncAddRule2() (fn js.Func[func(selector js.String) int32]) {
	bindings.FuncCSSStyleSheetAddRule2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRule2 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule2(selector js.String) (ret int32) {
	bindings.CallCSSStyleSheetAddRule2(
		this.ref, js.Pointer(&ret),
		selector.Ref(),
	)

	return
}

// TryAddRule2 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryAddRule2(selector js.String) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selector.Ref(),
	)

	return
}

// HasFuncAddRule3 returns true if the method "CSSStyleSheet.addRule" exists.
func (this CSSStyleSheet) HasFuncAddRule3() bool {
	return js.True == bindings.HasFuncCSSStyleSheetAddRule3(
		this.ref,
	)
}

// FuncAddRule3 returns the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) FuncAddRule3() (fn js.Func[func() int32]) {
	bindings.FuncCSSStyleSheetAddRule3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRule3 calls the method "CSSStyleSheet.addRule".
func (this CSSStyleSheet) AddRule3() (ret int32) {
	bindings.CallCSSStyleSheetAddRule3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAddRule3 calls the method "CSSStyleSheet.addRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryAddRule3() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetAddRule3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveRule returns true if the method "CSSStyleSheet.removeRule" exists.
func (this CSSStyleSheet) HasFuncRemoveRule() bool {
	return js.True == bindings.HasFuncCSSStyleSheetRemoveRule(
		this.ref,
	)
}

// FuncRemoveRule returns the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) FuncRemoveRule() (fn js.Func[func(index uint32)]) {
	bindings.FuncCSSStyleSheetRemoveRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRule calls the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRule(index uint32) (ret js.Void) {
	bindings.CallCSSStyleSheetRemoveRule(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveRule calls the method "CSSStyleSheet.removeRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryRemoveRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetRemoveRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncRemoveRule1 returns true if the method "CSSStyleSheet.removeRule" exists.
func (this CSSStyleSheet) HasFuncRemoveRule1() bool {
	return js.True == bindings.HasFuncCSSStyleSheetRemoveRule1(
		this.ref,
	)
}

// FuncRemoveRule1 returns the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) FuncRemoveRule1() (fn js.Func[func()]) {
	bindings.FuncCSSStyleSheetRemoveRule1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRule1 calls the method "CSSStyleSheet.removeRule".
func (this CSSStyleSheet) RemoveRule1() (ret js.Void) {
	bindings.CallCSSStyleSheetRemoveRule1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemoveRule1 calls the method "CSSStyleSheet.removeRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleSheet) TryRemoveRule1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleSheetRemoveRule1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type StyleSheetList struct {
	ref js.Ref
}

func (this StyleSheetList) Once() StyleSheetList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "StyleSheetList.length".
//
// It returns ok=false if there is no such property.
func (this StyleSheetList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStyleSheetListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "StyleSheetList.item" exists.
func (this StyleSheetList) HasFuncItem() bool {
	return js.True == bindings.HasFuncStyleSheetListItem(
		this.ref,
	)
}

// FuncItem returns the method "StyleSheetList.item".
func (this StyleSheetList) FuncItem() (fn js.Func[func(index uint32) CSSStyleSheet]) {
	bindings.FuncStyleSheetListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "StyleSheetList.item".
func (this StyleSheetList) Item(index uint32) (ret CSSStyleSheet) {
	bindings.CallStyleSheetListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "StyleSheetList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StyleSheetList) TryItem(index uint32) (ret CSSStyleSheet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStyleSheetListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type ShadowRoot struct {
	DocumentFragment
}

func (this ShadowRoot) Once() ShadowRoot {
	this.ref.Once()
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
	this.ref.Free()
}

// Mode returns the value of property "ShadowRoot.mode".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) Mode() (ret ShadowRootMode, ok bool) {
	ok = js.True == bindings.GetShadowRootMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DelegatesFocus returns the value of property "ShadowRoot.delegatesFocus".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) DelegatesFocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetShadowRootDelegatesFocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SlotAssignment returns the value of property "ShadowRoot.slotAssignment".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) SlotAssignment() (ret SlotAssignmentMode, ok bool) {
	ok = js.True == bindings.GetShadowRootSlotAssignment(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Host returns the value of property "ShadowRoot.host".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) Host() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InnerHTML returns the value of property "ShadowRoot.innerHTML".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) InnerHTML() (ret js.String, ok bool) {
	ok = js.True == bindings.GetShadowRootInnerHTML(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInnerHTML sets the value of property "ShadowRoot.innerHTML" to val.
//
// It returns false if the property cannot be set.
func (this ShadowRoot) SetInnerHTML(val js.String) bool {
	return js.True == bindings.SetShadowRootInnerHTML(
		this.ref,
		val.Ref(),
	)
}

// FullscreenElement returns the value of property "ShadowRoot.fullscreenElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) FullscreenElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootFullscreenElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StyleSheets returns the value of property "ShadowRoot.styleSheets".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) StyleSheets() (ret StyleSheetList, ok bool) {
	ok = js.True == bindings.GetShadowRootStyleSheets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AdoptedStyleSheets returns the value of property "ShadowRoot.adoptedStyleSheets".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) AdoptedStyleSheets() (ret js.ObservableArray[CSSStyleSheet], ok bool) {
	ok = js.True == bindings.GetShadowRootAdoptedStyleSheets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAdoptedStyleSheets sets the value of property "ShadowRoot.adoptedStyleSheets" to val.
//
// It returns false if the property cannot be set.
func (this ShadowRoot) SetAdoptedStyleSheets(val js.ObservableArray[CSSStyleSheet]) bool {
	return js.True == bindings.SetShadowRootAdoptedStyleSheets(
		this.ref,
		val.Ref(),
	)
}

// PictureInPictureElement returns the value of property "ShadowRoot.pictureInPictureElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) PictureInPictureElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootPictureInPictureElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActiveElement returns the value of property "ShadowRoot.activeElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) ActiveElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootActiveElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointerLockElement returns the value of property "ShadowRoot.pointerLockElement".
//
// It returns ok=false if there is no such property.
func (this ShadowRoot) PointerLockElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetShadowRootPointerLockElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetAnimations returns true if the method "ShadowRoot.getAnimations" exists.
func (this ShadowRoot) HasFuncGetAnimations() bool {
	return js.True == bindings.HasFuncShadowRootGetAnimations(
		this.ref,
	)
}

// FuncGetAnimations returns the method "ShadowRoot.getAnimations".
func (this ShadowRoot) FuncGetAnimations() (fn js.Func[func() js.Array[Animation]]) {
	bindings.FuncShadowRootGetAnimations(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAnimations calls the method "ShadowRoot.getAnimations".
func (this ShadowRoot) GetAnimations() (ret js.Array[Animation]) {
	bindings.CallShadowRootGetAnimations(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAnimations calls the method "ShadowRoot.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ShadowRoot) TryGetAnimations() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShadowRootGetAnimations(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ShadowRootInit) UpdateFrom(ref js.Ref) {
	bindings.ShadowRootInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShadowRootInit) Update(ref js.Ref) {
	bindings.ShadowRootInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShadowRootInit) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// AvailLeft returns the value of property "ScreenDetailed.availLeft".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) AvailLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedAvailLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AvailTop returns the value of property "ScreenDetailed.availTop".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) AvailTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedAvailTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Left returns the value of property "ScreenDetailed.left".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Left() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "ScreenDetailed.top".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Top() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "ScreenDetailed.isPrimary".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenDetailedIsPrimary(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsInternal returns the value of property "ScreenDetailed.isInternal".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) IsInternal() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenDetailedIsInternal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DevicePixelRatio returns the value of property "ScreenDetailed.devicePixelRatio".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) DevicePixelRatio() (ret float32, ok bool) {
	ok = js.True == bindings.GetScreenDetailedDevicePixelRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "ScreenDetailed.label".
//
// It returns ok=false if there is no such property.
func (this ScreenDetailed) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetScreenDetailedLabel(
		this.ref, js.Pointer(&ret),
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
func (p *FullscreenOptions) UpdateFrom(ref js.Ref) {
	bindings.FullscreenOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FullscreenOptions) Update(ref js.Ref) {
	bindings.FullscreenOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FullscreenOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Screen.Ref(),
	)
	p.Screen = p.Screen.FromRef(js.Undefined)
}

type CSSStyleValue struct {
	ref js.Ref
}

func (this CSSStyleValue) Once() CSSStyleValue {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncToString returns true if the method "CSSStyleValue.toString" exists.
func (this CSSStyleValue) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSStyleValueToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSStyleValue.toString".
func (this CSSStyleValue) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSStyleValueToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSStyleValue.toString".
func (this CSSStyleValue) ToString() (ret js.String) {
	bindings.CallCSSStyleValueToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSStyleValue.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleValue) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncParse returns true if the static method "CSSStyleValue.parse" exists.
func (this CSSStyleValue) HasFuncParse() bool {
	return js.True == bindings.HasFuncCSSStyleValueParse(
		this.ref,
	)
}

// FuncParse returns the static method "CSSStyleValue.parse".
func (this CSSStyleValue) FuncParse() (fn js.Func[func(property js.String, cssText js.String) CSSStyleValue]) {
	bindings.FuncCSSStyleValueParse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Parse calls the static method "CSSStyleValue.parse".
func (this CSSStyleValue) Parse(property js.String, cssText js.String) (ret CSSStyleValue) {
	bindings.CallCSSStyleValueParse(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// TryParse calls the static method "CSSStyleValue.parse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleValue) TryParse(property js.String, cssText js.String) (ret CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueParse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// HasFuncParseAll returns true if the static method "CSSStyleValue.parseAll" exists.
func (this CSSStyleValue) HasFuncParseAll() bool {
	return js.True == bindings.HasFuncCSSStyleValueParseAll(
		this.ref,
	)
}

// FuncParseAll returns the static method "CSSStyleValue.parseAll".
func (this CSSStyleValue) FuncParseAll() (fn js.Func[func(property js.String, cssText js.String) js.Array[CSSStyleValue]]) {
	bindings.FuncCSSStyleValueParseAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ParseAll calls the static method "CSSStyleValue.parseAll".
func (this CSSStyleValue) ParseAll(property js.String, cssText js.String) (ret js.Array[CSSStyleValue]) {
	bindings.CallCSSStyleValueParseAll(
		this.ref, js.Pointer(&ret),
		property.Ref(),
		cssText.Ref(),
	)

	return
}

// TryParseAll calls the static method "CSSStyleValue.parseAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSStyleValue) TryParseAll(property js.String, cssText js.String) (ret js.Array[CSSStyleValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSStyleValueParseAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Size returns the value of property "StylePropertyMapReadOnly.size".
//
// It returns ok=false if there is no such property.
func (this StylePropertyMapReadOnly) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStylePropertyMapReadOnlySize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "StylePropertyMapReadOnly.get" exists.
func (this StylePropertyMapReadOnly) HasFuncGet() bool {
	return js.True == bindings.HasFuncStylePropertyMapReadOnlyGet(
		this.ref,
	)
}

// FuncGet returns the method "StylePropertyMapReadOnly.get".
func (this StylePropertyMapReadOnly) FuncGet() (fn js.Func[func(property js.String) OneOf_undefined_CSSStyleValue]) {
	bindings.FuncStylePropertyMapReadOnlyGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "StylePropertyMapReadOnly.get".
func (this StylePropertyMapReadOnly) Get(property js.String) (ret OneOf_undefined_CSSStyleValue) {
	bindings.CallStylePropertyMapReadOnlyGet(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGet calls the method "StylePropertyMapReadOnly.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMapReadOnly) TryGet(property js.String) (ret OneOf_undefined_CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the method "StylePropertyMapReadOnly.getAll" exists.
func (this StylePropertyMapReadOnly) HasFuncGetAll() bool {
	return js.True == bindings.HasFuncStylePropertyMapReadOnlyGetAll(
		this.ref,
	)
}

// FuncGetAll returns the method "StylePropertyMapReadOnly.getAll".
func (this StylePropertyMapReadOnly) FuncGetAll() (fn js.Func[func(property js.String) js.Array[CSSStyleValue]]) {
	bindings.FuncStylePropertyMapReadOnlyGetAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll calls the method "StylePropertyMapReadOnly.getAll".
func (this StylePropertyMapReadOnly) GetAll(property js.String) (ret js.Array[CSSStyleValue]) {
	bindings.CallStylePropertyMapReadOnlyGetAll(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryGetAll calls the method "StylePropertyMapReadOnly.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMapReadOnly) TryGetAll(property js.String) (ret js.Array[CSSStyleValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyGetAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		property.Ref(),
	)

	return
}

// HasFuncHas returns true if the method "StylePropertyMapReadOnly.has" exists.
func (this StylePropertyMapReadOnly) HasFuncHas() bool {
	return js.True == bindings.HasFuncStylePropertyMapReadOnlyHas(
		this.ref,
	)
}

// FuncHas returns the method "StylePropertyMapReadOnly.has".
func (this StylePropertyMapReadOnly) FuncHas() (fn js.Func[func(property js.String) bool]) {
	bindings.FuncStylePropertyMapReadOnlyHas(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Has calls the method "StylePropertyMapReadOnly.has".
func (this StylePropertyMapReadOnly) Has(property js.String) (ret bool) {
	bindings.CallStylePropertyMapReadOnlyHas(
		this.ref, js.Pointer(&ret),
		property.Ref(),
	)

	return
}

// TryHas calls the method "StylePropertyMapReadOnly.has"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StylePropertyMapReadOnly) TryHas(property js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStylePropertyMapReadOnlyHas(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *FocusableAreasOption) UpdateFrom(ref js.Ref) {
	bindings.FocusableAreasOptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FocusableAreasOption) Update(ref js.Ref) {
	bindings.FocusableAreasOptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FocusableAreasOption) FreeMembers(recursive bool) {
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
func (p *SpatialNavigationSearchOptions) UpdateFrom(ref js.Ref) {
	bindings.SpatialNavigationSearchOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpatialNavigationSearchOptions) Update(ref js.Ref) {
	bindings.SpatialNavigationSearchOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpatialNavigationSearchOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Candidates.Ref(),
		p.Container.Ref(),
	)
	p.Candidates = p.Candidates.FromRef(js.Undefined)
	p.Container = p.Container.FromRef(js.Undefined)
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
func (p *DOMPointInit) UpdateFrom(ref js.Ref) {
	bindings.DOMPointInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMPointInit) Update(ref js.Ref) {
	bindings.DOMPointInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMPointInit) FreeMembers(recursive bool) {
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
func (p *DOMRectInit) UpdateFrom(ref js.Ref) {
	bindings.DOMRectInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMRectInit) Update(ref js.Ref) {
	bindings.DOMRectInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMRectInit) FreeMembers(recursive bool) {
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
func (p *DOMQuadInit) UpdateFrom(ref js.Ref) {
	bindings.DOMQuadInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMQuadInit) Update(ref js.Ref) {
	bindings.DOMQuadInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMQuadInit) FreeMembers(recursive bool) {
	if recursive {
		p.P1.FreeMembers(true)
		p.P2.FreeMembers(true)
		p.P3.FreeMembers(true)
		p.P4.FreeMembers(true)
	}
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
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "DOMRect.x".
//
// It returns ok=false if there is no such property.
func (this DOMRect) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "DOMRect.x" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetX(val float64) bool {
	return js.True == bindings.SetDOMRectX(
		this.ref,
		float64(val),
	)
}

// Y returns the value of property "DOMRect.y".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "DOMRect.y" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetY(val float64) bool {
	return js.True == bindings.SetDOMRectY(
		this.ref,
		float64(val),
	)
}

// Width returns the value of property "DOMRect.width".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "DOMRect.width" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetWidth(val float64) bool {
	return js.True == bindings.SetDOMRectWidth(
		this.ref,
		float64(val),
	)
}

// Height returns the value of property "DOMRect.height".
//
// It returns ok=false if there is no such property.
func (this DOMRect) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "DOMRect.height" to val.
//
// It returns false if the property cannot be set.
func (this DOMRect) SetHeight(val float64) bool {
	return js.True == bindings.SetDOMRectHeight(
		this.ref,
		float64(val),
	)
}

// HasFuncFromRect returns true if the static method "DOMRect.fromRect" exists.
func (this DOMRect) HasFuncFromRect() bool {
	return js.True == bindings.HasFuncDOMRectFromRect(
		this.ref,
	)
}

// FuncFromRect returns the static method "DOMRect.fromRect".
func (this DOMRect) FuncFromRect() (fn js.Func[func(other DOMRectInit) DOMRect]) {
	bindings.FuncDOMRectFromRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect calls the static method "DOMRect.fromRect".
func (this DOMRect) FromRect(other DOMRectInit) (ret DOMRect) {
	bindings.CallDOMRectFromRect(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the static method "DOMRect.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRect) TryFromRect(other DOMRectInit) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectFromRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromRect1 returns true if the static method "DOMRect.fromRect" exists.
func (this DOMRect) HasFuncFromRect1() bool {
	return js.True == bindings.HasFuncDOMRectFromRect1(
		this.ref,
	)
}

// FuncFromRect1 returns the static method "DOMRect.fromRect".
func (this DOMRect) FuncFromRect1() (fn js.Func[func() DOMRect]) {
	bindings.FuncDOMRectFromRect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect1 calls the static method "DOMRect.fromRect".
func (this DOMRect) FromRect1() (ret DOMRect) {
	bindings.CallDOMRectFromRect1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the static method "DOMRect.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRect) TryFromRect1() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectFromRect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "DOMPoint.x".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "DOMPoint.x" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetX(val float64) bool {
	return js.True == bindings.SetDOMPointX(
		this.ref,
		float64(val),
	)
}

// Y returns the value of property "DOMPoint.y".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "DOMPoint.y" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetY(val float64) bool {
	return js.True == bindings.SetDOMPointY(
		this.ref,
		float64(val),
	)
}

// Z returns the value of property "DOMPoint.z".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetZ sets the value of property "DOMPoint.z" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetZ(val float64) bool {
	return js.True == bindings.SetDOMPointZ(
		this.ref,
		float64(val),
	)
}

// W returns the value of property "DOMPoint.w".
//
// It returns ok=false if there is no such property.
func (this DOMPoint) W() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMPointW(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetW sets the value of property "DOMPoint.w" to val.
//
// It returns false if the property cannot be set.
func (this DOMPoint) SetW(val float64) bool {
	return js.True == bindings.SetDOMPointW(
		this.ref,
		float64(val),
	)
}

// HasFuncFromPoint returns true if the static method "DOMPoint.fromPoint" exists.
func (this DOMPoint) HasFuncFromPoint() bool {
	return js.True == bindings.HasFuncDOMPointFromPoint(
		this.ref,
	)
}

// FuncFromPoint returns the static method "DOMPoint.fromPoint".
func (this DOMPoint) FuncFromPoint() (fn js.Func[func(other DOMPointInit) DOMPoint]) {
	bindings.FuncDOMPointFromPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromPoint calls the static method "DOMPoint.fromPoint".
func (this DOMPoint) FromPoint(other DOMPointInit) (ret DOMPoint) {
	bindings.CallDOMPointFromPoint(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromPoint calls the static method "DOMPoint.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPoint) TryFromPoint(other DOMPointInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointFromPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromPoint1 returns true if the static method "DOMPoint.fromPoint" exists.
func (this DOMPoint) HasFuncFromPoint1() bool {
	return js.True == bindings.HasFuncDOMPointFromPoint1(
		this.ref,
	)
}

// FuncFromPoint1 returns the static method "DOMPoint.fromPoint".
func (this DOMPoint) FuncFromPoint1() (fn js.Func[func() DOMPoint]) {
	bindings.FuncDOMPointFromPoint1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromPoint1 calls the static method "DOMPoint.fromPoint".
func (this DOMPoint) FromPoint1() (ret DOMPoint) {
	bindings.CallDOMPointFromPoint1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromPoint1 calls the static method "DOMPoint.fromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMPoint) TryFromPoint1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMPointFromPoint1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// P1 returns the value of property "DOMQuad.p1".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P1() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// P2 returns the value of property "DOMQuad.p2".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P2() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP2(
		this.ref, js.Pointer(&ret),
	)
	return
}

// P3 returns the value of property "DOMQuad.p3".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P3() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP3(
		this.ref, js.Pointer(&ret),
	)
	return
}

// P4 returns the value of property "DOMQuad.p4".
//
// It returns ok=false if there is no such property.
func (this DOMQuad) P4() (ret DOMPoint, ok bool) {
	ok = js.True == bindings.GetDOMQuadP4(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFromRect returns true if the static method "DOMQuad.fromRect" exists.
func (this DOMQuad) HasFuncFromRect() bool {
	return js.True == bindings.HasFuncDOMQuadFromRect(
		this.ref,
	)
}

// FuncFromRect returns the static method "DOMQuad.fromRect".
func (this DOMQuad) FuncFromRect() (fn js.Func[func(other DOMRectInit) DOMQuad]) {
	bindings.FuncDOMQuadFromRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect calls the static method "DOMQuad.fromRect".
func (this DOMQuad) FromRect(other DOMRectInit) (ret DOMQuad) {
	bindings.CallDOMQuadFromRect(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the static method "DOMQuad.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryFromRect(other DOMRectInit) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromRect1 returns true if the static method "DOMQuad.fromRect" exists.
func (this DOMQuad) HasFuncFromRect1() bool {
	return js.True == bindings.HasFuncDOMQuadFromRect1(
		this.ref,
	)
}

// FuncFromRect1 returns the static method "DOMQuad.fromRect".
func (this DOMQuad) FuncFromRect1() (fn js.Func[func() DOMQuad]) {
	bindings.FuncDOMQuadFromRect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect1 calls the static method "DOMQuad.fromRect".
func (this DOMQuad) FromRect1() (ret DOMQuad) {
	bindings.CallDOMQuadFromRect1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the static method "DOMQuad.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryFromRect1() (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromRect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFromQuad returns true if the static method "DOMQuad.fromQuad" exists.
func (this DOMQuad) HasFuncFromQuad() bool {
	return js.True == bindings.HasFuncDOMQuadFromQuad(
		this.ref,
	)
}

// FuncFromQuad returns the static method "DOMQuad.fromQuad".
func (this DOMQuad) FuncFromQuad() (fn js.Func[func(other DOMQuadInit) DOMQuad]) {
	bindings.FuncDOMQuadFromQuad(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromQuad calls the static method "DOMQuad.fromQuad".
func (this DOMQuad) FromQuad(other DOMQuadInit) (ret DOMQuad) {
	bindings.CallDOMQuadFromQuad(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromQuad calls the static method "DOMQuad.fromQuad"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryFromQuad(other DOMQuadInit) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromQuad(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromQuad1 returns true if the static method "DOMQuad.fromQuad" exists.
func (this DOMQuad) HasFuncFromQuad1() bool {
	return js.True == bindings.HasFuncDOMQuadFromQuad1(
		this.ref,
	)
}

// FuncFromQuad1 returns the static method "DOMQuad.fromQuad".
func (this DOMQuad) FuncFromQuad1() (fn js.Func[func() DOMQuad]) {
	bindings.FuncDOMQuadFromQuad1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromQuad1 calls the static method "DOMQuad.fromQuad".
func (this DOMQuad) FromQuad1() (ret DOMQuad) {
	bindings.CallDOMQuadFromQuad1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromQuad1 calls the static method "DOMQuad.fromQuad"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryFromQuad1() (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadFromQuad1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBounds returns true if the method "DOMQuad.getBounds" exists.
func (this DOMQuad) HasFuncGetBounds() bool {
	return js.True == bindings.HasFuncDOMQuadGetBounds(
		this.ref,
	)
}

// FuncGetBounds returns the method "DOMQuad.getBounds".
func (this DOMQuad) FuncGetBounds() (fn js.Func[func() DOMRect]) {
	bindings.FuncDOMQuadGetBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBounds calls the method "DOMQuad.getBounds".
func (this DOMQuad) GetBounds() (ret DOMRect) {
	bindings.CallDOMQuadGetBounds(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBounds calls the method "DOMQuad.getBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryGetBounds() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadGetBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "DOMQuad.toJSON" exists.
func (this DOMQuad) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncDOMQuadToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "DOMQuad.toJSON".
func (this DOMQuad) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncDOMQuadToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "DOMQuad.toJSON".
func (this DOMQuad) ToJSON() (ret js.Object) {
	bindings.CallDOMQuadToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMQuad.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMQuad) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMQuadToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ConvertCoordinateOptions) UpdateFrom(ref js.Ref) {
	bindings.ConvertCoordinateOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConvertCoordinateOptions) Update(ref js.Ref) {
	bindings.ConvertCoordinateOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConvertCoordinateOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "DOMRectReadOnly.x".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DOMRectReadOnly.y".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "DOMRectReadOnly.width".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "DOMRectReadOnly.height".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "DOMRectReadOnly.top".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Top() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Right returns the value of property "DOMRectReadOnly.right".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Right() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyRight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Bottom returns the value of property "DOMRectReadOnly.bottom".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Bottom() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyBottom(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Left returns the value of property "DOMRectReadOnly.left".
//
// It returns ok=false if there is no such property.
func (this DOMRectReadOnly) Left() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMRectReadOnlyLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFromRect returns true if the static method "DOMRectReadOnly.fromRect" exists.
func (this DOMRectReadOnly) HasFuncFromRect() bool {
	return js.True == bindings.HasFuncDOMRectReadOnlyFromRect(
		this.ref,
	)
}

// FuncFromRect returns the static method "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FuncFromRect() (fn js.Func[func(other DOMRectInit) DOMRectReadOnly]) {
	bindings.FuncDOMRectReadOnlyFromRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect calls the static method "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRect(other DOMRectInit) (ret DOMRectReadOnly) {
	bindings.CallDOMRectReadOnlyFromRect(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromRect calls the static method "DOMRectReadOnly.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRectReadOnly) TryFromRect(other DOMRectInit) (ret DOMRectReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyFromRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromRect1 returns true if the static method "DOMRectReadOnly.fromRect" exists.
func (this DOMRectReadOnly) HasFuncFromRect1() bool {
	return js.True == bindings.HasFuncDOMRectReadOnlyFromRect1(
		this.ref,
	)
}

// FuncFromRect1 returns the static method "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FuncFromRect1() (fn js.Func[func() DOMRectReadOnly]) {
	bindings.FuncDOMRectReadOnlyFromRect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromRect1 calls the static method "DOMRectReadOnly.fromRect".
func (this DOMRectReadOnly) FromRect1() (ret DOMRectReadOnly) {
	bindings.CallDOMRectReadOnlyFromRect1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromRect1 calls the static method "DOMRectReadOnly.fromRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRectReadOnly) TryFromRect1() (ret DOMRectReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyFromRect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "DOMRectReadOnly.toJSON" exists.
func (this DOMRectReadOnly) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncDOMRectReadOnlyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "DOMRectReadOnly.toJSON".
func (this DOMRectReadOnly) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncDOMRectReadOnlyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "DOMRectReadOnly.toJSON".
func (this DOMRectReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMRectReadOnlyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMRectReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMRectReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMRectReadOnlyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *AssignedNodesOptions) UpdateFrom(ref js.Ref) {
	bindings.AssignedNodesOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssignedNodesOptions) Update(ref js.Ref) {
	bindings.AssignedNodesOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssignedNodesOptions) FreeMembers(recursive bool) {
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

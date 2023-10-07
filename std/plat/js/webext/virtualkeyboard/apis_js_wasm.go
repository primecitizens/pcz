// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package virtualkeyboard

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/virtualkeyboard/bindings"
)

type FeatureRestrictions struct {
	// AutoCompleteEnabled is "FeatureRestrictions.autoCompleteEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoCompleteEnabled MUST be set to true to make this field effective.
	AutoCompleteEnabled bool
	// AutoCorrectEnabled is "FeatureRestrictions.autoCorrectEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoCorrectEnabled MUST be set to true to make this field effective.
	AutoCorrectEnabled bool
	// HandwritingEnabled is "FeatureRestrictions.handwritingEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_HandwritingEnabled MUST be set to true to make this field effective.
	HandwritingEnabled bool
	// SpellCheckEnabled is "FeatureRestrictions.spellCheckEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpellCheckEnabled MUST be set to true to make this field effective.
	SpellCheckEnabled bool
	// VoiceInputEnabled is "FeatureRestrictions.voiceInputEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_VoiceInputEnabled MUST be set to true to make this field effective.
	VoiceInputEnabled bool

	FFI_USE_AutoCompleteEnabled bool // for AutoCompleteEnabled.
	FFI_USE_AutoCorrectEnabled  bool // for AutoCorrectEnabled.
	FFI_USE_HandwritingEnabled  bool // for HandwritingEnabled.
	FFI_USE_SpellCheckEnabled   bool // for SpellCheckEnabled.
	FFI_USE_VoiceInputEnabled   bool // for VoiceInputEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FeatureRestrictions with all fields set.
func (p FeatureRestrictions) FromRef(ref js.Ref) FeatureRestrictions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FeatureRestrictions in the application heap.
func (p FeatureRestrictions) New() js.Ref {
	return bindings.FeatureRestrictionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FeatureRestrictions) UpdateFrom(ref js.Ref) {
	bindings.FeatureRestrictionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FeatureRestrictions) Update(ref js.Ref) {
	bindings.FeatureRestrictionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FeatureRestrictions) FreeMembers(recursive bool) {
}

type RestrictFeaturesCallbackFunc func(this js.Ref, update *FeatureRestrictions) js.Ref

func (fn RestrictFeaturesCallbackFunc) Register() js.Func[func(update *FeatureRestrictions)] {
	return js.RegisterCallback[func(update *FeatureRestrictions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RestrictFeaturesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FeatureRestrictions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RestrictFeaturesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, update *FeatureRestrictions) js.Ref
	Arg T
}

func (cb *RestrictFeaturesCallback[T]) Register() js.Func[func(update *FeatureRestrictions)] {
	return js.RegisterCallback[func(update *FeatureRestrictions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RestrictFeaturesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FeatureRestrictions
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncRestrictFeatures returns true if the function "WEBEXT.virtualKeyboard.restrictFeatures" exists.
func HasFuncRestrictFeatures() bool {
	return js.True == bindings.HasFuncRestrictFeatures()
}

// FuncRestrictFeatures returns the function "WEBEXT.virtualKeyboard.restrictFeatures".
func FuncRestrictFeatures() (fn js.Func[func(restrictions FeatureRestrictions) js.Promise[FeatureRestrictions]]) {
	bindings.FuncRestrictFeatures(
		js.Pointer(&fn),
	)
	return
}

// RestrictFeatures calls the function "WEBEXT.virtualKeyboard.restrictFeatures" directly.
func RestrictFeatures(restrictions FeatureRestrictions) (ret js.Promise[FeatureRestrictions]) {
	bindings.CallRestrictFeatures(
		js.Pointer(&ret),
		js.Pointer(&restrictions),
	)

	return
}

// TryRestrictFeatures calls the function "WEBEXT.virtualKeyboard.restrictFeatures"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestrictFeatures(restrictions FeatureRestrictions) (ret js.Promise[FeatureRestrictions], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestrictFeatures(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&restrictions),
	)

	return
}

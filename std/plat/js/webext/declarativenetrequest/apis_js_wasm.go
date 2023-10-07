// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package declarativenetrequest

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/declarativenetrequest/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
)

type Ruleset struct {
	// Id is "Ruleset.id"
	//
	// Optional
	Id js.String
	// Path is "Ruleset.path"
	//
	// Optional
	Path js.String
	// Enabled is "Ruleset.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool

	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Ruleset with all fields set.
func (p Ruleset) FromRef(ref js.Ref) Ruleset {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Ruleset in the application heap.
func (p Ruleset) New() js.Ref {
	return bindings.RulesetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Ruleset) UpdateFrom(ref js.Ref) {
	bindings.RulesetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Ruleset) Update(ref js.Ref) {
	bindings.RulesetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Ruleset) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Path.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
}

type DNRInfo struct {
	// RuleResources is "DNRInfo.rule_resources"
	//
	// Optional
	RuleResources js.Array[Ruleset]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DNRInfo with all fields set.
func (p DNRInfo) FromRef(ref js.Ref) DNRInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DNRInfo in the application heap.
func (p DNRInfo) New() js.Ref {
	return bindings.DNRInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DNRInfo) UpdateFrom(ref js.Ref) {
	bindings.DNRInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DNRInfo) Update(ref js.Ref) {
	bindings.DNRInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DNRInfo) FreeMembers(recursive bool) {
	js.Free(
		p.RuleResources.Ref(),
	)
	p.RuleResources = p.RuleResources.FromRef(js.Undefined)
}

type DomainType uint32

const (
	_ DomainType = iota

	DomainType_FIRST_PARTY
	DomainType_THIRD_PARTY
)

func (DomainType) FromRef(str js.Ref) DomainType {
	return DomainType(bindings.ConstOfDomainType(str))
}

func (x DomainType) String() (string, bool) {
	switch x {
	case DomainType_FIRST_PARTY:
		return "firstParty", true
	case DomainType_THIRD_PARTY:
		return "thirdParty", true
	default:
		return "", false
	}
}

type EmptyCallbackFunc func(this js.Ref) js.Ref

func (fn EmptyCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EmptyCallbackFunc) DispatchCallback(
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

type EmptyCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *EmptyCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EmptyCallback[T]) DispatchCallback(
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

type TabActionCountUpdate struct {
	// TabId is "TabActionCountUpdate.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// Increment is "TabActionCountUpdate.increment"
	//
	// Optional
	//
	// NOTE: FFI_USE_Increment MUST be set to true to make this field effective.
	Increment int32

	FFI_USE_TabId     bool // for TabId.
	FFI_USE_Increment bool // for Increment.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TabActionCountUpdate with all fields set.
func (p TabActionCountUpdate) FromRef(ref js.Ref) TabActionCountUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TabActionCountUpdate in the application heap.
func (p TabActionCountUpdate) New() js.Ref {
	return bindings.TabActionCountUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TabActionCountUpdate) UpdateFrom(ref js.Ref) {
	bindings.TabActionCountUpdateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TabActionCountUpdate) Update(ref js.Ref) {
	bindings.TabActionCountUpdateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TabActionCountUpdate) FreeMembers(recursive bool) {
}

type ExtensionActionOptions struct {
	// DisplayActionCountAsBadgeText is "ExtensionActionOptions.displayActionCountAsBadgeText"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayActionCountAsBadgeText MUST be set to true to make this field effective.
	DisplayActionCountAsBadgeText bool
	// TabUpdate is "ExtensionActionOptions.tabUpdate"
	//
	// Optional
	//
	// NOTE: TabUpdate.FFI_USE MUST be set to true to get TabUpdate used.
	TabUpdate TabActionCountUpdate

	FFI_USE_DisplayActionCountAsBadgeText bool // for DisplayActionCountAsBadgeText.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionActionOptions with all fields set.
func (p ExtensionActionOptions) FromRef(ref js.Ref) ExtensionActionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionActionOptions in the application heap.
func (p ExtensionActionOptions) New() js.Ref {
	return bindings.ExtensionActionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionActionOptions) UpdateFrom(ref js.Ref) {
	bindings.ExtensionActionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionActionOptions) Update(ref js.Ref) {
	bindings.ExtensionActionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionActionOptions) FreeMembers(recursive bool) {
	if recursive {
		p.TabUpdate.FreeMembers(true)
	}
}

type GetAllowedPagesCallbackFunc func(this js.Ref, result js.Array[js.String]) js.Ref

func (fn GetAllowedPagesCallbackFunc) Register() js.Func[func(result js.Array[js.String])] {
	return js.RegisterCallback[func(result js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllowedPagesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllowedPagesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetAllowedPagesCallback[T]) Register() js.Func[func(result js.Array[js.String])] {
	return js.RegisterCallback[func(result js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllowedPagesCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAvailableStaticRuleCountCallbackFunc func(this js.Ref, count int32) js.Ref

func (fn GetAvailableStaticRuleCountCallbackFunc) Register() js.Func[func(count int32)] {
	return js.RegisterCallback[func(count int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAvailableStaticRuleCountCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAvailableStaticRuleCountCallback[T any] struct {
	Fn  func(arg T, this js.Ref, count int32) js.Ref
	Arg T
}

func (cb *GetAvailableStaticRuleCountCallback[T]) Register() js.Func[func(count int32)] {
	return js.RegisterCallback[func(count int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAvailableStaticRuleCountCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisabledRuleIdsCallbackFunc func(this js.Ref, disabledRuleIds js.Array[int32]) js.Ref

func (fn GetDisabledRuleIdsCallbackFunc) Register() js.Func[func(disabledRuleIds js.Array[int32])] {
	return js.RegisterCallback[func(disabledRuleIds js.Array[int32])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDisabledRuleIdsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisabledRuleIdsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, disabledRuleIds js.Array[int32]) js.Ref
	Arg T
}

func (cb *GetDisabledRuleIdsCallback[T]) Register() js.Func[func(disabledRuleIds js.Array[int32])] {
	return js.RegisterCallback[func(disabledRuleIds js.Array[int32])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDisabledRuleIdsCallback[T]) DispatchCallback(
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

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisabledRuleIdsOptions struct {
	// RulesetId is "GetDisabledRuleIdsOptions.rulesetId"
	//
	// Optional
	RulesetId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDisabledRuleIdsOptions with all fields set.
func (p GetDisabledRuleIdsOptions) FromRef(ref js.Ref) GetDisabledRuleIdsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDisabledRuleIdsOptions in the application heap.
func (p GetDisabledRuleIdsOptions) New() js.Ref {
	return bindings.GetDisabledRuleIdsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDisabledRuleIdsOptions) UpdateFrom(ref js.Ref) {
	bindings.GetDisabledRuleIdsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDisabledRuleIdsOptions) Update(ref js.Ref) {
	bindings.GetDisabledRuleIdsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDisabledRuleIdsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.RulesetId.Ref(),
	)
	p.RulesetId = p.RulesetId.FromRef(js.Undefined)
}

type GetEnabledRulesetsCallbackFunc func(this js.Ref, rulesetIds js.Array[js.String]) js.Ref

func (fn GetEnabledRulesetsCallbackFunc) Register() js.Func[func(rulesetIds js.Array[js.String])] {
	return js.RegisterCallback[func(rulesetIds js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetEnabledRulesetsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetEnabledRulesetsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rulesetIds js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetEnabledRulesetsCallback[T]) Register() js.Func[func(rulesetIds js.Array[js.String])] {
	return js.RegisterCallback[func(rulesetIds js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetEnabledRulesetsCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetMatchedRulesCallbackFunc func(this js.Ref, details *RulesMatchedDetails) js.Ref

func (fn GetMatchedRulesCallbackFunc) Register() js.Func[func(details *RulesMatchedDetails)] {
	return js.RegisterCallback[func(details *RulesMatchedDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetMatchedRulesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RulesMatchedDetails
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

type GetMatchedRulesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *RulesMatchedDetails) js.Ref
	Arg T
}

func (cb *GetMatchedRulesCallback[T]) Register() js.Func[func(details *RulesMatchedDetails)] {
	return js.RegisterCallback[func(details *RulesMatchedDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetMatchedRulesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RulesMatchedDetails
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

type MatchedRule struct {
	// RuleId is "MatchedRule.ruleId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RuleId MUST be set to true to make this field effective.
	RuleId int32
	// RulesetId is "MatchedRule.rulesetId"
	//
	// Optional
	RulesetId js.String

	FFI_USE_RuleId bool // for RuleId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchedRule with all fields set.
func (p MatchedRule) FromRef(ref js.Ref) MatchedRule {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchedRule in the application heap.
func (p MatchedRule) New() js.Ref {
	return bindings.MatchedRuleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchedRule) UpdateFrom(ref js.Ref) {
	bindings.MatchedRuleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchedRule) Update(ref js.Ref) {
	bindings.MatchedRuleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchedRule) FreeMembers(recursive bool) {
	js.Free(
		p.RulesetId.Ref(),
	)
	p.RulesetId = p.RulesetId.FromRef(js.Undefined)
}

type MatchedRuleInfo struct {
	// Rule is "MatchedRuleInfo.rule"
	//
	// Optional
	//
	// NOTE: Rule.FFI_USE MUST be set to true to get Rule used.
	Rule MatchedRule
	// TimeStamp is "MatchedRuleInfo.timeStamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeStamp MUST be set to true to make this field effective.
	TimeStamp float64
	// TabId is "MatchedRuleInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_TimeStamp bool // for TimeStamp.
	FFI_USE_TabId     bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchedRuleInfo with all fields set.
func (p MatchedRuleInfo) FromRef(ref js.Ref) MatchedRuleInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchedRuleInfo in the application heap.
func (p MatchedRuleInfo) New() js.Ref {
	return bindings.MatchedRuleInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchedRuleInfo) UpdateFrom(ref js.Ref) {
	bindings.MatchedRuleInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchedRuleInfo) Update(ref js.Ref) {
	bindings.MatchedRuleInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchedRuleInfo) FreeMembers(recursive bool) {
	if recursive {
		p.Rule.FreeMembers(true)
	}
}

type RulesMatchedDetails struct {
	// RulesMatchedInfo is "RulesMatchedDetails.rulesMatchedInfo"
	//
	// Optional
	RulesMatchedInfo js.Array[MatchedRuleInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RulesMatchedDetails with all fields set.
func (p RulesMatchedDetails) FromRef(ref js.Ref) RulesMatchedDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RulesMatchedDetails in the application heap.
func (p RulesMatchedDetails) New() js.Ref {
	return bindings.RulesMatchedDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RulesMatchedDetails) UpdateFrom(ref js.Ref) {
	bindings.RulesMatchedDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RulesMatchedDetails) Update(ref js.Ref) {
	bindings.RulesMatchedDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RulesMatchedDetails) FreeMembers(recursive bool) {
	js.Free(
		p.RulesMatchedInfo.Ref(),
	)
	p.RulesMatchedInfo = p.RulesMatchedInfo.FromRef(js.Undefined)
}

type GetRulesCallbackFunc func(this js.Ref, rules js.Array[Rule]) js.Ref

func (fn GetRulesCallbackFunc) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRulesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRulesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rules js.Array[Rule]) js.Ref
	Arg T
}

func (cb *GetRulesCallback[T]) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRulesCallback[T]) DispatchCallback(
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

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResourceType uint32

const (
	_ ResourceType = iota

	ResourceType_MAIN_FRAME
	ResourceType_SUB_FRAME
	ResourceType_STYLESHEET
	ResourceType_SCRIPT
	ResourceType_IMAGE
	ResourceType_FONT
	ResourceType_OBJECT
	ResourceType_XMLHTTPREQUEST
	ResourceType_PING
	ResourceType_CSP_REPORT
	ResourceType_MEDIA
	ResourceType_WEBSOCKET
	ResourceType_WEBTRANSPORT
	ResourceType_WEBBUNDLE
	ResourceType_OTHER
)

func (ResourceType) FromRef(str js.Ref) ResourceType {
	return ResourceType(bindings.ConstOfResourceType(str))
}

func (x ResourceType) String() (string, bool) {
	switch x {
	case ResourceType_MAIN_FRAME:
		return "main_frame", true
	case ResourceType_SUB_FRAME:
		return "sub_frame", true
	case ResourceType_STYLESHEET:
		return "stylesheet", true
	case ResourceType_SCRIPT:
		return "script", true
	case ResourceType_IMAGE:
		return "image", true
	case ResourceType_FONT:
		return "font", true
	case ResourceType_OBJECT:
		return "object", true
	case ResourceType_XMLHTTPREQUEST:
		return "xmlhttprequest", true
	case ResourceType_PING:
		return "ping", true
	case ResourceType_CSP_REPORT:
		return "csp_report", true
	case ResourceType_MEDIA:
		return "media", true
	case ResourceType_WEBSOCKET:
		return "websocket", true
	case ResourceType_WEBTRANSPORT:
		return "webtransport", true
	case ResourceType_WEBBUNDLE:
		return "webbundle", true
	case ResourceType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type RequestMethod uint32

const (
	_ RequestMethod = iota

	RequestMethod_CONNECT
	RequestMethod_DELETE
	RequestMethod_GET
	RequestMethod_HEAD
	RequestMethod_OPTIONS
	RequestMethod_PATCH
	RequestMethod_POST
	RequestMethod_PUT
	RequestMethod_OTHER
)

func (RequestMethod) FromRef(str js.Ref) RequestMethod {
	return RequestMethod(bindings.ConstOfRequestMethod(str))
}

func (x RequestMethod) String() (string, bool) {
	switch x {
	case RequestMethod_CONNECT:
		return "connect", true
	case RequestMethod_DELETE:
		return "delete", true
	case RequestMethod_GET:
		return "get", true
	case RequestMethod_HEAD:
		return "head", true
	case RequestMethod_OPTIONS:
		return "options", true
	case RequestMethod_PATCH:
		return "patch", true
	case RequestMethod_POST:
		return "post", true
	case RequestMethod_PUT:
		return "put", true
	case RequestMethod_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type RuleCondition struct {
	// UrlFilter is "RuleCondition.urlFilter"
	//
	// Optional
	UrlFilter js.String
	// RegexFilter is "RuleCondition.regexFilter"
	//
	// Optional
	RegexFilter js.String
	// IsUrlFilterCaseSensitive is "RuleCondition.isUrlFilterCaseSensitive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUrlFilterCaseSensitive MUST be set to true to make this field effective.
	IsUrlFilterCaseSensitive bool
	// InitiatorDomains is "RuleCondition.initiatorDomains"
	//
	// Optional
	InitiatorDomains js.Array[js.String]
	// ExcludedInitiatorDomains is "RuleCondition.excludedInitiatorDomains"
	//
	// Optional
	ExcludedInitiatorDomains js.Array[js.String]
	// RequestDomains is "RuleCondition.requestDomains"
	//
	// Optional
	RequestDomains js.Array[js.String]
	// ExcludedRequestDomains is "RuleCondition.excludedRequestDomains"
	//
	// Optional
	ExcludedRequestDomains js.Array[js.String]
	// Domains is "RuleCondition.domains"
	//
	// Optional
	Domains js.Array[js.String]
	// ExcludedDomains is "RuleCondition.excludedDomains"
	//
	// Optional
	ExcludedDomains js.Array[js.String]
	// ResourceTypes is "RuleCondition.resourceTypes"
	//
	// Optional
	ResourceTypes js.Array[ResourceType]
	// ExcludedResourceTypes is "RuleCondition.excludedResourceTypes"
	//
	// Optional
	ExcludedResourceTypes js.Array[ResourceType]
	// RequestMethods is "RuleCondition.requestMethods"
	//
	// Optional
	RequestMethods js.Array[RequestMethod]
	// ExcludedRequestMethods is "RuleCondition.excludedRequestMethods"
	//
	// Optional
	ExcludedRequestMethods js.Array[RequestMethod]
	// DomainType is "RuleCondition.domainType"
	//
	// Optional
	DomainType DomainType
	// TabIds is "RuleCondition.tabIds"
	//
	// Optional
	TabIds js.Array[int32]
	// ExcludedTabIds is "RuleCondition.excludedTabIds"
	//
	// Optional
	ExcludedTabIds js.Array[int32]

	FFI_USE_IsUrlFilterCaseSensitive bool // for IsUrlFilterCaseSensitive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RuleCondition with all fields set.
func (p RuleCondition) FromRef(ref js.Ref) RuleCondition {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RuleCondition in the application heap.
func (p RuleCondition) New() js.Ref {
	return bindings.RuleConditionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RuleCondition) UpdateFrom(ref js.Ref) {
	bindings.RuleConditionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RuleCondition) Update(ref js.Ref) {
	bindings.RuleConditionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RuleCondition) FreeMembers(recursive bool) {
	js.Free(
		p.UrlFilter.Ref(),
		p.RegexFilter.Ref(),
		p.InitiatorDomains.Ref(),
		p.ExcludedInitiatorDomains.Ref(),
		p.RequestDomains.Ref(),
		p.ExcludedRequestDomains.Ref(),
		p.Domains.Ref(),
		p.ExcludedDomains.Ref(),
		p.ResourceTypes.Ref(),
		p.ExcludedResourceTypes.Ref(),
		p.RequestMethods.Ref(),
		p.ExcludedRequestMethods.Ref(),
		p.TabIds.Ref(),
		p.ExcludedTabIds.Ref(),
	)
	p.UrlFilter = p.UrlFilter.FromRef(js.Undefined)
	p.RegexFilter = p.RegexFilter.FromRef(js.Undefined)
	p.InitiatorDomains = p.InitiatorDomains.FromRef(js.Undefined)
	p.ExcludedInitiatorDomains = p.ExcludedInitiatorDomains.FromRef(js.Undefined)
	p.RequestDomains = p.RequestDomains.FromRef(js.Undefined)
	p.ExcludedRequestDomains = p.ExcludedRequestDomains.FromRef(js.Undefined)
	p.Domains = p.Domains.FromRef(js.Undefined)
	p.ExcludedDomains = p.ExcludedDomains.FromRef(js.Undefined)
	p.ResourceTypes = p.ResourceTypes.FromRef(js.Undefined)
	p.ExcludedResourceTypes = p.ExcludedResourceTypes.FromRef(js.Undefined)
	p.RequestMethods = p.RequestMethods.FromRef(js.Undefined)
	p.ExcludedRequestMethods = p.ExcludedRequestMethods.FromRef(js.Undefined)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
	p.ExcludedTabIds = p.ExcludedTabIds.FromRef(js.Undefined)
}

type RuleActionType uint32

const (
	_ RuleActionType = iota

	RuleActionType_BLOCK
	RuleActionType_REDIRECT
	RuleActionType_ALLOW
	RuleActionType_UPGRADE_SCHEME
	RuleActionType_MODIFY_HEADERS
	RuleActionType_ALLOW_ALL_REQUESTS
)

func (RuleActionType) FromRef(str js.Ref) RuleActionType {
	return RuleActionType(bindings.ConstOfRuleActionType(str))
}

func (x RuleActionType) String() (string, bool) {
	switch x {
	case RuleActionType_BLOCK:
		return "block", true
	case RuleActionType_REDIRECT:
		return "redirect", true
	case RuleActionType_ALLOW:
		return "allow", true
	case RuleActionType_UPGRADE_SCHEME:
		return "upgradeScheme", true
	case RuleActionType_MODIFY_HEADERS:
		return "modifyHeaders", true
	case RuleActionType_ALLOW_ALL_REQUESTS:
		return "allowAllRequests", true
	default:
		return "", false
	}
}

type QueryKeyValue struct {
	// Key is "QueryKeyValue.key"
	//
	// Optional
	Key js.String
	// Value is "QueryKeyValue.value"
	//
	// Optional
	Value js.String
	// ReplaceOnly is "QueryKeyValue.replaceOnly"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReplaceOnly MUST be set to true to make this field effective.
	ReplaceOnly bool

	FFI_USE_ReplaceOnly bool // for ReplaceOnly.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryKeyValue with all fields set.
func (p QueryKeyValue) FromRef(ref js.Ref) QueryKeyValue {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryKeyValue in the application heap.
func (p QueryKeyValue) New() js.Ref {
	return bindings.QueryKeyValueJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryKeyValue) UpdateFrom(ref js.Ref) {
	bindings.QueryKeyValueJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryKeyValue) Update(ref js.Ref) {
	bindings.QueryKeyValueJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryKeyValue) FreeMembers(recursive bool) {
	js.Free(
		p.Key.Ref(),
		p.Value.Ref(),
	)
	p.Key = p.Key.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type QueryTransform struct {
	// RemoveParams is "QueryTransform.removeParams"
	//
	// Optional
	RemoveParams js.Array[js.String]
	// AddOrReplaceParams is "QueryTransform.addOrReplaceParams"
	//
	// Optional
	AddOrReplaceParams js.Array[QueryKeyValue]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryTransform with all fields set.
func (p QueryTransform) FromRef(ref js.Ref) QueryTransform {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryTransform in the application heap.
func (p QueryTransform) New() js.Ref {
	return bindings.QueryTransformJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryTransform) UpdateFrom(ref js.Ref) {
	bindings.QueryTransformJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryTransform) Update(ref js.Ref) {
	bindings.QueryTransformJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryTransform) FreeMembers(recursive bool) {
	js.Free(
		p.RemoveParams.Ref(),
		p.AddOrReplaceParams.Ref(),
	)
	p.RemoveParams = p.RemoveParams.FromRef(js.Undefined)
	p.AddOrReplaceParams = p.AddOrReplaceParams.FromRef(js.Undefined)
}

type URLTransform struct {
	// Scheme is "URLTransform.scheme"
	//
	// Optional
	Scheme js.String
	// Host is "URLTransform.host"
	//
	// Optional
	Host js.String
	// Port is "URLTransform.port"
	//
	// Optional
	Port js.String
	// Path is "URLTransform.path"
	//
	// Optional
	Path js.String
	// Query is "URLTransform.query"
	//
	// Optional
	Query js.String
	// QueryTransform is "URLTransform.queryTransform"
	//
	// Optional
	//
	// NOTE: QueryTransform.FFI_USE MUST be set to true to get QueryTransform used.
	QueryTransform QueryTransform
	// Fragment is "URLTransform.fragment"
	//
	// Optional
	Fragment js.String
	// Username is "URLTransform.username"
	//
	// Optional
	Username js.String
	// Password is "URLTransform.password"
	//
	// Optional
	Password js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a URLTransform with all fields set.
func (p URLTransform) FromRef(ref js.Ref) URLTransform {
	p.UpdateFrom(ref)
	return p
}

// New creates a new URLTransform in the application heap.
func (p URLTransform) New() js.Ref {
	return bindings.URLTransformJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *URLTransform) UpdateFrom(ref js.Ref) {
	bindings.URLTransformJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *URLTransform) Update(ref js.Ref) {
	bindings.URLTransformJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *URLTransform) FreeMembers(recursive bool) {
	js.Free(
		p.Scheme.Ref(),
		p.Host.Ref(),
		p.Port.Ref(),
		p.Path.Ref(),
		p.Query.Ref(),
		p.Fragment.Ref(),
		p.Username.Ref(),
		p.Password.Ref(),
	)
	p.Scheme = p.Scheme.FromRef(js.Undefined)
	p.Host = p.Host.FromRef(js.Undefined)
	p.Port = p.Port.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Query = p.Query.FromRef(js.Undefined)
	p.Fragment = p.Fragment.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	if recursive {
		p.QueryTransform.FreeMembers(true)
	}
}

type Redirect struct {
	// ExtensionPath is "Redirect.extensionPath"
	//
	// Optional
	ExtensionPath js.String
	// Transform is "Redirect.transform"
	//
	// Optional
	//
	// NOTE: Transform.FFI_USE MUST be set to true to get Transform used.
	Transform URLTransform
	// Url is "Redirect.url"
	//
	// Optional
	Url js.String
	// RegexSubstitution is "Redirect.regexSubstitution"
	//
	// Optional
	RegexSubstitution js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Redirect with all fields set.
func (p Redirect) FromRef(ref js.Ref) Redirect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Redirect in the application heap.
func (p Redirect) New() js.Ref {
	return bindings.RedirectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Redirect) UpdateFrom(ref js.Ref) {
	bindings.RedirectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Redirect) Update(ref js.Ref) {
	bindings.RedirectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Redirect) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionPath.Ref(),
		p.Url.Ref(),
		p.RegexSubstitution.Ref(),
	)
	p.ExtensionPath = p.ExtensionPath.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.RegexSubstitution = p.RegexSubstitution.FromRef(js.Undefined)
	if recursive {
		p.Transform.FreeMembers(true)
	}
}

type HeaderOperation uint32

const (
	_ HeaderOperation = iota

	HeaderOperation_APPEND
	HeaderOperation_SET
	HeaderOperation_REMOVE
)

func (HeaderOperation) FromRef(str js.Ref) HeaderOperation {
	return HeaderOperation(bindings.ConstOfHeaderOperation(str))
}

func (x HeaderOperation) String() (string, bool) {
	switch x {
	case HeaderOperation_APPEND:
		return "append", true
	case HeaderOperation_SET:
		return "set", true
	case HeaderOperation_REMOVE:
		return "remove", true
	default:
		return "", false
	}
}

type ModifyHeaderInfo struct {
	// Header is "ModifyHeaderInfo.header"
	//
	// Optional
	Header js.String
	// Operation is "ModifyHeaderInfo.operation"
	//
	// Optional
	Operation HeaderOperation
	// Value is "ModifyHeaderInfo.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ModifyHeaderInfo with all fields set.
func (p ModifyHeaderInfo) FromRef(ref js.Ref) ModifyHeaderInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ModifyHeaderInfo in the application heap.
func (p ModifyHeaderInfo) New() js.Ref {
	return bindings.ModifyHeaderInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ModifyHeaderInfo) UpdateFrom(ref js.Ref) {
	bindings.ModifyHeaderInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ModifyHeaderInfo) Update(ref js.Ref) {
	bindings.ModifyHeaderInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ModifyHeaderInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Header.Ref(),
		p.Value.Ref(),
	)
	p.Header = p.Header.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type RuleAction struct {
	// Type is "RuleAction.type"
	//
	// Optional
	Type RuleActionType
	// Redirect is "RuleAction.redirect"
	//
	// Optional
	//
	// NOTE: Redirect.FFI_USE MUST be set to true to get Redirect used.
	Redirect Redirect
	// RequestHeaders is "RuleAction.requestHeaders"
	//
	// Optional
	RequestHeaders js.Array[ModifyHeaderInfo]
	// ResponseHeaders is "RuleAction.responseHeaders"
	//
	// Optional
	ResponseHeaders js.Array[ModifyHeaderInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RuleAction with all fields set.
func (p RuleAction) FromRef(ref js.Ref) RuleAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RuleAction in the application heap.
func (p RuleAction) New() js.Ref {
	return bindings.RuleActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RuleAction) UpdateFrom(ref js.Ref) {
	bindings.RuleActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RuleAction) Update(ref js.Ref) {
	bindings.RuleActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RuleAction) FreeMembers(recursive bool) {
	js.Free(
		p.RequestHeaders.Ref(),
		p.ResponseHeaders.Ref(),
	)
	p.RequestHeaders = p.RequestHeaders.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	if recursive {
		p.Redirect.FreeMembers(true)
	}
}

type Rule struct {
	// Id is "Rule.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Priority is "Rule.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// Condition is "Rule.condition"
	//
	// Optional
	//
	// NOTE: Condition.FFI_USE MUST be set to true to get Condition used.
	Condition RuleCondition
	// Action is "Rule.action"
	//
	// Optional
	//
	// NOTE: Action.FFI_USE MUST be set to true to get Action used.
	Action RuleAction

	FFI_USE_Id       bool // for Id.
	FFI_USE_Priority bool // for Priority.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Rule with all fields set.
func (p Rule) FromRef(ref js.Ref) Rule {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Rule in the application heap.
func (p Rule) New() js.Ref {
	return bindings.RuleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Rule) UpdateFrom(ref js.Ref) {
	bindings.RuleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Rule) Update(ref js.Ref) {
	bindings.RuleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Rule) FreeMembers(recursive bool) {
	if recursive {
		p.Condition.FreeMembers(true)
		p.Action.FreeMembers(true)
	}
}

type GetRulesFilter struct {
	// RuleIds is "GetRulesFilter.ruleIds"
	//
	// Optional
	RuleIds js.Array[int32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetRulesFilter with all fields set.
func (p GetRulesFilter) FromRef(ref js.Ref) GetRulesFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetRulesFilter in the application heap.
func (p GetRulesFilter) New() js.Ref {
	return bindings.GetRulesFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetRulesFilter) UpdateFrom(ref js.Ref) {
	bindings.GetRulesFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetRulesFilter) Update(ref js.Ref) {
	bindings.GetRulesFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetRulesFilter) FreeMembers(recursive bool) {
	js.Free(
		p.RuleIds.Ref(),
	)
	p.RuleIds = p.RuleIds.FromRef(js.Undefined)
}

type IsRegexSupportedCallbackFunc func(this js.Ref, result *IsRegexSupportedResult) js.Ref

func (fn IsRegexSupportedCallbackFunc) Register() js.Func[func(result *IsRegexSupportedResult)] {
	return js.RegisterCallback[func(result *IsRegexSupportedResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsRegexSupportedCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 IsRegexSupportedResult
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

type IsRegexSupportedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *IsRegexSupportedResult) js.Ref
	Arg T
}

func (cb *IsRegexSupportedCallback[T]) Register() js.Func[func(result *IsRegexSupportedResult)] {
	return js.RegisterCallback[func(result *IsRegexSupportedResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsRegexSupportedCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 IsRegexSupportedResult
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

type UnsupportedRegexReason uint32

const (
	_ UnsupportedRegexReason = iota

	UnsupportedRegexReason_SYNTAX_ERROR
	UnsupportedRegexReason_MEMORY_LIMIT_EXCEEDED
)

func (UnsupportedRegexReason) FromRef(str js.Ref) UnsupportedRegexReason {
	return UnsupportedRegexReason(bindings.ConstOfUnsupportedRegexReason(str))
}

func (x UnsupportedRegexReason) String() (string, bool) {
	switch x {
	case UnsupportedRegexReason_SYNTAX_ERROR:
		return "syntaxError", true
	case UnsupportedRegexReason_MEMORY_LIMIT_EXCEEDED:
		return "memoryLimitExceeded", true
	default:
		return "", false
	}
}

type IsRegexSupportedResult struct {
	// IsSupported is "IsRegexSupportedResult.isSupported"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsSupported MUST be set to true to make this field effective.
	IsSupported bool
	// Reason is "IsRegexSupportedResult.reason"
	//
	// Optional
	Reason UnsupportedRegexReason

	FFI_USE_IsSupported bool // for IsSupported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsRegexSupportedResult with all fields set.
func (p IsRegexSupportedResult) FromRef(ref js.Ref) IsRegexSupportedResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IsRegexSupportedResult in the application heap.
func (p IsRegexSupportedResult) New() js.Ref {
	return bindings.IsRegexSupportedResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IsRegexSupportedResult) UpdateFrom(ref js.Ref) {
	bindings.IsRegexSupportedResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IsRegexSupportedResult) Update(ref js.Ref) {
	bindings.IsRegexSupportedResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IsRegexSupportedResult) FreeMembers(recursive bool) {
}

type ManifestKeys struct {
	// DeclarativeNetRequest is "ManifestKeys.declarative_net_request"
	//
	// Optional
	//
	// NOTE: DeclarativeNetRequest.FFI_USE MUST be set to true to get DeclarativeNetRequest used.
	DeclarativeNetRequest DNRInfo

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManifestKeys with all fields set.
func (p ManifestKeys) FromRef(ref js.Ref) ManifestKeys {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManifestKeys in the application heap.
func (p ManifestKeys) New() js.Ref {
	return bindings.ManifestKeysJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManifestKeys) UpdateFrom(ref js.Ref) {
	bindings.ManifestKeysJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManifestKeys) Update(ref js.Ref) {
	bindings.ManifestKeysJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManifestKeys) FreeMembers(recursive bool) {
	if recursive {
		p.DeclarativeNetRequest.FreeMembers(true)
	}
}

type RequestDetails struct {
	// RequestId is "RequestDetails.requestId"
	//
	// Optional
	RequestId js.String
	// Url is "RequestDetails.url"
	//
	// Optional
	Url js.String
	// Initiator is "RequestDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "RequestDetails.method"
	//
	// Optional
	Method js.String
	// FrameId is "RequestDetails.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int32
	// DocumentId is "RequestDetails.documentId"
	//
	// Optional
	DocumentId js.String
	// FrameType is "RequestDetails.frameType"
	//
	// Optional
	FrameType extensiontypes.FrameType
	// DocumentLifecycle is "RequestDetails.documentLifecycle"
	//
	// Optional
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// ParentFrameId is "RequestDetails.parentFrameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ParentFrameId MUST be set to true to make this field effective.
	ParentFrameId int32
	// ParentDocumentId is "RequestDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// TabId is "RequestDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// Type is "RequestDetails.type"
	//
	// Optional
	Type ResourceType

	FFI_USE_FrameId       bool // for FrameId.
	FFI_USE_ParentFrameId bool // for ParentFrameId.
	FFI_USE_TabId         bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestDetails with all fields set.
func (p RequestDetails) FromRef(ref js.Ref) RequestDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestDetails in the application heap.
func (p RequestDetails) New() js.Ref {
	return bindings.RequestDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestDetails) UpdateFrom(ref js.Ref) {
	bindings.RequestDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestDetails) Update(ref js.Ref) {
	bindings.RequestDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestDetails) FreeMembers(recursive bool) {
	js.Free(
		p.RequestId.Ref(),
		p.Url.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.DocumentId.Ref(),
		p.ParentDocumentId.Ref(),
	)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
}

type MatchedRuleInfoDebug struct {
	// Rule is "MatchedRuleInfoDebug.rule"
	//
	// Optional
	//
	// NOTE: Rule.FFI_USE MUST be set to true to get Rule used.
	Rule MatchedRule
	// Request is "MatchedRuleInfoDebug.request"
	//
	// Optional
	//
	// NOTE: Request.FFI_USE MUST be set to true to get Request used.
	Request RequestDetails

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchedRuleInfoDebug with all fields set.
func (p MatchedRuleInfoDebug) FromRef(ref js.Ref) MatchedRuleInfoDebug {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchedRuleInfoDebug in the application heap.
func (p MatchedRuleInfoDebug) New() js.Ref {
	return bindings.MatchedRuleInfoDebugJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchedRuleInfoDebug) UpdateFrom(ref js.Ref) {
	bindings.MatchedRuleInfoDebugJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchedRuleInfoDebug) Update(ref js.Ref) {
	bindings.MatchedRuleInfoDebugJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchedRuleInfoDebug) FreeMembers(recursive bool) {
	if recursive {
		p.Rule.FreeMembers(true)
		p.Request.FreeMembers(true)
	}
}

type MatchedRulesFilter struct {
	// TabId is "MatchedRulesFilter.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// MinTimeStamp is "MatchedRulesFilter.minTimeStamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinTimeStamp MUST be set to true to make this field effective.
	MinTimeStamp float64

	FFI_USE_TabId        bool // for TabId.
	FFI_USE_MinTimeStamp bool // for MinTimeStamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchedRulesFilter with all fields set.
func (p MatchedRulesFilter) FromRef(ref js.Ref) MatchedRulesFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchedRulesFilter in the application heap.
func (p MatchedRulesFilter) New() js.Ref {
	return bindings.MatchedRulesFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchedRulesFilter) UpdateFrom(ref js.Ref) {
	bindings.MatchedRulesFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchedRulesFilter) Update(ref js.Ref) {
	bindings.MatchedRulesFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchedRulesFilter) FreeMembers(recursive bool) {
}

type Properties struct {
	ref js.Ref
}

func (this Properties) Once() Properties {
	this.ref.Once()
	return this
}

func (this Properties) Ref() js.Ref {
	return this.ref
}

func (this Properties) FromRef(ref js.Ref) Properties {
	this.ref = ref
	return this
}

func (this Properties) Free() {
	this.ref.Free()
}

// HasFuncGUARANTEED_MINIMUM_STATIC_RULES returns true if the static method "Properties.GUARANTEED_MINIMUM_STATIC_RULES" exists.
func (this Properties) HasFuncGUARANTEED_MINIMUM_STATIC_RULES() bool {
	return js.True == bindings.HasFuncPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
		this.ref,
	)
}

// FuncGUARANTEED_MINIMUM_STATIC_RULES returns the static method "Properties.GUARANTEED_MINIMUM_STATIC_RULES".
func (this Properties) FuncGUARANTEED_MINIMUM_STATIC_RULES() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GUARANTEED_MINIMUM_STATIC_RULES calls the static method "Properties.GUARANTEED_MINIMUM_STATIC_RULES".
func (this Properties) GUARANTEED_MINIMUM_STATIC_RULES() (ret int32) {
	bindings.CallPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGUARANTEED_MINIMUM_STATIC_RULES calls the static method "Properties.GUARANTEED_MINIMUM_STATIC_RULES"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryGUARANTEED_MINIMUM_STATIC_RULES() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesGUARANTEED_MINIMUM_STATIC_RULES(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_NUMBER_OF_DYNAMIC_RULES returns true if the static method "Properties.MAX_NUMBER_OF_DYNAMIC_RULES" exists.
func (this Properties) HasFuncMAX_NUMBER_OF_DYNAMIC_RULES() bool {
	return js.True == bindings.HasFuncPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
		this.ref,
	)
}

// FuncMAX_NUMBER_OF_DYNAMIC_RULES returns the static method "Properties.MAX_NUMBER_OF_DYNAMIC_RULES".
func (this Properties) FuncMAX_NUMBER_OF_DYNAMIC_RULES() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_NUMBER_OF_DYNAMIC_RULES calls the static method "Properties.MAX_NUMBER_OF_DYNAMIC_RULES".
func (this Properties) MAX_NUMBER_OF_DYNAMIC_RULES() (ret int32) {
	bindings.CallPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_NUMBER_OF_DYNAMIC_RULES calls the static method "Properties.MAX_NUMBER_OF_DYNAMIC_RULES"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_NUMBER_OF_DYNAMIC_RULES() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_NUMBER_OF_DYNAMIC_RULES(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES returns true if the static method "Properties.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES" exists.
func (this Properties) HasFuncMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES() bool {
	return js.True == bindings.HasFuncPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
		this.ref,
	)
}

// FuncMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES returns the static method "Properties.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES".
func (this Properties) FuncMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES calls the static method "Properties.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES".
func (this Properties) MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES() (ret int32) {
	bindings.CallPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES calls the static method "Properties.MAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_NUMBER_OF_DYNAMIC_AND_SESSION_RULES(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGETMATCHEDRULES_QUOTA_INTERVAL returns true if the static method "Properties.GETMATCHEDRULES_QUOTA_INTERVAL" exists.
func (this Properties) HasFuncGETMATCHEDRULES_QUOTA_INTERVAL() bool {
	return js.True == bindings.HasFuncPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
		this.ref,
	)
}

// FuncGETMATCHEDRULES_QUOTA_INTERVAL returns the static method "Properties.GETMATCHEDRULES_QUOTA_INTERVAL".
func (this Properties) FuncGETMATCHEDRULES_QUOTA_INTERVAL() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GETMATCHEDRULES_QUOTA_INTERVAL calls the static method "Properties.GETMATCHEDRULES_QUOTA_INTERVAL".
func (this Properties) GETMATCHEDRULES_QUOTA_INTERVAL() (ret int32) {
	bindings.CallPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGETMATCHEDRULES_QUOTA_INTERVAL calls the static method "Properties.GETMATCHEDRULES_QUOTA_INTERVAL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryGETMATCHEDRULES_QUOTA_INTERVAL() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesGETMATCHEDRULES_QUOTA_INTERVAL(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL returns true if the static method "Properties.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL" exists.
func (this Properties) HasFuncMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL() bool {
	return js.True == bindings.HasFuncPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
		this.ref,
	)
}

// FuncMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL returns the static method "Properties.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL".
func (this Properties) FuncMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL calls the static method "Properties.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL".
func (this Properties) MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL() (ret int32) {
	bindings.CallPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL calls the static method "Properties.MAX_GETMATCHEDRULES_CALLS_PER_INTERVAL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_GETMATCHEDRULES_CALLS_PER_INTERVAL(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_NUMBER_OF_REGEX_RULES returns true if the static method "Properties.MAX_NUMBER_OF_REGEX_RULES" exists.
func (this Properties) HasFuncMAX_NUMBER_OF_REGEX_RULES() bool {
	return js.True == bindings.HasFuncPropertiesMAX_NUMBER_OF_REGEX_RULES(
		this.ref,
	)
}

// FuncMAX_NUMBER_OF_REGEX_RULES returns the static method "Properties.MAX_NUMBER_OF_REGEX_RULES".
func (this Properties) FuncMAX_NUMBER_OF_REGEX_RULES() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_NUMBER_OF_REGEX_RULES(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_NUMBER_OF_REGEX_RULES calls the static method "Properties.MAX_NUMBER_OF_REGEX_RULES".
func (this Properties) MAX_NUMBER_OF_REGEX_RULES() (ret int32) {
	bindings.CallPropertiesMAX_NUMBER_OF_REGEX_RULES(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_NUMBER_OF_REGEX_RULES calls the static method "Properties.MAX_NUMBER_OF_REGEX_RULES"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_NUMBER_OF_REGEX_RULES() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_NUMBER_OF_REGEX_RULES(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_NUMBER_OF_STATIC_RULESETS returns true if the static method "Properties.MAX_NUMBER_OF_STATIC_RULESETS" exists.
func (this Properties) HasFuncMAX_NUMBER_OF_STATIC_RULESETS() bool {
	return js.True == bindings.HasFuncPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
		this.ref,
	)
}

// FuncMAX_NUMBER_OF_STATIC_RULESETS returns the static method "Properties.MAX_NUMBER_OF_STATIC_RULESETS".
func (this Properties) FuncMAX_NUMBER_OF_STATIC_RULESETS() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_NUMBER_OF_STATIC_RULESETS calls the static method "Properties.MAX_NUMBER_OF_STATIC_RULESETS".
func (this Properties) MAX_NUMBER_OF_STATIC_RULESETS() (ret int32) {
	bindings.CallPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_NUMBER_OF_STATIC_RULESETS calls the static method "Properties.MAX_NUMBER_OF_STATIC_RULESETS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_NUMBER_OF_STATIC_RULESETS() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_NUMBER_OF_STATIC_RULESETS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_NUMBER_OF_ENABLED_STATIC_RULESETS returns true if the static method "Properties.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS" exists.
func (this Properties) HasFuncMAX_NUMBER_OF_ENABLED_STATIC_RULESETS() bool {
	return js.True == bindings.HasFuncPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
		this.ref,
	)
}

// FuncMAX_NUMBER_OF_ENABLED_STATIC_RULESETS returns the static method "Properties.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS".
func (this Properties) FuncMAX_NUMBER_OF_ENABLED_STATIC_RULESETS() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_NUMBER_OF_ENABLED_STATIC_RULESETS calls the static method "Properties.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS".
func (this Properties) MAX_NUMBER_OF_ENABLED_STATIC_RULESETS() (ret int32) {
	bindings.CallPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_NUMBER_OF_ENABLED_STATIC_RULESETS calls the static method "Properties.MAX_NUMBER_OF_ENABLED_STATIC_RULESETS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_NUMBER_OF_ENABLED_STATIC_RULESETS() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_NUMBER_OF_ENABLED_STATIC_RULESETS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDYNAMIC_RULESET_ID returns true if the static method "Properties.DYNAMIC_RULESET_ID" exists.
func (this Properties) HasFuncDYNAMIC_RULESET_ID() bool {
	return js.True == bindings.HasFuncPropertiesDYNAMIC_RULESET_ID(
		this.ref,
	)
}

// FuncDYNAMIC_RULESET_ID returns the static method "Properties.DYNAMIC_RULESET_ID".
func (this Properties) FuncDYNAMIC_RULESET_ID() (fn js.Func[func() js.String]) {
	bindings.FuncPropertiesDYNAMIC_RULESET_ID(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DYNAMIC_RULESET_ID calls the static method "Properties.DYNAMIC_RULESET_ID".
func (this Properties) DYNAMIC_RULESET_ID() (ret js.String) {
	bindings.CallPropertiesDYNAMIC_RULESET_ID(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDYNAMIC_RULESET_ID calls the static method "Properties.DYNAMIC_RULESET_ID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryDYNAMIC_RULESET_ID() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesDYNAMIC_RULESET_ID(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSESSION_RULESET_ID returns true if the static method "Properties.SESSION_RULESET_ID" exists.
func (this Properties) HasFuncSESSION_RULESET_ID() bool {
	return js.True == bindings.HasFuncPropertiesSESSION_RULESET_ID(
		this.ref,
	)
}

// FuncSESSION_RULESET_ID returns the static method "Properties.SESSION_RULESET_ID".
func (this Properties) FuncSESSION_RULESET_ID() (fn js.Func[func() js.String]) {
	bindings.FuncPropertiesSESSION_RULESET_ID(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SESSION_RULESET_ID calls the static method "Properties.SESSION_RULESET_ID".
func (this Properties) SESSION_RULESET_ID() (ret js.String) {
	bindings.CallPropertiesSESSION_RULESET_ID(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySESSION_RULESET_ID calls the static method "Properties.SESSION_RULESET_ID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TrySESSION_RULESET_ID() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesSESSION_RULESET_ID(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RegexOptions struct {
	// Regex is "RegexOptions.regex"
	//
	// Optional
	Regex js.String
	// IsCaseSensitive is "RegexOptions.isCaseSensitive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsCaseSensitive MUST be set to true to make this field effective.
	IsCaseSensitive bool
	// RequireCapturing is "RegexOptions.requireCapturing"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequireCapturing MUST be set to true to make this field effective.
	RequireCapturing bool

	FFI_USE_IsCaseSensitive  bool // for IsCaseSensitive.
	FFI_USE_RequireCapturing bool // for RequireCapturing.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegexOptions with all fields set.
func (p RegexOptions) FromRef(ref js.Ref) RegexOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RegexOptions in the application heap.
func (p RegexOptions) New() js.Ref {
	return bindings.RegexOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegexOptions) UpdateFrom(ref js.Ref) {
	bindings.RegexOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegexOptions) Update(ref js.Ref) {
	bindings.RegexOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegexOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Regex.Ref(),
	)
	p.Regex = p.Regex.FromRef(js.Undefined)
}

type TestMatchOutcomeCallbackFunc func(this js.Ref, result *TestMatchOutcomeResult) js.Ref

func (fn TestMatchOutcomeCallbackFunc) Register() js.Func[func(result *TestMatchOutcomeResult)] {
	return js.RegisterCallback[func(result *TestMatchOutcomeResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TestMatchOutcomeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TestMatchOutcomeResult
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

type TestMatchOutcomeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *TestMatchOutcomeResult) js.Ref
	Arg T
}

func (cb *TestMatchOutcomeCallback[T]) Register() js.Func[func(result *TestMatchOutcomeResult)] {
	return js.RegisterCallback[func(result *TestMatchOutcomeResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TestMatchOutcomeCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TestMatchOutcomeResult
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

type TestMatchOutcomeResult struct {
	// MatchedRules is "TestMatchOutcomeResult.matchedRules"
	//
	// Optional
	MatchedRules js.Array[MatchedRule]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TestMatchOutcomeResult with all fields set.
func (p TestMatchOutcomeResult) FromRef(ref js.Ref) TestMatchOutcomeResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TestMatchOutcomeResult in the application heap.
func (p TestMatchOutcomeResult) New() js.Ref {
	return bindings.TestMatchOutcomeResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TestMatchOutcomeResult) UpdateFrom(ref js.Ref) {
	bindings.TestMatchOutcomeResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TestMatchOutcomeResult) Update(ref js.Ref) {
	bindings.TestMatchOutcomeResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TestMatchOutcomeResult) FreeMembers(recursive bool) {
	js.Free(
		p.MatchedRules.Ref(),
	)
	p.MatchedRules = p.MatchedRules.FromRef(js.Undefined)
}

type TestMatchRequestDetails struct {
	// Url is "TestMatchRequestDetails.url"
	//
	// Optional
	Url js.String
	// Initiator is "TestMatchRequestDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "TestMatchRequestDetails.method"
	//
	// Optional
	Method RequestMethod
	// Type is "TestMatchRequestDetails.type"
	//
	// Optional
	Type ResourceType
	// TabId is "TestMatchRequestDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TestMatchRequestDetails with all fields set.
func (p TestMatchRequestDetails) FromRef(ref js.Ref) TestMatchRequestDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TestMatchRequestDetails in the application heap.
func (p TestMatchRequestDetails) New() js.Ref {
	return bindings.TestMatchRequestDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TestMatchRequestDetails) UpdateFrom(ref js.Ref) {
	bindings.TestMatchRequestDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TestMatchRequestDetails) Update(ref js.Ref) {
	bindings.TestMatchRequestDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TestMatchRequestDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Initiator.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
}

type UpdateRuleOptions struct {
	// RemoveRuleIds is "UpdateRuleOptions.removeRuleIds"
	//
	// Optional
	RemoveRuleIds js.Array[int32]
	// AddRules is "UpdateRuleOptions.addRules"
	//
	// Optional
	AddRules js.Array[Rule]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateRuleOptions with all fields set.
func (p UpdateRuleOptions) FromRef(ref js.Ref) UpdateRuleOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateRuleOptions in the application heap.
func (p UpdateRuleOptions) New() js.Ref {
	return bindings.UpdateRuleOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateRuleOptions) UpdateFrom(ref js.Ref) {
	bindings.UpdateRuleOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateRuleOptions) Update(ref js.Ref) {
	bindings.UpdateRuleOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateRuleOptions) FreeMembers(recursive bool) {
	js.Free(
		p.RemoveRuleIds.Ref(),
		p.AddRules.Ref(),
	)
	p.RemoveRuleIds = p.RemoveRuleIds.FromRef(js.Undefined)
	p.AddRules = p.AddRules.FromRef(js.Undefined)
}

type UpdateRulesetOptions struct {
	// DisableRulesetIds is "UpdateRulesetOptions.disableRulesetIds"
	//
	// Optional
	DisableRulesetIds js.Array[js.String]
	// EnableRulesetIds is "UpdateRulesetOptions.enableRulesetIds"
	//
	// Optional
	EnableRulesetIds js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateRulesetOptions with all fields set.
func (p UpdateRulesetOptions) FromRef(ref js.Ref) UpdateRulesetOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateRulesetOptions in the application heap.
func (p UpdateRulesetOptions) New() js.Ref {
	return bindings.UpdateRulesetOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateRulesetOptions) UpdateFrom(ref js.Ref) {
	bindings.UpdateRulesetOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateRulesetOptions) Update(ref js.Ref) {
	bindings.UpdateRulesetOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateRulesetOptions) FreeMembers(recursive bool) {
	js.Free(
		p.DisableRulesetIds.Ref(),
		p.EnableRulesetIds.Ref(),
	)
	p.DisableRulesetIds = p.DisableRulesetIds.FromRef(js.Undefined)
	p.EnableRulesetIds = p.EnableRulesetIds.FromRef(js.Undefined)
}

type UpdateStaticRulesOptions struct {
	// RulesetId is "UpdateStaticRulesOptions.rulesetId"
	//
	// Optional
	RulesetId js.String
	// DisableRuleIds is "UpdateStaticRulesOptions.disableRuleIds"
	//
	// Optional
	DisableRuleIds js.Array[int32]
	// EnableRuleIds is "UpdateStaticRulesOptions.enableRuleIds"
	//
	// Optional
	EnableRuleIds js.Array[int32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateStaticRulesOptions with all fields set.
func (p UpdateStaticRulesOptions) FromRef(ref js.Ref) UpdateStaticRulesOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateStaticRulesOptions in the application heap.
func (p UpdateStaticRulesOptions) New() js.Ref {
	return bindings.UpdateStaticRulesOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateStaticRulesOptions) UpdateFrom(ref js.Ref) {
	bindings.UpdateStaticRulesOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateStaticRulesOptions) Update(ref js.Ref) {
	bindings.UpdateStaticRulesOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateStaticRulesOptions) FreeMembers(recursive bool) {
	js.Free(
		p.RulesetId.Ref(),
		p.DisableRuleIds.Ref(),
		p.EnableRuleIds.Ref(),
	)
	p.RulesetId = p.RulesetId.FromRef(js.Undefined)
	p.DisableRuleIds = p.DisableRuleIds.FromRef(js.Undefined)
	p.EnableRuleIds = p.EnableRuleIds.FromRef(js.Undefined)
}

// HasFuncGetAvailableStaticRuleCount returns true if the function "WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount" exists.
func HasFuncGetAvailableStaticRuleCount() bool {
	return js.True == bindings.HasFuncGetAvailableStaticRuleCount()
}

// FuncGetAvailableStaticRuleCount returns the function "WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount".
func FuncGetAvailableStaticRuleCount() (fn js.Func[func() js.Promise[js.Number[int32]]]) {
	bindings.FuncGetAvailableStaticRuleCount(
		js.Pointer(&fn),
	)
	return
}

// GetAvailableStaticRuleCount calls the function "WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount" directly.
func GetAvailableStaticRuleCount() (ret js.Promise[js.Number[int32]]) {
	bindings.CallGetAvailableStaticRuleCount(
		js.Pointer(&ret),
	)

	return
}

// TryGetAvailableStaticRuleCount calls the function "WEBEXT.declarativeNetRequest.getAvailableStaticRuleCount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAvailableStaticRuleCount() (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAvailableStaticRuleCount(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDisabledRuleIds returns true if the function "WEBEXT.declarativeNetRequest.getDisabledRuleIds" exists.
func HasFuncGetDisabledRuleIds() bool {
	return js.True == bindings.HasFuncGetDisabledRuleIds()
}

// FuncGetDisabledRuleIds returns the function "WEBEXT.declarativeNetRequest.getDisabledRuleIds".
func FuncGetDisabledRuleIds() (fn js.Func[func(options GetDisabledRuleIdsOptions) js.Promise[js.Array[int32]]]) {
	bindings.FuncGetDisabledRuleIds(
		js.Pointer(&fn),
	)
	return
}

// GetDisabledRuleIds calls the function "WEBEXT.declarativeNetRequest.getDisabledRuleIds" directly.
func GetDisabledRuleIds(options GetDisabledRuleIdsOptions) (ret js.Promise[js.Array[int32]]) {
	bindings.CallGetDisabledRuleIds(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetDisabledRuleIds calls the function "WEBEXT.declarativeNetRequest.getDisabledRuleIds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisabledRuleIds(options GetDisabledRuleIdsOptions) (ret js.Promise[js.Array[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisabledRuleIds(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetDynamicRules returns true if the function "WEBEXT.declarativeNetRequest.getDynamicRules" exists.
func HasFuncGetDynamicRules() bool {
	return js.True == bindings.HasFuncGetDynamicRules()
}

// FuncGetDynamicRules returns the function "WEBEXT.declarativeNetRequest.getDynamicRules".
func FuncGetDynamicRules() (fn js.Func[func(filter GetRulesFilter) js.Promise[js.Array[Rule]]]) {
	bindings.FuncGetDynamicRules(
		js.Pointer(&fn),
	)
	return
}

// GetDynamicRules calls the function "WEBEXT.declarativeNetRequest.getDynamicRules" directly.
func GetDynamicRules(filter GetRulesFilter) (ret js.Promise[js.Array[Rule]]) {
	bindings.CallGetDynamicRules(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetDynamicRules calls the function "WEBEXT.declarativeNetRequest.getDynamicRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDynamicRules(filter GetRulesFilter) (ret js.Promise[js.Array[Rule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDynamicRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncGetEnabledRulesets returns true if the function "WEBEXT.declarativeNetRequest.getEnabledRulesets" exists.
func HasFuncGetEnabledRulesets() bool {
	return js.True == bindings.HasFuncGetEnabledRulesets()
}

// FuncGetEnabledRulesets returns the function "WEBEXT.declarativeNetRequest.getEnabledRulesets".
func FuncGetEnabledRulesets() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetEnabledRulesets(
		js.Pointer(&fn),
	)
	return
}

// GetEnabledRulesets calls the function "WEBEXT.declarativeNetRequest.getEnabledRulesets" directly.
func GetEnabledRulesets() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetEnabledRulesets(
		js.Pointer(&ret),
	)

	return
}

// TryGetEnabledRulesets calls the function "WEBEXT.declarativeNetRequest.getEnabledRulesets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetEnabledRulesets() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetEnabledRulesets(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMatchedRules returns true if the function "WEBEXT.declarativeNetRequest.getMatchedRules" exists.
func HasFuncGetMatchedRules() bool {
	return js.True == bindings.HasFuncGetMatchedRules()
}

// FuncGetMatchedRules returns the function "WEBEXT.declarativeNetRequest.getMatchedRules".
func FuncGetMatchedRules() (fn js.Func[func(filter MatchedRulesFilter) js.Promise[RulesMatchedDetails]]) {
	bindings.FuncGetMatchedRules(
		js.Pointer(&fn),
	)
	return
}

// GetMatchedRules calls the function "WEBEXT.declarativeNetRequest.getMatchedRules" directly.
func GetMatchedRules(filter MatchedRulesFilter) (ret js.Promise[RulesMatchedDetails]) {
	bindings.CallGetMatchedRules(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetMatchedRules calls the function "WEBEXT.declarativeNetRequest.getMatchedRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMatchedRules(filter MatchedRulesFilter) (ret js.Promise[RulesMatchedDetails], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMatchedRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncGetSessionRules returns true if the function "WEBEXT.declarativeNetRequest.getSessionRules" exists.
func HasFuncGetSessionRules() bool {
	return js.True == bindings.HasFuncGetSessionRules()
}

// FuncGetSessionRules returns the function "WEBEXT.declarativeNetRequest.getSessionRules".
func FuncGetSessionRules() (fn js.Func[func(filter GetRulesFilter) js.Promise[js.Array[Rule]]]) {
	bindings.FuncGetSessionRules(
		js.Pointer(&fn),
	)
	return
}

// GetSessionRules calls the function "WEBEXT.declarativeNetRequest.getSessionRules" directly.
func GetSessionRules(filter GetRulesFilter) (ret js.Promise[js.Array[Rule]]) {
	bindings.CallGetSessionRules(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetSessionRules calls the function "WEBEXT.declarativeNetRequest.getSessionRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSessionRules(filter GetRulesFilter) (ret js.Promise[js.Array[Rule]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSessionRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncIsRegexSupported returns true if the function "WEBEXT.declarativeNetRequest.isRegexSupported" exists.
func HasFuncIsRegexSupported() bool {
	return js.True == bindings.HasFuncIsRegexSupported()
}

// FuncIsRegexSupported returns the function "WEBEXT.declarativeNetRequest.isRegexSupported".
func FuncIsRegexSupported() (fn js.Func[func(regexOptions RegexOptions) js.Promise[IsRegexSupportedResult]]) {
	bindings.FuncIsRegexSupported(
		js.Pointer(&fn),
	)
	return
}

// IsRegexSupported calls the function "WEBEXT.declarativeNetRequest.isRegexSupported" directly.
func IsRegexSupported(regexOptions RegexOptions) (ret js.Promise[IsRegexSupportedResult]) {
	bindings.CallIsRegexSupported(
		js.Pointer(&ret),
		js.Pointer(&regexOptions),
	)

	return
}

// TryIsRegexSupported calls the function "WEBEXT.declarativeNetRequest.isRegexSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsRegexSupported(regexOptions RegexOptions) (ret js.Promise[IsRegexSupportedResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsRegexSupported(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&regexOptions),
	)

	return
}

type OnRuleMatchedDebugEventCallbackFunc func(this js.Ref, info *MatchedRuleInfoDebug) js.Ref

func (fn OnRuleMatchedDebugEventCallbackFunc) Register() js.Func[func(info *MatchedRuleInfoDebug)] {
	return js.RegisterCallback[func(info *MatchedRuleInfoDebug)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRuleMatchedDebugEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MatchedRuleInfoDebug
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

type OnRuleMatchedDebugEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *MatchedRuleInfoDebug) js.Ref
	Arg T
}

func (cb *OnRuleMatchedDebugEventCallback[T]) Register() js.Func[func(info *MatchedRuleInfoDebug)] {
	return js.RegisterCallback[func(info *MatchedRuleInfoDebug)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRuleMatchedDebugEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MatchedRuleInfoDebug
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

// HasFuncOnRuleMatchedDebug returns true if the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener" exists.
func HasFuncOnRuleMatchedDebug() bool {
	return js.True == bindings.HasFuncOnRuleMatchedDebug()
}

// FuncOnRuleMatchedDebug returns the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener".
func FuncOnRuleMatchedDebug() (fn js.Func[func(callback js.Func[func(info *MatchedRuleInfoDebug)])]) {
	bindings.FuncOnRuleMatchedDebug(
		js.Pointer(&fn),
	)
	return
}

// OnRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener" directly.
func OnRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret js.Void) {
	bindings.CallOnRuleMatchedDebug(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRuleMatchedDebug(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRuleMatchedDebug returns true if the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener" exists.
func HasFuncOffRuleMatchedDebug() bool {
	return js.True == bindings.HasFuncOffRuleMatchedDebug()
}

// FuncOffRuleMatchedDebug returns the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener".
func FuncOffRuleMatchedDebug() (fn js.Func[func(callback js.Func[func(info *MatchedRuleInfoDebug)])]) {
	bindings.FuncOffRuleMatchedDebug(
		js.Pointer(&fn),
	)
	return
}

// OffRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener" directly.
func OffRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret js.Void) {
	bindings.CallOffRuleMatchedDebug(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRuleMatchedDebug(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRuleMatchedDebug returns true if the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener" exists.
func HasFuncHasOnRuleMatchedDebug() bool {
	return js.True == bindings.HasFuncHasOnRuleMatchedDebug()
}

// FuncHasOnRuleMatchedDebug returns the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener".
func FuncHasOnRuleMatchedDebug() (fn js.Func[func(callback js.Func[func(info *MatchedRuleInfoDebug)]) bool]) {
	bindings.FuncHasOnRuleMatchedDebug(
		js.Pointer(&fn),
	)
	return
}

// HasOnRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener" directly.
func HasOnRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret bool) {
	bindings.CallHasOnRuleMatchedDebug(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRuleMatchedDebug calls the function "WEBEXT.declarativeNetRequest.onRuleMatchedDebug.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRuleMatchedDebug(callback js.Func[func(info *MatchedRuleInfoDebug)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRuleMatchedDebug(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetExtensionActionOptions returns true if the function "WEBEXT.declarativeNetRequest.setExtensionActionOptions" exists.
func HasFuncSetExtensionActionOptions() bool {
	return js.True == bindings.HasFuncSetExtensionActionOptions()
}

// FuncSetExtensionActionOptions returns the function "WEBEXT.declarativeNetRequest.setExtensionActionOptions".
func FuncSetExtensionActionOptions() (fn js.Func[func(options ExtensionActionOptions) js.Promise[js.Void]]) {
	bindings.FuncSetExtensionActionOptions(
		js.Pointer(&fn),
	)
	return
}

// SetExtensionActionOptions calls the function "WEBEXT.declarativeNetRequest.setExtensionActionOptions" directly.
func SetExtensionActionOptions(options ExtensionActionOptions) (ret js.Promise[js.Void]) {
	bindings.CallSetExtensionActionOptions(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySetExtensionActionOptions calls the function "WEBEXT.declarativeNetRequest.setExtensionActionOptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetExtensionActionOptions(options ExtensionActionOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetExtensionActionOptions(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncTestMatchOutcome returns true if the function "WEBEXT.declarativeNetRequest.testMatchOutcome" exists.
func HasFuncTestMatchOutcome() bool {
	return js.True == bindings.HasFuncTestMatchOutcome()
}

// FuncTestMatchOutcome returns the function "WEBEXT.declarativeNetRequest.testMatchOutcome".
func FuncTestMatchOutcome() (fn js.Func[func(request TestMatchRequestDetails) js.Promise[TestMatchOutcomeResult]]) {
	bindings.FuncTestMatchOutcome(
		js.Pointer(&fn),
	)
	return
}

// TestMatchOutcome calls the function "WEBEXT.declarativeNetRequest.testMatchOutcome" directly.
func TestMatchOutcome(request TestMatchRequestDetails) (ret js.Promise[TestMatchOutcomeResult]) {
	bindings.CallTestMatchOutcome(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryTestMatchOutcome calls the function "WEBEXT.declarativeNetRequest.testMatchOutcome"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTestMatchOutcome(request TestMatchRequestDetails) (ret js.Promise[TestMatchOutcomeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTestMatchOutcome(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncUpdateDynamicRules returns true if the function "WEBEXT.declarativeNetRequest.updateDynamicRules" exists.
func HasFuncUpdateDynamicRules() bool {
	return js.True == bindings.HasFuncUpdateDynamicRules()
}

// FuncUpdateDynamicRules returns the function "WEBEXT.declarativeNetRequest.updateDynamicRules".
func FuncUpdateDynamicRules() (fn js.Func[func(options UpdateRuleOptions) js.Promise[js.Void]]) {
	bindings.FuncUpdateDynamicRules(
		js.Pointer(&fn),
	)
	return
}

// UpdateDynamicRules calls the function "WEBEXT.declarativeNetRequest.updateDynamicRules" directly.
func UpdateDynamicRules(options UpdateRuleOptions) (ret js.Promise[js.Void]) {
	bindings.CallUpdateDynamicRules(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateDynamicRules calls the function "WEBEXT.declarativeNetRequest.updateDynamicRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateDynamicRules(options UpdateRuleOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateDynamicRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncUpdateEnabledRulesets returns true if the function "WEBEXT.declarativeNetRequest.updateEnabledRulesets" exists.
func HasFuncUpdateEnabledRulesets() bool {
	return js.True == bindings.HasFuncUpdateEnabledRulesets()
}

// FuncUpdateEnabledRulesets returns the function "WEBEXT.declarativeNetRequest.updateEnabledRulesets".
func FuncUpdateEnabledRulesets() (fn js.Func[func(options UpdateRulesetOptions) js.Promise[js.Void]]) {
	bindings.FuncUpdateEnabledRulesets(
		js.Pointer(&fn),
	)
	return
}

// UpdateEnabledRulesets calls the function "WEBEXT.declarativeNetRequest.updateEnabledRulesets" directly.
func UpdateEnabledRulesets(options UpdateRulesetOptions) (ret js.Promise[js.Void]) {
	bindings.CallUpdateEnabledRulesets(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateEnabledRulesets calls the function "WEBEXT.declarativeNetRequest.updateEnabledRulesets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateEnabledRulesets(options UpdateRulesetOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateEnabledRulesets(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncUpdateSessionRules returns true if the function "WEBEXT.declarativeNetRequest.updateSessionRules" exists.
func HasFuncUpdateSessionRules() bool {
	return js.True == bindings.HasFuncUpdateSessionRules()
}

// FuncUpdateSessionRules returns the function "WEBEXT.declarativeNetRequest.updateSessionRules".
func FuncUpdateSessionRules() (fn js.Func[func(options UpdateRuleOptions) js.Promise[js.Void]]) {
	bindings.FuncUpdateSessionRules(
		js.Pointer(&fn),
	)
	return
}

// UpdateSessionRules calls the function "WEBEXT.declarativeNetRequest.updateSessionRules" directly.
func UpdateSessionRules(options UpdateRuleOptions) (ret js.Promise[js.Void]) {
	bindings.CallUpdateSessionRules(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateSessionRules calls the function "WEBEXT.declarativeNetRequest.updateSessionRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateSessionRules(options UpdateRuleOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateSessionRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncUpdateStaticRules returns true if the function "WEBEXT.declarativeNetRequest.updateStaticRules" exists.
func HasFuncUpdateStaticRules() bool {
	return js.True == bindings.HasFuncUpdateStaticRules()
}

// FuncUpdateStaticRules returns the function "WEBEXT.declarativeNetRequest.updateStaticRules".
func FuncUpdateStaticRules() (fn js.Func[func(options UpdateStaticRulesOptions) js.Promise[js.Void]]) {
	bindings.FuncUpdateStaticRules(
		js.Pointer(&fn),
	)
	return
}

// UpdateStaticRules calls the function "WEBEXT.declarativeNetRequest.updateStaticRules" directly.
func UpdateStaticRules(options UpdateStaticRulesOptions) (ret js.Promise[js.Void]) {
	bindings.CallUpdateStaticRules(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateStaticRules calls the function "WEBEXT.declarativeNetRequest.updateStaticRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateStaticRules(options UpdateStaticRulesOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateStaticRules(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

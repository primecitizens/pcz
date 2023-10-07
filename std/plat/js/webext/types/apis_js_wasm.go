// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package types

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/types/bindings"
)

type LevelOfControl uint32

const (
	_ LevelOfControl = iota

	LevelOfControl_NOT_CONTROLLABLE
	LevelOfControl_CONTROLLED_BY_OTHER_EXTENSIONS
	LevelOfControl_CONTROLLABLE_BY_THIS_EXTENSION
	LevelOfControl_CONTROLLED_BY_THIS_EXTENSION
)

func (LevelOfControl) FromRef(str js.Ref) LevelOfControl {
	return LevelOfControl(bindings.ConstOfLevelOfControl(str))
}

func (x LevelOfControl) String() (string, bool) {
	switch x {
	case LevelOfControl_NOT_CONTROLLABLE:
		return "not_controllable", true
	case LevelOfControl_CONTROLLED_BY_OTHER_EXTENSIONS:
		return "controlled_by_other_extensions", true
	case LevelOfControl_CONTROLLABLE_BY_THIS_EXTENSION:
		return "controllable_by_this_extension", true
	case LevelOfControl_CONTROLLED_BY_THIS_EXTENSION:
		return "controlled_by_this_extension", true
	default:
		return "", false
	}
}

type GetReturnType struct {
	// IncognitoSpecific is "GetReturnType.incognitoSpecific"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncognitoSpecific MUST be set to true to make this field effective.
	IncognitoSpecific bool
	// LevelOfControl is "GetReturnType.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// Value is "GetReturnType.value"
	//
	// Required
	Value js.Any

	FFI_USE_IncognitoSpecific bool // for IncognitoSpecific.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetReturnType with all fields set.
func (p GetReturnType) FromRef(ref js.Ref) GetReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetReturnType in the application heap.
func (p GetReturnType) New() js.Ref {
	return bindings.GetReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetReturnType) Update(ref js.Ref) {
	bindings.GetReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type GetArgDetails struct {
	// Incognito is "GetArgDetails.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool

	FFI_USE_Incognito bool // for Incognito.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetArgDetails with all fields set.
func (p GetArgDetails) FromRef(ref js.Ref) GetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetArgDetails in the application heap.
func (p GetArgDetails) New() js.Ref {
	return bindings.GetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetArgDetails) Update(ref js.Ref) {
	bindings.GetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetArgDetails) FreeMembers(recursive bool) {
}

type ChromeSettingScope uint32

const (
	_ ChromeSettingScope = iota

	ChromeSettingScope_REGULAR
	ChromeSettingScope_REGULAR_ONLY
	ChromeSettingScope_INCOGNITO_PERSISTENT
	ChromeSettingScope_INCOGNITO_SESSION_ONLY
)

func (ChromeSettingScope) FromRef(str js.Ref) ChromeSettingScope {
	return ChromeSettingScope(bindings.ConstOfChromeSettingScope(str))
}

func (x ChromeSettingScope) String() (string, bool) {
	switch x {
	case ChromeSettingScope_REGULAR:
		return "regular", true
	case ChromeSettingScope_REGULAR_ONLY:
		return "regular_only", true
	case ChromeSettingScope_INCOGNITO_PERSISTENT:
		return "incognito_persistent", true
	case ChromeSettingScope_INCOGNITO_SESSION_ONLY:
		return "incognito_session_only", true
	default:
		return "", false
	}
}

type SetArgDetails struct {
	// Scope is "SetArgDetails.scope"
	//
	// Optional
	Scope ChromeSettingScope
	// Value is "SetArgDetails.value"
	//
	// Required
	Value js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetArgDetails with all fields set.
func (p SetArgDetails) FromRef(ref js.Ref) SetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetArgDetails in the application heap.
func (p SetArgDetails) New() js.Ref {
	return bindings.SetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetArgDetails) Update(ref js.Ref) {
	bindings.SetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type ClearArgDetails struct {
	// Scope is "ClearArgDetails.scope"
	//
	// Optional
	Scope ChromeSettingScope

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearArgDetails with all fields set.
func (p ClearArgDetails) FromRef(ref js.Ref) ClearArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearArgDetails in the application heap.
func (p ClearArgDetails) New() js.Ref {
	return bindings.ClearArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearArgDetails) Update(ref js.Ref) {
	bindings.ClearArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearArgDetails) FreeMembers(recursive bool) {
}

type ChromeSetting struct {
	ref js.Ref
}

func (this ChromeSetting) Once() ChromeSetting {
	this.ref.Once()
	return this
}

func (this ChromeSetting) Ref() js.Ref {
	return this.ref
}

func (this ChromeSetting) FromRef(ref js.Ref) ChromeSetting {
	this.ref = ref
	return this
}

func (this ChromeSetting) Free() {
	this.ref.Free()
}

// HasFuncGet returns true if the method "ChromeSetting.get" exists.
func (this ChromeSetting) HasFuncGet() bool {
	return js.True == bindings.HasFuncChromeSettingGet(
		this.ref,
	)
}

// FuncGet returns the method "ChromeSetting.get".
func (this ChromeSetting) FuncGet() (fn js.Func[func(details GetArgDetails) js.Promise[GetReturnType]]) {
	bindings.FuncChromeSettingGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "ChromeSetting.get".
func (this ChromeSetting) Get(details GetArgDetails) (ret js.Promise[GetReturnType]) {
	bindings.CallChromeSettingGet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGet calls the method "ChromeSetting.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ChromeSetting) TryGet(details GetArgDetails) (ret js.Promise[GetReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChromeSettingGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSet returns true if the method "ChromeSetting.set" exists.
func (this ChromeSetting) HasFuncSet() bool {
	return js.True == bindings.HasFuncChromeSettingSet(
		this.ref,
	)
}

// FuncSet returns the method "ChromeSetting.set".
func (this ChromeSetting) FuncSet() (fn js.Func[func(details SetArgDetails) js.Promise[js.Void]]) {
	bindings.FuncChromeSettingSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "ChromeSetting.set".
func (this ChromeSetting) Set(details SetArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallChromeSettingSet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySet calls the method "ChromeSetting.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ChromeSetting) TrySet(details SetArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChromeSettingSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncClear returns true if the method "ChromeSetting.clear" exists.
func (this ChromeSetting) HasFuncClear() bool {
	return js.True == bindings.HasFuncChromeSettingClear(
		this.ref,
	)
}

// FuncClear returns the method "ChromeSetting.clear".
func (this ChromeSetting) FuncClear() (fn js.Func[func(details ClearArgDetails) js.Promise[js.Void]]) {
	bindings.FuncChromeSettingClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "ChromeSetting.clear".
func (this ChromeSetting) Clear(details ClearArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallChromeSettingClear(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClear calls the method "ChromeSetting.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ChromeSetting) TryClear(details ClearArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChromeSettingClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

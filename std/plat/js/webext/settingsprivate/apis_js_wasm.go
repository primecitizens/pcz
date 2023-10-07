// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package settingsprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/settingsprivate/bindings"
)

type ControlledBy uint32

const (
	_ ControlledBy = iota

	ControlledBy_DEVICE_POLICY
	ControlledBy_USER_POLICY
	ControlledBy_OWNER
	ControlledBy_PRIMARY_USER
	ControlledBy_EXTENSION
	ControlledBy_PARENT
	ControlledBy_CHILD_RESTRICTION
)

func (ControlledBy) FromRef(str js.Ref) ControlledBy {
	return ControlledBy(bindings.ConstOfControlledBy(str))
}

func (x ControlledBy) String() (string, bool) {
	switch x {
	case ControlledBy_DEVICE_POLICY:
		return "DEVICE_POLICY", true
	case ControlledBy_USER_POLICY:
		return "USER_POLICY", true
	case ControlledBy_OWNER:
		return "OWNER", true
	case ControlledBy_PRIMARY_USER:
		return "PRIMARY_USER", true
	case ControlledBy_EXTENSION:
		return "EXTENSION", true
	case ControlledBy_PARENT:
		return "PARENT", true
	case ControlledBy_CHILD_RESTRICTION:
		return "CHILD_RESTRICTION", true
	default:
		return "", false
	}
}

type Enforcement uint32

const (
	_ Enforcement = iota

	Enforcement_ENFORCED
	Enforcement_RECOMMENDED
	Enforcement_PARENT_SUPERVISED
)

func (Enforcement) FromRef(str js.Ref) Enforcement {
	return Enforcement(bindings.ConstOfEnforcement(str))
}

func (x Enforcement) String() (string, bool) {
	switch x {
	case Enforcement_ENFORCED:
		return "ENFORCED", true
	case Enforcement_RECOMMENDED:
		return "RECOMMENDED", true
	case Enforcement_PARENT_SUPERVISED:
		return "PARENT_SUPERVISED", true
	default:
		return "", false
	}
}

type GetAllPrefsCallbackFunc func(this js.Ref, prefs js.Array[PrefObject]) js.Ref

func (fn GetAllPrefsCallbackFunc) Register() js.Func[func(prefs js.Array[PrefObject])] {
	return js.RegisterCallback[func(prefs js.Array[PrefObject])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllPrefsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PrefObject]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllPrefsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, prefs js.Array[PrefObject]) js.Ref
	Arg T
}

func (cb *GetAllPrefsCallback[T]) Register() js.Func[func(prefs js.Array[PrefObject])] {
	return js.RegisterCallback[func(prefs js.Array[PrefObject])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllPrefsCallback[T]) DispatchCallback(
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

		js.Array[PrefObject]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PrefType uint32

const (
	_ PrefType = iota

	PrefType_BOOLEAN
	PrefType_NUMBER
	PrefType_STRING
	PrefType_URL
	PrefType_LIST
	PrefType_DICTIONARY
)

func (PrefType) FromRef(str js.Ref) PrefType {
	return PrefType(bindings.ConstOfPrefType(str))
}

func (x PrefType) String() (string, bool) {
	switch x {
	case PrefType_BOOLEAN:
		return "BOOLEAN", true
	case PrefType_NUMBER:
		return "NUMBER", true
	case PrefType_STRING:
		return "STRING", true
	case PrefType_URL:
		return "URL", true
	case PrefType_LIST:
		return "LIST", true
	case PrefType_DICTIONARY:
		return "DICTIONARY", true
	default:
		return "", false
	}
}

type PrefObject struct {
	// Key is "PrefObject.key"
	//
	// Optional
	Key js.String
	// Type is "PrefObject.type"
	//
	// Optional
	Type PrefType
	// Value is "PrefObject.value"
	//
	// Optional
	Value js.Any
	// ControlledBy is "PrefObject.controlledBy"
	//
	// Optional
	ControlledBy ControlledBy
	// ControlledByName is "PrefObject.controlledByName"
	//
	// Optional
	ControlledByName js.String
	// Enforcement is "PrefObject.enforcement"
	//
	// Optional
	Enforcement Enforcement
	// RecommendedValue is "PrefObject.recommendedValue"
	//
	// Optional
	RecommendedValue js.Any
	// UserSelectableValues is "PrefObject.userSelectableValues"
	//
	// Optional
	UserSelectableValues js.Array[js.Any]
	// UserControlDisabled is "PrefObject.userControlDisabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserControlDisabled MUST be set to true to make this field effective.
	UserControlDisabled bool
	// ExtensionId is "PrefObject.extensionId"
	//
	// Optional
	ExtensionId js.String
	// ExtensionCanBeDisabled is "PrefObject.extensionCanBeDisabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExtensionCanBeDisabled MUST be set to true to make this field effective.
	ExtensionCanBeDisabled bool

	FFI_USE_UserControlDisabled    bool // for UserControlDisabled.
	FFI_USE_ExtensionCanBeDisabled bool // for ExtensionCanBeDisabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrefObject with all fields set.
func (p PrefObject) FromRef(ref js.Ref) PrefObject {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrefObject in the application heap.
func (p PrefObject) New() js.Ref {
	return bindings.PrefObjectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrefObject) UpdateFrom(ref js.Ref) {
	bindings.PrefObjectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrefObject) Update(ref js.Ref) {
	bindings.PrefObjectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrefObject) FreeMembers(recursive bool) {
	js.Free(
		p.Key.Ref(),
		p.Value.Ref(),
		p.ControlledByName.Ref(),
		p.RecommendedValue.Ref(),
		p.UserSelectableValues.Ref(),
		p.ExtensionId.Ref(),
	)
	p.Key = p.Key.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
	p.ControlledByName = p.ControlledByName.FromRef(js.Undefined)
	p.RecommendedValue = p.RecommendedValue.FromRef(js.Undefined)
	p.UserSelectableValues = p.UserSelectableValues.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
}

type GetDefaultZoomCallbackFunc func(this js.Ref, zoom float64) js.Ref

func (fn GetDefaultZoomCallbackFunc) Register() js.Func[func(zoom float64)] {
	return js.RegisterCallback[func(zoom float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDefaultZoomCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDefaultZoomCallback[T any] struct {
	Fn  func(arg T, this js.Ref, zoom float64) js.Ref
	Arg T
}

func (cb *GetDefaultZoomCallback[T]) Register() js.Func[func(zoom float64)] {
	return js.RegisterCallback[func(zoom float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDefaultZoomCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPrefCallbackFunc func(this js.Ref, pref *PrefObject) js.Ref

func (fn GetPrefCallbackFunc) Register() js.Func[func(pref *PrefObject)] {
	return js.RegisterCallback[func(pref *PrefObject)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPrefCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrefObject
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

type GetPrefCallback[T any] struct {
	Fn  func(arg T, this js.Ref, pref *PrefObject) js.Ref
	Arg T
}

func (cb *GetPrefCallback[T]) Register() js.Func[func(pref *PrefObject)] {
	return js.RegisterCallback[func(pref *PrefObject)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPrefCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrefObject
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

type OnPrefSetCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn OnPrefSetCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrefSetCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPrefSetCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *OnPrefSetCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrefSetCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDefaultZoomCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn SetDefaultZoomCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetDefaultZoomCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDefaultZoomCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *SetDefaultZoomCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetDefaultZoomCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetAllPrefs returns true if the function "WEBEXT.settingsPrivate.getAllPrefs" exists.
func HasFuncGetAllPrefs() bool {
	return js.True == bindings.HasFuncGetAllPrefs()
}

// FuncGetAllPrefs returns the function "WEBEXT.settingsPrivate.getAllPrefs".
func FuncGetAllPrefs() (fn js.Func[func() js.Promise[js.Array[PrefObject]]]) {
	bindings.FuncGetAllPrefs(
		js.Pointer(&fn),
	)
	return
}

// GetAllPrefs calls the function "WEBEXT.settingsPrivate.getAllPrefs" directly.
func GetAllPrefs() (ret js.Promise[js.Array[PrefObject]]) {
	bindings.CallGetAllPrefs(
		js.Pointer(&ret),
	)

	return
}

// TryGetAllPrefs calls the function "WEBEXT.settingsPrivate.getAllPrefs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllPrefs() (ret js.Promise[js.Array[PrefObject]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllPrefs(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDefaultZoom returns true if the function "WEBEXT.settingsPrivate.getDefaultZoom" exists.
func HasFuncGetDefaultZoom() bool {
	return js.True == bindings.HasFuncGetDefaultZoom()
}

// FuncGetDefaultZoom returns the function "WEBEXT.settingsPrivate.getDefaultZoom".
func FuncGetDefaultZoom() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	bindings.FuncGetDefaultZoom(
		js.Pointer(&fn),
	)
	return
}

// GetDefaultZoom calls the function "WEBEXT.settingsPrivate.getDefaultZoom" directly.
func GetDefaultZoom() (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetDefaultZoom(
		js.Pointer(&ret),
	)

	return
}

// TryGetDefaultZoom calls the function "WEBEXT.settingsPrivate.getDefaultZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDefaultZoom() (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDefaultZoom(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPref returns true if the function "WEBEXT.settingsPrivate.getPref" exists.
func HasFuncGetPref() bool {
	return js.True == bindings.HasFuncGetPref()
}

// FuncGetPref returns the function "WEBEXT.settingsPrivate.getPref".
func FuncGetPref() (fn js.Func[func(name js.String) js.Promise[PrefObject]]) {
	bindings.FuncGetPref(
		js.Pointer(&fn),
	)
	return
}

// GetPref calls the function "WEBEXT.settingsPrivate.getPref" directly.
func GetPref(name js.String) (ret js.Promise[PrefObject]) {
	bindings.CallGetPref(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetPref calls the function "WEBEXT.settingsPrivate.getPref"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPref(name js.String) (ret js.Promise[PrefObject], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPref(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type OnPrefsChangedEventCallbackFunc func(this js.Ref, prefs js.Array[PrefObject]) js.Ref

func (fn OnPrefsChangedEventCallbackFunc) Register() js.Func[func(prefs js.Array[PrefObject])] {
	return js.RegisterCallback[func(prefs js.Array[PrefObject])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrefsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PrefObject]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPrefsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, prefs js.Array[PrefObject]) js.Ref
	Arg T
}

func (cb *OnPrefsChangedEventCallback[T]) Register() js.Func[func(prefs js.Array[PrefObject])] {
	return js.RegisterCallback[func(prefs js.Array[PrefObject])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrefsChangedEventCallback[T]) DispatchCallback(
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

		js.Array[PrefObject]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPrefsChanged returns true if the function "WEBEXT.settingsPrivate.onPrefsChanged.addListener" exists.
func HasFuncOnPrefsChanged() bool {
	return js.True == bindings.HasFuncOnPrefsChanged()
}

// FuncOnPrefsChanged returns the function "WEBEXT.settingsPrivate.onPrefsChanged.addListener".
func FuncOnPrefsChanged() (fn js.Func[func(callback js.Func[func(prefs js.Array[PrefObject])])]) {
	bindings.FuncOnPrefsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.addListener" directly.
func OnPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret js.Void) {
	bindings.CallOnPrefsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPrefsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPrefsChanged returns true if the function "WEBEXT.settingsPrivate.onPrefsChanged.removeListener" exists.
func HasFuncOffPrefsChanged() bool {
	return js.True == bindings.HasFuncOffPrefsChanged()
}

// FuncOffPrefsChanged returns the function "WEBEXT.settingsPrivate.onPrefsChanged.removeListener".
func FuncOffPrefsChanged() (fn js.Func[func(callback js.Func[func(prefs js.Array[PrefObject])])]) {
	bindings.FuncOffPrefsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.removeListener" directly.
func OffPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret js.Void) {
	bindings.CallOffPrefsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPrefsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPrefsChanged returns true if the function "WEBEXT.settingsPrivate.onPrefsChanged.hasListener" exists.
func HasFuncHasOnPrefsChanged() bool {
	return js.True == bindings.HasFuncHasOnPrefsChanged()
}

// FuncHasOnPrefsChanged returns the function "WEBEXT.settingsPrivate.onPrefsChanged.hasListener".
func FuncHasOnPrefsChanged() (fn js.Func[func(callback js.Func[func(prefs js.Array[PrefObject])]) bool]) {
	bindings.FuncHasOnPrefsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.hasListener" directly.
func HasOnPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret bool) {
	bindings.CallHasOnPrefsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPrefsChanged calls the function "WEBEXT.settingsPrivate.onPrefsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPrefsChanged(callback js.Func[func(prefs js.Array[PrefObject])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPrefsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetDefaultZoom returns true if the function "WEBEXT.settingsPrivate.setDefaultZoom" exists.
func HasFuncSetDefaultZoom() bool {
	return js.True == bindings.HasFuncSetDefaultZoom()
}

// FuncSetDefaultZoom returns the function "WEBEXT.settingsPrivate.setDefaultZoom".
func FuncSetDefaultZoom() (fn js.Func[func(zoom float64) js.Promise[js.Boolean]]) {
	bindings.FuncSetDefaultZoom(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultZoom calls the function "WEBEXT.settingsPrivate.setDefaultZoom" directly.
func SetDefaultZoom(zoom float64) (ret js.Promise[js.Boolean]) {
	bindings.CallSetDefaultZoom(
		js.Pointer(&ret),
		float64(zoom),
	)

	return
}

// TrySetDefaultZoom calls the function "WEBEXT.settingsPrivate.setDefaultZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultZoom(zoom float64) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(zoom),
	)

	return
}

// HasFuncSetPref returns true if the function "WEBEXT.settingsPrivate.setPref" exists.
func HasFuncSetPref() bool {
	return js.True == bindings.HasFuncSetPref()
}

// FuncSetPref returns the function "WEBEXT.settingsPrivate.setPref".
func FuncSetPref() (fn js.Func[func(name js.String, value js.Any, pageId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncSetPref(
		js.Pointer(&fn),
	)
	return
}

// SetPref calls the function "WEBEXT.settingsPrivate.setPref" directly.
func SetPref(name js.String, value js.Any, pageId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallSetPref(
		js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
		pageId.Ref(),
	)

	return
}

// TrySetPref calls the function "WEBEXT.settingsPrivate.setPref"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPref(name js.String, value js.Any, pageId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPref(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
		pageId.Ref(),
	)

	return
}

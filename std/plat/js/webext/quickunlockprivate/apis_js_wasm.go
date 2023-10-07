// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package quickunlockprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/quickunlockprivate/bindings"
)

type BooleanResultCallbackFunc func(this js.Ref, value bool) js.Ref

func (fn BooleanResultCallbackFunc) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BooleanResultCallbackFunc) DispatchCallback(
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

type BooleanResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref, value bool) js.Ref
	Arg T
}

func (cb *BooleanResultCallback[T]) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BooleanResultCallback[T]) DispatchCallback(
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

type CredentialProblem uint32

const (
	_ CredentialProblem = iota

	CredentialProblem_TOO_SHORT
	CredentialProblem_TOO_LONG
	CredentialProblem_TOO_WEAK
	CredentialProblem_CONTAINS_NONDIGIT
)

func (CredentialProblem) FromRef(str js.Ref) CredentialProblem {
	return CredentialProblem(bindings.ConstOfCredentialProblem(str))
}

func (x CredentialProblem) String() (string, bool) {
	switch x {
	case CredentialProblem_TOO_SHORT:
		return "TOO_SHORT", true
	case CredentialProblem_TOO_LONG:
		return "TOO_LONG", true
	case CredentialProblem_TOO_WEAK:
		return "TOO_WEAK", true
	case CredentialProblem_CONTAINS_NONDIGIT:
		return "CONTAINS_NONDIGIT", true
	default:
		return "", false
	}
}

type CredentialCheck struct {
	// Errors is "CredentialCheck.errors"
	//
	// Optional
	Errors js.Array[CredentialProblem]
	// Warnings is "CredentialCheck.warnings"
	//
	// Optional
	Warnings js.Array[CredentialProblem]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialCheck with all fields set.
func (p CredentialCheck) FromRef(ref js.Ref) CredentialCheck {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialCheck in the application heap.
func (p CredentialCheck) New() js.Ref {
	return bindings.CredentialCheckJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialCheck) UpdateFrom(ref js.Ref) {
	bindings.CredentialCheckJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialCheck) Update(ref js.Ref) {
	bindings.CredentialCheckJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialCheck) FreeMembers(recursive bool) {
	js.Free(
		p.Errors.Ref(),
		p.Warnings.Ref(),
	)
	p.Errors = p.Errors.FromRef(js.Undefined)
	p.Warnings = p.Warnings.FromRef(js.Undefined)
}

type CredentialCheckCallbackFunc func(this js.Ref, check *CredentialCheck) js.Ref

func (fn CredentialCheckCallbackFunc) Register() js.Func[func(check *CredentialCheck)] {
	return js.RegisterCallback[func(check *CredentialCheck)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CredentialCheckCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CredentialCheck
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

type CredentialCheckCallback[T any] struct {
	Fn  func(arg T, this js.Ref, check *CredentialCheck) js.Ref
	Arg T
}

func (cb *CredentialCheckCallback[T]) Register() js.Func[func(check *CredentialCheck)] {
	return js.RegisterCallback[func(check *CredentialCheck)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CredentialCheckCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CredentialCheck
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

type CredentialRequirements struct {
	// MinLength is "CredentialRequirements.minLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinLength MUST be set to true to make this field effective.
	MinLength int32
	// MaxLength is "CredentialRequirements.maxLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxLength MUST be set to true to make this field effective.
	MaxLength int32

	FFI_USE_MinLength bool // for MinLength.
	FFI_USE_MaxLength bool // for MaxLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialRequirements with all fields set.
func (p CredentialRequirements) FromRef(ref js.Ref) CredentialRequirements {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialRequirements in the application heap.
func (p CredentialRequirements) New() js.Ref {
	return bindings.CredentialRequirementsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialRequirements) UpdateFrom(ref js.Ref) {
	bindings.CredentialRequirementsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialRequirements) Update(ref js.Ref) {
	bindings.CredentialRequirementsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialRequirements) FreeMembers(recursive bool) {
}

type CredentialRequirementsCallbackFunc func(this js.Ref, requirements *CredentialRequirements) js.Ref

func (fn CredentialRequirementsCallbackFunc) Register() js.Func[func(requirements *CredentialRequirements)] {
	return js.RegisterCallback[func(requirements *CredentialRequirements)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CredentialRequirementsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CredentialRequirements
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

type CredentialRequirementsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requirements *CredentialRequirements) js.Ref
	Arg T
}

func (cb *CredentialRequirementsCallback[T]) Register() js.Func[func(requirements *CredentialRequirements)] {
	return js.RegisterCallback[func(requirements *CredentialRequirements)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CredentialRequirementsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CredentialRequirements
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

type ModesCallbackFunc func(this js.Ref, modes js.Array[QuickUnlockMode]) js.Ref

func (fn ModesCallbackFunc) Register() js.Func[func(modes js.Array[QuickUnlockMode])] {
	return js.RegisterCallback[func(modes js.Array[QuickUnlockMode])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ModesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[QuickUnlockMode]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ModesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, modes js.Array[QuickUnlockMode]) js.Ref
	Arg T
}

func (cb *ModesCallback[T]) Register() js.Func[func(modes js.Array[QuickUnlockMode])] {
	return js.RegisterCallback[func(modes js.Array[QuickUnlockMode])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ModesCallback[T]) DispatchCallback(
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

		js.Array[QuickUnlockMode]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QuickUnlockMode uint32

const (
	_ QuickUnlockMode = iota

	QuickUnlockMode_PIN
)

func (QuickUnlockMode) FromRef(str js.Ref) QuickUnlockMode {
	return QuickUnlockMode(bindings.ConstOfQuickUnlockMode(str))
}

func (x QuickUnlockMode) String() (string, bool) {
	switch x {
	case QuickUnlockMode_PIN:
		return "PIN", true
	default:
		return "", false
	}
}

type TokenInfo struct {
	// Token is "TokenInfo.token"
	//
	// Optional
	Token js.String
	// LifetimeSeconds is "TokenInfo.lifetimeSeconds"
	//
	// Optional
	//
	// NOTE: FFI_USE_LifetimeSeconds MUST be set to true to make this field effective.
	LifetimeSeconds int32

	FFI_USE_LifetimeSeconds bool // for LifetimeSeconds.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TokenInfo with all fields set.
func (p TokenInfo) FromRef(ref js.Ref) TokenInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TokenInfo in the application heap.
func (p TokenInfo) New() js.Ref {
	return bindings.TokenInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TokenInfo) UpdateFrom(ref js.Ref) {
	bindings.TokenInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TokenInfo) Update(ref js.Ref) {
	bindings.TokenInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TokenInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Token.Ref(),
	)
	p.Token = p.Token.FromRef(js.Undefined)
}

type TokenResultCallbackFunc func(this js.Ref, result *TokenInfo) js.Ref

func (fn TokenResultCallbackFunc) Register() js.Func[func(result *TokenInfo)] {
	return js.RegisterCallback[func(result *TokenInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TokenResultCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TokenInfo
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

type TokenResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *TokenInfo) js.Ref
	Arg T
}

func (cb *TokenResultCallback[T]) Register() js.Func[func(result *TokenInfo)] {
	return js.RegisterCallback[func(result *TokenInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TokenResultCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TokenInfo
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

type VoidResultCallbackFunc func(this js.Ref) js.Ref

func (fn VoidResultCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidResultCallbackFunc) DispatchCallback(
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

type VoidResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidResultCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidResultCallback[T]) DispatchCallback(
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

// HasFuncCanAuthenticatePin returns true if the function "WEBEXT.quickUnlockPrivate.canAuthenticatePin" exists.
func HasFuncCanAuthenticatePin() bool {
	return js.True == bindings.HasFuncCanAuthenticatePin()
}

// FuncCanAuthenticatePin returns the function "WEBEXT.quickUnlockPrivate.canAuthenticatePin".
func FuncCanAuthenticatePin() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncCanAuthenticatePin(
		js.Pointer(&fn),
	)
	return
}

// CanAuthenticatePin calls the function "WEBEXT.quickUnlockPrivate.canAuthenticatePin" directly.
func CanAuthenticatePin() (ret js.Promise[js.Boolean]) {
	bindings.CallCanAuthenticatePin(
		js.Pointer(&ret),
	)

	return
}

// TryCanAuthenticatePin calls the function "WEBEXT.quickUnlockPrivate.canAuthenticatePin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCanAuthenticatePin() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanAuthenticatePin(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckCredential returns true if the function "WEBEXT.quickUnlockPrivate.checkCredential" exists.
func HasFuncCheckCredential() bool {
	return js.True == bindings.HasFuncCheckCredential()
}

// FuncCheckCredential returns the function "WEBEXT.quickUnlockPrivate.checkCredential".
func FuncCheckCredential() (fn js.Func[func(mode QuickUnlockMode, credential js.String) js.Promise[CredentialCheck]]) {
	bindings.FuncCheckCredential(
		js.Pointer(&fn),
	)
	return
}

// CheckCredential calls the function "WEBEXT.quickUnlockPrivate.checkCredential" directly.
func CheckCredential(mode QuickUnlockMode, credential js.String) (ret js.Promise[CredentialCheck]) {
	bindings.CallCheckCredential(
		js.Pointer(&ret),
		uint32(mode),
		credential.Ref(),
	)

	return
}

// TryCheckCredential calls the function "WEBEXT.quickUnlockPrivate.checkCredential"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCheckCredential(mode QuickUnlockMode, credential js.String) (ret js.Promise[CredentialCheck], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCheckCredential(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		credential.Ref(),
	)

	return
}

// HasFuncGetActiveModes returns true if the function "WEBEXT.quickUnlockPrivate.getActiveModes" exists.
func HasFuncGetActiveModes() bool {
	return js.True == bindings.HasFuncGetActiveModes()
}

// FuncGetActiveModes returns the function "WEBEXT.quickUnlockPrivate.getActiveModes".
func FuncGetActiveModes() (fn js.Func[func() js.Promise[js.Array[QuickUnlockMode]]]) {
	bindings.FuncGetActiveModes(
		js.Pointer(&fn),
	)
	return
}

// GetActiveModes calls the function "WEBEXT.quickUnlockPrivate.getActiveModes" directly.
func GetActiveModes() (ret js.Promise[js.Array[QuickUnlockMode]]) {
	bindings.CallGetActiveModes(
		js.Pointer(&ret),
	)

	return
}

// TryGetActiveModes calls the function "WEBEXT.quickUnlockPrivate.getActiveModes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetActiveModes() (ret js.Promise[js.Array[QuickUnlockMode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetActiveModes(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAuthToken returns true if the function "WEBEXT.quickUnlockPrivate.getAuthToken" exists.
func HasFuncGetAuthToken() bool {
	return js.True == bindings.HasFuncGetAuthToken()
}

// FuncGetAuthToken returns the function "WEBEXT.quickUnlockPrivate.getAuthToken".
func FuncGetAuthToken() (fn js.Func[func(accountPassword js.String) js.Promise[TokenInfo]]) {
	bindings.FuncGetAuthToken(
		js.Pointer(&fn),
	)
	return
}

// GetAuthToken calls the function "WEBEXT.quickUnlockPrivate.getAuthToken" directly.
func GetAuthToken(accountPassword js.String) (ret js.Promise[TokenInfo]) {
	bindings.CallGetAuthToken(
		js.Pointer(&ret),
		accountPassword.Ref(),
	)

	return
}

// TryGetAuthToken calls the function "WEBEXT.quickUnlockPrivate.getAuthToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAuthToken(accountPassword js.String) (ret js.Promise[TokenInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAuthToken(
		js.Pointer(&ret), js.Pointer(&exception),
		accountPassword.Ref(),
	)

	return
}

// HasFuncGetAvailableModes returns true if the function "WEBEXT.quickUnlockPrivate.getAvailableModes" exists.
func HasFuncGetAvailableModes() bool {
	return js.True == bindings.HasFuncGetAvailableModes()
}

// FuncGetAvailableModes returns the function "WEBEXT.quickUnlockPrivate.getAvailableModes".
func FuncGetAvailableModes() (fn js.Func[func() js.Promise[js.Array[QuickUnlockMode]]]) {
	bindings.FuncGetAvailableModes(
		js.Pointer(&fn),
	)
	return
}

// GetAvailableModes calls the function "WEBEXT.quickUnlockPrivate.getAvailableModes" directly.
func GetAvailableModes() (ret js.Promise[js.Array[QuickUnlockMode]]) {
	bindings.CallGetAvailableModes(
		js.Pointer(&ret),
	)

	return
}

// TryGetAvailableModes calls the function "WEBEXT.quickUnlockPrivate.getAvailableModes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAvailableModes() (ret js.Promise[js.Array[QuickUnlockMode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAvailableModes(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCredentialRequirements returns true if the function "WEBEXT.quickUnlockPrivate.getCredentialRequirements" exists.
func HasFuncGetCredentialRequirements() bool {
	return js.True == bindings.HasFuncGetCredentialRequirements()
}

// FuncGetCredentialRequirements returns the function "WEBEXT.quickUnlockPrivate.getCredentialRequirements".
func FuncGetCredentialRequirements() (fn js.Func[func(mode QuickUnlockMode) js.Promise[CredentialRequirements]]) {
	bindings.FuncGetCredentialRequirements(
		js.Pointer(&fn),
	)
	return
}

// GetCredentialRequirements calls the function "WEBEXT.quickUnlockPrivate.getCredentialRequirements" directly.
func GetCredentialRequirements(mode QuickUnlockMode) (ret js.Promise[CredentialRequirements]) {
	bindings.CallGetCredentialRequirements(
		js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryGetCredentialRequirements calls the function "WEBEXT.quickUnlockPrivate.getCredentialRequirements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCredentialRequirements(mode QuickUnlockMode) (ret js.Promise[CredentialRequirements], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCredentialRequirements(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

type OnActiveModesChangedEventCallbackFunc func(this js.Ref, activeModes js.Array[QuickUnlockMode]) js.Ref

func (fn OnActiveModesChangedEventCallbackFunc) Register() js.Func[func(activeModes js.Array[QuickUnlockMode])] {
	return js.RegisterCallback[func(activeModes js.Array[QuickUnlockMode])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActiveModesChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[QuickUnlockMode]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnActiveModesChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activeModes js.Array[QuickUnlockMode]) js.Ref
	Arg T
}

func (cb *OnActiveModesChangedEventCallback[T]) Register() js.Func[func(activeModes js.Array[QuickUnlockMode])] {
	return js.RegisterCallback[func(activeModes js.Array[QuickUnlockMode])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActiveModesChangedEventCallback[T]) DispatchCallback(
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

		js.Array[QuickUnlockMode]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnActiveModesChanged returns true if the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener" exists.
func HasFuncOnActiveModesChanged() bool {
	return js.True == bindings.HasFuncOnActiveModesChanged()
}

// FuncOnActiveModesChanged returns the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener".
func FuncOnActiveModesChanged() (fn js.Func[func(callback js.Func[func(activeModes js.Array[QuickUnlockMode])])]) {
	bindings.FuncOnActiveModesChanged(
		js.Pointer(&fn),
	)
	return
}

// OnActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener" directly.
func OnActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret js.Void) {
	bindings.CallOnActiveModesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActiveModesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActiveModesChanged returns true if the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener" exists.
func HasFuncOffActiveModesChanged() bool {
	return js.True == bindings.HasFuncOffActiveModesChanged()
}

// FuncOffActiveModesChanged returns the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener".
func FuncOffActiveModesChanged() (fn js.Func[func(callback js.Func[func(activeModes js.Array[QuickUnlockMode])])]) {
	bindings.FuncOffActiveModesChanged(
		js.Pointer(&fn),
	)
	return
}

// OffActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener" directly.
func OffActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret js.Void) {
	bindings.CallOffActiveModesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActiveModesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActiveModesChanged returns true if the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener" exists.
func HasFuncHasOnActiveModesChanged() bool {
	return js.True == bindings.HasFuncHasOnActiveModesChanged()
}

// FuncHasOnActiveModesChanged returns the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener".
func FuncHasOnActiveModesChanged() (fn js.Func[func(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) bool]) {
	bindings.FuncHasOnActiveModesChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener" directly.
func HasOnActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret bool) {
	bindings.CallHasOnActiveModesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActiveModesChanged calls the function "WEBEXT.quickUnlockPrivate.onActiveModesChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActiveModesChanged(callback js.Func[func(activeModes js.Array[QuickUnlockMode])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActiveModesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetLockScreenEnabled returns true if the function "WEBEXT.quickUnlockPrivate.setLockScreenEnabled" exists.
func HasFuncSetLockScreenEnabled() bool {
	return js.True == bindings.HasFuncSetLockScreenEnabled()
}

// FuncSetLockScreenEnabled returns the function "WEBEXT.quickUnlockPrivate.setLockScreenEnabled".
func FuncSetLockScreenEnabled() (fn js.Func[func(token js.String, enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetLockScreenEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetLockScreenEnabled calls the function "WEBEXT.quickUnlockPrivate.setLockScreenEnabled" directly.
func SetLockScreenEnabled(token js.String, enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetLockScreenEnabled(
		js.Pointer(&ret),
		token.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetLockScreenEnabled calls the function "WEBEXT.quickUnlockPrivate.setLockScreenEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetLockScreenEnabled(token js.String, enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetLockScreenEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetModes returns true if the function "WEBEXT.quickUnlockPrivate.setModes" exists.
func HasFuncSetModes() bool {
	return js.True == bindings.HasFuncSetModes()
}

// FuncSetModes returns the function "WEBEXT.quickUnlockPrivate.setModes".
func FuncSetModes() (fn js.Func[func(token js.String, modes js.Array[QuickUnlockMode], credentials js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncSetModes(
		js.Pointer(&fn),
	)
	return
}

// SetModes calls the function "WEBEXT.quickUnlockPrivate.setModes" directly.
func SetModes(token js.String, modes js.Array[QuickUnlockMode], credentials js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallSetModes(
		js.Pointer(&ret),
		token.Ref(),
		modes.Ref(),
		credentials.Ref(),
	)

	return
}

// TrySetModes calls the function "WEBEXT.quickUnlockPrivate.setModes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetModes(token js.String, modes js.Array[QuickUnlockMode], credentials js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetModes(
		js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		modes.Ref(),
		credentials.Ref(),
	)

	return
}

// HasFuncSetPinAutosubmitEnabled returns true if the function "WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled" exists.
func HasFuncSetPinAutosubmitEnabled() bool {
	return js.True == bindings.HasFuncSetPinAutosubmitEnabled()
}

// FuncSetPinAutosubmitEnabled returns the function "WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled".
func FuncSetPinAutosubmitEnabled() (fn js.Func[func(token js.String, pin js.String, enabled bool) js.Promise[js.Boolean]]) {
	bindings.FuncSetPinAutosubmitEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetPinAutosubmitEnabled calls the function "WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled" directly.
func SetPinAutosubmitEnabled(token js.String, pin js.String, enabled bool) (ret js.Promise[js.Boolean]) {
	bindings.CallSetPinAutosubmitEnabled(
		js.Pointer(&ret),
		token.Ref(),
		pin.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetPinAutosubmitEnabled calls the function "WEBEXT.quickUnlockPrivate.setPinAutosubmitEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPinAutosubmitEnabled(token js.String, pin js.String, enabled bool) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPinAutosubmitEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		token.Ref(),
		pin.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package identity

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/identity/bindings"
)

type AccountInfo struct {
	// Id is "AccountInfo.id"
	//
	// Optional
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccountInfo with all fields set.
func (p AccountInfo) FromRef(ref js.Ref) AccountInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AccountInfo in the application heap.
func (p AccountInfo) New() js.Ref {
	return bindings.AccountInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AccountInfo) UpdateFrom(ref js.Ref) {
	bindings.AccountInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AccountInfo) Update(ref js.Ref) {
	bindings.AccountInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AccountInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type AccountStatus uint32

const (
	_ AccountStatus = iota

	AccountStatus_SYNC
	AccountStatus_ANY
)

func (AccountStatus) FromRef(str js.Ref) AccountStatus {
	return AccountStatus(bindings.ConstOfAccountStatus(str))
}

func (x AccountStatus) String() (string, bool) {
	switch x {
	case AccountStatus_SYNC:
		return "SYNC", true
	case AccountStatus_ANY:
		return "ANY", true
	default:
		return "", false
	}
}

type ClearAllCachedAuthTokensCallbackFunc func(this js.Ref) js.Ref

func (fn ClearAllCachedAuthTokensCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ClearAllCachedAuthTokensCallbackFunc) DispatchCallback(
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

type ClearAllCachedAuthTokensCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ClearAllCachedAuthTokensCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ClearAllCachedAuthTokensCallback[T]) DispatchCallback(
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

type GetAccountsCallbackFunc func(this js.Ref, accounts js.Array[AccountInfo]) js.Ref

func (fn GetAccountsCallbackFunc) Register() js.Func[func(accounts js.Array[AccountInfo])] {
	return js.RegisterCallback[func(accounts js.Array[AccountInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAccountsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AccountInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAccountsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, accounts js.Array[AccountInfo]) js.Ref
	Arg T
}

func (cb *GetAccountsCallback[T]) Register() js.Func[func(accounts js.Array[AccountInfo])] {
	return js.RegisterCallback[func(accounts js.Array[AccountInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAccountsCallback[T]) DispatchCallback(
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

		js.Array[AccountInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAuthTokenCallbackFunc func(this js.Ref, result *GetAuthTokenResult) js.Ref

func (fn GetAuthTokenCallbackFunc) Register() js.Func[func(result *GetAuthTokenResult)] {
	return js.RegisterCallback[func(result *GetAuthTokenResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAuthTokenCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAuthTokenResult
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

type GetAuthTokenCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *GetAuthTokenResult) js.Ref
	Arg T
}

func (cb *GetAuthTokenCallback[T]) Register() js.Func[func(result *GetAuthTokenResult)] {
	return js.RegisterCallback[func(result *GetAuthTokenResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAuthTokenCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAuthTokenResult
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

type GetAuthTokenResult struct {
	// Token is "GetAuthTokenResult.token"
	//
	// Optional
	Token js.String
	// GrantedScopes is "GetAuthTokenResult.grantedScopes"
	//
	// Optional
	GrantedScopes js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAuthTokenResult with all fields set.
func (p GetAuthTokenResult) FromRef(ref js.Ref) GetAuthTokenResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAuthTokenResult in the application heap.
func (p GetAuthTokenResult) New() js.Ref {
	return bindings.GetAuthTokenResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAuthTokenResult) UpdateFrom(ref js.Ref) {
	bindings.GetAuthTokenResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAuthTokenResult) Update(ref js.Ref) {
	bindings.GetAuthTokenResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAuthTokenResult) FreeMembers(recursive bool) {
	js.Free(
		p.Token.Ref(),
		p.GrantedScopes.Ref(),
	)
	p.Token = p.Token.FromRef(js.Undefined)
	p.GrantedScopes = p.GrantedScopes.FromRef(js.Undefined)
}

type GetProfileUserInfoCallbackFunc func(this js.Ref, userInfo *ProfileUserInfo) js.Ref

func (fn GetProfileUserInfoCallbackFunc) Register() js.Func[func(userInfo *ProfileUserInfo)] {
	return js.RegisterCallback[func(userInfo *ProfileUserInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProfileUserInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileUserInfo
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

type GetProfileUserInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, userInfo *ProfileUserInfo) js.Ref
	Arg T
}

func (cb *GetProfileUserInfoCallback[T]) Register() js.Func[func(userInfo *ProfileUserInfo)] {
	return js.RegisterCallback[func(userInfo *ProfileUserInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProfileUserInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileUserInfo
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

type ProfileUserInfo struct {
	// Email is "ProfileUserInfo.email"
	//
	// Optional
	Email js.String
	// Id is "ProfileUserInfo.id"
	//
	// Optional
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfileUserInfo with all fields set.
func (p ProfileUserInfo) FromRef(ref js.Ref) ProfileUserInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfileUserInfo in the application heap.
func (p ProfileUserInfo) New() js.Ref {
	return bindings.ProfileUserInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfileUserInfo) UpdateFrom(ref js.Ref) {
	bindings.ProfileUserInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfileUserInfo) Update(ref js.Ref) {
	bindings.ProfileUserInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfileUserInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.Id.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type InvalidTokenDetails struct {
	// Token is "InvalidTokenDetails.token"
	//
	// Optional
	Token js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InvalidTokenDetails with all fields set.
func (p InvalidTokenDetails) FromRef(ref js.Ref) InvalidTokenDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InvalidTokenDetails in the application heap.
func (p InvalidTokenDetails) New() js.Ref {
	return bindings.InvalidTokenDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InvalidTokenDetails) UpdateFrom(ref js.Ref) {
	bindings.InvalidTokenDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InvalidTokenDetails) Update(ref js.Ref) {
	bindings.InvalidTokenDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InvalidTokenDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Token.Ref(),
	)
	p.Token = p.Token.FromRef(js.Undefined)
}

type InvalidateAuthTokenCallbackFunc func(this js.Ref) js.Ref

func (fn InvalidateAuthTokenCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InvalidateAuthTokenCallbackFunc) DispatchCallback(
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

type InvalidateAuthTokenCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *InvalidateAuthTokenCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InvalidateAuthTokenCallback[T]) DispatchCallback(
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

type LaunchWebAuthFlowCallbackFunc func(this js.Ref, responseUrl js.String) js.Ref

func (fn LaunchWebAuthFlowCallbackFunc) Register() js.Func[func(responseUrl js.String)] {
	return js.RegisterCallback[func(responseUrl js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LaunchWebAuthFlowCallbackFunc) DispatchCallback(
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

type LaunchWebAuthFlowCallback[T any] struct {
	Fn  func(arg T, this js.Ref, responseUrl js.String) js.Ref
	Arg T
}

func (cb *LaunchWebAuthFlowCallback[T]) Register() js.Func[func(responseUrl js.String)] {
	return js.RegisterCallback[func(responseUrl js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LaunchWebAuthFlowCallback[T]) DispatchCallback(
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

type ProfileDetails struct {
	// AccountStatus is "ProfileDetails.accountStatus"
	//
	// Optional
	AccountStatus AccountStatus

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfileDetails with all fields set.
func (p ProfileDetails) FromRef(ref js.Ref) ProfileDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfileDetails in the application heap.
func (p ProfileDetails) New() js.Ref {
	return bindings.ProfileDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfileDetails) UpdateFrom(ref js.Ref) {
	bindings.ProfileDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfileDetails) Update(ref js.Ref) {
	bindings.ProfileDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfileDetails) FreeMembers(recursive bool) {
}

type TokenDetails struct {
	// Interactive is "TokenDetails.interactive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Interactive MUST be set to true to make this field effective.
	Interactive bool
	// Account is "TokenDetails.account"
	//
	// Optional
	//
	// NOTE: Account.FFI_USE MUST be set to true to get Account used.
	Account AccountInfo
	// Scopes is "TokenDetails.scopes"
	//
	// Optional
	Scopes js.Array[js.String]
	// EnableGranularPermissions is "TokenDetails.enableGranularPermissions"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableGranularPermissions MUST be set to true to make this field effective.
	EnableGranularPermissions bool

	FFI_USE_Interactive               bool // for Interactive.
	FFI_USE_EnableGranularPermissions bool // for EnableGranularPermissions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TokenDetails with all fields set.
func (p TokenDetails) FromRef(ref js.Ref) TokenDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TokenDetails in the application heap.
func (p TokenDetails) New() js.Ref {
	return bindings.TokenDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TokenDetails) UpdateFrom(ref js.Ref) {
	bindings.TokenDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TokenDetails) Update(ref js.Ref) {
	bindings.TokenDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TokenDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Scopes.Ref(),
	)
	p.Scopes = p.Scopes.FromRef(js.Undefined)
	if recursive {
		p.Account.FreeMembers(true)
	}
}

type WebAuthFlowDetails struct {
	// Url is "WebAuthFlowDetails.url"
	//
	// Optional
	Url js.String
	// Interactive is "WebAuthFlowDetails.interactive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Interactive MUST be set to true to make this field effective.
	Interactive bool
	// AbortOnLoadForNonInteractive is "WebAuthFlowDetails.abortOnLoadForNonInteractive"
	//
	// Optional
	//
	// NOTE: FFI_USE_AbortOnLoadForNonInteractive MUST be set to true to make this field effective.
	AbortOnLoadForNonInteractive bool
	// TimeoutMsForNonInteractive is "WebAuthFlowDetails.timeoutMsForNonInteractive"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeoutMsForNonInteractive MUST be set to true to make this field effective.
	TimeoutMsForNonInteractive int32

	FFI_USE_Interactive                  bool // for Interactive.
	FFI_USE_AbortOnLoadForNonInteractive bool // for AbortOnLoadForNonInteractive.
	FFI_USE_TimeoutMsForNonInteractive   bool // for TimeoutMsForNonInteractive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebAuthFlowDetails with all fields set.
func (p WebAuthFlowDetails) FromRef(ref js.Ref) WebAuthFlowDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebAuthFlowDetails in the application heap.
func (p WebAuthFlowDetails) New() js.Ref {
	return bindings.WebAuthFlowDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebAuthFlowDetails) UpdateFrom(ref js.Ref) {
	bindings.WebAuthFlowDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebAuthFlowDetails) Update(ref js.Ref) {
	bindings.WebAuthFlowDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebAuthFlowDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncClearAllCachedAuthTokens returns true if the function "WEBEXT.identity.clearAllCachedAuthTokens" exists.
func HasFuncClearAllCachedAuthTokens() bool {
	return js.True == bindings.HasFuncClearAllCachedAuthTokens()
}

// FuncClearAllCachedAuthTokens returns the function "WEBEXT.identity.clearAllCachedAuthTokens".
func FuncClearAllCachedAuthTokens() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncClearAllCachedAuthTokens(
		js.Pointer(&fn),
	)
	return
}

// ClearAllCachedAuthTokens calls the function "WEBEXT.identity.clearAllCachedAuthTokens" directly.
func ClearAllCachedAuthTokens() (ret js.Promise[js.Void]) {
	bindings.CallClearAllCachedAuthTokens(
		js.Pointer(&ret),
	)

	return
}

// TryClearAllCachedAuthTokens calls the function "WEBEXT.identity.clearAllCachedAuthTokens"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearAllCachedAuthTokens() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearAllCachedAuthTokens(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAccounts returns true if the function "WEBEXT.identity.getAccounts" exists.
func HasFuncGetAccounts() bool {
	return js.True == bindings.HasFuncGetAccounts()
}

// FuncGetAccounts returns the function "WEBEXT.identity.getAccounts".
func FuncGetAccounts() (fn js.Func[func() js.Promise[js.Array[AccountInfo]]]) {
	bindings.FuncGetAccounts(
		js.Pointer(&fn),
	)
	return
}

// GetAccounts calls the function "WEBEXT.identity.getAccounts" directly.
func GetAccounts() (ret js.Promise[js.Array[AccountInfo]]) {
	bindings.CallGetAccounts(
		js.Pointer(&ret),
	)

	return
}

// TryGetAccounts calls the function "WEBEXT.identity.getAccounts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAccounts() (ret js.Promise[js.Array[AccountInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAccounts(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAuthToken returns true if the function "WEBEXT.identity.getAuthToken" exists.
func HasFuncGetAuthToken() bool {
	return js.True == bindings.HasFuncGetAuthToken()
}

// FuncGetAuthToken returns the function "WEBEXT.identity.getAuthToken".
func FuncGetAuthToken() (fn js.Func[func(details TokenDetails) js.Promise[GetAuthTokenResult]]) {
	bindings.FuncGetAuthToken(
		js.Pointer(&fn),
	)
	return
}

// GetAuthToken calls the function "WEBEXT.identity.getAuthToken" directly.
func GetAuthToken(details TokenDetails) (ret js.Promise[GetAuthTokenResult]) {
	bindings.CallGetAuthToken(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetAuthToken calls the function "WEBEXT.identity.getAuthToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAuthToken(details TokenDetails) (ret js.Promise[GetAuthTokenResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAuthToken(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetProfileUserInfo returns true if the function "WEBEXT.identity.getProfileUserInfo" exists.
func HasFuncGetProfileUserInfo() bool {
	return js.True == bindings.HasFuncGetProfileUserInfo()
}

// FuncGetProfileUserInfo returns the function "WEBEXT.identity.getProfileUserInfo".
func FuncGetProfileUserInfo() (fn js.Func[func(details ProfileDetails) js.Promise[ProfileUserInfo]]) {
	bindings.FuncGetProfileUserInfo(
		js.Pointer(&fn),
	)
	return
}

// GetProfileUserInfo calls the function "WEBEXT.identity.getProfileUserInfo" directly.
func GetProfileUserInfo(details ProfileDetails) (ret js.Promise[ProfileUserInfo]) {
	bindings.CallGetProfileUserInfo(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetProfileUserInfo calls the function "WEBEXT.identity.getProfileUserInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProfileUserInfo(details ProfileDetails) (ret js.Promise[ProfileUserInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProfileUserInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetRedirectURL returns true if the function "WEBEXT.identity.getRedirectURL" exists.
func HasFuncGetRedirectURL() bool {
	return js.True == bindings.HasFuncGetRedirectURL()
}

// FuncGetRedirectURL returns the function "WEBEXT.identity.getRedirectURL".
func FuncGetRedirectURL() (fn js.Func[func(path js.String) js.String]) {
	bindings.FuncGetRedirectURL(
		js.Pointer(&fn),
	)
	return
}

// GetRedirectURL calls the function "WEBEXT.identity.getRedirectURL" directly.
func GetRedirectURL(path js.String) (ret js.String) {
	bindings.CallGetRedirectURL(
		js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetRedirectURL calls the function "WEBEXT.identity.getRedirectURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRedirectURL(path js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRedirectURL(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncLaunchWebAuthFlow returns true if the function "WEBEXT.identity.launchWebAuthFlow" exists.
func HasFuncLaunchWebAuthFlow() bool {
	return js.True == bindings.HasFuncLaunchWebAuthFlow()
}

// FuncLaunchWebAuthFlow returns the function "WEBEXT.identity.launchWebAuthFlow".
func FuncLaunchWebAuthFlow() (fn js.Func[func(details WebAuthFlowDetails) js.Promise[js.String]]) {
	bindings.FuncLaunchWebAuthFlow(
		js.Pointer(&fn),
	)
	return
}

// LaunchWebAuthFlow calls the function "WEBEXT.identity.launchWebAuthFlow" directly.
func LaunchWebAuthFlow(details WebAuthFlowDetails) (ret js.Promise[js.String]) {
	bindings.CallLaunchWebAuthFlow(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryLaunchWebAuthFlow calls the function "WEBEXT.identity.launchWebAuthFlow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchWebAuthFlow(details WebAuthFlowDetails) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchWebAuthFlow(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

type OnSignInChangedEventCallbackFunc func(this js.Ref, account *AccountInfo, signedIn bool) js.Ref

func (fn OnSignInChangedEventCallbackFunc) Register() js.Func[func(account *AccountInfo, signedIn bool)] {
	return js.RegisterCallback[func(account *AccountInfo, signedIn bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSignInChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AccountInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSignInChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, account *AccountInfo, signedIn bool) js.Ref
	Arg T
}

func (cb *OnSignInChangedEventCallback[T]) Register() js.Func[func(account *AccountInfo, signedIn bool)] {
	return js.RegisterCallback[func(account *AccountInfo, signedIn bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSignInChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AccountInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSignInChanged returns true if the function "WEBEXT.identity.onSignInChanged.addListener" exists.
func HasFuncOnSignInChanged() bool {
	return js.True == bindings.HasFuncOnSignInChanged()
}

// FuncOnSignInChanged returns the function "WEBEXT.identity.onSignInChanged.addListener".
func FuncOnSignInChanged() (fn js.Func[func(callback js.Func[func(account *AccountInfo, signedIn bool)])]) {
	bindings.FuncOnSignInChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSignInChanged calls the function "WEBEXT.identity.onSignInChanged.addListener" directly.
func OnSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret js.Void) {
	bindings.CallOnSignInChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSignInChanged calls the function "WEBEXT.identity.onSignInChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSignInChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSignInChanged returns true if the function "WEBEXT.identity.onSignInChanged.removeListener" exists.
func HasFuncOffSignInChanged() bool {
	return js.True == bindings.HasFuncOffSignInChanged()
}

// FuncOffSignInChanged returns the function "WEBEXT.identity.onSignInChanged.removeListener".
func FuncOffSignInChanged() (fn js.Func[func(callback js.Func[func(account *AccountInfo, signedIn bool)])]) {
	bindings.FuncOffSignInChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSignInChanged calls the function "WEBEXT.identity.onSignInChanged.removeListener" directly.
func OffSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret js.Void) {
	bindings.CallOffSignInChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSignInChanged calls the function "WEBEXT.identity.onSignInChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSignInChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSignInChanged returns true if the function "WEBEXT.identity.onSignInChanged.hasListener" exists.
func HasFuncHasOnSignInChanged() bool {
	return js.True == bindings.HasFuncHasOnSignInChanged()
}

// FuncHasOnSignInChanged returns the function "WEBEXT.identity.onSignInChanged.hasListener".
func FuncHasOnSignInChanged() (fn js.Func[func(callback js.Func[func(account *AccountInfo, signedIn bool)]) bool]) {
	bindings.FuncHasOnSignInChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSignInChanged calls the function "WEBEXT.identity.onSignInChanged.hasListener" directly.
func HasOnSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret bool) {
	bindings.CallHasOnSignInChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSignInChanged calls the function "WEBEXT.identity.onSignInChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSignInChanged(callback js.Func[func(account *AccountInfo, signedIn bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSignInChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveCachedAuthToken returns true if the function "WEBEXT.identity.removeCachedAuthToken" exists.
func HasFuncRemoveCachedAuthToken() bool {
	return js.True == bindings.HasFuncRemoveCachedAuthToken()
}

// FuncRemoveCachedAuthToken returns the function "WEBEXT.identity.removeCachedAuthToken".
func FuncRemoveCachedAuthToken() (fn js.Func[func(details InvalidTokenDetails) js.Promise[js.Void]]) {
	bindings.FuncRemoveCachedAuthToken(
		js.Pointer(&fn),
	)
	return
}

// RemoveCachedAuthToken calls the function "WEBEXT.identity.removeCachedAuthToken" directly.
func RemoveCachedAuthToken(details InvalidTokenDetails) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCachedAuthToken(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryRemoveCachedAuthToken calls the function "WEBEXT.identity.removeCachedAuthToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCachedAuthToken(details InvalidTokenDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCachedAuthToken(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package login

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/login/bindings"
)

type SamlUserSessionProperties struct {
	// Email is "SamlUserSessionProperties.email"
	//
	// Optional
	Email js.String
	// GaiaId is "SamlUserSessionProperties.gaiaId"
	//
	// Optional
	GaiaId js.String
	// Password is "SamlUserSessionProperties.password"
	//
	// Optional
	Password js.String
	// OauthCode is "SamlUserSessionProperties.oauthCode"
	//
	// Optional
	OauthCode js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SamlUserSessionProperties with all fields set.
func (p SamlUserSessionProperties) FromRef(ref js.Ref) SamlUserSessionProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SamlUserSessionProperties in the application heap.
func (p SamlUserSessionProperties) New() js.Ref {
	return bindings.SamlUserSessionPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SamlUserSessionProperties) UpdateFrom(ref js.Ref) {
	bindings.SamlUserSessionPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SamlUserSessionProperties) Update(ref js.Ref) {
	bindings.SamlUserSessionPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SamlUserSessionProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.GaiaId.Ref(),
		p.Password.Ref(),
		p.OauthCode.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.GaiaId = p.GaiaId.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.OauthCode = p.OauthCode.FromRef(js.Undefined)
}

type StringCallbackFunc func(this js.Ref, result js.String) js.Ref

func (fn StringCallbackFunc) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StringCallbackFunc) DispatchCallback(
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

type StringCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.String) js.Ref
	Arg T
}

func (cb *StringCallback[T]) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StringCallback[T]) DispatchCallback(
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

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncEndSharedSession returns true if the function "WEBEXT.login.endSharedSession" exists.
func HasFuncEndSharedSession() bool {
	return js.True == bindings.HasFuncEndSharedSession()
}

// FuncEndSharedSession returns the function "WEBEXT.login.endSharedSession".
func FuncEndSharedSession() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncEndSharedSession(
		js.Pointer(&fn),
	)
	return
}

// EndSharedSession calls the function "WEBEXT.login.endSharedSession" directly.
func EndSharedSession() (ret js.Promise[js.Void]) {
	bindings.CallEndSharedSession(
		js.Pointer(&ret),
	)

	return
}

// TryEndSharedSession calls the function "WEBEXT.login.endSharedSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEndSharedSession() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEndSharedSession(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEnterSharedSession returns true if the function "WEBEXT.login.enterSharedSession" exists.
func HasFuncEnterSharedSession() bool {
	return js.True == bindings.HasFuncEnterSharedSession()
}

// FuncEnterSharedSession returns the function "WEBEXT.login.enterSharedSession".
func FuncEnterSharedSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncEnterSharedSession(
		js.Pointer(&fn),
	)
	return
}

// EnterSharedSession calls the function "WEBEXT.login.enterSharedSession" directly.
func EnterSharedSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallEnterSharedSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryEnterSharedSession calls the function "WEBEXT.login.enterSharedSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnterSharedSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnterSharedSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

// HasFuncExitCurrentSession returns true if the function "WEBEXT.login.exitCurrentSession" exists.
func HasFuncExitCurrentSession() bool {
	return js.True == bindings.HasFuncExitCurrentSession()
}

// FuncExitCurrentSession returns the function "WEBEXT.login.exitCurrentSession".
func FuncExitCurrentSession() (fn js.Func[func(dataForNextLoginAttempt js.String) js.Promise[js.Void]]) {
	bindings.FuncExitCurrentSession(
		js.Pointer(&fn),
	)
	return
}

// ExitCurrentSession calls the function "WEBEXT.login.exitCurrentSession" directly.
func ExitCurrentSession(dataForNextLoginAttempt js.String) (ret js.Promise[js.Void]) {
	bindings.CallExitCurrentSession(
		js.Pointer(&ret),
		dataForNextLoginAttempt.Ref(),
	)

	return
}

// TryExitCurrentSession calls the function "WEBEXT.login.exitCurrentSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExitCurrentSession(dataForNextLoginAttempt js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExitCurrentSession(
		js.Pointer(&ret), js.Pointer(&exception),
		dataForNextLoginAttempt.Ref(),
	)

	return
}

// HasFuncFetchDataForNextLoginAttempt returns true if the function "WEBEXT.login.fetchDataForNextLoginAttempt" exists.
func HasFuncFetchDataForNextLoginAttempt() bool {
	return js.True == bindings.HasFuncFetchDataForNextLoginAttempt()
}

// FuncFetchDataForNextLoginAttempt returns the function "WEBEXT.login.fetchDataForNextLoginAttempt".
func FuncFetchDataForNextLoginAttempt() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncFetchDataForNextLoginAttempt(
		js.Pointer(&fn),
	)
	return
}

// FetchDataForNextLoginAttempt calls the function "WEBEXT.login.fetchDataForNextLoginAttempt" directly.
func FetchDataForNextLoginAttempt() (ret js.Promise[js.String]) {
	bindings.CallFetchDataForNextLoginAttempt(
		js.Pointer(&ret),
	)

	return
}

// TryFetchDataForNextLoginAttempt calls the function "WEBEXT.login.fetchDataForNextLoginAttempt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFetchDataForNextLoginAttempt() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFetchDataForNextLoginAttempt(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLaunchManagedGuestSession returns true if the function "WEBEXT.login.launchManagedGuestSession" exists.
func HasFuncLaunchManagedGuestSession() bool {
	return js.True == bindings.HasFuncLaunchManagedGuestSession()
}

// FuncLaunchManagedGuestSession returns the function "WEBEXT.login.launchManagedGuestSession".
func FuncLaunchManagedGuestSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchManagedGuestSession(
		js.Pointer(&fn),
	)
	return
}

// LaunchManagedGuestSession calls the function "WEBEXT.login.launchManagedGuestSession" directly.
func LaunchManagedGuestSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchManagedGuestSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryLaunchManagedGuestSession calls the function "WEBEXT.login.launchManagedGuestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchManagedGuestSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchManagedGuestSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

// HasFuncLaunchSamlUserSession returns true if the function "WEBEXT.login.launchSamlUserSession" exists.
func HasFuncLaunchSamlUserSession() bool {
	return js.True == bindings.HasFuncLaunchSamlUserSession()
}

// FuncLaunchSamlUserSession returns the function "WEBEXT.login.launchSamlUserSession".
func FuncLaunchSamlUserSession() (fn js.Func[func(properties SamlUserSessionProperties) js.Promise[js.Void]]) {
	bindings.FuncLaunchSamlUserSession(
		js.Pointer(&fn),
	)
	return
}

// LaunchSamlUserSession calls the function "WEBEXT.login.launchSamlUserSession" directly.
func LaunchSamlUserSession(properties SamlUserSessionProperties) (ret js.Promise[js.Void]) {
	bindings.CallLaunchSamlUserSession(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryLaunchSamlUserSession calls the function "WEBEXT.login.launchSamlUserSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchSamlUserSession(properties SamlUserSessionProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchSamlUserSession(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncLaunchSharedManagedGuestSession returns true if the function "WEBEXT.login.launchSharedManagedGuestSession" exists.
func HasFuncLaunchSharedManagedGuestSession() bool {
	return js.True == bindings.HasFuncLaunchSharedManagedGuestSession()
}

// FuncLaunchSharedManagedGuestSession returns the function "WEBEXT.login.launchSharedManagedGuestSession".
func FuncLaunchSharedManagedGuestSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchSharedManagedGuestSession(
		js.Pointer(&fn),
	)
	return
}

// LaunchSharedManagedGuestSession calls the function "WEBEXT.login.launchSharedManagedGuestSession" directly.
func LaunchSharedManagedGuestSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchSharedManagedGuestSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryLaunchSharedManagedGuestSession calls the function "WEBEXT.login.launchSharedManagedGuestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchSharedManagedGuestSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchSharedManagedGuestSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

// HasFuncLockCurrentSession returns true if the function "WEBEXT.login.lockCurrentSession" exists.
func HasFuncLockCurrentSession() bool {
	return js.True == bindings.HasFuncLockCurrentSession()
}

// FuncLockCurrentSession returns the function "WEBEXT.login.lockCurrentSession".
func FuncLockCurrentSession() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncLockCurrentSession(
		js.Pointer(&fn),
	)
	return
}

// LockCurrentSession calls the function "WEBEXT.login.lockCurrentSession" directly.
func LockCurrentSession() (ret js.Promise[js.Void]) {
	bindings.CallLockCurrentSession(
		js.Pointer(&ret),
	)

	return
}

// TryLockCurrentSession calls the function "WEBEXT.login.lockCurrentSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLockCurrentSession() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockCurrentSession(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLockManagedGuestSession returns true if the function "WEBEXT.login.lockManagedGuestSession" exists.
func HasFuncLockManagedGuestSession() bool {
	return js.True == bindings.HasFuncLockManagedGuestSession()
}

// FuncLockManagedGuestSession returns the function "WEBEXT.login.lockManagedGuestSession".
func FuncLockManagedGuestSession() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncLockManagedGuestSession(
		js.Pointer(&fn),
	)
	return
}

// LockManagedGuestSession calls the function "WEBEXT.login.lockManagedGuestSession" directly.
func LockManagedGuestSession() (ret js.Promise[js.Void]) {
	bindings.CallLockManagedGuestSession(
		js.Pointer(&ret),
	)

	return
}

// TryLockManagedGuestSession calls the function "WEBEXT.login.lockManagedGuestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLockManagedGuestSession() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockManagedGuestSession(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNotifyExternalLogoutDone returns true if the function "WEBEXT.login.notifyExternalLogoutDone" exists.
func HasFuncNotifyExternalLogoutDone() bool {
	return js.True == bindings.HasFuncNotifyExternalLogoutDone()
}

// FuncNotifyExternalLogoutDone returns the function "WEBEXT.login.notifyExternalLogoutDone".
func FuncNotifyExternalLogoutDone() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNotifyExternalLogoutDone(
		js.Pointer(&fn),
	)
	return
}

// NotifyExternalLogoutDone calls the function "WEBEXT.login.notifyExternalLogoutDone" directly.
func NotifyExternalLogoutDone() (ret js.Promise[js.Void]) {
	bindings.CallNotifyExternalLogoutDone(
		js.Pointer(&ret),
	)

	return
}

// TryNotifyExternalLogoutDone calls the function "WEBEXT.login.notifyExternalLogoutDone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyExternalLogoutDone() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyExternalLogoutDone(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnExternalLogoutDoneEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnExternalLogoutDoneEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExternalLogoutDoneEventCallbackFunc) DispatchCallback(
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

type OnExternalLogoutDoneEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnExternalLogoutDoneEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExternalLogoutDoneEventCallback[T]) DispatchCallback(
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

// HasFuncOnExternalLogoutDone returns true if the function "WEBEXT.login.onExternalLogoutDone.addListener" exists.
func HasFuncOnExternalLogoutDone() bool {
	return js.True == bindings.HasFuncOnExternalLogoutDone()
}

// FuncOnExternalLogoutDone returns the function "WEBEXT.login.onExternalLogoutDone.addListener".
func FuncOnExternalLogoutDone() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnExternalLogoutDone(
		js.Pointer(&fn),
	)
	return
}

// OnExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.addListener" directly.
func OnExternalLogoutDone(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnExternalLogoutDone(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExternalLogoutDone(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExternalLogoutDone(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExternalLogoutDone returns true if the function "WEBEXT.login.onExternalLogoutDone.removeListener" exists.
func HasFuncOffExternalLogoutDone() bool {
	return js.True == bindings.HasFuncOffExternalLogoutDone()
}

// FuncOffExternalLogoutDone returns the function "WEBEXT.login.onExternalLogoutDone.removeListener".
func FuncOffExternalLogoutDone() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffExternalLogoutDone(
		js.Pointer(&fn),
	)
	return
}

// OffExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.removeListener" directly.
func OffExternalLogoutDone(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffExternalLogoutDone(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExternalLogoutDone(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExternalLogoutDone(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExternalLogoutDone returns true if the function "WEBEXT.login.onExternalLogoutDone.hasListener" exists.
func HasFuncHasOnExternalLogoutDone() bool {
	return js.True == bindings.HasFuncHasOnExternalLogoutDone()
}

// FuncHasOnExternalLogoutDone returns the function "WEBEXT.login.onExternalLogoutDone.hasListener".
func FuncHasOnExternalLogoutDone() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnExternalLogoutDone(
		js.Pointer(&fn),
	)
	return
}

// HasOnExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.hasListener" directly.
func HasOnExternalLogoutDone(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnExternalLogoutDone(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExternalLogoutDone calls the function "WEBEXT.login.onExternalLogoutDone.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExternalLogoutDone(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExternalLogoutDone(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRequestExternalLogoutEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnRequestExternalLogoutEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestExternalLogoutEventCallbackFunc) DispatchCallback(
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

type OnRequestExternalLogoutEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRequestExternalLogoutEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestExternalLogoutEventCallback[T]) DispatchCallback(
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

// HasFuncOnRequestExternalLogout returns true if the function "WEBEXT.login.onRequestExternalLogout.addListener" exists.
func HasFuncOnRequestExternalLogout() bool {
	return js.True == bindings.HasFuncOnRequestExternalLogout()
}

// FuncOnRequestExternalLogout returns the function "WEBEXT.login.onRequestExternalLogout.addListener".
func FuncOnRequestExternalLogout() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRequestExternalLogout(
		js.Pointer(&fn),
	)
	return
}

// OnRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.addListener" directly.
func OnRequestExternalLogout(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRequestExternalLogout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRequestExternalLogout(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRequestExternalLogout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRequestExternalLogout returns true if the function "WEBEXT.login.onRequestExternalLogout.removeListener" exists.
func HasFuncOffRequestExternalLogout() bool {
	return js.True == bindings.HasFuncOffRequestExternalLogout()
}

// FuncOffRequestExternalLogout returns the function "WEBEXT.login.onRequestExternalLogout.removeListener".
func FuncOffRequestExternalLogout() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRequestExternalLogout(
		js.Pointer(&fn),
	)
	return
}

// OffRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.removeListener" directly.
func OffRequestExternalLogout(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRequestExternalLogout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRequestExternalLogout(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRequestExternalLogout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRequestExternalLogout returns true if the function "WEBEXT.login.onRequestExternalLogout.hasListener" exists.
func HasFuncHasOnRequestExternalLogout() bool {
	return js.True == bindings.HasFuncHasOnRequestExternalLogout()
}

// FuncHasOnRequestExternalLogout returns the function "WEBEXT.login.onRequestExternalLogout.hasListener".
func FuncHasOnRequestExternalLogout() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRequestExternalLogout(
		js.Pointer(&fn),
	)
	return
}

// HasOnRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.hasListener" directly.
func HasOnRequestExternalLogout(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRequestExternalLogout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRequestExternalLogout calls the function "WEBEXT.login.onRequestExternalLogout.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRequestExternalLogout(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRequestExternalLogout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRequestExternalLogout returns true if the function "WEBEXT.login.requestExternalLogout" exists.
func HasFuncRequestExternalLogout() bool {
	return js.True == bindings.HasFuncRequestExternalLogout()
}

// FuncRequestExternalLogout returns the function "WEBEXT.login.requestExternalLogout".
func FuncRequestExternalLogout() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRequestExternalLogout(
		js.Pointer(&fn),
	)
	return
}

// RequestExternalLogout calls the function "WEBEXT.login.requestExternalLogout" directly.
func RequestExternalLogout() (ret js.Promise[js.Void]) {
	bindings.CallRequestExternalLogout(
		js.Pointer(&ret),
	)

	return
}

// TryRequestExternalLogout calls the function "WEBEXT.login.requestExternalLogout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestExternalLogout() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestExternalLogout(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetDataForNextLoginAttempt returns true if the function "WEBEXT.login.setDataForNextLoginAttempt" exists.
func HasFuncSetDataForNextLoginAttempt() bool {
	return js.True == bindings.HasFuncSetDataForNextLoginAttempt()
}

// FuncSetDataForNextLoginAttempt returns the function "WEBEXT.login.setDataForNextLoginAttempt".
func FuncSetDataForNextLoginAttempt() (fn js.Func[func(dataForNextLoginAttempt js.String) js.Promise[js.Void]]) {
	bindings.FuncSetDataForNextLoginAttempt(
		js.Pointer(&fn),
	)
	return
}

// SetDataForNextLoginAttempt calls the function "WEBEXT.login.setDataForNextLoginAttempt" directly.
func SetDataForNextLoginAttempt(dataForNextLoginAttempt js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetDataForNextLoginAttempt(
		js.Pointer(&ret),
		dataForNextLoginAttempt.Ref(),
	)

	return
}

// TrySetDataForNextLoginAttempt calls the function "WEBEXT.login.setDataForNextLoginAttempt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDataForNextLoginAttempt(dataForNextLoginAttempt js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDataForNextLoginAttempt(
		js.Pointer(&ret), js.Pointer(&exception),
		dataForNextLoginAttempt.Ref(),
	)

	return
}

// HasFuncUnlockCurrentSession returns true if the function "WEBEXT.login.unlockCurrentSession" exists.
func HasFuncUnlockCurrentSession() bool {
	return js.True == bindings.HasFuncUnlockCurrentSession()
}

// FuncUnlockCurrentSession returns the function "WEBEXT.login.unlockCurrentSession".
func FuncUnlockCurrentSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncUnlockCurrentSession(
		js.Pointer(&fn),
	)
	return
}

// UnlockCurrentSession calls the function "WEBEXT.login.unlockCurrentSession" directly.
func UnlockCurrentSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallUnlockCurrentSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryUnlockCurrentSession calls the function "WEBEXT.login.unlockCurrentSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnlockCurrentSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnlockCurrentSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

// HasFuncUnlockManagedGuestSession returns true if the function "WEBEXT.login.unlockManagedGuestSession" exists.
func HasFuncUnlockManagedGuestSession() bool {
	return js.True == bindings.HasFuncUnlockManagedGuestSession()
}

// FuncUnlockManagedGuestSession returns the function "WEBEXT.login.unlockManagedGuestSession".
func FuncUnlockManagedGuestSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncUnlockManagedGuestSession(
		js.Pointer(&fn),
	)
	return
}

// UnlockManagedGuestSession calls the function "WEBEXT.login.unlockManagedGuestSession" directly.
func UnlockManagedGuestSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallUnlockManagedGuestSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryUnlockManagedGuestSession calls the function "WEBEXT.login.unlockManagedGuestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnlockManagedGuestSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnlockManagedGuestSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

// HasFuncUnlockSharedSession returns true if the function "WEBEXT.login.unlockSharedSession" exists.
func HasFuncUnlockSharedSession() bool {
	return js.True == bindings.HasFuncUnlockSharedSession()
}

// FuncUnlockSharedSession returns the function "WEBEXT.login.unlockSharedSession".
func FuncUnlockSharedSession() (fn js.Func[func(password js.String) js.Promise[js.Void]]) {
	bindings.FuncUnlockSharedSession(
		js.Pointer(&fn),
	)
	return
}

// UnlockSharedSession calls the function "WEBEXT.login.unlockSharedSession" directly.
func UnlockSharedSession(password js.String) (ret js.Promise[js.Void]) {
	bindings.CallUnlockSharedSession(
		js.Pointer(&ret),
		password.Ref(),
	)

	return
}

// TryUnlockSharedSession calls the function "WEBEXT.login.unlockSharedSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnlockSharedSession(password js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnlockSharedSession(
		js.Pointer(&ret), js.Pointer(&exception),
		password.Ref(),
	)

	return
}

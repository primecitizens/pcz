// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package usersprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/usersprivate/bindings"
)

type IsUserInListCallbackFunc func(this js.Ref, found bool) js.Ref

func (fn IsUserInListCallbackFunc) Register() js.Func[func(found bool)] {
	return js.RegisterCallback[func(found bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsUserInListCallbackFunc) DispatchCallback(
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

type IsUserInListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, found bool) js.Ref
	Arg T
}

func (cb *IsUserInListCallback[T]) Register() js.Func[func(found bool)] {
	return js.RegisterCallback[func(found bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsUserInListCallback[T]) DispatchCallback(
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

type LoginStatusCallbackFunc func(this js.Ref, status *LoginStatusDict) js.Ref

func (fn LoginStatusCallbackFunc) Register() js.Func[func(status *LoginStatusDict)] {
	return js.RegisterCallback[func(status *LoginStatusDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoginStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoginStatusDict
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

type LoginStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *LoginStatusDict) js.Ref
	Arg T
}

func (cb *LoginStatusCallback[T]) Register() js.Func[func(status *LoginStatusDict)] {
	return js.RegisterCallback[func(status *LoginStatusDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoginStatusCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoginStatusDict
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

type LoginStatusDict struct {
	// IsLoggedIn is "LoginStatusDict.isLoggedIn"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLoggedIn MUST be set to true to make this field effective.
	IsLoggedIn bool
	// IsScreenLocked is "LoginStatusDict.isScreenLocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsScreenLocked MUST be set to true to make this field effective.
	IsScreenLocked bool

	FFI_USE_IsLoggedIn     bool // for IsLoggedIn.
	FFI_USE_IsScreenLocked bool // for IsScreenLocked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoginStatusDict with all fields set.
func (p LoginStatusDict) FromRef(ref js.Ref) LoginStatusDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoginStatusDict in the application heap.
func (p LoginStatusDict) New() js.Ref {
	return bindings.LoginStatusDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoginStatusDict) UpdateFrom(ref js.Ref) {
	bindings.LoginStatusDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoginStatusDict) Update(ref js.Ref) {
	bindings.LoginStatusDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoginStatusDict) FreeMembers(recursive bool) {
}

type ManagedCallbackFunc func(this js.Ref, managed bool) js.Ref

func (fn ManagedCallbackFunc) Register() js.Func[func(managed bool)] {
	return js.RegisterCallback[func(managed bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ManagedCallbackFunc) DispatchCallback(
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

type ManagedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, managed bool) js.Ref
	Arg T
}

func (cb *ManagedCallback[T]) Register() js.Func[func(managed bool)] {
	return js.RegisterCallback[func(managed bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ManagedCallback[T]) DispatchCallback(
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

type User struct {
	// Email is "User.email"
	//
	// Optional
	Email js.String
	// DisplayEmail is "User.displayEmail"
	//
	// Optional
	DisplayEmail js.String
	// Name is "User.name"
	//
	// Optional
	Name js.String
	// IsOwner is "User.isOwner"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsOwner MUST be set to true to make this field effective.
	IsOwner bool
	// IsChild is "User.isChild"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsChild MUST be set to true to make this field effective.
	IsChild bool

	FFI_USE_IsOwner bool // for IsOwner.
	FFI_USE_IsChild bool // for IsChild.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a User with all fields set.
func (p User) FromRef(ref js.Ref) User {
	p.UpdateFrom(ref)
	return p
}

// New creates a new User in the application heap.
func (p User) New() js.Ref {
	return bindings.UserJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *User) UpdateFrom(ref js.Ref) {
	bindings.UserJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *User) Update(ref js.Ref) {
	bindings.UserJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *User) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.DisplayEmail.Ref(),
		p.Name.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.DisplayEmail = p.DisplayEmail.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type UserAddedCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn UserAddedCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UserAddedCallbackFunc) DispatchCallback(
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

type UserAddedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *UserAddedCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UserAddedCallback[T]) DispatchCallback(
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

type UserCallbackFunc func(this js.Ref, user *User) js.Ref

func (fn UserCallbackFunc) Register() js.Func[func(user *User)] {
	return js.RegisterCallback[func(user *User)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UserCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 User
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

type UserCallback[T any] struct {
	Fn  func(arg T, this js.Ref, user *User) js.Ref
	Arg T
}

func (cb *UserCallback[T]) Register() js.Func[func(user *User)] {
	return js.RegisterCallback[func(user *User)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UserCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 User
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

type UserRemovedCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn UserRemovedCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UserRemovedCallbackFunc) DispatchCallback(
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

type UserRemovedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *UserRemovedCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UserRemovedCallback[T]) DispatchCallback(
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

type UsersCallbackFunc func(this js.Ref, users js.Array[User]) js.Ref

func (fn UsersCallbackFunc) Register() js.Func[func(users js.Array[User])] {
	return js.RegisterCallback[func(users js.Array[User])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UsersCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[User]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UsersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, users js.Array[User]) js.Ref
	Arg T
}

func (cb *UsersCallback[T]) Register() js.Func[func(users js.Array[User])] {
	return js.RegisterCallback[func(users js.Array[User])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UsersCallback[T]) DispatchCallback(
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

		js.Array[User]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncAddUser returns true if the function "WEBEXT.usersPrivate.addUser" exists.
func HasFuncAddUser() bool {
	return js.True == bindings.HasFuncAddUser()
}

// FuncAddUser returns the function "WEBEXT.usersPrivate.addUser".
func FuncAddUser() (fn js.Func[func(email js.String) js.Promise[js.Boolean]]) {
	bindings.FuncAddUser(
		js.Pointer(&fn),
	)
	return
}

// AddUser calls the function "WEBEXT.usersPrivate.addUser" directly.
func AddUser(email js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallAddUser(
		js.Pointer(&ret),
		email.Ref(),
	)

	return
}

// TryAddUser calls the function "WEBEXT.usersPrivate.addUser"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddUser(email js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddUser(
		js.Pointer(&ret), js.Pointer(&exception),
		email.Ref(),
	)

	return
}

// HasFuncGetCurrentUser returns true if the function "WEBEXT.usersPrivate.getCurrentUser" exists.
func HasFuncGetCurrentUser() bool {
	return js.True == bindings.HasFuncGetCurrentUser()
}

// FuncGetCurrentUser returns the function "WEBEXT.usersPrivate.getCurrentUser".
func FuncGetCurrentUser() (fn js.Func[func() js.Promise[User]]) {
	bindings.FuncGetCurrentUser(
		js.Pointer(&fn),
	)
	return
}

// GetCurrentUser calls the function "WEBEXT.usersPrivate.getCurrentUser" directly.
func GetCurrentUser() (ret js.Promise[User]) {
	bindings.CallGetCurrentUser(
		js.Pointer(&ret),
	)

	return
}

// TryGetCurrentUser calls the function "WEBEXT.usersPrivate.getCurrentUser"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCurrentUser() (ret js.Promise[User], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCurrentUser(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLoginStatus returns true if the function "WEBEXT.usersPrivate.getLoginStatus" exists.
func HasFuncGetLoginStatus() bool {
	return js.True == bindings.HasFuncGetLoginStatus()
}

// FuncGetLoginStatus returns the function "WEBEXT.usersPrivate.getLoginStatus".
func FuncGetLoginStatus() (fn js.Func[func() js.Promise[LoginStatusDict]]) {
	bindings.FuncGetLoginStatus(
		js.Pointer(&fn),
	)
	return
}

// GetLoginStatus calls the function "WEBEXT.usersPrivate.getLoginStatus" directly.
func GetLoginStatus() (ret js.Promise[LoginStatusDict]) {
	bindings.CallGetLoginStatus(
		js.Pointer(&ret),
	)

	return
}

// TryGetLoginStatus calls the function "WEBEXT.usersPrivate.getLoginStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLoginStatus() (ret js.Promise[LoginStatusDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLoginStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUsers returns true if the function "WEBEXT.usersPrivate.getUsers" exists.
func HasFuncGetUsers() bool {
	return js.True == bindings.HasFuncGetUsers()
}

// FuncGetUsers returns the function "WEBEXT.usersPrivate.getUsers".
func FuncGetUsers() (fn js.Func[func() js.Promise[js.Array[User]]]) {
	bindings.FuncGetUsers(
		js.Pointer(&fn),
	)
	return
}

// GetUsers calls the function "WEBEXT.usersPrivate.getUsers" directly.
func GetUsers() (ret js.Promise[js.Array[User]]) {
	bindings.CallGetUsers(
		js.Pointer(&ret),
	)

	return
}

// TryGetUsers calls the function "WEBEXT.usersPrivate.getUsers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUsers() (ret js.Promise[js.Array[User]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUsers(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsUserInList returns true if the function "WEBEXT.usersPrivate.isUserInList" exists.
func HasFuncIsUserInList() bool {
	return js.True == bindings.HasFuncIsUserInList()
}

// FuncIsUserInList returns the function "WEBEXT.usersPrivate.isUserInList".
func FuncIsUserInList() (fn js.Func[func(email js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsUserInList(
		js.Pointer(&fn),
	)
	return
}

// IsUserInList calls the function "WEBEXT.usersPrivate.isUserInList" directly.
func IsUserInList(email js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsUserInList(
		js.Pointer(&ret),
		email.Ref(),
	)

	return
}

// TryIsUserInList calls the function "WEBEXT.usersPrivate.isUserInList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsUserInList(email js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsUserInList(
		js.Pointer(&ret), js.Pointer(&exception),
		email.Ref(),
	)

	return
}

// HasFuncIsUserListManaged returns true if the function "WEBEXT.usersPrivate.isUserListManaged" exists.
func HasFuncIsUserListManaged() bool {
	return js.True == bindings.HasFuncIsUserListManaged()
}

// FuncIsUserListManaged returns the function "WEBEXT.usersPrivate.isUserListManaged".
func FuncIsUserListManaged() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsUserListManaged(
		js.Pointer(&fn),
	)
	return
}

// IsUserListManaged calls the function "WEBEXT.usersPrivate.isUserListManaged" directly.
func IsUserListManaged() (ret js.Promise[js.Boolean]) {
	bindings.CallIsUserListManaged(
		js.Pointer(&ret),
	)

	return
}

// TryIsUserListManaged calls the function "WEBEXT.usersPrivate.isUserListManaged"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsUserListManaged() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsUserListManaged(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveUser returns true if the function "WEBEXT.usersPrivate.removeUser" exists.
func HasFuncRemoveUser() bool {
	return js.True == bindings.HasFuncRemoveUser()
}

// FuncRemoveUser returns the function "WEBEXT.usersPrivate.removeUser".
func FuncRemoveUser() (fn js.Func[func(email js.String) js.Promise[js.Boolean]]) {
	bindings.FuncRemoveUser(
		js.Pointer(&fn),
	)
	return
}

// RemoveUser calls the function "WEBEXT.usersPrivate.removeUser" directly.
func RemoveUser(email js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallRemoveUser(
		js.Pointer(&ret),
		email.Ref(),
	)

	return
}

// TryRemoveUser calls the function "WEBEXT.usersPrivate.removeUser"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveUser(email js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveUser(
		js.Pointer(&ret), js.Pointer(&exception),
		email.Ref(),
	)

	return
}

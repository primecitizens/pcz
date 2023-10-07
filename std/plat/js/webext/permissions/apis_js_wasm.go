// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package permissions

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/permissions/bindings"
)

type Permissions struct {
	// Origins is "Permissions.origins"
	//
	// Optional
	Origins js.Array[js.String]
	// Permissions is "Permissions.permissions"
	//
	// Optional
	Permissions js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Permissions with all fields set.
func (p Permissions) FromRef(ref js.Ref) Permissions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Permissions in the application heap.
func (p Permissions) New() js.Ref {
	return bindings.PermissionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Permissions) UpdateFrom(ref js.Ref) {
	bindings.PermissionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Permissions) Update(ref js.Ref) {
	bindings.PermissionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Permissions) FreeMembers(recursive bool) {
	js.Free(
		p.Origins.Ref(),
		p.Permissions.Ref(),
	)
	p.Origins = p.Origins.FromRef(js.Undefined)
	p.Permissions = p.Permissions.FromRef(js.Undefined)
}

// HasFuncContains returns true if the function "WEBEXT.permissions.contains" exists.
func HasFuncContains() bool {
	return js.True == bindings.HasFuncContains()
}

// FuncContains returns the function "WEBEXT.permissions.contains".
func FuncContains() (fn js.Func[func(permissions Permissions) js.Promise[js.Boolean]]) {
	bindings.FuncContains(
		js.Pointer(&fn),
	)
	return
}

// Contains calls the function "WEBEXT.permissions.contains" directly.
func Contains(permissions Permissions) (ret js.Promise[js.Boolean]) {
	bindings.CallContains(
		js.Pointer(&ret),
		js.Pointer(&permissions),
	)

	return
}

// TryContains calls the function "WEBEXT.permissions.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContains(permissions Permissions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContains(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&permissions),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.permissions.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.permissions.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[Permissions]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.permissions.getAll" directly.
func GetAll() (ret js.Promise[Permissions]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.permissions.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[Permissions], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnAddedEventCallbackFunc func(this js.Ref, permissions *Permissions) js.Ref

func (fn OnAddedEventCallbackFunc) Register() js.Func[func(permissions *Permissions)] {
	return js.RegisterCallback[func(permissions *Permissions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Permissions
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

type OnAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, permissions *Permissions) js.Ref
	Arg T
}

func (cb *OnAddedEventCallback[T]) Register() js.Func[func(permissions *Permissions)] {
	return js.RegisterCallback[func(permissions *Permissions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Permissions
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

// HasFuncOnAdded returns true if the function "WEBEXT.permissions.onAdded.addListener" exists.
func HasFuncOnAdded() bool {
	return js.True == bindings.HasFuncOnAdded()
}

// FuncOnAdded returns the function "WEBEXT.permissions.onAdded.addListener".
func FuncOnAdded() (fn js.Func[func(callback js.Func[func(permissions *Permissions)])]) {
	bindings.FuncOnAdded(
		js.Pointer(&fn),
	)
	return
}

// OnAdded calls the function "WEBEXT.permissions.onAdded.addListener" directly.
func OnAdded(callback js.Func[func(permissions *Permissions)]) (ret js.Void) {
	bindings.CallOnAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAdded calls the function "WEBEXT.permissions.onAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAdded(callback js.Func[func(permissions *Permissions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAdded returns true if the function "WEBEXT.permissions.onAdded.removeListener" exists.
func HasFuncOffAdded() bool {
	return js.True == bindings.HasFuncOffAdded()
}

// FuncOffAdded returns the function "WEBEXT.permissions.onAdded.removeListener".
func FuncOffAdded() (fn js.Func[func(callback js.Func[func(permissions *Permissions)])]) {
	bindings.FuncOffAdded(
		js.Pointer(&fn),
	)
	return
}

// OffAdded calls the function "WEBEXT.permissions.onAdded.removeListener" directly.
func OffAdded(callback js.Func[func(permissions *Permissions)]) (ret js.Void) {
	bindings.CallOffAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAdded calls the function "WEBEXT.permissions.onAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAdded(callback js.Func[func(permissions *Permissions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAdded returns true if the function "WEBEXT.permissions.onAdded.hasListener" exists.
func HasFuncHasOnAdded() bool {
	return js.True == bindings.HasFuncHasOnAdded()
}

// FuncHasOnAdded returns the function "WEBEXT.permissions.onAdded.hasListener".
func FuncHasOnAdded() (fn js.Func[func(callback js.Func[func(permissions *Permissions)]) bool]) {
	bindings.FuncHasOnAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnAdded calls the function "WEBEXT.permissions.onAdded.hasListener" directly.
func HasOnAdded(callback js.Func[func(permissions *Permissions)]) (ret bool) {
	bindings.CallHasOnAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAdded calls the function "WEBEXT.permissions.onAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAdded(callback js.Func[func(permissions *Permissions)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemovedEventCallbackFunc func(this js.Ref, permissions *Permissions) js.Ref

func (fn OnRemovedEventCallbackFunc) Register() js.Func[func(permissions *Permissions)] {
	return js.RegisterCallback[func(permissions *Permissions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Permissions
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

type OnRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, permissions *Permissions) js.Ref
	Arg T
}

func (cb *OnRemovedEventCallback[T]) Register() js.Func[func(permissions *Permissions)] {
	return js.RegisterCallback[func(permissions *Permissions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Permissions
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

// HasFuncOnRemoved returns true if the function "WEBEXT.permissions.onRemoved.addListener" exists.
func HasFuncOnRemoved() bool {
	return js.True == bindings.HasFuncOnRemoved()
}

// FuncOnRemoved returns the function "WEBEXT.permissions.onRemoved.addListener".
func FuncOnRemoved() (fn js.Func[func(callback js.Func[func(permissions *Permissions)])]) {
	bindings.FuncOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnRemoved calls the function "WEBEXT.permissions.onRemoved.addListener" directly.
func OnRemoved(callback js.Func[func(permissions *Permissions)]) (ret js.Void) {
	bindings.CallOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoved calls the function "WEBEXT.permissions.onRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoved(callback js.Func[func(permissions *Permissions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoved returns true if the function "WEBEXT.permissions.onRemoved.removeListener" exists.
func HasFuncOffRemoved() bool {
	return js.True == bindings.HasFuncOffRemoved()
}

// FuncOffRemoved returns the function "WEBEXT.permissions.onRemoved.removeListener".
func FuncOffRemoved() (fn js.Func[func(callback js.Func[func(permissions *Permissions)])]) {
	bindings.FuncOffRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffRemoved calls the function "WEBEXT.permissions.onRemoved.removeListener" directly.
func OffRemoved(callback js.Func[func(permissions *Permissions)]) (ret js.Void) {
	bindings.CallOffRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoved calls the function "WEBEXT.permissions.onRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoved(callback js.Func[func(permissions *Permissions)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoved returns true if the function "WEBEXT.permissions.onRemoved.hasListener" exists.
func HasFuncHasOnRemoved() bool {
	return js.True == bindings.HasFuncHasOnRemoved()
}

// FuncHasOnRemoved returns the function "WEBEXT.permissions.onRemoved.hasListener".
func FuncHasOnRemoved() (fn js.Func[func(callback js.Func[func(permissions *Permissions)]) bool]) {
	bindings.FuncHasOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoved calls the function "WEBEXT.permissions.onRemoved.hasListener" directly.
func HasOnRemoved(callback js.Func[func(permissions *Permissions)]) (ret bool) {
	bindings.CallHasOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoved calls the function "WEBEXT.permissions.onRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoved(callback js.Func[func(permissions *Permissions)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.permissions.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.permissions.remove".
func FuncRemove() (fn js.Func[func(permissions Permissions) js.Promise[js.Boolean]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.permissions.remove" directly.
func Remove(permissions Permissions) (ret js.Promise[js.Boolean]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		js.Pointer(&permissions),
	)

	return
}

// TryRemove calls the function "WEBEXT.permissions.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(permissions Permissions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&permissions),
	)

	return
}

// HasFuncRequest returns true if the function "WEBEXT.permissions.request" exists.
func HasFuncRequest() bool {
	return js.True == bindings.HasFuncRequest()
}

// FuncRequest returns the function "WEBEXT.permissions.request".
func FuncRequest() (fn js.Func[func(permissions Permissions) js.Promise[js.Boolean]]) {
	bindings.FuncRequest(
		js.Pointer(&fn),
	)
	return
}

// Request calls the function "WEBEXT.permissions.request" directly.
func Request(permissions Permissions) (ret js.Promise[js.Boolean]) {
	bindings.CallRequest(
		js.Pointer(&ret),
		js.Pointer(&permissions),
	)

	return
}

// TryRequest calls the function "WEBEXT.permissions.request"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequest(permissions Permissions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&permissions),
	)

	return
}

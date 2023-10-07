// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package storage

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/storage/bindings"
)

type AccessLevel uint32

const (
	_ AccessLevel = iota

	AccessLevel_TRUSTED_CONTEXTS
	AccessLevel_TRUSTED_AND_UNTRUSTED_CONTEXTS
)

func (AccessLevel) FromRef(str js.Ref) AccessLevel {
	return AccessLevel(bindings.ConstOfAccessLevel(str))
}

func (x AccessLevel) String() (string, bool) {
	switch x {
	case AccessLevel_TRUSTED_CONTEXTS:
		return "TRUSTED_CONTEXTS", true
	case AccessLevel_TRUSTED_AND_UNTRUSTED_CONTEXTS:
		return "TRUSTED_AND_UNTRUSTED_CONTEXTS", true
	default:
		return "", false
	}
}

type SetAccessLevelArgAccessOptions struct {
	// AccessLevel is "SetAccessLevelArgAccessOptions.accessLevel"
	//
	// Required
	AccessLevel AccessLevel

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetAccessLevelArgAccessOptions with all fields set.
func (p SetAccessLevelArgAccessOptions) FromRef(ref js.Ref) SetAccessLevelArgAccessOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetAccessLevelArgAccessOptions in the application heap.
func (p SetAccessLevelArgAccessOptions) New() js.Ref {
	return bindings.SetAccessLevelArgAccessOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetAccessLevelArgAccessOptions) UpdateFrom(ref js.Ref) {
	bindings.SetAccessLevelArgAccessOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetAccessLevelArgAccessOptions) Update(ref js.Ref) {
	bindings.SetAccessLevelArgAccessOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetAccessLevelArgAccessOptions) FreeMembers(recursive bool) {
}

type OneOf_String_ArrayString_Any struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString_Any) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString_Any) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString_Any) FromRef(ref js.Ref) OneOf_String_ArrayString_Any {
	return OneOf_String_ArrayString_Any{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString_Any) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString_Any) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString_Any) Any() js.Any {
	return js.Any{}.FromRef(x.ref)
}

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type StorageArea struct {
	ref js.Ref
}

func (this StorageArea) Once() StorageArea {
	this.ref.Once()
	return this
}

func (this StorageArea) Ref() js.Ref {
	return this.ref
}

func (this StorageArea) FromRef(ref js.Ref) StorageArea {
	this.ref = ref
	return this
}

func (this StorageArea) Free() {
	this.ref.Free()
}

// HasFuncGet returns true if the method "StorageArea.get" exists.
func (this StorageArea) HasFuncGet() bool {
	return js.True == bindings.HasFuncStorageAreaGet(
		this.ref,
	)
}

// FuncGet returns the method "StorageArea.get".
func (this StorageArea) FuncGet() (fn js.Func[func(keys OneOf_String_ArrayString_Any) js.Promise[js.Any]]) {
	bindings.FuncStorageAreaGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "StorageArea.get".
func (this StorageArea) Get(keys OneOf_String_ArrayString_Any) (ret js.Promise[js.Any]) {
	bindings.CallStorageAreaGet(
		this.ref, js.Pointer(&ret),
		keys.Ref(),
	)

	return
}

// TryGet calls the method "StorageArea.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryGet(keys OneOf_String_ArrayString_Any) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keys.Ref(),
	)

	return
}

// HasFuncGet1 returns true if the method "StorageArea.get" exists.
func (this StorageArea) HasFuncGet1() bool {
	return js.True == bindings.HasFuncStorageAreaGet1(
		this.ref,
	)
}

// FuncGet1 returns the method "StorageArea.get".
func (this StorageArea) FuncGet1() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncStorageAreaGet1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get1 calls the method "StorageArea.get".
func (this StorageArea) Get1() (ret js.Promise[js.Any]) {
	bindings.CallStorageAreaGet1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGet1 calls the method "StorageArea.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryGet1() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaGet1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBytesInUse returns true if the method "StorageArea.getBytesInUse" exists.
func (this StorageArea) HasFuncGetBytesInUse() bool {
	return js.True == bindings.HasFuncStorageAreaGetBytesInUse(
		this.ref,
	)
}

// FuncGetBytesInUse returns the method "StorageArea.getBytesInUse".
func (this StorageArea) FuncGetBytesInUse() (fn js.Func[func(keys OneOf_String_ArrayString) js.Promise[js.BigInt[int64]]]) {
	bindings.FuncStorageAreaGetBytesInUse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBytesInUse calls the method "StorageArea.getBytesInUse".
func (this StorageArea) GetBytesInUse(keys OneOf_String_ArrayString) (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallStorageAreaGetBytesInUse(
		this.ref, js.Pointer(&ret),
		keys.Ref(),
	)

	return
}

// TryGetBytesInUse calls the method "StorageArea.getBytesInUse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryGetBytesInUse(keys OneOf_String_ArrayString) (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaGetBytesInUse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keys.Ref(),
	)

	return
}

// HasFuncGetBytesInUse1 returns true if the method "StorageArea.getBytesInUse" exists.
func (this StorageArea) HasFuncGetBytesInUse1() bool {
	return js.True == bindings.HasFuncStorageAreaGetBytesInUse1(
		this.ref,
	)
}

// FuncGetBytesInUse1 returns the method "StorageArea.getBytesInUse".
func (this StorageArea) FuncGetBytesInUse1() (fn js.Func[func() js.Promise[js.BigInt[int64]]]) {
	bindings.FuncStorageAreaGetBytesInUse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBytesInUse1 calls the method "StorageArea.getBytesInUse".
func (this StorageArea) GetBytesInUse1() (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallStorageAreaGetBytesInUse1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBytesInUse1 calls the method "StorageArea.getBytesInUse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryGetBytesInUse1() (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaGetBytesInUse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSet returns true if the method "StorageArea.set" exists.
func (this StorageArea) HasFuncSet() bool {
	return js.True == bindings.HasFuncStorageAreaSet(
		this.ref,
	)
}

// FuncSet returns the method "StorageArea.set".
func (this StorageArea) FuncSet() (fn js.Func[func(items js.Any) js.Promise[js.Void]]) {
	bindings.FuncStorageAreaSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "StorageArea.set".
func (this StorageArea) Set(items js.Any) (ret js.Promise[js.Void]) {
	bindings.CallStorageAreaSet(
		this.ref, js.Pointer(&ret),
		items.Ref(),
	)

	return
}

// TrySet calls the method "StorageArea.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TrySet(items js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		items.Ref(),
	)

	return
}

// HasFuncRemove returns true if the method "StorageArea.remove" exists.
func (this StorageArea) HasFuncRemove() bool {
	return js.True == bindings.HasFuncStorageAreaRemove(
		this.ref,
	)
}

// FuncRemove returns the method "StorageArea.remove".
func (this StorageArea) FuncRemove() (fn js.Func[func(keys OneOf_String_ArrayString) js.Promise[js.Void]]) {
	bindings.FuncStorageAreaRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "StorageArea.remove".
func (this StorageArea) Remove(keys OneOf_String_ArrayString) (ret js.Promise[js.Void]) {
	bindings.CallStorageAreaRemove(
		this.ref, js.Pointer(&ret),
		keys.Ref(),
	)

	return
}

// TryRemove calls the method "StorageArea.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryRemove(keys OneOf_String_ArrayString) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keys.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "StorageArea.clear" exists.
func (this StorageArea) HasFuncClear() bool {
	return js.True == bindings.HasFuncStorageAreaClear(
		this.ref,
	)
}

// FuncClear returns the method "StorageArea.clear".
func (this StorageArea) FuncClear() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStorageAreaClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "StorageArea.clear".
func (this StorageArea) Clear() (ret js.Promise[js.Void]) {
	bindings.CallStorageAreaClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "StorageArea.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TryClear() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetAccessLevel returns true if the method "StorageArea.setAccessLevel" exists.
func (this StorageArea) HasFuncSetAccessLevel() bool {
	return js.True == bindings.HasFuncStorageAreaSetAccessLevel(
		this.ref,
	)
}

// FuncSetAccessLevel returns the method "StorageArea.setAccessLevel".
func (this StorageArea) FuncSetAccessLevel() (fn js.Func[func(accessOptions SetAccessLevelArgAccessOptions) js.Promise[js.Void]]) {
	bindings.FuncStorageAreaSetAccessLevel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAccessLevel calls the method "StorageArea.setAccessLevel".
func (this StorageArea) SetAccessLevel(accessOptions SetAccessLevelArgAccessOptions) (ret js.Promise[js.Void]) {
	bindings.CallStorageAreaSetAccessLevel(
		this.ref, js.Pointer(&ret),
		js.Pointer(&accessOptions),
	)

	return
}

// TrySetAccessLevel calls the method "StorageArea.setAccessLevel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageArea) TrySetAccessLevel(accessOptions SetAccessLevelArgAccessOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageAreaSetAccessLevel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&accessOptions),
	)

	return
}

type StorageChange struct {
	// NewValue is "StorageChange.newValue"
	//
	// Optional
	NewValue js.Any
	// OldValue is "StorageChange.oldValue"
	//
	// Optional
	OldValue js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageChange with all fields set.
func (p StorageChange) FromRef(ref js.Ref) StorageChange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageChange in the application heap.
func (p StorageChange) New() js.Ref {
	return bindings.StorageChangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageChange) UpdateFrom(ref js.Ref) {
	bindings.StorageChangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageChange) Update(ref js.Ref) {
	bindings.StorageChangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageChange) FreeMembers(recursive bool) {
	js.Free(
		p.NewValue.Ref(),
		p.OldValue.Ref(),
	)
	p.NewValue = p.NewValue.FromRef(js.Undefined)
	p.OldValue = p.OldValue.FromRef(js.Undefined)
}

// Local returns the value of property "WEBEXT.storage.local".
//
// The returned bool will be false if there is no such property.
func Local() (ret StorageArea, ok bool) {
	ok = js.True == bindings.GetLocal(
		js.Pointer(&ret),
	)

	return
}

// SetLocal sets the value of property "WEBEXT.storage.local" to val.
//
// It returns false if the property cannot be set.
func SetLocal(val StorageArea) bool {
	return js.True == bindings.SetLocal(
		val.Ref())
}

// Managed returns the value of property "WEBEXT.storage.managed".
//
// The returned bool will be false if there is no such property.
func Managed() (ret StorageArea, ok bool) {
	ok = js.True == bindings.GetManaged(
		js.Pointer(&ret),
	)

	return
}

// SetManaged sets the value of property "WEBEXT.storage.managed" to val.
//
// It returns false if the property cannot be set.
func SetManaged(val StorageArea) bool {
	return js.True == bindings.SetManaged(
		val.Ref())
}

type OnChangedEventCallbackFunc func(this js.Ref, changes *StorageChange, areaName js.String) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func(changes *StorageChange, areaName js.String)] {
	return js.RegisterCallback[func(changes *StorageChange, areaName js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageChange
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, changes *StorageChange, areaName js.String) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func(changes *StorageChange, areaName js.String)] {
	return js.RegisterCallback[func(changes *StorageChange, areaName js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageChange
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnChanged returns true if the function "WEBEXT.storage.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.storage.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func(changes *StorageChange, areaName js.String)])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.storage.onChanged.addListener" directly.
func OnChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.storage.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.storage.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.storage.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func(changes *StorageChange, areaName js.String)])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.storage.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.storage.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.storage.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.storage.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func(changes *StorageChange, areaName js.String)]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.storage.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.storage.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func(changes *StorageChange, areaName js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// Session returns the value of property "WEBEXT.storage.session".
//
// The returned bool will be false if there is no such property.
func Session() (ret StorageArea, ok bool) {
	ok = js.True == bindings.GetSession(
		js.Pointer(&ret),
	)

	return
}

// SetSession sets the value of property "WEBEXT.storage.session" to val.
//
// It returns false if the property cannot be set.
func SetSession(val StorageArea) bool {
	return js.True == bindings.SetSession(
		val.Ref())
}

// Sync returns the value of property "WEBEXT.storage.sync".
//
// The returned bool will be false if there is no such property.
func Sync() (ret StorageArea, ok bool) {
	ok = js.True == bindings.GetSync(
		js.Pointer(&ret),
	)

	return
}

// SetSync sets the value of property "WEBEXT.storage.sync" to val.
//
// It returns false if the property cannot be set.
func SetSync(val StorageArea) bool {
	return js.True == bindings.SetSync(
		val.Ref())
}

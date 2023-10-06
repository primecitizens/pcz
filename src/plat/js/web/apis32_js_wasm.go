// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type SharedStorageWorklet struct {
	Worklet
}

func (this SharedStorageWorklet) Once() SharedStorageWorklet {
	this.Ref().Once()
	return this
}

func (this SharedStorageWorklet) Ref() js.Ref {
	return this.Worklet.Ref()
}

func (this SharedStorageWorklet) FromRef(ref js.Ref) SharedStorageWorklet {
	this.Worklet = this.Worklet.FromRef(ref)
	return this
}

func (this SharedStorageWorklet) Free() {
	this.Ref().Free()
}

type WindowSharedStorage struct {
	SharedStorage
}

func (this WindowSharedStorage) Once() WindowSharedStorage {
	this.Ref().Once()
	return this
}

func (this WindowSharedStorage) Ref() js.Ref {
	return this.SharedStorage.Ref()
}

func (this WindowSharedStorage) FromRef(ref js.Ref) WindowSharedStorage {
	this.SharedStorage = this.SharedStorage.FromRef(ref)
	return this
}

func (this WindowSharedStorage) Free() {
	this.Ref().Free()
}

// Worklet returns the value of property "WindowSharedStorage.worklet".
//
// It returns ok=false if there is no such property.
func (this WindowSharedStorage) Worklet() (ret SharedStorageWorklet, ok bool) {
	ok = js.True == bindings.GetWindowSharedStorageWorklet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRun returns true if the method "WindowSharedStorage.run" exists.
func (this WindowSharedStorage) HasRun() bool {
	return js.True == bindings.HasWindowSharedStorageRun(
		this.Ref(),
	)
}

// RunFunc returns the method "WindowSharedStorage.run".
func (this WindowSharedStorage) RunFunc() (fn js.Func[func(name js.String, options SharedStorageRunOperationMethodOptions) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageRunFunc(
			this.Ref(),
		),
	)
}

// Run calls the method "WindowSharedStorage.run".
func (this WindowSharedStorage) Run(name js.String, options SharedStorageRunOperationMethodOptions) (ret js.Promise[js.Any]) {
	bindings.CallWindowSharedStorageRun(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRun calls the method "WindowSharedStorage.run"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this WindowSharedStorage) TryRun(name js.String, options SharedStorageRunOperationMethodOptions) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSharedStorageRun(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRun1 returns true if the method "WindowSharedStorage.run" exists.
func (this WindowSharedStorage) HasRun1() bool {
	return js.True == bindings.HasWindowSharedStorageRun1(
		this.Ref(),
	)
}

// Run1Func returns the method "WindowSharedStorage.run".
func (this WindowSharedStorage) Run1Func() (fn js.Func[func(name js.String) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageRun1Func(
			this.Ref(),
		),
	)
}

// Run1 calls the method "WindowSharedStorage.run".
func (this WindowSharedStorage) Run1(name js.String) (ret js.Promise[js.Any]) {
	bindings.CallWindowSharedStorageRun1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRun1 calls the method "WindowSharedStorage.run"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this WindowSharedStorage) TryRun1(name js.String) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSharedStorageRun1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSelectURL returns true if the method "WindowSharedStorage.selectURL" exists.
func (this WindowSharedStorage) HasSelectURL() bool {
	return js.True == bindings.HasWindowSharedStorageSelectURL(
		this.Ref(),
	)
}

// SelectURLFunc returns the method "WindowSharedStorage.selectURL".
func (this WindowSharedStorage) SelectURLFunc() (fn js.Func[func(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata], options SharedStorageRunOperationMethodOptions) js.Promise[SharedStorageResponse]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageSelectURLFunc(
			this.Ref(),
		),
	)
}

// SelectURL calls the method "WindowSharedStorage.selectURL".
func (this WindowSharedStorage) SelectURL(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata], options SharedStorageRunOperationMethodOptions) (ret js.Promise[SharedStorageResponse]) {
	bindings.CallWindowSharedStorageSelectURL(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		urls.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySelectURL calls the method "WindowSharedStorage.selectURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this WindowSharedStorage) TrySelectURL(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata], options SharedStorageRunOperationMethodOptions) (ret js.Promise[SharedStorageResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSharedStorageSelectURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		urls.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSelectURL1 returns true if the method "WindowSharedStorage.selectURL" exists.
func (this WindowSharedStorage) HasSelectURL1() bool {
	return js.True == bindings.HasWindowSharedStorageSelectURL1(
		this.Ref(),
	)
}

// SelectURL1Func returns the method "WindowSharedStorage.selectURL".
func (this WindowSharedStorage) SelectURL1Func() (fn js.Func[func(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata]) js.Promise[SharedStorageResponse]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageSelectURL1Func(
			this.Ref(),
		),
	)
}

// SelectURL1 calls the method "WindowSharedStorage.selectURL".
func (this WindowSharedStorage) SelectURL1(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (ret js.Promise[SharedStorageResponse]) {
	bindings.CallWindowSharedStorageSelectURL1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		urls.Ref(),
	)

	return
}

// TrySelectURL1 calls the method "WindowSharedStorage.selectURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this WindowSharedStorage) TrySelectURL1(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (ret js.Promise[SharedStorageResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSharedStorageSelectURL1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		urls.Ref(),
	)

	return
}

type OneOf_Event_undefined struct {
	ref js.Ref
}

func (x OneOf_Event_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Event_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_Event_undefined) FromRef(ref js.Ref) OneOf_Event_undefined {
	return OneOf_Event_undefined{
		ref: ref,
	}
}

func (x OneOf_Event_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_Event_undefined) Event() Event {
	return Event{}.FromRef(x.ref)
}

type CookieSameSite uint32

const (
	_ CookieSameSite = iota

	CookieSameSite_STRICT
	CookieSameSite_LAX
	CookieSameSite_NONE
)

func (CookieSameSite) FromRef(str js.Ref) CookieSameSite {
	return CookieSameSite(bindings.ConstOfCookieSameSite(str))
}

func (x CookieSameSite) String() (string, bool) {
	switch x {
	case CookieSameSite_STRICT:
		return "strict", true
	case CookieSameSite_LAX:
		return "lax", true
	case CookieSameSite_NONE:
		return "none", true
	default:
		return "", false
	}
}

type CookieListItem struct {
	// Name is "CookieListItem.name"
	//
	// Optional
	Name js.String
	// Value is "CookieListItem.value"
	//
	// Optional
	Value js.String
	// Domain is "CookieListItem.domain"
	//
	// Optional
	Domain js.String
	// Path is "CookieListItem.path"
	//
	// Optional
	Path js.String
	// Expires is "CookieListItem.expires"
	//
	// Optional
	//
	// NOTE: FFI_USE_Expires MUST be set to true to make this field effective.
	Expires DOMHighResTimeStamp
	// Secure is "CookieListItem.secure"
	//
	// Optional
	//
	// NOTE: FFI_USE_Secure MUST be set to true to make this field effective.
	Secure bool
	// SameSite is "CookieListItem.sameSite"
	//
	// Optional
	SameSite CookieSameSite
	// Partitioned is "CookieListItem.partitioned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Partitioned MUST be set to true to make this field effective.
	Partitioned bool

	FFI_USE_Expires     bool // for Expires.
	FFI_USE_Secure      bool // for Secure.
	FFI_USE_Partitioned bool // for Partitioned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieListItem with all fields set.
func (p CookieListItem) FromRef(ref js.Ref) CookieListItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CookieListItem in the application heap.
func (p CookieListItem) New() js.Ref {
	return bindings.CookieListItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CookieListItem) UpdateFrom(ref js.Ref) {
	bindings.CookieListItemJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CookieListItem) Update(ref js.Ref) {
	bindings.CookieListItemJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CookieList = js.Array[CookieListItem]

type CookieInit struct {
	// Name is "CookieInit.name"
	//
	// Required
	Name js.String
	// Value is "CookieInit.value"
	//
	// Required
	Value js.String
	// Expires is "CookieInit.expires"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Expires MUST be set to true to make this field effective.
	Expires DOMHighResTimeStamp
	// Domain is "CookieInit.domain"
	//
	// Optional, defaults to null.
	Domain js.String
	// Path is "CookieInit.path"
	//
	// Optional, defaults to "/".
	Path js.String
	// SameSite is "CookieInit.sameSite"
	//
	// Optional, defaults to "strict".
	SameSite CookieSameSite
	// Partitioned is "CookieInit.partitioned"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Partitioned MUST be set to true to make this field effective.
	Partitioned bool

	FFI_USE_Expires     bool // for Expires.
	FFI_USE_Partitioned bool // for Partitioned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieInit with all fields set.
func (p CookieInit) FromRef(ref js.Ref) CookieInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CookieInit in the application heap.
func (p CookieInit) New() js.Ref {
	return bindings.CookieInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CookieInit) UpdateFrom(ref js.Ref) {
	bindings.CookieInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CookieInit) Update(ref js.Ref) {
	bindings.CookieInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CookieStoreDeleteOptions struct {
	// Name is "CookieStoreDeleteOptions.name"
	//
	// Required
	Name js.String
	// Domain is "CookieStoreDeleteOptions.domain"
	//
	// Optional, defaults to null.
	Domain js.String
	// Path is "CookieStoreDeleteOptions.path"
	//
	// Optional, defaults to "/".
	Path js.String
	// Partitioned is "CookieStoreDeleteOptions.partitioned"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Partitioned MUST be set to true to make this field effective.
	Partitioned bool

	FFI_USE_Partitioned bool // for Partitioned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieStoreDeleteOptions with all fields set.
func (p CookieStoreDeleteOptions) FromRef(ref js.Ref) CookieStoreDeleteOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CookieStoreDeleteOptions in the application heap.
func (p CookieStoreDeleteOptions) New() js.Ref {
	return bindings.CookieStoreDeleteOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CookieStoreDeleteOptions) UpdateFrom(ref js.Ref) {
	bindings.CookieStoreDeleteOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CookieStoreDeleteOptions) Update(ref js.Ref) {
	bindings.CookieStoreDeleteOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CookieStore struct {
	EventTarget
}

func (this CookieStore) Once() CookieStore {
	this.Ref().Once()
	return this
}

func (this CookieStore) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this CookieStore) FromRef(ref js.Ref) CookieStore {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this CookieStore) Free() {
	this.Ref().Free()
}

// HasGet returns true if the method "CookieStore.get" exists.
func (this CookieStore) HasGet() bool {
	return js.True == bindings.HasCookieStoreGet(
		this.Ref(),
	)
}

// GetFunc returns the method "CookieStore.get".
func (this CookieStore) GetFunc() (fn js.Func[func(name js.String) js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "CookieStore.get".
func (this CookieStore) Get(name js.String) (ret js.Promise[CookieListItem]) {
	bindings.CallCookieStoreGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "CookieStore.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGet(name js.String) (ret js.Promise[CookieListItem], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGet1 returns true if the method "CookieStore.get" exists.
func (this CookieStore) HasGet1() bool {
	return js.True == bindings.HasCookieStoreGet1(
		this.Ref(),
	)
}

// Get1Func returns the method "CookieStore.get".
func (this CookieStore) Get1Func() (fn js.Func[func(options CookieStoreGetOptions) js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGet1Func(
			this.Ref(),
		),
	)
}

// Get1 calls the method "CookieStore.get".
func (this CookieStore) Get1(options CookieStoreGetOptions) (ret js.Promise[CookieListItem]) {
	bindings.CallCookieStoreGet1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGet1 calls the method "CookieStore.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGet1(options CookieStoreGetOptions) (ret js.Promise[CookieListItem], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGet1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGet2 returns true if the method "CookieStore.get" exists.
func (this CookieStore) HasGet2() bool {
	return js.True == bindings.HasCookieStoreGet2(
		this.Ref(),
	)
}

// Get2Func returns the method "CookieStore.get".
func (this CookieStore) Get2Func() (fn js.Func[func() js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGet2Func(
			this.Ref(),
		),
	)
}

// Get2 calls the method "CookieStore.get".
func (this CookieStore) Get2() (ret js.Promise[CookieListItem]) {
	bindings.CallCookieStoreGet2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGet2 calls the method "CookieStore.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGet2() (ret js.Promise[CookieListItem], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGet2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAll returns true if the method "CookieStore.getAll" exists.
func (this CookieStore) HasGetAll() bool {
	return js.True == bindings.HasCookieStoreGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "CookieStore.getAll".
func (this CookieStore) GetAllFunc() (fn js.Func[func(name js.String) js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "CookieStore.getAll".
func (this CookieStore) GetAll(name js.String) (ret js.Promise[CookieList]) {
	bindings.CallCookieStoreGetAll(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetAll calls the method "CookieStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGetAll(name js.String) (ret js.Promise[CookieList], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetAll1 returns true if the method "CookieStore.getAll" exists.
func (this CookieStore) HasGetAll1() bool {
	return js.True == bindings.HasCookieStoreGetAll1(
		this.Ref(),
	)
}

// GetAll1Func returns the method "CookieStore.getAll".
func (this CookieStore) GetAll1Func() (fn js.Func[func(options CookieStoreGetOptions) js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "CookieStore.getAll".
func (this CookieStore) GetAll1(options CookieStoreGetOptions) (ret js.Promise[CookieList]) {
	bindings.CallCookieStoreGetAll1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetAll1 calls the method "CookieStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGetAll1(options CookieStoreGetOptions) (ret js.Promise[CookieList], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGetAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetAll2 returns true if the method "CookieStore.getAll" exists.
func (this CookieStore) HasGetAll2() bool {
	return js.True == bindings.HasCookieStoreGetAll2(
		this.Ref(),
	)
}

// GetAll2Func returns the method "CookieStore.getAll".
func (this CookieStore) GetAll2Func() (fn js.Func[func() js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAll2Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "CookieStore.getAll".
func (this CookieStore) GetAll2() (ret js.Promise[CookieList]) {
	bindings.CallCookieStoreGetAll2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAll2 calls the method "CookieStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryGetAll2() (ret js.Promise[CookieList], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreGetAll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSet returns true if the method "CookieStore.set" exists.
func (this CookieStore) HasSet() bool {
	return js.True == bindings.HasCookieStoreSet(
		this.Ref(),
	)
}

// SetFunc returns the method "CookieStore.set".
func (this CookieStore) SetFunc() (fn js.Func[func(name js.String, value js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CookieStore.set".
func (this CookieStore) Set(name js.String, value js.String) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreSet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "CookieStore.set"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TrySet(name js.String, value js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasSet1 returns true if the method "CookieStore.set" exists.
func (this CookieStore) HasSet1() bool {
	return js.True == bindings.HasCookieStoreSet1(
		this.Ref(),
	)
}

// Set1Func returns the method "CookieStore.set".
func (this CookieStore) Set1Func() (fn js.Func[func(options CookieInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreSet1Func(
			this.Ref(),
		),
	)
}

// Set1 calls the method "CookieStore.set".
func (this CookieStore) Set1(options CookieInit) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreSet1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySet1 calls the method "CookieStore.set"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TrySet1(options CookieInit) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreSet1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasDelete returns true if the method "CookieStore.delete" exists.
func (this CookieStore) HasDelete() bool {
	return js.True == bindings.HasCookieStoreDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "CookieStore.delete".
func (this CookieStore) DeleteFunc() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "CookieStore.delete".
func (this CookieStore) Delete(name js.String) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "CookieStore.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryDelete(name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasDelete1 returns true if the method "CookieStore.delete" exists.
func (this CookieStore) HasDelete1() bool {
	return js.True == bindings.HasCookieStoreDelete1(
		this.Ref(),
	)
}

// Delete1Func returns the method "CookieStore.delete".
func (this CookieStore) Delete1Func() (fn js.Func[func(options CookieStoreDeleteOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreDelete1Func(
			this.Ref(),
		),
	)
}

// Delete1 calls the method "CookieStore.delete".
func (this CookieStore) Delete1(options CookieStoreDeleteOptions) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreDelete1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryDelete1 calls the method "CookieStore.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CookieStore) TryDelete1(options CookieStoreDeleteOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreDelete1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type DocumentPictureInPictureOptions struct {
	// Width is "DocumentPictureInPictureOptions.width"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width uint64
	// Height is "DocumentPictureInPictureOptions.height"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height uint64

	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DocumentPictureInPictureOptions with all fields set.
func (p DocumentPictureInPictureOptions) FromRef(ref js.Ref) DocumentPictureInPictureOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DocumentPictureInPictureOptions in the application heap.
func (p DocumentPictureInPictureOptions) New() js.Ref {
	return bindings.DocumentPictureInPictureOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DocumentPictureInPictureOptions) UpdateFrom(ref js.Ref) {
	bindings.DocumentPictureInPictureOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DocumentPictureInPictureOptions) Update(ref js.Ref) {
	bindings.DocumentPictureInPictureOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DocumentPictureInPicture struct {
	EventTarget
}

func (this DocumentPictureInPicture) Once() DocumentPictureInPicture {
	this.Ref().Once()
	return this
}

func (this DocumentPictureInPicture) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this DocumentPictureInPicture) FromRef(ref js.Ref) DocumentPictureInPicture {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this DocumentPictureInPicture) Free() {
	this.Ref().Free()
}

// Window returns the value of property "DocumentPictureInPicture.window".
//
// It returns ok=false if there is no such property.
func (this DocumentPictureInPicture) Window() (ret Window, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestWindow returns true if the method "DocumentPictureInPicture.requestWindow" exists.
func (this DocumentPictureInPicture) HasRequestWindow() bool {
	return js.True == bindings.HasDocumentPictureInPictureRequestWindow(
		this.Ref(),
	)
}

// RequestWindowFunc returns the method "DocumentPictureInPicture.requestWindow".
func (this DocumentPictureInPicture) RequestWindowFunc() (fn js.Func[func(options DocumentPictureInPictureOptions) js.Promise[Window]]) {
	return fn.FromRef(
		bindings.DocumentPictureInPictureRequestWindowFunc(
			this.Ref(),
		),
	)
}

// RequestWindow calls the method "DocumentPictureInPicture.requestWindow".
func (this DocumentPictureInPicture) RequestWindow(options DocumentPictureInPictureOptions) (ret js.Promise[Window]) {
	bindings.CallDocumentPictureInPictureRequestWindow(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestWindow calls the method "DocumentPictureInPicture.requestWindow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentPictureInPicture) TryRequestWindow(options DocumentPictureInPictureOptions) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentPictureInPictureRequestWindow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestWindow1 returns true if the method "DocumentPictureInPicture.requestWindow" exists.
func (this DocumentPictureInPicture) HasRequestWindow1() bool {
	return js.True == bindings.HasDocumentPictureInPictureRequestWindow1(
		this.Ref(),
	)
}

// RequestWindow1Func returns the method "DocumentPictureInPicture.requestWindow".
func (this DocumentPictureInPicture) RequestWindow1Func() (fn js.Func[func() js.Promise[Window]]) {
	return fn.FromRef(
		bindings.DocumentPictureInPictureRequestWindow1Func(
			this.Ref(),
		),
	)
}

// RequestWindow1 calls the method "DocumentPictureInPicture.requestWindow".
func (this DocumentPictureInPicture) RequestWindow1() (ret js.Promise[Window]) {
	bindings.CallDocumentPictureInPictureRequestWindow1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestWindow1 calls the method "DocumentPictureInPicture.requestWindow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DocumentPictureInPicture) TryRequestWindow1() (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentPictureInPictureRequestWindow1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type External struct {
	ref js.Ref
}

func (this External) Once() External {
	this.Ref().Once()
	return this
}

func (this External) Ref() js.Ref {
	return this.ref
}

func (this External) FromRef(ref js.Ref) External {
	this.ref = ref
	return this
}

func (this External) Free() {
	this.Ref().Free()
}

// HasAddSearchProvider returns true if the method "External.AddSearchProvider" exists.
func (this External) HasAddSearchProvider() bool {
	return js.True == bindings.HasExternalAddSearchProvider(
		this.Ref(),
	)
}

// AddSearchProviderFunc returns the method "External.AddSearchProvider".
func (this External) AddSearchProviderFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ExternalAddSearchProviderFunc(
			this.Ref(),
		),
	)
}

// AddSearchProvider calls the method "External.AddSearchProvider".
func (this External) AddSearchProvider() (ret js.Void) {
	bindings.CallExternalAddSearchProvider(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAddSearchProvider calls the method "External.AddSearchProvider"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this External) TryAddSearchProvider() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryExternalAddSearchProvider(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsSearchProviderInstalled returns true if the method "External.IsSearchProviderInstalled" exists.
func (this External) HasIsSearchProviderInstalled() bool {
	return js.True == bindings.HasExternalIsSearchProviderInstalled(
		this.Ref(),
	)
}

// IsSearchProviderInstalledFunc returns the method "External.IsSearchProviderInstalled".
func (this External) IsSearchProviderInstalledFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ExternalIsSearchProviderInstalledFunc(
			this.Ref(),
		),
	)
}

// IsSearchProviderInstalled calls the method "External.IsSearchProviderInstalled".
func (this External) IsSearchProviderInstalled() (ret js.Void) {
	bindings.CallExternalIsSearchProviderInstalled(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsSearchProviderInstalled calls the method "External.IsSearchProviderInstalled"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this External) TryIsSearchProviderInstalled() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryExternalIsSearchProviderInstalled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SpeechSynthesisVoice struct {
	ref js.Ref
}

func (this SpeechSynthesisVoice) Once() SpeechSynthesisVoice {
	this.Ref().Once()
	return this
}

func (this SpeechSynthesisVoice) Ref() js.Ref {
	return this.ref
}

func (this SpeechSynthesisVoice) FromRef(ref js.Ref) SpeechSynthesisVoice {
	this.ref = ref
	return this
}

func (this SpeechSynthesisVoice) Free() {
	this.Ref().Free()
}

// VoiceURI returns the value of property "SpeechSynthesisVoice.voiceURI".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisVoice) VoiceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisVoiceVoiceURI(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "SpeechSynthesisVoice.name".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisVoice) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisVoiceName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Lang returns the value of property "SpeechSynthesisVoice.lang".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisVoice) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisVoiceLang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LocalService returns the value of property "SpeechSynthesisVoice.localService".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisVoice) LocalService() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisVoiceLocalService(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Default returns the value of property "SpeechSynthesisVoice.default".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisVoice) Default() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisVoiceDefault(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewSpeechSynthesisUtterance(text js.String) (ret SpeechSynthesisUtterance) {
	ret.ref = bindings.NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance(
		text.Ref())
	return
}

func NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance1() (ret SpeechSynthesisUtterance) {
	ret.ref = bindings.NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance1()
	return
}

type SpeechSynthesisUtterance struct {
	EventTarget
}

func (this SpeechSynthesisUtterance) Once() SpeechSynthesisUtterance {
	this.Ref().Once()
	return this
}

func (this SpeechSynthesisUtterance) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SpeechSynthesisUtterance) FromRef(ref js.Ref) SpeechSynthesisUtterance {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SpeechSynthesisUtterance) Free() {
	this.Ref().Free()
}

// Text returns the value of property "SpeechSynthesisUtterance.text".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtteranceText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "SpeechSynthesisUtterance.text" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetText(val js.String) bool {
	return js.True == bindings.SetSpeechSynthesisUtteranceText(
		this.Ref(),
		val.Ref(),
	)
}

// Lang returns the value of property "SpeechSynthesisUtterance.lang".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtteranceLang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLang sets the value of property "SpeechSynthesisUtterance.lang" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetLang(val js.String) bool {
	return js.True == bindings.SetSpeechSynthesisUtteranceLang(
		this.Ref(),
		val.Ref(),
	)
}

// Voice returns the value of property "SpeechSynthesisUtterance.voice".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Voice() (ret SpeechSynthesisVoice, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtteranceVoice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVoice sets the value of property "SpeechSynthesisUtterance.voice" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetVoice(val SpeechSynthesisVoice) bool {
	return js.True == bindings.SetSpeechSynthesisUtteranceVoice(
		this.Ref(),
		val.Ref(),
	)
}

// Volume returns the value of property "SpeechSynthesisUtterance.volume".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Volume() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtteranceVolume(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVolume sets the value of property "SpeechSynthesisUtterance.volume" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetVolume(val float32) bool {
	return js.True == bindings.SetSpeechSynthesisUtteranceVolume(
		this.Ref(),
		float32(val),
	)
}

// Rate returns the value of property "SpeechSynthesisUtterance.rate".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Rate() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtteranceRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRate sets the value of property "SpeechSynthesisUtterance.rate" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetRate(val float32) bool {
	return js.True == bindings.SetSpeechSynthesisUtteranceRate(
		this.Ref(),
		float32(val),
	)
}

// Pitch returns the value of property "SpeechSynthesisUtterance.pitch".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisUtterance) Pitch() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisUtterancePitch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPitch sets the value of property "SpeechSynthesisUtterance.pitch" to val.
//
// It returns false if the property cannot be set.
func (this SpeechSynthesisUtterance) SetPitch(val float32) bool {
	return js.True == bindings.SetSpeechSynthesisUtterancePitch(
		this.Ref(),
		float32(val),
	)
}

type SpeechSynthesis struct {
	EventTarget
}

func (this SpeechSynthesis) Once() SpeechSynthesis {
	this.Ref().Once()
	return this
}

func (this SpeechSynthesis) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SpeechSynthesis) FromRef(ref js.Ref) SpeechSynthesis {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SpeechSynthesis) Free() {
	this.Ref().Free()
}

// Pending returns the value of property "SpeechSynthesis.pending".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesis) Pending() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisPending(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Speaking returns the value of property "SpeechSynthesis.speaking".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesis) Speaking() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisSpeaking(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Paused returns the value of property "SpeechSynthesis.paused".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesis) Paused() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisPaused(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSpeak returns true if the method "SpeechSynthesis.speak" exists.
func (this SpeechSynthesis) HasSpeak() bool {
	return js.True == bindings.HasSpeechSynthesisSpeak(
		this.Ref(),
	)
}

// SpeakFunc returns the method "SpeechSynthesis.speak".
func (this SpeechSynthesis) SpeakFunc() (fn js.Func[func(utterance SpeechSynthesisUtterance)]) {
	return fn.FromRef(
		bindings.SpeechSynthesisSpeakFunc(
			this.Ref(),
		),
	)
}

// Speak calls the method "SpeechSynthesis.speak".
func (this SpeechSynthesis) Speak(utterance SpeechSynthesisUtterance) (ret js.Void) {
	bindings.CallSpeechSynthesisSpeak(
		this.Ref(), js.Pointer(&ret),
		utterance.Ref(),
	)

	return
}

// TrySpeak calls the method "SpeechSynthesis.speak"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SpeechSynthesis) TrySpeak(utterance SpeechSynthesisUtterance) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechSynthesisSpeak(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		utterance.Ref(),
	)

	return
}

// HasCancel returns true if the method "SpeechSynthesis.cancel" exists.
func (this SpeechSynthesis) HasCancel() bool {
	return js.True == bindings.HasSpeechSynthesisCancel(
		this.Ref(),
	)
}

// CancelFunc returns the method "SpeechSynthesis.cancel".
func (this SpeechSynthesis) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "SpeechSynthesis.cancel".
func (this SpeechSynthesis) Cancel() (ret js.Void) {
	bindings.CallSpeechSynthesisCancel(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "SpeechSynthesis.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SpeechSynthesis) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechSynthesisCancel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPause returns true if the method "SpeechSynthesis.pause" exists.
func (this SpeechSynthesis) HasPause() bool {
	return js.True == bindings.HasSpeechSynthesisPause(
		this.Ref(),
	)
}

// PauseFunc returns the method "SpeechSynthesis.pause".
func (this SpeechSynthesis) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisPauseFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "SpeechSynthesis.pause".
func (this SpeechSynthesis) Pause() (ret js.Void) {
	bindings.CallSpeechSynthesisPause(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "SpeechSynthesis.pause"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SpeechSynthesis) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechSynthesisPause(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResume returns true if the method "SpeechSynthesis.resume" exists.
func (this SpeechSynthesis) HasResume() bool {
	return js.True == bindings.HasSpeechSynthesisResume(
		this.Ref(),
	)
}

// ResumeFunc returns the method "SpeechSynthesis.resume".
func (this SpeechSynthesis) ResumeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisResumeFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "SpeechSynthesis.resume".
func (this SpeechSynthesis) Resume() (ret js.Void) {
	bindings.CallSpeechSynthesisResume(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResume calls the method "SpeechSynthesis.resume"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SpeechSynthesis) TryResume() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechSynthesisResume(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetVoices returns true if the method "SpeechSynthesis.getVoices" exists.
func (this SpeechSynthesis) HasGetVoices() bool {
	return js.True == bindings.HasSpeechSynthesisGetVoices(
		this.Ref(),
	)
}

// GetVoicesFunc returns the method "SpeechSynthesis.getVoices".
func (this SpeechSynthesis) GetVoicesFunc() (fn js.Func[func() js.Array[SpeechSynthesisVoice]]) {
	return fn.FromRef(
		bindings.SpeechSynthesisGetVoicesFunc(
			this.Ref(),
		),
	)
}

// GetVoices calls the method "SpeechSynthesis.getVoices".
func (this SpeechSynthesis) GetVoices() (ret js.Array[SpeechSynthesisVoice]) {
	bindings.CallSpeechSynthesisGetVoices(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetVoices calls the method "SpeechSynthesis.getVoices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SpeechSynthesis) TryGetVoices() (ret js.Array[SpeechSynthesisVoice], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechSynthesisGetVoices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SchedulerPostTaskCallbackFunc func(this js.Ref) js.Ref

func (fn SchedulerPostTaskCallbackFunc) Register() js.Func[func() js.Any] {
	return js.RegisterCallback[func() js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SchedulerPostTaskCallbackFunc) DispatchCallback(
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

type SchedulerPostTaskCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SchedulerPostTaskCallback[T]) Register() js.Func[func() js.Any] {
	return js.RegisterCallback[func() js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SchedulerPostTaskCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TaskPriority uint32

const (
	_ TaskPriority = iota

	TaskPriority_USER_BLOCKING
	TaskPriority_USER_VISIBLE
	TaskPriority_BACKGROUND
)

func (TaskPriority) FromRef(str js.Ref) TaskPriority {
	return TaskPriority(bindings.ConstOfTaskPriority(str))
}

func (x TaskPriority) String() (string, bool) {
	switch x {
	case TaskPriority_USER_BLOCKING:
		return "user-blocking", true
	case TaskPriority_USER_VISIBLE:
		return "user-visible", true
	case TaskPriority_BACKGROUND:
		return "background", true
	default:
		return "", false
	}
}

type SchedulerPostTaskOptions struct {
	// Signal is "SchedulerPostTaskOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// Priority is "SchedulerPostTaskOptions.priority"
	//
	// Optional
	Priority TaskPriority
	// Delay is "SchedulerPostTaskOptions.delay"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Delay MUST be set to true to make this field effective.
	Delay uint64

	FFI_USE_Delay bool // for Delay.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SchedulerPostTaskOptions with all fields set.
func (p SchedulerPostTaskOptions) FromRef(ref js.Ref) SchedulerPostTaskOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SchedulerPostTaskOptions in the application heap.
func (p SchedulerPostTaskOptions) New() js.Ref {
	return bindings.SchedulerPostTaskOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SchedulerPostTaskOptions) UpdateFrom(ref js.Ref) {
	bindings.SchedulerPostTaskOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SchedulerPostTaskOptions) Update(ref js.Ref) {
	bindings.SchedulerPostTaskOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Scheduler struct {
	ref js.Ref
}

func (this Scheduler) Once() Scheduler {
	this.Ref().Once()
	return this
}

func (this Scheduler) Ref() js.Ref {
	return this.ref
}

func (this Scheduler) FromRef(ref js.Ref) Scheduler {
	this.ref = ref
	return this
}

func (this Scheduler) Free() {
	this.Ref().Free()
}

// HasPostTask returns true if the method "Scheduler.postTask" exists.
func (this Scheduler) HasPostTask() bool {
	return js.True == bindings.HasSchedulerPostTask(
		this.Ref(),
	)
}

// PostTaskFunc returns the method "Scheduler.postTask".
func (this Scheduler) PostTaskFunc() (fn js.Func[func(callback js.Func[func() js.Any], options SchedulerPostTaskOptions) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SchedulerPostTaskFunc(
			this.Ref(),
		),
	)
}

// PostTask calls the method "Scheduler.postTask".
func (this Scheduler) PostTask(callback js.Func[func() js.Any], options SchedulerPostTaskOptions) (ret js.Promise[js.Any]) {
	bindings.CallSchedulerPostTask(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostTask calls the method "Scheduler.postTask"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Scheduler) TryPostTask(callback js.Func[func() js.Any], options SchedulerPostTaskOptions) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulerPostTask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostTask1 returns true if the method "Scheduler.postTask" exists.
func (this Scheduler) HasPostTask1() bool {
	return js.True == bindings.HasSchedulerPostTask1(
		this.Ref(),
	)
}

// PostTask1Func returns the method "Scheduler.postTask".
func (this Scheduler) PostTask1Func() (fn js.Func[func(callback js.Func[func() js.Any]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SchedulerPostTask1Func(
			this.Ref(),
		),
	)
}

// PostTask1 calls the method "Scheduler.postTask".
func (this Scheduler) PostTask1(callback js.Func[func() js.Any]) (ret js.Promise[js.Any]) {
	bindings.CallSchedulerPostTask1(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryPostTask1 calls the method "Scheduler.postTask"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Scheduler) TryPostTask1(callback js.Func[func() js.Any]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySchedulerPostTask1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type TrustedHTML struct {
	ref js.Ref
}

func (this TrustedHTML) Once() TrustedHTML {
	this.Ref().Once()
	return this
}

func (this TrustedHTML) Ref() js.Ref {
	return this.ref
}

func (this TrustedHTML) FromRef(ref js.Ref) TrustedHTML {
	this.ref = ref
	return this
}

func (this TrustedHTML) Free() {
	this.Ref().Free()
}

// HasToString returns true if the method "TrustedHTML.toString" exists.
func (this TrustedHTML) HasToString() bool {
	return js.True == bindings.HasTrustedHTMLToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "TrustedHTML.toString".
func (this TrustedHTML) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedHTMLToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "TrustedHTML.toString".
func (this TrustedHTML) ToString() (ret js.String) {
	bindings.CallTrustedHTMLToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "TrustedHTML.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedHTML) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedHTMLToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "TrustedHTML.toJSON" exists.
func (this TrustedHTML) HasToJSON() bool {
	return js.True == bindings.HasTrustedHTMLToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "TrustedHTML.toJSON".
func (this TrustedHTML) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedHTMLToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedHTML.toJSON".
func (this TrustedHTML) ToJSON() (ret js.String) {
	bindings.CallTrustedHTMLToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "TrustedHTML.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedHTML) TryToJSON() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedHTMLToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromLiteral returns true if the staticmethod "TrustedHTML.fromLiteral" exists.
func (this TrustedHTML) HasFromLiteral() bool {
	return js.True == bindings.HasTrustedHTMLFromLiteral(
		this.Ref(),
	)
}

// FromLiteralFunc returns the staticmethod "TrustedHTML.fromLiteral".
func (this TrustedHTML) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedHTML]) {
	return fn.FromRef(
		bindings.TrustedHTMLFromLiteralFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedHTML.fromLiteral".
func (this TrustedHTML) FromLiteral(templateStringsArray js.Object) (ret TrustedHTML) {
	bindings.CallTrustedHTMLFromLiteral(
		this.Ref(), js.Pointer(&ret),
		templateStringsArray.Ref(),
	)

	return
}

// TryFromLiteral calls the staticmethod "TrustedHTML.fromLiteral"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedHTML) TryFromLiteral(templateStringsArray js.Object) (ret TrustedHTML, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedHTMLFromLiteral(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		templateStringsArray.Ref(),
	)

	return
}

type TrustedScript struct {
	ref js.Ref
}

func (this TrustedScript) Once() TrustedScript {
	this.Ref().Once()
	return this
}

func (this TrustedScript) Ref() js.Ref {
	return this.ref
}

func (this TrustedScript) FromRef(ref js.Ref) TrustedScript {
	this.ref = ref
	return this
}

func (this TrustedScript) Free() {
	this.Ref().Free()
}

// HasToString returns true if the method "TrustedScript.toString" exists.
func (this TrustedScript) HasToString() bool {
	return js.True == bindings.HasTrustedScriptToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "TrustedScript.toString".
func (this TrustedScript) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "TrustedScript.toString".
func (this TrustedScript) ToString() (ret js.String) {
	bindings.CallTrustedScriptToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "TrustedScript.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScript) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "TrustedScript.toJSON" exists.
func (this TrustedScript) HasToJSON() bool {
	return js.True == bindings.HasTrustedScriptToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "TrustedScript.toJSON".
func (this TrustedScript) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedScript.toJSON".
func (this TrustedScript) ToJSON() (ret js.String) {
	bindings.CallTrustedScriptToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "TrustedScript.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScript) TryToJSON() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromLiteral returns true if the staticmethod "TrustedScript.fromLiteral" exists.
func (this TrustedScript) HasFromLiteral() bool {
	return js.True == bindings.HasTrustedScriptFromLiteral(
		this.Ref(),
	)
}

// FromLiteralFunc returns the staticmethod "TrustedScript.fromLiteral".
func (this TrustedScript) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedScript]) {
	return fn.FromRef(
		bindings.TrustedScriptFromLiteralFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedScript.fromLiteral".
func (this TrustedScript) FromLiteral(templateStringsArray js.Object) (ret TrustedScript) {
	bindings.CallTrustedScriptFromLiteral(
		this.Ref(), js.Pointer(&ret),
		templateStringsArray.Ref(),
	)

	return
}

// TryFromLiteral calls the staticmethod "TrustedScript.fromLiteral"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScript) TryFromLiteral(templateStringsArray js.Object) (ret TrustedScript, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptFromLiteral(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		templateStringsArray.Ref(),
	)

	return
}

type TrustedScriptURL struct {
	ref js.Ref
}

func (this TrustedScriptURL) Once() TrustedScriptURL {
	this.Ref().Once()
	return this
}

func (this TrustedScriptURL) Ref() js.Ref {
	return this.ref
}

func (this TrustedScriptURL) FromRef(ref js.Ref) TrustedScriptURL {
	this.ref = ref
	return this
}

func (this TrustedScriptURL) Free() {
	this.Ref().Free()
}

// HasToString returns true if the method "TrustedScriptURL.toString" exists.
func (this TrustedScriptURL) HasToString() bool {
	return js.True == bindings.HasTrustedScriptURLToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "TrustedScriptURL.toString".
func (this TrustedScriptURL) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptURLToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "TrustedScriptURL.toString".
func (this TrustedScriptURL) ToString() (ret js.String) {
	bindings.CallTrustedScriptURLToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "TrustedScriptURL.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScriptURL) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptURLToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "TrustedScriptURL.toJSON" exists.
func (this TrustedScriptURL) HasToJSON() bool {
	return js.True == bindings.HasTrustedScriptURLToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "TrustedScriptURL.toJSON".
func (this TrustedScriptURL) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptURLToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedScriptURL.toJSON".
func (this TrustedScriptURL) ToJSON() (ret js.String) {
	bindings.CallTrustedScriptURLToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "TrustedScriptURL.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScriptURL) TryToJSON() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptURLToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromLiteral returns true if the staticmethod "TrustedScriptURL.fromLiteral" exists.
func (this TrustedScriptURL) HasFromLiteral() bool {
	return js.True == bindings.HasTrustedScriptURLFromLiteral(
		this.Ref(),
	)
}

// FromLiteralFunc returns the staticmethod "TrustedScriptURL.fromLiteral".
func (this TrustedScriptURL) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedScriptURL]) {
	return fn.FromRef(
		bindings.TrustedScriptURLFromLiteralFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedScriptURL.fromLiteral".
func (this TrustedScriptURL) FromLiteral(templateStringsArray js.Object) (ret TrustedScriptURL) {
	bindings.CallTrustedScriptURLFromLiteral(
		this.Ref(), js.Pointer(&ret),
		templateStringsArray.Ref(),
	)

	return
}

// TryFromLiteral calls the staticmethod "TrustedScriptURL.fromLiteral"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedScriptURL) TryFromLiteral(templateStringsArray js.Object) (ret TrustedScriptURL, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedScriptURLFromLiteral(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		templateStringsArray.Ref(),
	)

	return
}

type TrustedTypePolicy struct {
	ref js.Ref
}

func (this TrustedTypePolicy) Once() TrustedTypePolicy {
	this.Ref().Once()
	return this
}

func (this TrustedTypePolicy) Ref() js.Ref {
	return this.ref
}

func (this TrustedTypePolicy) FromRef(ref js.Ref) TrustedTypePolicy {
	this.ref = ref
	return this
}

func (this TrustedTypePolicy) Free() {
	this.Ref().Free()
}

// Name returns the value of property "TrustedTypePolicy.name".
//
// It returns ok=false if there is no such property.
func (this TrustedTypePolicy) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTrustedTypePolicyName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCreateHTML returns true if the method "TrustedTypePolicy.createHTML" exists.
func (this TrustedTypePolicy) HasCreateHTML() bool {
	return js.True == bindings.HasTrustedTypePolicyCreateHTML(
		this.Ref(),
	)
}

// CreateHTMLFunc returns the method "TrustedTypePolicy.createHTML".
func (this TrustedTypePolicy) CreateHTMLFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedHTML]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateHTMLFunc(
			this.Ref(),
		),
	)
}

// CreateHTML calls the method "TrustedTypePolicy.createHTML".
func (this TrustedTypePolicy) CreateHTML(input js.String, arguments ...js.Any) (ret TrustedHTML) {
	bindings.CallTrustedTypePolicyCreateHTML(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TryCreateHTML calls the method "TrustedTypePolicy.createHTML"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicy) TryCreateHTML(input js.String, arguments ...js.Any) (ret TrustedHTML, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyCreateHTML(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasCreateScript returns true if the method "TrustedTypePolicy.createScript" exists.
func (this TrustedTypePolicy) HasCreateScript() bool {
	return js.True == bindings.HasTrustedTypePolicyCreateScript(
		this.Ref(),
	)
}

// CreateScriptFunc returns the method "TrustedTypePolicy.createScript".
func (this TrustedTypePolicy) CreateScriptFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedScript]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateScriptFunc(
			this.Ref(),
		),
	)
}

// CreateScript calls the method "TrustedTypePolicy.createScript".
func (this TrustedTypePolicy) CreateScript(input js.String, arguments ...js.Any) (ret TrustedScript) {
	bindings.CallTrustedTypePolicyCreateScript(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TryCreateScript calls the method "TrustedTypePolicy.createScript"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicy) TryCreateScript(input js.String, arguments ...js.Any) (ret TrustedScript, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyCreateScript(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasCreateScriptURL returns true if the method "TrustedTypePolicy.createScriptURL" exists.
func (this TrustedTypePolicy) HasCreateScriptURL() bool {
	return js.True == bindings.HasTrustedTypePolicyCreateScriptURL(
		this.Ref(),
	)
}

// CreateScriptURLFunc returns the method "TrustedTypePolicy.createScriptURL".
func (this TrustedTypePolicy) CreateScriptURLFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedScriptURL]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateScriptURLFunc(
			this.Ref(),
		),
	)
}

// CreateScriptURL calls the method "TrustedTypePolicy.createScriptURL".
func (this TrustedTypePolicy) CreateScriptURL(input js.String, arguments ...js.Any) (ret TrustedScriptURL) {
	bindings.CallTrustedTypePolicyCreateScriptURL(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TryCreateScriptURL calls the method "TrustedTypePolicy.createScriptURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicy) TryCreateScriptURL(input js.String, arguments ...js.Any) (ret TrustedScriptURL, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyCreateScriptURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

type CreateHTMLCallbackFunc func(this js.Ref, input js.String, arguments ...js.Any) js.Ref

func (fn CreateHTMLCallbackFunc) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateHTMLCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateHTMLCallback[T any] struct {
	Fn  func(arg T, this js.Ref, input js.String, arguments ...js.Any) js.Ref
	Arg T
}

func (cb *CreateHTMLCallback[T]) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateHTMLCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateScriptCallbackFunc func(this js.Ref, input js.String, arguments ...js.Any) js.Ref

func (fn CreateScriptCallbackFunc) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateScriptCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateScriptCallback[T any] struct {
	Fn  func(arg T, this js.Ref, input js.String, arguments ...js.Any) js.Ref
	Arg T
}

func (cb *CreateScriptCallback[T]) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateScriptCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateScriptURLCallbackFunc func(this js.Ref, input js.String, arguments ...js.Any) js.Ref

func (fn CreateScriptURLCallbackFunc) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateScriptURLCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateScriptURLCallback[T any] struct {
	Fn  func(arg T, this js.Ref, input js.String, arguments ...js.Any) js.Ref
	Arg T
}

func (cb *CreateScriptURLCallback[T]) Register() js.Func[func(input js.String, arguments ...js.Any) js.String] {
	return js.RegisterCallback[func(input js.String, arguments ...js.Any) js.String](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateScriptURLCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TrustedTypePolicyOptions struct {
	// CreateHTML is "TrustedTypePolicyOptions.createHTML"
	//
	// Optional
	CreateHTML js.Func[func(input js.String, arguments ...js.Any) js.String]
	// CreateScript is "TrustedTypePolicyOptions.createScript"
	//
	// Optional
	CreateScript js.Func[func(input js.String, arguments ...js.Any) js.String]
	// CreateScriptURL is "TrustedTypePolicyOptions.createScriptURL"
	//
	// Optional
	CreateScriptURL js.Func[func(input js.String, arguments ...js.Any) js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TrustedTypePolicyOptions with all fields set.
func (p TrustedTypePolicyOptions) FromRef(ref js.Ref) TrustedTypePolicyOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TrustedTypePolicyOptions in the application heap.
func (p TrustedTypePolicyOptions) New() js.Ref {
	return bindings.TrustedTypePolicyOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TrustedTypePolicyOptions) UpdateFrom(ref js.Ref) {
	bindings.TrustedTypePolicyOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TrustedTypePolicyOptions) Update(ref js.Ref) {
	bindings.TrustedTypePolicyOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TrustedTypePolicyFactory struct {
	ref js.Ref
}

func (this TrustedTypePolicyFactory) Once() TrustedTypePolicyFactory {
	this.Ref().Once()
	return this
}

func (this TrustedTypePolicyFactory) Ref() js.Ref {
	return this.ref
}

func (this TrustedTypePolicyFactory) FromRef(ref js.Ref) TrustedTypePolicyFactory {
	this.ref = ref
	return this
}

func (this TrustedTypePolicyFactory) Free() {
	this.Ref().Free()
}

// EmptyHTML returns the value of property "TrustedTypePolicyFactory.emptyHTML".
//
// It returns ok=false if there is no such property.
func (this TrustedTypePolicyFactory) EmptyHTML() (ret TrustedHTML, ok bool) {
	ok = js.True == bindings.GetTrustedTypePolicyFactoryEmptyHTML(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EmptyScript returns the value of property "TrustedTypePolicyFactory.emptyScript".
//
// It returns ok=false if there is no such property.
func (this TrustedTypePolicyFactory) EmptyScript() (ret TrustedScript, ok bool) {
	ok = js.True == bindings.GetTrustedTypePolicyFactoryEmptyScript(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultPolicy returns the value of property "TrustedTypePolicyFactory.defaultPolicy".
//
// It returns ok=false if there is no such property.
func (this TrustedTypePolicyFactory) DefaultPolicy() (ret TrustedTypePolicy, ok bool) {
	ok = js.True == bindings.GetTrustedTypePolicyFactoryDefaultPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCreatePolicy returns true if the method "TrustedTypePolicyFactory.createPolicy" exists.
func (this TrustedTypePolicyFactory) HasCreatePolicy() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryCreatePolicy(
		this.Ref(),
	)
}

// CreatePolicyFunc returns the method "TrustedTypePolicyFactory.createPolicy".
func (this TrustedTypePolicyFactory) CreatePolicyFunc() (fn js.Func[func(policyName js.String, policyOptions TrustedTypePolicyOptions) TrustedTypePolicy]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryCreatePolicyFunc(
			this.Ref(),
		),
	)
}

// CreatePolicy calls the method "TrustedTypePolicyFactory.createPolicy".
func (this TrustedTypePolicyFactory) CreatePolicy(policyName js.String, policyOptions TrustedTypePolicyOptions) (ret TrustedTypePolicy) {
	bindings.CallTrustedTypePolicyFactoryCreatePolicy(
		this.Ref(), js.Pointer(&ret),
		policyName.Ref(),
		js.Pointer(&policyOptions),
	)

	return
}

// TryCreatePolicy calls the method "TrustedTypePolicyFactory.createPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryCreatePolicy(policyName js.String, policyOptions TrustedTypePolicyOptions) (ret TrustedTypePolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryCreatePolicy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		policyName.Ref(),
		js.Pointer(&policyOptions),
	)

	return
}

// HasCreatePolicy1 returns true if the method "TrustedTypePolicyFactory.createPolicy" exists.
func (this TrustedTypePolicyFactory) HasCreatePolicy1() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryCreatePolicy1(
		this.Ref(),
	)
}

// CreatePolicy1Func returns the method "TrustedTypePolicyFactory.createPolicy".
func (this TrustedTypePolicyFactory) CreatePolicy1Func() (fn js.Func[func(policyName js.String) TrustedTypePolicy]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryCreatePolicy1Func(
			this.Ref(),
		),
	)
}

// CreatePolicy1 calls the method "TrustedTypePolicyFactory.createPolicy".
func (this TrustedTypePolicyFactory) CreatePolicy1(policyName js.String) (ret TrustedTypePolicy) {
	bindings.CallTrustedTypePolicyFactoryCreatePolicy1(
		this.Ref(), js.Pointer(&ret),
		policyName.Ref(),
	)

	return
}

// TryCreatePolicy1 calls the method "TrustedTypePolicyFactory.createPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryCreatePolicy1(policyName js.String) (ret TrustedTypePolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryCreatePolicy1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		policyName.Ref(),
	)

	return
}

// HasIsHTML returns true if the method "TrustedTypePolicyFactory.isHTML" exists.
func (this TrustedTypePolicyFactory) HasIsHTML() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryIsHTML(
		this.Ref(),
	)
}

// IsHTMLFunc returns the method "TrustedTypePolicyFactory.isHTML".
func (this TrustedTypePolicyFactory) IsHTMLFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsHTMLFunc(
			this.Ref(),
		),
	)
}

// IsHTML calls the method "TrustedTypePolicyFactory.isHTML".
func (this TrustedTypePolicyFactory) IsHTML(value js.Any) (ret bool) {
	bindings.CallTrustedTypePolicyFactoryIsHTML(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryIsHTML calls the method "TrustedTypePolicyFactory.isHTML"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryIsHTML(value js.Any) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryIsHTML(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasIsScript returns true if the method "TrustedTypePolicyFactory.isScript" exists.
func (this TrustedTypePolicyFactory) HasIsScript() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryIsScript(
		this.Ref(),
	)
}

// IsScriptFunc returns the method "TrustedTypePolicyFactory.isScript".
func (this TrustedTypePolicyFactory) IsScriptFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsScriptFunc(
			this.Ref(),
		),
	)
}

// IsScript calls the method "TrustedTypePolicyFactory.isScript".
func (this TrustedTypePolicyFactory) IsScript(value js.Any) (ret bool) {
	bindings.CallTrustedTypePolicyFactoryIsScript(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryIsScript calls the method "TrustedTypePolicyFactory.isScript"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryIsScript(value js.Any) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryIsScript(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasIsScriptURL returns true if the method "TrustedTypePolicyFactory.isScriptURL" exists.
func (this TrustedTypePolicyFactory) HasIsScriptURL() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryIsScriptURL(
		this.Ref(),
	)
}

// IsScriptURLFunc returns the method "TrustedTypePolicyFactory.isScriptURL".
func (this TrustedTypePolicyFactory) IsScriptURLFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsScriptURLFunc(
			this.Ref(),
		),
	)
}

// IsScriptURL calls the method "TrustedTypePolicyFactory.isScriptURL".
func (this TrustedTypePolicyFactory) IsScriptURL(value js.Any) (ret bool) {
	bindings.CallTrustedTypePolicyFactoryIsScriptURL(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryIsScriptURL calls the method "TrustedTypePolicyFactory.isScriptURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryIsScriptURL(value js.Any) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryIsScriptURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasGetAttributeType returns true if the method "TrustedTypePolicyFactory.getAttributeType" exists.
func (this TrustedTypePolicyFactory) HasGetAttributeType() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryGetAttributeType(
		this.Ref(),
	)
}

// GetAttributeTypeFunc returns the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeTypeFunc() (fn js.Func[func(tagName js.String, attribute js.String, elementNs js.String, attrNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeTypeFunc(
			this.Ref(),
		),
	)
}

// GetAttributeType calls the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeType(tagName js.String, attribute js.String, elementNs js.String, attrNs js.String) (ret js.String) {
	bindings.CallTrustedTypePolicyFactoryGetAttributeType(
		this.Ref(), js.Pointer(&ret),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
		attrNs.Ref(),
	)

	return
}

// TryGetAttributeType calls the method "TrustedTypePolicyFactory.getAttributeType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryGetAttributeType(tagName js.String, attribute js.String, elementNs js.String, attrNs js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryGetAttributeType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
		attrNs.Ref(),
	)

	return
}

// HasGetAttributeType1 returns true if the method "TrustedTypePolicyFactory.getAttributeType" exists.
func (this TrustedTypePolicyFactory) HasGetAttributeType1() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryGetAttributeType1(
		this.Ref(),
	)
}

// GetAttributeType1Func returns the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeType1Func() (fn js.Func[func(tagName js.String, attribute js.String, elementNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeType1Func(
			this.Ref(),
		),
	)
}

// GetAttributeType1 calls the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeType1(tagName js.String, attribute js.String, elementNs js.String) (ret js.String) {
	bindings.CallTrustedTypePolicyFactoryGetAttributeType1(
		this.Ref(), js.Pointer(&ret),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
	)

	return
}

// TryGetAttributeType1 calls the method "TrustedTypePolicyFactory.getAttributeType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryGetAttributeType1(tagName js.String, attribute js.String, elementNs js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryGetAttributeType1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
	)

	return
}

// HasGetAttributeType2 returns true if the method "TrustedTypePolicyFactory.getAttributeType" exists.
func (this TrustedTypePolicyFactory) HasGetAttributeType2() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryGetAttributeType2(
		this.Ref(),
	)
}

// GetAttributeType2Func returns the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeType2Func() (fn js.Func[func(tagName js.String, attribute js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeType2Func(
			this.Ref(),
		),
	)
}

// GetAttributeType2 calls the method "TrustedTypePolicyFactory.getAttributeType".
func (this TrustedTypePolicyFactory) GetAttributeType2(tagName js.String, attribute js.String) (ret js.String) {
	bindings.CallTrustedTypePolicyFactoryGetAttributeType2(
		this.Ref(), js.Pointer(&ret),
		tagName.Ref(),
		attribute.Ref(),
	)

	return
}

// TryGetAttributeType2 calls the method "TrustedTypePolicyFactory.getAttributeType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryGetAttributeType2(tagName js.String, attribute js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryGetAttributeType2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tagName.Ref(),
		attribute.Ref(),
	)

	return
}

// HasGetPropertyType returns true if the method "TrustedTypePolicyFactory.getPropertyType" exists.
func (this TrustedTypePolicyFactory) HasGetPropertyType() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryGetPropertyType(
		this.Ref(),
	)
}

// GetPropertyTypeFunc returns the method "TrustedTypePolicyFactory.getPropertyType".
func (this TrustedTypePolicyFactory) GetPropertyTypeFunc() (fn js.Func[func(tagName js.String, property js.String, elementNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetPropertyTypeFunc(
			this.Ref(),
		),
	)
}

// GetPropertyType calls the method "TrustedTypePolicyFactory.getPropertyType".
func (this TrustedTypePolicyFactory) GetPropertyType(tagName js.String, property js.String, elementNs js.String) (ret js.String) {
	bindings.CallTrustedTypePolicyFactoryGetPropertyType(
		this.Ref(), js.Pointer(&ret),
		tagName.Ref(),
		property.Ref(),
		elementNs.Ref(),
	)

	return
}

// TryGetPropertyType calls the method "TrustedTypePolicyFactory.getPropertyType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryGetPropertyType(tagName js.String, property js.String, elementNs js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryGetPropertyType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tagName.Ref(),
		property.Ref(),
		elementNs.Ref(),
	)

	return
}

// HasGetPropertyType1 returns true if the method "TrustedTypePolicyFactory.getPropertyType" exists.
func (this TrustedTypePolicyFactory) HasGetPropertyType1() bool {
	return js.True == bindings.HasTrustedTypePolicyFactoryGetPropertyType1(
		this.Ref(),
	)
}

// GetPropertyType1Func returns the method "TrustedTypePolicyFactory.getPropertyType".
func (this TrustedTypePolicyFactory) GetPropertyType1Func() (fn js.Func[func(tagName js.String, property js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetPropertyType1Func(
			this.Ref(),
		),
	)
}

// GetPropertyType1 calls the method "TrustedTypePolicyFactory.getPropertyType".
func (this TrustedTypePolicyFactory) GetPropertyType1(tagName js.String, property js.String) (ret js.String) {
	bindings.CallTrustedTypePolicyFactoryGetPropertyType1(
		this.Ref(), js.Pointer(&ret),
		tagName.Ref(),
		property.Ref(),
	)

	return
}

// TryGetPropertyType1 calls the method "TrustedTypePolicyFactory.getPropertyType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this TrustedTypePolicyFactory) TryGetPropertyType1(tagName js.String, property js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTrustedTypePolicyFactoryGetPropertyType1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tagName.Ref(),
		property.Ref(),
	)

	return
}

type KeyType uint32

const (
	_ KeyType = iota

	KeyType_PUBLIC
	KeyType_PRIVATE
	KeyType_SECRET
)

func (KeyType) FromRef(str js.Ref) KeyType {
	return KeyType(bindings.ConstOfKeyType(str))
}

func (x KeyType) String() (string, bool) {
	switch x {
	case KeyType_PUBLIC:
		return "public", true
	case KeyType_PRIVATE:
		return "private", true
	case KeyType_SECRET:
		return "secret", true
	default:
		return "", false
	}
}

type CryptoKey struct {
	ref js.Ref
}

func (this CryptoKey) Once() CryptoKey {
	this.Ref().Once()
	return this
}

func (this CryptoKey) Ref() js.Ref {
	return this.ref
}

func (this CryptoKey) FromRef(ref js.Ref) CryptoKey {
	this.ref = ref
	return this
}

func (this CryptoKey) Free() {
	this.Ref().Free()
}

// Type returns the value of property "CryptoKey.type".
//
// It returns ok=false if there is no such property.
func (this CryptoKey) Type() (ret KeyType, ok bool) {
	ok = js.True == bindings.GetCryptoKeyType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Extractable returns the value of property "CryptoKey.extractable".
//
// It returns ok=false if there is no such property.
func (this CryptoKey) Extractable() (ret bool, ok bool) {
	ok = js.True == bindings.GetCryptoKeyExtractable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Algorithm returns the value of property "CryptoKey.algorithm".
//
// It returns ok=false if there is no such property.
func (this CryptoKey) Algorithm() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetCryptoKeyAlgorithm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Usages returns the value of property "CryptoKey.usages".
//
// It returns ok=false if there is no such property.
func (this CryptoKey) Usages() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetCryptoKeyUsages(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type KeyUsage uint32

const (
	_ KeyUsage = iota

	KeyUsage_ENCRYPT
	KeyUsage_DECRYPT
	KeyUsage_SIGN
	KeyUsage_VERIFY
	KeyUsage_DERIVE_KEY
	KeyUsage_DERIVE_BITS
	KeyUsage_WRAP_KEY
	KeyUsage_UNWRAP_KEY
)

func (KeyUsage) FromRef(str js.Ref) KeyUsage {
	return KeyUsage(bindings.ConstOfKeyUsage(str))
}

func (x KeyUsage) String() (string, bool) {
	switch x {
	case KeyUsage_ENCRYPT:
		return "encrypt", true
	case KeyUsage_DECRYPT:
		return "decrypt", true
	case KeyUsage_SIGN:
		return "sign", true
	case KeyUsage_VERIFY:
		return "verify", true
	case KeyUsage_DERIVE_KEY:
		return "deriveKey", true
	case KeyUsage_DERIVE_BITS:
		return "deriveBits", true
	case KeyUsage_WRAP_KEY:
		return "wrapKey", true
	case KeyUsage_UNWRAP_KEY:
		return "unwrapKey", true
	default:
		return "", false
	}
}

type KeyFormat uint32

const (
	_ KeyFormat = iota

	KeyFormat_RAW
	KeyFormat_SPKI
	KeyFormat_PKCS8
	KeyFormat_JWK
)

func (KeyFormat) FromRef(str js.Ref) KeyFormat {
	return KeyFormat(bindings.ConstOfKeyFormat(str))
}

func (x KeyFormat) String() (string, bool) {
	switch x {
	case KeyFormat_RAW:
		return "raw", true
	case KeyFormat_SPKI:
		return "spki", true
	case KeyFormat_PKCS8:
		return "pkcs8", true
	case KeyFormat_JWK:
		return "jwk", true
	default:
		return "", false
	}
}

type RsaOtherPrimesInfo struct {
	// R is "RsaOtherPrimesInfo.r"
	//
	// Optional
	R js.String
	// D is "RsaOtherPrimesInfo.d"
	//
	// Optional
	D js.String
	// T is "RsaOtherPrimesInfo.t"
	//
	// Optional
	T js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RsaOtherPrimesInfo with all fields set.
func (p RsaOtherPrimesInfo) FromRef(ref js.Ref) RsaOtherPrimesInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RsaOtherPrimesInfo in the application heap.
func (p RsaOtherPrimesInfo) New() js.Ref {
	return bindings.RsaOtherPrimesInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RsaOtherPrimesInfo) UpdateFrom(ref js.Ref) {
	bindings.RsaOtherPrimesInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RsaOtherPrimesInfo) Update(ref js.Ref) {
	bindings.RsaOtherPrimesInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type JsonWebKey struct {
	// Kty is "JsonWebKey.kty"
	//
	// Optional
	Kty js.String
	// Use is "JsonWebKey.use"
	//
	// Optional
	Use js.String
	// KeyOps is "JsonWebKey.key_ops"
	//
	// Optional
	KeyOps js.Array[js.String]
	// Alg is "JsonWebKey.alg"
	//
	// Optional
	Alg js.String
	// Ext is "JsonWebKey.ext"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ext MUST be set to true to make this field effective.
	Ext bool
	// Crv is "JsonWebKey.crv"
	//
	// Optional
	Crv js.String
	// X is "JsonWebKey.x"
	//
	// Optional
	X js.String
	// Y is "JsonWebKey.y"
	//
	// Optional
	Y js.String
	// D is "JsonWebKey.d"
	//
	// Optional
	D js.String
	// N is "JsonWebKey.n"
	//
	// Optional
	N js.String
	// E is "JsonWebKey.e"
	//
	// Optional
	E js.String
	// P is "JsonWebKey.p"
	//
	// Optional
	P js.String
	// Q is "JsonWebKey.q"
	//
	// Optional
	Q js.String
	// Dp is "JsonWebKey.dp"
	//
	// Optional
	Dp js.String
	// Dq is "JsonWebKey.dq"
	//
	// Optional
	Dq js.String
	// Qi is "JsonWebKey.qi"
	//
	// Optional
	Qi js.String
	// Oth is "JsonWebKey.oth"
	//
	// Optional
	Oth js.Array[RsaOtherPrimesInfo]
	// K is "JsonWebKey.k"
	//
	// Optional
	K js.String

	FFI_USE_Ext bool // for Ext.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a JsonWebKey with all fields set.
func (p JsonWebKey) FromRef(ref js.Ref) JsonWebKey {
	p.UpdateFrom(ref)
	return p
}

// New creates a new JsonWebKey in the application heap.
func (p JsonWebKey) New() js.Ref {
	return bindings.JsonWebKeyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p JsonWebKey) UpdateFrom(ref js.Ref) {
	bindings.JsonWebKeyJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p JsonWebKey) Update(ref js.Ref) {
	bindings.JsonWebKeyJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey) JsonWebKey() JsonWebKey {
	var ret JsonWebKey
	ret.UpdateFrom(x.ref)
	return ret
}

type SubtleCrypto struct {
	ref js.Ref
}

func (this SubtleCrypto) Once() SubtleCrypto {
	this.Ref().Once()
	return this
}

func (this SubtleCrypto) Ref() js.Ref {
	return this.ref
}

func (this SubtleCrypto) FromRef(ref js.Ref) SubtleCrypto {
	this.ref = ref
	return this
}

func (this SubtleCrypto) Free() {
	this.Ref().Free()
}

// HasEncrypt returns true if the method "SubtleCrypto.encrypt" exists.
func (this SubtleCrypto) HasEncrypt() bool {
	return js.True == bindings.HasSubtleCryptoEncrypt(
		this.Ref(),
	)
}

// EncryptFunc returns the method "SubtleCrypto.encrypt".
func (this SubtleCrypto) EncryptFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoEncryptFunc(
			this.Ref(),
		),
	)
}

// Encrypt calls the method "SubtleCrypto.encrypt".
func (this SubtleCrypto) Encrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoEncrypt(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// TryEncrypt calls the method "SubtleCrypto.encrypt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryEncrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoEncrypt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// HasDecrypt returns true if the method "SubtleCrypto.decrypt" exists.
func (this SubtleCrypto) HasDecrypt() bool {
	return js.True == bindings.HasSubtleCryptoDecrypt(
		this.Ref(),
	)
}

// DecryptFunc returns the method "SubtleCrypto.decrypt".
func (this SubtleCrypto) DecryptFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDecryptFunc(
			this.Ref(),
		),
	)
}

// Decrypt calls the method "SubtleCrypto.decrypt".
func (this SubtleCrypto) Decrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoDecrypt(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// TryDecrypt calls the method "SubtleCrypto.decrypt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryDecrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoDecrypt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// HasSign returns true if the method "SubtleCrypto.sign" exists.
func (this SubtleCrypto) HasSign() bool {
	return js.True == bindings.HasSubtleCryptoSign(
		this.Ref(),
	)
}

// SignFunc returns the method "SubtleCrypto.sign".
func (this SubtleCrypto) SignFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoSignFunc(
			this.Ref(),
		),
	)
}

// Sign calls the method "SubtleCrypto.sign".
func (this SubtleCrypto) Sign(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoSign(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// TrySign calls the method "SubtleCrypto.sign"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TrySign(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoSign(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return
}

// HasVerify returns true if the method "SubtleCrypto.verify" exists.
func (this SubtleCrypto) HasVerify() bool {
	return js.True == bindings.HasSubtleCryptoVerify(
		this.Ref(),
	)
}

// VerifyFunc returns the method "SubtleCrypto.verify".
func (this SubtleCrypto) VerifyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, signature BufferSource, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoVerifyFunc(
			this.Ref(),
		),
	)
}

// Verify calls the method "SubtleCrypto.verify".
func (this SubtleCrypto) Verify(algorithm AlgorithmIdentifier, key CryptoKey, signature BufferSource, data BufferSource) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoVerify(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		key.Ref(),
		signature.Ref(),
		data.Ref(),
	)

	return
}

// TryVerify calls the method "SubtleCrypto.verify"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryVerify(algorithm AlgorithmIdentifier, key CryptoKey, signature BufferSource, data BufferSource) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoVerify(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		key.Ref(),
		signature.Ref(),
		data.Ref(),
	)

	return
}

// HasDigest returns true if the method "SubtleCrypto.digest" exists.
func (this SubtleCrypto) HasDigest() bool {
	return js.True == bindings.HasSubtleCryptoDigest(
		this.Ref(),
	)
}

// DigestFunc returns the method "SubtleCrypto.digest".
func (this SubtleCrypto) DigestFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDigestFunc(
			this.Ref(),
		),
	)
}

// Digest calls the method "SubtleCrypto.digest".
func (this SubtleCrypto) Digest(algorithm AlgorithmIdentifier, data BufferSource) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoDigest(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		data.Ref(),
	)

	return
}

// TryDigest calls the method "SubtleCrypto.digest"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryDigest(algorithm AlgorithmIdentifier, data BufferSource) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoDigest(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		data.Ref(),
	)

	return
}

// HasGenerateKey returns true if the method "SubtleCrypto.generateKey" exists.
func (this SubtleCrypto) HasGenerateKey() bool {
	return js.True == bindings.HasSubtleCryptoGenerateKey(
		this.Ref(),
	)
}

// GenerateKeyFunc returns the method "SubtleCrypto.generateKey".
func (this SubtleCrypto) GenerateKeyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoGenerateKeyFunc(
			this.Ref(),
		),
	)
}

// GenerateKey calls the method "SubtleCrypto.generateKey".
func (this SubtleCrypto) GenerateKey(algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoGenerateKey(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// TryGenerateKey calls the method "SubtleCrypto.generateKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryGenerateKey(algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoGenerateKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// HasDeriveKey returns true if the method "SubtleCrypto.deriveKey" exists.
func (this SubtleCrypto) HasDeriveKey() bool {
	return js.True == bindings.HasSubtleCryptoDeriveKey(
		this.Ref(),
	)
}

// DeriveKeyFunc returns the method "SubtleCrypto.deriveKey".
func (this SubtleCrypto) DeriveKeyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, baseKey CryptoKey, derivedKeyType AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDeriveKeyFunc(
			this.Ref(),
		),
	)
}

// DeriveKey calls the method "SubtleCrypto.deriveKey".
func (this SubtleCrypto) DeriveKey(algorithm AlgorithmIdentifier, baseKey CryptoKey, derivedKeyType AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoDeriveKey(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		baseKey.Ref(),
		derivedKeyType.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// TryDeriveKey calls the method "SubtleCrypto.deriveKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryDeriveKey(algorithm AlgorithmIdentifier, baseKey CryptoKey, derivedKeyType AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoDeriveKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		baseKey.Ref(),
		derivedKeyType.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// HasDeriveBits returns true if the method "SubtleCrypto.deriveBits" exists.
func (this SubtleCrypto) HasDeriveBits() bool {
	return js.True == bindings.HasSubtleCryptoDeriveBits(
		this.Ref(),
	)
}

// DeriveBitsFunc returns the method "SubtleCrypto.deriveBits".
func (this SubtleCrypto) DeriveBitsFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, baseKey CryptoKey, length uint32) js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDeriveBitsFunc(
			this.Ref(),
		),
	)
}

// DeriveBits calls the method "SubtleCrypto.deriveBits".
func (this SubtleCrypto) DeriveBits(algorithm AlgorithmIdentifier, baseKey CryptoKey, length uint32) (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallSubtleCryptoDeriveBits(
		this.Ref(), js.Pointer(&ret),
		algorithm.Ref(),
		baseKey.Ref(),
		uint32(length),
	)

	return
}

// TryDeriveBits calls the method "SubtleCrypto.deriveBits"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryDeriveBits(algorithm AlgorithmIdentifier, baseKey CryptoKey, length uint32) (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoDeriveBits(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		algorithm.Ref(),
		baseKey.Ref(),
		uint32(length),
	)

	return
}

// HasImportKey returns true if the method "SubtleCrypto.importKey" exists.
func (this SubtleCrypto) HasImportKey() bool {
	return js.True == bindings.HasSubtleCryptoImportKey(
		this.Ref(),
	)
}

// ImportKeyFunc returns the method "SubtleCrypto.importKey".
func (this SubtleCrypto) ImportKeyFunc() (fn js.Func[func(format KeyFormat, keyData OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey, algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[CryptoKey]]) {
	return fn.FromRef(
		bindings.SubtleCryptoImportKeyFunc(
			this.Ref(),
		),
	)
}

// ImportKey calls the method "SubtleCrypto.importKey".
func (this SubtleCrypto) ImportKey(format KeyFormat, keyData OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey, algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[CryptoKey]) {
	bindings.CallSubtleCryptoImportKey(
		this.Ref(), js.Pointer(&ret),
		uint32(format),
		keyData.Ref(),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// TryImportKey calls the method "SubtleCrypto.importKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryImportKey(format KeyFormat, keyData OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey, algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[CryptoKey], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoImportKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(format),
		keyData.Ref(),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// HasExportKey returns true if the method "SubtleCrypto.exportKey" exists.
func (this SubtleCrypto) HasExportKey() bool {
	return js.True == bindings.HasSubtleCryptoExportKey(
		this.Ref(),
	)
}

// ExportKeyFunc returns the method "SubtleCrypto.exportKey".
func (this SubtleCrypto) ExportKeyFunc() (fn js.Func[func(format KeyFormat, key CryptoKey) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoExportKeyFunc(
			this.Ref(),
		),
	)
}

// ExportKey calls the method "SubtleCrypto.exportKey".
func (this SubtleCrypto) ExportKey(format KeyFormat, key CryptoKey) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoExportKey(
		this.Ref(), js.Pointer(&ret),
		uint32(format),
		key.Ref(),
	)

	return
}

// TryExportKey calls the method "SubtleCrypto.exportKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryExportKey(format KeyFormat, key CryptoKey) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoExportKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(format),
		key.Ref(),
	)

	return
}

// HasWrapKey returns true if the method "SubtleCrypto.wrapKey" exists.
func (this SubtleCrypto) HasWrapKey() bool {
	return js.True == bindings.HasSubtleCryptoWrapKey(
		this.Ref(),
	)
}

// WrapKeyFunc returns the method "SubtleCrypto.wrapKey".
func (this SubtleCrypto) WrapKeyFunc() (fn js.Func[func(format KeyFormat, key CryptoKey, wrappingKey CryptoKey, wrapAlgorithm AlgorithmIdentifier) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoWrapKeyFunc(
			this.Ref(),
		),
	)
}

// WrapKey calls the method "SubtleCrypto.wrapKey".
func (this SubtleCrypto) WrapKey(format KeyFormat, key CryptoKey, wrappingKey CryptoKey, wrapAlgorithm AlgorithmIdentifier) (ret js.Promise[js.Any]) {
	bindings.CallSubtleCryptoWrapKey(
		this.Ref(), js.Pointer(&ret),
		uint32(format),
		key.Ref(),
		wrappingKey.Ref(),
		wrapAlgorithm.Ref(),
	)

	return
}

// TryWrapKey calls the method "SubtleCrypto.wrapKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryWrapKey(format KeyFormat, key CryptoKey, wrappingKey CryptoKey, wrapAlgorithm AlgorithmIdentifier) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoWrapKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(format),
		key.Ref(),
		wrappingKey.Ref(),
		wrapAlgorithm.Ref(),
	)

	return
}

// HasUnwrapKey returns true if the method "SubtleCrypto.unwrapKey" exists.
func (this SubtleCrypto) HasUnwrapKey() bool {
	return js.True == bindings.HasSubtleCryptoUnwrapKey(
		this.Ref(),
	)
}

// UnwrapKeyFunc returns the method "SubtleCrypto.unwrapKey".
func (this SubtleCrypto) UnwrapKeyFunc() (fn js.Func[func(format KeyFormat, wrappedKey BufferSource, unwrappingKey CryptoKey, unwrapAlgorithm AlgorithmIdentifier, unwrappedKeyAlgorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[CryptoKey]]) {
	return fn.FromRef(
		bindings.SubtleCryptoUnwrapKeyFunc(
			this.Ref(),
		),
	)
}

// UnwrapKey calls the method "SubtleCrypto.unwrapKey".
func (this SubtleCrypto) UnwrapKey(format KeyFormat, wrappedKey BufferSource, unwrappingKey CryptoKey, unwrapAlgorithm AlgorithmIdentifier, unwrappedKeyAlgorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[CryptoKey]) {
	bindings.CallSubtleCryptoUnwrapKey(
		this.Ref(), js.Pointer(&ret),
		uint32(format),
		wrappedKey.Ref(),
		unwrappingKey.Ref(),
		unwrapAlgorithm.Ref(),
		unwrappedKeyAlgorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

// TryUnwrapKey calls the method "SubtleCrypto.unwrapKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SubtleCrypto) TryUnwrapKey(format KeyFormat, wrappedKey BufferSource, unwrappingKey CryptoKey, unwrapAlgorithm AlgorithmIdentifier, unwrappedKeyAlgorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (ret js.Promise[CryptoKey], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCryptoUnwrapKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(format),
		wrappedKey.Ref(),
		unwrappingKey.Ref(),
		unwrapAlgorithm.Ref(),
		unwrappedKeyAlgorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return
}

type Crypto struct {
	ref js.Ref
}

func (this Crypto) Once() Crypto {
	this.Ref().Once()
	return this
}

func (this Crypto) Ref() js.Ref {
	return this.ref
}

func (this Crypto) FromRef(ref js.Ref) Crypto {
	this.ref = ref
	return this
}

func (this Crypto) Free() {
	this.Ref().Free()
}

// Subtle returns the value of property "Crypto.subtle".
//
// It returns ok=false if there is no such property.
func (this Crypto) Subtle() (ret SubtleCrypto, ok bool) {
	ok = js.True == bindings.GetCryptoSubtle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetRandomValues returns true if the method "Crypto.getRandomValues" exists.
func (this Crypto) HasGetRandomValues() bool {
	return js.True == bindings.HasCryptoGetRandomValues(
		this.Ref(),
	)
}

// GetRandomValuesFunc returns the method "Crypto.getRandomValues".
func (this Crypto) GetRandomValuesFunc() (fn js.Func[func(array js.ArrayBufferView) js.ArrayBufferView]) {
	return fn.FromRef(
		bindings.CryptoGetRandomValuesFunc(
			this.Ref(),
		),
	)
}

// GetRandomValues calls the method "Crypto.getRandomValues".
func (this Crypto) GetRandomValues(array js.ArrayBufferView) (ret js.ArrayBufferView) {
	bindings.CallCryptoGetRandomValues(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetRandomValues calls the method "Crypto.getRandomValues"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Crypto) TryGetRandomValues(array js.ArrayBufferView) (ret js.ArrayBufferView, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCryptoGetRandomValues(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasRandomUUID returns true if the method "Crypto.randomUUID" exists.
func (this Crypto) HasRandomUUID() bool {
	return js.True == bindings.HasCryptoRandomUUID(
		this.Ref(),
	)
}

// RandomUUIDFunc returns the method "Crypto.randomUUID".
func (this Crypto) RandomUUIDFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CryptoRandomUUIDFunc(
			this.Ref(),
		),
	)
}

// RandomUUID calls the method "Crypto.randomUUID".
func (this Crypto) RandomUUID() (ret js.String) {
	bindings.CallCryptoRandomUUID(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRandomUUID calls the method "Crypto.randomUUID"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Crypto) TryRandomUUID() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCryptoRandomUUID(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MemoryAttributionContainer struct {
	// Id is "MemoryAttributionContainer.id"
	//
	// Optional
	Id js.String
	// Src is "MemoryAttributionContainer.src"
	//
	// Optional
	Src js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryAttributionContainer with all fields set.
func (p MemoryAttributionContainer) FromRef(ref js.Ref) MemoryAttributionContainer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryAttributionContainer in the application heap.
func (p MemoryAttributionContainer) New() js.Ref {
	return bindings.MemoryAttributionContainerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MemoryAttributionContainer) UpdateFrom(ref js.Ref) {
	bindings.MemoryAttributionContainerJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MemoryAttributionContainer) Update(ref js.Ref) {
	bindings.MemoryAttributionContainerJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MemoryAttribution struct {
	// Url is "MemoryAttribution.url"
	//
	// Optional
	Url js.String
	// Container is "MemoryAttribution.container"
	//
	// Optional
	//
	// NOTE: Container.FFI_USE MUST be set to true to get Container used.
	Container MemoryAttributionContainer
	// Scope is "MemoryAttribution.scope"
	//
	// Optional
	Scope js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryAttribution with all fields set.
func (p MemoryAttribution) FromRef(ref js.Ref) MemoryAttribution {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryAttribution in the application heap.
func (p MemoryAttribution) New() js.Ref {
	return bindings.MemoryAttributionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MemoryAttribution) UpdateFrom(ref js.Ref) {
	bindings.MemoryAttributionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MemoryAttribution) Update(ref js.Ref) {
	bindings.MemoryAttributionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MemoryBreakdownEntry struct {
	// Bytes is "MemoryBreakdownEntry.bytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bytes MUST be set to true to make this field effective.
	Bytes uint64
	// Attribution is "MemoryBreakdownEntry.attribution"
	//
	// Optional
	Attribution js.Array[MemoryAttribution]
	// Types is "MemoryBreakdownEntry.types"
	//
	// Optional
	Types js.Array[js.String]

	FFI_USE_Bytes bool // for Bytes.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryBreakdownEntry with all fields set.
func (p MemoryBreakdownEntry) FromRef(ref js.Ref) MemoryBreakdownEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryBreakdownEntry in the application heap.
func (p MemoryBreakdownEntry) New() js.Ref {
	return bindings.MemoryBreakdownEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MemoryBreakdownEntry) UpdateFrom(ref js.Ref) {
	bindings.MemoryBreakdownEntryJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MemoryBreakdownEntry) Update(ref js.Ref) {
	bindings.MemoryBreakdownEntryJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MemoryMeasurement struct {
	// Bytes is "MemoryMeasurement.bytes"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bytes MUST be set to true to make this field effective.
	Bytes uint64
	// Breakdown is "MemoryMeasurement.breakdown"
	//
	// Optional
	Breakdown js.Array[MemoryBreakdownEntry]

	FFI_USE_Bytes bool // for Bytes.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryMeasurement with all fields set.
func (p MemoryMeasurement) FromRef(ref js.Ref) MemoryMeasurement {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryMeasurement in the application heap.
func (p MemoryMeasurement) New() js.Ref {
	return bindings.MemoryMeasurementJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MemoryMeasurement) UpdateFrom(ref js.Ref) {
	bindings.MemoryMeasurementJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MemoryMeasurement) Update(ref js.Ref) {
	bindings.MemoryMeasurementJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PerformanceMarkOptions struct {
	// Detail is "PerformanceMarkOptions.detail"
	//
	// Optional
	Detail js.Any
	// StartTime is "PerformanceMarkOptions.startTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartTime MUST be set to true to make this field effective.
	StartTime DOMHighResTimeStamp

	FFI_USE_StartTime bool // for StartTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceMarkOptions with all fields set.
func (p PerformanceMarkOptions) FromRef(ref js.Ref) PerformanceMarkOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerformanceMarkOptions in the application heap.
func (p PerformanceMarkOptions) New() js.Ref {
	return bindings.PerformanceMarkOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PerformanceMarkOptions) UpdateFrom(ref js.Ref) {
	bindings.PerformanceMarkOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PerformanceMarkOptions) Update(ref js.Ref) {
	bindings.PerformanceMarkOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPerformanceMark(markName js.String, markOptions PerformanceMarkOptions) (ret PerformanceMark) {
	ret.ref = bindings.NewPerformanceMarkByPerformanceMark(
		markName.Ref(),
		js.Pointer(&markOptions))
	return
}

func NewPerformanceMarkByPerformanceMark1(markName js.String) (ret PerformanceMark) {
	ret.ref = bindings.NewPerformanceMarkByPerformanceMark1(
		markName.Ref())
	return
}

type PerformanceMark struct {
	PerformanceEntry
}

func (this PerformanceMark) Once() PerformanceMark {
	this.Ref().Once()
	return this
}

func (this PerformanceMark) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceMark) FromRef(ref js.Ref) PerformanceMark {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceMark) Free() {
	this.Ref().Free()
}

// Detail returns the value of property "PerformanceMark.detail".
//
// It returns ok=false if there is no such property.
func (this PerformanceMark) Detail() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetPerformanceMarkDetail(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PerformanceMeasure struct {
	PerformanceEntry
}

func (this PerformanceMeasure) Once() PerformanceMeasure {
	this.Ref().Once()
	return this
}

func (this PerformanceMeasure) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceMeasure) FromRef(ref js.Ref) PerformanceMeasure {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceMeasure) Free() {
	this.Ref().Free()
}

// Detail returns the value of property "PerformanceMeasure.detail".
//
// It returns ok=false if there is no such property.
func (this PerformanceMeasure) Detail() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetPerformanceMeasureDetail(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_String_Float64 struct {
	ref js.Ref
}

func (x OneOf_String_Float64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Float64) Free() {
	x.ref.Free()
}

func (x OneOf_String_Float64) FromRef(ref js.Ref) OneOf_String_Float64 {
	return OneOf_String_Float64{
		ref: ref,
	}
}

func (x OneOf_String_Float64) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Float64) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

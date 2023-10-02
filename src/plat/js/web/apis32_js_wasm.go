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
// The returned bool will be false if there is no such property.
func (this WindowSharedStorage) Worklet() (SharedStorageWorklet, bool) {
	var _ok bool
	_ret := bindings.GetWindowSharedStorageWorklet(
		this.Ref(), js.Pointer(&_ok),
	)
	return SharedStorageWorklet{}.FromRef(_ret), _ok
}

// Run calls the method "WindowSharedStorage.run".
//
// The returned bool will be false if there is no such method.
func (this WindowSharedStorage) Run(name js.String, options SharedStorageRunOperationMethodOptions) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallWindowSharedStorageRun(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// RunFunc returns the method "WindowSharedStorage.run".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowSharedStorage) RunFunc() (fn js.Func[func(name js.String, options SharedStorageRunOperationMethodOptions) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageRunFunc(
			this.Ref(),
		),
	)
}

// Run1 calls the method "WindowSharedStorage.run".
//
// The returned bool will be false if there is no such method.
func (this WindowSharedStorage) Run1(name js.String) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallWindowSharedStorageRun1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// Run1Func returns the method "WindowSharedStorage.run".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowSharedStorage) Run1Func() (fn js.Func[func(name js.String) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageRun1Func(
			this.Ref(),
		),
	)
}

// SelectURL calls the method "WindowSharedStorage.selectURL".
//
// The returned bool will be false if there is no such method.
func (this WindowSharedStorage) SelectURL(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata], options SharedStorageRunOperationMethodOptions) (js.Promise[SharedStorageResponse], bool) {
	var _ok bool
	_ret := bindings.CallWindowSharedStorageSelectURL(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		urls.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[SharedStorageResponse]{}.FromRef(_ret), _ok
}

// SelectURLFunc returns the method "WindowSharedStorage.selectURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowSharedStorage) SelectURLFunc() (fn js.Func[func(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata], options SharedStorageRunOperationMethodOptions) js.Promise[SharedStorageResponse]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageSelectURLFunc(
			this.Ref(),
		),
	)
}

// SelectURL1 calls the method "WindowSharedStorage.selectURL".
//
// The returned bool will be false if there is no such method.
func (this WindowSharedStorage) SelectURL1(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (js.Promise[SharedStorageResponse], bool) {
	var _ok bool
	_ret := bindings.CallWindowSharedStorageSelectURL1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		urls.Ref(),
	)

	return js.Promise[SharedStorageResponse]{}.FromRef(_ret), _ok
}

// SelectURL1Func returns the method "WindowSharedStorage.selectURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowSharedStorage) SelectURL1Func() (fn js.Func[func(name js.String, urls js.FrozenArray[SharedStorageUrlWithMetadata]) js.Promise[SharedStorageResponse]]) {
	return fn.FromRef(
		bindings.WindowSharedStorageSelectURL1Func(
			this.Ref(),
		),
	)
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
	Expires DOMHighResTimeStamp
	// Secure is "CookieListItem.secure"
	//
	// Optional
	Secure bool
	// SameSite is "CookieListItem.sameSite"
	//
	// Optional
	SameSite CookieSameSite
	// Partitioned is "CookieListItem.partitioned"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 CookieListItem CookieListItem [// CookieListItem] [0x140007da3c0 0x140007da460 0x140007da500 0x140007da5a0 0x140007da640 0x140007da780 0x140007da8c0 0x140007da960 0x140007da6e0 0x140007da820 0x140007daa00] 0x1400081fad0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 CookieInit CookieInit [// CookieInit] [0x140007daaa0 0x140007dab40 0x140007dabe0 0x140007dad20 0x140007dadc0 0x140007dae60 0x140007daf00 0x140007dac80 0x140007dafa0] 0x1400081fba8 {0 0}} in the application heap.
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
	Partitioned bool

	FFI_USE_Partitioned bool // for Partitioned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieStoreDeleteOptions with all fields set.
func (p CookieStoreDeleteOptions) FromRef(ref js.Ref) CookieStoreDeleteOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CookieStoreDeleteOptions CookieStoreDeleteOptions [// CookieStoreDeleteOptions] [0x140007db040 0x140007db0e0 0x140007db180 0x140007db220 0x140007db2c0] 0x1400081fc08 {0 0}} in the application heap.
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

// Get calls the method "CookieStore.get".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Get(name js.String) (js.Promise[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[CookieListItem]{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CookieStore.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) GetFunc() (fn js.Func[func(name js.String) js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGetFunc(
			this.Ref(),
		),
	)
}

// Get1 calls the method "CookieStore.get".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Get1(options CookieStoreGetOptions) (js.Promise[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGet1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[CookieListItem]{}.FromRef(_ret), _ok
}

// Get1Func returns the method "CookieStore.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) Get1Func() (fn js.Func[func(options CookieStoreGetOptions) js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGet1Func(
			this.Ref(),
		),
	)
}

// Get2 calls the method "CookieStore.get".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Get2() (js.Promise[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGet2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[CookieListItem]{}.FromRef(_ret), _ok
}

// Get2Func returns the method "CookieStore.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) Get2Func() (fn js.Func[func() js.Promise[CookieListItem]]) {
	return fn.FromRef(
		bindings.CookieStoreGet2Func(
			this.Ref(),
		),
	)
}

// GetAll calls the method "CookieStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) GetAll(name js.String) (js.Promise[CookieList], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGetAll(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[CookieList]{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "CookieStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) GetAllFunc() (fn js.Func[func(name js.String) js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "CookieStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) GetAll1(options CookieStoreGetOptions) (js.Promise[CookieList], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGetAll1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[CookieList]{}.FromRef(_ret), _ok
}

// GetAll1Func returns the method "CookieStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) GetAll1Func() (fn js.Func[func(options CookieStoreGetOptions) js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "CookieStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) GetAll2() (js.Promise[CookieList], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreGetAll2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[CookieList]{}.FromRef(_ret), _ok
}

// GetAll2Func returns the method "CookieStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) GetAll2Func() (fn js.Func[func() js.Promise[CookieList]]) {
	return fn.FromRef(
		bindings.CookieStoreGetAll2Func(
			this.Ref(),
		),
	)
}

// Set calls the method "CookieStore.set".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Set(name js.String, value js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreSet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetFunc returns the method "CookieStore.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) SetFunc() (fn js.Func[func(name js.String, value js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreSetFunc(
			this.Ref(),
		),
	)
}

// Set1 calls the method "CookieStore.set".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Set1(options CookieInit) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreSet1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Set1Func returns the method "CookieStore.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) Set1Func() (fn js.Func[func(options CookieInit) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreSet1Func(
			this.Ref(),
		),
	)
}

// Delete calls the method "CookieStore.delete".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Delete(name js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreDelete(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "CookieStore.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) DeleteFunc() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete1 calls the method "CookieStore.delete".
//
// The returned bool will be false if there is no such method.
func (this CookieStore) Delete1(options CookieStoreDeleteOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreDelete1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Delete1Func returns the method "CookieStore.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStore) Delete1Func() (fn js.Func[func(options CookieStoreDeleteOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreDelete1Func(
			this.Ref(),
		),
	)
}

type DocumentPictureInPictureOptions struct {
	// Width is "DocumentPictureInPictureOptions.width"
	//
	// Optional, defaults to 0.
	Width uint64
	// Height is "DocumentPictureInPictureOptions.height"
	//
	// Optional, defaults to 0.
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

// New creates a new {0x140004cc0e0 DocumentPictureInPictureOptions DocumentPictureInPictureOptions [// DocumentPictureInPictureOptions] [0x140007db360 0x140007db4a0 0x140007db400 0x140007db540] 0x1400081fc80 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this DocumentPictureInPicture) Window() (Window, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPictureInPictureWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return Window{}.FromRef(_ret), _ok
}

// RequestWindow calls the method "DocumentPictureInPicture.requestWindow".
//
// The returned bool will be false if there is no such method.
func (this DocumentPictureInPicture) RequestWindow(options DocumentPictureInPictureOptions) (js.Promise[Window], bool) {
	var _ok bool
	_ret := bindings.CallDocumentPictureInPictureRequestWindow(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[Window]{}.FromRef(_ret), _ok
}

// RequestWindowFunc returns the method "DocumentPictureInPicture.requestWindow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentPictureInPicture) RequestWindowFunc() (fn js.Func[func(options DocumentPictureInPictureOptions) js.Promise[Window]]) {
	return fn.FromRef(
		bindings.DocumentPictureInPictureRequestWindowFunc(
			this.Ref(),
		),
	)
}

// RequestWindow1 calls the method "DocumentPictureInPicture.requestWindow".
//
// The returned bool will be false if there is no such method.
func (this DocumentPictureInPicture) RequestWindow1() (js.Promise[Window], bool) {
	var _ok bool
	_ret := bindings.CallDocumentPictureInPictureRequestWindow1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Window]{}.FromRef(_ret), _ok
}

// RequestWindow1Func returns the method "DocumentPictureInPicture.requestWindow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DocumentPictureInPicture) RequestWindow1Func() (fn js.Func[func() js.Promise[Window]]) {
	return fn.FromRef(
		bindings.DocumentPictureInPictureRequestWindow1Func(
			this.Ref(),
		),
	)
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

// AddSearchProvider calls the method "External.AddSearchProvider".
//
// The returned bool will be false if there is no such method.
func (this External) AddSearchProvider() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallExternalAddSearchProvider(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddSearchProviderFunc returns the method "External.AddSearchProvider".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this External) AddSearchProviderFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ExternalAddSearchProviderFunc(
			this.Ref(),
		),
	)
}

// IsSearchProviderInstalled calls the method "External.IsSearchProviderInstalled".
//
// The returned bool will be false if there is no such method.
func (this External) IsSearchProviderInstalled() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallExternalIsSearchProviderInstalled(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// IsSearchProviderInstalledFunc returns the method "External.IsSearchProviderInstalled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this External) IsSearchProviderInstalledFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ExternalIsSearchProviderInstalledFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisVoice) VoiceURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisVoiceVoiceURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "SpeechSynthesisVoice.name".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisVoice) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisVoiceName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lang returns the value of property "SpeechSynthesisVoice.lang".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisVoice) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisVoiceLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LocalService returns the value of property "SpeechSynthesisVoice.localService".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisVoice) LocalService() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisVoiceLocalService(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Default returns the value of property "SpeechSynthesisVoice.default".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisVoice) Default() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisVoiceDefault(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

func NewSpeechSynthesisUtterance(text js.String) SpeechSynthesisUtterance {
	return SpeechSynthesisUtterance{}.FromRef(
		bindings.NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance(
			text.Ref()),
	)
}

func NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance1() SpeechSynthesisUtterance {
	return SpeechSynthesisUtterance{}.FromRef(
		bindings.NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance1(),
	)
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtteranceText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "SpeechSynthesisUtterance.text" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtteranceLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lang sets the value of property "SpeechSynthesisUtterance.lang" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Voice() (SpeechSynthesisVoice, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtteranceVoice(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechSynthesisVoice{}.FromRef(_ret), _ok
}

// Voice sets the value of property "SpeechSynthesisUtterance.voice" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Volume() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtteranceVolume(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Volume sets the value of property "SpeechSynthesisUtterance.volume" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Rate() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtteranceRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Rate sets the value of property "SpeechSynthesisUtterance.rate" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisUtterance) Pitch() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisUtterancePitch(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Pitch sets the value of property "SpeechSynthesisUtterance.pitch" to val.
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
// The returned bool will be false if there is no such property.
func (this SpeechSynthesis) Pending() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisPending(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Speaking returns the value of property "SpeechSynthesis.speaking".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesis) Speaking() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisSpeaking(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Paused returns the value of property "SpeechSynthesis.paused".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesis) Paused() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisPaused(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Speak calls the method "SpeechSynthesis.speak".
//
// The returned bool will be false if there is no such method.
func (this SpeechSynthesis) Speak(utterance SpeechSynthesisUtterance) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechSynthesisSpeak(
		this.Ref(), js.Pointer(&_ok),
		utterance.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SpeakFunc returns the method "SpeechSynthesis.speak".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechSynthesis) SpeakFunc() (fn js.Func[func(utterance SpeechSynthesisUtterance)]) {
	return fn.FromRef(
		bindings.SpeechSynthesisSpeakFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "SpeechSynthesis.cancel".
//
// The returned bool will be false if there is no such method.
func (this SpeechSynthesis) Cancel() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechSynthesisCancel(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelFunc returns the method "SpeechSynthesis.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechSynthesis) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisCancelFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "SpeechSynthesis.pause".
//
// The returned bool will be false if there is no such method.
func (this SpeechSynthesis) Pause() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechSynthesisPause(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseFunc returns the method "SpeechSynthesis.pause".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechSynthesis) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisPauseFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "SpeechSynthesis.resume".
//
// The returned bool will be false if there is no such method.
func (this SpeechSynthesis) Resume() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechSynthesisResume(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResumeFunc returns the method "SpeechSynthesis.resume".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechSynthesis) ResumeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechSynthesisResumeFunc(
			this.Ref(),
		),
	)
}

// GetVoices calls the method "SpeechSynthesis.getVoices".
//
// The returned bool will be false if there is no such method.
func (this SpeechSynthesis) GetVoices() (js.Array[SpeechSynthesisVoice], bool) {
	var _ok bool
	_ret := bindings.CallSpeechSynthesisGetVoices(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[SpeechSynthesisVoice]{}.FromRef(_ret), _ok
}

// GetVoicesFunc returns the method "SpeechSynthesis.getVoices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechSynthesis) GetVoicesFunc() (fn js.Func[func() js.Array[SpeechSynthesisVoice]]) {
	return fn.FromRef(
		bindings.SpeechSynthesisGetVoicesFunc(
			this.Ref(),
		),
	)
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
	Delay uint64

	FFI_USE_Delay bool // for Delay.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SchedulerPostTaskOptions with all fields set.
func (p SchedulerPostTaskOptions) FromRef(ref js.Ref) SchedulerPostTaskOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SchedulerPostTaskOptions SchedulerPostTaskOptions [// SchedulerPostTaskOptions] [0x140007db7c0 0x140007db860 0x140007db900 0x140007db9a0] 0x1400081fcb0 {0 0}} in the application heap.
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

// PostTask calls the method "Scheduler.postTask".
//
// The returned bool will be false if there is no such method.
func (this Scheduler) PostTask(callback js.Func[func() js.Any], options SchedulerPostTaskOptions) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSchedulerPostTask(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// PostTaskFunc returns the method "Scheduler.postTask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Scheduler) PostTaskFunc() (fn js.Func[func(callback js.Func[func() js.Any], options SchedulerPostTaskOptions) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SchedulerPostTaskFunc(
			this.Ref(),
		),
	)
}

// PostTask1 calls the method "Scheduler.postTask".
//
// The returned bool will be false if there is no such method.
func (this Scheduler) PostTask1(callback js.Func[func() js.Any]) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSchedulerPostTask1(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// PostTask1Func returns the method "Scheduler.postTask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Scheduler) PostTask1Func() (fn js.Func[func(callback js.Func[func() js.Any]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SchedulerPostTask1Func(
			this.Ref(),
		),
	)
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

// ToString calls the method "TrustedHTML.toString".
//
// The returned bool will be false if there is no such method.
func (this TrustedHTML) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedHTMLToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "TrustedHTML.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedHTML) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedHTMLToStringFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedHTML.toJSON".
//
// The returned bool will be false if there is no such method.
func (this TrustedHTML) ToJSON() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedHTMLToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "TrustedHTML.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedHTML) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedHTMLToJSONFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedHTML.fromLiteral".
//
// The returned bool will be false if there is no such method.
func (this TrustedHTML) FromLiteral(templateStringsArray js.Object) (TrustedHTML, bool) {
	var _ok bool
	_ret := bindings.CallTrustedHTMLFromLiteral(
		this.Ref(), js.Pointer(&_ok),
		templateStringsArray.Ref(),
	)

	return TrustedHTML{}.FromRef(_ret), _ok
}

// FromLiteralFunc returns the staticmethod "TrustedHTML.fromLiteral".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedHTML) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedHTML]) {
	return fn.FromRef(
		bindings.TrustedHTMLFromLiteralFunc(
			this.Ref(),
		),
	)
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

// ToString calls the method "TrustedScript.toString".
//
// The returned bool will be false if there is no such method.
func (this TrustedScript) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "TrustedScript.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScript) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptToStringFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedScript.toJSON".
//
// The returned bool will be false if there is no such method.
func (this TrustedScript) ToJSON() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "TrustedScript.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScript) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptToJSONFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedScript.fromLiteral".
//
// The returned bool will be false if there is no such method.
func (this TrustedScript) FromLiteral(templateStringsArray js.Object) (TrustedScript, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptFromLiteral(
		this.Ref(), js.Pointer(&_ok),
		templateStringsArray.Ref(),
	)

	return TrustedScript{}.FromRef(_ret), _ok
}

// FromLiteralFunc returns the staticmethod "TrustedScript.fromLiteral".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScript) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedScript]) {
	return fn.FromRef(
		bindings.TrustedScriptFromLiteralFunc(
			this.Ref(),
		),
	)
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

// ToString calls the method "TrustedScriptURL.toString".
//
// The returned bool will be false if there is no such method.
func (this TrustedScriptURL) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptURLToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "TrustedScriptURL.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScriptURL) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptURLToStringFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TrustedScriptURL.toJSON".
//
// The returned bool will be false if there is no such method.
func (this TrustedScriptURL) ToJSON() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptURLToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "TrustedScriptURL.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScriptURL) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.TrustedScriptURLToJSONFunc(
			this.Ref(),
		),
	)
}

// FromLiteral calls the staticmethod "TrustedScriptURL.fromLiteral".
//
// The returned bool will be false if there is no such method.
func (this TrustedScriptURL) FromLiteral(templateStringsArray js.Object) (TrustedScriptURL, bool) {
	var _ok bool
	_ret := bindings.CallTrustedScriptURLFromLiteral(
		this.Ref(), js.Pointer(&_ok),
		templateStringsArray.Ref(),
	)

	return TrustedScriptURL{}.FromRef(_ret), _ok
}

// FromLiteralFunc returns the staticmethod "TrustedScriptURL.fromLiteral".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedScriptURL) FromLiteralFunc() (fn js.Func[func(templateStringsArray js.Object) TrustedScriptURL]) {
	return fn.FromRef(
		bindings.TrustedScriptURLFromLiteralFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this TrustedTypePolicy) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTrustedTypePolicyName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CreateHTML calls the method "TrustedTypePolicy.createHTML".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicy) CreateHTML(input js.String, arguments ...js.Any) (TrustedHTML, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyCreateHTML(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return TrustedHTML{}.FromRef(_ret), _ok
}

// CreateHTMLFunc returns the method "TrustedTypePolicy.createHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicy) CreateHTMLFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedHTML]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateHTMLFunc(
			this.Ref(),
		),
	)
}

// CreateScript calls the method "TrustedTypePolicy.createScript".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicy) CreateScript(input js.String, arguments ...js.Any) (TrustedScript, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyCreateScript(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return TrustedScript{}.FromRef(_ret), _ok
}

// CreateScriptFunc returns the method "TrustedTypePolicy.createScript".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicy) CreateScriptFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedScript]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateScriptFunc(
			this.Ref(),
		),
	)
}

// CreateScriptURL calls the method "TrustedTypePolicy.createScriptURL".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicy) CreateScriptURL(input js.String, arguments ...js.Any) (TrustedScriptURL, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyCreateScriptURL(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return TrustedScriptURL{}.FromRef(_ret), _ok
}

// CreateScriptURLFunc returns the method "TrustedTypePolicy.createScriptURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicy) CreateScriptURLFunc() (fn js.Func[func(input js.String, arguments ...js.Any) TrustedScriptURL]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyCreateScriptURLFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 TrustedTypePolicyOptions TrustedTypePolicyOptions [// TrustedTypePolicyOptions] [0x140007dba40 0x140007dbae0 0x140007dbb80] 0x1400081fd10 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this TrustedTypePolicyFactory) EmptyHTML() (TrustedHTML, bool) {
	var _ok bool
	_ret := bindings.GetTrustedTypePolicyFactoryEmptyHTML(
		this.Ref(), js.Pointer(&_ok),
	)
	return TrustedHTML{}.FromRef(_ret), _ok
}

// EmptyScript returns the value of property "TrustedTypePolicyFactory.emptyScript".
//
// The returned bool will be false if there is no such property.
func (this TrustedTypePolicyFactory) EmptyScript() (TrustedScript, bool) {
	var _ok bool
	_ret := bindings.GetTrustedTypePolicyFactoryEmptyScript(
		this.Ref(), js.Pointer(&_ok),
	)
	return TrustedScript{}.FromRef(_ret), _ok
}

// DefaultPolicy returns the value of property "TrustedTypePolicyFactory.defaultPolicy".
//
// The returned bool will be false if there is no such property.
func (this TrustedTypePolicyFactory) DefaultPolicy() (TrustedTypePolicy, bool) {
	var _ok bool
	_ret := bindings.GetTrustedTypePolicyFactoryDefaultPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return TrustedTypePolicy{}.FromRef(_ret), _ok
}

// CreatePolicy calls the method "TrustedTypePolicyFactory.createPolicy".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) CreatePolicy(policyName js.String, policyOptions TrustedTypePolicyOptions) (TrustedTypePolicy, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryCreatePolicy(
		this.Ref(), js.Pointer(&_ok),
		policyName.Ref(),
		js.Pointer(&policyOptions),
	)

	return TrustedTypePolicy{}.FromRef(_ret), _ok
}

// CreatePolicyFunc returns the method "TrustedTypePolicyFactory.createPolicy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) CreatePolicyFunc() (fn js.Func[func(policyName js.String, policyOptions TrustedTypePolicyOptions) TrustedTypePolicy]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryCreatePolicyFunc(
			this.Ref(),
		),
	)
}

// CreatePolicy1 calls the method "TrustedTypePolicyFactory.createPolicy".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) CreatePolicy1(policyName js.String) (TrustedTypePolicy, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryCreatePolicy1(
		this.Ref(), js.Pointer(&_ok),
		policyName.Ref(),
	)

	return TrustedTypePolicy{}.FromRef(_ret), _ok
}

// CreatePolicy1Func returns the method "TrustedTypePolicyFactory.createPolicy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) CreatePolicy1Func() (fn js.Func[func(policyName js.String) TrustedTypePolicy]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryCreatePolicy1Func(
			this.Ref(),
		),
	)
}

// IsHTML calls the method "TrustedTypePolicyFactory.isHTML".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) IsHTML(value js.Any) (bool, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryIsHTML(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return _ret == js.True, _ok
}

// IsHTMLFunc returns the method "TrustedTypePolicyFactory.isHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) IsHTMLFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsHTMLFunc(
			this.Ref(),
		),
	)
}

// IsScript calls the method "TrustedTypePolicyFactory.isScript".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) IsScript(value js.Any) (bool, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryIsScript(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return _ret == js.True, _ok
}

// IsScriptFunc returns the method "TrustedTypePolicyFactory.isScript".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) IsScriptFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsScriptFunc(
			this.Ref(),
		),
	)
}

// IsScriptURL calls the method "TrustedTypePolicyFactory.isScriptURL".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) IsScriptURL(value js.Any) (bool, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryIsScriptURL(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return _ret == js.True, _ok
}

// IsScriptURLFunc returns the method "TrustedTypePolicyFactory.isScriptURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) IsScriptURLFunc() (fn js.Func[func(value js.Any) bool]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryIsScriptURLFunc(
			this.Ref(),
		),
	)
}

// GetAttributeType calls the method "TrustedTypePolicyFactory.getAttributeType".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeType(tagName js.String, attribute js.String, elementNs js.String, attrNs js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryGetAttributeType(
		this.Ref(), js.Pointer(&_ok),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
		attrNs.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAttributeTypeFunc returns the method "TrustedTypePolicyFactory.getAttributeType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeTypeFunc() (fn js.Func[func(tagName js.String, attribute js.String, elementNs js.String, attrNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeTypeFunc(
			this.Ref(),
		),
	)
}

// GetAttributeType1 calls the method "TrustedTypePolicyFactory.getAttributeType".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeType1(tagName js.String, attribute js.String, elementNs js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryGetAttributeType1(
		this.Ref(), js.Pointer(&_ok),
		tagName.Ref(),
		attribute.Ref(),
		elementNs.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAttributeType1Func returns the method "TrustedTypePolicyFactory.getAttributeType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeType1Func() (fn js.Func[func(tagName js.String, attribute js.String, elementNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeType1Func(
			this.Ref(),
		),
	)
}

// GetAttributeType2 calls the method "TrustedTypePolicyFactory.getAttributeType".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeType2(tagName js.String, attribute js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryGetAttributeType2(
		this.Ref(), js.Pointer(&_ok),
		tagName.Ref(),
		attribute.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetAttributeType2Func returns the method "TrustedTypePolicyFactory.getAttributeType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) GetAttributeType2Func() (fn js.Func[func(tagName js.String, attribute js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetAttributeType2Func(
			this.Ref(),
		),
	)
}

// GetPropertyType calls the method "TrustedTypePolicyFactory.getPropertyType".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) GetPropertyType(tagName js.String, property js.String, elementNs js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryGetPropertyType(
		this.Ref(), js.Pointer(&_ok),
		tagName.Ref(),
		property.Ref(),
		elementNs.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetPropertyTypeFunc returns the method "TrustedTypePolicyFactory.getPropertyType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) GetPropertyTypeFunc() (fn js.Func[func(tagName js.String, property js.String, elementNs js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetPropertyTypeFunc(
			this.Ref(),
		),
	)
}

// GetPropertyType1 calls the method "TrustedTypePolicyFactory.getPropertyType".
//
// The returned bool will be false if there is no such method.
func (this TrustedTypePolicyFactory) GetPropertyType1(tagName js.String, property js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallTrustedTypePolicyFactoryGetPropertyType1(
		this.Ref(), js.Pointer(&_ok),
		tagName.Ref(),
		property.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetPropertyType1Func returns the method "TrustedTypePolicyFactory.getPropertyType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TrustedTypePolicyFactory) GetPropertyType1Func() (fn js.Func[func(tagName js.String, property js.String) js.String]) {
	return fn.FromRef(
		bindings.TrustedTypePolicyFactoryGetPropertyType1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this CryptoKey) Type() (KeyType, bool) {
	var _ok bool
	_ret := bindings.GetCryptoKeyType(
		this.Ref(), js.Pointer(&_ok),
	)
	return KeyType(_ret), _ok
}

// Extractable returns the value of property "CryptoKey.extractable".
//
// The returned bool will be false if there is no such property.
func (this CryptoKey) Extractable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCryptoKeyExtractable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Algorithm returns the value of property "CryptoKey.algorithm".
//
// The returned bool will be false if there is no such property.
func (this CryptoKey) Algorithm() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetCryptoKeyAlgorithm(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Usages returns the value of property "CryptoKey.usages".
//
// The returned bool will be false if there is no such property.
func (this CryptoKey) Usages() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetCryptoKeyUsages(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 RsaOtherPrimesInfo RsaOtherPrimesInfo [// RsaOtherPrimesInfo] [0x140007ee820 0x140007ee8c0 0x140007ee960] 0x14000780c78 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 JsonWebKey JsonWebKey [// JsonWebKey] [0x140007dbd60 0x140007dbe00 0x140007dbea0 0x140007dbf40 0x140007ee000 0x140007ee140 0x140007ee1e0 0x140007ee280 0x140007ee320 0x140007ee3c0 0x140007ee460 0x140007ee500 0x140007ee5a0 0x140007ee640 0x140007ee6e0 0x140007ee780 0x140007eea00 0x140007eeaa0 0x140007ee0a0] 0x1400081fe60 {0 0}} in the application heap.
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

// Encrypt calls the method "SubtleCrypto.encrypt".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) Encrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoEncrypt(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// EncryptFunc returns the method "SubtleCrypto.encrypt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) EncryptFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoEncryptFunc(
			this.Ref(),
		),
	)
}

// Decrypt calls the method "SubtleCrypto.decrypt".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) Decrypt(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoDecrypt(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// DecryptFunc returns the method "SubtleCrypto.decrypt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) DecryptFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDecryptFunc(
			this.Ref(),
		),
	)
}

// Sign calls the method "SubtleCrypto.sign".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) Sign(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoSign(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		key.Ref(),
		data.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// SignFunc returns the method "SubtleCrypto.sign".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) SignFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoSignFunc(
			this.Ref(),
		),
	)
}

// Verify calls the method "SubtleCrypto.verify".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) Verify(algorithm AlgorithmIdentifier, key CryptoKey, signature BufferSource, data BufferSource) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoVerify(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		key.Ref(),
		signature.Ref(),
		data.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// VerifyFunc returns the method "SubtleCrypto.verify".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) VerifyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, key CryptoKey, signature BufferSource, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoVerifyFunc(
			this.Ref(),
		),
	)
}

// Digest calls the method "SubtleCrypto.digest".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) Digest(algorithm AlgorithmIdentifier, data BufferSource) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoDigest(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		data.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// DigestFunc returns the method "SubtleCrypto.digest".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) DigestFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, data BufferSource) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDigestFunc(
			this.Ref(),
		),
	)
}

// GenerateKey calls the method "SubtleCrypto.generateKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) GenerateKey(algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoGenerateKey(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// GenerateKeyFunc returns the method "SubtleCrypto.generateKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) GenerateKeyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoGenerateKeyFunc(
			this.Ref(),
		),
	)
}

// DeriveKey calls the method "SubtleCrypto.deriveKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) DeriveKey(algorithm AlgorithmIdentifier, baseKey CryptoKey, derivedKeyType AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoDeriveKey(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		baseKey.Ref(),
		derivedKeyType.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// DeriveKeyFunc returns the method "SubtleCrypto.deriveKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) DeriveKeyFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, baseKey CryptoKey, derivedKeyType AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDeriveKeyFunc(
			this.Ref(),
		),
	)
}

// DeriveBits calls the method "SubtleCrypto.deriveBits".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) DeriveBits(algorithm AlgorithmIdentifier, baseKey CryptoKey, length uint32) (js.Promise[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoDeriveBits(
		this.Ref(), js.Pointer(&_ok),
		algorithm.Ref(),
		baseKey.Ref(),
		uint32(length),
	)

	return js.Promise[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// DeriveBitsFunc returns the method "SubtleCrypto.deriveBits".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) DeriveBitsFunc() (fn js.Func[func(algorithm AlgorithmIdentifier, baseKey CryptoKey, length uint32) js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.SubtleCryptoDeriveBitsFunc(
			this.Ref(),
		),
	)
}

// ImportKey calls the method "SubtleCrypto.importKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) ImportKey(format KeyFormat, keyData OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey, algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (js.Promise[CryptoKey], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoImportKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(format),
		keyData.Ref(),
		algorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return js.Promise[CryptoKey]{}.FromRef(_ret), _ok
}

// ImportKeyFunc returns the method "SubtleCrypto.importKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) ImportKeyFunc() (fn js.Func[func(format KeyFormat, keyData OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_JsonWebKey, algorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[CryptoKey]]) {
	return fn.FromRef(
		bindings.SubtleCryptoImportKeyFunc(
			this.Ref(),
		),
	)
}

// ExportKey calls the method "SubtleCrypto.exportKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) ExportKey(format KeyFormat, key CryptoKey) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoExportKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(format),
		key.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// ExportKeyFunc returns the method "SubtleCrypto.exportKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) ExportKeyFunc() (fn js.Func[func(format KeyFormat, key CryptoKey) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoExportKeyFunc(
			this.Ref(),
		),
	)
}

// WrapKey calls the method "SubtleCrypto.wrapKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) WrapKey(format KeyFormat, key CryptoKey, wrappingKey CryptoKey, wrapAlgorithm AlgorithmIdentifier) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoWrapKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(format),
		key.Ref(),
		wrappingKey.Ref(),
		wrapAlgorithm.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// WrapKeyFunc returns the method "SubtleCrypto.wrapKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) WrapKeyFunc() (fn js.Func[func(format KeyFormat, key CryptoKey, wrappingKey CryptoKey, wrapAlgorithm AlgorithmIdentifier) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SubtleCryptoWrapKeyFunc(
			this.Ref(),
		),
	)
}

// UnwrapKey calls the method "SubtleCrypto.unwrapKey".
//
// The returned bool will be false if there is no such method.
func (this SubtleCrypto) UnwrapKey(format KeyFormat, wrappedKey BufferSource, unwrappingKey CryptoKey, unwrapAlgorithm AlgorithmIdentifier, unwrappedKeyAlgorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) (js.Promise[CryptoKey], bool) {
	var _ok bool
	_ret := bindings.CallSubtleCryptoUnwrapKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(format),
		wrappedKey.Ref(),
		unwrappingKey.Ref(),
		unwrapAlgorithm.Ref(),
		unwrappedKeyAlgorithm.Ref(),
		js.Bool(bool(extractable)),
		keyUsages.Ref(),
	)

	return js.Promise[CryptoKey]{}.FromRef(_ret), _ok
}

// UnwrapKeyFunc returns the method "SubtleCrypto.unwrapKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SubtleCrypto) UnwrapKeyFunc() (fn js.Func[func(format KeyFormat, wrappedKey BufferSource, unwrappingKey CryptoKey, unwrapAlgorithm AlgorithmIdentifier, unwrappedKeyAlgorithm AlgorithmIdentifier, extractable bool, keyUsages js.Array[KeyUsage]) js.Promise[CryptoKey]]) {
	return fn.FromRef(
		bindings.SubtleCryptoUnwrapKeyFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this Crypto) Subtle() (SubtleCrypto, bool) {
	var _ok bool
	_ret := bindings.GetCryptoSubtle(
		this.Ref(), js.Pointer(&_ok),
	)
	return SubtleCrypto{}.FromRef(_ret), _ok
}

// GetRandomValues calls the method "Crypto.getRandomValues".
//
// The returned bool will be false if there is no such method.
func (this Crypto) GetRandomValues(array ArrayBufferView) (ArrayBufferView, bool) {
	var _ok bool
	_ret := bindings.CallCryptoGetRandomValues(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	return ArrayBufferView{}.FromRef(_ret), _ok
}

// GetRandomValuesFunc returns the method "Crypto.getRandomValues".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Crypto) GetRandomValuesFunc() (fn js.Func[func(array ArrayBufferView) ArrayBufferView]) {
	return fn.FromRef(
		bindings.CryptoGetRandomValuesFunc(
			this.Ref(),
		),
	)
}

// RandomUUID calls the method "Crypto.randomUUID".
//
// The returned bool will be false if there is no such method.
func (this Crypto) RandomUUID() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCryptoRandomUUID(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// RandomUUIDFunc returns the method "Crypto.randomUUID".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Crypto) RandomUUIDFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CryptoRandomUUIDFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 MemoryAttributionContainer MemoryAttributionContainer [// MemoryAttributionContainer] [0x140007eee60 0x140007eef00] 0x14000780db0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MemoryAttribution MemoryAttribution [// MemoryAttribution] [0x140007eedc0 0x140007eefa0 0x140007ef040] 0x14000780d80 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MemoryBreakdownEntry MemoryBreakdownEntry [// MemoryBreakdownEntry] [0x140007eec80 0x140007ef0e0 0x140007ef180 0x140007eed20] 0x14000780d50 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MemoryMeasurement MemoryMeasurement [// MemoryMeasurement] [0x140007eeb40 0x140007ef220 0x140007eebe0] 0x14000780d20 {0 0}} in the application heap.
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
	StartTime DOMHighResTimeStamp

	FFI_USE_StartTime bool // for StartTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceMarkOptions with all fields set.
func (p PerformanceMarkOptions) FromRef(ref js.Ref) PerformanceMarkOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PerformanceMarkOptions PerformanceMarkOptions [// PerformanceMarkOptions] [0x140007ef2c0 0x140007ef360 0x140007ef400] 0x14000780e28 {0 0}} in the application heap.
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

func NewPerformanceMark(markName js.String, markOptions PerformanceMarkOptions) PerformanceMark {
	return PerformanceMark{}.FromRef(
		bindings.NewPerformanceMarkByPerformanceMark(
			markName.Ref(),
			js.Pointer(&markOptions)),
	)
}

func NewPerformanceMarkByPerformanceMark1(markName js.String) PerformanceMark {
	return PerformanceMark{}.FromRef(
		bindings.NewPerformanceMarkByPerformanceMark1(
			markName.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceMark) Detail() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceMarkDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this PerformanceMeasure) Detail() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceMeasureDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
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

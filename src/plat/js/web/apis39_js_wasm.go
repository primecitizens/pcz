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

type IDBRequestReadyState uint32

const (
	_ IDBRequestReadyState = iota

	IDBRequestReadyState_PENDING
	IDBRequestReadyState_DONE
)

func (IDBRequestReadyState) FromRef(str js.Ref) IDBRequestReadyState {
	return IDBRequestReadyState(bindings.ConstOfIDBRequestReadyState(str))
}

func (x IDBRequestReadyState) String() (string, bool) {
	switch x {
	case IDBRequestReadyState_PENDING:
		return "pending", true
	case IDBRequestReadyState_DONE:
		return "done", true
	default:
		return "", false
	}
}

type IDBRequest struct {
	EventTarget
}

func (this IDBRequest) Once() IDBRequest {
	this.Ref().Once()
	return this
}

func (this IDBRequest) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this IDBRequest) FromRef(ref js.Ref) IDBRequest {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this IDBRequest) Free() {
	this.Ref().Free()
}

// Result returns the value of property "IDBRequest.result".
//
// The returned bool will be false if there is no such property.
func (this IDBRequest) Result() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBRequestResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Error returns the value of property "IDBRequest.error".
//
// The returned bool will be false if there is no such property.
func (this IDBRequest) Error() (DOMException, bool) {
	var _ok bool
	_ret := bindings.GetIDBRequestError(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMException{}.FromRef(_ret), _ok
}

// Source returns the value of property "IDBRequest.source".
//
// The returned bool will be false if there is no such property.
func (this IDBRequest) Source() (OneOf_IDBObjectStore_IDBIndex_IDBCursor, bool) {
	var _ok bool
	_ret := bindings.GetIDBRequestSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_IDBObjectStore_IDBIndex_IDBCursor{}.FromRef(_ret), _ok
}

// Transaction returns the value of property "IDBRequest.transaction".
//
// The returned bool will be false if there is no such property.
func (this IDBRequest) Transaction() (IDBTransaction, bool) {
	var _ok bool
	_ret := bindings.GetIDBRequestTransaction(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBTransaction{}.FromRef(_ret), _ok
}

// ReadyState returns the value of property "IDBRequest.readyState".
//
// The returned bool will be false if there is no such property.
func (this IDBRequest) ReadyState() (IDBRequestReadyState, bool) {
	var _ok bool
	_ret := bindings.GetIDBRequestReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBRequestReadyState(_ret), _ok
}

type OneOf_IDBObjectStore_IDBIndex struct {
	ref js.Ref
}

func (x OneOf_IDBObjectStore_IDBIndex) Ref() js.Ref {
	return x.ref
}

func (x OneOf_IDBObjectStore_IDBIndex) Free() {
	x.ref.Free()
}

func (x OneOf_IDBObjectStore_IDBIndex) FromRef(ref js.Ref) OneOf_IDBObjectStore_IDBIndex {
	return OneOf_IDBObjectStore_IDBIndex{
		ref: ref,
	}
}

func (x OneOf_IDBObjectStore_IDBIndex) IDBObjectStore() IDBObjectStore {
	return IDBObjectStore{}.FromRef(x.ref)
}

func (x OneOf_IDBObjectStore_IDBIndex) IDBIndex() IDBIndex {
	return IDBIndex{}.FromRef(x.ref)
}

type IDBCursor struct {
	ref js.Ref
}

func (this IDBCursor) Once() IDBCursor {
	this.Ref().Once()
	return this
}

func (this IDBCursor) Ref() js.Ref {
	return this.ref
}

func (this IDBCursor) FromRef(ref js.Ref) IDBCursor {
	this.ref = ref
	return this
}

func (this IDBCursor) Free() {
	this.Ref().Free()
}

// Source returns the value of property "IDBCursor.source".
//
// The returned bool will be false if there is no such property.
func (this IDBCursor) Source() (OneOf_IDBObjectStore_IDBIndex, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_IDBObjectStore_IDBIndex{}.FromRef(_ret), _ok
}

// Direction returns the value of property "IDBCursor.direction".
//
// The returned bool will be false if there is no such property.
func (this IDBCursor) Direction() (IDBCursorDirection, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBCursorDirection(_ret), _ok
}

// Key returns the value of property "IDBCursor.key".
//
// The returned bool will be false if there is no such property.
func (this IDBCursor) Key() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// PrimaryKey returns the value of property "IDBCursor.primaryKey".
//
// The returned bool will be false if there is no such property.
func (this IDBCursor) PrimaryKey() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorPrimaryKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Request returns the value of property "IDBCursor.request".
//
// The returned bool will be false if there is no such property.
func (this IDBCursor) Request() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBRequest{}.FromRef(_ret), _ok
}

// Advance calls the method "IDBCursor.advance".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) Advance(count uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorAdvance(
		this.Ref(), js.Pointer(&_ok),
		uint32(count),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AdvanceFunc returns the method "IDBCursor.advance".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) AdvanceFunc() (fn js.Func[func(count uint32)]) {
	return fn.FromRef(
		bindings.IDBCursorAdvanceFunc(
			this.Ref(),
		),
	)
}

// Continue calls the method "IDBCursor.continue".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) Continue(key js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorContinue(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ContinueFunc returns the method "IDBCursor.continue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) ContinueFunc() (fn js.Func[func(key js.Any)]) {
	return fn.FromRef(
		bindings.IDBCursorContinueFunc(
			this.Ref(),
		),
	)
}

// Continue1 calls the method "IDBCursor.continue".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) Continue1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorContinue1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Continue1Func returns the method "IDBCursor.continue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) Continue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBCursorContinue1Func(
			this.Ref(),
		),
	)
}

// ContinuePrimaryKey calls the method "IDBCursor.continuePrimaryKey".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) ContinuePrimaryKey(key js.Any, primaryKey js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorContinuePrimaryKey(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		primaryKey.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ContinuePrimaryKeyFunc returns the method "IDBCursor.continuePrimaryKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) ContinuePrimaryKeyFunc() (fn js.Func[func(key js.Any, primaryKey js.Any)]) {
	return fn.FromRef(
		bindings.IDBCursorContinuePrimaryKeyFunc(
			this.Ref(),
		),
	)
}

// Update calls the method "IDBCursor.update".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) Update(value js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorUpdate(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// UpdateFunc returns the method "IDBCursor.update".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) UpdateFunc() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBCursorUpdateFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "IDBCursor.delete".
//
// The returned bool will be false if there is no such method.
func (this IDBCursor) Delete() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBCursorDelete(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "IDBCursor.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBCursor) DeleteFunc() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBCursorDeleteFunc(
			this.Ref(),
		),
	)
}

type IDBCursorWithValue struct {
	IDBCursor
}

func (this IDBCursorWithValue) Once() IDBCursorWithValue {
	this.Ref().Once()
	return this
}

func (this IDBCursorWithValue) Ref() js.Ref {
	return this.IDBCursor.Ref()
}

func (this IDBCursorWithValue) FromRef(ref js.Ref) IDBCursorWithValue {
	this.IDBCursor = this.IDBCursor.FromRef(ref)
	return this
}

func (this IDBCursorWithValue) Free() {
	this.Ref().Free()
}

// Value returns the value of property "IDBCursorWithValue.value".
//
// The returned bool will be false if there is no such property.
func (this IDBCursorWithValue) Value() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBCursorWithValueValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

type IDBKeyRange struct {
	ref js.Ref
}

func (this IDBKeyRange) Once() IDBKeyRange {
	this.Ref().Once()
	return this
}

func (this IDBKeyRange) Ref() js.Ref {
	return this.ref
}

func (this IDBKeyRange) FromRef(ref js.Ref) IDBKeyRange {
	this.ref = ref
	return this
}

func (this IDBKeyRange) Free() {
	this.Ref().Free()
}

// Lower returns the value of property "IDBKeyRange.lower".
//
// The returned bool will be false if there is no such property.
func (this IDBKeyRange) Lower() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBKeyRangeLower(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Upper returns the value of property "IDBKeyRange.upper".
//
// The returned bool will be false if there is no such property.
func (this IDBKeyRange) Upper() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBKeyRangeUpper(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// LowerOpen returns the value of property "IDBKeyRange.lowerOpen".
//
// The returned bool will be false if there is no such property.
func (this IDBKeyRange) LowerOpen() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIDBKeyRangeLowerOpen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// UpperOpen returns the value of property "IDBKeyRange.upperOpen".
//
// The returned bool will be false if there is no such property.
func (this IDBKeyRange) UpperOpen() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIDBKeyRangeUpperOpen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Only calls the staticmethod "IDBKeyRange.only".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) Only(value js.Any) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeOnly(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// OnlyFunc returns the staticmethod "IDBKeyRange.only".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) OnlyFunc() (fn js.Func[func(value js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeOnlyFunc(
			this.Ref(),
		),
	)
}

// LowerBound calls the staticmethod "IDBKeyRange.lowerBound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) LowerBound(lower js.Any, open bool) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeLowerBound(
		this.Ref(), js.Pointer(&_ok),
		lower.Ref(),
		js.Bool(bool(open)),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// LowerBoundFunc returns the staticmethod "IDBKeyRange.lowerBound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) LowerBoundFunc() (fn js.Func[func(lower js.Any, open bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeLowerBoundFunc(
			this.Ref(),
		),
	)
}

// LowerBound1 calls the staticmethod "IDBKeyRange.lowerBound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) LowerBound1(lower js.Any) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeLowerBound1(
		this.Ref(), js.Pointer(&_ok),
		lower.Ref(),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// LowerBound1Func returns the staticmethod "IDBKeyRange.lowerBound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) LowerBound1Func() (fn js.Func[func(lower js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeLowerBound1Func(
			this.Ref(),
		),
	)
}

// UpperBound calls the staticmethod "IDBKeyRange.upperBound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) UpperBound(upper js.Any, open bool) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeUpperBound(
		this.Ref(), js.Pointer(&_ok),
		upper.Ref(),
		js.Bool(bool(open)),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// UpperBoundFunc returns the staticmethod "IDBKeyRange.upperBound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) UpperBoundFunc() (fn js.Func[func(upper js.Any, open bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeUpperBoundFunc(
			this.Ref(),
		),
	)
}

// UpperBound1 calls the staticmethod "IDBKeyRange.upperBound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) UpperBound1(upper js.Any) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeUpperBound1(
		this.Ref(), js.Pointer(&_ok),
		upper.Ref(),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// UpperBound1Func returns the staticmethod "IDBKeyRange.upperBound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) UpperBound1Func() (fn js.Func[func(upper js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeUpperBound1Func(
			this.Ref(),
		),
	)
}

// Bound calls the staticmethod "IDBKeyRange.bound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) Bound(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeBound(
		this.Ref(), js.Pointer(&_ok),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
		js.Bool(bool(upperOpen)),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// BoundFunc returns the staticmethod "IDBKeyRange.bound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) BoundFunc() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBoundFunc(
			this.Ref(),
		),
	)
}

// Bound1 calls the staticmethod "IDBKeyRange.bound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) Bound1(lower js.Any, upper js.Any, lowerOpen bool) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeBound1(
		this.Ref(), js.Pointer(&_ok),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// Bound1Func returns the staticmethod "IDBKeyRange.bound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) Bound1Func() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBound1Func(
			this.Ref(),
		),
	)
}

// Bound2 calls the staticmethod "IDBKeyRange.bound".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) Bound2(lower js.Any, upper js.Any) (IDBKeyRange, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeBound2(
		this.Ref(), js.Pointer(&_ok),
		lower.Ref(),
		upper.Ref(),
	)

	return IDBKeyRange{}.FromRef(_ret), _ok
}

// Bound2Func returns the staticmethod "IDBKeyRange.bound".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) Bound2Func() (fn js.Func[func(lower js.Any, upper js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBound2Func(
			this.Ref(),
		),
	)
}

// Includes calls the method "IDBKeyRange.includes".
//
// The returned bool will be false if there is no such method.
func (this IDBKeyRange) Includes(key js.Any) (bool, bool) {
	var _ok bool
	_ret := bindings.CallIDBKeyRangeIncludes(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	return _ret == js.True, _ok
}

// IncludesFunc returns the method "IDBKeyRange.includes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBKeyRange) IncludesFunc() (fn js.Func[func(key js.Any) bool]) {
	return fn.FromRef(
		bindings.IDBKeyRangeIncludesFunc(
			this.Ref(),
		),
	)
}

type IDBVersionChangeEventInit struct {
	// OldVersion is "IDBVersionChangeEventInit.oldVersion"
	//
	// Optional, defaults to 0.
	OldVersion uint64
	// NewVersion is "IDBVersionChangeEventInit.newVersion"
	//
	// Optional, defaults to null.
	NewVersion uint64
	// Bubbles is "IDBVersionChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "IDBVersionChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "IDBVersionChangeEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_OldVersion bool // for OldVersion.
	FFI_USE_NewVersion bool // for NewVersion.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBVersionChangeEventInit with all fields set.
func (p IDBVersionChangeEventInit) FromRef(ref js.Ref) IDBVersionChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IDBVersionChangeEventInit IDBVersionChangeEventInit [// IDBVersionChangeEventInit] [0x1400084f0e0 0x1400084f220 0x1400084f360 0x1400084f4a0 0x1400084f5e0 0x1400084f180 0x1400084f2c0 0x1400084f400 0x1400084f540 0x1400084f680] 0x14000aa2438 {0 0}} in the application heap.
func (p IDBVersionChangeEventInit) New() js.Ref {
	return bindings.IDBVersionChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IDBVersionChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.IDBVersionChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IDBVersionChangeEventInit) Update(ref js.Ref) {
	bindings.IDBVersionChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewIDBVersionChangeEvent(typ js.String, eventInitDict IDBVersionChangeEventInit) IDBVersionChangeEvent {
	return IDBVersionChangeEvent{}.FromRef(
		bindings.NewIDBVersionChangeEventByIDBVersionChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewIDBVersionChangeEventByIDBVersionChangeEvent1(typ js.String) IDBVersionChangeEvent {
	return IDBVersionChangeEvent{}.FromRef(
		bindings.NewIDBVersionChangeEventByIDBVersionChangeEvent1(
			typ.Ref()),
	)
}

type IDBVersionChangeEvent struct {
	Event
}

func (this IDBVersionChangeEvent) Once() IDBVersionChangeEvent {
	this.Ref().Once()
	return this
}

func (this IDBVersionChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this IDBVersionChangeEvent) FromRef(ref js.Ref) IDBVersionChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this IDBVersionChangeEvent) Free() {
	this.Ref().Free()
}

// OldVersion returns the value of property "IDBVersionChangeEvent.oldVersion".
//
// The returned bool will be false if there is no such property.
func (this IDBVersionChangeEvent) OldVersion() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetIDBVersionChangeEventOldVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// NewVersion returns the value of property "IDBVersionChangeEvent.newVersion".
//
// The returned bool will be false if there is no such property.
func (this IDBVersionChangeEvent) NewVersion() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetIDBVersionChangeEventNewVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

type IdentityCredential struct {
	Credential
}

func (this IdentityCredential) Once() IdentityCredential {
	this.Ref().Once()
	return this
}

func (this IdentityCredential) Ref() js.Ref {
	return this.Credential.Ref()
}

func (this IdentityCredential) FromRef(ref js.Ref) IdentityCredential {
	this.Credential = this.Credential.FromRef(ref)
	return this
}

func (this IdentityCredential) Free() {
	this.Ref().Free()
}

// Token returns the value of property "IdentityCredential.token".
//
// The returned bool will be false if there is no such property.
func (this IdentityCredential) Token() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetIdentityCredentialToken(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type IdentityUserInfo struct {
	// Email is "IdentityUserInfo.email"
	//
	// Optional
	Email js.String
	// Name is "IdentityUserInfo.name"
	//
	// Optional
	Name js.String
	// GivenName is "IdentityUserInfo.givenName"
	//
	// Optional
	GivenName js.String
	// Picture is "IdentityUserInfo.picture"
	//
	// Optional
	Picture js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityUserInfo with all fields set.
func (p IdentityUserInfo) FromRef(ref js.Ref) IdentityUserInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityUserInfo IdentityUserInfo [// IdentityUserInfo] [0x1400084f7c0 0x1400084f860 0x1400084f900 0x1400084f9a0] 0x14000aa31a0 {0 0}} in the application heap.
func (p IdentityUserInfo) New() js.Ref {
	return bindings.IdentityUserInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityUserInfo) UpdateFrom(ref js.Ref) {
	bindings.IdentityUserInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityUserInfo) Update(ref js.Ref) {
	bindings.IdentityUserInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProvider struct {
	ref js.Ref
}

func (this IdentityProvider) Once() IdentityProvider {
	this.Ref().Once()
	return this
}

func (this IdentityProvider) Ref() js.Ref {
	return this.ref
}

func (this IdentityProvider) FromRef(ref js.Ref) IdentityProvider {
	this.ref = ref
	return this
}

func (this IdentityProvider) Free() {
	this.Ref().Free()
}

// GetUserInfo calls the staticmethod "IdentityProvider.getUserInfo".
//
// The returned bool will be false if there is no such method.
func (this IdentityProvider) GetUserInfo(config IdentityProviderConfig) (js.Promise[js.Array[IdentityUserInfo]], bool) {
	var _ok bool
	_ret := bindings.CallIdentityProviderGetUserInfo(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	return js.Promise[js.Array[IdentityUserInfo]]{}.FromRef(_ret), _ok
}

// GetUserInfoFunc returns the staticmethod "IdentityProvider.getUserInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IdentityProvider) GetUserInfoFunc() (fn js.Func[func(config IdentityProviderConfig) js.Promise[js.Array[IdentityUserInfo]]]) {
	return fn.FromRef(
		bindings.IdentityProviderGetUserInfoFunc(
			this.Ref(),
		),
	)
}

type IdentityProviderIcon struct {
	// Url is "IdentityProviderIcon.url"
	//
	// Required
	Url js.String
	// Size is "IdentityProviderIcon.size"
	//
	// Optional
	Size uint32

	FFI_USE_Size bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderIcon with all fields set.
func (p IdentityProviderIcon) FromRef(ref js.Ref) IdentityProviderIcon {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderIcon IdentityProviderIcon [// IdentityProviderIcon] [0x1400084fd60 0x1400084fe00 0x1400084fea0] 0x14000aa3b78 {0 0}} in the application heap.
func (p IdentityProviderIcon) New() js.Ref {
	return bindings.IdentityProviderIconJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderIcon) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderIconJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderIcon) Update(ref js.Ref) {
	bindings.IdentityProviderIconJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderBranding struct {
	// BackgroundColor is "IdentityProviderBranding.background_color"
	//
	// Optional
	BackgroundColor js.String
	// Color is "IdentityProviderBranding.color"
	//
	// Optional
	Color js.String
	// Icons is "IdentityProviderBranding.icons"
	//
	// Optional
	Icons js.Array[IdentityProviderIcon]
	// Name is "IdentityProviderBranding.name"
	//
	// Optional
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderBranding with all fields set.
func (p IdentityProviderBranding) FromRef(ref js.Ref) IdentityProviderBranding {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderBranding IdentityProviderBranding [// IdentityProviderBranding] [0x1400084fc20 0x1400084fcc0 0x1400084ff40 0x14000862000] 0x14000aa3a10 {0 0}} in the application heap.
func (p IdentityProviderBranding) New() js.Ref {
	return bindings.IdentityProviderBrandingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderBranding) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderBrandingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderBranding) Update(ref js.Ref) {
	bindings.IdentityProviderBrandingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderAPIConfig struct {
	// AccountsEndpoint is "IdentityProviderAPIConfig.accounts_endpoint"
	//
	// Required
	AccountsEndpoint js.String
	// ClientMetadataEndpoint is "IdentityProviderAPIConfig.client_metadata_endpoint"
	//
	// Required
	ClientMetadataEndpoint js.String
	// IdAssertionEndpoint is "IdentityProviderAPIConfig.id_assertion_endpoint"
	//
	// Required
	IdAssertionEndpoint js.String
	// Branding is "IdentityProviderAPIConfig.branding"
	//
	// Optional
	Branding IdentityProviderBranding

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderAPIConfig with all fields set.
func (p IdentityProviderAPIConfig) FromRef(ref js.Ref) IdentityProviderAPIConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderAPIConfig IdentityProviderAPIConfig [// IdentityProviderAPIConfig] [0x1400084fa40 0x1400084fae0 0x1400084fb80 0x140008620a0] 0x14000aa39e0 {0 0}} in the application heap.
func (p IdentityProviderAPIConfig) New() js.Ref {
	return bindings.IdentityProviderAPIConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderAPIConfig) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAPIConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderAPIConfig) Update(ref js.Ref) {
	bindings.IdentityProviderAPIConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderAccount struct {
	// Id is "IdentityProviderAccount.id"
	//
	// Required
	Id js.String
	// Name is "IdentityProviderAccount.name"
	//
	// Required
	Name js.String
	// Email is "IdentityProviderAccount.email"
	//
	// Required
	Email js.String
	// GivenName is "IdentityProviderAccount.given_name"
	//
	// Optional
	GivenName js.String
	// Picture is "IdentityProviderAccount.picture"
	//
	// Optional
	Picture js.String
	// ApprovedClients is "IdentityProviderAccount.approved_clients"
	//
	// Optional
	ApprovedClients js.Array[js.String]
	// LoginHints is "IdentityProviderAccount.login_hints"
	//
	// Optional
	LoginHints js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderAccount with all fields set.
func (p IdentityProviderAccount) FromRef(ref js.Ref) IdentityProviderAccount {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderAccount IdentityProviderAccount [// IdentityProviderAccount] [0x14000862140 0x140008621e0 0x14000862280 0x14000862320 0x140008623c0 0x14000862460 0x14000862500] 0x14000aa3f50 {0 0}} in the application heap.
func (p IdentityProviderAccount) New() js.Ref {
	return bindings.IdentityProviderAccountJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderAccount) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAccountJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderAccount) Update(ref js.Ref) {
	bindings.IdentityProviderAccountJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderAccountList struct {
	// Accounts is "IdentityProviderAccountList.accounts"
	//
	// Optional
	Accounts js.Array[IdentityProviderAccount]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderAccountList with all fields set.
func (p IdentityProviderAccountList) FromRef(ref js.Ref) IdentityProviderAccountList {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderAccountList IdentityProviderAccountList [// IdentityProviderAccountList] [0x140008625a0] 0x14000aa3fe0 {0 0}} in the application heap.
func (p IdentityProviderAccountList) New() js.Ref {
	return bindings.IdentityProviderAccountListJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderAccountList) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAccountListJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderAccountList) Update(ref js.Ref) {
	bindings.IdentityProviderAccountListJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderClientMetadata struct {
	// PrivacyPolicyUrl is "IdentityProviderClientMetadata.privacy_policy_url"
	//
	// Optional
	PrivacyPolicyUrl js.String
	// TermsOfServiceUrl is "IdentityProviderClientMetadata.terms_of_service_url"
	//
	// Optional
	TermsOfServiceUrl js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderClientMetadata with all fields set.
func (p IdentityProviderClientMetadata) FromRef(ref js.Ref) IdentityProviderClientMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderClientMetadata IdentityProviderClientMetadata [// IdentityProviderClientMetadata] [0x14000862640 0x140008626e0] 0x14001c3a5d0 {0 0}} in the application heap.
func (p IdentityProviderClientMetadata) New() js.Ref {
	return bindings.IdentityProviderClientMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderClientMetadata) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderClientMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderClientMetadata) Update(ref js.Ref) {
	bindings.IdentityProviderClientMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderToken struct {
	// Token is "IdentityProviderToken.token"
	//
	// Required
	Token js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderToken with all fields set.
func (p IdentityProviderToken) FromRef(ref js.Ref) IdentityProviderToken {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderToken IdentityProviderToken [// IdentityProviderToken] [0x14000862780] 0x14001c3ae28 {0 0}} in the application heap.
func (p IdentityProviderToken) New() js.Ref {
	return bindings.IdentityProviderTokenJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderToken) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderTokenJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderToken) Update(ref js.Ref) {
	bindings.IdentityProviderTokenJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdentityProviderWellKnown struct {
	// ProviderUrls is "IdentityProviderWellKnown.provider_urls"
	//
	// Required
	ProviderUrls js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderWellKnown with all fields set.
func (p IdentityProviderWellKnown) FromRef(ref js.Ref) IdentityProviderWellKnown {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdentityProviderWellKnown IdentityProviderWellKnown [// IdentityProviderWellKnown] [0x14000862820] 0x14001c3ae88 {0 0}} in the application heap.
func (p IdentityProviderWellKnown) New() js.Ref {
	return bindings.IdentityProviderWellKnownJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdentityProviderWellKnown) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderWellKnownJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdentityProviderWellKnown) Update(ref js.Ref) {
	bindings.IdentityProviderWellKnownJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IdleOptions struct {
	// Threshold is "IdleOptions.threshold"
	//
	// Optional
	Threshold uint64
	// Signal is "IdleOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE_Threshold bool // for Threshold.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdleOptions with all fields set.
func (p IdleOptions) FromRef(ref js.Ref) IdleOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 IdleOptions IdleOptions [// IdleOptions] [0x140008628c0 0x14000862a00 0x14000862960] 0x14001c3bbc0 {0 0}} in the application heap.
func (p IdleOptions) New() js.Ref {
	return bindings.IdleOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IdleOptions) UpdateFrom(ref js.Ref) {
	bindings.IdleOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdleOptions) Update(ref js.Ref) {
	bindings.IdleOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type UserIdleState uint32

const (
	_ UserIdleState = iota

	UserIdleState_ACTIVE
	UserIdleState_IDLE
)

func (UserIdleState) FromRef(str js.Ref) UserIdleState {
	return UserIdleState(bindings.ConstOfUserIdleState(str))
}

func (x UserIdleState) String() (string, bool) {
	switch x {
	case UserIdleState_ACTIVE:
		return "active", true
	case UserIdleState_IDLE:
		return "idle", true
	default:
		return "", false
	}
}

type ScreenIdleState uint32

const (
	_ ScreenIdleState = iota

	ScreenIdleState_LOCKED
	ScreenIdleState_UNLOCKED
)

func (ScreenIdleState) FromRef(str js.Ref) ScreenIdleState {
	return ScreenIdleState(bindings.ConstOfScreenIdleState(str))
}

func (x ScreenIdleState) String() (string, bool) {
	switch x {
	case ScreenIdleState_LOCKED:
		return "locked", true
	case ScreenIdleState_UNLOCKED:
		return "unlocked", true
	default:
		return "", false
	}
}

type IdleDetector struct {
	EventTarget
}

func (this IdleDetector) Once() IdleDetector {
	this.Ref().Once()
	return this
}

func (this IdleDetector) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this IdleDetector) FromRef(ref js.Ref) IdleDetector {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this IdleDetector) Free() {
	this.Ref().Free()
}

// UserState returns the value of property "IdleDetector.userState".
//
// The returned bool will be false if there is no such property.
func (this IdleDetector) UserState() (UserIdleState, bool) {
	var _ok bool
	_ret := bindings.GetIdleDetectorUserState(
		this.Ref(), js.Pointer(&_ok),
	)
	return UserIdleState(_ret), _ok
}

// ScreenState returns the value of property "IdleDetector.screenState".
//
// The returned bool will be false if there is no such property.
func (this IdleDetector) ScreenState() (ScreenIdleState, bool) {
	var _ok bool
	_ret := bindings.GetIdleDetectorScreenState(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScreenIdleState(_ret), _ok
}

// RequestPermission calls the staticmethod "IdleDetector.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this IdleDetector) RequestPermission() (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallIdleDetectorRequestPermission(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// RequestPermissionFunc returns the staticmethod "IdleDetector.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IdleDetector) RequestPermissionFunc() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.IdleDetectorRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "IdleDetector.start".
//
// The returned bool will be false if there is no such method.
func (this IdleDetector) Start(options IdleOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallIdleDetectorStart(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// StartFunc returns the method "IdleDetector.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IdleDetector) StartFunc() (fn js.Func[func(options IdleOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.IdleDetectorStartFunc(
			this.Ref(),
		),
	)
}

// Start1 calls the method "IdleDetector.start".
//
// The returned bool will be false if there is no such method.
func (this IdleDetector) Start1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallIdleDetectorStart1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Start1Func returns the method "IdleDetector.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IdleDetector) Start1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.IdleDetectorStart1Func(
			this.Ref(),
		),
	)
}

type ImageBitmapRenderingContextSettings struct {
	// Alpha is "ImageBitmapRenderingContextSettings.alpha"
	//
	// Optional, defaults to true.
	Alpha bool

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageBitmapRenderingContextSettings with all fields set.
func (p ImageBitmapRenderingContextSettings) FromRef(ref js.Ref) ImageBitmapRenderingContextSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ImageBitmapRenderingContextSettings ImageBitmapRenderingContextSettings [// ImageBitmapRenderingContextSettings] [0x14000862b40 0x14000862be0] 0x14001c3bbf0 {0 0}} in the application heap.
func (p ImageBitmapRenderingContextSettings) New() js.Ref {
	return bindings.ImageBitmapRenderingContextSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageBitmapRenderingContextSettings) UpdateFrom(ref js.Ref) {
	bindings.ImageBitmapRenderingContextSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageBitmapRenderingContextSettings) Update(ref js.Ref) {
	bindings.ImageBitmapRenderingContextSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream) ReadableStream() ReadableStream {
	return ReadableStream{}.FromRef(x.ref)
}

type ImageBufferSource = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_ReadableStream

type PhotoSettings struct {
	// FillLightMode is "PhotoSettings.fillLightMode"
	//
	// Optional
	FillLightMode FillLightMode
	// ImageHeight is "PhotoSettings.imageHeight"
	//
	// Optional
	ImageHeight float64
	// ImageWidth is "PhotoSettings.imageWidth"
	//
	// Optional
	ImageWidth float64
	// RedEyeReduction is "PhotoSettings.redEyeReduction"
	//
	// Optional
	RedEyeReduction bool

	FFI_USE_ImageHeight     bool // for ImageHeight.
	FFI_USE_ImageWidth      bool // for ImageWidth.
	FFI_USE_RedEyeReduction bool // for RedEyeReduction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PhotoSettings with all fields set.
func (p PhotoSettings) FromRef(ref js.Ref) PhotoSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PhotoSettings PhotoSettings [// PhotoSettings] [0x14000862c80 0x14000862d20 0x14000862e60 0x14000862fa0 0x14000862dc0 0x14000862f00 0x14000863040] 0x14001c3bd88 {0 0}} in the application heap.
func (p PhotoSettings) New() js.Ref {
	return bindings.PhotoSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PhotoSettings) UpdateFrom(ref js.Ref) {
	bindings.PhotoSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PhotoSettings) Update(ref js.Ref) {
	bindings.PhotoSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RedEyeReduction uint32

const (
	_ RedEyeReduction = iota

	RedEyeReduction_NEVER
	RedEyeReduction_ALWAYS
	RedEyeReduction_CONTROLLABLE
)

func (RedEyeReduction) FromRef(str js.Ref) RedEyeReduction {
	return RedEyeReduction(bindings.ConstOfRedEyeReduction(str))
}

func (x RedEyeReduction) String() (string, bool) {
	switch x {
	case RedEyeReduction_NEVER:
		return "never", true
	case RedEyeReduction_ALWAYS:
		return "always", true
	case RedEyeReduction_CONTROLLABLE:
		return "controllable", true
	default:
		return "", false
	}
}

type PhotoCapabilities struct {
	// RedEyeReduction is "PhotoCapabilities.redEyeReduction"
	//
	// Optional
	RedEyeReduction RedEyeReduction
	// ImageHeight is "PhotoCapabilities.imageHeight"
	//
	// Optional
	ImageHeight MediaSettingsRange
	// ImageWidth is "PhotoCapabilities.imageWidth"
	//
	// Optional
	ImageWidth MediaSettingsRange
	// FillLightMode is "PhotoCapabilities.fillLightMode"
	//
	// Optional
	FillLightMode js.Array[FillLightMode]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PhotoCapabilities with all fields set.
func (p PhotoCapabilities) FromRef(ref js.Ref) PhotoCapabilities {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PhotoCapabilities PhotoCapabilities [// PhotoCapabilities] [0x140008630e0 0x14000863180 0x14000863220 0x140008632c0] 0x14001c3bf08 {0 0}} in the application heap.
func (p PhotoCapabilities) New() js.Ref {
	return bindings.PhotoCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PhotoCapabilities) UpdateFrom(ref js.Ref) {
	bindings.PhotoCapabilitiesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PhotoCapabilities) Update(ref js.Ref) {
	bindings.PhotoCapabilitiesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewImageCapture(videoTrack MediaStreamTrack) ImageCapture {
	return ImageCapture{}.FromRef(
		bindings.NewImageCaptureByImageCapture(
			videoTrack.Ref()),
	)
}

type ImageCapture struct {
	ref js.Ref
}

func (this ImageCapture) Once() ImageCapture {
	this.Ref().Once()
	return this
}

func (this ImageCapture) Ref() js.Ref {
	return this.ref
}

func (this ImageCapture) FromRef(ref js.Ref) ImageCapture {
	this.ref = ref
	return this
}

func (this ImageCapture) Free() {
	this.Ref().Free()
}

// Track returns the value of property "ImageCapture.track".
//
// The returned bool will be false if there is no such property.
func (this ImageCapture) Track() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.GetImageCaptureTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// TakePhoto calls the method "ImageCapture.takePhoto".
//
// The returned bool will be false if there is no such method.
func (this ImageCapture) TakePhoto(photoSettings PhotoSettings) (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallImageCaptureTakePhoto(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&photoSettings),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// TakePhotoFunc returns the method "ImageCapture.takePhoto".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageCapture) TakePhotoFunc() (fn js.Func[func(photoSettings PhotoSettings) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ImageCaptureTakePhotoFunc(
			this.Ref(),
		),
	)
}

// TakePhoto1 calls the method "ImageCapture.takePhoto".
//
// The returned bool will be false if there is no such method.
func (this ImageCapture) TakePhoto1() (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallImageCaptureTakePhoto1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// TakePhoto1Func returns the method "ImageCapture.takePhoto".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageCapture) TakePhoto1Func() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ImageCaptureTakePhoto1Func(
			this.Ref(),
		),
	)
}

// GetPhotoCapabilities calls the method "ImageCapture.getPhotoCapabilities".
//
// The returned bool will be false if there is no such method.
func (this ImageCapture) GetPhotoCapabilities() (js.Promise[PhotoCapabilities], bool) {
	var _ok bool
	_ret := bindings.CallImageCaptureGetPhotoCapabilities(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PhotoCapabilities]{}.FromRef(_ret), _ok
}

// GetPhotoCapabilitiesFunc returns the method "ImageCapture.getPhotoCapabilities".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageCapture) GetPhotoCapabilitiesFunc() (fn js.Func[func() js.Promise[PhotoCapabilities]]) {
	return fn.FromRef(
		bindings.ImageCaptureGetPhotoCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetPhotoSettings calls the method "ImageCapture.getPhotoSettings".
//
// The returned bool will be false if there is no such method.
func (this ImageCapture) GetPhotoSettings() (js.Promise[PhotoSettings], bool) {
	var _ok bool
	_ret := bindings.CallImageCaptureGetPhotoSettings(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PhotoSettings]{}.FromRef(_ret), _ok
}

// GetPhotoSettingsFunc returns the method "ImageCapture.getPhotoSettings".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageCapture) GetPhotoSettingsFunc() (fn js.Func[func() js.Promise[PhotoSettings]]) {
	return fn.FromRef(
		bindings.ImageCaptureGetPhotoSettingsFunc(
			this.Ref(),
		),
	)
}

// GrabFrame calls the method "ImageCapture.grabFrame".
//
// The returned bool will be false if there is no such method.
func (this ImageCapture) GrabFrame() (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallImageCaptureGrabFrame(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// GrabFrameFunc returns the method "ImageCapture.grabFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageCapture) GrabFrameFunc() (fn js.Func[func() js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.ImageCaptureGrabFrameFunc(
			this.Ref(),
		),
	)
}

type ImageDecodeOptions struct {
	// FrameIndex is "ImageDecodeOptions.frameIndex"
	//
	// Optional, defaults to 0.
	FrameIndex uint32
	// CompleteFramesOnly is "ImageDecodeOptions.completeFramesOnly"
	//
	// Optional, defaults to true.
	CompleteFramesOnly bool

	FFI_USE_FrameIndex         bool // for FrameIndex.
	FFI_USE_CompleteFramesOnly bool // for CompleteFramesOnly.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDecodeOptions with all fields set.
func (p ImageDecodeOptions) FromRef(ref js.Ref) ImageDecodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ImageDecodeOptions ImageDecodeOptions [// ImageDecodeOptions] [0x14000863360 0x140008634a0 0x14000863400 0x14000863540] 0x14001c3bf50 {0 0}} in the application heap.
func (p ImageDecodeOptions) New() js.Ref {
	return bindings.ImageDecodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageDecodeOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageDecodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageDecodeOptions) Update(ref js.Ref) {
	bindings.ImageDecodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ImageDecodeResult struct {
	// Image is "ImageDecodeResult.image"
	//
	// Required
	Image VideoFrame
	// Complete is "ImageDecodeResult.complete"
	//
	// Required
	Complete bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDecodeResult with all fields set.
func (p ImageDecodeResult) FromRef(ref js.Ref) ImageDecodeResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ImageDecodeResult ImageDecodeResult [// ImageDecodeResult] [0x140008635e0 0x14000863680] 0x1400107e060 {0 0}} in the application heap.
func (p ImageDecodeResult) New() js.Ref {
	return bindings.ImageDecodeResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageDecodeResult) UpdateFrom(ref js.Ref) {
	bindings.ImageDecodeResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageDecodeResult) Update(ref js.Ref) {
	bindings.ImageDecodeResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ImageDecoderInit struct {
	// Type is "ImageDecoderInit.type"
	//
	// Required
	Type js.String
	// Data is "ImageDecoderInit.data"
	//
	// Required
	Data ImageBufferSource
	// ColorSpaceConversion is "ImageDecoderInit.colorSpaceConversion"
	//
	// Optional, defaults to "default".
	ColorSpaceConversion ColorSpaceConversion
	// DesiredWidth is "ImageDecoderInit.desiredWidth"
	//
	// Optional
	DesiredWidth uint32
	// DesiredHeight is "ImageDecoderInit.desiredHeight"
	//
	// Optional
	DesiredHeight uint32
	// PreferAnimation is "ImageDecoderInit.preferAnimation"
	//
	// Optional
	PreferAnimation bool
	// Transfer is "ImageDecoderInit.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.ArrayBuffer]

	FFI_USE_DesiredWidth    bool // for DesiredWidth.
	FFI_USE_DesiredHeight   bool // for DesiredHeight.
	FFI_USE_PreferAnimation bool // for PreferAnimation.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDecoderInit with all fields set.
func (p ImageDecoderInit) FromRef(ref js.Ref) ImageDecoderInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ImageDecoderInit ImageDecoderInit [// ImageDecoderInit] [0x14000863720 0x140008637c0 0x14000863860 0x14000863900 0x14000863a40 0x14000863b80 0x14000863cc0 0x140008639a0 0x14000863ae0 0x14000863c20] 0x1400107e0d8 {0 0}} in the application heap.
func (p ImageDecoderInit) New() js.Ref {
	return bindings.ImageDecoderInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageDecoderInit) UpdateFrom(ref js.Ref) {
	bindings.ImageDecoderInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageDecoderInit) Update(ref js.Ref) {
	bindings.ImageDecoderInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ImageTrack struct {
	ref js.Ref
}

func (this ImageTrack) Once() ImageTrack {
	this.Ref().Once()
	return this
}

func (this ImageTrack) Ref() js.Ref {
	return this.ref
}

func (this ImageTrack) FromRef(ref js.Ref) ImageTrack {
	this.ref = ref
	return this
}

func (this ImageTrack) Free() {
	this.Ref().Free()
}

// Animated returns the value of property "ImageTrack.animated".
//
// The returned bool will be false if there is no such property.
func (this ImageTrack) Animated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackAnimated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FrameCount returns the value of property "ImageTrack.frameCount".
//
// The returned bool will be false if there is no such property.
func (this ImageTrack) FrameCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackFrameCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// RepetitionCount returns the value of property "ImageTrack.repetitionCount".
//
// The returned bool will be false if there is no such property.
func (this ImageTrack) RepetitionCount() (float32, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackRepetitionCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Selected returns the value of property "ImageTrack.selected".
//
// The returned bool will be false if there is no such property.
func (this ImageTrack) Selected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Selected sets the value of property "ImageTrack.selected" to val.
//
// It returns false if the property cannot be set.
func (this ImageTrack) SetSelected(val bool) bool {
	return js.True == bindings.SetImageTrackSelected(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

type ImageTrackList struct {
	ref js.Ref
}

func (this ImageTrackList) Once() ImageTrackList {
	this.Ref().Once()
	return this
}

func (this ImageTrackList) Ref() js.Ref {
	return this.ref
}

func (this ImageTrackList) FromRef(ref js.Ref) ImageTrackList {
	this.ref = ref
	return this
}

func (this ImageTrackList) Free() {
	this.Ref().Free()
}

// Ready returns the value of property "ImageTrackList.ready".
//
// The returned bool will be false if there is no such property.
func (this ImageTrackList) Ready() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetImageTrackListReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Length returns the value of property "ImageTrackList.length".
//
// The returned bool will be false if there is no such property.
func (this ImageTrackList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectedIndex returns the value of property "ImageTrackList.selectedIndex".
//
// The returned bool will be false if there is no such property.
func (this ImageTrackList) SelectedIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackListSelectedIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SelectedTrack returns the value of property "ImageTrackList.selectedTrack".
//
// The returned bool will be false if there is no such property.
func (this ImageTrackList) SelectedTrack() (ImageTrack, bool) {
	var _ok bool
	_ret := bindings.GetImageTrackListSelectedTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return ImageTrack{}.FromRef(_ret), _ok
}

// Get calls the method "ImageTrackList.".
//
// The returned bool will be false if there is no such method.
func (this ImageTrackList) Get(index uint32) (ImageTrack, bool) {
	var _ok bool
	_ret := bindings.CallImageTrackListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return ImageTrack{}.FromRef(_ret), _ok
}

// GetFunc returns the method "ImageTrackList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageTrackList) GetFunc() (fn js.Func[func(index uint32) ImageTrack]) {
	return fn.FromRef(
		bindings.ImageTrackListGetFunc(
			this.Ref(),
		),
	)
}

func NewImageDecoder(init ImageDecoderInit) ImageDecoder {
	return ImageDecoder{}.FromRef(
		bindings.NewImageDecoderByImageDecoder(
			js.Pointer(&init)),
	)
}

type ImageDecoder struct {
	ref js.Ref
}

func (this ImageDecoder) Once() ImageDecoder {
	this.Ref().Once()
	return this
}

func (this ImageDecoder) Ref() js.Ref {
	return this.ref
}

func (this ImageDecoder) FromRef(ref js.Ref) ImageDecoder {
	this.ref = ref
	return this
}

func (this ImageDecoder) Free() {
	this.Ref().Free()
}

// Type returns the value of property "ImageDecoder.type".
//
// The returned bool will be false if there is no such property.
func (this ImageDecoder) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetImageDecoderType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Complete returns the value of property "ImageDecoder.complete".
//
// The returned bool will be false if there is no such property.
func (this ImageDecoder) Complete() (bool, bool) {
	var _ok bool
	_ret := bindings.GetImageDecoderComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Completed returns the value of property "ImageDecoder.completed".
//
// The returned bool will be false if there is no such property.
func (this ImageDecoder) Completed() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetImageDecoderCompleted(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Tracks returns the value of property "ImageDecoder.tracks".
//
// The returned bool will be false if there is no such property.
func (this ImageDecoder) Tracks() (ImageTrackList, bool) {
	var _ok bool
	_ret := bindings.GetImageDecoderTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return ImageTrackList{}.FromRef(_ret), _ok
}

// Decode calls the method "ImageDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this ImageDecoder) Decode(options ImageDecodeOptions) (js.Promise[ImageDecodeResult], bool) {
	var _ok bool
	_ret := bindings.CallImageDecoderDecode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[ImageDecodeResult]{}.FromRef(_ret), _ok
}

// DecodeFunc returns the method "ImageDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageDecoder) DecodeFunc() (fn js.Func[func(options ImageDecodeOptions) js.Promise[ImageDecodeResult]]) {
	return fn.FromRef(
		bindings.ImageDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode1 calls the method "ImageDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this ImageDecoder) Decode1() (js.Promise[ImageDecodeResult], bool) {
	var _ok bool
	_ret := bindings.CallImageDecoderDecode1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ImageDecodeResult]{}.FromRef(_ret), _ok
}

// Decode1Func returns the method "ImageDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageDecoder) Decode1Func() (fn js.Func[func() js.Promise[ImageDecodeResult]]) {
	return fn.FromRef(
		bindings.ImageDecoderDecode1Func(
			this.Ref(),
		),
	)
}

// Reset calls the method "ImageDecoder.reset".
//
// The returned bool will be false if there is no such method.
func (this ImageDecoder) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallImageDecoderReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "ImageDecoder.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageDecoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageDecoderResetFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "ImageDecoder.close".
//
// The returned bool will be false if there is no such method.
func (this ImageDecoder) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallImageDecoderClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "ImageDecoder.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageDecoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageDecoderCloseFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "ImageDecoder.isTypeSupported".
//
// The returned bool will be false if there is no such method.
func (this ImageDecoder) IsTypeSupported(typ js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallImageDecoderIsTypeSupported(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsTypeSupportedFunc returns the staticmethod "ImageDecoder.isTypeSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageDecoder) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.ImageDecoderIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

type ImportExportKind uint32

const (
	_ ImportExportKind = iota

	ImportExportKind_FUNCTION
	ImportExportKind_TABLE
	ImportExportKind_MEMORY
	ImportExportKind_GLOBAL
)

func (ImportExportKind) FromRef(str js.Ref) ImportExportKind {
	return ImportExportKind(bindings.ConstOfImportExportKind(str))
}

func (x ImportExportKind) String() (string, bool) {
	switch x {
	case ImportExportKind_FUNCTION:
		return "function", true
	case ImportExportKind_TABLE:
		return "table", true
	case ImportExportKind_MEMORY:
		return "memory", true
	case ImportExportKind_GLOBAL:
		return "global", true
	default:
		return "", false
	}
}

type InputDeviceCapabilitiesInit struct {
	// FiresTouchEvents is "InputDeviceCapabilitiesInit.firesTouchEvents"
	//
	// Optional, defaults to false.
	FiresTouchEvents bool
	// PointerMovementScrolls is "InputDeviceCapabilitiesInit.pointerMovementScrolls"
	//
	// Optional, defaults to false.
	PointerMovementScrolls bool

	FFI_USE_FiresTouchEvents       bool // for FiresTouchEvents.
	FFI_USE_PointerMovementScrolls bool // for PointerMovementScrolls.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputDeviceCapabilitiesInit with all fields set.
func (p InputDeviceCapabilitiesInit) FromRef(ref js.Ref) InputDeviceCapabilitiesInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 InputDeviceCapabilitiesInit InputDeviceCapabilitiesInit [// InputDeviceCapabilitiesInit] [0x14000863f40 0x1400086a0a0 0x1400086a000 0x1400086a140] 0x1400107e228 {0 0}} in the application heap.
func (p InputDeviceCapabilitiesInit) New() js.Ref {
	return bindings.InputDeviceCapabilitiesInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p InputDeviceCapabilitiesInit) UpdateFrom(ref js.Ref) {
	bindings.InputDeviceCapabilitiesInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p InputDeviceCapabilitiesInit) Update(ref js.Ref) {
	bindings.InputDeviceCapabilitiesInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewInputDeviceCapabilities(deviceInitDict InputDeviceCapabilitiesInit) InputDeviceCapabilities {
	return InputDeviceCapabilities{}.FromRef(
		bindings.NewInputDeviceCapabilitiesByInputDeviceCapabilities(
			js.Pointer(&deviceInitDict)),
	)
}

func NewInputDeviceCapabilitiesByInputDeviceCapabilities1() InputDeviceCapabilities {
	return InputDeviceCapabilities{}.FromRef(
		bindings.NewInputDeviceCapabilitiesByInputDeviceCapabilities1(),
	)
}

type InputDeviceCapabilities struct {
	ref js.Ref
}

func (this InputDeviceCapabilities) Once() InputDeviceCapabilities {
	this.Ref().Once()
	return this
}

func (this InputDeviceCapabilities) Ref() js.Ref {
	return this.ref
}

func (this InputDeviceCapabilities) FromRef(ref js.Ref) InputDeviceCapabilities {
	this.ref = ref
	return this
}

func (this InputDeviceCapabilities) Free() {
	this.Ref().Free()
}

// FiresTouchEvents returns the value of property "InputDeviceCapabilities.firesTouchEvents".
//
// The returned bool will be false if there is no such property.
func (this InputDeviceCapabilities) FiresTouchEvents() (bool, bool) {
	var _ok bool
	_ret := bindings.GetInputDeviceCapabilitiesFiresTouchEvents(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// PointerMovementScrolls returns the value of property "InputDeviceCapabilities.pointerMovementScrolls".
//
// The returned bool will be false if there is no such property.
func (this InputDeviceCapabilities) PointerMovementScrolls() (bool, bool) {
	var _ok bool
	_ret := bindings.GetInputDeviceCapabilitiesPointerMovementScrolls(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type InputDeviceInfo struct {
	MediaDeviceInfo
}

func (this InputDeviceInfo) Once() InputDeviceInfo {
	this.Ref().Once()
	return this
}

func (this InputDeviceInfo) Ref() js.Ref {
	return this.MediaDeviceInfo.Ref()
}

func (this InputDeviceInfo) FromRef(ref js.Ref) InputDeviceInfo {
	this.MediaDeviceInfo = this.MediaDeviceInfo.FromRef(ref)
	return this
}

func (this InputDeviceInfo) Free() {
	this.Ref().Free()
}

// GetCapabilities calls the method "InputDeviceInfo.getCapabilities".
//
// The returned bool will be false if there is no such method.
func (this InputDeviceInfo) GetCapabilities() (MediaTrackCapabilities, bool) {
	var _ret MediaTrackCapabilities
	_ok := js.True == bindings.CallInputDeviceInfoGetCapabilities(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetCapabilitiesFunc returns the method "InputDeviceInfo.getCapabilities".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InputDeviceInfo) GetCapabilitiesFunc() (fn js.Func[func() MediaTrackCapabilities]) {
	return fn.FromRef(
		bindings.InputDeviceInfoGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

type InputEventInit struct {
	// Data is "InputEventInit.data"
	//
	// Optional, defaults to null.
	Data js.String
	// IsComposing is "InputEventInit.isComposing"
	//
	// Optional, defaults to false.
	IsComposing bool
	// InputType is "InputEventInit.inputType"
	//
	// Optional, defaults to "".
	InputType js.String
	// View is "InputEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "InputEventInit.detail"
	//
	// Optional, defaults to 0.
	Detail int32
	// Bubbles is "InputEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "InputEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "InputEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool
	// DataTransfer is "InputEventInit.dataTransfer"
	//
	// Optional, defaults to null.
	DataTransfer DataTransfer
	// TargetRanges is "InputEventInit.targetRanges"
	//
	// Optional, defaults to [].
	TargetRanges js.Array[StaticRange]

	FFI_USE_IsComposing bool // for IsComposing.
	FFI_USE_Detail      bool // for Detail.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputEventInit with all fields set.
func (p InputEventInit) FromRef(ref js.Ref) InputEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 InputEventInit InputEventInit [// InputEventInit] [0x1400086a280 0x1400086a320 0x1400086a500 0x1400086a5a0 0x1400086a640 0x1400086a780 0x1400086a8c0 0x1400086aa00 0x1400086ab40 0x1400086abe0 0x1400086a3c0 0x1400086a6e0 0x1400086a820 0x1400086a960 0x1400086aaa0] 0x1400107e330 {0 0}} in the application heap.
func (p InputEventInit) New() js.Ref {
	return bindings.InputEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p InputEventInit) UpdateFrom(ref js.Ref) {
	bindings.InputEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p InputEventInit) Update(ref js.Ref) {
	bindings.InputEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewInputEvent(typ js.String, eventInitDict InputEventInit) InputEvent {
	return InputEvent{}.FromRef(
		bindings.NewInputEventByInputEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewInputEventByInputEvent1(typ js.String) InputEvent {
	return InputEvent{}.FromRef(
		bindings.NewInputEventByInputEvent1(
			typ.Ref()),
	)
}

type InputEvent struct {
	UIEvent
}

func (this InputEvent) Once() InputEvent {
	this.Ref().Once()
	return this
}

func (this InputEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this InputEvent) FromRef(ref js.Ref) InputEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this InputEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "InputEvent.data".
//
// The returned bool will be false if there is no such property.
func (this InputEvent) Data() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetInputEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsComposing returns the value of property "InputEvent.isComposing".
//
// The returned bool will be false if there is no such property.
func (this InputEvent) IsComposing() (bool, bool) {
	var _ok bool
	_ret := bindings.GetInputEventIsComposing(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// InputType returns the value of property "InputEvent.inputType".
//
// The returned bool will be false if there is no such property.
func (this InputEvent) InputType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetInputEventInputType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DataTransfer returns the value of property "InputEvent.dataTransfer".
//
// The returned bool will be false if there is no such property.
func (this InputEvent) DataTransfer() (DataTransfer, bool) {
	var _ok bool
	_ret := bindings.GetInputEventDataTransfer(
		this.Ref(), js.Pointer(&_ok),
	)
	return DataTransfer{}.FromRef(_ret), _ok
}

// GetTargetRanges calls the method "InputEvent.getTargetRanges".
//
// The returned bool will be false if there is no such method.
func (this InputEvent) GetTargetRanges() (js.Array[StaticRange], bool) {
	var _ok bool
	_ret := bindings.CallInputEventGetTargetRanges(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[StaticRange]{}.FromRef(_ret), _ok
}

// GetTargetRangesFunc returns the method "InputEvent.getTargetRanges".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InputEvent) GetTargetRangesFunc() (fn js.Func[func() js.Array[StaticRange]]) {
	return fn.FromRef(
		bindings.InputEventGetTargetRangesFunc(
			this.Ref(),
		),
	)
}

type ModuleExportDescriptor struct {
	// Name is "ModuleExportDescriptor.name"
	//
	// Required
	Name js.String
	// Kind is "ModuleExportDescriptor.kind"
	//
	// Required
	Kind ImportExportKind

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ModuleExportDescriptor with all fields set.
func (p ModuleExportDescriptor) FromRef(ref js.Ref) ModuleExportDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ModuleExportDescriptor ModuleExportDescriptor [// ModuleExportDescriptor] [0x1400086ad20 0x1400086adc0] 0x1400107e3c0 {0 0}} in the application heap.
func (p ModuleExportDescriptor) New() js.Ref {
	return bindings.ModuleExportDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ModuleExportDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ModuleExportDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ModuleExportDescriptor) Update(ref js.Ref) {
	bindings.ModuleExportDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ModuleImportDescriptor struct {
	// Module is "ModuleImportDescriptor.module"
	//
	// Required
	Module js.String
	// Name is "ModuleImportDescriptor.name"
	//
	// Required
	Name js.String
	// Kind is "ModuleImportDescriptor.kind"
	//
	// Required
	Kind ImportExportKind

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ModuleImportDescriptor with all fields set.
func (p ModuleImportDescriptor) FromRef(ref js.Ref) ModuleImportDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ModuleImportDescriptor ModuleImportDescriptor [// ModuleImportDescriptor] [0x1400086ae60 0x1400086af00 0x1400086afa0] 0x1400107e540 {0 0}} in the application heap.
func (p ModuleImportDescriptor) New() js.Ref {
	return bindings.ModuleImportDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ModuleImportDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ModuleImportDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ModuleImportDescriptor) Update(ref js.Ref) {
	bindings.ModuleImportDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewModule(bytes BufferSource) Module {
	return Module{}.FromRef(
		bindings.NewModuleByModule(
			bytes.Ref()),
	)
}

type Module struct {
	ref js.Ref
}

func (this Module) Once() Module {
	this.Ref().Once()
	return this
}

func (this Module) Ref() js.Ref {
	return this.ref
}

func (this Module) FromRef(ref js.Ref) Module {
	this.ref = ref
	return this
}

func (this Module) Free() {
	this.Ref().Free()
}

// Exports calls the staticmethod "Module.exports".
//
// The returned bool will be false if there is no such method.
func (this Module) Exports(moduleObject Module) (js.Array[ModuleExportDescriptor], bool) {
	var _ok bool
	_ret := bindings.CallModuleExports(
		this.Ref(), js.Pointer(&_ok),
		moduleObject.Ref(),
	)

	return js.Array[ModuleExportDescriptor]{}.FromRef(_ret), _ok
}

// ExportsFunc returns the staticmethod "Module.exports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Module) ExportsFunc() (fn js.Func[func(moduleObject Module) js.Array[ModuleExportDescriptor]]) {
	return fn.FromRef(
		bindings.ModuleExportsFunc(
			this.Ref(),
		),
	)
}

// Imports calls the staticmethod "Module.imports".
//
// The returned bool will be false if there is no such method.
func (this Module) Imports(moduleObject Module) (js.Array[ModuleImportDescriptor], bool) {
	var _ok bool
	_ret := bindings.CallModuleImports(
		this.Ref(), js.Pointer(&_ok),
		moduleObject.Ref(),
	)

	return js.Array[ModuleImportDescriptor]{}.FromRef(_ret), _ok
}

// ImportsFunc returns the staticmethod "Module.imports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Module) ImportsFunc() (fn js.Func[func(moduleObject Module) js.Array[ModuleImportDescriptor]]) {
	return fn.FromRef(
		bindings.ModuleImportsFunc(
			this.Ref(),
		),
	)
}

// CustomSections calls the staticmethod "Module.customSections".
//
// The returned bool will be false if there is no such method.
func (this Module) CustomSections(moduleObject Module, sectionName js.String) (js.Array[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallModuleCustomSections(
		this.Ref(), js.Pointer(&_ok),
		moduleObject.Ref(),
		sectionName.Ref(),
	)

	return js.Array[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// CustomSectionsFunc returns the staticmethod "Module.customSections".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Module) CustomSectionsFunc() (fn js.Func[func(moduleObject Module, sectionName js.String) js.Array[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.ModuleCustomSectionsFunc(
			this.Ref(),
		),
	)
}

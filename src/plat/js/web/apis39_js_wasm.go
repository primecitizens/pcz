// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
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
// It returns ok=false if there is no such property.
func (this IDBRequest) Result() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBRequestResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "IDBRequest.error".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetIDBRequestError(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Source returns the value of property "IDBRequest.source".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Source() (ret OneOf_IDBObjectStore_IDBIndex_IDBCursor, ok bool) {
	ok = js.True == bindings.GetIDBRequestSource(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transaction returns the value of property "IDBRequest.transaction".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Transaction() (ret IDBTransaction, ok bool) {
	ok = js.True == bindings.GetIDBRequestTransaction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "IDBRequest.readyState".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) ReadyState() (ret IDBRequestReadyState, ok bool) {
	ok = js.True == bindings.GetIDBRequestReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this IDBCursor) Source() (ret OneOf_IDBObjectStore_IDBIndex, ok bool) {
	ok = js.True == bindings.GetIDBCursorSource(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "IDBCursor.direction".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Direction() (ret IDBCursorDirection, ok bool) {
	ok = js.True == bindings.GetIDBCursorDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "IDBCursor.key".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Key() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PrimaryKey returns the value of property "IDBCursor.primaryKey".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) PrimaryKey() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorPrimaryKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Request returns the value of property "IDBCursor.request".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Request() (ret IDBRequest, ok bool) {
	ok = js.True == bindings.GetIDBCursorRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAdvance returns true if the method "IDBCursor.advance" exists.
func (this IDBCursor) HasAdvance() bool {
	return js.True == bindings.HasIDBCursorAdvance(
		this.Ref(),
	)
}

// AdvanceFunc returns the method "IDBCursor.advance".
func (this IDBCursor) AdvanceFunc() (fn js.Func[func(count uint32)]) {
	return fn.FromRef(
		bindings.IDBCursorAdvanceFunc(
			this.Ref(),
		),
	)
}

// Advance calls the method "IDBCursor.advance".
func (this IDBCursor) Advance(count uint32) (ret js.Void) {
	bindings.CallIDBCursorAdvance(
		this.Ref(), js.Pointer(&ret),
		uint32(count),
	)

	return
}

// TryAdvance calls the method "IDBCursor.advance"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryAdvance(count uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorAdvance(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(count),
	)

	return
}

// HasContinue returns true if the method "IDBCursor.continue" exists.
func (this IDBCursor) HasContinue() bool {
	return js.True == bindings.HasIDBCursorContinue(
		this.Ref(),
	)
}

// ContinueFunc returns the method "IDBCursor.continue".
func (this IDBCursor) ContinueFunc() (fn js.Func[func(key js.Any)]) {
	return fn.FromRef(
		bindings.IDBCursorContinueFunc(
			this.Ref(),
		),
	)
}

// Continue calls the method "IDBCursor.continue".
func (this IDBCursor) Continue(key js.Any) (ret js.Void) {
	bindings.CallIDBCursorContinue(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryContinue calls the method "IDBCursor.continue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryContinue(key js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorContinue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasContinue1 returns true if the method "IDBCursor.continue" exists.
func (this IDBCursor) HasContinue1() bool {
	return js.True == bindings.HasIDBCursorContinue1(
		this.Ref(),
	)
}

// Continue1Func returns the method "IDBCursor.continue".
func (this IDBCursor) Continue1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBCursorContinue1Func(
			this.Ref(),
		),
	)
}

// Continue1 calls the method "IDBCursor.continue".
func (this IDBCursor) Continue1() (ret js.Void) {
	bindings.CallIDBCursorContinue1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryContinue1 calls the method "IDBCursor.continue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryContinue1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorContinue1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasContinuePrimaryKey returns true if the method "IDBCursor.continuePrimaryKey" exists.
func (this IDBCursor) HasContinuePrimaryKey() bool {
	return js.True == bindings.HasIDBCursorContinuePrimaryKey(
		this.Ref(),
	)
}

// ContinuePrimaryKeyFunc returns the method "IDBCursor.continuePrimaryKey".
func (this IDBCursor) ContinuePrimaryKeyFunc() (fn js.Func[func(key js.Any, primaryKey js.Any)]) {
	return fn.FromRef(
		bindings.IDBCursorContinuePrimaryKeyFunc(
			this.Ref(),
		),
	)
}

// ContinuePrimaryKey calls the method "IDBCursor.continuePrimaryKey".
func (this IDBCursor) ContinuePrimaryKey(key js.Any, primaryKey js.Any) (ret js.Void) {
	bindings.CallIDBCursorContinuePrimaryKey(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
		primaryKey.Ref(),
	)

	return
}

// TryContinuePrimaryKey calls the method "IDBCursor.continuePrimaryKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryContinuePrimaryKey(key js.Any, primaryKey js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorContinuePrimaryKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		primaryKey.Ref(),
	)

	return
}

// HasUpdate returns true if the method "IDBCursor.update" exists.
func (this IDBCursor) HasUpdate() bool {
	return js.True == bindings.HasIDBCursorUpdate(
		this.Ref(),
	)
}

// UpdateFunc returns the method "IDBCursor.update".
func (this IDBCursor) UpdateFunc() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBCursorUpdateFunc(
			this.Ref(),
		),
	)
}

// Update calls the method "IDBCursor.update".
func (this IDBCursor) Update(value js.Any) (ret IDBRequest) {
	bindings.CallIDBCursorUpdate(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryUpdate calls the method "IDBCursor.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryUpdate(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorUpdate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasDelete returns true if the method "IDBCursor.delete" exists.
func (this IDBCursor) HasDelete() bool {
	return js.True == bindings.HasIDBCursorDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "IDBCursor.delete".
func (this IDBCursor) DeleteFunc() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBCursorDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "IDBCursor.delete".
func (this IDBCursor) Delete() (ret IDBRequest) {
	bindings.CallIDBCursorDelete(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDelete calls the method "IDBCursor.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryDelete() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this IDBCursorWithValue) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorWithValueValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this IDBKeyRange) Lower() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeLower(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Upper returns the value of property "IDBKeyRange.upper".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) Upper() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeUpper(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LowerOpen returns the value of property "IDBKeyRange.lowerOpen".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) LowerOpen() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeLowerOpen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpperOpen returns the value of property "IDBKeyRange.upperOpen".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) UpperOpen() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeUpperOpen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasOnly returns true if the staticmethod "IDBKeyRange.only" exists.
func (this IDBKeyRange) HasOnly() bool {
	return js.True == bindings.HasIDBKeyRangeOnly(
		this.Ref(),
	)
}

// OnlyFunc returns the staticmethod "IDBKeyRange.only".
func (this IDBKeyRange) OnlyFunc() (fn js.Func[func(value js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeOnlyFunc(
			this.Ref(),
		),
	)
}

// Only calls the staticmethod "IDBKeyRange.only".
func (this IDBKeyRange) Only(value js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeOnly(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryOnly calls the staticmethod "IDBKeyRange.only"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryOnly(value js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeOnly(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasLowerBound returns true if the staticmethod "IDBKeyRange.lowerBound" exists.
func (this IDBKeyRange) HasLowerBound() bool {
	return js.True == bindings.HasIDBKeyRangeLowerBound(
		this.Ref(),
	)
}

// LowerBoundFunc returns the staticmethod "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBoundFunc() (fn js.Func[func(lower js.Any, open bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeLowerBoundFunc(
			this.Ref(),
		),
	)
}

// LowerBound calls the staticmethod "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBound(lower js.Any, open bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeLowerBound(
		this.Ref(), js.Pointer(&ret),
		lower.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// TryLowerBound calls the staticmethod "IDBKeyRange.lowerBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryLowerBound(lower js.Any, open bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeLowerBound(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// HasLowerBound1 returns true if the staticmethod "IDBKeyRange.lowerBound" exists.
func (this IDBKeyRange) HasLowerBound1() bool {
	return js.True == bindings.HasIDBKeyRangeLowerBound1(
		this.Ref(),
	)
}

// LowerBound1Func returns the staticmethod "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBound1Func() (fn js.Func[func(lower js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeLowerBound1Func(
			this.Ref(),
		),
	)
}

// LowerBound1 calls the staticmethod "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBound1(lower js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeLowerBound1(
		this.Ref(), js.Pointer(&ret),
		lower.Ref(),
	)

	return
}

// TryLowerBound1 calls the staticmethod "IDBKeyRange.lowerBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryLowerBound1(lower js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeLowerBound1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
	)

	return
}

// HasUpperBound returns true if the staticmethod "IDBKeyRange.upperBound" exists.
func (this IDBKeyRange) HasUpperBound() bool {
	return js.True == bindings.HasIDBKeyRangeUpperBound(
		this.Ref(),
	)
}

// UpperBoundFunc returns the staticmethod "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBoundFunc() (fn js.Func[func(upper js.Any, open bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeUpperBoundFunc(
			this.Ref(),
		),
	)
}

// UpperBound calls the staticmethod "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBound(upper js.Any, open bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeUpperBound(
		this.Ref(), js.Pointer(&ret),
		upper.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// TryUpperBound calls the staticmethod "IDBKeyRange.upperBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryUpperBound(upper js.Any, open bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeUpperBound(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		upper.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// HasUpperBound1 returns true if the staticmethod "IDBKeyRange.upperBound" exists.
func (this IDBKeyRange) HasUpperBound1() bool {
	return js.True == bindings.HasIDBKeyRangeUpperBound1(
		this.Ref(),
	)
}

// UpperBound1Func returns the staticmethod "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBound1Func() (fn js.Func[func(upper js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeUpperBound1Func(
			this.Ref(),
		),
	)
}

// UpperBound1 calls the staticmethod "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBound1(upper js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeUpperBound1(
		this.Ref(), js.Pointer(&ret),
		upper.Ref(),
	)

	return
}

// TryUpperBound1 calls the staticmethod "IDBKeyRange.upperBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryUpperBound1(upper js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeUpperBound1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		upper.Ref(),
	)

	return
}

// HasBound returns true if the staticmethod "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasBound() bool {
	return js.True == bindings.HasIDBKeyRangeBound(
		this.Ref(),
	)
}

// BoundFunc returns the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) BoundFunc() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBoundFunc(
			this.Ref(),
		),
	)
}

// Bound calls the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) Bound(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound(
		this.Ref(), js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
		js.Bool(bool(upperOpen)),
	)

	return
}

// TryBound calls the staticmethod "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
		js.Bool(bool(upperOpen)),
	)

	return
}

// HasBound1 returns true if the staticmethod "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasBound1() bool {
	return js.True == bindings.HasIDBKeyRangeBound1(
		this.Ref(),
	)
}

// Bound1Func returns the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) Bound1Func() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBound1Func(
			this.Ref(),
		),
	)
}

// Bound1 calls the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) Bound1(lower js.Any, upper js.Any, lowerOpen bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound1(
		this.Ref(), js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
	)

	return
}

// TryBound1 calls the staticmethod "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound1(lower js.Any, upper js.Any, lowerOpen bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
	)

	return
}

// HasBound2 returns true if the staticmethod "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasBound2() bool {
	return js.True == bindings.HasIDBKeyRangeBound2(
		this.Ref(),
	)
}

// Bound2Func returns the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) Bound2Func() (fn js.Func[func(lower js.Any, upper js.Any) IDBKeyRange]) {
	return fn.FromRef(
		bindings.IDBKeyRangeBound2Func(
			this.Ref(),
		),
	)
}

// Bound2 calls the staticmethod "IDBKeyRange.bound".
func (this IDBKeyRange) Bound2(lower js.Any, upper js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound2(
		this.Ref(), js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
	)

	return
}

// TryBound2 calls the staticmethod "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound2(lower js.Any, upper js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
	)

	return
}

// HasIncludes returns true if the method "IDBKeyRange.includes" exists.
func (this IDBKeyRange) HasIncludes() bool {
	return js.True == bindings.HasIDBKeyRangeIncludes(
		this.Ref(),
	)
}

// IncludesFunc returns the method "IDBKeyRange.includes".
func (this IDBKeyRange) IncludesFunc() (fn js.Func[func(key js.Any) bool]) {
	return fn.FromRef(
		bindings.IDBKeyRangeIncludesFunc(
			this.Ref(),
		),
	)
}

// Includes calls the method "IDBKeyRange.includes".
func (this IDBKeyRange) Includes(key js.Any) (ret bool) {
	bindings.CallIDBKeyRangeIncludes(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryIncludes calls the method "IDBKeyRange.includes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryIncludes(key js.Any) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeIncludes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

type IDBVersionChangeEventInit struct {
	// OldVersion is "IDBVersionChangeEventInit.oldVersion"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_OldVersion MUST be set to true to make this field effective.
	OldVersion uint64
	// NewVersion is "IDBVersionChangeEventInit.newVersion"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_NewVersion MUST be set to true to make this field effective.
	NewVersion uint64
	// Bubbles is "IDBVersionChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "IDBVersionChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "IDBVersionChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new IDBVersionChangeEventInit in the application heap.
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

func NewIDBVersionChangeEvent(typ js.String, eventInitDict IDBVersionChangeEventInit) (ret IDBVersionChangeEvent) {
	ret.ref = bindings.NewIDBVersionChangeEventByIDBVersionChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewIDBVersionChangeEventByIDBVersionChangeEvent1(typ js.String) (ret IDBVersionChangeEvent) {
	ret.ref = bindings.NewIDBVersionChangeEventByIDBVersionChangeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this IDBVersionChangeEvent) OldVersion() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBVersionChangeEventOldVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NewVersion returns the value of property "IDBVersionChangeEvent.newVersion".
//
// It returns ok=false if there is no such property.
func (this IDBVersionChangeEvent) NewVersion() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBVersionChangeEventNewVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this IdentityCredential) Token() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIdentityCredentialToken(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// New creates a new IdentityUserInfo in the application heap.
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

// HasGetUserInfo returns true if the staticmethod "IdentityProvider.getUserInfo" exists.
func (this IdentityProvider) HasGetUserInfo() bool {
	return js.True == bindings.HasIdentityProviderGetUserInfo(
		this.Ref(),
	)
}

// GetUserInfoFunc returns the staticmethod "IdentityProvider.getUserInfo".
func (this IdentityProvider) GetUserInfoFunc() (fn js.Func[func(config IdentityProviderConfig) js.Promise[js.Array[IdentityUserInfo]]]) {
	return fn.FromRef(
		bindings.IdentityProviderGetUserInfoFunc(
			this.Ref(),
		),
	)
}

// GetUserInfo calls the staticmethod "IdentityProvider.getUserInfo".
func (this IdentityProvider) GetUserInfo(config IdentityProviderConfig) (ret js.Promise[js.Array[IdentityUserInfo]]) {
	bindings.CallIdentityProviderGetUserInfo(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryGetUserInfo calls the staticmethod "IdentityProvider.getUserInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdentityProvider) TryGetUserInfo(config IdentityProviderConfig) (ret js.Promise[js.Array[IdentityUserInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdentityProviderGetUserInfo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

type IdentityProviderIcon struct {
	// Url is "IdentityProviderIcon.url"
	//
	// Required
	Url js.String
	// Size is "IdentityProviderIcon.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size uint32

	FFI_USE_Size bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderIcon with all fields set.
func (p IdentityProviderIcon) FromRef(ref js.Ref) IdentityProviderIcon {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IdentityProviderIcon in the application heap.
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

// New creates a new IdentityProviderBranding in the application heap.
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
	//
	// NOTE: Branding.FFI_USE MUST be set to true to get Branding used.
	Branding IdentityProviderBranding

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdentityProviderAPIConfig with all fields set.
func (p IdentityProviderAPIConfig) FromRef(ref js.Ref) IdentityProviderAPIConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IdentityProviderAPIConfig in the application heap.
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

// New creates a new IdentityProviderAccount in the application heap.
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

// New creates a new IdentityProviderAccountList in the application heap.
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

// New creates a new IdentityProviderClientMetadata in the application heap.
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

// New creates a new IdentityProviderToken in the application heap.
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

// New creates a new IdentityProviderWellKnown in the application heap.
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
	//
	// NOTE: FFI_USE_Threshold MUST be set to true to make this field effective.
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

// New creates a new IdleOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this IdleDetector) UserState() (ret UserIdleState, ok bool) {
	ok = js.True == bindings.GetIdleDetectorUserState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenState returns the value of property "IdleDetector.screenState".
//
// It returns ok=false if there is no such property.
func (this IdleDetector) ScreenState() (ret ScreenIdleState, ok bool) {
	ok = js.True == bindings.GetIdleDetectorScreenState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestPermission returns true if the staticmethod "IdleDetector.requestPermission" exists.
func (this IdleDetector) HasRequestPermission() bool {
	return js.True == bindings.HasIdleDetectorRequestPermission(
		this.Ref(),
	)
}

// RequestPermissionFunc returns the staticmethod "IdleDetector.requestPermission".
func (this IdleDetector) RequestPermissionFunc() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.IdleDetectorRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// RequestPermission calls the staticmethod "IdleDetector.requestPermission".
func (this IdleDetector) RequestPermission() (ret js.Promise[PermissionState]) {
	bindings.CallIdleDetectorRequestPermission(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPermission calls the staticmethod "IdleDetector.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryRequestPermission() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorRequestPermission(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStart returns true if the method "IdleDetector.start" exists.
func (this IdleDetector) HasStart() bool {
	return js.True == bindings.HasIdleDetectorStart(
		this.Ref(),
	)
}

// StartFunc returns the method "IdleDetector.start".
func (this IdleDetector) StartFunc() (fn js.Func[func(options IdleOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.IdleDetectorStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "IdleDetector.start".
func (this IdleDetector) Start(options IdleOptions) (ret js.Promise[js.Void]) {
	bindings.CallIdleDetectorStart(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryStart calls the method "IdleDetector.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryStart(options IdleOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasStart1 returns true if the method "IdleDetector.start" exists.
func (this IdleDetector) HasStart1() bool {
	return js.True == bindings.HasIdleDetectorStart1(
		this.Ref(),
	)
}

// Start1Func returns the method "IdleDetector.start".
func (this IdleDetector) Start1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.IdleDetectorStart1Func(
			this.Ref(),
		),
	)
}

// Start1 calls the method "IdleDetector.start".
func (this IdleDetector) Start1() (ret js.Promise[js.Void]) {
	bindings.CallIdleDetectorStart1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "IdleDetector.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryStart1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorStart1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ImageBitmapRenderingContextSettings struct {
	// Alpha is "ImageBitmapRenderingContextSettings.alpha"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha bool

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageBitmapRenderingContextSettings with all fields set.
func (p ImageBitmapRenderingContextSettings) FromRef(ref js.Ref) ImageBitmapRenderingContextSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageBitmapRenderingContextSettings in the application heap.
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
	//
	// NOTE: FFI_USE_ImageHeight MUST be set to true to make this field effective.
	ImageHeight float64
	// ImageWidth is "PhotoSettings.imageWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_ImageWidth MUST be set to true to make this field effective.
	ImageWidth float64
	// RedEyeReduction is "PhotoSettings.redEyeReduction"
	//
	// Optional
	//
	// NOTE: FFI_USE_RedEyeReduction MUST be set to true to make this field effective.
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

// New creates a new PhotoSettings in the application heap.
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
	//
	// NOTE: ImageHeight.FFI_USE MUST be set to true to get ImageHeight used.
	ImageHeight MediaSettingsRange
	// ImageWidth is "PhotoCapabilities.imageWidth"
	//
	// Optional
	//
	// NOTE: ImageWidth.FFI_USE MUST be set to true to get ImageWidth used.
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

// New creates a new PhotoCapabilities in the application heap.
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

func NewImageCapture(videoTrack MediaStreamTrack) (ret ImageCapture) {
	ret.ref = bindings.NewImageCaptureByImageCapture(
		videoTrack.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ImageCapture) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetImageCaptureTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasTakePhoto returns true if the method "ImageCapture.takePhoto" exists.
func (this ImageCapture) HasTakePhoto() bool {
	return js.True == bindings.HasImageCaptureTakePhoto(
		this.Ref(),
	)
}

// TakePhotoFunc returns the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhotoFunc() (fn js.Func[func(photoSettings PhotoSettings) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ImageCaptureTakePhotoFunc(
			this.Ref(),
		),
	)
}

// TakePhoto calls the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhoto(photoSettings PhotoSettings) (ret js.Promise[Blob]) {
	bindings.CallImageCaptureTakePhoto(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&photoSettings),
	)

	return
}

// TryTakePhoto calls the method "ImageCapture.takePhoto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryTakePhoto(photoSettings PhotoSettings) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureTakePhoto(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&photoSettings),
	)

	return
}

// HasTakePhoto1 returns true if the method "ImageCapture.takePhoto" exists.
func (this ImageCapture) HasTakePhoto1() bool {
	return js.True == bindings.HasImageCaptureTakePhoto1(
		this.Ref(),
	)
}

// TakePhoto1Func returns the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhoto1Func() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ImageCaptureTakePhoto1Func(
			this.Ref(),
		),
	)
}

// TakePhoto1 calls the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhoto1() (ret js.Promise[Blob]) {
	bindings.CallImageCaptureTakePhoto1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTakePhoto1 calls the method "ImageCapture.takePhoto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryTakePhoto1() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureTakePhoto1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPhotoCapabilities returns true if the method "ImageCapture.getPhotoCapabilities" exists.
func (this ImageCapture) HasGetPhotoCapabilities() bool {
	return js.True == bindings.HasImageCaptureGetPhotoCapabilities(
		this.Ref(),
	)
}

// GetPhotoCapabilitiesFunc returns the method "ImageCapture.getPhotoCapabilities".
func (this ImageCapture) GetPhotoCapabilitiesFunc() (fn js.Func[func() js.Promise[PhotoCapabilities]]) {
	return fn.FromRef(
		bindings.ImageCaptureGetPhotoCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetPhotoCapabilities calls the method "ImageCapture.getPhotoCapabilities".
func (this ImageCapture) GetPhotoCapabilities() (ret js.Promise[PhotoCapabilities]) {
	bindings.CallImageCaptureGetPhotoCapabilities(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPhotoCapabilities calls the method "ImageCapture.getPhotoCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGetPhotoCapabilities() (ret js.Promise[PhotoCapabilities], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGetPhotoCapabilities(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPhotoSettings returns true if the method "ImageCapture.getPhotoSettings" exists.
func (this ImageCapture) HasGetPhotoSettings() bool {
	return js.True == bindings.HasImageCaptureGetPhotoSettings(
		this.Ref(),
	)
}

// GetPhotoSettingsFunc returns the method "ImageCapture.getPhotoSettings".
func (this ImageCapture) GetPhotoSettingsFunc() (fn js.Func[func() js.Promise[PhotoSettings]]) {
	return fn.FromRef(
		bindings.ImageCaptureGetPhotoSettingsFunc(
			this.Ref(),
		),
	)
}

// GetPhotoSettings calls the method "ImageCapture.getPhotoSettings".
func (this ImageCapture) GetPhotoSettings() (ret js.Promise[PhotoSettings]) {
	bindings.CallImageCaptureGetPhotoSettings(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPhotoSettings calls the method "ImageCapture.getPhotoSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGetPhotoSettings() (ret js.Promise[PhotoSettings], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGetPhotoSettings(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGrabFrame returns true if the method "ImageCapture.grabFrame" exists.
func (this ImageCapture) HasGrabFrame() bool {
	return js.True == bindings.HasImageCaptureGrabFrame(
		this.Ref(),
	)
}

// GrabFrameFunc returns the method "ImageCapture.grabFrame".
func (this ImageCapture) GrabFrameFunc() (fn js.Func[func() js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.ImageCaptureGrabFrameFunc(
			this.Ref(),
		),
	)
}

// GrabFrame calls the method "ImageCapture.grabFrame".
func (this ImageCapture) GrabFrame() (ret js.Promise[ImageBitmap]) {
	bindings.CallImageCaptureGrabFrame(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGrabFrame calls the method "ImageCapture.grabFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGrabFrame() (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGrabFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ImageDecodeOptions struct {
	// FrameIndex is "ImageDecodeOptions.frameIndex"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_FrameIndex MUST be set to true to make this field effective.
	FrameIndex uint32
	// CompleteFramesOnly is "ImageDecodeOptions.completeFramesOnly"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_CompleteFramesOnly MUST be set to true to make this field effective.
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

// New creates a new ImageDecodeOptions in the application heap.
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

// New creates a new ImageDecodeResult in the application heap.
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
	//
	// NOTE: FFI_USE_DesiredWidth MUST be set to true to make this field effective.
	DesiredWidth uint32
	// DesiredHeight is "ImageDecoderInit.desiredHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DesiredHeight MUST be set to true to make this field effective.
	DesiredHeight uint32
	// PreferAnimation is "ImageDecoderInit.preferAnimation"
	//
	// Optional
	//
	// NOTE: FFI_USE_PreferAnimation MUST be set to true to make this field effective.
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

// New creates a new ImageDecoderInit in the application heap.
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
// It returns ok=false if there is no such property.
func (this ImageTrack) Animated() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageTrackAnimated(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FrameCount returns the value of property "ImageTrack.frameCount".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) FrameCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageTrackFrameCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RepetitionCount returns the value of property "ImageTrack.repetitionCount".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) RepetitionCount() (ret float32, ok bool) {
	ok = js.True == bindings.GetImageTrackRepetitionCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Selected returns the value of property "ImageTrack.selected".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageTrackSelected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "ImageTrack.selected" to val.
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
// It returns ok=false if there is no such property.
func (this ImageTrackList) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetImageTrackListReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "ImageTrackList.length".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageTrackListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectedIndex returns the value of property "ImageTrackList.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetImageTrackListSelectedIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectedTrack returns the value of property "ImageTrackList.selectedTrack".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) SelectedTrack() (ret ImageTrack, ok bool) {
	ok = js.True == bindings.GetImageTrackListSelectedTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "ImageTrackList." exists.
func (this ImageTrackList) HasGet() bool {
	return js.True == bindings.HasImageTrackListGet(
		this.Ref(),
	)
}

// GetFunc returns the method "ImageTrackList.".
func (this ImageTrackList) GetFunc() (fn js.Func[func(index uint32) ImageTrack]) {
	return fn.FromRef(
		bindings.ImageTrackListGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "ImageTrackList.".
func (this ImageTrackList) Get(index uint32) (ret ImageTrack) {
	bindings.CallImageTrackListGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "ImageTrackList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageTrackList) TryGet(index uint32) (ret ImageTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageTrackListGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewImageDecoder(init ImageDecoderInit) (ret ImageDecoder) {
	ret.ref = bindings.NewImageDecoderByImageDecoder(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this ImageDecoder) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetImageDecoderType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Complete returns the value of property "ImageDecoder.complete".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Complete() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageDecoderComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Completed returns the value of property "ImageDecoder.completed".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Completed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetImageDecoderCompleted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Tracks returns the value of property "ImageDecoder.tracks".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Tracks() (ret ImageTrackList, ok bool) {
	ok = js.True == bindings.GetImageDecoderTracks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasDecode returns true if the method "ImageDecoder.decode" exists.
func (this ImageDecoder) HasDecode() bool {
	return js.True == bindings.HasImageDecoderDecode(
		this.Ref(),
	)
}

// DecodeFunc returns the method "ImageDecoder.decode".
func (this ImageDecoder) DecodeFunc() (fn js.Func[func(options ImageDecodeOptions) js.Promise[ImageDecodeResult]]) {
	return fn.FromRef(
		bindings.ImageDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "ImageDecoder.decode".
func (this ImageDecoder) Decode(options ImageDecodeOptions) (ret js.Promise[ImageDecodeResult]) {
	bindings.CallImageDecoderDecode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryDecode calls the method "ImageDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryDecode(options ImageDecodeOptions) (ret js.Promise[ImageDecodeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderDecode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasDecode1 returns true if the method "ImageDecoder.decode" exists.
func (this ImageDecoder) HasDecode1() bool {
	return js.True == bindings.HasImageDecoderDecode1(
		this.Ref(),
	)
}

// Decode1Func returns the method "ImageDecoder.decode".
func (this ImageDecoder) Decode1Func() (fn js.Func[func() js.Promise[ImageDecodeResult]]) {
	return fn.FromRef(
		bindings.ImageDecoderDecode1Func(
			this.Ref(),
		),
	)
}

// Decode1 calls the method "ImageDecoder.decode".
func (this ImageDecoder) Decode1() (ret js.Promise[ImageDecodeResult]) {
	bindings.CallImageDecoderDecode1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDecode1 calls the method "ImageDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryDecode1() (ret js.Promise[ImageDecodeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderDecode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "ImageDecoder.reset" exists.
func (this ImageDecoder) HasReset() bool {
	return js.True == bindings.HasImageDecoderReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "ImageDecoder.reset".
func (this ImageDecoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageDecoderResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "ImageDecoder.reset".
func (this ImageDecoder) Reset() (ret js.Void) {
	bindings.CallImageDecoderReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "ImageDecoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "ImageDecoder.close" exists.
func (this ImageDecoder) HasClose() bool {
	return js.True == bindings.HasImageDecoderClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "ImageDecoder.close".
func (this ImageDecoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageDecoderCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "ImageDecoder.close".
func (this ImageDecoder) Close() (ret js.Void) {
	bindings.CallImageDecoderClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ImageDecoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsTypeSupported returns true if the staticmethod "ImageDecoder.isTypeSupported" exists.
func (this ImageDecoder) HasIsTypeSupported() bool {
	return js.True == bindings.HasImageDecoderIsTypeSupported(
		this.Ref(),
	)
}

// IsTypeSupportedFunc returns the staticmethod "ImageDecoder.isTypeSupported".
func (this ImageDecoder) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.ImageDecoderIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "ImageDecoder.isTypeSupported".
func (this ImageDecoder) IsTypeSupported(typ js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallImageDecoderIsTypeSupported(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryIsTypeSupported calls the staticmethod "ImageDecoder.isTypeSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryIsTypeSupported(typ js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderIsTypeSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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
	//
	// NOTE: FFI_USE_FiresTouchEvents MUST be set to true to make this field effective.
	FiresTouchEvents bool
	// PointerMovementScrolls is "InputDeviceCapabilitiesInit.pointerMovementScrolls"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PointerMovementScrolls MUST be set to true to make this field effective.
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

// New creates a new InputDeviceCapabilitiesInit in the application heap.
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

func NewInputDeviceCapabilities(deviceInitDict InputDeviceCapabilitiesInit) (ret InputDeviceCapabilities) {
	ret.ref = bindings.NewInputDeviceCapabilitiesByInputDeviceCapabilities(
		js.Pointer(&deviceInitDict))
	return
}

func NewInputDeviceCapabilitiesByInputDeviceCapabilities1() (ret InputDeviceCapabilities) {
	ret.ref = bindings.NewInputDeviceCapabilitiesByInputDeviceCapabilities1()
	return
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
// It returns ok=false if there is no such property.
func (this InputDeviceCapabilities) FiresTouchEvents() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputDeviceCapabilitiesFiresTouchEvents(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointerMovementScrolls returns the value of property "InputDeviceCapabilities.pointerMovementScrolls".
//
// It returns ok=false if there is no such property.
func (this InputDeviceCapabilities) PointerMovementScrolls() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputDeviceCapabilitiesPointerMovementScrolls(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasGetCapabilities returns true if the method "InputDeviceInfo.getCapabilities" exists.
func (this InputDeviceInfo) HasGetCapabilities() bool {
	return js.True == bindings.HasInputDeviceInfoGetCapabilities(
		this.Ref(),
	)
}

// GetCapabilitiesFunc returns the method "InputDeviceInfo.getCapabilities".
func (this InputDeviceInfo) GetCapabilitiesFunc() (fn js.Func[func() MediaTrackCapabilities]) {
	return fn.FromRef(
		bindings.InputDeviceInfoGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetCapabilities calls the method "InputDeviceInfo.getCapabilities".
func (this InputDeviceInfo) GetCapabilities() (ret MediaTrackCapabilities) {
	bindings.CallInputDeviceInfoGetCapabilities(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCapabilities calls the method "InputDeviceInfo.getCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InputDeviceInfo) TryGetCapabilities() (ret MediaTrackCapabilities, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInputDeviceInfoGetCapabilities(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type InputEventInit struct {
	// Data is "InputEventInit.data"
	//
	// Optional, defaults to null.
	Data js.String
	// IsComposing is "InputEventInit.isComposing"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsComposing MUST be set to true to make this field effective.
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
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "InputEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "InputEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "InputEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new InputEventInit in the application heap.
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

func NewInputEvent(typ js.String, eventInitDict InputEventInit) (ret InputEvent) {
	ret.ref = bindings.NewInputEventByInputEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewInputEventByInputEvent1(typ js.String) (ret InputEvent) {
	ret.ref = bindings.NewInputEventByInputEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this InputEvent) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInputEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsComposing returns the value of property "InputEvent.isComposing".
//
// It returns ok=false if there is no such property.
func (this InputEvent) IsComposing() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputEventIsComposing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InputType returns the value of property "InputEvent.inputType".
//
// It returns ok=false if there is no such property.
func (this InputEvent) InputType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInputEventInputType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DataTransfer returns the value of property "InputEvent.dataTransfer".
//
// It returns ok=false if there is no such property.
func (this InputEvent) DataTransfer() (ret DataTransfer, ok bool) {
	ok = js.True == bindings.GetInputEventDataTransfer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetTargetRanges returns true if the method "InputEvent.getTargetRanges" exists.
func (this InputEvent) HasGetTargetRanges() bool {
	return js.True == bindings.HasInputEventGetTargetRanges(
		this.Ref(),
	)
}

// GetTargetRangesFunc returns the method "InputEvent.getTargetRanges".
func (this InputEvent) GetTargetRangesFunc() (fn js.Func[func() js.Array[StaticRange]]) {
	return fn.FromRef(
		bindings.InputEventGetTargetRangesFunc(
			this.Ref(),
		),
	)
}

// GetTargetRanges calls the method "InputEvent.getTargetRanges".
func (this InputEvent) GetTargetRanges() (ret js.Array[StaticRange]) {
	bindings.CallInputEventGetTargetRanges(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTargetRanges calls the method "InputEvent.getTargetRanges"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InputEvent) TryGetTargetRanges() (ret js.Array[StaticRange], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInputEventGetTargetRanges(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new ModuleExportDescriptor in the application heap.
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

// New creates a new ModuleImportDescriptor in the application heap.
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

func NewModule(bytes BufferSource) (ret Module) {
	ret.ref = bindings.NewModuleByModule(
		bytes.Ref())
	return
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

// HasExports returns true if the staticmethod "Module.exports" exists.
func (this Module) HasExports() bool {
	return js.True == bindings.HasModuleExports(
		this.Ref(),
	)
}

// ExportsFunc returns the staticmethod "Module.exports".
func (this Module) ExportsFunc() (fn js.Func[func(moduleObject Module) js.Array[ModuleExportDescriptor]]) {
	return fn.FromRef(
		bindings.ModuleExportsFunc(
			this.Ref(),
		),
	)
}

// Exports calls the staticmethod "Module.exports".
func (this Module) Exports(moduleObject Module) (ret js.Array[ModuleExportDescriptor]) {
	bindings.CallModuleExports(
		this.Ref(), js.Pointer(&ret),
		moduleObject.Ref(),
	)

	return
}

// TryExports calls the staticmethod "Module.exports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryExports(moduleObject Module) (ret js.Array[ModuleExportDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleExports(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
	)

	return
}

// HasImports returns true if the staticmethod "Module.imports" exists.
func (this Module) HasImports() bool {
	return js.True == bindings.HasModuleImports(
		this.Ref(),
	)
}

// ImportsFunc returns the staticmethod "Module.imports".
func (this Module) ImportsFunc() (fn js.Func[func(moduleObject Module) js.Array[ModuleImportDescriptor]]) {
	return fn.FromRef(
		bindings.ModuleImportsFunc(
			this.Ref(),
		),
	)
}

// Imports calls the staticmethod "Module.imports".
func (this Module) Imports(moduleObject Module) (ret js.Array[ModuleImportDescriptor]) {
	bindings.CallModuleImports(
		this.Ref(), js.Pointer(&ret),
		moduleObject.Ref(),
	)

	return
}

// TryImports calls the staticmethod "Module.imports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryImports(moduleObject Module) (ret js.Array[ModuleImportDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleImports(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
	)

	return
}

// HasCustomSections returns true if the staticmethod "Module.customSections" exists.
func (this Module) HasCustomSections() bool {
	return js.True == bindings.HasModuleCustomSections(
		this.Ref(),
	)
}

// CustomSectionsFunc returns the staticmethod "Module.customSections".
func (this Module) CustomSectionsFunc() (fn js.Func[func(moduleObject Module, sectionName js.String) js.Array[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.ModuleCustomSectionsFunc(
			this.Ref(),
		),
	)
}

// CustomSections calls the staticmethod "Module.customSections".
func (this Module) CustomSections(moduleObject Module, sectionName js.String) (ret js.Array[js.ArrayBuffer]) {
	bindings.CallModuleCustomSections(
		this.Ref(), js.Pointer(&ret),
		moduleObject.Ref(),
		sectionName.Ref(),
	)

	return
}

// TryCustomSections calls the staticmethod "Module.customSections"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryCustomSections(moduleObject Module, sectionName js.String) (ret js.Array[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleCustomSections(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
		sectionName.Ref(),
	)

	return
}

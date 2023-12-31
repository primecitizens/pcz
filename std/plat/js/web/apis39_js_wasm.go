// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// Result returns the value of property "IDBRequest.result".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Result() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBRequestResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "IDBRequest.error".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetIDBRequestError(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Source returns the value of property "IDBRequest.source".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Source() (ret OneOf_IDBObjectStore_IDBIndex_IDBCursor, ok bool) {
	ok = js.True == bindings.GetIDBRequestSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transaction returns the value of property "IDBRequest.transaction".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) Transaction() (ret IDBTransaction, ok bool) {
	ok = js.True == bindings.GetIDBRequestTransaction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "IDBRequest.readyState".
//
// It returns ok=false if there is no such property.
func (this IDBRequest) ReadyState() (ret IDBRequestReadyState, ok bool) {
	ok = js.True == bindings.GetIDBRequestReadyState(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Source returns the value of property "IDBCursor.source".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Source() (ret OneOf_IDBObjectStore_IDBIndex, ok bool) {
	ok = js.True == bindings.GetIDBCursorSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "IDBCursor.direction".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Direction() (ret IDBCursorDirection, ok bool) {
	ok = js.True == bindings.GetIDBCursorDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "IDBCursor.key".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Key() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PrimaryKey returns the value of property "IDBCursor.primaryKey".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) PrimaryKey() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorPrimaryKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Request returns the value of property "IDBCursor.request".
//
// It returns ok=false if there is no such property.
func (this IDBCursor) Request() (ret IDBRequest, ok bool) {
	ok = js.True == bindings.GetIDBCursorRequest(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAdvance returns true if the method "IDBCursor.advance" exists.
func (this IDBCursor) HasFuncAdvance() bool {
	return js.True == bindings.HasFuncIDBCursorAdvance(
		this.ref,
	)
}

// FuncAdvance returns the method "IDBCursor.advance".
func (this IDBCursor) FuncAdvance() (fn js.Func[func(count uint32)]) {
	bindings.FuncIDBCursorAdvance(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Advance calls the method "IDBCursor.advance".
func (this IDBCursor) Advance(count uint32) (ret js.Void) {
	bindings.CallIDBCursorAdvance(
		this.ref, js.Pointer(&ret),
		uint32(count),
	)

	return
}

// TryAdvance calls the method "IDBCursor.advance"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryAdvance(count uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorAdvance(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(count),
	)

	return
}

// HasFuncContinue returns true if the method "IDBCursor.continue" exists.
func (this IDBCursor) HasFuncContinue() bool {
	return js.True == bindings.HasFuncIDBCursorContinue(
		this.ref,
	)
}

// FuncContinue returns the method "IDBCursor.continue".
func (this IDBCursor) FuncContinue() (fn js.Func[func(key js.Any)]) {
	bindings.FuncIDBCursorContinue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Continue calls the method "IDBCursor.continue".
func (this IDBCursor) Continue(key js.Any) (ret js.Void) {
	bindings.CallIDBCursorContinue(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryContinue calls the method "IDBCursor.continue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryContinue(key js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorContinue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncContinue1 returns true if the method "IDBCursor.continue" exists.
func (this IDBCursor) HasFuncContinue1() bool {
	return js.True == bindings.HasFuncIDBCursorContinue1(
		this.ref,
	)
}

// FuncContinue1 returns the method "IDBCursor.continue".
func (this IDBCursor) FuncContinue1() (fn js.Func[func()]) {
	bindings.FuncIDBCursorContinue1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Continue1 calls the method "IDBCursor.continue".
func (this IDBCursor) Continue1() (ret js.Void) {
	bindings.CallIDBCursorContinue1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryContinue1 calls the method "IDBCursor.continue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryContinue1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorContinue1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncContinuePrimaryKey returns true if the method "IDBCursor.continuePrimaryKey" exists.
func (this IDBCursor) HasFuncContinuePrimaryKey() bool {
	return js.True == bindings.HasFuncIDBCursorContinuePrimaryKey(
		this.ref,
	)
}

// FuncContinuePrimaryKey returns the method "IDBCursor.continuePrimaryKey".
func (this IDBCursor) FuncContinuePrimaryKey() (fn js.Func[func(key js.Any, primaryKey js.Any)]) {
	bindings.FuncIDBCursorContinuePrimaryKey(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ContinuePrimaryKey calls the method "IDBCursor.continuePrimaryKey".
func (this IDBCursor) ContinuePrimaryKey(key js.Any, primaryKey js.Any) (ret js.Void) {
	bindings.CallIDBCursorContinuePrimaryKey(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		primaryKey.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the method "IDBCursor.update" exists.
func (this IDBCursor) HasFuncUpdate() bool {
	return js.True == bindings.HasFuncIDBCursorUpdate(
		this.ref,
	)
}

// FuncUpdate returns the method "IDBCursor.update".
func (this IDBCursor) FuncUpdate() (fn js.Func[func(value js.Any) IDBRequest]) {
	bindings.FuncIDBCursorUpdate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Update calls the method "IDBCursor.update".
func (this IDBCursor) Update(value js.Any) (ret IDBRequest) {
	bindings.CallIDBCursorUpdate(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryUpdate calls the method "IDBCursor.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryUpdate(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorUpdate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "IDBCursor.delete" exists.
func (this IDBCursor) HasFuncDelete() bool {
	return js.True == bindings.HasFuncIDBCursorDelete(
		this.ref,
	)
}

// FuncDelete returns the method "IDBCursor.delete".
func (this IDBCursor) FuncDelete() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBCursorDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "IDBCursor.delete".
func (this IDBCursor) Delete() (ret IDBRequest) {
	bindings.CallIDBCursorDelete(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDelete calls the method "IDBCursor.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBCursor) TryDelete() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBCursorDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type IDBCursorWithValue struct {
	IDBCursor
}

func (this IDBCursorWithValue) Once() IDBCursorWithValue {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "IDBCursorWithValue.value".
//
// It returns ok=false if there is no such property.
func (this IDBCursorWithValue) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBCursorWithValueValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

type IDBKeyRange struct {
	ref js.Ref
}

func (this IDBKeyRange) Once() IDBKeyRange {
	this.ref.Once()
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
	this.ref.Free()
}

// Lower returns the value of property "IDBKeyRange.lower".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) Lower() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeLower(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Upper returns the value of property "IDBKeyRange.upper".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) Upper() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeUpper(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LowerOpen returns the value of property "IDBKeyRange.lowerOpen".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) LowerOpen() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeLowerOpen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UpperOpen returns the value of property "IDBKeyRange.upperOpen".
//
// It returns ok=false if there is no such property.
func (this IDBKeyRange) UpperOpen() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBKeyRangeUpperOpen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncOnly returns true if the static method "IDBKeyRange.only" exists.
func (this IDBKeyRange) HasFuncOnly() bool {
	return js.True == bindings.HasFuncIDBKeyRangeOnly(
		this.ref,
	)
}

// FuncOnly returns the static method "IDBKeyRange.only".
func (this IDBKeyRange) FuncOnly() (fn js.Func[func(value js.Any) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeOnly(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Only calls the static method "IDBKeyRange.only".
func (this IDBKeyRange) Only(value js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeOnly(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryOnly calls the static method "IDBKeyRange.only"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryOnly(value js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeOnly(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncLowerBound returns true if the static method "IDBKeyRange.lowerBound" exists.
func (this IDBKeyRange) HasFuncLowerBound() bool {
	return js.True == bindings.HasFuncIDBKeyRangeLowerBound(
		this.ref,
	)
}

// FuncLowerBound returns the static method "IDBKeyRange.lowerBound".
func (this IDBKeyRange) FuncLowerBound() (fn js.Func[func(lower js.Any, open bool) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeLowerBound(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LowerBound calls the static method "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBound(lower js.Any, open bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeLowerBound(
		this.ref, js.Pointer(&ret),
		lower.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// TryLowerBound calls the static method "IDBKeyRange.lowerBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryLowerBound(lower js.Any, open bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeLowerBound(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// HasFuncLowerBound1 returns true if the static method "IDBKeyRange.lowerBound" exists.
func (this IDBKeyRange) HasFuncLowerBound1() bool {
	return js.True == bindings.HasFuncIDBKeyRangeLowerBound1(
		this.ref,
	)
}

// FuncLowerBound1 returns the static method "IDBKeyRange.lowerBound".
func (this IDBKeyRange) FuncLowerBound1() (fn js.Func[func(lower js.Any) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeLowerBound1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LowerBound1 calls the static method "IDBKeyRange.lowerBound".
func (this IDBKeyRange) LowerBound1(lower js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeLowerBound1(
		this.ref, js.Pointer(&ret),
		lower.Ref(),
	)

	return
}

// TryLowerBound1 calls the static method "IDBKeyRange.lowerBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryLowerBound1(lower js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeLowerBound1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
	)

	return
}

// HasFuncUpperBound returns true if the static method "IDBKeyRange.upperBound" exists.
func (this IDBKeyRange) HasFuncUpperBound() bool {
	return js.True == bindings.HasFuncIDBKeyRangeUpperBound(
		this.ref,
	)
}

// FuncUpperBound returns the static method "IDBKeyRange.upperBound".
func (this IDBKeyRange) FuncUpperBound() (fn js.Func[func(upper js.Any, open bool) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeUpperBound(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpperBound calls the static method "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBound(upper js.Any, open bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeUpperBound(
		this.ref, js.Pointer(&ret),
		upper.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// TryUpperBound calls the static method "IDBKeyRange.upperBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryUpperBound(upper js.Any, open bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeUpperBound(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		upper.Ref(),
		js.Bool(bool(open)),
	)

	return
}

// HasFuncUpperBound1 returns true if the static method "IDBKeyRange.upperBound" exists.
func (this IDBKeyRange) HasFuncUpperBound1() bool {
	return js.True == bindings.HasFuncIDBKeyRangeUpperBound1(
		this.ref,
	)
}

// FuncUpperBound1 returns the static method "IDBKeyRange.upperBound".
func (this IDBKeyRange) FuncUpperBound1() (fn js.Func[func(upper js.Any) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeUpperBound1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpperBound1 calls the static method "IDBKeyRange.upperBound".
func (this IDBKeyRange) UpperBound1(upper js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeUpperBound1(
		this.ref, js.Pointer(&ret),
		upper.Ref(),
	)

	return
}

// TryUpperBound1 calls the static method "IDBKeyRange.upperBound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryUpperBound1(upper js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeUpperBound1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		upper.Ref(),
	)

	return
}

// HasFuncBound returns true if the static method "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasFuncBound() bool {
	return js.True == bindings.HasFuncIDBKeyRangeBound(
		this.ref,
	)
}

// FuncBound returns the static method "IDBKeyRange.bound".
func (this IDBKeyRange) FuncBound() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeBound(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Bound calls the static method "IDBKeyRange.bound".
func (this IDBKeyRange) Bound(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound(
		this.ref, js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
		js.Bool(bool(upperOpen)),
	)

	return
}

// TryBound calls the static method "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound(lower js.Any, upper js.Any, lowerOpen bool, upperOpen bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
		js.Bool(bool(upperOpen)),
	)

	return
}

// HasFuncBound1 returns true if the static method "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasFuncBound1() bool {
	return js.True == bindings.HasFuncIDBKeyRangeBound1(
		this.ref,
	)
}

// FuncBound1 returns the static method "IDBKeyRange.bound".
func (this IDBKeyRange) FuncBound1() (fn js.Func[func(lower js.Any, upper js.Any, lowerOpen bool) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeBound1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Bound1 calls the static method "IDBKeyRange.bound".
func (this IDBKeyRange) Bound1(lower js.Any, upper js.Any, lowerOpen bool) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound1(
		this.ref, js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
	)

	return
}

// TryBound1 calls the static method "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound1(lower js.Any, upper js.Any, lowerOpen bool) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
		js.Bool(bool(lowerOpen)),
	)

	return
}

// HasFuncBound2 returns true if the static method "IDBKeyRange.bound" exists.
func (this IDBKeyRange) HasFuncBound2() bool {
	return js.True == bindings.HasFuncIDBKeyRangeBound2(
		this.ref,
	)
}

// FuncBound2 returns the static method "IDBKeyRange.bound".
func (this IDBKeyRange) FuncBound2() (fn js.Func[func(lower js.Any, upper js.Any) IDBKeyRange]) {
	bindings.FuncIDBKeyRangeBound2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Bound2 calls the static method "IDBKeyRange.bound".
func (this IDBKeyRange) Bound2(lower js.Any, upper js.Any) (ret IDBKeyRange) {
	bindings.CallIDBKeyRangeBound2(
		this.ref, js.Pointer(&ret),
		lower.Ref(),
		upper.Ref(),
	)

	return
}

// TryBound2 calls the static method "IDBKeyRange.bound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryBound2(lower js.Any, upper js.Any) (ret IDBKeyRange, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeBound2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lower.Ref(),
		upper.Ref(),
	)

	return
}

// HasFuncIncludes returns true if the method "IDBKeyRange.includes" exists.
func (this IDBKeyRange) HasFuncIncludes() bool {
	return js.True == bindings.HasFuncIDBKeyRangeIncludes(
		this.ref,
	)
}

// FuncIncludes returns the method "IDBKeyRange.includes".
func (this IDBKeyRange) FuncIncludes() (fn js.Func[func(key js.Any) bool]) {
	bindings.FuncIDBKeyRangeIncludes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Includes calls the method "IDBKeyRange.includes".
func (this IDBKeyRange) Includes(key js.Any) (ret bool) {
	bindings.CallIDBKeyRangeIncludes(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryIncludes calls the method "IDBKeyRange.includes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBKeyRange) TryIncludes(key js.Any) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBKeyRangeIncludes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *IDBVersionChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.IDBVersionChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IDBVersionChangeEventInit) Update(ref js.Ref) {
	bindings.IDBVersionChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IDBVersionChangeEventInit) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// OldVersion returns the value of property "IDBVersionChangeEvent.oldVersion".
//
// It returns ok=false if there is no such property.
func (this IDBVersionChangeEvent) OldVersion() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBVersionChangeEventOldVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NewVersion returns the value of property "IDBVersionChangeEvent.newVersion".
//
// It returns ok=false if there is no such property.
func (this IDBVersionChangeEvent) NewVersion() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBVersionChangeEventNewVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

type IdentityCredential struct {
	Credential
}

func (this IdentityCredential) Once() IdentityCredential {
	this.ref.Once()
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
	this.ref.Free()
}

// Token returns the value of property "IdentityCredential.token".
//
// It returns ok=false if there is no such property.
func (this IdentityCredential) Token() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIdentityCredentialToken(
		this.ref, js.Pointer(&ret),
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
func (p *IdentityUserInfo) UpdateFrom(ref js.Ref) {
	bindings.IdentityUserInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityUserInfo) Update(ref js.Ref) {
	bindings.IdentityUserInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityUserInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.Name.Ref(),
		p.GivenName.Ref(),
		p.Picture.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.GivenName = p.GivenName.FromRef(js.Undefined)
	p.Picture = p.Picture.FromRef(js.Undefined)
}

type IdentityProvider struct {
	ref js.Ref
}

func (this IdentityProvider) Once() IdentityProvider {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetUserInfo returns true if the static method "IdentityProvider.getUserInfo" exists.
func (this IdentityProvider) HasFuncGetUserInfo() bool {
	return js.True == bindings.HasFuncIdentityProviderGetUserInfo(
		this.ref,
	)
}

// FuncGetUserInfo returns the static method "IdentityProvider.getUserInfo".
func (this IdentityProvider) FuncGetUserInfo() (fn js.Func[func(config IdentityProviderConfig) js.Promise[js.Array[IdentityUserInfo]]]) {
	bindings.FuncIdentityProviderGetUserInfo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUserInfo calls the static method "IdentityProvider.getUserInfo".
func (this IdentityProvider) GetUserInfo(config IdentityProviderConfig) (ret js.Promise[js.Array[IdentityUserInfo]]) {
	bindings.CallIdentityProviderGetUserInfo(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryGetUserInfo calls the static method "IdentityProvider.getUserInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdentityProvider) TryGetUserInfo(config IdentityProviderConfig) (ret js.Promise[js.Array[IdentityUserInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdentityProviderGetUserInfo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *IdentityProviderIcon) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderIconJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderIcon) Update(ref js.Ref) {
	bindings.IdentityProviderIconJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderIcon) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
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
func (p *IdentityProviderBranding) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderBrandingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderBranding) Update(ref js.Ref) {
	bindings.IdentityProviderBrandingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderBranding) FreeMembers(recursive bool) {
	js.Free(
		p.BackgroundColor.Ref(),
		p.Color.Ref(),
		p.Icons.Ref(),
		p.Name.Ref(),
	)
	p.BackgroundColor = p.BackgroundColor.FromRef(js.Undefined)
	p.Color = p.Color.FromRef(js.Undefined)
	p.Icons = p.Icons.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *IdentityProviderAPIConfig) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAPIConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderAPIConfig) Update(ref js.Ref) {
	bindings.IdentityProviderAPIConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderAPIConfig) FreeMembers(recursive bool) {
	js.Free(
		p.AccountsEndpoint.Ref(),
		p.ClientMetadataEndpoint.Ref(),
		p.IdAssertionEndpoint.Ref(),
	)
	p.AccountsEndpoint = p.AccountsEndpoint.FromRef(js.Undefined)
	p.ClientMetadataEndpoint = p.ClientMetadataEndpoint.FromRef(js.Undefined)
	p.IdAssertionEndpoint = p.IdAssertionEndpoint.FromRef(js.Undefined)
	if recursive {
		p.Branding.FreeMembers(true)
	}
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
func (p *IdentityProviderAccount) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAccountJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderAccount) Update(ref js.Ref) {
	bindings.IdentityProviderAccountJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderAccount) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
		p.Email.Ref(),
		p.GivenName.Ref(),
		p.Picture.Ref(),
		p.ApprovedClients.Ref(),
		p.LoginHints.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Email = p.Email.FromRef(js.Undefined)
	p.GivenName = p.GivenName.FromRef(js.Undefined)
	p.Picture = p.Picture.FromRef(js.Undefined)
	p.ApprovedClients = p.ApprovedClients.FromRef(js.Undefined)
	p.LoginHints = p.LoginHints.FromRef(js.Undefined)
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
func (p *IdentityProviderAccountList) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderAccountListJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderAccountList) Update(ref js.Ref) {
	bindings.IdentityProviderAccountListJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderAccountList) FreeMembers(recursive bool) {
	js.Free(
		p.Accounts.Ref(),
	)
	p.Accounts = p.Accounts.FromRef(js.Undefined)
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
func (p *IdentityProviderClientMetadata) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderClientMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderClientMetadata) Update(ref js.Ref) {
	bindings.IdentityProviderClientMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderClientMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.PrivacyPolicyUrl.Ref(),
		p.TermsOfServiceUrl.Ref(),
	)
	p.PrivacyPolicyUrl = p.PrivacyPolicyUrl.FromRef(js.Undefined)
	p.TermsOfServiceUrl = p.TermsOfServiceUrl.FromRef(js.Undefined)
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
func (p *IdentityProviderToken) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderTokenJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderToken) Update(ref js.Ref) {
	bindings.IdentityProviderTokenJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderToken) FreeMembers(recursive bool) {
	js.Free(
		p.Token.Ref(),
	)
	p.Token = p.Token.FromRef(js.Undefined)
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
func (p *IdentityProviderWellKnown) UpdateFrom(ref js.Ref) {
	bindings.IdentityProviderWellKnownJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdentityProviderWellKnown) Update(ref js.Ref) {
	bindings.IdentityProviderWellKnownJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdentityProviderWellKnown) FreeMembers(recursive bool) {
	js.Free(
		p.ProviderUrls.Ref(),
	)
	p.ProviderUrls = p.ProviderUrls.FromRef(js.Undefined)
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
func (p *IdleOptions) UpdateFrom(ref js.Ref) {
	bindings.IdleOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdleOptions) Update(ref js.Ref) {
	bindings.IdleOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdleOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// UserState returns the value of property "IdleDetector.userState".
//
// It returns ok=false if there is no such property.
func (this IdleDetector) UserState() (ret UserIdleState, ok bool) {
	ok = js.True == bindings.GetIdleDetectorUserState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenState returns the value of property "IdleDetector.screenState".
//
// It returns ok=false if there is no such property.
func (this IdleDetector) ScreenState() (ret ScreenIdleState, ok bool) {
	ok = js.True == bindings.GetIdleDetectorScreenState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestPermission returns true if the static method "IdleDetector.requestPermission" exists.
func (this IdleDetector) HasFuncRequestPermission() bool {
	return js.True == bindings.HasFuncIdleDetectorRequestPermission(
		this.ref,
	)
}

// FuncRequestPermission returns the static method "IdleDetector.requestPermission".
func (this IdleDetector) FuncRequestPermission() (fn js.Func[func() js.Promise[PermissionState]]) {
	bindings.FuncIdleDetectorRequestPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission calls the static method "IdleDetector.requestPermission".
func (this IdleDetector) RequestPermission() (ret js.Promise[PermissionState]) {
	bindings.CallIdleDetectorRequestPermission(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPermission calls the static method "IdleDetector.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryRequestPermission() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorRequestPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStart returns true if the method "IdleDetector.start" exists.
func (this IdleDetector) HasFuncStart() bool {
	return js.True == bindings.HasFuncIdleDetectorStart(
		this.ref,
	)
}

// FuncStart returns the method "IdleDetector.start".
func (this IdleDetector) FuncStart() (fn js.Func[func(options IdleOptions) js.Promise[js.Void]]) {
	bindings.FuncIdleDetectorStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "IdleDetector.start".
func (this IdleDetector) Start(options IdleOptions) (ret js.Promise[js.Void]) {
	bindings.CallIdleDetectorStart(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryStart calls the method "IdleDetector.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryStart(options IdleOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncStart1 returns true if the method "IdleDetector.start" exists.
func (this IdleDetector) HasFuncStart1() bool {
	return js.True == bindings.HasFuncIdleDetectorStart1(
		this.ref,
	)
}

// FuncStart1 returns the method "IdleDetector.start".
func (this IdleDetector) FuncStart1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncIdleDetectorStart1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start1 calls the method "IdleDetector.start".
func (this IdleDetector) Start1() (ret js.Promise[js.Void]) {
	bindings.CallIdleDetectorStart1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "IdleDetector.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDetector) TryStart1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDetectorStart1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ImageBitmapRenderingContextSettings) UpdateFrom(ref js.Ref) {
	bindings.ImageBitmapRenderingContextSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageBitmapRenderingContextSettings) Update(ref js.Ref) {
	bindings.ImageBitmapRenderingContextSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageBitmapRenderingContextSettings) FreeMembers(recursive bool) {
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
func (p *PhotoSettings) UpdateFrom(ref js.Ref) {
	bindings.PhotoSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PhotoSettings) Update(ref js.Ref) {
	bindings.PhotoSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PhotoSettings) FreeMembers(recursive bool) {
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
func (p *PhotoCapabilities) UpdateFrom(ref js.Ref) {
	bindings.PhotoCapabilitiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PhotoCapabilities) Update(ref js.Ref) {
	bindings.PhotoCapabilitiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PhotoCapabilities) FreeMembers(recursive bool) {
	js.Free(
		p.FillLightMode.Ref(),
	)
	p.FillLightMode = p.FillLightMode.FromRef(js.Undefined)
	if recursive {
		p.ImageHeight.FreeMembers(true)
		p.ImageWidth.FreeMembers(true)
	}
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
	this.ref.Once()
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
	this.ref.Free()
}

// Track returns the value of property "ImageCapture.track".
//
// It returns ok=false if there is no such property.
func (this ImageCapture) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetImageCaptureTrack(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncTakePhoto returns true if the method "ImageCapture.takePhoto" exists.
func (this ImageCapture) HasFuncTakePhoto() bool {
	return js.True == bindings.HasFuncImageCaptureTakePhoto(
		this.ref,
	)
}

// FuncTakePhoto returns the method "ImageCapture.takePhoto".
func (this ImageCapture) FuncTakePhoto() (fn js.Func[func(photoSettings PhotoSettings) js.Promise[Blob]]) {
	bindings.FuncImageCaptureTakePhoto(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TakePhoto calls the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhoto(photoSettings PhotoSettings) (ret js.Promise[Blob]) {
	bindings.CallImageCaptureTakePhoto(
		this.ref, js.Pointer(&ret),
		js.Pointer(&photoSettings),
	)

	return
}

// TryTakePhoto calls the method "ImageCapture.takePhoto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryTakePhoto(photoSettings PhotoSettings) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureTakePhoto(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&photoSettings),
	)

	return
}

// HasFuncTakePhoto1 returns true if the method "ImageCapture.takePhoto" exists.
func (this ImageCapture) HasFuncTakePhoto1() bool {
	return js.True == bindings.HasFuncImageCaptureTakePhoto1(
		this.ref,
	)
}

// FuncTakePhoto1 returns the method "ImageCapture.takePhoto".
func (this ImageCapture) FuncTakePhoto1() (fn js.Func[func() js.Promise[Blob]]) {
	bindings.FuncImageCaptureTakePhoto1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TakePhoto1 calls the method "ImageCapture.takePhoto".
func (this ImageCapture) TakePhoto1() (ret js.Promise[Blob]) {
	bindings.CallImageCaptureTakePhoto1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTakePhoto1 calls the method "ImageCapture.takePhoto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryTakePhoto1() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureTakePhoto1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPhotoCapabilities returns true if the method "ImageCapture.getPhotoCapabilities" exists.
func (this ImageCapture) HasFuncGetPhotoCapabilities() bool {
	return js.True == bindings.HasFuncImageCaptureGetPhotoCapabilities(
		this.ref,
	)
}

// FuncGetPhotoCapabilities returns the method "ImageCapture.getPhotoCapabilities".
func (this ImageCapture) FuncGetPhotoCapabilities() (fn js.Func[func() js.Promise[PhotoCapabilities]]) {
	bindings.FuncImageCaptureGetPhotoCapabilities(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPhotoCapabilities calls the method "ImageCapture.getPhotoCapabilities".
func (this ImageCapture) GetPhotoCapabilities() (ret js.Promise[PhotoCapabilities]) {
	bindings.CallImageCaptureGetPhotoCapabilities(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetPhotoCapabilities calls the method "ImageCapture.getPhotoCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGetPhotoCapabilities() (ret js.Promise[PhotoCapabilities], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGetPhotoCapabilities(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPhotoSettings returns true if the method "ImageCapture.getPhotoSettings" exists.
func (this ImageCapture) HasFuncGetPhotoSettings() bool {
	return js.True == bindings.HasFuncImageCaptureGetPhotoSettings(
		this.ref,
	)
}

// FuncGetPhotoSettings returns the method "ImageCapture.getPhotoSettings".
func (this ImageCapture) FuncGetPhotoSettings() (fn js.Func[func() js.Promise[PhotoSettings]]) {
	bindings.FuncImageCaptureGetPhotoSettings(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPhotoSettings calls the method "ImageCapture.getPhotoSettings".
func (this ImageCapture) GetPhotoSettings() (ret js.Promise[PhotoSettings]) {
	bindings.CallImageCaptureGetPhotoSettings(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetPhotoSettings calls the method "ImageCapture.getPhotoSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGetPhotoSettings() (ret js.Promise[PhotoSettings], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGetPhotoSettings(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGrabFrame returns true if the method "ImageCapture.grabFrame" exists.
func (this ImageCapture) HasFuncGrabFrame() bool {
	return js.True == bindings.HasFuncImageCaptureGrabFrame(
		this.ref,
	)
}

// FuncGrabFrame returns the method "ImageCapture.grabFrame".
func (this ImageCapture) FuncGrabFrame() (fn js.Func[func() js.Promise[ImageBitmap]]) {
	bindings.FuncImageCaptureGrabFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GrabFrame calls the method "ImageCapture.grabFrame".
func (this ImageCapture) GrabFrame() (ret js.Promise[ImageBitmap]) {
	bindings.CallImageCaptureGrabFrame(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGrabFrame calls the method "ImageCapture.grabFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageCapture) TryGrabFrame() (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageCaptureGrabFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ImageDecodeOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageDecodeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDecodeOptions) Update(ref js.Ref) {
	bindings.ImageDecodeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDecodeOptions) FreeMembers(recursive bool) {
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
func (p *ImageDecodeResult) UpdateFrom(ref js.Ref) {
	bindings.ImageDecodeResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDecodeResult) Update(ref js.Ref) {
	bindings.ImageDecodeResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDecodeResult) FreeMembers(recursive bool) {
	js.Free(
		p.Image.Ref(),
	)
	p.Image = p.Image.FromRef(js.Undefined)
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
func (p *ImageDecoderInit) UpdateFrom(ref js.Ref) {
	bindings.ImageDecoderInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDecoderInit) Update(ref js.Ref) {
	bindings.ImageDecoderInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDecoderInit) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Data.Ref(),
		p.Transfer.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Transfer = p.Transfer.FromRef(js.Undefined)
}

type ImageTrack struct {
	ref js.Ref
}

func (this ImageTrack) Once() ImageTrack {
	this.ref.Once()
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
	this.ref.Free()
}

// Animated returns the value of property "ImageTrack.animated".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) Animated() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageTrackAnimated(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FrameCount returns the value of property "ImageTrack.frameCount".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) FrameCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageTrackFrameCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RepetitionCount returns the value of property "ImageTrack.repetitionCount".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) RepetitionCount() (ret float32, ok bool) {
	ok = js.True == bindings.GetImageTrackRepetitionCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Selected returns the value of property "ImageTrack.selected".
//
// It returns ok=false if there is no such property.
func (this ImageTrack) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageTrackSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "ImageTrack.selected" to val.
//
// It returns false if the property cannot be set.
func (this ImageTrack) SetSelected(val bool) bool {
	return js.True == bindings.SetImageTrackSelected(
		this.ref,
		js.Bool(bool(val)),
	)
}

type ImageTrackList struct {
	ref js.Ref
}

func (this ImageTrackList) Once() ImageTrackList {
	this.ref.Once()
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
	this.ref.Free()
}

// Ready returns the value of property "ImageTrackList.ready".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetImageTrackListReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "ImageTrackList.length".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageTrackListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectedIndex returns the value of property "ImageTrackList.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetImageTrackListSelectedIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectedTrack returns the value of property "ImageTrackList.selectedTrack".
//
// It returns ok=false if there is no such property.
func (this ImageTrackList) SelectedTrack() (ret ImageTrack, ok bool) {
	ok = js.True == bindings.GetImageTrackListSelectedTrack(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "ImageTrackList." exists.
func (this ImageTrackList) HasFuncGet() bool {
	return js.True == bindings.HasFuncImageTrackListGet(
		this.ref,
	)
}

// FuncGet returns the method "ImageTrackList.".
func (this ImageTrackList) FuncGet() (fn js.Func[func(index uint32) ImageTrack]) {
	bindings.FuncImageTrackListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "ImageTrackList.".
func (this ImageTrackList) Get(index uint32) (ret ImageTrack) {
	bindings.CallImageTrackListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "ImageTrackList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageTrackList) TryGet(index uint32) (ret ImageTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageTrackListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "ImageDecoder.type".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetImageDecoderType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Complete returns the value of property "ImageDecoder.complete".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Complete() (ret bool, ok bool) {
	ok = js.True == bindings.GetImageDecoderComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Completed returns the value of property "ImageDecoder.completed".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Completed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetImageDecoderCompleted(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Tracks returns the value of property "ImageDecoder.tracks".
//
// It returns ok=false if there is no such property.
func (this ImageDecoder) Tracks() (ret ImageTrackList, ok bool) {
	ok = js.True == bindings.GetImageDecoderTracks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncDecode returns true if the method "ImageDecoder.decode" exists.
func (this ImageDecoder) HasFuncDecode() bool {
	return js.True == bindings.HasFuncImageDecoderDecode(
		this.ref,
	)
}

// FuncDecode returns the method "ImageDecoder.decode".
func (this ImageDecoder) FuncDecode() (fn js.Func[func(options ImageDecodeOptions) js.Promise[ImageDecodeResult]]) {
	bindings.FuncImageDecoderDecode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Decode calls the method "ImageDecoder.decode".
func (this ImageDecoder) Decode(options ImageDecodeOptions) (ret js.Promise[ImageDecodeResult]) {
	bindings.CallImageDecoderDecode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryDecode calls the method "ImageDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryDecode(options ImageDecodeOptions) (ret js.Promise[ImageDecodeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderDecode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncDecode1 returns true if the method "ImageDecoder.decode" exists.
func (this ImageDecoder) HasFuncDecode1() bool {
	return js.True == bindings.HasFuncImageDecoderDecode1(
		this.ref,
	)
}

// FuncDecode1 returns the method "ImageDecoder.decode".
func (this ImageDecoder) FuncDecode1() (fn js.Func[func() js.Promise[ImageDecodeResult]]) {
	bindings.FuncImageDecoderDecode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Decode1 calls the method "ImageDecoder.decode".
func (this ImageDecoder) Decode1() (ret js.Promise[ImageDecodeResult]) {
	bindings.CallImageDecoderDecode1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDecode1 calls the method "ImageDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryDecode1() (ret js.Promise[ImageDecodeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderDecode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "ImageDecoder.reset" exists.
func (this ImageDecoder) HasFuncReset() bool {
	return js.True == bindings.HasFuncImageDecoderReset(
		this.ref,
	)
}

// FuncReset returns the method "ImageDecoder.reset".
func (this ImageDecoder) FuncReset() (fn js.Func[func()]) {
	bindings.FuncImageDecoderReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "ImageDecoder.reset".
func (this ImageDecoder) Reset() (ret js.Void) {
	bindings.CallImageDecoderReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "ImageDecoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "ImageDecoder.close" exists.
func (this ImageDecoder) HasFuncClose() bool {
	return js.True == bindings.HasFuncImageDecoderClose(
		this.ref,
	)
}

// FuncClose returns the method "ImageDecoder.close".
func (this ImageDecoder) FuncClose() (fn js.Func[func()]) {
	bindings.FuncImageDecoderClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "ImageDecoder.close".
func (this ImageDecoder) Close() (ret js.Void) {
	bindings.CallImageDecoderClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ImageDecoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsTypeSupported returns true if the static method "ImageDecoder.isTypeSupported" exists.
func (this ImageDecoder) HasFuncIsTypeSupported() bool {
	return js.True == bindings.HasFuncImageDecoderIsTypeSupported(
		this.ref,
	)
}

// FuncIsTypeSupported returns the static method "ImageDecoder.isTypeSupported".
func (this ImageDecoder) FuncIsTypeSupported() (fn js.Func[func(typ js.String) js.Promise[js.Boolean]]) {
	bindings.FuncImageDecoderIsTypeSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTypeSupported calls the static method "ImageDecoder.isTypeSupported".
func (this ImageDecoder) IsTypeSupported(typ js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallImageDecoderIsTypeSupported(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryIsTypeSupported calls the static method "ImageDecoder.isTypeSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageDecoder) TryIsTypeSupported(typ js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageDecoderIsTypeSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *InputDeviceCapabilitiesInit) UpdateFrom(ref js.Ref) {
	bindings.InputDeviceCapabilitiesInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputDeviceCapabilitiesInit) Update(ref js.Ref) {
	bindings.InputDeviceCapabilitiesInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputDeviceCapabilitiesInit) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// FiresTouchEvents returns the value of property "InputDeviceCapabilities.firesTouchEvents".
//
// It returns ok=false if there is no such property.
func (this InputDeviceCapabilities) FiresTouchEvents() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputDeviceCapabilitiesFiresTouchEvents(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointerMovementScrolls returns the value of property "InputDeviceCapabilities.pointerMovementScrolls".
//
// It returns ok=false if there is no such property.
func (this InputDeviceCapabilities) PointerMovementScrolls() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputDeviceCapabilitiesPointerMovementScrolls(
		this.ref, js.Pointer(&ret),
	)
	return
}

type InputDeviceInfo struct {
	MediaDeviceInfo
}

func (this InputDeviceInfo) Once() InputDeviceInfo {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetCapabilities returns true if the method "InputDeviceInfo.getCapabilities" exists.
func (this InputDeviceInfo) HasFuncGetCapabilities() bool {
	return js.True == bindings.HasFuncInputDeviceInfoGetCapabilities(
		this.ref,
	)
}

// FuncGetCapabilities returns the method "InputDeviceInfo.getCapabilities".
func (this InputDeviceInfo) FuncGetCapabilities() (fn js.Func[func() MediaTrackCapabilities]) {
	bindings.FuncInputDeviceInfoGetCapabilities(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCapabilities calls the method "InputDeviceInfo.getCapabilities".
func (this InputDeviceInfo) GetCapabilities() (ret MediaTrackCapabilities) {
	bindings.CallInputDeviceInfoGetCapabilities(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCapabilities calls the method "InputDeviceInfo.getCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InputDeviceInfo) TryGetCapabilities() (ret MediaTrackCapabilities, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInputDeviceInfoGetCapabilities(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *InputEventInit) UpdateFrom(ref js.Ref) {
	bindings.InputEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputEventInit) Update(ref js.Ref) {
	bindings.InputEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.InputType.Ref(),
		p.View.Ref(),
		p.DataTransfer.Ref(),
		p.TargetRanges.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.InputType = p.InputType.FromRef(js.Undefined)
	p.View = p.View.FromRef(js.Undefined)
	p.DataTransfer = p.DataTransfer.FromRef(js.Undefined)
	p.TargetRanges = p.TargetRanges.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "InputEvent.data".
//
// It returns ok=false if there is no such property.
func (this InputEvent) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInputEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsComposing returns the value of property "InputEvent.isComposing".
//
// It returns ok=false if there is no such property.
func (this InputEvent) IsComposing() (ret bool, ok bool) {
	ok = js.True == bindings.GetInputEventIsComposing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputType returns the value of property "InputEvent.inputType".
//
// It returns ok=false if there is no such property.
func (this InputEvent) InputType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInputEventInputType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DataTransfer returns the value of property "InputEvent.dataTransfer".
//
// It returns ok=false if there is no such property.
func (this InputEvent) DataTransfer() (ret DataTransfer, ok bool) {
	ok = js.True == bindings.GetInputEventDataTransfer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetTargetRanges returns true if the method "InputEvent.getTargetRanges" exists.
func (this InputEvent) HasFuncGetTargetRanges() bool {
	return js.True == bindings.HasFuncInputEventGetTargetRanges(
		this.ref,
	)
}

// FuncGetTargetRanges returns the method "InputEvent.getTargetRanges".
func (this InputEvent) FuncGetTargetRanges() (fn js.Func[func() js.Array[StaticRange]]) {
	bindings.FuncInputEventGetTargetRanges(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTargetRanges calls the method "InputEvent.getTargetRanges".
func (this InputEvent) GetTargetRanges() (ret js.Array[StaticRange]) {
	bindings.CallInputEventGetTargetRanges(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetTargetRanges calls the method "InputEvent.getTargetRanges"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InputEvent) TryGetTargetRanges() (ret js.Array[StaticRange], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInputEventGetTargetRanges(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ModuleExportDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ModuleExportDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ModuleExportDescriptor) Update(ref js.Ref) {
	bindings.ModuleExportDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ModuleExportDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *ModuleImportDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ModuleImportDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ModuleImportDescriptor) Update(ref js.Ref) {
	bindings.ModuleImportDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ModuleImportDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Module.Ref(),
		p.Name.Ref(),
	)
	p.Module = p.Module.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncExports returns true if the static method "Module.exports" exists.
func (this Module) HasFuncExports() bool {
	return js.True == bindings.HasFuncModuleExports(
		this.ref,
	)
}

// FuncExports returns the static method "Module.exports".
func (this Module) FuncExports() (fn js.Func[func(moduleObject Module) js.Array[ModuleExportDescriptor]]) {
	bindings.FuncModuleExports(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Exports calls the static method "Module.exports".
func (this Module) Exports(moduleObject Module) (ret js.Array[ModuleExportDescriptor]) {
	bindings.CallModuleExports(
		this.ref, js.Pointer(&ret),
		moduleObject.Ref(),
	)

	return
}

// TryExports calls the static method "Module.exports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryExports(moduleObject Module) (ret js.Array[ModuleExportDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleExports(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
	)

	return
}

// HasFuncImports returns true if the static method "Module.imports" exists.
func (this Module) HasFuncImports() bool {
	return js.True == bindings.HasFuncModuleImports(
		this.ref,
	)
}

// FuncImports returns the static method "Module.imports".
func (this Module) FuncImports() (fn js.Func[func(moduleObject Module) js.Array[ModuleImportDescriptor]]) {
	bindings.FuncModuleImports(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Imports calls the static method "Module.imports".
func (this Module) Imports(moduleObject Module) (ret js.Array[ModuleImportDescriptor]) {
	bindings.CallModuleImports(
		this.ref, js.Pointer(&ret),
		moduleObject.Ref(),
	)

	return
}

// TryImports calls the static method "Module.imports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryImports(moduleObject Module) (ret js.Array[ModuleImportDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleImports(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
	)

	return
}

// HasFuncCustomSections returns true if the static method "Module.customSections" exists.
func (this Module) HasFuncCustomSections() bool {
	return js.True == bindings.HasFuncModuleCustomSections(
		this.ref,
	)
}

// FuncCustomSections returns the static method "Module.customSections".
func (this Module) FuncCustomSections() (fn js.Func[func(moduleObject Module, sectionName js.String) js.Array[js.ArrayBuffer]]) {
	bindings.FuncModuleCustomSections(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CustomSections calls the static method "Module.customSections".
func (this Module) CustomSections(moduleObject Module, sectionName js.String) (ret js.Array[js.ArrayBuffer]) {
	bindings.CallModuleCustomSections(
		this.ref, js.Pointer(&ret),
		moduleObject.Ref(),
		sectionName.Ref(),
	)

	return
}

// TryCustomSections calls the static method "Module.customSections"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Module) TryCustomSections(moduleObject Module, sectionName js.String) (ret js.Array[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryModuleCustomSections(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
		sectionName.Ref(),
	)

	return
}

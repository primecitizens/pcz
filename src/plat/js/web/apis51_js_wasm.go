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

type SVGAnimatedEnumeration struct {
	ref js.Ref
}

func (this SVGAnimatedEnumeration) Once() SVGAnimatedEnumeration {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedEnumeration) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedEnumeration) FromRef(ref js.Ref) SVGAnimatedEnumeration {
	this.ref = ref
	return this
}

func (this SVGAnimatedEnumeration) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedEnumeration.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedEnumeration) BaseVal() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedEnumerationBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedEnumeration.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedEnumeration) SetBaseVal(val uint16) bool {
	return js.True == bindings.SetSVGAnimatedEnumerationBaseVal(
		this.Ref(),
		uint32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedEnumeration.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedEnumeration) AnimVal() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedEnumerationAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGAnimatedInteger struct {
	ref js.Ref
}

func (this SVGAnimatedInteger) Once() SVGAnimatedInteger {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedInteger) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedInteger) FromRef(ref js.Ref) SVGAnimatedInteger {
	this.ref = ref
	return this
}

func (this SVGAnimatedInteger) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedInteger.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedInteger) BaseVal() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedIntegerBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedInteger.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedInteger) SetBaseVal(val int32) bool {
	return js.True == bindings.SetSVGAnimatedIntegerBaseVal(
		this.Ref(),
		int32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedInteger.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedInteger) AnimVal() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedIntegerAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGLengthList struct {
	ref js.Ref
}

func (this SVGLengthList) Once() SVGLengthList {
	this.Ref().Once()
	return this
}

func (this SVGLengthList) Ref() js.Ref {
	return this.ref
}

func (this SVGLengthList) FromRef(ref js.Ref) SVGLengthList {
	this.ref = ref
	return this
}

func (this SVGLengthList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SVGLengthList.length".
//
// It returns ok=false if there is no such property.
func (this SVGLengthList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGLengthListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGLengthList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGLengthList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGLengthListNumberOfItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClear returns true if the method "SVGLengthList.clear" exists.
func (this SVGLengthList) HasClear() bool {
	return js.True == bindings.HasSVGLengthListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "SVGLengthList.clear".
func (this SVGLengthList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGLengthListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SVGLengthList.clear".
func (this SVGLengthList) Clear() (ret js.Void) {
	bindings.CallSVGLengthListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGLengthList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitialize returns true if the method "SVGLengthList.initialize" exists.
func (this SVGLengthList) HasInitialize() bool {
	return js.True == bindings.HasSVGLengthListInitialize(
		this.Ref(),
	)
}

// InitializeFunc returns the method "SVGLengthList.initialize".
func (this SVGLengthList) InitializeFunc() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListInitializeFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGLengthList.initialize".
func (this SVGLengthList) Initialize(newItem SVGLength) (ret SVGLength) {
	bindings.CallSVGLengthListInitialize(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGLengthList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryInitialize(newItem SVGLength) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListInitialize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasGetItem returns true if the method "SVGLengthList.getItem" exists.
func (this SVGLengthList) HasGetItem() bool {
	return js.True == bindings.HasSVGLengthListGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "SVGLengthList.getItem".
func (this SVGLengthList) GetItemFunc() (fn js.Func[func(index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGLengthList.getItem".
func (this SVGLengthList) GetItem(index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListGetItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGLengthList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryGetItem(index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasInsertItemBefore returns true if the method "SVGLengthList.insertItemBefore" exists.
func (this SVGLengthList) HasInsertItemBefore() bool {
	return js.True == bindings.HasSVGLengthListInsertItemBefore(
		this.Ref(),
	)
}

// InsertItemBeforeFunc returns the method "SVGLengthList.insertItemBefore".
func (this SVGLengthList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGLengthList.insertItemBefore".
func (this SVGLengthList) InsertItemBefore(newItem SVGLength, index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListInsertItemBefore(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGLengthList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryInsertItemBefore(newItem SVGLength, index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListInsertItemBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasReplaceItem returns true if the method "SVGLengthList.replaceItem" exists.
func (this SVGLengthList) HasReplaceItem() bool {
	return js.True == bindings.HasSVGLengthListReplaceItem(
		this.Ref(),
	)
}

// ReplaceItemFunc returns the method "SVGLengthList.replaceItem".
func (this SVGLengthList) ReplaceItemFunc() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGLengthList.replaceItem".
func (this SVGLengthList) ReplaceItem(newItem SVGLength, index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListReplaceItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGLengthList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryReplaceItem(newItem SVGLength, index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListReplaceItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasRemoveItem returns true if the method "SVGLengthList.removeItem" exists.
func (this SVGLengthList) HasRemoveItem() bool {
	return js.True == bindings.HasSVGLengthListRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "SVGLengthList.removeItem".
func (this SVGLengthList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGLengthList.removeItem".
func (this SVGLengthList) RemoveItem(index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListRemoveItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGLengthList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryRemoveItem(index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendItem returns true if the method "SVGLengthList.appendItem" exists.
func (this SVGLengthList) HasAppendItem() bool {
	return js.True == bindings.HasSVGLengthListAppendItem(
		this.Ref(),
	)
}

// AppendItemFunc returns the method "SVGLengthList.appendItem".
func (this SVGLengthList) AppendItemFunc() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListAppendItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGLengthList.appendItem".
func (this SVGLengthList) AppendItem(newItem SVGLength) (ret SVGLength) {
	bindings.CallSVGLengthListAppendItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGLengthList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TryAppendItem(newItem SVGLength) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListAppendItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasSet returns true if the method "SVGLengthList." exists.
func (this SVGLengthList) HasSet() bool {
	return js.True == bindings.HasSVGLengthListSet(
		this.Ref(),
	)
}

// SetFunc returns the method "SVGLengthList.".
func (this SVGLengthList) SetFunc() (fn js.Func[func(index uint32, newItem SVGLength)]) {
	return fn.FromRef(
		bindings.SVGLengthListSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGLengthList.".
func (this SVGLengthList) Set(index uint32, newItem SVGLength) (ret js.Void) {
	bindings.CallSVGLengthListSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGLengthList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGLengthList) TrySet(index uint32, newItem SVGLength) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimatedLengthList struct {
	ref js.Ref
}

func (this SVGAnimatedLengthList) Once() SVGAnimatedLengthList {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedLengthList) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedLengthList) FromRef(ref js.Ref) SVGAnimatedLengthList {
	this.ref = ref
	return this
}

func (this SVGAnimatedLengthList) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedLengthList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLengthList) BaseVal() (ret SVGLengthList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthListBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedLengthList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLengthList) AnimVal() (ret SVGLengthList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthListAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGAnimatedNumber struct {
	ref js.Ref
}

func (this SVGAnimatedNumber) Once() SVGAnimatedNumber {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedNumber) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedNumber) FromRef(ref js.Ref) SVGAnimatedNumber {
	this.ref = ref
	return this
}

func (this SVGAnimatedNumber) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedNumber.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumber) BaseVal() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedNumber.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedNumber) SetBaseVal(val float32) bool {
	return js.True == bindings.SetSVGAnimatedNumberBaseVal(
		this.Ref(),
		float32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedNumber.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumber) AnimVal() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGNumberList struct {
	ref js.Ref
}

func (this SVGNumberList) Once() SVGNumberList {
	this.Ref().Once()
	return this
}

func (this SVGNumberList) Ref() js.Ref {
	return this.ref
}

func (this SVGNumberList) FromRef(ref js.Ref) SVGNumberList {
	this.ref = ref
	return this
}

func (this SVGNumberList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SVGNumberList.length".
//
// It returns ok=false if there is no such property.
func (this SVGNumberList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGNumberListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGNumberList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGNumberList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGNumberListNumberOfItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClear returns true if the method "SVGNumberList.clear" exists.
func (this SVGNumberList) HasClear() bool {
	return js.True == bindings.HasSVGNumberListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "SVGNumberList.clear".
func (this SVGNumberList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGNumberListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SVGNumberList.clear".
func (this SVGNumberList) Clear() (ret js.Void) {
	bindings.CallSVGNumberListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGNumberList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitialize returns true if the method "SVGNumberList.initialize" exists.
func (this SVGNumberList) HasInitialize() bool {
	return js.True == bindings.HasSVGNumberListInitialize(
		this.Ref(),
	)
}

// InitializeFunc returns the method "SVGNumberList.initialize".
func (this SVGNumberList) InitializeFunc() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListInitializeFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGNumberList.initialize".
func (this SVGNumberList) Initialize(newItem SVGNumber) (ret SVGNumber) {
	bindings.CallSVGNumberListInitialize(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGNumberList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryInitialize(newItem SVGNumber) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListInitialize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasGetItem returns true if the method "SVGNumberList.getItem" exists.
func (this SVGNumberList) HasGetItem() bool {
	return js.True == bindings.HasSVGNumberListGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "SVGNumberList.getItem".
func (this SVGNumberList) GetItemFunc() (fn js.Func[func(index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGNumberList.getItem".
func (this SVGNumberList) GetItem(index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListGetItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGNumberList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryGetItem(index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasInsertItemBefore returns true if the method "SVGNumberList.insertItemBefore" exists.
func (this SVGNumberList) HasInsertItemBefore() bool {
	return js.True == bindings.HasSVGNumberListInsertItemBefore(
		this.Ref(),
	)
}

// InsertItemBeforeFunc returns the method "SVGNumberList.insertItemBefore".
func (this SVGNumberList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGNumberList.insertItemBefore".
func (this SVGNumberList) InsertItemBefore(newItem SVGNumber, index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListInsertItemBefore(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGNumberList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryInsertItemBefore(newItem SVGNumber, index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListInsertItemBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasReplaceItem returns true if the method "SVGNumberList.replaceItem" exists.
func (this SVGNumberList) HasReplaceItem() bool {
	return js.True == bindings.HasSVGNumberListReplaceItem(
		this.Ref(),
	)
}

// ReplaceItemFunc returns the method "SVGNumberList.replaceItem".
func (this SVGNumberList) ReplaceItemFunc() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGNumberList.replaceItem".
func (this SVGNumberList) ReplaceItem(newItem SVGNumber, index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListReplaceItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGNumberList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryReplaceItem(newItem SVGNumber, index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListReplaceItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasRemoveItem returns true if the method "SVGNumberList.removeItem" exists.
func (this SVGNumberList) HasRemoveItem() bool {
	return js.True == bindings.HasSVGNumberListRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "SVGNumberList.removeItem".
func (this SVGNumberList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGNumberList.removeItem".
func (this SVGNumberList) RemoveItem(index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListRemoveItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGNumberList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryRemoveItem(index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendItem returns true if the method "SVGNumberList.appendItem" exists.
func (this SVGNumberList) HasAppendItem() bool {
	return js.True == bindings.HasSVGNumberListAppendItem(
		this.Ref(),
	)
}

// AppendItemFunc returns the method "SVGNumberList.appendItem".
func (this SVGNumberList) AppendItemFunc() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListAppendItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGNumberList.appendItem".
func (this SVGNumberList) AppendItem(newItem SVGNumber) (ret SVGNumber) {
	bindings.CallSVGNumberListAppendItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGNumberList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TryAppendItem(newItem SVGNumber) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListAppendItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasSet returns true if the method "SVGNumberList." exists.
func (this SVGNumberList) HasSet() bool {
	return js.True == bindings.HasSVGNumberListSet(
		this.Ref(),
	)
}

// SetFunc returns the method "SVGNumberList.".
func (this SVGNumberList) SetFunc() (fn js.Func[func(index uint32, newItem SVGNumber)]) {
	return fn.FromRef(
		bindings.SVGNumberListSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGNumberList.".
func (this SVGNumberList) Set(index uint32, newItem SVGNumber) (ret js.Void) {
	bindings.CallSVGNumberListSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGNumberList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGNumberList) TrySet(index uint32, newItem SVGNumber) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimatedNumberList struct {
	ref js.Ref
}

func (this SVGAnimatedNumberList) Once() SVGAnimatedNumberList {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedNumberList) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedNumberList) FromRef(ref js.Ref) SVGAnimatedNumberList {
	this.ref = ref
	return this
}

func (this SVGAnimatedNumberList) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedNumberList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumberList) BaseVal() (ret SVGNumberList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberListBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedNumberList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumberList) AnimVal() (ret SVGNumberList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberListAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGTransformList struct {
	ref js.Ref
}

func (this SVGTransformList) Once() SVGTransformList {
	this.Ref().Once()
	return this
}

func (this SVGTransformList) Ref() js.Ref {
	return this.ref
}

func (this SVGTransformList) FromRef(ref js.Ref) SVGTransformList {
	this.ref = ref
	return this
}

func (this SVGTransformList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SVGTransformList.length".
//
// It returns ok=false if there is no such property.
func (this SVGTransformList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGTransformListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGTransformList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGTransformList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGTransformListNumberOfItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClear returns true if the method "SVGTransformList.clear" exists.
func (this SVGTransformList) HasClear() bool {
	return js.True == bindings.HasSVGTransformListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "SVGTransformList.clear".
func (this SVGTransformList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGTransformListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SVGTransformList.clear".
func (this SVGTransformList) Clear() (ret js.Void) {
	bindings.CallSVGTransformListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGTransformList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitialize returns true if the method "SVGTransformList.initialize" exists.
func (this SVGTransformList) HasInitialize() bool {
	return js.True == bindings.HasSVGTransformListInitialize(
		this.Ref(),
	)
}

// InitializeFunc returns the method "SVGTransformList.initialize".
func (this SVGTransformList) InitializeFunc() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListInitializeFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGTransformList.initialize".
func (this SVGTransformList) Initialize(newItem SVGTransform) (ret SVGTransform) {
	bindings.CallSVGTransformListInitialize(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGTransformList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryInitialize(newItem SVGTransform) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListInitialize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasGetItem returns true if the method "SVGTransformList.getItem" exists.
func (this SVGTransformList) HasGetItem() bool {
	return js.True == bindings.HasSVGTransformListGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "SVGTransformList.getItem".
func (this SVGTransformList) GetItemFunc() (fn js.Func[func(index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGTransformList.getItem".
func (this SVGTransformList) GetItem(index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListGetItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGTransformList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryGetItem(index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasInsertItemBefore returns true if the method "SVGTransformList.insertItemBefore" exists.
func (this SVGTransformList) HasInsertItemBefore() bool {
	return js.True == bindings.HasSVGTransformListInsertItemBefore(
		this.Ref(),
	)
}

// InsertItemBeforeFunc returns the method "SVGTransformList.insertItemBefore".
func (this SVGTransformList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGTransformList.insertItemBefore".
func (this SVGTransformList) InsertItemBefore(newItem SVGTransform, index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListInsertItemBefore(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGTransformList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryInsertItemBefore(newItem SVGTransform, index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListInsertItemBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasReplaceItem returns true if the method "SVGTransformList.replaceItem" exists.
func (this SVGTransformList) HasReplaceItem() bool {
	return js.True == bindings.HasSVGTransformListReplaceItem(
		this.Ref(),
	)
}

// ReplaceItemFunc returns the method "SVGTransformList.replaceItem".
func (this SVGTransformList) ReplaceItemFunc() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGTransformList.replaceItem".
func (this SVGTransformList) ReplaceItem(newItem SVGTransform, index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListReplaceItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGTransformList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryReplaceItem(newItem SVGTransform, index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListReplaceItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasRemoveItem returns true if the method "SVGTransformList.removeItem" exists.
func (this SVGTransformList) HasRemoveItem() bool {
	return js.True == bindings.HasSVGTransformListRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "SVGTransformList.removeItem".
func (this SVGTransformList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGTransformList.removeItem".
func (this SVGTransformList) RemoveItem(index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListRemoveItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGTransformList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryRemoveItem(index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendItem returns true if the method "SVGTransformList.appendItem" exists.
func (this SVGTransformList) HasAppendItem() bool {
	return js.True == bindings.HasSVGTransformListAppendItem(
		this.Ref(),
	)
}

// AppendItemFunc returns the method "SVGTransformList.appendItem".
func (this SVGTransformList) AppendItemFunc() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListAppendItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGTransformList.appendItem".
func (this SVGTransformList) AppendItem(newItem SVGTransform) (ret SVGTransform) {
	bindings.CallSVGTransformListAppendItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGTransformList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryAppendItem(newItem SVGTransform) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListAppendItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasSet returns true if the method "SVGTransformList." exists.
func (this SVGTransformList) HasSet() bool {
	return js.True == bindings.HasSVGTransformListSet(
		this.Ref(),
	)
}

// SetFunc returns the method "SVGTransformList.".
func (this SVGTransformList) SetFunc() (fn js.Func[func(index uint32, newItem SVGTransform)]) {
	return fn.FromRef(
		bindings.SVGTransformListSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGTransformList.".
func (this SVGTransformList) Set(index uint32, newItem SVGTransform) (ret js.Void) {
	bindings.CallSVGTransformListSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGTransformList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TrySet(index uint32, newItem SVGTransform) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// HasCreateSVGTransformFromMatrix returns true if the method "SVGTransformList.createSVGTransformFromMatrix" exists.
func (this SVGTransformList) HasCreateSVGTransformFromMatrix() bool {
	return js.True == bindings.HasSVGTransformListCreateSVGTransformFromMatrix(
		this.Ref(),
	)
}

// CreateSVGTransformFromMatrixFunc returns the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListCreateSVGTransformFromMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix calls the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform) {
	bindings.CallSVGTransformListCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryCreateSVGTransformFromMatrix calls the method "SVGTransformList.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryCreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasCreateSVGTransformFromMatrix1 returns true if the method "SVGTransformList.createSVGTransformFromMatrix" exists.
func (this SVGTransformList) HasCreateSVGTransformFromMatrix1() bool {
	return js.True == bindings.HasSVGTransformListCreateSVGTransformFromMatrix1(
		this.Ref(),
	)
}

// CreateSVGTransformFromMatrix1Func returns the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrix1Func() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListCreateSVGTransformFromMatrix1Func(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix1 calls the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrix1() (ret SVGTransform) {
	bindings.CallSVGTransformListCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransformFromMatrix1 calls the method "SVGTransformList.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryCreateSVGTransformFromMatrix1() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConsolidate returns true if the method "SVGTransformList.consolidate" exists.
func (this SVGTransformList) HasConsolidate() bool {
	return js.True == bindings.HasSVGTransformListConsolidate(
		this.Ref(),
	)
}

// ConsolidateFunc returns the method "SVGTransformList.consolidate".
func (this SVGTransformList) ConsolidateFunc() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListConsolidateFunc(
			this.Ref(),
		),
	)
}

// Consolidate calls the method "SVGTransformList.consolidate".
func (this SVGTransformList) Consolidate() (ret SVGTransform) {
	bindings.CallSVGTransformListConsolidate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryConsolidate calls the method "SVGTransformList.consolidate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTransformList) TryConsolidate() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListConsolidate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGAnimatedTransformList struct {
	ref js.Ref
}

func (this SVGAnimatedTransformList) Once() SVGAnimatedTransformList {
	this.Ref().Once()
	return this
}

func (this SVGAnimatedTransformList) Ref() js.Ref {
	return this.ref
}

func (this SVGAnimatedTransformList) FromRef(ref js.Ref) SVGAnimatedTransformList {
	this.ref = ref
	return this
}

func (this SVGAnimatedTransformList) Free() {
	this.Ref().Free()
}

// BaseVal returns the value of property "SVGAnimatedTransformList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedTransformList) BaseVal() (ret SVGTransformList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedTransformListBaseVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedTransformList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedTransformList) AnimVal() (ret SVGTransformList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedTransformListAnimVal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGStringList struct {
	ref js.Ref
}

func (this SVGStringList) Once() SVGStringList {
	this.Ref().Once()
	return this
}

func (this SVGStringList) Ref() js.Ref {
	return this.ref
}

func (this SVGStringList) FromRef(ref js.Ref) SVGStringList {
	this.ref = ref
	return this
}

func (this SVGStringList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SVGStringList.length".
//
// It returns ok=false if there is no such property.
func (this SVGStringList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGStringListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGStringList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGStringList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGStringListNumberOfItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClear returns true if the method "SVGStringList.clear" exists.
func (this SVGStringList) HasClear() bool {
	return js.True == bindings.HasSVGStringListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "SVGStringList.clear".
func (this SVGStringList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGStringListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SVGStringList.clear".
func (this SVGStringList) Clear() (ret js.Void) {
	bindings.CallSVGStringListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGStringList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitialize returns true if the method "SVGStringList.initialize" exists.
func (this SVGStringList) HasInitialize() bool {
	return js.True == bindings.HasSVGStringListInitialize(
		this.Ref(),
	)
}

// InitializeFunc returns the method "SVGStringList.initialize".
func (this SVGStringList) InitializeFunc() (fn js.Func[func(newItem js.String) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListInitializeFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGStringList.initialize".
func (this SVGStringList) Initialize(newItem js.String) (ret js.String) {
	bindings.CallSVGStringListInitialize(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGStringList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryInitialize(newItem js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListInitialize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasGetItem returns true if the method "SVGStringList.getItem" exists.
func (this SVGStringList) HasGetItem() bool {
	return js.True == bindings.HasSVGStringListGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "SVGStringList.getItem".
func (this SVGStringList) GetItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGStringList.getItem".
func (this SVGStringList) GetItem(index uint32) (ret js.String) {
	bindings.CallSVGStringListGetItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGStringList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryGetItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasInsertItemBefore returns true if the method "SVGStringList.insertItemBefore" exists.
func (this SVGStringList) HasInsertItemBefore() bool {
	return js.True == bindings.HasSVGStringListInsertItemBefore(
		this.Ref(),
	)
}

// InsertItemBeforeFunc returns the method "SVGStringList.insertItemBefore".
func (this SVGStringList) InsertItemBeforeFunc() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGStringList.insertItemBefore".
func (this SVGStringList) InsertItemBefore(newItem js.String, index uint32) (ret js.String) {
	bindings.CallSVGStringListInsertItemBefore(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGStringList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryInsertItemBefore(newItem js.String, index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListInsertItemBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasReplaceItem returns true if the method "SVGStringList.replaceItem" exists.
func (this SVGStringList) HasReplaceItem() bool {
	return js.True == bindings.HasSVGStringListReplaceItem(
		this.Ref(),
	)
}

// ReplaceItemFunc returns the method "SVGStringList.replaceItem".
func (this SVGStringList) ReplaceItemFunc() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGStringList.replaceItem".
func (this SVGStringList) ReplaceItem(newItem js.String, index uint32) (ret js.String) {
	bindings.CallSVGStringListReplaceItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGStringList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryReplaceItem(newItem js.String, index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListReplaceItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasRemoveItem returns true if the method "SVGStringList.removeItem" exists.
func (this SVGStringList) HasRemoveItem() bool {
	return js.True == bindings.HasSVGStringListRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "SVGStringList.removeItem".
func (this SVGStringList) RemoveItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGStringList.removeItem".
func (this SVGStringList) RemoveItem(index uint32) (ret js.String) {
	bindings.CallSVGStringListRemoveItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGStringList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryRemoveItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendItem returns true if the method "SVGStringList.appendItem" exists.
func (this SVGStringList) HasAppendItem() bool {
	return js.True == bindings.HasSVGStringListAppendItem(
		this.Ref(),
	)
}

// AppendItemFunc returns the method "SVGStringList.appendItem".
func (this SVGStringList) AppendItemFunc() (fn js.Func[func(newItem js.String) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListAppendItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGStringList.appendItem".
func (this SVGStringList) AppendItem(newItem js.String) (ret js.String) {
	bindings.CallSVGStringListAppendItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGStringList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TryAppendItem(newItem js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListAppendItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasSet returns true if the method "SVGStringList." exists.
func (this SVGStringList) HasSet() bool {
	return js.True == bindings.HasSVGStringListSet(
		this.Ref(),
	)
}

// SetFunc returns the method "SVGStringList.".
func (this SVGStringList) SetFunc() (fn js.Func[func(index uint32, newItem js.String)]) {
	return fn.FromRef(
		bindings.SVGStringListSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGStringList.".
func (this SVGStringList) Set(index uint32, newItem js.String) (ret js.Void) {
	bindings.CallSVGStringListSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGStringList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGStringList) TrySet(index uint32, newItem js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimationElement struct {
	SVGElement
}

func (this SVGAnimationElement) Once() SVGAnimationElement {
	this.Ref().Once()
	return this
}

func (this SVGAnimationElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGAnimationElement) FromRef(ref js.Ref) SVGAnimationElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGAnimationElement) Free() {
	this.Ref().Free()
}

// TargetElement returns the value of property "SVGAnimationElement.targetElement".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) TargetElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementTargetElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RequiredExtensions returns the value of property "SVGAnimationElement.requiredExtensions".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) RequiredExtensions() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementRequiredExtensions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SystemLanguage returns the value of property "SVGAnimationElement.systemLanguage".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) SystemLanguage() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementSystemLanguage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetStartTime returns true if the method "SVGAnimationElement.getStartTime" exists.
func (this SVGAnimationElement) HasGetStartTime() bool {
	return js.True == bindings.HasSVGAnimationElementGetStartTime(
		this.Ref(),
	)
}

// GetStartTimeFunc returns the method "SVGAnimationElement.getStartTime".
func (this SVGAnimationElement) GetStartTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetStartTimeFunc(
			this.Ref(),
		),
	)
}

// GetStartTime calls the method "SVGAnimationElement.getStartTime".
func (this SVGAnimationElement) GetStartTime() (ret float32) {
	bindings.CallSVGAnimationElementGetStartTime(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStartTime calls the method "SVGAnimationElement.getStartTime"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryGetStartTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetStartTime(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetCurrentTime returns true if the method "SVGAnimationElement.getCurrentTime" exists.
func (this SVGAnimationElement) HasGetCurrentTime() bool {
	return js.True == bindings.HasSVGAnimationElementGetCurrentTime(
		this.Ref(),
	)
}

// GetCurrentTimeFunc returns the method "SVGAnimationElement.getCurrentTime".
func (this SVGAnimationElement) GetCurrentTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTime calls the method "SVGAnimationElement.getCurrentTime".
func (this SVGAnimationElement) GetCurrentTime() (ret float32) {
	bindings.CallSVGAnimationElementGetCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTime calls the method "SVGAnimationElement.getCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryGetCurrentTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetCurrentTime(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSimpleDuration returns true if the method "SVGAnimationElement.getSimpleDuration" exists.
func (this SVGAnimationElement) HasGetSimpleDuration() bool {
	return js.True == bindings.HasSVGAnimationElementGetSimpleDuration(
		this.Ref(),
	)
}

// GetSimpleDurationFunc returns the method "SVGAnimationElement.getSimpleDuration".
func (this SVGAnimationElement) GetSimpleDurationFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetSimpleDurationFunc(
			this.Ref(),
		),
	)
}

// GetSimpleDuration calls the method "SVGAnimationElement.getSimpleDuration".
func (this SVGAnimationElement) GetSimpleDuration() (ret float32) {
	bindings.CallSVGAnimationElementGetSimpleDuration(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSimpleDuration calls the method "SVGAnimationElement.getSimpleDuration"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryGetSimpleDuration() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetSimpleDuration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBeginElement returns true if the method "SVGAnimationElement.beginElement" exists.
func (this SVGAnimationElement) HasBeginElement() bool {
	return js.True == bindings.HasSVGAnimationElementBeginElement(
		this.Ref(),
	)
}

// BeginElementFunc returns the method "SVGAnimationElement.beginElement".
func (this SVGAnimationElement) BeginElementFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGAnimationElementBeginElementFunc(
			this.Ref(),
		),
	)
}

// BeginElement calls the method "SVGAnimationElement.beginElement".
func (this SVGAnimationElement) BeginElement() (ret js.Void) {
	bindings.CallSVGAnimationElementBeginElement(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBeginElement calls the method "SVGAnimationElement.beginElement"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryBeginElement() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementBeginElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBeginElementAt returns true if the method "SVGAnimationElement.beginElementAt" exists.
func (this SVGAnimationElement) HasBeginElementAt() bool {
	return js.True == bindings.HasSVGAnimationElementBeginElementAt(
		this.Ref(),
	)
}

// BeginElementAtFunc returns the method "SVGAnimationElement.beginElementAt".
func (this SVGAnimationElement) BeginElementAtFunc() (fn js.Func[func(offset float32)]) {
	return fn.FromRef(
		bindings.SVGAnimationElementBeginElementAtFunc(
			this.Ref(),
		),
	)
}

// BeginElementAt calls the method "SVGAnimationElement.beginElementAt".
func (this SVGAnimationElement) BeginElementAt(offset float32) (ret js.Void) {
	bindings.CallSVGAnimationElementBeginElementAt(
		this.Ref(), js.Pointer(&ret),
		float32(offset),
	)

	return
}

// TryBeginElementAt calls the method "SVGAnimationElement.beginElementAt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryBeginElementAt(offset float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementBeginElementAt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(offset),
	)

	return
}

// HasEndElement returns true if the method "SVGAnimationElement.endElement" exists.
func (this SVGAnimationElement) HasEndElement() bool {
	return js.True == bindings.HasSVGAnimationElementEndElement(
		this.Ref(),
	)
}

// EndElementFunc returns the method "SVGAnimationElement.endElement".
func (this SVGAnimationElement) EndElementFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGAnimationElementEndElementFunc(
			this.Ref(),
		),
	)
}

// EndElement calls the method "SVGAnimationElement.endElement".
func (this SVGAnimationElement) EndElement() (ret js.Void) {
	bindings.CallSVGAnimationElementEndElement(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEndElement calls the method "SVGAnimationElement.endElement"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryEndElement() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementEndElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasEndElementAt returns true if the method "SVGAnimationElement.endElementAt" exists.
func (this SVGAnimationElement) HasEndElementAt() bool {
	return js.True == bindings.HasSVGAnimationElementEndElementAt(
		this.Ref(),
	)
}

// EndElementAtFunc returns the method "SVGAnimationElement.endElementAt".
func (this SVGAnimationElement) EndElementAtFunc() (fn js.Func[func(offset float32)]) {
	return fn.FromRef(
		bindings.SVGAnimationElementEndElementAtFunc(
			this.Ref(),
		),
	)
}

// EndElementAt calls the method "SVGAnimationElement.endElementAt".
func (this SVGAnimationElement) EndElementAt(offset float32) (ret js.Void) {
	bindings.CallSVGAnimationElementEndElementAt(
		this.Ref(), js.Pointer(&ret),
		float32(offset),
	)

	return
}

// TryEndElementAt calls the method "SVGAnimationElement.endElementAt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGAnimationElement) TryEndElementAt(offset float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementEndElementAt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(offset),
	)

	return
}

type SVGBoundingBoxOptions struct {
	// Fill is "SVGBoundingBoxOptions.fill"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Fill MUST be set to true to make this field effective.
	Fill bool
	// Stroke is "SVGBoundingBoxOptions.stroke"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Stroke MUST be set to true to make this field effective.
	Stroke bool
	// Markers is "SVGBoundingBoxOptions.markers"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Markers MUST be set to true to make this field effective.
	Markers bool
	// Clipped is "SVGBoundingBoxOptions.clipped"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Clipped MUST be set to true to make this field effective.
	Clipped bool

	FFI_USE_Fill    bool // for Fill.
	FFI_USE_Stroke  bool // for Stroke.
	FFI_USE_Markers bool // for Markers.
	FFI_USE_Clipped bool // for Clipped.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SVGBoundingBoxOptions with all fields set.
func (p SVGBoundingBoxOptions) FromRef(ref js.Ref) SVGBoundingBoxOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SVGBoundingBoxOptions in the application heap.
func (p SVGBoundingBoxOptions) New() js.Ref {
	return bindings.SVGBoundingBoxOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SVGBoundingBoxOptions) UpdateFrom(ref js.Ref) {
	bindings.SVGBoundingBoxOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SVGBoundingBoxOptions) Update(ref js.Ref) {
	bindings.SVGBoundingBoxOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SVGCircleElement struct {
	SVGGeometryElement
}

func (this SVGCircleElement) Once() SVGCircleElement {
	this.Ref().Once()
	return this
}

func (this SVGCircleElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGCircleElement) FromRef(ref js.Ref) SVGCircleElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGCircleElement) Free() {
	this.Ref().Free()
}

// Cx returns the value of property "SVGCircleElement.cx".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) Cx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementCx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cy returns the value of property "SVGCircleElement.cy".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) Cy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementCy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// R returns the value of property "SVGCircleElement.r".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) R() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementR(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGClipPathElement struct {
	SVGElement
}

func (this SVGClipPathElement) Once() SVGClipPathElement {
	this.Ref().Once()
	return this
}

func (this SVGClipPathElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGClipPathElement) FromRef(ref js.Ref) SVGClipPathElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGClipPathElement) Free() {
	this.Ref().Free()
}

// ClipPathUnits returns the value of property "SVGClipPathElement.clipPathUnits".
//
// It returns ok=false if there is no such property.
func (this SVGClipPathElement) ClipPathUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGClipPathElementClipPathUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "SVGClipPathElement.transform".
//
// It returns ok=false if there is no such property.
func (this SVGClipPathElement) Transform() (ret SVGAnimatedTransformList, ok bool) {
	ok = js.True == bindings.GetSVGClipPathElementTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_UNKNOWN  uint16 = 0
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_IDENTITY uint16 = 1
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_TABLE    uint16 = 2
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_DISCRETE uint16 = 3
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_LINEAR   uint16 = 4
	SVGComponentTransferFunctionElement_SVG_FECOMPONENTTRANSFER_TYPE_GAMMA    uint16 = 5
)

type SVGComponentTransferFunctionElement struct {
	SVGElement
}

func (this SVGComponentTransferFunctionElement) Once() SVGComponentTransferFunctionElement {
	this.Ref().Once()
	return this
}

func (this SVGComponentTransferFunctionElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGComponentTransferFunctionElement) FromRef(ref js.Ref) SVGComponentTransferFunctionElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGComponentTransferFunctionElement) Free() {
	this.Ref().Free()
}

// Type returns the value of property "SVGComponentTransferFunctionElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Type() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TableValues returns the value of property "SVGComponentTransferFunctionElement.tableValues".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) TableValues() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementTableValues(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Slope returns the value of property "SVGComponentTransferFunctionElement.slope".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Slope() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementSlope(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Intercept returns the value of property "SVGComponentTransferFunctionElement.intercept".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Intercept() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementIntercept(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Amplitude returns the value of property "SVGComponentTransferFunctionElement.amplitude".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Amplitude() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementAmplitude(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Exponent returns the value of property "SVGComponentTransferFunctionElement.exponent".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Exponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementExponent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Offset returns the value of property "SVGComponentTransferFunctionElement.offset".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Offset() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGDefsElement struct {
	SVGGraphicsElement
}

func (this SVGDefsElement) Once() SVGDefsElement {
	this.Ref().Once()
	return this
}

func (this SVGDefsElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGDefsElement) FromRef(ref js.Ref) SVGDefsElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGDefsElement) Free() {
	this.Ref().Free()
}

type SVGDescElement struct {
	SVGElement
}

func (this SVGDescElement) Once() SVGDescElement {
	this.Ref().Once()
	return this
}

func (this SVGDescElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGDescElement) FromRef(ref js.Ref) SVGDescElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGDescElement) Free() {
	this.Ref().Free()
}

type SVGDiscardElement struct {
	SVGAnimationElement
}

func (this SVGDiscardElement) Once() SVGDiscardElement {
	this.Ref().Once()
	return this
}

func (this SVGDiscardElement) Ref() js.Ref {
	return this.SVGAnimationElement.Ref()
}

func (this SVGDiscardElement) FromRef(ref js.Ref) SVGDiscardElement {
	this.SVGAnimationElement = this.SVGAnimationElement.FromRef(ref)
	return this
}

func (this SVGDiscardElement) Free() {
	this.Ref().Free()
}

type SVGEllipseElement struct {
	SVGGeometryElement
}

func (this SVGEllipseElement) Once() SVGEllipseElement {
	this.Ref().Once()
	return this
}

func (this SVGEllipseElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGEllipseElement) FromRef(ref js.Ref) SVGEllipseElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGEllipseElement) Free() {
	this.Ref().Free()
}

// Cx returns the value of property "SVGEllipseElement.cx".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Cx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementCx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cy returns the value of property "SVGEllipseElement.cy".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Cy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementCy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rx returns the value of property "SVGEllipseElement.rx".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Rx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementRx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ry returns the value of property "SVGEllipseElement.ry".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Ry() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementRy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFEBlendElement_SVG_FEBLEND_MODE_UNKNOWN     uint16 = 0
	SVGFEBlendElement_SVG_FEBLEND_MODE_NORMAL      uint16 = 1
	SVGFEBlendElement_SVG_FEBLEND_MODE_MULTIPLY    uint16 = 2
	SVGFEBlendElement_SVG_FEBLEND_MODE_SCREEN      uint16 = 3
	SVGFEBlendElement_SVG_FEBLEND_MODE_DARKEN      uint16 = 4
	SVGFEBlendElement_SVG_FEBLEND_MODE_LIGHTEN     uint16 = 5
	SVGFEBlendElement_SVG_FEBLEND_MODE_OVERLAY     uint16 = 6
	SVGFEBlendElement_SVG_FEBLEND_MODE_COLOR_DODGE uint16 = 7
	SVGFEBlendElement_SVG_FEBLEND_MODE_COLOR_BURN  uint16 = 8
	SVGFEBlendElement_SVG_FEBLEND_MODE_HARD_LIGHT  uint16 = 9
	SVGFEBlendElement_SVG_FEBLEND_MODE_SOFT_LIGHT  uint16 = 10
	SVGFEBlendElement_SVG_FEBLEND_MODE_DIFFERENCE  uint16 = 11
	SVGFEBlendElement_SVG_FEBLEND_MODE_EXCLUSION   uint16 = 12
	SVGFEBlendElement_SVG_FEBLEND_MODE_HUE         uint16 = 13
	SVGFEBlendElement_SVG_FEBLEND_MODE_SATURATION  uint16 = 14
	SVGFEBlendElement_SVG_FEBLEND_MODE_COLOR       uint16 = 15
	SVGFEBlendElement_SVG_FEBLEND_MODE_LUMINOSITY  uint16 = 16
)

type SVGFEBlendElement struct {
	SVGElement
}

func (this SVGFEBlendElement) Once() SVGFEBlendElement {
	this.Ref().Once()
	return this
}

func (this SVGFEBlendElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEBlendElement) FromRef(ref js.Ref) SVGFEBlendElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEBlendElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEBlendElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFEBlendElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementIn2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "SVGFEBlendElement.mode".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Mode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEBlendElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEBlendElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEBlendElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEBlendElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEBlendElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFEColorMatrixElement_SVG_FECOLORMATRIX_TYPE_UNKNOWN          uint16 = 0
	SVGFEColorMatrixElement_SVG_FECOLORMATRIX_TYPE_MATRIX           uint16 = 1
	SVGFEColorMatrixElement_SVG_FECOLORMATRIX_TYPE_SATURATE         uint16 = 2
	SVGFEColorMatrixElement_SVG_FECOLORMATRIX_TYPE_HUEROTATE        uint16 = 3
	SVGFEColorMatrixElement_SVG_FECOLORMATRIX_TYPE_LUMINANCETOALPHA uint16 = 4
)

type SVGFEColorMatrixElement struct {
	SVGElement
}

func (this SVGFEColorMatrixElement) Once() SVGFEColorMatrixElement {
	this.Ref().Once()
	return this
}

func (this SVGFEColorMatrixElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEColorMatrixElement) FromRef(ref js.Ref) SVGFEColorMatrixElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEColorMatrixElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEColorMatrixElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "SVGFEColorMatrixElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Type() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Values returns the value of property "SVGFEColorMatrixElement.values".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Values() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementValues(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEColorMatrixElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEColorMatrixElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEColorMatrixElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEColorMatrixElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEColorMatrixElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEComponentTransferElement struct {
	SVGElement
}

func (this SVGFEComponentTransferElement) Once() SVGFEComponentTransferElement {
	this.Ref().Once()
	return this
}

func (this SVGFEComponentTransferElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEComponentTransferElement) FromRef(ref js.Ref) SVGFEComponentTransferElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEComponentTransferElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEComponentTransferElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEComponentTransferElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEComponentTransferElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEComponentTransferElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEComponentTransferElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEComponentTransferElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_UNKNOWN    uint16 = 0
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_OVER       uint16 = 1
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_IN         uint16 = 2
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_OUT        uint16 = 3
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_ATOP       uint16 = 4
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_XOR        uint16 = 5
	SVGFECompositeElement_SVG_FECOMPOSITE_OPERATOR_ARITHMETIC uint16 = 6
)

type SVGFECompositeElement struct {
	SVGElement
}

func (this SVGFECompositeElement) Once() SVGFECompositeElement {
	this.Ref().Once()
	return this
}

func (this SVGFECompositeElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFECompositeElement) FromRef(ref js.Ref) SVGFECompositeElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFECompositeElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFECompositeElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFECompositeElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementIn2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Operator returns the value of property "SVGFECompositeElement.operator".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Operator() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementOperator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// K1 returns the value of property "SVGFECompositeElement.k1".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K1() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// K2 returns the value of property "SVGFECompositeElement.k2".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K2() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// K3 returns the value of property "SVGFECompositeElement.k3".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K3() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK3(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// K4 returns the value of property "SVGFECompositeElement.k4".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K4() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK4(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFECompositeElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFECompositeElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFECompositeElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFECompositeElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFECompositeElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFEConvolveMatrixElement_SVG_EDGEMODE_UNKNOWN   uint16 = 0
	SVGFEConvolveMatrixElement_SVG_EDGEMODE_DUPLICATE uint16 = 1
	SVGFEConvolveMatrixElement_SVG_EDGEMODE_WRAP      uint16 = 2
	SVGFEConvolveMatrixElement_SVG_EDGEMODE_NONE      uint16 = 3
)

type SVGFEConvolveMatrixElement struct {
	SVGElement
}

func (this SVGFEConvolveMatrixElement) Once() SVGFEConvolveMatrixElement {
	this.Ref().Once()
	return this
}

func (this SVGFEConvolveMatrixElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEConvolveMatrixElement) FromRef(ref js.Ref) SVGFEConvolveMatrixElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEConvolveMatrixElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEConvolveMatrixElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OrderX returns the value of property "SVGFEConvolveMatrixElement.orderX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderX() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementOrderX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OrderY returns the value of property "SVGFEConvolveMatrixElement.orderY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderY() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementOrderY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelMatrix returns the value of property "SVGFEConvolveMatrixElement.kernelMatrix".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelMatrix() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelMatrix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Divisor returns the value of property "SVGFEConvolveMatrixElement.divisor".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Divisor() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementDivisor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Bias returns the value of property "SVGFEConvolveMatrixElement.bias".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Bias() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementBias(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TargetX returns the value of property "SVGFEConvolveMatrixElement.targetX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetX() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementTargetX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TargetY returns the value of property "SVGFEConvolveMatrixElement.targetY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetY() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementTargetY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EdgeMode returns the value of property "SVGFEConvolveMatrixElement.edgeMode".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) EdgeMode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementEdgeMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAlpha returns the value of property "SVGFEConvolveMatrixElement.preserveAlpha".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) PreserveAlpha() (ret SVGAnimatedBoolean, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementPreserveAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEConvolveMatrixElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEConvolveMatrixElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEConvolveMatrixElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEConvolveMatrixElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEConvolveMatrixElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEDiffuseLightingElement struct {
	SVGElement
}

func (this SVGFEDiffuseLightingElement) Once() SVGFEDiffuseLightingElement {
	this.Ref().Once()
	return this
}

func (this SVGFEDiffuseLightingElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEDiffuseLightingElement) FromRef(ref js.Ref) SVGFEDiffuseLightingElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEDiffuseLightingElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEDiffuseLightingElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SurfaceScale returns the value of property "SVGFEDiffuseLightingElement.surfaceScale".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) SurfaceScale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementSurfaceScale(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DiffuseConstant returns the value of property "SVGFEDiffuseLightingElement.diffuseConstant".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) DiffuseConstant() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementDiffuseConstant(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDiffuseLightingElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDiffuseLightingElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDiffuseLightingElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDiffuseLightingElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDiffuseLightingElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFEDisplacementMapElement_SVG_CHANNEL_UNKNOWN uint16 = 0
	SVGFEDisplacementMapElement_SVG_CHANNEL_R       uint16 = 1
	SVGFEDisplacementMapElement_SVG_CHANNEL_G       uint16 = 2
	SVGFEDisplacementMapElement_SVG_CHANNEL_B       uint16 = 3
	SVGFEDisplacementMapElement_SVG_CHANNEL_A       uint16 = 4
)

type SVGFEDisplacementMapElement struct {
	SVGElement
}

func (this SVGFEDisplacementMapElement) Once() SVGFEDisplacementMapElement {
	this.Ref().Once()
	return this
}

func (this SVGFEDisplacementMapElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEDisplacementMapElement) FromRef(ref js.Ref) SVGFEDisplacementMapElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEDisplacementMapElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEDisplacementMapElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFEDisplacementMapElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementIn2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scale returns the value of property "SVGFEDisplacementMapElement.scale".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Scale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementScale(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// XChannelSelector returns the value of property "SVGFEDisplacementMapElement.xChannelSelector".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) XChannelSelector() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementXChannelSelector(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// YChannelSelector returns the value of property "SVGFEDisplacementMapElement.yChannelSelector".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) YChannelSelector() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementYChannelSelector(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDisplacementMapElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDisplacementMapElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDisplacementMapElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDisplacementMapElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDisplacementMapElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEDistantLightElement struct {
	SVGElement
}

func (this SVGFEDistantLightElement) Once() SVGFEDistantLightElement {
	this.Ref().Once()
	return this
}

func (this SVGFEDistantLightElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEDistantLightElement) FromRef(ref js.Ref) SVGFEDistantLightElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEDistantLightElement) Free() {
	this.Ref().Free()
}

// Azimuth returns the value of property "SVGFEDistantLightElement.azimuth".
//
// It returns ok=false if there is no such property.
func (this SVGFEDistantLightElement) Azimuth() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDistantLightElementAzimuth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Elevation returns the value of property "SVGFEDistantLightElement.elevation".
//
// It returns ok=false if there is no such property.
func (this SVGFEDistantLightElement) Elevation() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDistantLightElementElevation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEDropShadowElement struct {
	SVGElement
}

func (this SVGFEDropShadowElement) Once() SVGFEDropShadowElement {
	this.Ref().Once()
	return this
}

func (this SVGFEDropShadowElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEDropShadowElement) FromRef(ref js.Ref) SVGFEDropShadowElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEDropShadowElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEDropShadowElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dx returns the value of property "SVGFEDropShadowElement.dx".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Dx() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementDx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dy returns the value of property "SVGFEDropShadowElement.dy".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Dy() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementDy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StdDeviationX returns the value of property "SVGFEDropShadowElement.stdDeviationX".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementStdDeviationX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StdDeviationY returns the value of property "SVGFEDropShadowElement.stdDeviationY".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementStdDeviationY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDropShadowElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDropShadowElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDropShadowElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDropShadowElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDropShadowElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetStdDeviation returns true if the method "SVGFEDropShadowElement.setStdDeviation" exists.
func (this SVGFEDropShadowElement) HasSetStdDeviation() bool {
	return js.True == bindings.HasSVGFEDropShadowElementSetStdDeviation(
		this.Ref(),
	)
}

// SetStdDeviationFunc returns the method "SVGFEDropShadowElement.setStdDeviation".
func (this SVGFEDropShadowElement) SetStdDeviationFunc() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	return fn.FromRef(
		bindings.SVGFEDropShadowElementSetStdDeviationFunc(
			this.Ref(),
		),
	)
}

// SetStdDeviation calls the method "SVGFEDropShadowElement.setStdDeviation".
func (this SVGFEDropShadowElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void) {
	bindings.CallSVGFEDropShadowElementSetStdDeviation(
		this.Ref(), js.Pointer(&ret),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

// TrySetStdDeviation calls the method "SVGFEDropShadowElement.setStdDeviation"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGFEDropShadowElement) TrySetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGFEDropShadowElementSetStdDeviation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

type SVGFEFloodElement struct {
	SVGElement
}

func (this SVGFEFloodElement) Once() SVGFEFloodElement {
	this.Ref().Once()
	return this
}

func (this SVGFEFloodElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEFloodElement) FromRef(ref js.Ref) SVGFEFloodElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEFloodElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGFEFloodElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEFloodElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEFloodElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEFloodElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEFloodElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEFuncAElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncAElement) Once() SVGFEFuncAElement {
	this.Ref().Once()
	return this
}

func (this SVGFEFuncAElement) Ref() js.Ref {
	return this.SVGComponentTransferFunctionElement.Ref()
}

func (this SVGFEFuncAElement) FromRef(ref js.Ref) SVGFEFuncAElement {
	this.SVGComponentTransferFunctionElement = this.SVGComponentTransferFunctionElement.FromRef(ref)
	return this
}

func (this SVGFEFuncAElement) Free() {
	this.Ref().Free()
}

type SVGFEFuncBElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncBElement) Once() SVGFEFuncBElement {
	this.Ref().Once()
	return this
}

func (this SVGFEFuncBElement) Ref() js.Ref {
	return this.SVGComponentTransferFunctionElement.Ref()
}

func (this SVGFEFuncBElement) FromRef(ref js.Ref) SVGFEFuncBElement {
	this.SVGComponentTransferFunctionElement = this.SVGComponentTransferFunctionElement.FromRef(ref)
	return this
}

func (this SVGFEFuncBElement) Free() {
	this.Ref().Free()
}

type SVGFEFuncGElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncGElement) Once() SVGFEFuncGElement {
	this.Ref().Once()
	return this
}

func (this SVGFEFuncGElement) Ref() js.Ref {
	return this.SVGComponentTransferFunctionElement.Ref()
}

func (this SVGFEFuncGElement) FromRef(ref js.Ref) SVGFEFuncGElement {
	this.SVGComponentTransferFunctionElement = this.SVGComponentTransferFunctionElement.FromRef(ref)
	return this
}

func (this SVGFEFuncGElement) Free() {
	this.Ref().Free()
}

type SVGFEFuncRElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncRElement) Once() SVGFEFuncRElement {
	this.Ref().Once()
	return this
}

func (this SVGFEFuncRElement) Ref() js.Ref {
	return this.SVGComponentTransferFunctionElement.Ref()
}

func (this SVGFEFuncRElement) FromRef(ref js.Ref) SVGFEFuncRElement {
	this.SVGComponentTransferFunctionElement = this.SVGComponentTransferFunctionElement.FromRef(ref)
	return this
}

func (this SVGFEFuncRElement) Free() {
	this.Ref().Free()
}

const (
	SVGFEGaussianBlurElement_SVG_EDGEMODE_UNKNOWN   uint16 = 0
	SVGFEGaussianBlurElement_SVG_EDGEMODE_DUPLICATE uint16 = 1
	SVGFEGaussianBlurElement_SVG_EDGEMODE_WRAP      uint16 = 2
	SVGFEGaussianBlurElement_SVG_EDGEMODE_NONE      uint16 = 3
)

type SVGFEGaussianBlurElement struct {
	SVGElement
}

func (this SVGFEGaussianBlurElement) Once() SVGFEGaussianBlurElement {
	this.Ref().Once()
	return this
}

func (this SVGFEGaussianBlurElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEGaussianBlurElement) FromRef(ref js.Ref) SVGFEGaussianBlurElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEGaussianBlurElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEGaussianBlurElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StdDeviationX returns the value of property "SVGFEGaussianBlurElement.stdDeviationX".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementStdDeviationX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StdDeviationY returns the value of property "SVGFEGaussianBlurElement.stdDeviationY".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementStdDeviationY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EdgeMode returns the value of property "SVGFEGaussianBlurElement.edgeMode".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) EdgeMode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementEdgeMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEGaussianBlurElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEGaussianBlurElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEGaussianBlurElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEGaussianBlurElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEGaussianBlurElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetStdDeviation returns true if the method "SVGFEGaussianBlurElement.setStdDeviation" exists.
func (this SVGFEGaussianBlurElement) HasSetStdDeviation() bool {
	return js.True == bindings.HasSVGFEGaussianBlurElementSetStdDeviation(
		this.Ref(),
	)
}

// SetStdDeviationFunc returns the method "SVGFEGaussianBlurElement.setStdDeviation".
func (this SVGFEGaussianBlurElement) SetStdDeviationFunc() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	return fn.FromRef(
		bindings.SVGFEGaussianBlurElementSetStdDeviationFunc(
			this.Ref(),
		),
	)
}

// SetStdDeviation calls the method "SVGFEGaussianBlurElement.setStdDeviation".
func (this SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void) {
	bindings.CallSVGFEGaussianBlurElementSetStdDeviation(
		this.Ref(), js.Pointer(&ret),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

// TrySetStdDeviation calls the method "SVGFEGaussianBlurElement.setStdDeviation"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGFEGaussianBlurElement) TrySetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGFEGaussianBlurElementSetStdDeviation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

type SVGFEImageElement struct {
	SVGElement
}

func (this SVGFEImageElement) Once() SVGFEImageElement {
	this.Ref().Once()
	return this
}

func (this SVGFEImageElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEImageElement) FromRef(ref js.Ref) SVGFEImageElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEImageElement) Free() {
	this.Ref().Free()
}

// PreserveAspectRatio returns the value of property "SVGFEImageElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CrossOrigin returns the value of property "SVGFEImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) CrossOrigin() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEImageElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEImageElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEImageElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEImageElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEImageElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGFEImageElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEMergeElement struct {
	SVGElement
}

func (this SVGFEMergeElement) Once() SVGFEMergeElement {
	this.Ref().Once()
	return this
}

func (this SVGFEMergeElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEMergeElement) FromRef(ref js.Ref) SVGFEMergeElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEMergeElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGFEMergeElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEMergeElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEMergeElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEMergeElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEMergeElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEMergeNodeElement struct {
	SVGElement
}

func (this SVGFEMergeNodeElement) Once() SVGFEMergeNodeElement {
	this.Ref().Once()
	return this
}

func (this SVGFEMergeNodeElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEMergeNodeElement) FromRef(ref js.Ref) SVGFEMergeNodeElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEMergeNodeElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEMergeNodeElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeNodeElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeNodeElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	SVGFEMorphologyElement_SVG_MORPHOLOGY_OPERATOR_UNKNOWN uint16 = 0
	SVGFEMorphologyElement_SVG_MORPHOLOGY_OPERATOR_ERODE   uint16 = 1
	SVGFEMorphologyElement_SVG_MORPHOLOGY_OPERATOR_DILATE  uint16 = 2
)

type SVGFEMorphologyElement struct {
	SVGElement
}

func (this SVGFEMorphologyElement) Once() SVGFEMorphologyElement {
	this.Ref().Once()
	return this
}

func (this SVGFEMorphologyElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEMorphologyElement) FromRef(ref js.Ref) SVGFEMorphologyElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEMorphologyElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEMorphologyElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Operator returns the value of property "SVGFEMorphologyElement.operator".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Operator() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementOperator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RadiusX returns the value of property "SVGFEMorphologyElement.radiusX".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) RadiusX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementRadiusX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RadiusY returns the value of property "SVGFEMorphologyElement.radiusY".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) RadiusY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementRadiusY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEMorphologyElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEMorphologyElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEMorphologyElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEMorphologyElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEMorphologyElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEOffsetElement struct {
	SVGElement
}

func (this SVGFEOffsetElement) Once() SVGFEOffsetElement {
	this.Ref().Once()
	return this
}

func (this SVGFEOffsetElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEOffsetElement) FromRef(ref js.Ref) SVGFEOffsetElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEOffsetElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFEOffsetElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dx returns the value of property "SVGFEOffsetElement.dx".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Dx() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementDx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dy returns the value of property "SVGFEOffsetElement.dy".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Dy() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementDy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEOffsetElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEOffsetElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEOffsetElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEOffsetElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEOffsetElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFEPointLightElement struct {
	SVGElement
}

func (this SVGFEPointLightElement) Once() SVGFEPointLightElement {
	this.Ref().Once()
	return this
}

func (this SVGFEPointLightElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFEPointLightElement) FromRef(ref js.Ref) SVGFEPointLightElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFEPointLightElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGFEPointLightElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) X() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEPointLightElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) Y() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "SVGFEPointLightElement.z".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) Z() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFESpecularLightingElement struct {
	SVGElement
}

func (this SVGFESpecularLightingElement) Once() SVGFESpecularLightingElement {
	this.Ref().Once()
	return this
}

func (this SVGFESpecularLightingElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFESpecularLightingElement) FromRef(ref js.Ref) SVGFESpecularLightingElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFESpecularLightingElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFESpecularLightingElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SurfaceScale returns the value of property "SVGFESpecularLightingElement.surfaceScale".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SurfaceScale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSurfaceScale(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SpecularConstant returns the value of property "SVGFESpecularLightingElement.specularConstant".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularConstant() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSpecularConstant(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SpecularExponent returns the value of property "SVGFESpecularLightingElement.specularExponent".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularExponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSpecularExponent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFESpecularLightingElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFESpecularLightingElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFESpecularLightingElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFESpecularLightingElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFESpecularLightingElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SVGFESpotLightElement struct {
	SVGElement
}

func (this SVGFESpotLightElement) Once() SVGFESpotLightElement {
	this.Ref().Once()
	return this
}

func (this SVGFESpotLightElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFESpotLightElement) FromRef(ref js.Ref) SVGFESpotLightElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFESpotLightElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGFESpotLightElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) X() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFESpotLightElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) Y() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "SVGFESpotLightElement.z".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) Z() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointsAtX returns the value of property "SVGFESpotLightElement.pointsAtX".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointsAtY returns the value of property "SVGFESpotLightElement.pointsAtY".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointsAtZ returns the value of property "SVGFESpotLightElement.pointsAtZ".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtZ() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SpecularExponent returns the value of property "SVGFESpotLightElement.specularExponent".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) SpecularExponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementSpecularExponent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LimitingConeAngle returns the value of property "SVGFESpotLightElement.limitingConeAngle".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) LimitingConeAngle() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementLimitingConeAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

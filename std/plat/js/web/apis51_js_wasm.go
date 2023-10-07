// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type SVGAnimatedEnumeration struct {
	ref js.Ref
}

func (this SVGAnimatedEnumeration) Once() SVGAnimatedEnumeration {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedEnumeration.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedEnumeration) BaseVal() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedEnumerationBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedEnumeration.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedEnumeration) SetBaseVal(val uint16) bool {
	return js.True == bindings.SetSVGAnimatedEnumerationBaseVal(
		this.ref,
		uint32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedEnumeration.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedEnumeration) AnimVal() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedEnumerationAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGAnimatedInteger struct {
	ref js.Ref
}

func (this SVGAnimatedInteger) Once() SVGAnimatedInteger {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedInteger.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedInteger) BaseVal() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedIntegerBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedInteger.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedInteger) SetBaseVal(val int32) bool {
	return js.True == bindings.SetSVGAnimatedIntegerBaseVal(
		this.ref,
		int32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedInteger.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedInteger) AnimVal() (ret int32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedIntegerAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGLengthList struct {
	ref js.Ref
}

func (this SVGLengthList) Once() SVGLengthList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SVGLengthList.length".
//
// It returns ok=false if there is no such property.
func (this SVGLengthList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGLengthListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGLengthList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGLengthList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGLengthListNumberOfItems(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClear returns true if the method "SVGLengthList.clear" exists.
func (this SVGLengthList) HasFuncClear() bool {
	return js.True == bindings.HasFuncSVGLengthListClear(
		this.ref,
	)
}

// FuncClear returns the method "SVGLengthList.clear".
func (this SVGLengthList) FuncClear() (fn js.Func[func()]) {
	bindings.FuncSVGLengthListClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "SVGLengthList.clear".
func (this SVGLengthList) Clear() (ret js.Void) {
	bindings.CallSVGLengthListClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGLengthList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitialize returns true if the method "SVGLengthList.initialize" exists.
func (this SVGLengthList) HasFuncInitialize() bool {
	return js.True == bindings.HasFuncSVGLengthListInitialize(
		this.ref,
	)
}

// FuncInitialize returns the method "SVGLengthList.initialize".
func (this SVGLengthList) FuncInitialize() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	bindings.FuncSVGLengthListInitialize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Initialize calls the method "SVGLengthList.initialize".
func (this SVGLengthList) Initialize(newItem SVGLength) (ret SVGLength) {
	bindings.CallSVGLengthListInitialize(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGLengthList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryInitialize(newItem SVGLength) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListInitialize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncGetItem returns true if the method "SVGLengthList.getItem" exists.
func (this SVGLengthList) HasFuncGetItem() bool {
	return js.True == bindings.HasFuncSVGLengthListGetItem(
		this.ref,
	)
}

// FuncGetItem returns the method "SVGLengthList.getItem".
func (this SVGLengthList) FuncGetItem() (fn js.Func[func(index uint32) SVGLength]) {
	bindings.FuncSVGLengthListGetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetItem calls the method "SVGLengthList.getItem".
func (this SVGLengthList) GetItem(index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListGetItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGLengthList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryGetItem(index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListGetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncInsertItemBefore returns true if the method "SVGLengthList.insertItemBefore" exists.
func (this SVGLengthList) HasFuncInsertItemBefore() bool {
	return js.True == bindings.HasFuncSVGLengthListInsertItemBefore(
		this.ref,
	)
}

// FuncInsertItemBefore returns the method "SVGLengthList.insertItemBefore".
func (this SVGLengthList) FuncInsertItemBefore() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	bindings.FuncSVGLengthListInsertItemBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertItemBefore calls the method "SVGLengthList.insertItemBefore".
func (this SVGLengthList) InsertItemBefore(newItem SVGLength, index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListInsertItemBefore(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGLengthList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryInsertItemBefore(newItem SVGLength, index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListInsertItemBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncReplaceItem returns true if the method "SVGLengthList.replaceItem" exists.
func (this SVGLengthList) HasFuncReplaceItem() bool {
	return js.True == bindings.HasFuncSVGLengthListReplaceItem(
		this.ref,
	)
}

// FuncReplaceItem returns the method "SVGLengthList.replaceItem".
func (this SVGLengthList) FuncReplaceItem() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	bindings.FuncSVGLengthListReplaceItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceItem calls the method "SVGLengthList.replaceItem".
func (this SVGLengthList) ReplaceItem(newItem SVGLength, index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListReplaceItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGLengthList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryReplaceItem(newItem SVGLength, index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListReplaceItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncRemoveItem returns true if the method "SVGLengthList.removeItem" exists.
func (this SVGLengthList) HasFuncRemoveItem() bool {
	return js.True == bindings.HasFuncSVGLengthListRemoveItem(
		this.ref,
	)
}

// FuncRemoveItem returns the method "SVGLengthList.removeItem".
func (this SVGLengthList) FuncRemoveItem() (fn js.Func[func(index uint32) SVGLength]) {
	bindings.FuncSVGLengthListRemoveItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveItem calls the method "SVGLengthList.removeItem".
func (this SVGLengthList) RemoveItem(index uint32) (ret SVGLength) {
	bindings.CallSVGLengthListRemoveItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGLengthList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryRemoveItem(index uint32) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListRemoveItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendItem returns true if the method "SVGLengthList.appendItem" exists.
func (this SVGLengthList) HasFuncAppendItem() bool {
	return js.True == bindings.HasFuncSVGLengthListAppendItem(
		this.ref,
	)
}

// FuncAppendItem returns the method "SVGLengthList.appendItem".
func (this SVGLengthList) FuncAppendItem() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	bindings.FuncSVGLengthListAppendItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendItem calls the method "SVGLengthList.appendItem".
func (this SVGLengthList) AppendItem(newItem SVGLength) (ret SVGLength) {
	bindings.CallSVGLengthListAppendItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGLengthList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TryAppendItem(newItem SVGLength) (ret SVGLength, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListAppendItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "SVGLengthList." exists.
func (this SVGLengthList) HasFuncSet() bool {
	return js.True == bindings.HasFuncSVGLengthListSet(
		this.ref,
	)
}

// FuncSet returns the method "SVGLengthList.".
func (this SVGLengthList) FuncSet() (fn js.Func[func(index uint32, newItem SVGLength)]) {
	bindings.FuncSVGLengthListSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "SVGLengthList.".
func (this SVGLengthList) Set(index uint32, newItem SVGLength) (ret js.Void) {
	bindings.CallSVGLengthListSet(
		this.ref, js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGLengthList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGLengthList) TrySet(index uint32, newItem SVGLength) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGLengthListSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimatedLengthList struct {
	ref js.Ref
}

func (this SVGAnimatedLengthList) Once() SVGAnimatedLengthList {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedLengthList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLengthList) BaseVal() (ret SVGLengthList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthListBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedLengthList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedLengthList) AnimVal() (ret SVGLengthList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedLengthListAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGAnimatedNumber struct {
	ref js.Ref
}

func (this SVGAnimatedNumber) Once() SVGAnimatedNumber {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedNumber.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumber) BaseVal() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBaseVal sets the value of property "SVGAnimatedNumber.baseVal" to val.
//
// It returns false if the property cannot be set.
func (this SVGAnimatedNumber) SetBaseVal(val float32) bool {
	return js.True == bindings.SetSVGAnimatedNumberBaseVal(
		this.ref,
		float32(val),
	)
}

// AnimVal returns the value of property "SVGAnimatedNumber.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumber) AnimVal() (ret float32, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGNumberList struct {
	ref js.Ref
}

func (this SVGNumberList) Once() SVGNumberList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SVGNumberList.length".
//
// It returns ok=false if there is no such property.
func (this SVGNumberList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGNumberListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGNumberList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGNumberList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGNumberListNumberOfItems(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClear returns true if the method "SVGNumberList.clear" exists.
func (this SVGNumberList) HasFuncClear() bool {
	return js.True == bindings.HasFuncSVGNumberListClear(
		this.ref,
	)
}

// FuncClear returns the method "SVGNumberList.clear".
func (this SVGNumberList) FuncClear() (fn js.Func[func()]) {
	bindings.FuncSVGNumberListClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "SVGNumberList.clear".
func (this SVGNumberList) Clear() (ret js.Void) {
	bindings.CallSVGNumberListClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGNumberList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitialize returns true if the method "SVGNumberList.initialize" exists.
func (this SVGNumberList) HasFuncInitialize() bool {
	return js.True == bindings.HasFuncSVGNumberListInitialize(
		this.ref,
	)
}

// FuncInitialize returns the method "SVGNumberList.initialize".
func (this SVGNumberList) FuncInitialize() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	bindings.FuncSVGNumberListInitialize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Initialize calls the method "SVGNumberList.initialize".
func (this SVGNumberList) Initialize(newItem SVGNumber) (ret SVGNumber) {
	bindings.CallSVGNumberListInitialize(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGNumberList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryInitialize(newItem SVGNumber) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListInitialize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncGetItem returns true if the method "SVGNumberList.getItem" exists.
func (this SVGNumberList) HasFuncGetItem() bool {
	return js.True == bindings.HasFuncSVGNumberListGetItem(
		this.ref,
	)
}

// FuncGetItem returns the method "SVGNumberList.getItem".
func (this SVGNumberList) FuncGetItem() (fn js.Func[func(index uint32) SVGNumber]) {
	bindings.FuncSVGNumberListGetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetItem calls the method "SVGNumberList.getItem".
func (this SVGNumberList) GetItem(index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListGetItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGNumberList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryGetItem(index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListGetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncInsertItemBefore returns true if the method "SVGNumberList.insertItemBefore" exists.
func (this SVGNumberList) HasFuncInsertItemBefore() bool {
	return js.True == bindings.HasFuncSVGNumberListInsertItemBefore(
		this.ref,
	)
}

// FuncInsertItemBefore returns the method "SVGNumberList.insertItemBefore".
func (this SVGNumberList) FuncInsertItemBefore() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	bindings.FuncSVGNumberListInsertItemBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertItemBefore calls the method "SVGNumberList.insertItemBefore".
func (this SVGNumberList) InsertItemBefore(newItem SVGNumber, index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListInsertItemBefore(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGNumberList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryInsertItemBefore(newItem SVGNumber, index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListInsertItemBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncReplaceItem returns true if the method "SVGNumberList.replaceItem" exists.
func (this SVGNumberList) HasFuncReplaceItem() bool {
	return js.True == bindings.HasFuncSVGNumberListReplaceItem(
		this.ref,
	)
}

// FuncReplaceItem returns the method "SVGNumberList.replaceItem".
func (this SVGNumberList) FuncReplaceItem() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	bindings.FuncSVGNumberListReplaceItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceItem calls the method "SVGNumberList.replaceItem".
func (this SVGNumberList) ReplaceItem(newItem SVGNumber, index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListReplaceItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGNumberList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryReplaceItem(newItem SVGNumber, index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListReplaceItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncRemoveItem returns true if the method "SVGNumberList.removeItem" exists.
func (this SVGNumberList) HasFuncRemoveItem() bool {
	return js.True == bindings.HasFuncSVGNumberListRemoveItem(
		this.ref,
	)
}

// FuncRemoveItem returns the method "SVGNumberList.removeItem".
func (this SVGNumberList) FuncRemoveItem() (fn js.Func[func(index uint32) SVGNumber]) {
	bindings.FuncSVGNumberListRemoveItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveItem calls the method "SVGNumberList.removeItem".
func (this SVGNumberList) RemoveItem(index uint32) (ret SVGNumber) {
	bindings.CallSVGNumberListRemoveItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGNumberList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryRemoveItem(index uint32) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListRemoveItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendItem returns true if the method "SVGNumberList.appendItem" exists.
func (this SVGNumberList) HasFuncAppendItem() bool {
	return js.True == bindings.HasFuncSVGNumberListAppendItem(
		this.ref,
	)
}

// FuncAppendItem returns the method "SVGNumberList.appendItem".
func (this SVGNumberList) FuncAppendItem() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	bindings.FuncSVGNumberListAppendItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendItem calls the method "SVGNumberList.appendItem".
func (this SVGNumberList) AppendItem(newItem SVGNumber) (ret SVGNumber) {
	bindings.CallSVGNumberListAppendItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGNumberList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TryAppendItem(newItem SVGNumber) (ret SVGNumber, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListAppendItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "SVGNumberList." exists.
func (this SVGNumberList) HasFuncSet() bool {
	return js.True == bindings.HasFuncSVGNumberListSet(
		this.ref,
	)
}

// FuncSet returns the method "SVGNumberList.".
func (this SVGNumberList) FuncSet() (fn js.Func[func(index uint32, newItem SVGNumber)]) {
	bindings.FuncSVGNumberListSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "SVGNumberList.".
func (this SVGNumberList) Set(index uint32, newItem SVGNumber) (ret js.Void) {
	bindings.CallSVGNumberListSet(
		this.ref, js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGNumberList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGNumberList) TrySet(index uint32, newItem SVGNumber) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGNumberListSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimatedNumberList struct {
	ref js.Ref
}

func (this SVGAnimatedNumberList) Once() SVGAnimatedNumberList {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedNumberList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumberList) BaseVal() (ret SVGNumberList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberListBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedNumberList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedNumberList) AnimVal() (ret SVGNumberList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedNumberListAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGTransformList struct {
	ref js.Ref
}

func (this SVGTransformList) Once() SVGTransformList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SVGTransformList.length".
//
// It returns ok=false if there is no such property.
func (this SVGTransformList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGTransformListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGTransformList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGTransformList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGTransformListNumberOfItems(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClear returns true if the method "SVGTransformList.clear" exists.
func (this SVGTransformList) HasFuncClear() bool {
	return js.True == bindings.HasFuncSVGTransformListClear(
		this.ref,
	)
}

// FuncClear returns the method "SVGTransformList.clear".
func (this SVGTransformList) FuncClear() (fn js.Func[func()]) {
	bindings.FuncSVGTransformListClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "SVGTransformList.clear".
func (this SVGTransformList) Clear() (ret js.Void) {
	bindings.CallSVGTransformListClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGTransformList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitialize returns true if the method "SVGTransformList.initialize" exists.
func (this SVGTransformList) HasFuncInitialize() bool {
	return js.True == bindings.HasFuncSVGTransformListInitialize(
		this.ref,
	)
}

// FuncInitialize returns the method "SVGTransformList.initialize".
func (this SVGTransformList) FuncInitialize() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	bindings.FuncSVGTransformListInitialize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Initialize calls the method "SVGTransformList.initialize".
func (this SVGTransformList) Initialize(newItem SVGTransform) (ret SVGTransform) {
	bindings.CallSVGTransformListInitialize(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGTransformList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryInitialize(newItem SVGTransform) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListInitialize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncGetItem returns true if the method "SVGTransformList.getItem" exists.
func (this SVGTransformList) HasFuncGetItem() bool {
	return js.True == bindings.HasFuncSVGTransformListGetItem(
		this.ref,
	)
}

// FuncGetItem returns the method "SVGTransformList.getItem".
func (this SVGTransformList) FuncGetItem() (fn js.Func[func(index uint32) SVGTransform]) {
	bindings.FuncSVGTransformListGetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetItem calls the method "SVGTransformList.getItem".
func (this SVGTransformList) GetItem(index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListGetItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGTransformList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryGetItem(index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListGetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncInsertItemBefore returns true if the method "SVGTransformList.insertItemBefore" exists.
func (this SVGTransformList) HasFuncInsertItemBefore() bool {
	return js.True == bindings.HasFuncSVGTransformListInsertItemBefore(
		this.ref,
	)
}

// FuncInsertItemBefore returns the method "SVGTransformList.insertItemBefore".
func (this SVGTransformList) FuncInsertItemBefore() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	bindings.FuncSVGTransformListInsertItemBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertItemBefore calls the method "SVGTransformList.insertItemBefore".
func (this SVGTransformList) InsertItemBefore(newItem SVGTransform, index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListInsertItemBefore(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGTransformList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryInsertItemBefore(newItem SVGTransform, index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListInsertItemBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncReplaceItem returns true if the method "SVGTransformList.replaceItem" exists.
func (this SVGTransformList) HasFuncReplaceItem() bool {
	return js.True == bindings.HasFuncSVGTransformListReplaceItem(
		this.ref,
	)
}

// FuncReplaceItem returns the method "SVGTransformList.replaceItem".
func (this SVGTransformList) FuncReplaceItem() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	bindings.FuncSVGTransformListReplaceItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceItem calls the method "SVGTransformList.replaceItem".
func (this SVGTransformList) ReplaceItem(newItem SVGTransform, index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListReplaceItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGTransformList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryReplaceItem(newItem SVGTransform, index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListReplaceItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncRemoveItem returns true if the method "SVGTransformList.removeItem" exists.
func (this SVGTransformList) HasFuncRemoveItem() bool {
	return js.True == bindings.HasFuncSVGTransformListRemoveItem(
		this.ref,
	)
}

// FuncRemoveItem returns the method "SVGTransformList.removeItem".
func (this SVGTransformList) FuncRemoveItem() (fn js.Func[func(index uint32) SVGTransform]) {
	bindings.FuncSVGTransformListRemoveItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveItem calls the method "SVGTransformList.removeItem".
func (this SVGTransformList) RemoveItem(index uint32) (ret SVGTransform) {
	bindings.CallSVGTransformListRemoveItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGTransformList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryRemoveItem(index uint32) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListRemoveItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendItem returns true if the method "SVGTransformList.appendItem" exists.
func (this SVGTransformList) HasFuncAppendItem() bool {
	return js.True == bindings.HasFuncSVGTransformListAppendItem(
		this.ref,
	)
}

// FuncAppendItem returns the method "SVGTransformList.appendItem".
func (this SVGTransformList) FuncAppendItem() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	bindings.FuncSVGTransformListAppendItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendItem calls the method "SVGTransformList.appendItem".
func (this SVGTransformList) AppendItem(newItem SVGTransform) (ret SVGTransform) {
	bindings.CallSVGTransformListAppendItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGTransformList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryAppendItem(newItem SVGTransform) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListAppendItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "SVGTransformList." exists.
func (this SVGTransformList) HasFuncSet() bool {
	return js.True == bindings.HasFuncSVGTransformListSet(
		this.ref,
	)
}

// FuncSet returns the method "SVGTransformList.".
func (this SVGTransformList) FuncSet() (fn js.Func[func(index uint32, newItem SVGTransform)]) {
	bindings.FuncSVGTransformListSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "SVGTransformList.".
func (this SVGTransformList) Set(index uint32, newItem SVGTransform) (ret js.Void) {
	bindings.CallSVGTransformListSet(
		this.ref, js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGTransformList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TrySet(index uint32, newItem SVGTransform) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// HasFuncCreateSVGTransformFromMatrix returns true if the method "SVGTransformList.createSVGTransformFromMatrix" exists.
func (this SVGTransformList) HasFuncCreateSVGTransformFromMatrix() bool {
	return js.True == bindings.HasFuncSVGTransformListCreateSVGTransformFromMatrix(
		this.ref,
	)
}

// FuncCreateSVGTransformFromMatrix returns the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) FuncCreateSVGTransformFromMatrix() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	bindings.FuncSVGTransformListCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGTransformFromMatrix calls the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform) {
	bindings.CallSVGTransformListCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&ret),
		js.Pointer(&matrix),
	)

	return
}

// TryCreateSVGTransformFromMatrix calls the method "SVGTransformList.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryCreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListCreateSVGTransformFromMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&matrix),
	)

	return
}

// HasFuncCreateSVGTransformFromMatrix1 returns true if the method "SVGTransformList.createSVGTransformFromMatrix" exists.
func (this SVGTransformList) HasFuncCreateSVGTransformFromMatrix1() bool {
	return js.True == bindings.HasFuncSVGTransformListCreateSVGTransformFromMatrix1(
		this.ref,
	)
}

// FuncCreateSVGTransformFromMatrix1 returns the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) FuncCreateSVGTransformFromMatrix1() (fn js.Func[func() SVGTransform]) {
	bindings.FuncSVGTransformListCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSVGTransformFromMatrix1 calls the method "SVGTransformList.createSVGTransformFromMatrix".
func (this SVGTransformList) CreateSVGTransformFromMatrix1() (ret SVGTransform) {
	bindings.CallSVGTransformListCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSVGTransformFromMatrix1 calls the method "SVGTransformList.createSVGTransformFromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryCreateSVGTransformFromMatrix1() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListCreateSVGTransformFromMatrix1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConsolidate returns true if the method "SVGTransformList.consolidate" exists.
func (this SVGTransformList) HasFuncConsolidate() bool {
	return js.True == bindings.HasFuncSVGTransformListConsolidate(
		this.ref,
	)
}

// FuncConsolidate returns the method "SVGTransformList.consolidate".
func (this SVGTransformList) FuncConsolidate() (fn js.Func[func() SVGTransform]) {
	bindings.FuncSVGTransformListConsolidate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Consolidate calls the method "SVGTransformList.consolidate".
func (this SVGTransformList) Consolidate() (ret SVGTransform) {
	bindings.CallSVGTransformListConsolidate(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryConsolidate calls the method "SVGTransformList.consolidate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGTransformList) TryConsolidate() (ret SVGTransform, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTransformListConsolidate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGAnimatedTransformList struct {
	ref js.Ref
}

func (this SVGAnimatedTransformList) Once() SVGAnimatedTransformList {
	this.ref.Once()
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
	this.ref.Free()
}

// BaseVal returns the value of property "SVGAnimatedTransformList.baseVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedTransformList) BaseVal() (ret SVGTransformList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedTransformListBaseVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnimVal returns the value of property "SVGAnimatedTransformList.animVal".
//
// It returns ok=false if there is no such property.
func (this SVGAnimatedTransformList) AnimVal() (ret SVGTransformList, ok bool) {
	ok = js.True == bindings.GetSVGAnimatedTransformListAnimVal(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGStringList struct {
	ref js.Ref
}

func (this SVGStringList) Once() SVGStringList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SVGStringList.length".
//
// It returns ok=false if there is no such property.
func (this SVGStringList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGStringListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGStringList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGStringList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGStringListNumberOfItems(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClear returns true if the method "SVGStringList.clear" exists.
func (this SVGStringList) HasFuncClear() bool {
	return js.True == bindings.HasFuncSVGStringListClear(
		this.ref,
	)
}

// FuncClear returns the method "SVGStringList.clear".
func (this SVGStringList) FuncClear() (fn js.Func[func()]) {
	bindings.FuncSVGStringListClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "SVGStringList.clear".
func (this SVGStringList) Clear() (ret js.Void) {
	bindings.CallSVGStringListClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGStringList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitialize returns true if the method "SVGStringList.initialize" exists.
func (this SVGStringList) HasFuncInitialize() bool {
	return js.True == bindings.HasFuncSVGStringListInitialize(
		this.ref,
	)
}

// FuncInitialize returns the method "SVGStringList.initialize".
func (this SVGStringList) FuncInitialize() (fn js.Func[func(newItem js.String) js.String]) {
	bindings.FuncSVGStringListInitialize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Initialize calls the method "SVGStringList.initialize".
func (this SVGStringList) Initialize(newItem js.String) (ret js.String) {
	bindings.CallSVGStringListInitialize(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGStringList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryInitialize(newItem js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListInitialize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncGetItem returns true if the method "SVGStringList.getItem" exists.
func (this SVGStringList) HasFuncGetItem() bool {
	return js.True == bindings.HasFuncSVGStringListGetItem(
		this.ref,
	)
}

// FuncGetItem returns the method "SVGStringList.getItem".
func (this SVGStringList) FuncGetItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncSVGStringListGetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetItem calls the method "SVGStringList.getItem".
func (this SVGStringList) GetItem(index uint32) (ret js.String) {
	bindings.CallSVGStringListGetItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGStringList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryGetItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListGetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncInsertItemBefore returns true if the method "SVGStringList.insertItemBefore" exists.
func (this SVGStringList) HasFuncInsertItemBefore() bool {
	return js.True == bindings.HasFuncSVGStringListInsertItemBefore(
		this.ref,
	)
}

// FuncInsertItemBefore returns the method "SVGStringList.insertItemBefore".
func (this SVGStringList) FuncInsertItemBefore() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	bindings.FuncSVGStringListInsertItemBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertItemBefore calls the method "SVGStringList.insertItemBefore".
func (this SVGStringList) InsertItemBefore(newItem js.String, index uint32) (ret js.String) {
	bindings.CallSVGStringListInsertItemBefore(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGStringList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryInsertItemBefore(newItem js.String, index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListInsertItemBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncReplaceItem returns true if the method "SVGStringList.replaceItem" exists.
func (this SVGStringList) HasFuncReplaceItem() bool {
	return js.True == bindings.HasFuncSVGStringListReplaceItem(
		this.ref,
	)
}

// FuncReplaceItem returns the method "SVGStringList.replaceItem".
func (this SVGStringList) FuncReplaceItem() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	bindings.FuncSVGStringListReplaceItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceItem calls the method "SVGStringList.replaceItem".
func (this SVGStringList) ReplaceItem(newItem js.String, index uint32) (ret js.String) {
	bindings.CallSVGStringListReplaceItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGStringList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryReplaceItem(newItem js.String, index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListReplaceItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasFuncRemoveItem returns true if the method "SVGStringList.removeItem" exists.
func (this SVGStringList) HasFuncRemoveItem() bool {
	return js.True == bindings.HasFuncSVGStringListRemoveItem(
		this.ref,
	)
}

// FuncRemoveItem returns the method "SVGStringList.removeItem".
func (this SVGStringList) FuncRemoveItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncSVGStringListRemoveItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveItem calls the method "SVGStringList.removeItem".
func (this SVGStringList) RemoveItem(index uint32) (ret js.String) {
	bindings.CallSVGStringListRemoveItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGStringList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryRemoveItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListRemoveItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendItem returns true if the method "SVGStringList.appendItem" exists.
func (this SVGStringList) HasFuncAppendItem() bool {
	return js.True == bindings.HasFuncSVGStringListAppendItem(
		this.ref,
	)
}

// FuncAppendItem returns the method "SVGStringList.appendItem".
func (this SVGStringList) FuncAppendItem() (fn js.Func[func(newItem js.String) js.String]) {
	bindings.FuncSVGStringListAppendItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendItem calls the method "SVGStringList.appendItem".
func (this SVGStringList) AppendItem(newItem js.String) (ret js.String) {
	bindings.CallSVGStringListAppendItem(
		this.ref, js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGStringList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TryAppendItem(newItem js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListAppendItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "SVGStringList." exists.
func (this SVGStringList) HasFuncSet() bool {
	return js.True == bindings.HasFuncSVGStringListSet(
		this.ref,
	)
}

// FuncSet returns the method "SVGStringList.".
func (this SVGStringList) FuncSet() (fn js.Func[func(index uint32, newItem js.String)]) {
	bindings.FuncSVGStringListSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "SVGStringList.".
func (this SVGStringList) Set(index uint32, newItem js.String) (ret js.Void) {
	bindings.CallSVGStringListSet(
		this.ref, js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGStringList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGStringList) TrySet(index uint32, newItem js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGStringListSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
}

type SVGAnimationElement struct {
	SVGElement
}

func (this SVGAnimationElement) Once() SVGAnimationElement {
	this.ref.Once()
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
	this.ref.Free()
}

// TargetElement returns the value of property "SVGAnimationElement.targetElement".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) TargetElement() (ret SVGElement, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementTargetElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RequiredExtensions returns the value of property "SVGAnimationElement.requiredExtensions".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) RequiredExtensions() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementRequiredExtensions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SystemLanguage returns the value of property "SVGAnimationElement.systemLanguage".
//
// It returns ok=false if there is no such property.
func (this SVGAnimationElement) SystemLanguage() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGAnimationElementSystemLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetStartTime returns true if the method "SVGAnimationElement.getStartTime" exists.
func (this SVGAnimationElement) HasFuncGetStartTime() bool {
	return js.True == bindings.HasFuncSVGAnimationElementGetStartTime(
		this.ref,
	)
}

// FuncGetStartTime returns the method "SVGAnimationElement.getStartTime".
func (this SVGAnimationElement) FuncGetStartTime() (fn js.Func[func() float32]) {
	bindings.FuncSVGAnimationElementGetStartTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetStartTime calls the method "SVGAnimationElement.getStartTime".
func (this SVGAnimationElement) GetStartTime() (ret float32) {
	bindings.CallSVGAnimationElementGetStartTime(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetStartTime calls the method "SVGAnimationElement.getStartTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryGetStartTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetStartTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCurrentTime returns true if the method "SVGAnimationElement.getCurrentTime" exists.
func (this SVGAnimationElement) HasFuncGetCurrentTime() bool {
	return js.True == bindings.HasFuncSVGAnimationElementGetCurrentTime(
		this.ref,
	)
}

// FuncGetCurrentTime returns the method "SVGAnimationElement.getCurrentTime".
func (this SVGAnimationElement) FuncGetCurrentTime() (fn js.Func[func() float32]) {
	bindings.FuncSVGAnimationElementGetCurrentTime(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentTime calls the method "SVGAnimationElement.getCurrentTime".
func (this SVGAnimationElement) GetCurrentTime() (ret float32) {
	bindings.CallSVGAnimationElementGetCurrentTime(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTime calls the method "SVGAnimationElement.getCurrentTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryGetCurrentTime() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetCurrentTime(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSimpleDuration returns true if the method "SVGAnimationElement.getSimpleDuration" exists.
func (this SVGAnimationElement) HasFuncGetSimpleDuration() bool {
	return js.True == bindings.HasFuncSVGAnimationElementGetSimpleDuration(
		this.ref,
	)
}

// FuncGetSimpleDuration returns the method "SVGAnimationElement.getSimpleDuration".
func (this SVGAnimationElement) FuncGetSimpleDuration() (fn js.Func[func() float32]) {
	bindings.FuncSVGAnimationElementGetSimpleDuration(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSimpleDuration calls the method "SVGAnimationElement.getSimpleDuration".
func (this SVGAnimationElement) GetSimpleDuration() (ret float32) {
	bindings.CallSVGAnimationElementGetSimpleDuration(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSimpleDuration calls the method "SVGAnimationElement.getSimpleDuration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryGetSimpleDuration() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementGetSimpleDuration(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBeginElement returns true if the method "SVGAnimationElement.beginElement" exists.
func (this SVGAnimationElement) HasFuncBeginElement() bool {
	return js.True == bindings.HasFuncSVGAnimationElementBeginElement(
		this.ref,
	)
}

// FuncBeginElement returns the method "SVGAnimationElement.beginElement".
func (this SVGAnimationElement) FuncBeginElement() (fn js.Func[func()]) {
	bindings.FuncSVGAnimationElementBeginElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginElement calls the method "SVGAnimationElement.beginElement".
func (this SVGAnimationElement) BeginElement() (ret js.Void) {
	bindings.CallSVGAnimationElementBeginElement(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBeginElement calls the method "SVGAnimationElement.beginElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryBeginElement() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementBeginElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBeginElementAt returns true if the method "SVGAnimationElement.beginElementAt" exists.
func (this SVGAnimationElement) HasFuncBeginElementAt() bool {
	return js.True == bindings.HasFuncSVGAnimationElementBeginElementAt(
		this.ref,
	)
}

// FuncBeginElementAt returns the method "SVGAnimationElement.beginElementAt".
func (this SVGAnimationElement) FuncBeginElementAt() (fn js.Func[func(offset float32)]) {
	bindings.FuncSVGAnimationElementBeginElementAt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginElementAt calls the method "SVGAnimationElement.beginElementAt".
func (this SVGAnimationElement) BeginElementAt(offset float32) (ret js.Void) {
	bindings.CallSVGAnimationElementBeginElementAt(
		this.ref, js.Pointer(&ret),
		float32(offset),
	)

	return
}

// TryBeginElementAt calls the method "SVGAnimationElement.beginElementAt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryBeginElementAt(offset float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementBeginElementAt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(offset),
	)

	return
}

// HasFuncEndElement returns true if the method "SVGAnimationElement.endElement" exists.
func (this SVGAnimationElement) HasFuncEndElement() bool {
	return js.True == bindings.HasFuncSVGAnimationElementEndElement(
		this.ref,
	)
}

// FuncEndElement returns the method "SVGAnimationElement.endElement".
func (this SVGAnimationElement) FuncEndElement() (fn js.Func[func()]) {
	bindings.FuncSVGAnimationElementEndElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndElement calls the method "SVGAnimationElement.endElement".
func (this SVGAnimationElement) EndElement() (ret js.Void) {
	bindings.CallSVGAnimationElementEndElement(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEndElement calls the method "SVGAnimationElement.endElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryEndElement() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementEndElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEndElementAt returns true if the method "SVGAnimationElement.endElementAt" exists.
func (this SVGAnimationElement) HasFuncEndElementAt() bool {
	return js.True == bindings.HasFuncSVGAnimationElementEndElementAt(
		this.ref,
	)
}

// FuncEndElementAt returns the method "SVGAnimationElement.endElementAt".
func (this SVGAnimationElement) FuncEndElementAt() (fn js.Func[func(offset float32)]) {
	bindings.FuncSVGAnimationElementEndElementAt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndElementAt calls the method "SVGAnimationElement.endElementAt".
func (this SVGAnimationElement) EndElementAt(offset float32) (ret js.Void) {
	bindings.CallSVGAnimationElementEndElementAt(
		this.ref, js.Pointer(&ret),
		float32(offset),
	)

	return
}

// TryEndElementAt calls the method "SVGAnimationElement.endElementAt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGAnimationElement) TryEndElementAt(offset float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGAnimationElementEndElementAt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *SVGBoundingBoxOptions) UpdateFrom(ref js.Ref) {
	bindings.SVGBoundingBoxOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SVGBoundingBoxOptions) Update(ref js.Ref) {
	bindings.SVGBoundingBoxOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SVGBoundingBoxOptions) FreeMembers(recursive bool) {
}

type SVGCircleElement struct {
	SVGGeometryElement
}

func (this SVGCircleElement) Once() SVGCircleElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Cx returns the value of property "SVGCircleElement.cx".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) Cx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementCx(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cy returns the value of property "SVGCircleElement.cy".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) Cy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementCy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// R returns the value of property "SVGCircleElement.r".
//
// It returns ok=false if there is no such property.
func (this SVGCircleElement) R() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGCircleElementR(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGClipPathElement struct {
	SVGElement
}

func (this SVGClipPathElement) Once() SVGClipPathElement {
	this.ref.Once()
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
	this.ref.Free()
}

// ClipPathUnits returns the value of property "SVGClipPathElement.clipPathUnits".
//
// It returns ok=false if there is no such property.
func (this SVGClipPathElement) ClipPathUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGClipPathElementClipPathUnits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "SVGClipPathElement.transform".
//
// It returns ok=false if there is no such property.
func (this SVGClipPathElement) Transform() (ret SVGAnimatedTransformList, ok bool) {
	ok = js.True == bindings.GetSVGClipPathElementTransform(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "SVGComponentTransferFunctionElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Type() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TableValues returns the value of property "SVGComponentTransferFunctionElement.tableValues".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) TableValues() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementTableValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Slope returns the value of property "SVGComponentTransferFunctionElement.slope".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Slope() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementSlope(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Intercept returns the value of property "SVGComponentTransferFunctionElement.intercept".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Intercept() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementIntercept(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Amplitude returns the value of property "SVGComponentTransferFunctionElement.amplitude".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Amplitude() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementAmplitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Exponent returns the value of property "SVGComponentTransferFunctionElement.exponent".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Exponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementExponent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Offset returns the value of property "SVGComponentTransferFunctionElement.offset".
//
// It returns ok=false if there is no such property.
func (this SVGComponentTransferFunctionElement) Offset() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGComponentTransferFunctionElementOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGDefsElement struct {
	SVGGraphicsElement
}

func (this SVGDefsElement) Once() SVGDefsElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGDescElement struct {
	SVGElement
}

func (this SVGDescElement) Once() SVGDescElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGDiscardElement struct {
	SVGAnimationElement
}

func (this SVGDiscardElement) Once() SVGDiscardElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGEllipseElement struct {
	SVGGeometryElement
}

func (this SVGEllipseElement) Once() SVGEllipseElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Cx returns the value of property "SVGEllipseElement.cx".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Cx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementCx(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cy returns the value of property "SVGEllipseElement.cy".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Cy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementCy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Rx returns the value of property "SVGEllipseElement.rx".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Rx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementRx(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ry returns the value of property "SVGEllipseElement.ry".
//
// It returns ok=false if there is no such property.
func (this SVGEllipseElement) Ry() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGEllipseElementRy(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEBlendElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFEBlendElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementIn2(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "SVGFEBlendElement.mode".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Mode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEBlendElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEBlendElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEBlendElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEBlendElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEBlendElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEBlendElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEBlendElementResult(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEColorMatrixElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "SVGFEColorMatrixElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Type() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Values returns the value of property "SVGFEColorMatrixElement.values".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Values() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEColorMatrixElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEColorMatrixElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEColorMatrixElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEColorMatrixElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEColorMatrixElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEColorMatrixElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEColorMatrixElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEComponentTransferElement struct {
	SVGElement
}

func (this SVGFEComponentTransferElement) Once() SVGFEComponentTransferElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEComponentTransferElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEComponentTransferElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEComponentTransferElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEComponentTransferElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEComponentTransferElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEComponentTransferElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEComponentTransferElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEComponentTransferElementResult(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFECompositeElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFECompositeElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementIn2(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Operator returns the value of property "SVGFECompositeElement.operator".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Operator() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementOperator(
		this.ref, js.Pointer(&ret),
	)
	return
}

// K1 returns the value of property "SVGFECompositeElement.k1".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K1() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// K2 returns the value of property "SVGFECompositeElement.k2".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K2() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK2(
		this.ref, js.Pointer(&ret),
	)
	return
}

// K3 returns the value of property "SVGFECompositeElement.k3".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K3() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK3(
		this.ref, js.Pointer(&ret),
	)
	return
}

// K4 returns the value of property "SVGFECompositeElement.k4".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) K4() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementK4(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFECompositeElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFECompositeElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFECompositeElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFECompositeElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFECompositeElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFECompositeElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFECompositeElementResult(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEConvolveMatrixElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OrderX returns the value of property "SVGFEConvolveMatrixElement.orderX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderX() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementOrderX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OrderY returns the value of property "SVGFEConvolveMatrixElement.orderY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderY() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementOrderY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelMatrix returns the value of property "SVGFEConvolveMatrixElement.kernelMatrix".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelMatrix() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Divisor returns the value of property "SVGFEConvolveMatrixElement.divisor".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Divisor() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementDivisor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Bias returns the value of property "SVGFEConvolveMatrixElement.bias".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Bias() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementBias(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TargetX returns the value of property "SVGFEConvolveMatrixElement.targetX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetX() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementTargetX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TargetY returns the value of property "SVGFEConvolveMatrixElement.targetY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetY() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementTargetY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EdgeMode returns the value of property "SVGFEConvolveMatrixElement.edgeMode".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) EdgeMode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementEdgeMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreserveAlpha returns the value of property "SVGFEConvolveMatrixElement.preserveAlpha".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) PreserveAlpha() (ret SVGAnimatedBoolean, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementPreserveAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEConvolveMatrixElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEConvolveMatrixElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEConvolveMatrixElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEConvolveMatrixElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEConvolveMatrixElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEConvolveMatrixElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEConvolveMatrixElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEDiffuseLightingElement struct {
	SVGElement
}

func (this SVGFEDiffuseLightingElement) Once() SVGFEDiffuseLightingElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEDiffuseLightingElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SurfaceScale returns the value of property "SVGFEDiffuseLightingElement.surfaceScale".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) SurfaceScale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementSurfaceScale(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DiffuseConstant returns the value of property "SVGFEDiffuseLightingElement.diffuseConstant".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) DiffuseConstant() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementDiffuseConstant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDiffuseLightingElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDiffuseLightingElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDiffuseLightingElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDiffuseLightingElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDiffuseLightingElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDiffuseLightingElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDiffuseLightingElementResult(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEDisplacementMapElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// In2 returns the value of property "SVGFEDisplacementMapElement.in2".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) In2() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementIn2(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scale returns the value of property "SVGFEDisplacementMapElement.scale".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Scale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementScale(
		this.ref, js.Pointer(&ret),
	)
	return
}

// XChannelSelector returns the value of property "SVGFEDisplacementMapElement.xChannelSelector".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) XChannelSelector() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementXChannelSelector(
		this.ref, js.Pointer(&ret),
	)
	return
}

// YChannelSelector returns the value of property "SVGFEDisplacementMapElement.yChannelSelector".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) YChannelSelector() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementYChannelSelector(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDisplacementMapElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDisplacementMapElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDisplacementMapElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDisplacementMapElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDisplacementMapElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDisplacementMapElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDisplacementMapElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEDistantLightElement struct {
	SVGElement
}

func (this SVGFEDistantLightElement) Once() SVGFEDistantLightElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Azimuth returns the value of property "SVGFEDistantLightElement.azimuth".
//
// It returns ok=false if there is no such property.
func (this SVGFEDistantLightElement) Azimuth() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDistantLightElementAzimuth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Elevation returns the value of property "SVGFEDistantLightElement.elevation".
//
// It returns ok=false if there is no such property.
func (this SVGFEDistantLightElement) Elevation() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDistantLightElementElevation(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEDropShadowElement struct {
	SVGElement
}

func (this SVGFEDropShadowElement) Once() SVGFEDropShadowElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEDropShadowElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dx returns the value of property "SVGFEDropShadowElement.dx".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Dx() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementDx(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dy returns the value of property "SVGFEDropShadowElement.dy".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Dy() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementDy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StdDeviationX returns the value of property "SVGFEDropShadowElement.stdDeviationX".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementStdDeviationX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StdDeviationY returns the value of property "SVGFEDropShadowElement.stdDeviationY".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementStdDeviationY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEDropShadowElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEDropShadowElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEDropShadowElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEDropShadowElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEDropShadowElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEDropShadowElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEDropShadowElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetStdDeviation returns true if the method "SVGFEDropShadowElement.setStdDeviation" exists.
func (this SVGFEDropShadowElement) HasFuncSetStdDeviation() bool {
	return js.True == bindings.HasFuncSVGFEDropShadowElementSetStdDeviation(
		this.ref,
	)
}

// FuncSetStdDeviation returns the method "SVGFEDropShadowElement.setStdDeviation".
func (this SVGFEDropShadowElement) FuncSetStdDeviation() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	bindings.FuncSVGFEDropShadowElementSetStdDeviation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStdDeviation calls the method "SVGFEDropShadowElement.setStdDeviation".
func (this SVGFEDropShadowElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void) {
	bindings.CallSVGFEDropShadowElementSetStdDeviation(
		this.ref, js.Pointer(&ret),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

// TrySetStdDeviation calls the method "SVGFEDropShadowElement.setStdDeviation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGFEDropShadowElement) TrySetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGFEDropShadowElementSetStdDeviation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

type SVGFEFloodElement struct {
	SVGElement
}

func (this SVGFEFloodElement) Once() SVGFEFloodElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGFEFloodElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEFloodElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEFloodElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEFloodElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEFloodElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEFloodElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEFloodElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEFuncAElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncAElement) Once() SVGFEFuncAElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGFEFuncBElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncBElement) Once() SVGFEFuncBElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGFEFuncGElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncGElement) Once() SVGFEFuncGElement {
	this.ref.Once()
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
	this.ref.Free()
}

type SVGFEFuncRElement struct {
	SVGComponentTransferFunctionElement
}

func (this SVGFEFuncRElement) Once() SVGFEFuncRElement {
	this.ref.Once()
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
	this.ref.Free()
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEGaussianBlurElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StdDeviationX returns the value of property "SVGFEGaussianBlurElement.stdDeviationX".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementStdDeviationX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StdDeviationY returns the value of property "SVGFEGaussianBlurElement.stdDeviationY".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementStdDeviationY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EdgeMode returns the value of property "SVGFEGaussianBlurElement.edgeMode".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) EdgeMode() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementEdgeMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEGaussianBlurElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEGaussianBlurElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEGaussianBlurElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEGaussianBlurElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEGaussianBlurElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEGaussianBlurElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEGaussianBlurElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetStdDeviation returns true if the method "SVGFEGaussianBlurElement.setStdDeviation" exists.
func (this SVGFEGaussianBlurElement) HasFuncSetStdDeviation() bool {
	return js.True == bindings.HasFuncSVGFEGaussianBlurElementSetStdDeviation(
		this.ref,
	)
}

// FuncSetStdDeviation returns the method "SVGFEGaussianBlurElement.setStdDeviation".
func (this SVGFEGaussianBlurElement) FuncSetStdDeviation() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	bindings.FuncSVGFEGaussianBlurElementSetStdDeviation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStdDeviation calls the method "SVGFEGaussianBlurElement.setStdDeviation".
func (this SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void) {
	bindings.CallSVGFEGaussianBlurElementSetStdDeviation(
		this.ref, js.Pointer(&ret),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

// TrySetStdDeviation calls the method "SVGFEGaussianBlurElement.setStdDeviation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SVGFEGaussianBlurElement) TrySetStdDeviation(stdDeviationX float32, stdDeviationY float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGFEGaussianBlurElementSetStdDeviation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	return
}

type SVGFEImageElement struct {
	SVGElement
}

func (this SVGFEImageElement) Once() SVGFEImageElement {
	this.ref.Once()
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
	this.ref.Free()
}

// PreserveAspectRatio returns the value of property "SVGFEImageElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementPreserveAspectRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CrossOrigin returns the value of property "SVGFEImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) CrossOrigin() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEImageElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEImageElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEImageElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEImageElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEImageElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGFEImageElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGFEImageElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEImageElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEMergeElement struct {
	SVGElement
}

func (this SVGFEMergeElement) Once() SVGFEMergeElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGFEMergeElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEMergeElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEMergeElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEMergeElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEMergeElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEMergeNodeElement struct {
	SVGElement
}

func (this SVGFEMergeNodeElement) Once() SVGFEMergeNodeElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEMergeNodeElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEMergeNodeElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMergeNodeElementIn1(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEMorphologyElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Operator returns the value of property "SVGFEMorphologyElement.operator".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Operator() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementOperator(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RadiusX returns the value of property "SVGFEMorphologyElement.radiusX".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) RadiusX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementRadiusX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RadiusY returns the value of property "SVGFEMorphologyElement.radiusY".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) RadiusY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementRadiusY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEMorphologyElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEMorphologyElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEMorphologyElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEMorphologyElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEMorphologyElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEMorphologyElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEMorphologyElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEOffsetElement struct {
	SVGElement
}

func (this SVGFEOffsetElement) Once() SVGFEOffsetElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFEOffsetElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dx returns the value of property "SVGFEOffsetElement.dx".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Dx() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementDx(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dy returns the value of property "SVGFEOffsetElement.dy".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Dy() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementDy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFEOffsetElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEOffsetElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFEOffsetElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFEOffsetElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFEOffsetElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFEOffsetElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFEOffsetElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFEPointLightElement struct {
	SVGElement
}

func (this SVGFEPointLightElement) Once() SVGFEPointLightElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGFEPointLightElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) X() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFEPointLightElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) Y() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "SVGFEPointLightElement.z".
//
// It returns ok=false if there is no such property.
func (this SVGFEPointLightElement) Z() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFEPointLightElementZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFESpecularLightingElement struct {
	SVGElement
}

func (this SVGFESpecularLightingElement) Once() SVGFESpecularLightingElement {
	this.ref.Once()
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
	this.ref.Free()
}

// In1 returns the value of property "SVGFESpecularLightingElement.in1".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementIn1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SurfaceScale returns the value of property "SVGFESpecularLightingElement.surfaceScale".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SurfaceScale() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSurfaceScale(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SpecularConstant returns the value of property "SVGFESpecularLightingElement.specularConstant".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularConstant() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSpecularConstant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SpecularExponent returns the value of property "SVGFESpecularLightingElement.specularExponent".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularExponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementSpecularExponent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthX returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthX".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementKernelUnitLengthX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KernelUnitLengthY returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthY".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementKernelUnitLengthY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFESpecularLightingElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFESpecularLightingElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFESpecularLightingElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFESpecularLightingElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFESpecularLightingElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFESpecularLightingElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFESpecularLightingElementResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SVGFESpotLightElement struct {
	SVGElement
}

func (this SVGFESpotLightElement) Once() SVGFESpotLightElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGFESpotLightElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) X() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFESpotLightElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) Y() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "SVGFESpotLightElement.z".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) Z() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointsAtX returns the value of property "SVGFESpotLightElement.pointsAtX".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointsAtY returns the value of property "SVGFESpotLightElement.pointsAtY".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointsAtZ returns the value of property "SVGFESpotLightElement.pointsAtZ".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) PointsAtZ() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementPointsAtZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SpecularExponent returns the value of property "SVGFESpotLightElement.specularExponent".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) SpecularExponent() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementSpecularExponent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LimitingConeAngle returns the value of property "SVGFESpotLightElement.limitingConeAngle".
//
// It returns ok=false if there is no such property.
func (this SVGFESpotLightElement) LimitingConeAngle() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFESpotLightElementLimitingConeAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

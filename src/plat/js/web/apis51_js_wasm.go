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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedEnumeration) BaseVal() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedEnumerationBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// BaseVal sets the value of property "SVGAnimatedEnumeration.baseVal" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedEnumeration) AnimVal() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedEnumerationAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedInteger) BaseVal() (int32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedIntegerBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// BaseVal sets the value of property "SVGAnimatedInteger.baseVal" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedInteger) AnimVal() (int32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedIntegerAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGLengthList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfItems returns the value of property "SVGLengthList.numberOfItems".
//
// The returned bool will be false if there is no such property.
func (this SVGLengthList) NumberOfItems() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGLengthListNumberOfItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Clear calls the method "SVGLengthList.clear".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "SVGLengthList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGLengthListClearFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGLengthList.initialize".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) Initialize(newItem SVGLength) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListInitialize(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// InitializeFunc returns the method "SVGLengthList.initialize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) InitializeFunc() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListInitializeFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGLengthList.getItem".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) GetItem(index uint32) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListGetItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "SVGLengthList.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) GetItemFunc() (fn js.Func[func(index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListGetItemFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGLengthList.insertItemBefore".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) InsertItemBefore(newItem SVGLength, index uint32) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListInsertItemBefore(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// InsertItemBeforeFunc returns the method "SVGLengthList.insertItemBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGLengthList.replaceItem".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) ReplaceItem(newItem SVGLength, index uint32) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListReplaceItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// ReplaceItemFunc returns the method "SVGLengthList.replaceItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) ReplaceItemFunc() (fn js.Func[func(newItem SVGLength, index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGLengthList.removeItem".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) RemoveItem(index uint32) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// RemoveItemFunc returns the method "SVGLengthList.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGLengthList.appendItem".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) AppendItem(newItem SVGLength) (SVGLength, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListAppendItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGLength{}.FromRef(_ret), _ok
}

// AppendItemFunc returns the method "SVGLengthList.appendItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) AppendItemFunc() (fn js.Func[func(newItem SVGLength) SVGLength]) {
	return fn.FromRef(
		bindings.SVGLengthListAppendItemFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGLengthList.".
//
// The returned bool will be false if there is no such method.
func (this SVGLengthList) Set(index uint32, newItem SVGLength) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGLengthListSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		newItem.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "SVGLengthList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGLengthList) SetFunc() (fn js.Func[func(index uint32, newItem SVGLength)]) {
	return fn.FromRef(
		bindings.SVGLengthListSetFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedLengthList) BaseVal() (SVGLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedLengthListBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGLengthList{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedLengthList.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedLengthList) AnimVal() (SVGLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedLengthListAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGLengthList{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedNumber) BaseVal() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedNumberBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// BaseVal sets the value of property "SVGAnimatedNumber.baseVal" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedNumber) AnimVal() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedNumberAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGNumberList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGNumberListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfItems returns the value of property "SVGNumberList.numberOfItems".
//
// The returned bool will be false if there is no such property.
func (this SVGNumberList) NumberOfItems() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGNumberListNumberOfItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Clear calls the method "SVGNumberList.clear".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "SVGNumberList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGNumberListClearFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGNumberList.initialize".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) Initialize(newItem SVGNumber) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListInitialize(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// InitializeFunc returns the method "SVGNumberList.initialize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) InitializeFunc() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListInitializeFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGNumberList.getItem".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) GetItem(index uint32) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListGetItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "SVGNumberList.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) GetItemFunc() (fn js.Func[func(index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListGetItemFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGNumberList.insertItemBefore".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) InsertItemBefore(newItem SVGNumber, index uint32) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListInsertItemBefore(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// InsertItemBeforeFunc returns the method "SVGNumberList.insertItemBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGNumberList.replaceItem".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) ReplaceItem(newItem SVGNumber, index uint32) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListReplaceItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// ReplaceItemFunc returns the method "SVGNumberList.replaceItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) ReplaceItemFunc() (fn js.Func[func(newItem SVGNumber, index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGNumberList.removeItem".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) RemoveItem(index uint32) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// RemoveItemFunc returns the method "SVGNumberList.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGNumberList.appendItem".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) AppendItem(newItem SVGNumber) (SVGNumber, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListAppendItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGNumber{}.FromRef(_ret), _ok
}

// AppendItemFunc returns the method "SVGNumberList.appendItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) AppendItemFunc() (fn js.Func[func(newItem SVGNumber) SVGNumber]) {
	return fn.FromRef(
		bindings.SVGNumberListAppendItemFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGNumberList.".
//
// The returned bool will be false if there is no such method.
func (this SVGNumberList) Set(index uint32, newItem SVGNumber) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGNumberListSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		newItem.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "SVGNumberList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGNumberList) SetFunc() (fn js.Func[func(index uint32, newItem SVGNumber)]) {
	return fn.FromRef(
		bindings.SVGNumberListSetFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedNumberList) BaseVal() (SVGNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedNumberListBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGNumberList{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedNumberList.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedNumberList) AnimVal() (SVGNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedNumberListAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGNumberList{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGTransformList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGTransformListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfItems returns the value of property "SVGTransformList.numberOfItems".
//
// The returned bool will be false if there is no such property.
func (this SVGTransformList) NumberOfItems() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGTransformListNumberOfItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Clear calls the method "SVGTransformList.clear".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "SVGTransformList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGTransformListClearFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGTransformList.initialize".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) Initialize(newItem SVGTransform) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListInitialize(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// InitializeFunc returns the method "SVGTransformList.initialize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) InitializeFunc() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListInitializeFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGTransformList.getItem".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) GetItem(index uint32) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListGetItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "SVGTransformList.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) GetItemFunc() (fn js.Func[func(index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListGetItemFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGTransformList.insertItemBefore".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) InsertItemBefore(newItem SVGTransform, index uint32) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListInsertItemBefore(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// InsertItemBeforeFunc returns the method "SVGTransformList.insertItemBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) InsertItemBeforeFunc() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGTransformList.replaceItem".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) ReplaceItem(newItem SVGTransform, index uint32) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListReplaceItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// ReplaceItemFunc returns the method "SVGTransformList.replaceItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) ReplaceItemFunc() (fn js.Func[func(newItem SVGTransform, index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGTransformList.removeItem".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) RemoveItem(index uint32) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// RemoveItemFunc returns the method "SVGTransformList.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) RemoveItemFunc() (fn js.Func[func(index uint32) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGTransformList.appendItem".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) AppendItem(newItem SVGTransform) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListAppendItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// AppendItemFunc returns the method "SVGTransformList.appendItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) AppendItemFunc() (fn js.Func[func(newItem SVGTransform) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListAppendItemFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGTransformList.".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) Set(index uint32, newItem SVGTransform) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		newItem.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "SVGTransformList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) SetFunc() (fn js.Func[func(index uint32, newItem SVGTransform)]) {
	return fn.FromRef(
		bindings.SVGTransformListSetFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix calls the method "SVGTransformList.createSVGTransformFromMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) CreateSVGTransformFromMatrix(matrix DOMMatrix2DInit) (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListCreateSVGTransformFromMatrix(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&matrix),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// CreateSVGTransformFromMatrixFunc returns the method "SVGTransformList.createSVGTransformFromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) CreateSVGTransformFromMatrixFunc() (fn js.Func[func(matrix DOMMatrix2DInit) SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListCreateSVGTransformFromMatrixFunc(
			this.Ref(),
		),
	)
}

// CreateSVGTransformFromMatrix1 calls the method "SVGTransformList.createSVGTransformFromMatrix".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) CreateSVGTransformFromMatrix1() (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListCreateSVGTransformFromMatrix1(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// CreateSVGTransformFromMatrix1Func returns the method "SVGTransformList.createSVGTransformFromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) CreateSVGTransformFromMatrix1Func() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListCreateSVGTransformFromMatrix1Func(
			this.Ref(),
		),
	)
}

// Consolidate calls the method "SVGTransformList.consolidate".
//
// The returned bool will be false if there is no such method.
func (this SVGTransformList) Consolidate() (SVGTransform, bool) {
	var _ok bool
	_ret := bindings.CallSVGTransformListConsolidate(
		this.Ref(), js.Pointer(&_ok),
	)

	return SVGTransform{}.FromRef(_ret), _ok
}

// ConsolidateFunc returns the method "SVGTransformList.consolidate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTransformList) ConsolidateFunc() (fn js.Func[func() SVGTransform]) {
	return fn.FromRef(
		bindings.SVGTransformListConsolidateFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGAnimatedTransformList) BaseVal() (SVGTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedTransformListBaseVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGTransformList{}.FromRef(_ret), _ok
}

// AnimVal returns the value of property "SVGAnimatedTransformList.animVal".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimatedTransformList) AnimVal() (SVGTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimatedTransformListAnimVal(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGTransformList{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGStringList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGStringListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfItems returns the value of property "SVGStringList.numberOfItems".
//
// The returned bool will be false if there is no such property.
func (this SVGStringList) NumberOfItems() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGStringListNumberOfItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Clear calls the method "SVGStringList.clear".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "SVGStringList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGStringListClearFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGStringList.initialize".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) Initialize(newItem js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListInitialize(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// InitializeFunc returns the method "SVGStringList.initialize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) InitializeFunc() (fn js.Func[func(newItem js.String) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListInitializeFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGStringList.getItem".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) GetItem(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListGetItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "SVGStringList.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) GetItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListGetItemFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGStringList.insertItemBefore".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) InsertItemBefore(newItem js.String, index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListInsertItemBefore(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// InsertItemBeforeFunc returns the method "SVGStringList.insertItemBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) InsertItemBeforeFunc() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGStringList.replaceItem".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) ReplaceItem(newItem js.String, index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListReplaceItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ReplaceItemFunc returns the method "SVGStringList.replaceItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) ReplaceItemFunc() (fn js.Func[func(newItem js.String, index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGStringList.removeItem".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) RemoveItem(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// RemoveItemFunc returns the method "SVGStringList.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) RemoveItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGStringList.appendItem".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) AppendItem(newItem js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListAppendItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// AppendItemFunc returns the method "SVGStringList.appendItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) AppendItemFunc() (fn js.Func[func(newItem js.String) js.String]) {
	return fn.FromRef(
		bindings.SVGStringListAppendItemFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGStringList.".
//
// The returned bool will be false if there is no such method.
func (this SVGStringList) Set(index uint32, newItem js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGStringListSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		newItem.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "SVGStringList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGStringList) SetFunc() (fn js.Func[func(index uint32, newItem js.String)]) {
	return fn.FromRef(
		bindings.SVGStringListSetFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGAnimationElement) TargetElement() (SVGElement, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimationElementTargetElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGElement{}.FromRef(_ret), _ok
}

// RequiredExtensions returns the value of property "SVGAnimationElement.requiredExtensions".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimationElement) RequiredExtensions() (SVGStringList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimationElementRequiredExtensions(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGStringList{}.FromRef(_ret), _ok
}

// SystemLanguage returns the value of property "SVGAnimationElement.systemLanguage".
//
// The returned bool will be false if there is no such property.
func (this SVGAnimationElement) SystemLanguage() (SVGStringList, bool) {
	var _ok bool
	_ret := bindings.GetSVGAnimationElementSystemLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGStringList{}.FromRef(_ret), _ok
}

// GetStartTime calls the method "SVGAnimationElement.getStartTime".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) GetStartTime() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementGetStartTime(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetStartTimeFunc returns the method "SVGAnimationElement.getStartTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) GetStartTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetStartTimeFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTime calls the method "SVGAnimationElement.getCurrentTime".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) GetCurrentTime() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementGetCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetCurrentTimeFunc returns the method "SVGAnimationElement.getCurrentTime".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) GetCurrentTimeFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetCurrentTimeFunc(
			this.Ref(),
		),
	)
}

// GetSimpleDuration calls the method "SVGAnimationElement.getSimpleDuration".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) GetSimpleDuration() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementGetSimpleDuration(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetSimpleDurationFunc returns the method "SVGAnimationElement.getSimpleDuration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) GetSimpleDurationFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGAnimationElementGetSimpleDurationFunc(
			this.Ref(),
		),
	)
}

// BeginElement calls the method "SVGAnimationElement.beginElement".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) BeginElement() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementBeginElement(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginElementFunc returns the method "SVGAnimationElement.beginElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) BeginElementFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGAnimationElementBeginElementFunc(
			this.Ref(),
		),
	)
}

// BeginElementAt calls the method "SVGAnimationElement.beginElementAt".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) BeginElementAt(offset float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementBeginElementAt(
		this.Ref(), js.Pointer(&_ok),
		float32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginElementAtFunc returns the method "SVGAnimationElement.beginElementAt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) BeginElementAtFunc() (fn js.Func[func(offset float32)]) {
	return fn.FromRef(
		bindings.SVGAnimationElementBeginElementAtFunc(
			this.Ref(),
		),
	)
}

// EndElement calls the method "SVGAnimationElement.endElement".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) EndElement() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementEndElement(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndElementFunc returns the method "SVGAnimationElement.endElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) EndElementFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGAnimationElementEndElementFunc(
			this.Ref(),
		),
	)
}

// EndElementAt calls the method "SVGAnimationElement.endElementAt".
//
// The returned bool will be false if there is no such method.
func (this SVGAnimationElement) EndElementAt(offset float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGAnimationElementEndElementAt(
		this.Ref(), js.Pointer(&_ok),
		float32(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndElementAtFunc returns the method "SVGAnimationElement.endElementAt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGAnimationElement) EndElementAtFunc() (fn js.Func[func(offset float32)]) {
	return fn.FromRef(
		bindings.SVGAnimationElementEndElementAtFunc(
			this.Ref(),
		),
	)
}

type SVGBoundingBoxOptions struct {
	// Fill is "SVGBoundingBoxOptions.fill"
	//
	// Optional, defaults to true.
	Fill bool
	// Stroke is "SVGBoundingBoxOptions.stroke"
	//
	// Optional, defaults to false.
	Stroke bool
	// Markers is "SVGBoundingBoxOptions.markers"
	//
	// Optional, defaults to false.
	Markers bool
	// Clipped is "SVGBoundingBoxOptions.clipped"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 SVGBoundingBoxOptions SVGBoundingBoxOptions [// SVGBoundingBoxOptions] [0x14000a01040 0x14000a01180 0x14000a012c0 0x14000a01400 0x14000a010e0 0x14000a01220 0x14000a01360 0x14000a014a0] 0x14000a02078 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this SVGCircleElement) Cx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGCircleElementCx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Cy returns the value of property "SVGCircleElement.cy".
//
// The returned bool will be false if there is no such property.
func (this SVGCircleElement) Cy() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGCircleElementCy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// R returns the value of property "SVGCircleElement.r".
//
// The returned bool will be false if there is no such property.
func (this SVGCircleElement) R() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGCircleElementR(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGClipPathElement) ClipPathUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGClipPathElementClipPathUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Transform returns the value of property "SVGClipPathElement.transform".
//
// The returned bool will be false if there is no such property.
func (this SVGClipPathElement) Transform() (SVGAnimatedTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGClipPathElementTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedTransformList{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Type() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// TableValues returns the value of property "SVGComponentTransferFunctionElement.tableValues".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) TableValues() (SVGAnimatedNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementTableValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumberList{}.FromRef(_ret), _ok
}

// Slope returns the value of property "SVGComponentTransferFunctionElement.slope".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Slope() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementSlope(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Intercept returns the value of property "SVGComponentTransferFunctionElement.intercept".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Intercept() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementIntercept(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Amplitude returns the value of property "SVGComponentTransferFunctionElement.amplitude".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Amplitude() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementAmplitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Exponent returns the value of property "SVGComponentTransferFunctionElement.exponent".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Exponent() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementExponent(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Offset returns the value of property "SVGComponentTransferFunctionElement.offset".
//
// The returned bool will be false if there is no such property.
func (this SVGComponentTransferFunctionElement) Offset() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGComponentTransferFunctionElementOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGEllipseElement) Cx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGEllipseElementCx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Cy returns the value of property "SVGEllipseElement.cy".
//
// The returned bool will be false if there is no such property.
func (this SVGEllipseElement) Cy() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGEllipseElementCy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Rx returns the value of property "SVGEllipseElement.rx".
//
// The returned bool will be false if there is no such property.
func (this SVGEllipseElement) Rx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGEllipseElementRx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Ry returns the value of property "SVGEllipseElement.ry".
//
// The returned bool will be false if there is no such property.
func (this SVGEllipseElement) Ry() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGEllipseElementRy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// In2 returns the value of property "SVGFEBlendElement.in2".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) In2() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementIn2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Mode returns the value of property "SVGFEBlendElement.mode".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) Mode() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEBlendElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEBlendElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEBlendElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEBlendElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEBlendElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEBlendElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEBlendElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Type returns the value of property "SVGFEColorMatrixElement.type".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Type() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Values returns the value of property "SVGFEColorMatrixElement.values".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Values() (SVGAnimatedNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumberList{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEColorMatrixElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEColorMatrixElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEColorMatrixElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEColorMatrixElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEColorMatrixElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEColorMatrixElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEColorMatrixElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEComponentTransferElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEComponentTransferElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEComponentTransferElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEComponentTransferElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEComponentTransferElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEComponentTransferElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEComponentTransferElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// In2 returns the value of property "SVGFECompositeElement.in2".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) In2() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementIn2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Operator returns the value of property "SVGFECompositeElement.operator".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) Operator() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementOperator(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// K1 returns the value of property "SVGFECompositeElement.k1".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) K1() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementK1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// K2 returns the value of property "SVGFECompositeElement.k2".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) K2() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementK2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// K3 returns the value of property "SVGFECompositeElement.k3".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) K3() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementK3(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// K4 returns the value of property "SVGFECompositeElement.k4".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) K4() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementK4(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFECompositeElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFECompositeElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFECompositeElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFECompositeElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFECompositeElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFECompositeElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFECompositeElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// OrderX returns the value of property "SVGFEConvolveMatrixElement.orderX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderX() (SVGAnimatedInteger, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementOrderX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedInteger{}.FromRef(_ret), _ok
}

// OrderY returns the value of property "SVGFEConvolveMatrixElement.orderY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) OrderY() (SVGAnimatedInteger, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementOrderY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedInteger{}.FromRef(_ret), _ok
}

// KernelMatrix returns the value of property "SVGFEConvolveMatrixElement.kernelMatrix".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelMatrix() (SVGAnimatedNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementKernelMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumberList{}.FromRef(_ret), _ok
}

// Divisor returns the value of property "SVGFEConvolveMatrixElement.divisor".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Divisor() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementDivisor(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Bias returns the value of property "SVGFEConvolveMatrixElement.bias".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Bias() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementBias(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// TargetX returns the value of property "SVGFEConvolveMatrixElement.targetX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetX() (SVGAnimatedInteger, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementTargetX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedInteger{}.FromRef(_ret), _ok
}

// TargetY returns the value of property "SVGFEConvolveMatrixElement.targetY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) TargetY() (SVGAnimatedInteger, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementTargetY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedInteger{}.FromRef(_ret), _ok
}

// EdgeMode returns the value of property "SVGFEConvolveMatrixElement.edgeMode".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) EdgeMode() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementEdgeMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// KernelUnitLengthX returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// KernelUnitLengthY returns the value of property "SVGFEConvolveMatrixElement.kernelUnitLengthY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) KernelUnitLengthY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// PreserveAlpha returns the value of property "SVGFEConvolveMatrixElement.preserveAlpha".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) PreserveAlpha() (SVGAnimatedBoolean, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementPreserveAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedBoolean{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEConvolveMatrixElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEConvolveMatrixElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEConvolveMatrixElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEConvolveMatrixElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEConvolveMatrixElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEConvolveMatrixElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEConvolveMatrixElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// SurfaceScale returns the value of property "SVGFEDiffuseLightingElement.surfaceScale".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) SurfaceScale() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementSurfaceScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// DiffuseConstant returns the value of property "SVGFEDiffuseLightingElement.diffuseConstant".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) DiffuseConstant() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementDiffuseConstant(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// KernelUnitLengthX returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// KernelUnitLengthY returns the value of property "SVGFEDiffuseLightingElement.kernelUnitLengthY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) KernelUnitLengthY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEDiffuseLightingElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEDiffuseLightingElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEDiffuseLightingElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEDiffuseLightingElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEDiffuseLightingElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDiffuseLightingElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDiffuseLightingElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// In2 returns the value of property "SVGFEDisplacementMapElement.in2".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) In2() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementIn2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Scale returns the value of property "SVGFEDisplacementMapElement.scale".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) Scale() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// XChannelSelector returns the value of property "SVGFEDisplacementMapElement.xChannelSelector".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) XChannelSelector() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementXChannelSelector(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// YChannelSelector returns the value of property "SVGFEDisplacementMapElement.yChannelSelector".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) YChannelSelector() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementYChannelSelector(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEDisplacementMapElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEDisplacementMapElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEDisplacementMapElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEDisplacementMapElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEDisplacementMapElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDisplacementMapElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDisplacementMapElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEDistantLightElement) Azimuth() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDistantLightElementAzimuth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Elevation returns the value of property "SVGFEDistantLightElement.elevation".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDistantLightElement) Elevation() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDistantLightElementElevation(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Dx returns the value of property "SVGFEDropShadowElement.dx".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Dx() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementDx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Dy returns the value of property "SVGFEDropShadowElement.dy".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Dy() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementDy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// StdDeviationX returns the value of property "SVGFEDropShadowElement.stdDeviationX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementStdDeviationX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// StdDeviationY returns the value of property "SVGFEDropShadowElement.stdDeviationY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) StdDeviationY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementStdDeviationY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEDropShadowElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEDropShadowElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEDropShadowElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEDropShadowElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEDropShadowElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEDropShadowElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEDropShadowElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// SetStdDeviation calls the method "SVGFEDropShadowElement.setStdDeviation".
//
// The returned bool will be false if there is no such method.
func (this SVGFEDropShadowElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGFEDropShadowElementSetStdDeviation(
		this.Ref(), js.Pointer(&_ok),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStdDeviationFunc returns the method "SVGFEDropShadowElement.setStdDeviation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGFEDropShadowElement) SetStdDeviationFunc() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	return fn.FromRef(
		bindings.SVGFEDropShadowElementSetStdDeviationFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGFEFloodElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEFloodElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEFloodElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEFloodElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEFloodElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEFloodElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEFloodElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEFloodElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEFloodElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEFloodElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEFloodElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEFloodElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEFloodElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEFloodElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// StdDeviationX returns the value of property "SVGFEGaussianBlurElement.stdDeviationX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementStdDeviationX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// StdDeviationY returns the value of property "SVGFEGaussianBlurElement.stdDeviationY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) StdDeviationY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementStdDeviationY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// EdgeMode returns the value of property "SVGFEGaussianBlurElement.edgeMode".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) EdgeMode() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementEdgeMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEGaussianBlurElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEGaussianBlurElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEGaussianBlurElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEGaussianBlurElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEGaussianBlurElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEGaussianBlurElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEGaussianBlurElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// SetStdDeviation calls the method "SVGFEGaussianBlurElement.setStdDeviation".
//
// The returned bool will be false if there is no such method.
func (this SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGFEGaussianBlurElementSetStdDeviation(
		this.Ref(), js.Pointer(&_ok),
		float32(stdDeviationX),
		float32(stdDeviationY),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStdDeviationFunc returns the method "SVGFEGaussianBlurElement.setStdDeviation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGFEGaussianBlurElement) SetStdDeviationFunc() (fn js.Func[func(stdDeviationX float32, stdDeviationY float32)]) {
	return fn.FromRef(
		bindings.SVGFEGaussianBlurElementSetStdDeviationFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

// CrossOrigin returns the value of property "SVGFEImageElement.crossOrigin".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) CrossOrigin() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEImageElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEImageElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEImageElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEImageElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEImageElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGFEImageElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGFEImageElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEImageElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEMergeElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEMergeElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMergeElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEMergeElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMergeElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEMergeElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMergeElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEMergeElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMergeElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEMergeNodeElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMergeNodeElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Operator returns the value of property "SVGFEMorphologyElement.operator".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) Operator() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementOperator(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// RadiusX returns the value of property "SVGFEMorphologyElement.radiusX".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) RadiusX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementRadiusX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// RadiusY returns the value of property "SVGFEMorphologyElement.radiusY".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) RadiusY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementRadiusY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEMorphologyElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEMorphologyElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEMorphologyElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEMorphologyElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEMorphologyElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEMorphologyElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEMorphologyElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// Dx returns the value of property "SVGFEOffsetElement.dx".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Dx() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementDx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Dy returns the value of property "SVGFEOffsetElement.dy".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Dy() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementDy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFEOffsetElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEOffsetElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFEOffsetElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFEOffsetElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFEOffsetElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFEOffsetElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEOffsetElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFEPointLightElement) X() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEPointLightElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFEPointLightElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFEPointLightElement) Y() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEPointLightElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Z returns the value of property "SVGFEPointLightElement.z".
//
// The returned bool will be false if there is no such property.
func (this SVGFEPointLightElement) Z() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFEPointLightElementZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// SurfaceScale returns the value of property "SVGFESpecularLightingElement.surfaceScale".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) SurfaceScale() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementSurfaceScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// SpecularConstant returns the value of property "SVGFESpecularLightingElement.specularConstant".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularConstant() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementSpecularConstant(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// SpecularExponent returns the value of property "SVGFESpecularLightingElement.specularExponent".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) SpecularExponent() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementSpecularExponent(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// KernelUnitLengthX returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthX".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementKernelUnitLengthX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// KernelUnitLengthY returns the value of property "SVGFESpecularLightingElement.kernelUnitLengthY".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) KernelUnitLengthY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementKernelUnitLengthY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFESpecularLightingElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFESpecularLightingElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFESpecularLightingElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFESpecularLightingElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFESpecularLightingElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpecularLightingElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpecularLightingElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) X() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFESpotLightElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) Y() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// Z returns the value of property "SVGFESpotLightElement.z".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) Z() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// PointsAtX returns the value of property "SVGFESpotLightElement.pointsAtX".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) PointsAtX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementPointsAtX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// PointsAtY returns the value of property "SVGFESpotLightElement.pointsAtY".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) PointsAtY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementPointsAtY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// PointsAtZ returns the value of property "SVGFESpotLightElement.pointsAtZ".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) PointsAtZ() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementPointsAtZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// SpecularExponent returns the value of property "SVGFESpotLightElement.specularExponent".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) SpecularExponent() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementSpecularExponent(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// LimitingConeAngle returns the value of property "SVGFESpotLightElement.limitingConeAngle".
//
// The returned bool will be false if there is no such property.
func (this SVGFESpotLightElement) LimitingConeAngle() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFESpotLightElementLimitingConeAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

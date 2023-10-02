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

type OneOf_HTMLCollection_Element struct {
	ref js.Ref
}

func (x OneOf_HTMLCollection_Element) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLCollection_Element) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLCollection_Element) FromRef(ref js.Ref) OneOf_HTMLCollection_Element {
	return OneOf_HTMLCollection_Element{
		ref: ref,
	}
}

func (x OneOf_HTMLCollection_Element) HTMLCollection() HTMLCollection {
	return HTMLCollection{}.FromRef(x.ref)
}

func (x OneOf_HTMLCollection_Element) Element() Element {
	return Element{}.FromRef(x.ref)
}

type HTMLAllCollection struct {
	ref js.Ref
}

func (this HTMLAllCollection) Once() HTMLAllCollection {
	this.Ref().Once()
	return this
}

func (this HTMLAllCollection) Ref() js.Ref {
	return this.ref
}

func (this HTMLAllCollection) FromRef(ref js.Ref) HTMLAllCollection {
	this.ref = ref
	return this
}

func (this HTMLAllCollection) Free() {
	this.Ref().Free()
}

// Length returns the value of property "HTMLAllCollection.length".
//
// The returned bool will be false if there is no such property.
func (this HTMLAllCollection) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAllCollectionLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "HTMLAllCollection.".
//
// The returned bool will be false if there is no such method.
func (this HTMLAllCollection) Get(index uint32) (Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLAllCollectionGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Element{}.FromRef(_ret), _ok
}

// GetFunc returns the method "HTMLAllCollection.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLAllCollection) GetFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionGetFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLAllCollection.namedItem".
//
// The returned bool will be false if there is no such method.
func (this HTMLAllCollection) NamedItem(name js.String) (OneOf_HTMLCollection_Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLAllCollectionNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return OneOf_HTMLCollection_Element{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "HTMLAllCollection.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLAllCollection) NamedItemFunc() (fn js.Func[func(name js.String) OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "HTMLAllCollection.item".
//
// The returned bool will be false if there is no such method.
func (this HTMLAllCollection) Item(nameOrIndex js.String) (OneOf_HTMLCollection_Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLAllCollectionItem(
		this.Ref(), js.Pointer(&_ok),
		nameOrIndex.Ref(),
	)

	return OneOf_HTMLCollection_Element{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "HTMLAllCollection.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLAllCollection) ItemFunc() (fn js.Func[func(nameOrIndex js.String) OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionItemFunc(
			this.Ref(),
		),
	)
}

// Item1 calls the method "HTMLAllCollection.item".
//
// The returned bool will be false if there is no such method.
func (this HTMLAllCollection) Item1() (OneOf_HTMLCollection_Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLAllCollectionItem1(
		this.Ref(), js.Pointer(&_ok),
	)

	return OneOf_HTMLCollection_Element{}.FromRef(_ret), _ok
}

// Item1Func returns the method "HTMLAllCollection.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLAllCollection) Item1Func() (fn js.Func[func() OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionItem1Func(
			this.Ref(),
		),
	)
}

type DOMStringList struct {
	ref js.Ref
}

func (this DOMStringList) Once() DOMStringList {
	this.Ref().Once()
	return this
}

func (this DOMStringList) Ref() js.Ref {
	return this.ref
}

func (this DOMStringList) FromRef(ref js.Ref) DOMStringList {
	this.ref = ref
	return this
}

func (this DOMStringList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "DOMStringList.length".
//
// The returned bool will be false if there is no such property.
func (this DOMStringList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDOMStringListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "DOMStringList.item".
//
// The returned bool will be false if there is no such method.
func (this DOMStringList) Item(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDOMStringListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "DOMStringList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMStringList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.DOMStringListItemFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "DOMStringList.contains".
//
// The returned bool will be false if there is no such method.
func (this DOMStringList) Contains(string js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDOMStringListContains(
		this.Ref(), js.Pointer(&_ok),
		string.Ref(),
	)

	return _ret == js.True, _ok
}

// ContainsFunc returns the method "DOMStringList.contains".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMStringList) ContainsFunc() (fn js.Func[func(string js.String) bool]) {
	return fn.FromRef(
		bindings.DOMStringListContainsFunc(
			this.Ref(),
		),
	)
}

type Location struct {
	ref js.Ref
}

func (this Location) Once() Location {
	this.Ref().Once()
	return this
}

func (this Location) Ref() js.Ref {
	return this.ref
}

func (this Location) FromRef(ref js.Ref) Location {
	this.ref = ref
	return this
}

func (this Location) Free() {
	this.Ref().Free()
}

// Href returns the value of property "Location.href".
//
// The returned bool will be false if there is no such property.
func (this Location) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "Location.href" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHref(val js.String) bool {
	return js.True == bindings.SetLocationHref(
		this.Ref(),
		val.Ref(),
	)
}

// Origin returns the value of property "Location.origin".
//
// The returned bool will be false if there is no such property.
func (this Location) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "Location.protocol".
//
// The returned bool will be false if there is no such property.
func (this Location) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol sets the value of property "Location.protocol" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetProtocol(val js.String) bool {
	return js.True == bindings.SetLocationProtocol(
		this.Ref(),
		val.Ref(),
	)
}

// Host returns the value of property "Location.host".
//
// The returned bool will be false if there is no such property.
func (this Location) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host sets the value of property "Location.host" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHost(val js.String) bool {
	return js.True == bindings.SetLocationHost(
		this.Ref(),
		val.Ref(),
	)
}

// Hostname returns the value of property "Location.hostname".
//
// The returned bool will be false if there is no such property.
func (this Location) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname sets the value of property "Location.hostname" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHostname(val js.String) bool {
	return js.True == bindings.SetLocationHostname(
		this.Ref(),
		val.Ref(),
	)
}

// Port returns the value of property "Location.port".
//
// The returned bool will be false if there is no such property.
func (this Location) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port sets the value of property "Location.port" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetPort(val js.String) bool {
	return js.True == bindings.SetLocationPort(
		this.Ref(),
		val.Ref(),
	)
}

// Pathname returns the value of property "Location.pathname".
//
// The returned bool will be false if there is no such property.
func (this Location) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname sets the value of property "Location.pathname" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetPathname(val js.String) bool {
	return js.True == bindings.SetLocationPathname(
		this.Ref(),
		val.Ref(),
	)
}

// Search returns the value of property "Location.search".
//
// The returned bool will be false if there is no such property.
func (this Location) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search sets the value of property "Location.search" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetSearch(val js.String) bool {
	return js.True == bindings.SetLocationSearch(
		this.Ref(),
		val.Ref(),
	)
}

// Hash returns the value of property "Location.hash".
//
// The returned bool will be false if there is no such property.
func (this Location) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLocationHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash sets the value of property "Location.hash" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHash(val js.String) bool {
	return js.True == bindings.SetLocationHash(
		this.Ref(),
		val.Ref(),
	)
}

// AncestorOrigins returns the value of property "Location.ancestorOrigins".
//
// The returned bool will be false if there is no such property.
func (this Location) AncestorOrigins() (DOMStringList, bool) {
	var _ok bool
	_ret := bindings.GetLocationAncestorOrigins(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringList{}.FromRef(_ret), _ok
}

// Assign calls the method "Location.assign".
//
// The returned bool will be false if there is no such method.
func (this Location) Assign(url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallLocationAssign(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AssignFunc returns the method "Location.assign".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Location) AssignFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.LocationAssignFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "Location.replace".
//
// The returned bool will be false if there is no such method.
func (this Location) Replace(url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallLocationReplace(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceFunc returns the method "Location.replace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Location) ReplaceFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.LocationReplaceFunc(
			this.Ref(),
		),
	)
}

// Reload calls the method "Location.reload".
//
// The returned bool will be false if there is no such method.
func (this Location) Reload() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallLocationReload(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReloadFunc returns the method "Location.reload".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Location) ReloadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.LocationReloadFunc(
			this.Ref(),
		),
	)
}

type DocumentReadyState uint32

const (
	_ DocumentReadyState = iota

	DocumentReadyState_LOADING
	DocumentReadyState_INTERACTIVE
	DocumentReadyState_COMPLETE
)

func (DocumentReadyState) FromRef(str js.Ref) DocumentReadyState {
	return DocumentReadyState(bindings.ConstOfDocumentReadyState(str))
}

func (x DocumentReadyState) String() (string, bool) {
	switch x {
	case DocumentReadyState_LOADING:
		return "loading", true
	case DocumentReadyState_INTERACTIVE:
		return "interactive", true
	case DocumentReadyState_COMPLETE:
		return "complete", true
	default:
		return "", false
	}
}

type EndingType uint32

const (
	_ EndingType = iota

	EndingType_TRANSPARENT
	EndingType_NATIVE
)

func (EndingType) FromRef(str js.Ref) EndingType {
	return EndingType(bindings.ConstOfEndingType(str))
}

func (x EndingType) String() (string, bool) {
	switch x {
	case EndingType_TRANSPARENT:
		return "transparent", true
	case EndingType_NATIVE:
		return "native", true
	default:
		return "", false
	}
}

type BlobPropertyBag struct {
	// Type is "BlobPropertyBag.type"
	//
	// Optional, defaults to "".
	Type js.String
	// Endings is "BlobPropertyBag.endings"
	//
	// Optional, defaults to "transparent".
	Endings EndingType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BlobPropertyBag with all fields set.
func (p BlobPropertyBag) FromRef(ref js.Ref) BlobPropertyBag {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BlobPropertyBag in the application heap.
func (p BlobPropertyBag) New() js.Ref {
	return bindings.BlobPropertyBagJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BlobPropertyBag) UpdateFrom(ref js.Ref) {
	bindings.BlobPropertyBagJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BlobPropertyBag) Update(ref js.Ref) {
	bindings.BlobPropertyBagJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type QueuingStrategySizeFunc func(this js.Ref, chunk js.Any) js.Ref

func (fn QueuingStrategySizeFunc) Register() js.Func[func(chunk js.Any) float64] {
	return js.RegisterCallback[func(chunk js.Any) float64](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn QueuingStrategySizeFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QueuingStrategySize[T any] struct {
	Fn  func(arg T, this js.Ref, chunk js.Any) js.Ref
	Arg T
}

func (cb *QueuingStrategySize[T]) Register() js.Func[func(chunk js.Any) float64] {
	return js.RegisterCallback[func(chunk js.Any) float64](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *QueuingStrategySize[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QueuingStrategy struct {
	// HighWaterMark is "QueuingStrategy.highWaterMark"
	//
	// Optional
	//
	// NOTE: FFI_USE_HighWaterMark MUST be set to true to make this field effective.
	HighWaterMark float64
	// Size is "QueuingStrategy.size"
	//
	// Optional
	Size js.Func[func(chunk js.Any) float64]

	FFI_USE_HighWaterMark bool // for HighWaterMark.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueuingStrategy with all fields set.
func (p QueuingStrategy) FromRef(ref js.Ref) QueuingStrategy {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueuingStrategy in the application heap.
func (p QueuingStrategy) New() js.Ref {
	return bindings.QueuingStrategyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p QueuingStrategy) UpdateFrom(ref js.Ref) {
	bindings.QueuingStrategyJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p QueuingStrategy) Update(ref js.Ref) {
	bindings.QueuingStrategyJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReadableStreamReadResult struct {
	// Value is "ReadableStreamReadResult.value"
	//
	// Optional
	Value js.Any
	// Done is "ReadableStreamReadResult.done"
	//
	// Optional
	//
	// NOTE: FFI_USE_Done MUST be set to true to make this field effective.
	Done bool

	FFI_USE_Done bool // for Done.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadableStreamReadResult with all fields set.
func (p ReadableStreamReadResult) FromRef(ref js.Ref) ReadableStreamReadResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadableStreamReadResult in the application heap.
func (p ReadableStreamReadResult) New() js.Ref {
	return bindings.ReadableStreamReadResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReadableStreamReadResult) UpdateFrom(ref js.Ref) {
	bindings.ReadableStreamReadResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReadableStreamReadResult) Update(ref js.Ref) {
	bindings.ReadableStreamReadResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewReadableStreamDefaultReader(stream ReadableStream) ReadableStreamDefaultReader {
	return ReadableStreamDefaultReader{}.FromRef(
		bindings.NewReadableStreamDefaultReaderByReadableStreamDefaultReader(
			stream.Ref()),
	)
}

type ReadableStreamDefaultReader struct {
	ref js.Ref
}

func (this ReadableStreamDefaultReader) Once() ReadableStreamDefaultReader {
	this.Ref().Once()
	return this
}

func (this ReadableStreamDefaultReader) Ref() js.Ref {
	return this.ref
}

func (this ReadableStreamDefaultReader) FromRef(ref js.Ref) ReadableStreamDefaultReader {
	this.ref = ref
	return this
}

func (this ReadableStreamDefaultReader) Free() {
	this.Ref().Free()
}

// Closed returns the value of property "ReadableStreamDefaultReader.closed".
//
// The returned bool will be false if there is no such property.
func (this ReadableStreamDefaultReader) Closed() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetReadableStreamDefaultReaderClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Read calls the method "ReadableStreamDefaultReader.read".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultReader) Read() (js.Promise[ReadableStreamReadResult], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultReaderRead(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ReadableStreamReadResult]{}.FromRef(_ret), _ok
}

// ReadFunc returns the method "ReadableStreamDefaultReader.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultReader) ReadFunc() (fn js.Func[func() js.Promise[ReadableStreamReadResult]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderReadFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "ReadableStreamDefaultReader.releaseLock".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultReader) ReleaseLock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultReaderReleaseLock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleaseLockFunc returns the method "ReadableStreamDefaultReader.releaseLock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultReader) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderReleaseLockFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStreamDefaultReader.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultReader) Cancel(reason js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultReaderCancel(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CancelFunc returns the method "ReadableStreamDefaultReader.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultReader) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStreamDefaultReader.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamDefaultReader) Cancel1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamDefaultReaderCancel1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Cancel1Func returns the method "ReadableStreamDefaultReader.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamDefaultReader) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderCancel1Func(
			this.Ref(),
		),
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

type ArrayBufferView = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView

func NewReadableStreamBYOBReader(stream ReadableStream) ReadableStreamBYOBReader {
	return ReadableStreamBYOBReader{}.FromRef(
		bindings.NewReadableStreamBYOBReaderByReadableStreamBYOBReader(
			stream.Ref()),
	)
}

type ReadableStreamBYOBReader struct {
	ref js.Ref
}

func (this ReadableStreamBYOBReader) Once() ReadableStreamBYOBReader {
	this.Ref().Once()
	return this
}

func (this ReadableStreamBYOBReader) Ref() js.Ref {
	return this.ref
}

func (this ReadableStreamBYOBReader) FromRef(ref js.Ref) ReadableStreamBYOBReader {
	this.ref = ref
	return this
}

func (this ReadableStreamBYOBReader) Free() {
	this.Ref().Free()
}

// Closed returns the value of property "ReadableStreamBYOBReader.closed".
//
// The returned bool will be false if there is no such property.
func (this ReadableStreamBYOBReader) Closed() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetReadableStreamBYOBReaderClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Read calls the method "ReadableStreamBYOBReader.read".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBReader) Read(view ArrayBufferView) (js.Promise[ReadableStreamReadResult], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBReaderRead(
		this.Ref(), js.Pointer(&_ok),
		view.Ref(),
	)

	return js.Promise[ReadableStreamReadResult]{}.FromRef(_ret), _ok
}

// ReadFunc returns the method "ReadableStreamBYOBReader.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBReader) ReadFunc() (fn js.Func[func(view ArrayBufferView) js.Promise[ReadableStreamReadResult]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderReadFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "ReadableStreamBYOBReader.releaseLock".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBReader) ReleaseLock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBReaderReleaseLock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleaseLockFunc returns the method "ReadableStreamBYOBReader.releaseLock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBReader) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderReleaseLockFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStreamBYOBReader.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBReader) Cancel(reason js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBReaderCancel(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CancelFunc returns the method "ReadableStreamBYOBReader.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBReader) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStreamBYOBReader.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStreamBYOBReader) Cancel1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamBYOBReaderCancel1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Cancel1Func returns the method "ReadableStreamBYOBReader.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStreamBYOBReader) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderCancel1Func(
			this.Ref(),
		),
	)
}

type OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader struct {
	ref js.Ref
}

func (x OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader) Free() {
	x.ref.Free()
}

func (x OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader) FromRef(ref js.Ref) OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader {
	return OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader{
		ref: ref,
	}
}

func (x OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader) ReadableStreamDefaultReader() ReadableStreamDefaultReader {
	return ReadableStreamDefaultReader{}.FromRef(x.ref)
}

func (x OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader) ReadableStreamBYOBReader() ReadableStreamBYOBReader {
	return ReadableStreamBYOBReader{}.FromRef(x.ref)
}

type ReadableStreamReader = OneOf_ReadableStreamDefaultReader_ReadableStreamBYOBReader

type ReadableStreamReaderMode uint32

const (
	_ ReadableStreamReaderMode = iota

	ReadableStreamReaderMode_BYOB
)

func (ReadableStreamReaderMode) FromRef(str js.Ref) ReadableStreamReaderMode {
	return ReadableStreamReaderMode(bindings.ConstOfReadableStreamReaderMode(str))
}

func (x ReadableStreamReaderMode) String() (string, bool) {
	switch x {
	case ReadableStreamReaderMode_BYOB:
		return "byob", true
	default:
		return "", false
	}
}

type ReadableStreamGetReaderOptions struct {
	// Mode is "ReadableStreamGetReaderOptions.mode"
	//
	// Optional
	Mode ReadableStreamReaderMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadableStreamGetReaderOptions with all fields set.
func (p ReadableStreamGetReaderOptions) FromRef(ref js.Ref) ReadableStreamGetReaderOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadableStreamGetReaderOptions in the application heap.
func (p ReadableStreamGetReaderOptions) New() js.Ref {
	return bindings.ReadableStreamGetReaderOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReadableStreamGetReaderOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadableStreamGetReaderOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReadableStreamGetReaderOptions) Update(ref js.Ref) {
	bindings.ReadableStreamGetReaderOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWritableStreamDefaultWriter(stream WritableStream) WritableStreamDefaultWriter {
	return WritableStreamDefaultWriter{}.FromRef(
		bindings.NewWritableStreamDefaultWriterByWritableStreamDefaultWriter(
			stream.Ref()),
	)
}

type WritableStreamDefaultWriter struct {
	ref js.Ref
}

func (this WritableStreamDefaultWriter) Once() WritableStreamDefaultWriter {
	this.Ref().Once()
	return this
}

func (this WritableStreamDefaultWriter) Ref() js.Ref {
	return this.ref
}

func (this WritableStreamDefaultWriter) FromRef(ref js.Ref) WritableStreamDefaultWriter {
	this.ref = ref
	return this
}

func (this WritableStreamDefaultWriter) Free() {
	this.Ref().Free()
}

// Closed returns the value of property "WritableStreamDefaultWriter.closed".
//
// The returned bool will be false if there is no such property.
func (this WritableStreamDefaultWriter) Closed() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetWritableStreamDefaultWriterClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DesiredSize returns the value of property "WritableStreamDefaultWriter.desiredSize".
//
// The returned bool will be false if there is no such property.
func (this WritableStreamDefaultWriter) DesiredSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWritableStreamDefaultWriterDesiredSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Ready returns the value of property "WritableStreamDefaultWriter.ready".
//
// The returned bool will be false if there is no such property.
func (this WritableStreamDefaultWriter) Ready() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetWritableStreamDefaultWriterReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Abort calls the method "WritableStreamDefaultWriter.abort".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) Abort(reason js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterAbort(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AbortFunc returns the method "WritableStreamDefaultWriter.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) AbortFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterAbortFunc(
			this.Ref(),
		),
	)
}

// Abort1 calls the method "WritableStreamDefaultWriter.abort".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) Abort1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterAbort1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Abort1Func returns the method "WritableStreamDefaultWriter.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) Abort1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterAbort1Func(
			this.Ref(),
		),
	)
}

// Close calls the method "WritableStreamDefaultWriter.close".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "WritableStreamDefaultWriter.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterCloseFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "WritableStreamDefaultWriter.releaseLock".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) ReleaseLock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterReleaseLock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleaseLockFunc returns the method "WritableStreamDefaultWriter.releaseLock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterReleaseLockFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "WritableStreamDefaultWriter.write".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) Write(chunk js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterWrite(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteFunc returns the method "WritableStreamDefaultWriter.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) WriteFunc() (fn js.Func[func(chunk js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterWriteFunc(
			this.Ref(),
		),
	)
}

// Write1 calls the method "WritableStreamDefaultWriter.write".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultWriter) Write1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultWriterWrite1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Write1Func returns the method "WritableStreamDefaultWriter.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultWriter) Write1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterWrite1Func(
			this.Ref(),
		),
	)
}

func NewWritableStream(underlyingSink js.Object, strategy QueuingStrategy) WritableStream {
	return WritableStream{}.FromRef(
		bindings.NewWritableStreamByWritableStream(
			underlyingSink.Ref(),
			js.Pointer(&strategy)),
	)
}

func NewWritableStreamByWritableStream1(underlyingSink js.Object) WritableStream {
	return WritableStream{}.FromRef(
		bindings.NewWritableStreamByWritableStream1(
			underlyingSink.Ref()),
	)
}

func NewWritableStreamByWritableStream2() WritableStream {
	return WritableStream{}.FromRef(
		bindings.NewWritableStreamByWritableStream2(),
	)
}

type WritableStream struct {
	ref js.Ref
}

func (this WritableStream) Once() WritableStream {
	this.Ref().Once()
	return this
}

func (this WritableStream) Ref() js.Ref {
	return this.ref
}

func (this WritableStream) FromRef(ref js.Ref) WritableStream {
	this.ref = ref
	return this
}

func (this WritableStream) Free() {
	this.Ref().Free()
}

// Locked returns the value of property "WritableStream.locked".
//
// The returned bool will be false if there is no such property.
func (this WritableStream) Locked() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWritableStreamLocked(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Abort calls the method "WritableStream.abort".
//
// The returned bool will be false if there is no such method.
func (this WritableStream) Abort(reason js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamAbort(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AbortFunc returns the method "WritableStream.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStream) AbortFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamAbortFunc(
			this.Ref(),
		),
	)
}

// Abort1 calls the method "WritableStream.abort".
//
// The returned bool will be false if there is no such method.
func (this WritableStream) Abort1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamAbort1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Abort1Func returns the method "WritableStream.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStream) Abort1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamAbort1Func(
			this.Ref(),
		),
	)
}

// Close calls the method "WritableStream.close".
//
// The returned bool will be false if there is no such method.
func (this WritableStream) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "WritableStream.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStream) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamCloseFunc(
			this.Ref(),
		),
	)
}

// GetWriter calls the method "WritableStream.getWriter".
//
// The returned bool will be false if there is no such method.
func (this WritableStream) GetWriter() (WritableStreamDefaultWriter, bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamGetWriter(
		this.Ref(), js.Pointer(&_ok),
	)

	return WritableStreamDefaultWriter{}.FromRef(_ret), _ok
}

// GetWriterFunc returns the method "WritableStream.getWriter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStream) GetWriterFunc() (fn js.Func[func() WritableStreamDefaultWriter]) {
	return fn.FromRef(
		bindings.WritableStreamGetWriterFunc(
			this.Ref(),
		),
	)
}

type ReadableWritablePair struct {
	// Readable is "ReadableWritablePair.readable"
	//
	// Required
	Readable ReadableStream
	// Writable is "ReadableWritablePair.writable"
	//
	// Required
	Writable WritableStream

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadableWritablePair with all fields set.
func (p ReadableWritablePair) FromRef(ref js.Ref) ReadableWritablePair {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadableWritablePair in the application heap.
func (p ReadableWritablePair) New() js.Ref {
	return bindings.ReadableWritablePairJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReadableWritablePair) UpdateFrom(ref js.Ref) {
	bindings.ReadableWritablePairJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReadableWritablePair) Update(ref js.Ref) {
	bindings.ReadableWritablePairJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type StreamPipeOptions struct {
	// PreventClose is "StreamPipeOptions.preventClose"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreventClose MUST be set to true to make this field effective.
	PreventClose bool
	// PreventAbort is "StreamPipeOptions.preventAbort"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreventAbort MUST be set to true to make this field effective.
	PreventAbort bool
	// PreventCancel is "StreamPipeOptions.preventCancel"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreventCancel MUST be set to true to make this field effective.
	PreventCancel bool
	// Signal is "StreamPipeOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE_PreventClose  bool // for PreventClose.
	FFI_USE_PreventAbort  bool // for PreventAbort.
	FFI_USE_PreventCancel bool // for PreventCancel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StreamPipeOptions with all fields set.
func (p StreamPipeOptions) FromRef(ref js.Ref) StreamPipeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StreamPipeOptions in the application heap.
func (p StreamPipeOptions) New() js.Ref {
	return bindings.StreamPipeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p StreamPipeOptions) UpdateFrom(ref js.Ref) {
	bindings.StreamPipeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StreamPipeOptions) Update(ref js.Ref) {
	bindings.StreamPipeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewReadableStream(underlyingSource js.Object, strategy QueuingStrategy) ReadableStream {
	return ReadableStream{}.FromRef(
		bindings.NewReadableStreamByReadableStream(
			underlyingSource.Ref(),
			js.Pointer(&strategy)),
	)
}

func NewReadableStreamByReadableStream1(underlyingSource js.Object) ReadableStream {
	return ReadableStream{}.FromRef(
		bindings.NewReadableStreamByReadableStream1(
			underlyingSource.Ref()),
	)
}

func NewReadableStreamByReadableStream2() ReadableStream {
	return ReadableStream{}.FromRef(
		bindings.NewReadableStreamByReadableStream2(),
	)
}

type ReadableStream struct {
	ref js.Ref
}

func (this ReadableStream) Once() ReadableStream {
	this.Ref().Once()
	return this
}

func (this ReadableStream) Ref() js.Ref {
	return this.ref
}

func (this ReadableStream) FromRef(ref js.Ref) ReadableStream {
	this.ref = ref
	return this
}

func (this ReadableStream) Free() {
	this.Ref().Free()
}

// Locked returns the value of property "ReadableStream.locked".
//
// The returned bool will be false if there is no such property.
func (this ReadableStream) Locked() (bool, bool) {
	var _ok bool
	_ret := bindings.GetReadableStreamLocked(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// From calls the staticmethod "ReadableStream.from".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) From(asyncIterable js.Any) (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamFrom(
		this.Ref(), js.Pointer(&_ok),
		asyncIterable.Ref(),
	)

	return ReadableStream{}.FromRef(_ret), _ok
}

// FromFunc returns the staticmethod "ReadableStream.from".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) FromFunc() (fn js.Func[func(asyncIterable js.Any) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamFromFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStream.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) Cancel(reason js.Any) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamCancel(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CancelFunc returns the method "ReadableStream.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStream.cancel".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) Cancel1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamCancel1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Cancel1Func returns the method "ReadableStream.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamCancel1Func(
			this.Ref(),
		),
	)
}

// GetReader calls the method "ReadableStream.getReader".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) GetReader(options ReadableStreamGetReaderOptions) (ReadableStreamReader, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamGetReader(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return ReadableStreamReader{}.FromRef(_ret), _ok
}

// GetReaderFunc returns the method "ReadableStream.getReader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) GetReaderFunc() (fn js.Func[func(options ReadableStreamGetReaderOptions) ReadableStreamReader]) {
	return fn.FromRef(
		bindings.ReadableStreamGetReaderFunc(
			this.Ref(),
		),
	)
}

// GetReader1 calls the method "ReadableStream.getReader".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) GetReader1() (ReadableStreamReader, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamGetReader1(
		this.Ref(), js.Pointer(&_ok),
	)

	return ReadableStreamReader{}.FromRef(_ret), _ok
}

// GetReader1Func returns the method "ReadableStream.getReader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) GetReader1Func() (fn js.Func[func() ReadableStreamReader]) {
	return fn.FromRef(
		bindings.ReadableStreamGetReader1Func(
			this.Ref(),
		),
	)
}

// PipeThrough calls the method "ReadableStream.pipeThrough".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) PipeThrough(transform ReadableWritablePair, options StreamPipeOptions) (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamPipeThrough(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
		js.Pointer(&options),
	)

	return ReadableStream{}.FromRef(_ret), _ok
}

// PipeThroughFunc returns the method "ReadableStream.pipeThrough".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) PipeThroughFunc() (fn js.Func[func(transform ReadableWritablePair, options StreamPipeOptions) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeThroughFunc(
			this.Ref(),
		),
	)
}

// PipeThrough1 calls the method "ReadableStream.pipeThrough".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) PipeThrough1(transform ReadableWritablePair) (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamPipeThrough1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
	)

	return ReadableStream{}.FromRef(_ret), _ok
}

// PipeThrough1Func returns the method "ReadableStream.pipeThrough".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) PipeThrough1Func() (fn js.Func[func(transform ReadableWritablePair) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeThrough1Func(
			this.Ref(),
		),
	)
}

// PipeTo calls the method "ReadableStream.pipeTo".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) PipeTo(destination WritableStream, options StreamPipeOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamPipeTo(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PipeToFunc returns the method "ReadableStream.pipeTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) PipeToFunc() (fn js.Func[func(destination WritableStream, options StreamPipeOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeToFunc(
			this.Ref(),
		),
	)
}

// PipeTo1 calls the method "ReadableStream.pipeTo".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) PipeTo1(destination WritableStream) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamPipeTo1(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PipeTo1Func returns the method "ReadableStream.pipeTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) PipeTo1Func() (fn js.Func[func(destination WritableStream) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeTo1Func(
			this.Ref(),
		),
	)
}

// Tee calls the method "ReadableStream.tee".
//
// The returned bool will be false if there is no such method.
func (this ReadableStream) Tee() (js.Array[ReadableStream], bool) {
	var _ok bool
	_ret := bindings.CallReadableStreamTee(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[ReadableStream]{}.FromRef(_ret), _ok
}

// TeeFunc returns the method "ReadableStream.tee".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ReadableStream) TeeFunc() (fn js.Func[func() js.Array[ReadableStream]]) {
	return fn.FromRef(
		bindings.ReadableStreamTeeFunc(
			this.Ref(),
		),
	)
}

func NewBlob(blobParts js.Array[BlobPart], options BlobPropertyBag) Blob {
	return Blob{}.FromRef(
		bindings.NewBlobByBlob(
			blobParts.Ref(),
			js.Pointer(&options)),
	)
}

func NewBlobByBlob1(blobParts js.Array[BlobPart]) Blob {
	return Blob{}.FromRef(
		bindings.NewBlobByBlob1(
			blobParts.Ref()),
	)
}

func NewBlobByBlob2() Blob {
	return Blob{}.FromRef(
		bindings.NewBlobByBlob2(),
	)
}

type Blob struct {
	ref js.Ref
}

func (this Blob) Once() Blob {
	this.Ref().Once()
	return this
}

func (this Blob) Ref() js.Ref {
	return this.ref
}

func (this Blob) FromRef(ref js.Ref) Blob {
	this.ref = ref
	return this
}

func (this Blob) Free() {
	this.Ref().Free()
}

// Size returns the value of property "Blob.size".
//
// The returned bool will be false if there is no such property.
func (this Blob) Size() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetBlobSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Type returns the value of property "Blob.type".
//
// The returned bool will be false if there is no such property.
func (this Blob) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBlobType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Slice calls the method "Blob.slice".
//
// The returned bool will be false if there is no such method.
func (this Blob) Slice(start int64, end int64, contentType js.String) (Blob, bool) {
	var _ok bool
	_ret := bindings.CallBlobSlice(
		this.Ref(), js.Pointer(&_ok),
		float64(start),
		float64(end),
		contentType.Ref(),
	)

	return Blob{}.FromRef(_ret), _ok
}

// SliceFunc returns the method "Blob.slice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) SliceFunc() (fn js.Func[func(start int64, end int64, contentType js.String) Blob]) {
	return fn.FromRef(
		bindings.BlobSliceFunc(
			this.Ref(),
		),
	)
}

// Slice1 calls the method "Blob.slice".
//
// The returned bool will be false if there is no such method.
func (this Blob) Slice1(start int64, end int64) (Blob, bool) {
	var _ok bool
	_ret := bindings.CallBlobSlice1(
		this.Ref(), js.Pointer(&_ok),
		float64(start),
		float64(end),
	)

	return Blob{}.FromRef(_ret), _ok
}

// Slice1Func returns the method "Blob.slice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) Slice1Func() (fn js.Func[func(start int64, end int64) Blob]) {
	return fn.FromRef(
		bindings.BlobSlice1Func(
			this.Ref(),
		),
	)
}

// Slice2 calls the method "Blob.slice".
//
// The returned bool will be false if there is no such method.
func (this Blob) Slice2(start int64) (Blob, bool) {
	var _ok bool
	_ret := bindings.CallBlobSlice2(
		this.Ref(), js.Pointer(&_ok),
		float64(start),
	)

	return Blob{}.FromRef(_ret), _ok
}

// Slice2Func returns the method "Blob.slice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) Slice2Func() (fn js.Func[func(start int64) Blob]) {
	return fn.FromRef(
		bindings.BlobSlice2Func(
			this.Ref(),
		),
	)
}

// Slice3 calls the method "Blob.slice".
//
// The returned bool will be false if there is no such method.
func (this Blob) Slice3() (Blob, bool) {
	var _ok bool
	_ret := bindings.CallBlobSlice3(
		this.Ref(), js.Pointer(&_ok),
	)

	return Blob{}.FromRef(_ret), _ok
}

// Slice3Func returns the method "Blob.slice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) Slice3Func() (fn js.Func[func() Blob]) {
	return fn.FromRef(
		bindings.BlobSlice3Func(
			this.Ref(),
		),
	)
}

// Stream calls the method "Blob.stream".
//
// The returned bool will be false if there is no such method.
func (this Blob) Stream() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.CallBlobStream(
		this.Ref(), js.Pointer(&_ok),
	)

	return ReadableStream{}.FromRef(_ret), _ok
}

// StreamFunc returns the method "Blob.stream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) StreamFunc() (fn js.Func[func() ReadableStream]) {
	return fn.FromRef(
		bindings.BlobStreamFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "Blob.text".
//
// The returned bool will be false if there is no such method.
func (this Blob) Text() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallBlobText(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// TextFunc returns the method "Blob.text".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.BlobTextFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Blob.arrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this Blob) ArrayBuffer() (js.Promise[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallBlobArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// ArrayBufferFunc returns the method "Blob.arrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Blob) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.BlobArrayBufferFunc(
			this.Ref(),
		),
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type BlobPart = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String

type FilePropertyBag struct {
	// LastModified is "FilePropertyBag.lastModified"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastModified MUST be set to true to make this field effective.
	LastModified int64
	// Type is "FilePropertyBag.type"
	//
	// Optional, defaults to "".
	Type js.String
	// Endings is "FilePropertyBag.endings"
	//
	// Optional, defaults to "transparent".
	Endings EndingType

	FFI_USE_LastModified bool // for LastModified.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FilePropertyBag with all fields set.
func (p FilePropertyBag) FromRef(ref js.Ref) FilePropertyBag {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FilePropertyBag in the application heap.
func (p FilePropertyBag) New() js.Ref {
	return bindings.FilePropertyBagJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FilePropertyBag) UpdateFrom(ref js.Ref) {
	bindings.FilePropertyBagJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FilePropertyBag) Update(ref js.Ref) {
	bindings.FilePropertyBagJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFile(fileBits js.Array[BlobPart], fileName js.String, options FilePropertyBag) File {
	return File{}.FromRef(
		bindings.NewFileByFile(
			fileBits.Ref(),
			fileName.Ref(),
			js.Pointer(&options)),
	)
}

func NewFileByFile1(fileBits js.Array[BlobPart], fileName js.String) File {
	return File{}.FromRef(
		bindings.NewFileByFile1(
			fileBits.Ref(),
			fileName.Ref()),
	)
}

type File struct {
	Blob
}

func (this File) Once() File {
	this.Ref().Once()
	return this
}

func (this File) Ref() js.Ref {
	return this.Blob.Ref()
}

func (this File) FromRef(ref js.Ref) File {
	this.Blob = this.Blob.FromRef(ref)
	return this
}

func (this File) Free() {
	this.Ref().Free()
}

// Name returns the value of property "File.name".
//
// The returned bool will be false if there is no such property.
func (this File) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LastModified returns the value of property "File.lastModified".
//
// The returned bool will be false if there is no such property.
func (this File) LastModified() (int64, bool) {
	var _ok bool
	_ret := bindings.GetFileLastModified(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// WebkitRelativePath returns the value of property "File.webkitRelativePath".
//
// The returned bool will be false if there is no such property.
func (this File) WebkitRelativePath() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileWebkitRelativePath(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type RadioNodeList struct {
	NodeList
}

func (this RadioNodeList) Once() RadioNodeList {
	this.Ref().Once()
	return this
}

func (this RadioNodeList) Ref() js.Ref {
	return this.NodeList.Ref()
}

func (this RadioNodeList) FromRef(ref js.Ref) RadioNodeList {
	this.NodeList = this.NodeList.FromRef(ref)
	return this
}

func (this RadioNodeList) Free() {
	this.Ref().Free()
}

// Value returns the value of property "RadioNodeList.value".
//
// The returned bool will be false if there is no such property.
func (this RadioNodeList) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRadioNodeListValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "RadioNodeList.value" to val.
//
// It returns false if the property cannot be set.
func (this RadioNodeList) SetValue(val js.String) bool {
	return js.True == bindings.SetRadioNodeListValue(
		this.Ref(),
		val.Ref(),
	)
}

type OneOf_RadioNodeList_Element struct {
	ref js.Ref
}

func (x OneOf_RadioNodeList_Element) Ref() js.Ref {
	return x.ref
}

func (x OneOf_RadioNodeList_Element) Free() {
	x.ref.Free()
}

func (x OneOf_RadioNodeList_Element) FromRef(ref js.Ref) OneOf_RadioNodeList_Element {
	return OneOf_RadioNodeList_Element{
		ref: ref,
	}
}

func (x OneOf_RadioNodeList_Element) RadioNodeList() RadioNodeList {
	return RadioNodeList{}.FromRef(x.ref)
}

func (x OneOf_RadioNodeList_Element) Element() Element {
	return Element{}.FromRef(x.ref)
}

type HTMLFormControlsCollection struct {
	HTMLCollection
}

func (this HTMLFormControlsCollection) Once() HTMLFormControlsCollection {
	this.Ref().Once()
	return this
}

func (this HTMLFormControlsCollection) Ref() js.Ref {
	return this.HTMLCollection.Ref()
}

func (this HTMLFormControlsCollection) FromRef(ref js.Ref) HTMLFormControlsCollection {
	this.HTMLCollection = this.HTMLCollection.FromRef(ref)
	return this
}

func (this HTMLFormControlsCollection) Free() {
	this.Ref().Free()
}

// NamedItem calls the method "HTMLFormControlsCollection.namedItem".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormControlsCollection) NamedItem(name js.String) (OneOf_RadioNodeList_Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormControlsCollectionNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return OneOf_RadioNodeList_Element{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "HTMLFormControlsCollection.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormControlsCollection) NamedItemFunc() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	return fn.FromRef(
		bindings.HTMLFormControlsCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

func NewHTMLFormElement() HTMLFormElement {
	return HTMLFormElement{}.FromRef(
		bindings.NewHTMLFormElementByHTMLFormElement(),
	)
}

type HTMLFormElement struct {
	HTMLElement
}

func (this HTMLFormElement) Once() HTMLFormElement {
	this.Ref().Once()
	return this
}

func (this HTMLFormElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFormElement) FromRef(ref js.Ref) HTMLFormElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFormElement) Free() {
	this.Ref().Free()
}

// AcceptCharset returns the value of property "HTMLFormElement.acceptCharset".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) AcceptCharset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementAcceptCharset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AcceptCharset sets the value of property "HTMLFormElement.acceptCharset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAcceptCharset(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAcceptCharset(
		this.Ref(),
		val.Ref(),
	)
}

// Action returns the value of property "HTMLFormElement.action".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Action() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Action sets the value of property "HTMLFormElement.action" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAction(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAction(
		this.Ref(),
		val.Ref(),
	)
}

// Autocomplete returns the value of property "HTMLFormElement.autocomplete".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Autocomplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementAutocomplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Autocomplete sets the value of property "HTMLFormElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAutocomplete(
		this.Ref(),
		val.Ref(),
	)
}

// Enctype returns the value of property "HTMLFormElement.enctype".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Enctype() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementEnctype(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Enctype sets the value of property "HTMLFormElement.enctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementEnctype(
		this.Ref(),
		val.Ref(),
	)
}

// Encoding returns the value of property "HTMLFormElement.encoding".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Encoding sets the value of property "HTMLFormElement.encoding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetEncoding(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementEncoding(
		this.Ref(),
		val.Ref(),
	)
}

// Method returns the value of property "HTMLFormElement.method".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Method() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Method sets the value of property "HTMLFormElement.method" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetMethod(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementMethod(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLFormElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLFormElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementName(
		this.Ref(),
		val.Ref(),
	)
}

// NoValidate returns the value of property "HTMLFormElement.noValidate".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) NoValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementNoValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoValidate sets the value of property "HTMLFormElement.noValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLFormElementNoValidate(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Target returns the value of property "HTMLFormElement.target".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target sets the value of property "HTMLFormElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLFormElement.rel".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Rel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementRel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rel sets the value of property "HTMLFormElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementRel(
		this.Ref(),
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLFormElement.relList".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) RelList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementRelList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Elements returns the value of property "HTMLFormElement.elements".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Elements() (HTMLFormControlsCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormControlsCollection{}.FromRef(_ret), _ok
}

// Length returns the value of property "HTMLFormElement.length".
//
// The returned bool will be false if there is no such property.
func (this HTMLFormElement) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFormElementLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "HTMLFormElement.".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) Get(index uint32) (Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return Element{}.FromRef(_ret), _ok
}

// GetFunc returns the method "HTMLFormElement.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) GetFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLFormElementGetFunc(
			this.Ref(),
		),
	)
}

// Get1 calls the method "HTMLFormElement.".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) Get1(name js.String) (OneOf_RadioNodeList_Element, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementGet1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return OneOf_RadioNodeList_Element{}.FromRef(_ret), _ok
}

// Get1Func returns the method "HTMLFormElement.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) Get1Func() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	return fn.FromRef(
		bindings.HTMLFormElementGet1Func(
			this.Ref(),
		),
	)
}

// Submit calls the method "HTMLFormElement.submit".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) Submit() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementSubmit(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SubmitFunc returns the method "HTMLFormElement.submit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) SubmitFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementSubmitFunc(
			this.Ref(),
		),
	)
}

// RequestSubmit calls the method "HTMLFormElement.requestSubmit".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) RequestSubmit(submitter HTMLElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementRequestSubmit(
		this.Ref(), js.Pointer(&_ok),
		submitter.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestSubmitFunc returns the method "HTMLFormElement.requestSubmit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) RequestSubmitFunc() (fn js.Func[func(submitter HTMLElement)]) {
	return fn.FromRef(
		bindings.HTMLFormElementRequestSubmitFunc(
			this.Ref(),
		),
	)
}

// RequestSubmit1 calls the method "HTMLFormElement.requestSubmit".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) RequestSubmit1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementRequestSubmit1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestSubmit1Func returns the method "HTMLFormElement.requestSubmit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) RequestSubmit1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementRequestSubmit1Func(
			this.Ref(),
		),
	)
}

// Reset calls the method "HTMLFormElement.reset".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "HTMLFormElement.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementResetFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLFormElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLFormElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFormElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLFormElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLFormElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFormElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLFormElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFormElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFormElementReportValidityFunc(
			this.Ref(),
		),
	)
}

type OneOf_File_String struct {
	ref js.Ref
}

func (x OneOf_File_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_File_String) Free() {
	x.ref.Free()
}

func (x OneOf_File_String) FromRef(ref js.Ref) OneOf_File_String {
	return OneOf_File_String{
		ref: ref,
	}
}

func (x OneOf_File_String) File() File {
	return File{}.FromRef(x.ref)
}

func (x OneOf_File_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type FormDataEntryValue = OneOf_File_String

func NewFormData(form HTMLFormElement, submitter HTMLElement) FormData {
	return FormData{}.FromRef(
		bindings.NewFormDataByFormData(
			form.Ref(),
			submitter.Ref()),
	)
}

func NewFormDataByFormData1(form HTMLFormElement) FormData {
	return FormData{}.FromRef(
		bindings.NewFormDataByFormData1(
			form.Ref()),
	)
}

func NewFormDataByFormData2() FormData {
	return FormData{}.FromRef(
		bindings.NewFormDataByFormData2(),
	)
}

type FormData struct {
	ref js.Ref
}

func (this FormData) Once() FormData {
	this.Ref().Once()
	return this
}

func (this FormData) Ref() js.Ref {
	return this.ref
}

func (this FormData) FromRef(ref js.Ref) FormData {
	this.ref = ref
	return this
}

func (this FormData) Free() {
	this.Ref().Free()
}

// Append calls the method "FormData.append".
//
// The returned bool will be false if there is no such method.
func (this FormData) Append(name js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataAppend(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "FormData.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) AppendFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.FormDataAppendFunc(
			this.Ref(),
		),
	)
}

// Append1 calls the method "FormData.append".
//
// The returned bool will be false if there is no such method.
func (this FormData) Append1(name js.String, blobValue Blob, filename js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataAppend1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Append1Func returns the method "FormData.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) Append1Func() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	return fn.FromRef(
		bindings.FormDataAppend1Func(
			this.Ref(),
		),
	)
}

// Append2 calls the method "FormData.append".
//
// The returned bool will be false if there is no such method.
func (this FormData) Append2(name js.String, blobValue Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataAppend2(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		blobValue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Append2Func returns the method "FormData.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) Append2Func() (fn js.Func[func(name js.String, blobValue Blob)]) {
	return fn.FromRef(
		bindings.FormDataAppend2Func(
			this.Ref(),
		),
	)
}

// Delete calls the method "FormData.delete".
//
// The returned bool will be false if there is no such method.
func (this FormData) Delete(name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataDelete(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFunc returns the method "FormData.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) DeleteFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.FormDataDeleteFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "FormData.get".
//
// The returned bool will be false if there is no such method.
func (this FormData) Get(name js.String) (FormDataEntryValue, bool) {
	var _ok bool
	_ret := bindings.CallFormDataGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return FormDataEntryValue{}.FromRef(_ret), _ok
}

// GetFunc returns the method "FormData.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) GetFunc() (fn js.Func[func(name js.String) FormDataEntryValue]) {
	return fn.FromRef(
		bindings.FormDataGetFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "FormData.getAll".
//
// The returned bool will be false if there is no such method.
func (this FormData) GetAll(name js.String) (js.Array[FormDataEntryValue], bool) {
	var _ok bool
	_ret := bindings.CallFormDataGetAll(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Array[FormDataEntryValue]{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "FormData.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) GetAllFunc() (fn js.Func[func(name js.String) js.Array[FormDataEntryValue]]) {
	return fn.FromRef(
		bindings.FormDataGetAllFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "FormData.has".
//
// The returned bool will be false if there is no such method.
func (this FormData) Has(name js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallFormDataHas(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return _ret == js.True, _ok
}

// HasFunc returns the method "FormData.has".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) HasFunc() (fn js.Func[func(name js.String) bool]) {
	return fn.FromRef(
		bindings.FormDataHasFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "FormData.set".
//
// The returned bool will be false if there is no such method.
func (this FormData) Set(name js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataSet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "FormData.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.FormDataSetFunc(
			this.Ref(),
		),
	)
}

// Set1 calls the method "FormData.set".
//
// The returned bool will be false if there is no such method.
func (this FormData) Set1(name js.String, blobValue Blob, filename js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataSet1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Set1Func returns the method "FormData.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) Set1Func() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	return fn.FromRef(
		bindings.FormDataSet1Func(
			this.Ref(),
		),
	)
}

// Set2 calls the method "FormData.set".
//
// The returned bool will be false if there is no such method.
func (this FormData) Set2(name js.String, blobValue Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFormDataSet2(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		blobValue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Set2Func returns the method "FormData.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FormData) Set2Func() (fn js.Func[func(name js.String, blobValue Blob)]) {
	return fn.FromRef(
		bindings.FormDataSet2Func(
			this.Ref(),
		),
	)
}

type OneOf_File_String_FormData struct {
	ref js.Ref
}

func (x OneOf_File_String_FormData) Ref() js.Ref {
	return x.ref
}

func (x OneOf_File_String_FormData) Free() {
	x.ref.Free()
}

func (x OneOf_File_String_FormData) FromRef(ref js.Ref) OneOf_File_String_FormData {
	return OneOf_File_String_FormData{
		ref: ref,
	}
}

func (x OneOf_File_String_FormData) File() File {
	return File{}.FromRef(x.ref)
}

func (x OneOf_File_String_FormData) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_File_String_FormData) FormData() FormData {
	return FormData{}.FromRef(x.ref)
}

type ValidityStateFlags struct {
	// ValueMissing is "ValidityStateFlags.valueMissing"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ValueMissing MUST be set to true to make this field effective.
	ValueMissing bool
	// TypeMismatch is "ValidityStateFlags.typeMismatch"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_TypeMismatch MUST be set to true to make this field effective.
	TypeMismatch bool
	// PatternMismatch is "ValidityStateFlags.patternMismatch"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PatternMismatch MUST be set to true to make this field effective.
	PatternMismatch bool
	// TooLong is "ValidityStateFlags.tooLong"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_TooLong MUST be set to true to make this field effective.
	TooLong bool
	// TooShort is "ValidityStateFlags.tooShort"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_TooShort MUST be set to true to make this field effective.
	TooShort bool
	// RangeUnderflow is "ValidityStateFlags.rangeUnderflow"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RangeUnderflow MUST be set to true to make this field effective.
	RangeUnderflow bool
	// RangeOverflow is "ValidityStateFlags.rangeOverflow"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RangeOverflow MUST be set to true to make this field effective.
	RangeOverflow bool
	// StepMismatch is "ValidityStateFlags.stepMismatch"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_StepMismatch MUST be set to true to make this field effective.
	StepMismatch bool
	// BadInput is "ValidityStateFlags.badInput"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_BadInput MUST be set to true to make this field effective.
	BadInput bool
	// CustomError is "ValidityStateFlags.customError"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CustomError MUST be set to true to make this field effective.
	CustomError bool

	FFI_USE_ValueMissing    bool // for ValueMissing.
	FFI_USE_TypeMismatch    bool // for TypeMismatch.
	FFI_USE_PatternMismatch bool // for PatternMismatch.
	FFI_USE_TooLong         bool // for TooLong.
	FFI_USE_TooShort        bool // for TooShort.
	FFI_USE_RangeUnderflow  bool // for RangeUnderflow.
	FFI_USE_RangeOverflow   bool // for RangeOverflow.
	FFI_USE_StepMismatch    bool // for StepMismatch.
	FFI_USE_BadInput        bool // for BadInput.
	FFI_USE_CustomError     bool // for CustomError.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ValidityStateFlags with all fields set.
func (p ValidityStateFlags) FromRef(ref js.Ref) ValidityStateFlags {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ValidityStateFlags in the application heap.
func (p ValidityStateFlags) New() js.Ref {
	return bindings.ValidityStateFlagsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ValidityStateFlags) UpdateFrom(ref js.Ref) {
	bindings.ValidityStateFlagsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ValidityStateFlags) Update(ref js.Ref) {
	bindings.ValidityStateFlagsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ValidityState struct {
	ref js.Ref
}

func (this ValidityState) Once() ValidityState {
	this.Ref().Once()
	return this
}

func (this ValidityState) Ref() js.Ref {
	return this.ref
}

func (this ValidityState) FromRef(ref js.Ref) ValidityState {
	this.ref = ref
	return this
}

func (this ValidityState) Free() {
	this.Ref().Free()
}

// ValueMissing returns the value of property "ValidityState.valueMissing".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) ValueMissing() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateValueMissing(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TypeMismatch returns the value of property "ValidityState.typeMismatch".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) TypeMismatch() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateTypeMismatch(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// PatternMismatch returns the value of property "ValidityState.patternMismatch".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) PatternMismatch() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStatePatternMismatch(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TooLong returns the value of property "ValidityState.tooLong".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) TooLong() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateTooLong(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TooShort returns the value of property "ValidityState.tooShort".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) TooShort() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateTooShort(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// RangeUnderflow returns the value of property "ValidityState.rangeUnderflow".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) RangeUnderflow() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateRangeUnderflow(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// RangeOverflow returns the value of property "ValidityState.rangeOverflow".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) RangeOverflow() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateRangeOverflow(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// StepMismatch returns the value of property "ValidityState.stepMismatch".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) StepMismatch() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateStepMismatch(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// BadInput returns the value of property "ValidityState.badInput".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) BadInput() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateBadInput(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CustomError returns the value of property "ValidityState.customError".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) CustomError() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateCustomError(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Valid returns the value of property "ValidityState.valid".
//
// The returned bool will be false if there is no such property.
func (this ValidityState) Valid() (bool, bool) {
	var _ok bool
	_ret := bindings.GetValidityStateValid(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type CustomStateSet struct {
	ref js.Ref
}

func (this CustomStateSet) Once() CustomStateSet {
	this.Ref().Once()
	return this
}

func (this CustomStateSet) Ref() js.Ref {
	return this.ref
}

func (this CustomStateSet) FromRef(ref js.Ref) CustomStateSet {
	this.ref = ref
	return this
}

func (this CustomStateSet) Free() {
	this.Ref().Free()
}

// Add calls the method "CustomStateSet.add".
//
// The returned bool will be false if there is no such method.
func (this CustomStateSet) Add(value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomStateSetAdd(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "CustomStateSet.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomStateSet) AddFunc() (fn js.Func[func(value js.String)]) {
	return fn.FromRef(
		bindings.CustomStateSetAddFunc(
			this.Ref(),
		),
	)
}

type ElementInternals struct {
	ref js.Ref
}

func (this ElementInternals) Once() ElementInternals {
	this.Ref().Once()
	return this
}

func (this ElementInternals) Ref() js.Ref {
	return this.ref
}

func (this ElementInternals) FromRef(ref js.Ref) ElementInternals {
	this.ref = ref
	return this
}

func (this ElementInternals) Free() {
	this.Ref().Free()
}

// ShadowRoot returns the value of property "ElementInternals.shadowRoot".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) ShadowRoot() (ShadowRoot, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsShadowRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return ShadowRoot{}.FromRef(_ret), _ok
}

// Form returns the value of property "ElementInternals.form".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// WillValidate returns the value of property "ElementInternals.willValidate".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "ElementInternals.validity".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "ElementInternals.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "ElementInternals.labels".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// States returns the value of property "ElementInternals.states".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) States() (CustomStateSet, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsStates(
		this.Ref(), js.Pointer(&_ok),
	)
	return CustomStateSet{}.FromRef(_ret), _ok
}

// Role returns the value of property "ElementInternals.role".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) Role() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsRole(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Role sets the value of property "ElementInternals.role" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetRole(val js.String) bool {
	return js.True == bindings.SetElementInternalsRole(
		this.Ref(),
		val.Ref(),
	)
}

// AriaActiveDescendantElement returns the value of property "ElementInternals.ariaActiveDescendantElement".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaActiveDescendantElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaActiveDescendantElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// AriaActiveDescendantElement sets the value of property "ElementInternals.ariaActiveDescendantElement" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaActiveDescendantElement(val Element) bool {
	return js.True == bindings.SetElementInternalsAriaActiveDescendantElement(
		this.Ref(),
		val.Ref(),
	)
}

// AriaAtomic returns the value of property "ElementInternals.ariaAtomic".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaAtomic() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaAtomic(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaAtomic sets the value of property "ElementInternals.ariaAtomic" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaAtomic(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaAtomic(
		this.Ref(),
		val.Ref(),
	)
}

// AriaAutoComplete returns the value of property "ElementInternals.ariaAutoComplete".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaAutoComplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaAutoComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaAutoComplete sets the value of property "ElementInternals.ariaAutoComplete" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaAutoComplete(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaAutoComplete(
		this.Ref(),
		val.Ref(),
	)
}

// AriaBusy returns the value of property "ElementInternals.ariaBusy".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaBusy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaBusy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaBusy sets the value of property "ElementInternals.ariaBusy" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaBusy(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaBusy(
		this.Ref(),
		val.Ref(),
	)
}

// AriaChecked returns the value of property "ElementInternals.ariaChecked".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaChecked() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaChecked(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaChecked sets the value of property "ElementInternals.ariaChecked" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaChecked(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaChecked(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColCount returns the value of property "ElementInternals.ariaColCount".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaColCount() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaColCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColCount sets the value of property "ElementInternals.ariaColCount" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColCount(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColCount(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColIndex returns the value of property "ElementInternals.ariaColIndex".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaColIndex() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaColIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColIndex sets the value of property "ElementInternals.ariaColIndex" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColIndex(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColIndex(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColIndexText returns the value of property "ElementInternals.ariaColIndexText".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaColIndexText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaColIndexText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColIndexText sets the value of property "ElementInternals.ariaColIndexText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColIndexText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColIndexText(
		this.Ref(),
		val.Ref(),
	)
}

// AriaColSpan returns the value of property "ElementInternals.ariaColSpan".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaColSpan() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaColSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaColSpan sets the value of property "ElementInternals.ariaColSpan" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColSpan(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColSpan(
		this.Ref(),
		val.Ref(),
	)
}

// AriaControlsElements returns the value of property "ElementInternals.ariaControlsElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaControlsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaControlsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaControlsElements sets the value of property "ElementInternals.ariaControlsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaControlsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaControlsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaCurrent returns the value of property "ElementInternals.ariaCurrent".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaCurrent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaCurrent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaCurrent sets the value of property "ElementInternals.ariaCurrent" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaCurrent(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaCurrent(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDescribedByElements returns the value of property "ElementInternals.ariaDescribedByElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaDescribedByElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaDescribedByElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaDescribedByElements sets the value of property "ElementInternals.ariaDescribedByElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDescribedByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaDescribedByElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDescription returns the value of property "ElementInternals.ariaDescription".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaDescription() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaDescription sets the value of property "ElementInternals.ariaDescription" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDescription(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaDescription(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDetailsElements returns the value of property "ElementInternals.ariaDetailsElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaDetailsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaDetailsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaDetailsElements sets the value of property "ElementInternals.ariaDetailsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDetailsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaDetailsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaDisabled returns the value of property "ElementInternals.ariaDisabled".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaDisabled() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaDisabled sets the value of property "ElementInternals.ariaDisabled" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDisabled(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaDisabled(
		this.Ref(),
		val.Ref(),
	)
}

// AriaErrorMessageElements returns the value of property "ElementInternals.ariaErrorMessageElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaErrorMessageElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaErrorMessageElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaErrorMessageElements sets the value of property "ElementInternals.ariaErrorMessageElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaErrorMessageElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaErrorMessageElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaExpanded returns the value of property "ElementInternals.ariaExpanded".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaExpanded() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaExpanded(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaExpanded sets the value of property "ElementInternals.ariaExpanded" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaExpanded(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaExpanded(
		this.Ref(),
		val.Ref(),
	)
}

// AriaFlowToElements returns the value of property "ElementInternals.ariaFlowToElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaFlowToElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaFlowToElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaFlowToElements sets the value of property "ElementInternals.ariaFlowToElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaFlowToElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaFlowToElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaHasPopup returns the value of property "ElementInternals.ariaHasPopup".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaHasPopup() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaHasPopup(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaHasPopup sets the value of property "ElementInternals.ariaHasPopup" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaHasPopup(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaHasPopup(
		this.Ref(),
		val.Ref(),
	)
}

// AriaHidden returns the value of property "ElementInternals.ariaHidden".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaHidden() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaHidden(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaHidden sets the value of property "ElementInternals.ariaHidden" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaHidden(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaHidden(
		this.Ref(),
		val.Ref(),
	)
}

// AriaInvalid returns the value of property "ElementInternals.ariaInvalid".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaInvalid() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaInvalid(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaInvalid sets the value of property "ElementInternals.ariaInvalid" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaInvalid(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaInvalid(
		this.Ref(),
		val.Ref(),
	)
}

// AriaKeyShortcuts returns the value of property "ElementInternals.ariaKeyShortcuts".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaKeyShortcuts() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaKeyShortcuts(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaKeyShortcuts sets the value of property "ElementInternals.ariaKeyShortcuts" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaKeyShortcuts(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaKeyShortcuts(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLabel returns the value of property "ElementInternals.ariaLabel".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaLabel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLabel sets the value of property "ElementInternals.ariaLabel" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLabel(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLabel(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLabelledByElements returns the value of property "ElementInternals.ariaLabelledByElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaLabelledByElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaLabelledByElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaLabelledByElements sets the value of property "ElementInternals.ariaLabelledByElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLabelledByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaLabelledByElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLevel returns the value of property "ElementInternals.ariaLevel".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaLevel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaLevel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLevel sets the value of property "ElementInternals.ariaLevel" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLevel(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLevel(
		this.Ref(),
		val.Ref(),
	)
}

// AriaLive returns the value of property "ElementInternals.ariaLive".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaLive() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaLive(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaLive sets the value of property "ElementInternals.ariaLive" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLive(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLive(
		this.Ref(),
		val.Ref(),
	)
}

// AriaModal returns the value of property "ElementInternals.ariaModal".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaModal() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaModal(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaModal sets the value of property "ElementInternals.ariaModal" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaModal(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaModal(
		this.Ref(),
		val.Ref(),
	)
}

// AriaMultiLine returns the value of property "ElementInternals.ariaMultiLine".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaMultiLine() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaMultiLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaMultiLine sets the value of property "ElementInternals.ariaMultiLine" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaMultiLine(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaMultiLine(
		this.Ref(),
		val.Ref(),
	)
}

// AriaMultiSelectable returns the value of property "ElementInternals.ariaMultiSelectable".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaMultiSelectable() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaMultiSelectable(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaMultiSelectable sets the value of property "ElementInternals.ariaMultiSelectable" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaMultiSelectable(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaMultiSelectable(
		this.Ref(),
		val.Ref(),
	)
}

// AriaOrientation returns the value of property "ElementInternals.ariaOrientation".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaOrientation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaOrientation sets the value of property "ElementInternals.ariaOrientation" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaOrientation(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaOrientation(
		this.Ref(),
		val.Ref(),
	)
}

// AriaOwnsElements returns the value of property "ElementInternals.ariaOwnsElements".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaOwnsElements() (js.FrozenArray[Element], bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaOwnsElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[Element]{}.FromRef(_ret), _ok
}

// AriaOwnsElements sets the value of property "ElementInternals.ariaOwnsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaOwnsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaOwnsElements(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPlaceholder returns the value of property "ElementInternals.ariaPlaceholder".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaPlaceholder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaPlaceholder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPlaceholder sets the value of property "ElementInternals.ariaPlaceholder" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPlaceholder(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPlaceholder(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPosInSet returns the value of property "ElementInternals.ariaPosInSet".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaPosInSet() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaPosInSet(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPosInSet sets the value of property "ElementInternals.ariaPosInSet" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPosInSet(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPosInSet(
		this.Ref(),
		val.Ref(),
	)
}

// AriaPressed returns the value of property "ElementInternals.ariaPressed".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaPressed() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaPressed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaPressed sets the value of property "ElementInternals.ariaPressed" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPressed(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPressed(
		this.Ref(),
		val.Ref(),
	)
}

// AriaReadOnly returns the value of property "ElementInternals.ariaReadOnly".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaReadOnly() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaReadOnly(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaReadOnly sets the value of property "ElementInternals.ariaReadOnly" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaReadOnly(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaReadOnly(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRequired returns the value of property "ElementInternals.ariaRequired".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRequired() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRequired(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRequired sets the value of property "ElementInternals.ariaRequired" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRequired(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRequired(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRoleDescription returns the value of property "ElementInternals.ariaRoleDescription".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRoleDescription() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRoleDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRoleDescription sets the value of property "ElementInternals.ariaRoleDescription" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRoleDescription(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRoleDescription(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowCount returns the value of property "ElementInternals.ariaRowCount".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRowCount() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRowCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowCount sets the value of property "ElementInternals.ariaRowCount" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowCount(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowCount(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowIndex returns the value of property "ElementInternals.ariaRowIndex".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRowIndex() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRowIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowIndex sets the value of property "ElementInternals.ariaRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowIndex(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowIndex(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowIndexText returns the value of property "ElementInternals.ariaRowIndexText".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRowIndexText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRowIndexText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowIndexText sets the value of property "ElementInternals.ariaRowIndexText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowIndexText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowIndexText(
		this.Ref(),
		val.Ref(),
	)
}

// AriaRowSpan returns the value of property "ElementInternals.ariaRowSpan".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaRowSpan() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaRowSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaRowSpan sets the value of property "ElementInternals.ariaRowSpan" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowSpan(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowSpan(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSelected returns the value of property "ElementInternals.ariaSelected".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaSelected() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSelected sets the value of property "ElementInternals.ariaSelected" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSelected(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSelected(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSetSize returns the value of property "ElementInternals.ariaSetSize".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaSetSize() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaSetSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSetSize sets the value of property "ElementInternals.ariaSetSize" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSetSize(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSetSize(
		this.Ref(),
		val.Ref(),
	)
}

// AriaSort returns the value of property "ElementInternals.ariaSort".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaSort() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaSort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaSort sets the value of property "ElementInternals.ariaSort" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSort(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSort(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueMax returns the value of property "ElementInternals.ariaValueMax".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaValueMax() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaValueMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueMax sets the value of property "ElementInternals.ariaValueMax" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueMax(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueMax(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueMin returns the value of property "ElementInternals.ariaValueMin".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaValueMin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaValueMin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueMin sets the value of property "ElementInternals.ariaValueMin" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueMin(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueMin(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueNow returns the value of property "ElementInternals.ariaValueNow".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaValueNow() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaValueNow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueNow sets the value of property "ElementInternals.ariaValueNow" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueNow(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueNow(
		this.Ref(),
		val.Ref(),
	)
}

// AriaValueText returns the value of property "ElementInternals.ariaValueText".
//
// The returned bool will be false if there is no such property.
func (this ElementInternals) AriaValueText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetElementInternalsAriaValueText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AriaValueText sets the value of property "ElementInternals.ariaValueText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueText(
		this.Ref(),
		val.Ref(),
	)
}

// SetFormValue calls the method "ElementInternals.setFormValue".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetFormValue(value OneOf_File_String_FormData, state OneOf_File_String_FormData) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetFormValue(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
		state.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFormValueFunc returns the method "ElementInternals.setFormValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetFormValueFunc() (fn js.Func[func(value OneOf_File_String_FormData, state OneOf_File_String_FormData)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetFormValueFunc(
			this.Ref(),
		),
	)
}

// SetFormValue1 calls the method "ElementInternals.setFormValue".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetFormValue1(value OneOf_File_String_FormData) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetFormValue1(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFormValue1Func returns the method "ElementInternals.setFormValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetFormValue1Func() (fn js.Func[func(value OneOf_File_String_FormData)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetFormValue1Func(
			this.Ref(),
		),
	)
}

// SetValidity calls the method "ElementInternals.setValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetValidity(flags ValidityStateFlags, message js.String, anchor HTMLElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetValidity(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&flags),
		message.Ref(),
		anchor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetValidityFunc returns the method "ElementInternals.setValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetValidityFunc() (fn js.Func[func(flags ValidityStateFlags, message js.String, anchor HTMLElement)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidityFunc(
			this.Ref(),
		),
	)
}

// SetValidity1 calls the method "ElementInternals.setValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetValidity1(flags ValidityStateFlags, message js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetValidity1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&flags),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetValidity1Func returns the method "ElementInternals.setValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetValidity1Func() (fn js.Func[func(flags ValidityStateFlags, message js.String)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity1Func(
			this.Ref(),
		),
	)
}

// SetValidity2 calls the method "ElementInternals.setValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetValidity2(flags ValidityStateFlags) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetValidity2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&flags),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetValidity2Func returns the method "ElementInternals.setValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetValidity2Func() (fn js.Func[func(flags ValidityStateFlags)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity2Func(
			this.Ref(),
		),
	)
}

// SetValidity3 calls the method "ElementInternals.setValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) SetValidity3() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsSetValidity3(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetValidity3Func returns the method "ElementInternals.setValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) SetValidity3Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity3Func(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "ElementInternals.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "ElementInternals.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementInternalsCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "ElementInternals.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this ElementInternals) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallElementInternalsReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "ElementInternals.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ElementInternals) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementInternalsReportValidityFunc(
			this.Ref(),
		),
	)
}

type OneOf_Bool_Float64_String struct {
	ref js.Ref
}

func (x OneOf_Bool_Float64_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_Float64_String) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_Float64_String) FromRef(ref js.Ref) OneOf_Bool_Float64_String {
	return OneOf_Bool_Float64_String{
		ref: ref,
	}
}

func (x OneOf_Bool_Float64_String) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_Float64_String) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Bool_Float64_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type EditContextInit struct {
	// Text is "EditContextInit.text"
	//
	// Optional
	Text js.String
	// SelectionStart is "EditContextInit.selectionStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionStart MUST be set to true to make this field effective.
	SelectionStart uint32
	// SelectionEnd is "EditContextInit.selectionEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionEnd MUST be set to true to make this field effective.
	SelectionEnd uint32

	FFI_USE_SelectionStart bool // for SelectionStart.
	FFI_USE_SelectionEnd   bool // for SelectionEnd.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EditContextInit with all fields set.
func (p EditContextInit) FromRef(ref js.Ref) EditContextInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EditContextInit in the application heap.
func (p EditContextInit) New() js.Ref {
	return bindings.EditContextInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EditContextInit) UpdateFrom(ref js.Ref) {
	bindings.EditContextInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EditContextInit) Update(ref js.Ref) {
	bindings.EditContextInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewEditContext(options EditContextInit) EditContext {
	return EditContext{}.FromRef(
		bindings.NewEditContextByEditContext(
			js.Pointer(&options)),
	)
}

func NewEditContextByEditContext1() EditContext {
	return EditContext{}.FromRef(
		bindings.NewEditContextByEditContext1(),
	)
}

type EditContext struct {
	EventTarget
}

func (this EditContext) Once() EditContext {
	this.Ref().Once()
	return this
}

func (this EditContext) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this EditContext) FromRef(ref js.Ref) EditContext {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this EditContext) Free() {
	this.Ref().Free()
}

// Text returns the value of property "EditContext.text".
//
// The returned bool will be false if there is no such property.
func (this EditContext) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetEditContextText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectionStart returns the value of property "EditContext.selectionStart".
//
// The returned bool will be false if there is no such property.
func (this EditContext) SelectionStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEditContextSelectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionEnd returns the value of property "EditContext.selectionEnd".
//
// The returned bool will be false if there is no such property.
func (this EditContext) SelectionEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEditContextSelectionEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CompositionRangeStart returns the value of property "EditContext.compositionRangeStart".
//
// The returned bool will be false if there is no such property.
func (this EditContext) CompositionRangeStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEditContextCompositionRangeStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CompositionRangeEnd returns the value of property "EditContext.compositionRangeEnd".
//
// The returned bool will be false if there is no such property.
func (this EditContext) CompositionRangeEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEditContextCompositionRangeEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IsComposing returns the value of property "EditContext.isComposing".
//
// The returned bool will be false if there is no such property.
func (this EditContext) IsComposing() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEditContextIsComposing(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ControlBounds returns the value of property "EditContext.controlBounds".
//
// The returned bool will be false if there is no such property.
func (this EditContext) ControlBounds() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.GetEditContextControlBounds(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRect{}.FromRef(_ret), _ok
}

// SelectionBounds returns the value of property "EditContext.selectionBounds".
//
// The returned bool will be false if there is no such property.
func (this EditContext) SelectionBounds() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.GetEditContextSelectionBounds(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRect{}.FromRef(_ret), _ok
}

// CharacterBoundsRangeStart returns the value of property "EditContext.characterBoundsRangeStart".
//
// The returned bool will be false if there is no such property.
func (this EditContext) CharacterBoundsRangeStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEditContextCharacterBoundsRangeStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// UpdateText calls the method "EditContext.updateText".
//
// The returned bool will be false if there is no such method.
func (this EditContext) UpdateText(rangeStart uint32, rangeEnd uint32, text js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEditContextUpdateText(
		this.Ref(), js.Pointer(&_ok),
		uint32(rangeStart),
		uint32(rangeEnd),
		text.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateTextFunc returns the method "EditContext.updateText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) UpdateTextFunc() (fn js.Func[func(rangeStart uint32, rangeEnd uint32, text js.String)]) {
	return fn.FromRef(
		bindings.EditContextUpdateTextFunc(
			this.Ref(),
		),
	)
}

// UpdateSelection calls the method "EditContext.updateSelection".
//
// The returned bool will be false if there is no such method.
func (this EditContext) UpdateSelection(start uint32, end uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEditContextUpdateSelection(
		this.Ref(), js.Pointer(&_ok),
		uint32(start),
		uint32(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateSelectionFunc returns the method "EditContext.updateSelection".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) UpdateSelectionFunc() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.EditContextUpdateSelectionFunc(
			this.Ref(),
		),
	)
}

// UpdateControlBounds calls the method "EditContext.updateControlBounds".
//
// The returned bool will be false if there is no such method.
func (this EditContext) UpdateControlBounds(controlBounds DOMRect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEditContextUpdateControlBounds(
		this.Ref(), js.Pointer(&_ok),
		controlBounds.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateControlBoundsFunc returns the method "EditContext.updateControlBounds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) UpdateControlBoundsFunc() (fn js.Func[func(controlBounds DOMRect)]) {
	return fn.FromRef(
		bindings.EditContextUpdateControlBoundsFunc(
			this.Ref(),
		),
	)
}

// UpdateSelectionBounds calls the method "EditContext.updateSelectionBounds".
//
// The returned bool will be false if there is no such method.
func (this EditContext) UpdateSelectionBounds(selectionBounds DOMRect) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEditContextUpdateSelectionBounds(
		this.Ref(), js.Pointer(&_ok),
		selectionBounds.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateSelectionBoundsFunc returns the method "EditContext.updateSelectionBounds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) UpdateSelectionBoundsFunc() (fn js.Func[func(selectionBounds DOMRect)]) {
	return fn.FromRef(
		bindings.EditContextUpdateSelectionBoundsFunc(
			this.Ref(),
		),
	)
}

// UpdateCharacterBounds calls the method "EditContext.updateCharacterBounds".
//
// The returned bool will be false if there is no such method.
func (this EditContext) UpdateCharacterBounds(rangeStart uint32, characterBounds js.Array[DOMRect]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEditContextUpdateCharacterBounds(
		this.Ref(), js.Pointer(&_ok),
		uint32(rangeStart),
		characterBounds.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateCharacterBoundsFunc returns the method "EditContext.updateCharacterBounds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) UpdateCharacterBoundsFunc() (fn js.Func[func(rangeStart uint32, characterBounds js.Array[DOMRect])]) {
	return fn.FromRef(
		bindings.EditContextUpdateCharacterBoundsFunc(
			this.Ref(),
		),
	)
}

// AttachedElements calls the method "EditContext.attachedElements".
//
// The returned bool will be false if there is no such method.
func (this EditContext) AttachedElements() (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallEditContextAttachedElements(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// AttachedElementsFunc returns the method "EditContext.attachedElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) AttachedElementsFunc() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.EditContextAttachedElementsFunc(
			this.Ref(),
		),
	)
}

// CharacterBounds calls the method "EditContext.characterBounds".
//
// The returned bool will be false if there is no such method.
func (this EditContext) CharacterBounds() (js.Array[DOMRect], bool) {
	var _ok bool
	_ret := bindings.CallEditContextCharacterBounds(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[DOMRect]{}.FromRef(_ret), _ok
}

// CharacterBoundsFunc returns the method "EditContext.characterBounds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EditContext) CharacterBoundsFunc() (fn js.Func[func() js.Array[DOMRect]]) {
	return fn.FromRef(
		bindings.EditContextCharacterBoundsFunc(
			this.Ref(),
		),
	)
}

func NewHTMLElement() HTMLElement {
	return HTMLElement{}.FromRef(
		bindings.NewHTMLElementByHTMLElement(),
	)
}

type HTMLElement struct {
	Element
}

func (this HTMLElement) Once() HTMLElement {
	this.Ref().Once()
	return this
}

func (this HTMLElement) Ref() js.Ref {
	return this.Element.Ref()
}

func (this HTMLElement) FromRef(ref js.Ref) HTMLElement {
	this.Element = this.Element.FromRef(ref)
	return this
}

func (this HTMLElement) Free() {
	this.Ref().Free()
}

// Title returns the value of property "HTMLElement.title".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Title sets the value of property "HTMLElement.title" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTitle(val js.String) bool {
	return js.True == bindings.SetHTMLElementTitle(
		this.Ref(),
		val.Ref(),
	)
}

// Lang returns the value of property "HTMLElement.lang".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lang sets the value of property "HTMLElement.lang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetLang(val js.String) bool {
	return js.True == bindings.SetHTMLElementLang(
		this.Ref(),
		val.Ref(),
	)
}

// Translate returns the value of property "HTMLElement.translate".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Translate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementTranslate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Translate sets the value of property "HTMLElement.translate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTranslate(val bool) bool {
	return js.True == bindings.SetHTMLElementTranslate(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Dir returns the value of property "HTMLElement.dir".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Dir() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementDir(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Dir sets the value of property "HTMLElement.dir" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetDir(val js.String) bool {
	return js.True == bindings.SetHTMLElementDir(
		this.Ref(),
		val.Ref(),
	)
}

// Hidden returns the value of property "HTMLElement.hidden".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Hidden() (OneOf_Bool_Float64_String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementHidden(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Bool_Float64_String{}.FromRef(_ret), _ok
}

// Hidden sets the value of property "HTMLElement.hidden" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetHidden(val OneOf_Bool_Float64_String) bool {
	return js.True == bindings.SetHTMLElementHidden(
		this.Ref(),
		val.Ref(),
	)
}

// Inert returns the value of property "HTMLElement.inert".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Inert() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementInert(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Inert sets the value of property "HTMLElement.inert" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInert(val bool) bool {
	return js.True == bindings.SetHTMLElementInert(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// AccessKey returns the value of property "HTMLElement.accessKey".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) AccessKey() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementAccessKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AccessKey sets the value of property "HTMLElement.accessKey" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAccessKey(val js.String) bool {
	return js.True == bindings.SetHTMLElementAccessKey(
		this.Ref(),
		val.Ref(),
	)
}

// AccessKeyLabel returns the value of property "HTMLElement.accessKeyLabel".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) AccessKeyLabel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementAccessKeyLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Draggable returns the value of property "HTMLElement.draggable".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Draggable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementDraggable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Draggable sets the value of property "HTMLElement.draggable" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetDraggable(val bool) bool {
	return js.True == bindings.SetHTMLElementDraggable(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Spellcheck returns the value of property "HTMLElement.spellcheck".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Spellcheck() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementSpellcheck(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Spellcheck sets the value of property "HTMLElement.spellcheck" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetSpellcheck(val bool) bool {
	return js.True == bindings.SetHTMLElementSpellcheck(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Autocapitalize returns the value of property "HTMLElement.autocapitalize".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Autocapitalize() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementAutocapitalize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Autocapitalize sets the value of property "HTMLElement.autocapitalize" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAutocapitalize(val js.String) bool {
	return js.True == bindings.SetHTMLElementAutocapitalize(
		this.Ref(),
		val.Ref(),
	)
}

// InnerText returns the value of property "HTMLElement.innerText".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) InnerText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementInnerText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InnerText sets the value of property "HTMLElement.innerText" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInnerText(val js.String) bool {
	return js.True == bindings.SetHTMLElementInnerText(
		this.Ref(),
		val.Ref(),
	)
}

// OuterText returns the value of property "HTMLElement.outerText".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OuterText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOuterText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OuterText sets the value of property "HTMLElement.outerText" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetOuterText(val js.String) bool {
	return js.True == bindings.SetHTMLElementOuterText(
		this.Ref(),
		val.Ref(),
	)
}

// Popover returns the value of property "HTMLElement.popover".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Popover() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementPopover(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Popover sets the value of property "HTMLElement.popover" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetPopover(val js.String) bool {
	return js.True == bindings.SetHTMLElementPopover(
		this.Ref(),
		val.Ref(),
	)
}

// OffsetParent returns the value of property "HTMLElement.offsetParent".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OffsetParent() (Element, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOffsetParent(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// OffsetTop returns the value of property "HTMLElement.offsetTop".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OffsetTop() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOffsetTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// OffsetLeft returns the value of property "HTMLElement.offsetLeft".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OffsetLeft() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOffsetLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// OffsetWidth returns the value of property "HTMLElement.offsetWidth".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OffsetWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOffsetWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// OffsetHeight returns the value of property "HTMLElement.offsetHeight".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) OffsetHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementOffsetHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// EditContext returns the value of property "HTMLElement.editContext".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) EditContext() (EditContext, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementEditContext(
		this.Ref(), js.Pointer(&_ok),
	)
	return EditContext{}.FromRef(_ret), _ok
}

// EditContext sets the value of property "HTMLElement.editContext" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetEditContext(val EditContext) bool {
	return js.True == bindings.SetHTMLElementEditContext(
		this.Ref(),
		val.Ref(),
	)
}

// Style returns the value of property "HTMLElement.style".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// AttributeStyleMap returns the value of property "HTMLElement.attributeStyleMap".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) AttributeStyleMap() (StylePropertyMap, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementAttributeStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return StylePropertyMap{}.FromRef(_ret), _ok
}

// Dataset returns the value of property "HTMLElement.dataset".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Dataset() (DOMStringMap, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementDataset(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringMap{}.FromRef(_ret), _ok
}

// Nonce returns the value of property "HTMLElement.nonce".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Nonce() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementNonce(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Nonce sets the value of property "HTMLElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetHTMLElementNonce(
		this.Ref(),
		val.Ref(),
	)
}

// Autofocus returns the value of property "HTMLElement.autofocus".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) Autofocus() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementAutofocus(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Autofocus sets the value of property "HTMLElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetHTMLElementAutofocus(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "HTMLElement.tabIndex".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) TabIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementTabIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// TabIndex sets the value of property "HTMLElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetHTMLElementTabIndex(
		this.Ref(),
		int32(val),
	)
}

// ContentEditable returns the value of property "HTMLElement.contentEditable".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) ContentEditable() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementContentEditable(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContentEditable sets the value of property "HTMLElement.contentEditable" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetContentEditable(val js.String) bool {
	return js.True == bindings.SetHTMLElementContentEditable(
		this.Ref(),
		val.Ref(),
	)
}

// EnterKeyHint returns the value of property "HTMLElement.enterKeyHint".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) EnterKeyHint() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementEnterKeyHint(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EnterKeyHint sets the value of property "HTMLElement.enterKeyHint" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetEnterKeyHint(val js.String) bool {
	return js.True == bindings.SetHTMLElementEnterKeyHint(
		this.Ref(),
		val.Ref(),
	)
}

// IsContentEditable returns the value of property "HTMLElement.isContentEditable".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) IsContentEditable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementIsContentEditable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// InputMode returns the value of property "HTMLElement.inputMode".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) InputMode() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementInputMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InputMode sets the value of property "HTMLElement.inputMode" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInputMode(val js.String) bool {
	return js.True == bindings.SetHTMLElementInputMode(
		this.Ref(),
		val.Ref(),
	)
}

// VirtualKeyboardPolicy returns the value of property "HTMLElement.virtualKeyboardPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLElement) VirtualKeyboardPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLElementVirtualKeyboardPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VirtualKeyboardPolicy sets the value of property "HTMLElement.virtualKeyboardPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetVirtualKeyboardPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLElementVirtualKeyboardPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Click calls the method "HTMLElement.click".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) Click() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementClick(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClickFunc returns the method "HTMLElement.click".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) ClickFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementClickFunc(
			this.Ref(),
		),
	)
}

// AttachInternals calls the method "HTMLElement.attachInternals".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) AttachInternals() (ElementInternals, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementAttachInternals(
		this.Ref(), js.Pointer(&_ok),
	)

	return ElementInternals{}.FromRef(_ret), _ok
}

// AttachInternalsFunc returns the method "HTMLElement.attachInternals".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) AttachInternalsFunc() (fn js.Func[func() ElementInternals]) {
	return fn.FromRef(
		bindings.HTMLElementAttachInternalsFunc(
			this.Ref(),
		),
	)
}

// ShowPopover calls the method "HTMLElement.showPopover".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) ShowPopover() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementShowPopover(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShowPopoverFunc returns the method "HTMLElement.showPopover".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) ShowPopoverFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementShowPopoverFunc(
			this.Ref(),
		),
	)
}

// HidePopover calls the method "HTMLElement.hidePopover".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) HidePopover() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementHidePopover(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// HidePopoverFunc returns the method "HTMLElement.hidePopover".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) HidePopoverFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementHidePopoverFunc(
			this.Ref(),
		),
	)
}

// TogglePopover calls the method "HTMLElement.togglePopover".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) TogglePopover(force bool) (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementTogglePopover(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(force)),
	)

	return _ret == js.True, _ok
}

// TogglePopoverFunc returns the method "HTMLElement.togglePopover".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) TogglePopoverFunc() (fn js.Func[func(force bool) bool]) {
	return fn.FromRef(
		bindings.HTMLElementTogglePopoverFunc(
			this.Ref(),
		),
	)
}

// TogglePopover1 calls the method "HTMLElement.togglePopover".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) TogglePopover1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementTogglePopover1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// TogglePopover1Func returns the method "HTMLElement.togglePopover".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) TogglePopover1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLElementTogglePopover1Func(
			this.Ref(),
		),
	)
}

// Focus calls the method "HTMLElement.focus".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) Focus(options FocusOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementFocus(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FocusFunc returns the method "HTMLElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.HTMLElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "HTMLElement.focus".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) Focus1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementFocus1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Focus1Func returns the method "HTMLElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementFocus1Func(
			this.Ref(),
		),
	)
}

// Blur calls the method "HTMLElement.blur".
//
// The returned bool will be false if there is no such method.
func (this HTMLElement) Blur() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLElementBlur(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlurFunc returns the method "HTMLElement.blur".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementBlurFunc(
			this.Ref(),
		),
	)
}

func NewHTMLHeadElement() HTMLHeadElement {
	return HTMLHeadElement{}.FromRef(
		bindings.NewHTMLHeadElementByHTMLHeadElement(),
	)
}

type HTMLHeadElement struct {
	HTMLElement
}

func (this HTMLHeadElement) Once() HTMLHeadElement {
	this.Ref().Once()
	return this
}

func (this HTMLHeadElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLHeadElement) FromRef(ref js.Ref) HTMLHeadElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLHeadElement) Free() {
	this.Ref().Free()
}

func NewHTMLScriptElement() HTMLScriptElement {
	return HTMLScriptElement{}.FromRef(
		bindings.NewHTMLScriptElementByHTMLScriptElement(),
	)
}

type HTMLScriptElement struct {
	HTMLElement
}

func (this HTMLScriptElement) Once() HTMLScriptElement {
	this.Ref().Once()
	return this
}

func (this HTMLScriptElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLScriptElement) FromRef(ref js.Ref) HTMLScriptElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLScriptElement) Free() {
	this.Ref().Free()
}

// Src returns the value of property "HTMLScriptElement.src".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLScriptElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLScriptElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLScriptElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementType(
		this.Ref(),
		val.Ref(),
	)
}

// NoModule returns the value of property "HTMLScriptElement.noModule".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) NoModule() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementNoModule(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoModule sets the value of property "HTMLScriptElement.noModule" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetNoModule(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementNoModule(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Async returns the value of property "HTMLScriptElement.async".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Async() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementAsync(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Async sets the value of property "HTMLScriptElement.async" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetAsync(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementAsync(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Defer returns the value of property "HTMLScriptElement.defer".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Defer() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementDefer(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Defer sets the value of property "HTMLScriptElement.defer" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetDefer(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementDefer(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// CrossOrigin returns the value of property "HTMLScriptElement.crossOrigin".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "HTMLScriptElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// Text returns the value of property "HTMLScriptElement.text".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "HTMLScriptElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementText(
		this.Ref(),
		val.Ref(),
	)
}

// Integrity returns the value of property "HTMLScriptElement.integrity".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Integrity() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementIntegrity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Integrity sets the value of property "HTMLScriptElement.integrity" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetIntegrity(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementIntegrity(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLScriptElement.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLScriptElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLScriptElement.blocking".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Blocking() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementBlocking(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// FetchPriority returns the value of property "HTMLScriptElement.fetchPriority".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) FetchPriority() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementFetchPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FetchPriority sets the value of property "HTMLScriptElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementFetchPriority(
		this.Ref(),
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLScriptElement.charset".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Charset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementCharset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Charset sets the value of property "HTMLScriptElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementCharset(
		this.Ref(),
		val.Ref(),
	)
}

// Event returns the value of property "HTMLScriptElement.event".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) Event() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementEvent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Event sets the value of property "HTMLScriptElement.event" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetEvent(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementEvent(
		this.Ref(),
		val.Ref(),
	)
}

// HtmlFor returns the value of property "HTMLScriptElement.htmlFor".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) HtmlFor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementHtmlFor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// HtmlFor sets the value of property "HTMLScriptElement.htmlFor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetHtmlFor(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementHtmlFor(
		this.Ref(),
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLScriptElement.attributionSrc".
//
// The returned bool will be false if there is no such property.
func (this HTMLScriptElement) AttributionSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLScriptElementAttributionSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttributionSrc sets the value of property "HTMLScriptElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementAttributionSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Supports calls the staticmethod "HTMLScriptElement.supports".
//
// The returned bool will be false if there is no such method.
func (this HTMLScriptElement) Supports(typ js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLScriptElementSupports(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return _ret == js.True, _ok
}

// SupportsFunc returns the staticmethod "HTMLScriptElement.supports".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLScriptElement) SupportsFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.HTMLScriptElementSupportsFunc(
			this.Ref(),
		),
	)
}

type SVGScriptElement struct {
	SVGElement
}

func (this SVGScriptElement) Once() SVGScriptElement {
	this.Ref().Once()
	return this
}

func (this SVGScriptElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGScriptElement) FromRef(ref js.Ref) SVGScriptElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGScriptElement) Free() {
	this.Ref().Free()
}

// Type returns the value of property "SVGScriptElement.type".
//
// The returned bool will be false if there is no such property.
func (this SVGScriptElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGScriptElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "SVGScriptElement.type" to val.
//
// It returns false if the property cannot be set.
func (this SVGScriptElement) SetType(val js.String) bool {
	return js.True == bindings.SetSVGScriptElementType(
		this.Ref(),
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "SVGScriptElement.crossOrigin".
//
// The returned bool will be false if there is no such property.
func (this SVGScriptElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGScriptElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "SVGScriptElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this SVGScriptElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetSVGScriptElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// Href returns the value of property "SVGScriptElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGScriptElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGScriptElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

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
// It returns ok=false if there is no such property.
func (this HTMLAllCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLAllCollectionLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "HTMLAllCollection." exists.
func (this HTMLAllCollection) HasGet() bool {
	return js.True == bindings.HasHTMLAllCollectionGet(
		this.Ref(),
	)
}

// GetFunc returns the method "HTMLAllCollection.".
func (this HTMLAllCollection) GetFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "HTMLAllCollection.".
func (this HTMLAllCollection) Get(index uint32) (ret Element) {
	bindings.CallHTMLAllCollectionGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "HTMLAllCollection."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryGet(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "HTMLAllCollection.namedItem" exists.
func (this HTMLAllCollection) HasNamedItem() bool {
	return js.True == bindings.HasHTMLAllCollectionNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "HTMLAllCollection.namedItem".
func (this HTMLAllCollection) NamedItemFunc() (fn js.Func[func(name js.String) OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLAllCollection.namedItem".
func (this HTMLAllCollection) NamedItem(name js.String) (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLAllCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryNamedItem(name js.String) (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasItem returns true if the method "HTMLAllCollection.item" exists.
func (this HTMLAllCollection) HasItem() bool {
	return js.True == bindings.HasHTMLAllCollectionItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "HTMLAllCollection.item".
func (this HTMLAllCollection) ItemFunc() (fn js.Func[func(nameOrIndex js.String) OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "HTMLAllCollection.item".
func (this HTMLAllCollection) Item(nameOrIndex js.String) (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionItem(
		this.Ref(), js.Pointer(&ret),
		nameOrIndex.Ref(),
	)

	return
}

// TryItem calls the method "HTMLAllCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryItem(nameOrIndex js.String) (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		nameOrIndex.Ref(),
	)

	return
}

// HasItem1 returns true if the method "HTMLAllCollection.item" exists.
func (this HTMLAllCollection) HasItem1() bool {
	return js.True == bindings.HasHTMLAllCollectionItem1(
		this.Ref(),
	)
}

// Item1Func returns the method "HTMLAllCollection.item".
func (this HTMLAllCollection) Item1Func() (fn js.Func[func() OneOf_HTMLCollection_Element]) {
	return fn.FromRef(
		bindings.HTMLAllCollectionItem1Func(
			this.Ref(),
		),
	)
}

// Item1 calls the method "HTMLAllCollection.item".
func (this HTMLAllCollection) Item1() (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionItem1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryItem1 calls the method "HTMLAllCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryItem1() (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionItem1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DOMStringList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMStringListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "DOMStringList.item" exists.
func (this DOMStringList) HasItem() bool {
	return js.True == bindings.HasDOMStringListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "DOMStringList.item".
func (this DOMStringList) ItemFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.DOMStringListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "DOMStringList.item".
func (this DOMStringList) Item(index uint32) (ret js.String) {
	bindings.CallDOMStringListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMStringList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasContains returns true if the method "DOMStringList.contains" exists.
func (this DOMStringList) HasContains() bool {
	return js.True == bindings.HasDOMStringListContains(
		this.Ref(),
	)
}

// ContainsFunc returns the method "DOMStringList.contains".
func (this DOMStringList) ContainsFunc() (fn js.Func[func(string js.String) bool]) {
	return fn.FromRef(
		bindings.DOMStringListContainsFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "DOMStringList.contains".
func (this DOMStringList) Contains(string js.String) (ret bool) {
	bindings.CallDOMStringListContains(
		this.Ref(), js.Pointer(&ret),
		string.Ref(),
	)

	return
}

// TryContains calls the method "DOMStringList.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringList) TryContains(string js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringListContains(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		string.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Location) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "Location.href" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "Location.protocol".
//
// It returns ok=false if there is no such property.
func (this Location) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "Location.protocol" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "Location.host" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHostname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "Location.hostname" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "Location.port" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationPathname(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "Location.pathname" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationSearch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "Location.search" to val.
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
// It returns ok=false if there is no such property.
func (this Location) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHash(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "Location.hash" to val.
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
// It returns ok=false if there is no such property.
func (this Location) AncestorOrigins() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetLocationAncestorOrigins(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAssign returns true if the method "Location.assign" exists.
func (this Location) HasAssign() bool {
	return js.True == bindings.HasLocationAssign(
		this.Ref(),
	)
}

// AssignFunc returns the method "Location.assign".
func (this Location) AssignFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.LocationAssignFunc(
			this.Ref(),
		),
	)
}

// Assign calls the method "Location.assign".
func (this Location) Assign(url js.String) (ret js.Void) {
	bindings.CallLocationAssign(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryAssign calls the method "Location.assign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryAssign(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationAssign(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasReplace returns true if the method "Location.replace" exists.
func (this Location) HasReplace() bool {
	return js.True == bindings.HasLocationReplace(
		this.Ref(),
	)
}

// ReplaceFunc returns the method "Location.replace".
func (this Location) ReplaceFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.LocationReplaceFunc(
			this.Ref(),
		),
	)
}

// Replace calls the method "Location.replace".
func (this Location) Replace(url js.String) (ret js.Void) {
	bindings.CallLocationReplace(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryReplace calls the method "Location.replace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryReplace(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationReplace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasReload returns true if the method "Location.reload" exists.
func (this Location) HasReload() bool {
	return js.True == bindings.HasLocationReload(
		this.Ref(),
	)
}

// ReloadFunc returns the method "Location.reload".
func (this Location) ReloadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.LocationReloadFunc(
			this.Ref(),
		),
	)
}

// Reload calls the method "Location.reload".
func (this Location) Reload() (ret js.Void) {
	bindings.CallLocationReload(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReload calls the method "Location.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryReload() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationReload(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewReadableStreamDefaultReader(stream ReadableStream) (ret ReadableStreamDefaultReader) {
	ret.ref = bindings.NewReadableStreamDefaultReaderByReadableStreamDefaultReader(
		stream.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ReadableStreamDefaultReader) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetReadableStreamDefaultReaderClosed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRead returns true if the method "ReadableStreamDefaultReader.read" exists.
func (this ReadableStreamDefaultReader) HasRead() bool {
	return js.True == bindings.HasReadableStreamDefaultReaderRead(
		this.Ref(),
	)
}

// ReadFunc returns the method "ReadableStreamDefaultReader.read".
func (this ReadableStreamDefaultReader) ReadFunc() (fn js.Func[func() js.Promise[ReadableStreamReadResult]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderReadFunc(
			this.Ref(),
		),
	)
}

// Read calls the method "ReadableStreamDefaultReader.read".
func (this ReadableStreamDefaultReader) Read() (ret js.Promise[ReadableStreamReadResult]) {
	bindings.CallReadableStreamDefaultReaderRead(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRead calls the method "ReadableStreamDefaultReader.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryRead() (ret js.Promise[ReadableStreamReadResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderRead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReleaseLock returns true if the method "ReadableStreamDefaultReader.releaseLock" exists.
func (this ReadableStreamDefaultReader) HasReleaseLock() bool {
	return js.True == bindings.HasReadableStreamDefaultReaderReleaseLock(
		this.Ref(),
	)
}

// ReleaseLockFunc returns the method "ReadableStreamDefaultReader.releaseLock".
func (this ReadableStreamDefaultReader) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderReleaseLockFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "ReadableStreamDefaultReader.releaseLock".
func (this ReadableStreamDefaultReader) ReleaseLock() (ret js.Void) {
	bindings.CallReadableStreamDefaultReaderReleaseLock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "ReadableStreamDefaultReader.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderReleaseLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCancel returns true if the method "ReadableStreamDefaultReader.cancel" exists.
func (this ReadableStreamDefaultReader) HasCancel() bool {
	return js.True == bindings.HasReadableStreamDefaultReaderCancel(
		this.Ref(),
	)
}

// CancelFunc returns the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamDefaultReaderCancel(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStreamDefaultReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderCancel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasCancel1 returns true if the method "ReadableStreamDefaultReader.cancel" exists.
func (this ReadableStreamDefaultReader) HasCancel1() bool {
	return js.True == bindings.HasReadableStreamDefaultReaderCancel1(
		this.Ref(),
	)
}

// Cancel1Func returns the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamDefaultReaderCancel1Func(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamDefaultReaderCancel1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStreamDefaultReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderCancel1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewReadableStreamBYOBReader(stream ReadableStream) (ret ReadableStreamBYOBReader) {
	ret.ref = bindings.NewReadableStreamBYOBReaderByReadableStreamBYOBReader(
		stream.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ReadableStreamBYOBReader) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetReadableStreamBYOBReaderClosed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRead returns true if the method "ReadableStreamBYOBReader.read" exists.
func (this ReadableStreamBYOBReader) HasRead() bool {
	return js.True == bindings.HasReadableStreamBYOBReaderRead(
		this.Ref(),
	)
}

// ReadFunc returns the method "ReadableStreamBYOBReader.read".
func (this ReadableStreamBYOBReader) ReadFunc() (fn js.Func[func(view js.ArrayBufferView) js.Promise[ReadableStreamReadResult]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderReadFunc(
			this.Ref(),
		),
	)
}

// Read calls the method "ReadableStreamBYOBReader.read".
func (this ReadableStreamBYOBReader) Read(view js.ArrayBufferView) (ret js.Promise[ReadableStreamReadResult]) {
	bindings.CallReadableStreamBYOBReaderRead(
		this.Ref(), js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryRead calls the method "ReadableStreamBYOBReader.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryRead(view js.ArrayBufferView) (ret js.Promise[ReadableStreamReadResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderRead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

// HasReleaseLock returns true if the method "ReadableStreamBYOBReader.releaseLock" exists.
func (this ReadableStreamBYOBReader) HasReleaseLock() bool {
	return js.True == bindings.HasReadableStreamBYOBReaderReleaseLock(
		this.Ref(),
	)
}

// ReleaseLockFunc returns the method "ReadableStreamBYOBReader.releaseLock".
func (this ReadableStreamBYOBReader) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderReleaseLockFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "ReadableStreamBYOBReader.releaseLock".
func (this ReadableStreamBYOBReader) ReleaseLock() (ret js.Void) {
	bindings.CallReadableStreamBYOBReaderReleaseLock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "ReadableStreamBYOBReader.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderReleaseLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCancel returns true if the method "ReadableStreamBYOBReader.cancel" exists.
func (this ReadableStreamBYOBReader) HasCancel() bool {
	return js.True == bindings.HasReadableStreamBYOBReaderCancel(
		this.Ref(),
	)
}

// CancelFunc returns the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamBYOBReaderCancel(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStreamBYOBReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderCancel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasCancel1 returns true if the method "ReadableStreamBYOBReader.cancel" exists.
func (this ReadableStreamBYOBReader) HasCancel1() bool {
	return js.True == bindings.HasReadableStreamBYOBReaderCancel1(
		this.Ref(),
	)
}

// Cancel1Func returns the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamBYOBReaderCancel1Func(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamBYOBReaderCancel1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStreamBYOBReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderCancel1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewWritableStreamDefaultWriter(stream WritableStream) (ret WritableStreamDefaultWriter) {
	ret.ref = bindings.NewWritableStreamDefaultWriterByWritableStreamDefaultWriter(
		stream.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterClosed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DesiredSize returns the value of property "WritableStreamDefaultWriter.desiredSize".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) DesiredSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterDesiredSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "WritableStreamDefaultWriter.ready".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAbort returns true if the method "WritableStreamDefaultWriter.abort" exists.
func (this WritableStreamDefaultWriter) HasAbort() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) AbortFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) Abort(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterAbort(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the method "WritableStreamDefaultWriter.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryAbort(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasAbort1 returns true if the method "WritableStreamDefaultWriter.abort" exists.
func (this WritableStreamDefaultWriter) HasAbort1() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterAbort1(
		this.Ref(),
	)
}

// Abort1Func returns the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) Abort1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterAbort1Func(
			this.Ref(),
		),
	)
}

// Abort1 calls the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) Abort1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterAbort1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the method "WritableStreamDefaultWriter.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryAbort1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterAbort1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "WritableStreamDefaultWriter.close" exists.
func (this WritableStreamDefaultWriter) HasClose() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "WritableStreamDefaultWriter.close".
func (this WritableStreamDefaultWriter) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "WritableStreamDefaultWriter.close".
func (this WritableStreamDefaultWriter) Close() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "WritableStreamDefaultWriter.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReleaseLock returns true if the method "WritableStreamDefaultWriter.releaseLock" exists.
func (this WritableStreamDefaultWriter) HasReleaseLock() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterReleaseLock(
		this.Ref(),
	)
}

// ReleaseLockFunc returns the method "WritableStreamDefaultWriter.releaseLock".
func (this WritableStreamDefaultWriter) ReleaseLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterReleaseLockFunc(
			this.Ref(),
		),
	)
}

// ReleaseLock calls the method "WritableStreamDefaultWriter.releaseLock".
func (this WritableStreamDefaultWriter) ReleaseLock() (ret js.Void) {
	bindings.CallWritableStreamDefaultWriterReleaseLock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "WritableStreamDefaultWriter.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterReleaseLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWrite returns true if the method "WritableStreamDefaultWriter.write" exists.
func (this WritableStreamDefaultWriter) HasWrite() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) WriteFunc() (fn js.Func[func(chunk js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) Write(chunk js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterWrite(
		this.Ref(), js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryWrite calls the method "WritableStreamDefaultWriter.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryWrite(chunk js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterWrite(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasWrite1 returns true if the method "WritableStreamDefaultWriter.write" exists.
func (this WritableStreamDefaultWriter) HasWrite1() bool {
	return js.True == bindings.HasWritableStreamDefaultWriterWrite1(
		this.Ref(),
	)
}

// Write1Func returns the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) Write1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultWriterWrite1Func(
			this.Ref(),
		),
	)
}

// Write1 calls the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) Write1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterWrite1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryWrite1 calls the method "WritableStreamDefaultWriter.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryWrite1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterWrite1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewWritableStream(underlyingSink js.Object, strategy QueuingStrategy) (ret WritableStream) {
	ret.ref = bindings.NewWritableStreamByWritableStream(
		underlyingSink.Ref(),
		js.Pointer(&strategy))
	return
}

func NewWritableStreamByWritableStream1(underlyingSink js.Object) (ret WritableStream) {
	ret.ref = bindings.NewWritableStreamByWritableStream1(
		underlyingSink.Ref())
	return
}

func NewWritableStreamByWritableStream2() (ret WritableStream) {
	ret.ref = bindings.NewWritableStreamByWritableStream2()
	return
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
// It returns ok=false if there is no such property.
func (this WritableStream) Locked() (ret bool, ok bool) {
	ok = js.True == bindings.GetWritableStreamLocked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAbort returns true if the method "WritableStream.abort" exists.
func (this WritableStream) HasAbort() bool {
	return js.True == bindings.HasWritableStreamAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "WritableStream.abort".
func (this WritableStream) AbortFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "WritableStream.abort".
func (this WritableStream) Abort(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamAbort(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the method "WritableStream.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryAbort(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasAbort1 returns true if the method "WritableStream.abort" exists.
func (this WritableStream) HasAbort1() bool {
	return js.True == bindings.HasWritableStreamAbort1(
		this.Ref(),
	)
}

// Abort1Func returns the method "WritableStream.abort".
func (this WritableStream) Abort1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamAbort1Func(
			this.Ref(),
		),
	)
}

// Abort1 calls the method "WritableStream.abort".
func (this WritableStream) Abort1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamAbort1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the method "WritableStream.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryAbort1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamAbort1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "WritableStream.close" exists.
func (this WritableStream) HasClose() bool {
	return js.True == bindings.HasWritableStreamClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "WritableStream.close".
func (this WritableStream) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WritableStreamCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "WritableStream.close".
func (this WritableStream) Close() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "WritableStream.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetWriter returns true if the method "WritableStream.getWriter" exists.
func (this WritableStream) HasGetWriter() bool {
	return js.True == bindings.HasWritableStreamGetWriter(
		this.Ref(),
	)
}

// GetWriterFunc returns the method "WritableStream.getWriter".
func (this WritableStream) GetWriterFunc() (fn js.Func[func() WritableStreamDefaultWriter]) {
	return fn.FromRef(
		bindings.WritableStreamGetWriterFunc(
			this.Ref(),
		),
	)
}

// GetWriter calls the method "WritableStream.getWriter".
func (this WritableStream) GetWriter() (ret WritableStreamDefaultWriter) {
	bindings.CallWritableStreamGetWriter(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetWriter calls the method "WritableStream.getWriter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryGetWriter() (ret WritableStreamDefaultWriter, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamGetWriter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewReadableStream(underlyingSource js.Object, strategy QueuingStrategy) (ret ReadableStream) {
	ret.ref = bindings.NewReadableStreamByReadableStream(
		underlyingSource.Ref(),
		js.Pointer(&strategy))
	return
}

func NewReadableStreamByReadableStream1(underlyingSource js.Object) (ret ReadableStream) {
	ret.ref = bindings.NewReadableStreamByReadableStream1(
		underlyingSource.Ref())
	return
}

func NewReadableStreamByReadableStream2() (ret ReadableStream) {
	ret.ref = bindings.NewReadableStreamByReadableStream2()
	return
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
// It returns ok=false if there is no such property.
func (this ReadableStream) Locked() (ret bool, ok bool) {
	ok = js.True == bindings.GetReadableStreamLocked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFrom returns true if the staticmethod "ReadableStream.from" exists.
func (this ReadableStream) HasFrom() bool {
	return js.True == bindings.HasReadableStreamFrom(
		this.Ref(),
	)
}

// FromFunc returns the staticmethod "ReadableStream.from".
func (this ReadableStream) FromFunc() (fn js.Func[func(asyncIterable js.Any) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamFromFunc(
			this.Ref(),
		),
	)
}

// From calls the staticmethod "ReadableStream.from".
func (this ReadableStream) From(asyncIterable js.Any) (ret ReadableStream) {
	bindings.CallReadableStreamFrom(
		this.Ref(), js.Pointer(&ret),
		asyncIterable.Ref(),
	)

	return
}

// TryFrom calls the staticmethod "ReadableStream.from"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryFrom(asyncIterable js.Any) (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamFrom(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		asyncIterable.Ref(),
	)

	return
}

// HasCancel returns true if the method "ReadableStream.cancel" exists.
func (this ReadableStream) HasCancel() bool {
	return js.True == bindings.HasReadableStreamCancel(
		this.Ref(),
	)
}

// CancelFunc returns the method "ReadableStream.cancel".
func (this ReadableStream) CancelFunc() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamCancelFunc(
			this.Ref(),
		),
	)
}

// Cancel calls the method "ReadableStream.cancel".
func (this ReadableStream) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamCancel(
		this.Ref(), js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStream.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamCancel(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasCancel1 returns true if the method "ReadableStream.cancel" exists.
func (this ReadableStream) HasCancel1() bool {
	return js.True == bindings.HasReadableStreamCancel1(
		this.Ref(),
	)
}

// Cancel1Func returns the method "ReadableStream.cancel".
func (this ReadableStream) Cancel1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamCancel1Func(
			this.Ref(),
		),
	)
}

// Cancel1 calls the method "ReadableStream.cancel".
func (this ReadableStream) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamCancel1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStream.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamCancel1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetReader returns true if the method "ReadableStream.getReader" exists.
func (this ReadableStream) HasGetReader() bool {
	return js.True == bindings.HasReadableStreamGetReader(
		this.Ref(),
	)
}

// GetReaderFunc returns the method "ReadableStream.getReader".
func (this ReadableStream) GetReaderFunc() (fn js.Func[func(options ReadableStreamGetReaderOptions) ReadableStreamReader]) {
	return fn.FromRef(
		bindings.ReadableStreamGetReaderFunc(
			this.Ref(),
		),
	)
}

// GetReader calls the method "ReadableStream.getReader".
func (this ReadableStream) GetReader(options ReadableStreamGetReaderOptions) (ret ReadableStreamReader) {
	bindings.CallReadableStreamGetReader(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetReader calls the method "ReadableStream.getReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryGetReader(options ReadableStreamGetReaderOptions) (ret ReadableStreamReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamGetReader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetReader1 returns true if the method "ReadableStream.getReader" exists.
func (this ReadableStream) HasGetReader1() bool {
	return js.True == bindings.HasReadableStreamGetReader1(
		this.Ref(),
	)
}

// GetReader1Func returns the method "ReadableStream.getReader".
func (this ReadableStream) GetReader1Func() (fn js.Func[func() ReadableStreamReader]) {
	return fn.FromRef(
		bindings.ReadableStreamGetReader1Func(
			this.Ref(),
		),
	)
}

// GetReader1 calls the method "ReadableStream.getReader".
func (this ReadableStream) GetReader1() (ret ReadableStreamReader) {
	bindings.CallReadableStreamGetReader1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetReader1 calls the method "ReadableStream.getReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryGetReader1() (ret ReadableStreamReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamGetReader1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPipeThrough returns true if the method "ReadableStream.pipeThrough" exists.
func (this ReadableStream) HasPipeThrough() bool {
	return js.True == bindings.HasReadableStreamPipeThrough(
		this.Ref(),
	)
}

// PipeThroughFunc returns the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThroughFunc() (fn js.Func[func(transform ReadableWritablePair, options StreamPipeOptions) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeThroughFunc(
			this.Ref(),
		),
	)
}

// PipeThrough calls the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThrough(transform ReadableWritablePair, options StreamPipeOptions) (ret ReadableStream) {
	bindings.CallReadableStreamPipeThrough(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&transform),
		js.Pointer(&options),
	)

	return
}

// TryPipeThrough calls the method "ReadableStream.pipeThrough"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeThrough(transform ReadableWritablePair, options StreamPipeOptions) (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeThrough(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
		js.Pointer(&options),
	)

	return
}

// HasPipeThrough1 returns true if the method "ReadableStream.pipeThrough" exists.
func (this ReadableStream) HasPipeThrough1() bool {
	return js.True == bindings.HasReadableStreamPipeThrough1(
		this.Ref(),
	)
}

// PipeThrough1Func returns the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThrough1Func() (fn js.Func[func(transform ReadableWritablePair) ReadableStream]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeThrough1Func(
			this.Ref(),
		),
	)
}

// PipeThrough1 calls the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThrough1(transform ReadableWritablePair) (ret ReadableStream) {
	bindings.CallReadableStreamPipeThrough1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TryPipeThrough1 calls the method "ReadableStream.pipeThrough"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeThrough1(transform ReadableWritablePair) (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeThrough1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasPipeTo returns true if the method "ReadableStream.pipeTo" exists.
func (this ReadableStream) HasPipeTo() bool {
	return js.True == bindings.HasReadableStreamPipeTo(
		this.Ref(),
	)
}

// PipeToFunc returns the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeToFunc() (fn js.Func[func(destination WritableStream, options StreamPipeOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeToFunc(
			this.Ref(),
		),
	)
}

// PipeTo calls the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeTo(destination WritableStream, options StreamPipeOptions) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamPipeTo(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPipeTo calls the method "ReadableStream.pipeTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeTo(destination WritableStream, options StreamPipeOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPipeTo1 returns true if the method "ReadableStream.pipeTo" exists.
func (this ReadableStream) HasPipeTo1() bool {
	return js.True == bindings.HasReadableStreamPipeTo1(
		this.Ref(),
	)
}

// PipeTo1Func returns the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeTo1Func() (fn js.Func[func(destination WritableStream) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ReadableStreamPipeTo1Func(
			this.Ref(),
		),
	)
}

// PipeTo1 calls the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeTo1(destination WritableStream) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamPipeTo1(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryPipeTo1 calls the method "ReadableStream.pipeTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeTo1(destination WritableStream) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeTo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
}

// HasTee returns true if the method "ReadableStream.tee" exists.
func (this ReadableStream) HasTee() bool {
	return js.True == bindings.HasReadableStreamTee(
		this.Ref(),
	)
}

// TeeFunc returns the method "ReadableStream.tee".
func (this ReadableStream) TeeFunc() (fn js.Func[func() js.Array[ReadableStream]]) {
	return fn.FromRef(
		bindings.ReadableStreamTeeFunc(
			this.Ref(),
		),
	)
}

// Tee calls the method "ReadableStream.tee".
func (this ReadableStream) Tee() (ret js.Array[ReadableStream]) {
	bindings.CallReadableStreamTee(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTee calls the method "ReadableStream.tee"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryTee() (ret js.Array[ReadableStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamTee(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewBlob(blobParts js.Array[BlobPart], options BlobPropertyBag) (ret Blob) {
	ret.ref = bindings.NewBlobByBlob(
		blobParts.Ref(),
		js.Pointer(&options))
	return
}

func NewBlobByBlob1(blobParts js.Array[BlobPart]) (ret Blob) {
	ret.ref = bindings.NewBlobByBlob1(
		blobParts.Ref())
	return
}

func NewBlobByBlob2() (ret Blob) {
	ret.ref = bindings.NewBlobByBlob2()
	return
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
// It returns ok=false if there is no such property.
func (this Blob) Size() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBlobSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Blob.type".
//
// It returns ok=false if there is no such property.
func (this Blob) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBlobType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSlice returns true if the method "Blob.slice" exists.
func (this Blob) HasSlice() bool {
	return js.True == bindings.HasBlobSlice(
		this.Ref(),
	)
}

// SliceFunc returns the method "Blob.slice".
func (this Blob) SliceFunc() (fn js.Func[func(start int64, end int64, contentType js.String) Blob]) {
	return fn.FromRef(
		bindings.BlobSliceFunc(
			this.Ref(),
		),
	)
}

// Slice calls the method "Blob.slice".
func (this Blob) Slice(start int64, end int64, contentType js.String) (ret Blob) {
	bindings.CallBlobSlice(
		this.Ref(), js.Pointer(&ret),
		float64(start),
		float64(end),
		contentType.Ref(),
	)

	return
}

// TrySlice calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice(start int64, end int64, contentType js.String) (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
		contentType.Ref(),
	)

	return
}

// HasSlice1 returns true if the method "Blob.slice" exists.
func (this Blob) HasSlice1() bool {
	return js.True == bindings.HasBlobSlice1(
		this.Ref(),
	)
}

// Slice1Func returns the method "Blob.slice".
func (this Blob) Slice1Func() (fn js.Func[func(start int64, end int64) Blob]) {
	return fn.FromRef(
		bindings.BlobSlice1Func(
			this.Ref(),
		),
	)
}

// Slice1 calls the method "Blob.slice".
func (this Blob) Slice1(start int64, end int64) (ret Blob) {
	bindings.CallBlobSlice1(
		this.Ref(), js.Pointer(&ret),
		float64(start),
		float64(end),
	)

	return
}

// TrySlice1 calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice1(start int64, end int64) (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
	)

	return
}

// HasSlice2 returns true if the method "Blob.slice" exists.
func (this Blob) HasSlice2() bool {
	return js.True == bindings.HasBlobSlice2(
		this.Ref(),
	)
}

// Slice2Func returns the method "Blob.slice".
func (this Blob) Slice2Func() (fn js.Func[func(start int64) Blob]) {
	return fn.FromRef(
		bindings.BlobSlice2Func(
			this.Ref(),
		),
	)
}

// Slice2 calls the method "Blob.slice".
func (this Blob) Slice2(start int64) (ret Blob) {
	bindings.CallBlobSlice2(
		this.Ref(), js.Pointer(&ret),
		float64(start),
	)

	return
}

// TrySlice2 calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice2(start int64) (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
	)

	return
}

// HasSlice3 returns true if the method "Blob.slice" exists.
func (this Blob) HasSlice3() bool {
	return js.True == bindings.HasBlobSlice3(
		this.Ref(),
	)
}

// Slice3Func returns the method "Blob.slice".
func (this Blob) Slice3Func() (fn js.Func[func() Blob]) {
	return fn.FromRef(
		bindings.BlobSlice3Func(
			this.Ref(),
		),
	)
}

// Slice3 calls the method "Blob.slice".
func (this Blob) Slice3() (ret Blob) {
	bindings.CallBlobSlice3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySlice3 calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice3() (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStream returns true if the method "Blob.stream" exists.
func (this Blob) HasStream() bool {
	return js.True == bindings.HasBlobStream(
		this.Ref(),
	)
}

// StreamFunc returns the method "Blob.stream".
func (this Blob) StreamFunc() (fn js.Func[func() ReadableStream]) {
	return fn.FromRef(
		bindings.BlobStreamFunc(
			this.Ref(),
		),
	)
}

// Stream calls the method "Blob.stream".
func (this Blob) Stream() (ret ReadableStream) {
	bindings.CallBlobStream(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStream calls the method "Blob.stream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryStream() (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasText returns true if the method "Blob.text" exists.
func (this Blob) HasText() bool {
	return js.True == bindings.HasBlobText(
		this.Ref(),
	)
}

// TextFunc returns the method "Blob.text".
func (this Blob) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.BlobTextFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "Blob.text".
func (this Blob) Text() (ret js.Promise[js.String]) {
	bindings.CallBlobText(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Blob.text"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasArrayBuffer returns true if the method "Blob.arrayBuffer" exists.
func (this Blob) HasArrayBuffer() bool {
	return js.True == bindings.HasBlobArrayBuffer(
		this.Ref(),
	)
}

// ArrayBufferFunc returns the method "Blob.arrayBuffer".
func (this Blob) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.BlobArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Blob.arrayBuffer".
func (this Blob) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallBlobArrayBuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Blob.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobArrayBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewFile(fileBits js.Array[BlobPart], fileName js.String, options FilePropertyBag) (ret File) {
	ret.ref = bindings.NewFileByFile(
		fileBits.Ref(),
		fileName.Ref(),
		js.Pointer(&options))
	return
}

func NewFileByFile1(fileBits js.Array[BlobPart], fileName js.String) (ret File) {
	ret.ref = bindings.NewFileByFile1(
		fileBits.Ref(),
		fileName.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this File) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastModified returns the value of property "File.lastModified".
//
// It returns ok=false if there is no such property.
func (this File) LastModified() (ret int64, ok bool) {
	ok = js.True == bindings.GetFileLastModified(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WebkitRelativePath returns the value of property "File.webkitRelativePath".
//
// It returns ok=false if there is no such property.
func (this File) WebkitRelativePath() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileWebkitRelativePath(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this RadioNodeList) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRadioNodeListValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "RadioNodeList.value" to val.
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

// HasNamedItem returns true if the method "HTMLFormControlsCollection.namedItem" exists.
func (this HTMLFormControlsCollection) HasNamedItem() bool {
	return js.True == bindings.HasHTMLFormControlsCollectionNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "HTMLFormControlsCollection.namedItem".
func (this HTMLFormControlsCollection) NamedItemFunc() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	return fn.FromRef(
		bindings.HTMLFormControlsCollectionNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLFormControlsCollection.namedItem".
func (this HTMLFormControlsCollection) NamedItem(name js.String) (ret OneOf_RadioNodeList_Element) {
	bindings.CallHTMLFormControlsCollectionNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLFormControlsCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormControlsCollection) TryNamedItem(name js.String) (ret OneOf_RadioNodeList_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormControlsCollectionNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

func NewHTMLFormElement() (ret HTMLFormElement) {
	ret.ref = bindings.NewHTMLFormElementByHTMLFormElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) AcceptCharset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAcceptCharset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAcceptCharset sets the value of property "HTMLFormElement.acceptCharset" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Action() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAction sets the value of property "HTMLFormElement.action" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAutocomplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLFormElement.autocomplete" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Enctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementEnctype(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEnctype sets the value of property "HTMLFormElement.enctype" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEncoding sets the value of property "HTMLFormElement.encoding" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Method() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMethod sets the value of property "HTMLFormElement.method" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFormElement.name" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) NoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementNoValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNoValidate sets the value of property "HTMLFormElement.noValidate" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLFormElement.target" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementRel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLFormElement.rel" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLFormElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementRelList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Elements returns the value of property "HTMLFormElement.elements".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Elements() (ret HTMLFormControlsCollection, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "HTMLFormElement.length".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "HTMLFormElement." exists.
func (this HTMLFormElement) HasGet() bool {
	return js.True == bindings.HasHTMLFormElementGet(
		this.Ref(),
	)
}

// GetFunc returns the method "HTMLFormElement.".
func (this HTMLFormElement) GetFunc() (fn js.Func[func(index uint32) Element]) {
	return fn.FromRef(
		bindings.HTMLFormElementGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "HTMLFormElement.".
func (this HTMLFormElement) Get(index uint32) (ret Element) {
	bindings.CallHTMLFormElementGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "HTMLFormElement."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryGet(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasGet1 returns true if the method "HTMLFormElement." exists.
func (this HTMLFormElement) HasGet1() bool {
	return js.True == bindings.HasHTMLFormElementGet1(
		this.Ref(),
	)
}

// Get1Func returns the method "HTMLFormElement.".
func (this HTMLFormElement) Get1Func() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	return fn.FromRef(
		bindings.HTMLFormElementGet1Func(
			this.Ref(),
		),
	)
}

// Get1 calls the method "HTMLFormElement.".
func (this HTMLFormElement) Get1(name js.String) (ret OneOf_RadioNodeList_Element) {
	bindings.CallHTMLFormElementGet1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet1 calls the method "HTMLFormElement."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryGet1(name js.String) (ret OneOf_RadioNodeList_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementGet1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSubmit returns true if the method "HTMLFormElement.submit" exists.
func (this HTMLFormElement) HasSubmit() bool {
	return js.True == bindings.HasHTMLFormElementSubmit(
		this.Ref(),
	)
}

// SubmitFunc returns the method "HTMLFormElement.submit".
func (this HTMLFormElement) SubmitFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementSubmitFunc(
			this.Ref(),
		),
	)
}

// Submit calls the method "HTMLFormElement.submit".
func (this HTMLFormElement) Submit() (ret js.Void) {
	bindings.CallHTMLFormElementSubmit(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySubmit calls the method "HTMLFormElement.submit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TrySubmit() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementSubmit(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestSubmit returns true if the method "HTMLFormElement.requestSubmit" exists.
func (this HTMLFormElement) HasRequestSubmit() bool {
	return js.True == bindings.HasHTMLFormElementRequestSubmit(
		this.Ref(),
	)
}

// RequestSubmitFunc returns the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmitFunc() (fn js.Func[func(submitter HTMLElement)]) {
	return fn.FromRef(
		bindings.HTMLFormElementRequestSubmitFunc(
			this.Ref(),
		),
	)
}

// RequestSubmit calls the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmit(submitter HTMLElement) (ret js.Void) {
	bindings.CallHTMLFormElementRequestSubmit(
		this.Ref(), js.Pointer(&ret),
		submitter.Ref(),
	)

	return
}

// TryRequestSubmit calls the method "HTMLFormElement.requestSubmit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryRequestSubmit(submitter HTMLElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementRequestSubmit(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		submitter.Ref(),
	)

	return
}

// HasRequestSubmit1 returns true if the method "HTMLFormElement.requestSubmit" exists.
func (this HTMLFormElement) HasRequestSubmit1() bool {
	return js.True == bindings.HasHTMLFormElementRequestSubmit1(
		this.Ref(),
	)
}

// RequestSubmit1Func returns the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmit1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementRequestSubmit1Func(
			this.Ref(),
		),
	)
}

// RequestSubmit1 calls the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmit1() (ret js.Void) {
	bindings.CallHTMLFormElementRequestSubmit1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestSubmit1 calls the method "HTMLFormElement.requestSubmit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryRequestSubmit1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementRequestSubmit1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "HTMLFormElement.reset" exists.
func (this HTMLFormElement) HasReset() bool {
	return js.True == bindings.HasHTMLFormElementReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "HTMLFormElement.reset".
func (this HTMLFormElement) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLFormElementResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "HTMLFormElement.reset".
func (this HTMLFormElement) Reset() (ret js.Void) {
	bindings.CallHTMLFormElementReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "HTMLFormElement.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCheckValidity returns true if the method "HTMLFormElement.checkValidity" exists.
func (this HTMLFormElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLFormElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLFormElement.checkValidity".
func (this HTMLFormElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFormElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLFormElement.checkValidity".
func (this HTMLFormElement) CheckValidity() (ret bool) {
	bindings.CallHTMLFormElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLFormElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLFormElement.reportValidity" exists.
func (this HTMLFormElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLFormElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLFormElement.reportValidity".
func (this HTMLFormElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFormElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLFormElement.reportValidity".
func (this HTMLFormElement) ReportValidity() (ret bool) {
	bindings.CallHTMLFormElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLFormElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewFormData(form HTMLFormElement, submitter HTMLElement) (ret FormData) {
	ret.ref = bindings.NewFormDataByFormData(
		form.Ref(),
		submitter.Ref())
	return
}

func NewFormDataByFormData1(form HTMLFormElement) (ret FormData) {
	ret.ref = bindings.NewFormDataByFormData1(
		form.Ref())
	return
}

func NewFormDataByFormData2() (ret FormData) {
	ret.ref = bindings.NewFormDataByFormData2()
	return
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

// HasAppend returns true if the method "FormData.append" exists.
func (this FormData) HasAppend() bool {
	return js.True == bindings.HasFormDataAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "FormData.append".
func (this FormData) AppendFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.FormDataAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "FormData.append".
func (this FormData) Append(name js.String, value js.String) (ret js.Void) {
	bindings.CallFormDataAppend(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TryAppend calls the method "FormData.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryAppend(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasAppend1 returns true if the method "FormData.append" exists.
func (this FormData) HasAppend1() bool {
	return js.True == bindings.HasFormDataAppend1(
		this.Ref(),
	)
}

// Append1Func returns the method "FormData.append".
func (this FormData) Append1Func() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	return fn.FromRef(
		bindings.FormDataAppend1Func(
			this.Ref(),
		),
	)
}

// Append1 calls the method "FormData.append".
func (this FormData) Append1(name js.String, blobValue Blob, filename js.String) (ret js.Void) {
	bindings.CallFormDataAppend1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// TryAppend1 calls the method "FormData.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryAppend1(name js.String, blobValue Blob, filename js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataAppend1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// HasAppend2 returns true if the method "FormData.append" exists.
func (this FormData) HasAppend2() bool {
	return js.True == bindings.HasFormDataAppend2(
		this.Ref(),
	)
}

// Append2Func returns the method "FormData.append".
func (this FormData) Append2Func() (fn js.Func[func(name js.String, blobValue Blob)]) {
	return fn.FromRef(
		bindings.FormDataAppend2Func(
			this.Ref(),
		),
	)
}

// Append2 calls the method "FormData.append".
func (this FormData) Append2(name js.String, blobValue Blob) (ret js.Void) {
	bindings.CallFormDataAppend2(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		blobValue.Ref(),
	)

	return
}

// TryAppend2 calls the method "FormData.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryAppend2(name js.String, blobValue Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataAppend2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
	)

	return
}

// HasDelete returns true if the method "FormData.delete" exists.
func (this FormData) HasDelete() bool {
	return js.True == bindings.HasFormDataDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "FormData.delete".
func (this FormData) DeleteFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.FormDataDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "FormData.delete".
func (this FormData) Delete(name js.String) (ret js.Void) {
	bindings.CallFormDataDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "FormData.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryDelete(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGet returns true if the method "FormData.get" exists.
func (this FormData) HasGet() bool {
	return js.True == bindings.HasFormDataGet(
		this.Ref(),
	)
}

// GetFunc returns the method "FormData.get".
func (this FormData) GetFunc() (fn js.Func[func(name js.String) FormDataEntryValue]) {
	return fn.FromRef(
		bindings.FormDataGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "FormData.get".
func (this FormData) Get(name js.String) (ret FormDataEntryValue) {
	bindings.CallFormDataGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "FormData.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryGet(name js.String) (ret FormDataEntryValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetAll returns true if the method "FormData.getAll" exists.
func (this FormData) HasGetAll() bool {
	return js.True == bindings.HasFormDataGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "FormData.getAll".
func (this FormData) GetAllFunc() (fn js.Func[func(name js.String) js.Array[FormDataEntryValue]]) {
	return fn.FromRef(
		bindings.FormDataGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "FormData.getAll".
func (this FormData) GetAll(name js.String) (ret js.Array[FormDataEntryValue]) {
	bindings.CallFormDataGetAll(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetAll calls the method "FormData.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryGetAll(name js.String) (ret js.Array[FormDataEntryValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasHas returns true if the method "FormData.has" exists.
func (this FormData) HasHas() bool {
	return js.True == bindings.HasFormDataHas(
		this.Ref(),
	)
}

// HasFunc returns the method "FormData.has".
func (this FormData) HasFunc() (fn js.Func[func(name js.String) bool]) {
	return fn.FromRef(
		bindings.FormDataHasFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "FormData.has".
func (this FormData) Has(name js.String) (ret bool) {
	bindings.CallFormDataHas(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryHas calls the method "FormData.has"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryHas(name js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataHas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSet returns true if the method "FormData.set" exists.
func (this FormData) HasSet() bool {
	return js.True == bindings.HasFormDataSet(
		this.Ref(),
	)
}

// SetFunc returns the method "FormData.set".
func (this FormData) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.FormDataSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "FormData.set".
func (this FormData) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallFormDataSet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "FormData.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TrySet(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasSet1 returns true if the method "FormData.set" exists.
func (this FormData) HasSet1() bool {
	return js.True == bindings.HasFormDataSet1(
		this.Ref(),
	)
}

// Set1Func returns the method "FormData.set".
func (this FormData) Set1Func() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	return fn.FromRef(
		bindings.FormDataSet1Func(
			this.Ref(),
		),
	)
}

// Set1 calls the method "FormData.set".
func (this FormData) Set1(name js.String, blobValue Blob, filename js.String) (ret js.Void) {
	bindings.CallFormDataSet1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// TrySet1 calls the method "FormData.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TrySet1(name js.String, blobValue Blob, filename js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataSet1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// HasSet2 returns true if the method "FormData.set" exists.
func (this FormData) HasSet2() bool {
	return js.True == bindings.HasFormDataSet2(
		this.Ref(),
	)
}

// Set2Func returns the method "FormData.set".
func (this FormData) Set2Func() (fn js.Func[func(name js.String, blobValue Blob)]) {
	return fn.FromRef(
		bindings.FormDataSet2Func(
			this.Ref(),
		),
	)
}

// Set2 calls the method "FormData.set".
func (this FormData) Set2(name js.String, blobValue Blob) (ret js.Void) {
	bindings.CallFormDataSet2(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		blobValue.Ref(),
	)

	return
}

// TrySet2 calls the method "FormData.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TrySet2(name js.String, blobValue Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataSet2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this ValidityState) ValueMissing() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateValueMissing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TypeMismatch returns the value of property "ValidityState.typeMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TypeMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTypeMismatch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PatternMismatch returns the value of property "ValidityState.patternMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) PatternMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStatePatternMismatch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TooLong returns the value of property "ValidityState.tooLong".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TooLong() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTooLong(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TooShort returns the value of property "ValidityState.tooShort".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TooShort() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTooShort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RangeUnderflow returns the value of property "ValidityState.rangeUnderflow".
//
// It returns ok=false if there is no such property.
func (this ValidityState) RangeUnderflow() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateRangeUnderflow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RangeOverflow returns the value of property "ValidityState.rangeOverflow".
//
// It returns ok=false if there is no such property.
func (this ValidityState) RangeOverflow() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateRangeOverflow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StepMismatch returns the value of property "ValidityState.stepMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) StepMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateStepMismatch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BadInput returns the value of property "ValidityState.badInput".
//
// It returns ok=false if there is no such property.
func (this ValidityState) BadInput() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateBadInput(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CustomError returns the value of property "ValidityState.customError".
//
// It returns ok=false if there is no such property.
func (this ValidityState) CustomError() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateCustomError(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Valid returns the value of property "ValidityState.valid".
//
// It returns ok=false if there is no such property.
func (this ValidityState) Valid() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateValid(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasAdd returns true if the method "CustomStateSet.add" exists.
func (this CustomStateSet) HasAdd() bool {
	return js.True == bindings.HasCustomStateSetAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "CustomStateSet.add".
func (this CustomStateSet) AddFunc() (fn js.Func[func(value js.String)]) {
	return fn.FromRef(
		bindings.CustomStateSetAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "CustomStateSet.add".
func (this CustomStateSet) Add(value js.String) (ret js.Void) {
	bindings.CallCustomStateSetAdd(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryAdd calls the method "CustomStateSet.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomStateSet) TryAdd(value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomStateSetAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this ElementInternals) ShadowRoot() (ret ShadowRoot, ok bool) {
	ok = js.True == bindings.GetElementInternalsShadowRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Form returns the value of property "ElementInternals.form".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetElementInternalsForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "ElementInternals.willValidate".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetElementInternalsWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "ElementInternals.validity".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetElementInternalsValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "ElementInternals.validationMessage".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "ElementInternals.labels".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetElementInternalsLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// States returns the value of property "ElementInternals.states".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) States() (ret CustomStateSet, ok bool) {
	ok = js.True == bindings.GetElementInternalsStates(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Role returns the value of property "ElementInternals.role".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Role() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsRole(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRole sets the value of property "ElementInternals.role" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaActiveDescendantElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaActiveDescendantElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaActiveDescendantElement sets the value of property "ElementInternals.ariaActiveDescendantElement" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaAtomic() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaAtomic(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaAtomic sets the value of property "ElementInternals.ariaAtomic" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaAutoComplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaAutoComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaAutoComplete sets the value of property "ElementInternals.ariaAutoComplete" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaBusy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaBusy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaBusy sets the value of property "ElementInternals.ariaBusy" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaChecked() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaChecked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaChecked sets the value of property "ElementInternals.ariaChecked" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColCount sets the value of property "ElementInternals.ariaColCount" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColIndex sets the value of property "ElementInternals.ariaColIndex" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColIndexText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColIndexText sets the value of property "ElementInternals.ariaColIndexText" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaColSpan sets the value of property "ElementInternals.ariaColSpan" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaControlsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaControlsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaControlsElements sets the value of property "ElementInternals.ariaControlsElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaCurrent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaCurrent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaCurrent sets the value of property "ElementInternals.ariaCurrent" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDescribedByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDescribedByElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDescribedByElements sets the value of property "ElementInternals.ariaDescribedByElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDescription sets the value of property "ElementInternals.ariaDescription" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDetailsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDetailsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDetailsElements sets the value of property "ElementInternals.ariaDetailsElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDisabled() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaDisabled sets the value of property "ElementInternals.ariaDisabled" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaErrorMessageElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaErrorMessageElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaErrorMessageElements sets the value of property "ElementInternals.ariaErrorMessageElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaExpanded() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaExpanded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaExpanded sets the value of property "ElementInternals.ariaExpanded" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaFlowToElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaFlowToElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaFlowToElements sets the value of property "ElementInternals.ariaFlowToElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaHasPopup() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaHasPopup(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaHasPopup sets the value of property "ElementInternals.ariaHasPopup" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaHidden() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaHidden(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaHidden sets the value of property "ElementInternals.ariaHidden" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaInvalid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaInvalid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaInvalid sets the value of property "ElementInternals.ariaInvalid" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaKeyShortcuts() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaKeyShortcuts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaKeyShortcuts sets the value of property "ElementInternals.ariaKeyShortcuts" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLabel sets the value of property "ElementInternals.ariaLabel" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLabelledByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLabelledByElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLabelledByElements sets the value of property "ElementInternals.ariaLabelledByElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLevel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLevel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLevel sets the value of property "ElementInternals.ariaLevel" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaLive sets the value of property "ElementInternals.ariaLive" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaModal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaModal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaModal sets the value of property "ElementInternals.ariaModal" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaMultiLine() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaMultiLine(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaMultiLine sets the value of property "ElementInternals.ariaMultiLine" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaMultiSelectable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaMultiSelectable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaMultiSelectable sets the value of property "ElementInternals.ariaMultiSelectable" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaOrientation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaOrientation sets the value of property "ElementInternals.ariaOrientation" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaOwnsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaOwnsElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaOwnsElements sets the value of property "ElementInternals.ariaOwnsElements" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPlaceholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPlaceholder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPlaceholder sets the value of property "ElementInternals.ariaPlaceholder" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPosInSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPosInSet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPosInSet sets the value of property "ElementInternals.ariaPosInSet" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPressed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPressed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaPressed sets the value of property "ElementInternals.ariaPressed" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaReadOnly() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaReadOnly(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaReadOnly sets the value of property "ElementInternals.ariaReadOnly" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRequired() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRequired(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRequired sets the value of property "ElementInternals.ariaRequired" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRoleDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRoleDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRoleDescription sets the value of property "ElementInternals.ariaRoleDescription" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowCount sets the value of property "ElementInternals.ariaRowCount" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndex sets the value of property "ElementInternals.ariaRowIndex" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowIndexText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndexText sets the value of property "ElementInternals.ariaRowIndexText" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaRowSpan sets the value of property "ElementInternals.ariaRowSpan" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSelected() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSelected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSelected sets the value of property "ElementInternals.ariaSelected" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSetSize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSetSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSetSize sets the value of property "ElementInternals.ariaSetSize" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSort() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaSort sets the value of property "ElementInternals.ariaSort" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueMax() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueMax sets the value of property "ElementInternals.ariaValueMax" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueMin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueMin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueMin sets the value of property "ElementInternals.ariaValueMin" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueNow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueNow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueNow sets the value of property "ElementInternals.ariaValueNow" to val.
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
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAriaValueText sets the value of property "ElementInternals.ariaValueText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueText(
		this.Ref(),
		val.Ref(),
	)
}

// HasSetFormValue returns true if the method "ElementInternals.setFormValue" exists.
func (this ElementInternals) HasSetFormValue() bool {
	return js.True == bindings.HasElementInternalsSetFormValue(
		this.Ref(),
	)
}

// SetFormValueFunc returns the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValueFunc() (fn js.Func[func(value OneOf_File_String_FormData, state OneOf_File_String_FormData)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetFormValueFunc(
			this.Ref(),
		),
	)
}

// SetFormValue calls the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValue(value OneOf_File_String_FormData, state OneOf_File_String_FormData) (ret js.Void) {
	bindings.CallElementInternalsSetFormValue(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
		state.Ref(),
	)

	return
}

// TrySetFormValue calls the method "ElementInternals.setFormValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetFormValue(value OneOf_File_String_FormData, state OneOf_File_String_FormData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetFormValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		state.Ref(),
	)

	return
}

// HasSetFormValue1 returns true if the method "ElementInternals.setFormValue" exists.
func (this ElementInternals) HasSetFormValue1() bool {
	return js.True == bindings.HasElementInternalsSetFormValue1(
		this.Ref(),
	)
}

// SetFormValue1Func returns the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValue1Func() (fn js.Func[func(value OneOf_File_String_FormData)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetFormValue1Func(
			this.Ref(),
		),
	)
}

// SetFormValue1 calls the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValue1(value OneOf_File_String_FormData) (ret js.Void) {
	bindings.CallElementInternalsSetFormValue1(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TrySetFormValue1 calls the method "ElementInternals.setFormValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetFormValue1(value OneOf_File_String_FormData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetFormValue1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasSetValidity returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasSetValidity() bool {
	return js.True == bindings.HasElementInternalsSetValidity(
		this.Ref(),
	)
}

// SetValidityFunc returns the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidityFunc() (fn js.Func[func(flags ValidityStateFlags, message js.String, anchor HTMLElement)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidityFunc(
			this.Ref(),
		),
	)
}

// SetValidity calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity(flags ValidityStateFlags, message js.String, anchor HTMLElement) (ret js.Void) {
	bindings.CallElementInternalsSetValidity(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&flags),
		message.Ref(),
		anchor.Ref(),
	)

	return
}

// TrySetValidity calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity(flags ValidityStateFlags, message js.String, anchor HTMLElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
		message.Ref(),
		anchor.Ref(),
	)

	return
}

// HasSetValidity1 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasSetValidity1() bool {
	return js.True == bindings.HasElementInternalsSetValidity1(
		this.Ref(),
	)
}

// SetValidity1Func returns the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity1Func() (fn js.Func[func(flags ValidityStateFlags, message js.String)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity1Func(
			this.Ref(),
		),
	)
}

// SetValidity1 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity1(flags ValidityStateFlags, message js.String) (ret js.Void) {
	bindings.CallElementInternalsSetValidity1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&flags),
		message.Ref(),
	)

	return
}

// TrySetValidity1 calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity1(flags ValidityStateFlags, message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
		message.Ref(),
	)

	return
}

// HasSetValidity2 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasSetValidity2() bool {
	return js.True == bindings.HasElementInternalsSetValidity2(
		this.Ref(),
	)
}

// SetValidity2Func returns the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity2Func() (fn js.Func[func(flags ValidityStateFlags)]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity2Func(
			this.Ref(),
		),
	)
}

// SetValidity2 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity2(flags ValidityStateFlags) (ret js.Void) {
	bindings.CallElementInternalsSetValidity2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&flags),
	)

	return
}

// TrySetValidity2 calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity2(flags ValidityStateFlags) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
	)

	return
}

// HasSetValidity3 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasSetValidity3() bool {
	return js.True == bindings.HasElementInternalsSetValidity3(
		this.Ref(),
	)
}

// SetValidity3Func returns the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity3Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ElementInternalsSetValidity3Func(
			this.Ref(),
		),
	)
}

// SetValidity3 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity3() (ret js.Void) {
	bindings.CallElementInternalsSetValidity3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetValidity3 calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity3() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCheckValidity returns true if the method "ElementInternals.checkValidity" exists.
func (this ElementInternals) HasCheckValidity() bool {
	return js.True == bindings.HasElementInternalsCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "ElementInternals.checkValidity".
func (this ElementInternals) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementInternalsCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "ElementInternals.checkValidity".
func (this ElementInternals) CheckValidity() (ret bool) {
	bindings.CallElementInternalsCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "ElementInternals.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "ElementInternals.reportValidity" exists.
func (this ElementInternals) HasReportValidity() bool {
	return js.True == bindings.HasElementInternalsReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "ElementInternals.reportValidity".
func (this ElementInternals) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.ElementInternalsReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "ElementInternals.reportValidity".
func (this ElementInternals) ReportValidity() (ret bool) {
	bindings.CallElementInternalsReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "ElementInternals.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewEditContext(options EditContextInit) (ret EditContext) {
	ret.ref = bindings.NewEditContextByEditContext(
		js.Pointer(&options))
	return
}

func NewEditContextByEditContext1() (ret EditContext) {
	ret.ref = bindings.NewEditContextByEditContext1()
	return
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
// It returns ok=false if there is no such property.
func (this EditContext) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEditContextText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "EditContext.selectionStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionEnd returns the value of property "EditContext.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CompositionRangeStart returns the value of property "EditContext.compositionRangeStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) CompositionRangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCompositionRangeStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CompositionRangeEnd returns the value of property "EditContext.compositionRangeEnd".
//
// It returns ok=false if there is no such property.
func (this EditContext) CompositionRangeEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCompositionRangeEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsComposing returns the value of property "EditContext.isComposing".
//
// It returns ok=false if there is no such property.
func (this EditContext) IsComposing() (ret bool, ok bool) {
	ok = js.True == bindings.GetEditContextIsComposing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ControlBounds returns the value of property "EditContext.controlBounds".
//
// It returns ok=false if there is no such property.
func (this EditContext) ControlBounds() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetEditContextControlBounds(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionBounds returns the value of property "EditContext.selectionBounds".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionBounds() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionBounds(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CharacterBoundsRangeStart returns the value of property "EditContext.characterBoundsRangeStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) CharacterBoundsRangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCharacterBoundsRangeStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUpdateText returns true if the method "EditContext.updateText" exists.
func (this EditContext) HasUpdateText() bool {
	return js.True == bindings.HasEditContextUpdateText(
		this.Ref(),
	)
}

// UpdateTextFunc returns the method "EditContext.updateText".
func (this EditContext) UpdateTextFunc() (fn js.Func[func(rangeStart uint32, rangeEnd uint32, text js.String)]) {
	return fn.FromRef(
		bindings.EditContextUpdateTextFunc(
			this.Ref(),
		),
	)
}

// UpdateText calls the method "EditContext.updateText".
func (this EditContext) UpdateText(rangeStart uint32, rangeEnd uint32, text js.String) (ret js.Void) {
	bindings.CallEditContextUpdateText(
		this.Ref(), js.Pointer(&ret),
		uint32(rangeStart),
		uint32(rangeEnd),
		text.Ref(),
	)

	return
}

// TryUpdateText calls the method "EditContext.updateText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateText(rangeStart uint32, rangeEnd uint32, text js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(rangeStart),
		uint32(rangeEnd),
		text.Ref(),
	)

	return
}

// HasUpdateSelection returns true if the method "EditContext.updateSelection" exists.
func (this EditContext) HasUpdateSelection() bool {
	return js.True == bindings.HasEditContextUpdateSelection(
		this.Ref(),
	)
}

// UpdateSelectionFunc returns the method "EditContext.updateSelection".
func (this EditContext) UpdateSelectionFunc() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.EditContextUpdateSelectionFunc(
			this.Ref(),
		),
	)
}

// UpdateSelection calls the method "EditContext.updateSelection".
func (this EditContext) UpdateSelection(start uint32, end uint32) (ret js.Void) {
	bindings.CallEditContextUpdateSelection(
		this.Ref(), js.Pointer(&ret),
		uint32(start),
		uint32(end),
	)

	return
}

// TryUpdateSelection calls the method "EditContext.updateSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateSelection(start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateSelection(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

// HasUpdateControlBounds returns true if the method "EditContext.updateControlBounds" exists.
func (this EditContext) HasUpdateControlBounds() bool {
	return js.True == bindings.HasEditContextUpdateControlBounds(
		this.Ref(),
	)
}

// UpdateControlBoundsFunc returns the method "EditContext.updateControlBounds".
func (this EditContext) UpdateControlBoundsFunc() (fn js.Func[func(controlBounds DOMRect)]) {
	return fn.FromRef(
		bindings.EditContextUpdateControlBoundsFunc(
			this.Ref(),
		),
	)
}

// UpdateControlBounds calls the method "EditContext.updateControlBounds".
func (this EditContext) UpdateControlBounds(controlBounds DOMRect) (ret js.Void) {
	bindings.CallEditContextUpdateControlBounds(
		this.Ref(), js.Pointer(&ret),
		controlBounds.Ref(),
	)

	return
}

// TryUpdateControlBounds calls the method "EditContext.updateControlBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateControlBounds(controlBounds DOMRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateControlBounds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		controlBounds.Ref(),
	)

	return
}

// HasUpdateSelectionBounds returns true if the method "EditContext.updateSelectionBounds" exists.
func (this EditContext) HasUpdateSelectionBounds() bool {
	return js.True == bindings.HasEditContextUpdateSelectionBounds(
		this.Ref(),
	)
}

// UpdateSelectionBoundsFunc returns the method "EditContext.updateSelectionBounds".
func (this EditContext) UpdateSelectionBoundsFunc() (fn js.Func[func(selectionBounds DOMRect)]) {
	return fn.FromRef(
		bindings.EditContextUpdateSelectionBoundsFunc(
			this.Ref(),
		),
	)
}

// UpdateSelectionBounds calls the method "EditContext.updateSelectionBounds".
func (this EditContext) UpdateSelectionBounds(selectionBounds DOMRect) (ret js.Void) {
	bindings.CallEditContextUpdateSelectionBounds(
		this.Ref(), js.Pointer(&ret),
		selectionBounds.Ref(),
	)

	return
}

// TryUpdateSelectionBounds calls the method "EditContext.updateSelectionBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateSelectionBounds(selectionBounds DOMRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateSelectionBounds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectionBounds.Ref(),
	)

	return
}

// HasUpdateCharacterBounds returns true if the method "EditContext.updateCharacterBounds" exists.
func (this EditContext) HasUpdateCharacterBounds() bool {
	return js.True == bindings.HasEditContextUpdateCharacterBounds(
		this.Ref(),
	)
}

// UpdateCharacterBoundsFunc returns the method "EditContext.updateCharacterBounds".
func (this EditContext) UpdateCharacterBoundsFunc() (fn js.Func[func(rangeStart uint32, characterBounds js.Array[DOMRect])]) {
	return fn.FromRef(
		bindings.EditContextUpdateCharacterBoundsFunc(
			this.Ref(),
		),
	)
}

// UpdateCharacterBounds calls the method "EditContext.updateCharacterBounds".
func (this EditContext) UpdateCharacterBounds(rangeStart uint32, characterBounds js.Array[DOMRect]) (ret js.Void) {
	bindings.CallEditContextUpdateCharacterBounds(
		this.Ref(), js.Pointer(&ret),
		uint32(rangeStart),
		characterBounds.Ref(),
	)

	return
}

// TryUpdateCharacterBounds calls the method "EditContext.updateCharacterBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateCharacterBounds(rangeStart uint32, characterBounds js.Array[DOMRect]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateCharacterBounds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(rangeStart),
		characterBounds.Ref(),
	)

	return
}

// HasAttachedElements returns true if the method "EditContext.attachedElements" exists.
func (this EditContext) HasAttachedElements() bool {
	return js.True == bindings.HasEditContextAttachedElements(
		this.Ref(),
	)
}

// AttachedElementsFunc returns the method "EditContext.attachedElements".
func (this EditContext) AttachedElementsFunc() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.EditContextAttachedElementsFunc(
			this.Ref(),
		),
	)
}

// AttachedElements calls the method "EditContext.attachedElements".
func (this EditContext) AttachedElements() (ret js.Array[Element]) {
	bindings.CallEditContextAttachedElements(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAttachedElements calls the method "EditContext.attachedElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryAttachedElements() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextAttachedElements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCharacterBounds returns true if the method "EditContext.characterBounds" exists.
func (this EditContext) HasCharacterBounds() bool {
	return js.True == bindings.HasEditContextCharacterBounds(
		this.Ref(),
	)
}

// CharacterBoundsFunc returns the method "EditContext.characterBounds".
func (this EditContext) CharacterBoundsFunc() (fn js.Func[func() js.Array[DOMRect]]) {
	return fn.FromRef(
		bindings.EditContextCharacterBoundsFunc(
			this.Ref(),
		),
	)
}

// CharacterBounds calls the method "EditContext.characterBounds".
func (this EditContext) CharacterBounds() (ret js.Array[DOMRect]) {
	bindings.CallEditContextCharacterBounds(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCharacterBounds calls the method "EditContext.characterBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryCharacterBounds() (ret js.Array[DOMRect], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextCharacterBounds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLElement() (ret HTMLElement) {
	ret.ref = bindings.NewHTMLElementByHTMLElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementTitle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "HTMLElement.title" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementLang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLang sets the value of property "HTMLElement.lang" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Translate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementTranslate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTranslate sets the value of property "HTMLElement.translate" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Dir() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementDir(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDir sets the value of property "HTMLElement.dir" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Hidden() (ret OneOf_Bool_Float64_String, ok bool) {
	ok = js.True == bindings.GetHTMLElementHidden(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHidden sets the value of property "HTMLElement.hidden" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Inert() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementInert(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInert sets the value of property "HTMLElement.inert" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) AccessKey() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAccessKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAccessKey sets the value of property "HTMLElement.accessKey" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) AccessKeyLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAccessKeyLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Draggable returns the value of property "HTMLElement.draggable".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Draggable() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementDraggable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDraggable sets the value of property "HTMLElement.draggable" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Spellcheck() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementSpellcheck(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSpellcheck sets the value of property "HTMLElement.spellcheck" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Autocapitalize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAutocapitalize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutocapitalize sets the value of property "HTMLElement.autocapitalize" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) InnerText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementInnerText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInnerText sets the value of property "HTMLElement.innerText" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) OuterText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementOuterText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOuterText sets the value of property "HTMLElement.outerText" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Popover() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementPopover(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPopover sets the value of property "HTMLElement.popover" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetParent() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetParent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetTop returns the value of property "HTMLElement.offsetTop".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetLeft returns the value of property "HTMLElement.offsetLeft".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetWidth returns the value of property "HTMLElement.offsetWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetHeight returns the value of property "HTMLElement.offsetHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EditContext returns the value of property "HTMLElement.editContext".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) EditContext() (ret EditContext, ok bool) {
	ok = js.True == bindings.GetHTMLElementEditContext(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEditContext sets the value of property "HTMLElement.editContext" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetHTMLElementStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "HTMLElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetHTMLElementAttributeStyleMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "HTMLElement.dataset".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetHTMLElementDataset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "HTMLElement.nonce".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementNonce(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "HTMLElement.nonce" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementAutofocus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "HTMLElement.autofocus" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementTabIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "HTMLElement.tabIndex" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) ContentEditable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementContentEditable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetContentEditable sets the value of property "HTMLElement.contentEditable" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) EnterKeyHint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementEnterKeyHint(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEnterKeyHint sets the value of property "HTMLElement.enterKeyHint" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) IsContentEditable() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementIsContentEditable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InputMode returns the value of property "HTMLElement.inputMode".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) InputMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementInputMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetInputMode sets the value of property "HTMLElement.inputMode" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLElement) VirtualKeyboardPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementVirtualKeyboardPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVirtualKeyboardPolicy sets the value of property "HTMLElement.virtualKeyboardPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetVirtualKeyboardPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLElementVirtualKeyboardPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// HasClick returns true if the method "HTMLElement.click" exists.
func (this HTMLElement) HasClick() bool {
	return js.True == bindings.HasHTMLElementClick(
		this.Ref(),
	)
}

// ClickFunc returns the method "HTMLElement.click".
func (this HTMLElement) ClickFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementClickFunc(
			this.Ref(),
		),
	)
}

// Click calls the method "HTMLElement.click".
func (this HTMLElement) Click() (ret js.Void) {
	bindings.CallHTMLElementClick(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClick calls the method "HTMLElement.click"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryClick() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementClick(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAttachInternals returns true if the method "HTMLElement.attachInternals" exists.
func (this HTMLElement) HasAttachInternals() bool {
	return js.True == bindings.HasHTMLElementAttachInternals(
		this.Ref(),
	)
}

// AttachInternalsFunc returns the method "HTMLElement.attachInternals".
func (this HTMLElement) AttachInternalsFunc() (fn js.Func[func() ElementInternals]) {
	return fn.FromRef(
		bindings.HTMLElementAttachInternalsFunc(
			this.Ref(),
		),
	)
}

// AttachInternals calls the method "HTMLElement.attachInternals".
func (this HTMLElement) AttachInternals() (ret ElementInternals) {
	bindings.CallHTMLElementAttachInternals(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAttachInternals calls the method "HTMLElement.attachInternals"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryAttachInternals() (ret ElementInternals, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementAttachInternals(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasShowPopover returns true if the method "HTMLElement.showPopover" exists.
func (this HTMLElement) HasShowPopover() bool {
	return js.True == bindings.HasHTMLElementShowPopover(
		this.Ref(),
	)
}

// ShowPopoverFunc returns the method "HTMLElement.showPopover".
func (this HTMLElement) ShowPopoverFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementShowPopoverFunc(
			this.Ref(),
		),
	)
}

// ShowPopover calls the method "HTMLElement.showPopover".
func (this HTMLElement) ShowPopover() (ret js.Void) {
	bindings.CallHTMLElementShowPopover(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowPopover calls the method "HTMLElement.showPopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryShowPopover() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementShowPopover(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHidePopover returns true if the method "HTMLElement.hidePopover" exists.
func (this HTMLElement) HasHidePopover() bool {
	return js.True == bindings.HasHTMLElementHidePopover(
		this.Ref(),
	)
}

// HidePopoverFunc returns the method "HTMLElement.hidePopover".
func (this HTMLElement) HidePopoverFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementHidePopoverFunc(
			this.Ref(),
		),
	)
}

// HidePopover calls the method "HTMLElement.hidePopover".
func (this HTMLElement) HidePopover() (ret js.Void) {
	bindings.CallHTMLElementHidePopover(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHidePopover calls the method "HTMLElement.hidePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryHidePopover() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementHidePopover(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTogglePopover returns true if the method "HTMLElement.togglePopover" exists.
func (this HTMLElement) HasTogglePopover() bool {
	return js.True == bindings.HasHTMLElementTogglePopover(
		this.Ref(),
	)
}

// TogglePopoverFunc returns the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopoverFunc() (fn js.Func[func(force bool) bool]) {
	return fn.FromRef(
		bindings.HTMLElementTogglePopoverFunc(
			this.Ref(),
		),
	)
}

// TogglePopover calls the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopover(force bool) (ret bool) {
	bindings.CallHTMLElementTogglePopover(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(force)),
	)

	return
}

// TryTogglePopover calls the method "HTMLElement.togglePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryTogglePopover(force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementTogglePopover(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(force)),
	)

	return
}

// HasTogglePopover1 returns true if the method "HTMLElement.togglePopover" exists.
func (this HTMLElement) HasTogglePopover1() bool {
	return js.True == bindings.HasHTMLElementTogglePopover1(
		this.Ref(),
	)
}

// TogglePopover1Func returns the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopover1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLElementTogglePopover1Func(
			this.Ref(),
		),
	)
}

// TogglePopover1 calls the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopover1() (ret bool) {
	bindings.CallHTMLElementTogglePopover1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTogglePopover1 calls the method "HTMLElement.togglePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryTogglePopover1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementTogglePopover1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFocus returns true if the method "HTMLElement.focus" exists.
func (this HTMLElement) HasFocus() bool {
	return js.True == bindings.HasHTMLElementFocus(
		this.Ref(),
	)
}

// FocusFunc returns the method "HTMLElement.focus".
func (this HTMLElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.HTMLElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "HTMLElement.focus".
func (this HTMLElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallHTMLElementFocus(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "HTMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFocus1 returns true if the method "HTMLElement.focus" exists.
func (this HTMLElement) HasFocus1() bool {
	return js.True == bindings.HasHTMLElementFocus1(
		this.Ref(),
	)
}

// Focus1Func returns the method "HTMLElement.focus".
func (this HTMLElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementFocus1Func(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "HTMLElement.focus".
func (this HTMLElement) Focus1() (ret js.Void) {
	bindings.CallHTMLElementFocus1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "HTMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementFocus1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlur returns true if the method "HTMLElement.blur" exists.
func (this HTMLElement) HasBlur() bool {
	return js.True == bindings.HasHTMLElementBlur(
		this.Ref(),
	)
}

// BlurFunc returns the method "HTMLElement.blur".
func (this HTMLElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLElementBlurFunc(
			this.Ref(),
		),
	)
}

// Blur calls the method "HTMLElement.blur".
func (this HTMLElement) Blur() (ret js.Void) {
	bindings.CallHTMLElementBlur(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "HTMLElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementBlur(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLHeadElement() (ret HTMLHeadElement) {
	ret.ref = bindings.NewHTMLHeadElementByHTMLHeadElement()
	return
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

func NewHTMLScriptElement() (ret HTMLScriptElement) {
	ret.ref = bindings.NewHTMLScriptElementByHTMLScriptElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLScriptElement.src" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLScriptElement.type" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) NoModule() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementNoModule(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNoModule sets the value of property "HTMLScriptElement.noModule" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Async() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementAsync(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAsync sets the value of property "HTMLScriptElement.async" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Defer() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementDefer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefer sets the value of property "HTMLScriptElement.defer" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLScriptElement.crossOrigin" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLScriptElement.text" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementIntegrity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIntegrity sets the value of property "HTMLScriptElement.integrity" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLScriptElement.referrerPolicy" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementBlocking(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FetchPriority returns the value of property "HTMLScriptElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementFetchPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLScriptElement.fetchPriority" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementCharset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCharset sets the value of property "HTMLScriptElement.charset" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Event() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementEvent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEvent sets the value of property "HTMLScriptElement.event" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) HtmlFor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementHtmlFor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHtmlFor sets the value of property "HTMLScriptElement.htmlFor" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) AttributionSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementAttributionSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAttributionSrc sets the value of property "HTMLScriptElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementAttributionSrc(
		this.Ref(),
		val.Ref(),
	)
}

// HasSupports returns true if the staticmethod "HTMLScriptElement.supports" exists.
func (this HTMLScriptElement) HasSupports() bool {
	return js.True == bindings.HasHTMLScriptElementSupports(
		this.Ref(),
	)
}

// SupportsFunc returns the staticmethod "HTMLScriptElement.supports".
func (this HTMLScriptElement) SupportsFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.HTMLScriptElementSupportsFunc(
			this.Ref(),
		),
	)
}

// Supports calls the staticmethod "HTMLScriptElement.supports".
func (this HTMLScriptElement) Supports(typ js.String) (ret bool) {
	bindings.CallHTMLScriptElementSupports(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TrySupports calls the staticmethod "HTMLScriptElement.supports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLScriptElement) TrySupports(typ js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLScriptElementSupports(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGScriptElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "SVGScriptElement.type" to val.
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
// It returns ok=false if there is no such property.
func (this SVGScriptElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "SVGScriptElement.crossOrigin" to val.
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
// It returns ok=false if there is no such property.
func (this SVGScriptElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_HTMLScriptElement_SVGScriptElement struct {
	ref js.Ref
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) FromRef(ref js.Ref) OneOf_HTMLScriptElement_SVGScriptElement {
	return OneOf_HTMLScriptElement_SVGScriptElement{
		ref: ref,
	}
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) HTMLScriptElement() HTMLScriptElement {
	return HTMLScriptElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) SVGScriptElement() SVGScriptElement {
	return SVGScriptElement{}.FromRef(x.ref)
}

type HTMLOrSVGScriptElement = OneOf_HTMLScriptElement_SVGScriptElement

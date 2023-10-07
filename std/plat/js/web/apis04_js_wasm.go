// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "HTMLAllCollection.length".
//
// It returns ok=false if there is no such property.
func (this HTMLAllCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLAllCollectionLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "HTMLAllCollection." exists.
func (this HTMLAllCollection) HasFuncGet() bool {
	return js.True == bindings.HasFuncHTMLAllCollectionGet(
		this.ref,
	)
}

// FuncGet returns the method "HTMLAllCollection.".
func (this HTMLAllCollection) FuncGet() (fn js.Func[func(index uint32) Element]) {
	bindings.FuncHTMLAllCollectionGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "HTMLAllCollection.".
func (this HTMLAllCollection) Get(index uint32) (ret Element) {
	bindings.CallHTMLAllCollectionGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "HTMLAllCollection."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryGet(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "HTMLAllCollection.namedItem" exists.
func (this HTMLAllCollection) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncHTMLAllCollectionNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "HTMLAllCollection.namedItem".
func (this HTMLAllCollection) FuncNamedItem() (fn js.Func[func(name js.String) OneOf_HTMLCollection_Element]) {
	bindings.FuncHTMLAllCollectionNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "HTMLAllCollection.namedItem".
func (this HTMLAllCollection) NamedItem(name js.String) (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLAllCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryNamedItem(name js.String) (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncItem returns true if the method "HTMLAllCollection.item" exists.
func (this HTMLAllCollection) HasFuncItem() bool {
	return js.True == bindings.HasFuncHTMLAllCollectionItem(
		this.ref,
	)
}

// FuncItem returns the method "HTMLAllCollection.item".
func (this HTMLAllCollection) FuncItem() (fn js.Func[func(nameOrIndex js.String) OneOf_HTMLCollection_Element]) {
	bindings.FuncHTMLAllCollectionItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "HTMLAllCollection.item".
func (this HTMLAllCollection) Item(nameOrIndex js.String) (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionItem(
		this.ref, js.Pointer(&ret),
		nameOrIndex.Ref(),
	)

	return
}

// TryItem calls the method "HTMLAllCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryItem(nameOrIndex js.String) (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		nameOrIndex.Ref(),
	)

	return
}

// HasFuncItem1 returns true if the method "HTMLAllCollection.item" exists.
func (this HTMLAllCollection) HasFuncItem1() bool {
	return js.True == bindings.HasFuncHTMLAllCollectionItem1(
		this.ref,
	)
}

// FuncItem1 returns the method "HTMLAllCollection.item".
func (this HTMLAllCollection) FuncItem1() (fn js.Func[func() OneOf_HTMLCollection_Element]) {
	bindings.FuncHTMLAllCollectionItem1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item1 calls the method "HTMLAllCollection.item".
func (this HTMLAllCollection) Item1() (ret OneOf_HTMLCollection_Element) {
	bindings.CallHTMLAllCollectionItem1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryItem1 calls the method "HTMLAllCollection.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLAllCollection) TryItem1() (ret OneOf_HTMLCollection_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLAllCollectionItem1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DOMStringList struct {
	ref js.Ref
}

func (this DOMStringList) Once() DOMStringList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "DOMStringList.length".
//
// It returns ok=false if there is no such property.
func (this DOMStringList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDOMStringListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "DOMStringList.item" exists.
func (this DOMStringList) HasFuncItem() bool {
	return js.True == bindings.HasFuncDOMStringListItem(
		this.ref,
	)
}

// FuncItem returns the method "DOMStringList.item".
func (this DOMStringList) FuncItem() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncDOMStringListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "DOMStringList.item".
func (this DOMStringList) Item(index uint32) (ret js.String) {
	bindings.CallDOMStringListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "DOMStringList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringList) TryItem(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncContains returns true if the method "DOMStringList.contains" exists.
func (this DOMStringList) HasFuncContains() bool {
	return js.True == bindings.HasFuncDOMStringListContains(
		this.ref,
	)
}

// FuncContains returns the method "DOMStringList.contains".
func (this DOMStringList) FuncContains() (fn js.Func[func(string js.String) bool]) {
	bindings.FuncDOMStringListContains(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Contains calls the method "DOMStringList.contains".
func (this DOMStringList) Contains(string js.String) (ret bool) {
	bindings.CallDOMStringListContains(
		this.ref, js.Pointer(&ret),
		string.Ref(),
	)

	return
}

// TryContains calls the method "DOMStringList.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMStringList) TryContains(string js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMStringListContains(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		string.Ref(),
	)

	return
}

type Location struct {
	ref js.Ref
}

func (this Location) Once() Location {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "Location.href".
//
// It returns ok=false if there is no such property.
func (this Location) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "Location.href" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHref(val js.String) bool {
	return js.True == bindings.SetLocationHref(
		this.ref,
		val.Ref(),
	)
}

// Origin returns the value of property "Location.origin".
//
// It returns ok=false if there is no such property.
func (this Location) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "Location.protocol".
//
// It returns ok=false if there is no such property.
func (this Location) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "Location.protocol" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetProtocol(val js.String) bool {
	return js.True == bindings.SetLocationProtocol(
		this.ref,
		val.Ref(),
	)
}

// Host returns the value of property "Location.host".
//
// It returns ok=false if there is no such property.
func (this Location) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "Location.host" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHost(val js.String) bool {
	return js.True == bindings.SetLocationHost(
		this.ref,
		val.Ref(),
	)
}

// Hostname returns the value of property "Location.hostname".
//
// It returns ok=false if there is no such property.
func (this Location) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "Location.hostname" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHostname(val js.String) bool {
	return js.True == bindings.SetLocationHostname(
		this.ref,
		val.Ref(),
	)
}

// Port returns the value of property "Location.port".
//
// It returns ok=false if there is no such property.
func (this Location) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "Location.port" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetPort(val js.String) bool {
	return js.True == bindings.SetLocationPort(
		this.ref,
		val.Ref(),
	)
}

// Pathname returns the value of property "Location.pathname".
//
// It returns ok=false if there is no such property.
func (this Location) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "Location.pathname" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetPathname(val js.String) bool {
	return js.True == bindings.SetLocationPathname(
		this.ref,
		val.Ref(),
	)
}

// Search returns the value of property "Location.search".
//
// It returns ok=false if there is no such property.
func (this Location) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "Location.search" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetSearch(val js.String) bool {
	return js.True == bindings.SetLocationSearch(
		this.ref,
		val.Ref(),
	)
}

// Hash returns the value of property "Location.hash".
//
// It returns ok=false if there is no such property.
func (this Location) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLocationHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "Location.hash" to val.
//
// It returns false if the property cannot be set.
func (this Location) SetHash(val js.String) bool {
	return js.True == bindings.SetLocationHash(
		this.ref,
		val.Ref(),
	)
}

// AncestorOrigins returns the value of property "Location.ancestorOrigins".
//
// It returns ok=false if there is no such property.
func (this Location) AncestorOrigins() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetLocationAncestorOrigins(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAssign returns true if the method "Location.assign" exists.
func (this Location) HasFuncAssign() bool {
	return js.True == bindings.HasFuncLocationAssign(
		this.ref,
	)
}

// FuncAssign returns the method "Location.assign".
func (this Location) FuncAssign() (fn js.Func[func(url js.String)]) {
	bindings.FuncLocationAssign(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Assign calls the method "Location.assign".
func (this Location) Assign(url js.String) (ret js.Void) {
	bindings.CallLocationAssign(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryAssign calls the method "Location.assign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryAssign(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationAssign(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncReplace returns true if the method "Location.replace" exists.
func (this Location) HasFuncReplace() bool {
	return js.True == bindings.HasFuncLocationReplace(
		this.ref,
	)
}

// FuncReplace returns the method "Location.replace".
func (this Location) FuncReplace() (fn js.Func[func(url js.String)]) {
	bindings.FuncLocationReplace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Replace calls the method "Location.replace".
func (this Location) Replace(url js.String) (ret js.Void) {
	bindings.CallLocationReplace(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryReplace calls the method "Location.replace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryReplace(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationReplace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncReload returns true if the method "Location.reload" exists.
func (this Location) HasFuncReload() bool {
	return js.True == bindings.HasFuncLocationReload(
		this.ref,
	)
}

// FuncReload returns the method "Location.reload".
func (this Location) FuncReload() (fn js.Func[func()]) {
	bindings.FuncLocationReload(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reload calls the method "Location.reload".
func (this Location) Reload() (ret js.Void) {
	bindings.CallLocationReload(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReload calls the method "Location.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Location) TryReload() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLocationReload(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BlobPropertyBag) UpdateFrom(ref js.Ref) {
	bindings.BlobPropertyBagJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BlobPropertyBag) Update(ref js.Ref) {
	bindings.BlobPropertyBagJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BlobPropertyBag) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *QueuingStrategy) UpdateFrom(ref js.Ref) {
	bindings.QueuingStrategyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueuingStrategy) Update(ref js.Ref) {
	bindings.QueuingStrategyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueuingStrategy) FreeMembers(recursive bool) {
	js.Free(
		p.Size.Ref(),
	)
	p.Size = p.Size.FromRef(js.Undefined)
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
func (p *ReadableStreamReadResult) UpdateFrom(ref js.Ref) {
	bindings.ReadableStreamReadResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadableStreamReadResult) Update(ref js.Ref) {
	bindings.ReadableStreamReadResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadableStreamReadResult) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Closed returns the value of property "ReadableStreamDefaultReader.closed".
//
// It returns ok=false if there is no such property.
func (this ReadableStreamDefaultReader) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetReadableStreamDefaultReaderClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRead returns true if the method "ReadableStreamDefaultReader.read" exists.
func (this ReadableStreamDefaultReader) HasFuncRead() bool {
	return js.True == bindings.HasFuncReadableStreamDefaultReaderRead(
		this.ref,
	)
}

// FuncRead returns the method "ReadableStreamDefaultReader.read".
func (this ReadableStreamDefaultReader) FuncRead() (fn js.Func[func() js.Promise[ReadableStreamReadResult]]) {
	bindings.FuncReadableStreamDefaultReaderRead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read calls the method "ReadableStreamDefaultReader.read".
func (this ReadableStreamDefaultReader) Read() (ret js.Promise[ReadableStreamReadResult]) {
	bindings.CallReadableStreamDefaultReaderRead(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRead calls the method "ReadableStreamDefaultReader.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryRead() (ret js.Promise[ReadableStreamReadResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderRead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReleaseLock returns true if the method "ReadableStreamDefaultReader.releaseLock" exists.
func (this ReadableStreamDefaultReader) HasFuncReleaseLock() bool {
	return js.True == bindings.HasFuncReadableStreamDefaultReaderReleaseLock(
		this.ref,
	)
}

// FuncReleaseLock returns the method "ReadableStreamDefaultReader.releaseLock".
func (this ReadableStreamDefaultReader) FuncReleaseLock() (fn js.Func[func()]) {
	bindings.FuncReadableStreamDefaultReaderReleaseLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseLock calls the method "ReadableStreamDefaultReader.releaseLock".
func (this ReadableStreamDefaultReader) ReleaseLock() (ret js.Void) {
	bindings.CallReadableStreamDefaultReaderReleaseLock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "ReadableStreamDefaultReader.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderReleaseLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCancel returns true if the method "ReadableStreamDefaultReader.cancel" exists.
func (this ReadableStreamDefaultReader) HasFuncCancel() bool {
	return js.True == bindings.HasFuncReadableStreamDefaultReaderCancel(
		this.ref,
	)
}

// FuncCancel returns the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) FuncCancel() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	bindings.FuncReadableStreamDefaultReaderCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamDefaultReaderCancel(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStreamDefaultReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncCancel1 returns true if the method "ReadableStreamDefaultReader.cancel" exists.
func (this ReadableStreamDefaultReader) HasFuncCancel1() bool {
	return js.True == bindings.HasFuncReadableStreamDefaultReaderCancel1(
		this.ref,
	)
}

// FuncCancel1 returns the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) FuncCancel1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncReadableStreamDefaultReaderCancel1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel1 calls the method "ReadableStreamDefaultReader.cancel".
func (this ReadableStreamDefaultReader) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamDefaultReaderCancel1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStreamDefaultReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamDefaultReader) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamDefaultReaderCancel1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Closed returns the value of property "ReadableStreamBYOBReader.closed".
//
// It returns ok=false if there is no such property.
func (this ReadableStreamBYOBReader) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetReadableStreamBYOBReaderClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRead returns true if the method "ReadableStreamBYOBReader.read" exists.
func (this ReadableStreamBYOBReader) HasFuncRead() bool {
	return js.True == bindings.HasFuncReadableStreamBYOBReaderRead(
		this.ref,
	)
}

// FuncRead returns the method "ReadableStreamBYOBReader.read".
func (this ReadableStreamBYOBReader) FuncRead() (fn js.Func[func(view js.ArrayBufferView) js.Promise[ReadableStreamReadResult]]) {
	bindings.FuncReadableStreamBYOBReaderRead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read calls the method "ReadableStreamBYOBReader.read".
func (this ReadableStreamBYOBReader) Read(view js.ArrayBufferView) (ret js.Promise[ReadableStreamReadResult]) {
	bindings.CallReadableStreamBYOBReaderRead(
		this.ref, js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryRead calls the method "ReadableStreamBYOBReader.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryRead(view js.ArrayBufferView) (ret js.Promise[ReadableStreamReadResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderRead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

// HasFuncReleaseLock returns true if the method "ReadableStreamBYOBReader.releaseLock" exists.
func (this ReadableStreamBYOBReader) HasFuncReleaseLock() bool {
	return js.True == bindings.HasFuncReadableStreamBYOBReaderReleaseLock(
		this.ref,
	)
}

// FuncReleaseLock returns the method "ReadableStreamBYOBReader.releaseLock".
func (this ReadableStreamBYOBReader) FuncReleaseLock() (fn js.Func[func()]) {
	bindings.FuncReadableStreamBYOBReaderReleaseLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseLock calls the method "ReadableStreamBYOBReader.releaseLock".
func (this ReadableStreamBYOBReader) ReleaseLock() (ret js.Void) {
	bindings.CallReadableStreamBYOBReaderReleaseLock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "ReadableStreamBYOBReader.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderReleaseLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCancel returns true if the method "ReadableStreamBYOBReader.cancel" exists.
func (this ReadableStreamBYOBReader) HasFuncCancel() bool {
	return js.True == bindings.HasFuncReadableStreamBYOBReaderCancel(
		this.ref,
	)
}

// FuncCancel returns the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) FuncCancel() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	bindings.FuncReadableStreamBYOBReaderCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamBYOBReaderCancel(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStreamBYOBReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncCancel1 returns true if the method "ReadableStreamBYOBReader.cancel" exists.
func (this ReadableStreamBYOBReader) HasFuncCancel1() bool {
	return js.True == bindings.HasFuncReadableStreamBYOBReaderCancel1(
		this.ref,
	)
}

// FuncCancel1 returns the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) FuncCancel1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncReadableStreamBYOBReaderCancel1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel1 calls the method "ReadableStreamBYOBReader.cancel".
func (this ReadableStreamBYOBReader) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamBYOBReaderCancel1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStreamBYOBReader.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStreamBYOBReader) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamBYOBReaderCancel1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ReadableStreamGetReaderOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadableStreamGetReaderOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadableStreamGetReaderOptions) Update(ref js.Ref) {
	bindings.ReadableStreamGetReaderOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadableStreamGetReaderOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Closed returns the value of property "WritableStreamDefaultWriter.closed".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) Closed() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DesiredSize returns the value of property "WritableStreamDefaultWriter.desiredSize".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) DesiredSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterDesiredSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "WritableStreamDefaultWriter.ready".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultWriter) Ready() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultWriterReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAbort returns true if the method "WritableStreamDefaultWriter.abort" exists.
func (this WritableStreamDefaultWriter) HasFuncAbort() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterAbort(
		this.ref,
	)
}

// FuncAbort returns the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) FuncAbort() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	bindings.FuncWritableStreamDefaultWriterAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) Abort(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterAbort(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the method "WritableStreamDefaultWriter.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryAbort(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncAbort1 returns true if the method "WritableStreamDefaultWriter.abort" exists.
func (this WritableStreamDefaultWriter) HasFuncAbort1() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterAbort1(
		this.ref,
	)
}

// FuncAbort1 returns the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) FuncAbort1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWritableStreamDefaultWriterAbort1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort1 calls the method "WritableStreamDefaultWriter.abort".
func (this WritableStreamDefaultWriter) Abort1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterAbort1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the method "WritableStreamDefaultWriter.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryAbort1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterAbort1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "WritableStreamDefaultWriter.close" exists.
func (this WritableStreamDefaultWriter) HasFuncClose() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterClose(
		this.ref,
	)
}

// FuncClose returns the method "WritableStreamDefaultWriter.close".
func (this WritableStreamDefaultWriter) FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWritableStreamDefaultWriterClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "WritableStreamDefaultWriter.close".
func (this WritableStreamDefaultWriter) Close() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "WritableStreamDefaultWriter.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReleaseLock returns true if the method "WritableStreamDefaultWriter.releaseLock" exists.
func (this WritableStreamDefaultWriter) HasFuncReleaseLock() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterReleaseLock(
		this.ref,
	)
}

// FuncReleaseLock returns the method "WritableStreamDefaultWriter.releaseLock".
func (this WritableStreamDefaultWriter) FuncReleaseLock() (fn js.Func[func()]) {
	bindings.FuncWritableStreamDefaultWriterReleaseLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseLock calls the method "WritableStreamDefaultWriter.releaseLock".
func (this WritableStreamDefaultWriter) ReleaseLock() (ret js.Void) {
	bindings.CallWritableStreamDefaultWriterReleaseLock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReleaseLock calls the method "WritableStreamDefaultWriter.releaseLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryReleaseLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterReleaseLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWrite returns true if the method "WritableStreamDefaultWriter.write" exists.
func (this WritableStreamDefaultWriter) HasFuncWrite() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterWrite(
		this.ref,
	)
}

// FuncWrite returns the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) FuncWrite() (fn js.Func[func(chunk js.Any) js.Promise[js.Void]]) {
	bindings.FuncWritableStreamDefaultWriterWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) Write(chunk js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterWrite(
		this.ref, js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryWrite calls the method "WritableStreamDefaultWriter.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryWrite(chunk js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasFuncWrite1 returns true if the method "WritableStreamDefaultWriter.write" exists.
func (this WritableStreamDefaultWriter) HasFuncWrite1() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultWriterWrite1(
		this.ref,
	)
}

// FuncWrite1 returns the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) FuncWrite1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWritableStreamDefaultWriterWrite1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write1 calls the method "WritableStreamDefaultWriter.write".
func (this WritableStreamDefaultWriter) Write1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamDefaultWriterWrite1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryWrite1 calls the method "WritableStreamDefaultWriter.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultWriter) TryWrite1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultWriterWrite1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Locked returns the value of property "WritableStream.locked".
//
// It returns ok=false if there is no such property.
func (this WritableStream) Locked() (ret bool, ok bool) {
	ok = js.True == bindings.GetWritableStreamLocked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAbort returns true if the method "WritableStream.abort" exists.
func (this WritableStream) HasFuncAbort() bool {
	return js.True == bindings.HasFuncWritableStreamAbort(
		this.ref,
	)
}

// FuncAbort returns the method "WritableStream.abort".
func (this WritableStream) FuncAbort() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	bindings.FuncWritableStreamAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "WritableStream.abort".
func (this WritableStream) Abort(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamAbort(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the method "WritableStream.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryAbort(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncAbort1 returns true if the method "WritableStream.abort" exists.
func (this WritableStream) HasFuncAbort1() bool {
	return js.True == bindings.HasFuncWritableStreamAbort1(
		this.ref,
	)
}

// FuncAbort1 returns the method "WritableStream.abort".
func (this WritableStream) FuncAbort1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWritableStreamAbort1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort1 calls the method "WritableStream.abort".
func (this WritableStream) Abort1() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamAbort1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the method "WritableStream.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryAbort1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamAbort1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "WritableStream.close" exists.
func (this WritableStream) HasFuncClose() bool {
	return js.True == bindings.HasFuncWritableStreamClose(
		this.ref,
	)
}

// FuncClose returns the method "WritableStream.close".
func (this WritableStream) FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWritableStreamClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "WritableStream.close".
func (this WritableStream) Close() (ret js.Promise[js.Void]) {
	bindings.CallWritableStreamClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "WritableStream.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetWriter returns true if the method "WritableStream.getWriter" exists.
func (this WritableStream) HasFuncGetWriter() bool {
	return js.True == bindings.HasFuncWritableStreamGetWriter(
		this.ref,
	)
}

// FuncGetWriter returns the method "WritableStream.getWriter".
func (this WritableStream) FuncGetWriter() (fn js.Func[func() WritableStreamDefaultWriter]) {
	bindings.FuncWritableStreamGetWriter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetWriter calls the method "WritableStream.getWriter".
func (this WritableStream) GetWriter() (ret WritableStreamDefaultWriter) {
	bindings.CallWritableStreamGetWriter(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetWriter calls the method "WritableStream.getWriter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStream) TryGetWriter() (ret WritableStreamDefaultWriter, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamGetWriter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ReadableWritablePair) UpdateFrom(ref js.Ref) {
	bindings.ReadableWritablePairJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadableWritablePair) Update(ref js.Ref) {
	bindings.ReadableWritablePairJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadableWritablePair) FreeMembers(recursive bool) {
	js.Free(
		p.Readable.Ref(),
		p.Writable.Ref(),
	)
	p.Readable = p.Readable.FromRef(js.Undefined)
	p.Writable = p.Writable.FromRef(js.Undefined)
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
func (p *StreamPipeOptions) UpdateFrom(ref js.Ref) {
	bindings.StreamPipeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StreamPipeOptions) Update(ref js.Ref) {
	bindings.StreamPipeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StreamPipeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Locked returns the value of property "ReadableStream.locked".
//
// It returns ok=false if there is no such property.
func (this ReadableStream) Locked() (ret bool, ok bool) {
	ok = js.True == bindings.GetReadableStreamLocked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFrom returns true if the static method "ReadableStream.from" exists.
func (this ReadableStream) HasFuncFrom() bool {
	return js.True == bindings.HasFuncReadableStreamFrom(
		this.ref,
	)
}

// FuncFrom returns the static method "ReadableStream.from".
func (this ReadableStream) FuncFrom() (fn js.Func[func(asyncIterable js.Any) ReadableStream]) {
	bindings.FuncReadableStreamFrom(
		this.ref, js.Pointer(&fn),
	)
	return
}

// From calls the static method "ReadableStream.from".
func (this ReadableStream) From(asyncIterable js.Any) (ret ReadableStream) {
	bindings.CallReadableStreamFrom(
		this.ref, js.Pointer(&ret),
		asyncIterable.Ref(),
	)

	return
}

// TryFrom calls the static method "ReadableStream.from"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryFrom(asyncIterable js.Any) (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamFrom(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		asyncIterable.Ref(),
	)

	return
}

// HasFuncCancel returns true if the method "ReadableStream.cancel" exists.
func (this ReadableStream) HasFuncCancel() bool {
	return js.True == bindings.HasFuncReadableStreamCancel(
		this.ref,
	)
}

// FuncCancel returns the method "ReadableStream.cancel".
func (this ReadableStream) FuncCancel() (fn js.Func[func(reason js.Any) js.Promise[js.Void]]) {
	bindings.FuncReadableStreamCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "ReadableStream.cancel".
func (this ReadableStream) Cancel(reason js.Any) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamCancel(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryCancel calls the method "ReadableStream.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryCancel(reason js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncCancel1 returns true if the method "ReadableStream.cancel" exists.
func (this ReadableStream) HasFuncCancel1() bool {
	return js.True == bindings.HasFuncReadableStreamCancel1(
		this.ref,
	)
}

// FuncCancel1 returns the method "ReadableStream.cancel".
func (this ReadableStream) FuncCancel1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncReadableStreamCancel1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel1 calls the method "ReadableStream.cancel".
func (this ReadableStream) Cancel1() (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamCancel1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel1 calls the method "ReadableStream.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryCancel1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamCancel1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetReader returns true if the method "ReadableStream.getReader" exists.
func (this ReadableStream) HasFuncGetReader() bool {
	return js.True == bindings.HasFuncReadableStreamGetReader(
		this.ref,
	)
}

// FuncGetReader returns the method "ReadableStream.getReader".
func (this ReadableStream) FuncGetReader() (fn js.Func[func(options ReadableStreamGetReaderOptions) ReadableStreamReader]) {
	bindings.FuncReadableStreamGetReader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetReader calls the method "ReadableStream.getReader".
func (this ReadableStream) GetReader(options ReadableStreamGetReaderOptions) (ret ReadableStreamReader) {
	bindings.CallReadableStreamGetReader(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetReader calls the method "ReadableStream.getReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryGetReader(options ReadableStreamGetReaderOptions) (ret ReadableStreamReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamGetReader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetReader1 returns true if the method "ReadableStream.getReader" exists.
func (this ReadableStream) HasFuncGetReader1() bool {
	return js.True == bindings.HasFuncReadableStreamGetReader1(
		this.ref,
	)
}

// FuncGetReader1 returns the method "ReadableStream.getReader".
func (this ReadableStream) FuncGetReader1() (fn js.Func[func() ReadableStreamReader]) {
	bindings.FuncReadableStreamGetReader1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetReader1 calls the method "ReadableStream.getReader".
func (this ReadableStream) GetReader1() (ret ReadableStreamReader) {
	bindings.CallReadableStreamGetReader1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetReader1 calls the method "ReadableStream.getReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryGetReader1() (ret ReadableStreamReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamGetReader1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPipeThrough returns true if the method "ReadableStream.pipeThrough" exists.
func (this ReadableStream) HasFuncPipeThrough() bool {
	return js.True == bindings.HasFuncReadableStreamPipeThrough(
		this.ref,
	)
}

// FuncPipeThrough returns the method "ReadableStream.pipeThrough".
func (this ReadableStream) FuncPipeThrough() (fn js.Func[func(transform ReadableWritablePair, options StreamPipeOptions) ReadableStream]) {
	bindings.FuncReadableStreamPipeThrough(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PipeThrough calls the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThrough(transform ReadableWritablePair, options StreamPipeOptions) (ret ReadableStream) {
	bindings.CallReadableStreamPipeThrough(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
		js.Pointer(&options),
	)

	return
}

// HasFuncPipeThrough1 returns true if the method "ReadableStream.pipeThrough" exists.
func (this ReadableStream) HasFuncPipeThrough1() bool {
	return js.True == bindings.HasFuncReadableStreamPipeThrough1(
		this.ref,
	)
}

// FuncPipeThrough1 returns the method "ReadableStream.pipeThrough".
func (this ReadableStream) FuncPipeThrough1() (fn js.Func[func(transform ReadableWritablePair) ReadableStream]) {
	bindings.FuncReadableStreamPipeThrough1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PipeThrough1 calls the method "ReadableStream.pipeThrough".
func (this ReadableStream) PipeThrough1(transform ReadableWritablePair) (ret ReadableStream) {
	bindings.CallReadableStreamPipeThrough1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TryPipeThrough1 calls the method "ReadableStream.pipeThrough"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeThrough1(transform ReadableWritablePair) (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeThrough1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasFuncPipeTo returns true if the method "ReadableStream.pipeTo" exists.
func (this ReadableStream) HasFuncPipeTo() bool {
	return js.True == bindings.HasFuncReadableStreamPipeTo(
		this.ref,
	)
}

// FuncPipeTo returns the method "ReadableStream.pipeTo".
func (this ReadableStream) FuncPipeTo() (fn js.Func[func(destination WritableStream, options StreamPipeOptions) js.Promise[js.Void]]) {
	bindings.FuncReadableStreamPipeTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PipeTo calls the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeTo(destination WritableStream, options StreamPipeOptions) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamPipeTo(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPipeTo1 returns true if the method "ReadableStream.pipeTo" exists.
func (this ReadableStream) HasFuncPipeTo1() bool {
	return js.True == bindings.HasFuncReadableStreamPipeTo1(
		this.ref,
	)
}

// FuncPipeTo1 returns the method "ReadableStream.pipeTo".
func (this ReadableStream) FuncPipeTo1() (fn js.Func[func(destination WritableStream) js.Promise[js.Void]]) {
	bindings.FuncReadableStreamPipeTo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PipeTo1 calls the method "ReadableStream.pipeTo".
func (this ReadableStream) PipeTo1(destination WritableStream) (ret js.Promise[js.Void]) {
	bindings.CallReadableStreamPipeTo1(
		this.ref, js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryPipeTo1 calls the method "ReadableStream.pipeTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryPipeTo1(destination WritableStream) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamPipeTo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
}

// HasFuncTee returns true if the method "ReadableStream.tee" exists.
func (this ReadableStream) HasFuncTee() bool {
	return js.True == bindings.HasFuncReadableStreamTee(
		this.ref,
	)
}

// FuncTee returns the method "ReadableStream.tee".
func (this ReadableStream) FuncTee() (fn js.Func[func() js.Array[ReadableStream]]) {
	bindings.FuncReadableStreamTee(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Tee calls the method "ReadableStream.tee".
func (this ReadableStream) Tee() (ret js.Array[ReadableStream]) {
	bindings.CallReadableStreamTee(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTee calls the method "ReadableStream.tee"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ReadableStream) TryTee() (ret js.Array[ReadableStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadableStreamTee(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Size returns the value of property "Blob.size".
//
// It returns ok=false if there is no such property.
func (this Blob) Size() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBlobSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Blob.type".
//
// It returns ok=false if there is no such property.
func (this Blob) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBlobType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSlice returns true if the method "Blob.slice" exists.
func (this Blob) HasFuncSlice() bool {
	return js.True == bindings.HasFuncBlobSlice(
		this.ref,
	)
}

// FuncSlice returns the method "Blob.slice".
func (this Blob) FuncSlice() (fn js.Func[func(start int64, end int64, contentType js.String) Blob]) {
	bindings.FuncBlobSlice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Slice calls the method "Blob.slice".
func (this Blob) Slice(start int64, end int64, contentType js.String) (ret Blob) {
	bindings.CallBlobSlice(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
		contentType.Ref(),
	)

	return
}

// HasFuncSlice1 returns true if the method "Blob.slice" exists.
func (this Blob) HasFuncSlice1() bool {
	return js.True == bindings.HasFuncBlobSlice1(
		this.ref,
	)
}

// FuncSlice1 returns the method "Blob.slice".
func (this Blob) FuncSlice1() (fn js.Func[func(start int64, end int64) Blob]) {
	bindings.FuncBlobSlice1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Slice1 calls the method "Blob.slice".
func (this Blob) Slice1(start int64, end int64) (ret Blob) {
	bindings.CallBlobSlice1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
	)

	return
}

// HasFuncSlice2 returns true if the method "Blob.slice" exists.
func (this Blob) HasFuncSlice2() bool {
	return js.True == bindings.HasFuncBlobSlice2(
		this.ref,
	)
}

// FuncSlice2 returns the method "Blob.slice".
func (this Blob) FuncSlice2() (fn js.Func[func(start int64) Blob]) {
	bindings.FuncBlobSlice2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Slice2 calls the method "Blob.slice".
func (this Blob) Slice2(start int64) (ret Blob) {
	bindings.CallBlobSlice2(
		this.ref, js.Pointer(&ret),
		float64(start),
	)

	return
}

// TrySlice2 calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice2(start int64) (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
	)

	return
}

// HasFuncSlice3 returns true if the method "Blob.slice" exists.
func (this Blob) HasFuncSlice3() bool {
	return js.True == bindings.HasFuncBlobSlice3(
		this.ref,
	)
}

// FuncSlice3 returns the method "Blob.slice".
func (this Blob) FuncSlice3() (fn js.Func[func() Blob]) {
	bindings.FuncBlobSlice3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Slice3 calls the method "Blob.slice".
func (this Blob) Slice3() (ret Blob) {
	bindings.CallBlobSlice3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySlice3 calls the method "Blob.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TrySlice3() (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobSlice3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStream returns true if the method "Blob.stream" exists.
func (this Blob) HasFuncStream() bool {
	return js.True == bindings.HasFuncBlobStream(
		this.ref,
	)
}

// FuncStream returns the method "Blob.stream".
func (this Blob) FuncStream() (fn js.Func[func() ReadableStream]) {
	bindings.FuncBlobStream(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stream calls the method "Blob.stream".
func (this Blob) Stream() (ret ReadableStream) {
	bindings.CallBlobStream(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStream calls the method "Blob.stream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryStream() (ret ReadableStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobStream(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncText returns true if the method "Blob.text" exists.
func (this Blob) HasFuncText() bool {
	return js.True == bindings.HasFuncBlobText(
		this.ref,
	)
}

// FuncText returns the method "Blob.text".
func (this Blob) FuncText() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncBlobText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Text calls the method "Blob.text".
func (this Blob) Text() (ret js.Promise[js.String]) {
	bindings.CallBlobText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Blob.text"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncArrayBuffer returns true if the method "Blob.arrayBuffer" exists.
func (this Blob) HasFuncArrayBuffer() bool {
	return js.True == bindings.HasFuncBlobArrayBuffer(
		this.ref,
	)
}

// FuncArrayBuffer returns the method "Blob.arrayBuffer".
func (this Blob) FuncArrayBuffer() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	bindings.FuncBlobArrayBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArrayBuffer calls the method "Blob.arrayBuffer".
func (this Blob) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallBlobArrayBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Blob.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Blob) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBlobArrayBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *FilePropertyBag) UpdateFrom(ref js.Ref) {
	bindings.FilePropertyBagJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FilePropertyBag) Update(ref js.Ref) {
	bindings.FilePropertyBagJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FilePropertyBag) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "File.name".
//
// It returns ok=false if there is no such property.
func (this File) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastModified returns the value of property "File.lastModified".
//
// It returns ok=false if there is no such property.
func (this File) LastModified() (ret int64, ok bool) {
	ok = js.True == bindings.GetFileLastModified(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WebkitRelativePath returns the value of property "File.webkitRelativePath".
//
// It returns ok=false if there is no such property.
func (this File) WebkitRelativePath() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileWebkitRelativePath(
		this.ref, js.Pointer(&ret),
	)
	return
}

type RadioNodeList struct {
	NodeList
}

func (this RadioNodeList) Once() RadioNodeList {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "RadioNodeList.value".
//
// It returns ok=false if there is no such property.
func (this RadioNodeList) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRadioNodeListValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "RadioNodeList.value" to val.
//
// It returns false if the property cannot be set.
func (this RadioNodeList) SetValue(val js.String) bool {
	return js.True == bindings.SetRadioNodeListValue(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncNamedItem returns true if the method "HTMLFormControlsCollection.namedItem" exists.
func (this HTMLFormControlsCollection) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncHTMLFormControlsCollectionNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "HTMLFormControlsCollection.namedItem".
func (this HTMLFormControlsCollection) FuncNamedItem() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	bindings.FuncHTMLFormControlsCollectionNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "HTMLFormControlsCollection.namedItem".
func (this HTMLFormControlsCollection) NamedItem(name js.String) (ret OneOf_RadioNodeList_Element) {
	bindings.CallHTMLFormControlsCollectionNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLFormControlsCollection.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormControlsCollection) TryNamedItem(name js.String) (ret OneOf_RadioNodeList_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormControlsCollectionNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type HTMLFormElement struct {
	HTMLElement
}

func (this HTMLFormElement) Once() HTMLFormElement {
	this.ref.Once()
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
	this.ref.Free()
}

// AcceptCharset returns the value of property "HTMLFormElement.acceptCharset".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) AcceptCharset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAcceptCharset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAcceptCharset sets the value of property "HTMLFormElement.acceptCharset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAcceptCharset(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAcceptCharset(
		this.ref,
		val.Ref(),
	)
}

// Action returns the value of property "HTMLFormElement.action".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Action() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAction sets the value of property "HTMLFormElement.action" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAction(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAction(
		this.ref,
		val.Ref(),
	)
}

// Autocomplete returns the value of property "HTMLFormElement.autocomplete".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementAutocomplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLFormElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementAutocomplete(
		this.ref,
		val.Ref(),
	)
}

// Enctype returns the value of property "HTMLFormElement.enctype".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Enctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementEnctype(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEnctype sets the value of property "HTMLFormElement.enctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementEnctype(
		this.ref,
		val.Ref(),
	)
}

// Encoding returns the value of property "HTMLFormElement.encoding".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementEncoding(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEncoding sets the value of property "HTMLFormElement.encoding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetEncoding(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementEncoding(
		this.ref,
		val.Ref(),
	)
}

// Method returns the value of property "HTMLFormElement.method".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Method() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementMethod(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMethod sets the value of property "HTMLFormElement.method" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetMethod(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementMethod(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLFormElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFormElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementName(
		this.ref,
		val.Ref(),
	)
}

// NoValidate returns the value of property "HTMLFormElement.noValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) NoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementNoValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoValidate sets the value of property "HTMLFormElement.noValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLFormElementNoValidate(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Target returns the value of property "HTMLFormElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLFormElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementTarget(
		this.ref,
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLFormElement.rel".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementRel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLFormElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFormElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLFormElementRel(
		this.ref,
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLFormElement.relList".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementRelList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Elements returns the value of property "HTMLFormElement.elements".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Elements() (ret HTMLFormControlsCollection, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "HTMLFormElement.length".
//
// It returns ok=false if there is no such property.
func (this HTMLFormElement) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLFormElementLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "HTMLFormElement." exists.
func (this HTMLFormElement) HasFuncGet() bool {
	return js.True == bindings.HasFuncHTMLFormElementGet(
		this.ref,
	)
}

// FuncGet returns the method "HTMLFormElement.".
func (this HTMLFormElement) FuncGet() (fn js.Func[func(index uint32) Element]) {
	bindings.FuncHTMLFormElementGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "HTMLFormElement.".
func (this HTMLFormElement) Get(index uint32) (ret Element) {
	bindings.CallHTMLFormElementGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "HTMLFormElement."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryGet(index uint32) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGet1 returns true if the method "HTMLFormElement." exists.
func (this HTMLFormElement) HasFuncGet1() bool {
	return js.True == bindings.HasFuncHTMLFormElementGet1(
		this.ref,
	)
}

// FuncGet1 returns the method "HTMLFormElement.".
func (this HTMLFormElement) FuncGet1() (fn js.Func[func(name js.String) OneOf_RadioNodeList_Element]) {
	bindings.FuncHTMLFormElementGet1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get1 calls the method "HTMLFormElement.".
func (this HTMLFormElement) Get1(name js.String) (ret OneOf_RadioNodeList_Element) {
	bindings.CallHTMLFormElementGet1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet1 calls the method "HTMLFormElement."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryGet1(name js.String) (ret OneOf_RadioNodeList_Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementGet1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncSubmit returns true if the method "HTMLFormElement.submit" exists.
func (this HTMLFormElement) HasFuncSubmit() bool {
	return js.True == bindings.HasFuncHTMLFormElementSubmit(
		this.ref,
	)
}

// FuncSubmit returns the method "HTMLFormElement.submit".
func (this HTMLFormElement) FuncSubmit() (fn js.Func[func()]) {
	bindings.FuncHTMLFormElementSubmit(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Submit calls the method "HTMLFormElement.submit".
func (this HTMLFormElement) Submit() (ret js.Void) {
	bindings.CallHTMLFormElementSubmit(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySubmit calls the method "HTMLFormElement.submit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TrySubmit() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementSubmit(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestSubmit returns true if the method "HTMLFormElement.requestSubmit" exists.
func (this HTMLFormElement) HasFuncRequestSubmit() bool {
	return js.True == bindings.HasFuncHTMLFormElementRequestSubmit(
		this.ref,
	)
}

// FuncRequestSubmit returns the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) FuncRequestSubmit() (fn js.Func[func(submitter HTMLElement)]) {
	bindings.FuncHTMLFormElementRequestSubmit(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestSubmit calls the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmit(submitter HTMLElement) (ret js.Void) {
	bindings.CallHTMLFormElementRequestSubmit(
		this.ref, js.Pointer(&ret),
		submitter.Ref(),
	)

	return
}

// TryRequestSubmit calls the method "HTMLFormElement.requestSubmit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryRequestSubmit(submitter HTMLElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementRequestSubmit(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		submitter.Ref(),
	)

	return
}

// HasFuncRequestSubmit1 returns true if the method "HTMLFormElement.requestSubmit" exists.
func (this HTMLFormElement) HasFuncRequestSubmit1() bool {
	return js.True == bindings.HasFuncHTMLFormElementRequestSubmit1(
		this.ref,
	)
}

// FuncRequestSubmit1 returns the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) FuncRequestSubmit1() (fn js.Func[func()]) {
	bindings.FuncHTMLFormElementRequestSubmit1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestSubmit1 calls the method "HTMLFormElement.requestSubmit".
func (this HTMLFormElement) RequestSubmit1() (ret js.Void) {
	bindings.CallHTMLFormElementRequestSubmit1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestSubmit1 calls the method "HTMLFormElement.requestSubmit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryRequestSubmit1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementRequestSubmit1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "HTMLFormElement.reset" exists.
func (this HTMLFormElement) HasFuncReset() bool {
	return js.True == bindings.HasFuncHTMLFormElementReset(
		this.ref,
	)
}

// FuncReset returns the method "HTMLFormElement.reset".
func (this HTMLFormElement) FuncReset() (fn js.Func[func()]) {
	bindings.FuncHTMLFormElementReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "HTMLFormElement.reset".
func (this HTMLFormElement) Reset() (ret js.Void) {
	bindings.CallHTMLFormElementReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "HTMLFormElement.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckValidity returns true if the method "HTMLFormElement.checkValidity" exists.
func (this HTMLFormElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLFormElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLFormElement.checkValidity".
func (this HTMLFormElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLFormElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLFormElement.checkValidity".
func (this HTMLFormElement) CheckValidity() (ret bool) {
	bindings.CallHTMLFormElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLFormElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLFormElement.reportValidity" exists.
func (this HTMLFormElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLFormElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLFormElement.reportValidity".
func (this HTMLFormElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLFormElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLFormElement.reportValidity".
func (this HTMLFormElement) ReportValidity() (ret bool) {
	bindings.CallHTMLFormElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLFormElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFormElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFormElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAppend returns true if the method "FormData.append" exists.
func (this FormData) HasFuncAppend() bool {
	return js.True == bindings.HasFuncFormDataAppend(
		this.ref,
	)
}

// FuncAppend returns the method "FormData.append".
func (this FormData) FuncAppend() (fn js.Func[func(name js.String, value js.String)]) {
	bindings.FuncFormDataAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "FormData.append".
func (this FormData) Append(name js.String, value js.String) (ret js.Void) {
	bindings.CallFormDataAppend(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncAppend1 returns true if the method "FormData.append" exists.
func (this FormData) HasFuncAppend1() bool {
	return js.True == bindings.HasFuncFormDataAppend1(
		this.ref,
	)
}

// FuncAppend1 returns the method "FormData.append".
func (this FormData) FuncAppend1() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	bindings.FuncFormDataAppend1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append1 calls the method "FormData.append".
func (this FormData) Append1(name js.String, blobValue Blob, filename js.String) (ret js.Void) {
	bindings.CallFormDataAppend1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// HasFuncAppend2 returns true if the method "FormData.append" exists.
func (this FormData) HasFuncAppend2() bool {
	return js.True == bindings.HasFuncFormDataAppend2(
		this.ref,
	)
}

// FuncAppend2 returns the method "FormData.append".
func (this FormData) FuncAppend2() (fn js.Func[func(name js.String, blobValue Blob)]) {
	bindings.FuncFormDataAppend2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append2 calls the method "FormData.append".
func (this FormData) Append2(name js.String, blobValue Blob) (ret js.Void) {
	bindings.CallFormDataAppend2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "FormData.delete" exists.
func (this FormData) HasFuncDelete() bool {
	return js.True == bindings.HasFuncFormDataDelete(
		this.ref,
	)
}

// FuncDelete returns the method "FormData.delete".
func (this FormData) FuncDelete() (fn js.Func[func(name js.String)]) {
	bindings.FuncFormDataDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "FormData.delete".
func (this FormData) Delete(name js.String) (ret js.Void) {
	bindings.CallFormDataDelete(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "FormData.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryDelete(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGet returns true if the method "FormData.get" exists.
func (this FormData) HasFuncGet() bool {
	return js.True == bindings.HasFuncFormDataGet(
		this.ref,
	)
}

// FuncGet returns the method "FormData.get".
func (this FormData) FuncGet() (fn js.Func[func(name js.String) FormDataEntryValue]) {
	bindings.FuncFormDataGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "FormData.get".
func (this FormData) Get(name js.String) (ret FormDataEntryValue) {
	bindings.CallFormDataGet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "FormData.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryGet(name js.String) (ret FormDataEntryValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the method "FormData.getAll" exists.
func (this FormData) HasFuncGetAll() bool {
	return js.True == bindings.HasFuncFormDataGetAll(
		this.ref,
	)
}

// FuncGetAll returns the method "FormData.getAll".
func (this FormData) FuncGetAll() (fn js.Func[func(name js.String) js.Array[FormDataEntryValue]]) {
	bindings.FuncFormDataGetAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll calls the method "FormData.getAll".
func (this FormData) GetAll(name js.String) (ret js.Array[FormDataEntryValue]) {
	bindings.CallFormDataGetAll(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetAll calls the method "FormData.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryGetAll(name js.String) (ret js.Array[FormDataEntryValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataGetAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncHas returns true if the method "FormData.has" exists.
func (this FormData) HasFuncHas() bool {
	return js.True == bindings.HasFuncFormDataHas(
		this.ref,
	)
}

// FuncHas returns the method "FormData.has".
func (this FormData) FuncHas() (fn js.Func[func(name js.String) bool]) {
	bindings.FuncFormDataHas(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Has calls the method "FormData.has".
func (this FormData) Has(name js.String) (ret bool) {
	bindings.CallFormDataHas(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryHas calls the method "FormData.has"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FormData) TryHas(name js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFormDataHas(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncSet returns true if the method "FormData.set" exists.
func (this FormData) HasFuncSet() bool {
	return js.True == bindings.HasFuncFormDataSet(
		this.ref,
	)
}

// FuncSet returns the method "FormData.set".
func (this FormData) FuncSet() (fn js.Func[func(name js.String, value js.String)]) {
	bindings.FuncFormDataSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "FormData.set".
func (this FormData) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallFormDataSet(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncSet1 returns true if the method "FormData.set" exists.
func (this FormData) HasFuncSet1() bool {
	return js.True == bindings.HasFuncFormDataSet1(
		this.ref,
	)
}

// FuncSet1 returns the method "FormData.set".
func (this FormData) FuncSet1() (fn js.Func[func(name js.String, blobValue Blob, filename js.String)]) {
	bindings.FuncFormDataSet1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set1 calls the method "FormData.set".
func (this FormData) Set1(name js.String, blobValue Blob, filename js.String) (ret js.Void) {
	bindings.CallFormDataSet1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		blobValue.Ref(),
		filename.Ref(),
	)

	return
}

// HasFuncSet2 returns true if the method "FormData.set" exists.
func (this FormData) HasFuncSet2() bool {
	return js.True == bindings.HasFuncFormDataSet2(
		this.ref,
	)
}

// FuncSet2 returns the method "FormData.set".
func (this FormData) FuncSet2() (fn js.Func[func(name js.String, blobValue Blob)]) {
	bindings.FuncFormDataSet2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set2 calls the method "FormData.set".
func (this FormData) Set2(name js.String, blobValue Blob) (ret js.Void) {
	bindings.CallFormDataSet2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ValidityStateFlags) UpdateFrom(ref js.Ref) {
	bindings.ValidityStateFlagsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ValidityStateFlags) Update(ref js.Ref) {
	bindings.ValidityStateFlagsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ValidityStateFlags) FreeMembers(recursive bool) {
}

type ValidityState struct {
	ref js.Ref
}

func (this ValidityState) Once() ValidityState {
	this.ref.Once()
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
	this.ref.Free()
}

// ValueMissing returns the value of property "ValidityState.valueMissing".
//
// It returns ok=false if there is no such property.
func (this ValidityState) ValueMissing() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateValueMissing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TypeMismatch returns the value of property "ValidityState.typeMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TypeMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTypeMismatch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PatternMismatch returns the value of property "ValidityState.patternMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) PatternMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStatePatternMismatch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TooLong returns the value of property "ValidityState.tooLong".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TooLong() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTooLong(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TooShort returns the value of property "ValidityState.tooShort".
//
// It returns ok=false if there is no such property.
func (this ValidityState) TooShort() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateTooShort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RangeUnderflow returns the value of property "ValidityState.rangeUnderflow".
//
// It returns ok=false if there is no such property.
func (this ValidityState) RangeUnderflow() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateRangeUnderflow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RangeOverflow returns the value of property "ValidityState.rangeOverflow".
//
// It returns ok=false if there is no such property.
func (this ValidityState) RangeOverflow() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateRangeOverflow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StepMismatch returns the value of property "ValidityState.stepMismatch".
//
// It returns ok=false if there is no such property.
func (this ValidityState) StepMismatch() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateStepMismatch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BadInput returns the value of property "ValidityState.badInput".
//
// It returns ok=false if there is no such property.
func (this ValidityState) BadInput() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateBadInput(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CustomError returns the value of property "ValidityState.customError".
//
// It returns ok=false if there is no such property.
func (this ValidityState) CustomError() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateCustomError(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Valid returns the value of property "ValidityState.valid".
//
// It returns ok=false if there is no such property.
func (this ValidityState) Valid() (ret bool, ok bool) {
	ok = js.True == bindings.GetValidityStateValid(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CustomStateSet struct {
	ref js.Ref
}

func (this CustomStateSet) Once() CustomStateSet {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAdd returns true if the method "CustomStateSet.add" exists.
func (this CustomStateSet) HasFuncAdd() bool {
	return js.True == bindings.HasFuncCustomStateSetAdd(
		this.ref,
	)
}

// FuncAdd returns the method "CustomStateSet.add".
func (this CustomStateSet) FuncAdd() (fn js.Func[func(value js.String)]) {
	bindings.FuncCustomStateSetAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "CustomStateSet.add".
func (this CustomStateSet) Add(value js.String) (ret js.Void) {
	bindings.CallCustomStateSetAdd(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryAdd calls the method "CustomStateSet.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomStateSet) TryAdd(value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomStateSetAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

type ElementInternals struct {
	ref js.Ref
}

func (this ElementInternals) Once() ElementInternals {
	this.ref.Once()
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
	this.ref.Free()
}

// ShadowRoot returns the value of property "ElementInternals.shadowRoot".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) ShadowRoot() (ret ShadowRoot, ok bool) {
	ok = js.True == bindings.GetElementInternalsShadowRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Form returns the value of property "ElementInternals.form".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetElementInternalsForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "ElementInternals.willValidate".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetElementInternalsWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "ElementInternals.validity".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetElementInternalsValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "ElementInternals.validationMessage".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "ElementInternals.labels".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetElementInternalsLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// States returns the value of property "ElementInternals.states".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) States() (ret CustomStateSet, ok bool) {
	ok = js.True == bindings.GetElementInternalsStates(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Role returns the value of property "ElementInternals.role".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) Role() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsRole(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRole sets the value of property "ElementInternals.role" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetRole(val js.String) bool {
	return js.True == bindings.SetElementInternalsRole(
		this.ref,
		val.Ref(),
	)
}

// AriaActiveDescendantElement returns the value of property "ElementInternals.ariaActiveDescendantElement".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaActiveDescendantElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaActiveDescendantElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaActiveDescendantElement sets the value of property "ElementInternals.ariaActiveDescendantElement" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaActiveDescendantElement(val Element) bool {
	return js.True == bindings.SetElementInternalsAriaActiveDescendantElement(
		this.ref,
		val.Ref(),
	)
}

// AriaAtomic returns the value of property "ElementInternals.ariaAtomic".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaAtomic() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaAtomic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaAtomic sets the value of property "ElementInternals.ariaAtomic" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaAtomic(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaAtomic(
		this.ref,
		val.Ref(),
	)
}

// AriaAutoComplete returns the value of property "ElementInternals.ariaAutoComplete".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaAutoComplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaAutoComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaAutoComplete sets the value of property "ElementInternals.ariaAutoComplete" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaAutoComplete(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaAutoComplete(
		this.ref,
		val.Ref(),
	)
}

// AriaBusy returns the value of property "ElementInternals.ariaBusy".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaBusy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaBusy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaBusy sets the value of property "ElementInternals.ariaBusy" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaBusy(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaBusy(
		this.ref,
		val.Ref(),
	)
}

// AriaChecked returns the value of property "ElementInternals.ariaChecked".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaChecked() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaChecked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaChecked sets the value of property "ElementInternals.ariaChecked" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaChecked(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaChecked(
		this.ref,
		val.Ref(),
	)
}

// AriaColCount returns the value of property "ElementInternals.ariaColCount".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColCount sets the value of property "ElementInternals.ariaColCount" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColCount(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColCount(
		this.ref,
		val.Ref(),
	)
}

// AriaColIndex returns the value of property "ElementInternals.ariaColIndex".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColIndex sets the value of property "ElementInternals.ariaColIndex" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColIndex(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColIndex(
		this.ref,
		val.Ref(),
	)
}

// AriaColIndexText returns the value of property "ElementInternals.ariaColIndexText".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColIndexText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColIndexText sets the value of property "ElementInternals.ariaColIndexText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColIndexText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColIndexText(
		this.ref,
		val.Ref(),
	)
}

// AriaColSpan returns the value of property "ElementInternals.ariaColSpan".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaColSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaColSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColSpan sets the value of property "ElementInternals.ariaColSpan" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaColSpan(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaColSpan(
		this.ref,
		val.Ref(),
	)
}

// AriaControlsElements returns the value of property "ElementInternals.ariaControlsElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaControlsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaControlsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaControlsElements sets the value of property "ElementInternals.ariaControlsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaControlsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaControlsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaCurrent returns the value of property "ElementInternals.ariaCurrent".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaCurrent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaCurrent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaCurrent sets the value of property "ElementInternals.ariaCurrent" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaCurrent(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaCurrent(
		this.ref,
		val.Ref(),
	)
}

// AriaDescribedByElements returns the value of property "ElementInternals.ariaDescribedByElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDescribedByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDescribedByElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDescribedByElements sets the value of property "ElementInternals.ariaDescribedByElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDescribedByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaDescribedByElements(
		this.ref,
		val.Ref(),
	)
}

// AriaDescription returns the value of property "ElementInternals.ariaDescription".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDescription sets the value of property "ElementInternals.ariaDescription" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDescription(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaDescription(
		this.ref,
		val.Ref(),
	)
}

// AriaDetailsElements returns the value of property "ElementInternals.ariaDetailsElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDetailsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDetailsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDetailsElements sets the value of property "ElementInternals.ariaDetailsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDetailsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaDetailsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaDisabled returns the value of property "ElementInternals.ariaDisabled".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaDisabled() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaDisabled sets the value of property "ElementInternals.ariaDisabled" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaDisabled(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaDisabled(
		this.ref,
		val.Ref(),
	)
}

// AriaErrorMessageElements returns the value of property "ElementInternals.ariaErrorMessageElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaErrorMessageElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaErrorMessageElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaErrorMessageElements sets the value of property "ElementInternals.ariaErrorMessageElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaErrorMessageElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaErrorMessageElements(
		this.ref,
		val.Ref(),
	)
}

// AriaExpanded returns the value of property "ElementInternals.ariaExpanded".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaExpanded() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaExpanded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaExpanded sets the value of property "ElementInternals.ariaExpanded" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaExpanded(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaExpanded(
		this.ref,
		val.Ref(),
	)
}

// AriaFlowToElements returns the value of property "ElementInternals.ariaFlowToElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaFlowToElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaFlowToElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaFlowToElements sets the value of property "ElementInternals.ariaFlowToElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaFlowToElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaFlowToElements(
		this.ref,
		val.Ref(),
	)
}

// AriaHasPopup returns the value of property "ElementInternals.ariaHasPopup".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaHasPopup() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaHasPopup(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaHasPopup sets the value of property "ElementInternals.ariaHasPopup" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaHasPopup(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaHasPopup(
		this.ref,
		val.Ref(),
	)
}

// AriaHidden returns the value of property "ElementInternals.ariaHidden".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaHidden() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaHidden(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaHidden sets the value of property "ElementInternals.ariaHidden" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaHidden(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaHidden(
		this.ref,
		val.Ref(),
	)
}

// AriaInvalid returns the value of property "ElementInternals.ariaInvalid".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaInvalid() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaInvalid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaInvalid sets the value of property "ElementInternals.ariaInvalid" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaInvalid(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaInvalid(
		this.ref,
		val.Ref(),
	)
}

// AriaKeyShortcuts returns the value of property "ElementInternals.ariaKeyShortcuts".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaKeyShortcuts() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaKeyShortcuts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaKeyShortcuts sets the value of property "ElementInternals.ariaKeyShortcuts" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaKeyShortcuts(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaKeyShortcuts(
		this.ref,
		val.Ref(),
	)
}

// AriaLabel returns the value of property "ElementInternals.ariaLabel".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLabel sets the value of property "ElementInternals.ariaLabel" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLabel(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLabel(
		this.ref,
		val.Ref(),
	)
}

// AriaLabelledByElements returns the value of property "ElementInternals.ariaLabelledByElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLabelledByElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLabelledByElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLabelledByElements sets the value of property "ElementInternals.ariaLabelledByElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLabelledByElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaLabelledByElements(
		this.ref,
		val.Ref(),
	)
}

// AriaLevel returns the value of property "ElementInternals.ariaLevel".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLevel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLevel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLevel sets the value of property "ElementInternals.ariaLevel" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLevel(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLevel(
		this.ref,
		val.Ref(),
	)
}

// AriaLive returns the value of property "ElementInternals.ariaLive".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaLive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaLive(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaLive sets the value of property "ElementInternals.ariaLive" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaLive(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaLive(
		this.ref,
		val.Ref(),
	)
}

// AriaModal returns the value of property "ElementInternals.ariaModal".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaModal() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaModal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaModal sets the value of property "ElementInternals.ariaModal" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaModal(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaModal(
		this.ref,
		val.Ref(),
	)
}

// AriaMultiLine returns the value of property "ElementInternals.ariaMultiLine".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaMultiLine() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaMultiLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaMultiLine sets the value of property "ElementInternals.ariaMultiLine" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaMultiLine(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaMultiLine(
		this.ref,
		val.Ref(),
	)
}

// AriaMultiSelectable returns the value of property "ElementInternals.ariaMultiSelectable".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaMultiSelectable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaMultiSelectable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaMultiSelectable sets the value of property "ElementInternals.ariaMultiSelectable" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaMultiSelectable(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaMultiSelectable(
		this.ref,
		val.Ref(),
	)
}

// AriaOrientation returns the value of property "ElementInternals.ariaOrientation".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaOrientation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaOrientation sets the value of property "ElementInternals.ariaOrientation" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaOrientation(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaOrientation(
		this.ref,
		val.Ref(),
	)
}

// AriaOwnsElements returns the value of property "ElementInternals.ariaOwnsElements".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaOwnsElements() (ret js.FrozenArray[Element], ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaOwnsElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaOwnsElements sets the value of property "ElementInternals.ariaOwnsElements" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaOwnsElements(val js.FrozenArray[Element]) bool {
	return js.True == bindings.SetElementInternalsAriaOwnsElements(
		this.ref,
		val.Ref(),
	)
}

// AriaPlaceholder returns the value of property "ElementInternals.ariaPlaceholder".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPlaceholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPlaceholder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPlaceholder sets the value of property "ElementInternals.ariaPlaceholder" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPlaceholder(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPlaceholder(
		this.ref,
		val.Ref(),
	)
}

// AriaPosInSet returns the value of property "ElementInternals.ariaPosInSet".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPosInSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPosInSet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPosInSet sets the value of property "ElementInternals.ariaPosInSet" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPosInSet(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPosInSet(
		this.ref,
		val.Ref(),
	)
}

// AriaPressed returns the value of property "ElementInternals.ariaPressed".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaPressed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaPressed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaPressed sets the value of property "ElementInternals.ariaPressed" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaPressed(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaPressed(
		this.ref,
		val.Ref(),
	)
}

// AriaReadOnly returns the value of property "ElementInternals.ariaReadOnly".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaReadOnly() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaReadOnly(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaReadOnly sets the value of property "ElementInternals.ariaReadOnly" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaReadOnly(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaReadOnly(
		this.ref,
		val.Ref(),
	)
}

// AriaRequired returns the value of property "ElementInternals.ariaRequired".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRequired() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRequired(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRequired sets the value of property "ElementInternals.ariaRequired" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRequired(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRequired(
		this.ref,
		val.Ref(),
	)
}

// AriaRoleDescription returns the value of property "ElementInternals.ariaRoleDescription".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRoleDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRoleDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRoleDescription sets the value of property "ElementInternals.ariaRoleDescription" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRoleDescription(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRoleDescription(
		this.ref,
		val.Ref(),
	)
}

// AriaRowCount returns the value of property "ElementInternals.ariaRowCount".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowCount() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowCount sets the value of property "ElementInternals.ariaRowCount" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowCount(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowCount(
		this.ref,
		val.Ref(),
	)
}

// AriaRowIndex returns the value of property "ElementInternals.ariaRowIndex".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowIndex() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndex sets the value of property "ElementInternals.ariaRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowIndex(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowIndex(
		this.ref,
		val.Ref(),
	)
}

// AriaRowIndexText returns the value of property "ElementInternals.ariaRowIndexText".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowIndexText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowIndexText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowIndexText sets the value of property "ElementInternals.ariaRowIndexText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowIndexText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowIndexText(
		this.ref,
		val.Ref(),
	)
}

// AriaRowSpan returns the value of property "ElementInternals.ariaRowSpan".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaRowSpan() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaRowSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowSpan sets the value of property "ElementInternals.ariaRowSpan" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaRowSpan(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaRowSpan(
		this.ref,
		val.Ref(),
	)
}

// AriaSelected returns the value of property "ElementInternals.ariaSelected".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSelected() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSelected sets the value of property "ElementInternals.ariaSelected" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSelected(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSelected(
		this.ref,
		val.Ref(),
	)
}

// AriaSetSize returns the value of property "ElementInternals.ariaSetSize".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSetSize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSetSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSetSize sets the value of property "ElementInternals.ariaSetSize" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSetSize(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSetSize(
		this.ref,
		val.Ref(),
	)
}

// AriaSort returns the value of property "ElementInternals.ariaSort".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaSort() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaSort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaSort sets the value of property "ElementInternals.ariaSort" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaSort(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaSort(
		this.ref,
		val.Ref(),
	)
}

// AriaValueMax returns the value of property "ElementInternals.ariaValueMax".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueMax() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueMax sets the value of property "ElementInternals.ariaValueMax" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueMax(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueMax(
		this.ref,
		val.Ref(),
	)
}

// AriaValueMin returns the value of property "ElementInternals.ariaValueMin".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueMin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueMin sets the value of property "ElementInternals.ariaValueMin" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueMin(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueMin(
		this.ref,
		val.Ref(),
	)
}

// AriaValueNow returns the value of property "ElementInternals.ariaValueNow".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueNow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueNow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueNow sets the value of property "ElementInternals.ariaValueNow" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueNow(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueNow(
		this.ref,
		val.Ref(),
	)
}

// AriaValueText returns the value of property "ElementInternals.ariaValueText".
//
// It returns ok=false if there is no such property.
func (this ElementInternals) AriaValueText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetElementInternalsAriaValueText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaValueText sets the value of property "ElementInternals.ariaValueText" to val.
//
// It returns false if the property cannot be set.
func (this ElementInternals) SetAriaValueText(val js.String) bool {
	return js.True == bindings.SetElementInternalsAriaValueText(
		this.ref,
		val.Ref(),
	)
}

// HasFuncSetFormValue returns true if the method "ElementInternals.setFormValue" exists.
func (this ElementInternals) HasFuncSetFormValue() bool {
	return js.True == bindings.HasFuncElementInternalsSetFormValue(
		this.ref,
	)
}

// FuncSetFormValue returns the method "ElementInternals.setFormValue".
func (this ElementInternals) FuncSetFormValue() (fn js.Func[func(value OneOf_File_String_FormData, state OneOf_File_String_FormData)]) {
	bindings.FuncElementInternalsSetFormValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetFormValue calls the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValue(value OneOf_File_String_FormData, state OneOf_File_String_FormData) (ret js.Void) {
	bindings.CallElementInternalsSetFormValue(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		state.Ref(),
	)

	return
}

// HasFuncSetFormValue1 returns true if the method "ElementInternals.setFormValue" exists.
func (this ElementInternals) HasFuncSetFormValue1() bool {
	return js.True == bindings.HasFuncElementInternalsSetFormValue1(
		this.ref,
	)
}

// FuncSetFormValue1 returns the method "ElementInternals.setFormValue".
func (this ElementInternals) FuncSetFormValue1() (fn js.Func[func(value OneOf_File_String_FormData)]) {
	bindings.FuncElementInternalsSetFormValue1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetFormValue1 calls the method "ElementInternals.setFormValue".
func (this ElementInternals) SetFormValue1(value OneOf_File_String_FormData) (ret js.Void) {
	bindings.CallElementInternalsSetFormValue1(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TrySetFormValue1 calls the method "ElementInternals.setFormValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetFormValue1(value OneOf_File_String_FormData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetFormValue1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncSetValidity returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasFuncSetValidity() bool {
	return js.True == bindings.HasFuncElementInternalsSetValidity(
		this.ref,
	)
}

// FuncSetValidity returns the method "ElementInternals.setValidity".
func (this ElementInternals) FuncSetValidity() (fn js.Func[func(flags ValidityStateFlags, message js.String, anchor HTMLElement)]) {
	bindings.FuncElementInternalsSetValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValidity calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity(flags ValidityStateFlags, message js.String, anchor HTMLElement) (ret js.Void) {
	bindings.CallElementInternalsSetValidity(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
		message.Ref(),
		anchor.Ref(),
	)

	return
}

// HasFuncSetValidity1 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasFuncSetValidity1() bool {
	return js.True == bindings.HasFuncElementInternalsSetValidity1(
		this.ref,
	)
}

// FuncSetValidity1 returns the method "ElementInternals.setValidity".
func (this ElementInternals) FuncSetValidity1() (fn js.Func[func(flags ValidityStateFlags, message js.String)]) {
	bindings.FuncElementInternalsSetValidity1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValidity1 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity1(flags ValidityStateFlags, message js.String) (ret js.Void) {
	bindings.CallElementInternalsSetValidity1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
		message.Ref(),
	)

	return
}

// HasFuncSetValidity2 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasFuncSetValidity2() bool {
	return js.True == bindings.HasFuncElementInternalsSetValidity2(
		this.ref,
	)
}

// FuncSetValidity2 returns the method "ElementInternals.setValidity".
func (this ElementInternals) FuncSetValidity2() (fn js.Func[func(flags ValidityStateFlags)]) {
	bindings.FuncElementInternalsSetValidity2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValidity2 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity2(flags ValidityStateFlags) (ret js.Void) {
	bindings.CallElementInternalsSetValidity2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&flags),
	)

	return
}

// TrySetValidity2 calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity2(flags ValidityStateFlags) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
	)

	return
}

// HasFuncSetValidity3 returns true if the method "ElementInternals.setValidity" exists.
func (this ElementInternals) HasFuncSetValidity3() bool {
	return js.True == bindings.HasFuncElementInternalsSetValidity3(
		this.ref,
	)
}

// FuncSetValidity3 returns the method "ElementInternals.setValidity".
func (this ElementInternals) FuncSetValidity3() (fn js.Func[func()]) {
	bindings.FuncElementInternalsSetValidity3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValidity3 calls the method "ElementInternals.setValidity".
func (this ElementInternals) SetValidity3() (ret js.Void) {
	bindings.CallElementInternalsSetValidity3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetValidity3 calls the method "ElementInternals.setValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TrySetValidity3() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsSetValidity3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckValidity returns true if the method "ElementInternals.checkValidity" exists.
func (this ElementInternals) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncElementInternalsCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "ElementInternals.checkValidity".
func (this ElementInternals) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncElementInternalsCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "ElementInternals.checkValidity".
func (this ElementInternals) CheckValidity() (ret bool) {
	bindings.CallElementInternalsCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "ElementInternals.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "ElementInternals.reportValidity" exists.
func (this ElementInternals) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncElementInternalsReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "ElementInternals.reportValidity".
func (this ElementInternals) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncElementInternalsReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "ElementInternals.reportValidity".
func (this ElementInternals) ReportValidity() (ret bool) {
	bindings.CallElementInternalsReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "ElementInternals.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ElementInternals) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryElementInternalsReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *EditContextInit) UpdateFrom(ref js.Ref) {
	bindings.EditContextInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EditContextInit) Update(ref js.Ref) {
	bindings.EditContextInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EditContextInit) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Text returns the value of property "EditContext.text".
//
// It returns ok=false if there is no such property.
func (this EditContext) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEditContextText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "EditContext.selectionStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectionEnd returns the value of property "EditContext.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CompositionRangeStart returns the value of property "EditContext.compositionRangeStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) CompositionRangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCompositionRangeStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CompositionRangeEnd returns the value of property "EditContext.compositionRangeEnd".
//
// It returns ok=false if there is no such property.
func (this EditContext) CompositionRangeEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCompositionRangeEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsComposing returns the value of property "EditContext.isComposing".
//
// It returns ok=false if there is no such property.
func (this EditContext) IsComposing() (ret bool, ok bool) {
	ok = js.True == bindings.GetEditContextIsComposing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ControlBounds returns the value of property "EditContext.controlBounds".
//
// It returns ok=false if there is no such property.
func (this EditContext) ControlBounds() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetEditContextControlBounds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectionBounds returns the value of property "EditContext.selectionBounds".
//
// It returns ok=false if there is no such property.
func (this EditContext) SelectionBounds() (ret DOMRect, ok bool) {
	ok = js.True == bindings.GetEditContextSelectionBounds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CharacterBoundsRangeStart returns the value of property "EditContext.characterBoundsRangeStart".
//
// It returns ok=false if there is no such property.
func (this EditContext) CharacterBoundsRangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEditContextCharacterBoundsRangeStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncUpdateText returns true if the method "EditContext.updateText" exists.
func (this EditContext) HasFuncUpdateText() bool {
	return js.True == bindings.HasFuncEditContextUpdateText(
		this.ref,
	)
}

// FuncUpdateText returns the method "EditContext.updateText".
func (this EditContext) FuncUpdateText() (fn js.Func[func(rangeStart uint32, rangeEnd uint32, text js.String)]) {
	bindings.FuncEditContextUpdateText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateText calls the method "EditContext.updateText".
func (this EditContext) UpdateText(rangeStart uint32, rangeEnd uint32, text js.String) (ret js.Void) {
	bindings.CallEditContextUpdateText(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(rangeStart),
		uint32(rangeEnd),
		text.Ref(),
	)

	return
}

// HasFuncUpdateSelection returns true if the method "EditContext.updateSelection" exists.
func (this EditContext) HasFuncUpdateSelection() bool {
	return js.True == bindings.HasFuncEditContextUpdateSelection(
		this.ref,
	)
}

// FuncUpdateSelection returns the method "EditContext.updateSelection".
func (this EditContext) FuncUpdateSelection() (fn js.Func[func(start uint32, end uint32)]) {
	bindings.FuncEditContextUpdateSelection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateSelection calls the method "EditContext.updateSelection".
func (this EditContext) UpdateSelection(start uint32, end uint32) (ret js.Void) {
	bindings.CallEditContextUpdateSelection(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

// HasFuncUpdateControlBounds returns true if the method "EditContext.updateControlBounds" exists.
func (this EditContext) HasFuncUpdateControlBounds() bool {
	return js.True == bindings.HasFuncEditContextUpdateControlBounds(
		this.ref,
	)
}

// FuncUpdateControlBounds returns the method "EditContext.updateControlBounds".
func (this EditContext) FuncUpdateControlBounds() (fn js.Func[func(controlBounds DOMRect)]) {
	bindings.FuncEditContextUpdateControlBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateControlBounds calls the method "EditContext.updateControlBounds".
func (this EditContext) UpdateControlBounds(controlBounds DOMRect) (ret js.Void) {
	bindings.CallEditContextUpdateControlBounds(
		this.ref, js.Pointer(&ret),
		controlBounds.Ref(),
	)

	return
}

// TryUpdateControlBounds calls the method "EditContext.updateControlBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateControlBounds(controlBounds DOMRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateControlBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		controlBounds.Ref(),
	)

	return
}

// HasFuncUpdateSelectionBounds returns true if the method "EditContext.updateSelectionBounds" exists.
func (this EditContext) HasFuncUpdateSelectionBounds() bool {
	return js.True == bindings.HasFuncEditContextUpdateSelectionBounds(
		this.ref,
	)
}

// FuncUpdateSelectionBounds returns the method "EditContext.updateSelectionBounds".
func (this EditContext) FuncUpdateSelectionBounds() (fn js.Func[func(selectionBounds DOMRect)]) {
	bindings.FuncEditContextUpdateSelectionBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateSelectionBounds calls the method "EditContext.updateSelectionBounds".
func (this EditContext) UpdateSelectionBounds(selectionBounds DOMRect) (ret js.Void) {
	bindings.CallEditContextUpdateSelectionBounds(
		this.ref, js.Pointer(&ret),
		selectionBounds.Ref(),
	)

	return
}

// TryUpdateSelectionBounds calls the method "EditContext.updateSelectionBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryUpdateSelectionBounds(selectionBounds DOMRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextUpdateSelectionBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectionBounds.Ref(),
	)

	return
}

// HasFuncUpdateCharacterBounds returns true if the method "EditContext.updateCharacterBounds" exists.
func (this EditContext) HasFuncUpdateCharacterBounds() bool {
	return js.True == bindings.HasFuncEditContextUpdateCharacterBounds(
		this.ref,
	)
}

// FuncUpdateCharacterBounds returns the method "EditContext.updateCharacterBounds".
func (this EditContext) FuncUpdateCharacterBounds() (fn js.Func[func(rangeStart uint32, characterBounds js.Array[DOMRect])]) {
	bindings.FuncEditContextUpdateCharacterBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateCharacterBounds calls the method "EditContext.updateCharacterBounds".
func (this EditContext) UpdateCharacterBounds(rangeStart uint32, characterBounds js.Array[DOMRect]) (ret js.Void) {
	bindings.CallEditContextUpdateCharacterBounds(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(rangeStart),
		characterBounds.Ref(),
	)

	return
}

// HasFuncAttachedElements returns true if the method "EditContext.attachedElements" exists.
func (this EditContext) HasFuncAttachedElements() bool {
	return js.True == bindings.HasFuncEditContextAttachedElements(
		this.ref,
	)
}

// FuncAttachedElements returns the method "EditContext.attachedElements".
func (this EditContext) FuncAttachedElements() (fn js.Func[func() js.Array[Element]]) {
	bindings.FuncEditContextAttachedElements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AttachedElements calls the method "EditContext.attachedElements".
func (this EditContext) AttachedElements() (ret js.Array[Element]) {
	bindings.CallEditContextAttachedElements(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAttachedElements calls the method "EditContext.attachedElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryAttachedElements() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextAttachedElements(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCharacterBounds returns true if the method "EditContext.characterBounds" exists.
func (this EditContext) HasFuncCharacterBounds() bool {
	return js.True == bindings.HasFuncEditContextCharacterBounds(
		this.ref,
	)
}

// FuncCharacterBounds returns the method "EditContext.characterBounds".
func (this EditContext) FuncCharacterBounds() (fn js.Func[func() js.Array[DOMRect]]) {
	bindings.FuncEditContextCharacterBounds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CharacterBounds calls the method "EditContext.characterBounds".
func (this EditContext) CharacterBounds() (ret js.Array[DOMRect]) {
	bindings.CallEditContextCharacterBounds(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCharacterBounds calls the method "EditContext.characterBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EditContext) TryCharacterBounds() (ret js.Array[DOMRect], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEditContextCharacterBounds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLElement struct {
	Element
}

func (this HTMLElement) Once() HTMLElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Title returns the value of property "HTMLElement.title".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementTitle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "HTMLElement.title" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTitle(val js.String) bool {
	return js.True == bindings.SetHTMLElementTitle(
		this.ref,
		val.Ref(),
	)
}

// Lang returns the value of property "HTMLElement.lang".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementLang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLang sets the value of property "HTMLElement.lang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetLang(val js.String) bool {
	return js.True == bindings.SetHTMLElementLang(
		this.ref,
		val.Ref(),
	)
}

// Translate returns the value of property "HTMLElement.translate".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Translate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementTranslate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTranslate sets the value of property "HTMLElement.translate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTranslate(val bool) bool {
	return js.True == bindings.SetHTMLElementTranslate(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Dir returns the value of property "HTMLElement.dir".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Dir() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementDir(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDir sets the value of property "HTMLElement.dir" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetDir(val js.String) bool {
	return js.True == bindings.SetHTMLElementDir(
		this.ref,
		val.Ref(),
	)
}

// Hidden returns the value of property "HTMLElement.hidden".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Hidden() (ret OneOf_Bool_Float64_String, ok bool) {
	ok = js.True == bindings.GetHTMLElementHidden(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHidden sets the value of property "HTMLElement.hidden" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetHidden(val OneOf_Bool_Float64_String) bool {
	return js.True == bindings.SetHTMLElementHidden(
		this.ref,
		val.Ref(),
	)
}

// Inert returns the value of property "HTMLElement.inert".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Inert() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementInert(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInert sets the value of property "HTMLElement.inert" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInert(val bool) bool {
	return js.True == bindings.SetHTMLElementInert(
		this.ref,
		js.Bool(bool(val)),
	)
}

// AccessKey returns the value of property "HTMLElement.accessKey".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) AccessKey() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAccessKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAccessKey sets the value of property "HTMLElement.accessKey" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAccessKey(val js.String) bool {
	return js.True == bindings.SetHTMLElementAccessKey(
		this.ref,
		val.Ref(),
	)
}

// AccessKeyLabel returns the value of property "HTMLElement.accessKeyLabel".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) AccessKeyLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAccessKeyLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Draggable returns the value of property "HTMLElement.draggable".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Draggable() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementDraggable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDraggable sets the value of property "HTMLElement.draggable" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetDraggable(val bool) bool {
	return js.True == bindings.SetHTMLElementDraggable(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Spellcheck returns the value of property "HTMLElement.spellcheck".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Spellcheck() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementSpellcheck(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpellcheck sets the value of property "HTMLElement.spellcheck" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetSpellcheck(val bool) bool {
	return js.True == bindings.SetHTMLElementSpellcheck(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Autocapitalize returns the value of property "HTMLElement.autocapitalize".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Autocapitalize() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementAutocapitalize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutocapitalize sets the value of property "HTMLElement.autocapitalize" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAutocapitalize(val js.String) bool {
	return js.True == bindings.SetHTMLElementAutocapitalize(
		this.ref,
		val.Ref(),
	)
}

// InnerText returns the value of property "HTMLElement.innerText".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) InnerText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementInnerText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInnerText sets the value of property "HTMLElement.innerText" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInnerText(val js.String) bool {
	return js.True == bindings.SetHTMLElementInnerText(
		this.ref,
		val.Ref(),
	)
}

// OuterText returns the value of property "HTMLElement.outerText".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OuterText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementOuterText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOuterText sets the value of property "HTMLElement.outerText" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetOuterText(val js.String) bool {
	return js.True == bindings.SetHTMLElementOuterText(
		this.ref,
		val.Ref(),
	)
}

// Popover returns the value of property "HTMLElement.popover".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Popover() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementPopover(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPopover sets the value of property "HTMLElement.popover" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetPopover(val js.String) bool {
	return js.True == bindings.SetHTMLElementPopover(
		this.ref,
		val.Ref(),
	)
}

// OffsetParent returns the value of property "HTMLElement.offsetParent".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetParent() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetTop returns the value of property "HTMLElement.offsetTop".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetLeft returns the value of property "HTMLElement.offsetLeft".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetWidth returns the value of property "HTMLElement.offsetWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetHeight returns the value of property "HTMLElement.offsetHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) OffsetHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementOffsetHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EditContext returns the value of property "HTMLElement.editContext".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) EditContext() (ret EditContext, ok bool) {
	ok = js.True == bindings.GetHTMLElementEditContext(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEditContext sets the value of property "HTMLElement.editContext" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetEditContext(val EditContext) bool {
	return js.True == bindings.SetHTMLElementEditContext(
		this.ref,
		val.Ref(),
	)
}

// Style returns the value of property "HTMLElement.style".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetHTMLElementStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "HTMLElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetHTMLElementAttributeStyleMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "HTMLElement.dataset".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetHTMLElementDataset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "HTMLElement.nonce".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementNonce(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "HTMLElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetHTMLElementNonce(
		this.ref,
		val.Ref(),
	)
}

// Autofocus returns the value of property "HTMLElement.autofocus".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementAutofocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "HTMLElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetHTMLElementAutofocus(
		this.ref,
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "HTMLElement.tabIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLElementTabIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "HTMLElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetHTMLElementTabIndex(
		this.ref,
		int32(val),
	)
}

// ContentEditable returns the value of property "HTMLElement.contentEditable".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) ContentEditable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementContentEditable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContentEditable sets the value of property "HTMLElement.contentEditable" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetContentEditable(val js.String) bool {
	return js.True == bindings.SetHTMLElementContentEditable(
		this.ref,
		val.Ref(),
	)
}

// EnterKeyHint returns the value of property "HTMLElement.enterKeyHint".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) EnterKeyHint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementEnterKeyHint(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEnterKeyHint sets the value of property "HTMLElement.enterKeyHint" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetEnterKeyHint(val js.String) bool {
	return js.True == bindings.SetHTMLElementEnterKeyHint(
		this.ref,
		val.Ref(),
	)
}

// IsContentEditable returns the value of property "HTMLElement.isContentEditable".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) IsContentEditable() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLElementIsContentEditable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputMode returns the value of property "HTMLElement.inputMode".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) InputMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementInputMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInputMode sets the value of property "HTMLElement.inputMode" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetInputMode(val js.String) bool {
	return js.True == bindings.SetHTMLElementInputMode(
		this.ref,
		val.Ref(),
	)
}

// VirtualKeyboardPolicy returns the value of property "HTMLElement.virtualKeyboardPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLElement) VirtualKeyboardPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLElementVirtualKeyboardPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVirtualKeyboardPolicy sets the value of property "HTMLElement.virtualKeyboardPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLElement) SetVirtualKeyboardPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLElementVirtualKeyboardPolicy(
		this.ref,
		val.Ref(),
	)
}

// HasFuncClick returns true if the method "HTMLElement.click" exists.
func (this HTMLElement) HasFuncClick() bool {
	return js.True == bindings.HasFuncHTMLElementClick(
		this.ref,
	)
}

// FuncClick returns the method "HTMLElement.click".
func (this HTMLElement) FuncClick() (fn js.Func[func()]) {
	bindings.FuncHTMLElementClick(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Click calls the method "HTMLElement.click".
func (this HTMLElement) Click() (ret js.Void) {
	bindings.CallHTMLElementClick(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClick calls the method "HTMLElement.click"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryClick() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementClick(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAttachInternals returns true if the method "HTMLElement.attachInternals" exists.
func (this HTMLElement) HasFuncAttachInternals() bool {
	return js.True == bindings.HasFuncHTMLElementAttachInternals(
		this.ref,
	)
}

// FuncAttachInternals returns the method "HTMLElement.attachInternals".
func (this HTMLElement) FuncAttachInternals() (fn js.Func[func() ElementInternals]) {
	bindings.FuncHTMLElementAttachInternals(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AttachInternals calls the method "HTMLElement.attachInternals".
func (this HTMLElement) AttachInternals() (ret ElementInternals) {
	bindings.CallHTMLElementAttachInternals(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAttachInternals calls the method "HTMLElement.attachInternals"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryAttachInternals() (ret ElementInternals, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementAttachInternals(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowPopover returns true if the method "HTMLElement.showPopover" exists.
func (this HTMLElement) HasFuncShowPopover() bool {
	return js.True == bindings.HasFuncHTMLElementShowPopover(
		this.ref,
	)
}

// FuncShowPopover returns the method "HTMLElement.showPopover".
func (this HTMLElement) FuncShowPopover() (fn js.Func[func()]) {
	bindings.FuncHTMLElementShowPopover(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowPopover calls the method "HTMLElement.showPopover".
func (this HTMLElement) ShowPopover() (ret js.Void) {
	bindings.CallHTMLElementShowPopover(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowPopover calls the method "HTMLElement.showPopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryShowPopover() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementShowPopover(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHidePopover returns true if the method "HTMLElement.hidePopover" exists.
func (this HTMLElement) HasFuncHidePopover() bool {
	return js.True == bindings.HasFuncHTMLElementHidePopover(
		this.ref,
	)
}

// FuncHidePopover returns the method "HTMLElement.hidePopover".
func (this HTMLElement) FuncHidePopover() (fn js.Func[func()]) {
	bindings.FuncHTMLElementHidePopover(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HidePopover calls the method "HTMLElement.hidePopover".
func (this HTMLElement) HidePopover() (ret js.Void) {
	bindings.CallHTMLElementHidePopover(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHidePopover calls the method "HTMLElement.hidePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryHidePopover() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementHidePopover(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTogglePopover returns true if the method "HTMLElement.togglePopover" exists.
func (this HTMLElement) HasFuncTogglePopover() bool {
	return js.True == bindings.HasFuncHTMLElementTogglePopover(
		this.ref,
	)
}

// FuncTogglePopover returns the method "HTMLElement.togglePopover".
func (this HTMLElement) FuncTogglePopover() (fn js.Func[func(force bool) bool]) {
	bindings.FuncHTMLElementTogglePopover(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TogglePopover calls the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopover(force bool) (ret bool) {
	bindings.CallHTMLElementTogglePopover(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(force)),
	)

	return
}

// TryTogglePopover calls the method "HTMLElement.togglePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryTogglePopover(force bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementTogglePopover(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(force)),
	)

	return
}

// HasFuncTogglePopover1 returns true if the method "HTMLElement.togglePopover" exists.
func (this HTMLElement) HasFuncTogglePopover1() bool {
	return js.True == bindings.HasFuncHTMLElementTogglePopover1(
		this.ref,
	)
}

// FuncTogglePopover1 returns the method "HTMLElement.togglePopover".
func (this HTMLElement) FuncTogglePopover1() (fn js.Func[func() bool]) {
	bindings.FuncHTMLElementTogglePopover1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TogglePopover1 calls the method "HTMLElement.togglePopover".
func (this HTMLElement) TogglePopover1() (ret bool) {
	bindings.CallHTMLElementTogglePopover1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTogglePopover1 calls the method "HTMLElement.togglePopover"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryTogglePopover1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementTogglePopover1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFocus returns true if the method "HTMLElement.focus" exists.
func (this HTMLElement) HasFuncFocus() bool {
	return js.True == bindings.HasFuncHTMLElementFocus(
		this.ref,
	)
}

// FuncFocus returns the method "HTMLElement.focus".
func (this HTMLElement) FuncFocus() (fn js.Func[func(options FocusOptions)]) {
	bindings.FuncHTMLElementFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "HTMLElement.focus".
func (this HTMLElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallHTMLElementFocus(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "HTMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncFocus1 returns true if the method "HTMLElement.focus" exists.
func (this HTMLElement) HasFuncFocus1() bool {
	return js.True == bindings.HasFuncHTMLElementFocus1(
		this.ref,
	)
}

// FuncFocus1 returns the method "HTMLElement.focus".
func (this HTMLElement) FuncFocus1() (fn js.Func[func()]) {
	bindings.FuncHTMLElementFocus1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus1 calls the method "HTMLElement.focus".
func (this HTMLElement) Focus1() (ret js.Void) {
	bindings.CallHTMLElementFocus1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "HTMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementFocus1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlur returns true if the method "HTMLElement.blur" exists.
func (this HTMLElement) HasFuncBlur() bool {
	return js.True == bindings.HasFuncHTMLElementBlur(
		this.ref,
	)
}

// FuncBlur returns the method "HTMLElement.blur".
func (this HTMLElement) FuncBlur() (fn js.Func[func()]) {
	bindings.FuncHTMLElementBlur(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blur calls the method "HTMLElement.blur".
func (this HTMLElement) Blur() (ret js.Void) {
	bindings.CallHTMLElementBlur(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "HTMLElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLElementBlur(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLHeadElement struct {
	HTMLElement
}

func (this HTMLHeadElement) Once() HTMLHeadElement {
	this.ref.Once()
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
	this.ref.Free()
}

type HTMLScriptElement struct {
	HTMLElement
}

func (this HTMLScriptElement) Once() HTMLScriptElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "HTMLScriptElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLScriptElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLScriptElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLScriptElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementType(
		this.ref,
		val.Ref(),
	)
}

// NoModule returns the value of property "HTMLScriptElement.noModule".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) NoModule() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementNoModule(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoModule sets the value of property "HTMLScriptElement.noModule" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetNoModule(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementNoModule(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Async returns the value of property "HTMLScriptElement.async".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Async() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementAsync(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAsync sets the value of property "HTMLScriptElement.async" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetAsync(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementAsync(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Defer returns the value of property "HTMLScriptElement.defer".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Defer() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementDefer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefer sets the value of property "HTMLScriptElement.defer" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetDefer(val bool) bool {
	return js.True == bindings.SetHTMLScriptElementDefer(
		this.ref,
		js.Bool(bool(val)),
	)
}

// CrossOrigin returns the value of property "HTMLScriptElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLScriptElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementCrossOrigin(
		this.ref,
		val.Ref(),
	)
}

// Text returns the value of property "HTMLScriptElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLScriptElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementText(
		this.ref,
		val.Ref(),
	)
}

// Integrity returns the value of property "HTMLScriptElement.integrity".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementIntegrity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIntegrity sets the value of property "HTMLScriptElement.integrity" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetIntegrity(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementIntegrity(
		this.ref,
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLScriptElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLScriptElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLScriptElement.blocking".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementBlocking(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FetchPriority returns the value of property "HTMLScriptElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementFetchPriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLScriptElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementFetchPriority(
		this.ref,
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLScriptElement.charset".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementCharset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCharset sets the value of property "HTMLScriptElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementCharset(
		this.ref,
		val.Ref(),
	)
}

// Event returns the value of property "HTMLScriptElement.event".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) Event() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementEvent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEvent sets the value of property "HTMLScriptElement.event" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetEvent(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementEvent(
		this.ref,
		val.Ref(),
	)
}

// HtmlFor returns the value of property "HTMLScriptElement.htmlFor".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) HtmlFor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementHtmlFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHtmlFor sets the value of property "HTMLScriptElement.htmlFor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetHtmlFor(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementHtmlFor(
		this.ref,
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLScriptElement.attributionSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLScriptElement) AttributionSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLScriptElementAttributionSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAttributionSrc sets the value of property "HTMLScriptElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLScriptElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLScriptElementAttributionSrc(
		this.ref,
		val.Ref(),
	)
}

// HasFuncSupports returns true if the static method "HTMLScriptElement.supports" exists.
func (this HTMLScriptElement) HasFuncSupports() bool {
	return js.True == bindings.HasFuncHTMLScriptElementSupports(
		this.ref,
	)
}

// FuncSupports returns the static method "HTMLScriptElement.supports".
func (this HTMLScriptElement) FuncSupports() (fn js.Func[func(typ js.String) bool]) {
	bindings.FuncHTMLScriptElementSupports(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Supports calls the static method "HTMLScriptElement.supports".
func (this HTMLScriptElement) Supports(typ js.String) (ret bool) {
	bindings.CallHTMLScriptElementSupports(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TrySupports calls the static method "HTMLScriptElement.supports"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLScriptElement) TrySupports(typ js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLScriptElementSupports(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

type SVGScriptElement struct {
	SVGElement
}

func (this SVGScriptElement) Once() SVGScriptElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "SVGScriptElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGScriptElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "SVGScriptElement.type" to val.
//
// It returns false if the property cannot be set.
func (this SVGScriptElement) SetType(val js.String) bool {
	return js.True == bindings.SetSVGScriptElementType(
		this.ref,
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "SVGScriptElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this SVGScriptElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "SVGScriptElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this SVGScriptElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetSVGScriptElementCrossOrigin(
		this.ref,
		val.Ref(),
	)
}

// Href returns the value of property "SVGScriptElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGScriptElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGScriptElementHref(
		this.ref, js.Pointer(&ret),
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

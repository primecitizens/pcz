// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package readinglist

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/readinglist/bindings"
)

type AddEntryCallbackFunc func(this js.Ref) js.Ref

func (fn AddEntryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddEntryCallbackFunc) DispatchCallback(
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

type AddEntryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *AddEntryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddEntryCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AddEntryOptions struct {
	// Url is "AddEntryOptions.url"
	//
	// Optional
	Url js.String
	// Title is "AddEntryOptions.title"
	//
	// Optional
	Title js.String
	// HasBeenRead is "AddEntryOptions.hasBeenRead"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasBeenRead MUST be set to true to make this field effective.
	HasBeenRead bool

	FFI_USE_HasBeenRead bool // for HasBeenRead.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddEntryOptions with all fields set.
func (p AddEntryOptions) FromRef(ref js.Ref) AddEntryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddEntryOptions in the application heap.
func (p AddEntryOptions) New() js.Ref {
	return bindings.AddEntryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddEntryOptions) UpdateFrom(ref js.Ref) {
	bindings.AddEntryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddEntryOptions) Update(ref js.Ref) {
	bindings.AddEntryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddEntryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Title.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type QueryCallbackFunc func(this js.Ref, entries js.Array[ReadingListEntry]) js.Ref

func (fn QueryCallbackFunc) Register() js.Func[func(entries js.Array[ReadingListEntry])] {
	return js.RegisterCallback[func(entries js.Array[ReadingListEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn QueryCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ReadingListEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QueryCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[ReadingListEntry]) js.Ref
	Arg T
}

func (cb *QueryCallback[T]) Register() js.Func[func(entries js.Array[ReadingListEntry])] {
	return js.RegisterCallback[func(entries js.Array[ReadingListEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *QueryCallback[T]) DispatchCallback(
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

		js.Array[ReadingListEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReadingListEntry struct {
	// Url is "ReadingListEntry.url"
	//
	// Optional
	Url js.String
	// Title is "ReadingListEntry.title"
	//
	// Optional
	Title js.String
	// HasBeenRead is "ReadingListEntry.hasBeenRead"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasBeenRead MUST be set to true to make this field effective.
	HasBeenRead bool
	// LastUpdateTime is "ReadingListEntry.lastUpdateTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastUpdateTime MUST be set to true to make this field effective.
	LastUpdateTime float64
	// CreationTime is "ReadingListEntry.creationTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CreationTime MUST be set to true to make this field effective.
	CreationTime float64

	FFI_USE_HasBeenRead    bool // for HasBeenRead.
	FFI_USE_LastUpdateTime bool // for LastUpdateTime.
	FFI_USE_CreationTime   bool // for CreationTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadingListEntry with all fields set.
func (p ReadingListEntry) FromRef(ref js.Ref) ReadingListEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadingListEntry in the application heap.
func (p ReadingListEntry) New() js.Ref {
	return bindings.ReadingListEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadingListEntry) UpdateFrom(ref js.Ref) {
	bindings.ReadingListEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadingListEntry) Update(ref js.Ref) {
	bindings.ReadingListEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadingListEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Title.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type QueryInfo struct {
	// Url is "QueryInfo.url"
	//
	// Optional
	Url js.String
	// Title is "QueryInfo.title"
	//
	// Optional
	Title js.String
	// HasBeenRead is "QueryInfo.hasBeenRead"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasBeenRead MUST be set to true to make this field effective.
	HasBeenRead bool

	FFI_USE_HasBeenRead bool // for HasBeenRead.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryInfo with all fields set.
func (p QueryInfo) FromRef(ref js.Ref) QueryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryInfo in the application heap.
func (p QueryInfo) New() js.Ref {
	return bindings.QueryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryInfo) UpdateFrom(ref js.Ref) {
	bindings.QueryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryInfo) Update(ref js.Ref) {
	bindings.QueryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Title.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type RemoveEntryCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveEntryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveEntryCallbackFunc) DispatchCallback(
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

type RemoveEntryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveEntryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveEntryCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemoveOptions struct {
	// Url is "RemoveOptions.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveOptions with all fields set.
func (p RemoveOptions) FromRef(ref js.Ref) RemoveOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveOptions in the application heap.
func (p RemoveOptions) New() js.Ref {
	return bindings.RemoveOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveOptions) UpdateFrom(ref js.Ref) {
	bindings.RemoveOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveOptions) Update(ref js.Ref) {
	bindings.RemoveOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type UpdateEntryCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateEntryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateEntryCallbackFunc) DispatchCallback(
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

type UpdateEntryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateEntryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateEntryCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateEntryOptions struct {
	// Url is "UpdateEntryOptions.url"
	//
	// Optional
	Url js.String
	// Title is "UpdateEntryOptions.title"
	//
	// Optional
	Title js.String
	// HasBeenRead is "UpdateEntryOptions.hasBeenRead"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasBeenRead MUST be set to true to make this field effective.
	HasBeenRead bool

	FFI_USE_HasBeenRead bool // for HasBeenRead.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateEntryOptions with all fields set.
func (p UpdateEntryOptions) FromRef(ref js.Ref) UpdateEntryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateEntryOptions in the application heap.
func (p UpdateEntryOptions) New() js.Ref {
	return bindings.UpdateEntryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateEntryOptions) UpdateFrom(ref js.Ref) {
	bindings.UpdateEntryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateEntryOptions) Update(ref js.Ref) {
	bindings.UpdateEntryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateEntryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Title.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

// HasFuncAddEntry returns true if the function "WEBEXT.readingList.addEntry" exists.
func HasFuncAddEntry() bool {
	return js.True == bindings.HasFuncAddEntry()
}

// FuncAddEntry returns the function "WEBEXT.readingList.addEntry".
func FuncAddEntry() (fn js.Func[func(entry AddEntryOptions) js.Promise[js.Void]]) {
	bindings.FuncAddEntry(
		js.Pointer(&fn),
	)
	return
}

// AddEntry calls the function "WEBEXT.readingList.addEntry" directly.
func AddEntry(entry AddEntryOptions) (ret js.Promise[js.Void]) {
	bindings.CallAddEntry(
		js.Pointer(&ret),
		js.Pointer(&entry),
	)

	return
}

// TryAddEntry calls the function "WEBEXT.readingList.addEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddEntry(entry AddEntryOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&entry),
	)

	return
}

type OnEntryAddedEventCallbackFunc func(this js.Ref, entry *ReadingListEntry) js.Ref

func (fn OnEntryAddedEventCallbackFunc) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEntryAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEntryAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry *ReadingListEntry) js.Ref
	Arg T
}

func (cb *OnEntryAddedEventCallback[T]) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEntryAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEntryAdded returns true if the function "WEBEXT.readingList.onEntryAdded.addListener" exists.
func HasFuncOnEntryAdded() bool {
	return js.True == bindings.HasFuncOnEntryAdded()
}

// FuncOnEntryAdded returns the function "WEBEXT.readingList.onEntryAdded.addListener".
func FuncOnEntryAdded() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOnEntryAdded(
		js.Pointer(&fn),
	)
	return
}

// OnEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.addListener" directly.
func OnEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOnEntryAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEntryAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEntryAdded returns true if the function "WEBEXT.readingList.onEntryAdded.removeListener" exists.
func HasFuncOffEntryAdded() bool {
	return js.True == bindings.HasFuncOffEntryAdded()
}

// FuncOffEntryAdded returns the function "WEBEXT.readingList.onEntryAdded.removeListener".
func FuncOffEntryAdded() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOffEntryAdded(
		js.Pointer(&fn),
	)
	return
}

// OffEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.removeListener" directly.
func OffEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOffEntryAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEntryAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEntryAdded returns true if the function "WEBEXT.readingList.onEntryAdded.hasListener" exists.
func HasFuncHasOnEntryAdded() bool {
	return js.True == bindings.HasFuncHasOnEntryAdded()
}

// FuncHasOnEntryAdded returns the function "WEBEXT.readingList.onEntryAdded.hasListener".
func FuncHasOnEntryAdded() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)]) bool]) {
	bindings.FuncHasOnEntryAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.hasListener" directly.
func HasOnEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret bool) {
	bindings.CallHasOnEntryAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEntryAdded calls the function "WEBEXT.readingList.onEntryAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEntryAdded(callback js.Func[func(entry *ReadingListEntry)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEntryAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEntryRemovedEventCallbackFunc func(this js.Ref, entry *ReadingListEntry) js.Ref

func (fn OnEntryRemovedEventCallbackFunc) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEntryRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEntryRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry *ReadingListEntry) js.Ref
	Arg T
}

func (cb *OnEntryRemovedEventCallback[T]) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEntryRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEntryRemoved returns true if the function "WEBEXT.readingList.onEntryRemoved.addListener" exists.
func HasFuncOnEntryRemoved() bool {
	return js.True == bindings.HasFuncOnEntryRemoved()
}

// FuncOnEntryRemoved returns the function "WEBEXT.readingList.onEntryRemoved.addListener".
func FuncOnEntryRemoved() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOnEntryRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.addListener" directly.
func OnEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOnEntryRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEntryRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEntryRemoved returns true if the function "WEBEXT.readingList.onEntryRemoved.removeListener" exists.
func HasFuncOffEntryRemoved() bool {
	return js.True == bindings.HasFuncOffEntryRemoved()
}

// FuncOffEntryRemoved returns the function "WEBEXT.readingList.onEntryRemoved.removeListener".
func FuncOffEntryRemoved() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOffEntryRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.removeListener" directly.
func OffEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOffEntryRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEntryRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEntryRemoved returns true if the function "WEBEXT.readingList.onEntryRemoved.hasListener" exists.
func HasFuncHasOnEntryRemoved() bool {
	return js.True == bindings.HasFuncHasOnEntryRemoved()
}

// FuncHasOnEntryRemoved returns the function "WEBEXT.readingList.onEntryRemoved.hasListener".
func FuncHasOnEntryRemoved() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)]) bool]) {
	bindings.FuncHasOnEntryRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.hasListener" directly.
func HasOnEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret bool) {
	bindings.CallHasOnEntryRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEntryRemoved calls the function "WEBEXT.readingList.onEntryRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEntryRemoved(callback js.Func[func(entry *ReadingListEntry)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEntryRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEntryUpdatedEventCallbackFunc func(this js.Ref, entry *ReadingListEntry) js.Ref

func (fn OnEntryUpdatedEventCallbackFunc) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEntryUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEntryUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry *ReadingListEntry) js.Ref
	Arg T
}

func (cb *OnEntryUpdatedEventCallback[T]) Register() js.Func[func(entry *ReadingListEntry)] {
	return js.RegisterCallback[func(entry *ReadingListEntry)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEntryUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadingListEntry
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEntryUpdated returns true if the function "WEBEXT.readingList.onEntryUpdated.addListener" exists.
func HasFuncOnEntryUpdated() bool {
	return js.True == bindings.HasFuncOnEntryUpdated()
}

// FuncOnEntryUpdated returns the function "WEBEXT.readingList.onEntryUpdated.addListener".
func FuncOnEntryUpdated() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOnEntryUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.addListener" directly.
func OnEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOnEntryUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEntryUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEntryUpdated returns true if the function "WEBEXT.readingList.onEntryUpdated.removeListener" exists.
func HasFuncOffEntryUpdated() bool {
	return js.True == bindings.HasFuncOffEntryUpdated()
}

// FuncOffEntryUpdated returns the function "WEBEXT.readingList.onEntryUpdated.removeListener".
func FuncOffEntryUpdated() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)])]) {
	bindings.FuncOffEntryUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.removeListener" directly.
func OffEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void) {
	bindings.CallOffEntryUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEntryUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEntryUpdated returns true if the function "WEBEXT.readingList.onEntryUpdated.hasListener" exists.
func HasFuncHasOnEntryUpdated() bool {
	return js.True == bindings.HasFuncHasOnEntryUpdated()
}

// FuncHasOnEntryUpdated returns the function "WEBEXT.readingList.onEntryUpdated.hasListener".
func FuncHasOnEntryUpdated() (fn js.Func[func(callback js.Func[func(entry *ReadingListEntry)]) bool]) {
	bindings.FuncHasOnEntryUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.hasListener" directly.
func HasOnEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret bool) {
	bindings.CallHasOnEntryUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEntryUpdated calls the function "WEBEXT.readingList.onEntryUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEntryUpdated(callback js.Func[func(entry *ReadingListEntry)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEntryUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncQuery returns true if the function "WEBEXT.readingList.query" exists.
func HasFuncQuery() bool {
	return js.True == bindings.HasFuncQuery()
}

// FuncQuery returns the function "WEBEXT.readingList.query".
func FuncQuery() (fn js.Func[func(info QueryInfo) js.Promise[js.Array[ReadingListEntry]]]) {
	bindings.FuncQuery(
		js.Pointer(&fn),
	)
	return
}

// Query calls the function "WEBEXT.readingList.query" directly.
func Query(info QueryInfo) (ret js.Promise[js.Array[ReadingListEntry]]) {
	bindings.CallQuery(
		js.Pointer(&ret),
		js.Pointer(&info),
	)

	return
}

// TryQuery calls the function "WEBEXT.readingList.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQuery(info QueryInfo) (ret js.Promise[js.Array[ReadingListEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQuery(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&info),
	)

	return
}

// HasFuncRemoveEntry returns true if the function "WEBEXT.readingList.removeEntry" exists.
func HasFuncRemoveEntry() bool {
	return js.True == bindings.HasFuncRemoveEntry()
}

// FuncRemoveEntry returns the function "WEBEXT.readingList.removeEntry".
func FuncRemoveEntry() (fn js.Func[func(info RemoveOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveEntry(
		js.Pointer(&fn),
	)
	return
}

// RemoveEntry calls the function "WEBEXT.readingList.removeEntry" directly.
func RemoveEntry(info RemoveOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveEntry(
		js.Pointer(&ret),
		js.Pointer(&info),
	)

	return
}

// TryRemoveEntry calls the function "WEBEXT.readingList.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveEntry(info RemoveOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&info),
	)

	return
}

// HasFuncUpdateEntry returns true if the function "WEBEXT.readingList.updateEntry" exists.
func HasFuncUpdateEntry() bool {
	return js.True == bindings.HasFuncUpdateEntry()
}

// FuncUpdateEntry returns the function "WEBEXT.readingList.updateEntry".
func FuncUpdateEntry() (fn js.Func[func(info UpdateEntryOptions) js.Promise[js.Void]]) {
	bindings.FuncUpdateEntry(
		js.Pointer(&fn),
	)
	return
}

// UpdateEntry calls the function "WEBEXT.readingList.updateEntry" directly.
func UpdateEntry(info UpdateEntryOptions) (ret js.Promise[js.Void]) {
	bindings.CallUpdateEntry(
		js.Pointer(&ret),
		js.Pointer(&info),
	)

	return
}

// TryUpdateEntry calls the function "WEBEXT.readingList.updateEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateEntry(info UpdateEntryOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&info),
	)

	return
}

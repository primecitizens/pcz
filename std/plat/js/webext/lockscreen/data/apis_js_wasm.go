// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package data

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/lockscreen/data/bindings"
)

type DataCallbackFunc func(this js.Ref, data js.ArrayBuffer) js.Ref

func (fn DataCallbackFunc) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *DataCallback[T]) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DataCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DataItemCallbackFunc func(this js.Ref, item *DataItemInfo) js.Ref

func (fn DataItemCallbackFunc) Register() js.Func[func(item *DataItemInfo)] {
	return js.RegisterCallback[func(item *DataItemInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DataItemCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DataItemInfo
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

type DataItemCallback[T any] struct {
	Fn  func(arg T, this js.Ref, item *DataItemInfo) js.Ref
	Arg T
}

func (cb *DataItemCallback[T]) Register() js.Func[func(item *DataItemInfo)] {
	return js.RegisterCallback[func(item *DataItemInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DataItemCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DataItemInfo
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

type DataItemInfo struct {
	// Id is "DataItemInfo.id"
	//
	// Optional
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DataItemInfo with all fields set.
func (p DataItemInfo) FromRef(ref js.Ref) DataItemInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DataItemInfo in the application heap.
func (p DataItemInfo) New() js.Ref {
	return bindings.DataItemInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DataItemInfo) UpdateFrom(ref js.Ref) {
	bindings.DataItemInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DataItemInfo) Update(ref js.Ref) {
	bindings.DataItemInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DataItemInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type DataItemListCallbackFunc func(this js.Ref, items js.Array[DataItemInfo]) js.Ref

func (fn DataItemListCallbackFunc) Register() js.Func[func(items js.Array[DataItemInfo])] {
	return js.RegisterCallback[func(items js.Array[DataItemInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DataItemListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DataItemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DataItemListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, items js.Array[DataItemInfo]) js.Ref
	Arg T
}

func (cb *DataItemListCallback[T]) Register() js.Func[func(items js.Array[DataItemInfo])] {
	return js.RegisterCallback[func(items js.Array[DataItemInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DataItemListCallback[T]) DispatchCallback(
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

		js.Array[DataItemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DataItemsAvailableEvent struct {
	// WasLocked is "DataItemsAvailableEvent.wasLocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_WasLocked MUST be set to true to make this field effective.
	WasLocked bool

	FFI_USE_WasLocked bool // for WasLocked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DataItemsAvailableEvent with all fields set.
func (p DataItemsAvailableEvent) FromRef(ref js.Ref) DataItemsAvailableEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DataItemsAvailableEvent in the application heap.
func (p DataItemsAvailableEvent) New() js.Ref {
	return bindings.DataItemsAvailableEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DataItemsAvailableEvent) UpdateFrom(ref js.Ref) {
	bindings.DataItemsAvailableEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DataItemsAvailableEvent) Update(ref js.Ref) {
	bindings.DataItemsAvailableEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DataItemsAvailableEvent) FreeMembers(recursive bool) {
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncCreate returns true if the function "WEBEXT.lockScreen.data.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.lockScreen.data.create".
func FuncCreate() (fn js.Func[func() js.Promise[DataItemInfo]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.lockScreen.data.create" directly.
func Create() (ret js.Promise[DataItemInfo]) {
	bindings.CallCreate(
		js.Pointer(&ret),
	)

	return
}

// TryCreate calls the function "WEBEXT.lockScreen.data.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate() (ret js.Promise[DataItemInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDelete returns true if the function "WEBEXT.lockScreen.data.delete" exists.
func HasFuncDelete() bool {
	return js.True == bindings.HasFuncDelete()
}

// FuncDelete returns the function "WEBEXT.lockScreen.data.delete".
func FuncDelete() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncDelete(
		js.Pointer(&fn),
	)
	return
}

// Delete calls the function "WEBEXT.lockScreen.data.delete" directly.
func Delete(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallDelete(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryDelete calls the function "WEBEXT.lockScreen.data.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDelete(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDelete(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.lockScreen.data.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.lockScreen.data.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Array[DataItemInfo]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.lockScreen.data.getAll" directly.
func GetAll() (ret js.Promise[js.Array[DataItemInfo]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.lockScreen.data.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Array[DataItemInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetContent returns true if the function "WEBEXT.lockScreen.data.getContent" exists.
func HasFuncGetContent() bool {
	return js.True == bindings.HasFuncGetContent()
}

// FuncGetContent returns the function "WEBEXT.lockScreen.data.getContent".
func FuncGetContent() (fn js.Func[func(id js.String) js.Promise[js.ArrayBuffer]]) {
	bindings.FuncGetContent(
		js.Pointer(&fn),
	)
	return
}

// GetContent calls the function "WEBEXT.lockScreen.data.getContent" directly.
func GetContent(id js.String) (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallGetContent(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetContent calls the function "WEBEXT.lockScreen.data.getContent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContent(id js.String) (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContent(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type OnDataItemsAvailableEventCallbackFunc func(this js.Ref, event *DataItemsAvailableEvent) js.Ref

func (fn OnDataItemsAvailableEventCallbackFunc) Register() js.Func[func(event *DataItemsAvailableEvent)] {
	return js.RegisterCallback[func(event *DataItemsAvailableEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDataItemsAvailableEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DataItemsAvailableEvent
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

type OnDataItemsAvailableEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *DataItemsAvailableEvent) js.Ref
	Arg T
}

func (cb *OnDataItemsAvailableEventCallback[T]) Register() js.Func[func(event *DataItemsAvailableEvent)] {
	return js.RegisterCallback[func(event *DataItemsAvailableEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDataItemsAvailableEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DataItemsAvailableEvent
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

// HasFuncOnDataItemsAvailable returns true if the function "WEBEXT.lockScreen.data.onDataItemsAvailable.addListener" exists.
func HasFuncOnDataItemsAvailable() bool {
	return js.True == bindings.HasFuncOnDataItemsAvailable()
}

// FuncOnDataItemsAvailable returns the function "WEBEXT.lockScreen.data.onDataItemsAvailable.addListener".
func FuncOnDataItemsAvailable() (fn js.Func[func(callback js.Func[func(event *DataItemsAvailableEvent)])]) {
	bindings.FuncOnDataItemsAvailable(
		js.Pointer(&fn),
	)
	return
}

// OnDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.addListener" directly.
func OnDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret js.Void) {
	bindings.CallOnDataItemsAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDataItemsAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDataItemsAvailable returns true if the function "WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener" exists.
func HasFuncOffDataItemsAvailable() bool {
	return js.True == bindings.HasFuncOffDataItemsAvailable()
}

// FuncOffDataItemsAvailable returns the function "WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener".
func FuncOffDataItemsAvailable() (fn js.Func[func(callback js.Func[func(event *DataItemsAvailableEvent)])]) {
	bindings.FuncOffDataItemsAvailable(
		js.Pointer(&fn),
	)
	return
}

// OffDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener" directly.
func OffDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret js.Void) {
	bindings.CallOffDataItemsAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDataItemsAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDataItemsAvailable returns true if the function "WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener" exists.
func HasFuncHasOnDataItemsAvailable() bool {
	return js.True == bindings.HasFuncHasOnDataItemsAvailable()
}

// FuncHasOnDataItemsAvailable returns the function "WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener".
func FuncHasOnDataItemsAvailable() (fn js.Func[func(callback js.Func[func(event *DataItemsAvailableEvent)]) bool]) {
	bindings.FuncHasOnDataItemsAvailable(
		js.Pointer(&fn),
	)
	return
}

// HasOnDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener" directly.
func HasOnDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret bool) {
	bindings.CallHasOnDataItemsAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDataItemsAvailable calls the function "WEBEXT.lockScreen.data.onDataItemsAvailable.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDataItemsAvailable(callback js.Func[func(event *DataItemsAvailableEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDataItemsAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetContent returns true if the function "WEBEXT.lockScreen.data.setContent" exists.
func HasFuncSetContent() bool {
	return js.True == bindings.HasFuncSetContent()
}

// FuncSetContent returns the function "WEBEXT.lockScreen.data.setContent".
func FuncSetContent() (fn js.Func[func(id js.String, data js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncSetContent(
		js.Pointer(&fn),
	)
	return
}

// SetContent calls the function "WEBEXT.lockScreen.data.setContent" directly.
func SetContent(id js.String, data js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallSetContent(
		js.Pointer(&ret),
		id.Ref(),
		data.Ref(),
	)

	return
}

// TrySetContent calls the function "WEBEXT.lockScreen.data.setContent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetContent(id js.String, data js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetContent(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		data.Ref(),
	)

	return
}

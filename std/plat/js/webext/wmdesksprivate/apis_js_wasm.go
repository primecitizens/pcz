// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package wmdesksprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/wmdesksprivate/bindings"
)

type Desk struct {
	// DeskUuid is "Desk.deskUuid"
	//
	// Optional
	DeskUuid js.String
	// DeskName is "Desk.deskName"
	//
	// Optional
	DeskName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Desk with all fields set.
func (p Desk) FromRef(ref js.Ref) Desk {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Desk in the application heap.
func (p Desk) New() js.Ref {
	return bindings.DeskJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Desk) UpdateFrom(ref js.Ref) {
	bindings.DeskJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Desk) Update(ref js.Ref) {
	bindings.DeskJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Desk) FreeMembers(recursive bool) {
	js.Free(
		p.DeskUuid.Ref(),
		p.DeskName.Ref(),
	)
	p.DeskUuid = p.DeskUuid.FromRef(js.Undefined)
	p.DeskName = p.DeskName.FromRef(js.Undefined)
}

type DeskIdCallbackFunc func(this js.Ref, deskId js.String) js.Ref

func (fn DeskIdCallbackFunc) Register() js.Func[func(deskId js.String)] {
	return js.RegisterCallback[func(deskId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DeskIdCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DeskIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deskId js.String) js.Ref
	Arg T
}

func (cb *DeskIdCallback[T]) Register() js.Func[func(deskId js.String)] {
	return js.RegisterCallback[func(deskId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DeskIdCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllDesksCallbackFunc func(this js.Ref, desks js.Array[Desk]) js.Ref

func (fn GetAllDesksCallbackFunc) Register() js.Func[func(desks js.Array[Desk])] {
	return js.RegisterCallback[func(desks js.Array[Desk])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllDesksCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Desk]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllDesksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, desks js.Array[Desk]) js.Ref
	Arg T
}

func (cb *GetAllDesksCallback[T]) Register() js.Func[func(desks js.Array[Desk])] {
	return js.RegisterCallback[func(desks js.Array[Desk])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllDesksCallback[T]) DispatchCallback(
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

		js.Array[Desk]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeskByIDCallbackFunc func(this js.Ref, desk *Desk) js.Ref

func (fn GetDeskByIDCallbackFunc) Register() js.Func[func(desk *Desk)] {
	return js.RegisterCallback[func(desk *Desk)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeskByIDCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Desk
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

type GetDeskByIDCallback[T any] struct {
	Fn  func(arg T, this js.Ref, desk *Desk) js.Ref
	Arg T
}

func (cb *GetDeskByIDCallback[T]) Register() js.Func[func(desk *Desk)] {
	return js.RegisterCallback[func(desk *Desk)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeskByIDCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Desk
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

type GetDeskTemplateJsonCallbackFunc func(this js.Ref, templateJson js.String) js.Ref

func (fn GetDeskTemplateJsonCallbackFunc) Register() js.Func[func(templateJson js.String)] {
	return js.RegisterCallback[func(templateJson js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeskTemplateJsonCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeskTemplateJsonCallback[T any] struct {
	Fn  func(arg T, this js.Ref, templateJson js.String) js.Ref
	Arg T
}

func (cb *GetDeskTemplateJsonCallback[T]) Register() js.Func[func(templateJson js.String)] {
	return js.RegisterCallback[func(templateJson js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeskTemplateJsonCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSavedDesksCallbackFunc func(this js.Ref, saveDesks js.Array[SavedDesk]) js.Ref

func (fn GetSavedDesksCallbackFunc) Register() js.Func[func(saveDesks js.Array[SavedDesk])] {
	return js.RegisterCallback[func(saveDesks js.Array[SavedDesk])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSavedDesksCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SavedDesk]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSavedDesksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, saveDesks js.Array[SavedDesk]) js.Ref
	Arg T
}

func (cb *GetSavedDesksCallback[T]) Register() js.Func[func(saveDesks js.Array[SavedDesk])] {
	return js.RegisterCallback[func(saveDesks js.Array[SavedDesk])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSavedDesksCallback[T]) DispatchCallback(
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

		js.Array[SavedDesk]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SavedDeskType uint32

const (
	_ SavedDeskType = iota

	SavedDeskType_K_TEMPLATE
	SavedDeskType_K_SAVE_AND_RECALL
	SavedDeskType_K_UNKNOWN
)

func (SavedDeskType) FromRef(str js.Ref) SavedDeskType {
	return SavedDeskType(bindings.ConstOfSavedDeskType(str))
}

func (x SavedDeskType) String() (string, bool) {
	switch x {
	case SavedDeskType_K_TEMPLATE:
		return "kTemplate", true
	case SavedDeskType_K_SAVE_AND_RECALL:
		return "kSaveAndRecall", true
	case SavedDeskType_K_UNKNOWN:
		return "kUnknown", true
	default:
		return "", false
	}
}

type SavedDesk struct {
	// SavedDeskUuid is "SavedDesk.savedDeskUuid"
	//
	// Optional
	SavedDeskUuid js.String
	// SavedDeskName is "SavedDesk.savedDeskName"
	//
	// Optional
	SavedDeskName js.String
	// SavedDeskType is "SavedDesk.savedDeskType"
	//
	// Optional
	SavedDeskType SavedDeskType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SavedDesk with all fields set.
func (p SavedDesk) FromRef(ref js.Ref) SavedDesk {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SavedDesk in the application heap.
func (p SavedDesk) New() js.Ref {
	return bindings.SavedDeskJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SavedDesk) UpdateFrom(ref js.Ref) {
	bindings.SavedDeskJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SavedDesk) Update(ref js.Ref) {
	bindings.SavedDeskJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SavedDesk) FreeMembers(recursive bool) {
	js.Free(
		p.SavedDeskUuid.Ref(),
		p.SavedDeskName.Ref(),
	)
	p.SavedDeskUuid = p.SavedDeskUuid.FromRef(js.Undefined)
	p.SavedDeskName = p.SavedDeskName.FromRef(js.Undefined)
}

type LaunchOptions struct {
	// DeskName is "LaunchOptions.deskName"
	//
	// Optional
	DeskName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LaunchOptions with all fields set.
func (p LaunchOptions) FromRef(ref js.Ref) LaunchOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LaunchOptions in the application heap.
func (p LaunchOptions) New() js.Ref {
	return bindings.LaunchOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LaunchOptions) UpdateFrom(ref js.Ref) {
	bindings.LaunchOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LaunchOptions) Update(ref js.Ref) {
	bindings.LaunchOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LaunchOptions) FreeMembers(recursive bool) {
	js.Free(
		p.DeskName.Ref(),
	)
	p.DeskName = p.DeskName.FromRef(js.Undefined)
}

type OnDeskAddedEventCallbackFunc func(this js.Ref, deskId js.String, fromUndo bool) js.Ref

func (fn OnDeskAddedEventCallbackFunc) Register() js.Func[func(deskId js.String, fromUndo bool)] {
	return js.RegisterCallback[func(deskId js.String, fromUndo bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeskAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeskAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deskId js.String, fromUndo bool) js.Ref
	Arg T
}

func (cb *OnDeskAddedEventCallback[T]) Register() js.Func[func(deskId js.String, fromUndo bool)] {
	return js.RegisterCallback[func(deskId js.String, fromUndo bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeskAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeskAdded returns true if the function "WEBEXT.wmDesksPrivate.OnDeskAdded.addListener" exists.
func HasFuncOnDeskAdded() bool {
	return js.True == bindings.HasFuncOnDeskAdded()
}

// FuncOnDeskAdded returns the function "WEBEXT.wmDesksPrivate.OnDeskAdded.addListener".
func FuncOnDeskAdded() (fn js.Func[func(callback js.Func[func(deskId js.String, fromUndo bool)])]) {
	bindings.FuncOnDeskAdded(
		js.Pointer(&fn),
	)
	return
}

// OnDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.addListener" directly.
func OnDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret js.Void) {
	bindings.CallOnDeskAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeskAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeskAdded returns true if the function "WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener" exists.
func HasFuncOffDeskAdded() bool {
	return js.True == bindings.HasFuncOffDeskAdded()
}

// FuncOffDeskAdded returns the function "WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener".
func FuncOffDeskAdded() (fn js.Func[func(callback js.Func[func(deskId js.String, fromUndo bool)])]) {
	bindings.FuncOffDeskAdded(
		js.Pointer(&fn),
	)
	return
}

// OffDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener" directly.
func OffDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret js.Void) {
	bindings.CallOffDeskAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeskAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeskAdded returns true if the function "WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener" exists.
func HasFuncHasOnDeskAdded() bool {
	return js.True == bindings.HasFuncHasOnDeskAdded()
}

// FuncHasOnDeskAdded returns the function "WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener".
func FuncHasOnDeskAdded() (fn js.Func[func(callback js.Func[func(deskId js.String, fromUndo bool)]) bool]) {
	bindings.FuncHasOnDeskAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener" directly.
func HasOnDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret bool) {
	bindings.CallHasOnDeskAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeskAdded calls the function "WEBEXT.wmDesksPrivate.OnDeskAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeskAdded(callback js.Func[func(deskId js.String, fromUndo bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeskAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeskRemovedEventCallbackFunc func(this js.Ref, deskId js.String) js.Ref

func (fn OnDeskRemovedEventCallbackFunc) Register() js.Func[func(deskId js.String)] {
	return js.RegisterCallback[func(deskId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeskRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeskRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deskId js.String) js.Ref
	Arg T
}

func (cb *OnDeskRemovedEventCallback[T]) Register() js.Func[func(deskId js.String)] {
	return js.RegisterCallback[func(deskId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeskRemovedEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeskRemoved returns true if the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener" exists.
func HasFuncOnDeskRemoved() bool {
	return js.True == bindings.HasFuncOnDeskRemoved()
}

// FuncOnDeskRemoved returns the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener".
func FuncOnDeskRemoved() (fn js.Func[func(callback js.Func[func(deskId js.String)])]) {
	bindings.FuncOnDeskRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener" directly.
func OnDeskRemoved(callback js.Func[func(deskId js.String)]) (ret js.Void) {
	bindings.CallOnDeskRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeskRemoved(callback js.Func[func(deskId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeskRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeskRemoved returns true if the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener" exists.
func HasFuncOffDeskRemoved() bool {
	return js.True == bindings.HasFuncOffDeskRemoved()
}

// FuncOffDeskRemoved returns the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener".
func FuncOffDeskRemoved() (fn js.Func[func(callback js.Func[func(deskId js.String)])]) {
	bindings.FuncOffDeskRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener" directly.
func OffDeskRemoved(callback js.Func[func(deskId js.String)]) (ret js.Void) {
	bindings.CallOffDeskRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeskRemoved(callback js.Func[func(deskId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeskRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeskRemoved returns true if the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener" exists.
func HasFuncHasOnDeskRemoved() bool {
	return js.True == bindings.HasFuncHasOnDeskRemoved()
}

// FuncHasOnDeskRemoved returns the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener".
func FuncHasOnDeskRemoved() (fn js.Func[func(callback js.Func[func(deskId js.String)]) bool]) {
	bindings.FuncHasOnDeskRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener" directly.
func HasOnDeskRemoved(callback js.Func[func(deskId js.String)]) (ret bool) {
	bindings.CallHasOnDeskRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeskRemoved calls the function "WEBEXT.wmDesksPrivate.OnDeskRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeskRemoved(callback js.Func[func(deskId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeskRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeskSwitchedEventCallbackFunc func(this js.Ref, activated js.String, deactivated js.String) js.Ref

func (fn OnDeskSwitchedEventCallbackFunc) Register() js.Func[func(activated js.String, deactivated js.String)] {
	return js.RegisterCallback[func(activated js.String, deactivated js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeskSwitchedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeskSwitchedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activated js.String, deactivated js.String) js.Ref
	Arg T
}

func (cb *OnDeskSwitchedEventCallback[T]) Register() js.Func[func(activated js.String, deactivated js.String)] {
	return js.RegisterCallback[func(activated js.String, deactivated js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeskSwitchedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeskSwitched returns true if the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener" exists.
func HasFuncOnDeskSwitched() bool {
	return js.True == bindings.HasFuncOnDeskSwitched()
}

// FuncOnDeskSwitched returns the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener".
func FuncOnDeskSwitched() (fn js.Func[func(callback js.Func[func(activated js.String, deactivated js.String)])]) {
	bindings.FuncOnDeskSwitched(
		js.Pointer(&fn),
	)
	return
}

// OnDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener" directly.
func OnDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret js.Void) {
	bindings.CallOnDeskSwitched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeskSwitched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeskSwitched returns true if the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener" exists.
func HasFuncOffDeskSwitched() bool {
	return js.True == bindings.HasFuncOffDeskSwitched()
}

// FuncOffDeskSwitched returns the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener".
func FuncOffDeskSwitched() (fn js.Func[func(callback js.Func[func(activated js.String, deactivated js.String)])]) {
	bindings.FuncOffDeskSwitched(
		js.Pointer(&fn),
	)
	return
}

// OffDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener" directly.
func OffDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret js.Void) {
	bindings.CallOffDeskSwitched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeskSwitched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeskSwitched returns true if the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener" exists.
func HasFuncHasOnDeskSwitched() bool {
	return js.True == bindings.HasFuncHasOnDeskSwitched()
}

// FuncHasOnDeskSwitched returns the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener".
func FuncHasOnDeskSwitched() (fn js.Func[func(callback js.Func[func(activated js.String, deactivated js.String)]) bool]) {
	bindings.FuncHasOnDeskSwitched(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener" directly.
func HasOnDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret bool) {
	bindings.CallHasOnDeskSwitched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeskSwitched calls the function "WEBEXT.wmDesksPrivate.OnDeskSwitched.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeskSwitched(callback js.Func[func(activated js.String, deactivated js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeskSwitched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type RemoveDeskOptions struct {
	// CombineDesks is "RemoveDeskOptions.combineDesks"
	//
	// Optional
	//
	// NOTE: FFI_USE_CombineDesks MUST be set to true to make this field effective.
	CombineDesks bool
	// AllowUndo is "RemoveDeskOptions.allowUndo"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowUndo MUST be set to true to make this field effective.
	AllowUndo bool

	FFI_USE_CombineDesks bool // for CombineDesks.
	FFI_USE_AllowUndo    bool // for AllowUndo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveDeskOptions with all fields set.
func (p RemoveDeskOptions) FromRef(ref js.Ref) RemoveDeskOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveDeskOptions in the application heap.
func (p RemoveDeskOptions) New() js.Ref {
	return bindings.RemoveDeskOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveDeskOptions) UpdateFrom(ref js.Ref) {
	bindings.RemoveDeskOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveDeskOptions) Update(ref js.Ref) {
	bindings.RemoveDeskOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveDeskOptions) FreeMembers(recursive bool) {
}

type SaveActiveDeskCallbackFunc func(this js.Ref, desk *SavedDesk) js.Ref

func (fn SaveActiveDeskCallbackFunc) Register() js.Func[func(desk *SavedDesk)] {
	return js.RegisterCallback[func(desk *SavedDesk)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SaveActiveDeskCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SavedDesk
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

type SaveActiveDeskCallback[T any] struct {
	Fn  func(arg T, this js.Ref, desk *SavedDesk) js.Ref
	Arg T
}

func (cb *SaveActiveDeskCallback[T]) Register() js.Func[func(desk *SavedDesk)] {
	return js.RegisterCallback[func(desk *SavedDesk)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SaveActiveDeskCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SavedDesk
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

type WindowProperties struct {
	// AllDesks is "WindowProperties.allDesks"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllDesks MUST be set to true to make this field effective.
	AllDesks bool

	FFI_USE_AllDesks bool // for AllDesks.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WindowProperties with all fields set.
func (p WindowProperties) FromRef(ref js.Ref) WindowProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WindowProperties in the application heap.
func (p WindowProperties) New() js.Ref {
	return bindings.WindowPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WindowProperties) UpdateFrom(ref js.Ref) {
	bindings.WindowPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WindowProperties) Update(ref js.Ref) {
	bindings.WindowPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WindowProperties) FreeMembers(recursive bool) {
}

// HasFuncDeleteSavedDesk returns true if the function "WEBEXT.wmDesksPrivate.deleteSavedDesk" exists.
func HasFuncDeleteSavedDesk() bool {
	return js.True == bindings.HasFuncDeleteSavedDesk()
}

// FuncDeleteSavedDesk returns the function "WEBEXT.wmDesksPrivate.deleteSavedDesk".
func FuncDeleteSavedDesk() (fn js.Func[func(savedDeskUuid js.String) js.Promise[js.Void]]) {
	bindings.FuncDeleteSavedDesk(
		js.Pointer(&fn),
	)
	return
}

// DeleteSavedDesk calls the function "WEBEXT.wmDesksPrivate.deleteSavedDesk" directly.
func DeleteSavedDesk(savedDeskUuid js.String) (ret js.Promise[js.Void]) {
	bindings.CallDeleteSavedDesk(
		js.Pointer(&ret),
		savedDeskUuid.Ref(),
	)

	return
}

// TryDeleteSavedDesk calls the function "WEBEXT.wmDesksPrivate.deleteSavedDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteSavedDesk(savedDeskUuid js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteSavedDesk(
		js.Pointer(&ret), js.Pointer(&exception),
		savedDeskUuid.Ref(),
	)

	return
}

// HasFuncGetActiveDesk returns true if the function "WEBEXT.wmDesksPrivate.getActiveDesk" exists.
func HasFuncGetActiveDesk() bool {
	return js.True == bindings.HasFuncGetActiveDesk()
}

// FuncGetActiveDesk returns the function "WEBEXT.wmDesksPrivate.getActiveDesk".
func FuncGetActiveDesk() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetActiveDesk(
		js.Pointer(&fn),
	)
	return
}

// GetActiveDesk calls the function "WEBEXT.wmDesksPrivate.getActiveDesk" directly.
func GetActiveDesk() (ret js.Promise[js.String]) {
	bindings.CallGetActiveDesk(
		js.Pointer(&ret),
	)

	return
}

// TryGetActiveDesk calls the function "WEBEXT.wmDesksPrivate.getActiveDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetActiveDesk() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetActiveDesk(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAllDesks returns true if the function "WEBEXT.wmDesksPrivate.getAllDesks" exists.
func HasFuncGetAllDesks() bool {
	return js.True == bindings.HasFuncGetAllDesks()
}

// FuncGetAllDesks returns the function "WEBEXT.wmDesksPrivate.getAllDesks".
func FuncGetAllDesks() (fn js.Func[func() js.Promise[js.Array[Desk]]]) {
	bindings.FuncGetAllDesks(
		js.Pointer(&fn),
	)
	return
}

// GetAllDesks calls the function "WEBEXT.wmDesksPrivate.getAllDesks" directly.
func GetAllDesks() (ret js.Promise[js.Array[Desk]]) {
	bindings.CallGetAllDesks(
		js.Pointer(&ret),
	)

	return
}

// TryGetAllDesks calls the function "WEBEXT.wmDesksPrivate.getAllDesks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllDesks() (ret js.Promise[js.Array[Desk]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllDesks(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeskByID returns true if the function "WEBEXT.wmDesksPrivate.getDeskByID" exists.
func HasFuncGetDeskByID() bool {
	return js.True == bindings.HasFuncGetDeskByID()
}

// FuncGetDeskByID returns the function "WEBEXT.wmDesksPrivate.getDeskByID".
func FuncGetDeskByID() (fn js.Func[func(deskUuid js.String) js.Promise[Desk]]) {
	bindings.FuncGetDeskByID(
		js.Pointer(&fn),
	)
	return
}

// GetDeskByID calls the function "WEBEXT.wmDesksPrivate.getDeskByID" directly.
func GetDeskByID(deskUuid js.String) (ret js.Promise[Desk]) {
	bindings.CallGetDeskByID(
		js.Pointer(&ret),
		deskUuid.Ref(),
	)

	return
}

// TryGetDeskByID calls the function "WEBEXT.wmDesksPrivate.getDeskByID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeskByID(deskUuid js.String) (ret js.Promise[Desk], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeskByID(
		js.Pointer(&ret), js.Pointer(&exception),
		deskUuid.Ref(),
	)

	return
}

// HasFuncGetDeskTemplateJson returns true if the function "WEBEXT.wmDesksPrivate.getDeskTemplateJson" exists.
func HasFuncGetDeskTemplateJson() bool {
	return js.True == bindings.HasFuncGetDeskTemplateJson()
}

// FuncGetDeskTemplateJson returns the function "WEBEXT.wmDesksPrivate.getDeskTemplateJson".
func FuncGetDeskTemplateJson() (fn js.Func[func(templateUuid js.String) js.Promise[js.String]]) {
	bindings.FuncGetDeskTemplateJson(
		js.Pointer(&fn),
	)
	return
}

// GetDeskTemplateJson calls the function "WEBEXT.wmDesksPrivate.getDeskTemplateJson" directly.
func GetDeskTemplateJson(templateUuid js.String) (ret js.Promise[js.String]) {
	bindings.CallGetDeskTemplateJson(
		js.Pointer(&ret),
		templateUuid.Ref(),
	)

	return
}

// TryGetDeskTemplateJson calls the function "WEBEXT.wmDesksPrivate.getDeskTemplateJson"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeskTemplateJson(templateUuid js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeskTemplateJson(
		js.Pointer(&ret), js.Pointer(&exception),
		templateUuid.Ref(),
	)

	return
}

// HasFuncGetSavedDesks returns true if the function "WEBEXT.wmDesksPrivate.getSavedDesks" exists.
func HasFuncGetSavedDesks() bool {
	return js.True == bindings.HasFuncGetSavedDesks()
}

// FuncGetSavedDesks returns the function "WEBEXT.wmDesksPrivate.getSavedDesks".
func FuncGetSavedDesks() (fn js.Func[func() js.Promise[js.Array[SavedDesk]]]) {
	bindings.FuncGetSavedDesks(
		js.Pointer(&fn),
	)
	return
}

// GetSavedDesks calls the function "WEBEXT.wmDesksPrivate.getSavedDesks" directly.
func GetSavedDesks() (ret js.Promise[js.Array[SavedDesk]]) {
	bindings.CallGetSavedDesks(
		js.Pointer(&ret),
	)

	return
}

// TryGetSavedDesks calls the function "WEBEXT.wmDesksPrivate.getSavedDesks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSavedDesks() (ret js.Promise[js.Array[SavedDesk]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSavedDesks(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLaunchDesk returns true if the function "WEBEXT.wmDesksPrivate.launchDesk" exists.
func HasFuncLaunchDesk() bool {
	return js.True == bindings.HasFuncLaunchDesk()
}

// FuncLaunchDesk returns the function "WEBEXT.wmDesksPrivate.launchDesk".
func FuncLaunchDesk() (fn js.Func[func(launchOptions LaunchOptions) js.Promise[js.String]]) {
	bindings.FuncLaunchDesk(
		js.Pointer(&fn),
	)
	return
}

// LaunchDesk calls the function "WEBEXT.wmDesksPrivate.launchDesk" directly.
func LaunchDesk(launchOptions LaunchOptions) (ret js.Promise[js.String]) {
	bindings.CallLaunchDesk(
		js.Pointer(&ret),
		js.Pointer(&launchOptions),
	)

	return
}

// TryLaunchDesk calls the function "WEBEXT.wmDesksPrivate.launchDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchDesk(launchOptions LaunchOptions) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchDesk(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&launchOptions),
	)

	return
}

// HasFuncRecallSavedDesk returns true if the function "WEBEXT.wmDesksPrivate.recallSavedDesk" exists.
func HasFuncRecallSavedDesk() bool {
	return js.True == bindings.HasFuncRecallSavedDesk()
}

// FuncRecallSavedDesk returns the function "WEBEXT.wmDesksPrivate.recallSavedDesk".
func FuncRecallSavedDesk() (fn js.Func[func(savedDeskUuid js.String) js.Promise[js.String]]) {
	bindings.FuncRecallSavedDesk(
		js.Pointer(&fn),
	)
	return
}

// RecallSavedDesk calls the function "WEBEXT.wmDesksPrivate.recallSavedDesk" directly.
func RecallSavedDesk(savedDeskUuid js.String) (ret js.Promise[js.String]) {
	bindings.CallRecallSavedDesk(
		js.Pointer(&ret),
		savedDeskUuid.Ref(),
	)

	return
}

// TryRecallSavedDesk calls the function "WEBEXT.wmDesksPrivate.recallSavedDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecallSavedDesk(savedDeskUuid js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecallSavedDesk(
		js.Pointer(&ret), js.Pointer(&exception),
		savedDeskUuid.Ref(),
	)

	return
}

// HasFuncRemoveDesk returns true if the function "WEBEXT.wmDesksPrivate.removeDesk" exists.
func HasFuncRemoveDesk() bool {
	return js.True == bindings.HasFuncRemoveDesk()
}

// FuncRemoveDesk returns the function "WEBEXT.wmDesksPrivate.removeDesk".
func FuncRemoveDesk() (fn js.Func[func(deskId js.String, removeDeskOptions RemoveDeskOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveDesk(
		js.Pointer(&fn),
	)
	return
}

// RemoveDesk calls the function "WEBEXT.wmDesksPrivate.removeDesk" directly.
func RemoveDesk(deskId js.String, removeDeskOptions RemoveDeskOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveDesk(
		js.Pointer(&ret),
		deskId.Ref(),
		js.Pointer(&removeDeskOptions),
	)

	return
}

// TryRemoveDesk calls the function "WEBEXT.wmDesksPrivate.removeDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveDesk(deskId js.String, removeDeskOptions RemoveDeskOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveDesk(
		js.Pointer(&ret), js.Pointer(&exception),
		deskId.Ref(),
		js.Pointer(&removeDeskOptions),
	)

	return
}

// HasFuncSaveActiveDesk returns true if the function "WEBEXT.wmDesksPrivate.saveActiveDesk" exists.
func HasFuncSaveActiveDesk() bool {
	return js.True == bindings.HasFuncSaveActiveDesk()
}

// FuncSaveActiveDesk returns the function "WEBEXT.wmDesksPrivate.saveActiveDesk".
func FuncSaveActiveDesk() (fn js.Func[func() js.Promise[SavedDesk]]) {
	bindings.FuncSaveActiveDesk(
		js.Pointer(&fn),
	)
	return
}

// SaveActiveDesk calls the function "WEBEXT.wmDesksPrivate.saveActiveDesk" directly.
func SaveActiveDesk() (ret js.Promise[SavedDesk]) {
	bindings.CallSaveActiveDesk(
		js.Pointer(&ret),
	)

	return
}

// TrySaveActiveDesk calls the function "WEBEXT.wmDesksPrivate.saveActiveDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySaveActiveDesk() (ret js.Promise[SavedDesk], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySaveActiveDesk(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetWindowProperties returns true if the function "WEBEXT.wmDesksPrivate.setWindowProperties" exists.
func HasFuncSetWindowProperties() bool {
	return js.True == bindings.HasFuncSetWindowProperties()
}

// FuncSetWindowProperties returns the function "WEBEXT.wmDesksPrivate.setWindowProperties".
func FuncSetWindowProperties() (fn js.Func[func(windowId int32, windowProperties WindowProperties) js.Promise[js.Void]]) {
	bindings.FuncSetWindowProperties(
		js.Pointer(&fn),
	)
	return
}

// SetWindowProperties calls the function "WEBEXT.wmDesksPrivate.setWindowProperties" directly.
func SetWindowProperties(windowId int32, windowProperties WindowProperties) (ret js.Promise[js.Void]) {
	bindings.CallSetWindowProperties(
		js.Pointer(&ret),
		int32(windowId),
		js.Pointer(&windowProperties),
	)

	return
}

// TrySetWindowProperties calls the function "WEBEXT.wmDesksPrivate.setWindowProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetWindowProperties(windowId int32, windowProperties WindowProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetWindowProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(windowId),
		js.Pointer(&windowProperties),
	)

	return
}

// HasFuncSwitchDesk returns true if the function "WEBEXT.wmDesksPrivate.switchDesk" exists.
func HasFuncSwitchDesk() bool {
	return js.True == bindings.HasFuncSwitchDesk()
}

// FuncSwitchDesk returns the function "WEBEXT.wmDesksPrivate.switchDesk".
func FuncSwitchDesk() (fn js.Func[func(deskUuid js.String) js.Promise[js.Void]]) {
	bindings.FuncSwitchDesk(
		js.Pointer(&fn),
	)
	return
}

// SwitchDesk calls the function "WEBEXT.wmDesksPrivate.switchDesk" directly.
func SwitchDesk(deskUuid js.String) (ret js.Promise[js.Void]) {
	bindings.CallSwitchDesk(
		js.Pointer(&ret),
		deskUuid.Ref(),
	)

	return
}

// TrySwitchDesk calls the function "WEBEXT.wmDesksPrivate.switchDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySwitchDesk(deskUuid js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySwitchDesk(
		js.Pointer(&ret), js.Pointer(&exception),
		deskUuid.Ref(),
	)

	return
}

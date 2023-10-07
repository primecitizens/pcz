// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tabgroups

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabgroups/bindings"
)

type Color uint32

const (
	_ Color = iota

	Color_GREY
	Color_BLUE
	Color_RED
	Color_YELLOW
	Color_GREEN
	Color_PINK
	Color_PURPLE
	Color_CYAN
	Color_ORANGE
)

func (Color) FromRef(str js.Ref) Color {
	return Color(bindings.ConstOfColor(str))
}

func (x Color) String() (string, bool) {
	switch x {
	case Color_GREY:
		return "grey", true
	case Color_BLUE:
		return "blue", true
	case Color_RED:
		return "red", true
	case Color_YELLOW:
		return "yellow", true
	case Color_GREEN:
		return "green", true
	case Color_PINK:
		return "pink", true
	case Color_PURPLE:
		return "purple", true
	case Color_CYAN:
		return "cyan", true
	case Color_ORANGE:
		return "orange", true
	default:
		return "", false
	}
}

type MoveArgMoveProperties struct {
	// Index is "MoveArgMoveProperties.index"
	//
	// Required
	Index int64
	// WindowId is "MoveArgMoveProperties.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MoveArgMoveProperties with all fields set.
func (p MoveArgMoveProperties) FromRef(ref js.Ref) MoveArgMoveProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MoveArgMoveProperties in the application heap.
func (p MoveArgMoveProperties) New() js.Ref {
	return bindings.MoveArgMovePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MoveArgMoveProperties) UpdateFrom(ref js.Ref) {
	bindings.MoveArgMovePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MoveArgMoveProperties) Update(ref js.Ref) {
	bindings.MoveArgMovePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MoveArgMoveProperties) FreeMembers(recursive bool) {
}

type QueryArgQueryInfo struct {
	// Collapsed is "QueryArgQueryInfo.collapsed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Collapsed MUST be set to true to make this field effective.
	Collapsed bool
	// Color is "QueryArgQueryInfo.color"
	//
	// Optional
	Color Color
	// Title is "QueryArgQueryInfo.title"
	//
	// Optional
	Title js.String
	// WindowId is "QueryArgQueryInfo.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_Collapsed bool // for Collapsed.
	FFI_USE_WindowId  bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryArgQueryInfo with all fields set.
func (p QueryArgQueryInfo) FromRef(ref js.Ref) QueryArgQueryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryArgQueryInfo in the application heap.
func (p QueryArgQueryInfo) New() js.Ref {
	return bindings.QueryArgQueryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryArgQueryInfo) UpdateFrom(ref js.Ref) {
	bindings.QueryArgQueryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryArgQueryInfo) Update(ref js.Ref) {
	bindings.QueryArgQueryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryArgQueryInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
}

// TAB_GROUP_ID_NONE returns the value of property "WEBEXT.tabGroups.TAB_GROUP_ID_NONE".
//
// The returned bool will be false if there is no such property.
func TAB_GROUP_ID_NONE() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTAB_GROUP_ID_NONE(
		js.Pointer(&ret),
	)

	return
}

// SetTAB_GROUP_ID_NONE sets the value of property "WEBEXT.tabGroups.TAB_GROUP_ID_NONE" to val.
//
// It returns false if the property cannot be set.
func SetTAB_GROUP_ID_NONE(val js.String) bool {
	return js.True == bindings.SetTAB_GROUP_ID_NONE(
		val.Ref())
}

type TabGroup struct {
	// Collapsed is "TabGroup.collapsed"
	//
	// Required
	Collapsed bool
	// Color is "TabGroup.color"
	//
	// Required
	Color Color
	// Id is "TabGroup.id"
	//
	// Required
	Id int64
	// Title is "TabGroup.title"
	//
	// Optional
	Title js.String
	// WindowId is "TabGroup.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TabGroup with all fields set.
func (p TabGroup) FromRef(ref js.Ref) TabGroup {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TabGroup in the application heap.
func (p TabGroup) New() js.Ref {
	return bindings.TabGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TabGroup) UpdateFrom(ref js.Ref) {
	bindings.TabGroupJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TabGroup) Update(ref js.Ref) {
	bindings.TabGroupJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TabGroup) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
}

type UpdateArgUpdateProperties struct {
	// Collapsed is "UpdateArgUpdateProperties.collapsed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Collapsed MUST be set to true to make this field effective.
	Collapsed bool
	// Color is "UpdateArgUpdateProperties.color"
	//
	// Optional
	Color Color
	// Title is "UpdateArgUpdateProperties.title"
	//
	// Optional
	Title js.String

	FFI_USE_Collapsed bool // for Collapsed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateArgUpdateProperties with all fields set.
func (p UpdateArgUpdateProperties) FromRef(ref js.Ref) UpdateArgUpdateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateArgUpdateProperties in the application heap.
func (p UpdateArgUpdateProperties) New() js.Ref {
	return bindings.UpdateArgUpdatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateArgUpdateProperties) UpdateFrom(ref js.Ref) {
	bindings.UpdateArgUpdatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateArgUpdateProperties) Update(ref js.Ref) {
	bindings.UpdateArgUpdatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateArgUpdateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
}

// HasFuncGet returns true if the function "WEBEXT.tabGroups.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.tabGroups.get".
func FuncGet() (fn js.Func[func(groupId int64) js.Promise[TabGroup]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.tabGroups.get" directly.
func Get(groupId int64) (ret js.Promise[TabGroup]) {
	bindings.CallGet(
		js.Pointer(&ret),
		float64(groupId),
	)

	return
}

// TryGet calls the function "WEBEXT.tabGroups.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(groupId int64) (ret js.Promise[TabGroup], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(groupId),
	)

	return
}

// HasFuncMove returns true if the function "WEBEXT.tabGroups.move" exists.
func HasFuncMove() bool {
	return js.True == bindings.HasFuncMove()
}

// FuncMove returns the function "WEBEXT.tabGroups.move".
func FuncMove() (fn js.Func[func(groupId int64, moveProperties MoveArgMoveProperties) js.Promise[TabGroup]]) {
	bindings.FuncMove(
		js.Pointer(&fn),
	)
	return
}

// Move calls the function "WEBEXT.tabGroups.move" directly.
func Move(groupId int64, moveProperties MoveArgMoveProperties) (ret js.Promise[TabGroup]) {
	bindings.CallMove(
		js.Pointer(&ret),
		float64(groupId),
		js.Pointer(&moveProperties),
	)

	return
}

// TryMove calls the function "WEBEXT.tabGroups.move"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMove(groupId int64, moveProperties MoveArgMoveProperties) (ret js.Promise[TabGroup], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMove(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(groupId),
		js.Pointer(&moveProperties),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, group *TabGroup) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

type OnCreatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, group *TabGroup) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

// HasFuncOnCreated returns true if the function "WEBEXT.tabGroups.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.tabGroups.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.tabGroups.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.tabGroups.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.tabGroups.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.tabGroups.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.tabGroups.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.tabGroups.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.tabGroups.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.tabGroups.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(group *TabGroup)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.tabGroups.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(group *TabGroup)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.tabGroups.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(group *TabGroup)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMovedEventCallbackFunc func(this js.Ref, group *TabGroup) js.Ref

func (fn OnMovedEventCallbackFunc) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

type OnMovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, group *TabGroup) js.Ref
	Arg T
}

func (cb *OnMovedEventCallback[T]) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

// HasFuncOnMoved returns true if the function "WEBEXT.tabGroups.onMoved.addListener" exists.
func HasFuncOnMoved() bool {
	return js.True == bindings.HasFuncOnMoved()
}

// FuncOnMoved returns the function "WEBEXT.tabGroups.onMoved.addListener".
func FuncOnMoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOnMoved(
		js.Pointer(&fn),
	)
	return
}

// OnMoved calls the function "WEBEXT.tabGroups.onMoved.addListener" directly.
func OnMoved(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMoved calls the function "WEBEXT.tabGroups.onMoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMoved(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMoved returns true if the function "WEBEXT.tabGroups.onMoved.removeListener" exists.
func HasFuncOffMoved() bool {
	return js.True == bindings.HasFuncOffMoved()
}

// FuncOffMoved returns the function "WEBEXT.tabGroups.onMoved.removeListener".
func FuncOffMoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOffMoved(
		js.Pointer(&fn),
	)
	return
}

// OffMoved calls the function "WEBEXT.tabGroups.onMoved.removeListener" directly.
func OffMoved(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOffMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMoved calls the function "WEBEXT.tabGroups.onMoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMoved(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMoved returns true if the function "WEBEXT.tabGroups.onMoved.hasListener" exists.
func HasFuncHasOnMoved() bool {
	return js.True == bindings.HasFuncHasOnMoved()
}

// FuncHasOnMoved returns the function "WEBEXT.tabGroups.onMoved.hasListener".
func FuncHasOnMoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)]) bool]) {
	bindings.FuncHasOnMoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnMoved calls the function "WEBEXT.tabGroups.onMoved.hasListener" directly.
func HasOnMoved(callback js.Func[func(group *TabGroup)]) (ret bool) {
	bindings.CallHasOnMoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMoved calls the function "WEBEXT.tabGroups.onMoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMoved(callback js.Func[func(group *TabGroup)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemovedEventCallbackFunc func(this js.Ref, group *TabGroup) js.Ref

func (fn OnRemovedEventCallbackFunc) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

type OnRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, group *TabGroup) js.Ref
	Arg T
}

func (cb *OnRemovedEventCallback[T]) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

// HasFuncOnRemoved returns true if the function "WEBEXT.tabGroups.onRemoved.addListener" exists.
func HasFuncOnRemoved() bool {
	return js.True == bindings.HasFuncOnRemoved()
}

// FuncOnRemoved returns the function "WEBEXT.tabGroups.onRemoved.addListener".
func FuncOnRemoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnRemoved calls the function "WEBEXT.tabGroups.onRemoved.addListener" directly.
func OnRemoved(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoved calls the function "WEBEXT.tabGroups.onRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoved(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoved returns true if the function "WEBEXT.tabGroups.onRemoved.removeListener" exists.
func HasFuncOffRemoved() bool {
	return js.True == bindings.HasFuncOffRemoved()
}

// FuncOffRemoved returns the function "WEBEXT.tabGroups.onRemoved.removeListener".
func FuncOffRemoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOffRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffRemoved calls the function "WEBEXT.tabGroups.onRemoved.removeListener" directly.
func OffRemoved(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOffRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoved calls the function "WEBEXT.tabGroups.onRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoved(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoved returns true if the function "WEBEXT.tabGroups.onRemoved.hasListener" exists.
func HasFuncHasOnRemoved() bool {
	return js.True == bindings.HasFuncHasOnRemoved()
}

// FuncHasOnRemoved returns the function "WEBEXT.tabGroups.onRemoved.hasListener".
func FuncHasOnRemoved() (fn js.Func[func(callback js.Func[func(group *TabGroup)]) bool]) {
	bindings.FuncHasOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoved calls the function "WEBEXT.tabGroups.onRemoved.hasListener" directly.
func HasOnRemoved(callback js.Func[func(group *TabGroup)]) (ret bool) {
	bindings.CallHasOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoved calls the function "WEBEXT.tabGroups.onRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoved(callback js.Func[func(group *TabGroup)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUpdatedEventCallbackFunc func(this js.Ref, group *TabGroup) js.Ref

func (fn OnUpdatedEventCallbackFunc) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

type OnUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, group *TabGroup) js.Ref
	Arg T
}

func (cb *OnUpdatedEventCallback[T]) Register() js.Func[func(group *TabGroup)] {
	return js.RegisterCallback[func(group *TabGroup)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUpdatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TabGroup
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

// HasFuncOnUpdated returns true if the function "WEBEXT.tabGroups.onUpdated.addListener" exists.
func HasFuncOnUpdated() bool {
	return js.True == bindings.HasFuncOnUpdated()
}

// FuncOnUpdated returns the function "WEBEXT.tabGroups.onUpdated.addListener".
func FuncOnUpdated() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnUpdated calls the function "WEBEXT.tabGroups.onUpdated.addListener" directly.
func OnUpdated(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUpdated calls the function "WEBEXT.tabGroups.onUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUpdated(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUpdated returns true if the function "WEBEXT.tabGroups.onUpdated.removeListener" exists.
func HasFuncOffUpdated() bool {
	return js.True == bindings.HasFuncOffUpdated()
}

// FuncOffUpdated returns the function "WEBEXT.tabGroups.onUpdated.removeListener".
func FuncOffUpdated() (fn js.Func[func(callback js.Func[func(group *TabGroup)])]) {
	bindings.FuncOffUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffUpdated calls the function "WEBEXT.tabGroups.onUpdated.removeListener" directly.
func OffUpdated(callback js.Func[func(group *TabGroup)]) (ret js.Void) {
	bindings.CallOffUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUpdated calls the function "WEBEXT.tabGroups.onUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUpdated(callback js.Func[func(group *TabGroup)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUpdated returns true if the function "WEBEXT.tabGroups.onUpdated.hasListener" exists.
func HasFuncHasOnUpdated() bool {
	return js.True == bindings.HasFuncHasOnUpdated()
}

// FuncHasOnUpdated returns the function "WEBEXT.tabGroups.onUpdated.hasListener".
func FuncHasOnUpdated() (fn js.Func[func(callback js.Func[func(group *TabGroup)]) bool]) {
	bindings.FuncHasOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnUpdated calls the function "WEBEXT.tabGroups.onUpdated.hasListener" directly.
func HasOnUpdated(callback js.Func[func(group *TabGroup)]) (ret bool) {
	bindings.CallHasOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUpdated calls the function "WEBEXT.tabGroups.onUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUpdated(callback js.Func[func(group *TabGroup)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncQuery returns true if the function "WEBEXT.tabGroups.query" exists.
func HasFuncQuery() bool {
	return js.True == bindings.HasFuncQuery()
}

// FuncQuery returns the function "WEBEXT.tabGroups.query".
func FuncQuery() (fn js.Func[func(queryInfo QueryArgQueryInfo) js.Promise[js.Array[TabGroup]]]) {
	bindings.FuncQuery(
		js.Pointer(&fn),
	)
	return
}

// Query calls the function "WEBEXT.tabGroups.query" directly.
func Query(queryInfo QueryArgQueryInfo) (ret js.Promise[js.Array[TabGroup]]) {
	bindings.CallQuery(
		js.Pointer(&ret),
		js.Pointer(&queryInfo),
	)

	return
}

// TryQuery calls the function "WEBEXT.tabGroups.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQuery(queryInfo QueryArgQueryInfo) (ret js.Promise[js.Array[TabGroup]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQuery(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryInfo),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.tabGroups.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.tabGroups.update".
func FuncUpdate() (fn js.Func[func(groupId int64, updateProperties UpdateArgUpdateProperties) js.Promise[TabGroup]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.tabGroups.update" directly.
func Update(groupId int64, updateProperties UpdateArgUpdateProperties) (ret js.Promise[TabGroup]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		float64(groupId),
		js.Pointer(&updateProperties),
	)

	return
}

// TryUpdate calls the function "WEBEXT.tabGroups.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(groupId int64, updateProperties UpdateArgUpdateProperties) (ret js.Promise[TabGroup], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(groupId),
		js.Pointer(&updateProperties),
	)

	return
}

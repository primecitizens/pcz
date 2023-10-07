// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package chromewebviewinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chromewebviewinternal/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/contextmenus"
)

type ContextMenuItem struct {
	// CommandId is "ContextMenuItem.commandId"
	//
	// Required
	CommandId int64
	// Label is "ContextMenuItem.label"
	//
	// Optional
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextMenuItem with all fields set.
func (p ContextMenuItem) FromRef(ref js.Ref) ContextMenuItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextMenuItem in the application heap.
func (p ContextMenuItem) New() js.Ref {
	return bindings.ContextMenuItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextMenuItem) UpdateFrom(ref js.Ref) {
	bindings.ContextMenuItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextMenuItem) Update(ref js.Ref) {
	bindings.ContextMenuItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextMenuItem) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

type ContextMenusCreateArgCallbackFunc func(this js.Ref) js.Ref

func (fn ContextMenusCreateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusCreateArgCallbackFunc) DispatchCallback(
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

type ContextMenusCreateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContextMenusCreateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusCreateArgCallback[T]) DispatchCallback(
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

type ContextMenusCreateArgCreatePropertiesFieldOnclickFunc func(this js.Ref, info *contextmenus.OnClickData) js.Ref

func (fn ContextMenusCreateArgCreatePropertiesFieldOnclickFunc) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusCreateArgCreatePropertiesFieldOnclickFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
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

type ContextMenusCreateArgCreatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref, info *contextmenus.OnClickData) js.Ref
	Arg T
}

func (cb *ContextMenusCreateArgCreatePropertiesFieldOnclick[T]) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusCreateArgCreatePropertiesFieldOnclick[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
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

type OneOf_Int64_String struct {
	ref js.Ref
}

func (x OneOf_Int64_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Int64_String) Free() {
	x.ref.Free()
}

func (x OneOf_Int64_String) FromRef(ref js.Ref) OneOf_Int64_String {
	return OneOf_Int64_String{
		ref: ref,
	}
}

func (x OneOf_Int64_String) Int64() int64 {
	return js.BigInt[int64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Int64_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type ContextMenusCreateArgCreateProperties struct {
	// Checked is "ContextMenusCreateArgCreateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "ContextMenusCreateArgCreateProperties.contexts"
	//
	// Optional
	Contexts js.Array[contextmenus.ContextType]
	// DocumentUrlPatterns is "ContextMenusCreateArgCreateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "ContextMenusCreateArgCreateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Id is "ContextMenusCreateArgCreateProperties.id"
	//
	// Optional
	Id js.String
	// Onclick is "ContextMenusCreateArgCreateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func(info *contextmenus.OnClickData)]
	// ParentId is "ContextMenusCreateArgCreateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "ContextMenusCreateArgCreateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "ContextMenusCreateArgCreateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "ContextMenusCreateArgCreateProperties.type"
	//
	// Optional
	Type contextmenus.ItemType
	// Visible is "ContextMenusCreateArgCreateProperties.visible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool

	FFI_USE_Checked bool // for Checked.
	FFI_USE_Enabled bool // for Enabled.
	FFI_USE_Visible bool // for Visible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextMenusCreateArgCreateProperties with all fields set.
func (p ContextMenusCreateArgCreateProperties) FromRef(ref js.Ref) ContextMenusCreateArgCreateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextMenusCreateArgCreateProperties in the application heap.
func (p ContextMenusCreateArgCreateProperties) New() js.Ref {
	return bindings.ContextMenusCreateArgCreatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextMenusCreateArgCreateProperties) UpdateFrom(ref js.Ref) {
	bindings.ContextMenusCreateArgCreatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextMenusCreateArgCreateProperties) Update(ref js.Ref) {
	bindings.ContextMenusCreateArgCreatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextMenusCreateArgCreateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Contexts.Ref(),
		p.DocumentUrlPatterns.Ref(),
		p.Id.Ref(),
		p.Onclick.Ref(),
		p.ParentId.Ref(),
		p.TargetUrlPatterns.Ref(),
		p.Title.Ref(),
	)
	p.Contexts = p.Contexts.FromRef(js.Undefined)
	p.DocumentUrlPatterns = p.DocumentUrlPatterns.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Onclick = p.Onclick.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.TargetUrlPatterns = p.TargetUrlPatterns.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type ContextMenusRemoveAllArgCallbackFunc func(this js.Ref) js.Ref

func (fn ContextMenusRemoveAllArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusRemoveAllArgCallbackFunc) DispatchCallback(
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

type ContextMenusRemoveAllArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContextMenusRemoveAllArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusRemoveAllArgCallback[T]) DispatchCallback(
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

type ContextMenusRemoveArgCallbackFunc func(this js.Ref) js.Ref

func (fn ContextMenusRemoveArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusRemoveArgCallbackFunc) DispatchCallback(
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

type ContextMenusRemoveArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContextMenusRemoveArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusRemoveArgCallback[T]) DispatchCallback(
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

type ContextMenusUpdateArgCallbackFunc func(this js.Ref) js.Ref

func (fn ContextMenusUpdateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusUpdateArgCallbackFunc) DispatchCallback(
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

type ContextMenusUpdateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContextMenusUpdateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusUpdateArgCallback[T]) DispatchCallback(
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

type ContextMenusUpdateArgUpdatePropertiesFieldOnclickFunc func(this js.Ref) js.Ref

func (fn ContextMenusUpdateArgUpdatePropertiesFieldOnclickFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenusUpdateArgUpdatePropertiesFieldOnclickFunc) DispatchCallback(
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

type ContextMenusUpdateArgUpdatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContextMenusUpdateArgUpdatePropertiesFieldOnclick[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenusUpdateArgUpdatePropertiesFieldOnclick[T]) DispatchCallback(
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

type ContextMenusUpdateArgUpdateProperties struct {
	// Checked is "ContextMenusUpdateArgUpdateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "ContextMenusUpdateArgUpdateProperties.contexts"
	//
	// Optional
	Contexts js.Array[contextmenus.ContextType]
	// DocumentUrlPatterns is "ContextMenusUpdateArgUpdateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "ContextMenusUpdateArgUpdateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Onclick is "ContextMenusUpdateArgUpdateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func()]
	// ParentId is "ContextMenusUpdateArgUpdateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "ContextMenusUpdateArgUpdateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "ContextMenusUpdateArgUpdateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "ContextMenusUpdateArgUpdateProperties.type"
	//
	// Optional
	Type contextmenus.ItemType
	// Visible is "ContextMenusUpdateArgUpdateProperties.visible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool

	FFI_USE_Checked bool // for Checked.
	FFI_USE_Enabled bool // for Enabled.
	FFI_USE_Visible bool // for Visible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextMenusUpdateArgUpdateProperties with all fields set.
func (p ContextMenusUpdateArgUpdateProperties) FromRef(ref js.Ref) ContextMenusUpdateArgUpdateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextMenusUpdateArgUpdateProperties in the application heap.
func (p ContextMenusUpdateArgUpdateProperties) New() js.Ref {
	return bindings.ContextMenusUpdateArgUpdatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextMenusUpdateArgUpdateProperties) UpdateFrom(ref js.Ref) {
	bindings.ContextMenusUpdateArgUpdatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextMenusUpdateArgUpdateProperties) Update(ref js.Ref) {
	bindings.ContextMenusUpdateArgUpdatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextMenusUpdateArgUpdateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Contexts.Ref(),
		p.DocumentUrlPatterns.Ref(),
		p.Onclick.Ref(),
		p.ParentId.Ref(),
		p.TargetUrlPatterns.Ref(),
		p.Title.Ref(),
	)
	p.Contexts = p.Contexts.FromRef(js.Undefined)
	p.DocumentUrlPatterns = p.DocumentUrlPatterns.FromRef(js.Undefined)
	p.Onclick = p.Onclick.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.TargetUrlPatterns = p.TargetUrlPatterns.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type OnShowArgEventFieldPreventDefaultFunc func(this js.Ref) js.Ref

func (fn OnShowArgEventFieldPreventDefaultFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnShowArgEventFieldPreventDefaultFunc) DispatchCallback(
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

type OnShowArgEventFieldPreventDefault[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnShowArgEventFieldPreventDefault[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnShowArgEventFieldPreventDefault[T]) DispatchCallback(
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

type OnShowArgEvent struct {
	// PreventDefault is "OnShowArgEvent.preventDefault"
	//
	// Required
	PreventDefault js.Func[func()]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnShowArgEvent with all fields set.
func (p OnShowArgEvent) FromRef(ref js.Ref) OnShowArgEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnShowArgEvent in the application heap.
func (p OnShowArgEvent) New() js.Ref {
	return bindings.OnShowArgEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnShowArgEvent) UpdateFrom(ref js.Ref) {
	bindings.OnShowArgEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnShowArgEvent) Update(ref js.Ref) {
	bindings.OnShowArgEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnShowArgEvent) FreeMembers(recursive bool) {
	js.Free(
		p.PreventDefault.Ref(),
	)
	p.PreventDefault = p.PreventDefault.FromRef(js.Undefined)
}

// HasFuncContextMenusCreate returns true if the function "WEBEXT.chromeWebViewInternal.contextMenusCreate" exists.
func HasFuncContextMenusCreate() bool {
	return js.True == bindings.HasFuncContextMenusCreate()
}

// FuncContextMenusCreate returns the function "WEBEXT.chromeWebViewInternal.contextMenusCreate".
func FuncContextMenusCreate() (fn js.Func[func(instanceId int64, createProperties ContextMenusCreateArgCreateProperties, callback js.Func[func()]) OneOf_Int64_String]) {
	bindings.FuncContextMenusCreate(
		js.Pointer(&fn),
	)
	return
}

// ContextMenusCreate calls the function "WEBEXT.chromeWebViewInternal.contextMenusCreate" directly.
func ContextMenusCreate(instanceId int64, createProperties ContextMenusCreateArgCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String) {
	bindings.CallContextMenusCreate(
		js.Pointer(&ret),
		float64(instanceId),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

// TryContextMenusCreate calls the function "WEBEXT.chromeWebViewInternal.contextMenusCreate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContextMenusCreate(instanceId int64, createProperties ContextMenusCreateArgCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

// HasFuncContextMenusRemove returns true if the function "WEBEXT.chromeWebViewInternal.contextMenusRemove" exists.
func HasFuncContextMenusRemove() bool {
	return js.True == bindings.HasFuncContextMenusRemove()
}

// FuncContextMenusRemove returns the function "WEBEXT.chromeWebViewInternal.contextMenusRemove".
func FuncContextMenusRemove() (fn js.Func[func(instanceId int64, menuItemId OneOf_Int64_String, callback js.Func[func()])]) {
	bindings.FuncContextMenusRemove(
		js.Pointer(&fn),
	)
	return
}

// ContextMenusRemove calls the function "WEBEXT.chromeWebViewInternal.contextMenusRemove" directly.
func ContextMenusRemove(instanceId int64, menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusRemove(
		js.Pointer(&ret),
		float64(instanceId),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// TryContextMenusRemove calls the function "WEBEXT.chromeWebViewInternal.contextMenusRemove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContextMenusRemove(instanceId int64, menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncContextMenusRemoveAll returns true if the function "WEBEXT.chromeWebViewInternal.contextMenusRemoveAll" exists.
func HasFuncContextMenusRemoveAll() bool {
	return js.True == bindings.HasFuncContextMenusRemoveAll()
}

// FuncContextMenusRemoveAll returns the function "WEBEXT.chromeWebViewInternal.contextMenusRemoveAll".
func FuncContextMenusRemoveAll() (fn js.Func[func(instanceId int64, callback js.Func[func()])]) {
	bindings.FuncContextMenusRemoveAll(
		js.Pointer(&fn),
	)
	return
}

// ContextMenusRemoveAll calls the function "WEBEXT.chromeWebViewInternal.contextMenusRemoveAll" directly.
func ContextMenusRemoveAll(instanceId int64, callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusRemoveAll(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryContextMenusRemoveAll calls the function "WEBEXT.chromeWebViewInternal.contextMenusRemoveAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContextMenusRemoveAll(instanceId int64, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusRemoveAll(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncContextMenusUpdate returns true if the function "WEBEXT.chromeWebViewInternal.contextMenusUpdate" exists.
func HasFuncContextMenusUpdate() bool {
	return js.True == bindings.HasFuncContextMenusUpdate()
}

// FuncContextMenusUpdate returns the function "WEBEXT.chromeWebViewInternal.contextMenusUpdate".
func FuncContextMenusUpdate() (fn js.Func[func(instanceId int64, id OneOf_Int64_String, updateProperties ContextMenusUpdateArgUpdateProperties, callback js.Func[func()])]) {
	bindings.FuncContextMenusUpdate(
		js.Pointer(&fn),
	)
	return
}

// ContextMenusUpdate calls the function "WEBEXT.chromeWebViewInternal.contextMenusUpdate" directly.
func ContextMenusUpdate(instanceId int64, id OneOf_Int64_String, updateProperties ContextMenusUpdateArgUpdateProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusUpdate(
		js.Pointer(&ret),
		float64(instanceId),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

// TryContextMenusUpdate calls the function "WEBEXT.chromeWebViewInternal.contextMenusUpdate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContextMenusUpdate(instanceId int64, id OneOf_Int64_String, updateProperties ContextMenusUpdateArgUpdateProperties, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

type OnClickedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnClickedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClickedEventCallbackFunc) DispatchCallback(
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

type OnClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnClickedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClickedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClicked returns true if the function "WEBEXT.chromeWebViewInternal.onClicked.addListener" exists.
func HasFuncOnClicked() bool {
	return js.True == bindings.HasFuncOnClicked()
}

// FuncOnClicked returns the function "WEBEXT.chromeWebViewInternal.onClicked.addListener".
func FuncOnClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClicked(
		js.Pointer(&fn),
	)
	return
}

// OnClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.addListener" directly.
func OnClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClicked returns true if the function "WEBEXT.chromeWebViewInternal.onClicked.removeListener" exists.
func HasFuncOffClicked() bool {
	return js.True == bindings.HasFuncOffClicked()
}

// FuncOffClicked returns the function "WEBEXT.chromeWebViewInternal.onClicked.removeListener".
func FuncOffClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClicked(
		js.Pointer(&fn),
	)
	return
}

// OffClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.removeListener" directly.
func OffClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClicked returns true if the function "WEBEXT.chromeWebViewInternal.onClicked.hasListener" exists.
func HasFuncHasOnClicked() bool {
	return js.True == bindings.HasFuncHasOnClicked()
}

// FuncHasOnClicked returns the function "WEBEXT.chromeWebViewInternal.onClicked.hasListener".
func FuncHasOnClicked() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.hasListener" directly.
func HasOnClicked(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClicked calls the function "WEBEXT.chromeWebViewInternal.onClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClicked(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnShowEventCallbackFunc func(this js.Ref, event *OnShowArgEvent) js.Ref

func (fn OnShowEventCallbackFunc) Register() js.Func[func(event *OnShowArgEvent)] {
	return js.RegisterCallback[func(event *OnShowArgEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnShowEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnShowArgEvent
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

type OnShowEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *OnShowArgEvent) js.Ref
	Arg T
}

func (cb *OnShowEventCallback[T]) Register() js.Func[func(event *OnShowArgEvent)] {
	return js.RegisterCallback[func(event *OnShowArgEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnShowEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnShowArgEvent
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

// HasFuncOnShow returns true if the function "WEBEXT.chromeWebViewInternal.onShow.addListener" exists.
func HasFuncOnShow() bool {
	return js.True == bindings.HasFuncOnShow()
}

// FuncOnShow returns the function "WEBEXT.chromeWebViewInternal.onShow.addListener".
func FuncOnShow() (fn js.Func[func(callback js.Func[func(event *OnShowArgEvent)])]) {
	bindings.FuncOnShow(
		js.Pointer(&fn),
	)
	return
}

// OnShow calls the function "WEBEXT.chromeWebViewInternal.onShow.addListener" directly.
func OnShow(callback js.Func[func(event *OnShowArgEvent)]) (ret js.Void) {
	bindings.CallOnShow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnShow calls the function "WEBEXT.chromeWebViewInternal.onShow.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnShow(callback js.Func[func(event *OnShowArgEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnShow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffShow returns true if the function "WEBEXT.chromeWebViewInternal.onShow.removeListener" exists.
func HasFuncOffShow() bool {
	return js.True == bindings.HasFuncOffShow()
}

// FuncOffShow returns the function "WEBEXT.chromeWebViewInternal.onShow.removeListener".
func FuncOffShow() (fn js.Func[func(callback js.Func[func(event *OnShowArgEvent)])]) {
	bindings.FuncOffShow(
		js.Pointer(&fn),
	)
	return
}

// OffShow calls the function "WEBEXT.chromeWebViewInternal.onShow.removeListener" directly.
func OffShow(callback js.Func[func(event *OnShowArgEvent)]) (ret js.Void) {
	bindings.CallOffShow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffShow calls the function "WEBEXT.chromeWebViewInternal.onShow.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffShow(callback js.Func[func(event *OnShowArgEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffShow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnShow returns true if the function "WEBEXT.chromeWebViewInternal.onShow.hasListener" exists.
func HasFuncHasOnShow() bool {
	return js.True == bindings.HasFuncHasOnShow()
}

// FuncHasOnShow returns the function "WEBEXT.chromeWebViewInternal.onShow.hasListener".
func FuncHasOnShow() (fn js.Func[func(callback js.Func[func(event *OnShowArgEvent)]) bool]) {
	bindings.FuncHasOnShow(
		js.Pointer(&fn),
	)
	return
}

// HasOnShow calls the function "WEBEXT.chromeWebViewInternal.onShow.hasListener" directly.
func HasOnShow(callback js.Func[func(event *OnShowArgEvent)]) (ret bool) {
	bindings.CallHasOnShow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnShow calls the function "WEBEXT.chromeWebViewInternal.onShow.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnShow(callback js.Func[func(event *OnShowArgEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnShow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncShowContextMenu returns true if the function "WEBEXT.chromeWebViewInternal.showContextMenu" exists.
func HasFuncShowContextMenu() bool {
	return js.True == bindings.HasFuncShowContextMenu()
}

// FuncShowContextMenu returns the function "WEBEXT.chromeWebViewInternal.showContextMenu".
func FuncShowContextMenu() (fn js.Func[func(instanceId int64, requestId int64, itemsToShow js.Array[ContextMenuItem])]) {
	bindings.FuncShowContextMenu(
		js.Pointer(&fn),
	)
	return
}

// ShowContextMenu calls the function "WEBEXT.chromeWebViewInternal.showContextMenu" directly.
func ShowContextMenu(instanceId int64, requestId int64, itemsToShow js.Array[ContextMenuItem]) (ret js.Void) {
	bindings.CallShowContextMenu(
		js.Pointer(&ret),
		float64(instanceId),
		float64(requestId),
		itemsToShow.Ref(),
	)

	return
}

// TryShowContextMenu calls the function "WEBEXT.chromeWebViewInternal.showContextMenu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowContextMenu(instanceId int64, requestId int64, itemsToShow js.Array[ContextMenuItem]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowContextMenu(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		float64(requestId),
		itemsToShow.Ref(),
	)

	return
}

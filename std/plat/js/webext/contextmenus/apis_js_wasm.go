// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package contextmenus

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contextmenus/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

// ACTION_MENU_TOP_LEVEL_LIMIT returns the value of property "WEBEXT.contextMenus.ACTION_MENU_TOP_LEVEL_LIMIT".
//
// The returned bool will be false if there is no such property.
func ACTION_MENU_TOP_LEVEL_LIMIT() (ret js.String, ok bool) {
	ok = js.True == bindings.GetACTION_MENU_TOP_LEVEL_LIMIT(
		js.Pointer(&ret),
	)

	return
}

// SetACTION_MENU_TOP_LEVEL_LIMIT sets the value of property "WEBEXT.contextMenus.ACTION_MENU_TOP_LEVEL_LIMIT" to val.
//
// It returns false if the property cannot be set.
func SetACTION_MENU_TOP_LEVEL_LIMIT(val js.String) bool {
	return js.True == bindings.SetACTION_MENU_TOP_LEVEL_LIMIT(
		val.Ref())
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

type OnClickData struct {
	// Checked is "OnClickData.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Editable is "OnClickData.editable"
	//
	// Required
	Editable bool
	// FrameId is "OnClickData.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// FrameUrl is "OnClickData.frameUrl"
	//
	// Optional
	FrameUrl js.String
	// LinkUrl is "OnClickData.linkUrl"
	//
	// Optional
	LinkUrl js.String
	// MediaType is "OnClickData.mediaType"
	//
	// Optional
	MediaType js.String
	// MenuItemId is "OnClickData.menuItemId"
	//
	// Required
	MenuItemId OneOf_Int64_String
	// PageUrl is "OnClickData.pageUrl"
	//
	// Optional
	PageUrl js.String
	// ParentMenuItemId is "OnClickData.parentMenuItemId"
	//
	// Optional
	ParentMenuItemId OneOf_Int64_String
	// SelectionText is "OnClickData.selectionText"
	//
	// Optional
	SelectionText js.String
	// SrcUrl is "OnClickData.srcUrl"
	//
	// Optional
	SrcUrl js.String
	// WasChecked is "OnClickData.wasChecked"
	//
	// Optional
	//
	// NOTE: FFI_USE_WasChecked MUST be set to true to make this field effective.
	WasChecked bool

	FFI_USE_Checked    bool // for Checked.
	FFI_USE_FrameId    bool // for FrameId.
	FFI_USE_WasChecked bool // for WasChecked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnClickData with all fields set.
func (p OnClickData) FromRef(ref js.Ref) OnClickData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnClickData in the application heap.
func (p OnClickData) New() js.Ref {
	return bindings.OnClickDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnClickData) UpdateFrom(ref js.Ref) {
	bindings.OnClickDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnClickData) Update(ref js.Ref) {
	bindings.OnClickDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnClickData) FreeMembers(recursive bool) {
	js.Free(
		p.FrameUrl.Ref(),
		p.LinkUrl.Ref(),
		p.MediaType.Ref(),
		p.MenuItemId.Ref(),
		p.PageUrl.Ref(),
		p.ParentMenuItemId.Ref(),
		p.SelectionText.Ref(),
		p.SrcUrl.Ref(),
	)
	p.FrameUrl = p.FrameUrl.FromRef(js.Undefined)
	p.LinkUrl = p.LinkUrl.FromRef(js.Undefined)
	p.MediaType = p.MediaType.FromRef(js.Undefined)
	p.MenuItemId = p.MenuItemId.FromRef(js.Undefined)
	p.PageUrl = p.PageUrl.FromRef(js.Undefined)
	p.ParentMenuItemId = p.ParentMenuItemId.FromRef(js.Undefined)
	p.SelectionText = p.SelectionText.FromRef(js.Undefined)
	p.SrcUrl = p.SrcUrl.FromRef(js.Undefined)
}

type ItemType uint32

const (
	_ ItemType = iota

	ItemType_NORMAL
	ItemType_CHECKBOX
	ItemType_RADIO
	ItemType_SEPARATOR
)

func (ItemType) FromRef(str js.Ref) ItemType {
	return ItemType(bindings.ConstOfItemType(str))
}

func (x ItemType) String() (string, bool) {
	switch x {
	case ItemType_NORMAL:
		return "normal", true
	case ItemType_CHECKBOX:
		return "checkbox", true
	case ItemType_RADIO:
		return "radio", true
	case ItemType_SEPARATOR:
		return "separator", true
	default:
		return "", false
	}
}

type ContextType uint32

const (
	_ ContextType = iota

	ContextType_ALL
	ContextType_PAGE
	ContextType_FRAME
	ContextType_SELECTION
	ContextType_LINK
	ContextType_EDITABLE
	ContextType_IMAGE
	ContextType_VIDEO
	ContextType_AUDIO
	ContextType_LAUNCHER
	ContextType_BROWSER_ACTION
	ContextType_PAGE_ACTION
	ContextType_ACTION
)

func (ContextType) FromRef(str js.Ref) ContextType {
	return ContextType(bindings.ConstOfContextType(str))
}

func (x ContextType) String() (string, bool) {
	switch x {
	case ContextType_ALL:
		return "all", true
	case ContextType_PAGE:
		return "page", true
	case ContextType_FRAME:
		return "frame", true
	case ContextType_SELECTION:
		return "selection", true
	case ContextType_LINK:
		return "link", true
	case ContextType_EDITABLE:
		return "editable", true
	case ContextType_IMAGE:
		return "image", true
	case ContextType_VIDEO:
		return "video", true
	case ContextType_AUDIO:
		return "audio", true
	case ContextType_LAUNCHER:
		return "launcher", true
	case ContextType_BROWSER_ACTION:
		return "browser_action", true
	case ContextType_PAGE_ACTION:
		return "page_action", true
	case ContextType_ACTION:
		return "action", true
	default:
		return "", false
	}
}

type CreateArgCallbackFunc func(this js.Ref) js.Ref

func (fn CreateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateArgCallbackFunc) DispatchCallback(
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

type CreateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CreateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateArgCallback[T]) DispatchCallback(
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

type CreateArgCreatePropertiesFieldOnclickFunc func(this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref

func (fn CreateArgCreatePropertiesFieldOnclickFunc) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateArgCreatePropertiesFieldOnclickFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateArgCreatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref
	Arg T
}

func (cb *CreateArgCreatePropertiesFieldOnclick[T]) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateArgCreatePropertiesFieldOnclick[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateArgCreateProperties struct {
	// Checked is "CreateArgCreateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "CreateArgCreateProperties.contexts"
	//
	// Optional
	Contexts js.Array[ContextType]
	// DocumentUrlPatterns is "CreateArgCreateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "CreateArgCreateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Id is "CreateArgCreateProperties.id"
	//
	// Optional
	Id js.String
	// Onclick is "CreateArgCreateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func(info *OnClickData, tab *tabs.Tab)]
	// ParentId is "CreateArgCreateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "CreateArgCreateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "CreateArgCreateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "CreateArgCreateProperties.type"
	//
	// Optional
	Type ItemType
	// Visible is "CreateArgCreateProperties.visible"
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

// FromRef calls UpdateFrom and returns a CreateArgCreateProperties with all fields set.
func (p CreateArgCreateProperties) FromRef(ref js.Ref) CreateArgCreateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateArgCreateProperties in the application heap.
func (p CreateArgCreateProperties) New() js.Ref {
	return bindings.CreateArgCreatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateArgCreateProperties) UpdateFrom(ref js.Ref) {
	bindings.CreateArgCreatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateArgCreateProperties) Update(ref js.Ref) {
	bindings.CreateArgCreatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateArgCreateProperties) FreeMembers(recursive bool) {
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

type RemoveAllArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveAllArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveAllArgCallbackFunc) DispatchCallback(
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

type RemoveAllArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveAllArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveAllArgCallback[T]) DispatchCallback(
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

type RemoveArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveArgCallbackFunc) DispatchCallback(
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

type RemoveArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveArgCallback[T]) DispatchCallback(
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

type UpdateArgCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateArgCallbackFunc) DispatchCallback(
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

type UpdateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateArgCallback[T]) DispatchCallback(
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

type UpdateArgUpdatePropertiesFieldOnclickFunc func(this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref

func (fn UpdateArgUpdatePropertiesFieldOnclickFunc) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateArgUpdatePropertiesFieldOnclickFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateArgUpdatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref
	Arg T
}

func (cb *UpdateArgUpdatePropertiesFieldOnclick[T]) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateArgUpdatePropertiesFieldOnclick[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateArgUpdateProperties struct {
	// Checked is "UpdateArgUpdateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "UpdateArgUpdateProperties.contexts"
	//
	// Optional
	Contexts js.Array[ContextType]
	// DocumentUrlPatterns is "UpdateArgUpdateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "UpdateArgUpdateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Onclick is "UpdateArgUpdateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func(info *OnClickData, tab *tabs.Tab)]
	// ParentId is "UpdateArgUpdateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "UpdateArgUpdateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "UpdateArgUpdateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "UpdateArgUpdateProperties.type"
	//
	// Optional
	Type ItemType
	// Visible is "UpdateArgUpdateProperties.visible"
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

// HasFuncCreate returns true if the function "WEBEXT.contextMenus.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.contextMenus.create".
func FuncCreate() (fn js.Func[func(createProperties CreateArgCreateProperties, callback js.Func[func()]) OneOf_Int64_String]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.contextMenus.create" directly.
func Create(createProperties CreateArgCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

// TryCreate calls the function "WEBEXT.contextMenus.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(createProperties CreateArgCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

type OnClickedEventCallbackFunc func(this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref

func (fn OnClickedEventCallbackFunc) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClickedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *OnClickData, tab *tabs.Tab) js.Ref
	Arg T
}

func (cb *OnClickedEventCallback[T]) Register() js.Func[func(info *OnClickData, tab *tabs.Tab)] {
	return js.RegisterCallback[func(info *OnClickData, tab *tabs.Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClickedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)
	var arg1 tabs.Tab
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnClicked returns true if the function "WEBEXT.contextMenus.onClicked.addListener" exists.
func HasFuncOnClicked() bool {
	return js.True == bindings.HasFuncOnClicked()
}

// FuncOnClicked returns the function "WEBEXT.contextMenus.onClicked.addListener".
func FuncOnClicked() (fn js.Func[func(callback js.Func[func(info *OnClickData, tab *tabs.Tab)])]) {
	bindings.FuncOnClicked(
		js.Pointer(&fn),
	)
	return
}

// OnClicked calls the function "WEBEXT.contextMenus.onClicked.addListener" directly.
func OnClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClicked calls the function "WEBEXT.contextMenus.onClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClicked returns true if the function "WEBEXT.contextMenus.onClicked.removeListener" exists.
func HasFuncOffClicked() bool {
	return js.True == bindings.HasFuncOffClicked()
}

// FuncOffClicked returns the function "WEBEXT.contextMenus.onClicked.removeListener".
func FuncOffClicked() (fn js.Func[func(callback js.Func[func(info *OnClickData, tab *tabs.Tab)])]) {
	bindings.FuncOffClicked(
		js.Pointer(&fn),
	)
	return
}

// OffClicked calls the function "WEBEXT.contextMenus.onClicked.removeListener" directly.
func OffClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOffClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClicked calls the function "WEBEXT.contextMenus.onClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClicked returns true if the function "WEBEXT.contextMenus.onClicked.hasListener" exists.
func HasFuncHasOnClicked() bool {
	return js.True == bindings.HasFuncHasOnClicked()
}

// FuncHasOnClicked returns the function "WEBEXT.contextMenus.onClicked.hasListener".
func FuncHasOnClicked() (fn js.Func[func(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) bool]) {
	bindings.FuncHasOnClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnClicked calls the function "WEBEXT.contextMenus.onClicked.hasListener" directly.
func HasOnClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret bool) {
	bindings.CallHasOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClicked calls the function "WEBEXT.contextMenus.onClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClicked(callback js.Func[func(info *OnClickData, tab *tabs.Tab)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.contextMenus.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.contextMenus.remove".
func FuncRemove() (fn js.Func[func(menuItemId OneOf_Int64_String, callback js.Func[func()])]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.contextMenus.remove" directly.
func Remove(menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallRemove(
		js.Pointer(&ret),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// TryRemove calls the function "WEBEXT.contextMenus.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveAll returns true if the function "WEBEXT.contextMenus.removeAll" exists.
func HasFuncRemoveAll() bool {
	return js.True == bindings.HasFuncRemoveAll()
}

// FuncRemoveAll returns the function "WEBEXT.contextMenus.removeAll".
func FuncRemoveAll() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncRemoveAll(
		js.Pointer(&fn),
	)
	return
}

// RemoveAll calls the function "WEBEXT.contextMenus.removeAll" directly.
func RemoveAll(callback js.Func[func()]) (ret js.Void) {
	bindings.CallRemoveAll(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRemoveAll calls the function "WEBEXT.contextMenus.removeAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveAll(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveAll(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.contextMenus.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.contextMenus.update".
func FuncUpdate() (fn js.Func[func(id OneOf_Int64_String, updateProperties UpdateArgUpdateProperties, callback js.Func[func()])]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.contextMenus.update" directly.
func Update(id OneOf_Int64_String, updateProperties UpdateArgUpdateProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

// TryUpdate calls the function "WEBEXT.contextMenus.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(id OneOf_Int64_String, updateProperties UpdateArgUpdateProperties, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

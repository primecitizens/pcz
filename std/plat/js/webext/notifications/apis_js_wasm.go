// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package notifications

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/notifications/bindings"
)

type ClearCallbackFunc func(this js.Ref, wasCleared bool) js.Ref

func (fn ClearCallbackFunc) Register() js.Func[func(wasCleared bool)] {
	return js.RegisterCallback[func(wasCleared bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ClearCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ClearCallback[T any] struct {
	Fn  func(arg T, this js.Ref, wasCleared bool) js.Ref
	Arg T
}

func (cb *ClearCallback[T]) Register() js.Func[func(wasCleared bool)] {
	return js.RegisterCallback[func(wasCleared bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ClearCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateCallbackFunc func(this js.Ref, notificationId js.String) js.Ref

func (fn CreateCallbackFunc) Register() js.Func[func(notificationId js.String)] {
	return js.RegisterCallback[func(notificationId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateCallbackFunc) DispatchCallback(
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

type CreateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notificationId js.String) js.Ref
	Arg T
}

func (cb *CreateCallback[T]) Register() js.Func[func(notificationId js.String)] {
	return js.RegisterCallback[func(notificationId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateCallback[T]) DispatchCallback(
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

type GetAllCallbackFunc func(this js.Ref, notifications js.Object) js.Ref

func (fn GetAllCallbackFunc) Register() js.Func[func(notifications js.Object)] {
	return js.RegisterCallback[func(notifications js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notifications js.Object) js.Ref
	Arg T
}

func (cb *GetAllCallback[T]) Register() js.Func[func(notifications js.Object)] {
	return js.RegisterCallback[func(notifications js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NotificationBitmap struct {
	// Width is "NotificationBitmap.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "NotificationBitmap.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// Data is "NotificationBitmap.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationBitmap with all fields set.
func (p NotificationBitmap) FromRef(ref js.Ref) NotificationBitmap {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationBitmap in the application heap.
func (p NotificationBitmap) New() js.Ref {
	return bindings.NotificationBitmapJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationBitmap) UpdateFrom(ref js.Ref) {
	bindings.NotificationBitmapJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationBitmap) Update(ref js.Ref) {
	bindings.NotificationBitmapJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationBitmap) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type NotificationButton struct {
	// Title is "NotificationButton.title"
	//
	// Optional
	Title js.String
	// IconUrl is "NotificationButton.iconUrl"
	//
	// Optional
	IconUrl js.String
	// IconBitmap is "NotificationButton.iconBitmap"
	//
	// Optional
	//
	// NOTE: IconBitmap.FFI_USE MUST be set to true to get IconBitmap used.
	IconBitmap NotificationBitmap

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationButton with all fields set.
func (p NotificationButton) FromRef(ref js.Ref) NotificationButton {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationButton in the application heap.
func (p NotificationButton) New() js.Ref {
	return bindings.NotificationButtonJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationButton) UpdateFrom(ref js.Ref) {
	bindings.NotificationButtonJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationButton) Update(ref js.Ref) {
	bindings.NotificationButtonJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationButton) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.IconUrl.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	if recursive {
		p.IconBitmap.FreeMembers(true)
	}
}

type NotificationItem struct {
	// Title is "NotificationItem.title"
	//
	// Optional
	Title js.String
	// Message is "NotificationItem.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationItem with all fields set.
func (p NotificationItem) FromRef(ref js.Ref) NotificationItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationItem in the application heap.
func (p NotificationItem) New() js.Ref {
	return bindings.NotificationItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationItem) UpdateFrom(ref js.Ref) {
	bindings.NotificationItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationItem) Update(ref js.Ref) {
	bindings.NotificationItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationItem) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.Message.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
}

type TemplateType uint32

const (
	_ TemplateType = iota

	TemplateType_BASIC
	TemplateType_IMAGE
	TemplateType_LIST
	TemplateType_PROGRESS
)

func (TemplateType) FromRef(str js.Ref) TemplateType {
	return TemplateType(bindings.ConstOfTemplateType(str))
}

func (x TemplateType) String() (string, bool) {
	switch x {
	case TemplateType_BASIC:
		return "basic", true
	case TemplateType_IMAGE:
		return "image", true
	case TemplateType_LIST:
		return "list", true
	case TemplateType_PROGRESS:
		return "progress", true
	default:
		return "", false
	}
}

type NotificationOptions struct {
	// Type is "NotificationOptions.type"
	//
	// Optional
	Type TemplateType
	// IconUrl is "NotificationOptions.iconUrl"
	//
	// Optional
	IconUrl js.String
	// IconBitmap is "NotificationOptions.iconBitmap"
	//
	// Optional
	//
	// NOTE: IconBitmap.FFI_USE MUST be set to true to get IconBitmap used.
	IconBitmap NotificationBitmap
	// AppIconMaskUrl is "NotificationOptions.appIconMaskUrl"
	//
	// Optional
	AppIconMaskUrl js.String
	// AppIconMaskBitmap is "NotificationOptions.appIconMaskBitmap"
	//
	// Optional
	//
	// NOTE: AppIconMaskBitmap.FFI_USE MUST be set to true to get AppIconMaskBitmap used.
	AppIconMaskBitmap NotificationBitmap
	// Title is "NotificationOptions.title"
	//
	// Optional
	Title js.String
	// Message is "NotificationOptions.message"
	//
	// Optional
	Message js.String
	// ContextMessage is "NotificationOptions.contextMessage"
	//
	// Optional
	ContextMessage js.String
	// Priority is "NotificationOptions.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// EventTime is "NotificationOptions.eventTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_EventTime MUST be set to true to make this field effective.
	EventTime float64
	// Buttons is "NotificationOptions.buttons"
	//
	// Optional
	Buttons js.Array[NotificationButton]
	// ExpandedMessage is "NotificationOptions.expandedMessage"
	//
	// Optional
	ExpandedMessage js.String
	// ImageUrl is "NotificationOptions.imageUrl"
	//
	// Optional
	ImageUrl js.String
	// ImageBitmap is "NotificationOptions.imageBitmap"
	//
	// Optional
	//
	// NOTE: ImageBitmap.FFI_USE MUST be set to true to get ImageBitmap used.
	ImageBitmap NotificationBitmap
	// Items is "NotificationOptions.items"
	//
	// Optional
	Items js.Array[NotificationItem]
	// Progress is "NotificationOptions.progress"
	//
	// Optional
	//
	// NOTE: FFI_USE_Progress MUST be set to true to make this field effective.
	Progress int32
	// IsClickable is "NotificationOptions.isClickable"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsClickable MUST be set to true to make this field effective.
	IsClickable bool
	// RequireInteraction is "NotificationOptions.requireInteraction"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequireInteraction MUST be set to true to make this field effective.
	RequireInteraction bool
	// Silent is "NotificationOptions.silent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Silent MUST be set to true to make this field effective.
	Silent bool

	FFI_USE_Priority           bool // for Priority.
	FFI_USE_EventTime          bool // for EventTime.
	FFI_USE_Progress           bool // for Progress.
	FFI_USE_IsClickable        bool // for IsClickable.
	FFI_USE_RequireInteraction bool // for RequireInteraction.
	FFI_USE_Silent             bool // for Silent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationOptions with all fields set.
func (p NotificationOptions) FromRef(ref js.Ref) NotificationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationOptions in the application heap.
func (p NotificationOptions) New() js.Ref {
	return bindings.NotificationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationOptions) UpdateFrom(ref js.Ref) {
	bindings.NotificationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationOptions) Update(ref js.Ref) {
	bindings.NotificationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.IconUrl.Ref(),
		p.AppIconMaskUrl.Ref(),
		p.Title.Ref(),
		p.Message.Ref(),
		p.ContextMessage.Ref(),
		p.Buttons.Ref(),
		p.ExpandedMessage.Ref(),
		p.ImageUrl.Ref(),
		p.Items.Ref(),
	)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	p.AppIconMaskUrl = p.AppIconMaskUrl.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
	p.ContextMessage = p.ContextMessage.FromRef(js.Undefined)
	p.Buttons = p.Buttons.FromRef(js.Undefined)
	p.ExpandedMessage = p.ExpandedMessage.FromRef(js.Undefined)
	p.ImageUrl = p.ImageUrl.FromRef(js.Undefined)
	p.Items = p.Items.FromRef(js.Undefined)
	if recursive {
		p.IconBitmap.FreeMembers(true)
		p.AppIconMaskBitmap.FreeMembers(true)
		p.ImageBitmap.FreeMembers(true)
	}
}

type PermissionLevel uint32

const (
	_ PermissionLevel = iota

	PermissionLevel_GRANTED
	PermissionLevel_DENIED
)

func (PermissionLevel) FromRef(str js.Ref) PermissionLevel {
	return PermissionLevel(bindings.ConstOfPermissionLevel(str))
}

func (x PermissionLevel) String() (string, bool) {
	switch x {
	case PermissionLevel_GRANTED:
		return "granted", true
	case PermissionLevel_DENIED:
		return "denied", true
	default:
		return "", false
	}
}

type PermissionLevelCallbackFunc func(this js.Ref, level PermissionLevel) js.Ref

func (fn PermissionLevelCallbackFunc) Register() js.Func[func(level PermissionLevel)] {
	return js.RegisterCallback[func(level PermissionLevel)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PermissionLevelCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		PermissionLevel(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PermissionLevelCallback[T any] struct {
	Fn  func(arg T, this js.Ref, level PermissionLevel) js.Ref
	Arg T
}

func (cb *PermissionLevelCallback[T]) Register() js.Func[func(level PermissionLevel)] {
	return js.RegisterCallback[func(level PermissionLevel)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PermissionLevelCallback[T]) DispatchCallback(
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

		PermissionLevel(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateCallbackFunc func(this js.Ref, wasUpdated bool) js.Ref

func (fn UpdateCallbackFunc) Register() js.Func[func(wasUpdated bool)] {
	return js.RegisterCallback[func(wasUpdated bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, wasUpdated bool) js.Ref
	Arg T
}

func (cb *UpdateCallback[T]) Register() js.Func[func(wasUpdated bool)] {
	return js.RegisterCallback[func(wasUpdated bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncClear returns true if the function "WEBEXT.notifications.clear" exists.
func HasFuncClear() bool {
	return js.True == bindings.HasFuncClear()
}

// FuncClear returns the function "WEBEXT.notifications.clear".
func FuncClear() (fn js.Func[func(notificationId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncClear(
		js.Pointer(&fn),
	)
	return
}

// Clear calls the function "WEBEXT.notifications.clear" directly.
func Clear(notificationId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallClear(
		js.Pointer(&ret),
		notificationId.Ref(),
	)

	return
}

// TryClear calls the function "WEBEXT.notifications.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClear(notificationId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClear(
		js.Pointer(&ret), js.Pointer(&exception),
		notificationId.Ref(),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.notifications.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.notifications.create".
func FuncCreate() (fn js.Func[func(notificationId js.String, options NotificationOptions) js.Promise[js.String]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.notifications.create" directly.
func Create(notificationId js.String, options NotificationOptions) (ret js.Promise[js.String]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		notificationId.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreate calls the function "WEBEXT.notifications.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(notificationId js.String, options NotificationOptions) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		notificationId.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.notifications.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.notifications.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Object]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.notifications.getAll" directly.
func GetAll() (ret js.Promise[js.Object]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.notifications.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPermissionLevel returns true if the function "WEBEXT.notifications.getPermissionLevel" exists.
func HasFuncGetPermissionLevel() bool {
	return js.True == bindings.HasFuncGetPermissionLevel()
}

// FuncGetPermissionLevel returns the function "WEBEXT.notifications.getPermissionLevel".
func FuncGetPermissionLevel() (fn js.Func[func() js.Promise[PermissionLevel]]) {
	bindings.FuncGetPermissionLevel(
		js.Pointer(&fn),
	)
	return
}

// GetPermissionLevel calls the function "WEBEXT.notifications.getPermissionLevel" directly.
func GetPermissionLevel() (ret js.Promise[PermissionLevel]) {
	bindings.CallGetPermissionLevel(
		js.Pointer(&ret),
	)

	return
}

// TryGetPermissionLevel calls the function "WEBEXT.notifications.getPermissionLevel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPermissionLevel() (ret js.Promise[PermissionLevel], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPermissionLevel(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnButtonClickedEventCallbackFunc func(this js.Ref, notificationId js.String, buttonIndex int32) js.Ref

func (fn OnButtonClickedEventCallbackFunc) Register() js.Func[func(notificationId js.String, buttonIndex int32)] {
	return js.RegisterCallback[func(notificationId js.String, buttonIndex int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnButtonClickedEventCallbackFunc) DispatchCallback(
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
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnButtonClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notificationId js.String, buttonIndex int32) js.Ref
	Arg T
}

func (cb *OnButtonClickedEventCallback[T]) Register() js.Func[func(notificationId js.String, buttonIndex int32)] {
	return js.RegisterCallback[func(notificationId js.String, buttonIndex int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnButtonClickedEventCallback[T]) DispatchCallback(
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
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnButtonClicked returns true if the function "WEBEXT.notifications.onButtonClicked.addListener" exists.
func HasFuncOnButtonClicked() bool {
	return js.True == bindings.HasFuncOnButtonClicked()
}

// FuncOnButtonClicked returns the function "WEBEXT.notifications.onButtonClicked.addListener".
func FuncOnButtonClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String, buttonIndex int32)])]) {
	bindings.FuncOnButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// OnButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.addListener" directly.
func OnButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret js.Void) {
	bindings.CallOnButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffButtonClicked returns true if the function "WEBEXT.notifications.onButtonClicked.removeListener" exists.
func HasFuncOffButtonClicked() bool {
	return js.True == bindings.HasFuncOffButtonClicked()
}

// FuncOffButtonClicked returns the function "WEBEXT.notifications.onButtonClicked.removeListener".
func FuncOffButtonClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String, buttonIndex int32)])]) {
	bindings.FuncOffButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// OffButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.removeListener" directly.
func OffButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret js.Void) {
	bindings.CallOffButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnButtonClicked returns true if the function "WEBEXT.notifications.onButtonClicked.hasListener" exists.
func HasFuncHasOnButtonClicked() bool {
	return js.True == bindings.HasFuncHasOnButtonClicked()
}

// FuncHasOnButtonClicked returns the function "WEBEXT.notifications.onButtonClicked.hasListener".
func FuncHasOnButtonClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String, buttonIndex int32)]) bool]) {
	bindings.FuncHasOnButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.hasListener" directly.
func HasOnButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret bool) {
	bindings.CallHasOnButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnButtonClicked calls the function "WEBEXT.notifications.onButtonClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnButtonClicked(callback js.Func[func(notificationId js.String, buttonIndex int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnClickedEventCallbackFunc func(this js.Ref, notificationId js.String) js.Ref

func (fn OnClickedEventCallbackFunc) Register() js.Func[func(notificationId js.String)] {
	return js.RegisterCallback[func(notificationId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClickedEventCallbackFunc) DispatchCallback(
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

type OnClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notificationId js.String) js.Ref
	Arg T
}

func (cb *OnClickedEventCallback[T]) Register() js.Func[func(notificationId js.String)] {
	return js.RegisterCallback[func(notificationId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClickedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClicked returns true if the function "WEBEXT.notifications.onClicked.addListener" exists.
func HasFuncOnClicked() bool {
	return js.True == bindings.HasFuncOnClicked()
}

// FuncOnClicked returns the function "WEBEXT.notifications.onClicked.addListener".
func FuncOnClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String)])]) {
	bindings.FuncOnClicked(
		js.Pointer(&fn),
	)
	return
}

// OnClicked calls the function "WEBEXT.notifications.onClicked.addListener" directly.
func OnClicked(callback js.Func[func(notificationId js.String)]) (ret js.Void) {
	bindings.CallOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClicked calls the function "WEBEXT.notifications.onClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClicked(callback js.Func[func(notificationId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClicked returns true if the function "WEBEXT.notifications.onClicked.removeListener" exists.
func HasFuncOffClicked() bool {
	return js.True == bindings.HasFuncOffClicked()
}

// FuncOffClicked returns the function "WEBEXT.notifications.onClicked.removeListener".
func FuncOffClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String)])]) {
	bindings.FuncOffClicked(
		js.Pointer(&fn),
	)
	return
}

// OffClicked calls the function "WEBEXT.notifications.onClicked.removeListener" directly.
func OffClicked(callback js.Func[func(notificationId js.String)]) (ret js.Void) {
	bindings.CallOffClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClicked calls the function "WEBEXT.notifications.onClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClicked(callback js.Func[func(notificationId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClicked returns true if the function "WEBEXT.notifications.onClicked.hasListener" exists.
func HasFuncHasOnClicked() bool {
	return js.True == bindings.HasFuncHasOnClicked()
}

// FuncHasOnClicked returns the function "WEBEXT.notifications.onClicked.hasListener".
func FuncHasOnClicked() (fn js.Func[func(callback js.Func[func(notificationId js.String)]) bool]) {
	bindings.FuncHasOnClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnClicked calls the function "WEBEXT.notifications.onClicked.hasListener" directly.
func HasOnClicked(callback js.Func[func(notificationId js.String)]) (ret bool) {
	bindings.CallHasOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClicked calls the function "WEBEXT.notifications.onClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClicked(callback js.Func[func(notificationId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnClosedEventCallbackFunc func(this js.Ref, notificationId js.String, byUser bool) js.Ref

func (fn OnClosedEventCallbackFunc) Register() js.Func[func(notificationId js.String, byUser bool)] {
	return js.RegisterCallback[func(notificationId js.String, byUser bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClosedEventCallbackFunc) DispatchCallback(
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

type OnClosedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notificationId js.String, byUser bool) js.Ref
	Arg T
}

func (cb *OnClosedEventCallback[T]) Register() js.Func[func(notificationId js.String, byUser bool)] {
	return js.RegisterCallback[func(notificationId js.String, byUser bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClosedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClosed returns true if the function "WEBEXT.notifications.onClosed.addListener" exists.
func HasFuncOnClosed() bool {
	return js.True == bindings.HasFuncOnClosed()
}

// FuncOnClosed returns the function "WEBEXT.notifications.onClosed.addListener".
func FuncOnClosed() (fn js.Func[func(callback js.Func[func(notificationId js.String, byUser bool)])]) {
	bindings.FuncOnClosed(
		js.Pointer(&fn),
	)
	return
}

// OnClosed calls the function "WEBEXT.notifications.onClosed.addListener" directly.
func OnClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret js.Void) {
	bindings.CallOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClosed calls the function "WEBEXT.notifications.onClosed.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClosed returns true if the function "WEBEXT.notifications.onClosed.removeListener" exists.
func HasFuncOffClosed() bool {
	return js.True == bindings.HasFuncOffClosed()
}

// FuncOffClosed returns the function "WEBEXT.notifications.onClosed.removeListener".
func FuncOffClosed() (fn js.Func[func(callback js.Func[func(notificationId js.String, byUser bool)])]) {
	bindings.FuncOffClosed(
		js.Pointer(&fn),
	)
	return
}

// OffClosed calls the function "WEBEXT.notifications.onClosed.removeListener" directly.
func OffClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret js.Void) {
	bindings.CallOffClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClosed calls the function "WEBEXT.notifications.onClosed.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClosed returns true if the function "WEBEXT.notifications.onClosed.hasListener" exists.
func HasFuncHasOnClosed() bool {
	return js.True == bindings.HasFuncHasOnClosed()
}

// FuncHasOnClosed returns the function "WEBEXT.notifications.onClosed.hasListener".
func FuncHasOnClosed() (fn js.Func[func(callback js.Func[func(notificationId js.String, byUser bool)]) bool]) {
	bindings.FuncHasOnClosed(
		js.Pointer(&fn),
	)
	return
}

// HasOnClosed calls the function "WEBEXT.notifications.onClosed.hasListener" directly.
func HasOnClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret bool) {
	bindings.CallHasOnClosed(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClosed calls the function "WEBEXT.notifications.onClosed.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClosed(callback js.Func[func(notificationId js.String, byUser bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPermissionLevelChangedEventCallbackFunc func(this js.Ref, level PermissionLevel) js.Ref

func (fn OnPermissionLevelChangedEventCallbackFunc) Register() js.Func[func(level PermissionLevel)] {
	return js.RegisterCallback[func(level PermissionLevel)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPermissionLevelChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		PermissionLevel(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPermissionLevelChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, level PermissionLevel) js.Ref
	Arg T
}

func (cb *OnPermissionLevelChangedEventCallback[T]) Register() js.Func[func(level PermissionLevel)] {
	return js.RegisterCallback[func(level PermissionLevel)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPermissionLevelChangedEventCallback[T]) DispatchCallback(
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

		PermissionLevel(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPermissionLevelChanged returns true if the function "WEBEXT.notifications.onPermissionLevelChanged.addListener" exists.
func HasFuncOnPermissionLevelChanged() bool {
	return js.True == bindings.HasFuncOnPermissionLevelChanged()
}

// FuncOnPermissionLevelChanged returns the function "WEBEXT.notifications.onPermissionLevelChanged.addListener".
func FuncOnPermissionLevelChanged() (fn js.Func[func(callback js.Func[func(level PermissionLevel)])]) {
	bindings.FuncOnPermissionLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.addListener" directly.
func OnPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret js.Void) {
	bindings.CallOnPermissionLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPermissionLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPermissionLevelChanged returns true if the function "WEBEXT.notifications.onPermissionLevelChanged.removeListener" exists.
func HasFuncOffPermissionLevelChanged() bool {
	return js.True == bindings.HasFuncOffPermissionLevelChanged()
}

// FuncOffPermissionLevelChanged returns the function "WEBEXT.notifications.onPermissionLevelChanged.removeListener".
func FuncOffPermissionLevelChanged() (fn js.Func[func(callback js.Func[func(level PermissionLevel)])]) {
	bindings.FuncOffPermissionLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.removeListener" directly.
func OffPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret js.Void) {
	bindings.CallOffPermissionLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPermissionLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPermissionLevelChanged returns true if the function "WEBEXT.notifications.onPermissionLevelChanged.hasListener" exists.
func HasFuncHasOnPermissionLevelChanged() bool {
	return js.True == bindings.HasFuncHasOnPermissionLevelChanged()
}

// FuncHasOnPermissionLevelChanged returns the function "WEBEXT.notifications.onPermissionLevelChanged.hasListener".
func FuncHasOnPermissionLevelChanged() (fn js.Func[func(callback js.Func[func(level PermissionLevel)]) bool]) {
	bindings.FuncHasOnPermissionLevelChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.hasListener" directly.
func HasOnPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret bool) {
	bindings.CallHasOnPermissionLevelChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPermissionLevelChanged calls the function "WEBEXT.notifications.onPermissionLevelChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPermissionLevelChanged(callback js.Func[func(level PermissionLevel)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPermissionLevelChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnShowSettingsEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnShowSettingsEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnShowSettingsEventCallbackFunc) DispatchCallback(
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

type OnShowSettingsEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnShowSettingsEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnShowSettingsEventCallback[T]) DispatchCallback(
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

// HasFuncOnShowSettings returns true if the function "WEBEXT.notifications.onShowSettings.addListener" exists.
func HasFuncOnShowSettings() bool {
	return js.True == bindings.HasFuncOnShowSettings()
}

// FuncOnShowSettings returns the function "WEBEXT.notifications.onShowSettings.addListener".
func FuncOnShowSettings() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnShowSettings(
		js.Pointer(&fn),
	)
	return
}

// OnShowSettings calls the function "WEBEXT.notifications.onShowSettings.addListener" directly.
func OnShowSettings(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnShowSettings(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnShowSettings calls the function "WEBEXT.notifications.onShowSettings.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnShowSettings(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnShowSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffShowSettings returns true if the function "WEBEXT.notifications.onShowSettings.removeListener" exists.
func HasFuncOffShowSettings() bool {
	return js.True == bindings.HasFuncOffShowSettings()
}

// FuncOffShowSettings returns the function "WEBEXT.notifications.onShowSettings.removeListener".
func FuncOffShowSettings() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffShowSettings(
		js.Pointer(&fn),
	)
	return
}

// OffShowSettings calls the function "WEBEXT.notifications.onShowSettings.removeListener" directly.
func OffShowSettings(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffShowSettings(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffShowSettings calls the function "WEBEXT.notifications.onShowSettings.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffShowSettings(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffShowSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnShowSettings returns true if the function "WEBEXT.notifications.onShowSettings.hasListener" exists.
func HasFuncHasOnShowSettings() bool {
	return js.True == bindings.HasFuncHasOnShowSettings()
}

// FuncHasOnShowSettings returns the function "WEBEXT.notifications.onShowSettings.hasListener".
func FuncHasOnShowSettings() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnShowSettings(
		js.Pointer(&fn),
	)
	return
}

// HasOnShowSettings calls the function "WEBEXT.notifications.onShowSettings.hasListener" directly.
func HasOnShowSettings(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnShowSettings(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnShowSettings calls the function "WEBEXT.notifications.onShowSettings.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnShowSettings(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnShowSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.notifications.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.notifications.update".
func FuncUpdate() (fn js.Func[func(notificationId js.String, options NotificationOptions) js.Promise[js.Boolean]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.notifications.update" directly.
func Update(notificationId js.String, options NotificationOptions) (ret js.Promise[js.Boolean]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		notificationId.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryUpdate calls the function "WEBEXT.notifications.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(notificationId js.String, options NotificationOptions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		notificationId.Ref(),
		js.Pointer(&options),
	)

	return
}

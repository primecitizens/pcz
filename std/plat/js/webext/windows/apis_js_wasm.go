// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package windows

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
	"github.com/primecitizens/pcz/std/plat/js/webext/windows/bindings"
)

type WindowState uint32

const (
	_ WindowState = iota

	WindowState_NORMAL
	WindowState_MINIMIZED
	WindowState_MAXIMIZED
	WindowState_FULLSCREEN
	WindowState_LOCKED_FULLSCREEN
)

func (WindowState) FromRef(str js.Ref) WindowState {
	return WindowState(bindings.ConstOfWindowState(str))
}

func (x WindowState) String() (string, bool) {
	switch x {
	case WindowState_NORMAL:
		return "normal", true
	case WindowState_MINIMIZED:
		return "minimized", true
	case WindowState_MAXIMIZED:
		return "maximized", true
	case WindowState_FULLSCREEN:
		return "fullscreen", true
	case WindowState_LOCKED_FULLSCREEN:
		return "locked-fullscreen", true
	default:
		return "", false
	}
}

type CreateType uint32

const (
	_ CreateType = iota

	CreateType_NORMAL
	CreateType_POPUP
	CreateType_PANEL
)

func (CreateType) FromRef(str js.Ref) CreateType {
	return CreateType(bindings.ConstOfCreateType(str))
}

func (x CreateType) String() (string, bool) {
	switch x {
	case CreateType_NORMAL:
		return "normal", true
	case CreateType_POPUP:
		return "popup", true
	case CreateType_PANEL:
		return "panel", true
	default:
		return "", false
	}
}

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type CreateArgCreateData struct {
	// Focused is "CreateArgCreateData.focused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Focused MUST be set to true to make this field effective.
	Focused bool
	// Height is "CreateArgCreateData.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int64
	// Incognito is "CreateArgCreateData.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// Left is "CreateArgCreateData.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int64
	// SetSelfAsOpener is "CreateArgCreateData.setSelfAsOpener"
	//
	// Optional
	//
	// NOTE: FFI_USE_SetSelfAsOpener MUST be set to true to make this field effective.
	SetSelfAsOpener bool
	// State is "CreateArgCreateData.state"
	//
	// Optional
	State WindowState
	// TabId is "CreateArgCreateData.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Top is "CreateArgCreateData.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int64
	// Type is "CreateArgCreateData.type"
	//
	// Optional
	Type CreateType
	// Url is "CreateArgCreateData.url"
	//
	// Optional
	Url OneOf_String_ArrayString
	// Width is "CreateArgCreateData.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int64

	FFI_USE_Focused         bool // for Focused.
	FFI_USE_Height          bool // for Height.
	FFI_USE_Incognito       bool // for Incognito.
	FFI_USE_Left            bool // for Left.
	FFI_USE_SetSelfAsOpener bool // for SetSelfAsOpener.
	FFI_USE_TabId           bool // for TabId.
	FFI_USE_Top             bool // for Top.
	FFI_USE_Width           bool // for Width.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateArgCreateData with all fields set.
func (p CreateArgCreateData) FromRef(ref js.Ref) CreateArgCreateData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateArgCreateData in the application heap.
func (p CreateArgCreateData) New() js.Ref {
	return bindings.CreateArgCreateDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateArgCreateData) UpdateFrom(ref js.Ref) {
	bindings.CreateArgCreateDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateArgCreateData) Update(ref js.Ref) {
	bindings.CreateArgCreateDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateArgCreateData) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type WindowType uint32

const (
	_ WindowType = iota

	WindowType_NORMAL
	WindowType_POPUP
	WindowType_PANEL
	WindowType_APP
	WindowType_DEVTOOLS
)

func (WindowType) FromRef(str js.Ref) WindowType {
	return WindowType(bindings.ConstOfWindowType(str))
}

func (x WindowType) String() (string, bool) {
	switch x {
	case WindowType_NORMAL:
		return "normal", true
	case WindowType_POPUP:
		return "popup", true
	case WindowType_PANEL:
		return "panel", true
	case WindowType_APP:
		return "app", true
	case WindowType_DEVTOOLS:
		return "devtools", true
	default:
		return "", false
	}
}

type Window struct {
	// AlwaysOnTop is "Window.alwaysOnTop"
	//
	// Required
	AlwaysOnTop bool
	// Focused is "Window.focused"
	//
	// Required
	Focused bool
	// Height is "Window.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int64
	// Id is "Window.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int64
	// Incognito is "Window.incognito"
	//
	// Required
	Incognito bool
	// Left is "Window.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int64
	// SessionId is "Window.sessionId"
	//
	// Optional
	SessionId js.String
	// State is "Window.state"
	//
	// Optional
	State WindowState
	// Tabs is "Window.tabs"
	//
	// Optional
	Tabs js.Array[tabs.Tab]
	// Top is "Window.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int64
	// Type is "Window.type"
	//
	// Optional
	Type WindowType
	// Width is "Window.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int64

	FFI_USE_Height bool // for Height.
	FFI_USE_Id     bool // for Id.
	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Window with all fields set.
func (p Window) FromRef(ref js.Ref) Window {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Window in the application heap.
func (p Window) New() js.Ref {
	return bindings.WindowJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Window) UpdateFrom(ref js.Ref) {
	bindings.WindowJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Window) Update(ref js.Ref) {
	bindings.WindowJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Window) FreeMembers(recursive bool) {
	js.Free(
		p.SessionId.Ref(),
		p.Tabs.Ref(),
	)
	p.SessionId = p.SessionId.FromRef(js.Undefined)
	p.Tabs = p.Tabs.FromRef(js.Undefined)
}

type QueryOptions struct {
	// Populate is "QueryOptions.populate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Populate MUST be set to true to make this field effective.
	Populate bool
	// WindowTypes is "QueryOptions.windowTypes"
	//
	// Optional
	WindowTypes js.Array[WindowType]

	FFI_USE_Populate bool // for Populate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryOptions with all fields set.
func (p QueryOptions) FromRef(ref js.Ref) QueryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryOptions in the application heap.
func (p QueryOptions) New() js.Ref {
	return bindings.QueryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryOptions) UpdateFrom(ref js.Ref) {
	bindings.QueryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryOptions) Update(ref js.Ref) {
	bindings.QueryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.WindowTypes.Ref(),
	)
	p.WindowTypes = p.WindowTypes.FromRef(js.Undefined)
}

type UpdateArgUpdateInfo struct {
	// DrawAttention is "UpdateArgUpdateInfo.drawAttention"
	//
	// Optional
	//
	// NOTE: FFI_USE_DrawAttention MUST be set to true to make this field effective.
	DrawAttention bool
	// Focused is "UpdateArgUpdateInfo.focused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Focused MUST be set to true to make this field effective.
	Focused bool
	// Height is "UpdateArgUpdateInfo.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int64
	// Left is "UpdateArgUpdateInfo.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int64
	// State is "UpdateArgUpdateInfo.state"
	//
	// Optional
	State WindowState
	// Top is "UpdateArgUpdateInfo.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int64
	// Width is "UpdateArgUpdateInfo.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int64

	FFI_USE_DrawAttention bool // for DrawAttention.
	FFI_USE_Focused       bool // for Focused.
	FFI_USE_Height        bool // for Height.
	FFI_USE_Left          bool // for Left.
	FFI_USE_Top           bool // for Top.
	FFI_USE_Width         bool // for Width.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateArgUpdateInfo with all fields set.
func (p UpdateArgUpdateInfo) FromRef(ref js.Ref) UpdateArgUpdateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateArgUpdateInfo in the application heap.
func (p UpdateArgUpdateInfo) New() js.Ref {
	return bindings.UpdateArgUpdateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateArgUpdateInfo) UpdateFrom(ref js.Ref) {
	bindings.UpdateArgUpdateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateArgUpdateInfo) Update(ref js.Ref) {
	bindings.UpdateArgUpdateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateArgUpdateInfo) FreeMembers(recursive bool) {
}

// WINDOW_ID_CURRENT returns the value of property "WEBEXT.windows.WINDOW_ID_CURRENT".
//
// The returned bool will be false if there is no such property.
func WINDOW_ID_CURRENT() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWINDOW_ID_CURRENT(
		js.Pointer(&ret),
	)

	return
}

// SetWINDOW_ID_CURRENT sets the value of property "WEBEXT.windows.WINDOW_ID_CURRENT" to val.
//
// It returns false if the property cannot be set.
func SetWINDOW_ID_CURRENT(val js.String) bool {
	return js.True == bindings.SetWINDOW_ID_CURRENT(
		val.Ref())
}

// WINDOW_ID_NONE returns the value of property "WEBEXT.windows.WINDOW_ID_NONE".
//
// The returned bool will be false if there is no such property.
func WINDOW_ID_NONE() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWINDOW_ID_NONE(
		js.Pointer(&ret),
	)

	return
}

// SetWINDOW_ID_NONE sets the value of property "WEBEXT.windows.WINDOW_ID_NONE" to val.
//
// It returns false if the property cannot be set.
func SetWINDOW_ID_NONE(val js.String) bool {
	return js.True == bindings.SetWINDOW_ID_NONE(
		val.Ref())
}

// HasFuncCreate returns true if the function "WEBEXT.windows.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.windows.create".
func FuncCreate() (fn js.Func[func(createData CreateArgCreateData) js.Promise[Window]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.windows.create" directly.
func Create(createData CreateArgCreateData) (ret js.Promise[Window]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&createData),
	)

	return
}

// TryCreate calls the function "WEBEXT.windows.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(createData CreateArgCreateData) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&createData),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.windows.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.windows.get".
func FuncGet() (fn js.Func[func(windowId int64, queryOptions QueryOptions) js.Promise[Window]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.windows.get" directly.
func Get(windowId int64, queryOptions QueryOptions) (ret js.Promise[Window]) {
	bindings.CallGet(
		js.Pointer(&ret),
		float64(windowId),
		js.Pointer(&queryOptions),
	)

	return
}

// TryGet calls the function "WEBEXT.windows.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(windowId int64, queryOptions QueryOptions) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
		js.Pointer(&queryOptions),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.windows.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.windows.getAll".
func FuncGetAll() (fn js.Func[func(queryOptions QueryOptions) js.Promise[js.Array[Window]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.windows.getAll" directly.
func GetAll(queryOptions QueryOptions) (ret js.Promise[js.Array[Window]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
		js.Pointer(&queryOptions),
	)

	return
}

// TryGetAll calls the function "WEBEXT.windows.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll(queryOptions QueryOptions) (ret js.Promise[js.Array[Window]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryOptions),
	)

	return
}

// HasFuncGetCurrent returns true if the function "WEBEXT.windows.getCurrent" exists.
func HasFuncGetCurrent() bool {
	return js.True == bindings.HasFuncGetCurrent()
}

// FuncGetCurrent returns the function "WEBEXT.windows.getCurrent".
func FuncGetCurrent() (fn js.Func[func(queryOptions QueryOptions) js.Promise[Window]]) {
	bindings.FuncGetCurrent(
		js.Pointer(&fn),
	)
	return
}

// GetCurrent calls the function "WEBEXT.windows.getCurrent" directly.
func GetCurrent(queryOptions QueryOptions) (ret js.Promise[Window]) {
	bindings.CallGetCurrent(
		js.Pointer(&ret),
		js.Pointer(&queryOptions),
	)

	return
}

// TryGetCurrent calls the function "WEBEXT.windows.getCurrent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCurrent(queryOptions QueryOptions) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCurrent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryOptions),
	)

	return
}

// HasFuncGetLastFocused returns true if the function "WEBEXT.windows.getLastFocused" exists.
func HasFuncGetLastFocused() bool {
	return js.True == bindings.HasFuncGetLastFocused()
}

// FuncGetLastFocused returns the function "WEBEXT.windows.getLastFocused".
func FuncGetLastFocused() (fn js.Func[func(queryOptions QueryOptions) js.Promise[Window]]) {
	bindings.FuncGetLastFocused(
		js.Pointer(&fn),
	)
	return
}

// GetLastFocused calls the function "WEBEXT.windows.getLastFocused" directly.
func GetLastFocused(queryOptions QueryOptions) (ret js.Promise[Window]) {
	bindings.CallGetLastFocused(
		js.Pointer(&ret),
		js.Pointer(&queryOptions),
	)

	return
}

// TryGetLastFocused calls the function "WEBEXT.windows.getLastFocused"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLastFocused(queryOptions QueryOptions) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLastFocused(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryOptions),
	)

	return
}

type OnBoundsChangedEventCallbackFunc func(this js.Ref, window *Window) js.Ref

func (fn OnBoundsChangedEventCallbackFunc) Register() js.Func[func(window *Window)] {
	return js.RegisterCallback[func(window *Window)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBoundsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Window
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

type OnBoundsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, window *Window) js.Ref
	Arg T
}

func (cb *OnBoundsChangedEventCallback[T]) Register() js.Func[func(window *Window)] {
	return js.RegisterCallback[func(window *Window)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBoundsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Window
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

// HasFuncOnBoundsChanged returns true if the function "WEBEXT.windows.onBoundsChanged.addListener" exists.
func HasFuncOnBoundsChanged() bool {
	return js.True == bindings.HasFuncOnBoundsChanged()
}

// FuncOnBoundsChanged returns the function "WEBEXT.windows.onBoundsChanged.addListener".
func FuncOnBoundsChanged() (fn js.Func[func(callback js.Func[func(window *Window)])]) {
	bindings.FuncOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.addListener" directly.
func OnBoundsChanged(callback js.Func[func(window *Window)]) (ret js.Void) {
	bindings.CallOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBoundsChanged(callback js.Func[func(window *Window)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBoundsChanged returns true if the function "WEBEXT.windows.onBoundsChanged.removeListener" exists.
func HasFuncOffBoundsChanged() bool {
	return js.True == bindings.HasFuncOffBoundsChanged()
}

// FuncOffBoundsChanged returns the function "WEBEXT.windows.onBoundsChanged.removeListener".
func FuncOffBoundsChanged() (fn js.Func[func(callback js.Func[func(window *Window)])]) {
	bindings.FuncOffBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.removeListener" directly.
func OffBoundsChanged(callback js.Func[func(window *Window)]) (ret js.Void) {
	bindings.CallOffBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBoundsChanged(callback js.Func[func(window *Window)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBoundsChanged returns true if the function "WEBEXT.windows.onBoundsChanged.hasListener" exists.
func HasFuncHasOnBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnBoundsChanged()
}

// FuncHasOnBoundsChanged returns the function "WEBEXT.windows.onBoundsChanged.hasListener".
func FuncHasOnBoundsChanged() (fn js.Func[func(callback js.Func[func(window *Window)]) bool]) {
	bindings.FuncHasOnBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.hasListener" directly.
func HasOnBoundsChanged(callback js.Func[func(window *Window)]) (ret bool) {
	bindings.CallHasOnBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBoundsChanged calls the function "WEBEXT.windows.onBoundsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBoundsChanged(callback js.Func[func(window *Window)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, window *Window) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(window *Window)] {
	return js.RegisterCallback[func(window *Window)](
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
	var arg0 Window
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
	Fn  func(arg T, this js.Ref, window *Window) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(window *Window)] {
	return js.RegisterCallback[func(window *Window)](
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
	var arg0 Window
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

// HasFuncOnCreated returns true if the function "WEBEXT.windows.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.windows.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(window *Window)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.windows.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(window *Window)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.windows.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(window *Window)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.windows.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.windows.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(window *Window)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.windows.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(window *Window)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.windows.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(window *Window)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.windows.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.windows.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(window *Window)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.windows.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(window *Window)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.windows.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(window *Window)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFocusChangedEventCallbackFunc func(this js.Ref, windowId int64) js.Ref

func (fn OnFocusChangedEventCallbackFunc) Register() js.Func[func(windowId int64)] {
	return js.RegisterCallback[func(windowId int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFocusChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnFocusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, windowId int64) js.Ref
	Arg T
}

func (cb *OnFocusChangedEventCallback[T]) Register() js.Func[func(windowId int64)] {
	return js.RegisterCallback[func(windowId int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFocusChangedEventCallback[T]) DispatchCallback(
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

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnFocusChanged returns true if the function "WEBEXT.windows.onFocusChanged.addListener" exists.
func HasFuncOnFocusChanged() bool {
	return js.True == bindings.HasFuncOnFocusChanged()
}

// FuncOnFocusChanged returns the function "WEBEXT.windows.onFocusChanged.addListener".
func FuncOnFocusChanged() (fn js.Func[func(callback js.Func[func(windowId int64)])]) {
	bindings.FuncOnFocusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnFocusChanged calls the function "WEBEXT.windows.onFocusChanged.addListener" directly.
func OnFocusChanged(callback js.Func[func(windowId int64)]) (ret js.Void) {
	bindings.CallOnFocusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFocusChanged calls the function "WEBEXT.windows.onFocusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFocusChanged(callback js.Func[func(windowId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFocusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFocusChanged returns true if the function "WEBEXT.windows.onFocusChanged.removeListener" exists.
func HasFuncOffFocusChanged() bool {
	return js.True == bindings.HasFuncOffFocusChanged()
}

// FuncOffFocusChanged returns the function "WEBEXT.windows.onFocusChanged.removeListener".
func FuncOffFocusChanged() (fn js.Func[func(callback js.Func[func(windowId int64)])]) {
	bindings.FuncOffFocusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffFocusChanged calls the function "WEBEXT.windows.onFocusChanged.removeListener" directly.
func OffFocusChanged(callback js.Func[func(windowId int64)]) (ret js.Void) {
	bindings.CallOffFocusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFocusChanged calls the function "WEBEXT.windows.onFocusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFocusChanged(callback js.Func[func(windowId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFocusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFocusChanged returns true if the function "WEBEXT.windows.onFocusChanged.hasListener" exists.
func HasFuncHasOnFocusChanged() bool {
	return js.True == bindings.HasFuncHasOnFocusChanged()
}

// FuncHasOnFocusChanged returns the function "WEBEXT.windows.onFocusChanged.hasListener".
func FuncHasOnFocusChanged() (fn js.Func[func(callback js.Func[func(windowId int64)]) bool]) {
	bindings.FuncHasOnFocusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnFocusChanged calls the function "WEBEXT.windows.onFocusChanged.hasListener" directly.
func HasOnFocusChanged(callback js.Func[func(windowId int64)]) (ret bool) {
	bindings.CallHasOnFocusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFocusChanged calls the function "WEBEXT.windows.onFocusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFocusChanged(callback js.Func[func(windowId int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFocusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemovedEventCallbackFunc func(this js.Ref, windowId int64) js.Ref

func (fn OnRemovedEventCallbackFunc) Register() js.Func[func(windowId int64)] {
	return js.RegisterCallback[func(windowId int64)](
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

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, windowId int64) js.Ref
	Arg T
}

func (cb *OnRemovedEventCallback[T]) Register() js.Func[func(windowId int64)] {
	return js.RegisterCallback[func(windowId int64)](
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

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRemoved returns true if the function "WEBEXT.windows.onRemoved.addListener" exists.
func HasFuncOnRemoved() bool {
	return js.True == bindings.HasFuncOnRemoved()
}

// FuncOnRemoved returns the function "WEBEXT.windows.onRemoved.addListener".
func FuncOnRemoved() (fn js.Func[func(callback js.Func[func(windowId int64)])]) {
	bindings.FuncOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnRemoved calls the function "WEBEXT.windows.onRemoved.addListener" directly.
func OnRemoved(callback js.Func[func(windowId int64)]) (ret js.Void) {
	bindings.CallOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoved calls the function "WEBEXT.windows.onRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoved(callback js.Func[func(windowId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoved returns true if the function "WEBEXT.windows.onRemoved.removeListener" exists.
func HasFuncOffRemoved() bool {
	return js.True == bindings.HasFuncOffRemoved()
}

// FuncOffRemoved returns the function "WEBEXT.windows.onRemoved.removeListener".
func FuncOffRemoved() (fn js.Func[func(callback js.Func[func(windowId int64)])]) {
	bindings.FuncOffRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffRemoved calls the function "WEBEXT.windows.onRemoved.removeListener" directly.
func OffRemoved(callback js.Func[func(windowId int64)]) (ret js.Void) {
	bindings.CallOffRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoved calls the function "WEBEXT.windows.onRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoved(callback js.Func[func(windowId int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoved returns true if the function "WEBEXT.windows.onRemoved.hasListener" exists.
func HasFuncHasOnRemoved() bool {
	return js.True == bindings.HasFuncHasOnRemoved()
}

// FuncHasOnRemoved returns the function "WEBEXT.windows.onRemoved.hasListener".
func FuncHasOnRemoved() (fn js.Func[func(callback js.Func[func(windowId int64)]) bool]) {
	bindings.FuncHasOnRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoved calls the function "WEBEXT.windows.onRemoved.hasListener" directly.
func HasOnRemoved(callback js.Func[func(windowId int64)]) (ret bool) {
	bindings.CallHasOnRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoved calls the function "WEBEXT.windows.onRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoved(callback js.Func[func(windowId int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.windows.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.windows.remove".
func FuncRemove() (fn js.Func[func(windowId int64) js.Promise[js.Void]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.windows.remove" directly.
func Remove(windowId int64) (ret js.Promise[js.Void]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		float64(windowId),
	)

	return
}

// TryRemove calls the function "WEBEXT.windows.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(windowId int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.windows.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.windows.update".
func FuncUpdate() (fn js.Func[func(windowId int64, updateInfo UpdateArgUpdateInfo) js.Promise[Window]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.windows.update" directly.
func Update(windowId int64, updateInfo UpdateArgUpdateInfo) (ret js.Promise[Window]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		float64(windowId),
		js.Pointer(&updateInfo),
	)

	return
}

// TryUpdate calls the function "WEBEXT.windows.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(windowId int64, updateInfo UpdateArgUpdateInfo) (ret js.Promise[Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
		js.Pointer(&updateInfo),
	)

	return
}

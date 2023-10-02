// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type OneOf_Client_undefined struct {
	ref js.Ref
}

func (x OneOf_Client_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Client_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_Client_undefined) FromRef(ref js.Ref) OneOf_Client_undefined {
	return OneOf_Client_undefined{
		ref: ref,
	}
}

func (x OneOf_Client_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_Client_undefined) Client() Client {
	return Client{}.FromRef(x.ref)
}

type WindowClient struct {
	Client
}

func (this WindowClient) Once() WindowClient {
	this.Ref().Once()
	return this
}

func (this WindowClient) Ref() js.Ref {
	return this.Client.Ref()
}

func (this WindowClient) FromRef(ref js.Ref) WindowClient {
	this.Client = this.Client.FromRef(ref)
	return this
}

func (this WindowClient) Free() {
	this.Ref().Free()
}

// VisibilityState returns the value of property "WindowClient.visibilityState".
//
// The returned bool will be false if there is no such property.
func (this WindowClient) VisibilityState() (DocumentVisibilityState, bool) {
	var _ok bool
	_ret := bindings.GetWindowClientVisibilityState(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentVisibilityState(_ret), _ok
}

// Focused returns the value of property "WindowClient.focused".
//
// The returned bool will be false if there is no such property.
func (this WindowClient) Focused() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowClientFocused(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AncestorOrigins returns the value of property "WindowClient.ancestorOrigins".
//
// The returned bool will be false if there is no such property.
func (this WindowClient) AncestorOrigins() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetWindowClientAncestorOrigins(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Focus calls the method "WindowClient.focus".
//
// The returned bool will be false if there is no such method.
func (this WindowClient) Focus() (js.Promise[WindowClient], bool) {
	var _ok bool
	_ret := bindings.CallWindowClientFocus(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[WindowClient]{}.FromRef(_ret), _ok
}

// FocusFunc returns the method "WindowClient.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowClient) FocusFunc() (fn js.Func[func() js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.WindowClientFocusFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "WindowClient.navigate".
//
// The returned bool will be false if there is no such method.
func (this WindowClient) Navigate(url js.String) (js.Promise[WindowClient], bool) {
	var _ok bool
	_ret := bindings.CallWindowClientNavigate(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return js.Promise[WindowClient]{}.FromRef(_ret), _ok
}

// NavigateFunc returns the method "WindowClient.navigate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WindowClient) NavigateFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.WindowClientNavigateFunc(
			this.Ref(),
		),
	)
}

type Clients struct {
	ref js.Ref
}

func (this Clients) Once() Clients {
	this.Ref().Once()
	return this
}

func (this Clients) Ref() js.Ref {
	return this.ref
}

func (this Clients) FromRef(ref js.Ref) Clients {
	this.ref = ref
	return this
}

func (this Clients) Free() {
	this.Ref().Free()
}

// Get calls the method "Clients.get".
//
// The returned bool will be false if there is no such method.
func (this Clients) Get(id js.String) (js.Promise[OneOf_Client_undefined], bool) {
	var _ok bool
	_ret := bindings.CallClientsGet(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return js.Promise[OneOf_Client_undefined]{}.FromRef(_ret), _ok
}

// GetFunc returns the method "Clients.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clients) GetFunc() (fn js.Func[func(id js.String) js.Promise[OneOf_Client_undefined]]) {
	return fn.FromRef(
		bindings.ClientsGetFunc(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "Clients.matchAll".
//
// The returned bool will be false if there is no such method.
func (this Clients) MatchAll(options ClientQueryOptions) (js.Promise[js.FrozenArray[Client]], bool) {
	var _ok bool
	_ret := bindings.CallClientsMatchAll(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.FrozenArray[Client]]{}.FromRef(_ret), _ok
}

// MatchAllFunc returns the method "Clients.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clients) MatchAllFunc() (fn js.Func[func(options ClientQueryOptions) js.Promise[js.FrozenArray[Client]]]) {
	return fn.FromRef(
		bindings.ClientsMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "Clients.matchAll".
//
// The returned bool will be false if there is no such method.
func (this Clients) MatchAll1() (js.Promise[js.FrozenArray[Client]], bool) {
	var _ok bool
	_ret := bindings.CallClientsMatchAll1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.FrozenArray[Client]]{}.FromRef(_ret), _ok
}

// MatchAll1Func returns the method "Clients.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clients) MatchAll1Func() (fn js.Func[func() js.Promise[js.FrozenArray[Client]]]) {
	return fn.FromRef(
		bindings.ClientsMatchAll1Func(
			this.Ref(),
		),
	)
}

// OpenWindow calls the method "Clients.openWindow".
//
// The returned bool will be false if there is no such method.
func (this Clients) OpenWindow(url js.String) (js.Promise[WindowClient], bool) {
	var _ok bool
	_ret := bindings.CallClientsOpenWindow(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return js.Promise[WindowClient]{}.FromRef(_ret), _ok
}

// OpenWindowFunc returns the method "Clients.openWindow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clients) OpenWindowFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.ClientsOpenWindowFunc(
			this.Ref(),
		),
	)
}

// Claim calls the method "Clients.claim".
//
// The returned bool will be false if there is no such method.
func (this Clients) Claim() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallClientsClaim(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ClaimFunc returns the method "Clients.claim".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clients) ClaimFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClientsClaimFunc(
			this.Ref(),
		),
	)
}

type OneOf_String_Blob struct {
	ref js.Ref
}

func (x OneOf_String_Blob) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Blob) Free() {
	x.ref.Free()
}

func (x OneOf_String_Blob) FromRef(ref js.Ref) OneOf_String_Blob {
	return OneOf_String_Blob{
		ref: ref,
	}
}

func (x OneOf_String_Blob) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Blob) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

type ClipboardItemData = js.Promise[OneOf_String_Blob]

type PresentationStyle uint32

const (
	_ PresentationStyle = iota

	PresentationStyle_UNSPECIFIED
	PresentationStyle_INLINE
	PresentationStyle_ATTACHMENT
)

func (PresentationStyle) FromRef(str js.Ref) PresentationStyle {
	return PresentationStyle(bindings.ConstOfPresentationStyle(str))
}

func (x PresentationStyle) String() (string, bool) {
	switch x {
	case PresentationStyle_UNSPECIFIED:
		return "unspecified", true
	case PresentationStyle_INLINE:
		return "inline", true
	case PresentationStyle_ATTACHMENT:
		return "attachment", true
	default:
		return "", false
	}
}

type ClipboardItemOptions struct {
	// PresentationStyle is "ClipboardItemOptions.presentationStyle"
	//
	// Optional, defaults to "unspecified".
	PresentationStyle PresentationStyle

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClipboardItemOptions with all fields set.
func (p ClipboardItemOptions) FromRef(ref js.Ref) ClipboardItemOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ClipboardItemOptions ClipboardItemOptions [// ClipboardItemOptions] [0x140005bfd60] 0x14000574630 {0 0}} in the application heap.
func (p ClipboardItemOptions) New() js.Ref {
	return bindings.ClipboardItemOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ClipboardItemOptions) UpdateFrom(ref js.Ref) {
	bindings.ClipboardItemOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ClipboardItemOptions) Update(ref js.Ref) {
	bindings.ClipboardItemOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewClipboardItem(items js.Record[ClipboardItemData], options ClipboardItemOptions) ClipboardItem {
	return ClipboardItem{}.FromRef(
		bindings.NewClipboardItemByClipboardItem(
			items.Ref(),
			js.Pointer(&options)),
	)
}

func NewClipboardItemByClipboardItem1(items js.Record[ClipboardItemData]) ClipboardItem {
	return ClipboardItem{}.FromRef(
		bindings.NewClipboardItemByClipboardItem1(
			items.Ref()),
	)
}

type ClipboardItem struct {
	ref js.Ref
}

func (this ClipboardItem) Once() ClipboardItem {
	this.Ref().Once()
	return this
}

func (this ClipboardItem) Ref() js.Ref {
	return this.ref
}

func (this ClipboardItem) FromRef(ref js.Ref) ClipboardItem {
	this.ref = ref
	return this
}

func (this ClipboardItem) Free() {
	this.Ref().Free()
}

// PresentationStyle returns the value of property "ClipboardItem.presentationStyle".
//
// The returned bool will be false if there is no such property.
func (this ClipboardItem) PresentationStyle() (PresentationStyle, bool) {
	var _ok bool
	_ret := bindings.GetClipboardItemPresentationStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationStyle(_ret), _ok
}

// Types returns the value of property "ClipboardItem.types".
//
// The returned bool will be false if there is no such property.
func (this ClipboardItem) Types() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetClipboardItemTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// GetType calls the method "ClipboardItem.getType".
//
// The returned bool will be false if there is no such method.
func (this ClipboardItem) GetType(typ js.String) (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallClipboardItemGetType(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// GetTypeFunc returns the method "ClipboardItem.getType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ClipboardItem) GetTypeFunc() (fn js.Func[func(typ js.String) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ClipboardItemGetTypeFunc(
			this.Ref(),
		),
	)
}

type ClipboardItems = js.Array[ClipboardItem]

type Clipboard struct {
	EventTarget
}

func (this Clipboard) Once() Clipboard {
	this.Ref().Once()
	return this
}

func (this Clipboard) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Clipboard) FromRef(ref js.Ref) Clipboard {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Clipboard) Free() {
	this.Ref().Free()
}

// Read calls the method "Clipboard.read".
//
// The returned bool will be false if there is no such method.
func (this Clipboard) Read() (js.Promise[ClipboardItems], bool) {
	var _ok bool
	_ret := bindings.CallClipboardRead(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ClipboardItems]{}.FromRef(_ret), _ok
}

// ReadFunc returns the method "Clipboard.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clipboard) ReadFunc() (fn js.Func[func() js.Promise[ClipboardItems]]) {
	return fn.FromRef(
		bindings.ClipboardReadFunc(
			this.Ref(),
		),
	)
}

// ReadText calls the method "Clipboard.readText".
//
// The returned bool will be false if there is no such method.
func (this Clipboard) ReadText() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallClipboardReadText(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// ReadTextFunc returns the method "Clipboard.readText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clipboard) ReadTextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.ClipboardReadTextFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "Clipboard.write".
//
// The returned bool will be false if there is no such method.
func (this Clipboard) Write(data ClipboardItems) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallClipboardWrite(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteFunc returns the method "Clipboard.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clipboard) WriteFunc() (fn js.Func[func(data ClipboardItems) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClipboardWriteFunc(
			this.Ref(),
		),
	)
}

// WriteText calls the method "Clipboard.writeText".
//
// The returned bool will be false if there is no such method.
func (this Clipboard) WriteText(data js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallClipboardWriteText(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteTextFunc returns the method "Clipboard.writeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Clipboard) WriteTextFunc() (fn js.Func[func(data js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClipboardWriteTextFunc(
			this.Ref(),
		),
	)
}

type FunctionStringCallbackFunc func(this js.Ref, data js.String) js.Ref

func (fn FunctionStringCallbackFunc) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FunctionStringCallbackFunc) DispatchCallback(
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

type FunctionStringCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.String) js.Ref
	Arg T
}

func (cb *FunctionStringCallback[T]) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FunctionStringCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PermissionState uint32

const (
	_ PermissionState = iota

	PermissionState_GRANTED
	PermissionState_DENIED
	PermissionState_PROMPT
)

func (PermissionState) FromRef(str js.Ref) PermissionState {
	return PermissionState(bindings.ConstOfPermissionState(str))
}

func (x PermissionState) String() (string, bool) {
	switch x {
	case PermissionState_GRANTED:
		return "granted", true
	case PermissionState_DENIED:
		return "denied", true
	case PermissionState_PROMPT:
		return "prompt", true
	default:
		return "", false
	}
}

type FileSystemPermissionMode uint32

const (
	_ FileSystemPermissionMode = iota

	FileSystemPermissionMode_READ
	FileSystemPermissionMode_READWRITE
)

func (FileSystemPermissionMode) FromRef(str js.Ref) FileSystemPermissionMode {
	return FileSystemPermissionMode(bindings.ConstOfFileSystemPermissionMode(str))
}

func (x FileSystemPermissionMode) String() (string, bool) {
	switch x {
	case FileSystemPermissionMode_READ:
		return "read", true
	case FileSystemPermissionMode_READWRITE:
		return "readwrite", true
	default:
		return "", false
	}
}

type FileSystemHandlePermissionDescriptor struct {
	// Mode is "FileSystemHandlePermissionDescriptor.mode"
	//
	// Optional, defaults to "read".
	Mode FileSystemPermissionMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemHandlePermissionDescriptor with all fields set.
func (p FileSystemHandlePermissionDescriptor) FromRef(ref js.Ref) FileSystemHandlePermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FileSystemHandlePermissionDescriptor FileSystemHandlePermissionDescriptor [// FileSystemHandlePermissionDescriptor] [0x140005bff40] 0x140005746a8 {0 0}} in the application heap.
func (p FileSystemHandlePermissionDescriptor) New() js.Ref {
	return bindings.FileSystemHandlePermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FileSystemHandlePermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.FileSystemHandlePermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemHandlePermissionDescriptor) Update(ref js.Ref) {
	bindings.FileSystemHandlePermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FileSystemHandleKind uint32

const (
	_ FileSystemHandleKind = iota

	FileSystemHandleKind_FILE
	FileSystemHandleKind_DIRECTORY
)

func (FileSystemHandleKind) FromRef(str js.Ref) FileSystemHandleKind {
	return FileSystemHandleKind(bindings.ConstOfFileSystemHandleKind(str))
}

func (x FileSystemHandleKind) String() (string, bool) {
	switch x {
	case FileSystemHandleKind_FILE:
		return "file", true
	case FileSystemHandleKind_DIRECTORY:
		return "directory", true
	default:
		return "", false
	}
}

type FileSystemHandle struct {
	ref js.Ref
}

func (this FileSystemHandle) Once() FileSystemHandle {
	this.Ref().Once()
	return this
}

func (this FileSystemHandle) Ref() js.Ref {
	return this.ref
}

func (this FileSystemHandle) FromRef(ref js.Ref) FileSystemHandle {
	this.ref = ref
	return this
}

func (this FileSystemHandle) Free() {
	this.Ref().Free()
}

// Kind returns the value of property "FileSystemHandle.kind".
//
// The returned bool will be false if there is no such property.
func (this FileSystemHandle) Kind() (FileSystemHandleKind, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemHandleKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return FileSystemHandleKind(_ret), _ok
}

// Name returns the value of property "FileSystemHandle.name".
//
// The returned bool will be false if there is no such property.
func (this FileSystemHandle) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemHandleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsSameEntry calls the method "FileSystemHandle.isSameEntry".
//
// The returned bool will be false if there is no such method.
func (this FileSystemHandle) IsSameEntry(other FileSystemHandle) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemHandleIsSameEntry(
		this.Ref(), js.Pointer(&_ok),
		other.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsSameEntryFunc returns the method "FileSystemHandle.isSameEntry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemHandle) IsSameEntryFunc() (fn js.Func[func(other FileSystemHandle) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.FileSystemHandleIsSameEntryFunc(
			this.Ref(),
		),
	)
}

// QueryPermission calls the method "FileSystemHandle.queryPermission".
//
// The returned bool will be false if there is no such method.
func (this FileSystemHandle) QueryPermission(descriptor FileSystemHandlePermissionDescriptor) (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemHandleQueryPermission(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// QueryPermissionFunc returns the method "FileSystemHandle.queryPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemHandle) QueryPermissionFunc() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleQueryPermissionFunc(
			this.Ref(),
		),
	)
}

// QueryPermission1 calls the method "FileSystemHandle.queryPermission".
//
// The returned bool will be false if there is no such method.
func (this FileSystemHandle) QueryPermission1() (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemHandleQueryPermission1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// QueryPermission1Func returns the method "FileSystemHandle.queryPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemHandle) QueryPermission1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleQueryPermission1Func(
			this.Ref(),
		),
	)
}

// RequestPermission calls the method "FileSystemHandle.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this FileSystemHandle) RequestPermission(descriptor FileSystemHandlePermissionDescriptor) (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemHandleRequestPermission(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// RequestPermissionFunc returns the method "FileSystemHandle.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemHandle) RequestPermissionFunc() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// RequestPermission1 calls the method "FileSystemHandle.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this FileSystemHandle) RequestPermission1() (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemHandleRequestPermission1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// RequestPermission1Func returns the method "FileSystemHandle.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemHandle) RequestPermission1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleRequestPermission1Func(
			this.Ref(),
		),
	)
}

type FileSystemEntryCallbackFunc func(this js.Ref, entry FileSystemEntry) js.Ref

func (fn FileSystemEntryCallbackFunc) Register() js.Func[func(entry FileSystemEntry)] {
	return js.RegisterCallback[func(entry FileSystemEntry)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FileSystemEntryCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		FileSystemEntry{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileSystemEntryCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entry FileSystemEntry) js.Ref
	Arg T
}

func (cb *FileSystemEntryCallback[T]) Register() js.Func[func(entry FileSystemEntry)] {
	return js.RegisterCallback[func(entry FileSystemEntry)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FileSystemEntryCallback[T]) DispatchCallback(
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

		FileSystemEntry{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ErrorCallbackFunc func(this js.Ref, err DOMException) js.Ref

func (fn ErrorCallbackFunc) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err DOMException) js.Ref
	Arg T
}

func (cb *ErrorCallback[T]) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ErrorCallback[T]) DispatchCallback(
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

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileSystemEntriesCallbackFunc func(this js.Ref, entries js.Array[FileSystemEntry]) js.Ref

func (fn FileSystemEntriesCallbackFunc) Register() js.Func[func(entries js.Array[FileSystemEntry])] {
	return js.RegisterCallback[func(entries js.Array[FileSystemEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FileSystemEntriesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[FileSystemEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileSystemEntriesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[FileSystemEntry]) js.Ref
	Arg T
}

func (cb *FileSystemEntriesCallback[T]) Register() js.Func[func(entries js.Array[FileSystemEntry])] {
	return js.RegisterCallback[func(entries js.Array[FileSystemEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FileSystemEntriesCallback[T]) DispatchCallback(
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

		js.Array[FileSystemEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileSystemDirectoryReader struct {
	ref js.Ref
}

func (this FileSystemDirectoryReader) Once() FileSystemDirectoryReader {
	this.Ref().Once()
	return this
}

func (this FileSystemDirectoryReader) Ref() js.Ref {
	return this.ref
}

func (this FileSystemDirectoryReader) FromRef(ref js.Ref) FileSystemDirectoryReader {
	this.ref = ref
	return this
}

func (this FileSystemDirectoryReader) Free() {
	this.Ref().Free()
}

// ReadEntries calls the method "FileSystemDirectoryReader.readEntries".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryReader) ReadEntries(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryReaderReadEntries(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadEntriesFunc returns the method "FileSystemDirectoryReader.readEntries".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryReader) ReadEntriesFunc() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryReaderReadEntriesFunc(
			this.Ref(),
		),
	)
}

// ReadEntries1 calls the method "FileSystemDirectoryReader.readEntries".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryReader) ReadEntries1(successCallback js.Func[func(entries js.Array[FileSystemEntry])]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryReaderReadEntries1(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadEntries1Func returns the method "FileSystemDirectoryReader.readEntries".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryReader) ReadEntries1Func() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryReaderReadEntries1Func(
			this.Ref(),
		),
	)
}

type FileSystemFlags struct {
	// Create is "FileSystemFlags.create"
	//
	// Optional, defaults to false.
	Create bool
	// Exclusive is "FileSystemFlags.exclusive"
	//
	// Optional, defaults to false.
	Exclusive bool

	FFI_USE_Create    bool // for Create.
	FFI_USE_Exclusive bool // for Exclusive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemFlags with all fields set.
func (p FileSystemFlags) FromRef(ref js.Ref) FileSystemFlags {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FileSystemFlags FileSystemFlags [// FileSystemFlags] [0x140005e0140 0x140005e0280 0x140005e01e0 0x140005e0320] 0x14000574738 {0 0}} in the application heap.
func (p FileSystemFlags) New() js.Ref {
	return bindings.FileSystemFlagsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FileSystemFlags) UpdateFrom(ref js.Ref) {
	bindings.FileSystemFlagsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemFlags) Update(ref js.Ref) {
	bindings.FileSystemFlagsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FileSystemDirectoryEntry struct {
	FileSystemEntry
}

func (this FileSystemDirectoryEntry) Once() FileSystemDirectoryEntry {
	this.Ref().Once()
	return this
}

func (this FileSystemDirectoryEntry) Ref() js.Ref {
	return this.FileSystemEntry.Ref()
}

func (this FileSystemDirectoryEntry) FromRef(ref js.Ref) FileSystemDirectoryEntry {
	this.FileSystemEntry = this.FileSystemEntry.FromRef(ref)
	return this
}

func (this FileSystemDirectoryEntry) Free() {
	this.Ref().Free()
}

// CreateReader calls the method "FileSystemDirectoryEntry.createReader".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) CreateReader() (FileSystemDirectoryReader, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryCreateReader(
		this.Ref(), js.Pointer(&_ok),
	)

	return FileSystemDirectoryReader{}.FromRef(_ret), _ok
}

// CreateReaderFunc returns the method "FileSystemDirectoryEntry.createReader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) CreateReaderFunc() (fn js.Func[func() FileSystemDirectoryReader]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryCreateReaderFunc(
			this.Ref(),
		),
	)
}

// GetFile calls the method "FileSystemDirectoryEntry.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetFile(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetFile(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFileFunc returns the method "FileSystemDirectoryEntry.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetFileFunc() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFileFunc(
			this.Ref(),
		),
	)
}

// GetFile1 calls the method "FileSystemDirectoryEntry.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetFile1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetFile1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFile1Func returns the method "FileSystemDirectoryEntry.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetFile1Func() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile1Func(
			this.Ref(),
		),
	)
}

// GetFile2 calls the method "FileSystemDirectoryEntry.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetFile2(path js.String, options FileSystemFlags) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetFile2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFile2Func returns the method "FileSystemDirectoryEntry.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetFile2Func() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile2Func(
			this.Ref(),
		),
	)
}

// GetFile3 calls the method "FileSystemDirectoryEntry.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetFile3(path js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetFile3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFile3Func returns the method "FileSystemDirectoryEntry.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetFile3Func() (fn js.Func[func(path js.String)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile3Func(
			this.Ref(),
		),
	)
}

// GetFile4 calls the method "FileSystemDirectoryEntry.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetFile4() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetFile4(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFile4Func returns the method "FileSystemDirectoryEntry.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetFile4Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile4Func(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "FileSystemDirectoryEntry.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetDirectory(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetDirectoryFunc returns the method "FileSystemDirectoryEntry.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectoryFunc() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectoryFunc(
			this.Ref(),
		),
	)
}

// GetDirectory1 calls the method "FileSystemDirectoryEntry.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetDirectory1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetDirectory1Func returns the method "FileSystemDirectoryEntry.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory1Func() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory1Func(
			this.Ref(),
		),
	)
}

// GetDirectory2 calls the method "FileSystemDirectoryEntry.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory2(path js.String, options FileSystemFlags) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetDirectory2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetDirectory2Func returns the method "FileSystemDirectoryEntry.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory2Func() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory2Func(
			this.Ref(),
		),
	)
}

// GetDirectory3 calls the method "FileSystemDirectoryEntry.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory3(path js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetDirectory3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetDirectory3Func returns the method "FileSystemDirectoryEntry.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory3Func() (fn js.Func[func(path js.String)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory3Func(
			this.Ref(),
		),
	)
}

// GetDirectory4 calls the method "FileSystemDirectoryEntry.getDirectory".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory4() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryEntryGetDirectory4(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetDirectory4Func returns the method "FileSystemDirectoryEntry.getDirectory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryEntry) GetDirectory4Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory4Func(
			this.Ref(),
		),
	)
}

type FileSystem struct {
	ref js.Ref
}

func (this FileSystem) Once() FileSystem {
	this.Ref().Once()
	return this
}

func (this FileSystem) Ref() js.Ref {
	return this.ref
}

func (this FileSystem) FromRef(ref js.Ref) FileSystem {
	this.ref = ref
	return this
}

func (this FileSystem) Free() {
	this.Ref().Free()
}

// Name returns the value of property "FileSystem.name".
//
// The returned bool will be false if there is no such property.
func (this FileSystem) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Root returns the value of property "FileSystem.root".
//
// The returned bool will be false if there is no such property.
func (this FileSystem) Root() (FileSystemDirectoryEntry, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return FileSystemDirectoryEntry{}.FromRef(_ret), _ok
}

type FileSystemEntry struct {
	ref js.Ref
}

func (this FileSystemEntry) Once() FileSystemEntry {
	this.Ref().Once()
	return this
}

func (this FileSystemEntry) Ref() js.Ref {
	return this.ref
}

func (this FileSystemEntry) FromRef(ref js.Ref) FileSystemEntry {
	this.ref = ref
	return this
}

func (this FileSystemEntry) Free() {
	this.Ref().Free()
}

// IsFile returns the value of property "FileSystemEntry.isFile".
//
// The returned bool will be false if there is no such property.
func (this FileSystemEntry) IsFile() (bool, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemEntryIsFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsDirectory returns the value of property "FileSystemEntry.isDirectory".
//
// The returned bool will be false if there is no such property.
func (this FileSystemEntry) IsDirectory() (bool, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemEntryIsDirectory(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Name returns the value of property "FileSystemEntry.name".
//
// The returned bool will be false if there is no such property.
func (this FileSystemEntry) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemEntryName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FullPath returns the value of property "FileSystemEntry.fullPath".
//
// The returned bool will be false if there is no such property.
func (this FileSystemEntry) FullPath() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemEntryFullPath(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Filesystem returns the value of property "FileSystemEntry.filesystem".
//
// The returned bool will be false if there is no such property.
func (this FileSystemEntry) Filesystem() (FileSystem, bool) {
	var _ok bool
	_ret := bindings.GetFileSystemEntryFilesystem(
		this.Ref(), js.Pointer(&_ok),
	)
	return FileSystem{}.FromRef(_ret), _ok
}

// GetParent calls the method "FileSystemEntry.getParent".
//
// The returned bool will be false if there is no such method.
func (this FileSystemEntry) GetParent(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemEntryGetParent(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetParentFunc returns the method "FileSystemEntry.getParent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemEntry) GetParentFunc() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParentFunc(
			this.Ref(),
		),
	)
}

// GetParent1 calls the method "FileSystemEntry.getParent".
//
// The returned bool will be false if there is no such method.
func (this FileSystemEntry) GetParent1(successCallback js.Func[func(entry FileSystemEntry)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemEntryGetParent1(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetParent1Func returns the method "FileSystemEntry.getParent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemEntry) GetParent1Func() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParent1Func(
			this.Ref(),
		),
	)
}

// GetParent2 calls the method "FileSystemEntry.getParent".
//
// The returned bool will be false if there is no such method.
func (this FileSystemEntry) GetParent2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemEntryGetParent2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetParent2Func returns the method "FileSystemEntry.getParent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemEntry) GetParent2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParent2Func(
			this.Ref(),
		),
	)
}

type DataTransferItem struct {
	ref js.Ref
}

func (this DataTransferItem) Once() DataTransferItem {
	this.Ref().Once()
	return this
}

func (this DataTransferItem) Ref() js.Ref {
	return this.ref
}

func (this DataTransferItem) FromRef(ref js.Ref) DataTransferItem {
	this.ref = ref
	return this
}

func (this DataTransferItem) Free() {
	this.Ref().Free()
}

// Kind returns the value of property "DataTransferItem.kind".
//
// The returned bool will be false if there is no such property.
func (this DataTransferItem) Kind() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferItemKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type returns the value of property "DataTransferItem.type".
//
// The returned bool will be false if there is no such property.
func (this DataTransferItem) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferItemType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GetAsString calls the method "DataTransferItem.getAsString".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItem) GetAsString(callback js.Func[func(data js.String)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemGetAsString(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetAsStringFunc returns the method "DataTransferItem.getAsString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItem) GetAsStringFunc() (fn js.Func[func(callback js.Func[func(data js.String)])]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsStringFunc(
			this.Ref(),
		),
	)
}

// GetAsFile calls the method "DataTransferItem.getAsFile".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItem) GetAsFile() (File, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemGetAsFile(
		this.Ref(), js.Pointer(&_ok),
	)

	return File{}.FromRef(_ret), _ok
}

// GetAsFileFunc returns the method "DataTransferItem.getAsFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItem) GetAsFileFunc() (fn js.Func[func() File]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsFileFunc(
			this.Ref(),
		),
	)
}

// GetAsFileSystemHandle calls the method "DataTransferItem.getAsFileSystemHandle".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItem) GetAsFileSystemHandle() (js.Promise[FileSystemHandle], bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemGetAsFileSystemHandle(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemHandle]{}.FromRef(_ret), _ok
}

// GetAsFileSystemHandleFunc returns the method "DataTransferItem.getAsFileSystemHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItem) GetAsFileSystemHandleFunc() (fn js.Func[func() js.Promise[FileSystemHandle]]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsFileSystemHandleFunc(
			this.Ref(),
		),
	)
}

// WebkitGetAsEntry calls the method "DataTransferItem.webkitGetAsEntry".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItem) WebkitGetAsEntry() (FileSystemEntry, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemWebkitGetAsEntry(
		this.Ref(), js.Pointer(&_ok),
	)

	return FileSystemEntry{}.FromRef(_ret), _ok
}

// WebkitGetAsEntryFunc returns the method "DataTransferItem.webkitGetAsEntry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItem) WebkitGetAsEntryFunc() (fn js.Func[func() FileSystemEntry]) {
	return fn.FromRef(
		bindings.DataTransferItemWebkitGetAsEntryFunc(
			this.Ref(),
		),
	)
}

type DataTransferItemList struct {
	ref js.Ref
}

func (this DataTransferItemList) Once() DataTransferItemList {
	this.Ref().Once()
	return this
}

func (this DataTransferItemList) Ref() js.Ref {
	return this.ref
}

func (this DataTransferItemList) FromRef(ref js.Ref) DataTransferItemList {
	this.ref = ref
	return this
}

func (this DataTransferItemList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "DataTransferItemList.length".
//
// The returned bool will be false if there is no such property.
func (this DataTransferItemList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferItemListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "DataTransferItemList.".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItemList) Get(index uint32) (DataTransferItem, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return DataTransferItem{}.FromRef(_ret), _ok
}

// GetFunc returns the method "DataTransferItemList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItemList) GetFunc() (fn js.Func[func(index uint32) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListGetFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "DataTransferItemList.add".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItemList) Add(data js.String, typ js.String) (DataTransferItem, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemListAdd(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		typ.Ref(),
	)

	return DataTransferItem{}.FromRef(_ret), _ok
}

// AddFunc returns the method "DataTransferItemList.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItemList) AddFunc() (fn js.Func[func(data js.String, typ js.String) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListAddFunc(
			this.Ref(),
		),
	)
}

// Add1 calls the method "DataTransferItemList.add".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItemList) Add1(data File) (DataTransferItem, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemListAdd1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return DataTransferItem{}.FromRef(_ret), _ok
}

// Add1Func returns the method "DataTransferItemList.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItemList) Add1Func() (fn js.Func[func(data File) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListAdd1Func(
			this.Ref(),
		),
	)
}

// Remove calls the method "DataTransferItemList.remove".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItemList) Remove(index uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemListRemove(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "DataTransferItemList.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItemList) RemoveFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.DataTransferItemListRemoveFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "DataTransferItemList.clear".
//
// The returned bool will be false if there is no such method.
func (this DataTransferItemList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferItemListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "DataTransferItemList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransferItemList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DataTransferItemListClearFunc(
			this.Ref(),
		),
	)
}

type FileList struct {
	ref js.Ref
}

func (this FileList) Once() FileList {
	this.Ref().Once()
	return this
}

func (this FileList) Ref() js.Ref {
	return this.ref
}

func (this FileList) FromRef(ref js.Ref) FileList {
	this.ref = ref
	return this
}

func (this FileList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "FileList.length".
//
// The returned bool will be false if there is no such property.
func (this FileList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetFileListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "FileList.item".
//
// The returned bool will be false if there is no such method.
func (this FileList) Item(index uint32) (File, bool) {
	var _ok bool
	_ret := bindings.CallFileListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return File{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "FileList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileList) ItemFunc() (fn js.Func[func(index uint32) File]) {
	return fn.FromRef(
		bindings.FileListItemFunc(
			this.Ref(),
		),
	)
}

type DataTransfer struct {
	ref js.Ref
}

func (this DataTransfer) Once() DataTransfer {
	this.Ref().Once()
	return this
}

func (this DataTransfer) Ref() js.Ref {
	return this.ref
}

func (this DataTransfer) FromRef(ref js.Ref) DataTransfer {
	this.ref = ref
	return this
}

func (this DataTransfer) Free() {
	this.Ref().Free()
}

// DropEffect returns the value of property "DataTransfer.dropEffect".
//
// The returned bool will be false if there is no such property.
func (this DataTransfer) DropEffect() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferDropEffect(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DropEffect sets the value of property "DataTransfer.dropEffect" to val.
//
// It returns false if the property cannot be set.
func (this DataTransfer) SetDropEffect(val js.String) bool {
	return js.True == bindings.SetDataTransferDropEffect(
		this.Ref(),
		val.Ref(),
	)
}

// EffectAllowed returns the value of property "DataTransfer.effectAllowed".
//
// The returned bool will be false if there is no such property.
func (this DataTransfer) EffectAllowed() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferEffectAllowed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EffectAllowed sets the value of property "DataTransfer.effectAllowed" to val.
//
// It returns false if the property cannot be set.
func (this DataTransfer) SetEffectAllowed(val js.String) bool {
	return js.True == bindings.SetDataTransferEffectAllowed(
		this.Ref(),
		val.Ref(),
	)
}

// Items returns the value of property "DataTransfer.items".
//
// The returned bool will be false if there is no such property.
func (this DataTransfer) Items() (DataTransferItemList, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return DataTransferItemList{}.FromRef(_ret), _ok
}

// Types returns the value of property "DataTransfer.types".
//
// The returned bool will be false if there is no such property.
func (this DataTransfer) Types() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetDataTransferTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Files returns the value of property "DataTransfer.files".
//
// The returned bool will be false if there is no such property.
func (this DataTransfer) Files() (FileList, bool) {
	var _ok bool
	_ret := bindings.GetDataTransferFiles(
		this.Ref(), js.Pointer(&_ok),
	)
	return FileList{}.FromRef(_ret), _ok
}

// SetDragImage calls the method "DataTransfer.setDragImage".
//
// The returned bool will be false if there is no such method.
func (this DataTransfer) SetDragImage(image Element, x int32, y int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferSetDragImage(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetDragImageFunc returns the method "DataTransfer.setDragImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransfer) SetDragImageFunc() (fn js.Func[func(image Element, x int32, y int32)]) {
	return fn.FromRef(
		bindings.DataTransferSetDragImageFunc(
			this.Ref(),
		),
	)
}

// GetData calls the method "DataTransfer.getData".
//
// The returned bool will be false if there is no such method.
func (this DataTransfer) GetData(format js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferGetData(
		this.Ref(), js.Pointer(&_ok),
		format.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetDataFunc returns the method "DataTransfer.getData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransfer) GetDataFunc() (fn js.Func[func(format js.String) js.String]) {
	return fn.FromRef(
		bindings.DataTransferGetDataFunc(
			this.Ref(),
		),
	)
}

// SetData calls the method "DataTransfer.setData".
//
// The returned bool will be false if there is no such method.
func (this DataTransfer) SetData(format js.String, data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferSetData(
		this.Ref(), js.Pointer(&_ok),
		format.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetDataFunc returns the method "DataTransfer.setData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransfer) SetDataFunc() (fn js.Func[func(format js.String, data js.String)]) {
	return fn.FromRef(
		bindings.DataTransferSetDataFunc(
			this.Ref(),
		),
	)
}

// ClearData calls the method "DataTransfer.clearData".
//
// The returned bool will be false if there is no such method.
func (this DataTransfer) ClearData(format js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferClearData(
		this.Ref(), js.Pointer(&_ok),
		format.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearDataFunc returns the method "DataTransfer.clearData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransfer) ClearDataFunc() (fn js.Func[func(format js.String)]) {
	return fn.FromRef(
		bindings.DataTransferClearDataFunc(
			this.Ref(),
		),
	)
}

// ClearData1 calls the method "DataTransfer.clearData".
//
// The returned bool will be false if there is no such method.
func (this DataTransfer) ClearData1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDataTransferClearData1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearData1Func returns the method "DataTransfer.clearData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DataTransfer) ClearData1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DataTransferClearData1Func(
			this.Ref(),
		),
	)
}

type ClipboardEventInit struct {
	// ClipboardData is "ClipboardEventInit.clipboardData"
	//
	// Optional, defaults to null.
	ClipboardData DataTransfer
	// Bubbles is "ClipboardEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "ClipboardEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "ClipboardEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClipboardEventInit with all fields set.
func (p ClipboardEventInit) FromRef(ref js.Ref) ClipboardEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ClipboardEventInit ClipboardEventInit [// ClipboardEventInit] [0x140005e0500 0x140005e05a0 0x140005e06e0 0x140005e0820 0x140005e0640 0x140005e0780 0x140005e08c0] 0x14000574678 {0 0}} in the application heap.
func (p ClipboardEventInit) New() js.Ref {
	return bindings.ClipboardEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ClipboardEventInit) UpdateFrom(ref js.Ref) {
	bindings.ClipboardEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ClipboardEventInit) Update(ref js.Ref) {
	bindings.ClipboardEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewClipboardEvent(typ js.String, eventInitDict ClipboardEventInit) ClipboardEvent {
	return ClipboardEvent{}.FromRef(
		bindings.NewClipboardEventByClipboardEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewClipboardEventByClipboardEvent1(typ js.String) ClipboardEvent {
	return ClipboardEvent{}.FromRef(
		bindings.NewClipboardEventByClipboardEvent1(
			typ.Ref()),
	)
}

type ClipboardEvent struct {
	Event
}

func (this ClipboardEvent) Once() ClipboardEvent {
	this.Ref().Once()
	return this
}

func (this ClipboardEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ClipboardEvent) FromRef(ref js.Ref) ClipboardEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ClipboardEvent) Free() {
	this.Ref().Free()
}

// ClipboardData returns the value of property "ClipboardEvent.clipboardData".
//
// The returned bool will be false if there is no such property.
func (this ClipboardEvent) ClipboardData() (DataTransfer, bool) {
	var _ok bool
	_ret := bindings.GetClipboardEventClipboardData(
		this.Ref(), js.Pointer(&_ok),
	)
	return DataTransfer{}.FromRef(_ret), _ok
}

type ClipboardPermissionDescriptor struct {
	// AllowWithoutGesture is "ClipboardPermissionDescriptor.allowWithoutGesture"
	//
	// Optional, defaults to false.
	AllowWithoutGesture bool
	// Name is "ClipboardPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE_AllowWithoutGesture bool // for AllowWithoutGesture.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClipboardPermissionDescriptor with all fields set.
func (p ClipboardPermissionDescriptor) FromRef(ref js.Ref) ClipboardPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ClipboardPermissionDescriptor ClipboardPermissionDescriptor [// ClipboardPermissionDescriptor] [0x140005e0960 0x140005e0aa0 0x140005e0a00] 0x140005747e0 {0 0}} in the application heap.
func (p ClipboardPermissionDescriptor) New() js.Ref {
	return bindings.ClipboardPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ClipboardPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ClipboardPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ClipboardPermissionDescriptor) Update(ref js.Ref) {
	bindings.ClipboardPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CloseEventInit struct {
	// WasClean is "CloseEventInit.wasClean"
	//
	// Optional, defaults to false.
	WasClean bool
	// Code is "CloseEventInit.code"
	//
	// Optional, defaults to 0.
	Code uint16
	// Reason is "CloseEventInit.reason"
	//
	// Optional, defaults to "".
	Reason js.String
	// Bubbles is "CloseEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CloseEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CloseEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_WasClean   bool // for WasClean.
	FFI_USE_Code       bool // for Code.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CloseEventInit with all fields set.
func (p CloseEventInit) FromRef(ref js.Ref) CloseEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CloseEventInit CloseEventInit [// CloseEventInit] [0x140005e0b40 0x140005e0c80 0x140005e0dc0 0x140005e0e60 0x140005e0fa0 0x140005e10e0 0x140005e0be0 0x140005e0d20 0x140005e0f00 0x140005e1040 0x140005e1180] 0x14000574840 {0 0}} in the application heap.
func (p CloseEventInit) New() js.Ref {
	return bindings.CloseEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CloseEventInit) UpdateFrom(ref js.Ref) {
	bindings.CloseEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CloseEventInit) Update(ref js.Ref) {
	bindings.CloseEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCloseEvent(typ js.String, eventInitDict CloseEventInit) CloseEvent {
	return CloseEvent{}.FromRef(
		bindings.NewCloseEventByCloseEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewCloseEventByCloseEvent1(typ js.String) CloseEvent {
	return CloseEvent{}.FromRef(
		bindings.NewCloseEventByCloseEvent1(
			typ.Ref()),
	)
}

type CloseEvent struct {
	Event
}

func (this CloseEvent) Once() CloseEvent {
	this.Ref().Once()
	return this
}

func (this CloseEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CloseEvent) FromRef(ref js.Ref) CloseEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CloseEvent) Free() {
	this.Ref().Free()
}

// WasClean returns the value of property "CloseEvent.wasClean".
//
// The returned bool will be false if there is no such property.
func (this CloseEvent) WasClean() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCloseEventWasClean(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Code returns the value of property "CloseEvent.code".
//
// The returned bool will be false if there is no such property.
func (this CloseEvent) Code() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetCloseEventCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Reason returns the value of property "CloseEvent.reason".
//
// The returned bool will be false if there is no such property.
func (this CloseEvent) Reason() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCloseEventReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CollectedClientAdditionalPaymentData struct {
	// RpId is "CollectedClientAdditionalPaymentData.rpId"
	//
	// Required
	RpId js.String
	// TopOrigin is "CollectedClientAdditionalPaymentData.topOrigin"
	//
	// Required
	TopOrigin js.String
	// PayeeName is "CollectedClientAdditionalPaymentData.payeeName"
	//
	// Optional
	PayeeName js.String
	// PayeeOrigin is "CollectedClientAdditionalPaymentData.payeeOrigin"
	//
	// Optional
	PayeeOrigin js.String
	// Total is "CollectedClientAdditionalPaymentData.total"
	//
	// Required
	Total PaymentCurrencyAmount
	// Instrument is "CollectedClientAdditionalPaymentData.instrument"
	//
	// Required
	Instrument PaymentCredentialInstrument

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientAdditionalPaymentData with all fields set.
func (p CollectedClientAdditionalPaymentData) FromRef(ref js.Ref) CollectedClientAdditionalPaymentData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CollectedClientAdditionalPaymentData CollectedClientAdditionalPaymentData [// CollectedClientAdditionalPaymentData] [0x140005e12c0 0x140005e1360 0x140005e1400 0x140005e14a0 0x140005e1540 0x140005e15e0] 0x14000574888 {0 0}} in the application heap.
func (p CollectedClientAdditionalPaymentData) New() js.Ref {
	return bindings.CollectedClientAdditionalPaymentDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CollectedClientAdditionalPaymentData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientAdditionalPaymentDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CollectedClientAdditionalPaymentData) Update(ref js.Ref) {
	bindings.CollectedClientAdditionalPaymentDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CollectedClientData struct {
	// Type is "CollectedClientData.type"
	//
	// Required
	Type js.String
	// Challenge is "CollectedClientData.challenge"
	//
	// Required
	Challenge js.String
	// Origin is "CollectedClientData.origin"
	//
	// Required
	Origin js.String
	// TopOrigin is "CollectedClientData.topOrigin"
	//
	// Optional
	TopOrigin js.String
	// CrossOrigin is "CollectedClientData.crossOrigin"
	//
	// Optional
	CrossOrigin bool

	FFI_USE_CrossOrigin bool // for CrossOrigin.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientData with all fields set.
func (p CollectedClientData) FromRef(ref js.Ref) CollectedClientData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CollectedClientData CollectedClientData [// CollectedClientData] [0x140005e1680 0x140005e1720 0x140005e17c0 0x140005e1860 0x140005e1900 0x140005e19a0] 0x140005748d0 {0 0}} in the application heap.
func (p CollectedClientData) New() js.Ref {
	return bindings.CollectedClientDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CollectedClientData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CollectedClientData) Update(ref js.Ref) {
	bindings.CollectedClientDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CollectedClientPaymentData struct {
	// Payment is "CollectedClientPaymentData.payment"
	//
	// Required
	Payment CollectedClientAdditionalPaymentData
	// Type is "CollectedClientPaymentData.type"
	//
	// Required
	Type js.String
	// Challenge is "CollectedClientPaymentData.challenge"
	//
	// Required
	Challenge js.String
	// Origin is "CollectedClientPaymentData.origin"
	//
	// Required
	Origin js.String
	// TopOrigin is "CollectedClientPaymentData.topOrigin"
	//
	// Optional
	TopOrigin js.String
	// CrossOrigin is "CollectedClientPaymentData.crossOrigin"
	//
	// Optional
	CrossOrigin bool

	FFI_USE_CrossOrigin bool // for CrossOrigin.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientPaymentData with all fields set.
func (p CollectedClientPaymentData) FromRef(ref js.Ref) CollectedClientPaymentData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CollectedClientPaymentData CollectedClientPaymentData [// CollectedClientPaymentData] [0x140005e1a40 0x140005e1ae0 0x140005e1b80 0x140005e1c20 0x140005e1cc0 0x140005e1d60 0x140005e1e00] 0x14000574918 {0 0}} in the application heap.
func (p CollectedClientPaymentData) New() js.Ref {
	return bindings.CollectedClientPaymentDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CollectedClientPaymentData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientPaymentDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CollectedClientPaymentData) Update(ref js.Ref) {
	bindings.CollectedClientPaymentDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ColorGamut uint32

const (
	_ ColorGamut = iota

	ColorGamut_SRGB
	ColorGamut_P3
	ColorGamut_REC2020
)

func (ColorGamut) FromRef(str js.Ref) ColorGamut {
	return ColorGamut(bindings.ConstOfColorGamut(str))
}

func (x ColorGamut) String() (string, bool) {
	switch x {
	case ColorGamut_SRGB:
		return "srgb", true
	case ColorGamut_P3:
		return "p3", true
	case ColorGamut_REC2020:
		return "rec2020", true
	default:
		return "", false
	}
}

type ColorSelectionOptions struct {
	// Signal is "ColorSelectionOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ColorSelectionOptions with all fields set.
func (p ColorSelectionOptions) FromRef(ref js.Ref) ColorSelectionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ColorSelectionOptions ColorSelectionOptions [// ColorSelectionOptions] [0x140005e1ea0] 0x14000574978 {0 0}} in the application heap.
func (p ColorSelectionOptions) New() js.Ref {
	return bindings.ColorSelectionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ColorSelectionOptions) UpdateFrom(ref js.Ref) {
	bindings.ColorSelectionOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ColorSelectionOptions) Update(ref js.Ref) {
	bindings.ColorSelectionOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ColorSelectionResult struct {
	// SRGBHex is "ColorSelectionResult.sRGBHex"
	//
	// Optional
	SRGBHex js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ColorSelectionResult with all fields set.
func (p ColorSelectionResult) FromRef(ref js.Ref) ColorSelectionResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ColorSelectionResult ColorSelectionResult [// ColorSelectionResult] [0x140005e1f40] 0x140005749a8 {0 0}} in the application heap.
func (p ColorSelectionResult) New() js.Ref {
	return bindings.ColorSelectionResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ColorSelectionResult) UpdateFrom(ref js.Ref) {
	bindings.ColorSelectionResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ColorSelectionResult) Update(ref js.Ref) {
	bindings.ColorSelectionResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ColorSpaceConversion uint32

const (
	_ ColorSpaceConversion = iota

	ColorSpaceConversion_NONE
	ColorSpaceConversion_DEFAULT
)

func (ColorSpaceConversion) FromRef(str js.Ref) ColorSpaceConversion {
	return ColorSpaceConversion(bindings.ConstOfColorSpaceConversion(str))
}

func (x ColorSpaceConversion) String() (string, bool) {
	switch x {
	case ColorSpaceConversion_NONE:
		return "none", true
	case ColorSpaceConversion_DEFAULT:
		return "default", true
	default:
		return "", false
	}
}

type CompositeOperation uint32

const (
	_ CompositeOperation = iota

	CompositeOperation_REPLACE
	CompositeOperation_ADD
	CompositeOperation_ACCUMULATE
)

func (CompositeOperation) FromRef(str js.Ref) CompositeOperation {
	return CompositeOperation(bindings.ConstOfCompositeOperation(str))
}

func (x CompositeOperation) String() (string, bool) {
	switch x {
	case CompositeOperation_REPLACE:
		return "replace", true
	case CompositeOperation_ADD:
		return "add", true
	case CompositeOperation_ACCUMULATE:
		return "accumulate", true
	default:
		return "", false
	}
}

type WindowPostMessageOptions struct {
	// TargetOrigin is "WindowPostMessageOptions.targetOrigin"
	//
	// Optional, defaults to "/".
	TargetOrigin js.String
	// Transfer is "WindowPostMessageOptions.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.Object]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WindowPostMessageOptions with all fields set.
func (p WindowPostMessageOptions) FromRef(ref js.Ref) WindowPostMessageOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 WindowPostMessageOptions WindowPostMessageOptions [// WindowPostMessageOptions] [0x140005f00a0 0x140005f0140] 0x14000574b28 {0 0}} in the application heap.
func (p WindowPostMessageOptions) New() js.Ref {
	return bindings.WindowPostMessageOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WindowPostMessageOptions) UpdateFrom(ref js.Ref) {
	bindings.WindowPostMessageOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WindowPostMessageOptions) Update(ref js.Ref) {
	bindings.WindowPostMessageOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaQueryList struct {
	EventTarget
}

func (this MediaQueryList) Once() MediaQueryList {
	this.Ref().Once()
	return this
}

func (this MediaQueryList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaQueryList) FromRef(ref js.Ref) MediaQueryList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaQueryList) Free() {
	this.Ref().Free()
}

// Media returns the value of property "MediaQueryList.media".
//
// The returned bool will be false if there is no such property.
func (this MediaQueryList) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaQueryListMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Matches returns the value of property "MediaQueryList.matches".
//
// The returned bool will be false if there is no such property.
func (this MediaQueryList) Matches() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaQueryListMatches(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AddListener calls the method "MediaQueryList.addListener".
//
// The returned bool will be false if there is no such method.
func (this MediaQueryList) AddListener(callback js.Func[func(event Event)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaQueryListAddListener(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddListenerFunc returns the method "MediaQueryList.addListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaQueryList) AddListenerFunc() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.MediaQueryListAddListenerFunc(
			this.Ref(),
		),
	)
}

// RemoveListener calls the method "MediaQueryList.removeListener".
//
// The returned bool will be false if there is no such method.
func (this MediaQueryList) RemoveListener(callback js.Func[func(event Event)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaQueryListRemoveListener(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveListenerFunc returns the method "MediaQueryList.removeListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaQueryList) RemoveListenerFunc() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.MediaQueryListRemoveListenerFunc(
			this.Ref(),
		),
	)
}

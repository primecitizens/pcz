// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
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
// It returns ok=false if there is no such property.
func (this WindowClient) VisibilityState() (ret DocumentVisibilityState, ok bool) {
	ok = js.True == bindings.GetWindowClientVisibilityState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Focused returns the value of property "WindowClient.focused".
//
// It returns ok=false if there is no such property.
func (this WindowClient) Focused() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowClientFocused(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AncestorOrigins returns the value of property "WindowClient.ancestorOrigins".
//
// It returns ok=false if there is no such property.
func (this WindowClient) AncestorOrigins() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetWindowClientAncestorOrigins(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFocus returns true if the method "WindowClient.focus" exists.
func (this WindowClient) HasFocus() bool {
	return js.True == bindings.HasWindowClientFocus(
		this.Ref(),
	)
}

// FocusFunc returns the method "WindowClient.focus".
func (this WindowClient) FocusFunc() (fn js.Func[func() js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.WindowClientFocusFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "WindowClient.focus".
func (this WindowClient) Focus() (ret js.Promise[WindowClient]) {
	bindings.CallWindowClientFocus(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "WindowClient.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WindowClient) TryFocus() (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClientFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasNavigate returns true if the method "WindowClient.navigate" exists.
func (this WindowClient) HasNavigate() bool {
	return js.True == bindings.HasWindowClientNavigate(
		this.Ref(),
	)
}

// NavigateFunc returns the method "WindowClient.navigate".
func (this WindowClient) NavigateFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.WindowClientNavigateFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "WindowClient.navigate".
func (this WindowClient) Navigate(url js.String) (ret js.Promise[WindowClient]) {
	bindings.CallWindowClientNavigate(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryNavigate calls the method "WindowClient.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WindowClient) TryNavigate(url js.String) (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClientNavigate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
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

// HasGet returns true if the method "Clients.get" exists.
func (this Clients) HasGet() bool {
	return js.True == bindings.HasClientsGet(
		this.Ref(),
	)
}

// GetFunc returns the method "Clients.get".
func (this Clients) GetFunc() (fn js.Func[func(id js.String) js.Promise[OneOf_Client_undefined]]) {
	return fn.FromRef(
		bindings.ClientsGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Clients.get".
func (this Clients) Get(id js.String) (ret js.Promise[OneOf_Client_undefined]) {
	bindings.CallClientsGet(
		this.Ref(), js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the method "Clients.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryGet(id js.String) (ret js.Promise[OneOf_Client_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasMatchAll returns true if the method "Clients.matchAll" exists.
func (this Clients) HasMatchAll() bool {
	return js.True == bindings.HasClientsMatchAll(
		this.Ref(),
	)
}

// MatchAllFunc returns the method "Clients.matchAll".
func (this Clients) MatchAllFunc() (fn js.Func[func(options ClientQueryOptions) js.Promise[js.FrozenArray[Client]]]) {
	return fn.FromRef(
		bindings.ClientsMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "Clients.matchAll".
func (this Clients) MatchAll(options ClientQueryOptions) (ret js.Promise[js.FrozenArray[Client]]) {
	bindings.CallClientsMatchAll(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryMatchAll calls the method "Clients.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryMatchAll(options ClientQueryOptions) (ret js.Promise[js.FrozenArray[Client]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsMatchAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasMatchAll1 returns true if the method "Clients.matchAll" exists.
func (this Clients) HasMatchAll1() bool {
	return js.True == bindings.HasClientsMatchAll1(
		this.Ref(),
	)
}

// MatchAll1Func returns the method "Clients.matchAll".
func (this Clients) MatchAll1Func() (fn js.Func[func() js.Promise[js.FrozenArray[Client]]]) {
	return fn.FromRef(
		bindings.ClientsMatchAll1Func(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "Clients.matchAll".
func (this Clients) MatchAll1() (ret js.Promise[js.FrozenArray[Client]]) {
	bindings.CallClientsMatchAll1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMatchAll1 calls the method "Clients.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryMatchAll1() (ret js.Promise[js.FrozenArray[Client]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsMatchAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpenWindow returns true if the method "Clients.openWindow" exists.
func (this Clients) HasOpenWindow() bool {
	return js.True == bindings.HasClientsOpenWindow(
		this.Ref(),
	)
}

// OpenWindowFunc returns the method "Clients.openWindow".
func (this Clients) OpenWindowFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.ClientsOpenWindowFunc(
			this.Ref(),
		),
	)
}

// OpenWindow calls the method "Clients.openWindow".
func (this Clients) OpenWindow(url js.String) (ret js.Promise[WindowClient]) {
	bindings.CallClientsOpenWindow(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpenWindow calls the method "Clients.openWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryOpenWindow(url js.String) (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsOpenWindow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasClaim returns true if the method "Clients.claim" exists.
func (this Clients) HasClaim() bool {
	return js.True == bindings.HasClientsClaim(
		this.Ref(),
	)
}

// ClaimFunc returns the method "Clients.claim".
func (this Clients) ClaimFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClientsClaimFunc(
			this.Ref(),
		),
	)
}

// Claim calls the method "Clients.claim".
func (this Clients) Claim() (ret js.Promise[js.Void]) {
	bindings.CallClientsClaim(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClaim calls the method "Clients.claim"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryClaim() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsClaim(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new ClipboardItemOptions in the application heap.
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

func NewClipboardItem(items js.Record[ClipboardItemData], options ClipboardItemOptions) (ret ClipboardItem) {
	ret.ref = bindings.NewClipboardItemByClipboardItem(
		items.Ref(),
		js.Pointer(&options))
	return
}

func NewClipboardItemByClipboardItem1(items js.Record[ClipboardItemData]) (ret ClipboardItem) {
	ret.ref = bindings.NewClipboardItemByClipboardItem1(
		items.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ClipboardItem) PresentationStyle() (ret PresentationStyle, ok bool) {
	ok = js.True == bindings.GetClipboardItemPresentationStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Types returns the value of property "ClipboardItem.types".
//
// It returns ok=false if there is no such property.
func (this ClipboardItem) Types() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetClipboardItemTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetType returns true if the method "ClipboardItem.getType" exists.
func (this ClipboardItem) HasGetType() bool {
	return js.True == bindings.HasClipboardItemGetType(
		this.Ref(),
	)
}

// GetTypeFunc returns the method "ClipboardItem.getType".
func (this ClipboardItem) GetTypeFunc() (fn js.Func[func(typ js.String) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ClipboardItemGetTypeFunc(
			this.Ref(),
		),
	)
}

// GetType calls the method "ClipboardItem.getType".
func (this ClipboardItem) GetType(typ js.String) (ret js.Promise[Blob]) {
	bindings.CallClipboardItemGetType(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetType calls the method "ClipboardItem.getType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ClipboardItem) TryGetType(typ js.String) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardItemGetType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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

// HasRead returns true if the method "Clipboard.read" exists.
func (this Clipboard) HasRead() bool {
	return js.True == bindings.HasClipboardRead(
		this.Ref(),
	)
}

// ReadFunc returns the method "Clipboard.read".
func (this Clipboard) ReadFunc() (fn js.Func[func() js.Promise[ClipboardItems]]) {
	return fn.FromRef(
		bindings.ClipboardReadFunc(
			this.Ref(),
		),
	)
}

// Read calls the method "Clipboard.read".
func (this Clipboard) Read() (ret js.Promise[ClipboardItems]) {
	bindings.CallClipboardRead(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRead calls the method "Clipboard.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryRead() (ret js.Promise[ClipboardItems], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardRead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReadText returns true if the method "Clipboard.readText" exists.
func (this Clipboard) HasReadText() bool {
	return js.True == bindings.HasClipboardReadText(
		this.Ref(),
	)
}

// ReadTextFunc returns the method "Clipboard.readText".
func (this Clipboard) ReadTextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.ClipboardReadTextFunc(
			this.Ref(),
		),
	)
}

// ReadText calls the method "Clipboard.readText".
func (this Clipboard) ReadText() (ret js.Promise[js.String]) {
	bindings.CallClipboardReadText(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReadText calls the method "Clipboard.readText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryReadText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardReadText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWrite returns true if the method "Clipboard.write" exists.
func (this Clipboard) HasWrite() bool {
	return js.True == bindings.HasClipboardWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "Clipboard.write".
func (this Clipboard) WriteFunc() (fn js.Func[func(data ClipboardItems) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClipboardWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "Clipboard.write".
func (this Clipboard) Write(data ClipboardItems) (ret js.Promise[js.Void]) {
	bindings.CallClipboardWrite(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWrite calls the method "Clipboard.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryWrite(data ClipboardItems) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardWrite(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasWriteText returns true if the method "Clipboard.writeText" exists.
func (this Clipboard) HasWriteText() bool {
	return js.True == bindings.HasClipboardWriteText(
		this.Ref(),
	)
}

// WriteTextFunc returns the method "Clipboard.writeText".
func (this Clipboard) WriteTextFunc() (fn js.Func[func(data js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ClipboardWriteTextFunc(
			this.Ref(),
		),
	)
}

// WriteText calls the method "Clipboard.writeText".
func (this Clipboard) WriteText(data js.String) (ret js.Promise[js.Void]) {
	bindings.CallClipboardWriteText(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWriteText calls the method "Clipboard.writeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryWriteText(data js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardWriteText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
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

// New creates a new FileSystemHandlePermissionDescriptor in the application heap.
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
// It returns ok=false if there is no such property.
func (this FileSystemHandle) Kind() (ret FileSystemHandleKind, ok bool) {
	ok = js.True == bindings.GetFileSystemHandleKind(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "FileSystemHandle.name".
//
// It returns ok=false if there is no such property.
func (this FileSystemHandle) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemHandleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasIsSameEntry returns true if the method "FileSystemHandle.isSameEntry" exists.
func (this FileSystemHandle) HasIsSameEntry() bool {
	return js.True == bindings.HasFileSystemHandleIsSameEntry(
		this.Ref(),
	)
}

// IsSameEntryFunc returns the method "FileSystemHandle.isSameEntry".
func (this FileSystemHandle) IsSameEntryFunc() (fn js.Func[func(other FileSystemHandle) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.FileSystemHandleIsSameEntryFunc(
			this.Ref(),
		),
	)
}

// IsSameEntry calls the method "FileSystemHandle.isSameEntry".
func (this FileSystemHandle) IsSameEntry(other FileSystemHandle) (ret js.Promise[js.Boolean]) {
	bindings.CallFileSystemHandleIsSameEntry(
		this.Ref(), js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryIsSameEntry calls the method "FileSystemHandle.isSameEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryIsSameEntry(other FileSystemHandle) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleIsSameEntry(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasQueryPermission returns true if the method "FileSystemHandle.queryPermission" exists.
func (this FileSystemHandle) HasQueryPermission() bool {
	return js.True == bindings.HasFileSystemHandleQueryPermission(
		this.Ref(),
	)
}

// QueryPermissionFunc returns the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermissionFunc() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleQueryPermissionFunc(
			this.Ref(),
		),
	)
}

// QueryPermission calls the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleQueryPermission(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryQueryPermission calls the method "FileSystemHandle.queryPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryQueryPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleQueryPermission(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasQueryPermission1 returns true if the method "FileSystemHandle.queryPermission" exists.
func (this FileSystemHandle) HasQueryPermission1() bool {
	return js.True == bindings.HasFileSystemHandleQueryPermission1(
		this.Ref(),
	)
}

// QueryPermission1Func returns the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermission1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleQueryPermission1Func(
			this.Ref(),
		),
	)
}

// QueryPermission1 calls the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermission1() (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleQueryPermission1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryQueryPermission1 calls the method "FileSystemHandle.queryPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryQueryPermission1() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleQueryPermission1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestPermission returns true if the method "FileSystemHandle.requestPermission" exists.
func (this FileSystemHandle) HasRequestPermission() bool {
	return js.True == bindings.HasFileSystemHandleRequestPermission(
		this.Ref(),
	)
}

// RequestPermissionFunc returns the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermissionFunc() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// RequestPermission calls the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleRequestPermission(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryRequestPermission calls the method "FileSystemHandle.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryRequestPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleRequestPermission(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasRequestPermission1 returns true if the method "FileSystemHandle.requestPermission" exists.
func (this FileSystemHandle) HasRequestPermission1() bool {
	return js.True == bindings.HasFileSystemHandleRequestPermission1(
		this.Ref(),
	)
}

// RequestPermission1Func returns the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermission1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.FileSystemHandleRequestPermission1Func(
			this.Ref(),
		),
	)
}

// RequestPermission1 calls the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermission1() (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleRequestPermission1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPermission1 calls the method "FileSystemHandle.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryRequestPermission1() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleRequestPermission1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasReadEntries returns true if the method "FileSystemDirectoryReader.readEntries" exists.
func (this FileSystemDirectoryReader) HasReadEntries() bool {
	return js.True == bindings.HasFileSystemDirectoryReaderReadEntries(
		this.Ref(),
	)
}

// ReadEntriesFunc returns the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntriesFunc() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryReaderReadEntriesFunc(
			this.Ref(),
		),
	)
}

// ReadEntries calls the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntries(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryReaderReadEntries(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryReadEntries calls the method "FileSystemDirectoryReader.readEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryReader) TryReadEntries(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryReaderReadEntries(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasReadEntries1 returns true if the method "FileSystemDirectoryReader.readEntries" exists.
func (this FileSystemDirectoryReader) HasReadEntries1() bool {
	return js.True == bindings.HasFileSystemDirectoryReaderReadEntries1(
		this.Ref(),
	)
}

// ReadEntries1Func returns the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntries1Func() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryReaderReadEntries1Func(
			this.Ref(),
		),
	)
}

// ReadEntries1 calls the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntries1(successCallback js.Func[func(entries js.Array[FileSystemEntry])]) (ret js.Void) {
	bindings.CallFileSystemDirectoryReaderReadEntries1(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryReadEntries1 calls the method "FileSystemDirectoryReader.readEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryReader) TryReadEntries1(successCallback js.Func[func(entries js.Array[FileSystemEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryReaderReadEntries1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
}

type FileSystemFlags struct {
	// Create is "FileSystemFlags.create"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Create MUST be set to true to make this field effective.
	Create bool
	// Exclusive is "FileSystemFlags.exclusive"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Exclusive MUST be set to true to make this field effective.
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

// New creates a new FileSystemFlags in the application heap.
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

// HasCreateReader returns true if the method "FileSystemDirectoryEntry.createReader" exists.
func (this FileSystemDirectoryEntry) HasCreateReader() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryCreateReader(
		this.Ref(),
	)
}

// CreateReaderFunc returns the method "FileSystemDirectoryEntry.createReader".
func (this FileSystemDirectoryEntry) CreateReaderFunc() (fn js.Func[func() FileSystemDirectoryReader]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryCreateReaderFunc(
			this.Ref(),
		),
	)
}

// CreateReader calls the method "FileSystemDirectoryEntry.createReader".
func (this FileSystemDirectoryEntry) CreateReader() (ret FileSystemDirectoryReader) {
	bindings.CallFileSystemDirectoryEntryCreateReader(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateReader calls the method "FileSystemDirectoryEntry.createReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryCreateReader() (ret FileSystemDirectoryReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryCreateReader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetFile returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasGetFile() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetFile(
		this.Ref(),
	)
}

// GetFileFunc returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFileFunc() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFileFunc(
			this.Ref(),
		),
	)
}

// GetFile calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetFile calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasGetFile1 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasGetFile1() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetFile1(
		this.Ref(),
	)
}

// GetFile1Func returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile1Func() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile1Func(
			this.Ref(),
		),
	)
}

// GetFile1 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// TryGetFile1 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// HasGetFile2 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasGetFile2() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetFile2(
		this.Ref(),
	)
}

// GetFile2Func returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile2Func() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile2Func(
			this.Ref(),
		),
	)
}

// GetFile2 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile2(path js.String, options FileSystemFlags) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetFile2 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile2(path js.String, options FileSystemFlags) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasGetFile3 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasGetFile3() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetFile3(
		this.Ref(),
	)
}

// GetFile3Func returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile3Func() (fn js.Func[func(path js.String)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile3Func(
			this.Ref(),
		),
	)
}

// GetFile3 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile3(path js.String) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetFile3 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile3(path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasGetFile4 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasGetFile4() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetFile4(
		this.Ref(),
	)
}

// GetFile4Func returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile4Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetFile4Func(
			this.Ref(),
		),
	)
}

// GetFile4 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile4() (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetFile4 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile4() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDirectory returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasGetDirectory() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetDirectory(
		this.Ref(),
	)
}

// GetDirectoryFunc returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectoryFunc() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectoryFunc(
			this.Ref(),
		),
	)
}

// GetDirectory calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetDirectory calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasGetDirectory1 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasGetDirectory1() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetDirectory1(
		this.Ref(),
	)
}

// GetDirectory1Func returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory1Func() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory1Func(
			this.Ref(),
		),
	)
}

// GetDirectory1 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// TryGetDirectory1 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// HasGetDirectory2 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasGetDirectory2() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetDirectory2(
		this.Ref(),
	)
}

// GetDirectory2Func returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory2Func() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory2Func(
			this.Ref(),
		),
	)
}

// GetDirectory2 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory2(path js.String, options FileSystemFlags) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetDirectory2 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory2(path js.String, options FileSystemFlags) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasGetDirectory3 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasGetDirectory3() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetDirectory3(
		this.Ref(),
	)
}

// GetDirectory3Func returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory3Func() (fn js.Func[func(path js.String)]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory3Func(
			this.Ref(),
		),
	)
}

// GetDirectory3 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory3(path js.String) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetDirectory3 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory3(path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasGetDirectory4 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasGetDirectory4() bool {
	return js.True == bindings.HasFileSystemDirectoryEntryGetDirectory4(
		this.Ref(),
	)
}

// GetDirectory4Func returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory4Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryEntryGetDirectory4Func(
			this.Ref(),
		),
	)
}

// GetDirectory4 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory4() (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDirectory4 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory4() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this FileSystem) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Root returns the value of property "FileSystem.root".
//
// It returns ok=false if there is no such property.
func (this FileSystem) Root() (ret FileSystemDirectoryEntry, ok bool) {
	ok = js.True == bindings.GetFileSystemRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this FileSystemEntry) IsFile() (ret bool, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryIsFile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsDirectory returns the value of property "FileSystemEntry.isDirectory".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) IsDirectory() (ret bool, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryIsDirectory(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "FileSystemEntry.name".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FullPath returns the value of property "FileSystemEntry.fullPath".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) FullPath() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryFullPath(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Filesystem returns the value of property "FileSystemEntry.filesystem".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) Filesystem() (ret FileSystem, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryFilesystem(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetParent returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasGetParent() bool {
	return js.True == bindings.HasFileSystemEntryGetParent(
		this.Ref(),
	)
}

// GetParentFunc returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParentFunc() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParentFunc(
			this.Ref(),
		),
	)
}

// GetParent calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemEntryGetParent(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetParent calls the method "FileSystemEntry.getParent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemEntry) TryGetParent(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemEntryGetParent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasGetParent1 returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasGetParent1() bool {
	return js.True == bindings.HasFileSystemEntryGetParent1(
		this.Ref(),
	)
}

// GetParent1Func returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent1Func() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)])]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParent1Func(
			this.Ref(),
		),
	)
}

// GetParent1 calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent1(successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemEntryGetParent1(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryGetParent1 calls the method "FileSystemEntry.getParent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemEntry) TryGetParent1(successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemEntryGetParent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
}

// HasGetParent2 returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasGetParent2() bool {
	return js.True == bindings.HasFileSystemEntryGetParent2(
		this.Ref(),
	)
}

// GetParent2Func returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemEntryGetParent2Func(
			this.Ref(),
		),
	)
}

// GetParent2 calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent2() (ret js.Void) {
	bindings.CallFileSystemEntryGetParent2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetParent2 calls the method "FileSystemEntry.getParent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemEntry) TryGetParent2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemEntryGetParent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DataTransferItem) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferItemKind(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "DataTransferItem.type".
//
// It returns ok=false if there is no such property.
func (this DataTransferItem) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferItemType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetAsString returns true if the method "DataTransferItem.getAsString" exists.
func (this DataTransferItem) HasGetAsString() bool {
	return js.True == bindings.HasDataTransferItemGetAsString(
		this.Ref(),
	)
}

// GetAsStringFunc returns the method "DataTransferItem.getAsString".
func (this DataTransferItem) GetAsStringFunc() (fn js.Func[func(callback js.Func[func(data js.String)])]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsStringFunc(
			this.Ref(),
		),
	)
}

// GetAsString calls the method "DataTransferItem.getAsString".
func (this DataTransferItem) GetAsString(callback js.Func[func(data js.String)]) (ret js.Void) {
	bindings.CallDataTransferItemGetAsString(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetAsString calls the method "DataTransferItem.getAsString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsString(callback js.Func[func(data js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasGetAsFile returns true if the method "DataTransferItem.getAsFile" exists.
func (this DataTransferItem) HasGetAsFile() bool {
	return js.True == bindings.HasDataTransferItemGetAsFile(
		this.Ref(),
	)
}

// GetAsFileFunc returns the method "DataTransferItem.getAsFile".
func (this DataTransferItem) GetAsFileFunc() (fn js.Func[func() File]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsFileFunc(
			this.Ref(),
		),
	)
}

// GetAsFile calls the method "DataTransferItem.getAsFile".
func (this DataTransferItem) GetAsFile() (ret File) {
	bindings.CallDataTransferItemGetAsFile(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAsFile calls the method "DataTransferItem.getAsFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsFile() (ret File, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsFile(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAsFileSystemHandle returns true if the method "DataTransferItem.getAsFileSystemHandle" exists.
func (this DataTransferItem) HasGetAsFileSystemHandle() bool {
	return js.True == bindings.HasDataTransferItemGetAsFileSystemHandle(
		this.Ref(),
	)
}

// GetAsFileSystemHandleFunc returns the method "DataTransferItem.getAsFileSystemHandle".
func (this DataTransferItem) GetAsFileSystemHandleFunc() (fn js.Func[func() js.Promise[FileSystemHandle]]) {
	return fn.FromRef(
		bindings.DataTransferItemGetAsFileSystemHandleFunc(
			this.Ref(),
		),
	)
}

// GetAsFileSystemHandle calls the method "DataTransferItem.getAsFileSystemHandle".
func (this DataTransferItem) GetAsFileSystemHandle() (ret js.Promise[FileSystemHandle]) {
	bindings.CallDataTransferItemGetAsFileSystemHandle(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAsFileSystemHandle calls the method "DataTransferItem.getAsFileSystemHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsFileSystemHandle() (ret js.Promise[FileSystemHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsFileSystemHandle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWebkitGetAsEntry returns true if the method "DataTransferItem.webkitGetAsEntry" exists.
func (this DataTransferItem) HasWebkitGetAsEntry() bool {
	return js.True == bindings.HasDataTransferItemWebkitGetAsEntry(
		this.Ref(),
	)
}

// WebkitGetAsEntryFunc returns the method "DataTransferItem.webkitGetAsEntry".
func (this DataTransferItem) WebkitGetAsEntryFunc() (fn js.Func[func() FileSystemEntry]) {
	return fn.FromRef(
		bindings.DataTransferItemWebkitGetAsEntryFunc(
			this.Ref(),
		),
	)
}

// WebkitGetAsEntry calls the method "DataTransferItem.webkitGetAsEntry".
func (this DataTransferItem) WebkitGetAsEntry() (ret FileSystemEntry) {
	bindings.CallDataTransferItemWebkitGetAsEntry(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryWebkitGetAsEntry calls the method "DataTransferItem.webkitGetAsEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryWebkitGetAsEntry() (ret FileSystemEntry, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemWebkitGetAsEntry(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DataTransferItemList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDataTransferItemListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "DataTransferItemList." exists.
func (this DataTransferItemList) HasGet() bool {
	return js.True == bindings.HasDataTransferItemListGet(
		this.Ref(),
	)
}

// GetFunc returns the method "DataTransferItemList.".
func (this DataTransferItemList) GetFunc() (fn js.Func[func(index uint32) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "DataTransferItemList.".
func (this DataTransferItemList) Get(index uint32) (ret DataTransferItem) {
	bindings.CallDataTransferItemListGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "DataTransferItemList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryGet(index uint32) (ret DataTransferItem, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAdd returns true if the method "DataTransferItemList.add" exists.
func (this DataTransferItemList) HasAdd() bool {
	return js.True == bindings.HasDataTransferItemListAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "DataTransferItemList.add".
func (this DataTransferItemList) AddFunc() (fn js.Func[func(data js.String, typ js.String) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "DataTransferItemList.add".
func (this DataTransferItemList) Add(data js.String, typ js.String) (ret DataTransferItem) {
	bindings.CallDataTransferItemListAdd(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		typ.Ref(),
	)

	return
}

// TryAdd calls the method "DataTransferItemList.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryAdd(data js.String, typ js.String) (ret DataTransferItem, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		typ.Ref(),
	)

	return
}

// HasAdd1 returns true if the method "DataTransferItemList.add" exists.
func (this DataTransferItemList) HasAdd1() bool {
	return js.True == bindings.HasDataTransferItemListAdd1(
		this.Ref(),
	)
}

// Add1Func returns the method "DataTransferItemList.add".
func (this DataTransferItemList) Add1Func() (fn js.Func[func(data File) DataTransferItem]) {
	return fn.FromRef(
		bindings.DataTransferItemListAdd1Func(
			this.Ref(),
		),
	)
}

// Add1 calls the method "DataTransferItemList.add".
func (this DataTransferItemList) Add1(data File) (ret DataTransferItem) {
	bindings.CallDataTransferItemListAdd1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAdd1 calls the method "DataTransferItemList.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryAdd1(data File) (ret DataTransferItem, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListAdd1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasRemove returns true if the method "DataTransferItemList.remove" exists.
func (this DataTransferItemList) HasRemove() bool {
	return js.True == bindings.HasDataTransferItemListRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "DataTransferItemList.remove".
func (this DataTransferItemList) RemoveFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.DataTransferItemListRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "DataTransferItemList.remove".
func (this DataTransferItemList) Remove(index uint32) (ret js.Void) {
	bindings.CallDataTransferItemListRemove(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemove calls the method "DataTransferItemList.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryRemove(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasClear returns true if the method "DataTransferItemList.clear" exists.
func (this DataTransferItemList) HasClear() bool {
	return js.True == bindings.HasDataTransferItemListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "DataTransferItemList.clear".
func (this DataTransferItemList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DataTransferItemListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "DataTransferItemList.clear".
func (this DataTransferItemList) Clear() (ret js.Void) {
	bindings.CallDataTransferItemListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "DataTransferItemList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this FileList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFileListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "FileList.item" exists.
func (this FileList) HasItem() bool {
	return js.True == bindings.HasFileListItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "FileList.item".
func (this FileList) ItemFunc() (fn js.Func[func(index uint32) File]) {
	return fn.FromRef(
		bindings.FileListItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "FileList.item".
func (this FileList) Item(index uint32) (ret File) {
	bindings.CallFileListItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "FileList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileList) TryItem(index uint32) (ret File, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileListItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DataTransfer) DropEffect() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferDropEffect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDropEffect sets the value of property "DataTransfer.dropEffect" to val.
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
// It returns ok=false if there is no such property.
func (this DataTransfer) EffectAllowed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferEffectAllowed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEffectAllowed sets the value of property "DataTransfer.effectAllowed" to val.
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
// It returns ok=false if there is no such property.
func (this DataTransfer) Items() (ret DataTransferItemList, ok bool) {
	ok = js.True == bindings.GetDataTransferItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Types returns the value of property "DataTransfer.types".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) Types() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetDataTransferTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "DataTransfer.files".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) Files() (ret FileList, ok bool) {
	ok = js.True == bindings.GetDataTransferFiles(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetDragImage returns true if the method "DataTransfer.setDragImage" exists.
func (this DataTransfer) HasSetDragImage() bool {
	return js.True == bindings.HasDataTransferSetDragImage(
		this.Ref(),
	)
}

// SetDragImageFunc returns the method "DataTransfer.setDragImage".
func (this DataTransfer) SetDragImageFunc() (fn js.Func[func(image Element, x int32, y int32)]) {
	return fn.FromRef(
		bindings.DataTransferSetDragImageFunc(
			this.Ref(),
		),
	)
}

// SetDragImage calls the method "DataTransfer.setDragImage".
func (this DataTransfer) SetDragImage(image Element, x int32, y int32) (ret js.Void) {
	bindings.CallDataTransferSetDragImage(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// TrySetDragImage calls the method "DataTransfer.setDragImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TrySetDragImage(image Element, x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferSetDragImage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasGetData returns true if the method "DataTransfer.getData" exists.
func (this DataTransfer) HasGetData() bool {
	return js.True == bindings.HasDataTransferGetData(
		this.Ref(),
	)
}

// GetDataFunc returns the method "DataTransfer.getData".
func (this DataTransfer) GetDataFunc() (fn js.Func[func(format js.String) js.String]) {
	return fn.FromRef(
		bindings.DataTransferGetDataFunc(
			this.Ref(),
		),
	)
}

// GetData calls the method "DataTransfer.getData".
func (this DataTransfer) GetData(format js.String) (ret js.String) {
	bindings.CallDataTransferGetData(
		this.Ref(), js.Pointer(&ret),
		format.Ref(),
	)

	return
}

// TryGetData calls the method "DataTransfer.getData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryGetData(format js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferGetData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
	)

	return
}

// HasSetData returns true if the method "DataTransfer.setData" exists.
func (this DataTransfer) HasSetData() bool {
	return js.True == bindings.HasDataTransferSetData(
		this.Ref(),
	)
}

// SetDataFunc returns the method "DataTransfer.setData".
func (this DataTransfer) SetDataFunc() (fn js.Func[func(format js.String, data js.String)]) {
	return fn.FromRef(
		bindings.DataTransferSetDataFunc(
			this.Ref(),
		),
	)
}

// SetData calls the method "DataTransfer.setData".
func (this DataTransfer) SetData(format js.String, data js.String) (ret js.Void) {
	bindings.CallDataTransferSetData(
		this.Ref(), js.Pointer(&ret),
		format.Ref(),
		data.Ref(),
	)

	return
}

// TrySetData calls the method "DataTransfer.setData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TrySetData(format js.String, data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferSetData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
		data.Ref(),
	)

	return
}

// HasClearData returns true if the method "DataTransfer.clearData" exists.
func (this DataTransfer) HasClearData() bool {
	return js.True == bindings.HasDataTransferClearData(
		this.Ref(),
	)
}

// ClearDataFunc returns the method "DataTransfer.clearData".
func (this DataTransfer) ClearDataFunc() (fn js.Func[func(format js.String)]) {
	return fn.FromRef(
		bindings.DataTransferClearDataFunc(
			this.Ref(),
		),
	)
}

// ClearData calls the method "DataTransfer.clearData".
func (this DataTransfer) ClearData(format js.String) (ret js.Void) {
	bindings.CallDataTransferClearData(
		this.Ref(), js.Pointer(&ret),
		format.Ref(),
	)

	return
}

// TryClearData calls the method "DataTransfer.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryClearData(format js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferClearData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
	)

	return
}

// HasClearData1 returns true if the method "DataTransfer.clearData" exists.
func (this DataTransfer) HasClearData1() bool {
	return js.True == bindings.HasDataTransferClearData1(
		this.Ref(),
	)
}

// ClearData1Func returns the method "DataTransfer.clearData".
func (this DataTransfer) ClearData1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DataTransferClearData1Func(
			this.Ref(),
		),
	)
}

// ClearData1 calls the method "DataTransfer.clearData".
func (this DataTransfer) ClearData1() (ret js.Void) {
	bindings.CallDataTransferClearData1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearData1 calls the method "DataTransfer.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryClearData1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferClearData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ClipboardEventInit struct {
	// ClipboardData is "ClipboardEventInit.clipboardData"
	//
	// Optional, defaults to null.
	ClipboardData DataTransfer
	// Bubbles is "ClipboardEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ClipboardEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ClipboardEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new ClipboardEventInit in the application heap.
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

func NewClipboardEvent(typ js.String, eventInitDict ClipboardEventInit) (ret ClipboardEvent) {
	ret.ref = bindings.NewClipboardEventByClipboardEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewClipboardEventByClipboardEvent1(typ js.String) (ret ClipboardEvent) {
	ret.ref = bindings.NewClipboardEventByClipboardEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ClipboardEvent) ClipboardData() (ret DataTransfer, ok bool) {
	ok = js.True == bindings.GetClipboardEventClipboardData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ClipboardPermissionDescriptor struct {
	// AllowWithoutGesture is "ClipboardPermissionDescriptor.allowWithoutGesture"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AllowWithoutGesture MUST be set to true to make this field effective.
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

// New creates a new ClipboardPermissionDescriptor in the application heap.
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
	//
	// NOTE: FFI_USE_WasClean MUST be set to true to make this field effective.
	WasClean bool
	// Code is "CloseEventInit.code"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Code MUST be set to true to make this field effective.
	Code uint16
	// Reason is "CloseEventInit.reason"
	//
	// Optional, defaults to "".
	Reason js.String
	// Bubbles is "CloseEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CloseEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CloseEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new CloseEventInit in the application heap.
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

func NewCloseEvent(typ js.String, eventInitDict CloseEventInit) (ret CloseEvent) {
	ret.ref = bindings.NewCloseEventByCloseEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewCloseEventByCloseEvent1(typ js.String) (ret CloseEvent) {
	ret.ref = bindings.NewCloseEventByCloseEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CloseEvent) WasClean() (ret bool, ok bool) {
	ok = js.True == bindings.GetCloseEventWasClean(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Code returns the value of property "CloseEvent.code".
//
// It returns ok=false if there is no such property.
func (this CloseEvent) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCloseEventCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Reason returns the value of property "CloseEvent.reason".
//
// It returns ok=false if there is no such property.
func (this CloseEvent) Reason() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCloseEventReason(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentCurrencyAmount
	// Instrument is "CollectedClientAdditionalPaymentData.instrument"
	//
	// Required
	//
	// NOTE: Instrument.FFI_USE MUST be set to true to get Instrument used.
	Instrument PaymentCredentialInstrument

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientAdditionalPaymentData with all fields set.
func (p CollectedClientAdditionalPaymentData) FromRef(ref js.Ref) CollectedClientAdditionalPaymentData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CollectedClientAdditionalPaymentData in the application heap.
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
	//
	// NOTE: FFI_USE_CrossOrigin MUST be set to true to make this field effective.
	CrossOrigin bool

	FFI_USE_CrossOrigin bool // for CrossOrigin.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientData with all fields set.
func (p CollectedClientData) FromRef(ref js.Ref) CollectedClientData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CollectedClientData in the application heap.
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
	//
	// NOTE: Payment.FFI_USE MUST be set to true to get Payment used.
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
	//
	// NOTE: FFI_USE_CrossOrigin MUST be set to true to make this field effective.
	CrossOrigin bool

	FFI_USE_CrossOrigin bool // for CrossOrigin.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CollectedClientPaymentData with all fields set.
func (p CollectedClientPaymentData) FromRef(ref js.Ref) CollectedClientPaymentData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CollectedClientPaymentData in the application heap.
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

// New creates a new ColorSelectionOptions in the application heap.
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

// New creates a new ColorSelectionResult in the application heap.
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

// New creates a new WindowPostMessageOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this MediaQueryList) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaQueryListMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Matches returns the value of property "MediaQueryList.matches".
//
// It returns ok=false if there is no such property.
func (this MediaQueryList) Matches() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaQueryListMatches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAddListener returns true if the method "MediaQueryList.addListener" exists.
func (this MediaQueryList) HasAddListener() bool {
	return js.True == bindings.HasMediaQueryListAddListener(
		this.Ref(),
	)
}

// AddListenerFunc returns the method "MediaQueryList.addListener".
func (this MediaQueryList) AddListenerFunc() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.MediaQueryListAddListenerFunc(
			this.Ref(),
		),
	)
}

// AddListener calls the method "MediaQueryList.addListener".
func (this MediaQueryList) AddListener(callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallMediaQueryListAddListener(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryAddListener calls the method "MediaQueryList.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaQueryList) TryAddListener(callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaQueryListAddListener(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasRemoveListener returns true if the method "MediaQueryList.removeListener" exists.
func (this MediaQueryList) HasRemoveListener() bool {
	return js.True == bindings.HasMediaQueryListRemoveListener(
		this.Ref(),
	)
}

// RemoveListenerFunc returns the method "MediaQueryList.removeListener".
func (this MediaQueryList) RemoveListenerFunc() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.MediaQueryListRemoveListenerFunc(
			this.Ref(),
		),
	)
}

// RemoveListener calls the method "MediaQueryList.removeListener".
func (this MediaQueryList) RemoveListener(callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallMediaQueryListRemoveListener(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRemoveListener calls the method "MediaQueryList.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaQueryList) TryRemoveListener(callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaQueryListRemoveListener(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

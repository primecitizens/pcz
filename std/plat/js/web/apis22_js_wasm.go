// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// VisibilityState returns the value of property "WindowClient.visibilityState".
//
// It returns ok=false if there is no such property.
func (this WindowClient) VisibilityState() (ret DocumentVisibilityState, ok bool) {
	ok = js.True == bindings.GetWindowClientVisibilityState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Focused returns the value of property "WindowClient.focused".
//
// It returns ok=false if there is no such property.
func (this WindowClient) Focused() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowClientFocused(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AncestorOrigins returns the value of property "WindowClient.ancestorOrigins".
//
// It returns ok=false if there is no such property.
func (this WindowClient) AncestorOrigins() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetWindowClientAncestorOrigins(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFocus returns true if the method "WindowClient.focus" exists.
func (this WindowClient) HasFuncFocus() bool {
	return js.True == bindings.HasFuncWindowClientFocus(
		this.ref,
	)
}

// FuncFocus returns the method "WindowClient.focus".
func (this WindowClient) FuncFocus() (fn js.Func[func() js.Promise[WindowClient]]) {
	bindings.FuncWindowClientFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "WindowClient.focus".
func (this WindowClient) Focus() (ret js.Promise[WindowClient]) {
	bindings.CallWindowClientFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "WindowClient.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WindowClient) TryFocus() (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClientFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNavigate returns true if the method "WindowClient.navigate" exists.
func (this WindowClient) HasFuncNavigate() bool {
	return js.True == bindings.HasFuncWindowClientNavigate(
		this.ref,
	)
}

// FuncNavigate returns the method "WindowClient.navigate".
func (this WindowClient) FuncNavigate() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	bindings.FuncWindowClientNavigate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Navigate calls the method "WindowClient.navigate".
func (this WindowClient) Navigate(url js.String) (ret js.Promise[WindowClient]) {
	bindings.CallWindowClientNavigate(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryNavigate calls the method "WindowClient.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WindowClient) TryNavigate(url js.String) (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClientNavigate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

type Clients struct {
	ref js.Ref
}

func (this Clients) Once() Clients {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGet returns true if the method "Clients.get" exists.
func (this Clients) HasFuncGet() bool {
	return js.True == bindings.HasFuncClientsGet(
		this.ref,
	)
}

// FuncGet returns the method "Clients.get".
func (this Clients) FuncGet() (fn js.Func[func(id js.String) js.Promise[OneOf_Client_undefined]]) {
	bindings.FuncClientsGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "Clients.get".
func (this Clients) Get(id js.String) (ret js.Promise[OneOf_Client_undefined]) {
	bindings.CallClientsGet(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the method "Clients.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryGet(id js.String) (ret js.Promise[OneOf_Client_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncMatchAll returns true if the method "Clients.matchAll" exists.
func (this Clients) HasFuncMatchAll() bool {
	return js.True == bindings.HasFuncClientsMatchAll(
		this.ref,
	)
}

// FuncMatchAll returns the method "Clients.matchAll".
func (this Clients) FuncMatchAll() (fn js.Func[func(options ClientQueryOptions) js.Promise[js.FrozenArray[Client]]]) {
	bindings.FuncClientsMatchAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchAll calls the method "Clients.matchAll".
func (this Clients) MatchAll(options ClientQueryOptions) (ret js.Promise[js.FrozenArray[Client]]) {
	bindings.CallClientsMatchAll(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryMatchAll calls the method "Clients.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryMatchAll(options ClientQueryOptions) (ret js.Promise[js.FrozenArray[Client]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsMatchAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncMatchAll1 returns true if the method "Clients.matchAll" exists.
func (this Clients) HasFuncMatchAll1() bool {
	return js.True == bindings.HasFuncClientsMatchAll1(
		this.ref,
	)
}

// FuncMatchAll1 returns the method "Clients.matchAll".
func (this Clients) FuncMatchAll1() (fn js.Func[func() js.Promise[js.FrozenArray[Client]]]) {
	bindings.FuncClientsMatchAll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchAll1 calls the method "Clients.matchAll".
func (this Clients) MatchAll1() (ret js.Promise[js.FrozenArray[Client]]) {
	bindings.CallClientsMatchAll1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMatchAll1 calls the method "Clients.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryMatchAll1() (ret js.Promise[js.FrozenArray[Client]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsMatchAll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenWindow returns true if the method "Clients.openWindow" exists.
func (this Clients) HasFuncOpenWindow() bool {
	return js.True == bindings.HasFuncClientsOpenWindow(
		this.ref,
	)
}

// FuncOpenWindow returns the method "Clients.openWindow".
func (this Clients) FuncOpenWindow() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	bindings.FuncClientsOpenWindow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenWindow calls the method "Clients.openWindow".
func (this Clients) OpenWindow(url js.String) (ret js.Promise[WindowClient]) {
	bindings.CallClientsOpenWindow(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpenWindow calls the method "Clients.openWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryOpenWindow(url js.String) (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsOpenWindow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncClaim returns true if the method "Clients.claim" exists.
func (this Clients) HasFuncClaim() bool {
	return js.True == bindings.HasFuncClientsClaim(
		this.ref,
	)
}

// FuncClaim returns the method "Clients.claim".
func (this Clients) FuncClaim() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncClientsClaim(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Claim calls the method "Clients.claim".
func (this Clients) Claim() (ret js.Promise[js.Void]) {
	bindings.CallClientsClaim(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClaim calls the method "Clients.claim"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clients) TryClaim() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientsClaim(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ClipboardItemOptions) UpdateFrom(ref js.Ref) {
	bindings.ClipboardItemOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClipboardItemOptions) Update(ref js.Ref) {
	bindings.ClipboardItemOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClipboardItemOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// PresentationStyle returns the value of property "ClipboardItem.presentationStyle".
//
// It returns ok=false if there is no such property.
func (this ClipboardItem) PresentationStyle() (ret PresentationStyle, ok bool) {
	ok = js.True == bindings.GetClipboardItemPresentationStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Types returns the value of property "ClipboardItem.types".
//
// It returns ok=false if there is no such property.
func (this ClipboardItem) Types() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetClipboardItemTypes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetType returns true if the method "ClipboardItem.getType" exists.
func (this ClipboardItem) HasFuncGetType() bool {
	return js.True == bindings.HasFuncClipboardItemGetType(
		this.ref,
	)
}

// FuncGetType returns the method "ClipboardItem.getType".
func (this ClipboardItem) FuncGetType() (fn js.Func[func(typ js.String) js.Promise[Blob]]) {
	bindings.FuncClipboardItemGetType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetType calls the method "ClipboardItem.getType".
func (this ClipboardItem) GetType(typ js.String) (ret js.Promise[Blob]) {
	bindings.CallClipboardItemGetType(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetType calls the method "ClipboardItem.getType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ClipboardItem) TryGetType(typ js.String) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardItemGetType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

type ClipboardItems = js.Array[ClipboardItem]

type Clipboard struct {
	EventTarget
}

func (this Clipboard) Once() Clipboard {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncRead returns true if the method "Clipboard.read" exists.
func (this Clipboard) HasFuncRead() bool {
	return js.True == bindings.HasFuncClipboardRead(
		this.ref,
	)
}

// FuncRead returns the method "Clipboard.read".
func (this Clipboard) FuncRead() (fn js.Func[func() js.Promise[ClipboardItems]]) {
	bindings.FuncClipboardRead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read calls the method "Clipboard.read".
func (this Clipboard) Read() (ret js.Promise[ClipboardItems]) {
	bindings.CallClipboardRead(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRead calls the method "Clipboard.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryRead() (ret js.Promise[ClipboardItems], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardRead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReadText returns true if the method "Clipboard.readText" exists.
func (this Clipboard) HasFuncReadText() bool {
	return js.True == bindings.HasFuncClipboardReadText(
		this.ref,
	)
}

// FuncReadText returns the method "Clipboard.readText".
func (this Clipboard) FuncReadText() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncClipboardReadText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadText calls the method "Clipboard.readText".
func (this Clipboard) ReadText() (ret js.Promise[js.String]) {
	bindings.CallClipboardReadText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReadText calls the method "Clipboard.readText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryReadText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardReadText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWrite returns true if the method "Clipboard.write" exists.
func (this Clipboard) HasFuncWrite() bool {
	return js.True == bindings.HasFuncClipboardWrite(
		this.ref,
	)
}

// FuncWrite returns the method "Clipboard.write".
func (this Clipboard) FuncWrite() (fn js.Func[func(data ClipboardItems) js.Promise[js.Void]]) {
	bindings.FuncClipboardWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "Clipboard.write".
func (this Clipboard) Write(data ClipboardItems) (ret js.Promise[js.Void]) {
	bindings.CallClipboardWrite(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWrite calls the method "Clipboard.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryWrite(data ClipboardItems) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncWriteText returns true if the method "Clipboard.writeText" exists.
func (this Clipboard) HasFuncWriteText() bool {
	return js.True == bindings.HasFuncClipboardWriteText(
		this.ref,
	)
}

// FuncWriteText returns the method "Clipboard.writeText".
func (this Clipboard) FuncWriteText() (fn js.Func[func(data js.String) js.Promise[js.Void]]) {
	bindings.FuncClipboardWriteText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteText calls the method "Clipboard.writeText".
func (this Clipboard) WriteText(data js.String) (ret js.Promise[js.Void]) {
	bindings.CallClipboardWriteText(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWriteText calls the method "Clipboard.writeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Clipboard) TryWriteText(data js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClipboardWriteText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *FileSystemHandlePermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.FileSystemHandlePermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemHandlePermissionDescriptor) Update(ref js.Ref) {
	bindings.FileSystemHandlePermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemHandlePermissionDescriptor) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Kind returns the value of property "FileSystemHandle.kind".
//
// It returns ok=false if there is no such property.
func (this FileSystemHandle) Kind() (ret FileSystemHandleKind, ok bool) {
	ok = js.True == bindings.GetFileSystemHandleKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "FileSystemHandle.name".
//
// It returns ok=false if there is no such property.
func (this FileSystemHandle) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemHandleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncIsSameEntry returns true if the method "FileSystemHandle.isSameEntry" exists.
func (this FileSystemHandle) HasFuncIsSameEntry() bool {
	return js.True == bindings.HasFuncFileSystemHandleIsSameEntry(
		this.ref,
	)
}

// FuncIsSameEntry returns the method "FileSystemHandle.isSameEntry".
func (this FileSystemHandle) FuncIsSameEntry() (fn js.Func[func(other FileSystemHandle) js.Promise[js.Boolean]]) {
	bindings.FuncFileSystemHandleIsSameEntry(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsSameEntry calls the method "FileSystemHandle.isSameEntry".
func (this FileSystemHandle) IsSameEntry(other FileSystemHandle) (ret js.Promise[js.Boolean]) {
	bindings.CallFileSystemHandleIsSameEntry(
		this.ref, js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryIsSameEntry calls the method "FileSystemHandle.isSameEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryIsSameEntry(other FileSystemHandle) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleIsSameEntry(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasFuncQueryPermission returns true if the method "FileSystemHandle.queryPermission" exists.
func (this FileSystemHandle) HasFuncQueryPermission() bool {
	return js.True == bindings.HasFuncFileSystemHandleQueryPermission(
		this.ref,
	)
}

// FuncQueryPermission returns the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) FuncQueryPermission() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	bindings.FuncFileSystemHandleQueryPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryPermission calls the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleQueryPermission(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryQueryPermission calls the method "FileSystemHandle.queryPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryQueryPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleQueryPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncQueryPermission1 returns true if the method "FileSystemHandle.queryPermission" exists.
func (this FileSystemHandle) HasFuncQueryPermission1() bool {
	return js.True == bindings.HasFuncFileSystemHandleQueryPermission1(
		this.ref,
	)
}

// FuncQueryPermission1 returns the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) FuncQueryPermission1() (fn js.Func[func() js.Promise[PermissionState]]) {
	bindings.FuncFileSystemHandleQueryPermission1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryPermission1 calls the method "FileSystemHandle.queryPermission".
func (this FileSystemHandle) QueryPermission1() (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleQueryPermission1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryQueryPermission1 calls the method "FileSystemHandle.queryPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryQueryPermission1() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleQueryPermission1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestPermission returns true if the method "FileSystemHandle.requestPermission" exists.
func (this FileSystemHandle) HasFuncRequestPermission() bool {
	return js.True == bindings.HasFuncFileSystemHandleRequestPermission(
		this.ref,
	)
}

// FuncRequestPermission returns the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) FuncRequestPermission() (fn js.Func[func(descriptor FileSystemHandlePermissionDescriptor) js.Promise[PermissionState]]) {
	bindings.FuncFileSystemHandleRequestPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission calls the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleRequestPermission(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryRequestPermission calls the method "FileSystemHandle.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryRequestPermission(descriptor FileSystemHandlePermissionDescriptor) (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleRequestPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncRequestPermission1 returns true if the method "FileSystemHandle.requestPermission" exists.
func (this FileSystemHandle) HasFuncRequestPermission1() bool {
	return js.True == bindings.HasFuncFileSystemHandleRequestPermission1(
		this.ref,
	)
}

// FuncRequestPermission1 returns the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) FuncRequestPermission1() (fn js.Func[func() js.Promise[PermissionState]]) {
	bindings.FuncFileSystemHandleRequestPermission1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission1 calls the method "FileSystemHandle.requestPermission".
func (this FileSystemHandle) RequestPermission1() (ret js.Promise[PermissionState]) {
	bindings.CallFileSystemHandleRequestPermission1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPermission1 calls the method "FileSystemHandle.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemHandle) TryRequestPermission1() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemHandleRequestPermission1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncReadEntries returns true if the method "FileSystemDirectoryReader.readEntries" exists.
func (this FileSystemDirectoryReader) HasFuncReadEntries() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryReaderReadEntries(
		this.ref,
	)
}

// FuncReadEntries returns the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) FuncReadEntries() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)])]) {
	bindings.FuncFileSystemDirectoryReaderReadEntries(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadEntries calls the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntries(successCallback js.Func[func(entries js.Array[FileSystemEntry])], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryReaderReadEntries(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncReadEntries1 returns true if the method "FileSystemDirectoryReader.readEntries" exists.
func (this FileSystemDirectoryReader) HasFuncReadEntries1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryReaderReadEntries1(
		this.ref,
	)
}

// FuncReadEntries1 returns the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) FuncReadEntries1() (fn js.Func[func(successCallback js.Func[func(entries js.Array[FileSystemEntry])])]) {
	bindings.FuncFileSystemDirectoryReaderReadEntries1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadEntries1 calls the method "FileSystemDirectoryReader.readEntries".
func (this FileSystemDirectoryReader) ReadEntries1(successCallback js.Func[func(entries js.Array[FileSystemEntry])]) (ret js.Void) {
	bindings.CallFileSystemDirectoryReaderReadEntries1(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryReadEntries1 calls the method "FileSystemDirectoryReader.readEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryReader) TryReadEntries1(successCallback js.Func[func(entries js.Array[FileSystemEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryReaderReadEntries1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *FileSystemFlags) UpdateFrom(ref js.Ref) {
	bindings.FileSystemFlagsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemFlags) Update(ref js.Ref) {
	bindings.FileSystemFlagsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemFlags) FreeMembers(recursive bool) {
}

type FileSystemDirectoryEntry struct {
	FileSystemEntry
}

func (this FileSystemDirectoryEntry) Once() FileSystemDirectoryEntry {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncCreateReader returns true if the method "FileSystemDirectoryEntry.createReader" exists.
func (this FileSystemDirectoryEntry) HasFuncCreateReader() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryCreateReader(
		this.ref,
	)
}

// FuncCreateReader returns the method "FileSystemDirectoryEntry.createReader".
func (this FileSystemDirectoryEntry) FuncCreateReader() (fn js.Func[func() FileSystemDirectoryReader]) {
	bindings.FuncFileSystemDirectoryEntryCreateReader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateReader calls the method "FileSystemDirectoryEntry.createReader".
func (this FileSystemDirectoryEntry) CreateReader() (ret FileSystemDirectoryReader) {
	bindings.CallFileSystemDirectoryEntryCreateReader(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateReader calls the method "FileSystemDirectoryEntry.createReader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryCreateReader() (ret FileSystemDirectoryReader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryCreateReader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetFile returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasFuncGetFile() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetFile(
		this.ref,
	)
}

// FuncGetFile returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) FuncGetFile() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	bindings.FuncFileSystemDirectoryEntryGetFile(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncGetFile1 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasFuncGetFile1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetFile1(
		this.ref,
	)
}

// FuncGetFile1 returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) FuncGetFile1() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	bindings.FuncFileSystemDirectoryEntryGetFile1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile1 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// HasFuncGetFile2 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasFuncGetFile2() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetFile2(
		this.ref,
	)
}

// FuncGetFile2 returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) FuncGetFile2() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	bindings.FuncFileSystemDirectoryEntryGetFile2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile2 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile2(path js.String, options FileSystemFlags) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetFile3 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasFuncGetFile3() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetFile3(
		this.ref,
	)
}

// FuncGetFile3 returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) FuncGetFile3() (fn js.Func[func(path js.String)]) {
	bindings.FuncFileSystemDirectoryEntryGetFile3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile3 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile3(path js.String) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile3(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetFile3 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile3(path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncGetFile4 returns true if the method "FileSystemDirectoryEntry.getFile" exists.
func (this FileSystemDirectoryEntry) HasFuncGetFile4() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetFile4(
		this.ref,
	)
}

// FuncGetFile4 returns the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) FuncGetFile4() (fn js.Func[func()]) {
	bindings.FuncFileSystemDirectoryEntryGetFile4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile4 calls the method "FileSystemDirectoryEntry.getFile".
func (this FileSystemDirectoryEntry) GetFile4() (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetFile4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetFile4 calls the method "FileSystemDirectoryEntry.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetFile4() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetFile4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDirectory returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasFuncGetDirectory() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetDirectory(
		this.ref,
	)
}

// FuncGetDirectory returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) FuncGetDirectory() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	bindings.FuncFileSystemDirectoryEntryGetDirectory(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncGetDirectory1 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasFuncGetDirectory1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetDirectory1(
		this.ref,
	)
}

// FuncGetDirectory1 returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) FuncGetDirectory1() (fn js.Func[func(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)])]) {
	bindings.FuncFileSystemDirectoryEntryGetDirectory1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory1 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory1(path js.String, options FileSystemFlags, successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
		successCallback.Ref(),
	)

	return
}

// HasFuncGetDirectory2 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasFuncGetDirectory2() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetDirectory2(
		this.ref,
	)
}

// FuncGetDirectory2 returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) FuncGetDirectory2() (fn js.Func[func(path js.String, options FileSystemFlags)]) {
	bindings.FuncFileSystemDirectoryEntryGetDirectory2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory2 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory2(path js.String, options FileSystemFlags) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetDirectory3 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasFuncGetDirectory3() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetDirectory3(
		this.ref,
	)
}

// FuncGetDirectory3 returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) FuncGetDirectory3() (fn js.Func[func(path js.String)]) {
	bindings.FuncFileSystemDirectoryEntryGetDirectory3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory3 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory3(path js.String) (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory3(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetDirectory3 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory3(path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncGetDirectory4 returns true if the method "FileSystemDirectoryEntry.getDirectory" exists.
func (this FileSystemDirectoryEntry) HasFuncGetDirectory4() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryEntryGetDirectory4(
		this.ref,
	)
}

// FuncGetDirectory4 returns the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) FuncGetDirectory4() (fn js.Func[func()]) {
	bindings.FuncFileSystemDirectoryEntryGetDirectory4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectory4 calls the method "FileSystemDirectoryEntry.getDirectory".
func (this FileSystemDirectoryEntry) GetDirectory4() (ret js.Void) {
	bindings.CallFileSystemDirectoryEntryGetDirectory4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDirectory4 calls the method "FileSystemDirectoryEntry.getDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryEntry) TryGetDirectory4() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryEntryGetDirectory4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FileSystem struct {
	ref js.Ref
}

func (this FileSystem) Once() FileSystem {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "FileSystem.name".
//
// It returns ok=false if there is no such property.
func (this FileSystem) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Root returns the value of property "FileSystem.root".
//
// It returns ok=false if there is no such property.
func (this FileSystem) Root() (ret FileSystemDirectoryEntry, ok bool) {
	ok = js.True == bindings.GetFileSystemRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

type FileSystemEntry struct {
	ref js.Ref
}

func (this FileSystemEntry) Once() FileSystemEntry {
	this.ref.Once()
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
	this.ref.Free()
}

// IsFile returns the value of property "FileSystemEntry.isFile".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) IsFile() (ret bool, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryIsFile(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsDirectory returns the value of property "FileSystemEntry.isDirectory".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) IsDirectory() (ret bool, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryIsDirectory(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "FileSystemEntry.name".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FullPath returns the value of property "FileSystemEntry.fullPath".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) FullPath() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryFullPath(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Filesystem returns the value of property "FileSystemEntry.filesystem".
//
// It returns ok=false if there is no such property.
func (this FileSystemEntry) Filesystem() (ret FileSystem, ok bool) {
	ok = js.True == bindings.GetFileSystemEntryFilesystem(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetParent returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasFuncGetParent() bool {
	return js.True == bindings.HasFuncFileSystemEntryGetParent(
		this.ref,
	)
}

// FuncGetParent returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) FuncGetParent() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)])]) {
	bindings.FuncFileSystemEntryGetParent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParent calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent(successCallback js.Func[func(entry FileSystemEntry)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemEntryGetParent(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncGetParent1 returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasFuncGetParent1() bool {
	return js.True == bindings.HasFuncFileSystemEntryGetParent1(
		this.ref,
	)
}

// FuncGetParent1 returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) FuncGetParent1() (fn js.Func[func(successCallback js.Func[func(entry FileSystemEntry)])]) {
	bindings.FuncFileSystemEntryGetParent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParent1 calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent1(successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void) {
	bindings.CallFileSystemEntryGetParent1(
		this.ref, js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryGetParent1 calls the method "FileSystemEntry.getParent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemEntry) TryGetParent1(successCallback js.Func[func(entry FileSystemEntry)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemEntryGetParent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
}

// HasFuncGetParent2 returns true if the method "FileSystemEntry.getParent" exists.
func (this FileSystemEntry) HasFuncGetParent2() bool {
	return js.True == bindings.HasFuncFileSystemEntryGetParent2(
		this.ref,
	)
}

// FuncGetParent2 returns the method "FileSystemEntry.getParent".
func (this FileSystemEntry) FuncGetParent2() (fn js.Func[func()]) {
	bindings.FuncFileSystemEntryGetParent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParent2 calls the method "FileSystemEntry.getParent".
func (this FileSystemEntry) GetParent2() (ret js.Void) {
	bindings.CallFileSystemEntryGetParent2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetParent2 calls the method "FileSystemEntry.getParent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemEntry) TryGetParent2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemEntryGetParent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DataTransferItem struct {
	ref js.Ref
}

func (this DataTransferItem) Once() DataTransferItem {
	this.ref.Once()
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
	this.ref.Free()
}

// Kind returns the value of property "DataTransferItem.kind".
//
// It returns ok=false if there is no such property.
func (this DataTransferItem) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferItemKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "DataTransferItem.type".
//
// It returns ok=false if there is no such property.
func (this DataTransferItem) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferItemType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetAsString returns true if the method "DataTransferItem.getAsString" exists.
func (this DataTransferItem) HasFuncGetAsString() bool {
	return js.True == bindings.HasFuncDataTransferItemGetAsString(
		this.ref,
	)
}

// FuncGetAsString returns the method "DataTransferItem.getAsString".
func (this DataTransferItem) FuncGetAsString() (fn js.Func[func(callback js.Func[func(data js.String)])]) {
	bindings.FuncDataTransferItemGetAsString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAsString calls the method "DataTransferItem.getAsString".
func (this DataTransferItem) GetAsString(callback js.Func[func(data js.String)]) (ret js.Void) {
	bindings.CallDataTransferItemGetAsString(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetAsString calls the method "DataTransferItem.getAsString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsString(callback js.Func[func(data js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetAsFile returns true if the method "DataTransferItem.getAsFile" exists.
func (this DataTransferItem) HasFuncGetAsFile() bool {
	return js.True == bindings.HasFuncDataTransferItemGetAsFile(
		this.ref,
	)
}

// FuncGetAsFile returns the method "DataTransferItem.getAsFile".
func (this DataTransferItem) FuncGetAsFile() (fn js.Func[func() File]) {
	bindings.FuncDataTransferItemGetAsFile(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAsFile calls the method "DataTransferItem.getAsFile".
func (this DataTransferItem) GetAsFile() (ret File) {
	bindings.CallDataTransferItemGetAsFile(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAsFile calls the method "DataTransferItem.getAsFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsFile() (ret File, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsFile(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAsFileSystemHandle returns true if the method "DataTransferItem.getAsFileSystemHandle" exists.
func (this DataTransferItem) HasFuncGetAsFileSystemHandle() bool {
	return js.True == bindings.HasFuncDataTransferItemGetAsFileSystemHandle(
		this.ref,
	)
}

// FuncGetAsFileSystemHandle returns the method "DataTransferItem.getAsFileSystemHandle".
func (this DataTransferItem) FuncGetAsFileSystemHandle() (fn js.Func[func() js.Promise[FileSystemHandle]]) {
	bindings.FuncDataTransferItemGetAsFileSystemHandle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAsFileSystemHandle calls the method "DataTransferItem.getAsFileSystemHandle".
func (this DataTransferItem) GetAsFileSystemHandle() (ret js.Promise[FileSystemHandle]) {
	bindings.CallDataTransferItemGetAsFileSystemHandle(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAsFileSystemHandle calls the method "DataTransferItem.getAsFileSystemHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryGetAsFileSystemHandle() (ret js.Promise[FileSystemHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemGetAsFileSystemHandle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWebkitGetAsEntry returns true if the method "DataTransferItem.webkitGetAsEntry" exists.
func (this DataTransferItem) HasFuncWebkitGetAsEntry() bool {
	return js.True == bindings.HasFuncDataTransferItemWebkitGetAsEntry(
		this.ref,
	)
}

// FuncWebkitGetAsEntry returns the method "DataTransferItem.webkitGetAsEntry".
func (this DataTransferItem) FuncWebkitGetAsEntry() (fn js.Func[func() FileSystemEntry]) {
	bindings.FuncDataTransferItemWebkitGetAsEntry(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WebkitGetAsEntry calls the method "DataTransferItem.webkitGetAsEntry".
func (this DataTransferItem) WebkitGetAsEntry() (ret FileSystemEntry) {
	bindings.CallDataTransferItemWebkitGetAsEntry(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryWebkitGetAsEntry calls the method "DataTransferItem.webkitGetAsEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItem) TryWebkitGetAsEntry() (ret FileSystemEntry, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemWebkitGetAsEntry(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DataTransferItemList struct {
	ref js.Ref
}

func (this DataTransferItemList) Once() DataTransferItemList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "DataTransferItemList.length".
//
// It returns ok=false if there is no such property.
func (this DataTransferItemList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDataTransferItemListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "DataTransferItemList." exists.
func (this DataTransferItemList) HasFuncGet() bool {
	return js.True == bindings.HasFuncDataTransferItemListGet(
		this.ref,
	)
}

// FuncGet returns the method "DataTransferItemList.".
func (this DataTransferItemList) FuncGet() (fn js.Func[func(index uint32) DataTransferItem]) {
	bindings.FuncDataTransferItemListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "DataTransferItemList.".
func (this DataTransferItemList) Get(index uint32) (ret DataTransferItem) {
	bindings.CallDataTransferItemListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "DataTransferItemList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryGet(index uint32) (ret DataTransferItem, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAdd returns true if the method "DataTransferItemList.add" exists.
func (this DataTransferItemList) HasFuncAdd() bool {
	return js.True == bindings.HasFuncDataTransferItemListAdd(
		this.ref,
	)
}

// FuncAdd returns the method "DataTransferItemList.add".
func (this DataTransferItemList) FuncAdd() (fn js.Func[func(data js.String, typ js.String) DataTransferItem]) {
	bindings.FuncDataTransferItemListAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "DataTransferItemList.add".
func (this DataTransferItemList) Add(data js.String, typ js.String) (ret DataTransferItem) {
	bindings.CallDataTransferItemListAdd(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		typ.Ref(),
	)

	return
}

// HasFuncAdd1 returns true if the method "DataTransferItemList.add" exists.
func (this DataTransferItemList) HasFuncAdd1() bool {
	return js.True == bindings.HasFuncDataTransferItemListAdd1(
		this.ref,
	)
}

// FuncAdd1 returns the method "DataTransferItemList.add".
func (this DataTransferItemList) FuncAdd1() (fn js.Func[func(data File) DataTransferItem]) {
	bindings.FuncDataTransferItemListAdd1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add1 calls the method "DataTransferItemList.add".
func (this DataTransferItemList) Add1(data File) (ret DataTransferItem) {
	bindings.CallDataTransferItemListAdd1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAdd1 calls the method "DataTransferItemList.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryAdd1(data File) (ret DataTransferItem, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListAdd1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncRemove returns true if the method "DataTransferItemList.remove" exists.
func (this DataTransferItemList) HasFuncRemove() bool {
	return js.True == bindings.HasFuncDataTransferItemListRemove(
		this.ref,
	)
}

// FuncRemove returns the method "DataTransferItemList.remove".
func (this DataTransferItemList) FuncRemove() (fn js.Func[func(index uint32)]) {
	bindings.FuncDataTransferItemListRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "DataTransferItemList.remove".
func (this DataTransferItemList) Remove(index uint32) (ret js.Void) {
	bindings.CallDataTransferItemListRemove(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemove calls the method "DataTransferItemList.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryRemove(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncClear returns true if the method "DataTransferItemList.clear" exists.
func (this DataTransferItemList) HasFuncClear() bool {
	return js.True == bindings.HasFuncDataTransferItemListClear(
		this.ref,
	)
}

// FuncClear returns the method "DataTransferItemList.clear".
func (this DataTransferItemList) FuncClear() (fn js.Func[func()]) {
	bindings.FuncDataTransferItemListClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "DataTransferItemList.clear".
func (this DataTransferItemList) Clear() (ret js.Void) {
	bindings.CallDataTransferItemListClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "DataTransferItemList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransferItemList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferItemListClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FileList struct {
	ref js.Ref
}

func (this FileList) Once() FileList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "FileList.length".
//
// It returns ok=false if there is no such property.
func (this FileList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFileListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "FileList.item" exists.
func (this FileList) HasFuncItem() bool {
	return js.True == bindings.HasFuncFileListItem(
		this.ref,
	)
}

// FuncItem returns the method "FileList.item".
func (this FileList) FuncItem() (fn js.Func[func(index uint32) File]) {
	bindings.FuncFileListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "FileList.item".
func (this FileList) Item(index uint32) (ret File) {
	bindings.CallFileListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "FileList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileList) TryItem(index uint32) (ret File, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type DataTransfer struct {
	ref js.Ref
}

func (this DataTransfer) Once() DataTransfer {
	this.ref.Once()
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
	this.ref.Free()
}

// DropEffect returns the value of property "DataTransfer.dropEffect".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) DropEffect() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferDropEffect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDropEffect sets the value of property "DataTransfer.dropEffect" to val.
//
// It returns false if the property cannot be set.
func (this DataTransfer) SetDropEffect(val js.String) bool {
	return js.True == bindings.SetDataTransferDropEffect(
		this.ref,
		val.Ref(),
	)
}

// EffectAllowed returns the value of property "DataTransfer.effectAllowed".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) EffectAllowed() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataTransferEffectAllowed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEffectAllowed sets the value of property "DataTransfer.effectAllowed" to val.
//
// It returns false if the property cannot be set.
func (this DataTransfer) SetEffectAllowed(val js.String) bool {
	return js.True == bindings.SetDataTransferEffectAllowed(
		this.ref,
		val.Ref(),
	)
}

// Items returns the value of property "DataTransfer.items".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) Items() (ret DataTransferItemList, ok bool) {
	ok = js.True == bindings.GetDataTransferItems(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Types returns the value of property "DataTransfer.types".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) Types() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetDataTransferTypes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "DataTransfer.files".
//
// It returns ok=false if there is no such property.
func (this DataTransfer) Files() (ret FileList, ok bool) {
	ok = js.True == bindings.GetDataTransferFiles(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetDragImage returns true if the method "DataTransfer.setDragImage" exists.
func (this DataTransfer) HasFuncSetDragImage() bool {
	return js.True == bindings.HasFuncDataTransferSetDragImage(
		this.ref,
	)
}

// FuncSetDragImage returns the method "DataTransfer.setDragImage".
func (this DataTransfer) FuncSetDragImage() (fn js.Func[func(image Element, x int32, y int32)]) {
	bindings.FuncDataTransferSetDragImage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetDragImage calls the method "DataTransfer.setDragImage".
func (this DataTransfer) SetDragImage(image Element, x int32, y int32) (ret js.Void) {
	bindings.CallDataTransferSetDragImage(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncGetData returns true if the method "DataTransfer.getData" exists.
func (this DataTransfer) HasFuncGetData() bool {
	return js.True == bindings.HasFuncDataTransferGetData(
		this.ref,
	)
}

// FuncGetData returns the method "DataTransfer.getData".
func (this DataTransfer) FuncGetData() (fn js.Func[func(format js.String) js.String]) {
	bindings.FuncDataTransferGetData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetData calls the method "DataTransfer.getData".
func (this DataTransfer) GetData(format js.String) (ret js.String) {
	bindings.CallDataTransferGetData(
		this.ref, js.Pointer(&ret),
		format.Ref(),
	)

	return
}

// TryGetData calls the method "DataTransfer.getData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryGetData(format js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferGetData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
	)

	return
}

// HasFuncSetData returns true if the method "DataTransfer.setData" exists.
func (this DataTransfer) HasFuncSetData() bool {
	return js.True == bindings.HasFuncDataTransferSetData(
		this.ref,
	)
}

// FuncSetData returns the method "DataTransfer.setData".
func (this DataTransfer) FuncSetData() (fn js.Func[func(format js.String, data js.String)]) {
	bindings.FuncDataTransferSetData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetData calls the method "DataTransfer.setData".
func (this DataTransfer) SetData(format js.String, data js.String) (ret js.Void) {
	bindings.CallDataTransferSetData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncClearData returns true if the method "DataTransfer.clearData" exists.
func (this DataTransfer) HasFuncClearData() bool {
	return js.True == bindings.HasFuncDataTransferClearData(
		this.ref,
	)
}

// FuncClearData returns the method "DataTransfer.clearData".
func (this DataTransfer) FuncClearData() (fn js.Func[func(format js.String)]) {
	bindings.FuncDataTransferClearData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearData calls the method "DataTransfer.clearData".
func (this DataTransfer) ClearData(format js.String) (ret js.Void) {
	bindings.CallDataTransferClearData(
		this.ref, js.Pointer(&ret),
		format.Ref(),
	)

	return
}

// TryClearData calls the method "DataTransfer.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryClearData(format js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferClearData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		format.Ref(),
	)

	return
}

// HasFuncClearData1 returns true if the method "DataTransfer.clearData" exists.
func (this DataTransfer) HasFuncClearData1() bool {
	return js.True == bindings.HasFuncDataTransferClearData1(
		this.ref,
	)
}

// FuncClearData1 returns the method "DataTransfer.clearData".
func (this DataTransfer) FuncClearData1() (fn js.Func[func()]) {
	bindings.FuncDataTransferClearData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearData1 calls the method "DataTransfer.clearData".
func (this DataTransfer) ClearData1() (ret js.Void) {
	bindings.CallDataTransferClearData1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearData1 calls the method "DataTransfer.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DataTransfer) TryClearData1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDataTransferClearData1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ClipboardEventInit) UpdateFrom(ref js.Ref) {
	bindings.ClipboardEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClipboardEventInit) Update(ref js.Ref) {
	bindings.ClipboardEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClipboardEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.ClipboardData.Ref(),
	)
	p.ClipboardData = p.ClipboardData.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// ClipboardData returns the value of property "ClipboardEvent.clipboardData".
//
// It returns ok=false if there is no such property.
func (this ClipboardEvent) ClipboardData() (ret DataTransfer, ok bool) {
	ok = js.True == bindings.GetClipboardEventClipboardData(
		this.ref, js.Pointer(&ret),
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
func (p *ClipboardPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.ClipboardPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClipboardPermissionDescriptor) Update(ref js.Ref) {
	bindings.ClipboardPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClipboardPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *CloseEventInit) UpdateFrom(ref js.Ref) {
	bindings.CloseEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CloseEventInit) Update(ref js.Ref) {
	bindings.CloseEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CloseEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Reason.Ref(),
	)
	p.Reason = p.Reason.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// WasClean returns the value of property "CloseEvent.wasClean".
//
// It returns ok=false if there is no such property.
func (this CloseEvent) WasClean() (ret bool, ok bool) {
	ok = js.True == bindings.GetCloseEventWasClean(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Code returns the value of property "CloseEvent.code".
//
// It returns ok=false if there is no such property.
func (this CloseEvent) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetCloseEventCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Reason returns the value of property "CloseEvent.reason".
//
// It returns ok=false if there is no such property.
func (this CloseEvent) Reason() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCloseEventReason(
		this.ref, js.Pointer(&ret),
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
func (p *CollectedClientAdditionalPaymentData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientAdditionalPaymentDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CollectedClientAdditionalPaymentData) Update(ref js.Ref) {
	bindings.CollectedClientAdditionalPaymentDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CollectedClientAdditionalPaymentData) FreeMembers(recursive bool) {
	js.Free(
		p.RpId.Ref(),
		p.TopOrigin.Ref(),
		p.PayeeName.Ref(),
		p.PayeeOrigin.Ref(),
	)
	p.RpId = p.RpId.FromRef(js.Undefined)
	p.TopOrigin = p.TopOrigin.FromRef(js.Undefined)
	p.PayeeName = p.PayeeName.FromRef(js.Undefined)
	p.PayeeOrigin = p.PayeeOrigin.FromRef(js.Undefined)
	if recursive {
		p.Total.FreeMembers(true)
		p.Instrument.FreeMembers(true)
	}
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
func (p *CollectedClientData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CollectedClientData) Update(ref js.Ref) {
	bindings.CollectedClientDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CollectedClientData) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Challenge.Ref(),
		p.Origin.Ref(),
		p.TopOrigin.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.TopOrigin = p.TopOrigin.FromRef(js.Undefined)
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
func (p *CollectedClientPaymentData) UpdateFrom(ref js.Ref) {
	bindings.CollectedClientPaymentDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CollectedClientPaymentData) Update(ref js.Ref) {
	bindings.CollectedClientPaymentDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CollectedClientPaymentData) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
		p.Challenge.Ref(),
		p.Origin.Ref(),
		p.TopOrigin.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.TopOrigin = p.TopOrigin.FromRef(js.Undefined)
	if recursive {
		p.Payment.FreeMembers(true)
	}
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
func (p *ColorSelectionOptions) UpdateFrom(ref js.Ref) {
	bindings.ColorSelectionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ColorSelectionOptions) Update(ref js.Ref) {
	bindings.ColorSelectionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ColorSelectionOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
func (p *ColorSelectionResult) UpdateFrom(ref js.Ref) {
	bindings.ColorSelectionResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ColorSelectionResult) Update(ref js.Ref) {
	bindings.ColorSelectionResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ColorSelectionResult) FreeMembers(recursive bool) {
	js.Free(
		p.SRGBHex.Ref(),
	)
	p.SRGBHex = p.SRGBHex.FromRef(js.Undefined)
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
func (p *WindowPostMessageOptions) UpdateFrom(ref js.Ref) {
	bindings.WindowPostMessageOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WindowPostMessageOptions) Update(ref js.Ref) {
	bindings.WindowPostMessageOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WindowPostMessageOptions) FreeMembers(recursive bool) {
	js.Free(
		p.TargetOrigin.Ref(),
		p.Transfer.Ref(),
	)
	p.TargetOrigin = p.TargetOrigin.FromRef(js.Undefined)
	p.Transfer = p.Transfer.FromRef(js.Undefined)
}

type MediaQueryList struct {
	EventTarget
}

func (this MediaQueryList) Once() MediaQueryList {
	this.ref.Once()
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
	this.ref.Free()
}

// Media returns the value of property "MediaQueryList.media".
//
// It returns ok=false if there is no such property.
func (this MediaQueryList) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaQueryListMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matches returns the value of property "MediaQueryList.matches".
//
// It returns ok=false if there is no such property.
func (this MediaQueryList) Matches() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaQueryListMatches(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAddListener returns true if the method "MediaQueryList.addListener" exists.
func (this MediaQueryList) HasFuncAddListener() bool {
	return js.True == bindings.HasFuncMediaQueryListAddListener(
		this.ref,
	)
}

// FuncAddListener returns the method "MediaQueryList.addListener".
func (this MediaQueryList) FuncAddListener() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	bindings.FuncMediaQueryListAddListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddListener calls the method "MediaQueryList.addListener".
func (this MediaQueryList) AddListener(callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallMediaQueryListAddListener(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryAddListener calls the method "MediaQueryList.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaQueryList) TryAddListener(callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaQueryListAddListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveListener returns true if the method "MediaQueryList.removeListener" exists.
func (this MediaQueryList) HasFuncRemoveListener() bool {
	return js.True == bindings.HasFuncMediaQueryListRemoveListener(
		this.ref,
	)
}

// FuncRemoveListener returns the method "MediaQueryList.removeListener".
func (this MediaQueryList) FuncRemoveListener() (fn js.Func[func(callback js.Func[func(event Event)])]) {
	bindings.FuncMediaQueryListRemoveListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveListener calls the method "MediaQueryList.removeListener".
func (this MediaQueryList) RemoveListener(callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallMediaQueryListRemoveListener(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRemoveListener calls the method "MediaQueryList.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaQueryList) TryRemoveListener(callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaQueryListRemoveListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

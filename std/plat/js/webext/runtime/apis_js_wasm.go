// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

type ConnectArgConnectInfo struct {
	// IncludeTlsChannelId is "ConnectArgConnectInfo.includeTlsChannelId"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncludeTlsChannelId MUST be set to true to make this field effective.
	IncludeTlsChannelId bool
	// Name is "ConnectArgConnectInfo.name"
	//
	// Optional
	Name js.String

	FFI_USE_IncludeTlsChannelId bool // for IncludeTlsChannelId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectArgConnectInfo with all fields set.
func (p ConnectArgConnectInfo) FromRef(ref js.Ref) ConnectArgConnectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectArgConnectInfo in the application heap.
func (p ConnectArgConnectInfo) New() js.Ref {
	return bindings.ConnectArgConnectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectArgConnectInfo) UpdateFrom(ref js.Ref) {
	bindings.ConnectArgConnectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectArgConnectInfo) Update(ref js.Ref) {
	bindings.ConnectArgConnectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectArgConnectInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type ContextType uint32

const (
	_ ContextType = iota

	ContextType_TAB
	ContextType_POPUP
	ContextType_BACKGROUND
	ContextType_OFFSCREEN_DOCUMENT
	ContextType_SIDE_PANEL
)

func (ContextType) FromRef(str js.Ref) ContextType {
	return ContextType(bindings.ConstOfContextType(str))
}

func (x ContextType) String() (string, bool) {
	switch x {
	case ContextType_TAB:
		return "TAB", true
	case ContextType_POPUP:
		return "POPUP", true
	case ContextType_BACKGROUND:
		return "BACKGROUND", true
	case ContextType_OFFSCREEN_DOCUMENT:
		return "OFFSCREEN_DOCUMENT", true
	case ContextType_SIDE_PANEL:
		return "SIDE_PANEL", true
	default:
		return "", false
	}
}

type ContextFilter struct {
	// ContextIds is "ContextFilter.contextIds"
	//
	// Optional
	ContextIds js.Array[js.String]
	// ContextTypes is "ContextFilter.contextTypes"
	//
	// Optional
	ContextTypes js.Array[ContextType]
	// DocumentIds is "ContextFilter.documentIds"
	//
	// Optional
	DocumentIds js.Array[js.String]
	// DocumentOrigins is "ContextFilter.documentOrigins"
	//
	// Optional
	DocumentOrigins js.Array[js.String]
	// DocumentUrls is "ContextFilter.documentUrls"
	//
	// Optional
	DocumentUrls js.Array[js.String]
	// FrameIds is "ContextFilter.frameIds"
	//
	// Optional
	FrameIds js.Array[int64]
	// Incognito is "ContextFilter.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// TabIds is "ContextFilter.tabIds"
	//
	// Optional
	TabIds js.Array[int64]
	// WindowIds is "ContextFilter.windowIds"
	//
	// Optional
	WindowIds js.Array[int64]

	FFI_USE_Incognito bool // for Incognito.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextFilter with all fields set.
func (p ContextFilter) FromRef(ref js.Ref) ContextFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextFilter in the application heap.
func (p ContextFilter) New() js.Ref {
	return bindings.ContextFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextFilter) UpdateFrom(ref js.Ref) {
	bindings.ContextFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextFilter) Update(ref js.Ref) {
	bindings.ContextFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextFilter) FreeMembers(recursive bool) {
	js.Free(
		p.ContextIds.Ref(),
		p.ContextTypes.Ref(),
		p.DocumentIds.Ref(),
		p.DocumentOrigins.Ref(),
		p.DocumentUrls.Ref(),
		p.FrameIds.Ref(),
		p.TabIds.Ref(),
		p.WindowIds.Ref(),
	)
	p.ContextIds = p.ContextIds.FromRef(js.Undefined)
	p.ContextTypes = p.ContextTypes.FromRef(js.Undefined)
	p.DocumentIds = p.DocumentIds.FromRef(js.Undefined)
	p.DocumentOrigins = p.DocumentOrigins.FromRef(js.Undefined)
	p.DocumentUrls = p.DocumentUrls.FromRef(js.Undefined)
	p.FrameIds = p.FrameIds.FromRef(js.Undefined)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
	p.WindowIds = p.WindowIds.FromRef(js.Undefined)
}

type ExtensionContext struct {
	// ContextId is "ExtensionContext.contextId"
	//
	// Required
	ContextId js.String
	// ContextType is "ExtensionContext.contextType"
	//
	// Required
	ContextType ContextType
	// DocumentId is "ExtensionContext.documentId"
	//
	// Optional
	DocumentId js.String
	// DocumentOrigin is "ExtensionContext.documentOrigin"
	//
	// Optional
	DocumentOrigin js.String
	// DocumentUrl is "ExtensionContext.documentUrl"
	//
	// Optional
	DocumentUrl js.String
	// FrameId is "ExtensionContext.frameId"
	//
	// Required
	FrameId int64
	// Incognito is "ExtensionContext.incognito"
	//
	// Required
	Incognito bool
	// TabId is "ExtensionContext.tabId"
	//
	// Required
	TabId int64
	// WindowId is "ExtensionContext.windowId"
	//
	// Required
	WindowId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionContext with all fields set.
func (p ExtensionContext) FromRef(ref js.Ref) ExtensionContext {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionContext in the application heap.
func (p ExtensionContext) New() js.Ref {
	return bindings.ExtensionContextJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionContext) UpdateFrom(ref js.Ref) {
	bindings.ExtensionContextJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionContext) Update(ref js.Ref) {
	bindings.ExtensionContextJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionContext) FreeMembers(recursive bool) {
	js.Free(
		p.ContextId.Ref(),
		p.DocumentId.Ref(),
		p.DocumentOrigin.Ref(),
		p.DocumentUrl.Ref(),
	)
	p.ContextId = p.ContextId.FromRef(js.Undefined)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.DocumentOrigin = p.DocumentOrigin.FromRef(js.Undefined)
	p.DocumentUrl = p.DocumentUrl.FromRef(js.Undefined)
}

type GetPackageDirectoryEntryArgCallbackFunc func(this js.Ref, directoryEntry js.Any) js.Ref

func (fn GetPackageDirectoryEntryArgCallbackFunc) Register() js.Func[func(directoryEntry js.Any)] {
	return js.RegisterCallback[func(directoryEntry js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPackageDirectoryEntryArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPackageDirectoryEntryArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, directoryEntry js.Any) js.Ref
	Arg T
}

func (cb *GetPackageDirectoryEntryArgCallback[T]) Register() js.Func[func(directoryEntry js.Any)] {
	return js.RegisterCallback[func(directoryEntry js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPackageDirectoryEntryArgCallback[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LastErrorProperty struct {
	// Message is "LastErrorProperty.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LastErrorProperty with all fields set.
func (p LastErrorProperty) FromRef(ref js.Ref) LastErrorProperty {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LastErrorProperty in the application heap.
func (p LastErrorProperty) New() js.Ref {
	return bindings.LastErrorPropertyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LastErrorProperty) UpdateFrom(ref js.Ref) {
	bindings.LastErrorPropertyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LastErrorProperty) Update(ref js.Ref) {
	bindings.LastErrorPropertyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LastErrorProperty) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type MessageSender struct {
	// DocumentId is "MessageSender.documentId"
	//
	// Optional
	DocumentId js.String
	// DocumentLifecycle is "MessageSender.documentLifecycle"
	//
	// Optional
	DocumentLifecycle js.String
	// FrameId is "MessageSender.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int64
	// GuestProcessId is "MessageSender.guestProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GuestProcessId MUST be set to true to make this field effective.
	GuestProcessId int64
	// GuestRenderFrameRoutingId is "MessageSender.guestRenderFrameRoutingId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GuestRenderFrameRoutingId MUST be set to true to make this field effective.
	GuestRenderFrameRoutingId int64
	// Id is "MessageSender.id"
	//
	// Optional
	Id js.String
	// NativeApplication is "MessageSender.nativeApplication"
	//
	// Optional
	NativeApplication js.String
	// Origin is "MessageSender.origin"
	//
	// Optional
	Origin js.String
	// Tab is "MessageSender.tab"
	//
	// Optional
	//
	// NOTE: Tab.FFI_USE MUST be set to true to get Tab used.
	Tab tabs.Tab
	// TlsChannelId is "MessageSender.tlsChannelId"
	//
	// Optional
	TlsChannelId js.String
	// Url is "MessageSender.url"
	//
	// Optional
	Url js.String

	FFI_USE_FrameId                   bool // for FrameId.
	FFI_USE_GuestProcessId            bool // for GuestProcessId.
	FFI_USE_GuestRenderFrameRoutingId bool // for GuestRenderFrameRoutingId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MessageSender with all fields set.
func (p MessageSender) FromRef(ref js.Ref) MessageSender {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MessageSender in the application heap.
func (p MessageSender) New() js.Ref {
	return bindings.MessageSenderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MessageSender) UpdateFrom(ref js.Ref) {
	bindings.MessageSenderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MessageSender) Update(ref js.Ref) {
	bindings.MessageSenderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MessageSender) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.DocumentLifecycle.Ref(),
		p.Id.Ref(),
		p.NativeApplication.Ref(),
		p.Origin.Ref(),
		p.TlsChannelId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.DocumentLifecycle = p.DocumentLifecycle.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.NativeApplication = p.NativeApplication.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.TlsChannelId = p.TlsChannelId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	if recursive {
		p.Tab.FreeMembers(true)
	}
}

type OnInstalledReason uint32

const (
	_ OnInstalledReason = iota

	OnInstalledReason_INSTALL
	OnInstalledReason_UPDATE
	OnInstalledReason_CHROME_UPDATE
	OnInstalledReason_SHARED_MODULE_UPDATE
)

func (OnInstalledReason) FromRef(str js.Ref) OnInstalledReason {
	return OnInstalledReason(bindings.ConstOfOnInstalledReason(str))
}

func (x OnInstalledReason) String() (string, bool) {
	switch x {
	case OnInstalledReason_INSTALL:
		return "install", true
	case OnInstalledReason_UPDATE:
		return "update", true
	case OnInstalledReason_CHROME_UPDATE:
		return "chrome_update", true
	case OnInstalledReason_SHARED_MODULE_UPDATE:
		return "shared_module_update", true
	default:
		return "", false
	}
}

type OnInstalledArgDetails struct {
	// Id is "OnInstalledArgDetails.id"
	//
	// Optional
	Id js.String
	// PreviousVersion is "OnInstalledArgDetails.previousVersion"
	//
	// Optional
	PreviousVersion js.String
	// Reason is "OnInstalledArgDetails.reason"
	//
	// Required
	Reason OnInstalledReason

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnInstalledArgDetails with all fields set.
func (p OnInstalledArgDetails) FromRef(ref js.Ref) OnInstalledArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnInstalledArgDetails in the application heap.
func (p OnInstalledArgDetails) New() js.Ref {
	return bindings.OnInstalledArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnInstalledArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnInstalledArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnInstalledArgDetails) Update(ref js.Ref) {
	bindings.OnInstalledArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnInstalledArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.PreviousVersion.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.PreviousVersion = p.PreviousVersion.FromRef(js.Undefined)
}

type OnMessageArgSendResponseFunc func(this js.Ref) js.Ref

func (fn OnMessageArgSendResponseFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageArgSendResponseFunc) DispatchCallback(
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

type OnMessageArgSendResponse[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnMessageArgSendResponse[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageArgSendResponse[T]) DispatchCallback(
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

type OnMessageExternalArgSendResponseFunc func(this js.Ref) js.Ref

func (fn OnMessageExternalArgSendResponseFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageExternalArgSendResponseFunc) DispatchCallback(
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

type OnMessageExternalArgSendResponse[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnMessageExternalArgSendResponse[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageExternalArgSendResponse[T]) DispatchCallback(
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

type OnRestartRequiredReason uint32

const (
	_ OnRestartRequiredReason = iota

	OnRestartRequiredReason_APP_UPDATE
	OnRestartRequiredReason_OS_UPDATE
	OnRestartRequiredReason_PERIODIC
)

func (OnRestartRequiredReason) FromRef(str js.Ref) OnRestartRequiredReason {
	return OnRestartRequiredReason(bindings.ConstOfOnRestartRequiredReason(str))
}

func (x OnRestartRequiredReason) String() (string, bool) {
	switch x {
	case OnRestartRequiredReason_APP_UPDATE:
		return "app_update", true
	case OnRestartRequiredReason_OS_UPDATE:
		return "os_update", true
	case OnRestartRequiredReason_PERIODIC:
		return "periodic", true
	default:
		return "", false
	}
}

type OnUserScriptMessageArgSendResponseFunc func(this js.Ref) js.Ref

func (fn OnUserScriptMessageArgSendResponseFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUserScriptMessageArgSendResponseFunc) DispatchCallback(
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

type OnUserScriptMessageArgSendResponse[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnUserScriptMessageArgSendResponse[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUserScriptMessageArgSendResponse[T]) DispatchCallback(
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

type PlatformArch uint32

const (
	_ PlatformArch = iota

	PlatformArch_ARM
	PlatformArch_ARM64
	PlatformArch_X86_32
	PlatformArch_X86_64
	PlatformArch_MIPS
	PlatformArch_MIPS64
)

func (PlatformArch) FromRef(str js.Ref) PlatformArch {
	return PlatformArch(bindings.ConstOfPlatformArch(str))
}

func (x PlatformArch) String() (string, bool) {
	switch x {
	case PlatformArch_ARM:
		return "arm", true
	case PlatformArch_ARM64:
		return "arm64", true
	case PlatformArch_X86_32:
		return "x86-32", true
	case PlatformArch_X86_64:
		return "x86-64", true
	case PlatformArch_MIPS:
		return "mips", true
	case PlatformArch_MIPS64:
		return "mips64", true
	default:
		return "", false
	}
}

type PlatformNaclArch uint32

const (
	_ PlatformNaclArch = iota

	PlatformNaclArch_ARM
	PlatformNaclArch_X86_32
	PlatformNaclArch_X86_64
	PlatformNaclArch_MIPS
	PlatformNaclArch_MIPS64
)

func (PlatformNaclArch) FromRef(str js.Ref) PlatformNaclArch {
	return PlatformNaclArch(bindings.ConstOfPlatformNaclArch(str))
}

func (x PlatformNaclArch) String() (string, bool) {
	switch x {
	case PlatformNaclArch_ARM:
		return "arm", true
	case PlatformNaclArch_X86_32:
		return "x86-32", true
	case PlatformNaclArch_X86_64:
		return "x86-64", true
	case PlatformNaclArch_MIPS:
		return "mips", true
	case PlatformNaclArch_MIPS64:
		return "mips64", true
	default:
		return "", false
	}
}

type PlatformOs uint32

const (
	_ PlatformOs = iota

	PlatformOs_MAC
	PlatformOs_WIN
	PlatformOs_ANDROID
	PlatformOs_CROS
	PlatformOs_LINUX
	PlatformOs_OPENBSD
	PlatformOs_FUCHSIA
)

func (PlatformOs) FromRef(str js.Ref) PlatformOs {
	return PlatformOs(bindings.ConstOfPlatformOs(str))
}

func (x PlatformOs) String() (string, bool) {
	switch x {
	case PlatformOs_MAC:
		return "mac", true
	case PlatformOs_WIN:
		return "win", true
	case PlatformOs_ANDROID:
		return "android", true
	case PlatformOs_CROS:
		return "cros", true
	case PlatformOs_LINUX:
		return "linux", true
	case PlatformOs_OPENBSD:
		return "openbsd", true
	case PlatformOs_FUCHSIA:
		return "fuchsia", true
	default:
		return "", false
	}
}

type PlatformInfo struct {
	// Arch is "PlatformInfo.arch"
	//
	// Required
	Arch PlatformArch
	// NaclArch is "PlatformInfo.nacl_arch"
	//
	// Required
	NaclArch PlatformNaclArch
	// Os is "PlatformInfo.os"
	//
	// Required
	Os PlatformOs

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PlatformInfo with all fields set.
func (p PlatformInfo) FromRef(ref js.Ref) PlatformInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PlatformInfo in the application heap.
func (p PlatformInfo) New() js.Ref {
	return bindings.PlatformInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PlatformInfo) UpdateFrom(ref js.Ref) {
	bindings.PlatformInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PlatformInfo) Update(ref js.Ref) {
	bindings.PlatformInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PlatformInfo) FreeMembers(recursive bool) {
}

type PortFieldDisconnectFunc func(this js.Ref) js.Ref

func (fn PortFieldDisconnectFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PortFieldDisconnectFunc) DispatchCallback(
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

type PortFieldDisconnect[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *PortFieldDisconnect[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PortFieldDisconnect[T]) DispatchCallback(
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

type PortFieldPostMessageFunc func(this js.Ref, message js.Any) js.Ref

func (fn PortFieldPostMessageFunc) Register() js.Func[func(message js.Any)] {
	return js.RegisterCallback[func(message js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PortFieldPostMessageFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PortFieldPostMessage[T any] struct {
	Fn  func(arg T, this js.Ref, message js.Any) js.Ref
	Arg T
}

func (cb *PortFieldPostMessage[T]) Register() js.Func[func(message js.Any)] {
	return js.RegisterCallback[func(message js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PortFieldPostMessage[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Port struct {
	// Disconnect is "Port.disconnect"
	//
	// Required
	Disconnect js.Func[func()]
	// Name is "Port.name"
	//
	// Required
	Name js.String
	// PostMessage is "Port.postMessage"
	//
	// Required
	PostMessage js.Func[func(message js.Any)]
	// Sender is "Port.sender"
	//
	// Optional
	//
	// NOTE: Sender.FFI_USE MUST be set to true to get Sender used.
	Sender MessageSender

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Port with all fields set.
func (p Port) FromRef(ref js.Ref) Port {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Port in the application heap.
func (p Port) New() js.Ref {
	return bindings.PortJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Port) UpdateFrom(ref js.Ref) {
	bindings.PortJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Port) Update(ref js.Ref) {
	bindings.PortJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Port) FreeMembers(recursive bool) {
	js.Free(
		p.Disconnect.Ref(),
		p.Name.Ref(),
		p.PostMessage.Ref(),
	)
	p.Disconnect = p.Disconnect.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PostMessage = p.PostMessage.FromRef(js.Undefined)
	if recursive {
		p.Sender.FreeMembers(true)
	}
}

type RequestUpdateCheckStatus uint32

const (
	_ RequestUpdateCheckStatus = iota

	RequestUpdateCheckStatus_THROTTLED
	RequestUpdateCheckStatus_NO_UPDATE
	RequestUpdateCheckStatus_UPDATE_AVAILABLE
)

func (RequestUpdateCheckStatus) FromRef(str js.Ref) RequestUpdateCheckStatus {
	return RequestUpdateCheckStatus(bindings.ConstOfRequestUpdateCheckStatus(str))
}

func (x RequestUpdateCheckStatus) String() (string, bool) {
	switch x {
	case RequestUpdateCheckStatus_THROTTLED:
		return "throttled", true
	case RequestUpdateCheckStatus_NO_UPDATE:
		return "no_update", true
	case RequestUpdateCheckStatus_UPDATE_AVAILABLE:
		return "update_available", true
	default:
		return "", false
	}
}

type RequestUpdateCheckReturnType struct {
	// Status is "RequestUpdateCheckReturnType.status"
	//
	// Required
	Status RequestUpdateCheckStatus
	// Version is "RequestUpdateCheckReturnType.version"
	//
	// Optional
	Version js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestUpdateCheckReturnType with all fields set.
func (p RequestUpdateCheckReturnType) FromRef(ref js.Ref) RequestUpdateCheckReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestUpdateCheckReturnType in the application heap.
func (p RequestUpdateCheckReturnType) New() js.Ref {
	return bindings.RequestUpdateCheckReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestUpdateCheckReturnType) UpdateFrom(ref js.Ref) {
	bindings.RequestUpdateCheckReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestUpdateCheckReturnType) Update(ref js.Ref) {
	bindings.RequestUpdateCheckReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestUpdateCheckReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Version.Ref(),
	)
	p.Version = p.Version.FromRef(js.Undefined)
}

type SendMessageArgOptions struct {
	// IncludeTlsChannelId is "SendMessageArgOptions.includeTlsChannelId"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncludeTlsChannelId MUST be set to true to make this field effective.
	IncludeTlsChannelId bool

	FFI_USE_IncludeTlsChannelId bool // for IncludeTlsChannelId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendMessageArgOptions with all fields set.
func (p SendMessageArgOptions) FromRef(ref js.Ref) SendMessageArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendMessageArgOptions in the application heap.
func (p SendMessageArgOptions) New() js.Ref {
	return bindings.SendMessageArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendMessageArgOptions) UpdateFrom(ref js.Ref) {
	bindings.SendMessageArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendMessageArgOptions) Update(ref js.Ref) {
	bindings.SendMessageArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendMessageArgOptions) FreeMembers(recursive bool) {
}

// HasFuncConnect returns true if the function "WEBEXT.runtime.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.runtime.connect".
func FuncConnect() (fn js.Func[func(extensionId js.String, connectInfo ConnectArgConnectInfo) Port]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.runtime.connect" directly.
func Connect(extensionId js.String, connectInfo ConnectArgConnectInfo) (ret Port) {
	bindings.CallConnect(
		js.Pointer(&ret),
		extensionId.Ref(),
		js.Pointer(&connectInfo),
	)

	return
}

// TryConnect calls the function "WEBEXT.runtime.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(extensionId js.String, connectInfo ConnectArgConnectInfo) (ret Port, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		js.Pointer(&connectInfo),
	)

	return
}

// HasFuncConnectNative returns true if the function "WEBEXT.runtime.connectNative" exists.
func HasFuncConnectNative() bool {
	return js.True == bindings.HasFuncConnectNative()
}

// FuncConnectNative returns the function "WEBEXT.runtime.connectNative".
func FuncConnectNative() (fn js.Func[func(application js.String) Port]) {
	bindings.FuncConnectNative(
		js.Pointer(&fn),
	)
	return
}

// ConnectNative calls the function "WEBEXT.runtime.connectNative" directly.
func ConnectNative(application js.String) (ret Port) {
	bindings.CallConnectNative(
		js.Pointer(&ret),
		application.Ref(),
	)

	return
}

// TryConnectNative calls the function "WEBEXT.runtime.connectNative"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnectNative(application js.String) (ret Port, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnectNative(
		js.Pointer(&ret), js.Pointer(&exception),
		application.Ref(),
	)

	return
}

// HasFuncGetBackgroundPage returns true if the function "WEBEXT.runtime.getBackgroundPage" exists.
func HasFuncGetBackgroundPage() bool {
	return js.True == bindings.HasFuncGetBackgroundPage()
}

// FuncGetBackgroundPage returns the function "WEBEXT.runtime.getBackgroundPage".
func FuncGetBackgroundPage() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncGetBackgroundPage(
		js.Pointer(&fn),
	)
	return
}

// GetBackgroundPage calls the function "WEBEXT.runtime.getBackgroundPage" directly.
func GetBackgroundPage() (ret js.Promise[js.Any]) {
	bindings.CallGetBackgroundPage(
		js.Pointer(&ret),
	)

	return
}

// TryGetBackgroundPage calls the function "WEBEXT.runtime.getBackgroundPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBackgroundPage() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBackgroundPage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetContexts returns true if the function "WEBEXT.runtime.getContexts" exists.
func HasFuncGetContexts() bool {
	return js.True == bindings.HasFuncGetContexts()
}

// FuncGetContexts returns the function "WEBEXT.runtime.getContexts".
func FuncGetContexts() (fn js.Func[func(filter ContextFilter) js.Promise[js.Array[ExtensionContext]]]) {
	bindings.FuncGetContexts(
		js.Pointer(&fn),
	)
	return
}

// GetContexts calls the function "WEBEXT.runtime.getContexts" directly.
func GetContexts(filter ContextFilter) (ret js.Promise[js.Array[ExtensionContext]]) {
	bindings.CallGetContexts(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetContexts calls the function "WEBEXT.runtime.getContexts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContexts(filter ContextFilter) (ret js.Promise[js.Array[ExtensionContext]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContexts(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncGetManifest returns true if the function "WEBEXT.runtime.getManifest" exists.
func HasFuncGetManifest() bool {
	return js.True == bindings.HasFuncGetManifest()
}

// FuncGetManifest returns the function "WEBEXT.runtime.getManifest".
func FuncGetManifest() (fn js.Func[func() js.Any]) {
	bindings.FuncGetManifest(
		js.Pointer(&fn),
	)
	return
}

// GetManifest calls the function "WEBEXT.runtime.getManifest" directly.
func GetManifest() (ret js.Any) {
	bindings.CallGetManifest(
		js.Pointer(&ret),
	)

	return
}

// TryGetManifest calls the function "WEBEXT.runtime.getManifest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetManifest() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetManifest(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPackageDirectoryEntry returns true if the function "WEBEXT.runtime.getPackageDirectoryEntry" exists.
func HasFuncGetPackageDirectoryEntry() bool {
	return js.True == bindings.HasFuncGetPackageDirectoryEntry()
}

// FuncGetPackageDirectoryEntry returns the function "WEBEXT.runtime.getPackageDirectoryEntry".
func FuncGetPackageDirectoryEntry() (fn js.Func[func(callback js.Func[func(directoryEntry js.Any)])]) {
	bindings.FuncGetPackageDirectoryEntry(
		js.Pointer(&fn),
	)
	return
}

// GetPackageDirectoryEntry calls the function "WEBEXT.runtime.getPackageDirectoryEntry" directly.
func GetPackageDirectoryEntry(callback js.Func[func(directoryEntry js.Any)]) (ret js.Void) {
	bindings.CallGetPackageDirectoryEntry(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetPackageDirectoryEntry calls the function "WEBEXT.runtime.getPackageDirectoryEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPackageDirectoryEntry(callback js.Func[func(directoryEntry js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPackageDirectoryEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetPlatformInfo returns true if the function "WEBEXT.runtime.getPlatformInfo" exists.
func HasFuncGetPlatformInfo() bool {
	return js.True == bindings.HasFuncGetPlatformInfo()
}

// FuncGetPlatformInfo returns the function "WEBEXT.runtime.getPlatformInfo".
func FuncGetPlatformInfo() (fn js.Func[func() js.Promise[PlatformInfo]]) {
	bindings.FuncGetPlatformInfo(
		js.Pointer(&fn),
	)
	return
}

// GetPlatformInfo calls the function "WEBEXT.runtime.getPlatformInfo" directly.
func GetPlatformInfo() (ret js.Promise[PlatformInfo]) {
	bindings.CallGetPlatformInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetPlatformInfo calls the function "WEBEXT.runtime.getPlatformInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPlatformInfo() (ret js.Promise[PlatformInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPlatformInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetURL returns true if the function "WEBEXT.runtime.getURL" exists.
func HasFuncGetURL() bool {
	return js.True == bindings.HasFuncGetURL()
}

// FuncGetURL returns the function "WEBEXT.runtime.getURL".
func FuncGetURL() (fn js.Func[func(path js.String) js.String]) {
	bindings.FuncGetURL(
		js.Pointer(&fn),
	)
	return
}

// GetURL calls the function "WEBEXT.runtime.getURL" directly.
func GetURL(path js.String) (ret js.String) {
	bindings.CallGetURL(
		js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetURL calls the function "WEBEXT.runtime.getURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetURL(path js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetURL(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// Id returns the value of property "WEBEXT.runtime.id".
//
// The returned bool will be false if there is no such property.
func Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetId(
		js.Pointer(&ret),
	)

	return
}

// SetId sets the value of property "WEBEXT.runtime.id" to val.
//
// It returns false if the property cannot be set.
func SetId(val js.String) bool {
	return js.True == bindings.SetId(
		val.Ref())
}

// LastError returns the value of property "WEBEXT.runtime.lastError".
//
// The returned bool will be false if there is no such property.
func LastError() (ret LastErrorProperty, ok bool) {
	ok = js.True == bindings.GetLastError(
		js.Pointer(&ret),
	)

	return
}

// SetLastError sets the value of property "WEBEXT.runtime.lastError" to val.
//
// It returns false if the property cannot be set.
func SetLastError(val LastErrorProperty) bool {
	return js.True == bindings.SetLastError(
		js.Pointer(&val))
}

type OnBrowserUpdateAvailableEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnBrowserUpdateAvailableEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBrowserUpdateAvailableEventCallbackFunc) DispatchCallback(
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

type OnBrowserUpdateAvailableEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnBrowserUpdateAvailableEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBrowserUpdateAvailableEventCallback[T]) DispatchCallback(
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

// HasFuncOnBrowserUpdateAvailable returns true if the function "WEBEXT.runtime.onBrowserUpdateAvailable.addListener" exists.
func HasFuncOnBrowserUpdateAvailable() bool {
	return js.True == bindings.HasFuncOnBrowserUpdateAvailable()
}

// FuncOnBrowserUpdateAvailable returns the function "WEBEXT.runtime.onBrowserUpdateAvailable.addListener".
func FuncOnBrowserUpdateAvailable() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnBrowserUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// OnBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.addListener" directly.
func OnBrowserUpdateAvailable(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnBrowserUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBrowserUpdateAvailable(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBrowserUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBrowserUpdateAvailable returns true if the function "WEBEXT.runtime.onBrowserUpdateAvailable.removeListener" exists.
func HasFuncOffBrowserUpdateAvailable() bool {
	return js.True == bindings.HasFuncOffBrowserUpdateAvailable()
}

// FuncOffBrowserUpdateAvailable returns the function "WEBEXT.runtime.onBrowserUpdateAvailable.removeListener".
func FuncOffBrowserUpdateAvailable() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffBrowserUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// OffBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.removeListener" directly.
func OffBrowserUpdateAvailable(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffBrowserUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBrowserUpdateAvailable(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBrowserUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBrowserUpdateAvailable returns true if the function "WEBEXT.runtime.onBrowserUpdateAvailable.hasListener" exists.
func HasFuncHasOnBrowserUpdateAvailable() bool {
	return js.True == bindings.HasFuncHasOnBrowserUpdateAvailable()
}

// FuncHasOnBrowserUpdateAvailable returns the function "WEBEXT.runtime.onBrowserUpdateAvailable.hasListener".
func FuncHasOnBrowserUpdateAvailable() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnBrowserUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// HasOnBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.hasListener" directly.
func HasOnBrowserUpdateAvailable(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnBrowserUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBrowserUpdateAvailable calls the function "WEBEXT.runtime.onBrowserUpdateAvailable.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBrowserUpdateAvailable(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBrowserUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConnectEventCallbackFunc func(this js.Ref, port *Port) js.Ref

func (fn OnConnectEventCallbackFunc) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConnectEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

type OnConnectEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, port *Port) js.Ref
	Arg T
}

func (cb *OnConnectEventCallback[T]) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConnectEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

// HasFuncOnConnect returns true if the function "WEBEXT.runtime.onConnect.addListener" exists.
func HasFuncOnConnect() bool {
	return js.True == bindings.HasFuncOnConnect()
}

// FuncOnConnect returns the function "WEBEXT.runtime.onConnect.addListener".
func FuncOnConnect() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOnConnect(
		js.Pointer(&fn),
	)
	return
}

// OnConnect calls the function "WEBEXT.runtime.onConnect.addListener" directly.
func OnConnect(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOnConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConnect calls the function "WEBEXT.runtime.onConnect.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConnect(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConnect returns true if the function "WEBEXT.runtime.onConnect.removeListener" exists.
func HasFuncOffConnect() bool {
	return js.True == bindings.HasFuncOffConnect()
}

// FuncOffConnect returns the function "WEBEXT.runtime.onConnect.removeListener".
func FuncOffConnect() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOffConnect(
		js.Pointer(&fn),
	)
	return
}

// OffConnect calls the function "WEBEXT.runtime.onConnect.removeListener" directly.
func OffConnect(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOffConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConnect calls the function "WEBEXT.runtime.onConnect.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConnect(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConnect returns true if the function "WEBEXT.runtime.onConnect.hasListener" exists.
func HasFuncHasOnConnect() bool {
	return js.True == bindings.HasFuncHasOnConnect()
}

// FuncHasOnConnect returns the function "WEBEXT.runtime.onConnect.hasListener".
func FuncHasOnConnect() (fn js.Func[func(callback js.Func[func(port *Port)]) bool]) {
	bindings.FuncHasOnConnect(
		js.Pointer(&fn),
	)
	return
}

// HasOnConnect calls the function "WEBEXT.runtime.onConnect.hasListener" directly.
func HasOnConnect(callback js.Func[func(port *Port)]) (ret bool) {
	bindings.CallHasOnConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConnect calls the function "WEBEXT.runtime.onConnect.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConnect(callback js.Func[func(port *Port)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConnectExternalEventCallbackFunc func(this js.Ref, port *Port) js.Ref

func (fn OnConnectExternalEventCallbackFunc) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConnectExternalEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

type OnConnectExternalEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, port *Port) js.Ref
	Arg T
}

func (cb *OnConnectExternalEventCallback[T]) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConnectExternalEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

// HasFuncOnConnectExternal returns true if the function "WEBEXT.runtime.onConnectExternal.addListener" exists.
func HasFuncOnConnectExternal() bool {
	return js.True == bindings.HasFuncOnConnectExternal()
}

// FuncOnConnectExternal returns the function "WEBEXT.runtime.onConnectExternal.addListener".
func FuncOnConnectExternal() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOnConnectExternal(
		js.Pointer(&fn),
	)
	return
}

// OnConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.addListener" directly.
func OnConnectExternal(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOnConnectExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConnectExternal(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConnectExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConnectExternal returns true if the function "WEBEXT.runtime.onConnectExternal.removeListener" exists.
func HasFuncOffConnectExternal() bool {
	return js.True == bindings.HasFuncOffConnectExternal()
}

// FuncOffConnectExternal returns the function "WEBEXT.runtime.onConnectExternal.removeListener".
func FuncOffConnectExternal() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOffConnectExternal(
		js.Pointer(&fn),
	)
	return
}

// OffConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.removeListener" directly.
func OffConnectExternal(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOffConnectExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConnectExternal(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConnectExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConnectExternal returns true if the function "WEBEXT.runtime.onConnectExternal.hasListener" exists.
func HasFuncHasOnConnectExternal() bool {
	return js.True == bindings.HasFuncHasOnConnectExternal()
}

// FuncHasOnConnectExternal returns the function "WEBEXT.runtime.onConnectExternal.hasListener".
func FuncHasOnConnectExternal() (fn js.Func[func(callback js.Func[func(port *Port)]) bool]) {
	bindings.FuncHasOnConnectExternal(
		js.Pointer(&fn),
	)
	return
}

// HasOnConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.hasListener" directly.
func HasOnConnectExternal(callback js.Func[func(port *Port)]) (ret bool) {
	bindings.CallHasOnConnectExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConnectExternal calls the function "WEBEXT.runtime.onConnectExternal.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConnectExternal(callback js.Func[func(port *Port)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConnectExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConnectNativeEventCallbackFunc func(this js.Ref, port *Port) js.Ref

func (fn OnConnectNativeEventCallbackFunc) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConnectNativeEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

type OnConnectNativeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, port *Port) js.Ref
	Arg T
}

func (cb *OnConnectNativeEventCallback[T]) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConnectNativeEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

// HasFuncOnConnectNative returns true if the function "WEBEXT.runtime.onConnectNative.addListener" exists.
func HasFuncOnConnectNative() bool {
	return js.True == bindings.HasFuncOnConnectNative()
}

// FuncOnConnectNative returns the function "WEBEXT.runtime.onConnectNative.addListener".
func FuncOnConnectNative() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOnConnectNative(
		js.Pointer(&fn),
	)
	return
}

// OnConnectNative calls the function "WEBEXT.runtime.onConnectNative.addListener" directly.
func OnConnectNative(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOnConnectNative(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConnectNative calls the function "WEBEXT.runtime.onConnectNative.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConnectNative(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConnectNative(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConnectNative returns true if the function "WEBEXT.runtime.onConnectNative.removeListener" exists.
func HasFuncOffConnectNative() bool {
	return js.True == bindings.HasFuncOffConnectNative()
}

// FuncOffConnectNative returns the function "WEBEXT.runtime.onConnectNative.removeListener".
func FuncOffConnectNative() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOffConnectNative(
		js.Pointer(&fn),
	)
	return
}

// OffConnectNative calls the function "WEBEXT.runtime.onConnectNative.removeListener" directly.
func OffConnectNative(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOffConnectNative(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConnectNative calls the function "WEBEXT.runtime.onConnectNative.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConnectNative(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConnectNative(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConnectNative returns true if the function "WEBEXT.runtime.onConnectNative.hasListener" exists.
func HasFuncHasOnConnectNative() bool {
	return js.True == bindings.HasFuncHasOnConnectNative()
}

// FuncHasOnConnectNative returns the function "WEBEXT.runtime.onConnectNative.hasListener".
func FuncHasOnConnectNative() (fn js.Func[func(callback js.Func[func(port *Port)]) bool]) {
	bindings.FuncHasOnConnectNative(
		js.Pointer(&fn),
	)
	return
}

// HasOnConnectNative calls the function "WEBEXT.runtime.onConnectNative.hasListener" directly.
func HasOnConnectNative(callback js.Func[func(port *Port)]) (ret bool) {
	bindings.CallHasOnConnectNative(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConnectNative calls the function "WEBEXT.runtime.onConnectNative.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConnectNative(callback js.Func[func(port *Port)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConnectNative(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInstalledEventCallbackFunc func(this js.Ref, details *OnInstalledArgDetails) js.Ref

func (fn OnInstalledEventCallbackFunc) Register() js.Func[func(details *OnInstalledArgDetails)] {
	return js.RegisterCallback[func(details *OnInstalledArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInstalledEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnInstalledArgDetails
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

type OnInstalledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnInstalledArgDetails) js.Ref
	Arg T
}

func (cb *OnInstalledEventCallback[T]) Register() js.Func[func(details *OnInstalledArgDetails)] {
	return js.RegisterCallback[func(details *OnInstalledArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInstalledEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnInstalledArgDetails
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

// HasFuncOnInstalled returns true if the function "WEBEXT.runtime.onInstalled.addListener" exists.
func HasFuncOnInstalled() bool {
	return js.True == bindings.HasFuncOnInstalled()
}

// FuncOnInstalled returns the function "WEBEXT.runtime.onInstalled.addListener".
func FuncOnInstalled() (fn js.Func[func(callback js.Func[func(details *OnInstalledArgDetails)])]) {
	bindings.FuncOnInstalled(
		js.Pointer(&fn),
	)
	return
}

// OnInstalled calls the function "WEBEXT.runtime.onInstalled.addListener" directly.
func OnInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret js.Void) {
	bindings.CallOnInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInstalled calls the function "WEBEXT.runtime.onInstalled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInstalled returns true if the function "WEBEXT.runtime.onInstalled.removeListener" exists.
func HasFuncOffInstalled() bool {
	return js.True == bindings.HasFuncOffInstalled()
}

// FuncOffInstalled returns the function "WEBEXT.runtime.onInstalled.removeListener".
func FuncOffInstalled() (fn js.Func[func(callback js.Func[func(details *OnInstalledArgDetails)])]) {
	bindings.FuncOffInstalled(
		js.Pointer(&fn),
	)
	return
}

// OffInstalled calls the function "WEBEXT.runtime.onInstalled.removeListener" directly.
func OffInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret js.Void) {
	bindings.CallOffInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInstalled calls the function "WEBEXT.runtime.onInstalled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInstalled returns true if the function "WEBEXT.runtime.onInstalled.hasListener" exists.
func HasFuncHasOnInstalled() bool {
	return js.True == bindings.HasFuncHasOnInstalled()
}

// FuncHasOnInstalled returns the function "WEBEXT.runtime.onInstalled.hasListener".
func FuncHasOnInstalled() (fn js.Func[func(callback js.Func[func(details *OnInstalledArgDetails)]) bool]) {
	bindings.FuncHasOnInstalled(
		js.Pointer(&fn),
	)
	return
}

// HasOnInstalled calls the function "WEBEXT.runtime.onInstalled.hasListener" directly.
func HasOnInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret bool) {
	bindings.CallHasOnInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInstalled calls the function "WEBEXT.runtime.onInstalled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInstalled(callback js.Func[func(details *OnInstalledArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMessageEventCallbackFunc func(this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref

func (fn OnMessageEventCallbackFunc) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref
	Arg T
}

func (cb *OnMessageEventCallback[T]) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMessage returns true if the function "WEBEXT.runtime.onMessage.addListener" exists.
func HasFuncOnMessage() bool {
	return js.True == bindings.HasFuncOnMessage()
}

// FuncOnMessage returns the function "WEBEXT.runtime.onMessage.addListener".
func FuncOnMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOnMessage(
		js.Pointer(&fn),
	)
	return
}

// OnMessage calls the function "WEBEXT.runtime.onMessage.addListener" directly.
func OnMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessage calls the function "WEBEXT.runtime.onMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessage returns true if the function "WEBEXT.runtime.onMessage.removeListener" exists.
func HasFuncOffMessage() bool {
	return js.True == bindings.HasFuncOffMessage()
}

// FuncOffMessage returns the function "WEBEXT.runtime.onMessage.removeListener".
func FuncOffMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOffMessage(
		js.Pointer(&fn),
	)
	return
}

// OffMessage calls the function "WEBEXT.runtime.onMessage.removeListener" directly.
func OffMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOffMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessage calls the function "WEBEXT.runtime.onMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessage returns true if the function "WEBEXT.runtime.onMessage.hasListener" exists.
func HasFuncHasOnMessage() bool {
	return js.True == bindings.HasFuncHasOnMessage()
}

// FuncHasOnMessage returns the function "WEBEXT.runtime.onMessage.hasListener".
func FuncHasOnMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) bool]) {
	bindings.FuncHasOnMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessage calls the function "WEBEXT.runtime.onMessage.hasListener" directly.
func HasOnMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool) {
	bindings.CallHasOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessage calls the function "WEBEXT.runtime.onMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMessageExternalEventCallbackFunc func(this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref

func (fn OnMessageExternalEventCallbackFunc) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageExternalEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMessageExternalEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref
	Arg T
}

func (cb *OnMessageExternalEventCallback[T]) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageExternalEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMessageExternal returns true if the function "WEBEXT.runtime.onMessageExternal.addListener" exists.
func HasFuncOnMessageExternal() bool {
	return js.True == bindings.HasFuncOnMessageExternal()
}

// FuncOnMessageExternal returns the function "WEBEXT.runtime.onMessageExternal.addListener".
func FuncOnMessageExternal() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOnMessageExternal(
		js.Pointer(&fn),
	)
	return
}

// OnMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.addListener" directly.
func OnMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOnMessageExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessageExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessageExternal returns true if the function "WEBEXT.runtime.onMessageExternal.removeListener" exists.
func HasFuncOffMessageExternal() bool {
	return js.True == bindings.HasFuncOffMessageExternal()
}

// FuncOffMessageExternal returns the function "WEBEXT.runtime.onMessageExternal.removeListener".
func FuncOffMessageExternal() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOffMessageExternal(
		js.Pointer(&fn),
	)
	return
}

// OffMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.removeListener" directly.
func OffMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOffMessageExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessageExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessageExternal returns true if the function "WEBEXT.runtime.onMessageExternal.hasListener" exists.
func HasFuncHasOnMessageExternal() bool {
	return js.True == bindings.HasFuncHasOnMessageExternal()
}

// FuncHasOnMessageExternal returns the function "WEBEXT.runtime.onMessageExternal.hasListener".
func FuncHasOnMessageExternal() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) bool]) {
	bindings.FuncHasOnMessageExternal(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.hasListener" directly.
func HasOnMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool) {
	bindings.CallHasOnMessageExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessageExternal calls the function "WEBEXT.runtime.onMessageExternal.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessageExternal(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessageExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRestartRequiredEventCallbackFunc func(this js.Ref, reason OnRestartRequiredReason) js.Ref

func (fn OnRestartRequiredEventCallbackFunc) Register() js.Func[func(reason OnRestartRequiredReason)] {
	return js.RegisterCallback[func(reason OnRestartRequiredReason)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRestartRequiredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		OnRestartRequiredReason(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRestartRequiredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reason OnRestartRequiredReason) js.Ref
	Arg T
}

func (cb *OnRestartRequiredEventCallback[T]) Register() js.Func[func(reason OnRestartRequiredReason)] {
	return js.RegisterCallback[func(reason OnRestartRequiredReason)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRestartRequiredEventCallback[T]) DispatchCallback(
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

		OnRestartRequiredReason(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRestartRequired returns true if the function "WEBEXT.runtime.onRestartRequired.addListener" exists.
func HasFuncOnRestartRequired() bool {
	return js.True == bindings.HasFuncOnRestartRequired()
}

// FuncOnRestartRequired returns the function "WEBEXT.runtime.onRestartRequired.addListener".
func FuncOnRestartRequired() (fn js.Func[func(callback js.Func[func(reason OnRestartRequiredReason)])]) {
	bindings.FuncOnRestartRequired(
		js.Pointer(&fn),
	)
	return
}

// OnRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.addListener" directly.
func OnRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret js.Void) {
	bindings.CallOnRestartRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRestartRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRestartRequired returns true if the function "WEBEXT.runtime.onRestartRequired.removeListener" exists.
func HasFuncOffRestartRequired() bool {
	return js.True == bindings.HasFuncOffRestartRequired()
}

// FuncOffRestartRequired returns the function "WEBEXT.runtime.onRestartRequired.removeListener".
func FuncOffRestartRequired() (fn js.Func[func(callback js.Func[func(reason OnRestartRequiredReason)])]) {
	bindings.FuncOffRestartRequired(
		js.Pointer(&fn),
	)
	return
}

// OffRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.removeListener" directly.
func OffRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret js.Void) {
	bindings.CallOffRestartRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRestartRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRestartRequired returns true if the function "WEBEXT.runtime.onRestartRequired.hasListener" exists.
func HasFuncHasOnRestartRequired() bool {
	return js.True == bindings.HasFuncHasOnRestartRequired()
}

// FuncHasOnRestartRequired returns the function "WEBEXT.runtime.onRestartRequired.hasListener".
func FuncHasOnRestartRequired() (fn js.Func[func(callback js.Func[func(reason OnRestartRequiredReason)]) bool]) {
	bindings.FuncHasOnRestartRequired(
		js.Pointer(&fn),
	)
	return
}

// HasOnRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.hasListener" directly.
func HasOnRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret bool) {
	bindings.CallHasOnRestartRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRestartRequired calls the function "WEBEXT.runtime.onRestartRequired.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRestartRequired(callback js.Func[func(reason OnRestartRequiredReason)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRestartRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStartupEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnStartupEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStartupEventCallbackFunc) DispatchCallback(
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

type OnStartupEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnStartupEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStartupEventCallback[T]) DispatchCallback(
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

// HasFuncOnStartup returns true if the function "WEBEXT.runtime.onStartup.addListener" exists.
func HasFuncOnStartup() bool {
	return js.True == bindings.HasFuncOnStartup()
}

// FuncOnStartup returns the function "WEBEXT.runtime.onStartup.addListener".
func FuncOnStartup() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnStartup(
		js.Pointer(&fn),
	)
	return
}

// OnStartup calls the function "WEBEXT.runtime.onStartup.addListener" directly.
func OnStartup(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnStartup(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStartup calls the function "WEBEXT.runtime.onStartup.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStartup(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStartup(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStartup returns true if the function "WEBEXT.runtime.onStartup.removeListener" exists.
func HasFuncOffStartup() bool {
	return js.True == bindings.HasFuncOffStartup()
}

// FuncOffStartup returns the function "WEBEXT.runtime.onStartup.removeListener".
func FuncOffStartup() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffStartup(
		js.Pointer(&fn),
	)
	return
}

// OffStartup calls the function "WEBEXT.runtime.onStartup.removeListener" directly.
func OffStartup(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffStartup(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStartup calls the function "WEBEXT.runtime.onStartup.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStartup(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStartup(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStartup returns true if the function "WEBEXT.runtime.onStartup.hasListener" exists.
func HasFuncHasOnStartup() bool {
	return js.True == bindings.HasFuncHasOnStartup()
}

// FuncHasOnStartup returns the function "WEBEXT.runtime.onStartup.hasListener".
func FuncHasOnStartup() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnStartup(
		js.Pointer(&fn),
	)
	return
}

// HasOnStartup calls the function "WEBEXT.runtime.onStartup.hasListener" directly.
func HasOnStartup(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnStartup(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStartup calls the function "WEBEXT.runtime.onStartup.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStartup(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStartup(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSuspendEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnSuspendEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSuspendEventCallbackFunc) DispatchCallback(
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

type OnSuspendEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnSuspendEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSuspendEventCallback[T]) DispatchCallback(
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

// HasFuncOnSuspend returns true if the function "WEBEXT.runtime.onSuspend.addListener" exists.
func HasFuncOnSuspend() bool {
	return js.True == bindings.HasFuncOnSuspend()
}

// FuncOnSuspend returns the function "WEBEXT.runtime.onSuspend.addListener".
func FuncOnSuspend() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnSuspend(
		js.Pointer(&fn),
	)
	return
}

// OnSuspend calls the function "WEBEXT.runtime.onSuspend.addListener" directly.
func OnSuspend(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnSuspend(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSuspend calls the function "WEBEXT.runtime.onSuspend.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSuspend(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSuspend(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSuspend returns true if the function "WEBEXT.runtime.onSuspend.removeListener" exists.
func HasFuncOffSuspend() bool {
	return js.True == bindings.HasFuncOffSuspend()
}

// FuncOffSuspend returns the function "WEBEXT.runtime.onSuspend.removeListener".
func FuncOffSuspend() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffSuspend(
		js.Pointer(&fn),
	)
	return
}

// OffSuspend calls the function "WEBEXT.runtime.onSuspend.removeListener" directly.
func OffSuspend(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffSuspend(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSuspend calls the function "WEBEXT.runtime.onSuspend.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSuspend(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSuspend(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSuspend returns true if the function "WEBEXT.runtime.onSuspend.hasListener" exists.
func HasFuncHasOnSuspend() bool {
	return js.True == bindings.HasFuncHasOnSuspend()
}

// FuncHasOnSuspend returns the function "WEBEXT.runtime.onSuspend.hasListener".
func FuncHasOnSuspend() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnSuspend(
		js.Pointer(&fn),
	)
	return
}

// HasOnSuspend calls the function "WEBEXT.runtime.onSuspend.hasListener" directly.
func HasOnSuspend(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnSuspend(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSuspend calls the function "WEBEXT.runtime.onSuspend.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSuspend(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSuspend(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSuspendCanceledEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnSuspendCanceledEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSuspendCanceledEventCallbackFunc) DispatchCallback(
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

type OnSuspendCanceledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnSuspendCanceledEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSuspendCanceledEventCallback[T]) DispatchCallback(
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

// HasFuncOnSuspendCanceled returns true if the function "WEBEXT.runtime.onSuspendCanceled.addListener" exists.
func HasFuncOnSuspendCanceled() bool {
	return js.True == bindings.HasFuncOnSuspendCanceled()
}

// FuncOnSuspendCanceled returns the function "WEBEXT.runtime.onSuspendCanceled.addListener".
func FuncOnSuspendCanceled() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnSuspendCanceled(
		js.Pointer(&fn),
	)
	return
}

// OnSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.addListener" directly.
func OnSuspendCanceled(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnSuspendCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSuspendCanceled(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSuspendCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSuspendCanceled returns true if the function "WEBEXT.runtime.onSuspendCanceled.removeListener" exists.
func HasFuncOffSuspendCanceled() bool {
	return js.True == bindings.HasFuncOffSuspendCanceled()
}

// FuncOffSuspendCanceled returns the function "WEBEXT.runtime.onSuspendCanceled.removeListener".
func FuncOffSuspendCanceled() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffSuspendCanceled(
		js.Pointer(&fn),
	)
	return
}

// OffSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.removeListener" directly.
func OffSuspendCanceled(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffSuspendCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSuspendCanceled(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSuspendCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSuspendCanceled returns true if the function "WEBEXT.runtime.onSuspendCanceled.hasListener" exists.
func HasFuncHasOnSuspendCanceled() bool {
	return js.True == bindings.HasFuncHasOnSuspendCanceled()
}

// FuncHasOnSuspendCanceled returns the function "WEBEXT.runtime.onSuspendCanceled.hasListener".
func FuncHasOnSuspendCanceled() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnSuspendCanceled(
		js.Pointer(&fn),
	)
	return
}

// HasOnSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.hasListener" directly.
func HasOnSuspendCanceled(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnSuspendCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSuspendCanceled calls the function "WEBEXT.runtime.onSuspendCanceled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSuspendCanceled(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSuspendCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUpdateAvailableEventCallbackFunc func(this js.Ref, details js.Any) js.Ref

func (fn OnUpdateAvailableEventCallbackFunc) Register() js.Func[func(details js.Any)] {
	return js.RegisterCallback[func(details js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUpdateAvailableEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUpdateAvailableEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details js.Any) js.Ref
	Arg T
}

func (cb *OnUpdateAvailableEventCallback[T]) Register() js.Func[func(details js.Any)] {
	return js.RegisterCallback[func(details js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUpdateAvailableEventCallback[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUpdateAvailable returns true if the function "WEBEXT.runtime.onUpdateAvailable.addListener" exists.
func HasFuncOnUpdateAvailable() bool {
	return js.True == bindings.HasFuncOnUpdateAvailable()
}

// FuncOnUpdateAvailable returns the function "WEBEXT.runtime.onUpdateAvailable.addListener".
func FuncOnUpdateAvailable() (fn js.Func[func(callback js.Func[func(details js.Any)])]) {
	bindings.FuncOnUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// OnUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.addListener" directly.
func OnUpdateAvailable(callback js.Func[func(details js.Any)]) (ret js.Void) {
	bindings.CallOnUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUpdateAvailable(callback js.Func[func(details js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUpdateAvailable returns true if the function "WEBEXT.runtime.onUpdateAvailable.removeListener" exists.
func HasFuncOffUpdateAvailable() bool {
	return js.True == bindings.HasFuncOffUpdateAvailable()
}

// FuncOffUpdateAvailable returns the function "WEBEXT.runtime.onUpdateAvailable.removeListener".
func FuncOffUpdateAvailable() (fn js.Func[func(callback js.Func[func(details js.Any)])]) {
	bindings.FuncOffUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// OffUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.removeListener" directly.
func OffUpdateAvailable(callback js.Func[func(details js.Any)]) (ret js.Void) {
	bindings.CallOffUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUpdateAvailable(callback js.Func[func(details js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUpdateAvailable returns true if the function "WEBEXT.runtime.onUpdateAvailable.hasListener" exists.
func HasFuncHasOnUpdateAvailable() bool {
	return js.True == bindings.HasFuncHasOnUpdateAvailable()
}

// FuncHasOnUpdateAvailable returns the function "WEBEXT.runtime.onUpdateAvailable.hasListener".
func FuncHasOnUpdateAvailable() (fn js.Func[func(callback js.Func[func(details js.Any)]) bool]) {
	bindings.FuncHasOnUpdateAvailable(
		js.Pointer(&fn),
	)
	return
}

// HasOnUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.hasListener" directly.
func HasOnUpdateAvailable(callback js.Func[func(details js.Any)]) (ret bool) {
	bindings.CallHasOnUpdateAvailable(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUpdateAvailable calls the function "WEBEXT.runtime.onUpdateAvailable.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUpdateAvailable(callback js.Func[func(details js.Any)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUpdateAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUserScriptConnectEventCallbackFunc func(this js.Ref, port *Port) js.Ref

func (fn OnUserScriptConnectEventCallbackFunc) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUserScriptConnectEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

type OnUserScriptConnectEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, port *Port) js.Ref
	Arg T
}

func (cb *OnUserScriptConnectEventCallback[T]) Register() js.Func[func(port *Port)] {
	return js.RegisterCallback[func(port *Port)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUserScriptConnectEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Port
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

// HasFuncOnUserScriptConnect returns true if the function "WEBEXT.runtime.onUserScriptConnect.addListener" exists.
func HasFuncOnUserScriptConnect() bool {
	return js.True == bindings.HasFuncOnUserScriptConnect()
}

// FuncOnUserScriptConnect returns the function "WEBEXT.runtime.onUserScriptConnect.addListener".
func FuncOnUserScriptConnect() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOnUserScriptConnect(
		js.Pointer(&fn),
	)
	return
}

// OnUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.addListener" directly.
func OnUserScriptConnect(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOnUserScriptConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUserScriptConnect(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUserScriptConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUserScriptConnect returns true if the function "WEBEXT.runtime.onUserScriptConnect.removeListener" exists.
func HasFuncOffUserScriptConnect() bool {
	return js.True == bindings.HasFuncOffUserScriptConnect()
}

// FuncOffUserScriptConnect returns the function "WEBEXT.runtime.onUserScriptConnect.removeListener".
func FuncOffUserScriptConnect() (fn js.Func[func(callback js.Func[func(port *Port)])]) {
	bindings.FuncOffUserScriptConnect(
		js.Pointer(&fn),
	)
	return
}

// OffUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.removeListener" directly.
func OffUserScriptConnect(callback js.Func[func(port *Port)]) (ret js.Void) {
	bindings.CallOffUserScriptConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUserScriptConnect(callback js.Func[func(port *Port)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUserScriptConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUserScriptConnect returns true if the function "WEBEXT.runtime.onUserScriptConnect.hasListener" exists.
func HasFuncHasOnUserScriptConnect() bool {
	return js.True == bindings.HasFuncHasOnUserScriptConnect()
}

// FuncHasOnUserScriptConnect returns the function "WEBEXT.runtime.onUserScriptConnect.hasListener".
func FuncHasOnUserScriptConnect() (fn js.Func[func(callback js.Func[func(port *Port)]) bool]) {
	bindings.FuncHasOnUserScriptConnect(
		js.Pointer(&fn),
	)
	return
}

// HasOnUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.hasListener" directly.
func HasOnUserScriptConnect(callback js.Func[func(port *Port)]) (ret bool) {
	bindings.CallHasOnUserScriptConnect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUserScriptConnect calls the function "WEBEXT.runtime.onUserScriptConnect.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUserScriptConnect(callback js.Func[func(port *Port)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUserScriptConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUserScriptMessageEventCallbackFunc func(this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref

func (fn OnUserScriptMessageEventCallbackFunc) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUserScriptMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUserScriptMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, message js.Any, sender *MessageSender, sendResponse js.Func[func()]) js.Ref
	Arg T
}

func (cb *OnUserScriptMessageEventCallback[T]) Register() js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool] {
	return js.RegisterCallback[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUserScriptMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUserScriptMessage returns true if the function "WEBEXT.runtime.onUserScriptMessage.addListener" exists.
func HasFuncOnUserScriptMessage() bool {
	return js.True == bindings.HasFuncOnUserScriptMessage()
}

// FuncOnUserScriptMessage returns the function "WEBEXT.runtime.onUserScriptMessage.addListener".
func FuncOnUserScriptMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOnUserScriptMessage(
		js.Pointer(&fn),
	)
	return
}

// OnUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.addListener" directly.
func OnUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOnUserScriptMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUserScriptMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUserScriptMessage returns true if the function "WEBEXT.runtime.onUserScriptMessage.removeListener" exists.
func HasFuncOffUserScriptMessage() bool {
	return js.True == bindings.HasFuncOffUserScriptMessage()
}

// FuncOffUserScriptMessage returns the function "WEBEXT.runtime.onUserScriptMessage.removeListener".
func FuncOffUserScriptMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool])]) {
	bindings.FuncOffUserScriptMessage(
		js.Pointer(&fn),
	)
	return
}

// OffUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.removeListener" directly.
func OffUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void) {
	bindings.CallOffUserScriptMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUserScriptMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUserScriptMessage returns true if the function "WEBEXT.runtime.onUserScriptMessage.hasListener" exists.
func HasFuncHasOnUserScriptMessage() bool {
	return js.True == bindings.HasFuncHasOnUserScriptMessage()
}

// FuncHasOnUserScriptMessage returns the function "WEBEXT.runtime.onUserScriptMessage.hasListener".
func FuncHasOnUserScriptMessage() (fn js.Func[func(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) bool]) {
	bindings.FuncHasOnUserScriptMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.hasListener" directly.
func HasOnUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool) {
	bindings.CallHasOnUserScriptMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUserScriptMessage calls the function "WEBEXT.runtime.onUserScriptMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUserScriptMessage(callback js.Func[func(message js.Any, sender *MessageSender, sendResponse js.Func[func()]) bool]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUserScriptMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenOptionsPage returns true if the function "WEBEXT.runtime.openOptionsPage" exists.
func HasFuncOpenOptionsPage() bool {
	return js.True == bindings.HasFuncOpenOptionsPage()
}

// FuncOpenOptionsPage returns the function "WEBEXT.runtime.openOptionsPage".
func FuncOpenOptionsPage() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncOpenOptionsPage(
		js.Pointer(&fn),
	)
	return
}

// OpenOptionsPage calls the function "WEBEXT.runtime.openOptionsPage" directly.
func OpenOptionsPage() (ret js.Promise[js.Void]) {
	bindings.CallOpenOptionsPage(
		js.Pointer(&ret),
	)

	return
}

// TryOpenOptionsPage calls the function "WEBEXT.runtime.openOptionsPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenOptionsPage() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenOptionsPage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReload returns true if the function "WEBEXT.runtime.reload" exists.
func HasFuncReload() bool {
	return js.True == bindings.HasFuncReload()
}

// FuncReload returns the function "WEBEXT.runtime.reload".
func FuncReload() (fn js.Func[func()]) {
	bindings.FuncReload(
		js.Pointer(&fn),
	)
	return
}

// Reload calls the function "WEBEXT.runtime.reload" directly.
func Reload() (ret js.Void) {
	bindings.CallReload(
		js.Pointer(&ret),
	)

	return
}

// TryReload calls the function "WEBEXT.runtime.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReload() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReload(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestUpdateCheck returns true if the function "WEBEXT.runtime.requestUpdateCheck" exists.
func HasFuncRequestUpdateCheck() bool {
	return js.True == bindings.HasFuncRequestUpdateCheck()
}

// FuncRequestUpdateCheck returns the function "WEBEXT.runtime.requestUpdateCheck".
func FuncRequestUpdateCheck() (fn js.Func[func() js.Promise[RequestUpdateCheckReturnType]]) {
	bindings.FuncRequestUpdateCheck(
		js.Pointer(&fn),
	)
	return
}

// RequestUpdateCheck calls the function "WEBEXT.runtime.requestUpdateCheck" directly.
func RequestUpdateCheck() (ret js.Promise[RequestUpdateCheckReturnType]) {
	bindings.CallRequestUpdateCheck(
		js.Pointer(&ret),
	)

	return
}

// TryRequestUpdateCheck calls the function "WEBEXT.runtime.requestUpdateCheck"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestUpdateCheck() (ret js.Promise[RequestUpdateCheckReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestUpdateCheck(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestart returns true if the function "WEBEXT.runtime.restart" exists.
func HasFuncRestart() bool {
	return js.True == bindings.HasFuncRestart()
}

// FuncRestart returns the function "WEBEXT.runtime.restart".
func FuncRestart() (fn js.Func[func()]) {
	bindings.FuncRestart(
		js.Pointer(&fn),
	)
	return
}

// Restart calls the function "WEBEXT.runtime.restart" directly.
func Restart() (ret js.Void) {
	bindings.CallRestart(
		js.Pointer(&ret),
	)

	return
}

// TryRestart calls the function "WEBEXT.runtime.restart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestart(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestartAfterDelay returns true if the function "WEBEXT.runtime.restartAfterDelay" exists.
func HasFuncRestartAfterDelay() bool {
	return js.True == bindings.HasFuncRestartAfterDelay()
}

// FuncRestartAfterDelay returns the function "WEBEXT.runtime.restartAfterDelay".
func FuncRestartAfterDelay() (fn js.Func[func(seconds int64) js.Promise[js.Void]]) {
	bindings.FuncRestartAfterDelay(
		js.Pointer(&fn),
	)
	return
}

// RestartAfterDelay calls the function "WEBEXT.runtime.restartAfterDelay" directly.
func RestartAfterDelay(seconds int64) (ret js.Promise[js.Void]) {
	bindings.CallRestartAfterDelay(
		js.Pointer(&ret),
		float64(seconds),
	)

	return
}

// TryRestartAfterDelay calls the function "WEBEXT.runtime.restartAfterDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestartAfterDelay(seconds int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestartAfterDelay(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(seconds),
	)

	return
}

// HasFuncSendMessage returns true if the function "WEBEXT.runtime.sendMessage" exists.
func HasFuncSendMessage() bool {
	return js.True == bindings.HasFuncSendMessage()
}

// FuncSendMessage returns the function "WEBEXT.runtime.sendMessage".
func FuncSendMessage() (fn js.Func[func(extensionId js.String, message js.Any, options SendMessageArgOptions) js.Promise[js.Any]]) {
	bindings.FuncSendMessage(
		js.Pointer(&fn),
	)
	return
}

// SendMessage calls the function "WEBEXT.runtime.sendMessage" directly.
func SendMessage(extensionId js.String, message js.Any, options SendMessageArgOptions) (ret js.Promise[js.Any]) {
	bindings.CallSendMessage(
		js.Pointer(&ret),
		extensionId.Ref(),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySendMessage calls the function "WEBEXT.runtime.sendMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendMessage(extensionId js.String, message js.Any, options SendMessageArgOptions) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSendNativeMessage returns true if the function "WEBEXT.runtime.sendNativeMessage" exists.
func HasFuncSendNativeMessage() bool {
	return js.True == bindings.HasFuncSendNativeMessage()
}

// FuncSendNativeMessage returns the function "WEBEXT.runtime.sendNativeMessage".
func FuncSendNativeMessage() (fn js.Func[func(application js.String, message js.Any) js.Promise[js.Any]]) {
	bindings.FuncSendNativeMessage(
		js.Pointer(&fn),
	)
	return
}

// SendNativeMessage calls the function "WEBEXT.runtime.sendNativeMessage" directly.
func SendNativeMessage(application js.String, message js.Any) (ret js.Promise[js.Any]) {
	bindings.CallSendNativeMessage(
		js.Pointer(&ret),
		application.Ref(),
		message.Ref(),
	)

	return
}

// TrySendNativeMessage calls the function "WEBEXT.runtime.sendNativeMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendNativeMessage(application js.String, message js.Any) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendNativeMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		application.Ref(),
		message.Ref(),
	)

	return
}

// HasFuncSetUninstallURL returns true if the function "WEBEXT.runtime.setUninstallURL" exists.
func HasFuncSetUninstallURL() bool {
	return js.True == bindings.HasFuncSetUninstallURL()
}

// FuncSetUninstallURL returns the function "WEBEXT.runtime.setUninstallURL".
func FuncSetUninstallURL() (fn js.Func[func(url js.String) js.Promise[js.Void]]) {
	bindings.FuncSetUninstallURL(
		js.Pointer(&fn),
	)
	return
}

// SetUninstallURL calls the function "WEBEXT.runtime.setUninstallURL" directly.
func SetUninstallURL(url js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetUninstallURL(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TrySetUninstallURL calls the function "WEBEXT.runtime.setUninstallURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetUninstallURL(url js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetUninstallURL(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

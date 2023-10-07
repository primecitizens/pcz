// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/runtime store_ConnectArgConnectInfo
//go:noescape
func ConnectArgConnectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_ConnectArgConnectInfo
//go:noescape
func ConnectArgConnectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime constof_ContextType
//go:noescape
func ConstOfContextType(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime store_ContextFilter
//go:noescape
func ContextFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_ContextFilter
//go:noescape
func ContextFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime store_ExtensionContext
//go:noescape
func ExtensionContextJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_ExtensionContext
//go:noescape
func ExtensionContextJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime store_LastErrorProperty
//go:noescape
func LastErrorPropertyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_LastErrorProperty
//go:noescape
func LastErrorPropertyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime store_MessageSender
//go:noescape
func MessageSenderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_MessageSender
//go:noescape
func MessageSenderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime constof_OnInstalledReason
//go:noescape
func ConstOfOnInstalledReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime store_OnInstalledArgDetails
//go:noescape
func OnInstalledArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_OnInstalledArgDetails
//go:noescape
func OnInstalledArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime constof_OnRestartRequiredReason
//go:noescape
func ConstOfOnRestartRequiredReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime constof_PlatformArch
//go:noescape
func ConstOfPlatformArch(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime constof_PlatformNaclArch
//go:noescape
func ConstOfPlatformNaclArch(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime constof_PlatformOs
//go:noescape
func ConstOfPlatformOs(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime store_PlatformInfo
//go:noescape
func PlatformInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_PlatformInfo
//go:noescape
func PlatformInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime store_Port
//go:noescape
func PortJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_Port
//go:noescape
func PortJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime constof_RequestUpdateCheckStatus
//go:noescape
func ConstOfRequestUpdateCheckStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/runtime store_RequestUpdateCheckReturnType
//go:noescape
func RequestUpdateCheckReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_RequestUpdateCheckReturnType
//go:noescape
func RequestUpdateCheckReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime store_SendMessageArgOptions
//go:noescape
func SendMessageArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/runtime load_SendMessageArgOptions
//go:noescape
func SendMessageArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	connectInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	connectInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_ConnectNative
//go:noescape
func HasFuncConnectNative() js.Ref

//go:wasmimport plat/js/webext/runtime func_ConnectNative
//go:noescape
func FuncConnectNative(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_ConnectNative
//go:noescape
func CallConnectNative(
	retPtr unsafe.Pointer,
	application js.Ref)

//go:wasmimport plat/js/webext/runtime try_ConnectNative
//go:noescape
func TryConnectNative(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	application js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetBackgroundPage
//go:noescape
func HasFuncGetBackgroundPage() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetBackgroundPage
//go:noescape
func FuncGetBackgroundPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetBackgroundPage
//go:noescape
func CallGetBackgroundPage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_GetBackgroundPage
//go:noescape
func TryGetBackgroundPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetContexts
//go:noescape
func HasFuncGetContexts() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetContexts
//go:noescape
func FuncGetContexts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetContexts
//go:noescape
func CallGetContexts(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_GetContexts
//go:noescape
func TryGetContexts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetManifest
//go:noescape
func HasFuncGetManifest() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetManifest
//go:noescape
func FuncGetManifest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetManifest
//go:noescape
func CallGetManifest(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_GetManifest
//go:noescape
func TryGetManifest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetPackageDirectoryEntry
//go:noescape
func HasFuncGetPackageDirectoryEntry() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetPackageDirectoryEntry
//go:noescape
func FuncGetPackageDirectoryEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetPackageDirectoryEntry
//go:noescape
func CallGetPackageDirectoryEntry(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_GetPackageDirectoryEntry
//go:noescape
func TryGetPackageDirectoryEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetPlatformInfo
//go:noescape
func HasFuncGetPlatformInfo() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetPlatformInfo
//go:noescape
func FuncGetPlatformInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetPlatformInfo
//go:noescape
func CallGetPlatformInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_GetPlatformInfo
//go:noescape
func TryGetPlatformInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_GetURL
//go:noescape
func HasFuncGetURL() js.Ref

//go:wasmimport plat/js/webext/runtime func_GetURL
//go:noescape
func FuncGetURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_GetURL
//go:noescape
func CallGetURL(
	retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/webext/runtime try_GetURL
//go:noescape
func TryGetURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime get_Id
//go:noescape
func GetId(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/runtime set_Id
//go:noescape
func SetId(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/runtime get_LastError
//go:noescape
func GetLastError(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/runtime set_LastError
//go:noescape
func SetLastError(
	val unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/runtime has_OnBrowserUpdateAvailable
//go:noescape
func HasFuncOnBrowserUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnBrowserUpdateAvailable
//go:noescape
func FuncOnBrowserUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnBrowserUpdateAvailable
//go:noescape
func CallOnBrowserUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnBrowserUpdateAvailable
//go:noescape
func TryOnBrowserUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffBrowserUpdateAvailable
//go:noescape
func HasFuncOffBrowserUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffBrowserUpdateAvailable
//go:noescape
func FuncOffBrowserUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffBrowserUpdateAvailable
//go:noescape
func CallOffBrowserUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffBrowserUpdateAvailable
//go:noescape
func TryOffBrowserUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnBrowserUpdateAvailable
//go:noescape
func HasFuncHasOnBrowserUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnBrowserUpdateAvailable
//go:noescape
func FuncHasOnBrowserUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnBrowserUpdateAvailable
//go:noescape
func CallHasOnBrowserUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnBrowserUpdateAvailable
//go:noescape
func TryHasOnBrowserUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnConnect
//go:noescape
func HasFuncOnConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnConnect
//go:noescape
func FuncOnConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnConnect
//go:noescape
func CallOnConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnConnect
//go:noescape
func TryOnConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffConnect
//go:noescape
func HasFuncOffConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffConnect
//go:noescape
func FuncOffConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffConnect
//go:noescape
func CallOffConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffConnect
//go:noescape
func TryOffConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnConnect
//go:noescape
func HasFuncHasOnConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnConnect
//go:noescape
func FuncHasOnConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnConnect
//go:noescape
func CallHasOnConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnConnect
//go:noescape
func TryHasOnConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnConnectExternal
//go:noescape
func HasFuncOnConnectExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnConnectExternal
//go:noescape
func FuncOnConnectExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnConnectExternal
//go:noescape
func CallOnConnectExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnConnectExternal
//go:noescape
func TryOnConnectExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffConnectExternal
//go:noescape
func HasFuncOffConnectExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffConnectExternal
//go:noescape
func FuncOffConnectExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffConnectExternal
//go:noescape
func CallOffConnectExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffConnectExternal
//go:noescape
func TryOffConnectExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnConnectExternal
//go:noescape
func HasFuncHasOnConnectExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnConnectExternal
//go:noescape
func FuncHasOnConnectExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnConnectExternal
//go:noescape
func CallHasOnConnectExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnConnectExternal
//go:noescape
func TryHasOnConnectExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnConnectNative
//go:noescape
func HasFuncOnConnectNative() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnConnectNative
//go:noescape
func FuncOnConnectNative(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnConnectNative
//go:noescape
func CallOnConnectNative(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnConnectNative
//go:noescape
func TryOnConnectNative(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffConnectNative
//go:noescape
func HasFuncOffConnectNative() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffConnectNative
//go:noescape
func FuncOffConnectNative(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffConnectNative
//go:noescape
func CallOffConnectNative(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffConnectNative
//go:noescape
func TryOffConnectNative(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnConnectNative
//go:noescape
func HasFuncHasOnConnectNative() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnConnectNative
//go:noescape
func FuncHasOnConnectNative(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnConnectNative
//go:noescape
func CallHasOnConnectNative(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnConnectNative
//go:noescape
func TryHasOnConnectNative(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnInstalled
//go:noescape
func HasFuncOnInstalled() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnInstalled
//go:noescape
func FuncOnInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnInstalled
//go:noescape
func CallOnInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnInstalled
//go:noescape
func TryOnInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffInstalled
//go:noescape
func HasFuncOffInstalled() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffInstalled
//go:noescape
func FuncOffInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffInstalled
//go:noescape
func CallOffInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffInstalled
//go:noescape
func TryOffInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnInstalled
//go:noescape
func HasFuncHasOnInstalled() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnInstalled
//go:noescape
func FuncHasOnInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnInstalled
//go:noescape
func CallHasOnInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnInstalled
//go:noescape
func TryHasOnInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnMessage
//go:noescape
func HasFuncOnMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnMessage
//go:noescape
func FuncOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnMessage
//go:noescape
func CallOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnMessage
//go:noescape
func TryOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffMessage
//go:noescape
func HasFuncOffMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffMessage
//go:noescape
func FuncOffMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffMessage
//go:noescape
func CallOffMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffMessage
//go:noescape
func TryOffMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnMessage
//go:noescape
func HasFuncHasOnMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnMessage
//go:noescape
func FuncHasOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnMessage
//go:noescape
func CallHasOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnMessage
//go:noescape
func TryHasOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnMessageExternal
//go:noescape
func HasFuncOnMessageExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnMessageExternal
//go:noescape
func FuncOnMessageExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnMessageExternal
//go:noescape
func CallOnMessageExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnMessageExternal
//go:noescape
func TryOnMessageExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffMessageExternal
//go:noescape
func HasFuncOffMessageExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffMessageExternal
//go:noescape
func FuncOffMessageExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffMessageExternal
//go:noescape
func CallOffMessageExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffMessageExternal
//go:noescape
func TryOffMessageExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnMessageExternal
//go:noescape
func HasFuncHasOnMessageExternal() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnMessageExternal
//go:noescape
func FuncHasOnMessageExternal(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnMessageExternal
//go:noescape
func CallHasOnMessageExternal(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnMessageExternal
//go:noescape
func TryHasOnMessageExternal(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnRestartRequired
//go:noescape
func HasFuncOnRestartRequired() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnRestartRequired
//go:noescape
func FuncOnRestartRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnRestartRequired
//go:noescape
func CallOnRestartRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnRestartRequired
//go:noescape
func TryOnRestartRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffRestartRequired
//go:noescape
func HasFuncOffRestartRequired() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffRestartRequired
//go:noescape
func FuncOffRestartRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffRestartRequired
//go:noescape
func CallOffRestartRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffRestartRequired
//go:noescape
func TryOffRestartRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnRestartRequired
//go:noescape
func HasFuncHasOnRestartRequired() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnRestartRequired
//go:noescape
func FuncHasOnRestartRequired(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnRestartRequired
//go:noescape
func CallHasOnRestartRequired(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnRestartRequired
//go:noescape
func TryHasOnRestartRequired(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnStartup
//go:noescape
func HasFuncOnStartup() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnStartup
//go:noescape
func FuncOnStartup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnStartup
//go:noescape
func CallOnStartup(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnStartup
//go:noescape
func TryOnStartup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffStartup
//go:noescape
func HasFuncOffStartup() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffStartup
//go:noescape
func FuncOffStartup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffStartup
//go:noescape
func CallOffStartup(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffStartup
//go:noescape
func TryOffStartup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnStartup
//go:noescape
func HasFuncHasOnStartup() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnStartup
//go:noescape
func FuncHasOnStartup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnStartup
//go:noescape
func CallHasOnStartup(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnStartup
//go:noescape
func TryHasOnStartup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnSuspend
//go:noescape
func HasFuncOnSuspend() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnSuspend
//go:noescape
func FuncOnSuspend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnSuspend
//go:noescape
func CallOnSuspend(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnSuspend
//go:noescape
func TryOnSuspend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffSuspend
//go:noescape
func HasFuncOffSuspend() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffSuspend
//go:noescape
func FuncOffSuspend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffSuspend
//go:noescape
func CallOffSuspend(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffSuspend
//go:noescape
func TryOffSuspend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnSuspend
//go:noescape
func HasFuncHasOnSuspend() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnSuspend
//go:noescape
func FuncHasOnSuspend(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnSuspend
//go:noescape
func CallHasOnSuspend(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnSuspend
//go:noescape
func TryHasOnSuspend(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnSuspendCanceled
//go:noescape
func HasFuncOnSuspendCanceled() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnSuspendCanceled
//go:noescape
func FuncOnSuspendCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnSuspendCanceled
//go:noescape
func CallOnSuspendCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnSuspendCanceled
//go:noescape
func TryOnSuspendCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffSuspendCanceled
//go:noescape
func HasFuncOffSuspendCanceled() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffSuspendCanceled
//go:noescape
func FuncOffSuspendCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffSuspendCanceled
//go:noescape
func CallOffSuspendCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffSuspendCanceled
//go:noescape
func TryOffSuspendCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnSuspendCanceled
//go:noescape
func HasFuncHasOnSuspendCanceled() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnSuspendCanceled
//go:noescape
func FuncHasOnSuspendCanceled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnSuspendCanceled
//go:noescape
func CallHasOnSuspendCanceled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnSuspendCanceled
//go:noescape
func TryHasOnSuspendCanceled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnUpdateAvailable
//go:noescape
func HasFuncOnUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnUpdateAvailable
//go:noescape
func FuncOnUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnUpdateAvailable
//go:noescape
func CallOnUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnUpdateAvailable
//go:noescape
func TryOnUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffUpdateAvailable
//go:noescape
func HasFuncOffUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffUpdateAvailable
//go:noescape
func FuncOffUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffUpdateAvailable
//go:noescape
func CallOffUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffUpdateAvailable
//go:noescape
func TryOffUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnUpdateAvailable
//go:noescape
func HasFuncHasOnUpdateAvailable() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnUpdateAvailable
//go:noescape
func FuncHasOnUpdateAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnUpdateAvailable
//go:noescape
func CallHasOnUpdateAvailable(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnUpdateAvailable
//go:noescape
func TryHasOnUpdateAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnUserScriptConnect
//go:noescape
func HasFuncOnUserScriptConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnUserScriptConnect
//go:noescape
func FuncOnUserScriptConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnUserScriptConnect
//go:noescape
func CallOnUserScriptConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnUserScriptConnect
//go:noescape
func TryOnUserScriptConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffUserScriptConnect
//go:noescape
func HasFuncOffUserScriptConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffUserScriptConnect
//go:noescape
func FuncOffUserScriptConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffUserScriptConnect
//go:noescape
func CallOffUserScriptConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffUserScriptConnect
//go:noescape
func TryOffUserScriptConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnUserScriptConnect
//go:noescape
func HasFuncHasOnUserScriptConnect() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnUserScriptConnect
//go:noescape
func FuncHasOnUserScriptConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnUserScriptConnect
//go:noescape
func CallHasOnUserScriptConnect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnUserScriptConnect
//go:noescape
func TryHasOnUserScriptConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OnUserScriptMessage
//go:noescape
func HasFuncOnUserScriptMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_OnUserScriptMessage
//go:noescape
func FuncOnUserScriptMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OnUserScriptMessage
//go:noescape
func CallOnUserScriptMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OnUserScriptMessage
//go:noescape
func TryOnUserScriptMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OffUserScriptMessage
//go:noescape
func HasFuncOffUserScriptMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_OffUserScriptMessage
//go:noescape
func FuncOffUserScriptMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OffUserScriptMessage
//go:noescape
func CallOffUserScriptMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_OffUserScriptMessage
//go:noescape
func TryOffUserScriptMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_HasOnUserScriptMessage
//go:noescape
func HasFuncHasOnUserScriptMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_HasOnUserScriptMessage
//go:noescape
func FuncHasOnUserScriptMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_HasOnUserScriptMessage
//go:noescape
func CallHasOnUserScriptMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/runtime try_HasOnUserScriptMessage
//go:noescape
func TryHasOnUserScriptMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_OpenOptionsPage
//go:noescape
func HasFuncOpenOptionsPage() js.Ref

//go:wasmimport plat/js/webext/runtime func_OpenOptionsPage
//go:noescape
func FuncOpenOptionsPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_OpenOptionsPage
//go:noescape
func CallOpenOptionsPage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_OpenOptionsPage
//go:noescape
func TryOpenOptionsPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_Reload
//go:noescape
func HasFuncReload() js.Ref

//go:wasmimport plat/js/webext/runtime func_Reload
//go:noescape
func FuncReload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_Reload
//go:noescape
func CallReload(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_Reload
//go:noescape
func TryReload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_RequestUpdateCheck
//go:noescape
func HasFuncRequestUpdateCheck() js.Ref

//go:wasmimport plat/js/webext/runtime func_RequestUpdateCheck
//go:noescape
func FuncRequestUpdateCheck(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_RequestUpdateCheck
//go:noescape
func CallRequestUpdateCheck(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_RequestUpdateCheck
//go:noescape
func TryRequestUpdateCheck(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_Restart
//go:noescape
func HasFuncRestart() js.Ref

//go:wasmimport plat/js/webext/runtime func_Restart
//go:noescape
func FuncRestart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_Restart
//go:noescape
func CallRestart(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_Restart
//go:noescape
func TryRestart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_RestartAfterDelay
//go:noescape
func HasFuncRestartAfterDelay() js.Ref

//go:wasmimport plat/js/webext/runtime func_RestartAfterDelay
//go:noescape
func FuncRestartAfterDelay(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_RestartAfterDelay
//go:noescape
func CallRestartAfterDelay(
	retPtr unsafe.Pointer,
	seconds float64)

//go:wasmimport plat/js/webext/runtime try_RestartAfterDelay
//go:noescape
func TryRestartAfterDelay(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	seconds float64) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_SendMessage
//go:noescape
func HasFuncSendMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_SendMessage
//go:noescape
func FuncSendMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_SendMessage
//go:noescape
func CallSendMessage(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime try_SendMessage
//go:noescape
func TrySendMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_SendNativeMessage
//go:noescape
func HasFuncSendNativeMessage() js.Ref

//go:wasmimport plat/js/webext/runtime func_SendNativeMessage
//go:noescape
func FuncSendNativeMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_SendNativeMessage
//go:noescape
func CallSendNativeMessage(
	retPtr unsafe.Pointer,
	application js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/runtime try_SendNativeMessage
//go:noescape
func TrySendNativeMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	application js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/runtime has_SetUninstallURL
//go:noescape
func HasFuncSetUninstallURL() js.Ref

//go:wasmimport plat/js/webext/runtime func_SetUninstallURL
//go:noescape
func FuncSetUninstallURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/runtime call_SetUninstallURL
//go:noescape
func CallSetUninstallURL(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/runtime try_SetUninstallURL
//go:noescape
func TrySetUninstallURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

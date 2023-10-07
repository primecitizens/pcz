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

//go:wasmimport plat/js/webext/webviewtag store_AttachArgWebview
//go:noescape
func AttachArgWebviewJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_AttachArgWebview
//go:noescape
func AttachArgWebviewJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_ClearDataOptions
//go:noescape
func ClearDataOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_ClearDataOptions
//go:noescape
func ClearDataOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_ClearDataTypeSet
//go:noescape
func ClearDataTypeSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_ClearDataTypeSet
//go:noescape
func ClearDataTypeSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_InjectionItems
//go:noescape
func InjectionItemsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_InjectionItems
//go:noescape
func InjectionItemsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_ContentScriptDetails
//go:noescape
func ContentScriptDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_ContentScriptDetails
//go:noescape
func ContentScriptDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_ContentWindowType_PostMessage
//go:noescape
func HasFuncContentWindowTypePostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContentWindowType_PostMessage
//go:noescape
func FuncContentWindowTypePostMessage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContentWindowType_PostMessage
//go:noescape
func CallContentWindowTypePostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContentWindowType_PostMessage
//go:noescape
func TryContentWindowTypePostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_ContextType
//go:noescape
func ConstOfContextType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag store_ContextMenuCreateProperties
//go:noescape
func ContextMenuCreatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_ContextMenuCreateProperties
//go:noescape
func ContextMenuCreatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_ContextMenuUpdateProperties
//go:noescape
func ContextMenuUpdatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_ContextMenuUpdateProperties
//go:noescape
func ContextMenuUpdatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Create
//go:noescape
func HasFuncContextMenusTypeCreate(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Create
//go:noescape
func FuncContextMenusTypeCreate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Create
//go:noescape
func CallContextMenusTypeCreate(
	this js.Ref, retPtr unsafe.Pointer,
	createProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Create
//go:noescape
func TryContextMenusTypeCreate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	createProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Create1
//go:noescape
func HasFuncContextMenusTypeCreate1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Create1
//go:noescape
func FuncContextMenusTypeCreate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Create1
//go:noescape
func CallContextMenusTypeCreate1(
	this js.Ref, retPtr unsafe.Pointer,
	createProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Create1
//go:noescape
func TryContextMenusTypeCreate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	createProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Update
//go:noescape
func HasFuncContextMenusTypeUpdate(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Update
//go:noescape
func FuncContextMenusTypeUpdate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Update
//go:noescape
func CallContextMenusTypeUpdate(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Update
//go:noescape
func TryContextMenusTypeUpdate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Update1
//go:noescape
func HasFuncContextMenusTypeUpdate1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Update1
//go:noescape
func FuncContextMenusTypeUpdate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Update1
//go:noescape
func CallContextMenusTypeUpdate1(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Update1
//go:noescape
func TryContextMenusTypeUpdate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	updateProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Remove
//go:noescape
func HasFuncContextMenusTypeRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Remove
//go:noescape
func FuncContextMenusTypeRemove(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Remove
//go:noescape
func CallContextMenusTypeRemove(
	this js.Ref, retPtr unsafe.Pointer,
	menuItemId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Remove
//go:noescape
func TryContextMenusTypeRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	menuItemId js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_Remove1
//go:noescape
func HasFuncContextMenusTypeRemove1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_Remove1
//go:noescape
func FuncContextMenusTypeRemove1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_Remove1
//go:noescape
func CallContextMenusTypeRemove1(
	this js.Ref, retPtr unsafe.Pointer,
	menuItemId js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_Remove1
//go:noescape
func TryContextMenusTypeRemove1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	menuItemId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_RemoveAll
//go:noescape
func HasFuncContextMenusTypeRemoveAll(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_RemoveAll
//go:noescape
func FuncContextMenusTypeRemoveAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_RemoveAll
//go:noescape
func CallContextMenusTypeRemoveAll(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_RemoveAll
//go:noescape
func TryContextMenusTypeRemoveAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ContextMenusType_RemoveAll1
//go:noescape
func HasFuncContextMenusTypeRemoveAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ContextMenusType_RemoveAll1
//go:noescape
func FuncContextMenusTypeRemoveAll1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ContextMenusType_RemoveAll1
//go:noescape
func CallContextMenusTypeRemoveAll1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_ContextMenusType_RemoveAll1
//go:noescape
func TryContextMenusTypeRemoveAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_DialogArgMessageType
//go:noescape
func ConstOfDialogArgMessageType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag has_DialogController_Ok
//go:noescape
func HasFuncDialogControllerOk(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_DialogController_Ok
//go:noescape
func FuncDialogControllerOk(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_DialogController_Ok
//go:noescape
func CallDialogControllerOk(
	this js.Ref, retPtr unsafe.Pointer,
	response js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_DialogController_Ok
//go:noescape
func TryDialogControllerOk(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	response js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_DialogController_Ok1
//go:noescape
func HasFuncDialogControllerOk1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_DialogController_Ok1
//go:noescape
func FuncDialogControllerOk1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_DialogController_Ok1
//go:noescape
func CallDialogControllerOk1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_DialogController_Ok1
//go:noescape
func TryDialogControllerOk1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_DialogController_Cancel
//go:noescape
func HasFuncDialogControllerCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_DialogController_Cancel
//go:noescape
func FuncDialogControllerCancel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_DialogController_Cancel
//go:noescape
func CallDialogControllerCancel(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_DialogController_Cancel
//go:noescape
func TryDialogControllerCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag get_DownloadPermissionRequest_RequestMethod
//go:noescape
func GetDownloadPermissionRequestRequestMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_DownloadPermissionRequest_RequestMethod
//go:noescape
func SetDownloadPermissionRequestRequestMethod(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_DownloadPermissionRequest_Url
//go:noescape
func GetDownloadPermissionRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_DownloadPermissionRequest_Url
//go:noescape
func SetDownloadPermissionRequestUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_DownloadPermissionRequest_Allow
//go:noescape
func HasFuncDownloadPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_DownloadPermissionRequest_Allow
//go:noescape
func FuncDownloadPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_DownloadPermissionRequest_Allow
//go:noescape
func CallDownloadPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_DownloadPermissionRequest_Allow
//go:noescape
func TryDownloadPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_DownloadPermissionRequest_Deny
//go:noescape
func HasFuncDownloadPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_DownloadPermissionRequest_Deny
//go:noescape
func FuncDownloadPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_DownloadPermissionRequest_Deny
//go:noescape
func CallDownloadPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_DownloadPermissionRequest_Deny
//go:noescape
func TryDownloadPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_ExitArgReason
//go:noescape
func ConstOfExitArgReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag get_FileSystemPermissionRequest_Url
//go:noescape
func GetFileSystemPermissionRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_FileSystemPermissionRequest_Url
//go:noescape
func SetFileSystemPermissionRequestUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_FileSystemPermissionRequest_Allow
//go:noescape
func HasFuncFileSystemPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_FileSystemPermissionRequest_Allow
//go:noescape
func FuncFileSystemPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_FileSystemPermissionRequest_Allow
//go:noescape
func CallFileSystemPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_FileSystemPermissionRequest_Allow
//go:noescape
func TryFileSystemPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_FileSystemPermissionRequest_Deny
//go:noescape
func HasFuncFileSystemPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_FileSystemPermissionRequest_Deny
//go:noescape
func FuncFileSystemPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_FileSystemPermissionRequest_Deny
//go:noescape
func CallFileSystemPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_FileSystemPermissionRequest_Deny
//go:noescape
func TryFileSystemPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag store_SelectionRect
//go:noescape
func SelectionRectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_SelectionRect
//go:noescape
func SelectionRectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_FindCallbackResults
//go:noescape
func FindCallbackResultsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_FindCallbackResults
//go:noescape
func FindCallbackResultsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag store_FindOptions
//go:noescape
func FindOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_FindOptions
//go:noescape
func FindOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_FullscreenPermissionRequest_Origin
//go:noescape
func GetFullscreenPermissionRequestOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_FullscreenPermissionRequest_Origin
//go:noescape
func SetFullscreenPermissionRequestOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_FullscreenPermissionRequest_Allow
//go:noescape
func HasFuncFullscreenPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_FullscreenPermissionRequest_Allow
//go:noescape
func FuncFullscreenPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_FullscreenPermissionRequest_Allow
//go:noescape
func CallFullscreenPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_FullscreenPermissionRequest_Allow
//go:noescape
func TryFullscreenPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_FullscreenPermissionRequest_Deny
//go:noescape
func HasFuncFullscreenPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_FullscreenPermissionRequest_Deny
//go:noescape
func FuncFullscreenPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_FullscreenPermissionRequest_Deny
//go:noescape
func CallFullscreenPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_FullscreenPermissionRequest_Deny
//go:noescape
func TryFullscreenPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag get_GeolocationPermissionRequest_Url
//go:noescape
func GetGeolocationPermissionRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_GeolocationPermissionRequest_Url
//go:noescape
func SetGeolocationPermissionRequestUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_GeolocationPermissionRequest_Allow
//go:noescape
func HasFuncGeolocationPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GeolocationPermissionRequest_Allow
//go:noescape
func FuncGeolocationPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GeolocationPermissionRequest_Allow
//go:noescape
func CallGeolocationPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_GeolocationPermissionRequest_Allow
//go:noescape
func TryGeolocationPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GeolocationPermissionRequest_Deny
//go:noescape
func HasFuncGeolocationPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GeolocationPermissionRequest_Deny
//go:noescape
func FuncGeolocationPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GeolocationPermissionRequest_Deny
//go:noescape
func CallGeolocationPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_GeolocationPermissionRequest_Deny
//go:noescape
func TryGeolocationPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_ZoomMode
//go:noescape
func ConstOfZoomMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag store_InjectDetails
//go:noescape
func InjectDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_InjectDetails
//go:noescape
func InjectDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_LoadPluginPermissionRequest_Identifier
//go:noescape
func GetLoadPluginPermissionRequestIdentifier(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_LoadPluginPermissionRequest_Identifier
//go:noescape
func SetLoadPluginPermissionRequestIdentifier(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_LoadPluginPermissionRequest_Name
//go:noescape
func GetLoadPluginPermissionRequestName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_LoadPluginPermissionRequest_Name
//go:noescape
func SetLoadPluginPermissionRequestName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_LoadPluginPermissionRequest_Allow
//go:noescape
func HasFuncLoadPluginPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_LoadPluginPermissionRequest_Allow
//go:noescape
func FuncLoadPluginPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_LoadPluginPermissionRequest_Allow
//go:noescape
func CallLoadPluginPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_LoadPluginPermissionRequest_Allow
//go:noescape
func TryLoadPluginPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_LoadPluginPermissionRequest_Deny
//go:noescape
func HasFuncLoadPluginPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_LoadPluginPermissionRequest_Deny
//go:noescape
func FuncLoadPluginPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_LoadPluginPermissionRequest_Deny
//go:noescape
func CallLoadPluginPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_LoadPluginPermissionRequest_Deny
//go:noescape
func TryLoadPluginPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_LoadabortArgReason
//go:noescape
func ConstOfLoadabortArgReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag get_MediaPermissionRequest_Url
//go:noescape
func GetMediaPermissionRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_MediaPermissionRequest_Url
//go:noescape
func SetMediaPermissionRequestUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_MediaPermissionRequest_Allow
//go:noescape
func HasFuncMediaPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_MediaPermissionRequest_Allow
//go:noescape
func FuncMediaPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_MediaPermissionRequest_Allow
//go:noescape
func CallMediaPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_MediaPermissionRequest_Allow
//go:noescape
func TryMediaPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_MediaPermissionRequest_Deny
//go:noescape
func HasFuncMediaPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_MediaPermissionRequest_Deny
//go:noescape
func FuncMediaPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_MediaPermissionRequest_Deny
//go:noescape
func CallMediaPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_MediaPermissionRequest_Deny
//go:noescape
func TryMediaPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_NewWindow_Attach
//go:noescape
func HasFuncNewWindowAttach(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_NewWindow_Attach
//go:noescape
func FuncNewWindowAttach(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_NewWindow_Attach
//go:noescape
func CallNewWindowAttach(
	this js.Ref, retPtr unsafe.Pointer,
	webview unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_NewWindow_Attach
//go:noescape
func TryNewWindowAttach(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webview unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_NewWindow_Discard
//go:noescape
func HasFuncNewWindowDiscard(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_NewWindow_Discard
//go:noescape
func FuncNewWindowDiscard(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_NewWindow_Discard
//go:noescape
func CallNewWindowDiscard(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_NewWindow_Discard
//go:noescape
func TryNewWindowDiscard(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_NewwindowArgWindowOpenDisposition
//go:noescape
func ConstOfNewwindowArgWindowOpenDisposition(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag constof_PermissionrequestArgPermission
//go:noescape
func ConstOfPermissionrequestArgPermission(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag store_PermissionrequestArgRequest
//go:noescape
func PermissionrequestArgRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_PermissionrequestArgRequest
//go:noescape
func PermissionrequestArgRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_PointerLockPermissionRequest_LastUnlockedBySelf
//go:noescape
func GetPointerLockPermissionRequestLastUnlockedBySelf(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_PointerLockPermissionRequest_LastUnlockedBySelf
//go:noescape
func SetPointerLockPermissionRequestLastUnlockedBySelf(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_PointerLockPermissionRequest_Url
//go:noescape
func GetPointerLockPermissionRequestUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_PointerLockPermissionRequest_Url
//go:noescape
func SetPointerLockPermissionRequestUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag get_PointerLockPermissionRequest_UserGesture
//go:noescape
func GetPointerLockPermissionRequestUserGesture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag set_PointerLockPermissionRequest_UserGesture
//go:noescape
func SetPointerLockPermissionRequestUserGesture(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_PointerLockPermissionRequest_Allow
//go:noescape
func HasFuncPointerLockPermissionRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_PointerLockPermissionRequest_Allow
//go:noescape
func FuncPointerLockPermissionRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_PointerLockPermissionRequest_Allow
//go:noescape
func CallPointerLockPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_PointerLockPermissionRequest_Allow
//go:noescape
func TryPointerLockPermissionRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_PointerLockPermissionRequest_Deny
//go:noescape
func HasFuncPointerLockPermissionRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag func_PointerLockPermissionRequest_Deny
//go:noescape
func FuncPointerLockPermissionRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_PointerLockPermissionRequest_Deny
//go:noescape
func CallPointerLockPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_PointerLockPermissionRequest_Deny
//go:noescape
func TryPointerLockPermissionRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag constof_StopFindingArgAction
//go:noescape
func ConstOfStopFindingArgAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewtag store_WebRequestEventInterface
//go:noescape
func WebRequestEventInterfaceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewtag load_WebRequestEventInterface
//go:noescape
func WebRequestEventInterfaceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_AddContentScripts
//go:noescape
func HasFuncAddContentScripts() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_AddContentScripts
//go:noescape
func FuncAddContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_AddContentScripts
//go:noescape
func CallAddContentScripts(
	retPtr unsafe.Pointer,
	contentScriptList js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_AddContentScripts
//go:noescape
func TryAddContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contentScriptList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Back
//go:noescape
func HasFuncBack() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Back
//go:noescape
func FuncBack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Back
//go:noescape
func CallBack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_Back
//go:noescape
func TryBack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_CanGoBack
//go:noescape
func HasFuncCanGoBack() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_CanGoBack
//go:noescape
func FuncCanGoBack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_CanGoBack
//go:noescape
func CallCanGoBack(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_CanGoBack
//go:noescape
func TryCanGoBack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_CanGoForward
//go:noescape
func HasFuncCanGoForward() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_CanGoForward
//go:noescape
func FuncCanGoForward(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_CanGoForward
//go:noescape
func CallCanGoForward(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_CanGoForward
//go:noescape
func TryCanGoForward(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_CaptureVisibleRegion
//go:noescape
func HasFuncCaptureVisibleRegion() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_CaptureVisibleRegion
//go:noescape
func FuncCaptureVisibleRegion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_CaptureVisibleRegion
//go:noescape
func CallCaptureVisibleRegion(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_CaptureVisibleRegion
//go:noescape
func TryCaptureVisibleRegion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ClearData
//go:noescape
func HasFuncClearData() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ClearData
//go:noescape
func FuncClearData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ClearData
//go:noescape
func CallClearData(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	types unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ClearData
//go:noescape
func TryClearData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	types unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnClose
//go:noescape
func HasFuncOnClose() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnClose
//go:noescape
func FuncOnClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnClose
//go:noescape
func CallOnClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnClose
//go:noescape
func TryOnClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffClose
//go:noescape
func HasFuncOffClose() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffClose
//go:noescape
func FuncOffClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffClose
//go:noescape
func CallOffClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffClose
//go:noescape
func TryOffClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnClose
//go:noescape
func HasFuncHasOnClose() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnClose
//go:noescape
func FuncHasOnClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnClose
//go:noescape
func CallHasOnClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnClose
//go:noescape
func TryHasOnClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnConsolemessage
//go:noescape
func HasFuncOnConsolemessage() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnConsolemessage
//go:noescape
func FuncOnConsolemessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnConsolemessage
//go:noescape
func CallOnConsolemessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnConsolemessage
//go:noescape
func TryOnConsolemessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffConsolemessage
//go:noescape
func HasFuncOffConsolemessage() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffConsolemessage
//go:noescape
func FuncOffConsolemessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffConsolemessage
//go:noescape
func CallOffConsolemessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffConsolemessage
//go:noescape
func TryOffConsolemessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnConsolemessage
//go:noescape
func HasFuncHasOnConsolemessage() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnConsolemessage
//go:noescape
func FuncHasOnConsolemessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnConsolemessage
//go:noescape
func CallHasOnConsolemessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnConsolemessage
//go:noescape
func TryHasOnConsolemessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag get_ContentWindow
//go:noescape
func GetContentWindow(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webviewtag set_ContentWindow
//go:noescape
func SetContentWindow(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_OnContentload
//go:noescape
func HasFuncOnContentload() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnContentload
//go:noescape
func FuncOnContentload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnContentload
//go:noescape
func CallOnContentload(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnContentload
//go:noescape
func TryOnContentload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffContentload
//go:noescape
func HasFuncOffContentload() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffContentload
//go:noescape
func FuncOffContentload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffContentload
//go:noescape
func CallOffContentload(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffContentload
//go:noescape
func TryOffContentload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnContentload
//go:noescape
func HasFuncHasOnContentload() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnContentload
//go:noescape
func FuncHasOnContentload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnContentload
//go:noescape
func CallHasOnContentload(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnContentload
//go:noescape
func TryHasOnContentload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag get_ContextMenus
//go:noescape
func GetContextMenus(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webviewtag set_ContextMenus
//go:noescape
func SetContextMenus(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_OnDialog
//go:noescape
func HasFuncOnDialog() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnDialog
//go:noescape
func FuncOnDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnDialog
//go:noescape
func CallOnDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnDialog
//go:noescape
func TryOnDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffDialog
//go:noescape
func HasFuncOffDialog() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffDialog
//go:noescape
func FuncOffDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffDialog
//go:noescape
func CallOffDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffDialog
//go:noescape
func TryOffDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnDialog
//go:noescape
func HasFuncHasOnDialog() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnDialog
//go:noescape
func FuncHasOnDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnDialog
//go:noescape
func CallHasOnDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnDialog
//go:noescape
func TryHasOnDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_ExecuteScript
//go:noescape
func HasFuncExecuteScript() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_ExecuteScript
//go:noescape
func FuncExecuteScript(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_ExecuteScript
//go:noescape
func CallExecuteScript(
	retPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_ExecuteScript
//go:noescape
func TryExecuteScript(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnExit
//go:noescape
func HasFuncOnExit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnExit
//go:noescape
func FuncOnExit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnExit
//go:noescape
func CallOnExit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnExit
//go:noescape
func TryOnExit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffExit
//go:noescape
func HasFuncOffExit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffExit
//go:noescape
func FuncOffExit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffExit
//go:noescape
func CallOffExit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffExit
//go:noescape
func TryOffExit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnExit
//go:noescape
func HasFuncHasOnExit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnExit
//go:noescape
func FuncHasOnExit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnExit
//go:noescape
func CallHasOnExit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnExit
//go:noescape
func TryHasOnExit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Find
//go:noescape
func HasFuncFind() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Find
//go:noescape
func FuncFind(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Find
//go:noescape
func CallFind(
	retPtr unsafe.Pointer,
	searchText js.Ref,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_Find
//go:noescape
func TryFind(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchText js.Ref,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnFindupdate
//go:noescape
func HasFuncOnFindupdate() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnFindupdate
//go:noescape
func FuncOnFindupdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnFindupdate
//go:noescape
func CallOnFindupdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnFindupdate
//go:noescape
func TryOnFindupdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffFindupdate
//go:noescape
func HasFuncOffFindupdate() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffFindupdate
//go:noescape
func FuncOffFindupdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffFindupdate
//go:noescape
func CallOffFindupdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffFindupdate
//go:noescape
func TryOffFindupdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnFindupdate
//go:noescape
func HasFuncHasOnFindupdate() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnFindupdate
//go:noescape
func FuncHasOnFindupdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnFindupdate
//go:noescape
func CallHasOnFindupdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnFindupdate
//go:noescape
func TryHasOnFindupdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Forward
//go:noescape
func HasFuncForward() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Forward
//go:noescape
func FuncForward(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Forward
//go:noescape
func CallForward(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_Forward
//go:noescape
func TryForward(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GetAudioState
//go:noescape
func HasFuncGetAudioState() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GetAudioState
//go:noescape
func FuncGetAudioState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GetAudioState
//go:noescape
func CallGetAudioState(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_GetAudioState
//go:noescape
func TryGetAudioState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GetProcessId
//go:noescape
func HasFuncGetProcessId() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GetProcessId
//go:noescape
func FuncGetProcessId(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GetProcessId
//go:noescape
func CallGetProcessId(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_GetProcessId
//go:noescape
func TryGetProcessId(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GetUserAgent
//go:noescape
func HasFuncGetUserAgent() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GetUserAgent
//go:noescape
func FuncGetUserAgent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GetUserAgent
//go:noescape
func CallGetUserAgent(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_GetUserAgent
//go:noescape
func TryGetUserAgent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GetZoom
//go:noescape
func HasFuncGetZoom() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GetZoom
//go:noescape
func FuncGetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GetZoom
//go:noescape
func CallGetZoom(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_GetZoom
//go:noescape
func TryGetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_GetZoomMode
//go:noescape
func HasFuncGetZoomMode() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_GetZoomMode
//go:noescape
func FuncGetZoomMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_GetZoomMode
//go:noescape
func CallGetZoomMode(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_GetZoomMode
//go:noescape
func TryGetZoomMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Go
//go:noescape
func HasFuncGo() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Go
//go:noescape
func FuncGo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Go
//go:noescape
func CallGo(
	retPtr unsafe.Pointer,
	relativeIndex float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_Go
//go:noescape
func TryGo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	relativeIndex float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_InsertCSS
//go:noescape
func HasFuncInsertCSS() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_InsertCSS
//go:noescape
func FuncInsertCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_InsertCSS
//go:noescape
func CallInsertCSS(
	retPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_InsertCSS
//go:noescape
func TryInsertCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_IsAudioMuted
//go:noescape
func HasFuncIsAudioMuted() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_IsAudioMuted
//go:noescape
func FuncIsAudioMuted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_IsAudioMuted
//go:noescape
func CallIsAudioMuted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_IsAudioMuted
//go:noescape
func TryIsAudioMuted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_IsSpatialNavigationEnabled
//go:noescape
func HasFuncIsSpatialNavigationEnabled() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_IsSpatialNavigationEnabled
//go:noescape
func FuncIsSpatialNavigationEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_IsSpatialNavigationEnabled
//go:noescape
func CallIsSpatialNavigationEnabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_IsSpatialNavigationEnabled
//go:noescape
func TryIsSpatialNavigationEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_IsUserAgentOverridden
//go:noescape
func HasFuncIsUserAgentOverridden() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_IsUserAgentOverridden
//go:noescape
func FuncIsUserAgentOverridden(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_IsUserAgentOverridden
//go:noescape
func CallIsUserAgentOverridden(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_IsUserAgentOverridden
//go:noescape
func TryIsUserAgentOverridden(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_LoadDataWithBaseUrl
//go:noescape
func HasFuncLoadDataWithBaseUrl() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_LoadDataWithBaseUrl
//go:noescape
func FuncLoadDataWithBaseUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_LoadDataWithBaseUrl
//go:noescape
func CallLoadDataWithBaseUrl(
	retPtr unsafe.Pointer,
	dataUrl js.Ref,
	baseUrl js.Ref,
	virtualUrl js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_LoadDataWithBaseUrl
//go:noescape
func TryLoadDataWithBaseUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataUrl js.Ref,
	baseUrl js.Ref,
	virtualUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnLoadabort
//go:noescape
func HasFuncOnLoadabort() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnLoadabort
//go:noescape
func FuncOnLoadabort(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnLoadabort
//go:noescape
func CallOnLoadabort(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnLoadabort
//go:noescape
func TryOnLoadabort(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffLoadabort
//go:noescape
func HasFuncOffLoadabort() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffLoadabort
//go:noescape
func FuncOffLoadabort(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffLoadabort
//go:noescape
func CallOffLoadabort(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffLoadabort
//go:noescape
func TryOffLoadabort(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnLoadabort
//go:noescape
func HasFuncHasOnLoadabort() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnLoadabort
//go:noescape
func FuncHasOnLoadabort(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnLoadabort
//go:noescape
func CallHasOnLoadabort(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnLoadabort
//go:noescape
func TryHasOnLoadabort(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnLoadcommit
//go:noescape
func HasFuncOnLoadcommit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnLoadcommit
//go:noescape
func FuncOnLoadcommit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnLoadcommit
//go:noescape
func CallOnLoadcommit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnLoadcommit
//go:noescape
func TryOnLoadcommit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffLoadcommit
//go:noescape
func HasFuncOffLoadcommit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffLoadcommit
//go:noescape
func FuncOffLoadcommit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffLoadcommit
//go:noescape
func CallOffLoadcommit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffLoadcommit
//go:noescape
func TryOffLoadcommit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnLoadcommit
//go:noescape
func HasFuncHasOnLoadcommit() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnLoadcommit
//go:noescape
func FuncHasOnLoadcommit(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnLoadcommit
//go:noescape
func CallHasOnLoadcommit(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnLoadcommit
//go:noescape
func TryHasOnLoadcommit(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnLoadredirect
//go:noescape
func HasFuncOnLoadredirect() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnLoadredirect
//go:noescape
func FuncOnLoadredirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnLoadredirect
//go:noescape
func CallOnLoadredirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnLoadredirect
//go:noescape
func TryOnLoadredirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffLoadredirect
//go:noescape
func HasFuncOffLoadredirect() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffLoadredirect
//go:noescape
func FuncOffLoadredirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffLoadredirect
//go:noescape
func CallOffLoadredirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffLoadredirect
//go:noescape
func TryOffLoadredirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnLoadredirect
//go:noescape
func HasFuncHasOnLoadredirect() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnLoadredirect
//go:noescape
func FuncHasOnLoadredirect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnLoadredirect
//go:noescape
func CallHasOnLoadredirect(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnLoadredirect
//go:noescape
func TryHasOnLoadredirect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnLoadstart
//go:noescape
func HasFuncOnLoadstart() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnLoadstart
//go:noescape
func FuncOnLoadstart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnLoadstart
//go:noescape
func CallOnLoadstart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnLoadstart
//go:noescape
func TryOnLoadstart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffLoadstart
//go:noescape
func HasFuncOffLoadstart() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffLoadstart
//go:noescape
func FuncOffLoadstart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffLoadstart
//go:noescape
func CallOffLoadstart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffLoadstart
//go:noescape
func TryOffLoadstart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnLoadstart
//go:noescape
func HasFuncHasOnLoadstart() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnLoadstart
//go:noescape
func FuncHasOnLoadstart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnLoadstart
//go:noescape
func CallHasOnLoadstart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnLoadstart
//go:noescape
func TryHasOnLoadstart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnLoadstop
//go:noescape
func HasFuncOnLoadstop() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnLoadstop
//go:noescape
func FuncOnLoadstop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnLoadstop
//go:noescape
func CallOnLoadstop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnLoadstop
//go:noescape
func TryOnLoadstop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffLoadstop
//go:noescape
func HasFuncOffLoadstop() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffLoadstop
//go:noescape
func FuncOffLoadstop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffLoadstop
//go:noescape
func CallOffLoadstop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffLoadstop
//go:noescape
func TryOffLoadstop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnLoadstop
//go:noescape
func HasFuncHasOnLoadstop() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnLoadstop
//go:noescape
func FuncHasOnLoadstop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnLoadstop
//go:noescape
func CallHasOnLoadstop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnLoadstop
//go:noescape
func TryHasOnLoadstop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnNewwindow
//go:noescape
func HasFuncOnNewwindow() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnNewwindow
//go:noescape
func FuncOnNewwindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnNewwindow
//go:noescape
func CallOnNewwindow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnNewwindow
//go:noescape
func TryOnNewwindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffNewwindow
//go:noescape
func HasFuncOffNewwindow() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffNewwindow
//go:noescape
func FuncOffNewwindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffNewwindow
//go:noescape
func CallOffNewwindow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffNewwindow
//go:noescape
func TryOffNewwindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnNewwindow
//go:noescape
func HasFuncHasOnNewwindow() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnNewwindow
//go:noescape
func FuncHasOnNewwindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnNewwindow
//go:noescape
func CallHasOnNewwindow(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnNewwindow
//go:noescape
func TryHasOnNewwindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnPermissionrequest
//go:noescape
func HasFuncOnPermissionrequest() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnPermissionrequest
//go:noescape
func FuncOnPermissionrequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnPermissionrequest
//go:noescape
func CallOnPermissionrequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnPermissionrequest
//go:noescape
func TryOnPermissionrequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffPermissionrequest
//go:noescape
func HasFuncOffPermissionrequest() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffPermissionrequest
//go:noescape
func FuncOffPermissionrequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffPermissionrequest
//go:noescape
func CallOffPermissionrequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffPermissionrequest
//go:noescape
func TryOffPermissionrequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnPermissionrequest
//go:noescape
func HasFuncHasOnPermissionrequest() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnPermissionrequest
//go:noescape
func FuncHasOnPermissionrequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnPermissionrequest
//go:noescape
func CallHasOnPermissionrequest(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnPermissionrequest
//go:noescape
func TryHasOnPermissionrequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Print
//go:noescape
func HasFuncPrint() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Print
//go:noescape
func FuncPrint(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Print
//go:noescape
func CallPrint(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_Print
//go:noescape
func TryPrint(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Reload
//go:noescape
func HasFuncReload() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Reload
//go:noescape
func FuncReload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Reload
//go:noescape
func CallReload(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_Reload
//go:noescape
func TryReload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_RemoveContentScripts
//go:noescape
func HasFuncRemoveContentScripts() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_RemoveContentScripts
//go:noescape
func FuncRemoveContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_RemoveContentScripts
//go:noescape
func CallRemoveContentScripts(
	retPtr unsafe.Pointer,
	scriptNameList js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_RemoveContentScripts
//go:noescape
func TryRemoveContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scriptNameList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag get_Request
//go:noescape
func GetRequest(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webviewtag set_Request
//go:noescape
func SetRequest(
	val unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webviewtag has_OnResponsive
//go:noescape
func HasFuncOnResponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnResponsive
//go:noescape
func FuncOnResponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnResponsive
//go:noescape
func CallOnResponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnResponsive
//go:noescape
func TryOnResponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffResponsive
//go:noescape
func HasFuncOffResponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffResponsive
//go:noescape
func FuncOffResponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffResponsive
//go:noescape
func CallOffResponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffResponsive
//go:noescape
func TryOffResponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnResponsive
//go:noescape
func HasFuncHasOnResponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnResponsive
//go:noescape
func FuncHasOnResponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnResponsive
//go:noescape
func CallHasOnResponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnResponsive
//go:noescape
func TryHasOnResponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_SetAudioMuted
//go:noescape
func HasFuncSetAudioMuted() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_SetAudioMuted
//go:noescape
func FuncSetAudioMuted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_SetAudioMuted
//go:noescape
func CallSetAudioMuted(
	retPtr unsafe.Pointer,
	mute js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_SetAudioMuted
//go:noescape
func TrySetAudioMuted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mute js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_SetSpatialNavigationEnabled
//go:noescape
func HasFuncSetSpatialNavigationEnabled() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_SetSpatialNavigationEnabled
//go:noescape
func FuncSetSpatialNavigationEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_SetSpatialNavigationEnabled
//go:noescape
func CallSetSpatialNavigationEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_SetSpatialNavigationEnabled
//go:noescape
func TrySetSpatialNavigationEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_SetUserAgentOverride
//go:noescape
func HasFuncSetUserAgentOverride() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_SetUserAgentOverride
//go:noescape
func FuncSetUserAgentOverride(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_SetUserAgentOverride
//go:noescape
func CallSetUserAgentOverride(
	retPtr unsafe.Pointer,
	userAgent js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_SetUserAgentOverride
//go:noescape
func TrySetUserAgentOverride(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	userAgent js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_SetZoom
//go:noescape
func HasFuncSetZoom() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_SetZoom
//go:noescape
func FuncSetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_SetZoom
//go:noescape
func CallSetZoom(
	retPtr unsafe.Pointer,
	zoomFactor float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_SetZoom
//go:noescape
func TrySetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	zoomFactor float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_SetZoomMode
//go:noescape
func HasFuncSetZoomMode() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_SetZoomMode
//go:noescape
func FuncSetZoomMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_SetZoomMode
//go:noescape
func CallSetZoomMode(
	retPtr unsafe.Pointer,
	ZoomMode uint32,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_SetZoomMode
//go:noescape
func TrySetZoomMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ZoomMode uint32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnSizechanged
//go:noescape
func HasFuncOnSizechanged() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnSizechanged
//go:noescape
func FuncOnSizechanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnSizechanged
//go:noescape
func CallOnSizechanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnSizechanged
//go:noescape
func TryOnSizechanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffSizechanged
//go:noescape
func HasFuncOffSizechanged() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffSizechanged
//go:noescape
func FuncOffSizechanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffSizechanged
//go:noescape
func CallOffSizechanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffSizechanged
//go:noescape
func TryOffSizechanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnSizechanged
//go:noescape
func HasFuncHasOnSizechanged() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnSizechanged
//go:noescape
func FuncHasOnSizechanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnSizechanged
//go:noescape
func CallHasOnSizechanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnSizechanged
//go:noescape
func TryHasOnSizechanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Stop
//go:noescape
func HasFuncStop() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Stop
//go:noescape
func FuncStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Stop
//go:noescape
func CallStop(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_Stop
//go:noescape
func TryStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_StopFinding
//go:noescape
func HasFuncStopFinding() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_StopFinding
//go:noescape
func FuncStopFinding(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_StopFinding
//go:noescape
func CallStopFinding(
	retPtr unsafe.Pointer,
	action uint32)

//go:wasmimport plat/js/webext/webviewtag try_StopFinding
//go:noescape
func TryStopFinding(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	action uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_Terminate
//go:noescape
func HasFuncTerminate() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_Terminate
//go:noescape
func FuncTerminate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_Terminate
//go:noescape
func CallTerminate(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag try_Terminate
//go:noescape
func TryTerminate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnUnresponsive
//go:noescape
func HasFuncOnUnresponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnUnresponsive
//go:noescape
func FuncOnUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnUnresponsive
//go:noescape
func CallOnUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnUnresponsive
//go:noescape
func TryOnUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffUnresponsive
//go:noescape
func HasFuncOffUnresponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffUnresponsive
//go:noescape
func FuncOffUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffUnresponsive
//go:noescape
func CallOffUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffUnresponsive
//go:noescape
func TryOffUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnUnresponsive
//go:noescape
func HasFuncHasOnUnresponsive() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnUnresponsive
//go:noescape
func FuncHasOnUnresponsive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnUnresponsive
//go:noescape
func CallHasOnUnresponsive(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnUnresponsive
//go:noescape
func TryHasOnUnresponsive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OnZoomchange
//go:noescape
func HasFuncOnZoomchange() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OnZoomchange
//go:noescape
func FuncOnZoomchange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OnZoomchange
//go:noescape
func CallOnZoomchange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OnZoomchange
//go:noescape
func TryOnZoomchange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_OffZoomchange
//go:noescape
func HasFuncOffZoomchange() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_OffZoomchange
//go:noescape
func FuncOffZoomchange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_OffZoomchange
//go:noescape
func CallOffZoomchange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_OffZoomchange
//go:noescape
func TryOffZoomchange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewtag has_HasOnZoomchange
//go:noescape
func HasFuncHasOnZoomchange() js.Ref

//go:wasmimport plat/js/webext/webviewtag func_HasOnZoomchange
//go:noescape
func FuncHasOnZoomchange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewtag call_HasOnZoomchange
//go:noescape
func CallHasOnZoomchange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewtag try_HasOnZoomchange
//go:noescape
func TryHasOnZoomchange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

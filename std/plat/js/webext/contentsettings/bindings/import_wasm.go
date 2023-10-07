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

//go:wasmimport plat/js/webext/contentsettings constof_AutoVerifyContentSetting
//go:noescape
func ConstOfAutoVerifyContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_CameraContentSetting
//go:noescape
func ConstOfCameraContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_Scope
//go:noescape
func ConstOfScope(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings store_ClearArgDetails
//go:noescape
func ClearArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentsettings load_ClearArgDetails
//go:noescape
func ClearArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings store_GetReturnType
//go:noescape
func GetReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentsettings load_GetReturnType
//go:noescape
func GetReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings store_ResourceIdentifier
//go:noescape
func ResourceIdentifierJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentsettings load_ResourceIdentifier
//go:noescape
func ResourceIdentifierJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings store_GetArgDetails
//go:noescape
func GetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentsettings load_GetArgDetails
//go:noescape
func GetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings store_SetArgDetails
//go:noescape
func SetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentsettings load_SetArgDetails
//go:noescape
func SetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings has_ContentSetting_Clear
//go:noescape
func HasFuncContentSettingClear(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings func_ContentSetting_Clear
//go:noescape
func FuncContentSettingClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings call_ContentSetting_Clear
//go:noescape
func CallContentSettingClear(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings try_ContentSetting_Clear
//go:noescape
func TryContentSettingClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/contentsettings has_ContentSetting_Get
//go:noescape
func HasFuncContentSettingGet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings func_ContentSetting_Get
//go:noescape
func FuncContentSettingGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings call_ContentSetting_Get
//go:noescape
func CallContentSettingGet(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings try_ContentSetting_Get
//go:noescape
func TryContentSettingGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/contentsettings has_ContentSetting_Set
//go:noescape
func HasFuncContentSettingSet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings func_ContentSetting_Set
//go:noescape
func FuncContentSettingSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings call_ContentSetting_Set
//go:noescape
func CallContentSettingSet(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings try_ContentSetting_Set
//go:noescape
func TryContentSettingSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/contentsettings has_ContentSetting_GetResourceIdentifiers
//go:noescape
func HasFuncContentSettingGetResourceIdentifiers(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings func_ContentSetting_GetResourceIdentifiers
//go:noescape
func FuncContentSettingGetResourceIdentifiers(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings call_ContentSetting_GetResourceIdentifiers
//go:noescape
func CallContentSettingGetResourceIdentifiers(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/contentsettings try_ContentSetting_GetResourceIdentifiers
//go:noescape
func TryContentSettingGetResourceIdentifiers(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/contentsettings constof_CookiesContentSetting
//go:noescape
func ConstOfCookiesContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_FullscreenContentSetting
//go:noescape
func ConstOfFullscreenContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_ImagesContentSetting
//go:noescape
func ConstOfImagesContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_JavascriptContentSetting
//go:noescape
func ConstOfJavascriptContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_LocationContentSetting
//go:noescape
func ConstOfLocationContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_MicrophoneContentSetting
//go:noescape
func ConstOfMicrophoneContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_MouselockContentSetting
//go:noescape
func ConstOfMouselockContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_MultipleAutomaticDownloadsContentSetting
//go:noescape
func ConstOfMultipleAutomaticDownloadsContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_NotificationsContentSetting
//go:noescape
func ConstOfNotificationsContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_PluginsContentSetting
//go:noescape
func ConstOfPluginsContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_PopupsContentSetting
//go:noescape
func ConstOfPopupsContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings constof_PpapiBrokerContentSetting
//go:noescape
func ConstOfPpapiBrokerContentSetting(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentsettings get_AutoVerify
//go:noescape
func GetAutoVerify(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_AutoVerify
//go:noescape
func SetAutoVerify(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_AutomaticDownloads
//go:noescape
func GetAutomaticDownloads(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_AutomaticDownloads
//go:noescape
func SetAutomaticDownloads(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Camera
//go:noescape
func GetCamera(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Camera
//go:noescape
func SetCamera(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Cookies
//go:noescape
func GetCookies(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Cookies
//go:noescape
func SetCookies(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Fullscreen
//go:noescape
func GetFullscreen(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Fullscreen
//go:noescape
func SetFullscreen(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Images
//go:noescape
func GetImages(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Images
//go:noescape
func SetImages(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Javascript
//go:noescape
func GetJavascript(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Javascript
//go:noescape
func SetJavascript(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Location
//go:noescape
func GetLocation(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Location
//go:noescape
func SetLocation(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Microphone
//go:noescape
func GetMicrophone(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Microphone
//go:noescape
func SetMicrophone(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Mouselock
//go:noescape
func GetMouselock(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Mouselock
//go:noescape
func SetMouselock(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Notifications
//go:noescape
func GetNotifications(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Notifications
//go:noescape
func SetNotifications(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Plugins
//go:noescape
func GetPlugins(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Plugins
//go:noescape
func SetPlugins(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_Popups
//go:noescape
func GetPopups(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_Popups
//go:noescape
func SetPopups(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentsettings get_UnsandboxedPlugins
//go:noescape
func GetUnsandboxedPlugins(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/contentsettings set_UnsandboxedPlugins
//go:noescape
func SetUnsandboxedPlugins(
	val js.Ref) js.Ref

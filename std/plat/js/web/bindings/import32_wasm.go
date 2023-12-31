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

//go:wasmimport plat/js/web get_WindowSharedStorage_Worklet
//go:noescape
func GetWindowSharedStorageWorklet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowSharedStorage_Run
//go:noescape
func HasFuncWindowSharedStorageRun(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_Run
//go:noescape
func FuncWindowSharedStorageRun(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowSharedStorage_Run
//go:noescape
func CallWindowSharedStorageRun(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WindowSharedStorage_Run
//go:noescape
func TryWindowSharedStorageRun(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowSharedStorage_Run1
//go:noescape
func HasFuncWindowSharedStorageRun1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_Run1
//go:noescape
func FuncWindowSharedStorageRun1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowSharedStorage_Run1
//go:noescape
func CallWindowSharedStorageRun1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_WindowSharedStorage_Run1
//go:noescape
func TryWindowSharedStorageRun1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowSharedStorage_SelectURL
//go:noescape
func HasFuncWindowSharedStorageSelectURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_SelectURL
//go:noescape
func FuncWindowSharedStorageSelectURL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowSharedStorage_SelectURL
//go:noescape
func CallWindowSharedStorageSelectURL(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	urls js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_WindowSharedStorage_SelectURL
//go:noescape
func TryWindowSharedStorageSelectURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	urls js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowSharedStorage_SelectURL1
//go:noescape
func HasFuncWindowSharedStorageSelectURL1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_SelectURL1
//go:noescape
func FuncWindowSharedStorageSelectURL1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowSharedStorage_SelectURL1
//go:noescape
func CallWindowSharedStorageSelectURL1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	urls js.Ref)

//go:wasmimport plat/js/web try_WindowSharedStorage_SelectURL1
//go:noescape
func TryWindowSharedStorageSelectURL1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_CookieSameSite
//go:noescape
func ConstOfCookieSameSite(str js.Ref) uint32

//go:wasmimport plat/js/web store_CookieListItem
//go:noescape
func CookieListItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieListItem
//go:noescape
func CookieListItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CookieInit
//go:noescape
func CookieInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieInit
//go:noescape
func CookieInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CookieStoreDeleteOptions
//go:noescape
func CookieStoreDeleteOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieStoreDeleteOptions
//go:noescape
func CookieStoreDeleteOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CookieStore_Get
//go:noescape
func HasFuncCookieStoreGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get
//go:noescape
func FuncCookieStoreGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Get
//go:noescape
func CallCookieStoreGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_CookieStore_Get
//go:noescape
func TryCookieStoreGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Get1
//go:noescape
func HasFuncCookieStoreGet1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get1
//go:noescape
func FuncCookieStoreGet1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Get1
//go:noescape
func CallCookieStoreGet1(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_Get1
//go:noescape
func TryCookieStoreGet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Get2
//go:noescape
func HasFuncCookieStoreGet2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get2
//go:noescape
func FuncCookieStoreGet2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Get2
//go:noescape
func CallCookieStoreGet2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_Get2
//go:noescape
func TryCookieStoreGet2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_GetAll
//go:noescape
func HasFuncCookieStoreGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll
//go:noescape
func FuncCookieStoreGetAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_GetAll
//go:noescape
func CallCookieStoreGetAll(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_CookieStore_GetAll
//go:noescape
func TryCookieStoreGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_GetAll1
//go:noescape
func HasFuncCookieStoreGetAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll1
//go:noescape
func FuncCookieStoreGetAll1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_GetAll1
//go:noescape
func CallCookieStoreGetAll1(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_GetAll1
//go:noescape
func TryCookieStoreGetAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_GetAll2
//go:noescape
func HasFuncCookieStoreGetAll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll2
//go:noescape
func FuncCookieStoreGetAll2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_GetAll2
//go:noescape
func CallCookieStoreGetAll2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_GetAll2
//go:noescape
func TryCookieStoreGetAll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Set
//go:noescape
func HasFuncCookieStoreSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Set
//go:noescape
func FuncCookieStoreSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Set
//go:noescape
func CallCookieStoreSet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_CookieStore_Set
//go:noescape
func TryCookieStoreSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Set1
//go:noescape
func HasFuncCookieStoreSet1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Set1
//go:noescape
func FuncCookieStoreSet1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Set1
//go:noescape
func CallCookieStoreSet1(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_Set1
//go:noescape
func TryCookieStoreSet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Delete
//go:noescape
func HasFuncCookieStoreDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Delete
//go:noescape
func FuncCookieStoreDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Delete
//go:noescape
func CallCookieStoreDelete(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_CookieStore_Delete
//go:noescape
func TryCookieStoreDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStore_Delete1
//go:noescape
func HasFuncCookieStoreDelete1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Delete1
//go:noescape
func FuncCookieStoreDelete1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CookieStore_Delete1
//go:noescape
func CallCookieStoreDelete1(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStore_Delete1
//go:noescape
func TryCookieStoreDelete1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_DocumentPictureInPictureOptions
//go:noescape
func DocumentPictureInPictureOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DocumentPictureInPictureOptions
//go:noescape
func DocumentPictureInPictureOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_DocumentPictureInPicture_Window
//go:noescape
func GetDocumentPictureInPictureWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentPictureInPicture_RequestWindow
//go:noescape
func HasFuncDocumentPictureInPictureRequestWindow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentPictureInPicture_RequestWindow
//go:noescape
func FuncDocumentPictureInPictureRequestWindow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DocumentPictureInPicture_RequestWindow
//go:noescape
func CallDocumentPictureInPictureRequestWindow(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_DocumentPictureInPicture_RequestWindow
//go:noescape
func TryDocumentPictureInPictureRequestWindow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentPictureInPicture_RequestWindow1
//go:noescape
func HasFuncDocumentPictureInPictureRequestWindow1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentPictureInPicture_RequestWindow1
//go:noescape
func FuncDocumentPictureInPictureRequestWindow1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DocumentPictureInPicture_RequestWindow1
//go:noescape
func CallDocumentPictureInPictureRequestWindow1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DocumentPictureInPicture_RequestWindow1
//go:noescape
func TryDocumentPictureInPictureRequestWindow1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_External_AddSearchProvider
//go:noescape
func HasFuncExternalAddSearchProvider(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_External_AddSearchProvider
//go:noescape
func FuncExternalAddSearchProvider(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_External_AddSearchProvider
//go:noescape
func CallExternalAddSearchProvider(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_External_AddSearchProvider
//go:noescape
func TryExternalAddSearchProvider(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_External_IsSearchProviderInstalled
//go:noescape
func HasFuncExternalIsSearchProviderInstalled(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_External_IsSearchProviderInstalled
//go:noescape
func FuncExternalIsSearchProviderInstalled(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_External_IsSearchProviderInstalled
//go:noescape
func CallExternalIsSearchProviderInstalled(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_External_IsSearchProviderInstalled
//go:noescape
func TryExternalIsSearchProviderInstalled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_VoiceURI
//go:noescape
func GetSpeechSynthesisVoiceVoiceURI(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Name
//go:noescape
func GetSpeechSynthesisVoiceName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Lang
//go:noescape
func GetSpeechSynthesisVoiceLang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_LocalService
//go:noescape
func GetSpeechSynthesisVoiceLocalService(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Default
//go:noescape
func GetSpeechSynthesisVoiceDefault(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_SpeechSynthesisUtterance_SpeechSynthesisUtterance
//go:noescape
func NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance(
	text js.Ref) js.Ref

//go:wasmimport plat/js/web new_SpeechSynthesisUtterance_SpeechSynthesisUtterance1
//go:noescape
func NewSpeechSynthesisUtteranceBySpeechSynthesisUtterance1() js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Text
//go:noescape
func GetSpeechSynthesisUtteranceText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Text
//go:noescape
func SetSpeechSynthesisUtteranceText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Lang
//go:noescape
func GetSpeechSynthesisUtteranceLang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Lang
//go:noescape
func SetSpeechSynthesisUtteranceLang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Voice
//go:noescape
func GetSpeechSynthesisUtteranceVoice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Voice
//go:noescape
func SetSpeechSynthesisUtteranceVoice(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Volume
//go:noescape
func GetSpeechSynthesisUtteranceVolume(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Volume
//go:noescape
func SetSpeechSynthesisUtteranceVolume(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Rate
//go:noescape
func GetSpeechSynthesisUtteranceRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Rate
//go:noescape
func SetSpeechSynthesisUtteranceRate(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Pitch
//go:noescape
func GetSpeechSynthesisUtterancePitch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Pitch
//go:noescape
func SetSpeechSynthesisUtterancePitch(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesis_Pending
//go:noescape
func GetSpeechSynthesisPending(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesis_Speaking
//go:noescape
func GetSpeechSynthesisSpeaking(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SpeechSynthesis_Paused
//go:noescape
func GetSpeechSynthesisPaused(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SpeechSynthesis_Speak
//go:noescape
func HasFuncSpeechSynthesisSpeak(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Speak
//go:noescape
func FuncSpeechSynthesisSpeak(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SpeechSynthesis_Speak
//go:noescape
func CallSpeechSynthesisSpeak(
	this js.Ref, retPtr unsafe.Pointer,
	utterance js.Ref)

//go:wasmimport plat/js/web try_SpeechSynthesis_Speak
//go:noescape
func TrySpeechSynthesisSpeak(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	utterance js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SpeechSynthesis_Cancel
//go:noescape
func HasFuncSpeechSynthesisCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Cancel
//go:noescape
func FuncSpeechSynthesisCancel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SpeechSynthesis_Cancel
//go:noescape
func CallSpeechSynthesisCancel(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SpeechSynthesis_Cancel
//go:noescape
func TrySpeechSynthesisCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SpeechSynthesis_Pause
//go:noescape
func HasFuncSpeechSynthesisPause(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Pause
//go:noescape
func FuncSpeechSynthesisPause(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SpeechSynthesis_Pause
//go:noescape
func CallSpeechSynthesisPause(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SpeechSynthesis_Pause
//go:noescape
func TrySpeechSynthesisPause(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SpeechSynthesis_Resume
//go:noescape
func HasFuncSpeechSynthesisResume(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Resume
//go:noescape
func FuncSpeechSynthesisResume(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SpeechSynthesis_Resume
//go:noescape
func CallSpeechSynthesisResume(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SpeechSynthesis_Resume
//go:noescape
func TrySpeechSynthesisResume(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SpeechSynthesis_GetVoices
//go:noescape
func HasFuncSpeechSynthesisGetVoices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_GetVoices
//go:noescape
func FuncSpeechSynthesisGetVoices(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SpeechSynthesis_GetVoices
//go:noescape
func CallSpeechSynthesisGetVoices(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SpeechSynthesis_GetVoices
//go:noescape
func TrySpeechSynthesisGetVoices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_TaskPriority
//go:noescape
func ConstOfTaskPriority(str js.Ref) uint32

//go:wasmimport plat/js/web store_SchedulerPostTaskOptions
//go:noescape
func SchedulerPostTaskOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SchedulerPostTaskOptions
//go:noescape
func SchedulerPostTaskOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_Scheduler_PostTask
//go:noescape
func HasFuncSchedulerPostTask(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Scheduler_PostTask
//go:noescape
func FuncSchedulerPostTask(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Scheduler_PostTask
//go:noescape
func CallSchedulerPostTask(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Scheduler_PostTask
//go:noescape
func TrySchedulerPostTask(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Scheduler_PostTask1
//go:noescape
func HasFuncSchedulerPostTask1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Scheduler_PostTask1
//go:noescape
func FuncSchedulerPostTask1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Scheduler_PostTask1
//go:noescape
func CallSchedulerPostTask1(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_Scheduler_PostTask1
//go:noescape
func TrySchedulerPostTask1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedHTML_ToString
//go:noescape
func HasFuncTrustedHTMLToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_ToString
//go:noescape
func FuncTrustedHTMLToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedHTML_ToString
//go:noescape
func CallTrustedHTMLToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedHTML_ToString
//go:noescape
func TryTrustedHTMLToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedHTML_ToJSON
//go:noescape
func HasFuncTrustedHTMLToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_ToJSON
//go:noescape
func FuncTrustedHTMLToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedHTML_ToJSON
//go:noescape
func CallTrustedHTMLToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedHTML_ToJSON
//go:noescape
func TryTrustedHTMLToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedHTML_FromLiteral
//go:noescape
func HasFuncTrustedHTMLFromLiteral(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_FromLiteral
//go:noescape
func FuncTrustedHTMLFromLiteral(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedHTML_FromLiteral
//go:noescape
func CallTrustedHTMLFromLiteral(
	this js.Ref, retPtr unsafe.Pointer,
	templateStringsArray js.Ref)

//go:wasmimport plat/js/web try_TrustedHTML_FromLiteral
//go:noescape
func TryTrustedHTMLFromLiteral(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	templateStringsArray js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScript_ToString
//go:noescape
func HasFuncTrustedScriptToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_ToString
//go:noescape
func FuncTrustedScriptToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScript_ToString
//go:noescape
func CallTrustedScriptToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedScript_ToString
//go:noescape
func TryTrustedScriptToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScript_ToJSON
//go:noescape
func HasFuncTrustedScriptToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_ToJSON
//go:noescape
func FuncTrustedScriptToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScript_ToJSON
//go:noescape
func CallTrustedScriptToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedScript_ToJSON
//go:noescape
func TryTrustedScriptToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScript_FromLiteral
//go:noescape
func HasFuncTrustedScriptFromLiteral(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_FromLiteral
//go:noescape
func FuncTrustedScriptFromLiteral(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScript_FromLiteral
//go:noescape
func CallTrustedScriptFromLiteral(
	this js.Ref, retPtr unsafe.Pointer,
	templateStringsArray js.Ref)

//go:wasmimport plat/js/web try_TrustedScript_FromLiteral
//go:noescape
func TryTrustedScriptFromLiteral(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	templateStringsArray js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScriptURL_ToString
//go:noescape
func HasFuncTrustedScriptURLToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_ToString
//go:noescape
func FuncTrustedScriptURLToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScriptURL_ToString
//go:noescape
func CallTrustedScriptURLToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedScriptURL_ToString
//go:noescape
func TryTrustedScriptURLToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScriptURL_ToJSON
//go:noescape
func HasFuncTrustedScriptURLToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_ToJSON
//go:noescape
func FuncTrustedScriptURLToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScriptURL_ToJSON
//go:noescape
func CallTrustedScriptURLToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedScriptURL_ToJSON
//go:noescape
func TryTrustedScriptURLToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedScriptURL_FromLiteral
//go:noescape
func HasFuncTrustedScriptURLFromLiteral(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_FromLiteral
//go:noescape
func FuncTrustedScriptURLFromLiteral(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedScriptURL_FromLiteral
//go:noescape
func CallTrustedScriptURLFromLiteral(
	this js.Ref, retPtr unsafe.Pointer,
	templateStringsArray js.Ref)

//go:wasmimport plat/js/web try_TrustedScriptURL_FromLiteral
//go:noescape
func TryTrustedScriptURLFromLiteral(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	templateStringsArray js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_TrustedTypePolicy_Name
//go:noescape
func GetTrustedTypePolicyName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicy_CreateHTML
//go:noescape
func HasFuncTrustedTypePolicyCreateHTML(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateHTML
//go:noescape
func FuncTrustedTypePolicyCreateHTML(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateHTML
//go:noescape
func CallTrustedTypePolicyCreateHTML(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_TrustedTypePolicy_CreateHTML
//go:noescape
func TryTrustedTypePolicyCreateHTML(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicy_CreateScript
//go:noescape
func HasFuncTrustedTypePolicyCreateScript(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateScript
//go:noescape
func FuncTrustedTypePolicyCreateScript(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateScript
//go:noescape
func CallTrustedTypePolicyCreateScript(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_TrustedTypePolicy_CreateScript
//go:noescape
func TryTrustedTypePolicyCreateScript(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicy_CreateScriptURL
//go:noescape
func HasFuncTrustedTypePolicyCreateScriptURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateScriptURL
//go:noescape
func FuncTrustedTypePolicyCreateScriptURL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateScriptURL
//go:noescape
func CallTrustedTypePolicyCreateScriptURL(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_TrustedTypePolicy_CreateScriptURL
//go:noescape
func TryTrustedTypePolicyCreateScriptURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_TrustedTypePolicyOptions
//go:noescape
func TrustedTypePolicyOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TrustedTypePolicyOptions
//go:noescape
func TrustedTypePolicyOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_TrustedTypePolicyFactory_EmptyHTML
//go:noescape
func GetTrustedTypePolicyFactoryEmptyHTML(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TrustedTypePolicyFactory_EmptyScript
//go:noescape
func GetTrustedTypePolicyFactoryEmptyScript(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TrustedTypePolicyFactory_DefaultPolicy
//go:noescape
func GetTrustedTypePolicyFactoryDefaultPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_CreatePolicy
//go:noescape
func HasFuncTrustedTypePolicyFactoryCreatePolicy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_CreatePolicy
//go:noescape
func FuncTrustedTypePolicyFactoryCreatePolicy(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_CreatePolicy
//go:noescape
func CallTrustedTypePolicyFactoryCreatePolicy(
	this js.Ref, retPtr unsafe.Pointer,
	policyName js.Ref,
	policyOptions unsafe.Pointer)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_CreatePolicy
//go:noescape
func TryTrustedTypePolicyFactoryCreatePolicy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	policyName js.Ref,
	policyOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape
func HasFuncTrustedTypePolicyFactoryCreatePolicy1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape
func FuncTrustedTypePolicyFactoryCreatePolicy1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape
func CallTrustedTypePolicyFactoryCreatePolicy1(
	this js.Ref, retPtr unsafe.Pointer,
	policyName js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape
func TryTrustedTypePolicyFactoryCreatePolicy1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	policyName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_IsHTML
//go:noescape
func HasFuncTrustedTypePolicyFactoryIsHTML(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsHTML
//go:noescape
func FuncTrustedTypePolicyFactoryIsHTML(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsHTML
//go:noescape
func CallTrustedTypePolicyFactoryIsHTML(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_IsHTML
//go:noescape
func TryTrustedTypePolicyFactoryIsHTML(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_IsScript
//go:noescape
func HasFuncTrustedTypePolicyFactoryIsScript(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsScript
//go:noescape
func FuncTrustedTypePolicyFactoryIsScript(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsScript
//go:noescape
func CallTrustedTypePolicyFactoryIsScript(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_IsScript
//go:noescape
func TryTrustedTypePolicyFactoryIsScript(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_IsScriptURL
//go:noescape
func HasFuncTrustedTypePolicyFactoryIsScriptURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsScriptURL
//go:noescape
func FuncTrustedTypePolicyFactoryIsScriptURL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsScriptURL
//go:noescape
func CallTrustedTypePolicyFactoryIsScriptURL(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_IsScriptURL
//go:noescape
func TryTrustedTypePolicyFactoryIsScriptURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_GetAttributeType
//go:noescape
func HasFuncTrustedTypePolicyFactoryGetAttributeType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType
//go:noescape
func FuncTrustedTypePolicyFactoryGetAttributeType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType
//go:noescape
func CallTrustedTypePolicyFactoryGetAttributeType(
	this js.Ref, retPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref,
	attrNs js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_GetAttributeType
//go:noescape
func TryTrustedTypePolicyFactoryGetAttributeType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref,
	attrNs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape
func HasFuncTrustedTypePolicyFactoryGetAttributeType1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape
func FuncTrustedTypePolicyFactoryGetAttributeType1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape
func CallTrustedTypePolicyFactoryGetAttributeType1(
	this js.Ref, retPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape
func TryTrustedTypePolicyFactoryGetAttributeType1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape
func HasFuncTrustedTypePolicyFactoryGetAttributeType2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape
func FuncTrustedTypePolicyFactoryGetAttributeType2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape
func CallTrustedTypePolicyFactoryGetAttributeType2(
	this js.Ref, retPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape
func TryTrustedTypePolicyFactoryGetAttributeType2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tagName js.Ref,
	attribute js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_GetPropertyType
//go:noescape
func HasFuncTrustedTypePolicyFactoryGetPropertyType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetPropertyType
//go:noescape
func FuncTrustedTypePolicyFactoryGetPropertyType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetPropertyType
//go:noescape
func CallTrustedTypePolicyFactoryGetPropertyType(
	this js.Ref, retPtr unsafe.Pointer,
	tagName js.Ref,
	property js.Ref,
	elementNs js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_GetPropertyType
//go:noescape
func TryTrustedTypePolicyFactoryGetPropertyType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tagName js.Ref,
	property js.Ref,
	elementNs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape
func HasFuncTrustedTypePolicyFactoryGetPropertyType1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape
func FuncTrustedTypePolicyFactoryGetPropertyType1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape
func CallTrustedTypePolicyFactoryGetPropertyType1(
	this js.Ref, retPtr unsafe.Pointer,
	tagName js.Ref,
	property js.Ref)

//go:wasmimport plat/js/web try_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape
func TryTrustedTypePolicyFactoryGetPropertyType1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tagName js.Ref,
	property js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_KeyType
//go:noescape
func ConstOfKeyType(str js.Ref) uint32

//go:wasmimport plat/js/web get_CryptoKey_Type
//go:noescape
func GetCryptoKeyType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CryptoKey_Extractable
//go:noescape
func GetCryptoKeyExtractable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CryptoKey_Algorithm
//go:noescape
func GetCryptoKeyAlgorithm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CryptoKey_Usages
//go:noescape
func GetCryptoKeyUsages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_KeyUsage
//go:noescape
func ConstOfKeyUsage(str js.Ref) uint32

//go:wasmimport plat/js/web constof_KeyFormat
//go:noescape
func ConstOfKeyFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_RsaOtherPrimesInfo
//go:noescape
func RsaOtherPrimesInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RsaOtherPrimesInfo
//go:noescape
func RsaOtherPrimesInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_JsonWebKey
//go:noescape
func JsonWebKeyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_JsonWebKey
//go:noescape
func JsonWebKeyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_SubtleCrypto_Encrypt
//go:noescape
func HasFuncSubtleCryptoEncrypt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Encrypt
//go:noescape
func FuncSubtleCryptoEncrypt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_Encrypt
//go:noescape
func CallSubtleCryptoEncrypt(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_Encrypt
//go:noescape
func TrySubtleCryptoEncrypt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_Decrypt
//go:noescape
func HasFuncSubtleCryptoDecrypt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Decrypt
//go:noescape
func FuncSubtleCryptoDecrypt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_Decrypt
//go:noescape
func CallSubtleCryptoDecrypt(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_Decrypt
//go:noescape
func TrySubtleCryptoDecrypt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_Sign
//go:noescape
func HasFuncSubtleCryptoSign(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Sign
//go:noescape
func FuncSubtleCryptoSign(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_Sign
//go:noescape
func CallSubtleCryptoSign(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_Sign
//go:noescape
func TrySubtleCryptoSign(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_Verify
//go:noescape
func HasFuncSubtleCryptoVerify(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Verify
//go:noescape
func FuncSubtleCryptoVerify(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_Verify
//go:noescape
func CallSubtleCryptoVerify(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	signature js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_Verify
//go:noescape
func TrySubtleCryptoVerify(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	key js.Ref,
	signature js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_Digest
//go:noescape
func HasFuncSubtleCryptoDigest(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Digest
//go:noescape
func FuncSubtleCryptoDigest(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_Digest
//go:noescape
func CallSubtleCryptoDigest(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_Digest
//go:noescape
func TrySubtleCryptoDigest(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_GenerateKey
//go:noescape
func HasFuncSubtleCryptoGenerateKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_GenerateKey
//go:noescape
func FuncSubtleCryptoGenerateKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_GenerateKey
//go:noescape
func CallSubtleCryptoGenerateKey(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_GenerateKey
//go:noescape
func TrySubtleCryptoGenerateKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_DeriveKey
//go:noescape
func HasFuncSubtleCryptoDeriveKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_DeriveKey
//go:noescape
func FuncSubtleCryptoDeriveKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_DeriveKey
//go:noescape
func CallSubtleCryptoDeriveKey(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	baseKey js.Ref,
	derivedKeyType js.Ref,
	extractable js.Ref,
	keyUsages js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_DeriveKey
//go:noescape
func TrySubtleCryptoDeriveKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	baseKey js.Ref,
	derivedKeyType js.Ref,
	extractable js.Ref,
	keyUsages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_DeriveBits
//go:noescape
func HasFuncSubtleCryptoDeriveBits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_DeriveBits
//go:noescape
func FuncSubtleCryptoDeriveBits(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_DeriveBits
//go:noescape
func CallSubtleCryptoDeriveBits(
	this js.Ref, retPtr unsafe.Pointer,
	algorithm js.Ref,
	baseKey js.Ref,
	length uint32)

//go:wasmimport plat/js/web try_SubtleCrypto_DeriveBits
//go:noescape
func TrySubtleCryptoDeriveBits(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	algorithm js.Ref,
	baseKey js.Ref,
	length uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_ImportKey
//go:noescape
func HasFuncSubtleCryptoImportKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_ImportKey
//go:noescape
func FuncSubtleCryptoImportKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_ImportKey
//go:noescape
func CallSubtleCryptoImportKey(
	this js.Ref, retPtr unsafe.Pointer,
	format uint32,
	keyData js.Ref,
	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_ImportKey
//go:noescape
func TrySubtleCryptoImportKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format uint32,
	keyData js.Ref,
	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_ExportKey
//go:noescape
func HasFuncSubtleCryptoExportKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_ExportKey
//go:noescape
func FuncSubtleCryptoExportKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_ExportKey
//go:noescape
func CallSubtleCryptoExportKey(
	this js.Ref, retPtr unsafe.Pointer,
	format uint32,
	key js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_ExportKey
//go:noescape
func TrySubtleCryptoExportKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format uint32,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_WrapKey
//go:noescape
func HasFuncSubtleCryptoWrapKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_WrapKey
//go:noescape
func FuncSubtleCryptoWrapKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_WrapKey
//go:noescape
func CallSubtleCryptoWrapKey(
	this js.Ref, retPtr unsafe.Pointer,
	format uint32,
	key js.Ref,
	wrappingKey js.Ref,
	wrapAlgorithm js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_WrapKey
//go:noescape
func TrySubtleCryptoWrapKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format uint32,
	key js.Ref,
	wrappingKey js.Ref,
	wrapAlgorithm js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SubtleCrypto_UnwrapKey
//go:noescape
func HasFuncSubtleCryptoUnwrapKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_UnwrapKey
//go:noescape
func FuncSubtleCryptoUnwrapKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SubtleCrypto_UnwrapKey
//go:noescape
func CallSubtleCryptoUnwrapKey(
	this js.Ref, retPtr unsafe.Pointer,
	format uint32,
	wrappedKey js.Ref,
	unwrappingKey js.Ref,
	unwrapAlgorithm js.Ref,
	unwrappedKeyAlgorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref)

//go:wasmimport plat/js/web try_SubtleCrypto_UnwrapKey
//go:noescape
func TrySubtleCryptoUnwrapKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format uint32,
	wrappedKey js.Ref,
	unwrappingKey js.Ref,
	unwrapAlgorithm js.Ref,
	unwrappedKeyAlgorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_Crypto_Subtle
//go:noescape
func GetCryptoSubtle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Crypto_GetRandomValues
//go:noescape
func HasFuncCryptoGetRandomValues(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Crypto_GetRandomValues
//go:noescape
func FuncCryptoGetRandomValues(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Crypto_GetRandomValues
//go:noescape
func CallCryptoGetRandomValues(
	this js.Ref, retPtr unsafe.Pointer,
	array js.Ref)

//go:wasmimport plat/js/web try_Crypto_GetRandomValues
//go:noescape
func TryCryptoGetRandomValues(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Crypto_RandomUUID
//go:noescape
func HasFuncCryptoRandomUUID(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Crypto_RandomUUID
//go:noescape
func FuncCryptoRandomUUID(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Crypto_RandomUUID
//go:noescape
func CallCryptoRandomUUID(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Crypto_RandomUUID
//go:noescape
func TryCryptoRandomUUID(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MemoryAttributionContainer
//go:noescape
func MemoryAttributionContainerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MemoryAttributionContainer
//go:noescape
func MemoryAttributionContainerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MemoryAttribution
//go:noescape
func MemoryAttributionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MemoryAttribution
//go:noescape
func MemoryAttributionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MemoryBreakdownEntry
//go:noescape
func MemoryBreakdownEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MemoryBreakdownEntry
//go:noescape
func MemoryBreakdownEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MemoryMeasurement
//go:noescape
func MemoryMeasurementJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MemoryMeasurement
//go:noescape
func MemoryMeasurementJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PerformanceMarkOptions
//go:noescape
func PerformanceMarkOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PerformanceMarkOptions
//go:noescape
func PerformanceMarkOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PerformanceMark_PerformanceMark
//go:noescape
func NewPerformanceMarkByPerformanceMark(
	markName js.Ref,
	markOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PerformanceMark_PerformanceMark1
//go:noescape
func NewPerformanceMarkByPerformanceMark1(
	markName js.Ref) js.Ref

//go:wasmimport plat/js/web get_PerformanceMark_Detail
//go:noescape
func GetPerformanceMarkDetail(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceMeasure_Detail
//go:noescape
func GetPerformanceMeasureDetail(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

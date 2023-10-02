// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web get_WindowSharedStorage_Worklet
//go:noescape

func GetWindowSharedStorageWorklet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WindowSharedStorage_Run
//go:noescape

func CallWindowSharedStorageRun(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_Run
//go:noescape

func WindowSharedStorageRunFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WindowSharedStorage_Run1
//go:noescape

func CallWindowSharedStorageRun1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_Run1
//go:noescape

func WindowSharedStorageRun1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WindowSharedStorage_SelectURL
//go:noescape

func CallWindowSharedStorageSelectURL(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	urls js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_SelectURL
//go:noescape

func WindowSharedStorageSelectURLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WindowSharedStorage_SelectURL1
//go:noescape

func CallWindowSharedStorageSelectURL1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	urls js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WindowSharedStorage_SelectURL1
//go:noescape

func WindowSharedStorageSelectURL1Func(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_CookieStore_Get
//go:noescape

func CallCookieStoreGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get
//go:noescape

func CookieStoreGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Get1
//go:noescape

func CallCookieStoreGet1(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get1
//go:noescape

func CookieStoreGet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Get2
//go:noescape

func CallCookieStoreGet2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Get2
//go:noescape

func CookieStoreGet2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_GetAll
//go:noescape

func CallCookieStoreGetAll(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll
//go:noescape

func CookieStoreGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_GetAll1
//go:noescape

func CallCookieStoreGetAll1(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll1
//go:noescape

func CookieStoreGetAll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_GetAll2
//go:noescape

func CallCookieStoreGetAll2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CookieStore_GetAll2
//go:noescape

func CookieStoreGetAll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Set
//go:noescape

func CallCookieStoreSet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Set
//go:noescape

func CookieStoreSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Set1
//go:noescape

func CallCookieStoreSet1(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Set1
//go:noescape

func CookieStoreSet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Delete
//go:noescape

func CallCookieStoreDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Delete
//go:noescape

func CookieStoreDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStore_Delete1
//go:noescape

func CallCookieStoreDelete1(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CookieStore_Delete1
//go:noescape

func CookieStoreDelete1Func(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_DocumentPictureInPicture_RequestWindow
//go:noescape

func CallDocumentPictureInPictureRequestWindow(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DocumentPictureInPicture_RequestWindow
//go:noescape

func DocumentPictureInPictureRequestWindowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentPictureInPicture_RequestWindow1
//go:noescape

func CallDocumentPictureInPictureRequestWindow1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DocumentPictureInPicture_RequestWindow1
//go:noescape

func DocumentPictureInPictureRequestWindow1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_External_AddSearchProvider
//go:noescape

func CallExternalAddSearchProvider(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_External_AddSearchProvider
//go:noescape

func ExternalAddSearchProviderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_External_IsSearchProviderInstalled
//go:noescape

func CallExternalIsSearchProviderInstalled(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_External_IsSearchProviderInstalled
//go:noescape

func ExternalIsSearchProviderInstalledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_VoiceURI
//go:noescape

func GetSpeechSynthesisVoiceVoiceURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Name
//go:noescape

func GetSpeechSynthesisVoiceName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Lang
//go:noescape

func GetSpeechSynthesisVoiceLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_LocalService
//go:noescape

func GetSpeechSynthesisVoiceLocalService(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisVoice_Default
//go:noescape

func GetSpeechSynthesisVoiceDefault(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Text
//go:noescape

func SetSpeechSynthesisUtteranceText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Lang
//go:noescape

func GetSpeechSynthesisUtteranceLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Lang
//go:noescape

func SetSpeechSynthesisUtteranceLang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Voice
//go:noescape

func GetSpeechSynthesisUtteranceVoice(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Voice
//go:noescape

func SetSpeechSynthesisUtteranceVoice(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Volume
//go:noescape

func GetSpeechSynthesisUtteranceVolume(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Volume
//go:noescape

func SetSpeechSynthesisUtteranceVolume(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Rate
//go:noescape

func GetSpeechSynthesisUtteranceRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Rate
//go:noescape

func SetSpeechSynthesisUtteranceRate(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisUtterance_Pitch
//go:noescape

func GetSpeechSynthesisUtterancePitch(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SpeechSynthesisUtterance_Pitch
//go:noescape

func SetSpeechSynthesisUtterancePitch(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesis_Pending
//go:noescape

func GetSpeechSynthesisPending(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesis_Speaking
//go:noescape

func GetSpeechSynthesisSpeaking(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesis_Paused
//go:noescape

func GetSpeechSynthesisPaused(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SpeechSynthesis_Speak
//go:noescape

func CallSpeechSynthesisSpeak(
	this js.Ref,
	ptr unsafe.Pointer,

	utterance js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Speak
//go:noescape

func SpeechSynthesisSpeakFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechSynthesis_Cancel
//go:noescape

func CallSpeechSynthesisCancel(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Cancel
//go:noescape

func SpeechSynthesisCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechSynthesis_Pause
//go:noescape

func CallSpeechSynthesisPause(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Pause
//go:noescape

func SpeechSynthesisPauseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechSynthesis_Resume
//go:noescape

func CallSpeechSynthesisResume(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_Resume
//go:noescape

func SpeechSynthesisResumeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechSynthesis_GetVoices
//go:noescape

func CallSpeechSynthesisGetVoices(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechSynthesis_GetVoices
//go:noescape

func SpeechSynthesisGetVoicesFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_Scheduler_PostTask
//go:noescape

func CallSchedulerPostTask(
	this js.Ref,
	ptr unsafe.Pointer,

	callback js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Scheduler_PostTask
//go:noescape

func SchedulerPostTaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Scheduler_PostTask1
//go:noescape

func CallSchedulerPostTask1(
	this js.Ref,
	ptr unsafe.Pointer,

	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Scheduler_PostTask1
//go:noescape

func SchedulerPostTask1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedHTML_ToString
//go:noescape

func CallTrustedHTMLToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_ToString
//go:noescape

func TrustedHTMLToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedHTML_ToJSON
//go:noescape

func CallTrustedHTMLToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_ToJSON
//go:noescape

func TrustedHTMLToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedHTML_FromLiteral
//go:noescape

func CallTrustedHTMLFromLiteral(
	this js.Ref,
	ptr unsafe.Pointer,

	templateStringsArray js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedHTML_FromLiteral
//go:noescape

func TrustedHTMLFromLiteralFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScript_ToString
//go:noescape

func CallTrustedScriptToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_ToString
//go:noescape

func TrustedScriptToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScript_ToJSON
//go:noescape

func CallTrustedScriptToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_ToJSON
//go:noescape

func TrustedScriptToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScript_FromLiteral
//go:noescape

func CallTrustedScriptFromLiteral(
	this js.Ref,
	ptr unsafe.Pointer,

	templateStringsArray js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedScript_FromLiteral
//go:noescape

func TrustedScriptFromLiteralFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScriptURL_ToString
//go:noescape

func CallTrustedScriptURLToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_ToString
//go:noescape

func TrustedScriptURLToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScriptURL_ToJSON
//go:noescape

func CallTrustedScriptURLToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_ToJSON
//go:noescape

func TrustedScriptURLToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedScriptURL_FromLiteral
//go:noescape

func CallTrustedScriptURLFromLiteral(
	this js.Ref,
	ptr unsafe.Pointer,

	templateStringsArray js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedScriptURL_FromLiteral
//go:noescape

func TrustedScriptURLFromLiteralFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TrustedTypePolicy_Name
//go:noescape

func GetTrustedTypePolicyName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateHTML
//go:noescape

func CallTrustedTypePolicyCreateHTML(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateHTML
//go:noescape

func TrustedTypePolicyCreateHTMLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateScript
//go:noescape

func CallTrustedTypePolicyCreateScript(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateScript
//go:noescape

func TrustedTypePolicyCreateScriptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicy_CreateScriptURL
//go:noescape

func CallTrustedTypePolicyCreateScriptURL(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicy_CreateScriptURL
//go:noescape

func TrustedTypePolicyCreateScriptURLFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TrustedTypePolicyFactory_EmptyScript
//go:noescape

func GetTrustedTypePolicyFactoryEmptyScript(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TrustedTypePolicyFactory_DefaultPolicy
//go:noescape

func GetTrustedTypePolicyFactoryDefaultPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_CreatePolicy
//go:noescape

func CallTrustedTypePolicyFactoryCreatePolicy(
	this js.Ref,
	ptr unsafe.Pointer,

	policyName js.Ref,
	policyOptions unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_CreatePolicy
//go:noescape

func TrustedTypePolicyFactoryCreatePolicyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape

func CallTrustedTypePolicyFactoryCreatePolicy1(
	this js.Ref,
	ptr unsafe.Pointer,

	policyName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_CreatePolicy1
//go:noescape

func TrustedTypePolicyFactoryCreatePolicy1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsHTML
//go:noescape

func CallTrustedTypePolicyFactoryIsHTML(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsHTML
//go:noescape

func TrustedTypePolicyFactoryIsHTMLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsScript
//go:noescape

func CallTrustedTypePolicyFactoryIsScript(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsScript
//go:noescape

func TrustedTypePolicyFactoryIsScriptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_IsScriptURL
//go:noescape

func CallTrustedTypePolicyFactoryIsScriptURL(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_IsScriptURL
//go:noescape

func TrustedTypePolicyFactoryIsScriptURLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType
//go:noescape

func CallTrustedTypePolicyFactoryGetAttributeType(
	this js.Ref,
	ptr unsafe.Pointer,

	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref,
	attrNs js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType
//go:noescape

func TrustedTypePolicyFactoryGetAttributeTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape

func CallTrustedTypePolicyFactoryGetAttributeType1(
	this js.Ref,
	ptr unsafe.Pointer,

	tagName js.Ref,
	attribute js.Ref,
	elementNs js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType1
//go:noescape

func TrustedTypePolicyFactoryGetAttributeType1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape

func CallTrustedTypePolicyFactoryGetAttributeType2(
	this js.Ref,
	ptr unsafe.Pointer,

	tagName js.Ref,
	attribute js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetAttributeType2
//go:noescape

func TrustedTypePolicyFactoryGetAttributeType2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetPropertyType
//go:noescape

func CallTrustedTypePolicyFactoryGetPropertyType(
	this js.Ref,
	ptr unsafe.Pointer,

	tagName js.Ref,
	property js.Ref,
	elementNs js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetPropertyType
//go:noescape

func TrustedTypePolicyFactoryGetPropertyTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape

func CallTrustedTypePolicyFactoryGetPropertyType1(
	this js.Ref,
	ptr unsafe.Pointer,

	tagName js.Ref,
	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TrustedTypePolicyFactory_GetPropertyType1
//go:noescape

func TrustedTypePolicyFactoryGetPropertyType1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_KeyType
//go:noescape

func ConstOfKeyType(str js.Ref) uint32

//go:wasmimport plat/js/web get_CryptoKey_Type
//go:noescape

func GetCryptoKeyType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CryptoKey_Extractable
//go:noescape

func GetCryptoKeyExtractable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CryptoKey_Algorithm
//go:noescape

func GetCryptoKeyAlgorithm(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CryptoKey_Usages
//go:noescape

func GetCryptoKeyUsages(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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

//go:wasmimport plat/js/web call_SubtleCrypto_Encrypt
//go:noescape

func CallSubtleCryptoEncrypt(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	key js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Encrypt
//go:noescape

func SubtleCryptoEncryptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_Decrypt
//go:noescape

func CallSubtleCryptoDecrypt(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	key js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Decrypt
//go:noescape

func SubtleCryptoDecryptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_Sign
//go:noescape

func CallSubtleCryptoSign(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	key js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Sign
//go:noescape

func SubtleCryptoSignFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_Verify
//go:noescape

func CallSubtleCryptoVerify(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	key js.Ref,
	signature js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Verify
//go:noescape

func SubtleCryptoVerifyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_Digest
//go:noescape

func CallSubtleCryptoDigest(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_Digest
//go:noescape

func SubtleCryptoDigestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_GenerateKey
//go:noescape

func CallSubtleCryptoGenerateKey(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_GenerateKey
//go:noescape

func SubtleCryptoGenerateKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_DeriveKey
//go:noescape

func CallSubtleCryptoDeriveKey(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	baseKey js.Ref,
	derivedKeyType js.Ref,
	extractable js.Ref,
	keyUsages js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_DeriveKey
//go:noescape

func SubtleCryptoDeriveKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_DeriveBits
//go:noescape

func CallSubtleCryptoDeriveBits(
	this js.Ref,
	ptr unsafe.Pointer,

	algorithm js.Ref,
	baseKey js.Ref,
	length uint32,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_DeriveBits
//go:noescape

func SubtleCryptoDeriveBitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_ImportKey
//go:noescape

func CallSubtleCryptoImportKey(
	this js.Ref,
	ptr unsafe.Pointer,

	format uint32,
	keyData js.Ref,
	algorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_ImportKey
//go:noescape

func SubtleCryptoImportKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_ExportKey
//go:noescape

func CallSubtleCryptoExportKey(
	this js.Ref,
	ptr unsafe.Pointer,

	format uint32,
	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_ExportKey
//go:noescape

func SubtleCryptoExportKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_WrapKey
//go:noescape

func CallSubtleCryptoWrapKey(
	this js.Ref,
	ptr unsafe.Pointer,

	format uint32,
	key js.Ref,
	wrappingKey js.Ref,
	wrapAlgorithm js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_WrapKey
//go:noescape

func SubtleCryptoWrapKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SubtleCrypto_UnwrapKey
//go:noescape

func CallSubtleCryptoUnwrapKey(
	this js.Ref,
	ptr unsafe.Pointer,

	format uint32,
	wrappedKey js.Ref,
	unwrappingKey js.Ref,
	unwrapAlgorithm js.Ref,
	unwrappedKeyAlgorithm js.Ref,
	extractable js.Ref,
	keyUsages js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SubtleCrypto_UnwrapKey
//go:noescape

func SubtleCryptoUnwrapKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_Crypto_Subtle
//go:noescape

func GetCryptoSubtle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Crypto_GetRandomValues
//go:noescape

func CallCryptoGetRandomValues(
	this js.Ref,
	ptr unsafe.Pointer,

	array js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Crypto_GetRandomValues
//go:noescape

func CryptoGetRandomValuesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Crypto_RandomUUID
//go:noescape

func CallCryptoRandomUUID(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Crypto_RandomUUID
//go:noescape

func CryptoRandomUUIDFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PerformanceMeasure_Detail
//go:noescape

func GetPerformanceMeasureDetail(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

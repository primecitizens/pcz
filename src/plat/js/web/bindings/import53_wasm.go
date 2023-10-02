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

//go:wasmimport plat/js/web store_ScrollOptions
//go:noescape

func ScrollOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScrollOptions
//go:noescape

func ScrollOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ScrollSetting
//go:noescape

func ConstOfScrollSetting(str js.Ref) uint32

//go:wasmimport plat/js/web store_ScrollTimelineOptions
//go:noescape

func ScrollTimelineOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScrollTimelineOptions
//go:noescape

func ScrollTimelineOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ScrollTimeline_ScrollTimeline
//go:noescape

func NewScrollTimelineByScrollTimeline(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ScrollTimeline_ScrollTimeline1
//go:noescape

func NewScrollTimelineByScrollTimeline1() js.Ref

//go:wasmimport plat/js/web get_ScrollTimeline_Source
//go:noescape

func GetScrollTimelineSource(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ScrollTimeline_Axis
//go:noescape

func GetScrollTimelineAxis(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_SecurePaymentConfirmationRequest
//go:noescape

func SecurePaymentConfirmationRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SecurePaymentConfirmationRequest
//go:noescape

func SecurePaymentConfirmationRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SecurityPolicyViolationEventInit
//go:noescape

func SecurityPolicyViolationEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SecurityPolicyViolationEventInit
//go:noescape

func SecurityPolicyViolationEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SecurityPolicyViolationEvent_SecurityPolicyViolationEvent
//go:noescape

func NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_SecurityPolicyViolationEvent_SecurityPolicyViolationEvent1
//go:noescape

func NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_DocumentURI
//go:noescape

func GetSecurityPolicyViolationEventDocumentURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_Referrer
//go:noescape

func GetSecurityPolicyViolationEventReferrer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_BlockedURI
//go:noescape

func GetSecurityPolicyViolationEventBlockedURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_EffectiveDirective
//go:noescape

func GetSecurityPolicyViolationEventEffectiveDirective(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_ViolatedDirective
//go:noescape

func GetSecurityPolicyViolationEventViolatedDirective(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_OriginalPolicy
//go:noescape

func GetSecurityPolicyViolationEventOriginalPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_SourceFile
//go:noescape

func GetSecurityPolicyViolationEventSourceFile(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_Sample
//go:noescape

func GetSecurityPolicyViolationEventSample(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_Disposition
//go:noescape

func GetSecurityPolicyViolationEventDisposition(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_StatusCode
//go:noescape

func GetSecurityPolicyViolationEventStatusCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_LineNumber
//go:noescape

func GetSecurityPolicyViolationEventLineNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SecurityPolicyViolationEvent_ColumnNumber
//go:noescape

func GetSecurityPolicyViolationEventColumnNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Sensor_Activated
//go:noescape

func GetSensorActivated(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Sensor_HasReading
//go:noescape

func GetSensorHasReading(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Sensor_Timestamp
//go:noescape

func GetSensorTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_Sensor_Start
//go:noescape

func CallSensorStart(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Sensor_Start
//go:noescape

func SensorStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Sensor_Stop
//go:noescape

func CallSensorStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Sensor_Stop
//go:noescape

func SensorStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_SensorErrorEventInit
//go:noescape

func SensorErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SensorErrorEventInit
//go:noescape

func SensorErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SensorErrorEvent_SensorErrorEvent
//go:noescape

func NewSensorErrorEventBySensorErrorEvent(
	typ js.Ref,
	errorEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SensorErrorEvent_Error
//go:noescape

func GetSensorErrorEventError(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web new_SequenceEffect_SequenceEffect
//go:noescape

func NewSequenceEffectBySequenceEffect(
	children js.Ref,
	timing js.Ref) js.Ref

//go:wasmimport plat/js/web new_SequenceEffect_SequenceEffect1
//go:noescape

func NewSequenceEffectBySequenceEffect1(
	children js.Ref) js.Ref

//go:wasmimport plat/js/web call_SequenceEffect_Clone
//go:noescape

func CallSequenceEffectClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SequenceEffect_Clone
//go:noescape

func SequenceEffectCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerGlobalScope_Clients
//go:noescape

func GetServiceWorkerGlobalScopeClients(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerGlobalScope_Registration
//go:noescape

func GetServiceWorkerGlobalScopeRegistration(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerGlobalScope_ServiceWorker
//go:noescape

func GetServiceWorkerGlobalScopeServiceWorker(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerGlobalScope_CookieStore
//go:noescape

func GetServiceWorkerGlobalScopeCookieStore(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerGlobalScope_SkipWaiting
//go:noescape

func CallServiceWorkerGlobalScopeSkipWaiting(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerGlobalScope_SkipWaiting
//go:noescape

func ServiceWorkerGlobalScopeSkipWaitingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_ShadowAnimation_ShadowAnimation
//go:noescape

func NewShadowAnimationByShadowAnimation(
	source js.Ref,
	newTarget js.Ref) js.Ref

//go:wasmimport plat/js/web get_ShadowAnimation_SourceAnimation
//go:noescape

func GetShadowAnimationSourceAnimation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_SharedStorageSetMethodOptions
//go:noescape

func SharedStorageSetMethodOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SharedStorageSetMethodOptions
//go:noescape

func SharedStorageSetMethodOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorage_Set
//go:noescape

func CallSharedStorageSet(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
	value js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorage_Set
//go:noescape

func SharedStorageSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorage_Set1
//go:noescape

func CallSharedStorageSet1(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorage_Set1
//go:noescape

func SharedStorageSet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorage_Append
//go:noescape

func CallSharedStorageAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorage_Append
//go:noescape

func SharedStorageAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorage_Delete
//go:noescape

func CallSharedStorageDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorage_Delete
//go:noescape

func SharedStorageDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorage_Clear
//go:noescape

func CallSharedStorageClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SharedStorage_Clear
//go:noescape

func SharedStorageClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorageRunOperation_Run
//go:noescape

func CallSharedStorageRunOperationRun(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorageRunOperation_Run
//go:noescape

func SharedStorageRunOperationRunFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SharedStorageSelectURLOperation_Run
//go:noescape

func CallSharedStorageSelectURLOperationRun(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
	urls js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorageSelectURLOperation_Run
//go:noescape

func SharedStorageSelectURLOperationRunFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkletSharedStorage_Get
//go:noescape

func CallWorkletSharedStorageGet(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WorkletSharedStorage_Get
//go:noescape

func WorkletSharedStorageGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkletSharedStorage_Length
//go:noescape

func CallWorkletSharedStorageLength(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkletSharedStorage_Length
//go:noescape

func WorkletSharedStorageLengthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WorkletSharedStorage_RemainingBudget
//go:noescape

func CallWorkletSharedStorageRemainingBudget(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WorkletSharedStorage_RemainingBudget
//go:noescape

func WorkletSharedStorageRemainingBudgetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SharedStorageWorkletGlobalScope_SharedStorage
//go:noescape

func GetSharedStorageWorkletGlobalScopeSharedStorage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SharedStorageWorkletGlobalScope_Register
//go:noescape

func CallSharedStorageWorkletGlobalScopeRegister(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	operationCtor js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SharedStorageWorkletGlobalScope_Register
//go:noescape

func SharedStorageWorkletGlobalScopeRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_SharedWorker_SharedWorker
//go:noescape

func NewSharedWorkerBySharedWorker(
	scriptURL js.Ref,
	options js.Ref) js.Ref

//go:wasmimport plat/js/web new_SharedWorker_SharedWorker1
//go:noescape

func NewSharedWorkerBySharedWorker1(
	scriptURL js.Ref) js.Ref

//go:wasmimport plat/js/web get_SharedWorker_Port
//go:noescape

func GetSharedWorkerPort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SharedWorkerGlobalScope_Name
//go:noescape

func GetSharedWorkerGlobalScopeName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SharedWorkerGlobalScope_Close
//go:noescape

func CallSharedWorkerGlobalScopeClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SharedWorkerGlobalScope_Close
//go:noescape

func SharedWorkerGlobalScopeCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SpeechGrammar_Src
//go:noescape

func GetSpeechGrammarSrc(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechGrammar_Src
//go:noescape

func SetSpeechGrammarSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechGrammar_Weight
//go:noescape

func GetSpeechGrammarWeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SpeechGrammar_Weight
//go:noescape

func SetSpeechGrammarWeight(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SpeechGrammarList_Length
//go:noescape

func GetSpeechGrammarListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SpeechGrammarList_Item
//go:noescape

func CallSpeechGrammarListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SpeechGrammarList_Item
//go:noescape

func SpeechGrammarListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechGrammarList_AddFromURI
//go:noescape

func CallSpeechGrammarListAddFromURI(
	this js.Ref,
	ptr unsafe.Pointer,

	src js.Ref,
	weight float32,
) js.Ref

//go:wasmimport plat/js/web func_SpeechGrammarList_AddFromURI
//go:noescape

func SpeechGrammarListAddFromURIFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechGrammarList_AddFromURI1
//go:noescape

func CallSpeechGrammarListAddFromURI1(
	this js.Ref,
	ptr unsafe.Pointer,

	src js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SpeechGrammarList_AddFromURI1
//go:noescape

func SpeechGrammarListAddFromURI1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechGrammarList_AddFromString
//go:noescape

func CallSpeechGrammarListAddFromString(
	this js.Ref,
	ptr unsafe.Pointer,

	string js.Ref,
	weight float32,
) js.Ref

//go:wasmimport plat/js/web func_SpeechGrammarList_AddFromString
//go:noescape

func SpeechGrammarListAddFromStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechGrammarList_AddFromString1
//go:noescape

func CallSpeechGrammarListAddFromString1(
	this js.Ref,
	ptr unsafe.Pointer,

	string js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SpeechGrammarList_AddFromString1
//go:noescape

func SpeechGrammarListAddFromString1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognition_Grammars
//go:noescape

func GetSpeechRecognitionGrammars(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechRecognition_Grammars
//go:noescape

func SetSpeechRecognitionGrammars(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognition_Lang
//go:noescape

func GetSpeechRecognitionLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechRecognition_Lang
//go:noescape

func SetSpeechRecognitionLang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognition_Continuous
//go:noescape

func GetSpeechRecognitionContinuous(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechRecognition_Continuous
//go:noescape

func SetSpeechRecognitionContinuous(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognition_InterimResults
//go:noescape

func GetSpeechRecognitionInterimResults(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SpeechRecognition_InterimResults
//go:noescape

func SetSpeechRecognitionInterimResults(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognition_MaxAlternatives
//go:noescape

func GetSpeechRecognitionMaxAlternatives(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_SpeechRecognition_MaxAlternatives
//go:noescape

func SetSpeechRecognitionMaxAlternatives(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_SpeechRecognition_Start
//go:noescape

func CallSpeechRecognitionStart(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechRecognition_Start
//go:noescape

func SpeechRecognitionStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechRecognition_Stop
//go:noescape

func CallSpeechRecognitionStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechRecognition_Stop
//go:noescape

func SpeechRecognitionStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SpeechRecognition_Abort
//go:noescape

func CallSpeechRecognitionAbort(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SpeechRecognition_Abort
//go:noescape

func SpeechRecognitionAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionAlternative_Transcript
//go:noescape

func GetSpeechRecognitionAlternativeTranscript(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionAlternative_Confidence
//go:noescape

func GetSpeechRecognitionAlternativeConfidence(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web constof_SpeechRecognitionErrorCode
//go:noescape

func ConstOfSpeechRecognitionErrorCode(str js.Ref) uint32

//go:wasmimport plat/js/web store_SpeechRecognitionErrorEventInit
//go:noescape

func SpeechRecognitionErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SpeechRecognitionErrorEventInit
//go:noescape

func SpeechRecognitionErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SpeechRecognitionErrorEvent_SpeechRecognitionErrorEvent
//go:noescape

func NewSpeechRecognitionErrorEventBySpeechRecognitionErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionErrorEvent_Error
//go:noescape

func GetSpeechRecognitionErrorEventError(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SpeechRecognitionErrorEvent_Message
//go:noescape

func GetSpeechRecognitionErrorEventMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionResult_Length
//go:noescape

func GetSpeechRecognitionResultLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SpeechRecognitionResult_IsFinal
//go:noescape

func GetSpeechRecognitionResultIsFinal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SpeechRecognitionResult_Item
//go:noescape

func CallSpeechRecognitionResultItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SpeechRecognitionResult_Item
//go:noescape

func SpeechRecognitionResultItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionResultList_Length
//go:noescape

func GetSpeechRecognitionResultListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SpeechRecognitionResultList_Item
//go:noescape

func CallSpeechRecognitionResultListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SpeechRecognitionResultList_Item
//go:noescape

func SpeechRecognitionResultListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_SpeechRecognitionEventInit
//go:noescape

func SpeechRecognitionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SpeechRecognitionEventInit
//go:noescape

func SpeechRecognitionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SpeechRecognitionEvent_SpeechRecognitionEvent
//go:noescape

func NewSpeechRecognitionEventBySpeechRecognitionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SpeechRecognitionEvent_ResultIndex
//go:noescape

func GetSpeechRecognitionEventResultIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SpeechRecognitionEvent_Results
//go:noescape

func GetSpeechRecognitionEventResults(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_SpeechSynthesisErrorCode
//go:noescape

func ConstOfSpeechSynthesisErrorCode(str js.Ref) uint32

//go:wasmimport plat/js/web store_SpeechSynthesisErrorEventInit
//go:noescape

func SpeechSynthesisErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SpeechSynthesisErrorEventInit
//go:noescape

func SpeechSynthesisErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SpeechSynthesisErrorEvent_SpeechSynthesisErrorEvent
//go:noescape

func NewSpeechSynthesisErrorEventBySpeechSynthesisErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisErrorEvent_Error
//go:noescape

func GetSpeechSynthesisErrorEventError(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_SpeechSynthesisEventInit
//go:noescape

func SpeechSynthesisEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SpeechSynthesisEventInit
//go:noescape

func SpeechSynthesisEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SpeechSynthesisEvent_SpeechSynthesisEvent
//go:noescape

func NewSpeechSynthesisEventBySpeechSynthesisEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisEvent_Utterance
//go:noescape

func GetSpeechSynthesisEventUtterance(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SpeechSynthesisEvent_CharIndex
//go:noescape

func GetSpeechSynthesisEventCharIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SpeechSynthesisEvent_CharLength
//go:noescape

func GetSpeechSynthesisEventCharLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SpeechSynthesisEvent_ElapsedTime
//go:noescape

func GetSpeechSynthesisEventElapsedTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_SpeechSynthesisEvent_Name
//go:noescape

func GetSpeechSynthesisEventName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_StorageEventInit
//go:noescape

func StorageEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StorageEventInit
//go:noescape

func StorageEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_StorageEvent_StorageEvent
//go:noescape

func NewStorageEventByStorageEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_StorageEvent_StorageEvent1
//go:noescape

func NewStorageEventByStorageEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_StorageEvent_Key
//go:noescape

func GetStorageEventKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageEvent_OldValue
//go:noescape

func GetStorageEventOldValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageEvent_NewValue
//go:noescape

func GetStorageEventNewValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageEvent_Url
//go:noescape

func GetStorageEventUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StorageEvent_StorageArea
//go:noescape

func GetStorageEventStorageArea(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent
//go:noescape

func CallStorageEventInitStorageEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	key js.Ref,
	oldValue js.Ref,
	newValue js.Ref,
	url js.Ref,
	storageArea js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent
//go:noescape

func StorageEventInitStorageEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent1
//go:noescape

func CallStorageEventInitStorageEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	key js.Ref,
	oldValue js.Ref,
	newValue js.Ref,
	url js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent1
//go:noescape

func StorageEventInitStorageEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent2
//go:noescape

func CallStorageEventInitStorageEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	key js.Ref,
	oldValue js.Ref,
	newValue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent2
//go:noescape

func StorageEventInitStorageEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent3
//go:noescape

func CallStorageEventInitStorageEvent3(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	key js.Ref,
	oldValue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent3
//go:noescape

func StorageEventInitStorageEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent4
//go:noescape

func CallStorageEventInitStorageEvent4(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent4
//go:noescape

func StorageEventInitStorageEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent5
//go:noescape

func CallStorageEventInitStorageEvent5(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent5
//go:noescape

func StorageEventInitStorageEvent5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent6
//go:noescape

func CallStorageEventInitStorageEvent6(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
	bubbles js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent6
//go:noescape

func StorageEventInitStorageEvent6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StorageEvent_InitStorageEvent7
//go:noescape

func CallStorageEventInitStorageEvent7(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StorageEvent_InitStorageEvent7
//go:noescape

func StorageEventInitStorageEvent7Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_Type
//go:noescape

func GetStyleSheetType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_Href
//go:noescape

func GetStyleSheetHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_OwnerNode
//go:noescape

func GetStyleSheetOwnerNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_ParentStyleSheet
//go:noescape

func GetStyleSheetParentStyleSheet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_Title
//go:noescape

func GetStyleSheetTitle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_Media
//go:noescape

func GetStyleSheetMedia(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_StyleSheet_Disabled
//go:noescape

func GetStyleSheetDisabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_StyleSheet_Disabled
//go:noescape

func SetStyleSheetDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_SubmitEventInit
//go:noescape

func SubmitEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SubmitEventInit
//go:noescape

func SubmitEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SubmitEvent_SubmitEvent
//go:noescape

func NewSubmitEventBySubmitEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_SubmitEvent_SubmitEvent1
//go:noescape

func NewSubmitEventBySubmitEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_SubmitEvent_Submitter
//go:noescape

func GetSubmitEventSubmitter(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

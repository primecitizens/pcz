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

//go:wasmimport plat/js/web new_Instance_Instance
//go:noescape

func NewInstanceByInstance(
	module js.Ref,
	importObject js.Ref) js.Ref

//go:wasmimport plat/js/web new_Instance_Instance1
//go:noescape

func NewInstanceByInstance1(
	module js.Ref) js.Ref

//go:wasmimport plat/js/web get_Instance_Exports
//go:noescape

func GetInstanceExports(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape

func CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
	this js.Ref,
	ptr unsafe.Pointer,

	generateBidOutput unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape

func InterestGroupBiddingScriptRunnerGlobalScopeSetBidFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape

func CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape

func InterestGroupBiddingScriptRunnerGlobalScopeSetBid1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape

func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
	this js.Ref,
	ptr unsafe.Pointer,

	priority float64,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape

func InterestGroupBiddingScriptRunnerGlobalScopeSetPriorityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape

func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
	priority float64,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape

func InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverrideFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape

func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape

func InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape

func CallInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
	this js.Ref,
	ptr unsafe.Pointer,

	url js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape

func InterestGroupReportingScriptRunnerGlobalScopeSendReportToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape

func CallInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
	this js.Ref,
	ptr unsafe.Pointer,

	mapping js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape

func InterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeaconFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_IntersectionObserverEntryInit
//go:noescape

func IntersectionObserverEntryInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IntersectionObserverEntryInit
//go:noescape

func IntersectionObserverEntryInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_IntersectionObserverEntry_IntersectionObserverEntry
//go:noescape

func NewIntersectionObserverEntryByIntersectionObserverEntry(
	intersectionObserverEntryInit unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserverEntry_Time
//go:noescape

func GetIntersectionObserverEntryTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_IntersectionObserverEntry_RootBounds
//go:noescape

func GetIntersectionObserverEntryRootBounds(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserverEntry_BoundingClientRect
//go:noescape

func GetIntersectionObserverEntryBoundingClientRect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IntersectionRect
//go:noescape

func GetIntersectionObserverEntryIntersectionRect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IsIntersecting
//go:noescape

func GetIntersectionObserverEntryIsIntersecting(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IntersectionRatio
//go:noescape

func GetIntersectionObserverEntryIntersectionRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_IntersectionObserverEntry_Target
//go:noescape

func GetIntersectionObserverEntryTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_IntersectionObserverInit
//go:noescape

func IntersectionObserverInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IntersectionObserverInit
//go:noescape

func IntersectionObserverInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_IntersectionObserver_IntersectionObserver
//go:noescape

func NewIntersectionObserverByIntersectionObserver(
	callback js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_IntersectionObserver_IntersectionObserver1
//go:noescape

func NewIntersectionObserverByIntersectionObserver1(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserver_Root
//go:noescape

func GetIntersectionObserverRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserver_RootMargin
//go:noescape

func GetIntersectionObserverRootMargin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_IntersectionObserver_Thresholds
//go:noescape

func GetIntersectionObserverThresholds(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_IntersectionObserver_Observe
//go:noescape

func CallIntersectionObserverObserve(
	this js.Ref,
	ptr unsafe.Pointer,

	target js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Observe
//go:noescape

func IntersectionObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IntersectionObserver_Unobserve
//go:noescape

func CallIntersectionObserverUnobserve(
	this js.Ref,
	ptr unsafe.Pointer,

	target js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Unobserve
//go:noescape

func IntersectionObserverUnobserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IntersectionObserver_Disconnect
//go:noescape

func CallIntersectionObserverDisconnect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Disconnect
//go:noescape

func IntersectionObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IntersectionObserver_TakeRecords
//go:noescape

func CallIntersectionObserverTakeRecords(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_TakeRecords
//go:noescape

func IntersectionObserverTakeRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_InterventionReportBody_Id
//go:noescape

func GetInterventionReportBodyId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_InterventionReportBody_Message
//go:noescape

func GetInterventionReportBodyMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_InterventionReportBody_SourceFile
//go:noescape

func GetInterventionReportBodySourceFile(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_InterventionReportBody_LineNumber
//go:noescape

func GetInterventionReportBodyLineNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_InterventionReportBody_ColumnNumber
//go:noescape

func GetInterventionReportBodyColumnNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_InterventionReportBody_ToJSON
//go:noescape

func CallInterventionReportBodyToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_InterventionReportBody_ToJSON
//go:noescape

func InterventionReportBodyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_IntrinsicSizesResultOptions
//go:noescape

func IntrinsicSizesResultOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IntrinsicSizesResultOptions
//go:noescape

func IntrinsicSizesResultOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_JsonLdEmbed
//go:noescape

func ConstOfJsonLdEmbed(str js.Ref) uint32

//go:wasmimport plat/js/web constof_JsonLdErrorCode
//go:noescape

func ConstOfJsonLdErrorCode(str js.Ref) uint32

//go:wasmimport plat/js/web store_JsonLdError
//go:noescape

func JsonLdErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_JsonLdError
//go:noescape

func JsonLdErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_JsonLdFramingErrorCode
//go:noescape

func ConstOfJsonLdFramingErrorCode(str js.Ref) uint32

//go:wasmimport plat/js/web store_JsonLdFramingError
//go:noescape

func JsonLdFramingErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_JsonLdFramingError
//go:noescape

func JsonLdFramingErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_ContentType
//go:noescape

func GetRemoteDocumentContentType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_ContextUrl
//go:noescape

func GetRemoteDocumentContextUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_Document
//go:noescape

func GetRemoteDocumentDocument(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RemoteDocument_Document
//go:noescape

func SetRemoteDocumentDocument(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_DocumentUrl
//go:noescape

func GetRemoteDocumentDocumentUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_Profile
//go:noescape

func GetRemoteDocumentProfile(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_LoadDocumentOptions
//go:noescape

func LoadDocumentOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LoadDocumentOptions
//go:noescape

func LoadDocumentOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_JsonLdOptions
//go:noescape

func JsonLdOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_JsonLdOptions
//go:noescape

func JsonLdOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RdfLiteral_Value
//go:noescape

func GetRdfLiteralValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RdfLiteral_Datatype
//go:noescape

func GetRdfLiteralDatatype(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RdfLiteral_Language
//go:noescape

func GetRdfLiteralLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RdfTriple_Subject
//go:noescape

func GetRdfTripleSubject(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RdfTriple_Predicate
//go:noescape

func GetRdfTriplePredicate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RdfTriple_Object
//go:noescape

func GetRdfTripleObject(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RdfGraph_Add
//go:noescape

func CallRdfGraphAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	triple js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RdfGraph_Add
//go:noescape

func RdfGraphAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_RdfDataset_DefaultGraph
//go:noescape

func GetRdfDatasetDefaultGraph(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RdfDataset_Add
//go:noescape

func CallRdfDatasetAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	graphName js.Ref,
	graph js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RdfDataset_Add
//go:noescape

func RdfDatasetAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact
//go:noescape

func CallJsonLdProcessorCompact(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	context js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact
//go:noescape

func JsonLdProcessorCompactFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact1
//go:noescape

func CallJsonLdProcessorCompact1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	context js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact1
//go:noescape

func JsonLdProcessorCompact1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact2
//go:noescape

func CallJsonLdProcessorCompact2(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact2
//go:noescape

func JsonLdProcessorCompact2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Expand
//go:noescape

func CallJsonLdProcessorExpand(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Expand
//go:noescape

func JsonLdProcessorExpandFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Expand1
//go:noescape

func CallJsonLdProcessorExpand1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Expand1
//go:noescape

func JsonLdProcessorExpand1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten
//go:noescape

func CallJsonLdProcessorFlatten(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	context js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten
//go:noescape

func JsonLdProcessorFlattenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten1
//go:noescape

func CallJsonLdProcessorFlatten1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	context js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten1
//go:noescape

func JsonLdProcessorFlatten1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten2
//go:noescape

func CallJsonLdProcessorFlatten2(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten2
//go:noescape

func JsonLdProcessorFlatten2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_FromRdf
//go:noescape

func CallJsonLdProcessorFromRdf(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_FromRdf
//go:noescape

func JsonLdProcessorFromRdfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_FromRdf1
//go:noescape

func CallJsonLdProcessorFromRdf1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_FromRdf1
//go:noescape

func JsonLdProcessorFromRdf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_ToRdf
//go:noescape

func CallJsonLdProcessorToRdf(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_ToRdf
//go:noescape

func JsonLdProcessorToRdfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_ToRdf1
//go:noescape

func CallJsonLdProcessorToRdf1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_ToRdf1
//go:noescape

func JsonLdProcessorToRdf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Frame
//go:noescape

func CallJsonLdProcessorFrame(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	frame js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Frame
//go:noescape

func JsonLdProcessorFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_JsonLdProcessor_Frame1
//go:noescape

func CallJsonLdProcessorFrame1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	frame js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Frame1
//go:noescape

func JsonLdProcessorFrame1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_KeyboardEventInit
//go:noescape

func KeyboardEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_KeyboardEventInit
//go:noescape

func KeyboardEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_KeyboardEvent_KeyboardEvent
//go:noescape

func NewKeyboardEventByKeyboardEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_KeyboardEvent_KeyboardEvent1
//go:noescape

func NewKeyboardEventByKeyboardEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_Key
//go:noescape

func GetKeyboardEventKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_Code
//go:noescape

func GetKeyboardEventCode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_Location
//go:noescape

func GetKeyboardEventLocation(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_KeyboardEvent_CtrlKey
//go:noescape

func GetKeyboardEventCtrlKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_ShiftKey
//go:noescape

func GetKeyboardEventShiftKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_AltKey
//go:noescape

func GetKeyboardEventAltKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_MetaKey
//go:noescape

func GetKeyboardEventMetaKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_Repeat
//go:noescape

func GetKeyboardEventRepeat(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_IsComposing
//go:noescape

func GetKeyboardEventIsComposing(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_KeyboardEvent_CharCode
//go:noescape

func GetKeyboardEventCharCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_KeyboardEvent_KeyCode
//go:noescape

func GetKeyboardEventKeyCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_KeyboardEvent_GetModifierState
//go:noescape

func CallKeyboardEventGetModifierState(
	this js.Ref,
	ptr unsafe.Pointer,

	keyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_GetModifierState
//go:noescape

func KeyboardEventGetModifierStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent
//go:noescape

func CallKeyboardEventInitKeyboardEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref,
	metaKey js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent
//go:noescape

func KeyboardEventInitKeyboardEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent1
//go:noescape

func CallKeyboardEventInitKeyboardEvent1(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent1
//go:noescape

func KeyboardEventInitKeyboardEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent2
//go:noescape

func CallKeyboardEventInitKeyboardEvent2(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent2
//go:noescape

func KeyboardEventInitKeyboardEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent3
//go:noescape

func CallKeyboardEventInitKeyboardEvent3(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent3
//go:noescape

func KeyboardEventInitKeyboardEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent4
//go:noescape

func CallKeyboardEventInitKeyboardEvent4(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent4
//go:noescape

func KeyboardEventInitKeyboardEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent5
//go:noescape

func CallKeyboardEventInitKeyboardEvent5(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent5
//go:noescape

func KeyboardEventInitKeyboardEvent5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent6
//go:noescape

func CallKeyboardEventInitKeyboardEvent6(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent6
//go:noescape

func KeyboardEventInitKeyboardEvent6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent7
//go:noescape

func CallKeyboardEventInitKeyboardEvent7(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent7
//go:noescape

func KeyboardEventInitKeyboardEvent7Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent8
//go:noescape

func CallKeyboardEventInitKeyboardEvent8(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
	bubblesArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent8
//go:noescape

func KeyboardEventInitKeyboardEvent8Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent9
//go:noescape

func CallKeyboardEventInitKeyboardEvent9(
	this js.Ref,
	ptr unsafe.Pointer,

	typeArg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent9
//go:noescape

func KeyboardEventInitKeyboardEvent9Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_KeyframeEffectOptions
//go:noescape

func KeyframeEffectOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_KeyframeEffectOptions
//go:noescape

func KeyframeEffectOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_KeyframeEffect_KeyframeEffect
//go:noescape

func NewKeyframeEffectByKeyframeEffect(
	target js.Ref,
	keyframes js.Ref,
	options js.Ref) js.Ref

//go:wasmimport plat/js/web new_KeyframeEffect_KeyframeEffect1
//go:noescape

func NewKeyframeEffectByKeyframeEffect1(
	target js.Ref,
	keyframes js.Ref) js.Ref

//go:wasmimport plat/js/web new_KeyframeEffect_KeyframeEffect2
//go:noescape

func NewKeyframeEffectByKeyframeEffect2(
	source js.Ref) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_Target
//go:noescape

func GetKeyframeEffectTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_KeyframeEffect_Target
//go:noescape

func SetKeyframeEffectTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_PseudoElement
//go:noescape

func GetKeyframeEffectPseudoElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_KeyframeEffect_PseudoElement
//go:noescape

func SetKeyframeEffectPseudoElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_Composite
//go:noescape

func GetKeyframeEffectComposite(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_KeyframeEffect_Composite
//go:noescape

func SetKeyframeEffectComposite(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_IterationComposite
//go:noescape

func GetKeyframeEffectIterationComposite(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_KeyframeEffect_IterationComposite
//go:noescape

func SetKeyframeEffectIterationComposite(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_KeyframeEffect_GetKeyframes
//go:noescape

func CallKeyframeEffectGetKeyframes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_KeyframeEffect_GetKeyframes
//go:noescape

func KeyframeEffectGetKeyframesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_KeyframeEffect_SetKeyframes
//go:noescape

func CallKeyframeEffectSetKeyframes(
	this js.Ref,
	ptr unsafe.Pointer,

	keyframes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_KeyframeEffect_SetKeyframes
//go:noescape

func KeyframeEffectSetKeyframesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_LargeBlobSupport
//go:noescape

func ConstOfLargeBlobSupport(str js.Ref) uint32

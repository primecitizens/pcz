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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape
func HasFuncInterestGroupBiddingScriptRunnerGlobalScopeSetBid(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape
func FuncInterestGroupBiddingScriptRunnerGlobalScopeSetBid(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape
func CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
	this js.Ref, retPtr unsafe.Pointer,
	generateBidOutput unsafe.Pointer)

//go:wasmimport plat/js/web try_InterestGroupBiddingScriptRunnerGlobalScope_SetBid
//go:noescape
func TryInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	generateBidOutput unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape
func HasFuncInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape
func FuncInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape
func CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_InterestGroupBiddingScriptRunnerGlobalScope_SetBid1
//go:noescape
func TryInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape
func HasFuncInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape
func FuncInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape
func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
	this js.Ref, retPtr unsafe.Pointer,
	priority float64)

//go:wasmimport plat/js/web try_InterestGroupBiddingScriptRunnerGlobalScope_SetPriority
//go:noescape
func TryInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	priority float64) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape
func HasFuncInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape
func FuncInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape
func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref,
	priority float64)

//go:wasmimport plat/js/web try_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride
//go:noescape
func TryInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref,
	priority float64) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape
func HasFuncInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape
func FuncInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape
func CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_InterestGroupBiddingScriptRunnerGlobalScope_SetPrioritySignalsOverride1
//go:noescape
func TryInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape
func HasFuncInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape
func FuncInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape
func CallInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_InterestGroupReportingScriptRunnerGlobalScope_SendReportTo
//go:noescape
func TryInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape
func HasFuncInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape
func FuncInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape
func CallInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
	this js.Ref, retPtr unsafe.Pointer,
	mapping js.Ref)

//go:wasmimport plat/js/web try_InterestGroupReportingScriptRunnerGlobalScope_RegisterAdBeacon
//go:noescape
func TryInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mapping js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_RootBounds
//go:noescape
func GetIntersectionObserverEntryRootBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_BoundingClientRect
//go:noescape
func GetIntersectionObserverEntryBoundingClientRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IntersectionRect
//go:noescape
func GetIntersectionObserverEntryIntersectionRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IsIntersecting
//go:noescape
func GetIntersectionObserverEntryIsIntersecting(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_IntersectionRatio
//go:noescape
func GetIntersectionObserverEntryIntersectionRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserverEntry_Target
//go:noescape
func GetIntersectionObserverEntryTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserver_RootMargin
//go:noescape
func GetIntersectionObserverRootMargin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntersectionObserver_Thresholds
//go:noescape
func GetIntersectionObserverThresholds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IntersectionObserver_Observe
//go:noescape
func HasFuncIntersectionObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Observe
//go:noescape
func FuncIntersectionObserverObserve(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IntersectionObserver_Observe
//go:noescape
func CallIntersectionObserverObserve(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref)

//go:wasmimport plat/js/web try_IntersectionObserver_Observe
//go:noescape
func TryIntersectionObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IntersectionObserver_Unobserve
//go:noescape
func HasFuncIntersectionObserverUnobserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Unobserve
//go:noescape
func FuncIntersectionObserverUnobserve(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IntersectionObserver_Unobserve
//go:noescape
func CallIntersectionObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref)

//go:wasmimport plat/js/web try_IntersectionObserver_Unobserve
//go:noescape
func TryIntersectionObserverUnobserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IntersectionObserver_Disconnect
//go:noescape
func HasFuncIntersectionObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_Disconnect
//go:noescape
func FuncIntersectionObserverDisconnect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IntersectionObserver_Disconnect
//go:noescape
func CallIntersectionObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IntersectionObserver_Disconnect
//go:noescape
func TryIntersectionObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IntersectionObserver_TakeRecords
//go:noescape
func HasFuncIntersectionObserverTakeRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IntersectionObserver_TakeRecords
//go:noescape
func FuncIntersectionObserverTakeRecords(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IntersectionObserver_TakeRecords
//go:noescape
func CallIntersectionObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IntersectionObserver_TakeRecords
//go:noescape
func TryIntersectionObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InterventionReportBody_Id
//go:noescape
func GetInterventionReportBodyId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InterventionReportBody_Message
//go:noescape
func GetInterventionReportBodyMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InterventionReportBody_SourceFile
//go:noescape
func GetInterventionReportBodySourceFile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InterventionReportBody_LineNumber
//go:noescape
func GetInterventionReportBodyLineNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InterventionReportBody_ColumnNumber
//go:noescape
func GetInterventionReportBodyColumnNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InterventionReportBody_ToJSON
//go:noescape
func HasFuncInterventionReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InterventionReportBody_ToJSON
//go:noescape
func FuncInterventionReportBodyToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InterventionReportBody_ToJSON
//go:noescape
func CallInterventionReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_InterventionReportBody_ToJSON
//go:noescape
func TryInterventionReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RemoteDocument_ContextUrl
//go:noescape
func GetRemoteDocumentContextUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RemoteDocument_Document
//go:noescape
func GetRemoteDocumentDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RemoteDocument_Document
//go:noescape
func SetRemoteDocumentDocument(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_RemoteDocument_DocumentUrl
//go:noescape
func GetRemoteDocumentDocumentUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RemoteDocument_Profile
//go:noescape
func GetRemoteDocumentProfile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfLiteral_Datatype
//go:noescape
func GetRdfLiteralDatatype(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfLiteral_Language
//go:noescape
func GetRdfLiteralLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfTriple_Subject
//go:noescape
func GetRdfTripleSubject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfTriple_Predicate
//go:noescape
func GetRdfTriplePredicate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfTriple_Object
//go:noescape
func GetRdfTripleObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RdfGraph_Add
//go:noescape
func HasFuncRdfGraphAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RdfGraph_Add
//go:noescape
func FuncRdfGraphAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RdfGraph_Add
//go:noescape
func CallRdfGraphAdd(
	this js.Ref, retPtr unsafe.Pointer,
	triple js.Ref)

//go:wasmimport plat/js/web try_RdfGraph_Add
//go:noescape
func TryRdfGraphAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	triple js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_RdfDataset_DefaultGraph
//go:noescape
func GetRdfDatasetDefaultGraph(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RdfDataset_Add
//go:noescape
func HasFuncRdfDatasetAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RdfDataset_Add
//go:noescape
func FuncRdfDatasetAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RdfDataset_Add
//go:noescape
func CallRdfDatasetAdd(
	this js.Ref, retPtr unsafe.Pointer,
	graphName js.Ref,
	graph js.Ref)

//go:wasmimport plat/js/web try_RdfDataset_Add
//go:noescape
func TryRdfDatasetAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	graphName js.Ref,
	graph js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Compact
//go:noescape
func HasFuncJsonLdProcessorCompact(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact
//go:noescape
func FuncJsonLdProcessorCompact(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact
//go:noescape
func CallJsonLdProcessorCompact(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_Compact
//go:noescape
func TryJsonLdProcessorCompact(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Compact1
//go:noescape
func HasFuncJsonLdProcessorCompact1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact1
//go:noescape
func FuncJsonLdProcessorCompact1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact1
//go:noescape
func CallJsonLdProcessorCompact1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Compact1
//go:noescape
func TryJsonLdProcessorCompact1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Compact2
//go:noescape
func HasFuncJsonLdProcessorCompact2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Compact2
//go:noescape
func FuncJsonLdProcessorCompact2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Compact2
//go:noescape
func CallJsonLdProcessorCompact2(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Compact2
//go:noescape
func TryJsonLdProcessorCompact2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Expand
//go:noescape
func HasFuncJsonLdProcessorExpand(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Expand
//go:noescape
func FuncJsonLdProcessorExpand(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Expand
//go:noescape
func CallJsonLdProcessorExpand(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_Expand
//go:noescape
func TryJsonLdProcessorExpand(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Expand1
//go:noescape
func HasFuncJsonLdProcessorExpand1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Expand1
//go:noescape
func FuncJsonLdProcessorExpand1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Expand1
//go:noescape
func CallJsonLdProcessorExpand1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Expand1
//go:noescape
func TryJsonLdProcessorExpand1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Flatten
//go:noescape
func HasFuncJsonLdProcessorFlatten(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten
//go:noescape
func FuncJsonLdProcessorFlatten(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten
//go:noescape
func CallJsonLdProcessorFlatten(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_Flatten
//go:noescape
func TryJsonLdProcessorFlatten(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Flatten1
//go:noescape
func HasFuncJsonLdProcessorFlatten1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten1
//go:noescape
func FuncJsonLdProcessorFlatten1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten1
//go:noescape
func CallJsonLdProcessorFlatten1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Flatten1
//go:noescape
func TryJsonLdProcessorFlatten1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	context js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Flatten2
//go:noescape
func HasFuncJsonLdProcessorFlatten2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Flatten2
//go:noescape
func FuncJsonLdProcessorFlatten2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Flatten2
//go:noescape
func CallJsonLdProcessorFlatten2(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Flatten2
//go:noescape
func TryJsonLdProcessorFlatten2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_FromRdf
//go:noescape
func HasFuncJsonLdProcessorFromRdf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_FromRdf
//go:noescape
func FuncJsonLdProcessorFromRdf(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_FromRdf
//go:noescape
func CallJsonLdProcessorFromRdf(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_FromRdf
//go:noescape
func TryJsonLdProcessorFromRdf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_FromRdf1
//go:noescape
func HasFuncJsonLdProcessorFromRdf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_FromRdf1
//go:noescape
func FuncJsonLdProcessorFromRdf1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_FromRdf1
//go:noescape
func CallJsonLdProcessorFromRdf1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_FromRdf1
//go:noescape
func TryJsonLdProcessorFromRdf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_ToRdf
//go:noescape
func HasFuncJsonLdProcessorToRdf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_ToRdf
//go:noescape
func FuncJsonLdProcessorToRdf(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_ToRdf
//go:noescape
func CallJsonLdProcessorToRdf(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_ToRdf
//go:noescape
func TryJsonLdProcessorToRdf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_ToRdf1
//go:noescape
func HasFuncJsonLdProcessorToRdf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_ToRdf1
//go:noescape
func FuncJsonLdProcessorToRdf1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_ToRdf1
//go:noescape
func CallJsonLdProcessorToRdf1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_ToRdf1
//go:noescape
func TryJsonLdProcessorToRdf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Frame
//go:noescape
func HasFuncJsonLdProcessorFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Frame
//go:noescape
func FuncJsonLdProcessorFrame(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Frame
//go:noescape
func CallJsonLdProcessorFrame(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	frame js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_JsonLdProcessor_Frame
//go:noescape
func TryJsonLdProcessorFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	frame js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_JsonLdProcessor_Frame1
//go:noescape
func HasFuncJsonLdProcessorFrame1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_JsonLdProcessor_Frame1
//go:noescape
func FuncJsonLdProcessorFrame1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_JsonLdProcessor_Frame1
//go:noescape
func CallJsonLdProcessorFrame1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	frame js.Ref)

//go:wasmimport plat/js/web try_JsonLdProcessor_Frame1
//go:noescape
func TryJsonLdProcessorFrame1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	frame js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_Code
//go:noescape
func GetKeyboardEventCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_Location
//go:noescape
func GetKeyboardEventLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_CtrlKey
//go:noescape
func GetKeyboardEventCtrlKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_ShiftKey
//go:noescape
func GetKeyboardEventShiftKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_AltKey
//go:noescape
func GetKeyboardEventAltKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_MetaKey
//go:noescape
func GetKeyboardEventMetaKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_Repeat
//go:noescape
func GetKeyboardEventRepeat(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_IsComposing
//go:noescape
func GetKeyboardEventIsComposing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_CharCode
//go:noescape
func GetKeyboardEventCharCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_KeyboardEvent_KeyCode
//go:noescape
func GetKeyboardEventKeyCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_GetModifierState
//go:noescape
func HasFuncKeyboardEventGetModifierState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_GetModifierState
//go:noescape
func FuncKeyboardEventGetModifierState(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_GetModifierState
//go:noescape
func CallKeyboardEventGetModifierState(
	this js.Ref, retPtr unsafe.Pointer,
	keyArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_GetModifierState
//go:noescape
func TryKeyboardEventGetModifierState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent
//go:noescape
func FuncKeyboardEventInitKeyboardEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent
//go:noescape
func CallKeyboardEventInitKeyboardEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref,
	metaKey js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent
//go:noescape
func TryKeyboardEventInitKeyboardEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref,
	metaKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent1
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent1
//go:noescape
func FuncKeyboardEventInitKeyboardEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent1
//go:noescape
func CallKeyboardEventInitKeyboardEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent1
//go:noescape
func TryKeyboardEventInitKeyboardEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref,
	shiftKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent2
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent2
//go:noescape
func FuncKeyboardEventInitKeyboardEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent2
//go:noescape
func CallKeyboardEventInitKeyboardEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent2
//go:noescape
func TryKeyboardEventInitKeyboardEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref,
	altKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent3
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent3
//go:noescape
func FuncKeyboardEventInitKeyboardEvent3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent3
//go:noescape
func CallKeyboardEventInitKeyboardEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent3
//go:noescape
func TryKeyboardEventInitKeyboardEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32,
	ctrlKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent4
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent4
//go:noescape
func FuncKeyboardEventInitKeyboardEvent4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent4
//go:noescape
func CallKeyboardEventInitKeyboardEvent4(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent4
//go:noescape
func TryKeyboardEventInitKeyboardEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref,
	locationArg uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent5
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent5
//go:noescape
func FuncKeyboardEventInitKeyboardEvent5(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent5
//go:noescape
func CallKeyboardEventInitKeyboardEvent5(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent5
//go:noescape
func TryKeyboardEventInitKeyboardEvent5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	keyArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent6
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent6
//go:noescape
func FuncKeyboardEventInitKeyboardEvent6(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent6
//go:noescape
func CallKeyboardEventInitKeyboardEvent6(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent6
//go:noescape
func TryKeyboardEventInitKeyboardEvent6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent7
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent7(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent7
//go:noescape
func FuncKeyboardEventInitKeyboardEvent7(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent7
//go:noescape
func CallKeyboardEventInitKeyboardEvent7(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent7
//go:noescape
func TryKeyboardEventInitKeyboardEvent7(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent8
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent8(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent8
//go:noescape
func FuncKeyboardEventInitKeyboardEvent8(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent8
//go:noescape
func CallKeyboardEventInitKeyboardEvent8(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent8
//go:noescape
func TryKeyboardEventInitKeyboardEvent8(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyboardEvent_InitKeyboardEvent9
//go:noescape
func HasFuncKeyboardEventInitKeyboardEvent9(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyboardEvent_InitKeyboardEvent9
//go:noescape
func FuncKeyboardEventInitKeyboardEvent9(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyboardEvent_InitKeyboardEvent9
//go:noescape
func CallKeyboardEventInitKeyboardEvent9(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref)

//go:wasmimport plat/js/web try_KeyboardEvent_InitKeyboardEvent9
//go:noescape
func TryKeyboardEventInitKeyboardEvent9(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_KeyframeEffect_Target
//go:noescape
func SetKeyframeEffectTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_PseudoElement
//go:noescape
func GetKeyframeEffectPseudoElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_KeyframeEffect_PseudoElement
//go:noescape
func SetKeyframeEffectPseudoElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_Composite
//go:noescape
func GetKeyframeEffectComposite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_KeyframeEffect_Composite
//go:noescape
func SetKeyframeEffectComposite(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_KeyframeEffect_IterationComposite
//go:noescape
func GetKeyframeEffectIterationComposite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_KeyframeEffect_IterationComposite
//go:noescape
func SetKeyframeEffectIterationComposite(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web has_KeyframeEffect_GetKeyframes
//go:noescape
func HasFuncKeyframeEffectGetKeyframes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyframeEffect_GetKeyframes
//go:noescape
func FuncKeyframeEffectGetKeyframes(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyframeEffect_GetKeyframes
//go:noescape
func CallKeyframeEffectGetKeyframes(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_KeyframeEffect_GetKeyframes
//go:noescape
func TryKeyframeEffectGetKeyframes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_KeyframeEffect_SetKeyframes
//go:noescape
func HasFuncKeyframeEffectSetKeyframes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_KeyframeEffect_SetKeyframes
//go:noescape
func FuncKeyframeEffectSetKeyframes(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_KeyframeEffect_SetKeyframes
//go:noescape
func CallKeyframeEffectSetKeyframes(
	this js.Ref, retPtr unsafe.Pointer,
	keyframes js.Ref)

//go:wasmimport plat/js/web try_KeyframeEffect_SetKeyframes
//go:noescape
func TryKeyframeEffectSetKeyframes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyframes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_LargeBlobSupport
//go:noescape
func ConstOfLargeBlobSupport(str js.Ref) uint32

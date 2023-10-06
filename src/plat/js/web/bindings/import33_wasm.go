// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web store_PerformanceMeasureOptions
//go:noescape
func PerformanceMeasureOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PerformanceMeasureOptions
//go:noescape
func PerformanceMeasureOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PerformanceEntry_Name
//go:noescape
func GetPerformanceEntryName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEntry_EntryType
//go:noescape
func GetPerformanceEntryEntryType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEntry_StartTime
//go:noescape
func GetPerformanceEntryStartTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEntry_Duration
//go:noescape
func GetPerformanceEntryDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceEntry_ToJSON
//go:noescape
func HasPerformanceEntryToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceEntry_ToJSON
//go:noescape
func PerformanceEntryToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceEntry_ToJSON
//go:noescape
func CallPerformanceEntryToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceEntry_ToJSON
//go:noescape
func TryPerformanceEntryToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_NavigationStart
//go:noescape
func GetPerformanceTimingNavigationStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_UnloadEventStart
//go:noescape
func GetPerformanceTimingUnloadEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_UnloadEventEnd
//go:noescape
func GetPerformanceTimingUnloadEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_RedirectStart
//go:noescape
func GetPerformanceTimingRedirectStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_RedirectEnd
//go:noescape
func GetPerformanceTimingRedirectEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_FetchStart
//go:noescape
func GetPerformanceTimingFetchStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomainLookupStart
//go:noescape
func GetPerformanceTimingDomainLookupStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomainLookupEnd
//go:noescape
func GetPerformanceTimingDomainLookupEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_ConnectStart
//go:noescape
func GetPerformanceTimingConnectStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_ConnectEnd
//go:noescape
func GetPerformanceTimingConnectEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_SecureConnectionStart
//go:noescape
func GetPerformanceTimingSecureConnectionStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_RequestStart
//go:noescape
func GetPerformanceTimingRequestStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_ResponseStart
//go:noescape
func GetPerformanceTimingResponseStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_ResponseEnd
//go:noescape
func GetPerformanceTimingResponseEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomLoading
//go:noescape
func GetPerformanceTimingDomLoading(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomInteractive
//go:noescape
func GetPerformanceTimingDomInteractive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomContentLoadedEventStart
//go:noescape
func GetPerformanceTimingDomContentLoadedEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomContentLoadedEventEnd
//go:noescape
func GetPerformanceTimingDomContentLoadedEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_DomComplete
//go:noescape
func GetPerformanceTimingDomComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_LoadEventStart
//go:noescape
func GetPerformanceTimingLoadEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceTiming_LoadEventEnd
//go:noescape
func GetPerformanceTimingLoadEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceTiming_ToJSON
//go:noescape
func HasPerformanceTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceTiming_ToJSON
//go:noescape
func PerformanceTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceTiming_ToJSON
//go:noescape
func CallPerformanceTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceTiming_ToJSON
//go:noescape
func TryPerformanceTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigation_Type
//go:noescape
func GetPerformanceNavigationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigation_RedirectCount
//go:noescape
func GetPerformanceNavigationRedirectCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceNavigation_ToJSON
//go:noescape
func HasPerformanceNavigationToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceNavigation_ToJSON
//go:noescape
func PerformanceNavigationToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceNavigation_ToJSON
//go:noescape
func CallPerformanceNavigationToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceNavigation_ToJSON
//go:noescape
func TryPerformanceNavigationToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Performance_TimeOrigin
//go:noescape
func GetPerformanceTimeOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Performance_Timing
//go:noescape
func GetPerformanceTiming(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Performance_Navigation
//go:noescape
func GetPerformanceNavigation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Performance_EventCounts
//go:noescape
func GetPerformanceEventCounts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Performance_InteractionCount
//go:noescape
func GetPerformanceInteractionCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Now
//go:noescape
func HasPerformanceNow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Now
//go:noescape
func PerformanceNowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Now
//go:noescape
func CallPerformanceNow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_Now
//go:noescape
func TryPerformanceNow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ToJSON
//go:noescape
func HasPerformanceToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ToJSON
//go:noescape
func PerformanceToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ToJSON
//go:noescape
func CallPerformanceToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_ToJSON
//go:noescape
func TryPerformanceToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_MeasureUserAgentSpecificMemory
//go:noescape
func HasPerformanceMeasureUserAgentSpecificMemory(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_MeasureUserAgentSpecificMemory
//go:noescape
func PerformanceMeasureUserAgentSpecificMemoryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_MeasureUserAgentSpecificMemory
//go:noescape
func CallPerformanceMeasureUserAgentSpecificMemory(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_MeasureUserAgentSpecificMemory
//go:noescape
func TryPerformanceMeasureUserAgentSpecificMemory(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Mark
//go:noescape
func HasPerformanceMark(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Mark
//go:noescape
func PerformanceMarkFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Mark
//go:noescape
func CallPerformanceMark(
	this js.Ref, retPtr unsafe.Pointer,
	markName js.Ref,
	markOptions unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_Mark
//go:noescape
func TryPerformanceMark(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markName js.Ref,
	markOptions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Mark1
//go:noescape
func HasPerformanceMark1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Mark1
//go:noescape
func PerformanceMark1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Mark1
//go:noescape
func CallPerformanceMark1(
	this js.Ref, retPtr unsafe.Pointer,
	markName js.Ref)

//go:wasmimport plat/js/web try_Performance_Mark1
//go:noescape
func TryPerformanceMark1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ClearMarks
//go:noescape
func HasPerformanceClearMarks(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ClearMarks
//go:noescape
func PerformanceClearMarksFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ClearMarks
//go:noescape
func CallPerformanceClearMarks(
	this js.Ref, retPtr unsafe.Pointer,
	markName js.Ref)

//go:wasmimport plat/js/web try_Performance_ClearMarks
//go:noescape
func TryPerformanceClearMarks(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	markName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ClearMarks1
//go:noescape
func HasPerformanceClearMarks1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ClearMarks1
//go:noescape
func PerformanceClearMarks1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ClearMarks1
//go:noescape
func CallPerformanceClearMarks1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_ClearMarks1
//go:noescape
func TryPerformanceClearMarks1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Measure
//go:noescape
func HasPerformanceMeasure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Measure
//go:noescape
func PerformanceMeasureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Measure
//go:noescape
func CallPerformanceMeasure(
	this js.Ref, retPtr unsafe.Pointer,
	measureName js.Ref,
	startOrMeasureOptions js.Ref,
	endMark js.Ref)

//go:wasmimport plat/js/web try_Performance_Measure
//go:noescape
func TryPerformanceMeasure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	measureName js.Ref,
	startOrMeasureOptions js.Ref,
	endMark js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Measure1
//go:noescape
func HasPerformanceMeasure1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Measure1
//go:noescape
func PerformanceMeasure1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Measure1
//go:noescape
func CallPerformanceMeasure1(
	this js.Ref, retPtr unsafe.Pointer,
	measureName js.Ref,
	startOrMeasureOptions js.Ref)

//go:wasmimport plat/js/web try_Performance_Measure1
//go:noescape
func TryPerformanceMeasure1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	measureName js.Ref,
	startOrMeasureOptions js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_Measure2
//go:noescape
func HasPerformanceMeasure2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_Measure2
//go:noescape
func PerformanceMeasure2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_Measure2
//go:noescape
func CallPerformanceMeasure2(
	this js.Ref, retPtr unsafe.Pointer,
	measureName js.Ref)

//go:wasmimport plat/js/web try_Performance_Measure2
//go:noescape
func TryPerformanceMeasure2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	measureName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ClearMeasures
//go:noescape
func HasPerformanceClearMeasures(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ClearMeasures
//go:noescape
func PerformanceClearMeasuresFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ClearMeasures
//go:noescape
func CallPerformanceClearMeasures(
	this js.Ref, retPtr unsafe.Pointer,
	measureName js.Ref)

//go:wasmimport plat/js/web try_Performance_ClearMeasures
//go:noescape
func TryPerformanceClearMeasures(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	measureName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ClearMeasures1
//go:noescape
func HasPerformanceClearMeasures1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ClearMeasures1
//go:noescape
func PerformanceClearMeasures1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ClearMeasures1
//go:noescape
func CallPerformanceClearMeasures1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_ClearMeasures1
//go:noescape
func TryPerformanceClearMeasures1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_ClearResourceTimings
//go:noescape
func HasPerformanceClearResourceTimings(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_ClearResourceTimings
//go:noescape
func PerformanceClearResourceTimingsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_ClearResourceTimings
//go:noescape
func CallPerformanceClearResourceTimings(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_ClearResourceTimings
//go:noescape
func TryPerformanceClearResourceTimings(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_SetResourceTimingBufferSize
//go:noescape
func HasPerformanceSetResourceTimingBufferSize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_SetResourceTimingBufferSize
//go:noescape
func PerformanceSetResourceTimingBufferSizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_SetResourceTimingBufferSize
//go:noescape
func CallPerformanceSetResourceTimingBufferSize(
	this js.Ref, retPtr unsafe.Pointer,
	maxSize uint32)

//go:wasmimport plat/js/web try_Performance_SetResourceTimingBufferSize
//go:noescape
func TryPerformanceSetResourceTimingBufferSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	maxSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_GetEntries
//go:noescape
func HasPerformanceGetEntries(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_GetEntries
//go:noescape
func PerformanceGetEntriesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_GetEntries
//go:noescape
func CallPerformanceGetEntries(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Performance_GetEntries
//go:noescape
func TryPerformanceGetEntries(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_GetEntriesByType
//go:noescape
func HasPerformanceGetEntriesByType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_GetEntriesByType
//go:noescape
func PerformanceGetEntriesByTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_GetEntriesByType
//go:noescape
func CallPerformanceGetEntriesByType(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_Performance_GetEntriesByType
//go:noescape
func TryPerformanceGetEntriesByType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_GetEntriesByName
//go:noescape
func HasPerformanceGetEntriesByName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_GetEntriesByName
//go:noescape
func PerformanceGetEntriesByNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_GetEntriesByName
//go:noescape
func CallPerformanceGetEntriesByName(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	typ js.Ref)

//go:wasmimport plat/js/web try_Performance_GetEntriesByName
//go:noescape
func TryPerformanceGetEntriesByName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Performance_GetEntriesByName1
//go:noescape
func HasPerformanceGetEntriesByName1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Performance_GetEntriesByName1
//go:noescape
func PerformanceGetEntriesByName1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Performance_GetEntriesByName1
//go:noescape
func CallPerformanceGetEntriesByName1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Performance_GetEntriesByName1
//go:noescape
func TryPerformanceGetEntriesByName1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_Storage_Length
//go:noescape
func GetStorageLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Storage_Key
//go:noescape
func HasStorageKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Storage_Key
//go:noescape
func StorageKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Storage_Key
//go:noescape
func CallStorageKey(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_Storage_Key
//go:noescape
func TryStorageKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Storage_GetItem
//go:noescape
func HasStorageGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Storage_GetItem
//go:noescape
func StorageGetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Storage_GetItem
//go:noescape
func CallStorageGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_Storage_GetItem
//go:noescape
func TryStorageGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Storage_SetItem
//go:noescape
func HasStorageSetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Storage_SetItem
//go:noescape
func StorageSetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Storage_SetItem
//go:noescape
func CallStorageSetItem(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_Storage_SetItem
//go:noescape
func TryStorageSetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Storage_RemoveItem
//go:noescape
func HasStorageRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Storage_RemoveItem
//go:noescape
func StorageRemoveItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Storage_RemoveItem
//go:noescape
func CallStorageRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_Storage_RemoveItem
//go:noescape
func TryStorageRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Storage_Clear
//go:noescape
func HasStorageClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Storage_Clear
//go:noescape
func StorageClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Storage_Clear
//go:noescape
func CallStorageClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Storage_Clear
//go:noescape
func TryStorageClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Window
//go:noescape
func GetWindowWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Self
//go:noescape
func GetWindowSelf(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Document
//go:noescape
func GetWindowDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Name
//go:noescape
func GetWindowName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Window_Name
//go:noescape
func SetWindowName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Window_Location
//go:noescape
func GetWindowLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_History
//go:noescape
func GetWindowHistory(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Navigation
//go:noescape
func GetWindowNavigation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_CustomElements
//go:noescape
func GetWindowCustomElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Locationbar
//go:noescape
func GetWindowLocationbar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Menubar
//go:noescape
func GetWindowMenubar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Personalbar
//go:noescape
func GetWindowPersonalbar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Scrollbars
//go:noescape
func GetWindowScrollbars(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Statusbar
//go:noescape
func GetWindowStatusbar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Toolbar
//go:noescape
func GetWindowToolbar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Status
//go:noescape
func GetWindowStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Window_Status
//go:noescape
func SetWindowStatus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Window_Closed
//go:noescape
func GetWindowClosed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Frames
//go:noescape
func GetWindowFrames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Length
//go:noescape
func GetWindowLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Top
//go:noescape
func GetWindowTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Opener
//go:noescape
func GetWindowOpener(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Window_Opener
//go:noescape
func SetWindowOpener(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Window_Parent
//go:noescape
func GetWindowParent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_FrameElement
//go:noescape
func GetWindowFrameElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Navigator
//go:noescape
func GetWindowNavigator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ClientInformation
//go:noescape
func GetWindowClientInformation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_OriginAgentCluster
//go:noescape
func GetWindowOriginAgentCluster(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Orientation
//go:noescape
func GetWindowOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_LaunchQueue
//go:noescape
func GetWindowLaunchQueue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Fence
//go:noescape
func GetWindowFence(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_PortalHost
//go:noescape
func GetWindowPortalHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Screen
//go:noescape
func GetWindowScreen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_VisualViewport
//go:noescape
func GetWindowVisualViewport(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_InnerWidth
//go:noescape
func GetWindowInnerWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_InnerHeight
//go:noescape
func GetWindowInnerHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScrollX
//go:noescape
func GetWindowScrollX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_PageXOffset
//go:noescape
func GetWindowPageXOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScrollY
//go:noescape
func GetWindowScrollY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_PageYOffset
//go:noescape
func GetWindowPageYOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScreenX
//go:noescape
func GetWindowScreenX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScreenLeft
//go:noescape
func GetWindowScreenLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScreenY
//go:noescape
func GetWindowScreenY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_ScreenTop
//go:noescape
func GetWindowScreenTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_OuterWidth
//go:noescape
func GetWindowOuterWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_OuterHeight
//go:noescape
func GetWindowOuterHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_DevicePixelRatio
//go:noescape
func GetWindowDevicePixelRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_SharedStorage
//go:noescape
func GetWindowSharedStorage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Event
//go:noescape
func GetWindowEvent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_CookieStore
//go:noescape
func GetWindowCookieStore(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_DocumentPictureInPicture
//go:noescape
func GetWindowDocumentPictureInPicture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_External
//go:noescape
func GetWindowExternal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_SpeechSynthesis
//go:noescape
func GetWindowSpeechSynthesis(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Origin
//go:noescape
func GetWindowOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_IsSecureContext
//go:noescape
func GetWindowIsSecureContext(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_CrossOriginIsolated
//go:noescape
func GetWindowCrossOriginIsolated(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Scheduler
//go:noescape
func GetWindowScheduler(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_TrustedTypes
//go:noescape
func GetWindowTrustedTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Caches
//go:noescape
func GetWindowCaches(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Crypto
//go:noescape
func GetWindowCrypto(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_IndexedDB
//go:noescape
func GetWindowIndexedDB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_Performance
//go:noescape
func GetWindowPerformance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_SessionStorage
//go:noescape
func GetWindowSessionStorage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Window_LocalStorage
//go:noescape
func GetWindowLocalStorage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Close
//go:noescape
func HasWindowClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Close
//go:noescape
func WindowCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Close
//go:noescape
func CallWindowClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Close
//go:noescape
func TryWindowClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Stop
//go:noescape
func HasWindowStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Stop
//go:noescape
func WindowStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Stop
//go:noescape
func CallWindowStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Stop
//go:noescape
func TryWindowStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Focus
//go:noescape
func HasWindowFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Focus
//go:noescape
func WindowFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Focus
//go:noescape
func CallWindowFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Focus
//go:noescape
func TryWindowFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Blur
//go:noescape
func HasWindowBlur(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Blur
//go:noescape
func WindowBlurFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Blur
//go:noescape
func CallWindowBlur(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Blur
//go:noescape
func TryWindowBlur(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Open
//go:noescape
func HasWindowOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Open
//go:noescape
func WindowOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Open
//go:noescape
func CallWindowOpen(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	target js.Ref,
	features js.Ref)

//go:wasmimport plat/js/web try_Window_Open
//go:noescape
func TryWindowOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	target js.Ref,
	features js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Open1
//go:noescape
func HasWindowOpen1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Open1
//go:noescape
func WindowOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Open1
//go:noescape
func CallWindowOpen1(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	target js.Ref)

//go:wasmimport plat/js/web try_Window_Open1
//go:noescape
func TryWindowOpen1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	target js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Open2
//go:noescape
func HasWindowOpen2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Open2
//go:noescape
func WindowOpen2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Open2
//go:noescape
func CallWindowOpen2(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Window_Open2
//go:noescape
func TryWindowOpen2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Open3
//go:noescape
func HasWindowOpen3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Open3
//go:noescape
func WindowOpen3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Open3
//go:noescape
func CallWindowOpen3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Open3
//go:noescape
func TryWindowOpen3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Get
//go:noescape
func HasWindowGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Get
//go:noescape
func WindowGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Get
//go:noescape
func CallWindowGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Window_Get
//go:noescape
func TryWindowGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Alert
//go:noescape
func HasWindowAlert(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Alert
//go:noescape
func WindowAlertFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Alert
//go:noescape
func CallWindowAlert(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Alert
//go:noescape
func TryWindowAlert(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Alert1
//go:noescape
func HasWindowAlert1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Alert1
//go:noescape
func WindowAlert1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Alert1
//go:noescape
func CallWindowAlert1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_Window_Alert1
//go:noescape
func TryWindowAlert1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Confirm
//go:noescape
func HasWindowConfirm(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Confirm
//go:noescape
func WindowConfirmFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Confirm
//go:noescape
func CallWindowConfirm(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_Window_Confirm
//go:noescape
func TryWindowConfirm(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Confirm1
//go:noescape
func HasWindowConfirm1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Confirm1
//go:noescape
func WindowConfirm1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Confirm1
//go:noescape
func CallWindowConfirm1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Confirm1
//go:noescape
func TryWindowConfirm1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Prompt
//go:noescape
func HasWindowPrompt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Prompt
//go:noescape
func WindowPromptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Prompt
//go:noescape
func CallWindowPrompt(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	def js.Ref)

//go:wasmimport plat/js/web try_Window_Prompt
//go:noescape
func TryWindowPrompt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	def js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Prompt1
//go:noescape
func HasWindowPrompt1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Prompt1
//go:noescape
func WindowPrompt1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Prompt1
//go:noescape
func CallWindowPrompt1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_Window_Prompt1
//go:noescape
func TryWindowPrompt1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Prompt2
//go:noescape
func HasWindowPrompt2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Prompt2
//go:noescape
func WindowPrompt2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Prompt2
//go:noescape
func CallWindowPrompt2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Prompt2
//go:noescape
func TryWindowPrompt2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Print
//go:noescape
func HasWindowPrint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Print
//go:noescape
func WindowPrintFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Print
//go:noescape
func CallWindowPrint(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Print
//go:noescape
func TryWindowPrint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_PostMessage
//go:noescape
func HasWindowPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_PostMessage
//go:noescape
func WindowPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_PostMessage
//go:noescape
func CallWindowPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref,
	transfer js.Ref)

//go:wasmimport plat/js/web try_Window_PostMessage
//go:noescape
func TryWindowPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref,
	transfer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_PostMessage1
//go:noescape
func HasWindowPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_PostMessage1
//go:noescape
func WindowPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_PostMessage1
//go:noescape
func CallWindowPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref)

//go:wasmimport plat/js/web try_Window_PostMessage1
//go:noescape
func TryWindowPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	targetOrigin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_PostMessage2
//go:noescape
func HasWindowPostMessage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_PostMessage2
//go:noescape
func WindowPostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_PostMessage2
//go:noescape
func CallWindowPostMessage2(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_PostMessage2
//go:noescape
func TryWindowPostMessage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_PostMessage3
//go:noescape
func HasWindowPostMessage3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_PostMessage3
//go:noescape
func WindowPostMessage3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_PostMessage3
//go:noescape
func CallWindowPostMessage3(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_Window_PostMessage3
//go:noescape
func TryWindowPostMessage3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_GetSelection
//go:noescape
func HasWindowGetSelection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_GetSelection
//go:noescape
func WindowGetSelectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_GetSelection
//go:noescape
func CallWindowGetSelection(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_GetSelection
//go:noescape
func TryWindowGetSelection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Navigate
//go:noescape
func HasWindowNavigate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Navigate
//go:noescape
func WindowNavigateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Navigate
//go:noescape
func CallWindowNavigate(
	this js.Ref, retPtr unsafe.Pointer,
	dir uint32)

//go:wasmimport plat/js/web try_Window_Navigate
//go:noescape
func TryWindowNavigate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dir uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_GetComputedStyle
//go:noescape
func HasWindowGetComputedStyle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_GetComputedStyle
//go:noescape
func WindowGetComputedStyleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_GetComputedStyle
//go:noescape
func CallWindowGetComputedStyle(
	this js.Ref, retPtr unsafe.Pointer,
	elt js.Ref,
	pseudoElt js.Ref)

//go:wasmimport plat/js/web try_Window_GetComputedStyle
//go:noescape
func TryWindowGetComputedStyle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	elt js.Ref,
	pseudoElt js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_GetComputedStyle1
//go:noescape
func HasWindowGetComputedStyle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_GetComputedStyle1
//go:noescape
func WindowGetComputedStyle1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_GetComputedStyle1
//go:noescape
func CallWindowGetComputedStyle1(
	this js.Ref, retPtr unsafe.Pointer,
	elt js.Ref)

//go:wasmimport plat/js/web try_Window_GetComputedStyle1
//go:noescape
func TryWindowGetComputedStyle1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	elt js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_MatchMedia
//go:noescape
func HasWindowMatchMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_MatchMedia
//go:noescape
func WindowMatchMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_MatchMedia
//go:noescape
func CallWindowMatchMedia(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_Window_MatchMedia
//go:noescape
func TryWindowMatchMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_MoveTo
//go:noescape
func HasWindowMoveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_MoveTo
//go:noescape
func WindowMoveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_MoveTo
//go:noescape
func CallWindowMoveTo(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32)

//go:wasmimport plat/js/web try_Window_MoveTo
//go:noescape
func TryWindowMoveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_MoveBy
//go:noescape
func HasWindowMoveBy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_MoveBy
//go:noescape
func WindowMoveByFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_MoveBy
//go:noescape
func CallWindowMoveBy(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32)

//go:wasmimport plat/js/web try_Window_MoveBy
//go:noescape
func TryWindowMoveBy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ResizeTo
//go:noescape
func HasWindowResizeTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ResizeTo
//go:noescape
func WindowResizeToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ResizeTo
//go:noescape
func CallWindowResizeTo(
	this js.Ref, retPtr unsafe.Pointer,
	width int32,
	height int32)

//go:wasmimport plat/js/web try_Window_ResizeTo
//go:noescape
func TryWindowResizeTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	width int32,
	height int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ResizeBy
//go:noescape
func HasWindowResizeBy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ResizeBy
//go:noescape
func WindowResizeByFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ResizeBy
//go:noescape
func CallWindowResizeBy(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32)

//go:wasmimport plat/js/web try_Window_ResizeBy
//go:noescape
func TryWindowResizeBy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Scroll
//go:noescape
func HasWindowScroll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Scroll
//go:noescape
func WindowScrollFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Scroll
//go:noescape
func CallWindowScroll(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Scroll
//go:noescape
func TryWindowScroll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Scroll1
//go:noescape
func HasWindowScroll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Scroll1
//go:noescape
func WindowScroll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Scroll1
//go:noescape
func CallWindowScroll1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Scroll1
//go:noescape
func TryWindowScroll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Scroll2
//go:noescape
func HasWindowScroll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Scroll2
//go:noescape
func WindowScroll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Scroll2
//go:noescape
func CallWindowScroll2(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Window_Scroll2
//go:noescape
func TryWindowScroll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollTo
//go:noescape
func HasWindowScrollTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollTo
//go:noescape
func WindowScrollToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollTo
//go:noescape
func CallWindowScrollTo(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ScrollTo
//go:noescape
func TryWindowScrollTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollTo1
//go:noescape
func HasWindowScrollTo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollTo1
//go:noescape
func WindowScrollTo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollTo1
//go:noescape
func CallWindowScrollTo1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ScrollTo1
//go:noescape
func TryWindowScrollTo1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollTo2
//go:noescape
func HasWindowScrollTo2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollTo2
//go:noescape
func WindowScrollTo2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollTo2
//go:noescape
func CallWindowScrollTo2(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Window_ScrollTo2
//go:noescape
func TryWindowScrollTo2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollBy
//go:noescape
func HasWindowScrollBy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollBy
//go:noescape
func WindowScrollByFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollBy
//go:noescape
func CallWindowScrollBy(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ScrollBy
//go:noescape
func TryWindowScrollBy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollBy1
//go:noescape
func HasWindowScrollBy1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollBy1
//go:noescape
func WindowScrollBy1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollBy1
//go:noescape
func CallWindowScrollBy1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ScrollBy1
//go:noescape
func TryWindowScrollBy1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ScrollBy2
//go:noescape
func HasWindowScrollBy2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ScrollBy2
//go:noescape
func WindowScrollBy2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ScrollBy2
//go:noescape
func CallWindowScrollBy2(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Window_ScrollBy2
//go:noescape
func TryWindowScrollBy2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_RequestIdleCallback
//go:noescape
func HasWindowRequestIdleCallback(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_RequestIdleCallback
//go:noescape
func WindowRequestIdleCallbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_RequestIdleCallback
//go:noescape
func CallWindowRequestIdleCallback(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_RequestIdleCallback
//go:noescape
func TryWindowRequestIdleCallback(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_RequestIdleCallback1
//go:noescape
func HasWindowRequestIdleCallback1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_RequestIdleCallback1
//go:noescape
func WindowRequestIdleCallback1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_RequestIdleCallback1
//go:noescape
func CallWindowRequestIdleCallback1(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_Window_RequestIdleCallback1
//go:noescape
func TryWindowRequestIdleCallback1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CancelIdleCallback
//go:noescape
func HasWindowCancelIdleCallback(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CancelIdleCallback
//go:noescape
func WindowCancelIdleCallbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CancelIdleCallback
//go:noescape
func CallWindowCancelIdleCallback(
	this js.Ref, retPtr unsafe.Pointer,
	handle uint32)

//go:wasmimport plat/js/web try_Window_CancelIdleCallback
//go:noescape
func TryWindowCancelIdleCallback(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowOpenFilePicker
//go:noescape
func HasWindowShowOpenFilePicker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowOpenFilePicker
//go:noescape
func WindowShowOpenFilePickerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowOpenFilePicker
//go:noescape
func CallWindowShowOpenFilePicker(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowOpenFilePicker
//go:noescape
func TryWindowShowOpenFilePicker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowOpenFilePicker1
//go:noescape
func HasWindowShowOpenFilePicker1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowOpenFilePicker1
//go:noescape
func WindowShowOpenFilePicker1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowOpenFilePicker1
//go:noescape
func CallWindowShowOpenFilePicker1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowOpenFilePicker1
//go:noescape
func TryWindowShowOpenFilePicker1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowSaveFilePicker
//go:noescape
func HasWindowShowSaveFilePicker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowSaveFilePicker
//go:noescape
func WindowShowSaveFilePickerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowSaveFilePicker
//go:noescape
func CallWindowShowSaveFilePicker(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowSaveFilePicker
//go:noescape
func TryWindowShowSaveFilePicker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowSaveFilePicker1
//go:noescape
func HasWindowShowSaveFilePicker1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowSaveFilePicker1
//go:noescape
func WindowShowSaveFilePicker1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowSaveFilePicker1
//go:noescape
func CallWindowShowSaveFilePicker1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowSaveFilePicker1
//go:noescape
func TryWindowShowSaveFilePicker1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowDirectoryPicker
//go:noescape
func HasWindowShowDirectoryPicker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowDirectoryPicker
//go:noescape
func WindowShowDirectoryPickerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowDirectoryPicker
//go:noescape
func CallWindowShowDirectoryPicker(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowDirectoryPicker
//go:noescape
func TryWindowShowDirectoryPicker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ShowDirectoryPicker1
//go:noescape
func HasWindowShowDirectoryPicker1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ShowDirectoryPicker1
//go:noescape
func WindowShowDirectoryPicker1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ShowDirectoryPicker1
//go:noescape
func CallWindowShowDirectoryPicker1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ShowDirectoryPicker1
//go:noescape
func TryWindowShowDirectoryPicker1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_QueryLocalFonts
//go:noescape
func HasWindowQueryLocalFonts(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_QueryLocalFonts
//go:noescape
func WindowQueryLocalFontsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_QueryLocalFonts
//go:noescape
func CallWindowQueryLocalFonts(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_QueryLocalFonts
//go:noescape
func TryWindowQueryLocalFonts(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_QueryLocalFonts1
//go:noescape
func HasWindowQueryLocalFonts1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_QueryLocalFonts1
//go:noescape
func WindowQueryLocalFonts1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_QueryLocalFonts1
//go:noescape
func CallWindowQueryLocalFonts1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_QueryLocalFonts1
//go:noescape
func TryWindowQueryLocalFonts1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_GetScreenDetails
//go:noescape
func HasWindowGetScreenDetails(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_GetScreenDetails
//go:noescape
func WindowGetScreenDetailsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_GetScreenDetails
//go:noescape
func CallWindowGetScreenDetails(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_GetScreenDetails
//go:noescape
func TryWindowGetScreenDetails(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_GetDigitalGoodsService
//go:noescape
func HasWindowGetDigitalGoodsService(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_GetDigitalGoodsService
//go:noescape
func WindowGetDigitalGoodsServiceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_GetDigitalGoodsService
//go:noescape
func CallWindowGetDigitalGoodsService(
	this js.Ref, retPtr unsafe.Pointer,
	serviceProvider js.Ref)

//go:wasmimport plat/js/web try_Window_GetDigitalGoodsService
//go:noescape
func TryWindowGetDigitalGoodsService(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serviceProvider js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CaptureEvents
//go:noescape
func HasWindowCaptureEvents(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CaptureEvents
//go:noescape
func WindowCaptureEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CaptureEvents
//go:noescape
func CallWindowCaptureEvents(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_CaptureEvents
//go:noescape
func TryWindowCaptureEvents(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ReleaseEvents
//go:noescape
func HasWindowReleaseEvents(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ReleaseEvents
//go:noescape
func WindowReleaseEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ReleaseEvents
//go:noescape
func CallWindowReleaseEvents(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ReleaseEvents
//go:noescape
func TryWindowReleaseEvents(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ReportError
//go:noescape
func HasWindowReportError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ReportError
//go:noescape
func WindowReportErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ReportError
//go:noescape
func CallWindowReportError(
	this js.Ref, retPtr unsafe.Pointer,
	e js.Ref)

//go:wasmimport plat/js/web try_Window_ReportError
//go:noescape
func TryWindowReportError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	e js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Btoa
//go:noescape
func HasWindowBtoa(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Btoa
//go:noescape
func WindowBtoaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Btoa
//go:noescape
func CallWindowBtoa(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Window_Btoa
//go:noescape
func TryWindowBtoa(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Atob
//go:noescape
func HasWindowAtob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Atob
//go:noescape
func WindowAtobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Atob
//go:noescape
func CallWindowAtob(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Window_Atob
//go:noescape
func TryWindowAtob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_SetTimeout
//go:noescape
func HasWindowSetTimeout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_SetTimeout
//go:noescape
func WindowSetTimeoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_SetTimeout
//go:noescape
func CallWindowSetTimeout(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_Window_SetTimeout
//go:noescape
func TryWindowSetTimeout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_SetTimeout1
//go:noescape
func HasWindowSetTimeout1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_SetTimeout1
//go:noescape
func WindowSetTimeout1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_SetTimeout1
//go:noescape
func CallWindowSetTimeout1(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref)

//go:wasmimport plat/js/web try_Window_SetTimeout1
//go:noescape
func TryWindowSetTimeout1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ClearTimeout
//go:noescape
func HasWindowClearTimeout(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ClearTimeout
//go:noescape
func WindowClearTimeoutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ClearTimeout
//go:noescape
func CallWindowClearTimeout(
	this js.Ref, retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/web try_Window_ClearTimeout
//go:noescape
func TryWindowClearTimeout(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ClearTimeout1
//go:noescape
func HasWindowClearTimeout1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ClearTimeout1
//go:noescape
func WindowClearTimeout1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ClearTimeout1
//go:noescape
func CallWindowClearTimeout1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ClearTimeout1
//go:noescape
func TryWindowClearTimeout1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_SetInterval
//go:noescape
func HasWindowSetInterval(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_SetInterval
//go:noescape
func WindowSetIntervalFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_SetInterval
//go:noescape
func CallWindowSetInterval(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32)

//go:wasmimport plat/js/web try_Window_SetInterval
//go:noescape
func TryWindowSetInterval(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref,
	timeout int32,
	argumentsPtr unsafe.Pointer,
	argumentsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_SetInterval1
//go:noescape
func HasWindowSetInterval1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_SetInterval1
//go:noescape
func WindowSetInterval1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_SetInterval1
//go:noescape
func CallWindowSetInterval1(
	this js.Ref, retPtr unsafe.Pointer,
	handler js.Ref)

//go:wasmimport plat/js/web try_Window_SetInterval1
//go:noescape
func TryWindowSetInterval1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ClearInterval
//go:noescape
func HasWindowClearInterval(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ClearInterval
//go:noescape
func WindowClearIntervalFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ClearInterval
//go:noescape
func CallWindowClearInterval(
	this js.Ref, retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/web try_Window_ClearInterval
//go:noescape
func TryWindowClearInterval(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_ClearInterval1
//go:noescape
func HasWindowClearInterval1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_ClearInterval1
//go:noescape
func WindowClearInterval1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_ClearInterval1
//go:noescape
func CallWindowClearInterval1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_ClearInterval1
//go:noescape
func TryWindowClearInterval1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_QueueMicrotask
//go:noescape
func HasWindowQueueMicrotask(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_QueueMicrotask
//go:noescape
func WindowQueueMicrotaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_QueueMicrotask
//go:noescape
func CallWindowQueueMicrotask(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_Window_QueueMicrotask
//go:noescape
func TryWindowQueueMicrotask(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CreateImageBitmap
//go:noescape
func HasWindowCreateImageBitmap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CreateImageBitmap
//go:noescape
func WindowCreateImageBitmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CreateImageBitmap
//go:noescape
func CallWindowCreateImageBitmap(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_CreateImageBitmap
//go:noescape
func TryWindowCreateImageBitmap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CreateImageBitmap1
//go:noescape
func HasWindowCreateImageBitmap1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CreateImageBitmap1
//go:noescape
func WindowCreateImageBitmap1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CreateImageBitmap1
//go:noescape
func CallWindowCreateImageBitmap1(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref)

//go:wasmimport plat/js/web try_Window_CreateImageBitmap1
//go:noescape
func TryWindowCreateImageBitmap1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CreateImageBitmap2
//go:noescape
func HasWindowCreateImageBitmap2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CreateImageBitmap2
//go:noescape
func WindowCreateImageBitmap2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CreateImageBitmap2
//go:noescape
func CallWindowCreateImageBitmap2(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_CreateImageBitmap2
//go:noescape
func TryWindowCreateImageBitmap2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CreateImageBitmap3
//go:noescape
func HasWindowCreateImageBitmap3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CreateImageBitmap3
//go:noescape
func WindowCreateImageBitmap3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CreateImageBitmap3
//go:noescape
func CallWindowCreateImageBitmap3(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_Window_CreateImageBitmap3
//go:noescape
func TryWindowCreateImageBitmap3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	sx int32,
	sy int32,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_StructuredClone
//go:noescape
func HasWindowStructuredClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_StructuredClone
//go:noescape
func WindowStructuredCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_StructuredClone
//go:noescape
func CallWindowStructuredClone(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_StructuredClone
//go:noescape
func TryWindowStructuredClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_StructuredClone1
//go:noescape
func HasWindowStructuredClone1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_StructuredClone1
//go:noescape
func WindowStructuredClone1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_StructuredClone1
//go:noescape
func CallWindowStructuredClone1(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_Window_StructuredClone1
//go:noescape
func TryWindowStructuredClone1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Fetch
//go:noescape
func HasWindowFetch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Fetch
//go:noescape
func WindowFetchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Fetch
//go:noescape
func CallWindowFetch(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_Window_Fetch
//go:noescape
func TryWindowFetch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_Fetch1
//go:noescape
func HasWindowFetch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_Fetch1
//go:noescape
func WindowFetch1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_Fetch1
//go:noescape
func CallWindowFetch1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_Window_Fetch1
//go:noescape
func TryWindowFetch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_RequestAnimationFrame
//go:noescape
func HasWindowRequestAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_RequestAnimationFrame
//go:noescape
func WindowRequestAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_RequestAnimationFrame
//go:noescape
func CallWindowRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_Window_RequestAnimationFrame
//go:noescape
func TryWindowRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Window_CancelAnimationFrame
//go:noescape
func HasWindowCancelAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Window_CancelAnimationFrame
//go:noescape
func WindowCancelAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Window_CancelAnimationFrame
//go:noescape
func CallWindowCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	handle uint32)

//go:wasmimport plat/js/web try_Window_CancelAnimationFrame
//go:noescape
func TryWindowCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_CompositionEventInit
//go:noescape
func CompositionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CompositionEventInit
//go:noescape
func CompositionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CompositionEvent_CompositionEvent
//go:noescape
func NewCompositionEventByCompositionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CompositionEvent_CompositionEvent1
//go:noescape
func NewCompositionEventByCompositionEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CompositionEvent_Data
//go:noescape
func GetCompositionEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CompositionEvent_InitCompositionEvent
//go:noescape
func HasCompositionEventInitCompositionEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CompositionEvent_InitCompositionEvent
//go:noescape
func CompositionEventInitCompositionEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CompositionEvent_InitCompositionEvent
//go:noescape
func CallCompositionEventInitCompositionEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	dataArg js.Ref)

//go:wasmimport plat/js/web try_CompositionEvent_InitCompositionEvent
//go:noescape
func TryCompositionEventInitCompositionEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	dataArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CompositionEvent_InitCompositionEvent1
//go:noescape
func HasCompositionEventInitCompositionEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CompositionEvent_InitCompositionEvent1
//go:noescape
func CompositionEventInitCompositionEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CompositionEvent_InitCompositionEvent1
//go:noescape
func CallCompositionEventInitCompositionEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref)

//go:wasmimport plat/js/web try_CompositionEvent_InitCompositionEvent1
//go:noescape
func TryCompositionEventInitCompositionEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CompositionEvent_InitCompositionEvent2
//go:noescape
func HasCompositionEventInitCompositionEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CompositionEvent_InitCompositionEvent2
//go:noescape
func CompositionEventInitCompositionEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CompositionEvent_InitCompositionEvent2
//go:noescape
func CallCompositionEventInitCompositionEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref)

//go:wasmimport plat/js/web try_CompositionEvent_InitCompositionEvent2
//go:noescape
func TryCompositionEventInitCompositionEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CompositionEvent_InitCompositionEvent3
//go:noescape
func HasCompositionEventInitCompositionEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CompositionEvent_InitCompositionEvent3
//go:noescape
func CompositionEventInitCompositionEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CompositionEvent_InitCompositionEvent3
//go:noescape
func CallCompositionEventInitCompositionEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref)

//go:wasmimport plat/js/web try_CompositionEvent_InitCompositionEvent3
//go:noescape
func TryCompositionEventInitCompositionEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CompositionEvent_InitCompositionEvent4
//go:noescape
func HasCompositionEventInitCompositionEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CompositionEvent_InitCompositionEvent4
//go:noescape
func CompositionEventInitCompositionEvent4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CompositionEvent_InitCompositionEvent4
//go:noescape
func CallCompositionEventInitCompositionEvent4(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref)

//go:wasmimport plat/js/web try_CompositionEvent_InitCompositionEvent4
//go:noescape
func TryCompositionEventInitCompositionEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_CompressionFormat
//go:noescape
func ConstOfCompressionFormat(str js.Ref) uint32

//go:wasmimport plat/js/web new_CompressionStream_CompressionStream
//go:noescape
func NewCompressionStreamByCompressionStream(
	format uint32) js.Ref

//go:wasmimport plat/js/web get_CompressionStream_Readable
//go:noescape
func GetCompressionStreamReadable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CompressionStream_Writable
//go:noescape
func GetCompressionStreamWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ContentIndexEventInit
//go:noescape
func ContentIndexEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ContentIndexEventInit
//go:noescape
func ContentIndexEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ContentIndexEvent_ContentIndexEvent
//go:noescape
func NewContentIndexEventByContentIndexEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_ContentIndexEvent_Id
//go:noescape
func GetContentIndexEventId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ContentVisibilityAutoStateChangeEventInit
//go:noescape
func ContentVisibilityAutoStateChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ContentVisibilityAutoStateChangeEventInit
//go:noescape
func ContentVisibilityAutoStateChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ContentVisibilityAutoStateChangeEvent_ContentVisibilityAutoStateChangeEvent
//go:noescape
func NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ContentVisibilityAutoStateChangeEvent_ContentVisibilityAutoStateChangeEvent1
//go:noescape
func NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ContentVisibilityAutoStateChangeEvent_Skipped
//go:noescape
func GetContentVisibilityAutoStateChangeEventSkipped(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CookieChangeEventInit
//go:noescape
func CookieChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieChangeEventInit
//go:noescape
func CookieChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CookieChangeEvent_CookieChangeEvent
//go:noescape
func NewCookieChangeEventByCookieChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CookieChangeEvent_CookieChangeEvent1
//go:noescape
func NewCookieChangeEventByCookieChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CookieChangeEvent_Changed
//go:noescape
func GetCookieChangeEventChanged(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CookieChangeEvent_Deleted
//go:noescape
func GetCookieChangeEventDeleted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CountQueuingStrategy_CountQueuingStrategy
//go:noescape
func NewCountQueuingStrategyByCountQueuingStrategy(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CountQueuingStrategy_HighWaterMark
//go:noescape
func GetCountQueuingStrategyHighWaterMark(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CountQueuingStrategy_Size
//go:noescape
func GetCountQueuingStrategySize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CrashReportBody_Reason
//go:noescape
func GetCrashReportBodyReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CrashReportBody_ToJSON
//go:noescape
func HasCrashReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CrashReportBody_ToJSON
//go:noescape
func CrashReportBodyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CrashReportBody_ToJSON
//go:noescape
func CallCrashReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CrashReportBody_ToJSON
//go:noescape
func TryCrashReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CredentialData
//go:noescape
func CredentialDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CredentialData
//go:noescape
func CredentialDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CryptoKeyPair
//go:noescape
func CryptoKeyPairJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CryptoKeyPair
//go:noescape
func CryptoKeyPairJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_CursorCaptureConstraint
//go:noescape
func ConstOfCursorCaptureConstraint(str js.Ref) uint32

//go:wasmimport plat/js/web store_CustomEventInit
//go:noescape
func CustomEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CustomEventInit
//go:noescape
func CustomEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CustomEvent_CustomEvent
//go:noescape
func NewCustomEventByCustomEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CustomEvent_CustomEvent1
//go:noescape
func NewCustomEventByCustomEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CustomEvent_Detail
//go:noescape
func GetCustomEventDetail(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomEvent_InitCustomEvent
//go:noescape
func HasCustomEventInitCustomEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomEvent_InitCustomEvent
//go:noescape
func CustomEventInitCustomEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CustomEvent_InitCustomEvent
//go:noescape
func CallCustomEventInitCustomEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	detail js.Ref)

//go:wasmimport plat/js/web try_CustomEvent_InitCustomEvent
//go:noescape
func TryCustomEventInitCustomEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref,
	detail js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomEvent_InitCustomEvent1
//go:noescape
func HasCustomEventInitCustomEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomEvent_InitCustomEvent1
//go:noescape
func CustomEventInitCustomEvent1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CustomEvent_InitCustomEvent1
//go:noescape
func CallCustomEventInitCustomEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref)

//go:wasmimport plat/js/web try_CustomEvent_InitCustomEvent1
//go:noescape
func TryCustomEventInitCustomEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref,
	cancelable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomEvent_InitCustomEvent2
//go:noescape
func HasCustomEventInitCustomEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomEvent_InitCustomEvent2
//go:noescape
func CustomEventInitCustomEvent2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CustomEvent_InitCustomEvent2
//go:noescape
func CallCustomEventInitCustomEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref)

//go:wasmimport plat/js/web try_CustomEvent_InitCustomEvent2
//go:noescape
func TryCustomEventInitCustomEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	bubbles js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomEvent_InitCustomEvent3
//go:noescape
func HasCustomEventInitCustomEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomEvent_InitCustomEvent3
//go:noescape
func CustomEventInitCustomEvent3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CustomEvent_InitCustomEvent3
//go:noescape
func CallCustomEventInitCustomEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_CustomEvent_InitCustomEvent3
//go:noescape
func TryCustomEventInitCustomEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_DOMParserSupportedType
//go:noescape
func ConstOfDOMParserSupportedType(str js.Ref) uint32

//go:wasmimport plat/js/web has_DOMParser_ParseFromString
//go:noescape
func HasDOMParserParseFromString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMParser_ParseFromString
//go:noescape
func DOMParserParseFromStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMParser_ParseFromString
//go:noescape
func CallDOMParserParseFromString(
	this js.Ref, retPtr unsafe.Pointer,
	string js.Ref,
	typ uint32)

//go:wasmimport plat/js/web try_DOMParser_ParseFromString
//go:noescape
func TryDOMParserParseFromString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	string js.Ref,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_DataCue_DataCue
//go:noescape
func NewDataCueByDataCue(
	startTime float64,
	endTime float64,
	value js.Ref,
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web new_DataCue_DataCue1
//go:noescape
func NewDataCueByDataCue1(
	startTime float64,
	endTime float64,
	value js.Ref) js.Ref

//go:wasmimport plat/js/web get_DataCue_Value
//go:noescape
func GetDataCueValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DataCue_Value
//go:noescape
func SetDataCueValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_DataCue_Type
//go:noescape
func GetDataCueType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_DecompressionStream_DecompressionStream
//go:noescape
func NewDecompressionStreamByDecompressionStream(
	format uint32) js.Ref

//go:wasmimport plat/js/web get_DecompressionStream_Readable
//go:noescape
func GetDecompressionStreamReadable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DecompressionStream_Writable
//go:noescape
func GetDecompressionStreamWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DedicatedWorkerGlobalScope_Name
//go:noescape
func GetDedicatedWorkerGlobalScopeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_PostMessage
//go:noescape
func HasDedicatedWorkerGlobalScopePostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_PostMessage
//go:noescape
func DedicatedWorkerGlobalScopePostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_PostMessage
//go:noescape
func CallDedicatedWorkerGlobalScopePostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_PostMessage
//go:noescape
func TryDedicatedWorkerGlobalScopePostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_PostMessage1
//go:noescape
func HasDedicatedWorkerGlobalScopePostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_PostMessage1
//go:noescape
func DedicatedWorkerGlobalScopePostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_PostMessage1
//go:noescape
func CallDedicatedWorkerGlobalScopePostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_PostMessage1
//go:noescape
func TryDedicatedWorkerGlobalScopePostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_PostMessage2
//go:noescape
func HasDedicatedWorkerGlobalScopePostMessage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_PostMessage2
//go:noescape
func DedicatedWorkerGlobalScopePostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_PostMessage2
//go:noescape
func CallDedicatedWorkerGlobalScopePostMessage2(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_PostMessage2
//go:noescape
func TryDedicatedWorkerGlobalScopePostMessage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_Close
//go:noescape
func HasDedicatedWorkerGlobalScopeClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_Close
//go:noescape
func DedicatedWorkerGlobalScopeCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_Close
//go:noescape
func CallDedicatedWorkerGlobalScopeClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_Close
//go:noescape
func TryDedicatedWorkerGlobalScopeClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_RequestAnimationFrame
//go:noescape
func HasDedicatedWorkerGlobalScopeRequestAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_RequestAnimationFrame
//go:noescape
func DedicatedWorkerGlobalScopeRequestAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_RequestAnimationFrame
//go:noescape
func CallDedicatedWorkerGlobalScopeRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_RequestAnimationFrame
//go:noescape
func TryDedicatedWorkerGlobalScopeRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DedicatedWorkerGlobalScope_CancelAnimationFrame
//go:noescape
func HasDedicatedWorkerGlobalScopeCancelAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DedicatedWorkerGlobalScope_CancelAnimationFrame
//go:noescape
func DedicatedWorkerGlobalScopeCancelAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DedicatedWorkerGlobalScope_CancelAnimationFrame
//go:noescape
func CallDedicatedWorkerGlobalScopeCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	handle uint32)

//go:wasmimport plat/js/web try_DedicatedWorkerGlobalScope_CancelAnimationFrame
//go:noescape
func TryDedicatedWorkerGlobalScopeCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_Id
//go:noescape
func GetDeprecationReportBodyId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_AnticipatedRemoval
//go:noescape
func GetDeprecationReportBodyAnticipatedRemoval(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_Message
//go:noescape
func GetDeprecationReportBodyMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_SourceFile
//go:noescape
func GetDeprecationReportBodySourceFile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_LineNumber
//go:noescape
func GetDeprecationReportBodyLineNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeprecationReportBody_ColumnNumber
//go:noescape
func GetDeprecationReportBodyColumnNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DeprecationReportBody_ToJSON
//go:noescape
func HasDeprecationReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DeprecationReportBody_ToJSON
//go:noescape
func DeprecationReportBodyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DeprecationReportBody_ToJSON
//go:noescape
func CallDeprecationReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DeprecationReportBody_ToJSON
//go:noescape
func TryDeprecationReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_LandmarkType
//go:noescape
func ConstOfLandmarkType(str js.Ref) uint32

//go:wasmimport plat/js/web store_Landmark
//go:noescape
func LandmarkJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Landmark
//go:noescape
func LandmarkJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DetectedFace
//go:noescape
func DetectedFaceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DetectedFace
//go:noescape
func DetectedFaceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DetectedText
//go:noescape
func DetectedTextJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DetectedText
//go:noescape
func DetectedTextJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DeviceMotionEventAccelerationInit
//go:noescape
func DeviceMotionEventAccelerationInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DeviceMotionEventAccelerationInit
//go:noescape
func DeviceMotionEventAccelerationInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DeviceMotionEventRotationRateInit
//go:noescape
func DeviceMotionEventRotationRateInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DeviceMotionEventRotationRateInit
//go:noescape
func DeviceMotionEventRotationRateInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DeviceMotionEventInit
//go:noescape
func DeviceMotionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DeviceMotionEventInit
//go:noescape
func DeviceMotionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_DeviceMotionEventAcceleration_X
//go:noescape
func GetDeviceMotionEventAccelerationX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEventAcceleration_Y
//go:noescape
func GetDeviceMotionEventAccelerationY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEventAcceleration_Z
//go:noescape
func GetDeviceMotionEventAccelerationZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEventRotationRate_Alpha
//go:noescape
func GetDeviceMotionEventRotationRateAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEventRotationRate_Beta
//go:noescape
func GetDeviceMotionEventRotationRateBeta(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEventRotationRate_Gamma
//go:noescape
func GetDeviceMotionEventRotationRateGamma(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_DeviceMotionEvent_DeviceMotionEvent
//go:noescape
func NewDeviceMotionEventByDeviceMotionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DeviceMotionEvent_DeviceMotionEvent1
//go:noescape
func NewDeviceMotionEventByDeviceMotionEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_DeviceMotionEvent_Acceleration
//go:noescape
func GetDeviceMotionEventAcceleration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEvent_AccelerationIncludingGravity
//go:noescape
func GetDeviceMotionEventAccelerationIncludingGravity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEvent_RotationRate
//go:noescape
func GetDeviceMotionEventRotationRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceMotionEvent_Interval
//go:noescape
func GetDeviceMotionEventInterval(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DeviceMotionEvent_RequestPermission
//go:noescape
func HasDeviceMotionEventRequestPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DeviceMotionEvent_RequestPermission
//go:noescape
func DeviceMotionEventRequestPermissionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DeviceMotionEvent_RequestPermission
//go:noescape
func CallDeviceMotionEventRequestPermission(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DeviceMotionEvent_RequestPermission
//go:noescape
func TryDeviceMotionEventRequestPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

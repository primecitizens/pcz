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

//go:wasmimport plat/js/webext/metricsprivate store_HistogramBucket
//go:noescape
func HistogramBucketJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/metricsprivate load_HistogramBucket
//go:noescape
func HistogramBucketJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/metricsprivate store_Histogram
//go:noescape
func HistogramJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/metricsprivate load_Histogram
//go:noescape
func HistogramJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/metricsprivate constof_MetricTypeType
//go:noescape
func ConstOfMetricTypeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/metricsprivate store_MetricType
//go:noescape
func MetricTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/metricsprivate load_MetricType
//go:noescape
func MetricTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/metricsprivate has_GetFieldTrial
//go:noescape
func HasFuncGetFieldTrial() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_GetFieldTrial
//go:noescape
func FuncGetFieldTrial(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_GetFieldTrial
//go:noescape
func CallGetFieldTrial(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_GetFieldTrial
//go:noescape
func TryGetFieldTrial(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_GetHistogram
//go:noescape
func HasFuncGetHistogram() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_GetHistogram
//go:noescape
func FuncGetHistogram(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_GetHistogram
//go:noescape
func CallGetHistogram(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_GetHistogram
//go:noescape
func TryGetHistogram(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_GetIsCrashReportingEnabled
//go:noescape
func HasFuncGetIsCrashReportingEnabled() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_GetIsCrashReportingEnabled
//go:noescape
func FuncGetIsCrashReportingEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_GetIsCrashReportingEnabled
//go:noescape
func CallGetIsCrashReportingEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate try_GetIsCrashReportingEnabled
//go:noescape
func TryGetIsCrashReportingEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_GetVariationParams
//go:noescape
func HasFuncGetVariationParams() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_GetVariationParams
//go:noescape
func FuncGetVariationParams(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_GetVariationParams
//go:noescape
func CallGetVariationParams(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_GetVariationParams
//go:noescape
func TryGetVariationParams(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordBoolean
//go:noescape
func HasFuncRecordBoolean() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordBoolean
//go:noescape
func FuncRecordBoolean(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordBoolean
//go:noescape
func CallRecordBoolean(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_RecordBoolean
//go:noescape
func TryRecordBoolean(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordCount
//go:noescape
func HasFuncRecordCount() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordCount
//go:noescape
func FuncRecordCount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordCount
//go:noescape
func CallRecordCount(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordCount
//go:noescape
func TryRecordCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordEnumerationValue
//go:noescape
func HasFuncRecordEnumerationValue() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordEnumerationValue
//go:noescape
func FuncRecordEnumerationValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordEnumerationValue
//go:noescape
func CallRecordEnumerationValue(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64,
	enumSize float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordEnumerationValue
//go:noescape
func TryRecordEnumerationValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64,
	enumSize float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordLongTime
//go:noescape
func HasFuncRecordLongTime() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordLongTime
//go:noescape
func FuncRecordLongTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordLongTime
//go:noescape
func CallRecordLongTime(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordLongTime
//go:noescape
func TryRecordLongTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordMediumCount
//go:noescape
func HasFuncRecordMediumCount() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordMediumCount
//go:noescape
func FuncRecordMediumCount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordMediumCount
//go:noescape
func CallRecordMediumCount(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordMediumCount
//go:noescape
func TryRecordMediumCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordMediumTime
//go:noescape
func HasFuncRecordMediumTime() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordMediumTime
//go:noescape
func FuncRecordMediumTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordMediumTime
//go:noescape
func CallRecordMediumTime(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordMediumTime
//go:noescape
func TryRecordMediumTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordPercentage
//go:noescape
func HasFuncRecordPercentage() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordPercentage
//go:noescape
func FuncRecordPercentage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordPercentage
//go:noescape
func CallRecordPercentage(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordPercentage
//go:noescape
func TryRecordPercentage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordSmallCount
//go:noescape
func HasFuncRecordSmallCount() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordSmallCount
//go:noescape
func FuncRecordSmallCount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordSmallCount
//go:noescape
func CallRecordSmallCount(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordSmallCount
//go:noescape
func TryRecordSmallCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordSparseValue
//go:noescape
func HasFuncRecordSparseValue() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordSparseValue
//go:noescape
func FuncRecordSparseValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordSparseValue
//go:noescape
func CallRecordSparseValue(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordSparseValue
//go:noescape
func TryRecordSparseValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordSparseValueWithHashMetricName
//go:noescape
func HasFuncRecordSparseValueWithHashMetricName() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordSparseValueWithHashMetricName
//go:noescape
func FuncRecordSparseValueWithHashMetricName(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordSparseValueWithHashMetricName
//go:noescape
func CallRecordSparseValueWithHashMetricName(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_RecordSparseValueWithHashMetricName
//go:noescape
func TryRecordSparseValueWithHashMetricName(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordSparseValueWithPersistentHash
//go:noescape
func HasFuncRecordSparseValueWithPersistentHash() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordSparseValueWithPersistentHash
//go:noescape
func FuncRecordSparseValueWithPersistentHash(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordSparseValueWithPersistentHash
//go:noescape
func CallRecordSparseValueWithPersistentHash(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_RecordSparseValueWithPersistentHash
//go:noescape
func TryRecordSparseValueWithPersistentHash(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordTime
//go:noescape
func HasFuncRecordTime() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordTime
//go:noescape
func FuncRecordTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordTime
//go:noescape
func CallRecordTime(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordTime
//go:noescape
func TryRecordTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordUserAction
//go:noescape
func HasFuncRecordUserAction() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordUserAction
//go:noescape
func FuncRecordUserAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordUserAction
//go:noescape
func CallRecordUserAction(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/metricsprivate try_RecordUserAction
//go:noescape
func TryRecordUserAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivate has_RecordValue
//go:noescape
func HasFuncRecordValue() js.Ref

//go:wasmimport plat/js/webext/metricsprivate func_RecordValue
//go:noescape
func FuncRecordValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivate call_RecordValue
//go:noescape
func CallRecordValue(
	retPtr unsafe.Pointer,
	metric unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/webext/metricsprivate try_RecordValue
//go:noescape
func TryRecordValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metric unsafe.Pointer,
	value float64) (ok js.Ref)

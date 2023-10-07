// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package metricsprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/metricsprivate/bindings"
)

type HistogramBucket struct {
	// Count is "HistogramBucket.count"
	//
	// Required
	Count int64
	// Max is "HistogramBucket.max"
	//
	// Required
	Max int64
	// Min is "HistogramBucket.min"
	//
	// Required
	Min int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HistogramBucket with all fields set.
func (p HistogramBucket) FromRef(ref js.Ref) HistogramBucket {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HistogramBucket in the application heap.
func (p HistogramBucket) New() js.Ref {
	return bindings.HistogramBucketJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HistogramBucket) UpdateFrom(ref js.Ref) {
	bindings.HistogramBucketJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HistogramBucket) Update(ref js.Ref) {
	bindings.HistogramBucketJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HistogramBucket) FreeMembers(recursive bool) {
}

type Histogram struct {
	// Buckets is "Histogram.buckets"
	//
	// Required
	Buckets js.Array[HistogramBucket]
	// Sum is "Histogram.sum"
	//
	// Required
	Sum float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Histogram with all fields set.
func (p Histogram) FromRef(ref js.Ref) Histogram {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Histogram in the application heap.
func (p Histogram) New() js.Ref {
	return bindings.HistogramJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Histogram) UpdateFrom(ref js.Ref) {
	bindings.HistogramJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Histogram) Update(ref js.Ref) {
	bindings.HistogramJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Histogram) FreeMembers(recursive bool) {
	js.Free(
		p.Buckets.Ref(),
	)
	p.Buckets = p.Buckets.FromRef(js.Undefined)
}

type MetricTypeType uint32

const (
	_ MetricTypeType = iota

	MetricTypeType_HISTOGRAM_LOG
	MetricTypeType_HISTOGRAM_LINEAR
)

func (MetricTypeType) FromRef(str js.Ref) MetricTypeType {
	return MetricTypeType(bindings.ConstOfMetricTypeType(str))
}

func (x MetricTypeType) String() (string, bool) {
	switch x {
	case MetricTypeType_HISTOGRAM_LOG:
		return "histogram-log", true
	case MetricTypeType_HISTOGRAM_LINEAR:
		return "histogram-linear", true
	default:
		return "", false
	}
}

type MetricType struct {
	// Buckets is "MetricType.buckets"
	//
	// Required
	Buckets int64
	// Max is "MetricType.max"
	//
	// Required
	Max int64
	// MetricName is "MetricType.metricName"
	//
	// Required
	MetricName js.String
	// Min is "MetricType.min"
	//
	// Required
	Min int64
	// Type is "MetricType.type"
	//
	// Required
	Type MetricTypeType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MetricType with all fields set.
func (p MetricType) FromRef(ref js.Ref) MetricType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MetricType in the application heap.
func (p MetricType) New() js.Ref {
	return bindings.MetricTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MetricType) UpdateFrom(ref js.Ref) {
	bindings.MetricTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MetricType) Update(ref js.Ref) {
	bindings.MetricTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MetricType) FreeMembers(recursive bool) {
	js.Free(
		p.MetricName.Ref(),
	)
	p.MetricName = p.MetricName.FromRef(js.Undefined)
}

// HasFuncGetFieldTrial returns true if the function "WEBEXT.metricsPrivate.getFieldTrial" exists.
func HasFuncGetFieldTrial() bool {
	return js.True == bindings.HasFuncGetFieldTrial()
}

// FuncGetFieldTrial returns the function "WEBEXT.metricsPrivate.getFieldTrial".
func FuncGetFieldTrial() (fn js.Func[func(name js.String) js.Promise[js.String]]) {
	bindings.FuncGetFieldTrial(
		js.Pointer(&fn),
	)
	return
}

// GetFieldTrial calls the function "WEBEXT.metricsPrivate.getFieldTrial" directly.
func GetFieldTrial(name js.String) (ret js.Promise[js.String]) {
	bindings.CallGetFieldTrial(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetFieldTrial calls the function "WEBEXT.metricsPrivate.getFieldTrial"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFieldTrial(name js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFieldTrial(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetHistogram returns true if the function "WEBEXT.metricsPrivate.getHistogram" exists.
func HasFuncGetHistogram() bool {
	return js.True == bindings.HasFuncGetHistogram()
}

// FuncGetHistogram returns the function "WEBEXT.metricsPrivate.getHistogram".
func FuncGetHistogram() (fn js.Func[func(name js.String) js.Promise[Histogram]]) {
	bindings.FuncGetHistogram(
		js.Pointer(&fn),
	)
	return
}

// GetHistogram calls the function "WEBEXT.metricsPrivate.getHistogram" directly.
func GetHistogram(name js.String) (ret js.Promise[Histogram]) {
	bindings.CallGetHistogram(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetHistogram calls the function "WEBEXT.metricsPrivate.getHistogram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetHistogram(name js.String) (ret js.Promise[Histogram], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetHistogram(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetIsCrashReportingEnabled returns true if the function "WEBEXT.metricsPrivate.getIsCrashReportingEnabled" exists.
func HasFuncGetIsCrashReportingEnabled() bool {
	return js.True == bindings.HasFuncGetIsCrashReportingEnabled()
}

// FuncGetIsCrashReportingEnabled returns the function "WEBEXT.metricsPrivate.getIsCrashReportingEnabled".
func FuncGetIsCrashReportingEnabled() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncGetIsCrashReportingEnabled(
		js.Pointer(&fn),
	)
	return
}

// GetIsCrashReportingEnabled calls the function "WEBEXT.metricsPrivate.getIsCrashReportingEnabled" directly.
func GetIsCrashReportingEnabled() (ret js.Promise[js.Boolean]) {
	bindings.CallGetIsCrashReportingEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryGetIsCrashReportingEnabled calls the function "WEBEXT.metricsPrivate.getIsCrashReportingEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIsCrashReportingEnabled() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIsCrashReportingEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVariationParams returns true if the function "WEBEXT.metricsPrivate.getVariationParams" exists.
func HasFuncGetVariationParams() bool {
	return js.True == bindings.HasFuncGetVariationParams()
}

// FuncGetVariationParams returns the function "WEBEXT.metricsPrivate.getVariationParams".
func FuncGetVariationParams() (fn js.Func[func(name js.String) js.Promise[js.String]]) {
	bindings.FuncGetVariationParams(
		js.Pointer(&fn),
	)
	return
}

// GetVariationParams calls the function "WEBEXT.metricsPrivate.getVariationParams" directly.
func GetVariationParams(name js.String) (ret js.Promise[js.String]) {
	bindings.CallGetVariationParams(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetVariationParams calls the function "WEBEXT.metricsPrivate.getVariationParams"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVariationParams(name js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVariationParams(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncRecordBoolean returns true if the function "WEBEXT.metricsPrivate.recordBoolean" exists.
func HasFuncRecordBoolean() bool {
	return js.True == bindings.HasFuncRecordBoolean()
}

// FuncRecordBoolean returns the function "WEBEXT.metricsPrivate.recordBoolean".
func FuncRecordBoolean() (fn js.Func[func(metricName js.String, value bool)]) {
	bindings.FuncRecordBoolean(
		js.Pointer(&fn),
	)
	return
}

// RecordBoolean calls the function "WEBEXT.metricsPrivate.recordBoolean" directly.
func RecordBoolean(metricName js.String, value bool) (ret js.Void) {
	bindings.CallRecordBoolean(
		js.Pointer(&ret),
		metricName.Ref(),
		js.Bool(bool(value)),
	)

	return
}

// TryRecordBoolean calls the function "WEBEXT.metricsPrivate.recordBoolean"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordBoolean(metricName js.String, value bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordBoolean(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		js.Bool(bool(value)),
	)

	return
}

// HasFuncRecordCount returns true if the function "WEBEXT.metricsPrivate.recordCount" exists.
func HasFuncRecordCount() bool {
	return js.True == bindings.HasFuncRecordCount()
}

// FuncRecordCount returns the function "WEBEXT.metricsPrivate.recordCount".
func FuncRecordCount() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordCount(
		js.Pointer(&fn),
	)
	return
}

// RecordCount calls the function "WEBEXT.metricsPrivate.recordCount" directly.
func RecordCount(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordCount(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordCount calls the function "WEBEXT.metricsPrivate.recordCount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordCount(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordCount(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordEnumerationValue returns true if the function "WEBEXT.metricsPrivate.recordEnumerationValue" exists.
func HasFuncRecordEnumerationValue() bool {
	return js.True == bindings.HasFuncRecordEnumerationValue()
}

// FuncRecordEnumerationValue returns the function "WEBEXT.metricsPrivate.recordEnumerationValue".
func FuncRecordEnumerationValue() (fn js.Func[func(metricName js.String, value int64, enumSize int64)]) {
	bindings.FuncRecordEnumerationValue(
		js.Pointer(&fn),
	)
	return
}

// RecordEnumerationValue calls the function "WEBEXT.metricsPrivate.recordEnumerationValue" directly.
func RecordEnumerationValue(metricName js.String, value int64, enumSize int64) (ret js.Void) {
	bindings.CallRecordEnumerationValue(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
		float64(enumSize),
	)

	return
}

// TryRecordEnumerationValue calls the function "WEBEXT.metricsPrivate.recordEnumerationValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordEnumerationValue(metricName js.String, value int64, enumSize int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordEnumerationValue(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
		float64(enumSize),
	)

	return
}

// HasFuncRecordLongTime returns true if the function "WEBEXT.metricsPrivate.recordLongTime" exists.
func HasFuncRecordLongTime() bool {
	return js.True == bindings.HasFuncRecordLongTime()
}

// FuncRecordLongTime returns the function "WEBEXT.metricsPrivate.recordLongTime".
func FuncRecordLongTime() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordLongTime(
		js.Pointer(&fn),
	)
	return
}

// RecordLongTime calls the function "WEBEXT.metricsPrivate.recordLongTime" directly.
func RecordLongTime(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordLongTime(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordLongTime calls the function "WEBEXT.metricsPrivate.recordLongTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordLongTime(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordLongTime(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordMediumCount returns true if the function "WEBEXT.metricsPrivate.recordMediumCount" exists.
func HasFuncRecordMediumCount() bool {
	return js.True == bindings.HasFuncRecordMediumCount()
}

// FuncRecordMediumCount returns the function "WEBEXT.metricsPrivate.recordMediumCount".
func FuncRecordMediumCount() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordMediumCount(
		js.Pointer(&fn),
	)
	return
}

// RecordMediumCount calls the function "WEBEXT.metricsPrivate.recordMediumCount" directly.
func RecordMediumCount(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordMediumCount(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordMediumCount calls the function "WEBEXT.metricsPrivate.recordMediumCount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordMediumCount(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordMediumCount(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordMediumTime returns true if the function "WEBEXT.metricsPrivate.recordMediumTime" exists.
func HasFuncRecordMediumTime() bool {
	return js.True == bindings.HasFuncRecordMediumTime()
}

// FuncRecordMediumTime returns the function "WEBEXT.metricsPrivate.recordMediumTime".
func FuncRecordMediumTime() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordMediumTime(
		js.Pointer(&fn),
	)
	return
}

// RecordMediumTime calls the function "WEBEXT.metricsPrivate.recordMediumTime" directly.
func RecordMediumTime(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordMediumTime(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordMediumTime calls the function "WEBEXT.metricsPrivate.recordMediumTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordMediumTime(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordMediumTime(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordPercentage returns true if the function "WEBEXT.metricsPrivate.recordPercentage" exists.
func HasFuncRecordPercentage() bool {
	return js.True == bindings.HasFuncRecordPercentage()
}

// FuncRecordPercentage returns the function "WEBEXT.metricsPrivate.recordPercentage".
func FuncRecordPercentage() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordPercentage(
		js.Pointer(&fn),
	)
	return
}

// RecordPercentage calls the function "WEBEXT.metricsPrivate.recordPercentage" directly.
func RecordPercentage(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordPercentage(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordPercentage calls the function "WEBEXT.metricsPrivate.recordPercentage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordPercentage(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordPercentage(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordSmallCount returns true if the function "WEBEXT.metricsPrivate.recordSmallCount" exists.
func HasFuncRecordSmallCount() bool {
	return js.True == bindings.HasFuncRecordSmallCount()
}

// FuncRecordSmallCount returns the function "WEBEXT.metricsPrivate.recordSmallCount".
func FuncRecordSmallCount() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordSmallCount(
		js.Pointer(&fn),
	)
	return
}

// RecordSmallCount calls the function "WEBEXT.metricsPrivate.recordSmallCount" directly.
func RecordSmallCount(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordSmallCount(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordSmallCount calls the function "WEBEXT.metricsPrivate.recordSmallCount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordSmallCount(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordSmallCount(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordSparseValue returns true if the function "WEBEXT.metricsPrivate.recordSparseValue" exists.
func HasFuncRecordSparseValue() bool {
	return js.True == bindings.HasFuncRecordSparseValue()
}

// FuncRecordSparseValue returns the function "WEBEXT.metricsPrivate.recordSparseValue".
func FuncRecordSparseValue() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordSparseValue(
		js.Pointer(&fn),
	)
	return
}

// RecordSparseValue calls the function "WEBEXT.metricsPrivate.recordSparseValue" directly.
func RecordSparseValue(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordSparseValue(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordSparseValue calls the function "WEBEXT.metricsPrivate.recordSparseValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordSparseValue(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordSparseValue(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordSparseValueWithHashMetricName returns true if the function "WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName" exists.
func HasFuncRecordSparseValueWithHashMetricName() bool {
	return js.True == bindings.HasFuncRecordSparseValueWithHashMetricName()
}

// FuncRecordSparseValueWithHashMetricName returns the function "WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName".
func FuncRecordSparseValueWithHashMetricName() (fn js.Func[func(metricName js.String, value js.String)]) {
	bindings.FuncRecordSparseValueWithHashMetricName(
		js.Pointer(&fn),
	)
	return
}

// RecordSparseValueWithHashMetricName calls the function "WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName" directly.
func RecordSparseValueWithHashMetricName(metricName js.String, value js.String) (ret js.Void) {
	bindings.CallRecordSparseValueWithHashMetricName(
		js.Pointer(&ret),
		metricName.Ref(),
		value.Ref(),
	)

	return
}

// TryRecordSparseValueWithHashMetricName calls the function "WEBEXT.metricsPrivate.recordSparseValueWithHashMetricName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordSparseValueWithHashMetricName(metricName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordSparseValueWithHashMetricName(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncRecordSparseValueWithPersistentHash returns true if the function "WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash" exists.
func HasFuncRecordSparseValueWithPersistentHash() bool {
	return js.True == bindings.HasFuncRecordSparseValueWithPersistentHash()
}

// FuncRecordSparseValueWithPersistentHash returns the function "WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash".
func FuncRecordSparseValueWithPersistentHash() (fn js.Func[func(metricName js.String, value js.String)]) {
	bindings.FuncRecordSparseValueWithPersistentHash(
		js.Pointer(&fn),
	)
	return
}

// RecordSparseValueWithPersistentHash calls the function "WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash" directly.
func RecordSparseValueWithPersistentHash(metricName js.String, value js.String) (ret js.Void) {
	bindings.CallRecordSparseValueWithPersistentHash(
		js.Pointer(&ret),
		metricName.Ref(),
		value.Ref(),
	)

	return
}

// TryRecordSparseValueWithPersistentHash calls the function "WEBEXT.metricsPrivate.recordSparseValueWithPersistentHash"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordSparseValueWithPersistentHash(metricName js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordSparseValueWithPersistentHash(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncRecordTime returns true if the function "WEBEXT.metricsPrivate.recordTime" exists.
func HasFuncRecordTime() bool {
	return js.True == bindings.HasFuncRecordTime()
}

// FuncRecordTime returns the function "WEBEXT.metricsPrivate.recordTime".
func FuncRecordTime() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordTime(
		js.Pointer(&fn),
	)
	return
}

// RecordTime calls the function "WEBEXT.metricsPrivate.recordTime" directly.
func RecordTime(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordTime(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordTime calls the function "WEBEXT.metricsPrivate.recordTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordTime(metricName js.String, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordTime(
		js.Pointer(&ret), js.Pointer(&exception),
		metricName.Ref(),
		float64(value),
	)

	return
}

// HasFuncRecordUserAction returns true if the function "WEBEXT.metricsPrivate.recordUserAction" exists.
func HasFuncRecordUserAction() bool {
	return js.True == bindings.HasFuncRecordUserAction()
}

// FuncRecordUserAction returns the function "WEBEXT.metricsPrivate.recordUserAction".
func FuncRecordUserAction() (fn js.Func[func(name js.String)]) {
	bindings.FuncRecordUserAction(
		js.Pointer(&fn),
	)
	return
}

// RecordUserAction calls the function "WEBEXT.metricsPrivate.recordUserAction" directly.
func RecordUserAction(name js.String) (ret js.Void) {
	bindings.CallRecordUserAction(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRecordUserAction calls the function "WEBEXT.metricsPrivate.recordUserAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordUserAction(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordUserAction(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncRecordValue returns true if the function "WEBEXT.metricsPrivate.recordValue" exists.
func HasFuncRecordValue() bool {
	return js.True == bindings.HasFuncRecordValue()
}

// FuncRecordValue returns the function "WEBEXT.metricsPrivate.recordValue".
func FuncRecordValue() (fn js.Func[func(metric MetricType, value int64)]) {
	bindings.FuncRecordValue(
		js.Pointer(&fn),
	)
	return
}

// RecordValue calls the function "WEBEXT.metricsPrivate.recordValue" directly.
func RecordValue(metric MetricType, value int64) (ret js.Void) {
	bindings.CallRecordValue(
		js.Pointer(&ret),
		js.Pointer(&metric),
		float64(value),
	)

	return
}

// TryRecordValue calls the function "WEBEXT.metricsPrivate.recordValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordValue(metric MetricType, value int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordValue(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&metric),
		float64(value),
	)

	return
}

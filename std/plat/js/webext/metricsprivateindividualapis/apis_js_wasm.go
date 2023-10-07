// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package metricsprivateindividualapis

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/metricsprivateindividualapis/bindings"
)

// HasFuncRecordEnumerationValue returns true if the function "WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue" exists.
func HasFuncRecordEnumerationValue() bool {
	return js.True == bindings.HasFuncRecordEnumerationValue()
}

// FuncRecordEnumerationValue returns the function "WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue".
func FuncRecordEnumerationValue() (fn js.Func[func(metricName js.String, value int64, enumSize int64)]) {
	bindings.FuncRecordEnumerationValue(
		js.Pointer(&fn),
	)
	return
}

// RecordEnumerationValue calls the function "WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue" directly.
func RecordEnumerationValue(metricName js.String, value int64, enumSize int64) (ret js.Void) {
	bindings.CallRecordEnumerationValue(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
		float64(enumSize),
	)

	return
}

// TryRecordEnumerationValue calls the function "WEBEXT.metricsPrivateIndividualApis.recordEnumerationValue"
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

// HasFuncRecordMediumTime returns true if the function "WEBEXT.metricsPrivateIndividualApis.recordMediumTime" exists.
func HasFuncRecordMediumTime() bool {
	return js.True == bindings.HasFuncRecordMediumTime()
}

// FuncRecordMediumTime returns the function "WEBEXT.metricsPrivateIndividualApis.recordMediumTime".
func FuncRecordMediumTime() (fn js.Func[func(metricName js.String, value int64)]) {
	bindings.FuncRecordMediumTime(
		js.Pointer(&fn),
	)
	return
}

// RecordMediumTime calls the function "WEBEXT.metricsPrivateIndividualApis.recordMediumTime" directly.
func RecordMediumTime(metricName js.String, value int64) (ret js.Void) {
	bindings.CallRecordMediumTime(
		js.Pointer(&ret),
		metricName.Ref(),
		float64(value),
	)

	return
}

// TryRecordMediumTime calls the function "WEBEXT.metricsPrivateIndividualApis.recordMediumTime"
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

// HasFuncRecordUserAction returns true if the function "WEBEXT.metricsPrivateIndividualApis.recordUserAction" exists.
func HasFuncRecordUserAction() bool {
	return js.True == bindings.HasFuncRecordUserAction()
}

// FuncRecordUserAction returns the function "WEBEXT.metricsPrivateIndividualApis.recordUserAction".
func FuncRecordUserAction() (fn js.Func[func(name js.String)]) {
	bindings.FuncRecordUserAction(
		js.Pointer(&fn),
	)
	return
}

// RecordUserAction calls the function "WEBEXT.metricsPrivateIndividualApis.recordUserAction" directly.
func RecordUserAction(name js.String) (ret js.Void) {
	bindings.CallRecordUserAction(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRecordUserAction calls the function "WEBEXT.metricsPrivateIndividualApis.recordUserAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordUserAction(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordUserAction(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

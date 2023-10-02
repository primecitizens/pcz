// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type PerformanceMeasureOptions struct {
	// Detail is "PerformanceMeasureOptions.detail"
	//
	// Optional
	Detail js.Any
	// Start is "PerformanceMeasureOptions.start"
	//
	// Optional
	Start OneOf_String_Float64
	// Duration is "PerformanceMeasureOptions.duration"
	//
	// Optional
	Duration DOMHighResTimeStamp
	// End is "PerformanceMeasureOptions.end"
	//
	// Optional
	End OneOf_String_Float64

	FFI_USE_Duration bool // for Duration.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceMeasureOptions with all fields set.
func (p PerformanceMeasureOptions) FromRef(ref js.Ref) PerformanceMeasureOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PerformanceMeasureOptions PerformanceMeasureOptions [// PerformanceMeasureOptions] [0x140007ef4a0 0x140007ef540 0x140007ef5e0 0x140007ef720 0x140007ef680] 0x14000780e88 {0 0}} in the application heap.
func (p PerformanceMeasureOptions) New() js.Ref {
	return bindings.PerformanceMeasureOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PerformanceMeasureOptions) UpdateFrom(ref js.Ref) {
	bindings.PerformanceMeasureOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PerformanceMeasureOptions) Update(ref js.Ref) {
	bindings.PerformanceMeasureOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_String_PerformanceMeasureOptions struct {
	ref js.Ref
}

func (x OneOf_String_PerformanceMeasureOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_PerformanceMeasureOptions) Free() {
	x.ref.Free()
}

func (x OneOf_String_PerformanceMeasureOptions) FromRef(ref js.Ref) OneOf_String_PerformanceMeasureOptions {
	return OneOf_String_PerformanceMeasureOptions{
		ref: ref,
	}
}

func (x OneOf_String_PerformanceMeasureOptions) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_PerformanceMeasureOptions) PerformanceMeasureOptions() PerformanceMeasureOptions {
	var ret PerformanceMeasureOptions
	ret.UpdateFrom(x.ref)
	return ret
}

type PerformanceEntry struct {
	ref js.Ref
}

func (this PerformanceEntry) Once() PerformanceEntry {
	this.Ref().Once()
	return this
}

func (this PerformanceEntry) Ref() js.Ref {
	return this.ref
}

func (this PerformanceEntry) FromRef(ref js.Ref) PerformanceEntry {
	this.ref = ref
	return this
}

func (this PerformanceEntry) Free() {
	this.Ref().Free()
}

// Name returns the value of property "PerformanceEntry.name".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEntry) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEntryName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EntryType returns the value of property "PerformanceEntry.entryType".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEntry) EntryType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEntryEntryType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// StartTime returns the value of property "PerformanceEntry.startTime".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEntry) StartTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEntryStartTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Duration returns the value of property "PerformanceEntry.duration".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEntry) Duration() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEntryDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ToJSON calls the method "PerformanceEntry.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceEntry) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceEntryToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceEntry.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceEntry) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceEntryToJSONFunc(
			this.Ref(),
		),
	)
}

type PerformanceEntryList = js.Array[PerformanceEntry]

type PerformanceTiming struct {
	ref js.Ref
}

func (this PerformanceTiming) Once() PerformanceTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceTiming) Ref() js.Ref {
	return this.ref
}

func (this PerformanceTiming) FromRef(ref js.Ref) PerformanceTiming {
	this.ref = ref
	return this
}

func (this PerformanceTiming) Free() {
	this.Ref().Free()
}

// NavigationStart returns the value of property "PerformanceTiming.navigationStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) NavigationStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingNavigationStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// UnloadEventStart returns the value of property "PerformanceTiming.unloadEventStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) UnloadEventStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingUnloadEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// UnloadEventEnd returns the value of property "PerformanceTiming.unloadEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) UnloadEventEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingUnloadEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// RedirectStart returns the value of property "PerformanceTiming.redirectStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) RedirectStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingRedirectStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// RedirectEnd returns the value of property "PerformanceTiming.redirectEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) RedirectEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingRedirectEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// FetchStart returns the value of property "PerformanceTiming.fetchStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) FetchStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingFetchStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomainLookupStart returns the value of property "PerformanceTiming.domainLookupStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomainLookupStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomainLookupStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomainLookupEnd returns the value of property "PerformanceTiming.domainLookupEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomainLookupEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomainLookupEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ConnectStart returns the value of property "PerformanceTiming.connectStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) ConnectStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingConnectStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ConnectEnd returns the value of property "PerformanceTiming.connectEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) ConnectEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingConnectEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// SecureConnectionStart returns the value of property "PerformanceTiming.secureConnectionStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) SecureConnectionStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingSecureConnectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// RequestStart returns the value of property "PerformanceTiming.requestStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) RequestStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingRequestStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ResponseStart returns the value of property "PerformanceTiming.responseStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) ResponseStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingResponseStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ResponseEnd returns the value of property "PerformanceTiming.responseEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) ResponseEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingResponseEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomLoading returns the value of property "PerformanceTiming.domLoading".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomLoading() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomLoading(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomInteractive returns the value of property "PerformanceTiming.domInteractive".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomInteractive() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomInteractive(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomContentLoadedEventStart returns the value of property "PerformanceTiming.domContentLoadedEventStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomContentLoadedEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomContentLoadedEventEnd returns the value of property "PerformanceTiming.domContentLoadedEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomContentLoadedEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DomComplete returns the value of property "PerformanceTiming.domComplete".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) DomComplete() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingDomComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// LoadEventStart returns the value of property "PerformanceTiming.loadEventStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) LoadEventStart() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingLoadEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// LoadEventEnd returns the value of property "PerformanceTiming.loadEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceTiming) LoadEventEnd() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimingLoadEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ToJSON calls the method "PerformanceTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceTimingToJSONFunc(
			this.Ref(),
		),
	)
}

const (
	PerformanceNavigation_TYPE_NAVIGATE     uint16 = 0
	PerformanceNavigation_TYPE_RELOAD       uint16 = 1
	PerformanceNavigation_TYPE_BACK_FORWARD uint16 = 2
	PerformanceNavigation_TYPE_RESERVED     uint16 = 255
)

type PerformanceNavigation struct {
	ref js.Ref
}

func (this PerformanceNavigation) Once() PerformanceNavigation {
	this.Ref().Once()
	return this
}

func (this PerformanceNavigation) Ref() js.Ref {
	return this.ref
}

func (this PerformanceNavigation) FromRef(ref js.Ref) PerformanceNavigation {
	this.ref = ref
	return this
}

func (this PerformanceNavigation) Free() {
	this.Ref().Free()
}

// Type returns the value of property "PerformanceNavigation.type".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigation) Type() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// RedirectCount returns the value of property "PerformanceNavigation.redirectCount".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigation) RedirectCount() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationRedirectCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// ToJSON calls the method "PerformanceNavigation.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceNavigation) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceNavigationToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceNavigation.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceNavigation) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceNavigationToJSONFunc(
			this.Ref(),
		),
	)
}

type EventCounts struct {
	ref js.Ref
}

func (this EventCounts) Once() EventCounts {
	this.Ref().Once()
	return this
}

func (this EventCounts) Ref() js.Ref {
	return this.ref
}

func (this EventCounts) FromRef(ref js.Ref) EventCounts {
	this.ref = ref
	return this
}

func (this EventCounts) Free() {
	this.Ref().Free()
}

type Performance struct {
	EventTarget
}

func (this Performance) Once() Performance {
	this.Ref().Once()
	return this
}

func (this Performance) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Performance) FromRef(ref js.Ref) Performance {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Performance) Free() {
	this.Ref().Free()
}

// TimeOrigin returns the value of property "Performance.timeOrigin".
//
// The returned bool will be false if there is no such property.
func (this Performance) TimeOrigin() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTimeOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Timing returns the value of property "Performance.timing".
//
// The returned bool will be false if there is no such property.
func (this Performance) Timing() (PerformanceTiming, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceTiming(
		this.Ref(), js.Pointer(&_ok),
	)
	return PerformanceTiming{}.FromRef(_ret), _ok
}

// Navigation returns the value of property "Performance.navigation".
//
// The returned bool will be false if there is no such property.
func (this Performance) Navigation() (PerformanceNavigation, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigation(
		this.Ref(), js.Pointer(&_ok),
	)
	return PerformanceNavigation{}.FromRef(_ret), _ok
}

// EventCounts returns the value of property "Performance.eventCounts".
//
// The returned bool will be false if there is no such property.
func (this Performance) EventCounts() (EventCounts, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventCounts(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventCounts{}.FromRef(_ret), _ok
}

// InteractionCount returns the value of property "Performance.interactionCount".
//
// The returned bool will be false if there is no such property.
func (this Performance) InteractionCount() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceInteractionCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Now calls the method "Performance.now".
//
// The returned bool will be false if there is no such method.
func (this Performance) Now() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceNow(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMHighResTimeStamp(_ret), _ok
}

// NowFunc returns the method "Performance.now".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) NowFunc() (fn js.Func[func() DOMHighResTimeStamp]) {
	return fn.FromRef(
		bindings.PerformanceNowFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "Performance.toJSON".
//
// The returned bool will be false if there is no such method.
func (this Performance) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "Performance.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceToJSONFunc(
			this.Ref(),
		),
	)
}

// MeasureUserAgentSpecificMemory calls the method "Performance.measureUserAgentSpecificMemory".
//
// The returned bool will be false if there is no such method.
func (this Performance) MeasureUserAgentSpecificMemory() (js.Promise[MemoryMeasurement], bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMeasureUserAgentSpecificMemory(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MemoryMeasurement]{}.FromRef(_ret), _ok
}

// MeasureUserAgentSpecificMemoryFunc returns the method "Performance.measureUserAgentSpecificMemory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) MeasureUserAgentSpecificMemoryFunc() (fn js.Func[func() js.Promise[MemoryMeasurement]]) {
	return fn.FromRef(
		bindings.PerformanceMeasureUserAgentSpecificMemoryFunc(
			this.Ref(),
		),
	)
}

// Mark calls the method "Performance.mark".
//
// The returned bool will be false if there is no such method.
func (this Performance) Mark(markName js.String, markOptions PerformanceMarkOptions) (PerformanceMark, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMark(
		this.Ref(), js.Pointer(&_ok),
		markName.Ref(),
		js.Pointer(&markOptions),
	)

	return PerformanceMark{}.FromRef(_ret), _ok
}

// MarkFunc returns the method "Performance.mark".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) MarkFunc() (fn js.Func[func(markName js.String, markOptions PerformanceMarkOptions) PerformanceMark]) {
	return fn.FromRef(
		bindings.PerformanceMarkFunc(
			this.Ref(),
		),
	)
}

// Mark1 calls the method "Performance.mark".
//
// The returned bool will be false if there is no such method.
func (this Performance) Mark1(markName js.String) (PerformanceMark, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMark1(
		this.Ref(), js.Pointer(&_ok),
		markName.Ref(),
	)

	return PerformanceMark{}.FromRef(_ret), _ok
}

// Mark1Func returns the method "Performance.mark".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) Mark1Func() (fn js.Func[func(markName js.String) PerformanceMark]) {
	return fn.FromRef(
		bindings.PerformanceMark1Func(
			this.Ref(),
		),
	)
}

// ClearMarks calls the method "Performance.clearMarks".
//
// The returned bool will be false if there is no such method.
func (this Performance) ClearMarks(markName js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceClearMarks(
		this.Ref(), js.Pointer(&_ok),
		markName.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearMarksFunc returns the method "Performance.clearMarks".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ClearMarksFunc() (fn js.Func[func(markName js.String)]) {
	return fn.FromRef(
		bindings.PerformanceClearMarksFunc(
			this.Ref(),
		),
	)
}

// ClearMarks1 calls the method "Performance.clearMarks".
//
// The returned bool will be false if there is no such method.
func (this Performance) ClearMarks1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceClearMarks1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearMarks1Func returns the method "Performance.clearMarks".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ClearMarks1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearMarks1Func(
			this.Ref(),
		),
	)
}

// Measure calls the method "Performance.measure".
//
// The returned bool will be false if there is no such method.
func (this Performance) Measure(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) (PerformanceMeasure, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMeasure(
		this.Ref(), js.Pointer(&_ok),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
		endMark.Ref(),
	)

	return PerformanceMeasure{}.FromRef(_ret), _ok
}

// MeasureFunc returns the method "Performance.measure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) MeasureFunc() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasureFunc(
			this.Ref(),
		),
	)
}

// Measure1 calls the method "Performance.measure".
//
// The returned bool will be false if there is no such method.
func (this Performance) Measure1(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) (PerformanceMeasure, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMeasure1(
		this.Ref(), js.Pointer(&_ok),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
	)

	return PerformanceMeasure{}.FromRef(_ret), _ok
}

// Measure1Func returns the method "Performance.measure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) Measure1Func() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasure1Func(
			this.Ref(),
		),
	)
}

// Measure2 calls the method "Performance.measure".
//
// The returned bool will be false if there is no such method.
func (this Performance) Measure2(measureName js.String) (PerformanceMeasure, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceMeasure2(
		this.Ref(), js.Pointer(&_ok),
		measureName.Ref(),
	)

	return PerformanceMeasure{}.FromRef(_ret), _ok
}

// Measure2Func returns the method "Performance.measure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) Measure2Func() (fn js.Func[func(measureName js.String) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasure2Func(
			this.Ref(),
		),
	)
}

// ClearMeasures calls the method "Performance.clearMeasures".
//
// The returned bool will be false if there is no such method.
func (this Performance) ClearMeasures(measureName js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceClearMeasures(
		this.Ref(), js.Pointer(&_ok),
		measureName.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearMeasuresFunc returns the method "Performance.clearMeasures".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ClearMeasuresFunc() (fn js.Func[func(measureName js.String)]) {
	return fn.FromRef(
		bindings.PerformanceClearMeasuresFunc(
			this.Ref(),
		),
	)
}

// ClearMeasures1 calls the method "Performance.clearMeasures".
//
// The returned bool will be false if there is no such method.
func (this Performance) ClearMeasures1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceClearMeasures1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearMeasures1Func returns the method "Performance.clearMeasures".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ClearMeasures1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearMeasures1Func(
			this.Ref(),
		),
	)
}

// ClearResourceTimings calls the method "Performance.clearResourceTimings".
//
// The returned bool will be false if there is no such method.
func (this Performance) ClearResourceTimings() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceClearResourceTimings(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearResourceTimingsFunc returns the method "Performance.clearResourceTimings".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) ClearResourceTimingsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearResourceTimingsFunc(
			this.Ref(),
		),
	)
}

// SetResourceTimingBufferSize calls the method "Performance.setResourceTimingBufferSize".
//
// The returned bool will be false if there is no such method.
func (this Performance) SetResourceTimingBufferSize(maxSize uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceSetResourceTimingBufferSize(
		this.Ref(), js.Pointer(&_ok),
		uint32(maxSize),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetResourceTimingBufferSizeFunc returns the method "Performance.setResourceTimingBufferSize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) SetResourceTimingBufferSizeFunc() (fn js.Func[func(maxSize uint32)]) {
	return fn.FromRef(
		bindings.PerformanceSetResourceTimingBufferSizeFunc(
			this.Ref(),
		),
	)
}

// GetEntries calls the method "Performance.getEntries".
//
// The returned bool will be false if there is no such method.
func (this Performance) GetEntries() (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceGetEntries(
		this.Ref(), js.Pointer(&_ok),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesFunc returns the method "Performance.getEntries".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) GetEntriesFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByType calls the method "Performance.getEntriesByType".
//
// The returned bool will be false if there is no such method.
func (this Performance) GetEntriesByType(typ js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceGetEntriesByType(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByTypeFunc returns the method "Performance.getEntriesByType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) GetEntriesByTypeFunc() (fn js.Func[func(typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByTypeFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName calls the method "Performance.getEntriesByName".
//
// The returned bool will be false if there is no such method.
func (this Performance) GetEntriesByName(name js.String, typ js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceGetEntriesByName(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		typ.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByNameFunc returns the method "Performance.getEntriesByName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) GetEntriesByNameFunc() (fn js.Func[func(name js.String, typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByNameFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName1 calls the method "Performance.getEntriesByName".
//
// The returned bool will be false if there is no such method.
func (this Performance) GetEntriesByName1(name js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceGetEntriesByName1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByName1Func returns the method "Performance.getEntriesByName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Performance) GetEntriesByName1Func() (fn js.Func[func(name js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByName1Func(
			this.Ref(),
		),
	)
}

type Storage struct {
	ref js.Ref
}

func (this Storage) Once() Storage {
	this.Ref().Once()
	return this
}

func (this Storage) Ref() js.Ref {
	return this.ref
}

func (this Storage) FromRef(ref js.Ref) Storage {
	this.ref = ref
	return this
}

func (this Storage) Free() {
	this.Ref().Free()
}

// Length returns the value of property "Storage.length".
//
// The returned bool will be false if there is no such property.
func (this Storage) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetStorageLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Key calls the method "Storage.key".
//
// The returned bool will be false if there is no such method.
func (this Storage) Key(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallStorageKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// KeyFunc returns the method "Storage.key".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Storage) KeyFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.StorageKeyFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "Storage.getItem".
//
// The returned bool will be false if there is no such method.
func (this Storage) GetItem(key js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallStorageGetItem(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "Storage.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Storage) GetItemFunc() (fn js.Func[func(key js.String) js.String]) {
	return fn.FromRef(
		bindings.StorageGetItemFunc(
			this.Ref(),
		),
	)
}

// SetItem calls the method "Storage.setItem".
//
// The returned bool will be false if there is no such method.
func (this Storage) SetItem(key js.String, value js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageSetItem(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetItemFunc returns the method "Storage.setItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Storage) SetItemFunc() (fn js.Func[func(key js.String, value js.String)]) {
	return fn.FromRef(
		bindings.StorageSetItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "Storage.removeItem".
//
// The returned bool will be false if there is no such method.
func (this Storage) RemoveItem(key js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveItemFunc returns the method "Storage.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Storage) RemoveItemFunc() (fn js.Func[func(key js.String)]) {
	return fn.FromRef(
		bindings.StorageRemoveItemFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "Storage.clear".
//
// The returned bool will be false if there is no such method.
func (this Storage) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "Storage.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Storage) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.StorageClearFunc(
			this.Ref(),
		),
	)
}

type Window struct {
	EventTarget
}

func (this Window) Once() Window {
	this.Ref().Once()
	return this
}

func (this Window) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Window) FromRef(ref js.Ref) Window {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Window) Free() {
	this.Ref().Free()
}

// Window returns the value of property "Window.window".
//
// The returned bool will be false if there is no such property.
func (this Window) Window() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetWindowWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Self returns the value of property "Window.self".
//
// The returned bool will be false if there is no such property.
func (this Window) Self() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetWindowSelf(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Document returns the value of property "Window.document".
//
// The returned bool will be false if there is no such property.
func (this Window) Document() (Document, bool) {
	var _ok bool
	_ret := bindings.GetWindowDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// Name returns the value of property "Window.name".
//
// The returned bool will be false if there is no such property.
func (this Window) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWindowName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "Window.name" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetName(val js.String) bool {
	return js.True == bindings.SetWindowName(
		this.Ref(),
		val.Ref(),
	)
}

// Location returns the value of property "Window.location".
//
// The returned bool will be false if there is no such property.
func (this Window) Location() (Location, bool) {
	var _ok bool
	_ret := bindings.GetWindowLocation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Location{}.FromRef(_ret), _ok
}

// History returns the value of property "Window.history".
//
// The returned bool will be false if there is no such property.
func (this Window) History() (History, bool) {
	var _ok bool
	_ret := bindings.GetWindowHistory(
		this.Ref(), js.Pointer(&_ok),
	)
	return History{}.FromRef(_ret), _ok
}

// Navigation returns the value of property "Window.navigation".
//
// The returned bool will be false if there is no such property.
func (this Window) Navigation() (Navigation, bool) {
	var _ok bool
	_ret := bindings.GetWindowNavigation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Navigation{}.FromRef(_ret), _ok
}

// CustomElements returns the value of property "Window.customElements".
//
// The returned bool will be false if there is no such property.
func (this Window) CustomElements() (CustomElementRegistry, bool) {
	var _ok bool
	_ret := bindings.GetWindowCustomElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return CustomElementRegistry{}.FromRef(_ret), _ok
}

// Locationbar returns the value of property "Window.locationbar".
//
// The returned bool will be false if there is no such property.
func (this Window) Locationbar() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowLocationbar(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Menubar returns the value of property "Window.menubar".
//
// The returned bool will be false if there is no such property.
func (this Window) Menubar() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowMenubar(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Personalbar returns the value of property "Window.personalbar".
//
// The returned bool will be false if there is no such property.
func (this Window) Personalbar() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowPersonalbar(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Scrollbars returns the value of property "Window.scrollbars".
//
// The returned bool will be false if there is no such property.
func (this Window) Scrollbars() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowScrollbars(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Statusbar returns the value of property "Window.statusbar".
//
// The returned bool will be false if there is no such property.
func (this Window) Statusbar() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowStatusbar(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Toolbar returns the value of property "Window.toolbar".
//
// The returned bool will be false if there is no such property.
func (this Window) Toolbar() (BarProp, bool) {
	var _ok bool
	_ret := bindings.GetWindowToolbar(
		this.Ref(), js.Pointer(&_ok),
	)
	return BarProp{}.FromRef(_ret), _ok
}

// Status returns the value of property "Window.status".
//
// The returned bool will be false if there is no such property.
func (this Window) Status() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWindowStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Status sets the value of property "Window.status" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetStatus(val js.String) bool {
	return js.True == bindings.SetWindowStatus(
		this.Ref(),
		val.Ref(),
	)
}

// Closed returns the value of property "Window.closed".
//
// The returned bool will be false if there is no such property.
func (this Window) Closed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Frames returns the value of property "Window.frames".
//
// The returned bool will be false if there is no such property.
func (this Window) Frames() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetWindowFrames(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Length returns the value of property "Window.length".
//
// The returned bool will be false if there is no such property.
func (this Window) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetWindowLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Top returns the value of property "Window.top".
//
// The returned bool will be false if there is no such property.
func (this Window) Top() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetWindowTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Opener returns the value of property "Window.opener".
//
// The returned bool will be false if there is no such property.
func (this Window) Opener() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetWindowOpener(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Opener sets the value of property "Window.opener" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetOpener(val js.Any) bool {
	return js.True == bindings.SetWindowOpener(
		this.Ref(),
		val.Ref(),
	)
}

// Parent returns the value of property "Window.parent".
//
// The returned bool will be false if there is no such property.
func (this Window) Parent() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetWindowParent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// FrameElement returns the value of property "Window.frameElement".
//
// The returned bool will be false if there is no such property.
func (this Window) FrameElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetWindowFrameElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Navigator returns the value of property "Window.navigator".
//
// The returned bool will be false if there is no such property.
func (this Window) Navigator() (Navigator, bool) {
	var _ok bool
	_ret := bindings.GetWindowNavigator(
		this.Ref(), js.Pointer(&_ok),
	)
	return Navigator{}.FromRef(_ret), _ok
}

// ClientInformation returns the value of property "Window.clientInformation".
//
// The returned bool will be false if there is no such property.
func (this Window) ClientInformation() (Navigator, bool) {
	var _ok bool
	_ret := bindings.GetWindowClientInformation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Navigator{}.FromRef(_ret), _ok
}

// OriginAgentCluster returns the value of property "Window.originAgentCluster".
//
// The returned bool will be false if there is no such property.
func (this Window) OriginAgentCluster() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowOriginAgentCluster(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Orientation returns the value of property "Window.orientation".
//
// The returned bool will be false if there is no such property.
func (this Window) Orientation() (int16, bool) {
	var _ok bool
	_ret := bindings.GetWindowOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return int16(_ret), _ok
}

// LaunchQueue returns the value of property "Window.launchQueue".
//
// The returned bool will be false if there is no such property.
func (this Window) LaunchQueue() (LaunchQueue, bool) {
	var _ok bool
	_ret := bindings.GetWindowLaunchQueue(
		this.Ref(), js.Pointer(&_ok),
	)
	return LaunchQueue{}.FromRef(_ret), _ok
}

// Fence returns the value of property "Window.fence".
//
// The returned bool will be false if there is no such property.
func (this Window) Fence() (Fence, bool) {
	var _ok bool
	_ret := bindings.GetWindowFence(
		this.Ref(), js.Pointer(&_ok),
	)
	return Fence{}.FromRef(_ret), _ok
}

// PortalHost returns the value of property "Window.portalHost".
//
// The returned bool will be false if there is no such property.
func (this Window) PortalHost() (PortalHost, bool) {
	var _ok bool
	_ret := bindings.GetWindowPortalHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return PortalHost{}.FromRef(_ret), _ok
}

// Screen returns the value of property "Window.screen".
//
// The returned bool will be false if there is no such property.
func (this Window) Screen() (Screen, bool) {
	var _ok bool
	_ret := bindings.GetWindowScreen(
		this.Ref(), js.Pointer(&_ok),
	)
	return Screen{}.FromRef(_ret), _ok
}

// VisualViewport returns the value of property "Window.visualViewport".
//
// The returned bool will be false if there is no such property.
func (this Window) VisualViewport() (VisualViewport, bool) {
	var _ok bool
	_ret := bindings.GetWindowVisualViewport(
		this.Ref(), js.Pointer(&_ok),
	)
	return VisualViewport{}.FromRef(_ret), _ok
}

// InnerWidth returns the value of property "Window.innerWidth".
//
// The returned bool will be false if there is no such property.
func (this Window) InnerWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowInnerWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// InnerHeight returns the value of property "Window.innerHeight".
//
// The returned bool will be false if there is no such property.
func (this Window) InnerHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowInnerHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScrollX returns the value of property "Window.scrollX".
//
// The returned bool will be false if there is no such property.
func (this Window) ScrollX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWindowScrollX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageXOffset returns the value of property "Window.pageXOffset".
//
// The returned bool will be false if there is no such property.
func (this Window) PageXOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWindowPageXOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ScrollY returns the value of property "Window.scrollY".
//
// The returned bool will be false if there is no such property.
func (this Window) ScrollY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWindowScrollY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageYOffset returns the value of property "Window.pageYOffset".
//
// The returned bool will be false if there is no such property.
func (this Window) PageYOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWindowPageYOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ScreenX returns the value of property "Window.screenX".
//
// The returned bool will be false if there is no such property.
func (this Window) ScreenX() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowScreenX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScreenLeft returns the value of property "Window.screenLeft".
//
// The returned bool will be false if there is no such property.
func (this Window) ScreenLeft() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowScreenLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScreenY returns the value of property "Window.screenY".
//
// The returned bool will be false if there is no such property.
func (this Window) ScreenY() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowScreenY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScreenTop returns the value of property "Window.screenTop".
//
// The returned bool will be false if there is no such property.
func (this Window) ScreenTop() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowScreenTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// OuterWidth returns the value of property "Window.outerWidth".
//
// The returned bool will be false if there is no such property.
func (this Window) OuterWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowOuterWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// OuterHeight returns the value of property "Window.outerHeight".
//
// The returned bool will be false if there is no such property.
func (this Window) OuterHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetWindowOuterHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// DevicePixelRatio returns the value of property "Window.devicePixelRatio".
//
// The returned bool will be false if there is no such property.
func (this Window) DevicePixelRatio() (float64, bool) {
	var _ok bool
	_ret := bindings.GetWindowDevicePixelRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// SharedStorage returns the value of property "Window.sharedStorage".
//
// The returned bool will be false if there is no such property.
func (this Window) SharedStorage() (WindowSharedStorage, bool) {
	var _ok bool
	_ret := bindings.GetWindowSharedStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return WindowSharedStorage{}.FromRef(_ret), _ok
}

// Event returns the value of property "Window.event".
//
// The returned bool will be false if there is no such property.
func (this Window) Event() (OneOf_Event_undefined, bool) {
	var _ok bool
	_ret := bindings.GetWindowEvent(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Event_undefined{}.FromRef(_ret), _ok
}

// CookieStore returns the value of property "Window.cookieStore".
//
// The returned bool will be false if there is no such property.
func (this Window) CookieStore() (CookieStore, bool) {
	var _ok bool
	_ret := bindings.GetWindowCookieStore(
		this.Ref(), js.Pointer(&_ok),
	)
	return CookieStore{}.FromRef(_ret), _ok
}

// DocumentPictureInPicture returns the value of property "Window.documentPictureInPicture".
//
// The returned bool will be false if there is no such property.
func (this Window) DocumentPictureInPicture() (DocumentPictureInPicture, bool) {
	var _ok bool
	_ret := bindings.GetWindowDocumentPictureInPicture(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentPictureInPicture{}.FromRef(_ret), _ok
}

// External returns the value of property "Window.external".
//
// The returned bool will be false if there is no such property.
func (this Window) External() (External, bool) {
	var _ok bool
	_ret := bindings.GetWindowExternal(
		this.Ref(), js.Pointer(&_ok),
	)
	return External{}.FromRef(_ret), _ok
}

// SpeechSynthesis returns the value of property "Window.speechSynthesis".
//
// The returned bool will be false if there is no such property.
func (this Window) SpeechSynthesis() (SpeechSynthesis, bool) {
	var _ok bool
	_ret := bindings.GetWindowSpeechSynthesis(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechSynthesis{}.FromRef(_ret), _ok
}

// Origin returns the value of property "Window.origin".
//
// The returned bool will be false if there is no such property.
func (this Window) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWindowOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsSecureContext returns the value of property "Window.isSecureContext".
//
// The returned bool will be false if there is no such property.
func (this Window) IsSecureContext() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowIsSecureContext(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CrossOriginIsolated returns the value of property "Window.crossOriginIsolated".
//
// The returned bool will be false if there is no such property.
func (this Window) CrossOriginIsolated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetWindowCrossOriginIsolated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Scheduler returns the value of property "Window.scheduler".
//
// The returned bool will be false if there is no such property.
func (this Window) Scheduler() (Scheduler, bool) {
	var _ok bool
	_ret := bindings.GetWindowScheduler(
		this.Ref(), js.Pointer(&_ok),
	)
	return Scheduler{}.FromRef(_ret), _ok
}

// TrustedTypes returns the value of property "Window.trustedTypes".
//
// The returned bool will be false if there is no such property.
func (this Window) TrustedTypes() (TrustedTypePolicyFactory, bool) {
	var _ok bool
	_ret := bindings.GetWindowTrustedTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return TrustedTypePolicyFactory{}.FromRef(_ret), _ok
}

// Caches returns the value of property "Window.caches".
//
// The returned bool will be false if there is no such property.
func (this Window) Caches() (CacheStorage, bool) {
	var _ok bool
	_ret := bindings.GetWindowCaches(
		this.Ref(), js.Pointer(&_ok),
	)
	return CacheStorage{}.FromRef(_ret), _ok
}

// Crypto returns the value of property "Window.crypto".
//
// The returned bool will be false if there is no such property.
func (this Window) Crypto() (Crypto, bool) {
	var _ok bool
	_ret := bindings.GetWindowCrypto(
		this.Ref(), js.Pointer(&_ok),
	)
	return Crypto{}.FromRef(_ret), _ok
}

// IndexedDB returns the value of property "Window.indexedDB".
//
// The returned bool will be false if there is no such property.
func (this Window) IndexedDB() (IDBFactory, bool) {
	var _ok bool
	_ret := bindings.GetWindowIndexedDB(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBFactory{}.FromRef(_ret), _ok
}

// Performance returns the value of property "Window.performance".
//
// The returned bool will be false if there is no such property.
func (this Window) Performance() (Performance, bool) {
	var _ok bool
	_ret := bindings.GetWindowPerformance(
		this.Ref(), js.Pointer(&_ok),
	)
	return Performance{}.FromRef(_ret), _ok
}

// SessionStorage returns the value of property "Window.sessionStorage".
//
// The returned bool will be false if there is no such property.
func (this Window) SessionStorage() (Storage, bool) {
	var _ok bool
	_ret := bindings.GetWindowSessionStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return Storage{}.FromRef(_ret), _ok
}

// LocalStorage returns the value of property "Window.localStorage".
//
// The returned bool will be false if there is no such property.
func (this Window) LocalStorage() (Storage, bool) {
	var _ok bool
	_ret := bindings.GetWindowLocalStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return Storage{}.FromRef(_ret), _ok
}

// Close calls the method "Window.close".
//
// The returned bool will be false if there is no such method.
func (this Window) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "Window.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowCloseFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "Window.stop".
//
// The returned bool will be false if there is no such method.
func (this Window) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "Window.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowStopFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "Window.focus".
//
// The returned bool will be false if there is no such method.
func (this Window) Focus() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowFocus(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FocusFunc returns the method "Window.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) FocusFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowFocusFunc(
			this.Ref(),
		),
	)
}

// Blur calls the method "Window.blur".
//
// The returned bool will be false if there is no such method.
func (this Window) Blur() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowBlur(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlurFunc returns the method "Window.blur".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowBlurFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "Window.open".
//
// The returned bool will be false if there is no such method.
func (this Window) Open(url js.String, target js.String, features js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWindowOpen(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		target.Ref(),
		features.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "Window.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) OpenFunc() (fn js.Func[func(url js.String, target js.String, features js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "Window.open".
//
// The returned bool will be false if there is no such method.
func (this Window) Open1(url js.String, target js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWindowOpen1(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		target.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// Open1Func returns the method "Window.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Open1Func() (fn js.Func[func(url js.String, target js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen1Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "Window.open".
//
// The returned bool will be false if there is no such method.
func (this Window) Open2(url js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWindowOpen2(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// Open2Func returns the method "Window.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Open2Func() (fn js.Func[func(url js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen2Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "Window.open".
//
// The returned bool will be false if there is no such method.
func (this Window) Open3() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWindowOpen3(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// Open3Func returns the method "Window.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Open3Func() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen3Func(
			this.Ref(),
		),
	)
}

// Get calls the method "Window.".
//
// The returned bool will be false if there is no such method.
func (this Window) Get(name js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWindowGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// GetFunc returns the method "Window.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowGetFunc(
			this.Ref(),
		),
	)
}

// Alert calls the method "Window.alert".
//
// The returned bool will be false if there is no such method.
func (this Window) Alert() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowAlert(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AlertFunc returns the method "Window.alert".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) AlertFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowAlertFunc(
			this.Ref(),
		),
	)
}

// Alert1 calls the method "Window.alert".
//
// The returned bool will be false if there is no such method.
func (this Window) Alert1(message js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowAlert1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Alert1Func returns the method "Window.alert".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Alert1Func() (fn js.Func[func(message js.String)]) {
	return fn.FromRef(
		bindings.WindowAlert1Func(
			this.Ref(),
		),
	)
}

// Confirm calls the method "Window.confirm".
//
// The returned bool will be false if there is no such method.
func (this Window) Confirm(message js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallWindowConfirm(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	return _ret == js.True, _ok
}

// ConfirmFunc returns the method "Window.confirm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ConfirmFunc() (fn js.Func[func(message js.String) bool]) {
	return fn.FromRef(
		bindings.WindowConfirmFunc(
			this.Ref(),
		),
	)
}

// Confirm1 calls the method "Window.confirm".
//
// The returned bool will be false if there is no such method.
func (this Window) Confirm1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallWindowConfirm1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// Confirm1Func returns the method "Window.confirm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Confirm1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WindowConfirm1Func(
			this.Ref(),
		),
	)
}

// Prompt calls the method "Window.prompt".
//
// The returned bool will be false if there is no such method.
func (this Window) Prompt(message js.String, def js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWindowPrompt(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		def.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// PromptFunc returns the method "Window.prompt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PromptFunc() (fn js.Func[func(message js.String, def js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowPromptFunc(
			this.Ref(),
		),
	)
}

// Prompt1 calls the method "Window.prompt".
//
// The returned bool will be false if there is no such method.
func (this Window) Prompt1(message js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWindowPrompt1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// Prompt1Func returns the method "Window.prompt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Prompt1Func() (fn js.Func[func(message js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowPrompt1Func(
			this.Ref(),
		),
	)
}

// Prompt2 calls the method "Window.prompt".
//
// The returned bool will be false if there is no such method.
func (this Window) Prompt2() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWindowPrompt2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// Prompt2Func returns the method "Window.prompt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Prompt2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.WindowPrompt2Func(
			this.Ref(),
		),
	)
}

// Print calls the method "Window.print".
//
// The returned bool will be false if there is no such method.
func (this Window) Print() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowPrint(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PrintFunc returns the method "Window.print".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PrintFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowPrintFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "Window.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Window) PostMessage(message js.Any, targetOrigin js.String, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		targetOrigin.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "Window.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PostMessageFunc() (fn js.Func[func(message js.Any, targetOrigin js.String, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.WindowPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Window.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Window) PostMessage1(message js.Any, targetOrigin js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		targetOrigin.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "Window.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PostMessage1Func() (fn js.Func[func(message js.Any, targetOrigin js.String)]) {
	return fn.FromRef(
		bindings.WindowPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Window.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Window) PostMessage2(message js.Any, options WindowPostMessageOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowPostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "Window.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PostMessage2Func() (fn js.Func[func(message js.Any, options WindowPostMessageOptions)]) {
	return fn.FromRef(
		bindings.WindowPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage3 calls the method "Window.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Window) PostMessage3(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowPostMessage3(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage3Func returns the method "Window.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) PostMessage3Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.WindowPostMessage3Func(
			this.Ref(),
		),
	)
}

// GetSelection calls the method "Window.getSelection".
//
// The returned bool will be false if there is no such method.
func (this Window) GetSelection() (Selection, bool) {
	var _ok bool
	_ret := bindings.CallWindowGetSelection(
		this.Ref(), js.Pointer(&_ok),
	)

	return Selection{}.FromRef(_ret), _ok
}

// GetSelectionFunc returns the method "Window.getSelection".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetSelectionFunc() (fn js.Func[func() Selection]) {
	return fn.FromRef(
		bindings.WindowGetSelectionFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "Window.navigate".
//
// The returned bool will be false if there is no such method.
func (this Window) Navigate(dir SpatialNavigationDirection) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowNavigate(
		this.Ref(), js.Pointer(&_ok),
		uint32(dir),
	)

	_ = _ret
	return js.Void{}, _ok
}

// NavigateFunc returns the method "Window.navigate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) NavigateFunc() (fn js.Func[func(dir SpatialNavigationDirection)]) {
	return fn.FromRef(
		bindings.WindowNavigateFunc(
			this.Ref(),
		),
	)
}

// GetComputedStyle calls the method "Window.getComputedStyle".
//
// The returned bool will be false if there is no such method.
func (this Window) GetComputedStyle(elt Element, pseudoElt js.String) (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.CallWindowGetComputedStyle(
		this.Ref(), js.Pointer(&_ok),
		elt.Ref(),
		pseudoElt.Ref(),
	)

	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// GetComputedStyleFunc returns the method "Window.getComputedStyle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetComputedStyleFunc() (fn js.Func[func(elt Element, pseudoElt js.String) CSSStyleDeclaration]) {
	return fn.FromRef(
		bindings.WindowGetComputedStyleFunc(
			this.Ref(),
		),
	)
}

// GetComputedStyle1 calls the method "Window.getComputedStyle".
//
// The returned bool will be false if there is no such method.
func (this Window) GetComputedStyle1(elt Element) (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.CallWindowGetComputedStyle1(
		this.Ref(), js.Pointer(&_ok),
		elt.Ref(),
	)

	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// GetComputedStyle1Func returns the method "Window.getComputedStyle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetComputedStyle1Func() (fn js.Func[func(elt Element) CSSStyleDeclaration]) {
	return fn.FromRef(
		bindings.WindowGetComputedStyle1Func(
			this.Ref(),
		),
	)
}

// MatchMedia calls the method "Window.matchMedia".
//
// The returned bool will be false if there is no such method.
func (this Window) MatchMedia(query js.String) (MediaQueryList, bool) {
	var _ok bool
	_ret := bindings.CallWindowMatchMedia(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return MediaQueryList{}.FromRef(_ret), _ok
}

// MatchMediaFunc returns the method "Window.matchMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) MatchMediaFunc() (fn js.Func[func(query js.String) MediaQueryList]) {
	return fn.FromRef(
		bindings.WindowMatchMediaFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "Window.moveTo".
//
// The returned bool will be false if there is no such method.
func (this Window) MoveTo(x int32, y int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowMoveTo(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveToFunc returns the method "Window.moveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) MoveToFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowMoveToFunc(
			this.Ref(),
		),
	)
}

// MoveBy calls the method "Window.moveBy".
//
// The returned bool will be false if there is no such method.
func (this Window) MoveBy(x int32, y int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowMoveBy(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveByFunc returns the method "Window.moveBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) MoveByFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowMoveByFunc(
			this.Ref(),
		),
	)
}

// ResizeTo calls the method "Window.resizeTo".
//
// The returned bool will be false if there is no such method.
func (this Window) ResizeTo(width int32, height int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowResizeTo(
		this.Ref(), js.Pointer(&_ok),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResizeToFunc returns the method "Window.resizeTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ResizeToFunc() (fn js.Func[func(width int32, height int32)]) {
	return fn.FromRef(
		bindings.WindowResizeToFunc(
			this.Ref(),
		),
	)
}

// ResizeBy calls the method "Window.resizeBy".
//
// The returned bool will be false if there is no such method.
func (this Window) ResizeBy(x int32, y int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowResizeBy(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResizeByFunc returns the method "Window.resizeBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ResizeByFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowResizeByFunc(
			this.Ref(),
		),
	)
}

// Scroll calls the method "Window.scroll".
//
// The returned bool will be false if there is no such method.
func (this Window) Scroll(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScroll(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollFunc returns the method "Window.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollFunc(
			this.Ref(),
		),
	)
}

// Scroll1 calls the method "Window.scroll".
//
// The returned bool will be false if there is no such method.
func (this Window) Scroll1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScroll1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Scroll1Func returns the method "Window.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Scroll1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScroll1Func(
			this.Ref(),
		),
	)
}

// Scroll2 calls the method "Window.scroll".
//
// The returned bool will be false if there is no such method.
func (this Window) Scroll2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScroll2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Scroll2Func returns the method "Window.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Scroll2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScroll2Func(
			this.Ref(),
		),
	)
}

// ScrollTo calls the method "Window.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollTo(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollTo(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollToFunc returns the method "Window.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollToFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollToFunc(
			this.Ref(),
		),
	)
}

// ScrollTo1 calls the method "Window.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollTo1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollTo1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollTo1Func returns the method "Window.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollTo1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScrollTo1Func(
			this.Ref(),
		),
	)
}

// ScrollTo2 calls the method "Window.scrollTo".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollTo2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollTo2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollTo2Func returns the method "Window.scrollTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollTo2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScrollTo2Func(
			this.Ref(),
		),
	)
}

// ScrollBy calls the method "Window.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollBy(options ScrollToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollBy(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollByFunc returns the method "Window.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollByFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollByFunc(
			this.Ref(),
		),
	)
}

// ScrollBy1 calls the method "Window.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollBy1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollBy1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollBy1Func returns the method "Window.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollBy1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScrollBy1Func(
			this.Ref(),
		),
	)
}

// ScrollBy2 calls the method "Window.scrollBy".
//
// The returned bool will be false if there is no such method.
func (this Window) ScrollBy2(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowScrollBy2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollBy2Func returns the method "Window.scrollBy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ScrollBy2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScrollBy2Func(
			this.Ref(),
		),
	)
}

// RequestIdleCallback calls the method "Window.requestIdleCallback".
//
// The returned bool will be false if there is no such method.
func (this Window) RequestIdleCallback(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallWindowRequestIdleCallback(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
		js.Pointer(&options),
	)

	return uint32(_ret), _ok
}

// RequestIdleCallbackFunc returns the method "Window.requestIdleCallback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) RequestIdleCallbackFunc() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestIdleCallbackFunc(
			this.Ref(),
		),
	)
}

// RequestIdleCallback1 calls the method "Window.requestIdleCallback".
//
// The returned bool will be false if there is no such method.
func (this Window) RequestIdleCallback1(callback js.Func[func(deadline IdleDeadline)]) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallWindowRequestIdleCallback1(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return uint32(_ret), _ok
}

// RequestIdleCallback1Func returns the method "Window.requestIdleCallback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) RequestIdleCallback1Func() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)]) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestIdleCallback1Func(
			this.Ref(),
		),
	)
}

// CancelIdleCallback calls the method "Window.cancelIdleCallback".
//
// The returned bool will be false if there is no such method.
func (this Window) CancelIdleCallback(handle uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowCancelIdleCallback(
		this.Ref(), js.Pointer(&_ok),
		uint32(handle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelIdleCallbackFunc returns the method "Window.cancelIdleCallback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CancelIdleCallbackFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.WindowCancelIdleCallbackFunc(
			this.Ref(),
		),
	)
}

// ShowOpenFilePicker calls the method "Window.showOpenFilePicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowOpenFilePicker(options OpenFilePickerOptions) (js.Promise[js.Array[FileSystemFileHandle]], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowOpenFilePicker(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[FileSystemFileHandle]]{}.FromRef(_ret), _ok
}

// ShowOpenFilePickerFunc returns the method "Window.showOpenFilePicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowOpenFilePickerFunc() (fn js.Func[func(options OpenFilePickerOptions) js.Promise[js.Array[FileSystemFileHandle]]]) {
	return fn.FromRef(
		bindings.WindowShowOpenFilePickerFunc(
			this.Ref(),
		),
	)
}

// ShowOpenFilePicker1 calls the method "Window.showOpenFilePicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowOpenFilePicker1() (js.Promise[js.Array[FileSystemFileHandle]], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowOpenFilePicker1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[FileSystemFileHandle]]{}.FromRef(_ret), _ok
}

// ShowOpenFilePicker1Func returns the method "Window.showOpenFilePicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowOpenFilePicker1Func() (fn js.Func[func() js.Promise[js.Array[FileSystemFileHandle]]]) {
	return fn.FromRef(
		bindings.WindowShowOpenFilePicker1Func(
			this.Ref(),
		),
	)
}

// ShowSaveFilePicker calls the method "Window.showSaveFilePicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowSaveFilePicker(options SaveFilePickerOptions) (js.Promise[FileSystemFileHandle], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowSaveFilePicker(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[FileSystemFileHandle]{}.FromRef(_ret), _ok
}

// ShowSaveFilePickerFunc returns the method "Window.showSaveFilePicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowSaveFilePickerFunc() (fn js.Func[func(options SaveFilePickerOptions) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.WindowShowSaveFilePickerFunc(
			this.Ref(),
		),
	)
}

// ShowSaveFilePicker1 calls the method "Window.showSaveFilePicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowSaveFilePicker1() (js.Promise[FileSystemFileHandle], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowSaveFilePicker1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemFileHandle]{}.FromRef(_ret), _ok
}

// ShowSaveFilePicker1Func returns the method "Window.showSaveFilePicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowSaveFilePicker1Func() (fn js.Func[func() js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.WindowShowSaveFilePicker1Func(
			this.Ref(),
		),
	)
}

// ShowDirectoryPicker calls the method "Window.showDirectoryPicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowDirectoryPicker(options DirectoryPickerOptions) (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowDirectoryPicker(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// ShowDirectoryPickerFunc returns the method "Window.showDirectoryPicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowDirectoryPickerFunc() (fn js.Func[func(options DirectoryPickerOptions) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.WindowShowDirectoryPickerFunc(
			this.Ref(),
		),
	)
}

// ShowDirectoryPicker1 calls the method "Window.showDirectoryPicker".
//
// The returned bool will be false if there is no such method.
func (this Window) ShowDirectoryPicker1() (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallWindowShowDirectoryPicker1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// ShowDirectoryPicker1Func returns the method "Window.showDirectoryPicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ShowDirectoryPicker1Func() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.WindowShowDirectoryPicker1Func(
			this.Ref(),
		),
	)
}

// QueryLocalFonts calls the method "Window.queryLocalFonts".
//
// The returned bool will be false if there is no such method.
func (this Window) QueryLocalFonts(options QueryOptions) (js.Promise[js.Array[FontData]], bool) {
	var _ok bool
	_ret := bindings.CallWindowQueryLocalFonts(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[FontData]]{}.FromRef(_ret), _ok
}

// QueryLocalFontsFunc returns the method "Window.queryLocalFonts".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) QueryLocalFontsFunc() (fn js.Func[func(options QueryOptions) js.Promise[js.Array[FontData]]]) {
	return fn.FromRef(
		bindings.WindowQueryLocalFontsFunc(
			this.Ref(),
		),
	)
}

// QueryLocalFonts1 calls the method "Window.queryLocalFonts".
//
// The returned bool will be false if there is no such method.
func (this Window) QueryLocalFonts1() (js.Promise[js.Array[FontData]], bool) {
	var _ok bool
	_ret := bindings.CallWindowQueryLocalFonts1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[FontData]]{}.FromRef(_ret), _ok
}

// QueryLocalFonts1Func returns the method "Window.queryLocalFonts".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) QueryLocalFonts1Func() (fn js.Func[func() js.Promise[js.Array[FontData]]]) {
	return fn.FromRef(
		bindings.WindowQueryLocalFonts1Func(
			this.Ref(),
		),
	)
}

// GetScreenDetails calls the method "Window.getScreenDetails".
//
// The returned bool will be false if there is no such method.
func (this Window) GetScreenDetails() (js.Promise[ScreenDetails], bool) {
	var _ok bool
	_ret := bindings.CallWindowGetScreenDetails(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ScreenDetails]{}.FromRef(_ret), _ok
}

// GetScreenDetailsFunc returns the method "Window.getScreenDetails".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetScreenDetailsFunc() (fn js.Func[func() js.Promise[ScreenDetails]]) {
	return fn.FromRef(
		bindings.WindowGetScreenDetailsFunc(
			this.Ref(),
		),
	)
}

// GetDigitalGoodsService calls the method "Window.getDigitalGoodsService".
//
// The returned bool will be false if there is no such method.
func (this Window) GetDigitalGoodsService(serviceProvider js.String) (js.Promise[DigitalGoodsService], bool) {
	var _ok bool
	_ret := bindings.CallWindowGetDigitalGoodsService(
		this.Ref(), js.Pointer(&_ok),
		serviceProvider.Ref(),
	)

	return js.Promise[DigitalGoodsService]{}.FromRef(_ret), _ok
}

// GetDigitalGoodsServiceFunc returns the method "Window.getDigitalGoodsService".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) GetDigitalGoodsServiceFunc() (fn js.Func[func(serviceProvider js.String) js.Promise[DigitalGoodsService]]) {
	return fn.FromRef(
		bindings.WindowGetDigitalGoodsServiceFunc(
			this.Ref(),
		),
	)
}

// CaptureEvents calls the method "Window.captureEvents".
//
// The returned bool will be false if there is no such method.
func (this Window) CaptureEvents() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowCaptureEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CaptureEventsFunc returns the method "Window.captureEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CaptureEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowCaptureEventsFunc(
			this.Ref(),
		),
	)
}

// ReleaseEvents calls the method "Window.releaseEvents".
//
// The returned bool will be false if there is no such method.
func (this Window) ReleaseEvents() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowReleaseEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleaseEventsFunc returns the method "Window.releaseEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ReleaseEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowReleaseEventsFunc(
			this.Ref(),
		),
	)
}

// ReportError calls the method "Window.reportError".
//
// The returned bool will be false if there is no such method.
func (this Window) ReportError(e js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowReportError(
		this.Ref(), js.Pointer(&_ok),
		e.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReportErrorFunc returns the method "Window.reportError".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ReportErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.WindowReportErrorFunc(
			this.Ref(),
		),
	)
}

// Btoa calls the method "Window.btoa".
//
// The returned bool will be false if there is no such method.
func (this Window) Btoa(data js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWindowBtoa(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// BtoaFunc returns the method "Window.btoa".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) BtoaFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowBtoaFunc(
			this.Ref(),
		),
	)
}

// Atob calls the method "Window.atob".
//
// The returned bool will be false if there is no such method.
func (this Window) Atob(data js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWindowAtob(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// AtobFunc returns the method "Window.atob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) AtobFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowAtobFunc(
			this.Ref(),
		),
	)
}

// SetTimeout calls the method "Window.setTimeout".
//
// The returned bool will be false if there is no such method.
func (this Window) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWindowSetTimeout(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return int32(_ret), _ok
}

// SetTimeoutFunc returns the method "Window.setTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) SetTimeoutFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WindowSetTimeoutFunc(
			this.Ref(),
		),
	)
}

// SetTimeout1 calls the method "Window.setTimeout".
//
// The returned bool will be false if there is no such method.
func (this Window) SetTimeout1(handler TimerHandler) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWindowSetTimeout1(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
	)

	return int32(_ret), _ok
}

// SetTimeout1Func returns the method "Window.setTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) SetTimeout1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WindowSetTimeout1Func(
			this.Ref(),
		),
	)
}

// ClearTimeout calls the method "Window.clearTimeout".
//
// The returned bool will be false if there is no such method.
func (this Window) ClearTimeout(id int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowClearTimeout(
		this.Ref(), js.Pointer(&_ok),
		int32(id),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearTimeoutFunc returns the method "Window.clearTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ClearTimeoutFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WindowClearTimeoutFunc(
			this.Ref(),
		),
	)
}

// ClearTimeout1 calls the method "Window.clearTimeout".
//
// The returned bool will be false if there is no such method.
func (this Window) ClearTimeout1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowClearTimeout1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearTimeout1Func returns the method "Window.clearTimeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ClearTimeout1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowClearTimeout1Func(
			this.Ref(),
		),
	)
}

// SetInterval calls the method "Window.setInterval".
//
// The returned bool will be false if there is no such method.
func (this Window) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWindowSetInterval(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return int32(_ret), _ok
}

// SetIntervalFunc returns the method "Window.setInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) SetIntervalFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WindowSetIntervalFunc(
			this.Ref(),
		),
	)
}

// SetInterval1 calls the method "Window.setInterval".
//
// The returned bool will be false if there is no such method.
func (this Window) SetInterval1(handler TimerHandler) (int32, bool) {
	var _ok bool
	_ret := bindings.CallWindowSetInterval1(
		this.Ref(), js.Pointer(&_ok),
		handler.Ref(),
	)

	return int32(_ret), _ok
}

// SetInterval1Func returns the method "Window.setInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) SetInterval1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WindowSetInterval1Func(
			this.Ref(),
		),
	)
}

// ClearInterval calls the method "Window.clearInterval".
//
// The returned bool will be false if there is no such method.
func (this Window) ClearInterval(id int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowClearInterval(
		this.Ref(), js.Pointer(&_ok),
		int32(id),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearIntervalFunc returns the method "Window.clearInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ClearIntervalFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WindowClearIntervalFunc(
			this.Ref(),
		),
	)
}

// ClearInterval1 calls the method "Window.clearInterval".
//
// The returned bool will be false if there is no such method.
func (this Window) ClearInterval1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowClearInterval1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearInterval1Func returns the method "Window.clearInterval".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) ClearInterval1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowClearInterval1Func(
			this.Ref(),
		),
	)
}

// QueueMicrotask calls the method "Window.queueMicrotask".
//
// The returned bool will be false if there is no such method.
func (this Window) QueueMicrotask(callback js.Func[func()]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowQueueMicrotask(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QueueMicrotaskFunc returns the method "Window.queueMicrotask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) QueueMicrotaskFunc() (fn js.Func[func(callback js.Func[func()])]) {
	return fn.FromRef(
		bindings.WindowQueueMicrotaskFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap calls the method "Window.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this Window) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWindowCreateImageBitmap(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmapFunc returns the method "Window.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CreateImageBitmapFunc() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmapFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap1 calls the method "Window.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this Window) CreateImageBitmap1(image ImageBitmapSource) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWindowCreateImageBitmap1(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap1Func returns the method "Window.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CreateImageBitmap1Func() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap1Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap2 calls the method "Window.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this Window) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWindowCreateImageBitmap2(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap2Func returns the method "Window.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CreateImageBitmap2Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap2Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap3 calls the method "Window.createImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this Window) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (js.Promise[ImageBitmap], bool) {
	var _ok bool
	_ret := bindings.CallWindowCreateImageBitmap3(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return js.Promise[ImageBitmap]{}.FromRef(_ret), _ok
}

// CreateImageBitmap3Func returns the method "Window.createImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CreateImageBitmap3Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap3Func(
			this.Ref(),
		),
	)
}

// StructuredClone calls the method "Window.structuredClone".
//
// The returned bool will be false if there is no such method.
func (this Window) StructuredClone(value js.Any, options StructuredSerializeOptions) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWindowStructuredClone(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
		js.Pointer(&options),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// StructuredCloneFunc returns the method "Window.structuredClone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) StructuredCloneFunc() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	return fn.FromRef(
		bindings.WindowStructuredCloneFunc(
			this.Ref(),
		),
	)
}

// StructuredClone1 calls the method "Window.structuredClone".
//
// The returned bool will be false if there is no such method.
func (this Window) StructuredClone1(value js.Any) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWindowStructuredClone1(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// StructuredClone1Func returns the method "Window.structuredClone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) StructuredClone1Func() (fn js.Func[func(value js.Any) js.Any]) {
	return fn.FromRef(
		bindings.WindowStructuredClone1Func(
			this.Ref(),
		),
	)
}

// Fetch calls the method "Window.fetch".
//
// The returned bool will be false if there is no such method.
func (this Window) Fetch(input RequestInfo, init RequestInit) (js.Promise[Response], bool) {
	var _ok bool
	_ret := bindings.CallWindowFetch(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&init),
	)

	return js.Promise[Response]{}.FromRef(_ret), _ok
}

// FetchFunc returns the method "Window.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) FetchFunc() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WindowFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "Window.fetch".
//
// The returned bool will be false if there is no such method.
func (this Window) Fetch1(input RequestInfo) (js.Promise[Response], bool) {
	var _ok bool
	_ret := bindings.CallWindowFetch1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[Response]{}.FromRef(_ret), _ok
}

// Fetch1Func returns the method "Window.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) Fetch1Func() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WindowFetch1Func(
			this.Ref(),
		),
	)
}

// RequestAnimationFrame calls the method "Window.requestAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this Window) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallWindowRequestAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return uint32(_ret), _ok
}

// RequestAnimationFrameFunc returns the method "Window.requestAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) RequestAnimationFrameFunc() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// CancelAnimationFrame calls the method "Window.cancelAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this Window) CancelAnimationFrame(handle uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWindowCancelAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		uint32(handle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelAnimationFrameFunc returns the method "Window.cancelAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Window) CancelAnimationFrameFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.WindowCancelAnimationFrameFunc(
			this.Ref(),
		),
	)
}

type CompositionEventInit struct {
	// Data is "CompositionEventInit.data"
	//
	// Optional, defaults to "".
	Data js.String
	// View is "CompositionEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "CompositionEventInit.detail"
	//
	// Optional, defaults to 0.
	Detail int32
	// Bubbles is "CompositionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CompositionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CompositionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Detail     bool // for Detail.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CompositionEventInit with all fields set.
func (p CompositionEventInit) FromRef(ref js.Ref) CompositionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CompositionEventInit CompositionEventInit [// CompositionEventInit] [0x140005f0000 0x140007efa40 0x140007efae0 0x140007efc20 0x140007efd60 0x140007efea0 0x140007efb80 0x140007efcc0 0x140007efe00 0x140007eff40] 0x140005749f0 {0 0}} in the application heap.
func (p CompositionEventInit) New() js.Ref {
	return bindings.CompositionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CompositionEventInit) UpdateFrom(ref js.Ref) {
	bindings.CompositionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CompositionEventInit) Update(ref js.Ref) {
	bindings.CompositionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCompositionEvent(typ js.String, eventInitDict CompositionEventInit) CompositionEvent {
	return CompositionEvent{}.FromRef(
		bindings.NewCompositionEventByCompositionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewCompositionEventByCompositionEvent1(typ js.String) CompositionEvent {
	return CompositionEvent{}.FromRef(
		bindings.NewCompositionEventByCompositionEvent1(
			typ.Ref()),
	)
}

type CompositionEvent struct {
	UIEvent
}

func (this CompositionEvent) Once() CompositionEvent {
	this.Ref().Once()
	return this
}

func (this CompositionEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this CompositionEvent) FromRef(ref js.Ref) CompositionEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this CompositionEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "CompositionEvent.data".
//
// The returned bool will be false if there is no such property.
func (this CompositionEvent) Data() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCompositionEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InitCompositionEvent calls the method "CompositionEvent.initCompositionEvent".
//
// The returned bool will be false if there is no such method.
func (this CompositionEvent) InitCompositionEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCompositionEventInitCompositionEvent(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		dataArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCompositionEventFunc returns the method "CompositionEvent.initCompositionEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CompositionEvent) InitCompositionEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEventFunc(
			this.Ref(),
		),
	)
}

// InitCompositionEvent1 calls the method "CompositionEvent.initCompositionEvent".
//
// The returned bool will be false if there is no such method.
func (this CompositionEvent) InitCompositionEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCompositionEventInitCompositionEvent1(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCompositionEvent1Func returns the method "CompositionEvent.initCompositionEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CompositionEvent) InitCompositionEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent1Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent2 calls the method "CompositionEvent.initCompositionEvent".
//
// The returned bool will be false if there is no such method.
func (this CompositionEvent) InitCompositionEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCompositionEventInitCompositionEvent2(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCompositionEvent2Func returns the method "CompositionEvent.initCompositionEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CompositionEvent) InitCompositionEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent2Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent3 calls the method "CompositionEvent.initCompositionEvent".
//
// The returned bool will be false if there is no such method.
func (this CompositionEvent) InitCompositionEvent3(typeArg js.String, bubblesArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCompositionEventInitCompositionEvent3(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCompositionEvent3Func returns the method "CompositionEvent.initCompositionEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CompositionEvent) InitCompositionEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent3Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent4 calls the method "CompositionEvent.initCompositionEvent".
//
// The returned bool will be false if there is no such method.
func (this CompositionEvent) InitCompositionEvent4(typeArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCompositionEventInitCompositionEvent4(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCompositionEvent4Func returns the method "CompositionEvent.initCompositionEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CompositionEvent) InitCompositionEvent4Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent4Func(
			this.Ref(),
		),
	)
}

type CompressionFormat uint32

const (
	_ CompressionFormat = iota

	CompressionFormat_DEFLATE
	CompressionFormat_DEFLATE_RAW
	CompressionFormat_GZIP
)

func (CompressionFormat) FromRef(str js.Ref) CompressionFormat {
	return CompressionFormat(bindings.ConstOfCompressionFormat(str))
}

func (x CompressionFormat) String() (string, bool) {
	switch x {
	case CompressionFormat_DEFLATE:
		return "deflate", true
	case CompressionFormat_DEFLATE_RAW:
		return "deflate-raw", true
	case CompressionFormat_GZIP:
		return "gzip", true
	default:
		return "", false
	}
}

func NewCompressionStream(format CompressionFormat) CompressionStream {
	return CompressionStream{}.FromRef(
		bindings.NewCompressionStreamByCompressionStream(
			uint32(format)),
	)
}

type CompressionStream struct {
	ref js.Ref
}

func (this CompressionStream) Once() CompressionStream {
	this.Ref().Once()
	return this
}

func (this CompressionStream) Ref() js.Ref {
	return this.ref
}

func (this CompressionStream) FromRef(ref js.Ref) CompressionStream {
	this.ref = ref
	return this
}

func (this CompressionStream) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "CompressionStream.readable".
//
// The returned bool will be false if there is no such property.
func (this CompressionStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetCompressionStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "CompressionStream.writable".
//
// The returned bool will be false if there is no such property.
func (this CompressionStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetCompressionStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

type ContentIndexEventInit struct {
	// Id is "ContentIndexEventInit.id"
	//
	// Required
	Id js.String
	// Bubbles is "ContentIndexEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "ContentIndexEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "ContentIndexEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentIndexEventInit with all fields set.
func (p ContentIndexEventInit) FromRef(ref js.Ref) ContentIndexEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ContentIndexEventInit ContentIndexEventInit [// ContentIndexEventInit] [0x140007f40a0 0x140007f4140 0x140007f4280 0x140007f43c0 0x140007f41e0 0x140007f4320 0x140007f4460] 0x14000781068 {0 0}} in the application heap.
func (p ContentIndexEventInit) New() js.Ref {
	return bindings.ContentIndexEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ContentIndexEventInit) UpdateFrom(ref js.Ref) {
	bindings.ContentIndexEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ContentIndexEventInit) Update(ref js.Ref) {
	bindings.ContentIndexEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewContentIndexEvent(typ js.String, init ContentIndexEventInit) ContentIndexEvent {
	return ContentIndexEvent{}.FromRef(
		bindings.NewContentIndexEventByContentIndexEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
}

type ContentIndexEvent struct {
	ExtendableEvent
}

func (this ContentIndexEvent) Once() ContentIndexEvent {
	this.Ref().Once()
	return this
}

func (this ContentIndexEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this ContentIndexEvent) FromRef(ref js.Ref) ContentIndexEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this ContentIndexEvent) Free() {
	this.Ref().Free()
}

// Id returns the value of property "ContentIndexEvent.id".
//
// The returned bool will be false if there is no such property.
func (this ContentIndexEvent) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetContentIndexEventId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type ContentVisibilityAutoStateChangeEventInit struct {
	// Skipped is "ContentVisibilityAutoStateChangeEventInit.skipped"
	//
	// Optional, defaults to false.
	Skipped bool
	// Bubbles is "ContentVisibilityAutoStateChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "ContentVisibilityAutoStateChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "ContentVisibilityAutoStateChangeEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Skipped    bool // for Skipped.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentVisibilityAutoStateChangeEventInit with all fields set.
func (p ContentVisibilityAutoStateChangeEventInit) FromRef(ref js.Ref) ContentVisibilityAutoStateChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ContentVisibilityAutoStateChangeEventInit ContentVisibilityAutoStateChangeEventInit [// ContentVisibilityAutoStateChangeEventInit] [0x140007f4500 0x140007f4640 0x140007f4780 0x140007f48c0 0x140007f45a0 0x140007f46e0 0x140007f4820 0x140007f4960] 0x140007810c8 {0 0}} in the application heap.
func (p ContentVisibilityAutoStateChangeEventInit) New() js.Ref {
	return bindings.ContentVisibilityAutoStateChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ContentVisibilityAutoStateChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.ContentVisibilityAutoStateChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ContentVisibilityAutoStateChangeEventInit) Update(ref js.Ref) {
	bindings.ContentVisibilityAutoStateChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewContentVisibilityAutoStateChangeEvent(typ js.String, eventInitDict ContentVisibilityAutoStateChangeEventInit) ContentVisibilityAutoStateChangeEvent {
	return ContentVisibilityAutoStateChangeEvent{}.FromRef(
		bindings.NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent1(typ js.String) ContentVisibilityAutoStateChangeEvent {
	return ContentVisibilityAutoStateChangeEvent{}.FromRef(
		bindings.NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent1(
			typ.Ref()),
	)
}

type ContentVisibilityAutoStateChangeEvent struct {
	Event
}

func (this ContentVisibilityAutoStateChangeEvent) Once() ContentVisibilityAutoStateChangeEvent {
	this.Ref().Once()
	return this
}

func (this ContentVisibilityAutoStateChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ContentVisibilityAutoStateChangeEvent) FromRef(ref js.Ref) ContentVisibilityAutoStateChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ContentVisibilityAutoStateChangeEvent) Free() {
	this.Ref().Free()
}

// Skipped returns the value of property "ContentVisibilityAutoStateChangeEvent.skipped".
//
// The returned bool will be false if there is no such property.
func (this ContentVisibilityAutoStateChangeEvent) Skipped() (bool, bool) {
	var _ok bool
	_ret := bindings.GetContentVisibilityAutoStateChangeEventSkipped(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type CookieChangeEventInit struct {
	// Changed is "CookieChangeEventInit.changed"
	//
	// Optional
	Changed CookieList
	// Deleted is "CookieChangeEventInit.deleted"
	//
	// Optional
	Deleted CookieList
	// Bubbles is "CookieChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CookieChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CookieChangeEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieChangeEventInit with all fields set.
func (p CookieChangeEventInit) FromRef(ref js.Ref) CookieChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CookieChangeEventInit CookieChangeEventInit [// CookieChangeEventInit] [0x140007f4a00 0x140007f4aa0 0x140007f4b40 0x140007f4c80 0x140007f4dc0 0x140007f4be0 0x140007f4d20 0x140007f4e60] 0x14000781128 {0 0}} in the application heap.
func (p CookieChangeEventInit) New() js.Ref {
	return bindings.CookieChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CookieChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.CookieChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CookieChangeEventInit) Update(ref js.Ref) {
	bindings.CookieChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCookieChangeEvent(typ js.String, eventInitDict CookieChangeEventInit) CookieChangeEvent {
	return CookieChangeEvent{}.FromRef(
		bindings.NewCookieChangeEventByCookieChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewCookieChangeEventByCookieChangeEvent1(typ js.String) CookieChangeEvent {
	return CookieChangeEvent{}.FromRef(
		bindings.NewCookieChangeEventByCookieChangeEvent1(
			typ.Ref()),
	)
}

type CookieChangeEvent struct {
	Event
}

func (this CookieChangeEvent) Once() CookieChangeEvent {
	this.Ref().Once()
	return this
}

func (this CookieChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CookieChangeEvent) FromRef(ref js.Ref) CookieChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CookieChangeEvent) Free() {
	this.Ref().Free()
}

// Changed returns the value of property "CookieChangeEvent.changed".
//
// The returned bool will be false if there is no such property.
func (this CookieChangeEvent) Changed() (js.FrozenArray[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.GetCookieChangeEventChanged(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CookieListItem]{}.FromRef(_ret), _ok
}

// Deleted returns the value of property "CookieChangeEvent.deleted".
//
// The returned bool will be false if there is no such property.
func (this CookieChangeEvent) Deleted() (js.FrozenArray[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.GetCookieChangeEventDeleted(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CookieListItem]{}.FromRef(_ret), _ok
}

func NewCountQueuingStrategy(init QueuingStrategyInit) CountQueuingStrategy {
	return CountQueuingStrategy{}.FromRef(
		bindings.NewCountQueuingStrategyByCountQueuingStrategy(
			js.Pointer(&init)),
	)
}

type CountQueuingStrategy struct {
	ref js.Ref
}

func (this CountQueuingStrategy) Once() CountQueuingStrategy {
	this.Ref().Once()
	return this
}

func (this CountQueuingStrategy) Ref() js.Ref {
	return this.ref
}

func (this CountQueuingStrategy) FromRef(ref js.Ref) CountQueuingStrategy {
	this.ref = ref
	return this
}

func (this CountQueuingStrategy) Free() {
	this.Ref().Free()
}

// HighWaterMark returns the value of property "CountQueuingStrategy.highWaterMark".
//
// The returned bool will be false if there is no such property.
func (this CountQueuingStrategy) HighWaterMark() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCountQueuingStrategyHighWaterMark(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Size returns the value of property "CountQueuingStrategy.size".
//
// The returned bool will be false if there is no such property.
func (this CountQueuingStrategy) Size() (js.Func[func(arguments ...js.Any) js.Any], bool) {
	var _ok bool
	_ret := bindings.GetCountQueuingStrategySize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Func[func(arguments ...js.Any) js.Any]{}.FromRef(_ret), _ok
}

type CrashReportBody struct {
	ReportBody
}

func (this CrashReportBody) Once() CrashReportBody {
	this.Ref().Once()
	return this
}

func (this CrashReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this CrashReportBody) FromRef(ref js.Ref) CrashReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this CrashReportBody) Free() {
	this.Ref().Free()
}

// Reason returns the value of property "CrashReportBody.reason".
//
// The returned bool will be false if there is no such property.
func (this CrashReportBody) Reason() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCrashReportBodyReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "CrashReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this CrashReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallCrashReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "CrashReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CrashReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.CrashReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

type CredentialData struct {
	// Id is "CredentialData.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialData with all fields set.
func (p CredentialData) FromRef(ref js.Ref) CredentialData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CredentialData CredentialData [// CredentialData] [0x140007f5040] 0x14000781188 {0 0}} in the application heap.
func (p CredentialData) New() js.Ref {
	return bindings.CredentialDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CredentialData) UpdateFrom(ref js.Ref) {
	bindings.CredentialDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CredentialData) Update(ref js.Ref) {
	bindings.CredentialDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Uint64_BigIntUint64 struct {
	ref js.Ref
}

func (x OneOf_Uint64_BigIntUint64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Uint64_BigIntUint64) Free() {
	x.ref.Free()
}

func (x OneOf_Uint64_BigIntUint64) FromRef(ref js.Ref) OneOf_Uint64_BigIntUint64 {
	return OneOf_Uint64_BigIntUint64{
		ref: ref,
	}
}

func (x OneOf_Uint64_BigIntUint64) Uint64() uint64 {
	return js.BigInt[uint64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Uint64_BigIntUint64) BigIntUint64() js.BigInt[uint64] {
	return js.BigInt[uint64]{}.FromRef(x.ref)
}

type CryptoKeyID = OneOf_Uint64_BigIntUint64

type CryptoKeyPair struct {
	// PublicKey is "CryptoKeyPair.publicKey"
	//
	// Optional
	PublicKey CryptoKey
	// PrivateKey is "CryptoKeyPair.privateKey"
	//
	// Optional
	PrivateKey CryptoKey

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CryptoKeyPair with all fields set.
func (p CryptoKeyPair) FromRef(ref js.Ref) CryptoKeyPair {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CryptoKeyPair CryptoKeyPair [// CryptoKeyPair] [0x140007f50e0 0x140007f5180] 0x140007811b8 {0 0}} in the application heap.
func (p CryptoKeyPair) New() js.Ref {
	return bindings.CryptoKeyPairJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CryptoKeyPair) UpdateFrom(ref js.Ref) {
	bindings.CryptoKeyPairJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CryptoKeyPair) Update(ref js.Ref) {
	bindings.CryptoKeyPairJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CursorCaptureConstraint uint32

const (
	_ CursorCaptureConstraint = iota

	CursorCaptureConstraint_NEVER
	CursorCaptureConstraint_ALWAYS
	CursorCaptureConstraint_MOTION
)

func (CursorCaptureConstraint) FromRef(str js.Ref) CursorCaptureConstraint {
	return CursorCaptureConstraint(bindings.ConstOfCursorCaptureConstraint(str))
}

func (x CursorCaptureConstraint) String() (string, bool) {
	switch x {
	case CursorCaptureConstraint_NEVER:
		return "never", true
	case CursorCaptureConstraint_ALWAYS:
		return "always", true
	case CursorCaptureConstraint_MOTION:
		return "motion", true
	default:
		return "", false
	}
}

type CustomEventInit struct {
	// Detail is "CustomEventInit.detail"
	//
	// Optional, defaults to null.
	Detail js.Any
	// Bubbles is "CustomEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CustomEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CustomEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CustomEventInit with all fields set.
func (p CustomEventInit) FromRef(ref js.Ref) CustomEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CustomEventInit CustomEventInit [// CustomEventInit] [0x140007f5220 0x140007f52c0 0x140007f5400 0x140007f5540 0x140007f5360 0x140007f54a0 0x140007f55e0] 0x14000781200 {0 0}} in the application heap.
func (p CustomEventInit) New() js.Ref {
	return bindings.CustomEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CustomEventInit) UpdateFrom(ref js.Ref) {
	bindings.CustomEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CustomEventInit) Update(ref js.Ref) {
	bindings.CustomEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCustomEvent(typ js.String, eventInitDict CustomEventInit) CustomEvent {
	return CustomEvent{}.FromRef(
		bindings.NewCustomEventByCustomEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewCustomEventByCustomEvent1(typ js.String) CustomEvent {
	return CustomEvent{}.FromRef(
		bindings.NewCustomEventByCustomEvent1(
			typ.Ref()),
	)
}

type CustomEvent struct {
	Event
}

func (this CustomEvent) Once() CustomEvent {
	this.Ref().Once()
	return this
}

func (this CustomEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CustomEvent) FromRef(ref js.Ref) CustomEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CustomEvent) Free() {
	this.Ref().Free()
}

// Detail returns the value of property "CustomEvent.detail".
//
// The returned bool will be false if there is no such property.
func (this CustomEvent) Detail() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetCustomEventDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// InitCustomEvent calls the method "CustomEvent.initCustomEvent".
//
// The returned bool will be false if there is no such method.
func (this CustomEvent) InitCustomEvent(typ js.String, bubbles bool, cancelable bool, detail js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomEventInitCustomEvent(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		detail.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCustomEventFunc returns the method "CustomEvent.initCustomEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomEvent) InitCustomEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, detail js.Any)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEventFunc(
			this.Ref(),
		),
	)
}

// InitCustomEvent1 calls the method "CustomEvent.initCustomEvent".
//
// The returned bool will be false if there is no such method.
func (this CustomEvent) InitCustomEvent1(typ js.String, bubbles bool, cancelable bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomEventInitCustomEvent1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCustomEvent1Func returns the method "CustomEvent.initCustomEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomEvent) InitCustomEvent1Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent1Func(
			this.Ref(),
		),
	)
}

// InitCustomEvent2 calls the method "CustomEvent.initCustomEvent".
//
// The returned bool will be false if there is no such method.
func (this CustomEvent) InitCustomEvent2(typ js.String, bubbles bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomEventInitCustomEvent2(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCustomEvent2Func returns the method "CustomEvent.initCustomEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomEvent) InitCustomEvent2Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent2Func(
			this.Ref(),
		),
	)
}

// InitCustomEvent3 calls the method "CustomEvent.initCustomEvent".
//
// The returned bool will be false if there is no such method.
func (this CustomEvent) InitCustomEvent3(typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomEventInitCustomEvent3(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitCustomEvent3Func returns the method "CustomEvent.initCustomEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomEvent) InitCustomEvent3Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent3Func(
			this.Ref(),
		),
	)
}

type DOMParserSupportedType uint32

const (
	_ DOMParserSupportedType = iota

	DOMParserSupportedType_TEXT_HTML
	DOMParserSupportedType_TEXT_XML
	DOMParserSupportedType_APPLICATION_XML
	DOMParserSupportedType_APPLICATION_XHTML_XML
	DOMParserSupportedType_IMAGE_SVG_XML
)

func (DOMParserSupportedType) FromRef(str js.Ref) DOMParserSupportedType {
	return DOMParserSupportedType(bindings.ConstOfDOMParserSupportedType(str))
}

func (x DOMParserSupportedType) String() (string, bool) {
	switch x {
	case DOMParserSupportedType_TEXT_HTML:
		return "text/html", true
	case DOMParserSupportedType_TEXT_XML:
		return "text/xml", true
	case DOMParserSupportedType_APPLICATION_XML:
		return "application/xml", true
	case DOMParserSupportedType_APPLICATION_XHTML_XML:
		return "application/xhtml+xml", true
	case DOMParserSupportedType_IMAGE_SVG_XML:
		return "image/svg+xml", true
	default:
		return "", false
	}
}

type DOMParser struct {
	ref js.Ref
}

func (this DOMParser) Once() DOMParser {
	this.Ref().Once()
	return this
}

func (this DOMParser) Ref() js.Ref {
	return this.ref
}

func (this DOMParser) FromRef(ref js.Ref) DOMParser {
	this.ref = ref
	return this
}

func (this DOMParser) Free() {
	this.Ref().Free()
}

// ParseFromString calls the method "DOMParser.parseFromString".
//
// The returned bool will be false if there is no such method.
func (this DOMParser) ParseFromString(string js.String, typ DOMParserSupportedType) (Document, bool) {
	var _ok bool
	_ret := bindings.CallDOMParserParseFromString(
		this.Ref(), js.Pointer(&_ok),
		string.Ref(),
		uint32(typ),
	)

	return Document{}.FromRef(_ret), _ok
}

// ParseFromStringFunc returns the method "DOMParser.parseFromString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMParser) ParseFromStringFunc() (fn js.Func[func(string js.String, typ DOMParserSupportedType) Document]) {
	return fn.FromRef(
		bindings.DOMParserParseFromStringFunc(
			this.Ref(),
		),
	)
}

func NewDataCue(startTime float64, endTime float64, value js.Any, typ js.String) DataCue {
	return DataCue{}.FromRef(
		bindings.NewDataCueByDataCue(
			float64(startTime),
			float64(endTime),
			value.Ref(),
			typ.Ref()),
	)
}

func NewDataCueByDataCue1(startTime float64, endTime float64, value js.Any) DataCue {
	return DataCue{}.FromRef(
		bindings.NewDataCueByDataCue1(
			float64(startTime),
			float64(endTime),
			value.Ref()),
	)
}

type DataCue struct {
	TextTrackCue
}

func (this DataCue) Once() DataCue {
	this.Ref().Once()
	return this
}

func (this DataCue) Ref() js.Ref {
	return this.TextTrackCue.Ref()
}

func (this DataCue) FromRef(ref js.Ref) DataCue {
	this.TextTrackCue = this.TextTrackCue.FromRef(ref)
	return this
}

func (this DataCue) Free() {
	this.Ref().Free()
}

// Value returns the value of property "DataCue.value".
//
// The returned bool will be false if there is no such property.
func (this DataCue) Value() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetDataCueValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Value sets the value of property "DataCue.value" to val.
//
// It returns false if the property cannot be set.
func (this DataCue) SetValue(val js.Any) bool {
	return js.True == bindings.SetDataCueValue(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "DataCue.type".
//
// The returned bool will be false if there is no such property.
func (this DataCue) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDataCueType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

func NewDecompressionStream(format CompressionFormat) DecompressionStream {
	return DecompressionStream{}.FromRef(
		bindings.NewDecompressionStreamByDecompressionStream(
			uint32(format)),
	)
}

type DecompressionStream struct {
	ref js.Ref
}

func (this DecompressionStream) Once() DecompressionStream {
	this.Ref().Once()
	return this
}

func (this DecompressionStream) Ref() js.Ref {
	return this.ref
}

func (this DecompressionStream) FromRef(ref js.Ref) DecompressionStream {
	this.ref = ref
	return this
}

func (this DecompressionStream) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "DecompressionStream.readable".
//
// The returned bool will be false if there is no such property.
func (this DecompressionStream) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetDecompressionStreamReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Writable returns the value of property "DecompressionStream.writable".
//
// The returned bool will be false if there is no such property.
func (this DecompressionStream) Writable() (WritableStream, bool) {
	var _ok bool
	_ret := bindings.GetDecompressionStreamWritable(
		this.Ref(), js.Pointer(&_ok),
	)
	return WritableStream{}.FromRef(_ret), _ok
}

type DedicatedWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this DedicatedWorkerGlobalScope) Once() DedicatedWorkerGlobalScope {
	this.Ref().Once()
	return this
}

func (this DedicatedWorkerGlobalScope) Ref() js.Ref {
	return this.WorkerGlobalScope.Ref()
}

func (this DedicatedWorkerGlobalScope) FromRef(ref js.Ref) DedicatedWorkerGlobalScope {
	this.WorkerGlobalScope = this.WorkerGlobalScope.FromRef(ref)
	return this
}

func (this DedicatedWorkerGlobalScope) Free() {
	this.Ref().Free()
}

// Name returns the value of property "DedicatedWorkerGlobalScope.name".
//
// The returned bool will be false if there is no such property.
func (this DedicatedWorkerGlobalScope) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDedicatedWorkerGlobalScopeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PostMessage calls the method "DedicatedWorkerGlobalScope.postMessage".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessage(message js.Any, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopePostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "DedicatedWorkerGlobalScope.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "DedicatedWorkerGlobalScope.postMessage".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessage1(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopePostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "DedicatedWorkerGlobalScope.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "DedicatedWorkerGlobalScope.postMessage".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessage2(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopePostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "DedicatedWorkerGlobalScope.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessage2Func(
			this.Ref(),
		),
	)
}

// Close calls the method "DedicatedWorkerGlobalScope.close".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopeClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "DedicatedWorkerGlobalScope.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeCloseFunc(
			this.Ref(),
		),
	)
}

// RequestAnimationFrame calls the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return uint32(_ret), _ok
}

// RequestAnimationFrameFunc returns the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) RequestAnimationFrameFunc() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeRequestAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// CancelAnimationFrame calls the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this DedicatedWorkerGlobalScope) CancelAnimationFrame(handle uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		uint32(handle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelAnimationFrameFunc returns the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DedicatedWorkerGlobalScope) CancelAnimationFrameFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeCancelAnimationFrameFunc(
			this.Ref(),
		),
	)
}

type DeprecationReportBody struct {
	ReportBody
}

func (this DeprecationReportBody) Once() DeprecationReportBody {
	this.Ref().Once()
	return this
}

func (this DeprecationReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this DeprecationReportBody) FromRef(ref js.Ref) DeprecationReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this DeprecationReportBody) Free() {
	this.Ref().Free()
}

// Id returns the value of property "DeprecationReportBody.id".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodyId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AnticipatedRemoval returns the value of property "DeprecationReportBody.anticipatedRemoval".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) AnticipatedRemoval() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodyAnticipatedRemoval(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Message returns the value of property "DeprecationReportBody.message".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodyMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SourceFile returns the value of property "DeprecationReportBody.sourceFile".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) SourceFile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodySourceFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LineNumber returns the value of property "DeprecationReportBody.lineNumber".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) LineNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodyLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColumnNumber returns the value of property "DeprecationReportBody.columnNumber".
//
// The returned bool will be false if there is no such property.
func (this DeprecationReportBody) ColumnNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDeprecationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ToJSON calls the method "DeprecationReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this DeprecationReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDeprecationReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "DeprecationReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DeprecationReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DeprecationReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

type LandmarkType uint32

const (
	_ LandmarkType = iota

	LandmarkType_MOUTH
	LandmarkType_EYE
	LandmarkType_NOSE
)

func (LandmarkType) FromRef(str js.Ref) LandmarkType {
	return LandmarkType(bindings.ConstOfLandmarkType(str))
}

func (x LandmarkType) String() (string, bool) {
	switch x {
	case LandmarkType_MOUTH:
		return "mouth", true
	case LandmarkType_EYE:
		return "eye", true
	case LandmarkType_NOSE:
		return "nose", true
	default:
		return "", false
	}
}

type Landmark struct {
	// Locations is "Landmark.locations"
	//
	// Required
	Locations js.FrozenArray[Point2D]
	// Type is "Landmark.type"
	//
	// Optional
	Type LandmarkType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Landmark with all fields set.
func (p Landmark) FromRef(ref js.Ref) Landmark {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 Landmark Landmark [// Landmark] [0x140007f5900 0x140007f59a0] 0x14000781350 {0 0}} in the application heap.
func (p Landmark) New() js.Ref {
	return bindings.LandmarkJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p Landmark) UpdateFrom(ref js.Ref) {
	bindings.LandmarkJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p Landmark) Update(ref js.Ref) {
	bindings.LandmarkJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DetectedFace struct {
	// BoundingBox is "DetectedFace.boundingBox"
	//
	// Required
	BoundingBox DOMRectReadOnly
	// Landmarks is "DetectedFace.landmarks"
	//
	// Required
	Landmarks js.FrozenArray[Landmark]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DetectedFace with all fields set.
func (p DetectedFace) FromRef(ref js.Ref) DetectedFace {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DetectedFace DetectedFace [// DetectedFace] [0x140007f5860 0x140007f5a40] 0x14000781338 {0 0}} in the application heap.
func (p DetectedFace) New() js.Ref {
	return bindings.DetectedFaceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DetectedFace) UpdateFrom(ref js.Ref) {
	bindings.DetectedFaceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DetectedFace) Update(ref js.Ref) {
	bindings.DetectedFaceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DetectedText struct {
	// BoundingBox is "DetectedText.boundingBox"
	//
	// Required
	BoundingBox DOMRectReadOnly
	// RawValue is "DetectedText.rawValue"
	//
	// Required
	RawValue js.String
	// CornerPoints is "DetectedText.cornerPoints"
	//
	// Required
	CornerPoints js.FrozenArray[Point2D]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DetectedText with all fields set.
func (p DetectedText) FromRef(ref js.Ref) DetectedText {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DetectedText DetectedText [// DetectedText] [0x140007f5ae0 0x140007f5b80 0x140007f5c20] 0x14000781380 {0 0}} in the application heap.
func (p DetectedText) New() js.Ref {
	return bindings.DetectedTextJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DetectedText) UpdateFrom(ref js.Ref) {
	bindings.DetectedTextJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DetectedText) Update(ref js.Ref) {
	bindings.DetectedTextJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DeviceMotionEventAccelerationInit struct {
	// X is "DeviceMotionEventAccelerationInit.x"
	//
	// Optional, defaults to null.
	X float64
	// Y is "DeviceMotionEventAccelerationInit.y"
	//
	// Optional, defaults to null.
	Y float64
	// Z is "DeviceMotionEventAccelerationInit.z"
	//
	// Optional, defaults to null.
	Z float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.
	FFI_USE_Z bool // for Z.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceMotionEventAccelerationInit with all fields set.
func (p DeviceMotionEventAccelerationInit) FromRef(ref js.Ref) DeviceMotionEventAccelerationInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DeviceMotionEventAccelerationInit DeviceMotionEventAccelerationInit [// DeviceMotionEventAccelerationInit] [0x140007f5cc0 0x140007f5e00 0x140007f5f40 0x140007f5d60 0x140007f5ea0 0x140007fa000] 0x140007813e0 {0 0}} in the application heap.
func (p DeviceMotionEventAccelerationInit) New() js.Ref {
	return bindings.DeviceMotionEventAccelerationInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DeviceMotionEventAccelerationInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventAccelerationInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DeviceMotionEventAccelerationInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventAccelerationInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DeviceMotionEventRotationRateInit struct {
	// Alpha is "DeviceMotionEventRotationRateInit.alpha"
	//
	// Optional, defaults to null.
	Alpha float64
	// Beta is "DeviceMotionEventRotationRateInit.beta"
	//
	// Optional, defaults to null.
	Beta float64
	// Gamma is "DeviceMotionEventRotationRateInit.gamma"
	//
	// Optional, defaults to null.
	Gamma float64

	FFI_USE_Alpha bool // for Alpha.
	FFI_USE_Beta  bool // for Beta.
	FFI_USE_Gamma bool // for Gamma.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceMotionEventRotationRateInit with all fields set.
func (p DeviceMotionEventRotationRateInit) FromRef(ref js.Ref) DeviceMotionEventRotationRateInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DeviceMotionEventRotationRateInit DeviceMotionEventRotationRateInit [// DeviceMotionEventRotationRateInit] [0x140007fa1e0 0x140007fa320 0x140007fa460 0x140007fa280 0x140007fa3c0 0x140007fa500] 0x14000781440 {0 0}} in the application heap.
func (p DeviceMotionEventRotationRateInit) New() js.Ref {
	return bindings.DeviceMotionEventRotationRateInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DeviceMotionEventRotationRateInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventRotationRateInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DeviceMotionEventRotationRateInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventRotationRateInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DeviceMotionEventInit struct {
	// Acceleration is "DeviceMotionEventInit.acceleration"
	//
	// Optional
	Acceleration DeviceMotionEventAccelerationInit
	// AccelerationIncludingGravity is "DeviceMotionEventInit.accelerationIncludingGravity"
	//
	// Optional
	AccelerationIncludingGravity DeviceMotionEventAccelerationInit
	// RotationRate is "DeviceMotionEventInit.rotationRate"
	//
	// Optional
	RotationRate DeviceMotionEventRotationRateInit
	// Interval is "DeviceMotionEventInit.interval"
	//
	// Optional, defaults to 0.
	Interval float64
	// Bubbles is "DeviceMotionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "DeviceMotionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "DeviceMotionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Interval   bool // for Interval.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceMotionEventInit with all fields set.
func (p DeviceMotionEventInit) FromRef(ref js.Ref) DeviceMotionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DeviceMotionEventInit DeviceMotionEventInit [// DeviceMotionEventInit] [0x140007fa0a0 0x140007fa140 0x140007fa5a0 0x140007fa640 0x140007fa780 0x140007fa8c0 0x140007faa00 0x140007fa6e0 0x140007fa820 0x140007fa960 0x140007faaa0] 0x140007813c8 {0 0}} in the application heap.
func (p DeviceMotionEventInit) New() js.Ref {
	return bindings.DeviceMotionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DeviceMotionEventInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DeviceMotionEventInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DeviceMotionEventAcceleration struct {
	ref js.Ref
}

func (this DeviceMotionEventAcceleration) Once() DeviceMotionEventAcceleration {
	this.Ref().Once()
	return this
}

func (this DeviceMotionEventAcceleration) Ref() js.Ref {
	return this.ref
}

func (this DeviceMotionEventAcceleration) FromRef(ref js.Ref) DeviceMotionEventAcceleration {
	this.ref = ref
	return this
}

func (this DeviceMotionEventAcceleration) Free() {
	this.Ref().Free()
}

// X returns the value of property "DeviceMotionEventAcceleration.x".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventAcceleration) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventAccelerationX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "DeviceMotionEventAcceleration.y".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventAcceleration) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventAccelerationY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "DeviceMotionEventAcceleration.z".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventAcceleration) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventAccelerationZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type DeviceMotionEventRotationRate struct {
	ref js.Ref
}

func (this DeviceMotionEventRotationRate) Once() DeviceMotionEventRotationRate {
	this.Ref().Once()
	return this
}

func (this DeviceMotionEventRotationRate) Ref() js.Ref {
	return this.ref
}

func (this DeviceMotionEventRotationRate) FromRef(ref js.Ref) DeviceMotionEventRotationRate {
	this.ref = ref
	return this
}

func (this DeviceMotionEventRotationRate) Free() {
	this.Ref().Free()
}

// Alpha returns the value of property "DeviceMotionEventRotationRate.alpha".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventRotationRate) Alpha() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventRotationRateAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Beta returns the value of property "DeviceMotionEventRotationRate.beta".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventRotationRate) Beta() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventRotationRateBeta(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Gamma returns the value of property "DeviceMotionEventRotationRate.gamma".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEventRotationRate) Gamma() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventRotationRateGamma(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

func NewDeviceMotionEvent(typ js.String, eventInitDict DeviceMotionEventInit) DeviceMotionEvent {
	return DeviceMotionEvent{}.FromRef(
		bindings.NewDeviceMotionEventByDeviceMotionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewDeviceMotionEventByDeviceMotionEvent1(typ js.String) DeviceMotionEvent {
	return DeviceMotionEvent{}.FromRef(
		bindings.NewDeviceMotionEventByDeviceMotionEvent1(
			typ.Ref()),
	)
}

type DeviceMotionEvent struct {
	Event
}

func (this DeviceMotionEvent) Once() DeviceMotionEvent {
	this.Ref().Once()
	return this
}

func (this DeviceMotionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this DeviceMotionEvent) FromRef(ref js.Ref) DeviceMotionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this DeviceMotionEvent) Free() {
	this.Ref().Free()
}

// Acceleration returns the value of property "DeviceMotionEvent.acceleration".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEvent) Acceleration() (DeviceMotionEventAcceleration, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventAcceleration(
		this.Ref(), js.Pointer(&_ok),
	)
	return DeviceMotionEventAcceleration{}.FromRef(_ret), _ok
}

// AccelerationIncludingGravity returns the value of property "DeviceMotionEvent.accelerationIncludingGravity".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEvent) AccelerationIncludingGravity() (DeviceMotionEventAcceleration, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventAccelerationIncludingGravity(
		this.Ref(), js.Pointer(&_ok),
	)
	return DeviceMotionEventAcceleration{}.FromRef(_ret), _ok
}

// RotationRate returns the value of property "DeviceMotionEvent.rotationRate".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEvent) RotationRate() (DeviceMotionEventRotationRate, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventRotationRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return DeviceMotionEventRotationRate{}.FromRef(_ret), _ok
}

// Interval returns the value of property "DeviceMotionEvent.interval".
//
// The returned bool will be false if there is no such property.
func (this DeviceMotionEvent) Interval() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDeviceMotionEventInterval(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RequestPermission calls the staticmethod "DeviceMotionEvent.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this DeviceMotionEvent) RequestPermission() (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallDeviceMotionEventRequestPermission(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// RequestPermissionFunc returns the staticmethod "DeviceMotionEvent.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DeviceMotionEvent) RequestPermissionFunc() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.DeviceMotionEventRequestPermissionFunc(
			this.Ref(),
		),
	)
}

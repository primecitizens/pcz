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
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
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

// New creates a new PerformanceMeasureOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this PerformanceEntry) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EntryType returns the value of property "PerformanceEntry.entryType".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) EntryType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryEntryType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StartTime returns the value of property "PerformanceEntry.startTime".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) StartTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryStartTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "PerformanceEntry.duration".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) Duration() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceEntry.toJSON" exists.
func (this PerformanceEntry) HasToJSON() bool {
	return js.True == bindings.HasPerformanceEntryToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceEntry.toJSON".
func (this PerformanceEntry) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceEntryToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceEntry.toJSON".
func (this PerformanceEntry) ToJSON() (ret js.Object) {
	bindings.CallPerformanceEntryToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceEntry.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PerformanceEntry) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceEntryToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this PerformanceTiming) NavigationStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingNavigationStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnloadEventStart returns the value of property "PerformanceTiming.unloadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) UnloadEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingUnloadEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnloadEventEnd returns the value of property "PerformanceTiming.unloadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) UnloadEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingUnloadEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectStart returns the value of property "PerformanceTiming.redirectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RedirectStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRedirectStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectEnd returns the value of property "PerformanceTiming.redirectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RedirectEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRedirectEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FetchStart returns the value of property "PerformanceTiming.fetchStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) FetchStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingFetchStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomainLookupStart returns the value of property "PerformanceTiming.domainLookupStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomainLookupStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomainLookupStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomainLookupEnd returns the value of property "PerformanceTiming.domainLookupEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomainLookupEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomainLookupEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ConnectStart returns the value of property "PerformanceTiming.connectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ConnectStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingConnectStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ConnectEnd returns the value of property "PerformanceTiming.connectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ConnectEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingConnectEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SecureConnectionStart returns the value of property "PerformanceTiming.secureConnectionStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) SecureConnectionStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingSecureConnectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RequestStart returns the value of property "PerformanceTiming.requestStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RequestStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRequestStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseStart returns the value of property "PerformanceTiming.responseStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ResponseStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingResponseStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseEnd returns the value of property "PerformanceTiming.responseEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ResponseEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingResponseEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomLoading returns the value of property "PerformanceTiming.domLoading".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomLoading() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomLoading(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomInteractive returns the value of property "PerformanceTiming.domInteractive".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomInteractive() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomInteractive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventStart returns the value of property "PerformanceTiming.domContentLoadedEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomContentLoadedEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventEnd returns the value of property "PerformanceTiming.domContentLoadedEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomContentLoadedEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomComplete returns the value of property "PerformanceTiming.domComplete".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomComplete() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadEventStart returns the value of property "PerformanceTiming.loadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) LoadEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingLoadEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadEventEnd returns the value of property "PerformanceTiming.loadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) LoadEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingLoadEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceTiming.toJSON" exists.
func (this PerformanceTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceTiming.toJSON".
func (this PerformanceTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceTiming.toJSON".
func (this PerformanceTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PerformanceTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this PerformanceNavigation) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectCount returns the value of property "PerformanceNavigation.redirectCount".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigation) RedirectCount() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationRedirectCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceNavigation.toJSON" exists.
func (this PerformanceNavigation) HasToJSON() bool {
	return js.True == bindings.HasPerformanceNavigationToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceNavigation.toJSON".
func (this PerformanceNavigation) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceNavigationToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceNavigation.toJSON".
func (this PerformanceNavigation) ToJSON() (ret js.Object) {
	bindings.CallPerformanceNavigationToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceNavigation.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PerformanceNavigation) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceNavigationToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Performance) TimeOrigin() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceTimeOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timing returns the value of property "Performance.timing".
//
// It returns ok=false if there is no such property.
func (this Performance) Timing() (ret PerformanceTiming, ok bool) {
	ok = js.True == bindings.GetPerformanceTiming(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Navigation returns the value of property "Performance.navigation".
//
// It returns ok=false if there is no such property.
func (this Performance) Navigation() (ret PerformanceNavigation, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EventCounts returns the value of property "Performance.eventCounts".
//
// It returns ok=false if there is no such property.
func (this Performance) EventCounts() (ret EventCounts, ok bool) {
	ok = js.True == bindings.GetPerformanceEventCounts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InteractionCount returns the value of property "Performance.interactionCount".
//
// It returns ok=false if there is no such property.
func (this Performance) InteractionCount() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceInteractionCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasNow returns true if the method "Performance.now" exists.
func (this Performance) HasNow() bool {
	return js.True == bindings.HasPerformanceNow(
		this.Ref(),
	)
}

// NowFunc returns the method "Performance.now".
func (this Performance) NowFunc() (fn js.Func[func() DOMHighResTimeStamp]) {
	return fn.FromRef(
		bindings.PerformanceNowFunc(
			this.Ref(),
		),
	)
}

// Now calls the method "Performance.now".
func (this Performance) Now() (ret DOMHighResTimeStamp) {
	bindings.CallPerformanceNow(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryNow calls the method "Performance.now"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryNow() (ret DOMHighResTimeStamp, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceNow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "Performance.toJSON" exists.
func (this Performance) HasToJSON() bool {
	return js.True == bindings.HasPerformanceToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "Performance.toJSON".
func (this Performance) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "Performance.toJSON".
func (this Performance) ToJSON() (ret js.Object) {
	bindings.CallPerformanceToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "Performance.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMeasureUserAgentSpecificMemory returns true if the method "Performance.measureUserAgentSpecificMemory" exists.
func (this Performance) HasMeasureUserAgentSpecificMemory() bool {
	return js.True == bindings.HasPerformanceMeasureUserAgentSpecificMemory(
		this.Ref(),
	)
}

// MeasureUserAgentSpecificMemoryFunc returns the method "Performance.measureUserAgentSpecificMemory".
func (this Performance) MeasureUserAgentSpecificMemoryFunc() (fn js.Func[func() js.Promise[MemoryMeasurement]]) {
	return fn.FromRef(
		bindings.PerformanceMeasureUserAgentSpecificMemoryFunc(
			this.Ref(),
		),
	)
}

// MeasureUserAgentSpecificMemory calls the method "Performance.measureUserAgentSpecificMemory".
func (this Performance) MeasureUserAgentSpecificMemory() (ret js.Promise[MemoryMeasurement]) {
	bindings.CallPerformanceMeasureUserAgentSpecificMemory(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMeasureUserAgentSpecificMemory calls the method "Performance.measureUserAgentSpecificMemory"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMeasureUserAgentSpecificMemory() (ret js.Promise[MemoryMeasurement], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasureUserAgentSpecificMemory(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMark returns true if the method "Performance.mark" exists.
func (this Performance) HasMark() bool {
	return js.True == bindings.HasPerformanceMark(
		this.Ref(),
	)
}

// MarkFunc returns the method "Performance.mark".
func (this Performance) MarkFunc() (fn js.Func[func(markName js.String, markOptions PerformanceMarkOptions) PerformanceMark]) {
	return fn.FromRef(
		bindings.PerformanceMarkFunc(
			this.Ref(),
		),
	)
}

// Mark calls the method "Performance.mark".
func (this Performance) Mark(markName js.String, markOptions PerformanceMarkOptions) (ret PerformanceMark) {
	bindings.CallPerformanceMark(
		this.Ref(), js.Pointer(&ret),
		markName.Ref(),
		js.Pointer(&markOptions),
	)

	return
}

// TryMark calls the method "Performance.mark"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMark(markName js.String, markOptions PerformanceMarkOptions) (ret PerformanceMark, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMark(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
		js.Pointer(&markOptions),
	)

	return
}

// HasMark1 returns true if the method "Performance.mark" exists.
func (this Performance) HasMark1() bool {
	return js.True == bindings.HasPerformanceMark1(
		this.Ref(),
	)
}

// Mark1Func returns the method "Performance.mark".
func (this Performance) Mark1Func() (fn js.Func[func(markName js.String) PerformanceMark]) {
	return fn.FromRef(
		bindings.PerformanceMark1Func(
			this.Ref(),
		),
	)
}

// Mark1 calls the method "Performance.mark".
func (this Performance) Mark1(markName js.String) (ret PerformanceMark) {
	bindings.CallPerformanceMark1(
		this.Ref(), js.Pointer(&ret),
		markName.Ref(),
	)

	return
}

// TryMark1 calls the method "Performance.mark"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMark1(markName js.String) (ret PerformanceMark, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMark1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
	)

	return
}

// HasClearMarks returns true if the method "Performance.clearMarks" exists.
func (this Performance) HasClearMarks() bool {
	return js.True == bindings.HasPerformanceClearMarks(
		this.Ref(),
	)
}

// ClearMarksFunc returns the method "Performance.clearMarks".
func (this Performance) ClearMarksFunc() (fn js.Func[func(markName js.String)]) {
	return fn.FromRef(
		bindings.PerformanceClearMarksFunc(
			this.Ref(),
		),
	)
}

// ClearMarks calls the method "Performance.clearMarks".
func (this Performance) ClearMarks(markName js.String) (ret js.Void) {
	bindings.CallPerformanceClearMarks(
		this.Ref(), js.Pointer(&ret),
		markName.Ref(),
	)

	return
}

// TryClearMarks calls the method "Performance.clearMarks"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryClearMarks(markName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMarks(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
	)

	return
}

// HasClearMarks1 returns true if the method "Performance.clearMarks" exists.
func (this Performance) HasClearMarks1() bool {
	return js.True == bindings.HasPerformanceClearMarks1(
		this.Ref(),
	)
}

// ClearMarks1Func returns the method "Performance.clearMarks".
func (this Performance) ClearMarks1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearMarks1Func(
			this.Ref(),
		),
	)
}

// ClearMarks1 calls the method "Performance.clearMarks".
func (this Performance) ClearMarks1() (ret js.Void) {
	bindings.CallPerformanceClearMarks1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearMarks1 calls the method "Performance.clearMarks"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryClearMarks1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMarks1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMeasure returns true if the method "Performance.measure" exists.
func (this Performance) HasMeasure() bool {
	return js.True == bindings.HasPerformanceMeasure(
		this.Ref(),
	)
}

// MeasureFunc returns the method "Performance.measure".
func (this Performance) MeasureFunc() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasureFunc(
			this.Ref(),
		),
	)
}

// Measure calls the method "Performance.measure".
func (this Performance) Measure(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure(
		this.Ref(), js.Pointer(&ret),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
		endMark.Ref(),
	)

	return
}

// TryMeasure calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMeasure(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
		endMark.Ref(),
	)

	return
}

// HasMeasure1 returns true if the method "Performance.measure" exists.
func (this Performance) HasMeasure1() bool {
	return js.True == bindings.HasPerformanceMeasure1(
		this.Ref(),
	)
}

// Measure1Func returns the method "Performance.measure".
func (this Performance) Measure1Func() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasure1Func(
			this.Ref(),
		),
	)
}

// Measure1 calls the method "Performance.measure".
func (this Performance) Measure1(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure1(
		this.Ref(), js.Pointer(&ret),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
	)

	return
}

// TryMeasure1 calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMeasure1(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
	)

	return
}

// HasMeasure2 returns true if the method "Performance.measure" exists.
func (this Performance) HasMeasure2() bool {
	return js.True == bindings.HasPerformanceMeasure2(
		this.Ref(),
	)
}

// Measure2Func returns the method "Performance.measure".
func (this Performance) Measure2Func() (fn js.Func[func(measureName js.String) PerformanceMeasure]) {
	return fn.FromRef(
		bindings.PerformanceMeasure2Func(
			this.Ref(),
		),
	)
}

// Measure2 calls the method "Performance.measure".
func (this Performance) Measure2(measureName js.String) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure2(
		this.Ref(), js.Pointer(&ret),
		measureName.Ref(),
	)

	return
}

// TryMeasure2 calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryMeasure2(measureName js.String) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
	)

	return
}

// HasClearMeasures returns true if the method "Performance.clearMeasures" exists.
func (this Performance) HasClearMeasures() bool {
	return js.True == bindings.HasPerformanceClearMeasures(
		this.Ref(),
	)
}

// ClearMeasuresFunc returns the method "Performance.clearMeasures".
func (this Performance) ClearMeasuresFunc() (fn js.Func[func(measureName js.String)]) {
	return fn.FromRef(
		bindings.PerformanceClearMeasuresFunc(
			this.Ref(),
		),
	)
}

// ClearMeasures calls the method "Performance.clearMeasures".
func (this Performance) ClearMeasures(measureName js.String) (ret js.Void) {
	bindings.CallPerformanceClearMeasures(
		this.Ref(), js.Pointer(&ret),
		measureName.Ref(),
	)

	return
}

// TryClearMeasures calls the method "Performance.clearMeasures"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryClearMeasures(measureName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMeasures(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
	)

	return
}

// HasClearMeasures1 returns true if the method "Performance.clearMeasures" exists.
func (this Performance) HasClearMeasures1() bool {
	return js.True == bindings.HasPerformanceClearMeasures1(
		this.Ref(),
	)
}

// ClearMeasures1Func returns the method "Performance.clearMeasures".
func (this Performance) ClearMeasures1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearMeasures1Func(
			this.Ref(),
		),
	)
}

// ClearMeasures1 calls the method "Performance.clearMeasures".
func (this Performance) ClearMeasures1() (ret js.Void) {
	bindings.CallPerformanceClearMeasures1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearMeasures1 calls the method "Performance.clearMeasures"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryClearMeasures1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMeasures1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClearResourceTimings returns true if the method "Performance.clearResourceTimings" exists.
func (this Performance) HasClearResourceTimings() bool {
	return js.True == bindings.HasPerformanceClearResourceTimings(
		this.Ref(),
	)
}

// ClearResourceTimingsFunc returns the method "Performance.clearResourceTimings".
func (this Performance) ClearResourceTimingsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceClearResourceTimingsFunc(
			this.Ref(),
		),
	)
}

// ClearResourceTimings calls the method "Performance.clearResourceTimings".
func (this Performance) ClearResourceTimings() (ret js.Void) {
	bindings.CallPerformanceClearResourceTimings(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearResourceTimings calls the method "Performance.clearResourceTimings"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryClearResourceTimings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearResourceTimings(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetResourceTimingBufferSize returns true if the method "Performance.setResourceTimingBufferSize" exists.
func (this Performance) HasSetResourceTimingBufferSize() bool {
	return js.True == bindings.HasPerformanceSetResourceTimingBufferSize(
		this.Ref(),
	)
}

// SetResourceTimingBufferSizeFunc returns the method "Performance.setResourceTimingBufferSize".
func (this Performance) SetResourceTimingBufferSizeFunc() (fn js.Func[func(maxSize uint32)]) {
	return fn.FromRef(
		bindings.PerformanceSetResourceTimingBufferSizeFunc(
			this.Ref(),
		),
	)
}

// SetResourceTimingBufferSize calls the method "Performance.setResourceTimingBufferSize".
func (this Performance) SetResourceTimingBufferSize(maxSize uint32) (ret js.Void) {
	bindings.CallPerformanceSetResourceTimingBufferSize(
		this.Ref(), js.Pointer(&ret),
		uint32(maxSize),
	)

	return
}

// TrySetResourceTimingBufferSize calls the method "Performance.setResourceTimingBufferSize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TrySetResourceTimingBufferSize(maxSize uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceSetResourceTimingBufferSize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(maxSize),
	)

	return
}

// HasGetEntries returns true if the method "Performance.getEntries" exists.
func (this Performance) HasGetEntries() bool {
	return js.True == bindings.HasPerformanceGetEntries(
		this.Ref(),
	)
}

// GetEntriesFunc returns the method "Performance.getEntries".
func (this Performance) GetEntriesFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesFunc(
			this.Ref(),
		),
	)
}

// GetEntries calls the method "Performance.getEntries".
func (this Performance) GetEntries() (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntries(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetEntries calls the method "Performance.getEntries"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryGetEntries() (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntries(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetEntriesByType returns true if the method "Performance.getEntriesByType" exists.
func (this Performance) HasGetEntriesByType() bool {
	return js.True == bindings.HasPerformanceGetEntriesByType(
		this.Ref(),
	)
}

// GetEntriesByTypeFunc returns the method "Performance.getEntriesByType".
func (this Performance) GetEntriesByTypeFunc() (fn js.Func[func(typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByTypeFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByType calls the method "Performance.getEntriesByType".
func (this Performance) GetEntriesByType(typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByType(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByType calls the method "Performance.getEntriesByType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryGetEntriesByType(typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasGetEntriesByName returns true if the method "Performance.getEntriesByName" exists.
func (this Performance) HasGetEntriesByName() bool {
	return js.True == bindings.HasPerformanceGetEntriesByName(
		this.Ref(),
	)
}

// GetEntriesByNameFunc returns the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByNameFunc() (fn js.Func[func(name js.String, typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByNameFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName calls the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByName(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByName calls the method "Performance.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryGetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// HasGetEntriesByName1 returns true if the method "Performance.getEntriesByName" exists.
func (this Performance) HasGetEntriesByName1() bool {
	return js.True == bindings.HasPerformanceGetEntriesByName1(
		this.Ref(),
	)
}

// GetEntriesByName1Func returns the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByName1Func() (fn js.Func[func(name js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceGetEntriesByName1Func(
			this.Ref(),
		),
	)
}

// GetEntriesByName1 calls the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByName1(name js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByName1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetEntriesByName1 calls the method "Performance.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Performance) TryGetEntriesByName1(name js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByName1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Storage) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStorageLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasKey returns true if the method "Storage.key" exists.
func (this Storage) HasKey() bool {
	return js.True == bindings.HasStorageKey(
		this.Ref(),
	)
}

// KeyFunc returns the method "Storage.key".
func (this Storage) KeyFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.StorageKeyFunc(
			this.Ref(),
		),
	)
}

// Key calls the method "Storage.key".
func (this Storage) Key(index uint32) (ret js.String) {
	bindings.CallStorageKey(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryKey calls the method "Storage.key"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Storage) TryKey(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasGetItem returns true if the method "Storage.getItem" exists.
func (this Storage) HasGetItem() bool {
	return js.True == bindings.HasStorageGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "Storage.getItem".
func (this Storage) GetItemFunc() (fn js.Func[func(key js.String) js.String]) {
	return fn.FromRef(
		bindings.StorageGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "Storage.getItem".
func (this Storage) GetItem(key js.String) (ret js.String) {
	bindings.CallStorageGetItem(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryGetItem calls the method "Storage.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Storage) TryGetItem(key js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasSetItem returns true if the method "Storage.setItem" exists.
func (this Storage) HasSetItem() bool {
	return js.True == bindings.HasStorageSetItem(
		this.Ref(),
	)
}

// SetItemFunc returns the method "Storage.setItem".
func (this Storage) SetItemFunc() (fn js.Func[func(key js.String, value js.String)]) {
	return fn.FromRef(
		bindings.StorageSetItemFunc(
			this.Ref(),
		),
	)
}

// SetItem calls the method "Storage.setItem".
func (this Storage) SetItem(key js.String, value js.String) (ret js.Void) {
	bindings.CallStorageSetItem(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
		value.Ref(),
	)

	return
}

// TrySetItem calls the method "Storage.setItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Storage) TrySetItem(key js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageSetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		value.Ref(),
	)

	return
}

// HasRemoveItem returns true if the method "Storage.removeItem" exists.
func (this Storage) HasRemoveItem() bool {
	return js.True == bindings.HasStorageRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "Storage.removeItem".
func (this Storage) RemoveItemFunc() (fn js.Func[func(key js.String)]) {
	return fn.FromRef(
		bindings.StorageRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "Storage.removeItem".
func (this Storage) RemoveItem(key js.String) (ret js.Void) {
	bindings.CallStorageRemoveItem(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryRemoveItem calls the method "Storage.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Storage) TryRemoveItem(key js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasClear returns true if the method "Storage.clear" exists.
func (this Storage) HasClear() bool {
	return js.True == bindings.HasStorageClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "Storage.clear".
func (this Storage) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.StorageClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "Storage.clear".
func (this Storage) Clear() (ret js.Void) {
	bindings.CallStorageClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "Storage.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Storage) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Window) Window() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Self returns the value of property "Window.self".
//
// It returns ok=false if there is no such property.
func (this Window) Self() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowSelf(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Document returns the value of property "Window.document".
//
// It returns ok=false if there is no such property.
func (this Window) Document() (ret Document, ok bool) {
	ok = js.True == bindings.GetWindowDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "Window.name".
//
// It returns ok=false if there is no such property.
func (this Window) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "Window.name" to val.
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
// It returns ok=false if there is no such property.
func (this Window) Location() (ret Location, ok bool) {
	ok = js.True == bindings.GetWindowLocation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// History returns the value of property "Window.history".
//
// It returns ok=false if there is no such property.
func (this Window) History() (ret History, ok bool) {
	ok = js.True == bindings.GetWindowHistory(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Navigation returns the value of property "Window.navigation".
//
// It returns ok=false if there is no such property.
func (this Window) Navigation() (ret Navigation, ok bool) {
	ok = js.True == bindings.GetWindowNavigation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CustomElements returns the value of property "Window.customElements".
//
// It returns ok=false if there is no such property.
func (this Window) CustomElements() (ret CustomElementRegistry, ok bool) {
	ok = js.True == bindings.GetWindowCustomElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Locationbar returns the value of property "Window.locationbar".
//
// It returns ok=false if there is no such property.
func (this Window) Locationbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowLocationbar(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Menubar returns the value of property "Window.menubar".
//
// It returns ok=false if there is no such property.
func (this Window) Menubar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowMenubar(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Personalbar returns the value of property "Window.personalbar".
//
// It returns ok=false if there is no such property.
func (this Window) Personalbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowPersonalbar(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scrollbars returns the value of property "Window.scrollbars".
//
// It returns ok=false if there is no such property.
func (this Window) Scrollbars() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowScrollbars(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Statusbar returns the value of property "Window.statusbar".
//
// It returns ok=false if there is no such property.
func (this Window) Statusbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowStatusbar(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Toolbar returns the value of property "Window.toolbar".
//
// It returns ok=false if there is no such property.
func (this Window) Toolbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowToolbar(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "Window.status".
//
// It returns ok=false if there is no such property.
func (this Window) Status() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStatus sets the value of property "Window.status" to val.
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
// It returns ok=false if there is no such property.
func (this Window) Closed() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowClosed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Frames returns the value of property "Window.frames".
//
// It returns ok=false if there is no such property.
func (this Window) Frames() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowFrames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "Window.length".
//
// It returns ok=false if there is no such property.
func (this Window) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWindowLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "Window.top".
//
// It returns ok=false if there is no such property.
func (this Window) Top() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Opener returns the value of property "Window.opener".
//
// It returns ok=false if there is no such property.
func (this Window) Opener() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetWindowOpener(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOpener sets the value of property "Window.opener" to val.
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
// It returns ok=false if there is no such property.
func (this Window) Parent() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowParent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FrameElement returns the value of property "Window.frameElement".
//
// It returns ok=false if there is no such property.
func (this Window) FrameElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetWindowFrameElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Navigator returns the value of property "Window.navigator".
//
// It returns ok=false if there is no such property.
func (this Window) Navigator() (ret Navigator, ok bool) {
	ok = js.True == bindings.GetWindowNavigator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientInformation returns the value of property "Window.clientInformation".
//
// It returns ok=false if there is no such property.
func (this Window) ClientInformation() (ret Navigator, ok bool) {
	ok = js.True == bindings.GetWindowClientInformation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OriginAgentCluster returns the value of property "Window.originAgentCluster".
//
// It returns ok=false if there is no such property.
func (this Window) OriginAgentCluster() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowOriginAgentCluster(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "Window.orientation".
//
// It returns ok=false if there is no such property.
func (this Window) Orientation() (ret int16, ok bool) {
	ok = js.True == bindings.GetWindowOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LaunchQueue returns the value of property "Window.launchQueue".
//
// It returns ok=false if there is no such property.
func (this Window) LaunchQueue() (ret LaunchQueue, ok bool) {
	ok = js.True == bindings.GetWindowLaunchQueue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fence returns the value of property "Window.fence".
//
// It returns ok=false if there is no such property.
func (this Window) Fence() (ret Fence, ok bool) {
	ok = js.True == bindings.GetWindowFence(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PortalHost returns the value of property "Window.portalHost".
//
// It returns ok=false if there is no such property.
func (this Window) PortalHost() (ret PortalHost, ok bool) {
	ok = js.True == bindings.GetWindowPortalHost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Screen returns the value of property "Window.screen".
//
// It returns ok=false if there is no such property.
func (this Window) Screen() (ret Screen, ok bool) {
	ok = js.True == bindings.GetWindowScreen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VisualViewport returns the value of property "Window.visualViewport".
//
// It returns ok=false if there is no such property.
func (this Window) VisualViewport() (ret VisualViewport, ok bool) {
	ok = js.True == bindings.GetWindowVisualViewport(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InnerWidth returns the value of property "Window.innerWidth".
//
// It returns ok=false if there is no such property.
func (this Window) InnerWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowInnerWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InnerHeight returns the value of property "Window.innerHeight".
//
// It returns ok=false if there is no such property.
func (this Window) InnerHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowInnerHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScrollX returns the value of property "Window.scrollX".
//
// It returns ok=false if there is no such property.
func (this Window) ScrollX() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowScrollX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageXOffset returns the value of property "Window.pageXOffset".
//
// It returns ok=false if there is no such property.
func (this Window) PageXOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowPageXOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScrollY returns the value of property "Window.scrollY".
//
// It returns ok=false if there is no such property.
func (this Window) ScrollY() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowScrollY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageYOffset returns the value of property "Window.pageYOffset".
//
// It returns ok=false if there is no such property.
func (this Window) PageYOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowPageYOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenX returns the value of property "Window.screenX".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenX() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenLeft returns the value of property "Window.screenLeft".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenY returns the value of property "Window.screenY".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenY() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenTop returns the value of property "Window.screenTop".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OuterWidth returns the value of property "Window.outerWidth".
//
// It returns ok=false if there is no such property.
func (this Window) OuterWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowOuterWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OuterHeight returns the value of property "Window.outerHeight".
//
// It returns ok=false if there is no such property.
func (this Window) OuterHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowOuterHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DevicePixelRatio returns the value of property "Window.devicePixelRatio".
//
// It returns ok=false if there is no such property.
func (this Window) DevicePixelRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowDevicePixelRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SharedStorage returns the value of property "Window.sharedStorage".
//
// It returns ok=false if there is no such property.
func (this Window) SharedStorage() (ret WindowSharedStorage, ok bool) {
	ok = js.True == bindings.GetWindowSharedStorage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Event returns the value of property "Window.event".
//
// It returns ok=false if there is no such property.
func (this Window) Event() (ret OneOf_Event_undefined, ok bool) {
	ok = js.True == bindings.GetWindowEvent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CookieStore returns the value of property "Window.cookieStore".
//
// It returns ok=false if there is no such property.
func (this Window) CookieStore() (ret CookieStore, ok bool) {
	ok = js.True == bindings.GetWindowCookieStore(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DocumentPictureInPicture returns the value of property "Window.documentPictureInPicture".
//
// It returns ok=false if there is no such property.
func (this Window) DocumentPictureInPicture() (ret DocumentPictureInPicture, ok bool) {
	ok = js.True == bindings.GetWindowDocumentPictureInPicture(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// External returns the value of property "Window.external".
//
// It returns ok=false if there is no such property.
func (this Window) External() (ret External, ok bool) {
	ok = js.True == bindings.GetWindowExternal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SpeechSynthesis returns the value of property "Window.speechSynthesis".
//
// It returns ok=false if there is no such property.
func (this Window) SpeechSynthesis() (ret SpeechSynthesis, ok bool) {
	ok = js.True == bindings.GetWindowSpeechSynthesis(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "Window.origin".
//
// It returns ok=false if there is no such property.
func (this Window) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsSecureContext returns the value of property "Window.isSecureContext".
//
// It returns ok=false if there is no such property.
func (this Window) IsSecureContext() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowIsSecureContext(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CrossOriginIsolated returns the value of property "Window.crossOriginIsolated".
//
// It returns ok=false if there is no such property.
func (this Window) CrossOriginIsolated() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowCrossOriginIsolated(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scheduler returns the value of property "Window.scheduler".
//
// It returns ok=false if there is no such property.
func (this Window) Scheduler() (ret Scheduler, ok bool) {
	ok = js.True == bindings.GetWindowScheduler(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TrustedTypes returns the value of property "Window.trustedTypes".
//
// It returns ok=false if there is no such property.
func (this Window) TrustedTypes() (ret TrustedTypePolicyFactory, ok bool) {
	ok = js.True == bindings.GetWindowTrustedTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "Window.caches".
//
// It returns ok=false if there is no such property.
func (this Window) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetWindowCaches(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Crypto returns the value of property "Window.crypto".
//
// It returns ok=false if there is no such property.
func (this Window) Crypto() (ret Crypto, ok bool) {
	ok = js.True == bindings.GetWindowCrypto(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "Window.indexedDB".
//
// It returns ok=false if there is no such property.
func (this Window) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetWindowIndexedDB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Performance returns the value of property "Window.performance".
//
// It returns ok=false if there is no such property.
func (this Window) Performance() (ret Performance, ok bool) {
	ok = js.True == bindings.GetWindowPerformance(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SessionStorage returns the value of property "Window.sessionStorage".
//
// It returns ok=false if there is no such property.
func (this Window) SessionStorage() (ret Storage, ok bool) {
	ok = js.True == bindings.GetWindowSessionStorage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LocalStorage returns the value of property "Window.localStorage".
//
// It returns ok=false if there is no such property.
func (this Window) LocalStorage() (ret Storage, ok bool) {
	ok = js.True == bindings.GetWindowLocalStorage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClose returns true if the method "Window.close" exists.
func (this Window) HasClose() bool {
	return js.True == bindings.HasWindowClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "Window.close".
func (this Window) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "Window.close".
func (this Window) Close() (ret js.Void) {
	bindings.CallWindowClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "Window.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "Window.stop" exists.
func (this Window) HasStop() bool {
	return js.True == bindings.HasWindowStop(
		this.Ref(),
	)
}

// StopFunc returns the method "Window.stop".
func (this Window) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "Window.stop".
func (this Window) Stop() (ret js.Void) {
	bindings.CallWindowStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "Window.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFocus returns true if the method "Window.focus" exists.
func (this Window) HasFocus() bool {
	return js.True == bindings.HasWindowFocus(
		this.Ref(),
	)
}

// FocusFunc returns the method "Window.focus".
func (this Window) FocusFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowFocusFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "Window.focus".
func (this Window) Focus() (ret js.Void) {
	bindings.CallWindowFocus(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "Window.focus"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlur returns true if the method "Window.blur" exists.
func (this Window) HasBlur() bool {
	return js.True == bindings.HasWindowBlur(
		this.Ref(),
	)
}

// BlurFunc returns the method "Window.blur".
func (this Window) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowBlurFunc(
			this.Ref(),
		),
	)
}

// Blur calls the method "Window.blur".
func (this Window) Blur() (ret js.Void) {
	bindings.CallWindowBlur(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "Window.blur"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowBlur(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpen returns true if the method "Window.open" exists.
func (this Window) HasOpen() bool {
	return js.True == bindings.HasWindowOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "Window.open".
func (this Window) OpenFunc() (fn js.Func[func(url js.String, target js.String, features js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "Window.open".
func (this Window) Open(url js.String, target js.String, features js.String) (ret js.Object) {
	bindings.CallWindowOpen(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
		target.Ref(),
		features.Ref(),
	)

	return
}

// TryOpen calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryOpen(url js.String, target js.String, features js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		target.Ref(),
		features.Ref(),
	)

	return
}

// HasOpen1 returns true if the method "Window.open" exists.
func (this Window) HasOpen1() bool {
	return js.True == bindings.HasWindowOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "Window.open".
func (this Window) Open1Func() (fn js.Func[func(url js.String, target js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "Window.open".
func (this Window) Open1(url js.String, target js.String) (ret js.Object) {
	bindings.CallWindowOpen1(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
		target.Ref(),
	)

	return
}

// TryOpen1 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryOpen1(url js.String, target js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		target.Ref(),
	)

	return
}

// HasOpen2 returns true if the method "Window.open" exists.
func (this Window) HasOpen2() bool {
	return js.True == bindings.HasWindowOpen2(
		this.Ref(),
	)
}

// Open2Func returns the method "Window.open".
func (this Window) Open2Func() (fn js.Func[func(url js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen2Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "Window.open".
func (this Window) Open2(url js.String) (ret js.Object) {
	bindings.CallWindowOpen2(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpen2 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryOpen2(url js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasOpen3 returns true if the method "Window.open" exists.
func (this Window) HasOpen3() bool {
	return js.True == bindings.HasWindowOpen3(
		this.Ref(),
	)
}

// Open3Func returns the method "Window.open".
func (this Window) Open3Func() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.WindowOpen3Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "Window.open".
func (this Window) Open3() (ret js.Object) {
	bindings.CallWindowOpen3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpen3 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryOpen3() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGet returns true if the method "Window." exists.
func (this Window) HasGet() bool {
	return js.True == bindings.HasWindowGet(
		this.Ref(),
	)
}

// GetFunc returns the method "Window.".
func (this Window) GetFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WindowGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Window.".
func (this Window) Get(name js.String) (ret js.Object) {
	bindings.CallWindowGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "Window."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGet(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasAlert returns true if the method "Window.alert" exists.
func (this Window) HasAlert() bool {
	return js.True == bindings.HasWindowAlert(
		this.Ref(),
	)
}

// AlertFunc returns the method "Window.alert".
func (this Window) AlertFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowAlertFunc(
			this.Ref(),
		),
	)
}

// Alert calls the method "Window.alert".
func (this Window) Alert() (ret js.Void) {
	bindings.CallWindowAlert(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAlert calls the method "Window.alert"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryAlert() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAlert(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAlert1 returns true if the method "Window.alert" exists.
func (this Window) HasAlert1() bool {
	return js.True == bindings.HasWindowAlert1(
		this.Ref(),
	)
}

// Alert1Func returns the method "Window.alert".
func (this Window) Alert1Func() (fn js.Func[func(message js.String)]) {
	return fn.FromRef(
		bindings.WindowAlert1Func(
			this.Ref(),
		),
	)
}

// Alert1 calls the method "Window.alert".
func (this Window) Alert1(message js.String) (ret js.Void) {
	bindings.CallWindowAlert1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryAlert1 calls the method "Window.alert"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryAlert1(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAlert1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasConfirm returns true if the method "Window.confirm" exists.
func (this Window) HasConfirm() bool {
	return js.True == bindings.HasWindowConfirm(
		this.Ref(),
	)
}

// ConfirmFunc returns the method "Window.confirm".
func (this Window) ConfirmFunc() (fn js.Func[func(message js.String) bool]) {
	return fn.FromRef(
		bindings.WindowConfirmFunc(
			this.Ref(),
		),
	)
}

// Confirm calls the method "Window.confirm".
func (this Window) Confirm(message js.String) (ret bool) {
	bindings.CallWindowConfirm(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryConfirm calls the method "Window.confirm"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryConfirm(message js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowConfirm(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasConfirm1 returns true if the method "Window.confirm" exists.
func (this Window) HasConfirm1() bool {
	return js.True == bindings.HasWindowConfirm1(
		this.Ref(),
	)
}

// Confirm1Func returns the method "Window.confirm".
func (this Window) Confirm1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WindowConfirm1Func(
			this.Ref(),
		),
	)
}

// Confirm1 calls the method "Window.confirm".
func (this Window) Confirm1() (ret bool) {
	bindings.CallWindowConfirm1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryConfirm1 calls the method "Window.confirm"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryConfirm1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowConfirm1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPrompt returns true if the method "Window.prompt" exists.
func (this Window) HasPrompt() bool {
	return js.True == bindings.HasWindowPrompt(
		this.Ref(),
	)
}

// PromptFunc returns the method "Window.prompt".
func (this Window) PromptFunc() (fn js.Func[func(message js.String, def js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowPromptFunc(
			this.Ref(),
		),
	)
}

// Prompt calls the method "Window.prompt".
func (this Window) Prompt(message js.String, def js.String) (ret js.String) {
	bindings.CallWindowPrompt(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		def.Ref(),
	)

	return
}

// TryPrompt calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPrompt(message js.String, def js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		def.Ref(),
	)

	return
}

// HasPrompt1 returns true if the method "Window.prompt" exists.
func (this Window) HasPrompt1() bool {
	return js.True == bindings.HasWindowPrompt1(
		this.Ref(),
	)
}

// Prompt1Func returns the method "Window.prompt".
func (this Window) Prompt1Func() (fn js.Func[func(message js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowPrompt1Func(
			this.Ref(),
		),
	)
}

// Prompt1 calls the method "Window.prompt".
func (this Window) Prompt1(message js.String) (ret js.String) {
	bindings.CallWindowPrompt1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPrompt1 calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPrompt1(message js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasPrompt2 returns true if the method "Window.prompt" exists.
func (this Window) HasPrompt2() bool {
	return js.True == bindings.HasWindowPrompt2(
		this.Ref(),
	)
}

// Prompt2Func returns the method "Window.prompt".
func (this Window) Prompt2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.WindowPrompt2Func(
			this.Ref(),
		),
	)
}

// Prompt2 calls the method "Window.prompt".
func (this Window) Prompt2() (ret js.String) {
	bindings.CallWindowPrompt2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPrompt2 calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPrompt2() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPrint returns true if the method "Window.print" exists.
func (this Window) HasPrint() bool {
	return js.True == bindings.HasWindowPrint(
		this.Ref(),
	)
}

// PrintFunc returns the method "Window.print".
func (this Window) PrintFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowPrintFunc(
			this.Ref(),
		),
	)
}

// Print calls the method "Window.print".
func (this Window) Print() (ret js.Void) {
	bindings.CallWindowPrint(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPrint calls the method "Window.print"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPrint() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPostMessage returns true if the method "Window.postMessage" exists.
func (this Window) HasPostMessage() bool {
	return js.True == bindings.HasWindowPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "Window.postMessage".
func (this Window) PostMessageFunc() (fn js.Func[func(message js.Any, targetOrigin js.String, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.WindowPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "Window.postMessage".
func (this Window) PostMessage(message js.Any, targetOrigin js.String, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallWindowPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		targetOrigin.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPostMessage(message js.Any, targetOrigin js.String, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		targetOrigin.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "Window.postMessage" exists.
func (this Window) HasPostMessage1() bool {
	return js.True == bindings.HasWindowPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "Window.postMessage".
func (this Window) PostMessage1Func() (fn js.Func[func(message js.Any, targetOrigin js.String)]) {
	return fn.FromRef(
		bindings.WindowPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Window.postMessage".
func (this Window) PostMessage1(message js.Any, targetOrigin js.String) (ret js.Void) {
	bindings.CallWindowPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPostMessage1(message js.Any, targetOrigin js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

// HasPostMessage2 returns true if the method "Window.postMessage" exists.
func (this Window) HasPostMessage2() bool {
	return js.True == bindings.HasWindowPostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "Window.postMessage".
func (this Window) PostMessage2Func() (fn js.Func[func(message js.Any, options WindowPostMessageOptions)]) {
	return fn.FromRef(
		bindings.WindowPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Window.postMessage".
func (this Window) PostMessage2(message js.Any, options WindowPostMessageOptions) (ret js.Void) {
	bindings.CallWindowPostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage2 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPostMessage2(message js.Any, options WindowPostMessageOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage3 returns true if the method "Window.postMessage" exists.
func (this Window) HasPostMessage3() bool {
	return js.True == bindings.HasWindowPostMessage3(
		this.Ref(),
	)
}

// PostMessage3Func returns the method "Window.postMessage".
func (this Window) PostMessage3Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.WindowPostMessage3Func(
			this.Ref(),
		),
	)
}

// PostMessage3 calls the method "Window.postMessage".
func (this Window) PostMessage3(message js.Any) (ret js.Void) {
	bindings.CallWindowPostMessage3(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage3 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryPostMessage3(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasGetSelection returns true if the method "Window.getSelection" exists.
func (this Window) HasGetSelection() bool {
	return js.True == bindings.HasWindowGetSelection(
		this.Ref(),
	)
}

// GetSelectionFunc returns the method "Window.getSelection".
func (this Window) GetSelectionFunc() (fn js.Func[func() Selection]) {
	return fn.FromRef(
		bindings.WindowGetSelectionFunc(
			this.Ref(),
		),
	)
}

// GetSelection calls the method "Window.getSelection".
func (this Window) GetSelection() (ret Selection) {
	bindings.CallWindowGetSelection(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSelection calls the method "Window.getSelection"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGetSelection() (ret Selection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetSelection(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasNavigate returns true if the method "Window.navigate" exists.
func (this Window) HasNavigate() bool {
	return js.True == bindings.HasWindowNavigate(
		this.Ref(),
	)
}

// NavigateFunc returns the method "Window.navigate".
func (this Window) NavigateFunc() (fn js.Func[func(dir SpatialNavigationDirection)]) {
	return fn.FromRef(
		bindings.WindowNavigateFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "Window.navigate".
func (this Window) Navigate(dir SpatialNavigationDirection) (ret js.Void) {
	bindings.CallWindowNavigate(
		this.Ref(), js.Pointer(&ret),
		uint32(dir),
	)

	return
}

// TryNavigate calls the method "Window.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryNavigate(dir SpatialNavigationDirection) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowNavigate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
	)

	return
}

// HasGetComputedStyle returns true if the method "Window.getComputedStyle" exists.
func (this Window) HasGetComputedStyle() bool {
	return js.True == bindings.HasWindowGetComputedStyle(
		this.Ref(),
	)
}

// GetComputedStyleFunc returns the method "Window.getComputedStyle".
func (this Window) GetComputedStyleFunc() (fn js.Func[func(elt Element, pseudoElt js.String) CSSStyleDeclaration]) {
	return fn.FromRef(
		bindings.WindowGetComputedStyleFunc(
			this.Ref(),
		),
	)
}

// GetComputedStyle calls the method "Window.getComputedStyle".
func (this Window) GetComputedStyle(elt Element, pseudoElt js.String) (ret CSSStyleDeclaration) {
	bindings.CallWindowGetComputedStyle(
		this.Ref(), js.Pointer(&ret),
		elt.Ref(),
		pseudoElt.Ref(),
	)

	return
}

// TryGetComputedStyle calls the method "Window.getComputedStyle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGetComputedStyle(elt Element, pseudoElt js.String) (ret CSSStyleDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetComputedStyle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elt.Ref(),
		pseudoElt.Ref(),
	)

	return
}

// HasGetComputedStyle1 returns true if the method "Window.getComputedStyle" exists.
func (this Window) HasGetComputedStyle1() bool {
	return js.True == bindings.HasWindowGetComputedStyle1(
		this.Ref(),
	)
}

// GetComputedStyle1Func returns the method "Window.getComputedStyle".
func (this Window) GetComputedStyle1Func() (fn js.Func[func(elt Element) CSSStyleDeclaration]) {
	return fn.FromRef(
		bindings.WindowGetComputedStyle1Func(
			this.Ref(),
		),
	)
}

// GetComputedStyle1 calls the method "Window.getComputedStyle".
func (this Window) GetComputedStyle1(elt Element) (ret CSSStyleDeclaration) {
	bindings.CallWindowGetComputedStyle1(
		this.Ref(), js.Pointer(&ret),
		elt.Ref(),
	)

	return
}

// TryGetComputedStyle1 calls the method "Window.getComputedStyle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGetComputedStyle1(elt Element) (ret CSSStyleDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetComputedStyle1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elt.Ref(),
	)

	return
}

// HasMatchMedia returns true if the method "Window.matchMedia" exists.
func (this Window) HasMatchMedia() bool {
	return js.True == bindings.HasWindowMatchMedia(
		this.Ref(),
	)
}

// MatchMediaFunc returns the method "Window.matchMedia".
func (this Window) MatchMediaFunc() (fn js.Func[func(query js.String) MediaQueryList]) {
	return fn.FromRef(
		bindings.WindowMatchMediaFunc(
			this.Ref(),
		),
	)
}

// MatchMedia calls the method "Window.matchMedia".
func (this Window) MatchMedia(query js.String) (ret MediaQueryList) {
	bindings.CallWindowMatchMedia(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryMatchMedia calls the method "Window.matchMedia"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryMatchMedia(query js.String) (ret MediaQueryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMatchMedia(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasMoveTo returns true if the method "Window.moveTo" exists.
func (this Window) HasMoveTo() bool {
	return js.True == bindings.HasWindowMoveTo(
		this.Ref(),
	)
}

// MoveToFunc returns the method "Window.moveTo".
func (this Window) MoveToFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowMoveToFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "Window.moveTo".
func (this Window) MoveTo(x int32, y int32) (ret js.Void) {
	bindings.CallWindowMoveTo(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryMoveTo calls the method "Window.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryMoveTo(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMoveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasMoveBy returns true if the method "Window.moveBy" exists.
func (this Window) HasMoveBy() bool {
	return js.True == bindings.HasWindowMoveBy(
		this.Ref(),
	)
}

// MoveByFunc returns the method "Window.moveBy".
func (this Window) MoveByFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowMoveByFunc(
			this.Ref(),
		),
	)
}

// MoveBy calls the method "Window.moveBy".
func (this Window) MoveBy(x int32, y int32) (ret js.Void) {
	bindings.CallWindowMoveBy(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryMoveBy calls the method "Window.moveBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryMoveBy(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMoveBy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasResizeTo returns true if the method "Window.resizeTo" exists.
func (this Window) HasResizeTo() bool {
	return js.True == bindings.HasWindowResizeTo(
		this.Ref(),
	)
}

// ResizeToFunc returns the method "Window.resizeTo".
func (this Window) ResizeToFunc() (fn js.Func[func(width int32, height int32)]) {
	return fn.FromRef(
		bindings.WindowResizeToFunc(
			this.Ref(),
		),
	)
}

// ResizeTo calls the method "Window.resizeTo".
func (this Window) ResizeTo(width int32, height int32) (ret js.Void) {
	bindings.CallWindowResizeTo(
		this.Ref(), js.Pointer(&ret),
		int32(width),
		int32(height),
	)

	return
}

// TryResizeTo calls the method "Window.resizeTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryResizeTo(width int32, height int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowResizeTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(width),
		int32(height),
	)

	return
}

// HasResizeBy returns true if the method "Window.resizeBy" exists.
func (this Window) HasResizeBy() bool {
	return js.True == bindings.HasWindowResizeBy(
		this.Ref(),
	)
}

// ResizeByFunc returns the method "Window.resizeBy".
func (this Window) ResizeByFunc() (fn js.Func[func(x int32, y int32)]) {
	return fn.FromRef(
		bindings.WindowResizeByFunc(
			this.Ref(),
		),
	)
}

// ResizeBy calls the method "Window.resizeBy".
func (this Window) ResizeBy(x int32, y int32) (ret js.Void) {
	bindings.CallWindowResizeBy(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryResizeBy calls the method "Window.resizeBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryResizeBy(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowResizeBy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasScroll returns true if the method "Window.scroll" exists.
func (this Window) HasScroll() bool {
	return js.True == bindings.HasWindowScroll(
		this.Ref(),
	)
}

// ScrollFunc returns the method "Window.scroll".
func (this Window) ScrollFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollFunc(
			this.Ref(),
		),
	)
}

// Scroll calls the method "Window.scroll".
func (this Window) Scroll(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScroll(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScroll calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScroll(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScroll1 returns true if the method "Window.scroll" exists.
func (this Window) HasScroll1() bool {
	return js.True == bindings.HasWindowScroll1(
		this.Ref(),
	)
}

// Scroll1Func returns the method "Window.scroll".
func (this Window) Scroll1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScroll1Func(
			this.Ref(),
		),
	)
}

// Scroll1 calls the method "Window.scroll".
func (this Window) Scroll1() (ret js.Void) {
	bindings.CallWindowScroll1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScroll1 calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScroll1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScroll2 returns true if the method "Window.scroll" exists.
func (this Window) HasScroll2() bool {
	return js.True == bindings.HasWindowScroll2(
		this.Ref(),
	)
}

// Scroll2Func returns the method "Window.scroll".
func (this Window) Scroll2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScroll2Func(
			this.Ref(),
		),
	)
}

// Scroll2 calls the method "Window.scroll".
func (this Window) Scroll2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScroll2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScroll2 calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScroll2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasScrollTo returns true if the method "Window.scrollTo" exists.
func (this Window) HasScrollTo() bool {
	return js.True == bindings.HasWindowScrollTo(
		this.Ref(),
	)
}

// ScrollToFunc returns the method "Window.scrollTo".
func (this Window) ScrollToFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollToFunc(
			this.Ref(),
		),
	)
}

// ScrollTo calls the method "Window.scrollTo".
func (this Window) ScrollTo(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScrollTo(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollTo calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollTo(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScrollTo1 returns true if the method "Window.scrollTo" exists.
func (this Window) HasScrollTo1() bool {
	return js.True == bindings.HasWindowScrollTo1(
		this.Ref(),
	)
}

// ScrollTo1Func returns the method "Window.scrollTo".
func (this Window) ScrollTo1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScrollTo1Func(
			this.Ref(),
		),
	)
}

// ScrollTo1 calls the method "Window.scrollTo".
func (this Window) ScrollTo1() (ret js.Void) {
	bindings.CallWindowScrollTo1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollTo1 calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollTo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollTo2 returns true if the method "Window.scrollTo" exists.
func (this Window) HasScrollTo2() bool {
	return js.True == bindings.HasWindowScrollTo2(
		this.Ref(),
	)
}

// ScrollTo2Func returns the method "Window.scrollTo".
func (this Window) ScrollTo2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScrollTo2Func(
			this.Ref(),
		),
	)
}

// ScrollTo2 calls the method "Window.scrollTo".
func (this Window) ScrollTo2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScrollTo2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollTo2 calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollTo2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasScrollBy returns true if the method "Window.scrollBy" exists.
func (this Window) HasScrollBy() bool {
	return js.True == bindings.HasWindowScrollBy(
		this.Ref(),
	)
}

// ScrollByFunc returns the method "Window.scrollBy".
func (this Window) ScrollByFunc() (fn js.Func[func(options ScrollToOptions)]) {
	return fn.FromRef(
		bindings.WindowScrollByFunc(
			this.Ref(),
		),
	)
}

// ScrollBy calls the method "Window.scrollBy".
func (this Window) ScrollBy(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScrollBy(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollBy calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollBy(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScrollBy1 returns true if the method "Window.scrollBy" exists.
func (this Window) HasScrollBy1() bool {
	return js.True == bindings.HasWindowScrollBy1(
		this.Ref(),
	)
}

// ScrollBy1Func returns the method "Window.scrollBy".
func (this Window) ScrollBy1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowScrollBy1Func(
			this.Ref(),
		),
	)
}

// ScrollBy1 calls the method "Window.scrollBy".
func (this Window) ScrollBy1() (ret js.Void) {
	bindings.CallWindowScrollBy1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollBy1 calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollBy1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollBy2 returns true if the method "Window.scrollBy" exists.
func (this Window) HasScrollBy2() bool {
	return js.True == bindings.HasWindowScrollBy2(
		this.Ref(),
	)
}

// ScrollBy2Func returns the method "Window.scrollBy".
func (this Window) ScrollBy2Func() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.WindowScrollBy2Func(
			this.Ref(),
		),
	)
}

// ScrollBy2 calls the method "Window.scrollBy".
func (this Window) ScrollBy2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScrollBy2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollBy2 calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryScrollBy2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRequestIdleCallback returns true if the method "Window.requestIdleCallback" exists.
func (this Window) HasRequestIdleCallback() bool {
	return js.True == bindings.HasWindowRequestIdleCallback(
		this.Ref(),
	)
}

// RequestIdleCallbackFunc returns the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallbackFunc() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestIdleCallbackFunc(
			this.Ref(),
		),
	)
}

// RequestIdleCallback calls the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallback(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) (ret uint32) {
	bindings.CallWindowRequestIdleCallback(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRequestIdleCallback calls the method "Window.requestIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryRequestIdleCallback(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestIdleCallback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRequestIdleCallback1 returns true if the method "Window.requestIdleCallback" exists.
func (this Window) HasRequestIdleCallback1() bool {
	return js.True == bindings.HasWindowRequestIdleCallback1(
		this.Ref(),
	)
}

// RequestIdleCallback1Func returns the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallback1Func() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)]) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestIdleCallback1Func(
			this.Ref(),
		),
	)
}

// RequestIdleCallback1 calls the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallback1(callback js.Func[func(deadline IdleDeadline)]) (ret uint32) {
	bindings.CallWindowRequestIdleCallback1(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestIdleCallback1 calls the method "Window.requestIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryRequestIdleCallback1(callback js.Func[func(deadline IdleDeadline)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestIdleCallback1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCancelIdleCallback returns true if the method "Window.cancelIdleCallback" exists.
func (this Window) HasCancelIdleCallback() bool {
	return js.True == bindings.HasWindowCancelIdleCallback(
		this.Ref(),
	)
}

// CancelIdleCallbackFunc returns the method "Window.cancelIdleCallback".
func (this Window) CancelIdleCallbackFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.WindowCancelIdleCallbackFunc(
			this.Ref(),
		),
	)
}

// CancelIdleCallback calls the method "Window.cancelIdleCallback".
func (this Window) CancelIdleCallback(handle uint32) (ret js.Void) {
	bindings.CallWindowCancelIdleCallback(
		this.Ref(), js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelIdleCallback calls the method "Window.cancelIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCancelIdleCallback(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCancelIdleCallback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

// HasShowOpenFilePicker returns true if the method "Window.showOpenFilePicker" exists.
func (this Window) HasShowOpenFilePicker() bool {
	return js.True == bindings.HasWindowShowOpenFilePicker(
		this.Ref(),
	)
}

// ShowOpenFilePickerFunc returns the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePickerFunc() (fn js.Func[func(options OpenFilePickerOptions) js.Promise[js.Array[FileSystemFileHandle]]]) {
	return fn.FromRef(
		bindings.WindowShowOpenFilePickerFunc(
			this.Ref(),
		),
	)
}

// ShowOpenFilePicker calls the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePicker(options OpenFilePickerOptions) (ret js.Promise[js.Array[FileSystemFileHandle]]) {
	bindings.CallWindowShowOpenFilePicker(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowOpenFilePicker calls the method "Window.showOpenFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowOpenFilePicker(options OpenFilePickerOptions) (ret js.Promise[js.Array[FileSystemFileHandle]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowOpenFilePicker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasShowOpenFilePicker1 returns true if the method "Window.showOpenFilePicker" exists.
func (this Window) HasShowOpenFilePicker1() bool {
	return js.True == bindings.HasWindowShowOpenFilePicker1(
		this.Ref(),
	)
}

// ShowOpenFilePicker1Func returns the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePicker1Func() (fn js.Func[func() js.Promise[js.Array[FileSystemFileHandle]]]) {
	return fn.FromRef(
		bindings.WindowShowOpenFilePicker1Func(
			this.Ref(),
		),
	)
}

// ShowOpenFilePicker1 calls the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePicker1() (ret js.Promise[js.Array[FileSystemFileHandle]]) {
	bindings.CallWindowShowOpenFilePicker1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowOpenFilePicker1 calls the method "Window.showOpenFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowOpenFilePicker1() (ret js.Promise[js.Array[FileSystemFileHandle]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowOpenFilePicker1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasShowSaveFilePicker returns true if the method "Window.showSaveFilePicker" exists.
func (this Window) HasShowSaveFilePicker() bool {
	return js.True == bindings.HasWindowShowSaveFilePicker(
		this.Ref(),
	)
}

// ShowSaveFilePickerFunc returns the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePickerFunc() (fn js.Func[func(options SaveFilePickerOptions) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.WindowShowSaveFilePickerFunc(
			this.Ref(),
		),
	)
}

// ShowSaveFilePicker calls the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePicker(options SaveFilePickerOptions) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallWindowShowSaveFilePicker(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowSaveFilePicker calls the method "Window.showSaveFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowSaveFilePicker(options SaveFilePickerOptions) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowSaveFilePicker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasShowSaveFilePicker1 returns true if the method "Window.showSaveFilePicker" exists.
func (this Window) HasShowSaveFilePicker1() bool {
	return js.True == bindings.HasWindowShowSaveFilePicker1(
		this.Ref(),
	)
}

// ShowSaveFilePicker1Func returns the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePicker1Func() (fn js.Func[func() js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.WindowShowSaveFilePicker1Func(
			this.Ref(),
		),
	)
}

// ShowSaveFilePicker1 calls the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePicker1() (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallWindowShowSaveFilePicker1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowSaveFilePicker1 calls the method "Window.showSaveFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowSaveFilePicker1() (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowSaveFilePicker1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasShowDirectoryPicker returns true if the method "Window.showDirectoryPicker" exists.
func (this Window) HasShowDirectoryPicker() bool {
	return js.True == bindings.HasWindowShowDirectoryPicker(
		this.Ref(),
	)
}

// ShowDirectoryPickerFunc returns the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPickerFunc() (fn js.Func[func(options DirectoryPickerOptions) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.WindowShowDirectoryPickerFunc(
			this.Ref(),
		),
	)
}

// ShowDirectoryPicker calls the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPicker(options DirectoryPickerOptions) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallWindowShowDirectoryPicker(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowDirectoryPicker calls the method "Window.showDirectoryPicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowDirectoryPicker(options DirectoryPickerOptions) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowDirectoryPicker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasShowDirectoryPicker1 returns true if the method "Window.showDirectoryPicker" exists.
func (this Window) HasShowDirectoryPicker1() bool {
	return js.True == bindings.HasWindowShowDirectoryPicker1(
		this.Ref(),
	)
}

// ShowDirectoryPicker1Func returns the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPicker1Func() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.WindowShowDirectoryPicker1Func(
			this.Ref(),
		),
	)
}

// ShowDirectoryPicker1 calls the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPicker1() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallWindowShowDirectoryPicker1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowDirectoryPicker1 calls the method "Window.showDirectoryPicker"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryShowDirectoryPicker1() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowDirectoryPicker1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasQueryLocalFonts returns true if the method "Window.queryLocalFonts" exists.
func (this Window) HasQueryLocalFonts() bool {
	return js.True == bindings.HasWindowQueryLocalFonts(
		this.Ref(),
	)
}

// QueryLocalFontsFunc returns the method "Window.queryLocalFonts".
func (this Window) QueryLocalFontsFunc() (fn js.Func[func(options QueryOptions) js.Promise[js.Array[FontData]]]) {
	return fn.FromRef(
		bindings.WindowQueryLocalFontsFunc(
			this.Ref(),
		),
	)
}

// QueryLocalFonts calls the method "Window.queryLocalFonts".
func (this Window) QueryLocalFonts(options QueryOptions) (ret js.Promise[js.Array[FontData]]) {
	bindings.CallWindowQueryLocalFonts(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryQueryLocalFonts calls the method "Window.queryLocalFonts"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryQueryLocalFonts(options QueryOptions) (ret js.Promise[js.Array[FontData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueryLocalFonts(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasQueryLocalFonts1 returns true if the method "Window.queryLocalFonts" exists.
func (this Window) HasQueryLocalFonts1() bool {
	return js.True == bindings.HasWindowQueryLocalFonts1(
		this.Ref(),
	)
}

// QueryLocalFonts1Func returns the method "Window.queryLocalFonts".
func (this Window) QueryLocalFonts1Func() (fn js.Func[func() js.Promise[js.Array[FontData]]]) {
	return fn.FromRef(
		bindings.WindowQueryLocalFonts1Func(
			this.Ref(),
		),
	)
}

// QueryLocalFonts1 calls the method "Window.queryLocalFonts".
func (this Window) QueryLocalFonts1() (ret js.Promise[js.Array[FontData]]) {
	bindings.CallWindowQueryLocalFonts1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryQueryLocalFonts1 calls the method "Window.queryLocalFonts"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryQueryLocalFonts1() (ret js.Promise[js.Array[FontData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueryLocalFonts1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetScreenDetails returns true if the method "Window.getScreenDetails" exists.
func (this Window) HasGetScreenDetails() bool {
	return js.True == bindings.HasWindowGetScreenDetails(
		this.Ref(),
	)
}

// GetScreenDetailsFunc returns the method "Window.getScreenDetails".
func (this Window) GetScreenDetailsFunc() (fn js.Func[func() js.Promise[ScreenDetails]]) {
	return fn.FromRef(
		bindings.WindowGetScreenDetailsFunc(
			this.Ref(),
		),
	)
}

// GetScreenDetails calls the method "Window.getScreenDetails".
func (this Window) GetScreenDetails() (ret js.Promise[ScreenDetails]) {
	bindings.CallWindowGetScreenDetails(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetScreenDetails calls the method "Window.getScreenDetails"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGetScreenDetails() (ret js.Promise[ScreenDetails], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetScreenDetails(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDigitalGoodsService returns true if the method "Window.getDigitalGoodsService" exists.
func (this Window) HasGetDigitalGoodsService() bool {
	return js.True == bindings.HasWindowGetDigitalGoodsService(
		this.Ref(),
	)
}

// GetDigitalGoodsServiceFunc returns the method "Window.getDigitalGoodsService".
func (this Window) GetDigitalGoodsServiceFunc() (fn js.Func[func(serviceProvider js.String) js.Promise[DigitalGoodsService]]) {
	return fn.FromRef(
		bindings.WindowGetDigitalGoodsServiceFunc(
			this.Ref(),
		),
	)
}

// GetDigitalGoodsService calls the method "Window.getDigitalGoodsService".
func (this Window) GetDigitalGoodsService(serviceProvider js.String) (ret js.Promise[DigitalGoodsService]) {
	bindings.CallWindowGetDigitalGoodsService(
		this.Ref(), js.Pointer(&ret),
		serviceProvider.Ref(),
	)

	return
}

// TryGetDigitalGoodsService calls the method "Window.getDigitalGoodsService"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryGetDigitalGoodsService(serviceProvider js.String) (ret js.Promise[DigitalGoodsService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetDigitalGoodsService(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		serviceProvider.Ref(),
	)

	return
}

// HasCaptureEvents returns true if the method "Window.captureEvents" exists.
func (this Window) HasCaptureEvents() bool {
	return js.True == bindings.HasWindowCaptureEvents(
		this.Ref(),
	)
}

// CaptureEventsFunc returns the method "Window.captureEvents".
func (this Window) CaptureEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowCaptureEventsFunc(
			this.Ref(),
		),
	)
}

// CaptureEvents calls the method "Window.captureEvents".
func (this Window) CaptureEvents() (ret js.Void) {
	bindings.CallWindowCaptureEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCaptureEvents calls the method "Window.captureEvents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCaptureEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCaptureEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReleaseEvents returns true if the method "Window.releaseEvents" exists.
func (this Window) HasReleaseEvents() bool {
	return js.True == bindings.HasWindowReleaseEvents(
		this.Ref(),
	)
}

// ReleaseEventsFunc returns the method "Window.releaseEvents".
func (this Window) ReleaseEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowReleaseEventsFunc(
			this.Ref(),
		),
	)
}

// ReleaseEvents calls the method "Window.releaseEvents".
func (this Window) ReleaseEvents() (ret js.Void) {
	bindings.CallWindowReleaseEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReleaseEvents calls the method "Window.releaseEvents"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryReleaseEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowReleaseEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportError returns true if the method "Window.reportError" exists.
func (this Window) HasReportError() bool {
	return js.True == bindings.HasWindowReportError(
		this.Ref(),
	)
}

// ReportErrorFunc returns the method "Window.reportError".
func (this Window) ReportErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.WindowReportErrorFunc(
			this.Ref(),
		),
	)
}

// ReportError calls the method "Window.reportError".
func (this Window) ReportError(e js.Any) (ret js.Void) {
	bindings.CallWindowReportError(
		this.Ref(), js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryReportError calls the method "Window.reportError"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryReportError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowReportError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasBtoa returns true if the method "Window.btoa" exists.
func (this Window) HasBtoa() bool {
	return js.True == bindings.HasWindowBtoa(
		this.Ref(),
	)
}

// BtoaFunc returns the method "Window.btoa".
func (this Window) BtoaFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowBtoaFunc(
			this.Ref(),
		),
	)
}

// Btoa calls the method "Window.btoa".
func (this Window) Btoa(data js.String) (ret js.String) {
	bindings.CallWindowBtoa(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryBtoa calls the method "Window.btoa"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryBtoa(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowBtoa(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasAtob returns true if the method "Window.atob" exists.
func (this Window) HasAtob() bool {
	return js.True == bindings.HasWindowAtob(
		this.Ref(),
	)
}

// AtobFunc returns the method "Window.atob".
func (this Window) AtobFunc() (fn js.Func[func(data js.String) js.String]) {
	return fn.FromRef(
		bindings.WindowAtobFunc(
			this.Ref(),
		),
	)
}

// Atob calls the method "Window.atob".
func (this Window) Atob(data js.String) (ret js.String) {
	bindings.CallWindowAtob(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAtob calls the method "Window.atob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryAtob(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAtob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasSetTimeout returns true if the method "Window.setTimeout" exists.
func (this Window) HasSetTimeout() bool {
	return js.True == bindings.HasWindowSetTimeout(
		this.Ref(),
	)
}

// SetTimeoutFunc returns the method "Window.setTimeout".
func (this Window) SetTimeoutFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WindowSetTimeoutFunc(
			this.Ref(),
		),
	)
}

// SetTimeout calls the method "Window.setTimeout".
func (this Window) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWindowSetTimeout(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetTimeout calls the method "Window.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TrySetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetTimeout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasSetTimeout1 returns true if the method "Window.setTimeout" exists.
func (this Window) HasSetTimeout1() bool {
	return js.True == bindings.HasWindowSetTimeout1(
		this.Ref(),
	)
}

// SetTimeout1Func returns the method "Window.setTimeout".
func (this Window) SetTimeout1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WindowSetTimeout1Func(
			this.Ref(),
		),
	)
}

// SetTimeout1 calls the method "Window.setTimeout".
func (this Window) SetTimeout1(handler TimerHandler) (ret int32) {
	bindings.CallWindowSetTimeout1(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetTimeout1 calls the method "Window.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TrySetTimeout1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetTimeout1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasClearTimeout returns true if the method "Window.clearTimeout" exists.
func (this Window) HasClearTimeout() bool {
	return js.True == bindings.HasWindowClearTimeout(
		this.Ref(),
	)
}

// ClearTimeoutFunc returns the method "Window.clearTimeout".
func (this Window) ClearTimeoutFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WindowClearTimeoutFunc(
			this.Ref(),
		),
	)
}

// ClearTimeout calls the method "Window.clearTimeout".
func (this Window) ClearTimeout(id int32) (ret js.Void) {
	bindings.CallWindowClearTimeout(
		this.Ref(), js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearTimeout calls the method "Window.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryClearTimeout(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearTimeout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasClearTimeout1 returns true if the method "Window.clearTimeout" exists.
func (this Window) HasClearTimeout1() bool {
	return js.True == bindings.HasWindowClearTimeout1(
		this.Ref(),
	)
}

// ClearTimeout1Func returns the method "Window.clearTimeout".
func (this Window) ClearTimeout1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowClearTimeout1Func(
			this.Ref(),
		),
	)
}

// ClearTimeout1 calls the method "Window.clearTimeout".
func (this Window) ClearTimeout1() (ret js.Void) {
	bindings.CallWindowClearTimeout1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearTimeout1 calls the method "Window.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryClearTimeout1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearTimeout1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetInterval returns true if the method "Window.setInterval" exists.
func (this Window) HasSetInterval() bool {
	return js.True == bindings.HasWindowSetInterval(
		this.Ref(),
	)
}

// SetIntervalFunc returns the method "Window.setInterval".
func (this Window) SetIntervalFunc() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	return fn.FromRef(
		bindings.WindowSetIntervalFunc(
			this.Ref(),
		),
	)
}

// SetInterval calls the method "Window.setInterval".
func (this Window) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWindowSetInterval(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetInterval calls the method "Window.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TrySetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetInterval(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasSetInterval1 returns true if the method "Window.setInterval" exists.
func (this Window) HasSetInterval1() bool {
	return js.True == bindings.HasWindowSetInterval1(
		this.Ref(),
	)
}

// SetInterval1Func returns the method "Window.setInterval".
func (this Window) SetInterval1Func() (fn js.Func[func(handler TimerHandler) int32]) {
	return fn.FromRef(
		bindings.WindowSetInterval1Func(
			this.Ref(),
		),
	)
}

// SetInterval1 calls the method "Window.setInterval".
func (this Window) SetInterval1(handler TimerHandler) (ret int32) {
	bindings.CallWindowSetInterval1(
		this.Ref(), js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetInterval1 calls the method "Window.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TrySetInterval1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetInterval1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasClearInterval returns true if the method "Window.clearInterval" exists.
func (this Window) HasClearInterval() bool {
	return js.True == bindings.HasWindowClearInterval(
		this.Ref(),
	)
}

// ClearIntervalFunc returns the method "Window.clearInterval".
func (this Window) ClearIntervalFunc() (fn js.Func[func(id int32)]) {
	return fn.FromRef(
		bindings.WindowClearIntervalFunc(
			this.Ref(),
		),
	)
}

// ClearInterval calls the method "Window.clearInterval".
func (this Window) ClearInterval(id int32) (ret js.Void) {
	bindings.CallWindowClearInterval(
		this.Ref(), js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearInterval calls the method "Window.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryClearInterval(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearInterval(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasClearInterval1 returns true if the method "Window.clearInterval" exists.
func (this Window) HasClearInterval1() bool {
	return js.True == bindings.HasWindowClearInterval1(
		this.Ref(),
	)
}

// ClearInterval1Func returns the method "Window.clearInterval".
func (this Window) ClearInterval1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WindowClearInterval1Func(
			this.Ref(),
		),
	)
}

// ClearInterval1 calls the method "Window.clearInterval".
func (this Window) ClearInterval1() (ret js.Void) {
	bindings.CallWindowClearInterval1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearInterval1 calls the method "Window.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryClearInterval1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearInterval1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasQueueMicrotask returns true if the method "Window.queueMicrotask" exists.
func (this Window) HasQueueMicrotask() bool {
	return js.True == bindings.HasWindowQueueMicrotask(
		this.Ref(),
	)
}

// QueueMicrotaskFunc returns the method "Window.queueMicrotask".
func (this Window) QueueMicrotaskFunc() (fn js.Func[func(callback js.Func[func()])]) {
	return fn.FromRef(
		bindings.WindowQueueMicrotaskFunc(
			this.Ref(),
		),
	)
}

// QueueMicrotask calls the method "Window.queueMicrotask".
func (this Window) QueueMicrotask(callback js.Func[func()]) (ret js.Void) {
	bindings.CallWindowQueueMicrotask(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryQueueMicrotask calls the method "Window.queueMicrotask"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryQueueMicrotask(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueueMicrotask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCreateImageBitmap returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasCreateImageBitmap() bool {
	return js.True == bindings.HasWindowCreateImageBitmap(
		this.Ref(),
	)
}

// CreateImageBitmapFunc returns the method "Window.createImageBitmap".
func (this Window) CreateImageBitmapFunc() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmapFunc(
			this.Ref(),
		),
	)
}

// CreateImageBitmap calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateImageBitmap calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCreateImageBitmap1 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasCreateImageBitmap1() bool {
	return js.True == bindings.HasWindowCreateImageBitmap1(
		this.Ref(),
	)
}

// CreateImageBitmap1Func returns the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap1Func() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap1Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap1 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap1(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryCreateImageBitmap1 calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
}

// HasCreateImageBitmap2 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasCreateImageBitmap2() bool {
	return js.True == bindings.HasWindowCreateImageBitmap2(
		this.Ref(),
	)
}

// CreateImageBitmap2Func returns the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap2Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap2Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap2 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap2(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// TryCreateImageBitmap2 calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// HasCreateImageBitmap3 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasCreateImageBitmap3() bool {
	return js.True == bindings.HasWindowCreateImageBitmap3(
		this.Ref(),
	)
}

// CreateImageBitmap3Func returns the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap3Func() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	return fn.FromRef(
		bindings.WindowCreateImageBitmap3Func(
			this.Ref(),
		),
	)
}

// CreateImageBitmap3 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap3(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// TryCreateImageBitmap3 calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasStructuredClone returns true if the method "Window.structuredClone" exists.
func (this Window) HasStructuredClone() bool {
	return js.True == bindings.HasWindowStructuredClone(
		this.Ref(),
	)
}

// StructuredCloneFunc returns the method "Window.structuredClone".
func (this Window) StructuredCloneFunc() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	return fn.FromRef(
		bindings.WindowStructuredCloneFunc(
			this.Ref(),
		),
	)
}

// StructuredClone calls the method "Window.structuredClone".
func (this Window) StructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any) {
	bindings.CallWindowStructuredClone(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryStructuredClone calls the method "Window.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryStructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStructuredClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasStructuredClone1 returns true if the method "Window.structuredClone" exists.
func (this Window) HasStructuredClone1() bool {
	return js.True == bindings.HasWindowStructuredClone1(
		this.Ref(),
	)
}

// StructuredClone1Func returns the method "Window.structuredClone".
func (this Window) StructuredClone1Func() (fn js.Func[func(value js.Any) js.Any]) {
	return fn.FromRef(
		bindings.WindowStructuredClone1Func(
			this.Ref(),
		),
	)
}

// StructuredClone1 calls the method "Window.structuredClone".
func (this Window) StructuredClone1(value js.Any) (ret js.Any) {
	bindings.CallWindowStructuredClone1(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryStructuredClone1 calls the method "Window.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryStructuredClone1(value js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStructuredClone1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFetch returns true if the method "Window.fetch" exists.
func (this Window) HasFetch() bool {
	return js.True == bindings.HasWindowFetch(
		this.Ref(),
	)
}

// FetchFunc returns the method "Window.fetch".
func (this Window) FetchFunc() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WindowFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch calls the method "Window.fetch".
func (this Window) Fetch(input RequestInfo, init RequestInit) (ret js.Promise[Response]) {
	bindings.CallWindowFetch(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryFetch calls the method "Window.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryFetch(input RequestInfo, init RequestInit) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFetch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFetch1 returns true if the method "Window.fetch" exists.
func (this Window) HasFetch1() bool {
	return js.True == bindings.HasWindowFetch1(
		this.Ref(),
	)
}

// Fetch1Func returns the method "Window.fetch".
func (this Window) Fetch1Func() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	return fn.FromRef(
		bindings.WindowFetch1Func(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "Window.fetch".
func (this Window) Fetch1(input RequestInfo) (ret js.Promise[Response]) {
	bindings.CallWindowFetch1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFetch1 calls the method "Window.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryFetch1(input RequestInfo) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFetch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasRequestAnimationFrame returns true if the method "Window.requestAnimationFrame" exists.
func (this Window) HasRequestAnimationFrame() bool {
	return js.True == bindings.HasWindowRequestAnimationFrame(
		this.Ref(),
	)
}

// RequestAnimationFrameFunc returns the method "Window.requestAnimationFrame".
func (this Window) RequestAnimationFrameFunc() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	return fn.FromRef(
		bindings.WindowRequestAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// RequestAnimationFrame calls the method "Window.requestAnimationFrame".
func (this Window) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32) {
	bindings.CallWindowRequestAnimationFrame(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestAnimationFrame calls the method "Window.requestAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryRequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestAnimationFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCancelAnimationFrame returns true if the method "Window.cancelAnimationFrame" exists.
func (this Window) HasCancelAnimationFrame() bool {
	return js.True == bindings.HasWindowCancelAnimationFrame(
		this.Ref(),
	)
}

// CancelAnimationFrameFunc returns the method "Window.cancelAnimationFrame".
func (this Window) CancelAnimationFrameFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.WindowCancelAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// CancelAnimationFrame calls the method "Window.cancelAnimationFrame".
func (this Window) CancelAnimationFrame(handle uint32) (ret js.Void) {
	bindings.CallWindowCancelAnimationFrame(
		this.Ref(), js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelAnimationFrame calls the method "Window.cancelAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Window) TryCancelAnimationFrame(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCancelAnimationFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
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
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "CompositionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CompositionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CompositionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new CompositionEventInit in the application heap.
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

func NewCompositionEvent(typ js.String, eventInitDict CompositionEventInit) (ret CompositionEvent) {
	ret.ref = bindings.NewCompositionEventByCompositionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewCompositionEventByCompositionEvent1(typ js.String) (ret CompositionEvent) {
	ret.ref = bindings.NewCompositionEventByCompositionEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CompositionEvent) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCompositionEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInitCompositionEvent returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasInitCompositionEvent() bool {
	return js.True == bindings.HasCompositionEventInitCompositionEvent(
		this.Ref(),
	)
}

// InitCompositionEventFunc returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEventFunc(
			this.Ref(),
		),
	)
}

// InitCompositionEvent calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		dataArg.Ref(),
	)

	return
}

// TryInitCompositionEvent calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		dataArg.Ref(),
	)

	return
}

// HasInitCompositionEvent1 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasInitCompositionEvent1() bool {
	return js.True == bindings.HasCompositionEventInitCompositionEvent1(
		this.Ref(),
	)
}

// InitCompositionEvent1Func returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent1Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent1 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent1(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitCompositionEvent1 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasInitCompositionEvent2 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasInitCompositionEvent2() bool {
	return js.True == bindings.HasCompositionEventInitCompositionEvent2(
		this.Ref(),
	)
}

// InitCompositionEvent2Func returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent2Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent2 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent2(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitCompositionEvent2 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasInitCompositionEvent3 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasInitCompositionEvent3() bool {
	return js.True == bindings.HasCompositionEventInitCompositionEvent3(
		this.Ref(),
	)
}

// InitCompositionEvent3Func returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent3Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent3 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent3(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent3(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitCompositionEvent3 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent3(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasInitCompositionEvent4 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasInitCompositionEvent4() bool {
	return js.True == bindings.HasCompositionEventInitCompositionEvent4(
		this.Ref(),
	)
}

// InitCompositionEvent4Func returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent4Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.CompositionEventInitCompositionEvent4Func(
			this.Ref(),
		),
	)
}

// InitCompositionEvent4 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent4(typeArg js.String) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent4(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitCompositionEvent4 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent4(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
	)

	return
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

func NewCompressionStream(format CompressionFormat) (ret CompressionStream) {
	ret.ref = bindings.NewCompressionStreamByCompressionStream(
		uint32(format))
	return
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
// It returns ok=false if there is no such property.
func (this CompressionStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetCompressionStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "CompressionStream.writable".
//
// It returns ok=false if there is no such property.
func (this CompressionStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetCompressionStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ContentIndexEventInit struct {
	// Id is "ContentIndexEventInit.id"
	//
	// Required
	Id js.String
	// Bubbles is "ContentIndexEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ContentIndexEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ContentIndexEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new ContentIndexEventInit in the application heap.
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

func NewContentIndexEvent(typ js.String, init ContentIndexEventInit) (ret ContentIndexEvent) {
	ret.ref = bindings.NewContentIndexEventByContentIndexEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this ContentIndexEvent) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContentIndexEventId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ContentVisibilityAutoStateChangeEventInit struct {
	// Skipped is "ContentVisibilityAutoStateChangeEventInit.skipped"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Skipped MUST be set to true to make this field effective.
	Skipped bool
	// Bubbles is "ContentVisibilityAutoStateChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ContentVisibilityAutoStateChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ContentVisibilityAutoStateChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new ContentVisibilityAutoStateChangeEventInit in the application heap.
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

func NewContentVisibilityAutoStateChangeEvent(typ js.String, eventInitDict ContentVisibilityAutoStateChangeEventInit) (ret ContentVisibilityAutoStateChangeEvent) {
	ret.ref = bindings.NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent1(typ js.String) (ret ContentVisibilityAutoStateChangeEvent) {
	ret.ref = bindings.NewContentVisibilityAutoStateChangeEventByContentVisibilityAutoStateChangeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ContentVisibilityAutoStateChangeEvent) Skipped() (ret bool, ok bool) {
	ok = js.True == bindings.GetContentVisibilityAutoStateChangeEventSkipped(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CookieChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CookieChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new CookieChangeEventInit in the application heap.
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

func NewCookieChangeEvent(typ js.String, eventInitDict CookieChangeEventInit) (ret CookieChangeEvent) {
	ret.ref = bindings.NewCookieChangeEventByCookieChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewCookieChangeEventByCookieChangeEvent1(typ js.String) (ret CookieChangeEvent) {
	ret.ref = bindings.NewCookieChangeEventByCookieChangeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CookieChangeEvent) Changed() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetCookieChangeEventChanged(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Deleted returns the value of property "CookieChangeEvent.deleted".
//
// It returns ok=false if there is no such property.
func (this CookieChangeEvent) Deleted() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetCookieChangeEventDeleted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCountQueuingStrategy(init QueuingStrategyInit) (ret CountQueuingStrategy) {
	ret.ref = bindings.NewCountQueuingStrategyByCountQueuingStrategy(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this CountQueuingStrategy) HighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetCountQueuingStrategyHighWaterMark(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Size returns the value of property "CountQueuingStrategy.size".
//
// It returns ok=false if there is no such property.
func (this CountQueuingStrategy) Size() (ret js.Func[func(arguments ...js.Any) js.Any], ok bool) {
	ok = js.True == bindings.GetCountQueuingStrategySize(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CrashReportBody) Reason() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCrashReportBodyReason(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "CrashReportBody.toJSON" exists.
func (this CrashReportBody) HasToJSON() bool {
	return js.True == bindings.HasCrashReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "CrashReportBody.toJSON".
func (this CrashReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.CrashReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "CrashReportBody.toJSON".
func (this CrashReportBody) ToJSON() (ret js.Object) {
	bindings.CallCrashReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "CrashReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CrashReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCrashReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new CredentialData in the application heap.
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

// New creates a new CryptoKeyPair in the application heap.
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CustomEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CustomEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new CustomEventInit in the application heap.
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

func NewCustomEvent(typ js.String, eventInitDict CustomEventInit) (ret CustomEvent) {
	ret.ref = bindings.NewCustomEventByCustomEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewCustomEventByCustomEvent1(typ js.String) (ret CustomEvent) {
	ret.ref = bindings.NewCustomEventByCustomEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CustomEvent) Detail() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetCustomEventDetail(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInitCustomEvent returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasInitCustomEvent() bool {
	return js.True == bindings.HasCustomEventInitCustomEvent(
		this.Ref(),
	)
}

// InitCustomEventFunc returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, detail js.Any)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEventFunc(
			this.Ref(),
		),
	)
}

// InitCustomEvent calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent(typ js.String, bubbles bool, cancelable bool, detail js.Any) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		detail.Ref(),
	)

	return
}

// TryInitCustomEvent calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CustomEvent) TryInitCustomEvent(typ js.String, bubbles bool, cancelable bool, detail js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		detail.Ref(),
	)

	return
}

// HasInitCustomEvent1 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasInitCustomEvent1() bool {
	return js.True == bindings.HasCustomEventInitCustomEvent1(
		this.Ref(),
	)
}

// InitCustomEvent1Func returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent1Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent1Func(
			this.Ref(),
		),
	)
}

// InitCustomEvent1 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent1(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent1(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitCustomEvent1 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CustomEvent) TryInitCustomEvent1(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasInitCustomEvent2 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasInitCustomEvent2() bool {
	return js.True == bindings.HasCustomEventInitCustomEvent2(
		this.Ref(),
	)
}

// InitCustomEvent2Func returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent2Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent2Func(
			this.Ref(),
		),
	)
}

// InitCustomEvent2 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent2(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent2(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitCustomEvent2 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CustomEvent) TryInitCustomEvent2(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasInitCustomEvent3 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasInitCustomEvent3() bool {
	return js.True == bindings.HasCustomEventInitCustomEvent3(
		this.Ref(),
	)
}

// InitCustomEvent3Func returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent3Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.CustomEventInitCustomEvent3Func(
			this.Ref(),
		),
	)
}

// InitCustomEvent3 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent3(typ js.String) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent3(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitCustomEvent3 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CustomEvent) TryInitCustomEvent3(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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

// HasParseFromString returns true if the method "DOMParser.parseFromString" exists.
func (this DOMParser) HasParseFromString() bool {
	return js.True == bindings.HasDOMParserParseFromString(
		this.Ref(),
	)
}

// ParseFromStringFunc returns the method "DOMParser.parseFromString".
func (this DOMParser) ParseFromStringFunc() (fn js.Func[func(string js.String, typ DOMParserSupportedType) Document]) {
	return fn.FromRef(
		bindings.DOMParserParseFromStringFunc(
			this.Ref(),
		),
	)
}

// ParseFromString calls the method "DOMParser.parseFromString".
func (this DOMParser) ParseFromString(string js.String, typ DOMParserSupportedType) (ret Document) {
	bindings.CallDOMParserParseFromString(
		this.Ref(), js.Pointer(&ret),
		string.Ref(),
		uint32(typ),
	)

	return
}

// TryParseFromString calls the method "DOMParser.parseFromString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DOMParser) TryParseFromString(string js.String, typ DOMParserSupportedType) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMParserParseFromString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		string.Ref(),
		uint32(typ),
	)

	return
}

func NewDataCue(startTime float64, endTime float64, value js.Any, typ js.String) (ret DataCue) {
	ret.ref = bindings.NewDataCueByDataCue(
		float64(startTime),
		float64(endTime),
		value.Ref(),
		typ.Ref())
	return
}

func NewDataCueByDataCue1(startTime float64, endTime float64, value js.Any) (ret DataCue) {
	ret.ref = bindings.NewDataCueByDataCue1(
		float64(startTime),
		float64(endTime),
		value.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this DataCue) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetDataCueValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "DataCue.value" to val.
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
// It returns ok=false if there is no such property.
func (this DataCue) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataCueType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewDecompressionStream(format CompressionFormat) (ret DecompressionStream) {
	ret.ref = bindings.NewDecompressionStreamByDecompressionStream(
		uint32(format))
	return
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
// It returns ok=false if there is no such property.
func (this DecompressionStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetDecompressionStreamReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "DecompressionStream.writable".
//
// It returns ok=false if there is no such property.
func (this DecompressionStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetDecompressionStreamWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this DedicatedWorkerGlobalScope) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDedicatedWorkerGlobalScopeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPostMessage returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasPostMessage() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopePostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasPostMessage1() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopePostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage2 returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasPostMessage2() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopePostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopePostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasClose returns true if the method "DedicatedWorkerGlobalScope.close" exists.
func (this DedicatedWorkerGlobalScope) HasClose() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopeClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "DedicatedWorkerGlobalScope.close".
func (this DedicatedWorkerGlobalScope) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "DedicatedWorkerGlobalScope.close".
func (this DedicatedWorkerGlobalScope) Close() (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopeClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "DedicatedWorkerGlobalScope.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestAnimationFrame returns true if the method "DedicatedWorkerGlobalScope.requestAnimationFrame" exists.
func (this DedicatedWorkerGlobalScope) HasRequestAnimationFrame() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.Ref(),
	)
}

// RequestAnimationFrameFunc returns the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
func (this DedicatedWorkerGlobalScope) RequestAnimationFrameFunc() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeRequestAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// RequestAnimationFrame calls the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
func (this DedicatedWorkerGlobalScope) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32) {
	bindings.CallDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestAnimationFrame calls the method "DedicatedWorkerGlobalScope.requestAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryRequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCancelAnimationFrame returns true if the method "DedicatedWorkerGlobalScope.cancelAnimationFrame" exists.
func (this DedicatedWorkerGlobalScope) HasCancelAnimationFrame() bool {
	return js.True == bindings.HasDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.Ref(),
	)
}

// CancelAnimationFrameFunc returns the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
func (this DedicatedWorkerGlobalScope) CancelAnimationFrameFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.DedicatedWorkerGlobalScopeCancelAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// CancelAnimationFrame calls the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
func (this DedicatedWorkerGlobalScope) CancelAnimationFrame(handle uint32) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.Ref(), js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelAnimationFrame calls the method "DedicatedWorkerGlobalScope.cancelAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryCancelAnimationFrame(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
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
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnticipatedRemoval returns the value of property "DeprecationReportBody.anticipatedRemoval".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) AnticipatedRemoval() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyAnticipatedRemoval(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "DeprecationReportBody.message".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "DeprecationReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodySourceFile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "DeprecationReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyLineNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "DeprecationReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "DeprecationReportBody.toJSON" exists.
func (this DeprecationReportBody) HasToJSON() bool {
	return js.True == bindings.HasDeprecationReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "DeprecationReportBody.toJSON".
func (this DeprecationReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DeprecationReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DeprecationReportBody.toJSON".
func (this DeprecationReportBody) ToJSON() (ret js.Object) {
	bindings.CallDeprecationReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DeprecationReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DeprecationReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeprecationReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new Landmark in the application heap.
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

// New creates a new DetectedFace in the application heap.
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

// New creates a new DetectedText in the application heap.
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
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "DeviceMotionEventAccelerationInit.y"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64
	// Z is "DeviceMotionEventAccelerationInit.z"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Z MUST be set to true to make this field effective.
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

// New creates a new DeviceMotionEventAccelerationInit in the application heap.
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
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float64
	// Beta is "DeviceMotionEventRotationRateInit.beta"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Beta MUST be set to true to make this field effective.
	Beta float64
	// Gamma is "DeviceMotionEventRotationRateInit.gamma"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Gamma MUST be set to true to make this field effective.
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

// New creates a new DeviceMotionEventRotationRateInit in the application heap.
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
	//
	// NOTE: Acceleration.FFI_USE MUST be set to true to get Acceleration used.
	Acceleration DeviceMotionEventAccelerationInit
	// AccelerationIncludingGravity is "DeviceMotionEventInit.accelerationIncludingGravity"
	//
	// Optional
	//
	// NOTE: AccelerationIncludingGravity.FFI_USE MUST be set to true to get AccelerationIncludingGravity used.
	AccelerationIncludingGravity DeviceMotionEventAccelerationInit
	// RotationRate is "DeviceMotionEventInit.rotationRate"
	//
	// Optional
	//
	// NOTE: RotationRate.FFI_USE MUST be set to true to get RotationRate used.
	RotationRate DeviceMotionEventRotationRateInit
	// Interval is "DeviceMotionEventInit.interval"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Interval MUST be set to true to make this field effective.
	Interval float64
	// Bubbles is "DeviceMotionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "DeviceMotionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "DeviceMotionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new DeviceMotionEventInit in the application heap.
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
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DeviceMotionEventAcceleration.y".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "DeviceMotionEventAcceleration.z".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Alpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Beta returns the value of property "DeviceMotionEventRotationRate.beta".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Beta() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateBeta(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Gamma returns the value of property "DeviceMotionEventRotationRate.gamma".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Gamma() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateGamma(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewDeviceMotionEvent(typ js.String, eventInitDict DeviceMotionEventInit) (ret DeviceMotionEvent) {
	ret.ref = bindings.NewDeviceMotionEventByDeviceMotionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewDeviceMotionEventByDeviceMotionEvent1(typ js.String) (ret DeviceMotionEvent) {
	ret.ref = bindings.NewDeviceMotionEventByDeviceMotionEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) Acceleration() (ret DeviceMotionEventAcceleration, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAcceleration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AccelerationIncludingGravity returns the value of property "DeviceMotionEvent.accelerationIncludingGravity".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) AccelerationIncludingGravity() (ret DeviceMotionEventAcceleration, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationIncludingGravity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RotationRate returns the value of property "DeviceMotionEvent.rotationRate".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) RotationRate() (ret DeviceMotionEventRotationRate, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Interval returns the value of property "DeviceMotionEvent.interval".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) Interval() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventInterval(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestPermission returns true if the staticmethod "DeviceMotionEvent.requestPermission" exists.
func (this DeviceMotionEvent) HasRequestPermission() bool {
	return js.True == bindings.HasDeviceMotionEventRequestPermission(
		this.Ref(),
	)
}

// RequestPermissionFunc returns the staticmethod "DeviceMotionEvent.requestPermission".
func (this DeviceMotionEvent) RequestPermissionFunc() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.DeviceMotionEventRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// RequestPermission calls the staticmethod "DeviceMotionEvent.requestPermission".
func (this DeviceMotionEvent) RequestPermission() (ret js.Promise[PermissionState]) {
	bindings.CallDeviceMotionEventRequestPermission(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPermission calls the staticmethod "DeviceMotionEvent.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DeviceMotionEvent) TryRequestPermission() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeviceMotionEventRequestPermission(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

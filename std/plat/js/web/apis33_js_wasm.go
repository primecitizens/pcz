// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *PerformanceMeasureOptions) UpdateFrom(ref js.Ref) {
	bindings.PerformanceMeasureOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PerformanceMeasureOptions) Update(ref js.Ref) {
	bindings.PerformanceMeasureOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PerformanceMeasureOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Detail.Ref(),
		p.Start.Ref(),
		p.End.Ref(),
	)
	p.Detail = p.Detail.FromRef(js.Undefined)
	p.Start = p.Start.FromRef(js.Undefined)
	p.End = p.End.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "PerformanceEntry.name".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EntryType returns the value of property "PerformanceEntry.entryType".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) EntryType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryEntryType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StartTime returns the value of property "PerformanceEntry.startTime".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) StartTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryStartTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "PerformanceEntry.duration".
//
// It returns ok=false if there is no such property.
func (this PerformanceEntry) Duration() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEntryDuration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "PerformanceEntry.toJSON" exists.
func (this PerformanceEntry) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPerformanceEntryToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "PerformanceEntry.toJSON".
func (this PerformanceEntry) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncPerformanceEntryToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "PerformanceEntry.toJSON".
func (this PerformanceEntry) ToJSON() (ret js.Object) {
	bindings.CallPerformanceEntryToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceEntry.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceEntry) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceEntryToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceEntryList = js.Array[PerformanceEntry]

type PerformanceTiming struct {
	ref js.Ref
}

func (this PerformanceTiming) Once() PerformanceTiming {
	this.ref.Once()
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
	this.ref.Free()
}

// NavigationStart returns the value of property "PerformanceTiming.navigationStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) NavigationStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingNavigationStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UnloadEventStart returns the value of property "PerformanceTiming.unloadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) UnloadEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingUnloadEventStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UnloadEventEnd returns the value of property "PerformanceTiming.unloadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) UnloadEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingUnloadEventEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RedirectStart returns the value of property "PerformanceTiming.redirectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RedirectStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRedirectStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RedirectEnd returns the value of property "PerformanceTiming.redirectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RedirectEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRedirectEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FetchStart returns the value of property "PerformanceTiming.fetchStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) FetchStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingFetchStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomainLookupStart returns the value of property "PerformanceTiming.domainLookupStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomainLookupStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomainLookupStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomainLookupEnd returns the value of property "PerformanceTiming.domainLookupEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomainLookupEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomainLookupEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ConnectStart returns the value of property "PerformanceTiming.connectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ConnectStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingConnectStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ConnectEnd returns the value of property "PerformanceTiming.connectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ConnectEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingConnectEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SecureConnectionStart returns the value of property "PerformanceTiming.secureConnectionStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) SecureConnectionStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingSecureConnectionStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RequestStart returns the value of property "PerformanceTiming.requestStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) RequestStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingRequestStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseStart returns the value of property "PerformanceTiming.responseStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ResponseStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingResponseStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseEnd returns the value of property "PerformanceTiming.responseEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) ResponseEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingResponseEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomLoading returns the value of property "PerformanceTiming.domLoading".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomLoading() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomLoading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomInteractive returns the value of property "PerformanceTiming.domInteractive".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomInteractive() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomInteractive(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventStart returns the value of property "PerformanceTiming.domContentLoadedEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomContentLoadedEventStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventEnd returns the value of property "PerformanceTiming.domContentLoadedEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomContentLoadedEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomContentLoadedEventEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomComplete returns the value of property "PerformanceTiming.domComplete".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) DomComplete() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingDomComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LoadEventStart returns the value of property "PerformanceTiming.loadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) LoadEventStart() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingLoadEventStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LoadEventEnd returns the value of property "PerformanceTiming.loadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceTiming) LoadEventEnd() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceTimingLoadEventEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "PerformanceTiming.toJSON" exists.
func (this PerformanceTiming) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPerformanceTimingToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "PerformanceTiming.toJSON".
func (this PerformanceTiming) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncPerformanceTimingToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "PerformanceTiming.toJSON".
func (this PerformanceTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceTimingToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceTimingToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "PerformanceNavigation.type".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigation) Type() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RedirectCount returns the value of property "PerformanceNavigation.redirectCount".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigation) RedirectCount() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationRedirectCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "PerformanceNavigation.toJSON" exists.
func (this PerformanceNavigation) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPerformanceNavigationToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "PerformanceNavigation.toJSON".
func (this PerformanceNavigation) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncPerformanceNavigationToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "PerformanceNavigation.toJSON".
func (this PerformanceNavigation) ToJSON() (ret js.Object) {
	bindings.CallPerformanceNavigationToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceNavigation.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceNavigation) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceNavigationToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type EventCounts struct {
	ref js.Ref
}

func (this EventCounts) Once() EventCounts {
	this.ref.Once()
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
	this.ref.Free()
}

type Performance struct {
	EventTarget
}

func (this Performance) Once() Performance {
	this.ref.Once()
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
	this.ref.Free()
}

// TimeOrigin returns the value of property "Performance.timeOrigin".
//
// It returns ok=false if there is no such property.
func (this Performance) TimeOrigin() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceTimeOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timing returns the value of property "Performance.timing".
//
// It returns ok=false if there is no such property.
func (this Performance) Timing() (ret PerformanceTiming, ok bool) {
	ok = js.True == bindings.GetPerformanceTiming(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Navigation returns the value of property "Performance.navigation".
//
// It returns ok=false if there is no such property.
func (this Performance) Navigation() (ret PerformanceNavigation, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EventCounts returns the value of property "Performance.eventCounts".
//
// It returns ok=false if there is no such property.
func (this Performance) EventCounts() (ret EventCounts, ok bool) {
	ok = js.True == bindings.GetPerformanceEventCounts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InteractionCount returns the value of property "Performance.interactionCount".
//
// It returns ok=false if there is no such property.
func (this Performance) InteractionCount() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceInteractionCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncNow returns true if the method "Performance.now" exists.
func (this Performance) HasFuncNow() bool {
	return js.True == bindings.HasFuncPerformanceNow(
		this.ref,
	)
}

// FuncNow returns the method "Performance.now".
func (this Performance) FuncNow() (fn js.Func[func() DOMHighResTimeStamp]) {
	bindings.FuncPerformanceNow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Now calls the method "Performance.now".
func (this Performance) Now() (ret DOMHighResTimeStamp) {
	bindings.CallPerformanceNow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryNow calls the method "Performance.now"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryNow() (ret DOMHighResTimeStamp, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceNow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "Performance.toJSON" exists.
func (this Performance) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPerformanceToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "Performance.toJSON".
func (this Performance) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncPerformanceToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "Performance.toJSON".
func (this Performance) ToJSON() (ret js.Object) {
	bindings.CallPerformanceToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "Performance.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMeasureUserAgentSpecificMemory returns true if the method "Performance.measureUserAgentSpecificMemory" exists.
func (this Performance) HasFuncMeasureUserAgentSpecificMemory() bool {
	return js.True == bindings.HasFuncPerformanceMeasureUserAgentSpecificMemory(
		this.ref,
	)
}

// FuncMeasureUserAgentSpecificMemory returns the method "Performance.measureUserAgentSpecificMemory".
func (this Performance) FuncMeasureUserAgentSpecificMemory() (fn js.Func[func() js.Promise[MemoryMeasurement]]) {
	bindings.FuncPerformanceMeasureUserAgentSpecificMemory(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MeasureUserAgentSpecificMemory calls the method "Performance.measureUserAgentSpecificMemory".
func (this Performance) MeasureUserAgentSpecificMemory() (ret js.Promise[MemoryMeasurement]) {
	bindings.CallPerformanceMeasureUserAgentSpecificMemory(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMeasureUserAgentSpecificMemory calls the method "Performance.measureUserAgentSpecificMemory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMeasureUserAgentSpecificMemory() (ret js.Promise[MemoryMeasurement], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasureUserAgentSpecificMemory(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMark returns true if the method "Performance.mark" exists.
func (this Performance) HasFuncMark() bool {
	return js.True == bindings.HasFuncPerformanceMark(
		this.ref,
	)
}

// FuncMark returns the method "Performance.mark".
func (this Performance) FuncMark() (fn js.Func[func(markName js.String, markOptions PerformanceMarkOptions) PerformanceMark]) {
	bindings.FuncPerformanceMark(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Mark calls the method "Performance.mark".
func (this Performance) Mark(markName js.String, markOptions PerformanceMarkOptions) (ret PerformanceMark) {
	bindings.CallPerformanceMark(
		this.ref, js.Pointer(&ret),
		markName.Ref(),
		js.Pointer(&markOptions),
	)

	return
}

// TryMark calls the method "Performance.mark"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMark(markName js.String, markOptions PerformanceMarkOptions) (ret PerformanceMark, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMark(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
		js.Pointer(&markOptions),
	)

	return
}

// HasFuncMark1 returns true if the method "Performance.mark" exists.
func (this Performance) HasFuncMark1() bool {
	return js.True == bindings.HasFuncPerformanceMark1(
		this.ref,
	)
}

// FuncMark1 returns the method "Performance.mark".
func (this Performance) FuncMark1() (fn js.Func[func(markName js.String) PerformanceMark]) {
	bindings.FuncPerformanceMark1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Mark1 calls the method "Performance.mark".
func (this Performance) Mark1(markName js.String) (ret PerformanceMark) {
	bindings.CallPerformanceMark1(
		this.ref, js.Pointer(&ret),
		markName.Ref(),
	)

	return
}

// TryMark1 calls the method "Performance.mark"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMark1(markName js.String) (ret PerformanceMark, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMark1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
	)

	return
}

// HasFuncClearMarks returns true if the method "Performance.clearMarks" exists.
func (this Performance) HasFuncClearMarks() bool {
	return js.True == bindings.HasFuncPerformanceClearMarks(
		this.ref,
	)
}

// FuncClearMarks returns the method "Performance.clearMarks".
func (this Performance) FuncClearMarks() (fn js.Func[func(markName js.String)]) {
	bindings.FuncPerformanceClearMarks(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearMarks calls the method "Performance.clearMarks".
func (this Performance) ClearMarks(markName js.String) (ret js.Void) {
	bindings.CallPerformanceClearMarks(
		this.ref, js.Pointer(&ret),
		markName.Ref(),
	)

	return
}

// TryClearMarks calls the method "Performance.clearMarks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryClearMarks(markName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMarks(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		markName.Ref(),
	)

	return
}

// HasFuncClearMarks1 returns true if the method "Performance.clearMarks" exists.
func (this Performance) HasFuncClearMarks1() bool {
	return js.True == bindings.HasFuncPerformanceClearMarks1(
		this.ref,
	)
}

// FuncClearMarks1 returns the method "Performance.clearMarks".
func (this Performance) FuncClearMarks1() (fn js.Func[func()]) {
	bindings.FuncPerformanceClearMarks1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearMarks1 calls the method "Performance.clearMarks".
func (this Performance) ClearMarks1() (ret js.Void) {
	bindings.CallPerformanceClearMarks1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearMarks1 calls the method "Performance.clearMarks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryClearMarks1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMarks1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMeasure returns true if the method "Performance.measure" exists.
func (this Performance) HasFuncMeasure() bool {
	return js.True == bindings.HasFuncPerformanceMeasure(
		this.ref,
	)
}

// FuncMeasure returns the method "Performance.measure".
func (this Performance) FuncMeasure() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) PerformanceMeasure]) {
	bindings.FuncPerformanceMeasure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Measure calls the method "Performance.measure".
func (this Performance) Measure(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure(
		this.ref, js.Pointer(&ret),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
		endMark.Ref(),
	)

	return
}

// TryMeasure calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMeasure(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions, endMark js.String) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
		endMark.Ref(),
	)

	return
}

// HasFuncMeasure1 returns true if the method "Performance.measure" exists.
func (this Performance) HasFuncMeasure1() bool {
	return js.True == bindings.HasFuncPerformanceMeasure1(
		this.ref,
	)
}

// FuncMeasure1 returns the method "Performance.measure".
func (this Performance) FuncMeasure1() (fn js.Func[func(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) PerformanceMeasure]) {
	bindings.FuncPerformanceMeasure1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Measure1 calls the method "Performance.measure".
func (this Performance) Measure1(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure1(
		this.ref, js.Pointer(&ret),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
	)

	return
}

// TryMeasure1 calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMeasure1(measureName js.String, startOrMeasureOptions OneOf_String_PerformanceMeasureOptions) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
		startOrMeasureOptions.Ref(),
	)

	return
}

// HasFuncMeasure2 returns true if the method "Performance.measure" exists.
func (this Performance) HasFuncMeasure2() bool {
	return js.True == bindings.HasFuncPerformanceMeasure2(
		this.ref,
	)
}

// FuncMeasure2 returns the method "Performance.measure".
func (this Performance) FuncMeasure2() (fn js.Func[func(measureName js.String) PerformanceMeasure]) {
	bindings.FuncPerformanceMeasure2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Measure2 calls the method "Performance.measure".
func (this Performance) Measure2(measureName js.String) (ret PerformanceMeasure) {
	bindings.CallPerformanceMeasure2(
		this.ref, js.Pointer(&ret),
		measureName.Ref(),
	)

	return
}

// TryMeasure2 calls the method "Performance.measure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryMeasure2(measureName js.String) (ret PerformanceMeasure, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceMeasure2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
	)

	return
}

// HasFuncClearMeasures returns true if the method "Performance.clearMeasures" exists.
func (this Performance) HasFuncClearMeasures() bool {
	return js.True == bindings.HasFuncPerformanceClearMeasures(
		this.ref,
	)
}

// FuncClearMeasures returns the method "Performance.clearMeasures".
func (this Performance) FuncClearMeasures() (fn js.Func[func(measureName js.String)]) {
	bindings.FuncPerformanceClearMeasures(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearMeasures calls the method "Performance.clearMeasures".
func (this Performance) ClearMeasures(measureName js.String) (ret js.Void) {
	bindings.CallPerformanceClearMeasures(
		this.ref, js.Pointer(&ret),
		measureName.Ref(),
	)

	return
}

// TryClearMeasures calls the method "Performance.clearMeasures"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryClearMeasures(measureName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMeasures(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		measureName.Ref(),
	)

	return
}

// HasFuncClearMeasures1 returns true if the method "Performance.clearMeasures" exists.
func (this Performance) HasFuncClearMeasures1() bool {
	return js.True == bindings.HasFuncPerformanceClearMeasures1(
		this.ref,
	)
}

// FuncClearMeasures1 returns the method "Performance.clearMeasures".
func (this Performance) FuncClearMeasures1() (fn js.Func[func()]) {
	bindings.FuncPerformanceClearMeasures1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearMeasures1 calls the method "Performance.clearMeasures".
func (this Performance) ClearMeasures1() (ret js.Void) {
	bindings.CallPerformanceClearMeasures1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearMeasures1 calls the method "Performance.clearMeasures"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryClearMeasures1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearMeasures1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClearResourceTimings returns true if the method "Performance.clearResourceTimings" exists.
func (this Performance) HasFuncClearResourceTimings() bool {
	return js.True == bindings.HasFuncPerformanceClearResourceTimings(
		this.ref,
	)
}

// FuncClearResourceTimings returns the method "Performance.clearResourceTimings".
func (this Performance) FuncClearResourceTimings() (fn js.Func[func()]) {
	bindings.FuncPerformanceClearResourceTimings(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearResourceTimings calls the method "Performance.clearResourceTimings".
func (this Performance) ClearResourceTimings() (ret js.Void) {
	bindings.CallPerformanceClearResourceTimings(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearResourceTimings calls the method "Performance.clearResourceTimings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryClearResourceTimings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceClearResourceTimings(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetResourceTimingBufferSize returns true if the method "Performance.setResourceTimingBufferSize" exists.
func (this Performance) HasFuncSetResourceTimingBufferSize() bool {
	return js.True == bindings.HasFuncPerformanceSetResourceTimingBufferSize(
		this.ref,
	)
}

// FuncSetResourceTimingBufferSize returns the method "Performance.setResourceTimingBufferSize".
func (this Performance) FuncSetResourceTimingBufferSize() (fn js.Func[func(maxSize uint32)]) {
	bindings.FuncPerformanceSetResourceTimingBufferSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetResourceTimingBufferSize calls the method "Performance.setResourceTimingBufferSize".
func (this Performance) SetResourceTimingBufferSize(maxSize uint32) (ret js.Void) {
	bindings.CallPerformanceSetResourceTimingBufferSize(
		this.ref, js.Pointer(&ret),
		uint32(maxSize),
	)

	return
}

// TrySetResourceTimingBufferSize calls the method "Performance.setResourceTimingBufferSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TrySetResourceTimingBufferSize(maxSize uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceSetResourceTimingBufferSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(maxSize),
	)

	return
}

// HasFuncGetEntries returns true if the method "Performance.getEntries" exists.
func (this Performance) HasFuncGetEntries() bool {
	return js.True == bindings.HasFuncPerformanceGetEntries(
		this.ref,
	)
}

// FuncGetEntries returns the method "Performance.getEntries".
func (this Performance) FuncGetEntries() (fn js.Func[func() PerformanceEntryList]) {
	bindings.FuncPerformanceGetEntries(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetEntries calls the method "Performance.getEntries".
func (this Performance) GetEntries() (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntries(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetEntries calls the method "Performance.getEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryGetEntries() (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntries(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetEntriesByType returns true if the method "Performance.getEntriesByType" exists.
func (this Performance) HasFuncGetEntriesByType() bool {
	return js.True == bindings.HasFuncPerformanceGetEntriesByType(
		this.ref,
	)
}

// FuncGetEntriesByType returns the method "Performance.getEntriesByType".
func (this Performance) FuncGetEntriesByType() (fn js.Func[func(typ js.String) PerformanceEntryList]) {
	bindings.FuncPerformanceGetEntriesByType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetEntriesByType calls the method "Performance.getEntriesByType".
func (this Performance) GetEntriesByType(typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByType(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByType calls the method "Performance.getEntriesByType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryGetEntriesByType(typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncGetEntriesByName returns true if the method "Performance.getEntriesByName" exists.
func (this Performance) HasFuncGetEntriesByName() bool {
	return js.True == bindings.HasFuncPerformanceGetEntriesByName(
		this.ref,
	)
}

// FuncGetEntriesByName returns the method "Performance.getEntriesByName".
func (this Performance) FuncGetEntriesByName() (fn js.Func[func(name js.String, typ js.String) PerformanceEntryList]) {
	bindings.FuncPerformanceGetEntriesByName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetEntriesByName calls the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByName(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByName calls the method "Performance.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryGetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// HasFuncGetEntriesByName1 returns true if the method "Performance.getEntriesByName" exists.
func (this Performance) HasFuncGetEntriesByName1() bool {
	return js.True == bindings.HasFuncPerformanceGetEntriesByName1(
		this.ref,
	)
}

// FuncGetEntriesByName1 returns the method "Performance.getEntriesByName".
func (this Performance) FuncGetEntriesByName1() (fn js.Func[func(name js.String) PerformanceEntryList]) {
	bindings.FuncPerformanceGetEntriesByName1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetEntriesByName1 calls the method "Performance.getEntriesByName".
func (this Performance) GetEntriesByName1(name js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceGetEntriesByName1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetEntriesByName1 calls the method "Performance.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Performance) TryGetEntriesByName1(name js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceGetEntriesByName1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type Storage struct {
	ref js.Ref
}

func (this Storage) Once() Storage {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "Storage.length".
//
// It returns ok=false if there is no such property.
func (this Storage) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetStorageLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncKey returns true if the method "Storage.key" exists.
func (this Storage) HasFuncKey() bool {
	return js.True == bindings.HasFuncStorageKey(
		this.ref,
	)
}

// FuncKey returns the method "Storage.key".
func (this Storage) FuncKey() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncStorageKey(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Key calls the method "Storage.key".
func (this Storage) Key(index uint32) (ret js.String) {
	bindings.CallStorageKey(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryKey calls the method "Storage.key"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Storage) TryKey(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageKey(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetItem returns true if the method "Storage.getItem" exists.
func (this Storage) HasFuncGetItem() bool {
	return js.True == bindings.HasFuncStorageGetItem(
		this.ref,
	)
}

// FuncGetItem returns the method "Storage.getItem".
func (this Storage) FuncGetItem() (fn js.Func[func(key js.String) js.String]) {
	bindings.FuncStorageGetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetItem calls the method "Storage.getItem".
func (this Storage) GetItem(key js.String) (ret js.String) {
	bindings.CallStorageGetItem(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryGetItem calls the method "Storage.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Storage) TryGetItem(key js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageGetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncSetItem returns true if the method "Storage.setItem" exists.
func (this Storage) HasFuncSetItem() bool {
	return js.True == bindings.HasFuncStorageSetItem(
		this.ref,
	)
}

// FuncSetItem returns the method "Storage.setItem".
func (this Storage) FuncSetItem() (fn js.Func[func(key js.String, value js.String)]) {
	bindings.FuncStorageSetItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetItem calls the method "Storage.setItem".
func (this Storage) SetItem(key js.String, value js.String) (ret js.Void) {
	bindings.CallStorageSetItem(
		this.ref, js.Pointer(&ret),
		key.Ref(),
		value.Ref(),
	)

	return
}

// TrySetItem calls the method "Storage.setItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Storage) TrySetItem(key js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageSetItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncRemoveItem returns true if the method "Storage.removeItem" exists.
func (this Storage) HasFuncRemoveItem() bool {
	return js.True == bindings.HasFuncStorageRemoveItem(
		this.ref,
	)
}

// FuncRemoveItem returns the method "Storage.removeItem".
func (this Storage) FuncRemoveItem() (fn js.Func[func(key js.String)]) {
	bindings.FuncStorageRemoveItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveItem calls the method "Storage.removeItem".
func (this Storage) RemoveItem(key js.String) (ret js.Void) {
	bindings.CallStorageRemoveItem(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryRemoveItem calls the method "Storage.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Storage) TryRemoveItem(key js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageRemoveItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "Storage.clear" exists.
func (this Storage) HasFuncClear() bool {
	return js.True == bindings.HasFuncStorageClear(
		this.ref,
	)
}

// FuncClear returns the method "Storage.clear".
func (this Storage) FuncClear() (fn js.Func[func()]) {
	bindings.FuncStorageClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "Storage.clear".
func (this Storage) Clear() (ret js.Void) {
	bindings.CallStorageClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "Storage.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Storage) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Window struct {
	EventTarget
}

func (this Window) Once() Window {
	this.ref.Once()
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
	this.ref.Free()
}

// Window returns the value of property "Window.window".
//
// It returns ok=false if there is no such property.
func (this Window) Window() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Self returns the value of property "Window.self".
//
// It returns ok=false if there is no such property.
func (this Window) Self() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowSelf(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Document returns the value of property "Window.document".
//
// It returns ok=false if there is no such property.
func (this Window) Document() (ret Document, ok bool) {
	ok = js.True == bindings.GetWindowDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "Window.name".
//
// It returns ok=false if there is no such property.
func (this Window) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "Window.name" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetName(val js.String) bool {
	return js.True == bindings.SetWindowName(
		this.ref,
		val.Ref(),
	)
}

// Location returns the value of property "Window.location".
//
// It returns ok=false if there is no such property.
func (this Window) Location() (ret Location, ok bool) {
	ok = js.True == bindings.GetWindowLocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// History returns the value of property "Window.history".
//
// It returns ok=false if there is no such property.
func (this Window) History() (ret History, ok bool) {
	ok = js.True == bindings.GetWindowHistory(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Navigation returns the value of property "Window.navigation".
//
// It returns ok=false if there is no such property.
func (this Window) Navigation() (ret Navigation, ok bool) {
	ok = js.True == bindings.GetWindowNavigation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CustomElements returns the value of property "Window.customElements".
//
// It returns ok=false if there is no such property.
func (this Window) CustomElements() (ret CustomElementRegistry, ok bool) {
	ok = js.True == bindings.GetWindowCustomElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Locationbar returns the value of property "Window.locationbar".
//
// It returns ok=false if there is no such property.
func (this Window) Locationbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowLocationbar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Menubar returns the value of property "Window.menubar".
//
// It returns ok=false if there is no such property.
func (this Window) Menubar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowMenubar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Personalbar returns the value of property "Window.personalbar".
//
// It returns ok=false if there is no such property.
func (this Window) Personalbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowPersonalbar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scrollbars returns the value of property "Window.scrollbars".
//
// It returns ok=false if there is no such property.
func (this Window) Scrollbars() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowScrollbars(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Statusbar returns the value of property "Window.statusbar".
//
// It returns ok=false if there is no such property.
func (this Window) Statusbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowStatusbar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Toolbar returns the value of property "Window.toolbar".
//
// It returns ok=false if there is no such property.
func (this Window) Toolbar() (ret BarProp, ok bool) {
	ok = js.True == bindings.GetWindowToolbar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "Window.status".
//
// It returns ok=false if there is no such property.
func (this Window) Status() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStatus sets the value of property "Window.status" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetStatus(val js.String) bool {
	return js.True == bindings.SetWindowStatus(
		this.ref,
		val.Ref(),
	)
}

// Closed returns the value of property "Window.closed".
//
// It returns ok=false if there is no such property.
func (this Window) Closed() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Frames returns the value of property "Window.frames".
//
// It returns ok=false if there is no such property.
func (this Window) Frames() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowFrames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "Window.length".
//
// It returns ok=false if there is no such property.
func (this Window) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetWindowLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Top returns the value of property "Window.top".
//
// It returns ok=false if there is no such property.
func (this Window) Top() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Opener returns the value of property "Window.opener".
//
// It returns ok=false if there is no such property.
func (this Window) Opener() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetWindowOpener(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOpener sets the value of property "Window.opener" to val.
//
// It returns false if the property cannot be set.
func (this Window) SetOpener(val js.Any) bool {
	return js.True == bindings.SetWindowOpener(
		this.ref,
		val.Ref(),
	)
}

// Parent returns the value of property "Window.parent".
//
// It returns ok=false if there is no such property.
func (this Window) Parent() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetWindowParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FrameElement returns the value of property "Window.frameElement".
//
// It returns ok=false if there is no such property.
func (this Window) FrameElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetWindowFrameElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Navigator returns the value of property "Window.navigator".
//
// It returns ok=false if there is no such property.
func (this Window) Navigator() (ret Navigator, ok bool) {
	ok = js.True == bindings.GetWindowNavigator(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientInformation returns the value of property "Window.clientInformation".
//
// It returns ok=false if there is no such property.
func (this Window) ClientInformation() (ret Navigator, ok bool) {
	ok = js.True == bindings.GetWindowClientInformation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OriginAgentCluster returns the value of property "Window.originAgentCluster".
//
// It returns ok=false if there is no such property.
func (this Window) OriginAgentCluster() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowOriginAgentCluster(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "Window.orientation".
//
// It returns ok=false if there is no such property.
func (this Window) Orientation() (ret int16, ok bool) {
	ok = js.True == bindings.GetWindowOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LaunchQueue returns the value of property "Window.launchQueue".
//
// It returns ok=false if there is no such property.
func (this Window) LaunchQueue() (ret LaunchQueue, ok bool) {
	ok = js.True == bindings.GetWindowLaunchQueue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Fence returns the value of property "Window.fence".
//
// It returns ok=false if there is no such property.
func (this Window) Fence() (ret Fence, ok bool) {
	ok = js.True == bindings.GetWindowFence(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PortalHost returns the value of property "Window.portalHost".
//
// It returns ok=false if there is no such property.
func (this Window) PortalHost() (ret PortalHost, ok bool) {
	ok = js.True == bindings.GetWindowPortalHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Screen returns the value of property "Window.screen".
//
// It returns ok=false if there is no such property.
func (this Window) Screen() (ret Screen, ok bool) {
	ok = js.True == bindings.GetWindowScreen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VisualViewport returns the value of property "Window.visualViewport".
//
// It returns ok=false if there is no such property.
func (this Window) VisualViewport() (ret VisualViewport, ok bool) {
	ok = js.True == bindings.GetWindowVisualViewport(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InnerWidth returns the value of property "Window.innerWidth".
//
// It returns ok=false if there is no such property.
func (this Window) InnerWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowInnerWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InnerHeight returns the value of property "Window.innerHeight".
//
// It returns ok=false if there is no such property.
func (this Window) InnerHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowInnerHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScrollX returns the value of property "Window.scrollX".
//
// It returns ok=false if there is no such property.
func (this Window) ScrollX() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowScrollX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageXOffset returns the value of property "Window.pageXOffset".
//
// It returns ok=false if there is no such property.
func (this Window) PageXOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowPageXOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScrollY returns the value of property "Window.scrollY".
//
// It returns ok=false if there is no such property.
func (this Window) ScrollY() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowScrollY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageYOffset returns the value of property "Window.pageYOffset".
//
// It returns ok=false if there is no such property.
func (this Window) PageYOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowPageYOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenX returns the value of property "Window.screenX".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenX() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenLeft returns the value of property "Window.screenLeft".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenLeft() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenY returns the value of property "Window.screenY".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenY() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenTop returns the value of property "Window.screenTop".
//
// It returns ok=false if there is no such property.
func (this Window) ScreenTop() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowScreenTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OuterWidth returns the value of property "Window.outerWidth".
//
// It returns ok=false if there is no such property.
func (this Window) OuterWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowOuterWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OuterHeight returns the value of property "Window.outerHeight".
//
// It returns ok=false if there is no such property.
func (this Window) OuterHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetWindowOuterHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DevicePixelRatio returns the value of property "Window.devicePixelRatio".
//
// It returns ok=false if there is no such property.
func (this Window) DevicePixelRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetWindowDevicePixelRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SharedStorage returns the value of property "Window.sharedStorage".
//
// It returns ok=false if there is no such property.
func (this Window) SharedStorage() (ret WindowSharedStorage, ok bool) {
	ok = js.True == bindings.GetWindowSharedStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Event returns the value of property "Window.event".
//
// It returns ok=false if there is no such property.
func (this Window) Event() (ret OneOf_Event_undefined, ok bool) {
	ok = js.True == bindings.GetWindowEvent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CookieStore returns the value of property "Window.cookieStore".
//
// It returns ok=false if there is no such property.
func (this Window) CookieStore() (ret CookieStore, ok bool) {
	ok = js.True == bindings.GetWindowCookieStore(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DocumentPictureInPicture returns the value of property "Window.documentPictureInPicture".
//
// It returns ok=false if there is no such property.
func (this Window) DocumentPictureInPicture() (ret DocumentPictureInPicture, ok bool) {
	ok = js.True == bindings.GetWindowDocumentPictureInPicture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// External returns the value of property "Window.external".
//
// It returns ok=false if there is no such property.
func (this Window) External() (ret External, ok bool) {
	ok = js.True == bindings.GetWindowExternal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SpeechSynthesis returns the value of property "Window.speechSynthesis".
//
// It returns ok=false if there is no such property.
func (this Window) SpeechSynthesis() (ret SpeechSynthesis, ok bool) {
	ok = js.True == bindings.GetWindowSpeechSynthesis(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "Window.origin".
//
// It returns ok=false if there is no such property.
func (this Window) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetWindowOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsSecureContext returns the value of property "Window.isSecureContext".
//
// It returns ok=false if there is no such property.
func (this Window) IsSecureContext() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowIsSecureContext(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CrossOriginIsolated returns the value of property "Window.crossOriginIsolated".
//
// It returns ok=false if there is no such property.
func (this Window) CrossOriginIsolated() (ret bool, ok bool) {
	ok = js.True == bindings.GetWindowCrossOriginIsolated(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scheduler returns the value of property "Window.scheduler".
//
// It returns ok=false if there is no such property.
func (this Window) Scheduler() (ret Scheduler, ok bool) {
	ok = js.True == bindings.GetWindowScheduler(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TrustedTypes returns the value of property "Window.trustedTypes".
//
// It returns ok=false if there is no such property.
func (this Window) TrustedTypes() (ret TrustedTypePolicyFactory, ok bool) {
	ok = js.True == bindings.GetWindowTrustedTypes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Caches returns the value of property "Window.caches".
//
// It returns ok=false if there is no such property.
func (this Window) Caches() (ret CacheStorage, ok bool) {
	ok = js.True == bindings.GetWindowCaches(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Crypto returns the value of property "Window.crypto".
//
// It returns ok=false if there is no such property.
func (this Window) Crypto() (ret Crypto, ok bool) {
	ok = js.True == bindings.GetWindowCrypto(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IndexedDB returns the value of property "Window.indexedDB".
//
// It returns ok=false if there is no such property.
func (this Window) IndexedDB() (ret IDBFactory, ok bool) {
	ok = js.True == bindings.GetWindowIndexedDB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Performance returns the value of property "Window.performance".
//
// It returns ok=false if there is no such property.
func (this Window) Performance() (ret Performance, ok bool) {
	ok = js.True == bindings.GetWindowPerformance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SessionStorage returns the value of property "Window.sessionStorage".
//
// It returns ok=false if there is no such property.
func (this Window) SessionStorage() (ret Storage, ok bool) {
	ok = js.True == bindings.GetWindowSessionStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LocalStorage returns the value of property "Window.localStorage".
//
// It returns ok=false if there is no such property.
func (this Window) LocalStorage() (ret Storage, ok bool) {
	ok = js.True == bindings.GetWindowLocalStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClose returns true if the method "Window.close" exists.
func (this Window) HasFuncClose() bool {
	return js.True == bindings.HasFuncWindowClose(
		this.ref,
	)
}

// FuncClose returns the method "Window.close".
func (this Window) FuncClose() (fn js.Func[func()]) {
	bindings.FuncWindowClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "Window.close".
func (this Window) Close() (ret js.Void) {
	bindings.CallWindowClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "Window.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "Window.stop" exists.
func (this Window) HasFuncStop() bool {
	return js.True == bindings.HasFuncWindowStop(
		this.ref,
	)
}

// FuncStop returns the method "Window.stop".
func (this Window) FuncStop() (fn js.Func[func()]) {
	bindings.FuncWindowStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "Window.stop".
func (this Window) Stop() (ret js.Void) {
	bindings.CallWindowStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "Window.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFocus returns true if the method "Window.focus" exists.
func (this Window) HasFuncFocus() bool {
	return js.True == bindings.HasFuncWindowFocus(
		this.ref,
	)
}

// FuncFocus returns the method "Window.focus".
func (this Window) FuncFocus() (fn js.Func[func()]) {
	bindings.FuncWindowFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "Window.focus".
func (this Window) Focus() (ret js.Void) {
	bindings.CallWindowFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "Window.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlur returns true if the method "Window.blur" exists.
func (this Window) HasFuncBlur() bool {
	return js.True == bindings.HasFuncWindowBlur(
		this.ref,
	)
}

// FuncBlur returns the method "Window.blur".
func (this Window) FuncBlur() (fn js.Func[func()]) {
	bindings.FuncWindowBlur(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blur calls the method "Window.blur".
func (this Window) Blur() (ret js.Void) {
	bindings.CallWindowBlur(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "Window.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowBlur(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpen returns true if the method "Window.open" exists.
func (this Window) HasFuncOpen() bool {
	return js.True == bindings.HasFuncWindowOpen(
		this.ref,
	)
}

// FuncOpen returns the method "Window.open".
func (this Window) FuncOpen() (fn js.Func[func(url js.String, target js.String, features js.String) js.Object]) {
	bindings.FuncWindowOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "Window.open".
func (this Window) Open(url js.String, target js.String, features js.String) (ret js.Object) {
	bindings.CallWindowOpen(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		target.Ref(),
		features.Ref(),
	)

	return
}

// TryOpen calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryOpen(url js.String, target js.String, features js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		target.Ref(),
		features.Ref(),
	)

	return
}

// HasFuncOpen1 returns true if the method "Window.open" exists.
func (this Window) HasFuncOpen1() bool {
	return js.True == bindings.HasFuncWindowOpen1(
		this.ref,
	)
}

// FuncOpen1 returns the method "Window.open".
func (this Window) FuncOpen1() (fn js.Func[func(url js.String, target js.String) js.Object]) {
	bindings.FuncWindowOpen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open1 calls the method "Window.open".
func (this Window) Open1(url js.String, target js.String) (ret js.Object) {
	bindings.CallWindowOpen1(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		target.Ref(),
	)

	return
}

// TryOpen1 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryOpen1(url js.String, target js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		target.Ref(),
	)

	return
}

// HasFuncOpen2 returns true if the method "Window.open" exists.
func (this Window) HasFuncOpen2() bool {
	return js.True == bindings.HasFuncWindowOpen2(
		this.ref,
	)
}

// FuncOpen2 returns the method "Window.open".
func (this Window) FuncOpen2() (fn js.Func[func(url js.String) js.Object]) {
	bindings.FuncWindowOpen2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open2 calls the method "Window.open".
func (this Window) Open2(url js.String) (ret js.Object) {
	bindings.CallWindowOpen2(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpen2 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryOpen2(url js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncOpen3 returns true if the method "Window.open" exists.
func (this Window) HasFuncOpen3() bool {
	return js.True == bindings.HasFuncWindowOpen3(
		this.ref,
	)
}

// FuncOpen3 returns the method "Window.open".
func (this Window) FuncOpen3() (fn js.Func[func() js.Object]) {
	bindings.FuncWindowOpen3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open3 calls the method "Window.open".
func (this Window) Open3() (ret js.Object) {
	bindings.CallWindowOpen3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpen3 calls the method "Window.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryOpen3() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowOpen3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGet returns true if the method "Window." exists.
func (this Window) HasFuncGet() bool {
	return js.True == bindings.HasFuncWindowGet(
		this.ref,
	)
}

// FuncGet returns the method "Window.".
func (this Window) FuncGet() (fn js.Func[func(name js.String) js.Object]) {
	bindings.FuncWindowGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "Window.".
func (this Window) Get(name js.String) (ret js.Object) {
	bindings.CallWindowGet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "Window."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGet(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncAlert returns true if the method "Window.alert" exists.
func (this Window) HasFuncAlert() bool {
	return js.True == bindings.HasFuncWindowAlert(
		this.ref,
	)
}

// FuncAlert returns the method "Window.alert".
func (this Window) FuncAlert() (fn js.Func[func()]) {
	bindings.FuncWindowAlert(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Alert calls the method "Window.alert".
func (this Window) Alert() (ret js.Void) {
	bindings.CallWindowAlert(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAlert calls the method "Window.alert"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryAlert() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAlert(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAlert1 returns true if the method "Window.alert" exists.
func (this Window) HasFuncAlert1() bool {
	return js.True == bindings.HasFuncWindowAlert1(
		this.ref,
	)
}

// FuncAlert1 returns the method "Window.alert".
func (this Window) FuncAlert1() (fn js.Func[func(message js.String)]) {
	bindings.FuncWindowAlert1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Alert1 calls the method "Window.alert".
func (this Window) Alert1(message js.String) (ret js.Void) {
	bindings.CallWindowAlert1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryAlert1 calls the method "Window.alert"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryAlert1(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAlert1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncConfirm returns true if the method "Window.confirm" exists.
func (this Window) HasFuncConfirm() bool {
	return js.True == bindings.HasFuncWindowConfirm(
		this.ref,
	)
}

// FuncConfirm returns the method "Window.confirm".
func (this Window) FuncConfirm() (fn js.Func[func(message js.String) bool]) {
	bindings.FuncWindowConfirm(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Confirm calls the method "Window.confirm".
func (this Window) Confirm(message js.String) (ret bool) {
	bindings.CallWindowConfirm(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryConfirm calls the method "Window.confirm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryConfirm(message js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowConfirm(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncConfirm1 returns true if the method "Window.confirm" exists.
func (this Window) HasFuncConfirm1() bool {
	return js.True == bindings.HasFuncWindowConfirm1(
		this.ref,
	)
}

// FuncConfirm1 returns the method "Window.confirm".
func (this Window) FuncConfirm1() (fn js.Func[func() bool]) {
	bindings.FuncWindowConfirm1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Confirm1 calls the method "Window.confirm".
func (this Window) Confirm1() (ret bool) {
	bindings.CallWindowConfirm1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryConfirm1 calls the method "Window.confirm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryConfirm1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowConfirm1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPrompt returns true if the method "Window.prompt" exists.
func (this Window) HasFuncPrompt() bool {
	return js.True == bindings.HasFuncWindowPrompt(
		this.ref,
	)
}

// FuncPrompt returns the method "Window.prompt".
func (this Window) FuncPrompt() (fn js.Func[func(message js.String, def js.String) js.String]) {
	bindings.FuncWindowPrompt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prompt calls the method "Window.prompt".
func (this Window) Prompt(message js.String, def js.String) (ret js.String) {
	bindings.CallWindowPrompt(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		def.Ref(),
	)

	return
}

// TryPrompt calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPrompt(message js.String, def js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		def.Ref(),
	)

	return
}

// HasFuncPrompt1 returns true if the method "Window.prompt" exists.
func (this Window) HasFuncPrompt1() bool {
	return js.True == bindings.HasFuncWindowPrompt1(
		this.ref,
	)
}

// FuncPrompt1 returns the method "Window.prompt".
func (this Window) FuncPrompt1() (fn js.Func[func(message js.String) js.String]) {
	bindings.FuncWindowPrompt1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prompt1 calls the method "Window.prompt".
func (this Window) Prompt1(message js.String) (ret js.String) {
	bindings.CallWindowPrompt1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPrompt1 calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPrompt1(message js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncPrompt2 returns true if the method "Window.prompt" exists.
func (this Window) HasFuncPrompt2() bool {
	return js.True == bindings.HasFuncWindowPrompt2(
		this.ref,
	)
}

// FuncPrompt2 returns the method "Window.prompt".
func (this Window) FuncPrompt2() (fn js.Func[func() js.String]) {
	bindings.FuncWindowPrompt2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prompt2 calls the method "Window.prompt".
func (this Window) Prompt2() (ret js.String) {
	bindings.CallWindowPrompt2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPrompt2 calls the method "Window.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPrompt2() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrompt2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPrint returns true if the method "Window.print" exists.
func (this Window) HasFuncPrint() bool {
	return js.True == bindings.HasFuncWindowPrint(
		this.ref,
	)
}

// FuncPrint returns the method "Window.print".
func (this Window) FuncPrint() (fn js.Func[func()]) {
	bindings.FuncWindowPrint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Print calls the method "Window.print".
func (this Window) Print() (ret js.Void) {
	bindings.CallWindowPrint(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPrint calls the method "Window.print"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPrint() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPrint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPostMessage returns true if the method "Window.postMessage" exists.
func (this Window) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncWindowPostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "Window.postMessage".
func (this Window) FuncPostMessage() (fn js.Func[func(message js.Any, targetOrigin js.String, transfer js.Array[js.Object])]) {
	bindings.FuncWindowPostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "Window.postMessage".
func (this Window) PostMessage(message js.Any, targetOrigin js.String, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallWindowPostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		targetOrigin.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPostMessage(message js.Any, targetOrigin js.String, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		targetOrigin.Ref(),
		transfer.Ref(),
	)

	return
}

// HasFuncPostMessage1 returns true if the method "Window.postMessage" exists.
func (this Window) HasFuncPostMessage1() bool {
	return js.True == bindings.HasFuncWindowPostMessage1(
		this.ref,
	)
}

// FuncPostMessage1 returns the method "Window.postMessage".
func (this Window) FuncPostMessage1() (fn js.Func[func(message js.Any, targetOrigin js.String)]) {
	bindings.FuncWindowPostMessage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage1 calls the method "Window.postMessage".
func (this Window) PostMessage1(message js.Any, targetOrigin js.String) (ret js.Void) {
	bindings.CallWindowPostMessage1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPostMessage1(message js.Any, targetOrigin js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

// HasFuncPostMessage2 returns true if the method "Window.postMessage" exists.
func (this Window) HasFuncPostMessage2() bool {
	return js.True == bindings.HasFuncWindowPostMessage2(
		this.ref,
	)
}

// FuncPostMessage2 returns the method "Window.postMessage".
func (this Window) FuncPostMessage2() (fn js.Func[func(message js.Any, options WindowPostMessageOptions)]) {
	bindings.FuncWindowPostMessage2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage2 calls the method "Window.postMessage".
func (this Window) PostMessage2(message js.Any, options WindowPostMessageOptions) (ret js.Void) {
	bindings.CallWindowPostMessage2(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage2 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPostMessage2(message js.Any, options WindowPostMessageOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPostMessage3 returns true if the method "Window.postMessage" exists.
func (this Window) HasFuncPostMessage3() bool {
	return js.True == bindings.HasFuncWindowPostMessage3(
		this.ref,
	)
}

// FuncPostMessage3 returns the method "Window.postMessage".
func (this Window) FuncPostMessage3() (fn js.Func[func(message js.Any)]) {
	bindings.FuncWindowPostMessage3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage3 calls the method "Window.postMessage".
func (this Window) PostMessage3(message js.Any) (ret js.Void) {
	bindings.CallWindowPostMessage3(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage3 calls the method "Window.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryPostMessage3(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowPostMessage3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncGetSelection returns true if the method "Window.getSelection" exists.
func (this Window) HasFuncGetSelection() bool {
	return js.True == bindings.HasFuncWindowGetSelection(
		this.ref,
	)
}

// FuncGetSelection returns the method "Window.getSelection".
func (this Window) FuncGetSelection() (fn js.Func[func() Selection]) {
	bindings.FuncWindowGetSelection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSelection calls the method "Window.getSelection".
func (this Window) GetSelection() (ret Selection) {
	bindings.CallWindowGetSelection(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSelection calls the method "Window.getSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGetSelection() (ret Selection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetSelection(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNavigate returns true if the method "Window.navigate" exists.
func (this Window) HasFuncNavigate() bool {
	return js.True == bindings.HasFuncWindowNavigate(
		this.ref,
	)
}

// FuncNavigate returns the method "Window.navigate".
func (this Window) FuncNavigate() (fn js.Func[func(dir SpatialNavigationDirection)]) {
	bindings.FuncWindowNavigate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Navigate calls the method "Window.navigate".
func (this Window) Navigate(dir SpatialNavigationDirection) (ret js.Void) {
	bindings.CallWindowNavigate(
		this.ref, js.Pointer(&ret),
		uint32(dir),
	)

	return
}

// TryNavigate calls the method "Window.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryNavigate(dir SpatialNavigationDirection) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowNavigate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(dir),
	)

	return
}

// HasFuncGetComputedStyle returns true if the method "Window.getComputedStyle" exists.
func (this Window) HasFuncGetComputedStyle() bool {
	return js.True == bindings.HasFuncWindowGetComputedStyle(
		this.ref,
	)
}

// FuncGetComputedStyle returns the method "Window.getComputedStyle".
func (this Window) FuncGetComputedStyle() (fn js.Func[func(elt Element, pseudoElt js.String) CSSStyleDeclaration]) {
	bindings.FuncWindowGetComputedStyle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetComputedStyle calls the method "Window.getComputedStyle".
func (this Window) GetComputedStyle(elt Element, pseudoElt js.String) (ret CSSStyleDeclaration) {
	bindings.CallWindowGetComputedStyle(
		this.ref, js.Pointer(&ret),
		elt.Ref(),
		pseudoElt.Ref(),
	)

	return
}

// TryGetComputedStyle calls the method "Window.getComputedStyle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGetComputedStyle(elt Element, pseudoElt js.String) (ret CSSStyleDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetComputedStyle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elt.Ref(),
		pseudoElt.Ref(),
	)

	return
}

// HasFuncGetComputedStyle1 returns true if the method "Window.getComputedStyle" exists.
func (this Window) HasFuncGetComputedStyle1() bool {
	return js.True == bindings.HasFuncWindowGetComputedStyle1(
		this.ref,
	)
}

// FuncGetComputedStyle1 returns the method "Window.getComputedStyle".
func (this Window) FuncGetComputedStyle1() (fn js.Func[func(elt Element) CSSStyleDeclaration]) {
	bindings.FuncWindowGetComputedStyle1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetComputedStyle1 calls the method "Window.getComputedStyle".
func (this Window) GetComputedStyle1(elt Element) (ret CSSStyleDeclaration) {
	bindings.CallWindowGetComputedStyle1(
		this.ref, js.Pointer(&ret),
		elt.Ref(),
	)

	return
}

// TryGetComputedStyle1 calls the method "Window.getComputedStyle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGetComputedStyle1(elt Element) (ret CSSStyleDeclaration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetComputedStyle1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elt.Ref(),
	)

	return
}

// HasFuncMatchMedia returns true if the method "Window.matchMedia" exists.
func (this Window) HasFuncMatchMedia() bool {
	return js.True == bindings.HasFuncWindowMatchMedia(
		this.ref,
	)
}

// FuncMatchMedia returns the method "Window.matchMedia".
func (this Window) FuncMatchMedia() (fn js.Func[func(query js.String) MediaQueryList]) {
	bindings.FuncWindowMatchMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchMedia calls the method "Window.matchMedia".
func (this Window) MatchMedia(query js.String) (ret MediaQueryList) {
	bindings.CallWindowMatchMedia(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryMatchMedia calls the method "Window.matchMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryMatchMedia(query js.String) (ret MediaQueryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMatchMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncMoveTo returns true if the method "Window.moveTo" exists.
func (this Window) HasFuncMoveTo() bool {
	return js.True == bindings.HasFuncWindowMoveTo(
		this.ref,
	)
}

// FuncMoveTo returns the method "Window.moveTo".
func (this Window) FuncMoveTo() (fn js.Func[func(x int32, y int32)]) {
	bindings.FuncWindowMoveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveTo calls the method "Window.moveTo".
func (this Window) MoveTo(x int32, y int32) (ret js.Void) {
	bindings.CallWindowMoveTo(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryMoveTo calls the method "Window.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryMoveTo(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMoveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncMoveBy returns true if the method "Window.moveBy" exists.
func (this Window) HasFuncMoveBy() bool {
	return js.True == bindings.HasFuncWindowMoveBy(
		this.ref,
	)
}

// FuncMoveBy returns the method "Window.moveBy".
func (this Window) FuncMoveBy() (fn js.Func[func(x int32, y int32)]) {
	bindings.FuncWindowMoveBy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveBy calls the method "Window.moveBy".
func (this Window) MoveBy(x int32, y int32) (ret js.Void) {
	bindings.CallWindowMoveBy(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryMoveBy calls the method "Window.moveBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryMoveBy(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowMoveBy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncResizeTo returns true if the method "Window.resizeTo" exists.
func (this Window) HasFuncResizeTo() bool {
	return js.True == bindings.HasFuncWindowResizeTo(
		this.ref,
	)
}

// FuncResizeTo returns the method "Window.resizeTo".
func (this Window) FuncResizeTo() (fn js.Func[func(width int32, height int32)]) {
	bindings.FuncWindowResizeTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResizeTo calls the method "Window.resizeTo".
func (this Window) ResizeTo(width int32, height int32) (ret js.Void) {
	bindings.CallWindowResizeTo(
		this.ref, js.Pointer(&ret),
		int32(width),
		int32(height),
	)

	return
}

// TryResizeTo calls the method "Window.resizeTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryResizeTo(width int32, height int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowResizeTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncResizeBy returns true if the method "Window.resizeBy" exists.
func (this Window) HasFuncResizeBy() bool {
	return js.True == bindings.HasFuncWindowResizeBy(
		this.ref,
	)
}

// FuncResizeBy returns the method "Window.resizeBy".
func (this Window) FuncResizeBy() (fn js.Func[func(x int32, y int32)]) {
	bindings.FuncWindowResizeBy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResizeBy calls the method "Window.resizeBy".
func (this Window) ResizeBy(x int32, y int32) (ret js.Void) {
	bindings.CallWindowResizeBy(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryResizeBy calls the method "Window.resizeBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryResizeBy(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowResizeBy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncScroll returns true if the method "Window.scroll" exists.
func (this Window) HasFuncScroll() bool {
	return js.True == bindings.HasFuncWindowScroll(
		this.ref,
	)
}

// FuncScroll returns the method "Window.scroll".
func (this Window) FuncScroll() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncWindowScroll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll calls the method "Window.scroll".
func (this Window) Scroll(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScroll(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScroll calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScroll(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScroll1 returns true if the method "Window.scroll" exists.
func (this Window) HasFuncScroll1() bool {
	return js.True == bindings.HasFuncWindowScroll1(
		this.ref,
	)
}

// FuncScroll1 returns the method "Window.scroll".
func (this Window) FuncScroll1() (fn js.Func[func()]) {
	bindings.FuncWindowScroll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll1 calls the method "Window.scroll".
func (this Window) Scroll1() (ret js.Void) {
	bindings.CallWindowScroll1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScroll1 calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScroll1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScroll2 returns true if the method "Window.scroll" exists.
func (this Window) HasFuncScroll2() bool {
	return js.True == bindings.HasFuncWindowScroll2(
		this.ref,
	)
}

// FuncScroll2 returns the method "Window.scroll".
func (this Window) FuncScroll2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncWindowScroll2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scroll2 calls the method "Window.scroll".
func (this Window) Scroll2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScroll2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScroll2 calls the method "Window.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScroll2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScroll2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncScrollTo returns true if the method "Window.scrollTo" exists.
func (this Window) HasFuncScrollTo() bool {
	return js.True == bindings.HasFuncWindowScrollTo(
		this.ref,
	)
}

// FuncScrollTo returns the method "Window.scrollTo".
func (this Window) FuncScrollTo() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncWindowScrollTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo calls the method "Window.scrollTo".
func (this Window) ScrollTo(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScrollTo(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollTo calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollTo(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScrollTo1 returns true if the method "Window.scrollTo" exists.
func (this Window) HasFuncScrollTo1() bool {
	return js.True == bindings.HasFuncWindowScrollTo1(
		this.ref,
	)
}

// FuncScrollTo1 returns the method "Window.scrollTo".
func (this Window) FuncScrollTo1() (fn js.Func[func()]) {
	bindings.FuncWindowScrollTo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo1 calls the method "Window.scrollTo".
func (this Window) ScrollTo1() (ret js.Void) {
	bindings.CallWindowScrollTo1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollTo1 calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollTo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollTo2 returns true if the method "Window.scrollTo" exists.
func (this Window) HasFuncScrollTo2() bool {
	return js.True == bindings.HasFuncWindowScrollTo2(
		this.ref,
	)
}

// FuncScrollTo2 returns the method "Window.scrollTo".
func (this Window) FuncScrollTo2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncWindowScrollTo2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollTo2 calls the method "Window.scrollTo".
func (this Window) ScrollTo2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScrollTo2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollTo2 calls the method "Window.scrollTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollTo2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollTo2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncScrollBy returns true if the method "Window.scrollBy" exists.
func (this Window) HasFuncScrollBy() bool {
	return js.True == bindings.HasFuncWindowScrollBy(
		this.ref,
	)
}

// FuncScrollBy returns the method "Window.scrollBy".
func (this Window) FuncScrollBy() (fn js.Func[func(options ScrollToOptions)]) {
	bindings.FuncWindowScrollBy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy calls the method "Window.scrollBy".
func (this Window) ScrollBy(options ScrollToOptions) (ret js.Void) {
	bindings.CallWindowScrollBy(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScrollBy calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollBy(options ScrollToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScrollBy1 returns true if the method "Window.scrollBy" exists.
func (this Window) HasFuncScrollBy1() bool {
	return js.True == bindings.HasFuncWindowScrollBy1(
		this.ref,
	)
}

// FuncScrollBy1 returns the method "Window.scrollBy".
func (this Window) FuncScrollBy1() (fn js.Func[func()]) {
	bindings.FuncWindowScrollBy1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy1 calls the method "Window.scrollBy".
func (this Window) ScrollBy1() (ret js.Void) {
	bindings.CallWindowScrollBy1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollBy1 calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollBy1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollBy2 returns true if the method "Window.scrollBy" exists.
func (this Window) HasFuncScrollBy2() bool {
	return js.True == bindings.HasFuncWindowScrollBy2(
		this.ref,
	)
}

// FuncScrollBy2 returns the method "Window.scrollBy".
func (this Window) FuncScrollBy2() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncWindowScrollBy2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBy2 calls the method "Window.scrollBy".
func (this Window) ScrollBy2(x float64, y float64) (ret js.Void) {
	bindings.CallWindowScrollBy2(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScrollBy2 calls the method "Window.scrollBy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryScrollBy2(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowScrollBy2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRequestIdleCallback returns true if the method "Window.requestIdleCallback" exists.
func (this Window) HasFuncRequestIdleCallback() bool {
	return js.True == bindings.HasFuncWindowRequestIdleCallback(
		this.ref,
	)
}

// FuncRequestIdleCallback returns the method "Window.requestIdleCallback".
func (this Window) FuncRequestIdleCallback() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) uint32]) {
	bindings.FuncWindowRequestIdleCallback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestIdleCallback calls the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallback(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) (ret uint32) {
	bindings.CallWindowRequestIdleCallback(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRequestIdleCallback calls the method "Window.requestIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryRequestIdleCallback(callback js.Func[func(deadline IdleDeadline)], options IdleRequestOptions) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestIdleCallback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestIdleCallback1 returns true if the method "Window.requestIdleCallback" exists.
func (this Window) HasFuncRequestIdleCallback1() bool {
	return js.True == bindings.HasFuncWindowRequestIdleCallback1(
		this.ref,
	)
}

// FuncRequestIdleCallback1 returns the method "Window.requestIdleCallback".
func (this Window) FuncRequestIdleCallback1() (fn js.Func[func(callback js.Func[func(deadline IdleDeadline)]) uint32]) {
	bindings.FuncWindowRequestIdleCallback1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestIdleCallback1 calls the method "Window.requestIdleCallback".
func (this Window) RequestIdleCallback1(callback js.Func[func(deadline IdleDeadline)]) (ret uint32) {
	bindings.CallWindowRequestIdleCallback1(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestIdleCallback1 calls the method "Window.requestIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryRequestIdleCallback1(callback js.Func[func(deadline IdleDeadline)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestIdleCallback1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCancelIdleCallback returns true if the method "Window.cancelIdleCallback" exists.
func (this Window) HasFuncCancelIdleCallback() bool {
	return js.True == bindings.HasFuncWindowCancelIdleCallback(
		this.ref,
	)
}

// FuncCancelIdleCallback returns the method "Window.cancelIdleCallback".
func (this Window) FuncCancelIdleCallback() (fn js.Func[func(handle uint32)]) {
	bindings.FuncWindowCancelIdleCallback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelIdleCallback calls the method "Window.cancelIdleCallback".
func (this Window) CancelIdleCallback(handle uint32) (ret js.Void) {
	bindings.CallWindowCancelIdleCallback(
		this.ref, js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelIdleCallback calls the method "Window.cancelIdleCallback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCancelIdleCallback(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCancelIdleCallback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

// HasFuncShowOpenFilePicker returns true if the method "Window.showOpenFilePicker" exists.
func (this Window) HasFuncShowOpenFilePicker() bool {
	return js.True == bindings.HasFuncWindowShowOpenFilePicker(
		this.ref,
	)
}

// FuncShowOpenFilePicker returns the method "Window.showOpenFilePicker".
func (this Window) FuncShowOpenFilePicker() (fn js.Func[func(options OpenFilePickerOptions) js.Promise[js.Array[FileSystemFileHandle]]]) {
	bindings.FuncWindowShowOpenFilePicker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowOpenFilePicker calls the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePicker(options OpenFilePickerOptions) (ret js.Promise[js.Array[FileSystemFileHandle]]) {
	bindings.CallWindowShowOpenFilePicker(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowOpenFilePicker calls the method "Window.showOpenFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowOpenFilePicker(options OpenFilePickerOptions) (ret js.Promise[js.Array[FileSystemFileHandle]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowOpenFilePicker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncShowOpenFilePicker1 returns true if the method "Window.showOpenFilePicker" exists.
func (this Window) HasFuncShowOpenFilePicker1() bool {
	return js.True == bindings.HasFuncWindowShowOpenFilePicker1(
		this.ref,
	)
}

// FuncShowOpenFilePicker1 returns the method "Window.showOpenFilePicker".
func (this Window) FuncShowOpenFilePicker1() (fn js.Func[func() js.Promise[js.Array[FileSystemFileHandle]]]) {
	bindings.FuncWindowShowOpenFilePicker1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowOpenFilePicker1 calls the method "Window.showOpenFilePicker".
func (this Window) ShowOpenFilePicker1() (ret js.Promise[js.Array[FileSystemFileHandle]]) {
	bindings.CallWindowShowOpenFilePicker1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowOpenFilePicker1 calls the method "Window.showOpenFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowOpenFilePicker1() (ret js.Promise[js.Array[FileSystemFileHandle]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowOpenFilePicker1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowSaveFilePicker returns true if the method "Window.showSaveFilePicker" exists.
func (this Window) HasFuncShowSaveFilePicker() bool {
	return js.True == bindings.HasFuncWindowShowSaveFilePicker(
		this.ref,
	)
}

// FuncShowSaveFilePicker returns the method "Window.showSaveFilePicker".
func (this Window) FuncShowSaveFilePicker() (fn js.Func[func(options SaveFilePickerOptions) js.Promise[FileSystemFileHandle]]) {
	bindings.FuncWindowShowSaveFilePicker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowSaveFilePicker calls the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePicker(options SaveFilePickerOptions) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallWindowShowSaveFilePicker(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowSaveFilePicker calls the method "Window.showSaveFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowSaveFilePicker(options SaveFilePickerOptions) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowSaveFilePicker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncShowSaveFilePicker1 returns true if the method "Window.showSaveFilePicker" exists.
func (this Window) HasFuncShowSaveFilePicker1() bool {
	return js.True == bindings.HasFuncWindowShowSaveFilePicker1(
		this.ref,
	)
}

// FuncShowSaveFilePicker1 returns the method "Window.showSaveFilePicker".
func (this Window) FuncShowSaveFilePicker1() (fn js.Func[func() js.Promise[FileSystemFileHandle]]) {
	bindings.FuncWindowShowSaveFilePicker1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowSaveFilePicker1 calls the method "Window.showSaveFilePicker".
func (this Window) ShowSaveFilePicker1() (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallWindowShowSaveFilePicker1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowSaveFilePicker1 calls the method "Window.showSaveFilePicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowSaveFilePicker1() (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowSaveFilePicker1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowDirectoryPicker returns true if the method "Window.showDirectoryPicker" exists.
func (this Window) HasFuncShowDirectoryPicker() bool {
	return js.True == bindings.HasFuncWindowShowDirectoryPicker(
		this.ref,
	)
}

// FuncShowDirectoryPicker returns the method "Window.showDirectoryPicker".
func (this Window) FuncShowDirectoryPicker() (fn js.Func[func(options DirectoryPickerOptions) js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncWindowShowDirectoryPicker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowDirectoryPicker calls the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPicker(options DirectoryPickerOptions) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallWindowShowDirectoryPicker(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryShowDirectoryPicker calls the method "Window.showDirectoryPicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowDirectoryPicker(options DirectoryPickerOptions) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowDirectoryPicker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncShowDirectoryPicker1 returns true if the method "Window.showDirectoryPicker" exists.
func (this Window) HasFuncShowDirectoryPicker1() bool {
	return js.True == bindings.HasFuncWindowShowDirectoryPicker1(
		this.ref,
	)
}

// FuncShowDirectoryPicker1 returns the method "Window.showDirectoryPicker".
func (this Window) FuncShowDirectoryPicker1() (fn js.Func[func() js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncWindowShowDirectoryPicker1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowDirectoryPicker1 calls the method "Window.showDirectoryPicker".
func (this Window) ShowDirectoryPicker1() (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallWindowShowDirectoryPicker1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowDirectoryPicker1 calls the method "Window.showDirectoryPicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryShowDirectoryPicker1() (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowShowDirectoryPicker1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncQueryLocalFonts returns true if the method "Window.queryLocalFonts" exists.
func (this Window) HasFuncQueryLocalFonts() bool {
	return js.True == bindings.HasFuncWindowQueryLocalFonts(
		this.ref,
	)
}

// FuncQueryLocalFonts returns the method "Window.queryLocalFonts".
func (this Window) FuncQueryLocalFonts() (fn js.Func[func(options QueryOptions) js.Promise[js.Array[FontData]]]) {
	bindings.FuncWindowQueryLocalFonts(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryLocalFonts calls the method "Window.queryLocalFonts".
func (this Window) QueryLocalFonts(options QueryOptions) (ret js.Promise[js.Array[FontData]]) {
	bindings.CallWindowQueryLocalFonts(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryQueryLocalFonts calls the method "Window.queryLocalFonts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryQueryLocalFonts(options QueryOptions) (ret js.Promise[js.Array[FontData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueryLocalFonts(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncQueryLocalFonts1 returns true if the method "Window.queryLocalFonts" exists.
func (this Window) HasFuncQueryLocalFonts1() bool {
	return js.True == bindings.HasFuncWindowQueryLocalFonts1(
		this.ref,
	)
}

// FuncQueryLocalFonts1 returns the method "Window.queryLocalFonts".
func (this Window) FuncQueryLocalFonts1() (fn js.Func[func() js.Promise[js.Array[FontData]]]) {
	bindings.FuncWindowQueryLocalFonts1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryLocalFonts1 calls the method "Window.queryLocalFonts".
func (this Window) QueryLocalFonts1() (ret js.Promise[js.Array[FontData]]) {
	bindings.CallWindowQueryLocalFonts1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryQueryLocalFonts1 calls the method "Window.queryLocalFonts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryQueryLocalFonts1() (ret js.Promise[js.Array[FontData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueryLocalFonts1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetScreenDetails returns true if the method "Window.getScreenDetails" exists.
func (this Window) HasFuncGetScreenDetails() bool {
	return js.True == bindings.HasFuncWindowGetScreenDetails(
		this.ref,
	)
}

// FuncGetScreenDetails returns the method "Window.getScreenDetails".
func (this Window) FuncGetScreenDetails() (fn js.Func[func() js.Promise[ScreenDetails]]) {
	bindings.FuncWindowGetScreenDetails(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetScreenDetails calls the method "Window.getScreenDetails".
func (this Window) GetScreenDetails() (ret js.Promise[ScreenDetails]) {
	bindings.CallWindowGetScreenDetails(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetScreenDetails calls the method "Window.getScreenDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGetScreenDetails() (ret js.Promise[ScreenDetails], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetScreenDetails(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDigitalGoodsService returns true if the method "Window.getDigitalGoodsService" exists.
func (this Window) HasFuncGetDigitalGoodsService() bool {
	return js.True == bindings.HasFuncWindowGetDigitalGoodsService(
		this.ref,
	)
}

// FuncGetDigitalGoodsService returns the method "Window.getDigitalGoodsService".
func (this Window) FuncGetDigitalGoodsService() (fn js.Func[func(serviceProvider js.String) js.Promise[DigitalGoodsService]]) {
	bindings.FuncWindowGetDigitalGoodsService(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDigitalGoodsService calls the method "Window.getDigitalGoodsService".
func (this Window) GetDigitalGoodsService(serviceProvider js.String) (ret js.Promise[DigitalGoodsService]) {
	bindings.CallWindowGetDigitalGoodsService(
		this.ref, js.Pointer(&ret),
		serviceProvider.Ref(),
	)

	return
}

// TryGetDigitalGoodsService calls the method "Window.getDigitalGoodsService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryGetDigitalGoodsService(serviceProvider js.String) (ret js.Promise[DigitalGoodsService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowGetDigitalGoodsService(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		serviceProvider.Ref(),
	)

	return
}

// HasFuncCaptureEvents returns true if the method "Window.captureEvents" exists.
func (this Window) HasFuncCaptureEvents() bool {
	return js.True == bindings.HasFuncWindowCaptureEvents(
		this.ref,
	)
}

// FuncCaptureEvents returns the method "Window.captureEvents".
func (this Window) FuncCaptureEvents() (fn js.Func[func()]) {
	bindings.FuncWindowCaptureEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CaptureEvents calls the method "Window.captureEvents".
func (this Window) CaptureEvents() (ret js.Void) {
	bindings.CallWindowCaptureEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCaptureEvents calls the method "Window.captureEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCaptureEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCaptureEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReleaseEvents returns true if the method "Window.releaseEvents" exists.
func (this Window) HasFuncReleaseEvents() bool {
	return js.True == bindings.HasFuncWindowReleaseEvents(
		this.ref,
	)
}

// FuncReleaseEvents returns the method "Window.releaseEvents".
func (this Window) FuncReleaseEvents() (fn js.Func[func()]) {
	bindings.FuncWindowReleaseEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseEvents calls the method "Window.releaseEvents".
func (this Window) ReleaseEvents() (ret js.Void) {
	bindings.CallWindowReleaseEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReleaseEvents calls the method "Window.releaseEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryReleaseEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowReleaseEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportError returns true if the method "Window.reportError" exists.
func (this Window) HasFuncReportError() bool {
	return js.True == bindings.HasFuncWindowReportError(
		this.ref,
	)
}

// FuncReportError returns the method "Window.reportError".
func (this Window) FuncReportError() (fn js.Func[func(e js.Any)]) {
	bindings.FuncWindowReportError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportError calls the method "Window.reportError".
func (this Window) ReportError(e js.Any) (ret js.Void) {
	bindings.CallWindowReportError(
		this.ref, js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryReportError calls the method "Window.reportError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryReportError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowReportError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasFuncBtoa returns true if the method "Window.btoa" exists.
func (this Window) HasFuncBtoa() bool {
	return js.True == bindings.HasFuncWindowBtoa(
		this.ref,
	)
}

// FuncBtoa returns the method "Window.btoa".
func (this Window) FuncBtoa() (fn js.Func[func(data js.String) js.String]) {
	bindings.FuncWindowBtoa(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Btoa calls the method "Window.btoa".
func (this Window) Btoa(data js.String) (ret js.String) {
	bindings.CallWindowBtoa(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryBtoa calls the method "Window.btoa"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryBtoa(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowBtoa(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncAtob returns true if the method "Window.atob" exists.
func (this Window) HasFuncAtob() bool {
	return js.True == bindings.HasFuncWindowAtob(
		this.ref,
	)
}

// FuncAtob returns the method "Window.atob".
func (this Window) FuncAtob() (fn js.Func[func(data js.String) js.String]) {
	bindings.FuncWindowAtob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Atob calls the method "Window.atob".
func (this Window) Atob(data js.String) (ret js.String) {
	bindings.CallWindowAtob(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAtob calls the method "Window.atob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryAtob(data js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowAtob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSetTimeout returns true if the method "Window.setTimeout" exists.
func (this Window) HasFuncSetTimeout() bool {
	return js.True == bindings.HasFuncWindowSetTimeout(
		this.ref,
	)
}

// FuncSetTimeout returns the method "Window.setTimeout".
func (this Window) FuncSetTimeout() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	bindings.FuncWindowSetTimeout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTimeout calls the method "Window.setTimeout".
func (this Window) SetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWindowSetTimeout(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetTimeout calls the method "Window.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TrySetTimeout(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetTimeout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasFuncSetTimeout1 returns true if the method "Window.setTimeout" exists.
func (this Window) HasFuncSetTimeout1() bool {
	return js.True == bindings.HasFuncWindowSetTimeout1(
		this.ref,
	)
}

// FuncSetTimeout1 returns the method "Window.setTimeout".
func (this Window) FuncSetTimeout1() (fn js.Func[func(handler TimerHandler) int32]) {
	bindings.FuncWindowSetTimeout1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTimeout1 calls the method "Window.setTimeout".
func (this Window) SetTimeout1(handler TimerHandler) (ret int32) {
	bindings.CallWindowSetTimeout1(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetTimeout1 calls the method "Window.setTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TrySetTimeout1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetTimeout1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasFuncClearTimeout returns true if the method "Window.clearTimeout" exists.
func (this Window) HasFuncClearTimeout() bool {
	return js.True == bindings.HasFuncWindowClearTimeout(
		this.ref,
	)
}

// FuncClearTimeout returns the method "Window.clearTimeout".
func (this Window) FuncClearTimeout() (fn js.Func[func(id int32)]) {
	bindings.FuncWindowClearTimeout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearTimeout calls the method "Window.clearTimeout".
func (this Window) ClearTimeout(id int32) (ret js.Void) {
	bindings.CallWindowClearTimeout(
		this.ref, js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearTimeout calls the method "Window.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryClearTimeout(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearTimeout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncClearTimeout1 returns true if the method "Window.clearTimeout" exists.
func (this Window) HasFuncClearTimeout1() bool {
	return js.True == bindings.HasFuncWindowClearTimeout1(
		this.ref,
	)
}

// FuncClearTimeout1 returns the method "Window.clearTimeout".
func (this Window) FuncClearTimeout1() (fn js.Func[func()]) {
	bindings.FuncWindowClearTimeout1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearTimeout1 calls the method "Window.clearTimeout".
func (this Window) ClearTimeout1() (ret js.Void) {
	bindings.CallWindowClearTimeout1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearTimeout1 calls the method "Window.clearTimeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryClearTimeout1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearTimeout1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetInterval returns true if the method "Window.setInterval" exists.
func (this Window) HasFuncSetInterval() bool {
	return js.True == bindings.HasFuncWindowSetInterval(
		this.ref,
	)
}

// FuncSetInterval returns the method "Window.setInterval".
func (this Window) FuncSetInterval() (fn js.Func[func(handler TimerHandler, timeout int32, arguments ...js.Any) int32]) {
	bindings.FuncWindowSetInterval(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetInterval calls the method "Window.setInterval".
func (this Window) SetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32) {
	bindings.CallWindowSetInterval(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// TrySetInterval calls the method "Window.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TrySetInterval(handler TimerHandler, timeout int32, arguments ...js.Any) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetInterval(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
		int32(timeout),
		js.SliceData(arguments),
		js.SizeU(len(arguments)),
	)

	return
}

// HasFuncSetInterval1 returns true if the method "Window.setInterval" exists.
func (this Window) HasFuncSetInterval1() bool {
	return js.True == bindings.HasFuncWindowSetInterval1(
		this.ref,
	)
}

// FuncSetInterval1 returns the method "Window.setInterval".
func (this Window) FuncSetInterval1() (fn js.Func[func(handler TimerHandler) int32]) {
	bindings.FuncWindowSetInterval1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetInterval1 calls the method "Window.setInterval".
func (this Window) SetInterval1(handler TimerHandler) (ret int32) {
	bindings.CallWindowSetInterval1(
		this.ref, js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetInterval1 calls the method "Window.setInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TrySetInterval1(handler TimerHandler) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowSetInterval1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasFuncClearInterval returns true if the method "Window.clearInterval" exists.
func (this Window) HasFuncClearInterval() bool {
	return js.True == bindings.HasFuncWindowClearInterval(
		this.ref,
	)
}

// FuncClearInterval returns the method "Window.clearInterval".
func (this Window) FuncClearInterval() (fn js.Func[func(id int32)]) {
	bindings.FuncWindowClearInterval(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearInterval calls the method "Window.clearInterval".
func (this Window) ClearInterval(id int32) (ret js.Void) {
	bindings.CallWindowClearInterval(
		this.ref, js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryClearInterval calls the method "Window.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryClearInterval(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearInterval(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncClearInterval1 returns true if the method "Window.clearInterval" exists.
func (this Window) HasFuncClearInterval1() bool {
	return js.True == bindings.HasFuncWindowClearInterval1(
		this.ref,
	)
}

// FuncClearInterval1 returns the method "Window.clearInterval".
func (this Window) FuncClearInterval1() (fn js.Func[func()]) {
	bindings.FuncWindowClearInterval1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearInterval1 calls the method "Window.clearInterval".
func (this Window) ClearInterval1() (ret js.Void) {
	bindings.CallWindowClearInterval1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearInterval1 calls the method "Window.clearInterval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryClearInterval1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowClearInterval1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncQueueMicrotask returns true if the method "Window.queueMicrotask" exists.
func (this Window) HasFuncQueueMicrotask() bool {
	return js.True == bindings.HasFuncWindowQueueMicrotask(
		this.ref,
	)
}

// FuncQueueMicrotask returns the method "Window.queueMicrotask".
func (this Window) FuncQueueMicrotask() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncWindowQueueMicrotask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueueMicrotask calls the method "Window.queueMicrotask".
func (this Window) QueueMicrotask(callback js.Func[func()]) (ret js.Void) {
	bindings.CallWindowQueueMicrotask(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryQueueMicrotask calls the method "Window.queueMicrotask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryQueueMicrotask(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowQueueMicrotask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCreateImageBitmap returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasFuncCreateImageBitmap() bool {
	return js.True == bindings.HasFuncWindowCreateImageBitmap(
		this.ref,
	)
}

// FuncCreateImageBitmap returns the method "Window.createImageBitmap".
func (this Window) FuncCreateImageBitmap() (fn js.Func[func(image ImageBitmapSource, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	bindings.FuncWindowCreateImageBitmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap(
		this.ref, js.Pointer(&ret),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateImageBitmap calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCreateImageBitmap(image ImageBitmapSource, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateImageBitmap1 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasFuncCreateImageBitmap1() bool {
	return js.True == bindings.HasFuncWindowCreateImageBitmap1(
		this.ref,
	)
}

// FuncCreateImageBitmap1 returns the method "Window.createImageBitmap".
func (this Window) FuncCreateImageBitmap1() (fn js.Func[func(image ImageBitmapSource) js.Promise[ImageBitmap]]) {
	bindings.FuncWindowCreateImageBitmap1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap1 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap1(
		this.ref, js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryCreateImageBitmap1 calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCreateImageBitmap1(image ImageBitmapSource) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
}

// HasFuncCreateImageBitmap2 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasFuncCreateImageBitmap2() bool {
	return js.True == bindings.HasFuncWindowCreateImageBitmap2(
		this.ref,
	)
}

// FuncCreateImageBitmap2 returns the method "Window.createImageBitmap".
func (this Window) FuncCreateImageBitmap2() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) js.Promise[ImageBitmap]]) {
	bindings.FuncWindowCreateImageBitmap2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap2 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap2(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCreateImageBitmap2(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32, options ImageBitmapOptions) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateImageBitmap3 returns true if the method "Window.createImageBitmap" exists.
func (this Window) HasFuncCreateImageBitmap3() bool {
	return js.True == bindings.HasFuncWindowCreateImageBitmap3(
		this.ref,
	)
}

// FuncCreateImageBitmap3 returns the method "Window.createImageBitmap".
func (this Window) FuncCreateImageBitmap3() (fn js.Func[func(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) js.Promise[ImageBitmap]]) {
	bindings.FuncWindowCreateImageBitmap3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageBitmap3 calls the method "Window.createImageBitmap".
func (this Window) CreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap]) {
	bindings.CallWindowCreateImageBitmap3(
		this.ref, js.Pointer(&ret),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// TryCreateImageBitmap3 calls the method "Window.createImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCreateImageBitmap3(image ImageBitmapSource, sx int32, sy int32, sw int32, sh int32) (ret js.Promise[ImageBitmap], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCreateImageBitmap3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasFuncStructuredClone returns true if the method "Window.structuredClone" exists.
func (this Window) HasFuncStructuredClone() bool {
	return js.True == bindings.HasFuncWindowStructuredClone(
		this.ref,
	)
}

// FuncStructuredClone returns the method "Window.structuredClone".
func (this Window) FuncStructuredClone() (fn js.Func[func(value js.Any, options StructuredSerializeOptions) js.Any]) {
	bindings.FuncWindowStructuredClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StructuredClone calls the method "Window.structuredClone".
func (this Window) StructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any) {
	bindings.CallWindowStructuredClone(
		this.ref, js.Pointer(&ret),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryStructuredClone calls the method "Window.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryStructuredClone(value js.Any, options StructuredSerializeOptions) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStructuredClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncStructuredClone1 returns true if the method "Window.structuredClone" exists.
func (this Window) HasFuncStructuredClone1() bool {
	return js.True == bindings.HasFuncWindowStructuredClone1(
		this.ref,
	)
}

// FuncStructuredClone1 returns the method "Window.structuredClone".
func (this Window) FuncStructuredClone1() (fn js.Func[func(value js.Any) js.Any]) {
	bindings.FuncWindowStructuredClone1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StructuredClone1 calls the method "Window.structuredClone".
func (this Window) StructuredClone1(value js.Any) (ret js.Any) {
	bindings.CallWindowStructuredClone1(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryStructuredClone1 calls the method "Window.structuredClone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryStructuredClone1(value js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowStructuredClone1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncFetch returns true if the method "Window.fetch" exists.
func (this Window) HasFuncFetch() bool {
	return js.True == bindings.HasFuncWindowFetch(
		this.ref,
	)
}

// FuncFetch returns the method "Window.fetch".
func (this Window) FuncFetch() (fn js.Func[func(input RequestInfo, init RequestInit) js.Promise[Response]]) {
	bindings.FuncWindowFetch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch calls the method "Window.fetch".
func (this Window) Fetch(input RequestInfo, init RequestInit) (ret js.Promise[Response]) {
	bindings.CallWindowFetch(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryFetch calls the method "Window.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryFetch(input RequestInfo, init RequestInit) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFetch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncFetch1 returns true if the method "Window.fetch" exists.
func (this Window) HasFuncFetch1() bool {
	return js.True == bindings.HasFuncWindowFetch1(
		this.ref,
	)
}

// FuncFetch1 returns the method "Window.fetch".
func (this Window) FuncFetch1() (fn js.Func[func(input RequestInfo) js.Promise[Response]]) {
	bindings.FuncWindowFetch1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch1 calls the method "Window.fetch".
func (this Window) Fetch1(input RequestInfo) (ret js.Promise[Response]) {
	bindings.CallWindowFetch1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFetch1 calls the method "Window.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryFetch1(input RequestInfo) (ret js.Promise[Response], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowFetch1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncRequestAnimationFrame returns true if the method "Window.requestAnimationFrame" exists.
func (this Window) HasFuncRequestAnimationFrame() bool {
	return js.True == bindings.HasFuncWindowRequestAnimationFrame(
		this.ref,
	)
}

// FuncRequestAnimationFrame returns the method "Window.requestAnimationFrame".
func (this Window) FuncRequestAnimationFrame() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	bindings.FuncWindowRequestAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAnimationFrame calls the method "Window.requestAnimationFrame".
func (this Window) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32) {
	bindings.CallWindowRequestAnimationFrame(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestAnimationFrame calls the method "Window.requestAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryRequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowRequestAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCancelAnimationFrame returns true if the method "Window.cancelAnimationFrame" exists.
func (this Window) HasFuncCancelAnimationFrame() bool {
	return js.True == bindings.HasFuncWindowCancelAnimationFrame(
		this.ref,
	)
}

// FuncCancelAnimationFrame returns the method "Window.cancelAnimationFrame".
func (this Window) FuncCancelAnimationFrame() (fn js.Func[func(handle uint32)]) {
	bindings.FuncWindowCancelAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelAnimationFrame calls the method "Window.cancelAnimationFrame".
func (this Window) CancelAnimationFrame(handle uint32) (ret js.Void) {
	bindings.CallWindowCancelAnimationFrame(
		this.ref, js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelAnimationFrame calls the method "Window.cancelAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Window) TryCancelAnimationFrame(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWindowCancelAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CompositionEventInit) UpdateFrom(ref js.Ref) {
	bindings.CompositionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CompositionEventInit) Update(ref js.Ref) {
	bindings.CompositionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CompositionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.View.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.View = p.View.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "CompositionEvent.data".
//
// It returns ok=false if there is no such property.
func (this CompositionEvent) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCompositionEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitCompositionEvent returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasFuncInitCompositionEvent() bool {
	return js.True == bindings.HasFuncCompositionEventInitCompositionEvent(
		this.ref,
	)
}

// FuncInitCompositionEvent returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) FuncInitCompositionEvent() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String)]) {
	bindings.FuncCompositionEventInitCompositionEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCompositionEvent calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		dataArg.Ref(),
	)

	return
}

// TryInitCompositionEvent calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object, dataArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		dataArg.Ref(),
	)

	return
}

// HasFuncInitCompositionEvent1 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasFuncInitCompositionEvent1() bool {
	return js.True == bindings.HasFuncCompositionEventInitCompositionEvent1(
		this.ref,
	)
}

// FuncInitCompositionEvent1 returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) FuncInitCompositionEvent1() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object)]) {
	bindings.FuncCompositionEventInitCompositionEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCompositionEvent1 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent1(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitCompositionEvent1 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasFuncInitCompositionEvent2 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasFuncInitCompositionEvent2() bool {
	return js.True == bindings.HasFuncCompositionEventInitCompositionEvent2(
		this.ref,
	)
}

// FuncInitCompositionEvent2 returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) FuncInitCompositionEvent2() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	bindings.FuncCompositionEventInitCompositionEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCompositionEvent2 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent2(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitCompositionEvent2 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasFuncInitCompositionEvent3 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasFuncInitCompositionEvent3() bool {
	return js.True == bindings.HasFuncCompositionEventInitCompositionEvent3(
		this.ref,
	)
}

// FuncInitCompositionEvent3 returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) FuncInitCompositionEvent3() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	bindings.FuncCompositionEventInitCompositionEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCompositionEvent3 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent3(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent3(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitCompositionEvent3 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent3(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasFuncInitCompositionEvent4 returns true if the method "CompositionEvent.initCompositionEvent" exists.
func (this CompositionEvent) HasFuncInitCompositionEvent4() bool {
	return js.True == bindings.HasFuncCompositionEventInitCompositionEvent4(
		this.ref,
	)
}

// FuncInitCompositionEvent4 returns the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) FuncInitCompositionEvent4() (fn js.Func[func(typeArg js.String)]) {
	bindings.FuncCompositionEventInitCompositionEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCompositionEvent4 calls the method "CompositionEvent.initCompositionEvent".
func (this CompositionEvent) InitCompositionEvent4(typeArg js.String) (ret js.Void) {
	bindings.CallCompositionEventInitCompositionEvent4(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitCompositionEvent4 calls the method "CompositionEvent.initCompositionEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CompositionEvent) TryInitCompositionEvent4(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompositionEventInitCompositionEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Readable returns the value of property "CompressionStream.readable".
//
// It returns ok=false if there is no such property.
func (this CompressionStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetCompressionStreamReadable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "CompressionStream.writable".
//
// It returns ok=false if there is no such property.
func (this CompressionStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetCompressionStreamWritable(
		this.ref, js.Pointer(&ret),
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
func (p *ContentIndexEventInit) UpdateFrom(ref js.Ref) {
	bindings.ContentIndexEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentIndexEventInit) Update(ref js.Ref) {
	bindings.ContentIndexEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentIndexEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "ContentIndexEvent.id".
//
// It returns ok=false if there is no such property.
func (this ContentIndexEvent) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetContentIndexEventId(
		this.ref, js.Pointer(&ret),
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
func (p *ContentVisibilityAutoStateChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.ContentVisibilityAutoStateChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentVisibilityAutoStateChangeEventInit) Update(ref js.Ref) {
	bindings.ContentVisibilityAutoStateChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentVisibilityAutoStateChangeEventInit) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Skipped returns the value of property "ContentVisibilityAutoStateChangeEvent.skipped".
//
// It returns ok=false if there is no such property.
func (this ContentVisibilityAutoStateChangeEvent) Skipped() (ret bool, ok bool) {
	ok = js.True == bindings.GetContentVisibilityAutoStateChangeEventSkipped(
		this.ref, js.Pointer(&ret),
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
func (p *CookieChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.CookieChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CookieChangeEventInit) Update(ref js.Ref) {
	bindings.CookieChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CookieChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Changed.Ref(),
		p.Deleted.Ref(),
	)
	p.Changed = p.Changed.FromRef(js.Undefined)
	p.Deleted = p.Deleted.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Changed returns the value of property "CookieChangeEvent.changed".
//
// It returns ok=false if there is no such property.
func (this CookieChangeEvent) Changed() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetCookieChangeEventChanged(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Deleted returns the value of property "CookieChangeEvent.deleted".
//
// It returns ok=false if there is no such property.
func (this CookieChangeEvent) Deleted() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetCookieChangeEventDeleted(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HighWaterMark returns the value of property "CountQueuingStrategy.highWaterMark".
//
// It returns ok=false if there is no such property.
func (this CountQueuingStrategy) HighWaterMark() (ret float64, ok bool) {
	ok = js.True == bindings.GetCountQueuingStrategyHighWaterMark(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Size returns the value of property "CountQueuingStrategy.size".
//
// It returns ok=false if there is no such property.
func (this CountQueuingStrategy) Size() (ret js.Func[func(arguments ...js.Any) js.Any], ok bool) {
	ok = js.True == bindings.GetCountQueuingStrategySize(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CrashReportBody struct {
	ReportBody
}

func (this CrashReportBody) Once() CrashReportBody {
	this.ref.Once()
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
	this.ref.Free()
}

// Reason returns the value of property "CrashReportBody.reason".
//
// It returns ok=false if there is no such property.
func (this CrashReportBody) Reason() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCrashReportBodyReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "CrashReportBody.toJSON" exists.
func (this CrashReportBody) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncCrashReportBodyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "CrashReportBody.toJSON".
func (this CrashReportBody) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncCrashReportBodyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "CrashReportBody.toJSON".
func (this CrashReportBody) ToJSON() (ret js.Object) {
	bindings.CallCrashReportBodyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "CrashReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CrashReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCrashReportBodyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CredentialData) UpdateFrom(ref js.Ref) {
	bindings.CredentialDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialData) Update(ref js.Ref) {
	bindings.CredentialDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialData) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
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
func (p *CryptoKeyPair) UpdateFrom(ref js.Ref) {
	bindings.CryptoKeyPairJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CryptoKeyPair) Update(ref js.Ref) {
	bindings.CryptoKeyPairJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CryptoKeyPair) FreeMembers(recursive bool) {
	js.Free(
		p.PublicKey.Ref(),
		p.PrivateKey.Ref(),
	)
	p.PublicKey = p.PublicKey.FromRef(js.Undefined)
	p.PrivateKey = p.PrivateKey.FromRef(js.Undefined)
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
func (p *CustomEventInit) UpdateFrom(ref js.Ref) {
	bindings.CustomEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CustomEventInit) Update(ref js.Ref) {
	bindings.CustomEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CustomEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Detail.Ref(),
	)
	p.Detail = p.Detail.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Detail returns the value of property "CustomEvent.detail".
//
// It returns ok=false if there is no such property.
func (this CustomEvent) Detail() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetCustomEventDetail(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitCustomEvent returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasFuncInitCustomEvent() bool {
	return js.True == bindings.HasFuncCustomEventInitCustomEvent(
		this.ref,
	)
}

// FuncInitCustomEvent returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) FuncInitCustomEvent() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, detail js.Any)]) {
	bindings.FuncCustomEventInitCustomEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCustomEvent calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent(typ js.String, bubbles bool, cancelable bool, detail js.Any) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		detail.Ref(),
	)

	return
}

// TryInitCustomEvent calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomEvent) TryInitCustomEvent(typ js.String, bubbles bool, cancelable bool, detail js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		detail.Ref(),
	)

	return
}

// HasFuncInitCustomEvent1 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasFuncInitCustomEvent1() bool {
	return js.True == bindings.HasFuncCustomEventInitCustomEvent1(
		this.ref,
	)
}

// FuncInitCustomEvent1 returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) FuncInitCustomEvent1() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	bindings.FuncCustomEventInitCustomEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCustomEvent1 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent1(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitCustomEvent1 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomEvent) TryInitCustomEvent1(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasFuncInitCustomEvent2 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasFuncInitCustomEvent2() bool {
	return js.True == bindings.HasFuncCustomEventInitCustomEvent2(
		this.ref,
	)
}

// FuncInitCustomEvent2 returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) FuncInitCustomEvent2() (fn js.Func[func(typ js.String, bubbles bool)]) {
	bindings.FuncCustomEventInitCustomEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCustomEvent2 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent2(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent2(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitCustomEvent2 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomEvent) TryInitCustomEvent2(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasFuncInitCustomEvent3 returns true if the method "CustomEvent.initCustomEvent" exists.
func (this CustomEvent) HasFuncInitCustomEvent3() bool {
	return js.True == bindings.HasFuncCustomEventInitCustomEvent3(
		this.ref,
	)
}

// FuncInitCustomEvent3 returns the method "CustomEvent.initCustomEvent".
func (this CustomEvent) FuncInitCustomEvent3() (fn js.Func[func(typ js.String)]) {
	bindings.FuncCustomEventInitCustomEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitCustomEvent3 calls the method "CustomEvent.initCustomEvent".
func (this CustomEvent) InitCustomEvent3(typ js.String) (ret js.Void) {
	bindings.CallCustomEventInitCustomEvent3(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitCustomEvent3 calls the method "CustomEvent.initCustomEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomEvent) TryInitCustomEvent3(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomEventInitCustomEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncParseFromString returns true if the method "DOMParser.parseFromString" exists.
func (this DOMParser) HasFuncParseFromString() bool {
	return js.True == bindings.HasFuncDOMParserParseFromString(
		this.ref,
	)
}

// FuncParseFromString returns the method "DOMParser.parseFromString".
func (this DOMParser) FuncParseFromString() (fn js.Func[func(string js.String, typ DOMParserSupportedType) Document]) {
	bindings.FuncDOMParserParseFromString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ParseFromString calls the method "DOMParser.parseFromString".
func (this DOMParser) ParseFromString(string js.String, typ DOMParserSupportedType) (ret Document) {
	bindings.CallDOMParserParseFromString(
		this.ref, js.Pointer(&ret),
		string.Ref(),
		uint32(typ),
	)

	return
}

// TryParseFromString calls the method "DOMParser.parseFromString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMParser) TryParseFromString(string js.String, typ DOMParserSupportedType) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMParserParseFromString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "DataCue.value".
//
// It returns ok=false if there is no such property.
func (this DataCue) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetDataCueValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "DataCue.value" to val.
//
// It returns false if the property cannot be set.
func (this DataCue) SetValue(val js.Any) bool {
	return js.True == bindings.SetDataCueValue(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "DataCue.type".
//
// It returns ok=false if there is no such property.
func (this DataCue) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDataCueType(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Readable returns the value of property "DecompressionStream.readable".
//
// It returns ok=false if there is no such property.
func (this DecompressionStream) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetDecompressionStreamReadable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Writable returns the value of property "DecompressionStream.writable".
//
// It returns ok=false if there is no such property.
func (this DecompressionStream) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetDecompressionStreamWritable(
		this.ref, js.Pointer(&ret),
	)
	return
}

type DedicatedWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this DedicatedWorkerGlobalScope) Once() DedicatedWorkerGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "DedicatedWorkerGlobalScope.name".
//
// It returns ok=false if there is no such property.
func (this DedicatedWorkerGlobalScope) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDedicatedWorkerGlobalScopeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPostMessage returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopePostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) FuncPostMessage() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	bindings.FuncDedicatedWorkerGlobalScopePostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasFuncPostMessage1 returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasFuncPostMessage1() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopePostMessage1(
		this.ref,
	)
}

// FuncPostMessage1 returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) FuncPostMessage1() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	bindings.FuncDedicatedWorkerGlobalScopePostMessage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage1 calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPostMessage2 returns true if the method "DedicatedWorkerGlobalScope.postMessage" exists.
func (this DedicatedWorkerGlobalScope) HasFuncPostMessage2() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopePostMessage2(
		this.ref,
	)
}

// FuncPostMessage2 returns the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) FuncPostMessage2() (fn js.Func[func(message js.Any)]) {
	bindings.FuncDedicatedWorkerGlobalScopePostMessage2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage2 calls the method "DedicatedWorkerGlobalScope.postMessage".
func (this DedicatedWorkerGlobalScope) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopePostMessage2(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "DedicatedWorkerGlobalScope.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopePostMessage2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncClose returns true if the method "DedicatedWorkerGlobalScope.close" exists.
func (this DedicatedWorkerGlobalScope) HasFuncClose() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopeClose(
		this.ref,
	)
}

// FuncClose returns the method "DedicatedWorkerGlobalScope.close".
func (this DedicatedWorkerGlobalScope) FuncClose() (fn js.Func[func()]) {
	bindings.FuncDedicatedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "DedicatedWorkerGlobalScope.close".
func (this DedicatedWorkerGlobalScope) Close() (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "DedicatedWorkerGlobalScope.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestAnimationFrame returns true if the method "DedicatedWorkerGlobalScope.requestAnimationFrame" exists.
func (this DedicatedWorkerGlobalScope) HasFuncRequestAnimationFrame() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.ref,
	)
}

// FuncRequestAnimationFrame returns the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
func (this DedicatedWorkerGlobalScope) FuncRequestAnimationFrame() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp)]) uint32]) {
	bindings.FuncDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAnimationFrame calls the method "DedicatedWorkerGlobalScope.requestAnimationFrame".
func (this DedicatedWorkerGlobalScope) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32) {
	bindings.CallDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestAnimationFrame calls the method "DedicatedWorkerGlobalScope.requestAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryRequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeRequestAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCancelAnimationFrame returns true if the method "DedicatedWorkerGlobalScope.cancelAnimationFrame" exists.
func (this DedicatedWorkerGlobalScope) HasFuncCancelAnimationFrame() bool {
	return js.True == bindings.HasFuncDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.ref,
	)
}

// FuncCancelAnimationFrame returns the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
func (this DedicatedWorkerGlobalScope) FuncCancelAnimationFrame() (fn js.Func[func(handle uint32)]) {
	bindings.FuncDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelAnimationFrame calls the method "DedicatedWorkerGlobalScope.cancelAnimationFrame".
func (this DedicatedWorkerGlobalScope) CancelAnimationFrame(handle uint32) (ret js.Void) {
	bindings.CallDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.ref, js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelAnimationFrame calls the method "DedicatedWorkerGlobalScope.cancelAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DedicatedWorkerGlobalScope) TryCancelAnimationFrame(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDedicatedWorkerGlobalScopeCancelAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

type DeprecationReportBody struct {
	ReportBody
}

func (this DeprecationReportBody) Once() DeprecationReportBody {
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "DeprecationReportBody.id".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AnticipatedRemoval returns the value of property "DeprecationReportBody.anticipatedRemoval".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) AnticipatedRemoval() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyAnticipatedRemoval(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "DeprecationReportBody.message".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "DeprecationReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodySourceFile(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "DeprecationReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyLineNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "DeprecationReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this DeprecationReportBody) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDeprecationReportBodyColumnNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "DeprecationReportBody.toJSON" exists.
func (this DeprecationReportBody) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncDeprecationReportBodyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "DeprecationReportBody.toJSON".
func (this DeprecationReportBody) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncDeprecationReportBodyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "DeprecationReportBody.toJSON".
func (this DeprecationReportBody) ToJSON() (ret js.Object) {
	bindings.CallDeprecationReportBodyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DeprecationReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DeprecationReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeprecationReportBodyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *Landmark) UpdateFrom(ref js.Ref) {
	bindings.LandmarkJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Landmark) Update(ref js.Ref) {
	bindings.LandmarkJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Landmark) FreeMembers(recursive bool) {
	js.Free(
		p.Locations.Ref(),
	)
	p.Locations = p.Locations.FromRef(js.Undefined)
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
func (p *DetectedFace) UpdateFrom(ref js.Ref) {
	bindings.DetectedFaceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DetectedFace) Update(ref js.Ref) {
	bindings.DetectedFaceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DetectedFace) FreeMembers(recursive bool) {
	js.Free(
		p.BoundingBox.Ref(),
		p.Landmarks.Ref(),
	)
	p.BoundingBox = p.BoundingBox.FromRef(js.Undefined)
	p.Landmarks = p.Landmarks.FromRef(js.Undefined)
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
func (p *DetectedText) UpdateFrom(ref js.Ref) {
	bindings.DetectedTextJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DetectedText) Update(ref js.Ref) {
	bindings.DetectedTextJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DetectedText) FreeMembers(recursive bool) {
	js.Free(
		p.BoundingBox.Ref(),
		p.RawValue.Ref(),
		p.CornerPoints.Ref(),
	)
	p.BoundingBox = p.BoundingBox.FromRef(js.Undefined)
	p.RawValue = p.RawValue.FromRef(js.Undefined)
	p.CornerPoints = p.CornerPoints.FromRef(js.Undefined)
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
func (p *DeviceMotionEventAccelerationInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventAccelerationInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceMotionEventAccelerationInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventAccelerationInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceMotionEventAccelerationInit) FreeMembers(recursive bool) {
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
func (p *DeviceMotionEventRotationRateInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventRotationRateInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceMotionEventRotationRateInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventRotationRateInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceMotionEventRotationRateInit) FreeMembers(recursive bool) {
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
func (p *DeviceMotionEventInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceMotionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceMotionEventInit) Update(ref js.Ref) {
	bindings.DeviceMotionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceMotionEventInit) FreeMembers(recursive bool) {
	if recursive {
		p.Acceleration.FreeMembers(true)
		p.AccelerationIncludingGravity.FreeMembers(true)
		p.RotationRate.FreeMembers(true)
	}
}

type DeviceMotionEventAcceleration struct {
	ref js.Ref
}

func (this DeviceMotionEventAcceleration) Once() DeviceMotionEventAcceleration {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "DeviceMotionEventAcceleration.x".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "DeviceMotionEventAcceleration.y".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "DeviceMotionEventAcceleration.z".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventAcceleration) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

type DeviceMotionEventRotationRate struct {
	ref js.Ref
}

func (this DeviceMotionEventRotationRate) Once() DeviceMotionEventRotationRate {
	this.ref.Once()
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
	this.ref.Free()
}

// Alpha returns the value of property "DeviceMotionEventRotationRate.alpha".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Alpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Beta returns the value of property "DeviceMotionEventRotationRate.beta".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Beta() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateBeta(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gamma returns the value of property "DeviceMotionEventRotationRate.gamma".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEventRotationRate) Gamma() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRateGamma(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Acceleration returns the value of property "DeviceMotionEvent.acceleration".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) Acceleration() (ret DeviceMotionEventAcceleration, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAcceleration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AccelerationIncludingGravity returns the value of property "DeviceMotionEvent.accelerationIncludingGravity".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) AccelerationIncludingGravity() (ret DeviceMotionEventAcceleration, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventAccelerationIncludingGravity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RotationRate returns the value of property "DeviceMotionEvent.rotationRate".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) RotationRate() (ret DeviceMotionEventRotationRate, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventRotationRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Interval returns the value of property "DeviceMotionEvent.interval".
//
// It returns ok=false if there is no such property.
func (this DeviceMotionEvent) Interval() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceMotionEventInterval(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestPermission returns true if the static method "DeviceMotionEvent.requestPermission" exists.
func (this DeviceMotionEvent) HasFuncRequestPermission() bool {
	return js.True == bindings.HasFuncDeviceMotionEventRequestPermission(
		this.ref,
	)
}

// FuncRequestPermission returns the static method "DeviceMotionEvent.requestPermission".
func (this DeviceMotionEvent) FuncRequestPermission() (fn js.Func[func() js.Promise[PermissionState]]) {
	bindings.FuncDeviceMotionEventRequestPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission calls the static method "DeviceMotionEvent.requestPermission".
func (this DeviceMotionEvent) RequestPermission() (ret js.Promise[PermissionState]) {
	bindings.CallDeviceMotionEventRequestPermission(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPermission calls the static method "DeviceMotionEvent.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DeviceMotionEvent) TryRequestPermission() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeviceMotionEventRequestPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

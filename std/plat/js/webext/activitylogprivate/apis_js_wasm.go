// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package activitylogprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/activitylogprivate/bindings"
)

type ExtensionActivityType uint32

const (
	_ ExtensionActivityType = iota

	ExtensionActivityType_API_CALL
	ExtensionActivityType_API_EVENT
	ExtensionActivityType_CONTENT_SCRIPT
	ExtensionActivityType_DOM_ACCESS
	ExtensionActivityType_DOM_EVENT
	ExtensionActivityType_WEB_REQUEST
)

func (ExtensionActivityType) FromRef(str js.Ref) ExtensionActivityType {
	return ExtensionActivityType(bindings.ConstOfExtensionActivityType(str))
}

func (x ExtensionActivityType) String() (string, bool) {
	switch x {
	case ExtensionActivityType_API_CALL:
		return "api_call", true
	case ExtensionActivityType_API_EVENT:
		return "api_event", true
	case ExtensionActivityType_CONTENT_SCRIPT:
		return "content_script", true
	case ExtensionActivityType_DOM_ACCESS:
		return "dom_access", true
	case ExtensionActivityType_DOM_EVENT:
		return "dom_event", true
	case ExtensionActivityType_WEB_REQUEST:
		return "web_request", true
	default:
		return "", false
	}
}

type ExtensionActivityDomVerb uint32

const (
	_ ExtensionActivityDomVerb = iota

	ExtensionActivityDomVerb_GETTER
	ExtensionActivityDomVerb_SETTER
	ExtensionActivityDomVerb_METHOD
	ExtensionActivityDomVerb_INSERTED
	ExtensionActivityDomVerb_XHR
	ExtensionActivityDomVerb_WEBREQUEST
	ExtensionActivityDomVerb_MODIFIED
)

func (ExtensionActivityDomVerb) FromRef(str js.Ref) ExtensionActivityDomVerb {
	return ExtensionActivityDomVerb(bindings.ConstOfExtensionActivityDomVerb(str))
}

func (x ExtensionActivityDomVerb) String() (string, bool) {
	switch x {
	case ExtensionActivityDomVerb_GETTER:
		return "getter", true
	case ExtensionActivityDomVerb_SETTER:
		return "setter", true
	case ExtensionActivityDomVerb_METHOD:
		return "method", true
	case ExtensionActivityDomVerb_INSERTED:
		return "inserted", true
	case ExtensionActivityDomVerb_XHR:
		return "xhr", true
	case ExtensionActivityDomVerb_WEBREQUEST:
		return "webrequest", true
	case ExtensionActivityDomVerb_MODIFIED:
		return "modified", true
	default:
		return "", false
	}
}

type ExtensionActivityFieldOther struct {
	// DomVerb is "ExtensionActivityFieldOther.domVerb"
	//
	// Optional
	DomVerb ExtensionActivityDomVerb
	// Extra is "ExtensionActivityFieldOther.extra"
	//
	// Optional
	Extra js.String
	// Prerender is "ExtensionActivityFieldOther.prerender"
	//
	// Optional
	//
	// NOTE: FFI_USE_Prerender MUST be set to true to make this field effective.
	Prerender bool
	// WebRequest is "ExtensionActivityFieldOther.webRequest"
	//
	// Optional
	WebRequest js.String

	FFI_USE_Prerender bool // for Prerender.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionActivityFieldOther with all fields set.
func (p ExtensionActivityFieldOther) FromRef(ref js.Ref) ExtensionActivityFieldOther {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionActivityFieldOther in the application heap.
func (p ExtensionActivityFieldOther) New() js.Ref {
	return bindings.ExtensionActivityFieldOtherJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionActivityFieldOther) UpdateFrom(ref js.Ref) {
	bindings.ExtensionActivityFieldOtherJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionActivityFieldOther) Update(ref js.Ref) {
	bindings.ExtensionActivityFieldOtherJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionActivityFieldOther) FreeMembers(recursive bool) {
	js.Free(
		p.Extra.Ref(),
		p.WebRequest.Ref(),
	)
	p.Extra = p.Extra.FromRef(js.Undefined)
	p.WebRequest = p.WebRequest.FromRef(js.Undefined)
}

type ExtensionActivity struct {
	// ActivityId is "ExtensionActivity.activityId"
	//
	// Optional
	ActivityId js.String
	// ActivityType is "ExtensionActivity.activityType"
	//
	// Required
	ActivityType ExtensionActivityType
	// ApiCall is "ExtensionActivity.apiCall"
	//
	// Optional
	ApiCall js.String
	// ArgUrl is "ExtensionActivity.argUrl"
	//
	// Optional
	ArgUrl js.String
	// Args is "ExtensionActivity.args"
	//
	// Optional
	Args js.String
	// Count is "ExtensionActivity.count"
	//
	// Optional
	//
	// NOTE: FFI_USE_Count MUST be set to true to make this field effective.
	Count float64
	// ExtensionId is "ExtensionActivity.extensionId"
	//
	// Optional
	ExtensionId js.String
	// Other is "ExtensionActivity.other"
	//
	// Optional
	//
	// NOTE: Other.FFI_USE MUST be set to true to get Other used.
	Other ExtensionActivityFieldOther
	// PageTitle is "ExtensionActivity.pageTitle"
	//
	// Optional
	PageTitle js.String
	// PageUrl is "ExtensionActivity.pageUrl"
	//
	// Optional
	PageUrl js.String
	// Time is "ExtensionActivity.time"
	//
	// Optional
	//
	// NOTE: FFI_USE_Time MUST be set to true to make this field effective.
	Time float64

	FFI_USE_Count bool // for Count.
	FFI_USE_Time  bool // for Time.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionActivity with all fields set.
func (p ExtensionActivity) FromRef(ref js.Ref) ExtensionActivity {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionActivity in the application heap.
func (p ExtensionActivity) New() js.Ref {
	return bindings.ExtensionActivityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionActivity) UpdateFrom(ref js.Ref) {
	bindings.ExtensionActivityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionActivity) Update(ref js.Ref) {
	bindings.ExtensionActivityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionActivity) FreeMembers(recursive bool) {
	js.Free(
		p.ActivityId.Ref(),
		p.ApiCall.Ref(),
		p.ArgUrl.Ref(),
		p.Args.Ref(),
		p.ExtensionId.Ref(),
		p.PageTitle.Ref(),
		p.PageUrl.Ref(),
	)
	p.ActivityId = p.ActivityId.FromRef(js.Undefined)
	p.ApiCall = p.ApiCall.FromRef(js.Undefined)
	p.ArgUrl = p.ArgUrl.FromRef(js.Undefined)
	p.Args = p.Args.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.PageTitle = p.PageTitle.FromRef(js.Undefined)
	p.PageUrl = p.PageUrl.FromRef(js.Undefined)
	if recursive {
		p.Other.FreeMembers(true)
	}
}

type ActivityResultSet struct {
	// Activities is "ActivityResultSet.activities"
	//
	// Required
	Activities js.Array[ExtensionActivity]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ActivityResultSet with all fields set.
func (p ActivityResultSet) FromRef(ref js.Ref) ActivityResultSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ActivityResultSet in the application heap.
func (p ActivityResultSet) New() js.Ref {
	return bindings.ActivityResultSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ActivityResultSet) UpdateFrom(ref js.Ref) {
	bindings.ActivityResultSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ActivityResultSet) Update(ref js.Ref) {
	bindings.ActivityResultSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ActivityResultSet) FreeMembers(recursive bool) {
	js.Free(
		p.Activities.Ref(),
	)
	p.Activities = p.Activities.FromRef(js.Undefined)
}

type ExtensionActivityFilter uint32

const (
	_ ExtensionActivityFilter = iota

	ExtensionActivityFilter_API_CALL
	ExtensionActivityFilter_API_EVENT
	ExtensionActivityFilter_CONTENT_SCRIPT
	ExtensionActivityFilter_DOM_ACCESS
	ExtensionActivityFilter_DOM_EVENT
	ExtensionActivityFilter_WEB_REQUEST
	ExtensionActivityFilter_ANY
)

func (ExtensionActivityFilter) FromRef(str js.Ref) ExtensionActivityFilter {
	return ExtensionActivityFilter(bindings.ConstOfExtensionActivityFilter(str))
}

func (x ExtensionActivityFilter) String() (string, bool) {
	switch x {
	case ExtensionActivityFilter_API_CALL:
		return "api_call", true
	case ExtensionActivityFilter_API_EVENT:
		return "api_event", true
	case ExtensionActivityFilter_CONTENT_SCRIPT:
		return "content_script", true
	case ExtensionActivityFilter_DOM_ACCESS:
		return "dom_access", true
	case ExtensionActivityFilter_DOM_EVENT:
		return "dom_event", true
	case ExtensionActivityFilter_WEB_REQUEST:
		return "web_request", true
	case ExtensionActivityFilter_ANY:
		return "any", true
	default:
		return "", false
	}
}

type Filter struct {
	// ActivityType is "Filter.activityType"
	//
	// Required
	ActivityType ExtensionActivityFilter
	// ApiCall is "Filter.apiCall"
	//
	// Optional
	ApiCall js.String
	// ArgUrl is "Filter.argUrl"
	//
	// Optional
	ArgUrl js.String
	// DaysAgo is "Filter.daysAgo"
	//
	// Optional
	//
	// NOTE: FFI_USE_DaysAgo MUST be set to true to make this field effective.
	DaysAgo int64
	// ExtensionId is "Filter.extensionId"
	//
	// Optional
	ExtensionId js.String
	// PageUrl is "Filter.pageUrl"
	//
	// Optional
	PageUrl js.String

	FFI_USE_DaysAgo bool // for DaysAgo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Filter with all fields set.
func (p Filter) FromRef(ref js.Ref) Filter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Filter in the application heap.
func (p Filter) New() js.Ref {
	return bindings.FilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Filter) UpdateFrom(ref js.Ref) {
	bindings.FilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Filter) Update(ref js.Ref) {
	bindings.FilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Filter) FreeMembers(recursive bool) {
	js.Free(
		p.ApiCall.Ref(),
		p.ArgUrl.Ref(),
		p.ExtensionId.Ref(),
		p.PageUrl.Ref(),
	)
	p.ApiCall = p.ApiCall.FromRef(js.Undefined)
	p.ArgUrl = p.ArgUrl.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.PageUrl = p.PageUrl.FromRef(js.Undefined)
}

// HasFuncDeleteActivities returns true if the function "WEBEXT.activityLogPrivate.deleteActivities" exists.
func HasFuncDeleteActivities() bool {
	return js.True == bindings.HasFuncDeleteActivities()
}

// FuncDeleteActivities returns the function "WEBEXT.activityLogPrivate.deleteActivities".
func FuncDeleteActivities() (fn js.Func[func(activityIds js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncDeleteActivities(
		js.Pointer(&fn),
	)
	return
}

// DeleteActivities calls the function "WEBEXT.activityLogPrivate.deleteActivities" directly.
func DeleteActivities(activityIds js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallDeleteActivities(
		js.Pointer(&ret),
		activityIds.Ref(),
	)

	return
}

// TryDeleteActivities calls the function "WEBEXT.activityLogPrivate.deleteActivities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteActivities(activityIds js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteActivities(
		js.Pointer(&ret), js.Pointer(&exception),
		activityIds.Ref(),
	)

	return
}

// HasFuncDeleteActivitiesByExtension returns true if the function "WEBEXT.activityLogPrivate.deleteActivitiesByExtension" exists.
func HasFuncDeleteActivitiesByExtension() bool {
	return js.True == bindings.HasFuncDeleteActivitiesByExtension()
}

// FuncDeleteActivitiesByExtension returns the function "WEBEXT.activityLogPrivate.deleteActivitiesByExtension".
func FuncDeleteActivitiesByExtension() (fn js.Func[func(extensionId js.String) js.Promise[js.Void]]) {
	bindings.FuncDeleteActivitiesByExtension(
		js.Pointer(&fn),
	)
	return
}

// DeleteActivitiesByExtension calls the function "WEBEXT.activityLogPrivate.deleteActivitiesByExtension" directly.
func DeleteActivitiesByExtension(extensionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallDeleteActivitiesByExtension(
		js.Pointer(&ret),
		extensionId.Ref(),
	)

	return
}

// TryDeleteActivitiesByExtension calls the function "WEBEXT.activityLogPrivate.deleteActivitiesByExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteActivitiesByExtension(extensionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteActivitiesByExtension(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
	)

	return
}

// HasFuncDeleteDatabase returns true if the function "WEBEXT.activityLogPrivate.deleteDatabase" exists.
func HasFuncDeleteDatabase() bool {
	return js.True == bindings.HasFuncDeleteDatabase()
}

// FuncDeleteDatabase returns the function "WEBEXT.activityLogPrivate.deleteDatabase".
func FuncDeleteDatabase() (fn js.Func[func()]) {
	bindings.FuncDeleteDatabase(
		js.Pointer(&fn),
	)
	return
}

// DeleteDatabase calls the function "WEBEXT.activityLogPrivate.deleteDatabase" directly.
func DeleteDatabase() (ret js.Void) {
	bindings.CallDeleteDatabase(
		js.Pointer(&ret),
	)

	return
}

// TryDeleteDatabase calls the function "WEBEXT.activityLogPrivate.deleteDatabase"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteDatabase() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteDatabase(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteUrls returns true if the function "WEBEXT.activityLogPrivate.deleteUrls" exists.
func HasFuncDeleteUrls() bool {
	return js.True == bindings.HasFuncDeleteUrls()
}

// FuncDeleteUrls returns the function "WEBEXT.activityLogPrivate.deleteUrls".
func FuncDeleteUrls() (fn js.Func[func(urls js.Array[js.String])]) {
	bindings.FuncDeleteUrls(
		js.Pointer(&fn),
	)
	return
}

// DeleteUrls calls the function "WEBEXT.activityLogPrivate.deleteUrls" directly.
func DeleteUrls(urls js.Array[js.String]) (ret js.Void) {
	bindings.CallDeleteUrls(
		js.Pointer(&ret),
		urls.Ref(),
	)

	return
}

// TryDeleteUrls calls the function "WEBEXT.activityLogPrivate.deleteUrls"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteUrls(urls js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteUrls(
		js.Pointer(&ret), js.Pointer(&exception),
		urls.Ref(),
	)

	return
}

// HasFuncGetExtensionActivities returns true if the function "WEBEXT.activityLogPrivate.getExtensionActivities" exists.
func HasFuncGetExtensionActivities() bool {
	return js.True == bindings.HasFuncGetExtensionActivities()
}

// FuncGetExtensionActivities returns the function "WEBEXT.activityLogPrivate.getExtensionActivities".
func FuncGetExtensionActivities() (fn js.Func[func(filter Filter) js.Promise[ActivityResultSet]]) {
	bindings.FuncGetExtensionActivities(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionActivities calls the function "WEBEXT.activityLogPrivate.getExtensionActivities" directly.
func GetExtensionActivities(filter Filter) (ret js.Promise[ActivityResultSet]) {
	bindings.CallGetExtensionActivities(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetExtensionActivities calls the function "WEBEXT.activityLogPrivate.getExtensionActivities"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionActivities(filter Filter) (ret js.Promise[ActivityResultSet], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionActivities(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

type OnExtensionActivityEventCallbackFunc func(this js.Ref, activity *ExtensionActivity) js.Ref

func (fn OnExtensionActivityEventCallbackFunc) Register() js.Func[func(activity *ExtensionActivity)] {
	return js.RegisterCallback[func(activity *ExtensionActivity)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExtensionActivityEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionActivity
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnExtensionActivityEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activity *ExtensionActivity) js.Ref
	Arg T
}

func (cb *OnExtensionActivityEventCallback[T]) Register() js.Func[func(activity *ExtensionActivity)] {
	return js.RegisterCallback[func(activity *ExtensionActivity)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExtensionActivityEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionActivity
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnExtensionActivity returns true if the function "WEBEXT.activityLogPrivate.onExtensionActivity.addListener" exists.
func HasFuncOnExtensionActivity() bool {
	return js.True == bindings.HasFuncOnExtensionActivity()
}

// FuncOnExtensionActivity returns the function "WEBEXT.activityLogPrivate.onExtensionActivity.addListener".
func FuncOnExtensionActivity() (fn js.Func[func(callback js.Func[func(activity *ExtensionActivity)])]) {
	bindings.FuncOnExtensionActivity(
		js.Pointer(&fn),
	)
	return
}

// OnExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.addListener" directly.
func OnExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret js.Void) {
	bindings.CallOnExtensionActivity(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExtensionActivity(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExtensionActivity returns true if the function "WEBEXT.activityLogPrivate.onExtensionActivity.removeListener" exists.
func HasFuncOffExtensionActivity() bool {
	return js.True == bindings.HasFuncOffExtensionActivity()
}

// FuncOffExtensionActivity returns the function "WEBEXT.activityLogPrivate.onExtensionActivity.removeListener".
func FuncOffExtensionActivity() (fn js.Func[func(callback js.Func[func(activity *ExtensionActivity)])]) {
	bindings.FuncOffExtensionActivity(
		js.Pointer(&fn),
	)
	return
}

// OffExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.removeListener" directly.
func OffExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret js.Void) {
	bindings.CallOffExtensionActivity(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExtensionActivity(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExtensionActivity returns true if the function "WEBEXT.activityLogPrivate.onExtensionActivity.hasListener" exists.
func HasFuncHasOnExtensionActivity() bool {
	return js.True == bindings.HasFuncHasOnExtensionActivity()
}

// FuncHasOnExtensionActivity returns the function "WEBEXT.activityLogPrivate.onExtensionActivity.hasListener".
func FuncHasOnExtensionActivity() (fn js.Func[func(callback js.Func[func(activity *ExtensionActivity)]) bool]) {
	bindings.FuncHasOnExtensionActivity(
		js.Pointer(&fn),
	)
	return
}

// HasOnExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.hasListener" directly.
func HasOnExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret bool) {
	bindings.CallHasOnExtensionActivity(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExtensionActivity calls the function "WEBEXT.activityLogPrivate.onExtensionActivity.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExtensionActivity(callback js.Func[func(activity *ExtensionActivity)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExtensionActivity(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

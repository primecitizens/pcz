// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sessions

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sessions/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
	"github.com/primecitizens/pcz/std/plat/js/webext/windows"
)

type Session struct {
	// LastModified is "Session.lastModified"
	//
	// Required
	LastModified int64
	// Tab is "Session.tab"
	//
	// Optional
	//
	// NOTE: Tab.FFI_USE MUST be set to true to get Tab used.
	Tab tabs.Tab
	// Window is "Session.window"
	//
	// Optional
	//
	// NOTE: Window.FFI_USE MUST be set to true to get Window used.
	Window windows.Window

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Session with all fields set.
func (p Session) FromRef(ref js.Ref) Session {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Session in the application heap.
func (p Session) New() js.Ref {
	return bindings.SessionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Session) UpdateFrom(ref js.Ref) {
	bindings.SessionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Session) Update(ref js.Ref) {
	bindings.SessionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Session) FreeMembers(recursive bool) {
	if recursive {
		p.Tab.FreeMembers(true)
		p.Window.FreeMembers(true)
	}
}

type Device struct {
	// DeviceName is "Device.deviceName"
	//
	// Required
	DeviceName js.String
	// Info is "Device.info"
	//
	// Required
	Info js.String
	// Sessions is "Device.sessions"
	//
	// Required
	Sessions js.Array[Session]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Device with all fields set.
func (p Device) FromRef(ref js.Ref) Device {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Device in the application heap.
func (p Device) New() js.Ref {
	return bindings.DeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Device) UpdateFrom(ref js.Ref) {
	bindings.DeviceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Device) Update(ref js.Ref) {
	bindings.DeviceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Device) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceName.Ref(),
		p.Info.Ref(),
		p.Sessions.Ref(),
	)
	p.DeviceName = p.DeviceName.FromRef(js.Undefined)
	p.Info = p.Info.FromRef(js.Undefined)
	p.Sessions = p.Sessions.FromRef(js.Undefined)
}

type Filter struct {
	// MaxResults is "Filter.maxResults"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxResults MUST be set to true to make this field effective.
	MaxResults int64

	FFI_USE_MaxResults bool // for MaxResults.

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
}

// MAX_SESSION_RESULTS returns the value of property "WEBEXT.sessions.MAX_SESSION_RESULTS".
//
// The returned bool will be false if there is no such property.
func MAX_SESSION_RESULTS() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_SESSION_RESULTS(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_SESSION_RESULTS sets the value of property "WEBEXT.sessions.MAX_SESSION_RESULTS" to val.
//
// It returns false if the property cannot be set.
func SetMAX_SESSION_RESULTS(val js.String) bool {
	return js.True == bindings.SetMAX_SESSION_RESULTS(
		val.Ref())
}

// HasFuncGetDevices returns true if the function "WEBEXT.sessions.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.sessions.getDevices".
func FuncGetDevices() (fn js.Func[func(filter Filter) js.Promise[js.Array[Device]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.sessions.getDevices" directly.
func GetDevices(filter Filter) (ret js.Promise[js.Array[Device]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.sessions.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices(filter Filter) (ret js.Promise[js.Array[Device]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncGetRecentlyClosed returns true if the function "WEBEXT.sessions.getRecentlyClosed" exists.
func HasFuncGetRecentlyClosed() bool {
	return js.True == bindings.HasFuncGetRecentlyClosed()
}

// FuncGetRecentlyClosed returns the function "WEBEXT.sessions.getRecentlyClosed".
func FuncGetRecentlyClosed() (fn js.Func[func(filter Filter) js.Promise[js.Array[Session]]]) {
	bindings.FuncGetRecentlyClosed(
		js.Pointer(&fn),
	)
	return
}

// GetRecentlyClosed calls the function "WEBEXT.sessions.getRecentlyClosed" directly.
func GetRecentlyClosed(filter Filter) (ret js.Promise[js.Array[Session]]) {
	bindings.CallGetRecentlyClosed(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetRecentlyClosed calls the function "WEBEXT.sessions.getRecentlyClosed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRecentlyClosed(filter Filter) (ret js.Promise[js.Array[Session]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRecentlyClosed(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

type OnChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnChanged returns true if the function "WEBEXT.sessions.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.sessions.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.sessions.onChanged.addListener" directly.
func OnChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.sessions.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.sessions.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.sessions.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.sessions.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.sessions.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.sessions.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.sessions.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.sessions.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.sessions.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRestore returns true if the function "WEBEXT.sessions.restore" exists.
func HasFuncRestore() bool {
	return js.True == bindings.HasFuncRestore()
}

// FuncRestore returns the function "WEBEXT.sessions.restore".
func FuncRestore() (fn js.Func[func(sessionId js.String) js.Promise[Session]]) {
	bindings.FuncRestore(
		js.Pointer(&fn),
	)
	return
}

// Restore calls the function "WEBEXT.sessions.restore" directly.
func Restore(sessionId js.String) (ret js.Promise[Session]) {
	bindings.CallRestore(
		js.Pointer(&ret),
		sessionId.Ref(),
	)

	return
}

// TryRestore calls the function "WEBEXT.sessions.restore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestore(sessionId js.String) (ret js.Promise[Session], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestore(
		js.Pointer(&ret), js.Pointer(&exception),
		sessionId.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package debugger

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/debugger/bindings"
)

type Debuggee struct {
	// ExtensionId is "Debuggee.extensionId"
	//
	// Optional
	ExtensionId js.String
	// TabId is "Debuggee.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// TargetId is "Debuggee.targetId"
	//
	// Optional
	TargetId js.String

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Debuggee with all fields set.
func (p Debuggee) FromRef(ref js.Ref) Debuggee {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Debuggee in the application heap.
func (p Debuggee) New() js.Ref {
	return bindings.DebuggeeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Debuggee) UpdateFrom(ref js.Ref) {
	bindings.DebuggeeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Debuggee) Update(ref js.Ref) {
	bindings.DebuggeeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Debuggee) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.TargetId.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.TargetId = p.TargetId.FromRef(js.Undefined)
}

type DetachReason uint32

const (
	_ DetachReason = iota

	DetachReason_TARGET_CLOSED
	DetachReason_CANCELED_BY_USER
)

func (DetachReason) FromRef(str js.Ref) DetachReason {
	return DetachReason(bindings.ConstOfDetachReason(str))
}

func (x DetachReason) String() (string, bool) {
	switch x {
	case DetachReason_TARGET_CLOSED:
		return "target_closed", true
	case DetachReason_CANCELED_BY_USER:
		return "canceled_by_user", true
	default:
		return "", false
	}
}

type TargetInfoType uint32

const (
	_ TargetInfoType = iota

	TargetInfoType_PAGE
	TargetInfoType_BACKGROUND_PAGE
	TargetInfoType_WORKER
	TargetInfoType_OTHER
)

func (TargetInfoType) FromRef(str js.Ref) TargetInfoType {
	return TargetInfoType(bindings.ConstOfTargetInfoType(str))
}

func (x TargetInfoType) String() (string, bool) {
	switch x {
	case TargetInfoType_PAGE:
		return "page", true
	case TargetInfoType_BACKGROUND_PAGE:
		return "background_page", true
	case TargetInfoType_WORKER:
		return "worker", true
	case TargetInfoType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type TargetInfo struct {
	// Attached is "TargetInfo.attached"
	//
	// Required
	Attached bool
	// ExtensionId is "TargetInfo.extensionId"
	//
	// Optional
	ExtensionId js.String
	// FaviconUrl is "TargetInfo.faviconUrl"
	//
	// Optional
	FaviconUrl js.String
	// Id is "TargetInfo.id"
	//
	// Required
	Id js.String
	// TabId is "TargetInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Title is "TargetInfo.title"
	//
	// Required
	Title js.String
	// Type is "TargetInfo.type"
	//
	// Required
	Type TargetInfoType
	// Url is "TargetInfo.url"
	//
	// Required
	Url js.String

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TargetInfo with all fields set.
func (p TargetInfo) FromRef(ref js.Ref) TargetInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TargetInfo in the application heap.
func (p TargetInfo) New() js.Ref {
	return bindings.TargetInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TargetInfo) UpdateFrom(ref js.Ref) {
	bindings.TargetInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TargetInfo) Update(ref js.Ref) {
	bindings.TargetInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TargetInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.FaviconUrl.Ref(),
		p.Id.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.FaviconUrl = p.FaviconUrl.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncAttach returns true if the function "WEBEXT.debugger.attach" exists.
func HasFuncAttach() bool {
	return js.True == bindings.HasFuncAttach()
}

// FuncAttach returns the function "WEBEXT.debugger.attach".
func FuncAttach() (fn js.Func[func(target Debuggee, requiredVersion js.String) js.Promise[js.Void]]) {
	bindings.FuncAttach(
		js.Pointer(&fn),
	)
	return
}

// Attach calls the function "WEBEXT.debugger.attach" directly.
func Attach(target Debuggee, requiredVersion js.String) (ret js.Promise[js.Void]) {
	bindings.CallAttach(
		js.Pointer(&ret),
		js.Pointer(&target),
		requiredVersion.Ref(),
	)

	return
}

// TryAttach calls the function "WEBEXT.debugger.attach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAttach(target Debuggee, requiredVersion js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAttach(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&target),
		requiredVersion.Ref(),
	)

	return
}

// HasFuncDetach returns true if the function "WEBEXT.debugger.detach" exists.
func HasFuncDetach() bool {
	return js.True == bindings.HasFuncDetach()
}

// FuncDetach returns the function "WEBEXT.debugger.detach".
func FuncDetach() (fn js.Func[func(target Debuggee) js.Promise[js.Void]]) {
	bindings.FuncDetach(
		js.Pointer(&fn),
	)
	return
}

// Detach calls the function "WEBEXT.debugger.detach" directly.
func Detach(target Debuggee) (ret js.Promise[js.Void]) {
	bindings.CallDetach(
		js.Pointer(&ret),
		js.Pointer(&target),
	)

	return
}

// TryDetach calls the function "WEBEXT.debugger.detach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDetach(target Debuggee) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDetach(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&target),
	)

	return
}

// HasFuncGetTargets returns true if the function "WEBEXT.debugger.getTargets" exists.
func HasFuncGetTargets() bool {
	return js.True == bindings.HasFuncGetTargets()
}

// FuncGetTargets returns the function "WEBEXT.debugger.getTargets".
func FuncGetTargets() (fn js.Func[func() js.Promise[js.Array[TargetInfo]]]) {
	bindings.FuncGetTargets(
		js.Pointer(&fn),
	)
	return
}

// GetTargets calls the function "WEBEXT.debugger.getTargets" directly.
func GetTargets() (ret js.Promise[js.Array[TargetInfo]]) {
	bindings.CallGetTargets(
		js.Pointer(&ret),
	)

	return
}

// TryGetTargets calls the function "WEBEXT.debugger.getTargets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTargets() (ret js.Promise[js.Array[TargetInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTargets(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnDetachEventCallbackFunc func(this js.Ref, source *Debuggee, reason DetachReason) js.Ref

func (fn OnDetachEventCallbackFunc) Register() js.Func[func(source *Debuggee, reason DetachReason)] {
	return js.RegisterCallback[func(source *Debuggee, reason DetachReason)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDetachEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Debuggee
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		DetachReason(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDetachEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, source *Debuggee, reason DetachReason) js.Ref
	Arg T
}

func (cb *OnDetachEventCallback[T]) Register() js.Func[func(source *Debuggee, reason DetachReason)] {
	return js.RegisterCallback[func(source *Debuggee, reason DetachReason)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDetachEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Debuggee
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		DetachReason(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDetach returns true if the function "WEBEXT.debugger.onDetach.addListener" exists.
func HasFuncOnDetach() bool {
	return js.True == bindings.HasFuncOnDetach()
}

// FuncOnDetach returns the function "WEBEXT.debugger.onDetach.addListener".
func FuncOnDetach() (fn js.Func[func(callback js.Func[func(source *Debuggee, reason DetachReason)])]) {
	bindings.FuncOnDetach(
		js.Pointer(&fn),
	)
	return
}

// OnDetach calls the function "WEBEXT.debugger.onDetach.addListener" directly.
func OnDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret js.Void) {
	bindings.CallOnDetach(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDetach calls the function "WEBEXT.debugger.onDetach.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDetach(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDetach returns true if the function "WEBEXT.debugger.onDetach.removeListener" exists.
func HasFuncOffDetach() bool {
	return js.True == bindings.HasFuncOffDetach()
}

// FuncOffDetach returns the function "WEBEXT.debugger.onDetach.removeListener".
func FuncOffDetach() (fn js.Func[func(callback js.Func[func(source *Debuggee, reason DetachReason)])]) {
	bindings.FuncOffDetach(
		js.Pointer(&fn),
	)
	return
}

// OffDetach calls the function "WEBEXT.debugger.onDetach.removeListener" directly.
func OffDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret js.Void) {
	bindings.CallOffDetach(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDetach calls the function "WEBEXT.debugger.onDetach.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDetach(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDetach returns true if the function "WEBEXT.debugger.onDetach.hasListener" exists.
func HasFuncHasOnDetach() bool {
	return js.True == bindings.HasFuncHasOnDetach()
}

// FuncHasOnDetach returns the function "WEBEXT.debugger.onDetach.hasListener".
func FuncHasOnDetach() (fn js.Func[func(callback js.Func[func(source *Debuggee, reason DetachReason)]) bool]) {
	bindings.FuncHasOnDetach(
		js.Pointer(&fn),
	)
	return
}

// HasOnDetach calls the function "WEBEXT.debugger.onDetach.hasListener" directly.
func HasOnDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret bool) {
	bindings.CallHasOnDetach(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDetach calls the function "WEBEXT.debugger.onDetach.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDetach(callback js.Func[func(source *Debuggee, reason DetachReason)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDetach(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEventEventCallbackFunc func(this js.Ref, source *Debuggee, method js.String, params js.Any) js.Ref

func (fn OnEventEventCallbackFunc) Register() js.Func[func(source *Debuggee, method js.String, params js.Any)] {
	return js.RegisterCallback[func(source *Debuggee, method js.String, params js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Debuggee
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
		js.Any{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, source *Debuggee, method js.String, params js.Any) js.Ref
	Arg T
}

func (cb *OnEventEventCallback[T]) Register() js.Func[func(source *Debuggee, method js.String, params js.Any)] {
	return js.RegisterCallback[func(source *Debuggee, method js.String, params js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Debuggee
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.String{}.FromRef(args[1+1]),
		js.Any{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEvent returns true if the function "WEBEXT.debugger.onEvent.addListener" exists.
func HasFuncOnEvent() bool {
	return js.True == bindings.HasFuncOnEvent()
}

// FuncOnEvent returns the function "WEBEXT.debugger.onEvent.addListener".
func FuncOnEvent() (fn js.Func[func(callback js.Func[func(source *Debuggee, method js.String, params js.Any)])]) {
	bindings.FuncOnEvent(
		js.Pointer(&fn),
	)
	return
}

// OnEvent calls the function "WEBEXT.debugger.onEvent.addListener" directly.
func OnEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret js.Void) {
	bindings.CallOnEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEvent calls the function "WEBEXT.debugger.onEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEvent returns true if the function "WEBEXT.debugger.onEvent.removeListener" exists.
func HasFuncOffEvent() bool {
	return js.True == bindings.HasFuncOffEvent()
}

// FuncOffEvent returns the function "WEBEXT.debugger.onEvent.removeListener".
func FuncOffEvent() (fn js.Func[func(callback js.Func[func(source *Debuggee, method js.String, params js.Any)])]) {
	bindings.FuncOffEvent(
		js.Pointer(&fn),
	)
	return
}

// OffEvent calls the function "WEBEXT.debugger.onEvent.removeListener" directly.
func OffEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret js.Void) {
	bindings.CallOffEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEvent calls the function "WEBEXT.debugger.onEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEvent returns true if the function "WEBEXT.debugger.onEvent.hasListener" exists.
func HasFuncHasOnEvent() bool {
	return js.True == bindings.HasFuncHasOnEvent()
}

// FuncHasOnEvent returns the function "WEBEXT.debugger.onEvent.hasListener".
func FuncHasOnEvent() (fn js.Func[func(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) bool]) {
	bindings.FuncHasOnEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnEvent calls the function "WEBEXT.debugger.onEvent.hasListener" directly.
func HasOnEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret bool) {
	bindings.CallHasOnEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEvent calls the function "WEBEXT.debugger.onEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEvent(callback js.Func[func(source *Debuggee, method js.String, params js.Any)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendCommand returns true if the function "WEBEXT.debugger.sendCommand" exists.
func HasFuncSendCommand() bool {
	return js.True == bindings.HasFuncSendCommand()
}

// FuncSendCommand returns the function "WEBEXT.debugger.sendCommand".
func FuncSendCommand() (fn js.Func[func(target Debuggee, method js.String, commandParams js.Any) js.Promise[js.Any]]) {
	bindings.FuncSendCommand(
		js.Pointer(&fn),
	)
	return
}

// SendCommand calls the function "WEBEXT.debugger.sendCommand" directly.
func SendCommand(target Debuggee, method js.String, commandParams js.Any) (ret js.Promise[js.Any]) {
	bindings.CallSendCommand(
		js.Pointer(&ret),
		js.Pointer(&target),
		method.Ref(),
		commandParams.Ref(),
	)

	return
}

// TrySendCommand calls the function "WEBEXT.debugger.sendCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendCommand(target Debuggee, method js.String, commandParams js.Any) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&target),
		method.Ref(),
		commandParams.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package power

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/power/bindings"
)

type Level uint32

const (
	_ Level = iota

	Level_SYSTEM
	Level_DISPLAY
)

func (Level) FromRef(str js.Ref) Level {
	return Level(bindings.ConstOfLevel(str))
}

func (x Level) String() (string, bool) {
	switch x {
	case Level_SYSTEM:
		return "system", true
	case Level_DISPLAY:
		return "display", true
	default:
		return "", false
	}
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncReleaseKeepAwake returns true if the function "WEBEXT.power.releaseKeepAwake" exists.
func HasFuncReleaseKeepAwake() bool {
	return js.True == bindings.HasFuncReleaseKeepAwake()
}

// FuncReleaseKeepAwake returns the function "WEBEXT.power.releaseKeepAwake".
func FuncReleaseKeepAwake() (fn js.Func[func()]) {
	bindings.FuncReleaseKeepAwake(
		js.Pointer(&fn),
	)
	return
}

// ReleaseKeepAwake calls the function "WEBEXT.power.releaseKeepAwake" directly.
func ReleaseKeepAwake() (ret js.Void) {
	bindings.CallReleaseKeepAwake(
		js.Pointer(&ret),
	)

	return
}

// TryReleaseKeepAwake calls the function "WEBEXT.power.releaseKeepAwake"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReleaseKeepAwake() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReleaseKeepAwake(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportActivity returns true if the function "WEBEXT.power.reportActivity" exists.
func HasFuncReportActivity() bool {
	return js.True == bindings.HasFuncReportActivity()
}

// FuncReportActivity returns the function "WEBEXT.power.reportActivity".
func FuncReportActivity() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncReportActivity(
		js.Pointer(&fn),
	)
	return
}

// ReportActivity calls the function "WEBEXT.power.reportActivity" directly.
func ReportActivity() (ret js.Promise[js.Void]) {
	bindings.CallReportActivity(
		js.Pointer(&ret),
	)

	return
}

// TryReportActivity calls the function "WEBEXT.power.reportActivity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportActivity() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportActivity(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestKeepAwake returns true if the function "WEBEXT.power.requestKeepAwake" exists.
func HasFuncRequestKeepAwake() bool {
	return js.True == bindings.HasFuncRequestKeepAwake()
}

// FuncRequestKeepAwake returns the function "WEBEXT.power.requestKeepAwake".
func FuncRequestKeepAwake() (fn js.Func[func(level Level)]) {
	bindings.FuncRequestKeepAwake(
		js.Pointer(&fn),
	)
	return
}

// RequestKeepAwake calls the function "WEBEXT.power.requestKeepAwake" directly.
func RequestKeepAwake(level Level) (ret js.Void) {
	bindings.CallRequestKeepAwake(
		js.Pointer(&ret),
		uint32(level),
	)

	return
}

// TryRequestKeepAwake calls the function "WEBEXT.power.requestKeepAwake"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestKeepAwake(level Level) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestKeepAwake(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(level),
	)

	return
}

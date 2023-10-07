// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package app

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/app/bindings"
)

type DOMWindow struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMWindow with all fields set.
func (p DOMWindow) FromRef(ref js.Ref) DOMWindow {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMWindow in the application heap.
func (p DOMWindow) New() js.Ref {
	return bindings.DOMWindowJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DOMWindow) UpdateFrom(ref js.Ref) {
	bindings.DOMWindowJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMWindow) Update(ref js.Ref) {
	bindings.DOMWindowJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMWindow) FreeMembers(recursive bool) {
}

type Details struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Details with all fields set.
func (p Details) FromRef(ref js.Ref) Details {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Details in the application heap.
func (p Details) New() js.Ref {
	return bindings.DetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Details) UpdateFrom(ref js.Ref) {
	bindings.DetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Details) Update(ref js.Ref) {
	bindings.DetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Details) FreeMembers(recursive bool) {
}

type InstallStateType uint32

const (
	_ InstallStateType = iota

	InstallStateType_NOT_INSTALLED
	InstallStateType_INSTALLED
	InstallStateType_DISABLED
)

func (InstallStateType) FromRef(str js.Ref) InstallStateType {
	return InstallStateType(bindings.ConstOfInstallStateType(str))
}

func (x InstallStateType) String() (string, bool) {
	switch x {
	case InstallStateType_NOT_INSTALLED:
		return "not_installed", true
	case InstallStateType_INSTALLED:
		return "installed", true
	case InstallStateType_DISABLED:
		return "disabled", true
	default:
		return "", false
	}
}

type RunningStateType uint32

const (
	_ RunningStateType = iota

	RunningStateType_RUNNING
	RunningStateType_CANNOT_RUN
	RunningStateType_READY_TO_RUN
)

func (RunningStateType) FromRef(str js.Ref) RunningStateType {
	return RunningStateType(bindings.ConstOfRunningStateType(str))
}

func (x RunningStateType) String() (string, bool) {
	switch x {
	case RunningStateType_RUNNING:
		return "running", true
	case RunningStateType_CANNOT_RUN:
		return "cannot_run", true
	case RunningStateType_READY_TO_RUN:
		return "ready_to_run", true
	default:
		return "", false
	}
}

// HasFuncGetDetails returns true if the function "WEBEXT.app.getDetails" exists.
func HasFuncGetDetails() bool {
	return js.True == bindings.HasFuncGetDetails()
}

// FuncGetDetails returns the function "WEBEXT.app.getDetails".
func FuncGetDetails() (fn js.Func[func() Details]) {
	bindings.FuncGetDetails(
		js.Pointer(&fn),
	)
	return
}

// GetDetails calls the function "WEBEXT.app.getDetails" directly.
func GetDetails() (ret Details) {
	bindings.CallGetDetails(
		js.Pointer(&ret),
	)

	return
}

// TryGetDetails calls the function "WEBEXT.app.getDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDetails() (ret Details, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDetails(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetIsInstalled returns true if the function "WEBEXT.app.getIsInstalled" exists.
func HasFuncGetIsInstalled() bool {
	return js.True == bindings.HasFuncGetIsInstalled()
}

// FuncGetIsInstalled returns the function "WEBEXT.app.getIsInstalled".
func FuncGetIsInstalled() (fn js.Func[func() bool]) {
	bindings.FuncGetIsInstalled(
		js.Pointer(&fn),
	)
	return
}

// GetIsInstalled calls the function "WEBEXT.app.getIsInstalled" directly.
func GetIsInstalled() (ret bool) {
	bindings.CallGetIsInstalled(
		js.Pointer(&ret),
	)

	return
}

// TryGetIsInstalled calls the function "WEBEXT.app.getIsInstalled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIsInstalled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIsInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInstallState returns true if the function "WEBEXT.app.installState" exists.
func HasFuncInstallState() bool {
	return js.True == bindings.HasFuncInstallState()
}

// FuncInstallState returns the function "WEBEXT.app.installState".
func FuncInstallState() (fn js.Func[func() js.Promise[InstallStateType]]) {
	bindings.FuncInstallState(
		js.Pointer(&fn),
	)
	return
}

// InstallState calls the function "WEBEXT.app.installState" directly.
func InstallState() (ret js.Promise[InstallStateType]) {
	bindings.CallInstallState(
		js.Pointer(&ret),
	)

	return
}

// TryInstallState calls the function "WEBEXT.app.installState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallState() (ret js.Promise[InstallStateType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunningState returns true if the function "WEBEXT.app.runningState" exists.
func HasFuncRunningState() bool {
	return js.True == bindings.HasFuncRunningState()
}

// FuncRunningState returns the function "WEBEXT.app.runningState".
func FuncRunningState() (fn js.Func[func() RunningStateType]) {
	bindings.FuncRunningState(
		js.Pointer(&fn),
	)
	return
}

// RunningState calls the function "WEBEXT.app.runningState" directly.
func RunningState() (ret RunningStateType) {
	bindings.CallRunningState(
		js.Pointer(&ret),
	)

	return
}

// TryRunningState calls the function "WEBEXT.app.runningState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunningState() (ret RunningStateType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunningState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

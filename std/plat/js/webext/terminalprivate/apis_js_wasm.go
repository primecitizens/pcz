// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package terminalprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/terminalprivate/bindings"
)

type GetOSInfoReturnType struct {
	// AlternativeEmulator is "GetOSInfoReturnType.alternative_emulator"
	//
	// Required
	AlternativeEmulator bool
	// Tast is "GetOSInfoReturnType.tast"
	//
	// Required
	Tast bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetOSInfoReturnType with all fields set.
func (p GetOSInfoReturnType) FromRef(ref js.Ref) GetOSInfoReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetOSInfoReturnType in the application heap.
func (p GetOSInfoReturnType) New() js.Ref {
	return bindings.GetOSInfoReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetOSInfoReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetOSInfoReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetOSInfoReturnType) Update(ref js.Ref) {
	bindings.GetOSInfoReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetOSInfoReturnType) FreeMembers(recursive bool) {
}

type OpenWindowArgData struct {
	// AsTab is "OpenWindowArgData.asTab"
	//
	// Optional
	//
	// NOTE: FFI_USE_AsTab MUST be set to true to make this field effective.
	AsTab bool
	// Url is "OpenWindowArgData.url"
	//
	// Optional
	Url js.String

	FFI_USE_AsTab bool // for AsTab.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenWindowArgData with all fields set.
func (p OpenWindowArgData) FromRef(ref js.Ref) OpenWindowArgData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenWindowArgData in the application heap.
func (p OpenWindowArgData) New() js.Ref {
	return bindings.OpenWindowArgDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenWindowArgData) UpdateFrom(ref js.Ref) {
	bindings.OpenWindowArgDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenWindowArgData) Update(ref js.Ref) {
	bindings.OpenWindowArgDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenWindowArgData) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OutputType uint32

const (
	_ OutputType = iota

	OutputType_STDOUT
	OutputType_STDERR
	OutputType_EXIT
)

func (OutputType) FromRef(str js.Ref) OutputType {
	return OutputType(bindings.ConstOfOutputType(str))
}

func (x OutputType) String() (string, bool) {
	switch x {
	case OutputType_STDOUT:
		return "stdout", true
	case OutputType_STDERR:
		return "stderr", true
	case OutputType_EXIT:
		return "exit", true
	default:
		return "", false
	}
}

// HasFuncAckOutput returns true if the function "WEBEXT.terminalPrivate.ackOutput" exists.
func HasFuncAckOutput() bool {
	return js.True == bindings.HasFuncAckOutput()
}

// FuncAckOutput returns the function "WEBEXT.terminalPrivate.ackOutput".
func FuncAckOutput() (fn js.Func[func(id js.String)]) {
	bindings.FuncAckOutput(
		js.Pointer(&fn),
	)
	return
}

// AckOutput calls the function "WEBEXT.terminalPrivate.ackOutput" directly.
func AckOutput(id js.String) (ret js.Void) {
	bindings.CallAckOutput(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryAckOutput calls the function "WEBEXT.terminalPrivate.ackOutput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAckOutput(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAckOutput(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncCloseTerminalProcess returns true if the function "WEBEXT.terminalPrivate.closeTerminalProcess" exists.
func HasFuncCloseTerminalProcess() bool {
	return js.True == bindings.HasFuncCloseTerminalProcess()
}

// FuncCloseTerminalProcess returns the function "WEBEXT.terminalPrivate.closeTerminalProcess".
func FuncCloseTerminalProcess() (fn js.Func[func(id js.String) js.Promise[js.Boolean]]) {
	bindings.FuncCloseTerminalProcess(
		js.Pointer(&fn),
	)
	return
}

// CloseTerminalProcess calls the function "WEBEXT.terminalPrivate.closeTerminalProcess" directly.
func CloseTerminalProcess(id js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallCloseTerminalProcess(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryCloseTerminalProcess calls the function "WEBEXT.terminalPrivate.closeTerminalProcess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseTerminalProcess(id js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseTerminalProcess(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetOSInfo returns true if the function "WEBEXT.terminalPrivate.getOSInfo" exists.
func HasFuncGetOSInfo() bool {
	return js.True == bindings.HasFuncGetOSInfo()
}

// FuncGetOSInfo returns the function "WEBEXT.terminalPrivate.getOSInfo".
func FuncGetOSInfo() (fn js.Func[func() js.Promise[GetOSInfoReturnType]]) {
	bindings.FuncGetOSInfo(
		js.Pointer(&fn),
	)
	return
}

// GetOSInfo calls the function "WEBEXT.terminalPrivate.getOSInfo" directly.
func GetOSInfo() (ret js.Promise[GetOSInfoReturnType]) {
	bindings.CallGetOSInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetOSInfo calls the function "WEBEXT.terminalPrivate.getOSInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOSInfo() (ret js.Promise[GetOSInfoReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOSInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPrefs returns true if the function "WEBEXT.terminalPrivate.getPrefs" exists.
func HasFuncGetPrefs() bool {
	return js.True == bindings.HasFuncGetPrefs()
}

// FuncGetPrefs returns the function "WEBEXT.terminalPrivate.getPrefs".
func FuncGetPrefs() (fn js.Func[func(paths js.Array[js.String]) js.Promise[js.Any]]) {
	bindings.FuncGetPrefs(
		js.Pointer(&fn),
	)
	return
}

// GetPrefs calls the function "WEBEXT.terminalPrivate.getPrefs" directly.
func GetPrefs(paths js.Array[js.String]) (ret js.Promise[js.Any]) {
	bindings.CallGetPrefs(
		js.Pointer(&ret),
		paths.Ref(),
	)

	return
}

// TryGetPrefs calls the function "WEBEXT.terminalPrivate.getPrefs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrefs(paths js.Array[js.String]) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrefs(
		js.Pointer(&ret), js.Pointer(&exception),
		paths.Ref(),
	)

	return
}

type OnPrefChangedEventCallbackFunc func(this js.Ref, prefs js.Any) js.Ref

func (fn OnPrefChangedEventCallbackFunc) Register() js.Func[func(prefs js.Any)] {
	return js.RegisterCallback[func(prefs js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrefChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPrefChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, prefs js.Any) js.Ref
	Arg T
}

func (cb *OnPrefChangedEventCallback[T]) Register() js.Func[func(prefs js.Any)] {
	return js.RegisterCallback[func(prefs js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrefChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPrefChanged returns true if the function "WEBEXT.terminalPrivate.onPrefChanged.addListener" exists.
func HasFuncOnPrefChanged() bool {
	return js.True == bindings.HasFuncOnPrefChanged()
}

// FuncOnPrefChanged returns the function "WEBEXT.terminalPrivate.onPrefChanged.addListener".
func FuncOnPrefChanged() (fn js.Func[func(callback js.Func[func(prefs js.Any)])]) {
	bindings.FuncOnPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.addListener" directly.
func OnPrefChanged(callback js.Func[func(prefs js.Any)]) (ret js.Void) {
	bindings.CallOnPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPrefChanged(callback js.Func[func(prefs js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPrefChanged returns true if the function "WEBEXT.terminalPrivate.onPrefChanged.removeListener" exists.
func HasFuncOffPrefChanged() bool {
	return js.True == bindings.HasFuncOffPrefChanged()
}

// FuncOffPrefChanged returns the function "WEBEXT.terminalPrivate.onPrefChanged.removeListener".
func FuncOffPrefChanged() (fn js.Func[func(callback js.Func[func(prefs js.Any)])]) {
	bindings.FuncOffPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.removeListener" directly.
func OffPrefChanged(callback js.Func[func(prefs js.Any)]) (ret js.Void) {
	bindings.CallOffPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPrefChanged(callback js.Func[func(prefs js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPrefChanged returns true if the function "WEBEXT.terminalPrivate.onPrefChanged.hasListener" exists.
func HasFuncHasOnPrefChanged() bool {
	return js.True == bindings.HasFuncHasOnPrefChanged()
}

// FuncHasOnPrefChanged returns the function "WEBEXT.terminalPrivate.onPrefChanged.hasListener".
func FuncHasOnPrefChanged() (fn js.Func[func(callback js.Func[func(prefs js.Any)]) bool]) {
	bindings.FuncHasOnPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.hasListener" directly.
func HasOnPrefChanged(callback js.Func[func(prefs js.Any)]) (ret bool) {
	bindings.CallHasOnPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPrefChanged calls the function "WEBEXT.terminalPrivate.onPrefChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPrefChanged(callback js.Func[func(prefs js.Any)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnProcessOutputEventCallbackFunc func(this js.Ref, id js.String, typ OutputType, data js.TypedArray[uint8]) js.Ref

func (fn OnProcessOutputEventCallbackFunc) Register() js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])] {
	return js.RegisterCallback[func(id js.String, typ OutputType, data js.TypedArray[uint8])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnProcessOutputEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		OutputType(0).FromRef(args[1+1]),
		js.TypedArray[uint8]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnProcessOutputEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, typ OutputType, data js.TypedArray[uint8]) js.Ref
	Arg T
}

func (cb *OnProcessOutputEventCallback[T]) Register() js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])] {
	return js.RegisterCallback[func(id js.String, typ OutputType, data js.TypedArray[uint8])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnProcessOutputEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		OutputType(0).FromRef(args[1+1]),
		js.TypedArray[uint8]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnProcessOutput returns true if the function "WEBEXT.terminalPrivate.onProcessOutput.addListener" exists.
func HasFuncOnProcessOutput() bool {
	return js.True == bindings.HasFuncOnProcessOutput()
}

// FuncOnProcessOutput returns the function "WEBEXT.terminalPrivate.onProcessOutput.addListener".
func FuncOnProcessOutput() (fn js.Func[func(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])])]) {
	bindings.FuncOnProcessOutput(
		js.Pointer(&fn),
	)
	return
}

// OnProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.addListener" directly.
func OnProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret js.Void) {
	bindings.CallOnProcessOutput(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnProcessOutput(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffProcessOutput returns true if the function "WEBEXT.terminalPrivate.onProcessOutput.removeListener" exists.
func HasFuncOffProcessOutput() bool {
	return js.True == bindings.HasFuncOffProcessOutput()
}

// FuncOffProcessOutput returns the function "WEBEXT.terminalPrivate.onProcessOutput.removeListener".
func FuncOffProcessOutput() (fn js.Func[func(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])])]) {
	bindings.FuncOffProcessOutput(
		js.Pointer(&fn),
	)
	return
}

// OffProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.removeListener" directly.
func OffProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret js.Void) {
	bindings.CallOffProcessOutput(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffProcessOutput(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnProcessOutput returns true if the function "WEBEXT.terminalPrivate.onProcessOutput.hasListener" exists.
func HasFuncHasOnProcessOutput() bool {
	return js.True == bindings.HasFuncHasOnProcessOutput()
}

// FuncHasOnProcessOutput returns the function "WEBEXT.terminalPrivate.onProcessOutput.hasListener".
func FuncHasOnProcessOutput() (fn js.Func[func(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) bool]) {
	bindings.FuncHasOnProcessOutput(
		js.Pointer(&fn),
	)
	return
}

// HasOnProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.hasListener" directly.
func HasOnProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret bool) {
	bindings.CallHasOnProcessOutput(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnProcessOutput calls the function "WEBEXT.terminalPrivate.onProcessOutput.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnProcessOutput(callback js.Func[func(id js.String, typ OutputType, data js.TypedArray[uint8])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnProcessOutput(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOnTerminalResize returns true if the function "WEBEXT.terminalPrivate.onTerminalResize" exists.
func HasFuncOnTerminalResize() bool {
	return js.True == bindings.HasFuncOnTerminalResize()
}

// FuncOnTerminalResize returns the function "WEBEXT.terminalPrivate.onTerminalResize".
func FuncOnTerminalResize() (fn js.Func[func(id js.String, width int64, height int64) js.Promise[js.Boolean]]) {
	bindings.FuncOnTerminalResize(
		js.Pointer(&fn),
	)
	return
}

// OnTerminalResize calls the function "WEBEXT.terminalPrivate.onTerminalResize" directly.
func OnTerminalResize(id js.String, width int64, height int64) (ret js.Promise[js.Boolean]) {
	bindings.CallOnTerminalResize(
		js.Pointer(&ret),
		id.Ref(),
		float64(width),
		float64(height),
	)

	return
}

// TryOnTerminalResize calls the function "WEBEXT.terminalPrivate.onTerminalResize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTerminalResize(id js.String, width int64, height int64) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTerminalResize(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		float64(width),
		float64(height),
	)

	return
}

// HasFuncOpenOptionsPage returns true if the function "WEBEXT.terminalPrivate.openOptionsPage" exists.
func HasFuncOpenOptionsPage() bool {
	return js.True == bindings.HasFuncOpenOptionsPage()
}

// FuncOpenOptionsPage returns the function "WEBEXT.terminalPrivate.openOptionsPage".
func FuncOpenOptionsPage() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncOpenOptionsPage(
		js.Pointer(&fn),
	)
	return
}

// OpenOptionsPage calls the function "WEBEXT.terminalPrivate.openOptionsPage" directly.
func OpenOptionsPage() (ret js.Promise[js.Void]) {
	bindings.CallOpenOptionsPage(
		js.Pointer(&ret),
	)

	return
}

// TryOpenOptionsPage calls the function "WEBEXT.terminalPrivate.openOptionsPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenOptionsPage() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenOptionsPage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenSettingsSubpage returns true if the function "WEBEXT.terminalPrivate.openSettingsSubpage" exists.
func HasFuncOpenSettingsSubpage() bool {
	return js.True == bindings.HasFuncOpenSettingsSubpage()
}

// FuncOpenSettingsSubpage returns the function "WEBEXT.terminalPrivate.openSettingsSubpage".
func FuncOpenSettingsSubpage() (fn js.Func[func(subpage js.String) js.Promise[js.Void]]) {
	bindings.FuncOpenSettingsSubpage(
		js.Pointer(&fn),
	)
	return
}

// OpenSettingsSubpage calls the function "WEBEXT.terminalPrivate.openSettingsSubpage" directly.
func OpenSettingsSubpage(subpage js.String) (ret js.Promise[js.Void]) {
	bindings.CallOpenSettingsSubpage(
		js.Pointer(&ret),
		subpage.Ref(),
	)

	return
}

// TryOpenSettingsSubpage calls the function "WEBEXT.terminalPrivate.openSettingsSubpage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSettingsSubpage(subpage js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSettingsSubpage(
		js.Pointer(&ret), js.Pointer(&exception),
		subpage.Ref(),
	)

	return
}

// HasFuncOpenTerminalProcess returns true if the function "WEBEXT.terminalPrivate.openTerminalProcess" exists.
func HasFuncOpenTerminalProcess() bool {
	return js.True == bindings.HasFuncOpenTerminalProcess()
}

// FuncOpenTerminalProcess returns the function "WEBEXT.terminalPrivate.openTerminalProcess".
func FuncOpenTerminalProcess() (fn js.Func[func(processName js.String, args js.Array[js.String]) js.Promise[js.String]]) {
	bindings.FuncOpenTerminalProcess(
		js.Pointer(&fn),
	)
	return
}

// OpenTerminalProcess calls the function "WEBEXT.terminalPrivate.openTerminalProcess" directly.
func OpenTerminalProcess(processName js.String, args js.Array[js.String]) (ret js.Promise[js.String]) {
	bindings.CallOpenTerminalProcess(
		js.Pointer(&ret),
		processName.Ref(),
		args.Ref(),
	)

	return
}

// TryOpenTerminalProcess calls the function "WEBEXT.terminalPrivate.openTerminalProcess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenTerminalProcess(processName js.String, args js.Array[js.String]) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenTerminalProcess(
		js.Pointer(&ret), js.Pointer(&exception),
		processName.Ref(),
		args.Ref(),
	)

	return
}

// HasFuncOpenVmshellProcess returns true if the function "WEBEXT.terminalPrivate.openVmshellProcess" exists.
func HasFuncOpenVmshellProcess() bool {
	return js.True == bindings.HasFuncOpenVmshellProcess()
}

// FuncOpenVmshellProcess returns the function "WEBEXT.terminalPrivate.openVmshellProcess".
func FuncOpenVmshellProcess() (fn js.Func[func(args js.Array[js.String]) js.Promise[js.String]]) {
	bindings.FuncOpenVmshellProcess(
		js.Pointer(&fn),
	)
	return
}

// OpenVmshellProcess calls the function "WEBEXT.terminalPrivate.openVmshellProcess" directly.
func OpenVmshellProcess(args js.Array[js.String]) (ret js.Promise[js.String]) {
	bindings.CallOpenVmshellProcess(
		js.Pointer(&ret),
		args.Ref(),
	)

	return
}

// TryOpenVmshellProcess calls the function "WEBEXT.terminalPrivate.openVmshellProcess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenVmshellProcess(args js.Array[js.String]) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenVmshellProcess(
		js.Pointer(&ret), js.Pointer(&exception),
		args.Ref(),
	)

	return
}

// HasFuncOpenWindow returns true if the function "WEBEXT.terminalPrivate.openWindow" exists.
func HasFuncOpenWindow() bool {
	return js.True == bindings.HasFuncOpenWindow()
}

// FuncOpenWindow returns the function "WEBEXT.terminalPrivate.openWindow".
func FuncOpenWindow() (fn js.Func[func(data OpenWindowArgData)]) {
	bindings.FuncOpenWindow(
		js.Pointer(&fn),
	)
	return
}

// OpenWindow calls the function "WEBEXT.terminalPrivate.openWindow" directly.
func OpenWindow(data OpenWindowArgData) (ret js.Void) {
	bindings.CallOpenWindow(
		js.Pointer(&ret),
		js.Pointer(&data),
	)

	return
}

// TryOpenWindow calls the function "WEBEXT.terminalPrivate.openWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenWindow(data OpenWindowArgData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&data),
	)

	return
}

// HasFuncSendInput returns true if the function "WEBEXT.terminalPrivate.sendInput" exists.
func HasFuncSendInput() bool {
	return js.True == bindings.HasFuncSendInput()
}

// FuncSendInput returns the function "WEBEXT.terminalPrivate.sendInput".
func FuncSendInput() (fn js.Func[func(id js.String, input js.String) js.Promise[js.Boolean]]) {
	bindings.FuncSendInput(
		js.Pointer(&fn),
	)
	return
}

// SendInput calls the function "WEBEXT.terminalPrivate.sendInput" directly.
func SendInput(id js.String, input js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallSendInput(
		js.Pointer(&ret),
		id.Ref(),
		input.Ref(),
	)

	return
}

// TrySendInput calls the function "WEBEXT.terminalPrivate.sendInput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendInput(id js.String, input js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendInput(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		input.Ref(),
	)

	return
}

// HasFuncSetPrefs returns true if the function "WEBEXT.terminalPrivate.setPrefs" exists.
func HasFuncSetPrefs() bool {
	return js.True == bindings.HasFuncSetPrefs()
}

// FuncSetPrefs returns the function "WEBEXT.terminalPrivate.setPrefs".
func FuncSetPrefs() (fn js.Func[func(prefs js.Any) js.Promise[js.Void]]) {
	bindings.FuncSetPrefs(
		js.Pointer(&fn),
	)
	return
}

// SetPrefs calls the function "WEBEXT.terminalPrivate.setPrefs" directly.
func SetPrefs(prefs js.Any) (ret js.Promise[js.Void]) {
	bindings.CallSetPrefs(
		js.Pointer(&ret),
		prefs.Ref(),
	)

	return
}

// TrySetPrefs calls the function "WEBEXT.terminalPrivate.setPrefs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPrefs(prefs js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPrefs(
		js.Pointer(&ret), js.Pointer(&exception),
		prefs.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package processes

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/processes/bindings"
)

type Cache struct {
	// Size is "Cache.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size float64
	// LiveSize is "Cache.liveSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_LiveSize MUST be set to true to make this field effective.
	LiveSize float64

	FFI_USE_Size     bool // for Size.
	FFI_USE_LiveSize bool // for LiveSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Cache with all fields set.
func (p Cache) FromRef(ref js.Ref) Cache {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Cache in the application heap.
func (p Cache) New() js.Ref {
	return bindings.CacheJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Cache) UpdateFrom(ref js.Ref) {
	bindings.CacheJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Cache) Update(ref js.Ref) {
	bindings.CacheJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Cache) FreeMembers(recursive bool) {
}

type GetProcessIdForTabCallbackFunc func(this js.Ref, processId int32) js.Ref

func (fn GetProcessIdForTabCallbackFunc) Register() js.Func[func(processId int32)] {
	return js.RegisterCallback[func(processId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProcessIdForTabCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProcessIdForTabCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processId int32) js.Ref
	Arg T
}

func (cb *GetProcessIdForTabCallback[T]) Register() js.Func[func(processId int32)] {
	return js.RegisterCallback[func(processId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProcessIdForTabCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProcessInfoCallbackFunc func(this js.Ref, processes js.Object) js.Ref

func (fn GetProcessInfoCallbackFunc) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProcessInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProcessInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processes js.Object) js.Ref
	Arg T
}

func (cb *GetProcessInfoCallback[T]) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProcessInfoCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProcessType uint32

const (
	_ ProcessType = iota

	ProcessType_BROWSER
	ProcessType_RENDERER
	ProcessType_EXTENSION
	ProcessType_NOTIFICATION
	ProcessType_PLUGIN
	ProcessType_WORKER
	ProcessType_NACL
	ProcessType_SERVICE_WORKER
	ProcessType_UTILITY
	ProcessType_GPU
	ProcessType_OTHER
)

func (ProcessType) FromRef(str js.Ref) ProcessType {
	return ProcessType(bindings.ConstOfProcessType(str))
}

func (x ProcessType) String() (string, bool) {
	switch x {
	case ProcessType_BROWSER:
		return "browser", true
	case ProcessType_RENDERER:
		return "renderer", true
	case ProcessType_EXTENSION:
		return "extension", true
	case ProcessType_NOTIFICATION:
		return "notification", true
	case ProcessType_PLUGIN:
		return "plugin", true
	case ProcessType_WORKER:
		return "worker", true
	case ProcessType_NACL:
		return "nacl", true
	case ProcessType_SERVICE_WORKER:
		return "service_worker", true
	case ProcessType_UTILITY:
		return "utility", true
	case ProcessType_GPU:
		return "gpu", true
	case ProcessType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type TaskInfo struct {
	// Title is "TaskInfo.title"
	//
	// Optional
	Title js.String
	// TabId is "TaskInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TaskInfo with all fields set.
func (p TaskInfo) FromRef(ref js.Ref) TaskInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TaskInfo in the application heap.
func (p TaskInfo) New() js.Ref {
	return bindings.TaskInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TaskInfo) UpdateFrom(ref js.Ref) {
	bindings.TaskInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TaskInfo) Update(ref js.Ref) {
	bindings.TaskInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TaskInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
}

type Process struct {
	// Id is "Process.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// OsProcessId is "Process.osProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_OsProcessId MUST be set to true to make this field effective.
	OsProcessId int32
	// Type is "Process.type"
	//
	// Optional
	Type ProcessType
	// Profile is "Process.profile"
	//
	// Optional
	Profile js.String
	// NaclDebugPort is "Process.naclDebugPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_NaclDebugPort MUST be set to true to make this field effective.
	NaclDebugPort int32
	// Tasks is "Process.tasks"
	//
	// Optional
	Tasks js.Array[TaskInfo]
	// Cpu is "Process.cpu"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cpu MUST be set to true to make this field effective.
	Cpu float64
	// Network is "Process.network"
	//
	// Optional
	//
	// NOTE: FFI_USE_Network MUST be set to true to make this field effective.
	Network float64
	// PrivateMemory is "Process.privateMemory"
	//
	// Optional
	//
	// NOTE: FFI_USE_PrivateMemory MUST be set to true to make this field effective.
	PrivateMemory float64
	// JsMemoryAllocated is "Process.jsMemoryAllocated"
	//
	// Optional
	//
	// NOTE: FFI_USE_JsMemoryAllocated MUST be set to true to make this field effective.
	JsMemoryAllocated float64
	// JsMemoryUsed is "Process.jsMemoryUsed"
	//
	// Optional
	//
	// NOTE: FFI_USE_JsMemoryUsed MUST be set to true to make this field effective.
	JsMemoryUsed float64
	// SqliteMemory is "Process.sqliteMemory"
	//
	// Optional
	//
	// NOTE: FFI_USE_SqliteMemory MUST be set to true to make this field effective.
	SqliteMemory float64
	// ImageCache is "Process.imageCache"
	//
	// Optional
	//
	// NOTE: ImageCache.FFI_USE MUST be set to true to get ImageCache used.
	ImageCache Cache
	// ScriptCache is "Process.scriptCache"
	//
	// Optional
	//
	// NOTE: ScriptCache.FFI_USE MUST be set to true to get ScriptCache used.
	ScriptCache Cache
	// CssCache is "Process.cssCache"
	//
	// Optional
	//
	// NOTE: CssCache.FFI_USE MUST be set to true to get CssCache used.
	CssCache Cache

	FFI_USE_Id                bool // for Id.
	FFI_USE_OsProcessId       bool // for OsProcessId.
	FFI_USE_NaclDebugPort     bool // for NaclDebugPort.
	FFI_USE_Cpu               bool // for Cpu.
	FFI_USE_Network           bool // for Network.
	FFI_USE_PrivateMemory     bool // for PrivateMemory.
	FFI_USE_JsMemoryAllocated bool // for JsMemoryAllocated.
	FFI_USE_JsMemoryUsed      bool // for JsMemoryUsed.
	FFI_USE_SqliteMemory      bool // for SqliteMemory.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Process with all fields set.
func (p Process) FromRef(ref js.Ref) Process {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Process in the application heap.
func (p Process) New() js.Ref {
	return bindings.ProcessJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Process) UpdateFrom(ref js.Ref) {
	bindings.ProcessJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Process) Update(ref js.Ref) {
	bindings.ProcessJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Process) FreeMembers(recursive bool) {
	js.Free(
		p.Profile.Ref(),
		p.Tasks.Ref(),
	)
	p.Profile = p.Profile.FromRef(js.Undefined)
	p.Tasks = p.Tasks.FromRef(js.Undefined)
	if recursive {
		p.ImageCache.FreeMembers(true)
		p.ScriptCache.FreeMembers(true)
		p.CssCache.FreeMembers(true)
	}
}

type TerminateCallbackFunc func(this js.Ref, didTerminate bool) js.Ref

func (fn TerminateCallbackFunc) Register() js.Func[func(didTerminate bool)] {
	return js.RegisterCallback[func(didTerminate bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TerminateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TerminateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, didTerminate bool) js.Ref
	Arg T
}

func (cb *TerminateCallback[T]) Register() js.Func[func(didTerminate bool)] {
	return js.RegisterCallback[func(didTerminate bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TerminateCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetProcessIdForTab returns true if the function "WEBEXT.processes.getProcessIdForTab" exists.
func HasFuncGetProcessIdForTab() bool {
	return js.True == bindings.HasFuncGetProcessIdForTab()
}

// FuncGetProcessIdForTab returns the function "WEBEXT.processes.getProcessIdForTab".
func FuncGetProcessIdForTab() (fn js.Func[func(tabId int32) js.Promise[js.Number[int32]]]) {
	bindings.FuncGetProcessIdForTab(
		js.Pointer(&fn),
	)
	return
}

// GetProcessIdForTab calls the function "WEBEXT.processes.getProcessIdForTab" directly.
func GetProcessIdForTab(tabId int32) (ret js.Promise[js.Number[int32]]) {
	bindings.CallGetProcessIdForTab(
		js.Pointer(&ret),
		int32(tabId),
	)

	return
}

// TryGetProcessIdForTab calls the function "WEBEXT.processes.getProcessIdForTab"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProcessIdForTab(tabId int32) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProcessIdForTab(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(tabId),
	)

	return
}

type OneOf_Int32_ArrayInt32 struct {
	ref js.Ref
}

func (x OneOf_Int32_ArrayInt32) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Int32_ArrayInt32) Free() {
	x.ref.Free()
}

func (x OneOf_Int32_ArrayInt32) FromRef(ref js.Ref) OneOf_Int32_ArrayInt32 {
	return OneOf_Int32_ArrayInt32{
		ref: ref,
	}
}

func (x OneOf_Int32_ArrayInt32) Int32() int32 {
	return js.Number[int32]{}.FromRef(x.ref).Get()
}

func (x OneOf_Int32_ArrayInt32) ArrayInt32() js.Array[int32] {
	return js.Array[int32]{}.FromRef(x.ref)
}

// HasFuncGetProcessInfo returns true if the function "WEBEXT.processes.getProcessInfo" exists.
func HasFuncGetProcessInfo() bool {
	return js.True == bindings.HasFuncGetProcessInfo()
}

// FuncGetProcessInfo returns the function "WEBEXT.processes.getProcessInfo".
func FuncGetProcessInfo() (fn js.Func[func(processIds OneOf_Int32_ArrayInt32, includeMemory bool) js.Promise[js.Object]]) {
	bindings.FuncGetProcessInfo(
		js.Pointer(&fn),
	)
	return
}

// GetProcessInfo calls the function "WEBEXT.processes.getProcessInfo" directly.
func GetProcessInfo(processIds OneOf_Int32_ArrayInt32, includeMemory bool) (ret js.Promise[js.Object]) {
	bindings.CallGetProcessInfo(
		js.Pointer(&ret),
		processIds.Ref(),
		js.Bool(bool(includeMemory)),
	)

	return
}

// TryGetProcessInfo calls the function "WEBEXT.processes.getProcessInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProcessInfo(processIds OneOf_Int32_ArrayInt32, includeMemory bool) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProcessInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		processIds.Ref(),
		js.Bool(bool(includeMemory)),
	)

	return
}

type OnCreatedEventCallbackFunc func(this js.Ref, process *Process) js.Ref

func (fn OnCreatedEventCallbackFunc) Register() js.Func[func(process *Process)] {
	return js.RegisterCallback[func(process *Process)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Process
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

type OnCreatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, process *Process) js.Ref
	Arg T
}

func (cb *OnCreatedEventCallback[T]) Register() js.Func[func(process *Process)] {
	return js.RegisterCallback[func(process *Process)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Process
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

// HasFuncOnCreated returns true if the function "WEBEXT.processes.onCreated.addListener" exists.
func HasFuncOnCreated() bool {
	return js.True == bindings.HasFuncOnCreated()
}

// FuncOnCreated returns the function "WEBEXT.processes.onCreated.addListener".
func FuncOnCreated() (fn js.Func[func(callback js.Func[func(process *Process)])]) {
	bindings.FuncOnCreated(
		js.Pointer(&fn),
	)
	return
}

// OnCreated calls the function "WEBEXT.processes.onCreated.addListener" directly.
func OnCreated(callback js.Func[func(process *Process)]) (ret js.Void) {
	bindings.CallOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreated calls the function "WEBEXT.processes.onCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreated(callback js.Func[func(process *Process)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreated returns true if the function "WEBEXT.processes.onCreated.removeListener" exists.
func HasFuncOffCreated() bool {
	return js.True == bindings.HasFuncOffCreated()
}

// FuncOffCreated returns the function "WEBEXT.processes.onCreated.removeListener".
func FuncOffCreated() (fn js.Func[func(callback js.Func[func(process *Process)])]) {
	bindings.FuncOffCreated(
		js.Pointer(&fn),
	)
	return
}

// OffCreated calls the function "WEBEXT.processes.onCreated.removeListener" directly.
func OffCreated(callback js.Func[func(process *Process)]) (ret js.Void) {
	bindings.CallOffCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreated calls the function "WEBEXT.processes.onCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreated(callback js.Func[func(process *Process)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreated returns true if the function "WEBEXT.processes.onCreated.hasListener" exists.
func HasFuncHasOnCreated() bool {
	return js.True == bindings.HasFuncHasOnCreated()
}

// FuncHasOnCreated returns the function "WEBEXT.processes.onCreated.hasListener".
func FuncHasOnCreated() (fn js.Func[func(callback js.Func[func(process *Process)]) bool]) {
	bindings.FuncHasOnCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreated calls the function "WEBEXT.processes.onCreated.hasListener" directly.
func HasOnCreated(callback js.Func[func(process *Process)]) (ret bool) {
	bindings.CallHasOnCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreated calls the function "WEBEXT.processes.onCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreated(callback js.Func[func(process *Process)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnExitedEventCallbackFunc func(this js.Ref, processId int32, exitType int32, exitCode int32) js.Ref

func (fn OnExitedEventCallbackFunc) Register() js.Func[func(processId int32, exitType int32, exitCode int32)] {
	return js.RegisterCallback[func(processId int32, exitType int32, exitCode int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExitedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnExitedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processId int32, exitType int32, exitCode int32) js.Ref
	Arg T
}

func (cb *OnExitedEventCallback[T]) Register() js.Func[func(processId int32, exitType int32, exitCode int32)] {
	return js.RegisterCallback[func(processId int32, exitType int32, exitCode int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExitedEventCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.Number[int32]{}.FromRef(args[1+1]).Get(),
		js.Number[int32]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnExited returns true if the function "WEBEXT.processes.onExited.addListener" exists.
func HasFuncOnExited() bool {
	return js.True == bindings.HasFuncOnExited()
}

// FuncOnExited returns the function "WEBEXT.processes.onExited.addListener".
func FuncOnExited() (fn js.Func[func(callback js.Func[func(processId int32, exitType int32, exitCode int32)])]) {
	bindings.FuncOnExited(
		js.Pointer(&fn),
	)
	return
}

// OnExited calls the function "WEBEXT.processes.onExited.addListener" directly.
func OnExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret js.Void) {
	bindings.CallOnExited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExited calls the function "WEBEXT.processes.onExited.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExited returns true if the function "WEBEXT.processes.onExited.removeListener" exists.
func HasFuncOffExited() bool {
	return js.True == bindings.HasFuncOffExited()
}

// FuncOffExited returns the function "WEBEXT.processes.onExited.removeListener".
func FuncOffExited() (fn js.Func[func(callback js.Func[func(processId int32, exitType int32, exitCode int32)])]) {
	bindings.FuncOffExited(
		js.Pointer(&fn),
	)
	return
}

// OffExited calls the function "WEBEXT.processes.onExited.removeListener" directly.
func OffExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret js.Void) {
	bindings.CallOffExited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExited calls the function "WEBEXT.processes.onExited.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExited returns true if the function "WEBEXT.processes.onExited.hasListener" exists.
func HasFuncHasOnExited() bool {
	return js.True == bindings.HasFuncHasOnExited()
}

// FuncHasOnExited returns the function "WEBEXT.processes.onExited.hasListener".
func FuncHasOnExited() (fn js.Func[func(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) bool]) {
	bindings.FuncHasOnExited(
		js.Pointer(&fn),
	)
	return
}

// HasOnExited calls the function "WEBEXT.processes.onExited.hasListener" directly.
func HasOnExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret bool) {
	bindings.CallHasOnExited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExited calls the function "WEBEXT.processes.onExited.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExited(callback js.Func[func(processId int32, exitType int32, exitCode int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUnresponsiveEventCallbackFunc func(this js.Ref, process *Process) js.Ref

func (fn OnUnresponsiveEventCallbackFunc) Register() js.Func[func(process *Process)] {
	return js.RegisterCallback[func(process *Process)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUnresponsiveEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Process
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

type OnUnresponsiveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, process *Process) js.Ref
	Arg T
}

func (cb *OnUnresponsiveEventCallback[T]) Register() js.Func[func(process *Process)] {
	return js.RegisterCallback[func(process *Process)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUnresponsiveEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Process
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

// HasFuncOnUnresponsive returns true if the function "WEBEXT.processes.onUnresponsive.addListener" exists.
func HasFuncOnUnresponsive() bool {
	return js.True == bindings.HasFuncOnUnresponsive()
}

// FuncOnUnresponsive returns the function "WEBEXT.processes.onUnresponsive.addListener".
func FuncOnUnresponsive() (fn js.Func[func(callback js.Func[func(process *Process)])]) {
	bindings.FuncOnUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// OnUnresponsive calls the function "WEBEXT.processes.onUnresponsive.addListener" directly.
func OnUnresponsive(callback js.Func[func(process *Process)]) (ret js.Void) {
	bindings.CallOnUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUnresponsive calls the function "WEBEXT.processes.onUnresponsive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUnresponsive(callback js.Func[func(process *Process)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUnresponsive returns true if the function "WEBEXT.processes.onUnresponsive.removeListener" exists.
func HasFuncOffUnresponsive() bool {
	return js.True == bindings.HasFuncOffUnresponsive()
}

// FuncOffUnresponsive returns the function "WEBEXT.processes.onUnresponsive.removeListener".
func FuncOffUnresponsive() (fn js.Func[func(callback js.Func[func(process *Process)])]) {
	bindings.FuncOffUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// OffUnresponsive calls the function "WEBEXT.processes.onUnresponsive.removeListener" directly.
func OffUnresponsive(callback js.Func[func(process *Process)]) (ret js.Void) {
	bindings.CallOffUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUnresponsive calls the function "WEBEXT.processes.onUnresponsive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUnresponsive(callback js.Func[func(process *Process)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUnresponsive returns true if the function "WEBEXT.processes.onUnresponsive.hasListener" exists.
func HasFuncHasOnUnresponsive() bool {
	return js.True == bindings.HasFuncHasOnUnresponsive()
}

// FuncHasOnUnresponsive returns the function "WEBEXT.processes.onUnresponsive.hasListener".
func FuncHasOnUnresponsive() (fn js.Func[func(callback js.Func[func(process *Process)]) bool]) {
	bindings.FuncHasOnUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// HasOnUnresponsive calls the function "WEBEXT.processes.onUnresponsive.hasListener" directly.
func HasOnUnresponsive(callback js.Func[func(process *Process)]) (ret bool) {
	bindings.CallHasOnUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUnresponsive calls the function "WEBEXT.processes.onUnresponsive.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUnresponsive(callback js.Func[func(process *Process)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUpdatedEventCallbackFunc func(this js.Ref, processes js.Object) js.Ref

func (fn OnUpdatedEventCallbackFunc) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUpdatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUpdatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processes js.Object) js.Ref
	Arg T
}

func (cb *OnUpdatedEventCallback[T]) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUpdatedEventCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUpdated returns true if the function "WEBEXT.processes.onUpdated.addListener" exists.
func HasFuncOnUpdated() bool {
	return js.True == bindings.HasFuncOnUpdated()
}

// FuncOnUpdated returns the function "WEBEXT.processes.onUpdated.addListener".
func FuncOnUpdated() (fn js.Func[func(callback js.Func[func(processes js.Object)])]) {
	bindings.FuncOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// OnUpdated calls the function "WEBEXT.processes.onUpdated.addListener" directly.
func OnUpdated(callback js.Func[func(processes js.Object)]) (ret js.Void) {
	bindings.CallOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUpdated calls the function "WEBEXT.processes.onUpdated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUpdated(callback js.Func[func(processes js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUpdated returns true if the function "WEBEXT.processes.onUpdated.removeListener" exists.
func HasFuncOffUpdated() bool {
	return js.True == bindings.HasFuncOffUpdated()
}

// FuncOffUpdated returns the function "WEBEXT.processes.onUpdated.removeListener".
func FuncOffUpdated() (fn js.Func[func(callback js.Func[func(processes js.Object)])]) {
	bindings.FuncOffUpdated(
		js.Pointer(&fn),
	)
	return
}

// OffUpdated calls the function "WEBEXT.processes.onUpdated.removeListener" directly.
func OffUpdated(callback js.Func[func(processes js.Object)]) (ret js.Void) {
	bindings.CallOffUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUpdated calls the function "WEBEXT.processes.onUpdated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUpdated(callback js.Func[func(processes js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUpdated returns true if the function "WEBEXT.processes.onUpdated.hasListener" exists.
func HasFuncHasOnUpdated() bool {
	return js.True == bindings.HasFuncHasOnUpdated()
}

// FuncHasOnUpdated returns the function "WEBEXT.processes.onUpdated.hasListener".
func FuncHasOnUpdated() (fn js.Func[func(callback js.Func[func(processes js.Object)]) bool]) {
	bindings.FuncHasOnUpdated(
		js.Pointer(&fn),
	)
	return
}

// HasOnUpdated calls the function "WEBEXT.processes.onUpdated.hasListener" directly.
func HasOnUpdated(callback js.Func[func(processes js.Object)]) (ret bool) {
	bindings.CallHasOnUpdated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUpdated calls the function "WEBEXT.processes.onUpdated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUpdated(callback js.Func[func(processes js.Object)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUpdated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUpdatedWithMemoryEventCallbackFunc func(this js.Ref, processes js.Object) js.Ref

func (fn OnUpdatedWithMemoryEventCallbackFunc) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUpdatedWithMemoryEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUpdatedWithMemoryEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processes js.Object) js.Ref
	Arg T
}

func (cb *OnUpdatedWithMemoryEventCallback[T]) Register() js.Func[func(processes js.Object)] {
	return js.RegisterCallback[func(processes js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUpdatedWithMemoryEventCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUpdatedWithMemory returns true if the function "WEBEXT.processes.onUpdatedWithMemory.addListener" exists.
func HasFuncOnUpdatedWithMemory() bool {
	return js.True == bindings.HasFuncOnUpdatedWithMemory()
}

// FuncOnUpdatedWithMemory returns the function "WEBEXT.processes.onUpdatedWithMemory.addListener".
func FuncOnUpdatedWithMemory() (fn js.Func[func(callback js.Func[func(processes js.Object)])]) {
	bindings.FuncOnUpdatedWithMemory(
		js.Pointer(&fn),
	)
	return
}

// OnUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.addListener" directly.
func OnUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret js.Void) {
	bindings.CallOnUpdatedWithMemory(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUpdatedWithMemory(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUpdatedWithMemory returns true if the function "WEBEXT.processes.onUpdatedWithMemory.removeListener" exists.
func HasFuncOffUpdatedWithMemory() bool {
	return js.True == bindings.HasFuncOffUpdatedWithMemory()
}

// FuncOffUpdatedWithMemory returns the function "WEBEXT.processes.onUpdatedWithMemory.removeListener".
func FuncOffUpdatedWithMemory() (fn js.Func[func(callback js.Func[func(processes js.Object)])]) {
	bindings.FuncOffUpdatedWithMemory(
		js.Pointer(&fn),
	)
	return
}

// OffUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.removeListener" directly.
func OffUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret js.Void) {
	bindings.CallOffUpdatedWithMemory(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUpdatedWithMemory(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUpdatedWithMemory returns true if the function "WEBEXT.processes.onUpdatedWithMemory.hasListener" exists.
func HasFuncHasOnUpdatedWithMemory() bool {
	return js.True == bindings.HasFuncHasOnUpdatedWithMemory()
}

// FuncHasOnUpdatedWithMemory returns the function "WEBEXT.processes.onUpdatedWithMemory.hasListener".
func FuncHasOnUpdatedWithMemory() (fn js.Func[func(callback js.Func[func(processes js.Object)]) bool]) {
	bindings.FuncHasOnUpdatedWithMemory(
		js.Pointer(&fn),
	)
	return
}

// HasOnUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.hasListener" directly.
func HasOnUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret bool) {
	bindings.CallHasOnUpdatedWithMemory(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUpdatedWithMemory calls the function "WEBEXT.processes.onUpdatedWithMemory.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUpdatedWithMemory(callback js.Func[func(processes js.Object)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUpdatedWithMemory(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncTerminate returns true if the function "WEBEXT.processes.terminate" exists.
func HasFuncTerminate() bool {
	return js.True == bindings.HasFuncTerminate()
}

// FuncTerminate returns the function "WEBEXT.processes.terminate".
func FuncTerminate() (fn js.Func[func(processId int32) js.Promise[js.Boolean]]) {
	bindings.FuncTerminate(
		js.Pointer(&fn),
	)
	return
}

// Terminate calls the function "WEBEXT.processes.terminate" directly.
func Terminate(processId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallTerminate(
		js.Pointer(&ret),
		int32(processId),
	)

	return
}

// TryTerminate calls the function "WEBEXT.processes.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTerminate(processId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTerminate(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(processId),
	)

	return
}

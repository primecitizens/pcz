// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package test

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/test/bindings"
)

type AssertPromiseRejectsArgExpectedMessageChoice1 struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssertPromiseRejectsArgExpectedMessageChoice1 with all fields set.
func (p AssertPromiseRejectsArgExpectedMessageChoice1) FromRef(ref js.Ref) AssertPromiseRejectsArgExpectedMessageChoice1 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssertPromiseRejectsArgExpectedMessageChoice1 in the application heap.
func (p AssertPromiseRejectsArgExpectedMessageChoice1) New() js.Ref {
	return bindings.AssertPromiseRejectsArgExpectedMessageChoice1JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssertPromiseRejectsArgExpectedMessageChoice1) UpdateFrom(ref js.Ref) {
	bindings.AssertPromiseRejectsArgExpectedMessageChoice1JSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssertPromiseRejectsArgExpectedMessageChoice1) Update(ref js.Ref) {
	bindings.AssertPromiseRejectsArgExpectedMessageChoice1JSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssertPromiseRejectsArgExpectedMessageChoice1) FreeMembers(recursive bool) {
}

type AssertPromiseRejectsArgPromise struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssertPromiseRejectsArgPromise with all fields set.
func (p AssertPromiseRejectsArgPromise) FromRef(ref js.Ref) AssertPromiseRejectsArgPromise {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssertPromiseRejectsArgPromise in the application heap.
func (p AssertPromiseRejectsArgPromise) New() js.Ref {
	return bindings.AssertPromiseRejectsArgPromiseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssertPromiseRejectsArgPromise) UpdateFrom(ref js.Ref) {
	bindings.AssertPromiseRejectsArgPromiseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssertPromiseRejectsArgPromise) Update(ref js.Ref) {
	bindings.AssertPromiseRejectsArgPromiseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssertPromiseRejectsArgPromise) FreeMembers(recursive bool) {
}

type AssertPromiseRejectsReturnType struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssertPromiseRejectsReturnType with all fields set.
func (p AssertPromiseRejectsReturnType) FromRef(ref js.Ref) AssertPromiseRejectsReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssertPromiseRejectsReturnType in the application heap.
func (p AssertPromiseRejectsReturnType) New() js.Ref {
	return bindings.AssertPromiseRejectsReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssertPromiseRejectsReturnType) UpdateFrom(ref js.Ref) {
	bindings.AssertPromiseRejectsReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssertPromiseRejectsReturnType) Update(ref js.Ref) {
	bindings.AssertPromiseRejectsReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssertPromiseRejectsReturnType) FreeMembers(recursive bool) {
}

type AssertThrowsArgFnFunc func(this js.Ref) js.Ref

func (fn AssertThrowsArgFnFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AssertThrowsArgFnFunc) DispatchCallback(
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

type AssertThrowsArgFn[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *AssertThrowsArgFn[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AssertThrowsArgFn[T]) DispatchCallback(
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

type AssertThrowsArgMessageChoice1 struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssertThrowsArgMessageChoice1 with all fields set.
func (p AssertThrowsArgMessageChoice1) FromRef(ref js.Ref) AssertThrowsArgMessageChoice1 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssertThrowsArgMessageChoice1 in the application heap.
func (p AssertThrowsArgMessageChoice1) New() js.Ref {
	return bindings.AssertThrowsArgMessageChoice1JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssertThrowsArgMessageChoice1) UpdateFrom(ref js.Ref) {
	bindings.AssertThrowsArgMessageChoice1JSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssertThrowsArgMessageChoice1) Update(ref js.Ref) {
	bindings.AssertThrowsArgMessageChoice1JSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssertThrowsArgMessageChoice1) FreeMembers(recursive bool) {
}

type CallbackArgFuncFunc func(this js.Ref) js.Ref

func (fn CallbackArgFuncFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CallbackArgFuncFunc) DispatchCallback(
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

type CallbackArgFunc[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CallbackArgFunc[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CallbackArgFunc[T]) DispatchCallback(
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

type CallbackFailArgFuncFunc func(this js.Ref) js.Ref

func (fn CallbackFailArgFuncFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CallbackFailArgFuncFunc) DispatchCallback(
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

type CallbackFailArgFunc[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CallbackFailArgFunc[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CallbackFailArgFunc[T]) DispatchCallback(
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

type CallbackPassArgFuncFunc func(this js.Ref) js.Ref

func (fn CallbackPassArgFuncFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CallbackPassArgFuncFunc) DispatchCallback(
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

type CallbackPassArgFunc[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CallbackPassArgFunc[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CallbackPassArgFunc[T]) DispatchCallback(
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

type GetConfigReturnTypeFieldFtpServer struct {
	// Port is "GetConfigReturnTypeFieldFtpServer.port"
	//
	// Required
	Port int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetConfigReturnTypeFieldFtpServer with all fields set.
func (p GetConfigReturnTypeFieldFtpServer) FromRef(ref js.Ref) GetConfigReturnTypeFieldFtpServer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetConfigReturnTypeFieldFtpServer in the application heap.
func (p GetConfigReturnTypeFieldFtpServer) New() js.Ref {
	return bindings.GetConfigReturnTypeFieldFtpServerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetConfigReturnTypeFieldFtpServer) UpdateFrom(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldFtpServerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetConfigReturnTypeFieldFtpServer) Update(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldFtpServerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetConfigReturnTypeFieldFtpServer) FreeMembers(recursive bool) {
}

type GetConfigReturnTypeFieldLoginStatus struct {
	// IsLoggedIn is "GetConfigReturnTypeFieldLoginStatus.isLoggedIn"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLoggedIn MUST be set to true to make this field effective.
	IsLoggedIn bool
	// IsScreenLocked is "GetConfigReturnTypeFieldLoginStatus.isScreenLocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsScreenLocked MUST be set to true to make this field effective.
	IsScreenLocked bool

	FFI_USE_IsLoggedIn     bool // for IsLoggedIn.
	FFI_USE_IsScreenLocked bool // for IsScreenLocked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetConfigReturnTypeFieldLoginStatus with all fields set.
func (p GetConfigReturnTypeFieldLoginStatus) FromRef(ref js.Ref) GetConfigReturnTypeFieldLoginStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetConfigReturnTypeFieldLoginStatus in the application heap.
func (p GetConfigReturnTypeFieldLoginStatus) New() js.Ref {
	return bindings.GetConfigReturnTypeFieldLoginStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetConfigReturnTypeFieldLoginStatus) UpdateFrom(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldLoginStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetConfigReturnTypeFieldLoginStatus) Update(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldLoginStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetConfigReturnTypeFieldLoginStatus) FreeMembers(recursive bool) {
}

type GetConfigReturnTypeFieldTestServer struct {
	// Port is "GetConfigReturnTypeFieldTestServer.port"
	//
	// Required
	Port int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetConfigReturnTypeFieldTestServer with all fields set.
func (p GetConfigReturnTypeFieldTestServer) FromRef(ref js.Ref) GetConfigReturnTypeFieldTestServer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetConfigReturnTypeFieldTestServer in the application heap.
func (p GetConfigReturnTypeFieldTestServer) New() js.Ref {
	return bindings.GetConfigReturnTypeFieldTestServerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetConfigReturnTypeFieldTestServer) UpdateFrom(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldTestServerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetConfigReturnTypeFieldTestServer) Update(ref js.Ref) {
	bindings.GetConfigReturnTypeFieldTestServerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetConfigReturnTypeFieldTestServer) FreeMembers(recursive bool) {
}

type GetConfigReturnType struct {
	// CustomArg is "GetConfigReturnType.customArg"
	//
	// Optional
	CustomArg js.String
	// FtpServer is "GetConfigReturnType.ftpServer"
	//
	// Optional
	//
	// NOTE: FtpServer.FFI_USE MUST be set to true to get FtpServer used.
	FtpServer GetConfigReturnTypeFieldFtpServer
	// LoginStatus is "GetConfigReturnType.loginStatus"
	//
	// Optional
	//
	// NOTE: LoginStatus.FFI_USE MUST be set to true to get LoginStatus used.
	LoginStatus GetConfigReturnTypeFieldLoginStatus
	// TestDataDirectory is "GetConfigReturnType.testDataDirectory"
	//
	// Optional
	TestDataDirectory js.String
	// TestServer is "GetConfigReturnType.testServer"
	//
	// Optional
	//
	// NOTE: TestServer.FFI_USE MUST be set to true to get TestServer used.
	TestServer GetConfigReturnTypeFieldTestServer
	// TestWebSocketPort is "GetConfigReturnType.testWebSocketPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_TestWebSocketPort MUST be set to true to make this field effective.
	TestWebSocketPort int64
	// TestWebTransportPort is "GetConfigReturnType.testWebTransportPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_TestWebTransportPort MUST be set to true to make this field effective.
	TestWebTransportPort int64

	FFI_USE_TestWebSocketPort    bool // for TestWebSocketPort.
	FFI_USE_TestWebTransportPort bool // for TestWebTransportPort.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetConfigReturnType with all fields set.
func (p GetConfigReturnType) FromRef(ref js.Ref) GetConfigReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetConfigReturnType in the application heap.
func (p GetConfigReturnType) New() js.Ref {
	return bindings.GetConfigReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetConfigReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetConfigReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetConfigReturnType) Update(ref js.Ref) {
	bindings.GetConfigReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetConfigReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.CustomArg.Ref(),
		p.TestDataDirectory.Ref(),
	)
	p.CustomArg = p.CustomArg.FromRef(js.Undefined)
	p.TestDataDirectory = p.TestDataDirectory.FromRef(js.Undefined)
	if recursive {
		p.FtpServer.FreeMembers(true)
		p.LoginStatus.FreeMembers(true)
		p.TestServer.FreeMembers(true)
	}
}

type GetWakeEventPageReturnTypeFunc func(this js.Ref) js.Ref

func (fn GetWakeEventPageReturnTypeFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetWakeEventPageReturnTypeFunc) DispatchCallback(
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

type GetWakeEventPageReturnType[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *GetWakeEventPageReturnType[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetWakeEventPageReturnType[T]) DispatchCallback(
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

type ListenForeverArgFuncFunc func(this js.Ref) js.Ref

func (fn ListenForeverArgFuncFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListenForeverArgFuncFunc) DispatchCallback(
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

type ListenForeverArgFunc[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ListenForeverArgFunc[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListenForeverArgFunc[T]) DispatchCallback(
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

type ListenOnceArgFuncFunc func(this js.Ref) js.Ref

func (fn ListenOnceArgFuncFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListenOnceArgFuncFunc) DispatchCallback(
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

type ListenOnceArgFunc[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ListenOnceArgFunc[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListenOnceArgFunc[T]) DispatchCallback(
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

type LoadScriptReturnType struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoadScriptReturnType with all fields set.
func (p LoadScriptReturnType) FromRef(ref js.Ref) LoadScriptReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoadScriptReturnType in the application heap.
func (p LoadScriptReturnType) New() js.Ref {
	return bindings.LoadScriptReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoadScriptReturnType) UpdateFrom(ref js.Ref) {
	bindings.LoadScriptReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoadScriptReturnType) Update(ref js.Ref) {
	bindings.LoadScriptReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoadScriptReturnType) FreeMembers(recursive bool) {
}

type OnMessageArgInfo struct {
	// Data is "OnMessageArgInfo.data"
	//
	// Required
	Data js.String
	// LastMessage is "OnMessageArgInfo.lastMessage"
	//
	// Required
	LastMessage bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnMessageArgInfo with all fields set.
func (p OnMessageArgInfo) FromRef(ref js.Ref) OnMessageArgInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnMessageArgInfo in the application heap.
func (p OnMessageArgInfo) New() js.Ref {
	return bindings.OnMessageArgInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnMessageArgInfo) UpdateFrom(ref js.Ref) {
	bindings.OnMessageArgInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnMessageArgInfo) Update(ref js.Ref) {
	bindings.OnMessageArgInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnMessageArgInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type RunTestsArgTestsElemFunc func(this js.Ref) js.Ref

func (fn RunTestsArgTestsElemFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RunTestsArgTestsElemFunc) DispatchCallback(
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

type RunTestsArgTestsElem[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RunTestsArgTestsElem[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RunTestsArgTestsElem[T]) DispatchCallback(
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

type RunWithUserGestureArgFunctionToRunFunc func(this js.Ref) js.Ref

func (fn RunWithUserGestureArgFunctionToRunFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RunWithUserGestureArgFunctionToRunFunc) DispatchCallback(
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

type RunWithUserGestureArgFunctionToRun[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RunWithUserGestureArgFunctionToRun[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RunWithUserGestureArgFunctionToRun[T]) DispatchCallback(
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

type SetExceptionHandlerArgHandlerFunc func(this js.Ref, message js.String, exception js.Any) js.Ref

func (fn SetExceptionHandlerArgHandlerFunc) Register() js.Func[func(message js.String, exception js.Any)] {
	return js.RegisterCallback[func(message js.String, exception js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetExceptionHandlerArgHandlerFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetExceptionHandlerArgHandler[T any] struct {
	Fn  func(arg T, this js.Ref, message js.String, exception js.Any) js.Ref
	Arg T
}

func (cb *SetExceptionHandlerArgHandler[T]) Register() js.Func[func(message js.String, exception js.Any)] {
	return js.RegisterCallback[func(message js.String, exception js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetExceptionHandlerArgHandler[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncAssertEq returns true if the function "WEBEXT.test.assertEq" exists.
func HasFuncAssertEq() bool {
	return js.True == bindings.HasFuncAssertEq()
}

// FuncAssertEq returns the function "WEBEXT.test.assertEq".
func FuncAssertEq() (fn js.Func[func(expected js.Any, actual js.Any, message js.String)]) {
	bindings.FuncAssertEq(
		js.Pointer(&fn),
	)
	return
}

// AssertEq calls the function "WEBEXT.test.assertEq" directly.
func AssertEq(expected js.Any, actual js.Any, message js.String) (ret js.Void) {
	bindings.CallAssertEq(
		js.Pointer(&ret),
		expected.Ref(),
		actual.Ref(),
		message.Ref(),
	)

	return
}

// TryAssertEq calls the function "WEBEXT.test.assertEq"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertEq(expected js.Any, actual js.Any, message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertEq(
		js.Pointer(&ret), js.Pointer(&exception),
		expected.Ref(),
		actual.Ref(),
		message.Ref(),
	)

	return
}

type OneOf_String_Bool struct {
	ref js.Ref
}

func (x OneOf_String_Bool) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Bool) Free() {
	x.ref.Free()
}

func (x OneOf_String_Bool) FromRef(ref js.Ref) OneOf_String_Bool {
	return OneOf_String_Bool{
		ref: ref,
	}
}

func (x OneOf_String_Bool) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Bool) Bool() bool {
	return x.ref == js.True
}

// HasFuncAssertFalse returns true if the function "WEBEXT.test.assertFalse" exists.
func HasFuncAssertFalse() bool {
	return js.True == bindings.HasFuncAssertFalse()
}

// FuncAssertFalse returns the function "WEBEXT.test.assertFalse".
func FuncAssertFalse() (fn js.Func[func(test OneOf_String_Bool, message js.String)]) {
	bindings.FuncAssertFalse(
		js.Pointer(&fn),
	)
	return
}

// AssertFalse calls the function "WEBEXT.test.assertFalse" directly.
func AssertFalse(test OneOf_String_Bool, message js.String) (ret js.Void) {
	bindings.CallAssertFalse(
		js.Pointer(&ret),
		test.Ref(),
		message.Ref(),
	)

	return
}

// TryAssertFalse calls the function "WEBEXT.test.assertFalse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertFalse(test OneOf_String_Bool, message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertFalse(
		js.Pointer(&ret), js.Pointer(&exception),
		test.Ref(),
		message.Ref(),
	)

	return
}

// HasFuncAssertLastError returns true if the function "WEBEXT.test.assertLastError" exists.
func HasFuncAssertLastError() bool {
	return js.True == bindings.HasFuncAssertLastError()
}

// FuncAssertLastError returns the function "WEBEXT.test.assertLastError".
func FuncAssertLastError() (fn js.Func[func(expectedError js.String)]) {
	bindings.FuncAssertLastError(
		js.Pointer(&fn),
	)
	return
}

// AssertLastError calls the function "WEBEXT.test.assertLastError" directly.
func AssertLastError(expectedError js.String) (ret js.Void) {
	bindings.CallAssertLastError(
		js.Pointer(&ret),
		expectedError.Ref(),
	)

	return
}

// TryAssertLastError calls the function "WEBEXT.test.assertLastError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertLastError(expectedError js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertLastError(
		js.Pointer(&ret), js.Pointer(&exception),
		expectedError.Ref(),
	)

	return
}

// HasFuncAssertNe returns true if the function "WEBEXT.test.assertNe" exists.
func HasFuncAssertNe() bool {
	return js.True == bindings.HasFuncAssertNe()
}

// FuncAssertNe returns the function "WEBEXT.test.assertNe".
func FuncAssertNe() (fn js.Func[func(expected js.Any, actual js.Any, message js.String)]) {
	bindings.FuncAssertNe(
		js.Pointer(&fn),
	)
	return
}

// AssertNe calls the function "WEBEXT.test.assertNe" directly.
func AssertNe(expected js.Any, actual js.Any, message js.String) (ret js.Void) {
	bindings.CallAssertNe(
		js.Pointer(&ret),
		expected.Ref(),
		actual.Ref(),
		message.Ref(),
	)

	return
}

// TryAssertNe calls the function "WEBEXT.test.assertNe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertNe(expected js.Any, actual js.Any, message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertNe(
		js.Pointer(&ret), js.Pointer(&exception),
		expected.Ref(),
		actual.Ref(),
		message.Ref(),
	)

	return
}

// HasFuncAssertNoLastError returns true if the function "WEBEXT.test.assertNoLastError" exists.
func HasFuncAssertNoLastError() bool {
	return js.True == bindings.HasFuncAssertNoLastError()
}

// FuncAssertNoLastError returns the function "WEBEXT.test.assertNoLastError".
func FuncAssertNoLastError() (fn js.Func[func()]) {
	bindings.FuncAssertNoLastError(
		js.Pointer(&fn),
	)
	return
}

// AssertNoLastError calls the function "WEBEXT.test.assertNoLastError" directly.
func AssertNoLastError() (ret js.Void) {
	bindings.CallAssertNoLastError(
		js.Pointer(&ret),
	)

	return
}

// TryAssertNoLastError calls the function "WEBEXT.test.assertNoLastError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertNoLastError() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertNoLastError(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1 struct {
	ref js.Ref
}

func (x OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) Free() {
	x.ref.Free()
}

func (x OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) FromRef(ref js.Ref) OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1 {
	return OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1{
		ref: ref,
	}
}

func (x OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) AssertPromiseRejectsArgExpectedMessageChoice1() AssertPromiseRejectsArgExpectedMessageChoice1 {
	var ret AssertPromiseRejectsArgExpectedMessageChoice1
	ret.UpdateFrom(x.ref)
	return ret
}

// HasFuncAssertPromiseRejects returns true if the function "WEBEXT.test.assertPromiseRejects" exists.
func HasFuncAssertPromiseRejects() bool {
	return js.True == bindings.HasFuncAssertPromiseRejects()
}

// FuncAssertPromiseRejects returns the function "WEBEXT.test.assertPromiseRejects".
func FuncAssertPromiseRejects() (fn js.Func[func(promise AssertPromiseRejectsArgPromise, expectedMessage OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) AssertPromiseRejectsReturnType]) {
	bindings.FuncAssertPromiseRejects(
		js.Pointer(&fn),
	)
	return
}

// AssertPromiseRejects calls the function "WEBEXT.test.assertPromiseRejects" directly.
func AssertPromiseRejects(promise AssertPromiseRejectsArgPromise, expectedMessage OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) (ret AssertPromiseRejectsReturnType) {
	bindings.CallAssertPromiseRejects(
		js.Pointer(&ret),
		js.Pointer(&promise),
		expectedMessage.Ref(),
	)

	return
}

// TryAssertPromiseRejects calls the function "WEBEXT.test.assertPromiseRejects"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertPromiseRejects(promise AssertPromiseRejectsArgPromise, expectedMessage OneOf_String_AssertPromiseRejectsArgExpectedMessageChoice1) (ret AssertPromiseRejectsReturnType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertPromiseRejects(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&promise),
		expectedMessage.Ref(),
	)

	return
}

type OneOf_String_AssertThrowsArgMessageChoice1 struct {
	ref js.Ref
}

func (x OneOf_String_AssertThrowsArgMessageChoice1) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_AssertThrowsArgMessageChoice1) Free() {
	x.ref.Free()
}

func (x OneOf_String_AssertThrowsArgMessageChoice1) FromRef(ref js.Ref) OneOf_String_AssertThrowsArgMessageChoice1 {
	return OneOf_String_AssertThrowsArgMessageChoice1{
		ref: ref,
	}
}

func (x OneOf_String_AssertThrowsArgMessageChoice1) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_AssertThrowsArgMessageChoice1) AssertThrowsArgMessageChoice1() AssertThrowsArgMessageChoice1 {
	var ret AssertThrowsArgMessageChoice1
	ret.UpdateFrom(x.ref)
	return ret
}

// HasFuncAssertThrows returns true if the function "WEBEXT.test.assertThrows" exists.
func HasFuncAssertThrows() bool {
	return js.True == bindings.HasFuncAssertThrows()
}

// FuncAssertThrows returns the function "WEBEXT.test.assertThrows".
func FuncAssertThrows() (fn js.Func[func(fn js.Func[func()], self js.Any, args js.Array[js.Any], message OneOf_String_AssertThrowsArgMessageChoice1)]) {
	bindings.FuncAssertThrows(
		js.Pointer(&fn),
	)
	return
}

// AssertThrows calls the function "WEBEXT.test.assertThrows" directly.
func AssertThrows(fn js.Func[func()], self js.Any, args js.Array[js.Any], message OneOf_String_AssertThrowsArgMessageChoice1) (ret js.Void) {
	bindings.CallAssertThrows(
		js.Pointer(&ret),
		fn.Ref(),
		self.Ref(),
		args.Ref(),
		message.Ref(),
	)

	return
}

// TryAssertThrows calls the function "WEBEXT.test.assertThrows"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertThrows(fn js.Func[func()], self js.Any, args js.Array[js.Any], message OneOf_String_AssertThrowsArgMessageChoice1) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertThrows(
		js.Pointer(&ret), js.Pointer(&exception),
		fn.Ref(),
		self.Ref(),
		args.Ref(),
		message.Ref(),
	)

	return
}

// HasFuncAssertTrue returns true if the function "WEBEXT.test.assertTrue" exists.
func HasFuncAssertTrue() bool {
	return js.True == bindings.HasFuncAssertTrue()
}

// FuncAssertTrue returns the function "WEBEXT.test.assertTrue".
func FuncAssertTrue() (fn js.Func[func(test OneOf_String_Bool, message js.String)]) {
	bindings.FuncAssertTrue(
		js.Pointer(&fn),
	)
	return
}

// AssertTrue calls the function "WEBEXT.test.assertTrue" directly.
func AssertTrue(test OneOf_String_Bool, message js.String) (ret js.Void) {
	bindings.CallAssertTrue(
		js.Pointer(&ret),
		test.Ref(),
		message.Ref(),
	)

	return
}

// TryAssertTrue calls the function "WEBEXT.test.assertTrue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAssertTrue(test OneOf_String_Bool, message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAssertTrue(
		js.Pointer(&ret), js.Pointer(&exception),
		test.Ref(),
		message.Ref(),
	)

	return
}

// HasFuncCallback returns true if the function "WEBEXT.test.callback" exists.
func HasFuncCallback() bool {
	return js.True == bindings.HasFuncCallback()
}

// FuncCallback returns the function "WEBEXT.test.callback".
func FuncCallback() (fn js.Func[func(fn js.Func[func()], expectedError js.String)]) {
	bindings.FuncCallback(
		js.Pointer(&fn),
	)
	return
}

// Callback calls the function "WEBEXT.test.callback" directly.
func Callback(fn js.Func[func()], expectedError js.String) (ret js.Void) {
	bindings.CallCallback(
		js.Pointer(&ret),
		fn.Ref(),
		expectedError.Ref(),
	)

	return
}

// TryCallback calls the function "WEBEXT.test.callback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCallback(fn js.Func[func()], expectedError js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCallback(
		js.Pointer(&ret), js.Pointer(&exception),
		fn.Ref(),
		expectedError.Ref(),
	)

	return
}

// HasFuncCallbackAdded returns true if the function "WEBEXT.test.callbackAdded" exists.
func HasFuncCallbackAdded() bool {
	return js.True == bindings.HasFuncCallbackAdded()
}

// FuncCallbackAdded returns the function "WEBEXT.test.callbackAdded".
func FuncCallbackAdded() (fn js.Func[func()]) {
	bindings.FuncCallbackAdded(
		js.Pointer(&fn),
	)
	return
}

// CallbackAdded calls the function "WEBEXT.test.callbackAdded" directly.
func CallbackAdded() (ret js.Void) {
	bindings.CallCallbackAdded(
		js.Pointer(&ret),
	)

	return
}

// TryCallbackAdded calls the function "WEBEXT.test.callbackAdded"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCallbackAdded() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCallbackAdded(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCallbackFail returns true if the function "WEBEXT.test.callbackFail" exists.
func HasFuncCallbackFail() bool {
	return js.True == bindings.HasFuncCallbackFail()
}

// FuncCallbackFail returns the function "WEBEXT.test.callbackFail".
func FuncCallbackFail() (fn js.Func[func(expectedError js.String, fn js.Func[func()])]) {
	bindings.FuncCallbackFail(
		js.Pointer(&fn),
	)
	return
}

// CallbackFail calls the function "WEBEXT.test.callbackFail" directly.
func CallbackFail(expectedError js.String, fn js.Func[func()]) (ret js.Void) {
	bindings.CallCallbackFail(
		js.Pointer(&ret),
		expectedError.Ref(),
		fn.Ref(),
	)

	return
}

// TryCallbackFail calls the function "WEBEXT.test.callbackFail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCallbackFail(expectedError js.String, fn js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCallbackFail(
		js.Pointer(&ret), js.Pointer(&exception),
		expectedError.Ref(),
		fn.Ref(),
	)

	return
}

// HasFuncCallbackPass returns true if the function "WEBEXT.test.callbackPass" exists.
func HasFuncCallbackPass() bool {
	return js.True == bindings.HasFuncCallbackPass()
}

// FuncCallbackPass returns the function "WEBEXT.test.callbackPass".
func FuncCallbackPass() (fn js.Func[func(fn js.Func[func()])]) {
	bindings.FuncCallbackPass(
		js.Pointer(&fn),
	)
	return
}

// CallbackPass calls the function "WEBEXT.test.callbackPass" directly.
func CallbackPass(fn js.Func[func()]) (ret js.Void) {
	bindings.CallCallbackPass(
		js.Pointer(&ret),
		fn.Ref(),
	)

	return
}

// TryCallbackPass calls the function "WEBEXT.test.callbackPass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCallbackPass(fn js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCallbackPass(
		js.Pointer(&ret), js.Pointer(&exception),
		fn.Ref(),
	)

	return
}

// HasFuncCheckDeepEq returns true if the function "WEBEXT.test.checkDeepEq" exists.
func HasFuncCheckDeepEq() bool {
	return js.True == bindings.HasFuncCheckDeepEq()
}

// FuncCheckDeepEq returns the function "WEBEXT.test.checkDeepEq".
func FuncCheckDeepEq() (fn js.Func[func(expected js.Any, actual js.Any)]) {
	bindings.FuncCheckDeepEq(
		js.Pointer(&fn),
	)
	return
}

// CheckDeepEq calls the function "WEBEXT.test.checkDeepEq" directly.
func CheckDeepEq(expected js.Any, actual js.Any) (ret js.Void) {
	bindings.CallCheckDeepEq(
		js.Pointer(&ret),
		expected.Ref(),
		actual.Ref(),
	)

	return
}

// TryCheckDeepEq calls the function "WEBEXT.test.checkDeepEq"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCheckDeepEq(expected js.Any, actual js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCheckDeepEq(
		js.Pointer(&ret), js.Pointer(&exception),
		expected.Ref(),
		actual.Ref(),
	)

	return
}

// HasFuncFail returns true if the function "WEBEXT.test.fail" exists.
func HasFuncFail() bool {
	return js.True == bindings.HasFuncFail()
}

// FuncFail returns the function "WEBEXT.test.fail".
func FuncFail() (fn js.Func[func(message js.Any)]) {
	bindings.FuncFail(
		js.Pointer(&fn),
	)
	return
}

// Fail calls the function "WEBEXT.test.fail" directly.
func Fail(message js.Any) (ret js.Void) {
	bindings.CallFail(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryFail calls the function "WEBEXT.test.fail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFail(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFail(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncGetApiDefinitions returns true if the function "WEBEXT.test.getApiDefinitions" exists.
func HasFuncGetApiDefinitions() bool {
	return js.True == bindings.HasFuncGetApiDefinitions()
}

// FuncGetApiDefinitions returns the function "WEBEXT.test.getApiDefinitions".
func FuncGetApiDefinitions() (fn js.Func[func(apiNames js.Array[js.String])]) {
	bindings.FuncGetApiDefinitions(
		js.Pointer(&fn),
	)
	return
}

// GetApiDefinitions calls the function "WEBEXT.test.getApiDefinitions" directly.
func GetApiDefinitions(apiNames js.Array[js.String]) (ret js.Void) {
	bindings.CallGetApiDefinitions(
		js.Pointer(&ret),
		apiNames.Ref(),
	)

	return
}

// TryGetApiDefinitions calls the function "WEBEXT.test.getApiDefinitions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetApiDefinitions(apiNames js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetApiDefinitions(
		js.Pointer(&ret), js.Pointer(&exception),
		apiNames.Ref(),
	)

	return
}

// HasFuncGetApiFeatures returns true if the function "WEBEXT.test.getApiFeatures" exists.
func HasFuncGetApiFeatures() bool {
	return js.True == bindings.HasFuncGetApiFeatures()
}

// FuncGetApiFeatures returns the function "WEBEXT.test.getApiFeatures".
func FuncGetApiFeatures() (fn js.Func[func()]) {
	bindings.FuncGetApiFeatures(
		js.Pointer(&fn),
	)
	return
}

// GetApiFeatures calls the function "WEBEXT.test.getApiFeatures" directly.
func GetApiFeatures() (ret js.Void) {
	bindings.CallGetApiFeatures(
		js.Pointer(&ret),
	)

	return
}

// TryGetApiFeatures calls the function "WEBEXT.test.getApiFeatures"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetApiFeatures() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetApiFeatures(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetConfig returns true if the function "WEBEXT.test.getConfig" exists.
func HasFuncGetConfig() bool {
	return js.True == bindings.HasFuncGetConfig()
}

// FuncGetConfig returns the function "WEBEXT.test.getConfig".
func FuncGetConfig() (fn js.Func[func() js.Promise[GetConfigReturnType]]) {
	bindings.FuncGetConfig(
		js.Pointer(&fn),
	)
	return
}

// GetConfig calls the function "WEBEXT.test.getConfig" directly.
func GetConfig() (ret js.Promise[GetConfigReturnType]) {
	bindings.CallGetConfig(
		js.Pointer(&ret),
	)

	return
}

// TryGetConfig calls the function "WEBEXT.test.getConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetConfig() (ret js.Promise[GetConfigReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetConfig(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetModuleSystem returns true if the function "WEBEXT.test.getModuleSystem" exists.
func HasFuncGetModuleSystem() bool {
	return js.True == bindings.HasFuncGetModuleSystem()
}

// FuncGetModuleSystem returns the function "WEBEXT.test.getModuleSystem".
func FuncGetModuleSystem() (fn js.Func[func(context js.Any) js.Any]) {
	bindings.FuncGetModuleSystem(
		js.Pointer(&fn),
	)
	return
}

// GetModuleSystem calls the function "WEBEXT.test.getModuleSystem" directly.
func GetModuleSystem(context js.Any) (ret js.Any) {
	bindings.CallGetModuleSystem(
		js.Pointer(&ret),
		context.Ref(),
	)

	return
}

// TryGetModuleSystem calls the function "WEBEXT.test.getModuleSystem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetModuleSystem(context js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetModuleSystem(
		js.Pointer(&ret), js.Pointer(&exception),
		context.Ref(),
	)

	return
}

// HasFuncGetWakeEventPage returns true if the function "WEBEXT.test.getWakeEventPage" exists.
func HasFuncGetWakeEventPage() bool {
	return js.True == bindings.HasFuncGetWakeEventPage()
}

// FuncGetWakeEventPage returns the function "WEBEXT.test.getWakeEventPage".
func FuncGetWakeEventPage() (fn js.Func[func() js.Func[func()]]) {
	bindings.FuncGetWakeEventPage(
		js.Pointer(&fn),
	)
	return
}

// GetWakeEventPage calls the function "WEBEXT.test.getWakeEventPage" directly.
func GetWakeEventPage() (ret js.Func[func()]) {
	bindings.CallGetWakeEventPage(
		js.Pointer(&ret),
	)

	return
}

// TryGetWakeEventPage calls the function "WEBEXT.test.getWakeEventPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetWakeEventPage() (ret js.Func[func()], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetWakeEventPage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsProcessingUserGesture returns true if the function "WEBEXT.test.isProcessingUserGesture" exists.
func HasFuncIsProcessingUserGesture() bool {
	return js.True == bindings.HasFuncIsProcessingUserGesture()
}

// FuncIsProcessingUserGesture returns the function "WEBEXT.test.isProcessingUserGesture".
func FuncIsProcessingUserGesture() (fn js.Func[func()]) {
	bindings.FuncIsProcessingUserGesture(
		js.Pointer(&fn),
	)
	return
}

// IsProcessingUserGesture calls the function "WEBEXT.test.isProcessingUserGesture" directly.
func IsProcessingUserGesture() (ret js.Void) {
	bindings.CallIsProcessingUserGesture(
		js.Pointer(&ret),
	)

	return
}

// TryIsProcessingUserGesture calls the function "WEBEXT.test.isProcessingUserGesture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsProcessingUserGesture() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsProcessingUserGesture(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncListenForever returns true if the function "WEBEXT.test.listenForever" exists.
func HasFuncListenForever() bool {
	return js.True == bindings.HasFuncListenForever()
}

// FuncListenForever returns the function "WEBEXT.test.listenForever".
func FuncListenForever() (fn js.Func[func(event js.Any, fn js.Func[func()])]) {
	bindings.FuncListenForever(
		js.Pointer(&fn),
	)
	return
}

// ListenForever calls the function "WEBEXT.test.listenForever" directly.
func ListenForever(event js.Any, fn js.Func[func()]) (ret js.Void) {
	bindings.CallListenForever(
		js.Pointer(&ret),
		event.Ref(),
		fn.Ref(),
	)

	return
}

// TryListenForever calls the function "WEBEXT.test.listenForever"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListenForever(event js.Any, fn js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryListenForever(
		js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
		fn.Ref(),
	)

	return
}

// HasFuncListenOnce returns true if the function "WEBEXT.test.listenOnce" exists.
func HasFuncListenOnce() bool {
	return js.True == bindings.HasFuncListenOnce()
}

// FuncListenOnce returns the function "WEBEXT.test.listenOnce".
func FuncListenOnce() (fn js.Func[func(event js.Any, fn js.Func[func()])]) {
	bindings.FuncListenOnce(
		js.Pointer(&fn),
	)
	return
}

// ListenOnce calls the function "WEBEXT.test.listenOnce" directly.
func ListenOnce(event js.Any, fn js.Func[func()]) (ret js.Void) {
	bindings.CallListenOnce(
		js.Pointer(&ret),
		event.Ref(),
		fn.Ref(),
	)

	return
}

// TryListenOnce calls the function "WEBEXT.test.listenOnce"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListenOnce(event js.Any, fn js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryListenOnce(
		js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
		fn.Ref(),
	)

	return
}

// HasFuncLoadScript returns true if the function "WEBEXT.test.loadScript" exists.
func HasFuncLoadScript() bool {
	return js.True == bindings.HasFuncLoadScript()
}

// FuncLoadScript returns the function "WEBEXT.test.loadScript".
func FuncLoadScript() (fn js.Func[func(scriptUrl js.String) LoadScriptReturnType]) {
	bindings.FuncLoadScript(
		js.Pointer(&fn),
	)
	return
}

// LoadScript calls the function "WEBEXT.test.loadScript" directly.
func LoadScript(scriptUrl js.String) (ret LoadScriptReturnType) {
	bindings.CallLoadScript(
		js.Pointer(&ret),
		scriptUrl.Ref(),
	)

	return
}

// TryLoadScript calls the function "WEBEXT.test.loadScript"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadScript(scriptUrl js.String) (ret LoadScriptReturnType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadScript(
		js.Pointer(&ret), js.Pointer(&exception),
		scriptUrl.Ref(),
	)

	return
}

// HasFuncLog returns true if the function "WEBEXT.test.log" exists.
func HasFuncLog() bool {
	return js.True == bindings.HasFuncLog()
}

// FuncLog returns the function "WEBEXT.test.log".
func FuncLog() (fn js.Func[func(message js.String)]) {
	bindings.FuncLog(
		js.Pointer(&fn),
	)
	return
}

// Log calls the function "WEBEXT.test.log" directly.
func Log(message js.String) (ret js.Void) {
	bindings.CallLog(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryLog calls the function "WEBEXT.test.log"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLog(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLog(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncNotifyFail returns true if the function "WEBEXT.test.notifyFail" exists.
func HasFuncNotifyFail() bool {
	return js.True == bindings.HasFuncNotifyFail()
}

// FuncNotifyFail returns the function "WEBEXT.test.notifyFail".
func FuncNotifyFail() (fn js.Func[func(message js.String)]) {
	bindings.FuncNotifyFail(
		js.Pointer(&fn),
	)
	return
}

// NotifyFail calls the function "WEBEXT.test.notifyFail" directly.
func NotifyFail(message js.String) (ret js.Void) {
	bindings.CallNotifyFail(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryNotifyFail calls the function "WEBEXT.test.notifyFail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyFail(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyFail(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncNotifyPass returns true if the function "WEBEXT.test.notifyPass" exists.
func HasFuncNotifyPass() bool {
	return js.True == bindings.HasFuncNotifyPass()
}

// FuncNotifyPass returns the function "WEBEXT.test.notifyPass".
func FuncNotifyPass() (fn js.Func[func(message js.String)]) {
	bindings.FuncNotifyPass(
		js.Pointer(&fn),
	)
	return
}

// NotifyPass calls the function "WEBEXT.test.notifyPass" directly.
func NotifyPass(message js.String) (ret js.Void) {
	bindings.CallNotifyPass(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryNotifyPass calls the function "WEBEXT.test.notifyPass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyPass(message js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyPass(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

type OnMessageEventCallbackFunc func(this js.Ref, info *OnMessageArgInfo) js.Ref

func (fn OnMessageEventCallbackFunc) Register() js.Func[func(info *OnMessageArgInfo)] {
	return js.RegisterCallback[func(info *OnMessageArgInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgInfo
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

type OnMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *OnMessageArgInfo) js.Ref
	Arg T
}

func (cb *OnMessageEventCallback[T]) Register() js.Func[func(info *OnMessageArgInfo)] {
	return js.RegisterCallback[func(info *OnMessageArgInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgInfo
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

// HasFuncOnMessage returns true if the function "WEBEXT.test.onMessage.addListener" exists.
func HasFuncOnMessage() bool {
	return js.True == bindings.HasFuncOnMessage()
}

// FuncOnMessage returns the function "WEBEXT.test.onMessage.addListener".
func FuncOnMessage() (fn js.Func[func(callback js.Func[func(info *OnMessageArgInfo)])]) {
	bindings.FuncOnMessage(
		js.Pointer(&fn),
	)
	return
}

// OnMessage calls the function "WEBEXT.test.onMessage.addListener" directly.
func OnMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret js.Void) {
	bindings.CallOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessage calls the function "WEBEXT.test.onMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessage returns true if the function "WEBEXT.test.onMessage.removeListener" exists.
func HasFuncOffMessage() bool {
	return js.True == bindings.HasFuncOffMessage()
}

// FuncOffMessage returns the function "WEBEXT.test.onMessage.removeListener".
func FuncOffMessage() (fn js.Func[func(callback js.Func[func(info *OnMessageArgInfo)])]) {
	bindings.FuncOffMessage(
		js.Pointer(&fn),
	)
	return
}

// OffMessage calls the function "WEBEXT.test.onMessage.removeListener" directly.
func OffMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret js.Void) {
	bindings.CallOffMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessage calls the function "WEBEXT.test.onMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessage returns true if the function "WEBEXT.test.onMessage.hasListener" exists.
func HasFuncHasOnMessage() bool {
	return js.True == bindings.HasFuncHasOnMessage()
}

// FuncHasOnMessage returns the function "WEBEXT.test.onMessage.hasListener".
func FuncHasOnMessage() (fn js.Func[func(callback js.Func[func(info *OnMessageArgInfo)]) bool]) {
	bindings.FuncHasOnMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessage calls the function "WEBEXT.test.onMessage.hasListener" directly.
func HasOnMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret bool) {
	bindings.CallHasOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessage calls the function "WEBEXT.test.onMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessage(callback js.Func[func(info *OnMessageArgInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenFileUrl returns true if the function "WEBEXT.test.openFileUrl" exists.
func HasFuncOpenFileUrl() bool {
	return js.True == bindings.HasFuncOpenFileUrl()
}

// FuncOpenFileUrl returns the function "WEBEXT.test.openFileUrl".
func FuncOpenFileUrl() (fn js.Func[func(url js.String)]) {
	bindings.FuncOpenFileUrl(
		js.Pointer(&fn),
	)
	return
}

// OpenFileUrl calls the function "WEBEXT.test.openFileUrl" directly.
func OpenFileUrl(url js.String) (ret js.Void) {
	bindings.CallOpenFileUrl(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpenFileUrl calls the function "WEBEXT.test.openFileUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenFileUrl(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenFileUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncRunTests returns true if the function "WEBEXT.test.runTests" exists.
func HasFuncRunTests() bool {
	return js.True == bindings.HasFuncRunTests()
}

// FuncRunTests returns the function "WEBEXT.test.runTests".
func FuncRunTests() (fn js.Func[func(tests js.Array[js.Func[func()]])]) {
	bindings.FuncRunTests(
		js.Pointer(&fn),
	)
	return
}

// RunTests calls the function "WEBEXT.test.runTests" directly.
func RunTests(tests js.Array[js.Func[func()]]) (ret js.Void) {
	bindings.CallRunTests(
		js.Pointer(&ret),
		tests.Ref(),
	)

	return
}

// TryRunTests calls the function "WEBEXT.test.runTests"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunTests(tests js.Array[js.Func[func()]]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunTests(
		js.Pointer(&ret), js.Pointer(&exception),
		tests.Ref(),
	)

	return
}

// HasFuncRunWithUserGesture returns true if the function "WEBEXT.test.runWithUserGesture" exists.
func HasFuncRunWithUserGesture() bool {
	return js.True == bindings.HasFuncRunWithUserGesture()
}

// FuncRunWithUserGesture returns the function "WEBEXT.test.runWithUserGesture".
func FuncRunWithUserGesture() (fn js.Func[func(functionToRun js.Func[func()])]) {
	bindings.FuncRunWithUserGesture(
		js.Pointer(&fn),
	)
	return
}

// RunWithUserGesture calls the function "WEBEXT.test.runWithUserGesture" directly.
func RunWithUserGesture(functionToRun js.Func[func()]) (ret js.Void) {
	bindings.CallRunWithUserGesture(
		js.Pointer(&ret),
		functionToRun.Ref(),
	)

	return
}

// TryRunWithUserGesture calls the function "WEBEXT.test.runWithUserGesture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunWithUserGesture(functionToRun js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunWithUserGesture(
		js.Pointer(&ret), js.Pointer(&exception),
		functionToRun.Ref(),
	)

	return
}

// HasFuncSendMessage returns true if the function "WEBEXT.test.sendMessage" exists.
func HasFuncSendMessage() bool {
	return js.True == bindings.HasFuncSendMessage()
}

// FuncSendMessage returns the function "WEBEXT.test.sendMessage".
func FuncSendMessage() (fn js.Func[func(message js.String) js.Promise[js.String]]) {
	bindings.FuncSendMessage(
		js.Pointer(&fn),
	)
	return
}

// SendMessage calls the function "WEBEXT.test.sendMessage" directly.
func SendMessage(message js.String) (ret js.Promise[js.String]) {
	bindings.CallSendMessage(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TrySendMessage calls the function "WEBEXT.test.sendMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendMessage(message js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncSendScriptResult returns true if the function "WEBEXT.test.sendScriptResult" exists.
func HasFuncSendScriptResult() bool {
	return js.True == bindings.HasFuncSendScriptResult()
}

// FuncSendScriptResult returns the function "WEBEXT.test.sendScriptResult".
func FuncSendScriptResult() (fn js.Func[func(result js.Any) js.Promise[js.Void]]) {
	bindings.FuncSendScriptResult(
		js.Pointer(&fn),
	)
	return
}

// SendScriptResult calls the function "WEBEXT.test.sendScriptResult" directly.
func SendScriptResult(result js.Any) (ret js.Promise[js.Void]) {
	bindings.CallSendScriptResult(
		js.Pointer(&ret),
		result.Ref(),
	)

	return
}

// TrySendScriptResult calls the function "WEBEXT.test.sendScriptResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendScriptResult(result js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendScriptResult(
		js.Pointer(&ret), js.Pointer(&exception),
		result.Ref(),
	)

	return
}

// HasFuncSetExceptionHandler returns true if the function "WEBEXT.test.setExceptionHandler" exists.
func HasFuncSetExceptionHandler() bool {
	return js.True == bindings.HasFuncSetExceptionHandler()
}

// FuncSetExceptionHandler returns the function "WEBEXT.test.setExceptionHandler".
func FuncSetExceptionHandler() (fn js.Func[func(handler js.Func[func(message js.String, exception js.Any)])]) {
	bindings.FuncSetExceptionHandler(
		js.Pointer(&fn),
	)
	return
}

// SetExceptionHandler calls the function "WEBEXT.test.setExceptionHandler" directly.
func SetExceptionHandler(handler js.Func[func(message js.String, exception js.Any)]) (ret js.Void) {
	bindings.CallSetExceptionHandler(
		js.Pointer(&ret),
		handler.Ref(),
	)

	return
}

// TrySetExceptionHandler calls the function "WEBEXT.test.setExceptionHandler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetExceptionHandler(handler js.Func[func(message js.String, exception js.Any)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetExceptionHandler(
		js.Pointer(&ret), js.Pointer(&exception),
		handler.Ref(),
	)

	return
}

// HasFuncSucceed returns true if the function "WEBEXT.test.succeed" exists.
func HasFuncSucceed() bool {
	return js.True == bindings.HasFuncSucceed()
}

// FuncSucceed returns the function "WEBEXT.test.succeed".
func FuncSucceed() (fn js.Func[func(message js.Any)]) {
	bindings.FuncSucceed(
		js.Pointer(&fn),
	)
	return
}

// Succeed calls the function "WEBEXT.test.succeed" directly.
func Succeed(message js.Any) (ret js.Void) {
	bindings.CallSucceed(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TrySucceed calls the function "WEBEXT.test.succeed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySucceed(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySucceed(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncWaitForRoundTrip returns true if the function "WEBEXT.test.waitForRoundTrip" exists.
func HasFuncWaitForRoundTrip() bool {
	return js.True == bindings.HasFuncWaitForRoundTrip()
}

// FuncWaitForRoundTrip returns the function "WEBEXT.test.waitForRoundTrip".
func FuncWaitForRoundTrip() (fn js.Func[func(message js.String) js.Promise[js.String]]) {
	bindings.FuncWaitForRoundTrip(
		js.Pointer(&fn),
	)
	return
}

// WaitForRoundTrip calls the function "WEBEXT.test.waitForRoundTrip" directly.
func WaitForRoundTrip(message js.String) (ret js.Promise[js.String]) {
	bindings.CallWaitForRoundTrip(
		js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryWaitForRoundTrip calls the function "WEBEXT.test.waitForRoundTrip"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForRoundTrip(message js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForRoundTrip(
		js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

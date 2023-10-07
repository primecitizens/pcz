// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package crashreportprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/crashreportprivate/bindings"
)

type ErrorInfo struct {
	// Message is "ErrorInfo.message"
	//
	// Optional
	Message js.String
	// Url is "ErrorInfo.url"
	//
	// Optional
	Url js.String
	// Product is "ErrorInfo.product"
	//
	// Optional
	Product js.String
	// Version is "ErrorInfo.version"
	//
	// Optional
	Version js.String
	// LineNumber is "ErrorInfo.lineNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_LineNumber MUST be set to true to make this field effective.
	LineNumber int32
	// ColumnNumber is "ErrorInfo.columnNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_ColumnNumber MUST be set to true to make this field effective.
	ColumnNumber int32
	// DebugId is "ErrorInfo.debugId"
	//
	// Optional
	DebugId js.String
	// StackTrace is "ErrorInfo.stackTrace"
	//
	// Optional
	StackTrace js.String

	FFI_USE_LineNumber   bool // for LineNumber.
	FFI_USE_ColumnNumber bool // for ColumnNumber.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ErrorInfo with all fields set.
func (p ErrorInfo) FromRef(ref js.Ref) ErrorInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ErrorInfo in the application heap.
func (p ErrorInfo) New() js.Ref {
	return bindings.ErrorInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ErrorInfo) UpdateFrom(ref js.Ref) {
	bindings.ErrorInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ErrorInfo) Update(ref js.Ref) {
	bindings.ErrorInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ErrorInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
		p.Url.Ref(),
		p.Product.Ref(),
		p.Version.Ref(),
		p.DebugId.Ref(),
		p.StackTrace.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Product = p.Product.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.DebugId = p.DebugId.FromRef(js.Undefined)
	p.StackTrace = p.StackTrace.FromRef(js.Undefined)
}

type ReportCallbackFunc func(this js.Ref) js.Ref

func (fn ReportCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReportCallbackFunc) DispatchCallback(
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

type ReportCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ReportCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReportCallback[T]) DispatchCallback(
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

// HasFuncReportError returns true if the function "WEBEXT.crashReportPrivate.reportError" exists.
func HasFuncReportError() bool {
	return js.True == bindings.HasFuncReportError()
}

// FuncReportError returns the function "WEBEXT.crashReportPrivate.reportError".
func FuncReportError() (fn js.Func[func(info ErrorInfo) js.Promise[js.Void]]) {
	bindings.FuncReportError(
		js.Pointer(&fn),
	)
	return
}

// ReportError calls the function "WEBEXT.crashReportPrivate.reportError" directly.
func ReportError(info ErrorInfo) (ret js.Promise[js.Void]) {
	bindings.CallReportError(
		js.Pointer(&ret),
		js.Pointer(&info),
	)

	return
}

// TryReportError calls the function "WEBEXT.crashReportPrivate.reportError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportError(info ErrorInfo) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportError(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&info),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package certificateproviderinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/certificateprovider"
	"github.com/primecitizens/pcz/std/plat/js/webext/certificateproviderinternal/bindings"
)

type DoneCallbackFunc func(this js.Ref) js.Ref

func (fn DoneCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DoneCallbackFunc) DispatchCallback(
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

type DoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DoneCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DoneCallback[T]) DispatchCallback(
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

type ResultCallbackFunc func(this js.Ref, rejectedCertificates js.Array[js.ArrayBuffer]) js.Ref

func (fn ResultCallbackFunc) Register() js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(rejectedCertificates js.Array[js.ArrayBuffer])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResultCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.ArrayBuffer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rejectedCertificates js.Array[js.ArrayBuffer]) js.Ref
	Arg T
}

func (cb *ResultCallback[T]) Register() js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(rejectedCertificates js.Array[js.ArrayBuffer])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResultCallback[T]) DispatchCallback(
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

		js.Array[js.ArrayBuffer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncReportCertificates returns true if the function "WEBEXT.certificateProviderInternal.reportCertificates" exists.
func HasFuncReportCertificates() bool {
	return js.True == bindings.HasFuncReportCertificates()
}

// FuncReportCertificates returns the function "WEBEXT.certificateProviderInternal.reportCertificates".
func FuncReportCertificates() (fn js.Func[func(requestId int32, certificates js.Array[certificateprovider.CertificateInfo]) js.Promise[js.Array[js.ArrayBuffer]]]) {
	bindings.FuncReportCertificates(
		js.Pointer(&fn),
	)
	return
}

// ReportCertificates calls the function "WEBEXT.certificateProviderInternal.reportCertificates" directly.
func ReportCertificates(requestId int32, certificates js.Array[certificateprovider.CertificateInfo]) (ret js.Promise[js.Array[js.ArrayBuffer]]) {
	bindings.CallReportCertificates(
		js.Pointer(&ret),
		int32(requestId),
		certificates.Ref(),
	)

	return
}

// TryReportCertificates calls the function "WEBEXT.certificateProviderInternal.reportCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportCertificates(requestId int32, certificates js.Array[certificateprovider.CertificateInfo]) (ret js.Promise[js.Array[js.ArrayBuffer]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportCertificates(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		certificates.Ref(),
	)

	return
}

// HasFuncReportSignature returns true if the function "WEBEXT.certificateProviderInternal.reportSignature" exists.
func HasFuncReportSignature() bool {
	return js.True == bindings.HasFuncReportSignature()
}

// FuncReportSignature returns the function "WEBEXT.certificateProviderInternal.reportSignature".
func FuncReportSignature() (fn js.Func[func(requestId int32, signature js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncReportSignature(
		js.Pointer(&fn),
	)
	return
}

// ReportSignature calls the function "WEBEXT.certificateProviderInternal.reportSignature" directly.
func ReportSignature(requestId int32, signature js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallReportSignature(
		js.Pointer(&ret),
		int32(requestId),
		signature.Ref(),
	)

	return
}

// TryReportSignature calls the function "WEBEXT.certificateProviderInternal.reportSignature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportSignature(requestId int32, signature js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportSignature(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		signature.Ref(),
	)

	return
}

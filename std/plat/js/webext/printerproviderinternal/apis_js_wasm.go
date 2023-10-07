// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package printerproviderinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/printerprovider"
	"github.com/primecitizens/pcz/std/plat/js/webext/printerproviderinternal/bindings"
)

type BlobCallbackFunc func(this js.Ref, blob js.Object) js.Ref

func (fn BlobCallbackFunc) Register() js.Func[func(blob js.Object)] {
	return js.RegisterCallback[func(blob js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BlobCallbackFunc) DispatchCallback(
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

type BlobCallback[T any] struct {
	Fn  func(arg T, this js.Ref, blob js.Object) js.Ref
	Arg T
}

func (cb *BlobCallback[T]) Register() js.Func[func(blob js.Object)] {
	return js.RegisterCallback[func(blob js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BlobCallback[T]) DispatchCallback(
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

type PrintError uint32

const (
	_ PrintError = iota

	PrintError_OK
	PrintError_FAILED
	PrintError_INVALID_TICKET
	PrintError_INVALID_DATA
)

func (PrintError) FromRef(str js.Ref) PrintError {
	return PrintError(bindings.ConstOfPrintError(str))
}

func (x PrintError) String() (string, bool) {
	switch x {
	case PrintError_OK:
		return "OK", true
	case PrintError_FAILED:
		return "FAILED", true
	case PrintError_INVALID_TICKET:
		return "INVALID_TICKET", true
	case PrintError_INVALID_DATA:
		return "INVALID_DATA", true
	default:
		return "", false
	}
}

// HasFuncGetPrintData returns true if the function "WEBEXT.printerProviderInternal.getPrintData" exists.
func HasFuncGetPrintData() bool {
	return js.True == bindings.HasFuncGetPrintData()
}

// FuncGetPrintData returns the function "WEBEXT.printerProviderInternal.getPrintData".
func FuncGetPrintData() (fn js.Func[func(requestId int32) js.Promise[js.Object]]) {
	bindings.FuncGetPrintData(
		js.Pointer(&fn),
	)
	return
}

// GetPrintData calls the function "WEBEXT.printerProviderInternal.getPrintData" directly.
func GetPrintData(requestId int32) (ret js.Promise[js.Object]) {
	bindings.CallGetPrintData(
		js.Pointer(&ret),
		int32(requestId),
	)

	return
}

// TryGetPrintData calls the function "WEBEXT.printerProviderInternal.getPrintData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrintData(requestId int32) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrintData(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
	)

	return
}

// HasFuncReportPrintResult returns true if the function "WEBEXT.printerProviderInternal.reportPrintResult" exists.
func HasFuncReportPrintResult() bool {
	return js.True == bindings.HasFuncReportPrintResult()
}

// FuncReportPrintResult returns the function "WEBEXT.printerProviderInternal.reportPrintResult".
func FuncReportPrintResult() (fn js.Func[func(request_id int32, err PrintError)]) {
	bindings.FuncReportPrintResult(
		js.Pointer(&fn),
	)
	return
}

// ReportPrintResult calls the function "WEBEXT.printerProviderInternal.reportPrintResult" directly.
func ReportPrintResult(request_id int32, err PrintError) (ret js.Void) {
	bindings.CallReportPrintResult(
		js.Pointer(&ret),
		int32(request_id),
		uint32(err),
	)

	return
}

// TryReportPrintResult calls the function "WEBEXT.printerProviderInternal.reportPrintResult"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportPrintResult(request_id int32, err PrintError) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportPrintResult(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(request_id),
		uint32(err),
	)

	return
}

// HasFuncReportPrinterCapability returns true if the function "WEBEXT.printerProviderInternal.reportPrinterCapability" exists.
func HasFuncReportPrinterCapability() bool {
	return js.True == bindings.HasFuncReportPrinterCapability()
}

// FuncReportPrinterCapability returns the function "WEBEXT.printerProviderInternal.reportPrinterCapability".
func FuncReportPrinterCapability() (fn js.Func[func(request_id int32, capability js.Object)]) {
	bindings.FuncReportPrinterCapability(
		js.Pointer(&fn),
	)
	return
}

// ReportPrinterCapability calls the function "WEBEXT.printerProviderInternal.reportPrinterCapability" directly.
func ReportPrinterCapability(request_id int32, capability js.Object) (ret js.Void) {
	bindings.CallReportPrinterCapability(
		js.Pointer(&ret),
		int32(request_id),
		capability.Ref(),
	)

	return
}

// TryReportPrinterCapability calls the function "WEBEXT.printerProviderInternal.reportPrinterCapability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportPrinterCapability(request_id int32, capability js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportPrinterCapability(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(request_id),
		capability.Ref(),
	)

	return
}

// HasFuncReportPrinters returns true if the function "WEBEXT.printerProviderInternal.reportPrinters" exists.
func HasFuncReportPrinters() bool {
	return js.True == bindings.HasFuncReportPrinters()
}

// FuncReportPrinters returns the function "WEBEXT.printerProviderInternal.reportPrinters".
func FuncReportPrinters() (fn js.Func[func(requestId int32, printers js.Array[printerprovider.PrinterInfo])]) {
	bindings.FuncReportPrinters(
		js.Pointer(&fn),
	)
	return
}

// ReportPrinters calls the function "WEBEXT.printerProviderInternal.reportPrinters" directly.
func ReportPrinters(requestId int32, printers js.Array[printerprovider.PrinterInfo]) (ret js.Void) {
	bindings.CallReportPrinters(
		js.Pointer(&ret),
		int32(requestId),
		printers.Ref(),
	)

	return
}

// TryReportPrinters calls the function "WEBEXT.printerProviderInternal.reportPrinters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportPrinters(requestId int32, printers js.Array[printerprovider.PrinterInfo]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportPrinters(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		printers.Ref(),
	)

	return
}

// HasFuncReportUsbPrinterInfo returns true if the function "WEBEXT.printerProviderInternal.reportUsbPrinterInfo" exists.
func HasFuncReportUsbPrinterInfo() bool {
	return js.True == bindings.HasFuncReportUsbPrinterInfo()
}

// FuncReportUsbPrinterInfo returns the function "WEBEXT.printerProviderInternal.reportUsbPrinterInfo".
func FuncReportUsbPrinterInfo() (fn js.Func[func(requestId int32, printerInfo printerprovider.PrinterInfo)]) {
	bindings.FuncReportUsbPrinterInfo(
		js.Pointer(&fn),
	)
	return
}

// ReportUsbPrinterInfo calls the function "WEBEXT.printerProviderInternal.reportUsbPrinterInfo" directly.
func ReportUsbPrinterInfo(requestId int32, printerInfo printerprovider.PrinterInfo) (ret js.Void) {
	bindings.CallReportUsbPrinterInfo(
		js.Pointer(&ret),
		int32(requestId),
		js.Pointer(&printerInfo),
	)

	return
}

// TryReportUsbPrinterInfo calls the function "WEBEXT.printerProviderInternal.reportUsbPrinterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportUsbPrinterInfo(requestId int32, printerInfo printerprovider.PrinterInfo) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportUsbPrinterInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		js.Pointer(&printerInfo),
	)

	return
}

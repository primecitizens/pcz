// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package printerprovider

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/printerprovider/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/usb"
)

type CapabilitiesCallbackFunc func(this js.Ref, capabilities js.Object) js.Ref

func (fn CapabilitiesCallbackFunc) Register() js.Func[func(capabilities js.Object)] {
	return js.RegisterCallback[func(capabilities js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CapabilitiesCallbackFunc) DispatchCallback(
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

type CapabilitiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, capabilities js.Object) js.Ref
	Arg T
}

func (cb *CapabilitiesCallback[T]) Register() js.Func[func(capabilities js.Object)] {
	return js.RegisterCallback[func(capabilities js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CapabilitiesCallback[T]) DispatchCallback(
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

type PrintCallbackFunc func(this js.Ref, result PrintError) js.Ref

func (fn PrintCallbackFunc) Register() js.Func[func(result PrintError)] {
	return js.RegisterCallback[func(result PrintError)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PrintCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		PrintError(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PrintCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result PrintError) js.Ref
	Arg T
}

func (cb *PrintCallback[T]) Register() js.Func[func(result PrintError)] {
	return js.RegisterCallback[func(result PrintError)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PrintCallback[T]) DispatchCallback(
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

		PrintError(0).FromRef(args[0+1]),
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

type PrintJob struct {
	// PrinterId is "PrintJob.printerId"
	//
	// Optional
	PrinterId js.String
	// Title is "PrintJob.title"
	//
	// Optional
	Title js.String
	// Ticket is "PrintJob.ticket"
	//
	// Optional
	Ticket js.Object
	// ContentType is "PrintJob.contentType"
	//
	// Optional
	ContentType js.String
	// Document is "PrintJob.document"
	//
	// Optional
	Document js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrintJob with all fields set.
func (p PrintJob) FromRef(ref js.Ref) PrintJob {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrintJob in the application heap.
func (p PrintJob) New() js.Ref {
	return bindings.PrintJobJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrintJob) UpdateFrom(ref js.Ref) {
	bindings.PrintJobJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrintJob) Update(ref js.Ref) {
	bindings.PrintJobJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrintJob) FreeMembers(recursive bool) {
	js.Free(
		p.PrinterId.Ref(),
		p.Title.Ref(),
		p.Ticket.Ref(),
		p.ContentType.Ref(),
		p.Document.Ref(),
	)
	p.PrinterId = p.PrinterId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Ticket = p.Ticket.FromRef(js.Undefined)
	p.ContentType = p.ContentType.FromRef(js.Undefined)
	p.Document = p.Document.FromRef(js.Undefined)
}

type PrinterInfo struct {
	// Id is "PrinterInfo.id"
	//
	// Optional
	Id js.String
	// Name is "PrinterInfo.name"
	//
	// Optional
	Name js.String
	// Description is "PrinterInfo.description"
	//
	// Optional
	Description js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrinterInfo with all fields set.
func (p PrinterInfo) FromRef(ref js.Ref) PrinterInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrinterInfo in the application heap.
func (p PrinterInfo) New() js.Ref {
	return bindings.PrinterInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrinterInfo) UpdateFrom(ref js.Ref) {
	bindings.PrinterInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrinterInfo) Update(ref js.Ref) {
	bindings.PrinterInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrinterInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
		p.Description.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
}

type PrinterInfoCallbackFunc func(this js.Ref, printerInfo *PrinterInfo) js.Ref

func (fn PrinterInfoCallbackFunc) Register() js.Func[func(printerInfo *PrinterInfo)] {
	return js.RegisterCallback[func(printerInfo *PrinterInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PrinterInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrinterInfo
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

type PrinterInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printerInfo *PrinterInfo) js.Ref
	Arg T
}

func (cb *PrinterInfoCallback[T]) Register() js.Func[func(printerInfo *PrinterInfo)] {
	return js.RegisterCallback[func(printerInfo *PrinterInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PrinterInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrinterInfo
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

type PrintersCallbackFunc func(this js.Ref, printerInfo js.Array[PrinterInfo]) js.Ref

func (fn PrintersCallbackFunc) Register() js.Func[func(printerInfo js.Array[PrinterInfo])] {
	return js.RegisterCallback[func(printerInfo js.Array[PrinterInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PrintersCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PrinterInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PrintersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printerInfo js.Array[PrinterInfo]) js.Ref
	Arg T
}

func (cb *PrintersCallback[T]) Register() js.Func[func(printerInfo js.Array[PrinterInfo])] {
	return js.RegisterCallback[func(printerInfo js.Array[PrinterInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PrintersCallback[T]) DispatchCallback(
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

		js.Array[PrinterInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetCapabilityRequestedEventCallbackFunc func(this js.Ref, printerId js.String, resultCallback js.Func[func(capabilities js.Object)]) js.Ref

func (fn OnGetCapabilityRequestedEventCallbackFunc) Register() js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])] {
	return js.RegisterCallback[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetCapabilityRequestedEventCallbackFunc) DispatchCallback(
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
		js.Func[func(capabilities js.Object)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetCapabilityRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printerId js.String, resultCallback js.Func[func(capabilities js.Object)]) js.Ref
	Arg T
}

func (cb *OnGetCapabilityRequestedEventCallback[T]) Register() js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])] {
	return js.RegisterCallback[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetCapabilityRequestedEventCallback[T]) DispatchCallback(
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
		js.Func[func(capabilities js.Object)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetCapabilityRequested returns true if the function "WEBEXT.printerProvider.onGetCapabilityRequested.addListener" exists.
func HasFuncOnGetCapabilityRequested() bool {
	return js.True == bindings.HasFuncOnGetCapabilityRequested()
}

// FuncOnGetCapabilityRequested returns the function "WEBEXT.printerProvider.onGetCapabilityRequested.addListener".
func FuncOnGetCapabilityRequested() (fn js.Func[func(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])])]) {
	bindings.FuncOnGetCapabilityRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.addListener" directly.
func OnGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret js.Void) {
	bindings.CallOnGetCapabilityRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetCapabilityRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetCapabilityRequested returns true if the function "WEBEXT.printerProvider.onGetCapabilityRequested.removeListener" exists.
func HasFuncOffGetCapabilityRequested() bool {
	return js.True == bindings.HasFuncOffGetCapabilityRequested()
}

// FuncOffGetCapabilityRequested returns the function "WEBEXT.printerProvider.onGetCapabilityRequested.removeListener".
func FuncOffGetCapabilityRequested() (fn js.Func[func(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])])]) {
	bindings.FuncOffGetCapabilityRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.removeListener" directly.
func OffGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret js.Void) {
	bindings.CallOffGetCapabilityRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetCapabilityRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetCapabilityRequested returns true if the function "WEBEXT.printerProvider.onGetCapabilityRequested.hasListener" exists.
func HasFuncHasOnGetCapabilityRequested() bool {
	return js.True == bindings.HasFuncHasOnGetCapabilityRequested()
}

// FuncHasOnGetCapabilityRequested returns the function "WEBEXT.printerProvider.onGetCapabilityRequested.hasListener".
func FuncHasOnGetCapabilityRequested() (fn js.Func[func(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) bool]) {
	bindings.FuncHasOnGetCapabilityRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.hasListener" directly.
func HasOnGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret bool) {
	bindings.CallHasOnGetCapabilityRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetCapabilityRequested calls the function "WEBEXT.printerProvider.onGetCapabilityRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetCapabilityRequested(callback js.Func[func(printerId js.String, resultCallback js.Func[func(capabilities js.Object)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetCapabilityRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetPrintersRequestedEventCallbackFunc func(this js.Ref, resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])]) js.Ref

func (fn OnGetPrintersRequestedEventCallbackFunc) Register() js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])] {
	return js.RegisterCallback[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetPrintersRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Func[func(printerInfo js.Array[PrinterInfo])]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetPrintersRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])]) js.Ref
	Arg T
}

func (cb *OnGetPrintersRequestedEventCallback[T]) Register() js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])] {
	return js.RegisterCallback[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetPrintersRequestedEventCallback[T]) DispatchCallback(
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

		js.Func[func(printerInfo js.Array[PrinterInfo])]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetPrintersRequested returns true if the function "WEBEXT.printerProvider.onGetPrintersRequested.addListener" exists.
func HasFuncOnGetPrintersRequested() bool {
	return js.True == bindings.HasFuncOnGetPrintersRequested()
}

// FuncOnGetPrintersRequested returns the function "WEBEXT.printerProvider.onGetPrintersRequested.addListener".
func FuncOnGetPrintersRequested() (fn js.Func[func(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])])]) {
	bindings.FuncOnGetPrintersRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.addListener" directly.
func OnGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret js.Void) {
	bindings.CallOnGetPrintersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetPrintersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetPrintersRequested returns true if the function "WEBEXT.printerProvider.onGetPrintersRequested.removeListener" exists.
func HasFuncOffGetPrintersRequested() bool {
	return js.True == bindings.HasFuncOffGetPrintersRequested()
}

// FuncOffGetPrintersRequested returns the function "WEBEXT.printerProvider.onGetPrintersRequested.removeListener".
func FuncOffGetPrintersRequested() (fn js.Func[func(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])])]) {
	bindings.FuncOffGetPrintersRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.removeListener" directly.
func OffGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret js.Void) {
	bindings.CallOffGetPrintersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetPrintersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetPrintersRequested returns true if the function "WEBEXT.printerProvider.onGetPrintersRequested.hasListener" exists.
func HasFuncHasOnGetPrintersRequested() bool {
	return js.True == bindings.HasFuncHasOnGetPrintersRequested()
}

// FuncHasOnGetPrintersRequested returns the function "WEBEXT.printerProvider.onGetPrintersRequested.hasListener".
func FuncHasOnGetPrintersRequested() (fn js.Func[func(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) bool]) {
	bindings.FuncHasOnGetPrintersRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.hasListener" directly.
func HasOnGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret bool) {
	bindings.CallHasOnGetPrintersRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetPrintersRequested calls the function "WEBEXT.printerProvider.onGetPrintersRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetPrintersRequested(callback js.Func[func(resultCallback js.Func[func(printerInfo js.Array[PrinterInfo])])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetPrintersRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetUsbPrinterInfoRequestedEventCallbackFunc func(this js.Ref, device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)]) js.Ref

func (fn OnGetUsbPrinterInfoRequestedEventCallbackFunc) Register() js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])] {
	return js.RegisterCallback[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetUsbPrinterInfoRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 usb.Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(printerInfo *PrinterInfo)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnGetUsbPrinterInfoRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)]) js.Ref
	Arg T
}

func (cb *OnGetUsbPrinterInfoRequestedEventCallback[T]) Register() js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])] {
	return js.RegisterCallback[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetUsbPrinterInfoRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 usb.Device
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(printerInfo *PrinterInfo)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnGetUsbPrinterInfoRequested returns true if the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener" exists.
func HasFuncOnGetUsbPrinterInfoRequested() bool {
	return js.True == bindings.HasFuncOnGetUsbPrinterInfoRequested()
}

// FuncOnGetUsbPrinterInfoRequested returns the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener".
func FuncOnGetUsbPrinterInfoRequested() (fn js.Func[func(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])])]) {
	bindings.FuncOnGetUsbPrinterInfoRequested(
		js.Pointer(&fn),
	)
	return
}

// OnGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener" directly.
func OnGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret js.Void) {
	bindings.CallOnGetUsbPrinterInfoRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetUsbPrinterInfoRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetUsbPrinterInfoRequested returns true if the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener" exists.
func HasFuncOffGetUsbPrinterInfoRequested() bool {
	return js.True == bindings.HasFuncOffGetUsbPrinterInfoRequested()
}

// FuncOffGetUsbPrinterInfoRequested returns the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener".
func FuncOffGetUsbPrinterInfoRequested() (fn js.Func[func(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])])]) {
	bindings.FuncOffGetUsbPrinterInfoRequested(
		js.Pointer(&fn),
	)
	return
}

// OffGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener" directly.
func OffGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret js.Void) {
	bindings.CallOffGetUsbPrinterInfoRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetUsbPrinterInfoRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetUsbPrinterInfoRequested returns true if the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener" exists.
func HasFuncHasOnGetUsbPrinterInfoRequested() bool {
	return js.True == bindings.HasFuncHasOnGetUsbPrinterInfoRequested()
}

// FuncHasOnGetUsbPrinterInfoRequested returns the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener".
func FuncHasOnGetUsbPrinterInfoRequested() (fn js.Func[func(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) bool]) {
	bindings.FuncHasOnGetUsbPrinterInfoRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener" directly.
func HasOnGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret bool) {
	bindings.CallHasOnGetUsbPrinterInfoRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetUsbPrinterInfoRequested calls the function "WEBEXT.printerProvider.onGetUsbPrinterInfoRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetUsbPrinterInfoRequested(callback js.Func[func(device *usb.Device, resultCallback js.Func[func(printerInfo *PrinterInfo)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetUsbPrinterInfoRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPrintRequestedEventCallbackFunc func(this js.Ref, printJob *PrintJob, resultCallback js.Func[func(result PrintError)]) js.Ref

func (fn OnPrintRequestedEventCallbackFunc) Register() js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])] {
	return js.RegisterCallback[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrintRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrintJob
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(result PrintError)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPrintRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printJob *PrintJob, resultCallback js.Func[func(result PrintError)]) js.Ref
	Arg T
}

func (cb *OnPrintRequestedEventCallback[T]) Register() js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])] {
	return js.RegisterCallback[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrintRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrintJob
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(result PrintError)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPrintRequested returns true if the function "WEBEXT.printerProvider.onPrintRequested.addListener" exists.
func HasFuncOnPrintRequested() bool {
	return js.True == bindings.HasFuncOnPrintRequested()
}

// FuncOnPrintRequested returns the function "WEBEXT.printerProvider.onPrintRequested.addListener".
func FuncOnPrintRequested() (fn js.Func[func(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])])]) {
	bindings.FuncOnPrintRequested(
		js.Pointer(&fn),
	)
	return
}

// OnPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.addListener" directly.
func OnPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret js.Void) {
	bindings.CallOnPrintRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPrintRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPrintRequested returns true if the function "WEBEXT.printerProvider.onPrintRequested.removeListener" exists.
func HasFuncOffPrintRequested() bool {
	return js.True == bindings.HasFuncOffPrintRequested()
}

// FuncOffPrintRequested returns the function "WEBEXT.printerProvider.onPrintRequested.removeListener".
func FuncOffPrintRequested() (fn js.Func[func(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])])]) {
	bindings.FuncOffPrintRequested(
		js.Pointer(&fn),
	)
	return
}

// OffPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.removeListener" directly.
func OffPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret js.Void) {
	bindings.CallOffPrintRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPrintRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPrintRequested returns true if the function "WEBEXT.printerProvider.onPrintRequested.hasListener" exists.
func HasFuncHasOnPrintRequested() bool {
	return js.True == bindings.HasFuncHasOnPrintRequested()
}

// FuncHasOnPrintRequested returns the function "WEBEXT.printerProvider.onPrintRequested.hasListener".
func FuncHasOnPrintRequested() (fn js.Func[func(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) bool]) {
	bindings.FuncHasOnPrintRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.hasListener" directly.
func HasOnPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret bool) {
	bindings.CallHasOnPrintRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPrintRequested calls the function "WEBEXT.printerProvider.onPrintRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPrintRequested(callback js.Func[func(printJob *PrintJob, resultCallback js.Func[func(result PrintError)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPrintRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

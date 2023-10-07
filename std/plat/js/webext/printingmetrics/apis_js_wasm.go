// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package printingmetrics

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/printing"
	"github.com/primecitizens/pcz/std/plat/js/webext/printingmetrics/bindings"
)

type ColorMode uint32

const (
	_ ColorMode = iota

	ColorMode_BLACK_AND_WHITE
	ColorMode_COLOR
)

func (ColorMode) FromRef(str js.Ref) ColorMode {
	return ColorMode(bindings.ConstOfColorMode(str))
}

func (x ColorMode) String() (string, bool) {
	switch x {
	case ColorMode_BLACK_AND_WHITE:
		return "BLACK_AND_WHITE", true
	case ColorMode_COLOR:
		return "COLOR", true
	default:
		return "", false
	}
}

type DuplexMode uint32

const (
	_ DuplexMode = iota

	DuplexMode_ONE_SIDED
	DuplexMode_TWO_SIDED_LONG_EDGE
	DuplexMode_TWO_SIDED_SHORT_EDGE
)

func (DuplexMode) FromRef(str js.Ref) DuplexMode {
	return DuplexMode(bindings.ConstOfDuplexMode(str))
}

func (x DuplexMode) String() (string, bool) {
	switch x {
	case DuplexMode_ONE_SIDED:
		return "ONE_SIDED", true
	case DuplexMode_TWO_SIDED_LONG_EDGE:
		return "TWO_SIDED_LONG_EDGE", true
	case DuplexMode_TWO_SIDED_SHORT_EDGE:
		return "TWO_SIDED_SHORT_EDGE", true
	default:
		return "", false
	}
}

type GetPrintJobsCallbackFunc func(this js.Ref, jobs js.Array[PrintJobInfo]) js.Ref

func (fn GetPrintJobsCallbackFunc) Register() js.Func[func(jobs js.Array[PrintJobInfo])] {
	return js.RegisterCallback[func(jobs js.Array[PrintJobInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPrintJobsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PrintJobInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPrintJobsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, jobs js.Array[PrintJobInfo]) js.Ref
	Arg T
}

func (cb *GetPrintJobsCallback[T]) Register() js.Func[func(jobs js.Array[PrintJobInfo])] {
	return js.RegisterCallback[func(jobs js.Array[PrintJobInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPrintJobsCallback[T]) DispatchCallback(
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

		js.Array[PrintJobInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PrintJobSource uint32

const (
	_ PrintJobSource = iota

	PrintJobSource_PRINT_PREVIEW
	PrintJobSource_ANDROID_APP
	PrintJobSource_EXTENSION
)

func (PrintJobSource) FromRef(str js.Ref) PrintJobSource {
	return PrintJobSource(bindings.ConstOfPrintJobSource(str))
}

func (x PrintJobSource) String() (string, bool) {
	switch x {
	case PrintJobSource_PRINT_PREVIEW:
		return "PRINT_PREVIEW", true
	case PrintJobSource_ANDROID_APP:
		return "ANDROID_APP", true
	case PrintJobSource_EXTENSION:
		return "EXTENSION", true
	default:
		return "", false
	}
}

type PrintJobStatus uint32

const (
	_ PrintJobStatus = iota

	PrintJobStatus_FAILED
	PrintJobStatus_CANCELED
	PrintJobStatus_PRINTED
)

func (PrintJobStatus) FromRef(str js.Ref) PrintJobStatus {
	return PrintJobStatus(bindings.ConstOfPrintJobStatus(str))
}

func (x PrintJobStatus) String() (string, bool) {
	switch x {
	case PrintJobStatus_FAILED:
		return "FAILED", true
	case PrintJobStatus_CANCELED:
		return "CANCELED", true
	case PrintJobStatus_PRINTED:
		return "PRINTED", true
	default:
		return "", false
	}
}

type PrinterSource uint32

const (
	_ PrinterSource = iota

	PrinterSource_USER
	PrinterSource_POLICY
)

func (PrinterSource) FromRef(str js.Ref) PrinterSource {
	return PrinterSource(bindings.ConstOfPrinterSource(str))
}

func (x PrinterSource) String() (string, bool) {
	switch x {
	case PrinterSource_USER:
		return "USER", true
	case PrinterSource_POLICY:
		return "POLICY", true
	default:
		return "", false
	}
}

type Printer struct {
	// Name is "Printer.name"
	//
	// Optional
	Name js.String
	// Uri is "Printer.uri"
	//
	// Optional
	Uri js.String
	// Source is "Printer.source"
	//
	// Optional
	Source PrinterSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Printer with all fields set.
func (p Printer) FromRef(ref js.Ref) Printer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Printer in the application heap.
func (p Printer) New() js.Ref {
	return bindings.PrinterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Printer) UpdateFrom(ref js.Ref) {
	bindings.PrinterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Printer) Update(ref js.Ref) {
	bindings.PrinterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Printer) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Uri.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Uri = p.Uri.FromRef(js.Undefined)
}

type MediaSize struct {
	// Width is "MediaSize.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "MediaSize.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// VendorId is "MediaSize.vendorId"
	//
	// Optional
	VendorId js.String

	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaSize with all fields set.
func (p MediaSize) FromRef(ref js.Ref) MediaSize {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaSize in the application heap.
func (p MediaSize) New() js.Ref {
	return bindings.MediaSizeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaSize) UpdateFrom(ref js.Ref) {
	bindings.MediaSizeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaSize) Update(ref js.Ref) {
	bindings.MediaSizeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaSize) FreeMembers(recursive bool) {
	js.Free(
		p.VendorId.Ref(),
	)
	p.VendorId = p.VendorId.FromRef(js.Undefined)
}

type PrintSettings struct {
	// Color is "PrintSettings.color"
	//
	// Optional
	Color ColorMode
	// Duplex is "PrintSettings.duplex"
	//
	// Optional
	Duplex DuplexMode
	// MediaSize is "PrintSettings.mediaSize"
	//
	// Optional
	//
	// NOTE: MediaSize.FFI_USE MUST be set to true to get MediaSize used.
	MediaSize MediaSize
	// Copies is "PrintSettings.copies"
	//
	// Optional
	//
	// NOTE: FFI_USE_Copies MUST be set to true to make this field effective.
	Copies int32

	FFI_USE_Copies bool // for Copies.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrintSettings with all fields set.
func (p PrintSettings) FromRef(ref js.Ref) PrintSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrintSettings in the application heap.
func (p PrintSettings) New() js.Ref {
	return bindings.PrintSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrintSettings) UpdateFrom(ref js.Ref) {
	bindings.PrintSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrintSettings) Update(ref js.Ref) {
	bindings.PrintSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrintSettings) FreeMembers(recursive bool) {
	if recursive {
		p.MediaSize.FreeMembers(true)
	}
}

type PrintJobInfo struct {
	// Id is "PrintJobInfo.id"
	//
	// Optional
	Id js.String
	// Title is "PrintJobInfo.title"
	//
	// Optional
	Title js.String
	// Source is "PrintJobInfo.source"
	//
	// Optional
	Source PrintJobSource
	// SourceId is "PrintJobInfo.sourceId"
	//
	// Optional
	SourceId js.String
	// Status is "PrintJobInfo.status"
	//
	// Optional
	Status PrintJobStatus
	// CreationTime is "PrintJobInfo.creationTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CreationTime MUST be set to true to make this field effective.
	CreationTime float64
	// CompletionTime is "PrintJobInfo.completionTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CompletionTime MUST be set to true to make this field effective.
	CompletionTime float64
	// Printer is "PrintJobInfo.printer"
	//
	// Optional
	//
	// NOTE: Printer.FFI_USE MUST be set to true to get Printer used.
	Printer Printer
	// Settings is "PrintJobInfo.settings"
	//
	// Optional
	//
	// NOTE: Settings.FFI_USE MUST be set to true to get Settings used.
	Settings PrintSettings
	// NumberOfPages is "PrintJobInfo.numberOfPages"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumberOfPages MUST be set to true to make this field effective.
	NumberOfPages int32
	// PrinterStatus is "PrintJobInfo.printer_status"
	//
	// Optional
	PrinterStatus printing.PrinterStatus

	FFI_USE_CreationTime   bool // for CreationTime.
	FFI_USE_CompletionTime bool // for CompletionTime.
	FFI_USE_NumberOfPages  bool // for NumberOfPages.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrintJobInfo with all fields set.
func (p PrintJobInfo) FromRef(ref js.Ref) PrintJobInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrintJobInfo in the application heap.
func (p PrintJobInfo) New() js.Ref {
	return bindings.PrintJobInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrintJobInfo) UpdateFrom(ref js.Ref) {
	bindings.PrintJobInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrintJobInfo) Update(ref js.Ref) {
	bindings.PrintJobInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrintJobInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Title.Ref(),
		p.SourceId.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.SourceId = p.SourceId.FromRef(js.Undefined)
	if recursive {
		p.Printer.FreeMembers(true)
		p.Settings.FreeMembers(true)
	}
}

// HasFuncGetPrintJobs returns true if the function "WEBEXT.printingMetrics.getPrintJobs" exists.
func HasFuncGetPrintJobs() bool {
	return js.True == bindings.HasFuncGetPrintJobs()
}

// FuncGetPrintJobs returns the function "WEBEXT.printingMetrics.getPrintJobs".
func FuncGetPrintJobs() (fn js.Func[func() js.Promise[js.Array[PrintJobInfo]]]) {
	bindings.FuncGetPrintJobs(
		js.Pointer(&fn),
	)
	return
}

// GetPrintJobs calls the function "WEBEXT.printingMetrics.getPrintJobs" directly.
func GetPrintJobs() (ret js.Promise[js.Array[PrintJobInfo]]) {
	bindings.CallGetPrintJobs(
		js.Pointer(&ret),
	)

	return
}

// TryGetPrintJobs calls the function "WEBEXT.printingMetrics.getPrintJobs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrintJobs() (ret js.Promise[js.Array[PrintJobInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrintJobs(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnPrintJobFinishedEventCallbackFunc func(this js.Ref, jobInfo *PrintJobInfo) js.Ref

func (fn OnPrintJobFinishedEventCallbackFunc) Register() js.Func[func(jobInfo *PrintJobInfo)] {
	return js.RegisterCallback[func(jobInfo *PrintJobInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPrintJobFinishedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrintJobInfo
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

type OnPrintJobFinishedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, jobInfo *PrintJobInfo) js.Ref
	Arg T
}

func (cb *OnPrintJobFinishedEventCallback[T]) Register() js.Func[func(jobInfo *PrintJobInfo)] {
	return js.RegisterCallback[func(jobInfo *PrintJobInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPrintJobFinishedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PrintJobInfo
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

// HasFuncOnPrintJobFinished returns true if the function "WEBEXT.printingMetrics.onPrintJobFinished.addListener" exists.
func HasFuncOnPrintJobFinished() bool {
	return js.True == bindings.HasFuncOnPrintJobFinished()
}

// FuncOnPrintJobFinished returns the function "WEBEXT.printingMetrics.onPrintJobFinished.addListener".
func FuncOnPrintJobFinished() (fn js.Func[func(callback js.Func[func(jobInfo *PrintJobInfo)])]) {
	bindings.FuncOnPrintJobFinished(
		js.Pointer(&fn),
	)
	return
}

// OnPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.addListener" directly.
func OnPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret js.Void) {
	bindings.CallOnPrintJobFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPrintJobFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPrintJobFinished returns true if the function "WEBEXT.printingMetrics.onPrintJobFinished.removeListener" exists.
func HasFuncOffPrintJobFinished() bool {
	return js.True == bindings.HasFuncOffPrintJobFinished()
}

// FuncOffPrintJobFinished returns the function "WEBEXT.printingMetrics.onPrintJobFinished.removeListener".
func FuncOffPrintJobFinished() (fn js.Func[func(callback js.Func[func(jobInfo *PrintJobInfo)])]) {
	bindings.FuncOffPrintJobFinished(
		js.Pointer(&fn),
	)
	return
}

// OffPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.removeListener" directly.
func OffPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret js.Void) {
	bindings.CallOffPrintJobFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPrintJobFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPrintJobFinished returns true if the function "WEBEXT.printingMetrics.onPrintJobFinished.hasListener" exists.
func HasFuncHasOnPrintJobFinished() bool {
	return js.True == bindings.HasFuncHasOnPrintJobFinished()
}

// FuncHasOnPrintJobFinished returns the function "WEBEXT.printingMetrics.onPrintJobFinished.hasListener".
func FuncHasOnPrintJobFinished() (fn js.Func[func(callback js.Func[func(jobInfo *PrintJobInfo)]) bool]) {
	bindings.FuncHasOnPrintJobFinished(
		js.Pointer(&fn),
	)
	return
}

// HasOnPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.hasListener" directly.
func HasOnPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret bool) {
	bindings.CallHasOnPrintJobFinished(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPrintJobFinished calls the function "WEBEXT.printingMetrics.onPrintJobFinished.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPrintJobFinished(callback js.Func[func(jobInfo *PrintJobInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPrintJobFinished(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

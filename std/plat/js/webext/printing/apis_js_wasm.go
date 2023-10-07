// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package printing

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/printerprovider"
	"github.com/primecitizens/pcz/std/plat/js/webext/printing/bindings"
)

type CancelJobCallbackFunc func(this js.Ref) js.Ref

func (fn CancelJobCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CancelJobCallbackFunc) DispatchCallback(
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

type CancelJobCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CancelJobCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CancelJobCallback[T]) DispatchCallback(
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

type PrinterStatus uint32

const (
	_ PrinterStatus = iota

	PrinterStatus_DOOR_OPEN
	PrinterStatus_TRAY_MISSING
	PrinterStatus_OUT_OF_INK
	PrinterStatus_OUT_OF_PAPER
	PrinterStatus_OUTPUT_FULL
	PrinterStatus_PAPER_JAM
	PrinterStatus_GENERIC_ISSUE
	PrinterStatus_STOPPED
	PrinterStatus_UNREACHABLE
	PrinterStatus_EXPIRED_CERTIFICATE
	PrinterStatus_AVAILABLE
)

func (PrinterStatus) FromRef(str js.Ref) PrinterStatus {
	return PrinterStatus(bindings.ConstOfPrinterStatus(str))
}

func (x PrinterStatus) String() (string, bool) {
	switch x {
	case PrinterStatus_DOOR_OPEN:
		return "DOOR_OPEN", true
	case PrinterStatus_TRAY_MISSING:
		return "TRAY_MISSING", true
	case PrinterStatus_OUT_OF_INK:
		return "OUT_OF_INK", true
	case PrinterStatus_OUT_OF_PAPER:
		return "OUT_OF_PAPER", true
	case PrinterStatus_OUTPUT_FULL:
		return "OUTPUT_FULL", true
	case PrinterStatus_PAPER_JAM:
		return "PAPER_JAM", true
	case PrinterStatus_GENERIC_ISSUE:
		return "GENERIC_ISSUE", true
	case PrinterStatus_STOPPED:
		return "STOPPED", true
	case PrinterStatus_UNREACHABLE:
		return "UNREACHABLE", true
	case PrinterStatus_EXPIRED_CERTIFICATE:
		return "EXPIRED_CERTIFICATE", true
	case PrinterStatus_AVAILABLE:
		return "AVAILABLE", true
	default:
		return "", false
	}
}

type GetPrinterInfoCallbackFunc func(this js.Ref, response *GetPrinterInfoResponse) js.Ref

func (fn GetPrinterInfoCallbackFunc) Register() js.Func[func(response *GetPrinterInfoResponse)] {
	return js.RegisterCallback[func(response *GetPrinterInfoResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPrinterInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetPrinterInfoResponse
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

type GetPrinterInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *GetPrinterInfoResponse) js.Ref
	Arg T
}

func (cb *GetPrinterInfoCallback[T]) Register() js.Func[func(response *GetPrinterInfoResponse)] {
	return js.RegisterCallback[func(response *GetPrinterInfoResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPrinterInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetPrinterInfoResponse
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

type GetPrinterInfoResponse struct {
	// Capabilities is "GetPrinterInfoResponse.capabilities"
	//
	// Optional
	Capabilities js.Object
	// Status is "GetPrinterInfoResponse.status"
	//
	// Optional
	Status PrinterStatus

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetPrinterInfoResponse with all fields set.
func (p GetPrinterInfoResponse) FromRef(ref js.Ref) GetPrinterInfoResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetPrinterInfoResponse in the application heap.
func (p GetPrinterInfoResponse) New() js.Ref {
	return bindings.GetPrinterInfoResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetPrinterInfoResponse) UpdateFrom(ref js.Ref) {
	bindings.GetPrinterInfoResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetPrinterInfoResponse) Update(ref js.Ref) {
	bindings.GetPrinterInfoResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetPrinterInfoResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Capabilities.Ref(),
	)
	p.Capabilities = p.Capabilities.FromRef(js.Undefined)
}

type GetPrintersCallbackFunc func(this js.Ref, printers js.Array[Printer]) js.Ref

func (fn GetPrintersCallbackFunc) Register() js.Func[func(printers js.Array[Printer])] {
	return js.RegisterCallback[func(printers js.Array[Printer])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPrintersCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Printer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPrintersCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printers js.Array[Printer]) js.Ref
	Arg T
}

func (cb *GetPrintersCallback[T]) Register() js.Func[func(printers js.Array[Printer])] {
	return js.RegisterCallback[func(printers js.Array[Printer])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPrintersCallback[T]) DispatchCallback(
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

		js.Array[Printer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
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
	// Id is "Printer.id"
	//
	// Optional
	Id js.String
	// Name is "Printer.name"
	//
	// Optional
	Name js.String
	// Description is "Printer.description"
	//
	// Optional
	Description js.String
	// Uri is "Printer.uri"
	//
	// Optional
	Uri js.String
	// Source is "Printer.source"
	//
	// Optional
	Source PrinterSource
	// IsDefault is "Printer.isDefault"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDefault MUST be set to true to make this field effective.
	IsDefault bool
	// RecentlyUsedRank is "Printer.recentlyUsedRank"
	//
	// Optional
	//
	// NOTE: FFI_USE_RecentlyUsedRank MUST be set to true to make this field effective.
	RecentlyUsedRank int32

	FFI_USE_IsDefault        bool // for IsDefault.
	FFI_USE_RecentlyUsedRank bool // for RecentlyUsedRank.

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
		p.Id.Ref(),
		p.Name.Ref(),
		p.Description.Ref(),
		p.Uri.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Uri = p.Uri.FromRef(js.Undefined)
}

type JobStatus uint32

const (
	_ JobStatus = iota

	JobStatus_PENDING
	JobStatus_IN_PROGRESS
	JobStatus_FAILED
	JobStatus_CANCELED
	JobStatus_PRINTED
)

func (JobStatus) FromRef(str js.Ref) JobStatus {
	return JobStatus(bindings.ConstOfJobStatus(str))
}

func (x JobStatus) String() (string, bool) {
	switch x {
	case JobStatus_PENDING:
		return "PENDING", true
	case JobStatus_IN_PROGRESS:
		return "IN_PROGRESS", true
	case JobStatus_FAILED:
		return "FAILED", true
	case JobStatus_CANCELED:
		return "CANCELED", true
	case JobStatus_PRINTED:
		return "PRINTED", true
	default:
		return "", false
	}
}

type Properties struct {
	ref js.Ref
}

func (this Properties) Once() Properties {
	this.ref.Once()
	return this
}

func (this Properties) Ref() js.Ref {
	return this.ref
}

func (this Properties) FromRef(ref js.Ref) Properties {
	this.ref = ref
	return this
}

func (this Properties) Free() {
	this.ref.Free()
}

// HasFuncMAX_SUBMIT_JOB_CALLS_PER_MINUTE returns true if the static method "Properties.MAX_SUBMIT_JOB_CALLS_PER_MINUTE" exists.
func (this Properties) HasFuncMAX_SUBMIT_JOB_CALLS_PER_MINUTE() bool {
	return js.True == bindings.HasFuncPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
		this.ref,
	)
}

// FuncMAX_SUBMIT_JOB_CALLS_PER_MINUTE returns the static method "Properties.MAX_SUBMIT_JOB_CALLS_PER_MINUTE".
func (this Properties) FuncMAX_SUBMIT_JOB_CALLS_PER_MINUTE() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_SUBMIT_JOB_CALLS_PER_MINUTE calls the static method "Properties.MAX_SUBMIT_JOB_CALLS_PER_MINUTE".
func (this Properties) MAX_SUBMIT_JOB_CALLS_PER_MINUTE() (ret int32) {
	bindings.CallPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_SUBMIT_JOB_CALLS_PER_MINUTE calls the static method "Properties.MAX_SUBMIT_JOB_CALLS_PER_MINUTE"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_SUBMIT_JOB_CALLS_PER_MINUTE() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE returns true if the static method "Properties.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE" exists.
func (this Properties) HasFuncMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE() bool {
	return js.True == bindings.HasFuncPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
		this.ref,
	)
}

// FuncMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE returns the static method "Properties.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE".
func (this Properties) FuncMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE calls the static method "Properties.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE".
func (this Properties) MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE() (ret int32) {
	bindings.CallPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE calls the static method "Properties.MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SubmitJobCallbackFunc func(this js.Ref, response *SubmitJobResponse) js.Ref

func (fn SubmitJobCallbackFunc) Register() js.Func[func(response *SubmitJobResponse)] {
	return js.RegisterCallback[func(response *SubmitJobResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SubmitJobCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SubmitJobResponse
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

type SubmitJobCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *SubmitJobResponse) js.Ref
	Arg T
}

func (cb *SubmitJobCallback[T]) Register() js.Func[func(response *SubmitJobResponse)] {
	return js.RegisterCallback[func(response *SubmitJobResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SubmitJobCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SubmitJobResponse
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

type SubmitJobStatus uint32

const (
	_ SubmitJobStatus = iota

	SubmitJobStatus_OK
	SubmitJobStatus_USER_REJECTED
)

func (SubmitJobStatus) FromRef(str js.Ref) SubmitJobStatus {
	return SubmitJobStatus(bindings.ConstOfSubmitJobStatus(str))
}

func (x SubmitJobStatus) String() (string, bool) {
	switch x {
	case SubmitJobStatus_OK:
		return "OK", true
	case SubmitJobStatus_USER_REJECTED:
		return "USER_REJECTED", true
	default:
		return "", false
	}
}

type SubmitJobResponse struct {
	// Status is "SubmitJobResponse.status"
	//
	// Optional
	Status SubmitJobStatus
	// JobId is "SubmitJobResponse.jobId"
	//
	// Optional
	JobId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SubmitJobResponse with all fields set.
func (p SubmitJobResponse) FromRef(ref js.Ref) SubmitJobResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SubmitJobResponse in the application heap.
func (p SubmitJobResponse) New() js.Ref {
	return bindings.SubmitJobResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SubmitJobResponse) UpdateFrom(ref js.Ref) {
	bindings.SubmitJobResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SubmitJobResponse) Update(ref js.Ref) {
	bindings.SubmitJobResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SubmitJobResponse) FreeMembers(recursive bool) {
	js.Free(
		p.JobId.Ref(),
	)
	p.JobId = p.JobId.FromRef(js.Undefined)
}

type SubmitJobRequest struct {
	// Job is "SubmitJobRequest.job"
	//
	// Optional
	//
	// NOTE: Job.FFI_USE MUST be set to true to get Job used.
	Job printerprovider.PrintJob
	// DocumentBlobUuid is "SubmitJobRequest.documentBlobUuid"
	//
	// Optional
	DocumentBlobUuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SubmitJobRequest with all fields set.
func (p SubmitJobRequest) FromRef(ref js.Ref) SubmitJobRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SubmitJobRequest in the application heap.
func (p SubmitJobRequest) New() js.Ref {
	return bindings.SubmitJobRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SubmitJobRequest) UpdateFrom(ref js.Ref) {
	bindings.SubmitJobRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SubmitJobRequest) Update(ref js.Ref) {
	bindings.SubmitJobRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SubmitJobRequest) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentBlobUuid.Ref(),
	)
	p.DocumentBlobUuid = p.DocumentBlobUuid.FromRef(js.Undefined)
	if recursive {
		p.Job.FreeMembers(true)
	}
}

// HasFuncCancelJob returns true if the function "WEBEXT.printing.cancelJob" exists.
func HasFuncCancelJob() bool {
	return js.True == bindings.HasFuncCancelJob()
}

// FuncCancelJob returns the function "WEBEXT.printing.cancelJob".
func FuncCancelJob() (fn js.Func[func(jobId js.String) js.Promise[js.Void]]) {
	bindings.FuncCancelJob(
		js.Pointer(&fn),
	)
	return
}

// CancelJob calls the function "WEBEXT.printing.cancelJob" directly.
func CancelJob(jobId js.String) (ret js.Promise[js.Void]) {
	bindings.CallCancelJob(
		js.Pointer(&ret),
		jobId.Ref(),
	)

	return
}

// TryCancelJob calls the function "WEBEXT.printing.cancelJob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelJob(jobId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelJob(
		js.Pointer(&ret), js.Pointer(&exception),
		jobId.Ref(),
	)

	return
}

// HasFuncGetPrinterInfo returns true if the function "WEBEXT.printing.getPrinterInfo" exists.
func HasFuncGetPrinterInfo() bool {
	return js.True == bindings.HasFuncGetPrinterInfo()
}

// FuncGetPrinterInfo returns the function "WEBEXT.printing.getPrinterInfo".
func FuncGetPrinterInfo() (fn js.Func[func(printerId js.String) js.Promise[GetPrinterInfoResponse]]) {
	bindings.FuncGetPrinterInfo(
		js.Pointer(&fn),
	)
	return
}

// GetPrinterInfo calls the function "WEBEXT.printing.getPrinterInfo" directly.
func GetPrinterInfo(printerId js.String) (ret js.Promise[GetPrinterInfoResponse]) {
	bindings.CallGetPrinterInfo(
		js.Pointer(&ret),
		printerId.Ref(),
	)

	return
}

// TryGetPrinterInfo calls the function "WEBEXT.printing.getPrinterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrinterInfo(printerId js.String) (ret js.Promise[GetPrinterInfoResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrinterInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		printerId.Ref(),
	)

	return
}

// HasFuncGetPrinters returns true if the function "WEBEXT.printing.getPrinters" exists.
func HasFuncGetPrinters() bool {
	return js.True == bindings.HasFuncGetPrinters()
}

// FuncGetPrinters returns the function "WEBEXT.printing.getPrinters".
func FuncGetPrinters() (fn js.Func[func() js.Promise[js.Array[Printer]]]) {
	bindings.FuncGetPrinters(
		js.Pointer(&fn),
	)
	return
}

// GetPrinters calls the function "WEBEXT.printing.getPrinters" directly.
func GetPrinters() (ret js.Promise[js.Array[Printer]]) {
	bindings.CallGetPrinters(
		js.Pointer(&ret),
	)

	return
}

// TryGetPrinters calls the function "WEBEXT.printing.getPrinters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrinters() (ret js.Promise[js.Array[Printer]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrinters(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnJobStatusChangedEventCallbackFunc func(this js.Ref, jobId js.String, status JobStatus) js.Ref

func (fn OnJobStatusChangedEventCallbackFunc) Register() js.Func[func(jobId js.String, status JobStatus)] {
	return js.RegisterCallback[func(jobId js.String, status JobStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnJobStatusChangedEventCallbackFunc) DispatchCallback(
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
		JobStatus(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnJobStatusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, jobId js.String, status JobStatus) js.Ref
	Arg T
}

func (cb *OnJobStatusChangedEventCallback[T]) Register() js.Func[func(jobId js.String, status JobStatus)] {
	return js.RegisterCallback[func(jobId js.String, status JobStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnJobStatusChangedEventCallback[T]) DispatchCallback(
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
		JobStatus(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnJobStatusChanged returns true if the function "WEBEXT.printing.onJobStatusChanged.addListener" exists.
func HasFuncOnJobStatusChanged() bool {
	return js.True == bindings.HasFuncOnJobStatusChanged()
}

// FuncOnJobStatusChanged returns the function "WEBEXT.printing.onJobStatusChanged.addListener".
func FuncOnJobStatusChanged() (fn js.Func[func(callback js.Func[func(jobId js.String, status JobStatus)])]) {
	bindings.FuncOnJobStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.addListener" directly.
func OnJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret js.Void) {
	bindings.CallOnJobStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnJobStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffJobStatusChanged returns true if the function "WEBEXT.printing.onJobStatusChanged.removeListener" exists.
func HasFuncOffJobStatusChanged() bool {
	return js.True == bindings.HasFuncOffJobStatusChanged()
}

// FuncOffJobStatusChanged returns the function "WEBEXT.printing.onJobStatusChanged.removeListener".
func FuncOffJobStatusChanged() (fn js.Func[func(callback js.Func[func(jobId js.String, status JobStatus)])]) {
	bindings.FuncOffJobStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.removeListener" directly.
func OffJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret js.Void) {
	bindings.CallOffJobStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffJobStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnJobStatusChanged returns true if the function "WEBEXT.printing.onJobStatusChanged.hasListener" exists.
func HasFuncHasOnJobStatusChanged() bool {
	return js.True == bindings.HasFuncHasOnJobStatusChanged()
}

// FuncHasOnJobStatusChanged returns the function "WEBEXT.printing.onJobStatusChanged.hasListener".
func FuncHasOnJobStatusChanged() (fn js.Func[func(callback js.Func[func(jobId js.String, status JobStatus)]) bool]) {
	bindings.FuncHasOnJobStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.hasListener" directly.
func HasOnJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret bool) {
	bindings.CallHasOnJobStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnJobStatusChanged calls the function "WEBEXT.printing.onJobStatusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnJobStatusChanged(callback js.Func[func(jobId js.String, status JobStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnJobStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSubmitJob returns true if the function "WEBEXT.printing.submitJob" exists.
func HasFuncSubmitJob() bool {
	return js.True == bindings.HasFuncSubmitJob()
}

// FuncSubmitJob returns the function "WEBEXT.printing.submitJob".
func FuncSubmitJob() (fn js.Func[func(request SubmitJobRequest) js.Promise[SubmitJobResponse]]) {
	bindings.FuncSubmitJob(
		js.Pointer(&fn),
	)
	return
}

// SubmitJob calls the function "WEBEXT.printing.submitJob" directly.
func SubmitJob(request SubmitJobRequest) (ret js.Promise[SubmitJobResponse]) {
	bindings.CallSubmitJob(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TrySubmitJob calls the function "WEBEXT.printing.submitJob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySubmitJob(request SubmitJobRequest) (ret js.Promise[SubmitJobResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubmitJob(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

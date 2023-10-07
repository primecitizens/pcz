// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/printing constof_PrinterStatus
//go:noescape
func ConstOfPrinterStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/printing store_GetPrinterInfoResponse
//go:noescape
func GetPrinterInfoResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printing load_GetPrinterInfoResponse
//go:noescape
func GetPrinterInfoResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing constof_PrinterSource
//go:noescape
func ConstOfPrinterSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/printing store_Printer
//go:noescape
func PrinterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printing load_Printer
//go:noescape
func PrinterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing constof_JobStatus
//go:noescape
func ConstOfJobStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/printing has_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE
//go:noescape
func HasFuncPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing func_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE
//go:noescape
func FuncPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE
//go:noescape
func CallPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/printing try_Properties_MAX_SUBMIT_JOB_CALLS_PER_MINUTE
//go:noescape
func TryPropertiesMAX_SUBMIT_JOB_CALLS_PER_MINUTE(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE
//go:noescape
func HasFuncPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing func_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE
//go:noescape
func FuncPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE
//go:noescape
func CallPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/printing try_Properties_MAX_GET_PRINTER_INFO_CALLS_PER_MINUTE
//go:noescape
func TryPropertiesMAX_GET_PRINTER_INFO_CALLS_PER_MINUTE(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/printing constof_SubmitJobStatus
//go:noescape
func ConstOfSubmitJobStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/printing store_SubmitJobResponse
//go:noescape
func SubmitJobResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printing load_SubmitJobResponse
//go:noescape
func SubmitJobResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing store_SubmitJobRequest
//go:noescape
func SubmitJobRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printing load_SubmitJobRequest
//go:noescape
func SubmitJobRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printing has_CancelJob
//go:noescape
func HasFuncCancelJob() js.Ref

//go:wasmimport plat/js/webext/printing func_CancelJob
//go:noescape
func FuncCancelJob(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_CancelJob
//go:noescape
func CallCancelJob(
	retPtr unsafe.Pointer,
	jobId js.Ref)

//go:wasmimport plat/js/webext/printing try_CancelJob
//go:noescape
func TryCancelJob(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	jobId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_GetPrinterInfo
//go:noescape
func HasFuncGetPrinterInfo() js.Ref

//go:wasmimport plat/js/webext/printing func_GetPrinterInfo
//go:noescape
func FuncGetPrinterInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_GetPrinterInfo
//go:noescape
func CallGetPrinterInfo(
	retPtr unsafe.Pointer,
	printerId js.Ref)

//go:wasmimport plat/js/webext/printing try_GetPrinterInfo
//go:noescape
func TryGetPrinterInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	printerId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_GetPrinters
//go:noescape
func HasFuncGetPrinters() js.Ref

//go:wasmimport plat/js/webext/printing func_GetPrinters
//go:noescape
func FuncGetPrinters(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_GetPrinters
//go:noescape
func CallGetPrinters(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/printing try_GetPrinters
//go:noescape
func TryGetPrinters(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_OnJobStatusChanged
//go:noescape
func HasFuncOnJobStatusChanged() js.Ref

//go:wasmimport plat/js/webext/printing func_OnJobStatusChanged
//go:noescape
func FuncOnJobStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_OnJobStatusChanged
//go:noescape
func CallOnJobStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printing try_OnJobStatusChanged
//go:noescape
func TryOnJobStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_OffJobStatusChanged
//go:noescape
func HasFuncOffJobStatusChanged() js.Ref

//go:wasmimport plat/js/webext/printing func_OffJobStatusChanged
//go:noescape
func FuncOffJobStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_OffJobStatusChanged
//go:noescape
func CallOffJobStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printing try_OffJobStatusChanged
//go:noescape
func TryOffJobStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_HasOnJobStatusChanged
//go:noescape
func HasFuncHasOnJobStatusChanged() js.Ref

//go:wasmimport plat/js/webext/printing func_HasOnJobStatusChanged
//go:noescape
func FuncHasOnJobStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_HasOnJobStatusChanged
//go:noescape
func CallHasOnJobStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printing try_HasOnJobStatusChanged
//go:noescape
func TryHasOnJobStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printing has_SubmitJob
//go:noescape
func HasFuncSubmitJob() js.Ref

//go:wasmimport plat/js/webext/printing func_SubmitJob
//go:noescape
func FuncSubmitJob(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printing call_SubmitJob
//go:noescape
func CallSubmitJob(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/printing try_SubmitJob
//go:noescape
func TrySubmitJob(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/webext/printerproviderinternal constof_PrintError
//go:noescape
func ConstOfPrintError(str js.Ref) uint32

//go:wasmimport plat/js/webext/printerproviderinternal has_GetPrintData
//go:noescape
func HasFuncGetPrintData() js.Ref

//go:wasmimport plat/js/webext/printerproviderinternal func_GetPrintData
//go:noescape
func FuncGetPrintData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal call_GetPrintData
//go:noescape
func CallGetPrintData(
	retPtr unsafe.Pointer,
	requestId int32)

//go:wasmimport plat/js/webext/printerproviderinternal try_GetPrintData
//go:noescape
func TryGetPrintData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal has_ReportPrintResult
//go:noescape
func HasFuncReportPrintResult() js.Ref

//go:wasmimport plat/js/webext/printerproviderinternal func_ReportPrintResult
//go:noescape
func FuncReportPrintResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal call_ReportPrintResult
//go:noescape
func CallReportPrintResult(
	retPtr unsafe.Pointer,
	request_id int32,
	err uint32)

//go:wasmimport plat/js/webext/printerproviderinternal try_ReportPrintResult
//go:noescape
func TryReportPrintResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request_id int32,
	err uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal has_ReportPrinterCapability
//go:noescape
func HasFuncReportPrinterCapability() js.Ref

//go:wasmimport plat/js/webext/printerproviderinternal func_ReportPrinterCapability
//go:noescape
func FuncReportPrinterCapability(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal call_ReportPrinterCapability
//go:noescape
func CallReportPrinterCapability(
	retPtr unsafe.Pointer,
	request_id int32,
	capability js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal try_ReportPrinterCapability
//go:noescape
func TryReportPrinterCapability(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request_id int32,
	capability js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal has_ReportPrinters
//go:noescape
func HasFuncReportPrinters() js.Ref

//go:wasmimport plat/js/webext/printerproviderinternal func_ReportPrinters
//go:noescape
func FuncReportPrinters(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal call_ReportPrinters
//go:noescape
func CallReportPrinters(
	retPtr unsafe.Pointer,
	requestId int32,
	printers js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal try_ReportPrinters
//go:noescape
func TryReportPrinters(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	printers js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerproviderinternal has_ReportUsbPrinterInfo
//go:noescape
func HasFuncReportUsbPrinterInfo() js.Ref

//go:wasmimport plat/js/webext/printerproviderinternal func_ReportUsbPrinterInfo
//go:noescape
func FuncReportUsbPrinterInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal call_ReportUsbPrinterInfo
//go:noescape
func CallReportUsbPrinterInfo(
	retPtr unsafe.Pointer,
	requestId int32,
	printerInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/printerproviderinternal try_ReportUsbPrinterInfo
//go:noescape
func TryReportUsbPrinterInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	printerInfo unsafe.Pointer) (ok js.Ref)

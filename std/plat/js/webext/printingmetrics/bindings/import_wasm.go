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

//go:wasmimport plat/js/webext/printingmetrics constof_ColorMode
//go:noescape
func ConstOfColorMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/printingmetrics constof_DuplexMode
//go:noescape
func ConstOfDuplexMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/printingmetrics constof_PrintJobSource
//go:noescape
func ConstOfPrintJobSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/printingmetrics constof_PrintJobStatus
//go:noescape
func ConstOfPrintJobStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/printingmetrics constof_PrinterSource
//go:noescape
func ConstOfPrinterSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/printingmetrics store_Printer
//go:noescape
func PrinterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printingmetrics load_Printer
//go:noescape
func PrinterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printingmetrics store_MediaSize
//go:noescape
func MediaSizeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printingmetrics load_MediaSize
//go:noescape
func MediaSizeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printingmetrics store_PrintSettings
//go:noescape
func PrintSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printingmetrics load_PrintSettings
//go:noescape
func PrintSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printingmetrics store_PrintJobInfo
//go:noescape
func PrintJobInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printingmetrics load_PrintJobInfo
//go:noescape
func PrintJobInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printingmetrics has_GetPrintJobs
//go:noescape
func HasFuncGetPrintJobs() js.Ref

//go:wasmimport plat/js/webext/printingmetrics func_GetPrintJobs
//go:noescape
func FuncGetPrintJobs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printingmetrics call_GetPrintJobs
//go:noescape
func CallGetPrintJobs(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/printingmetrics try_GetPrintJobs
//go:noescape
func TryGetPrintJobs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/printingmetrics has_OnPrintJobFinished
//go:noescape
func HasFuncOnPrintJobFinished() js.Ref

//go:wasmimport plat/js/webext/printingmetrics func_OnPrintJobFinished
//go:noescape
func FuncOnPrintJobFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printingmetrics call_OnPrintJobFinished
//go:noescape
func CallOnPrintJobFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printingmetrics try_OnPrintJobFinished
//go:noescape
func TryOnPrintJobFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printingmetrics has_OffPrintJobFinished
//go:noescape
func HasFuncOffPrintJobFinished() js.Ref

//go:wasmimport plat/js/webext/printingmetrics func_OffPrintJobFinished
//go:noescape
func FuncOffPrintJobFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printingmetrics call_OffPrintJobFinished
//go:noescape
func CallOffPrintJobFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printingmetrics try_OffPrintJobFinished
//go:noescape
func TryOffPrintJobFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printingmetrics has_HasOnPrintJobFinished
//go:noescape
func HasFuncHasOnPrintJobFinished() js.Ref

//go:wasmimport plat/js/webext/printingmetrics func_HasOnPrintJobFinished
//go:noescape
func FuncHasOnPrintJobFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printingmetrics call_HasOnPrintJobFinished
//go:noescape
func CallHasOnPrintJobFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printingmetrics try_HasOnPrintJobFinished
//go:noescape
func TryHasOnPrintJobFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/printerprovider constof_PrintError
//go:noescape
func ConstOfPrintError(str js.Ref) uint32

//go:wasmimport plat/js/webext/printerprovider store_PrintJob
//go:noescape
func PrintJobJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printerprovider load_PrintJob
//go:noescape
func PrintJobJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printerprovider store_PrinterInfo
//go:noescape
func PrinterInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/printerprovider load_PrinterInfo
//go:noescape
func PrinterInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/printerprovider has_OnGetCapabilityRequested
//go:noescape
func HasFuncOnGetCapabilityRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OnGetCapabilityRequested
//go:noescape
func FuncOnGetCapabilityRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OnGetCapabilityRequested
//go:noescape
func CallOnGetCapabilityRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OnGetCapabilityRequested
//go:noescape
func TryOnGetCapabilityRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OffGetCapabilityRequested
//go:noescape
func HasFuncOffGetCapabilityRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OffGetCapabilityRequested
//go:noescape
func FuncOffGetCapabilityRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OffGetCapabilityRequested
//go:noescape
func CallOffGetCapabilityRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OffGetCapabilityRequested
//go:noescape
func TryOffGetCapabilityRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_HasOnGetCapabilityRequested
//go:noescape
func HasFuncHasOnGetCapabilityRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_HasOnGetCapabilityRequested
//go:noescape
func FuncHasOnGetCapabilityRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_HasOnGetCapabilityRequested
//go:noescape
func CallHasOnGetCapabilityRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_HasOnGetCapabilityRequested
//go:noescape
func TryHasOnGetCapabilityRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OnGetPrintersRequested
//go:noescape
func HasFuncOnGetPrintersRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OnGetPrintersRequested
//go:noescape
func FuncOnGetPrintersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OnGetPrintersRequested
//go:noescape
func CallOnGetPrintersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OnGetPrintersRequested
//go:noescape
func TryOnGetPrintersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OffGetPrintersRequested
//go:noescape
func HasFuncOffGetPrintersRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OffGetPrintersRequested
//go:noescape
func FuncOffGetPrintersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OffGetPrintersRequested
//go:noescape
func CallOffGetPrintersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OffGetPrintersRequested
//go:noescape
func TryOffGetPrintersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_HasOnGetPrintersRequested
//go:noescape
func HasFuncHasOnGetPrintersRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_HasOnGetPrintersRequested
//go:noescape
func FuncHasOnGetPrintersRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_HasOnGetPrintersRequested
//go:noescape
func CallHasOnGetPrintersRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_HasOnGetPrintersRequested
//go:noescape
func TryHasOnGetPrintersRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OnGetUsbPrinterInfoRequested
//go:noescape
func HasFuncOnGetUsbPrinterInfoRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OnGetUsbPrinterInfoRequested
//go:noescape
func FuncOnGetUsbPrinterInfoRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OnGetUsbPrinterInfoRequested
//go:noescape
func CallOnGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OnGetUsbPrinterInfoRequested
//go:noescape
func TryOnGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OffGetUsbPrinterInfoRequested
//go:noescape
func HasFuncOffGetUsbPrinterInfoRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OffGetUsbPrinterInfoRequested
//go:noescape
func FuncOffGetUsbPrinterInfoRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OffGetUsbPrinterInfoRequested
//go:noescape
func CallOffGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OffGetUsbPrinterInfoRequested
//go:noescape
func TryOffGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_HasOnGetUsbPrinterInfoRequested
//go:noescape
func HasFuncHasOnGetUsbPrinterInfoRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_HasOnGetUsbPrinterInfoRequested
//go:noescape
func FuncHasOnGetUsbPrinterInfoRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_HasOnGetUsbPrinterInfoRequested
//go:noescape
func CallHasOnGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_HasOnGetUsbPrinterInfoRequested
//go:noescape
func TryHasOnGetUsbPrinterInfoRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OnPrintRequested
//go:noescape
func HasFuncOnPrintRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OnPrintRequested
//go:noescape
func FuncOnPrintRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OnPrintRequested
//go:noescape
func CallOnPrintRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OnPrintRequested
//go:noescape
func TryOnPrintRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_OffPrintRequested
//go:noescape
func HasFuncOffPrintRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_OffPrintRequested
//go:noescape
func FuncOffPrintRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_OffPrintRequested
//go:noescape
func CallOffPrintRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_OffPrintRequested
//go:noescape
func TryOffPrintRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/printerprovider has_HasOnPrintRequested
//go:noescape
func HasFuncHasOnPrintRequested() js.Ref

//go:wasmimport plat/js/webext/printerprovider func_HasOnPrintRequested
//go:noescape
func FuncHasOnPrintRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/printerprovider call_HasOnPrintRequested
//go:noescape
func CallHasOnPrintRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/printerprovider try_HasOnPrintRequested
//go:noescape
func TryHasOnPrintRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/crashreportprivate store_ErrorInfo
//go:noescape
func ErrorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/crashreportprivate load_ErrorInfo
//go:noescape
func ErrorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/crashreportprivate has_ReportError
//go:noescape
func HasFuncReportError() js.Ref

//go:wasmimport plat/js/webext/crashreportprivate func_ReportError
//go:noescape
func FuncReportError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/crashreportprivate call_ReportError
//go:noescape
func CallReportError(
	retPtr unsafe.Pointer,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/crashreportprivate try_ReportError
//go:noescape
func TryReportError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	info unsafe.Pointer) (ok js.Ref)

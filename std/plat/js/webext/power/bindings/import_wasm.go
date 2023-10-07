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

//go:wasmimport plat/js/webext/power constof_Level
//go:noescape
func ConstOfLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/power has_ReleaseKeepAwake
//go:noescape
func HasFuncReleaseKeepAwake() js.Ref

//go:wasmimport plat/js/webext/power func_ReleaseKeepAwake
//go:noescape
func FuncReleaseKeepAwake(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/power call_ReleaseKeepAwake
//go:noescape
func CallReleaseKeepAwake(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/power try_ReleaseKeepAwake
//go:noescape
func TryReleaseKeepAwake(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/power has_ReportActivity
//go:noescape
func HasFuncReportActivity() js.Ref

//go:wasmimport plat/js/webext/power func_ReportActivity
//go:noescape
func FuncReportActivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/power call_ReportActivity
//go:noescape
func CallReportActivity(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/power try_ReportActivity
//go:noescape
func TryReportActivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/power has_RequestKeepAwake
//go:noescape
func HasFuncRequestKeepAwake() js.Ref

//go:wasmimport plat/js/webext/power func_RequestKeepAwake
//go:noescape
func FuncRequestKeepAwake(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/power call_RequestKeepAwake
//go:noescape
func CallRequestKeepAwake(
	retPtr unsafe.Pointer,
	level uint32)

//go:wasmimport plat/js/webext/power try_RequestKeepAwake
//go:noescape
func TryRequestKeepAwake(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	level uint32) (ok js.Ref)

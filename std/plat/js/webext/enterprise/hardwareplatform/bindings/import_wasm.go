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

//go:wasmimport plat/js/webext/enterprise/hardwareplatform store_HardwarePlatformInfo
//go:noescape
func HardwarePlatformInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/hardwareplatform load_HardwarePlatformInfo
//go:noescape
func HardwarePlatformInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/hardwareplatform has_GetHardwarePlatformInfo
//go:noescape
func HasFuncGetHardwarePlatformInfo() js.Ref

//go:wasmimport plat/js/webext/enterprise/hardwareplatform func_GetHardwarePlatformInfo
//go:noescape
func FuncGetHardwarePlatformInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/hardwareplatform call_GetHardwarePlatformInfo
//go:noescape
func CallGetHardwarePlatformInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/hardwareplatform try_GetHardwarePlatformInfo
//go:noescape
func TryGetHardwarePlatformInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

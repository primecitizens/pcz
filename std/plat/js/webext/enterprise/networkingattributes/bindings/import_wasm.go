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

//go:wasmimport plat/js/webext/enterprise/networkingattributes store_NetworkDetails
//go:noescape
func NetworkDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/networkingattributes load_NetworkDetails
//go:noescape
func NetworkDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/networkingattributes has_GetNetworkDetails
//go:noescape
func HasFuncGetNetworkDetails() js.Ref

//go:wasmimport plat/js/webext/enterprise/networkingattributes func_GetNetworkDetails
//go:noescape
func FuncGetNetworkDetails(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/networkingattributes call_GetNetworkDetails
//go:noescape
func CallGetNetworkDetails(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/networkingattributes try_GetNetworkDetails
//go:noescape
func TryGetNetworkDetails(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

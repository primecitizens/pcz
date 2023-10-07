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

//go:wasmimport plat/js/webext/system/network store_NetworkInterface
//go:noescape
func NetworkInterfaceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/network load_NetworkInterface
//go:noescape
func NetworkInterfaceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/network has_GetNetworkInterfaces
//go:noescape
func HasFuncGetNetworkInterfaces() js.Ref

//go:wasmimport plat/js/webext/system/network func_GetNetworkInterfaces
//go:noescape
func FuncGetNetworkInterfaces(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/network call_GetNetworkInterfaces
//go:noescape
func CallGetNetworkInterfaces(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/system/network try_GetNetworkInterfaces
//go:noescape
func TryGetNetworkInterfaces(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

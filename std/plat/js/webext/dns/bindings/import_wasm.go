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

//go:wasmimport plat/js/webext/dns store_ResolveCallbackResolveInfo
//go:noescape
func ResolveCallbackResolveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/dns load_ResolveCallbackResolveInfo
//go:noescape
func ResolveCallbackResolveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/dns has_Resolve
//go:noescape
func HasFuncResolve() js.Ref

//go:wasmimport plat/js/webext/dns func_Resolve
//go:noescape
func FuncResolve(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/dns call_Resolve
//go:noescape
func CallResolve(
	retPtr unsafe.Pointer,
	hostname js.Ref)

//go:wasmimport plat/js/webext/dns try_Resolve
//go:noescape
func TryResolve(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	hostname js.Ref) (ok js.Ref)

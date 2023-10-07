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

//go:wasmimport plat/js/webext/proxy constof_Mode
//go:noescape
func ConstOfMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/proxy store_OnProxyErrorArgDetails
//go:noescape
func OnProxyErrorArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/proxy load_OnProxyErrorArgDetails
//go:noescape
func OnProxyErrorArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/proxy store_PacScript
//go:noescape
func PacScriptJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/proxy load_PacScript
//go:noescape
func PacScriptJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/proxy constof_Scheme
//go:noescape
func ConstOfScheme(str js.Ref) uint32

//go:wasmimport plat/js/webext/proxy store_ProxyServer
//go:noescape
func ProxyServerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/proxy load_ProxyServer
//go:noescape
func ProxyServerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/proxy store_ProxyRules
//go:noescape
func ProxyRulesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/proxy load_ProxyRules
//go:noescape
func ProxyRulesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/proxy store_ProxyConfig
//go:noescape
func ProxyConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/proxy load_ProxyConfig
//go:noescape
func ProxyConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/proxy has_OnProxyError
//go:noescape
func HasFuncOnProxyError() js.Ref

//go:wasmimport plat/js/webext/proxy func_OnProxyError
//go:noescape
func FuncOnProxyError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/proxy call_OnProxyError
//go:noescape
func CallOnProxyError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/proxy try_OnProxyError
//go:noescape
func TryOnProxyError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/proxy has_OffProxyError
//go:noescape
func HasFuncOffProxyError() js.Ref

//go:wasmimport plat/js/webext/proxy func_OffProxyError
//go:noescape
func FuncOffProxyError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/proxy call_OffProxyError
//go:noescape
func CallOffProxyError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/proxy try_OffProxyError
//go:noescape
func TryOffProxyError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/proxy has_HasOnProxyError
//go:noescape
func HasFuncHasOnProxyError() js.Ref

//go:wasmimport plat/js/webext/proxy func_HasOnProxyError
//go:noescape
func FuncHasOnProxyError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/proxy call_HasOnProxyError
//go:noescape
func CallHasOnProxyError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/proxy try_HasOnProxyError
//go:noescape
func TryHasOnProxyError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/proxy get_Settings
//go:noescape
func GetSettings(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/proxy set_Settings
//go:noescape
func SetSettings(
	val js.Ref) js.Ref

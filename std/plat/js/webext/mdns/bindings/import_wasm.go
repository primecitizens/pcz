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

//go:wasmimport plat/js/webext/mdns store_MDnsService
//go:noescape
func MDnsServiceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mdns load_MDnsService
//go:noescape
func MDnsServiceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mdns has_Properties_MAX_SERVICE_INSTANCES_PER_EVENT
//go:noescape
func HasFuncPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/mdns func_Properties_MAX_SERVICE_INSTANCES_PER_EVENT
//go:noescape
func FuncPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns call_Properties_MAX_SERVICE_INSTANCES_PER_EVENT
//go:noescape
func CallPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns try_Properties_MAX_SERVICE_INSTANCES_PER_EVENT
//go:noescape
func TryPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mdns has_ForceDiscovery
//go:noescape
func HasFuncForceDiscovery() js.Ref

//go:wasmimport plat/js/webext/mdns func_ForceDiscovery
//go:noescape
func FuncForceDiscovery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns call_ForceDiscovery
//go:noescape
func CallForceDiscovery(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns try_ForceDiscovery
//go:noescape
func TryForceDiscovery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mdns has_OnServiceList
//go:noescape
func HasFuncOnServiceList() js.Ref

//go:wasmimport plat/js/webext/mdns func_OnServiceList
//go:noescape
func FuncOnServiceList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns call_OnServiceList
//go:noescape
func CallOnServiceList(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mdns try_OnServiceList
//go:noescape
func TryOnServiceList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mdns has_OffServiceList
//go:noescape
func HasFuncOffServiceList() js.Ref

//go:wasmimport plat/js/webext/mdns func_OffServiceList
//go:noescape
func FuncOffServiceList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns call_OffServiceList
//go:noescape
func CallOffServiceList(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mdns try_OffServiceList
//go:noescape
func TryOffServiceList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mdns has_HasOnServiceList
//go:noescape
func HasFuncHasOnServiceList() js.Ref

//go:wasmimport plat/js/webext/mdns func_HasOnServiceList
//go:noescape
func FuncHasOnServiceList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mdns call_HasOnServiceList
//go:noescape
func CallHasOnServiceList(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mdns try_HasOnServiceList
//go:noescape
func TryHasOnServiceList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

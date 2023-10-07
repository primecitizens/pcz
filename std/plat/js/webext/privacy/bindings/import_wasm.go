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

//go:wasmimport plat/js/webext/privacy constof_IPHandlingPolicy
//go:noescape
func ConstOfIPHandlingPolicy(str js.Ref) uint32

//go:wasmimport plat/js/webext/privacy store_NetworkProperty
//go:noescape
func NetworkPropertyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/privacy load_NetworkProperty
//go:noescape
func NetworkPropertyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/privacy store_ServicesProperty
//go:noescape
func ServicesPropertyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/privacy load_ServicesProperty
//go:noescape
func ServicesPropertyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/privacy store_WebsitesProperty
//go:noescape
func WebsitesPropertyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/privacy load_WebsitesProperty
//go:noescape
func WebsitesPropertyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/privacy get_Network
//go:noescape
func GetNetwork(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/privacy set_Network
//go:noescape
func SetNetwork(
	val unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/privacy get_Services
//go:noescape
func GetServices(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/privacy set_Services
//go:noescape
func SetServices(
	val unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/privacy get_Websites
//go:noescape
func GetWebsites(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/privacy set_Websites
//go:noescape
func SetWebsites(
	val unsafe.Pointer) js.Ref

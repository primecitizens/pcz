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

//go:wasmimport plat/js/webext/guestviewinternal store_Size
//go:noescape
func SizeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/guestviewinternal load_Size
//go:noescape
func SizeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/guestviewinternal store_SizeParams
//go:noescape
func SizeParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/guestviewinternal load_SizeParams
//go:noescape
func SizeParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/guestviewinternal has_CreateGuest
//go:noescape
func HasFuncCreateGuest() js.Ref

//go:wasmimport plat/js/webext/guestviewinternal func_CreateGuest
//go:noescape
func FuncCreateGuest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/guestviewinternal call_CreateGuest
//go:noescape
func CallCreateGuest(
	retPtr unsafe.Pointer,
	viewType js.Ref,
	ownerRoutingId float64,
	createParams js.Ref)

//go:wasmimport plat/js/webext/guestviewinternal try_CreateGuest
//go:noescape
func TryCreateGuest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	viewType js.Ref,
	ownerRoutingId float64,
	createParams js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/guestviewinternal has_DestroyUnattachedGuest
//go:noescape
func HasFuncDestroyUnattachedGuest() js.Ref

//go:wasmimport plat/js/webext/guestviewinternal func_DestroyUnattachedGuest
//go:noescape
func FuncDestroyUnattachedGuest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/guestviewinternal call_DestroyUnattachedGuest
//go:noescape
func CallDestroyUnattachedGuest(
	retPtr unsafe.Pointer,
	instanceId float64)

//go:wasmimport plat/js/webext/guestviewinternal try_DestroyUnattachedGuest
//go:noescape
func TryDestroyUnattachedGuest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/guestviewinternal has_SetSize
//go:noescape
func HasFuncSetSize() js.Ref

//go:wasmimport plat/js/webext/guestviewinternal func_SetSize
//go:noescape
func FuncSetSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/guestviewinternal call_SetSize
//go:noescape
func CallSetSize(
	retPtr unsafe.Pointer,
	instanceId float64,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/guestviewinternal try_SetSize
//go:noescape
func TrySetSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	params unsafe.Pointer) (ok js.Ref)

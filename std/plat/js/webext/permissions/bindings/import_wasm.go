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

//go:wasmimport plat/js/webext/permissions store_Permissions
//go:noescape
func PermissionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/permissions load_Permissions
//go:noescape
func PermissionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/permissions has_Contains
//go:noescape
func HasFuncContains() js.Ref

//go:wasmimport plat/js/webext/permissions func_Contains
//go:noescape
func FuncContains(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_Contains
//go:noescape
func CallContains(
	retPtr unsafe.Pointer,
	permissions unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions try_Contains
//go:noescape
func TryContains(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/permissions func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_OnAdded
//go:noescape
func HasFuncOnAdded() js.Ref

//go:wasmimport plat/js/webext/permissions func_OnAdded
//go:noescape
func FuncOnAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_OnAdded
//go:noescape
func CallOnAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_OnAdded
//go:noescape
func TryOnAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_OffAdded
//go:noescape
func HasFuncOffAdded() js.Ref

//go:wasmimport plat/js/webext/permissions func_OffAdded
//go:noescape
func FuncOffAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_OffAdded
//go:noescape
func CallOffAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_OffAdded
//go:noescape
func TryOffAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_HasOnAdded
//go:noescape
func HasFuncHasOnAdded() js.Ref

//go:wasmimport plat/js/webext/permissions func_HasOnAdded
//go:noescape
func FuncHasOnAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_HasOnAdded
//go:noescape
func CallHasOnAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_HasOnAdded
//go:noescape
func TryHasOnAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_OnRemoved
//go:noescape
func HasFuncOnRemoved() js.Ref

//go:wasmimport plat/js/webext/permissions func_OnRemoved
//go:noescape
func FuncOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_OnRemoved
//go:noescape
func CallOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_OnRemoved
//go:noescape
func TryOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_OffRemoved
//go:noescape
func HasFuncOffRemoved() js.Ref

//go:wasmimport plat/js/webext/permissions func_OffRemoved
//go:noescape
func FuncOffRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_OffRemoved
//go:noescape
func CallOffRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_OffRemoved
//go:noescape
func TryOffRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_HasOnRemoved
//go:noescape
func HasFuncHasOnRemoved() js.Ref

//go:wasmimport plat/js/webext/permissions func_HasOnRemoved
//go:noescape
func FuncHasOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_HasOnRemoved
//go:noescape
func CallHasOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/permissions try_HasOnRemoved
//go:noescape
func TryHasOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/permissions func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	permissions unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissions unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/permissions has_Request
//go:noescape
func HasFuncRequest() js.Ref

//go:wasmimport plat/js/webext/permissions func_Request
//go:noescape
func FuncRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions call_Request
//go:noescape
func CallRequest(
	retPtr unsafe.Pointer,
	permissions unsafe.Pointer)

//go:wasmimport plat/js/webext/permissions try_Request
//go:noescape
func TryRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissions unsafe.Pointer) (ok js.Ref)

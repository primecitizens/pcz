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

//go:wasmimport plat/js/webext/usersprivate store_LoginStatusDict
//go:noescape
func LoginStatusDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usersprivate load_LoginStatusDict
//go:noescape
func LoginStatusDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usersprivate store_User
//go:noescape
func UserJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/usersprivate load_User
//go:noescape
func UserJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/usersprivate has_AddUser
//go:noescape
func HasFuncAddUser() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_AddUser
//go:noescape
func FuncAddUser(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_AddUser
//go:noescape
func CallAddUser(
	retPtr unsafe.Pointer,
	email js.Ref)

//go:wasmimport plat/js/webext/usersprivate try_AddUser
//go:noescape
func TryAddUser(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	email js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_GetCurrentUser
//go:noescape
func HasFuncGetCurrentUser() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_GetCurrentUser
//go:noescape
func FuncGetCurrentUser(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_GetCurrentUser
//go:noescape
func CallGetCurrentUser(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate try_GetCurrentUser
//go:noescape
func TryGetCurrentUser(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_GetLoginStatus
//go:noescape
func HasFuncGetLoginStatus() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_GetLoginStatus
//go:noescape
func FuncGetLoginStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_GetLoginStatus
//go:noescape
func CallGetLoginStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate try_GetLoginStatus
//go:noescape
func TryGetLoginStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_GetUsers
//go:noescape
func HasFuncGetUsers() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_GetUsers
//go:noescape
func FuncGetUsers(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_GetUsers
//go:noescape
func CallGetUsers(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate try_GetUsers
//go:noescape
func TryGetUsers(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_IsUserInList
//go:noescape
func HasFuncIsUserInList() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_IsUserInList
//go:noescape
func FuncIsUserInList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_IsUserInList
//go:noescape
func CallIsUserInList(
	retPtr unsafe.Pointer,
	email js.Ref)

//go:wasmimport plat/js/webext/usersprivate try_IsUserInList
//go:noescape
func TryIsUserInList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	email js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_IsUserListManaged
//go:noescape
func HasFuncIsUserListManaged() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_IsUserListManaged
//go:noescape
func FuncIsUserListManaged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_IsUserListManaged
//go:noescape
func CallIsUserListManaged(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate try_IsUserListManaged
//go:noescape
func TryIsUserListManaged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/usersprivate has_RemoveUser
//go:noescape
func HasFuncRemoveUser() js.Ref

//go:wasmimport plat/js/webext/usersprivate func_RemoveUser
//go:noescape
func FuncRemoveUser(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/usersprivate call_RemoveUser
//go:noescape
func CallRemoveUser(
	retPtr unsafe.Pointer,
	email js.Ref)

//go:wasmimport plat/js/webext/usersprivate try_RemoveUser
//go:noescape
func TryRemoveUser(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	email js.Ref) (ok js.Ref)

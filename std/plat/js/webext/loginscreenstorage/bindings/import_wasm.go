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

//go:wasmimport plat/js/webext/loginscreenstorage has_RetrieveCredentials
//go:noescape
func HasFuncRetrieveCredentials() js.Ref

//go:wasmimport plat/js/webext/loginscreenstorage func_RetrieveCredentials
//go:noescape
func FuncRetrieveCredentials(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenstorage call_RetrieveCredentials
//go:noescape
func CallRetrieveCredentials(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenstorage try_RetrieveCredentials
//go:noescape
func TryRetrieveCredentials(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage has_RetrievePersistentData
//go:noescape
func HasFuncRetrievePersistentData() js.Ref

//go:wasmimport plat/js/webext/loginscreenstorage func_RetrievePersistentData
//go:noescape
func FuncRetrievePersistentData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenstorage call_RetrievePersistentData
//go:noescape
func CallRetrievePersistentData(
	retPtr unsafe.Pointer,
	ownerId js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage try_RetrievePersistentData
//go:noescape
func TryRetrievePersistentData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ownerId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage has_StoreCredentials
//go:noescape
func HasFuncStoreCredentials() js.Ref

//go:wasmimport plat/js/webext/loginscreenstorage func_StoreCredentials
//go:noescape
func FuncStoreCredentials(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenstorage call_StoreCredentials
//go:noescape
func CallStoreCredentials(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	credentials js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage try_StoreCredentials
//go:noescape
func TryStoreCredentials(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	credentials js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage has_StorePersistentData
//go:noescape
func HasFuncStorePersistentData() js.Ref

//go:wasmimport plat/js/webext/loginscreenstorage func_StorePersistentData
//go:noescape
func FuncStorePersistentData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenstorage call_StorePersistentData
//go:noescape
func CallStorePersistentData(
	retPtr unsafe.Pointer,
	extensionIds js.Ref,
	data js.Ref)

//go:wasmimport plat/js/webext/loginscreenstorage try_StorePersistentData
//go:noescape
func TryStorePersistentData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionIds js.Ref,
	data js.Ref) (ok js.Ref)

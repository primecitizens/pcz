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

//go:wasmimport plat/js/webext/browsingdata store_DataTypeSet
//go:noescape
func DataTypeSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/browsingdata load_DataTypeSet
//go:noescape
func DataTypeSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/browsingdata store_RemovalOptionsFieldOriginTypes
//go:noescape
func RemovalOptionsFieldOriginTypesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/browsingdata load_RemovalOptionsFieldOriginTypes
//go:noescape
func RemovalOptionsFieldOriginTypesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/browsingdata store_RemovalOptions
//go:noescape
func RemovalOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/browsingdata load_RemovalOptions
//go:noescape
func RemovalOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/browsingdata store_SettingsReturnType
//go:noescape
func SettingsReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/browsingdata load_SettingsReturnType
//go:noescape
func SettingsReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/browsingdata has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	dataToRemove unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	dataToRemove unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveAppcache
//go:noescape
func HasFuncRemoveAppcache() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveAppcache
//go:noescape
func FuncRemoveAppcache(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveAppcache
//go:noescape
func CallRemoveAppcache(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveAppcache
//go:noescape
func TryRemoveAppcache(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveCache
//go:noescape
func HasFuncRemoveCache() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveCache
//go:noescape
func FuncRemoveCache(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveCache
//go:noescape
func CallRemoveCache(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveCache
//go:noescape
func TryRemoveCache(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveCacheStorage
//go:noescape
func HasFuncRemoveCacheStorage() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveCacheStorage
//go:noescape
func FuncRemoveCacheStorage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveCacheStorage
//go:noescape
func CallRemoveCacheStorage(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveCacheStorage
//go:noescape
func TryRemoveCacheStorage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveCookies
//go:noescape
func HasFuncRemoveCookies() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveCookies
//go:noescape
func FuncRemoveCookies(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveCookies
//go:noescape
func CallRemoveCookies(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveCookies
//go:noescape
func TryRemoveCookies(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveDownloads
//go:noescape
func HasFuncRemoveDownloads() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveDownloads
//go:noescape
func FuncRemoveDownloads(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveDownloads
//go:noescape
func CallRemoveDownloads(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveDownloads
//go:noescape
func TryRemoveDownloads(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveFileSystems
//go:noescape
func HasFuncRemoveFileSystems() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveFileSystems
//go:noescape
func FuncRemoveFileSystems(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveFileSystems
//go:noescape
func CallRemoveFileSystems(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveFileSystems
//go:noescape
func TryRemoveFileSystems(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveFormData
//go:noescape
func HasFuncRemoveFormData() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveFormData
//go:noescape
func FuncRemoveFormData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveFormData
//go:noescape
func CallRemoveFormData(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveFormData
//go:noescape
func TryRemoveFormData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveHistory
//go:noescape
func HasFuncRemoveHistory() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveHistory
//go:noescape
func FuncRemoveHistory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveHistory
//go:noescape
func CallRemoveHistory(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveHistory
//go:noescape
func TryRemoveHistory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveIndexedDB
//go:noescape
func HasFuncRemoveIndexedDB() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveIndexedDB
//go:noescape
func FuncRemoveIndexedDB(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveIndexedDB
//go:noescape
func CallRemoveIndexedDB(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveIndexedDB
//go:noescape
func TryRemoveIndexedDB(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveLocalStorage
//go:noescape
func HasFuncRemoveLocalStorage() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveLocalStorage
//go:noescape
func FuncRemoveLocalStorage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveLocalStorage
//go:noescape
func CallRemoveLocalStorage(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveLocalStorage
//go:noescape
func TryRemoveLocalStorage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemovePasswords
//go:noescape
func HasFuncRemovePasswords() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemovePasswords
//go:noescape
func FuncRemovePasswords(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemovePasswords
//go:noescape
func CallRemovePasswords(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemovePasswords
//go:noescape
func TryRemovePasswords(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemovePluginData
//go:noescape
func HasFuncRemovePluginData() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemovePluginData
//go:noescape
func FuncRemovePluginData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemovePluginData
//go:noescape
func CallRemovePluginData(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemovePluginData
//go:noescape
func TryRemovePluginData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveServiceWorkers
//go:noescape
func HasFuncRemoveServiceWorkers() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveServiceWorkers
//go:noescape
func FuncRemoveServiceWorkers(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveServiceWorkers
//go:noescape
func CallRemoveServiceWorkers(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveServiceWorkers
//go:noescape
func TryRemoveServiceWorkers(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_RemoveWebSQL
//go:noescape
func HasFuncRemoveWebSQL() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_RemoveWebSQL
//go:noescape
func FuncRemoveWebSQL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_RemoveWebSQL
//go:noescape
func CallRemoveWebSQL(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_RemoveWebSQL
//go:noescape
func TryRemoveWebSQL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/browsingdata has_Settings
//go:noescape
func HasFuncSettings() js.Ref

//go:wasmimport plat/js/webext/browsingdata func_Settings
//go:noescape
func FuncSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata call_Settings
//go:noescape
func CallSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/browsingdata try_Settings
//go:noescape
func TrySettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

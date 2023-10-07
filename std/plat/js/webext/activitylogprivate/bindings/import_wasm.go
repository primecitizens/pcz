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

//go:wasmimport plat/js/webext/activitylogprivate constof_ExtensionActivityType
//go:noescape
func ConstOfExtensionActivityType(str js.Ref) uint32

//go:wasmimport plat/js/webext/activitylogprivate constof_ExtensionActivityDomVerb
//go:noescape
func ConstOfExtensionActivityDomVerb(str js.Ref) uint32

//go:wasmimport plat/js/webext/activitylogprivate store_ExtensionActivityFieldOther
//go:noescape
func ExtensionActivityFieldOtherJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate load_ExtensionActivityFieldOther
//go:noescape
func ExtensionActivityFieldOtherJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/activitylogprivate store_ExtensionActivity
//go:noescape
func ExtensionActivityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate load_ExtensionActivity
//go:noescape
func ExtensionActivityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/activitylogprivate store_ActivityResultSet
//go:noescape
func ActivityResultSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate load_ActivityResultSet
//go:noescape
func ActivityResultSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/activitylogprivate constof_ExtensionActivityFilter
//go:noescape
func ConstOfExtensionActivityFilter(str js.Ref) uint32

//go:wasmimport plat/js/webext/activitylogprivate store_Filter
//go:noescape
func FilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate load_Filter
//go:noescape
func FilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/activitylogprivate has_DeleteActivities
//go:noescape
func HasFuncDeleteActivities() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_DeleteActivities
//go:noescape
func FuncDeleteActivities(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_DeleteActivities
//go:noescape
func CallDeleteActivities(
	retPtr unsafe.Pointer,
	activityIds js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_DeleteActivities
//go:noescape
func TryDeleteActivities(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	activityIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_DeleteActivitiesByExtension
//go:noescape
func HasFuncDeleteActivitiesByExtension() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_DeleteActivitiesByExtension
//go:noescape
func FuncDeleteActivitiesByExtension(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_DeleteActivitiesByExtension
//go:noescape
func CallDeleteActivitiesByExtension(
	retPtr unsafe.Pointer,
	extensionId js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_DeleteActivitiesByExtension
//go:noescape
func TryDeleteActivitiesByExtension(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_DeleteDatabase
//go:noescape
func HasFuncDeleteDatabase() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_DeleteDatabase
//go:noescape
func FuncDeleteDatabase(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_DeleteDatabase
//go:noescape
func CallDeleteDatabase(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate try_DeleteDatabase
//go:noescape
func TryDeleteDatabase(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_DeleteUrls
//go:noescape
func HasFuncDeleteUrls() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_DeleteUrls
//go:noescape
func FuncDeleteUrls(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_DeleteUrls
//go:noescape
func CallDeleteUrls(
	retPtr unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_DeleteUrls
//go:noescape
func TryDeleteUrls(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_GetExtensionActivities
//go:noescape
func HasFuncGetExtensionActivities() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_GetExtensionActivities
//go:noescape
func FuncGetExtensionActivities(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_GetExtensionActivities
//go:noescape
func CallGetExtensionActivities(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate try_GetExtensionActivities
//go:noescape
func TryGetExtensionActivities(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_OnExtensionActivity
//go:noescape
func HasFuncOnExtensionActivity() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_OnExtensionActivity
//go:noescape
func FuncOnExtensionActivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_OnExtensionActivity
//go:noescape
func CallOnExtensionActivity(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_OnExtensionActivity
//go:noescape
func TryOnExtensionActivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_OffExtensionActivity
//go:noescape
func HasFuncOffExtensionActivity() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_OffExtensionActivity
//go:noescape
func FuncOffExtensionActivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_OffExtensionActivity
//go:noescape
func CallOffExtensionActivity(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_OffExtensionActivity
//go:noescape
func TryOffExtensionActivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate has_HasOnExtensionActivity
//go:noescape
func HasFuncHasOnExtensionActivity() js.Ref

//go:wasmimport plat/js/webext/activitylogprivate func_HasOnExtensionActivity
//go:noescape
func FuncHasOnExtensionActivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/activitylogprivate call_HasOnExtensionActivity
//go:noescape
func CallHasOnExtensionActivity(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/activitylogprivate try_HasOnExtensionActivity
//go:noescape
func TryHasOnExtensionActivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

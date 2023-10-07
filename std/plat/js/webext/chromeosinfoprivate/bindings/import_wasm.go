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

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_AssistantStatus
//go:noescape
func ConstOfAssistantStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_DeviceType
//go:noescape
func ConstOfDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_ManagedDeviceStatus
//go:noescape
func ConstOfManagedDeviceStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_PlayStoreStatus
//go:noescape
func ConstOfPlayStoreStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_SessionType
//go:noescape
func ConstOfSessionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_StylusStatus
//go:noescape
func ConstOfStylusStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate store_GetReturnType
//go:noescape
func GetReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeosinfoprivate load_GetReturnType
//go:noescape
func GetReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeosinfoprivate constof_PropertyName
//go:noescape
func ConstOfPropertyName(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeosinfoprivate has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/chromeosinfoprivate func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeosinfoprivate call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	propertyNames js.Ref)

//go:wasmimport plat/js/webext/chromeosinfoprivate try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	propertyNames js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeosinfoprivate has_IsTabletModeEnabled
//go:noescape
func HasFuncIsTabletModeEnabled() js.Ref

//go:wasmimport plat/js/webext/chromeosinfoprivate func_IsTabletModeEnabled
//go:noescape
func FuncIsTabletModeEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeosinfoprivate call_IsTabletModeEnabled
//go:noescape
func CallIsTabletModeEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeosinfoprivate try_IsTabletModeEnabled
//go:noescape
func TryIsTabletModeEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeosinfoprivate has_Set
//go:noescape
func HasFuncSet() js.Ref

//go:wasmimport plat/js/webext/chromeosinfoprivate func_Set
//go:noescape
func FuncSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeosinfoprivate call_Set
//go:noescape
func CallSet(
	retPtr unsafe.Pointer,
	propertyName uint32,
	propertyValue js.Ref)

//go:wasmimport plat/js/webext/chromeosinfoprivate try_Set
//go:noescape
func TrySet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	propertyName uint32,
	propertyValue js.Ref) (ok js.Ref)

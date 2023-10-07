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

//go:wasmimport plat/js/webext/audio constof_StreamType
//go:noescape
func ConstOfStreamType(str js.Ref) uint32

//go:wasmimport plat/js/webext/audio constof_DeviceType
//go:noescape
func ConstOfDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/audio store_AudioDeviceInfo
//go:noescape
func AudioDeviceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_AudioDeviceInfo
//go:noescape
func AudioDeviceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio store_DeviceFilter
//go:noescape
func DeviceFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_DeviceFilter
//go:noescape
func DeviceFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio store_DeviceIdLists
//go:noescape
func DeviceIdListsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_DeviceIdLists
//go:noescape
func DeviceIdListsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio store_DeviceProperties
//go:noescape
func DevicePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_DeviceProperties
//go:noescape
func DevicePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio store_LevelChangedEvent
//go:noescape
func LevelChangedEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_LevelChangedEvent
//go:noescape
func LevelChangedEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio store_MuteChangedEvent
//go:noescape
func MuteChangedEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/audio load_MuteChangedEvent
//go:noescape
func MuteChangedEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/audio has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/audio func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/audio try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_GetMute
//go:noescape
func HasFuncGetMute() js.Ref

//go:wasmimport plat/js/webext/audio func_GetMute
//go:noescape
func FuncGetMute(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_GetMute
//go:noescape
func CallGetMute(
	retPtr unsafe.Pointer,
	streamType uint32)

//go:wasmimport plat/js/webext/audio try_GetMute
//go:noescape
func TryGetMute(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	streamType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OnDeviceListChanged
//go:noescape
func HasFuncOnDeviceListChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OnDeviceListChanged
//go:noescape
func FuncOnDeviceListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OnDeviceListChanged
//go:noescape
func CallOnDeviceListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OnDeviceListChanged
//go:noescape
func TryOnDeviceListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OffDeviceListChanged
//go:noescape
func HasFuncOffDeviceListChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OffDeviceListChanged
//go:noescape
func FuncOffDeviceListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OffDeviceListChanged
//go:noescape
func CallOffDeviceListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OffDeviceListChanged
//go:noescape
func TryOffDeviceListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_HasOnDeviceListChanged
//go:noescape
func HasFuncHasOnDeviceListChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_HasOnDeviceListChanged
//go:noescape
func FuncHasOnDeviceListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_HasOnDeviceListChanged
//go:noescape
func CallHasOnDeviceListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_HasOnDeviceListChanged
//go:noescape
func TryHasOnDeviceListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OnLevelChanged
//go:noescape
func HasFuncOnLevelChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OnLevelChanged
//go:noescape
func FuncOnLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OnLevelChanged
//go:noescape
func CallOnLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OnLevelChanged
//go:noescape
func TryOnLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OffLevelChanged
//go:noescape
func HasFuncOffLevelChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OffLevelChanged
//go:noescape
func FuncOffLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OffLevelChanged
//go:noescape
func CallOffLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OffLevelChanged
//go:noescape
func TryOffLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_HasOnLevelChanged
//go:noescape
func HasFuncHasOnLevelChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_HasOnLevelChanged
//go:noescape
func FuncHasOnLevelChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_HasOnLevelChanged
//go:noescape
func CallHasOnLevelChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_HasOnLevelChanged
//go:noescape
func TryHasOnLevelChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OnMuteChanged
//go:noescape
func HasFuncOnMuteChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OnMuteChanged
//go:noescape
func FuncOnMuteChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OnMuteChanged
//go:noescape
func CallOnMuteChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OnMuteChanged
//go:noescape
func TryOnMuteChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_OffMuteChanged
//go:noescape
func HasFuncOffMuteChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_OffMuteChanged
//go:noescape
func FuncOffMuteChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_OffMuteChanged
//go:noescape
func CallOffMuteChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_OffMuteChanged
//go:noescape
func TryOffMuteChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_HasOnMuteChanged
//go:noescape
func HasFuncHasOnMuteChanged() js.Ref

//go:wasmimport plat/js/webext/audio func_HasOnMuteChanged
//go:noescape
func FuncHasOnMuteChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_HasOnMuteChanged
//go:noescape
func CallHasOnMuteChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/audio try_HasOnMuteChanged
//go:noescape
func TryHasOnMuteChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_SetActiveDevices
//go:noescape
func HasFuncSetActiveDevices() js.Ref

//go:wasmimport plat/js/webext/audio func_SetActiveDevices
//go:noescape
func FuncSetActiveDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_SetActiveDevices
//go:noescape
func CallSetActiveDevices(
	retPtr unsafe.Pointer,
	ids unsafe.Pointer)

//go:wasmimport plat/js/webext/audio try_SetActiveDevices
//go:noescape
func TrySetActiveDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ids unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_SetMute
//go:noescape
func HasFuncSetMute() js.Ref

//go:wasmimport plat/js/webext/audio func_SetMute
//go:noescape
func FuncSetMute(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_SetMute
//go:noescape
func CallSetMute(
	retPtr unsafe.Pointer,
	streamType uint32,
	isMuted js.Ref)

//go:wasmimport plat/js/webext/audio try_SetMute
//go:noescape
func TrySetMute(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	streamType uint32,
	isMuted js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/audio has_SetProperties
//go:noescape
func HasFuncSetProperties() js.Ref

//go:wasmimport plat/js/webext/audio func_SetProperties
//go:noescape
func FuncSetProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/audio call_SetProperties
//go:noescape
func CallSetProperties(
	retPtr unsafe.Pointer,
	id js.Ref,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/audio try_SetProperties
//go:noescape
func TrySetProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	properties unsafe.Pointer) (ok js.Ref)

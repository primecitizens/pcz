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

//go:wasmimport plat/js/webext/enterprise/deviceattributes has_GetDeviceAnnotatedLocation
//go:noescape
func HasFuncGetDeviceAnnotatedLocation() js.Ref

//go:wasmimport plat/js/webext/enterprise/deviceattributes func_GetDeviceAnnotatedLocation
//go:noescape
func FuncGetDeviceAnnotatedLocation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes call_GetDeviceAnnotatedLocation
//go:noescape
func CallGetDeviceAnnotatedLocation(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes try_GetDeviceAnnotatedLocation
//go:noescape
func TryGetDeviceAnnotatedLocation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/deviceattributes has_GetDeviceAssetId
//go:noescape
func HasFuncGetDeviceAssetId() js.Ref

//go:wasmimport plat/js/webext/enterprise/deviceattributes func_GetDeviceAssetId
//go:noescape
func FuncGetDeviceAssetId(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes call_GetDeviceAssetId
//go:noescape
func CallGetDeviceAssetId(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes try_GetDeviceAssetId
//go:noescape
func TryGetDeviceAssetId(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/deviceattributes has_GetDeviceHostname
//go:noescape
func HasFuncGetDeviceHostname() js.Ref

//go:wasmimport plat/js/webext/enterprise/deviceattributes func_GetDeviceHostname
//go:noescape
func FuncGetDeviceHostname(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes call_GetDeviceHostname
//go:noescape
func CallGetDeviceHostname(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes try_GetDeviceHostname
//go:noescape
func TryGetDeviceHostname(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/deviceattributes has_GetDeviceSerialNumber
//go:noescape
func HasFuncGetDeviceSerialNumber() js.Ref

//go:wasmimport plat/js/webext/enterprise/deviceattributes func_GetDeviceSerialNumber
//go:noescape
func FuncGetDeviceSerialNumber(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes call_GetDeviceSerialNumber
//go:noescape
func CallGetDeviceSerialNumber(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes try_GetDeviceSerialNumber
//go:noescape
func TryGetDeviceSerialNumber(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/deviceattributes has_GetDirectoryDeviceId
//go:noescape
func HasFuncGetDirectoryDeviceId() js.Ref

//go:wasmimport plat/js/webext/enterprise/deviceattributes func_GetDirectoryDeviceId
//go:noescape
func FuncGetDirectoryDeviceId(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes call_GetDirectoryDeviceId
//go:noescape
func CallGetDirectoryDeviceId(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/deviceattributes try_GetDirectoryDeviceId
//go:noescape
func TryGetDirectoryDeviceId(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

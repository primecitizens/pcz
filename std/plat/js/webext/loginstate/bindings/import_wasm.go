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

//go:wasmimport plat/js/webext/loginstate constof_ProfileType
//go:noescape
func ConstOfProfileType(str js.Ref) uint32

//go:wasmimport plat/js/webext/loginstate constof_SessionState
//go:noescape
func ConstOfSessionState(str js.Ref) uint32

//go:wasmimport plat/js/webext/loginstate has_GetProfileType
//go:noescape
func HasFuncGetProfileType() js.Ref

//go:wasmimport plat/js/webext/loginstate func_GetProfileType
//go:noescape
func FuncGetProfileType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate call_GetProfileType
//go:noescape
func CallGetProfileType(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate try_GetProfileType
//go:noescape
func TryGetProfileType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/loginstate has_GetSessionState
//go:noescape
func HasFuncGetSessionState() js.Ref

//go:wasmimport plat/js/webext/loginstate func_GetSessionState
//go:noescape
func FuncGetSessionState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate call_GetSessionState
//go:noescape
func CallGetSessionState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate try_GetSessionState
//go:noescape
func TryGetSessionState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/loginstate has_OnSessionStateChanged
//go:noescape
func HasFuncOnSessionStateChanged() js.Ref

//go:wasmimport plat/js/webext/loginstate func_OnSessionStateChanged
//go:noescape
func FuncOnSessionStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate call_OnSessionStateChanged
//go:noescape
func CallOnSessionStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/loginstate try_OnSessionStateChanged
//go:noescape
func TryOnSessionStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/loginstate has_OffSessionStateChanged
//go:noescape
func HasFuncOffSessionStateChanged() js.Ref

//go:wasmimport plat/js/webext/loginstate func_OffSessionStateChanged
//go:noescape
func FuncOffSessionStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate call_OffSessionStateChanged
//go:noescape
func CallOffSessionStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/loginstate try_OffSessionStateChanged
//go:noescape
func TryOffSessionStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/loginstate has_HasOnSessionStateChanged
//go:noescape
func HasFuncHasOnSessionStateChanged() js.Ref

//go:wasmimport plat/js/webext/loginstate func_HasOnSessionStateChanged
//go:noescape
func FuncHasOnSessionStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginstate call_HasOnSessionStateChanged
//go:noescape
func CallHasOnSessionStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/loginstate try_HasOnSessionStateChanged
//go:noescape
func TryHasOnSessionStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

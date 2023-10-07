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

//go:wasmimport plat/js/webext/webrtcaudioprivate store_SinkInfo
//go:noescape
func SinkInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate load_SinkInfo
//go:noescape
func SinkInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate has_GetAssociatedSink
//go:noescape
func HasFuncGetAssociatedSink() js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate func_GetAssociatedSink
//go:noescape
func FuncGetAssociatedSink(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate call_GetAssociatedSink
//go:noescape
func CallGetAssociatedSink(
	retPtr unsafe.Pointer,
	securityOrigin js.Ref,
	sourceIdInOrigin js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate try_GetAssociatedSink
//go:noescape
func TryGetAssociatedSink(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	securityOrigin js.Ref,
	sourceIdInOrigin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate has_GetSinks
//go:noescape
func HasFuncGetSinks() js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate func_GetSinks
//go:noescape
func FuncGetSinks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate call_GetSinks
//go:noescape
func CallGetSinks(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate try_GetSinks
//go:noescape
func TryGetSinks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate has_OnSinksChanged
//go:noescape
func HasFuncOnSinksChanged() js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate func_OnSinksChanged
//go:noescape
func FuncOnSinksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate call_OnSinksChanged
//go:noescape
func CallOnSinksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate try_OnSinksChanged
//go:noescape
func TryOnSinksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate has_OffSinksChanged
//go:noescape
func HasFuncOffSinksChanged() js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate func_OffSinksChanged
//go:noescape
func FuncOffSinksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate call_OffSinksChanged
//go:noescape
func CallOffSinksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate try_OffSinksChanged
//go:noescape
func TryOffSinksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate has_HasOnSinksChanged
//go:noescape
func HasFuncHasOnSinksChanged() js.Ref

//go:wasmimport plat/js/webext/webrtcaudioprivate func_HasOnSinksChanged
//go:noescape
func FuncHasOnSinksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcaudioprivate call_HasOnSinksChanged
//go:noescape
func CallHasOnSinksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcaudioprivate try_HasOnSinksChanged
//go:noescape
func TryHasOnSinksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

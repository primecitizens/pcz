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

//go:wasmimport plat/js/webext/mediaplayerprivate has_OnNextTrack
//go:noescape
func HasFuncOnNextTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OnNextTrack
//go:noescape
func FuncOnNextTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OnNextTrack
//go:noescape
func CallOnNextTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OnNextTrack
//go:noescape
func TryOnNextTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_OffNextTrack
//go:noescape
func HasFuncOffNextTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OffNextTrack
//go:noescape
func FuncOffNextTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OffNextTrack
//go:noescape
func CallOffNextTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OffNextTrack
//go:noescape
func TryOffNextTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_HasOnNextTrack
//go:noescape
func HasFuncHasOnNextTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_HasOnNextTrack
//go:noescape
func FuncHasOnNextTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_HasOnNextTrack
//go:noescape
func CallHasOnNextTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_HasOnNextTrack
//go:noescape
func TryHasOnNextTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_OnPrevTrack
//go:noescape
func HasFuncOnPrevTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OnPrevTrack
//go:noescape
func FuncOnPrevTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OnPrevTrack
//go:noescape
func CallOnPrevTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OnPrevTrack
//go:noescape
func TryOnPrevTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_OffPrevTrack
//go:noescape
func HasFuncOffPrevTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OffPrevTrack
//go:noescape
func FuncOffPrevTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OffPrevTrack
//go:noescape
func CallOffPrevTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OffPrevTrack
//go:noescape
func TryOffPrevTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_HasOnPrevTrack
//go:noescape
func HasFuncHasOnPrevTrack() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_HasOnPrevTrack
//go:noescape
func FuncHasOnPrevTrack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_HasOnPrevTrack
//go:noescape
func CallHasOnPrevTrack(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_HasOnPrevTrack
//go:noescape
func TryHasOnPrevTrack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_OnTogglePlayState
//go:noescape
func HasFuncOnTogglePlayState() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OnTogglePlayState
//go:noescape
func FuncOnTogglePlayState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OnTogglePlayState
//go:noescape
func CallOnTogglePlayState(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OnTogglePlayState
//go:noescape
func TryOnTogglePlayState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_OffTogglePlayState
//go:noescape
func HasFuncOffTogglePlayState() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_OffTogglePlayState
//go:noescape
func FuncOffTogglePlayState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_OffTogglePlayState
//go:noescape
func CallOffTogglePlayState(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_OffTogglePlayState
//go:noescape
func TryOffTogglePlayState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate has_HasOnTogglePlayState
//go:noescape
func HasFuncHasOnTogglePlayState() js.Ref

//go:wasmimport plat/js/webext/mediaplayerprivate func_HasOnTogglePlayState
//go:noescape
func FuncHasOnTogglePlayState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mediaplayerprivate call_HasOnTogglePlayState
//go:noescape
func CallHasOnTogglePlayState(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mediaplayerprivate try_HasOnTogglePlayState
//go:noescape
func TryHasOnTogglePlayState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/diagnostics store_SendPacketResult
//go:noescape
func SendPacketResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/diagnostics load_SendPacketResult
//go:noescape
func SendPacketResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/diagnostics store_SendPacketOptions
//go:noescape
func SendPacketOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/diagnostics load_SendPacketOptions
//go:noescape
func SendPacketOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/diagnostics has_SendPacket
//go:noescape
func HasFuncSendPacket() js.Ref

//go:wasmimport plat/js/webext/diagnostics func_SendPacket
//go:noescape
func FuncSendPacket(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/diagnostics call_SendPacket
//go:noescape
func CallSendPacket(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/diagnostics try_SendPacket
//go:noescape
func TrySendPacket(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/webext/appviewguestinternal has_AttachFrame
//go:noescape
func HasFuncAttachFrame() js.Ref

//go:wasmimport plat/js/webext/appviewguestinternal func_AttachFrame
//go:noescape
func FuncAttachFrame(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewguestinternal call_AttachFrame
//go:noescape
func CallAttachFrame(
	retPtr unsafe.Pointer,
	url js.Ref,
	guestInstanceId float64)

//go:wasmimport plat/js/webext/appviewguestinternal try_AttachFrame
//go:noescape
func TryAttachFrame(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	guestInstanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/appviewguestinternal has_DenyRequest
//go:noescape
func HasFuncDenyRequest() js.Ref

//go:wasmimport plat/js/webext/appviewguestinternal func_DenyRequest
//go:noescape
func FuncDenyRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewguestinternal call_DenyRequest
//go:noescape
func CallDenyRequest(
	retPtr unsafe.Pointer,
	guestInstanceId float64)

//go:wasmimport plat/js/webext/appviewguestinternal try_DenyRequest
//go:noescape
func TryDenyRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	guestInstanceId float64) (ok js.Ref)

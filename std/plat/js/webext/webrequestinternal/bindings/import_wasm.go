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

//go:wasmimport plat/js/webext/webrequestinternal constof_AddEventListenerOptions
//go:noescape
func ConstOfAddEventListenerOptions(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrequestinternal has_AddEventListener
//go:noescape
func HasFuncAddEventListener() js.Ref

//go:wasmimport plat/js/webext/webrequestinternal func_AddEventListener
//go:noescape
func FuncAddEventListener(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequestinternal call_AddEventListener
//go:noescape
func CallAddEventListener(
	retPtr unsafe.Pointer,
	callback js.Ref,
	filter unsafe.Pointer,
	extraInfoSpec js.Ref,
	eventName js.Ref,
	subEventName js.Ref,
	webViewInstanceId float64)

//go:wasmimport plat/js/webext/webrequestinternal try_AddEventListener
//go:noescape
func TryAddEventListener(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref,
	filter unsafe.Pointer,
	extraInfoSpec js.Ref,
	eventName js.Ref,
	subEventName js.Ref,
	webViewInstanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/webrequestinternal has_EventHandled
//go:noescape
func HasFuncEventHandled() js.Ref

//go:wasmimport plat/js/webext/webrequestinternal func_EventHandled
//go:noescape
func FuncEventHandled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequestinternal call_EventHandled
//go:noescape
func CallEventHandled(
	retPtr unsafe.Pointer,
	eventName js.Ref,
	subEventName js.Ref,
	requestId js.Ref,
	webViewInstanceId float64,
	response unsafe.Pointer)

//go:wasmimport plat/js/webext/webrequestinternal try_EventHandled
//go:noescape
func TryEventHandled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	subEventName js.Ref,
	requestId js.Ref,
	webViewInstanceId float64,
	response unsafe.Pointer) (ok js.Ref)

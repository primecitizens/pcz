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

//go:wasmimport plat/js/webext/accessibilityserviceprivate has_SpeakSelectedText
//go:noescape
func HasFuncSpeakSelectedText() js.Ref

//go:wasmimport plat/js/webext/accessibilityserviceprivate func_SpeakSelectedText
//go:noescape
func FuncSpeakSelectedText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityserviceprivate call_SpeakSelectedText
//go:noescape
func CallSpeakSelectedText(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityserviceprivate try_SpeakSelectedText
//go:noescape
func TrySpeakSelectedText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/webext/downloadsinternal has_DetermineFilename
//go:noescape
func HasFuncDetermineFilename() js.Ref

//go:wasmimport plat/js/webext/downloadsinternal func_DetermineFilename
//go:noescape
func FuncDetermineFilename(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloadsinternal call_DetermineFilename
//go:noescape
func CallDetermineFilename(
	retPtr unsafe.Pointer,
	downloadId int32,
	filename js.Ref,
	conflict_action js.Ref)

//go:wasmimport plat/js/webext/downloadsinternal try_DetermineFilename
//go:noescape
func TryDetermineFilename(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32,
	filename js.Ref,
	conflict_action js.Ref) (ok js.Ref)

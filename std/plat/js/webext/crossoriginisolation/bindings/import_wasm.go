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

//go:wasmimport plat/js/webext/crossoriginisolation store_ResponseHeader
//go:noescape
func ResponseHeaderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/crossoriginisolation load_ResponseHeader
//go:noescape
func ResponseHeaderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/crossoriginisolation store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/crossoriginisolation load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

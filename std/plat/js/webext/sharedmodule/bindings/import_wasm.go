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

//go:wasmimport plat/js/webext/sharedmodule store_Export
//go:noescape
func ExportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sharedmodule load_Export
//go:noescape
func ExportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sharedmodule store_Import
//go:noescape
func ImportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sharedmodule load_Import
//go:noescape
func ImportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sharedmodule store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sharedmodule load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

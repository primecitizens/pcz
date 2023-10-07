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

//go:wasmimport plat/js/webext/contentscripts constof_RunAt
//go:noescape
func ConstOfRunAt(str js.Ref) uint32

//go:wasmimport plat/js/webext/contentscripts store_ContentScript
//go:noescape
func ContentScriptJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentscripts load_ContentScript
//go:noescape
func ContentScriptJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/contentscripts store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/contentscripts load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

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

//go:wasmimport plat/js/webext/chrome_url_overrides store_UrlOverrideInfo
//go:noescape
func UrlOverrideInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chrome_url_overrides load_UrlOverrideInfo
//go:noescape
func UrlOverrideInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chrome_url_overrides store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chrome_url_overrides load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

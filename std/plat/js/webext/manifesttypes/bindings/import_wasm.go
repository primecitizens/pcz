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

//go:wasmimport plat/js/webext/manifesttypes store_ChromeSettingsOverridesFieldSearchProvider
//go:noescape
func ChromeSettingsOverridesFieldSearchProviderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/manifesttypes load_ChromeSettingsOverridesFieldSearchProvider
//go:noescape
func ChromeSettingsOverridesFieldSearchProviderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/manifesttypes store_ChromeSettingsOverrides
//go:noescape
func ChromeSettingsOverridesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/manifesttypes load_ChromeSettingsOverrides
//go:noescape
func ChromeSettingsOverridesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/manifesttypes constof_FileSystemProviderSource
//go:noescape
func ConstOfFileSystemProviderSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/manifesttypes store_FileSystemProviderCapabilities
//go:noescape
func FileSystemProviderCapabilitiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/manifesttypes load_FileSystemProviderCapabilities
//go:noescape
func FileSystemProviderCapabilitiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

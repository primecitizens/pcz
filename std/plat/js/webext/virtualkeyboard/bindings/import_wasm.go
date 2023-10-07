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

//go:wasmimport plat/js/webext/virtualkeyboard store_FeatureRestrictions
//go:noescape
func FeatureRestrictionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboard load_FeatureRestrictions
//go:noescape
func FeatureRestrictionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboard has_RestrictFeatures
//go:noescape
func HasFuncRestrictFeatures() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboard func_RestrictFeatures
//go:noescape
func FuncRestrictFeatures(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboard call_RestrictFeatures
//go:noescape
func CallRestrictFeatures(
	retPtr unsafe.Pointer,
	restrictions unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboard try_RestrictFeatures
//go:noescape
func TryRestrictFeatures(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	restrictions unsafe.Pointer) (ok js.Ref)

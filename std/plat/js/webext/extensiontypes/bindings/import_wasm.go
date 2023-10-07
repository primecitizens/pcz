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

//go:wasmimport plat/js/webext/extensiontypes constof_CSSOrigin
//go:noescape
func ConstOfCSSOrigin(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes constof_ExecutionWorld
//go:noescape
func ConstOfExecutionWorld(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes constof_RunAt
//go:noescape
func ConstOfRunAt(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes store_DeleteInjectionDetails
//go:noescape
func DeleteInjectionDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensiontypes load_DeleteInjectionDetails
//go:noescape
func DeleteInjectionDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensiontypes constof_DocumentLifecycle
//go:noescape
func ConstOfDocumentLifecycle(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes constof_FrameType
//go:noescape
func ConstOfFrameType(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes constof_ImageFormat
//go:noescape
func ConstOfImageFormat(str js.Ref) uint32

//go:wasmimport plat/js/webext/extensiontypes store_ImageDetails
//go:noescape
func ImageDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensiontypes load_ImageDetails
//go:noescape
func ImageDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensiontypes store_InjectDetails
//go:noescape
func InjectDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensiontypes load_InjectDetails
//go:noescape
func InjectDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

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

//go:wasmimport plat/js/webext/metricsprivateindividualapis has_RecordEnumerationValue
//go:noescape
func HasFuncRecordEnumerationValue() js.Ref

//go:wasmimport plat/js/webext/metricsprivateindividualapis func_RecordEnumerationValue
//go:noescape
func FuncRecordEnumerationValue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivateindividualapis call_RecordEnumerationValue
//go:noescape
func CallRecordEnumerationValue(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64,
	enumSize float64)

//go:wasmimport plat/js/webext/metricsprivateindividualapis try_RecordEnumerationValue
//go:noescape
func TryRecordEnumerationValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64,
	enumSize float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivateindividualapis has_RecordMediumTime
//go:noescape
func HasFuncRecordMediumTime() js.Ref

//go:wasmimport plat/js/webext/metricsprivateindividualapis func_RecordMediumTime
//go:noescape
func FuncRecordMediumTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivateindividualapis call_RecordMediumTime
//go:noescape
func CallRecordMediumTime(
	retPtr unsafe.Pointer,
	metricName js.Ref,
	value float64)

//go:wasmimport plat/js/webext/metricsprivateindividualapis try_RecordMediumTime
//go:noescape
func TryRecordMediumTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	metricName js.Ref,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/webext/metricsprivateindividualapis has_RecordUserAction
//go:noescape
func HasFuncRecordUserAction() js.Ref

//go:wasmimport plat/js/webext/metricsprivateindividualapis func_RecordUserAction
//go:noescape
func FuncRecordUserAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/metricsprivateindividualapis call_RecordUserAction
//go:noescape
func CallRecordUserAction(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/metricsprivateindividualapis try_RecordUserAction
//go:noescape
func TryRecordUserAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

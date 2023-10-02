// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package async

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js/bindings"
)

//go:wasmimport ffi/js/async promise
//go:noescape
func Promise(fn bindings.Ref) bindings.Ref

//go:wasmimport ffi/js/async resolved
//go:noescape
func Resolved(data bindings.Ref) bindings.Ref

//go:wasmimport ffi/js/async rejected
//go:noescape
func Rejected(reason bindings.Ref) bindings.Ref

//go:wasmimport ffi/js/async bg
//go:noescape
func Bg(promise bindings.Ref) bindings.Ref

//go:wasmimport ffi/js/async shouldwait
//go:noescape
func ShouldWait(
	take bindings.Ref,
	promise bindings.Ref,
	pFulfilled unsafe.Pointer,
	pRefData unsafe.Pointer,
) (shouldWait bindings.Ref)

//go:wasmimport ffi/js/async takewaited
//go:noescape
func TakeWaited(ref bindings.Ref, pRefData unsafe.Pointer) (fulfilled bindings.Ref)

//go:wasmimport ffi/js/async allok
//go:noescape
func AllFulfilled(free bindings.Ref, pRefs unsafe.Pointer, n uint32) bindings.Ref

//go:wasmimport ffi/js/async all
//go:noescape
func AllFinished(free bindings.Ref, pRefs unsafe.Pointer, n uint32) bindings.Ref

//go:wasmimport ffi/js/async first
//go:noescape
func FirstFinished(free bindings.Ref, pRefs unsafe.Pointer, n uint32) bindings.Ref

//go:wasmimport ffi/js/async then
//go:noescape
func Then(onfulfilled, onrejected, finally bindings.Ref)

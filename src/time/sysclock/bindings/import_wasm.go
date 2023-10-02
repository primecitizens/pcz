// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"
)

//go:wasmimport time/sysclock timeOriginMS
func TimeOriginMS(pI64 unsafe.Pointer)

//go:wasmimport time/sysclock walltime
func Walltime(pI64 unsafe.Pointer)

//go:wasmimport time/sysclock millitime
func Millitime(pI64 unsafe.Pointer)

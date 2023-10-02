// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

//go:wasmimport time/sysclock timeOriginMS
func TimeOriginMS() float64

//go:wasmimport time/sysclock walltime
func Walltime() float64

//go:wasmimport time/sysclock millitime
func Millitime() float64

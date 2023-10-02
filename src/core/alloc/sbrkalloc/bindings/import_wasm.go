// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

//go:wasmimport core/alloc/sbrkalloc resetMemoryDataView
//go:noescape
func ResetMemoryDataView()

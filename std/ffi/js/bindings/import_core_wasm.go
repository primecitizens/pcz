// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

//go:wasmimport ffi/js.core exit
//go:noescape
func Exit(code int32)

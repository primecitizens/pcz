// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js && wasm

package js

//go:noescape
func callDispatcher(
	recv, targetPC uintptr, ctx *CallbackContext, dispFn uintptr,
)

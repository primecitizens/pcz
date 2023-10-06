// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package ffi

// CallbackDispatcher defines the generic interface to dispatch ffi callbacks.
//
// The CallbackContext is specific to each foreign language.
type CallbackDispatcher[CallbackContext any] interface {
	// DispatchCallback dispatches a call from foreign function.
	//
	// targetPC is the pc provided when registering the target function.
	// ctx is the foreign function specific callback context.
	DispatchCallback(targetPC uintptr, ctx *CallbackContext)
}

// DispatchFunc is the function signature of
// CallbackDispatcher.DispatchCallback
type DispatchFunc[T any, CallbackContext any] func(
	recv T, targetPC uintptr, ctx *CallbackContext,
)

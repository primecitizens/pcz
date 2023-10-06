// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos || !(js && wasm)

package js

import (
	"github.com/primecitizens/pcz/std/core/assert"
)

func callDispatcher(
	recv, targetPC uintptr, ctx *CallbackContext, fn uintptr,
) {
	assert.Throw("unsupported", "operation")
}

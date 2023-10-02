// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package assert

import (
	"github.com/primecitizens/std/core/debug"
)

func Eq[T comparable](a, b T) bool {
	if a == b {
		return true
	}

	println("assert.Eq", "failed:", a, "!=", b)
	debug.Abort()
	return false
}

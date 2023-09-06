// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package assert

import (
	"github.com/primecitizens/std/core/debug"
	"github.com/primecitizens/std/core/mark"
)

func Panic[T any](v ...T) {
	Throw("panic")
}

func Unreachable() {
	Throw("unreachable")
}

func TODO(s ...string) {
	Throw("TODO")
}

// Throw triggers a fatal error that dumps a stack trace and exits.
//
//go:nosplit
func Throw(s ...string) {
	s = mark.NoEscapeSlice(s)

	switch len(s) {
	case 0:
		println("fatal", "error")
	case 1:
		println("fatal", "error:", s[0])
	default:
		print("fatal", " ", "error:")
		for _, s := range s {
			print(" ", s)
		}
		println()
	}

	// TODO: print stacktrace

	debug.Abort()
}

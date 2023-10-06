// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js && wasm

package assert

import (
	"github.com/primecitizens/pcz/std/core/assert/bindings"
	jsbindings "github.com/primecitizens/pcz/std/ffi/js/bindings"
)

// Throw triggers a fatal error that dumps a stack trace and exits.
//
//go:nosplit
func Throw(s ...string) {
	switch len(s) {
	case 0:
		bindings.Append("fatal", " ", "error")
	case 1:
		bindings.Append("fatal", " ", "error", ":", " ", s[0])
	default:
		bindings.Append("fatal", " ", "error", ":")
		for _, s := range s {
			bindings.Append(" ", s)
		}
	}

	jsbindings.Exit(1)
	bindings.Throw()
}

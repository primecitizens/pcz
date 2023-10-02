// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package thread

import (
	stdgo "github.com/primecitizens/std/builtin/go"
	"github.com/primecitizens/std/core/assert"
)

// GetTLSBaseAddress is not supported for non-pcz toolchain
func GetTLSBaseAddress() uintptr {
	assert.Throw("not", "supported")
	return 0
}

// G is not supported for non-pcz toolchain
func G() *stdgo.GHead {
	assert.Throw("not", "supported")
	return nil
}

// SetG is not supported for non-pcz toolchain
func SetG(g *stdgo.GHead) {
	assert.Throw("not", "supported")
}

// Topframe is not supported for non-pcz toolchain
func Topframe(arg1, sp, pc uintptr) {
	assert.Throw("not", "supported")
}

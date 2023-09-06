// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alloc

import (
	"github.com/primecitizens/std/core/assert"
)

// Use golang allocation primitives when compiled with official go runtime

const (
	useGoNativeAlloc = true
)

func malloc(alignment, sz uintptr) (p uintptr, err int) {
	assert.Unreachable()
	return
}

func free(p, sz uintptr) (err int) {
	assert.Unreachable()
	return
}

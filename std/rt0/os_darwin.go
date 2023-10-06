// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && darwin

package rt0

import (
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/alloc/libcalloc"
	"github.com/primecitizens/pcz/std/plat/libc"
)

var _alloc libcalloc.T

func malloc() alloc.M { return &_alloc }
func palloc() alloc.P { return &_alloc }

func exit0() {
	libc.Exit(0)
}

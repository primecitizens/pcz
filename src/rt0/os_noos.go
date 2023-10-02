// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos

package rt0

import (
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/alloc"
)

var _alloc alloc.DenyNoneZero

func malloc() alloc.M { return &_alloc }
func palloc() alloc.P { return &_alloc }

//go:linkname exit0 main.exit0
func exit0()

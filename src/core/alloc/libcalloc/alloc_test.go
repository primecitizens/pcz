// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package libcalloc

import (
	"github.com/primecitizens/std/core/alloc"
)

var (
	_ alloc.M = T{}
	_ alloc.P = T{}
)

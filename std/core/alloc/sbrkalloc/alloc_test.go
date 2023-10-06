// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sbrkalloc

import "github.com/primecitizens/pcz/std/core/alloc"

var (
	_ alloc.M = (*T)(nil)
	_ alloc.P = (*T)(nil)
)

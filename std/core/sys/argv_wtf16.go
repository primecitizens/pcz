// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && windows

package sys

import (
	_ "unsafe" // for go:linkname
)

var (
	//go:linkname _args rt0.args
	_args []*uint16
	args  []string
)

func ntharg(i int) string {
	return args[i]
}

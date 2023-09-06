// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && windows

package sys

import (
	_ "unsafe" // for go:linkname

	stdstring "github.com/primecitizens/std/builtin/string"
)

var (
	//go:linkname args rt0.args
	args []*uint16
)

func ntharg(i int) string {
	return stdstring.FromByteArray((*byte)(unsafe.Pointer(args[i])))
}

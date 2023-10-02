// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (darwin || linux || freebsd || netbsd || openbsd || solaris || dragonfly)

package sys

import (
	_ "unsafe" // for go:linkname

	stdstring "github.com/primecitizens/std/builtin/string"
)

var (
	//go:linkname _args rt0.args
	_args []*byte
)

func ntharg(i int) string {
	return stdstring.FromByteArray(_args[i])
}

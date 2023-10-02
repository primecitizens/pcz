// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (darwin || linux || freebsd || netbsd || openbsd || solaris || dragonfly)

package sys

import (
	stdstring "github.com/primecitizens/std/builtin/string"
)

var envs []*byte

func nthenv(i int) string {
	return stdstring.FromByteArray(envs[i])
}

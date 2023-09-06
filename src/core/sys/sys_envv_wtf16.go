// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && windows

package sys

import (
	"unsafe"

	stdstring "github.com/primecitizens/std/builtin/string"
)

var envs []*uint16

func nthenv(i int) string {
	return stdstring.FromByteArray((*byte)(unsafe.Pointer(envs[i])))
}

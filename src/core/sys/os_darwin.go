// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && darwin

package sys

import (
	"unsafe"

	stdptr "github.com/primecitizens/std/builtin/ptr"
	stdstring "github.com/primecitizens/std/builtin/string"
	"github.com/primecitizens/std/core/arch"
)

var (
	//go:linkname executablePath
	executablePath string
)

func ExecutablePath() string { return executablePath }

func init() {
	base := unsafe.SliceData(args)
	if base == nil { // no args
		return
	}

	// base is the envv
	base = stdptr.Add(base, arch.PtrSize*uintptr(len(args)+1 /* NULL sep */))
	n := 0
	for p := base; *p != nil; p = stdptr.Add(p, arch.PtrSize) {
		n++
	}
	envs = unsafe.Slice(base, n)
	// envs[n] = nil
	// envs[n+1] = "executable_path="
	executablePath = stdstring.FromByteArray(*stdptr.Add(base, arch.PtrSize*uintptr(n+1)))

	// strip "executable_path=" prefix if available, it's added after OS X 10.11.
	const prefix = "executable_path="
	if len(executablePath) > len(prefix) && executablePath[:len(prefix)] == prefix {
		executablePath = executablePath[len(prefix):]
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (linux || freebsd || netbsd || solaris || dragonfly)

package sys

var (
	auxvs []uintptr
)

func init() {
	// base is the envv
	base := stdptr.Add(unsafe.SliceData(args), arch.PtrSize*uintptr(len(args)+1 /* NULL sep */))
	n := 0
	for p := base; *p != nil; p = stdptr.Add(p, arch.PtrSize) {
		n++
	}
	envs = unsafe.Slice(base, n)

	// auxvv
	base = stdptr.Add(base, arch.PtrSize*uintptr(n+1 /* NULL sep */))
	n = 0
	for p := base; *p != nil; p = stdptr.Add(p, arch.PtrSize) {
		n++
	}
	auxvs = unsafe.Slice((*uintptr)(unsafe.Pointer(base)), n)
}

type Auxv uintptr

func nthauxv(i int) Auxv {
	return auxvs[i]
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/os"
)

const (
	// GOOS is the running program's operating system target:
	// one of darwin, freebsd, linux, and so on.
	// To view possible combinations of GOOS and GOARCH, run "go tool dist list".
	GOOS string = os.GOOS

	// GOARCH is the running program's architecture target:
	// one of 386, amd64, arm, s390x, and so on.
	GOARCH string = arch.GOARCH
)

// Version returns the Go tree's version string.
// It is either the commit hash and date at the time of the build or,
// when possible, a release tag like "go1.3".
//
// If any GOEXPERIMENTs are set to non-default values, it will include
// "X:<GOEXPERIMENT>".
func Version() string {
	return buildVersion
}

// GOROOT returns the root used during the Go build.
func GOROOT() string {
	return defaultGOROOT
}

// GOARM returns the GOARM environ used to build this application.
func GOARM() uint8 {
	return goarm
}

// GOMAXPROCS returns 1.
func GOMAXPROCS(int) int {
	return 1
}

// Mark KeepAlive as noinline so that it is easily detectable as an intrinsic.
//
//go:noinline

// KeepAlive marks its argument as currently reachable.
// This ensures that the object is not freed, and its finalizer is not run,
// before the point in the program where KeepAlive is called.
//
// A very simplified example showing where KeepAlive is required:
//
//	type File struct { d int }
//	d, err := syscall.Open("/file/path", syscall.O_RDONLY, 0)
//	// ... do something if err != nil ...
//	p := &File{d}
//	runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
//	var buf [10]byte
//	n, err := syscall.Read(p.d, buf[:])
//	// Ensure p is not finalized until Read returns.
//	runtime.KeepAlive(p)
//	// No more uses of p after this point.
//
// Without the KeepAlive call, the finalizer could run at the start of
// syscall.Read, closing the file descriptor before syscall.Read makes
// the actual system call.
//
// Note: KeepAlive should only be used to prevent finalizers from
// running prematurely. In particular, when used with unsafe.Pointer,
// the rules for valid uses of unsafe.Pointer still apply.
func KeepAlive(x any) {
	// Introduce a use of x that the compiler can't eliminate.
	// This makes sure x is alive on entry. We need x to be alive
	// on entry for "defer runtime.KeepAlive(x)"; see issue 21402.
	if cgoAlwaysFalse {
		println(x)
	}
}

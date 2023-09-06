// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Package thread provides low-level access to machine threads
package thread

import (
	_ "unsafe" // for go:linkname

	stdgo "github.com/primecitizens/std/builtin/go"
)

// GetTLSBaseAddress returns the start address of thread local storage
func GetTLSBaseAddress() uintptr

//go:linkname GetG runtime._getg
func GetG() *stdgo.GHead

//go:noescape
func SetG(g *stdgo.GHead)

// Topframe jumps to the funcpc as top-frame.
//
//go:noescape
func Topframe(arg1, sp, pc uintptr)

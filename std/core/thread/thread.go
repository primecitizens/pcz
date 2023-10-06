// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package thread

import (
	_ "unsafe" // for go:linkname

	stdgo "github.com/primecitizens/pcz/std/builtin/go"
)

// GetTLSBaseAddress returns the start address of thread local storage
func GetTLSBaseAddress() uintptr

// G returns the G
//
//go:linkname G runtime._getg
func G() *stdgo.GHead

// SetG
//
// NOTE: When changing g, the new g should return the same id from g.G().ID().
//
//go:noescape
func SetG(g *stdgo.GHead)

// Topframe calls pc(arg1) and sets sp, its callframe is marked as top-frame.
//
//go:noescape
func Topframe(arg1, sp, pc uintptr)

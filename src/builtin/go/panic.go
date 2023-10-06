// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdgo

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/atomic"
)

// A Panic holds information about an active panic.
//
// A Panic value must only ever live on the stack.
//
// The argp and link fields are stack pointers, but don't need special
// handling during stack growth: because they are pointer-typed and
// Panic values only live on the stack, regular stack pointer
// adjustment takes care of them.
//
// see $GOROOT/src/runtime/runtime2.go#type:_panic
type Panic struct {
	Argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	Arg       any            // argument to panic
	Link      *Panic         // link to earlier panic
	PC        uintptr        // where to return to in runtime if this panic is bypassed
	SP        unsafe.Pointer // where to return to in runtime if this panic is bypassed
	Recovered bool           // whether this panic is over
	Aborted   bool           // the panic was aborted
	Goexit    bool
}

var panicking atomic.Uint32[uint32]

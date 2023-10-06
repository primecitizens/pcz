// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdgo

import (
	"unsafe"
)

// A Defer holds an entry on the list of deferred calls.
// If you add a field here, add code to clear it in deferProcStack.
// This struct must match the code in cmd/compile/internal/ssagen/ssa.go:deferstruct
// and cmd/compile/internal/ssagen/ssa.go:(*state).call.
// Some defers will be allocated on the stack and some on the heap.
// All defers are logically part of the stack, so write barriers to
// initialize them are not required. All defers must be manually scanned,
// and for heap defers, marked.
//
// See $GOROOT/src/runtime/runtime2.go#type:_defer
type Defer struct {
	Started bool
	Heap    bool
	// OpenDefer indicates that this _defer is for a frame with open-coded
	// defers. We have only one defer record for the entire frame (which may
	// currently have 0, 1, or more defers active).
	OpenDefer bool
	SP        uintptr // sp at time of defer
	PC        uintptr // pc at time of defer
	Fn        func()  // can be nil for open-coded defers
	Panic     *Panic  // panic that is running defer
	Link      *Defer  // next defer on G; can point to either heap or stack!

	// If openDefer is true, the fields below record values about the stack
	// frame and associated function that has the open-coded defer(s). sp
	// above will be the sp for the frame, and pc will be address of the
	// deferreturn call in the function.
	FuncData unsafe.Pointer // funcdata for the function associated with the frame
	Varp     uintptr        // value of varp for the stack frame
	// FramePC is the current pc associated with the stack frame. Together,
	// with sp above (which is the sp associated with the stack frame),
	// FramePC/sp can be used as pc/sp pair to continue a stack trace via
	// gentraceback().
	FramePC uintptr
}

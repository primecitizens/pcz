// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
)

// See ${GOROOT}/src/runtime/panic.go

// deferreturn runs deferred functions for the caller's frame.
// The compiler inserts a call to this at the end of any
// function which calls defer.
func deferreturn() {
	gp := getg()
	for {
		d := gp.Defer
		if d == nil {
			return
		}
		sp := getcallersp()
		if d.SP != sp {
			return
		}
		if d.OpenDefer {
			done := runOpenDeferFrame(d)
			if !done {
				assert.Throw("unfinished", "open-coded", "defers", "in", "deferreturn")
			}
			gp.Defer = d.Link
			freedefer(d)
			// If this frame uses open defers, then this
			// must be the only defer record for the
			// frame, so we can just return.
			return
		}

		fn := d.Fn
		d.Fn = nil
		gp.Defer = d.Link
		freedefer(d)
		fn()
	}
}

// runOpenDeferFrame runs the active open-coded defers in the frame specified by
// d. It normally processes all active defers in the frame, but stops immediately
// if a defer does a successful recover. It returns true if there are no
// remaining defers to run in the frame.
func runOpenDeferFrame(d *_defer) bool {
	done := true
	fd := d.FuncData

	deferBitsOffset, fd := readvarintUnsafe(fd)
	nDefers, fd := readvarintUnsafe(fd)
	deferBits := *(*uint8)(unsafe.Pointer(d.Varp - uintptr(deferBitsOffset)))

	for i := int(nDefers) - 1; i >= 0; i-- {
		// read the funcdata info for this defer
		var closureOffset uint32
		closureOffset, fd = readvarintUnsafe(fd)
		if deferBits&(1<<i) == 0 {
			continue
		}
		closure := *(*func())(unsafe.Pointer(d.Varp - uintptr(closureOffset)))
		d.Fn = closure
		deferBits = deferBits &^ (1 << i)
		*(*uint8)(unsafe.Pointer(d.Varp - uintptr(deferBitsOffset))) = deferBits
		p := d.Panic
		// Call the defer. Note that this can change d.Varp if
		// the stack moves.
		deferCallSave(p, d.Fn)
		if p != nil && p.Aborted {
			break
		}
		d.Fn = nil
		if d.Panic != nil && d.Panic.Recovered {
			done = deferBits == 0
			break
		}
	}

	return done
}

// deferCallSave calls fn() after saving the caller's pc and sp in the
// panic record. This allows the runtime to return to the Goexit defer
// processing loop, in the unusual case where the Goexit may be
// bypassed by a successful recover.
//
// This is marked as a wrapper by the compiler so it doesn't appear in
// tracebacks.
func deferCallSave(p *_panic, fn func()) {
	if p != nil {
		p.Argp = unsafe.Pointer(getargp())
		p.PC = getcallerpc()
		p.SP = unsafe.Pointer(getcallersp())
	}
	fn()
	if p != nil {
		p.PC = 0
		p.SP = unsafe.Pointer(nil)
	}
}

// getargp returns the location where the caller
// writes outgoing function call arguments.
//
//go:nosplit
//go:noinline
func getargp() uintptr {
	return getcallersp() + arch.MinFrameSize
}

// Free the given defer.
// The defer cannot be used after this call.
//
// This is nosplit because the incoming defer is in a perilous state.
// It's not on any defer list, so stack copying won't adjust stack
// pointers in it (namely, d.link). Hence, if we were to copy the
// stack, d could then contain a stale pointer.
//
//go:nosplit
func freedefer(d *_defer) {
	d.Link = nil
	// After this point we can copy the stack.

	if d.Panic != nil {
		freedeferpanic()
	}
	if d.Fn != nil {
		freedeferfn()
	}
	if !d.Heap {
		return
	}

	// mp := acquirem()
	// pp := mp.p.ptr()
	// if len(pp.deferpool) == cap(pp.deferpool) {
	// 	// Transfer half of local cache to the central cache.
	// 	var first, last *_defer
	// 	for len(pp.deferpool) > cap(pp.deferpool)/2 {
	// 		n := len(pp.deferpool)
	// 		d := pp.deferpool[n-1]
	// 		pp.deferpool[n-1] = nil
	// 		pp.deferpool = pp.deferpool[:n-1]
	// 		if first == nil {
	// 			first = d
	// 		} else {
	// 			last.link = d
	// 		}
	// 		last = d
	// 	}
	// 	lock(&sched.deferlock)
	// 	last.link = sched.deferpool
	// 	sched.deferpool = first
	// 	unlock(&sched.deferlock)
	// }

	*d = _defer{}

	// pp.deferpool = append(pp.deferpool, d)
	//
	// releasem(mp)
	// mp, pp = nil, nil
}

// Separate function so that it can split stack.
// Windows otherwise runs out of stack space.
func freedeferpanic() {
	// _panic must be cleared before d is unlinked from gp.
	assert.Throw("freedefer", "with", "d.Panic", "!=", "nil")
}

func freedeferfn() {
	// fn must be cleared before d is unlinked from gp.
	assert.Throw("freedefer", "with", "d.Fn", "!=", "nil")
}

// readvarintUnsafe reads the uint32 in varint format starting at fd, and returns the
// uint32 and a pointer to the byte following the varint.
//
// There is a similar function runtime.readvarint, which takes a slice of bytes,
// rather than an unsafe pointer. These functions are duplicated, because one of
// the two use cases for the functions would get slower if the functions were
// combined.
func readvarintUnsafe(fd unsafe.Pointer) (uint32, unsafe.Pointer) {
	var r uint32
	var shift int
	for {
		b := *(*uint8)((unsafe.Pointer(fd)))
		fd = unsafe.Add(fd, unsafe.Sizeof(b))
		if b < 128 {
			return r + uint32(b)<<shift, fd
		}
		r += ((uint32(b) &^ 128) << shift)
		shift += 7
		if shift > 28 {
			assert.Throw("bad", "varint")
		}
	}
}

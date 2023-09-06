// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdgo

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
)

type Chan struct {
	Qcount    uint           // total data in the queue
	DataqSize uint           // size of the circular queue
	Buf       unsafe.Pointer // points to an array of dataqsiz elements
	ElemSize  uint16
	Closed    uint32
	ElemType  *abi.Type // element type
	Sendx     uint      // send index
	Recvx     uint      // receive index

	// NOTE: linker checks `recvq` and `sendq` field for dwarf
	// 		DO NOT CHANGE THESE NAMES
	recvq waitq // list of recv waiters
	sendq waitq // list of send waiters

	// Lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this Lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	Lock Mutex
}

type Waitq struct {
	// NOTE: linker checks `first` and `last` field for dwarf
	// 		DO NOT CHANGE THESE NAMES
	first *Sudog
	last  *Sudog
}

func MakeChan(chanType *abi.ChanType, sz int) *hchan {
	return nil
}

// Recv receives on channel c and writes the received data to ep.
//
// ep may be nil, in which case received data is ignored.
//
// If block == false and no elements are available, returns (false, false).
// Otherwise, if c is closed, zeros *ep and returns (true, false).
// Otherwise, fills in *ep with an element and returns (true, true).
// A non-nil ep must point to the heap or the caller's stack.
//
//go:nosplit
func Recv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	return
}

// generic single channel send
//
// If block is not nil, then the protocol will not sleep but return if it could
// not complete.
//
// sleep can wake up with g.param == nil when a channel involved in the sleep
// has been closed. it is easiest to loop and re-run the operation;
// we'll see that it's now closed.
//
//go:nosplit
func Send(c *hchan, elem unsafe.Pointer, block bool, callerpc uintptr) bool {
	return false
}

// Select case descriptor.
// Known to compiler.
// Changes here must also be made in src/cmd/compile/internal/walk/select.go's scasetype.
type SelectCase struct {
	C    *hchan         // chan
	Elem unsafe.Pointer // data element
}

// Select implements the select statement.
//
// cas0 points to an array of type [ncases]scase, and order0 points to
// an array of type [2*ncases]uint16 where ncases must be <= 65536.
// Both reside on the goroutine's stack (regardless of any escaping in
// Select).
//
// For race detector builds, pc0 points to an array of type
// [ncases]uintptr (also on the stack); for other builds, it's set to
// nil.
//
// Select returns the index of the chosen scase, which matches the
// ordinal position of its respective select{recv,send,default} call.
// Also, if the chosen scase was a receive operation, it reports whether
// a value was received.
func Select(cas0 *SelectCase, order0 *uint16, pc0 *uintptr, nsends, nrecvs int, block bool) (int, bool) {
	return 0, false
}

func Block() {
	// TODO: Park
}

func CloseChan(c *hchan) {}

func Recover(argp uintptr) any {
	// TODO
	return nil
}

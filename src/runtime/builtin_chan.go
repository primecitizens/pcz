// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"

	stdgo "github.com/primecitizens/std/builtin/go"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
)

// see $GOROOT/src/runtime/chan.go

//
// channel
//

func makechan64(t *abi.ChanType, size int64) *stdgo.Chan {
	if int64(int(size)) != size {
		assert.Panic("makechan: size out of range")
	}

	return makechan(t, int(size))
}

func makechan(t *abi.ChanType, sz int) *stdgo.Chan { return stdgo.MakeChan(t, sz) }

// entry point for x := <- ch from compiled code.
//
//go:nosplit
func chanrecv1(c *stdgo.Chan, elem unsafe.Pointer) { stdgo.Recv(c, elem, true) }

// entry point for x, ok := <- ch from compiled code.
//
//go:nosplit
func chanrecv2(c *stdgo.Chan, elem unsafe.Pointer) (received bool) {
	_, received = stdgo.Recv(c, elem, true)
	return received
}

// entry point for ch <- x from compiled code.
//
//go:nosplit
func chansend1(c *stdgo.Chan, elem unsafe.Pointer) {
	stdgo.Send(c, elem, true, getcallerpc())
}

// entry point for close(ch) from compiled code.
//
//go:nosplit
func closechan(c *stdgo.Chan) {
	stdgo.CloseChan(c)
}

//
// select
//

// compiler implements
//
//	select {
//	case c <- v:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selectnbsend(c, v) {
//		... foo
//	} else {
//		... bar
//	}
func selectnbsend(c *stdgo.Chan, elem unsafe.Pointer) bool {
	return stdgo.Send(c, elem, false, getcallerpc())
}

// compiler implements
//
//	select {
//	case v, ok = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selected, ok = selectnbrecv(&v, c); selected {
//		... foo
//	} else {
//		... bar
//	}
func selectnbrecv(ep unsafe.Pointer, c *stdgo.Chan) (bool, bool) {
	return stdgo.Recv(c, ep, false)
}

func selectsetpc(pc *uintptr) { *pc = getcallerpc() }

func selectgo(cas0 *stdgo.SelectCase, order0 *uint16, pc0 *uintptr, nsends int, nrecvs int, block bool) (int, bool) {
	return stdgo.Select(cas0, order0, pc0, nsends, nrecvs, block)
}

func block() { stdgo.Block() }

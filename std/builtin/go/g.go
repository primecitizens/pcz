// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdgo

import (
	"unsafe"

	stdtype "github.com/primecitizens/pcz/std/builtin/type"
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/mark"
)

// G defines required methods for a goroutine implementation.
type G interface {
	ID() uint64
	Rand32() uint32
	Rand64() uint64

	// Status returns the status of this goroutine.
	Status() Status

	// DefaultAlloc returns the default allocator in this goroutine.
	DefaultAlloc() alloc.M

	// PersistantAlloc returns the persistant allocator in this goroutine.
	PersistantAlloc() alloc.P

	// TODO:
	// StackAllocator() alloc.S
	// HeapAllocator() alloc.H
	// ArenaAllocator() alloc.A

	// // Park puts the current goroutine into a waiting state and calls unlockf
	// // on the system stack.
	// //
	// // If unlockf returns false, the goroutine is resumed.
	// //
	// // unlockf must not access this G's stack, as it may be moved between
	// // the call to gopark and the call to unlockf.
	// //
	// // Note that because unlockf is called after putting the G into a waiting
	// // state, the G may have already been readied by the time unlockf is called
	// // unless there is external synchronization preventing the G from being
	// // readied. If unlockf returns false, it must guarantee that the G cannot be
	// // externally readied.
	// //
	// // Reason explains why the goroutine has been parked. It is displayed in stack
	// // traces and heap dumps. Reasons should be unique and descriptive. Do not
	// // re-use reasons, add new ones.
	// Park(unlockf func(g G, lock unsafe.Pointer) bool, lock unsafe.Pointer, reason WaitReason, traceReason trace.BlockReason)
}

// GHead contains required per-G data.
//
// See ${GOROOT}/src/runtime/runtime2.go#type:g
//
// NOTE:
//   - unsafe.Offsetof(T.GHead.Stack) MUST be 0.
//   - (dev) MUST be kept as non-generic struct type.
type GHead struct {
	// Stack parameters.
	//
	// Stack describes the actual Stack memory: [Stack.lo, Stack.hi).
	// stackguard0 is the Stack pointer compared in the Go Stack growth prologue.
	// It is Stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the Stack pointer compared in the C Stack growth prologue.
	// It is Stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	Stack       Stack   // offset known to runtime/cgo
	Stackguard0 uintptr // offset known to liblink
	Stackguard1 uintptr // offset known to liblink

	Panic      *Panic         // innermost panic - offset known to liblink
	Defer      *Defer         // innermost defer
	M          unsafe.Pointer // current m; offset known to arm liblink
	Sched      Gobuf
	SyscallSP  uintptr // if status==Gsyscall, syscallsp = Sched.sp to use during gc
	SyscallPC  uintptr // if status==Gsyscall, syscallpc = Sched.pc to use during gc
	StacktopSP uintptr // expected sp at top of stack, to check in traceback

	// NOTE(pcz): following fields are not known to compiler/link

	ID_ uint64

	// Itab for the G interface value.
	Itab *abi.Itab
}

func (gp *GHead) Guintptr() Guintptr { return Guintptr(unsafe.Pointer(gp)) }

// G returns the G implementation
//
//go:nosplit
func (gp *GHead) G() G {
	var g G
	iface := stdtype.IfaceOf(mark.NoEscape(&g))
	iface.Itab, iface.Data = mark.NoEscape(gp.Itab), unsafe.Pointer(mark.NoEscape(gp))
	return g
}

func (gp *GHead) ID() uint64 { return gp.ID_ }

// Stack describes a Go execution Stack.
// The bounds of the Stack are exactly [lo, hi),
// with no implicit data structures on either side.
type Stack struct {
	Lo uintptr
	Hi uintptr
}

func (stk Stack) PointerOnStack(ptr uintptr) bool {
	return stk.Lo <= ptr && ptr < stk.Hi
}

// A Guintptr holds a goroutine pointer, but typed as a uintptr
// to bypass write barriers. It is used in the Gobuf goroutine state
// and in scheduling lists that are manipulated without a P.
//
// The Gobuf.G goroutine pointer is almost always updated by assembly code.
// In one of the few places it is updated by Go code - func save - it must be
// treated as a uintptr to avoid a write barrier being emitted at a bad time.
// Instead of figuring out how to emit the write barriers missing in the
// assembly manipulation, we change the type of the field to uintptr,
// so that it does not require write barriers at all.
//
// Goroutine structs are published in the allg list and never freed.
// That will keep the goroutine structs from being collected.
// There is never a time that Gobuf.g's contain the only references
// to a goroutine: the publishing of the goroutine in allg comes first.
// Goroutine pointers are also kept in non-GC-visible places like TLS,
// so I can't see them ever moving. If we did want to start moving data
// in the GC, we'd need to allocate the goroutine structs from an
// alternate arena. Using Guintptr doesn't make that problem any worse.
// Note that pollDesc.rg, pollDesc.wg also store g in uintptr form,
// so they would need to be updated too if g's start moving.
type Guintptr uintptr

//go:nosplit
func (gp Guintptr) Ptr() *GHead { return (*GHead)(unsafe.Pointer(gp)) }

//go:nosplit
func (gp *Guintptr) Set(g *GHead) { *gp = Guintptr(unsafe.Pointer(g)) }

// See ${GOROOT}/src/runtime/runtime2.go#type:gobuf
type Gobuf struct {
	// The offsets of SP, PC, and G are known to (hard-coded in) libmach.
	//
	// Ctxt is unusual with respect to GC: it may be a
	// heap-allocated funcval, so GC needs to track it, but it
	// needs to be set and cleared from assembly, where it's
	// difficult to have write barriers. However, ctxt is really a
	// saved, live register, and we only ever exchange it between
	// the real register and the gobuf. Hence, we treat it as a
	// root during stack scanning, which means assembly that saves
	// and restores it doesn't need write barriers. It's still
	// typed as a pointer so that any other writes from Go get
	// write barriers.
	SP   uintptr
	PC   uintptr
	G    Guintptr
	Ctxt unsafe.Pointer
	Ret  uintptr
	LR   uintptr
	BP   uintptr // for framepointer-enabled architectures
}

// Sudog represents a g in a wait list, such as for sending/receiving
// on a channel.
//
// Sudog is necessary because the g ↔ synchronization object relation
// is many-to-many. A g can be on many wait lists, so there may be
// many sudogs for one g; and many gs may be waiting on the same
// synchronization object, so there may be many sudogs for one object.
//
// sudogs are allocated from a special pool. Use acquireSudog and
// releaseSudog to allocate and free them.
//
// See ${GOROOT}/src/runtime/runtime2.go#type:sudog
type Sudog struct {

	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this for sudogs involved in channel ops.

	g *GHead

	next *Sudog
	prev *Sudog

	// NOTE: linker checks `elem` field for dwarf
	// 		DO NOT CHANGE THESE NAMES
	elem unsafe.Pointer // data element (may point to stack)

	// The following fields are never accessed concurrently.
	// For channels, waitlink is only accessed by g.
	// For semaphores, all fields (including the ones above)
	// are only accessed when holding a semaRoot lock.

	acquiretime int64
	releasetime int64
	ticket      uint32

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool

	// success indicates whether communication over channel c
	// succeeded. It is true if the goroutine was awoken because a
	// value was delivered over channel c, and false if awoken
	// because c was closed.
	success bool

	parent   *Sudog // semaRoot binary tree
	waitlink *Sudog // g.waiting list or semaRoot
	waittail *Sudog // semaRoot
	c        *hchan // channel
}

// G status
//
// Beyond indicating the general state of a G, the G status
// acts like a lock on the goroutine's stack (and hence its
// ability to execute user code).
//
// If you add to this list, add to the list
// of "okay during garbage collection" status
// in mgcmark.go too.
//
// TODO(austin): The StatusGscan bit could be much lighter-weight.
// For example, we could choose not to run StatusGscanrunnable
// goroutines found in the run queue, rather than CAS-looping
// until they become StatusGrunnable. And transitions like
// StatusGscanwaiting -> StatusGscanrunnable are actually okay because
// they don't affect stack ownership.
type Status = uint32

const (
	// StatusGidle means this goroutine was just allocated and has not
	// yet been initialized.
	StatusGidle Status = iota // 0

	// StatusGrunnable means this goroutine is on a run queue. It is
	// not currently executing user code. The stack is not owned.
	StatusGrunnable // 1

	// StatusGrunning means this goroutine may execute user code. The
	// stack is owned by this goroutine. It is not on a run queue.
	// It is assigned an M and a P (g.m and g.m.p are valid).
	StatusGrunning // 2

	// StatusGsyscall means this goroutine is executing a system call.
	// It is not executing user code. The stack is owned by this
	// goroutine. It is not on a run queue. It is assigned an M.
	StatusGsyscall // 3

	// StatusGwaiting means this goroutine is blocked in the runtime.
	// It is not executing user code. It is not on a run queue,
	// but should be recorded somewhere (e.g., a channel wait
	// queue) so it can be ready()d when necessary. The stack is
	// not owned *except* that a channel operation may read or
	// write parts of the stack under the appropriate channel
	// lock. Otherwise, it is not safe to access the stack after a
	// goroutine enters StatusGwaiting (e.g., it may get moved).
	StatusGwaiting // 4

	// StatusGmoribund_unused is currently unused, but hardcoded in gdb
	// scripts.
	StatusGmoribund_unused // 5

	// StatusGdead means this goroutine is currently unused. It may be
	// just exited, on a free list, or just being initialized. It
	// is not executing user code. It may or may not have a stack
	// allocated. The G and its stack (if any) are owned by the M
	// that is exiting the G or that obtained the G from the free
	// list.
	StatusGdead // 6

	// StatusGenqueue_unused is currently unused.
	StatusGenqueue_unused // 7

	// StatusGcopystack means this goroutine's stack is being moved. It
	// is not executing user code and is not on a run queue. The
	// stack is owned by the goroutine that put it in StatusGcopystack.
	StatusGcopystack // 8

	// StatusGpreempted means this goroutine stopped itself for a
	// suspendG preemption. It is like StatusGwaiting, but nothing is
	// yet responsible for ready()ing it. Some suspendG must CAS
	// the status to StatusGwaiting to take responsibility for
	// ready()ing this G.
	StatusGpreempted // 9

	// StatusGscan combined with one of the above states other than
	// StatusGrunning indicates that GC is scanning the stack. The
	// goroutine is not executing user code and the stack is owned
	// by the goroutine that set the StatusGscan bit.
	//
	// StatusGscanrunning is different: it is used to briefly block
	// state transitions while GC signals the G to scan its own
	// stack. This is otherwise like StatusGrunning.
	//
	// atomicstatus&~Gscan gives the state the goroutine will
	// return to when the scan completes.
	StatusGscan          Status = 0x1000
	StatusGscanrunnable  Status = StatusGscan + StatusGrunnable  // 0x1001
	StatusGscanrunning   Status = StatusGscan + StatusGrunning   // 0x1002
	StatusGscansyscall   Status = StatusGscan + StatusGsyscall   // 0x1003
	StatusGscanwaiting   Status = StatusGscan + StatusGwaiting   // 0x1004
	StatusGscanpreempted Status = StatusGscan + StatusGpreempted // 0x1009
)

// A WaitReason explains why a goroutine has been stopped.
// See gopark. Do not re-use waitReasons, add new ones.
type WaitReason uint8

const (
	WaitReasonZero                  WaitReason = iota // ""
	WaitReasonGCAssistMarking                         // "GC assist marking"
	WaitReasonIOWait                                  // "IO wait"
	WaitReasonChanReceiveNilChan                      // "chan receive (nil chan)"
	WaitReasonChanSendNilChan                         // "chan send (nil chan)"
	WaitReasonDumpingHeap                             // "dumping heap"
	WaitReasonGarbageCollection                       // "garbage collection"
	WaitReasonGarbageCollectionScan                   // "garbage collection scan"
	WaitReasonPanicWait                               // "panicwait"
	WaitReasonSelect                                  // "select"
	WaitReasonSelectNoCases                           // "select (no cases)"
	WaitReasonGCAssistWait                            // "GC assist wait"
	WaitReasonGCSweepWait                             // "GC sweep wait"
	WaitReasonGCScavengeWait                          // "GC scavenge wait"
	WaitReasonChanReceive                             // "chan receive"
	WaitReasonChanSend                                // "chan send"
	WaitReasonFinalizerWait                           // "finalizer wait"
	WaitReasonForceGCIdle                             // "force gc (idle)"
	WaitReasonSemacquire                              // "semacquire"
	WaitReasonSleep                                   // "sleep"
	WaitReasonSyncCondWait                            // "sync.Cond.Wait"
	WaitReasonSyncMutexLock                           // "sync.Mutex.Lock"
	WaitReasonSyncRWMutexRLock                        // "sync.RWMutex.RLock"
	WaitReasonSyncRWMutexLock                         // "sync.RWMutex.Lock"
	WaitReasonTraceReaderBlocked                      // "trace reader (blocked)"
	WaitReasonWaitForGCCycle                          // "wait for GC cycle"
	WaitReasonGCWorkerIdle                            // "GC worker (idle)"
	WaitReasonGCWorkerActive                          // "GC worker (active)"
	WaitReasonPreempted                               // "preempted"
	WaitReasonDebugCall                               // "debug call"
	WaitReasonGCMarkTermination                       // "GC mark termination"
	WaitReasonStoppingTheWorld                        // "stopping the world"
)

func (w WaitReason) IsMutexWait() bool {
	return w == WaitReasonSyncMutexLock ||
		w == WaitReasonSyncRWMutexRLock ||
		w == WaitReasonSyncRWMutexLock
}

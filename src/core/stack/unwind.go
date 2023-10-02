// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package stack

import (
	"unsafe"

	stdgo "github.com/primecitizens/std/builtin/go"
	stdprint "github.com/primecitizens/std/builtin/print"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/core/os"
	"github.com/primecitizens/std/core/thread"
)

type hex = uint64

type pcvalueCache struct {
	entries [2][8]pcvalueCacheEnt
}

type pcvalueCacheEnt struct {
	// targetpc and off together are the key of this cache entry.
	targetpc uintptr
	off      uint32
	// val is the value of this cached pcvalue entry.
	val int32
}

// pcvalueCacheKey returns the outermost index in a pcvalueCache to use for targetpc.
// It must be very cheap to calculate.
// For now, align to arch.PtrSize and reduce mod the number of entries.
// In practice, this appears to be fairly randomly and evenly distributed.
func pcvalueCacheKey(targetpc uintptr) uintptr {
	return (targetpc / arch.PtrSize) % uintptr(len(pcvalueCache{}.entries))
}

const debugPcln = false

func funcspdelta(f abi.FuncInfo, targetpc uintptr, cache *pcvalueCache) int32 {
	x, _ := pcvalue(f, f.PCSP, targetpc, cache, true)
	if debugPcln && x&(arch.PtrSize-1) != 0 {
		print("invalid spdelta ", f.Name(), " ", hex(f.Entry()), " ", hex(targetpc), " ", hex(f.PCSP), " ", x, "\n")
		assert.Throw("bad", "spdelta")
	}
	return x
}

// Returns the PCData value, and the PC where this value starts.
// TODO: the start PC is returned only when cache is nil.
func pcvalue(f abi.FuncInfo, off uint32, targetpc uintptr, cache *pcvalueCache, strict bool) (int32, uintptr) {
	if off == 0 {
		return -1, 0
	}

	// Check the cache. This speeds up walks of deep stacks, which
	// tend to have the same recursive functions over and over.
	//
	// This cache is small enough that full associativity is
	// cheaper than doing the hashing for a less associative
	// cache.
	if cache != nil {
		x := pcvalueCacheKey(targetpc)
		for i := range cache.entries[x] {
			// We check off first because we're more
			// likely to have multiple entries with
			// different offsets for the same targetpc
			// than the other way around, so we'll usually
			// fail in the first clause.
			ent := &cache.entries[x][i]
			if ent.off == off && ent.targetpc == targetpc {
				return ent.val, 0
			}
		}
	}

	if !f.Valid() {
		if strict && panicking.Load() == 0 {
			println("runtime:", "no", "module", "data", "for", hex(f.Entry()))
			assert.Throw("no", "module", "data")
		}
		return -1, 0
	}
	datap := f.Datap
	p := datap.PCTab[off:]
	pc := f.Entry()
	prevpc := pc
	val := int32(-1)
	for {
		var ok bool
		p, ok = step(p, &pc, &val, pc == f.Entry())
		if !ok {
			break
		}
		if targetpc < pc {
			// Replace a random entry in the cache. Random
			// replacement prevents a performance cliff if
			// a recursive stack's cycle is slightly
			// larger than the cache.
			// Put the new element at the beginning,
			// since it is the most likely to be newly used.
			if cache != nil {
				x := pcvalueCacheKey(targetpc)
				e := &cache.entries[x]
				ci := fastrandn(uint32(len(cache.entries[x])))
				e[ci] = e[0]
				e[0] = pcvalueCacheEnt{
					targetpc: targetpc,
					off:      off,
					val:      val,
				}
			}

			return val, prevpc
		}
		prevpc = pc
	}

	// If there was a table, it should have covered all program counters.
	// If not, something is wrong.
	if panicking.Load() != 0 || !strict {
		return -1, 0
	}

	print(
		"runtime:", " ", "invalid", " ", "pc-encoded", " ", "table",
		" ", "f=", f.Name(),
		" ", "pc=", hex(pc),
		" ", "targetpc=", hex(targetpc),
		" ", "tab=", p, "\n",
	)

	p = datap.PCTab[off:]
	pc = f.Entry()
	val = -1
	for {
		var ok bool
		p, ok = step(p, &pc, &val, pc == f.Entry())
		if !ok {
			break
		}
		print("\tvalue=", val, " until pc=", hex(pc), "\n")
	}

	assert.Throw("invalid runtime symbol table")
	return -1, 0
}

// step advances to the next pc, value pair in the encoded table.
func step(p []byte, pc *uintptr, val *int32, first bool) (newp []byte, ok bool) {
	// For both uvdelta and pcdelta, the common case (~70%)
	// is that they are a single byte. If so, avoid calling readvarint.
	uvdelta := uint32(p[0])
	if uvdelta == 0 && !first {
		return nil, false
	}
	n := uint32(1)
	if uvdelta&0x80 != 0 {
		n, uvdelta = readvarint(p)
	}
	*val += int32(-(uvdelta & 1) ^ (uvdelta >> 1))
	p = p[n:]

	pcdelta := uint32(p[0])
	n = 1
	if pcdelta&0x80 != 0 {
		n, pcdelta = readvarint(p)
	}
	p = p[n:]
	*pc += uintptr(pcdelta * arch.PCQuantum)
	return p, true
}

// readvarint reads a varint from p.
func readvarint(p []byte) (read uint32, val uint32) {
	var v, shift, n uint32
	for {
		b := p[n]
		n++
		v |= uint32(b&0x7F) << (shift & 31)
		if b&0x80 == 0 {
			break
		}
		shift += 7
	}
	return n, v
}

// The code in this file implements stack trace walking for all architectures.
// The most important fact about a given architecture is whether it uses a link register.
// On systems with link registers, the prologue for a non-leaf function stores the
// incoming value of LR at the bottom of the newly allocated stack frame.
// On systems without link registers (x86), the architecture pushes a return PC during
// the call instruction, so the return PC ends up above the stack frame.
// In this file, the return PC is always called LR, no matter how it was found.

const usesLR = arch.MinFrameSize > 0

const (
	// tracebackInnerFrames is the number of innermost frames to print in a
	// stack trace. The total maximum frames is tracebackInnerFrames +
	// tracebackOuterFrames.
	tracebackInnerFrames = 50

	// tracebackOuterFrames is the number of outermost frames to print in a
	// stack trace.
	tracebackOuterFrames = 50
)

// unwindFlags control the behavior of various unwinders.
type unwindFlags uint8

const (
	// unwindPrintErrors indicates that if unwinding encounters an error, it
	// should print a message and stop without throwing. This is used for things
	// like stack printing, where it's better to get incomplete information than
	// to crash. This is also used in situations where everything may not be
	// stopped nicely and the stack walk may not be able to complete, such as
	// during profiling signals or during a crash.
	//
	// If neither unwindPrintErrors or unwindSilentErrors are set, unwinding
	// performs extra consistency checks and throws on any error.
	//
	// Note that there are a small number of fatal situations that will throw
	// regardless of unwindPrintErrors or unwindSilentErrors.
	unwindPrintErrors unwindFlags = 1 << iota

	// unwindSilentErrors silently ignores errors during unwinding.
	unwindSilentErrors

	// unwindTrap indicates that the initial PC and SP are from a trap, not a
	// return PC from a call.
	//
	// The unwindTrap flag is updated during unwinding. If set, frame.pc is the
	// address of a faulting instruction instead of the return address of a
	// call. It also means the liveness at pc may not be known.
	//
	// TODO: Distinguish frame.continpc, which is really the stack map PC, from
	// the actual continuation PC, which is computed differently depending on
	// this flag and a few other things.
	unwindTrap

	// unwindJumpStack indicates that, if the traceback is on a system stack, it
	// should resume tracing at the user stack when the system stack is
	// exhausted.
	unwindJumpStack
)

// An unwinder iterates the physical stack frames of a Go stack.
//
// Typical use of an unwinder looks like:
//
//	var u unwinder
//	for u.init(gp, 0); u.Valid(); u.next() {
//		// ... use frame info in u ...
//	}
//
// Implementation note: This is carefully structured to be pointer-free because
// tracebacks happen in places that disallow write barriers (e.g., signals).
// Even if this is stack-allocated, its pointer-receiver methods don't know that
// their receiver is on the stack, so they still emit write barriers. Here we
// address that by carefully avoiding any pointers in this type. Another
// approach would be to split this into a mutable part that's passed by pointer
// but contains no pointers itself and an immutable part that's passed and
// returned by value and can contain pointers. We could potentially hide that
// we're doing that in trivial methods that are inlined into the caller that has
// the stack allocation, but that's fragile.
type unwinder struct {
	// frame is the current physical stack frame, or all 0s if
	// there is no frame.
	frame Frame

	// g is the G who's stack is being unwound. If the
	// unwindJumpStack flag is set and the unwinder jumps stacks,
	// this will be different from the initial G.
	g stdgo.Guintptr

	// cgoCtxt is the index into g.cgoCtxt of the next frame on the cgo stack.
	// The cgo stack is unwound in tandem with the Go stack as we find marker frames.
	cgoCtxt int

	// calleeFuncID is the function ID of the caller of the current
	// frame.
	calleeFuncID abi.FuncID

	// flags are the flags to this unwind. Some of these are updated as we
	// unwind (see the flags documentation).
	flags unwindFlags

	// cache is used to cache pcvalue lookups.
	cache pcvalueCache
}

// init initializes u to start unwinding gp's stack and positions the
// iterator on gp's innermost frame. gp must not be the current G.
//
// A single unwinder can be reused for multiple unwinds.
func (u *unwinder) init(gp *stdgo.GHead, flags unwindFlags) {
	// Implementation note: This starts the iterator on the first frame and we
	// provide a "valid" method. Alternatively, this could start in a "before
	// the first frame" state and "next" could return whether it was able to
	// move to the next frame, but that's both more awkward to use in a "for"
	// loop and is harder to implement because we have to do things differently
	// for the first frame.
	u.initAt(^uintptr(0), ^uintptr(0), ^uintptr(0), gp, flags)
}

func (u *unwinder) initAt(pc0, sp0, lr0 uintptr, gp *stdgo.GHead, flags unwindFlags) {
	// Don't call this "g"; it's too easy get "g" and "gp" confused.
	if ourg := thread.G(); ourg == gp && ourg == ourg.M.curg {
		// The starting sp has been passed in as a uintptr, and the caller may
		// have other uintptr-typed stack references as well.
		// If during one of the calls that got us here or during one of the
		// callbacks below the stack must be grown, all these uintptr references
		// to the stack will not be updated, and traceback will continue
		// to inspect the old stack memory, which may no longer be valid.
		// Even if all the variables were updated correctly, it is not clear that
		// we want to expose a traceback that begins on one stack and ends
		// on another stack. That could confuse callers quite a bit.
		// Instead, we require that initAt and any other function that
		// accepts an sp for the current goroutine (typically obtained by
		// calling getcallersp) must not run on that goroutine's stack but
		// instead on the g0 stack.
		assert.Throw("cannot", "trace", "user", "goroutine", "on", "its", "own", "stack")
	}

	if pc0 == ^uintptr(0) && sp0 == ^uintptr(0) { // Signal to fetch saved values from gp.
		if gp.SyscallSP != 0 {
			pc0 = gp.SyscallPC
			sp0 = gp.SyscallSP
			if usesLR {
				lr0 = 0
			}
		} else {
			pc0 = gp.Sched.PC
			sp0 = gp.Sched.SP
			if usesLR {
				lr0 = gp.Sched.LR
			}
		}
	}

	var frame Frame
	frame.pc = pc0
	frame.sp = sp0
	if usesLR {
		frame.lr = lr0
	}

	// If the PC is zero, it's likely a nil function call.
	// Start in the caller's frame.
	if frame.pc == 0 {
		if usesLR {
			frame.pc = *(*uintptr)(unsafe.Pointer(frame.sp))
			frame.lr = 0
		} else {
			frame.pc = uintptr(*(*uintptr)(unsafe.Pointer(frame.sp)))
			frame.sp += arch.PtrSize
		}
	}

	// runtime/internal/atomic functions call into kernel helpers on
	// arm < 7. See runtime/internal/atomic/sys_linux_arm.s.
	//
	// Start in the caller's frame.
	if arch.IsArm&os.IsLinux != 0 && goarm < 7 && frame.pc&0xffff0000 == 0xffff0000 {
		// Note that the calls are simple BL without pushing the return
		// address, so we use LR directly.
		//
		// The kernel helpers are frameless leaf functions, so SP and
		// LR are not touched.
		frame.pc = frame.lr
		frame.lr = 0
	}

	f := abi.FindFunc(frame.pc)
	if !f.Valid() {
		if flags&unwindSilentErrors == 0 {
			print("runtime: g ", gp.ID, ": unknown pc ", hex(frame.pc), "\n")
			tracebackHexdump(gp.Stack, &frame, 0)
		}
		if flags&(unwindPrintErrors|unwindSilentErrors) == 0 {
			assert.Throw("unknown", "pc")
		}
		*u = unwinder{}
		return
	}
	frame.fn = f

	// Populate the unwinder.
	*u = unwinder{
		frame:        frame,
		g:            gp.Guintptr(),
		cgoCtxt:      len(gp.cgoCtxt) - 1,
		calleeFuncID: abi.FuncIDNormal,
		flags:        flags,
	}

	isSyscall := frame.pc == pc0 && frame.sp == sp0 && pc0 == gp.SyscallPC && sp0 == gp.SyscallSP
	u.resolveInternal(true, isSyscall)
}

func (u *unwinder) valid() bool {
	return u.frame.pc != 0
}

// resolveInternal fills in u.frame based on u.frame.fn, pc, and sp.
//
// innermost indicates that this is the first resolve on this stack. If
// innermost is set, isSyscall indicates that the PC/SP was retrieved from
// gp.syscall*; this is otherwise ignored.
//
// On entry, u.frame contains:
//   - fn is the running function.
//   - pc is the PC in the running function.
//   - sp is the stack pointer at that program counter.
//   - For the innermost frame on LR machines, lr is the program counter that called fn.
//
// On return, u.frame contains:
//   - fp is the stack pointer of the caller.
//   - lr is the program counter that called fn.
//   - varp, argp, and continpc are populated for the current frame.
//
// If fn is a stack-jumping function, resolveInternal can change the entire
// frame state to follow that stack jump.
//
// This is internal to unwinder.
func (u *unwinder) resolveInternal(innermost, isSyscall bool) {
	frame := &u.frame
	gp := u.g.Ptr()

	f := frame.fn
	if f.PCSP == 0 {
		// No frame information, must be external function, like race support.
		// See golang.org/issue/13568.
		u.finishInternal()
		return
	}

	// Compute function info flags.
	flag := f.Flag
	if f.FuncID == abi.FuncID_cgocallback {
		// cgocallback does write SP to switch from the g0 to the curg stack,
		// but it carefully arranges that during the transition BOTH stacks
		// have cgocallback frame valid for unwinding through.
		// So we don't need to exclude it with the other SP-writing functions.
		flag &^= abi.FuncFlagSPWrite
	}
	if isSyscall {
		// Some Syscall functions write to SP, but they do so only after
		// saving the entry PC/SP using entersyscall.
		// Since we are using the entry PC/SP, the later SP write doesn't matter.
		flag &^= abi.FuncFlagSPWrite
	}

	// Found an actual function.
	// Derive frame pointer.
	if frame.fp == 0 {
		// Jump over system stack transitions. If we're on g0 and there's a user
		// goroutine, try to jump. Otherwise this is a regular call.
		// We also defensively check that this won't switch M's on us,
		// which could happen at critical points in the scheduler.
		// This ensures gp.m doesn't change from a stack jump.
		if u.flags&unwindJumpStack != 0 && gp == gp.M.g0 && gp.M.curg != nil && gp.M.curg.M == gp.M {
			switch f.FuncID {
			case abi.FuncID_morestack:
				// morestack does not return normally -- newstack()
				// gogo's to curg.sched. Match that.
				// This keeps morestack() from showing up in the backtrace,
				// but that makes some sense since it'll never be returned
				// to.
				gp = gp.M.curg
				u.g.Set(gp)
				frame.pc = gp.Sched.PC
				frame.fn = abi.FindFunc(frame.pc)
				f = frame.fn
				flag = f.Flag
				frame.lr = gp.Sched.LR
				frame.sp = gp.Sched.SP
				u.cgoCtxt = len(gp.cgoCtxt) - 1
			case abi.FuncID_systemstack:
				// systemstack returns normally, so just follow the
				// stack transition.
				if usesLR && funcspdelta(f, frame.pc, &u.cache) == 0 {
					// We're at the function prologue and the stack
					// switch hasn't happened, or epilogue where we're
					// about to return. Just unwind normally.
					// Do this only on LR machines because on x86
					// systemstack doesn't have an SP delta (the CALL
					// instruction opens the frame), therefore no way
					// to check.
					flag &^= abi.FuncFlagSPWrite
					break
				}
				gp = gp.M.curg
				u.g.Set(gp)
				frame.sp = gp.Sched.SP
				u.cgoCtxt = len(gp.cgoCtxt) - 1
				flag &^= abi.FuncFlagSPWrite
			}
		}
		frame.fp = frame.sp + uintptr(funcspdelta(f, frame.pc, &u.cache))
		if !usesLR {
			// On x86, call instruction pushes return PC before entering new function.
			frame.fp += arch.PtrSize
		}
	}

	// Derive link register.
	if flag&abi.FuncFlagTopFrame != 0 {
		// This function marks the top of the stack. Stop the traceback.
		frame.lr = 0
	} else if flag&abi.FuncFlagSPWrite != 0 {
		// The function we are in does a write to SP that we don't know
		// how to encode in the spdelta table. Examples include context
		// switch routines like runtime.gogo but also any code that switches
		// to the g0 stack to run host C code.
		if u.flags&(unwindPrintErrors|unwindSilentErrors) != 0 {
			// We can't reliably unwind the SP (we might
			// not even be on the stack we think we are),
			// so stop the traceback here.
			frame.lr = 0
		} else {
			// For a GC stack traversal, we should only see
			// an SPWRITE function when it has voluntarily preempted itself on entry
			// during the stack growth check. In that case, the function has
			// not yet had a chance to do any writes to SP and is safe to unwind.
			// isAsyncSafePoint does not allow assembly functions to be async preempted,
			// and preemptPark double-checks that SPWRITE functions are not async preempted.
			// So for GC stack traversal, we can safely ignore SPWRITE for the innermost frame,
			// but farther up the stack we'd better not find any.
			if !innermost {
				println("traceback:", "unexpected", "SPWRITE", "function", f.Name())
				assert.Throw("traceback")
			}
		}
	} else {
		var lrPtr uintptr
		if usesLR {
			if innermost && frame.sp < frame.fp || frame.lr == 0 {
				lrPtr = frame.sp
				frame.lr = *(*uintptr)(unsafe.Pointer(lrPtr))
			}
		} else {
			if frame.lr == 0 {
				lrPtr = frame.fp - arch.PtrSize
				frame.lr = *(*uintptr)(unsafe.Pointer(lrPtr))
			}
		}
	}

	frame.varp = frame.fp
	if !usesLR {
		// On x86, call instruction pushes return PC before entering new function.
		frame.varp -= arch.PtrSize
	}

	// For architectures with frame pointers, if there's
	// a frame, then there's a saved frame pointer here.
	//
	// NOTE: This code is not as general as it looks.
	// On x86, the ABI is to save the frame pointer word at the
	// top of the stack frame, so we have to back down over it.
	// On arm64, the frame pointer should be at the bottom of
	// the stack (with R29 (aka FP) = RSP), in which case we would
	// not want to do the subtraction here. But we started out without
	// any frame pointer, and when we wanted to add it, we didn't
	// want to break all the assembly doing direct writes to 8(RSP)
	// to set the first parameter to a called function.
	// So we decided to write the FP link *below* the stack pointer
	// (with R29 = RSP - 8 in Go functions).
	// This is technically ABI-compatible but not standard.
	// And it happens to end up mimicking the x86 layout.
	// Other architectures may make different decisions.
	if frame.varp > frame.sp && arch.FramePointerEnabled {
		frame.varp -= arch.PtrSize
	}

	frame.argp = frame.fp + arch.MinFrameSize

	// Determine frame's 'continuation PC', where it can continue.
	// Normally this is the return address on the stack, but if sigpanic
	// is immediately below this function on the stack, then the frame
	// stopped executing due to a trap, and frame.pc is probably not
	// a safe point for looking up liveness information. In this panicking case,
	// the function either doesn't return at all (if it has no defers or if the
	// defers do not recover) or it returns from one of the calls to
	// deferproc a second time (if the corresponding deferred func recovers).
	// In the latter case, use a deferreturn call site as the continuation pc.
	frame.continpc = frame.pc
	if u.calleeFuncID == abi.FuncID_sigpanic {
		if frame.fn.DeferReturn != 0 {
			frame.continpc = frame.fn.Entry() + uintptr(frame.fn.DeferReturn) + 1
			// Note: this may perhaps keep return variables alive longer than
			// strictly necessary, as we are using "function has a defer statement"
			// as a proxy for "function actually deferred something". It seems
			// to be a minor drawback. (We used to actually look through the
			// gp._defer for a defer corresponding to this function, but that
			// is hard to do with defer records on the stack during a stack copy.)
			// Note: the +1 is to offset the -1 that
			// stack.go:getStackMap does to back up a return
			// address make sure the pc is in the CALL instruction.
		} else {
			frame.continpc = 0
		}
	}
}

func (u *unwinder) next() {
	frame := &u.frame
	f := frame.fn
	gp := u.g.Ptr()

	// Do not unwind past the bottom of the stack.
	if frame.lr == 0 {
		u.finishInternal()
		return
	}
	flr := abi.FindFunc(frame.lr)
	if !flr.Valid() {
		// This happens if you get a profiling interrupt at just the wrong time.
		// In that context it is okay to stop early.
		// But if no error flags are set, we're doing a garbage collection and must
		// get everything, so crash loudly.
		fail := u.flags&(unwindPrintErrors|unwindSilentErrors) == 0
		doPrint := u.flags&unwindSilentErrors == 0
		if doPrint && gp.M.incgo && f.FuncID == abi.FuncID_sigpanic {
			// We can inject sigpanic
			// calls directly into C code,
			// in which case we'll see a C
			// return PC. Don't complain.
			doPrint = false
		}
		if fail || doPrint {
			print("runtime: g ", gp.ID_, ": unexpected return pc for ", f.Name(), " called from ", hex(frame.lr), "\n")
			tracebackHexdump(gp.Stack, frame, 0)
		}
		if fail {
			assert.Throw("unknown", "caller", "pc")
		}
		frame.lr = 0
		u.finishInternal()
		return
	}

	if frame.pc == frame.lr && frame.sp == frame.fp {
		// If the next frame is identical to the current frame, we cannot make progress.
		print("runtime: traceback stuck. pc=", hex(frame.pc), " sp=", hex(frame.sp), "\n")
		tracebackHexdump(gp.Stack, frame, frame.sp)
		assert.Throw("traceback", "stuck")
	}

	injectedCall := f.FuncID == abi.FuncID_sigpanic || f.FuncID == abi.FuncID_asyncPreempt || f.FuncID == abi.FuncID_debugCallV2
	if injectedCall {
		u.flags |= unwindTrap
	} else {
		u.flags &^= unwindTrap
	}

	// Unwind to next frame.
	u.calleeFuncID = f.FuncID
	frame.fn = flr
	frame.pc = frame.lr
	frame.lr = 0
	frame.sp = frame.fp
	frame.fp = 0

	// On link register architectures, sighandler saves the LR on stack
	// before faking a call.
	if usesLR && injectedCall {
		x := *(*uintptr)(unsafe.Pointer(frame.sp))
		frame.sp += num.AlignUp[uintptr](arch.MinFrameSize, arch.StackAlign)
		f = abi.FindFunc(frame.pc)
		frame.fn = f
		if !f.Valid() {
			frame.pc = x
		} else if funcspdelta(f, frame.pc, &u.cache) == 0 {
			frame.lr = x
		}
	}

	u.resolveInternal(false, false)
}

// tracebackHexdump hexdumps part of stk around frame.sp and frame.fp
// for debugging purposes. If the address bad is included in the
// hexdumped range, it will mark it as well.
func tracebackHexdump(stk stdgo.Stack, frame *Frame, bad uintptr) {
	const expand = 32 * arch.PtrSize
	const maxExpand = 256 * arch.PtrSize
	// Start around frame.sp.
	lo, hi := frame.sp, frame.sp
	// Expand to include frame.fp.
	if frame.fp != 0 && frame.fp < lo {
		lo = frame.fp
	}
	if frame.fp != 0 && frame.fp > hi {
		hi = frame.fp
	}
	// Expand a bit more.
	lo, hi = lo-expand, hi+expand
	// But don't go too far from frame.sp.
	if lo < frame.sp-maxExpand {
		lo = frame.sp - maxExpand
	}
	if hi > frame.sp+maxExpand {
		hi = frame.sp + maxExpand
	}
	// And don't go outside the stack bounds.
	if lo < stk.Lo {
		lo = stk.Lo
	}
	if hi > stk.Hi {
		hi = stk.Hi
	}

	// Print the hex dump.
	print("stack: frame={sp:", hex(frame.sp), ", fp:", hex(frame.fp), "} stack=[", hex(stk.Lo), ",", hex(stk.hi), ")\n")
	hexdumpWords(lo, hi, func(p uintptr) byte {
		switch p {
		case frame.fp:
			return '>'
		case frame.sp:
			return '<'
		case bad:
			return '!'
		}
		return 0
	})
}

// hexdumpWords prints a word-oriented hex dump of [p, end).
//
// If mark != nil, it will be called with each printed word's address
// and should return a character mark to appear just before that
// word's value. It can return 0 to indicate no mark.
func hexdumpWords(p, end uintptr, mark func(uintptr) byte) {
	stdprint.PrintLock()
	var markbuf [1]byte
	markbuf[0] = ' '

	for i := uintptr(0); p+i < end; i += arch.PtrSize {
		if i%16 == 0 {
			if i != 0 {
				println()
			}
			print(hex(p+i), ": ")
		}

		if mark != nil {
			markbuf[0] = mark(p + i)
			if markbuf[0] == 0 {
				markbuf[0] = ' '
			}
		}
		gwrite(markbuf[:])
		val := *(*uintptr)(unsafe.Pointer(p + i))
		print(hex(val))
		print(" ")

		// Can we symbolize val?
		fn := abi.FindFunc(val)
		if fn.Valid() {
			print("<", fn.Name(), "+", hex(val-fn.Entry()), "> ")
		}
	}

	println()
	stdprint.PrintUnlock()
}

// finishInternal is an unwinder-internal helper called after the stack has been
// exhausted. It sets the unwinder to an invalid state and checks that it
// successfully unwound the entire stack.
func (u *unwinder) finishInternal() {
	u.frame.pc = 0

	// Note that panic != nil is okay here: there can be leftover panics,
	// because the defers on the panic stack do not nest in frame order as
	// they do on the defer stack. If you have:
	//
	//	frame 1 defers d1
	//	frame 2 defers d2
	//	frame 3 defers d3
	//	frame 4 panics
	//	frame 4's panic starts running defers
	//	frame 5, running d3, defers d4
	//	frame 5 panics
	//	frame 5's panic starts running defers
	//	frame 6, running d4, garbage collects
	//	frame 6, running d2, garbage collects
	//
	// During the execution of d4, the panic stack is d4 -> d3, which
	// is nested properly, and we'll treat frame 3 as resumable, because we
	// can find d3. (And in fact frame 3 is resumable. If d4 recovers
	// and frame 5 continues running, d3, d3 can recover and we'll
	// resume execution in (returning from) frame 3.)
	//
	// During the execution of d2, however, the panic stack is d2 -> d3,
	// which is inverted. The scan will match d2 to frame 2 but having
	// d2 on the stack until then means it will not match d3 to frame 3.
	// This is okay: if we're running d2, then all the defers after d2 have
	// completed and their corresponding frames are dead. Not finding d3
	// for frame 3 means we'll set frame 3's continpc == 0, which is correct
	// (frame 3 is dead). At the end of the walk the panic stack can thus
	// contain defers (d3 in this case) for dead frames. The inversion here
	// always indicates a dead frame, and the effect of the inversion on the
	// scan is to hide those dead frames, so the scan is still okay:
	// what's left on the panic stack are exactly (and only) the dead frames.
	//
	// We require callback != nil here because only when callback != nil
	// do we know that gentraceback is being called in a "must be correct"
	// context as opposed to a "best effort" context. The tracebacks with
	// callbacks only happen when everything is stopped nicely.
	// At other times, such as when gathering a stack for a profiling signal
	// or when printing a traceback during a crash, everything may not be
	// stopped nicely, and the stack walk may not be able to complete.
	gp := u.g.Ptr()
	if u.flags&(unwindPrintErrors|unwindSilentErrors) == 0 && u.frame.sp != gp.StacktopSP {
		print("runtime: g", gp.ID_, ": frame.sp=", hex(u.frame.sp), " top=", hex(gp.StacktopSP), "\n")
		print("\tstack=[", hex(gp.Stack.Lo), "-", hex(gp.Stack.Hi), "\n")
		assert.Throw("traceback", "did", "not", "unwind", "completely")
	}
}

// symPC returns the PC that should be used for symbolizing the current frame.
// Specifically, this is the PC of the last instruction executed in this frame.
//
// If this frame did a normal call, then frame.pc is a return PC, so this will
// return frame.pc-1, which points into the CALL instruction. If the frame was
// interrupted by a signal (e.g., profiler, segv, etc) then frame.pc is for the
// trapped instruction, so this returns frame.pc. See issue #34123. Finally,
// frame.pc can be at function entry when the frame is initialized without
// actually running code, like in runtime.mstart, in which case this returns
// frame.pc because that's the best we can do.
func (u *unwinder) symPC() uintptr {
	if u.flags&unwindTrap == 0 && u.frame.pc > u.frame.fn.Entry() {
		// Regular call.
		return u.frame.pc - 1
	}
	// Trapping instruction or we're at the function entry point.
	return u.frame.pc
}

// cgoCallers populates pcBuf with the cgo callers of the current frame using
// the registered cgo unwinder. It returns the number of PCs written to pcBuf.
// If the current frame is not a cgo frame or if there's no registered cgo
// unwinder, it returns 0.
func (u *unwinder) cgoCallers(pcBuf []uintptr) int {
	if cgoTraceback == nil || u.frame.fn.FuncID != abi.FuncID_cgocallback || u.cgoCtxt < 0 {
		// We don't have a cgo unwinder (typical case), or we do but we're not
		// in a cgo frame or we're out of cgo context.
		return 0
	}

	ctxt := u.g.Ptr().cgoCtxt[u.cgoCtxt]
	u.cgoCtxt--
	cgoContextPCs(ctxt, pcBuf)
	for i, pc := range pcBuf {
		if pc == 0 {
			return i
		}
	}
	return len(pcBuf)
}

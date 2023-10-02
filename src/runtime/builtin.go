// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	stdcomplex "github.com/primecitizens/std/builtin/complex"
	stdgo "github.com/primecitizens/std/builtin/go"
	stdprint "github.com/primecitizens/std/builtin/print"
	stdptr "github.com/primecitizens/std/builtin/ptr"
	stdtype "github.com/primecitizens/std/builtin/type"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/asan"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/cover"
	"github.com/primecitizens/std/core/emu64"
	"github.com/primecitizens/std/core/fuzz"
	"github.com/primecitizens/std/core/hash"
	"github.com/primecitizens/std/core/math"
	"github.com/primecitizens/std/core/mem"
	"github.com/primecitizens/std/core/msan"
	"github.com/primecitizens/std/core/race"
)

//
// rand
//

func fastrand() uint32 {
	return getg().G().Rand32()
}

//
// new
//

func newobject(typ *abi.Type) unsafe.Pointer {
	if typ.Size_ == 0 {
		return alloc.ZeroSized()
	}

	return getg().G().DefaultAlloc().Malloc(typ, 1, true)
}

// > Allocate an object of size bytes.
// > Small objects are allocated from the per-P cache's free lists.
// > Large objects (> 32 kB) are allocated straight from the heap.
//
// See ${GOROOT}/src/runtime/malloc.go#func:mallocgc
func mallocgc(size uintptr, typ *abi.Type, needszero bool) unsafe.Pointer {
	if typ.Size_ == 0 {
		return alloc.ZeroSized()
	}

	assert.Panic("invalid", "implicit", "allocation")
	return nil
}

//
// goroutine
//

func gopanic(v any) {
	assert.Panic(v)
}

// GoSchedGuarded yields the processor like gosched, but also checks
// for forbidden states and opts out of the yield in those cases.
//
//go:nosplit
func goschedguarded() { stdgo.Sched(true) }

// The implementation of the predeclared function recover.
// Cannot split the stack because it needs to reliably
// find the stack segment of its caller.
//
// TODO(rsc): Once we commit to CopyStackAlways,
// this doesn't need to be nosplit.
//
//go:nosplit
func gorecover(argp uintptr) any { return stdgo.Recover(argp) }

//
// print
//

func printbool(b bool)              { stdprint.PrintBool(b) }
func printfloat(n float64)          { stdprint.PrintFloat(n) }
func printint(n int64)              { stdprint.PrintInt(n) }
func printhex(n uint64)             { stdprint.PrintHex(n) }
func printuint(n uint64)            { stdprint.PrintUint(n) }
func printcomplex(n complex128)     { stdprint.PrintComplex(n) }
func printstring(a string)          { stdprint.PrintString(a) }
func printpointer(p unsafe.Pointer) { stdprint.PrintPointer(p) }
func printuintptr(p uintptr)        { stdprint.PrintUintptr(p) }
func printiface(i stdtype.Iface)    { stdprint.PrintIface(i) }
func printeface(e stdtype.Eface)    { stdprint.PrintEface(e) }
func printslice(s []byte)           { stdprint.PrintSlice(s) }
func printnl()                      { stdprint.PrintNewline() }
func printsp()                      { stdprint.PrintSpace() }
func printlock()                    { stdprint.PrintLock() }
func printunlock()                  { stdprint.PrintUnlock() }

//
// mem
//

func memmove(to, from unsafe.Pointer, n uintptr) {
	mem.Move(to, from, n)
}

// See ${GOROOT}/src/runtime/mbarrier.go#func:typedmemmove
func typedmemmove(typ *abi.Type, dst, src unsafe.Pointer) {
	if dst == src {
		return
	}
	if writeBarrier.needed && typ.PtrBytes != 0 {
		bulkBarrierPreWrite(uintptr(dst), uintptr(src), typ.PtrBytes)
	}
	// There's a race here: if some other goroutine can write to
	// src, it may change some pointer in src after we've
	// performed the write barrier but before we perform the
	// memory copy. This safe because the write performed by that
	// other goroutine must also be accompanied by a write
	// barrier, so at worst we've unnecessarily greyed the old
	// pointer that was in src.
	mem.Move(dst, src, typ.Size_)
}

// See ${GOROOT}/src/runtime/stubs.go#func:memclrNoHeapPointers
func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr) {
	mem.Clear(ptr, n)
}

// See ${GOROOT}/src/runtime/mbarrier.go#func:memclrHasPointers
func memclrHasPointers(ptr unsafe.Pointer, n uintptr) {
	bulkBarrierPreWrite(uintptr(ptr), 0, n)
	mem.Clear(ptr, n)
}

// typedmemclr clears the typed memory at ptr with type typ. The
// memory at ptr must already be initialized (and hence in type-safe
// state). If the memory is being initialized for the first time, see
// memclrNoHeapPointers.
//
// If the caller knows that typ has pointers, it can alternatively
// call memclrHasPointers.
//
// TODO: A "go:nosplitrec" annotation would be perfect for this.
//
// See ${GOROOT}/src/runtime/mbarrier.go#func:typedmemclr
//
//go:nosplit
func typedmemclr(typ *abi.Type, ptr unsafe.Pointer) {
	if writeBarrier.needed && typ.PtrBytes != 0 {
		bulkBarrierPreWrite(uintptr(ptr), 0, typ.PtrBytes)
	}

	mem.Clear(ptr, typ.Size_)
}

// See: ${GOROOT}/src/runtime/mbarrier.go#func:typedslicecopy
func typedslicecopy(typ *abi.Type, dstPtr unsafe.Pointer, dstLen int, srcPtr unsafe.Pointer, srcLen int) int {
	n := dstLen
	if n > srcLen {
		n = srcLen
	}
	if n == 0 {
		return 0
	}

	// The compiler emits calls to typedslicecopy before
	// instrumentation runs, so unlike the other copying and
	// assignment operations, it's not instrumented in the calling
	// code and needs its own instrumentation.
	if race.Enabled {
		callerpc := getcallerpc()
		pc := abi.FuncPCABIInternal(slicecopy)
		race.WriteRangePC(dstPtr, uintptr(n)*typ.Size_, callerpc, pc)
		race.ReadRangePC(srcPtr, uintptr(n)*typ.Size_, callerpc, pc)
	}
	if msan.Enabled {
		msan.Write(dstPtr, uintptr(n)*typ.Size_)
		msan.Read(srcPtr, uintptr(n)*typ.Size_)
	}
	if asan.Enabled {
		asan.Write(dstPtr, uintptr(n)*typ.Size_)
		asan.Read(srcPtr, uintptr(n)*typ.Size_)
	}

	if dstPtr == srcPtr {
		return n
	}

	// Note: No point in checking typ.PtrBytes here:
	// compiler only emits calls to typedslicecopy for types with pointers,
	// and growslice and reflect_typedslicecopy check for pointers
	// before calling typedslicecopy.
	size := uintptr(n) * typ.Size_
	if writeBarrier.needed {
		pwsize := size - typ.Size_ + typ.PtrBytes
		bulkBarrierPreWrite(uintptr(dstPtr), uintptr(srcPtr), pwsize)
	}
	// See typedmemmove for a discussion of the race between the
	// barrier and memmove.
	mem.Move(dstPtr, srcPtr, size)
	return n
}

//
// math
//

func mulUintptr(a, b uintptr) (uintptr, bool) {
	return math.MulUintptr(a, b)
}

func complex128div(num, den complex128) (quo complex128) {
	return stdcomplex.Complex128Div(num, den)
}

//
// 64bit emulation for 32bit platforms
//

func int64div(n int64, d int64) int64     { return emu64.Int64Div(n, d) }
func uint64div(n uint64, d uint64) uint64 { return emu64.Uint64Div(n, d) }
func int64mod(n int64, d int64) int64     { return emu64.Int64Div(n, d) }
func uint64mod(n uint64, d uint64) uint64 { return emu64.Uint64Mod(n, d) }
func float64toint64(n float64) uint64     { return emu64.Float64ToInt64(n) }
func float64touint64(n float64) uint64    { return emu64.Float64ToUint64(n) }
func float64touint32(n float64) uint32    { return emu64.Float64ToUint32(n) }
func int64tofloat64(n int64) float64      { return emu64.Int64ToFloat64(n) }
func int64tofloat32(n int64) float32      { return emu64.Int64ToFloat32(n) }
func uint64tofloat64(n uint64) float64    { return emu64.Uint64ToFloat64(n) }
func uint64tofloat32(n uint64) float32    { return emu64.Uint64ToFloat32(n) }
func uint32tofloat64(n uint32) float64    { return emu64.Uint32ToFloat64(n) }

//
// caller
//

// implemented as a compiler intrinsic on all platforms except riscv64
//
// see core/caller for riscv64 source code.
//
//go:linkname getcallerpc
//go:noescape
func getcallerpc() uintptr

// implemented as a compiler intrinsic on all platforms
//
//go:linkname getcallersp
//go:noescape
func getcallersp() uintptr

// getclosureptr returns the pointer to the current closure.
// getclosureptr can only be used in an assignment statement
// at the entry of a function. Moreover, go:nosplit directive
// must be specified at the declaration of caller function,
// so that the function prolog does not clobber the closure register.
// for example:
//
//	//go:nosplit
//	func f(arg1, arg2, arg3 int) {
//		dx := getclosureptr()
//	}
//
// The compiler rewrites calls to this function into instructions that fetch the
// pointer from a well-known register (DX on x86 architecture, etc.) directly.
//
//go:linkname getclosureptr
//go:noescape
func getclosureptr() uintptr

//
// equal
//

// implemented inside core/mem
func memequal_varlen(p, q unsafe.Pointer) bool
func memequal(p, q unsafe.Pointer, sz uintptr) bool { return mem.Equal(p, q, sz) }
func memequal0(p, q unsafe.Pointer) bool            { return true }
func memequal8(p, q unsafe.Pointer) bool            { return *(*int8)(p) == *(*int8)(q) }
func memequal16(p, q unsafe.Pointer) bool           { return *(*int16)(p) == *(*int16)(q) }
func memequal32(p, q unsafe.Pointer) bool           { return *(*int32)(p) == *(*int32)(q) }
func memequal64(p, q unsafe.Pointer) bool           { return *(*int64)(p) == *(*int64)(q) }
func memequal128(p, q unsafe.Pointer) bool          { return *(*[2]int64)(p) == *(*[2]int64)(q) }
func f32equal(p, q unsafe.Pointer) bool             { return *(*float32)(p) == *(*float32)(q) }
func f64equal(p, q unsafe.Pointer) bool             { return *(*float64)(p) == *(*float64)(q) }
func c64equal(p, q unsafe.Pointer) bool             { return *(*complex64)(p) == *(*complex64)(q) }
func c128equal(p, q unsafe.Pointer) bool            { return *(*complex128)(p) == *(*complex128)(q) }
func strequal(p, q unsafe.Pointer) bool             { return *(*string)(p) == *(*string)(q) }

func interequal(p, q unsafe.Pointer) bool {
	x := *(*stdtype.Iface)(p)
	y := *(*stdtype.Iface)(q)
	return x.Itab == y.Itab && ifaceeq(x.Itab, x.Data, y.Data)
}

func nilinterequal(p, q unsafe.Pointer) bool {
	x := *(*stdtype.Eface)(p)
	y := *(*stdtype.Eface)(q)
	return x.Type == y.Type && efaceeq(x.Type, x.Data, y.Data)
}

//
// hashing
//

func memhash(p unsafe.Pointer, h uintptr, sz uintptr) uintptr { return hash.MemHash(p, h, sz) }
func memhash0(p unsafe.Pointer, h uintptr) uintptr            { return h }
func memhash8(p unsafe.Pointer, h uintptr) uintptr            { return hash.MemHash(p, h, 1) }
func memhash16(p unsafe.Pointer, h uintptr) uintptr           { return hash.MemHash(p, h, 2) }
func memhash32(p unsafe.Pointer, h uintptr) uintptr           { return hash.MemHash32(p, h) }
func memhash64(p unsafe.Pointer, h uintptr) uintptr           { return hash.MemHash64(p, h) }
func memhash128(p unsafe.Pointer, h uintptr) uintptr          { return hash.MemHash(p, h, 16) }

// See ${GOROOT}/src/runtime/alg.go#func:memhash_varlen
//
//go:nosplit
func memhash_varlen(p unsafe.Pointer, h uintptr) uintptr {
	ptr := getclosureptr()
	size := *(*uintptr)(unsafe.Pointer(ptr + unsafe.Sizeof(h)))
	return memhash(p, h, size)
}

func f32hash(p unsafe.Pointer, h uintptr) uintptr      { return hash.Float32Hash(p, h) }
func f64hash(p unsafe.Pointer, h uintptr) uintptr      { return hash.Float64Hash(p, h) }
func c64hash(p unsafe.Pointer, h uintptr) uintptr      { return hash.Complex64Hash(p, h) }
func c128hash(p unsafe.Pointer, h uintptr) uintptr     { return hash.Complex128Hash(p, h) }
func strhash(a unsafe.Pointer, h uintptr) uintptr      { return hash.StringHash(a, h) }
func interhash(p unsafe.Pointer, h uintptr) uintptr    { return hash.InterfaceHash(p, h) }
func nilinterhash(p unsafe.Pointer, h uintptr) uintptr { return hash.NilInterfaceHash(p, h) }

//
// race detection
//

func racefuncenter(p uintptr)                        { race.FuncEnter(p) }
func racefuncexit()                                  { race.FuncExit() }
func raceread(p unsafe.Pointer)                      { race.Read(p) }
func racewrite(p unsafe.Pointer)                     { race.Write(p) }
func racereadrange(addr unsafe.Pointer, sz uintptr)  { race.ReadRange(addr, sz) }
func racewriterange(addr unsafe.Pointer, sz uintptr) { race.WriteRange(addr, sz) }

//
// memory sanitizer
//

func msanread(addr unsafe.Pointer, sz uintptr)  { msan.Read(addr, sz) }
func msanwrite(addr unsafe.Pointer, sz uintptr) { msan.Write(addr, sz) }
func msanmove(dst, src, sz uintptr)             { msan.Move(dst, src, sz) }

//
// address sanitizer
//

func asanread(addr unsafe.Pointer, size uintptr)  { asan.Read(addr, size) }
func asanwrite(addr unsafe.Pointer, size uintptr) { asan.Write(addr, size) }

//
// ptr
//

func checkptrAlignment(p unsafe.Pointer, elem *abi.Type, n uintptr) {
	stdptr.CheckAlignment(p, elem, n)
}

func checkptrArithmetic(p unsafe.Pointer, originals []unsafe.Pointer) {
	stdptr.CheckArithmetic(p, originals)
}

//
// fuzz
//

func libfuzzerTraceCmp1(arg0, arg1 uint8, fakepc int)       { fuzz.TraceCmp1(arg0, arg1, fakepc) }
func libfuzzerTraceCmp2(arg0, arg1 uint16, fakepc int)      { fuzz.TraceCmp2(arg0, arg1, fakepc) }
func libfuzzerTraceCmp4(arg0, arg1 uint32, fakepc int)      { fuzz.TraceCmp4(arg0, arg1, fakepc) }
func libfuzzerTraceCmp8(arg0, arg1 uint64, fakepc int)      { fuzz.TraceCmp8(arg0, arg1, fakepc) }
func libfuzzerTraceConstCmp1(arg0, arg1 uint8, fakepc int)  { fuzz.TraceConstCmp1(arg0, arg1, fakepc) }
func libfuzzerTraceConstCmp2(arg0, arg1 uint16, fakepc int) { fuzz.TraceConstCmp2(arg0, arg1, fakepc) }
func libfuzzerTraceConstCmp4(arg0, arg1 uint32, fakepc int) { fuzz.TraceConstCmp4(arg0, arg1, fakepc) }
func libfuzzerTraceConstCmp8(arg0, arg1 uint64, fakepc int) { fuzz.TraceConstCmp8(arg0, arg1, fakepc) }
func libfuzzerHookStrCmp(arg0, arg1 string, fakepc int)     { fuzz.HookStrCmp(arg0, arg1, fakepc) }
func libfuzzerHookEqualFold(arg0, arg1 string, fakepc int)  { fuzz.HookEqualFold(arg0, arg1, fakepc) }

//
// coverage
//

func addCovMeta(p unsafe.Pointer, dlen uint32, hash [16]byte, pkpath string, pkid int, cmode uint8, cgran uint8) uint32 {
	return cover.AddCovMeta(p, dlen, hash, pkpath, pkid, cmode, cgran)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package rt0

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/pcz/std/algo/rand/fastrand"
	stdgo "github.com/primecitizens/pcz/std/builtin/go"
	stdtype "github.com/primecitizens/pcz/std/builtin/type"
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/cpu"
	"github.com/primecitizens/pcz/std/core/stack"
	"github.com/primecitizens/pcz/std/core/thread"
)

var (
	g0       rt0G
	_g0iface stdgo.G = &g0

	// mainG points to the global G that is used by:
	//  - the main thread (goroutine in this std)
	//	- ffi callback
	//
	// It points to a global rt0G by default.
	//
	//go:linkname mainG
	mainG *stdgo.GHead = &g0.head

	//go:linkname runtime_inittasks runtime.runtime_inittasks
	runtime_inittasks []*abi.InitTask

	// for windows, it's actually []*uint16, see ../sys/sys_windows.go
	//
	//go:linkname args
	args []*byte

	stacksize uintptr
)

const (
	offsetsOK = unsafe.Offsetof(rt0G{}.head) == 0 &&
		unsafe.Offsetof(rt0G{}.head.Stack) == 0 &&
		unsafe.Offsetof(rt0G{}.head.Stack.Lo) == 0 &&
		unsafe.Offsetof(rt0G{}.head.Stack.Hi) == arch.PtrSize
)

// rt0G is the default G implementation used for early initialization of the
// program.
type rt0G struct {
	head  stdgo.GHead
	frand fastrand.Source
}

//go:nosplit
func (gp *rt0G) ID() uint64 { return 0 }

//go:nosplit
func (gp *rt0G) Rand32() uint32 { return gp.frand.Rand32() }

//go:nosplit
func (gp *rt0G) Rand64() uint64 { return gp.frand.Rand64() }

//go:nosplit
func (gp *rt0G) Status() uint32 { return stdgo.StatusGrunning }

//go:nosplit
func (gp *rt0G) DefaultAlloc() alloc.M { return malloc() }

//go:nosplit
func (gp *rt0G) PersistantAlloc() alloc.P { return palloc() }

// rt0 starts the program, this function should be considered as the mstart1
// in the official go runtime.
//
//go:nosplit
//go:noinline
func rt0(argc int32, argv **byte) {
	if !offsetsOK {
		assert.Throw("invalid", "offsets", "in", "rt0G")
	}

	if argc > 0 {
		args = unsafe.Slice(argv, argc)
	}

	// TODO: get actual size of the process stack
	//
	// - linux/darwin: check RLIMIT_STACK
	if arch.IsWasm == 1 {
		stacksize = 8192
	} else {
		stacksize = 4 * 1024 * 1024
	}

	mainG.Stack.Hi = stack.GetSP()
	mainG.Stack.Lo = mainG.Stack.Hi - stacksize
	mainG.Stackguard0 = mainG.Stack.Lo + stack.StackGuard
	mainG.Stackguard1 = mainG.Stackguard0

	// NOTE: code next to this comment is critical for dwarf generation
	// 		 it preserves relocation symbol `type:uintptr` for symbol `go:info.uintptr`
	//
	// It prevents the go tool link from throwing error:
	//
	// 	go:info.uintptr: unreachable sym in relocation: type:uintptr
	//
	// go linker doesn't mark `type:uintptr` reachable during deadcode analysis if it's
	// not converted to an interface type, and that's the case for this runtime
	mainG.Itab = stdtype.IfaceOf(&_g0iface).Itab

	thread.SetG(mainG)

	if !preinit() {
		cpu.Initialize("")

		for _, tsk := range runtime_inittasks {
			doInit(tsk)
		}

		for i, iter := 0, abi.ModuleIter(0); ; i++ {
			md, ok := iter.Nth(i)
			if !ok {
				break
			}

			for _, tsk := range md.InitTasks {
				doInit(tsk)
			}
		}
	}

	// verify g works
	if thread.G().G().Status() != stdgo.StatusGrunning {
		assert.Throw("goroutine", "not", "working", "properly")
	}

	main()

	exit0()
}

// `main.main` is the symbol of the main function from the main package
// created by the linker.
//
//go:linkname main main.main
func main()

// TODO(toolchain): create symbol main.preinit for tasks before doInit
func preinit() (skip bool) {
	return false
}

// doInit runs all init() functions in the tsk.
//
//go:nosplit
func doInit(tsk *abi.InitTask) {
	switch tsk.State {
	case 2: // fully initialized
		return
	case 1: // initialization in progress
		assert.Throw("recursive", "call", "during", "initialization", "-", "linker", "skew")
	}

	// not initialized yet

	tsk.State = 1 // mark in progress

	if tsk.NFns == 0 {
		// go toolchain should have pruned all of these in the linker.
		assert.Throw("inittask", "with", "no", "functions")
	}

	for i := uint32(0); i < tsk.NFns; i++ {
		tsk.Nth(i)()
	}

	tsk.State = 2 // mark done
}

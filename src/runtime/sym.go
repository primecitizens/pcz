// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"unsafe"

	stdgo "github.com/primecitizens/std/builtin/go"
	stdmap "github.com/primecitizens/std/builtin/map"
	stdtype "github.com/primecitizens/std/builtin/type"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"

	// _ "github.com/primecitizens/std/core/gc"    // import implementation of gcWritBarrier
	_ "github.com/primecitizens/std/core/stack" // import implementation of morestack and morestack_noctxt
	_ "github.com/primecitizens/std/rt0"        // import the magic line
)

var (
	// Set by the linker so the runtime can determine the buildmode.
	islibrary bool // -buildmode=c-shared
	isarchive bool // -buildmode=c-archive

	goarm uint8 // set by cmd/link on arm systems

	// base address for all 0-byte allocations
	zerobase uintptr

	lastmoduledatap *moduledata // linker symbol
	firstmoduledata moduledata  // linker symbol

	// This slice records the initializing tasks that need to be
	// done to start up the runtime. It is built by the linker.
	//
	// See $GOROOT/src/runtime/proc.go#var:runtime_inittasks
	// NOTE: It is used by rt0.start().
	runtime_inittasks []*abi.InitTask

	// The compiler knows about this variable.
	// If you change it, you must change builtin/runtime.go, too.
	// If you change the first four bytes, you must also change the write
	// barrier insertion code.
	writeBarrier struct {
		enabled bool    // compiler emits a check of this before calling write barrier
		pad     [3]byte // compiler uses 32-bit load for "enabled" field
		needed  bool    // whether we need a write barrier for current GC phase
		cgo     bool    // whether we need a write barrier for a cgo check
		alignme uint64  // guarantee alignment so that compiler can use a 32 or 64-bit load
	}

	// pinnedTypemaps are the map[typeOff]*_type from the moduledata objects.
	//
	// These typemap objects are allocated at run time on the heap, but the
	// only direct reference to them is in the moduledata, created by the
	// linker and marked SNOPTRDATA so it is ignored by the GC.
	//
	// To make sure the map isn't collected, we keep a second reference here.
	pinnedTypemaps []map[abi.TypeOff]*abi.Type

	// set using cmd/go/internal/modload.ModInfoProg
	modinfo string

	// buildVersion is the Go tree's version string at build time.
	//
	// This is set by the linker.
	//
	// This is accessed by "go version <binary>".
	buildVersion string

	defaultGOROOT string // set by cmd/link

	// cgoAlwaysFalse is a boolean value that is always false.
	// The cgo-generated code says if cgoAlwaysFalse { cgoUse(p) }.
	// The compiler cannot see that cgoAlwaysFalse is always false,
	// so it emits the test and keeps the call, giving the desired
	// escape analysis result. The test is cheaper than the call.
	cgoAlwaysFalse bool
)

// The linker redirects a reference of a method that it determined
// unreachable to a reference to this function, so it will throw if
// ever called.
//
//go:nosplit
func unreachableMethod() {
	assert.Unreachable()
}

// This is exported as ABI0 via linkname so obj can call it.
//
//go:nosplit
//go:linkname morestackc
func morestackc() {
	assert.Throw("attempt to execute system stack code on user stack")
}

type (
	_panic = stdgo.Panic
	_defer = stdgo.Defer
	hiter  = stdmap.HashIter
	g      = stdgo.GHead
)

// getg returns the pointer to the current g.
// The compiler rewrites calls to this function into instructions
// that fetch the g directly (from TLS or from the dedicated register).
func getg() *g

// function symbols required by linker (defined in core/stack/morestack.s)
func morestack()
func morestack_noctxt()

// Called from compiled code; declared for vet; do NOT call from Go.
//
// see $GOROOT/src/runtime/
func gcWriteBarrier1()
func gcWriteBarrier2()
func gcWriteBarrier3()
func gcWriteBarrier4()
func gcWriteBarrier5()
func gcWriteBarrier6()
func gcWriteBarrier7()
func gcWriteBarrier8()

// mapinitnoop is a no-op function known the Go linker; if a given global
// map (of the right size) is determined to be dead, the linker will
// rewrite the relocation (from the package init func) from the outlined
// map init function to this symbol.
//
// Defined in assembly so as to avoid complications with
// instrumentation (coverage, etc).
func mapinitnoop()

// type symbols required by linker
type moduledata abi.ModuleData

// type symbols required by dwarf genertion
//
// See src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type (
	eface stdtype.Eface
	iface stdtype.Iface

	// Variant with *byte pointer type for DWARF debugging.
	stringStructDWARF struct {
		str *byte
		len int
	}

	// here we keep a copy of stdslice.SliceHeader as we want `array` field instead fo `Array`
	slice struct {
		// NOTE: linker checks `array` field for dwarf
		// 		DO NOT CHANGE THESE NAMES
		array unsafe.Pointer
		len   int
		cap   int
	}

	hmap stdmap.HashMap
	bmap stdmap.MapBucket

	sudog stdgo.Sudog
	waitq stdgo.Waitq
	hchan stdgo.Chan

	itab abi.Itab

	// initTask abi.InitTask

	// The compiler knows that a print of a value of this type
	// should use printhex instead of printuint (decimal).
	hex uint64
)

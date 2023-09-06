// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import (
	"unsafe"

	"github.com/primecitizens/std/core/mark"
)

type FuncVal struct {
	Fn uintptr
	// variable-size, fn-specific data here
}

// Layout of in-memory per-function information prepared by linker
// See https://golang.org/s/go12symtab.
// Keep in sync with linker (../cmd/link/internal/ld/pcln.go:/pclntab)
// and with package debug/gosym and with symtab.go in package runtime.
//
// see $GOROOT/src/runtime/runtime2.go#type:_func
type Func struct {
	_ mark.NotInHeap // Only in static data

	entryOff uint32  // start pc, as offset from moduledata.text/pcHeader.textStart
	nameOff  NameOff // function name, as index into moduledata.funcnametab.

	args        int32  // in/out args size
	DeferReturn uint32 // offset of start of a deferreturn call instruction from entry, if any.

	PCSP      uint32
	PCFile    uint32
	PCLn      uint32
	npcdata   uint32
	cuOffset  uint32 // runtime.cutab offset of this function's CU
	StartLine int32  // line number of start of function (func keyword/TEXT directive)
	FuncID    FuncID // set for certain special runtime functions
	Flag      FuncFlag
	_         [1]byte // pad
	nfuncdata uint8   // must be last, must end on a uint32-aligned boundary

	// The end of the struct is followed immediately by two variable-length
	// arrays that reference the pcdata and funcdata locations for this
	// function.

	// pcdata contains the offset into moduledata.pctab for the start of
	// that index's table. e.g.,
	// &moduledata.pctab[_func.pcdata[_PCDATA_UnsafePoint]] is the start of
	// the unsafe point table.
	//
	// An offset of 0 indicates that there is no table.
	//
	// pcdata [npcdata]uint32

	// funcdata contains the offset past moduledata.gofunc which contains a
	// pointer to that index's funcdata. e.g.,
	// *(moduledata.gofunc +  _func.funcdata[_FUNCDATA_ArgsPointerMaps]) is
	// the argument pointer map.
	//
	// An offset of ^uint32(0) indicates that there is no entry.
	//
	// funcdata [nfuncdata]uint32
}

// IsInlined reports whether f should be re-interpreted as a *funcinl.
func (f *Func) IsInlined() bool {
	return f.entryOff == ^uint32(0) // see comment for funcinl.ones
}

// Pseudo-Func that is returned for PCs that occur in inlined code.
// A *Func can be either a *_func or a *FuncInl, and they are distinguished
// by the first uintptr.
//
// TODO(austin): Can we merge this with inlinedCall?
type FuncInl struct {
	ones      uint32  // set to ^0 to distinguish from _func
	entry     uintptr // entry of the real (the "outermost") frame
	name      string
	file      string
	line      int32
	startLine int32
}

type FuncInfo struct {
	*Func
	datap *ModuleData
}

func (f FuncInfo) Valid() bool {
	return f.Func != nil
}

// Name is something like "main.(*T).F".
func (f FuncInfo) Name() string {
	if !f.Valid() {
		return ""
	}

	return f.datap.funcName(f.nameOff)
}

// Entry returns the entry PC for f.
func (f FuncInfo) Entry() uintptr {
	return f.datap.textAddr(f.entryOff)
}

// findfuncbucket is an array of these structures.
// Each bucket represents 4096 bytes of the text segment.
// Each subbucket represents 256 bytes of the text segment.
// To find a function given a pc, locate the bucket and subbucket for
// that pc. Add together the idx and subbucket value to obtain a
// function index. Then scan the functab array starting at that
// index to find the target function.
// This table uses 20 bytes for every 4096 bytes of code, or ~0.5% overhead.
type findfuncbucket struct {
	idx        uint32
	subbuckets [16]byte
}

// FindFunc looks up function metadata for a PC.
//
// It is nosplit because it's part of the isgoexception
// implementation.
//
//go:nosplit
func FindFunc(pc uintptr) FuncInfo {
	datap := findModuleForPC(pc)
	if datap == nil {
		return FuncInfo{}
	}

	const nsub = uintptr(len(findfuncbucket{}.subbuckets))

	pcOff, ok := datap.textOff(pc)
	if !ok {
		return FuncInfo{}
	}

	x := uintptr(pcOff) + datap.Text - datap.MinPC // TODO: are datap.text and datap.minpc always equal?
	b := x / pcbucketsize
	i := x % pcbucketsize / (pcbucketsize / nsub)

	ffb := (*findfuncbucket)(unsafe.Add(unsafe.Pointer(datap.FindFuncTab), b*unsafe.Sizeof(findfuncbucket{})))
	idx := ffb.idx + uint32(ffb.subbuckets[i])

	// Find the ftab entry.
	for datap.FTab[idx+1].Entryoff <= pcOff {
		idx++
	}

	funcoff := datap.FTab[idx].Funcoff
	return FuncInfo{
		Func:  (*Func)(unsafe.Pointer(&datap.PCLnTable[funcoff])),
		datap: datap,
	}
}

func findModuleForPC(pc uintptr) (md *ModuleData) {
	for i, iter := 0, ModuleIter(0); ; i++ {
		md, ok := iter.Nth(i)
		if !ok {
			break
		}

		if md.MinPC <= pc && pc < md.MaxPC {
			return md
		}
	}

	return nil
}

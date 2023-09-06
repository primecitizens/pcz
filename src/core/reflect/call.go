// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/mem"
)

// Call calls fn with arguments described by stackArgs, stackArgsSize,
// frameSize, and regArgs.
//
// Arguments passed on the stack and space for return values passed on the stack
// must be laid out at the space pointed to by stackArgs (with total length
// stackArgsSize) according to the ABI.
//
// stackRetOffset must be some value <= stackArgsSize that indicates the
// offset within stackArgs where the return value space begins.
//
// frameSize is the total size of the argument frame at stackArgs and must
// therefore be >= stackArgsSize. It must include additional space for spilling
// register arguments for stack growth and preemption.
//
// TODO(mknyszek): Once we don't need the additional spill space, remove frameSize,
// since frameSize will be redundant with stackArgsSize.
//
// Arguments passed in registers must be laid out in regArgs according to the ABI.
// regArgs will hold any return values passed in registers after the call.
//
// Call copies stack arguments from stackArgs to the goroutine stack, and
// then copies back stackArgsSize-stackRetOffset bytes back to the return space
// in stackArgs once fn has completed. It also "unspills" argument registers from
// regArgs before calling fn, and spills them back into regArgs immediately
// following the call to fn. If there are results being returned on the stack,
// the caller should pass the argument frame type as stackArgsType so that
// reflectcall can execute appropriate write barriers during the copy.
//
// Call expects regArgs.ReturnIsPtr to be populated indicating which
// registers on the return path will contain Go pointers. It will then store
// these pointers in regArgs.Ptrs such that they are visible to the GC.
//
// Arguments passed through to reflectcall do not escape. The type is used
// only in a very limited callee of reflectcall, the stackArgs are copied, and
// regArgs is only used in the reflectcall frame.
//
//go:noescape
func Call(stackArgsType *abi.Type, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)

// in asm_*.s
// not called directly; definitions here supply type information for traceback.
// These must have the same signature (arg pointer map) as reflectcall.

func call16(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call32(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call64(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call128(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call256(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call512(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call1024(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call2048(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call4096(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call8192(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call16384(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call32768(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call65536(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call131072(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call262144(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call524288(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call1048576(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call2097152(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call4194304(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call8388608(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call16777216(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call33554432(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call67108864(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call134217728(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call268435456(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call536870912(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)
func call1073741824(typ, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs)

func badreflectcall() {
	assert.Throw("arg size to reflect.call more than 1GB")
}

// reflectcallmove is invoked by reflectcall to copy the return values
// out of the stack and into the heap, invoking the necessary write
// barriers. dst, src, and size describe the return value area to
// copy. typ describes the entire frame (not just the return values).
// typ may be nil, which indicates write barriers are not needed.
//
// It must be nosplit and must only call nosplit functions because the
// stack map of reflectcall is wrong.
//
//go:nosplit
func reflectcallmove(typ *abi.Type, dst, src unsafe.Pointer, size uintptr, regs *abi.RegArgs) {
	// if writeBarrier.needed && typ != nil && typ.ptrdata != 0 && size >= goarch.PtrSize {
	// 	bulkBarrierPreWrite(uintptr(dst), uintptr(src), size)
	// }

	mem.Memmove(dst, src, size)

	// Move pointers returned in registers to a place where the GC can see them.
	for i := range regs.Ints {
		if regs.ReturnIsPtr.Get(i) {
			regs.Ptrs[i] = unsafe.Pointer(regs.Ints[i])
		}
	}
}

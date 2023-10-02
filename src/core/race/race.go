// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package race

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
)

// TODO

// runtime.raceinit
func Init() (uintptr, uintptr) {
	if Enabled {
		assert.TODO()
	}

	return 0, 0
}

// runtime.racefini
func Fini() {
	if Enabled {
		assert.TODO()
	}
}

// runtime.raceproccreate
func ProcCreate() uintptr {
	if Enabled {
		assert.TODO()
	}

	return 0
}

// runtime.raceprocdestroy
func ProcDestroy(ctx uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racefuncenter
func FuncEnter(pc uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racefuncenterfp
func FuncEnterFP(fp uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racefuncexit
func FuncExit() {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racemapshadow
func MapShadow(addr unsafe.Pointer, size uintptr) {
	if Enabled {
		assert.TODO()
	}
}
func Read(p unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racereadrange
func ReadRange(addr unsafe.Pointer, size uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racereadpc
func ReadPC(addr unsafe.Pointer, callpc, pc uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racereadrangepc
func ReadRangePC(addr unsafe.Pointer, sz, callerpc, pc uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.raceReadObjectPC
//
// For all functions accepting callerpc and pc,
// callerpc is a return PC of the function that calls this function,
// pc is start PC of the function that calls this function.
func ReadObjectPC(t *abi.Type, addr unsafe.Pointer, callerpc, pc uintptr) {
	switch t.Kind() {
	case abi.KindArray, abi.KindStruct:
		// for composite objects we have to read every address
		// because a write might happen to any subobject.
		ReadRangePC(addr, t.Size_, callerpc, pc)
	default:
		// for non-composite objects we can read just the start
		// address, as any write must write the first byte.
		ReadPC(addr, callerpc, pc)
	}
}

// runtime.racewrite
func Write(p unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racewriterange
func WriteRange(addr unsafe.Pointer, size uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racewritepc
func WritePC(addr unsafe.Pointer, callpc, pc uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racewriterangepc
func WriteRangePC(addr unsafe.Pointer, sz, callerpc, pc uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.raceWriteObjectPC
func WriteObjectPC(t *abi.Type, addr unsafe.Pointer, callerpc, pc uintptr) {
	switch t.Kind() {
	case abi.KindArray, abi.KindStruct:
		// for composite objects we have to write every address
		// because a write might happen to any subobject.
		WriteRangePC(addr, t.Size_, callerpc, pc)
	default:
		// for non-composite objects we can write just the start
		// address, as any write must write the first byte.
		WritePC(addr, callerpc, pc)
	}
}

// runtime.raceacquire
func Acquire(addr unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.raceacquirectx
func AcquireCtx(racectx uintptr, addr unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racerelease
func Release(addr unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racereleaseacquire
func ReleaseAcquire(addr unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.releasemerge
func ReleaseMerge(addr unsafe.Pointer) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racemalloc
func Malloc(p unsafe.Pointer, sz uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racefree
func Free(p unsafe.Pointer, sz uintptr) {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racefingo
func Fingo() {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racegostart
func Gostart(pc uintptr) uintptr {
	if Enabled {
		assert.TODO()
	}
	return 0
}

// runtime.racegoend
func Goend() {
	if Enabled {
		assert.TODO()
	}
}

// runtime.racectxend
func Ctxend(racectx uintptr) {
	if Enabled {
		assert.TODO()
	}
}

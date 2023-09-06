// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alloc

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/mem"
	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/core/zr"
)

func MallocGC(size uintptr, typ *abi.Type, needszero bool) unsafe.Pointer {
	assert.TODO()
	return nil
}

func Malloc(alignment, sz uintptr, zeroize bool) (ret unsafe.Pointer) {
	p, err := malloc(alignment, sz)
	if err != 0 || p == 0 {
		return nil
	}

	ret = unsafe.Pointer(p)
	if zeroize {
		mem.Memclr(ret, sz)
	}

	return
}

// New is the go builtin `new` equivalent
func New[T any]() (ret *T) {
	if useGoNativeAlloc {
		return new(T)
	}

	var x T
	return (*T)(Malloc(unsafe.Alignof(x), unsafe.Sizeof(x), true))
}

// Make is the go builtin `make([]T, sz, kap)` equivalent
func Make[T any, Size, Cap num.Integer](s Size, k Cap) (ret []T) {
	if useGoNativeAlloc {
		return make([]T, s, k)
	}

	if k == 0 {
		return unsafe.Slice(zr.Pointer[T](), 0)
	}

	var elem T
	ret = unsafe.Slice(
		(*T)(Malloc(unsafe.Alignof(elem), uintptr(k)*unsafe.Sizeof(elem), true)),
		k,
	)
	return ret[:s:k]
}

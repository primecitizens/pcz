package libcalloc

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/assert"
)

type T struct{}

func (T) Malloc(typ *abi.Type, n uintptr, zeroize bool) unsafe.Pointer {
	if n == 0 || (typ != nil && typ.Size_ == 0) {
		return alloc.ZeroSized()
	}

	assert.TODO()
	return nil
}

func (T) Palloc(size uintptr) unsafe.Pointer {
	if size == 0 {
		return alloc.ZeroSized()
	}

	assert.TODO()
	return nil
}

func (T) Free(typ *abi.Type, n uintptr, ptr unsafe.Pointer) alloc.Hint {
	assert.TODO()
	return 0
}

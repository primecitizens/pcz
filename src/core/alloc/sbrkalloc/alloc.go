package sbrkalloc

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/num"
	"github.com/primecitizens/pcz/std/core/os"
)

// TODO: need mutex protection once the std supports multi-threading.

// initialized by platform specific init()
var (
	_start    uintptr // _start is the start of heap memory
	base      uintptr // base is the base address for allocating new memory chunks
	end       uintptr // end is the end of heap memory
	lastFree  *head   // pointer to the last freed memory chunk
	lastAlloc *head   // pointer to the last allocated memory chunk, for adjusting size for alignment
)

const (
	headSz = unsafe.Sizeof(head{})
)

// A head is an allocation record put at the start of allocated memory
type head struct {
	nextfree *head
	// total sz of the allocated node (including this node)
	sz uintptr
}

func findVacant(sz uintptr) unsafe.Pointer {
	if lastFree == nil {
		return nil
	}

	cur := lastFree
	var fit, fitPrev, last *head
	for ; cur != nil; cur, last = cur.nextfree, cur {
		// TODO: merge neighbor chunks
		if cur.sz < sz {
			continue
		}

		if fit == nil || fit.sz > cur.sz {
			fit, fitPrev = cur, last
		}

		if cur.sz == sz {
			break
		}
	}

	if fit == nil {
		return nil
	}

	if fit.sz-sz > headSz {
		// can host at least one more record in this chunk
		frag := (*head)(unsafe.Add(unsafe.Pointer(fit), sz))
		*frag = head{
			nextfree: fit.nextfree,
			sz:       fit.sz - sz,
		}
		fit.sz = sz

		if fitPrev == nil {
			lastFree = frag
		} else {
			fitPrev.nextfree = frag
		}
	} else {
		if fitPrev == nil {
			lastFree = fit.nextfree
		} else {
			fitPrev.nextfree = fit.nextfree
		}
	}

	fit.nextfree = nil
	return unsafe.Pointer(uintptr(unsafe.Pointer(fit)) + headSz)
}

// T is the type implementating sbrk allocation.
type T struct{}

func (*T) Malloc(typ *abi.Type, n uintptr, zeroize bool) unsafe.Pointer {
	if n == 0 || (typ != nil && typ.Size_ == 0) {
		return alloc.ZeroSized()
	}

	var (
		szReq uintptr // total bytes requested by the caller
		sz    uintptr // total bytes including the head and szReq
	)
	if typ == nil {
		szReq = n
		sz = num.AlignUp(n+headSz, arch.Int64Align)
	} else {
		szReq = n * typ.Size_
		sz = num.AlignUp(n*typ.Size_+headSz, uintptr(max(arch.Int64Align, typ.Align_)))
	}

	if sz < szReq || sz > os.MaxAlloc {
		assert.Throw("size", "too", "large")
	}

	if luck := findVacant(sz); luck != nil {
		return luck
	}

	if end-base < sz {
		end = uintptr(Sbrk(num.AlignUp(sz, arch.DefaultPhysPageSize)))
	}

	lastAlloc = (*head)(unsafe.Pointer(base))
	*lastAlloc = head{
		nextfree: nil,
		sz:       sz,
	}

	ptr := unsafe.Pointer(base + headSz)
	base += sz

	if zeroize {
		clear(unsafe.Slice((*byte)(ptr), szReq))
	}

	return ptr
}

func (x *T) Palloc(size uintptr) unsafe.Pointer {
	if size == 0 {
		return alloc.ZeroSized()
	}

	return x.Malloc(nil, size, true)
}

// TODO: optimize for freeing last allocation

func (*T) Free(typ *abi.Type, n uintptr, ptr unsafe.Pointer) alloc.Hint {
	if n == 0 || (typ != nil && typ.Size_ == 0) {
		return 0
	}

	addr := uintptr(ptr) - headSz
	if addr < _start || addr > (base-headSz) {
		assert.Throw("memory", "not", "managed", "by", "this", "allocator")
	}

	x := (*head)(unsafe.Pointer(addr))
	if x.nextfree != nil {
		assert.Throw("double", "free")
	}

	if x.sz == 0 || x.sz > os.MaxAlloc {
		assert.Throw("invalid", "allocation", "record")
	}

	x.nextfree = lastFree
	lastFree = x
	return 0
}

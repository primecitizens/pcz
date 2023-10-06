package mem

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
)

// TypedMove copies a value of type typ to dst from src.
// Must be nosplit, see #16026.
//
// TODO: Perfect for go:nosplitrec since we can't have a safe point
// anywhere in the bulk barrier or Move.
//
// See ${GOROOT}/src/runtime/mbarrier.go#func:typedmemmove
//
//go:nosplit
func TypedMove(typ *abi.Type, dst, src unsafe.Pointer) {
	if dst == src {
		return
	}

	// TODO(pcz): implement typed move for self pointing fields
	var _ mark.SelfPointing

	// There's a race here: if some other goroutine can write to
	// src, it may change some pointer in src after we've
	// performed the write barrier but before we perform the
	// memory copy. This safe because the write performed by that
	// other goroutine must also be accompanied by a write
	// barrier, so at worst we've unnecessarily greyed the old
	// pointer that was in src.
	typedmemmove(typ, dst, src)
}

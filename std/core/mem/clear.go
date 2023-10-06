// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package mem

import (
	"unsafe"
)

// Clear clears n bytes starting at ptr.
//
// Usually you should use TypedClear. Clear should be
// used only when the caller knows that *ptr contains no heap pointers
// because either:
//
// *ptr is initialized memory and its type is pointer-free, or
//
// *ptr is uninitialized memory (e.g., memory that's being reused
// for a new allocation) and hence contains only "junk".
//
// Clear ensures that if ptr is pointer-aligned, and n
// is a multiple of the pointer size, then any pointer-aligned,
// pointer-sized portion is cleared atomically. Despite the function
// name, this is necessary because this function is the underlying
// implementation of TypedClear and memclrHasPointers. See the doc of
// memmove for more details.
//
// The (CPU-specific) implementations of this function are in memclr_*.s.
//
// See ${GOROOT}/src/runtime/stubs.go#func:memclrNoHeapPointers
//
//go:noescape
func Clear(ptr unsafe.Pointer, n uintptr)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/asan"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/cmp"
	"github.com/primecitizens/pcz/std/core/mem"
	"github.com/primecitizens/pcz/std/core/msan"
	"github.com/primecitizens/pcz/std/core/os"
	"github.com/primecitizens/pcz/std/core/race"
	"github.com/primecitizens/pcz/std/text/unicode/common"
	"github.com/primecitizens/pcz/std/text/unicode/utf8"
)

// The constant is known to the compiler.
// There is no fundamental theory behind this number.
const tmpStringBufSize = 32

type tmpBuf [tmpStringBufSize]byte

//
// string operations
//

func concatstring2(buf *tmpBuf, a0, a1 string) string {
	return concatstrings(buf, []string{a0, a1})
}

func concatstring3(buf *tmpBuf, a0, a1, a2 string) string {
	return concatstrings(buf, []string{a0, a1, a2})
}

func concatstring4(buf *tmpBuf, a0, a1, a2, a3 string) string {
	return concatstrings(buf, []string{a0, a1, a2, a3})
}

func concatstring5(buf *tmpBuf, a0, a1, a2, a3, a4 string) string {
	return concatstrings(buf, []string{a0, a1, a2, a3, a4})
}

func concatstrings(buf *tmpBuf, a []string) string {
	idx := 0
	l := 0
	count := 0
	for i, x := range a {
		n := len(x)
		if n == 0 {
			continue
		}
		if l+n < l {
			assert.Throw("string", "concatenation", "too", "long")
		}
		l += n
		count++
		idx = i
	}
	if count == 0 {
		return ""
	}

	// If there is just one string and either it is not on the stack
	// or our result does not escape the calling frame (buf != nil),
	// then we can return that string directly.
	if count == 1 && (buf != nil || !stringDataOnStack(a[idx])) {
		return a[idx]
	}
	s, b := rawstringtmp(buf, l)
	for _, x := range a {
		copy(b, x)
		b = b[len(x):]
	}
	return s
}

func cmpstring(a0, a1 string) int { return cmp.String(a0, a1) }

func intstring(buf *[4]byte, v int64) (s string) {
	var b []byte
	if buf != nil {
		b = buf[:]
		s = slicebytetostringtmp(&b[0], len(b))
	} else {
		s, b = rawstring(4)
	}
	if int64(rune(v)) != v {
		v = common.RuneError
	}

	_, n := utf8.EncodeRune(b[:0], rune(v))
	return s[:n]
}

// slicebytetostring converts a byte slice to a string.
// It is inserted by the compiler into generated code.
// ptr is a pointer to the first element of the slice;
// n is the length of the slice.
// Buf is a fixed-size buffer for the result,
// it is not nil if the result does not escape.
func slicebytetostring(buf *tmpBuf, ptr *byte, n int) string {
	if n == 0 {
		// Turns out to be a relatively common case.
		// Consider that you want to parse out data between parens in "foo()bar",
		// you find the indices and convert the subslice to string.
		return ""
	}

	if race.Enabled {
		race.ReadRangePC(unsafe.Pointer(ptr),
			uintptr(n),
			getcallerpc(),
			abi.FuncPCABIInternal(slicebytetostring))
	}
	if msan.Enabled {
		msan.Read(unsafe.Pointer(ptr), uintptr(n))
	}
	if asan.Enabled {
		asan.Read(unsafe.Pointer(ptr), uintptr(n))
	}

	if n == 1 {
		p := unsafe.Pointer(&staticuint64s[*ptr])
		if arch.BigEndian {
			p = unsafe.Add(p, 7)
		}
		return unsafe.String((*byte)(p), 1)
	}

	var p unsafe.Pointer
	if buf != nil && n <= len(buf) {
		p = unsafe.Pointer(buf)
	} else {
		p = explicit_mallocgc(nil, uintptr(n), false)
	}

	mem.Move(p, unsafe.Pointer(ptr), uintptr(n))
	return unsafe.String((*byte)(p), n)
}

// slicebytetostringtmp returns a "string" referring to the actual []byte bytes.
//
// Callers need to ensure that the returned string will not be used after
// the calling goroutine modifies the original slice or synchronizes with
// another goroutine.
//
// The function is only called when instrumenting
// and otherwise intrinsified by the compiler.
//
// Some internal compiler optimizations use this function.
//   - Used for m[T1{... Tn{..., string(k), ...} ...}] and m[string(k)]
//     where k is []byte, T1 to Tn is a nesting of struct and array literals.
//   - Used for "<"+string(b)+">" concatenation where b is []byte.
//   - Used for string(b)=="foo" comparison where b is []byte.
//
// NOTE: this function is implemented as a compiler intrinsic when not instrumenting.
// See $GOROOT/src/cmd/compile/internal/ssagen/ssa.go#func:InitTables
func slicebytetostringtmp(ptr *byte, n int) string {
	if race.Enabled && n > 0 {
		race.ReadRangePC(
			unsafe.Pointer(ptr),
			uintptr(n),
			getcallerpc(),
			abi.FuncPCABIInternal(slicebytetostringtmp),
		)
	}
	if msan.Enabled && n > 0 {
		msanread(unsafe.Pointer(ptr), uintptr(n))
	}
	if asan.Enabled && n > 0 {
		asanread(unsafe.Pointer(ptr), uintptr(n))
	}

	return unsafe.String(ptr, n)
}

func slicerunetostring(buf *tmpBuf, a []rune) string {
	if race.Enabled && len(a) > 0 {
		race.ReadRangePC(
			unsafe.Pointer(&a[0]),
			uintptr(len(a))*unsafe.Sizeof(a[0]),
			getcallerpc(),
			abi.FuncPCABIInternal(slicerunetostring),
		)
	}
	if msan.Enabled && len(a) > 0 {
		msanread(unsafe.Pointer(&a[0]), uintptr(len(a))*unsafe.Sizeof(a[0]))
	}
	if asan.Enabled && len(a) > 0 {
		asanread(unsafe.Pointer(&a[0]), uintptr(len(a))*unsafe.Sizeof(a[0]))
	}

	var (
		size1, size2, n int

		s, b = rawstringtmp(buf, size1+3)
	)
	for _, r := range a {
		size1 += utf8.RuneLen(r)
	}

	for _, r := range a {
		// check for race
		if size2 >= size1 {
			break
		}
		b, n = utf8.EncodeRune(b, r)
		size2 += n
	}

	return s[:size2]
}

func stringtoslicebyte(buf *tmpBuf, s string) []byte {
	var b []byte
	if buf != nil && len(s) <= len(buf) {
		*buf = tmpBuf{}
		b = buf[:len(s)]
	} else {
		b = rawbyteslice(len(s))
	}
	copy(b, s)
	return b
}

func stringtoslicerune(buf *[32]rune, s string) []rune {
	// two passes.
	// unlike slicerunetostring, no race because strings are immutable.
	n := 0
	for range s {
		n++
	}

	var a []rune
	if buf != nil && n <= len(buf) {
		*buf = [tmpStringBufSize]rune{}
		a = buf[:n]
	} else {
		a = rawruneslice(n)
	}

	n = 0
	for _, r := range s {
		a[n] = r
		n++
	}
	return a
}

func decoderune(s string, k int) (rune, int) { return utf8.IterRune(s, k) }
func countrunes(s string) int                { return utf8.Count(s) }

func rawstringtmp(buf *tmpBuf, l int) (s string, b []byte) {
	if buf != nil && l <= len(buf) {
		b = buf[:l]
		s = slicebytetostringtmp(&b[0], len(b))
	} else {
		s, b = rawstring(l)
	}
	return
}

// stringDataOnStack reports whether the string's data is
// stored on the current goroutine's stack.
func stringDataOnStack(s string) bool {
	return getg().Stack.PointerOnStack(uintptr(unsafe.Pointer(unsafe.StringData(s))))
}

// rawruneslice allocates a new rune slice. The rune slice is not zeroed.
func rawruneslice(size int) (b []rune) {
	if uintptr(size) > os.MaxAlloc/4 {
		assert.Throw("out", "of", "memory")
	}
	sz := uintptr(size) * 4
	p := explicit_mallocgc(nil, sz, false)
	if sz != uintptr(size)*4 {
		mem.Clear(unsafe.Add(p, uintptr(size)*4), sz-uintptr(size)*4)
	}

	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, int(sz / 4)}
	return
}

// rawbyteslice allocates a new byte slice. The byte slice is not zeroed.
func rawbyteslice(size int) (b []byte) {
	p := explicit_mallocgc(nil, uintptr(size), false)

	*(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}
	return
}

// rawstring allocates storage for a new string. The returned
// string and byte slice both refer to the same storage.
// The storage is not zeroed. Callers should use
// b to set the string contents and then drop b.
func rawstring(size int) (s string, b []byte) {
	p := explicit_mallocgc(nil, uintptr(size), false)
	return unsafe.String((*byte)(p), size), unsafe.Slice((*byte)(p), size)
}

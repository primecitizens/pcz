// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && windows

package sys

import (
	"unsafe"

	stdconst "github.com/primecitizens/pcz/std/builtin/const"
	stdptr "github.com/primecitizens/pcz/std/builtin/ptr"
	stdslice "github.com/primecitizens/pcz/std/builtin/slice"
	stdstring "github.com/primecitizens/pcz/std/builtin/string"
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/text/unicode/wtf16"
)

func init() {
	base := unsafe.SliceData(_args)
	if base == nil { // no args
		return
	}

	args = wtf16Decode(_args)

	// base is the envv
	base = stdptr.Add(base, arch.PtrSize*uintptr(len(_args)+1 /* NULL sep */))
	n := 0
	for p := base; *p != nil; p = stdptr.Add(p, arch.PtrSize) {
		n++
	}

	envs = wtf16Decode(unsafe.Slice(base, n))
}

func wtf16Decode(slice []*uint16) []string {
	strsliceBytes := stdconst.SizeStringType * uintptr(len(slice))
	totalBytes := strsliceBytes
	for _, p := range slice {
		x := unsafe.Slice(p, stdstring.FindNull2(p))
		wtf8bytes, canDecodeInPlace := wtf16.WTF8DecodedSize(x...)

		if !canDecodeInPlace {
			totalBytes += uintptr(wtf8bytes)
		}
	}

	buf := make([]byte, totalBytes)
	ret := unsafe.Slice((*string)(unsafe.Pointer(unsafe.SliceData(buf))), len(slice))

	buf = unsafe.Slice((*byte)(unsafe.SliceData(buf[strsliceBytes:])), totalBytes-strsliceBytes)[:0]
	for i, p := range slice {
		x := unsafe.Slice(p, stdstring.FindNull2(p))
		wtf8bytes, canDecodeInPlace := wtf16.WTF8DecodedSize(x...)

		var decoded []byte
		if canDecodeInPlace {
			decoded, _ = wtf16.WTF8Decode(unsafe.Slice((*byte)(unsafe.Pointer(p)), wtf8bytes)[:0], x...)
		} else {
			decoded, _ = wtf16.WTF8Decode(buf, x...)
			decoded, buf = stdslice.Cut(decoded)
		}

		if len(decoded) != wtf8bytes {
			assert.Throw("bad", "wtf16", "decoding")
		}

		ret[i] = stdstring.FromBytes(decoded)
	}

	return ret
}

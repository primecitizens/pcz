// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

// NewString creates a js string for the string, return the reference to
// the created string.
func NewString(s string) String {
	return String{
		ref: bindings.DecodeUTF8(
			StringData(s),
			SizeU(len(s)),
		),
	}
}

// A String is a reference to a javascript string.
type String struct {
	ref bindings.Ref
}

func (s String) FromRef(ref Ref) String {
	return String{
		ref: bindings.Ref(ref),
	}
}

func (s String) Ref() Ref {
	return Ref(s.ref)
}

func (s String) Once() String {
	bindings.Once(s.ref)
	return s
}

func (s String) Free() {
	bindings.Free(s.ref)
}

func (s String) Equals(other string) bool {
	return bindings.Ref(True) == bindings.EqualsUTF8(
		s.ref,
		StringData(other),
		SizeU(len(other)),
	)
}

func (s String) Prepend(x string) String {
	bindings.PrependUTF8(
		s.ref,
		StringData(x),
		SizeU(len(x)),
	)
	return s
}

func (s String) Append(x string) String {
	bindings.AppendUTF8(
		s.ref,
		StringData(x),
		SizeU(len(x)),
	)
	return s
}

// Read encodes the js string to buf (UTF-8 encoding) and
// returns the count of bytes encoded.
func (s String) Read(buf []byte) int {
	n := bindings.EncodeUTF8(
		s.ref,
		SliceData(buf),
		SizeU(len(buf)),
	)
	return int(n)
}

// Len returns the utf-8 length of the js string.
func (s String) Len() int {
	return int(bindings.SizeUTF8(s.ref))
}

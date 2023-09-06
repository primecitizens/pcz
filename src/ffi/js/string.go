// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/ffi/js/bindings"
)

// NewString creates a js string for the string, return the reference to
// the created string.
func NewString(s string) String {
	return String{
		bindings.StringFromUTF8(
			StringData(s),
			SizeU(len(s)),
		),
	}
}

// A String is a reference to a javascript string.
type String struct {
	ref bindings.Ref
}

func (s String) Ref() Ref {
	return Ref(s.ref)
}

func (s String) Equals(other string) bool {
	return bindings.StringEqualsUTF8(
		s.ref,
		StringData(other),
		SizeU(len(other)),
	) == bindings.Ref(True)
}

func (s String) Into(buf []byte) int {
	n := bindings.StringIntoUTF8(
		s.ref,
		SliceData(buf),
		SizeU(len(buf)),
	)
	return int(n)
}

// Len returns the utf-8 length of the js string.
func (s String) Len() int {
	return int(bindings.StringSizeUTF8(s.ref))
}

func (s String) Free() {
	bindings.Free(s.ref)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdslice

import (
	"unsafe"
)

type Header struct {
	Array unsafe.Pointer
	Len   int
	Cap   int
}

// Split returns actual members of a slice.
func Split[Elem any](s []Elem) (arr *Elem, sz, kap int) {
	return unsafe.SliceData(s), len(s), cap(s)
}

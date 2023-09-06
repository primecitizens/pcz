// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/ffi/js/bindings"
)

const smallIntCacheSize = 128

const (
	Undefined          Ref = iota
	Null                   //
	True                   //
	False                  //
	Global                 //
	firstSmallIntCache     //
	lastSmallIntCache  = firstSmallIntCache + smallIntCacheSize
)

type Ref bindings.Ref

func (r Ref) AsFunc() Func {
	return Func{ref: bindings.Ref(r)}
}

func (r Ref) AsArray() Array {
	return Array{ref: bindings.Ref(r)}
}

func (r Ref) AsString() String {
	return String{ref: bindings.Ref(r)}
}

func (r Ref) AsObject() Object {
	return Object{ref: bindings.Ref(r)}
}

func (r Ref) InstanceOf(o Ref) bool {
	return bindings.Instanceof(bindings.Ref(r), bindings.Ref(o)) == bindings.Ref(True)
}

// Free the referenced js value.
func (r Ref) Free() {
	bindings.Free(bindings.Ref(r))
}

func Free[T ~uint32](refs ...T) {
	bindings.BatchFree(
		SliceData(refs),
		SizeU(len(refs)),
	)
}

type Any struct{}

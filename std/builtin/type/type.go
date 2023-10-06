// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdtype

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
)

// Eface is the definition for interfaces with no method.
//
//   - interface{}
//   - any
//   - type Foo interface{}
type Eface struct {
	Type *abi.Type
	Data unsafe.Pointer
}

func TypeOf(typ any) *abi.Type {
	return mark.NoEscape(EfaceOf(mark.NoEscape(&typ)).Type)
}

func EfaceOf(ep *any) *Eface {
	return (*Eface)(unsafe.Pointer(ep))
}

// Iface is the definition for interfaces with methods.
type Iface struct {
	Itab *abi.Itab
	Data unsafe.Pointer
}

// NOTE: T MUST be of an interface type.
func IfaceOf[T any](ip *T) *Iface {
	return (*Iface)(unsafe.Pointer(ip))
}

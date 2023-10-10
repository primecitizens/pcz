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

func Type[T any]() *abi.Type {
	var x any = (*T)(nil)
	return mark.NoEscape(
		EfaceOf(mark.NoEscape(&x)).Type.PointerTypeUnsafe().Elem,
	)
}

// TypeOf returns the runtime type of x.
func TypeOf(x any) *abi.Type {
	return mark.NoEscape(EfaceOf(mark.NoEscape(&x)).Type)
}

// ValueOf returns the data in x.
func ValueOf(x any) unsafe.Pointer {
	return EfaceOf(mark.NoEscape(&x)).Data
}

// EfaceOf returns the underlay Eface data of *ep.
func EfaceOf(ep *any) *Eface {
	return (*Eface)(mark.NoEscapePointer(ep))
}

func FromEface(ep *Eface) any {
	return *(*any)(mark.NoEscapePointer(ep))
}

func FormatEface(typ *abi.Type, data unsafe.Pointer) any {
	eface := Eface{
		Type: typ,
		Data: data,
	}

	return FromEface((*Eface)(mark.NoEscapePointer(&eface)))
}

// Iface is the definition for interfaces with at least one method.
type Iface struct {
	Itab *abi.Itab
	Data unsafe.Pointer
}

// IfaceOf returns the underlay Iface data of T.
//
// NOTE: T MUST be an interface type with at least one method.
func IfaceOf[T any](ip *T) *Iface {
	return (*Iface)(mark.NoEscapePointer(ip))
}

// NOTE: T MUST be an interface type with at least one method.
func FromIface[T any](ip *Iface) T {
	return *(*T)(mark.NoEscapePointer(ip))
}

// NOTE: T MUST be an interface type with at least one method.
func FormatIface[T any](itab *abi.Itab, data unsafe.Pointer) T {
	iface := Iface{
		Itab: itab,
		Data: data,
	}

	return FromIface[T]((*Iface)(mark.NoEscapePointer(&iface)))
}

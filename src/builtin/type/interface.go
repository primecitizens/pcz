// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdtype

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/asan"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/debug"
	"github.com/primecitizens/std/core/mem"
	"github.com/primecitizens/std/core/msan"
	"github.com/primecitizens/std/core/race"
	"github.com/primecitizens/std/core/zr"
)

func InterfaceEqual(p, q unsafe.Pointer) bool {
	x := *(*Iface)(p)
	y := *(*Iface)(q)
	return x.Itab == y.Itab && IfaceEq(x.Itab, x.Data, y.Data)
}

func NilInterfaceEqual(p, q unsafe.Pointer) bool {
	x := *(*Eface)(p)
	y := *(*Eface)(q)
	return x.Type == y.Type && EfaceEq(x.Type, x.Data, y.Data)
}

// Eface is the definition for interfaces with no method.
//
//   - interface{}
//   - any
//   - type Foo interface{}
type Eface struct {
	Type *abi.Type
	Data unsafe.Pointer
}

// Direct returns true if the value is directly stored as iface.Data.
func (f Eface) Direct() bool {
	return f.Type.Kind_&abi.KindDirectIface != 0
}

func EfaceOf(ep *any) *Eface { return (*Eface)(unsafe.Pointer(ep)) }

// Iface is the definition for interfaces with methods.
type Iface struct {
	Itab *abi.Itab
	Data unsafe.Pointer
}

// // Direct returns true if the value is directly stored as iface.Data.
// func (f Iface) Direct() bool {
// 	return f.Tab.Type.Kind_&abi.KindDirectIface != 0
// }

func getitab(inter *abi.InterfaceType, typ *abi.Type, canfail bool) (itab *abi.Itab) {
	itab, code := abi.GetItab(inter, typ)
	switch code {
	case abi.GetItabResult_OK:
		return itab
	case abi.GetItabResult_UncommonType:
		if canfail {
			return nil
		}
		name := inter.Type.NameOff(inter.Methods[0].Name)
		assert.Panic(TypeAssertionError{
			_interface:    nil,
			concrete:      typ,
			asserted:      &inter.Type,
			missingMethod: name.Name(),
		})
		return nil
	case abi.GetItabResult_:
		if canfail {
			return nil
		}
		assert.Panic(TypeAssertionError{
			concrete:      typ,
			asserted:      &inter.Type,
			missingMethod: itab.Init(),
		})
		return nil
	default:
		assert.Unreachable()
		return nil
	}
}

//
// Non-empty-interface to non-empty-interface conversion.
//

// convI2I returns the new itab to be used for the destination value
// when converting a value with itab src to the dst interface.
func ConvI2I(dst *abi.InterfaceType, src *abi.Itab) *abi.Itab {
	if src == nil {
		return nil
	}
	if src.Inter == dst {
		return src
	}

	return getitab(dst, src.Type, false)
}

// Convert non-interface type to the data word of a (empty or nonempty) interface.

// ConvT converts a value of type t, which is pointed to by v, to a pointer that can
// be used as the second word of an interface value.
func ConvT(t *abi.Type, v unsafe.Pointer) unsafe.Pointer {
	if race.Enabled {
		race.ReadObjectPC(t, v, debug.GetCallerPC(), abi.FuncPCABIInternal(ConvT))
	}
	if msan.Enabled {
		msan.Read(v, t.Size_)
	}
	if asan.Enabled {
		asan.Read(v, t.Size_)
	}

	if t.Size_ == 0 {
		return zr.ZeroSized()
	}

	x := mallocgc(t.Size_, t, true)
	mem.TypedMemmove(t, x, v)
	return x
}

// Same as convT, for types with no pointers in them.
func ConvTNoPtr(t *abi.Type, v unsafe.Pointer) unsafe.Pointer {
	// TODO: maybe take size instead of type?
	if race.Enabled {
		race.ReadObjectPC(t, v, debug.GetCallerPC(), abi.FuncPCABIInternal(ConvTNoPtr))
	}
	if msan.Enabled {
		msan.Read(v, t.Size_)
	}
	if asan.Enabled {
		asan.Read(v, t.Size_)
	}

	if t.Size_ == 0 {
		return zr.ZeroSized()
	}

	x := mallocgc(t.Size_, t, false)
	mem.Memmove(x, v, t.Size_)
	return x
}

// The specialized convTx routines need a type descriptor to use when calling mallocgc.
// We don't need the type to be exact, just to have the correct size, alignment, and pointer-ness.
// However, when debugging, it'd be nice to have some indication in mallocgc where the types came from,
// so we use named types here.
// We then construct interface values of these types,
// and then extract the type word to use as needed.
type (
	uint16InterfacePtr uint16
	uint32InterfacePtr uint32
	uint64InterfacePtr uint64
	stringInterfacePtr string
	sliceInterfacePtr  []byte
)

var (
	uint16Eface any = uint16InterfacePtr(0)
	uint32Eface any = uint32InterfacePtr(0)
	uint64Eface any = uint64InterfacePtr(0)
	stringEface any = stringInterfacePtr("")
	sliceEface  any = sliceInterfacePtr(nil)

	uint16Type *abi.Type = EfaceOf(&uint16Eface).Type
	uint32Type *abi.Type = EfaceOf(&uint32Eface).Type
	uint64Type *abi.Type = EfaceOf(&uint64Eface).Type
	stringType *abi.Type = EfaceOf(&stringEface).Type
	sliceType  *abi.Type = EfaceOf(&sliceEface).Type
)

// staticuint64s is used to avoid allocating in convTx for small integer values.
//
//go:linkname staticuint64s runtime.staticuint64s
var staticuint64s [256]uint64

// Specialized versions of convT for specific types.
// These functions take concrete types in the runtime. But they may
// be used for a wider range of types, which have the same memory
// layout as the parameter type. The compiler converts the
// to-be-converted type to the parameter type before calling the
// runtime function. This way, the call is ABI-insensitive.

func ConvT16(val uint16) (x unsafe.Pointer) {
	if val < uint16(len(staticuint64s)) {
		x = unsafe.Pointer(&staticuint64s[val])
		if arch.BigEndian {
			x = unsafe.Pointer(uintptr(x) + 6)
		}
	} else {
		x = mallocgc(2, uint16Type, false)
		*(*uint16)(x) = val
	}
	return
}

func ConvT32(val uint32) (x unsafe.Pointer) {
	if val < uint32(len(staticuint64s)) {
		x = unsafe.Pointer(&staticuint64s[val])
		if arch.BigEndian {
			x = unsafe.Pointer(uintptr(x) + 4)
		}
	} else {
		x = mallocgc(4, uint32Type, false)
		*(*uint32)(x) = val
	}
	return
}

func ConvT64(val uint64) (x unsafe.Pointer) {
	if val < uint64(len(staticuint64s)) {
		x = unsafe.Pointer(&staticuint64s[val])
	} else {
		x = mallocgc(8, uint64Type, false)
		*(*uint64)(x) = val
	}
	return
}

func ConvTstring(val string) (x unsafe.Pointer) {
	if len(val) == 0 {
		x = zr.UnsafePointer()
	} else {
		x = mallocgc(unsafe.Sizeof(val), stringType, true)
		*(*string)(x) = val
	}
	return
}

func ConvTslice(val []byte) (x unsafe.Pointer) {
	// Note: this must work for any element type, not just byte.
	if unsafe.SliceData(val) == nil {
		x = zr.UnsafePointer()
	} else {
		x = mallocgc(unsafe.Sizeof(val), sliceType, true)
		*(*[]byte)(x) = val
	}
	return
}

// interface type assertions x.(T)

func AssertI2I(inter *abi.InterfaceType, tab *abi.Itab) *abi.Itab {
	if tab == nil {
		// explicit conversions require non-nil interface value.
		assert.Panic(TypeAssertionError{nil, nil, &inter.Type, ""})
	}
	if tab.Inter == inter {
		return tab
	}
	return getitab(inter, tab.Type, false)
}

func AssertI2I2(inter *abi.InterfaceType, i Iface) (r Iface) {
	tab := i.Itab
	if tab == nil {
		return
	}
	if tab.Inter != inter {
		tab = getitab(inter, tab.Type, true)
		if tab == nil {
			return
		}
	}
	r.Itab = tab
	r.Data = i.Data
	return
}

func AssertE2I(inter *abi.InterfaceType, t *abi.Type) *abi.Itab {
	if t == nil {
		// explicit conversions require non-nil interface value.
		assert.Panic(TypeAssertionError{nil, nil, &inter.Type, ""})
	}
	return getitab(inter, t, false)
}

func AssertE2I2(inter *abi.InterfaceType, e Eface) (r Iface) {
	t := e.Type
	if t == nil {
		return
	}
	tab := getitab(inter, t, true)
	if tab == nil {
		return
	}
	r.Itab = tab
	r.Data = e.Data
	return
}

// interface equality. Type/itab pointers are already known to be equal, so
// we only need to pass one.

func IfaceEq(tab *abi.Itab, x, y unsafe.Pointer) bool {
	if tab == nil {
		return true
	}
	t := tab.Type
	eq := t.Equal
	if eq == nil {
		assert.Panic("comparing uncomparable type ", t.String())
	}
	if t.IsDirectIface() {
		// See comment in EfaceEq.
		return x == y
	}
	return eq(x, y)
}

func EfaceEq(t *abi.Type, x, y unsafe.Pointer) bool {
	if t == nil {
		return true
	}
	eq := t.Equal
	if eq == nil {
		assert.Panic("comparing uncomparable type ", t.String())
	}
	if t.IsDirectIface() {
		// Direct interface types are ptr, chan, map, func, and single-element structs/arrays thereof.
		// Maps and funcs are not comparable, so they can't reach here.
		// Ptrs, chans, and single-element items can be compared directly using ==.
		return x == y
	}
	return eq(x, y)
}

// PanicDotTypeE is called when doing an e.(T) conversion and the conversion fails.
// have = the dynamic type we have.
// want = the static type we're trying to convert to.
// iface = the static type we're converting from.
func PanicDotTypeE(have, want, iface *abi.Type) {
	assert.Panic(TypeAssertionError{
		_interface:    iface,
		concrete:      have,
		asserted:      want,
		missingMethod: "",
	})
}

// PanicDotTypeI is called when doing an i.(T) conversion and the conversion fails.
// Same args as panicdottypeE, but "have" is the dynamic itab we have.
func PanicDotTypeI(have *abi.Itab, want, iface *abi.Type) {
	var t *abi.Type
	if have != nil {
		t = have.Type
	}
	PanicDotTypeE(t, want, iface)
}

// PanicNilDotType is called when doing a i.(T) conversion and the interface i is nil.
// want = the static type we're trying to convert to.
func PanicNilDotType(want *abi.Type) {
	assert.Panic(TypeAssertionError{
		_interface:    nil,
		concrete:      nil,
		asserted:      want,
		missingMethod: "",
	})
	// TODO: Add the static type we're converting from as well.
	// It might generate a better error message.
	// Just to match other nil conversion errors, we don't for now.
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	stdconst "github.com/primecitizens/std/builtin/const"
	stdtype "github.com/primecitizens/std/builtin/type"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/asan"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/mem"
	"github.com/primecitizens/std/core/msan"
	"github.com/primecitizens/std/core/race"
)

//
// interface type operations
//

// Convert non-interface type to the data word of a (empty or nonempty) interface.
func convT(typ *abi.Type, elem unsafe.Pointer) unsafe.Pointer {
	if race.Enabled {
		race.ReadObjectPC(typ, elem, getcallerpc(), abi.FuncPCABIInternal(convT))
	}
	if msan.Enabled {
		msan.Read(elem, typ.Size_)
	}
	if asan.Enabled {
		asan.Read(elem, typ.Size_)
	}

	x := mallocgc(typ.Size_, typ, true)
	mem.TypedMove(typ, x, elem)
	return x
}

// Same as convT, for types with no pointers in them.
func convTnoptr(typ *abi.Type, elem unsafe.Pointer) unsafe.Pointer {
	// TODO: maybe take size instead of type?
	if race.Enabled {
		race.ReadObjectPC(typ, elem, getcallerpc(), abi.FuncPCABIInternal(convTnoptr))
	}
	if msan.Enabled {
		msan.Read(elem, typ.Size_)
	}
	if asan.Enabled {
		asan.Read(elem, typ.Size_)
	}

	x := mallocgc(typ.Size_, typ, false)
	mem.Move(x, elem, typ.Size_)
	return x
}

// Specialized versions of convT for specific types.
// These functions take concrete types in the runtime. But they may
// be used for a wider range of types, which have the same memory
// layout as the parameter type. The compiler converts the
// to-be-converted type to the parameter type before calling the
// runtime function. This way, the call is ABI-insensitive.

// staticuint64s is used to avoid allocating in convTx for small integer values.
var staticuint64s = [...]uint64{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
	0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
	0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
	0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
	0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57,
	0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
	0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
	0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
	0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
	0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
	0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
	0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
	0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
	0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7,
	0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
	0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7,
	0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf,
	0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7,
	0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf,
	0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7,
	0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf,
	0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7,
	0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef,
	0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7,
	0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff,
}

func elemTypeOf(ptr any) *abi.Type {
	return stdtype.TypeOf(ptr).PointerTypeUnsafe().Elem
}

func convT16(val uint16) (x unsafe.Pointer) {
	if val < uint16(len(staticuint64s)) {
		if arch.BigEndian {
			return unsafe.Pointer(uintptr(x) + 6)
		}
		return unsafe.Pointer(&staticuint64s[val])
	}

	x = mallocgc(stdconst.SizeUint16Type, elemTypeOf((*uint16)(nil)), false)
	*(*uint16)(x) = val
	return
}

func convT32(val uint32) (x unsafe.Pointer) {
	if val < uint32(len(staticuint64s)) {
		if arch.BigEndian {
			return unsafe.Pointer(uintptr(x) + 4)
		}
		return unsafe.Pointer(&staticuint64s[val])
	}

	x = mallocgc(stdconst.SizeUint32Type, elemTypeOf((*uint32)(nil)), false)
	*(*uint32)(x) = val
	return
}

func convT64(val uint64) unsafe.Pointer {
	if val < uint64(len(staticuint64s)) {
		return unsafe.Pointer(&staticuint64s[val])
	}

	x := mallocgc(stdconst.SizeUint64Type, elemTypeOf((*uint64)(nil)), false)
	*(*uint64)(x) = val
	return x
}

var (
	zeroVal [max(stdconst.SizeSliceType, stdconst.SizeStringType)]byte
)

func convTstring(val string) unsafe.Pointer {
	if len(val) == 0 {
		return unsafe.Pointer(&zeroVal[0])
	}

	x := mallocgc(stdconst.SizeStringType, elemTypeOf((*string)(nil)), true)
	*(*string)(x) = val
	return x
}

func convTslice(val []byte) unsafe.Pointer {
	// Note: this must work for any element type, not just byte.
	if unsafe.SliceData(val) == nil {
		return unsafe.Pointer(&zeroVal[0])
	}

	x := mallocgc(stdconst.SizeSliceType, elemTypeOf((*[]byte)(nil)), true)
	*(*[]byte)(x) = val
	return x
}

// panicdottypeE is called when doing an e.(T) conversion and the conversion fails.
// have = the dynamic type we have.
// want = the static type we're trying to convert to.
// iface = the static type we're converting from.
func panicdottypeE(have, want, iface *abi.Type) {
	assert.Panic(stdtype.TypeAssertionError{
		Interface:     iface,
		Concrete:      have,
		Asserted:      want,
		MissingMethod: "",
	})
}

// panicdottypeI is called when doing an i.(T) conversion and the conversion fails.
// Same args as panicdottypeE, but "have" is the dynamic itab we have.
func panicdottypeI(have *abi.Itab, want, iface *abi.Type) {
	var t *abi.Type
	if have != nil {
		t = have.Type
	}
	panicdottypeE(t, want, iface)
}

// panicnildottype is called when doing a i.(T) conversion and the interface i is nil.
// want = the static type we're trying to convert to.
func panicnildottype(want *abi.Type) {
	assert.Panic(stdtype.TypeAssertionError{
		Interface:     nil,
		Concrete:      nil,
		Asserted:      want,
		MissingMethod: "",
	})
	// TODO: Add the static type we're converting from as well.
	// It might generate a better error message.
	// Just to match other nil conversion errors, we don't for now.
}

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
		assert.Panic(stdtype.TypeAssertionError{
			Interface:     nil,
			Concrete:      typ,
			Asserted:      &inter.Type,
			MissingMethod: name.Name(),
		})
		return nil
	case abi.GetItabResult_:
		if canfail {
			return nil
		}
		assert.Panic(stdtype.TypeAssertionError{
			Concrete:      typ,
			Asserted:      &inter.Type,
			MissingMethod: itab.Init(),
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
func convI2I(dst *abi.InterfaceType, src *abi.Itab) *abi.Itab {
	if src == nil {
		return nil
	}
	if src.Inter == dst {
		return src
	}

	return getitab(dst, src.Type, false)
}

// interface type assertions x.(T)

func assertI2I(inter *abi.InterfaceType, tab *abi.Itab) *abi.Itab {
	if tab == nil {
		// explicit conversions require non-nil interface value.
		assert.Panic(stdtype.TypeAssertionError{
			Interface:     nil,
			Concrete:      nil,
			Asserted:      &inter.Type,
			MissingMethod: "",
		})
	}
	if tab.Inter == inter {
		return tab
	}
	return getitab(inter, tab.Type, false)
}

func assertI2I2(inter *abi.InterfaceType, i stdtype.Iface) (r stdtype.Iface) {
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

func assertE2I(inter *abi.InterfaceType, t *abi.Type) *abi.Itab {
	if t == nil {
		// explicit conversions require non-nil interface value.
		assert.Panic(stdtype.TypeAssertionError{
			Interface:     nil,
			Concrete:      nil,
			Asserted:      &inter.Type,
			MissingMethod: "",
		})
	}
	return getitab(inter, t, false)
}

func assertE2I2(inter *abi.InterfaceType, e stdtype.Eface) (r stdtype.Iface) {
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

func ifaceeq(tab *abi.Itab, x, y unsafe.Pointer) bool {
	if tab == nil {
		return true
	}
	t := tab.Type
	eq := t.Equal
	if eq == nil {
		assert.Panic("comparing", "uncomparable", "type ", t.String())
	}
	if t.IsDirectIface() {
		// See comment in EfaceEq.
		return x == y
	}
	return eq(x, y)
}

func efaceeq(t *abi.Type, x, y unsafe.Pointer) bool {
	if t == nil {
		return true
	}
	eq := t.Equal
	if eq == nil {
		assert.Panic("comparing", "uncomparable", "type ", t.String())
	}
	if t.IsDirectIface() {
		// Direct interface types are ptr, chan, map, func, and single-element structs/arrays thereof.
		// Maps and funcs are not comparable, so they can't reach here.
		// Ptrs, chans, and single-element items can be compared directly using ==.
		return x == y
	}
	return eq(x, y)
}

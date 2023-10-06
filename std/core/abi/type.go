// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/assert"
)

// Type is the runtime representation of a Go type.
//
// Type is also referenced implicitly
// (in the form of expressions involving constants and arch.PtrSize)
// in cmd/compile/internal/reflectdata/reflect.go
// and cmd/link/internal/ld/decodesym.go
// (e.g. data[2*arch.PtrSize+4] references the TFlag field)
// unsafe.OffsetOf(Type{}.TFlag) cannot be used directly in those
// places because it varies with cross compilation and experiments.
//
// NOTE: The type name has to be Type, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type Type struct {
	Size_       uintptr
	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	Hash        uint32  // hash of type; avoids computation in hash tables
	TFlag       TFlag   // extra type information flags
	Align_      uint8   // alignment of variable with this type
	FieldAlign_ uint8   // alignment of struct field with this type
	Kind_       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	// GCData stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, GCData is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	GCData    *byte
	Str       NameOff // string form
	PtrToThis TypeOff // type for pointer to this type, may be zero
}

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	KindInvalid Kind = iota
	KindBool
	KindInt
	KindInt8
	KindInt16
	KindInt32
	KindInt64
	KindUint
	KindUint8
	KindUint16
	KindUint32
	KindUint64
	KindUintptr
	KindFloat32
	KindFloat64
	KindComplex64
	KindComplex128
	KindArray
	KindChan
	KindKindFunc
	KindInterface
	KindMap
	KindPointer
	KindSlice
	KindString
	KindStruct
	KindUnsafePointer
)

const (
	// TODO (khr, drchase) why aren't these in TFlag?  Investigate, fix if possible.
	KindDirectIface = 1 << 5
	KindGCProg      = 1 << 6 // Type.gc points to GC program
	KindMask        = (1 << 5) - 1
)

// TFlag is used by a Type to signal what extra type information is
// available in the memory directly following the Type value.
type TFlag uint8

const (
	// TFlagUncommon means that there is a data with a type, UncommonType,
	// just beyond the shared-per-type common data.  That is, the data
	// for struct types will store their UncommonType at one offset, the
	// data for interface types will store their UncommonType at a different
	// offset.  UncommonType is always accessed via a pointer that is computed
	// using trust-us-we-are-the-implementors pointer arithmetic.
	//
	// For example, if t.Kind() == Struct and t.tflag&TFlagUncommon != 0,
	// then t has UncommonType data and it can be accessed as:
	//
	//	type structTypeUncommon struct {
	//		structType
	//		u UncommonType
	//	}
	//	u := &(*structTypeUncommon)(unsafe.Pointer(t)).u
	TFlagUncommon TFlag = 1 << 0

	// TFlagExtraStar means the name in the str field has an
	// extraneous '*' prefix. This is because for most types T in
	// a program, the type *T also exists and reusing the str data
	// saves binary size.
	TFlagExtraStar TFlag = 1 << 1

	// TFlagNamed means the type has a name.
	TFlagNamed TFlag = 1 << 2

	// TFlagRegularMemory means that equal and hash functions can treat
	// this type as a single region of t.size bytes.
	TFlagRegularMemory TFlag = 1 << 3
)

// NameOff is the offset to a name from moduledata.types.  See resolveNameOff in runtime.
type NameOff int32

// TypeOff is the offset to a type from moduledata.types.  See resolveTypeOff in runtime.
type TypeOff int32

// TextOff is an offset from the top of a text section.  See (rtype).textOff in runtime.
type TextOff int32

func (t *Type) Kind() Kind { return Kind(t.Kind_ & KindMask) }

func (t *Type) HasName() bool {
	return t.TFlag&TFlagNamed != 0
}

func (t *Type) Pointers() bool { return t.PtrBytes != 0 }

// IfaceIndir reports whether t is stored indirectly in an interface value.
func (t *Type) IfaceIndir() bool {
	return t.Kind_&KindDirectIface == 0
}

// IsDirectIface reports whether t is stored directly in an interface value.
//
// To find out what types stores value directly as iface.Data,
// See ${GOROOT}/src/cmd/internal/types/type.go#func:IsDirectIface
//
// As of go1.21, following types return true:
//   - pointer types whose elem type doesn't have mark.NotInHeap
//   - chan, map, func, unsafe.Pointer
//   - struct/array of 1 field/elem of above types
func (t *Type) IsDirectIface() bool {
	return t.Kind_&KindDirectIface != 0
}

func (t *Type) GcSlice(begin, end uintptr) []byte {
	return unsafe.Slice(t.GCData, int(end))[begin:]
}

func (t *Type) NameOff(off NameOff) Name {
	return resolveNameOff(unsafe.Pointer(t), off)
}

func (t *Type) TypeOff(off TypeOff) *Type {
	return resolveTypeOff(unsafe.Pointer(t), off)
}

func (t *Type) TextOff(off TextOff) unsafe.Pointer {
	return resolveTextOff(unsafe.Pointer(t), off)
}

func (t *Type) String() string {
	s := t.NameOff(t.Str).Name()
	if t.TFlag&TFlagExtraStar != 0 {
		return s[1:]
	}
	return s
}

func (t *Type) Name() string {
	if t.TFlag&TFlagNamed == 0 {
		return ""
	}
	s := t.String()
	i := len(s) - 1
	sqBrackets := 0
	for i >= 0 && (s[i] != '.' || sqBrackets != 0) {
		switch s[i] {
		case ']':
			sqBrackets++
		case '[':
			sqBrackets--
		}
		i--
	}
	return s[i+1:]
}

// PkgPath returns the path of the package where t was defined, if
// available. This is not the same as the reflect package's PkgPath
// method, in that it returns the package path for struct and interface
// types, not just named types.
func (t *Type) PkgPath() string {
	if u := t.Uncommon(); u != nil {
		return t.NameOff(u.PkgPath).Name()
	}
	switch t.Kind() {
	case KindStruct:
		st := (*StructType)(unsafe.Pointer(t))
		return st.PkgPath.Name()
	case KindInterface:
		it := (*InterfaceType)(unsafe.Pointer(t))
		return it.PkgPath.Name()
	}
	return ""
}

// Method on non-interface type
type Method struct {
	Name NameOff // name of method
	Mtyp TypeOff // method type (without receiver)
	Ifn  TextOff // fn used in interface call (one-word receiver)
	Tfn  TextOff // fn used for normal method call
}

// UncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
type UncommonType struct {
	PkgPath NameOff // import path; empty for built-in types like int, string
	Mcount  uint16  // number of methods
	Xcount  uint16  // number of exported methods
	Moff    uint32  // offset from this uncommontype to [mcount]Method
	_       uint32  // unused
}

func (t *UncommonType) Methods() []Method {
	if t.Mcount == 0 {
		return nil
	}
	return unsafe.Slice((*Method)(addChecked(unsafe.Pointer(t), uintptr(t.Moff), "t.mcount > 0")), t.Mcount)
}

func (t *UncommonType) ExportedMethods() []Method {
	if t.Xcount == 0 {
		return nil
	}
	return unsafe.Slice((*Method)(addChecked(unsafe.Pointer(t), uintptr(t.Moff), "t.xcount > 0")), t.Xcount)
}

// addChecked returns p+x.
//
// The whySafe string is ignored, so that the function still inlines
// as efficiently as p+x, but all call sites should use the string to
// record why the addition is safe, which is to say why the addition
// does not cause x to advance to the very end of p's allocation
// and therefore point incorrectly at the next block in memory.
func addChecked(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

// Imethod represents a method on an interface type
//
// NOTE: The type name has to be `Imethod`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type Imethod struct {
	Name NameOff // name of method
	Typ  TypeOff // .(*FuncType) underneath
}

// ArrayType represents a fixed array type.
//
// NOTE: The type name has to be ArrayType, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type ArrayType struct {
	Type
	Elem  *Type // array element type
	Slice *Type // slice type
	Len   uintptr
}

// Len returns the length of t if t is an array type, otherwise 0
func (t *Type) Len() int {
	if t.Kind() == KindArray {
		return int((*ArrayType)(unsafe.Pointer(t)).Len)
	}
	return 0
}

func (t *Type) Common() *Type {
	return t
}

type ChanDir int

const (
	ChanDirRecv    ChanDir = 1 << iota                 // <-chan
	ChanDirSend                                        // chan<-
	ChanDirBoth            = ChanDirRecv | ChanDirSend // chan
	ChanDirInvalid ChanDir = 0
)

// ChanType represents a channel type
//
// NOTE: The type name has to be ChanType, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type ChanType struct {
	Type
	Elem *Type
	Dir  ChanDir
}

type structTypeUncommon struct {
	StructType
	u UncommonType
}

// ChanDir returns the direction of t if t is a channel type, otherwise InvalidDir (0).
func (t *Type) ChanDir() ChanDir {
	if t.Kind() == KindChan {
		ch := (*ChanType)(unsafe.Pointer(t))
		return ch.Dir
	}
	return ChanDirInvalid
}

// Uncommon returns a pointer to T's "uncommon" data if there is any, otherwise nil
func (t *Type) Uncommon() *UncommonType {
	if t.TFlag&TFlagUncommon == 0 {
		return nil
	}
	switch t.Kind() {
	case KindStruct:
		return &(*structTypeUncommon)(unsafe.Pointer(t)).u
	case KindPointer:
		type u struct {
			PtrType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindKindFunc:
		type u struct {
			FuncType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindSlice:
		type u struct {
			SliceType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindArray:
		type u struct {
			ArrayType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindChan:
		type u struct {
			ChanType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindMap:
		type u struct {
			MapType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	case KindInterface:
		type u struct {
			InterfaceType
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	default:
		type u struct {
			Type
			u UncommonType
		}
		return &(*u)(unsafe.Pointer(t)).u
	}
}

// Elem returns the element type for t if t is an array, channel, map, pointer, or slice, otherwise nil.
func (t *Type) Elem() *Type {
	switch t.Kind() {
	case KindArray:
		return t.ArrayTypeUnsafe().Elem
	case KindChan:
		return t.ChanTypeUnsafe().Elem
	case KindMap:
		return t.MapTypeUnsafe().Elem
	case KindPointer:
		return t.PointerTypeUnsafe().Elem
	case KindSlice:
		return t.SliceTypeUnsafe().Elem
	}
	return nil
}

// ChanType returns t cast to a *ChanType, or nil if its tag does not match.
func (t *Type) ChanType() *ChanType {
	if t.Kind() != KindStruct {
		return nil
	}
	return t.ChanTypeUnsafe()
}

func (t *Type) ChanTypeUnsafe() *ChanType {
	return (*ChanType)(unsafe.Pointer(t))
}

// StructType returns t cast to a *StructType, or nil if its tag does not match.
func (t *Type) StructType() *StructType {
	if t.Kind() != KindStruct {
		return nil
	}
	return t.StructTypeUnsafe()
}

func (t *Type) StructTypeUnsafe() *StructType {
	return (*StructType)(unsafe.Pointer(t))
}

// MapType returns t cast to a *MapType, or nil if its tag does not match.
func (t *Type) MapType() *MapType {
	if t.Kind() != KindMap {
		return nil
	}
	return (*MapType)(unsafe.Pointer(t))
}

func (t *Type) MapTypeUnsafe() *MapType {
	return (*MapType)(unsafe.Pointer(t))
}

// SliceType returns t cast to a *SliceType, or nil if its tag does not match.
func (t *Type) SliceType() *SliceType {
	if t.Kind() != KindSlice {
		return nil
	}
	return t.SliceTypeUnsafe()
}

func (t *Type) SliceTypeUnsafe() *SliceType {
	return (*SliceType)(unsafe.Pointer(t))
}

// ArrayType returns t cast to a *ArrayType, or nil if its tag does not match.
func (t *Type) ArrayType() *ArrayType {
	if t.Kind() != KindArray {
		return nil
	}
	return t.ArrayTypeUnsafe()
}

func (t *Type) ArrayTypeUnsafe() *ArrayType {
	return (*ArrayType)(unsafe.Pointer(t))
}

// FuncType returns t cast to a *FuncType, or nil if its tag does not match.
func (t *Type) FuncType() *FuncType {
	if t.Kind() != KindKindFunc {
		return nil
	}
	return t.FuncTypeUnsafe()
}

func (t *Type) FuncTypeUnsafe() *FuncType {
	return (*FuncType)(unsafe.Pointer(t))
}

// InterfaceType returns t cast to a *InterfaceType, or nil if its tag does not match.
func (t *Type) InterfaceType() *InterfaceType {
	if t.Kind() != KindInterface {
		return nil
	}
	return t.InterfaceTypeUnsafe()
}

func (t *Type) InterfaceTypeUnsafe() *InterfaceType {
	return (*InterfaceType)(unsafe.Pointer(t))
}

// PointerType returns t cast to a *InterfaceType, or nil if its tag does not match.
func (t *Type) PointerType() *PtrType {
	if t.Kind() != KindPointer {
		return nil
	}

	return t.PointerTypeUnsafe()
}

func (t *Type) PointerTypeUnsafe() *PtrType {
	return (*PtrType)(unsafe.Pointer(t))
}

// Size returns the size of data with type t.
func (t *Type) Size() uintptr { return t.Size_ }

// Align returns the alignment of data with type t.
func (t *Type) Align() int { return int(t.Align_) }

func (t *Type) FieldAlign() int { return int(t.FieldAlign_) }

// NOTE: The type name has to be `InterfaceType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type InterfaceType struct {
	Type
	PkgPath Name      // import path
	Methods []Imethod // sorted by hash
}

func (t *Type) ExportedMethods() []Method {
	ut := t.Uncommon()
	if ut == nil {
		return nil
	}
	return ut.ExportedMethods()
}

func (t *Type) NumMethod() int {
	if t.Kind() == KindInterface {
		tt := (*InterfaceType)(unsafe.Pointer(t))
		return tt.NumMethod()
	}
	return len(t.ExportedMethods())
}

// NumMethod returns the number of interface methods in the type's method set.
func (t *InterfaceType) NumMethod() int { return len(t.Methods) }

// NOTE: The type name has to be `MapType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type MapType struct {
	Type
	Key    *Type
	Elem   *Type
	Bucket *Type // internal type representing a hash bucket
	// function for hashing keys (ptr to key, seed) -> hash
	Hasher     func(unsafe.Pointer, uintptr) uintptr
	KeySize    uint8  // size of key slot
	ValueSize  uint8  // size of elem slot
	BucketSize uint16 // size of bucket
	Flags      uint32
}

// Note: flag values must match those used in the TMAP case
// in ../cmd/compile/internal/reflectdata/reflect.go:writeType.
func (mt *MapType) IndirectKey() bool { // store ptr to key instead of key itself
	return mt.Flags&1 != 0
}
func (mt *MapType) IndirectElem() bool { // store ptr to elem instead of elem itself
	return mt.Flags&2 != 0
}
func (mt *MapType) ReflexiveKey() bool { // true if k==k for all keys
	return mt.Flags&4 != 0
}
func (mt *MapType) NeedKeyUpdate() bool { // true if we need to update key on an overwrite
	return mt.Flags&8 != 0
}
func (mt *MapType) HashMightPanic() bool { // true if hash function might panic
	return mt.Flags&16 != 0
}

func (t *Type) Key() *Type {
	if t.Kind() == KindMap {
		return (*MapType)(unsafe.Pointer(t)).Key
	}
	return nil
}

// NOTE: The type name has to be `SliceType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type SliceType struct {
	Type
	Elem *Type // slice element type
}

// funcType represents a function type.
//
// A *Type for each in and out parameter is stored in an array that
// directly follows the funcType (and possibly its uncommonType). So
// a function type with one method, one input, and one output is:
//
//	struct {
//		funcType
//		uncommonType
//		[2]*rtype    // [0] is in, [1] is out
//	}
//
// NOTE: The type name has to be `FuncType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type FuncType struct {
	Type
	InCount  uint16
	OutCount uint16 // top bit is set if last input parameter is ...
}

func (t *FuncType) In(i int) *Type {
	return t.InSlice()[i]
}

func (t *FuncType) NumIn() int {
	return int(t.InCount)
}

func (t *FuncType) NumOut() int {
	return int(t.OutCount & (1<<15 - 1))
}

func (t *FuncType) Out(i int) *Type {
	return (t.OutSlice()[i])
}

func (t *FuncType) InSlice() []*Type {
	uadd := unsafe.Sizeof(*t)
	if t.TFlag&TFlagUncommon != 0 {
		uadd += unsafe.Sizeof(UncommonType{})
	}
	if t.InCount == 0 {
		return nil
	}
	return unsafe.Slice((**Type)(addChecked(unsafe.Pointer(t), uadd, "t.inCount > 0")), t.InCount)
}

func (t *FuncType) OutSlice() []*Type {
	outCount := uint16(t.NumOut())
	if outCount == 0 {
		return nil
	}
	uadd := unsafe.Sizeof(*t)
	if t.TFlag&TFlagUncommon != 0 {
		uadd += unsafe.Sizeof(UncommonType{})
	}
	return (*[1 << 17]*Type)(addChecked(unsafe.Pointer(t), uadd, "outCount > 0"))[t.InCount : t.InCount+outCount : t.InCount+outCount]
}

func (t *FuncType) IsVariadic() bool {
	return t.OutCount&(1<<15) != 0
}

// NOTE: The type name has to be `PtrType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type PtrType struct {
	Type
	Elem *Type // pointer element (pointed at) type
}

type StructField struct {
	Name   Name    // name is always non-empty
	Typ    *Type   // type of field
	Offset uintptr // byte offset of field
}

func (f *StructField) Embedded() bool {
	return f.Name.IsEmbedded()
}

// NOTE: The type name has to be `StructType`, as expected by dwarf
// See ${GOROOT}/src/cmd/link/internal/ld/dwarf.go#func:dwarfGenerateDebugInfo
type StructType struct {
	Type
	PkgPath Name
	Fields  []StructField
}

// Name is an encoded type Name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the name is exported
//	1<<1 tag data follows the name
//	1<<2 pkgPath nameOff follows the name and tag
//	1<<3 the name is of an embedded (a.k.a. anonymous) field
//
// Following that, there is a varint-encoded length of the name,
// followed by the name itself.
//
// If tag data is present, it also has a varint-encoded length
// followed by the tag itself.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
//
// Note: this encoding must match here and in:
//
//	cmd/compile/internal/reflectdata/reflect.go
//	cmd/link/internal/ld/decodesym.go
type Name struct {
	Bytes *byte
}

// DataChecked does pointer arithmetic on n's Bytes, and that arithmetic is asserted to
// be safe for the reason in whySafe (which can appear in a backtrace, etc.)
func (n Name) DataChecked(off int, whySafe string) *byte {
	return (*byte)(addChecked(unsafe.Pointer(n.Bytes), uintptr(off), whySafe))
}

// Data does pointer arithmetic on n's Bytes, and that arithmetic is asserted to
// be safe because the runtime made the call (other packages use DataChecked)
func (n Name) Data(off int) *byte {
	return (*byte)(addChecked(unsafe.Pointer(n.Bytes), uintptr(off), "the runtime doesn't need to give you a reason"))
}

// IsExported returns "is n exported?"
func (n Name) IsExported() bool {
	return (*n.Bytes)&(1<<0) != 0
}

// HasTag returns true iff there is tag data following this name
func (n Name) HasTag() bool {
	return (*n.Bytes)&(1<<1) != 0
}

// IsEmbedded returns true iff n is embedded (an anonymous field).
func (n Name) IsEmbedded() bool {
	return (*n.Bytes)&(1<<3) != 0
}

// ReadVarint parses a varint as encoded by encoding/binary.
// It returns the number of encoded bytes and the encoded value.
func (n Name) ReadVarint(off int) (int, int) {
	v := 0
	for i := 0; ; i++ {
		x := *n.DataChecked(off+i, "read varint")
		v += int(x&0x7f) << (7 * i)
		if x&0x80 == 0 {
			return i + 1, v
		}
	}
}

// IsBlank indicates whether n is "_".
func (n Name) IsBlank() bool {
	if n.Bytes == nil {
		return false
	}
	_, l := n.ReadVarint(1)
	return l == 1 && *n.Data(2) == '_'
}

// writeVarint writes n to buf in varint form. Returns the
// number of bytes written. n must be nonnegative.
// Writes at most 10 bytes.
func writeVarint(buf []byte, n int) int {
	for i := 0; ; i++ {
		b := byte(n & 0x7f)
		n >>= 7
		if n == 0 {
			buf[i] = b
			return i + 1
		}
		buf[i] = b | 0x80
	}
}

// Name returns the name string for n, or empty if there is none.
func (n Name) Name() string {
	if n.Bytes == nil {
		return ""
	}
	i, l := n.ReadVarint(1)
	return unsafe.String(n.DataChecked(1+i, "non-empty string"), l)
}

// Tag returns the tag string for n, or empty if there is none.
func (n Name) Tag() string {
	if !n.HasTag() {
		return ""
	}
	i, l := n.ReadVarint(1)
	i2, l2 := n.ReadVarint(1 + i + l)
	return unsafe.String(n.DataChecked(1+i+l+i2, "non-empty string"), l2)
}

func NewName(n, tag string, exported, embedded bool) Name {
	if len(n) >= 1<<29 {
		assert.Throw("abi.NewName:", "name", "too", "long:", n[:1024], "...")
	}
	if len(tag) >= 1<<29 {
		assert.Throw("abi.NewName:", "tag", "too", "long:", tag[:1024], "...")
	}
	var nameLen [10]byte
	var tagLen [10]byte
	nameLenLen := writeVarint(nameLen[:], len(n))
	tagLenLen := writeVarint(tagLen[:], len(tag))

	var bits byte
	l := 1 + nameLenLen + len(n)
	if exported {
		bits |= 1 << 0
	}
	if len(tag) > 0 {
		l += tagLenLen + len(tag)
		bits |= 1 << 1
	}
	if embedded {
		bits |= 1 << 3
	}

	b := make([]byte, l)
	b[0] = bits
	copy(b[1:], nameLen[:nameLenLen])
	copy(b[1+nameLenLen:], n)
	if len(tag) > 0 {
		tb := b[1+nameLenLen+len(n):]
		copy(tb, tagLen[:tagLenLen])
		copy(tb[tagLenLen:], tag)
	}

	return Name{Bytes: &b[0]}
}

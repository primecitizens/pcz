// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

func New(constructor Func[Any], args ...Ref) Object {
	return Object{
		ref: bindings.New(
			constructor.ref,
			SliceData(args),
			SizeU(len(args)),
		),
	}
}

// TODO
type Record[T any] struct {
	ref bindings.Ref
}

func (r Record[T]) FromRef(ref Ref) Record[T] {
	return Record[T]{
		ref: bindings.Ref(ref),
	}
}

func (r Record[T]) Ref() Ref {
	return Ref(r.ref)
}

func (r Record[T]) Once() Record[T] {
	bindings.Once(r.ref)
	return r
}

func (r Record[T]) Free() {
	bindings.Free(r.ref)
}

// An Object is a reference to a js object.
type Object struct {
	ref bindings.Ref
}

func (o Object) FromRef(ref Ref) Object {
	return Object{
		ref: bindings.Ref(ref),
	}
}

func (o Object) Ref() Ref {
	return Ref(o.ref)
}

func (o Object) Once() Object {
	bindings.Once(o.ref)
	return o
}

func (o Object) Free() {
	bindings.Free(o.ref)
}

func (o Object) Prototype() (proto Object, hasPrototype bool) {
	r := bindings.GetPrototype(o.ref)
	if r == bindings.Ref(Null) {
		return
	}

	proto.ref = r
	hasPrototype = true
	return
}

// SetPrototype sets o's prototype to proto
//
// When free is ture, proto is freed after the call.
func (o Object) SetPrototype(free bool, proto Ref) bool {
	return bindings.Ref(True) == bindings.SetPrototype(
		o.ref,
		bindings.Ref(Bool(free)),
		bindings.Ref(proto),
	)
}

type PropDesc uint32

const (
	PropDesc_exists       PropDesc = 1 << iota
	PropDesc_writable              //
	PropDesc_configurable          //
	PropDesc_enumerable            //
	PropDesc_getter                //
	PropDesc_setter                //
)

func (o Object) PropDesc(name string) PropDesc {
	return PropDesc(bindings.GetPropDesc(
		o.ref,
		StringData(name),
		SizeU(len(name)),
	))
}

func (o Object) DefineProp(
	name string,
	flags PropDesc,
	getter Func[func() Any],
	setter Func[func(Any) bool],
) bool {
	return bindings.Ref(True) == bindings.DefineProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
		uint32(flags),
		getter.ref,
		setter.ref,
	)
}

func (o Object) DeleteProp(name string) bool {
	return bindings.Ref(True) == bindings.DeleteProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
	)
}

func (o Object) Prop(name string) Any {
	return Any{
		ref: bindings.Prop(
			o.ref,
			StringData(name),
			SizeU(len(name)),
		),
	}
}

func (o Object) StringProp(name string) String {
	return o.Prop(name).AsString()
}

func (o Object) NumProp(name string) float64 {
	return bindings.NumProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
	)
}

func (o Object) BoolProp(name string) bool {
	return bindings.Ref(True) == bindings.BoolProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
	)
}

func (o Object) PropByString(name String) Any {
	return Any{
		ref: bindings.PropByString(o.ref, name.ref),
	}
}

func (o Object) StringPropByString(name String) String {
	return o.PropByString(name).AsString()
}

func (o Object) NumPropByString(name String) float64 {
	return bindings.NumPropByString(o.ref, name.ref)
}

func (o Object) BoolPropByString(name String) bool {
	return bindings.Ref(True) == bindings.BoolPropByString(o.ref, name.ref)
}

func (o Object) SetProp(name string, free bool, val Ref) bool {
	return bindings.Ref(True) == bindings.SetProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
		bindings.Ref(Bool(free)),
		bindings.Ref(val),
	)
}

func (o Object) SetNumProp(name string, val float64) bool {
	return bindings.Ref(True) == bindings.SetNumProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
		val,
	)
}

func (o Object) SetBoolProp(name string, val bool) bool {
	return bindings.Ref(True) == bindings.SetBoolProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
		bindings.Ref(Bool(val)),
	)
}

func (o Object) SetStringProp(name, val string) bool {
	return bindings.Ref(True) == bindings.SetStringProp(
		o.ref,
		StringData(name),
		SizeU(len(name)),
		StringData(val),
		SizeU(len(val)),
	)
}

func (o Object) SetPropByString(prop String, free bool, val Ref) bool {
	return bindings.Ref(True) == bindings.SetPropByString(
		o.ref,
		prop.ref,
		bindings.Ref(Bool(free)),
		bindings.Ref(val),
	)
}

func (o Object) SetNumPropByString(prop String, val float64) bool {
	return bindings.Ref(True) == bindings.SetNumPropByString(
		o.ref, prop.ref, val,
	)
}

func (o Object) SetBoolPropByString(s String, val bool) bool {
	return bindings.Ref(True) == bindings.SetBoolPropByString(
		o.ref,
		s.ref,
		bindings.Ref(Bool(val)),
	)
}

func (o Object) SetStringPropByString(name String, free bool, val String) bool {
	return bindings.Ref(True) == bindings.SetStringPropByString(
		o.ref,
		name.ref,
		bindings.Ref(Bool(free)),
		val.ref,
	)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js_test

import (
	"github.com/primecitizens/std/ffi/js"
)

type JSValue interface {
	Ref() js.Ref
	Free()
}

var (
	_ JSValue = js.String{}
	_ JSValue = js.Array{}.Ref().AsArray()
	_ JSValue = js.Func{}.Ref().AsFunc()
	_ JSValue = js.NewPromise[int32](js.Global.AsFunc()).
		Then(js.Undefined.AsFunc(), js.Undefined.AsFunc()).
		Catch(js.False.AsFunc()).
		Finally(js.True.AsFunc())
)

var _ int = js.String{}.Into([]byte{})
var _ int = js.String{}.Len()
var _ bool = js.String{}.Equals("other string")

var (
	_ bool = js.Object{}.DefineProp(
		"foo", js.PropDesc_configurable, js.Undefined.AsFunc(), js.Null.AsFunc(),
	)
	_ bool = js.Object{}.DeleteProp("foo")
	_ bool = js.Object{}.SetProp("foo", true, js.NewString("woo").Ref())
	_ bool = js.Object{}.SetPropByString(
		js.NewString("prop"), true, js.NewPromise[uint32](js.Global.AsFunc()).Ref(),
	)
	_ bool = js.Object{}.SetNumProp("num", 10)
	_ bool = js.Object{}.SetNumPropByString(js.NewString("num"), 10)
	_ bool = js.Object{}.SetBoolProp("bool", true)
	_ bool = js.Object{}.SetBoolPropByString(js.NewString("bool"), true)
	_ bool = js.Object{}.SetStringProp("string", "a")
	_ bool = js.Object{}.SetStringPropByString(
		js.NewString("string"), true, js.NewString("a"),
	)
	_ js.Object = js.Object{}.Prop("prop").AsObject()
	_ js.Object = js.Object{}.PropByString(js.NewString("prop")).AsObject()
	_ js.Object = js.Object{}.StringProp("foo").Ref().AsObject()
	_ js.Object = js.Object{}.StringPropByString(js.NewString("str")).Ref().AsObject()

	_ bool      = js.Object{}.BoolProp("bool")
	_ js.Number = js.Object{}.NumProp("num")
)

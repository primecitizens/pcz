// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js_test

import (
	"github.com/primecitizens/pcz/std/ffi/js"
)

type JSValue interface {
	Ref() js.Ref
	Free()
}

var (
	_ JSValue = js.String{}
	_ JSValue = js.Array[js.Any]{}.Ref().Any().AsArray()
	_ JSValue = js.Func[js.Any]{}.Ref().Any().AsFunc()
)

var _ int = js.String{}.Read([]byte{})
var _ int = js.String{}.Len()
var _ bool = js.String{}.Equals("other string")

var (
	_ bool = js.Object{}.DefineProp(
		"foo", js.PropDesc_configurable,
		js.Func[func() js.Any]{}.FromRef(js.Undefined),
		js.Func[func(js.Any) bool]{}.FromRef(js.Undefined),
	)
	_ bool = js.Object{}.DeleteProp("foo")
	_ bool = js.Object{}.SetProp("foo", true, js.NewString("woo").Ref())
	_ bool = js.Object{}.SetPropByString(js.NewString("prop"), true, 0)

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
	_ js.Object = js.Object{}.StringProp("foo").Ref().Any().AsObject()
	_ js.Object = js.Object{}.StringPropByString(js.NewString("str")).Ref().Any().AsObject()

	_ bool    = js.Object{}.BoolProp("bool")
	_ float64 = js.Object{}.NumProp("num")
)

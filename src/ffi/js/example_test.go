package js_test

import "github.com/primecitizens/pcz/std/ffi/js"

type MyType struct {
}

func (T MyType) PromiseExecute(this js.Any) (value, reason js.Ref, fulfilled bool) {
	return
}

func ExampleGoPromise() {
	x := js.Executor[MyType]{}
	js.Promise[js.String]{}.FromRef(
		x.Reset(MyType{}, MyType.PromiseExecute),
	)
}

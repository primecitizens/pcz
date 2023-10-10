package main

import (
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/sys"
	"github.com/primecitizens/pcz/std/core/yield"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
	"github.com/primecitizens/pcz/std/time"
	"github.com/primecitizens/pcz/std/time/sysclock"

	// unless your application contains symbols expected by the go linker
	// add a runtime package
	_ "github.com/primecitizens/pcz/std/runtime"
)

var (
	popup         PopupPage
	serviceWorker ServiceWorker

	console   web.Console
	startTime int64
)

func init() {
	startTime = sysclock.Nanotime()

	comp, ok := sys.Arg(0)
	if !ok {
		assert.Throw("unexpected no component name")
	}

	switch comp {
	case "service-worker":
		serviceWorker.Init()
	case "popup":
		popup.Init()
	default:
		assert.Throw("unknown component name")
	}
}

func main() {
	println("wasm start, extension id:")
	console.Log(js.Must(runtime.Id()).Ref().Any().Once())

	println("Startup Time (ms)", (sysclock.Nanotime()-startTime)/time.Millisecond)

	// wait for events
	yield.Thread()
}

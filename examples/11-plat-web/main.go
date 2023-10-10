// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	_ "embed" // for go:embed
	"unsafe"

	"github.com/primecitizens/pcz/std/core/yield"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
	"github.com/primecitizens/pcz/std/time"
	"github.com/primecitizens/pcz/std/time/sysclock"
	"github.com/primecitizens/pcz/std/ui/html"

	// unless your application contains symbols expected by the go linker
	// add a runtime package
	_ "github.com/primecitizens/pcz/std/runtime"
)

var (
	//go:embed body.html
	htmlBody string
)

// global objects
var (
	window    = web.Window{}.FromRef(js.Global.Any().AsObject().Prop("window").Ref())
	navigator = js.Must(window.Navigator())
	document  = js.Must(window.Document())
	head      = js.Must(document.Head())
	body      = js.Must(document.Body())

	emptyString = js.NewString("")

	css     web.CSS
	console web.Console
)

var startTime int64

var (
	exampleWebGPU ExampleWebGPU
	exampleEvent  ExampleEvent

	activeExample interface {
		Stop()
	}

	btnExampleEvent  web.HTMLButtonElement
	btnExampleFetch  web.HTMLButtonElement
	btnExampleWebGPU web.HTMLButtonElement

	displayArea   web.HTMLDivElement
	displayBuffer html.Builder
)

func init() {
	startTime = sysclock.Nanotime()

	// Using `htmlBody` is only intended to showcase you can do so
	// and build components this way.
	//
	// but we don't recommend doing it in wasm.
	body.SetInnerHTML(js.NewString(htmlBody).Once())

	html.FindElement(&btnExampleEvent, document, "btn-event")
	html.FindElement(&btnExampleFetch, document, "btn-fetch")
	html.FindElement(&btnExampleWebGPU, document, "btn-webgpu")
	html.FindElement(&displayArea, document, "example")

	displayBuffer.Reset(displayArea.Ref())

	click := js.NewString("click")
	defer click.Free()

	t := web.OneOf_AddEventListenerOptions_Bool{}.FromRef(js.True)

	listener := web.EventListenerFunc(selectExample).Register()
	defer listener.Free()

	for _, elem := range []web.EventTarget{
		btnExampleEvent.EventTarget,
		btnExampleFetch.EventTarget,
		btnExampleWebGPU.EventTarget,
	} {
		elem.AddEventListener(click, listener, t)
	}
}

func main() {
	println("Startup Time (ms):", (sysclock.Nanotime()-startTime)/time.Millisecond)

	// actually we only need to call yield.Thread() once ... but anyway, this also works.
	for {
		println("Waiting user events...")
		yield.Thread()
	}
}

func selectExample(this js.Ref, event web.Event) js.Ref {
	var buf [64]byte
	id := js.Must(web.HTMLButtonElement{}.FromRef(js.Must(event.Target()).Ref().Once()).Id()).Once()
	x := buf[:id.Read(buf[:])]
	if activeExample != nil {
		activeExample.Stop()
	}

	elemID := unsafe.String(unsafe.SliceData(x), len(x))

	displayBuffer.Discard().Flush(false)

	switch elemID {
	case "btn-event":
		activeExample = &exampleEvent
		exampleEvent.Run()
	case "btn-fetch":
		activeExample = nil
		ExampleFetch()
	case "btn-webgpu":
		activeExample = &exampleWebGPU
		exampleWebGPU.Run()
	default:
		println("unknown target", elemID)
	}

	return js.Undefined
}

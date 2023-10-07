// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
)

var (
	evtMousemove = js.NewString("mousemove")
)

type ExampleEvent struct {
	initialized bool
	running     bool
	listener    js.Func[func(web.Event)]

	callbackContext web.EventListener[*ExampleEvent]
}

func (x *ExampleEvent) Run() {
	if x.running {
		return
	}

	if !x.initialized {
		x.callbackContext = web.EventListener[*ExampleEvent]{
			Fn:  (*ExampleEvent).HandleMouseEvent,
			Arg: x,
		}
		x.listener = x.callbackContext.Register()
	}

	x.running = true
	window.AddEventListener1(evtMousemove, x.listener)
	return
}

func (x *ExampleEvent) HandleMouseEvent(this js.Ref, event web.Event) js.Ref {
	evt := web.MouseEvent{}.FromRef(event.Ref())
	displayBuffer.Html("<p>").
		Text("x = ").Int(int64(js.Must(evt.ClientX())), 10).
		Text(", y = ").Int(int64(js.Must(evt.ClientY())), 10).
		Html("</p>").Flush(true)

	return js.Undefined
}

func (x *ExampleEvent) Stop() {
	if !x.running {
		return
	}

	x.running = false
	window.RemoveEventListener1(evtMousemove, x.listener)
}

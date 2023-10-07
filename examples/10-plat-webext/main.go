// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/sys"
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
	console web.Console

	startTime int64
)

func init() {
	startTime = sysclock.Nanotime()
}

func main() {
	println("wasm start, extension id:")
	console.Log(js.Must(runtime.Id()).Ref().Any().Once())

	println("Startup Time (ms)", (sysclock.Nanotime()-startTime)/time.Millisecond)

	mode, ok := sys.Arg(0)
	if !ok {
		assert.Throw("unexpected no args set")
	}

	println("mode", mode)
	switch mode {
	case "--service-worker":
		serviceWorker()
	case "--popup":
		popup()
	default:
		println("unknown mode")
	}
}

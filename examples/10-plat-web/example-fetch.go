// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/thread"
	"github.com/primecitizens/std/encoding/binfmt/wasm"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web"
	"github.com/primecitizens/std/time"
	"github.com/primecitizens/std/time/sysclock"
)

var (
	// link time variable to be set using `-X main.wasmPath=/web/server/path/to/wasm`
	//
	// if not set, defaults to "/app.wasm"
	wasmPath string
)

func ExampleFetch() {
	displayBuffer.TextBlock("<p>", "</p>",
		"fetching wasm from ", wasmPath, " ...",
	).Flush(false)

	start := sysclock.Nanotime()

	if len(wasmPath) == 0 {
		wasmPath = "/app.wasm"
	}

	resp, err, ok := js.Must(window.Fetch(
		web.RequestInfo{}.FromRef(js.NewString(wasmPath).Ref().Once()),
		web.RequestInit{
			Method:            js.NewString("GET").Once(),
			ReferrerPolicy:    web.ReferrerPolicy_NO_REFERRER,
			Mode:              web.RequestMode_NO_CORS,
			Credentials:       web.RequestCredentials_OMIT,
			Cache:             web.RequestCache_NO_CACHE,
			Redirect:          web.RequestRedirect_ERROR,
			FFI_USE_Keepalive: false,
			Keepalive:         false,
			Priority:          web.RequestPriority_HIGH,
		},
	)).Await(true)
	if !ok {
		console.Log(err)
		return
	}

	blob, err, ok := js.Must(resp.Once().Blob()).Await(true)
	if !ok {
		console.Log(err)
		return
	}

	arr, err, ok := js.Must(blob.Once().ArrayBuffer()).Await(true)
	if !ok {
		console.Log(err)
		return
	}

	data := js.TypedArray[byte]{}.FromArrayBuffer(true, arr)
	sz := data.ByteLength()
	displayBuffer.HTML("<p>").
		Text("fetched ", wasmPath, ", size = ").Uint(uint64(sz), 10).
		Text(", took ").Uint(uint64((sysclock.Nanotime()-start)/time.Millisecond), 10).Text("ms").
		HTML("</p>").
		Flush(true)

	malloc := thread.G().G().DefaultAlloc()
	buf := alloc.Make[byte](malloc, sz, sz)
	defer alloc.Free(malloc, buf)

	if data.Copy(0, buf) != sz {
		console.Log(js.Any(js.NewString("not all bytes copied").Once()))
		return
	}

	displayBuffer.TextBlock("<p>", "</p>", "inspecting wasm...").Flush(true)

	var (
		hdr  wasm.Header
		sect wasm.Section
		imp  wasm.ImportSection
		ent  wasm.ImportEntry
	)
	buf, ok = hdr.Decode(buf)

	displayBuffer.HTML("<p>").
		Text("wasm binary format version = ").Uint(uint64(hdr.Version.Get()), 10).
		HTML("</p>").Flush(true)

	for ok {
		buf, ok = sect.Decode(buf)
		switch sect.ID {
		case wasm.Section_CODE:
		case wasm.Section_IMPORT:
			var ok bool
			for sect.Data, ok = imp.Decode(sect.Data); ok; {
				if sect.Data, ok = ent.Parse(sect.Data); !ok {
					break
				}

				displayBuffer.TextBlock("<p>", "</p>", "imported &", ent.Field, "@", ent.Module).Flush(true)
			}
		}
	}
}

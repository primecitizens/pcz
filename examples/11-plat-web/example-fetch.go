// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/thread"
	"github.com/primecitizens/pcz/std/encoding/binfmt/wasm"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
	"github.com/primecitizens/pcz/std/time"
	"github.com/primecitizens/pcz/std/time/sysclock"
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

	resp, err, ok := window.Fetch(
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
	).Await(true)
	if !ok {
		console.Log(err.Once())
		return
	}

	blob, err, ok := resp.Once().Blob().Await(true)
	if !ok {
		console.Log(err.Once())
		return
	}

	arr, err, ok := blob.Once().ArrayBuffer().Await(true)
	if !ok {
		console.Log(err.Once())
		return
	}
	defer arr.Free()

	data := js.TypedArray[byte]{}.FromArrayBuffer(true, arr)
	sz := data.ByteLength()
	displayBuffer.Html("<p>").
		Text("fetched ", wasmPath, ", size = ").Uint(uint64(sz), 10).
		Text(", took ").Uint(uint64((sysclock.Nanotime()-start)/time.Millisecond), 10).Text("ms").
		Html("</p>").
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

	displayBuffer.Html("<p>").
		Text("wasm binary format version = ").Uint(uint64(hdr.Version.Get()), 10).
		Html("</p>").Flush(true)

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

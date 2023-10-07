// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/core/yield"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
	"github.com/primecitizens/pcz/std/plat/js/webext/bookmarks"
)

var (
	document = web.Document{}.FromRef(js.Global.Any().AsObject().Prop("document").Ref())
	body, _  = document.Body()
	strHref  = js.NewString("href")
	strEmpty = js.NewString("")
	str_ul   = js.NewString("ul")
	str_li   = js.NewString("li")
	strA     = js.NewString("a")
)

func popup() {
	println("running popup")

	document.AddEventListener1(
		js.NewString("DOMContentLoaded").Once(),
		web.EventListenerFunc(onLoad).Register().Once(),
	)

	onLoad(js.Undefined, web.Event{})

	for {
		println("waiting for popup events")
		yield.Thread()
	}
}

func onLoad(js.Ref, web.Event) js.Ref {
	tree, err, ok := bookmarks.GetTree().Await(true)
	if !ok {
		console.Log(err.Once())
		return js.Undefined
	}

	body.SetInnerHTML(strEmpty)

	ul := web.HTMLUListElement{}.FromRef(document.CreateElement1(str_ul).Ref())
	body.AppendChild(ul.Node).Free()
	listBookmarks(ul, tree)
	return js.Undefined
}

func listBookmarks(dst web.HTMLUListElement, tree js.Array[bookmarks.BookmarkTreeNode]) {
	if tree.Ref().Falsy() {
		return
	}

	for i := 0; ; i++ {
		ref, ok := tree.Nth(i)
		if !ok {
			break
		}

		appendBookmark(dst, ref)
		ref.Free()
	}
}

func appendBookmark(dst web.HTMLUListElement, ref js.Ref) {
	if ref.Falsy() {
		return
	}

	entry := bookmarks.BookmarkTreeNode{}.FromRef(ref)
	defer entry.FreeMembers(true)

	if entry.Url.Ref().Truthy() {
		li := web.HTMLLIElement{}.FromRef(document.CreateElement1(str_li).Ref())

		link := web.HTMLElement{}.FromRef(document.CreateElement1(strA).Ref())
		link.SetAttribute(strHref, entry.Url)
		link.SetInnerText(entry.Title)

		dst.AppendChild(li.Once().AppendChild(link.Node.Once()).Once()).Free()
	}

	listBookmarks(dst, entry.Children)
}

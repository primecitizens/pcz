// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web"
	"github.com/primecitizens/pcz/std/plat/js/webext/bookmarks"
)

var (
	str_empty = js.NewString("")
	str_href  = js.NewString("href")
	str_ul    = js.NewString("ul")
	str_li    = js.NewString("li")
	str_a     = js.NewString("a")
)

type PopupPage struct {
	document  web.Document
	onLoadCtx web.EventListener[*PopupPage]
}

func (p *PopupPage) Init() {
	p.onLoadCtx = web.EventListener[*PopupPage]{
		Fn:  (*PopupPage).onLoad,
		Arg: p,
	}
	p.document = js.Must(web.Window{}.FromRef(js.Global).Document())

	p.document.AddEventListener1(
		js.NewString("DOMContentLoaded").Once(),
		p.onLoadCtx.Register().Once(),
	)
}

func (p *PopupPage) onLoad(this js.Ref, event web.Event) js.Ref {
	tree, err, ok := bookmarks.GetTree().Await(true)
	if !ok {
		console.Log(err.Once())
		return js.Undefined
	}

	body := js.Must(p.document.Body())
	body.SetInnerHTML(str_empty)

	ul := web.HTMLUListElement{}.FromRef(p.document.CreateElement1(str_ul).Ref())
	body.AppendChild(ul.Node).Free()
	p.listBookmarks(ul, tree)
	return js.Undefined
}

func (p *PopupPage) listBookmarks(dst web.HTMLUListElement, tree js.Array[bookmarks.BookmarkTreeNode]) {
	if tree.Ref().Falsy() {
		return
	}

	for i := 0; ; i++ {
		ref, ok := tree.Nth(i)
		if !ok {
			break
		}

		p.appendBookmark(dst, ref)
		ref.Free()
	}
}

func (p *PopupPage) appendBookmark(dst web.HTMLUListElement, ref js.Ref) {
	if ref.Falsy() {
		return
	}

	entry := bookmarks.BookmarkTreeNode{}.FromRef(ref)
	defer entry.FreeMembers(true)

	if entry.Url.Ref().Truthy() {
		li := web.HTMLLIElement{}.FromRef(p.document.CreateElement1(str_li).Ref())

		link := web.HTMLElement{}.FromRef(p.document.CreateElement1(str_a).Ref())
		link.SetAttribute(str_href, entry.Url)
		link.SetInnerText(entry.Title)

		dst.AppendChild(li.Once().AppendChild(link.Node.Once()).Once()).Free()
	}

	p.listBookmarks(dst, entry.Children)
}

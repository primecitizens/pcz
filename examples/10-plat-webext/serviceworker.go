// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/core/yield"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contextmenus"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

func serviceWorker() {
	println("running serviceworker")

	allTabs, err, ok := tabs.Query(tabs.QueryArgQueryInfo{}).Await(true)
	if !ok {
		println("failed", "to", "list", "tabs")
		console.Log(err.Once())
	}

	var t tabs.Tab
	for i := 0; ; i++ {
		tabRef, ok := allTabs.Nth(i)
		if !ok {
			break
		}

		t.UpdateFrom(tabRef.Once())

		println("id =", t.Id, ", active =", t.Active)
		t.FreeMembers(true)
	}
	allTabs.Free()

	tabs.OnRemoved(tabs.OnRemovedEventCallbackFunc(onTabRemoved).Register().Once())
	tabs.OnCreated(tabs.OnCreatedEventCallbackFunc(onTabCreated).Register().Once())

	// TODO: keepalive
	contextmenus.Create(
		contextmenus.CreateArgCreateProperties{
			Contexts: js.NewArrayOf[contextmenus.ContextType](true,
				js.NewString(js.Must(contextmenus.ContextType_SELECTION.String())).Ref().Once(),
			),
			// FFI_USE_Enabled:     false,
			// Enabled:             false,
			Id:    js.NewString("example-wasm-contextmenu").Once(),
			Title: js.NewString("Example WASM ContextMenu").Once(),
			Type:  contextmenus.ItemType_NORMAL,
			// FFI_USE_Visible:     false,
			// Visible:             false,
			// FFI_USE_Checked: false,
			// Checked:         false,
		},
		contextmenus.CreateArgCallbackFunc(onExampleContextMenuClick).Register().Once(),
	)
	contextmenus.OnClicked(contextmenus.OnClickedEventCallbackFunc(onContextMenuClick).Register().Once())
	runtime.OnSuspend(runtime.OnSuspendEventCallbackFunc(onSuspend).Register().Once())
	for {
		println("waiting for service-worker events")
		yield.Thread()
	}
}

func onExampleContextMenuClick(js.Ref) js.Ref {
	println("item clicked")
	return js.Undefined
}

func onContextMenuClick(this js.Ref, info *contextmenus.OnClickData, tab *tabs.Tab) js.Ref {
	println("menu clicked", tab.Id)
	console.Log(info.MenuItemId.Ref().Any())
	return js.Undefined
}

func onSuspend(this js.Ref) js.Ref {
	println("suspend")
	return js.Undefined
}

func onTabRemoved(this js.Ref, tabId int64, removeInfo *tabs.OnRemovedArgRemoveInfo) js.Ref {
	println("tab removed id =", tabId)
	tabs.Create(tabs.CreateArgCreateProperties{
		FFI_USE_Active: false,
		Active:         false,

		FFI_USE_Index: false,
		Index:         0,

		FFI_USE_OpenerTabId: false,
		OpenerTabId:         tabId,

		FFI_USE_Pinned: false,
		Pinned:         false,

		FFI_USE_Selected: false,
		Selected:         false,

		Url: js.NewString("chrome://extensions").Once(),

		FFI_USE_WindowId: false,
		WindowId:         0,
	})
	return js.Undefined
}

func onTabCreated(this js.Ref, tab *tabs.Tab) js.Ref {
	println("tab created, id = ", tab.Id)
	return js.Undefined
}

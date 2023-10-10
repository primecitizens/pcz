// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contextmenus"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

var (
	url_extension = js.NewString("chrome://extensions")
)

type ServiceWorker struct {
	onInstalledCtx runtime.OnInstalledEventCallback[*ServiceWorker]
}

func (w *ServiceWorker) Init() {
	w.onInstalledCtx = runtime.OnInstalledEventCallback[*ServiceWorker]{
		Fn:  (*ServiceWorker).onInstalled,
		Arg: w,
	}

	// Actions only need to be done once after the extension installation
	// should be done inside the callback of WEBEXT.runtime.onInstalled.addListener.
	runtime.OnInstalled(serviceWorker.onInstalledCtx.Register().Once())

	// When running as a ServiceWorker, extension event listeners
	// MUST be registered synchronously.
	//
	// See ../docs/user/11-plat-webext.md#general-guide

	tabs.OnRemoved(
		tabs.OnRemovedEventCallbackFunc(onTabRemoved).Register().Once(),
	)
	tabs.OnCreated(
		tabs.OnCreatedEventCallbackFunc(onTabCreated).Register().Once(),
	)
	contextmenus.OnClicked(
		contextmenus.OnClickedEventCallbackFunc(onContextMenuClick).Register().Once(),
	)
}

func (w *ServiceWorker) onInstalled(this js.Ref, details *runtime.OnInstalledArgDetails) js.Ref {
	contextmenus.Create(
		contextmenus.CreateArgCreateProperties{
			Contexts: js.NewArrayOf[contextmenus.ContextType](true,
				js.NewString(js.Must(contextmenus.ContextType_SELECTION.String())).Ref().Once(),
			),
			Id:    js.NewString("example-go-contextmenu").Once(),
			Title: js.NewString("Example Go ContextMenu").Once(),
			Type:  contextmenus.ItemType_NORMAL,
		},
		js.Func[func()]{}, // undefined
	).Free()

	return js.Undefined
}

func onContextMenuClick(this js.Ref, info *contextmenus.OnClickData, tab *tabs.Tab) js.Ref {
	println("context menu clicked from tab =", tab.Id)
	console.Log(info.MenuItemId.Ref().Any())
	return js.Undefined
}

func onTabCreated(this js.Ref, tab *tabs.Tab) js.Ref {
	println("tab created, id = ", tab.Id)
	return js.Undefined
}

func onTabRemoved(this js.Ref, tabId int64, removeInfo *tabs.OnRemovedArgRemoveInfo) js.Ref {
	println("tab removed id =", tabId)
	tabs.Create(tabs.CreateArgCreateProperties{
		Url: url_extension,
	}).Free()
	return js.Undefined
}

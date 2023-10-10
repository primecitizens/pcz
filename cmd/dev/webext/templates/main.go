// + package main
package templates

import (
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/sys"
	"github.com/primecitizens/pcz/std/core/yield"

	// {{- if or .ServiceWorker .Popup .Settings .SidePanel }}
	"github.com/primecitizens/pcz/std/ffi/js"
	// {{- end }}

	// {{- if or .Popup .Settings }}
	"github.com/primecitizens/pcz/std/plat/js/web"
	// {{- end }}

	// {{- if .ServiceWorker }}
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
	// {{- end }}

	// {{- if .SidePanel }}
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
	// {{- end }}

	_ "github.com/primecitizens/pcz/std/runtime"
)

var (
	// {{- if .Popup }}
	popup PopupPage
	// {{- end }}

	// {{- if .ServiceWorker }}
	serviceWorker ServiceWorker
	// {{- end }}

	// {{- if .Settings }}
	appSettings SettingsPage
	// {{- end }}

	// {{- if .SidePanel }}
	sidePanel SidePanel
	// {{- end }}
)

func init() {
	comp, ok := sys.Arg(0)
	if !ok {
		assert.Throw("unexpected no component name")
	}

	switch comp {
	// {{- if .ServiceWorker }}
	case "service-worker":
		serviceWorker.Init()
	// {{- end }}
	// {{- if .Popup }}
	case "popup":
		popup.Init()
	// {{- end }}
	// {{- if .Settings }}
	case "settings":
		appSettings.Init()
	// {{- end }}
	// {{- if .SidePanel }}
	case "sidepanel":
		sidePanel.Init()
	// {{- end }}
	default:
		assert.Throw("unknown component name")
	}
}

func main() {
	// wait for events
	yield.Thread()
}

// {{-  if .ServiceWorker }}

type ServiceWorker struct {
	// TODO: add callback contexts for extension events

	onInstalledCtx runtime.OnInstalledEventCallback[*ServiceWorker]
}

func (w *ServiceWorker) Init() {
	w.onInstalledCtx = runtime.OnInstalledEventCallback[*ServiceWorker]{
		Fn:  (*ServiceWorker).onInstalled,
		Arg: w,
	}
	runtime.OnInstalled(serviceWorker.onInstalledCtx.Register().Once())

	// TODO: register extension event listeners synchronously
}

func (w *ServiceWorker) onInstalled(this js.Ref, details *runtime.OnInstalledArgDetails) js.Ref {
	// TODO: add one-time operations after extension installation
	return js.Undefined
}

// {{- end }}

// {{- if .Popup }}

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
	body := js.Must(p.document.Body())

	// TODO: add widgets and listeners to your popup page
	_ = body

	return js.Undefined
}

// {{- end }}

// {{- if .Settings }}

type SettingsPage struct {
	document web.Document

	exampleEventCtx web.EventListener[*SettingsPage]
}

func (p *SettingsPage) Init() {
	// TODO: add widgets and listeners to your options page
	p.document = js.Must(web.Window{}.FromRef(js.Global).Document())

	p.exampleEventCtx = web.EventListener[*SettingsPage]{
		Fn:  (*SettingsPage).exampleOnButtonClick,
		Arg: p,
	}

	exampleButton := web.HTMLButtonElement{}.FromRef(
		p.document.CreateElement1(js.NewString("button").Once()).Ref(),
	)
	exampleButton.SetInnerText(js.NewString("Example Button").Once())
	exampleButton.AddEventListener1(
		js.NewString("click").Once(),
		p.exampleEventCtx.Register().Once(),
	)

	js.Must(p.document.Body()).AppendChild(exampleButton.Node.Once()).Free()
}

func (p *SettingsPage) exampleOnButtonClick(this js.Ref, event web.Event) js.Ref {
	// TODO: handle button click
	println("btn click")
	return js.Undefined
}

// {{- end }}

// {{- if .SidePanel }}

type SidePanel struct {
	onTabsUpdateCtx tabs.OnUpdatedEventCallback[*SidePanel]
}

func (p *SidePanel) Init() {
	// TODO: register extension event listeners synchronously

	p.onTabsUpdateCtx = tabs.OnUpdatedEventCallback[*SidePanel]{
		Fn:  (*SidePanel).onTabsUpdate,
		Arg: p,
	}

	tabs.OnUpdated(p.onTabsUpdateCtx.Register().Once())
}

func (arg *SidePanel) onTabsUpdate(this js.Ref, tabId int64, changeInfo *tabs.OnUpdatedArgChangeInfo, tab *tabs.Tab) js.Ref {
	// TODO: react to tab updates (e.g. site specific panel display)
	return js.Undefined
}

// {{- end }}

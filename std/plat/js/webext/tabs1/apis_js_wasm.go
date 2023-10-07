// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tabs1

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs1/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/windows"
)

// HasFuncConnect returns true if the function "WEBEXT.tabs.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.tabs.connect".
func FuncConnect() (fn js.Func[func(tabId int64, connectInfo tabs.ConnectArgConnectInfo) runtime.Port]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.tabs.connect" directly.
func Connect(tabId int64, connectInfo tabs.ConnectArgConnectInfo) (ret runtime.Port) {
	bindings.CallConnect(
		js.Pointer(&ret),
		float64(tabId),
		js.Pointer(&connectInfo),
	)

	return
}

// TryConnect calls the function "WEBEXT.tabs.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(tabId int64, connectInfo tabs.ConnectArgConnectInfo) (ret runtime.Port, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
		js.Pointer(&connectInfo),
	)

	return
}

// HasFuncHighlight returns true if the function "WEBEXT.tabs.highlight" exists.
func HasFuncHighlight() bool {
	return js.True == bindings.HasFuncHighlight()
}

// FuncHighlight returns the function "WEBEXT.tabs.highlight".
func FuncHighlight() (fn js.Func[func(highlightInfo tabs.HighlightArgHighlightInfo) js.Promise[windows.Window]]) {
	bindings.FuncHighlight(
		js.Pointer(&fn),
	)
	return
}

// Highlight calls the function "WEBEXT.tabs.highlight" directly.
func Highlight(highlightInfo tabs.HighlightArgHighlightInfo) (ret js.Promise[windows.Window]) {
	bindings.CallHighlight(
		js.Pointer(&ret),
		js.Pointer(&highlightInfo),
	)

	return
}

// TryHighlight calls the function "WEBEXT.tabs.highlight"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHighlight(highlightInfo tabs.HighlightArgHighlightInfo) (ret js.Promise[windows.Window], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHighlight(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&highlightInfo),
	)

	return
}

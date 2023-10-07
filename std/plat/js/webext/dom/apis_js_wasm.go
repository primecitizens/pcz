// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package dom

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/dom/bindings"
)

// HasFuncOpenOrClosedShadowRoot returns true if the function "WEBEXT.dom.openOrClosedShadowRoot" exists.
func HasFuncOpenOrClosedShadowRoot() bool {
	return js.True == bindings.HasFuncOpenOrClosedShadowRoot()
}

// FuncOpenOrClosedShadowRoot returns the function "WEBEXT.dom.openOrClosedShadowRoot".
func FuncOpenOrClosedShadowRoot() (fn js.Func[func(element js.Any) js.Any]) {
	bindings.FuncOpenOrClosedShadowRoot(
		js.Pointer(&fn),
	)
	return
}

// OpenOrClosedShadowRoot calls the function "WEBEXT.dom.openOrClosedShadowRoot" directly.
func OpenOrClosedShadowRoot(element js.Any) (ret js.Any) {
	bindings.CallOpenOrClosedShadowRoot(
		js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryOpenOrClosedShadowRoot calls the function "WEBEXT.dom.openOrClosedShadowRoot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenOrClosedShadowRoot(element js.Any) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenOrClosedShadowRoot(
		js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

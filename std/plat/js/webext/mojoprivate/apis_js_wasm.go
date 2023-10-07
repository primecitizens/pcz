// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mojoprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/mojoprivate/bindings"
)

// HasFuncRequireAsync returns true if the function "WEBEXT.mojoPrivate.requireAsync" exists.
func HasFuncRequireAsync() bool {
	return js.True == bindings.HasFuncRequireAsync()
}

// FuncRequireAsync returns the function "WEBEXT.mojoPrivate.requireAsync".
func FuncRequireAsync() (fn js.Func[func(name js.String) js.Any]) {
	bindings.FuncRequireAsync(
		js.Pointer(&fn),
	)
	return
}

// RequireAsync calls the function "WEBEXT.mojoPrivate.requireAsync" directly.
func RequireAsync(name js.String) (ret js.Any) {
	bindings.CallRequireAsync(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRequireAsync calls the function "WEBEXT.mojoPrivate.requireAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequireAsync(name js.String) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequireAsync(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

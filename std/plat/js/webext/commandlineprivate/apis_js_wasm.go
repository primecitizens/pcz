// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package commandlineprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/commandlineprivate/bindings"
)

// HasFuncHasSwitch returns true if the function "WEBEXT.commandLinePrivate.hasSwitch" exists.
func HasFuncHasSwitch() bool {
	return js.True == bindings.HasFuncHasSwitch()
}

// FuncHasSwitch returns the function "WEBEXT.commandLinePrivate.hasSwitch".
func FuncHasSwitch() (fn js.Func[func(name js.String) js.Promise[js.Boolean]]) {
	bindings.FuncHasSwitch(
		js.Pointer(&fn),
	)
	return
}

// HasSwitch calls the function "WEBEXT.commandLinePrivate.hasSwitch" directly.
func HasSwitch(name js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallHasSwitch(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryHasSwitch calls the function "WEBEXT.commandLinePrivate.hasSwitch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasSwitch(name js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasSwitch(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

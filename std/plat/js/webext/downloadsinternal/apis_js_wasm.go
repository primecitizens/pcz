// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package downloadsinternal

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/downloadsinternal/bindings"
)

// HasFuncDetermineFilename returns true if the function "WEBEXT.downloadsInternal.determineFilename" exists.
func HasFuncDetermineFilename() bool {
	return js.True == bindings.HasFuncDetermineFilename()
}

// FuncDetermineFilename returns the function "WEBEXT.downloadsInternal.determineFilename".
func FuncDetermineFilename() (fn js.Func[func(downloadId int32, filename js.String, conflict_action js.String)]) {
	bindings.FuncDetermineFilename(
		js.Pointer(&fn),
	)
	return
}

// DetermineFilename calls the function "WEBEXT.downloadsInternal.determineFilename" directly.
func DetermineFilename(downloadId int32, filename js.String, conflict_action js.String) (ret js.Void) {
	bindings.CallDetermineFilename(
		js.Pointer(&ret),
		int32(downloadId),
		filename.Ref(),
		conflict_action.Ref(),
	)

	return
}

// TryDetermineFilename calls the function "WEBEXT.downloadsInternal.determineFilename"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDetermineFilename(downloadId int32, filename js.String, conflict_action js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDetermineFilename(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(downloadId),
		filename.Ref(),
		conflict_action.Ref(),
	)

	return
}

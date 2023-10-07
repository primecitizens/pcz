// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package appviewguestinternal

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/appviewguestinternal/bindings"
)

// HasFuncAttachFrame returns true if the function "WEBEXT.appViewGuestInternal.attachFrame" exists.
func HasFuncAttachFrame() bool {
	return js.True == bindings.HasFuncAttachFrame()
}

// FuncAttachFrame returns the function "WEBEXT.appViewGuestInternal.attachFrame".
func FuncAttachFrame() (fn js.Func[func(url js.String, guestInstanceId int64) js.Promise[js.BigInt[int64]]]) {
	bindings.FuncAttachFrame(
		js.Pointer(&fn),
	)
	return
}

// AttachFrame calls the function "WEBEXT.appViewGuestInternal.attachFrame" directly.
func AttachFrame(url js.String, guestInstanceId int64) (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallAttachFrame(
		js.Pointer(&ret),
		url.Ref(),
		float64(guestInstanceId),
	)

	return
}

// TryAttachFrame calls the function "WEBEXT.appViewGuestInternal.attachFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAttachFrame(url js.String, guestInstanceId int64) (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAttachFrame(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		float64(guestInstanceId),
	)

	return
}

// HasFuncDenyRequest returns true if the function "WEBEXT.appViewGuestInternal.denyRequest" exists.
func HasFuncDenyRequest() bool {
	return js.True == bindings.HasFuncDenyRequest()
}

// FuncDenyRequest returns the function "WEBEXT.appViewGuestInternal.denyRequest".
func FuncDenyRequest() (fn js.Func[func(guestInstanceId int64)]) {
	bindings.FuncDenyRequest(
		js.Pointer(&fn),
	)
	return
}

// DenyRequest calls the function "WEBEXT.appViewGuestInternal.denyRequest" directly.
func DenyRequest(guestInstanceId int64) (ret js.Void) {
	bindings.CallDenyRequest(
		js.Pointer(&ret),
		float64(guestInstanceId),
	)

	return
}

// TryDenyRequest calls the function "WEBEXT.appViewGuestInternal.denyRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDenyRequest(guestInstanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDenyRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(guestInstanceId),
	)

	return
}

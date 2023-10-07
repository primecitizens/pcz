// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package pagecapture

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/pagecapture/bindings"
)

type SaveAsMHTMLArgDetails struct {
	// TabId is "SaveAsMHTMLArgDetails.tabId"
	//
	// Required
	TabId int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SaveAsMHTMLArgDetails with all fields set.
func (p SaveAsMHTMLArgDetails) FromRef(ref js.Ref) SaveAsMHTMLArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SaveAsMHTMLArgDetails in the application heap.
func (p SaveAsMHTMLArgDetails) New() js.Ref {
	return bindings.SaveAsMHTMLArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SaveAsMHTMLArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SaveAsMHTMLArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SaveAsMHTMLArgDetails) Update(ref js.Ref) {
	bindings.SaveAsMHTMLArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SaveAsMHTMLArgDetails) FreeMembers(recursive bool) {
}

// HasFuncSaveAsMHTML returns true if the function "WEBEXT.pageCapture.saveAsMHTML" exists.
func HasFuncSaveAsMHTML() bool {
	return js.True == bindings.HasFuncSaveAsMHTML()
}

// FuncSaveAsMHTML returns the function "WEBEXT.pageCapture.saveAsMHTML".
func FuncSaveAsMHTML() (fn js.Func[func(details SaveAsMHTMLArgDetails) js.Promise[js.TypedArray[uint8]]]) {
	bindings.FuncSaveAsMHTML(
		js.Pointer(&fn),
	)
	return
}

// SaveAsMHTML calls the function "WEBEXT.pageCapture.saveAsMHTML" directly.
func SaveAsMHTML(details SaveAsMHTMLArgDetails) (ret js.Promise[js.TypedArray[uint8]]) {
	bindings.CallSaveAsMHTML(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySaveAsMHTML calls the function "WEBEXT.pageCapture.saveAsMHTML"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySaveAsMHTML(details SaveAsMHTMLArgDetails) (ret js.Promise[js.TypedArray[uint8]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySaveAsMHTML(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

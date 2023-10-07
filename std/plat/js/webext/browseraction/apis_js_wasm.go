// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package browseraction

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/browseraction/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

type ColorArray = js.Array[int64]

type ImageDataType struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDataType with all fields set.
func (p ImageDataType) FromRef(ref js.Ref) ImageDataType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageDataType in the application heap.
func (p ImageDataType) New() js.Ref {
	return bindings.ImageDataTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImageDataType) UpdateFrom(ref js.Ref) {
	bindings.ImageDataTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDataType) Update(ref js.Ref) {
	bindings.ImageDataTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDataType) FreeMembers(recursive bool) {
}

type OneOf_String_ArrayInt64 struct {
	ref js.Ref
}

func (x OneOf_String_ArrayInt64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayInt64) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayInt64) FromRef(ref js.Ref) OneOf_String_ArrayInt64 {
	return OneOf_String_ArrayInt64{
		ref: ref,
	}
}

func (x OneOf_String_ArrayInt64) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayInt64) ArrayInt64() js.Array[int64] {
	return js.Array[int64]{}.FromRef(x.ref)
}

type SetBadgeBackgroundColorArgDetails struct {
	// Color is "SetBadgeBackgroundColorArgDetails.color"
	//
	// Required
	Color OneOf_String_ArrayInt64
	// TabId is "SetBadgeBackgroundColorArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetBadgeBackgroundColorArgDetails with all fields set.
func (p SetBadgeBackgroundColorArgDetails) FromRef(ref js.Ref) SetBadgeBackgroundColorArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetBadgeBackgroundColorArgDetails in the application heap.
func (p SetBadgeBackgroundColorArgDetails) New() js.Ref {
	return bindings.SetBadgeBackgroundColorArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetBadgeBackgroundColorArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetBadgeBackgroundColorArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetBadgeBackgroundColorArgDetails) Update(ref js.Ref) {
	bindings.SetBadgeBackgroundColorArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetBadgeBackgroundColorArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Color.Ref(),
	)
	p.Color = p.Color.FromRef(js.Undefined)
}

type SetBadgeTextArgDetails struct {
	// TabId is "SetBadgeTextArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Text is "SetBadgeTextArgDetails.text"
	//
	// Optional
	Text js.String

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetBadgeTextArgDetails with all fields set.
func (p SetBadgeTextArgDetails) FromRef(ref js.Ref) SetBadgeTextArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetBadgeTextArgDetails in the application heap.
func (p SetBadgeTextArgDetails) New() js.Ref {
	return bindings.SetBadgeTextArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetBadgeTextArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetBadgeTextArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetBadgeTextArgDetails) Update(ref js.Ref) {
	bindings.SetBadgeTextArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetBadgeTextArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

type OneOf_ImageDataType_Any struct {
	ref js.Ref
}

func (x OneOf_ImageDataType_Any) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ImageDataType_Any) Free() {
	x.ref.Free()
}

func (x OneOf_ImageDataType_Any) FromRef(ref js.Ref) OneOf_ImageDataType_Any {
	return OneOf_ImageDataType_Any{
		ref: ref,
	}
}

func (x OneOf_ImageDataType_Any) ImageDataType() ImageDataType {
	var ret ImageDataType
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_ImageDataType_Any) Any() js.Any {
	return js.Any{}.FromRef(x.ref)
}

type OneOf_String_Any struct {
	ref js.Ref
}

func (x OneOf_String_Any) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Any) Free() {
	x.ref.Free()
}

func (x OneOf_String_Any) FromRef(ref js.Ref) OneOf_String_Any {
	return OneOf_String_Any{
		ref: ref,
	}
}

func (x OneOf_String_Any) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Any) Any() js.Any {
	return js.Any{}.FromRef(x.ref)
}

type SetIconArgDetails struct {
	// ImageData is "SetIconArgDetails.imageData"
	//
	// Optional
	ImageData OneOf_ImageDataType_Any
	// Path is "SetIconArgDetails.path"
	//
	// Optional
	Path OneOf_String_Any
	// TabId is "SetIconArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetIconArgDetails with all fields set.
func (p SetIconArgDetails) FromRef(ref js.Ref) SetIconArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetIconArgDetails in the application heap.
func (p SetIconArgDetails) New() js.Ref {
	return bindings.SetIconArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetIconArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetIconArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetIconArgDetails) Update(ref js.Ref) {
	bindings.SetIconArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetIconArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ImageData.Ref(),
		p.Path.Ref(),
	)
	p.ImageData = p.ImageData.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
}

type SetPopupArgDetails struct {
	// Popup is "SetPopupArgDetails.popup"
	//
	// Required
	Popup js.String
	// TabId is "SetPopupArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetPopupArgDetails with all fields set.
func (p SetPopupArgDetails) FromRef(ref js.Ref) SetPopupArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetPopupArgDetails in the application heap.
func (p SetPopupArgDetails) New() js.Ref {
	return bindings.SetPopupArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetPopupArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetPopupArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetPopupArgDetails) Update(ref js.Ref) {
	bindings.SetPopupArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetPopupArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Popup.Ref(),
	)
	p.Popup = p.Popup.FromRef(js.Undefined)
}

type SetTitleArgDetails struct {
	// TabId is "SetTitleArgDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Title is "SetTitleArgDetails.title"
	//
	// Required
	Title js.String

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetTitleArgDetails with all fields set.
func (p SetTitleArgDetails) FromRef(ref js.Ref) SetTitleArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetTitleArgDetails in the application heap.
func (p SetTitleArgDetails) New() js.Ref {
	return bindings.SetTitleArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetTitleArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetTitleArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetTitleArgDetails) Update(ref js.Ref) {
	bindings.SetTitleArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetTitleArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
}

type TabDetails struct {
	// TabId is "TabDetails.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TabDetails with all fields set.
func (p TabDetails) FromRef(ref js.Ref) TabDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TabDetails in the application heap.
func (p TabDetails) New() js.Ref {
	return bindings.TabDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TabDetails) UpdateFrom(ref js.Ref) {
	bindings.TabDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TabDetails) Update(ref js.Ref) {
	bindings.TabDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TabDetails) FreeMembers(recursive bool) {
}

// HasFuncDisable returns true if the function "WEBEXT.browserAction.disable" exists.
func HasFuncDisable() bool {
	return js.True == bindings.HasFuncDisable()
}

// FuncDisable returns the function "WEBEXT.browserAction.disable".
func FuncDisable() (fn js.Func[func(tabId int64) js.Promise[js.Void]]) {
	bindings.FuncDisable(
		js.Pointer(&fn),
	)
	return
}

// Disable calls the function "WEBEXT.browserAction.disable" directly.
func Disable(tabId int64) (ret js.Promise[js.Void]) {
	bindings.CallDisable(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryDisable calls the function "WEBEXT.browserAction.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisable(tabId int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisable(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncEnable returns true if the function "WEBEXT.browserAction.enable" exists.
func HasFuncEnable() bool {
	return js.True == bindings.HasFuncEnable()
}

// FuncEnable returns the function "WEBEXT.browserAction.enable".
func FuncEnable() (fn js.Func[func(tabId int64) js.Promise[js.Void]]) {
	bindings.FuncEnable(
		js.Pointer(&fn),
	)
	return
}

// Enable calls the function "WEBEXT.browserAction.enable" directly.
func Enable(tabId int64) (ret js.Promise[js.Void]) {
	bindings.CallEnable(
		js.Pointer(&ret),
		float64(tabId),
	)

	return
}

// TryEnable calls the function "WEBEXT.browserAction.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnable(tabId int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnable(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(tabId),
	)

	return
}

// HasFuncGetBadgeBackgroundColor returns true if the function "WEBEXT.browserAction.getBadgeBackgroundColor" exists.
func HasFuncGetBadgeBackgroundColor() bool {
	return js.True == bindings.HasFuncGetBadgeBackgroundColor()
}

// FuncGetBadgeBackgroundColor returns the function "WEBEXT.browserAction.getBadgeBackgroundColor".
func FuncGetBadgeBackgroundColor() (fn js.Func[func(details TabDetails) js.Promise[ColorArray]]) {
	bindings.FuncGetBadgeBackgroundColor(
		js.Pointer(&fn),
	)
	return
}

// GetBadgeBackgroundColor calls the function "WEBEXT.browserAction.getBadgeBackgroundColor" directly.
func GetBadgeBackgroundColor(details TabDetails) (ret js.Promise[ColorArray]) {
	bindings.CallGetBadgeBackgroundColor(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetBadgeBackgroundColor calls the function "WEBEXT.browserAction.getBadgeBackgroundColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBadgeBackgroundColor(details TabDetails) (ret js.Promise[ColorArray], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBadgeBackgroundColor(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetBadgeText returns true if the function "WEBEXT.browserAction.getBadgeText" exists.
func HasFuncGetBadgeText() bool {
	return js.True == bindings.HasFuncGetBadgeText()
}

// FuncGetBadgeText returns the function "WEBEXT.browserAction.getBadgeText".
func FuncGetBadgeText() (fn js.Func[func(details TabDetails) js.Promise[js.String]]) {
	bindings.FuncGetBadgeText(
		js.Pointer(&fn),
	)
	return
}

// GetBadgeText calls the function "WEBEXT.browserAction.getBadgeText" directly.
func GetBadgeText(details TabDetails) (ret js.Promise[js.String]) {
	bindings.CallGetBadgeText(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetBadgeText calls the function "WEBEXT.browserAction.getBadgeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBadgeText(details TabDetails) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBadgeText(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetPopup returns true if the function "WEBEXT.browserAction.getPopup" exists.
func HasFuncGetPopup() bool {
	return js.True == bindings.HasFuncGetPopup()
}

// FuncGetPopup returns the function "WEBEXT.browserAction.getPopup".
func FuncGetPopup() (fn js.Func[func(details TabDetails) js.Promise[js.String]]) {
	bindings.FuncGetPopup(
		js.Pointer(&fn),
	)
	return
}

// GetPopup calls the function "WEBEXT.browserAction.getPopup" directly.
func GetPopup(details TabDetails) (ret js.Promise[js.String]) {
	bindings.CallGetPopup(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetPopup calls the function "WEBEXT.browserAction.getPopup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPopup(details TabDetails) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPopup(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetTitle returns true if the function "WEBEXT.browserAction.getTitle" exists.
func HasFuncGetTitle() bool {
	return js.True == bindings.HasFuncGetTitle()
}

// FuncGetTitle returns the function "WEBEXT.browserAction.getTitle".
func FuncGetTitle() (fn js.Func[func(details TabDetails) js.Promise[js.String]]) {
	bindings.FuncGetTitle(
		js.Pointer(&fn),
	)
	return
}

// GetTitle calls the function "WEBEXT.browserAction.getTitle" directly.
func GetTitle(details TabDetails) (ret js.Promise[js.String]) {
	bindings.CallGetTitle(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetTitle calls the function "WEBEXT.browserAction.getTitle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTitle(details TabDetails) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTitle(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

type OnClickedEventCallbackFunc func(this js.Ref, tab *tabs.Tab) js.Ref

func (fn OnClickedEventCallbackFunc) Register() js.Func[func(tab *tabs.Tab)] {
	return js.RegisterCallback[func(tab *tabs.Tab)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClickedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 tabs.Tab
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tab *tabs.Tab) js.Ref
	Arg T
}

func (cb *OnClickedEventCallback[T]) Register() js.Func[func(tab *tabs.Tab)] {
	return js.RegisterCallback[func(tab *tabs.Tab)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClickedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 tabs.Tab
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnClicked returns true if the function "WEBEXT.browserAction.onClicked.addListener" exists.
func HasFuncOnClicked() bool {
	return js.True == bindings.HasFuncOnClicked()
}

// FuncOnClicked returns the function "WEBEXT.browserAction.onClicked.addListener".
func FuncOnClicked() (fn js.Func[func(callback js.Func[func(tab *tabs.Tab)])]) {
	bindings.FuncOnClicked(
		js.Pointer(&fn),
	)
	return
}

// OnClicked calls the function "WEBEXT.browserAction.onClicked.addListener" directly.
func OnClicked(callback js.Func[func(tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClicked calls the function "WEBEXT.browserAction.onClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClicked(callback js.Func[func(tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClicked returns true if the function "WEBEXT.browserAction.onClicked.removeListener" exists.
func HasFuncOffClicked() bool {
	return js.True == bindings.HasFuncOffClicked()
}

// FuncOffClicked returns the function "WEBEXT.browserAction.onClicked.removeListener".
func FuncOffClicked() (fn js.Func[func(callback js.Func[func(tab *tabs.Tab)])]) {
	bindings.FuncOffClicked(
		js.Pointer(&fn),
	)
	return
}

// OffClicked calls the function "WEBEXT.browserAction.onClicked.removeListener" directly.
func OffClicked(callback js.Func[func(tab *tabs.Tab)]) (ret js.Void) {
	bindings.CallOffClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClicked calls the function "WEBEXT.browserAction.onClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClicked(callback js.Func[func(tab *tabs.Tab)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClicked returns true if the function "WEBEXT.browserAction.onClicked.hasListener" exists.
func HasFuncHasOnClicked() bool {
	return js.True == bindings.HasFuncHasOnClicked()
}

// FuncHasOnClicked returns the function "WEBEXT.browserAction.onClicked.hasListener".
func FuncHasOnClicked() (fn js.Func[func(callback js.Func[func(tab *tabs.Tab)]) bool]) {
	bindings.FuncHasOnClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnClicked calls the function "WEBEXT.browserAction.onClicked.hasListener" directly.
func HasOnClicked(callback js.Func[func(tab *tabs.Tab)]) (ret bool) {
	bindings.CallHasOnClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClicked calls the function "WEBEXT.browserAction.onClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClicked(callback js.Func[func(tab *tabs.Tab)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenPopup returns true if the function "WEBEXT.browserAction.openPopup" exists.
func HasFuncOpenPopup() bool {
	return js.True == bindings.HasFuncOpenPopup()
}

// FuncOpenPopup returns the function "WEBEXT.browserAction.openPopup".
func FuncOpenPopup() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncOpenPopup(
		js.Pointer(&fn),
	)
	return
}

// OpenPopup calls the function "WEBEXT.browserAction.openPopup" directly.
func OpenPopup() (ret js.Promise[js.Any]) {
	bindings.CallOpenPopup(
		js.Pointer(&ret),
	)

	return
}

// TryOpenPopup calls the function "WEBEXT.browserAction.openPopup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenPopup() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenPopup(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetBadgeBackgroundColor returns true if the function "WEBEXT.browserAction.setBadgeBackgroundColor" exists.
func HasFuncSetBadgeBackgroundColor() bool {
	return js.True == bindings.HasFuncSetBadgeBackgroundColor()
}

// FuncSetBadgeBackgroundColor returns the function "WEBEXT.browserAction.setBadgeBackgroundColor".
func FuncSetBadgeBackgroundColor() (fn js.Func[func(details SetBadgeBackgroundColorArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetBadgeBackgroundColor(
		js.Pointer(&fn),
	)
	return
}

// SetBadgeBackgroundColor calls the function "WEBEXT.browserAction.setBadgeBackgroundColor" directly.
func SetBadgeBackgroundColor(details SetBadgeBackgroundColorArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetBadgeBackgroundColor(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetBadgeBackgroundColor calls the function "WEBEXT.browserAction.setBadgeBackgroundColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetBadgeBackgroundColor(details SetBadgeBackgroundColorArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetBadgeBackgroundColor(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetBadgeText returns true if the function "WEBEXT.browserAction.setBadgeText" exists.
func HasFuncSetBadgeText() bool {
	return js.True == bindings.HasFuncSetBadgeText()
}

// FuncSetBadgeText returns the function "WEBEXT.browserAction.setBadgeText".
func FuncSetBadgeText() (fn js.Func[func(details SetBadgeTextArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetBadgeText(
		js.Pointer(&fn),
	)
	return
}

// SetBadgeText calls the function "WEBEXT.browserAction.setBadgeText" directly.
func SetBadgeText(details SetBadgeTextArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetBadgeText(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetBadgeText calls the function "WEBEXT.browserAction.setBadgeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetBadgeText(details SetBadgeTextArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetBadgeText(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetIcon returns true if the function "WEBEXT.browserAction.setIcon" exists.
func HasFuncSetIcon() bool {
	return js.True == bindings.HasFuncSetIcon()
}

// FuncSetIcon returns the function "WEBEXT.browserAction.setIcon".
func FuncSetIcon() (fn js.Func[func(details SetIconArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetIcon(
		js.Pointer(&fn),
	)
	return
}

// SetIcon calls the function "WEBEXT.browserAction.setIcon" directly.
func SetIcon(details SetIconArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetIcon(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetIcon calls the function "WEBEXT.browserAction.setIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetIcon(details SetIconArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetIcon(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetPopup returns true if the function "WEBEXT.browserAction.setPopup" exists.
func HasFuncSetPopup() bool {
	return js.True == bindings.HasFuncSetPopup()
}

// FuncSetPopup returns the function "WEBEXT.browserAction.setPopup".
func FuncSetPopup() (fn js.Func[func(details SetPopupArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetPopup(
		js.Pointer(&fn),
	)
	return
}

// SetPopup calls the function "WEBEXT.browserAction.setPopup" directly.
func SetPopup(details SetPopupArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetPopup(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetPopup calls the function "WEBEXT.browserAction.setPopup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPopup(details SetPopupArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPopup(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetTitle returns true if the function "WEBEXT.browserAction.setTitle" exists.
func HasFuncSetTitle() bool {
	return js.True == bindings.HasFuncSetTitle()
}

// FuncSetTitle returns the function "WEBEXT.browserAction.setTitle".
func FuncSetTitle() (fn js.Func[func(details SetTitleArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetTitle(
		js.Pointer(&fn),
	)
	return
}

// SetTitle calls the function "WEBEXT.browserAction.setTitle" directly.
func SetTitle(details SetTitleArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetTitle(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetTitle calls the function "WEBEXT.browserAction.setTitle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTitle(details SetTitleArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTitle(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

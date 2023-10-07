// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package offscreen

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/offscreen/bindings"
)

type BooleanCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn BooleanCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BooleanCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BooleanCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *BooleanCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BooleanCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Reason uint32

const (
	_ Reason = iota

	Reason_TESTING
	Reason_AUDIO_PLAYBACK
	Reason_IFRAME_SCRIPTING
	Reason_DOM_SCRAPING
	Reason_BLOBS
	Reason_DOM_PARSER
	Reason_USER_MEDIA
	Reason_DISPLAY_MEDIA
	Reason_WEB_RTC
	Reason_CLIPBOARD
	Reason_LOCAL_STORAGE
	Reason_WORKERS
	Reason_BATTERY_STATUS
	Reason_MATCH_MEDIA
	Reason_GEOLOCATION
)

func (Reason) FromRef(str js.Ref) Reason {
	return Reason(bindings.ConstOfReason(str))
}

func (x Reason) String() (string, bool) {
	switch x {
	case Reason_TESTING:
		return "TESTING", true
	case Reason_AUDIO_PLAYBACK:
		return "AUDIO_PLAYBACK", true
	case Reason_IFRAME_SCRIPTING:
		return "IFRAME_SCRIPTING", true
	case Reason_DOM_SCRAPING:
		return "DOM_SCRAPING", true
	case Reason_BLOBS:
		return "BLOBS", true
	case Reason_DOM_PARSER:
		return "DOM_PARSER", true
	case Reason_USER_MEDIA:
		return "USER_MEDIA", true
	case Reason_DISPLAY_MEDIA:
		return "DISPLAY_MEDIA", true
	case Reason_WEB_RTC:
		return "WEB_RTC", true
	case Reason_CLIPBOARD:
		return "CLIPBOARD", true
	case Reason_LOCAL_STORAGE:
		return "LOCAL_STORAGE", true
	case Reason_WORKERS:
		return "WORKERS", true
	case Reason_BATTERY_STATUS:
		return "BATTERY_STATUS", true
	case Reason_MATCH_MEDIA:
		return "MATCH_MEDIA", true
	case Reason_GEOLOCATION:
		return "GEOLOCATION", true
	default:
		return "", false
	}
}

type CreateParameters struct {
	// Reasons is "CreateParameters.reasons"
	//
	// Optional
	Reasons js.Array[Reason]
	// Url is "CreateParameters.url"
	//
	// Optional
	Url js.String
	// Justification is "CreateParameters.justification"
	//
	// Optional
	Justification js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateParameters with all fields set.
func (p CreateParameters) FromRef(ref js.Ref) CreateParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateParameters in the application heap.
func (p CreateParameters) New() js.Ref {
	return bindings.CreateParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateParameters) UpdateFrom(ref js.Ref) {
	bindings.CreateParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateParameters) Update(ref js.Ref) {
	bindings.CreateParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Reasons.Ref(),
		p.Url.Ref(),
		p.Justification.Ref(),
	)
	p.Reasons = p.Reasons.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Justification = p.Justification.FromRef(js.Undefined)
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncCloseDocument returns true if the function "WEBEXT.offscreen.closeDocument" exists.
func HasFuncCloseDocument() bool {
	return js.True == bindings.HasFuncCloseDocument()
}

// FuncCloseDocument returns the function "WEBEXT.offscreen.closeDocument".
func FuncCloseDocument() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncCloseDocument(
		js.Pointer(&fn),
	)
	return
}

// CloseDocument calls the function "WEBEXT.offscreen.closeDocument" directly.
func CloseDocument() (ret js.Promise[js.Void]) {
	bindings.CallCloseDocument(
		js.Pointer(&ret),
	)

	return
}

// TryCloseDocument calls the function "WEBEXT.offscreen.closeDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseDocument() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseDocument(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateDocument returns true if the function "WEBEXT.offscreen.createDocument" exists.
func HasFuncCreateDocument() bool {
	return js.True == bindings.HasFuncCreateDocument()
}

// FuncCreateDocument returns the function "WEBEXT.offscreen.createDocument".
func FuncCreateDocument() (fn js.Func[func(parameters CreateParameters) js.Promise[js.Void]]) {
	bindings.FuncCreateDocument(
		js.Pointer(&fn),
	)
	return
}

// CreateDocument calls the function "WEBEXT.offscreen.createDocument" directly.
func CreateDocument(parameters CreateParameters) (ret js.Promise[js.Void]) {
	bindings.CallCreateDocument(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryCreateDocument calls the function "WEBEXT.offscreen.createDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateDocument(parameters CreateParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateDocument(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncHasDocument returns true if the function "WEBEXT.offscreen.hasDocument" exists.
func HasFuncHasDocument() bool {
	return js.True == bindings.HasFuncHasDocument()
}

// FuncHasDocument returns the function "WEBEXT.offscreen.hasDocument".
func FuncHasDocument() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncHasDocument(
		js.Pointer(&fn),
	)
	return
}

// HasDocument calls the function "WEBEXT.offscreen.hasDocument" directly.
func HasDocument() (ret js.Promise[js.Boolean]) {
	bindings.CallHasDocument(
		js.Pointer(&ret),
	)

	return
}

// TryHasDocument calls the function "WEBEXT.offscreen.hasDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasDocument() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasDocument(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

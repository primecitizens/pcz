// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package omnibox

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/omnibox/bindings"
)

type DescriptionStyleType uint32

const (
	_ DescriptionStyleType = iota

	DescriptionStyleType_URL
	DescriptionStyleType_MATCH
	DescriptionStyleType_DIM
)

func (DescriptionStyleType) FromRef(str js.Ref) DescriptionStyleType {
	return DescriptionStyleType(bindings.ConstOfDescriptionStyleType(str))
}

func (x DescriptionStyleType) String() (string, bool) {
	switch x {
	case DescriptionStyleType_URL:
		return "url", true
	case DescriptionStyleType_MATCH:
		return "match", true
	case DescriptionStyleType_DIM:
		return "dim", true
	default:
		return "", false
	}
}

type MatchClassification struct {
	// Length is "MatchClassification.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length int64
	// Offset is "MatchClassification.offset"
	//
	// Required
	Offset int64
	// Type is "MatchClassification.type"
	//
	// Required
	Type DescriptionStyleType

	FFI_USE_Length bool // for Length.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchClassification with all fields set.
func (p MatchClassification) FromRef(ref js.Ref) MatchClassification {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchClassification in the application heap.
func (p MatchClassification) New() js.Ref {
	return bindings.MatchClassificationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchClassification) UpdateFrom(ref js.Ref) {
	bindings.MatchClassificationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchClassification) Update(ref js.Ref) {
	bindings.MatchClassificationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchClassification) FreeMembers(recursive bool) {
}

type DefaultSuggestResult struct {
	// Description is "DefaultSuggestResult.description"
	//
	// Required
	Description js.String
	// DescriptionStyles is "DefaultSuggestResult.descriptionStyles"
	//
	// Optional
	DescriptionStyles js.Array[MatchClassification]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DefaultSuggestResult with all fields set.
func (p DefaultSuggestResult) FromRef(ref js.Ref) DefaultSuggestResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DefaultSuggestResult in the application heap.
func (p DefaultSuggestResult) New() js.Ref {
	return bindings.DefaultSuggestResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DefaultSuggestResult) UpdateFrom(ref js.Ref) {
	bindings.DefaultSuggestResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DefaultSuggestResult) Update(ref js.Ref) {
	bindings.DefaultSuggestResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DefaultSuggestResult) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.DescriptionStyles.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.DescriptionStyles = p.DescriptionStyles.FromRef(js.Undefined)
}

type OnInputChangedArgSuggestFunc func(this js.Ref, suggestResults js.Array[SuggestResult]) js.Ref

func (fn OnInputChangedArgSuggestFunc) Register() js.Func[func(suggestResults js.Array[SuggestResult])] {
	return js.RegisterCallback[func(suggestResults js.Array[SuggestResult])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputChangedArgSuggestFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SuggestResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnInputChangedArgSuggest[T any] struct {
	Fn  func(arg T, this js.Ref, suggestResults js.Array[SuggestResult]) js.Ref
	Arg T
}

func (cb *OnInputChangedArgSuggest[T]) Register() js.Func[func(suggestResults js.Array[SuggestResult])] {
	return js.RegisterCallback[func(suggestResults js.Array[SuggestResult])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputChangedArgSuggest[T]) DispatchCallback(
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

		js.Array[SuggestResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SuggestResult struct {
	// Content is "SuggestResult.content"
	//
	// Required
	Content js.String
	// Deletable is "SuggestResult.deletable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Deletable MUST be set to true to make this field effective.
	Deletable bool
	// Description is "SuggestResult.description"
	//
	// Required
	Description js.String
	// DescriptionStyles is "SuggestResult.descriptionStyles"
	//
	// Optional
	DescriptionStyles js.Array[MatchClassification]

	FFI_USE_Deletable bool // for Deletable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SuggestResult with all fields set.
func (p SuggestResult) FromRef(ref js.Ref) SuggestResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SuggestResult in the application heap.
func (p SuggestResult) New() js.Ref {
	return bindings.SuggestResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SuggestResult) UpdateFrom(ref js.Ref) {
	bindings.SuggestResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SuggestResult) Update(ref js.Ref) {
	bindings.SuggestResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SuggestResult) FreeMembers(recursive bool) {
	js.Free(
		p.Content.Ref(),
		p.Description.Ref(),
		p.DescriptionStyles.Ref(),
	)
	p.Content = p.Content.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.DescriptionStyles = p.DescriptionStyles.FromRef(js.Undefined)
}

type OnInputEnteredDisposition uint32

const (
	_ OnInputEnteredDisposition = iota

	OnInputEnteredDisposition_CURRENT_TAB
	OnInputEnteredDisposition_NEW_FOREGROUND_TAB
	OnInputEnteredDisposition_NEW_BACKGROUND_TAB
)

func (OnInputEnteredDisposition) FromRef(str js.Ref) OnInputEnteredDisposition {
	return OnInputEnteredDisposition(bindings.ConstOfOnInputEnteredDisposition(str))
}

func (x OnInputEnteredDisposition) String() (string, bool) {
	switch x {
	case OnInputEnteredDisposition_CURRENT_TAB:
		return "currentTab", true
	case OnInputEnteredDisposition_NEW_FOREGROUND_TAB:
		return "newForegroundTab", true
	case OnInputEnteredDisposition_NEW_BACKGROUND_TAB:
		return "newBackgroundTab", true
	default:
		return "", false
	}
}

type OnDeleteSuggestionEventCallbackFunc func(this js.Ref, text js.String) js.Ref

func (fn OnDeleteSuggestionEventCallbackFunc) Register() js.Func[func(text js.String)] {
	return js.RegisterCallback[func(text js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeleteSuggestionEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDeleteSuggestionEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, text js.String) js.Ref
	Arg T
}

func (cb *OnDeleteSuggestionEventCallback[T]) Register() js.Func[func(text js.String)] {
	return js.RegisterCallback[func(text js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeleteSuggestionEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDeleteSuggestion returns true if the function "WEBEXT.omnibox.onDeleteSuggestion.addListener" exists.
func HasFuncOnDeleteSuggestion() bool {
	return js.True == bindings.HasFuncOnDeleteSuggestion()
}

// FuncOnDeleteSuggestion returns the function "WEBEXT.omnibox.onDeleteSuggestion.addListener".
func FuncOnDeleteSuggestion() (fn js.Func[func(callback js.Func[func(text js.String)])]) {
	bindings.FuncOnDeleteSuggestion(
		js.Pointer(&fn),
	)
	return
}

// OnDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.addListener" directly.
func OnDeleteSuggestion(callback js.Func[func(text js.String)]) (ret js.Void) {
	bindings.CallOnDeleteSuggestion(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeleteSuggestion(callback js.Func[func(text js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeleteSuggestion(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeleteSuggestion returns true if the function "WEBEXT.omnibox.onDeleteSuggestion.removeListener" exists.
func HasFuncOffDeleteSuggestion() bool {
	return js.True == bindings.HasFuncOffDeleteSuggestion()
}

// FuncOffDeleteSuggestion returns the function "WEBEXT.omnibox.onDeleteSuggestion.removeListener".
func FuncOffDeleteSuggestion() (fn js.Func[func(callback js.Func[func(text js.String)])]) {
	bindings.FuncOffDeleteSuggestion(
		js.Pointer(&fn),
	)
	return
}

// OffDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.removeListener" directly.
func OffDeleteSuggestion(callback js.Func[func(text js.String)]) (ret js.Void) {
	bindings.CallOffDeleteSuggestion(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeleteSuggestion(callback js.Func[func(text js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeleteSuggestion(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeleteSuggestion returns true if the function "WEBEXT.omnibox.onDeleteSuggestion.hasListener" exists.
func HasFuncHasOnDeleteSuggestion() bool {
	return js.True == bindings.HasFuncHasOnDeleteSuggestion()
}

// FuncHasOnDeleteSuggestion returns the function "WEBEXT.omnibox.onDeleteSuggestion.hasListener".
func FuncHasOnDeleteSuggestion() (fn js.Func[func(callback js.Func[func(text js.String)]) bool]) {
	bindings.FuncHasOnDeleteSuggestion(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.hasListener" directly.
func HasOnDeleteSuggestion(callback js.Func[func(text js.String)]) (ret bool) {
	bindings.CallHasOnDeleteSuggestion(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeleteSuggestion calls the function "WEBEXT.omnibox.onDeleteSuggestion.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeleteSuggestion(callback js.Func[func(text js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeleteSuggestion(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputCancelledEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnInputCancelledEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputCancelledEventCallbackFunc) DispatchCallback(
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

type OnInputCancelledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnInputCancelledEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputCancelledEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputCancelled returns true if the function "WEBEXT.omnibox.onInputCancelled.addListener" exists.
func HasFuncOnInputCancelled() bool {
	return js.True == bindings.HasFuncOnInputCancelled()
}

// FuncOnInputCancelled returns the function "WEBEXT.omnibox.onInputCancelled.addListener".
func FuncOnInputCancelled() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnInputCancelled(
		js.Pointer(&fn),
	)
	return
}

// OnInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.addListener" directly.
func OnInputCancelled(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnInputCancelled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputCancelled(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputCancelled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputCancelled returns true if the function "WEBEXT.omnibox.onInputCancelled.removeListener" exists.
func HasFuncOffInputCancelled() bool {
	return js.True == bindings.HasFuncOffInputCancelled()
}

// FuncOffInputCancelled returns the function "WEBEXT.omnibox.onInputCancelled.removeListener".
func FuncOffInputCancelled() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffInputCancelled(
		js.Pointer(&fn),
	)
	return
}

// OffInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.removeListener" directly.
func OffInputCancelled(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffInputCancelled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputCancelled(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputCancelled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputCancelled returns true if the function "WEBEXT.omnibox.onInputCancelled.hasListener" exists.
func HasFuncHasOnInputCancelled() bool {
	return js.True == bindings.HasFuncHasOnInputCancelled()
}

// FuncHasOnInputCancelled returns the function "WEBEXT.omnibox.onInputCancelled.hasListener".
func FuncHasOnInputCancelled() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnInputCancelled(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.hasListener" directly.
func HasOnInputCancelled(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnInputCancelled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputCancelled calls the function "WEBEXT.omnibox.onInputCancelled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputCancelled(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputCancelled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputChangedEventCallbackFunc func(this js.Ref, text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])]) js.Ref

func (fn OnInputChangedEventCallbackFunc) Register() js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])] {
	return js.RegisterCallback[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Func[func(suggestResults js.Array[SuggestResult])]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnInputChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])]) js.Ref
	Arg T
}

func (cb *OnInputChangedEventCallback[T]) Register() js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])] {
	return js.RegisterCallback[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.Func[func(suggestResults js.Array[SuggestResult])]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnInputChanged returns true if the function "WEBEXT.omnibox.onInputChanged.addListener" exists.
func HasFuncOnInputChanged() bool {
	return js.True == bindings.HasFuncOnInputChanged()
}

// FuncOnInputChanged returns the function "WEBEXT.omnibox.onInputChanged.addListener".
func FuncOnInputChanged() (fn js.Func[func(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])])]) {
	bindings.FuncOnInputChanged(
		js.Pointer(&fn),
	)
	return
}

// OnInputChanged calls the function "WEBEXT.omnibox.onInputChanged.addListener" directly.
func OnInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret js.Void) {
	bindings.CallOnInputChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputChanged calls the function "WEBEXT.omnibox.onInputChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputChanged returns true if the function "WEBEXT.omnibox.onInputChanged.removeListener" exists.
func HasFuncOffInputChanged() bool {
	return js.True == bindings.HasFuncOffInputChanged()
}

// FuncOffInputChanged returns the function "WEBEXT.omnibox.onInputChanged.removeListener".
func FuncOffInputChanged() (fn js.Func[func(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])])]) {
	bindings.FuncOffInputChanged(
		js.Pointer(&fn),
	)
	return
}

// OffInputChanged calls the function "WEBEXT.omnibox.onInputChanged.removeListener" directly.
func OffInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret js.Void) {
	bindings.CallOffInputChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputChanged calls the function "WEBEXT.omnibox.onInputChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputChanged returns true if the function "WEBEXT.omnibox.onInputChanged.hasListener" exists.
func HasFuncHasOnInputChanged() bool {
	return js.True == bindings.HasFuncHasOnInputChanged()
}

// FuncHasOnInputChanged returns the function "WEBEXT.omnibox.onInputChanged.hasListener".
func FuncHasOnInputChanged() (fn js.Func[func(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) bool]) {
	bindings.FuncHasOnInputChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputChanged calls the function "WEBEXT.omnibox.onInputChanged.hasListener" directly.
func HasOnInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret bool) {
	bindings.CallHasOnInputChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputChanged calls the function "WEBEXT.omnibox.onInputChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputChanged(callback js.Func[func(text js.String, suggest js.Func[func(suggestResults js.Array[SuggestResult])])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputEnteredEventCallbackFunc func(this js.Ref, text js.String, disposition OnInputEnteredDisposition) js.Ref

func (fn OnInputEnteredEventCallbackFunc) Register() js.Func[func(text js.String, disposition OnInputEnteredDisposition)] {
	return js.RegisterCallback[func(text js.String, disposition OnInputEnteredDisposition)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputEnteredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		OnInputEnteredDisposition(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnInputEnteredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, text js.String, disposition OnInputEnteredDisposition) js.Ref
	Arg T
}

func (cb *OnInputEnteredEventCallback[T]) Register() js.Func[func(text js.String, disposition OnInputEnteredDisposition)] {
	return js.RegisterCallback[func(text js.String, disposition OnInputEnteredDisposition)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputEnteredEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		OnInputEnteredDisposition(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnInputEntered returns true if the function "WEBEXT.omnibox.onInputEntered.addListener" exists.
func HasFuncOnInputEntered() bool {
	return js.True == bindings.HasFuncOnInputEntered()
}

// FuncOnInputEntered returns the function "WEBEXT.omnibox.onInputEntered.addListener".
func FuncOnInputEntered() (fn js.Func[func(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)])]) {
	bindings.FuncOnInputEntered(
		js.Pointer(&fn),
	)
	return
}

// OnInputEntered calls the function "WEBEXT.omnibox.onInputEntered.addListener" directly.
func OnInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret js.Void) {
	bindings.CallOnInputEntered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputEntered calls the function "WEBEXT.omnibox.onInputEntered.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputEntered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputEntered returns true if the function "WEBEXT.omnibox.onInputEntered.removeListener" exists.
func HasFuncOffInputEntered() bool {
	return js.True == bindings.HasFuncOffInputEntered()
}

// FuncOffInputEntered returns the function "WEBEXT.omnibox.onInputEntered.removeListener".
func FuncOffInputEntered() (fn js.Func[func(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)])]) {
	bindings.FuncOffInputEntered(
		js.Pointer(&fn),
	)
	return
}

// OffInputEntered calls the function "WEBEXT.omnibox.onInputEntered.removeListener" directly.
func OffInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret js.Void) {
	bindings.CallOffInputEntered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputEntered calls the function "WEBEXT.omnibox.onInputEntered.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputEntered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputEntered returns true if the function "WEBEXT.omnibox.onInputEntered.hasListener" exists.
func HasFuncHasOnInputEntered() bool {
	return js.True == bindings.HasFuncHasOnInputEntered()
}

// FuncHasOnInputEntered returns the function "WEBEXT.omnibox.onInputEntered.hasListener".
func FuncHasOnInputEntered() (fn js.Func[func(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) bool]) {
	bindings.FuncHasOnInputEntered(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputEntered calls the function "WEBEXT.omnibox.onInputEntered.hasListener" directly.
func HasOnInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret bool) {
	bindings.CallHasOnInputEntered(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputEntered calls the function "WEBEXT.omnibox.onInputEntered.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputEntered(callback js.Func[func(text js.String, disposition OnInputEnteredDisposition)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputEntered(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputStartedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnInputStartedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputStartedEventCallbackFunc) DispatchCallback(
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

type OnInputStartedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnInputStartedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputStartedEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputStarted returns true if the function "WEBEXT.omnibox.onInputStarted.addListener" exists.
func HasFuncOnInputStarted() bool {
	return js.True == bindings.HasFuncOnInputStarted()
}

// FuncOnInputStarted returns the function "WEBEXT.omnibox.onInputStarted.addListener".
func FuncOnInputStarted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnInputStarted(
		js.Pointer(&fn),
	)
	return
}

// OnInputStarted calls the function "WEBEXT.omnibox.onInputStarted.addListener" directly.
func OnInputStarted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnInputStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputStarted calls the function "WEBEXT.omnibox.onInputStarted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputStarted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputStarted returns true if the function "WEBEXT.omnibox.onInputStarted.removeListener" exists.
func HasFuncOffInputStarted() bool {
	return js.True == bindings.HasFuncOffInputStarted()
}

// FuncOffInputStarted returns the function "WEBEXT.omnibox.onInputStarted.removeListener".
func FuncOffInputStarted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffInputStarted(
		js.Pointer(&fn),
	)
	return
}

// OffInputStarted calls the function "WEBEXT.omnibox.onInputStarted.removeListener" directly.
func OffInputStarted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffInputStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputStarted calls the function "WEBEXT.omnibox.onInputStarted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputStarted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputStarted returns true if the function "WEBEXT.omnibox.onInputStarted.hasListener" exists.
func HasFuncHasOnInputStarted() bool {
	return js.True == bindings.HasFuncHasOnInputStarted()
}

// FuncHasOnInputStarted returns the function "WEBEXT.omnibox.onInputStarted.hasListener".
func FuncHasOnInputStarted() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnInputStarted(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputStarted calls the function "WEBEXT.omnibox.onInputStarted.hasListener" directly.
func HasOnInputStarted(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnInputStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputStarted calls the function "WEBEXT.omnibox.onInputStarted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputStarted(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendSuggestions returns true if the function "WEBEXT.omnibox.sendSuggestions" exists.
func HasFuncSendSuggestions() bool {
	return js.True == bindings.HasFuncSendSuggestions()
}

// FuncSendSuggestions returns the function "WEBEXT.omnibox.sendSuggestions".
func FuncSendSuggestions() (fn js.Func[func(requestId int64, suggestResults js.Array[SuggestResult])]) {
	bindings.FuncSendSuggestions(
		js.Pointer(&fn),
	)
	return
}

// SendSuggestions calls the function "WEBEXT.omnibox.sendSuggestions" directly.
func SendSuggestions(requestId int64, suggestResults js.Array[SuggestResult]) (ret js.Void) {
	bindings.CallSendSuggestions(
		js.Pointer(&ret),
		float64(requestId),
		suggestResults.Ref(),
	)

	return
}

// TrySendSuggestions calls the function "WEBEXT.omnibox.sendSuggestions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendSuggestions(requestId int64, suggestResults js.Array[SuggestResult]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendSuggestions(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(requestId),
		suggestResults.Ref(),
	)

	return
}

// HasFuncSetDefaultSuggestion returns true if the function "WEBEXT.omnibox.setDefaultSuggestion" exists.
func HasFuncSetDefaultSuggestion() bool {
	return js.True == bindings.HasFuncSetDefaultSuggestion()
}

// FuncSetDefaultSuggestion returns the function "WEBEXT.omnibox.setDefaultSuggestion".
func FuncSetDefaultSuggestion() (fn js.Func[func(suggestion DefaultSuggestResult) js.Promise[js.Void]]) {
	bindings.FuncSetDefaultSuggestion(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultSuggestion calls the function "WEBEXT.omnibox.setDefaultSuggestion" directly.
func SetDefaultSuggestion(suggestion DefaultSuggestResult) (ret js.Promise[js.Void]) {
	bindings.CallSetDefaultSuggestion(
		js.Pointer(&ret),
		js.Pointer(&suggestion),
	)

	return
}

// TrySetDefaultSuggestion calls the function "WEBEXT.omnibox.setDefaultSuggestion"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultSuggestion(suggestion DefaultSuggestResult) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultSuggestion(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&suggestion),
	)

	return
}

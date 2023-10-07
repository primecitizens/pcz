// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package languagesettingsprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/languagesettingsprivate/bindings"
)

type GetAlwaysTranslateLanguagesCallbackFunc func(this js.Ref, languageCodes js.Array[js.String]) js.Ref

func (fn GetAlwaysTranslateLanguagesCallbackFunc) Register() js.Func[func(languageCodes js.Array[js.String])] {
	return js.RegisterCallback[func(languageCodes js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAlwaysTranslateLanguagesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAlwaysTranslateLanguagesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, languageCodes js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetAlwaysTranslateLanguagesCallback[T]) Register() js.Func[func(languageCodes js.Array[js.String])] {
	return js.RegisterCallback[func(languageCodes js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAlwaysTranslateLanguagesCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetInputMethodListsCallbackFunc func(this js.Ref, lists *InputMethodLists) js.Ref

func (fn GetInputMethodListsCallbackFunc) Register() js.Func[func(lists *InputMethodLists)] {
	return js.RegisterCallback[func(lists *InputMethodLists)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetInputMethodListsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InputMethodLists
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

type GetInputMethodListsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, lists *InputMethodLists) js.Ref
	Arg T
}

func (cb *GetInputMethodListsCallback[T]) Register() js.Func[func(lists *InputMethodLists)] {
	return js.RegisterCallback[func(lists *InputMethodLists)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetInputMethodListsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InputMethodLists
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

type InputMethod struct {
	// Id is "InputMethod.id"
	//
	// Optional
	Id js.String
	// DisplayName is "InputMethod.displayName"
	//
	// Optional
	DisplayName js.String
	// LanguageCodes is "InputMethod.languageCodes"
	//
	// Optional
	LanguageCodes js.Array[js.String]
	// Tags is "InputMethod.tags"
	//
	// Optional
	Tags js.Array[js.String]
	// Enabled is "InputMethod.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// HasOptionsPage is "InputMethod.hasOptionsPage"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasOptionsPage MUST be set to true to make this field effective.
	HasOptionsPage bool
	// IsProhibitedByPolicy is "InputMethod.isProhibitedByPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsProhibitedByPolicy MUST be set to true to make this field effective.
	IsProhibitedByPolicy bool

	FFI_USE_Enabled              bool // for Enabled.
	FFI_USE_HasOptionsPage       bool // for HasOptionsPage.
	FFI_USE_IsProhibitedByPolicy bool // for IsProhibitedByPolicy.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputMethod with all fields set.
func (p InputMethod) FromRef(ref js.Ref) InputMethod {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InputMethod in the application heap.
func (p InputMethod) New() js.Ref {
	return bindings.InputMethodJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InputMethod) UpdateFrom(ref js.Ref) {
	bindings.InputMethodJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputMethod) Update(ref js.Ref) {
	bindings.InputMethodJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputMethod) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.DisplayName.Ref(),
		p.LanguageCodes.Ref(),
		p.Tags.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.LanguageCodes = p.LanguageCodes.FromRef(js.Undefined)
	p.Tags = p.Tags.FromRef(js.Undefined)
}

type InputMethodLists struct {
	// ComponentExtensionImes is "InputMethodLists.componentExtensionImes"
	//
	// Optional
	ComponentExtensionImes js.Array[InputMethod]
	// ThirdPartyExtensionImes is "InputMethodLists.thirdPartyExtensionImes"
	//
	// Optional
	ThirdPartyExtensionImes js.Array[InputMethod]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputMethodLists with all fields set.
func (p InputMethodLists) FromRef(ref js.Ref) InputMethodLists {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InputMethodLists in the application heap.
func (p InputMethodLists) New() js.Ref {
	return bindings.InputMethodListsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InputMethodLists) UpdateFrom(ref js.Ref) {
	bindings.InputMethodListsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputMethodLists) Update(ref js.Ref) {
	bindings.InputMethodListsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputMethodLists) FreeMembers(recursive bool) {
	js.Free(
		p.ComponentExtensionImes.Ref(),
		p.ThirdPartyExtensionImes.Ref(),
	)
	p.ComponentExtensionImes = p.ComponentExtensionImes.FromRef(js.Undefined)
	p.ThirdPartyExtensionImes = p.ThirdPartyExtensionImes.FromRef(js.Undefined)
}

type GetLanguageListCallbackFunc func(this js.Ref, languages js.Array[Language]) js.Ref

func (fn GetLanguageListCallbackFunc) Register() js.Func[func(languages js.Array[Language])] {
	return js.RegisterCallback[func(languages js.Array[Language])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLanguageListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Language]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetLanguageListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, languages js.Array[Language]) js.Ref
	Arg T
}

func (cb *GetLanguageListCallback[T]) Register() js.Func[func(languages js.Array[Language])] {
	return js.RegisterCallback[func(languages js.Array[Language])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLanguageListCallback[T]) DispatchCallback(
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

		js.Array[Language]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Language struct {
	// Code is "Language.code"
	//
	// Optional
	Code js.String
	// DisplayName is "Language.displayName"
	//
	// Optional
	DisplayName js.String
	// NativeDisplayName is "Language.nativeDisplayName"
	//
	// Optional
	NativeDisplayName js.String
	// SupportsUI is "Language.supportsUI"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportsUI MUST be set to true to make this field effective.
	SupportsUI bool
	// SupportsSpellcheck is "Language.supportsSpellcheck"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportsSpellcheck MUST be set to true to make this field effective.
	SupportsSpellcheck bool
	// SupportsTranslate is "Language.supportsTranslate"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportsTranslate MUST be set to true to make this field effective.
	SupportsTranslate bool
	// IsProhibitedLanguage is "Language.isProhibitedLanguage"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsProhibitedLanguage MUST be set to true to make this field effective.
	IsProhibitedLanguage bool

	FFI_USE_SupportsUI           bool // for SupportsUI.
	FFI_USE_SupportsSpellcheck   bool // for SupportsSpellcheck.
	FFI_USE_SupportsTranslate    bool // for SupportsTranslate.
	FFI_USE_IsProhibitedLanguage bool // for IsProhibitedLanguage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Language with all fields set.
func (p Language) FromRef(ref js.Ref) Language {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Language in the application heap.
func (p Language) New() js.Ref {
	return bindings.LanguageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Language) UpdateFrom(ref js.Ref) {
	bindings.LanguageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Language) Update(ref js.Ref) {
	bindings.LanguageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Language) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.DisplayName.Ref(),
		p.NativeDisplayName.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.NativeDisplayName = p.NativeDisplayName.FromRef(js.Undefined)
}

type GetNeverTranslateLanguagesCallbackFunc func(this js.Ref, languageCodes js.Array[js.String]) js.Ref

func (fn GetNeverTranslateLanguagesCallbackFunc) Register() js.Func[func(languageCodes js.Array[js.String])] {
	return js.RegisterCallback[func(languageCodes js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetNeverTranslateLanguagesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetNeverTranslateLanguagesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, languageCodes js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetNeverTranslateLanguagesCallback[T]) Register() js.Func[func(languageCodes js.Array[js.String])] {
	return js.RegisterCallback[func(languageCodes js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetNeverTranslateLanguagesCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSpellcheckDictionaryStatusesCallbackFunc func(this js.Ref, status js.Array[SpellcheckDictionaryStatus]) js.Ref

func (fn GetSpellcheckDictionaryStatusesCallbackFunc) Register() js.Func[func(status js.Array[SpellcheckDictionaryStatus])] {
	return js.RegisterCallback[func(status js.Array[SpellcheckDictionaryStatus])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSpellcheckDictionaryStatusesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SpellcheckDictionaryStatus]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSpellcheckDictionaryStatusesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status js.Array[SpellcheckDictionaryStatus]) js.Ref
	Arg T
}

func (cb *GetSpellcheckDictionaryStatusesCallback[T]) Register() js.Func[func(status js.Array[SpellcheckDictionaryStatus])] {
	return js.RegisterCallback[func(status js.Array[SpellcheckDictionaryStatus])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSpellcheckDictionaryStatusesCallback[T]) DispatchCallback(
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

		js.Array[SpellcheckDictionaryStatus]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SpellcheckDictionaryStatus struct {
	// LanguageCode is "SpellcheckDictionaryStatus.languageCode"
	//
	// Optional
	LanguageCode js.String
	// IsReady is "SpellcheckDictionaryStatus.isReady"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsReady MUST be set to true to make this field effective.
	IsReady bool
	// IsDownloading is "SpellcheckDictionaryStatus.isDownloading"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDownloading MUST be set to true to make this field effective.
	IsDownloading bool
	// DownloadFailed is "SpellcheckDictionaryStatus.downloadFailed"
	//
	// Optional
	//
	// NOTE: FFI_USE_DownloadFailed MUST be set to true to make this field effective.
	DownloadFailed bool

	FFI_USE_IsReady        bool // for IsReady.
	FFI_USE_IsDownloading  bool // for IsDownloading.
	FFI_USE_DownloadFailed bool // for DownloadFailed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpellcheckDictionaryStatus with all fields set.
func (p SpellcheckDictionaryStatus) FromRef(ref js.Ref) SpellcheckDictionaryStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpellcheckDictionaryStatus in the application heap.
func (p SpellcheckDictionaryStatus) New() js.Ref {
	return bindings.SpellcheckDictionaryStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpellcheckDictionaryStatus) UpdateFrom(ref js.Ref) {
	bindings.SpellcheckDictionaryStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpellcheckDictionaryStatus) Update(ref js.Ref) {
	bindings.SpellcheckDictionaryStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpellcheckDictionaryStatus) FreeMembers(recursive bool) {
	js.Free(
		p.LanguageCode.Ref(),
	)
	p.LanguageCode = p.LanguageCode.FromRef(js.Undefined)
}

type GetSpellcheckWordsCallbackFunc func(this js.Ref, words js.Array[js.String]) js.Ref

func (fn GetSpellcheckWordsCallbackFunc) Register() js.Func[func(words js.Array[js.String])] {
	return js.RegisterCallback[func(words js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSpellcheckWordsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSpellcheckWordsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, words js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetSpellcheckWordsCallback[T]) Register() js.Func[func(words js.Array[js.String])] {
	return js.RegisterCallback[func(words js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSpellcheckWordsCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetTranslateTargetLanguageCallbackFunc func(this js.Ref, languageCode js.String) js.Ref

func (fn GetTranslateTargetLanguageCallbackFunc) Register() js.Func[func(languageCode js.String)] {
	return js.RegisterCallback[func(languageCode js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetTranslateTargetLanguageCallbackFunc) DispatchCallback(
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

type GetTranslateTargetLanguageCallback[T any] struct {
	Fn  func(arg T, this js.Ref, languageCode js.String) js.Ref
	Arg T
}

func (cb *GetTranslateTargetLanguageCallback[T]) Register() js.Func[func(languageCode js.String)] {
	return js.RegisterCallback[func(languageCode js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetTranslateTargetLanguageCallback[T]) DispatchCallback(
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

type MoveType uint32

const (
	_ MoveType = iota

	MoveType_TOP
	MoveType_UP
	MoveType_DOWN
	MoveType_UNKNOWN
)

func (MoveType) FromRef(str js.Ref) MoveType {
	return MoveType(bindings.ConstOfMoveType(str))
}

func (x MoveType) String() (string, bool) {
	switch x {
	case MoveType_TOP:
		return "TOP", true
	case MoveType_UP:
		return "UP", true
	case MoveType_DOWN:
		return "DOWN", true
	case MoveType_UNKNOWN:
		return "UNKNOWN", true
	default:
		return "", false
	}
}

// HasFuncAddInputMethod returns true if the function "WEBEXT.languageSettingsPrivate.addInputMethod" exists.
func HasFuncAddInputMethod() bool {
	return js.True == bindings.HasFuncAddInputMethod()
}

// FuncAddInputMethod returns the function "WEBEXT.languageSettingsPrivate.addInputMethod".
func FuncAddInputMethod() (fn js.Func[func(inputMethodId js.String)]) {
	bindings.FuncAddInputMethod(
		js.Pointer(&fn),
	)
	return
}

// AddInputMethod calls the function "WEBEXT.languageSettingsPrivate.addInputMethod" directly.
func AddInputMethod(inputMethodId js.String) (ret js.Void) {
	bindings.CallAddInputMethod(
		js.Pointer(&ret),
		inputMethodId.Ref(),
	)

	return
}

// TryAddInputMethod calls the function "WEBEXT.languageSettingsPrivate.addInputMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddInputMethod(inputMethodId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddInputMethod(
		js.Pointer(&ret), js.Pointer(&exception),
		inputMethodId.Ref(),
	)

	return
}

// HasFuncAddSpellcheckWord returns true if the function "WEBEXT.languageSettingsPrivate.addSpellcheckWord" exists.
func HasFuncAddSpellcheckWord() bool {
	return js.True == bindings.HasFuncAddSpellcheckWord()
}

// FuncAddSpellcheckWord returns the function "WEBEXT.languageSettingsPrivate.addSpellcheckWord".
func FuncAddSpellcheckWord() (fn js.Func[func(word js.String)]) {
	bindings.FuncAddSpellcheckWord(
		js.Pointer(&fn),
	)
	return
}

// AddSpellcheckWord calls the function "WEBEXT.languageSettingsPrivate.addSpellcheckWord" directly.
func AddSpellcheckWord(word js.String) (ret js.Void) {
	bindings.CallAddSpellcheckWord(
		js.Pointer(&ret),
		word.Ref(),
	)

	return
}

// TryAddSpellcheckWord calls the function "WEBEXT.languageSettingsPrivate.addSpellcheckWord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddSpellcheckWord(word js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddSpellcheckWord(
		js.Pointer(&ret), js.Pointer(&exception),
		word.Ref(),
	)

	return
}

// HasFuncDisableLanguage returns true if the function "WEBEXT.languageSettingsPrivate.disableLanguage" exists.
func HasFuncDisableLanguage() bool {
	return js.True == bindings.HasFuncDisableLanguage()
}

// FuncDisableLanguage returns the function "WEBEXT.languageSettingsPrivate.disableLanguage".
func FuncDisableLanguage() (fn js.Func[func(languageCode js.String)]) {
	bindings.FuncDisableLanguage(
		js.Pointer(&fn),
	)
	return
}

// DisableLanguage calls the function "WEBEXT.languageSettingsPrivate.disableLanguage" directly.
func DisableLanguage(languageCode js.String) (ret js.Void) {
	bindings.CallDisableLanguage(
		js.Pointer(&ret),
		languageCode.Ref(),
	)

	return
}

// TryDisableLanguage calls the function "WEBEXT.languageSettingsPrivate.disableLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisableLanguage(languageCode js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisableLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
	)

	return
}

// HasFuncEnableLanguage returns true if the function "WEBEXT.languageSettingsPrivate.enableLanguage" exists.
func HasFuncEnableLanguage() bool {
	return js.True == bindings.HasFuncEnableLanguage()
}

// FuncEnableLanguage returns the function "WEBEXT.languageSettingsPrivate.enableLanguage".
func FuncEnableLanguage() (fn js.Func[func(languageCode js.String)]) {
	bindings.FuncEnableLanguage(
		js.Pointer(&fn),
	)
	return
}

// EnableLanguage calls the function "WEBEXT.languageSettingsPrivate.enableLanguage" directly.
func EnableLanguage(languageCode js.String) (ret js.Void) {
	bindings.CallEnableLanguage(
		js.Pointer(&ret),
		languageCode.Ref(),
	)

	return
}

// TryEnableLanguage calls the function "WEBEXT.languageSettingsPrivate.enableLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableLanguage(languageCode js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
	)

	return
}

// HasFuncGetAlwaysTranslateLanguages returns true if the function "WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages" exists.
func HasFuncGetAlwaysTranslateLanguages() bool {
	return js.True == bindings.HasFuncGetAlwaysTranslateLanguages()
}

// FuncGetAlwaysTranslateLanguages returns the function "WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages".
func FuncGetAlwaysTranslateLanguages() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetAlwaysTranslateLanguages(
		js.Pointer(&fn),
	)
	return
}

// GetAlwaysTranslateLanguages calls the function "WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages" directly.
func GetAlwaysTranslateLanguages() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetAlwaysTranslateLanguages(
		js.Pointer(&ret),
	)

	return
}

// TryGetAlwaysTranslateLanguages calls the function "WEBEXT.languageSettingsPrivate.getAlwaysTranslateLanguages"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAlwaysTranslateLanguages() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAlwaysTranslateLanguages(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInputMethodLists returns true if the function "WEBEXT.languageSettingsPrivate.getInputMethodLists" exists.
func HasFuncGetInputMethodLists() bool {
	return js.True == bindings.HasFuncGetInputMethodLists()
}

// FuncGetInputMethodLists returns the function "WEBEXT.languageSettingsPrivate.getInputMethodLists".
func FuncGetInputMethodLists() (fn js.Func[func() js.Promise[InputMethodLists]]) {
	bindings.FuncGetInputMethodLists(
		js.Pointer(&fn),
	)
	return
}

// GetInputMethodLists calls the function "WEBEXT.languageSettingsPrivate.getInputMethodLists" directly.
func GetInputMethodLists() (ret js.Promise[InputMethodLists]) {
	bindings.CallGetInputMethodLists(
		js.Pointer(&ret),
	)

	return
}

// TryGetInputMethodLists calls the function "WEBEXT.languageSettingsPrivate.getInputMethodLists"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInputMethodLists() (ret js.Promise[InputMethodLists], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInputMethodLists(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLanguageList returns true if the function "WEBEXT.languageSettingsPrivate.getLanguageList" exists.
func HasFuncGetLanguageList() bool {
	return js.True == bindings.HasFuncGetLanguageList()
}

// FuncGetLanguageList returns the function "WEBEXT.languageSettingsPrivate.getLanguageList".
func FuncGetLanguageList() (fn js.Func[func() js.Promise[js.Array[Language]]]) {
	bindings.FuncGetLanguageList(
		js.Pointer(&fn),
	)
	return
}

// GetLanguageList calls the function "WEBEXT.languageSettingsPrivate.getLanguageList" directly.
func GetLanguageList() (ret js.Promise[js.Array[Language]]) {
	bindings.CallGetLanguageList(
		js.Pointer(&ret),
	)

	return
}

// TryGetLanguageList calls the function "WEBEXT.languageSettingsPrivate.getLanguageList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLanguageList() (ret js.Promise[js.Array[Language]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLanguageList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetNeverTranslateLanguages returns true if the function "WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages" exists.
func HasFuncGetNeverTranslateLanguages() bool {
	return js.True == bindings.HasFuncGetNeverTranslateLanguages()
}

// FuncGetNeverTranslateLanguages returns the function "WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages".
func FuncGetNeverTranslateLanguages() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetNeverTranslateLanguages(
		js.Pointer(&fn),
	)
	return
}

// GetNeverTranslateLanguages calls the function "WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages" directly.
func GetNeverTranslateLanguages() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetNeverTranslateLanguages(
		js.Pointer(&ret),
	)

	return
}

// TryGetNeverTranslateLanguages calls the function "WEBEXT.languageSettingsPrivate.getNeverTranslateLanguages"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNeverTranslateLanguages() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNeverTranslateLanguages(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSpellcheckDictionaryStatuses returns true if the function "WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses" exists.
func HasFuncGetSpellcheckDictionaryStatuses() bool {
	return js.True == bindings.HasFuncGetSpellcheckDictionaryStatuses()
}

// FuncGetSpellcheckDictionaryStatuses returns the function "WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses".
func FuncGetSpellcheckDictionaryStatuses() (fn js.Func[func() js.Promise[js.Array[SpellcheckDictionaryStatus]]]) {
	bindings.FuncGetSpellcheckDictionaryStatuses(
		js.Pointer(&fn),
	)
	return
}

// GetSpellcheckDictionaryStatuses calls the function "WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses" directly.
func GetSpellcheckDictionaryStatuses() (ret js.Promise[js.Array[SpellcheckDictionaryStatus]]) {
	bindings.CallGetSpellcheckDictionaryStatuses(
		js.Pointer(&ret),
	)

	return
}

// TryGetSpellcheckDictionaryStatuses calls the function "WEBEXT.languageSettingsPrivate.getSpellcheckDictionaryStatuses"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSpellcheckDictionaryStatuses() (ret js.Promise[js.Array[SpellcheckDictionaryStatus]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSpellcheckDictionaryStatuses(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSpellcheckWords returns true if the function "WEBEXT.languageSettingsPrivate.getSpellcheckWords" exists.
func HasFuncGetSpellcheckWords() bool {
	return js.True == bindings.HasFuncGetSpellcheckWords()
}

// FuncGetSpellcheckWords returns the function "WEBEXT.languageSettingsPrivate.getSpellcheckWords".
func FuncGetSpellcheckWords() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetSpellcheckWords(
		js.Pointer(&fn),
	)
	return
}

// GetSpellcheckWords calls the function "WEBEXT.languageSettingsPrivate.getSpellcheckWords" directly.
func GetSpellcheckWords() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetSpellcheckWords(
		js.Pointer(&ret),
	)

	return
}

// TryGetSpellcheckWords calls the function "WEBEXT.languageSettingsPrivate.getSpellcheckWords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSpellcheckWords() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSpellcheckWords(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetTranslateTargetLanguage returns true if the function "WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage" exists.
func HasFuncGetTranslateTargetLanguage() bool {
	return js.True == bindings.HasFuncGetTranslateTargetLanguage()
}

// FuncGetTranslateTargetLanguage returns the function "WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage".
func FuncGetTranslateTargetLanguage() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetTranslateTargetLanguage(
		js.Pointer(&fn),
	)
	return
}

// GetTranslateTargetLanguage calls the function "WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage" directly.
func GetTranslateTargetLanguage() (ret js.Promise[js.String]) {
	bindings.CallGetTranslateTargetLanguage(
		js.Pointer(&ret),
	)

	return
}

// TryGetTranslateTargetLanguage calls the function "WEBEXT.languageSettingsPrivate.getTranslateTargetLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTranslateTargetLanguage() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTranslateTargetLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveLanguage returns true if the function "WEBEXT.languageSettingsPrivate.moveLanguage" exists.
func HasFuncMoveLanguage() bool {
	return js.True == bindings.HasFuncMoveLanguage()
}

// FuncMoveLanguage returns the function "WEBEXT.languageSettingsPrivate.moveLanguage".
func FuncMoveLanguage() (fn js.Func[func(languageCode js.String, moveType MoveType)]) {
	bindings.FuncMoveLanguage(
		js.Pointer(&fn),
	)
	return
}

// MoveLanguage calls the function "WEBEXT.languageSettingsPrivate.moveLanguage" directly.
func MoveLanguage(languageCode js.String, moveType MoveType) (ret js.Void) {
	bindings.CallMoveLanguage(
		js.Pointer(&ret),
		languageCode.Ref(),
		uint32(moveType),
	)

	return
}

// TryMoveLanguage calls the function "WEBEXT.languageSettingsPrivate.moveLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMoveLanguage(languageCode js.String, moveType MoveType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMoveLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
		uint32(moveType),
	)

	return
}

type OnCustomDictionaryChangedEventCallbackFunc func(this js.Ref, wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String]) js.Ref

func (fn OnCustomDictionaryChangedEventCallbackFunc) Register() js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])] {
	return js.RegisterCallback[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCustomDictionaryChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
		js.Array[js.String]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCustomDictionaryChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnCustomDictionaryChangedEventCallback[T]) Register() js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])] {
	return js.RegisterCallback[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCustomDictionaryChangedEventCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
		js.Array[js.String]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCustomDictionaryChanged returns true if the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener" exists.
func HasFuncOnCustomDictionaryChanged() bool {
	return js.True == bindings.HasFuncOnCustomDictionaryChanged()
}

// FuncOnCustomDictionaryChanged returns the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener".
func FuncOnCustomDictionaryChanged() (fn js.Func[func(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])])]) {
	bindings.FuncOnCustomDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// OnCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener" directly.
func OnCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnCustomDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCustomDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCustomDictionaryChanged returns true if the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener" exists.
func HasFuncOffCustomDictionaryChanged() bool {
	return js.True == bindings.HasFuncOffCustomDictionaryChanged()
}

// FuncOffCustomDictionaryChanged returns the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener".
func FuncOffCustomDictionaryChanged() (fn js.Func[func(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])])]) {
	bindings.FuncOffCustomDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// OffCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener" directly.
func OffCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffCustomDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCustomDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCustomDictionaryChanged returns true if the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener" exists.
func HasFuncHasOnCustomDictionaryChanged() bool {
	return js.True == bindings.HasFuncHasOnCustomDictionaryChanged()
}

// FuncHasOnCustomDictionaryChanged returns the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener".
func FuncHasOnCustomDictionaryChanged() (fn js.Func[func(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) bool]) {
	bindings.FuncHasOnCustomDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener" directly.
func HasOnCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnCustomDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCustomDictionaryChanged calls the function "WEBEXT.languageSettingsPrivate.onCustomDictionaryChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCustomDictionaryChanged(callback js.Func[func(wordsAdded js.Array[js.String], wordsRemoved js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCustomDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputMethodAddedEventCallbackFunc func(this js.Ref, inputMethodId js.String) js.Ref

func (fn OnInputMethodAddedEventCallbackFunc) Register() js.Func[func(inputMethodId js.String)] {
	return js.RegisterCallback[func(inputMethodId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputMethodAddedEventCallbackFunc) DispatchCallback(
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

type OnInputMethodAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, inputMethodId js.String) js.Ref
	Arg T
}

func (cb *OnInputMethodAddedEventCallback[T]) Register() js.Func[func(inputMethodId js.String)] {
	return js.RegisterCallback[func(inputMethodId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputMethodAddedEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputMethodAdded returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener" exists.
func HasFuncOnInputMethodAdded() bool {
	return js.True == bindings.HasFuncOnInputMethodAdded()
}

// FuncOnInputMethodAdded returns the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener".
func FuncOnInputMethodAdded() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)])]) {
	bindings.FuncOnInputMethodAdded(
		js.Pointer(&fn),
	)
	return
}

// OnInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener" directly.
func OnInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret js.Void) {
	bindings.CallOnInputMethodAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputMethodAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputMethodAdded returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener" exists.
func HasFuncOffInputMethodAdded() bool {
	return js.True == bindings.HasFuncOffInputMethodAdded()
}

// FuncOffInputMethodAdded returns the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener".
func FuncOffInputMethodAdded() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)])]) {
	bindings.FuncOffInputMethodAdded(
		js.Pointer(&fn),
	)
	return
}

// OffInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener" directly.
func OffInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret js.Void) {
	bindings.CallOffInputMethodAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputMethodAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputMethodAdded returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener" exists.
func HasFuncHasOnInputMethodAdded() bool {
	return js.True == bindings.HasFuncHasOnInputMethodAdded()
}

// FuncHasOnInputMethodAdded returns the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener".
func FuncHasOnInputMethodAdded() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)]) bool]) {
	bindings.FuncHasOnInputMethodAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener" directly.
func HasOnInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret bool) {
	bindings.CallHasOnInputMethodAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputMethodAdded calls the function "WEBEXT.languageSettingsPrivate.onInputMethodAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputMethodAdded(callback js.Func[func(inputMethodId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputMethodAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputMethodRemovedEventCallbackFunc func(this js.Ref, inputMethodId js.String) js.Ref

func (fn OnInputMethodRemovedEventCallbackFunc) Register() js.Func[func(inputMethodId js.String)] {
	return js.RegisterCallback[func(inputMethodId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputMethodRemovedEventCallbackFunc) DispatchCallback(
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

type OnInputMethodRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, inputMethodId js.String) js.Ref
	Arg T
}

func (cb *OnInputMethodRemovedEventCallback[T]) Register() js.Func[func(inputMethodId js.String)] {
	return js.RegisterCallback[func(inputMethodId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputMethodRemovedEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputMethodRemoved returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener" exists.
func HasFuncOnInputMethodRemoved() bool {
	return js.True == bindings.HasFuncOnInputMethodRemoved()
}

// FuncOnInputMethodRemoved returns the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener".
func FuncOnInputMethodRemoved() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)])]) {
	bindings.FuncOnInputMethodRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener" directly.
func OnInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret js.Void) {
	bindings.CallOnInputMethodRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputMethodRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputMethodRemoved returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener" exists.
func HasFuncOffInputMethodRemoved() bool {
	return js.True == bindings.HasFuncOffInputMethodRemoved()
}

// FuncOffInputMethodRemoved returns the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener".
func FuncOffInputMethodRemoved() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)])]) {
	bindings.FuncOffInputMethodRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener" directly.
func OffInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret js.Void) {
	bindings.CallOffInputMethodRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputMethodRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputMethodRemoved returns true if the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener" exists.
func HasFuncHasOnInputMethodRemoved() bool {
	return js.True == bindings.HasFuncHasOnInputMethodRemoved()
}

// FuncHasOnInputMethodRemoved returns the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener".
func FuncHasOnInputMethodRemoved() (fn js.Func[func(callback js.Func[func(inputMethodId js.String)]) bool]) {
	bindings.FuncHasOnInputMethodRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener" directly.
func HasOnInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret bool) {
	bindings.CallHasOnInputMethodRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputMethodRemoved calls the function "WEBEXT.languageSettingsPrivate.onInputMethodRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputMethodRemoved(callback js.Func[func(inputMethodId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputMethodRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSpellcheckDictionariesChangedEventCallbackFunc func(this js.Ref, statuses js.Array[SpellcheckDictionaryStatus]) js.Ref

func (fn OnSpellcheckDictionariesChangedEventCallbackFunc) Register() js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])] {
	return js.RegisterCallback[func(statuses js.Array[SpellcheckDictionaryStatus])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpellcheckDictionariesChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SpellcheckDictionaryStatus]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSpellcheckDictionariesChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, statuses js.Array[SpellcheckDictionaryStatus]) js.Ref
	Arg T
}

func (cb *OnSpellcheckDictionariesChangedEventCallback[T]) Register() js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])] {
	return js.RegisterCallback[func(statuses js.Array[SpellcheckDictionaryStatus])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpellcheckDictionariesChangedEventCallback[T]) DispatchCallback(
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

		js.Array[SpellcheckDictionaryStatus]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSpellcheckDictionariesChanged returns true if the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener" exists.
func HasFuncOnSpellcheckDictionariesChanged() bool {
	return js.True == bindings.HasFuncOnSpellcheckDictionariesChanged()
}

// FuncOnSpellcheckDictionariesChanged returns the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener".
func FuncOnSpellcheckDictionariesChanged() (fn js.Func[func(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])])]) {
	bindings.FuncOnSpellcheckDictionariesChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener" directly.
func OnSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret js.Void) {
	bindings.CallOnSpellcheckDictionariesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSpellcheckDictionariesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSpellcheckDictionariesChanged returns true if the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener" exists.
func HasFuncOffSpellcheckDictionariesChanged() bool {
	return js.True == bindings.HasFuncOffSpellcheckDictionariesChanged()
}

// FuncOffSpellcheckDictionariesChanged returns the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener".
func FuncOffSpellcheckDictionariesChanged() (fn js.Func[func(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])])]) {
	bindings.FuncOffSpellcheckDictionariesChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener" directly.
func OffSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret js.Void) {
	bindings.CallOffSpellcheckDictionariesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSpellcheckDictionariesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSpellcheckDictionariesChanged returns true if the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener" exists.
func HasFuncHasOnSpellcheckDictionariesChanged() bool {
	return js.True == bindings.HasFuncHasOnSpellcheckDictionariesChanged()
}

// FuncHasOnSpellcheckDictionariesChanged returns the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener".
func FuncHasOnSpellcheckDictionariesChanged() (fn js.Func[func(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) bool]) {
	bindings.FuncHasOnSpellcheckDictionariesChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener" directly.
func HasOnSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret bool) {
	bindings.CallHasOnSpellcheckDictionariesChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSpellcheckDictionariesChanged calls the function "WEBEXT.languageSettingsPrivate.onSpellcheckDictionariesChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSpellcheckDictionariesChanged(callback js.Func[func(statuses js.Array[SpellcheckDictionaryStatus])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSpellcheckDictionariesChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveInputMethod returns true if the function "WEBEXT.languageSettingsPrivate.removeInputMethod" exists.
func HasFuncRemoveInputMethod() bool {
	return js.True == bindings.HasFuncRemoveInputMethod()
}

// FuncRemoveInputMethod returns the function "WEBEXT.languageSettingsPrivate.removeInputMethod".
func FuncRemoveInputMethod() (fn js.Func[func(inputMethodId js.String)]) {
	bindings.FuncRemoveInputMethod(
		js.Pointer(&fn),
	)
	return
}

// RemoveInputMethod calls the function "WEBEXT.languageSettingsPrivate.removeInputMethod" directly.
func RemoveInputMethod(inputMethodId js.String) (ret js.Void) {
	bindings.CallRemoveInputMethod(
		js.Pointer(&ret),
		inputMethodId.Ref(),
	)

	return
}

// TryRemoveInputMethod calls the function "WEBEXT.languageSettingsPrivate.removeInputMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveInputMethod(inputMethodId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveInputMethod(
		js.Pointer(&ret), js.Pointer(&exception),
		inputMethodId.Ref(),
	)

	return
}

// HasFuncRemoveSpellcheckWord returns true if the function "WEBEXT.languageSettingsPrivate.removeSpellcheckWord" exists.
func HasFuncRemoveSpellcheckWord() bool {
	return js.True == bindings.HasFuncRemoveSpellcheckWord()
}

// FuncRemoveSpellcheckWord returns the function "WEBEXT.languageSettingsPrivate.removeSpellcheckWord".
func FuncRemoveSpellcheckWord() (fn js.Func[func(word js.String)]) {
	bindings.FuncRemoveSpellcheckWord(
		js.Pointer(&fn),
	)
	return
}

// RemoveSpellcheckWord calls the function "WEBEXT.languageSettingsPrivate.removeSpellcheckWord" directly.
func RemoveSpellcheckWord(word js.String) (ret js.Void) {
	bindings.CallRemoveSpellcheckWord(
		js.Pointer(&ret),
		word.Ref(),
	)

	return
}

// TryRemoveSpellcheckWord calls the function "WEBEXT.languageSettingsPrivate.removeSpellcheckWord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveSpellcheckWord(word js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveSpellcheckWord(
		js.Pointer(&ret), js.Pointer(&exception),
		word.Ref(),
	)

	return
}

// HasFuncRetryDownloadDictionary returns true if the function "WEBEXT.languageSettingsPrivate.retryDownloadDictionary" exists.
func HasFuncRetryDownloadDictionary() bool {
	return js.True == bindings.HasFuncRetryDownloadDictionary()
}

// FuncRetryDownloadDictionary returns the function "WEBEXT.languageSettingsPrivate.retryDownloadDictionary".
func FuncRetryDownloadDictionary() (fn js.Func[func(languageCode js.String)]) {
	bindings.FuncRetryDownloadDictionary(
		js.Pointer(&fn),
	)
	return
}

// RetryDownloadDictionary calls the function "WEBEXT.languageSettingsPrivate.retryDownloadDictionary" directly.
func RetryDownloadDictionary(languageCode js.String) (ret js.Void) {
	bindings.CallRetryDownloadDictionary(
		js.Pointer(&ret),
		languageCode.Ref(),
	)

	return
}

// TryRetryDownloadDictionary calls the function "WEBEXT.languageSettingsPrivate.retryDownloadDictionary"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRetryDownloadDictionary(languageCode js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRetryDownloadDictionary(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
	)

	return
}

// HasFuncSetEnableTranslationForLanguage returns true if the function "WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage" exists.
func HasFuncSetEnableTranslationForLanguage() bool {
	return js.True == bindings.HasFuncSetEnableTranslationForLanguage()
}

// FuncSetEnableTranslationForLanguage returns the function "WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage".
func FuncSetEnableTranslationForLanguage() (fn js.Func[func(languageCode js.String, enable bool)]) {
	bindings.FuncSetEnableTranslationForLanguage(
		js.Pointer(&fn),
	)
	return
}

// SetEnableTranslationForLanguage calls the function "WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage" directly.
func SetEnableTranslationForLanguage(languageCode js.String, enable bool) (ret js.Void) {
	bindings.CallSetEnableTranslationForLanguage(
		js.Pointer(&ret),
		languageCode.Ref(),
		js.Bool(bool(enable)),
	)

	return
}

// TrySetEnableTranslationForLanguage calls the function "WEBEXT.languageSettingsPrivate.setEnableTranslationForLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetEnableTranslationForLanguage(languageCode js.String, enable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetEnableTranslationForLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
		js.Bool(bool(enable)),
	)

	return
}

// HasFuncSetLanguageAlwaysTranslateState returns true if the function "WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState" exists.
func HasFuncSetLanguageAlwaysTranslateState() bool {
	return js.True == bindings.HasFuncSetLanguageAlwaysTranslateState()
}

// FuncSetLanguageAlwaysTranslateState returns the function "WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState".
func FuncSetLanguageAlwaysTranslateState() (fn js.Func[func(languageCode js.String, alwaysTranslate bool)]) {
	bindings.FuncSetLanguageAlwaysTranslateState(
		js.Pointer(&fn),
	)
	return
}

// SetLanguageAlwaysTranslateState calls the function "WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState" directly.
func SetLanguageAlwaysTranslateState(languageCode js.String, alwaysTranslate bool) (ret js.Void) {
	bindings.CallSetLanguageAlwaysTranslateState(
		js.Pointer(&ret),
		languageCode.Ref(),
		js.Bool(bool(alwaysTranslate)),
	)

	return
}

// TrySetLanguageAlwaysTranslateState calls the function "WEBEXT.languageSettingsPrivate.setLanguageAlwaysTranslateState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetLanguageAlwaysTranslateState(languageCode js.String, alwaysTranslate bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetLanguageAlwaysTranslateState(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
		js.Bool(bool(alwaysTranslate)),
	)

	return
}

// HasFuncSetTranslateTargetLanguage returns true if the function "WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage" exists.
func HasFuncSetTranslateTargetLanguage() bool {
	return js.True == bindings.HasFuncSetTranslateTargetLanguage()
}

// FuncSetTranslateTargetLanguage returns the function "WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage".
func FuncSetTranslateTargetLanguage() (fn js.Func[func(languageCode js.String)]) {
	bindings.FuncSetTranslateTargetLanguage(
		js.Pointer(&fn),
	)
	return
}

// SetTranslateTargetLanguage calls the function "WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage" directly.
func SetTranslateTargetLanguage(languageCode js.String) (ret js.Void) {
	bindings.CallSetTranslateTargetLanguage(
		js.Pointer(&ret),
		languageCode.Ref(),
	)

	return
}

// TrySetTranslateTargetLanguage calls the function "WEBEXT.languageSettingsPrivate.setTranslateTargetLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTranslateTargetLanguage(languageCode js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTranslateTargetLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		languageCode.Ref(),
	)

	return
}

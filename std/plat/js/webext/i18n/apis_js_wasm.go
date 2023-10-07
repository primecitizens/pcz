// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package i18n

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/i18n/bindings"
)

type LanguageCode = js.String

type DetectLanguageReturnTypeFieldLanguagesElem struct {
	// Language is "DetectLanguageReturnTypeFieldLanguagesElem.language"
	//
	// Required
	Language LanguageCode
	// Percentage is "DetectLanguageReturnTypeFieldLanguagesElem.percentage"
	//
	// Required
	Percentage int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DetectLanguageReturnTypeFieldLanguagesElem with all fields set.
func (p DetectLanguageReturnTypeFieldLanguagesElem) FromRef(ref js.Ref) DetectLanguageReturnTypeFieldLanguagesElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DetectLanguageReturnTypeFieldLanguagesElem in the application heap.
func (p DetectLanguageReturnTypeFieldLanguagesElem) New() js.Ref {
	return bindings.DetectLanguageReturnTypeFieldLanguagesElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DetectLanguageReturnTypeFieldLanguagesElem) UpdateFrom(ref js.Ref) {
	bindings.DetectLanguageReturnTypeFieldLanguagesElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DetectLanguageReturnTypeFieldLanguagesElem) Update(ref js.Ref) {
	bindings.DetectLanguageReturnTypeFieldLanguagesElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DetectLanguageReturnTypeFieldLanguagesElem) FreeMembers(recursive bool) {
	js.Free(
		p.Language.Ref(),
	)
	p.Language = p.Language.FromRef(js.Undefined)
}

type DetectLanguageReturnType struct {
	// IsReliable is "DetectLanguageReturnType.isReliable"
	//
	// Required
	IsReliable bool
	// Languages is "DetectLanguageReturnType.languages"
	//
	// Required
	Languages js.Array[DetectLanguageReturnTypeFieldLanguagesElem]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DetectLanguageReturnType with all fields set.
func (p DetectLanguageReturnType) FromRef(ref js.Ref) DetectLanguageReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DetectLanguageReturnType in the application heap.
func (p DetectLanguageReturnType) New() js.Ref {
	return bindings.DetectLanguageReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DetectLanguageReturnType) UpdateFrom(ref js.Ref) {
	bindings.DetectLanguageReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DetectLanguageReturnType) Update(ref js.Ref) {
	bindings.DetectLanguageReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DetectLanguageReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Languages.Ref(),
	)
	p.Languages = p.Languages.FromRef(js.Undefined)
}

type GetMessageArgOptions struct {
	// EscapeLt is "GetMessageArgOptions.escapeLt"
	//
	// Optional
	//
	// NOTE: FFI_USE_EscapeLt MUST be set to true to make this field effective.
	EscapeLt bool

	FFI_USE_EscapeLt bool // for EscapeLt.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetMessageArgOptions with all fields set.
func (p GetMessageArgOptions) FromRef(ref js.Ref) GetMessageArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetMessageArgOptions in the application heap.
func (p GetMessageArgOptions) New() js.Ref {
	return bindings.GetMessageArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetMessageArgOptions) UpdateFrom(ref js.Ref) {
	bindings.GetMessageArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetMessageArgOptions) Update(ref js.Ref) {
	bindings.GetMessageArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetMessageArgOptions) FreeMembers(recursive bool) {
}

// HasFuncDetectLanguage returns true if the function "WEBEXT.i18n.detectLanguage" exists.
func HasFuncDetectLanguage() bool {
	return js.True == bindings.HasFuncDetectLanguage()
}

// FuncDetectLanguage returns the function "WEBEXT.i18n.detectLanguage".
func FuncDetectLanguage() (fn js.Func[func(text js.String) js.Promise[DetectLanguageReturnType]]) {
	bindings.FuncDetectLanguage(
		js.Pointer(&fn),
	)
	return
}

// DetectLanguage calls the function "WEBEXT.i18n.detectLanguage" directly.
func DetectLanguage(text js.String) (ret js.Promise[DetectLanguageReturnType]) {
	bindings.CallDetectLanguage(
		js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryDetectLanguage calls the function "WEBEXT.i18n.detectLanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDetectLanguage(text js.String) (ret js.Promise[DetectLanguageReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDetectLanguage(
		js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasFuncGetAcceptLanguages returns true if the function "WEBEXT.i18n.getAcceptLanguages" exists.
func HasFuncGetAcceptLanguages() bool {
	return js.True == bindings.HasFuncGetAcceptLanguages()
}

// FuncGetAcceptLanguages returns the function "WEBEXT.i18n.getAcceptLanguages".
func FuncGetAcceptLanguages() (fn js.Func[func() js.Promise[js.Array[LanguageCode]]]) {
	bindings.FuncGetAcceptLanguages(
		js.Pointer(&fn),
	)
	return
}

// GetAcceptLanguages calls the function "WEBEXT.i18n.getAcceptLanguages" directly.
func GetAcceptLanguages() (ret js.Promise[js.Array[LanguageCode]]) {
	bindings.CallGetAcceptLanguages(
		js.Pointer(&ret),
	)

	return
}

// TryGetAcceptLanguages calls the function "WEBEXT.i18n.getAcceptLanguages"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAcceptLanguages() (ret js.Promise[js.Array[LanguageCode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAcceptLanguages(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMessage returns true if the function "WEBEXT.i18n.getMessage" exists.
func HasFuncGetMessage() bool {
	return js.True == bindings.HasFuncGetMessage()
}

// FuncGetMessage returns the function "WEBEXT.i18n.getMessage".
func FuncGetMessage() (fn js.Func[func(messageName js.String, substitutions js.Any, options GetMessageArgOptions) js.String]) {
	bindings.FuncGetMessage(
		js.Pointer(&fn),
	)
	return
}

// GetMessage calls the function "WEBEXT.i18n.getMessage" directly.
func GetMessage(messageName js.String, substitutions js.Any, options GetMessageArgOptions) (ret js.String) {
	bindings.CallGetMessage(
		js.Pointer(&ret),
		messageName.Ref(),
		substitutions.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetMessage calls the function "WEBEXT.i18n.getMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMessage(messageName js.String, substitutions js.Any, options GetMessageArgOptions) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		messageName.Ref(),
		substitutions.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetUILanguage returns true if the function "WEBEXT.i18n.getUILanguage" exists.
func HasFuncGetUILanguage() bool {
	return js.True == bindings.HasFuncGetUILanguage()
}

// FuncGetUILanguage returns the function "WEBEXT.i18n.getUILanguage".
func FuncGetUILanguage() (fn js.Func[func() js.String]) {
	bindings.FuncGetUILanguage(
		js.Pointer(&fn),
	)
	return
}

// GetUILanguage calls the function "WEBEXT.i18n.getUILanguage" directly.
func GetUILanguage() (ret js.String) {
	bindings.CallGetUILanguage(
		js.Pointer(&ret),
	)

	return
}

// TryGetUILanguage calls the function "WEBEXT.i18n.getUILanguage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUILanguage() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUILanguage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package inputmethodprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/inputmethodprivate/bindings"
)

type AutoCapitalizeType uint32

const (
	_ AutoCapitalizeType = iota

	AutoCapitalizeType_OFF
	AutoCapitalizeType_CHARACTERS
	AutoCapitalizeType_WORDS
	AutoCapitalizeType_SENTENCES
)

func (AutoCapitalizeType) FromRef(str js.Ref) AutoCapitalizeType {
	return AutoCapitalizeType(bindings.ConstOfAutoCapitalizeType(str))
}

func (x AutoCapitalizeType) String() (string, bool) {
	switch x {
	case AutoCapitalizeType_OFF:
		return "off", true
	case AutoCapitalizeType_CHARACTERS:
		return "characters", true
	case AutoCapitalizeType_WORDS:
		return "words", true
	case AutoCapitalizeType_SENTENCES:
		return "sentences", true
	default:
		return "", false
	}
}

type FinishComposingTextArgParameters struct {
	// ContextID is "FinishComposingTextArgParameters.contextID"
	//
	// Required
	ContextID int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FinishComposingTextArgParameters with all fields set.
func (p FinishComposingTextArgParameters) FromRef(ref js.Ref) FinishComposingTextArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FinishComposingTextArgParameters in the application heap.
func (p FinishComposingTextArgParameters) New() js.Ref {
	return bindings.FinishComposingTextArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FinishComposingTextArgParameters) UpdateFrom(ref js.Ref) {
	bindings.FinishComposingTextArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FinishComposingTextArgParameters) Update(ref js.Ref) {
	bindings.FinishComposingTextArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FinishComposingTextArgParameters) FreeMembers(recursive bool) {
}

type FocusReason uint32

const (
	_ FocusReason = iota

	FocusReason_MOUSE
	FocusReason_TOUCH
	FocusReason_PEN
	FocusReason_OTHER
)

func (FocusReason) FromRef(str js.Ref) FocusReason {
	return FocusReason(bindings.ConstOfFocusReason(str))
}

func (x FocusReason) String() (string, bool) {
	switch x {
	case FocusReason_MOUSE:
		return "mouse", true
	case FocusReason_TOUCH:
		return "touch", true
	case FocusReason_PEN:
		return "pen", true
	case FocusReason_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type GetInputMethodConfigReturnType struct {
	// IsImeMenuActivated is "GetInputMethodConfigReturnType.isImeMenuActivated"
	//
	// Required
	IsImeMenuActivated bool
	// IsPhysicalKeyboardAutocorrectEnabled is "GetInputMethodConfigReturnType.isPhysicalKeyboardAutocorrectEnabled"
	//
	// Required
	IsPhysicalKeyboardAutocorrectEnabled bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetInputMethodConfigReturnType with all fields set.
func (p GetInputMethodConfigReturnType) FromRef(ref js.Ref) GetInputMethodConfigReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetInputMethodConfigReturnType in the application heap.
func (p GetInputMethodConfigReturnType) New() js.Ref {
	return bindings.GetInputMethodConfigReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetInputMethodConfigReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetInputMethodConfigReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetInputMethodConfigReturnType) Update(ref js.Ref) {
	bindings.GetInputMethodConfigReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetInputMethodConfigReturnType) FreeMembers(recursive bool) {
}

type GetInputMethodsReturnTypeElem struct {
	// Id is "GetInputMethodsReturnTypeElem.id"
	//
	// Required
	Id js.String
	// Indicator is "GetInputMethodsReturnTypeElem.indicator"
	//
	// Required
	Indicator js.String
	// Name is "GetInputMethodsReturnTypeElem.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetInputMethodsReturnTypeElem with all fields set.
func (p GetInputMethodsReturnTypeElem) FromRef(ref js.Ref) GetInputMethodsReturnTypeElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetInputMethodsReturnTypeElem in the application heap.
func (p GetInputMethodsReturnTypeElem) New() js.Ref {
	return bindings.GetInputMethodsReturnTypeElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetInputMethodsReturnTypeElem) UpdateFrom(ref js.Ref) {
	bindings.GetInputMethodsReturnTypeElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetInputMethodsReturnTypeElem) Update(ref js.Ref) {
	bindings.GetInputMethodsReturnTypeElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetInputMethodsReturnTypeElem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Indicator.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Indicator = p.Indicator.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type GetSurroundingTextReturnType struct {
	// After is "GetSurroundingTextReturnType.after"
	//
	// Required
	After js.String
	// Before is "GetSurroundingTextReturnType.before"
	//
	// Required
	Before js.String
	// Selected is "GetSurroundingTextReturnType.selected"
	//
	// Required
	Selected js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetSurroundingTextReturnType with all fields set.
func (p GetSurroundingTextReturnType) FromRef(ref js.Ref) GetSurroundingTextReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetSurroundingTextReturnType in the application heap.
func (p GetSurroundingTextReturnType) New() js.Ref {
	return bindings.GetSurroundingTextReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetSurroundingTextReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetSurroundingTextReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetSurroundingTextReturnType) Update(ref js.Ref) {
	bindings.GetSurroundingTextReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetSurroundingTextReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.After.Ref(),
		p.Before.Ref(),
		p.Selected.Ref(),
	)
	p.After = p.After.FromRef(js.Undefined)
	p.Before = p.Before.FromRef(js.Undefined)
	p.Selected = p.Selected.FromRef(js.Undefined)
}

type GetTextFieldBoundsArgParameters struct {
	// ContextID is "GetTextFieldBoundsArgParameters.contextID"
	//
	// Required
	ContextID int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetTextFieldBoundsArgParameters with all fields set.
func (p GetTextFieldBoundsArgParameters) FromRef(ref js.Ref) GetTextFieldBoundsArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetTextFieldBoundsArgParameters in the application heap.
func (p GetTextFieldBoundsArgParameters) New() js.Ref {
	return bindings.GetTextFieldBoundsArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetTextFieldBoundsArgParameters) UpdateFrom(ref js.Ref) {
	bindings.GetTextFieldBoundsArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetTextFieldBoundsArgParameters) Update(ref js.Ref) {
	bindings.GetTextFieldBoundsArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetTextFieldBoundsArgParameters) FreeMembers(recursive bool) {
}

type GetTextFieldBoundsReturnType struct {
	// Height is "GetTextFieldBoundsReturnType.height"
	//
	// Required
	Height int64
	// Width is "GetTextFieldBoundsReturnType.width"
	//
	// Required
	Width int64
	// X is "GetTextFieldBoundsReturnType.x"
	//
	// Required
	X int64
	// Y is "GetTextFieldBoundsReturnType.y"
	//
	// Required
	Y int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetTextFieldBoundsReturnType with all fields set.
func (p GetTextFieldBoundsReturnType) FromRef(ref js.Ref) GetTextFieldBoundsReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetTextFieldBoundsReturnType in the application heap.
func (p GetTextFieldBoundsReturnType) New() js.Ref {
	return bindings.GetTextFieldBoundsReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetTextFieldBoundsReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetTextFieldBoundsReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetTextFieldBoundsReturnType) Update(ref js.Ref) {
	bindings.GetTextFieldBoundsReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetTextFieldBoundsReturnType) FreeMembers(recursive bool) {
}

type InputModeType uint32

const (
	_ InputModeType = iota

	InputModeType_NO_KEYBOARD
	InputModeType_TEXT
	InputModeType_TEL
	InputModeType_URL
	InputModeType_EMAIL
	InputModeType_NUMERIC
	InputModeType_DECIMAL
	InputModeType_SEARCH
)

func (InputModeType) FromRef(str js.Ref) InputModeType {
	return InputModeType(bindings.ConstOfInputModeType(str))
}

func (x InputModeType) String() (string, bool) {
	switch x {
	case InputModeType_NO_KEYBOARD:
		return "noKeyboard", true
	case InputModeType_TEXT:
		return "text", true
	case InputModeType_TEL:
		return "tel", true
	case InputModeType_URL:
		return "url", true
	case InputModeType_EMAIL:
		return "email", true
	case InputModeType_NUMERIC:
		return "numeric", true
	case InputModeType_DECIMAL:
		return "decimal", true
	case InputModeType_SEARCH:
		return "search", true
	default:
		return "", false
	}
}

type InputContextType uint32

const (
	_ InputContextType = iota

	InputContextType_TEXT
	InputContextType_SEARCH
	InputContextType_TEL
	InputContextType_URL
	InputContextType_EMAIL
	InputContextType_NUMBER
	InputContextType_PASSWORD
	InputContextType_NULL
)

func (InputContextType) FromRef(str js.Ref) InputContextType {
	return InputContextType(bindings.ConstOfInputContextType(str))
}

func (x InputContextType) String() (string, bool) {
	switch x {
	case InputContextType_TEXT:
		return "text", true
	case InputContextType_SEARCH:
		return "search", true
	case InputContextType_TEL:
		return "tel", true
	case InputContextType_URL:
		return "url", true
	case InputContextType_EMAIL:
		return "email", true
	case InputContextType_NUMBER:
		return "number", true
	case InputContextType_PASSWORD:
		return "password", true
	case InputContextType_NULL:
		return "null", true
	default:
		return "", false
	}
}

type InputContext struct {
	// AppKey is "InputContext.appKey"
	//
	// Optional
	AppKey js.String
	// AutoCapitalize is "InputContext.autoCapitalize"
	//
	// Required
	AutoCapitalize AutoCapitalizeType
	// AutoComplete is "InputContext.autoComplete"
	//
	// Required
	AutoComplete bool
	// AutoCorrect is "InputContext.autoCorrect"
	//
	// Required
	AutoCorrect bool
	// ContextID is "InputContext.contextID"
	//
	// Required
	ContextID int64
	// FocusReason is "InputContext.focusReason"
	//
	// Required
	FocusReason FocusReason
	// Mode is "InputContext.mode"
	//
	// Required
	Mode InputModeType
	// ShouldDoLearning is "InputContext.shouldDoLearning"
	//
	// Required
	ShouldDoLearning bool
	// SpellCheck is "InputContext.spellCheck"
	//
	// Required
	SpellCheck bool
	// Type is "InputContext.type"
	//
	// Required
	Type InputContextType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputContext with all fields set.
func (p InputContext) FromRef(ref js.Ref) InputContext {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InputContext in the application heap.
func (p InputContext) New() js.Ref {
	return bindings.InputContextJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InputContext) UpdateFrom(ref js.Ref) {
	bindings.InputContextJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputContext) Update(ref js.Ref) {
	bindings.InputContextJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputContext) FreeMembers(recursive bool) {
	js.Free(
		p.AppKey.Ref(),
	)
	p.AppKey = p.AppKey.FromRef(js.Undefined)
}

type InputMethodSettingsFieldPinyinFuzzyConfig struct {
	// AnAng is "InputMethodSettingsFieldPinyinFuzzyConfig.an_ang"
	//
	// Optional
	//
	// NOTE: FFI_USE_AnAng MUST be set to true to make this field effective.
	AnAng bool
	// CCh is "InputMethodSettingsFieldPinyinFuzzyConfig.c_ch"
	//
	// Optional
	//
	// NOTE: FFI_USE_CCh MUST be set to true to make this field effective.
	CCh bool
	// EnEng is "InputMethodSettingsFieldPinyinFuzzyConfig.en_eng"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnEng MUST be set to true to make this field effective.
	EnEng bool
	// FH is "InputMethodSettingsFieldPinyinFuzzyConfig.f_h"
	//
	// Optional
	//
	// NOTE: FFI_USE_FH MUST be set to true to make this field effective.
	FH bool
	// IanIang is "InputMethodSettingsFieldPinyinFuzzyConfig.ian_iang"
	//
	// Optional
	//
	// NOTE: FFI_USE_IanIang MUST be set to true to make this field effective.
	IanIang bool
	// InIng is "InputMethodSettingsFieldPinyinFuzzyConfig.in_ing"
	//
	// Optional
	//
	// NOTE: FFI_USE_InIng MUST be set to true to make this field effective.
	InIng bool
	// KG is "InputMethodSettingsFieldPinyinFuzzyConfig.k_g"
	//
	// Optional
	//
	// NOTE: FFI_USE_KG MUST be set to true to make this field effective.
	KG bool
	// LN is "InputMethodSettingsFieldPinyinFuzzyConfig.l_n"
	//
	// Optional
	//
	// NOTE: FFI_USE_LN MUST be set to true to make this field effective.
	LN bool
	// RL is "InputMethodSettingsFieldPinyinFuzzyConfig.r_l"
	//
	// Optional
	//
	// NOTE: FFI_USE_RL MUST be set to true to make this field effective.
	RL bool
	// SSh is "InputMethodSettingsFieldPinyinFuzzyConfig.s_sh"
	//
	// Optional
	//
	// NOTE: FFI_USE_SSh MUST be set to true to make this field effective.
	SSh bool
	// UanUang is "InputMethodSettingsFieldPinyinFuzzyConfig.uan_uang"
	//
	// Optional
	//
	// NOTE: FFI_USE_UanUang MUST be set to true to make this field effective.
	UanUang bool
	// ZZh is "InputMethodSettingsFieldPinyinFuzzyConfig.z_zh"
	//
	// Optional
	//
	// NOTE: FFI_USE_ZZh MUST be set to true to make this field effective.
	ZZh bool

	FFI_USE_AnAng   bool // for AnAng.
	FFI_USE_CCh     bool // for CCh.
	FFI_USE_EnEng   bool // for EnEng.
	FFI_USE_FH      bool // for FH.
	FFI_USE_IanIang bool // for IanIang.
	FFI_USE_InIng   bool // for InIng.
	FFI_USE_KG      bool // for KG.
	FFI_USE_LN      bool // for LN.
	FFI_USE_RL      bool // for RL.
	FFI_USE_SSh     bool // for SSh.
	FFI_USE_UanUang bool // for UanUang.
	FFI_USE_ZZh     bool // for ZZh.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputMethodSettingsFieldPinyinFuzzyConfig with all fields set.
func (p InputMethodSettingsFieldPinyinFuzzyConfig) FromRef(ref js.Ref) InputMethodSettingsFieldPinyinFuzzyConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InputMethodSettingsFieldPinyinFuzzyConfig in the application heap.
func (p InputMethodSettingsFieldPinyinFuzzyConfig) New() js.Ref {
	return bindings.InputMethodSettingsFieldPinyinFuzzyConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InputMethodSettingsFieldPinyinFuzzyConfig) UpdateFrom(ref js.Ref) {
	bindings.InputMethodSettingsFieldPinyinFuzzyConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputMethodSettingsFieldPinyinFuzzyConfig) Update(ref js.Ref) {
	bindings.InputMethodSettingsFieldPinyinFuzzyConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputMethodSettingsFieldPinyinFuzzyConfig) FreeMembers(recursive bool) {
}

type InputMethodSettings struct {
	// EnableCompletion is "InputMethodSettings.enableCompletion"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableCompletion MUST be set to true to make this field effective.
	EnableCompletion bool
	// EnableDoubleSpacePeriod is "InputMethodSettings.enableDoubleSpacePeriod"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableDoubleSpacePeriod MUST be set to true to make this field effective.
	EnableDoubleSpacePeriod bool
	// EnableGestureTyping is "InputMethodSettings.enableGestureTyping"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableGestureTyping MUST be set to true to make this field effective.
	EnableGestureTyping bool
	// EnablePrediction is "InputMethodSettings.enablePrediction"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnablePrediction MUST be set to true to make this field effective.
	EnablePrediction bool
	// EnableSoundOnKeypress is "InputMethodSettings.enableSoundOnKeypress"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableSoundOnKeypress MUST be set to true to make this field effective.
	EnableSoundOnKeypress bool
	// KoreanEnableSyllableInput is "InputMethodSettings.koreanEnableSyllableInput"
	//
	// Optional
	//
	// NOTE: FFI_USE_KoreanEnableSyllableInput MUST be set to true to make this field effective.
	KoreanEnableSyllableInput bool
	// KoreanKeyboardLayout is "InputMethodSettings.koreanKeyboardLayout"
	//
	// Optional
	KoreanKeyboardLayout js.String
	// KoreanShowHangulCandidate is "InputMethodSettings.koreanShowHangulCandidate"
	//
	// Optional
	//
	// NOTE: FFI_USE_KoreanShowHangulCandidate MUST be set to true to make this field effective.
	KoreanShowHangulCandidate bool
	// PhysicalKeyboardAutoCorrectionEnabledByDefault is "InputMethodSettings.physicalKeyboardAutoCorrectionEnabledByDefault"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalKeyboardAutoCorrectionEnabledByDefault MUST be set to true to make this field effective.
	PhysicalKeyboardAutoCorrectionEnabledByDefault bool
	// PhysicalKeyboardAutoCorrectionLevel is "InputMethodSettings.physicalKeyboardAutoCorrectionLevel"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalKeyboardAutoCorrectionLevel MUST be set to true to make this field effective.
	PhysicalKeyboardAutoCorrectionLevel int64
	// PhysicalKeyboardEnableCapitalization is "InputMethodSettings.physicalKeyboardEnableCapitalization"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalKeyboardEnableCapitalization MUST be set to true to make this field effective.
	PhysicalKeyboardEnableCapitalization bool
	// PhysicalKeyboardEnableDiacriticsOnLongpress is "InputMethodSettings.physicalKeyboardEnableDiacriticsOnLongpress"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalKeyboardEnableDiacriticsOnLongpress MUST be set to true to make this field effective.
	PhysicalKeyboardEnableDiacriticsOnLongpress bool
	// PhysicalKeyboardEnablePredictiveWriting is "InputMethodSettings.physicalKeyboardEnablePredictiveWriting"
	//
	// Optional
	//
	// NOTE: FFI_USE_PhysicalKeyboardEnablePredictiveWriting MUST be set to true to make this field effective.
	PhysicalKeyboardEnablePredictiveWriting bool
	// PinyinChinesePunctuation is "InputMethodSettings.pinyinChinesePunctuation"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinChinesePunctuation MUST be set to true to make this field effective.
	PinyinChinesePunctuation bool
	// PinyinDefaultChinese is "InputMethodSettings.pinyinDefaultChinese"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinDefaultChinese MUST be set to true to make this field effective.
	PinyinDefaultChinese bool
	// PinyinEnableFuzzy is "InputMethodSettings.pinyinEnableFuzzy"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinEnableFuzzy MUST be set to true to make this field effective.
	PinyinEnableFuzzy bool
	// PinyinEnableLowerPaging is "InputMethodSettings.pinyinEnableLowerPaging"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinEnableLowerPaging MUST be set to true to make this field effective.
	PinyinEnableLowerPaging bool
	// PinyinEnableUpperPaging is "InputMethodSettings.pinyinEnableUpperPaging"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinEnableUpperPaging MUST be set to true to make this field effective.
	PinyinEnableUpperPaging bool
	// PinyinFullWidthCharacter is "InputMethodSettings.pinyinFullWidthCharacter"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinyinFullWidthCharacter MUST be set to true to make this field effective.
	PinyinFullWidthCharacter bool
	// PinyinFuzzyConfig is "InputMethodSettings.pinyinFuzzyConfig"
	//
	// Optional
	//
	// NOTE: PinyinFuzzyConfig.FFI_USE MUST be set to true to get PinyinFuzzyConfig used.
	PinyinFuzzyConfig InputMethodSettingsFieldPinyinFuzzyConfig
	// VietnameseTelexAllowFlexibleDiacritics is "InputMethodSettings.vietnameseTelexAllowFlexibleDiacritics"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseTelexAllowFlexibleDiacritics MUST be set to true to make this field effective.
	VietnameseTelexAllowFlexibleDiacritics bool
	// VietnameseTelexInsertDoubleHornOnUo is "InputMethodSettings.vietnameseTelexInsertDoubleHornOnUo"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseTelexInsertDoubleHornOnUo MUST be set to true to make this field effective.
	VietnameseTelexInsertDoubleHornOnUo bool
	// VietnameseTelexInsertUHornOnW is "InputMethodSettings.vietnameseTelexInsertUHornOnW"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseTelexInsertUHornOnW MUST be set to true to make this field effective.
	VietnameseTelexInsertUHornOnW bool
	// VietnameseTelexNewStyleToneMarkPlacement is "InputMethodSettings.vietnameseTelexNewStyleToneMarkPlacement"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseTelexNewStyleToneMarkPlacement MUST be set to true to make this field effective.
	VietnameseTelexNewStyleToneMarkPlacement bool
	// VietnameseTelexShowUnderline is "InputMethodSettings.vietnameseTelexShowUnderline"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseTelexShowUnderline MUST be set to true to make this field effective.
	VietnameseTelexShowUnderline bool
	// VietnameseVniAllowFlexibleDiacritics is "InputMethodSettings.vietnameseVniAllowFlexibleDiacritics"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseVniAllowFlexibleDiacritics MUST be set to true to make this field effective.
	VietnameseVniAllowFlexibleDiacritics bool
	// VietnameseVniInsertDoubleHornOnUo is "InputMethodSettings.vietnameseVniInsertDoubleHornOnUo"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseVniInsertDoubleHornOnUo MUST be set to true to make this field effective.
	VietnameseVniInsertDoubleHornOnUo bool
	// VietnameseVniNewStyleToneMarkPlacement is "InputMethodSettings.vietnameseVniNewStyleToneMarkPlacement"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseVniNewStyleToneMarkPlacement MUST be set to true to make this field effective.
	VietnameseVniNewStyleToneMarkPlacement bool
	// VietnameseVniShowUnderline is "InputMethodSettings.vietnameseVniShowUnderline"
	//
	// Optional
	//
	// NOTE: FFI_USE_VietnameseVniShowUnderline MUST be set to true to make this field effective.
	VietnameseVniShowUnderline bool
	// VirtualKeyboardAutoCorrectionLevel is "InputMethodSettings.virtualKeyboardAutoCorrectionLevel"
	//
	// Optional
	//
	// NOTE: FFI_USE_VirtualKeyboardAutoCorrectionLevel MUST be set to true to make this field effective.
	VirtualKeyboardAutoCorrectionLevel int64
	// VirtualKeyboardEnableCapitalization is "InputMethodSettings.virtualKeyboardEnableCapitalization"
	//
	// Optional
	//
	// NOTE: FFI_USE_VirtualKeyboardEnableCapitalization MUST be set to true to make this field effective.
	VirtualKeyboardEnableCapitalization bool
	// XkbLayout is "InputMethodSettings.xkbLayout"
	//
	// Optional
	XkbLayout js.String
	// ZhuyinKeyboardLayout is "InputMethodSettings.zhuyinKeyboardLayout"
	//
	// Optional
	ZhuyinKeyboardLayout js.String
	// ZhuyinPageSize is "InputMethodSettings.zhuyinPageSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_ZhuyinPageSize MUST be set to true to make this field effective.
	ZhuyinPageSize int64
	// ZhuyinSelectKeys is "InputMethodSettings.zhuyinSelectKeys"
	//
	// Optional
	ZhuyinSelectKeys js.String

	FFI_USE_EnableCompletion                               bool // for EnableCompletion.
	FFI_USE_EnableDoubleSpacePeriod                        bool // for EnableDoubleSpacePeriod.
	FFI_USE_EnableGestureTyping                            bool // for EnableGestureTyping.
	FFI_USE_EnablePrediction                               bool // for EnablePrediction.
	FFI_USE_EnableSoundOnKeypress                          bool // for EnableSoundOnKeypress.
	FFI_USE_KoreanEnableSyllableInput                      bool // for KoreanEnableSyllableInput.
	FFI_USE_KoreanShowHangulCandidate                      bool // for KoreanShowHangulCandidate.
	FFI_USE_PhysicalKeyboardAutoCorrectionEnabledByDefault bool // for PhysicalKeyboardAutoCorrectionEnabledByDefault.
	FFI_USE_PhysicalKeyboardAutoCorrectionLevel            bool // for PhysicalKeyboardAutoCorrectionLevel.
	FFI_USE_PhysicalKeyboardEnableCapitalization           bool // for PhysicalKeyboardEnableCapitalization.
	FFI_USE_PhysicalKeyboardEnableDiacriticsOnLongpress    bool // for PhysicalKeyboardEnableDiacriticsOnLongpress.
	FFI_USE_PhysicalKeyboardEnablePredictiveWriting        bool // for PhysicalKeyboardEnablePredictiveWriting.
	FFI_USE_PinyinChinesePunctuation                       bool // for PinyinChinesePunctuation.
	FFI_USE_PinyinDefaultChinese                           bool // for PinyinDefaultChinese.
	FFI_USE_PinyinEnableFuzzy                              bool // for PinyinEnableFuzzy.
	FFI_USE_PinyinEnableLowerPaging                        bool // for PinyinEnableLowerPaging.
	FFI_USE_PinyinEnableUpperPaging                        bool // for PinyinEnableUpperPaging.
	FFI_USE_PinyinFullWidthCharacter                       bool // for PinyinFullWidthCharacter.
	FFI_USE_VietnameseTelexAllowFlexibleDiacritics         bool // for VietnameseTelexAllowFlexibleDiacritics.
	FFI_USE_VietnameseTelexInsertDoubleHornOnUo            bool // for VietnameseTelexInsertDoubleHornOnUo.
	FFI_USE_VietnameseTelexInsertUHornOnW                  bool // for VietnameseTelexInsertUHornOnW.
	FFI_USE_VietnameseTelexNewStyleToneMarkPlacement       bool // for VietnameseTelexNewStyleToneMarkPlacement.
	FFI_USE_VietnameseTelexShowUnderline                   bool // for VietnameseTelexShowUnderline.
	FFI_USE_VietnameseVniAllowFlexibleDiacritics           bool // for VietnameseVniAllowFlexibleDiacritics.
	FFI_USE_VietnameseVniInsertDoubleHornOnUo              bool // for VietnameseVniInsertDoubleHornOnUo.
	FFI_USE_VietnameseVniNewStyleToneMarkPlacement         bool // for VietnameseVniNewStyleToneMarkPlacement.
	FFI_USE_VietnameseVniShowUnderline                     bool // for VietnameseVniShowUnderline.
	FFI_USE_VirtualKeyboardAutoCorrectionLevel             bool // for VirtualKeyboardAutoCorrectionLevel.
	FFI_USE_VirtualKeyboardEnableCapitalization            bool // for VirtualKeyboardEnableCapitalization.
	FFI_USE_ZhuyinPageSize                                 bool // for ZhuyinPageSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InputMethodSettings with all fields set.
func (p InputMethodSettings) FromRef(ref js.Ref) InputMethodSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InputMethodSettings in the application heap.
func (p InputMethodSettings) New() js.Ref {
	return bindings.InputMethodSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InputMethodSettings) UpdateFrom(ref js.Ref) {
	bindings.InputMethodSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InputMethodSettings) Update(ref js.Ref) {
	bindings.InputMethodSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InputMethodSettings) FreeMembers(recursive bool) {
	js.Free(
		p.KoreanKeyboardLayout.Ref(),
		p.XkbLayout.Ref(),
		p.ZhuyinKeyboardLayout.Ref(),
		p.ZhuyinSelectKeys.Ref(),
	)
	p.KoreanKeyboardLayout = p.KoreanKeyboardLayout.FromRef(js.Undefined)
	p.XkbLayout = p.XkbLayout.FromRef(js.Undefined)
	p.ZhuyinKeyboardLayout = p.ZhuyinKeyboardLayout.FromRef(js.Undefined)
	p.ZhuyinSelectKeys = p.ZhuyinSelectKeys.FromRef(js.Undefined)
	if recursive {
		p.PinyinFuzzyConfig.FreeMembers(true)
	}
}

type MenuItemStyle uint32

const (
	_ MenuItemStyle = iota

	MenuItemStyle_CHECK
	MenuItemStyle_RADIO
	MenuItemStyle_SEPARATOR
)

func (MenuItemStyle) FromRef(str js.Ref) MenuItemStyle {
	return MenuItemStyle(bindings.ConstOfMenuItemStyle(str))
}

func (x MenuItemStyle) String() (string, bool) {
	switch x {
	case MenuItemStyle_CHECK:
		return "check", true
	case MenuItemStyle_RADIO:
		return "radio", true
	case MenuItemStyle_SEPARATOR:
		return "separator", true
	default:
		return "", false
	}
}

type MenuItem struct {
	// Checked is "MenuItem.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Enabled is "MenuItem.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Id is "MenuItem.id"
	//
	// Required
	Id js.String
	// Label is "MenuItem.label"
	//
	// Optional
	Label js.String
	// Style is "MenuItem.style"
	//
	// Optional
	Style MenuItemStyle
	// Visible is "MenuItem.visible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool

	FFI_USE_Checked bool // for Checked.
	FFI_USE_Enabled bool // for Enabled.
	FFI_USE_Visible bool // for Visible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MenuItem with all fields set.
func (p MenuItem) FromRef(ref js.Ref) MenuItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MenuItem in the application heap.
func (p MenuItem) New() js.Ref {
	return bindings.MenuItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MenuItem) UpdateFrom(ref js.Ref) {
	bindings.MenuItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MenuItem) Update(ref js.Ref) {
	bindings.MenuItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MenuItem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Label.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type OnAutocorrectArgParameters struct {
	// ContextID is "OnAutocorrectArgParameters.contextID"
	//
	// Required
	ContextID int64
	// CorrectedWord is "OnAutocorrectArgParameters.correctedWord"
	//
	// Required
	CorrectedWord js.String
	// StartIndex is "OnAutocorrectArgParameters.startIndex"
	//
	// Required
	StartIndex int64
	// TypedWord is "OnAutocorrectArgParameters.typedWord"
	//
	// Required
	TypedWord js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnAutocorrectArgParameters with all fields set.
func (p OnAutocorrectArgParameters) FromRef(ref js.Ref) OnAutocorrectArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnAutocorrectArgParameters in the application heap.
func (p OnAutocorrectArgParameters) New() js.Ref {
	return bindings.OnAutocorrectArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnAutocorrectArgParameters) UpdateFrom(ref js.Ref) {
	bindings.OnAutocorrectArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnAutocorrectArgParameters) Update(ref js.Ref) {
	bindings.OnAutocorrectArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnAutocorrectArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.CorrectedWord.Ref(),
		p.TypedWord.Ref(),
	)
	p.CorrectedWord = p.CorrectedWord.FromRef(js.Undefined)
	p.TypedWord = p.TypedWord.FromRef(js.Undefined)
}

type OnCaretBoundsChangedArgCaretBounds struct {
	// H is "OnCaretBoundsChangedArgCaretBounds.h"
	//
	// Required
	H int64
	// W is "OnCaretBoundsChangedArgCaretBounds.w"
	//
	// Required
	W int64
	// X is "OnCaretBoundsChangedArgCaretBounds.x"
	//
	// Required
	X int64
	// Y is "OnCaretBoundsChangedArgCaretBounds.y"
	//
	// Required
	Y int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnCaretBoundsChangedArgCaretBounds with all fields set.
func (p OnCaretBoundsChangedArgCaretBounds) FromRef(ref js.Ref) OnCaretBoundsChangedArgCaretBounds {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnCaretBoundsChangedArgCaretBounds in the application heap.
func (p OnCaretBoundsChangedArgCaretBounds) New() js.Ref {
	return bindings.OnCaretBoundsChangedArgCaretBoundsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnCaretBoundsChangedArgCaretBounds) UpdateFrom(ref js.Ref) {
	bindings.OnCaretBoundsChangedArgCaretBoundsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnCaretBoundsChangedArgCaretBounds) Update(ref js.Ref) {
	bindings.OnCaretBoundsChangedArgCaretBoundsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnCaretBoundsChangedArgCaretBounds) FreeMembers(recursive bool) {
}

type UnderlineStyle uint32

const (
	_ UnderlineStyle = iota

	UnderlineStyle_UNDERLINE
	UnderlineStyle_DOUBLE_UNDERLINE
	UnderlineStyle_NO_UNDERLINE
)

func (UnderlineStyle) FromRef(str js.Ref) UnderlineStyle {
	return UnderlineStyle(bindings.ConstOfUnderlineStyle(str))
}

func (x UnderlineStyle) String() (string, bool) {
	switch x {
	case UnderlineStyle_UNDERLINE:
		return "underline", true
	case UnderlineStyle_DOUBLE_UNDERLINE:
		return "doubleUnderline", true
	case UnderlineStyle_NO_UNDERLINE:
		return "noUnderline", true
	default:
		return "", false
	}
}

type SetCompositionRangeArgParametersFieldSegmentsElem struct {
	// End is "SetCompositionRangeArgParametersFieldSegmentsElem.end"
	//
	// Required
	End int64
	// Start is "SetCompositionRangeArgParametersFieldSegmentsElem.start"
	//
	// Required
	Start int64
	// Style is "SetCompositionRangeArgParametersFieldSegmentsElem.style"
	//
	// Required
	Style UnderlineStyle

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCompositionRangeArgParametersFieldSegmentsElem with all fields set.
func (p SetCompositionRangeArgParametersFieldSegmentsElem) FromRef(ref js.Ref) SetCompositionRangeArgParametersFieldSegmentsElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCompositionRangeArgParametersFieldSegmentsElem in the application heap.
func (p SetCompositionRangeArgParametersFieldSegmentsElem) New() js.Ref {
	return bindings.SetCompositionRangeArgParametersFieldSegmentsElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCompositionRangeArgParametersFieldSegmentsElem) UpdateFrom(ref js.Ref) {
	bindings.SetCompositionRangeArgParametersFieldSegmentsElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCompositionRangeArgParametersFieldSegmentsElem) Update(ref js.Ref) {
	bindings.SetCompositionRangeArgParametersFieldSegmentsElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCompositionRangeArgParametersFieldSegmentsElem) FreeMembers(recursive bool) {
}

type SetCompositionRangeArgParameters struct {
	// ContextID is "SetCompositionRangeArgParameters.contextID"
	//
	// Required
	ContextID int64
	// Segments is "SetCompositionRangeArgParameters.segments"
	//
	// Optional
	Segments js.Array[SetCompositionRangeArgParametersFieldSegmentsElem]
	// SelectionAfter is "SetCompositionRangeArgParameters.selectionAfter"
	//
	// Required
	SelectionAfter int64
	// SelectionBefore is "SetCompositionRangeArgParameters.selectionBefore"
	//
	// Required
	SelectionBefore int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCompositionRangeArgParameters with all fields set.
func (p SetCompositionRangeArgParameters) FromRef(ref js.Ref) SetCompositionRangeArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCompositionRangeArgParameters in the application heap.
func (p SetCompositionRangeArgParameters) New() js.Ref {
	return bindings.SetCompositionRangeArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCompositionRangeArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetCompositionRangeArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCompositionRangeArgParameters) Update(ref js.Ref) {
	bindings.SetCompositionRangeArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCompositionRangeArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Segments.Ref(),
	)
	p.Segments = p.Segments.FromRef(js.Undefined)
}

// HasFuncAddWordToDictionary returns true if the function "WEBEXT.inputMethodPrivate.addWordToDictionary" exists.
func HasFuncAddWordToDictionary() bool {
	return js.True == bindings.HasFuncAddWordToDictionary()
}

// FuncAddWordToDictionary returns the function "WEBEXT.inputMethodPrivate.addWordToDictionary".
func FuncAddWordToDictionary() (fn js.Func[func(word js.String) js.Promise[js.Void]]) {
	bindings.FuncAddWordToDictionary(
		js.Pointer(&fn),
	)
	return
}

// AddWordToDictionary calls the function "WEBEXT.inputMethodPrivate.addWordToDictionary" directly.
func AddWordToDictionary(word js.String) (ret js.Promise[js.Void]) {
	bindings.CallAddWordToDictionary(
		js.Pointer(&ret),
		word.Ref(),
	)

	return
}

// TryAddWordToDictionary calls the function "WEBEXT.inputMethodPrivate.addWordToDictionary"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddWordToDictionary(word js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddWordToDictionary(
		js.Pointer(&ret), js.Pointer(&exception),
		word.Ref(),
	)

	return
}

// HasFuncFetchAllDictionaryWords returns true if the function "WEBEXT.inputMethodPrivate.fetchAllDictionaryWords" exists.
func HasFuncFetchAllDictionaryWords() bool {
	return js.True == bindings.HasFuncFetchAllDictionaryWords()
}

// FuncFetchAllDictionaryWords returns the function "WEBEXT.inputMethodPrivate.fetchAllDictionaryWords".
func FuncFetchAllDictionaryWords() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncFetchAllDictionaryWords(
		js.Pointer(&fn),
	)
	return
}

// FetchAllDictionaryWords calls the function "WEBEXT.inputMethodPrivate.fetchAllDictionaryWords" directly.
func FetchAllDictionaryWords() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallFetchAllDictionaryWords(
		js.Pointer(&ret),
	)

	return
}

// TryFetchAllDictionaryWords calls the function "WEBEXT.inputMethodPrivate.fetchAllDictionaryWords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFetchAllDictionaryWords() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFetchAllDictionaryWords(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFinishComposingText returns true if the function "WEBEXT.inputMethodPrivate.finishComposingText" exists.
func HasFuncFinishComposingText() bool {
	return js.True == bindings.HasFuncFinishComposingText()
}

// FuncFinishComposingText returns the function "WEBEXT.inputMethodPrivate.finishComposingText".
func FuncFinishComposingText() (fn js.Func[func(parameters FinishComposingTextArgParameters) js.Promise[js.Void]]) {
	bindings.FuncFinishComposingText(
		js.Pointer(&fn),
	)
	return
}

// FinishComposingText calls the function "WEBEXT.inputMethodPrivate.finishComposingText" directly.
func FinishComposingText(parameters FinishComposingTextArgParameters) (ret js.Promise[js.Void]) {
	bindings.CallFinishComposingText(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryFinishComposingText calls the function "WEBEXT.inputMethodPrivate.finishComposingText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFinishComposingText(parameters FinishComposingTextArgParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFinishComposingText(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncGetCurrentInputMethod returns true if the function "WEBEXT.inputMethodPrivate.getCurrentInputMethod" exists.
func HasFuncGetCurrentInputMethod() bool {
	return js.True == bindings.HasFuncGetCurrentInputMethod()
}

// FuncGetCurrentInputMethod returns the function "WEBEXT.inputMethodPrivate.getCurrentInputMethod".
func FuncGetCurrentInputMethod() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetCurrentInputMethod(
		js.Pointer(&fn),
	)
	return
}

// GetCurrentInputMethod calls the function "WEBEXT.inputMethodPrivate.getCurrentInputMethod" directly.
func GetCurrentInputMethod() (ret js.Promise[js.String]) {
	bindings.CallGetCurrentInputMethod(
		js.Pointer(&ret),
	)

	return
}

// TryGetCurrentInputMethod calls the function "WEBEXT.inputMethodPrivate.getCurrentInputMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCurrentInputMethod() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCurrentInputMethod(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInputMethodConfig returns true if the function "WEBEXT.inputMethodPrivate.getInputMethodConfig" exists.
func HasFuncGetInputMethodConfig() bool {
	return js.True == bindings.HasFuncGetInputMethodConfig()
}

// FuncGetInputMethodConfig returns the function "WEBEXT.inputMethodPrivate.getInputMethodConfig".
func FuncGetInputMethodConfig() (fn js.Func[func() js.Promise[GetInputMethodConfigReturnType]]) {
	bindings.FuncGetInputMethodConfig(
		js.Pointer(&fn),
	)
	return
}

// GetInputMethodConfig calls the function "WEBEXT.inputMethodPrivate.getInputMethodConfig" directly.
func GetInputMethodConfig() (ret js.Promise[GetInputMethodConfigReturnType]) {
	bindings.CallGetInputMethodConfig(
		js.Pointer(&ret),
	)

	return
}

// TryGetInputMethodConfig calls the function "WEBEXT.inputMethodPrivate.getInputMethodConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInputMethodConfig() (ret js.Promise[GetInputMethodConfigReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInputMethodConfig(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInputMethods returns true if the function "WEBEXT.inputMethodPrivate.getInputMethods" exists.
func HasFuncGetInputMethods() bool {
	return js.True == bindings.HasFuncGetInputMethods()
}

// FuncGetInputMethods returns the function "WEBEXT.inputMethodPrivate.getInputMethods".
func FuncGetInputMethods() (fn js.Func[func() js.Promise[js.Array[GetInputMethodsReturnTypeElem]]]) {
	bindings.FuncGetInputMethods(
		js.Pointer(&fn),
	)
	return
}

// GetInputMethods calls the function "WEBEXT.inputMethodPrivate.getInputMethods" directly.
func GetInputMethods() (ret js.Promise[js.Array[GetInputMethodsReturnTypeElem]]) {
	bindings.CallGetInputMethods(
		js.Pointer(&ret),
	)

	return
}

// TryGetInputMethods calls the function "WEBEXT.inputMethodPrivate.getInputMethods"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInputMethods() (ret js.Promise[js.Array[GetInputMethodsReturnTypeElem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInputMethods(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSettings returns true if the function "WEBEXT.inputMethodPrivate.getSettings" exists.
func HasFuncGetSettings() bool {
	return js.True == bindings.HasFuncGetSettings()
}

// FuncGetSettings returns the function "WEBEXT.inputMethodPrivate.getSettings".
func FuncGetSettings() (fn js.Func[func(engineID js.String) js.Promise[InputMethodSettings]]) {
	bindings.FuncGetSettings(
		js.Pointer(&fn),
	)
	return
}

// GetSettings calls the function "WEBEXT.inputMethodPrivate.getSettings" directly.
func GetSettings(engineID js.String) (ret js.Promise[InputMethodSettings]) {
	bindings.CallGetSettings(
		js.Pointer(&ret),
		engineID.Ref(),
	)

	return
}

// TryGetSettings calls the function "WEBEXT.inputMethodPrivate.getSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSettings(engineID js.String) (ret js.Promise[InputMethodSettings], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		engineID.Ref(),
	)

	return
}

// HasFuncGetSurroundingText returns true if the function "WEBEXT.inputMethodPrivate.getSurroundingText" exists.
func HasFuncGetSurroundingText() bool {
	return js.True == bindings.HasFuncGetSurroundingText()
}

// FuncGetSurroundingText returns the function "WEBEXT.inputMethodPrivate.getSurroundingText".
func FuncGetSurroundingText() (fn js.Func[func(beforeLength int64, afterLength int64) js.Promise[GetSurroundingTextReturnType]]) {
	bindings.FuncGetSurroundingText(
		js.Pointer(&fn),
	)
	return
}

// GetSurroundingText calls the function "WEBEXT.inputMethodPrivate.getSurroundingText" directly.
func GetSurroundingText(beforeLength int64, afterLength int64) (ret js.Promise[GetSurroundingTextReturnType]) {
	bindings.CallGetSurroundingText(
		js.Pointer(&ret),
		float64(beforeLength),
		float64(afterLength),
	)

	return
}

// TryGetSurroundingText calls the function "WEBEXT.inputMethodPrivate.getSurroundingText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSurroundingText(beforeLength int64, afterLength int64) (ret js.Promise[GetSurroundingTextReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSurroundingText(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(beforeLength),
		float64(afterLength),
	)

	return
}

// HasFuncGetTextFieldBounds returns true if the function "WEBEXT.inputMethodPrivate.getTextFieldBounds" exists.
func HasFuncGetTextFieldBounds() bool {
	return js.True == bindings.HasFuncGetTextFieldBounds()
}

// FuncGetTextFieldBounds returns the function "WEBEXT.inputMethodPrivate.getTextFieldBounds".
func FuncGetTextFieldBounds() (fn js.Func[func(parameters GetTextFieldBoundsArgParameters) js.Promise[GetTextFieldBoundsReturnType]]) {
	bindings.FuncGetTextFieldBounds(
		js.Pointer(&fn),
	)
	return
}

// GetTextFieldBounds calls the function "WEBEXT.inputMethodPrivate.getTextFieldBounds" directly.
func GetTextFieldBounds(parameters GetTextFieldBoundsArgParameters) (ret js.Promise[GetTextFieldBoundsReturnType]) {
	bindings.CallGetTextFieldBounds(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryGetTextFieldBounds calls the function "WEBEXT.inputMethodPrivate.getTextFieldBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTextFieldBounds(parameters GetTextFieldBoundsArgParameters) (ret js.Promise[GetTextFieldBoundsReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTextFieldBounds(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncHideInputView returns true if the function "WEBEXT.inputMethodPrivate.hideInputView" exists.
func HasFuncHideInputView() bool {
	return js.True == bindings.HasFuncHideInputView()
}

// FuncHideInputView returns the function "WEBEXT.inputMethodPrivate.hideInputView".
func FuncHideInputView() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHideInputView(
		js.Pointer(&fn),
	)
	return
}

// HideInputView calls the function "WEBEXT.inputMethodPrivate.hideInputView" directly.
func HideInputView() (ret js.Promise[js.Void]) {
	bindings.CallHideInputView(
		js.Pointer(&ret),
	)

	return
}

// TryHideInputView calls the function "WEBEXT.inputMethodPrivate.hideInputView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHideInputView() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHideInputView(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNotifyInputMethodReadyForTesting returns true if the function "WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting" exists.
func HasFuncNotifyInputMethodReadyForTesting() bool {
	return js.True == bindings.HasFuncNotifyInputMethodReadyForTesting()
}

// FuncNotifyInputMethodReadyForTesting returns the function "WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting".
func FuncNotifyInputMethodReadyForTesting() (fn js.Func[func()]) {
	bindings.FuncNotifyInputMethodReadyForTesting(
		js.Pointer(&fn),
	)
	return
}

// NotifyInputMethodReadyForTesting calls the function "WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting" directly.
func NotifyInputMethodReadyForTesting() (ret js.Void) {
	bindings.CallNotifyInputMethodReadyForTesting(
		js.Pointer(&ret),
	)

	return
}

// TryNotifyInputMethodReadyForTesting calls the function "WEBEXT.inputMethodPrivate.notifyInputMethodReadyForTesting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyInputMethodReadyForTesting() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyInputMethodReadyForTesting(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOnAutocorrect returns true if the function "WEBEXT.inputMethodPrivate.onAutocorrect" exists.
func HasFuncOnAutocorrect() bool {
	return js.True == bindings.HasFuncOnAutocorrect()
}

// FuncOnAutocorrect returns the function "WEBEXT.inputMethodPrivate.onAutocorrect".
func FuncOnAutocorrect() (fn js.Func[func(parameters OnAutocorrectArgParameters)]) {
	bindings.FuncOnAutocorrect(
		js.Pointer(&fn),
	)
	return
}

// OnAutocorrect calls the function "WEBEXT.inputMethodPrivate.onAutocorrect" directly.
func OnAutocorrect(parameters OnAutocorrectArgParameters) (ret js.Void) {
	bindings.CallOnAutocorrect(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryOnAutocorrect calls the function "WEBEXT.inputMethodPrivate.onAutocorrect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAutocorrect(parameters OnAutocorrectArgParameters) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAutocorrect(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

type OnCaretBoundsChangedEventCallbackFunc func(this js.Ref, caretBounds *OnCaretBoundsChangedArgCaretBounds) js.Ref

func (fn OnCaretBoundsChangedEventCallbackFunc) Register() js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)] {
	return js.RegisterCallback[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCaretBoundsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCaretBoundsChangedArgCaretBounds
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

type OnCaretBoundsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, caretBounds *OnCaretBoundsChangedArgCaretBounds) js.Ref
	Arg T
}

func (cb *OnCaretBoundsChangedEventCallback[T]) Register() js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)] {
	return js.RegisterCallback[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCaretBoundsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCaretBoundsChangedArgCaretBounds
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

// HasFuncOnCaretBoundsChanged returns true if the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener" exists.
func HasFuncOnCaretBoundsChanged() bool {
	return js.True == bindings.HasFuncOnCaretBoundsChanged()
}

// FuncOnCaretBoundsChanged returns the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener".
func FuncOnCaretBoundsChanged() (fn js.Func[func(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)])]) {
	bindings.FuncOnCaretBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener" directly.
func OnCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret js.Void) {
	bindings.CallOnCaretBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCaretBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCaretBoundsChanged returns true if the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener" exists.
func HasFuncOffCaretBoundsChanged() bool {
	return js.True == bindings.HasFuncOffCaretBoundsChanged()
}

// FuncOffCaretBoundsChanged returns the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener".
func FuncOffCaretBoundsChanged() (fn js.Func[func(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)])]) {
	bindings.FuncOffCaretBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener" directly.
func OffCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret js.Void) {
	bindings.CallOffCaretBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCaretBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCaretBoundsChanged returns true if the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener" exists.
func HasFuncHasOnCaretBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnCaretBoundsChanged()
}

// FuncHasOnCaretBoundsChanged returns the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener".
func FuncHasOnCaretBoundsChanged() (fn js.Func[func(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) bool]) {
	bindings.FuncHasOnCaretBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener" directly.
func HasOnCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret bool) {
	bindings.CallHasOnCaretBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCaretBoundsChanged calls the function "WEBEXT.inputMethodPrivate.onCaretBoundsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCaretBoundsChanged(callback js.Func[func(caretBounds *OnCaretBoundsChangedArgCaretBounds)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCaretBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnChangedEventCallbackFunc func(this js.Ref, newInputMethodId js.String) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func(newInputMethodId js.String)] {
	return js.RegisterCallback[func(newInputMethodId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
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

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, newInputMethodId js.String) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func(newInputMethodId js.String)] {
	return js.RegisterCallback[func(newInputMethodId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnChanged returns true if the function "WEBEXT.inputMethodPrivate.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.inputMethodPrivate.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func(newInputMethodId js.String)])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.addListener" directly.
func OnChanged(callback js.Func[func(newInputMethodId js.String)]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func(newInputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.inputMethodPrivate.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.inputMethodPrivate.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func(newInputMethodId js.String)])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func(newInputMethodId js.String)]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func(newInputMethodId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.inputMethodPrivate.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.inputMethodPrivate.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func(newInputMethodId js.String)]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func(newInputMethodId js.String)]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.inputMethodPrivate.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func(newInputMethodId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDictionaryChangedEventCallbackFunc func(this js.Ref, added js.Array[js.String], removed js.Array[js.String]) js.Ref

func (fn OnDictionaryChangedEventCallbackFunc) Register() js.Func[func(added js.Array[js.String], removed js.Array[js.String])] {
	return js.RegisterCallback[func(added js.Array[js.String], removed js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDictionaryChangedEventCallbackFunc) DispatchCallback(
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

type OnDictionaryChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, added js.Array[js.String], removed js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnDictionaryChangedEventCallback[T]) Register() js.Func[func(added js.Array[js.String], removed js.Array[js.String])] {
	return js.RegisterCallback[func(added js.Array[js.String], removed js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDictionaryChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDictionaryChanged returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener" exists.
func HasFuncOnDictionaryChanged() bool {
	return js.True == bindings.HasFuncOnDictionaryChanged()
}

// FuncOnDictionaryChanged returns the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener".
func FuncOnDictionaryChanged() (fn js.Func[func(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])])]) {
	bindings.FuncOnDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener" directly.
func OnDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDictionaryChanged returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener" exists.
func HasFuncOffDictionaryChanged() bool {
	return js.True == bindings.HasFuncOffDictionaryChanged()
}

// FuncOffDictionaryChanged returns the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener".
func FuncOffDictionaryChanged() (fn js.Func[func(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])])]) {
	bindings.FuncOffDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener" directly.
func OffDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDictionaryChanged returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener" exists.
func HasFuncHasOnDictionaryChanged() bool {
	return js.True == bindings.HasFuncHasOnDictionaryChanged()
}

// FuncHasOnDictionaryChanged returns the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener".
func FuncHasOnDictionaryChanged() (fn js.Func[func(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) bool]) {
	bindings.FuncHasOnDictionaryChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener" directly.
func HasOnDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnDictionaryChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDictionaryChanged calls the function "WEBEXT.inputMethodPrivate.onDictionaryChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDictionaryChanged(callback js.Func[func(added js.Array[js.String], removed js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDictionaryChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDictionaryLoadedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnDictionaryLoadedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDictionaryLoadedEventCallbackFunc) DispatchCallback(
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

type OnDictionaryLoadedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnDictionaryLoadedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDictionaryLoadedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDictionaryLoaded returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener" exists.
func HasFuncOnDictionaryLoaded() bool {
	return js.True == bindings.HasFuncOnDictionaryLoaded()
}

// FuncOnDictionaryLoaded returns the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener".
func FuncOnDictionaryLoaded() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnDictionaryLoaded(
		js.Pointer(&fn),
	)
	return
}

// OnDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener" directly.
func OnDictionaryLoaded(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnDictionaryLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDictionaryLoaded(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDictionaryLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDictionaryLoaded returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener" exists.
func HasFuncOffDictionaryLoaded() bool {
	return js.True == bindings.HasFuncOffDictionaryLoaded()
}

// FuncOffDictionaryLoaded returns the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener".
func FuncOffDictionaryLoaded() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffDictionaryLoaded(
		js.Pointer(&fn),
	)
	return
}

// OffDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener" directly.
func OffDictionaryLoaded(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffDictionaryLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDictionaryLoaded(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDictionaryLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDictionaryLoaded returns true if the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener" exists.
func HasFuncHasOnDictionaryLoaded() bool {
	return js.True == bindings.HasFuncHasOnDictionaryLoaded()
}

// FuncHasOnDictionaryLoaded returns the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener".
func FuncHasOnDictionaryLoaded() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnDictionaryLoaded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener" directly.
func HasOnDictionaryLoaded(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnDictionaryLoaded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDictionaryLoaded calls the function "WEBEXT.inputMethodPrivate.onDictionaryLoaded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDictionaryLoaded(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDictionaryLoaded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFocusEventCallbackFunc func(this js.Ref, context *InputContext) js.Ref

func (fn OnFocusEventCallbackFunc) Register() js.Func[func(context *InputContext)] {
	return js.RegisterCallback[func(context *InputContext)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFocusEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InputContext
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

type OnFocusEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, context *InputContext) js.Ref
	Arg T
}

func (cb *OnFocusEventCallback[T]) Register() js.Func[func(context *InputContext)] {
	return js.RegisterCallback[func(context *InputContext)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFocusEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InputContext
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

// HasFuncOnFocus returns true if the function "WEBEXT.inputMethodPrivate.onFocus.addListener" exists.
func HasFuncOnFocus() bool {
	return js.True == bindings.HasFuncOnFocus()
}

// FuncOnFocus returns the function "WEBEXT.inputMethodPrivate.onFocus.addListener".
func FuncOnFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOnFocus(
		js.Pointer(&fn),
	)
	return
}

// OnFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.addListener" directly.
func OnFocus(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOnFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFocus(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFocus returns true if the function "WEBEXT.inputMethodPrivate.onFocus.removeListener" exists.
func HasFuncOffFocus() bool {
	return js.True == bindings.HasFuncOffFocus()
}

// FuncOffFocus returns the function "WEBEXT.inputMethodPrivate.onFocus.removeListener".
func FuncOffFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOffFocus(
		js.Pointer(&fn),
	)
	return
}

// OffFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.removeListener" directly.
func OffFocus(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOffFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFocus(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFocus returns true if the function "WEBEXT.inputMethodPrivate.onFocus.hasListener" exists.
func HasFuncHasOnFocus() bool {
	return js.True == bindings.HasFuncHasOnFocus()
}

// FuncHasOnFocus returns the function "WEBEXT.inputMethodPrivate.onFocus.hasListener".
func FuncHasOnFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)]) bool]) {
	bindings.FuncHasOnFocus(
		js.Pointer(&fn),
	)
	return
}

// HasOnFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.hasListener" directly.
func HasOnFocus(callback js.Func[func(context *InputContext)]) (ret bool) {
	bindings.CallHasOnFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFocus calls the function "WEBEXT.inputMethodPrivate.onFocus.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFocus(callback js.Func[func(context *InputContext)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnImeMenuActivationChangedEventCallbackFunc func(this js.Ref, activation bool) js.Ref

func (fn OnImeMenuActivationChangedEventCallbackFunc) Register() js.Func[func(activation bool)] {
	return js.RegisterCallback[func(activation bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnImeMenuActivationChangedEventCallbackFunc) DispatchCallback(
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

type OnImeMenuActivationChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activation bool) js.Ref
	Arg T
}

func (cb *OnImeMenuActivationChangedEventCallback[T]) Register() js.Func[func(activation bool)] {
	return js.RegisterCallback[func(activation bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnImeMenuActivationChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnImeMenuActivationChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener" exists.
func HasFuncOnImeMenuActivationChanged() bool {
	return js.True == bindings.HasFuncOnImeMenuActivationChanged()
}

// FuncOnImeMenuActivationChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener".
func FuncOnImeMenuActivationChanged() (fn js.Func[func(callback js.Func[func(activation bool)])]) {
	bindings.FuncOnImeMenuActivationChanged(
		js.Pointer(&fn),
	)
	return
}

// OnImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener" directly.
func OnImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret js.Void) {
	bindings.CallOnImeMenuActivationChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnImeMenuActivationChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffImeMenuActivationChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener" exists.
func HasFuncOffImeMenuActivationChanged() bool {
	return js.True == bindings.HasFuncOffImeMenuActivationChanged()
}

// FuncOffImeMenuActivationChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener".
func FuncOffImeMenuActivationChanged() (fn js.Func[func(callback js.Func[func(activation bool)])]) {
	bindings.FuncOffImeMenuActivationChanged(
		js.Pointer(&fn),
	)
	return
}

// OffImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener" directly.
func OffImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret js.Void) {
	bindings.CallOffImeMenuActivationChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffImeMenuActivationChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnImeMenuActivationChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener" exists.
func HasFuncHasOnImeMenuActivationChanged() bool {
	return js.True == bindings.HasFuncHasOnImeMenuActivationChanged()
}

// FuncHasOnImeMenuActivationChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener".
func FuncHasOnImeMenuActivationChanged() (fn js.Func[func(callback js.Func[func(activation bool)]) bool]) {
	bindings.FuncHasOnImeMenuActivationChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener" directly.
func HasOnImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret bool) {
	bindings.CallHasOnImeMenuActivationChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnImeMenuActivationChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuActivationChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnImeMenuActivationChanged(callback js.Func[func(activation bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnImeMenuActivationChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnImeMenuItemsChangedEventCallbackFunc func(this js.Ref, engineID js.String, items js.Array[MenuItem]) js.Ref

func (fn OnImeMenuItemsChangedEventCallbackFunc) Register() js.Func[func(engineID js.String, items js.Array[MenuItem])] {
	return js.RegisterCallback[func(engineID js.String, items js.Array[MenuItem])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnImeMenuItemsChangedEventCallbackFunc) DispatchCallback(
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
		js.Array[MenuItem]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnImeMenuItemsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, items js.Array[MenuItem]) js.Ref
	Arg T
}

func (cb *OnImeMenuItemsChangedEventCallback[T]) Register() js.Func[func(engineID js.String, items js.Array[MenuItem])] {
	return js.RegisterCallback[func(engineID js.String, items js.Array[MenuItem])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnImeMenuItemsChangedEventCallback[T]) DispatchCallback(
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
		js.Array[MenuItem]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnImeMenuItemsChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener" exists.
func HasFuncOnImeMenuItemsChanged() bool {
	return js.True == bindings.HasFuncOnImeMenuItemsChanged()
}

// FuncOnImeMenuItemsChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener".
func FuncOnImeMenuItemsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, items js.Array[MenuItem])])]) {
	bindings.FuncOnImeMenuItemsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener" directly.
func OnImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret js.Void) {
	bindings.CallOnImeMenuItemsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnImeMenuItemsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffImeMenuItemsChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener" exists.
func HasFuncOffImeMenuItemsChanged() bool {
	return js.True == bindings.HasFuncOffImeMenuItemsChanged()
}

// FuncOffImeMenuItemsChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener".
func FuncOffImeMenuItemsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, items js.Array[MenuItem])])]) {
	bindings.FuncOffImeMenuItemsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener" directly.
func OffImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret js.Void) {
	bindings.CallOffImeMenuItemsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffImeMenuItemsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnImeMenuItemsChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener" exists.
func HasFuncHasOnImeMenuItemsChanged() bool {
	return js.True == bindings.HasFuncHasOnImeMenuItemsChanged()
}

// FuncHasOnImeMenuItemsChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener".
func FuncHasOnImeMenuItemsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) bool]) {
	bindings.FuncHasOnImeMenuItemsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener" directly.
func HasOnImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret bool) {
	bindings.CallHasOnImeMenuItemsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnImeMenuItemsChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuItemsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnImeMenuItemsChanged(callback js.Func[func(engineID js.String, items js.Array[MenuItem])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnImeMenuItemsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnImeMenuListChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnImeMenuListChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnImeMenuListChangedEventCallbackFunc) DispatchCallback(
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

type OnImeMenuListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnImeMenuListChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnImeMenuListChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnImeMenuListChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener" exists.
func HasFuncOnImeMenuListChanged() bool {
	return js.True == bindings.HasFuncOnImeMenuListChanged()
}

// FuncOnImeMenuListChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener".
func FuncOnImeMenuListChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnImeMenuListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener" directly.
func OnImeMenuListChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnImeMenuListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnImeMenuListChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnImeMenuListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffImeMenuListChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener" exists.
func HasFuncOffImeMenuListChanged() bool {
	return js.True == bindings.HasFuncOffImeMenuListChanged()
}

// FuncOffImeMenuListChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener".
func FuncOffImeMenuListChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffImeMenuListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener" directly.
func OffImeMenuListChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffImeMenuListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffImeMenuListChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffImeMenuListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnImeMenuListChanged returns true if the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener" exists.
func HasFuncHasOnImeMenuListChanged() bool {
	return js.True == bindings.HasFuncHasOnImeMenuListChanged()
}

// FuncHasOnImeMenuListChanged returns the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener".
func FuncHasOnImeMenuListChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnImeMenuListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener" directly.
func HasOnImeMenuListChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnImeMenuListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnImeMenuListChanged calls the function "WEBEXT.inputMethodPrivate.onImeMenuListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnImeMenuListChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnImeMenuListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputMethodOptionsChangedEventCallbackFunc func(this js.Ref, engineID js.String) js.Ref

func (fn OnInputMethodOptionsChangedEventCallbackFunc) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputMethodOptionsChangedEventCallbackFunc) DispatchCallback(
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

type OnInputMethodOptionsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String) js.Ref
	Arg T
}

func (cb *OnInputMethodOptionsChangedEventCallback[T]) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputMethodOptionsChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputMethodOptionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener" exists.
func HasFuncOnInputMethodOptionsChanged() bool {
	return js.True == bindings.HasFuncOnInputMethodOptionsChanged()
}

// FuncOnInputMethodOptionsChanged returns the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener".
func FuncOnInputMethodOptionsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOnInputMethodOptionsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener" directly.
func OnInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOnInputMethodOptionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputMethodOptionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputMethodOptionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener" exists.
func HasFuncOffInputMethodOptionsChanged() bool {
	return js.True == bindings.HasFuncOffInputMethodOptionsChanged()
}

// FuncOffInputMethodOptionsChanged returns the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener".
func FuncOffInputMethodOptionsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOffInputMethodOptionsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener" directly.
func OffInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOffInputMethodOptionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputMethodOptionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputMethodOptionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener" exists.
func HasFuncHasOnInputMethodOptionsChanged() bool {
	return js.True == bindings.HasFuncHasOnInputMethodOptionsChanged()
}

// FuncHasOnInputMethodOptionsChanged returns the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener".
func FuncHasOnInputMethodOptionsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String)]) bool]) {
	bindings.FuncHasOnInputMethodOptionsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener" directly.
func HasOnInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret bool) {
	bindings.CallHasOnInputMethodOptionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputMethodOptionsChanged calls the function "WEBEXT.inputMethodPrivate.onInputMethodOptionsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputMethodOptionsChanged(callback js.Func[func(engineID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputMethodOptionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnScreenProjectionChangedEventCallbackFunc func(this js.Ref, isProjected bool) js.Ref

func (fn OnScreenProjectionChangedEventCallbackFunc) Register() js.Func[func(isProjected bool)] {
	return js.RegisterCallback[func(isProjected bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnScreenProjectionChangedEventCallbackFunc) DispatchCallback(
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

type OnScreenProjectionChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isProjected bool) js.Ref
	Arg T
}

func (cb *OnScreenProjectionChangedEventCallback[T]) Register() js.Func[func(isProjected bool)] {
	return js.RegisterCallback[func(isProjected bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnScreenProjectionChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnScreenProjectionChanged returns true if the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener" exists.
func HasFuncOnScreenProjectionChanged() bool {
	return js.True == bindings.HasFuncOnScreenProjectionChanged()
}

// FuncOnScreenProjectionChanged returns the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener".
func FuncOnScreenProjectionChanged() (fn js.Func[func(callback js.Func[func(isProjected bool)])]) {
	bindings.FuncOnScreenProjectionChanged(
		js.Pointer(&fn),
	)
	return
}

// OnScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener" directly.
func OnScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret js.Void) {
	bindings.CallOnScreenProjectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnScreenProjectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffScreenProjectionChanged returns true if the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener" exists.
func HasFuncOffScreenProjectionChanged() bool {
	return js.True == bindings.HasFuncOffScreenProjectionChanged()
}

// FuncOffScreenProjectionChanged returns the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener".
func FuncOffScreenProjectionChanged() (fn js.Func[func(callback js.Func[func(isProjected bool)])]) {
	bindings.FuncOffScreenProjectionChanged(
		js.Pointer(&fn),
	)
	return
}

// OffScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener" directly.
func OffScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret js.Void) {
	bindings.CallOffScreenProjectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffScreenProjectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnScreenProjectionChanged returns true if the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener" exists.
func HasFuncHasOnScreenProjectionChanged() bool {
	return js.True == bindings.HasFuncHasOnScreenProjectionChanged()
}

// FuncHasOnScreenProjectionChanged returns the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener".
func FuncHasOnScreenProjectionChanged() (fn js.Func[func(callback js.Func[func(isProjected bool)]) bool]) {
	bindings.FuncHasOnScreenProjectionChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener" directly.
func HasOnScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret bool) {
	bindings.CallHasOnScreenProjectionChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnScreenProjectionChanged calls the function "WEBEXT.inputMethodPrivate.onScreenProjectionChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnScreenProjectionChanged(callback js.Func[func(isProjected bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnScreenProjectionChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSettingsChangedEventCallbackFunc func(this js.Ref, engineID js.String, settings *InputMethodSettings) js.Ref

func (fn OnSettingsChangedEventCallbackFunc) Register() js.Func[func(engineID js.String, settings *InputMethodSettings)] {
	return js.RegisterCallback[func(engineID js.String, settings *InputMethodSettings)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSettingsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 InputMethodSettings
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSettingsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, settings *InputMethodSettings) js.Ref
	Arg T
}

func (cb *OnSettingsChangedEventCallback[T]) Register() js.Func[func(engineID js.String, settings *InputMethodSettings)] {
	return js.RegisterCallback[func(engineID js.String, settings *InputMethodSettings)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSettingsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 InputMethodSettings
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSettingsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSettingsChanged.addListener" exists.
func HasFuncOnSettingsChanged() bool {
	return js.True == bindings.HasFuncOnSettingsChanged()
}

// FuncOnSettingsChanged returns the function "WEBEXT.inputMethodPrivate.onSettingsChanged.addListener".
func FuncOnSettingsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, settings *InputMethodSettings)])]) {
	bindings.FuncOnSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.addListener" directly.
func OnSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret js.Void) {
	bindings.CallOnSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSettingsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener" exists.
func HasFuncOffSettingsChanged() bool {
	return js.True == bindings.HasFuncOffSettingsChanged()
}

// FuncOffSettingsChanged returns the function "WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener".
func FuncOffSettingsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, settings *InputMethodSettings)])]) {
	bindings.FuncOffSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener" directly.
func OffSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret js.Void) {
	bindings.CallOffSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSettingsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener" exists.
func HasFuncHasOnSettingsChanged() bool {
	return js.True == bindings.HasFuncHasOnSettingsChanged()
}

// FuncHasOnSettingsChanged returns the function "WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener".
func FuncHasOnSettingsChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) bool]) {
	bindings.FuncHasOnSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener" directly.
func HasOnSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret bool) {
	bindings.CallHasOnSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSettingsChanged calls the function "WEBEXT.inputMethodPrivate.onSettingsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSettingsChanged(callback js.Func[func(engineID js.String, settings *InputMethodSettings)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSuggestionsChangedEventCallbackFunc func(this js.Ref, suggestions js.Array[js.String]) js.Ref

func (fn OnSuggestionsChangedEventCallbackFunc) Register() js.Func[func(suggestions js.Array[js.String])] {
	return js.RegisterCallback[func(suggestions js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSuggestionsChangedEventCallbackFunc) DispatchCallback(
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

type OnSuggestionsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, suggestions js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnSuggestionsChangedEventCallback[T]) Register() js.Func[func(suggestions js.Array[js.String])] {
	return js.RegisterCallback[func(suggestions js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSuggestionsChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnSuggestionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener" exists.
func HasFuncOnSuggestionsChanged() bool {
	return js.True == bindings.HasFuncOnSuggestionsChanged()
}

// FuncOnSuggestionsChanged returns the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener".
func FuncOnSuggestionsChanged() (fn js.Func[func(callback js.Func[func(suggestions js.Array[js.String])])]) {
	bindings.FuncOnSuggestionsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener" directly.
func OnSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnSuggestionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSuggestionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSuggestionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener" exists.
func HasFuncOffSuggestionsChanged() bool {
	return js.True == bindings.HasFuncOffSuggestionsChanged()
}

// FuncOffSuggestionsChanged returns the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener".
func FuncOffSuggestionsChanged() (fn js.Func[func(callback js.Func[func(suggestions js.Array[js.String])])]) {
	bindings.FuncOffSuggestionsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener" directly.
func OffSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffSuggestionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSuggestionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSuggestionsChanged returns true if the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener" exists.
func HasFuncHasOnSuggestionsChanged() bool {
	return js.True == bindings.HasFuncHasOnSuggestionsChanged()
}

// FuncHasOnSuggestionsChanged returns the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener".
func FuncHasOnSuggestionsChanged() (fn js.Func[func(callback js.Func[func(suggestions js.Array[js.String])]) bool]) {
	bindings.FuncHasOnSuggestionsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener" directly.
func HasOnSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnSuggestionsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSuggestionsChanged calls the function "WEBEXT.inputMethodPrivate.onSuggestionsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSuggestionsChanged(callback js.Func[func(suggestions js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSuggestionsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchEventCallbackFunc func(this js.Ref, pointerType FocusReason) js.Ref

func (fn OnTouchEventCallbackFunc) Register() js.Func[func(pointerType FocusReason)] {
	return js.RegisterCallback[func(pointerType FocusReason)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		FocusReason(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnTouchEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, pointerType FocusReason) js.Ref
	Arg T
}

func (cb *OnTouchEventCallback[T]) Register() js.Func[func(pointerType FocusReason)] {
	return js.RegisterCallback[func(pointerType FocusReason)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchEventCallback[T]) DispatchCallback(
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

		FocusReason(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnTouch returns true if the function "WEBEXT.inputMethodPrivate.onTouch.addListener" exists.
func HasFuncOnTouch() bool {
	return js.True == bindings.HasFuncOnTouch()
}

// FuncOnTouch returns the function "WEBEXT.inputMethodPrivate.onTouch.addListener".
func FuncOnTouch() (fn js.Func[func(callback js.Func[func(pointerType FocusReason)])]) {
	bindings.FuncOnTouch(
		js.Pointer(&fn),
	)
	return
}

// OnTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.addListener" directly.
func OnTouch(callback js.Func[func(pointerType FocusReason)]) (ret js.Void) {
	bindings.CallOnTouch(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouch(callback js.Func[func(pointerType FocusReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouch(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouch returns true if the function "WEBEXT.inputMethodPrivate.onTouch.removeListener" exists.
func HasFuncOffTouch() bool {
	return js.True == bindings.HasFuncOffTouch()
}

// FuncOffTouch returns the function "WEBEXT.inputMethodPrivate.onTouch.removeListener".
func FuncOffTouch() (fn js.Func[func(callback js.Func[func(pointerType FocusReason)])]) {
	bindings.FuncOffTouch(
		js.Pointer(&fn),
	)
	return
}

// OffTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.removeListener" directly.
func OffTouch(callback js.Func[func(pointerType FocusReason)]) (ret js.Void) {
	bindings.CallOffTouch(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouch(callback js.Func[func(pointerType FocusReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouch(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouch returns true if the function "WEBEXT.inputMethodPrivate.onTouch.hasListener" exists.
func HasFuncHasOnTouch() bool {
	return js.True == bindings.HasFuncHasOnTouch()
}

// FuncHasOnTouch returns the function "WEBEXT.inputMethodPrivate.onTouch.hasListener".
func FuncHasOnTouch() (fn js.Func[func(callback js.Func[func(pointerType FocusReason)]) bool]) {
	bindings.FuncHasOnTouch(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.hasListener" directly.
func HasOnTouch(callback js.Func[func(pointerType FocusReason)]) (ret bool) {
	bindings.CallHasOnTouch(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouch calls the function "WEBEXT.inputMethodPrivate.onTouch.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouch(callback js.Func[func(pointerType FocusReason)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouch(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenOptionsPage returns true if the function "WEBEXT.inputMethodPrivate.openOptionsPage" exists.
func HasFuncOpenOptionsPage() bool {
	return js.True == bindings.HasFuncOpenOptionsPage()
}

// FuncOpenOptionsPage returns the function "WEBEXT.inputMethodPrivate.openOptionsPage".
func FuncOpenOptionsPage() (fn js.Func[func(inputMethodId js.String)]) {
	bindings.FuncOpenOptionsPage(
		js.Pointer(&fn),
	)
	return
}

// OpenOptionsPage calls the function "WEBEXT.inputMethodPrivate.openOptionsPage" directly.
func OpenOptionsPage(inputMethodId js.String) (ret js.Void) {
	bindings.CallOpenOptionsPage(
		js.Pointer(&ret),
		inputMethodId.Ref(),
	)

	return
}

// TryOpenOptionsPage calls the function "WEBEXT.inputMethodPrivate.openOptionsPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenOptionsPage(inputMethodId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenOptionsPage(
		js.Pointer(&ret), js.Pointer(&exception),
		inputMethodId.Ref(),
	)

	return
}

// HasFuncReset returns true if the function "WEBEXT.inputMethodPrivate.reset" exists.
func HasFuncReset() bool {
	return js.True == bindings.HasFuncReset()
}

// FuncReset returns the function "WEBEXT.inputMethodPrivate.reset".
func FuncReset() (fn js.Func[func()]) {
	bindings.FuncReset(
		js.Pointer(&fn),
	)
	return
}

// Reset calls the function "WEBEXT.inputMethodPrivate.reset" directly.
func Reset() (ret js.Void) {
	bindings.CallReset(
		js.Pointer(&ret),
	)

	return
}

// TryReset calls the function "WEBEXT.inputMethodPrivate.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReset(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCompositionRange returns true if the function "WEBEXT.inputMethodPrivate.setCompositionRange" exists.
func HasFuncSetCompositionRange() bool {
	return js.True == bindings.HasFuncSetCompositionRange()
}

// FuncSetCompositionRange returns the function "WEBEXT.inputMethodPrivate.setCompositionRange".
func FuncSetCompositionRange() (fn js.Func[func(parameters SetCompositionRangeArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetCompositionRange(
		js.Pointer(&fn),
	)
	return
}

// SetCompositionRange calls the function "WEBEXT.inputMethodPrivate.setCompositionRange" directly.
func SetCompositionRange(parameters SetCompositionRangeArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetCompositionRange(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetCompositionRange calls the function "WEBEXT.inputMethodPrivate.setCompositionRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCompositionRange(parameters SetCompositionRangeArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCompositionRange(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetCurrentInputMethod returns true if the function "WEBEXT.inputMethodPrivate.setCurrentInputMethod" exists.
func HasFuncSetCurrentInputMethod() bool {
	return js.True == bindings.HasFuncSetCurrentInputMethod()
}

// FuncSetCurrentInputMethod returns the function "WEBEXT.inputMethodPrivate.setCurrentInputMethod".
func FuncSetCurrentInputMethod() (fn js.Func[func(inputMethodId js.String) js.Promise[js.Void]]) {
	bindings.FuncSetCurrentInputMethod(
		js.Pointer(&fn),
	)
	return
}

// SetCurrentInputMethod calls the function "WEBEXT.inputMethodPrivate.setCurrentInputMethod" directly.
func SetCurrentInputMethod(inputMethodId js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetCurrentInputMethod(
		js.Pointer(&ret),
		inputMethodId.Ref(),
	)

	return
}

// TrySetCurrentInputMethod calls the function "WEBEXT.inputMethodPrivate.setCurrentInputMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCurrentInputMethod(inputMethodId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCurrentInputMethod(
		js.Pointer(&ret), js.Pointer(&exception),
		inputMethodId.Ref(),
	)

	return
}

// HasFuncSetSettings returns true if the function "WEBEXT.inputMethodPrivate.setSettings" exists.
func HasFuncSetSettings() bool {
	return js.True == bindings.HasFuncSetSettings()
}

// FuncSetSettings returns the function "WEBEXT.inputMethodPrivate.setSettings".
func FuncSetSettings() (fn js.Func[func(engineID js.String, settings InputMethodSettings) js.Promise[js.Void]]) {
	bindings.FuncSetSettings(
		js.Pointer(&fn),
	)
	return
}

// SetSettings calls the function "WEBEXT.inputMethodPrivate.setSettings" directly.
func SetSettings(engineID js.String, settings InputMethodSettings) (ret js.Promise[js.Void]) {
	bindings.CallSetSettings(
		js.Pointer(&ret),
		engineID.Ref(),
		js.Pointer(&settings),
	)

	return
}

// TrySetSettings calls the function "WEBEXT.inputMethodPrivate.setSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSettings(engineID js.String, settings InputMethodSettings) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		engineID.Ref(),
		js.Pointer(&settings),
	)

	return
}

// HasFuncSetXkbLayout returns true if the function "WEBEXT.inputMethodPrivate.setXkbLayout" exists.
func HasFuncSetXkbLayout() bool {
	return js.True == bindings.HasFuncSetXkbLayout()
}

// FuncSetXkbLayout returns the function "WEBEXT.inputMethodPrivate.setXkbLayout".
func FuncSetXkbLayout() (fn js.Func[func(xkb_name js.String) js.Promise[js.Void]]) {
	bindings.FuncSetXkbLayout(
		js.Pointer(&fn),
	)
	return
}

// SetXkbLayout calls the function "WEBEXT.inputMethodPrivate.setXkbLayout" directly.
func SetXkbLayout(xkb_name js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetXkbLayout(
		js.Pointer(&ret),
		xkb_name.Ref(),
	)

	return
}

// TrySetXkbLayout calls the function "WEBEXT.inputMethodPrivate.setXkbLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetXkbLayout(xkb_name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetXkbLayout(
		js.Pointer(&ret), js.Pointer(&exception),
		xkb_name.Ref(),
	)

	return
}

// HasFuncShowInputView returns true if the function "WEBEXT.inputMethodPrivate.showInputView" exists.
func HasFuncShowInputView() bool {
	return js.True == bindings.HasFuncShowInputView()
}

// FuncShowInputView returns the function "WEBEXT.inputMethodPrivate.showInputView".
func FuncShowInputView() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncShowInputView(
		js.Pointer(&fn),
	)
	return
}

// ShowInputView calls the function "WEBEXT.inputMethodPrivate.showInputView" directly.
func ShowInputView() (ret js.Promise[js.Void]) {
	bindings.CallShowInputView(
		js.Pointer(&ret),
	)

	return
}

// TryShowInputView calls the function "WEBEXT.inputMethodPrivate.showInputView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowInputView() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowInputView(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSwitchToLastUsedInputMethod returns true if the function "WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod" exists.
func HasFuncSwitchToLastUsedInputMethod() bool {
	return js.True == bindings.HasFuncSwitchToLastUsedInputMethod()
}

// FuncSwitchToLastUsedInputMethod returns the function "WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod".
func FuncSwitchToLastUsedInputMethod() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncSwitchToLastUsedInputMethod(
		js.Pointer(&fn),
	)
	return
}

// SwitchToLastUsedInputMethod calls the function "WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod" directly.
func SwitchToLastUsedInputMethod() (ret js.Promise[js.Void]) {
	bindings.CallSwitchToLastUsedInputMethod(
		js.Pointer(&ret),
	)

	return
}

// TrySwitchToLastUsedInputMethod calls the function "WEBEXT.inputMethodPrivate.switchToLastUsedInputMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySwitchToLastUsedInputMethod() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySwitchToLastUsedInputMethod(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

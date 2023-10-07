// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package accessibilityprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/accessibilityprivate/bindings"
)

type AcceleratorAction uint32

const (
	_ AcceleratorAction = iota

	AcceleratorAction_FOCUS_PREVIOUS_PANE
	AcceleratorAction_FOCUS_NEXT_PANE
)

func (AcceleratorAction) FromRef(str js.Ref) AcceleratorAction {
	return AcceleratorAction(bindings.ConstOfAcceleratorAction(str))
}

func (x AcceleratorAction) String() (string, bool) {
	switch x {
	case AcceleratorAction_FOCUS_PREVIOUS_PANE:
		return "focusPreviousPane", true
	case AcceleratorAction_FOCUS_NEXT_PANE:
		return "focusNextPane", true
	default:
		return "", false
	}
}

type AccessibilityFeature uint32

const (
	_ AccessibilityFeature = iota

	AccessibilityFeature_GOOGLE_TTS_LANGUAGE_PACKS
	AccessibilityFeature_DICTATION_CONTEXT_CHECKING
	AccessibilityFeature_CHROMEVOX_SETTINGS_MIGRATION
	AccessibilityFeature_GAME_FACE_INTEGRATION
	AccessibilityFeature_GOOGLE_TTS_HIGH_QUALITY_VOICES
)

func (AccessibilityFeature) FromRef(str js.Ref) AccessibilityFeature {
	return AccessibilityFeature(bindings.ConstOfAccessibilityFeature(str))
}

func (x AccessibilityFeature) String() (string, bool) {
	switch x {
	case AccessibilityFeature_GOOGLE_TTS_LANGUAGE_PACKS:
		return "googleTtsLanguagePacks", true
	case AccessibilityFeature_DICTATION_CONTEXT_CHECKING:
		return "dictationContextChecking", true
	case AccessibilityFeature_CHROMEVOX_SETTINGS_MIGRATION:
		return "chromevoxSettingsMigration", true
	case AccessibilityFeature_GAME_FACE_INTEGRATION:
		return "gameFaceIntegration", true
	case AccessibilityFeature_GOOGLE_TTS_HIGH_QUALITY_VOICES:
		return "googleTtsHighQualityVoices", true
	default:
		return "", false
	}
}

type AlertInfo struct {
	// Message is "AlertInfo.message"
	//
	// Required
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AlertInfo with all fields set.
func (p AlertInfo) FromRef(ref js.Ref) AlertInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AlertInfo in the application heap.
func (p AlertInfo) New() js.Ref {
	return bindings.AlertInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AlertInfo) UpdateFrom(ref js.Ref) {
	bindings.AlertInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AlertInfo) Update(ref js.Ref) {
	bindings.AlertInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AlertInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type AssistiveTechnologyType uint32

const (
	_ AssistiveTechnologyType = iota

	AssistiveTechnologyType_CHROME_VOX
	AssistiveTechnologyType_SELECT_TO_SPEAK
	AssistiveTechnologyType_SWITCH_ACCESS
	AssistiveTechnologyType_AUTO_CLICK
	AssistiveTechnologyType_MAGNIFIER
	AssistiveTechnologyType_DICTATION
)

func (AssistiveTechnologyType) FromRef(str js.Ref) AssistiveTechnologyType {
	return AssistiveTechnologyType(bindings.ConstOfAssistiveTechnologyType(str))
}

func (x AssistiveTechnologyType) String() (string, bool) {
	switch x {
	case AssistiveTechnologyType_CHROME_VOX:
		return "chromeVox", true
	case AssistiveTechnologyType_SELECT_TO_SPEAK:
		return "selectToSpeak", true
	case AssistiveTechnologyType_SWITCH_ACCESS:
		return "switchAccess", true
	case AssistiveTechnologyType_AUTO_CLICK:
		return "autoClick", true
	case AssistiveTechnologyType_MAGNIFIER:
		return "magnifier", true
	case AssistiveTechnologyType_DICTATION:
		return "dictation", true
	default:
		return "", false
	}
}

type DictationBubbleHintType uint32

const (
	_ DictationBubbleHintType = iota

	DictationBubbleHintType_TRY_SAYING
	DictationBubbleHintType_TYPE
	DictationBubbleHintType_DELETE
	DictationBubbleHintType_SELECT_ALL
	DictationBubbleHintType_UNDO
	DictationBubbleHintType_HELP
	DictationBubbleHintType_UNSELECT
	DictationBubbleHintType_COPY
)

func (DictationBubbleHintType) FromRef(str js.Ref) DictationBubbleHintType {
	return DictationBubbleHintType(bindings.ConstOfDictationBubbleHintType(str))
}

func (x DictationBubbleHintType) String() (string, bool) {
	switch x {
	case DictationBubbleHintType_TRY_SAYING:
		return "trySaying", true
	case DictationBubbleHintType_TYPE:
		return "type", true
	case DictationBubbleHintType_DELETE:
		return "delete", true
	case DictationBubbleHintType_SELECT_ALL:
		return "selectAll", true
	case DictationBubbleHintType_UNDO:
		return "undo", true
	case DictationBubbleHintType_HELP:
		return "help", true
	case DictationBubbleHintType_UNSELECT:
		return "unselect", true
	case DictationBubbleHintType_COPY:
		return "copy", true
	default:
		return "", false
	}
}

type DictationBubbleIconType uint32

const (
	_ DictationBubbleIconType = iota

	DictationBubbleIconType_HIDDEN
	DictationBubbleIconType_STANDBY
	DictationBubbleIconType_MACRO_SUCCESS
	DictationBubbleIconType_MACRO_FAIL
)

func (DictationBubbleIconType) FromRef(str js.Ref) DictationBubbleIconType {
	return DictationBubbleIconType(bindings.ConstOfDictationBubbleIconType(str))
}

func (x DictationBubbleIconType) String() (string, bool) {
	switch x {
	case DictationBubbleIconType_HIDDEN:
		return "hidden", true
	case DictationBubbleIconType_STANDBY:
		return "standby", true
	case DictationBubbleIconType_MACRO_SUCCESS:
		return "macroSuccess", true
	case DictationBubbleIconType_MACRO_FAIL:
		return "macroFail", true
	default:
		return "", false
	}
}

type DictationBubbleProperties struct {
	// Hints is "DictationBubbleProperties.hints"
	//
	// Optional
	Hints js.Array[DictationBubbleHintType]
	// Icon is "DictationBubbleProperties.icon"
	//
	// Required
	Icon DictationBubbleIconType
	// Text is "DictationBubbleProperties.text"
	//
	// Optional
	Text js.String
	// Visible is "DictationBubbleProperties.visible"
	//
	// Required
	Visible bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DictationBubbleProperties with all fields set.
func (p DictationBubbleProperties) FromRef(ref js.Ref) DictationBubbleProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DictationBubbleProperties in the application heap.
func (p DictationBubbleProperties) New() js.Ref {
	return bindings.DictationBubblePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DictationBubbleProperties) UpdateFrom(ref js.Ref) {
	bindings.DictationBubblePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DictationBubbleProperties) Update(ref js.Ref) {
	bindings.DictationBubblePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DictationBubbleProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Hints.Ref(),
		p.Text.Ref(),
	)
	p.Hints = p.Hints.FromRef(js.Undefined)
	p.Text = p.Text.FromRef(js.Undefined)
}

type DlcType uint32

const (
	_ DlcType = iota

	DlcType_TTS_BN_BD
	DlcType_TTS_CS_CZ
	DlcType_TTS_DA_DK
	DlcType_TTS_DE_DE
	DlcType_TTS_EL_GR
	DlcType_TTS_EN_AU
	DlcType_TTS_EN_GB
	DlcType_TTS_EN_US
	DlcType_TTS_ES_ES
	DlcType_TTS_ES_US
	DlcType_TTS_FI_FI
	DlcType_TTS_FIL_PH
	DlcType_TTS_FR_FR
	DlcType_TTS_HI_IN
	DlcType_TTS_HU_HU
	DlcType_TTS_ID_ID
	DlcType_TTS_IT_IT
	DlcType_TTS_JA_JP
	DlcType_TTS_KM_KH
	DlcType_TTS_KO_KR
	DlcType_TTS_NB_NO
	DlcType_TTS_NE_NP
	DlcType_TTS_NL_NL
	DlcType_TTS_PL_PL
	DlcType_TTS_PT_BR
	DlcType_TTS_SI_LK
	DlcType_TTS_SK_SK
	DlcType_TTS_SV_SE
	DlcType_TTS_TH_TH
	DlcType_TTS_TR_TR
	DlcType_TTS_UK_UA
	DlcType_TTS_VI_VN
	DlcType_TTS_YUE_HK
)

func (DlcType) FromRef(str js.Ref) DlcType {
	return DlcType(bindings.ConstOfDlcType(str))
}

func (x DlcType) String() (string, bool) {
	switch x {
	case DlcType_TTS_BN_BD:
		return "ttsBnBd", true
	case DlcType_TTS_CS_CZ:
		return "ttsCsCz", true
	case DlcType_TTS_DA_DK:
		return "ttsDaDk", true
	case DlcType_TTS_DE_DE:
		return "ttsDeDe", true
	case DlcType_TTS_EL_GR:
		return "ttsElGr", true
	case DlcType_TTS_EN_AU:
		return "ttsEnAu", true
	case DlcType_TTS_EN_GB:
		return "ttsEnGb", true
	case DlcType_TTS_EN_US:
		return "ttsEnUs", true
	case DlcType_TTS_ES_ES:
		return "ttsEsEs", true
	case DlcType_TTS_ES_US:
		return "ttsEsUs", true
	case DlcType_TTS_FI_FI:
		return "ttsFiFi", true
	case DlcType_TTS_FIL_PH:
		return "ttsFilPh", true
	case DlcType_TTS_FR_FR:
		return "ttsFrFr", true
	case DlcType_TTS_HI_IN:
		return "ttsHiIn", true
	case DlcType_TTS_HU_HU:
		return "ttsHuHu", true
	case DlcType_TTS_ID_ID:
		return "ttsIdId", true
	case DlcType_TTS_IT_IT:
		return "ttsItIt", true
	case DlcType_TTS_JA_JP:
		return "ttsJaJp", true
	case DlcType_TTS_KM_KH:
		return "ttsKmKh", true
	case DlcType_TTS_KO_KR:
		return "ttsKoKr", true
	case DlcType_TTS_NB_NO:
		return "ttsNbNo", true
	case DlcType_TTS_NE_NP:
		return "ttsNeNp", true
	case DlcType_TTS_NL_NL:
		return "ttsNlNl", true
	case DlcType_TTS_PL_PL:
		return "ttsPlPl", true
	case DlcType_TTS_PT_BR:
		return "ttsPtBr", true
	case DlcType_TTS_SI_LK:
		return "ttsSiLk", true
	case DlcType_TTS_SK_SK:
		return "ttsSkSk", true
	case DlcType_TTS_SV_SE:
		return "ttsSvSe", true
	case DlcType_TTS_TH_TH:
		return "ttsThTh", true
	case DlcType_TTS_TR_TR:
		return "ttsTrTr", true
	case DlcType_TTS_UK_UA:
		return "ttsUkUa", true
	case DlcType_TTS_VI_VN:
		return "ttsViVn", true
	case DlcType_TTS_YUE_HK:
		return "ttsYueHk", true
	default:
		return "", false
	}
}

type ScreenRect struct {
	// Height is "ScreenRect.height"
	//
	// Required
	Height int64
	// Left is "ScreenRect.left"
	//
	// Required
	Left int64
	// Top is "ScreenRect.top"
	//
	// Required
	Top int64
	// Width is "ScreenRect.width"
	//
	// Required
	Width int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScreenRect with all fields set.
func (p ScreenRect) FromRef(ref js.Ref) ScreenRect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScreenRect in the application heap.
func (p ScreenRect) New() js.Ref {
	return bindings.ScreenRectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScreenRect) UpdateFrom(ref js.Ref) {
	bindings.ScreenRectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScreenRect) Update(ref js.Ref) {
	bindings.ScreenRectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScreenRect) FreeMembers(recursive bool) {
}

type FocusRingStackingOrder uint32

const (
	_ FocusRingStackingOrder = iota

	FocusRingStackingOrder_ABOVE_ACCESSIBILITY_BUBBLES
	FocusRingStackingOrder_BELOW_ACCESSIBILITY_BUBBLES
)

func (FocusRingStackingOrder) FromRef(str js.Ref) FocusRingStackingOrder {
	return FocusRingStackingOrder(bindings.ConstOfFocusRingStackingOrder(str))
}

func (x FocusRingStackingOrder) String() (string, bool) {
	switch x {
	case FocusRingStackingOrder_ABOVE_ACCESSIBILITY_BUBBLES:
		return "aboveAccessibilityBubbles", true
	case FocusRingStackingOrder_BELOW_ACCESSIBILITY_BUBBLES:
		return "belowAccessibilityBubbles", true
	default:
		return "", false
	}
}

type FocusType uint32

const (
	_ FocusType = iota

	FocusType_GLOW
	FocusType_SOLID
	FocusType_DASHED
)

func (FocusType) FromRef(str js.Ref) FocusType {
	return FocusType(bindings.ConstOfFocusType(str))
}

func (x FocusType) String() (string, bool) {
	switch x {
	case FocusType_GLOW:
		return "glow", true
	case FocusType_SOLID:
		return "solid", true
	case FocusType_DASHED:
		return "dashed", true
	default:
		return "", false
	}
}

type FocusRingInfo struct {
	// BackgroundColor is "FocusRingInfo.backgroundColor"
	//
	// Optional
	BackgroundColor js.String
	// Color is "FocusRingInfo.color"
	//
	// Required
	Color js.String
	// Id is "FocusRingInfo.id"
	//
	// Optional
	Id js.String
	// Rects is "FocusRingInfo.rects"
	//
	// Required
	Rects js.Array[ScreenRect]
	// SecondaryColor is "FocusRingInfo.secondaryColor"
	//
	// Optional
	SecondaryColor js.String
	// StackingOrder is "FocusRingInfo.stackingOrder"
	//
	// Optional
	StackingOrder FocusRingStackingOrder
	// Type is "FocusRingInfo.type"
	//
	// Required
	Type FocusType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FocusRingInfo with all fields set.
func (p FocusRingInfo) FromRef(ref js.Ref) FocusRingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FocusRingInfo in the application heap.
func (p FocusRingInfo) New() js.Ref {
	return bindings.FocusRingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FocusRingInfo) UpdateFrom(ref js.Ref) {
	bindings.FocusRingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FocusRingInfo) Update(ref js.Ref) {
	bindings.FocusRingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FocusRingInfo) FreeMembers(recursive bool) {
	js.Free(
		p.BackgroundColor.Ref(),
		p.Color.Ref(),
		p.Id.Ref(),
		p.Rects.Ref(),
		p.SecondaryColor.Ref(),
	)
	p.BackgroundColor = p.BackgroundColor.FromRef(js.Undefined)
	p.Color = p.Color.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Rects = p.Rects.FromRef(js.Undefined)
	p.SecondaryColor = p.SecondaryColor.FromRef(js.Undefined)
}

type Gesture uint32

const (
	_ Gesture = iota

	Gesture_CLICK
	Gesture_SWIPE_LEFT1
	Gesture_SWIPE_UP1
	Gesture_SWIPE_RIGHT1
	Gesture_SWIPE_DOWN1
	Gesture_SWIPE_LEFT2
	Gesture_SWIPE_UP2
	Gesture_SWIPE_RIGHT2
	Gesture_SWIPE_DOWN2
	Gesture_SWIPE_LEFT3
	Gesture_SWIPE_UP3
	Gesture_SWIPE_RIGHT3
	Gesture_SWIPE_DOWN3
	Gesture_SWIPE_LEFT4
	Gesture_SWIPE_UP4
	Gesture_SWIPE_RIGHT4
	Gesture_SWIPE_DOWN4
	Gesture_TAP2
	Gesture_TAP3
	Gesture_TAP4
	Gesture_TOUCH_EXPLORE
)

func (Gesture) FromRef(str js.Ref) Gesture {
	return Gesture(bindings.ConstOfGesture(str))
}

func (x Gesture) String() (string, bool) {
	switch x {
	case Gesture_CLICK:
		return "click", true
	case Gesture_SWIPE_LEFT1:
		return "swipeLeft1", true
	case Gesture_SWIPE_UP1:
		return "swipeUp1", true
	case Gesture_SWIPE_RIGHT1:
		return "swipeRight1", true
	case Gesture_SWIPE_DOWN1:
		return "swipeDown1", true
	case Gesture_SWIPE_LEFT2:
		return "swipeLeft2", true
	case Gesture_SWIPE_UP2:
		return "swipeUp2", true
	case Gesture_SWIPE_RIGHT2:
		return "swipeRight2", true
	case Gesture_SWIPE_DOWN2:
		return "swipeDown2", true
	case Gesture_SWIPE_LEFT3:
		return "swipeLeft3", true
	case Gesture_SWIPE_UP3:
		return "swipeUp3", true
	case Gesture_SWIPE_RIGHT3:
		return "swipeRight3", true
	case Gesture_SWIPE_DOWN3:
		return "swipeDown3", true
	case Gesture_SWIPE_LEFT4:
		return "swipeLeft4", true
	case Gesture_SWIPE_UP4:
		return "swipeUp4", true
	case Gesture_SWIPE_RIGHT4:
		return "swipeRight4", true
	case Gesture_SWIPE_DOWN4:
		return "swipeDown4", true
	case Gesture_TAP2:
		return "tap2", true
	case Gesture_TAP3:
		return "tap3", true
	case Gesture_TAP4:
		return "tap4", true
	case Gesture_TOUCH_EXPLORE:
		return "touchExplore", true
	default:
		return "", false
	}
}

// IS_DEFAULT_EVENT_SOURCE_TOUCH returns the value of property "WEBEXT.accessibilityPrivate.IS_DEFAULT_EVENT_SOURCE_TOUCH".
//
// The returned bool will be false if there is no such property.
func IS_DEFAULT_EVENT_SOURCE_TOUCH() (ret int64, ok bool) {
	ok = js.True == bindings.GetIS_DEFAULT_EVENT_SOURCE_TOUCH(
		js.Pointer(&ret),
	)

	return
}

// SetIS_DEFAULT_EVENT_SOURCE_TOUCH sets the value of property "WEBEXT.accessibilityPrivate.IS_DEFAULT_EVENT_SOURCE_TOUCH" to val.
//
// It returns false if the property cannot be set.
func SetIS_DEFAULT_EVENT_SOURCE_TOUCH(val int64) bool {
	return js.True == bindings.SetIS_DEFAULT_EVENT_SOURCE_TOUCH(
		float64(val))
}

type MagnifierCommand uint32

const (
	_ MagnifierCommand = iota

	MagnifierCommand_MOVE_STOP
	MagnifierCommand_MOVE_UP
	MagnifierCommand_MOVE_DOWN
	MagnifierCommand_MOVE_LEFT
	MagnifierCommand_MOVE_RIGHT
)

func (MagnifierCommand) FromRef(str js.Ref) MagnifierCommand {
	return MagnifierCommand(bindings.ConstOfMagnifierCommand(str))
}

func (x MagnifierCommand) String() (string, bool) {
	switch x {
	case MagnifierCommand_MOVE_STOP:
		return "moveStop", true
	case MagnifierCommand_MOVE_UP:
		return "moveUp", true
	case MagnifierCommand_MOVE_DOWN:
		return "moveDown", true
	case MagnifierCommand_MOVE_LEFT:
		return "moveLeft", true
	case MagnifierCommand_MOVE_RIGHT:
		return "moveRight", true
	default:
		return "", false
	}
}

type PointScanPoint struct {
	// X is "PointScanPoint.x"
	//
	// Required
	X float64
	// Y is "PointScanPoint.y"
	//
	// Required
	Y float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PointScanPoint with all fields set.
func (p PointScanPoint) FromRef(ref js.Ref) PointScanPoint {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PointScanPoint in the application heap.
func (p PointScanPoint) New() js.Ref {
	return bindings.PointScanPointJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PointScanPoint) UpdateFrom(ref js.Ref) {
	bindings.PointScanPointJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PointScanPoint) Update(ref js.Ref) {
	bindings.PointScanPointJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PointScanPoint) FreeMembers(recursive bool) {
}

type PointScanState uint32

const (
	_ PointScanState = iota

	PointScanState_START
	PointScanState_STOP
)

func (PointScanState) FromRef(str js.Ref) PointScanState {
	return PointScanState(bindings.ConstOfPointScanState(str))
}

func (x PointScanState) String() (string, bool) {
	switch x {
	case PointScanState_START:
		return "start", true
	case PointScanState_STOP:
		return "stop", true
	default:
		return "", false
	}
}

type PumpkinData struct {
	// DeDeActionConfigBinarypb is "PumpkinData.de_de_action_config_binarypb"
	//
	// Required
	DeDeActionConfigBinarypb js.TypedArray[uint8]
	// DeDePumpkinConfigBinarypb is "PumpkinData.de_de_pumpkin_config_binarypb"
	//
	// Required
	DeDePumpkinConfigBinarypb js.TypedArray[uint8]
	// EnUsActionConfigBinarypb is "PumpkinData.en_us_action_config_binarypb"
	//
	// Required
	EnUsActionConfigBinarypb js.TypedArray[uint8]
	// EnUsPumpkinConfigBinarypb is "PumpkinData.en_us_pumpkin_config_binarypb"
	//
	// Required
	EnUsPumpkinConfigBinarypb js.TypedArray[uint8]
	// EsEsActionConfigBinarypb is "PumpkinData.es_es_action_config_binarypb"
	//
	// Required
	EsEsActionConfigBinarypb js.TypedArray[uint8]
	// EsEsPumpkinConfigBinarypb is "PumpkinData.es_es_pumpkin_config_binarypb"
	//
	// Required
	EsEsPumpkinConfigBinarypb js.TypedArray[uint8]
	// FrFrActionConfigBinarypb is "PumpkinData.fr_fr_action_config_binarypb"
	//
	// Required
	FrFrActionConfigBinarypb js.TypedArray[uint8]
	// FrFrPumpkinConfigBinarypb is "PumpkinData.fr_fr_pumpkin_config_binarypb"
	//
	// Required
	FrFrPumpkinConfigBinarypb js.TypedArray[uint8]
	// ItItActionConfigBinarypb is "PumpkinData.it_it_action_config_binarypb"
	//
	// Required
	ItItActionConfigBinarypb js.TypedArray[uint8]
	// ItItPumpkinConfigBinarypb is "PumpkinData.it_it_pumpkin_config_binarypb"
	//
	// Required
	ItItPumpkinConfigBinarypb js.TypedArray[uint8]
	// JsPumpkinTaggerBinJs is "PumpkinData.js_pumpkin_tagger_bin_js"
	//
	// Required
	JsPumpkinTaggerBinJs js.TypedArray[uint8]
	// TaggerWasmMainJs is "PumpkinData.tagger_wasm_main_js"
	//
	// Required
	TaggerWasmMainJs js.TypedArray[uint8]
	// TaggerWasmMainWasm is "PumpkinData.tagger_wasm_main_wasm"
	//
	// Required
	TaggerWasmMainWasm js.TypedArray[uint8]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PumpkinData with all fields set.
func (p PumpkinData) FromRef(ref js.Ref) PumpkinData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PumpkinData in the application heap.
func (p PumpkinData) New() js.Ref {
	return bindings.PumpkinDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PumpkinData) UpdateFrom(ref js.Ref) {
	bindings.PumpkinDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PumpkinData) Update(ref js.Ref) {
	bindings.PumpkinDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PumpkinData) FreeMembers(recursive bool) {
	js.Free(
		p.DeDeActionConfigBinarypb.Ref(),
		p.DeDePumpkinConfigBinarypb.Ref(),
		p.EnUsActionConfigBinarypb.Ref(),
		p.EnUsPumpkinConfigBinarypb.Ref(),
		p.EsEsActionConfigBinarypb.Ref(),
		p.EsEsPumpkinConfigBinarypb.Ref(),
		p.FrFrActionConfigBinarypb.Ref(),
		p.FrFrPumpkinConfigBinarypb.Ref(),
		p.ItItActionConfigBinarypb.Ref(),
		p.ItItPumpkinConfigBinarypb.Ref(),
		p.JsPumpkinTaggerBinJs.Ref(),
		p.TaggerWasmMainJs.Ref(),
		p.TaggerWasmMainWasm.Ref(),
	)
	p.DeDeActionConfigBinarypb = p.DeDeActionConfigBinarypb.FromRef(js.Undefined)
	p.DeDePumpkinConfigBinarypb = p.DeDePumpkinConfigBinarypb.FromRef(js.Undefined)
	p.EnUsActionConfigBinarypb = p.EnUsActionConfigBinarypb.FromRef(js.Undefined)
	p.EnUsPumpkinConfigBinarypb = p.EnUsPumpkinConfigBinarypb.FromRef(js.Undefined)
	p.EsEsActionConfigBinarypb = p.EsEsActionConfigBinarypb.FromRef(js.Undefined)
	p.EsEsPumpkinConfigBinarypb = p.EsEsPumpkinConfigBinarypb.FromRef(js.Undefined)
	p.FrFrActionConfigBinarypb = p.FrFrActionConfigBinarypb.FromRef(js.Undefined)
	p.FrFrPumpkinConfigBinarypb = p.FrFrPumpkinConfigBinarypb.FromRef(js.Undefined)
	p.ItItActionConfigBinarypb = p.ItItActionConfigBinarypb.FromRef(js.Undefined)
	p.ItItPumpkinConfigBinarypb = p.ItItPumpkinConfigBinarypb.FromRef(js.Undefined)
	p.JsPumpkinTaggerBinJs = p.JsPumpkinTaggerBinJs.FromRef(js.Undefined)
	p.TaggerWasmMainJs = p.TaggerWasmMainJs.FromRef(js.Undefined)
	p.TaggerWasmMainWasm = p.TaggerWasmMainWasm.FromRef(js.Undefined)
}

type ScreenPoint struct {
	// X is "ScreenPoint.x"
	//
	// Required
	X int64
	// Y is "ScreenPoint.y"
	//
	// Required
	Y int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScreenPoint with all fields set.
func (p ScreenPoint) FromRef(ref js.Ref) ScreenPoint {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScreenPoint in the application heap.
func (p ScreenPoint) New() js.Ref {
	return bindings.ScreenPointJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScreenPoint) UpdateFrom(ref js.Ref) {
	bindings.ScreenPointJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScreenPoint) Update(ref js.Ref) {
	bindings.ScreenPointJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScreenPoint) FreeMembers(recursive bool) {
}

type SelectToSpeakPanelAction uint32

const (
	_ SelectToSpeakPanelAction = iota

	SelectToSpeakPanelAction_PREVIOUS_PARAGRAPH
	SelectToSpeakPanelAction_PREVIOUS_SENTENCE
	SelectToSpeakPanelAction_PAUSE
	SelectToSpeakPanelAction_RESUME
	SelectToSpeakPanelAction_NEXT_SENTENCE
	SelectToSpeakPanelAction_NEXT_PARAGRAPH
	SelectToSpeakPanelAction_EXIT
	SelectToSpeakPanelAction_CHANGE_SPEED
)

func (SelectToSpeakPanelAction) FromRef(str js.Ref) SelectToSpeakPanelAction {
	return SelectToSpeakPanelAction(bindings.ConstOfSelectToSpeakPanelAction(str))
}

func (x SelectToSpeakPanelAction) String() (string, bool) {
	switch x {
	case SelectToSpeakPanelAction_PREVIOUS_PARAGRAPH:
		return "previousParagraph", true
	case SelectToSpeakPanelAction_PREVIOUS_SENTENCE:
		return "previousSentence", true
	case SelectToSpeakPanelAction_PAUSE:
		return "pause", true
	case SelectToSpeakPanelAction_RESUME:
		return "resume", true
	case SelectToSpeakPanelAction_NEXT_SENTENCE:
		return "nextSentence", true
	case SelectToSpeakPanelAction_NEXT_PARAGRAPH:
		return "nextParagraph", true
	case SelectToSpeakPanelAction_EXIT:
		return "exit", true
	case SelectToSpeakPanelAction_CHANGE_SPEED:
		return "changeSpeed", true
	default:
		return "", false
	}
}

type SelectToSpeakState uint32

const (
	_ SelectToSpeakState = iota

	SelectToSpeakState_SELECTING
	SelectToSpeakState_SPEAKING
	SelectToSpeakState_INACTIVE
)

func (SelectToSpeakState) FromRef(str js.Ref) SelectToSpeakState {
	return SelectToSpeakState(bindings.ConstOfSelectToSpeakState(str))
}

func (x SelectToSpeakState) String() (string, bool) {
	switch x {
	case SelectToSpeakState_SELECTING:
		return "selecting", true
	case SelectToSpeakState_SPEAKING:
		return "speaking", true
	case SelectToSpeakState_INACTIVE:
		return "inactive", true
	default:
		return "", false
	}
}

type SetNativeChromeVoxResponse uint32

const (
	_ SetNativeChromeVoxResponse = iota

	SetNativeChromeVoxResponse_SUCCESS
	SetNativeChromeVoxResponse_TALKBACK_NOT_INSTALLED
	SetNativeChromeVoxResponse_WINDOW_NOT_FOUND
	SetNativeChromeVoxResponse_FAILURE
	SetNativeChromeVoxResponse_NEED_DEPRECATION_CONFIRMATION
)

func (SetNativeChromeVoxResponse) FromRef(str js.Ref) SetNativeChromeVoxResponse {
	return SetNativeChromeVoxResponse(bindings.ConstOfSetNativeChromeVoxResponse(str))
}

func (x SetNativeChromeVoxResponse) String() (string, bool) {
	switch x {
	case SetNativeChromeVoxResponse_SUCCESS:
		return "success", true
	case SetNativeChromeVoxResponse_TALKBACK_NOT_INSTALLED:
		return "talkbackNotInstalled", true
	case SetNativeChromeVoxResponse_WINDOW_NOT_FOUND:
		return "windowNotFound", true
	case SetNativeChromeVoxResponse_FAILURE:
		return "failure", true
	case SetNativeChromeVoxResponse_NEED_DEPRECATION_CONFIRMATION:
		return "needDeprecationConfirmation", true
	default:
		return "", false
	}
}

type SwitchAccessBubble uint32

const (
	_ SwitchAccessBubble = iota

	SwitchAccessBubble_BACK_BUTTON
	SwitchAccessBubble_MENU
)

func (SwitchAccessBubble) FromRef(str js.Ref) SwitchAccessBubble {
	return SwitchAccessBubble(bindings.ConstOfSwitchAccessBubble(str))
}

func (x SwitchAccessBubble) String() (string, bool) {
	switch x {
	case SwitchAccessBubble_BACK_BUTTON:
		return "backButton", true
	case SwitchAccessBubble_MENU:
		return "menu", true
	default:
		return "", false
	}
}

type SwitchAccessCommand uint32

const (
	_ SwitchAccessCommand = iota

	SwitchAccessCommand_SELECT
	SwitchAccessCommand_NEXT
	SwitchAccessCommand_PREVIOUS
)

func (SwitchAccessCommand) FromRef(str js.Ref) SwitchAccessCommand {
	return SwitchAccessCommand(bindings.ConstOfSwitchAccessCommand(str))
}

func (x SwitchAccessCommand) String() (string, bool) {
	switch x {
	case SwitchAccessCommand_SELECT:
		return "select", true
	case SwitchAccessCommand_NEXT:
		return "next", true
	case SwitchAccessCommand_PREVIOUS:
		return "previous", true
	default:
		return "", false
	}
}

type SwitchAccessMenuAction uint32

const (
	_ SwitchAccessMenuAction = iota

	SwitchAccessMenuAction_COPY
	SwitchAccessMenuAction_CUT
	SwitchAccessMenuAction_DECREMENT
	SwitchAccessMenuAction_DICTATION
	SwitchAccessMenuAction_END_TEXT_SELECTION
	SwitchAccessMenuAction_INCREMENT
	SwitchAccessMenuAction_ITEM_SCAN
	SwitchAccessMenuAction_JUMP_TO_BEGINNING_OF_TEXT
	SwitchAccessMenuAction_JUMP_TO_END_OF_TEXT
	SwitchAccessMenuAction_KEYBOARD
	SwitchAccessMenuAction_LEFT_CLICK
	SwitchAccessMenuAction_MOVE_BACKWARD_ONE_CHAR_OF_TEXT
	SwitchAccessMenuAction_MOVE_BACKWARD_ONE_WORD_OF_TEXT
	SwitchAccessMenuAction_MOVE_CURSOR
	SwitchAccessMenuAction_MOVE_DOWN_ONE_LINE_OF_TEXT
	SwitchAccessMenuAction_MOVE_FORWARD_ONE_CHAR_OF_TEXT
	SwitchAccessMenuAction_MOVE_FORWARD_ONE_WORD_OF_TEXT
	SwitchAccessMenuAction_MOVE_UP_ONE_LINE_OF_TEXT
	SwitchAccessMenuAction_PASTE
	SwitchAccessMenuAction_POINT_SCAN
	SwitchAccessMenuAction_RIGHT_CLICK
	SwitchAccessMenuAction_SCROLL_DOWN
	SwitchAccessMenuAction_SCROLL_LEFT
	SwitchAccessMenuAction_SCROLL_RIGHT
	SwitchAccessMenuAction_SCROLL_UP
	SwitchAccessMenuAction_SELECT
	SwitchAccessMenuAction_SETTINGS
	SwitchAccessMenuAction_START_TEXT_SELECTION
)

func (SwitchAccessMenuAction) FromRef(str js.Ref) SwitchAccessMenuAction {
	return SwitchAccessMenuAction(bindings.ConstOfSwitchAccessMenuAction(str))
}

func (x SwitchAccessMenuAction) String() (string, bool) {
	switch x {
	case SwitchAccessMenuAction_COPY:
		return "copy", true
	case SwitchAccessMenuAction_CUT:
		return "cut", true
	case SwitchAccessMenuAction_DECREMENT:
		return "decrement", true
	case SwitchAccessMenuAction_DICTATION:
		return "dictation", true
	case SwitchAccessMenuAction_END_TEXT_SELECTION:
		return "endTextSelection", true
	case SwitchAccessMenuAction_INCREMENT:
		return "increment", true
	case SwitchAccessMenuAction_ITEM_SCAN:
		return "itemScan", true
	case SwitchAccessMenuAction_JUMP_TO_BEGINNING_OF_TEXT:
		return "jumpToBeginningOfText", true
	case SwitchAccessMenuAction_JUMP_TO_END_OF_TEXT:
		return "jumpToEndOfText", true
	case SwitchAccessMenuAction_KEYBOARD:
		return "keyboard", true
	case SwitchAccessMenuAction_LEFT_CLICK:
		return "leftClick", true
	case SwitchAccessMenuAction_MOVE_BACKWARD_ONE_CHAR_OF_TEXT:
		return "moveBackwardOneCharOfText", true
	case SwitchAccessMenuAction_MOVE_BACKWARD_ONE_WORD_OF_TEXT:
		return "moveBackwardOneWordOfText", true
	case SwitchAccessMenuAction_MOVE_CURSOR:
		return "moveCursor", true
	case SwitchAccessMenuAction_MOVE_DOWN_ONE_LINE_OF_TEXT:
		return "moveDownOneLineOfText", true
	case SwitchAccessMenuAction_MOVE_FORWARD_ONE_CHAR_OF_TEXT:
		return "moveForwardOneCharOfText", true
	case SwitchAccessMenuAction_MOVE_FORWARD_ONE_WORD_OF_TEXT:
		return "moveForwardOneWordOfText", true
	case SwitchAccessMenuAction_MOVE_UP_ONE_LINE_OF_TEXT:
		return "moveUpOneLineOfText", true
	case SwitchAccessMenuAction_PASTE:
		return "paste", true
	case SwitchAccessMenuAction_POINT_SCAN:
		return "pointScan", true
	case SwitchAccessMenuAction_RIGHT_CLICK:
		return "rightClick", true
	case SwitchAccessMenuAction_SCROLL_DOWN:
		return "scrollDown", true
	case SwitchAccessMenuAction_SCROLL_LEFT:
		return "scrollLeft", true
	case SwitchAccessMenuAction_SCROLL_RIGHT:
		return "scrollRight", true
	case SwitchAccessMenuAction_SCROLL_UP:
		return "scrollUp", true
	case SwitchAccessMenuAction_SELECT:
		return "select", true
	case SwitchAccessMenuAction_SETTINGS:
		return "settings", true
	case SwitchAccessMenuAction_START_TEXT_SELECTION:
		return "startTextSelection", true
	default:
		return "", false
	}
}

type SyntheticKeyboardModifiers struct {
	// Alt is "SyntheticKeyboardModifiers.alt"
	//
	// Optional
	//
	// NOTE: FFI_USE_Alt MUST be set to true to make this field effective.
	Alt bool
	// Ctrl is "SyntheticKeyboardModifiers.ctrl"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ctrl MUST be set to true to make this field effective.
	Ctrl bool
	// Search is "SyntheticKeyboardModifiers.search"
	//
	// Optional
	//
	// NOTE: FFI_USE_Search MUST be set to true to make this field effective.
	Search bool
	// Shift is "SyntheticKeyboardModifiers.shift"
	//
	// Optional
	//
	// NOTE: FFI_USE_Shift MUST be set to true to make this field effective.
	Shift bool

	FFI_USE_Alt    bool // for Alt.
	FFI_USE_Ctrl   bool // for Ctrl.
	FFI_USE_Search bool // for Search.
	FFI_USE_Shift  bool // for Shift.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SyntheticKeyboardModifiers with all fields set.
func (p SyntheticKeyboardModifiers) FromRef(ref js.Ref) SyntheticKeyboardModifiers {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SyntheticKeyboardModifiers in the application heap.
func (p SyntheticKeyboardModifiers) New() js.Ref {
	return bindings.SyntheticKeyboardModifiersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SyntheticKeyboardModifiers) UpdateFrom(ref js.Ref) {
	bindings.SyntheticKeyboardModifiersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SyntheticKeyboardModifiers) Update(ref js.Ref) {
	bindings.SyntheticKeyboardModifiersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SyntheticKeyboardModifiers) FreeMembers(recursive bool) {
}

type SyntheticKeyboardEventType uint32

const (
	_ SyntheticKeyboardEventType = iota

	SyntheticKeyboardEventType_KEYUP
	SyntheticKeyboardEventType_KEYDOWN
)

func (SyntheticKeyboardEventType) FromRef(str js.Ref) SyntheticKeyboardEventType {
	return SyntheticKeyboardEventType(bindings.ConstOfSyntheticKeyboardEventType(str))
}

func (x SyntheticKeyboardEventType) String() (string, bool) {
	switch x {
	case SyntheticKeyboardEventType_KEYUP:
		return "keyup", true
	case SyntheticKeyboardEventType_KEYDOWN:
		return "keydown", true
	default:
		return "", false
	}
}

type SyntheticKeyboardEvent struct {
	// KeyCode is "SyntheticKeyboardEvent.keyCode"
	//
	// Required
	KeyCode int64
	// Modifiers is "SyntheticKeyboardEvent.modifiers"
	//
	// Optional
	//
	// NOTE: Modifiers.FFI_USE MUST be set to true to get Modifiers used.
	Modifiers SyntheticKeyboardModifiers
	// Type is "SyntheticKeyboardEvent.type"
	//
	// Required
	Type SyntheticKeyboardEventType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SyntheticKeyboardEvent with all fields set.
func (p SyntheticKeyboardEvent) FromRef(ref js.Ref) SyntheticKeyboardEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SyntheticKeyboardEvent in the application heap.
func (p SyntheticKeyboardEvent) New() js.Ref {
	return bindings.SyntheticKeyboardEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SyntheticKeyboardEvent) UpdateFrom(ref js.Ref) {
	bindings.SyntheticKeyboardEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SyntheticKeyboardEvent) Update(ref js.Ref) {
	bindings.SyntheticKeyboardEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SyntheticKeyboardEvent) FreeMembers(recursive bool) {
	if recursive {
		p.Modifiers.FreeMembers(true)
	}
}

type SyntheticMouseEventButton uint32

const (
	_ SyntheticMouseEventButton = iota

	SyntheticMouseEventButton_LEFT
	SyntheticMouseEventButton_MIDDLE
	SyntheticMouseEventButton_RIGHT
	SyntheticMouseEventButton_BACK
	SyntheticMouseEventButton_FOWARD
)

func (SyntheticMouseEventButton) FromRef(str js.Ref) SyntheticMouseEventButton {
	return SyntheticMouseEventButton(bindings.ConstOfSyntheticMouseEventButton(str))
}

func (x SyntheticMouseEventButton) String() (string, bool) {
	switch x {
	case SyntheticMouseEventButton_LEFT:
		return "left", true
	case SyntheticMouseEventButton_MIDDLE:
		return "middle", true
	case SyntheticMouseEventButton_RIGHT:
		return "right", true
	case SyntheticMouseEventButton_BACK:
		return "back", true
	case SyntheticMouseEventButton_FOWARD:
		return "foward", true
	default:
		return "", false
	}
}

type SyntheticMouseEventType uint32

const (
	_ SyntheticMouseEventType = iota

	SyntheticMouseEventType_PRESS
	SyntheticMouseEventType_RELEASE
	SyntheticMouseEventType_DRAG
	SyntheticMouseEventType_MOVE
	SyntheticMouseEventType_ENTER
	SyntheticMouseEventType_EXIT
)

func (SyntheticMouseEventType) FromRef(str js.Ref) SyntheticMouseEventType {
	return SyntheticMouseEventType(bindings.ConstOfSyntheticMouseEventType(str))
}

func (x SyntheticMouseEventType) String() (string, bool) {
	switch x {
	case SyntheticMouseEventType_PRESS:
		return "press", true
	case SyntheticMouseEventType_RELEASE:
		return "release", true
	case SyntheticMouseEventType_DRAG:
		return "drag", true
	case SyntheticMouseEventType_MOVE:
		return "move", true
	case SyntheticMouseEventType_ENTER:
		return "enter", true
	case SyntheticMouseEventType_EXIT:
		return "exit", true
	default:
		return "", false
	}
}

type SyntheticMouseEvent struct {
	// MouseButton is "SyntheticMouseEvent.mouseButton"
	//
	// Optional
	MouseButton SyntheticMouseEventButton
	// TouchAccessibility is "SyntheticMouseEvent.touchAccessibility"
	//
	// Optional
	//
	// NOTE: FFI_USE_TouchAccessibility MUST be set to true to make this field effective.
	TouchAccessibility bool
	// Type is "SyntheticMouseEvent.type"
	//
	// Required
	Type SyntheticMouseEventType
	// X is "SyntheticMouseEvent.x"
	//
	// Required
	X int64
	// Y is "SyntheticMouseEvent.y"
	//
	// Required
	Y int64

	FFI_USE_TouchAccessibility bool // for TouchAccessibility.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SyntheticMouseEvent with all fields set.
func (p SyntheticMouseEvent) FromRef(ref js.Ref) SyntheticMouseEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SyntheticMouseEvent in the application heap.
func (p SyntheticMouseEvent) New() js.Ref {
	return bindings.SyntheticMouseEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SyntheticMouseEvent) UpdateFrom(ref js.Ref) {
	bindings.SyntheticMouseEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SyntheticMouseEvent) Update(ref js.Ref) {
	bindings.SyntheticMouseEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SyntheticMouseEvent) FreeMembers(recursive bool) {
}

type ToastType uint32

const (
	_ ToastType = iota

	ToastType_DICTATION_NO_FOCUSED_TEXT_FIELD
)

func (ToastType) FromRef(str js.Ref) ToastType {
	return ToastType(bindings.ConstOfToastType(str))
}

func (x ToastType) String() (string, bool) {
	switch x {
	case ToastType_DICTATION_NO_FOCUSED_TEXT_FIELD:
		return "dictationNoFocusedTextField", true
	default:
		return "", false
	}
}

// HasFuncDarkenScreen returns true if the function "WEBEXT.accessibilityPrivate.darkenScreen" exists.
func HasFuncDarkenScreen() bool {
	return js.True == bindings.HasFuncDarkenScreen()
}

// FuncDarkenScreen returns the function "WEBEXT.accessibilityPrivate.darkenScreen".
func FuncDarkenScreen() (fn js.Func[func(enabled bool)]) {
	bindings.FuncDarkenScreen(
		js.Pointer(&fn),
	)
	return
}

// DarkenScreen calls the function "WEBEXT.accessibilityPrivate.darkenScreen" directly.
func DarkenScreen(enabled bool) (ret js.Void) {
	bindings.CallDarkenScreen(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TryDarkenScreen calls the function "WEBEXT.accessibilityPrivate.darkenScreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDarkenScreen(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDarkenScreen(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncEnableMouseEvents returns true if the function "WEBEXT.accessibilityPrivate.enableMouseEvents" exists.
func HasFuncEnableMouseEvents() bool {
	return js.True == bindings.HasFuncEnableMouseEvents()
}

// FuncEnableMouseEvents returns the function "WEBEXT.accessibilityPrivate.enableMouseEvents".
func FuncEnableMouseEvents() (fn js.Func[func(enabled bool)]) {
	bindings.FuncEnableMouseEvents(
		js.Pointer(&fn),
	)
	return
}

// EnableMouseEvents calls the function "WEBEXT.accessibilityPrivate.enableMouseEvents" directly.
func EnableMouseEvents(enabled bool) (ret js.Void) {
	bindings.CallEnableMouseEvents(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TryEnableMouseEvents calls the function "WEBEXT.accessibilityPrivate.enableMouseEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableMouseEvents(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableMouseEvents(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncForwardKeyEventsToSwitchAccess returns true if the function "WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess" exists.
func HasFuncForwardKeyEventsToSwitchAccess() bool {
	return js.True == bindings.HasFuncForwardKeyEventsToSwitchAccess()
}

// FuncForwardKeyEventsToSwitchAccess returns the function "WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess".
func FuncForwardKeyEventsToSwitchAccess() (fn js.Func[func(shouldForward bool)]) {
	bindings.FuncForwardKeyEventsToSwitchAccess(
		js.Pointer(&fn),
	)
	return
}

// ForwardKeyEventsToSwitchAccess calls the function "WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess" directly.
func ForwardKeyEventsToSwitchAccess(shouldForward bool) (ret js.Void) {
	bindings.CallForwardKeyEventsToSwitchAccess(
		js.Pointer(&ret),
		js.Bool(bool(shouldForward)),
	)

	return
}

// TryForwardKeyEventsToSwitchAccess calls the function "WEBEXT.accessibilityPrivate.forwardKeyEventsToSwitchAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForwardKeyEventsToSwitchAccess(shouldForward bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryForwardKeyEventsToSwitchAccess(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(shouldForward)),
	)

	return
}

// HasFuncGetBatteryDescription returns true if the function "WEBEXT.accessibilityPrivate.getBatteryDescription" exists.
func HasFuncGetBatteryDescription() bool {
	return js.True == bindings.HasFuncGetBatteryDescription()
}

// FuncGetBatteryDescription returns the function "WEBEXT.accessibilityPrivate.getBatteryDescription".
func FuncGetBatteryDescription() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetBatteryDescription(
		js.Pointer(&fn),
	)
	return
}

// GetBatteryDescription calls the function "WEBEXT.accessibilityPrivate.getBatteryDescription" directly.
func GetBatteryDescription() (ret js.Promise[js.String]) {
	bindings.CallGetBatteryDescription(
		js.Pointer(&ret),
	)

	return
}

// TryGetBatteryDescription calls the function "WEBEXT.accessibilityPrivate.getBatteryDescription"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBatteryDescription() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBatteryDescription(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDisplayNameForLocale returns true if the function "WEBEXT.accessibilityPrivate.getDisplayNameForLocale" exists.
func HasFuncGetDisplayNameForLocale() bool {
	return js.True == bindings.HasFuncGetDisplayNameForLocale()
}

// FuncGetDisplayNameForLocale returns the function "WEBEXT.accessibilityPrivate.getDisplayNameForLocale".
func FuncGetDisplayNameForLocale() (fn js.Func[func(localeCodeToTranslate js.String, displayLocaleCode js.String) js.String]) {
	bindings.FuncGetDisplayNameForLocale(
		js.Pointer(&fn),
	)
	return
}

// GetDisplayNameForLocale calls the function "WEBEXT.accessibilityPrivate.getDisplayNameForLocale" directly.
func GetDisplayNameForLocale(localeCodeToTranslate js.String, displayLocaleCode js.String) (ret js.String) {
	bindings.CallGetDisplayNameForLocale(
		js.Pointer(&ret),
		localeCodeToTranslate.Ref(),
		displayLocaleCode.Ref(),
	)

	return
}

// TryGetDisplayNameForLocale calls the function "WEBEXT.accessibilityPrivate.getDisplayNameForLocale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplayNameForLocale(localeCodeToTranslate js.String, displayLocaleCode js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplayNameForLocale(
		js.Pointer(&ret), js.Pointer(&exception),
		localeCodeToTranslate.Ref(),
		displayLocaleCode.Ref(),
	)

	return
}

// HasFuncGetDlcContents returns true if the function "WEBEXT.accessibilityPrivate.getDlcContents" exists.
func HasFuncGetDlcContents() bool {
	return js.True == bindings.HasFuncGetDlcContents()
}

// FuncGetDlcContents returns the function "WEBEXT.accessibilityPrivate.getDlcContents".
func FuncGetDlcContents() (fn js.Func[func(dlc DlcType) js.Promise[js.TypedArray[uint8]]]) {
	bindings.FuncGetDlcContents(
		js.Pointer(&fn),
	)
	return
}

// GetDlcContents calls the function "WEBEXT.accessibilityPrivate.getDlcContents" directly.
func GetDlcContents(dlc DlcType) (ret js.Promise[js.TypedArray[uint8]]) {
	bindings.CallGetDlcContents(
		js.Pointer(&ret),
		uint32(dlc),
	)

	return
}

// TryGetDlcContents calls the function "WEBEXT.accessibilityPrivate.getDlcContents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDlcContents(dlc DlcType) (ret js.Promise[js.TypedArray[uint8]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDlcContents(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(dlc),
	)

	return
}

// HasFuncGetLocalizedDomKeyStringForKeyCode returns true if the function "WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode" exists.
func HasFuncGetLocalizedDomKeyStringForKeyCode() bool {
	return js.True == bindings.HasFuncGetLocalizedDomKeyStringForKeyCode()
}

// FuncGetLocalizedDomKeyStringForKeyCode returns the function "WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode".
func FuncGetLocalizedDomKeyStringForKeyCode() (fn js.Func[func(keyCode int64) js.Promise[js.String]]) {
	bindings.FuncGetLocalizedDomKeyStringForKeyCode(
		js.Pointer(&fn),
	)
	return
}

// GetLocalizedDomKeyStringForKeyCode calls the function "WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode" directly.
func GetLocalizedDomKeyStringForKeyCode(keyCode int64) (ret js.Promise[js.String]) {
	bindings.CallGetLocalizedDomKeyStringForKeyCode(
		js.Pointer(&ret),
		float64(keyCode),
	)

	return
}

// TryGetLocalizedDomKeyStringForKeyCode calls the function "WEBEXT.accessibilityPrivate.getLocalizedDomKeyStringForKeyCode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLocalizedDomKeyStringForKeyCode(keyCode int64) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLocalizedDomKeyStringForKeyCode(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(keyCode),
	)

	return
}

// HasFuncHandleScrollableBoundsForPointFound returns true if the function "WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound" exists.
func HasFuncHandleScrollableBoundsForPointFound() bool {
	return js.True == bindings.HasFuncHandleScrollableBoundsForPointFound()
}

// FuncHandleScrollableBoundsForPointFound returns the function "WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound".
func FuncHandleScrollableBoundsForPointFound() (fn js.Func[func(rect ScreenRect)]) {
	bindings.FuncHandleScrollableBoundsForPointFound(
		js.Pointer(&fn),
	)
	return
}

// HandleScrollableBoundsForPointFound calls the function "WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound" directly.
func HandleScrollableBoundsForPointFound(rect ScreenRect) (ret js.Void) {
	bindings.CallHandleScrollableBoundsForPointFound(
		js.Pointer(&ret),
		js.Pointer(&rect),
	)

	return
}

// TryHandleScrollableBoundsForPointFound calls the function "WEBEXT.accessibilityPrivate.handleScrollableBoundsForPointFound"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHandleScrollableBoundsForPointFound(rect ScreenRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHandleScrollableBoundsForPointFound(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&rect),
	)

	return
}

// HasFuncInstallPumpkinForDictation returns true if the function "WEBEXT.accessibilityPrivate.installPumpkinForDictation" exists.
func HasFuncInstallPumpkinForDictation() bool {
	return js.True == bindings.HasFuncInstallPumpkinForDictation()
}

// FuncInstallPumpkinForDictation returns the function "WEBEXT.accessibilityPrivate.installPumpkinForDictation".
func FuncInstallPumpkinForDictation() (fn js.Func[func() js.Promise[PumpkinData]]) {
	bindings.FuncInstallPumpkinForDictation(
		js.Pointer(&fn),
	)
	return
}

// InstallPumpkinForDictation calls the function "WEBEXT.accessibilityPrivate.installPumpkinForDictation" directly.
func InstallPumpkinForDictation() (ret js.Promise[PumpkinData]) {
	bindings.CallInstallPumpkinForDictation(
		js.Pointer(&ret),
	)

	return
}

// TryInstallPumpkinForDictation calls the function "WEBEXT.accessibilityPrivate.installPumpkinForDictation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallPumpkinForDictation() (ret js.Promise[PumpkinData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallPumpkinForDictation(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsFeatureEnabled returns true if the function "WEBEXT.accessibilityPrivate.isFeatureEnabled" exists.
func HasFuncIsFeatureEnabled() bool {
	return js.True == bindings.HasFuncIsFeatureEnabled()
}

// FuncIsFeatureEnabled returns the function "WEBEXT.accessibilityPrivate.isFeatureEnabled".
func FuncIsFeatureEnabled() (fn js.Func[func(feature AccessibilityFeature) js.Promise[js.Boolean]]) {
	bindings.FuncIsFeatureEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsFeatureEnabled calls the function "WEBEXT.accessibilityPrivate.isFeatureEnabled" directly.
func IsFeatureEnabled(feature AccessibilityFeature) (ret js.Promise[js.Boolean]) {
	bindings.CallIsFeatureEnabled(
		js.Pointer(&ret),
		uint32(feature),
	)

	return
}

// TryIsFeatureEnabled calls the function "WEBEXT.accessibilityPrivate.isFeatureEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsFeatureEnabled(feature AccessibilityFeature) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsFeatureEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(feature),
	)

	return
}

// HasFuncIsLacrosPrimary returns true if the function "WEBEXT.accessibilityPrivate.isLacrosPrimary" exists.
func HasFuncIsLacrosPrimary() bool {
	return js.True == bindings.HasFuncIsLacrosPrimary()
}

// FuncIsLacrosPrimary returns the function "WEBEXT.accessibilityPrivate.isLacrosPrimary".
func FuncIsLacrosPrimary() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsLacrosPrimary(
		js.Pointer(&fn),
	)
	return
}

// IsLacrosPrimary calls the function "WEBEXT.accessibilityPrivate.isLacrosPrimary" directly.
func IsLacrosPrimary() (ret js.Promise[js.Boolean]) {
	bindings.CallIsLacrosPrimary(
		js.Pointer(&ret),
	)

	return
}

// TryIsLacrosPrimary calls the function "WEBEXT.accessibilityPrivate.isLacrosPrimary"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsLacrosPrimary() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsLacrosPrimary(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMagnifierCenterOnPoint returns true if the function "WEBEXT.accessibilityPrivate.magnifierCenterOnPoint" exists.
func HasFuncMagnifierCenterOnPoint() bool {
	return js.True == bindings.HasFuncMagnifierCenterOnPoint()
}

// FuncMagnifierCenterOnPoint returns the function "WEBEXT.accessibilityPrivate.magnifierCenterOnPoint".
func FuncMagnifierCenterOnPoint() (fn js.Func[func(point ScreenPoint)]) {
	bindings.FuncMagnifierCenterOnPoint(
		js.Pointer(&fn),
	)
	return
}

// MagnifierCenterOnPoint calls the function "WEBEXT.accessibilityPrivate.magnifierCenterOnPoint" directly.
func MagnifierCenterOnPoint(point ScreenPoint) (ret js.Void) {
	bindings.CallMagnifierCenterOnPoint(
		js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryMagnifierCenterOnPoint calls the function "WEBEXT.accessibilityPrivate.magnifierCenterOnPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMagnifierCenterOnPoint(point ScreenPoint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMagnifierCenterOnPoint(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasFuncMoveMagnifierToRect returns true if the function "WEBEXT.accessibilityPrivate.moveMagnifierToRect" exists.
func HasFuncMoveMagnifierToRect() bool {
	return js.True == bindings.HasFuncMoveMagnifierToRect()
}

// FuncMoveMagnifierToRect returns the function "WEBEXT.accessibilityPrivate.moveMagnifierToRect".
func FuncMoveMagnifierToRect() (fn js.Func[func(rect ScreenRect)]) {
	bindings.FuncMoveMagnifierToRect(
		js.Pointer(&fn),
	)
	return
}

// MoveMagnifierToRect calls the function "WEBEXT.accessibilityPrivate.moveMagnifierToRect" directly.
func MoveMagnifierToRect(rect ScreenRect) (ret js.Void) {
	bindings.CallMoveMagnifierToRect(
		js.Pointer(&ret),
		js.Pointer(&rect),
	)

	return
}

// TryMoveMagnifierToRect calls the function "WEBEXT.accessibilityPrivate.moveMagnifierToRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMoveMagnifierToRect(rect ScreenRect) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMoveMagnifierToRect(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&rect),
	)

	return
}

type OnAccessibilityGestureEventCallbackFunc func(this js.Ref, gesture Gesture, x int64, y int64) js.Ref

func (fn OnAccessibilityGestureEventCallbackFunc) Register() js.Func[func(gesture Gesture, x int64, y int64)] {
	return js.RegisterCallback[func(gesture Gesture, x int64, y int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAccessibilityGestureEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Gesture(0).FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAccessibilityGestureEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, gesture Gesture, x int64, y int64) js.Ref
	Arg T
}

func (cb *OnAccessibilityGestureEventCallback[T]) Register() js.Func[func(gesture Gesture, x int64, y int64)] {
	return js.RegisterCallback[func(gesture Gesture, x int64, y int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAccessibilityGestureEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		Gesture(0).FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAccessibilityGesture returns true if the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener" exists.
func HasFuncOnAccessibilityGesture() bool {
	return js.True == bindings.HasFuncOnAccessibilityGesture()
}

// FuncOnAccessibilityGesture returns the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener".
func FuncOnAccessibilityGesture() (fn js.Func[func(callback js.Func[func(gesture Gesture, x int64, y int64)])]) {
	bindings.FuncOnAccessibilityGesture(
		js.Pointer(&fn),
	)
	return
}

// OnAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener" directly.
func OnAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret js.Void) {
	bindings.CallOnAccessibilityGesture(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccessibilityGesture(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccessibilityGesture returns true if the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener" exists.
func HasFuncOffAccessibilityGesture() bool {
	return js.True == bindings.HasFuncOffAccessibilityGesture()
}

// FuncOffAccessibilityGesture returns the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener".
func FuncOffAccessibilityGesture() (fn js.Func[func(callback js.Func[func(gesture Gesture, x int64, y int64)])]) {
	bindings.FuncOffAccessibilityGesture(
		js.Pointer(&fn),
	)
	return
}

// OffAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener" directly.
func OffAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret js.Void) {
	bindings.CallOffAccessibilityGesture(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccessibilityGesture(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccessibilityGesture returns true if the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener" exists.
func HasFuncHasOnAccessibilityGesture() bool {
	return js.True == bindings.HasFuncHasOnAccessibilityGesture()
}

// FuncHasOnAccessibilityGesture returns the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener".
func FuncHasOnAccessibilityGesture() (fn js.Func[func(callback js.Func[func(gesture Gesture, x int64, y int64)]) bool]) {
	bindings.FuncHasOnAccessibilityGesture(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener" directly.
func HasOnAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret bool) {
	bindings.CallHasOnAccessibilityGesture(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccessibilityGesture calls the function "WEBEXT.accessibilityPrivate.onAccessibilityGesture.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccessibilityGesture(callback js.Func[func(gesture Gesture, x int64, y int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccessibilityGesture(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAnnounceForAccessibilityEventCallbackFunc func(this js.Ref, announceText js.Array[js.String]) js.Ref

func (fn OnAnnounceForAccessibilityEventCallbackFunc) Register() js.Func[func(announceText js.Array[js.String])] {
	return js.RegisterCallback[func(announceText js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAnnounceForAccessibilityEventCallbackFunc) DispatchCallback(
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

type OnAnnounceForAccessibilityEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, announceText js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnAnnounceForAccessibilityEventCallback[T]) Register() js.Func[func(announceText js.Array[js.String])] {
	return js.RegisterCallback[func(announceText js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAnnounceForAccessibilityEventCallback[T]) DispatchCallback(
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

// HasFuncOnAnnounceForAccessibility returns true if the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener" exists.
func HasFuncOnAnnounceForAccessibility() bool {
	return js.True == bindings.HasFuncOnAnnounceForAccessibility()
}

// FuncOnAnnounceForAccessibility returns the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener".
func FuncOnAnnounceForAccessibility() (fn js.Func[func(callback js.Func[func(announceText js.Array[js.String])])]) {
	bindings.FuncOnAnnounceForAccessibility(
		js.Pointer(&fn),
	)
	return
}

// OnAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener" directly.
func OnAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnAnnounceForAccessibility(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAnnounceForAccessibility(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAnnounceForAccessibility returns true if the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener" exists.
func HasFuncOffAnnounceForAccessibility() bool {
	return js.True == bindings.HasFuncOffAnnounceForAccessibility()
}

// FuncOffAnnounceForAccessibility returns the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener".
func FuncOffAnnounceForAccessibility() (fn js.Func[func(callback js.Func[func(announceText js.Array[js.String])])]) {
	bindings.FuncOffAnnounceForAccessibility(
		js.Pointer(&fn),
	)
	return
}

// OffAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener" directly.
func OffAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffAnnounceForAccessibility(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAnnounceForAccessibility(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAnnounceForAccessibility returns true if the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener" exists.
func HasFuncHasOnAnnounceForAccessibility() bool {
	return js.True == bindings.HasFuncHasOnAnnounceForAccessibility()
}

// FuncHasOnAnnounceForAccessibility returns the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener".
func FuncHasOnAnnounceForAccessibility() (fn js.Func[func(callback js.Func[func(announceText js.Array[js.String])]) bool]) {
	bindings.FuncHasOnAnnounceForAccessibility(
		js.Pointer(&fn),
	)
	return
}

// HasOnAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener" directly.
func HasOnAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnAnnounceForAccessibility(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAnnounceForAccessibility calls the function "WEBEXT.accessibilityPrivate.onAnnounceForAccessibility.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAnnounceForAccessibility(callback js.Func[func(announceText js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAnnounceForAccessibility(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCustomSpokenFeedbackToggledEventCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn OnCustomSpokenFeedbackToggledEventCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCustomSpokenFeedbackToggledEventCallbackFunc) DispatchCallback(
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

type OnCustomSpokenFeedbackToggledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *OnCustomSpokenFeedbackToggledEventCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCustomSpokenFeedbackToggledEventCallback[T]) DispatchCallback(
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

// HasFuncOnCustomSpokenFeedbackToggled returns true if the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener" exists.
func HasFuncOnCustomSpokenFeedbackToggled() bool {
	return js.True == bindings.HasFuncOnCustomSpokenFeedbackToggled()
}

// FuncOnCustomSpokenFeedbackToggled returns the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener".
func FuncOnCustomSpokenFeedbackToggled() (fn js.Func[func(callback js.Func[func(enabled bool)])]) {
	bindings.FuncOnCustomSpokenFeedbackToggled(
		js.Pointer(&fn),
	)
	return
}

// OnCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener" directly.
func OnCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret js.Void) {
	bindings.CallOnCustomSpokenFeedbackToggled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCustomSpokenFeedbackToggled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCustomSpokenFeedbackToggled returns true if the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener" exists.
func HasFuncOffCustomSpokenFeedbackToggled() bool {
	return js.True == bindings.HasFuncOffCustomSpokenFeedbackToggled()
}

// FuncOffCustomSpokenFeedbackToggled returns the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener".
func FuncOffCustomSpokenFeedbackToggled() (fn js.Func[func(callback js.Func[func(enabled bool)])]) {
	bindings.FuncOffCustomSpokenFeedbackToggled(
		js.Pointer(&fn),
	)
	return
}

// OffCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener" directly.
func OffCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret js.Void) {
	bindings.CallOffCustomSpokenFeedbackToggled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCustomSpokenFeedbackToggled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCustomSpokenFeedbackToggled returns true if the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener" exists.
func HasFuncHasOnCustomSpokenFeedbackToggled() bool {
	return js.True == bindings.HasFuncHasOnCustomSpokenFeedbackToggled()
}

// FuncHasOnCustomSpokenFeedbackToggled returns the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener".
func FuncHasOnCustomSpokenFeedbackToggled() (fn js.Func[func(callback js.Func[func(enabled bool)]) bool]) {
	bindings.FuncHasOnCustomSpokenFeedbackToggled(
		js.Pointer(&fn),
	)
	return
}

// HasOnCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener" directly.
func HasOnCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret bool) {
	bindings.CallHasOnCustomSpokenFeedbackToggled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCustomSpokenFeedbackToggled calls the function "WEBEXT.accessibilityPrivate.onCustomSpokenFeedbackToggled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCustomSpokenFeedbackToggled(callback js.Func[func(enabled bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCustomSpokenFeedbackToggled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnIntroduceChromeVoxEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnIntroduceChromeVoxEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnIntroduceChromeVoxEventCallbackFunc) DispatchCallback(
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

type OnIntroduceChromeVoxEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnIntroduceChromeVoxEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnIntroduceChromeVoxEventCallback[T]) DispatchCallback(
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

// HasFuncOnIntroduceChromeVox returns true if the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener" exists.
func HasFuncOnIntroduceChromeVox() bool {
	return js.True == bindings.HasFuncOnIntroduceChromeVox()
}

// FuncOnIntroduceChromeVox returns the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener".
func FuncOnIntroduceChromeVox() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnIntroduceChromeVox(
		js.Pointer(&fn),
	)
	return
}

// OnIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener" directly.
func OnIntroduceChromeVox(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnIntroduceChromeVox(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnIntroduceChromeVox(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnIntroduceChromeVox(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffIntroduceChromeVox returns true if the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener" exists.
func HasFuncOffIntroduceChromeVox() bool {
	return js.True == bindings.HasFuncOffIntroduceChromeVox()
}

// FuncOffIntroduceChromeVox returns the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener".
func FuncOffIntroduceChromeVox() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffIntroduceChromeVox(
		js.Pointer(&fn),
	)
	return
}

// OffIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener" directly.
func OffIntroduceChromeVox(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffIntroduceChromeVox(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffIntroduceChromeVox(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffIntroduceChromeVox(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnIntroduceChromeVox returns true if the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener" exists.
func HasFuncHasOnIntroduceChromeVox() bool {
	return js.True == bindings.HasFuncHasOnIntroduceChromeVox()
}

// FuncHasOnIntroduceChromeVox returns the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener".
func FuncHasOnIntroduceChromeVox() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnIntroduceChromeVox(
		js.Pointer(&fn),
	)
	return
}

// HasOnIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener" directly.
func HasOnIntroduceChromeVox(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnIntroduceChromeVox(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnIntroduceChromeVox calls the function "WEBEXT.accessibilityPrivate.onIntroduceChromeVox.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnIntroduceChromeVox(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnIntroduceChromeVox(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMagnifierBoundsChangedEventCallbackFunc func(this js.Ref, magnifierBounds *ScreenRect) js.Ref

func (fn OnMagnifierBoundsChangedEventCallbackFunc) Register() js.Func[func(magnifierBounds *ScreenRect)] {
	return js.RegisterCallback[func(magnifierBounds *ScreenRect)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMagnifierBoundsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScreenRect
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

type OnMagnifierBoundsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, magnifierBounds *ScreenRect) js.Ref
	Arg T
}

func (cb *OnMagnifierBoundsChangedEventCallback[T]) Register() js.Func[func(magnifierBounds *ScreenRect)] {
	return js.RegisterCallback[func(magnifierBounds *ScreenRect)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMagnifierBoundsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScreenRect
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

// HasFuncOnMagnifierBoundsChanged returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener" exists.
func HasFuncOnMagnifierBoundsChanged() bool {
	return js.True == bindings.HasFuncOnMagnifierBoundsChanged()
}

// FuncOnMagnifierBoundsChanged returns the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener".
func FuncOnMagnifierBoundsChanged() (fn js.Func[func(callback js.Func[func(magnifierBounds *ScreenRect)])]) {
	bindings.FuncOnMagnifierBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener" directly.
func OnMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret js.Void) {
	bindings.CallOnMagnifierBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMagnifierBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMagnifierBoundsChanged returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener" exists.
func HasFuncOffMagnifierBoundsChanged() bool {
	return js.True == bindings.HasFuncOffMagnifierBoundsChanged()
}

// FuncOffMagnifierBoundsChanged returns the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener".
func FuncOffMagnifierBoundsChanged() (fn js.Func[func(callback js.Func[func(magnifierBounds *ScreenRect)])]) {
	bindings.FuncOffMagnifierBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener" directly.
func OffMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret js.Void) {
	bindings.CallOffMagnifierBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMagnifierBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMagnifierBoundsChanged returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener" exists.
func HasFuncHasOnMagnifierBoundsChanged() bool {
	return js.True == bindings.HasFuncHasOnMagnifierBoundsChanged()
}

// FuncHasOnMagnifierBoundsChanged returns the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener".
func FuncHasOnMagnifierBoundsChanged() (fn js.Func[func(callback js.Func[func(magnifierBounds *ScreenRect)]) bool]) {
	bindings.FuncHasOnMagnifierBoundsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener" directly.
func HasOnMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret bool) {
	bindings.CallHasOnMagnifierBoundsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMagnifierBoundsChanged calls the function "WEBEXT.accessibilityPrivate.onMagnifierBoundsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMagnifierBoundsChanged(callback js.Func[func(magnifierBounds *ScreenRect)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMagnifierBoundsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMagnifierCommandEventCallbackFunc func(this js.Ref, command MagnifierCommand) js.Ref

func (fn OnMagnifierCommandEventCallbackFunc) Register() js.Func[func(command MagnifierCommand)] {
	return js.RegisterCallback[func(command MagnifierCommand)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMagnifierCommandEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		MagnifierCommand(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMagnifierCommandEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, command MagnifierCommand) js.Ref
	Arg T
}

func (cb *OnMagnifierCommandEventCallback[T]) Register() js.Func[func(command MagnifierCommand)] {
	return js.RegisterCallback[func(command MagnifierCommand)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMagnifierCommandEventCallback[T]) DispatchCallback(
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

		MagnifierCommand(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMagnifierCommand returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener" exists.
func HasFuncOnMagnifierCommand() bool {
	return js.True == bindings.HasFuncOnMagnifierCommand()
}

// FuncOnMagnifierCommand returns the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener".
func FuncOnMagnifierCommand() (fn js.Func[func(callback js.Func[func(command MagnifierCommand)])]) {
	bindings.FuncOnMagnifierCommand(
		js.Pointer(&fn),
	)
	return
}

// OnMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener" directly.
func OnMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret js.Void) {
	bindings.CallOnMagnifierCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMagnifierCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMagnifierCommand returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener" exists.
func HasFuncOffMagnifierCommand() bool {
	return js.True == bindings.HasFuncOffMagnifierCommand()
}

// FuncOffMagnifierCommand returns the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener".
func FuncOffMagnifierCommand() (fn js.Func[func(callback js.Func[func(command MagnifierCommand)])]) {
	bindings.FuncOffMagnifierCommand(
		js.Pointer(&fn),
	)
	return
}

// OffMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener" directly.
func OffMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret js.Void) {
	bindings.CallOffMagnifierCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMagnifierCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMagnifierCommand returns true if the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener" exists.
func HasFuncHasOnMagnifierCommand() bool {
	return js.True == bindings.HasFuncHasOnMagnifierCommand()
}

// FuncHasOnMagnifierCommand returns the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener".
func FuncHasOnMagnifierCommand() (fn js.Func[func(callback js.Func[func(command MagnifierCommand)]) bool]) {
	bindings.FuncHasOnMagnifierCommand(
		js.Pointer(&fn),
	)
	return
}

// HasOnMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener" directly.
func HasOnMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret bool) {
	bindings.CallHasOnMagnifierCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMagnifierCommand calls the function "WEBEXT.accessibilityPrivate.onMagnifierCommand.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMagnifierCommand(callback js.Func[func(command MagnifierCommand)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMagnifierCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPointScanSetEventCallbackFunc func(this js.Ref, point *PointScanPoint) js.Ref

func (fn OnPointScanSetEventCallbackFunc) Register() js.Func[func(point *PointScanPoint)] {
	return js.RegisterCallback[func(point *PointScanPoint)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPointScanSetEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PointScanPoint
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

type OnPointScanSetEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, point *PointScanPoint) js.Ref
	Arg T
}

func (cb *OnPointScanSetEventCallback[T]) Register() js.Func[func(point *PointScanPoint)] {
	return js.RegisterCallback[func(point *PointScanPoint)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPointScanSetEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PointScanPoint
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

// HasFuncOnPointScanSet returns true if the function "WEBEXT.accessibilityPrivate.onPointScanSet.addListener" exists.
func HasFuncOnPointScanSet() bool {
	return js.True == bindings.HasFuncOnPointScanSet()
}

// FuncOnPointScanSet returns the function "WEBEXT.accessibilityPrivate.onPointScanSet.addListener".
func FuncOnPointScanSet() (fn js.Func[func(callback js.Func[func(point *PointScanPoint)])]) {
	bindings.FuncOnPointScanSet(
		js.Pointer(&fn),
	)
	return
}

// OnPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.addListener" directly.
func OnPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret js.Void) {
	bindings.CallOnPointScanSet(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPointScanSet(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPointScanSet returns true if the function "WEBEXT.accessibilityPrivate.onPointScanSet.removeListener" exists.
func HasFuncOffPointScanSet() bool {
	return js.True == bindings.HasFuncOffPointScanSet()
}

// FuncOffPointScanSet returns the function "WEBEXT.accessibilityPrivate.onPointScanSet.removeListener".
func FuncOffPointScanSet() (fn js.Func[func(callback js.Func[func(point *PointScanPoint)])]) {
	bindings.FuncOffPointScanSet(
		js.Pointer(&fn),
	)
	return
}

// OffPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.removeListener" directly.
func OffPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret js.Void) {
	bindings.CallOffPointScanSet(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPointScanSet(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPointScanSet returns true if the function "WEBEXT.accessibilityPrivate.onPointScanSet.hasListener" exists.
func HasFuncHasOnPointScanSet() bool {
	return js.True == bindings.HasFuncHasOnPointScanSet()
}

// FuncHasOnPointScanSet returns the function "WEBEXT.accessibilityPrivate.onPointScanSet.hasListener".
func FuncHasOnPointScanSet() (fn js.Func[func(callback js.Func[func(point *PointScanPoint)]) bool]) {
	bindings.FuncHasOnPointScanSet(
		js.Pointer(&fn),
	)
	return
}

// HasOnPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.hasListener" directly.
func HasOnPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret bool) {
	bindings.CallHasOnPointScanSet(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPointScanSet calls the function "WEBEXT.accessibilityPrivate.onPointScanSet.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPointScanSet(callback js.Func[func(point *PointScanPoint)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPointScanSet(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnScrollableBoundsForPointRequestedEventCallbackFunc func(this js.Ref, x float64, y float64) js.Ref

func (fn OnScrollableBoundsForPointRequestedEventCallbackFunc) Register() js.Func[func(x float64, y float64)] {
	return js.RegisterCallback[func(x float64, y float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnScrollableBoundsForPointRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnScrollableBoundsForPointRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, x float64, y float64) js.Ref
	Arg T
}

func (cb *OnScrollableBoundsForPointRequestedEventCallback[T]) Register() js.Func[func(x float64, y float64)] {
	return js.RegisterCallback[func(x float64, y float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnScrollableBoundsForPointRequestedEventCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnScrollableBoundsForPointRequested returns true if the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener" exists.
func HasFuncOnScrollableBoundsForPointRequested() bool {
	return js.True == bindings.HasFuncOnScrollableBoundsForPointRequested()
}

// FuncOnScrollableBoundsForPointRequested returns the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener".
func FuncOnScrollableBoundsForPointRequested() (fn js.Func[func(callback js.Func[func(x float64, y float64)])]) {
	bindings.FuncOnScrollableBoundsForPointRequested(
		js.Pointer(&fn),
	)
	return
}

// OnScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener" directly.
func OnScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret js.Void) {
	bindings.CallOnScrollableBoundsForPointRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnScrollableBoundsForPointRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffScrollableBoundsForPointRequested returns true if the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener" exists.
func HasFuncOffScrollableBoundsForPointRequested() bool {
	return js.True == bindings.HasFuncOffScrollableBoundsForPointRequested()
}

// FuncOffScrollableBoundsForPointRequested returns the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener".
func FuncOffScrollableBoundsForPointRequested() (fn js.Func[func(callback js.Func[func(x float64, y float64)])]) {
	bindings.FuncOffScrollableBoundsForPointRequested(
		js.Pointer(&fn),
	)
	return
}

// OffScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener" directly.
func OffScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret js.Void) {
	bindings.CallOffScrollableBoundsForPointRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffScrollableBoundsForPointRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnScrollableBoundsForPointRequested returns true if the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener" exists.
func HasFuncHasOnScrollableBoundsForPointRequested() bool {
	return js.True == bindings.HasFuncHasOnScrollableBoundsForPointRequested()
}

// FuncHasOnScrollableBoundsForPointRequested returns the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener".
func FuncHasOnScrollableBoundsForPointRequested() (fn js.Func[func(callback js.Func[func(x float64, y float64)]) bool]) {
	bindings.FuncHasOnScrollableBoundsForPointRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener" directly.
func HasOnScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret bool) {
	bindings.CallHasOnScrollableBoundsForPointRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnScrollableBoundsForPointRequested calls the function "WEBEXT.accessibilityPrivate.onScrollableBoundsForPointRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnScrollableBoundsForPointRequested(callback js.Func[func(x float64, y float64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnScrollableBoundsForPointRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectToSpeakContextMenuClickedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnSelectToSpeakContextMenuClickedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectToSpeakContextMenuClickedEventCallbackFunc) DispatchCallback(
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

type OnSelectToSpeakContextMenuClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnSelectToSpeakContextMenuClickedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectToSpeakContextMenuClickedEventCallback[T]) DispatchCallback(
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

// HasFuncOnSelectToSpeakContextMenuClicked returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener" exists.
func HasFuncOnSelectToSpeakContextMenuClicked() bool {
	return js.True == bindings.HasFuncOnSelectToSpeakContextMenuClicked()
}

// FuncOnSelectToSpeakContextMenuClicked returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener".
func FuncOnSelectToSpeakContextMenuClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnSelectToSpeakContextMenuClicked(
		js.Pointer(&fn),
	)
	return
}

// OnSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener" directly.
func OnSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnSelectToSpeakContextMenuClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectToSpeakContextMenuClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectToSpeakContextMenuClicked returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener" exists.
func HasFuncOffSelectToSpeakContextMenuClicked() bool {
	return js.True == bindings.HasFuncOffSelectToSpeakContextMenuClicked()
}

// FuncOffSelectToSpeakContextMenuClicked returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener".
func FuncOffSelectToSpeakContextMenuClicked() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffSelectToSpeakContextMenuClicked(
		js.Pointer(&fn),
	)
	return
}

// OffSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener" directly.
func OffSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffSelectToSpeakContextMenuClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectToSpeakContextMenuClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectToSpeakContextMenuClicked returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener" exists.
func HasFuncHasOnSelectToSpeakContextMenuClicked() bool {
	return js.True == bindings.HasFuncHasOnSelectToSpeakContextMenuClicked()
}

// FuncHasOnSelectToSpeakContextMenuClicked returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener".
func FuncHasOnSelectToSpeakContextMenuClicked() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnSelectToSpeakContextMenuClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener" directly.
func HasOnSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnSelectToSpeakContextMenuClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectToSpeakContextMenuClicked calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakContextMenuClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectToSpeakContextMenuClicked(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectToSpeakContextMenuClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectToSpeakKeysPressedChangedEventCallbackFunc func(this js.Ref, keyCodes js.Array[int64]) js.Ref

func (fn OnSelectToSpeakKeysPressedChangedEventCallbackFunc) Register() js.Func[func(keyCodes js.Array[int64])] {
	return js.RegisterCallback[func(keyCodes js.Array[int64])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectToSpeakKeysPressedChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[int64]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSelectToSpeakKeysPressedChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, keyCodes js.Array[int64]) js.Ref
	Arg T
}

func (cb *OnSelectToSpeakKeysPressedChangedEventCallback[T]) Register() js.Func[func(keyCodes js.Array[int64])] {
	return js.RegisterCallback[func(keyCodes js.Array[int64])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectToSpeakKeysPressedChangedEventCallback[T]) DispatchCallback(
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

		js.Array[int64]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSelectToSpeakKeysPressedChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener" exists.
func HasFuncOnSelectToSpeakKeysPressedChanged() bool {
	return js.True == bindings.HasFuncOnSelectToSpeakKeysPressedChanged()
}

// FuncOnSelectToSpeakKeysPressedChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener".
func FuncOnSelectToSpeakKeysPressedChanged() (fn js.Func[func(callback js.Func[func(keyCodes js.Array[int64])])]) {
	bindings.FuncOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener" directly.
func OnSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret js.Void) {
	bindings.CallOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectToSpeakKeysPressedChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener" exists.
func HasFuncOffSelectToSpeakKeysPressedChanged() bool {
	return js.True == bindings.HasFuncOffSelectToSpeakKeysPressedChanged()
}

// FuncOffSelectToSpeakKeysPressedChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener".
func FuncOffSelectToSpeakKeysPressedChanged() (fn js.Func[func(callback js.Func[func(keyCodes js.Array[int64])])]) {
	bindings.FuncOffSelectToSpeakKeysPressedChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener" directly.
func OffSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret js.Void) {
	bindings.CallOffSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectToSpeakKeysPressedChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener" exists.
func HasFuncHasOnSelectToSpeakKeysPressedChanged() bool {
	return js.True == bindings.HasFuncHasOnSelectToSpeakKeysPressedChanged()
}

// FuncHasOnSelectToSpeakKeysPressedChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener".
func FuncHasOnSelectToSpeakKeysPressedChanged() (fn js.Func[func(callback js.Func[func(keyCodes js.Array[int64])]) bool]) {
	bindings.FuncHasOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener" directly.
func HasOnSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret bool) {
	bindings.CallHasOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectToSpeakKeysPressedChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakKeysPressedChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectToSpeakKeysPressedChanged(callback js.Func[func(keyCodes js.Array[int64])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectToSpeakKeysPressedChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectToSpeakMouseChangedEventCallbackFunc func(this js.Ref, typ SyntheticMouseEventType, x int64, y int64) js.Ref

func (fn OnSelectToSpeakMouseChangedEventCallbackFunc) Register() js.Func[func(typ SyntheticMouseEventType, x int64, y int64)] {
	return js.RegisterCallback[func(typ SyntheticMouseEventType, x int64, y int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectToSpeakMouseChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SyntheticMouseEventType(0).FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSelectToSpeakMouseChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, typ SyntheticMouseEventType, x int64, y int64) js.Ref
	Arg T
}

func (cb *OnSelectToSpeakMouseChangedEventCallback[T]) Register() js.Func[func(typ SyntheticMouseEventType, x int64, y int64)] {
	return js.RegisterCallback[func(typ SyntheticMouseEventType, x int64, y int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectToSpeakMouseChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		SyntheticMouseEventType(0).FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSelectToSpeakMouseChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener" exists.
func HasFuncOnSelectToSpeakMouseChanged() bool {
	return js.True == bindings.HasFuncOnSelectToSpeakMouseChanged()
}

// FuncOnSelectToSpeakMouseChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener".
func FuncOnSelectToSpeakMouseChanged() (fn js.Func[func(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)])]) {
	bindings.FuncOnSelectToSpeakMouseChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener" directly.
func OnSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret js.Void) {
	bindings.CallOnSelectToSpeakMouseChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectToSpeakMouseChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectToSpeakMouseChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener" exists.
func HasFuncOffSelectToSpeakMouseChanged() bool {
	return js.True == bindings.HasFuncOffSelectToSpeakMouseChanged()
}

// FuncOffSelectToSpeakMouseChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener".
func FuncOffSelectToSpeakMouseChanged() (fn js.Func[func(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)])]) {
	bindings.FuncOffSelectToSpeakMouseChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener" directly.
func OffSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret js.Void) {
	bindings.CallOffSelectToSpeakMouseChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectToSpeakMouseChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectToSpeakMouseChanged returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener" exists.
func HasFuncHasOnSelectToSpeakMouseChanged() bool {
	return js.True == bindings.HasFuncHasOnSelectToSpeakMouseChanged()
}

// FuncHasOnSelectToSpeakMouseChanged returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener".
func FuncHasOnSelectToSpeakMouseChanged() (fn js.Func[func(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) bool]) {
	bindings.FuncHasOnSelectToSpeakMouseChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener" directly.
func HasOnSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret bool) {
	bindings.CallHasOnSelectToSpeakMouseChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectToSpeakMouseChanged calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakMouseChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectToSpeakMouseChanged(callback js.Func[func(typ SyntheticMouseEventType, x int64, y int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectToSpeakMouseChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectToSpeakPanelActionEventCallbackFunc func(this js.Ref, action SelectToSpeakPanelAction, value float64) js.Ref

func (fn OnSelectToSpeakPanelActionEventCallbackFunc) Register() js.Func[func(action SelectToSpeakPanelAction, value float64)] {
	return js.RegisterCallback[func(action SelectToSpeakPanelAction, value float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectToSpeakPanelActionEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SelectToSpeakPanelAction(0).FromRef(args[0+1]),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSelectToSpeakPanelActionEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, action SelectToSpeakPanelAction, value float64) js.Ref
	Arg T
}

func (cb *OnSelectToSpeakPanelActionEventCallback[T]) Register() js.Func[func(action SelectToSpeakPanelAction, value float64)] {
	return js.RegisterCallback[func(action SelectToSpeakPanelAction, value float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectToSpeakPanelActionEventCallback[T]) DispatchCallback(
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

		SelectToSpeakPanelAction(0).FromRef(args[0+1]),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSelectToSpeakPanelAction returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener" exists.
func HasFuncOnSelectToSpeakPanelAction() bool {
	return js.True == bindings.HasFuncOnSelectToSpeakPanelAction()
}

// FuncOnSelectToSpeakPanelAction returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener".
func FuncOnSelectToSpeakPanelAction() (fn js.Func[func(callback js.Func[func(action SelectToSpeakPanelAction, value float64)])]) {
	bindings.FuncOnSelectToSpeakPanelAction(
		js.Pointer(&fn),
	)
	return
}

// OnSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener" directly.
func OnSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret js.Void) {
	bindings.CallOnSelectToSpeakPanelAction(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectToSpeakPanelAction(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectToSpeakPanelAction returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener" exists.
func HasFuncOffSelectToSpeakPanelAction() bool {
	return js.True == bindings.HasFuncOffSelectToSpeakPanelAction()
}

// FuncOffSelectToSpeakPanelAction returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener".
func FuncOffSelectToSpeakPanelAction() (fn js.Func[func(callback js.Func[func(action SelectToSpeakPanelAction, value float64)])]) {
	bindings.FuncOffSelectToSpeakPanelAction(
		js.Pointer(&fn),
	)
	return
}

// OffSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener" directly.
func OffSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret js.Void) {
	bindings.CallOffSelectToSpeakPanelAction(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectToSpeakPanelAction(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectToSpeakPanelAction returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener" exists.
func HasFuncHasOnSelectToSpeakPanelAction() bool {
	return js.True == bindings.HasFuncHasOnSelectToSpeakPanelAction()
}

// FuncHasOnSelectToSpeakPanelAction returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener".
func FuncHasOnSelectToSpeakPanelAction() (fn js.Func[func(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) bool]) {
	bindings.FuncHasOnSelectToSpeakPanelAction(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener" directly.
func HasOnSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret bool) {
	bindings.CallHasOnSelectToSpeakPanelAction(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectToSpeakPanelAction calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakPanelAction.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectToSpeakPanelAction(callback js.Func[func(action SelectToSpeakPanelAction, value float64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectToSpeakPanelAction(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSelectToSpeakStateChangeRequestedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnSelectToSpeakStateChangeRequestedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSelectToSpeakStateChangeRequestedEventCallbackFunc) DispatchCallback(
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

type OnSelectToSpeakStateChangeRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnSelectToSpeakStateChangeRequestedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSelectToSpeakStateChangeRequestedEventCallback[T]) DispatchCallback(
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

// HasFuncOnSelectToSpeakStateChangeRequested returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener" exists.
func HasFuncOnSelectToSpeakStateChangeRequested() bool {
	return js.True == bindings.HasFuncOnSelectToSpeakStateChangeRequested()
}

// FuncOnSelectToSpeakStateChangeRequested returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener".
func FuncOnSelectToSpeakStateChangeRequested() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnSelectToSpeakStateChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// OnSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener" directly.
func OnSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnSelectToSpeakStateChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSelectToSpeakStateChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSelectToSpeakStateChangeRequested returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener" exists.
func HasFuncOffSelectToSpeakStateChangeRequested() bool {
	return js.True == bindings.HasFuncOffSelectToSpeakStateChangeRequested()
}

// FuncOffSelectToSpeakStateChangeRequested returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener".
func FuncOffSelectToSpeakStateChangeRequested() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffSelectToSpeakStateChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// OffSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener" directly.
func OffSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffSelectToSpeakStateChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSelectToSpeakStateChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSelectToSpeakStateChangeRequested returns true if the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener" exists.
func HasFuncHasOnSelectToSpeakStateChangeRequested() bool {
	return js.True == bindings.HasFuncHasOnSelectToSpeakStateChangeRequested()
}

// FuncHasOnSelectToSpeakStateChangeRequested returns the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener".
func FuncHasOnSelectToSpeakStateChangeRequested() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnSelectToSpeakStateChangeRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener" directly.
func HasOnSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnSelectToSpeakStateChangeRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSelectToSpeakStateChangeRequested calls the function "WEBEXT.accessibilityPrivate.onSelectToSpeakStateChangeRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSelectToSpeakStateChangeRequested(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSelectToSpeakStateChangeRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnShowChromeVoxTutorialEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnShowChromeVoxTutorialEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnShowChromeVoxTutorialEventCallbackFunc) DispatchCallback(
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

type OnShowChromeVoxTutorialEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnShowChromeVoxTutorialEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnShowChromeVoxTutorialEventCallback[T]) DispatchCallback(
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

// HasFuncOnShowChromeVoxTutorial returns true if the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener" exists.
func HasFuncOnShowChromeVoxTutorial() bool {
	return js.True == bindings.HasFuncOnShowChromeVoxTutorial()
}

// FuncOnShowChromeVoxTutorial returns the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener".
func FuncOnShowChromeVoxTutorial() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnShowChromeVoxTutorial(
		js.Pointer(&fn),
	)
	return
}

// OnShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener" directly.
func OnShowChromeVoxTutorial(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnShowChromeVoxTutorial(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnShowChromeVoxTutorial(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnShowChromeVoxTutorial(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffShowChromeVoxTutorial returns true if the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener" exists.
func HasFuncOffShowChromeVoxTutorial() bool {
	return js.True == bindings.HasFuncOffShowChromeVoxTutorial()
}

// FuncOffShowChromeVoxTutorial returns the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener".
func FuncOffShowChromeVoxTutorial() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffShowChromeVoxTutorial(
		js.Pointer(&fn),
	)
	return
}

// OffShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener" directly.
func OffShowChromeVoxTutorial(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffShowChromeVoxTutorial(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffShowChromeVoxTutorial(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffShowChromeVoxTutorial(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnShowChromeVoxTutorial returns true if the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener" exists.
func HasFuncHasOnShowChromeVoxTutorial() bool {
	return js.True == bindings.HasFuncHasOnShowChromeVoxTutorial()
}

// FuncHasOnShowChromeVoxTutorial returns the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener".
func FuncHasOnShowChromeVoxTutorial() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnShowChromeVoxTutorial(
		js.Pointer(&fn),
	)
	return
}

// HasOnShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener" directly.
func HasOnShowChromeVoxTutorial(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnShowChromeVoxTutorial(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnShowChromeVoxTutorial calls the function "WEBEXT.accessibilityPrivate.onShowChromeVoxTutorial.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnShowChromeVoxTutorial(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnShowChromeVoxTutorial(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSwitchAccessCommandEventCallbackFunc func(this js.Ref, command SwitchAccessCommand) js.Ref

func (fn OnSwitchAccessCommandEventCallbackFunc) Register() js.Func[func(command SwitchAccessCommand)] {
	return js.RegisterCallback[func(command SwitchAccessCommand)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSwitchAccessCommandEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SwitchAccessCommand(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSwitchAccessCommandEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, command SwitchAccessCommand) js.Ref
	Arg T
}

func (cb *OnSwitchAccessCommandEventCallback[T]) Register() js.Func[func(command SwitchAccessCommand)] {
	return js.RegisterCallback[func(command SwitchAccessCommand)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSwitchAccessCommandEventCallback[T]) DispatchCallback(
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

		SwitchAccessCommand(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSwitchAccessCommand returns true if the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener" exists.
func HasFuncOnSwitchAccessCommand() bool {
	return js.True == bindings.HasFuncOnSwitchAccessCommand()
}

// FuncOnSwitchAccessCommand returns the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener".
func FuncOnSwitchAccessCommand() (fn js.Func[func(callback js.Func[func(command SwitchAccessCommand)])]) {
	bindings.FuncOnSwitchAccessCommand(
		js.Pointer(&fn),
	)
	return
}

// OnSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener" directly.
func OnSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret js.Void) {
	bindings.CallOnSwitchAccessCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSwitchAccessCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSwitchAccessCommand returns true if the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener" exists.
func HasFuncOffSwitchAccessCommand() bool {
	return js.True == bindings.HasFuncOffSwitchAccessCommand()
}

// FuncOffSwitchAccessCommand returns the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener".
func FuncOffSwitchAccessCommand() (fn js.Func[func(callback js.Func[func(command SwitchAccessCommand)])]) {
	bindings.FuncOffSwitchAccessCommand(
		js.Pointer(&fn),
	)
	return
}

// OffSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener" directly.
func OffSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret js.Void) {
	bindings.CallOffSwitchAccessCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSwitchAccessCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSwitchAccessCommand returns true if the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener" exists.
func HasFuncHasOnSwitchAccessCommand() bool {
	return js.True == bindings.HasFuncHasOnSwitchAccessCommand()
}

// FuncHasOnSwitchAccessCommand returns the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener".
func FuncHasOnSwitchAccessCommand() (fn js.Func[func(callback js.Func[func(command SwitchAccessCommand)]) bool]) {
	bindings.FuncHasOnSwitchAccessCommand(
		js.Pointer(&fn),
	)
	return
}

// HasOnSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener" directly.
func HasOnSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret bool) {
	bindings.CallHasOnSwitchAccessCommand(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSwitchAccessCommand calls the function "WEBEXT.accessibilityPrivate.onSwitchAccessCommand.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSwitchAccessCommand(callback js.Func[func(command SwitchAccessCommand)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSwitchAccessCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnToggleDictationEventCallbackFunc func(this js.Ref, activated bool) js.Ref

func (fn OnToggleDictationEventCallbackFunc) Register() js.Func[func(activated bool)] {
	return js.RegisterCallback[func(activated bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnToggleDictationEventCallbackFunc) DispatchCallback(
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

type OnToggleDictationEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, activated bool) js.Ref
	Arg T
}

func (cb *OnToggleDictationEventCallback[T]) Register() js.Func[func(activated bool)] {
	return js.RegisterCallback[func(activated bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnToggleDictationEventCallback[T]) DispatchCallback(
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

// HasFuncOnToggleDictation returns true if the function "WEBEXT.accessibilityPrivate.onToggleDictation.addListener" exists.
func HasFuncOnToggleDictation() bool {
	return js.True == bindings.HasFuncOnToggleDictation()
}

// FuncOnToggleDictation returns the function "WEBEXT.accessibilityPrivate.onToggleDictation.addListener".
func FuncOnToggleDictation() (fn js.Func[func(callback js.Func[func(activated bool)])]) {
	bindings.FuncOnToggleDictation(
		js.Pointer(&fn),
	)
	return
}

// OnToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.addListener" directly.
func OnToggleDictation(callback js.Func[func(activated bool)]) (ret js.Void) {
	bindings.CallOnToggleDictation(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnToggleDictation(callback js.Func[func(activated bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnToggleDictation(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffToggleDictation returns true if the function "WEBEXT.accessibilityPrivate.onToggleDictation.removeListener" exists.
func HasFuncOffToggleDictation() bool {
	return js.True == bindings.HasFuncOffToggleDictation()
}

// FuncOffToggleDictation returns the function "WEBEXT.accessibilityPrivate.onToggleDictation.removeListener".
func FuncOffToggleDictation() (fn js.Func[func(callback js.Func[func(activated bool)])]) {
	bindings.FuncOffToggleDictation(
		js.Pointer(&fn),
	)
	return
}

// OffToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.removeListener" directly.
func OffToggleDictation(callback js.Func[func(activated bool)]) (ret js.Void) {
	bindings.CallOffToggleDictation(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffToggleDictation(callback js.Func[func(activated bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffToggleDictation(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnToggleDictation returns true if the function "WEBEXT.accessibilityPrivate.onToggleDictation.hasListener" exists.
func HasFuncHasOnToggleDictation() bool {
	return js.True == bindings.HasFuncHasOnToggleDictation()
}

// FuncHasOnToggleDictation returns the function "WEBEXT.accessibilityPrivate.onToggleDictation.hasListener".
func FuncHasOnToggleDictation() (fn js.Func[func(callback js.Func[func(activated bool)]) bool]) {
	bindings.FuncHasOnToggleDictation(
		js.Pointer(&fn),
	)
	return
}

// HasOnToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.hasListener" directly.
func HasOnToggleDictation(callback js.Func[func(activated bool)]) (ret bool) {
	bindings.CallHasOnToggleDictation(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnToggleDictation calls the function "WEBEXT.accessibilityPrivate.onToggleDictation.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnToggleDictation(callback js.Func[func(activated bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnToggleDictation(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTwoFingerTouchStartEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnTwoFingerTouchStartEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTwoFingerTouchStartEventCallbackFunc) DispatchCallback(
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

type OnTwoFingerTouchStartEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnTwoFingerTouchStartEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTwoFingerTouchStartEventCallback[T]) DispatchCallback(
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

// HasFuncOnTwoFingerTouchStart returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener" exists.
func HasFuncOnTwoFingerTouchStart() bool {
	return js.True == bindings.HasFuncOnTwoFingerTouchStart()
}

// FuncOnTwoFingerTouchStart returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener".
func FuncOnTwoFingerTouchStart() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnTwoFingerTouchStart(
		js.Pointer(&fn),
	)
	return
}

// OnTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener" directly.
func OnTwoFingerTouchStart(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnTwoFingerTouchStart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTwoFingerTouchStart(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTwoFingerTouchStart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTwoFingerTouchStart returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener" exists.
func HasFuncOffTwoFingerTouchStart() bool {
	return js.True == bindings.HasFuncOffTwoFingerTouchStart()
}

// FuncOffTwoFingerTouchStart returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener".
func FuncOffTwoFingerTouchStart() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffTwoFingerTouchStart(
		js.Pointer(&fn),
	)
	return
}

// OffTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener" directly.
func OffTwoFingerTouchStart(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffTwoFingerTouchStart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTwoFingerTouchStart(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTwoFingerTouchStart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTwoFingerTouchStart returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener" exists.
func HasFuncHasOnTwoFingerTouchStart() bool {
	return js.True == bindings.HasFuncHasOnTwoFingerTouchStart()
}

// FuncHasOnTwoFingerTouchStart returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener".
func FuncHasOnTwoFingerTouchStart() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnTwoFingerTouchStart(
		js.Pointer(&fn),
	)
	return
}

// HasOnTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener" directly.
func HasOnTwoFingerTouchStart(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnTwoFingerTouchStart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTwoFingerTouchStart calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStart.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTwoFingerTouchStart(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTwoFingerTouchStart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTwoFingerTouchStopEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnTwoFingerTouchStopEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTwoFingerTouchStopEventCallbackFunc) DispatchCallback(
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

type OnTwoFingerTouchStopEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnTwoFingerTouchStopEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTwoFingerTouchStopEventCallback[T]) DispatchCallback(
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

// HasFuncOnTwoFingerTouchStop returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener" exists.
func HasFuncOnTwoFingerTouchStop() bool {
	return js.True == bindings.HasFuncOnTwoFingerTouchStop()
}

// FuncOnTwoFingerTouchStop returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener".
func FuncOnTwoFingerTouchStop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnTwoFingerTouchStop(
		js.Pointer(&fn),
	)
	return
}

// OnTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener" directly.
func OnTwoFingerTouchStop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnTwoFingerTouchStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTwoFingerTouchStop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTwoFingerTouchStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTwoFingerTouchStop returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener" exists.
func HasFuncOffTwoFingerTouchStop() bool {
	return js.True == bindings.HasFuncOffTwoFingerTouchStop()
}

// FuncOffTwoFingerTouchStop returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener".
func FuncOffTwoFingerTouchStop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffTwoFingerTouchStop(
		js.Pointer(&fn),
	)
	return
}

// OffTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener" directly.
func OffTwoFingerTouchStop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffTwoFingerTouchStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTwoFingerTouchStop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTwoFingerTouchStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTwoFingerTouchStop returns true if the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener" exists.
func HasFuncHasOnTwoFingerTouchStop() bool {
	return js.True == bindings.HasFuncHasOnTwoFingerTouchStop()
}

// FuncHasOnTwoFingerTouchStop returns the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener".
func FuncHasOnTwoFingerTouchStop() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnTwoFingerTouchStop(
		js.Pointer(&fn),
	)
	return
}

// HasOnTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener" directly.
func HasOnTwoFingerTouchStop(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnTwoFingerTouchStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTwoFingerTouchStop calls the function "WEBEXT.accessibilityPrivate.onTwoFingerTouchStop.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTwoFingerTouchStop(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTwoFingerTouchStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenSettingsSubpage returns true if the function "WEBEXT.accessibilityPrivate.openSettingsSubpage" exists.
func HasFuncOpenSettingsSubpage() bool {
	return js.True == bindings.HasFuncOpenSettingsSubpage()
}

// FuncOpenSettingsSubpage returns the function "WEBEXT.accessibilityPrivate.openSettingsSubpage".
func FuncOpenSettingsSubpage() (fn js.Func[func(subpage js.String)]) {
	bindings.FuncOpenSettingsSubpage(
		js.Pointer(&fn),
	)
	return
}

// OpenSettingsSubpage calls the function "WEBEXT.accessibilityPrivate.openSettingsSubpage" directly.
func OpenSettingsSubpage(subpage js.String) (ret js.Void) {
	bindings.CallOpenSettingsSubpage(
		js.Pointer(&ret),
		subpage.Ref(),
	)

	return
}

// TryOpenSettingsSubpage calls the function "WEBEXT.accessibilityPrivate.openSettingsSubpage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSettingsSubpage(subpage js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSettingsSubpage(
		js.Pointer(&ret), js.Pointer(&exception),
		subpage.Ref(),
	)

	return
}

// HasFuncPerformAcceleratorAction returns true if the function "WEBEXT.accessibilityPrivate.performAcceleratorAction" exists.
func HasFuncPerformAcceleratorAction() bool {
	return js.True == bindings.HasFuncPerformAcceleratorAction()
}

// FuncPerformAcceleratorAction returns the function "WEBEXT.accessibilityPrivate.performAcceleratorAction".
func FuncPerformAcceleratorAction() (fn js.Func[func(acceleratorAction AcceleratorAction)]) {
	bindings.FuncPerformAcceleratorAction(
		js.Pointer(&fn),
	)
	return
}

// PerformAcceleratorAction calls the function "WEBEXT.accessibilityPrivate.performAcceleratorAction" directly.
func PerformAcceleratorAction(acceleratorAction AcceleratorAction) (ret js.Void) {
	bindings.CallPerformAcceleratorAction(
		js.Pointer(&ret),
		uint32(acceleratorAction),
	)

	return
}

// TryPerformAcceleratorAction calls the function "WEBEXT.accessibilityPrivate.performAcceleratorAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPerformAcceleratorAction(acceleratorAction AcceleratorAction) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformAcceleratorAction(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(acceleratorAction),
	)

	return
}

// HasFuncSendSyntheticKeyEvent returns true if the function "WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent" exists.
func HasFuncSendSyntheticKeyEvent() bool {
	return js.True == bindings.HasFuncSendSyntheticKeyEvent()
}

// FuncSendSyntheticKeyEvent returns the function "WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent".
func FuncSendSyntheticKeyEvent() (fn js.Func[func(keyEvent SyntheticKeyboardEvent, useRewriters bool)]) {
	bindings.FuncSendSyntheticKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// SendSyntheticKeyEvent calls the function "WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent" directly.
func SendSyntheticKeyEvent(keyEvent SyntheticKeyboardEvent, useRewriters bool) (ret js.Void) {
	bindings.CallSendSyntheticKeyEvent(
		js.Pointer(&ret),
		js.Pointer(&keyEvent),
		js.Bool(bool(useRewriters)),
	)

	return
}

// TrySendSyntheticKeyEvent calls the function "WEBEXT.accessibilityPrivate.sendSyntheticKeyEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendSyntheticKeyEvent(keyEvent SyntheticKeyboardEvent, useRewriters bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendSyntheticKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&keyEvent),
		js.Bool(bool(useRewriters)),
	)

	return
}

// HasFuncSendSyntheticMouseEvent returns true if the function "WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent" exists.
func HasFuncSendSyntheticMouseEvent() bool {
	return js.True == bindings.HasFuncSendSyntheticMouseEvent()
}

// FuncSendSyntheticMouseEvent returns the function "WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent".
func FuncSendSyntheticMouseEvent() (fn js.Func[func(mouseEvent SyntheticMouseEvent)]) {
	bindings.FuncSendSyntheticMouseEvent(
		js.Pointer(&fn),
	)
	return
}

// SendSyntheticMouseEvent calls the function "WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent" directly.
func SendSyntheticMouseEvent(mouseEvent SyntheticMouseEvent) (ret js.Void) {
	bindings.CallSendSyntheticMouseEvent(
		js.Pointer(&ret),
		js.Pointer(&mouseEvent),
	)

	return
}

// TrySendSyntheticMouseEvent calls the function "WEBEXT.accessibilityPrivate.sendSyntheticMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendSyntheticMouseEvent(mouseEvent SyntheticMouseEvent) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendSyntheticMouseEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&mouseEvent),
	)

	return
}

// HasFuncSetFocusRings returns true if the function "WEBEXT.accessibilityPrivate.setFocusRings" exists.
func HasFuncSetFocusRings() bool {
	return js.True == bindings.HasFuncSetFocusRings()
}

// FuncSetFocusRings returns the function "WEBEXT.accessibilityPrivate.setFocusRings".
func FuncSetFocusRings() (fn js.Func[func(focusRings js.Array[FocusRingInfo], atType AssistiveTechnologyType)]) {
	bindings.FuncSetFocusRings(
		js.Pointer(&fn),
	)
	return
}

// SetFocusRings calls the function "WEBEXT.accessibilityPrivate.setFocusRings" directly.
func SetFocusRings(focusRings js.Array[FocusRingInfo], atType AssistiveTechnologyType) (ret js.Void) {
	bindings.CallSetFocusRings(
		js.Pointer(&ret),
		focusRings.Ref(),
		uint32(atType),
	)

	return
}

// TrySetFocusRings calls the function "WEBEXT.accessibilityPrivate.setFocusRings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetFocusRings(focusRings js.Array[FocusRingInfo], atType AssistiveTechnologyType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetFocusRings(
		js.Pointer(&ret), js.Pointer(&exception),
		focusRings.Ref(),
		uint32(atType),
	)

	return
}

// HasFuncSetHighlights returns true if the function "WEBEXT.accessibilityPrivate.setHighlights" exists.
func HasFuncSetHighlights() bool {
	return js.True == bindings.HasFuncSetHighlights()
}

// FuncSetHighlights returns the function "WEBEXT.accessibilityPrivate.setHighlights".
func FuncSetHighlights() (fn js.Func[func(rects js.Array[ScreenRect], color js.String)]) {
	bindings.FuncSetHighlights(
		js.Pointer(&fn),
	)
	return
}

// SetHighlights calls the function "WEBEXT.accessibilityPrivate.setHighlights" directly.
func SetHighlights(rects js.Array[ScreenRect], color js.String) (ret js.Void) {
	bindings.CallSetHighlights(
		js.Pointer(&ret),
		rects.Ref(),
		color.Ref(),
	)

	return
}

// TrySetHighlights calls the function "WEBEXT.accessibilityPrivate.setHighlights"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetHighlights(rects js.Array[ScreenRect], color js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetHighlights(
		js.Pointer(&ret), js.Pointer(&exception),
		rects.Ref(),
		color.Ref(),
	)

	return
}

// HasFuncSetKeyboardListener returns true if the function "WEBEXT.accessibilityPrivate.setKeyboardListener" exists.
func HasFuncSetKeyboardListener() bool {
	return js.True == bindings.HasFuncSetKeyboardListener()
}

// FuncSetKeyboardListener returns the function "WEBEXT.accessibilityPrivate.setKeyboardListener".
func FuncSetKeyboardListener() (fn js.Func[func(enabled bool, capture bool)]) {
	bindings.FuncSetKeyboardListener(
		js.Pointer(&fn),
	)
	return
}

// SetKeyboardListener calls the function "WEBEXT.accessibilityPrivate.setKeyboardListener" directly.
func SetKeyboardListener(enabled bool, capture bool) (ret js.Void) {
	bindings.CallSetKeyboardListener(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
		js.Bool(bool(capture)),
	)

	return
}

// TrySetKeyboardListener calls the function "WEBEXT.accessibilityPrivate.setKeyboardListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetKeyboardListener(enabled bool, capture bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetKeyboardListener(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
		js.Bool(bool(capture)),
	)

	return
}

// HasFuncSetNativeAccessibilityEnabled returns true if the function "WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled" exists.
func HasFuncSetNativeAccessibilityEnabled() bool {
	return js.True == bindings.HasFuncSetNativeAccessibilityEnabled()
}

// FuncSetNativeAccessibilityEnabled returns the function "WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled".
func FuncSetNativeAccessibilityEnabled() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetNativeAccessibilityEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetNativeAccessibilityEnabled calls the function "WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled" directly.
func SetNativeAccessibilityEnabled(enabled bool) (ret js.Void) {
	bindings.CallSetNativeAccessibilityEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetNativeAccessibilityEnabled calls the function "WEBEXT.accessibilityPrivate.setNativeAccessibilityEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetNativeAccessibilityEnabled(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetNativeAccessibilityEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetNativeChromeVoxArcSupportForCurrentApp returns true if the function "WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp" exists.
func HasFuncSetNativeChromeVoxArcSupportForCurrentApp() bool {
	return js.True == bindings.HasFuncSetNativeChromeVoxArcSupportForCurrentApp()
}

// FuncSetNativeChromeVoxArcSupportForCurrentApp returns the function "WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp".
func FuncSetNativeChromeVoxArcSupportForCurrentApp() (fn js.Func[func(enabled bool) js.Promise[SetNativeChromeVoxResponse]]) {
	bindings.FuncSetNativeChromeVoxArcSupportForCurrentApp(
		js.Pointer(&fn),
	)
	return
}

// SetNativeChromeVoxArcSupportForCurrentApp calls the function "WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp" directly.
func SetNativeChromeVoxArcSupportForCurrentApp(enabled bool) (ret js.Promise[SetNativeChromeVoxResponse]) {
	bindings.CallSetNativeChromeVoxArcSupportForCurrentApp(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetNativeChromeVoxArcSupportForCurrentApp calls the function "WEBEXT.accessibilityPrivate.setNativeChromeVoxArcSupportForCurrentApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetNativeChromeVoxArcSupportForCurrentApp(enabled bool) (ret js.Promise[SetNativeChromeVoxResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetNativeChromeVoxArcSupportForCurrentApp(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetPointScanState returns true if the function "WEBEXT.accessibilityPrivate.setPointScanState" exists.
func HasFuncSetPointScanState() bool {
	return js.True == bindings.HasFuncSetPointScanState()
}

// FuncSetPointScanState returns the function "WEBEXT.accessibilityPrivate.setPointScanState".
func FuncSetPointScanState() (fn js.Func[func(state PointScanState)]) {
	bindings.FuncSetPointScanState(
		js.Pointer(&fn),
	)
	return
}

// SetPointScanState calls the function "WEBEXT.accessibilityPrivate.setPointScanState" directly.
func SetPointScanState(state PointScanState) (ret js.Void) {
	bindings.CallSetPointScanState(
		js.Pointer(&ret),
		uint32(state),
	)

	return
}

// TrySetPointScanState calls the function "WEBEXT.accessibilityPrivate.setPointScanState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPointScanState(state PointScanState) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPointScanState(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(state),
	)

	return
}

// HasFuncSetSelectToSpeakState returns true if the function "WEBEXT.accessibilityPrivate.setSelectToSpeakState" exists.
func HasFuncSetSelectToSpeakState() bool {
	return js.True == bindings.HasFuncSetSelectToSpeakState()
}

// FuncSetSelectToSpeakState returns the function "WEBEXT.accessibilityPrivate.setSelectToSpeakState".
func FuncSetSelectToSpeakState() (fn js.Func[func(state SelectToSpeakState)]) {
	bindings.FuncSetSelectToSpeakState(
		js.Pointer(&fn),
	)
	return
}

// SetSelectToSpeakState calls the function "WEBEXT.accessibilityPrivate.setSelectToSpeakState" directly.
func SetSelectToSpeakState(state SelectToSpeakState) (ret js.Void) {
	bindings.CallSetSelectToSpeakState(
		js.Pointer(&ret),
		uint32(state),
	)

	return
}

// TrySetSelectToSpeakState calls the function "WEBEXT.accessibilityPrivate.setSelectToSpeakState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSelectToSpeakState(state SelectToSpeakState) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSelectToSpeakState(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(state),
	)

	return
}

// HasFuncSetVirtualKeyboardVisible returns true if the function "WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible" exists.
func HasFuncSetVirtualKeyboardVisible() bool {
	return js.True == bindings.HasFuncSetVirtualKeyboardVisible()
}

// FuncSetVirtualKeyboardVisible returns the function "WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible".
func FuncSetVirtualKeyboardVisible() (fn js.Func[func(isVisible bool)]) {
	bindings.FuncSetVirtualKeyboardVisible(
		js.Pointer(&fn),
	)
	return
}

// SetVirtualKeyboardVisible calls the function "WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible" directly.
func SetVirtualKeyboardVisible(isVisible bool) (ret js.Void) {
	bindings.CallSetVirtualKeyboardVisible(
		js.Pointer(&ret),
		js.Bool(bool(isVisible)),
	)

	return
}

// TrySetVirtualKeyboardVisible calls the function "WEBEXT.accessibilityPrivate.setVirtualKeyboardVisible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetVirtualKeyboardVisible(isVisible bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetVirtualKeyboardVisible(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(isVisible)),
	)

	return
}

// HasFuncShowConfirmationDialog returns true if the function "WEBEXT.accessibilityPrivate.showConfirmationDialog" exists.
func HasFuncShowConfirmationDialog() bool {
	return js.True == bindings.HasFuncShowConfirmationDialog()
}

// FuncShowConfirmationDialog returns the function "WEBEXT.accessibilityPrivate.showConfirmationDialog".
func FuncShowConfirmationDialog() (fn js.Func[func(title js.String, description js.String, cancelName js.String) js.Promise[js.Boolean]]) {
	bindings.FuncShowConfirmationDialog(
		js.Pointer(&fn),
	)
	return
}

// ShowConfirmationDialog calls the function "WEBEXT.accessibilityPrivate.showConfirmationDialog" directly.
func ShowConfirmationDialog(title js.String, description js.String, cancelName js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallShowConfirmationDialog(
		js.Pointer(&ret),
		title.Ref(),
		description.Ref(),
		cancelName.Ref(),
	)

	return
}

// TryShowConfirmationDialog calls the function "WEBEXT.accessibilityPrivate.showConfirmationDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowConfirmationDialog(title js.String, description js.String, cancelName js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowConfirmationDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		title.Ref(),
		description.Ref(),
		cancelName.Ref(),
	)

	return
}

// HasFuncShowToast returns true if the function "WEBEXT.accessibilityPrivate.showToast" exists.
func HasFuncShowToast() bool {
	return js.True == bindings.HasFuncShowToast()
}

// FuncShowToast returns the function "WEBEXT.accessibilityPrivate.showToast".
func FuncShowToast() (fn js.Func[func(typ ToastType)]) {
	bindings.FuncShowToast(
		js.Pointer(&fn),
	)
	return
}

// ShowToast calls the function "WEBEXT.accessibilityPrivate.showToast" directly.
func ShowToast(typ ToastType) (ret js.Void) {
	bindings.CallShowToast(
		js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryShowToast calls the function "WEBEXT.accessibilityPrivate.showToast"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowToast(typ ToastType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowToast(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncSilenceSpokenFeedback returns true if the function "WEBEXT.accessibilityPrivate.silenceSpokenFeedback" exists.
func HasFuncSilenceSpokenFeedback() bool {
	return js.True == bindings.HasFuncSilenceSpokenFeedback()
}

// FuncSilenceSpokenFeedback returns the function "WEBEXT.accessibilityPrivate.silenceSpokenFeedback".
func FuncSilenceSpokenFeedback() (fn js.Func[func()]) {
	bindings.FuncSilenceSpokenFeedback(
		js.Pointer(&fn),
	)
	return
}

// SilenceSpokenFeedback calls the function "WEBEXT.accessibilityPrivate.silenceSpokenFeedback" directly.
func SilenceSpokenFeedback() (ret js.Void) {
	bindings.CallSilenceSpokenFeedback(
		js.Pointer(&ret),
	)

	return
}

// TrySilenceSpokenFeedback calls the function "WEBEXT.accessibilityPrivate.silenceSpokenFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySilenceSpokenFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySilenceSpokenFeedback(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToggleDictation returns true if the function "WEBEXT.accessibilityPrivate.toggleDictation" exists.
func HasFuncToggleDictation() bool {
	return js.True == bindings.HasFuncToggleDictation()
}

// FuncToggleDictation returns the function "WEBEXT.accessibilityPrivate.toggleDictation".
func FuncToggleDictation() (fn js.Func[func()]) {
	bindings.FuncToggleDictation(
		js.Pointer(&fn),
	)
	return
}

// ToggleDictation calls the function "WEBEXT.accessibilityPrivate.toggleDictation" directly.
func ToggleDictation() (ret js.Void) {
	bindings.CallToggleDictation(
		js.Pointer(&ret),
	)

	return
}

// TryToggleDictation calls the function "WEBEXT.accessibilityPrivate.toggleDictation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryToggleDictation() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryToggleDictation(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUpdateDictationBubble returns true if the function "WEBEXT.accessibilityPrivate.updateDictationBubble" exists.
func HasFuncUpdateDictationBubble() bool {
	return js.True == bindings.HasFuncUpdateDictationBubble()
}

// FuncUpdateDictationBubble returns the function "WEBEXT.accessibilityPrivate.updateDictationBubble".
func FuncUpdateDictationBubble() (fn js.Func[func(properties DictationBubbleProperties)]) {
	bindings.FuncUpdateDictationBubble(
		js.Pointer(&fn),
	)
	return
}

// UpdateDictationBubble calls the function "WEBEXT.accessibilityPrivate.updateDictationBubble" directly.
func UpdateDictationBubble(properties DictationBubbleProperties) (ret js.Void) {
	bindings.CallUpdateDictationBubble(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryUpdateDictationBubble calls the function "WEBEXT.accessibilityPrivate.updateDictationBubble"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateDictationBubble(properties DictationBubbleProperties) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateDictationBubble(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncUpdateSelectToSpeakPanel returns true if the function "WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel" exists.
func HasFuncUpdateSelectToSpeakPanel() bool {
	return js.True == bindings.HasFuncUpdateSelectToSpeakPanel()
}

// FuncUpdateSelectToSpeakPanel returns the function "WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel".
func FuncUpdateSelectToSpeakPanel() (fn js.Func[func(show bool, anchor ScreenRect, isPaused bool, speed float64)]) {
	bindings.FuncUpdateSelectToSpeakPanel(
		js.Pointer(&fn),
	)
	return
}

// UpdateSelectToSpeakPanel calls the function "WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel" directly.
func UpdateSelectToSpeakPanel(show bool, anchor ScreenRect, isPaused bool, speed float64) (ret js.Void) {
	bindings.CallUpdateSelectToSpeakPanel(
		js.Pointer(&ret),
		js.Bool(bool(show)),
		js.Pointer(&anchor),
		js.Bool(bool(isPaused)),
		float64(speed),
	)

	return
}

// TryUpdateSelectToSpeakPanel calls the function "WEBEXT.accessibilityPrivate.updateSelectToSpeakPanel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateSelectToSpeakPanel(show bool, anchor ScreenRect, isPaused bool, speed float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateSelectToSpeakPanel(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(show)),
		js.Pointer(&anchor),
		js.Bool(bool(isPaused)),
		float64(speed),
	)

	return
}

// HasFuncUpdateSwitchAccessBubble returns true if the function "WEBEXT.accessibilityPrivate.updateSwitchAccessBubble" exists.
func HasFuncUpdateSwitchAccessBubble() bool {
	return js.True == bindings.HasFuncUpdateSwitchAccessBubble()
}

// FuncUpdateSwitchAccessBubble returns the function "WEBEXT.accessibilityPrivate.updateSwitchAccessBubble".
func FuncUpdateSwitchAccessBubble() (fn js.Func[func(bubble SwitchAccessBubble, show bool, anchor ScreenRect, actions js.Array[SwitchAccessMenuAction])]) {
	bindings.FuncUpdateSwitchAccessBubble(
		js.Pointer(&fn),
	)
	return
}

// UpdateSwitchAccessBubble calls the function "WEBEXT.accessibilityPrivate.updateSwitchAccessBubble" directly.
func UpdateSwitchAccessBubble(bubble SwitchAccessBubble, show bool, anchor ScreenRect, actions js.Array[SwitchAccessMenuAction]) (ret js.Void) {
	bindings.CallUpdateSwitchAccessBubble(
		js.Pointer(&ret),
		uint32(bubble),
		js.Bool(bool(show)),
		js.Pointer(&anchor),
		actions.Ref(),
	)

	return
}

// TryUpdateSwitchAccessBubble calls the function "WEBEXT.accessibilityPrivate.updateSwitchAccessBubble"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateSwitchAccessBubble(bubble SwitchAccessBubble, show bool, anchor ScreenRect, actions js.Array[SwitchAccessMenuAction]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateSwitchAccessBubble(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(bubble),
		js.Bool(bool(show)),
		js.Pointer(&anchor),
		actions.Ref(),
	)

	return
}

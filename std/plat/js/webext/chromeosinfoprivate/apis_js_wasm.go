// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package chromeosinfoprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chromeosinfoprivate/bindings"
)

type AssistantStatus uint32

const (
	_ AssistantStatus = iota

	AssistantStatus_UNSUPPORTED
	AssistantStatus_SUPPORTED
)

func (AssistantStatus) FromRef(str js.Ref) AssistantStatus {
	return AssistantStatus(bindings.ConstOfAssistantStatus(str))
}

func (x AssistantStatus) String() (string, bool) {
	switch x {
	case AssistantStatus_UNSUPPORTED:
		return "unsupported", true
	case AssistantStatus_SUPPORTED:
		return "supported", true
	default:
		return "", false
	}
}

type DeviceType uint32

const (
	_ DeviceType = iota

	DeviceType_CHROMEBASE
	DeviceType_CHROMEBIT
	DeviceType_CHROMEBOOK
	DeviceType_CHROMEBOX
	DeviceType_CHROMEDEVICE
)

func (DeviceType) FromRef(str js.Ref) DeviceType {
	return DeviceType(bindings.ConstOfDeviceType(str))
}

func (x DeviceType) String() (string, bool) {
	switch x {
	case DeviceType_CHROMEBASE:
		return "chromebase", true
	case DeviceType_CHROMEBIT:
		return "chromebit", true
	case DeviceType_CHROMEBOOK:
		return "chromebook", true
	case DeviceType_CHROMEBOX:
		return "chromebox", true
	case DeviceType_CHROMEDEVICE:
		return "chromedevice", true
	default:
		return "", false
	}
}

type ManagedDeviceStatus uint32

const (
	_ ManagedDeviceStatus = iota

	ManagedDeviceStatus_MANAGED
	ManagedDeviceStatus_NOT_MANAGED
)

func (ManagedDeviceStatus) FromRef(str js.Ref) ManagedDeviceStatus {
	return ManagedDeviceStatus(bindings.ConstOfManagedDeviceStatus(str))
}

func (x ManagedDeviceStatus) String() (string, bool) {
	switch x {
	case ManagedDeviceStatus_MANAGED:
		return "managed", true
	case ManagedDeviceStatus_NOT_MANAGED:
		return "not managed", true
	default:
		return "", false
	}
}

type PlayStoreStatus uint32

const (
	_ PlayStoreStatus = iota

	PlayStoreStatus_NOT_AVAILABLE
	PlayStoreStatus_AVAILABLE
	PlayStoreStatus_ENABLED
)

func (PlayStoreStatus) FromRef(str js.Ref) PlayStoreStatus {
	return PlayStoreStatus(bindings.ConstOfPlayStoreStatus(str))
}

func (x PlayStoreStatus) String() (string, bool) {
	switch x {
	case PlayStoreStatus_NOT_AVAILABLE:
		return "not available", true
	case PlayStoreStatus_AVAILABLE:
		return "available", true
	case PlayStoreStatus_ENABLED:
		return "enabled", true
	default:
		return "", false
	}
}

type SessionType uint32

const (
	_ SessionType = iota

	SessionType_NORMAL
	SessionType_KIOSK
	SessionType_PUBLIC_SESSION
)

func (SessionType) FromRef(str js.Ref) SessionType {
	return SessionType(bindings.ConstOfSessionType(str))
}

func (x SessionType) String() (string, bool) {
	switch x {
	case SessionType_NORMAL:
		return "normal", true
	case SessionType_KIOSK:
		return "kiosk", true
	case SessionType_PUBLIC_SESSION:
		return "public session", true
	default:
		return "", false
	}
}

type StylusStatus uint32

const (
	_ StylusStatus = iota

	StylusStatus_UNSUPPORTED
	StylusStatus_SUPPORTED
	StylusStatus_SEEN
)

func (StylusStatus) FromRef(str js.Ref) StylusStatus {
	return StylusStatus(bindings.ConstOfStylusStatus(str))
}

func (x StylusStatus) String() (string, bool) {
	switch x {
	case StylusStatus_UNSUPPORTED:
		return "unsupported", true
	case StylusStatus_SUPPORTED:
		return "supported", true
	case StylusStatus_SEEN:
		return "seen", true
	default:
		return "", false
	}
}

type GetReturnType struct {
	// A11yAutoClickEnabled is "GetReturnType.a11yAutoClickEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yAutoClickEnabled MUST be set to true to make this field effective.
	A11yAutoClickEnabled bool
	// A11yCaretHighlightEnabled is "GetReturnType.a11yCaretHighlightEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yCaretHighlightEnabled MUST be set to true to make this field effective.
	A11yCaretHighlightEnabled bool
	// A11yCursorColorEnabled is "GetReturnType.a11yCursorColorEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yCursorColorEnabled MUST be set to true to make this field effective.
	A11yCursorColorEnabled bool
	// A11yCursorHighlightEnabled is "GetReturnType.a11yCursorHighlightEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yCursorHighlightEnabled MUST be set to true to make this field effective.
	A11yCursorHighlightEnabled bool
	// A11yDockedMagnifierEnabled is "GetReturnType.a11yDockedMagnifierEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yDockedMagnifierEnabled MUST be set to true to make this field effective.
	A11yDockedMagnifierEnabled bool
	// A11yFocusHighlightEnabled is "GetReturnType.a11yFocusHighlightEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yFocusHighlightEnabled MUST be set to true to make this field effective.
	A11yFocusHighlightEnabled bool
	// A11yHighContrastEnabled is "GetReturnType.a11yHighContrastEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yHighContrastEnabled MUST be set to true to make this field effective.
	A11yHighContrastEnabled bool
	// A11yLargeCursorEnabled is "GetReturnType.a11yLargeCursorEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yLargeCursorEnabled MUST be set to true to make this field effective.
	A11yLargeCursorEnabled bool
	// A11yScreenMagnifierEnabled is "GetReturnType.a11yScreenMagnifierEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yScreenMagnifierEnabled MUST be set to true to make this field effective.
	A11yScreenMagnifierEnabled bool
	// A11ySelectToSpeakEnabled is "GetReturnType.a11ySelectToSpeakEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11ySelectToSpeakEnabled MUST be set to true to make this field effective.
	A11ySelectToSpeakEnabled bool
	// A11ySpokenFeedbackEnabled is "GetReturnType.a11ySpokenFeedbackEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11ySpokenFeedbackEnabled MUST be set to true to make this field effective.
	A11ySpokenFeedbackEnabled bool
	// A11yStickyKeysEnabled is "GetReturnType.a11yStickyKeysEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yStickyKeysEnabled MUST be set to true to make this field effective.
	A11yStickyKeysEnabled bool
	// A11ySwitchAccessEnabled is "GetReturnType.a11ySwitchAccessEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11ySwitchAccessEnabled MUST be set to true to make this field effective.
	A11ySwitchAccessEnabled bool
	// A11yVirtualKeyboardEnabled is "GetReturnType.a11yVirtualKeyboardEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_A11yVirtualKeyboardEnabled MUST be set to true to make this field effective.
	A11yVirtualKeyboardEnabled bool
	// AssistantStatus is "GetReturnType.assistantStatus"
	//
	// Optional
	AssistantStatus AssistantStatus
	// Board is "GetReturnType.board"
	//
	// Optional
	Board js.String
	// ClientId is "GetReturnType.clientId"
	//
	// Optional
	ClientId js.String
	// CustomizationId is "GetReturnType.customizationId"
	//
	// Optional
	CustomizationId js.String
	// DeviceType is "GetReturnType.deviceType"
	//
	// Optional
	DeviceType DeviceType
	// HomeProvider is "GetReturnType.homeProvider"
	//
	// Optional
	HomeProvider js.String
	// Hwid is "GetReturnType.hwid"
	//
	// Optional
	Hwid js.String
	// InitialLocale is "GetReturnType.initialLocale"
	//
	// Optional
	InitialLocale js.String
	// IsMeetDevice is "GetReturnType.isMeetDevice"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMeetDevice MUST be set to true to make this field effective.
	IsMeetDevice bool
	// IsOwner is "GetReturnType.isOwner"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsOwner MUST be set to true to make this field effective.
	IsOwner bool
	// ManagedDeviceStatus is "GetReturnType.managedDeviceStatus"
	//
	// Optional
	ManagedDeviceStatus ManagedDeviceStatus
	// PlayStoreStatus is "GetReturnType.playStoreStatus"
	//
	// Optional
	PlayStoreStatus PlayStoreStatus
	// SendFunctionKeys is "GetReturnType.sendFunctionKeys"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendFunctionKeys MUST be set to true to make this field effective.
	SendFunctionKeys bool
	// SessionType is "GetReturnType.sessionType"
	//
	// Optional
	SessionType SessionType
	// StylusStatus is "GetReturnType.stylusStatus"
	//
	// Optional
	StylusStatus StylusStatus
	// SupportedTimezones is "GetReturnType.supportedTimezones"
	//
	// Optional
	SupportedTimezones js.Array[js.Array[js.String]]
	// Timezone is "GetReturnType.timezone"
	//
	// Optional
	Timezone js.String

	FFI_USE_A11yAutoClickEnabled       bool // for A11yAutoClickEnabled.
	FFI_USE_A11yCaretHighlightEnabled  bool // for A11yCaretHighlightEnabled.
	FFI_USE_A11yCursorColorEnabled     bool // for A11yCursorColorEnabled.
	FFI_USE_A11yCursorHighlightEnabled bool // for A11yCursorHighlightEnabled.
	FFI_USE_A11yDockedMagnifierEnabled bool // for A11yDockedMagnifierEnabled.
	FFI_USE_A11yFocusHighlightEnabled  bool // for A11yFocusHighlightEnabled.
	FFI_USE_A11yHighContrastEnabled    bool // for A11yHighContrastEnabled.
	FFI_USE_A11yLargeCursorEnabled     bool // for A11yLargeCursorEnabled.
	FFI_USE_A11yScreenMagnifierEnabled bool // for A11yScreenMagnifierEnabled.
	FFI_USE_A11ySelectToSpeakEnabled   bool // for A11ySelectToSpeakEnabled.
	FFI_USE_A11ySpokenFeedbackEnabled  bool // for A11ySpokenFeedbackEnabled.
	FFI_USE_A11yStickyKeysEnabled      bool // for A11yStickyKeysEnabled.
	FFI_USE_A11ySwitchAccessEnabled    bool // for A11ySwitchAccessEnabled.
	FFI_USE_A11yVirtualKeyboardEnabled bool // for A11yVirtualKeyboardEnabled.
	FFI_USE_IsMeetDevice               bool // for IsMeetDevice.
	FFI_USE_IsOwner                    bool // for IsOwner.
	FFI_USE_SendFunctionKeys           bool // for SendFunctionKeys.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetReturnType with all fields set.
func (p GetReturnType) FromRef(ref js.Ref) GetReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetReturnType in the application heap.
func (p GetReturnType) New() js.Ref {
	return bindings.GetReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetReturnType) Update(ref js.Ref) {
	bindings.GetReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Board.Ref(),
		p.ClientId.Ref(),
		p.CustomizationId.Ref(),
		p.HomeProvider.Ref(),
		p.Hwid.Ref(),
		p.InitialLocale.Ref(),
		p.SupportedTimezones.Ref(),
		p.Timezone.Ref(),
	)
	p.Board = p.Board.FromRef(js.Undefined)
	p.ClientId = p.ClientId.FromRef(js.Undefined)
	p.CustomizationId = p.CustomizationId.FromRef(js.Undefined)
	p.HomeProvider = p.HomeProvider.FromRef(js.Undefined)
	p.Hwid = p.Hwid.FromRef(js.Undefined)
	p.InitialLocale = p.InitialLocale.FromRef(js.Undefined)
	p.SupportedTimezones = p.SupportedTimezones.FromRef(js.Undefined)
	p.Timezone = p.Timezone.FromRef(js.Undefined)
}

type PropertyName uint32

const (
	_ PropertyName = iota

	PropertyName_TIMEZONE
	PropertyName_A_11Y_LARGE_CURSOR_ENABLED
	PropertyName_A_11Y_STICKY_KEYS_ENABLED
	PropertyName_A_11Y_SPOKEN_FEEDBACK_ENABLED
	PropertyName_A_11Y_HIGH_CONTRAST_ENABLED
	PropertyName_A_11Y_SCREEN_MAGNIFIER_ENABLED
	PropertyName_A_11Y_AUTO_CLICK_ENABLED
	PropertyName_A_11Y_VIRTUAL_KEYBOARD_ENABLED
	PropertyName_A_11Y_CARET_HIGHLIGHT_ENABLED
	PropertyName_A_11Y_CURSOR_HIGHLIGHT_ENABLED
	PropertyName_A_11Y_FOCUS_HIGHLIGHT_ENABLED
	PropertyName_A_11Y_SELECT_TO_SPEAK_ENABLED
	PropertyName_A_11Y_SWITCH_ACCESS_ENABLED
	PropertyName_A_11Y_CURSOR_COLOR_ENABLED
	PropertyName_A_11Y_DOCKED_MAGNIFIER_ENABLED
	PropertyName_SEND_FUNCTION_KEYS
)

func (PropertyName) FromRef(str js.Ref) PropertyName {
	return PropertyName(bindings.ConstOfPropertyName(str))
}

func (x PropertyName) String() (string, bool) {
	switch x {
	case PropertyName_TIMEZONE:
		return "timezone", true
	case PropertyName_A_11Y_LARGE_CURSOR_ENABLED:
		return "a11yLargeCursorEnabled", true
	case PropertyName_A_11Y_STICKY_KEYS_ENABLED:
		return "a11yStickyKeysEnabled", true
	case PropertyName_A_11Y_SPOKEN_FEEDBACK_ENABLED:
		return "a11ySpokenFeedbackEnabled", true
	case PropertyName_A_11Y_HIGH_CONTRAST_ENABLED:
		return "a11yHighContrastEnabled", true
	case PropertyName_A_11Y_SCREEN_MAGNIFIER_ENABLED:
		return "a11yScreenMagnifierEnabled", true
	case PropertyName_A_11Y_AUTO_CLICK_ENABLED:
		return "a11yAutoClickEnabled", true
	case PropertyName_A_11Y_VIRTUAL_KEYBOARD_ENABLED:
		return "a11yVirtualKeyboardEnabled", true
	case PropertyName_A_11Y_CARET_HIGHLIGHT_ENABLED:
		return "a11yCaretHighlightEnabled", true
	case PropertyName_A_11Y_CURSOR_HIGHLIGHT_ENABLED:
		return "a11yCursorHighlightEnabled", true
	case PropertyName_A_11Y_FOCUS_HIGHLIGHT_ENABLED:
		return "a11yFocusHighlightEnabled", true
	case PropertyName_A_11Y_SELECT_TO_SPEAK_ENABLED:
		return "a11ySelectToSpeakEnabled", true
	case PropertyName_A_11Y_SWITCH_ACCESS_ENABLED:
		return "a11ySwitchAccessEnabled", true
	case PropertyName_A_11Y_CURSOR_COLOR_ENABLED:
		return "a11yCursorColorEnabled", true
	case PropertyName_A_11Y_DOCKED_MAGNIFIER_ENABLED:
		return "a11yDockedMagnifierEnabled", true
	case PropertyName_SEND_FUNCTION_KEYS:
		return "sendFunctionKeys", true
	default:
		return "", false
	}
}

// HasFuncGet returns true if the function "WEBEXT.chromeosInfoPrivate.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.chromeosInfoPrivate.get".
func FuncGet() (fn js.Func[func(propertyNames js.Array[js.String]) js.Promise[GetReturnType]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.chromeosInfoPrivate.get" directly.
func Get(propertyNames js.Array[js.String]) (ret js.Promise[GetReturnType]) {
	bindings.CallGet(
		js.Pointer(&ret),
		propertyNames.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.chromeosInfoPrivate.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(propertyNames js.Array[js.String]) (ret js.Promise[GetReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		propertyNames.Ref(),
	)

	return
}

// HasFuncIsTabletModeEnabled returns true if the function "WEBEXT.chromeosInfoPrivate.isTabletModeEnabled" exists.
func HasFuncIsTabletModeEnabled() bool {
	return js.True == bindings.HasFuncIsTabletModeEnabled()
}

// FuncIsTabletModeEnabled returns the function "WEBEXT.chromeosInfoPrivate.isTabletModeEnabled".
func FuncIsTabletModeEnabled() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsTabletModeEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsTabletModeEnabled calls the function "WEBEXT.chromeosInfoPrivate.isTabletModeEnabled" directly.
func IsTabletModeEnabled() (ret js.Promise[js.Boolean]) {
	bindings.CallIsTabletModeEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryIsTabletModeEnabled calls the function "WEBEXT.chromeosInfoPrivate.isTabletModeEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsTabletModeEnabled() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsTabletModeEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSet returns true if the function "WEBEXT.chromeosInfoPrivate.set" exists.
func HasFuncSet() bool {
	return js.True == bindings.HasFuncSet()
}

// FuncSet returns the function "WEBEXT.chromeosInfoPrivate.set".
func FuncSet() (fn js.Func[func(propertyName PropertyName, propertyValue js.Any)]) {
	bindings.FuncSet(
		js.Pointer(&fn),
	)
	return
}

// Set calls the function "WEBEXT.chromeosInfoPrivate.set" directly.
func Set(propertyName PropertyName, propertyValue js.Any) (ret js.Void) {
	bindings.CallSet(
		js.Pointer(&ret),
		uint32(propertyName),
		propertyValue.Ref(),
	)

	return
}

// TrySet calls the function "WEBEXT.chromeosInfoPrivate.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySet(propertyName PropertyName, propertyValue js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySet(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(propertyName),
		propertyValue.Ref(),
	)

	return
}

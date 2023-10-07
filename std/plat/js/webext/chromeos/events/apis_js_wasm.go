// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package events

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chromeos/events/bindings"
)

type AudioJackDeviceType uint32

const (
	_ AudioJackDeviceType = iota

	AudioJackDeviceType_HEADPHONE
	AudioJackDeviceType_MICROPHONE
)

func (AudioJackDeviceType) FromRef(str js.Ref) AudioJackDeviceType {
	return AudioJackDeviceType(bindings.ConstOfAudioJackDeviceType(str))
}

func (x AudioJackDeviceType) String() (string, bool) {
	switch x {
	case AudioJackDeviceType_HEADPHONE:
		return "headphone", true
	case AudioJackDeviceType_MICROPHONE:
		return "microphone", true
	default:
		return "", false
	}
}

type AudioJackEvent uint32

const (
	_ AudioJackEvent = iota

	AudioJackEvent_CONNECTED
	AudioJackEvent_DISCONNECTED
)

func (AudioJackEvent) FromRef(str js.Ref) AudioJackEvent {
	return AudioJackEvent(bindings.ConstOfAudioJackEvent(str))
}

func (x AudioJackEvent) String() (string, bool) {
	switch x {
	case AudioJackEvent_CONNECTED:
		return "connected", true
	case AudioJackEvent_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type AudioJackEventInfo struct {
	// Event is "AudioJackEventInfo.event"
	//
	// Optional
	Event AudioJackEvent
	// DeviceType is "AudioJackEventInfo.deviceType"
	//
	// Optional
	DeviceType AudioJackDeviceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioJackEventInfo with all fields set.
func (p AudioJackEventInfo) FromRef(ref js.Ref) AudioJackEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioJackEventInfo in the application heap.
func (p AudioJackEventInfo) New() js.Ref {
	return bindings.AudioJackEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioJackEventInfo) UpdateFrom(ref js.Ref) {
	bindings.AudioJackEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioJackEventInfo) Update(ref js.Ref) {
	bindings.AudioJackEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioJackEventInfo) FreeMembers(recursive bool) {
}

type DisplayInputType uint32

const (
	_ DisplayInputType = iota

	DisplayInputType_UNKNOWN
	DisplayInputType_DIGITAL
	DisplayInputType_ANALOG
)

func (DisplayInputType) FromRef(str js.Ref) DisplayInputType {
	return DisplayInputType(bindings.ConstOfDisplayInputType(str))
}

func (x DisplayInputType) String() (string, bool) {
	switch x {
	case DisplayInputType_UNKNOWN:
		return "unknown", true
	case DisplayInputType_DIGITAL:
		return "digital", true
	case DisplayInputType_ANALOG:
		return "analog", true
	default:
		return "", false
	}
}

type EOFFunc func(this js.Ref) js.Ref

func (fn EOFFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EOFFunc) DispatchCallback(
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

type EOF[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *EOF[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EOF[T]) DispatchCallback(
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

type EventCategory uint32

const (
	_ EventCategory = iota

	EventCategory_AUDIO_JACK
	EventCategory_LID
	EventCategory_USB
	EventCategory_SD_CARD
	EventCategory_POWER
	EventCategory_KEYBOARD_DIAGNOSTIC
	EventCategory_STYLUS_GARAGE
	EventCategory_TOUCHPAD_BUTTON
	EventCategory_TOUCHPAD_TOUCH
	EventCategory_TOUCHPAD_CONNECTED
	EventCategory_TOUCHSCREEN_TOUCH
	EventCategory_TOUCHSCREEN_CONNECTED
	EventCategory_EXTERNAL_DISPLAY
	EventCategory_STYLUS_TOUCH
	EventCategory_STYLUS_CONNECTED
)

func (EventCategory) FromRef(str js.Ref) EventCategory {
	return EventCategory(bindings.ConstOfEventCategory(str))
}

func (x EventCategory) String() (string, bool) {
	switch x {
	case EventCategory_AUDIO_JACK:
		return "audio_jack", true
	case EventCategory_LID:
		return "lid", true
	case EventCategory_USB:
		return "usb", true
	case EventCategory_SD_CARD:
		return "sd_card", true
	case EventCategory_POWER:
		return "power", true
	case EventCategory_KEYBOARD_DIAGNOSTIC:
		return "keyboard_diagnostic", true
	case EventCategory_STYLUS_GARAGE:
		return "stylus_garage", true
	case EventCategory_TOUCHPAD_BUTTON:
		return "touchpad_button", true
	case EventCategory_TOUCHPAD_TOUCH:
		return "touchpad_touch", true
	case EventCategory_TOUCHPAD_CONNECTED:
		return "touchpad_connected", true
	case EventCategory_TOUCHSCREEN_TOUCH:
		return "touchscreen_touch", true
	case EventCategory_TOUCHSCREEN_CONNECTED:
		return "touchscreen_connected", true
	case EventCategory_EXTERNAL_DISPLAY:
		return "external_display", true
	case EventCategory_STYLUS_TOUCH:
		return "stylus_touch", true
	case EventCategory_STYLUS_CONNECTED:
		return "stylus_connected", true
	default:
		return "", false
	}
}

type EventSupportStatus uint32

const (
	_ EventSupportStatus = iota

	EventSupportStatus_SUPPORTED
	EventSupportStatus_UNSUPPORTED
)

func (EventSupportStatus) FromRef(str js.Ref) EventSupportStatus {
	return EventSupportStatus(bindings.ConstOfEventSupportStatus(str))
}

func (x EventSupportStatus) String() (string, bool) {
	switch x {
	case EventSupportStatus_SUPPORTED:
		return "supported", true
	case EventSupportStatus_UNSUPPORTED:
		return "unsupported", true
	default:
		return "", false
	}
}

type EventSupportStatusInfo struct {
	// Status is "EventSupportStatusInfo.status"
	//
	// Optional
	Status EventSupportStatus

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EventSupportStatusInfo with all fields set.
func (p EventSupportStatusInfo) FromRef(ref js.Ref) EventSupportStatusInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventSupportStatusInfo in the application heap.
func (p EventSupportStatusInfo) New() js.Ref {
	return bindings.EventSupportStatusInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EventSupportStatusInfo) UpdateFrom(ref js.Ref) {
	bindings.EventSupportStatusInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EventSupportStatusInfo) Update(ref js.Ref) {
	bindings.EventSupportStatusInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EventSupportStatusInfo) FreeMembers(recursive bool) {
}

type EventSupportStatusInfoCallbackFunc func(this js.Ref, info *EventSupportStatusInfo) js.Ref

func (fn EventSupportStatusInfoCallbackFunc) Register() js.Func[func(info *EventSupportStatusInfo)] {
	return js.RegisterCallback[func(info *EventSupportStatusInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EventSupportStatusInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EventSupportStatusInfo
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

type EventSupportStatusInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *EventSupportStatusInfo) js.Ref
	Arg T
}

func (cb *EventSupportStatusInfoCallback[T]) Register() js.Func[func(info *EventSupportStatusInfo)] {
	return js.RegisterCallback[func(info *EventSupportStatusInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EventSupportStatusInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EventSupportStatusInfo
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

type ExternalDisplayEvent uint32

const (
	_ ExternalDisplayEvent = iota

	ExternalDisplayEvent_CONNECTED
	ExternalDisplayEvent_DISCONNECTED
)

func (ExternalDisplayEvent) FromRef(str js.Ref) ExternalDisplayEvent {
	return ExternalDisplayEvent(bindings.ConstOfExternalDisplayEvent(str))
}

func (x ExternalDisplayEvent) String() (string, bool) {
	switch x {
	case ExternalDisplayEvent_CONNECTED:
		return "connected", true
	case ExternalDisplayEvent_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type ExternalDisplayInfo struct {
	// DisplayWidth is "ExternalDisplayInfo.displayWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayWidth MUST be set to true to make this field effective.
	DisplayWidth int32
	// DisplayHeight is "ExternalDisplayInfo.displayHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayHeight MUST be set to true to make this field effective.
	DisplayHeight int32
	// ResolutionHorizontal is "ExternalDisplayInfo.resolutionHorizontal"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionHorizontal MUST be set to true to make this field effective.
	ResolutionHorizontal int32
	// ResolutionVertical is "ExternalDisplayInfo.resolutionVertical"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionVertical MUST be set to true to make this field effective.
	ResolutionVertical int32
	// RefreshRate is "ExternalDisplayInfo.refreshRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_RefreshRate MUST be set to true to make this field effective.
	RefreshRate float64
	// Manufacturer is "ExternalDisplayInfo.manufacturer"
	//
	// Optional
	Manufacturer js.String
	// ModelId is "ExternalDisplayInfo.modelId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModelId MUST be set to true to make this field effective.
	ModelId int32
	// SerialNumber is "ExternalDisplayInfo.serialNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_SerialNumber MUST be set to true to make this field effective.
	SerialNumber int32
	// ManufactureWeek is "ExternalDisplayInfo.manufactureWeek"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureWeek MUST be set to true to make this field effective.
	ManufactureWeek int32
	// ManufactureYear is "ExternalDisplayInfo.manufactureYear"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureYear MUST be set to true to make this field effective.
	ManufactureYear int32
	// EdidVersion is "ExternalDisplayInfo.edidVersion"
	//
	// Optional
	EdidVersion js.String
	// InputType is "ExternalDisplayInfo.inputType"
	//
	// Optional
	InputType DisplayInputType
	// DisplayName is "ExternalDisplayInfo.displayName"
	//
	// Optional
	DisplayName js.String

	FFI_USE_DisplayWidth         bool // for DisplayWidth.
	FFI_USE_DisplayHeight        bool // for DisplayHeight.
	FFI_USE_ResolutionHorizontal bool // for ResolutionHorizontal.
	FFI_USE_ResolutionVertical   bool // for ResolutionVertical.
	FFI_USE_RefreshRate          bool // for RefreshRate.
	FFI_USE_ModelId              bool // for ModelId.
	FFI_USE_SerialNumber         bool // for SerialNumber.
	FFI_USE_ManufactureWeek      bool // for ManufactureWeek.
	FFI_USE_ManufactureYear      bool // for ManufactureYear.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExternalDisplayInfo with all fields set.
func (p ExternalDisplayInfo) FromRef(ref js.Ref) ExternalDisplayInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExternalDisplayInfo in the application heap.
func (p ExternalDisplayInfo) New() js.Ref {
	return bindings.ExternalDisplayInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExternalDisplayInfo) UpdateFrom(ref js.Ref) {
	bindings.ExternalDisplayInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExternalDisplayInfo) Update(ref js.Ref) {
	bindings.ExternalDisplayInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExternalDisplayInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Manufacturer.Ref(),
		p.EdidVersion.Ref(),
		p.DisplayName.Ref(),
	)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
	p.EdidVersion = p.EdidVersion.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type ExternalDisplayEventInfo struct {
	// Event is "ExternalDisplayEventInfo.event"
	//
	// Optional
	Event ExternalDisplayEvent
	// DisplayInfo is "ExternalDisplayEventInfo.displayInfo"
	//
	// Optional
	//
	// NOTE: DisplayInfo.FFI_USE MUST be set to true to get DisplayInfo used.
	DisplayInfo ExternalDisplayInfo

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExternalDisplayEventInfo with all fields set.
func (p ExternalDisplayEventInfo) FromRef(ref js.Ref) ExternalDisplayEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExternalDisplayEventInfo in the application heap.
func (p ExternalDisplayEventInfo) New() js.Ref {
	return bindings.ExternalDisplayEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExternalDisplayEventInfo) UpdateFrom(ref js.Ref) {
	bindings.ExternalDisplayEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExternalDisplayEventInfo) Update(ref js.Ref) {
	bindings.ExternalDisplayEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExternalDisplayEventInfo) FreeMembers(recursive bool) {
	if recursive {
		p.DisplayInfo.FreeMembers(true)
	}
}

type InputTouchButton uint32

const (
	_ InputTouchButton = iota

	InputTouchButton_LEFT
	InputTouchButton_MIDDLE
	InputTouchButton_RIGHT
)

func (InputTouchButton) FromRef(str js.Ref) InputTouchButton {
	return InputTouchButton(bindings.ConstOfInputTouchButton(str))
}

func (x InputTouchButton) String() (string, bool) {
	switch x {
	case InputTouchButton_LEFT:
		return "left", true
	case InputTouchButton_MIDDLE:
		return "middle", true
	case InputTouchButton_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type InputTouchButtonState uint32

const (
	_ InputTouchButtonState = iota

	InputTouchButtonState_PRESSED
	InputTouchButtonState_RELEASED
)

func (InputTouchButtonState) FromRef(str js.Ref) InputTouchButtonState {
	return InputTouchButtonState(bindings.ConstOfInputTouchButtonState(str))
}

func (x InputTouchButtonState) String() (string, bool) {
	switch x {
	case InputTouchButtonState_PRESSED:
		return "pressed", true
	case InputTouchButtonState_RELEASED:
		return "released", true
	default:
		return "", false
	}
}

type KeyboardConnectionType uint32

const (
	_ KeyboardConnectionType = iota

	KeyboardConnectionType_INTERNAL
	KeyboardConnectionType_USB
	KeyboardConnectionType_BLUETOOTH
	KeyboardConnectionType_UNKNOWN
)

func (KeyboardConnectionType) FromRef(str js.Ref) KeyboardConnectionType {
	return KeyboardConnectionType(bindings.ConstOfKeyboardConnectionType(str))
}

func (x KeyboardConnectionType) String() (string, bool) {
	switch x {
	case KeyboardConnectionType_INTERNAL:
		return "internal", true
	case KeyboardConnectionType_USB:
		return "usb", true
	case KeyboardConnectionType_BLUETOOTH:
		return "bluetooth", true
	case KeyboardConnectionType_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type PhysicalKeyboardLayout uint32

const (
	_ PhysicalKeyboardLayout = iota

	PhysicalKeyboardLayout_UNKNOWN
	PhysicalKeyboardLayout_CHROME_OS
)

func (PhysicalKeyboardLayout) FromRef(str js.Ref) PhysicalKeyboardLayout {
	return PhysicalKeyboardLayout(bindings.ConstOfPhysicalKeyboardLayout(str))
}

func (x PhysicalKeyboardLayout) String() (string, bool) {
	switch x {
	case PhysicalKeyboardLayout_UNKNOWN:
		return "unknown", true
	case PhysicalKeyboardLayout_CHROME_OS:
		return "chrome_os", true
	default:
		return "", false
	}
}

type MechanicalKeyboardLayout uint32

const (
	_ MechanicalKeyboardLayout = iota

	MechanicalKeyboardLayout_UNKNOWN
	MechanicalKeyboardLayout_ANSI
	MechanicalKeyboardLayout_ISO
	MechanicalKeyboardLayout_JIS
)

func (MechanicalKeyboardLayout) FromRef(str js.Ref) MechanicalKeyboardLayout {
	return MechanicalKeyboardLayout(bindings.ConstOfMechanicalKeyboardLayout(str))
}

func (x MechanicalKeyboardLayout) String() (string, bool) {
	switch x {
	case MechanicalKeyboardLayout_UNKNOWN:
		return "unknown", true
	case MechanicalKeyboardLayout_ANSI:
		return "ansi", true
	case MechanicalKeyboardLayout_ISO:
		return "iso", true
	case MechanicalKeyboardLayout_JIS:
		return "jis", true
	default:
		return "", false
	}
}

type KeyboardNumberPadPresence uint32

const (
	_ KeyboardNumberPadPresence = iota

	KeyboardNumberPadPresence_UNKNOWN
	KeyboardNumberPadPresence_PRESENT
	KeyboardNumberPadPresence_NOT_PRESENT
)

func (KeyboardNumberPadPresence) FromRef(str js.Ref) KeyboardNumberPadPresence {
	return KeyboardNumberPadPresence(bindings.ConstOfKeyboardNumberPadPresence(str))
}

func (x KeyboardNumberPadPresence) String() (string, bool) {
	switch x {
	case KeyboardNumberPadPresence_UNKNOWN:
		return "unknown", true
	case KeyboardNumberPadPresence_PRESENT:
		return "present", true
	case KeyboardNumberPadPresence_NOT_PRESENT:
		return "not_present", true
	default:
		return "", false
	}
}

type KeyboardTopRowKey uint32

const (
	_ KeyboardTopRowKey = iota

	KeyboardTopRowKey_NO_KEY
	KeyboardTopRowKey_UNKNOWN
	KeyboardTopRowKey_BACK
	KeyboardTopRowKey_FORWARD
	KeyboardTopRowKey_REFRESH
	KeyboardTopRowKey_FULLSCREEN
	KeyboardTopRowKey_OVERVIEW
	KeyboardTopRowKey_SCREENSHOT
	KeyboardTopRowKey_SCREEN_BRIGHTNESS_DOWN
	KeyboardTopRowKey_SCREEN_BRIGHTNESS_UP
	KeyboardTopRowKey_PRIVACY_SCREEN_TOGGLE
	KeyboardTopRowKey_MICROPHONE_MUTE
	KeyboardTopRowKey_VOLUME_MUTE
	KeyboardTopRowKey_VOLUME_DOWN
	KeyboardTopRowKey_VOLUME_UP
	KeyboardTopRowKey_KEYBOARD_BACKLIGHT_TOGGLE
	KeyboardTopRowKey_KEYBOARD_BACKLIGHT_DOWN
	KeyboardTopRowKey_KEYBOARD_BACKLIGHT_UP
	KeyboardTopRowKey_NEXT_TRACK
	KeyboardTopRowKey_PREVIOUS_TRACK
	KeyboardTopRowKey_PLAY_PAUSE
	KeyboardTopRowKey_SCREEN_MIRROR
	KeyboardTopRowKey_DELETE
)

func (KeyboardTopRowKey) FromRef(str js.Ref) KeyboardTopRowKey {
	return KeyboardTopRowKey(bindings.ConstOfKeyboardTopRowKey(str))
}

func (x KeyboardTopRowKey) String() (string, bool) {
	switch x {
	case KeyboardTopRowKey_NO_KEY:
		return "no_key", true
	case KeyboardTopRowKey_UNKNOWN:
		return "unknown", true
	case KeyboardTopRowKey_BACK:
		return "back", true
	case KeyboardTopRowKey_FORWARD:
		return "forward", true
	case KeyboardTopRowKey_REFRESH:
		return "refresh", true
	case KeyboardTopRowKey_FULLSCREEN:
		return "fullscreen", true
	case KeyboardTopRowKey_OVERVIEW:
		return "overview", true
	case KeyboardTopRowKey_SCREENSHOT:
		return "screenshot", true
	case KeyboardTopRowKey_SCREEN_BRIGHTNESS_DOWN:
		return "screen_brightness_down", true
	case KeyboardTopRowKey_SCREEN_BRIGHTNESS_UP:
		return "screen_brightness_up", true
	case KeyboardTopRowKey_PRIVACY_SCREEN_TOGGLE:
		return "privacy_screen_toggle", true
	case KeyboardTopRowKey_MICROPHONE_MUTE:
		return "microphone_mute", true
	case KeyboardTopRowKey_VOLUME_MUTE:
		return "volume_mute", true
	case KeyboardTopRowKey_VOLUME_DOWN:
		return "volume_down", true
	case KeyboardTopRowKey_VOLUME_UP:
		return "volume_up", true
	case KeyboardTopRowKey_KEYBOARD_BACKLIGHT_TOGGLE:
		return "keyboard_backlight_toggle", true
	case KeyboardTopRowKey_KEYBOARD_BACKLIGHT_DOWN:
		return "keyboard_backlight_down", true
	case KeyboardTopRowKey_KEYBOARD_BACKLIGHT_UP:
		return "keyboard_backlight_up", true
	case KeyboardTopRowKey_NEXT_TRACK:
		return "next_track", true
	case KeyboardTopRowKey_PREVIOUS_TRACK:
		return "previous_track", true
	case KeyboardTopRowKey_PLAY_PAUSE:
		return "play_pause", true
	case KeyboardTopRowKey_SCREEN_MIRROR:
		return "screen_mirror", true
	case KeyboardTopRowKey_DELETE:
		return "delete", true
	default:
		return "", false
	}
}

type KeyboardTopRightKey uint32

const (
	_ KeyboardTopRightKey = iota

	KeyboardTopRightKey_UNKNOWN
	KeyboardTopRightKey_POWER
	KeyboardTopRightKey_LOCK
	KeyboardTopRightKey_CONTROL_PANEL
)

func (KeyboardTopRightKey) FromRef(str js.Ref) KeyboardTopRightKey {
	return KeyboardTopRightKey(bindings.ConstOfKeyboardTopRightKey(str))
}

func (x KeyboardTopRightKey) String() (string, bool) {
	switch x {
	case KeyboardTopRightKey_UNKNOWN:
		return "unknown", true
	case KeyboardTopRightKey_POWER:
		return "power", true
	case KeyboardTopRightKey_LOCK:
		return "lock", true
	case KeyboardTopRightKey_CONTROL_PANEL:
		return "control_panel", true
	default:
		return "", false
	}
}

type KeyboardInfo struct {
	// Id is "KeyboardInfo.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// ConnectionType is "KeyboardInfo.connectionType"
	//
	// Optional
	ConnectionType KeyboardConnectionType
	// Name is "KeyboardInfo.name"
	//
	// Optional
	Name js.String
	// PhysicalLayout is "KeyboardInfo.physicalLayout"
	//
	// Optional
	PhysicalLayout PhysicalKeyboardLayout
	// MechanicalLayout is "KeyboardInfo.mechanicalLayout"
	//
	// Optional
	MechanicalLayout MechanicalKeyboardLayout
	// RegionCode is "KeyboardInfo.regionCode"
	//
	// Optional
	RegionCode js.String
	// NumberPadPresent is "KeyboardInfo.numberPadPresent"
	//
	// Optional
	NumberPadPresent KeyboardNumberPadPresence
	// TopRowKeys is "KeyboardInfo.topRowKeys"
	//
	// Optional
	TopRowKeys js.Array[KeyboardTopRowKey]
	// TopRightKey is "KeyboardInfo.topRightKey"
	//
	// Optional
	TopRightKey KeyboardTopRightKey
	// HasAssistantKey is "KeyboardInfo.hasAssistantKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasAssistantKey MUST be set to true to make this field effective.
	HasAssistantKey bool

	FFI_USE_Id              bool // for Id.
	FFI_USE_HasAssistantKey bool // for HasAssistantKey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyboardInfo with all fields set.
func (p KeyboardInfo) FromRef(ref js.Ref) KeyboardInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyboardInfo in the application heap.
func (p KeyboardInfo) New() js.Ref {
	return bindings.KeyboardInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KeyboardInfo) UpdateFrom(ref js.Ref) {
	bindings.KeyboardInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyboardInfo) Update(ref js.Ref) {
	bindings.KeyboardInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyboardInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.RegionCode.Ref(),
		p.TopRowKeys.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.RegionCode = p.RegionCode.FromRef(js.Undefined)
	p.TopRowKeys = p.TopRowKeys.FromRef(js.Undefined)
}

type KeyboardDiagnosticEventInfo struct {
	// KeyboardInfo is "KeyboardDiagnosticEventInfo.keyboardInfo"
	//
	// Optional
	//
	// NOTE: KeyboardInfo.FFI_USE MUST be set to true to get KeyboardInfo used.
	KeyboardInfo KeyboardInfo
	// TestedKeys is "KeyboardDiagnosticEventInfo.testedKeys"
	//
	// Optional
	TestedKeys js.Array[int32]
	// TestedTopRowKeys is "KeyboardDiagnosticEventInfo.testedTopRowKeys"
	//
	// Optional
	TestedTopRowKeys js.Array[int32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyboardDiagnosticEventInfo with all fields set.
func (p KeyboardDiagnosticEventInfo) FromRef(ref js.Ref) KeyboardDiagnosticEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyboardDiagnosticEventInfo in the application heap.
func (p KeyboardDiagnosticEventInfo) New() js.Ref {
	return bindings.KeyboardDiagnosticEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KeyboardDiagnosticEventInfo) UpdateFrom(ref js.Ref) {
	bindings.KeyboardDiagnosticEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyboardDiagnosticEventInfo) Update(ref js.Ref) {
	bindings.KeyboardDiagnosticEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyboardDiagnosticEventInfo) FreeMembers(recursive bool) {
	js.Free(
		p.TestedKeys.Ref(),
		p.TestedTopRowKeys.Ref(),
	)
	p.TestedKeys = p.TestedKeys.FromRef(js.Undefined)
	p.TestedTopRowKeys = p.TestedTopRowKeys.FromRef(js.Undefined)
	if recursive {
		p.KeyboardInfo.FreeMembers(true)
	}
}

type LidEvent uint32

const (
	_ LidEvent = iota

	LidEvent_CLOSED
	LidEvent_OPENED
)

func (LidEvent) FromRef(str js.Ref) LidEvent {
	return LidEvent(bindings.ConstOfLidEvent(str))
}

func (x LidEvent) String() (string, bool) {
	switch x {
	case LidEvent_CLOSED:
		return "closed", true
	case LidEvent_OPENED:
		return "opened", true
	default:
		return "", false
	}
}

type LidEventInfo struct {
	// Event is "LidEventInfo.event"
	//
	// Optional
	Event LidEvent

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LidEventInfo with all fields set.
func (p LidEventInfo) FromRef(ref js.Ref) LidEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LidEventInfo in the application heap.
func (p LidEventInfo) New() js.Ref {
	return bindings.LidEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LidEventInfo) UpdateFrom(ref js.Ref) {
	bindings.LidEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LidEventInfo) Update(ref js.Ref) {
	bindings.LidEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LidEventInfo) FreeMembers(recursive bool) {
}

type PowerEvent uint32

const (
	_ PowerEvent = iota

	PowerEvent_AC_INSERTED
	PowerEvent_AC_REMOVED
	PowerEvent_OS_SUSPEND
	PowerEvent_OS_RESUME
)

func (PowerEvent) FromRef(str js.Ref) PowerEvent {
	return PowerEvent(bindings.ConstOfPowerEvent(str))
}

func (x PowerEvent) String() (string, bool) {
	switch x {
	case PowerEvent_AC_INSERTED:
		return "ac_inserted", true
	case PowerEvent_AC_REMOVED:
		return "ac_removed", true
	case PowerEvent_OS_SUSPEND:
		return "os_suspend", true
	case PowerEvent_OS_RESUME:
		return "os_resume", true
	default:
		return "", false
	}
}

type PowerEventInfo struct {
	// Event is "PowerEventInfo.event"
	//
	// Optional
	Event PowerEvent

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PowerEventInfo with all fields set.
func (p PowerEventInfo) FromRef(ref js.Ref) PowerEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PowerEventInfo in the application heap.
func (p PowerEventInfo) New() js.Ref {
	return bindings.PowerEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PowerEventInfo) UpdateFrom(ref js.Ref) {
	bindings.PowerEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PowerEventInfo) Update(ref js.Ref) {
	bindings.PowerEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PowerEventInfo) FreeMembers(recursive bool) {
}

type SdCardEvent uint32

const (
	_ SdCardEvent = iota

	SdCardEvent_CONNECTED
	SdCardEvent_DISCONNECTED
)

func (SdCardEvent) FromRef(str js.Ref) SdCardEvent {
	return SdCardEvent(bindings.ConstOfSdCardEvent(str))
}

func (x SdCardEvent) String() (string, bool) {
	switch x {
	case SdCardEvent_CONNECTED:
		return "connected", true
	case SdCardEvent_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type SdCardEventInfo struct {
	// Event is "SdCardEventInfo.event"
	//
	// Optional
	Event SdCardEvent

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SdCardEventInfo with all fields set.
func (p SdCardEventInfo) FromRef(ref js.Ref) SdCardEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SdCardEventInfo in the application heap.
func (p SdCardEventInfo) New() js.Ref {
	return bindings.SdCardEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SdCardEventInfo) UpdateFrom(ref js.Ref) {
	bindings.SdCardEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SdCardEventInfo) Update(ref js.Ref) {
	bindings.SdCardEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SdCardEventInfo) FreeMembers(recursive bool) {
}

type StylusConnectedEventInfo struct {
	// MaxX is "StylusConnectedEventInfo.max_x"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxX MUST be set to true to make this field effective.
	MaxX int32
	// MaxY is "StylusConnectedEventInfo.max_y"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxY MUST be set to true to make this field effective.
	MaxY int32
	// MaxPressure is "StylusConnectedEventInfo.max_pressure"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxPressure MUST be set to true to make this field effective.
	MaxPressure int32

	FFI_USE_MaxX        bool // for MaxX.
	FFI_USE_MaxY        bool // for MaxY.
	FFI_USE_MaxPressure bool // for MaxPressure.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StylusConnectedEventInfo with all fields set.
func (p StylusConnectedEventInfo) FromRef(ref js.Ref) StylusConnectedEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StylusConnectedEventInfo in the application heap.
func (p StylusConnectedEventInfo) New() js.Ref {
	return bindings.StylusConnectedEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StylusConnectedEventInfo) UpdateFrom(ref js.Ref) {
	bindings.StylusConnectedEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StylusConnectedEventInfo) Update(ref js.Ref) {
	bindings.StylusConnectedEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StylusConnectedEventInfo) FreeMembers(recursive bool) {
}

type StylusGarageEvent uint32

const (
	_ StylusGarageEvent = iota

	StylusGarageEvent_INSERTED
	StylusGarageEvent_REMOVED
)

func (StylusGarageEvent) FromRef(str js.Ref) StylusGarageEvent {
	return StylusGarageEvent(bindings.ConstOfStylusGarageEvent(str))
}

func (x StylusGarageEvent) String() (string, bool) {
	switch x {
	case StylusGarageEvent_INSERTED:
		return "inserted", true
	case StylusGarageEvent_REMOVED:
		return "removed", true
	default:
		return "", false
	}
}

type StylusGarageEventInfo struct {
	// Event is "StylusGarageEventInfo.event"
	//
	// Optional
	Event StylusGarageEvent

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StylusGarageEventInfo with all fields set.
func (p StylusGarageEventInfo) FromRef(ref js.Ref) StylusGarageEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StylusGarageEventInfo in the application heap.
func (p StylusGarageEventInfo) New() js.Ref {
	return bindings.StylusGarageEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StylusGarageEventInfo) UpdateFrom(ref js.Ref) {
	bindings.StylusGarageEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StylusGarageEventInfo) Update(ref js.Ref) {
	bindings.StylusGarageEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StylusGarageEventInfo) FreeMembers(recursive bool) {
}

type StylusTouchPointInfo struct {
	// X is "StylusTouchPointInfo.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "StylusTouchPointInfo.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32
	// Pressure is "StylusTouchPointInfo.pressure"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pressure MUST be set to true to make this field effective.
	Pressure int32

	FFI_USE_X        bool // for X.
	FFI_USE_Y        bool // for Y.
	FFI_USE_Pressure bool // for Pressure.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StylusTouchPointInfo with all fields set.
func (p StylusTouchPointInfo) FromRef(ref js.Ref) StylusTouchPointInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StylusTouchPointInfo in the application heap.
func (p StylusTouchPointInfo) New() js.Ref {
	return bindings.StylusTouchPointInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StylusTouchPointInfo) UpdateFrom(ref js.Ref) {
	bindings.StylusTouchPointInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StylusTouchPointInfo) Update(ref js.Ref) {
	bindings.StylusTouchPointInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StylusTouchPointInfo) FreeMembers(recursive bool) {
}

type StylusTouchEventInfo struct {
	// TouchPoint is "StylusTouchEventInfo.touch_point"
	//
	// Optional
	//
	// NOTE: TouchPoint.FFI_USE MUST be set to true to get TouchPoint used.
	TouchPoint StylusTouchPointInfo

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StylusTouchEventInfo with all fields set.
func (p StylusTouchEventInfo) FromRef(ref js.Ref) StylusTouchEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StylusTouchEventInfo in the application heap.
func (p StylusTouchEventInfo) New() js.Ref {
	return bindings.StylusTouchEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StylusTouchEventInfo) UpdateFrom(ref js.Ref) {
	bindings.StylusTouchEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StylusTouchEventInfo) Update(ref js.Ref) {
	bindings.StylusTouchEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StylusTouchEventInfo) FreeMembers(recursive bool) {
	if recursive {
		p.TouchPoint.FreeMembers(true)
	}
}

type TouchPointInfo struct {
	// TrackingId is "TouchPointInfo.trackingId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TrackingId MUST be set to true to make this field effective.
	TrackingId int32
	// X is "TouchPointInfo.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "TouchPointInfo.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32
	// Pressure is "TouchPointInfo.pressure"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pressure MUST be set to true to make this field effective.
	Pressure int32
	// TouchMajor is "TouchPointInfo.touchMajor"
	//
	// Optional
	//
	// NOTE: FFI_USE_TouchMajor MUST be set to true to make this field effective.
	TouchMajor int32
	// TouchMinor is "TouchPointInfo.touchMinor"
	//
	// Optional
	//
	// NOTE: FFI_USE_TouchMinor MUST be set to true to make this field effective.
	TouchMinor int32

	FFI_USE_TrackingId bool // for TrackingId.
	FFI_USE_X          bool // for X.
	FFI_USE_Y          bool // for Y.
	FFI_USE_Pressure   bool // for Pressure.
	FFI_USE_TouchMajor bool // for TouchMajor.
	FFI_USE_TouchMinor bool // for TouchMinor.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchPointInfo with all fields set.
func (p TouchPointInfo) FromRef(ref js.Ref) TouchPointInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchPointInfo in the application heap.
func (p TouchPointInfo) New() js.Ref {
	return bindings.TouchPointInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchPointInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchPointInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchPointInfo) Update(ref js.Ref) {
	bindings.TouchPointInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchPointInfo) FreeMembers(recursive bool) {
}

type TouchpadButtonEventInfo struct {
	// Button is "TouchpadButtonEventInfo.button"
	//
	// Optional
	Button InputTouchButton
	// State is "TouchpadButtonEventInfo.state"
	//
	// Optional
	State InputTouchButtonState

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchpadButtonEventInfo with all fields set.
func (p TouchpadButtonEventInfo) FromRef(ref js.Ref) TouchpadButtonEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchpadButtonEventInfo in the application heap.
func (p TouchpadButtonEventInfo) New() js.Ref {
	return bindings.TouchpadButtonEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchpadButtonEventInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchpadButtonEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchpadButtonEventInfo) Update(ref js.Ref) {
	bindings.TouchpadButtonEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchpadButtonEventInfo) FreeMembers(recursive bool) {
}

type TouchpadConnectedEventInfo struct {
	// MaxX is "TouchpadConnectedEventInfo.maxX"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxX MUST be set to true to make this field effective.
	MaxX int32
	// MaxY is "TouchpadConnectedEventInfo.maxY"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxY MUST be set to true to make this field effective.
	MaxY int32
	// MaxPressure is "TouchpadConnectedEventInfo.maxPressure"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxPressure MUST be set to true to make this field effective.
	MaxPressure int32
	// Buttons is "TouchpadConnectedEventInfo.buttons"
	//
	// Optional
	Buttons js.Array[InputTouchButton]

	FFI_USE_MaxX        bool // for MaxX.
	FFI_USE_MaxY        bool // for MaxY.
	FFI_USE_MaxPressure bool // for MaxPressure.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchpadConnectedEventInfo with all fields set.
func (p TouchpadConnectedEventInfo) FromRef(ref js.Ref) TouchpadConnectedEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchpadConnectedEventInfo in the application heap.
func (p TouchpadConnectedEventInfo) New() js.Ref {
	return bindings.TouchpadConnectedEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchpadConnectedEventInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchpadConnectedEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchpadConnectedEventInfo) Update(ref js.Ref) {
	bindings.TouchpadConnectedEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchpadConnectedEventInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Buttons.Ref(),
	)
	p.Buttons = p.Buttons.FromRef(js.Undefined)
}

type TouchpadTouchEventInfo struct {
	// TouchPoints is "TouchpadTouchEventInfo.touchPoints"
	//
	// Optional
	TouchPoints js.Array[TouchPointInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchpadTouchEventInfo with all fields set.
func (p TouchpadTouchEventInfo) FromRef(ref js.Ref) TouchpadTouchEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchpadTouchEventInfo in the application heap.
func (p TouchpadTouchEventInfo) New() js.Ref {
	return bindings.TouchpadTouchEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchpadTouchEventInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchpadTouchEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchpadTouchEventInfo) Update(ref js.Ref) {
	bindings.TouchpadTouchEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchpadTouchEventInfo) FreeMembers(recursive bool) {
	js.Free(
		p.TouchPoints.Ref(),
	)
	p.TouchPoints = p.TouchPoints.FromRef(js.Undefined)
}

type TouchscreenConnectedEventInfo struct {
	// MaxX is "TouchscreenConnectedEventInfo.maxX"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxX MUST be set to true to make this field effective.
	MaxX int32
	// MaxY is "TouchscreenConnectedEventInfo.maxY"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxY MUST be set to true to make this field effective.
	MaxY int32
	// MaxPressure is "TouchscreenConnectedEventInfo.maxPressure"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxPressure MUST be set to true to make this field effective.
	MaxPressure int32

	FFI_USE_MaxX        bool // for MaxX.
	FFI_USE_MaxY        bool // for MaxY.
	FFI_USE_MaxPressure bool // for MaxPressure.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchscreenConnectedEventInfo with all fields set.
func (p TouchscreenConnectedEventInfo) FromRef(ref js.Ref) TouchscreenConnectedEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchscreenConnectedEventInfo in the application heap.
func (p TouchscreenConnectedEventInfo) New() js.Ref {
	return bindings.TouchscreenConnectedEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchscreenConnectedEventInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchscreenConnectedEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchscreenConnectedEventInfo) Update(ref js.Ref) {
	bindings.TouchscreenConnectedEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchscreenConnectedEventInfo) FreeMembers(recursive bool) {
}

type TouchscreenTouchEventInfo struct {
	// TouchPoints is "TouchscreenTouchEventInfo.touchPoints"
	//
	// Optional
	TouchPoints js.Array[TouchPointInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchscreenTouchEventInfo with all fields set.
func (p TouchscreenTouchEventInfo) FromRef(ref js.Ref) TouchscreenTouchEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchscreenTouchEventInfo in the application heap.
func (p TouchscreenTouchEventInfo) New() js.Ref {
	return bindings.TouchscreenTouchEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchscreenTouchEventInfo) UpdateFrom(ref js.Ref) {
	bindings.TouchscreenTouchEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchscreenTouchEventInfo) Update(ref js.Ref) {
	bindings.TouchscreenTouchEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchscreenTouchEventInfo) FreeMembers(recursive bool) {
	js.Free(
		p.TouchPoints.Ref(),
	)
	p.TouchPoints = p.TouchPoints.FromRef(js.Undefined)
}

type UsbEvent uint32

const (
	_ UsbEvent = iota

	UsbEvent_CONNECTED
	UsbEvent_DISCONNECTED
)

func (UsbEvent) FromRef(str js.Ref) UsbEvent {
	return UsbEvent(bindings.ConstOfUsbEvent(str))
}

func (x UsbEvent) String() (string, bool) {
	switch x {
	case UsbEvent_CONNECTED:
		return "connected", true
	case UsbEvent_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type UsbEventInfo struct {
	// Vendor is "UsbEventInfo.vendor"
	//
	// Optional
	Vendor js.String
	// Name is "UsbEventInfo.name"
	//
	// Optional
	Name js.String
	// Vid is "UsbEventInfo.vid"
	//
	// Optional
	//
	// NOTE: FFI_USE_Vid MUST be set to true to make this field effective.
	Vid int32
	// Pid is "UsbEventInfo.pid"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pid MUST be set to true to make this field effective.
	Pid int32
	// Categories is "UsbEventInfo.categories"
	//
	// Optional
	Categories js.Array[js.String]
	// Event is "UsbEventInfo.event"
	//
	// Optional
	Event UsbEvent

	FFI_USE_Vid bool // for Vid.
	FFI_USE_Pid bool // for Pid.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbEventInfo with all fields set.
func (p UsbEventInfo) FromRef(ref js.Ref) UsbEventInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbEventInfo in the application heap.
func (p UsbEventInfo) New() js.Ref {
	return bindings.UsbEventInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbEventInfo) UpdateFrom(ref js.Ref) {
	bindings.UsbEventInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbEventInfo) Update(ref js.Ref) {
	bindings.UsbEventInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbEventInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Vendor.Ref(),
		p.Name.Ref(),
		p.Categories.Ref(),
	)
	p.Vendor = p.Vendor.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Categories = p.Categories.FromRef(js.Undefined)
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

// HasFuncIsEventSupported returns true if the function "WEBEXT.os.events.isEventSupported" exists.
func HasFuncIsEventSupported() bool {
	return js.True == bindings.HasFuncIsEventSupported()
}

// FuncIsEventSupported returns the function "WEBEXT.os.events.isEventSupported".
func FuncIsEventSupported() (fn js.Func[func(category EventCategory) js.Promise[EventSupportStatusInfo]]) {
	bindings.FuncIsEventSupported(
		js.Pointer(&fn),
	)
	return
}

// IsEventSupported calls the function "WEBEXT.os.events.isEventSupported" directly.
func IsEventSupported(category EventCategory) (ret js.Promise[EventSupportStatusInfo]) {
	bindings.CallIsEventSupported(
		js.Pointer(&ret),
		uint32(category),
	)

	return
}

// TryIsEventSupported calls the function "WEBEXT.os.events.isEventSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsEventSupported(category EventCategory) (ret js.Promise[EventSupportStatusInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsEventSupported(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(category),
	)

	return
}

type OnAudioJackEventEventCallbackFunc func(this js.Ref, event_info *AudioJackEventInfo) js.Ref

func (fn OnAudioJackEventEventCallbackFunc) Register() js.Func[func(event_info *AudioJackEventInfo)] {
	return js.RegisterCallback[func(event_info *AudioJackEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAudioJackEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioJackEventInfo
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

type OnAudioJackEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *AudioJackEventInfo) js.Ref
	Arg T
}

func (cb *OnAudioJackEventEventCallback[T]) Register() js.Func[func(event_info *AudioJackEventInfo)] {
	return js.RegisterCallback[func(event_info *AudioJackEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAudioJackEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioJackEventInfo
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

// HasFuncOnAudioJackEvent returns true if the function "WEBEXT.os.events.onAudioJackEvent.addListener" exists.
func HasFuncOnAudioJackEvent() bool {
	return js.True == bindings.HasFuncOnAudioJackEvent()
}

// FuncOnAudioJackEvent returns the function "WEBEXT.os.events.onAudioJackEvent.addListener".
func FuncOnAudioJackEvent() (fn js.Func[func(callback js.Func[func(event_info *AudioJackEventInfo)])]) {
	bindings.FuncOnAudioJackEvent(
		js.Pointer(&fn),
	)
	return
}

// OnAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.addListener" directly.
func OnAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret js.Void) {
	bindings.CallOnAudioJackEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAudioJackEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAudioJackEvent returns true if the function "WEBEXT.os.events.onAudioJackEvent.removeListener" exists.
func HasFuncOffAudioJackEvent() bool {
	return js.True == bindings.HasFuncOffAudioJackEvent()
}

// FuncOffAudioJackEvent returns the function "WEBEXT.os.events.onAudioJackEvent.removeListener".
func FuncOffAudioJackEvent() (fn js.Func[func(callback js.Func[func(event_info *AudioJackEventInfo)])]) {
	bindings.FuncOffAudioJackEvent(
		js.Pointer(&fn),
	)
	return
}

// OffAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.removeListener" directly.
func OffAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret js.Void) {
	bindings.CallOffAudioJackEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAudioJackEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAudioJackEvent returns true if the function "WEBEXT.os.events.onAudioJackEvent.hasListener" exists.
func HasFuncHasOnAudioJackEvent() bool {
	return js.True == bindings.HasFuncHasOnAudioJackEvent()
}

// FuncHasOnAudioJackEvent returns the function "WEBEXT.os.events.onAudioJackEvent.hasListener".
func FuncHasOnAudioJackEvent() (fn js.Func[func(callback js.Func[func(event_info *AudioJackEventInfo)]) bool]) {
	bindings.FuncHasOnAudioJackEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.hasListener" directly.
func HasOnAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret bool) {
	bindings.CallHasOnAudioJackEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAudioJackEvent calls the function "WEBEXT.os.events.onAudioJackEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAudioJackEvent(callback js.Func[func(event_info *AudioJackEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAudioJackEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnExternalDisplayEventEventCallbackFunc func(this js.Ref, event_info *ExternalDisplayEventInfo) js.Ref

func (fn OnExternalDisplayEventEventCallbackFunc) Register() js.Func[func(event_info *ExternalDisplayEventInfo)] {
	return js.RegisterCallback[func(event_info *ExternalDisplayEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnExternalDisplayEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExternalDisplayEventInfo
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

type OnExternalDisplayEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *ExternalDisplayEventInfo) js.Ref
	Arg T
}

func (cb *OnExternalDisplayEventEventCallback[T]) Register() js.Func[func(event_info *ExternalDisplayEventInfo)] {
	return js.RegisterCallback[func(event_info *ExternalDisplayEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnExternalDisplayEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExternalDisplayEventInfo
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

// HasFuncOnExternalDisplayEvent returns true if the function "WEBEXT.os.events.onExternalDisplayEvent.addListener" exists.
func HasFuncOnExternalDisplayEvent() bool {
	return js.True == bindings.HasFuncOnExternalDisplayEvent()
}

// FuncOnExternalDisplayEvent returns the function "WEBEXT.os.events.onExternalDisplayEvent.addListener".
func FuncOnExternalDisplayEvent() (fn js.Func[func(callback js.Func[func(event_info *ExternalDisplayEventInfo)])]) {
	bindings.FuncOnExternalDisplayEvent(
		js.Pointer(&fn),
	)
	return
}

// OnExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.addListener" directly.
func OnExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret js.Void) {
	bindings.CallOnExternalDisplayEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExternalDisplayEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExternalDisplayEvent returns true if the function "WEBEXT.os.events.onExternalDisplayEvent.removeListener" exists.
func HasFuncOffExternalDisplayEvent() bool {
	return js.True == bindings.HasFuncOffExternalDisplayEvent()
}

// FuncOffExternalDisplayEvent returns the function "WEBEXT.os.events.onExternalDisplayEvent.removeListener".
func FuncOffExternalDisplayEvent() (fn js.Func[func(callback js.Func[func(event_info *ExternalDisplayEventInfo)])]) {
	bindings.FuncOffExternalDisplayEvent(
		js.Pointer(&fn),
	)
	return
}

// OffExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.removeListener" directly.
func OffExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret js.Void) {
	bindings.CallOffExternalDisplayEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExternalDisplayEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExternalDisplayEvent returns true if the function "WEBEXT.os.events.onExternalDisplayEvent.hasListener" exists.
func HasFuncHasOnExternalDisplayEvent() bool {
	return js.True == bindings.HasFuncHasOnExternalDisplayEvent()
}

// FuncHasOnExternalDisplayEvent returns the function "WEBEXT.os.events.onExternalDisplayEvent.hasListener".
func FuncHasOnExternalDisplayEvent() (fn js.Func[func(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) bool]) {
	bindings.FuncHasOnExternalDisplayEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.hasListener" directly.
func HasOnExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret bool) {
	bindings.CallHasOnExternalDisplayEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExternalDisplayEvent calls the function "WEBEXT.os.events.onExternalDisplayEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExternalDisplayEvent(callback js.Func[func(event_info *ExternalDisplayEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExternalDisplayEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnKeyboardDiagnosticEventEventCallbackFunc func(this js.Ref, event_info *KeyboardDiagnosticEventInfo) js.Ref

func (fn OnKeyboardDiagnosticEventEventCallbackFunc) Register() js.Func[func(event_info *KeyboardDiagnosticEventInfo)] {
	return js.RegisterCallback[func(event_info *KeyboardDiagnosticEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnKeyboardDiagnosticEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyboardDiagnosticEventInfo
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

type OnKeyboardDiagnosticEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *KeyboardDiagnosticEventInfo) js.Ref
	Arg T
}

func (cb *OnKeyboardDiagnosticEventEventCallback[T]) Register() js.Func[func(event_info *KeyboardDiagnosticEventInfo)] {
	return js.RegisterCallback[func(event_info *KeyboardDiagnosticEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnKeyboardDiagnosticEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 KeyboardDiagnosticEventInfo
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

// HasFuncOnKeyboardDiagnosticEvent returns true if the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener" exists.
func HasFuncOnKeyboardDiagnosticEvent() bool {
	return js.True == bindings.HasFuncOnKeyboardDiagnosticEvent()
}

// FuncOnKeyboardDiagnosticEvent returns the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener".
func FuncOnKeyboardDiagnosticEvent() (fn js.Func[func(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)])]) {
	bindings.FuncOnKeyboardDiagnosticEvent(
		js.Pointer(&fn),
	)
	return
}

// OnKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener" directly.
func OnKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret js.Void) {
	bindings.CallOnKeyboardDiagnosticEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnKeyboardDiagnosticEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffKeyboardDiagnosticEvent returns true if the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener" exists.
func HasFuncOffKeyboardDiagnosticEvent() bool {
	return js.True == bindings.HasFuncOffKeyboardDiagnosticEvent()
}

// FuncOffKeyboardDiagnosticEvent returns the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener".
func FuncOffKeyboardDiagnosticEvent() (fn js.Func[func(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)])]) {
	bindings.FuncOffKeyboardDiagnosticEvent(
		js.Pointer(&fn),
	)
	return
}

// OffKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener" directly.
func OffKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret js.Void) {
	bindings.CallOffKeyboardDiagnosticEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffKeyboardDiagnosticEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnKeyboardDiagnosticEvent returns true if the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener" exists.
func HasFuncHasOnKeyboardDiagnosticEvent() bool {
	return js.True == bindings.HasFuncHasOnKeyboardDiagnosticEvent()
}

// FuncHasOnKeyboardDiagnosticEvent returns the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener".
func FuncHasOnKeyboardDiagnosticEvent() (fn js.Func[func(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) bool]) {
	bindings.FuncHasOnKeyboardDiagnosticEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener" directly.
func HasOnKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret bool) {
	bindings.CallHasOnKeyboardDiagnosticEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnKeyboardDiagnosticEvent calls the function "WEBEXT.os.events.onKeyboardDiagnosticEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnKeyboardDiagnosticEvent(callback js.Func[func(event_info *KeyboardDiagnosticEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnKeyboardDiagnosticEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnLidEventEventCallbackFunc func(this js.Ref, event_info *LidEventInfo) js.Ref

func (fn OnLidEventEventCallbackFunc) Register() js.Func[func(event_info *LidEventInfo)] {
	return js.RegisterCallback[func(event_info *LidEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnLidEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LidEventInfo
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

type OnLidEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *LidEventInfo) js.Ref
	Arg T
}

func (cb *OnLidEventEventCallback[T]) Register() js.Func[func(event_info *LidEventInfo)] {
	return js.RegisterCallback[func(event_info *LidEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnLidEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LidEventInfo
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

// HasFuncOnLidEvent returns true if the function "WEBEXT.os.events.onLidEvent.addListener" exists.
func HasFuncOnLidEvent() bool {
	return js.True == bindings.HasFuncOnLidEvent()
}

// FuncOnLidEvent returns the function "WEBEXT.os.events.onLidEvent.addListener".
func FuncOnLidEvent() (fn js.Func[func(callback js.Func[func(event_info *LidEventInfo)])]) {
	bindings.FuncOnLidEvent(
		js.Pointer(&fn),
	)
	return
}

// OnLidEvent calls the function "WEBEXT.os.events.onLidEvent.addListener" directly.
func OnLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret js.Void) {
	bindings.CallOnLidEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLidEvent calls the function "WEBEXT.os.events.onLidEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLidEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLidEvent returns true if the function "WEBEXT.os.events.onLidEvent.removeListener" exists.
func HasFuncOffLidEvent() bool {
	return js.True == bindings.HasFuncOffLidEvent()
}

// FuncOffLidEvent returns the function "WEBEXT.os.events.onLidEvent.removeListener".
func FuncOffLidEvent() (fn js.Func[func(callback js.Func[func(event_info *LidEventInfo)])]) {
	bindings.FuncOffLidEvent(
		js.Pointer(&fn),
	)
	return
}

// OffLidEvent calls the function "WEBEXT.os.events.onLidEvent.removeListener" directly.
func OffLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret js.Void) {
	bindings.CallOffLidEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLidEvent calls the function "WEBEXT.os.events.onLidEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLidEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLidEvent returns true if the function "WEBEXT.os.events.onLidEvent.hasListener" exists.
func HasFuncHasOnLidEvent() bool {
	return js.True == bindings.HasFuncHasOnLidEvent()
}

// FuncHasOnLidEvent returns the function "WEBEXT.os.events.onLidEvent.hasListener".
func FuncHasOnLidEvent() (fn js.Func[func(callback js.Func[func(event_info *LidEventInfo)]) bool]) {
	bindings.FuncHasOnLidEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnLidEvent calls the function "WEBEXT.os.events.onLidEvent.hasListener" directly.
func HasOnLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret bool) {
	bindings.CallHasOnLidEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLidEvent calls the function "WEBEXT.os.events.onLidEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLidEvent(callback js.Func[func(event_info *LidEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLidEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPowerEventEventCallbackFunc func(this js.Ref, event_info *PowerEventInfo) js.Ref

func (fn OnPowerEventEventCallbackFunc) Register() js.Func[func(event_info *PowerEventInfo)] {
	return js.RegisterCallback[func(event_info *PowerEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPowerEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PowerEventInfo
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

type OnPowerEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *PowerEventInfo) js.Ref
	Arg T
}

func (cb *OnPowerEventEventCallback[T]) Register() js.Func[func(event_info *PowerEventInfo)] {
	return js.RegisterCallback[func(event_info *PowerEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPowerEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PowerEventInfo
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

// HasFuncOnPowerEvent returns true if the function "WEBEXT.os.events.onPowerEvent.addListener" exists.
func HasFuncOnPowerEvent() bool {
	return js.True == bindings.HasFuncOnPowerEvent()
}

// FuncOnPowerEvent returns the function "WEBEXT.os.events.onPowerEvent.addListener".
func FuncOnPowerEvent() (fn js.Func[func(callback js.Func[func(event_info *PowerEventInfo)])]) {
	bindings.FuncOnPowerEvent(
		js.Pointer(&fn),
	)
	return
}

// OnPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.addListener" directly.
func OnPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret js.Void) {
	bindings.CallOnPowerEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPowerEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPowerEvent returns true if the function "WEBEXT.os.events.onPowerEvent.removeListener" exists.
func HasFuncOffPowerEvent() bool {
	return js.True == bindings.HasFuncOffPowerEvent()
}

// FuncOffPowerEvent returns the function "WEBEXT.os.events.onPowerEvent.removeListener".
func FuncOffPowerEvent() (fn js.Func[func(callback js.Func[func(event_info *PowerEventInfo)])]) {
	bindings.FuncOffPowerEvent(
		js.Pointer(&fn),
	)
	return
}

// OffPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.removeListener" directly.
func OffPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret js.Void) {
	bindings.CallOffPowerEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPowerEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPowerEvent returns true if the function "WEBEXT.os.events.onPowerEvent.hasListener" exists.
func HasFuncHasOnPowerEvent() bool {
	return js.True == bindings.HasFuncHasOnPowerEvent()
}

// FuncHasOnPowerEvent returns the function "WEBEXT.os.events.onPowerEvent.hasListener".
func FuncHasOnPowerEvent() (fn js.Func[func(callback js.Func[func(event_info *PowerEventInfo)]) bool]) {
	bindings.FuncHasOnPowerEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.hasListener" directly.
func HasOnPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret bool) {
	bindings.CallHasOnPowerEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPowerEvent calls the function "WEBEXT.os.events.onPowerEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPowerEvent(callback js.Func[func(event_info *PowerEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPowerEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSdCardEventEventCallbackFunc func(this js.Ref, event_info *SdCardEventInfo) js.Ref

func (fn OnSdCardEventEventCallbackFunc) Register() js.Func[func(event_info *SdCardEventInfo)] {
	return js.RegisterCallback[func(event_info *SdCardEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSdCardEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SdCardEventInfo
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

type OnSdCardEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *SdCardEventInfo) js.Ref
	Arg T
}

func (cb *OnSdCardEventEventCallback[T]) Register() js.Func[func(event_info *SdCardEventInfo)] {
	return js.RegisterCallback[func(event_info *SdCardEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSdCardEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SdCardEventInfo
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

// HasFuncOnSdCardEvent returns true if the function "WEBEXT.os.events.onSdCardEvent.addListener" exists.
func HasFuncOnSdCardEvent() bool {
	return js.True == bindings.HasFuncOnSdCardEvent()
}

// FuncOnSdCardEvent returns the function "WEBEXT.os.events.onSdCardEvent.addListener".
func FuncOnSdCardEvent() (fn js.Func[func(callback js.Func[func(event_info *SdCardEventInfo)])]) {
	bindings.FuncOnSdCardEvent(
		js.Pointer(&fn),
	)
	return
}

// OnSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.addListener" directly.
func OnSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret js.Void) {
	bindings.CallOnSdCardEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSdCardEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSdCardEvent returns true if the function "WEBEXT.os.events.onSdCardEvent.removeListener" exists.
func HasFuncOffSdCardEvent() bool {
	return js.True == bindings.HasFuncOffSdCardEvent()
}

// FuncOffSdCardEvent returns the function "WEBEXT.os.events.onSdCardEvent.removeListener".
func FuncOffSdCardEvent() (fn js.Func[func(callback js.Func[func(event_info *SdCardEventInfo)])]) {
	bindings.FuncOffSdCardEvent(
		js.Pointer(&fn),
	)
	return
}

// OffSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.removeListener" directly.
func OffSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret js.Void) {
	bindings.CallOffSdCardEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSdCardEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSdCardEvent returns true if the function "WEBEXT.os.events.onSdCardEvent.hasListener" exists.
func HasFuncHasOnSdCardEvent() bool {
	return js.True == bindings.HasFuncHasOnSdCardEvent()
}

// FuncHasOnSdCardEvent returns the function "WEBEXT.os.events.onSdCardEvent.hasListener".
func FuncHasOnSdCardEvent() (fn js.Func[func(callback js.Func[func(event_info *SdCardEventInfo)]) bool]) {
	bindings.FuncHasOnSdCardEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.hasListener" directly.
func HasOnSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret bool) {
	bindings.CallHasOnSdCardEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSdCardEvent calls the function "WEBEXT.os.events.onSdCardEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSdCardEvent(callback js.Func[func(event_info *SdCardEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSdCardEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStylusConnectedEventEventCallbackFunc func(this js.Ref, event_info *StylusConnectedEventInfo) js.Ref

func (fn OnStylusConnectedEventEventCallbackFunc) Register() js.Func[func(event_info *StylusConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusConnectedEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStylusConnectedEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusConnectedEventInfo
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

type OnStylusConnectedEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *StylusConnectedEventInfo) js.Ref
	Arg T
}

func (cb *OnStylusConnectedEventEventCallback[T]) Register() js.Func[func(event_info *StylusConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusConnectedEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStylusConnectedEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusConnectedEventInfo
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

// HasFuncOnStylusConnectedEvent returns true if the function "WEBEXT.os.events.onStylusConnectedEvent.addListener" exists.
func HasFuncOnStylusConnectedEvent() bool {
	return js.True == bindings.HasFuncOnStylusConnectedEvent()
}

// FuncOnStylusConnectedEvent returns the function "WEBEXT.os.events.onStylusConnectedEvent.addListener".
func FuncOnStylusConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusConnectedEventInfo)])]) {
	bindings.FuncOnStylusConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OnStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.addListener" directly.
func OnStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOnStylusConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStylusConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStylusConnectedEvent returns true if the function "WEBEXT.os.events.onStylusConnectedEvent.removeListener" exists.
func HasFuncOffStylusConnectedEvent() bool {
	return js.True == bindings.HasFuncOffStylusConnectedEvent()
}

// FuncOffStylusConnectedEvent returns the function "WEBEXT.os.events.onStylusConnectedEvent.removeListener".
func FuncOffStylusConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusConnectedEventInfo)])]) {
	bindings.FuncOffStylusConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OffStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.removeListener" directly.
func OffStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOffStylusConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStylusConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStylusConnectedEvent returns true if the function "WEBEXT.os.events.onStylusConnectedEvent.hasListener" exists.
func HasFuncHasOnStylusConnectedEvent() bool {
	return js.True == bindings.HasFuncHasOnStylusConnectedEvent()
}

// FuncHasOnStylusConnectedEvent returns the function "WEBEXT.os.events.onStylusConnectedEvent.hasListener".
func FuncHasOnStylusConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusConnectedEventInfo)]) bool]) {
	bindings.FuncHasOnStylusConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.hasListener" directly.
func HasOnStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret bool) {
	bindings.CallHasOnStylusConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStylusConnectedEvent calls the function "WEBEXT.os.events.onStylusConnectedEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStylusConnectedEvent(callback js.Func[func(event_info *StylusConnectedEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStylusConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStylusGarageEventEventCallbackFunc func(this js.Ref, event_info *StylusGarageEventInfo) js.Ref

func (fn OnStylusGarageEventEventCallbackFunc) Register() js.Func[func(event_info *StylusGarageEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusGarageEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStylusGarageEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusGarageEventInfo
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

type OnStylusGarageEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *StylusGarageEventInfo) js.Ref
	Arg T
}

func (cb *OnStylusGarageEventEventCallback[T]) Register() js.Func[func(event_info *StylusGarageEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusGarageEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStylusGarageEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusGarageEventInfo
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

// HasFuncOnStylusGarageEvent returns true if the function "WEBEXT.os.events.onStylusGarageEvent.addListener" exists.
func HasFuncOnStylusGarageEvent() bool {
	return js.True == bindings.HasFuncOnStylusGarageEvent()
}

// FuncOnStylusGarageEvent returns the function "WEBEXT.os.events.onStylusGarageEvent.addListener".
func FuncOnStylusGarageEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusGarageEventInfo)])]) {
	bindings.FuncOnStylusGarageEvent(
		js.Pointer(&fn),
	)
	return
}

// OnStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.addListener" directly.
func OnStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret js.Void) {
	bindings.CallOnStylusGarageEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStylusGarageEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStylusGarageEvent returns true if the function "WEBEXT.os.events.onStylusGarageEvent.removeListener" exists.
func HasFuncOffStylusGarageEvent() bool {
	return js.True == bindings.HasFuncOffStylusGarageEvent()
}

// FuncOffStylusGarageEvent returns the function "WEBEXT.os.events.onStylusGarageEvent.removeListener".
func FuncOffStylusGarageEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusGarageEventInfo)])]) {
	bindings.FuncOffStylusGarageEvent(
		js.Pointer(&fn),
	)
	return
}

// OffStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.removeListener" directly.
func OffStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret js.Void) {
	bindings.CallOffStylusGarageEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStylusGarageEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStylusGarageEvent returns true if the function "WEBEXT.os.events.onStylusGarageEvent.hasListener" exists.
func HasFuncHasOnStylusGarageEvent() bool {
	return js.True == bindings.HasFuncHasOnStylusGarageEvent()
}

// FuncHasOnStylusGarageEvent returns the function "WEBEXT.os.events.onStylusGarageEvent.hasListener".
func FuncHasOnStylusGarageEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusGarageEventInfo)]) bool]) {
	bindings.FuncHasOnStylusGarageEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.hasListener" directly.
func HasOnStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret bool) {
	bindings.CallHasOnStylusGarageEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStylusGarageEvent calls the function "WEBEXT.os.events.onStylusGarageEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStylusGarageEvent(callback js.Func[func(event_info *StylusGarageEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStylusGarageEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStylusTouchEventEventCallbackFunc func(this js.Ref, event_info *StylusTouchEventInfo) js.Ref

func (fn OnStylusTouchEventEventCallbackFunc) Register() js.Func[func(event_info *StylusTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusTouchEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStylusTouchEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusTouchEventInfo
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

type OnStylusTouchEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *StylusTouchEventInfo) js.Ref
	Arg T
}

func (cb *OnStylusTouchEventEventCallback[T]) Register() js.Func[func(event_info *StylusTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *StylusTouchEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStylusTouchEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StylusTouchEventInfo
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

// HasFuncOnStylusTouchEvent returns true if the function "WEBEXT.os.events.onStylusTouchEvent.addListener" exists.
func HasFuncOnStylusTouchEvent() bool {
	return js.True == bindings.HasFuncOnStylusTouchEvent()
}

// FuncOnStylusTouchEvent returns the function "WEBEXT.os.events.onStylusTouchEvent.addListener".
func FuncOnStylusTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusTouchEventInfo)])]) {
	bindings.FuncOnStylusTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OnStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.addListener" directly.
func OnStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret js.Void) {
	bindings.CallOnStylusTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStylusTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStylusTouchEvent returns true if the function "WEBEXT.os.events.onStylusTouchEvent.removeListener" exists.
func HasFuncOffStylusTouchEvent() bool {
	return js.True == bindings.HasFuncOffStylusTouchEvent()
}

// FuncOffStylusTouchEvent returns the function "WEBEXT.os.events.onStylusTouchEvent.removeListener".
func FuncOffStylusTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusTouchEventInfo)])]) {
	bindings.FuncOffStylusTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OffStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.removeListener" directly.
func OffStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret js.Void) {
	bindings.CallOffStylusTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStylusTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStylusTouchEvent returns true if the function "WEBEXT.os.events.onStylusTouchEvent.hasListener" exists.
func HasFuncHasOnStylusTouchEvent() bool {
	return js.True == bindings.HasFuncHasOnStylusTouchEvent()
}

// FuncHasOnStylusTouchEvent returns the function "WEBEXT.os.events.onStylusTouchEvent.hasListener".
func FuncHasOnStylusTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *StylusTouchEventInfo)]) bool]) {
	bindings.FuncHasOnStylusTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.hasListener" directly.
func HasOnStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret bool) {
	bindings.CallHasOnStylusTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStylusTouchEvent calls the function "WEBEXT.os.events.onStylusTouchEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStylusTouchEvent(callback js.Func[func(event_info *StylusTouchEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStylusTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchpadButtonEventEventCallbackFunc func(this js.Ref, event_info *TouchpadButtonEventInfo) js.Ref

func (fn OnTouchpadButtonEventEventCallbackFunc) Register() js.Func[func(event_info *TouchpadButtonEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadButtonEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchpadButtonEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadButtonEventInfo
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

type OnTouchpadButtonEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *TouchpadButtonEventInfo) js.Ref
	Arg T
}

func (cb *OnTouchpadButtonEventEventCallback[T]) Register() js.Func[func(event_info *TouchpadButtonEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadButtonEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchpadButtonEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadButtonEventInfo
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

// HasFuncOnTouchpadButtonEvent returns true if the function "WEBEXT.os.events.onTouchpadButtonEvent.addListener" exists.
func HasFuncOnTouchpadButtonEvent() bool {
	return js.True == bindings.HasFuncOnTouchpadButtonEvent()
}

// FuncOnTouchpadButtonEvent returns the function "WEBEXT.os.events.onTouchpadButtonEvent.addListener".
func FuncOnTouchpadButtonEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadButtonEventInfo)])]) {
	bindings.FuncOnTouchpadButtonEvent(
		js.Pointer(&fn),
	)
	return
}

// OnTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.addListener" directly.
func OnTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret js.Void) {
	bindings.CallOnTouchpadButtonEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouchpadButtonEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouchpadButtonEvent returns true if the function "WEBEXT.os.events.onTouchpadButtonEvent.removeListener" exists.
func HasFuncOffTouchpadButtonEvent() bool {
	return js.True == bindings.HasFuncOffTouchpadButtonEvent()
}

// FuncOffTouchpadButtonEvent returns the function "WEBEXT.os.events.onTouchpadButtonEvent.removeListener".
func FuncOffTouchpadButtonEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadButtonEventInfo)])]) {
	bindings.FuncOffTouchpadButtonEvent(
		js.Pointer(&fn),
	)
	return
}

// OffTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.removeListener" directly.
func OffTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret js.Void) {
	bindings.CallOffTouchpadButtonEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouchpadButtonEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouchpadButtonEvent returns true if the function "WEBEXT.os.events.onTouchpadButtonEvent.hasListener" exists.
func HasFuncHasOnTouchpadButtonEvent() bool {
	return js.True == bindings.HasFuncHasOnTouchpadButtonEvent()
}

// FuncHasOnTouchpadButtonEvent returns the function "WEBEXT.os.events.onTouchpadButtonEvent.hasListener".
func FuncHasOnTouchpadButtonEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) bool]) {
	bindings.FuncHasOnTouchpadButtonEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.hasListener" directly.
func HasOnTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret bool) {
	bindings.CallHasOnTouchpadButtonEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouchpadButtonEvent calls the function "WEBEXT.os.events.onTouchpadButtonEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouchpadButtonEvent(callback js.Func[func(event_info *TouchpadButtonEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouchpadButtonEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchpadConnectedEventEventCallbackFunc func(this js.Ref, event_info *TouchpadConnectedEventInfo) js.Ref

func (fn OnTouchpadConnectedEventEventCallbackFunc) Register() js.Func[func(event_info *TouchpadConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadConnectedEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchpadConnectedEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadConnectedEventInfo
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

type OnTouchpadConnectedEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *TouchpadConnectedEventInfo) js.Ref
	Arg T
}

func (cb *OnTouchpadConnectedEventEventCallback[T]) Register() js.Func[func(event_info *TouchpadConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadConnectedEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchpadConnectedEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadConnectedEventInfo
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

// HasFuncOnTouchpadConnectedEvent returns true if the function "WEBEXT.os.events.onTouchpadConnectedEvent.addListener" exists.
func HasFuncOnTouchpadConnectedEvent() bool {
	return js.True == bindings.HasFuncOnTouchpadConnectedEvent()
}

// FuncOnTouchpadConnectedEvent returns the function "WEBEXT.os.events.onTouchpadConnectedEvent.addListener".
func FuncOnTouchpadConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadConnectedEventInfo)])]) {
	bindings.FuncOnTouchpadConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OnTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.addListener" directly.
func OnTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOnTouchpadConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouchpadConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouchpadConnectedEvent returns true if the function "WEBEXT.os.events.onTouchpadConnectedEvent.removeListener" exists.
func HasFuncOffTouchpadConnectedEvent() bool {
	return js.True == bindings.HasFuncOffTouchpadConnectedEvent()
}

// FuncOffTouchpadConnectedEvent returns the function "WEBEXT.os.events.onTouchpadConnectedEvent.removeListener".
func FuncOffTouchpadConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadConnectedEventInfo)])]) {
	bindings.FuncOffTouchpadConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OffTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.removeListener" directly.
func OffTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOffTouchpadConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouchpadConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouchpadConnectedEvent returns true if the function "WEBEXT.os.events.onTouchpadConnectedEvent.hasListener" exists.
func HasFuncHasOnTouchpadConnectedEvent() bool {
	return js.True == bindings.HasFuncHasOnTouchpadConnectedEvent()
}

// FuncHasOnTouchpadConnectedEvent returns the function "WEBEXT.os.events.onTouchpadConnectedEvent.hasListener".
func FuncHasOnTouchpadConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) bool]) {
	bindings.FuncHasOnTouchpadConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.hasListener" directly.
func HasOnTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret bool) {
	bindings.CallHasOnTouchpadConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouchpadConnectedEvent calls the function "WEBEXT.os.events.onTouchpadConnectedEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouchpadConnectedEvent(callback js.Func[func(event_info *TouchpadConnectedEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouchpadConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchpadTouchEventEventCallbackFunc func(this js.Ref, event_info *TouchpadTouchEventInfo) js.Ref

func (fn OnTouchpadTouchEventEventCallbackFunc) Register() js.Func[func(event_info *TouchpadTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadTouchEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchpadTouchEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadTouchEventInfo
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

type OnTouchpadTouchEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *TouchpadTouchEventInfo) js.Ref
	Arg T
}

func (cb *OnTouchpadTouchEventEventCallback[T]) Register() js.Func[func(event_info *TouchpadTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchpadTouchEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchpadTouchEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchpadTouchEventInfo
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

// HasFuncOnTouchpadTouchEvent returns true if the function "WEBEXT.os.events.onTouchpadTouchEvent.addListener" exists.
func HasFuncOnTouchpadTouchEvent() bool {
	return js.True == bindings.HasFuncOnTouchpadTouchEvent()
}

// FuncOnTouchpadTouchEvent returns the function "WEBEXT.os.events.onTouchpadTouchEvent.addListener".
func FuncOnTouchpadTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadTouchEventInfo)])]) {
	bindings.FuncOnTouchpadTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OnTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.addListener" directly.
func OnTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret js.Void) {
	bindings.CallOnTouchpadTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouchpadTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouchpadTouchEvent returns true if the function "WEBEXT.os.events.onTouchpadTouchEvent.removeListener" exists.
func HasFuncOffTouchpadTouchEvent() bool {
	return js.True == bindings.HasFuncOffTouchpadTouchEvent()
}

// FuncOffTouchpadTouchEvent returns the function "WEBEXT.os.events.onTouchpadTouchEvent.removeListener".
func FuncOffTouchpadTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadTouchEventInfo)])]) {
	bindings.FuncOffTouchpadTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OffTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.removeListener" directly.
func OffTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret js.Void) {
	bindings.CallOffTouchpadTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouchpadTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouchpadTouchEvent returns true if the function "WEBEXT.os.events.onTouchpadTouchEvent.hasListener" exists.
func HasFuncHasOnTouchpadTouchEvent() bool {
	return js.True == bindings.HasFuncHasOnTouchpadTouchEvent()
}

// FuncHasOnTouchpadTouchEvent returns the function "WEBEXT.os.events.onTouchpadTouchEvent.hasListener".
func FuncHasOnTouchpadTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) bool]) {
	bindings.FuncHasOnTouchpadTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.hasListener" directly.
func HasOnTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret bool) {
	bindings.CallHasOnTouchpadTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouchpadTouchEvent calls the function "WEBEXT.os.events.onTouchpadTouchEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouchpadTouchEvent(callback js.Func[func(event_info *TouchpadTouchEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouchpadTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchscreenConnectedEventEventCallbackFunc func(this js.Ref, event_info *TouchscreenConnectedEventInfo) js.Ref

func (fn OnTouchscreenConnectedEventEventCallbackFunc) Register() js.Func[func(event_info *TouchscreenConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchscreenConnectedEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchscreenConnectedEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchscreenConnectedEventInfo
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

type OnTouchscreenConnectedEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *TouchscreenConnectedEventInfo) js.Ref
	Arg T
}

func (cb *OnTouchscreenConnectedEventEventCallback[T]) Register() js.Func[func(event_info *TouchscreenConnectedEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchscreenConnectedEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchscreenConnectedEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchscreenConnectedEventInfo
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

// HasFuncOnTouchscreenConnectedEvent returns true if the function "WEBEXT.os.events.onTouchscreenConnectedEvent.addListener" exists.
func HasFuncOnTouchscreenConnectedEvent() bool {
	return js.True == bindings.HasFuncOnTouchscreenConnectedEvent()
}

// FuncOnTouchscreenConnectedEvent returns the function "WEBEXT.os.events.onTouchscreenConnectedEvent.addListener".
func FuncOnTouchscreenConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)])]) {
	bindings.FuncOnTouchscreenConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OnTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.addListener" directly.
func OnTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOnTouchscreenConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouchscreenConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouchscreenConnectedEvent returns true if the function "WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener" exists.
func HasFuncOffTouchscreenConnectedEvent() bool {
	return js.True == bindings.HasFuncOffTouchscreenConnectedEvent()
}

// FuncOffTouchscreenConnectedEvent returns the function "WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener".
func FuncOffTouchscreenConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)])]) {
	bindings.FuncOffTouchscreenConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// OffTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener" directly.
func OffTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret js.Void) {
	bindings.CallOffTouchscreenConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouchscreenConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouchscreenConnectedEvent returns true if the function "WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener" exists.
func HasFuncHasOnTouchscreenConnectedEvent() bool {
	return js.True == bindings.HasFuncHasOnTouchscreenConnectedEvent()
}

// FuncHasOnTouchscreenConnectedEvent returns the function "WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener".
func FuncHasOnTouchscreenConnectedEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) bool]) {
	bindings.FuncHasOnTouchscreenConnectedEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener" directly.
func HasOnTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret bool) {
	bindings.CallHasOnTouchscreenConnectedEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouchscreenConnectedEvent calls the function "WEBEXT.os.events.onTouchscreenConnectedEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouchscreenConnectedEvent(callback js.Func[func(event_info *TouchscreenConnectedEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouchscreenConnectedEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnTouchscreenTouchEventEventCallbackFunc func(this js.Ref, event_info *TouchscreenTouchEventInfo) js.Ref

func (fn OnTouchscreenTouchEventEventCallbackFunc) Register() js.Func[func(event_info *TouchscreenTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchscreenTouchEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTouchscreenTouchEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchscreenTouchEventInfo
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

type OnTouchscreenTouchEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *TouchscreenTouchEventInfo) js.Ref
	Arg T
}

func (cb *OnTouchscreenTouchEventEventCallback[T]) Register() js.Func[func(event_info *TouchscreenTouchEventInfo)] {
	return js.RegisterCallback[func(event_info *TouchscreenTouchEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTouchscreenTouchEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TouchscreenTouchEventInfo
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

// HasFuncOnTouchscreenTouchEvent returns true if the function "WEBEXT.os.events.onTouchscreenTouchEvent.addListener" exists.
func HasFuncOnTouchscreenTouchEvent() bool {
	return js.True == bindings.HasFuncOnTouchscreenTouchEvent()
}

// FuncOnTouchscreenTouchEvent returns the function "WEBEXT.os.events.onTouchscreenTouchEvent.addListener".
func FuncOnTouchscreenTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenTouchEventInfo)])]) {
	bindings.FuncOnTouchscreenTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OnTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.addListener" directly.
func OnTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret js.Void) {
	bindings.CallOnTouchscreenTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTouchscreenTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTouchscreenTouchEvent returns true if the function "WEBEXT.os.events.onTouchscreenTouchEvent.removeListener" exists.
func HasFuncOffTouchscreenTouchEvent() bool {
	return js.True == bindings.HasFuncOffTouchscreenTouchEvent()
}

// FuncOffTouchscreenTouchEvent returns the function "WEBEXT.os.events.onTouchscreenTouchEvent.removeListener".
func FuncOffTouchscreenTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenTouchEventInfo)])]) {
	bindings.FuncOffTouchscreenTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// OffTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.removeListener" directly.
func OffTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret js.Void) {
	bindings.CallOffTouchscreenTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTouchscreenTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTouchscreenTouchEvent returns true if the function "WEBEXT.os.events.onTouchscreenTouchEvent.hasListener" exists.
func HasFuncHasOnTouchscreenTouchEvent() bool {
	return js.True == bindings.HasFuncHasOnTouchscreenTouchEvent()
}

// FuncHasOnTouchscreenTouchEvent returns the function "WEBEXT.os.events.onTouchscreenTouchEvent.hasListener".
func FuncHasOnTouchscreenTouchEvent() (fn js.Func[func(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) bool]) {
	bindings.FuncHasOnTouchscreenTouchEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.hasListener" directly.
func HasOnTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret bool) {
	bindings.CallHasOnTouchscreenTouchEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTouchscreenTouchEvent calls the function "WEBEXT.os.events.onTouchscreenTouchEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTouchscreenTouchEvent(callback js.Func[func(event_info *TouchscreenTouchEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTouchscreenTouchEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUsbEventEventCallbackFunc func(this js.Ref, event_info *UsbEventInfo) js.Ref

func (fn OnUsbEventEventCallbackFunc) Register() js.Func[func(event_info *UsbEventInfo)] {
	return js.RegisterCallback[func(event_info *UsbEventInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUsbEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UsbEventInfo
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

type OnUsbEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event_info *UsbEventInfo) js.Ref
	Arg T
}

func (cb *OnUsbEventEventCallback[T]) Register() js.Func[func(event_info *UsbEventInfo)] {
	return js.RegisterCallback[func(event_info *UsbEventInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUsbEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UsbEventInfo
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

// HasFuncOnUsbEvent returns true if the function "WEBEXT.os.events.onUsbEvent.addListener" exists.
func HasFuncOnUsbEvent() bool {
	return js.True == bindings.HasFuncOnUsbEvent()
}

// FuncOnUsbEvent returns the function "WEBEXT.os.events.onUsbEvent.addListener".
func FuncOnUsbEvent() (fn js.Func[func(callback js.Func[func(event_info *UsbEventInfo)])]) {
	bindings.FuncOnUsbEvent(
		js.Pointer(&fn),
	)
	return
}

// OnUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.addListener" directly.
func OnUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret js.Void) {
	bindings.CallOnUsbEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUsbEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUsbEvent returns true if the function "WEBEXT.os.events.onUsbEvent.removeListener" exists.
func HasFuncOffUsbEvent() bool {
	return js.True == bindings.HasFuncOffUsbEvent()
}

// FuncOffUsbEvent returns the function "WEBEXT.os.events.onUsbEvent.removeListener".
func FuncOffUsbEvent() (fn js.Func[func(callback js.Func[func(event_info *UsbEventInfo)])]) {
	bindings.FuncOffUsbEvent(
		js.Pointer(&fn),
	)
	return
}

// OffUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.removeListener" directly.
func OffUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret js.Void) {
	bindings.CallOffUsbEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUsbEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUsbEvent returns true if the function "WEBEXT.os.events.onUsbEvent.hasListener" exists.
func HasFuncHasOnUsbEvent() bool {
	return js.True == bindings.HasFuncHasOnUsbEvent()
}

// FuncHasOnUsbEvent returns the function "WEBEXT.os.events.onUsbEvent.hasListener".
func FuncHasOnUsbEvent() (fn js.Func[func(callback js.Func[func(event_info *UsbEventInfo)]) bool]) {
	bindings.FuncHasOnUsbEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.hasListener" directly.
func HasOnUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret bool) {
	bindings.CallHasOnUsbEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUsbEvent calls the function "WEBEXT.os.events.onUsbEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUsbEvent(callback js.Func[func(event_info *UsbEventInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUsbEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncStartCapturingEvents returns true if the function "WEBEXT.os.events.startCapturingEvents" exists.
func HasFuncStartCapturingEvents() bool {
	return js.True == bindings.HasFuncStartCapturingEvents()
}

// FuncStartCapturingEvents returns the function "WEBEXT.os.events.startCapturingEvents".
func FuncStartCapturingEvents() (fn js.Func[func(category EventCategory) js.Promise[js.Void]]) {
	bindings.FuncStartCapturingEvents(
		js.Pointer(&fn),
	)
	return
}

// StartCapturingEvents calls the function "WEBEXT.os.events.startCapturingEvents" directly.
func StartCapturingEvents(category EventCategory) (ret js.Promise[js.Void]) {
	bindings.CallStartCapturingEvents(
		js.Pointer(&ret),
		uint32(category),
	)

	return
}

// TryStartCapturingEvents calls the function "WEBEXT.os.events.startCapturingEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartCapturingEvents(category EventCategory) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartCapturingEvents(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(category),
	)

	return
}

// HasFuncStopCapturingEvents returns true if the function "WEBEXT.os.events.stopCapturingEvents" exists.
func HasFuncStopCapturingEvents() bool {
	return js.True == bindings.HasFuncStopCapturingEvents()
}

// FuncStopCapturingEvents returns the function "WEBEXT.os.events.stopCapturingEvents".
func FuncStopCapturingEvents() (fn js.Func[func(category EventCategory) js.Promise[js.Void]]) {
	bindings.FuncStopCapturingEvents(
		js.Pointer(&fn),
	)
	return
}

// StopCapturingEvents calls the function "WEBEXT.os.events.stopCapturingEvents" directly.
func StopCapturingEvents(category EventCategory) (ret js.Promise[js.Void]) {
	bindings.CallStopCapturingEvents(
		js.Pointer(&ret),
		uint32(category),
	)

	return
}

// TryStopCapturingEvents calls the function "WEBEXT.os.events.stopCapturingEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopCapturingEvents(category EventCategory) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopCapturingEvents(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(category),
	)

	return
}

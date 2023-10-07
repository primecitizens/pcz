// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package ime

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/input/ime/bindings"
)

type AssistiveWindowButton uint32

const (
	_ AssistiveWindowButton = iota

	AssistiveWindowButton_UNDO
	AssistiveWindowButton_ADD_TO_DICTIONARY
)

func (AssistiveWindowButton) FromRef(str js.Ref) AssistiveWindowButton {
	return AssistiveWindowButton(bindings.ConstOfAssistiveWindowButton(str))
}

func (x AssistiveWindowButton) String() (string, bool) {
	switch x {
	case AssistiveWindowButton_UNDO:
		return "undo", true
	case AssistiveWindowButton_ADD_TO_DICTIONARY:
		return "addToDictionary", true
	default:
		return "", false
	}
}

type AssistiveWindowType uint32

const (
	_ AssistiveWindowType = iota

	AssistiveWindowType_UNDO
)

func (AssistiveWindowType) FromRef(str js.Ref) AssistiveWindowType {
	return AssistiveWindowType(bindings.ConstOfAssistiveWindowType(str))
}

func (x AssistiveWindowType) String() (string, bool) {
	switch x {
	case AssistiveWindowType_UNDO:
		return "undo", true
	default:
		return "", false
	}
}

type AssistiveWindowProperties struct {
	// AnnounceString is "AssistiveWindowProperties.announceString"
	//
	// Optional
	AnnounceString js.String
	// Type is "AssistiveWindowProperties.type"
	//
	// Required
	Type AssistiveWindowType
	// Visible is "AssistiveWindowProperties.visible"
	//
	// Required
	Visible bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssistiveWindowProperties with all fields set.
func (p AssistiveWindowProperties) FromRef(ref js.Ref) AssistiveWindowProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssistiveWindowProperties in the application heap.
func (p AssistiveWindowProperties) New() js.Ref {
	return bindings.AssistiveWindowPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssistiveWindowProperties) UpdateFrom(ref js.Ref) {
	bindings.AssistiveWindowPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssistiveWindowProperties) Update(ref js.Ref) {
	bindings.AssistiveWindowPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssistiveWindowProperties) FreeMembers(recursive bool) {
	js.Free(
		p.AnnounceString.Ref(),
	)
	p.AnnounceString = p.AnnounceString.FromRef(js.Undefined)
}

type AutoCapitalizeType uint32

const (
	_ AutoCapitalizeType = iota

	AutoCapitalizeType_CHARACTERS
	AutoCapitalizeType_WORDS
	AutoCapitalizeType_SENTENCES
)

func (AutoCapitalizeType) FromRef(str js.Ref) AutoCapitalizeType {
	return AutoCapitalizeType(bindings.ConstOfAutoCapitalizeType(str))
}

func (x AutoCapitalizeType) String() (string, bool) {
	switch x {
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

type ClearCompositionArgParameters struct {
	// ContextID is "ClearCompositionArgParameters.contextID"
	//
	// Required
	ContextID int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearCompositionArgParameters with all fields set.
func (p ClearCompositionArgParameters) FromRef(ref js.Ref) ClearCompositionArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearCompositionArgParameters in the application heap.
func (p ClearCompositionArgParameters) New() js.Ref {
	return bindings.ClearCompositionArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearCompositionArgParameters) UpdateFrom(ref js.Ref) {
	bindings.ClearCompositionArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearCompositionArgParameters) Update(ref js.Ref) {
	bindings.ClearCompositionArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearCompositionArgParameters) FreeMembers(recursive bool) {
}

type CommitTextArgParameters struct {
	// ContextID is "CommitTextArgParameters.contextID"
	//
	// Required
	ContextID int64
	// Text is "CommitTextArgParameters.text"
	//
	// Required
	Text js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CommitTextArgParameters with all fields set.
func (p CommitTextArgParameters) FromRef(ref js.Ref) CommitTextArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CommitTextArgParameters in the application heap.
func (p CommitTextArgParameters) New() js.Ref {
	return bindings.CommitTextArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CommitTextArgParameters) UpdateFrom(ref js.Ref) {
	bindings.CommitTextArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CommitTextArgParameters) Update(ref js.Ref) {
	bindings.CommitTextArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CommitTextArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

type DeleteSurroundingTextArgParameters struct {
	// ContextID is "DeleteSurroundingTextArgParameters.contextID"
	//
	// Required
	ContextID int64
	// EngineID is "DeleteSurroundingTextArgParameters.engineID"
	//
	// Required
	EngineID js.String
	// Length is "DeleteSurroundingTextArgParameters.length"
	//
	// Required
	Length int64
	// Offset is "DeleteSurroundingTextArgParameters.offset"
	//
	// Required
	Offset int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteSurroundingTextArgParameters with all fields set.
func (p DeleteSurroundingTextArgParameters) FromRef(ref js.Ref) DeleteSurroundingTextArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteSurroundingTextArgParameters in the application heap.
func (p DeleteSurroundingTextArgParameters) New() js.Ref {
	return bindings.DeleteSurroundingTextArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteSurroundingTextArgParameters) UpdateFrom(ref js.Ref) {
	bindings.DeleteSurroundingTextArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteSurroundingTextArgParameters) Update(ref js.Ref) {
	bindings.DeleteSurroundingTextArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteSurroundingTextArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.EngineID.Ref(),
	)
	p.EngineID = p.EngineID.FromRef(js.Undefined)
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
}

type KeyboardEventType uint32

const (
	_ KeyboardEventType = iota

	KeyboardEventType_KEYUP
	KeyboardEventType_KEYDOWN
)

func (KeyboardEventType) FromRef(str js.Ref) KeyboardEventType {
	return KeyboardEventType(bindings.ConstOfKeyboardEventType(str))
}

func (x KeyboardEventType) String() (string, bool) {
	switch x {
	case KeyboardEventType_KEYUP:
		return "keyup", true
	case KeyboardEventType_KEYDOWN:
		return "keydown", true
	default:
		return "", false
	}
}

type KeyboardEvent struct {
	// AltKey is "KeyboardEvent.altKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// AltgrKey is "KeyboardEvent.altgrKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_AltgrKey MUST be set to true to make this field effective.
	AltgrKey bool
	// CapsLock is "KeyboardEvent.capsLock"
	//
	// Optional
	//
	// NOTE: FFI_USE_CapsLock MUST be set to true to make this field effective.
	CapsLock bool
	// Code is "KeyboardEvent.code"
	//
	// Required
	Code js.String
	// CtrlKey is "KeyboardEvent.ctrlKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ExtensionId is "KeyboardEvent.extensionId"
	//
	// Optional
	ExtensionId js.String
	// Key is "KeyboardEvent.key"
	//
	// Required
	Key js.String
	// KeyCode is "KeyboardEvent.keyCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_KeyCode MUST be set to true to make this field effective.
	KeyCode int64
	// RequestId is "KeyboardEvent.requestId"
	//
	// Optional
	RequestId js.String
	// ShiftKey is "KeyboardEvent.shiftKey"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// Type is "KeyboardEvent.type"
	//
	// Required
	Type KeyboardEventType

	FFI_USE_AltKey   bool // for AltKey.
	FFI_USE_AltgrKey bool // for AltgrKey.
	FFI_USE_CapsLock bool // for CapsLock.
	FFI_USE_CtrlKey  bool // for CtrlKey.
	FFI_USE_KeyCode  bool // for KeyCode.
	FFI_USE_ShiftKey bool // for ShiftKey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyboardEvent with all fields set.
func (p KeyboardEvent) FromRef(ref js.Ref) KeyboardEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyboardEvent in the application heap.
func (p KeyboardEvent) New() js.Ref {
	return bindings.KeyboardEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KeyboardEvent) UpdateFrom(ref js.Ref) {
	bindings.KeyboardEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyboardEvent) Update(ref js.Ref) {
	bindings.KeyboardEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyboardEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.ExtensionId.Ref(),
		p.Key.Ref(),
		p.RequestId.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Key = p.Key.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
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

type MenuParameters struct {
	// EngineID is "MenuParameters.engineID"
	//
	// Required
	EngineID js.String
	// Items is "MenuParameters.items"
	//
	// Required
	Items js.Array[MenuItem]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MenuParameters with all fields set.
func (p MenuParameters) FromRef(ref js.Ref) MenuParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MenuParameters in the application heap.
func (p MenuParameters) New() js.Ref {
	return bindings.MenuParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MenuParameters) UpdateFrom(ref js.Ref) {
	bindings.MenuParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MenuParameters) Update(ref js.Ref) {
	bindings.MenuParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MenuParameters) FreeMembers(recursive bool) {
	js.Free(
		p.EngineID.Ref(),
		p.Items.Ref(),
	)
	p.EngineID = p.EngineID.FromRef(js.Undefined)
	p.Items = p.Items.FromRef(js.Undefined)
}

type MouseButton uint32

const (
	_ MouseButton = iota

	MouseButton_LEFT
	MouseButton_MIDDLE
	MouseButton_RIGHT
)

func (MouseButton) FromRef(str js.Ref) MouseButton {
	return MouseButton(bindings.ConstOfMouseButton(str))
}

func (x MouseButton) String() (string, bool) {
	switch x {
	case MouseButton_LEFT:
		return "left", true
	case MouseButton_MIDDLE:
		return "middle", true
	case MouseButton_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type OnAssistiveWindowButtonClickedArgDetails struct {
	// ButtonID is "OnAssistiveWindowButtonClickedArgDetails.buttonID"
	//
	// Required
	ButtonID AssistiveWindowButton
	// WindowType is "OnAssistiveWindowButtonClickedArgDetails.windowType"
	//
	// Required
	WindowType AssistiveWindowType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnAssistiveWindowButtonClickedArgDetails with all fields set.
func (p OnAssistiveWindowButtonClickedArgDetails) FromRef(ref js.Ref) OnAssistiveWindowButtonClickedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnAssistiveWindowButtonClickedArgDetails in the application heap.
func (p OnAssistiveWindowButtonClickedArgDetails) New() js.Ref {
	return bindings.OnAssistiveWindowButtonClickedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnAssistiveWindowButtonClickedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnAssistiveWindowButtonClickedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnAssistiveWindowButtonClickedArgDetails) Update(ref js.Ref) {
	bindings.OnAssistiveWindowButtonClickedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnAssistiveWindowButtonClickedArgDetails) FreeMembers(recursive bool) {
}

type OnSurroundingTextChangedArgSurroundingInfo struct {
	// Anchor is "OnSurroundingTextChangedArgSurroundingInfo.anchor"
	//
	// Required
	Anchor int64
	// Focus is "OnSurroundingTextChangedArgSurroundingInfo.focus"
	//
	// Required
	Focus int64
	// Offset is "OnSurroundingTextChangedArgSurroundingInfo.offset"
	//
	// Required
	Offset int64
	// Text is "OnSurroundingTextChangedArgSurroundingInfo.text"
	//
	// Required
	Text js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnSurroundingTextChangedArgSurroundingInfo with all fields set.
func (p OnSurroundingTextChangedArgSurroundingInfo) FromRef(ref js.Ref) OnSurroundingTextChangedArgSurroundingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnSurroundingTextChangedArgSurroundingInfo in the application heap.
func (p OnSurroundingTextChangedArgSurroundingInfo) New() js.Ref {
	return bindings.OnSurroundingTextChangedArgSurroundingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnSurroundingTextChangedArgSurroundingInfo) UpdateFrom(ref js.Ref) {
	bindings.OnSurroundingTextChangedArgSurroundingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnSurroundingTextChangedArgSurroundingInfo) Update(ref js.Ref) {
	bindings.OnSurroundingTextChangedArgSurroundingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnSurroundingTextChangedArgSurroundingInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

type ScreenType uint32

const (
	_ ScreenType = iota

	ScreenType_NORMAL
	ScreenType_LOGIN
	ScreenType_LOCK
	ScreenType_SECONDARY_LOGIN
)

func (ScreenType) FromRef(str js.Ref) ScreenType {
	return ScreenType(bindings.ConstOfScreenType(str))
}

func (x ScreenType) String() (string, bool) {
	switch x {
	case ScreenType_NORMAL:
		return "normal", true
	case ScreenType_LOGIN:
		return "login", true
	case ScreenType_LOCK:
		return "lock", true
	case ScreenType_SECONDARY_LOGIN:
		return "secondary-login", true
	default:
		return "", false
	}
}

type SendKeyEventsArgParameters struct {
	// ContextID is "SendKeyEventsArgParameters.contextID"
	//
	// Required
	ContextID int64
	// KeyData is "SendKeyEventsArgParameters.keyData"
	//
	// Required
	KeyData js.Array[KeyboardEvent]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendKeyEventsArgParameters with all fields set.
func (p SendKeyEventsArgParameters) FromRef(ref js.Ref) SendKeyEventsArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendKeyEventsArgParameters in the application heap.
func (p SendKeyEventsArgParameters) New() js.Ref {
	return bindings.SendKeyEventsArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendKeyEventsArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SendKeyEventsArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendKeyEventsArgParameters) Update(ref js.Ref) {
	bindings.SendKeyEventsArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendKeyEventsArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.KeyData.Ref(),
	)
	p.KeyData = p.KeyData.FromRef(js.Undefined)
}

type SetAssistiveWindowButtonHighlightedArgParameters struct {
	// AnnounceString is "SetAssistiveWindowButtonHighlightedArgParameters.announceString"
	//
	// Optional
	AnnounceString js.String
	// ButtonID is "SetAssistiveWindowButtonHighlightedArgParameters.buttonID"
	//
	// Required
	ButtonID AssistiveWindowButton
	// ContextID is "SetAssistiveWindowButtonHighlightedArgParameters.contextID"
	//
	// Required
	ContextID int64
	// Highlighted is "SetAssistiveWindowButtonHighlightedArgParameters.highlighted"
	//
	// Required
	Highlighted bool
	// WindowType is "SetAssistiveWindowButtonHighlightedArgParameters.windowType"
	//
	// Required
	WindowType AssistiveWindowType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetAssistiveWindowButtonHighlightedArgParameters with all fields set.
func (p SetAssistiveWindowButtonHighlightedArgParameters) FromRef(ref js.Ref) SetAssistiveWindowButtonHighlightedArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetAssistiveWindowButtonHighlightedArgParameters in the application heap.
func (p SetAssistiveWindowButtonHighlightedArgParameters) New() js.Ref {
	return bindings.SetAssistiveWindowButtonHighlightedArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetAssistiveWindowButtonHighlightedArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetAssistiveWindowButtonHighlightedArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetAssistiveWindowButtonHighlightedArgParameters) Update(ref js.Ref) {
	bindings.SetAssistiveWindowButtonHighlightedArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetAssistiveWindowButtonHighlightedArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.AnnounceString.Ref(),
	)
	p.AnnounceString = p.AnnounceString.FromRef(js.Undefined)
}

type SetAssistiveWindowPropertiesArgParameters struct {
	// ContextID is "SetAssistiveWindowPropertiesArgParameters.contextID"
	//
	// Required
	ContextID int64
	// Properties is "SetAssistiveWindowPropertiesArgParameters.properties"
	//
	// Required
	//
	// NOTE: Properties.FFI_USE MUST be set to true to get Properties used.
	Properties AssistiveWindowProperties

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetAssistiveWindowPropertiesArgParameters with all fields set.
func (p SetAssistiveWindowPropertiesArgParameters) FromRef(ref js.Ref) SetAssistiveWindowPropertiesArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetAssistiveWindowPropertiesArgParameters in the application heap.
func (p SetAssistiveWindowPropertiesArgParameters) New() js.Ref {
	return bindings.SetAssistiveWindowPropertiesArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetAssistiveWindowPropertiesArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetAssistiveWindowPropertiesArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetAssistiveWindowPropertiesArgParameters) Update(ref js.Ref) {
	bindings.SetAssistiveWindowPropertiesArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetAssistiveWindowPropertiesArgParameters) FreeMembers(recursive bool) {
	if recursive {
		p.Properties.FreeMembers(true)
	}
}

type WindowPosition uint32

const (
	_ WindowPosition = iota

	WindowPosition_CURSOR
	WindowPosition_COMPOSITION
)

func (WindowPosition) FromRef(str js.Ref) WindowPosition {
	return WindowPosition(bindings.ConstOfWindowPosition(str))
}

func (x WindowPosition) String() (string, bool) {
	switch x {
	case WindowPosition_CURSOR:
		return "cursor", true
	case WindowPosition_COMPOSITION:
		return "composition", true
	default:
		return "", false
	}
}

type SetCandidateWindowPropertiesArgParametersFieldProperties struct {
	// AuxiliaryText is "SetCandidateWindowPropertiesArgParametersFieldProperties.auxiliaryText"
	//
	// Optional
	AuxiliaryText js.String
	// AuxiliaryTextVisible is "SetCandidateWindowPropertiesArgParametersFieldProperties.auxiliaryTextVisible"
	//
	// Optional
	//
	// NOTE: FFI_USE_AuxiliaryTextVisible MUST be set to true to make this field effective.
	AuxiliaryTextVisible bool
	// CurrentCandidateIndex is "SetCandidateWindowPropertiesArgParametersFieldProperties.currentCandidateIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentCandidateIndex MUST be set to true to make this field effective.
	CurrentCandidateIndex int64
	// CursorVisible is "SetCandidateWindowPropertiesArgParametersFieldProperties.cursorVisible"
	//
	// Optional
	//
	// NOTE: FFI_USE_CursorVisible MUST be set to true to make this field effective.
	CursorVisible bool
	// PageSize is "SetCandidateWindowPropertiesArgParametersFieldProperties.pageSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_PageSize MUST be set to true to make this field effective.
	PageSize int64
	// TotalCandidates is "SetCandidateWindowPropertiesArgParametersFieldProperties.totalCandidates"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalCandidates MUST be set to true to make this field effective.
	TotalCandidates int64
	// Vertical is "SetCandidateWindowPropertiesArgParametersFieldProperties.vertical"
	//
	// Optional
	//
	// NOTE: FFI_USE_Vertical MUST be set to true to make this field effective.
	Vertical bool
	// Visible is "SetCandidateWindowPropertiesArgParametersFieldProperties.visible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool
	// WindowPosition is "SetCandidateWindowPropertiesArgParametersFieldProperties.windowPosition"
	//
	// Optional
	WindowPosition WindowPosition

	FFI_USE_AuxiliaryTextVisible  bool // for AuxiliaryTextVisible.
	FFI_USE_CurrentCandidateIndex bool // for CurrentCandidateIndex.
	FFI_USE_CursorVisible         bool // for CursorVisible.
	FFI_USE_PageSize              bool // for PageSize.
	FFI_USE_TotalCandidates       bool // for TotalCandidates.
	FFI_USE_Vertical              bool // for Vertical.
	FFI_USE_Visible               bool // for Visible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCandidateWindowPropertiesArgParametersFieldProperties with all fields set.
func (p SetCandidateWindowPropertiesArgParametersFieldProperties) FromRef(ref js.Ref) SetCandidateWindowPropertiesArgParametersFieldProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCandidateWindowPropertiesArgParametersFieldProperties in the application heap.
func (p SetCandidateWindowPropertiesArgParametersFieldProperties) New() js.Ref {
	return bindings.SetCandidateWindowPropertiesArgParametersFieldPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCandidateWindowPropertiesArgParametersFieldProperties) UpdateFrom(ref js.Ref) {
	bindings.SetCandidateWindowPropertiesArgParametersFieldPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCandidateWindowPropertiesArgParametersFieldProperties) Update(ref js.Ref) {
	bindings.SetCandidateWindowPropertiesArgParametersFieldPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCandidateWindowPropertiesArgParametersFieldProperties) FreeMembers(recursive bool) {
	js.Free(
		p.AuxiliaryText.Ref(),
	)
	p.AuxiliaryText = p.AuxiliaryText.FromRef(js.Undefined)
}

type SetCandidateWindowPropertiesArgParameters struct {
	// EngineID is "SetCandidateWindowPropertiesArgParameters.engineID"
	//
	// Required
	EngineID js.String
	// Properties is "SetCandidateWindowPropertiesArgParameters.properties"
	//
	// Required
	//
	// NOTE: Properties.FFI_USE MUST be set to true to get Properties used.
	Properties SetCandidateWindowPropertiesArgParametersFieldProperties

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCandidateWindowPropertiesArgParameters with all fields set.
func (p SetCandidateWindowPropertiesArgParameters) FromRef(ref js.Ref) SetCandidateWindowPropertiesArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCandidateWindowPropertiesArgParameters in the application heap.
func (p SetCandidateWindowPropertiesArgParameters) New() js.Ref {
	return bindings.SetCandidateWindowPropertiesArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCandidateWindowPropertiesArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetCandidateWindowPropertiesArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCandidateWindowPropertiesArgParameters) Update(ref js.Ref) {
	bindings.SetCandidateWindowPropertiesArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCandidateWindowPropertiesArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.EngineID.Ref(),
	)
	p.EngineID = p.EngineID.FromRef(js.Undefined)
	if recursive {
		p.Properties.FreeMembers(true)
	}
}

type SetCandidatesArgParametersFieldCandidatesElemFieldUsage struct {
	// Body is "SetCandidatesArgParametersFieldCandidatesElemFieldUsage.body"
	//
	// Required
	Body js.String
	// Title is "SetCandidatesArgParametersFieldCandidatesElemFieldUsage.title"
	//
	// Required
	Title js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCandidatesArgParametersFieldCandidatesElemFieldUsage with all fields set.
func (p SetCandidatesArgParametersFieldCandidatesElemFieldUsage) FromRef(ref js.Ref) SetCandidatesArgParametersFieldCandidatesElemFieldUsage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCandidatesArgParametersFieldCandidatesElemFieldUsage in the application heap.
func (p SetCandidatesArgParametersFieldCandidatesElemFieldUsage) New() js.Ref {
	return bindings.SetCandidatesArgParametersFieldCandidatesElemFieldUsageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCandidatesArgParametersFieldCandidatesElemFieldUsage) UpdateFrom(ref js.Ref) {
	bindings.SetCandidatesArgParametersFieldCandidatesElemFieldUsageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCandidatesArgParametersFieldCandidatesElemFieldUsage) Update(ref js.Ref) {
	bindings.SetCandidatesArgParametersFieldCandidatesElemFieldUsageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCandidatesArgParametersFieldCandidatesElemFieldUsage) FreeMembers(recursive bool) {
	js.Free(
		p.Body.Ref(),
		p.Title.Ref(),
	)
	p.Body = p.Body.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type SetCandidatesArgParametersFieldCandidatesElem struct {
	// Annotation is "SetCandidatesArgParametersFieldCandidatesElem.annotation"
	//
	// Optional
	Annotation js.String
	// Candidate is "SetCandidatesArgParametersFieldCandidatesElem.candidate"
	//
	// Required
	Candidate js.String
	// Id is "SetCandidatesArgParametersFieldCandidatesElem.id"
	//
	// Required
	Id int64
	// Label is "SetCandidatesArgParametersFieldCandidatesElem.label"
	//
	// Optional
	Label js.String
	// ParentId is "SetCandidatesArgParametersFieldCandidatesElem.parentId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ParentId MUST be set to true to make this field effective.
	ParentId int64
	// Usage is "SetCandidatesArgParametersFieldCandidatesElem.usage"
	//
	// Optional
	//
	// NOTE: Usage.FFI_USE MUST be set to true to get Usage used.
	Usage SetCandidatesArgParametersFieldCandidatesElemFieldUsage

	FFI_USE_ParentId bool // for ParentId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCandidatesArgParametersFieldCandidatesElem with all fields set.
func (p SetCandidatesArgParametersFieldCandidatesElem) FromRef(ref js.Ref) SetCandidatesArgParametersFieldCandidatesElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCandidatesArgParametersFieldCandidatesElem in the application heap.
func (p SetCandidatesArgParametersFieldCandidatesElem) New() js.Ref {
	return bindings.SetCandidatesArgParametersFieldCandidatesElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCandidatesArgParametersFieldCandidatesElem) UpdateFrom(ref js.Ref) {
	bindings.SetCandidatesArgParametersFieldCandidatesElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCandidatesArgParametersFieldCandidatesElem) Update(ref js.Ref) {
	bindings.SetCandidatesArgParametersFieldCandidatesElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCandidatesArgParametersFieldCandidatesElem) FreeMembers(recursive bool) {
	js.Free(
		p.Annotation.Ref(),
		p.Candidate.Ref(),
		p.Label.Ref(),
	)
	p.Annotation = p.Annotation.FromRef(js.Undefined)
	p.Candidate = p.Candidate.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
	if recursive {
		p.Usage.FreeMembers(true)
	}
}

type SetCandidatesArgParameters struct {
	// Candidates is "SetCandidatesArgParameters.candidates"
	//
	// Required
	Candidates js.Array[SetCandidatesArgParametersFieldCandidatesElem]
	// ContextID is "SetCandidatesArgParameters.contextID"
	//
	// Required
	ContextID int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCandidatesArgParameters with all fields set.
func (p SetCandidatesArgParameters) FromRef(ref js.Ref) SetCandidatesArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCandidatesArgParameters in the application heap.
func (p SetCandidatesArgParameters) New() js.Ref {
	return bindings.SetCandidatesArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCandidatesArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetCandidatesArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCandidatesArgParameters) Update(ref js.Ref) {
	bindings.SetCandidatesArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCandidatesArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Candidates.Ref(),
	)
	p.Candidates = p.Candidates.FromRef(js.Undefined)
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

type SetCompositionArgParametersFieldSegmentsElem struct {
	// End is "SetCompositionArgParametersFieldSegmentsElem.end"
	//
	// Required
	End int64
	// Start is "SetCompositionArgParametersFieldSegmentsElem.start"
	//
	// Required
	Start int64
	// Style is "SetCompositionArgParametersFieldSegmentsElem.style"
	//
	// Required
	Style UnderlineStyle

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCompositionArgParametersFieldSegmentsElem with all fields set.
func (p SetCompositionArgParametersFieldSegmentsElem) FromRef(ref js.Ref) SetCompositionArgParametersFieldSegmentsElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCompositionArgParametersFieldSegmentsElem in the application heap.
func (p SetCompositionArgParametersFieldSegmentsElem) New() js.Ref {
	return bindings.SetCompositionArgParametersFieldSegmentsElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCompositionArgParametersFieldSegmentsElem) UpdateFrom(ref js.Ref) {
	bindings.SetCompositionArgParametersFieldSegmentsElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCompositionArgParametersFieldSegmentsElem) Update(ref js.Ref) {
	bindings.SetCompositionArgParametersFieldSegmentsElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCompositionArgParametersFieldSegmentsElem) FreeMembers(recursive bool) {
}

type SetCompositionArgParameters struct {
	// ContextID is "SetCompositionArgParameters.contextID"
	//
	// Required
	ContextID int64
	// Cursor is "SetCompositionArgParameters.cursor"
	//
	// Required
	Cursor int64
	// Segments is "SetCompositionArgParameters.segments"
	//
	// Optional
	Segments js.Array[SetCompositionArgParametersFieldSegmentsElem]
	// SelectionEnd is "SetCompositionArgParameters.selectionEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionEnd MUST be set to true to make this field effective.
	SelectionEnd int64
	// SelectionStart is "SetCompositionArgParameters.selectionStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_SelectionStart MUST be set to true to make this field effective.
	SelectionStart int64
	// Text is "SetCompositionArgParameters.text"
	//
	// Required
	Text js.String

	FFI_USE_SelectionEnd   bool // for SelectionEnd.
	FFI_USE_SelectionStart bool // for SelectionStart.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCompositionArgParameters with all fields set.
func (p SetCompositionArgParameters) FromRef(ref js.Ref) SetCompositionArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCompositionArgParameters in the application heap.
func (p SetCompositionArgParameters) New() js.Ref {
	return bindings.SetCompositionArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCompositionArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetCompositionArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCompositionArgParameters) Update(ref js.Ref) {
	bindings.SetCompositionArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCompositionArgParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Segments.Ref(),
		p.Text.Ref(),
	)
	p.Segments = p.Segments.FromRef(js.Undefined)
	p.Text = p.Text.FromRef(js.Undefined)
}

type SetCursorPositionArgParameters struct {
	// CandidateID is "SetCursorPositionArgParameters.candidateID"
	//
	// Required
	CandidateID int64
	// ContextID is "SetCursorPositionArgParameters.contextID"
	//
	// Required
	ContextID int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCursorPositionArgParameters with all fields set.
func (p SetCursorPositionArgParameters) FromRef(ref js.Ref) SetCursorPositionArgParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCursorPositionArgParameters in the application heap.
func (p SetCursorPositionArgParameters) New() js.Ref {
	return bindings.SetCursorPositionArgParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCursorPositionArgParameters) UpdateFrom(ref js.Ref) {
	bindings.SetCursorPositionArgParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCursorPositionArgParameters) Update(ref js.Ref) {
	bindings.SetCursorPositionArgParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCursorPositionArgParameters) FreeMembers(recursive bool) {
}

// HasFuncClearComposition returns true if the function "WEBEXT.input.ime.clearComposition" exists.
func HasFuncClearComposition() bool {
	return js.True == bindings.HasFuncClearComposition()
}

// FuncClearComposition returns the function "WEBEXT.input.ime.clearComposition".
func FuncClearComposition() (fn js.Func[func(parameters ClearCompositionArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncClearComposition(
		js.Pointer(&fn),
	)
	return
}

// ClearComposition calls the function "WEBEXT.input.ime.clearComposition" directly.
func ClearComposition(parameters ClearCompositionArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallClearComposition(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryClearComposition calls the function "WEBEXT.input.ime.clearComposition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearComposition(parameters ClearCompositionArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearComposition(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncCommitText returns true if the function "WEBEXT.input.ime.commitText" exists.
func HasFuncCommitText() bool {
	return js.True == bindings.HasFuncCommitText()
}

// FuncCommitText returns the function "WEBEXT.input.ime.commitText".
func FuncCommitText() (fn js.Func[func(parameters CommitTextArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncCommitText(
		js.Pointer(&fn),
	)
	return
}

// CommitText calls the function "WEBEXT.input.ime.commitText" directly.
func CommitText(parameters CommitTextArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallCommitText(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryCommitText calls the function "WEBEXT.input.ime.commitText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCommitText(parameters CommitTextArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCommitText(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncDeleteSurroundingText returns true if the function "WEBEXT.input.ime.deleteSurroundingText" exists.
func HasFuncDeleteSurroundingText() bool {
	return js.True == bindings.HasFuncDeleteSurroundingText()
}

// FuncDeleteSurroundingText returns the function "WEBEXT.input.ime.deleteSurroundingText".
func FuncDeleteSurroundingText() (fn js.Func[func(parameters DeleteSurroundingTextArgParameters) js.Promise[js.Void]]) {
	bindings.FuncDeleteSurroundingText(
		js.Pointer(&fn),
	)
	return
}

// DeleteSurroundingText calls the function "WEBEXT.input.ime.deleteSurroundingText" directly.
func DeleteSurroundingText(parameters DeleteSurroundingTextArgParameters) (ret js.Promise[js.Void]) {
	bindings.CallDeleteSurroundingText(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryDeleteSurroundingText calls the function "WEBEXT.input.ime.deleteSurroundingText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteSurroundingText(parameters DeleteSurroundingTextArgParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteSurroundingText(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncHideInputView returns true if the function "WEBEXT.input.ime.hideInputView" exists.
func HasFuncHideInputView() bool {
	return js.True == bindings.HasFuncHideInputView()
}

// FuncHideInputView returns the function "WEBEXT.input.ime.hideInputView".
func FuncHideInputView() (fn js.Func[func()]) {
	bindings.FuncHideInputView(
		js.Pointer(&fn),
	)
	return
}

// HideInputView calls the function "WEBEXT.input.ime.hideInputView" directly.
func HideInputView() (ret js.Void) {
	bindings.CallHideInputView(
		js.Pointer(&ret),
	)

	return
}

// TryHideInputView calls the function "WEBEXT.input.ime.hideInputView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHideInputView() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHideInputView(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncKeyEventHandled returns true if the function "WEBEXT.input.ime.keyEventHandled" exists.
func HasFuncKeyEventHandled() bool {
	return js.True == bindings.HasFuncKeyEventHandled()
}

// FuncKeyEventHandled returns the function "WEBEXT.input.ime.keyEventHandled".
func FuncKeyEventHandled() (fn js.Func[func(requestId js.String, response bool)]) {
	bindings.FuncKeyEventHandled(
		js.Pointer(&fn),
	)
	return
}

// KeyEventHandled calls the function "WEBEXT.input.ime.keyEventHandled" directly.
func KeyEventHandled(requestId js.String, response bool) (ret js.Void) {
	bindings.CallKeyEventHandled(
		js.Pointer(&ret),
		requestId.Ref(),
		js.Bool(bool(response)),
	)

	return
}

// TryKeyEventHandled calls the function "WEBEXT.input.ime.keyEventHandled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryKeyEventHandled(requestId js.String, response bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyEventHandled(
		js.Pointer(&ret), js.Pointer(&exception),
		requestId.Ref(),
		js.Bool(bool(response)),
	)

	return
}

type OnActivateEventCallbackFunc func(this js.Ref, engineID js.String, screen ScreenType) js.Ref

func (fn OnActivateEventCallbackFunc) Register() js.Func[func(engineID js.String, screen ScreenType)] {
	return js.RegisterCallback[func(engineID js.String, screen ScreenType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActivateEventCallbackFunc) DispatchCallback(
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
		ScreenType(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnActivateEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, screen ScreenType) js.Ref
	Arg T
}

func (cb *OnActivateEventCallback[T]) Register() js.Func[func(engineID js.String, screen ScreenType)] {
	return js.RegisterCallback[func(engineID js.String, screen ScreenType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActivateEventCallback[T]) DispatchCallback(
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
		ScreenType(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnActivate returns true if the function "WEBEXT.input.ime.onActivate.addListener" exists.
func HasFuncOnActivate() bool {
	return js.True == bindings.HasFuncOnActivate()
}

// FuncOnActivate returns the function "WEBEXT.input.ime.onActivate.addListener".
func FuncOnActivate() (fn js.Func[func(callback js.Func[func(engineID js.String, screen ScreenType)])]) {
	bindings.FuncOnActivate(
		js.Pointer(&fn),
	)
	return
}

// OnActivate calls the function "WEBEXT.input.ime.onActivate.addListener" directly.
func OnActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret js.Void) {
	bindings.CallOnActivate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActivate calls the function "WEBEXT.input.ime.onActivate.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActivate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActivate returns true if the function "WEBEXT.input.ime.onActivate.removeListener" exists.
func HasFuncOffActivate() bool {
	return js.True == bindings.HasFuncOffActivate()
}

// FuncOffActivate returns the function "WEBEXT.input.ime.onActivate.removeListener".
func FuncOffActivate() (fn js.Func[func(callback js.Func[func(engineID js.String, screen ScreenType)])]) {
	bindings.FuncOffActivate(
		js.Pointer(&fn),
	)
	return
}

// OffActivate calls the function "WEBEXT.input.ime.onActivate.removeListener" directly.
func OffActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret js.Void) {
	bindings.CallOffActivate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActivate calls the function "WEBEXT.input.ime.onActivate.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActivate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActivate returns true if the function "WEBEXT.input.ime.onActivate.hasListener" exists.
func HasFuncHasOnActivate() bool {
	return js.True == bindings.HasFuncHasOnActivate()
}

// FuncHasOnActivate returns the function "WEBEXT.input.ime.onActivate.hasListener".
func FuncHasOnActivate() (fn js.Func[func(callback js.Func[func(engineID js.String, screen ScreenType)]) bool]) {
	bindings.FuncHasOnActivate(
		js.Pointer(&fn),
	)
	return
}

// HasOnActivate calls the function "WEBEXT.input.ime.onActivate.hasListener" directly.
func HasOnActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret bool) {
	bindings.CallHasOnActivate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActivate calls the function "WEBEXT.input.ime.onActivate.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActivate(callback js.Func[func(engineID js.String, screen ScreenType)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActivate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAssistiveWindowButtonClickedEventCallbackFunc func(this js.Ref, details *OnAssistiveWindowButtonClickedArgDetails) js.Ref

func (fn OnAssistiveWindowButtonClickedEventCallbackFunc) Register() js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)] {
	return js.RegisterCallback[func(details *OnAssistiveWindowButtonClickedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAssistiveWindowButtonClickedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnAssistiveWindowButtonClickedArgDetails
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

type OnAssistiveWindowButtonClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnAssistiveWindowButtonClickedArgDetails) js.Ref
	Arg T
}

func (cb *OnAssistiveWindowButtonClickedEventCallback[T]) Register() js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)] {
	return js.RegisterCallback[func(details *OnAssistiveWindowButtonClickedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAssistiveWindowButtonClickedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnAssistiveWindowButtonClickedArgDetails
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

// HasFuncOnAssistiveWindowButtonClicked returns true if the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener" exists.
func HasFuncOnAssistiveWindowButtonClicked() bool {
	return js.True == bindings.HasFuncOnAssistiveWindowButtonClicked()
}

// FuncOnAssistiveWindowButtonClicked returns the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener".
func FuncOnAssistiveWindowButtonClicked() (fn js.Func[func(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)])]) {
	bindings.FuncOnAssistiveWindowButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// OnAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener" directly.
func OnAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret js.Void) {
	bindings.CallOnAssistiveWindowButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAssistiveWindowButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAssistiveWindowButtonClicked returns true if the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener" exists.
func HasFuncOffAssistiveWindowButtonClicked() bool {
	return js.True == bindings.HasFuncOffAssistiveWindowButtonClicked()
}

// FuncOffAssistiveWindowButtonClicked returns the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener".
func FuncOffAssistiveWindowButtonClicked() (fn js.Func[func(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)])]) {
	bindings.FuncOffAssistiveWindowButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// OffAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener" directly.
func OffAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret js.Void) {
	bindings.CallOffAssistiveWindowButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAssistiveWindowButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAssistiveWindowButtonClicked returns true if the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener" exists.
func HasFuncHasOnAssistiveWindowButtonClicked() bool {
	return js.True == bindings.HasFuncHasOnAssistiveWindowButtonClicked()
}

// FuncHasOnAssistiveWindowButtonClicked returns the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener".
func FuncHasOnAssistiveWindowButtonClicked() (fn js.Func[func(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) bool]) {
	bindings.FuncHasOnAssistiveWindowButtonClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener" directly.
func HasOnAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret bool) {
	bindings.CallHasOnAssistiveWindowButtonClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAssistiveWindowButtonClicked calls the function "WEBEXT.input.ime.onAssistiveWindowButtonClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAssistiveWindowButtonClicked(callback js.Func[func(details *OnAssistiveWindowButtonClickedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAssistiveWindowButtonClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBlurEventCallbackFunc func(this js.Ref, contextID int64) js.Ref

func (fn OnBlurEventCallbackFunc) Register() js.Func[func(contextID int64)] {
	return js.RegisterCallback[func(contextID int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBlurEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnBlurEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, contextID int64) js.Ref
	Arg T
}

func (cb *OnBlurEventCallback[T]) Register() js.Func[func(contextID int64)] {
	return js.RegisterCallback[func(contextID int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBlurEventCallback[T]) DispatchCallback(
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

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnBlur returns true if the function "WEBEXT.input.ime.onBlur.addListener" exists.
func HasFuncOnBlur() bool {
	return js.True == bindings.HasFuncOnBlur()
}

// FuncOnBlur returns the function "WEBEXT.input.ime.onBlur.addListener".
func FuncOnBlur() (fn js.Func[func(callback js.Func[func(contextID int64)])]) {
	bindings.FuncOnBlur(
		js.Pointer(&fn),
	)
	return
}

// OnBlur calls the function "WEBEXT.input.ime.onBlur.addListener" directly.
func OnBlur(callback js.Func[func(contextID int64)]) (ret js.Void) {
	bindings.CallOnBlur(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBlur calls the function "WEBEXT.input.ime.onBlur.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBlur(callback js.Func[func(contextID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBlur(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBlur returns true if the function "WEBEXT.input.ime.onBlur.removeListener" exists.
func HasFuncOffBlur() bool {
	return js.True == bindings.HasFuncOffBlur()
}

// FuncOffBlur returns the function "WEBEXT.input.ime.onBlur.removeListener".
func FuncOffBlur() (fn js.Func[func(callback js.Func[func(contextID int64)])]) {
	bindings.FuncOffBlur(
		js.Pointer(&fn),
	)
	return
}

// OffBlur calls the function "WEBEXT.input.ime.onBlur.removeListener" directly.
func OffBlur(callback js.Func[func(contextID int64)]) (ret js.Void) {
	bindings.CallOffBlur(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBlur calls the function "WEBEXT.input.ime.onBlur.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBlur(callback js.Func[func(contextID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBlur(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBlur returns true if the function "WEBEXT.input.ime.onBlur.hasListener" exists.
func HasFuncHasOnBlur() bool {
	return js.True == bindings.HasFuncHasOnBlur()
}

// FuncHasOnBlur returns the function "WEBEXT.input.ime.onBlur.hasListener".
func FuncHasOnBlur() (fn js.Func[func(callback js.Func[func(contextID int64)]) bool]) {
	bindings.FuncHasOnBlur(
		js.Pointer(&fn),
	)
	return
}

// HasOnBlur calls the function "WEBEXT.input.ime.onBlur.hasListener" directly.
func HasOnBlur(callback js.Func[func(contextID int64)]) (ret bool) {
	bindings.CallHasOnBlur(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBlur calls the function "WEBEXT.input.ime.onBlur.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBlur(callback js.Func[func(contextID int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBlur(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCandidateClickedEventCallbackFunc func(this js.Ref, engineID js.String, candidateID int64, button MouseButton) js.Ref

func (fn OnCandidateClickedEventCallbackFunc) Register() js.Func[func(engineID js.String, candidateID int64, button MouseButton)] {
	return js.RegisterCallback[func(engineID js.String, candidateID int64, button MouseButton)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCandidateClickedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		MouseButton(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCandidateClickedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, candidateID int64, button MouseButton) js.Ref
	Arg T
}

func (cb *OnCandidateClickedEventCallback[T]) Register() js.Func[func(engineID js.String, candidateID int64, button MouseButton)] {
	return js.RegisterCallback[func(engineID js.String, candidateID int64, button MouseButton)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCandidateClickedEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		MouseButton(0).FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCandidateClicked returns true if the function "WEBEXT.input.ime.onCandidateClicked.addListener" exists.
func HasFuncOnCandidateClicked() bool {
	return js.True == bindings.HasFuncOnCandidateClicked()
}

// FuncOnCandidateClicked returns the function "WEBEXT.input.ime.onCandidateClicked.addListener".
func FuncOnCandidateClicked() (fn js.Func[func(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)])]) {
	bindings.FuncOnCandidateClicked(
		js.Pointer(&fn),
	)
	return
}

// OnCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.addListener" directly.
func OnCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret js.Void) {
	bindings.CallOnCandidateClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCandidateClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCandidateClicked returns true if the function "WEBEXT.input.ime.onCandidateClicked.removeListener" exists.
func HasFuncOffCandidateClicked() bool {
	return js.True == bindings.HasFuncOffCandidateClicked()
}

// FuncOffCandidateClicked returns the function "WEBEXT.input.ime.onCandidateClicked.removeListener".
func FuncOffCandidateClicked() (fn js.Func[func(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)])]) {
	bindings.FuncOffCandidateClicked(
		js.Pointer(&fn),
	)
	return
}

// OffCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.removeListener" directly.
func OffCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret js.Void) {
	bindings.CallOffCandidateClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCandidateClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCandidateClicked returns true if the function "WEBEXT.input.ime.onCandidateClicked.hasListener" exists.
func HasFuncHasOnCandidateClicked() bool {
	return js.True == bindings.HasFuncHasOnCandidateClicked()
}

// FuncHasOnCandidateClicked returns the function "WEBEXT.input.ime.onCandidateClicked.hasListener".
func FuncHasOnCandidateClicked() (fn js.Func[func(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) bool]) {
	bindings.FuncHasOnCandidateClicked(
		js.Pointer(&fn),
	)
	return
}

// HasOnCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.hasListener" directly.
func HasOnCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret bool) {
	bindings.CallHasOnCandidateClicked(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCandidateClicked calls the function "WEBEXT.input.ime.onCandidateClicked.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCandidateClicked(callback js.Func[func(engineID js.String, candidateID int64, button MouseButton)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCandidateClicked(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeactivatedEventCallbackFunc func(this js.Ref, engineID js.String) js.Ref

func (fn OnDeactivatedEventCallbackFunc) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeactivatedEventCallbackFunc) DispatchCallback(
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

type OnDeactivatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String) js.Ref
	Arg T
}

func (cb *OnDeactivatedEventCallback[T]) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeactivatedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDeactivated returns true if the function "WEBEXT.input.ime.onDeactivated.addListener" exists.
func HasFuncOnDeactivated() bool {
	return js.True == bindings.HasFuncOnDeactivated()
}

// FuncOnDeactivated returns the function "WEBEXT.input.ime.onDeactivated.addListener".
func FuncOnDeactivated() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOnDeactivated(
		js.Pointer(&fn),
	)
	return
}

// OnDeactivated calls the function "WEBEXT.input.ime.onDeactivated.addListener" directly.
func OnDeactivated(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOnDeactivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeactivated calls the function "WEBEXT.input.ime.onDeactivated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeactivated(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeactivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeactivated returns true if the function "WEBEXT.input.ime.onDeactivated.removeListener" exists.
func HasFuncOffDeactivated() bool {
	return js.True == bindings.HasFuncOffDeactivated()
}

// FuncOffDeactivated returns the function "WEBEXT.input.ime.onDeactivated.removeListener".
func FuncOffDeactivated() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOffDeactivated(
		js.Pointer(&fn),
	)
	return
}

// OffDeactivated calls the function "WEBEXT.input.ime.onDeactivated.removeListener" directly.
func OffDeactivated(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOffDeactivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeactivated calls the function "WEBEXT.input.ime.onDeactivated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeactivated(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeactivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeactivated returns true if the function "WEBEXT.input.ime.onDeactivated.hasListener" exists.
func HasFuncHasOnDeactivated() bool {
	return js.True == bindings.HasFuncHasOnDeactivated()
}

// FuncHasOnDeactivated returns the function "WEBEXT.input.ime.onDeactivated.hasListener".
func FuncHasOnDeactivated() (fn js.Func[func(callback js.Func[func(engineID js.String)]) bool]) {
	bindings.FuncHasOnDeactivated(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeactivated calls the function "WEBEXT.input.ime.onDeactivated.hasListener" directly.
func HasOnDeactivated(callback js.Func[func(engineID js.String)]) (ret bool) {
	bindings.CallHasOnDeactivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeactivated calls the function "WEBEXT.input.ime.onDeactivated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeactivated(callback js.Func[func(engineID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeactivated(
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

// HasFuncOnFocus returns true if the function "WEBEXT.input.ime.onFocus.addListener" exists.
func HasFuncOnFocus() bool {
	return js.True == bindings.HasFuncOnFocus()
}

// FuncOnFocus returns the function "WEBEXT.input.ime.onFocus.addListener".
func FuncOnFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOnFocus(
		js.Pointer(&fn),
	)
	return
}

// OnFocus calls the function "WEBEXT.input.ime.onFocus.addListener" directly.
func OnFocus(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOnFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFocus calls the function "WEBEXT.input.ime.onFocus.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFocus(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFocus returns true if the function "WEBEXT.input.ime.onFocus.removeListener" exists.
func HasFuncOffFocus() bool {
	return js.True == bindings.HasFuncOffFocus()
}

// FuncOffFocus returns the function "WEBEXT.input.ime.onFocus.removeListener".
func FuncOffFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOffFocus(
		js.Pointer(&fn),
	)
	return
}

// OffFocus calls the function "WEBEXT.input.ime.onFocus.removeListener" directly.
func OffFocus(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOffFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFocus calls the function "WEBEXT.input.ime.onFocus.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFocus(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFocus returns true if the function "WEBEXT.input.ime.onFocus.hasListener" exists.
func HasFuncHasOnFocus() bool {
	return js.True == bindings.HasFuncHasOnFocus()
}

// FuncHasOnFocus returns the function "WEBEXT.input.ime.onFocus.hasListener".
func FuncHasOnFocus() (fn js.Func[func(callback js.Func[func(context *InputContext)]) bool]) {
	bindings.FuncHasOnFocus(
		js.Pointer(&fn),
	)
	return
}

// HasOnFocus calls the function "WEBEXT.input.ime.onFocus.hasListener" directly.
func HasOnFocus(callback js.Func[func(context *InputContext)]) (ret bool) {
	bindings.CallHasOnFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFocus calls the function "WEBEXT.input.ime.onFocus.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFocus(callback js.Func[func(context *InputContext)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInputContextUpdateEventCallbackFunc func(this js.Ref, context *InputContext) js.Ref

func (fn OnInputContextUpdateEventCallbackFunc) Register() js.Func[func(context *InputContext)] {
	return js.RegisterCallback[func(context *InputContext)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInputContextUpdateEventCallbackFunc) DispatchCallback(
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

type OnInputContextUpdateEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, context *InputContext) js.Ref
	Arg T
}

func (cb *OnInputContextUpdateEventCallback[T]) Register() js.Func[func(context *InputContext)] {
	return js.RegisterCallback[func(context *InputContext)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInputContextUpdateEventCallback[T]) DispatchCallback(
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

// HasFuncOnInputContextUpdate returns true if the function "WEBEXT.input.ime.onInputContextUpdate.addListener" exists.
func HasFuncOnInputContextUpdate() bool {
	return js.True == bindings.HasFuncOnInputContextUpdate()
}

// FuncOnInputContextUpdate returns the function "WEBEXT.input.ime.onInputContextUpdate.addListener".
func FuncOnInputContextUpdate() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOnInputContextUpdate(
		js.Pointer(&fn),
	)
	return
}

// OnInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.addListener" directly.
func OnInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOnInputContextUpdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInputContextUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInputContextUpdate returns true if the function "WEBEXT.input.ime.onInputContextUpdate.removeListener" exists.
func HasFuncOffInputContextUpdate() bool {
	return js.True == bindings.HasFuncOffInputContextUpdate()
}

// FuncOffInputContextUpdate returns the function "WEBEXT.input.ime.onInputContextUpdate.removeListener".
func FuncOffInputContextUpdate() (fn js.Func[func(callback js.Func[func(context *InputContext)])]) {
	bindings.FuncOffInputContextUpdate(
		js.Pointer(&fn),
	)
	return
}

// OffInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.removeListener" directly.
func OffInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret js.Void) {
	bindings.CallOffInputContextUpdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInputContextUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInputContextUpdate returns true if the function "WEBEXT.input.ime.onInputContextUpdate.hasListener" exists.
func HasFuncHasOnInputContextUpdate() bool {
	return js.True == bindings.HasFuncHasOnInputContextUpdate()
}

// FuncHasOnInputContextUpdate returns the function "WEBEXT.input.ime.onInputContextUpdate.hasListener".
func FuncHasOnInputContextUpdate() (fn js.Func[func(callback js.Func[func(context *InputContext)]) bool]) {
	bindings.FuncHasOnInputContextUpdate(
		js.Pointer(&fn),
	)
	return
}

// HasOnInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.hasListener" directly.
func HasOnInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret bool) {
	bindings.CallHasOnInputContextUpdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInputContextUpdate calls the function "WEBEXT.input.ime.onInputContextUpdate.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInputContextUpdate(callback js.Func[func(context *InputContext)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInputContextUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnKeyEventEventCallbackFunc func(this js.Ref, engineID js.String, keyData *KeyboardEvent, requestId js.String) js.Ref

func (fn OnKeyEventEventCallbackFunc) Register() js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool] {
	return js.RegisterCallback[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnKeyEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 KeyboardEvent
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnKeyEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, keyData *KeyboardEvent, requestId js.String) js.Ref
	Arg T
}

func (cb *OnKeyEventEventCallback[T]) Register() js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool] {
	return js.RegisterCallback[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnKeyEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 KeyboardEvent
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnKeyEvent returns true if the function "WEBEXT.input.ime.onKeyEvent.addListener" exists.
func HasFuncOnKeyEvent() bool {
	return js.True == bindings.HasFuncOnKeyEvent()
}

// FuncOnKeyEvent returns the function "WEBEXT.input.ime.onKeyEvent.addListener".
func FuncOnKeyEvent() (fn js.Func[func(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool])]) {
	bindings.FuncOnKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// OnKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.addListener" directly.
func OnKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret js.Void) {
	bindings.CallOnKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffKeyEvent returns true if the function "WEBEXT.input.ime.onKeyEvent.removeListener" exists.
func HasFuncOffKeyEvent() bool {
	return js.True == bindings.HasFuncOffKeyEvent()
}

// FuncOffKeyEvent returns the function "WEBEXT.input.ime.onKeyEvent.removeListener".
func FuncOffKeyEvent() (fn js.Func[func(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool])]) {
	bindings.FuncOffKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// OffKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.removeListener" directly.
func OffKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret js.Void) {
	bindings.CallOffKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnKeyEvent returns true if the function "WEBEXT.input.ime.onKeyEvent.hasListener" exists.
func HasFuncHasOnKeyEvent() bool {
	return js.True == bindings.HasFuncHasOnKeyEvent()
}

// FuncHasOnKeyEvent returns the function "WEBEXT.input.ime.onKeyEvent.hasListener".
func FuncHasOnKeyEvent() (fn js.Func[func(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) bool]) {
	bindings.FuncHasOnKeyEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.hasListener" directly.
func HasOnKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret bool) {
	bindings.CallHasOnKeyEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnKeyEvent calls the function "WEBEXT.input.ime.onKeyEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnKeyEvent(callback js.Func[func(engineID js.String, keyData *KeyboardEvent, requestId js.String) bool]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnKeyEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMenuItemActivatedEventCallbackFunc func(this js.Ref, engineID js.String, name js.String) js.Ref

func (fn OnMenuItemActivatedEventCallbackFunc) Register() js.Func[func(engineID js.String, name js.String)] {
	return js.RegisterCallback[func(engineID js.String, name js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMenuItemActivatedEventCallbackFunc) DispatchCallback(
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
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMenuItemActivatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, name js.String) js.Ref
	Arg T
}

func (cb *OnMenuItemActivatedEventCallback[T]) Register() js.Func[func(engineID js.String, name js.String)] {
	return js.RegisterCallback[func(engineID js.String, name js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMenuItemActivatedEventCallback[T]) DispatchCallback(
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
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMenuItemActivated returns true if the function "WEBEXT.input.ime.onMenuItemActivated.addListener" exists.
func HasFuncOnMenuItemActivated() bool {
	return js.True == bindings.HasFuncOnMenuItemActivated()
}

// FuncOnMenuItemActivated returns the function "WEBEXT.input.ime.onMenuItemActivated.addListener".
func FuncOnMenuItemActivated() (fn js.Func[func(callback js.Func[func(engineID js.String, name js.String)])]) {
	bindings.FuncOnMenuItemActivated(
		js.Pointer(&fn),
	)
	return
}

// OnMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.addListener" directly.
func OnMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret js.Void) {
	bindings.CallOnMenuItemActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMenuItemActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMenuItemActivated returns true if the function "WEBEXT.input.ime.onMenuItemActivated.removeListener" exists.
func HasFuncOffMenuItemActivated() bool {
	return js.True == bindings.HasFuncOffMenuItemActivated()
}

// FuncOffMenuItemActivated returns the function "WEBEXT.input.ime.onMenuItemActivated.removeListener".
func FuncOffMenuItemActivated() (fn js.Func[func(callback js.Func[func(engineID js.String, name js.String)])]) {
	bindings.FuncOffMenuItemActivated(
		js.Pointer(&fn),
	)
	return
}

// OffMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.removeListener" directly.
func OffMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret js.Void) {
	bindings.CallOffMenuItemActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMenuItemActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMenuItemActivated returns true if the function "WEBEXT.input.ime.onMenuItemActivated.hasListener" exists.
func HasFuncHasOnMenuItemActivated() bool {
	return js.True == bindings.HasFuncHasOnMenuItemActivated()
}

// FuncHasOnMenuItemActivated returns the function "WEBEXT.input.ime.onMenuItemActivated.hasListener".
func FuncHasOnMenuItemActivated() (fn js.Func[func(callback js.Func[func(engineID js.String, name js.String)]) bool]) {
	bindings.FuncHasOnMenuItemActivated(
		js.Pointer(&fn),
	)
	return
}

// HasOnMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.hasListener" directly.
func HasOnMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret bool) {
	bindings.CallHasOnMenuItemActivated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMenuItemActivated calls the function "WEBEXT.input.ime.onMenuItemActivated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMenuItemActivated(callback js.Func[func(engineID js.String, name js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMenuItemActivated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnResetEventCallbackFunc func(this js.Ref, engineID js.String) js.Ref

func (fn OnResetEventCallbackFunc) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnResetEventCallbackFunc) DispatchCallback(
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

type OnResetEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String) js.Ref
	Arg T
}

func (cb *OnResetEventCallback[T]) Register() js.Func[func(engineID js.String)] {
	return js.RegisterCallback[func(engineID js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnResetEventCallback[T]) DispatchCallback(
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

// HasFuncOnReset returns true if the function "WEBEXT.input.ime.onReset.addListener" exists.
func HasFuncOnReset() bool {
	return js.True == bindings.HasFuncOnReset()
}

// FuncOnReset returns the function "WEBEXT.input.ime.onReset.addListener".
func FuncOnReset() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOnReset(
		js.Pointer(&fn),
	)
	return
}

// OnReset calls the function "WEBEXT.input.ime.onReset.addListener" directly.
func OnReset(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOnReset(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReset calls the function "WEBEXT.input.ime.onReset.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReset(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReset(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReset returns true if the function "WEBEXT.input.ime.onReset.removeListener" exists.
func HasFuncOffReset() bool {
	return js.True == bindings.HasFuncOffReset()
}

// FuncOffReset returns the function "WEBEXT.input.ime.onReset.removeListener".
func FuncOffReset() (fn js.Func[func(callback js.Func[func(engineID js.String)])]) {
	bindings.FuncOffReset(
		js.Pointer(&fn),
	)
	return
}

// OffReset calls the function "WEBEXT.input.ime.onReset.removeListener" directly.
func OffReset(callback js.Func[func(engineID js.String)]) (ret js.Void) {
	bindings.CallOffReset(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReset calls the function "WEBEXT.input.ime.onReset.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReset(callback js.Func[func(engineID js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReset(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReset returns true if the function "WEBEXT.input.ime.onReset.hasListener" exists.
func HasFuncHasOnReset() bool {
	return js.True == bindings.HasFuncHasOnReset()
}

// FuncHasOnReset returns the function "WEBEXT.input.ime.onReset.hasListener".
func FuncHasOnReset() (fn js.Func[func(callback js.Func[func(engineID js.String)]) bool]) {
	bindings.FuncHasOnReset(
		js.Pointer(&fn),
	)
	return
}

// HasOnReset calls the function "WEBEXT.input.ime.onReset.hasListener" directly.
func HasOnReset(callback js.Func[func(engineID js.String)]) (ret bool) {
	bindings.CallHasOnReset(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReset calls the function "WEBEXT.input.ime.onReset.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReset(callback js.Func[func(engineID js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReset(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSurroundingTextChangedEventCallbackFunc func(this js.Ref, engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo) js.Ref

func (fn OnSurroundingTextChangedEventCallbackFunc) Register() js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)] {
	return js.RegisterCallback[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSurroundingTextChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnSurroundingTextChangedArgSurroundingInfo
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

type OnSurroundingTextChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo) js.Ref
	Arg T
}

func (cb *OnSurroundingTextChangedEventCallback[T]) Register() js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)] {
	return js.RegisterCallback[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSurroundingTextChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 OnSurroundingTextChangedArgSurroundingInfo
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

// HasFuncOnSurroundingTextChanged returns true if the function "WEBEXT.input.ime.onSurroundingTextChanged.addListener" exists.
func HasFuncOnSurroundingTextChanged() bool {
	return js.True == bindings.HasFuncOnSurroundingTextChanged()
}

// FuncOnSurroundingTextChanged returns the function "WEBEXT.input.ime.onSurroundingTextChanged.addListener".
func FuncOnSurroundingTextChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)])]) {
	bindings.FuncOnSurroundingTextChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.addListener" directly.
func OnSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret js.Void) {
	bindings.CallOnSurroundingTextChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSurroundingTextChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSurroundingTextChanged returns true if the function "WEBEXT.input.ime.onSurroundingTextChanged.removeListener" exists.
func HasFuncOffSurroundingTextChanged() bool {
	return js.True == bindings.HasFuncOffSurroundingTextChanged()
}

// FuncOffSurroundingTextChanged returns the function "WEBEXT.input.ime.onSurroundingTextChanged.removeListener".
func FuncOffSurroundingTextChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)])]) {
	bindings.FuncOffSurroundingTextChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.removeListener" directly.
func OffSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret js.Void) {
	bindings.CallOffSurroundingTextChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSurroundingTextChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSurroundingTextChanged returns true if the function "WEBEXT.input.ime.onSurroundingTextChanged.hasListener" exists.
func HasFuncHasOnSurroundingTextChanged() bool {
	return js.True == bindings.HasFuncHasOnSurroundingTextChanged()
}

// FuncHasOnSurroundingTextChanged returns the function "WEBEXT.input.ime.onSurroundingTextChanged.hasListener".
func FuncHasOnSurroundingTextChanged() (fn js.Func[func(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) bool]) {
	bindings.FuncHasOnSurroundingTextChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.hasListener" directly.
func HasOnSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret bool) {
	bindings.CallHasOnSurroundingTextChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSurroundingTextChanged calls the function "WEBEXT.input.ime.onSurroundingTextChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSurroundingTextChanged(callback js.Func[func(engineID js.String, surroundingInfo *OnSurroundingTextChangedArgSurroundingInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSurroundingTextChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendKeyEvents returns true if the function "WEBEXT.input.ime.sendKeyEvents" exists.
func HasFuncSendKeyEvents() bool {
	return js.True == bindings.HasFuncSendKeyEvents()
}

// FuncSendKeyEvents returns the function "WEBEXT.input.ime.sendKeyEvents".
func FuncSendKeyEvents() (fn js.Func[func(parameters SendKeyEventsArgParameters) js.Promise[js.Void]]) {
	bindings.FuncSendKeyEvents(
		js.Pointer(&fn),
	)
	return
}

// SendKeyEvents calls the function "WEBEXT.input.ime.sendKeyEvents" directly.
func SendKeyEvents(parameters SendKeyEventsArgParameters) (ret js.Promise[js.Void]) {
	bindings.CallSendKeyEvents(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySendKeyEvents calls the function "WEBEXT.input.ime.sendKeyEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendKeyEvents(parameters SendKeyEventsArgParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendKeyEvents(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetAssistiveWindowButtonHighlighted returns true if the function "WEBEXT.input.ime.setAssistiveWindowButtonHighlighted" exists.
func HasFuncSetAssistiveWindowButtonHighlighted() bool {
	return js.True == bindings.HasFuncSetAssistiveWindowButtonHighlighted()
}

// FuncSetAssistiveWindowButtonHighlighted returns the function "WEBEXT.input.ime.setAssistiveWindowButtonHighlighted".
func FuncSetAssistiveWindowButtonHighlighted() (fn js.Func[func(parameters SetAssistiveWindowButtonHighlightedArgParameters) js.Promise[js.Void]]) {
	bindings.FuncSetAssistiveWindowButtonHighlighted(
		js.Pointer(&fn),
	)
	return
}

// SetAssistiveWindowButtonHighlighted calls the function "WEBEXT.input.ime.setAssistiveWindowButtonHighlighted" directly.
func SetAssistiveWindowButtonHighlighted(parameters SetAssistiveWindowButtonHighlightedArgParameters) (ret js.Promise[js.Void]) {
	bindings.CallSetAssistiveWindowButtonHighlighted(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetAssistiveWindowButtonHighlighted calls the function "WEBEXT.input.ime.setAssistiveWindowButtonHighlighted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAssistiveWindowButtonHighlighted(parameters SetAssistiveWindowButtonHighlightedArgParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAssistiveWindowButtonHighlighted(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetAssistiveWindowProperties returns true if the function "WEBEXT.input.ime.setAssistiveWindowProperties" exists.
func HasFuncSetAssistiveWindowProperties() bool {
	return js.True == bindings.HasFuncSetAssistiveWindowProperties()
}

// FuncSetAssistiveWindowProperties returns the function "WEBEXT.input.ime.setAssistiveWindowProperties".
func FuncSetAssistiveWindowProperties() (fn js.Func[func(parameters SetAssistiveWindowPropertiesArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetAssistiveWindowProperties(
		js.Pointer(&fn),
	)
	return
}

// SetAssistiveWindowProperties calls the function "WEBEXT.input.ime.setAssistiveWindowProperties" directly.
func SetAssistiveWindowProperties(parameters SetAssistiveWindowPropertiesArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetAssistiveWindowProperties(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetAssistiveWindowProperties calls the function "WEBEXT.input.ime.setAssistiveWindowProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAssistiveWindowProperties(parameters SetAssistiveWindowPropertiesArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAssistiveWindowProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetCandidateWindowProperties returns true if the function "WEBEXT.input.ime.setCandidateWindowProperties" exists.
func HasFuncSetCandidateWindowProperties() bool {
	return js.True == bindings.HasFuncSetCandidateWindowProperties()
}

// FuncSetCandidateWindowProperties returns the function "WEBEXT.input.ime.setCandidateWindowProperties".
func FuncSetCandidateWindowProperties() (fn js.Func[func(parameters SetCandidateWindowPropertiesArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetCandidateWindowProperties(
		js.Pointer(&fn),
	)
	return
}

// SetCandidateWindowProperties calls the function "WEBEXT.input.ime.setCandidateWindowProperties" directly.
func SetCandidateWindowProperties(parameters SetCandidateWindowPropertiesArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetCandidateWindowProperties(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetCandidateWindowProperties calls the function "WEBEXT.input.ime.setCandidateWindowProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCandidateWindowProperties(parameters SetCandidateWindowPropertiesArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCandidateWindowProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetCandidates returns true if the function "WEBEXT.input.ime.setCandidates" exists.
func HasFuncSetCandidates() bool {
	return js.True == bindings.HasFuncSetCandidates()
}

// FuncSetCandidates returns the function "WEBEXT.input.ime.setCandidates".
func FuncSetCandidates() (fn js.Func[func(parameters SetCandidatesArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetCandidates(
		js.Pointer(&fn),
	)
	return
}

// SetCandidates calls the function "WEBEXT.input.ime.setCandidates" directly.
func SetCandidates(parameters SetCandidatesArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetCandidates(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetCandidates calls the function "WEBEXT.input.ime.setCandidates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCandidates(parameters SetCandidatesArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCandidates(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetComposition returns true if the function "WEBEXT.input.ime.setComposition" exists.
func HasFuncSetComposition() bool {
	return js.True == bindings.HasFuncSetComposition()
}

// FuncSetComposition returns the function "WEBEXT.input.ime.setComposition".
func FuncSetComposition() (fn js.Func[func(parameters SetCompositionArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetComposition(
		js.Pointer(&fn),
	)
	return
}

// SetComposition calls the function "WEBEXT.input.ime.setComposition" directly.
func SetComposition(parameters SetCompositionArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetComposition(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetComposition calls the function "WEBEXT.input.ime.setComposition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetComposition(parameters SetCompositionArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetComposition(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetCursorPosition returns true if the function "WEBEXT.input.ime.setCursorPosition" exists.
func HasFuncSetCursorPosition() bool {
	return js.True == bindings.HasFuncSetCursorPosition()
}

// FuncSetCursorPosition returns the function "WEBEXT.input.ime.setCursorPosition".
func FuncSetCursorPosition() (fn js.Func[func(parameters SetCursorPositionArgParameters) js.Promise[js.Boolean]]) {
	bindings.FuncSetCursorPosition(
		js.Pointer(&fn),
	)
	return
}

// SetCursorPosition calls the function "WEBEXT.input.ime.setCursorPosition" directly.
func SetCursorPosition(parameters SetCursorPositionArgParameters) (ret js.Promise[js.Boolean]) {
	bindings.CallSetCursorPosition(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetCursorPosition calls the function "WEBEXT.input.ime.setCursorPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCursorPosition(parameters SetCursorPositionArgParameters) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCursorPosition(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncSetMenuItems returns true if the function "WEBEXT.input.ime.setMenuItems" exists.
func HasFuncSetMenuItems() bool {
	return js.True == bindings.HasFuncSetMenuItems()
}

// FuncSetMenuItems returns the function "WEBEXT.input.ime.setMenuItems".
func FuncSetMenuItems() (fn js.Func[func(parameters MenuParameters) js.Promise[js.Void]]) {
	bindings.FuncSetMenuItems(
		js.Pointer(&fn),
	)
	return
}

// SetMenuItems calls the function "WEBEXT.input.ime.setMenuItems" directly.
func SetMenuItems(parameters MenuParameters) (ret js.Promise[js.Void]) {
	bindings.CallSetMenuItems(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetMenuItems calls the function "WEBEXT.input.ime.setMenuItems"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMenuItems(parameters MenuParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMenuItems(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// HasFuncUpdateMenuItems returns true if the function "WEBEXT.input.ime.updateMenuItems" exists.
func HasFuncUpdateMenuItems() bool {
	return js.True == bindings.HasFuncUpdateMenuItems()
}

// FuncUpdateMenuItems returns the function "WEBEXT.input.ime.updateMenuItems".
func FuncUpdateMenuItems() (fn js.Func[func(parameters MenuParameters) js.Promise[js.Void]]) {
	bindings.FuncUpdateMenuItems(
		js.Pointer(&fn),
	)
	return
}

// UpdateMenuItems calls the function "WEBEXT.input.ime.updateMenuItems" directly.
func UpdateMenuItems(parameters MenuParameters) (ret js.Promise[js.Void]) {
	bindings.CallUpdateMenuItems(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TryUpdateMenuItems calls the function "WEBEXT.input.ime.updateMenuItems"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateMenuItems(parameters MenuParameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateMenuItems(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package developerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/developerprivate/bindings"
)

type AccessModifier struct {
	// IsEnabled is "AccessModifier.isEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsEnabled MUST be set to true to make this field effective.
	IsEnabled bool
	// IsActive is "AccessModifier.isActive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsActive MUST be set to true to make this field effective.
	IsActive bool

	FFI_USE_IsEnabled bool // for IsEnabled.
	FFI_USE_IsActive  bool // for IsActive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccessModifier with all fields set.
func (p AccessModifier) FromRef(ref js.Ref) AccessModifier {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AccessModifier in the application heap.
func (p AccessModifier) New() js.Ref {
	return bindings.AccessModifierJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AccessModifier) UpdateFrom(ref js.Ref) {
	bindings.AccessModifierJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AccessModifier) Update(ref js.Ref) {
	bindings.AccessModifierJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AccessModifier) FreeMembers(recursive bool) {
}

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

type CommandScope uint32

const (
	_ CommandScope = iota

	CommandScope_GLOBAL
	CommandScope_CHROME
)

func (CommandScope) FromRef(str js.Ref) CommandScope {
	return CommandScope(bindings.ConstOfCommandScope(str))
}

func (x CommandScope) String() (string, bool) {
	switch x {
	case CommandScope_GLOBAL:
		return "GLOBAL", true
	case CommandScope_CHROME:
		return "CHROME", true
	default:
		return "", false
	}
}

type Command struct {
	// Description is "Command.description"
	//
	// Optional
	Description js.String
	// Keybinding is "Command.keybinding"
	//
	// Optional
	Keybinding js.String
	// Name is "Command.name"
	//
	// Optional
	Name js.String
	// IsActive is "Command.isActive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsActive MUST be set to true to make this field effective.
	IsActive bool
	// Scope is "Command.scope"
	//
	// Optional
	Scope CommandScope
	// IsExtensionAction is "Command.isExtensionAction"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsExtensionAction MUST be set to true to make this field effective.
	IsExtensionAction bool

	FFI_USE_IsActive          bool // for IsActive.
	FFI_USE_IsExtensionAction bool // for IsExtensionAction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Command with all fields set.
func (p Command) FromRef(ref js.Ref) Command {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Command in the application heap.
func (p Command) New() js.Ref {
	return bindings.CommandJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Command) UpdateFrom(ref js.Ref) {
	bindings.CommandJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Command) Update(ref js.Ref) {
	bindings.CommandJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Command) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Keybinding.Ref(),
		p.Name.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Keybinding = p.Keybinding.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type ControlledInfo struct {
	// Text is "ControlledInfo.text"
	//
	// Optional
	Text js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ControlledInfo with all fields set.
func (p ControlledInfo) FromRef(ref js.Ref) ControlledInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ControlledInfo in the application heap.
func (p ControlledInfo) New() js.Ref {
	return bindings.ControlledInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ControlledInfo) UpdateFrom(ref js.Ref) {
	bindings.ControlledInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ControlledInfo) Update(ref js.Ref) {
	bindings.ControlledInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ControlledInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

type ErrorType uint32

const (
	_ ErrorType = iota

	ErrorType_MANIFEST
	ErrorType_RUNTIME
)

func (ErrorType) FromRef(str js.Ref) ErrorType {
	return ErrorType(bindings.ConstOfErrorType(str))
}

func (x ErrorType) String() (string, bool) {
	switch x {
	case ErrorType_MANIFEST:
		return "MANIFEST", true
	case ErrorType_RUNTIME:
		return "RUNTIME", true
	default:
		return "", false
	}
}

type DeleteExtensionErrorsProperties struct {
	// ExtensionId is "DeleteExtensionErrorsProperties.extensionId"
	//
	// Optional
	ExtensionId js.String
	// ErrorIds is "DeleteExtensionErrorsProperties.errorIds"
	//
	// Optional
	ErrorIds js.Array[int32]
	// Type is "DeleteExtensionErrorsProperties.type"
	//
	// Optional
	Type ErrorType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteExtensionErrorsProperties with all fields set.
func (p DeleteExtensionErrorsProperties) FromRef(ref js.Ref) DeleteExtensionErrorsProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteExtensionErrorsProperties in the application heap.
func (p DeleteExtensionErrorsProperties) New() js.Ref {
	return bindings.DeleteExtensionErrorsPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteExtensionErrorsProperties) UpdateFrom(ref js.Ref) {
	bindings.DeleteExtensionErrorsPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteExtensionErrorsProperties) Update(ref js.Ref) {
	bindings.DeleteExtensionErrorsPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteExtensionErrorsProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.ErrorIds.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.ErrorIds = p.ErrorIds.FromRef(js.Undefined)
}

type DependentExtension struct {
	// Id is "DependentExtension.id"
	//
	// Optional
	Id js.String
	// Name is "DependentExtension.name"
	//
	// Optional
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DependentExtension with all fields set.
func (p DependentExtension) FromRef(ref js.Ref) DependentExtension {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DependentExtension in the application heap.
func (p DependentExtension) New() js.Ref {
	return bindings.DependentExtensionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DependentExtension) UpdateFrom(ref js.Ref) {
	bindings.DependentExtensionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DependentExtension) Update(ref js.Ref) {
	bindings.DependentExtensionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DependentExtension) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type DisableReasons struct {
	// SuspiciousInstall is "DisableReasons.suspiciousInstall"
	//
	// Optional
	//
	// NOTE: FFI_USE_SuspiciousInstall MUST be set to true to make this field effective.
	SuspiciousInstall bool
	// CorruptInstall is "DisableReasons.corruptInstall"
	//
	// Optional
	//
	// NOTE: FFI_USE_CorruptInstall MUST be set to true to make this field effective.
	CorruptInstall bool
	// UpdateRequired is "DisableReasons.updateRequired"
	//
	// Optional
	//
	// NOTE: FFI_USE_UpdateRequired MUST be set to true to make this field effective.
	UpdateRequired bool
	// PublishedInStoreRequired is "DisableReasons.publishedInStoreRequired"
	//
	// Optional
	//
	// NOTE: FFI_USE_PublishedInStoreRequired MUST be set to true to make this field effective.
	PublishedInStoreRequired bool
	// BlockedByPolicy is "DisableReasons.blockedByPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_BlockedByPolicy MUST be set to true to make this field effective.
	BlockedByPolicy bool
	// Reloading is "DisableReasons.reloading"
	//
	// Optional
	//
	// NOTE: FFI_USE_Reloading MUST be set to true to make this field effective.
	Reloading bool
	// CustodianApprovalRequired is "DisableReasons.custodianApprovalRequired"
	//
	// Optional
	//
	// NOTE: FFI_USE_CustodianApprovalRequired MUST be set to true to make this field effective.
	CustodianApprovalRequired bool
	// ParentDisabledPermissions is "DisableReasons.parentDisabledPermissions"
	//
	// Optional
	//
	// NOTE: FFI_USE_ParentDisabledPermissions MUST be set to true to make this field effective.
	ParentDisabledPermissions bool

	FFI_USE_SuspiciousInstall         bool // for SuspiciousInstall.
	FFI_USE_CorruptInstall            bool // for CorruptInstall.
	FFI_USE_UpdateRequired            bool // for UpdateRequired.
	FFI_USE_PublishedInStoreRequired  bool // for PublishedInStoreRequired.
	FFI_USE_BlockedByPolicy           bool // for BlockedByPolicy.
	FFI_USE_Reloading                 bool // for Reloading.
	FFI_USE_CustodianApprovalRequired bool // for CustodianApprovalRequired.
	FFI_USE_ParentDisabledPermissions bool // for ParentDisabledPermissions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisableReasons with all fields set.
func (p DisableReasons) FromRef(ref js.Ref) DisableReasons {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisableReasons in the application heap.
func (p DisableReasons) New() js.Ref {
	return bindings.DisableReasonsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisableReasons) UpdateFrom(ref js.Ref) {
	bindings.DisableReasonsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisableReasons) Update(ref js.Ref) {
	bindings.DisableReasonsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisableReasons) FreeMembers(recursive bool) {
}

type DragInstallInProgressCallbackFunc func(this js.Ref, loadGuid js.String) js.Ref

func (fn DragInstallInProgressCallbackFunc) Register() js.Func[func(loadGuid js.String)] {
	return js.RegisterCallback[func(loadGuid js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DragInstallInProgressCallbackFunc) DispatchCallback(
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

type DragInstallInProgressCallback[T any] struct {
	Fn  func(arg T, this js.Ref, loadGuid js.String) js.Ref
	Arg T
}

func (cb *DragInstallInProgressCallback[T]) Register() js.Func[func(loadGuid js.String)] {
	return js.RegisterCallback[func(loadGuid js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DragInstallInProgressCallback[T]) DispatchCallback(
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

type ErrorFileSource struct {
	// BeforeHighlight is "ErrorFileSource.beforeHighlight"
	//
	// Optional
	BeforeHighlight js.String
	// Highlight is "ErrorFileSource.highlight"
	//
	// Optional
	Highlight js.String
	// AfterHighlight is "ErrorFileSource.afterHighlight"
	//
	// Optional
	AfterHighlight js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ErrorFileSource with all fields set.
func (p ErrorFileSource) FromRef(ref js.Ref) ErrorFileSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ErrorFileSource in the application heap.
func (p ErrorFileSource) New() js.Ref {
	return bindings.ErrorFileSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ErrorFileSource) UpdateFrom(ref js.Ref) {
	bindings.ErrorFileSourceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ErrorFileSource) Update(ref js.Ref) {
	bindings.ErrorFileSourceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ErrorFileSource) FreeMembers(recursive bool) {
	js.Free(
		p.BeforeHighlight.Ref(),
		p.Highlight.Ref(),
		p.AfterHighlight.Ref(),
	)
	p.BeforeHighlight = p.BeforeHighlight.FromRef(js.Undefined)
	p.Highlight = p.Highlight.FromRef(js.Undefined)
	p.AfterHighlight = p.AfterHighlight.FromRef(js.Undefined)
}

type ErrorLevel uint32

const (
	_ ErrorLevel = iota

	ErrorLevel_LOG
	ErrorLevel_WARN
	ErrorLevel_ERROR
)

func (ErrorLevel) FromRef(str js.Ref) ErrorLevel {
	return ErrorLevel(bindings.ConstOfErrorLevel(str))
}

func (x ErrorLevel) String() (string, bool) {
	switch x {
	case ErrorLevel_LOG:
		return "LOG", true
	case ErrorLevel_WARN:
		return "WARN", true
	case ErrorLevel_ERROR:
		return "ERROR", true
	default:
		return "", false
	}
}

type EventType uint32

const (
	_ EventType = iota

	EventType_INSTALLED
	EventType_UNINSTALLED
	EventType_LOADED
	EventType_UNLOADED
	EventType_VIEW_REGISTERED
	EventType_VIEW_UNREGISTERED
	EventType_ERROR_ADDED
	EventType_ERRORS_REMOVED
	EventType_PREFS_CHANGED
	EventType_WARNINGS_CHANGED
	EventType_COMMAND_ADDED
	EventType_COMMAND_REMOVED
	EventType_PERMISSIONS_CHANGED
	EventType_SERVICE_WORKER_STARTED
	EventType_SERVICE_WORKER_STOPPED
	EventType_CONFIGURATION_CHANGED
	EventType_PINNED_ACTIONS_CHANGED
)

func (EventType) FromRef(str js.Ref) EventType {
	return EventType(bindings.ConstOfEventType(str))
}

func (x EventType) String() (string, bool) {
	switch x {
	case EventType_INSTALLED:
		return "INSTALLED", true
	case EventType_UNINSTALLED:
		return "UNINSTALLED", true
	case EventType_LOADED:
		return "LOADED", true
	case EventType_UNLOADED:
		return "UNLOADED", true
	case EventType_VIEW_REGISTERED:
		return "VIEW_REGISTERED", true
	case EventType_VIEW_UNREGISTERED:
		return "VIEW_UNREGISTERED", true
	case EventType_ERROR_ADDED:
		return "ERROR_ADDED", true
	case EventType_ERRORS_REMOVED:
		return "ERRORS_REMOVED", true
	case EventType_PREFS_CHANGED:
		return "PREFS_CHANGED", true
	case EventType_WARNINGS_CHANGED:
		return "WARNINGS_CHANGED", true
	case EventType_COMMAND_ADDED:
		return "COMMAND_ADDED", true
	case EventType_COMMAND_REMOVED:
		return "COMMAND_REMOVED", true
	case EventType_PERMISSIONS_CHANGED:
		return "PERMISSIONS_CHANGED", true
	case EventType_SERVICE_WORKER_STARTED:
		return "SERVICE_WORKER_STARTED", true
	case EventType_SERVICE_WORKER_STOPPED:
		return "SERVICE_WORKER_STOPPED", true
	case EventType_CONFIGURATION_CHANGED:
		return "CONFIGURATION_CHANGED", true
	case EventType_PINNED_ACTIONS_CHANGED:
		return "PINNED_ACTIONS_CHANGED", true
	default:
		return "", false
	}
}

type SafetyCheckStrings struct {
	// PanelString is "SafetyCheckStrings.panelString"
	//
	// Optional
	PanelString js.String
	// DetailString is "SafetyCheckStrings.detailString"
	//
	// Optional
	DetailString js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SafetyCheckStrings with all fields set.
func (p SafetyCheckStrings) FromRef(ref js.Ref) SafetyCheckStrings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SafetyCheckStrings in the application heap.
func (p SafetyCheckStrings) New() js.Ref {
	return bindings.SafetyCheckStringsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SafetyCheckStrings) UpdateFrom(ref js.Ref) {
	bindings.SafetyCheckStringsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SafetyCheckStrings) Update(ref js.Ref) {
	bindings.SafetyCheckStringsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SafetyCheckStrings) FreeMembers(recursive bool) {
	js.Free(
		p.PanelString.Ref(),
		p.DetailString.Ref(),
	)
	p.PanelString = p.PanelString.FromRef(js.Undefined)
	p.DetailString = p.DetailString.FromRef(js.Undefined)
}

type HomePage struct {
	// Url is "HomePage.url"
	//
	// Optional
	Url js.String
	// Specified is "HomePage.specified"
	//
	// Optional
	//
	// NOTE: FFI_USE_Specified MUST be set to true to make this field effective.
	Specified bool

	FFI_USE_Specified bool // for Specified.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HomePage with all fields set.
func (p HomePage) FromRef(ref js.Ref) HomePage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HomePage in the application heap.
func (p HomePage) New() js.Ref {
	return bindings.HomePageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HomePage) UpdateFrom(ref js.Ref) {
	bindings.HomePageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HomePage) Update(ref js.Ref) {
	bindings.HomePageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HomePage) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type Location uint32

const (
	_ Location = iota

	Location_FROM_STORE
	Location_UNPACKED
	Location_THIRD_PARTY
	Location_INSTALLED_BY_DEFAULT
	Location_UNKNOWN
)

func (Location) FromRef(str js.Ref) Location {
	return Location(bindings.ConstOfLocation(str))
}

func (x Location) String() (string, bool) {
	switch x {
	case Location_FROM_STORE:
		return "FROM_STORE", true
	case Location_UNPACKED:
		return "UNPACKED", true
	case Location_THIRD_PARTY:
		return "THIRD_PARTY", true
	case Location_INSTALLED_BY_DEFAULT:
		return "INSTALLED_BY_DEFAULT", true
	case Location_UNKNOWN:
		return "UNKNOWN", true
	default:
		return "", false
	}
}

type ManifestError struct {
	// Type is "ManifestError.type"
	//
	// Optional
	Type ErrorType
	// ExtensionId is "ManifestError.extensionId"
	//
	// Optional
	ExtensionId js.String
	// FromIncognito is "ManifestError.fromIncognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_FromIncognito MUST be set to true to make this field effective.
	FromIncognito bool
	// Source is "ManifestError.source"
	//
	// Optional
	Source js.String
	// Message is "ManifestError.message"
	//
	// Optional
	Message js.String
	// Id is "ManifestError.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// ManifestKey is "ManifestError.manifestKey"
	//
	// Optional
	ManifestKey js.String
	// ManifestSpecific is "ManifestError.manifestSpecific"
	//
	// Optional
	ManifestSpecific js.String

	FFI_USE_FromIncognito bool // for FromIncognito.
	FFI_USE_Id            bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManifestError with all fields set.
func (p ManifestError) FromRef(ref js.Ref) ManifestError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManifestError in the application heap.
func (p ManifestError) New() js.Ref {
	return bindings.ManifestErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManifestError) UpdateFrom(ref js.Ref) {
	bindings.ManifestErrorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManifestError) Update(ref js.Ref) {
	bindings.ManifestErrorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManifestError) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.Source.Ref(),
		p.Message.Ref(),
		p.ManifestKey.Ref(),
		p.ManifestSpecific.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
	p.ManifestKey = p.ManifestKey.FromRef(js.Undefined)
	p.ManifestSpecific = p.ManifestSpecific.FromRef(js.Undefined)
}

type OptionsPage struct {
	// OpenInTab is "OptionsPage.openInTab"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenInTab MUST be set to true to make this field effective.
	OpenInTab bool
	// Url is "OptionsPage.url"
	//
	// Optional
	Url js.String

	FFI_USE_OpenInTab bool // for OpenInTab.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OptionsPage with all fields set.
func (p OptionsPage) FromRef(ref js.Ref) OptionsPage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OptionsPage in the application heap.
func (p OptionsPage) New() js.Ref {
	return bindings.OptionsPageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OptionsPage) UpdateFrom(ref js.Ref) {
	bindings.OptionsPageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OptionsPage) Update(ref js.Ref) {
	bindings.OptionsPageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OptionsPage) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type Permission struct {
	// Message is "Permission.message"
	//
	// Optional
	Message js.String
	// Submessages is "Permission.submessages"
	//
	// Optional
	Submessages js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Permission with all fields set.
func (p Permission) FromRef(ref js.Ref) Permission {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Permission in the application heap.
func (p Permission) New() js.Ref {
	return bindings.PermissionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Permission) UpdateFrom(ref js.Ref) {
	bindings.PermissionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Permission) Update(ref js.Ref) {
	bindings.PermissionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Permission) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
		p.Submessages.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
	p.Submessages = p.Submessages.FromRef(js.Undefined)
}

type HostAccess uint32

const (
	_ HostAccess = iota

	HostAccess_ON_CLICK
	HostAccess_ON_SPECIFIC_SITES
	HostAccess_ON_ALL_SITES
)

func (HostAccess) FromRef(str js.Ref) HostAccess {
	return HostAccess(bindings.ConstOfHostAccess(str))
}

func (x HostAccess) String() (string, bool) {
	switch x {
	case HostAccess_ON_CLICK:
		return "ON_CLICK", true
	case HostAccess_ON_SPECIFIC_SITES:
		return "ON_SPECIFIC_SITES", true
	case HostAccess_ON_ALL_SITES:
		return "ON_ALL_SITES", true
	default:
		return "", false
	}
}

type SiteControl struct {
	// Host is "SiteControl.host"
	//
	// Optional
	Host js.String
	// Granted is "SiteControl.granted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Granted MUST be set to true to make this field effective.
	Granted bool

	FFI_USE_Granted bool // for Granted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SiteControl with all fields set.
func (p SiteControl) FromRef(ref js.Ref) SiteControl {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SiteControl in the application heap.
func (p SiteControl) New() js.Ref {
	return bindings.SiteControlJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SiteControl) UpdateFrom(ref js.Ref) {
	bindings.SiteControlJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SiteControl) Update(ref js.Ref) {
	bindings.SiteControlJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SiteControl) FreeMembers(recursive bool) {
	js.Free(
		p.Host.Ref(),
	)
	p.Host = p.Host.FromRef(js.Undefined)
}

type RuntimeHostPermissions struct {
	// HasAllHosts is "RuntimeHostPermissions.hasAllHosts"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasAllHosts MUST be set to true to make this field effective.
	HasAllHosts bool
	// HostAccess is "RuntimeHostPermissions.hostAccess"
	//
	// Optional
	HostAccess HostAccess
	// Hosts is "RuntimeHostPermissions.hosts"
	//
	// Optional
	Hosts js.Array[SiteControl]

	FFI_USE_HasAllHosts bool // for HasAllHosts.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RuntimeHostPermissions with all fields set.
func (p RuntimeHostPermissions) FromRef(ref js.Ref) RuntimeHostPermissions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RuntimeHostPermissions in the application heap.
func (p RuntimeHostPermissions) New() js.Ref {
	return bindings.RuntimeHostPermissionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RuntimeHostPermissions) UpdateFrom(ref js.Ref) {
	bindings.RuntimeHostPermissionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RuntimeHostPermissions) Update(ref js.Ref) {
	bindings.RuntimeHostPermissionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RuntimeHostPermissions) FreeMembers(recursive bool) {
	js.Free(
		p.Hosts.Ref(),
	)
	p.Hosts = p.Hosts.FromRef(js.Undefined)
}

type Permissions struct {
	// SimplePermissions is "Permissions.simplePermissions"
	//
	// Optional
	SimplePermissions js.Array[Permission]
	// RuntimeHostPermissions is "Permissions.runtimeHostPermissions"
	//
	// Optional
	//
	// NOTE: RuntimeHostPermissions.FFI_USE MUST be set to true to get RuntimeHostPermissions used.
	RuntimeHostPermissions RuntimeHostPermissions
	// CanAccessSiteData is "Permissions.canAccessSiteData"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanAccessSiteData MUST be set to true to make this field effective.
	CanAccessSiteData bool

	FFI_USE_CanAccessSiteData bool // for CanAccessSiteData.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Permissions with all fields set.
func (p Permissions) FromRef(ref js.Ref) Permissions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Permissions in the application heap.
func (p Permissions) New() js.Ref {
	return bindings.PermissionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Permissions) UpdateFrom(ref js.Ref) {
	bindings.PermissionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Permissions) Update(ref js.Ref) {
	bindings.PermissionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Permissions) FreeMembers(recursive bool) {
	js.Free(
		p.SimplePermissions.Ref(),
	)
	p.SimplePermissions = p.SimplePermissions.FromRef(js.Undefined)
	if recursive {
		p.RuntimeHostPermissions.FreeMembers(true)
	}
}

type StackFrame struct {
	// LineNumber is "StackFrame.lineNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_LineNumber MUST be set to true to make this field effective.
	LineNumber int32
	// ColumnNumber is "StackFrame.columnNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_ColumnNumber MUST be set to true to make this field effective.
	ColumnNumber int32
	// Url is "StackFrame.url"
	//
	// Optional
	Url js.String
	// FunctionName is "StackFrame.functionName"
	//
	// Optional
	FunctionName js.String

	FFI_USE_LineNumber   bool // for LineNumber.
	FFI_USE_ColumnNumber bool // for ColumnNumber.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StackFrame with all fields set.
func (p StackFrame) FromRef(ref js.Ref) StackFrame {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StackFrame in the application heap.
func (p StackFrame) New() js.Ref {
	return bindings.StackFrameJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StackFrame) UpdateFrom(ref js.Ref) {
	bindings.StackFrameJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StackFrame) Update(ref js.Ref) {
	bindings.StackFrameJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StackFrame) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.FunctionName.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.FunctionName = p.FunctionName.FromRef(js.Undefined)
}

type RuntimeError struct {
	// Type is "RuntimeError.type"
	//
	// Optional
	Type ErrorType
	// ExtensionId is "RuntimeError.extensionId"
	//
	// Optional
	ExtensionId js.String
	// FromIncognito is "RuntimeError.fromIncognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_FromIncognito MUST be set to true to make this field effective.
	FromIncognito bool
	// Source is "RuntimeError.source"
	//
	// Optional
	Source js.String
	// Message is "RuntimeError.message"
	//
	// Optional
	Message js.String
	// Id is "RuntimeError.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Severity is "RuntimeError.severity"
	//
	// Optional
	Severity ErrorLevel
	// ContextUrl is "RuntimeError.contextUrl"
	//
	// Optional
	ContextUrl js.String
	// Occurrences is "RuntimeError.occurrences"
	//
	// Optional
	//
	// NOTE: FFI_USE_Occurrences MUST be set to true to make this field effective.
	Occurrences int32
	// RenderViewId is "RuntimeError.renderViewId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderViewId MUST be set to true to make this field effective.
	RenderViewId int32
	// RenderProcessId is "RuntimeError.renderProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderProcessId MUST be set to true to make this field effective.
	RenderProcessId int32
	// CanInspect is "RuntimeError.canInspect"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanInspect MUST be set to true to make this field effective.
	CanInspect bool
	// StackTrace is "RuntimeError.stackTrace"
	//
	// Optional
	StackTrace js.Array[StackFrame]

	FFI_USE_FromIncognito   bool // for FromIncognito.
	FFI_USE_Id              bool // for Id.
	FFI_USE_Occurrences     bool // for Occurrences.
	FFI_USE_RenderViewId    bool // for RenderViewId.
	FFI_USE_RenderProcessId bool // for RenderProcessId.
	FFI_USE_CanInspect      bool // for CanInspect.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RuntimeError with all fields set.
func (p RuntimeError) FromRef(ref js.Ref) RuntimeError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RuntimeError in the application heap.
func (p RuntimeError) New() js.Ref {
	return bindings.RuntimeErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RuntimeError) UpdateFrom(ref js.Ref) {
	bindings.RuntimeErrorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RuntimeError) Update(ref js.Ref) {
	bindings.RuntimeErrorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RuntimeError) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.Source.Ref(),
		p.Message.Ref(),
		p.ContextUrl.Ref(),
		p.StackTrace.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
	p.ContextUrl = p.ContextUrl.FromRef(js.Undefined)
	p.StackTrace = p.StackTrace.FromRef(js.Undefined)
}

type ExtensionState uint32

const (
	_ ExtensionState = iota

	ExtensionState_ENABLED
	ExtensionState_DISABLED
	ExtensionState_TERMINATED
	ExtensionState_BLACKLISTED
)

func (ExtensionState) FromRef(str js.Ref) ExtensionState {
	return ExtensionState(bindings.ConstOfExtensionState(str))
}

func (x ExtensionState) String() (string, bool) {
	switch x {
	case ExtensionState_ENABLED:
		return "ENABLED", true
	case ExtensionState_DISABLED:
		return "DISABLED", true
	case ExtensionState_TERMINATED:
		return "TERMINATED", true
	case ExtensionState_BLACKLISTED:
		return "BLACKLISTED", true
	default:
		return "", false
	}
}

type ExtensionType uint32

const (
	_ ExtensionType = iota

	ExtensionType_HOSTED_APP
	ExtensionType_PLATFORM_APP
	ExtensionType_LEGACY_PACKAGED_APP
	ExtensionType_EXTENSION
	ExtensionType_THEME
	ExtensionType_SHARED_MODULE
)

func (ExtensionType) FromRef(str js.Ref) ExtensionType {
	return ExtensionType(bindings.ConstOfExtensionType(str))
}

func (x ExtensionType) String() (string, bool) {
	switch x {
	case ExtensionType_HOSTED_APP:
		return "HOSTED_APP", true
	case ExtensionType_PLATFORM_APP:
		return "PLATFORM_APP", true
	case ExtensionType_LEGACY_PACKAGED_APP:
		return "LEGACY_PACKAGED_APP", true
	case ExtensionType_EXTENSION:
		return "EXTENSION", true
	case ExtensionType_THEME:
		return "THEME", true
	case ExtensionType_SHARED_MODULE:
		return "SHARED_MODULE", true
	default:
		return "", false
	}
}

type ViewType uint32

const (
	_ ViewType = iota

	ViewType_APP_WINDOW
	ViewType_BACKGROUND_CONTENTS
	ViewType_COMPONENT
	ViewType_EXTENSION_BACKGROUND_PAGE
	ViewType_EXTENSION_DIALOG
	ViewType_EXTENSION_GUEST
	ViewType_EXTENSION_POPUP
	ViewType_EXTENSION_SERVICE_WORKER_BACKGROUND
	ViewType_TAB_CONTENTS
	ViewType_OFFSCREEN_DOCUMENT
	ViewType_EXTENSION_SIDE_PANEL
)

func (ViewType) FromRef(str js.Ref) ViewType {
	return ViewType(bindings.ConstOfViewType(str))
}

func (x ViewType) String() (string, bool) {
	switch x {
	case ViewType_APP_WINDOW:
		return "APP_WINDOW", true
	case ViewType_BACKGROUND_CONTENTS:
		return "BACKGROUND_CONTENTS", true
	case ViewType_COMPONENT:
		return "COMPONENT", true
	case ViewType_EXTENSION_BACKGROUND_PAGE:
		return "EXTENSION_BACKGROUND_PAGE", true
	case ViewType_EXTENSION_DIALOG:
		return "EXTENSION_DIALOG", true
	case ViewType_EXTENSION_GUEST:
		return "EXTENSION_GUEST", true
	case ViewType_EXTENSION_POPUP:
		return "EXTENSION_POPUP", true
	case ViewType_EXTENSION_SERVICE_WORKER_BACKGROUND:
		return "EXTENSION_SERVICE_WORKER_BACKGROUND", true
	case ViewType_TAB_CONTENTS:
		return "TAB_CONTENTS", true
	case ViewType_OFFSCREEN_DOCUMENT:
		return "OFFSCREEN_DOCUMENT", true
	case ViewType_EXTENSION_SIDE_PANEL:
		return "EXTENSION_SIDE_PANEL", true
	default:
		return "", false
	}
}

type ExtensionView struct {
	// Url is "ExtensionView.url"
	//
	// Optional
	Url js.String
	// RenderProcessId is "ExtensionView.renderProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderProcessId MUST be set to true to make this field effective.
	RenderProcessId int32
	// RenderViewId is "ExtensionView.renderViewId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderViewId MUST be set to true to make this field effective.
	RenderViewId int32
	// Incognito is "ExtensionView.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// IsIframe is "ExtensionView.isIframe"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsIframe MUST be set to true to make this field effective.
	IsIframe bool
	// Type is "ExtensionView.type"
	//
	// Optional
	Type ViewType

	FFI_USE_RenderProcessId bool // for RenderProcessId.
	FFI_USE_RenderViewId    bool // for RenderViewId.
	FFI_USE_Incognito       bool // for Incognito.
	FFI_USE_IsIframe        bool // for IsIframe.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionView with all fields set.
func (p ExtensionView) FromRef(ref js.Ref) ExtensionView {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionView in the application heap.
func (p ExtensionView) New() js.Ref {
	return bindings.ExtensionViewJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionView) UpdateFrom(ref js.Ref) {
	bindings.ExtensionViewJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionView) Update(ref js.Ref) {
	bindings.ExtensionViewJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionView) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type ExtensionInfo struct {
	// BlacklistText is "ExtensionInfo.blacklistText"
	//
	// Optional
	BlacklistText js.String
	// SafetyCheckText is "ExtensionInfo.safetyCheckText"
	//
	// Optional
	//
	// NOTE: SafetyCheckText.FFI_USE MUST be set to true to get SafetyCheckText used.
	SafetyCheckText SafetyCheckStrings
	// Commands is "ExtensionInfo.commands"
	//
	// Optional
	Commands js.Array[Command]
	// ControlledInfo is "ExtensionInfo.controlledInfo"
	//
	// Optional
	//
	// NOTE: ControlledInfo.FFI_USE MUST be set to true to get ControlledInfo used.
	ControlledInfo ControlledInfo
	// DependentExtensions is "ExtensionInfo.dependentExtensions"
	//
	// Optional
	DependentExtensions js.Array[DependentExtension]
	// Description is "ExtensionInfo.description"
	//
	// Optional
	Description js.String
	// DisableReasons is "ExtensionInfo.disableReasons"
	//
	// Optional
	//
	// NOTE: DisableReasons.FFI_USE MUST be set to true to get DisableReasons used.
	DisableReasons DisableReasons
	// ErrorCollection is "ExtensionInfo.errorCollection"
	//
	// Optional
	//
	// NOTE: ErrorCollection.FFI_USE MUST be set to true to get ErrorCollection used.
	ErrorCollection AccessModifier
	// FileAccess is "ExtensionInfo.fileAccess"
	//
	// Optional
	//
	// NOTE: FileAccess.FFI_USE MUST be set to true to get FileAccess used.
	FileAccess AccessModifier
	// HomePage is "ExtensionInfo.homePage"
	//
	// Optional
	//
	// NOTE: HomePage.FFI_USE MUST be set to true to get HomePage used.
	HomePage HomePage
	// IconUrl is "ExtensionInfo.iconUrl"
	//
	// Optional
	IconUrl js.String
	// Id is "ExtensionInfo.id"
	//
	// Optional
	Id js.String
	// IncognitoAccess is "ExtensionInfo.incognitoAccess"
	//
	// Optional
	//
	// NOTE: IncognitoAccess.FFI_USE MUST be set to true to get IncognitoAccess used.
	IncognitoAccess AccessModifier
	// InstallWarnings is "ExtensionInfo.installWarnings"
	//
	// Optional
	InstallWarnings js.Array[js.String]
	// LaunchUrl is "ExtensionInfo.launchUrl"
	//
	// Optional
	LaunchUrl js.String
	// Location is "ExtensionInfo.location"
	//
	// Optional
	Location Location
	// LocationText is "ExtensionInfo.locationText"
	//
	// Optional
	LocationText js.String
	// ManifestErrors is "ExtensionInfo.manifestErrors"
	//
	// Optional
	ManifestErrors js.Array[ManifestError]
	// ManifestHomePageUrl is "ExtensionInfo.manifestHomePageUrl"
	//
	// Optional
	ManifestHomePageUrl js.String
	// MustRemainInstalled is "ExtensionInfo.mustRemainInstalled"
	//
	// Optional
	//
	// NOTE: FFI_USE_MustRemainInstalled MUST be set to true to make this field effective.
	MustRemainInstalled bool
	// Name is "ExtensionInfo.name"
	//
	// Optional
	Name js.String
	// OfflineEnabled is "ExtensionInfo.offlineEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfflineEnabled MUST be set to true to make this field effective.
	OfflineEnabled bool
	// OptionsPage is "ExtensionInfo.optionsPage"
	//
	// Optional
	//
	// NOTE: OptionsPage.FFI_USE MUST be set to true to get OptionsPage used.
	OptionsPage OptionsPage
	// Path is "ExtensionInfo.path"
	//
	// Optional
	Path js.String
	// Permissions is "ExtensionInfo.permissions"
	//
	// Optional
	//
	// NOTE: Permissions.FFI_USE MUST be set to true to get Permissions used.
	Permissions Permissions
	// PrettifiedPath is "ExtensionInfo.prettifiedPath"
	//
	// Optional
	PrettifiedPath js.String
	// RuntimeErrors is "ExtensionInfo.runtimeErrors"
	//
	// Optional
	RuntimeErrors js.Array[RuntimeError]
	// RuntimeWarnings is "ExtensionInfo.runtimeWarnings"
	//
	// Optional
	RuntimeWarnings js.Array[js.String]
	// State is "ExtensionInfo.state"
	//
	// Optional
	State ExtensionState
	// Type is "ExtensionInfo.type"
	//
	// Optional
	Type ExtensionType
	// UpdateUrl is "ExtensionInfo.updateUrl"
	//
	// Optional
	UpdateUrl js.String
	// UserMayModify is "ExtensionInfo.userMayModify"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserMayModify MUST be set to true to make this field effective.
	UserMayModify bool
	// Version is "ExtensionInfo.version"
	//
	// Optional
	Version js.String
	// Views is "ExtensionInfo.views"
	//
	// Optional
	Views js.Array[ExtensionView]
	// WebStoreUrl is "ExtensionInfo.webStoreUrl"
	//
	// Optional
	WebStoreUrl js.String
	// ShowSafeBrowsingAllowlistWarning is "ExtensionInfo.showSafeBrowsingAllowlistWarning"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowSafeBrowsingAllowlistWarning MUST be set to true to make this field effective.
	ShowSafeBrowsingAllowlistWarning bool
	// ShowAccessRequestsInToolbar is "ExtensionInfo.showAccessRequestsInToolbar"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowAccessRequestsInToolbar MUST be set to true to make this field effective.
	ShowAccessRequestsInToolbar bool
	// AcknowledgeSafetyCheckWarning is "ExtensionInfo.acknowledgeSafetyCheckWarning"
	//
	// Optional
	//
	// NOTE: FFI_USE_AcknowledgeSafetyCheckWarning MUST be set to true to make this field effective.
	AcknowledgeSafetyCheckWarning bool
	// PinnedToToolbar is "ExtensionInfo.pinnedToToolbar"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinnedToToolbar MUST be set to true to make this field effective.
	PinnedToToolbar bool

	FFI_USE_MustRemainInstalled              bool // for MustRemainInstalled.
	FFI_USE_OfflineEnabled                   bool // for OfflineEnabled.
	FFI_USE_UserMayModify                    bool // for UserMayModify.
	FFI_USE_ShowSafeBrowsingAllowlistWarning bool // for ShowSafeBrowsingAllowlistWarning.
	FFI_USE_ShowAccessRequestsInToolbar      bool // for ShowAccessRequestsInToolbar.
	FFI_USE_AcknowledgeSafetyCheckWarning    bool // for AcknowledgeSafetyCheckWarning.
	FFI_USE_PinnedToToolbar                  bool // for PinnedToToolbar.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionInfo with all fields set.
func (p ExtensionInfo) FromRef(ref js.Ref) ExtensionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionInfo in the application heap.
func (p ExtensionInfo) New() js.Ref {
	return bindings.ExtensionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionInfo) UpdateFrom(ref js.Ref) {
	bindings.ExtensionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionInfo) Update(ref js.Ref) {
	bindings.ExtensionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.BlacklistText.Ref(),
		p.Commands.Ref(),
		p.DependentExtensions.Ref(),
		p.Description.Ref(),
		p.IconUrl.Ref(),
		p.Id.Ref(),
		p.InstallWarnings.Ref(),
		p.LaunchUrl.Ref(),
		p.LocationText.Ref(),
		p.ManifestErrors.Ref(),
		p.ManifestHomePageUrl.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.PrettifiedPath.Ref(),
		p.RuntimeErrors.Ref(),
		p.RuntimeWarnings.Ref(),
		p.UpdateUrl.Ref(),
		p.Version.Ref(),
		p.Views.Ref(),
		p.WebStoreUrl.Ref(),
	)
	p.BlacklistText = p.BlacklistText.FromRef(js.Undefined)
	p.Commands = p.Commands.FromRef(js.Undefined)
	p.DependentExtensions = p.DependentExtensions.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.InstallWarnings = p.InstallWarnings.FromRef(js.Undefined)
	p.LaunchUrl = p.LaunchUrl.FromRef(js.Undefined)
	p.LocationText = p.LocationText.FromRef(js.Undefined)
	p.ManifestErrors = p.ManifestErrors.FromRef(js.Undefined)
	p.ManifestHomePageUrl = p.ManifestHomePageUrl.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.PrettifiedPath = p.PrettifiedPath.FromRef(js.Undefined)
	p.RuntimeErrors = p.RuntimeErrors.FromRef(js.Undefined)
	p.RuntimeWarnings = p.RuntimeWarnings.FromRef(js.Undefined)
	p.UpdateUrl = p.UpdateUrl.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.Views = p.Views.FromRef(js.Undefined)
	p.WebStoreUrl = p.WebStoreUrl.FromRef(js.Undefined)
	if recursive {
		p.SafetyCheckText.FreeMembers(true)
		p.ControlledInfo.FreeMembers(true)
		p.DisableReasons.FreeMembers(true)
		p.ErrorCollection.FreeMembers(true)
		p.FileAccess.FreeMembers(true)
		p.HomePage.FreeMembers(true)
		p.IncognitoAccess.FreeMembers(true)
		p.OptionsPage.FreeMembers(true)
		p.Permissions.FreeMembers(true)
	}
}

type EventData struct {
	// EventType is "EventData.event_type"
	//
	// Optional
	EventType EventType
	// ItemId is "EventData.item_id"
	//
	// Optional
	ItemId js.String
	// ExtensionInfo is "EventData.extensionInfo"
	//
	// Optional
	//
	// NOTE: ExtensionInfo.FFI_USE MUST be set to true to get ExtensionInfo used.
	ExtensionInfo ExtensionInfo

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EventData with all fields set.
func (p EventData) FromRef(ref js.Ref) EventData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventData in the application heap.
func (p EventData) New() js.Ref {
	return bindings.EventDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EventData) UpdateFrom(ref js.Ref) {
	bindings.EventDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EventData) Update(ref js.Ref) {
	bindings.EventDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EventData) FreeMembers(recursive bool) {
	js.Free(
		p.ItemId.Ref(),
	)
	p.ItemId = p.ItemId.FromRef(js.Undefined)
	if recursive {
		p.ExtensionInfo.FreeMembers(true)
	}
}

type ExtensionCommandUpdate struct {
	// ExtensionId is "ExtensionCommandUpdate.extensionId"
	//
	// Optional
	ExtensionId js.String
	// CommandName is "ExtensionCommandUpdate.commandName"
	//
	// Optional
	CommandName js.String
	// Scope is "ExtensionCommandUpdate.scope"
	//
	// Optional
	Scope CommandScope
	// Keybinding is "ExtensionCommandUpdate.keybinding"
	//
	// Optional
	Keybinding js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionCommandUpdate with all fields set.
func (p ExtensionCommandUpdate) FromRef(ref js.Ref) ExtensionCommandUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionCommandUpdate in the application heap.
func (p ExtensionCommandUpdate) New() js.Ref {
	return bindings.ExtensionCommandUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionCommandUpdate) UpdateFrom(ref js.Ref) {
	bindings.ExtensionCommandUpdateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionCommandUpdate) Update(ref js.Ref) {
	bindings.ExtensionCommandUpdateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionCommandUpdate) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.CommandName.Ref(),
		p.Keybinding.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.CommandName = p.CommandName.FromRef(js.Undefined)
	p.Keybinding = p.Keybinding.FromRef(js.Undefined)
}

type ExtensionConfigurationUpdate struct {
	// ExtensionId is "ExtensionConfigurationUpdate.extensionId"
	//
	// Optional
	ExtensionId js.String
	// FileAccess is "ExtensionConfigurationUpdate.fileAccess"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileAccess MUST be set to true to make this field effective.
	FileAccess bool
	// IncognitoAccess is "ExtensionConfigurationUpdate.incognitoAccess"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncognitoAccess MUST be set to true to make this field effective.
	IncognitoAccess bool
	// ErrorCollection is "ExtensionConfigurationUpdate.errorCollection"
	//
	// Optional
	//
	// NOTE: FFI_USE_ErrorCollection MUST be set to true to make this field effective.
	ErrorCollection bool
	// HostAccess is "ExtensionConfigurationUpdate.hostAccess"
	//
	// Optional
	HostAccess HostAccess
	// ShowAccessRequestsInToolbar is "ExtensionConfigurationUpdate.showAccessRequestsInToolbar"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowAccessRequestsInToolbar MUST be set to true to make this field effective.
	ShowAccessRequestsInToolbar bool
	// AcknowledgeSafetyCheckWarning is "ExtensionConfigurationUpdate.acknowledgeSafetyCheckWarning"
	//
	// Optional
	//
	// NOTE: FFI_USE_AcknowledgeSafetyCheckWarning MUST be set to true to make this field effective.
	AcknowledgeSafetyCheckWarning bool
	// PinnedToToolbar is "ExtensionConfigurationUpdate.pinnedToToolbar"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinnedToToolbar MUST be set to true to make this field effective.
	PinnedToToolbar bool

	FFI_USE_FileAccess                    bool // for FileAccess.
	FFI_USE_IncognitoAccess               bool // for IncognitoAccess.
	FFI_USE_ErrorCollection               bool // for ErrorCollection.
	FFI_USE_ShowAccessRequestsInToolbar   bool // for ShowAccessRequestsInToolbar.
	FFI_USE_AcknowledgeSafetyCheckWarning bool // for AcknowledgeSafetyCheckWarning.
	FFI_USE_PinnedToToolbar               bool // for PinnedToToolbar.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionConfigurationUpdate with all fields set.
func (p ExtensionConfigurationUpdate) FromRef(ref js.Ref) ExtensionConfigurationUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionConfigurationUpdate in the application heap.
func (p ExtensionConfigurationUpdate) New() js.Ref {
	return bindings.ExtensionConfigurationUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionConfigurationUpdate) UpdateFrom(ref js.Ref) {
	bindings.ExtensionConfigurationUpdateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionConfigurationUpdate) Update(ref js.Ref) {
	bindings.ExtensionConfigurationUpdateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionConfigurationUpdate) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
}

type ExtensionInfoCallbackFunc func(this js.Ref, result *ExtensionInfo) js.Ref

func (fn ExtensionInfoCallbackFunc) Register() js.Func[func(result *ExtensionInfo)] {
	return js.RegisterCallback[func(result *ExtensionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExtensionInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionInfo
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

type ExtensionInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *ExtensionInfo) js.Ref
	Arg T
}

func (cb *ExtensionInfoCallback[T]) Register() js.Func[func(result *ExtensionInfo)] {
	return js.RegisterCallback[func(result *ExtensionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExtensionInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionInfo
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

type ExtensionInfosCallbackFunc func(this js.Ref, result js.Array[ExtensionInfo]) js.Ref

func (fn ExtensionInfosCallbackFunc) Register() js.Func[func(result js.Array[ExtensionInfo])] {
	return js.RegisterCallback[func(result js.Array[ExtensionInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExtensionInfosCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ExtensionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExtensionInfosCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[ExtensionInfo]) js.Ref
	Arg T
}

func (cb *ExtensionInfosCallback[T]) Register() js.Func[func(result js.Array[ExtensionInfo])] {
	return js.RegisterCallback[func(result js.Array[ExtensionInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExtensionInfosCallback[T]) DispatchCallback(
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

		js.Array[ExtensionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExtensionSiteAccessUpdate struct {
	// Id is "ExtensionSiteAccessUpdate.id"
	//
	// Optional
	Id js.String
	// SiteAccess is "ExtensionSiteAccessUpdate.siteAccess"
	//
	// Optional
	SiteAccess HostAccess

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionSiteAccessUpdate with all fields set.
func (p ExtensionSiteAccessUpdate) FromRef(ref js.Ref) ExtensionSiteAccessUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionSiteAccessUpdate in the application heap.
func (p ExtensionSiteAccessUpdate) New() js.Ref {
	return bindings.ExtensionSiteAccessUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionSiteAccessUpdate) UpdateFrom(ref js.Ref) {
	bindings.ExtensionSiteAccessUpdateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionSiteAccessUpdate) Update(ref js.Ref) {
	bindings.ExtensionSiteAccessUpdateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionSiteAccessUpdate) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type FileType uint32

const (
	_ FileType = iota

	FileType_LOAD
	FileType_PEM
)

func (FileType) FromRef(str js.Ref) FileType {
	return FileType(bindings.ConstOfFileType(str))
}

func (x FileType) String() (string, bool) {
	switch x {
	case FileType_LOAD:
		return "LOAD", true
	case FileType_PEM:
		return "PEM", true
	default:
		return "", false
	}
}

type GetExtensionsInfoOptions struct {
	// IncludeDisabled is "GetExtensionsInfoOptions.includeDisabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncludeDisabled MUST be set to true to make this field effective.
	IncludeDisabled bool
	// IncludeTerminated is "GetExtensionsInfoOptions.includeTerminated"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncludeTerminated MUST be set to true to make this field effective.
	IncludeTerminated bool

	FFI_USE_IncludeDisabled   bool // for IncludeDisabled.
	FFI_USE_IncludeTerminated bool // for IncludeTerminated.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetExtensionsInfoOptions with all fields set.
func (p GetExtensionsInfoOptions) FromRef(ref js.Ref) GetExtensionsInfoOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetExtensionsInfoOptions in the application heap.
func (p GetExtensionsInfoOptions) New() js.Ref {
	return bindings.GetExtensionsInfoOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetExtensionsInfoOptions) UpdateFrom(ref js.Ref) {
	bindings.GetExtensionsInfoOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetExtensionsInfoOptions) Update(ref js.Ref) {
	bindings.GetExtensionsInfoOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetExtensionsInfoOptions) FreeMembers(recursive bool) {
}

type GetMatchingExtensionsForSiteCallbackFunc func(this js.Ref, matchingExtensions js.Array[MatchingExtensionInfo]) js.Ref

func (fn GetMatchingExtensionsForSiteCallbackFunc) Register() js.Func[func(matchingExtensions js.Array[MatchingExtensionInfo])] {
	return js.RegisterCallback[func(matchingExtensions js.Array[MatchingExtensionInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetMatchingExtensionsForSiteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[MatchingExtensionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetMatchingExtensionsForSiteCallback[T any] struct {
	Fn  func(arg T, this js.Ref, matchingExtensions js.Array[MatchingExtensionInfo]) js.Ref
	Arg T
}

func (cb *GetMatchingExtensionsForSiteCallback[T]) Register() js.Func[func(matchingExtensions js.Array[MatchingExtensionInfo])] {
	return js.RegisterCallback[func(matchingExtensions js.Array[MatchingExtensionInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetMatchingExtensionsForSiteCallback[T]) DispatchCallback(
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

		js.Array[MatchingExtensionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MatchingExtensionInfo struct {
	// Id is "MatchingExtensionInfo.id"
	//
	// Optional
	Id js.String
	// SiteAccess is "MatchingExtensionInfo.siteAccess"
	//
	// Optional
	SiteAccess HostAccess
	// CanRequestAllSites is "MatchingExtensionInfo.canRequestAllSites"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanRequestAllSites MUST be set to true to make this field effective.
	CanRequestAllSites bool

	FFI_USE_CanRequestAllSites bool // for CanRequestAllSites.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MatchingExtensionInfo with all fields set.
func (p MatchingExtensionInfo) FromRef(ref js.Ref) MatchingExtensionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MatchingExtensionInfo in the application heap.
func (p MatchingExtensionInfo) New() js.Ref {
	return bindings.MatchingExtensionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MatchingExtensionInfo) UpdateFrom(ref js.Ref) {
	bindings.MatchingExtensionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MatchingExtensionInfo) Update(ref js.Ref) {
	bindings.MatchingExtensionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MatchingExtensionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type GetProjectsInfoCallbackFunc func(this js.Ref, result js.Array[ProjectInfo]) js.Ref

func (fn GetProjectsInfoCallbackFunc) Register() js.Func[func(result js.Array[ProjectInfo])] {
	return js.RegisterCallback[func(result js.Array[ProjectInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetProjectsInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ProjectInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetProjectsInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[ProjectInfo]) js.Ref
	Arg T
}

func (cb *GetProjectsInfoCallback[T]) Register() js.Func[func(result js.Array[ProjectInfo])] {
	return js.RegisterCallback[func(result js.Array[ProjectInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetProjectsInfoCallback[T]) DispatchCallback(
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

		js.Array[ProjectInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProjectInfo struct {
	// Name is "ProjectInfo.name"
	//
	// Optional
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProjectInfo with all fields set.
func (p ProjectInfo) FromRef(ref js.Ref) ProjectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProjectInfo in the application heap.
func (p ProjectInfo) New() js.Ref {
	return bindings.ProjectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProjectInfo) UpdateFrom(ref js.Ref) {
	bindings.ProjectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProjectInfo) Update(ref js.Ref) {
	bindings.ProjectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProjectInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type OneOf_String_Int32 struct {
	ref js.Ref
}

func (x OneOf_String_Int32) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Int32) Free() {
	x.ref.Free()
}

func (x OneOf_String_Int32) FromRef(ref js.Ref) OneOf_String_Int32 {
	return OneOf_String_Int32{
		ref: ref,
	}
}

func (x OneOf_String_Int32) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Int32) Int32() int32 {
	return js.Number[int32]{}.FromRef(x.ref).Get()
}

type InspectOptions struct {
	// ExtensionId is "InspectOptions.extension_id"
	//
	// Optional
	ExtensionId js.String
	// RenderProcessId is "InspectOptions.render_process_id"
	//
	// Optional
	RenderProcessId OneOf_String_Int32
	// RenderViewId is "InspectOptions.render_view_id"
	//
	// Optional
	RenderViewId OneOf_String_Int32
	// Incognito is "InspectOptions.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool

	FFI_USE_Incognito bool // for Incognito.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InspectOptions with all fields set.
func (p InspectOptions) FromRef(ref js.Ref) InspectOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InspectOptions in the application heap.
func (p InspectOptions) New() js.Ref {
	return bindings.InspectOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InspectOptions) UpdateFrom(ref js.Ref) {
	bindings.InspectOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InspectOptions) Update(ref js.Ref) {
	bindings.InspectOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InspectOptions) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.RenderProcessId.Ref(),
		p.RenderViewId.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.RenderProcessId = p.RenderProcessId.FromRef(js.Undefined)
	p.RenderViewId = p.RenderViewId.FromRef(js.Undefined)
}

type InstallWarning struct {
	// Message is "InstallWarning.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InstallWarning with all fields set.
func (p InstallWarning) FromRef(ref js.Ref) InstallWarning {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InstallWarning in the application heap.
func (p InstallWarning) New() js.Ref {
	return bindings.InstallWarningJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InstallWarning) UpdateFrom(ref js.Ref) {
	bindings.InstallWarningJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InstallWarning) Update(ref js.Ref) {
	bindings.InstallWarningJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InstallWarning) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type ItemType uint32

const (
	_ ItemType = iota

	ItemType_HOSTED_APP
	ItemType_PACKAGED_APP
	ItemType_LEGACY_PACKAGED_APP
	ItemType_EXTENSION
	ItemType_THEME
)

func (ItemType) FromRef(str js.Ref) ItemType {
	return ItemType(bindings.ConstOfItemType(str))
}

func (x ItemType) String() (string, bool) {
	switch x {
	case ItemType_HOSTED_APP:
		return "hosted_app", true
	case ItemType_PACKAGED_APP:
		return "packaged_app", true
	case ItemType_LEGACY_PACKAGED_APP:
		return "legacy_packaged_app", true
	case ItemType_EXTENSION:
		return "extension", true
	case ItemType_THEME:
		return "theme", true
	default:
		return "", false
	}
}

type ItemInspectView struct {
	// Path is "ItemInspectView.path"
	//
	// Optional
	Path js.String
	// RenderProcessId is "ItemInspectView.render_process_id"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderProcessId MUST be set to true to make this field effective.
	RenderProcessId int32
	// RenderViewId is "ItemInspectView.render_view_id"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderViewId MUST be set to true to make this field effective.
	RenderViewId int32
	// Incognito is "ItemInspectView.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// GeneratedBackgroundPage is "ItemInspectView.generatedBackgroundPage"
	//
	// Optional
	//
	// NOTE: FFI_USE_GeneratedBackgroundPage MUST be set to true to make this field effective.
	GeneratedBackgroundPage bool

	FFI_USE_RenderProcessId         bool // for RenderProcessId.
	FFI_USE_RenderViewId            bool // for RenderViewId.
	FFI_USE_Incognito               bool // for Incognito.
	FFI_USE_GeneratedBackgroundPage bool // for GeneratedBackgroundPage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ItemInspectView with all fields set.
func (p ItemInspectView) FromRef(ref js.Ref) ItemInspectView {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ItemInspectView in the application heap.
func (p ItemInspectView) New() js.Ref {
	return bindings.ItemInspectViewJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ItemInspectView) UpdateFrom(ref js.Ref) {
	bindings.ItemInspectViewJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ItemInspectView) Update(ref js.Ref) {
	bindings.ItemInspectViewJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ItemInspectView) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
}

type ItemInfo struct {
	// Id is "ItemInfo.id"
	//
	// Optional
	Id js.String
	// Name is "ItemInfo.name"
	//
	// Optional
	Name js.String
	// Version is "ItemInfo.version"
	//
	// Optional
	Version js.String
	// Description is "ItemInfo.description"
	//
	// Optional
	Description js.String
	// MayDisable is "ItemInfo.may_disable"
	//
	// Optional
	//
	// NOTE: FFI_USE_MayDisable MUST be set to true to make this field effective.
	MayDisable bool
	// Enabled is "ItemInfo.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// IsApp is "ItemInfo.isApp"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsApp MUST be set to true to make this field effective.
	IsApp bool
	// Type is "ItemInfo.type"
	//
	// Optional
	Type ItemType
	// AllowActivity is "ItemInfo.allow_activity"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowActivity MUST be set to true to make this field effective.
	AllowActivity bool
	// AllowFileAccess is "ItemInfo.allow_file_access"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowFileAccess MUST be set to true to make this field effective.
	AllowFileAccess bool
	// WantsFileAccess is "ItemInfo.wants_file_access"
	//
	// Optional
	//
	// NOTE: FFI_USE_WantsFileAccess MUST be set to true to make this field effective.
	WantsFileAccess bool
	// IncognitoEnabled is "ItemInfo.incognito_enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncognitoEnabled MUST be set to true to make this field effective.
	IncognitoEnabled bool
	// IsUnpacked is "ItemInfo.is_unpacked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUnpacked MUST be set to true to make this field effective.
	IsUnpacked bool
	// AllowReload is "ItemInfo.allow_reload"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowReload MUST be set to true to make this field effective.
	AllowReload bool
	// Terminated is "ItemInfo.terminated"
	//
	// Optional
	//
	// NOTE: FFI_USE_Terminated MUST be set to true to make this field effective.
	Terminated bool
	// AllowIncognito is "ItemInfo.allow_incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowIncognito MUST be set to true to make this field effective.
	AllowIncognito bool
	// IconUrl is "ItemInfo.icon_url"
	//
	// Optional
	IconUrl js.String
	// Path is "ItemInfo.path"
	//
	// Optional
	Path js.String
	// OptionsUrl is "ItemInfo.options_url"
	//
	// Optional
	OptionsUrl js.String
	// AppLaunchUrl is "ItemInfo.app_launch_url"
	//
	// Optional
	AppLaunchUrl js.String
	// HomepageUrl is "ItemInfo.homepage_url"
	//
	// Optional
	HomepageUrl js.String
	// UpdateUrl is "ItemInfo.update_url"
	//
	// Optional
	UpdateUrl js.String
	// InstallWarnings is "ItemInfo.install_warnings"
	//
	// Optional
	InstallWarnings js.Array[InstallWarning]
	// ManifestErrors is "ItemInfo.manifest_errors"
	//
	// Optional
	ManifestErrors js.Array[js.Any]
	// RuntimeErrors is "ItemInfo.runtime_errors"
	//
	// Optional
	RuntimeErrors js.Array[js.Any]
	// OfflineEnabled is "ItemInfo.offline_enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_OfflineEnabled MUST be set to true to make this field effective.
	OfflineEnabled bool
	// Views is "ItemInfo.views"
	//
	// Optional
	Views js.Array[ItemInspectView]

	FFI_USE_MayDisable       bool // for MayDisable.
	FFI_USE_Enabled          bool // for Enabled.
	FFI_USE_IsApp            bool // for IsApp.
	FFI_USE_AllowActivity    bool // for AllowActivity.
	FFI_USE_AllowFileAccess  bool // for AllowFileAccess.
	FFI_USE_WantsFileAccess  bool // for WantsFileAccess.
	FFI_USE_IncognitoEnabled bool // for IncognitoEnabled.
	FFI_USE_IsUnpacked       bool // for IsUnpacked.
	FFI_USE_AllowReload      bool // for AllowReload.
	FFI_USE_Terminated       bool // for Terminated.
	FFI_USE_AllowIncognito   bool // for AllowIncognito.
	FFI_USE_OfflineEnabled   bool // for OfflineEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ItemInfo with all fields set.
func (p ItemInfo) FromRef(ref js.Ref) ItemInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ItemInfo in the application heap.
func (p ItemInfo) New() js.Ref {
	return bindings.ItemInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ItemInfo) UpdateFrom(ref js.Ref) {
	bindings.ItemInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ItemInfo) Update(ref js.Ref) {
	bindings.ItemInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ItemInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
		p.Version.Ref(),
		p.Description.Ref(),
		p.IconUrl.Ref(),
		p.Path.Ref(),
		p.OptionsUrl.Ref(),
		p.AppLaunchUrl.Ref(),
		p.HomepageUrl.Ref(),
		p.UpdateUrl.Ref(),
		p.InstallWarnings.Ref(),
		p.ManifestErrors.Ref(),
		p.RuntimeErrors.Ref(),
		p.Views.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.OptionsUrl = p.OptionsUrl.FromRef(js.Undefined)
	p.AppLaunchUrl = p.AppLaunchUrl.FromRef(js.Undefined)
	p.HomepageUrl = p.HomepageUrl.FromRef(js.Undefined)
	p.UpdateUrl = p.UpdateUrl.FromRef(js.Undefined)
	p.InstallWarnings = p.InstallWarnings.FromRef(js.Undefined)
	p.ManifestErrors = p.ManifestErrors.FromRef(js.Undefined)
	p.RuntimeErrors = p.RuntimeErrors.FromRef(js.Undefined)
	p.Views = p.Views.FromRef(js.Undefined)
}

type ItemsInfoCallbackFunc func(this js.Ref, result js.Array[ItemInfo]) js.Ref

func (fn ItemsInfoCallbackFunc) Register() js.Func[func(result js.Array[ItemInfo])] {
	return js.RegisterCallback[func(result js.Array[ItemInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ItemsInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ItemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ItemsInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[ItemInfo]) js.Ref
	Arg T
}

func (cb *ItemsInfoCallback[T]) Register() js.Func[func(result js.Array[ItemInfo])] {
	return js.RegisterCallback[func(result js.Array[ItemInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ItemsInfoCallback[T]) DispatchCallback(
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

		js.Array[ItemInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadError struct {
	// Error is "LoadError.error"
	//
	// Optional
	Error js.String
	// Path is "LoadError.path"
	//
	// Optional
	Path js.String
	// Source is "LoadError.source"
	//
	// Optional
	//
	// NOTE: Source.FFI_USE MUST be set to true to get Source used.
	Source ErrorFileSource
	// RetryGuid is "LoadError.retryGuid"
	//
	// Optional
	RetryGuid js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoadError with all fields set.
func (p LoadError) FromRef(ref js.Ref) LoadError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoadError in the application heap.
func (p LoadError) New() js.Ref {
	return bindings.LoadErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoadError) UpdateFrom(ref js.Ref) {
	bindings.LoadErrorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoadError) Update(ref js.Ref) {
	bindings.LoadErrorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoadError) FreeMembers(recursive bool) {
	js.Free(
		p.Error.Ref(),
		p.Path.Ref(),
		p.RetryGuid.Ref(),
	)
	p.Error = p.Error.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.RetryGuid = p.RetryGuid.FromRef(js.Undefined)
	if recursive {
		p.Source.FreeMembers(true)
	}
}

type LoadErrorCallbackFunc func(this js.Ref, err *LoadError) js.Ref

func (fn LoadErrorCallbackFunc) Register() js.Func[func(err *LoadError)] {
	return js.RegisterCallback[func(err *LoadError)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoadError
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

type LoadErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err *LoadError) js.Ref
	Arg T
}

func (cb *LoadErrorCallback[T]) Register() js.Func[func(err *LoadError)] {
	return js.RegisterCallback[func(err *LoadError)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadErrorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoadError
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

type LoadUnpackedOptions struct {
	// FailQuietly is "LoadUnpackedOptions.failQuietly"
	//
	// Optional
	//
	// NOTE: FFI_USE_FailQuietly MUST be set to true to make this field effective.
	FailQuietly bool
	// PopulateError is "LoadUnpackedOptions.populateError"
	//
	// Optional
	//
	// NOTE: FFI_USE_PopulateError MUST be set to true to make this field effective.
	PopulateError bool
	// RetryGuid is "LoadUnpackedOptions.retryGuid"
	//
	// Optional
	RetryGuid js.String
	// UseDraggedPath is "LoadUnpackedOptions.useDraggedPath"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseDraggedPath MUST be set to true to make this field effective.
	UseDraggedPath bool

	FFI_USE_FailQuietly    bool // for FailQuietly.
	FFI_USE_PopulateError  bool // for PopulateError.
	FFI_USE_UseDraggedPath bool // for UseDraggedPath.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoadUnpackedOptions with all fields set.
func (p LoadUnpackedOptions) FromRef(ref js.Ref) LoadUnpackedOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoadUnpackedOptions in the application heap.
func (p LoadUnpackedOptions) New() js.Ref {
	return bindings.LoadUnpackedOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoadUnpackedOptions) UpdateFrom(ref js.Ref) {
	bindings.LoadUnpackedOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoadUnpackedOptions) Update(ref js.Ref) {
	bindings.LoadUnpackedOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoadUnpackedOptions) FreeMembers(recursive bool) {
	js.Free(
		p.RetryGuid.Ref(),
	)
	p.RetryGuid = p.RetryGuid.FromRef(js.Undefined)
}

type OpenDevToolsProperties struct {
	// ExtensionId is "OpenDevToolsProperties.extensionId"
	//
	// Optional
	ExtensionId js.String
	// RenderViewId is "OpenDevToolsProperties.renderViewId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderViewId MUST be set to true to make this field effective.
	RenderViewId int32
	// RenderProcessId is "OpenDevToolsProperties.renderProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderProcessId MUST be set to true to make this field effective.
	RenderProcessId int32
	// IsServiceWorker is "OpenDevToolsProperties.isServiceWorker"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsServiceWorker MUST be set to true to make this field effective.
	IsServiceWorker bool
	// Incognito is "OpenDevToolsProperties.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// Url is "OpenDevToolsProperties.url"
	//
	// Optional
	Url js.String
	// LineNumber is "OpenDevToolsProperties.lineNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_LineNumber MUST be set to true to make this field effective.
	LineNumber int32
	// ColumnNumber is "OpenDevToolsProperties.columnNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_ColumnNumber MUST be set to true to make this field effective.
	ColumnNumber int32

	FFI_USE_RenderViewId    bool // for RenderViewId.
	FFI_USE_RenderProcessId bool // for RenderProcessId.
	FFI_USE_IsServiceWorker bool // for IsServiceWorker.
	FFI_USE_Incognito       bool // for Incognito.
	FFI_USE_LineNumber      bool // for LineNumber.
	FFI_USE_ColumnNumber    bool // for ColumnNumber.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenDevToolsProperties with all fields set.
func (p OpenDevToolsProperties) FromRef(ref js.Ref) OpenDevToolsProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenDevToolsProperties in the application heap.
func (p OpenDevToolsProperties) New() js.Ref {
	return bindings.OpenDevToolsPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenDevToolsProperties) UpdateFrom(ref js.Ref) {
	bindings.OpenDevToolsPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenDevToolsProperties) Update(ref js.Ref) {
	bindings.OpenDevToolsPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenDevToolsProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.Url.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type PackCallbackFunc func(this js.Ref, response *PackDirectoryResponse) js.Ref

func (fn PackCallbackFunc) Register() js.Func[func(response *PackDirectoryResponse)] {
	return js.RegisterCallback[func(response *PackDirectoryResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PackCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PackDirectoryResponse
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

type PackCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *PackDirectoryResponse) js.Ref
	Arg T
}

func (cb *PackCallback[T]) Register() js.Func[func(response *PackDirectoryResponse)] {
	return js.RegisterCallback[func(response *PackDirectoryResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PackCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PackDirectoryResponse
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

type PackStatus uint32

const (
	_ PackStatus = iota

	PackStatus_SUCCESS
	PackStatus_ERROR
	PackStatus_WARNING
)

func (PackStatus) FromRef(str js.Ref) PackStatus {
	return PackStatus(bindings.ConstOfPackStatus(str))
}

func (x PackStatus) String() (string, bool) {
	switch x {
	case PackStatus_SUCCESS:
		return "SUCCESS", true
	case PackStatus_ERROR:
		return "ERROR", true
	case PackStatus_WARNING:
		return "WARNING", true
	default:
		return "", false
	}
}

type PackDirectoryResponse struct {
	// Message is "PackDirectoryResponse.message"
	//
	// Optional
	Message js.String
	// ItemPath is "PackDirectoryResponse.item_path"
	//
	// Optional
	ItemPath js.String
	// PemPath is "PackDirectoryResponse.pem_path"
	//
	// Optional
	PemPath js.String
	// OverrideFlags is "PackDirectoryResponse.override_flags"
	//
	// Optional
	//
	// NOTE: FFI_USE_OverrideFlags MUST be set to true to make this field effective.
	OverrideFlags int32
	// Status is "PackDirectoryResponse.status"
	//
	// Optional
	Status PackStatus

	FFI_USE_OverrideFlags bool // for OverrideFlags.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PackDirectoryResponse with all fields set.
func (p PackDirectoryResponse) FromRef(ref js.Ref) PackDirectoryResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PackDirectoryResponse in the application heap.
func (p PackDirectoryResponse) New() js.Ref {
	return bindings.PackDirectoryResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PackDirectoryResponse) UpdateFrom(ref js.Ref) {
	bindings.PackDirectoryResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PackDirectoryResponse) Update(ref js.Ref) {
	bindings.PackDirectoryResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PackDirectoryResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
		p.ItemPath.Ref(),
		p.PemPath.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
	p.ItemPath = p.ItemPath.FromRef(js.Undefined)
	p.PemPath = p.PemPath.FromRef(js.Undefined)
}

type ProfileConfigurationUpdate struct {
	// InDeveloperMode is "ProfileConfigurationUpdate.inDeveloperMode"
	//
	// Optional
	//
	// NOTE: FFI_USE_InDeveloperMode MUST be set to true to make this field effective.
	InDeveloperMode bool

	FFI_USE_InDeveloperMode bool // for InDeveloperMode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfileConfigurationUpdate with all fields set.
func (p ProfileConfigurationUpdate) FromRef(ref js.Ref) ProfileConfigurationUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfileConfigurationUpdate in the application heap.
func (p ProfileConfigurationUpdate) New() js.Ref {
	return bindings.ProfileConfigurationUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfileConfigurationUpdate) UpdateFrom(ref js.Ref) {
	bindings.ProfileConfigurationUpdateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfileConfigurationUpdate) Update(ref js.Ref) {
	bindings.ProfileConfigurationUpdateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfileConfigurationUpdate) FreeMembers(recursive bool) {
}

type ProfileInfo struct {
	// CanLoadUnpacked is "ProfileInfo.canLoadUnpacked"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanLoadUnpacked MUST be set to true to make this field effective.
	CanLoadUnpacked bool
	// InDeveloperMode is "ProfileInfo.inDeveloperMode"
	//
	// Optional
	//
	// NOTE: FFI_USE_InDeveloperMode MUST be set to true to make this field effective.
	InDeveloperMode bool
	// IsDeveloperModeControlledByPolicy is "ProfileInfo.isDeveloperModeControlledByPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDeveloperModeControlledByPolicy MUST be set to true to make this field effective.
	IsDeveloperModeControlledByPolicy bool
	// IsIncognitoAvailable is "ProfileInfo.isIncognitoAvailable"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsIncognitoAvailable MUST be set to true to make this field effective.
	IsIncognitoAvailable bool
	// IsChildAccount is "ProfileInfo.isChildAccount"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsChildAccount MUST be set to true to make this field effective.
	IsChildAccount bool

	FFI_USE_CanLoadUnpacked                   bool // for CanLoadUnpacked.
	FFI_USE_InDeveloperMode                   bool // for InDeveloperMode.
	FFI_USE_IsDeveloperModeControlledByPolicy bool // for IsDeveloperModeControlledByPolicy.
	FFI_USE_IsIncognitoAvailable              bool // for IsIncognitoAvailable.
	FFI_USE_IsChildAccount                    bool // for IsChildAccount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfileInfo with all fields set.
func (p ProfileInfo) FromRef(ref js.Ref) ProfileInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfileInfo in the application heap.
func (p ProfileInfo) New() js.Ref {
	return bindings.ProfileInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfileInfo) UpdateFrom(ref js.Ref) {
	bindings.ProfileInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfileInfo) Update(ref js.Ref) {
	bindings.ProfileInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfileInfo) FreeMembers(recursive bool) {
}

type ProfileInfoCallbackFunc func(this js.Ref, info *ProfileInfo) js.Ref

func (fn ProfileInfoCallbackFunc) Register() js.Func[func(info *ProfileInfo)] {
	return js.RegisterCallback[func(info *ProfileInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ProfileInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileInfo
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

type ProfileInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ProfileInfo) js.Ref
	Arg T
}

func (cb *ProfileInfoCallback[T]) Register() js.Func[func(info *ProfileInfo)] {
	return js.RegisterCallback[func(info *ProfileInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ProfileInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileInfo
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

type ReloadOptions struct {
	// FailQuietly is "ReloadOptions.failQuietly"
	//
	// Optional
	//
	// NOTE: FFI_USE_FailQuietly MUST be set to true to make this field effective.
	FailQuietly bool
	// PopulateErrorForUnpacked is "ReloadOptions.populateErrorForUnpacked"
	//
	// Optional
	//
	// NOTE: FFI_USE_PopulateErrorForUnpacked MUST be set to true to make this field effective.
	PopulateErrorForUnpacked bool

	FFI_USE_FailQuietly              bool // for FailQuietly.
	FFI_USE_PopulateErrorForUnpacked bool // for PopulateErrorForUnpacked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReloadOptions with all fields set.
func (p ReloadOptions) FromRef(ref js.Ref) ReloadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReloadOptions in the application heap.
func (p ReloadOptions) New() js.Ref {
	return bindings.ReloadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReloadOptions) UpdateFrom(ref js.Ref) {
	bindings.ReloadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReloadOptions) Update(ref js.Ref) {
	bindings.ReloadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReloadOptions) FreeMembers(recursive bool) {
}

type RequestFileSourceCallbackFunc func(this js.Ref, response *RequestFileSourceResponse) js.Ref

func (fn RequestFileSourceCallbackFunc) Register() js.Func[func(response *RequestFileSourceResponse)] {
	return js.RegisterCallback[func(response *RequestFileSourceResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RequestFileSourceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RequestFileSourceResponse
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

type RequestFileSourceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *RequestFileSourceResponse) js.Ref
	Arg T
}

func (cb *RequestFileSourceCallback[T]) Register() js.Func[func(response *RequestFileSourceResponse)] {
	return js.RegisterCallback[func(response *RequestFileSourceResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RequestFileSourceCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RequestFileSourceResponse
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

type RequestFileSourceResponse struct {
	// Highlight is "RequestFileSourceResponse.highlight"
	//
	// Optional
	Highlight js.String
	// BeforeHighlight is "RequestFileSourceResponse.beforeHighlight"
	//
	// Optional
	BeforeHighlight js.String
	// AfterHighlight is "RequestFileSourceResponse.afterHighlight"
	//
	// Optional
	AfterHighlight js.String
	// Title is "RequestFileSourceResponse.title"
	//
	// Optional
	Title js.String
	// Message is "RequestFileSourceResponse.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestFileSourceResponse with all fields set.
func (p RequestFileSourceResponse) FromRef(ref js.Ref) RequestFileSourceResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestFileSourceResponse in the application heap.
func (p RequestFileSourceResponse) New() js.Ref {
	return bindings.RequestFileSourceResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestFileSourceResponse) UpdateFrom(ref js.Ref) {
	bindings.RequestFileSourceResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestFileSourceResponse) Update(ref js.Ref) {
	bindings.RequestFileSourceResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestFileSourceResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Highlight.Ref(),
		p.BeforeHighlight.Ref(),
		p.AfterHighlight.Ref(),
		p.Title.Ref(),
		p.Message.Ref(),
	)
	p.Highlight = p.Highlight.FromRef(js.Undefined)
	p.BeforeHighlight = p.BeforeHighlight.FromRef(js.Undefined)
	p.AfterHighlight = p.AfterHighlight.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
}

type RequestFileSourceProperties struct {
	// ExtensionId is "RequestFileSourceProperties.extensionId"
	//
	// Optional
	ExtensionId js.String
	// PathSuffix is "RequestFileSourceProperties.pathSuffix"
	//
	// Optional
	PathSuffix js.String
	// Message is "RequestFileSourceProperties.message"
	//
	// Optional
	Message js.String
	// ManifestKey is "RequestFileSourceProperties.manifestKey"
	//
	// Optional
	ManifestKey js.String
	// ManifestSpecific is "RequestFileSourceProperties.manifestSpecific"
	//
	// Optional
	ManifestSpecific js.String
	// LineNumber is "RequestFileSourceProperties.lineNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_LineNumber MUST be set to true to make this field effective.
	LineNumber int32

	FFI_USE_LineNumber bool // for LineNumber.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestFileSourceProperties with all fields set.
func (p RequestFileSourceProperties) FromRef(ref js.Ref) RequestFileSourceProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestFileSourceProperties in the application heap.
func (p RequestFileSourceProperties) New() js.Ref {
	return bindings.RequestFileSourcePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestFileSourceProperties) UpdateFrom(ref js.Ref) {
	bindings.RequestFileSourcePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestFileSourceProperties) Update(ref js.Ref) {
	bindings.RequestFileSourcePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestFileSourceProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionId.Ref(),
		p.PathSuffix.Ref(),
		p.Message.Ref(),
		p.ManifestKey.Ref(),
		p.ManifestSpecific.Ref(),
	)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.PathSuffix = p.PathSuffix.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
	p.ManifestKey = p.ManifestKey.FromRef(js.Undefined)
	p.ManifestSpecific = p.ManifestSpecific.FromRef(js.Undefined)
}

type SelectType uint32

const (
	_ SelectType = iota

	SelectType_FILE
	SelectType_FOLDER
)

func (SelectType) FromRef(str js.Ref) SelectType {
	return SelectType(bindings.ConstOfSelectType(str))
}

func (x SelectType) String() (string, bool) {
	switch x {
	case SelectType_FILE:
		return "FILE", true
	case SelectType_FOLDER:
		return "FOLDER", true
	default:
		return "", false
	}
}

type SiteSet uint32

const (
	_ SiteSet = iota

	SiteSet_USER_PERMITTED
	SiteSet_USER_RESTRICTED
	SiteSet_EXTENSION_SPECIFIED
)

func (SiteSet) FromRef(str js.Ref) SiteSet {
	return SiteSet(bindings.ConstOfSiteSet(str))
}

func (x SiteSet) String() (string, bool) {
	switch x {
	case SiteSet_USER_PERMITTED:
		return "USER_PERMITTED", true
	case SiteSet_USER_RESTRICTED:
		return "USER_RESTRICTED", true
	case SiteSet_EXTENSION_SPECIFIED:
		return "EXTENSION_SPECIFIED", true
	default:
		return "", false
	}
}

type SiteInfo struct {
	// SiteSet is "SiteInfo.siteSet"
	//
	// Optional
	SiteSet SiteSet
	// NumExtensions is "SiteInfo.numExtensions"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumExtensions MUST be set to true to make this field effective.
	NumExtensions int32
	// Site is "SiteInfo.site"
	//
	// Optional
	Site js.String

	FFI_USE_NumExtensions bool // for NumExtensions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SiteInfo with all fields set.
func (p SiteInfo) FromRef(ref js.Ref) SiteInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SiteInfo in the application heap.
func (p SiteInfo) New() js.Ref {
	return bindings.SiteInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SiteInfo) UpdateFrom(ref js.Ref) {
	bindings.SiteInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SiteInfo) Update(ref js.Ref) {
	bindings.SiteInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SiteInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Site.Ref(),
	)
	p.Site = p.Site.FromRef(js.Undefined)
}

type SiteGroup struct {
	// EtldPlusOne is "SiteGroup.etldPlusOne"
	//
	// Optional
	EtldPlusOne js.String
	// NumExtensions is "SiteGroup.numExtensions"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumExtensions MUST be set to true to make this field effective.
	NumExtensions int32
	// Sites is "SiteGroup.sites"
	//
	// Optional
	Sites js.Array[SiteInfo]

	FFI_USE_NumExtensions bool // for NumExtensions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SiteGroup with all fields set.
func (p SiteGroup) FromRef(ref js.Ref) SiteGroup {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SiteGroup in the application heap.
func (p SiteGroup) New() js.Ref {
	return bindings.SiteGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SiteGroup) UpdateFrom(ref js.Ref) {
	bindings.SiteGroupJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SiteGroup) Update(ref js.Ref) {
	bindings.SiteGroupJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SiteGroup) FreeMembers(recursive bool) {
	js.Free(
		p.EtldPlusOne.Ref(),
		p.Sites.Ref(),
	)
	p.EtldPlusOne = p.EtldPlusOne.FromRef(js.Undefined)
	p.Sites = p.Sites.FromRef(js.Undefined)
}

type StringCallbackFunc func(this js.Ref, string js.String) js.Ref

func (fn StringCallbackFunc) Register() js.Func[func(string js.String)] {
	return js.RegisterCallback[func(string js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StringCallbackFunc) DispatchCallback(
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

type StringCallback[T any] struct {
	Fn  func(arg T, this js.Ref, string js.String) js.Ref
	Arg T
}

func (cb *StringCallback[T]) Register() js.Func[func(string js.String)] {
	return js.RegisterCallback[func(string js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StringCallback[T]) DispatchCallback(
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

type UserAndExtensionSitesByEtldCallbackFunc func(this js.Ref, siteGroups js.Array[SiteGroup]) js.Ref

func (fn UserAndExtensionSitesByEtldCallbackFunc) Register() js.Func[func(siteGroups js.Array[SiteGroup])] {
	return js.RegisterCallback[func(siteGroups js.Array[SiteGroup])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UserAndExtensionSitesByEtldCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SiteGroup]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UserAndExtensionSitesByEtldCallback[T any] struct {
	Fn  func(arg T, this js.Ref, siteGroups js.Array[SiteGroup]) js.Ref
	Arg T
}

func (cb *UserAndExtensionSitesByEtldCallback[T]) Register() js.Func[func(siteGroups js.Array[SiteGroup])] {
	return js.RegisterCallback[func(siteGroups js.Array[SiteGroup])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UserAndExtensionSitesByEtldCallback[T]) DispatchCallback(
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

		js.Array[SiteGroup]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UserSiteSettings struct {
	// PermittedSites is "UserSiteSettings.permittedSites"
	//
	// Optional
	PermittedSites js.Array[js.String]
	// RestrictedSites is "UserSiteSettings.restrictedSites"
	//
	// Optional
	RestrictedSites js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UserSiteSettings with all fields set.
func (p UserSiteSettings) FromRef(ref js.Ref) UserSiteSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UserSiteSettings in the application heap.
func (p UserSiteSettings) New() js.Ref {
	return bindings.UserSiteSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UserSiteSettings) UpdateFrom(ref js.Ref) {
	bindings.UserSiteSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UserSiteSettings) Update(ref js.Ref) {
	bindings.UserSiteSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UserSiteSettings) FreeMembers(recursive bool) {
	js.Free(
		p.PermittedSites.Ref(),
		p.RestrictedSites.Ref(),
	)
	p.PermittedSites = p.PermittedSites.FromRef(js.Undefined)
	p.RestrictedSites = p.RestrictedSites.FromRef(js.Undefined)
}

type UserSiteSettingsCallbackFunc func(this js.Ref, settings *UserSiteSettings) js.Ref

func (fn UserSiteSettingsCallbackFunc) Register() js.Func[func(settings *UserSiteSettings)] {
	return js.RegisterCallback[func(settings *UserSiteSettings)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UserSiteSettingsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UserSiteSettings
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

type UserSiteSettingsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, settings *UserSiteSettings) js.Ref
	Arg T
}

func (cb *UserSiteSettingsCallback[T]) Register() js.Func[func(settings *UserSiteSettings)] {
	return js.RegisterCallback[func(settings *UserSiteSettings)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UserSiteSettingsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UserSiteSettings
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

type UserSiteSettingsOptions struct {
	// SiteSet is "UserSiteSettingsOptions.siteSet"
	//
	// Optional
	SiteSet SiteSet
	// Hosts is "UserSiteSettingsOptions.hosts"
	//
	// Optional
	Hosts js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UserSiteSettingsOptions with all fields set.
func (p UserSiteSettingsOptions) FromRef(ref js.Ref) UserSiteSettingsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UserSiteSettingsOptions in the application heap.
func (p UserSiteSettingsOptions) New() js.Ref {
	return bindings.UserSiteSettingsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UserSiteSettingsOptions) UpdateFrom(ref js.Ref) {
	bindings.UserSiteSettingsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UserSiteSettingsOptions) Update(ref js.Ref) {
	bindings.UserSiteSettingsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UserSiteSettingsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Hosts.Ref(),
	)
	p.Hosts = p.Hosts.FromRef(js.Undefined)
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

// HasFuncAddHostPermission returns true if the function "WEBEXT.developerPrivate.addHostPermission" exists.
func HasFuncAddHostPermission() bool {
	return js.True == bindings.HasFuncAddHostPermission()
}

// FuncAddHostPermission returns the function "WEBEXT.developerPrivate.addHostPermission".
func FuncAddHostPermission() (fn js.Func[func(extensionId js.String, host js.String) js.Promise[js.Void]]) {
	bindings.FuncAddHostPermission(
		js.Pointer(&fn),
	)
	return
}

// AddHostPermission calls the function "WEBEXT.developerPrivate.addHostPermission" directly.
func AddHostPermission(extensionId js.String, host js.String) (ret js.Promise[js.Void]) {
	bindings.CallAddHostPermission(
		js.Pointer(&ret),
		extensionId.Ref(),
		host.Ref(),
	)

	return
}

// TryAddHostPermission calls the function "WEBEXT.developerPrivate.addHostPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddHostPermission(extensionId js.String, host js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddHostPermission(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		host.Ref(),
	)

	return
}

// HasFuncAddUserSpecifiedSites returns true if the function "WEBEXT.developerPrivate.addUserSpecifiedSites" exists.
func HasFuncAddUserSpecifiedSites() bool {
	return js.True == bindings.HasFuncAddUserSpecifiedSites()
}

// FuncAddUserSpecifiedSites returns the function "WEBEXT.developerPrivate.addUserSpecifiedSites".
func FuncAddUserSpecifiedSites() (fn js.Func[func(options UserSiteSettingsOptions) js.Promise[js.Void]]) {
	bindings.FuncAddUserSpecifiedSites(
		js.Pointer(&fn),
	)
	return
}

// AddUserSpecifiedSites calls the function "WEBEXT.developerPrivate.addUserSpecifiedSites" directly.
func AddUserSpecifiedSites(options UserSiteSettingsOptions) (ret js.Promise[js.Void]) {
	bindings.CallAddUserSpecifiedSites(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAddUserSpecifiedSites calls the function "WEBEXT.developerPrivate.addUserSpecifiedSites"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddUserSpecifiedSites(options UserSiteSettingsOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddUserSpecifiedSites(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncAutoUpdate returns true if the function "WEBEXT.developerPrivate.autoUpdate" exists.
func HasFuncAutoUpdate() bool {
	return js.True == bindings.HasFuncAutoUpdate()
}

// FuncAutoUpdate returns the function "WEBEXT.developerPrivate.autoUpdate".
func FuncAutoUpdate() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncAutoUpdate(
		js.Pointer(&fn),
	)
	return
}

// AutoUpdate calls the function "WEBEXT.developerPrivate.autoUpdate" directly.
func AutoUpdate() (ret js.Promise[js.Void]) {
	bindings.CallAutoUpdate(
		js.Pointer(&ret),
	)

	return
}

// TryAutoUpdate calls the function "WEBEXT.developerPrivate.autoUpdate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAutoUpdate() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutoUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncChoosePath returns true if the function "WEBEXT.developerPrivate.choosePath" exists.
func HasFuncChoosePath() bool {
	return js.True == bindings.HasFuncChoosePath()
}

// FuncChoosePath returns the function "WEBEXT.developerPrivate.choosePath".
func FuncChoosePath() (fn js.Func[func(selectType SelectType, fileType FileType) js.Promise[js.String]]) {
	bindings.FuncChoosePath(
		js.Pointer(&fn),
	)
	return
}

// ChoosePath calls the function "WEBEXT.developerPrivate.choosePath" directly.
func ChoosePath(selectType SelectType, fileType FileType) (ret js.Promise[js.String]) {
	bindings.CallChoosePath(
		js.Pointer(&ret),
		uint32(selectType),
		uint32(fileType),
	)

	return
}

// TryChoosePath calls the function "WEBEXT.developerPrivate.choosePath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChoosePath(selectType SelectType, fileType FileType) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChoosePath(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(selectType),
		uint32(fileType),
	)

	return
}

// HasFuncDeleteExtensionErrors returns true if the function "WEBEXT.developerPrivate.deleteExtensionErrors" exists.
func HasFuncDeleteExtensionErrors() bool {
	return js.True == bindings.HasFuncDeleteExtensionErrors()
}

// FuncDeleteExtensionErrors returns the function "WEBEXT.developerPrivate.deleteExtensionErrors".
func FuncDeleteExtensionErrors() (fn js.Func[func(properties DeleteExtensionErrorsProperties) js.Promise[js.Void]]) {
	bindings.FuncDeleteExtensionErrors(
		js.Pointer(&fn),
	)
	return
}

// DeleteExtensionErrors calls the function "WEBEXT.developerPrivate.deleteExtensionErrors" directly.
func DeleteExtensionErrors(properties DeleteExtensionErrorsProperties) (ret js.Promise[js.Void]) {
	bindings.CallDeleteExtensionErrors(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryDeleteExtensionErrors calls the function "WEBEXT.developerPrivate.deleteExtensionErrors"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteExtensionErrors(properties DeleteExtensionErrorsProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteExtensionErrors(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncGetExtensionInfo returns true if the function "WEBEXT.developerPrivate.getExtensionInfo" exists.
func HasFuncGetExtensionInfo() bool {
	return js.True == bindings.HasFuncGetExtensionInfo()
}

// FuncGetExtensionInfo returns the function "WEBEXT.developerPrivate.getExtensionInfo".
func FuncGetExtensionInfo() (fn js.Func[func(id js.String) js.Promise[ExtensionInfo]]) {
	bindings.FuncGetExtensionInfo(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionInfo calls the function "WEBEXT.developerPrivate.getExtensionInfo" directly.
func GetExtensionInfo(id js.String) (ret js.Promise[ExtensionInfo]) {
	bindings.CallGetExtensionInfo(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetExtensionInfo calls the function "WEBEXT.developerPrivate.getExtensionInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionInfo(id js.String) (ret js.Promise[ExtensionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetExtensionSize returns true if the function "WEBEXT.developerPrivate.getExtensionSize" exists.
func HasFuncGetExtensionSize() bool {
	return js.True == bindings.HasFuncGetExtensionSize()
}

// FuncGetExtensionSize returns the function "WEBEXT.developerPrivate.getExtensionSize".
func FuncGetExtensionSize() (fn js.Func[func(id js.String) js.Promise[js.String]]) {
	bindings.FuncGetExtensionSize(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionSize calls the function "WEBEXT.developerPrivate.getExtensionSize" directly.
func GetExtensionSize(id js.String) (ret js.Promise[js.String]) {
	bindings.CallGetExtensionSize(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetExtensionSize calls the function "WEBEXT.developerPrivate.getExtensionSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionSize(id js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionSize(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetExtensionsInfo returns true if the function "WEBEXT.developerPrivate.getExtensionsInfo" exists.
func HasFuncGetExtensionsInfo() bool {
	return js.True == bindings.HasFuncGetExtensionsInfo()
}

// FuncGetExtensionsInfo returns the function "WEBEXT.developerPrivate.getExtensionsInfo".
func FuncGetExtensionsInfo() (fn js.Func[func(options GetExtensionsInfoOptions) js.Promise[js.Array[ExtensionInfo]]]) {
	bindings.FuncGetExtensionsInfo(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionsInfo calls the function "WEBEXT.developerPrivate.getExtensionsInfo" directly.
func GetExtensionsInfo(options GetExtensionsInfoOptions) (ret js.Promise[js.Array[ExtensionInfo]]) {
	bindings.CallGetExtensionsInfo(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetExtensionsInfo calls the function "WEBEXT.developerPrivate.getExtensionsInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionsInfo(options GetExtensionsInfoOptions) (ret js.Promise[js.Array[ExtensionInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionsInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetMatchingExtensionsForSite returns true if the function "WEBEXT.developerPrivate.getMatchingExtensionsForSite" exists.
func HasFuncGetMatchingExtensionsForSite() bool {
	return js.True == bindings.HasFuncGetMatchingExtensionsForSite()
}

// FuncGetMatchingExtensionsForSite returns the function "WEBEXT.developerPrivate.getMatchingExtensionsForSite".
func FuncGetMatchingExtensionsForSite() (fn js.Func[func(site js.String) js.Promise[js.Array[MatchingExtensionInfo]]]) {
	bindings.FuncGetMatchingExtensionsForSite(
		js.Pointer(&fn),
	)
	return
}

// GetMatchingExtensionsForSite calls the function "WEBEXT.developerPrivate.getMatchingExtensionsForSite" directly.
func GetMatchingExtensionsForSite(site js.String) (ret js.Promise[js.Array[MatchingExtensionInfo]]) {
	bindings.CallGetMatchingExtensionsForSite(
		js.Pointer(&ret),
		site.Ref(),
	)

	return
}

// TryGetMatchingExtensionsForSite calls the function "WEBEXT.developerPrivate.getMatchingExtensionsForSite"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMatchingExtensionsForSite(site js.String) (ret js.Promise[js.Array[MatchingExtensionInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMatchingExtensionsForSite(
		js.Pointer(&ret), js.Pointer(&exception),
		site.Ref(),
	)

	return
}

// HasFuncGetProfileConfiguration returns true if the function "WEBEXT.developerPrivate.getProfileConfiguration" exists.
func HasFuncGetProfileConfiguration() bool {
	return js.True == bindings.HasFuncGetProfileConfiguration()
}

// FuncGetProfileConfiguration returns the function "WEBEXT.developerPrivate.getProfileConfiguration".
func FuncGetProfileConfiguration() (fn js.Func[func() js.Promise[ProfileInfo]]) {
	bindings.FuncGetProfileConfiguration(
		js.Pointer(&fn),
	)
	return
}

// GetProfileConfiguration calls the function "WEBEXT.developerPrivate.getProfileConfiguration" directly.
func GetProfileConfiguration() (ret js.Promise[ProfileInfo]) {
	bindings.CallGetProfileConfiguration(
		js.Pointer(&ret),
	)

	return
}

// TryGetProfileConfiguration calls the function "WEBEXT.developerPrivate.getProfileConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProfileConfiguration() (ret js.Promise[ProfileInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProfileConfiguration(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUserAndExtensionSitesByEtld returns true if the function "WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld" exists.
func HasFuncGetUserAndExtensionSitesByEtld() bool {
	return js.True == bindings.HasFuncGetUserAndExtensionSitesByEtld()
}

// FuncGetUserAndExtensionSitesByEtld returns the function "WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld".
func FuncGetUserAndExtensionSitesByEtld() (fn js.Func[func() js.Promise[js.Array[SiteGroup]]]) {
	bindings.FuncGetUserAndExtensionSitesByEtld(
		js.Pointer(&fn),
	)
	return
}

// GetUserAndExtensionSitesByEtld calls the function "WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld" directly.
func GetUserAndExtensionSitesByEtld() (ret js.Promise[js.Array[SiteGroup]]) {
	bindings.CallGetUserAndExtensionSitesByEtld(
		js.Pointer(&ret),
	)

	return
}

// TryGetUserAndExtensionSitesByEtld calls the function "WEBEXT.developerPrivate.getUserAndExtensionSitesByEtld"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserAndExtensionSitesByEtld() (ret js.Promise[js.Array[SiteGroup]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserAndExtensionSitesByEtld(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUserSiteSettings returns true if the function "WEBEXT.developerPrivate.getUserSiteSettings" exists.
func HasFuncGetUserSiteSettings() bool {
	return js.True == bindings.HasFuncGetUserSiteSettings()
}

// FuncGetUserSiteSettings returns the function "WEBEXT.developerPrivate.getUserSiteSettings".
func FuncGetUserSiteSettings() (fn js.Func[func() js.Promise[UserSiteSettings]]) {
	bindings.FuncGetUserSiteSettings(
		js.Pointer(&fn),
	)
	return
}

// GetUserSiteSettings calls the function "WEBEXT.developerPrivate.getUserSiteSettings" directly.
func GetUserSiteSettings() (ret js.Promise[UserSiteSettings]) {
	bindings.CallGetUserSiteSettings(
		js.Pointer(&ret),
	)

	return
}

// TryGetUserSiteSettings calls the function "WEBEXT.developerPrivate.getUserSiteSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserSiteSettings() (ret js.Promise[UserSiteSettings], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserSiteSettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInspect returns true if the function "WEBEXT.developerPrivate.inspect" exists.
func HasFuncInspect() bool {
	return js.True == bindings.HasFuncInspect()
}

// FuncInspect returns the function "WEBEXT.developerPrivate.inspect".
func FuncInspect() (fn js.Func[func(options InspectOptions, callback js.Func[func()])]) {
	bindings.FuncInspect(
		js.Pointer(&fn),
	)
	return
}

// Inspect calls the function "WEBEXT.developerPrivate.inspect" directly.
func Inspect(options InspectOptions, callback js.Func[func()]) (ret js.Void) {
	bindings.CallInspect(
		js.Pointer(&ret),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryInspect calls the function "WEBEXT.developerPrivate.inspect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInspect(options InspectOptions, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInspect(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncInstallDroppedFile returns true if the function "WEBEXT.developerPrivate.installDroppedFile" exists.
func HasFuncInstallDroppedFile() bool {
	return js.True == bindings.HasFuncInstallDroppedFile()
}

// FuncInstallDroppedFile returns the function "WEBEXT.developerPrivate.installDroppedFile".
func FuncInstallDroppedFile() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncInstallDroppedFile(
		js.Pointer(&fn),
	)
	return
}

// InstallDroppedFile calls the function "WEBEXT.developerPrivate.installDroppedFile" directly.
func InstallDroppedFile() (ret js.Promise[js.Void]) {
	bindings.CallInstallDroppedFile(
		js.Pointer(&ret),
	)

	return
}

// TryInstallDroppedFile calls the function "WEBEXT.developerPrivate.installDroppedFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallDroppedFile() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallDroppedFile(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsProfileManaged returns true if the function "WEBEXT.developerPrivate.isProfileManaged" exists.
func HasFuncIsProfileManaged() bool {
	return js.True == bindings.HasFuncIsProfileManaged()
}

// FuncIsProfileManaged returns the function "WEBEXT.developerPrivate.isProfileManaged".
func FuncIsProfileManaged() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsProfileManaged(
		js.Pointer(&fn),
	)
	return
}

// IsProfileManaged calls the function "WEBEXT.developerPrivate.isProfileManaged" directly.
func IsProfileManaged() (ret js.Promise[js.Boolean]) {
	bindings.CallIsProfileManaged(
		js.Pointer(&ret),
	)

	return
}

// TryIsProfileManaged calls the function "WEBEXT.developerPrivate.isProfileManaged"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsProfileManaged() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsProfileManaged(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLoadDirectory returns true if the function "WEBEXT.developerPrivate.loadDirectory" exists.
func HasFuncLoadDirectory() bool {
	return js.True == bindings.HasFuncLoadDirectory()
}

// FuncLoadDirectory returns the function "WEBEXT.developerPrivate.loadDirectory".
func FuncLoadDirectory() (fn js.Func[func(directory js.Object) js.Promise[js.String]]) {
	bindings.FuncLoadDirectory(
		js.Pointer(&fn),
	)
	return
}

// LoadDirectory calls the function "WEBEXT.developerPrivate.loadDirectory" directly.
func LoadDirectory(directory js.Object) (ret js.Promise[js.String]) {
	bindings.CallLoadDirectory(
		js.Pointer(&ret),
		directory.Ref(),
	)

	return
}

// TryLoadDirectory calls the function "WEBEXT.developerPrivate.loadDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadDirectory(directory js.Object) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadDirectory(
		js.Pointer(&ret), js.Pointer(&exception),
		directory.Ref(),
	)

	return
}

// HasFuncLoadUnpacked returns true if the function "WEBEXT.developerPrivate.loadUnpacked" exists.
func HasFuncLoadUnpacked() bool {
	return js.True == bindings.HasFuncLoadUnpacked()
}

// FuncLoadUnpacked returns the function "WEBEXT.developerPrivate.loadUnpacked".
func FuncLoadUnpacked() (fn js.Func[func(options LoadUnpackedOptions) js.Promise[LoadError]]) {
	bindings.FuncLoadUnpacked(
		js.Pointer(&fn),
	)
	return
}

// LoadUnpacked calls the function "WEBEXT.developerPrivate.loadUnpacked" directly.
func LoadUnpacked(options LoadUnpackedOptions) (ret js.Promise[LoadError]) {
	bindings.CallLoadUnpacked(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryLoadUnpacked calls the function "WEBEXT.developerPrivate.loadUnpacked"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadUnpacked(options LoadUnpackedOptions) (ret js.Promise[LoadError], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadUnpacked(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncNotifyDragInstallInProgress returns true if the function "WEBEXT.developerPrivate.notifyDragInstallInProgress" exists.
func HasFuncNotifyDragInstallInProgress() bool {
	return js.True == bindings.HasFuncNotifyDragInstallInProgress()
}

// FuncNotifyDragInstallInProgress returns the function "WEBEXT.developerPrivate.notifyDragInstallInProgress".
func FuncNotifyDragInstallInProgress() (fn js.Func[func()]) {
	bindings.FuncNotifyDragInstallInProgress(
		js.Pointer(&fn),
	)
	return
}

// NotifyDragInstallInProgress calls the function "WEBEXT.developerPrivate.notifyDragInstallInProgress" directly.
func NotifyDragInstallInProgress() (ret js.Void) {
	bindings.CallNotifyDragInstallInProgress(
		js.Pointer(&ret),
	)

	return
}

// TryNotifyDragInstallInProgress calls the function "WEBEXT.developerPrivate.notifyDragInstallInProgress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyDragInstallInProgress() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyDragInstallInProgress(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnItemStateChangedEventCallbackFunc func(this js.Ref, response *EventData) js.Ref

func (fn OnItemStateChangedEventCallbackFunc) Register() js.Func[func(response *EventData)] {
	return js.RegisterCallback[func(response *EventData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnItemStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EventData
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

type OnItemStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *EventData) js.Ref
	Arg T
}

func (cb *OnItemStateChangedEventCallback[T]) Register() js.Func[func(response *EventData)] {
	return js.RegisterCallback[func(response *EventData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnItemStateChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 EventData
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

// HasFuncOnItemStateChanged returns true if the function "WEBEXT.developerPrivate.onItemStateChanged.addListener" exists.
func HasFuncOnItemStateChanged() bool {
	return js.True == bindings.HasFuncOnItemStateChanged()
}

// FuncOnItemStateChanged returns the function "WEBEXT.developerPrivate.onItemStateChanged.addListener".
func FuncOnItemStateChanged() (fn js.Func[func(callback js.Func[func(response *EventData)])]) {
	bindings.FuncOnItemStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.addListener" directly.
func OnItemStateChanged(callback js.Func[func(response *EventData)]) (ret js.Void) {
	bindings.CallOnItemStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnItemStateChanged(callback js.Func[func(response *EventData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnItemStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffItemStateChanged returns true if the function "WEBEXT.developerPrivate.onItemStateChanged.removeListener" exists.
func HasFuncOffItemStateChanged() bool {
	return js.True == bindings.HasFuncOffItemStateChanged()
}

// FuncOffItemStateChanged returns the function "WEBEXT.developerPrivate.onItemStateChanged.removeListener".
func FuncOffItemStateChanged() (fn js.Func[func(callback js.Func[func(response *EventData)])]) {
	bindings.FuncOffItemStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.removeListener" directly.
func OffItemStateChanged(callback js.Func[func(response *EventData)]) (ret js.Void) {
	bindings.CallOffItemStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffItemStateChanged(callback js.Func[func(response *EventData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffItemStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnItemStateChanged returns true if the function "WEBEXT.developerPrivate.onItemStateChanged.hasListener" exists.
func HasFuncHasOnItemStateChanged() bool {
	return js.True == bindings.HasFuncHasOnItemStateChanged()
}

// FuncHasOnItemStateChanged returns the function "WEBEXT.developerPrivate.onItemStateChanged.hasListener".
func FuncHasOnItemStateChanged() (fn js.Func[func(callback js.Func[func(response *EventData)]) bool]) {
	bindings.FuncHasOnItemStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.hasListener" directly.
func HasOnItemStateChanged(callback js.Func[func(response *EventData)]) (ret bool) {
	bindings.CallHasOnItemStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnItemStateChanged calls the function "WEBEXT.developerPrivate.onItemStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnItemStateChanged(callback js.Func[func(response *EventData)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnItemStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnProfileStateChangedEventCallbackFunc func(this js.Ref, info *ProfileInfo) js.Ref

func (fn OnProfileStateChangedEventCallbackFunc) Register() js.Func[func(info *ProfileInfo)] {
	return js.RegisterCallback[func(info *ProfileInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnProfileStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileInfo
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

type OnProfileStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ProfileInfo) js.Ref
	Arg T
}

func (cb *OnProfileStateChangedEventCallback[T]) Register() js.Func[func(info *ProfileInfo)] {
	return js.RegisterCallback[func(info *ProfileInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnProfileStateChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProfileInfo
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

// HasFuncOnProfileStateChanged returns true if the function "WEBEXT.developerPrivate.onProfileStateChanged.addListener" exists.
func HasFuncOnProfileStateChanged() bool {
	return js.True == bindings.HasFuncOnProfileStateChanged()
}

// FuncOnProfileStateChanged returns the function "WEBEXT.developerPrivate.onProfileStateChanged.addListener".
func FuncOnProfileStateChanged() (fn js.Func[func(callback js.Func[func(info *ProfileInfo)])]) {
	bindings.FuncOnProfileStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.addListener" directly.
func OnProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret js.Void) {
	bindings.CallOnProfileStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnProfileStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffProfileStateChanged returns true if the function "WEBEXT.developerPrivate.onProfileStateChanged.removeListener" exists.
func HasFuncOffProfileStateChanged() bool {
	return js.True == bindings.HasFuncOffProfileStateChanged()
}

// FuncOffProfileStateChanged returns the function "WEBEXT.developerPrivate.onProfileStateChanged.removeListener".
func FuncOffProfileStateChanged() (fn js.Func[func(callback js.Func[func(info *ProfileInfo)])]) {
	bindings.FuncOffProfileStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.removeListener" directly.
func OffProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret js.Void) {
	bindings.CallOffProfileStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffProfileStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnProfileStateChanged returns true if the function "WEBEXT.developerPrivate.onProfileStateChanged.hasListener" exists.
func HasFuncHasOnProfileStateChanged() bool {
	return js.True == bindings.HasFuncHasOnProfileStateChanged()
}

// FuncHasOnProfileStateChanged returns the function "WEBEXT.developerPrivate.onProfileStateChanged.hasListener".
func FuncHasOnProfileStateChanged() (fn js.Func[func(callback js.Func[func(info *ProfileInfo)]) bool]) {
	bindings.FuncHasOnProfileStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.hasListener" directly.
func HasOnProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret bool) {
	bindings.CallHasOnProfileStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnProfileStateChanged calls the function "WEBEXT.developerPrivate.onProfileStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnProfileStateChanged(callback js.Func[func(info *ProfileInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnProfileStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUserSiteSettingsChangedEventCallbackFunc func(this js.Ref, settings *UserSiteSettings) js.Ref

func (fn OnUserSiteSettingsChangedEventCallbackFunc) Register() js.Func[func(settings *UserSiteSettings)] {
	return js.RegisterCallback[func(settings *UserSiteSettings)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUserSiteSettingsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UserSiteSettings
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

type OnUserSiteSettingsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, settings *UserSiteSettings) js.Ref
	Arg T
}

func (cb *OnUserSiteSettingsChangedEventCallback[T]) Register() js.Func[func(settings *UserSiteSettings)] {
	return js.RegisterCallback[func(settings *UserSiteSettings)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUserSiteSettingsChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UserSiteSettings
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

// HasFuncOnUserSiteSettingsChanged returns true if the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener" exists.
func HasFuncOnUserSiteSettingsChanged() bool {
	return js.True == bindings.HasFuncOnUserSiteSettingsChanged()
}

// FuncOnUserSiteSettingsChanged returns the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener".
func FuncOnUserSiteSettingsChanged() (fn js.Func[func(callback js.Func[func(settings *UserSiteSettings)])]) {
	bindings.FuncOnUserSiteSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener" directly.
func OnUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret js.Void) {
	bindings.CallOnUserSiteSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUserSiteSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUserSiteSettingsChanged returns true if the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener" exists.
func HasFuncOffUserSiteSettingsChanged() bool {
	return js.True == bindings.HasFuncOffUserSiteSettingsChanged()
}

// FuncOffUserSiteSettingsChanged returns the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener".
func FuncOffUserSiteSettingsChanged() (fn js.Func[func(callback js.Func[func(settings *UserSiteSettings)])]) {
	bindings.FuncOffUserSiteSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener" directly.
func OffUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret js.Void) {
	bindings.CallOffUserSiteSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUserSiteSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUserSiteSettingsChanged returns true if the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener" exists.
func HasFuncHasOnUserSiteSettingsChanged() bool {
	return js.True == bindings.HasFuncHasOnUserSiteSettingsChanged()
}

// FuncHasOnUserSiteSettingsChanged returns the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener".
func FuncHasOnUserSiteSettingsChanged() (fn js.Func[func(callback js.Func[func(settings *UserSiteSettings)]) bool]) {
	bindings.FuncHasOnUserSiteSettingsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener" directly.
func HasOnUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret bool) {
	bindings.CallHasOnUserSiteSettingsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUserSiteSettingsChanged calls the function "WEBEXT.developerPrivate.onUserSiteSettingsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUserSiteSettingsChanged(callback js.Func[func(settings *UserSiteSettings)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUserSiteSettingsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenDevTools returns true if the function "WEBEXT.developerPrivate.openDevTools" exists.
func HasFuncOpenDevTools() bool {
	return js.True == bindings.HasFuncOpenDevTools()
}

// FuncOpenDevTools returns the function "WEBEXT.developerPrivate.openDevTools".
func FuncOpenDevTools() (fn js.Func[func(properties OpenDevToolsProperties) js.Promise[js.Void]]) {
	bindings.FuncOpenDevTools(
		js.Pointer(&fn),
	)
	return
}

// OpenDevTools calls the function "WEBEXT.developerPrivate.openDevTools" directly.
func OpenDevTools(properties OpenDevToolsProperties) (ret js.Promise[js.Void]) {
	bindings.CallOpenDevTools(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryOpenDevTools calls the function "WEBEXT.developerPrivate.openDevTools"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenDevTools(properties OpenDevToolsProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenDevTools(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncPackDirectory returns true if the function "WEBEXT.developerPrivate.packDirectory" exists.
func HasFuncPackDirectory() bool {
	return js.True == bindings.HasFuncPackDirectory()
}

// FuncPackDirectory returns the function "WEBEXT.developerPrivate.packDirectory".
func FuncPackDirectory() (fn js.Func[func(path js.String, privateKeyPath js.String, flags int32) js.Promise[PackDirectoryResponse]]) {
	bindings.FuncPackDirectory(
		js.Pointer(&fn),
	)
	return
}

// PackDirectory calls the function "WEBEXT.developerPrivate.packDirectory" directly.
func PackDirectory(path js.String, privateKeyPath js.String, flags int32) (ret js.Promise[PackDirectoryResponse]) {
	bindings.CallPackDirectory(
		js.Pointer(&ret),
		path.Ref(),
		privateKeyPath.Ref(),
		int32(flags),
	)

	return
}

// TryPackDirectory calls the function "WEBEXT.developerPrivate.packDirectory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPackDirectory(path js.String, privateKeyPath js.String, flags int32) (ret js.Promise[PackDirectoryResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPackDirectory(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		privateKeyPath.Ref(),
		int32(flags),
	)

	return
}

// HasFuncReload returns true if the function "WEBEXT.developerPrivate.reload" exists.
func HasFuncReload() bool {
	return js.True == bindings.HasFuncReload()
}

// FuncReload returns the function "WEBEXT.developerPrivate.reload".
func FuncReload() (fn js.Func[func(extensionId js.String, options ReloadOptions) js.Promise[LoadError]]) {
	bindings.FuncReload(
		js.Pointer(&fn),
	)
	return
}

// Reload calls the function "WEBEXT.developerPrivate.reload" directly.
func Reload(extensionId js.String, options ReloadOptions) (ret js.Promise[LoadError]) {
	bindings.CallReload(
		js.Pointer(&ret),
		extensionId.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReload calls the function "WEBEXT.developerPrivate.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReload(extensionId js.String, options ReloadOptions) (ret js.Promise[LoadError], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReload(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveHostPermission returns true if the function "WEBEXT.developerPrivate.removeHostPermission" exists.
func HasFuncRemoveHostPermission() bool {
	return js.True == bindings.HasFuncRemoveHostPermission()
}

// FuncRemoveHostPermission returns the function "WEBEXT.developerPrivate.removeHostPermission".
func FuncRemoveHostPermission() (fn js.Func[func(extensionId js.String, host js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveHostPermission(
		js.Pointer(&fn),
	)
	return
}

// RemoveHostPermission calls the function "WEBEXT.developerPrivate.removeHostPermission" directly.
func RemoveHostPermission(extensionId js.String, host js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveHostPermission(
		js.Pointer(&ret),
		extensionId.Ref(),
		host.Ref(),
	)

	return
}

// TryRemoveHostPermission calls the function "WEBEXT.developerPrivate.removeHostPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveHostPermission(extensionId js.String, host js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveHostPermission(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		host.Ref(),
	)

	return
}

// HasFuncRemoveMultipleExtensions returns true if the function "WEBEXT.developerPrivate.removeMultipleExtensions" exists.
func HasFuncRemoveMultipleExtensions() bool {
	return js.True == bindings.HasFuncRemoveMultipleExtensions()
}

// FuncRemoveMultipleExtensions returns the function "WEBEXT.developerPrivate.removeMultipleExtensions".
func FuncRemoveMultipleExtensions() (fn js.Func[func(extensionIds js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncRemoveMultipleExtensions(
		js.Pointer(&fn),
	)
	return
}

// RemoveMultipleExtensions calls the function "WEBEXT.developerPrivate.removeMultipleExtensions" directly.
func RemoveMultipleExtensions(extensionIds js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallRemoveMultipleExtensions(
		js.Pointer(&ret),
		extensionIds.Ref(),
	)

	return
}

// TryRemoveMultipleExtensions calls the function "WEBEXT.developerPrivate.removeMultipleExtensions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveMultipleExtensions(extensionIds js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveMultipleExtensions(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionIds.Ref(),
	)

	return
}

// HasFuncRemoveUserSpecifiedSites returns true if the function "WEBEXT.developerPrivate.removeUserSpecifiedSites" exists.
func HasFuncRemoveUserSpecifiedSites() bool {
	return js.True == bindings.HasFuncRemoveUserSpecifiedSites()
}

// FuncRemoveUserSpecifiedSites returns the function "WEBEXT.developerPrivate.removeUserSpecifiedSites".
func FuncRemoveUserSpecifiedSites() (fn js.Func[func(options UserSiteSettingsOptions) js.Promise[js.Void]]) {
	bindings.FuncRemoveUserSpecifiedSites(
		js.Pointer(&fn),
	)
	return
}

// RemoveUserSpecifiedSites calls the function "WEBEXT.developerPrivate.removeUserSpecifiedSites" directly.
func RemoveUserSpecifiedSites(options UserSiteSettingsOptions) (ret js.Promise[js.Void]) {
	bindings.CallRemoveUserSpecifiedSites(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRemoveUserSpecifiedSites calls the function "WEBEXT.developerPrivate.removeUserSpecifiedSites"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveUserSpecifiedSites(options UserSiteSettingsOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveUserSpecifiedSites(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRepairExtension returns true if the function "WEBEXT.developerPrivate.repairExtension" exists.
func HasFuncRepairExtension() bool {
	return js.True == bindings.HasFuncRepairExtension()
}

// FuncRepairExtension returns the function "WEBEXT.developerPrivate.repairExtension".
func FuncRepairExtension() (fn js.Func[func(extensionId js.String) js.Promise[js.Void]]) {
	bindings.FuncRepairExtension(
		js.Pointer(&fn),
	)
	return
}

// RepairExtension calls the function "WEBEXT.developerPrivate.repairExtension" directly.
func RepairExtension(extensionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallRepairExtension(
		js.Pointer(&ret),
		extensionId.Ref(),
	)

	return
}

// TryRepairExtension calls the function "WEBEXT.developerPrivate.repairExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRepairExtension(extensionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRepairExtension(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
	)

	return
}

// HasFuncRequestFileSource returns true if the function "WEBEXT.developerPrivate.requestFileSource" exists.
func HasFuncRequestFileSource() bool {
	return js.True == bindings.HasFuncRequestFileSource()
}

// FuncRequestFileSource returns the function "WEBEXT.developerPrivate.requestFileSource".
func FuncRequestFileSource() (fn js.Func[func(properties RequestFileSourceProperties) js.Promise[RequestFileSourceResponse]]) {
	bindings.FuncRequestFileSource(
		js.Pointer(&fn),
	)
	return
}

// RequestFileSource calls the function "WEBEXT.developerPrivate.requestFileSource" directly.
func RequestFileSource(properties RequestFileSourceProperties) (ret js.Promise[RequestFileSourceResponse]) {
	bindings.CallRequestFileSource(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryRequestFileSource calls the function "WEBEXT.developerPrivate.requestFileSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestFileSource(properties RequestFileSourceProperties) (ret js.Promise[RequestFileSourceResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestFileSource(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncSetShortcutHandlingSuspended returns true if the function "WEBEXT.developerPrivate.setShortcutHandlingSuspended" exists.
func HasFuncSetShortcutHandlingSuspended() bool {
	return js.True == bindings.HasFuncSetShortcutHandlingSuspended()
}

// FuncSetShortcutHandlingSuspended returns the function "WEBEXT.developerPrivate.setShortcutHandlingSuspended".
func FuncSetShortcutHandlingSuspended() (fn js.Func[func(isSuspended bool) js.Promise[js.Void]]) {
	bindings.FuncSetShortcutHandlingSuspended(
		js.Pointer(&fn),
	)
	return
}

// SetShortcutHandlingSuspended calls the function "WEBEXT.developerPrivate.setShortcutHandlingSuspended" directly.
func SetShortcutHandlingSuspended(isSuspended bool) (ret js.Promise[js.Void]) {
	bindings.CallSetShortcutHandlingSuspended(
		js.Pointer(&ret),
		js.Bool(bool(isSuspended)),
	)

	return
}

// TrySetShortcutHandlingSuspended calls the function "WEBEXT.developerPrivate.setShortcutHandlingSuspended"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShortcutHandlingSuspended(isSuspended bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShortcutHandlingSuspended(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(isSuspended)),
	)

	return
}

// HasFuncShowOptions returns true if the function "WEBEXT.developerPrivate.showOptions" exists.
func HasFuncShowOptions() bool {
	return js.True == bindings.HasFuncShowOptions()
}

// FuncShowOptions returns the function "WEBEXT.developerPrivate.showOptions".
func FuncShowOptions() (fn js.Func[func(extensionId js.String) js.Promise[js.Void]]) {
	bindings.FuncShowOptions(
		js.Pointer(&fn),
	)
	return
}

// ShowOptions calls the function "WEBEXT.developerPrivate.showOptions" directly.
func ShowOptions(extensionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallShowOptions(
		js.Pointer(&ret),
		extensionId.Ref(),
	)

	return
}

// TryShowOptions calls the function "WEBEXT.developerPrivate.showOptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowOptions(extensionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowOptions(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
	)

	return
}

// HasFuncShowPath returns true if the function "WEBEXT.developerPrivate.showPath" exists.
func HasFuncShowPath() bool {
	return js.True == bindings.HasFuncShowPath()
}

// FuncShowPath returns the function "WEBEXT.developerPrivate.showPath".
func FuncShowPath() (fn js.Func[func(extensionId js.String) js.Promise[js.Void]]) {
	bindings.FuncShowPath(
		js.Pointer(&fn),
	)
	return
}

// ShowPath calls the function "WEBEXT.developerPrivate.showPath" directly.
func ShowPath(extensionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallShowPath(
		js.Pointer(&ret),
		extensionId.Ref(),
	)

	return
}

// TryShowPath calls the function "WEBEXT.developerPrivate.showPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowPath(extensionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowPath(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
	)

	return
}

// HasFuncUpdateExtensionCommand returns true if the function "WEBEXT.developerPrivate.updateExtensionCommand" exists.
func HasFuncUpdateExtensionCommand() bool {
	return js.True == bindings.HasFuncUpdateExtensionCommand()
}

// FuncUpdateExtensionCommand returns the function "WEBEXT.developerPrivate.updateExtensionCommand".
func FuncUpdateExtensionCommand() (fn js.Func[func(update ExtensionCommandUpdate) js.Promise[js.Void]]) {
	bindings.FuncUpdateExtensionCommand(
		js.Pointer(&fn),
	)
	return
}

// UpdateExtensionCommand calls the function "WEBEXT.developerPrivate.updateExtensionCommand" directly.
func UpdateExtensionCommand(update ExtensionCommandUpdate) (ret js.Promise[js.Void]) {
	bindings.CallUpdateExtensionCommand(
		js.Pointer(&ret),
		js.Pointer(&update),
	)

	return
}

// TryUpdateExtensionCommand calls the function "WEBEXT.developerPrivate.updateExtensionCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateExtensionCommand(update ExtensionCommandUpdate) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateExtensionCommand(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&update),
	)

	return
}

// HasFuncUpdateExtensionConfiguration returns true if the function "WEBEXT.developerPrivate.updateExtensionConfiguration" exists.
func HasFuncUpdateExtensionConfiguration() bool {
	return js.True == bindings.HasFuncUpdateExtensionConfiguration()
}

// FuncUpdateExtensionConfiguration returns the function "WEBEXT.developerPrivate.updateExtensionConfiguration".
func FuncUpdateExtensionConfiguration() (fn js.Func[func(update ExtensionConfigurationUpdate) js.Promise[js.Void]]) {
	bindings.FuncUpdateExtensionConfiguration(
		js.Pointer(&fn),
	)
	return
}

// UpdateExtensionConfiguration calls the function "WEBEXT.developerPrivate.updateExtensionConfiguration" directly.
func UpdateExtensionConfiguration(update ExtensionConfigurationUpdate) (ret js.Promise[js.Void]) {
	bindings.CallUpdateExtensionConfiguration(
		js.Pointer(&ret),
		js.Pointer(&update),
	)

	return
}

// TryUpdateExtensionConfiguration calls the function "WEBEXT.developerPrivate.updateExtensionConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateExtensionConfiguration(update ExtensionConfigurationUpdate) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateExtensionConfiguration(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&update),
	)

	return
}

// HasFuncUpdateProfileConfiguration returns true if the function "WEBEXT.developerPrivate.updateProfileConfiguration" exists.
func HasFuncUpdateProfileConfiguration() bool {
	return js.True == bindings.HasFuncUpdateProfileConfiguration()
}

// FuncUpdateProfileConfiguration returns the function "WEBEXT.developerPrivate.updateProfileConfiguration".
func FuncUpdateProfileConfiguration() (fn js.Func[func(update ProfileConfigurationUpdate) js.Promise[js.Void]]) {
	bindings.FuncUpdateProfileConfiguration(
		js.Pointer(&fn),
	)
	return
}

// UpdateProfileConfiguration calls the function "WEBEXT.developerPrivate.updateProfileConfiguration" directly.
func UpdateProfileConfiguration(update ProfileConfigurationUpdate) (ret js.Promise[js.Void]) {
	bindings.CallUpdateProfileConfiguration(
		js.Pointer(&ret),
		js.Pointer(&update),
	)

	return
}

// TryUpdateProfileConfiguration calls the function "WEBEXT.developerPrivate.updateProfileConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateProfileConfiguration(update ProfileConfigurationUpdate) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateProfileConfiguration(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&update),
	)

	return
}

// HasFuncUpdateSiteAccess returns true if the function "WEBEXT.developerPrivate.updateSiteAccess" exists.
func HasFuncUpdateSiteAccess() bool {
	return js.True == bindings.HasFuncUpdateSiteAccess()
}

// FuncUpdateSiteAccess returns the function "WEBEXT.developerPrivate.updateSiteAccess".
func FuncUpdateSiteAccess() (fn js.Func[func(site js.String, updates js.Array[ExtensionSiteAccessUpdate]) js.Promise[js.Void]]) {
	bindings.FuncUpdateSiteAccess(
		js.Pointer(&fn),
	)
	return
}

// UpdateSiteAccess calls the function "WEBEXT.developerPrivate.updateSiteAccess" directly.
func UpdateSiteAccess(site js.String, updates js.Array[ExtensionSiteAccessUpdate]) (ret js.Promise[js.Void]) {
	bindings.CallUpdateSiteAccess(
		js.Pointer(&ret),
		site.Ref(),
		updates.Ref(),
	)

	return
}

// TryUpdateSiteAccess calls the function "WEBEXT.developerPrivate.updateSiteAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateSiteAccess(site js.String, updates js.Array[ExtensionSiteAccessUpdate]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateSiteAccess(
		js.Pointer(&ret), js.Pointer(&exception),
		site.Ref(),
		updates.Ref(),
	)

	return
}

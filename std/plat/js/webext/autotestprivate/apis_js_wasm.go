// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package autotestprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/autotestprivate/bindings"
)

type Accelerator struct {
	// KeyCode is "Accelerator.keyCode"
	//
	// Optional
	KeyCode js.String
	// Shift is "Accelerator.shift"
	//
	// Optional
	//
	// NOTE: FFI_USE_Shift MUST be set to true to make this field effective.
	Shift bool
	// Control is "Accelerator.control"
	//
	// Optional
	//
	// NOTE: FFI_USE_Control MUST be set to true to make this field effective.
	Control bool
	// Alt is "Accelerator.alt"
	//
	// Optional
	//
	// NOTE: FFI_USE_Alt MUST be set to true to make this field effective.
	Alt bool
	// Search is "Accelerator.search"
	//
	// Optional
	//
	// NOTE: FFI_USE_Search MUST be set to true to make this field effective.
	Search bool
	// Pressed is "Accelerator.pressed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pressed MUST be set to true to make this field effective.
	Pressed bool

	FFI_USE_Shift   bool // for Shift.
	FFI_USE_Control bool // for Control.
	FFI_USE_Alt     bool // for Alt.
	FFI_USE_Search  bool // for Search.
	FFI_USE_Pressed bool // for Pressed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Accelerator with all fields set.
func (p Accelerator) FromRef(ref js.Ref) Accelerator {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Accelerator in the application heap.
func (p Accelerator) New() js.Ref {
	return bindings.AcceleratorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Accelerator) UpdateFrom(ref js.Ref) {
	bindings.AcceleratorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Accelerator) Update(ref js.Ref) {
	bindings.AcceleratorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Accelerator) FreeMembers(recursive bool) {
	js.Free(
		p.KeyCode.Ref(),
	)
	p.KeyCode = p.KeyCode.FromRef(js.Undefined)
}

type AcceleratorCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn AcceleratorCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AcceleratorCallbackFunc) DispatchCallback(
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

type AcceleratorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *AcceleratorCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AcceleratorCallback[T]) DispatchCallback(
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

type AllEnterprisePoliciesCallbackFunc func(this js.Ref, all_policies js.Any) js.Ref

func (fn AllEnterprisePoliciesCallbackFunc) Register() js.Func[func(all_policies js.Any)] {
	return js.RegisterCallback[func(all_policies js.Any)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AllEnterprisePoliciesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AllEnterprisePoliciesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, all_policies js.Any) js.Ref
	Arg T
}

func (cb *AllEnterprisePoliciesCallback[T]) Register() js.Func[func(all_policies js.Any)] {
	return js.RegisterCallback[func(all_policies js.Any)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AllEnterprisePoliciesCallback[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AppType uint32

const (
	_ AppType = iota

	AppType_ARC
	AppType_BUILT_IN
	AppType_CROSTINI
	AppType_EXTENSION
	AppType_WEB
	AppType_MAC_OS
	AppType_PLUGIN_VM
	AppType_STANDALONE_BROWSER
	AppType_REMOTE
	AppType_BOREALIS
	AppType_BRUSCHETTA
)

func (AppType) FromRef(str js.Ref) AppType {
	return AppType(bindings.ConstOfAppType(str))
}

func (x AppType) String() (string, bool) {
	switch x {
	case AppType_ARC:
		return "Arc", true
	case AppType_BUILT_IN:
		return "BuiltIn", true
	case AppType_CROSTINI:
		return "Crostini", true
	case AppType_EXTENSION:
		return "Extension", true
	case AppType_WEB:
		return "Web", true
	case AppType_MAC_OS:
		return "MacOS", true
	case AppType_PLUGIN_VM:
		return "PluginVm", true
	case AppType_STANDALONE_BROWSER:
		return "StandaloneBrowser", true
	case AppType_REMOTE:
		return "Remote", true
	case AppType_BOREALIS:
		return "Borealis", true
	case AppType_BRUSCHETTA:
		return "Bruschetta", true
	default:
		return "", false
	}
}

type AppInstallSource uint32

const (
	_ AppInstallSource = iota

	AppInstallSource_UNKNOWN
	AppInstallSource_SYSTEM
	AppInstallSource_POLICY
	AppInstallSource_OEM
	AppInstallSource_DEFAULT
	AppInstallSource_SYNC
	AppInstallSource_USER
	AppInstallSource_SUB_APP
	AppInstallSource_KIOSK
	AppInstallSource_COMMAND_LINE
)

func (AppInstallSource) FromRef(str js.Ref) AppInstallSource {
	return AppInstallSource(bindings.ConstOfAppInstallSource(str))
}

func (x AppInstallSource) String() (string, bool) {
	switch x {
	case AppInstallSource_UNKNOWN:
		return "Unknown", true
	case AppInstallSource_SYSTEM:
		return "System", true
	case AppInstallSource_POLICY:
		return "Policy", true
	case AppInstallSource_OEM:
		return "Oem", true
	case AppInstallSource_DEFAULT:
		return "Default", true
	case AppInstallSource_SYNC:
		return "Sync", true
	case AppInstallSource_USER:
		return "User", true
	case AppInstallSource_SUB_APP:
		return "SubApp", true
	case AppInstallSource_KIOSK:
		return "Kiosk", true
	case AppInstallSource_COMMAND_LINE:
		return "CommandLine", true
	default:
		return "", false
	}
}

type AppReadiness uint32

const (
	_ AppReadiness = iota

	AppReadiness_READY
	AppReadiness_DISABLED_BY_BLACKLIST
	AppReadiness_DISABLED_BY_POLICY
	AppReadiness_DISABLED_BY_USER
	AppReadiness_TERMINATED
	AppReadiness_UNINSTALLED_BY_USER
	AppReadiness_REMOVED
	AppReadiness_UNINSTALLED_BY_MIGRATION
)

func (AppReadiness) FromRef(str js.Ref) AppReadiness {
	return AppReadiness(bindings.ConstOfAppReadiness(str))
}

func (x AppReadiness) String() (string, bool) {
	switch x {
	case AppReadiness_READY:
		return "Ready", true
	case AppReadiness_DISABLED_BY_BLACKLIST:
		return "DisabledByBlacklist", true
	case AppReadiness_DISABLED_BY_POLICY:
		return "DisabledByPolicy", true
	case AppReadiness_DISABLED_BY_USER:
		return "DisabledByUser", true
	case AppReadiness_TERMINATED:
		return "Terminated", true
	case AppReadiness_UNINSTALLED_BY_USER:
		return "UninstalledByUser", true
	case AppReadiness_REMOVED:
		return "Removed", true
	case AppReadiness_UNINSTALLED_BY_MIGRATION:
		return "UninstalledByMigration", true
	default:
		return "", false
	}
}

type App struct {
	// AppId is "App.appId"
	//
	// Optional
	AppId js.String
	// Name is "App.name"
	//
	// Optional
	Name js.String
	// ShortName is "App.shortName"
	//
	// Optional
	ShortName js.String
	// PublisherId is "App.publisherId"
	//
	// Optional
	PublisherId js.String
	// Type is "App.type"
	//
	// Optional
	Type AppType
	// InstallSource is "App.installSource"
	//
	// Optional
	InstallSource AppInstallSource
	// Readiness is "App.readiness"
	//
	// Optional
	Readiness AppReadiness
	// AdditionalSearchTerms is "App.additionalSearchTerms"
	//
	// Optional
	AdditionalSearchTerms js.Array[js.String]
	// ShowInLauncher is "App.showInLauncher"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowInLauncher MUST be set to true to make this field effective.
	ShowInLauncher bool
	// ShowInSearch is "App.showInSearch"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowInSearch MUST be set to true to make this field effective.
	ShowInSearch bool

	FFI_USE_ShowInLauncher bool // for ShowInLauncher.
	FFI_USE_ShowInSearch   bool // for ShowInSearch.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a App with all fields set.
func (p App) FromRef(ref js.Ref) App {
	p.UpdateFrom(ref)
	return p
}

// New creates a new App in the application heap.
func (p App) New() js.Ref {
	return bindings.AppJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *App) UpdateFrom(ref js.Ref) {
	bindings.AppJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *App) Update(ref js.Ref) {
	bindings.AppJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *App) FreeMembers(recursive bool) {
	js.Free(
		p.AppId.Ref(),
		p.Name.Ref(),
		p.ShortName.Ref(),
		p.PublisherId.Ref(),
		p.AdditionalSearchTerms.Ref(),
	)
	p.AppId = p.AppId.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.ShortName = p.ShortName.FromRef(js.Undefined)
	p.PublisherId = p.PublisherId.FromRef(js.Undefined)
	p.AdditionalSearchTerms = p.AdditionalSearchTerms.FromRef(js.Undefined)
}

type AppWindowType uint32

const (
	_ AppWindowType = iota

	AppWindowType_BROWSER
	AppWindowType_CHROME_APP
	AppWindowType_ARC_APP
	AppWindowType_CROSTINI_APP
	AppWindowType_SYSTEM_APP
	AppWindowType_EXTENSION_APP
	AppWindowType_LACROS
)

func (AppWindowType) FromRef(str js.Ref) AppWindowType {
	return AppWindowType(bindings.ConstOfAppWindowType(str))
}

func (x AppWindowType) String() (string, bool) {
	switch x {
	case AppWindowType_BROWSER:
		return "Browser", true
	case AppWindowType_CHROME_APP:
		return "ChromeApp", true
	case AppWindowType_ARC_APP:
		return "ArcApp", true
	case AppWindowType_CROSTINI_APP:
		return "CrostiniApp", true
	case AppWindowType_SYSTEM_APP:
		return "SystemApp", true
	case AppWindowType_EXTENSION_APP:
		return "ExtensionApp", true
	case AppWindowType_LACROS:
		return "Lacros", true
	default:
		return "", false
	}
}

type WindowStateType uint32

const (
	_ WindowStateType = iota

	WindowStateType_NORMAL
	WindowStateType_MINIMIZED
	WindowStateType_MAXIMIZED
	WindowStateType_FULLSCREEN
	WindowStateType_PRIMARY_SNAPPED
	WindowStateType_SECONDARY_SNAPPED
	WindowStateType_PIP
	WindowStateType_FLOATED
)

func (WindowStateType) FromRef(str js.Ref) WindowStateType {
	return WindowStateType(bindings.ConstOfWindowStateType(str))
}

func (x WindowStateType) String() (string, bool) {
	switch x {
	case WindowStateType_NORMAL:
		return "Normal", true
	case WindowStateType_MINIMIZED:
		return "Minimized", true
	case WindowStateType_MAXIMIZED:
		return "Maximized", true
	case WindowStateType_FULLSCREEN:
		return "Fullscreen", true
	case WindowStateType_PRIMARY_SNAPPED:
		return "PrimarySnapped", true
	case WindowStateType_SECONDARY_SNAPPED:
		return "SecondarySnapped", true
	case WindowStateType_PIP:
		return "PIP", true
	case WindowStateType_FLOATED:
		return "Floated", true
	default:
		return "", false
	}
}

type Bounds struct {
	// Left is "Bounds.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left float64
	// Top is "Bounds.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top float64
	// Width is "Bounds.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float64
	// Height is "Bounds.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float64

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Bounds with all fields set.
func (p Bounds) FromRef(ref js.Ref) Bounds {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Bounds in the application heap.
func (p Bounds) New() js.Ref {
	return bindings.BoundsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Bounds) UpdateFrom(ref js.Ref) {
	bindings.BoundsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Bounds) Update(ref js.Ref) {
	bindings.BoundsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Bounds) FreeMembers(recursive bool) {
}

type FrameMode uint32

const (
	_ FrameMode = iota

	FrameMode_NORMAL
	FrameMode_IMMERSIVE
)

func (FrameMode) FromRef(str js.Ref) FrameMode {
	return FrameMode(bindings.ConstOfFrameMode(str))
}

func (x FrameMode) String() (string, bool) {
	switch x {
	case FrameMode_NORMAL:
		return "Normal", true
	case FrameMode_IMMERSIVE:
		return "Immersive", true
	default:
		return "", false
	}
}

type OverviewInfo struct {
	// Bounds is "OverviewInfo.bounds"
	//
	// Optional
	//
	// NOTE: Bounds.FFI_USE MUST be set to true to get Bounds used.
	Bounds Bounds
	// IsDragged is "OverviewInfo.isDragged"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDragged MUST be set to true to make this field effective.
	IsDragged bool

	FFI_USE_IsDragged bool // for IsDragged.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OverviewInfo with all fields set.
func (p OverviewInfo) FromRef(ref js.Ref) OverviewInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OverviewInfo in the application heap.
func (p OverviewInfo) New() js.Ref {
	return bindings.OverviewInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OverviewInfo) UpdateFrom(ref js.Ref) {
	bindings.OverviewInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OverviewInfo) Update(ref js.Ref) {
	bindings.OverviewInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OverviewInfo) FreeMembers(recursive bool) {
	if recursive {
		p.Bounds.FreeMembers(true)
	}
}

type AppWindowInfo struct {
	// Id is "AppWindowInfo.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Name is "AppWindowInfo.name"
	//
	// Optional
	Name js.String
	// WindowType is "AppWindowInfo.windowType"
	//
	// Optional
	WindowType AppWindowType
	// StateType is "AppWindowInfo.stateType"
	//
	// Optional
	StateType WindowStateType
	// BoundsInRoot is "AppWindowInfo.boundsInRoot"
	//
	// Optional
	//
	// NOTE: BoundsInRoot.FFI_USE MUST be set to true to get BoundsInRoot used.
	BoundsInRoot Bounds
	// DisplayId is "AppWindowInfo.displayId"
	//
	// Optional
	DisplayId js.String
	// IsVisible is "AppWindowInfo.isVisible"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsVisible MUST be set to true to make this field effective.
	IsVisible bool
	// CanFocus is "AppWindowInfo.canFocus"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanFocus MUST be set to true to make this field effective.
	CanFocus bool
	// Title is "AppWindowInfo.title"
	//
	// Optional
	Title js.String
	// IsAnimating is "AppWindowInfo.isAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAnimating MUST be set to true to make this field effective.
	IsAnimating bool
	// TargetBounds is "AppWindowInfo.targetBounds"
	//
	// Optional
	//
	// NOTE: TargetBounds.FFI_USE MUST be set to true to get TargetBounds used.
	TargetBounds Bounds
	// TargetVisibility is "AppWindowInfo.targetVisibility"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetVisibility MUST be set to true to make this field effective.
	TargetVisibility bool
	// IsActive is "AppWindowInfo.isActive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsActive MUST be set to true to make this field effective.
	IsActive bool
	// HasFocus is "AppWindowInfo.hasFocus"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasFocus MUST be set to true to make this field effective.
	HasFocus bool
	// OnActiveDesk is "AppWindowInfo.onActiveDesk"
	//
	// Optional
	//
	// NOTE: FFI_USE_OnActiveDesk MUST be set to true to make this field effective.
	OnActiveDesk bool
	// HasCapture is "AppWindowInfo.hasCapture"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasCapture MUST be set to true to make this field effective.
	HasCapture bool
	// CanResize is "AppWindowInfo.canResize"
	//
	// Optional
	//
	// NOTE: FFI_USE_CanResize MUST be set to true to make this field effective.
	CanResize bool
	// StackingOrder is "AppWindowInfo.stackingOrder"
	//
	// Optional
	//
	// NOTE: FFI_USE_StackingOrder MUST be set to true to make this field effective.
	StackingOrder int32
	// FrameMode is "AppWindowInfo.frameMode"
	//
	// Optional
	FrameMode FrameMode
	// IsFrameVisible is "AppWindowInfo.isFrameVisible"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsFrameVisible MUST be set to true to make this field effective.
	IsFrameVisible bool
	// CaptionHeight is "AppWindowInfo.captionHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_CaptionHeight MUST be set to true to make this field effective.
	CaptionHeight int32
	// CaptionButtonEnabledStatus is "AppWindowInfo.captionButtonEnabledStatus"
	//
	// Optional
	//
	// NOTE: FFI_USE_CaptionButtonEnabledStatus MUST be set to true to make this field effective.
	CaptionButtonEnabledStatus int32
	// CaptionButtonVisibleStatus is "AppWindowInfo.captionButtonVisibleStatus"
	//
	// Optional
	//
	// NOTE: FFI_USE_CaptionButtonVisibleStatus MUST be set to true to make this field effective.
	CaptionButtonVisibleStatus int32
	// ArcPackageName is "AppWindowInfo.arcPackageName"
	//
	// Optional
	ArcPackageName js.String
	// OverviewInfo is "AppWindowInfo.overviewInfo"
	//
	// Optional
	//
	// NOTE: OverviewInfo.FFI_USE MUST be set to true to get OverviewInfo used.
	OverviewInfo OverviewInfo
	// FullRestoreWindowAppId is "AppWindowInfo.fullRestoreWindowAppId"
	//
	// Optional
	FullRestoreWindowAppId js.String
	// AppId is "AppWindowInfo.appId"
	//
	// Optional
	AppId js.String

	FFI_USE_Id                         bool // for Id.
	FFI_USE_IsVisible                  bool // for IsVisible.
	FFI_USE_CanFocus                   bool // for CanFocus.
	FFI_USE_IsAnimating                bool // for IsAnimating.
	FFI_USE_TargetVisibility           bool // for TargetVisibility.
	FFI_USE_IsActive                   bool // for IsActive.
	FFI_USE_HasFocus                   bool // for HasFocus.
	FFI_USE_OnActiveDesk               bool // for OnActiveDesk.
	FFI_USE_HasCapture                 bool // for HasCapture.
	FFI_USE_CanResize                  bool // for CanResize.
	FFI_USE_StackingOrder              bool // for StackingOrder.
	FFI_USE_IsFrameVisible             bool // for IsFrameVisible.
	FFI_USE_CaptionHeight              bool // for CaptionHeight.
	FFI_USE_CaptionButtonEnabledStatus bool // for CaptionButtonEnabledStatus.
	FFI_USE_CaptionButtonVisibleStatus bool // for CaptionButtonVisibleStatus.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AppWindowInfo with all fields set.
func (p AppWindowInfo) FromRef(ref js.Ref) AppWindowInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AppWindowInfo in the application heap.
func (p AppWindowInfo) New() js.Ref {
	return bindings.AppWindowInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AppWindowInfo) UpdateFrom(ref js.Ref) {
	bindings.AppWindowInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AppWindowInfo) Update(ref js.Ref) {
	bindings.AppWindowInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AppWindowInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.DisplayId.Ref(),
		p.Title.Ref(),
		p.ArcPackageName.Ref(),
		p.FullRestoreWindowAppId.Ref(),
		p.AppId.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.DisplayId = p.DisplayId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.ArcPackageName = p.ArcPackageName.FromRef(js.Undefined)
	p.FullRestoreWindowAppId = p.FullRestoreWindowAppId.FromRef(js.Undefined)
	p.AppId = p.AppId.FromRef(js.Undefined)
	if recursive {
		p.BoundsInRoot.FreeMembers(true)
		p.TargetBounds.FreeMembers(true)
		p.OverviewInfo.FreeMembers(true)
	}
}

type ArcAppDict struct {
	// Name is "ArcAppDict.name"
	//
	// Optional
	Name js.String
	// PackageName is "ArcAppDict.packageName"
	//
	// Optional
	PackageName js.String
	// Activity is "ArcAppDict.activity"
	//
	// Optional
	Activity js.String
	// IntentUri is "ArcAppDict.intentUri"
	//
	// Optional
	IntentUri js.String
	// IconResourceId is "ArcAppDict.iconResourceId"
	//
	// Optional
	IconResourceId js.String
	// LastLaunchTime is "ArcAppDict.lastLaunchTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastLaunchTime MUST be set to true to make this field effective.
	LastLaunchTime float64
	// InstallTime is "ArcAppDict.installTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_InstallTime MUST be set to true to make this field effective.
	InstallTime float64
	// Sticky is "ArcAppDict.sticky"
	//
	// Optional
	//
	// NOTE: FFI_USE_Sticky MUST be set to true to make this field effective.
	Sticky bool
	// NotificationsEnabled is "ArcAppDict.notificationsEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_NotificationsEnabled MUST be set to true to make this field effective.
	NotificationsEnabled bool
	// Ready is "ArcAppDict.ready"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ready MUST be set to true to make this field effective.
	Ready bool
	// Suspended is "ArcAppDict.suspended"
	//
	// Optional
	//
	// NOTE: FFI_USE_Suspended MUST be set to true to make this field effective.
	Suspended bool
	// ShowInLauncher is "ArcAppDict.showInLauncher"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowInLauncher MUST be set to true to make this field effective.
	ShowInLauncher bool
	// Shortcut is "ArcAppDict.shortcut"
	//
	// Optional
	//
	// NOTE: FFI_USE_Shortcut MUST be set to true to make this field effective.
	Shortcut bool
	// Launchable is "ArcAppDict.launchable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Launchable MUST be set to true to make this field effective.
	Launchable bool

	FFI_USE_LastLaunchTime       bool // for LastLaunchTime.
	FFI_USE_InstallTime          bool // for InstallTime.
	FFI_USE_Sticky               bool // for Sticky.
	FFI_USE_NotificationsEnabled bool // for NotificationsEnabled.
	FFI_USE_Ready                bool // for Ready.
	FFI_USE_Suspended            bool // for Suspended.
	FFI_USE_ShowInLauncher       bool // for ShowInLauncher.
	FFI_USE_Shortcut             bool // for Shortcut.
	FFI_USE_Launchable           bool // for Launchable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ArcAppDict with all fields set.
func (p ArcAppDict) FromRef(ref js.Ref) ArcAppDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ArcAppDict in the application heap.
func (p ArcAppDict) New() js.Ref {
	return bindings.ArcAppDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ArcAppDict) UpdateFrom(ref js.Ref) {
	bindings.ArcAppDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ArcAppDict) Update(ref js.Ref) {
	bindings.ArcAppDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ArcAppDict) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.PackageName.Ref(),
		p.Activity.Ref(),
		p.IntentUri.Ref(),
		p.IconResourceId.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PackageName = p.PackageName.FromRef(js.Undefined)
	p.Activity = p.Activity.FromRef(js.Undefined)
	p.IntentUri = p.IntentUri.FromRef(js.Undefined)
	p.IconResourceId = p.IconResourceId.FromRef(js.Undefined)
}

type ArcAppKillsDict struct {
	// Oom is "ArcAppKillsDict.oom"
	//
	// Optional
	//
	// NOTE: FFI_USE_Oom MUST be set to true to make this field effective.
	Oom float64
	// LmkdForeground is "ArcAppKillsDict.lmkdForeground"
	//
	// Optional
	//
	// NOTE: FFI_USE_LmkdForeground MUST be set to true to make this field effective.
	LmkdForeground float64
	// LmkdPerceptible is "ArcAppKillsDict.lmkdPerceptible"
	//
	// Optional
	//
	// NOTE: FFI_USE_LmkdPerceptible MUST be set to true to make this field effective.
	LmkdPerceptible float64
	// LmkdCached is "ArcAppKillsDict.lmkdCached"
	//
	// Optional
	//
	// NOTE: FFI_USE_LmkdCached MUST be set to true to make this field effective.
	LmkdCached float64
	// PressureForeground is "ArcAppKillsDict.pressureForeground"
	//
	// Optional
	//
	// NOTE: FFI_USE_PressureForeground MUST be set to true to make this field effective.
	PressureForeground float64
	// PressurePerceptible is "ArcAppKillsDict.pressurePerceptible"
	//
	// Optional
	//
	// NOTE: FFI_USE_PressurePerceptible MUST be set to true to make this field effective.
	PressurePerceptible float64
	// PressureCached is "ArcAppKillsDict.pressureCached"
	//
	// Optional
	//
	// NOTE: FFI_USE_PressureCached MUST be set to true to make this field effective.
	PressureCached float64

	FFI_USE_Oom                 bool // for Oom.
	FFI_USE_LmkdForeground      bool // for LmkdForeground.
	FFI_USE_LmkdPerceptible     bool // for LmkdPerceptible.
	FFI_USE_LmkdCached          bool // for LmkdCached.
	FFI_USE_PressureForeground  bool // for PressureForeground.
	FFI_USE_PressurePerceptible bool // for PressurePerceptible.
	FFI_USE_PressureCached      bool // for PressureCached.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ArcAppKillsDict with all fields set.
func (p ArcAppKillsDict) FromRef(ref js.Ref) ArcAppKillsDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ArcAppKillsDict in the application heap.
func (p ArcAppKillsDict) New() js.Ref {
	return bindings.ArcAppKillsDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ArcAppKillsDict) UpdateFrom(ref js.Ref) {
	bindings.ArcAppKillsDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ArcAppKillsDict) Update(ref js.Ref) {
	bindings.ArcAppKillsDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ArcAppKillsDict) FreeMembers(recursive bool) {
}

type ArcAppTracingCallbackFunc func(this js.Ref, info *ArcAppTracingInfo) js.Ref

func (fn ArcAppTracingCallbackFunc) Register() js.Func[func(info *ArcAppTracingInfo)] {
	return js.RegisterCallback[func(info *ArcAppTracingInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ArcAppTracingCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppTracingInfo
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

type ArcAppTracingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ArcAppTracingInfo) js.Ref
	Arg T
}

func (cb *ArcAppTracingCallback[T]) Register() js.Func[func(info *ArcAppTracingInfo)] {
	return js.RegisterCallback[func(info *ArcAppTracingInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ArcAppTracingCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppTracingInfo
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

type ArcAppTracingInfo struct {
	// Success is "ArcAppTracingInfo.success"
	//
	// Optional
	//
	// NOTE: FFI_USE_Success MUST be set to true to make this field effective.
	Success bool
	// Fps is "ArcAppTracingInfo.fps"
	//
	// Optional
	//
	// NOTE: FFI_USE_Fps MUST be set to true to make this field effective.
	Fps float64
	// CommitDeviation is "ArcAppTracingInfo.commitDeviation"
	//
	// Optional
	//
	// NOTE: FFI_USE_CommitDeviation MUST be set to true to make this field effective.
	CommitDeviation float64
	// RenderQuality is "ArcAppTracingInfo.renderQuality"
	//
	// Optional
	//
	// NOTE: FFI_USE_RenderQuality MUST be set to true to make this field effective.
	RenderQuality float64

	FFI_USE_Success         bool // for Success.
	FFI_USE_Fps             bool // for Fps.
	FFI_USE_CommitDeviation bool // for CommitDeviation.
	FFI_USE_RenderQuality   bool // for RenderQuality.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ArcAppTracingInfo with all fields set.
func (p ArcAppTracingInfo) FromRef(ref js.Ref) ArcAppTracingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ArcAppTracingInfo in the application heap.
func (p ArcAppTracingInfo) New() js.Ref {
	return bindings.ArcAppTracingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ArcAppTracingInfo) UpdateFrom(ref js.Ref) {
	bindings.ArcAppTracingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ArcAppTracingInfo) Update(ref js.Ref) {
	bindings.ArcAppTracingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ArcAppTracingInfo) FreeMembers(recursive bool) {
}

type ArcPackageDict struct {
	// PackageName is "ArcPackageDict.packageName"
	//
	// Optional
	PackageName js.String
	// PackageVersion is "ArcPackageDict.packageVersion"
	//
	// Optional
	//
	// NOTE: FFI_USE_PackageVersion MUST be set to true to make this field effective.
	PackageVersion int32
	// LastBackupAndroidId is "ArcPackageDict.lastBackupAndroidId"
	//
	// Optional
	LastBackupAndroidId js.String
	// LastBackupTime is "ArcPackageDict.lastBackupTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastBackupTime MUST be set to true to make this field effective.
	LastBackupTime float64
	// ShouldSync is "ArcPackageDict.shouldSync"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShouldSync MUST be set to true to make this field effective.
	ShouldSync bool
	// VpnProvider is "ArcPackageDict.vpnProvider"
	//
	// Optional
	//
	// NOTE: FFI_USE_VpnProvider MUST be set to true to make this field effective.
	VpnProvider bool

	FFI_USE_PackageVersion bool // for PackageVersion.
	FFI_USE_LastBackupTime bool // for LastBackupTime.
	FFI_USE_ShouldSync     bool // for ShouldSync.
	FFI_USE_VpnProvider    bool // for VpnProvider.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ArcPackageDict with all fields set.
func (p ArcPackageDict) FromRef(ref js.Ref) ArcPackageDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ArcPackageDict in the application heap.
func (p ArcPackageDict) New() js.Ref {
	return bindings.ArcPackageDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ArcPackageDict) UpdateFrom(ref js.Ref) {
	bindings.ArcPackageDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ArcPackageDict) Update(ref js.Ref) {
	bindings.ArcPackageDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ArcPackageDict) FreeMembers(recursive bool) {
	js.Free(
		p.PackageName.Ref(),
		p.LastBackupAndroidId.Ref(),
	)
	p.PackageName = p.PackageName.FromRef(js.Undefined)
	p.LastBackupAndroidId = p.LastBackupAndroidId.FromRef(js.Undefined)
}

type ArcStartTimeCallbackFunc func(this js.Ref, startTicks float64) js.Ref

func (fn ArcStartTimeCallbackFunc) Register() js.Func[func(startTicks float64)] {
	return js.RegisterCallback[func(startTicks float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ArcStartTimeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ArcStartTimeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, startTicks float64) js.Ref
	Arg T
}

func (cb *ArcStartTimeCallback[T]) Register() js.Func[func(startTicks float64)] {
	return js.RegisterCallback[func(startTicks float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ArcStartTimeCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ArcState struct {
	// Provisioned is "ArcState.provisioned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Provisioned MUST be set to true to make this field effective.
	Provisioned bool
	// TosNeeded is "ArcState.tosNeeded"
	//
	// Optional
	//
	// NOTE: FFI_USE_TosNeeded MUST be set to true to make this field effective.
	TosNeeded bool
	// PreStartTime is "ArcState.preStartTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_PreStartTime MUST be set to true to make this field effective.
	PreStartTime float64
	// StartTime is "ArcState.startTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartTime MUST be set to true to make this field effective.
	StartTime float64

	FFI_USE_Provisioned  bool // for Provisioned.
	FFI_USE_TosNeeded    bool // for TosNeeded.
	FFI_USE_PreStartTime bool // for PreStartTime.
	FFI_USE_StartTime    bool // for StartTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ArcState with all fields set.
func (p ArcState) FromRef(ref js.Ref) ArcState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ArcState in the application heap.
func (p ArcState) New() js.Ref {
	return bindings.ArcStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ArcState) UpdateFrom(ref js.Ref) {
	bindings.ArcStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ArcState) Update(ref js.Ref) {
	bindings.ArcStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ArcState) FreeMembers(recursive bool) {
}

type ArcStateCallbackFunc func(this js.Ref, result *ArcState) js.Ref

func (fn ArcStateCallbackFunc) Register() js.Func[func(result *ArcState)] {
	return js.RegisterCallback[func(result *ArcState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ArcStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcState
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

type ArcStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *ArcState) js.Ref
	Arg T
}

func (cb *ArcStateCallback[T]) Register() js.Func[func(result *ArcState)] {
	return js.RegisterCallback[func(result *ArcState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ArcStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcState
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

type AssistantQueryResponse struct {
	// Text is "AssistantQueryResponse.text"
	//
	// Optional
	Text js.String
	// HtmlResponse is "AssistantQueryResponse.htmlResponse"
	//
	// Optional
	HtmlResponse js.String
	// OpenUrl is "AssistantQueryResponse.openUrl"
	//
	// Optional
	OpenUrl js.String
	// OpenAppResponse is "AssistantQueryResponse.openAppResponse"
	//
	// Optional
	OpenAppResponse js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssistantQueryResponse with all fields set.
func (p AssistantQueryResponse) FromRef(ref js.Ref) AssistantQueryResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssistantQueryResponse in the application heap.
func (p AssistantQueryResponse) New() js.Ref {
	return bindings.AssistantQueryResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssistantQueryResponse) UpdateFrom(ref js.Ref) {
	bindings.AssistantQueryResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssistantQueryResponse) Update(ref js.Ref) {
	bindings.AssistantQueryResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssistantQueryResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
		p.HtmlResponse.Ref(),
		p.OpenUrl.Ref(),
		p.OpenAppResponse.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
	p.HtmlResponse = p.HtmlResponse.FromRef(js.Undefined)
	p.OpenUrl = p.OpenUrl.FromRef(js.Undefined)
	p.OpenAppResponse = p.OpenAppResponse.FromRef(js.Undefined)
}

type AssistantQueryStatus struct {
	// IsMicOpen is "AssistantQueryStatus.isMicOpen"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMicOpen MUST be set to true to make this field effective.
	IsMicOpen bool
	// QueryText is "AssistantQueryStatus.queryText"
	//
	// Optional
	QueryText js.String
	// QueryResponse is "AssistantQueryStatus.queryResponse"
	//
	// Optional
	//
	// NOTE: QueryResponse.FFI_USE MUST be set to true to get QueryResponse used.
	QueryResponse AssistantQueryResponse

	FFI_USE_IsMicOpen bool // for IsMicOpen.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AssistantQueryStatus with all fields set.
func (p AssistantQueryStatus) FromRef(ref js.Ref) AssistantQueryStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AssistantQueryStatus in the application heap.
func (p AssistantQueryStatus) New() js.Ref {
	return bindings.AssistantQueryStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AssistantQueryStatus) UpdateFrom(ref js.Ref) {
	bindings.AssistantQueryStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AssistantQueryStatus) Update(ref js.Ref) {
	bindings.AssistantQueryStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AssistantQueryStatus) FreeMembers(recursive bool) {
	js.Free(
		p.QueryText.Ref(),
	)
	p.QueryText = p.QueryText.FromRef(js.Undefined)
	if recursive {
		p.QueryResponse.FreeMembers(true)
	}
}

type AssistantQueryStatusCallbackFunc func(this js.Ref, status *AssistantQueryStatus) js.Ref

func (fn AssistantQueryStatusCallbackFunc) Register() js.Func[func(status *AssistantQueryStatus)] {
	return js.RegisterCallback[func(status *AssistantQueryStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AssistantQueryStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AssistantQueryStatus
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

type AssistantQueryStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *AssistantQueryStatus) js.Ref
	Arg T
}

func (cb *AssistantQueryStatusCallback[T]) Register() js.Func[func(status *AssistantQueryStatus)] {
	return js.RegisterCallback[func(status *AssistantQueryStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AssistantQueryStatusCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AssistantQueryStatus
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

type CouldAllowCrostiniCallbackFunc func(this js.Ref, canBeAllowed bool) js.Ref

func (fn CouldAllowCrostiniCallbackFunc) Register() js.Func[func(canBeAllowed bool)] {
	return js.RegisterCallback[func(canBeAllowed bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CouldAllowCrostiniCallbackFunc) DispatchCallback(
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

type CouldAllowCrostiniCallback[T any] struct {
	Fn  func(arg T, this js.Ref, canBeAllowed bool) js.Ref
	Arg T
}

func (cb *CouldAllowCrostiniCallback[T]) Register() js.Func[func(canBeAllowed bool)] {
	return js.RegisterCallback[func(canBeAllowed bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CouldAllowCrostiniCallback[T]) DispatchCallback(
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

type CryptohomeRecoveryDataDict struct {
	// ReauthProofToken is "CryptohomeRecoveryDataDict.reauthProofToken"
	//
	// Optional
	ReauthProofToken js.String
	// RefreshToken is "CryptohomeRecoveryDataDict.refreshToken"
	//
	// Optional
	RefreshToken js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CryptohomeRecoveryDataDict with all fields set.
func (p CryptohomeRecoveryDataDict) FromRef(ref js.Ref) CryptohomeRecoveryDataDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CryptohomeRecoveryDataDict in the application heap.
func (p CryptohomeRecoveryDataDict) New() js.Ref {
	return bindings.CryptohomeRecoveryDataDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CryptohomeRecoveryDataDict) UpdateFrom(ref js.Ref) {
	bindings.CryptohomeRecoveryDataDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CryptohomeRecoveryDataDict) Update(ref js.Ref) {
	bindings.CryptohomeRecoveryDataDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CryptohomeRecoveryDataDict) FreeMembers(recursive bool) {
	js.Free(
		p.ReauthProofToken.Ref(),
		p.RefreshToken.Ref(),
	)
	p.ReauthProofToken = p.ReauthProofToken.FromRef(js.Undefined)
	p.RefreshToken = p.RefreshToken.FromRef(js.Undefined)
}

type DOMStringCallbackFunc func(this js.Ref, data js.String) js.Ref

func (fn DOMStringCallbackFunc) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DOMStringCallbackFunc) DispatchCallback(
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

type DOMStringCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.String) js.Ref
	Arg T
}

func (cb *DOMStringCallback[T]) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DOMStringCallback[T]) DispatchCallback(
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

type DesksCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn DesksCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DesksCallbackFunc) DispatchCallback(
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

type DesksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *DesksCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DesksCallback[T]) DispatchCallback(
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

type DesksInfo struct {
	// ActiveDeskIndex is "DesksInfo.activeDeskIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_ActiveDeskIndex MUST be set to true to make this field effective.
	ActiveDeskIndex int32
	// NumDesks is "DesksInfo.numDesks"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumDesks MUST be set to true to make this field effective.
	NumDesks int32
	// IsAnimating is "DesksInfo.isAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAnimating MUST be set to true to make this field effective.
	IsAnimating bool
	// DeskContainers is "DesksInfo.deskContainers"
	//
	// Optional
	DeskContainers js.Array[js.String]

	FFI_USE_ActiveDeskIndex bool // for ActiveDeskIndex.
	FFI_USE_NumDesks        bool // for NumDesks.
	FFI_USE_IsAnimating     bool // for IsAnimating.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DesksInfo with all fields set.
func (p DesksInfo) FromRef(ref js.Ref) DesksInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DesksInfo in the application heap.
func (p DesksInfo) New() js.Ref {
	return bindings.DesksInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DesksInfo) UpdateFrom(ref js.Ref) {
	bindings.DesksInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DesksInfo) Update(ref js.Ref) {
	bindings.DesksInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DesksInfo) FreeMembers(recursive bool) {
	js.Free(
		p.DeskContainers.Ref(),
	)
	p.DeskContainers = p.DeskContainers.FromRef(js.Undefined)
}

type DisplaySmoothnessData struct {
	// FramesExpected is "DisplaySmoothnessData.framesExpected"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesExpected MUST be set to true to make this field effective.
	FramesExpected int32
	// FramesProduced is "DisplaySmoothnessData.framesProduced"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesProduced MUST be set to true to make this field effective.
	FramesProduced int32
	// JankCount is "DisplaySmoothnessData.jankCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_JankCount MUST be set to true to make this field effective.
	JankCount int32
	// Throughput is "DisplaySmoothnessData.throughput"
	//
	// Optional
	Throughput js.Array[int32]

	FFI_USE_FramesExpected bool // for FramesExpected.
	FFI_USE_FramesProduced bool // for FramesProduced.
	FFI_USE_JankCount      bool // for JankCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplaySmoothnessData with all fields set.
func (p DisplaySmoothnessData) FromRef(ref js.Ref) DisplaySmoothnessData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplaySmoothnessData in the application heap.
func (p DisplaySmoothnessData) New() js.Ref {
	return bindings.DisplaySmoothnessDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplaySmoothnessData) UpdateFrom(ref js.Ref) {
	bindings.DisplaySmoothnessDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplaySmoothnessData) Update(ref js.Ref) {
	bindings.DisplaySmoothnessDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplaySmoothnessData) FreeMembers(recursive bool) {
	js.Free(
		p.Throughput.Ref(),
	)
	p.Throughput = p.Throughput.FromRef(js.Undefined)
}

type ExtensionInfoDict struct {
	// Id is "ExtensionInfoDict.id"
	//
	// Optional
	Id js.String
	// Version is "ExtensionInfoDict.version"
	//
	// Optional
	Version js.String
	// Name is "ExtensionInfoDict.name"
	//
	// Optional
	Name js.String
	// PublicKey is "ExtensionInfoDict.publicKey"
	//
	// Optional
	PublicKey js.String
	// Description is "ExtensionInfoDict.description"
	//
	// Optional
	Description js.String
	// BackgroundUrl is "ExtensionInfoDict.backgroundUrl"
	//
	// Optional
	BackgroundUrl js.String
	// OptionsUrl is "ExtensionInfoDict.optionsUrl"
	//
	// Optional
	OptionsUrl js.String
	// HostPermissions is "ExtensionInfoDict.hostPermissions"
	//
	// Optional
	HostPermissions js.Array[js.String]
	// EffectiveHostPermissions is "ExtensionInfoDict.effectiveHostPermissions"
	//
	// Optional
	EffectiveHostPermissions js.Array[js.String]
	// ApiPermissions is "ExtensionInfoDict.apiPermissions"
	//
	// Optional
	ApiPermissions js.Array[js.String]
	// IsComponent is "ExtensionInfoDict.isComponent"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsComponent MUST be set to true to make this field effective.
	IsComponent bool
	// IsInternal is "ExtensionInfoDict.isInternal"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsInternal MUST be set to true to make this field effective.
	IsInternal bool
	// IsUserInstalled is "ExtensionInfoDict.isUserInstalled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUserInstalled MUST be set to true to make this field effective.
	IsUserInstalled bool
	// IsEnabled is "ExtensionInfoDict.isEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsEnabled MUST be set to true to make this field effective.
	IsEnabled bool
	// AllowedInIncognito is "ExtensionInfoDict.allowedInIncognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowedInIncognito MUST be set to true to make this field effective.
	AllowedInIncognito bool
	// HasPageAction is "ExtensionInfoDict.hasPageAction"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasPageAction MUST be set to true to make this field effective.
	HasPageAction bool

	FFI_USE_IsComponent        bool // for IsComponent.
	FFI_USE_IsInternal         bool // for IsInternal.
	FFI_USE_IsUserInstalled    bool // for IsUserInstalled.
	FFI_USE_IsEnabled          bool // for IsEnabled.
	FFI_USE_AllowedInIncognito bool // for AllowedInIncognito.
	FFI_USE_HasPageAction      bool // for HasPageAction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionInfoDict with all fields set.
func (p ExtensionInfoDict) FromRef(ref js.Ref) ExtensionInfoDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionInfoDict in the application heap.
func (p ExtensionInfoDict) New() js.Ref {
	return bindings.ExtensionInfoDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionInfoDict) UpdateFrom(ref js.Ref) {
	bindings.ExtensionInfoDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionInfoDict) Update(ref js.Ref) {
	bindings.ExtensionInfoDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionInfoDict) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Version.Ref(),
		p.Name.Ref(),
		p.PublicKey.Ref(),
		p.Description.Ref(),
		p.BackgroundUrl.Ref(),
		p.OptionsUrl.Ref(),
		p.HostPermissions.Ref(),
		p.EffectiveHostPermissions.Ref(),
		p.ApiPermissions.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PublicKey = p.PublicKey.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.BackgroundUrl = p.BackgroundUrl.FromRef(js.Undefined)
	p.OptionsUrl = p.OptionsUrl.FromRef(js.Undefined)
	p.HostPermissions = p.HostPermissions.FromRef(js.Undefined)
	p.EffectiveHostPermissions = p.EffectiveHostPermissions.FromRef(js.Undefined)
	p.ApiPermissions = p.ApiPermissions.FromRef(js.Undefined)
}

type ExtensionsInfoArray struct {
	// Extensions is "ExtensionsInfoArray.extensions"
	//
	// Optional
	Extensions js.Array[ExtensionInfoDict]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExtensionsInfoArray with all fields set.
func (p ExtensionsInfoArray) FromRef(ref js.Ref) ExtensionsInfoArray {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtensionsInfoArray in the application heap.
func (p ExtensionsInfoArray) New() js.Ref {
	return bindings.ExtensionsInfoArrayJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExtensionsInfoArray) UpdateFrom(ref js.Ref) {
	bindings.ExtensionsInfoArrayJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExtensionsInfoArray) Update(ref js.Ref) {
	bindings.ExtensionsInfoArrayJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExtensionsInfoArray) FreeMembers(recursive bool) {
	js.Free(
		p.Extensions.Ref(),
	)
	p.Extensions = p.Extensions.FromRef(js.Undefined)
}

type ExtensionsInfoCallbackFunc func(this js.Ref, info *ExtensionsInfoArray) js.Ref

func (fn ExtensionsInfoCallbackFunc) Register() js.Func[func(info *ExtensionsInfoArray)] {
	return js.RegisterCallback[func(info *ExtensionsInfoArray)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExtensionsInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionsInfoArray
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

type ExtensionsInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ExtensionsInfoArray) js.Ref
	Arg T
}

func (cb *ExtensionsInfoCallback[T]) Register() js.Func[func(info *ExtensionsInfoArray)] {
	return js.RegisterCallback[func(info *ExtensionsInfoArray)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExtensionsInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ExtensionsInfoArray
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

type FrameCountingPerSinkData struct {
	// SinkType is "FrameCountingPerSinkData.sinkType"
	//
	// Optional
	SinkType js.String
	// IsRoot is "FrameCountingPerSinkData.isRoot"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRoot MUST be set to true to make this field effective.
	IsRoot bool
	// DebugLabel is "FrameCountingPerSinkData.debugLabel"
	//
	// Optional
	DebugLabel js.String
	// PresentedFrames is "FrameCountingPerSinkData.presentedFrames"
	//
	// Optional
	PresentedFrames js.Array[int32]

	FFI_USE_IsRoot bool // for IsRoot.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FrameCountingPerSinkData with all fields set.
func (p FrameCountingPerSinkData) FromRef(ref js.Ref) FrameCountingPerSinkData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FrameCountingPerSinkData in the application heap.
func (p FrameCountingPerSinkData) New() js.Ref {
	return bindings.FrameCountingPerSinkDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FrameCountingPerSinkData) UpdateFrom(ref js.Ref) {
	bindings.FrameCountingPerSinkDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FrameCountingPerSinkData) Update(ref js.Ref) {
	bindings.FrameCountingPerSinkDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FrameCountingPerSinkData) FreeMembers(recursive bool) {
	js.Free(
		p.SinkType.Ref(),
		p.DebugLabel.Ref(),
		p.PresentedFrames.Ref(),
	)
	p.SinkType = p.SinkType.FromRef(js.Undefined)
	p.DebugLabel = p.DebugLabel.FromRef(js.Undefined)
	p.PresentedFrames = p.PresentedFrames.FromRef(js.Undefined)
}

type GetAccessTokenCallbackFunc func(this js.Ref, data *GetAccessTokenData) js.Ref

func (fn GetAccessTokenCallbackFunc) Register() js.Func[func(data *GetAccessTokenData)] {
	return js.RegisterCallback[func(data *GetAccessTokenData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAccessTokenCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAccessTokenData
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

type GetAccessTokenCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data *GetAccessTokenData) js.Ref
	Arg T
}

func (cb *GetAccessTokenCallback[T]) Register() js.Func[func(data *GetAccessTokenData)] {
	return js.RegisterCallback[func(data *GetAccessTokenData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAccessTokenCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetAccessTokenData
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

type GetAccessTokenData struct {
	// AccessToken is "GetAccessTokenData.accessToken"
	//
	// Optional
	AccessToken js.String
	// ExpirationTimeUnixMs is "GetAccessTokenData.expirationTimeUnixMs"
	//
	// Optional
	ExpirationTimeUnixMs js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAccessTokenData with all fields set.
func (p GetAccessTokenData) FromRef(ref js.Ref) GetAccessTokenData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAccessTokenData in the application heap.
func (p GetAccessTokenData) New() js.Ref {
	return bindings.GetAccessTokenDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAccessTokenData) UpdateFrom(ref js.Ref) {
	bindings.GetAccessTokenDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAccessTokenData) Update(ref js.Ref) {
	bindings.GetAccessTokenDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAccessTokenData) FreeMembers(recursive bool) {
	js.Free(
		p.AccessToken.Ref(),
		p.ExpirationTimeUnixMs.Ref(),
	)
	p.AccessToken = p.AccessToken.FromRef(js.Undefined)
	p.ExpirationTimeUnixMs = p.ExpirationTimeUnixMs.FromRef(js.Undefined)
}

type GetAccessTokenParams struct {
	// Email is "GetAccessTokenParams.email"
	//
	// Optional
	Email js.String
	// Scopes is "GetAccessTokenParams.scopes"
	//
	// Optional
	Scopes js.Array[js.String]
	// TimeoutMs is "GetAccessTokenParams.timeoutMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeoutMs MUST be set to true to make this field effective.
	TimeoutMs int32

	FFI_USE_TimeoutMs bool // for TimeoutMs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAccessTokenParams with all fields set.
func (p GetAccessTokenParams) FromRef(ref js.Ref) GetAccessTokenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAccessTokenParams in the application heap.
func (p GetAccessTokenParams) New() js.Ref {
	return bindings.GetAccessTokenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAccessTokenParams) UpdateFrom(ref js.Ref) {
	bindings.GetAccessTokenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAccessTokenParams) Update(ref js.Ref) {
	bindings.GetAccessTokenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAccessTokenParams) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.Scopes.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.Scopes = p.Scopes.FromRef(js.Undefined)
}

type GetAllInstalledAppsCallbackFunc func(this js.Ref, apps js.Array[App]) js.Ref

func (fn GetAllInstalledAppsCallbackFunc) Register() js.Func[func(apps js.Array[App])] {
	return js.RegisterCallback[func(apps js.Array[App])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAllInstalledAppsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[App]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAllInstalledAppsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, apps js.Array[App]) js.Ref
	Arg T
}

func (cb *GetAllInstalledAppsCallback[T]) Register() js.Func[func(apps js.Array[App])] {
	return js.RegisterCallback[func(apps js.Array[App])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAllInstalledAppsCallback[T]) DispatchCallback(
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

		js.Array[App]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAppWindowListCallbackFunc func(this js.Ref, window_list js.Array[AppWindowInfo]) js.Ref

func (fn GetAppWindowListCallbackFunc) Register() js.Func[func(window_list js.Array[AppWindowInfo])] {
	return js.RegisterCallback[func(window_list js.Array[AppWindowInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAppWindowListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AppWindowInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAppWindowListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, window_list js.Array[AppWindowInfo]) js.Ref
	Arg T
}

func (cb *GetAppWindowListCallback[T]) Register() js.Func[func(window_list js.Array[AppWindowInfo])] {
	return js.RegisterCallback[func(window_list js.Array[AppWindowInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAppWindowListCallback[T]) DispatchCallback(
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

		js.Array[AppWindowInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetArcAppCallbackFunc func(this js.Ref, pkg *ArcAppDict) js.Ref

func (fn GetArcAppCallbackFunc) Register() js.Func[func(pkg *ArcAppDict)] {
	return js.RegisterCallback[func(pkg *ArcAppDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetArcAppCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppDict
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

type GetArcAppCallback[T any] struct {
	Fn  func(arg T, this js.Ref, pkg *ArcAppDict) js.Ref
	Arg T
}

func (cb *GetArcAppCallback[T]) Register() js.Func[func(pkg *ArcAppDict)] {
	return js.RegisterCallback[func(pkg *ArcAppDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetArcAppCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppDict
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

type GetArcAppKillsCallbackFunc func(this js.Ref, counts *ArcAppKillsDict) js.Ref

func (fn GetArcAppKillsCallbackFunc) Register() js.Func[func(counts *ArcAppKillsDict)] {
	return js.RegisterCallback[func(counts *ArcAppKillsDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetArcAppKillsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppKillsDict
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

type GetArcAppKillsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, counts *ArcAppKillsDict) js.Ref
	Arg T
}

func (cb *GetArcAppKillsCallback[T]) Register() js.Func[func(counts *ArcAppKillsDict)] {
	return js.RegisterCallback[func(counts *ArcAppKillsDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetArcAppKillsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcAppKillsDict
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

type GetArcPackageCallbackFunc func(this js.Ref, pkg *ArcPackageDict) js.Ref

func (fn GetArcPackageCallbackFunc) Register() js.Func[func(pkg *ArcPackageDict)] {
	return js.RegisterCallback[func(pkg *ArcPackageDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetArcPackageCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcPackageDict
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

type GetArcPackageCallback[T any] struct {
	Fn  func(arg T, this js.Ref, pkg *ArcPackageDict) js.Ref
	Arg T
}

func (cb *GetArcPackageCallback[T]) Register() js.Func[func(pkg *ArcPackageDict)] {
	return js.RegisterCallback[func(pkg *ArcPackageDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetArcPackageCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ArcPackageDict
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

type GetCryptohomeRecoveryDataCallbackFunc func(this js.Ref, dict *CryptohomeRecoveryDataDict) js.Ref

func (fn GetCryptohomeRecoveryDataCallbackFunc) Register() js.Func[func(dict *CryptohomeRecoveryDataDict)] {
	return js.RegisterCallback[func(dict *CryptohomeRecoveryDataDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCryptohomeRecoveryDataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CryptohomeRecoveryDataDict
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

type GetCryptohomeRecoveryDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dict *CryptohomeRecoveryDataDict) js.Ref
	Arg T
}

func (cb *GetCryptohomeRecoveryDataCallback[T]) Register() js.Func[func(dict *CryptohomeRecoveryDataDict)] {
	return js.RegisterCallback[func(dict *CryptohomeRecoveryDataDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCryptohomeRecoveryDataCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CryptohomeRecoveryDataDict
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

type GetCurrentInputMethodDescriptorCallbackFunc func(this js.Ref, data *GetCurrentInputMethodDescriptorData) js.Ref

func (fn GetCurrentInputMethodDescriptorCallbackFunc) Register() js.Func[func(data *GetCurrentInputMethodDescriptorData)] {
	return js.RegisterCallback[func(data *GetCurrentInputMethodDescriptorData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCurrentInputMethodDescriptorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetCurrentInputMethodDescriptorData
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

type GetCurrentInputMethodDescriptorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data *GetCurrentInputMethodDescriptorData) js.Ref
	Arg T
}

func (cb *GetCurrentInputMethodDescriptorCallback[T]) Register() js.Func[func(data *GetCurrentInputMethodDescriptorData)] {
	return js.RegisterCallback[func(data *GetCurrentInputMethodDescriptorData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCurrentInputMethodDescriptorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetCurrentInputMethodDescriptorData
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

type GetCurrentInputMethodDescriptorData struct {
	// KeyboardLayout is "GetCurrentInputMethodDescriptorData.keyboardLayout"
	//
	// Optional
	KeyboardLayout js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetCurrentInputMethodDescriptorData with all fields set.
func (p GetCurrentInputMethodDescriptorData) FromRef(ref js.Ref) GetCurrentInputMethodDescriptorData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetCurrentInputMethodDescriptorData in the application heap.
func (p GetCurrentInputMethodDescriptorData) New() js.Ref {
	return bindings.GetCurrentInputMethodDescriptorDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetCurrentInputMethodDescriptorData) UpdateFrom(ref js.Ref) {
	bindings.GetCurrentInputMethodDescriptorDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetCurrentInputMethodDescriptorData) Update(ref js.Ref) {
	bindings.GetCurrentInputMethodDescriptorDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetCurrentInputMethodDescriptorData) FreeMembers(recursive bool) {
	js.Free(
		p.KeyboardLayout.Ref(),
	)
	p.KeyboardLayout = p.KeyboardLayout.FromRef(js.Undefined)
}

type GetDefaultPinnedAppIdsCallbackFunc func(this js.Ref, items js.Array[js.String]) js.Ref

func (fn GetDefaultPinnedAppIdsCallbackFunc) Register() js.Func[func(items js.Array[js.String])] {
	return js.RegisterCallback[func(items js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDefaultPinnedAppIdsCallbackFunc) DispatchCallback(
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

type GetDefaultPinnedAppIdsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, items js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetDefaultPinnedAppIdsCallback[T]) Register() js.Func[func(items js.Array[js.String])] {
	return js.RegisterCallback[func(items js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDefaultPinnedAppIdsCallback[T]) DispatchCallback(
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

type GetDeskCountCallbackFunc func(this js.Ref, count int32) js.Ref

func (fn GetDeskCountCallbackFunc) Register() js.Func[func(count int32)] {
	return js.RegisterCallback[func(count int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeskCountCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeskCountCallback[T any] struct {
	Fn  func(arg T, this js.Ref, count int32) js.Ref
	Arg T
}

func (cb *GetDeskCountCallback[T]) Register() js.Func[func(count int32)] {
	return js.RegisterCallback[func(count int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeskCountCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDesksInfoCallbackFunc func(this js.Ref, desks *DesksInfo) js.Ref

func (fn GetDesksInfoCallbackFunc) Register() js.Func[func(desks *DesksInfo)] {
	return js.RegisterCallback[func(desks *DesksInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDesksInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DesksInfo
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

type GetDesksInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, desks *DesksInfo) js.Ref
	Arg T
}

func (cb *GetDesksInfoCallback[T]) Register() js.Func[func(desks *DesksInfo)] {
	return js.RegisterCallback[func(desks *DesksInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDesksInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DesksInfo
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

type GetDisplaySmoothnessCallbackFunc func(this js.Ref, smoothness int32) js.Ref

func (fn GetDisplaySmoothnessCallbackFunc) Register() js.Func[func(smoothness int32)] {
	return js.RegisterCallback[func(smoothness int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDisplaySmoothnessCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDisplaySmoothnessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, smoothness int32) js.Ref
	Arg T
}

func (cb *GetDisplaySmoothnessCallback[T]) Register() js.Func[func(smoothness int32)] {
	return js.RegisterCallback[func(smoothness int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDisplaySmoothnessCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetLacrosInfoCallbackFunc func(this js.Ref, info *LacrosInfo) js.Ref

func (fn GetLacrosInfoCallbackFunc) Register() js.Func[func(info *LacrosInfo)] {
	return js.RegisterCallback[func(info *LacrosInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLacrosInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LacrosInfo
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

type GetLacrosInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *LacrosInfo) js.Ref
	Arg T
}

func (cb *GetLacrosInfoCallback[T]) Register() js.Func[func(info *LacrosInfo)] {
	return js.RegisterCallback[func(info *LacrosInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLacrosInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LacrosInfo
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

type LacrosState uint32

const (
	_ LacrosState = iota

	LacrosState_NOT_INITIALIZED
	LacrosState_RELOADING
	LacrosState_MOUNTING
	LacrosState_UNAVAILABLE
	LacrosState_STOPPED
	LacrosState_CREATING_LOG_FILE
	LacrosState_PRE_LAUNCHED
	LacrosState_STARTING
	LacrosState_RUNNING
	LacrosState_TERMINATING
)

func (LacrosState) FromRef(str js.Ref) LacrosState {
	return LacrosState(bindings.ConstOfLacrosState(str))
}

func (x LacrosState) String() (string, bool) {
	switch x {
	case LacrosState_NOT_INITIALIZED:
		return "NotInitialized", true
	case LacrosState_RELOADING:
		return "Reloading", true
	case LacrosState_MOUNTING:
		return "Mounting", true
	case LacrosState_UNAVAILABLE:
		return "Unavailable", true
	case LacrosState_STOPPED:
		return "Stopped", true
	case LacrosState_CREATING_LOG_FILE:
		return "CreatingLogFile", true
	case LacrosState_PRE_LAUNCHED:
		return "PreLaunched", true
	case LacrosState_STARTING:
		return "Starting", true
	case LacrosState_RUNNING:
		return "Running", true
	case LacrosState_TERMINATING:
		return "Terminating", true
	default:
		return "", false
	}
}

type LacrosMode uint32

const (
	_ LacrosMode = iota

	LacrosMode_DISABLED
	LacrosMode_ONLY
)

func (LacrosMode) FromRef(str js.Ref) LacrosMode {
	return LacrosMode(bindings.ConstOfLacrosMode(str))
}

func (x LacrosMode) String() (string, bool) {
	switch x {
	case LacrosMode_DISABLED:
		return "Disabled", true
	case LacrosMode_ONLY:
		return "Only", true
	default:
		return "", false
	}
}

type LacrosInfo struct {
	// State is "LacrosInfo.state"
	//
	// Optional
	State LacrosState
	// IsKeepAlive is "LacrosInfo.isKeepAlive"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsKeepAlive MUST be set to true to make this field effective.
	IsKeepAlive bool
	// LacrosPath is "LacrosInfo.lacrosPath"
	//
	// Optional
	LacrosPath js.String
	// Mode is "LacrosInfo.mode"
	//
	// Optional
	Mode LacrosMode

	FFI_USE_IsKeepAlive bool // for IsKeepAlive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LacrosInfo with all fields set.
func (p LacrosInfo) FromRef(ref js.Ref) LacrosInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LacrosInfo in the application heap.
func (p LacrosInfo) New() js.Ref {
	return bindings.LacrosInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LacrosInfo) UpdateFrom(ref js.Ref) {
	bindings.LacrosInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LacrosInfo) Update(ref js.Ref) {
	bindings.LacrosInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LacrosInfo) FreeMembers(recursive bool) {
	js.Free(
		p.LacrosPath.Ref(),
	)
	p.LacrosPath = p.LacrosPath.FromRef(js.Undefined)
}

type GetLauncherSearchBoxStateCallbackFunc func(this js.Ref, state *LauncherSearchBoxState) js.Ref

func (fn GetLauncherSearchBoxStateCallbackFunc) Register() js.Func[func(state *LauncherSearchBoxState)] {
	return js.RegisterCallback[func(state *LauncherSearchBoxState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLauncherSearchBoxStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LauncherSearchBoxState
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

type GetLauncherSearchBoxStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state *LauncherSearchBoxState) js.Ref
	Arg T
}

func (cb *GetLauncherSearchBoxStateCallback[T]) Register() js.Func[func(state *LauncherSearchBoxState)] {
	return js.RegisterCallback[func(state *LauncherSearchBoxState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLauncherSearchBoxStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LauncherSearchBoxState
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

type LauncherSearchBoxState struct {
	// GhostText is "LauncherSearchBoxState.ghostText"
	//
	// Optional
	GhostText js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LauncherSearchBoxState with all fields set.
func (p LauncherSearchBoxState) FromRef(ref js.Ref) LauncherSearchBoxState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LauncherSearchBoxState in the application heap.
func (p LauncherSearchBoxState) New() js.Ref {
	return bindings.LauncherSearchBoxStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LauncherSearchBoxState) UpdateFrom(ref js.Ref) {
	bindings.LauncherSearchBoxStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LauncherSearchBoxState) Update(ref js.Ref) {
	bindings.LauncherSearchBoxStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LauncherSearchBoxState) FreeMembers(recursive bool) {
	js.Free(
		p.GhostText.Ref(),
	)
	p.GhostText = p.GhostText.FromRef(js.Undefined)
}

type GetLoginEventRecorderLoginEventsCallbackFunc func(this js.Ref, data js.Array[LoginEventRecorderData]) js.Ref

func (fn GetLoginEventRecorderLoginEventsCallbackFunc) Register() js.Func[func(data js.Array[LoginEventRecorderData])] {
	return js.RegisterCallback[func(data js.Array[LoginEventRecorderData])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetLoginEventRecorderLoginEventsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[LoginEventRecorderData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetLoginEventRecorderLoginEventsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.Array[LoginEventRecorderData]) js.Ref
	Arg T
}

func (cb *GetLoginEventRecorderLoginEventsCallback[T]) Register() js.Func[func(data js.Array[LoginEventRecorderData])] {
	return js.RegisterCallback[func(data js.Array[LoginEventRecorderData])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetLoginEventRecorderLoginEventsCallback[T]) DispatchCallback(
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

		js.Array[LoginEventRecorderData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoginEventRecorderData struct {
	// Name is "LoginEventRecorderData.name"
	//
	// Optional
	Name js.String
	// MicrosecnodsSinceUnixEpoch is "LoginEventRecorderData.microsecnods_since_unix_epoch"
	//
	// Optional
	//
	// NOTE: FFI_USE_MicrosecnodsSinceUnixEpoch MUST be set to true to make this field effective.
	MicrosecnodsSinceUnixEpoch int32

	FFI_USE_MicrosecnodsSinceUnixEpoch bool // for MicrosecnodsSinceUnixEpoch.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoginEventRecorderData with all fields set.
func (p LoginEventRecorderData) FromRef(ref js.Ref) LoginEventRecorderData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoginEventRecorderData in the application heap.
func (p LoginEventRecorderData) New() js.Ref {
	return bindings.LoginEventRecorderDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoginEventRecorderData) UpdateFrom(ref js.Ref) {
	bindings.LoginEventRecorderDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoginEventRecorderData) Update(ref js.Ref) {
	bindings.LoginEventRecorderDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoginEventRecorderData) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type GetPrimaryDisplayScaleFactorCallbackFunc func(this js.Ref, scaleFactor float64) js.Ref

func (fn GetPrimaryDisplayScaleFactorCallbackFunc) Register() js.Func[func(scaleFactor float64)] {
	return js.RegisterCallback[func(scaleFactor float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPrimaryDisplayScaleFactorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPrimaryDisplayScaleFactorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, scaleFactor float64) js.Ref
	Arg T
}

func (cb *GetPrimaryDisplayScaleFactorCallback[T]) Register() js.Func[func(scaleFactor float64)] {
	return js.RegisterCallback[func(scaleFactor float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPrimaryDisplayScaleFactorCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRegisteredSystemWebAppsCallbackFunc func(this js.Ref, systemWebApps js.Array[SystemWebApp]) js.Ref

func (fn GetRegisteredSystemWebAppsCallbackFunc) Register() js.Func[func(systemWebApps js.Array[SystemWebApp])] {
	return js.RegisterCallback[func(systemWebApps js.Array[SystemWebApp])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRegisteredSystemWebAppsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SystemWebApp]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRegisteredSystemWebAppsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, systemWebApps js.Array[SystemWebApp]) js.Ref
	Arg T
}

func (cb *GetRegisteredSystemWebAppsCallback[T]) Register() js.Func[func(systemWebApps js.Array[SystemWebApp])] {
	return js.RegisterCallback[func(systemWebApps js.Array[SystemWebApp])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRegisteredSystemWebAppsCallback[T]) DispatchCallback(
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

		js.Array[SystemWebApp]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SystemWebApp struct {
	// InternalName is "SystemWebApp.internalName"
	//
	// Optional
	InternalName js.String
	// Url is "SystemWebApp.url"
	//
	// Optional
	Url js.String
	// Name is "SystemWebApp.name"
	//
	// Optional
	Name js.String
	// StartUrl is "SystemWebApp.startUrl"
	//
	// Optional
	StartUrl js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SystemWebApp with all fields set.
func (p SystemWebApp) FromRef(ref js.Ref) SystemWebApp {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SystemWebApp in the application heap.
func (p SystemWebApp) New() js.Ref {
	return bindings.SystemWebAppJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SystemWebApp) UpdateFrom(ref js.Ref) {
	bindings.SystemWebAppJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SystemWebApp) Update(ref js.Ref) {
	bindings.SystemWebAppJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SystemWebApp) FreeMembers(recursive bool) {
	js.Free(
		p.InternalName.Ref(),
		p.Url.Ref(),
		p.Name.Ref(),
		p.StartUrl.Ref(),
	)
	p.InternalName = p.InternalName.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.StartUrl = p.StartUrl.FromRef(js.Undefined)
}

type GetScrollableShelfInfoForStateCallbackFunc func(this js.Ref, info *ScrollableShelfInfo) js.Ref

func (fn GetScrollableShelfInfoForStateCallbackFunc) Register() js.Func[func(info *ScrollableShelfInfo)] {
	return js.RegisterCallback[func(info *ScrollableShelfInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetScrollableShelfInfoForStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScrollableShelfInfo
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

type GetScrollableShelfInfoForStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ScrollableShelfInfo) js.Ref
	Arg T
}

func (cb *GetScrollableShelfInfoForStateCallback[T]) Register() js.Func[func(info *ScrollableShelfInfo)] {
	return js.RegisterCallback[func(info *ScrollableShelfInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetScrollableShelfInfoForStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScrollableShelfInfo
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

type ScrollableShelfInfo struct {
	// MainAxisOffset is "ScrollableShelfInfo.mainAxisOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_MainAxisOffset MUST be set to true to make this field effective.
	MainAxisOffset float64
	// PageOffset is "ScrollableShelfInfo.pageOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_PageOffset MUST be set to true to make this field effective.
	PageOffset float64
	// TargetMainAxisOffset is "ScrollableShelfInfo.targetMainAxisOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetMainAxisOffset MUST be set to true to make this field effective.
	TargetMainAxisOffset float64
	// LeftArrowBounds is "ScrollableShelfInfo.leftArrowBounds"
	//
	// Optional
	//
	// NOTE: LeftArrowBounds.FFI_USE MUST be set to true to get LeftArrowBounds used.
	LeftArrowBounds Bounds
	// RightArrowBounds is "ScrollableShelfInfo.rightArrowBounds"
	//
	// Optional
	//
	// NOTE: RightArrowBounds.FFI_USE MUST be set to true to get RightArrowBounds used.
	RightArrowBounds Bounds
	// IsAnimating is "ScrollableShelfInfo.isAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAnimating MUST be set to true to make this field effective.
	IsAnimating bool
	// IconsUnderAnimation is "ScrollableShelfInfo.iconsUnderAnimation"
	//
	// Optional
	//
	// NOTE: FFI_USE_IconsUnderAnimation MUST be set to true to make this field effective.
	IconsUnderAnimation bool
	// IsOverflow is "ScrollableShelfInfo.isOverflow"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsOverflow MUST be set to true to make this field effective.
	IsOverflow bool
	// IconsBoundsInScreen is "ScrollableShelfInfo.iconsBoundsInScreen"
	//
	// Optional
	IconsBoundsInScreen js.Array[Bounds]
	// IsShelfWidgetAnimating is "ScrollableShelfInfo.isShelfWidgetAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsShelfWidgetAnimating MUST be set to true to make this field effective.
	IsShelfWidgetAnimating bool

	FFI_USE_MainAxisOffset         bool // for MainAxisOffset.
	FFI_USE_PageOffset             bool // for PageOffset.
	FFI_USE_TargetMainAxisOffset   bool // for TargetMainAxisOffset.
	FFI_USE_IsAnimating            bool // for IsAnimating.
	FFI_USE_IconsUnderAnimation    bool // for IconsUnderAnimation.
	FFI_USE_IsOverflow             bool // for IsOverflow.
	FFI_USE_IsShelfWidgetAnimating bool // for IsShelfWidgetAnimating.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollableShelfInfo with all fields set.
func (p ScrollableShelfInfo) FromRef(ref js.Ref) ScrollableShelfInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollableShelfInfo in the application heap.
func (p ScrollableShelfInfo) New() js.Ref {
	return bindings.ScrollableShelfInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollableShelfInfo) UpdateFrom(ref js.Ref) {
	bindings.ScrollableShelfInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollableShelfInfo) Update(ref js.Ref) {
	bindings.ScrollableShelfInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollableShelfInfo) FreeMembers(recursive bool) {
	js.Free(
		p.IconsBoundsInScreen.Ref(),
	)
	p.IconsBoundsInScreen = p.IconsBoundsInScreen.FromRef(js.Undefined)
	if recursive {
		p.LeftArrowBounds.FreeMembers(true)
		p.RightArrowBounds.FreeMembers(true)
	}
}

type GetShelfAlignmentCallbackFunc func(this js.Ref, alignment ShelfAlignmentType) js.Ref

func (fn GetShelfAlignmentCallbackFunc) Register() js.Func[func(alignment ShelfAlignmentType)] {
	return js.RegisterCallback[func(alignment ShelfAlignmentType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetShelfAlignmentCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ShelfAlignmentType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetShelfAlignmentCallback[T any] struct {
	Fn  func(arg T, this js.Ref, alignment ShelfAlignmentType) js.Ref
	Arg T
}

func (cb *GetShelfAlignmentCallback[T]) Register() js.Func[func(alignment ShelfAlignmentType)] {
	return js.RegisterCallback[func(alignment ShelfAlignmentType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetShelfAlignmentCallback[T]) DispatchCallback(
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

		ShelfAlignmentType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ShelfAlignmentType uint32

const (
	_ ShelfAlignmentType = iota

	ShelfAlignmentType_BOTTOM
	ShelfAlignmentType_LEFT
	ShelfAlignmentType_RIGHT
)

func (ShelfAlignmentType) FromRef(str js.Ref) ShelfAlignmentType {
	return ShelfAlignmentType(bindings.ConstOfShelfAlignmentType(str))
}

func (x ShelfAlignmentType) String() (string, bool) {
	switch x {
	case ShelfAlignmentType_BOTTOM:
		return "Bottom", true
	case ShelfAlignmentType_LEFT:
		return "Left", true
	case ShelfAlignmentType_RIGHT:
		return "Right", true
	default:
		return "", false
	}
}

type GetShelfAutoHideBehaviorCallbackFunc func(this js.Ref, behavior js.String) js.Ref

func (fn GetShelfAutoHideBehaviorCallbackFunc) Register() js.Func[func(behavior js.String)] {
	return js.RegisterCallback[func(behavior js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetShelfAutoHideBehaviorCallbackFunc) DispatchCallback(
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

type GetShelfAutoHideBehaviorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, behavior js.String) js.Ref
	Arg T
}

func (cb *GetShelfAutoHideBehaviorCallback[T]) Register() js.Func[func(behavior js.String)] {
	return js.RegisterCallback[func(behavior js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetShelfAutoHideBehaviorCallback[T]) DispatchCallback(
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

type GetShelfItemsCallbackFunc func(this js.Ref, items js.Array[ShelfItem]) js.Ref

func (fn GetShelfItemsCallbackFunc) Register() js.Func[func(items js.Array[ShelfItem])] {
	return js.RegisterCallback[func(items js.Array[ShelfItem])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetShelfItemsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ShelfItem]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetShelfItemsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, items js.Array[ShelfItem]) js.Ref
	Arg T
}

func (cb *GetShelfItemsCallback[T]) Register() js.Func[func(items js.Array[ShelfItem])] {
	return js.RegisterCallback[func(items js.Array[ShelfItem])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetShelfItemsCallback[T]) DispatchCallback(
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

		js.Array[ShelfItem]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ShelfItemType uint32

const (
	_ ShelfItemType = iota

	ShelfItemType_PINNED_APP
	ShelfItemType_BROWSER_SHORTCUT
	ShelfItemType_APP
	ShelfItemType_UNPINNED_BROWSER_SHORTCUT
	ShelfItemType_DIALOG
)

func (ShelfItemType) FromRef(str js.Ref) ShelfItemType {
	return ShelfItemType(bindings.ConstOfShelfItemType(str))
}

func (x ShelfItemType) String() (string, bool) {
	switch x {
	case ShelfItemType_PINNED_APP:
		return "PinnedApp", true
	case ShelfItemType_BROWSER_SHORTCUT:
		return "BrowserShortcut", true
	case ShelfItemType_APP:
		return "App", true
	case ShelfItemType_UNPINNED_BROWSER_SHORTCUT:
		return "UnpinnedBrowserShortcut", true
	case ShelfItemType_DIALOG:
		return "Dialog", true
	default:
		return "", false
	}
}

type ShelfItemStatus uint32

const (
	_ ShelfItemStatus = iota

	ShelfItemStatus_CLOSED
	ShelfItemStatus_RUNNING
	ShelfItemStatus_ATTENTION
)

func (ShelfItemStatus) FromRef(str js.Ref) ShelfItemStatus {
	return ShelfItemStatus(bindings.ConstOfShelfItemStatus(str))
}

func (x ShelfItemStatus) String() (string, bool) {
	switch x {
	case ShelfItemStatus_CLOSED:
		return "Closed", true
	case ShelfItemStatus_RUNNING:
		return "Running", true
	case ShelfItemStatus_ATTENTION:
		return "Attention", true
	default:
		return "", false
	}
}

type ShelfItem struct {
	// AppId is "ShelfItem.appId"
	//
	// Optional
	AppId js.String
	// LaunchId is "ShelfItem.launchId"
	//
	// Optional
	LaunchId js.String
	// Title is "ShelfItem.title"
	//
	// Optional
	Title js.String
	// Type is "ShelfItem.type"
	//
	// Optional
	Type ShelfItemType
	// Status is "ShelfItem.status"
	//
	// Optional
	Status ShelfItemStatus
	// ShowsTooltip is "ShelfItem.showsTooltip"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowsTooltip MUST be set to true to make this field effective.
	ShowsTooltip bool
	// PinnedByPolicy is "ShelfItem.pinnedByPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinnedByPolicy MUST be set to true to make this field effective.
	PinnedByPolicy bool
	// PinStateForcedByType is "ShelfItem.pinStateForcedByType"
	//
	// Optional
	//
	// NOTE: FFI_USE_PinStateForcedByType MUST be set to true to make this field effective.
	PinStateForcedByType bool
	// HasNotification is "ShelfItem.hasNotification"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasNotification MUST be set to true to make this field effective.
	HasNotification bool

	FFI_USE_ShowsTooltip         bool // for ShowsTooltip.
	FFI_USE_PinnedByPolicy       bool // for PinnedByPolicy.
	FFI_USE_PinStateForcedByType bool // for PinStateForcedByType.
	FFI_USE_HasNotification      bool // for HasNotification.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShelfItem with all fields set.
func (p ShelfItem) FromRef(ref js.Ref) ShelfItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShelfItem in the application heap.
func (p ShelfItem) New() js.Ref {
	return bindings.ShelfItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShelfItem) UpdateFrom(ref js.Ref) {
	bindings.ShelfItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShelfItem) Update(ref js.Ref) {
	bindings.ShelfItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShelfItem) FreeMembers(recursive bool) {
	js.Free(
		p.AppId.Ref(),
		p.LaunchId.Ref(),
		p.Title.Ref(),
	)
	p.AppId = p.AppId.FromRef(js.Undefined)
	p.LaunchId = p.LaunchId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type GetShelfUIInfoForStateCallbackFunc func(this js.Ref, info *ShelfUIInfo) js.Ref

func (fn GetShelfUIInfoForStateCallbackFunc) Register() js.Func[func(info *ShelfUIInfo)] {
	return js.RegisterCallback[func(info *ShelfUIInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetShelfUIInfoForStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ShelfUIInfo
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

type GetShelfUIInfoForStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ShelfUIInfo) js.Ref
	Arg T
}

func (cb *GetShelfUIInfoForStateCallback[T]) Register() js.Func[func(info *ShelfUIInfo)] {
	return js.RegisterCallback[func(info *ShelfUIInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetShelfUIInfoForStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ShelfUIInfo
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

type Location struct {
	// X is "Location.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "Location.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Location with all fields set.
func (p Location) FromRef(ref js.Ref) Location {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Location in the application heap.
func (p Location) New() js.Ref {
	return bindings.LocationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Location) UpdateFrom(ref js.Ref) {
	bindings.LocationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Location) Update(ref js.Ref) {
	bindings.LocationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Location) FreeMembers(recursive bool) {
}

type HotseatSwipeDescriptor struct {
	// SwipeStartLocation is "HotseatSwipeDescriptor.swipeStartLocation"
	//
	// Optional
	//
	// NOTE: SwipeStartLocation.FFI_USE MUST be set to true to get SwipeStartLocation used.
	SwipeStartLocation Location
	// SwipeEndLocation is "HotseatSwipeDescriptor.swipeEndLocation"
	//
	// Optional
	//
	// NOTE: SwipeEndLocation.FFI_USE MUST be set to true to get SwipeEndLocation used.
	SwipeEndLocation Location

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HotseatSwipeDescriptor with all fields set.
func (p HotseatSwipeDescriptor) FromRef(ref js.Ref) HotseatSwipeDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HotseatSwipeDescriptor in the application heap.
func (p HotseatSwipeDescriptor) New() js.Ref {
	return bindings.HotseatSwipeDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HotseatSwipeDescriptor) UpdateFrom(ref js.Ref) {
	bindings.HotseatSwipeDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HotseatSwipeDescriptor) Update(ref js.Ref) {
	bindings.HotseatSwipeDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HotseatSwipeDescriptor) FreeMembers(recursive bool) {
	if recursive {
		p.SwipeStartLocation.FreeMembers(true)
		p.SwipeEndLocation.FreeMembers(true)
	}
}

type HotseatState uint32

const (
	_ HotseatState = iota

	HotseatState_HIDDEN
	HotseatState_SHOWN_CLAM_SHELL
	HotseatState_SHOWN_HOME_LAUNCHER
	HotseatState_EXTENDED
)

func (HotseatState) FromRef(str js.Ref) HotseatState {
	return HotseatState(bindings.ConstOfHotseatState(str))
}

func (x HotseatState) String() (string, bool) {
	switch x {
	case HotseatState_HIDDEN:
		return "Hidden", true
	case HotseatState_SHOWN_CLAM_SHELL:
		return "ShownClamShell", true
	case HotseatState_SHOWN_HOME_LAUNCHER:
		return "ShownHomeLauncher", true
	case HotseatState_EXTENDED:
		return "Extended", true
	default:
		return "", false
	}
}

type HotseatInfo struct {
	// SwipeUp is "HotseatInfo.swipeUp"
	//
	// Optional
	//
	// NOTE: SwipeUp.FFI_USE MUST be set to true to get SwipeUp used.
	SwipeUp HotseatSwipeDescriptor
	// State is "HotseatInfo.state"
	//
	// Optional
	State HotseatState
	// IsAnimating is "HotseatInfo.isAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAnimating MUST be set to true to make this field effective.
	IsAnimating bool
	// IsAutoHidden is "HotseatInfo.isAutoHidden"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAutoHidden MUST be set to true to make this field effective.
	IsAutoHidden bool

	FFI_USE_IsAnimating  bool // for IsAnimating.
	FFI_USE_IsAutoHidden bool // for IsAutoHidden.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HotseatInfo with all fields set.
func (p HotseatInfo) FromRef(ref js.Ref) HotseatInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HotseatInfo in the application heap.
func (p HotseatInfo) New() js.Ref {
	return bindings.HotseatInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HotseatInfo) UpdateFrom(ref js.Ref) {
	bindings.HotseatInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HotseatInfo) Update(ref js.Ref) {
	bindings.HotseatInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HotseatInfo) FreeMembers(recursive bool) {
	if recursive {
		p.SwipeUp.FreeMembers(true)
	}
}

type ShelfUIInfo struct {
	// HotseatInfo is "ShelfUIInfo.hotseatInfo"
	//
	// Optional
	//
	// NOTE: HotseatInfo.FFI_USE MUST be set to true to get HotseatInfo used.
	HotseatInfo HotseatInfo
	// ScrollableShelfInfo is "ShelfUIInfo.scrollableShelfInfo"
	//
	// Optional
	//
	// NOTE: ScrollableShelfInfo.FFI_USE MUST be set to true to get ScrollableShelfInfo used.
	ScrollableShelfInfo ScrollableShelfInfo

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShelfUIInfo with all fields set.
func (p ShelfUIInfo) FromRef(ref js.Ref) ShelfUIInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShelfUIInfo in the application heap.
func (p ShelfUIInfo) New() js.Ref {
	return bindings.ShelfUIInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShelfUIInfo) UpdateFrom(ref js.Ref) {
	bindings.ShelfUIInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShelfUIInfo) Update(ref js.Ref) {
	bindings.ShelfUIInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShelfUIInfo) FreeMembers(recursive bool) {
	if recursive {
		p.HotseatInfo.FreeMembers(true)
		p.ScrollableShelfInfo.FreeMembers(true)
	}
}

type GetThroughtputTrackerDataCallbackFunc func(this js.Ref, data js.Array[ThroughputTrackerAnimationData]) js.Ref

func (fn GetThroughtputTrackerDataCallbackFunc) Register() js.Func[func(data js.Array[ThroughputTrackerAnimationData])] {
	return js.RegisterCallback[func(data js.Array[ThroughputTrackerAnimationData])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetThroughtputTrackerDataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ThroughputTrackerAnimationData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetThroughtputTrackerDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.Array[ThroughputTrackerAnimationData]) js.Ref
	Arg T
}

func (cb *GetThroughtputTrackerDataCallback[T]) Register() js.Func[func(data js.Array[ThroughputTrackerAnimationData])] {
	return js.RegisterCallback[func(data js.Array[ThroughputTrackerAnimationData])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetThroughtputTrackerDataCallback[T]) DispatchCallback(
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

		js.Array[ThroughputTrackerAnimationData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ThroughputTrackerAnimationData struct {
	// StartOffsetMs is "ThroughputTrackerAnimationData.startOffsetMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartOffsetMs MUST be set to true to make this field effective.
	StartOffsetMs int32
	// StopOffsetMs is "ThroughputTrackerAnimationData.stopOffsetMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_StopOffsetMs MUST be set to true to make this field effective.
	StopOffsetMs int32
	// FramesExpected is "ThroughputTrackerAnimationData.framesExpected"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesExpected MUST be set to true to make this field effective.
	FramesExpected int32
	// FramesProduced is "ThroughputTrackerAnimationData.framesProduced"
	//
	// Optional
	//
	// NOTE: FFI_USE_FramesProduced MUST be set to true to make this field effective.
	FramesProduced int32
	// JankCount is "ThroughputTrackerAnimationData.jankCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_JankCount MUST be set to true to make this field effective.
	JankCount int32

	FFI_USE_StartOffsetMs  bool // for StartOffsetMs.
	FFI_USE_StopOffsetMs   bool // for StopOffsetMs.
	FFI_USE_FramesExpected bool // for FramesExpected.
	FFI_USE_FramesProduced bool // for FramesProduced.
	FFI_USE_JankCount      bool // for JankCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ThroughputTrackerAnimationData with all fields set.
func (p ThroughputTrackerAnimationData) FromRef(ref js.Ref) ThroughputTrackerAnimationData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ThroughputTrackerAnimationData in the application heap.
func (p ThroughputTrackerAnimationData) New() js.Ref {
	return bindings.ThroughputTrackerAnimationDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ThroughputTrackerAnimationData) UpdateFrom(ref js.Ref) {
	bindings.ThroughputTrackerAnimationDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ThroughputTrackerAnimationData) Update(ref js.Ref) {
	bindings.ThroughputTrackerAnimationDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ThroughputTrackerAnimationData) FreeMembers(recursive bool) {
}

type InstallPWAForCurrentURLCallbackFunc func(this js.Ref, appId js.String) js.Ref

func (fn InstallPWAForCurrentURLCallbackFunc) Register() js.Func[func(appId js.String)] {
	return js.RegisterCallback[func(appId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InstallPWAForCurrentURLCallbackFunc) DispatchCallback(
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

type InstallPWAForCurrentURLCallback[T any] struct {
	Fn  func(arg T, this js.Ref, appId js.String) js.Ref
	Arg T
}

func (cb *InstallPWAForCurrentURLCallback[T]) Register() js.Func[func(appId js.String)] {
	return js.RegisterCallback[func(appId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InstallPWAForCurrentURLCallback[T]) DispatchCallback(
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

type IsAppShownCallbackFunc func(this js.Ref, appShown bool) js.Ref

func (fn IsAppShownCallbackFunc) Register() js.Func[func(appShown bool)] {
	return js.RegisterCallback[func(appShown bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsAppShownCallbackFunc) DispatchCallback(
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

type IsAppShownCallback[T any] struct {
	Fn  func(arg T, this js.Ref, appShown bool) js.Ref
	Arg T
}

func (cb *IsAppShownCallback[T]) Register() js.Func[func(appShown bool)] {
	return js.RegisterCallback[func(appShown bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsAppShownCallback[T]) DispatchCallback(
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

type IsArcPackageListInitialRefreshedCallbackFunc func(this js.Ref, refreshed bool) js.Ref

func (fn IsArcPackageListInitialRefreshedCallbackFunc) Register() js.Func[func(refreshed bool)] {
	return js.RegisterCallback[func(refreshed bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsArcPackageListInitialRefreshedCallbackFunc) DispatchCallback(
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

type IsArcPackageListInitialRefreshedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, refreshed bool) js.Ref
	Arg T
}

func (cb *IsArcPackageListInitialRefreshedCallback[T]) Register() js.Func[func(refreshed bool)] {
	return js.RegisterCallback[func(refreshed bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsArcPackageListInitialRefreshedCallback[T]) DispatchCallback(
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

type IsArcProvisionedCallbackFunc func(this js.Ref, arcProvisioned bool) js.Ref

func (fn IsArcProvisionedCallbackFunc) Register() js.Func[func(arcProvisioned bool)] {
	return js.RegisterCallback[func(arcProvisioned bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsArcProvisionedCallbackFunc) DispatchCallback(
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

type IsArcProvisionedCallback[T any] struct {
	Fn  func(arg T, this js.Ref, arcProvisioned bool) js.Ref
	Arg T
}

func (cb *IsArcProvisionedCallback[T]) Register() js.Func[func(arcProvisioned bool)] {
	return js.RegisterCallback[func(arcProvisioned bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsArcProvisionedCallback[T]) DispatchCallback(
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

type IsFeatureEnabledCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn IsFeatureEnabledCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsFeatureEnabledCallbackFunc) DispatchCallback(
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

type IsFeatureEnabledCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *IsFeatureEnabledCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsFeatureEnabledCallback[T]) DispatchCallback(
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

type IsInputMethodReadyForTestingCallbackFunc func(this js.Ref, isReady bool) js.Ref

func (fn IsInputMethodReadyForTestingCallbackFunc) Register() js.Func[func(isReady bool)] {
	return js.RegisterCallback[func(isReady bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsInputMethodReadyForTestingCallbackFunc) DispatchCallback(
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

type IsInputMethodReadyForTestingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isReady bool) js.Ref
	Arg T
}

func (cb *IsInputMethodReadyForTestingCallback[T]) Register() js.Func[func(isReady bool)] {
	return js.RegisterCallback[func(isReady bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsInputMethodReadyForTestingCallback[T]) DispatchCallback(
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

type IsSystemWebAppOpenCallbackFunc func(this js.Ref, isOpen bool) js.Ref

func (fn IsSystemWebAppOpenCallbackFunc) Register() js.Func[func(isOpen bool)] {
	return js.RegisterCallback[func(isOpen bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsSystemWebAppOpenCallbackFunc) DispatchCallback(
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

type IsSystemWebAppOpenCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isOpen bool) js.Ref
	Arg T
}

func (cb *IsSystemWebAppOpenCallback[T]) Register() js.Func[func(isOpen bool)] {
	return js.RegisterCallback[func(isOpen bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsSystemWebAppOpenCallback[T]) DispatchCallback(
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

type IsTabletModeEnabledCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn IsTabletModeEnabledCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsTabletModeEnabledCallbackFunc) DispatchCallback(
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

type IsTabletModeEnabledCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *IsTabletModeEnabledCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsTabletModeEnabledCallback[T]) DispatchCallback(
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

type LauncherStateType uint32

const (
	_ LauncherStateType = iota

	LauncherStateType_CLOSED
	LauncherStateType_FULLSCREEN_ALL_APPS
	LauncherStateType_FULLSCREEN_SEARCH
)

func (LauncherStateType) FromRef(str js.Ref) LauncherStateType {
	return LauncherStateType(bindings.ConstOfLauncherStateType(str))
}

func (x LauncherStateType) String() (string, bool) {
	switch x {
	case LauncherStateType_CLOSED:
		return "Closed", true
	case LauncherStateType_FULLSCREEN_ALL_APPS:
		return "FullscreenAllApps", true
	case LauncherStateType_FULLSCREEN_SEARCH:
		return "FullscreenSearch", true
	default:
		return "", false
	}
}

type LoginStatusCallbackFunc func(this js.Ref, status *LoginStatusDict) js.Ref

func (fn LoginStatusCallbackFunc) Register() js.Func[func(status *LoginStatusDict)] {
	return js.RegisterCallback[func(status *LoginStatusDict)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoginStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoginStatusDict
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

type LoginStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *LoginStatusDict) js.Ref
	Arg T
}

func (cb *LoginStatusCallback[T]) Register() js.Func[func(status *LoginStatusDict)] {
	return js.RegisterCallback[func(status *LoginStatusDict)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoginStatusCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LoginStatusDict
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

type LoginStatusDict struct {
	// IsLoggedIn is "LoginStatusDict.isLoggedIn"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLoggedIn MUST be set to true to make this field effective.
	IsLoggedIn bool
	// IsOwner is "LoginStatusDict.isOwner"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsOwner MUST be set to true to make this field effective.
	IsOwner bool
	// IsScreenLocked is "LoginStatusDict.isScreenLocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsScreenLocked MUST be set to true to make this field effective.
	IsScreenLocked bool
	// IsLockscreenWallpaperAnimating is "LoginStatusDict.isLockscreenWallpaperAnimating"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLockscreenWallpaperAnimating MUST be set to true to make this field effective.
	IsLockscreenWallpaperAnimating bool
	// IsReadyForPassword is "LoginStatusDict.isReadyForPassword"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsReadyForPassword MUST be set to true to make this field effective.
	IsReadyForPassword bool
	// AreAllUserImagesLoaded is "LoginStatusDict.areAllUserImagesLoaded"
	//
	// Optional
	//
	// NOTE: FFI_USE_AreAllUserImagesLoaded MUST be set to true to make this field effective.
	AreAllUserImagesLoaded bool
	// IsRegularUser is "LoginStatusDict.isRegularUser"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRegularUser MUST be set to true to make this field effective.
	IsRegularUser bool
	// IsGuest is "LoginStatusDict.isGuest"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsGuest MUST be set to true to make this field effective.
	IsGuest bool
	// IsKiosk is "LoginStatusDict.isKiosk"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsKiosk MUST be set to true to make this field effective.
	IsKiosk bool
	// Email is "LoginStatusDict.email"
	//
	// Optional
	Email js.String
	// DisplayEmail is "LoginStatusDict.displayEmail"
	//
	// Optional
	DisplayEmail js.String
	// DisplayName is "LoginStatusDict.displayName"
	//
	// Optional
	DisplayName js.String
	// UserImage is "LoginStatusDict.userImage"
	//
	// Optional
	UserImage js.String
	// HasValidOauth2Token is "LoginStatusDict.hasValidOauth2Token"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasValidOauth2Token MUST be set to true to make this field effective.
	HasValidOauth2Token bool

	FFI_USE_IsLoggedIn                     bool // for IsLoggedIn.
	FFI_USE_IsOwner                        bool // for IsOwner.
	FFI_USE_IsScreenLocked                 bool // for IsScreenLocked.
	FFI_USE_IsLockscreenWallpaperAnimating bool // for IsLockscreenWallpaperAnimating.
	FFI_USE_IsReadyForPassword             bool // for IsReadyForPassword.
	FFI_USE_AreAllUserImagesLoaded         bool // for AreAllUserImagesLoaded.
	FFI_USE_IsRegularUser                  bool // for IsRegularUser.
	FFI_USE_IsGuest                        bool // for IsGuest.
	FFI_USE_IsKiosk                        bool // for IsKiosk.
	FFI_USE_HasValidOauth2Token            bool // for HasValidOauth2Token.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoginStatusDict with all fields set.
func (p LoginStatusDict) FromRef(ref js.Ref) LoginStatusDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoginStatusDict in the application heap.
func (p LoginStatusDict) New() js.Ref {
	return bindings.LoginStatusDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LoginStatusDict) UpdateFrom(ref js.Ref) {
	bindings.LoginStatusDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LoginStatusDict) Update(ref js.Ref) {
	bindings.LoginStatusDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LoginStatusDict) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
		p.DisplayEmail.Ref(),
		p.DisplayName.Ref(),
		p.UserImage.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
	p.DisplayEmail = p.DisplayEmail.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.UserImage = p.UserImage.FromRef(js.Undefined)
}

type MakeFuseboxTempDirCallbackFunc func(this js.Ref, data *MakeFuseboxTempDirData) js.Ref

func (fn MakeFuseboxTempDirCallbackFunc) Register() js.Func[func(data *MakeFuseboxTempDirData)] {
	return js.RegisterCallback[func(data *MakeFuseboxTempDirData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MakeFuseboxTempDirCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MakeFuseboxTempDirData
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

type MakeFuseboxTempDirCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data *MakeFuseboxTempDirData) js.Ref
	Arg T
}

func (cb *MakeFuseboxTempDirCallback[T]) Register() js.Func[func(data *MakeFuseboxTempDirData)] {
	return js.RegisterCallback[func(data *MakeFuseboxTempDirData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MakeFuseboxTempDirCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MakeFuseboxTempDirData
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

type MakeFuseboxTempDirData struct {
	// FuseboxFilePath is "MakeFuseboxTempDirData.fuseboxFilePath"
	//
	// Optional
	FuseboxFilePath js.String
	// UnderlyingFilePath is "MakeFuseboxTempDirData.underlyingFilePath"
	//
	// Optional
	UnderlyingFilePath js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MakeFuseboxTempDirData with all fields set.
func (p MakeFuseboxTempDirData) FromRef(ref js.Ref) MakeFuseboxTempDirData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MakeFuseboxTempDirData in the application heap.
func (p MakeFuseboxTempDirData) New() js.Ref {
	return bindings.MakeFuseboxTempDirDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MakeFuseboxTempDirData) UpdateFrom(ref js.Ref) {
	bindings.MakeFuseboxTempDirDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MakeFuseboxTempDirData) Update(ref js.Ref) {
	bindings.MakeFuseboxTempDirDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MakeFuseboxTempDirData) FreeMembers(recursive bool) {
	js.Free(
		p.FuseboxFilePath.Ref(),
		p.UnderlyingFilePath.Ref(),
	)
	p.FuseboxFilePath = p.FuseboxFilePath.FromRef(js.Undefined)
	p.UnderlyingFilePath = p.UnderlyingFilePath.FromRef(js.Undefined)
}

type MouseButton uint32

const (
	_ MouseButton = iota

	MouseButton_LEFT
	MouseButton_MIDDLE
	MouseButton_RIGHT
	MouseButton_BACK
	MouseButton_FORWARD
)

func (MouseButton) FromRef(str js.Ref) MouseButton {
	return MouseButton(bindings.ConstOfMouseButton(str))
}

func (x MouseButton) String() (string, bool) {
	switch x {
	case MouseButton_LEFT:
		return "Left", true
	case MouseButton_MIDDLE:
		return "Middle", true
	case MouseButton_RIGHT:
		return "Right", true
	case MouseButton_BACK:
		return "Back", true
	case MouseButton_FORWARD:
		return "Forward", true
	default:
		return "", false
	}
}

type Notification struct {
	// Id is "Notification.id"
	//
	// Optional
	Id js.String
	// Type is "Notification.type"
	//
	// Optional
	Type js.String
	// Title is "Notification.title"
	//
	// Optional
	Title js.String
	// Message is "Notification.message"
	//
	// Optional
	Message js.String
	// Priority is "Notification.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// Progress is "Notification.progress"
	//
	// Optional
	//
	// NOTE: FFI_USE_Progress MUST be set to true to make this field effective.
	Progress int32

	FFI_USE_Priority bool // for Priority.
	FFI_USE_Progress bool // for Progress.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Notification with all fields set.
func (p Notification) FromRef(ref js.Ref) Notification {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Notification in the application heap.
func (p Notification) New() js.Ref {
	return bindings.NotificationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Notification) UpdateFrom(ref js.Ref) {
	bindings.NotificationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Notification) Update(ref js.Ref) {
	bindings.NotificationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Notification) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Type.Ref(),
		p.Title.Ref(),
		p.Message.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
}

type NotificationArrayCallbackFunc func(this js.Ref, notifications js.Array[Notification]) js.Ref

func (fn NotificationArrayCallbackFunc) Register() js.Func[func(notifications js.Array[Notification])] {
	return js.RegisterCallback[func(notifications js.Array[Notification])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NotificationArrayCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Notification]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NotificationArrayCallback[T any] struct {
	Fn  func(arg T, this js.Ref, notifications js.Array[Notification]) js.Ref
	Arg T
}

func (cb *NotificationArrayCallback[T]) Register() js.Func[func(notifications js.Array[Notification])] {
	return js.RegisterCallback[func(notifications js.Array[Notification])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NotificationArrayCallback[T]) DispatchCallback(
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

		js.Array[Notification]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OverviewStateType uint32

const (
	_ OverviewStateType = iota

	OverviewStateType_SHOWN
	OverviewStateType_HIDDEN
)

func (OverviewStateType) FromRef(str js.Ref) OverviewStateType {
	return OverviewStateType(bindings.ConstOfOverviewStateType(str))
}

func (x OverviewStateType) String() (string, bool) {
	switch x {
	case OverviewStateType_SHOWN:
		return "Shown", true
	case OverviewStateType_HIDDEN:
		return "Hidden", true
	default:
		return "", false
	}
}

type PlayStoreState struct {
	// Allowed is "PlayStoreState.allowed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Allowed MUST be set to true to make this field effective.
	Allowed bool
	// Enabled is "PlayStoreState.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Managed is "PlayStoreState.managed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Managed MUST be set to true to make this field effective.
	Managed bool

	FFI_USE_Allowed bool // for Allowed.
	FFI_USE_Enabled bool // for Enabled.
	FFI_USE_Managed bool // for Managed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PlayStoreState with all fields set.
func (p PlayStoreState) FromRef(ref js.Ref) PlayStoreState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PlayStoreState in the application heap.
func (p PlayStoreState) New() js.Ref {
	return bindings.PlayStoreStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PlayStoreState) UpdateFrom(ref js.Ref) {
	bindings.PlayStoreStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PlayStoreState) Update(ref js.Ref) {
	bindings.PlayStoreStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PlayStoreState) FreeMembers(recursive bool) {
}

type PlayStoreStateCallbackFunc func(this js.Ref, result *PlayStoreState) js.Ref

func (fn PlayStoreStateCallbackFunc) Register() js.Func[func(result *PlayStoreState)] {
	return js.RegisterCallback[func(result *PlayStoreState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PlayStoreStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PlayStoreState
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

type PlayStoreStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *PlayStoreState) js.Ref
	Arg T
}

func (cb *PlayStoreStateCallback[T]) Register() js.Func[func(result *PlayStoreState)] {
	return js.RegisterCallback[func(result *PlayStoreState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PlayStoreStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PlayStoreState
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

type Printer struct {
	// PrinterName is "Printer.printerName"
	//
	// Optional
	PrinterName js.String
	// PrinterId is "Printer.printerId"
	//
	// Optional
	PrinterId js.String
	// PrinterType is "Printer.printerType"
	//
	// Optional
	PrinterType js.String
	// PrinterDesc is "Printer.printerDesc"
	//
	// Optional
	PrinterDesc js.String
	// PrinterMakeAndModel is "Printer.printerMakeAndModel"
	//
	// Optional
	PrinterMakeAndModel js.String
	// PrinterUri is "Printer.printerUri"
	//
	// Optional
	PrinterUri js.String
	// PrinterPpd is "Printer.printerPpd"
	//
	// Optional
	PrinterPpd js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Printer with all fields set.
func (p Printer) FromRef(ref js.Ref) Printer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Printer in the application heap.
func (p Printer) New() js.Ref {
	return bindings.PrinterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Printer) UpdateFrom(ref js.Ref) {
	bindings.PrinterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Printer) Update(ref js.Ref) {
	bindings.PrinterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Printer) FreeMembers(recursive bool) {
	js.Free(
		p.PrinterName.Ref(),
		p.PrinterId.Ref(),
		p.PrinterType.Ref(),
		p.PrinterDesc.Ref(),
		p.PrinterMakeAndModel.Ref(),
		p.PrinterUri.Ref(),
		p.PrinterPpd.Ref(),
	)
	p.PrinterName = p.PrinterName.FromRef(js.Undefined)
	p.PrinterId = p.PrinterId.FromRef(js.Undefined)
	p.PrinterType = p.PrinterType.FromRef(js.Undefined)
	p.PrinterDesc = p.PrinterDesc.FromRef(js.Undefined)
	p.PrinterMakeAndModel = p.PrinterMakeAndModel.FromRef(js.Undefined)
	p.PrinterUri = p.PrinterUri.FromRef(js.Undefined)
	p.PrinterPpd = p.PrinterPpd.FromRef(js.Undefined)
}

type PrinterArrayCallbackFunc func(this js.Ref, printers js.Array[Printer]) js.Ref

func (fn PrinterArrayCallbackFunc) Register() js.Func[func(printers js.Array[Printer])] {
	return js.RegisterCallback[func(printers js.Array[Printer])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PrinterArrayCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Printer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PrinterArrayCallback[T any] struct {
	Fn  func(arg T, this js.Ref, printers js.Array[Printer]) js.Ref
	Arg T
}

func (cb *PrinterArrayCallback[T]) Register() js.Func[func(printers js.Array[Printer])] {
	return js.RegisterCallback[func(printers js.Array[Printer])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PrinterArrayCallback[T]) DispatchCallback(
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

		js.Array[Printer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemoveFuseboxTempDirCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveFuseboxTempDirCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveFuseboxTempDirCallbackFunc) DispatchCallback(
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

type RemoveFuseboxTempDirCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveFuseboxTempDirCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveFuseboxTempDirCallback[T]) DispatchCallback(
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

type ResetHoldingSpaceOptions struct {
	// MarkTimeOfFirstAdd is "ResetHoldingSpaceOptions.markTimeOfFirstAdd"
	//
	// Optional
	//
	// NOTE: FFI_USE_MarkTimeOfFirstAdd MUST be set to true to make this field effective.
	MarkTimeOfFirstAdd bool

	FFI_USE_MarkTimeOfFirstAdd bool // for MarkTimeOfFirstAdd.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResetHoldingSpaceOptions with all fields set.
func (p ResetHoldingSpaceOptions) FromRef(ref js.Ref) ResetHoldingSpaceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResetHoldingSpaceOptions in the application heap.
func (p ResetHoldingSpaceOptions) New() js.Ref {
	return bindings.ResetHoldingSpaceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResetHoldingSpaceOptions) UpdateFrom(ref js.Ref) {
	bindings.ResetHoldingSpaceOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResetHoldingSpaceOptions) Update(ref js.Ref) {
	bindings.ResetHoldingSpaceOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResetHoldingSpaceOptions) FreeMembers(recursive bool) {
}

type RotationType uint32

const (
	_ RotationType = iota

	RotationType_ROTATE_ANY
	RotationType_ROTATE0
	RotationType_ROTATE90
	RotationType_ROTATE180
	RotationType_ROTATE270
)

func (RotationType) FromRef(str js.Ref) RotationType {
	return RotationType(bindings.ConstOfRotationType(str))
}

func (x RotationType) String() (string, bool) {
	switch x {
	case RotationType_ROTATE_ANY:
		return "RotateAny", true
	case RotationType_ROTATE0:
		return "Rotate0", true
	case RotationType_ROTATE90:
		return "Rotate90", true
	case RotationType_ROTATE180:
		return "Rotate180", true
	case RotationType_ROTATE270:
		return "Rotate270", true
	default:
		return "", false
	}
}

type ScrollableShelfState struct {
	// ScrollDistance is "ScrollableShelfState.scrollDistance"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScrollDistance MUST be set to true to make this field effective.
	ScrollDistance float64

	FFI_USE_ScrollDistance bool // for ScrollDistance.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollableShelfState with all fields set.
func (p ScrollableShelfState) FromRef(ref js.Ref) ScrollableShelfState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScrollableShelfState in the application heap.
func (p ScrollableShelfState) New() js.Ref {
	return bindings.ScrollableShelfStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollableShelfState) UpdateFrom(ref js.Ref) {
	bindings.ScrollableShelfStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollableShelfState) Update(ref js.Ref) {
	bindings.ScrollableShelfStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollableShelfState) FreeMembers(recursive bool) {
}

type SendArcOverlayColorCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn SendArcOverlayColorCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendArcOverlayColorCallbackFunc) DispatchCallback(
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

type SendArcOverlayColorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *SendArcOverlayColorCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendArcOverlayColorCallback[T]) DispatchCallback(
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

type SetOverviewModeStateCallbackFunc func(this js.Ref, finished bool) js.Ref

func (fn SetOverviewModeStateCallbackFunc) Register() js.Func[func(finished bool)] {
	return js.RegisterCallback[func(finished bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetOverviewModeStateCallbackFunc) DispatchCallback(
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

type SetOverviewModeStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, finished bool) js.Ref
	Arg T
}

func (cb *SetOverviewModeStateCallback[T]) Register() js.Func[func(finished bool)] {
	return js.RegisterCallback[func(finished bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetOverviewModeStateCallback[T]) DispatchCallback(
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

type SetShelfIconPinCallbackFunc func(this js.Ref, results js.Array[js.String]) js.Ref

func (fn SetShelfIconPinCallbackFunc) Register() js.Func[func(results js.Array[js.String])] {
	return js.RegisterCallback[func(results js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetShelfIconPinCallbackFunc) DispatchCallback(
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

type SetShelfIconPinCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results js.Array[js.String]) js.Ref
	Arg T
}

func (cb *SetShelfIconPinCallback[T]) Register() js.Func[func(results js.Array[js.String])] {
	return js.RegisterCallback[func(results js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetShelfIconPinCallback[T]) DispatchCallback(
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

type SetTabletModeEnabledCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn SetTabletModeEnabledCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetTabletModeEnabledCallbackFunc) DispatchCallback(
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

type SetTabletModeEnabledCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *SetTabletModeEnabledCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetTabletModeEnabledCallback[T]) DispatchCallback(
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

type SetWindowBoundsResult struct {
	// Bounds is "SetWindowBoundsResult.bounds"
	//
	// Optional
	//
	// NOTE: Bounds.FFI_USE MUST be set to true to get Bounds used.
	Bounds Bounds
	// DisplayId is "SetWindowBoundsResult.displayId"
	//
	// Optional
	DisplayId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetWindowBoundsResult with all fields set.
func (p SetWindowBoundsResult) FromRef(ref js.Ref) SetWindowBoundsResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetWindowBoundsResult in the application heap.
func (p SetWindowBoundsResult) New() js.Ref {
	return bindings.SetWindowBoundsResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetWindowBoundsResult) UpdateFrom(ref js.Ref) {
	bindings.SetWindowBoundsResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetWindowBoundsResult) Update(ref js.Ref) {
	bindings.SetWindowBoundsResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetWindowBoundsResult) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayId.Ref(),
	)
	p.DisplayId = p.DisplayId.FromRef(js.Undefined)
	if recursive {
		p.Bounds.FreeMembers(true)
	}
}

type ShelfIconPinUpdateParam struct {
	// AppId is "ShelfIconPinUpdateParam.appId"
	//
	// Optional
	AppId js.String
	// Pinned is "ShelfIconPinUpdateParam.pinned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pinned MUST be set to true to make this field effective.
	Pinned bool

	FFI_USE_Pinned bool // for Pinned.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShelfIconPinUpdateParam with all fields set.
func (p ShelfIconPinUpdateParam) FromRef(ref js.Ref) ShelfIconPinUpdateParam {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShelfIconPinUpdateParam in the application heap.
func (p ShelfIconPinUpdateParam) New() js.Ref {
	return bindings.ShelfIconPinUpdateParamJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShelfIconPinUpdateParam) UpdateFrom(ref js.Ref) {
	bindings.ShelfIconPinUpdateParamJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShelfIconPinUpdateParam) Update(ref js.Ref) {
	bindings.ShelfIconPinUpdateParamJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShelfIconPinUpdateParam) FreeMembers(recursive bool) {
	js.Free(
		p.AppId.Ref(),
	)
	p.AppId = p.AppId.FromRef(js.Undefined)
}

type ShelfState struct {
	// ScrollDistance is "ShelfState.scrollDistance"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScrollDistance MUST be set to true to make this field effective.
	ScrollDistance float64

	FFI_USE_ScrollDistance bool // for ScrollDistance.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShelfState with all fields set.
func (p ShelfState) FromRef(ref js.Ref) ShelfState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShelfState in the application heap.
func (p ShelfState) New() js.Ref {
	return bindings.ShelfStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShelfState) UpdateFrom(ref js.Ref) {
	bindings.ShelfStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShelfState) Update(ref js.Ref) {
	bindings.ShelfStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShelfState) FreeMembers(recursive bool) {
}

type StopFrameCountingCallbackFunc func(this js.Ref, data js.Array[FrameCountingPerSinkData]) js.Ref

func (fn StopFrameCountingCallbackFunc) Register() js.Func[func(data js.Array[FrameCountingPerSinkData])] {
	return js.RegisterCallback[func(data js.Array[FrameCountingPerSinkData])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StopFrameCountingCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[FrameCountingPerSinkData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StopFrameCountingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.Array[FrameCountingPerSinkData]) js.Ref
	Arg T
}

func (cb *StopFrameCountingCallback[T]) Register() js.Func[func(data js.Array[FrameCountingPerSinkData])] {
	return js.RegisterCallback[func(data js.Array[FrameCountingPerSinkData])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StopFrameCountingCallback[T]) DispatchCallback(
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

		js.Array[FrameCountingPerSinkData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StopSmoothnessTrackingCallbackFunc func(this js.Ref, data *DisplaySmoothnessData) js.Ref

func (fn StopSmoothnessTrackingCallbackFunc) Register() js.Func[func(data *DisplaySmoothnessData)] {
	return js.RegisterCallback[func(data *DisplaySmoothnessData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StopSmoothnessTrackingCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplaySmoothnessData
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

type StopSmoothnessTrackingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data *DisplaySmoothnessData) js.Ref
	Arg T
}

func (cb *StopSmoothnessTrackingCallback[T]) Register() js.Func[func(data *DisplaySmoothnessData)] {
	return js.RegisterCallback[func(data *DisplaySmoothnessData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StopSmoothnessTrackingCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DisplaySmoothnessData
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

type StopThroughputTrackerDataCollectionCallbackFunc func(this js.Ref, data js.Array[ThroughputTrackerAnimationData]) js.Ref

func (fn StopThroughputTrackerDataCollectionCallbackFunc) Register() js.Func[func(data js.Array[ThroughputTrackerAnimationData])] {
	return js.RegisterCallback[func(data js.Array[ThroughputTrackerAnimationData])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StopThroughputTrackerDataCollectionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ThroughputTrackerAnimationData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StopThroughputTrackerDataCollectionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.Array[ThroughputTrackerAnimationData]) js.Ref
	Arg T
}

func (cb *StopThroughputTrackerDataCollectionCallback[T]) Register() js.Func[func(data js.Array[ThroughputTrackerAnimationData])] {
	return js.RegisterCallback[func(data js.Array[ThroughputTrackerAnimationData])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StopThroughputTrackerDataCollectionCallback[T]) DispatchCallback(
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

		js.Array[ThroughputTrackerAnimationData]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TakeScreenshotCallbackFunc func(this js.Ref, base64Png js.String) js.Ref

func (fn TakeScreenshotCallbackFunc) Register() js.Func[func(base64Png js.String)] {
	return js.RegisterCallback[func(base64Png js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TakeScreenshotCallbackFunc) DispatchCallback(
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

type TakeScreenshotCallback[T any] struct {
	Fn  func(arg T, this js.Ref, base64Png js.String) js.Ref
	Arg T
}

func (cb *TakeScreenshotCallback[T]) Register() js.Func[func(base64Png js.String)] {
	return js.RegisterCallback[func(base64Png js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TakeScreenshotCallback[T]) DispatchCallback(
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

type ThemeStyle uint32

const (
	_ ThemeStyle = iota

	ThemeStyle_TONAL_SPOT
	ThemeStyle_VIBRANT
	ThemeStyle_EXPRESSIVE
	ThemeStyle_SPRITZ
	ThemeStyle_RAINBOW
	ThemeStyle_FRUIT_SALAD
)

func (ThemeStyle) FromRef(str js.Ref) ThemeStyle {
	return ThemeStyle(bindings.ConstOfThemeStyle(str))
}

func (x ThemeStyle) String() (string, bool) {
	switch x {
	case ThemeStyle_TONAL_SPOT:
		return "TonalSpot", true
	case ThemeStyle_VIBRANT:
		return "Vibrant", true
	case ThemeStyle_EXPRESSIVE:
		return "Expressive", true
	case ThemeStyle_SPRITZ:
		return "Spritz", true
	case ThemeStyle_RAINBOW:
		return "Rainbow", true
	case ThemeStyle_FRUIT_SALAD:
		return "FruitSalad", true
	default:
		return "", false
	}
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

type WMEventType uint32

const (
	_ WMEventType = iota

	WMEventType_WM_EVENT_NORMAL
	WMEventType_WM_EVENT_MAXIMIZE
	WMEventType_WM_EVENT_MINIMIZE
	WMEventType_WM_EVENT_FULLSCREEN
	WMEventType_WM_EVENT_SNAP_PRIMARY
	WMEventType_WM_EVENT_SNAP_SECONDARY
	WMEventType_WM_EVENT_FLOAT
)

func (WMEventType) FromRef(str js.Ref) WMEventType {
	return WMEventType(bindings.ConstOfWMEventType(str))
}

func (x WMEventType) String() (string, bool) {
	switch x {
	case WMEventType_WM_EVENT_NORMAL:
		return "WMEventNormal", true
	case WMEventType_WM_EVENT_MAXIMIZE:
		return "WMEventMaximize", true
	case WMEventType_WM_EVENT_MINIMIZE:
		return "WMEventMinimize", true
	case WMEventType_WM_EVENT_FULLSCREEN:
		return "WMEventFullscreen", true
	case WMEventType_WM_EVENT_SNAP_PRIMARY:
		return "WMEventSnapPrimary", true
	case WMEventType_WM_EVENT_SNAP_SECONDARY:
		return "WMEventSnapSecondary", true
	case WMEventType_WM_EVENT_FLOAT:
		return "WMEventFloat", true
	default:
		return "", false
	}
}

type WaitForDisplayRotationCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn WaitForDisplayRotationCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WaitForDisplayRotationCallbackFunc) DispatchCallback(
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

type WaitForDisplayRotationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *WaitForDisplayRotationCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WaitForDisplayRotationCallback[T]) DispatchCallback(
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

type WindowBoundsCallbackFunc func(this js.Ref, result *SetWindowBoundsResult) js.Ref

func (fn WindowBoundsCallbackFunc) Register() js.Func[func(result *SetWindowBoundsResult)] {
	return js.RegisterCallback[func(result *SetWindowBoundsResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WindowBoundsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SetWindowBoundsResult
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

type WindowBoundsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *SetWindowBoundsResult) js.Ref
	Arg T
}

func (cb *WindowBoundsCallback[T]) Register() js.Func[func(result *SetWindowBoundsResult)] {
	return js.RegisterCallback[func(result *SetWindowBoundsResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WindowBoundsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SetWindowBoundsResult
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

type WindowStateCallbackFunc func(this js.Ref, currentType WindowStateType) js.Ref

func (fn WindowStateCallbackFunc) Register() js.Func[func(currentType WindowStateType)] {
	return js.RegisterCallback[func(currentType WindowStateType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WindowStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		WindowStateType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WindowStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, currentType WindowStateType) js.Ref
	Arg T
}

func (cb *WindowStateCallback[T]) Register() js.Func[func(currentType WindowStateType)] {
	return js.RegisterCallback[func(currentType WindowStateType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WindowStateCallback[T]) DispatchCallback(
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

		WindowStateType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WindowStateChangeDict struct {
	// EventType is "WindowStateChangeDict.eventType"
	//
	// Optional
	EventType WMEventType
	// FailIfNoChange is "WindowStateChangeDict.failIfNoChange"
	//
	// Optional
	//
	// NOTE: FFI_USE_FailIfNoChange MUST be set to true to make this field effective.
	FailIfNoChange bool

	FFI_USE_FailIfNoChange bool // for FailIfNoChange.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WindowStateChangeDict with all fields set.
func (p WindowStateChangeDict) FromRef(ref js.Ref) WindowStateChangeDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WindowStateChangeDict in the application heap.
func (p WindowStateChangeDict) New() js.Ref {
	return bindings.WindowStateChangeDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WindowStateChangeDict) UpdateFrom(ref js.Ref) {
	bindings.WindowStateChangeDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WindowStateChangeDict) Update(ref js.Ref) {
	bindings.WindowStateChangeDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WindowStateChangeDict) FreeMembers(recursive bool) {
}

// HasFuncActivateAccelerator returns true if the function "WEBEXT.autotestPrivate.activateAccelerator" exists.
func HasFuncActivateAccelerator() bool {
	return js.True == bindings.HasFuncActivateAccelerator()
}

// FuncActivateAccelerator returns the function "WEBEXT.autotestPrivate.activateAccelerator".
func FuncActivateAccelerator() (fn js.Func[func(accelerator Accelerator) js.Promise[js.Boolean]]) {
	bindings.FuncActivateAccelerator(
		js.Pointer(&fn),
	)
	return
}

// ActivateAccelerator calls the function "WEBEXT.autotestPrivate.activateAccelerator" directly.
func ActivateAccelerator(accelerator Accelerator) (ret js.Promise[js.Boolean]) {
	bindings.CallActivateAccelerator(
		js.Pointer(&ret),
		js.Pointer(&accelerator),
	)

	return
}

// TryActivateAccelerator calls the function "WEBEXT.autotestPrivate.activateAccelerator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryActivateAccelerator(accelerator Accelerator) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryActivateAccelerator(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&accelerator),
	)

	return
}

// HasFuncActivateAdjacentDesksToTargetIndex returns true if the function "WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex" exists.
func HasFuncActivateAdjacentDesksToTargetIndex() bool {
	return js.True == bindings.HasFuncActivateAdjacentDesksToTargetIndex()
}

// FuncActivateAdjacentDesksToTargetIndex returns the function "WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex".
func FuncActivateAdjacentDesksToTargetIndex() (fn js.Func[func(index int32) js.Promise[js.Boolean]]) {
	bindings.FuncActivateAdjacentDesksToTargetIndex(
		js.Pointer(&fn),
	)
	return
}

// ActivateAdjacentDesksToTargetIndex calls the function "WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex" directly.
func ActivateAdjacentDesksToTargetIndex(index int32) (ret js.Promise[js.Boolean]) {
	bindings.CallActivateAdjacentDesksToTargetIndex(
		js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryActivateAdjacentDesksToTargetIndex calls the function "WEBEXT.autotestPrivate.activateAdjacentDesksToTargetIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryActivateAdjacentDesksToTargetIndex(index int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryActivateAdjacentDesksToTargetIndex(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncActivateAppWindow returns true if the function "WEBEXT.autotestPrivate.activateAppWindow" exists.
func HasFuncActivateAppWindow() bool {
	return js.True == bindings.HasFuncActivateAppWindow()
}

// FuncActivateAppWindow returns the function "WEBEXT.autotestPrivate.activateAppWindow".
func FuncActivateAppWindow() (fn js.Func[func(id int32) js.Promise[js.Void]]) {
	bindings.FuncActivateAppWindow(
		js.Pointer(&fn),
	)
	return
}

// ActivateAppWindow calls the function "WEBEXT.autotestPrivate.activateAppWindow" directly.
func ActivateAppWindow(id int32) (ret js.Promise[js.Void]) {
	bindings.CallActivateAppWindow(
		js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryActivateAppWindow calls the function "WEBEXT.autotestPrivate.activateAppWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryActivateAppWindow(id int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryActivateAppWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncActivateDeskAtIndex returns true if the function "WEBEXT.autotestPrivate.activateDeskAtIndex" exists.
func HasFuncActivateDeskAtIndex() bool {
	return js.True == bindings.HasFuncActivateDeskAtIndex()
}

// FuncActivateDeskAtIndex returns the function "WEBEXT.autotestPrivate.activateDeskAtIndex".
func FuncActivateDeskAtIndex() (fn js.Func[func(index int32) js.Promise[js.Boolean]]) {
	bindings.FuncActivateDeskAtIndex(
		js.Pointer(&fn),
	)
	return
}

// ActivateDeskAtIndex calls the function "WEBEXT.autotestPrivate.activateDeskAtIndex" directly.
func ActivateDeskAtIndex(index int32) (ret js.Promise[js.Boolean]) {
	bindings.CallActivateDeskAtIndex(
		js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryActivateDeskAtIndex calls the function "WEBEXT.autotestPrivate.activateDeskAtIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryActivateDeskAtIndex(index int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryActivateDeskAtIndex(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncAddLoginEventForTesting returns true if the function "WEBEXT.autotestPrivate.addLoginEventForTesting" exists.
func HasFuncAddLoginEventForTesting() bool {
	return js.True == bindings.HasFuncAddLoginEventForTesting()
}

// FuncAddLoginEventForTesting returns the function "WEBEXT.autotestPrivate.addLoginEventForTesting".
func FuncAddLoginEventForTesting() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncAddLoginEventForTesting(
		js.Pointer(&fn),
	)
	return
}

// AddLoginEventForTesting calls the function "WEBEXT.autotestPrivate.addLoginEventForTesting" directly.
func AddLoginEventForTesting() (ret js.Promise[js.Void]) {
	bindings.CallAddLoginEventForTesting(
		js.Pointer(&ret),
	)

	return
}

// TryAddLoginEventForTesting calls the function "WEBEXT.autotestPrivate.addLoginEventForTesting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddLoginEventForTesting() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddLoginEventForTesting(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncArcAppTracingStart returns true if the function "WEBEXT.autotestPrivate.arcAppTracingStart" exists.
func HasFuncArcAppTracingStart() bool {
	return js.True == bindings.HasFuncArcAppTracingStart()
}

// FuncArcAppTracingStart returns the function "WEBEXT.autotestPrivate.arcAppTracingStart".
func FuncArcAppTracingStart() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncArcAppTracingStart(
		js.Pointer(&fn),
	)
	return
}

// ArcAppTracingStart calls the function "WEBEXT.autotestPrivate.arcAppTracingStart" directly.
func ArcAppTracingStart() (ret js.Promise[js.Void]) {
	bindings.CallArcAppTracingStart(
		js.Pointer(&ret),
	)

	return
}

// TryArcAppTracingStart calls the function "WEBEXT.autotestPrivate.arcAppTracingStart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryArcAppTracingStart() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryArcAppTracingStart(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncArcAppTracingStopAndAnalyze returns true if the function "WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze" exists.
func HasFuncArcAppTracingStopAndAnalyze() bool {
	return js.True == bindings.HasFuncArcAppTracingStopAndAnalyze()
}

// FuncArcAppTracingStopAndAnalyze returns the function "WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze".
func FuncArcAppTracingStopAndAnalyze() (fn js.Func[func() js.Promise[ArcAppTracingInfo]]) {
	bindings.FuncArcAppTracingStopAndAnalyze(
		js.Pointer(&fn),
	)
	return
}

// ArcAppTracingStopAndAnalyze calls the function "WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze" directly.
func ArcAppTracingStopAndAnalyze() (ret js.Promise[ArcAppTracingInfo]) {
	bindings.CallArcAppTracingStopAndAnalyze(
		js.Pointer(&ret),
	)

	return
}

// TryArcAppTracingStopAndAnalyze calls the function "WEBEXT.autotestPrivate.arcAppTracingStopAndAnalyze"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryArcAppTracingStopAndAnalyze() (ret js.Promise[ArcAppTracingInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryArcAppTracingStopAndAnalyze(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBootstrapMachineLearningService returns true if the function "WEBEXT.autotestPrivate.bootstrapMachineLearningService" exists.
func HasFuncBootstrapMachineLearningService() bool {
	return js.True == bindings.HasFuncBootstrapMachineLearningService()
}

// FuncBootstrapMachineLearningService returns the function "WEBEXT.autotestPrivate.bootstrapMachineLearningService".
func FuncBootstrapMachineLearningService() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncBootstrapMachineLearningService(
		js.Pointer(&fn),
	)
	return
}

// BootstrapMachineLearningService calls the function "WEBEXT.autotestPrivate.bootstrapMachineLearningService" directly.
func BootstrapMachineLearningService() (ret js.Promise[js.Void]) {
	bindings.CallBootstrapMachineLearningService(
		js.Pointer(&ret),
	)

	return
}

// TryBootstrapMachineLearningService calls the function "WEBEXT.autotestPrivate.bootstrapMachineLearningService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryBootstrapMachineLearningService() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBootstrapMachineLearningService(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCloseApp returns true if the function "WEBEXT.autotestPrivate.closeApp" exists.
func HasFuncCloseApp() bool {
	return js.True == bindings.HasFuncCloseApp()
}

// FuncCloseApp returns the function "WEBEXT.autotestPrivate.closeApp".
func FuncCloseApp() (fn js.Func[func(appId js.String) js.Promise[js.Void]]) {
	bindings.FuncCloseApp(
		js.Pointer(&fn),
	)
	return
}

// CloseApp calls the function "WEBEXT.autotestPrivate.closeApp" directly.
func CloseApp(appId js.String) (ret js.Promise[js.Void]) {
	bindings.CallCloseApp(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryCloseApp calls the function "WEBEXT.autotestPrivate.closeApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseApp(appId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseApp(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncCloseAppWindow returns true if the function "WEBEXT.autotestPrivate.closeAppWindow" exists.
func HasFuncCloseAppWindow() bool {
	return js.True == bindings.HasFuncCloseAppWindow()
}

// FuncCloseAppWindow returns the function "WEBEXT.autotestPrivate.closeAppWindow".
func FuncCloseAppWindow() (fn js.Func[func(id int32) js.Promise[js.Void]]) {
	bindings.FuncCloseAppWindow(
		js.Pointer(&fn),
	)
	return
}

// CloseAppWindow calls the function "WEBEXT.autotestPrivate.closeAppWindow" directly.
func CloseAppWindow(id int32) (ret js.Promise[js.Void]) {
	bindings.CallCloseAppWindow(
		js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryCloseAppWindow calls the function "WEBEXT.autotestPrivate.closeAppWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseAppWindow(id int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseAppWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncCouldAllowCrostini returns true if the function "WEBEXT.autotestPrivate.couldAllowCrostini" exists.
func HasFuncCouldAllowCrostini() bool {
	return js.True == bindings.HasFuncCouldAllowCrostini()
}

// FuncCouldAllowCrostini returns the function "WEBEXT.autotestPrivate.couldAllowCrostini".
func FuncCouldAllowCrostini() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncCouldAllowCrostini(
		js.Pointer(&fn),
	)
	return
}

// CouldAllowCrostini calls the function "WEBEXT.autotestPrivate.couldAllowCrostini" directly.
func CouldAllowCrostini() (ret js.Promise[js.Boolean]) {
	bindings.CallCouldAllowCrostini(
		js.Pointer(&ret),
	)

	return
}

// TryCouldAllowCrostini calls the function "WEBEXT.autotestPrivate.couldAllowCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCouldAllowCrostini() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCouldAllowCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateNewDesk returns true if the function "WEBEXT.autotestPrivate.createNewDesk" exists.
func HasFuncCreateNewDesk() bool {
	return js.True == bindings.HasFuncCreateNewDesk()
}

// FuncCreateNewDesk returns the function "WEBEXT.autotestPrivate.createNewDesk".
func FuncCreateNewDesk() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncCreateNewDesk(
		js.Pointer(&fn),
	)
	return
}

// CreateNewDesk calls the function "WEBEXT.autotestPrivate.createNewDesk" directly.
func CreateNewDesk() (ret js.Promise[js.Boolean]) {
	bindings.CallCreateNewDesk(
		js.Pointer(&ret),
	)

	return
}

// TryCreateNewDesk calls the function "WEBEXT.autotestPrivate.createNewDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateNewDesk() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateNewDesk(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDisableAutomation returns true if the function "WEBEXT.autotestPrivate.disableAutomation" exists.
func HasFuncDisableAutomation() bool {
	return js.True == bindings.HasFuncDisableAutomation()
}

// FuncDisableAutomation returns the function "WEBEXT.autotestPrivate.disableAutomation".
func FuncDisableAutomation() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDisableAutomation(
		js.Pointer(&fn),
	)
	return
}

// DisableAutomation calls the function "WEBEXT.autotestPrivate.disableAutomation" directly.
func DisableAutomation() (ret js.Promise[js.Void]) {
	bindings.CallDisableAutomation(
		js.Pointer(&ret),
	)

	return
}

// TryDisableAutomation calls the function "WEBEXT.autotestPrivate.disableAutomation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisableAutomation() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisableAutomation(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDisableSwitchAccessDialog returns true if the function "WEBEXT.autotestPrivate.disableSwitchAccessDialog" exists.
func HasFuncDisableSwitchAccessDialog() bool {
	return js.True == bindings.HasFuncDisableSwitchAccessDialog()
}

// FuncDisableSwitchAccessDialog returns the function "WEBEXT.autotestPrivate.disableSwitchAccessDialog".
func FuncDisableSwitchAccessDialog() (fn js.Func[func()]) {
	bindings.FuncDisableSwitchAccessDialog(
		js.Pointer(&fn),
	)
	return
}

// DisableSwitchAccessDialog calls the function "WEBEXT.autotestPrivate.disableSwitchAccessDialog" directly.
func DisableSwitchAccessDialog() (ret js.Void) {
	bindings.CallDisableSwitchAccessDialog(
		js.Pointer(&ret),
	)

	return
}

// TryDisableSwitchAccessDialog calls the function "WEBEXT.autotestPrivate.disableSwitchAccessDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisableSwitchAccessDialog() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisableSwitchAccessDialog(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncEnableAssistantAndWaitForReady returns true if the function "WEBEXT.autotestPrivate.enableAssistantAndWaitForReady" exists.
func HasFuncEnableAssistantAndWaitForReady() bool {
	return js.True == bindings.HasFuncEnableAssistantAndWaitForReady()
}

// FuncEnableAssistantAndWaitForReady returns the function "WEBEXT.autotestPrivate.enableAssistantAndWaitForReady".
func FuncEnableAssistantAndWaitForReady() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncEnableAssistantAndWaitForReady(
		js.Pointer(&fn),
	)
	return
}

// EnableAssistantAndWaitForReady calls the function "WEBEXT.autotestPrivate.enableAssistantAndWaitForReady" directly.
func EnableAssistantAndWaitForReady() (ret js.Promise[js.Void]) {
	bindings.CallEnableAssistantAndWaitForReady(
		js.Pointer(&ret),
	)

	return
}

// TryEnableAssistantAndWaitForReady calls the function "WEBEXT.autotestPrivate.enableAssistantAndWaitForReady"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableAssistantAndWaitForReady() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableAssistantAndWaitForReady(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExportCrostini returns true if the function "WEBEXT.autotestPrivate.exportCrostini" exists.
func HasFuncExportCrostini() bool {
	return js.True == bindings.HasFuncExportCrostini()
}

// FuncExportCrostini returns the function "WEBEXT.autotestPrivate.exportCrostini".
func FuncExportCrostini() (fn js.Func[func(path js.String) js.Promise[js.Void]]) {
	bindings.FuncExportCrostini(
		js.Pointer(&fn),
	)
	return
}

// ExportCrostini calls the function "WEBEXT.autotestPrivate.exportCrostini" directly.
func ExportCrostini(path js.String) (ret js.Promise[js.Void]) {
	bindings.CallExportCrostini(
		js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryExportCrostini calls the function "WEBEXT.autotestPrivate.exportCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExportCrostini(path js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExportCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncForceAutoThemeMode returns true if the function "WEBEXT.autotestPrivate.forceAutoThemeMode" exists.
func HasFuncForceAutoThemeMode() bool {
	return js.True == bindings.HasFuncForceAutoThemeMode()
}

// FuncForceAutoThemeMode returns the function "WEBEXT.autotestPrivate.forceAutoThemeMode".
func FuncForceAutoThemeMode() (fn js.Func[func(darkModeEnabled bool) js.Promise[js.Void]]) {
	bindings.FuncForceAutoThemeMode(
		js.Pointer(&fn),
	)
	return
}

// ForceAutoThemeMode calls the function "WEBEXT.autotestPrivate.forceAutoThemeMode" directly.
func ForceAutoThemeMode(darkModeEnabled bool) (ret js.Promise[js.Void]) {
	bindings.CallForceAutoThemeMode(
		js.Pointer(&ret),
		js.Bool(bool(darkModeEnabled)),
	)

	return
}

// TryForceAutoThemeMode calls the function "WEBEXT.autotestPrivate.forceAutoThemeMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForceAutoThemeMode(darkModeEnabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryForceAutoThemeMode(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(darkModeEnabled)),
	)

	return
}

// HasFuncGetAccessToken returns true if the function "WEBEXT.autotestPrivate.getAccessToken" exists.
func HasFuncGetAccessToken() bool {
	return js.True == bindings.HasFuncGetAccessToken()
}

// FuncGetAccessToken returns the function "WEBEXT.autotestPrivate.getAccessToken".
func FuncGetAccessToken() (fn js.Func[func(accessTokenParams GetAccessTokenParams) js.Promise[GetAccessTokenData]]) {
	bindings.FuncGetAccessToken(
		js.Pointer(&fn),
	)
	return
}

// GetAccessToken calls the function "WEBEXT.autotestPrivate.getAccessToken" directly.
func GetAccessToken(accessTokenParams GetAccessTokenParams) (ret js.Promise[GetAccessTokenData]) {
	bindings.CallGetAccessToken(
		js.Pointer(&ret),
		js.Pointer(&accessTokenParams),
	)

	return
}

// TryGetAccessToken calls the function "WEBEXT.autotestPrivate.getAccessToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAccessToken(accessTokenParams GetAccessTokenParams) (ret js.Promise[GetAccessTokenData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAccessToken(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&accessTokenParams),
	)

	return
}

// HasFuncGetAllEnterprisePolicies returns true if the function "WEBEXT.autotestPrivate.getAllEnterprisePolicies" exists.
func HasFuncGetAllEnterprisePolicies() bool {
	return js.True == bindings.HasFuncGetAllEnterprisePolicies()
}

// FuncGetAllEnterprisePolicies returns the function "WEBEXT.autotestPrivate.getAllEnterprisePolicies".
func FuncGetAllEnterprisePolicies() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncGetAllEnterprisePolicies(
		js.Pointer(&fn),
	)
	return
}

// GetAllEnterprisePolicies calls the function "WEBEXT.autotestPrivate.getAllEnterprisePolicies" directly.
func GetAllEnterprisePolicies() (ret js.Promise[js.Any]) {
	bindings.CallGetAllEnterprisePolicies(
		js.Pointer(&ret),
	)

	return
}

// TryGetAllEnterprisePolicies calls the function "WEBEXT.autotestPrivate.getAllEnterprisePolicies"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllEnterprisePolicies() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllEnterprisePolicies(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAllInstalledApps returns true if the function "WEBEXT.autotestPrivate.getAllInstalledApps" exists.
func HasFuncGetAllInstalledApps() bool {
	return js.True == bindings.HasFuncGetAllInstalledApps()
}

// FuncGetAllInstalledApps returns the function "WEBEXT.autotestPrivate.getAllInstalledApps".
func FuncGetAllInstalledApps() (fn js.Func[func() js.Promise[js.Array[App]]]) {
	bindings.FuncGetAllInstalledApps(
		js.Pointer(&fn),
	)
	return
}

// GetAllInstalledApps calls the function "WEBEXT.autotestPrivate.getAllInstalledApps" directly.
func GetAllInstalledApps() (ret js.Promise[js.Array[App]]) {
	bindings.CallGetAllInstalledApps(
		js.Pointer(&ret),
	)

	return
}

// TryGetAllInstalledApps calls the function "WEBEXT.autotestPrivate.getAllInstalledApps"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllInstalledApps() (ret js.Promise[js.Array[App]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllInstalledApps(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAppWindowList returns true if the function "WEBEXT.autotestPrivate.getAppWindowList" exists.
func HasFuncGetAppWindowList() bool {
	return js.True == bindings.HasFuncGetAppWindowList()
}

// FuncGetAppWindowList returns the function "WEBEXT.autotestPrivate.getAppWindowList".
func FuncGetAppWindowList() (fn js.Func[func() js.Promise[js.Array[AppWindowInfo]]]) {
	bindings.FuncGetAppWindowList(
		js.Pointer(&fn),
	)
	return
}

// GetAppWindowList calls the function "WEBEXT.autotestPrivate.getAppWindowList" directly.
func GetAppWindowList() (ret js.Promise[js.Array[AppWindowInfo]]) {
	bindings.CallGetAppWindowList(
		js.Pointer(&ret),
	)

	return
}

// TryGetAppWindowList calls the function "WEBEXT.autotestPrivate.getAppWindowList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAppWindowList() (ret js.Promise[js.Array[AppWindowInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAppWindowList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetArcApp returns true if the function "WEBEXT.autotestPrivate.getArcApp" exists.
func HasFuncGetArcApp() bool {
	return js.True == bindings.HasFuncGetArcApp()
}

// FuncGetArcApp returns the function "WEBEXT.autotestPrivate.getArcApp".
func FuncGetArcApp() (fn js.Func[func(appId js.String) js.Promise[ArcAppDict]]) {
	bindings.FuncGetArcApp(
		js.Pointer(&fn),
	)
	return
}

// GetArcApp calls the function "WEBEXT.autotestPrivate.getArcApp" directly.
func GetArcApp(appId js.String) (ret js.Promise[ArcAppDict]) {
	bindings.CallGetArcApp(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryGetArcApp calls the function "WEBEXT.autotestPrivate.getArcApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcApp(appId js.String) (ret js.Promise[ArcAppDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcApp(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncGetArcAppKills returns true if the function "WEBEXT.autotestPrivate.getArcAppKills" exists.
func HasFuncGetArcAppKills() bool {
	return js.True == bindings.HasFuncGetArcAppKills()
}

// FuncGetArcAppKills returns the function "WEBEXT.autotestPrivate.getArcAppKills".
func FuncGetArcAppKills() (fn js.Func[func() js.Promise[ArcAppKillsDict]]) {
	bindings.FuncGetArcAppKills(
		js.Pointer(&fn),
	)
	return
}

// GetArcAppKills calls the function "WEBEXT.autotestPrivate.getArcAppKills" directly.
func GetArcAppKills() (ret js.Promise[ArcAppKillsDict]) {
	bindings.CallGetArcAppKills(
		js.Pointer(&ret),
	)

	return
}

// TryGetArcAppKills calls the function "WEBEXT.autotestPrivate.getArcAppKills"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcAppKills() (ret js.Promise[ArcAppKillsDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcAppKills(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetArcPackage returns true if the function "WEBEXT.autotestPrivate.getArcPackage" exists.
func HasFuncGetArcPackage() bool {
	return js.True == bindings.HasFuncGetArcPackage()
}

// FuncGetArcPackage returns the function "WEBEXT.autotestPrivate.getArcPackage".
func FuncGetArcPackage() (fn js.Func[func(packageName js.String) js.Promise[ArcPackageDict]]) {
	bindings.FuncGetArcPackage(
		js.Pointer(&fn),
	)
	return
}

// GetArcPackage calls the function "WEBEXT.autotestPrivate.getArcPackage" directly.
func GetArcPackage(packageName js.String) (ret js.Promise[ArcPackageDict]) {
	bindings.CallGetArcPackage(
		js.Pointer(&ret),
		packageName.Ref(),
	)

	return
}

// TryGetArcPackage calls the function "WEBEXT.autotestPrivate.getArcPackage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcPackage(packageName js.String) (ret js.Promise[ArcPackageDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcPackage(
		js.Pointer(&ret), js.Pointer(&exception),
		packageName.Ref(),
	)

	return
}

// HasFuncGetArcStartTime returns true if the function "WEBEXT.autotestPrivate.getArcStartTime" exists.
func HasFuncGetArcStartTime() bool {
	return js.True == bindings.HasFuncGetArcStartTime()
}

// FuncGetArcStartTime returns the function "WEBEXT.autotestPrivate.getArcStartTime".
func FuncGetArcStartTime() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	bindings.FuncGetArcStartTime(
		js.Pointer(&fn),
	)
	return
}

// GetArcStartTime calls the function "WEBEXT.autotestPrivate.getArcStartTime" directly.
func GetArcStartTime() (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetArcStartTime(
		js.Pointer(&ret),
	)

	return
}

// TryGetArcStartTime calls the function "WEBEXT.autotestPrivate.getArcStartTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcStartTime() (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcStartTime(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetArcState returns true if the function "WEBEXT.autotestPrivate.getArcState" exists.
func HasFuncGetArcState() bool {
	return js.True == bindings.HasFuncGetArcState()
}

// FuncGetArcState returns the function "WEBEXT.autotestPrivate.getArcState".
func FuncGetArcState() (fn js.Func[func() js.Promise[ArcState]]) {
	bindings.FuncGetArcState(
		js.Pointer(&fn),
	)
	return
}

// GetArcState calls the function "WEBEXT.autotestPrivate.getArcState" directly.
func GetArcState() (ret js.Promise[ArcState]) {
	bindings.CallGetArcState(
		js.Pointer(&ret),
	)

	return
}

// TryGetArcState calls the function "WEBEXT.autotestPrivate.getArcState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcState() (ret js.Promise[ArcState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetClipboardTextData returns true if the function "WEBEXT.autotestPrivate.getClipboardTextData" exists.
func HasFuncGetClipboardTextData() bool {
	return js.True == bindings.HasFuncGetClipboardTextData()
}

// FuncGetClipboardTextData returns the function "WEBEXT.autotestPrivate.getClipboardTextData".
func FuncGetClipboardTextData() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetClipboardTextData(
		js.Pointer(&fn),
	)
	return
}

// GetClipboardTextData calls the function "WEBEXT.autotestPrivate.getClipboardTextData" directly.
func GetClipboardTextData() (ret js.Promise[js.String]) {
	bindings.CallGetClipboardTextData(
		js.Pointer(&ret),
	)

	return
}

// TryGetClipboardTextData calls the function "WEBEXT.autotestPrivate.getClipboardTextData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetClipboardTextData() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetClipboardTextData(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCryptohomeRecoveryData returns true if the function "WEBEXT.autotestPrivate.getCryptohomeRecoveryData" exists.
func HasFuncGetCryptohomeRecoveryData() bool {
	return js.True == bindings.HasFuncGetCryptohomeRecoveryData()
}

// FuncGetCryptohomeRecoveryData returns the function "WEBEXT.autotestPrivate.getCryptohomeRecoveryData".
func FuncGetCryptohomeRecoveryData() (fn js.Func[func() js.Promise[CryptohomeRecoveryDataDict]]) {
	bindings.FuncGetCryptohomeRecoveryData(
		js.Pointer(&fn),
	)
	return
}

// GetCryptohomeRecoveryData calls the function "WEBEXT.autotestPrivate.getCryptohomeRecoveryData" directly.
func GetCryptohomeRecoveryData() (ret js.Promise[CryptohomeRecoveryDataDict]) {
	bindings.CallGetCryptohomeRecoveryData(
		js.Pointer(&ret),
	)

	return
}

// TryGetCryptohomeRecoveryData calls the function "WEBEXT.autotestPrivate.getCryptohomeRecoveryData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCryptohomeRecoveryData() (ret js.Promise[CryptohomeRecoveryDataDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCryptohomeRecoveryData(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCurrentInputMethodDescriptor returns true if the function "WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor" exists.
func HasFuncGetCurrentInputMethodDescriptor() bool {
	return js.True == bindings.HasFuncGetCurrentInputMethodDescriptor()
}

// FuncGetCurrentInputMethodDescriptor returns the function "WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor".
func FuncGetCurrentInputMethodDescriptor() (fn js.Func[func() js.Promise[GetCurrentInputMethodDescriptorData]]) {
	bindings.FuncGetCurrentInputMethodDescriptor(
		js.Pointer(&fn),
	)
	return
}

// GetCurrentInputMethodDescriptor calls the function "WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor" directly.
func GetCurrentInputMethodDescriptor() (ret js.Promise[GetCurrentInputMethodDescriptorData]) {
	bindings.CallGetCurrentInputMethodDescriptor(
		js.Pointer(&ret),
	)

	return
}

// TryGetCurrentInputMethodDescriptor calls the function "WEBEXT.autotestPrivate.getCurrentInputMethodDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCurrentInputMethodDescriptor() (ret js.Promise[GetCurrentInputMethodDescriptorData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCurrentInputMethodDescriptor(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDefaultPinnedAppIds returns true if the function "WEBEXT.autotestPrivate.getDefaultPinnedAppIds" exists.
func HasFuncGetDefaultPinnedAppIds() bool {
	return js.True == bindings.HasFuncGetDefaultPinnedAppIds()
}

// FuncGetDefaultPinnedAppIds returns the function "WEBEXT.autotestPrivate.getDefaultPinnedAppIds".
func FuncGetDefaultPinnedAppIds() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetDefaultPinnedAppIds(
		js.Pointer(&fn),
	)
	return
}

// GetDefaultPinnedAppIds calls the function "WEBEXT.autotestPrivate.getDefaultPinnedAppIds" directly.
func GetDefaultPinnedAppIds() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetDefaultPinnedAppIds(
		js.Pointer(&ret),
	)

	return
}

// TryGetDefaultPinnedAppIds calls the function "WEBEXT.autotestPrivate.getDefaultPinnedAppIds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDefaultPinnedAppIds() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDefaultPinnedAppIds(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeskCount returns true if the function "WEBEXT.autotestPrivate.getDeskCount" exists.
func HasFuncGetDeskCount() bool {
	return js.True == bindings.HasFuncGetDeskCount()
}

// FuncGetDeskCount returns the function "WEBEXT.autotestPrivate.getDeskCount".
func FuncGetDeskCount() (fn js.Func[func() js.Promise[js.Number[int32]]]) {
	bindings.FuncGetDeskCount(
		js.Pointer(&fn),
	)
	return
}

// GetDeskCount calls the function "WEBEXT.autotestPrivate.getDeskCount" directly.
func GetDeskCount() (ret js.Promise[js.Number[int32]]) {
	bindings.CallGetDeskCount(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeskCount calls the function "WEBEXT.autotestPrivate.getDeskCount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeskCount() (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeskCount(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDesksInfo returns true if the function "WEBEXT.autotestPrivate.getDesksInfo" exists.
func HasFuncGetDesksInfo() bool {
	return js.True == bindings.HasFuncGetDesksInfo()
}

// FuncGetDesksInfo returns the function "WEBEXT.autotestPrivate.getDesksInfo".
func FuncGetDesksInfo() (fn js.Func[func() js.Promise[DesksInfo]]) {
	bindings.FuncGetDesksInfo(
		js.Pointer(&fn),
	)
	return
}

// GetDesksInfo calls the function "WEBEXT.autotestPrivate.getDesksInfo" directly.
func GetDesksInfo() (ret js.Promise[DesksInfo]) {
	bindings.CallGetDesksInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetDesksInfo calls the function "WEBEXT.autotestPrivate.getDesksInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDesksInfo() (ret js.Promise[DesksInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDesksInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDisplaySmoothness returns true if the function "WEBEXT.autotestPrivate.getDisplaySmoothness" exists.
func HasFuncGetDisplaySmoothness() bool {
	return js.True == bindings.HasFuncGetDisplaySmoothness()
}

// FuncGetDisplaySmoothness returns the function "WEBEXT.autotestPrivate.getDisplaySmoothness".
func FuncGetDisplaySmoothness() (fn js.Func[func(displayId js.String) js.Promise[js.Number[int32]]]) {
	bindings.FuncGetDisplaySmoothness(
		js.Pointer(&fn),
	)
	return
}

// GetDisplaySmoothness calls the function "WEBEXT.autotestPrivate.getDisplaySmoothness" directly.
func GetDisplaySmoothness(displayId js.String) (ret js.Promise[js.Number[int32]]) {
	bindings.CallGetDisplaySmoothness(
		js.Pointer(&ret),
		displayId.Ref(),
	)

	return
}

// TryGetDisplaySmoothness calls the function "WEBEXT.autotestPrivate.getDisplaySmoothness"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplaySmoothness(displayId js.String) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplaySmoothness(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
	)

	return
}

// HasFuncGetExtensionsInfo returns true if the function "WEBEXT.autotestPrivate.getExtensionsInfo" exists.
func HasFuncGetExtensionsInfo() bool {
	return js.True == bindings.HasFuncGetExtensionsInfo()
}

// FuncGetExtensionsInfo returns the function "WEBEXT.autotestPrivate.getExtensionsInfo".
func FuncGetExtensionsInfo() (fn js.Func[func() js.Promise[ExtensionsInfoArray]]) {
	bindings.FuncGetExtensionsInfo(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionsInfo calls the function "WEBEXT.autotestPrivate.getExtensionsInfo" directly.
func GetExtensionsInfo() (ret js.Promise[ExtensionsInfoArray]) {
	bindings.CallGetExtensionsInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetExtensionsInfo calls the function "WEBEXT.autotestPrivate.getExtensionsInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionsInfo() (ret js.Promise[ExtensionsInfoArray], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionsInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLacrosInfo returns true if the function "WEBEXT.autotestPrivate.getLacrosInfo" exists.
func HasFuncGetLacrosInfo() bool {
	return js.True == bindings.HasFuncGetLacrosInfo()
}

// FuncGetLacrosInfo returns the function "WEBEXT.autotestPrivate.getLacrosInfo".
func FuncGetLacrosInfo() (fn js.Func[func() js.Promise[LacrosInfo]]) {
	bindings.FuncGetLacrosInfo(
		js.Pointer(&fn),
	)
	return
}

// GetLacrosInfo calls the function "WEBEXT.autotestPrivate.getLacrosInfo" directly.
func GetLacrosInfo() (ret js.Promise[LacrosInfo]) {
	bindings.CallGetLacrosInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetLacrosInfo calls the function "WEBEXT.autotestPrivate.getLacrosInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLacrosInfo() (ret js.Promise[LacrosInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLacrosInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLauncherSearchBoxState returns true if the function "WEBEXT.autotestPrivate.getLauncherSearchBoxState" exists.
func HasFuncGetLauncherSearchBoxState() bool {
	return js.True == bindings.HasFuncGetLauncherSearchBoxState()
}

// FuncGetLauncherSearchBoxState returns the function "WEBEXT.autotestPrivate.getLauncherSearchBoxState".
func FuncGetLauncherSearchBoxState() (fn js.Func[func() js.Promise[LauncherSearchBoxState]]) {
	bindings.FuncGetLauncherSearchBoxState(
		js.Pointer(&fn),
	)
	return
}

// GetLauncherSearchBoxState calls the function "WEBEXT.autotestPrivate.getLauncherSearchBoxState" directly.
func GetLauncherSearchBoxState() (ret js.Promise[LauncherSearchBoxState]) {
	bindings.CallGetLauncherSearchBoxState(
		js.Pointer(&ret),
	)

	return
}

// TryGetLauncherSearchBoxState calls the function "WEBEXT.autotestPrivate.getLauncherSearchBoxState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLauncherSearchBoxState() (ret js.Promise[LauncherSearchBoxState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLauncherSearchBoxState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetLoginEventRecorderLoginEvents returns true if the function "WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents" exists.
func HasFuncGetLoginEventRecorderLoginEvents() bool {
	return js.True == bindings.HasFuncGetLoginEventRecorderLoginEvents()
}

// FuncGetLoginEventRecorderLoginEvents returns the function "WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents".
func FuncGetLoginEventRecorderLoginEvents() (fn js.Func[func() js.Promise[js.Array[LoginEventRecorderData]]]) {
	bindings.FuncGetLoginEventRecorderLoginEvents(
		js.Pointer(&fn),
	)
	return
}

// GetLoginEventRecorderLoginEvents calls the function "WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents" directly.
func GetLoginEventRecorderLoginEvents() (ret js.Promise[js.Array[LoginEventRecorderData]]) {
	bindings.CallGetLoginEventRecorderLoginEvents(
		js.Pointer(&ret),
	)

	return
}

// TryGetLoginEventRecorderLoginEvents calls the function "WEBEXT.autotestPrivate.getLoginEventRecorderLoginEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetLoginEventRecorderLoginEvents() (ret js.Promise[js.Array[LoginEventRecorderData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetLoginEventRecorderLoginEvents(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPlayStoreState returns true if the function "WEBEXT.autotestPrivate.getPlayStoreState" exists.
func HasFuncGetPlayStoreState() bool {
	return js.True == bindings.HasFuncGetPlayStoreState()
}

// FuncGetPlayStoreState returns the function "WEBEXT.autotestPrivate.getPlayStoreState".
func FuncGetPlayStoreState() (fn js.Func[func() js.Promise[PlayStoreState]]) {
	bindings.FuncGetPlayStoreState(
		js.Pointer(&fn),
	)
	return
}

// GetPlayStoreState calls the function "WEBEXT.autotestPrivate.getPlayStoreState" directly.
func GetPlayStoreState() (ret js.Promise[PlayStoreState]) {
	bindings.CallGetPlayStoreState(
		js.Pointer(&ret),
	)

	return
}

// TryGetPlayStoreState calls the function "WEBEXT.autotestPrivate.getPlayStoreState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPlayStoreState() (ret js.Promise[PlayStoreState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPlayStoreState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPrimaryDisplayScaleFactor returns true if the function "WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor" exists.
func HasFuncGetPrimaryDisplayScaleFactor() bool {
	return js.True == bindings.HasFuncGetPrimaryDisplayScaleFactor()
}

// FuncGetPrimaryDisplayScaleFactor returns the function "WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor".
func FuncGetPrimaryDisplayScaleFactor() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	bindings.FuncGetPrimaryDisplayScaleFactor(
		js.Pointer(&fn),
	)
	return
}

// GetPrimaryDisplayScaleFactor calls the function "WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor" directly.
func GetPrimaryDisplayScaleFactor() (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetPrimaryDisplayScaleFactor(
		js.Pointer(&ret),
	)

	return
}

// TryGetPrimaryDisplayScaleFactor calls the function "WEBEXT.autotestPrivate.getPrimaryDisplayScaleFactor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrimaryDisplayScaleFactor() (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrimaryDisplayScaleFactor(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPrinterList returns true if the function "WEBEXT.autotestPrivate.getPrinterList" exists.
func HasFuncGetPrinterList() bool {
	return js.True == bindings.HasFuncGetPrinterList()
}

// FuncGetPrinterList returns the function "WEBEXT.autotestPrivate.getPrinterList".
func FuncGetPrinterList() (fn js.Func[func() js.Promise[js.Array[Printer]]]) {
	bindings.FuncGetPrinterList(
		js.Pointer(&fn),
	)
	return
}

// GetPrinterList calls the function "WEBEXT.autotestPrivate.getPrinterList" directly.
func GetPrinterList() (ret js.Promise[js.Array[Printer]]) {
	bindings.CallGetPrinterList(
		js.Pointer(&ret),
	)

	return
}

// TryGetPrinterList calls the function "WEBEXT.autotestPrivate.getPrinterList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPrinterList() (ret js.Promise[js.Array[Printer]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPrinterList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRegisteredSystemWebApps returns true if the function "WEBEXT.autotestPrivate.getRegisteredSystemWebApps" exists.
func HasFuncGetRegisteredSystemWebApps() bool {
	return js.True == bindings.HasFuncGetRegisteredSystemWebApps()
}

// FuncGetRegisteredSystemWebApps returns the function "WEBEXT.autotestPrivate.getRegisteredSystemWebApps".
func FuncGetRegisteredSystemWebApps() (fn js.Func[func() js.Promise[js.Array[SystemWebApp]]]) {
	bindings.FuncGetRegisteredSystemWebApps(
		js.Pointer(&fn),
	)
	return
}

// GetRegisteredSystemWebApps calls the function "WEBEXT.autotestPrivate.getRegisteredSystemWebApps" directly.
func GetRegisteredSystemWebApps() (ret js.Promise[js.Array[SystemWebApp]]) {
	bindings.CallGetRegisteredSystemWebApps(
		js.Pointer(&ret),
	)

	return
}

// TryGetRegisteredSystemWebApps calls the function "WEBEXT.autotestPrivate.getRegisteredSystemWebApps"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRegisteredSystemWebApps() (ret js.Promise[js.Array[SystemWebApp]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRegisteredSystemWebApps(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetScrollableShelfInfoForState returns true if the function "WEBEXT.autotestPrivate.getScrollableShelfInfoForState" exists.
func HasFuncGetScrollableShelfInfoForState() bool {
	return js.True == bindings.HasFuncGetScrollableShelfInfoForState()
}

// FuncGetScrollableShelfInfoForState returns the function "WEBEXT.autotestPrivate.getScrollableShelfInfoForState".
func FuncGetScrollableShelfInfoForState() (fn js.Func[func(state ScrollableShelfState) js.Promise[ScrollableShelfInfo]]) {
	bindings.FuncGetScrollableShelfInfoForState(
		js.Pointer(&fn),
	)
	return
}

// GetScrollableShelfInfoForState calls the function "WEBEXT.autotestPrivate.getScrollableShelfInfoForState" directly.
func GetScrollableShelfInfoForState(state ScrollableShelfState) (ret js.Promise[ScrollableShelfInfo]) {
	bindings.CallGetScrollableShelfInfoForState(
		js.Pointer(&ret),
		js.Pointer(&state),
	)

	return
}

// TryGetScrollableShelfInfoForState calls the function "WEBEXT.autotestPrivate.getScrollableShelfInfoForState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetScrollableShelfInfoForState(state ScrollableShelfState) (ret js.Promise[ScrollableShelfInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetScrollableShelfInfoForState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&state),
	)

	return
}

// HasFuncGetShelfAlignment returns true if the function "WEBEXT.autotestPrivate.getShelfAlignment" exists.
func HasFuncGetShelfAlignment() bool {
	return js.True == bindings.HasFuncGetShelfAlignment()
}

// FuncGetShelfAlignment returns the function "WEBEXT.autotestPrivate.getShelfAlignment".
func FuncGetShelfAlignment() (fn js.Func[func(displayId js.String) js.Promise[ShelfAlignmentType]]) {
	bindings.FuncGetShelfAlignment(
		js.Pointer(&fn),
	)
	return
}

// GetShelfAlignment calls the function "WEBEXT.autotestPrivate.getShelfAlignment" directly.
func GetShelfAlignment(displayId js.String) (ret js.Promise[ShelfAlignmentType]) {
	bindings.CallGetShelfAlignment(
		js.Pointer(&ret),
		displayId.Ref(),
	)

	return
}

// TryGetShelfAlignment calls the function "WEBEXT.autotestPrivate.getShelfAlignment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetShelfAlignment(displayId js.String) (ret js.Promise[ShelfAlignmentType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetShelfAlignment(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
	)

	return
}

// HasFuncGetShelfAutoHideBehavior returns true if the function "WEBEXT.autotestPrivate.getShelfAutoHideBehavior" exists.
func HasFuncGetShelfAutoHideBehavior() bool {
	return js.True == bindings.HasFuncGetShelfAutoHideBehavior()
}

// FuncGetShelfAutoHideBehavior returns the function "WEBEXT.autotestPrivate.getShelfAutoHideBehavior".
func FuncGetShelfAutoHideBehavior() (fn js.Func[func(displayId js.String) js.Promise[js.String]]) {
	bindings.FuncGetShelfAutoHideBehavior(
		js.Pointer(&fn),
	)
	return
}

// GetShelfAutoHideBehavior calls the function "WEBEXT.autotestPrivate.getShelfAutoHideBehavior" directly.
func GetShelfAutoHideBehavior(displayId js.String) (ret js.Promise[js.String]) {
	bindings.CallGetShelfAutoHideBehavior(
		js.Pointer(&ret),
		displayId.Ref(),
	)

	return
}

// TryGetShelfAutoHideBehavior calls the function "WEBEXT.autotestPrivate.getShelfAutoHideBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetShelfAutoHideBehavior(displayId js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetShelfAutoHideBehavior(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
	)

	return
}

// HasFuncGetShelfItems returns true if the function "WEBEXT.autotestPrivate.getShelfItems" exists.
func HasFuncGetShelfItems() bool {
	return js.True == bindings.HasFuncGetShelfItems()
}

// FuncGetShelfItems returns the function "WEBEXT.autotestPrivate.getShelfItems".
func FuncGetShelfItems() (fn js.Func[func() js.Promise[js.Array[ShelfItem]]]) {
	bindings.FuncGetShelfItems(
		js.Pointer(&fn),
	)
	return
}

// GetShelfItems calls the function "WEBEXT.autotestPrivate.getShelfItems" directly.
func GetShelfItems() (ret js.Promise[js.Array[ShelfItem]]) {
	bindings.CallGetShelfItems(
		js.Pointer(&ret),
	)

	return
}

// TryGetShelfItems calls the function "WEBEXT.autotestPrivate.getShelfItems"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetShelfItems() (ret js.Promise[js.Array[ShelfItem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetShelfItems(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetShelfUIInfoForState returns true if the function "WEBEXT.autotestPrivate.getShelfUIInfoForState" exists.
func HasFuncGetShelfUIInfoForState() bool {
	return js.True == bindings.HasFuncGetShelfUIInfoForState()
}

// FuncGetShelfUIInfoForState returns the function "WEBEXT.autotestPrivate.getShelfUIInfoForState".
func FuncGetShelfUIInfoForState() (fn js.Func[func(state ShelfState) js.Promise[ShelfUIInfo]]) {
	bindings.FuncGetShelfUIInfoForState(
		js.Pointer(&fn),
	)
	return
}

// GetShelfUIInfoForState calls the function "WEBEXT.autotestPrivate.getShelfUIInfoForState" directly.
func GetShelfUIInfoForState(state ShelfState) (ret js.Promise[ShelfUIInfo]) {
	bindings.CallGetShelfUIInfoForState(
		js.Pointer(&ret),
		js.Pointer(&state),
	)

	return
}

// TryGetShelfUIInfoForState calls the function "WEBEXT.autotestPrivate.getShelfUIInfoForState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetShelfUIInfoForState(state ShelfState) (ret js.Promise[ShelfUIInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetShelfUIInfoForState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&state),
	)

	return
}

// HasFuncGetThroughputTrackerData returns true if the function "WEBEXT.autotestPrivate.getThroughputTrackerData" exists.
func HasFuncGetThroughputTrackerData() bool {
	return js.True == bindings.HasFuncGetThroughputTrackerData()
}

// FuncGetThroughputTrackerData returns the function "WEBEXT.autotestPrivate.getThroughputTrackerData".
func FuncGetThroughputTrackerData() (fn js.Func[func() js.Promise[js.Array[ThroughputTrackerAnimationData]]]) {
	bindings.FuncGetThroughputTrackerData(
		js.Pointer(&fn),
	)
	return
}

// GetThroughputTrackerData calls the function "WEBEXT.autotestPrivate.getThroughputTrackerData" directly.
func GetThroughputTrackerData() (ret js.Promise[js.Array[ThroughputTrackerAnimationData]]) {
	bindings.CallGetThroughputTrackerData(
		js.Pointer(&ret),
	)

	return
}

// TryGetThroughputTrackerData calls the function "WEBEXT.autotestPrivate.getThroughputTrackerData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetThroughputTrackerData() (ret js.Promise[js.Array[ThroughputTrackerAnimationData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetThroughputTrackerData(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVisibleNotifications returns true if the function "WEBEXT.autotestPrivate.getVisibleNotifications" exists.
func HasFuncGetVisibleNotifications() bool {
	return js.True == bindings.HasFuncGetVisibleNotifications()
}

// FuncGetVisibleNotifications returns the function "WEBEXT.autotestPrivate.getVisibleNotifications".
func FuncGetVisibleNotifications() (fn js.Func[func() js.Promise[js.Array[Notification]]]) {
	bindings.FuncGetVisibleNotifications(
		js.Pointer(&fn),
	)
	return
}

// GetVisibleNotifications calls the function "WEBEXT.autotestPrivate.getVisibleNotifications" directly.
func GetVisibleNotifications() (ret js.Promise[js.Array[Notification]]) {
	bindings.CallGetVisibleNotifications(
		js.Pointer(&ret),
	)

	return
}

// TryGetVisibleNotifications calls the function "WEBEXT.autotestPrivate.getVisibleNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVisibleNotifications() (ret js.Promise[js.Array[Notification]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVisibleNotifications(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncImportCrostini returns true if the function "WEBEXT.autotestPrivate.importCrostini" exists.
func HasFuncImportCrostini() bool {
	return js.True == bindings.HasFuncImportCrostini()
}

// FuncImportCrostini returns the function "WEBEXT.autotestPrivate.importCrostini".
func FuncImportCrostini() (fn js.Func[func(path js.String) js.Promise[js.Void]]) {
	bindings.FuncImportCrostini(
		js.Pointer(&fn),
	)
	return
}

// ImportCrostini calls the function "WEBEXT.autotestPrivate.importCrostini" directly.
func ImportCrostini(path js.String) (ret js.Promise[js.Void]) {
	bindings.CallImportCrostini(
		js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryImportCrostini calls the function "WEBEXT.autotestPrivate.importCrostini"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImportCrostini(path js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImportCrostini(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncInitializeEvents returns true if the function "WEBEXT.autotestPrivate.initializeEvents" exists.
func HasFuncInitializeEvents() bool {
	return js.True == bindings.HasFuncInitializeEvents()
}

// FuncInitializeEvents returns the function "WEBEXT.autotestPrivate.initializeEvents".
func FuncInitializeEvents() (fn js.Func[func()]) {
	bindings.FuncInitializeEvents(
		js.Pointer(&fn),
	)
	return
}

// InitializeEvents calls the function "WEBEXT.autotestPrivate.initializeEvents" directly.
func InitializeEvents() (ret js.Void) {
	bindings.CallInitializeEvents(
		js.Pointer(&ret),
	)

	return
}

// TryInitializeEvents calls the function "WEBEXT.autotestPrivate.initializeEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInitializeEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInitializeEvents(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInstallBorealis returns true if the function "WEBEXT.autotestPrivate.installBorealis" exists.
func HasFuncInstallBorealis() bool {
	return js.True == bindings.HasFuncInstallBorealis()
}

// FuncInstallBorealis returns the function "WEBEXT.autotestPrivate.installBorealis".
func FuncInstallBorealis() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncInstallBorealis(
		js.Pointer(&fn),
	)
	return
}

// InstallBorealis calls the function "WEBEXT.autotestPrivate.installBorealis" directly.
func InstallBorealis() (ret js.Promise[js.Void]) {
	bindings.CallInstallBorealis(
		js.Pointer(&ret),
	)

	return
}

// TryInstallBorealis calls the function "WEBEXT.autotestPrivate.installBorealis"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallBorealis() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallBorealis(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInstallBruschetta returns true if the function "WEBEXT.autotestPrivate.installBruschetta" exists.
func HasFuncInstallBruschetta() bool {
	return js.True == bindings.HasFuncInstallBruschetta()
}

// FuncInstallBruschetta returns the function "WEBEXT.autotestPrivate.installBruschetta".
func FuncInstallBruschetta() (fn js.Func[func(vm_name js.String) js.Promise[js.Void]]) {
	bindings.FuncInstallBruschetta(
		js.Pointer(&fn),
	)
	return
}

// InstallBruschetta calls the function "WEBEXT.autotestPrivate.installBruschetta" directly.
func InstallBruschetta(vm_name js.String) (ret js.Promise[js.Void]) {
	bindings.CallInstallBruschetta(
		js.Pointer(&ret),
		vm_name.Ref(),
	)

	return
}

// TryInstallBruschetta calls the function "WEBEXT.autotestPrivate.installBruschetta"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallBruschetta(vm_name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallBruschetta(
		js.Pointer(&ret), js.Pointer(&exception),
		vm_name.Ref(),
	)

	return
}

// HasFuncInstallPWAForCurrentURL returns true if the function "WEBEXT.autotestPrivate.installPWAForCurrentURL" exists.
func HasFuncInstallPWAForCurrentURL() bool {
	return js.True == bindings.HasFuncInstallPWAForCurrentURL()
}

// FuncInstallPWAForCurrentURL returns the function "WEBEXT.autotestPrivate.installPWAForCurrentURL".
func FuncInstallPWAForCurrentURL() (fn js.Func[func(timeout_ms int32) js.Promise[js.String]]) {
	bindings.FuncInstallPWAForCurrentURL(
		js.Pointer(&fn),
	)
	return
}

// InstallPWAForCurrentURL calls the function "WEBEXT.autotestPrivate.installPWAForCurrentURL" directly.
func InstallPWAForCurrentURL(timeout_ms int32) (ret js.Promise[js.String]) {
	bindings.CallInstallPWAForCurrentURL(
		js.Pointer(&ret),
		int32(timeout_ms),
	)

	return
}

// TryInstallPWAForCurrentURL calls the function "WEBEXT.autotestPrivate.installPWAForCurrentURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallPWAForCurrentURL(timeout_ms int32) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallPWAForCurrentURL(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(timeout_ms),
	)

	return
}

// HasFuncIsAppShown returns true if the function "WEBEXT.autotestPrivate.isAppShown" exists.
func HasFuncIsAppShown() bool {
	return js.True == bindings.HasFuncIsAppShown()
}

// FuncIsAppShown returns the function "WEBEXT.autotestPrivate.isAppShown".
func FuncIsAppShown() (fn js.Func[func(appId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsAppShown(
		js.Pointer(&fn),
	)
	return
}

// IsAppShown calls the function "WEBEXT.autotestPrivate.isAppShown" directly.
func IsAppShown(appId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsAppShown(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryIsAppShown calls the function "WEBEXT.autotestPrivate.isAppShown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAppShown(appId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAppShown(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncIsArcPackageListInitialRefreshed returns true if the function "WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed" exists.
func HasFuncIsArcPackageListInitialRefreshed() bool {
	return js.True == bindings.HasFuncIsArcPackageListInitialRefreshed()
}

// FuncIsArcPackageListInitialRefreshed returns the function "WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed".
func FuncIsArcPackageListInitialRefreshed() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsArcPackageListInitialRefreshed(
		js.Pointer(&fn),
	)
	return
}

// IsArcPackageListInitialRefreshed calls the function "WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed" directly.
func IsArcPackageListInitialRefreshed() (ret js.Promise[js.Boolean]) {
	bindings.CallIsArcPackageListInitialRefreshed(
		js.Pointer(&ret),
	)

	return
}

// TryIsArcPackageListInitialRefreshed calls the function "WEBEXT.autotestPrivate.isArcPackageListInitialRefreshed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsArcPackageListInitialRefreshed() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsArcPackageListInitialRefreshed(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsArcProvisioned returns true if the function "WEBEXT.autotestPrivate.isArcProvisioned" exists.
func HasFuncIsArcProvisioned() bool {
	return js.True == bindings.HasFuncIsArcProvisioned()
}

// FuncIsArcProvisioned returns the function "WEBEXT.autotestPrivate.isArcProvisioned".
func FuncIsArcProvisioned() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsArcProvisioned(
		js.Pointer(&fn),
	)
	return
}

// IsArcProvisioned calls the function "WEBEXT.autotestPrivate.isArcProvisioned" directly.
func IsArcProvisioned() (ret js.Promise[js.Boolean]) {
	bindings.CallIsArcProvisioned(
		js.Pointer(&ret),
	)

	return
}

// TryIsArcProvisioned calls the function "WEBEXT.autotestPrivate.isArcProvisioned"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsArcProvisioned() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsArcProvisioned(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsFeatureEnabled returns true if the function "WEBEXT.autotestPrivate.isFeatureEnabled" exists.
func HasFuncIsFeatureEnabled() bool {
	return js.True == bindings.HasFuncIsFeatureEnabled()
}

// FuncIsFeatureEnabled returns the function "WEBEXT.autotestPrivate.isFeatureEnabled".
func FuncIsFeatureEnabled() (fn js.Func[func(feature_name js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsFeatureEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsFeatureEnabled calls the function "WEBEXT.autotestPrivate.isFeatureEnabled" directly.
func IsFeatureEnabled(feature_name js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsFeatureEnabled(
		js.Pointer(&ret),
		feature_name.Ref(),
	)

	return
}

// TryIsFeatureEnabled calls the function "WEBEXT.autotestPrivate.isFeatureEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsFeatureEnabled(feature_name js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsFeatureEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		feature_name.Ref(),
	)

	return
}

// HasFuncIsInputMethodReadyForTesting returns true if the function "WEBEXT.autotestPrivate.isInputMethodReadyForTesting" exists.
func HasFuncIsInputMethodReadyForTesting() bool {
	return js.True == bindings.HasFuncIsInputMethodReadyForTesting()
}

// FuncIsInputMethodReadyForTesting returns the function "WEBEXT.autotestPrivate.isInputMethodReadyForTesting".
func FuncIsInputMethodReadyForTesting() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsInputMethodReadyForTesting(
		js.Pointer(&fn),
	)
	return
}

// IsInputMethodReadyForTesting calls the function "WEBEXT.autotestPrivate.isInputMethodReadyForTesting" directly.
func IsInputMethodReadyForTesting() (ret js.Promise[js.Boolean]) {
	bindings.CallIsInputMethodReadyForTesting(
		js.Pointer(&ret),
	)

	return
}

// TryIsInputMethodReadyForTesting calls the function "WEBEXT.autotestPrivate.isInputMethodReadyForTesting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsInputMethodReadyForTesting() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsInputMethodReadyForTesting(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsSystemWebAppOpen returns true if the function "WEBEXT.autotestPrivate.isSystemWebAppOpen" exists.
func HasFuncIsSystemWebAppOpen() bool {
	return js.True == bindings.HasFuncIsSystemWebAppOpen()
}

// FuncIsSystemWebAppOpen returns the function "WEBEXT.autotestPrivate.isSystemWebAppOpen".
func FuncIsSystemWebAppOpen() (fn js.Func[func(appId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsSystemWebAppOpen(
		js.Pointer(&fn),
	)
	return
}

// IsSystemWebAppOpen calls the function "WEBEXT.autotestPrivate.isSystemWebAppOpen" directly.
func IsSystemWebAppOpen(appId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsSystemWebAppOpen(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryIsSystemWebAppOpen calls the function "WEBEXT.autotestPrivate.isSystemWebAppOpen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsSystemWebAppOpen(appId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsSystemWebAppOpen(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncIsTabletModeEnabled returns true if the function "WEBEXT.autotestPrivate.isTabletModeEnabled" exists.
func HasFuncIsTabletModeEnabled() bool {
	return js.True == bindings.HasFuncIsTabletModeEnabled()
}

// FuncIsTabletModeEnabled returns the function "WEBEXT.autotestPrivate.isTabletModeEnabled".
func FuncIsTabletModeEnabled() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsTabletModeEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsTabletModeEnabled calls the function "WEBEXT.autotestPrivate.isTabletModeEnabled" directly.
func IsTabletModeEnabled() (ret js.Promise[js.Boolean]) {
	bindings.CallIsTabletModeEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryIsTabletModeEnabled calls the function "WEBEXT.autotestPrivate.isTabletModeEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsTabletModeEnabled() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsTabletModeEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLaunchApp returns true if the function "WEBEXT.autotestPrivate.launchApp" exists.
func HasFuncLaunchApp() bool {
	return js.True == bindings.HasFuncLaunchApp()
}

// FuncLaunchApp returns the function "WEBEXT.autotestPrivate.launchApp".
func FuncLaunchApp() (fn js.Func[func(appId js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchApp(
		js.Pointer(&fn),
	)
	return
}

// LaunchApp calls the function "WEBEXT.autotestPrivate.launchApp" directly.
func LaunchApp(appId js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchApp(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryLaunchApp calls the function "WEBEXT.autotestPrivate.launchApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchApp(appId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchApp(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncLaunchFilesAppToPath returns true if the function "WEBEXT.autotestPrivate.launchFilesAppToPath" exists.
func HasFuncLaunchFilesAppToPath() bool {
	return js.True == bindings.HasFuncLaunchFilesAppToPath()
}

// FuncLaunchFilesAppToPath returns the function "WEBEXT.autotestPrivate.launchFilesAppToPath".
func FuncLaunchFilesAppToPath() (fn js.Func[func(absolutePath js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchFilesAppToPath(
		js.Pointer(&fn),
	)
	return
}

// LaunchFilesAppToPath calls the function "WEBEXT.autotestPrivate.launchFilesAppToPath" directly.
func LaunchFilesAppToPath(absolutePath js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchFilesAppToPath(
		js.Pointer(&ret),
		absolutePath.Ref(),
	)

	return
}

// TryLaunchFilesAppToPath calls the function "WEBEXT.autotestPrivate.launchFilesAppToPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchFilesAppToPath(absolutePath js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchFilesAppToPath(
		js.Pointer(&ret), js.Pointer(&exception),
		absolutePath.Ref(),
	)

	return
}

// HasFuncLaunchSystemWebApp returns true if the function "WEBEXT.autotestPrivate.launchSystemWebApp" exists.
func HasFuncLaunchSystemWebApp() bool {
	return js.True == bindings.HasFuncLaunchSystemWebApp()
}

// FuncLaunchSystemWebApp returns the function "WEBEXT.autotestPrivate.launchSystemWebApp".
func FuncLaunchSystemWebApp() (fn js.Func[func(appName js.String, url js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchSystemWebApp(
		js.Pointer(&fn),
	)
	return
}

// LaunchSystemWebApp calls the function "WEBEXT.autotestPrivate.launchSystemWebApp" directly.
func LaunchSystemWebApp(appName js.String, url js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchSystemWebApp(
		js.Pointer(&ret),
		appName.Ref(),
		url.Ref(),
	)

	return
}

// TryLaunchSystemWebApp calls the function "WEBEXT.autotestPrivate.launchSystemWebApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchSystemWebApp(appName js.String, url js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchSystemWebApp(
		js.Pointer(&ret), js.Pointer(&exception),
		appName.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncLoadSmartDimComponent returns true if the function "WEBEXT.autotestPrivate.loadSmartDimComponent" exists.
func HasFuncLoadSmartDimComponent() bool {
	return js.True == bindings.HasFuncLoadSmartDimComponent()
}

// FuncLoadSmartDimComponent returns the function "WEBEXT.autotestPrivate.loadSmartDimComponent".
func FuncLoadSmartDimComponent() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncLoadSmartDimComponent(
		js.Pointer(&fn),
	)
	return
}

// LoadSmartDimComponent calls the function "WEBEXT.autotestPrivate.loadSmartDimComponent" directly.
func LoadSmartDimComponent() (ret js.Promise[js.Void]) {
	bindings.CallLoadSmartDimComponent(
		js.Pointer(&ret),
	)

	return
}

// TryLoadSmartDimComponent calls the function "WEBEXT.autotestPrivate.loadSmartDimComponent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadSmartDimComponent() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadSmartDimComponent(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLockScreen returns true if the function "WEBEXT.autotestPrivate.lockScreen" exists.
func HasFuncLockScreen() bool {
	return js.True == bindings.HasFuncLockScreen()
}

// FuncLockScreen returns the function "WEBEXT.autotestPrivate.lockScreen".
func FuncLockScreen() (fn js.Func[func()]) {
	bindings.FuncLockScreen(
		js.Pointer(&fn),
	)
	return
}

// LockScreen calls the function "WEBEXT.autotestPrivate.lockScreen" directly.
func LockScreen() (ret js.Void) {
	bindings.CallLockScreen(
		js.Pointer(&ret),
	)

	return
}

// TryLockScreen calls the function "WEBEXT.autotestPrivate.lockScreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLockScreen() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLockScreen(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLoginStatus returns true if the function "WEBEXT.autotestPrivate.loginStatus" exists.
func HasFuncLoginStatus() bool {
	return js.True == bindings.HasFuncLoginStatus()
}

// FuncLoginStatus returns the function "WEBEXT.autotestPrivate.loginStatus".
func FuncLoginStatus() (fn js.Func[func() js.Promise[LoginStatusDict]]) {
	bindings.FuncLoginStatus(
		js.Pointer(&fn),
	)
	return
}

// LoginStatus calls the function "WEBEXT.autotestPrivate.loginStatus" directly.
func LoginStatus() (ret js.Promise[LoginStatusDict]) {
	bindings.CallLoginStatus(
		js.Pointer(&ret),
	)

	return
}

// TryLoginStatus calls the function "WEBEXT.autotestPrivate.loginStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoginStatus() (ret js.Promise[LoginStatusDict], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoginStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLogout returns true if the function "WEBEXT.autotestPrivate.logout" exists.
func HasFuncLogout() bool {
	return js.True == bindings.HasFuncLogout()
}

// FuncLogout returns the function "WEBEXT.autotestPrivate.logout".
func FuncLogout() (fn js.Func[func()]) {
	bindings.FuncLogout(
		js.Pointer(&fn),
	)
	return
}

// Logout calls the function "WEBEXT.autotestPrivate.logout" directly.
func Logout() (ret js.Void) {
	bindings.CallLogout(
		js.Pointer(&ret),
	)

	return
}

// TryLogout calls the function "WEBEXT.autotestPrivate.logout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLogout() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLogout(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMakeFuseboxTempDir returns true if the function "WEBEXT.autotestPrivate.makeFuseboxTempDir" exists.
func HasFuncMakeFuseboxTempDir() bool {
	return js.True == bindings.HasFuncMakeFuseboxTempDir()
}

// FuncMakeFuseboxTempDir returns the function "WEBEXT.autotestPrivate.makeFuseboxTempDir".
func FuncMakeFuseboxTempDir() (fn js.Func[func() js.Promise[MakeFuseboxTempDirData]]) {
	bindings.FuncMakeFuseboxTempDir(
		js.Pointer(&fn),
	)
	return
}

// MakeFuseboxTempDir calls the function "WEBEXT.autotestPrivate.makeFuseboxTempDir" directly.
func MakeFuseboxTempDir() (ret js.Promise[MakeFuseboxTempDirData]) {
	bindings.CallMakeFuseboxTempDir(
		js.Pointer(&ret),
	)

	return
}

// TryMakeFuseboxTempDir calls the function "WEBEXT.autotestPrivate.makeFuseboxTempDir"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMakeFuseboxTempDir() (ret js.Promise[MakeFuseboxTempDirData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMakeFuseboxTempDir(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMouseClick returns true if the function "WEBEXT.autotestPrivate.mouseClick" exists.
func HasFuncMouseClick() bool {
	return js.True == bindings.HasFuncMouseClick()
}

// FuncMouseClick returns the function "WEBEXT.autotestPrivate.mouseClick".
func FuncMouseClick() (fn js.Func[func(button MouseButton) js.Promise[js.Void]]) {
	bindings.FuncMouseClick(
		js.Pointer(&fn),
	)
	return
}

// MouseClick calls the function "WEBEXT.autotestPrivate.mouseClick" directly.
func MouseClick(button MouseButton) (ret js.Promise[js.Void]) {
	bindings.CallMouseClick(
		js.Pointer(&ret),
		uint32(button),
	)

	return
}

// TryMouseClick calls the function "WEBEXT.autotestPrivate.mouseClick"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMouseClick(button MouseButton) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseClick(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(button),
	)

	return
}

// HasFuncMouseMove returns true if the function "WEBEXT.autotestPrivate.mouseMove" exists.
func HasFuncMouseMove() bool {
	return js.True == bindings.HasFuncMouseMove()
}

// FuncMouseMove returns the function "WEBEXT.autotestPrivate.mouseMove".
func FuncMouseMove() (fn js.Func[func(location Location, duration_in_ms float64) js.Promise[js.Void]]) {
	bindings.FuncMouseMove(
		js.Pointer(&fn),
	)
	return
}

// MouseMove calls the function "WEBEXT.autotestPrivate.mouseMove" directly.
func MouseMove(location Location, duration_in_ms float64) (ret js.Promise[js.Void]) {
	bindings.CallMouseMove(
		js.Pointer(&ret),
		js.Pointer(&location),
		float64(duration_in_ms),
	)

	return
}

// TryMouseMove calls the function "WEBEXT.autotestPrivate.mouseMove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMouseMove(location Location, duration_in_ms float64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseMove(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&location),
		float64(duration_in_ms),
	)

	return
}

// HasFuncMousePress returns true if the function "WEBEXT.autotestPrivate.mousePress" exists.
func HasFuncMousePress() bool {
	return js.True == bindings.HasFuncMousePress()
}

// FuncMousePress returns the function "WEBEXT.autotestPrivate.mousePress".
func FuncMousePress() (fn js.Func[func(button MouseButton) js.Promise[js.Void]]) {
	bindings.FuncMousePress(
		js.Pointer(&fn),
	)
	return
}

// MousePress calls the function "WEBEXT.autotestPrivate.mousePress" directly.
func MousePress(button MouseButton) (ret js.Promise[js.Void]) {
	bindings.CallMousePress(
		js.Pointer(&ret),
		uint32(button),
	)

	return
}

// TryMousePress calls the function "WEBEXT.autotestPrivate.mousePress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMousePress(button MouseButton) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMousePress(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(button),
	)

	return
}

// HasFuncMouseRelease returns true if the function "WEBEXT.autotestPrivate.mouseRelease" exists.
func HasFuncMouseRelease() bool {
	return js.True == bindings.HasFuncMouseRelease()
}

// FuncMouseRelease returns the function "WEBEXT.autotestPrivate.mouseRelease".
func FuncMouseRelease() (fn js.Func[func(button MouseButton) js.Promise[js.Void]]) {
	bindings.FuncMouseRelease(
		js.Pointer(&fn),
	)
	return
}

// MouseRelease calls the function "WEBEXT.autotestPrivate.mouseRelease" directly.
func MouseRelease(button MouseButton) (ret js.Promise[js.Void]) {
	bindings.CallMouseRelease(
		js.Pointer(&ret),
		uint32(button),
	)

	return
}

// TryMouseRelease calls the function "WEBEXT.autotestPrivate.mouseRelease"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMouseRelease(button MouseButton) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseRelease(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(button),
	)

	return
}

type OnClipboardDataChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnClipboardDataChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClipboardDataChangedEventCallbackFunc) DispatchCallback(
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

type OnClipboardDataChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnClipboardDataChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClipboardDataChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClipboardDataChanged returns true if the function "WEBEXT.autotestPrivate.onClipboardDataChanged.addListener" exists.
func HasFuncOnClipboardDataChanged() bool {
	return js.True == bindings.HasFuncOnClipboardDataChanged()
}

// FuncOnClipboardDataChanged returns the function "WEBEXT.autotestPrivate.onClipboardDataChanged.addListener".
func FuncOnClipboardDataChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OnClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.addListener" directly.
func OnClipboardDataChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClipboardDataChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClipboardDataChanged returns true if the function "WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener" exists.
func HasFuncOffClipboardDataChanged() bool {
	return js.True == bindings.HasFuncOffClipboardDataChanged()
}

// FuncOffClipboardDataChanged returns the function "WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener".
func FuncOffClipboardDataChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OffClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener" directly.
func OffClipboardDataChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClipboardDataChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClipboardDataChanged returns true if the function "WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener" exists.
func HasFuncHasOnClipboardDataChanged() bool {
	return js.True == bindings.HasFuncHasOnClipboardDataChanged()
}

// FuncHasOnClipboardDataChanged returns the function "WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener".
func FuncHasOnClipboardDataChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener" directly.
func HasOnClipboardDataChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClipboardDataChanged calls the function "WEBEXT.autotestPrivate.onClipboardDataChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClipboardDataChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncPinShelfIcon returns true if the function "WEBEXT.autotestPrivate.pinShelfIcon" exists.
func HasFuncPinShelfIcon() bool {
	return js.True == bindings.HasFuncPinShelfIcon()
}

// FuncPinShelfIcon returns the function "WEBEXT.autotestPrivate.pinShelfIcon".
func FuncPinShelfIcon() (fn js.Func[func(appId js.String) js.Promise[js.Void]]) {
	bindings.FuncPinShelfIcon(
		js.Pointer(&fn),
	)
	return
}

// PinShelfIcon calls the function "WEBEXT.autotestPrivate.pinShelfIcon" directly.
func PinShelfIcon(appId js.String) (ret js.Promise[js.Void]) {
	bindings.CallPinShelfIcon(
		js.Pointer(&ret),
		appId.Ref(),
	)

	return
}

// TryPinShelfIcon calls the function "WEBEXT.autotestPrivate.pinShelfIcon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPinShelfIcon(appId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPinShelfIcon(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
	)

	return
}

// HasFuncRefreshEnterprisePolicies returns true if the function "WEBEXT.autotestPrivate.refreshEnterprisePolicies" exists.
func HasFuncRefreshEnterprisePolicies() bool {
	return js.True == bindings.HasFuncRefreshEnterprisePolicies()
}

// FuncRefreshEnterprisePolicies returns the function "WEBEXT.autotestPrivate.refreshEnterprisePolicies".
func FuncRefreshEnterprisePolicies() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRefreshEnterprisePolicies(
		js.Pointer(&fn),
	)
	return
}

// RefreshEnterprisePolicies calls the function "WEBEXT.autotestPrivate.refreshEnterprisePolicies" directly.
func RefreshEnterprisePolicies() (ret js.Promise[js.Void]) {
	bindings.CallRefreshEnterprisePolicies(
		js.Pointer(&ret),
	)

	return
}

// TryRefreshEnterprisePolicies calls the function "WEBEXT.autotestPrivate.refreshEnterprisePolicies"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRefreshEnterprisePolicies() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRefreshEnterprisePolicies(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRefreshRemoteCommands returns true if the function "WEBEXT.autotestPrivate.refreshRemoteCommands" exists.
func HasFuncRefreshRemoteCommands() bool {
	return js.True == bindings.HasFuncRefreshRemoteCommands()
}

// FuncRefreshRemoteCommands returns the function "WEBEXT.autotestPrivate.refreshRemoteCommands".
func FuncRefreshRemoteCommands() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRefreshRemoteCommands(
		js.Pointer(&fn),
	)
	return
}

// RefreshRemoteCommands calls the function "WEBEXT.autotestPrivate.refreshRemoteCommands" directly.
func RefreshRemoteCommands() (ret js.Promise[js.Void]) {
	bindings.CallRefreshRemoteCommands(
		js.Pointer(&ret),
	)

	return
}

// TryRefreshRemoteCommands calls the function "WEBEXT.autotestPrivate.refreshRemoteCommands"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRefreshRemoteCommands() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRefreshRemoteCommands(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRegisterComponent returns true if the function "WEBEXT.autotestPrivate.registerComponent" exists.
func HasFuncRegisterComponent() bool {
	return js.True == bindings.HasFuncRegisterComponent()
}

// FuncRegisterComponent returns the function "WEBEXT.autotestPrivate.registerComponent".
func FuncRegisterComponent() (fn js.Func[func(name js.String, path js.String)]) {
	bindings.FuncRegisterComponent(
		js.Pointer(&fn),
	)
	return
}

// RegisterComponent calls the function "WEBEXT.autotestPrivate.registerComponent" directly.
func RegisterComponent(name js.String, path js.String) (ret js.Void) {
	bindings.CallRegisterComponent(
		js.Pointer(&ret),
		name.Ref(),
		path.Ref(),
	)

	return
}

// TryRegisterComponent calls the function "WEBEXT.autotestPrivate.registerComponent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegisterComponent(name js.String, path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegisterComponent(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		path.Ref(),
	)

	return
}

// HasFuncRemoveActiveDesk returns true if the function "WEBEXT.autotestPrivate.removeActiveDesk" exists.
func HasFuncRemoveActiveDesk() bool {
	return js.True == bindings.HasFuncRemoveActiveDesk()
}

// FuncRemoveActiveDesk returns the function "WEBEXT.autotestPrivate.removeActiveDesk".
func FuncRemoveActiveDesk() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncRemoveActiveDesk(
		js.Pointer(&fn),
	)
	return
}

// RemoveActiveDesk calls the function "WEBEXT.autotestPrivate.removeActiveDesk" directly.
func RemoveActiveDesk() (ret js.Promise[js.Boolean]) {
	bindings.CallRemoveActiveDesk(
		js.Pointer(&ret),
	)

	return
}

// TryRemoveActiveDesk calls the function "WEBEXT.autotestPrivate.removeActiveDesk"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveActiveDesk() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveActiveDesk(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveAllNotifications returns true if the function "WEBEXT.autotestPrivate.removeAllNotifications" exists.
func HasFuncRemoveAllNotifications() bool {
	return js.True == bindings.HasFuncRemoveAllNotifications()
}

// FuncRemoveAllNotifications returns the function "WEBEXT.autotestPrivate.removeAllNotifications".
func FuncRemoveAllNotifications() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRemoveAllNotifications(
		js.Pointer(&fn),
	)
	return
}

// RemoveAllNotifications calls the function "WEBEXT.autotestPrivate.removeAllNotifications" directly.
func RemoveAllNotifications() (ret js.Promise[js.Void]) {
	bindings.CallRemoveAllNotifications(
		js.Pointer(&ret),
	)

	return
}

// TryRemoveAllNotifications calls the function "WEBEXT.autotestPrivate.removeAllNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveAllNotifications() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveAllNotifications(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveBruschetta returns true if the function "WEBEXT.autotestPrivate.removeBruschetta" exists.
func HasFuncRemoveBruschetta() bool {
	return js.True == bindings.HasFuncRemoveBruschetta()
}

// FuncRemoveBruschetta returns the function "WEBEXT.autotestPrivate.removeBruschetta".
func FuncRemoveBruschetta() (fn js.Func[func(vm_name js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveBruschetta(
		js.Pointer(&fn),
	)
	return
}

// RemoveBruschetta calls the function "WEBEXT.autotestPrivate.removeBruschetta" directly.
func RemoveBruschetta(vm_name js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveBruschetta(
		js.Pointer(&ret),
		vm_name.Ref(),
	)

	return
}

// TryRemoveBruschetta calls the function "WEBEXT.autotestPrivate.removeBruschetta"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveBruschetta(vm_name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveBruschetta(
		js.Pointer(&ret), js.Pointer(&exception),
		vm_name.Ref(),
	)

	return
}

// HasFuncRemoveComponentExtension returns true if the function "WEBEXT.autotestPrivate.removeComponentExtension" exists.
func HasFuncRemoveComponentExtension() bool {
	return js.True == bindings.HasFuncRemoveComponentExtension()
}

// FuncRemoveComponentExtension returns the function "WEBEXT.autotestPrivate.removeComponentExtension".
func FuncRemoveComponentExtension() (fn js.Func[func(extensionId js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveComponentExtension(
		js.Pointer(&fn),
	)
	return
}

// RemoveComponentExtension calls the function "WEBEXT.autotestPrivate.removeComponentExtension" directly.
func RemoveComponentExtension(extensionId js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveComponentExtension(
		js.Pointer(&ret),
		extensionId.Ref(),
	)

	return
}

// TryRemoveComponentExtension calls the function "WEBEXT.autotestPrivate.removeComponentExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveComponentExtension(extensionId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveComponentExtension(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
	)

	return
}

// HasFuncRemoveFuseboxTempDir returns true if the function "WEBEXT.autotestPrivate.removeFuseboxTempDir" exists.
func HasFuncRemoveFuseboxTempDir() bool {
	return js.True == bindings.HasFuncRemoveFuseboxTempDir()
}

// FuncRemoveFuseboxTempDir returns the function "WEBEXT.autotestPrivate.removeFuseboxTempDir".
func FuncRemoveFuseboxTempDir() (fn js.Func[func(fuseboxFilePath js.String) js.Promise[js.Void]]) {
	bindings.FuncRemoveFuseboxTempDir(
		js.Pointer(&fn),
	)
	return
}

// RemoveFuseboxTempDir calls the function "WEBEXT.autotestPrivate.removeFuseboxTempDir" directly.
func RemoveFuseboxTempDir(fuseboxFilePath js.String) (ret js.Promise[js.Void]) {
	bindings.CallRemoveFuseboxTempDir(
		js.Pointer(&ret),
		fuseboxFilePath.Ref(),
	)

	return
}

// TryRemoveFuseboxTempDir calls the function "WEBEXT.autotestPrivate.removeFuseboxTempDir"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveFuseboxTempDir(fuseboxFilePath js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveFuseboxTempDir(
		js.Pointer(&ret), js.Pointer(&exception),
		fuseboxFilePath.Ref(),
	)

	return
}

// HasFuncRemovePrinter returns true if the function "WEBEXT.autotestPrivate.removePrinter" exists.
func HasFuncRemovePrinter() bool {
	return js.True == bindings.HasFuncRemovePrinter()
}

// FuncRemovePrinter returns the function "WEBEXT.autotestPrivate.removePrinter".
func FuncRemovePrinter() (fn js.Func[func(printerId js.String)]) {
	bindings.FuncRemovePrinter(
		js.Pointer(&fn),
	)
	return
}

// RemovePrinter calls the function "WEBEXT.autotestPrivate.removePrinter" directly.
func RemovePrinter(printerId js.String) (ret js.Void) {
	bindings.CallRemovePrinter(
		js.Pointer(&ret),
		printerId.Ref(),
	)

	return
}

// TryRemovePrinter calls the function "WEBEXT.autotestPrivate.removePrinter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemovePrinter(printerId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemovePrinter(
		js.Pointer(&ret), js.Pointer(&exception),
		printerId.Ref(),
	)

	return
}

// HasFuncResetHoldingSpace returns true if the function "WEBEXT.autotestPrivate.resetHoldingSpace" exists.
func HasFuncResetHoldingSpace() bool {
	return js.True == bindings.HasFuncResetHoldingSpace()
}

// FuncResetHoldingSpace returns the function "WEBEXT.autotestPrivate.resetHoldingSpace".
func FuncResetHoldingSpace() (fn js.Func[func(options ResetHoldingSpaceOptions) js.Promise[js.Void]]) {
	bindings.FuncResetHoldingSpace(
		js.Pointer(&fn),
	)
	return
}

// ResetHoldingSpace calls the function "WEBEXT.autotestPrivate.resetHoldingSpace" directly.
func ResetHoldingSpace(options ResetHoldingSpaceOptions) (ret js.Promise[js.Void]) {
	bindings.CallResetHoldingSpace(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryResetHoldingSpace calls the function "WEBEXT.autotestPrivate.resetHoldingSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResetHoldingSpace(options ResetHoldingSpaceOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResetHoldingSpace(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRestart returns true if the function "WEBEXT.autotestPrivate.restart" exists.
func HasFuncRestart() bool {
	return js.True == bindings.HasFuncRestart()
}

// FuncRestart returns the function "WEBEXT.autotestPrivate.restart".
func FuncRestart() (fn js.Func[func()]) {
	bindings.FuncRestart(
		js.Pointer(&fn),
	)
	return
}

// Restart calls the function "WEBEXT.autotestPrivate.restart" directly.
func Restart() (ret js.Void) {
	bindings.CallRestart(
		js.Pointer(&ret),
	)

	return
}

// TryRestart calls the function "WEBEXT.autotestPrivate.restart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestart(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunCrostiniInstaller returns true if the function "WEBEXT.autotestPrivate.runCrostiniInstaller" exists.
func HasFuncRunCrostiniInstaller() bool {
	return js.True == bindings.HasFuncRunCrostiniInstaller()
}

// FuncRunCrostiniInstaller returns the function "WEBEXT.autotestPrivate.runCrostiniInstaller".
func FuncRunCrostiniInstaller() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRunCrostiniInstaller(
		js.Pointer(&fn),
	)
	return
}

// RunCrostiniInstaller calls the function "WEBEXT.autotestPrivate.runCrostiniInstaller" directly.
func RunCrostiniInstaller() (ret js.Promise[js.Void]) {
	bindings.CallRunCrostiniInstaller(
		js.Pointer(&ret),
	)

	return
}

// TryRunCrostiniInstaller calls the function "WEBEXT.autotestPrivate.runCrostiniInstaller"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCrostiniInstaller() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCrostiniInstaller(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunCrostiniUninstaller returns true if the function "WEBEXT.autotestPrivate.runCrostiniUninstaller" exists.
func HasFuncRunCrostiniUninstaller() bool {
	return js.True == bindings.HasFuncRunCrostiniUninstaller()
}

// FuncRunCrostiniUninstaller returns the function "WEBEXT.autotestPrivate.runCrostiniUninstaller".
func FuncRunCrostiniUninstaller() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncRunCrostiniUninstaller(
		js.Pointer(&fn),
	)
	return
}

// RunCrostiniUninstaller calls the function "WEBEXT.autotestPrivate.runCrostiniUninstaller" directly.
func RunCrostiniUninstaller() (ret js.Promise[js.Void]) {
	bindings.CallRunCrostiniUninstaller(
		js.Pointer(&ret),
	)

	return
}

// TryRunCrostiniUninstaller calls the function "WEBEXT.autotestPrivate.runCrostiniUninstaller"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRunCrostiniUninstaller() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRunCrostiniUninstaller(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendArcOverlayColor returns true if the function "WEBEXT.autotestPrivate.sendArcOverlayColor" exists.
func HasFuncSendArcOverlayColor() bool {
	return js.True == bindings.HasFuncSendArcOverlayColor()
}

// FuncSendArcOverlayColor returns the function "WEBEXT.autotestPrivate.sendArcOverlayColor".
func FuncSendArcOverlayColor() (fn js.Func[func(color int32, theme ThemeStyle) js.Promise[js.Boolean]]) {
	bindings.FuncSendArcOverlayColor(
		js.Pointer(&fn),
	)
	return
}

// SendArcOverlayColor calls the function "WEBEXT.autotestPrivate.sendArcOverlayColor" directly.
func SendArcOverlayColor(color int32, theme ThemeStyle) (ret js.Promise[js.Boolean]) {
	bindings.CallSendArcOverlayColor(
		js.Pointer(&ret),
		int32(color),
		uint32(theme),
	)

	return
}

// TrySendArcOverlayColor calls the function "WEBEXT.autotestPrivate.sendArcOverlayColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendArcOverlayColor(color int32, theme ThemeStyle) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendArcOverlayColor(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(color),
		uint32(theme),
	)

	return
}

// HasFuncSendAssistantTextQuery returns true if the function "WEBEXT.autotestPrivate.sendAssistantTextQuery" exists.
func HasFuncSendAssistantTextQuery() bool {
	return js.True == bindings.HasFuncSendAssistantTextQuery()
}

// FuncSendAssistantTextQuery returns the function "WEBEXT.autotestPrivate.sendAssistantTextQuery".
func FuncSendAssistantTextQuery() (fn js.Func[func(query js.String, timeout_ms int32) js.Promise[AssistantQueryStatus]]) {
	bindings.FuncSendAssistantTextQuery(
		js.Pointer(&fn),
	)
	return
}

// SendAssistantTextQuery calls the function "WEBEXT.autotestPrivate.sendAssistantTextQuery" directly.
func SendAssistantTextQuery(query js.String, timeout_ms int32) (ret js.Promise[AssistantQueryStatus]) {
	bindings.CallSendAssistantTextQuery(
		js.Pointer(&ret),
		query.Ref(),
		int32(timeout_ms),
	)

	return
}

// TrySendAssistantTextQuery calls the function "WEBEXT.autotestPrivate.sendAssistantTextQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendAssistantTextQuery(query js.String, timeout_ms int32) (ret js.Promise[AssistantQueryStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendAssistantTextQuery(
		js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		int32(timeout_ms),
	)

	return
}

// HasFuncSetAllowedPref returns true if the function "WEBEXT.autotestPrivate.setAllowedPref" exists.
func HasFuncSetAllowedPref() bool {
	return js.True == bindings.HasFuncSetAllowedPref()
}

// FuncSetAllowedPref returns the function "WEBEXT.autotestPrivate.setAllowedPref".
func FuncSetAllowedPref() (fn js.Func[func(pref_name js.String, value js.Any) js.Promise[js.Void]]) {
	bindings.FuncSetAllowedPref(
		js.Pointer(&fn),
	)
	return
}

// SetAllowedPref calls the function "WEBEXT.autotestPrivate.setAllowedPref" directly.
func SetAllowedPref(pref_name js.String, value js.Any) (ret js.Promise[js.Void]) {
	bindings.CallSetAllowedPref(
		js.Pointer(&ret),
		pref_name.Ref(),
		value.Ref(),
	)

	return
}

// TrySetAllowedPref calls the function "WEBEXT.autotestPrivate.setAllowedPref"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAllowedPref(pref_name js.String, value js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAllowedPref(
		js.Pointer(&ret), js.Pointer(&exception),
		pref_name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncSetAppWindowState returns true if the function "WEBEXT.autotestPrivate.setAppWindowState" exists.
func HasFuncSetAppWindowState() bool {
	return js.True == bindings.HasFuncSetAppWindowState()
}

// FuncSetAppWindowState returns the function "WEBEXT.autotestPrivate.setAppWindowState".
func FuncSetAppWindowState() (fn js.Func[func(id int32, change WindowStateChangeDict, wait bool) js.Promise[WindowStateType]]) {
	bindings.FuncSetAppWindowState(
		js.Pointer(&fn),
	)
	return
}

// SetAppWindowState calls the function "WEBEXT.autotestPrivate.setAppWindowState" directly.
func SetAppWindowState(id int32, change WindowStateChangeDict, wait bool) (ret js.Promise[WindowStateType]) {
	bindings.CallSetAppWindowState(
		js.Pointer(&ret),
		int32(id),
		js.Pointer(&change),
		js.Bool(bool(wait)),
	)

	return
}

// TrySetAppWindowState calls the function "WEBEXT.autotestPrivate.setAppWindowState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAppWindowState(id int32, change WindowStateChangeDict, wait bool) (ret js.Promise[WindowStateType], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAppWindowState(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
		js.Pointer(&change),
		js.Bool(bool(wait)),
	)

	return
}

// HasFuncSetArcAppWindowFocus returns true if the function "WEBEXT.autotestPrivate.setArcAppWindowFocus" exists.
func HasFuncSetArcAppWindowFocus() bool {
	return js.True == bindings.HasFuncSetArcAppWindowFocus()
}

// FuncSetArcAppWindowFocus returns the function "WEBEXT.autotestPrivate.setArcAppWindowFocus".
func FuncSetArcAppWindowFocus() (fn js.Func[func(packageName js.String) js.Promise[js.Void]]) {
	bindings.FuncSetArcAppWindowFocus(
		js.Pointer(&fn),
	)
	return
}

// SetArcAppWindowFocus calls the function "WEBEXT.autotestPrivate.setArcAppWindowFocus" directly.
func SetArcAppWindowFocus(packageName js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetArcAppWindowFocus(
		js.Pointer(&ret),
		packageName.Ref(),
	)

	return
}

// TrySetArcAppWindowFocus calls the function "WEBEXT.autotestPrivate.setArcAppWindowFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetArcAppWindowFocus(packageName js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetArcAppWindowFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		packageName.Ref(),
	)

	return
}

// HasFuncSetArcInteractiveState returns true if the function "WEBEXT.autotestPrivate.setArcInteractiveState" exists.
func HasFuncSetArcInteractiveState() bool {
	return js.True == bindings.HasFuncSetArcInteractiveState()
}

// FuncSetArcInteractiveState returns the function "WEBEXT.autotestPrivate.setArcInteractiveState".
func FuncSetArcInteractiveState() (fn js.Func[func(enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetArcInteractiveState(
		js.Pointer(&fn),
	)
	return
}

// SetArcInteractiveState calls the function "WEBEXT.autotestPrivate.setArcInteractiveState" directly.
func SetArcInteractiveState(enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetArcInteractiveState(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetArcInteractiveState calls the function "WEBEXT.autotestPrivate.setArcInteractiveState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetArcInteractiveState(enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetArcInteractiveState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetArcTouchMode returns true if the function "WEBEXT.autotestPrivate.setArcTouchMode" exists.
func HasFuncSetArcTouchMode() bool {
	return js.True == bindings.HasFuncSetArcTouchMode()
}

// FuncSetArcTouchMode returns the function "WEBEXT.autotestPrivate.setArcTouchMode".
func FuncSetArcTouchMode() (fn js.Func[func(enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetArcTouchMode(
		js.Pointer(&fn),
	)
	return
}

// SetArcTouchMode calls the function "WEBEXT.autotestPrivate.setArcTouchMode" directly.
func SetArcTouchMode(enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetArcTouchMode(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetArcTouchMode calls the function "WEBEXT.autotestPrivate.setArcTouchMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetArcTouchMode(enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetArcTouchMode(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetAssistantEnabled returns true if the function "WEBEXT.autotestPrivate.setAssistantEnabled" exists.
func HasFuncSetAssistantEnabled() bool {
	return js.True == bindings.HasFuncSetAssistantEnabled()
}

// FuncSetAssistantEnabled returns the function "WEBEXT.autotestPrivate.setAssistantEnabled".
func FuncSetAssistantEnabled() (fn js.Func[func(enabled bool, timeout_ms int32) js.Promise[js.Void]]) {
	bindings.FuncSetAssistantEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetAssistantEnabled calls the function "WEBEXT.autotestPrivate.setAssistantEnabled" directly.
func SetAssistantEnabled(enabled bool, timeout_ms int32) (ret js.Promise[js.Void]) {
	bindings.CallSetAssistantEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
		int32(timeout_ms),
	)

	return
}

// TrySetAssistantEnabled calls the function "WEBEXT.autotestPrivate.setAssistantEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAssistantEnabled(enabled bool, timeout_ms int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAssistantEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
		int32(timeout_ms),
	)

	return
}

// HasFuncSetClipboardTextData returns true if the function "WEBEXT.autotestPrivate.setClipboardTextData" exists.
func HasFuncSetClipboardTextData() bool {
	return js.True == bindings.HasFuncSetClipboardTextData()
}

// FuncSetClipboardTextData returns the function "WEBEXT.autotestPrivate.setClipboardTextData".
func FuncSetClipboardTextData() (fn js.Func[func(data js.String) js.Promise[js.Void]]) {
	bindings.FuncSetClipboardTextData(
		js.Pointer(&fn),
	)
	return
}

// SetClipboardTextData calls the function "WEBEXT.autotestPrivate.setClipboardTextData" directly.
func SetClipboardTextData(data js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetClipboardTextData(
		js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySetClipboardTextData calls the function "WEBEXT.autotestPrivate.setClipboardTextData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetClipboardTextData(data js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetClipboardTextData(
		js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSetCrostiniAppScaled returns true if the function "WEBEXT.autotestPrivate.setCrostiniAppScaled" exists.
func HasFuncSetCrostiniAppScaled() bool {
	return js.True == bindings.HasFuncSetCrostiniAppScaled()
}

// FuncSetCrostiniAppScaled returns the function "WEBEXT.autotestPrivate.setCrostiniAppScaled".
func FuncSetCrostiniAppScaled() (fn js.Func[func(appId js.String, scaled bool) js.Promise[js.Void]]) {
	bindings.FuncSetCrostiniAppScaled(
		js.Pointer(&fn),
	)
	return
}

// SetCrostiniAppScaled calls the function "WEBEXT.autotestPrivate.setCrostiniAppScaled" directly.
func SetCrostiniAppScaled(appId js.String, scaled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetCrostiniAppScaled(
		js.Pointer(&ret),
		appId.Ref(),
		js.Bool(bool(scaled)),
	)

	return
}

// TrySetCrostiniAppScaled calls the function "WEBEXT.autotestPrivate.setCrostiniAppScaled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCrostiniAppScaled(appId js.String, scaled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCrostiniAppScaled(
		js.Pointer(&ret), js.Pointer(&exception),
		appId.Ref(),
		js.Bool(bool(scaled)),
	)

	return
}

// HasFuncSetCrostiniEnabled returns true if the function "WEBEXT.autotestPrivate.setCrostiniEnabled" exists.
func HasFuncSetCrostiniEnabled() bool {
	return js.True == bindings.HasFuncSetCrostiniEnabled()
}

// FuncSetCrostiniEnabled returns the function "WEBEXT.autotestPrivate.setCrostiniEnabled".
func FuncSetCrostiniEnabled() (fn js.Func[func(enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetCrostiniEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetCrostiniEnabled calls the function "WEBEXT.autotestPrivate.setCrostiniEnabled" directly.
func SetCrostiniEnabled(enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetCrostiniEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetCrostiniEnabled calls the function "WEBEXT.autotestPrivate.setCrostiniEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCrostiniEnabled(enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCrostiniEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetMetricsEnabled returns true if the function "WEBEXT.autotestPrivate.setMetricsEnabled" exists.
func HasFuncSetMetricsEnabled() bool {
	return js.True == bindings.HasFuncSetMetricsEnabled()
}

// FuncSetMetricsEnabled returns the function "WEBEXT.autotestPrivate.setMetricsEnabled".
func FuncSetMetricsEnabled() (fn js.Func[func(enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetMetricsEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetMetricsEnabled calls the function "WEBEXT.autotestPrivate.setMetricsEnabled" directly.
func SetMetricsEnabled(enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetMetricsEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetMetricsEnabled calls the function "WEBEXT.autotestPrivate.setMetricsEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMetricsEnabled(enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMetricsEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetMouseReverseScroll returns true if the function "WEBEXT.autotestPrivate.setMouseReverseScroll" exists.
func HasFuncSetMouseReverseScroll() bool {
	return js.True == bindings.HasFuncSetMouseReverseScroll()
}

// FuncSetMouseReverseScroll returns the function "WEBEXT.autotestPrivate.setMouseReverseScroll".
func FuncSetMouseReverseScroll() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetMouseReverseScroll(
		js.Pointer(&fn),
	)
	return
}

// SetMouseReverseScroll calls the function "WEBEXT.autotestPrivate.setMouseReverseScroll" directly.
func SetMouseReverseScroll(enabled bool) (ret js.Void) {
	bindings.CallSetMouseReverseScroll(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetMouseReverseScroll calls the function "WEBEXT.autotestPrivate.setMouseReverseScroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMouseReverseScroll(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMouseReverseScroll(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetMouseSensitivity returns true if the function "WEBEXT.autotestPrivate.setMouseSensitivity" exists.
func HasFuncSetMouseSensitivity() bool {
	return js.True == bindings.HasFuncSetMouseSensitivity()
}

// FuncSetMouseSensitivity returns the function "WEBEXT.autotestPrivate.setMouseSensitivity".
func FuncSetMouseSensitivity() (fn js.Func[func(value int32)]) {
	bindings.FuncSetMouseSensitivity(
		js.Pointer(&fn),
	)
	return
}

// SetMouseSensitivity calls the function "WEBEXT.autotestPrivate.setMouseSensitivity" directly.
func SetMouseSensitivity(value int32) (ret js.Void) {
	bindings.CallSetMouseSensitivity(
		js.Pointer(&ret),
		int32(value),
	)

	return
}

// TrySetMouseSensitivity calls the function "WEBEXT.autotestPrivate.setMouseSensitivity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMouseSensitivity(value int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMouseSensitivity(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(value),
	)

	return
}

// HasFuncSetNaturalScroll returns true if the function "WEBEXT.autotestPrivate.setNaturalScroll" exists.
func HasFuncSetNaturalScroll() bool {
	return js.True == bindings.HasFuncSetNaturalScroll()
}

// FuncSetNaturalScroll returns the function "WEBEXT.autotestPrivate.setNaturalScroll".
func FuncSetNaturalScroll() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetNaturalScroll(
		js.Pointer(&fn),
	)
	return
}

// SetNaturalScroll calls the function "WEBEXT.autotestPrivate.setNaturalScroll" directly.
func SetNaturalScroll(enabled bool) (ret js.Void) {
	bindings.CallSetNaturalScroll(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetNaturalScroll calls the function "WEBEXT.autotestPrivate.setNaturalScroll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetNaturalScroll(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetNaturalScroll(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetOverviewModeState returns true if the function "WEBEXT.autotestPrivate.setOverviewModeState" exists.
func HasFuncSetOverviewModeState() bool {
	return js.True == bindings.HasFuncSetOverviewModeState()
}

// FuncSetOverviewModeState returns the function "WEBEXT.autotestPrivate.setOverviewModeState".
func FuncSetOverviewModeState() (fn js.Func[func(start bool) js.Promise[js.Boolean]]) {
	bindings.FuncSetOverviewModeState(
		js.Pointer(&fn),
	)
	return
}

// SetOverviewModeState calls the function "WEBEXT.autotestPrivate.setOverviewModeState" directly.
func SetOverviewModeState(start bool) (ret js.Promise[js.Boolean]) {
	bindings.CallSetOverviewModeState(
		js.Pointer(&ret),
		js.Bool(bool(start)),
	)

	return
}

// TrySetOverviewModeState calls the function "WEBEXT.autotestPrivate.setOverviewModeState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetOverviewModeState(start bool) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetOverviewModeState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(start)),
	)

	return
}

// HasFuncSetPlayStoreEnabled returns true if the function "WEBEXT.autotestPrivate.setPlayStoreEnabled" exists.
func HasFuncSetPlayStoreEnabled() bool {
	return js.True == bindings.HasFuncSetPlayStoreEnabled()
}

// FuncSetPlayStoreEnabled returns the function "WEBEXT.autotestPrivate.setPlayStoreEnabled".
func FuncSetPlayStoreEnabled() (fn js.Func[func(enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetPlayStoreEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetPlayStoreEnabled calls the function "WEBEXT.autotestPrivate.setPlayStoreEnabled" directly.
func SetPlayStoreEnabled(enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetPlayStoreEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetPlayStoreEnabled calls the function "WEBEXT.autotestPrivate.setPlayStoreEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPlayStoreEnabled(enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPlayStoreEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetPluginVMPolicy returns true if the function "WEBEXT.autotestPrivate.setPluginVMPolicy" exists.
func HasFuncSetPluginVMPolicy() bool {
	return js.True == bindings.HasFuncSetPluginVMPolicy()
}

// FuncSetPluginVMPolicy returns the function "WEBEXT.autotestPrivate.setPluginVMPolicy".
func FuncSetPluginVMPolicy() (fn js.Func[func(imageUrl js.String, imageHash js.String, licenseKey js.String)]) {
	bindings.FuncSetPluginVMPolicy(
		js.Pointer(&fn),
	)
	return
}

// SetPluginVMPolicy calls the function "WEBEXT.autotestPrivate.setPluginVMPolicy" directly.
func SetPluginVMPolicy(imageUrl js.String, imageHash js.String, licenseKey js.String) (ret js.Void) {
	bindings.CallSetPluginVMPolicy(
		js.Pointer(&ret),
		imageUrl.Ref(),
		imageHash.Ref(),
		licenseKey.Ref(),
	)

	return
}

// TrySetPluginVMPolicy calls the function "WEBEXT.autotestPrivate.setPluginVMPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPluginVMPolicy(imageUrl js.String, imageHash js.String, licenseKey js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPluginVMPolicy(
		js.Pointer(&ret), js.Pointer(&exception),
		imageUrl.Ref(),
		imageHash.Ref(),
		licenseKey.Ref(),
	)

	return
}

// HasFuncSetPrimaryButtonRight returns true if the function "WEBEXT.autotestPrivate.setPrimaryButtonRight" exists.
func HasFuncSetPrimaryButtonRight() bool {
	return js.True == bindings.HasFuncSetPrimaryButtonRight()
}

// FuncSetPrimaryButtonRight returns the function "WEBEXT.autotestPrivate.setPrimaryButtonRight".
func FuncSetPrimaryButtonRight() (fn js.Func[func(right bool)]) {
	bindings.FuncSetPrimaryButtonRight(
		js.Pointer(&fn),
	)
	return
}

// SetPrimaryButtonRight calls the function "WEBEXT.autotestPrivate.setPrimaryButtonRight" directly.
func SetPrimaryButtonRight(right bool) (ret js.Void) {
	bindings.CallSetPrimaryButtonRight(
		js.Pointer(&ret),
		js.Bool(bool(right)),
	)

	return
}

// TrySetPrimaryButtonRight calls the function "WEBEXT.autotestPrivate.setPrimaryButtonRight"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPrimaryButtonRight(right bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPrimaryButtonRight(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(right)),
	)

	return
}

// HasFuncSetShelfAlignment returns true if the function "WEBEXT.autotestPrivate.setShelfAlignment" exists.
func HasFuncSetShelfAlignment() bool {
	return js.True == bindings.HasFuncSetShelfAlignment()
}

// FuncSetShelfAlignment returns the function "WEBEXT.autotestPrivate.setShelfAlignment".
func FuncSetShelfAlignment() (fn js.Func[func(displayId js.String, alignment ShelfAlignmentType) js.Promise[js.Void]]) {
	bindings.FuncSetShelfAlignment(
		js.Pointer(&fn),
	)
	return
}

// SetShelfAlignment calls the function "WEBEXT.autotestPrivate.setShelfAlignment" directly.
func SetShelfAlignment(displayId js.String, alignment ShelfAlignmentType) (ret js.Promise[js.Void]) {
	bindings.CallSetShelfAlignment(
		js.Pointer(&ret),
		displayId.Ref(),
		uint32(alignment),
	)

	return
}

// TrySetShelfAlignment calls the function "WEBEXT.autotestPrivate.setShelfAlignment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShelfAlignment(displayId js.String, alignment ShelfAlignmentType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShelfAlignment(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
		uint32(alignment),
	)

	return
}

// HasFuncSetShelfAutoHideBehavior returns true if the function "WEBEXT.autotestPrivate.setShelfAutoHideBehavior" exists.
func HasFuncSetShelfAutoHideBehavior() bool {
	return js.True == bindings.HasFuncSetShelfAutoHideBehavior()
}

// FuncSetShelfAutoHideBehavior returns the function "WEBEXT.autotestPrivate.setShelfAutoHideBehavior".
func FuncSetShelfAutoHideBehavior() (fn js.Func[func(displayId js.String, behavior js.String) js.Promise[js.Void]]) {
	bindings.FuncSetShelfAutoHideBehavior(
		js.Pointer(&fn),
	)
	return
}

// SetShelfAutoHideBehavior calls the function "WEBEXT.autotestPrivate.setShelfAutoHideBehavior" directly.
func SetShelfAutoHideBehavior(displayId js.String, behavior js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetShelfAutoHideBehavior(
		js.Pointer(&ret),
		displayId.Ref(),
		behavior.Ref(),
	)

	return
}

// TrySetShelfAutoHideBehavior calls the function "WEBEXT.autotestPrivate.setShelfAutoHideBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShelfAutoHideBehavior(displayId js.String, behavior js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShelfAutoHideBehavior(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
		behavior.Ref(),
	)

	return
}

// HasFuncSetShelfIconPin returns true if the function "WEBEXT.autotestPrivate.setShelfIconPin" exists.
func HasFuncSetShelfIconPin() bool {
	return js.True == bindings.HasFuncSetShelfIconPin()
}

// FuncSetShelfIconPin returns the function "WEBEXT.autotestPrivate.setShelfIconPin".
func FuncSetShelfIconPin() (fn js.Func[func(updateParams js.Array[ShelfIconPinUpdateParam]) js.Promise[js.Array[js.String]]]) {
	bindings.FuncSetShelfIconPin(
		js.Pointer(&fn),
	)
	return
}

// SetShelfIconPin calls the function "WEBEXT.autotestPrivate.setShelfIconPin" directly.
func SetShelfIconPin(updateParams js.Array[ShelfIconPinUpdateParam]) (ret js.Promise[js.Array[js.String]]) {
	bindings.CallSetShelfIconPin(
		js.Pointer(&ret),
		updateParams.Ref(),
	)

	return
}

// TrySetShelfIconPin calls the function "WEBEXT.autotestPrivate.setShelfIconPin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShelfIconPin(updateParams js.Array[ShelfIconPinUpdateParam]) (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShelfIconPin(
		js.Pointer(&ret), js.Pointer(&exception),
		updateParams.Ref(),
	)

	return
}

// HasFuncSetTabletModeEnabled returns true if the function "WEBEXT.autotestPrivate.setTabletModeEnabled" exists.
func HasFuncSetTabletModeEnabled() bool {
	return js.True == bindings.HasFuncSetTabletModeEnabled()
}

// FuncSetTabletModeEnabled returns the function "WEBEXT.autotestPrivate.setTabletModeEnabled".
func FuncSetTabletModeEnabled() (fn js.Func[func(enabled bool) js.Promise[js.Boolean]]) {
	bindings.FuncSetTabletModeEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetTabletModeEnabled calls the function "WEBEXT.autotestPrivate.setTabletModeEnabled" directly.
func SetTabletModeEnabled(enabled bool) (ret js.Promise[js.Boolean]) {
	bindings.CallSetTabletModeEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetTabletModeEnabled calls the function "WEBEXT.autotestPrivate.setTabletModeEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTabletModeEnabled(enabled bool) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTabletModeEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetTapDragging returns true if the function "WEBEXT.autotestPrivate.setTapDragging" exists.
func HasFuncSetTapDragging() bool {
	return js.True == bindings.HasFuncSetTapDragging()
}

// FuncSetTapDragging returns the function "WEBEXT.autotestPrivate.setTapDragging".
func FuncSetTapDragging() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetTapDragging(
		js.Pointer(&fn),
	)
	return
}

// SetTapDragging calls the function "WEBEXT.autotestPrivate.setTapDragging" directly.
func SetTapDragging(enabled bool) (ret js.Void) {
	bindings.CallSetTapDragging(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetTapDragging calls the function "WEBEXT.autotestPrivate.setTapDragging"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTapDragging(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTapDragging(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetTapToClick returns true if the function "WEBEXT.autotestPrivate.setTapToClick" exists.
func HasFuncSetTapToClick() bool {
	return js.True == bindings.HasFuncSetTapToClick()
}

// FuncSetTapToClick returns the function "WEBEXT.autotestPrivate.setTapToClick".
func FuncSetTapToClick() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetTapToClick(
		js.Pointer(&fn),
	)
	return
}

// SetTapToClick calls the function "WEBEXT.autotestPrivate.setTapToClick" directly.
func SetTapToClick(enabled bool) (ret js.Void) {
	bindings.CallSetTapToClick(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetTapToClick calls the function "WEBEXT.autotestPrivate.setTapToClick"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTapToClick(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTapToClick(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetThreeFingerClick returns true if the function "WEBEXT.autotestPrivate.setThreeFingerClick" exists.
func HasFuncSetThreeFingerClick() bool {
	return js.True == bindings.HasFuncSetThreeFingerClick()
}

// FuncSetThreeFingerClick returns the function "WEBEXT.autotestPrivate.setThreeFingerClick".
func FuncSetThreeFingerClick() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetThreeFingerClick(
		js.Pointer(&fn),
	)
	return
}

// SetThreeFingerClick calls the function "WEBEXT.autotestPrivate.setThreeFingerClick" directly.
func SetThreeFingerClick(enabled bool) (ret js.Void) {
	bindings.CallSetThreeFingerClick(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetThreeFingerClick calls the function "WEBEXT.autotestPrivate.setThreeFingerClick"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetThreeFingerClick(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetThreeFingerClick(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetTouchpadSensitivity returns true if the function "WEBEXT.autotestPrivate.setTouchpadSensitivity" exists.
func HasFuncSetTouchpadSensitivity() bool {
	return js.True == bindings.HasFuncSetTouchpadSensitivity()
}

// FuncSetTouchpadSensitivity returns the function "WEBEXT.autotestPrivate.setTouchpadSensitivity".
func FuncSetTouchpadSensitivity() (fn js.Func[func(value int32)]) {
	bindings.FuncSetTouchpadSensitivity(
		js.Pointer(&fn),
	)
	return
}

// SetTouchpadSensitivity calls the function "WEBEXT.autotestPrivate.setTouchpadSensitivity" directly.
func SetTouchpadSensitivity(value int32) (ret js.Void) {
	bindings.CallSetTouchpadSensitivity(
		js.Pointer(&ret),
		int32(value),
	)

	return
}

// TrySetTouchpadSensitivity calls the function "WEBEXT.autotestPrivate.setTouchpadSensitivity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetTouchpadSensitivity(value int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetTouchpadSensitivity(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(value),
	)

	return
}

// HasFuncSetWhitelistedPref returns true if the function "WEBEXT.autotestPrivate.setWhitelistedPref" exists.
func HasFuncSetWhitelistedPref() bool {
	return js.True == bindings.HasFuncSetWhitelistedPref()
}

// FuncSetWhitelistedPref returns the function "WEBEXT.autotestPrivate.setWhitelistedPref".
func FuncSetWhitelistedPref() (fn js.Func[func(pref_name js.String, value js.Any) js.Promise[js.Void]]) {
	bindings.FuncSetWhitelistedPref(
		js.Pointer(&fn),
	)
	return
}

// SetWhitelistedPref calls the function "WEBEXT.autotestPrivate.setWhitelistedPref" directly.
func SetWhitelistedPref(pref_name js.String, value js.Any) (ret js.Promise[js.Void]) {
	bindings.CallSetWhitelistedPref(
		js.Pointer(&ret),
		pref_name.Ref(),
		value.Ref(),
	)

	return
}

// TrySetWhitelistedPref calls the function "WEBEXT.autotestPrivate.setWhitelistedPref"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetWhitelistedPref(pref_name js.String, value js.Any) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetWhitelistedPref(
		js.Pointer(&ret), js.Pointer(&exception),
		pref_name.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncSetWindowBounds returns true if the function "WEBEXT.autotestPrivate.setWindowBounds" exists.
func HasFuncSetWindowBounds() bool {
	return js.True == bindings.HasFuncSetWindowBounds()
}

// FuncSetWindowBounds returns the function "WEBEXT.autotestPrivate.setWindowBounds".
func FuncSetWindowBounds() (fn js.Func[func(id int32, bounds Bounds, displayId js.String) js.Promise[SetWindowBoundsResult]]) {
	bindings.FuncSetWindowBounds(
		js.Pointer(&fn),
	)
	return
}

// SetWindowBounds calls the function "WEBEXT.autotestPrivate.setWindowBounds" directly.
func SetWindowBounds(id int32, bounds Bounds, displayId js.String) (ret js.Promise[SetWindowBoundsResult]) {
	bindings.CallSetWindowBounds(
		js.Pointer(&ret),
		int32(id),
		js.Pointer(&bounds),
		displayId.Ref(),
	)

	return
}

// TrySetWindowBounds calls the function "WEBEXT.autotestPrivate.setWindowBounds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetWindowBounds(id int32, bounds Bounds, displayId js.String) (ret js.Promise[SetWindowBoundsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetWindowBounds(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
		js.Pointer(&bounds),
		displayId.Ref(),
	)

	return
}

// HasFuncShowPluginVMInstaller returns true if the function "WEBEXT.autotestPrivate.showPluginVMInstaller" exists.
func HasFuncShowPluginVMInstaller() bool {
	return js.True == bindings.HasFuncShowPluginVMInstaller()
}

// FuncShowPluginVMInstaller returns the function "WEBEXT.autotestPrivate.showPluginVMInstaller".
func FuncShowPluginVMInstaller() (fn js.Func[func()]) {
	bindings.FuncShowPluginVMInstaller(
		js.Pointer(&fn),
	)
	return
}

// ShowPluginVMInstaller calls the function "WEBEXT.autotestPrivate.showPluginVMInstaller" directly.
func ShowPluginVMInstaller() (ret js.Void) {
	bindings.CallShowPluginVMInstaller(
		js.Pointer(&ret),
	)

	return
}

// TryShowPluginVMInstaller calls the function "WEBEXT.autotestPrivate.showPluginVMInstaller"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowPluginVMInstaller() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowPluginVMInstaller(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowVirtualKeyboardIfEnabled returns true if the function "WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled" exists.
func HasFuncShowVirtualKeyboardIfEnabled() bool {
	return js.True == bindings.HasFuncShowVirtualKeyboardIfEnabled()
}

// FuncShowVirtualKeyboardIfEnabled returns the function "WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled".
func FuncShowVirtualKeyboardIfEnabled() (fn js.Func[func()]) {
	bindings.FuncShowVirtualKeyboardIfEnabled(
		js.Pointer(&fn),
	)
	return
}

// ShowVirtualKeyboardIfEnabled calls the function "WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled" directly.
func ShowVirtualKeyboardIfEnabled() (ret js.Void) {
	bindings.CallShowVirtualKeyboardIfEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryShowVirtualKeyboardIfEnabled calls the function "WEBEXT.autotestPrivate.showVirtualKeyboardIfEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowVirtualKeyboardIfEnabled() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowVirtualKeyboardIfEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShutdown returns true if the function "WEBEXT.autotestPrivate.shutdown" exists.
func HasFuncShutdown() bool {
	return js.True == bindings.HasFuncShutdown()
}

// FuncShutdown returns the function "WEBEXT.autotestPrivate.shutdown".
func FuncShutdown() (fn js.Func[func(force bool)]) {
	bindings.FuncShutdown(
		js.Pointer(&fn),
	)
	return
}

// Shutdown calls the function "WEBEXT.autotestPrivate.shutdown" directly.
func Shutdown(force bool) (ret js.Void) {
	bindings.CallShutdown(
		js.Pointer(&ret),
		js.Bool(bool(force)),
	)

	return
}

// TryShutdown calls the function "WEBEXT.autotestPrivate.shutdown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShutdown(force bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShutdown(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(force)),
	)

	return
}

// HasFuncSimulateAsanMemoryBug returns true if the function "WEBEXT.autotestPrivate.simulateAsanMemoryBug" exists.
func HasFuncSimulateAsanMemoryBug() bool {
	return js.True == bindings.HasFuncSimulateAsanMemoryBug()
}

// FuncSimulateAsanMemoryBug returns the function "WEBEXT.autotestPrivate.simulateAsanMemoryBug".
func FuncSimulateAsanMemoryBug() (fn js.Func[func()]) {
	bindings.FuncSimulateAsanMemoryBug(
		js.Pointer(&fn),
	)
	return
}

// SimulateAsanMemoryBug calls the function "WEBEXT.autotestPrivate.simulateAsanMemoryBug" directly.
func SimulateAsanMemoryBug() (ret js.Void) {
	bindings.CallSimulateAsanMemoryBug(
		js.Pointer(&ret),
	)

	return
}

// TrySimulateAsanMemoryBug calls the function "WEBEXT.autotestPrivate.simulateAsanMemoryBug"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySimulateAsanMemoryBug() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySimulateAsanMemoryBug(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStartArc returns true if the function "WEBEXT.autotestPrivate.startArc" exists.
func HasFuncStartArc() bool {
	return js.True == bindings.HasFuncStartArc()
}

// FuncStartArc returns the function "WEBEXT.autotestPrivate.startArc".
func FuncStartArc() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStartArc(
		js.Pointer(&fn),
	)
	return
}

// StartArc calls the function "WEBEXT.autotestPrivate.startArc" directly.
func StartArc() (ret js.Promise[js.Void]) {
	bindings.CallStartArc(
		js.Pointer(&ret),
	)

	return
}

// TryStartArc calls the function "WEBEXT.autotestPrivate.startArc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartArc() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartArc(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStartFrameCounting returns true if the function "WEBEXT.autotestPrivate.startFrameCounting" exists.
func HasFuncStartFrameCounting() bool {
	return js.True == bindings.HasFuncStartFrameCounting()
}

// FuncStartFrameCounting returns the function "WEBEXT.autotestPrivate.startFrameCounting".
func FuncStartFrameCounting() (fn js.Func[func(bucketSizeInSeconds int32) js.Promise[js.Void]]) {
	bindings.FuncStartFrameCounting(
		js.Pointer(&fn),
	)
	return
}

// StartFrameCounting calls the function "WEBEXT.autotestPrivate.startFrameCounting" directly.
func StartFrameCounting(bucketSizeInSeconds int32) (ret js.Promise[js.Void]) {
	bindings.CallStartFrameCounting(
		js.Pointer(&ret),
		int32(bucketSizeInSeconds),
	)

	return
}

// TryStartFrameCounting calls the function "WEBEXT.autotestPrivate.startFrameCounting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartFrameCounting(bucketSizeInSeconds int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartFrameCounting(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(bucketSizeInSeconds),
	)

	return
}

// HasFuncStartLoginEventRecorderDataCollection returns true if the function "WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection" exists.
func HasFuncStartLoginEventRecorderDataCollection() bool {
	return js.True == bindings.HasFuncStartLoginEventRecorderDataCollection()
}

// FuncStartLoginEventRecorderDataCollection returns the function "WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection".
func FuncStartLoginEventRecorderDataCollection() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStartLoginEventRecorderDataCollection(
		js.Pointer(&fn),
	)
	return
}

// StartLoginEventRecorderDataCollection calls the function "WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection" directly.
func StartLoginEventRecorderDataCollection() (ret js.Promise[js.Void]) {
	bindings.CallStartLoginEventRecorderDataCollection(
		js.Pointer(&ret),
	)

	return
}

// TryStartLoginEventRecorderDataCollection calls the function "WEBEXT.autotestPrivate.startLoginEventRecorderDataCollection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartLoginEventRecorderDataCollection() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartLoginEventRecorderDataCollection(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStartSmoothnessTracking returns true if the function "WEBEXT.autotestPrivate.startSmoothnessTracking" exists.
func HasFuncStartSmoothnessTracking() bool {
	return js.True == bindings.HasFuncStartSmoothnessTracking()
}

// FuncStartSmoothnessTracking returns the function "WEBEXT.autotestPrivate.startSmoothnessTracking".
func FuncStartSmoothnessTracking() (fn js.Func[func(displayId js.String, throughputIntervalMs int32) js.Promise[js.Void]]) {
	bindings.FuncStartSmoothnessTracking(
		js.Pointer(&fn),
	)
	return
}

// StartSmoothnessTracking calls the function "WEBEXT.autotestPrivate.startSmoothnessTracking" directly.
func StartSmoothnessTracking(displayId js.String, throughputIntervalMs int32) (ret js.Promise[js.Void]) {
	bindings.CallStartSmoothnessTracking(
		js.Pointer(&ret),
		displayId.Ref(),
		int32(throughputIntervalMs),
	)

	return
}

// TryStartSmoothnessTracking calls the function "WEBEXT.autotestPrivate.startSmoothnessTracking"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartSmoothnessTracking(displayId js.String, throughputIntervalMs int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartSmoothnessTracking(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
		int32(throughputIntervalMs),
	)

	return
}

// HasFuncStartThroughputTrackerDataCollection returns true if the function "WEBEXT.autotestPrivate.startThroughputTrackerDataCollection" exists.
func HasFuncStartThroughputTrackerDataCollection() bool {
	return js.True == bindings.HasFuncStartThroughputTrackerDataCollection()
}

// FuncStartThroughputTrackerDataCollection returns the function "WEBEXT.autotestPrivate.startThroughputTrackerDataCollection".
func FuncStartThroughputTrackerDataCollection() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStartThroughputTrackerDataCollection(
		js.Pointer(&fn),
	)
	return
}

// StartThroughputTrackerDataCollection calls the function "WEBEXT.autotestPrivate.startThroughputTrackerDataCollection" directly.
func StartThroughputTrackerDataCollection() (ret js.Promise[js.Void]) {
	bindings.CallStartThroughputTrackerDataCollection(
		js.Pointer(&ret),
	)

	return
}

// TryStartThroughputTrackerDataCollection calls the function "WEBEXT.autotestPrivate.startThroughputTrackerDataCollection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartThroughputTrackerDataCollection() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartThroughputTrackerDataCollection(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopArc returns true if the function "WEBEXT.autotestPrivate.stopArc" exists.
func HasFuncStopArc() bool {
	return js.True == bindings.HasFuncStopArc()
}

// FuncStopArc returns the function "WEBEXT.autotestPrivate.stopArc".
func FuncStopArc() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStopArc(
		js.Pointer(&fn),
	)
	return
}

// StopArc calls the function "WEBEXT.autotestPrivate.stopArc" directly.
func StopArc() (ret js.Promise[js.Void]) {
	bindings.CallStopArc(
		js.Pointer(&ret),
	)

	return
}

// TryStopArc calls the function "WEBEXT.autotestPrivate.stopArc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopArc() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopArc(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopFrameCounting returns true if the function "WEBEXT.autotestPrivate.stopFrameCounting" exists.
func HasFuncStopFrameCounting() bool {
	return js.True == bindings.HasFuncStopFrameCounting()
}

// FuncStopFrameCounting returns the function "WEBEXT.autotestPrivate.stopFrameCounting".
func FuncStopFrameCounting() (fn js.Func[func() js.Promise[js.Array[FrameCountingPerSinkData]]]) {
	bindings.FuncStopFrameCounting(
		js.Pointer(&fn),
	)
	return
}

// StopFrameCounting calls the function "WEBEXT.autotestPrivate.stopFrameCounting" directly.
func StopFrameCounting() (ret js.Promise[js.Array[FrameCountingPerSinkData]]) {
	bindings.CallStopFrameCounting(
		js.Pointer(&ret),
	)

	return
}

// TryStopFrameCounting calls the function "WEBEXT.autotestPrivate.stopFrameCounting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopFrameCounting() (ret js.Promise[js.Array[FrameCountingPerSinkData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopFrameCounting(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopSmoothnessTracking returns true if the function "WEBEXT.autotestPrivate.stopSmoothnessTracking" exists.
func HasFuncStopSmoothnessTracking() bool {
	return js.True == bindings.HasFuncStopSmoothnessTracking()
}

// FuncStopSmoothnessTracking returns the function "WEBEXT.autotestPrivate.stopSmoothnessTracking".
func FuncStopSmoothnessTracking() (fn js.Func[func(displayId js.String) js.Promise[DisplaySmoothnessData]]) {
	bindings.FuncStopSmoothnessTracking(
		js.Pointer(&fn),
	)
	return
}

// StopSmoothnessTracking calls the function "WEBEXT.autotestPrivate.stopSmoothnessTracking" directly.
func StopSmoothnessTracking(displayId js.String) (ret js.Promise[DisplaySmoothnessData]) {
	bindings.CallStopSmoothnessTracking(
		js.Pointer(&ret),
		displayId.Ref(),
	)

	return
}

// TryStopSmoothnessTracking calls the function "WEBEXT.autotestPrivate.stopSmoothnessTracking"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopSmoothnessTracking(displayId js.String) (ret js.Promise[DisplaySmoothnessData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopSmoothnessTracking(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
	)

	return
}

// HasFuncStopThroughputTrackerDataCollection returns true if the function "WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection" exists.
func HasFuncStopThroughputTrackerDataCollection() bool {
	return js.True == bindings.HasFuncStopThroughputTrackerDataCollection()
}

// FuncStopThroughputTrackerDataCollection returns the function "WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection".
func FuncStopThroughputTrackerDataCollection() (fn js.Func[func() js.Promise[js.Array[ThroughputTrackerAnimationData]]]) {
	bindings.FuncStopThroughputTrackerDataCollection(
		js.Pointer(&fn),
	)
	return
}

// StopThroughputTrackerDataCollection calls the function "WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection" directly.
func StopThroughputTrackerDataCollection() (ret js.Promise[js.Array[ThroughputTrackerAnimationData]]) {
	bindings.CallStopThroughputTrackerDataCollection(
		js.Pointer(&ret),
	)

	return
}

// TryStopThroughputTrackerDataCollection calls the function "WEBEXT.autotestPrivate.stopThroughputTrackerDataCollection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopThroughputTrackerDataCollection() (ret js.Promise[js.Array[ThroughputTrackerAnimationData]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopThroughputTrackerDataCollection(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSwapWindowsInSplitView returns true if the function "WEBEXT.autotestPrivate.swapWindowsInSplitView" exists.
func HasFuncSwapWindowsInSplitView() bool {
	return js.True == bindings.HasFuncSwapWindowsInSplitView()
}

// FuncSwapWindowsInSplitView returns the function "WEBEXT.autotestPrivate.swapWindowsInSplitView".
func FuncSwapWindowsInSplitView() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncSwapWindowsInSplitView(
		js.Pointer(&fn),
	)
	return
}

// SwapWindowsInSplitView calls the function "WEBEXT.autotestPrivate.swapWindowsInSplitView" directly.
func SwapWindowsInSplitView() (ret js.Promise[js.Void]) {
	bindings.CallSwapWindowsInSplitView(
		js.Pointer(&ret),
	)

	return
}

// TrySwapWindowsInSplitView calls the function "WEBEXT.autotestPrivate.swapWindowsInSplitView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySwapWindowsInSplitView() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySwapWindowsInSplitView(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTakeScreenshot returns true if the function "WEBEXT.autotestPrivate.takeScreenshot" exists.
func HasFuncTakeScreenshot() bool {
	return js.True == bindings.HasFuncTakeScreenshot()
}

// FuncTakeScreenshot returns the function "WEBEXT.autotestPrivate.takeScreenshot".
func FuncTakeScreenshot() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncTakeScreenshot(
		js.Pointer(&fn),
	)
	return
}

// TakeScreenshot calls the function "WEBEXT.autotestPrivate.takeScreenshot" directly.
func TakeScreenshot() (ret js.Promise[js.String]) {
	bindings.CallTakeScreenshot(
		js.Pointer(&ret),
	)

	return
}

// TryTakeScreenshot calls the function "WEBEXT.autotestPrivate.takeScreenshot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTakeScreenshot() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTakeScreenshot(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTakeScreenshotForDisplay returns true if the function "WEBEXT.autotestPrivate.takeScreenshotForDisplay" exists.
func HasFuncTakeScreenshotForDisplay() bool {
	return js.True == bindings.HasFuncTakeScreenshotForDisplay()
}

// FuncTakeScreenshotForDisplay returns the function "WEBEXT.autotestPrivate.takeScreenshotForDisplay".
func FuncTakeScreenshotForDisplay() (fn js.Func[func(display_id js.String) js.Promise[js.String]]) {
	bindings.FuncTakeScreenshotForDisplay(
		js.Pointer(&fn),
	)
	return
}

// TakeScreenshotForDisplay calls the function "WEBEXT.autotestPrivate.takeScreenshotForDisplay" directly.
func TakeScreenshotForDisplay(display_id js.String) (ret js.Promise[js.String]) {
	bindings.CallTakeScreenshotForDisplay(
		js.Pointer(&ret),
		display_id.Ref(),
	)

	return
}

// TryTakeScreenshotForDisplay calls the function "WEBEXT.autotestPrivate.takeScreenshotForDisplay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTakeScreenshotForDisplay(display_id js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryTakeScreenshotForDisplay(
		js.Pointer(&ret), js.Pointer(&exception),
		display_id.Ref(),
	)

	return
}

// HasFuncUpdatePrinter returns true if the function "WEBEXT.autotestPrivate.updatePrinter" exists.
func HasFuncUpdatePrinter() bool {
	return js.True == bindings.HasFuncUpdatePrinter()
}

// FuncUpdatePrinter returns the function "WEBEXT.autotestPrivate.updatePrinter".
func FuncUpdatePrinter() (fn js.Func[func(printer Printer)]) {
	bindings.FuncUpdatePrinter(
		js.Pointer(&fn),
	)
	return
}

// UpdatePrinter calls the function "WEBEXT.autotestPrivate.updatePrinter" directly.
func UpdatePrinter(printer Printer) (ret js.Void) {
	bindings.CallUpdatePrinter(
		js.Pointer(&ret),
		js.Pointer(&printer),
	)

	return
}

// TryUpdatePrinter calls the function "WEBEXT.autotestPrivate.updatePrinter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdatePrinter(printer Printer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdatePrinter(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&printer),
	)

	return
}

// HasFuncWaitForAmbientPhotoAnimation returns true if the function "WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation" exists.
func HasFuncWaitForAmbientPhotoAnimation() bool {
	return js.True == bindings.HasFuncWaitForAmbientPhotoAnimation()
}

// FuncWaitForAmbientPhotoAnimation returns the function "WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation".
func FuncWaitForAmbientPhotoAnimation() (fn js.Func[func(numCompletions int32, timeout int32) js.Promise[js.Void]]) {
	bindings.FuncWaitForAmbientPhotoAnimation(
		js.Pointer(&fn),
	)
	return
}

// WaitForAmbientPhotoAnimation calls the function "WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation" directly.
func WaitForAmbientPhotoAnimation(numCompletions int32, timeout int32) (ret js.Promise[js.Void]) {
	bindings.CallWaitForAmbientPhotoAnimation(
		js.Pointer(&ret),
		int32(numCompletions),
		int32(timeout),
	)

	return
}

// TryWaitForAmbientPhotoAnimation calls the function "WEBEXT.autotestPrivate.waitForAmbientPhotoAnimation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForAmbientPhotoAnimation(numCompletions int32, timeout int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForAmbientPhotoAnimation(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(numCompletions),
		int32(timeout),
	)

	return
}

// HasFuncWaitForAmbientVideo returns true if the function "WEBEXT.autotestPrivate.waitForAmbientVideo" exists.
func HasFuncWaitForAmbientVideo() bool {
	return js.True == bindings.HasFuncWaitForAmbientVideo()
}

// FuncWaitForAmbientVideo returns the function "WEBEXT.autotestPrivate.waitForAmbientVideo".
func FuncWaitForAmbientVideo() (fn js.Func[func(timeout int32) js.Promise[js.Void]]) {
	bindings.FuncWaitForAmbientVideo(
		js.Pointer(&fn),
	)
	return
}

// WaitForAmbientVideo calls the function "WEBEXT.autotestPrivate.waitForAmbientVideo" directly.
func WaitForAmbientVideo(timeout int32) (ret js.Promise[js.Void]) {
	bindings.CallWaitForAmbientVideo(
		js.Pointer(&ret),
		int32(timeout),
	)

	return
}

// TryWaitForAmbientVideo calls the function "WEBEXT.autotestPrivate.waitForAmbientVideo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForAmbientVideo(timeout int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForAmbientVideo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(timeout),
	)

	return
}

// HasFuncWaitForAssistantQueryStatus returns true if the function "WEBEXT.autotestPrivate.waitForAssistantQueryStatus" exists.
func HasFuncWaitForAssistantQueryStatus() bool {
	return js.True == bindings.HasFuncWaitForAssistantQueryStatus()
}

// FuncWaitForAssistantQueryStatus returns the function "WEBEXT.autotestPrivate.waitForAssistantQueryStatus".
func FuncWaitForAssistantQueryStatus() (fn js.Func[func(timeout_s int32) js.Promise[AssistantQueryStatus]]) {
	bindings.FuncWaitForAssistantQueryStatus(
		js.Pointer(&fn),
	)
	return
}

// WaitForAssistantQueryStatus calls the function "WEBEXT.autotestPrivate.waitForAssistantQueryStatus" directly.
func WaitForAssistantQueryStatus(timeout_s int32) (ret js.Promise[AssistantQueryStatus]) {
	bindings.CallWaitForAssistantQueryStatus(
		js.Pointer(&ret),
		int32(timeout_s),
	)

	return
}

// TryWaitForAssistantQueryStatus calls the function "WEBEXT.autotestPrivate.waitForAssistantQueryStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForAssistantQueryStatus(timeout_s int32) (ret js.Promise[AssistantQueryStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForAssistantQueryStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(timeout_s),
	)

	return
}

// HasFuncWaitForDisplayRotation returns true if the function "WEBEXT.autotestPrivate.waitForDisplayRotation" exists.
func HasFuncWaitForDisplayRotation() bool {
	return js.True == bindings.HasFuncWaitForDisplayRotation()
}

// FuncWaitForDisplayRotation returns the function "WEBEXT.autotestPrivate.waitForDisplayRotation".
func FuncWaitForDisplayRotation() (fn js.Func[func(displayId js.String, rotation RotationType) js.Promise[js.Boolean]]) {
	bindings.FuncWaitForDisplayRotation(
		js.Pointer(&fn),
	)
	return
}

// WaitForDisplayRotation calls the function "WEBEXT.autotestPrivate.waitForDisplayRotation" directly.
func WaitForDisplayRotation(displayId js.String, rotation RotationType) (ret js.Promise[js.Boolean]) {
	bindings.CallWaitForDisplayRotation(
		js.Pointer(&ret),
		displayId.Ref(),
		uint32(rotation),
	)

	return
}

// TryWaitForDisplayRotation calls the function "WEBEXT.autotestPrivate.waitForDisplayRotation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForDisplayRotation(displayId js.String, rotation RotationType) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForDisplayRotation(
		js.Pointer(&ret), js.Pointer(&exception),
		displayId.Ref(),
		uint32(rotation),
	)

	return
}

// HasFuncWaitForLauncherState returns true if the function "WEBEXT.autotestPrivate.waitForLauncherState" exists.
func HasFuncWaitForLauncherState() bool {
	return js.True == bindings.HasFuncWaitForLauncherState()
}

// FuncWaitForLauncherState returns the function "WEBEXT.autotestPrivate.waitForLauncherState".
func FuncWaitForLauncherState() (fn js.Func[func(launcherState LauncherStateType) js.Promise[js.Void]]) {
	bindings.FuncWaitForLauncherState(
		js.Pointer(&fn),
	)
	return
}

// WaitForLauncherState calls the function "WEBEXT.autotestPrivate.waitForLauncherState" directly.
func WaitForLauncherState(launcherState LauncherStateType) (ret js.Promise[js.Void]) {
	bindings.CallWaitForLauncherState(
		js.Pointer(&ret),
		uint32(launcherState),
	)

	return
}

// TryWaitForLauncherState calls the function "WEBEXT.autotestPrivate.waitForLauncherState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForLauncherState(launcherState LauncherStateType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForLauncherState(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(launcherState),
	)

	return
}

// HasFuncWaitForOverviewState returns true if the function "WEBEXT.autotestPrivate.waitForOverviewState" exists.
func HasFuncWaitForOverviewState() bool {
	return js.True == bindings.HasFuncWaitForOverviewState()
}

// FuncWaitForOverviewState returns the function "WEBEXT.autotestPrivate.waitForOverviewState".
func FuncWaitForOverviewState() (fn js.Func[func(overviewState OverviewStateType) js.Promise[js.Void]]) {
	bindings.FuncWaitForOverviewState(
		js.Pointer(&fn),
	)
	return
}

// WaitForOverviewState calls the function "WEBEXT.autotestPrivate.waitForOverviewState" directly.
func WaitForOverviewState(overviewState OverviewStateType) (ret js.Promise[js.Void]) {
	bindings.CallWaitForOverviewState(
		js.Pointer(&ret),
		uint32(overviewState),
	)

	return
}

// TryWaitForOverviewState calls the function "WEBEXT.autotestPrivate.waitForOverviewState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForOverviewState(overviewState OverviewStateType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForOverviewState(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(overviewState),
	)

	return
}

// HasFuncWaitForSystemWebAppsInstall returns true if the function "WEBEXT.autotestPrivate.waitForSystemWebAppsInstall" exists.
func HasFuncWaitForSystemWebAppsInstall() bool {
	return js.True == bindings.HasFuncWaitForSystemWebAppsInstall()
}

// FuncWaitForSystemWebAppsInstall returns the function "WEBEXT.autotestPrivate.waitForSystemWebAppsInstall".
func FuncWaitForSystemWebAppsInstall() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWaitForSystemWebAppsInstall(
		js.Pointer(&fn),
	)
	return
}

// WaitForSystemWebAppsInstall calls the function "WEBEXT.autotestPrivate.waitForSystemWebAppsInstall" directly.
func WaitForSystemWebAppsInstall() (ret js.Promise[js.Void]) {
	bindings.CallWaitForSystemWebAppsInstall(
		js.Pointer(&ret),
	)

	return
}

// TryWaitForSystemWebAppsInstall calls the function "WEBEXT.autotestPrivate.waitForSystemWebAppsInstall"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWaitForSystemWebAppsInstall() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWaitForSystemWebAppsInstall(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

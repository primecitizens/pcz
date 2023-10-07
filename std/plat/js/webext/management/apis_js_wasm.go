// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package management

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/management/bindings"
)

type ExtensionDisabledReason uint32

const (
	_ ExtensionDisabledReason = iota

	ExtensionDisabledReason_UNKNOWN
	ExtensionDisabledReason_PERMISSIONS_INCREASE
)

func (ExtensionDisabledReason) FromRef(str js.Ref) ExtensionDisabledReason {
	return ExtensionDisabledReason(bindings.ConstOfExtensionDisabledReason(str))
}

func (x ExtensionDisabledReason) String() (string, bool) {
	switch x {
	case ExtensionDisabledReason_UNKNOWN:
		return "unknown", true
	case ExtensionDisabledReason_PERMISSIONS_INCREASE:
		return "permissions_increase", true
	default:
		return "", false
	}
}

type LaunchType uint32

const (
	_ LaunchType = iota

	LaunchType_OPEN_AS_REGULAR_TAB
	LaunchType_OPEN_AS_PINNED_TAB
	LaunchType_OPEN_AS_WINDOW
	LaunchType_OPEN_FULL_SCREEN
)

func (LaunchType) FromRef(str js.Ref) LaunchType {
	return LaunchType(bindings.ConstOfLaunchType(str))
}

func (x LaunchType) String() (string, bool) {
	switch x {
	case LaunchType_OPEN_AS_REGULAR_TAB:
		return "OPEN_AS_REGULAR_TAB", true
	case LaunchType_OPEN_AS_PINNED_TAB:
		return "OPEN_AS_PINNED_TAB", true
	case LaunchType_OPEN_AS_WINDOW:
		return "OPEN_AS_WINDOW", true
	case LaunchType_OPEN_FULL_SCREEN:
		return "OPEN_FULL_SCREEN", true
	default:
		return "", false
	}
}

type IconInfo struct {
	// Size is "IconInfo.size"
	//
	// Required
	Size int64
	// Url is "IconInfo.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IconInfo with all fields set.
func (p IconInfo) FromRef(ref js.Ref) IconInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IconInfo in the application heap.
func (p IconInfo) New() js.Ref {
	return bindings.IconInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IconInfo) UpdateFrom(ref js.Ref) {
	bindings.IconInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IconInfo) Update(ref js.Ref) {
	bindings.IconInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IconInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type ExtensionInstallType uint32

const (
	_ ExtensionInstallType = iota

	ExtensionInstallType_ADMIN
	ExtensionInstallType_DEVELOPMENT
	ExtensionInstallType_NORMAL
	ExtensionInstallType_SIDELOAD
	ExtensionInstallType_OTHER
)

func (ExtensionInstallType) FromRef(str js.Ref) ExtensionInstallType {
	return ExtensionInstallType(bindings.ConstOfExtensionInstallType(str))
}

func (x ExtensionInstallType) String() (string, bool) {
	switch x {
	case ExtensionInstallType_ADMIN:
		return "admin", true
	case ExtensionInstallType_DEVELOPMENT:
		return "development", true
	case ExtensionInstallType_NORMAL:
		return "normal", true
	case ExtensionInstallType_SIDELOAD:
		return "sideload", true
	case ExtensionInstallType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type ExtensionType uint32

const (
	_ ExtensionType = iota

	ExtensionType_EXTENSION
	ExtensionType_HOSTED_APP
	ExtensionType_PACKAGED_APP
	ExtensionType_LEGACY_PACKAGED_APP
	ExtensionType_THEME
	ExtensionType_LOGIN_SCREEN_EXTENSION
)

func (ExtensionType) FromRef(str js.Ref) ExtensionType {
	return ExtensionType(bindings.ConstOfExtensionType(str))
}

func (x ExtensionType) String() (string, bool) {
	switch x {
	case ExtensionType_EXTENSION:
		return "extension", true
	case ExtensionType_HOSTED_APP:
		return "hosted_app", true
	case ExtensionType_PACKAGED_APP:
		return "packaged_app", true
	case ExtensionType_LEGACY_PACKAGED_APP:
		return "legacy_packaged_app", true
	case ExtensionType_THEME:
		return "theme", true
	case ExtensionType_LOGIN_SCREEN_EXTENSION:
		return "login_screen_extension", true
	default:
		return "", false
	}
}

type ExtensionInfo struct {
	// AppLaunchUrl is "ExtensionInfo.appLaunchUrl"
	//
	// Optional
	AppLaunchUrl js.String
	// AvailableLaunchTypes is "ExtensionInfo.availableLaunchTypes"
	//
	// Optional
	AvailableLaunchTypes js.Array[LaunchType]
	// Description is "ExtensionInfo.description"
	//
	// Required
	Description js.String
	// DisabledReason is "ExtensionInfo.disabledReason"
	//
	// Optional
	DisabledReason ExtensionDisabledReason
	// Enabled is "ExtensionInfo.enabled"
	//
	// Required
	Enabled bool
	// HomepageUrl is "ExtensionInfo.homepageUrl"
	//
	// Optional
	HomepageUrl js.String
	// HostPermissions is "ExtensionInfo.hostPermissions"
	//
	// Required
	HostPermissions js.Array[js.String]
	// Icons is "ExtensionInfo.icons"
	//
	// Optional
	Icons js.Array[IconInfo]
	// Id is "ExtensionInfo.id"
	//
	// Required
	Id js.String
	// InstallType is "ExtensionInfo.installType"
	//
	// Required
	InstallType ExtensionInstallType
	// IsApp is "ExtensionInfo.isApp"
	//
	// Required
	IsApp bool
	// LaunchType is "ExtensionInfo.launchType"
	//
	// Optional
	LaunchType LaunchType
	// MayDisable is "ExtensionInfo.mayDisable"
	//
	// Required
	MayDisable bool
	// MayEnable is "ExtensionInfo.mayEnable"
	//
	// Optional
	//
	// NOTE: FFI_USE_MayEnable MUST be set to true to make this field effective.
	MayEnable bool
	// Name is "ExtensionInfo.name"
	//
	// Required
	Name js.String
	// OfflineEnabled is "ExtensionInfo.offlineEnabled"
	//
	// Required
	OfflineEnabled bool
	// OptionsUrl is "ExtensionInfo.optionsUrl"
	//
	// Required
	OptionsUrl js.String
	// Permissions is "ExtensionInfo.permissions"
	//
	// Required
	Permissions js.Array[js.String]
	// ShortName is "ExtensionInfo.shortName"
	//
	// Required
	ShortName js.String
	// Type is "ExtensionInfo.type"
	//
	// Required
	Type ExtensionType
	// UpdateUrl is "ExtensionInfo.updateUrl"
	//
	// Optional
	UpdateUrl js.String
	// Version is "ExtensionInfo.version"
	//
	// Required
	Version js.String
	// VersionName is "ExtensionInfo.versionName"
	//
	// Optional
	VersionName js.String

	FFI_USE_MayEnable bool // for MayEnable.

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
		p.AppLaunchUrl.Ref(),
		p.AvailableLaunchTypes.Ref(),
		p.Description.Ref(),
		p.HomepageUrl.Ref(),
		p.HostPermissions.Ref(),
		p.Icons.Ref(),
		p.Id.Ref(),
		p.Name.Ref(),
		p.OptionsUrl.Ref(),
		p.Permissions.Ref(),
		p.ShortName.Ref(),
		p.UpdateUrl.Ref(),
		p.Version.Ref(),
		p.VersionName.Ref(),
	)
	p.AppLaunchUrl = p.AppLaunchUrl.FromRef(js.Undefined)
	p.AvailableLaunchTypes = p.AvailableLaunchTypes.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.HomepageUrl = p.HomepageUrl.FromRef(js.Undefined)
	p.HostPermissions = p.HostPermissions.FromRef(js.Undefined)
	p.Icons = p.Icons.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.OptionsUrl = p.OptionsUrl.FromRef(js.Undefined)
	p.Permissions = p.Permissions.FromRef(js.Undefined)
	p.ShortName = p.ShortName.FromRef(js.Undefined)
	p.UpdateUrl = p.UpdateUrl.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
	p.VersionName = p.VersionName.FromRef(js.Undefined)
}

type UninstallOptions struct {
	// ShowConfirmDialog is "UninstallOptions.showConfirmDialog"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowConfirmDialog MUST be set to true to make this field effective.
	ShowConfirmDialog bool

	FFI_USE_ShowConfirmDialog bool // for ShowConfirmDialog.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UninstallOptions with all fields set.
func (p UninstallOptions) FromRef(ref js.Ref) UninstallOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UninstallOptions in the application heap.
func (p UninstallOptions) New() js.Ref {
	return bindings.UninstallOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UninstallOptions) UpdateFrom(ref js.Ref) {
	bindings.UninstallOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UninstallOptions) Update(ref js.Ref) {
	bindings.UninstallOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UninstallOptions) FreeMembers(recursive bool) {
}

// HasFuncCreateAppShortcut returns true if the function "WEBEXT.management.createAppShortcut" exists.
func HasFuncCreateAppShortcut() bool {
	return js.True == bindings.HasFuncCreateAppShortcut()
}

// FuncCreateAppShortcut returns the function "WEBEXT.management.createAppShortcut".
func FuncCreateAppShortcut() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncCreateAppShortcut(
		js.Pointer(&fn),
	)
	return
}

// CreateAppShortcut calls the function "WEBEXT.management.createAppShortcut" directly.
func CreateAppShortcut(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallCreateAppShortcut(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryCreateAppShortcut calls the function "WEBEXT.management.createAppShortcut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateAppShortcut(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateAppShortcut(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGenerateAppForLink returns true if the function "WEBEXT.management.generateAppForLink" exists.
func HasFuncGenerateAppForLink() bool {
	return js.True == bindings.HasFuncGenerateAppForLink()
}

// FuncGenerateAppForLink returns the function "WEBEXT.management.generateAppForLink".
func FuncGenerateAppForLink() (fn js.Func[func(url js.String, title js.String) js.Promise[ExtensionInfo]]) {
	bindings.FuncGenerateAppForLink(
		js.Pointer(&fn),
	)
	return
}

// GenerateAppForLink calls the function "WEBEXT.management.generateAppForLink" directly.
func GenerateAppForLink(url js.String, title js.String) (ret js.Promise[ExtensionInfo]) {
	bindings.CallGenerateAppForLink(
		js.Pointer(&ret),
		url.Ref(),
		title.Ref(),
	)

	return
}

// TryGenerateAppForLink calls the function "WEBEXT.management.generateAppForLink"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGenerateAppForLink(url js.String, title js.String) (ret js.Promise[ExtensionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGenerateAppForLink(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		title.Ref(),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.management.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.management.get".
func FuncGet() (fn js.Func[func(id js.String) js.Promise[ExtensionInfo]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.management.get" directly.
func Get(id js.String) (ret js.Promise[ExtensionInfo]) {
	bindings.CallGet(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.management.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(id js.String) (ret js.Promise[ExtensionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.management.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.management.getAll".
func FuncGetAll() (fn js.Func[func() js.Promise[js.Array[ExtensionInfo]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.management.getAll" directly.
func GetAll() (ret js.Promise[js.Array[ExtensionInfo]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the function "WEBEXT.management.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll() (ret js.Promise[js.Array[ExtensionInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPermissionWarningsById returns true if the function "WEBEXT.management.getPermissionWarningsById" exists.
func HasFuncGetPermissionWarningsById() bool {
	return js.True == bindings.HasFuncGetPermissionWarningsById()
}

// FuncGetPermissionWarningsById returns the function "WEBEXT.management.getPermissionWarningsById".
func FuncGetPermissionWarningsById() (fn js.Func[func(id js.String) js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetPermissionWarningsById(
		js.Pointer(&fn),
	)
	return
}

// GetPermissionWarningsById calls the function "WEBEXT.management.getPermissionWarningsById" directly.
func GetPermissionWarningsById(id js.String) (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetPermissionWarningsById(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetPermissionWarningsById calls the function "WEBEXT.management.getPermissionWarningsById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPermissionWarningsById(id js.String) (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPermissionWarningsById(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetPermissionWarningsByManifest returns true if the function "WEBEXT.management.getPermissionWarningsByManifest" exists.
func HasFuncGetPermissionWarningsByManifest() bool {
	return js.True == bindings.HasFuncGetPermissionWarningsByManifest()
}

// FuncGetPermissionWarningsByManifest returns the function "WEBEXT.management.getPermissionWarningsByManifest".
func FuncGetPermissionWarningsByManifest() (fn js.Func[func(manifestStr js.String) js.Promise[js.Array[js.String]]]) {
	bindings.FuncGetPermissionWarningsByManifest(
		js.Pointer(&fn),
	)
	return
}

// GetPermissionWarningsByManifest calls the function "WEBEXT.management.getPermissionWarningsByManifest" directly.
func GetPermissionWarningsByManifest(manifestStr js.String) (ret js.Promise[js.Array[js.String]]) {
	bindings.CallGetPermissionWarningsByManifest(
		js.Pointer(&ret),
		manifestStr.Ref(),
	)

	return
}

// TryGetPermissionWarningsByManifest calls the function "WEBEXT.management.getPermissionWarningsByManifest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPermissionWarningsByManifest(manifestStr js.String) (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPermissionWarningsByManifest(
		js.Pointer(&ret), js.Pointer(&exception),
		manifestStr.Ref(),
	)

	return
}

// HasFuncGetSelf returns true if the function "WEBEXT.management.getSelf" exists.
func HasFuncGetSelf() bool {
	return js.True == bindings.HasFuncGetSelf()
}

// FuncGetSelf returns the function "WEBEXT.management.getSelf".
func FuncGetSelf() (fn js.Func[func() js.Promise[ExtensionInfo]]) {
	bindings.FuncGetSelf(
		js.Pointer(&fn),
	)
	return
}

// GetSelf calls the function "WEBEXT.management.getSelf" directly.
func GetSelf() (ret js.Promise[ExtensionInfo]) {
	bindings.CallGetSelf(
		js.Pointer(&ret),
	)

	return
}

// TryGetSelf calls the function "WEBEXT.management.getSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSelf() (ret js.Promise[ExtensionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSelf(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInstallReplacementWebApp returns true if the function "WEBEXT.management.installReplacementWebApp" exists.
func HasFuncInstallReplacementWebApp() bool {
	return js.True == bindings.HasFuncInstallReplacementWebApp()
}

// FuncInstallReplacementWebApp returns the function "WEBEXT.management.installReplacementWebApp".
func FuncInstallReplacementWebApp() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncInstallReplacementWebApp(
		js.Pointer(&fn),
	)
	return
}

// InstallReplacementWebApp calls the function "WEBEXT.management.installReplacementWebApp" directly.
func InstallReplacementWebApp() (ret js.Promise[js.Void]) {
	bindings.CallInstallReplacementWebApp(
		js.Pointer(&ret),
	)

	return
}

// TryInstallReplacementWebApp calls the function "WEBEXT.management.installReplacementWebApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstallReplacementWebApp() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstallReplacementWebApp(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLaunchApp returns true if the function "WEBEXT.management.launchApp" exists.
func HasFuncLaunchApp() bool {
	return js.True == bindings.HasFuncLaunchApp()
}

// FuncLaunchApp returns the function "WEBEXT.management.launchApp".
func FuncLaunchApp() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncLaunchApp(
		js.Pointer(&fn),
	)
	return
}

// LaunchApp calls the function "WEBEXT.management.launchApp" directly.
func LaunchApp(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallLaunchApp(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryLaunchApp calls the function "WEBEXT.management.launchApp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLaunchApp(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchApp(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type OnDisabledEventCallbackFunc func(this js.Ref, info *ExtensionInfo) js.Ref

func (fn OnDisabledEventCallbackFunc) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDisabledEventCallbackFunc) DispatchCallback(
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

type OnDisabledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ExtensionInfo) js.Ref
	Arg T
}

func (cb *OnDisabledEventCallback[T]) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDisabledEventCallback[T]) DispatchCallback(
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

// HasFuncOnDisabled returns true if the function "WEBEXT.management.onDisabled.addListener" exists.
func HasFuncOnDisabled() bool {
	return js.True == bindings.HasFuncOnDisabled()
}

// FuncOnDisabled returns the function "WEBEXT.management.onDisabled.addListener".
func FuncOnDisabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOnDisabled(
		js.Pointer(&fn),
	)
	return
}

// OnDisabled calls the function "WEBEXT.management.onDisabled.addListener" directly.
func OnDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOnDisabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDisabled calls the function "WEBEXT.management.onDisabled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDisabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDisabled returns true if the function "WEBEXT.management.onDisabled.removeListener" exists.
func HasFuncOffDisabled() bool {
	return js.True == bindings.HasFuncOffDisabled()
}

// FuncOffDisabled returns the function "WEBEXT.management.onDisabled.removeListener".
func FuncOffDisabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOffDisabled(
		js.Pointer(&fn),
	)
	return
}

// OffDisabled calls the function "WEBEXT.management.onDisabled.removeListener" directly.
func OffDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOffDisabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDisabled calls the function "WEBEXT.management.onDisabled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDisabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDisabled returns true if the function "WEBEXT.management.onDisabled.hasListener" exists.
func HasFuncHasOnDisabled() bool {
	return js.True == bindings.HasFuncHasOnDisabled()
}

// FuncHasOnDisabled returns the function "WEBEXT.management.onDisabled.hasListener".
func FuncHasOnDisabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)]) bool]) {
	bindings.FuncHasOnDisabled(
		js.Pointer(&fn),
	)
	return
}

// HasOnDisabled calls the function "WEBEXT.management.onDisabled.hasListener" directly.
func HasOnDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret bool) {
	bindings.CallHasOnDisabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDisabled calls the function "WEBEXT.management.onDisabled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDisabled(callback js.Func[func(info *ExtensionInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDisabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnEnabledEventCallbackFunc func(this js.Ref, info *ExtensionInfo) js.Ref

func (fn OnEnabledEventCallbackFunc) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEnabledEventCallbackFunc) DispatchCallback(
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

type OnEnabledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ExtensionInfo) js.Ref
	Arg T
}

func (cb *OnEnabledEventCallback[T]) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEnabledEventCallback[T]) DispatchCallback(
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

// HasFuncOnEnabled returns true if the function "WEBEXT.management.onEnabled.addListener" exists.
func HasFuncOnEnabled() bool {
	return js.True == bindings.HasFuncOnEnabled()
}

// FuncOnEnabled returns the function "WEBEXT.management.onEnabled.addListener".
func FuncOnEnabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOnEnabled(
		js.Pointer(&fn),
	)
	return
}

// OnEnabled calls the function "WEBEXT.management.onEnabled.addListener" directly.
func OnEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOnEnabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEnabled calls the function "WEBEXT.management.onEnabled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEnabled returns true if the function "WEBEXT.management.onEnabled.removeListener" exists.
func HasFuncOffEnabled() bool {
	return js.True == bindings.HasFuncOffEnabled()
}

// FuncOffEnabled returns the function "WEBEXT.management.onEnabled.removeListener".
func FuncOffEnabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOffEnabled(
		js.Pointer(&fn),
	)
	return
}

// OffEnabled calls the function "WEBEXT.management.onEnabled.removeListener" directly.
func OffEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOffEnabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEnabled calls the function "WEBEXT.management.onEnabled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEnabled returns true if the function "WEBEXT.management.onEnabled.hasListener" exists.
func HasFuncHasOnEnabled() bool {
	return js.True == bindings.HasFuncHasOnEnabled()
}

// FuncHasOnEnabled returns the function "WEBEXT.management.onEnabled.hasListener".
func FuncHasOnEnabled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)]) bool]) {
	bindings.FuncHasOnEnabled(
		js.Pointer(&fn),
	)
	return
}

// HasOnEnabled calls the function "WEBEXT.management.onEnabled.hasListener" directly.
func HasOnEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret bool) {
	bindings.CallHasOnEnabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEnabled calls the function "WEBEXT.management.onEnabled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEnabled(callback js.Func[func(info *ExtensionInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInstalledEventCallbackFunc func(this js.Ref, info *ExtensionInfo) js.Ref

func (fn OnInstalledEventCallbackFunc) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInstalledEventCallbackFunc) DispatchCallback(
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

type OnInstalledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ExtensionInfo) js.Ref
	Arg T
}

func (cb *OnInstalledEventCallback[T]) Register() js.Func[func(info *ExtensionInfo)] {
	return js.RegisterCallback[func(info *ExtensionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInstalledEventCallback[T]) DispatchCallback(
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

// HasFuncOnInstalled returns true if the function "WEBEXT.management.onInstalled.addListener" exists.
func HasFuncOnInstalled() bool {
	return js.True == bindings.HasFuncOnInstalled()
}

// FuncOnInstalled returns the function "WEBEXT.management.onInstalled.addListener".
func FuncOnInstalled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOnInstalled(
		js.Pointer(&fn),
	)
	return
}

// OnInstalled calls the function "WEBEXT.management.onInstalled.addListener" directly.
func OnInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOnInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInstalled calls the function "WEBEXT.management.onInstalled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInstalled returns true if the function "WEBEXT.management.onInstalled.removeListener" exists.
func HasFuncOffInstalled() bool {
	return js.True == bindings.HasFuncOffInstalled()
}

// FuncOffInstalled returns the function "WEBEXT.management.onInstalled.removeListener".
func FuncOffInstalled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)])]) {
	bindings.FuncOffInstalled(
		js.Pointer(&fn),
	)
	return
}

// OffInstalled calls the function "WEBEXT.management.onInstalled.removeListener" directly.
func OffInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void) {
	bindings.CallOffInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInstalled calls the function "WEBEXT.management.onInstalled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInstalled returns true if the function "WEBEXT.management.onInstalled.hasListener" exists.
func HasFuncHasOnInstalled() bool {
	return js.True == bindings.HasFuncHasOnInstalled()
}

// FuncHasOnInstalled returns the function "WEBEXT.management.onInstalled.hasListener".
func FuncHasOnInstalled() (fn js.Func[func(callback js.Func[func(info *ExtensionInfo)]) bool]) {
	bindings.FuncHasOnInstalled(
		js.Pointer(&fn),
	)
	return
}

// HasOnInstalled calls the function "WEBEXT.management.onInstalled.hasListener" directly.
func HasOnInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret bool) {
	bindings.CallHasOnInstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInstalled calls the function "WEBEXT.management.onInstalled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInstalled(callback js.Func[func(info *ExtensionInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUninstalledEventCallbackFunc func(this js.Ref, id js.String) js.Ref

func (fn OnUninstalledEventCallbackFunc) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUninstalledEventCallbackFunc) DispatchCallback(
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

type OnUninstalledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String) js.Ref
	Arg T
}

func (cb *OnUninstalledEventCallback[T]) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUninstalledEventCallback[T]) DispatchCallback(
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

// HasFuncOnUninstalled returns true if the function "WEBEXT.management.onUninstalled.addListener" exists.
func HasFuncOnUninstalled() bool {
	return js.True == bindings.HasFuncOnUninstalled()
}

// FuncOnUninstalled returns the function "WEBEXT.management.onUninstalled.addListener".
func FuncOnUninstalled() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOnUninstalled(
		js.Pointer(&fn),
	)
	return
}

// OnUninstalled calls the function "WEBEXT.management.onUninstalled.addListener" directly.
func OnUninstalled(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOnUninstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUninstalled calls the function "WEBEXT.management.onUninstalled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUninstalled(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUninstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUninstalled returns true if the function "WEBEXT.management.onUninstalled.removeListener" exists.
func HasFuncOffUninstalled() bool {
	return js.True == bindings.HasFuncOffUninstalled()
}

// FuncOffUninstalled returns the function "WEBEXT.management.onUninstalled.removeListener".
func FuncOffUninstalled() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOffUninstalled(
		js.Pointer(&fn),
	)
	return
}

// OffUninstalled calls the function "WEBEXT.management.onUninstalled.removeListener" directly.
func OffUninstalled(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOffUninstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUninstalled calls the function "WEBEXT.management.onUninstalled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUninstalled(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUninstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUninstalled returns true if the function "WEBEXT.management.onUninstalled.hasListener" exists.
func HasFuncHasOnUninstalled() bool {
	return js.True == bindings.HasFuncHasOnUninstalled()
}

// FuncHasOnUninstalled returns the function "WEBEXT.management.onUninstalled.hasListener".
func FuncHasOnUninstalled() (fn js.Func[func(callback js.Func[func(id js.String)]) bool]) {
	bindings.FuncHasOnUninstalled(
		js.Pointer(&fn),
	)
	return
}

// HasOnUninstalled calls the function "WEBEXT.management.onUninstalled.hasListener" directly.
func HasOnUninstalled(callback js.Func[func(id js.String)]) (ret bool) {
	bindings.CallHasOnUninstalled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUninstalled calls the function "WEBEXT.management.onUninstalled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUninstalled(callback js.Func[func(id js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUninstalled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetEnabled returns true if the function "WEBEXT.management.setEnabled" exists.
func HasFuncSetEnabled() bool {
	return js.True == bindings.HasFuncSetEnabled()
}

// FuncSetEnabled returns the function "WEBEXT.management.setEnabled".
func FuncSetEnabled() (fn js.Func[func(id js.String, enabled bool) js.Promise[js.Void]]) {
	bindings.FuncSetEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetEnabled calls the function "WEBEXT.management.setEnabled" directly.
func SetEnabled(id js.String, enabled bool) (ret js.Promise[js.Void]) {
	bindings.CallSetEnabled(
		js.Pointer(&ret),
		id.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetEnabled calls the function "WEBEXT.management.setEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetEnabled(id js.String, enabled bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetLaunchType returns true if the function "WEBEXT.management.setLaunchType" exists.
func HasFuncSetLaunchType() bool {
	return js.True == bindings.HasFuncSetLaunchType()
}

// FuncSetLaunchType returns the function "WEBEXT.management.setLaunchType".
func FuncSetLaunchType() (fn js.Func[func(id js.String, launchType LaunchType) js.Promise[js.Void]]) {
	bindings.FuncSetLaunchType(
		js.Pointer(&fn),
	)
	return
}

// SetLaunchType calls the function "WEBEXT.management.setLaunchType" directly.
func SetLaunchType(id js.String, launchType LaunchType) (ret js.Promise[js.Void]) {
	bindings.CallSetLaunchType(
		js.Pointer(&ret),
		id.Ref(),
		uint32(launchType),
	)

	return
}

// TrySetLaunchType calls the function "WEBEXT.management.setLaunchType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetLaunchType(id js.String, launchType LaunchType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetLaunchType(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		uint32(launchType),
	)

	return
}

// HasFuncUninstall returns true if the function "WEBEXT.management.uninstall" exists.
func HasFuncUninstall() bool {
	return js.True == bindings.HasFuncUninstall()
}

// FuncUninstall returns the function "WEBEXT.management.uninstall".
func FuncUninstall() (fn js.Func[func(id js.String, options UninstallOptions) js.Promise[js.Void]]) {
	bindings.FuncUninstall(
		js.Pointer(&fn),
	)
	return
}

// Uninstall calls the function "WEBEXT.management.uninstall" directly.
func Uninstall(id js.String, options UninstallOptions) (ret js.Promise[js.Void]) {
	bindings.CallUninstall(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryUninstall calls the function "WEBEXT.management.uninstall"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUninstall(id js.String, options UninstallOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUninstall(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncUninstallSelf returns true if the function "WEBEXT.management.uninstallSelf" exists.
func HasFuncUninstallSelf() bool {
	return js.True == bindings.HasFuncUninstallSelf()
}

// FuncUninstallSelf returns the function "WEBEXT.management.uninstallSelf".
func FuncUninstallSelf() (fn js.Func[func(options UninstallOptions) js.Promise[js.Void]]) {
	bindings.FuncUninstallSelf(
		js.Pointer(&fn),
	)
	return
}

// UninstallSelf calls the function "WEBEXT.management.uninstallSelf" directly.
func UninstallSelf(options UninstallOptions) (ret js.Promise[js.Void]) {
	bindings.CallUninstallSelf(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUninstallSelf calls the function "WEBEXT.management.uninstallSelf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUninstallSelf(options UninstallOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUninstallSelf(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

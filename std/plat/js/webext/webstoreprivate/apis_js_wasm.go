// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webstoreprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webstoreprivate/bindings"
)

type ExtensionInstallStatus uint32

const (
	_ ExtensionInstallStatus = iota

	ExtensionInstallStatus_CAN_REQUEST
	ExtensionInstallStatus_REQUEST_PENDING
	ExtensionInstallStatus_BLOCKED_BY_POLICY
	ExtensionInstallStatus_INSTALLABLE
	ExtensionInstallStatus_ENABLED
	ExtensionInstallStatus_DISABLED
	ExtensionInstallStatus_TERMINATED
	ExtensionInstallStatus_BLACKLISTED
	ExtensionInstallStatus_CUSTODIAN_APPROVAL_REQUIRED
	ExtensionInstallStatus_FORCE_INSTALLED
)

func (ExtensionInstallStatus) FromRef(str js.Ref) ExtensionInstallStatus {
	return ExtensionInstallStatus(bindings.ConstOfExtensionInstallStatus(str))
}

func (x ExtensionInstallStatus) String() (string, bool) {
	switch x {
	case ExtensionInstallStatus_CAN_REQUEST:
		return "can_request", true
	case ExtensionInstallStatus_REQUEST_PENDING:
		return "request_pending", true
	case ExtensionInstallStatus_BLOCKED_BY_POLICY:
		return "blocked_by_policy", true
	case ExtensionInstallStatus_INSTALLABLE:
		return "installable", true
	case ExtensionInstallStatus_ENABLED:
		return "enabled", true
	case ExtensionInstallStatus_DISABLED:
		return "disabled", true
	case ExtensionInstallStatus_TERMINATED:
		return "terminated", true
	case ExtensionInstallStatus_BLACKLISTED:
		return "blacklisted", true
	case ExtensionInstallStatus_CUSTODIAN_APPROVAL_REQUIRED:
		return "custodian_approval_required", true
	case ExtensionInstallStatus_FORCE_INSTALLED:
		return "force_installed", true
	default:
		return "", false
	}
}

type GetBrowserLoginReturnType struct {
	// Login is "GetBrowserLoginReturnType.login"
	//
	// Required
	Login js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetBrowserLoginReturnType with all fields set.
func (p GetBrowserLoginReturnType) FromRef(ref js.Ref) GetBrowserLoginReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetBrowserLoginReturnType in the application heap.
func (p GetBrowserLoginReturnType) New() js.Ref {
	return bindings.GetBrowserLoginReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetBrowserLoginReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetBrowserLoginReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetBrowserLoginReturnType) Update(ref js.Ref) {
	bindings.GetBrowserLoginReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetBrowserLoginReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Login.Ref(),
	)
	p.Login = p.Login.FromRef(js.Undefined)
}

type Result uint32

const (
	_ Result = iota

	Result_
	Result_SUCCESS
	Result_USER_GESTURE_REQUIRED
	Result_UNKNOWN_ERROR
	Result_FEATURE_DISABLED
	Result_UNSUPPORTED_EXTENSION_TYPE
	Result_MISSING_DEPENDENCIES
	Result_INSTALL_ERROR
	Result_USER_CANCELLED
	Result_INVALID_ID
	Result_BLACKLISTED
	Result_BLOCKED_BY_POLICY
	Result_INSTALL_IN_PROGRESS
	Result_LAUNCH_IN_PROGRESS
	Result_MANIFEST_ERROR
	Result_ICON_ERROR
	Result_INVALID_ICON_URL
	Result_ALREADY_INSTALLED
	Result_BLOCKED_FOR_CHILD_ACCOUNT
)

func (Result) FromRef(str js.Ref) Result {
	return Result(bindings.ConstOfResult(str))
}

func (x Result) String() (string, bool) {
	switch x {
	case Result_:
		return "", true
	case Result_SUCCESS:
		return "success", true
	case Result_USER_GESTURE_REQUIRED:
		return "user_gesture_required", true
	case Result_UNKNOWN_ERROR:
		return "unknown_error", true
	case Result_FEATURE_DISABLED:
		return "feature_disabled", true
	case Result_UNSUPPORTED_EXTENSION_TYPE:
		return "unsupported_extension_type", true
	case Result_MISSING_DEPENDENCIES:
		return "missing_dependencies", true
	case Result_INSTALL_ERROR:
		return "install_error", true
	case Result_USER_CANCELLED:
		return "user_cancelled", true
	case Result_INVALID_ID:
		return "invalid_id", true
	case Result_BLACKLISTED:
		return "blacklisted", true
	case Result_BLOCKED_BY_POLICY:
		return "blocked_by_policy", true
	case Result_INSTALL_IN_PROGRESS:
		return "install_in_progress", true
	case Result_LAUNCH_IN_PROGRESS:
		return "launch_in_progress", true
	case Result_MANIFEST_ERROR:
		return "manifest_error", true
	case Result_ICON_ERROR:
		return "icon_error", true
	case Result_INVALID_ICON_URL:
		return "invalid_icon_url", true
	case Result_ALREADY_INSTALLED:
		return "already_installed", true
	case Result_BLOCKED_FOR_CHILD_ACCOUNT:
		return "blocked_for_child_account", true
	default:
		return "", false
	}
}

type WebGlStatus uint32

const (
	_ WebGlStatus = iota

	WebGlStatus_WEBGL_ALLOWED
	WebGlStatus_WEBGL_BLOCKED
)

func (WebGlStatus) FromRef(str js.Ref) WebGlStatus {
	return WebGlStatus(bindings.ConstOfWebGlStatus(str))
}

func (x WebGlStatus) String() (string, bool) {
	switch x {
	case WebGlStatus_WEBGL_ALLOWED:
		return "webgl_allowed", true
	case WebGlStatus_WEBGL_BLOCKED:
		return "webgl_blocked", true
	default:
		return "", false
	}
}

// HasFuncBeginInstallWithManifest3 returns true if the function "WEBEXT.webstorePrivate.beginInstallWithManifest3" exists.
func HasFuncBeginInstallWithManifest3() bool {
	return js.True == bindings.HasFuncBeginInstallWithManifest3()
}

// FuncBeginInstallWithManifest3 returns the function "WEBEXT.webstorePrivate.beginInstallWithManifest3".
func FuncBeginInstallWithManifest3() (fn js.Func[func(details js.Any) js.Promise[Result]]) {
	bindings.FuncBeginInstallWithManifest3(
		js.Pointer(&fn),
	)
	return
}

// BeginInstallWithManifest3 calls the function "WEBEXT.webstorePrivate.beginInstallWithManifest3" directly.
func BeginInstallWithManifest3(details js.Any) (ret js.Promise[Result]) {
	bindings.CallBeginInstallWithManifest3(
		js.Pointer(&ret),
		details.Ref(),
	)

	return
}

// TryBeginInstallWithManifest3 calls the function "WEBEXT.webstorePrivate.beginInstallWithManifest3"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryBeginInstallWithManifest3(details js.Any) (ret js.Promise[Result], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBeginInstallWithManifest3(
		js.Pointer(&ret), js.Pointer(&exception),
		details.Ref(),
	)

	return
}

// HasFuncCompleteInstall returns true if the function "WEBEXT.webstorePrivate.completeInstall" exists.
func HasFuncCompleteInstall() bool {
	return js.True == bindings.HasFuncCompleteInstall()
}

// FuncCompleteInstall returns the function "WEBEXT.webstorePrivate.completeInstall".
func FuncCompleteInstall() (fn js.Func[func(expected_id js.String) js.Promise[js.Void]]) {
	bindings.FuncCompleteInstall(
		js.Pointer(&fn),
	)
	return
}

// CompleteInstall calls the function "WEBEXT.webstorePrivate.completeInstall" directly.
func CompleteInstall(expected_id js.String) (ret js.Promise[js.Void]) {
	bindings.CallCompleteInstall(
		js.Pointer(&ret),
		expected_id.Ref(),
	)

	return
}

// TryCompleteInstall calls the function "WEBEXT.webstorePrivate.completeInstall"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCompleteInstall(expected_id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompleteInstall(
		js.Pointer(&ret), js.Pointer(&exception),
		expected_id.Ref(),
	)

	return
}

// HasFuncEnableAppLauncher returns true if the function "WEBEXT.webstorePrivate.enableAppLauncher" exists.
func HasFuncEnableAppLauncher() bool {
	return js.True == bindings.HasFuncEnableAppLauncher()
}

// FuncEnableAppLauncher returns the function "WEBEXT.webstorePrivate.enableAppLauncher".
func FuncEnableAppLauncher() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncEnableAppLauncher(
		js.Pointer(&fn),
	)
	return
}

// EnableAppLauncher calls the function "WEBEXT.webstorePrivate.enableAppLauncher" directly.
func EnableAppLauncher() (ret js.Promise[js.Void]) {
	bindings.CallEnableAppLauncher(
		js.Pointer(&ret),
	)

	return
}

// TryEnableAppLauncher calls the function "WEBEXT.webstorePrivate.enableAppLauncher"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableAppLauncher() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableAppLauncher(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBrowserLogin returns true if the function "WEBEXT.webstorePrivate.getBrowserLogin" exists.
func HasFuncGetBrowserLogin() bool {
	return js.True == bindings.HasFuncGetBrowserLogin()
}

// FuncGetBrowserLogin returns the function "WEBEXT.webstorePrivate.getBrowserLogin".
func FuncGetBrowserLogin() (fn js.Func[func() js.Promise[GetBrowserLoginReturnType]]) {
	bindings.FuncGetBrowserLogin(
		js.Pointer(&fn),
	)
	return
}

// GetBrowserLogin calls the function "WEBEXT.webstorePrivate.getBrowserLogin" directly.
func GetBrowserLogin() (ret js.Promise[GetBrowserLoginReturnType]) {
	bindings.CallGetBrowserLogin(
		js.Pointer(&ret),
	)

	return
}

// TryGetBrowserLogin calls the function "WEBEXT.webstorePrivate.getBrowserLogin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBrowserLogin() (ret js.Promise[GetBrowserLoginReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBrowserLogin(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetExtensionStatus returns true if the function "WEBEXT.webstorePrivate.getExtensionStatus" exists.
func HasFuncGetExtensionStatus() bool {
	return js.True == bindings.HasFuncGetExtensionStatus()
}

// FuncGetExtensionStatus returns the function "WEBEXT.webstorePrivate.getExtensionStatus".
func FuncGetExtensionStatus() (fn js.Func[func(id js.String, manifest js.String) js.Promise[ExtensionInstallStatus]]) {
	bindings.FuncGetExtensionStatus(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionStatus calls the function "WEBEXT.webstorePrivate.getExtensionStatus" directly.
func GetExtensionStatus(id js.String, manifest js.String) (ret js.Promise[ExtensionInstallStatus]) {
	bindings.CallGetExtensionStatus(
		js.Pointer(&ret),
		id.Ref(),
		manifest.Ref(),
	)

	return
}

// TryGetExtensionStatus calls the function "WEBEXT.webstorePrivate.getExtensionStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionStatus(id js.String, manifest js.String) (ret js.Promise[ExtensionInstallStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		manifest.Ref(),
	)

	return
}

// HasFuncGetIsLauncherEnabled returns true if the function "WEBEXT.webstorePrivate.getIsLauncherEnabled" exists.
func HasFuncGetIsLauncherEnabled() bool {
	return js.True == bindings.HasFuncGetIsLauncherEnabled()
}

// FuncGetIsLauncherEnabled returns the function "WEBEXT.webstorePrivate.getIsLauncherEnabled".
func FuncGetIsLauncherEnabled() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncGetIsLauncherEnabled(
		js.Pointer(&fn),
	)
	return
}

// GetIsLauncherEnabled calls the function "WEBEXT.webstorePrivate.getIsLauncherEnabled" directly.
func GetIsLauncherEnabled() (ret js.Promise[js.Boolean]) {
	bindings.CallGetIsLauncherEnabled(
		js.Pointer(&ret),
	)

	return
}

// TryGetIsLauncherEnabled calls the function "WEBEXT.webstorePrivate.getIsLauncherEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIsLauncherEnabled() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIsLauncherEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetReferrerChain returns true if the function "WEBEXT.webstorePrivate.getReferrerChain" exists.
func HasFuncGetReferrerChain() bool {
	return js.True == bindings.HasFuncGetReferrerChain()
}

// FuncGetReferrerChain returns the function "WEBEXT.webstorePrivate.getReferrerChain".
func FuncGetReferrerChain() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetReferrerChain(
		js.Pointer(&fn),
	)
	return
}

// GetReferrerChain calls the function "WEBEXT.webstorePrivate.getReferrerChain" directly.
func GetReferrerChain() (ret js.Promise[js.String]) {
	bindings.CallGetReferrerChain(
		js.Pointer(&ret),
	)

	return
}

// TryGetReferrerChain calls the function "WEBEXT.webstorePrivate.getReferrerChain"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetReferrerChain() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetReferrerChain(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetStoreLogin returns true if the function "WEBEXT.webstorePrivate.getStoreLogin" exists.
func HasFuncGetStoreLogin() bool {
	return js.True == bindings.HasFuncGetStoreLogin()
}

// FuncGetStoreLogin returns the function "WEBEXT.webstorePrivate.getStoreLogin".
func FuncGetStoreLogin() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetStoreLogin(
		js.Pointer(&fn),
	)
	return
}

// GetStoreLogin calls the function "WEBEXT.webstorePrivate.getStoreLogin" directly.
func GetStoreLogin() (ret js.Promise[js.String]) {
	bindings.CallGetStoreLogin(
		js.Pointer(&ret),
	)

	return
}

// TryGetStoreLogin calls the function "WEBEXT.webstorePrivate.getStoreLogin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetStoreLogin() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetStoreLogin(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetWebGLStatus returns true if the function "WEBEXT.webstorePrivate.getWebGLStatus" exists.
func HasFuncGetWebGLStatus() bool {
	return js.True == bindings.HasFuncGetWebGLStatus()
}

// FuncGetWebGLStatus returns the function "WEBEXT.webstorePrivate.getWebGLStatus".
func FuncGetWebGLStatus() (fn js.Func[func() js.Promise[WebGlStatus]]) {
	bindings.FuncGetWebGLStatus(
		js.Pointer(&fn),
	)
	return
}

// GetWebGLStatus calls the function "WEBEXT.webstorePrivate.getWebGLStatus" directly.
func GetWebGLStatus() (ret js.Promise[WebGlStatus]) {
	bindings.CallGetWebGLStatus(
		js.Pointer(&ret),
	)

	return
}

// TryGetWebGLStatus calls the function "WEBEXT.webstorePrivate.getWebGLStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetWebGLStatus() (ret js.Promise[WebGlStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetWebGLStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInstall returns true if the function "WEBEXT.webstorePrivate.install" exists.
func HasFuncInstall() bool {
	return js.True == bindings.HasFuncInstall()
}

// FuncInstall returns the function "WEBEXT.webstorePrivate.install".
func FuncInstall() (fn js.Func[func(expected_id js.String) js.Promise[js.Void]]) {
	bindings.FuncInstall(
		js.Pointer(&fn),
	)
	return
}

// Install calls the function "WEBEXT.webstorePrivate.install" directly.
func Install(expected_id js.String) (ret js.Promise[js.Void]) {
	bindings.CallInstall(
		js.Pointer(&ret),
		expected_id.Ref(),
	)

	return
}

// TryInstall calls the function "WEBEXT.webstorePrivate.install"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInstall(expected_id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInstall(
		js.Pointer(&ret), js.Pointer(&exception),
		expected_id.Ref(),
	)

	return
}

// HasFuncIsInIncognitoMode returns true if the function "WEBEXT.webstorePrivate.isInIncognitoMode" exists.
func HasFuncIsInIncognitoMode() bool {
	return js.True == bindings.HasFuncIsInIncognitoMode()
}

// FuncIsInIncognitoMode returns the function "WEBEXT.webstorePrivate.isInIncognitoMode".
func FuncIsInIncognitoMode() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsInIncognitoMode(
		js.Pointer(&fn),
	)
	return
}

// IsInIncognitoMode calls the function "WEBEXT.webstorePrivate.isInIncognitoMode" directly.
func IsInIncognitoMode() (ret js.Promise[js.Boolean]) {
	bindings.CallIsInIncognitoMode(
		js.Pointer(&ret),
	)

	return
}

// TryIsInIncognitoMode calls the function "WEBEXT.webstorePrivate.isInIncognitoMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsInIncognitoMode() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsInIncognitoMode(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsPendingCustodianApproval returns true if the function "WEBEXT.webstorePrivate.isPendingCustodianApproval" exists.
func HasFuncIsPendingCustodianApproval() bool {
	return js.True == bindings.HasFuncIsPendingCustodianApproval()
}

// FuncIsPendingCustodianApproval returns the function "WEBEXT.webstorePrivate.isPendingCustodianApproval".
func FuncIsPendingCustodianApproval() (fn js.Func[func(id js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsPendingCustodianApproval(
		js.Pointer(&fn),
	)
	return
}

// IsPendingCustodianApproval calls the function "WEBEXT.webstorePrivate.isPendingCustodianApproval" directly.
func IsPendingCustodianApproval(id js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsPendingCustodianApproval(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryIsPendingCustodianApproval calls the function "WEBEXT.webstorePrivate.isPendingCustodianApproval"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsPendingCustodianApproval(id js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsPendingCustodianApproval(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncSetStoreLogin returns true if the function "WEBEXT.webstorePrivate.setStoreLogin" exists.
func HasFuncSetStoreLogin() bool {
	return js.True == bindings.HasFuncSetStoreLogin()
}

// FuncSetStoreLogin returns the function "WEBEXT.webstorePrivate.setStoreLogin".
func FuncSetStoreLogin() (fn js.Func[func(login js.String) js.Promise[js.Void]]) {
	bindings.FuncSetStoreLogin(
		js.Pointer(&fn),
	)
	return
}

// SetStoreLogin calls the function "WEBEXT.webstorePrivate.setStoreLogin" directly.
func SetStoreLogin(login js.String) (ret js.Promise[js.Void]) {
	bindings.CallSetStoreLogin(
		js.Pointer(&ret),
		login.Ref(),
	)

	return
}

// TrySetStoreLogin calls the function "WEBEXT.webstorePrivate.setStoreLogin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetStoreLogin(login js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetStoreLogin(
		js.Pointer(&ret), js.Pointer(&exception),
		login.Ref(),
	)

	return
}

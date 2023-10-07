// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package dashboardprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/dashboardprivate/bindings"
)

type Result uint32

const (
	_ Result = iota

	Result_
	Result_UNKNOWN_ERROR
	Result_USER_CANCELLED
	Result_INVALID_ID
	Result_MANIFEST_ERROR
	Result_ICON_ERROR
	Result_INVALID_ICON_URL
)

func (Result) FromRef(str js.Ref) Result {
	return Result(bindings.ConstOfResult(str))
}

func (x Result) String() (string, bool) {
	switch x {
	case Result_:
		return "", true
	case Result_UNKNOWN_ERROR:
		return "unknown_error", true
	case Result_USER_CANCELLED:
		return "user_cancelled", true
	case Result_INVALID_ID:
		return "invalid_id", true
	case Result_MANIFEST_ERROR:
		return "manifest_error", true
	case Result_ICON_ERROR:
		return "icon_error", true
	case Result_INVALID_ICON_URL:
		return "invalid_icon_url", true
	default:
		return "", false
	}
}

type ShowPermissionPromptForDelegatedInstallArgDetails struct {
	// DelegatedUser is "ShowPermissionPromptForDelegatedInstallArgDetails.delegatedUser"
	//
	// Required
	DelegatedUser js.String
	// IconUrl is "ShowPermissionPromptForDelegatedInstallArgDetails.iconUrl"
	//
	// Optional
	IconUrl js.String
	// Id is "ShowPermissionPromptForDelegatedInstallArgDetails.id"
	//
	// Required
	Id js.String
	// LocalizedName is "ShowPermissionPromptForDelegatedInstallArgDetails.localizedName"
	//
	// Optional
	LocalizedName js.String
	// Manifest is "ShowPermissionPromptForDelegatedInstallArgDetails.manifest"
	//
	// Required
	Manifest js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShowPermissionPromptForDelegatedInstallArgDetails with all fields set.
func (p ShowPermissionPromptForDelegatedInstallArgDetails) FromRef(ref js.Ref) ShowPermissionPromptForDelegatedInstallArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShowPermissionPromptForDelegatedInstallArgDetails in the application heap.
func (p ShowPermissionPromptForDelegatedInstallArgDetails) New() js.Ref {
	return bindings.ShowPermissionPromptForDelegatedInstallArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShowPermissionPromptForDelegatedInstallArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ShowPermissionPromptForDelegatedInstallArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShowPermissionPromptForDelegatedInstallArgDetails) Update(ref js.Ref) {
	bindings.ShowPermissionPromptForDelegatedInstallArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShowPermissionPromptForDelegatedInstallArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DelegatedUser.Ref(),
		p.IconUrl.Ref(),
		p.Id.Ref(),
		p.LocalizedName.Ref(),
		p.Manifest.Ref(),
	)
	p.DelegatedUser = p.DelegatedUser.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.LocalizedName = p.LocalizedName.FromRef(js.Undefined)
	p.Manifest = p.Manifest.FromRef(js.Undefined)
}

// HasFuncShowPermissionPromptForDelegatedInstall returns true if the function "WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall" exists.
func HasFuncShowPermissionPromptForDelegatedInstall() bool {
	return js.True == bindings.HasFuncShowPermissionPromptForDelegatedInstall()
}

// FuncShowPermissionPromptForDelegatedInstall returns the function "WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall".
func FuncShowPermissionPromptForDelegatedInstall() (fn js.Func[func(details ShowPermissionPromptForDelegatedInstallArgDetails) js.Promise[Result]]) {
	bindings.FuncShowPermissionPromptForDelegatedInstall(
		js.Pointer(&fn),
	)
	return
}

// ShowPermissionPromptForDelegatedInstall calls the function "WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall" directly.
func ShowPermissionPromptForDelegatedInstall(details ShowPermissionPromptForDelegatedInstallArgDetails) (ret js.Promise[Result]) {
	bindings.CallShowPermissionPromptForDelegatedInstall(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryShowPermissionPromptForDelegatedInstall calls the function "WEBEXT.dashboardPrivate.showPermissionPromptForDelegatedInstall"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowPermissionPromptForDelegatedInstall(details ShowPermissionPromptForDelegatedInstallArgDetails) (ret js.Promise[Result], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowPermissionPromptForDelegatedInstall(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

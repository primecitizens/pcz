// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package systemprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/systemprivate/bindings"
)

type GetIncognitoModeAvailabilityValue uint32

const (
	_ GetIncognitoModeAvailabilityValue = iota

	GetIncognitoModeAvailabilityValue_ENABLED
	GetIncognitoModeAvailabilityValue_DISABLED
	GetIncognitoModeAvailabilityValue_FORCED
)

func (GetIncognitoModeAvailabilityValue) FromRef(str js.Ref) GetIncognitoModeAvailabilityValue {
	return GetIncognitoModeAvailabilityValue(bindings.ConstOfGetIncognitoModeAvailabilityValue(str))
}

func (x GetIncognitoModeAvailabilityValue) String() (string, bool) {
	switch x {
	case GetIncognitoModeAvailabilityValue_ENABLED:
		return "enabled", true
	case GetIncognitoModeAvailabilityValue_DISABLED:
		return "disabled", true
	case GetIncognitoModeAvailabilityValue_FORCED:
		return "forced", true
	default:
		return "", false
	}
}

type UpdateStatusState uint32

const (
	_ UpdateStatusState = iota

	UpdateStatusState_NOT_AVAILABLE
	UpdateStatusState_UPDATING
	UpdateStatusState_NEED_RESTART
)

func (UpdateStatusState) FromRef(str js.Ref) UpdateStatusState {
	return UpdateStatusState(bindings.ConstOfUpdateStatusState(str))
}

func (x UpdateStatusState) String() (string, bool) {
	switch x {
	case UpdateStatusState_NOT_AVAILABLE:
		return "NotAvailable", true
	case UpdateStatusState_UPDATING:
		return "Updating", true
	case UpdateStatusState_NEED_RESTART:
		return "NeedRestart", true
	default:
		return "", false
	}
}

type UpdateStatus struct {
	// DownloadProgress is "UpdateStatus.downloadProgress"
	//
	// Required
	DownloadProgress float64
	// State is "UpdateStatus.state"
	//
	// Required
	State UpdateStatusState

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UpdateStatus with all fields set.
func (p UpdateStatus) FromRef(ref js.Ref) UpdateStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UpdateStatus in the application heap.
func (p UpdateStatus) New() js.Ref {
	return bindings.UpdateStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UpdateStatus) UpdateFrom(ref js.Ref) {
	bindings.UpdateStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UpdateStatus) Update(ref js.Ref) {
	bindings.UpdateStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UpdateStatus) FreeMembers(recursive bool) {
}

// HasFuncGetApiKey returns true if the function "WEBEXT.systemPrivate.getApiKey" exists.
func HasFuncGetApiKey() bool {
	return js.True == bindings.HasFuncGetApiKey()
}

// FuncGetApiKey returns the function "WEBEXT.systemPrivate.getApiKey".
func FuncGetApiKey() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetApiKey(
		js.Pointer(&fn),
	)
	return
}

// GetApiKey calls the function "WEBEXT.systemPrivate.getApiKey" directly.
func GetApiKey() (ret js.Promise[js.String]) {
	bindings.CallGetApiKey(
		js.Pointer(&ret),
	)

	return
}

// TryGetApiKey calls the function "WEBEXT.systemPrivate.getApiKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetApiKey() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetApiKey(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetIncognitoModeAvailability returns true if the function "WEBEXT.systemPrivate.getIncognitoModeAvailability" exists.
func HasFuncGetIncognitoModeAvailability() bool {
	return js.True == bindings.HasFuncGetIncognitoModeAvailability()
}

// FuncGetIncognitoModeAvailability returns the function "WEBEXT.systemPrivate.getIncognitoModeAvailability".
func FuncGetIncognitoModeAvailability() (fn js.Func[func() js.Promise[GetIncognitoModeAvailabilityValue]]) {
	bindings.FuncGetIncognitoModeAvailability(
		js.Pointer(&fn),
	)
	return
}

// GetIncognitoModeAvailability calls the function "WEBEXT.systemPrivate.getIncognitoModeAvailability" directly.
func GetIncognitoModeAvailability() (ret js.Promise[GetIncognitoModeAvailabilityValue]) {
	bindings.CallGetIncognitoModeAvailability(
		js.Pointer(&ret),
	)

	return
}

// TryGetIncognitoModeAvailability calls the function "WEBEXT.systemPrivate.getIncognitoModeAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIncognitoModeAvailability() (ret js.Promise[GetIncognitoModeAvailabilityValue], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIncognitoModeAvailability(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUpdateStatus returns true if the function "WEBEXT.systemPrivate.getUpdateStatus" exists.
func HasFuncGetUpdateStatus() bool {
	return js.True == bindings.HasFuncGetUpdateStatus()
}

// FuncGetUpdateStatus returns the function "WEBEXT.systemPrivate.getUpdateStatus".
func FuncGetUpdateStatus() (fn js.Func[func() js.Promise[UpdateStatus]]) {
	bindings.FuncGetUpdateStatus(
		js.Pointer(&fn),
	)
	return
}

// GetUpdateStatus calls the function "WEBEXT.systemPrivate.getUpdateStatus" directly.
func GetUpdateStatus() (ret js.Promise[UpdateStatus]) {
	bindings.CallGetUpdateStatus(
		js.Pointer(&ret),
	)

	return
}

// TryGetUpdateStatus calls the function "WEBEXT.systemPrivate.getUpdateStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUpdateStatus() (ret js.Promise[UpdateStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUpdateStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

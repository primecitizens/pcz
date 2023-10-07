// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package reportingprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/reportingprivate/bindings"
)

type AntiVirusProductState uint32

const (
	_ AntiVirusProductState = iota

	AntiVirusProductState_ON
	AntiVirusProductState_OFF
	AntiVirusProductState_SNOOZED
	AntiVirusProductState_EXPIRED
)

func (AntiVirusProductState) FromRef(str js.Ref) AntiVirusProductState {
	return AntiVirusProductState(bindings.ConstOfAntiVirusProductState(str))
}

func (x AntiVirusProductState) String() (string, bool) {
	switch x {
	case AntiVirusProductState_ON:
		return "ON", true
	case AntiVirusProductState_OFF:
		return "OFF", true
	case AntiVirusProductState_SNOOZED:
		return "SNOOZED", true
	case AntiVirusProductState_EXPIRED:
		return "EXPIRED", true
	default:
		return "", false
	}
}

type AntiVirusSignal struct {
	// DisplayName is "AntiVirusSignal.displayName"
	//
	// Optional
	DisplayName js.String
	// ProductId is "AntiVirusSignal.productId"
	//
	// Optional
	ProductId js.String
	// State is "AntiVirusSignal.state"
	//
	// Optional
	State AntiVirusProductState

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AntiVirusSignal with all fields set.
func (p AntiVirusSignal) FromRef(ref js.Ref) AntiVirusSignal {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AntiVirusSignal in the application heap.
func (p AntiVirusSignal) New() js.Ref {
	return bindings.AntiVirusSignalJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AntiVirusSignal) UpdateFrom(ref js.Ref) {
	bindings.AntiVirusSignalJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AntiVirusSignal) Update(ref js.Ref) {
	bindings.AntiVirusSignalJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AntiVirusSignal) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayName.Ref(),
		p.ProductId.Ref(),
	)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.ProductId = p.ProductId.FromRef(js.Undefined)
}

type AvInfoCallbackFunc func(this js.Ref, avSignals js.Array[AntiVirusSignal]) js.Ref

func (fn AvInfoCallbackFunc) Register() js.Func[func(avSignals js.Array[AntiVirusSignal])] {
	return js.RegisterCallback[func(avSignals js.Array[AntiVirusSignal])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AvInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AntiVirusSignal]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AvInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, avSignals js.Array[AntiVirusSignal]) js.Ref
	Arg T
}

func (cb *AvInfoCallback[T]) Register() js.Func[func(avSignals js.Array[AntiVirusSignal])] {
	return js.RegisterCallback[func(avSignals js.Array[AntiVirusSignal])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AvInfoCallback[T]) DispatchCallback(
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

		js.Array[AntiVirusSignal]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CertificateStatus uint32

const (
	_ CertificateStatus = iota

	CertificateStatus_OK
	CertificateStatus_POLICY_UNSET
)

func (CertificateStatus) FromRef(str js.Ref) CertificateStatus {
	return CertificateStatus(bindings.ConstOfCertificateStatus(str))
}

func (x CertificateStatus) String() (string, bool) {
	switch x {
	case CertificateStatus_OK:
		return "OK", true
	case CertificateStatus_POLICY_UNSET:
		return "POLICY_UNSET", true
	default:
		return "", false
	}
}

type Certificate struct {
	// Status is "Certificate.status"
	//
	// Optional
	Status CertificateStatus
	// EncodedCertificate is "Certificate.encodedCertificate"
	//
	// Optional
	EncodedCertificate js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Certificate with all fields set.
func (p Certificate) FromRef(ref js.Ref) Certificate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Certificate in the application heap.
func (p Certificate) New() js.Ref {
	return bindings.CertificateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Certificate) UpdateFrom(ref js.Ref) {
	bindings.CertificateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Certificate) Update(ref js.Ref) {
	bindings.CertificateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Certificate) FreeMembers(recursive bool) {
	js.Free(
		p.EncodedCertificate.Ref(),
	)
	p.EncodedCertificate = p.EncodedCertificate.FromRef(js.Undefined)
}

type CertificateCallbackFunc func(this js.Ref, certificate *Certificate) js.Ref

func (fn CertificateCallbackFunc) Register() js.Func[func(certificate *Certificate)] {
	return js.RegisterCallback[func(certificate *Certificate)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CertificateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Certificate
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

type CertificateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, certificate *Certificate) js.Ref
	Arg T
}

func (cb *CertificateCallback[T]) Register() js.Func[func(certificate *Certificate)] {
	return js.RegisterCallback[func(certificate *Certificate)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CertificateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Certificate
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

type RealtimeUrlCheckMode uint32

const (
	_ RealtimeUrlCheckMode = iota

	RealtimeUrlCheckMode_DISABLED
	RealtimeUrlCheckMode_ENABLED_MAIN_FRAME
)

func (RealtimeUrlCheckMode) FromRef(str js.Ref) RealtimeUrlCheckMode {
	return RealtimeUrlCheckMode(bindings.ConstOfRealtimeUrlCheckMode(str))
}

func (x RealtimeUrlCheckMode) String() (string, bool) {
	switch x {
	case RealtimeUrlCheckMode_DISABLED:
		return "DISABLED", true
	case RealtimeUrlCheckMode_ENABLED_MAIN_FRAME:
		return "ENABLED_MAIN_FRAME", true
	default:
		return "", false
	}
}

type SafeBrowsingLevel uint32

const (
	_ SafeBrowsingLevel = iota

	SafeBrowsingLevel_DISABLED
	SafeBrowsingLevel_STANDARD
	SafeBrowsingLevel_ENHANCED
)

func (SafeBrowsingLevel) FromRef(str js.Ref) SafeBrowsingLevel {
	return SafeBrowsingLevel(bindings.ConstOfSafeBrowsingLevel(str))
}

func (x SafeBrowsingLevel) String() (string, bool) {
	switch x {
	case SafeBrowsingLevel_DISABLED:
		return "DISABLED", true
	case SafeBrowsingLevel_STANDARD:
		return "STANDARD", true
	case SafeBrowsingLevel_ENHANCED:
		return "ENHANCED", true
	default:
		return "", false
	}
}

type PasswordProtectionTrigger uint32

const (
	_ PasswordProtectionTrigger = iota

	PasswordProtectionTrigger_PASSWORD_PROTECTION_OFF
	PasswordProtectionTrigger_PASSWORD_REUSE
	PasswordProtectionTrigger_PHISHING_REUSE
	PasswordProtectionTrigger_POLICY_UNSET
)

func (PasswordProtectionTrigger) FromRef(str js.Ref) PasswordProtectionTrigger {
	return PasswordProtectionTrigger(bindings.ConstOfPasswordProtectionTrigger(str))
}

func (x PasswordProtectionTrigger) String() (string, bool) {
	switch x {
	case PasswordProtectionTrigger_PASSWORD_PROTECTION_OFF:
		return "PASSWORD_PROTECTION_OFF", true
	case PasswordProtectionTrigger_PASSWORD_REUSE:
		return "PASSWORD_REUSE", true
	case PasswordProtectionTrigger_PHISHING_REUSE:
		return "PHISHING_REUSE", true
	case PasswordProtectionTrigger_POLICY_UNSET:
		return "POLICY_UNSET", true
	default:
		return "", false
	}
}

type SettingValue uint32

const (
	_ SettingValue = iota

	SettingValue_UNKNOWN
	SettingValue_DISABLED
	SettingValue_ENABLED
)

func (SettingValue) FromRef(str js.Ref) SettingValue {
	return SettingValue(bindings.ConstOfSettingValue(str))
}

func (x SettingValue) String() (string, bool) {
	switch x {
	case SettingValue_UNKNOWN:
		return "UNKNOWN", true
	case SettingValue_DISABLED:
		return "DISABLED", true
	case SettingValue_ENABLED:
		return "ENABLED", true
	default:
		return "", false
	}
}

type ContextInfo struct {
	// BrowserAffiliationIds is "ContextInfo.browserAffiliationIds"
	//
	// Optional
	BrowserAffiliationIds js.Array[js.String]
	// ProfileAffiliationIds is "ContextInfo.profileAffiliationIds"
	//
	// Optional
	ProfileAffiliationIds js.Array[js.String]
	// OnFileAttachedProviders is "ContextInfo.onFileAttachedProviders"
	//
	// Optional
	OnFileAttachedProviders js.Array[js.String]
	// OnFileDownloadedProviders is "ContextInfo.onFileDownloadedProviders"
	//
	// Optional
	OnFileDownloadedProviders js.Array[js.String]
	// OnBulkDataEntryProviders is "ContextInfo.onBulkDataEntryProviders"
	//
	// Optional
	OnBulkDataEntryProviders js.Array[js.String]
	// OnPrintProviders is "ContextInfo.onPrintProviders"
	//
	// Optional
	OnPrintProviders js.Array[js.String]
	// RealtimeUrlCheckMode is "ContextInfo.realtimeUrlCheckMode"
	//
	// Optional
	RealtimeUrlCheckMode RealtimeUrlCheckMode
	// OnSecurityEventProviders is "ContextInfo.onSecurityEventProviders"
	//
	// Optional
	OnSecurityEventProviders js.Array[js.String]
	// BrowserVersion is "ContextInfo.browserVersion"
	//
	// Optional
	BrowserVersion js.String
	// SafeBrowsingProtectionLevel is "ContextInfo.safeBrowsingProtectionLevel"
	//
	// Optional
	SafeBrowsingProtectionLevel SafeBrowsingLevel
	// SiteIsolationEnabled is "ContextInfo.siteIsolationEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_SiteIsolationEnabled MUST be set to true to make this field effective.
	SiteIsolationEnabled bool
	// BuiltInDnsClientEnabled is "ContextInfo.builtInDnsClientEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_BuiltInDnsClientEnabled MUST be set to true to make this field effective.
	BuiltInDnsClientEnabled bool
	// PasswordProtectionWarningTrigger is "ContextInfo.passwordProtectionWarningTrigger"
	//
	// Optional
	PasswordProtectionWarningTrigger PasswordProtectionTrigger
	// ChromeRemoteDesktopAppBlocked is "ContextInfo.chromeRemoteDesktopAppBlocked"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChromeRemoteDesktopAppBlocked MUST be set to true to make this field effective.
	ChromeRemoteDesktopAppBlocked bool
	// ThirdPartyBlockingEnabled is "ContextInfo.thirdPartyBlockingEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_ThirdPartyBlockingEnabled MUST be set to true to make this field effective.
	ThirdPartyBlockingEnabled bool
	// OsFirewall is "ContextInfo.osFirewall"
	//
	// Optional
	OsFirewall SettingValue
	// SystemDnsServers is "ContextInfo.systemDnsServers"
	//
	// Optional
	SystemDnsServers js.Array[js.String]
	// EnterpriseProfileId is "ContextInfo.enterpriseProfileId"
	//
	// Optional
	EnterpriseProfileId js.String

	FFI_USE_SiteIsolationEnabled          bool // for SiteIsolationEnabled.
	FFI_USE_BuiltInDnsClientEnabled       bool // for BuiltInDnsClientEnabled.
	FFI_USE_ChromeRemoteDesktopAppBlocked bool // for ChromeRemoteDesktopAppBlocked.
	FFI_USE_ThirdPartyBlockingEnabled     bool // for ThirdPartyBlockingEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextInfo with all fields set.
func (p ContextInfo) FromRef(ref js.Ref) ContextInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextInfo in the application heap.
func (p ContextInfo) New() js.Ref {
	return bindings.ContextInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextInfo) UpdateFrom(ref js.Ref) {
	bindings.ContextInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextInfo) Update(ref js.Ref) {
	bindings.ContextInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextInfo) FreeMembers(recursive bool) {
	js.Free(
		p.BrowserAffiliationIds.Ref(),
		p.ProfileAffiliationIds.Ref(),
		p.OnFileAttachedProviders.Ref(),
		p.OnFileDownloadedProviders.Ref(),
		p.OnBulkDataEntryProviders.Ref(),
		p.OnPrintProviders.Ref(),
		p.OnSecurityEventProviders.Ref(),
		p.BrowserVersion.Ref(),
		p.SystemDnsServers.Ref(),
		p.EnterpriseProfileId.Ref(),
	)
	p.BrowserAffiliationIds = p.BrowserAffiliationIds.FromRef(js.Undefined)
	p.ProfileAffiliationIds = p.ProfileAffiliationIds.FromRef(js.Undefined)
	p.OnFileAttachedProviders = p.OnFileAttachedProviders.FromRef(js.Undefined)
	p.OnFileDownloadedProviders = p.OnFileDownloadedProviders.FromRef(js.Undefined)
	p.OnBulkDataEntryProviders = p.OnBulkDataEntryProviders.FromRef(js.Undefined)
	p.OnPrintProviders = p.OnPrintProviders.FromRef(js.Undefined)
	p.OnSecurityEventProviders = p.OnSecurityEventProviders.FromRef(js.Undefined)
	p.BrowserVersion = p.BrowserVersion.FromRef(js.Undefined)
	p.SystemDnsServers = p.SystemDnsServers.FromRef(js.Undefined)
	p.EnterpriseProfileId = p.EnterpriseProfileId.FromRef(js.Undefined)
}

type DeviceInfo struct {
	// OsName is "DeviceInfo.osName"
	//
	// Optional
	OsName js.String
	// OsVersion is "DeviceInfo.osVersion"
	//
	// Optional
	OsVersion js.String
	// DeviceHostName is "DeviceInfo.deviceHostName"
	//
	// Optional
	DeviceHostName js.String
	// DeviceModel is "DeviceInfo.deviceModel"
	//
	// Optional
	DeviceModel js.String
	// SerialNumber is "DeviceInfo.serialNumber"
	//
	// Optional
	SerialNumber js.String
	// ScreenLockSecured is "DeviceInfo.screenLockSecured"
	//
	// Optional
	ScreenLockSecured SettingValue
	// DiskEncrypted is "DeviceInfo.diskEncrypted"
	//
	// Optional
	DiskEncrypted SettingValue
	// MacAddresses is "DeviceInfo.macAddresses"
	//
	// Optional
	MacAddresses js.Array[js.String]
	// WindowsMachineDomain is "DeviceInfo.windowsMachineDomain"
	//
	// Optional
	WindowsMachineDomain js.String
	// WindowsUserDomain is "DeviceInfo.windowsUserDomain"
	//
	// Optional
	WindowsUserDomain js.String
	// SecurityPatchLevel is "DeviceInfo.securityPatchLevel"
	//
	// Optional
	SecurityPatchLevel js.String
	// SecureBootEnabled is "DeviceInfo.secureBootEnabled"
	//
	// Optional
	SecureBootEnabled SettingValue

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceInfo with all fields set.
func (p DeviceInfo) FromRef(ref js.Ref) DeviceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceInfo in the application heap.
func (p DeviceInfo) New() js.Ref {
	return bindings.DeviceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceInfo) UpdateFrom(ref js.Ref) {
	bindings.DeviceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceInfo) Update(ref js.Ref) {
	bindings.DeviceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.OsName.Ref(),
		p.OsVersion.Ref(),
		p.DeviceHostName.Ref(),
		p.DeviceModel.Ref(),
		p.SerialNumber.Ref(),
		p.MacAddresses.Ref(),
		p.WindowsMachineDomain.Ref(),
		p.WindowsUserDomain.Ref(),
		p.SecurityPatchLevel.Ref(),
	)
	p.OsName = p.OsName.FromRef(js.Undefined)
	p.OsVersion = p.OsVersion.FromRef(js.Undefined)
	p.DeviceHostName = p.DeviceHostName.FromRef(js.Undefined)
	p.DeviceModel = p.DeviceModel.FromRef(js.Undefined)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
	p.MacAddresses = p.MacAddresses.FromRef(js.Undefined)
	p.WindowsMachineDomain = p.WindowsMachineDomain.FromRef(js.Undefined)
	p.WindowsUserDomain = p.WindowsUserDomain.FromRef(js.Undefined)
	p.SecurityPatchLevel = p.SecurityPatchLevel.FromRef(js.Undefined)
}

type DoneCallbackFunc func(this js.Ref) js.Ref

func (fn DoneCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DoneCallbackFunc) DispatchCallback(
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

type DoneCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DoneCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DoneCallback[T]) DispatchCallback(
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

type EventType uint32

const (
	_ EventType = iota

	EventType_DEVICE
	EventType_USER
)

func (EventType) FromRef(str js.Ref) EventType {
	return EventType(bindings.ConstOfEventType(str))
}

func (x EventType) String() (string, bool) {
	switch x {
	case EventType_DEVICE:
		return "DEVICE", true
	case EventType_USER:
		return "USER", true
	default:
		return "", false
	}
}

type EnqueueRecordRequest struct {
	// RecordData is "EnqueueRecordRequest.recordData"
	//
	// Optional
	RecordData js.ArrayBufferView
	// Priority is "EnqueueRecordRequest.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// EventType is "EnqueueRecordRequest.eventType"
	//
	// Optional
	EventType EventType

	FFI_USE_Priority bool // for Priority.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EnqueueRecordRequest with all fields set.
func (p EnqueueRecordRequest) FromRef(ref js.Ref) EnqueueRecordRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EnqueueRecordRequest in the application heap.
func (p EnqueueRecordRequest) New() js.Ref {
	return bindings.EnqueueRecordRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EnqueueRecordRequest) UpdateFrom(ref js.Ref) {
	bindings.EnqueueRecordRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EnqueueRecordRequest) Update(ref js.Ref) {
	bindings.EnqueueRecordRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EnqueueRecordRequest) FreeMembers(recursive bool) {
	js.Free(
		p.RecordData.Ref(),
	)
	p.RecordData = p.RecordData.FromRef(js.Undefined)
}

type FileSystemInfoCallbackFunc func(this js.Ref, fileSystemSignals js.Array[GetFileSystemInfoResponse]) js.Ref

func (fn FileSystemInfoCallbackFunc) Register() js.Func[func(fileSystemSignals js.Array[GetFileSystemInfoResponse])] {
	return js.RegisterCallback[func(fileSystemSignals js.Array[GetFileSystemInfoResponse])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FileSystemInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[GetFileSystemInfoResponse]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileSystemInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, fileSystemSignals js.Array[GetFileSystemInfoResponse]) js.Ref
	Arg T
}

func (cb *FileSystemInfoCallback[T]) Register() js.Func[func(fileSystemSignals js.Array[GetFileSystemInfoResponse])] {
	return js.RegisterCallback[func(fileSystemSignals js.Array[GetFileSystemInfoResponse])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FileSystemInfoCallback[T]) DispatchCallback(
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

		js.Array[GetFileSystemInfoResponse]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PresenceValue uint32

const (
	_ PresenceValue = iota

	PresenceValue_UNSPECIFIED
	PresenceValue_ACCESS_DENIED
	PresenceValue_NOT_FOUND
	PresenceValue_FOUND
)

func (PresenceValue) FromRef(str js.Ref) PresenceValue {
	return PresenceValue(bindings.ConstOfPresenceValue(str))
}

func (x PresenceValue) String() (string, bool) {
	switch x {
	case PresenceValue_UNSPECIFIED:
		return "UNSPECIFIED", true
	case PresenceValue_ACCESS_DENIED:
		return "ACCESS_DENIED", true
	case PresenceValue_NOT_FOUND:
		return "NOT_FOUND", true
	case PresenceValue_FOUND:
		return "FOUND", true
	default:
		return "", false
	}
}

type GetFileSystemInfoResponse struct {
	// Path is "GetFileSystemInfoResponse.path"
	//
	// Optional
	Path js.String
	// Presence is "GetFileSystemInfoResponse.presence"
	//
	// Optional
	Presence PresenceValue
	// Sha256Hash is "GetFileSystemInfoResponse.sha256Hash"
	//
	// Optional
	Sha256Hash js.String
	// IsRunning is "GetFileSystemInfoResponse.isRunning"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRunning MUST be set to true to make this field effective.
	IsRunning bool
	// PublicKeysHashes is "GetFileSystemInfoResponse.publicKeysHashes"
	//
	// Optional
	PublicKeysHashes js.Array[js.String]
	// ProductName is "GetFileSystemInfoResponse.productName"
	//
	// Optional
	ProductName js.String
	// Version is "GetFileSystemInfoResponse.version"
	//
	// Optional
	Version js.String

	FFI_USE_IsRunning bool // for IsRunning.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFileSystemInfoResponse with all fields set.
func (p GetFileSystemInfoResponse) FromRef(ref js.Ref) GetFileSystemInfoResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFileSystemInfoResponse in the application heap.
func (p GetFileSystemInfoResponse) New() js.Ref {
	return bindings.GetFileSystemInfoResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFileSystemInfoResponse) UpdateFrom(ref js.Ref) {
	bindings.GetFileSystemInfoResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFileSystemInfoResponse) Update(ref js.Ref) {
	bindings.GetFileSystemInfoResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFileSystemInfoResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
		p.Sha256Hash.Ref(),
		p.PublicKeysHashes.Ref(),
		p.ProductName.Ref(),
		p.Version.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Sha256Hash = p.Sha256Hash.FromRef(js.Undefined)
	p.PublicKeysHashes = p.PublicKeysHashes.FromRef(js.Undefined)
	p.ProductName = p.ProductName.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
}

type GetContextInfoCallbackFunc func(this js.Ref, contextInfo *ContextInfo) js.Ref

func (fn GetContextInfoCallbackFunc) Register() js.Func[func(contextInfo *ContextInfo)] {
	return js.RegisterCallback[func(contextInfo *ContextInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetContextInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ContextInfo
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

type GetContextInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, contextInfo *ContextInfo) js.Ref
	Arg T
}

func (cb *GetContextInfoCallback[T]) Register() js.Func[func(contextInfo *ContextInfo)] {
	return js.RegisterCallback[func(contextInfo *ContextInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetContextInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ContextInfo
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

type GetDeviceDataCallbackFunc func(this js.Ref, data js.ArrayBuffer) js.Ref

func (fn GetDeviceDataCallbackFunc) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceDataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeviceDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *GetDeviceDataCallback[T]) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceDataCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeviceIdCallbackFunc func(this js.Ref, id js.String) js.Ref

func (fn GetDeviceIdCallbackFunc) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceIdCallbackFunc) DispatchCallback(
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

type GetDeviceIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String) js.Ref
	Arg T
}

func (cb *GetDeviceIdCallback[T]) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceIdCallback[T]) DispatchCallback(
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

type GetDeviceInfoCallbackFunc func(this js.Ref, deviceInfo *DeviceInfo) js.Ref

func (fn GetDeviceInfoCallbackFunc) Register() js.Func[func(deviceInfo *DeviceInfo)] {
	return js.RegisterCallback[func(deviceInfo *DeviceInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceInfo
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

type GetDeviceInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deviceInfo *DeviceInfo) js.Ref
	Arg T
}

func (cb *GetDeviceInfoCallback[T]) Register() js.Func[func(deviceInfo *DeviceInfo)] {
	return js.RegisterCallback[func(deviceInfo *DeviceInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceInfo
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

type GetFileSystemInfoOptions struct {
	// Path is "GetFileSystemInfoOptions.path"
	//
	// Optional
	Path js.String
	// ComputeSha256 is "GetFileSystemInfoOptions.computeSha256"
	//
	// Optional
	//
	// NOTE: FFI_USE_ComputeSha256 MUST be set to true to make this field effective.
	ComputeSha256 bool
	// ComputeExecutableMetadata is "GetFileSystemInfoOptions.computeExecutableMetadata"
	//
	// Optional
	//
	// NOTE: FFI_USE_ComputeExecutableMetadata MUST be set to true to make this field effective.
	ComputeExecutableMetadata bool

	FFI_USE_ComputeSha256             bool // for ComputeSha256.
	FFI_USE_ComputeExecutableMetadata bool // for ComputeExecutableMetadata.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFileSystemInfoOptions with all fields set.
func (p GetFileSystemInfoOptions) FromRef(ref js.Ref) GetFileSystemInfoOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFileSystemInfoOptions in the application heap.
func (p GetFileSystemInfoOptions) New() js.Ref {
	return bindings.GetFileSystemInfoOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFileSystemInfoOptions) UpdateFrom(ref js.Ref) {
	bindings.GetFileSystemInfoOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFileSystemInfoOptions) Update(ref js.Ref) {
	bindings.GetFileSystemInfoOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFileSystemInfoOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
}

type UserContext struct {
	// UserId is "UserContext.userId"
	//
	// Optional
	UserId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UserContext with all fields set.
func (p UserContext) FromRef(ref js.Ref) UserContext {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UserContext in the application heap.
func (p UserContext) New() js.Ref {
	return bindings.UserContextJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UserContext) UpdateFrom(ref js.Ref) {
	bindings.UserContextJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UserContext) Update(ref js.Ref) {
	bindings.UserContextJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UserContext) FreeMembers(recursive bool) {
	js.Free(
		p.UserId.Ref(),
	)
	p.UserId = p.UserId.FromRef(js.Undefined)
}

type GetFileSystemInfoRequest struct {
	// UserContext is "GetFileSystemInfoRequest.userContext"
	//
	// Optional
	//
	// NOTE: UserContext.FFI_USE MUST be set to true to get UserContext used.
	UserContext UserContext
	// Options is "GetFileSystemInfoRequest.options"
	//
	// Optional
	Options js.Array[GetFileSystemInfoOptions]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFileSystemInfoRequest with all fields set.
func (p GetFileSystemInfoRequest) FromRef(ref js.Ref) GetFileSystemInfoRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFileSystemInfoRequest in the application heap.
func (p GetFileSystemInfoRequest) New() js.Ref {
	return bindings.GetFileSystemInfoRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFileSystemInfoRequest) UpdateFrom(ref js.Ref) {
	bindings.GetFileSystemInfoRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFileSystemInfoRequest) Update(ref js.Ref) {
	bindings.GetFileSystemInfoRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFileSystemInfoRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Options.Ref(),
	)
	p.Options = p.Options.FromRef(js.Undefined)
	if recursive {
		p.UserContext.FreeMembers(true)
	}
}

type GetPersistentSecretCallbackFunc func(this js.Ref, secret js.ArrayBuffer) js.Ref

func (fn GetPersistentSecretCallbackFunc) Register() js.Func[func(secret js.ArrayBuffer)] {
	return js.RegisterCallback[func(secret js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPersistentSecretCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPersistentSecretCallback[T any] struct {
	Fn  func(arg T, this js.Ref, secret js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *GetPersistentSecretCallback[T]) Register() js.Func[func(secret js.ArrayBuffer)] {
	return js.RegisterCallback[func(secret js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPersistentSecretCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RegistryHive uint32

const (
	_ RegistryHive = iota

	RegistryHive_HKEY_CLASSES_ROOT
	RegistryHive_HKEY_LOCAL_MACHINE
	RegistryHive_HKEY_CURRENT_USER
)

func (RegistryHive) FromRef(str js.Ref) RegistryHive {
	return RegistryHive(bindings.ConstOfRegistryHive(str))
}

func (x RegistryHive) String() (string, bool) {
	switch x {
	case RegistryHive_HKEY_CLASSES_ROOT:
		return "HKEY_CLASSES_ROOT", true
	case RegistryHive_HKEY_LOCAL_MACHINE:
		return "HKEY_LOCAL_MACHINE", true
	case RegistryHive_HKEY_CURRENT_USER:
		return "HKEY_CURRENT_USER", true
	default:
		return "", false
	}
}

type GetSettingsOptions struct {
	// Path is "GetSettingsOptions.path"
	//
	// Optional
	Path js.String
	// Key is "GetSettingsOptions.key"
	//
	// Optional
	Key js.String
	// GetValue is "GetSettingsOptions.getValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_GetValue MUST be set to true to make this field effective.
	GetValue bool
	// Hive is "GetSettingsOptions.hive"
	//
	// Optional
	Hive RegistryHive

	FFI_USE_GetValue bool // for GetValue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetSettingsOptions with all fields set.
func (p GetSettingsOptions) FromRef(ref js.Ref) GetSettingsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetSettingsOptions in the application heap.
func (p GetSettingsOptions) New() js.Ref {
	return bindings.GetSettingsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetSettingsOptions) UpdateFrom(ref js.Ref) {
	bindings.GetSettingsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetSettingsOptions) Update(ref js.Ref) {
	bindings.GetSettingsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetSettingsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
		p.Key.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Key = p.Key.FromRef(js.Undefined)
}

type GetSettingsRequest struct {
	// UserContext is "GetSettingsRequest.userContext"
	//
	// Optional
	//
	// NOTE: UserContext.FFI_USE MUST be set to true to get UserContext used.
	UserContext UserContext
	// Options is "GetSettingsRequest.options"
	//
	// Optional
	Options js.Array[GetSettingsOptions]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetSettingsRequest with all fields set.
func (p GetSettingsRequest) FromRef(ref js.Ref) GetSettingsRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetSettingsRequest in the application heap.
func (p GetSettingsRequest) New() js.Ref {
	return bindings.GetSettingsRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetSettingsRequest) UpdateFrom(ref js.Ref) {
	bindings.GetSettingsRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetSettingsRequest) Update(ref js.Ref) {
	bindings.GetSettingsRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetSettingsRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Options.Ref(),
	)
	p.Options = p.Options.FromRef(js.Undefined)
	if recursive {
		p.UserContext.FreeMembers(true)
	}
}

type GetSettingsResponse struct {
	// Path is "GetSettingsResponse.path"
	//
	// Optional
	Path js.String
	// Key is "GetSettingsResponse.key"
	//
	// Optional
	Key js.String
	// Hive is "GetSettingsResponse.hive"
	//
	// Optional
	Hive RegistryHive
	// Presence is "GetSettingsResponse.presence"
	//
	// Optional
	Presence PresenceValue
	// Value is "GetSettingsResponse.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetSettingsResponse with all fields set.
func (p GetSettingsResponse) FromRef(ref js.Ref) GetSettingsResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetSettingsResponse in the application heap.
func (p GetSettingsResponse) New() js.Ref {
	return bindings.GetSettingsResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetSettingsResponse) UpdateFrom(ref js.Ref) {
	bindings.GetSettingsResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetSettingsResponse) Update(ref js.Ref) {
	bindings.GetSettingsResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetSettingsResponse) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
		p.Key.Ref(),
		p.Value.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Key = p.Key.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type HotfixSignal struct {
	// HotfixId is "HotfixSignal.hotfixId"
	//
	// Optional
	HotfixId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HotfixSignal with all fields set.
func (p HotfixSignal) FromRef(ref js.Ref) HotfixSignal {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HotfixSignal in the application heap.
func (p HotfixSignal) New() js.Ref {
	return bindings.HotfixSignalJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HotfixSignal) UpdateFrom(ref js.Ref) {
	bindings.HotfixSignalJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HotfixSignal) Update(ref js.Ref) {
	bindings.HotfixSignalJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HotfixSignal) FreeMembers(recursive bool) {
	js.Free(
		p.HotfixId.Ref(),
	)
	p.HotfixId = p.HotfixId.FromRef(js.Undefined)
}

type HotfixesCallbackFunc func(this js.Ref, hotfixSignals js.Array[HotfixSignal]) js.Ref

func (fn HotfixesCallbackFunc) Register() js.Func[func(hotfixSignals js.Array[HotfixSignal])] {
	return js.RegisterCallback[func(hotfixSignals js.Array[HotfixSignal])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn HotfixesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[HotfixSignal]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type HotfixesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, hotfixSignals js.Array[HotfixSignal]) js.Ref
	Arg T
}

func (cb *HotfixesCallback[T]) Register() js.Func[func(hotfixSignals js.Array[HotfixSignal])] {
	return js.RegisterCallback[func(hotfixSignals js.Array[HotfixSignal])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *HotfixesCallback[T]) DispatchCallback(
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

		js.Array[HotfixSignal]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SettingsCallbackFunc func(this js.Ref, settings js.Array[GetSettingsResponse]) js.Ref

func (fn SettingsCallbackFunc) Register() js.Func[func(settings js.Array[GetSettingsResponse])] {
	return js.RegisterCallback[func(settings js.Array[GetSettingsResponse])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SettingsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[GetSettingsResponse]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SettingsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, settings js.Array[GetSettingsResponse]) js.Ref
	Arg T
}

func (cb *SettingsCallback[T]) Register() js.Func[func(settings js.Array[GetSettingsResponse])] {
	return js.RegisterCallback[func(settings js.Array[GetSettingsResponse])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SettingsCallback[T]) DispatchCallback(
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

		js.Array[GetSettingsResponse]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncEnqueueRecord returns true if the function "WEBEXT.enterprise.reportingPrivate.enqueueRecord" exists.
func HasFuncEnqueueRecord() bool {
	return js.True == bindings.HasFuncEnqueueRecord()
}

// FuncEnqueueRecord returns the function "WEBEXT.enterprise.reportingPrivate.enqueueRecord".
func FuncEnqueueRecord() (fn js.Func[func(request EnqueueRecordRequest) js.Promise[js.Void]]) {
	bindings.FuncEnqueueRecord(
		js.Pointer(&fn),
	)
	return
}

// EnqueueRecord calls the function "WEBEXT.enterprise.reportingPrivate.enqueueRecord" directly.
func EnqueueRecord(request EnqueueRecordRequest) (ret js.Promise[js.Void]) {
	bindings.CallEnqueueRecord(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryEnqueueRecord calls the function "WEBEXT.enterprise.reportingPrivate.enqueueRecord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnqueueRecord(request EnqueueRecordRequest) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnqueueRecord(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncGetAvInfo returns true if the function "WEBEXT.enterprise.reportingPrivate.getAvInfo" exists.
func HasFuncGetAvInfo() bool {
	return js.True == bindings.HasFuncGetAvInfo()
}

// FuncGetAvInfo returns the function "WEBEXT.enterprise.reportingPrivate.getAvInfo".
func FuncGetAvInfo() (fn js.Func[func(userContext UserContext) js.Promise[js.Array[AntiVirusSignal]]]) {
	bindings.FuncGetAvInfo(
		js.Pointer(&fn),
	)
	return
}

// GetAvInfo calls the function "WEBEXT.enterprise.reportingPrivate.getAvInfo" directly.
func GetAvInfo(userContext UserContext) (ret js.Promise[js.Array[AntiVirusSignal]]) {
	bindings.CallGetAvInfo(
		js.Pointer(&ret),
		js.Pointer(&userContext),
	)

	return
}

// TryGetAvInfo calls the function "WEBEXT.enterprise.reportingPrivate.getAvInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAvInfo(userContext UserContext) (ret js.Promise[js.Array[AntiVirusSignal]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAvInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&userContext),
	)

	return
}

// HasFuncGetCertificate returns true if the function "WEBEXT.enterprise.reportingPrivate.getCertificate" exists.
func HasFuncGetCertificate() bool {
	return js.True == bindings.HasFuncGetCertificate()
}

// FuncGetCertificate returns the function "WEBEXT.enterprise.reportingPrivate.getCertificate".
func FuncGetCertificate() (fn js.Func[func(url js.String) js.Promise[Certificate]]) {
	bindings.FuncGetCertificate(
		js.Pointer(&fn),
	)
	return
}

// GetCertificate calls the function "WEBEXT.enterprise.reportingPrivate.getCertificate" directly.
func GetCertificate(url js.String) (ret js.Promise[Certificate]) {
	bindings.CallGetCertificate(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetCertificate calls the function "WEBEXT.enterprise.reportingPrivate.getCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCertificate(url js.String) (ret js.Promise[Certificate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCertificate(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetContextInfo returns true if the function "WEBEXT.enterprise.reportingPrivate.getContextInfo" exists.
func HasFuncGetContextInfo() bool {
	return js.True == bindings.HasFuncGetContextInfo()
}

// FuncGetContextInfo returns the function "WEBEXT.enterprise.reportingPrivate.getContextInfo".
func FuncGetContextInfo() (fn js.Func[func() js.Promise[ContextInfo]]) {
	bindings.FuncGetContextInfo(
		js.Pointer(&fn),
	)
	return
}

// GetContextInfo calls the function "WEBEXT.enterprise.reportingPrivate.getContextInfo" directly.
func GetContextInfo() (ret js.Promise[ContextInfo]) {
	bindings.CallGetContextInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetContextInfo calls the function "WEBEXT.enterprise.reportingPrivate.getContextInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetContextInfo() (ret js.Promise[ContextInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetContextInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeviceData returns true if the function "WEBEXT.enterprise.reportingPrivate.getDeviceData" exists.
func HasFuncGetDeviceData() bool {
	return js.True == bindings.HasFuncGetDeviceData()
}

// FuncGetDeviceData returns the function "WEBEXT.enterprise.reportingPrivate.getDeviceData".
func FuncGetDeviceData() (fn js.Func[func(id js.String) js.Promise[js.ArrayBuffer]]) {
	bindings.FuncGetDeviceData(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceData calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceData" directly.
func GetDeviceData(id js.String) (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallGetDeviceData(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetDeviceData calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceData(id js.String) (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceData(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetDeviceId returns true if the function "WEBEXT.enterprise.reportingPrivate.getDeviceId" exists.
func HasFuncGetDeviceId() bool {
	return js.True == bindings.HasFuncGetDeviceId()
}

// FuncGetDeviceId returns the function "WEBEXT.enterprise.reportingPrivate.getDeviceId".
func FuncGetDeviceId() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDeviceId(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceId calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceId" directly.
func GetDeviceId() (ret js.Promise[js.String]) {
	bindings.CallGetDeviceId(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceId calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceId"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceId() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceId(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeviceInfo returns true if the function "WEBEXT.enterprise.reportingPrivate.getDeviceInfo" exists.
func HasFuncGetDeviceInfo() bool {
	return js.True == bindings.HasFuncGetDeviceInfo()
}

// FuncGetDeviceInfo returns the function "WEBEXT.enterprise.reportingPrivate.getDeviceInfo".
func FuncGetDeviceInfo() (fn js.Func[func() js.Promise[DeviceInfo]]) {
	bindings.FuncGetDeviceInfo(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceInfo calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceInfo" directly.
func GetDeviceInfo() (ret js.Promise[DeviceInfo]) {
	bindings.CallGetDeviceInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceInfo calls the function "WEBEXT.enterprise.reportingPrivate.getDeviceInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceInfo() (ret js.Promise[DeviceInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetFileSystemInfo returns true if the function "WEBEXT.enterprise.reportingPrivate.getFileSystemInfo" exists.
func HasFuncGetFileSystemInfo() bool {
	return js.True == bindings.HasFuncGetFileSystemInfo()
}

// FuncGetFileSystemInfo returns the function "WEBEXT.enterprise.reportingPrivate.getFileSystemInfo".
func FuncGetFileSystemInfo() (fn js.Func[func(request GetFileSystemInfoRequest) js.Promise[js.Array[GetFileSystemInfoResponse]]]) {
	bindings.FuncGetFileSystemInfo(
		js.Pointer(&fn),
	)
	return
}

// GetFileSystemInfo calls the function "WEBEXT.enterprise.reportingPrivate.getFileSystemInfo" directly.
func GetFileSystemInfo(request GetFileSystemInfoRequest) (ret js.Promise[js.Array[GetFileSystemInfoResponse]]) {
	bindings.CallGetFileSystemInfo(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryGetFileSystemInfo calls the function "WEBEXT.enterprise.reportingPrivate.getFileSystemInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFileSystemInfo(request GetFileSystemInfoRequest) (ret js.Promise[js.Array[GetFileSystemInfoResponse]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFileSystemInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncGetHotfixes returns true if the function "WEBEXT.enterprise.reportingPrivate.getHotfixes" exists.
func HasFuncGetHotfixes() bool {
	return js.True == bindings.HasFuncGetHotfixes()
}

// FuncGetHotfixes returns the function "WEBEXT.enterprise.reportingPrivate.getHotfixes".
func FuncGetHotfixes() (fn js.Func[func(userContext UserContext) js.Promise[js.Array[HotfixSignal]]]) {
	bindings.FuncGetHotfixes(
		js.Pointer(&fn),
	)
	return
}

// GetHotfixes calls the function "WEBEXT.enterprise.reportingPrivate.getHotfixes" directly.
func GetHotfixes(userContext UserContext) (ret js.Promise[js.Array[HotfixSignal]]) {
	bindings.CallGetHotfixes(
		js.Pointer(&ret),
		js.Pointer(&userContext),
	)

	return
}

// TryGetHotfixes calls the function "WEBEXT.enterprise.reportingPrivate.getHotfixes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetHotfixes(userContext UserContext) (ret js.Promise[js.Array[HotfixSignal]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetHotfixes(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&userContext),
	)

	return
}

// HasFuncGetPersistentSecret returns true if the function "WEBEXT.enterprise.reportingPrivate.getPersistentSecret" exists.
func HasFuncGetPersistentSecret() bool {
	return js.True == bindings.HasFuncGetPersistentSecret()
}

// FuncGetPersistentSecret returns the function "WEBEXT.enterprise.reportingPrivate.getPersistentSecret".
func FuncGetPersistentSecret() (fn js.Func[func(resetSecret bool) js.Promise[js.ArrayBuffer]]) {
	bindings.FuncGetPersistentSecret(
		js.Pointer(&fn),
	)
	return
}

// GetPersistentSecret calls the function "WEBEXT.enterprise.reportingPrivate.getPersistentSecret" directly.
func GetPersistentSecret(resetSecret bool) (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallGetPersistentSecret(
		js.Pointer(&ret),
		js.Bool(bool(resetSecret)),
	)

	return
}

// TryGetPersistentSecret calls the function "WEBEXT.enterprise.reportingPrivate.getPersistentSecret"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPersistentSecret(resetSecret bool) (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPersistentSecret(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(resetSecret)),
	)

	return
}

// HasFuncGetSettings returns true if the function "WEBEXT.enterprise.reportingPrivate.getSettings" exists.
func HasFuncGetSettings() bool {
	return js.True == bindings.HasFuncGetSettings()
}

// FuncGetSettings returns the function "WEBEXT.enterprise.reportingPrivate.getSettings".
func FuncGetSettings() (fn js.Func[func(request GetSettingsRequest) js.Promise[js.Array[GetSettingsResponse]]]) {
	bindings.FuncGetSettings(
		js.Pointer(&fn),
	)
	return
}

// GetSettings calls the function "WEBEXT.enterprise.reportingPrivate.getSettings" directly.
func GetSettings(request GetSettingsRequest) (ret js.Promise[js.Array[GetSettingsResponse]]) {
	bindings.CallGetSettings(
		js.Pointer(&ret),
		js.Pointer(&request),
	)

	return
}

// TryGetSettings calls the function "WEBEXT.enterprise.reportingPrivate.getSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSettings(request GetSettingsRequest) (ret js.Promise[js.Array[GetSettingsResponse]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSettings(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&request),
	)

	return
}

// HasFuncSetDeviceData returns true if the function "WEBEXT.enterprise.reportingPrivate.setDeviceData" exists.
func HasFuncSetDeviceData() bool {
	return js.True == bindings.HasFuncSetDeviceData()
}

// FuncSetDeviceData returns the function "WEBEXT.enterprise.reportingPrivate.setDeviceData".
func FuncSetDeviceData() (fn js.Func[func(id js.String, data js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncSetDeviceData(
		js.Pointer(&fn),
	)
	return
}

// SetDeviceData calls the function "WEBEXT.enterprise.reportingPrivate.setDeviceData" directly.
func SetDeviceData(id js.String, data js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallSetDeviceData(
		js.Pointer(&ret),
		id.Ref(),
		data.Ref(),
	)

	return
}

// TrySetDeviceData calls the function "WEBEXT.enterprise.reportingPrivate.setDeviceData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDeviceData(id js.String, data js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDeviceData(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		data.Ref(),
	)

	return
}

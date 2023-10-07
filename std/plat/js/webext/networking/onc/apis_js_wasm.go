// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package onc

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/networking/onc/bindings"
)

type ActivationStateType uint32

const (
	_ ActivationStateType = iota

	ActivationStateType_ACTIVATED
	ActivationStateType_ACTIVATING
	ActivationStateType_NOT_ACTIVATED
	ActivationStateType_PARTIALLY_ACTIVATED
)

func (ActivationStateType) FromRef(str js.Ref) ActivationStateType {
	return ActivationStateType(bindings.ConstOfActivationStateType(str))
}

func (x ActivationStateType) String() (string, bool) {
	switch x {
	case ActivationStateType_ACTIVATED:
		return "Activated", true
	case ActivationStateType_ACTIVATING:
		return "Activating", true
	case ActivationStateType_NOT_ACTIVATED:
		return "NotActivated", true
	case ActivationStateType_PARTIALLY_ACTIVATED:
		return "PartiallyActivated", true
	default:
		return "", false
	}
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

type CaptivePortalStatus uint32

const (
	_ CaptivePortalStatus = iota

	CaptivePortalStatus_UNKNOWN
	CaptivePortalStatus_OFFLINE
	CaptivePortalStatus_ONLINE
	CaptivePortalStatus_PORTAL
	CaptivePortalStatus_PROXY_AUTH_REQUIRED
)

func (CaptivePortalStatus) FromRef(str js.Ref) CaptivePortalStatus {
	return CaptivePortalStatus(bindings.ConstOfCaptivePortalStatus(str))
}

func (x CaptivePortalStatus) String() (string, bool) {
	switch x {
	case CaptivePortalStatus_UNKNOWN:
		return "Unknown", true
	case CaptivePortalStatus_OFFLINE:
		return "Offline", true
	case CaptivePortalStatus_ONLINE:
		return "Online", true
	case CaptivePortalStatus_PORTAL:
		return "Portal", true
	case CaptivePortalStatus_PROXY_AUTH_REQUIRED:
		return "ProxyAuthRequired", true
	default:
		return "", false
	}
}

type CaptivePortalStatusCallbackFunc func(this js.Ref, result CaptivePortalStatus) js.Ref

func (fn CaptivePortalStatusCallbackFunc) Register() js.Func[func(result CaptivePortalStatus)] {
	return js.RegisterCallback[func(result CaptivePortalStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CaptivePortalStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		CaptivePortalStatus(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CaptivePortalStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result CaptivePortalStatus) js.Ref
	Arg T
}

func (cb *CaptivePortalStatusCallback[T]) Register() js.Func[func(result CaptivePortalStatus)] {
	return js.RegisterCallback[func(result CaptivePortalStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CaptivePortalStatusCallback[T]) DispatchCallback(
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

		CaptivePortalStatus(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FoundNetworkProperties struct {
	// Status is "FoundNetworkProperties.Status"
	//
	// Optional
	Status js.String
	// NetworkId is "FoundNetworkProperties.NetworkId"
	//
	// Optional
	NetworkId js.String
	// Technology is "FoundNetworkProperties.Technology"
	//
	// Optional
	Technology js.String
	// ShortName is "FoundNetworkProperties.ShortName"
	//
	// Optional
	ShortName js.String
	// LongName is "FoundNetworkProperties.LongName"
	//
	// Optional
	LongName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FoundNetworkProperties with all fields set.
func (p FoundNetworkProperties) FromRef(ref js.Ref) FoundNetworkProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FoundNetworkProperties in the application heap.
func (p FoundNetworkProperties) New() js.Ref {
	return bindings.FoundNetworkPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FoundNetworkProperties) UpdateFrom(ref js.Ref) {
	bindings.FoundNetworkPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FoundNetworkProperties) Update(ref js.Ref) {
	bindings.FoundNetworkPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FoundNetworkProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Status.Ref(),
		p.NetworkId.Ref(),
		p.Technology.Ref(),
		p.ShortName.Ref(),
		p.LongName.Ref(),
	)
	p.Status = p.Status.FromRef(js.Undefined)
	p.NetworkId = p.NetworkId.FromRef(js.Undefined)
	p.Technology = p.Technology.FromRef(js.Undefined)
	p.ShortName = p.ShortName.FromRef(js.Undefined)
	p.LongName = p.LongName.FromRef(js.Undefined)
}

type CellularProviderProperties struct {
	// Name is "CellularProviderProperties.Name"
	//
	// Optional
	Name js.String
	// Code is "CellularProviderProperties.Code"
	//
	// Optional
	Code js.String
	// Country is "CellularProviderProperties.Country"
	//
	// Optional
	Country js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CellularProviderProperties with all fields set.
func (p CellularProviderProperties) FromRef(ref js.Ref) CellularProviderProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CellularProviderProperties in the application heap.
func (p CellularProviderProperties) New() js.Ref {
	return bindings.CellularProviderPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CellularProviderProperties) UpdateFrom(ref js.Ref) {
	bindings.CellularProviderPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CellularProviderProperties) Update(ref js.Ref) {
	bindings.CellularProviderPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CellularProviderProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Code.Ref(),
		p.Country.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Code = p.Code.FromRef(js.Undefined)
	p.Country = p.Country.FromRef(js.Undefined)
}

type PaymentPortal struct {
	// Method is "PaymentPortal.Method"
	//
	// Optional
	Method js.String
	// PostData is "PaymentPortal.PostData"
	//
	// Optional
	PostData js.String
	// Url is "PaymentPortal.Url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentPortal with all fields set.
func (p PaymentPortal) FromRef(ref js.Ref) PaymentPortal {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentPortal in the application heap.
func (p PaymentPortal) New() js.Ref {
	return bindings.PaymentPortalJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PaymentPortal) UpdateFrom(ref js.Ref) {
	bindings.PaymentPortalJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PaymentPortal) Update(ref js.Ref) {
	bindings.PaymentPortalJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PaymentPortal) FreeMembers(recursive bool) {
	js.Free(
		p.Method.Ref(),
		p.PostData.Ref(),
		p.Url.Ref(),
	)
	p.Method = p.Method.FromRef(js.Undefined)
	p.PostData = p.PostData.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type SIMLockStatus struct {
	// LockType is "SIMLockStatus.LockType"
	//
	// Optional
	LockType js.String
	// LockEnabled is "SIMLockStatus.LockEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_LockEnabled MUST be set to true to make this field effective.
	LockEnabled bool
	// RetriesLeft is "SIMLockStatus.RetriesLeft"
	//
	// Optional
	//
	// NOTE: FFI_USE_RetriesLeft MUST be set to true to make this field effective.
	RetriesLeft int32

	FFI_USE_LockEnabled bool // for LockEnabled.
	FFI_USE_RetriesLeft bool // for RetriesLeft.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SIMLockStatus with all fields set.
func (p SIMLockStatus) FromRef(ref js.Ref) SIMLockStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SIMLockStatus in the application heap.
func (p SIMLockStatus) New() js.Ref {
	return bindings.SIMLockStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SIMLockStatus) UpdateFrom(ref js.Ref) {
	bindings.SIMLockStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SIMLockStatus) Update(ref js.Ref) {
	bindings.SIMLockStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SIMLockStatus) FreeMembers(recursive bool) {
	js.Free(
		p.LockType.Ref(),
	)
	p.LockType = p.LockType.FromRef(js.Undefined)
}

type CellularProperties struct {
	// AutoConnect is "CellularProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoConnect MUST be set to true to make this field effective.
	AutoConnect bool
	// ActivationType is "CellularProperties.ActivationType"
	//
	// Optional
	ActivationType js.String
	// ActivationState is "CellularProperties.ActivationState"
	//
	// Optional
	ActivationState ActivationStateType
	// AllowRoaming is "CellularProperties.AllowRoaming"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowRoaming MUST be set to true to make this field effective.
	AllowRoaming bool
	// Family is "CellularProperties.Family"
	//
	// Optional
	Family js.String
	// FirmwareRevision is "CellularProperties.FirmwareRevision"
	//
	// Optional
	FirmwareRevision js.String
	// FoundNetworks is "CellularProperties.FoundNetworks"
	//
	// Optional
	FoundNetworks js.Array[FoundNetworkProperties]
	// HardwareRevision is "CellularProperties.HardwareRevision"
	//
	// Optional
	HardwareRevision js.String
	// HomeProvider is "CellularProperties.HomeProvider"
	//
	// Optional
	//
	// NOTE: HomeProvider.FFI_USE MUST be set to true to get HomeProvider used.
	HomeProvider CellularProviderProperties
	// Manufacturer is "CellularProperties.Manufacturer"
	//
	// Optional
	Manufacturer js.String
	// ModelID is "CellularProperties.ModelID"
	//
	// Optional
	ModelID js.String
	// NetworkTechnology is "CellularProperties.NetworkTechnology"
	//
	// Optional
	NetworkTechnology js.String
	// PaymentPortal is "CellularProperties.PaymentPortal"
	//
	// Optional
	//
	// NOTE: PaymentPortal.FFI_USE MUST be set to true to get PaymentPortal used.
	PaymentPortal PaymentPortal
	// RoamingState is "CellularProperties.RoamingState"
	//
	// Optional
	RoamingState js.String
	// Scanning is "CellularProperties.Scanning"
	//
	// Optional
	//
	// NOTE: FFI_USE_Scanning MUST be set to true to make this field effective.
	Scanning bool
	// ServingOperator is "CellularProperties.ServingOperator"
	//
	// Optional
	//
	// NOTE: ServingOperator.FFI_USE MUST be set to true to get ServingOperator used.
	ServingOperator CellularProviderProperties
	// SIMLockStatus is "CellularProperties.SIMLockStatus"
	//
	// Optional
	//
	// NOTE: SIMLockStatus.FFI_USE MUST be set to true to get SIMLockStatus used.
	SIMLockStatus SIMLockStatus
	// SIMPresent is "CellularProperties.SIMPresent"
	//
	// Optional
	//
	// NOTE: FFI_USE_SIMPresent MUST be set to true to make this field effective.
	SIMPresent bool
	// SignalStrength is "CellularProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32
	// SupportNetworkScan is "CellularProperties.SupportNetworkScan"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportNetworkScan MUST be set to true to make this field effective.
	SupportNetworkScan bool

	FFI_USE_AutoConnect        bool // for AutoConnect.
	FFI_USE_AllowRoaming       bool // for AllowRoaming.
	FFI_USE_Scanning           bool // for Scanning.
	FFI_USE_SIMPresent         bool // for SIMPresent.
	FFI_USE_SignalStrength     bool // for SignalStrength.
	FFI_USE_SupportNetworkScan bool // for SupportNetworkScan.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CellularProperties with all fields set.
func (p CellularProperties) FromRef(ref js.Ref) CellularProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CellularProperties in the application heap.
func (p CellularProperties) New() js.Ref {
	return bindings.CellularPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CellularProperties) UpdateFrom(ref js.Ref) {
	bindings.CellularPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CellularProperties) Update(ref js.Ref) {
	bindings.CellularPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CellularProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ActivationType.Ref(),
		p.Family.Ref(),
		p.FirmwareRevision.Ref(),
		p.FoundNetworks.Ref(),
		p.HardwareRevision.Ref(),
		p.Manufacturer.Ref(),
		p.ModelID.Ref(),
		p.NetworkTechnology.Ref(),
		p.RoamingState.Ref(),
	)
	p.ActivationType = p.ActivationType.FromRef(js.Undefined)
	p.Family = p.Family.FromRef(js.Undefined)
	p.FirmwareRevision = p.FirmwareRevision.FromRef(js.Undefined)
	p.FoundNetworks = p.FoundNetworks.FromRef(js.Undefined)
	p.HardwareRevision = p.HardwareRevision.FromRef(js.Undefined)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
	p.ModelID = p.ModelID.FromRef(js.Undefined)
	p.NetworkTechnology = p.NetworkTechnology.FromRef(js.Undefined)
	p.RoamingState = p.RoamingState.FromRef(js.Undefined)
	if recursive {
		p.HomeProvider.FreeMembers(true)
		p.PaymentPortal.FreeMembers(true)
		p.ServingOperator.FreeMembers(true)
		p.SIMLockStatus.FreeMembers(true)
	}
}

type CellularStateProperties struct {
	// ActivationState is "CellularStateProperties.ActivationState"
	//
	// Optional
	ActivationState ActivationStateType
	// NetworkTechnology is "CellularStateProperties.NetworkTechnology"
	//
	// Optional
	NetworkTechnology js.String
	// RoamingState is "CellularStateProperties.RoamingState"
	//
	// Optional
	RoamingState js.String
	// SIMPresent is "CellularStateProperties.SIMPresent"
	//
	// Optional
	//
	// NOTE: FFI_USE_SIMPresent MUST be set to true to make this field effective.
	SIMPresent bool
	// SignalStrength is "CellularStateProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32

	FFI_USE_SIMPresent     bool // for SIMPresent.
	FFI_USE_SignalStrength bool // for SignalStrength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CellularStateProperties with all fields set.
func (p CellularStateProperties) FromRef(ref js.Ref) CellularStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CellularStateProperties in the application heap.
func (p CellularStateProperties) New() js.Ref {
	return bindings.CellularStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CellularStateProperties) UpdateFrom(ref js.Ref) {
	bindings.CellularStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CellularStateProperties) Update(ref js.Ref) {
	bindings.CellularStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CellularStateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.NetworkTechnology.Ref(),
		p.RoamingState.Ref(),
	)
	p.NetworkTechnology = p.NetworkTechnology.FromRef(js.Undefined)
	p.RoamingState = p.RoamingState.FromRef(js.Undefined)
}

type IssuerSubjectPattern struct {
	// CommonName is "IssuerSubjectPattern.CommonName"
	//
	// Optional
	CommonName js.String
	// Locality is "IssuerSubjectPattern.Locality"
	//
	// Optional
	Locality js.String
	// Organization is "IssuerSubjectPattern.Organization"
	//
	// Optional
	Organization js.String
	// OrganizationalUnit is "IssuerSubjectPattern.OrganizationalUnit"
	//
	// Optional
	OrganizationalUnit js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IssuerSubjectPattern with all fields set.
func (p IssuerSubjectPattern) FromRef(ref js.Ref) IssuerSubjectPattern {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IssuerSubjectPattern in the application heap.
func (p IssuerSubjectPattern) New() js.Ref {
	return bindings.IssuerSubjectPatternJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IssuerSubjectPattern) UpdateFrom(ref js.Ref) {
	bindings.IssuerSubjectPatternJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IssuerSubjectPattern) Update(ref js.Ref) {
	bindings.IssuerSubjectPatternJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IssuerSubjectPattern) FreeMembers(recursive bool) {
	js.Free(
		p.CommonName.Ref(),
		p.Locality.Ref(),
		p.Organization.Ref(),
		p.OrganizationalUnit.Ref(),
	)
	p.CommonName = p.CommonName.FromRef(js.Undefined)
	p.Locality = p.Locality.FromRef(js.Undefined)
	p.Organization = p.Organization.FromRef(js.Undefined)
	p.OrganizationalUnit = p.OrganizationalUnit.FromRef(js.Undefined)
}

type CertificatePattern struct {
	// EnrollmentURI is "CertificatePattern.EnrollmentURI"
	//
	// Optional
	EnrollmentURI js.Array[js.String]
	// Issuer is "CertificatePattern.Issuer"
	//
	// Optional
	//
	// NOTE: Issuer.FFI_USE MUST be set to true to get Issuer used.
	Issuer IssuerSubjectPattern
	// IssuerCARef is "CertificatePattern.IssuerCARef"
	//
	// Optional
	IssuerCARef js.Array[js.String]
	// Subject is "CertificatePattern.Subject"
	//
	// Optional
	//
	// NOTE: Subject.FFI_USE MUST be set to true to get Subject used.
	Subject IssuerSubjectPattern

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CertificatePattern with all fields set.
func (p CertificatePattern) FromRef(ref js.Ref) CertificatePattern {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CertificatePattern in the application heap.
func (p CertificatePattern) New() js.Ref {
	return bindings.CertificatePatternJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CertificatePattern) UpdateFrom(ref js.Ref) {
	bindings.CertificatePatternJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CertificatePattern) Update(ref js.Ref) {
	bindings.CertificatePatternJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CertificatePattern) FreeMembers(recursive bool) {
	js.Free(
		p.EnrollmentURI.Ref(),
		p.IssuerCARef.Ref(),
	)
	p.EnrollmentURI = p.EnrollmentURI.FromRef(js.Undefined)
	p.IssuerCARef = p.IssuerCARef.FromRef(js.Undefined)
	if recursive {
		p.Issuer.FreeMembers(true)
		p.Subject.FreeMembers(true)
	}
}

type ClientCertificateType uint32

const (
	_ ClientCertificateType = iota

	ClientCertificateType_REF
	ClientCertificateType_PATTERN
)

func (ClientCertificateType) FromRef(str js.Ref) ClientCertificateType {
	return ClientCertificateType(bindings.ConstOfClientCertificateType(str))
}

func (x ClientCertificateType) String() (string, bool) {
	switch x {
	case ClientCertificateType_REF:
		return "Ref", true
	case ClientCertificateType_PATTERN:
		return "Pattern", true
	default:
		return "", false
	}
}

type ConnectionStateType uint32

const (
	_ ConnectionStateType = iota

	ConnectionStateType_CONNECTED
	ConnectionStateType_CONNECTING
	ConnectionStateType_NOT_CONNECTED
)

func (ConnectionStateType) FromRef(str js.Ref) ConnectionStateType {
	return ConnectionStateType(bindings.ConstOfConnectionStateType(str))
}

func (x ConnectionStateType) String() (string, bool) {
	switch x {
	case ConnectionStateType_CONNECTED:
		return "Connected", true
	case ConnectionStateType_CONNECTING:
		return "Connecting", true
	case ConnectionStateType_NOT_CONNECTED:
		return "NotConnected", true
	default:
		return "", false
	}
}

type DeviceStateType uint32

const (
	_ DeviceStateType = iota

	DeviceStateType_UNINITIALIZED
	DeviceStateType_DISABLED
	DeviceStateType_ENABLING
	DeviceStateType_ENABLED
	DeviceStateType_PROHIBITED
)

func (DeviceStateType) FromRef(str js.Ref) DeviceStateType {
	return DeviceStateType(bindings.ConstOfDeviceStateType(str))
}

func (x DeviceStateType) String() (string, bool) {
	switch x {
	case DeviceStateType_UNINITIALIZED:
		return "Uninitialized", true
	case DeviceStateType_DISABLED:
		return "Disabled", true
	case DeviceStateType_ENABLING:
		return "Enabling", true
	case DeviceStateType_ENABLED:
		return "Enabled", true
	case DeviceStateType_PROHIBITED:
		return "Prohibited", true
	default:
		return "", false
	}
}

type NetworkType uint32

const (
	_ NetworkType = iota

	NetworkType_ALL
	NetworkType_CELLULAR
	NetworkType_ETHERNET
	NetworkType_TETHER
	NetworkType_VPN
	NetworkType_WIRELESS
	NetworkType_WI_FI
)

func (NetworkType) FromRef(str js.Ref) NetworkType {
	return NetworkType(bindings.ConstOfNetworkType(str))
}

func (x NetworkType) String() (string, bool) {
	switch x {
	case NetworkType_ALL:
		return "All", true
	case NetworkType_CELLULAR:
		return "Cellular", true
	case NetworkType_ETHERNET:
		return "Ethernet", true
	case NetworkType_TETHER:
		return "Tether", true
	case NetworkType_VPN:
		return "VPN", true
	case NetworkType_WIRELESS:
		return "Wireless", true
	case NetworkType_WI_FI:
		return "WiFi", true
	default:
		return "", false
	}
}

type DeviceStateProperties struct {
	// Scanning is "DeviceStateProperties.Scanning"
	//
	// Optional
	//
	// NOTE: FFI_USE_Scanning MUST be set to true to make this field effective.
	Scanning bool
	// SIMLockStatus is "DeviceStateProperties.SIMLockStatus"
	//
	// Optional
	//
	// NOTE: SIMLockStatus.FFI_USE MUST be set to true to get SIMLockStatus used.
	SIMLockStatus SIMLockStatus
	// SIMPresent is "DeviceStateProperties.SIMPresent"
	//
	// Optional
	//
	// NOTE: FFI_USE_SIMPresent MUST be set to true to make this field effective.
	SIMPresent bool
	// State is "DeviceStateProperties.State"
	//
	// Optional
	State DeviceStateType
	// Type is "DeviceStateProperties.Type"
	//
	// Optional
	Type NetworkType

	FFI_USE_Scanning   bool // for Scanning.
	FFI_USE_SIMPresent bool // for SIMPresent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceStateProperties with all fields set.
func (p DeviceStateProperties) FromRef(ref js.Ref) DeviceStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceStateProperties in the application heap.
func (p DeviceStateProperties) New() js.Ref {
	return bindings.DeviceStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceStateProperties) UpdateFrom(ref js.Ref) {
	bindings.DeviceStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceStateProperties) Update(ref js.Ref) {
	bindings.DeviceStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceStateProperties) FreeMembers(recursive bool) {
	if recursive {
		p.SIMLockStatus.FreeMembers(true)
	}
}

type ManagedDOMString struct {
	// Active is "ManagedDOMString.Active"
	//
	// Optional
	Active js.String
	// Effective is "ManagedDOMString.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedDOMString.UserPolicy"
	//
	// Optional
	UserPolicy js.String
	// DevicePolicy is "ManagedDOMString.DevicePolicy"
	//
	// Optional
	DevicePolicy js.String
	// UserSetting is "ManagedDOMString.UserSetting"
	//
	// Optional
	UserSetting js.String
	// SharedSetting is "ManagedDOMString.SharedSetting"
	//
	// Optional
	SharedSetting js.String
	// UserEditable is "ManagedDOMString.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedDOMString.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedDOMString with all fields set.
func (p ManagedDOMString) FromRef(ref js.Ref) ManagedDOMString {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedDOMString in the application heap.
func (p ManagedDOMString) New() js.Ref {
	return bindings.ManagedDOMStringJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedDOMString) UpdateFrom(ref js.Ref) {
	bindings.ManagedDOMStringJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedDOMString) Update(ref js.Ref) {
	bindings.ManagedDOMStringJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedDOMString) FreeMembers(recursive bool) {
	js.Free(
		p.Active.Ref(),
		p.Effective.Ref(),
		p.UserPolicy.Ref(),
		p.DevicePolicy.Ref(),
		p.UserSetting.Ref(),
		p.SharedSetting.Ref(),
	)
	p.Active = p.Active.FromRef(js.Undefined)
	p.Effective = p.Effective.FromRef(js.Undefined)
	p.UserPolicy = p.UserPolicy.FromRef(js.Undefined)
	p.DevicePolicy = p.DevicePolicy.FromRef(js.Undefined)
	p.UserSetting = p.UserSetting.FromRef(js.Undefined)
	p.SharedSetting = p.SharedSetting.FromRef(js.Undefined)
}

type EAPProperties struct {
	// AnonymousIdentity is "EAPProperties.AnonymousIdentity"
	//
	// Optional
	AnonymousIdentity js.String
	// ClientCertPattern is "EAPProperties.ClientCertPattern"
	//
	// Optional
	//
	// NOTE: ClientCertPattern.FFI_USE MUST be set to true to get ClientCertPattern used.
	ClientCertPattern CertificatePattern
	// ClientCertPKCS11Id is "EAPProperties.ClientCertPKCS11Id"
	//
	// Optional
	ClientCertPKCS11Id js.String
	// ClientCertProvisioningProfileId is "EAPProperties.ClientCertProvisioningProfileId"
	//
	// Optional
	ClientCertProvisioningProfileId js.String
	// ClientCertRef is "EAPProperties.ClientCertRef"
	//
	// Optional
	ClientCertRef js.String
	// ClientCertType is "EAPProperties.ClientCertType"
	//
	// Optional
	ClientCertType ClientCertificateType
	// Identity is "EAPProperties.Identity"
	//
	// Optional
	Identity js.String
	// Inner is "EAPProperties.Inner"
	//
	// Optional
	Inner js.String
	// Outer is "EAPProperties.Outer"
	//
	// Optional
	Outer js.String
	// Password is "EAPProperties.Password"
	//
	// Optional
	Password js.String
	// SaveCredentials is "EAPProperties.SaveCredentials"
	//
	// Optional
	//
	// NOTE: FFI_USE_SaveCredentials MUST be set to true to make this field effective.
	SaveCredentials bool
	// ServerCAPEMs is "EAPProperties.ServerCAPEMs"
	//
	// Optional
	ServerCAPEMs js.Array[js.String]
	// ServerCARefs is "EAPProperties.ServerCARefs"
	//
	// Optional
	ServerCARefs js.Array[js.String]
	// SubjectMatch is "EAPProperties.SubjectMatch"
	//
	// Optional
	//
	// NOTE: SubjectMatch.FFI_USE MUST be set to true to get SubjectMatch used.
	SubjectMatch ManagedDOMString
	// UseProactiveKeyCaching is "EAPProperties.UseProactiveKeyCaching"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseProactiveKeyCaching MUST be set to true to make this field effective.
	UseProactiveKeyCaching bool
	// UseSystemCAs is "EAPProperties.UseSystemCAs"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseSystemCAs MUST be set to true to make this field effective.
	UseSystemCAs bool

	FFI_USE_SaveCredentials        bool // for SaveCredentials.
	FFI_USE_UseProactiveKeyCaching bool // for UseProactiveKeyCaching.
	FFI_USE_UseSystemCAs           bool // for UseSystemCAs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EAPProperties with all fields set.
func (p EAPProperties) FromRef(ref js.Ref) EAPProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EAPProperties in the application heap.
func (p EAPProperties) New() js.Ref {
	return bindings.EAPPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EAPProperties) UpdateFrom(ref js.Ref) {
	bindings.EAPPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EAPProperties) Update(ref js.Ref) {
	bindings.EAPPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EAPProperties) FreeMembers(recursive bool) {
	js.Free(
		p.AnonymousIdentity.Ref(),
		p.ClientCertPKCS11Id.Ref(),
		p.ClientCertProvisioningProfileId.Ref(),
		p.ClientCertRef.Ref(),
		p.Identity.Ref(),
		p.Inner.Ref(),
		p.Outer.Ref(),
		p.Password.Ref(),
		p.ServerCAPEMs.Ref(),
		p.ServerCARefs.Ref(),
	)
	p.AnonymousIdentity = p.AnonymousIdentity.FromRef(js.Undefined)
	p.ClientCertPKCS11Id = p.ClientCertPKCS11Id.FromRef(js.Undefined)
	p.ClientCertProvisioningProfileId = p.ClientCertProvisioningProfileId.FromRef(js.Undefined)
	p.ClientCertRef = p.ClientCertRef.FromRef(js.Undefined)
	p.Identity = p.Identity.FromRef(js.Undefined)
	p.Inner = p.Inner.FromRef(js.Undefined)
	p.Outer = p.Outer.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.ServerCAPEMs = p.ServerCAPEMs.FromRef(js.Undefined)
	p.ServerCARefs = p.ServerCARefs.FromRef(js.Undefined)
	if recursive {
		p.ClientCertPattern.FreeMembers(true)
		p.SubjectMatch.FreeMembers(true)
	}
}

type EthernetProperties struct {
	// AutoConnect is "EthernetProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoConnect MUST be set to true to make this field effective.
	AutoConnect bool
	// Authentication is "EthernetProperties.Authentication"
	//
	// Optional
	Authentication js.String
	// EAP is "EthernetProperties.EAP"
	//
	// Optional
	//
	// NOTE: EAP.FFI_USE MUST be set to true to get EAP used.
	EAP EAPProperties

	FFI_USE_AutoConnect bool // for AutoConnect.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EthernetProperties with all fields set.
func (p EthernetProperties) FromRef(ref js.Ref) EthernetProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EthernetProperties in the application heap.
func (p EthernetProperties) New() js.Ref {
	return bindings.EthernetPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EthernetProperties) UpdateFrom(ref js.Ref) {
	bindings.EthernetPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EthernetProperties) Update(ref js.Ref) {
	bindings.EthernetPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EthernetProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Authentication.Ref(),
	)
	p.Authentication = p.Authentication.FromRef(js.Undefined)
	if recursive {
		p.EAP.FreeMembers(true)
	}
}

type EthernetStateProperties struct {
	// Authentication is "EthernetStateProperties.Authentication"
	//
	// Optional
	Authentication js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EthernetStateProperties with all fields set.
func (p EthernetStateProperties) FromRef(ref js.Ref) EthernetStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EthernetStateProperties in the application heap.
func (p EthernetStateProperties) New() js.Ref {
	return bindings.EthernetStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EthernetStateProperties) UpdateFrom(ref js.Ref) {
	bindings.EthernetStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EthernetStateProperties) Update(ref js.Ref) {
	bindings.EthernetStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EthernetStateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Authentication.Ref(),
	)
	p.Authentication = p.Authentication.FromRef(js.Undefined)
}

type GetDeviceStatesCallbackFunc func(this js.Ref, result js.Array[DeviceStateProperties]) js.Ref

func (fn GetDeviceStatesCallbackFunc) Register() js.Func[func(result js.Array[DeviceStateProperties])] {
	return js.RegisterCallback[func(result js.Array[DeviceStateProperties])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceStatesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DeviceStateProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDeviceStatesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[DeviceStateProperties]) js.Ref
	Arg T
}

func (cb *GetDeviceStatesCallback[T]) Register() js.Func[func(result js.Array[DeviceStateProperties])] {
	return js.RegisterCallback[func(result js.Array[DeviceStateProperties])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceStatesCallback[T]) DispatchCallback(
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

		js.Array[DeviceStateProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetEnabledNetworkTypesCallbackFunc func(this js.Ref, result js.Array[NetworkType]) js.Ref

func (fn GetEnabledNetworkTypesCallbackFunc) Register() js.Func[func(result js.Array[NetworkType])] {
	return js.RegisterCallback[func(result js.Array[NetworkType])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetEnabledNetworkTypesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[NetworkType]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetEnabledNetworkTypesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[NetworkType]) js.Ref
	Arg T
}

func (cb *GetEnabledNetworkTypesCallback[T]) Register() js.Func[func(result js.Array[NetworkType])] {
	return js.RegisterCallback[func(result js.Array[NetworkType])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetEnabledNetworkTypesCallback[T]) DispatchCallback(
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

		js.Array[NetworkType]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetGlobalPolicyCallbackFunc func(this js.Ref, result *GlobalPolicy) js.Ref

func (fn GetGlobalPolicyCallbackFunc) Register() js.Func[func(result *GlobalPolicy)] {
	return js.RegisterCallback[func(result *GlobalPolicy)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetGlobalPolicyCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GlobalPolicy
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

type GetGlobalPolicyCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *GlobalPolicy) js.Ref
	Arg T
}

func (cb *GetGlobalPolicyCallback[T]) Register() js.Func[func(result *GlobalPolicy)] {
	return js.RegisterCallback[func(result *GlobalPolicy)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetGlobalPolicyCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GlobalPolicy
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

type GlobalPolicy struct {
	// AllowOnlyPolicyNetworksToAutoconnect is "GlobalPolicy.AllowOnlyPolicyNetworksToAutoconnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowOnlyPolicyNetworksToAutoconnect MUST be set to true to make this field effective.
	AllowOnlyPolicyNetworksToAutoconnect bool
	// AllowOnlyPolicyNetworksToConnect is "GlobalPolicy.AllowOnlyPolicyNetworksToConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowOnlyPolicyNetworksToConnect MUST be set to true to make this field effective.
	AllowOnlyPolicyNetworksToConnect bool
	// AllowOnlyPolicyNetworksToConnectIfAvailable is "GlobalPolicy.AllowOnlyPolicyNetworksToConnectIfAvailable"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowOnlyPolicyNetworksToConnectIfAvailable MUST be set to true to make this field effective.
	AllowOnlyPolicyNetworksToConnectIfAvailable bool
	// BlockedHexSSIDs is "GlobalPolicy.BlockedHexSSIDs"
	//
	// Optional
	BlockedHexSSIDs js.Array[js.String]

	FFI_USE_AllowOnlyPolicyNetworksToAutoconnect        bool // for AllowOnlyPolicyNetworksToAutoconnect.
	FFI_USE_AllowOnlyPolicyNetworksToConnect            bool // for AllowOnlyPolicyNetworksToConnect.
	FFI_USE_AllowOnlyPolicyNetworksToConnectIfAvailable bool // for AllowOnlyPolicyNetworksToConnectIfAvailable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GlobalPolicy with all fields set.
func (p GlobalPolicy) FromRef(ref js.Ref) GlobalPolicy {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GlobalPolicy in the application heap.
func (p GlobalPolicy) New() js.Ref {
	return bindings.GlobalPolicyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GlobalPolicy) UpdateFrom(ref js.Ref) {
	bindings.GlobalPolicyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GlobalPolicy) Update(ref js.Ref) {
	bindings.GlobalPolicyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GlobalPolicy) FreeMembers(recursive bool) {
	js.Free(
		p.BlockedHexSSIDs.Ref(),
	)
	p.BlockedHexSSIDs = p.BlockedHexSSIDs.FromRef(js.Undefined)
}

type GetManagedPropertiesCallbackFunc func(this js.Ref, result *ManagedProperties) js.Ref

func (fn GetManagedPropertiesCallbackFunc) Register() js.Func[func(result *ManagedProperties)] {
	return js.RegisterCallback[func(result *ManagedProperties)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetManagedPropertiesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ManagedProperties
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

type GetManagedPropertiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *ManagedProperties) js.Ref
	Arg T
}

func (cb *GetManagedPropertiesCallback[T]) Register() js.Func[func(result *ManagedProperties)] {
	return js.RegisterCallback[func(result *ManagedProperties)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetManagedPropertiesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ManagedProperties
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

type ManagedBoolean struct {
	// Active is "ManagedBoolean.Active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// Effective is "ManagedBoolean.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedBoolean.UserPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserPolicy MUST be set to true to make this field effective.
	UserPolicy bool
	// DevicePolicy is "ManagedBoolean.DevicePolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_DevicePolicy MUST be set to true to make this field effective.
	DevicePolicy bool
	// UserSetting is "ManagedBoolean.UserSetting"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserSetting MUST be set to true to make this field effective.
	UserSetting bool
	// SharedSetting is "ManagedBoolean.SharedSetting"
	//
	// Optional
	//
	// NOTE: FFI_USE_SharedSetting MUST be set to true to make this field effective.
	SharedSetting bool
	// UserEditable is "ManagedBoolean.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedBoolean.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_Active         bool // for Active.
	FFI_USE_UserPolicy     bool // for UserPolicy.
	FFI_USE_DevicePolicy   bool // for DevicePolicy.
	FFI_USE_UserSetting    bool // for UserSetting.
	FFI_USE_SharedSetting  bool // for SharedSetting.
	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedBoolean with all fields set.
func (p ManagedBoolean) FromRef(ref js.Ref) ManagedBoolean {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedBoolean in the application heap.
func (p ManagedBoolean) New() js.Ref {
	return bindings.ManagedBooleanJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedBoolean) UpdateFrom(ref js.Ref) {
	bindings.ManagedBooleanJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedBoolean) Update(ref js.Ref) {
	bindings.ManagedBooleanJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedBoolean) FreeMembers(recursive bool) {
	js.Free(
		p.Effective.Ref(),
	)
	p.Effective = p.Effective.FromRef(js.Undefined)
}

type ManagedCellularProperties struct {
	// AutoConnect is "ManagedCellularProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: AutoConnect.FFI_USE MUST be set to true to get AutoConnect used.
	AutoConnect ManagedBoolean
	// ActivationType is "ManagedCellularProperties.ActivationType"
	//
	// Optional
	ActivationType js.String
	// ActivationState is "ManagedCellularProperties.ActivationState"
	//
	// Optional
	ActivationState ActivationStateType
	// AllowRoaming is "ManagedCellularProperties.AllowRoaming"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowRoaming MUST be set to true to make this field effective.
	AllowRoaming bool
	// Family is "ManagedCellularProperties.Family"
	//
	// Optional
	Family js.String
	// FirmwareRevision is "ManagedCellularProperties.FirmwareRevision"
	//
	// Optional
	FirmwareRevision js.String
	// FoundNetworks is "ManagedCellularProperties.FoundNetworks"
	//
	// Optional
	FoundNetworks js.Array[FoundNetworkProperties]
	// HardwareRevision is "ManagedCellularProperties.HardwareRevision"
	//
	// Optional
	HardwareRevision js.String
	// HomeProvider is "ManagedCellularProperties.HomeProvider"
	//
	// Optional
	HomeProvider js.Array[CellularProviderProperties]
	// Manufacturer is "ManagedCellularProperties.Manufacturer"
	//
	// Optional
	Manufacturer js.String
	// ModelID is "ManagedCellularProperties.ModelID"
	//
	// Optional
	ModelID js.String
	// NetworkTechnology is "ManagedCellularProperties.NetworkTechnology"
	//
	// Optional
	NetworkTechnology js.String
	// PaymentPortal is "ManagedCellularProperties.PaymentPortal"
	//
	// Optional
	//
	// NOTE: PaymentPortal.FFI_USE MUST be set to true to get PaymentPortal used.
	PaymentPortal PaymentPortal
	// RoamingState is "ManagedCellularProperties.RoamingState"
	//
	// Optional
	RoamingState js.String
	// Scanning is "ManagedCellularProperties.Scanning"
	//
	// Optional
	//
	// NOTE: FFI_USE_Scanning MUST be set to true to make this field effective.
	Scanning bool
	// ServingOperator is "ManagedCellularProperties.ServingOperator"
	//
	// Optional
	//
	// NOTE: ServingOperator.FFI_USE MUST be set to true to get ServingOperator used.
	ServingOperator CellularProviderProperties
	// SIMLockStatus is "ManagedCellularProperties.SIMLockStatus"
	//
	// Optional
	//
	// NOTE: SIMLockStatus.FFI_USE MUST be set to true to get SIMLockStatus used.
	SIMLockStatus SIMLockStatus
	// SIMPresent is "ManagedCellularProperties.SIMPresent"
	//
	// Optional
	//
	// NOTE: FFI_USE_SIMPresent MUST be set to true to make this field effective.
	SIMPresent bool
	// SignalStrength is "ManagedCellularProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32
	// SupportNetworkScan is "ManagedCellularProperties.SupportNetworkScan"
	//
	// Optional
	//
	// NOTE: FFI_USE_SupportNetworkScan MUST be set to true to make this field effective.
	SupportNetworkScan bool

	FFI_USE_AllowRoaming       bool // for AllowRoaming.
	FFI_USE_Scanning           bool // for Scanning.
	FFI_USE_SIMPresent         bool // for SIMPresent.
	FFI_USE_SignalStrength     bool // for SignalStrength.
	FFI_USE_SupportNetworkScan bool // for SupportNetworkScan.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedCellularProperties with all fields set.
func (p ManagedCellularProperties) FromRef(ref js.Ref) ManagedCellularProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedCellularProperties in the application heap.
func (p ManagedCellularProperties) New() js.Ref {
	return bindings.ManagedCellularPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedCellularProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedCellularPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedCellularProperties) Update(ref js.Ref) {
	bindings.ManagedCellularPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedCellularProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ActivationType.Ref(),
		p.Family.Ref(),
		p.FirmwareRevision.Ref(),
		p.FoundNetworks.Ref(),
		p.HardwareRevision.Ref(),
		p.HomeProvider.Ref(),
		p.Manufacturer.Ref(),
		p.ModelID.Ref(),
		p.NetworkTechnology.Ref(),
		p.RoamingState.Ref(),
	)
	p.ActivationType = p.ActivationType.FromRef(js.Undefined)
	p.Family = p.Family.FromRef(js.Undefined)
	p.FirmwareRevision = p.FirmwareRevision.FromRef(js.Undefined)
	p.FoundNetworks = p.FoundNetworks.FromRef(js.Undefined)
	p.HardwareRevision = p.HardwareRevision.FromRef(js.Undefined)
	p.HomeProvider = p.HomeProvider.FromRef(js.Undefined)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
	p.ModelID = p.ModelID.FromRef(js.Undefined)
	p.NetworkTechnology = p.NetworkTechnology.FromRef(js.Undefined)
	p.RoamingState = p.RoamingState.FromRef(js.Undefined)
	if recursive {
		p.AutoConnect.FreeMembers(true)
		p.PaymentPortal.FreeMembers(true)
		p.ServingOperator.FreeMembers(true)
		p.SIMLockStatus.FreeMembers(true)
	}
}

type ManagedEthernetProperties struct {
	// AutoConnect is "ManagedEthernetProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: AutoConnect.FFI_USE MUST be set to true to get AutoConnect used.
	AutoConnect ManagedBoolean
	// Authentication is "ManagedEthernetProperties.Authentication"
	//
	// Optional
	//
	// NOTE: Authentication.FFI_USE MUST be set to true to get Authentication used.
	Authentication ManagedDOMString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedEthernetProperties with all fields set.
func (p ManagedEthernetProperties) FromRef(ref js.Ref) ManagedEthernetProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedEthernetProperties in the application heap.
func (p ManagedEthernetProperties) New() js.Ref {
	return bindings.ManagedEthernetPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedEthernetProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedEthernetPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedEthernetProperties) Update(ref js.Ref) {
	bindings.ManagedEthernetPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedEthernetProperties) FreeMembers(recursive bool) {
	if recursive {
		p.AutoConnect.FreeMembers(true)
		p.Authentication.FreeMembers(true)
	}
}

type IPConfigType uint32

const (
	_ IPConfigType = iota

	IPConfigType_DHCP
	IPConfigType_STATIC
)

func (IPConfigType) FromRef(str js.Ref) IPConfigType {
	return IPConfigType(bindings.ConstOfIPConfigType(str))
}

func (x IPConfigType) String() (string, bool) {
	switch x {
	case IPConfigType_DHCP:
		return "DHCP", true
	case IPConfigType_STATIC:
		return "Static", true
	default:
		return "", false
	}
}

type ManagedIPConfigType struct {
	// Active is "ManagedIPConfigType.Active"
	//
	// Optional
	Active IPConfigType
	// Effective is "ManagedIPConfigType.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedIPConfigType.UserPolicy"
	//
	// Optional
	UserPolicy IPConfigType
	// DevicePolicy is "ManagedIPConfigType.DevicePolicy"
	//
	// Optional
	DevicePolicy IPConfigType
	// UserSetting is "ManagedIPConfigType.UserSetting"
	//
	// Optional
	UserSetting IPConfigType
	// SharedSetting is "ManagedIPConfigType.SharedSetting"
	//
	// Optional
	SharedSetting IPConfigType
	// UserEditable is "ManagedIPConfigType.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedIPConfigType.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedIPConfigType with all fields set.
func (p ManagedIPConfigType) FromRef(ref js.Ref) ManagedIPConfigType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedIPConfigType in the application heap.
func (p ManagedIPConfigType) New() js.Ref {
	return bindings.ManagedIPConfigTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedIPConfigType) UpdateFrom(ref js.Ref) {
	bindings.ManagedIPConfigTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedIPConfigType) Update(ref js.Ref) {
	bindings.ManagedIPConfigTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedIPConfigType) FreeMembers(recursive bool) {
	js.Free(
		p.Effective.Ref(),
	)
	p.Effective = p.Effective.FromRef(js.Undefined)
}

type IPConfigProperties struct {
	// Gateway is "IPConfigProperties.Gateway"
	//
	// Optional
	Gateway js.String
	// IPAddress is "IPConfigProperties.IPAddress"
	//
	// Optional
	IPAddress js.String
	// ExcludedRoutes is "IPConfigProperties.ExcludedRoutes"
	//
	// Optional
	ExcludedRoutes js.Array[js.String]
	// IncludedRoutes is "IPConfigProperties.IncludedRoutes"
	//
	// Optional
	IncludedRoutes js.Array[js.String]
	// NameServers is "IPConfigProperties.NameServers"
	//
	// Optional
	NameServers js.Array[js.String]
	// SearchDomains is "IPConfigProperties.SearchDomains"
	//
	// Optional
	SearchDomains js.Array[js.String]
	// RoutingPrefix is "IPConfigProperties.RoutingPrefix"
	//
	// Optional
	//
	// NOTE: FFI_USE_RoutingPrefix MUST be set to true to make this field effective.
	RoutingPrefix int32
	// Type is "IPConfigProperties.Type"
	//
	// Optional
	Type js.String
	// WebProxyAutoDiscoveryUrl is "IPConfigProperties.WebProxyAutoDiscoveryUrl"
	//
	// Optional
	WebProxyAutoDiscoveryUrl js.String

	FFI_USE_RoutingPrefix bool // for RoutingPrefix.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IPConfigProperties with all fields set.
func (p IPConfigProperties) FromRef(ref js.Ref) IPConfigProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IPConfigProperties in the application heap.
func (p IPConfigProperties) New() js.Ref {
	return bindings.IPConfigPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IPConfigProperties) UpdateFrom(ref js.Ref) {
	bindings.IPConfigPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IPConfigProperties) Update(ref js.Ref) {
	bindings.IPConfigPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IPConfigProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Gateway.Ref(),
		p.IPAddress.Ref(),
		p.ExcludedRoutes.Ref(),
		p.IncludedRoutes.Ref(),
		p.NameServers.Ref(),
		p.SearchDomains.Ref(),
		p.Type.Ref(),
		p.WebProxyAutoDiscoveryUrl.Ref(),
	)
	p.Gateway = p.Gateway.FromRef(js.Undefined)
	p.IPAddress = p.IPAddress.FromRef(js.Undefined)
	p.ExcludedRoutes = p.ExcludedRoutes.FromRef(js.Undefined)
	p.IncludedRoutes = p.IncludedRoutes.FromRef(js.Undefined)
	p.NameServers = p.NameServers.FromRef(js.Undefined)
	p.SearchDomains = p.SearchDomains.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	p.WebProxyAutoDiscoveryUrl = p.WebProxyAutoDiscoveryUrl.FromRef(js.Undefined)
}

type ManagedLong struct {
	// Active is "ManagedLong.Active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active int32
	// Effective is "ManagedLong.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedLong.UserPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserPolicy MUST be set to true to make this field effective.
	UserPolicy int32
	// DevicePolicy is "ManagedLong.DevicePolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_DevicePolicy MUST be set to true to make this field effective.
	DevicePolicy int32
	// UserSetting is "ManagedLong.UserSetting"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserSetting MUST be set to true to make this field effective.
	UserSetting int32
	// SharedSetting is "ManagedLong.SharedSetting"
	//
	// Optional
	//
	// NOTE: FFI_USE_SharedSetting MUST be set to true to make this field effective.
	SharedSetting int32
	// UserEditable is "ManagedLong.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedLong.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_Active         bool // for Active.
	FFI_USE_UserPolicy     bool // for UserPolicy.
	FFI_USE_DevicePolicy   bool // for DevicePolicy.
	FFI_USE_UserSetting    bool // for UserSetting.
	FFI_USE_SharedSetting  bool // for SharedSetting.
	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedLong with all fields set.
func (p ManagedLong) FromRef(ref js.Ref) ManagedLong {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedLong in the application heap.
func (p ManagedLong) New() js.Ref {
	return bindings.ManagedLongJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedLong) UpdateFrom(ref js.Ref) {
	bindings.ManagedLongJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedLong) Update(ref js.Ref) {
	bindings.ManagedLongJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedLong) FreeMembers(recursive bool) {
	js.Free(
		p.Effective.Ref(),
	)
	p.Effective = p.Effective.FromRef(js.Undefined)
}

type ProxySettingsType uint32

const (
	_ ProxySettingsType = iota

	ProxySettingsType_DIRECT
	ProxySettingsType_MANUAL
	ProxySettingsType_PAC
	ProxySettingsType_WPAD
)

func (ProxySettingsType) FromRef(str js.Ref) ProxySettingsType {
	return ProxySettingsType(bindings.ConstOfProxySettingsType(str))
}

func (x ProxySettingsType) String() (string, bool) {
	switch x {
	case ProxySettingsType_DIRECT:
		return "Direct", true
	case ProxySettingsType_MANUAL:
		return "Manual", true
	case ProxySettingsType_PAC:
		return "PAC", true
	case ProxySettingsType_WPAD:
		return "WPAD", true
	default:
		return "", false
	}
}

type ManagedProxySettingsType struct {
	// Active is "ManagedProxySettingsType.Active"
	//
	// Optional
	Active ProxySettingsType
	// Effective is "ManagedProxySettingsType.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedProxySettingsType.UserPolicy"
	//
	// Optional
	UserPolicy ProxySettingsType
	// DevicePolicy is "ManagedProxySettingsType.DevicePolicy"
	//
	// Optional
	DevicePolicy ProxySettingsType
	// UserSetting is "ManagedProxySettingsType.UserSetting"
	//
	// Optional
	UserSetting ProxySettingsType
	// SharedSetting is "ManagedProxySettingsType.SharedSetting"
	//
	// Optional
	SharedSetting ProxySettingsType
	// UserEditable is "ManagedProxySettingsType.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedProxySettingsType.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedProxySettingsType with all fields set.
func (p ManagedProxySettingsType) FromRef(ref js.Ref) ManagedProxySettingsType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedProxySettingsType in the application heap.
func (p ManagedProxySettingsType) New() js.Ref {
	return bindings.ManagedProxySettingsTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedProxySettingsType) UpdateFrom(ref js.Ref) {
	bindings.ManagedProxySettingsTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedProxySettingsType) Update(ref js.Ref) {
	bindings.ManagedProxySettingsTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedProxySettingsType) FreeMembers(recursive bool) {
	js.Free(
		p.Effective.Ref(),
	)
	p.Effective = p.Effective.FromRef(js.Undefined)
}

type ManagedProxyLocation struct {
	// Host is "ManagedProxyLocation.Host"
	//
	// Optional
	//
	// NOTE: Host.FFI_USE MUST be set to true to get Host used.
	Host ManagedDOMString
	// Port is "ManagedProxyLocation.Port"
	//
	// Optional
	//
	// NOTE: Port.FFI_USE MUST be set to true to get Port used.
	Port ManagedLong

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedProxyLocation with all fields set.
func (p ManagedProxyLocation) FromRef(ref js.Ref) ManagedProxyLocation {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedProxyLocation in the application heap.
func (p ManagedProxyLocation) New() js.Ref {
	return bindings.ManagedProxyLocationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedProxyLocation) UpdateFrom(ref js.Ref) {
	bindings.ManagedProxyLocationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedProxyLocation) Update(ref js.Ref) {
	bindings.ManagedProxyLocationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedProxyLocation) FreeMembers(recursive bool) {
	if recursive {
		p.Host.FreeMembers(true)
		p.Port.FreeMembers(true)
	}
}

type ManagedManualProxySettings struct {
	// HTTPProxy is "ManagedManualProxySettings.HTTPProxy"
	//
	// Optional
	//
	// NOTE: HTTPProxy.FFI_USE MUST be set to true to get HTTPProxy used.
	HTTPProxy ManagedProxyLocation
	// SecureHTTPProxy is "ManagedManualProxySettings.SecureHTTPProxy"
	//
	// Optional
	//
	// NOTE: SecureHTTPProxy.FFI_USE MUST be set to true to get SecureHTTPProxy used.
	SecureHTTPProxy ManagedProxyLocation
	// FTPProxy is "ManagedManualProxySettings.FTPProxy"
	//
	// Optional
	//
	// NOTE: FTPProxy.FFI_USE MUST be set to true to get FTPProxy used.
	FTPProxy ManagedProxyLocation
	// SOCKS is "ManagedManualProxySettings.SOCKS"
	//
	// Optional
	//
	// NOTE: SOCKS.FFI_USE MUST be set to true to get SOCKS used.
	SOCKS ManagedProxyLocation

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedManualProxySettings with all fields set.
func (p ManagedManualProxySettings) FromRef(ref js.Ref) ManagedManualProxySettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedManualProxySettings in the application heap.
func (p ManagedManualProxySettings) New() js.Ref {
	return bindings.ManagedManualProxySettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedManualProxySettings) UpdateFrom(ref js.Ref) {
	bindings.ManagedManualProxySettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedManualProxySettings) Update(ref js.Ref) {
	bindings.ManagedManualProxySettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedManualProxySettings) FreeMembers(recursive bool) {
	if recursive {
		p.HTTPProxy.FreeMembers(true)
		p.SecureHTTPProxy.FreeMembers(true)
		p.FTPProxy.FreeMembers(true)
		p.SOCKS.FreeMembers(true)
	}
}

type ManagedDOMStringList struct {
	// Active is "ManagedDOMStringList.Active"
	//
	// Optional
	Active js.Array[js.String]
	// Effective is "ManagedDOMStringList.Effective"
	//
	// Optional
	Effective js.String
	// UserPolicy is "ManagedDOMStringList.UserPolicy"
	//
	// Optional
	UserPolicy js.Array[js.String]
	// DevicePolicy is "ManagedDOMStringList.DevicePolicy"
	//
	// Optional
	DevicePolicy js.Array[js.String]
	// UserSetting is "ManagedDOMStringList.UserSetting"
	//
	// Optional
	UserSetting js.Array[js.String]
	// SharedSetting is "ManagedDOMStringList.SharedSetting"
	//
	// Optional
	SharedSetting js.Array[js.String]
	// UserEditable is "ManagedDOMStringList.UserEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_UserEditable MUST be set to true to make this field effective.
	UserEditable bool
	// DeviceEditable is "ManagedDOMStringList.DeviceEditable"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceEditable MUST be set to true to make this field effective.
	DeviceEditable bool

	FFI_USE_UserEditable   bool // for UserEditable.
	FFI_USE_DeviceEditable bool // for DeviceEditable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedDOMStringList with all fields set.
func (p ManagedDOMStringList) FromRef(ref js.Ref) ManagedDOMStringList {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedDOMStringList in the application heap.
func (p ManagedDOMStringList) New() js.Ref {
	return bindings.ManagedDOMStringListJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedDOMStringList) UpdateFrom(ref js.Ref) {
	bindings.ManagedDOMStringListJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedDOMStringList) Update(ref js.Ref) {
	bindings.ManagedDOMStringListJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedDOMStringList) FreeMembers(recursive bool) {
	js.Free(
		p.Active.Ref(),
		p.Effective.Ref(),
		p.UserPolicy.Ref(),
		p.DevicePolicy.Ref(),
		p.UserSetting.Ref(),
		p.SharedSetting.Ref(),
	)
	p.Active = p.Active.FromRef(js.Undefined)
	p.Effective = p.Effective.FromRef(js.Undefined)
	p.UserPolicy = p.UserPolicy.FromRef(js.Undefined)
	p.DevicePolicy = p.DevicePolicy.FromRef(js.Undefined)
	p.UserSetting = p.UserSetting.FromRef(js.Undefined)
	p.SharedSetting = p.SharedSetting.FromRef(js.Undefined)
}

type ManagedProxySettings struct {
	// Type is "ManagedProxySettings.Type"
	//
	// Optional
	//
	// NOTE: Type.FFI_USE MUST be set to true to get Type used.
	Type ManagedProxySettingsType
	// Manual is "ManagedProxySettings.Manual"
	//
	// Optional
	//
	// NOTE: Manual.FFI_USE MUST be set to true to get Manual used.
	Manual ManagedManualProxySettings
	// ExcludeDomains is "ManagedProxySettings.ExcludeDomains"
	//
	// Optional
	//
	// NOTE: ExcludeDomains.FFI_USE MUST be set to true to get ExcludeDomains used.
	ExcludeDomains ManagedDOMStringList
	// PAC is "ManagedProxySettings.PAC"
	//
	// Optional
	//
	// NOTE: PAC.FFI_USE MUST be set to true to get PAC used.
	PAC ManagedDOMString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedProxySettings with all fields set.
func (p ManagedProxySettings) FromRef(ref js.Ref) ManagedProxySettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedProxySettings in the application heap.
func (p ManagedProxySettings) New() js.Ref {
	return bindings.ManagedProxySettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedProxySettings) UpdateFrom(ref js.Ref) {
	bindings.ManagedProxySettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedProxySettings) Update(ref js.Ref) {
	bindings.ManagedProxySettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedProxySettings) FreeMembers(recursive bool) {
	if recursive {
		p.Type.FreeMembers(true)
		p.Manual.FreeMembers(true)
		p.ExcludeDomains.FreeMembers(true)
		p.PAC.FreeMembers(true)
	}
}

type ManagedIPConfigProperties struct {
	// Gateway is "ManagedIPConfigProperties.Gateway"
	//
	// Optional
	//
	// NOTE: Gateway.FFI_USE MUST be set to true to get Gateway used.
	Gateway ManagedDOMString
	// IPAddress is "ManagedIPConfigProperties.IPAddress"
	//
	// Optional
	//
	// NOTE: IPAddress.FFI_USE MUST be set to true to get IPAddress used.
	IPAddress ManagedDOMString
	// NameServers is "ManagedIPConfigProperties.NameServers"
	//
	// Optional
	//
	// NOTE: NameServers.FFI_USE MUST be set to true to get NameServers used.
	NameServers ManagedDOMStringList
	// RoutingPrefix is "ManagedIPConfigProperties.RoutingPrefix"
	//
	// Optional
	//
	// NOTE: RoutingPrefix.FFI_USE MUST be set to true to get RoutingPrefix used.
	RoutingPrefix ManagedLong
	// Type is "ManagedIPConfigProperties.Type"
	//
	// Optional
	//
	// NOTE: Type.FFI_USE MUST be set to true to get Type used.
	Type ManagedDOMString
	// WebProxyAutoDiscoveryUrl is "ManagedIPConfigProperties.WebProxyAutoDiscoveryUrl"
	//
	// Optional
	//
	// NOTE: WebProxyAutoDiscoveryUrl.FFI_USE MUST be set to true to get WebProxyAutoDiscoveryUrl used.
	WebProxyAutoDiscoveryUrl ManagedDOMString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedIPConfigProperties with all fields set.
func (p ManagedIPConfigProperties) FromRef(ref js.Ref) ManagedIPConfigProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedIPConfigProperties in the application heap.
func (p ManagedIPConfigProperties) New() js.Ref {
	return bindings.ManagedIPConfigPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedIPConfigProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedIPConfigPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedIPConfigProperties) Update(ref js.Ref) {
	bindings.ManagedIPConfigPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedIPConfigProperties) FreeMembers(recursive bool) {
	if recursive {
		p.Gateway.FreeMembers(true)
		p.IPAddress.FreeMembers(true)
		p.NameServers.FreeMembers(true)
		p.RoutingPrefix.FreeMembers(true)
		p.Type.FreeMembers(true)
		p.WebProxyAutoDiscoveryUrl.FreeMembers(true)
	}
}

type ManagedVPNProperties struct {
	// AutoConnect is "ManagedVPNProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: AutoConnect.FFI_USE MUST be set to true to get AutoConnect used.
	AutoConnect ManagedBoolean
	// Host is "ManagedVPNProperties.Host"
	//
	// Optional
	//
	// NOTE: Host.FFI_USE MUST be set to true to get Host used.
	Host ManagedDOMString
	// Type is "ManagedVPNProperties.Type"
	//
	// Optional
	//
	// NOTE: Type.FFI_USE MUST be set to true to get Type used.
	Type ManagedDOMString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedVPNProperties with all fields set.
func (p ManagedVPNProperties) FromRef(ref js.Ref) ManagedVPNProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedVPNProperties in the application heap.
func (p ManagedVPNProperties) New() js.Ref {
	return bindings.ManagedVPNPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedVPNProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedVPNPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedVPNProperties) Update(ref js.Ref) {
	bindings.ManagedVPNPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedVPNProperties) FreeMembers(recursive bool) {
	if recursive {
		p.AutoConnect.FreeMembers(true)
		p.Host.FreeMembers(true)
		p.Type.FreeMembers(true)
	}
}

type ManagedWiFiProperties struct {
	// AllowGatewayARPPolling is "ManagedWiFiProperties.AllowGatewayARPPolling"
	//
	// Optional
	//
	// NOTE: AllowGatewayARPPolling.FFI_USE MUST be set to true to get AllowGatewayARPPolling used.
	AllowGatewayARPPolling ManagedBoolean
	// AutoConnect is "ManagedWiFiProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: AutoConnect.FFI_USE MUST be set to true to get AutoConnect used.
	AutoConnect ManagedBoolean
	// BSSID is "ManagedWiFiProperties.BSSID"
	//
	// Optional
	BSSID js.String
	// Frequency is "ManagedWiFiProperties.Frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency int32
	// FrequencyList is "ManagedWiFiProperties.FrequencyList"
	//
	// Optional
	FrequencyList js.Array[int32]
	// HexSSID is "ManagedWiFiProperties.HexSSID"
	//
	// Optional
	//
	// NOTE: HexSSID.FFI_USE MUST be set to true to get HexSSID used.
	HexSSID ManagedDOMString
	// HiddenSSID is "ManagedWiFiProperties.HiddenSSID"
	//
	// Optional
	//
	// NOTE: HiddenSSID.FFI_USE MUST be set to true to get HiddenSSID used.
	HiddenSSID ManagedBoolean
	// RoamThreshold is "ManagedWiFiProperties.RoamThreshold"
	//
	// Optional
	//
	// NOTE: RoamThreshold.FFI_USE MUST be set to true to get RoamThreshold used.
	RoamThreshold ManagedLong
	// SSID is "ManagedWiFiProperties.SSID"
	//
	// Optional
	//
	// NOTE: SSID.FFI_USE MUST be set to true to get SSID used.
	SSID ManagedDOMString
	// Security is "ManagedWiFiProperties.Security"
	//
	// Optional
	//
	// NOTE: Security.FFI_USE MUST be set to true to get Security used.
	Security ManagedDOMString
	// SignalStrength is "ManagedWiFiProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32

	FFI_USE_Frequency      bool // for Frequency.
	FFI_USE_SignalStrength bool // for SignalStrength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedWiFiProperties with all fields set.
func (p ManagedWiFiProperties) FromRef(ref js.Ref) ManagedWiFiProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedWiFiProperties in the application heap.
func (p ManagedWiFiProperties) New() js.Ref {
	return bindings.ManagedWiFiPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedWiFiProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedWiFiPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedWiFiProperties) Update(ref js.Ref) {
	bindings.ManagedWiFiPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedWiFiProperties) FreeMembers(recursive bool) {
	js.Free(
		p.BSSID.Ref(),
		p.FrequencyList.Ref(),
	)
	p.BSSID = p.BSSID.FromRef(js.Undefined)
	p.FrequencyList = p.FrequencyList.FromRef(js.Undefined)
	if recursive {
		p.AllowGatewayARPPolling.FreeMembers(true)
		p.AutoConnect.FreeMembers(true)
		p.HexSSID.FreeMembers(true)
		p.HiddenSSID.FreeMembers(true)
		p.RoamThreshold.FreeMembers(true)
		p.SSID.FreeMembers(true)
		p.Security.FreeMembers(true)
	}
}

type ManagedProperties struct {
	// Cellular is "ManagedProperties.Cellular"
	//
	// Optional
	//
	// NOTE: Cellular.FFI_USE MUST be set to true to get Cellular used.
	Cellular ManagedCellularProperties
	// Connectable is "ManagedProperties.Connectable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connectable MUST be set to true to make this field effective.
	Connectable bool
	// ConnectionState is "ManagedProperties.ConnectionState"
	//
	// Optional
	ConnectionState ConnectionStateType
	// ErrorState is "ManagedProperties.ErrorState"
	//
	// Optional
	ErrorState js.String
	// Ethernet is "ManagedProperties.Ethernet"
	//
	// Optional
	//
	// NOTE: Ethernet.FFI_USE MUST be set to true to get Ethernet used.
	Ethernet ManagedEthernetProperties
	// GUID is "ManagedProperties.GUID"
	//
	// Optional
	GUID js.String
	// IPAddressConfigType is "ManagedProperties.IPAddressConfigType"
	//
	// Optional
	//
	// NOTE: IPAddressConfigType.FFI_USE MUST be set to true to get IPAddressConfigType used.
	IPAddressConfigType ManagedIPConfigType
	// IPConfigs is "ManagedProperties.IPConfigs"
	//
	// Optional
	IPConfigs js.Array[IPConfigProperties]
	// MacAddress is "ManagedProperties.MacAddress"
	//
	// Optional
	MacAddress js.String
	// Metered is "ManagedProperties.Metered"
	//
	// Optional
	//
	// NOTE: Metered.FFI_USE MUST be set to true to get Metered used.
	Metered ManagedBoolean
	// Name is "ManagedProperties.Name"
	//
	// Optional
	//
	// NOTE: Name.FFI_USE MUST be set to true to get Name used.
	Name ManagedDOMString
	// NameServersConfigType is "ManagedProperties.NameServersConfigType"
	//
	// Optional
	//
	// NOTE: NameServersConfigType.FFI_USE MUST be set to true to get NameServersConfigType used.
	NameServersConfigType ManagedIPConfigType
	// Priority is "ManagedProperties.Priority"
	//
	// Optional
	//
	// NOTE: Priority.FFI_USE MUST be set to true to get Priority used.
	Priority ManagedLong
	// ProxySettings is "ManagedProperties.ProxySettings"
	//
	// Optional
	//
	// NOTE: ProxySettings.FFI_USE MUST be set to true to get ProxySettings used.
	ProxySettings ManagedProxySettings
	// RestrictedConnectivity is "ManagedProperties.RestrictedConnectivity"
	//
	// Optional
	//
	// NOTE: FFI_USE_RestrictedConnectivity MUST be set to true to make this field effective.
	RestrictedConnectivity bool
	// StaticIPConfig is "ManagedProperties.StaticIPConfig"
	//
	// Optional
	//
	// NOTE: StaticIPConfig.FFI_USE MUST be set to true to get StaticIPConfig used.
	StaticIPConfig ManagedIPConfigProperties
	// SavedIPConfig is "ManagedProperties.SavedIPConfig"
	//
	// Optional
	//
	// NOTE: SavedIPConfig.FFI_USE MUST be set to true to get SavedIPConfig used.
	SavedIPConfig IPConfigProperties
	// Source is "ManagedProperties.Source"
	//
	// Optional
	Source js.String
	// Type is "ManagedProperties.Type"
	//
	// Optional
	Type NetworkType
	// VPN is "ManagedProperties.VPN"
	//
	// Optional
	//
	// NOTE: VPN.FFI_USE MUST be set to true to get VPN used.
	VPN ManagedVPNProperties
	// WiFi is "ManagedProperties.WiFi"
	//
	// Optional
	//
	// NOTE: WiFi.FFI_USE MUST be set to true to get WiFi used.
	WiFi ManagedWiFiProperties

	FFI_USE_Connectable            bool // for Connectable.
	FFI_USE_RestrictedConnectivity bool // for RestrictedConnectivity.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedProperties with all fields set.
func (p ManagedProperties) FromRef(ref js.Ref) ManagedProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedProperties in the application heap.
func (p ManagedProperties) New() js.Ref {
	return bindings.ManagedPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedProperties) Update(ref js.Ref) {
	bindings.ManagedPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ErrorState.Ref(),
		p.GUID.Ref(),
		p.IPConfigs.Ref(),
		p.MacAddress.Ref(),
		p.Source.Ref(),
	)
	p.ErrorState = p.ErrorState.FromRef(js.Undefined)
	p.GUID = p.GUID.FromRef(js.Undefined)
	p.IPConfigs = p.IPConfigs.FromRef(js.Undefined)
	p.MacAddress = p.MacAddress.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	if recursive {
		p.Cellular.FreeMembers(true)
		p.Ethernet.FreeMembers(true)
		p.IPAddressConfigType.FreeMembers(true)
		p.Metered.FreeMembers(true)
		p.Name.FreeMembers(true)
		p.NameServersConfigType.FreeMembers(true)
		p.Priority.FreeMembers(true)
		p.ProxySettings.FreeMembers(true)
		p.StaticIPConfig.FreeMembers(true)
		p.SavedIPConfig.FreeMembers(true)
		p.VPN.FreeMembers(true)
		p.WiFi.FreeMembers(true)
	}
}

type GetNetworksCallbackFunc func(this js.Ref, result js.Array[NetworkStateProperties]) js.Ref

func (fn GetNetworksCallbackFunc) Register() js.Func[func(result js.Array[NetworkStateProperties])] {
	return js.RegisterCallback[func(result js.Array[NetworkStateProperties])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetNetworksCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[NetworkStateProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetNetworksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[NetworkStateProperties]) js.Ref
	Arg T
}

func (cb *GetNetworksCallback[T]) Register() js.Func[func(result js.Array[NetworkStateProperties])] {
	return js.RegisterCallback[func(result js.Array[NetworkStateProperties])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetNetworksCallback[T]) DispatchCallback(
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

		js.Array[NetworkStateProperties]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VPNStateProperties struct {
	// Type is "VPNStateProperties.Type"
	//
	// Optional
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VPNStateProperties with all fields set.
func (p VPNStateProperties) FromRef(ref js.Ref) VPNStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VPNStateProperties in the application heap.
func (p VPNStateProperties) New() js.Ref {
	return bindings.VPNStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VPNStateProperties) UpdateFrom(ref js.Ref) {
	bindings.VPNStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VPNStateProperties) Update(ref js.Ref) {
	bindings.VPNStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VPNStateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
}

type WiFiStateProperties struct {
	// BSSID is "WiFiStateProperties.BSSID"
	//
	// Optional
	BSSID js.String
	// Frequency is "WiFiStateProperties.Frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency int32
	// HexSSID is "WiFiStateProperties.HexSSID"
	//
	// Optional
	HexSSID js.String
	// Security is "WiFiStateProperties.Security"
	//
	// Optional
	Security js.String
	// SignalStrength is "WiFiStateProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32
	// SSID is "WiFiStateProperties.SSID"
	//
	// Optional
	SSID js.String

	FFI_USE_Frequency      bool // for Frequency.
	FFI_USE_SignalStrength bool // for SignalStrength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WiFiStateProperties with all fields set.
func (p WiFiStateProperties) FromRef(ref js.Ref) WiFiStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WiFiStateProperties in the application heap.
func (p WiFiStateProperties) New() js.Ref {
	return bindings.WiFiStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WiFiStateProperties) UpdateFrom(ref js.Ref) {
	bindings.WiFiStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WiFiStateProperties) Update(ref js.Ref) {
	bindings.WiFiStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WiFiStateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.BSSID.Ref(),
		p.HexSSID.Ref(),
		p.Security.Ref(),
		p.SSID.Ref(),
	)
	p.BSSID = p.BSSID.FromRef(js.Undefined)
	p.HexSSID = p.HexSSID.FromRef(js.Undefined)
	p.Security = p.Security.FromRef(js.Undefined)
	p.SSID = p.SSID.FromRef(js.Undefined)
}

type NetworkStateProperties struct {
	// Cellular is "NetworkStateProperties.Cellular"
	//
	// Optional
	//
	// NOTE: Cellular.FFI_USE MUST be set to true to get Cellular used.
	Cellular CellularStateProperties
	// Connectable is "NetworkStateProperties.Connectable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connectable MUST be set to true to make this field effective.
	Connectable bool
	// ConnectionState is "NetworkStateProperties.ConnectionState"
	//
	// Optional
	ConnectionState ConnectionStateType
	// Ethernet is "NetworkStateProperties.Ethernet"
	//
	// Optional
	//
	// NOTE: Ethernet.FFI_USE MUST be set to true to get Ethernet used.
	Ethernet EthernetStateProperties
	// ErrorState is "NetworkStateProperties.ErrorState"
	//
	// Optional
	ErrorState js.String
	// GUID is "NetworkStateProperties.GUID"
	//
	// Optional
	GUID js.String
	// Name is "NetworkStateProperties.Name"
	//
	// Optional
	Name js.String
	// Priority is "NetworkStateProperties.Priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// Source is "NetworkStateProperties.Source"
	//
	// Optional
	Source js.String
	// Type is "NetworkStateProperties.Type"
	//
	// Optional
	Type NetworkType
	// VPN is "NetworkStateProperties.VPN"
	//
	// Optional
	//
	// NOTE: VPN.FFI_USE MUST be set to true to get VPN used.
	VPN VPNStateProperties
	// WiFi is "NetworkStateProperties.WiFi"
	//
	// Optional
	//
	// NOTE: WiFi.FFI_USE MUST be set to true to get WiFi used.
	WiFi WiFiStateProperties

	FFI_USE_Connectable bool // for Connectable.
	FFI_USE_Priority    bool // for Priority.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkStateProperties with all fields set.
func (p NetworkStateProperties) FromRef(ref js.Ref) NetworkStateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkStateProperties in the application heap.
func (p NetworkStateProperties) New() js.Ref {
	return bindings.NetworkStatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkStateProperties) UpdateFrom(ref js.Ref) {
	bindings.NetworkStatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkStateProperties) Update(ref js.Ref) {
	bindings.NetworkStatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkStateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ErrorState.Ref(),
		p.GUID.Ref(),
		p.Name.Ref(),
		p.Source.Ref(),
	)
	p.ErrorState = p.ErrorState.FromRef(js.Undefined)
	p.GUID = p.GUID.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	if recursive {
		p.Cellular.FreeMembers(true)
		p.Ethernet.FreeMembers(true)
		p.VPN.FreeMembers(true)
		p.WiFi.FreeMembers(true)
	}
}

type GetPropertiesCallbackFunc func(this js.Ref, result *NetworkProperties) js.Ref

func (fn GetPropertiesCallbackFunc) Register() js.Func[func(result *NetworkProperties)] {
	return js.RegisterCallback[func(result *NetworkProperties)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPropertiesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkProperties
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

type GetPropertiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *NetworkProperties) js.Ref
	Arg T
}

func (cb *GetPropertiesCallback[T]) Register() js.Func[func(result *NetworkProperties)] {
	return js.RegisterCallback[func(result *NetworkProperties)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPropertiesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkProperties
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

type ProxyLocation struct {
	// Host is "ProxyLocation.Host"
	//
	// Optional
	Host js.String
	// Port is "ProxyLocation.Port"
	//
	// Optional
	//
	// NOTE: FFI_USE_Port MUST be set to true to make this field effective.
	Port int32

	FFI_USE_Port bool // for Port.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProxyLocation with all fields set.
func (p ProxyLocation) FromRef(ref js.Ref) ProxyLocation {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProxyLocation in the application heap.
func (p ProxyLocation) New() js.Ref {
	return bindings.ProxyLocationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProxyLocation) UpdateFrom(ref js.Ref) {
	bindings.ProxyLocationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProxyLocation) Update(ref js.Ref) {
	bindings.ProxyLocationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProxyLocation) FreeMembers(recursive bool) {
	js.Free(
		p.Host.Ref(),
	)
	p.Host = p.Host.FromRef(js.Undefined)
}

type ManualProxySettings struct {
	// HTTPProxy is "ManualProxySettings.HTTPProxy"
	//
	// Optional
	//
	// NOTE: HTTPProxy.FFI_USE MUST be set to true to get HTTPProxy used.
	HTTPProxy ProxyLocation
	// SecureHTTPProxy is "ManualProxySettings.SecureHTTPProxy"
	//
	// Optional
	//
	// NOTE: SecureHTTPProxy.FFI_USE MUST be set to true to get SecureHTTPProxy used.
	SecureHTTPProxy ProxyLocation
	// FTPProxy is "ManualProxySettings.FTPProxy"
	//
	// Optional
	//
	// NOTE: FTPProxy.FFI_USE MUST be set to true to get FTPProxy used.
	FTPProxy ProxyLocation
	// SOCKS is "ManualProxySettings.SOCKS"
	//
	// Optional
	//
	// NOTE: SOCKS.FFI_USE MUST be set to true to get SOCKS used.
	SOCKS ProxyLocation

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManualProxySettings with all fields set.
func (p ManualProxySettings) FromRef(ref js.Ref) ManualProxySettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManualProxySettings in the application heap.
func (p ManualProxySettings) New() js.Ref {
	return bindings.ManualProxySettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManualProxySettings) UpdateFrom(ref js.Ref) {
	bindings.ManualProxySettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManualProxySettings) Update(ref js.Ref) {
	bindings.ManualProxySettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManualProxySettings) FreeMembers(recursive bool) {
	if recursive {
		p.HTTPProxy.FreeMembers(true)
		p.SecureHTTPProxy.FreeMembers(true)
		p.FTPProxy.FreeMembers(true)
		p.SOCKS.FreeMembers(true)
	}
}

type ProxySettings struct {
	// Type is "ProxySettings.Type"
	//
	// Optional
	Type ProxySettingsType
	// Manual is "ProxySettings.Manual"
	//
	// Optional
	//
	// NOTE: Manual.FFI_USE MUST be set to true to get Manual used.
	Manual ManualProxySettings
	// ExcludeDomains is "ProxySettings.ExcludeDomains"
	//
	// Optional
	ExcludeDomains js.Array[js.String]
	// PAC is "ProxySettings.PAC"
	//
	// Optional
	PAC js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProxySettings with all fields set.
func (p ProxySettings) FromRef(ref js.Ref) ProxySettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProxySettings in the application heap.
func (p ProxySettings) New() js.Ref {
	return bindings.ProxySettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProxySettings) UpdateFrom(ref js.Ref) {
	bindings.ProxySettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProxySettings) Update(ref js.Ref) {
	bindings.ProxySettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProxySettings) FreeMembers(recursive bool) {
	js.Free(
		p.ExcludeDomains.Ref(),
		p.PAC.Ref(),
	)
	p.ExcludeDomains = p.ExcludeDomains.FromRef(js.Undefined)
	p.PAC = p.PAC.FromRef(js.Undefined)
	if recursive {
		p.Manual.FreeMembers(true)
	}
}

type VPNProperties struct {
	// AutoConnect is "VPNProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoConnect MUST be set to true to make this field effective.
	AutoConnect bool
	// Host is "VPNProperties.Host"
	//
	// Optional
	Host js.String
	// Type is "VPNProperties.Type"
	//
	// Optional
	Type js.String

	FFI_USE_AutoConnect bool // for AutoConnect.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VPNProperties with all fields set.
func (p VPNProperties) FromRef(ref js.Ref) VPNProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VPNProperties in the application heap.
func (p VPNProperties) New() js.Ref {
	return bindings.VPNPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VPNProperties) UpdateFrom(ref js.Ref) {
	bindings.VPNPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VPNProperties) Update(ref js.Ref) {
	bindings.VPNPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VPNProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Host.Ref(),
		p.Type.Ref(),
	)
	p.Host = p.Host.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
}

type WiFiProperties struct {
	// AllowGatewayARPPolling is "WiFiProperties.AllowGatewayARPPolling"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowGatewayARPPolling MUST be set to true to make this field effective.
	AllowGatewayARPPolling bool
	// AutoConnect is "WiFiProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoConnect MUST be set to true to make this field effective.
	AutoConnect bool
	// BSSID is "WiFiProperties.BSSID"
	//
	// Optional
	BSSID js.String
	// EAP is "WiFiProperties.EAP"
	//
	// Optional
	//
	// NOTE: EAP.FFI_USE MUST be set to true to get EAP used.
	EAP EAPProperties
	// Frequency is "WiFiProperties.Frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency int32
	// FrequencyList is "WiFiProperties.FrequencyList"
	//
	// Optional
	FrequencyList js.Array[int32]
	// HexSSID is "WiFiProperties.HexSSID"
	//
	// Optional
	HexSSID js.String
	// HiddenSSID is "WiFiProperties.HiddenSSID"
	//
	// Optional
	//
	// NOTE: FFI_USE_HiddenSSID MUST be set to true to make this field effective.
	HiddenSSID bool
	// Passphrase is "WiFiProperties.Passphrase"
	//
	// Optional
	Passphrase js.String
	// RoamThreshold is "WiFiProperties.RoamThreshold"
	//
	// Optional
	//
	// NOTE: FFI_USE_RoamThreshold MUST be set to true to make this field effective.
	RoamThreshold int32
	// SSID is "WiFiProperties.SSID"
	//
	// Optional
	SSID js.String
	// Security is "WiFiProperties.Security"
	//
	// Optional
	Security js.String
	// SignalStrength is "WiFiProperties.SignalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength int32

	FFI_USE_AllowGatewayARPPolling bool // for AllowGatewayARPPolling.
	FFI_USE_AutoConnect            bool // for AutoConnect.
	FFI_USE_Frequency              bool // for Frequency.
	FFI_USE_HiddenSSID             bool // for HiddenSSID.
	FFI_USE_RoamThreshold          bool // for RoamThreshold.
	FFI_USE_SignalStrength         bool // for SignalStrength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WiFiProperties with all fields set.
func (p WiFiProperties) FromRef(ref js.Ref) WiFiProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WiFiProperties in the application heap.
func (p WiFiProperties) New() js.Ref {
	return bindings.WiFiPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WiFiProperties) UpdateFrom(ref js.Ref) {
	bindings.WiFiPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WiFiProperties) Update(ref js.Ref) {
	bindings.WiFiPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WiFiProperties) FreeMembers(recursive bool) {
	js.Free(
		p.BSSID.Ref(),
		p.FrequencyList.Ref(),
		p.HexSSID.Ref(),
		p.Passphrase.Ref(),
		p.SSID.Ref(),
		p.Security.Ref(),
	)
	p.BSSID = p.BSSID.FromRef(js.Undefined)
	p.FrequencyList = p.FrequencyList.FromRef(js.Undefined)
	p.HexSSID = p.HexSSID.FromRef(js.Undefined)
	p.Passphrase = p.Passphrase.FromRef(js.Undefined)
	p.SSID = p.SSID.FromRef(js.Undefined)
	p.Security = p.Security.FromRef(js.Undefined)
	if recursive {
		p.EAP.FreeMembers(true)
	}
}

type NetworkProperties struct {
	// Cellular is "NetworkProperties.Cellular"
	//
	// Optional
	//
	// NOTE: Cellular.FFI_USE MUST be set to true to get Cellular used.
	Cellular CellularProperties
	// Connectable is "NetworkProperties.Connectable"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connectable MUST be set to true to make this field effective.
	Connectable bool
	// ConnectionState is "NetworkProperties.ConnectionState"
	//
	// Optional
	ConnectionState ConnectionStateType
	// ErrorState is "NetworkProperties.ErrorState"
	//
	// Optional
	ErrorState js.String
	// Ethernet is "NetworkProperties.Ethernet"
	//
	// Optional
	//
	// NOTE: Ethernet.FFI_USE MUST be set to true to get Ethernet used.
	Ethernet EthernetProperties
	// GUID is "NetworkProperties.GUID"
	//
	// Optional
	GUID js.String
	// IPAddressConfigType is "NetworkProperties.IPAddressConfigType"
	//
	// Optional
	IPAddressConfigType IPConfigType
	// IPConfigs is "NetworkProperties.IPConfigs"
	//
	// Optional
	IPConfigs js.Array[IPConfigProperties]
	// MacAddress is "NetworkProperties.MacAddress"
	//
	// Optional
	MacAddress js.String
	// Metered is "NetworkProperties.Metered"
	//
	// Optional
	//
	// NOTE: FFI_USE_Metered MUST be set to true to make this field effective.
	Metered bool
	// Name is "NetworkProperties.Name"
	//
	// Optional
	Name js.String
	// NameServersConfigType is "NetworkProperties.NameServersConfigType"
	//
	// Optional
	NameServersConfigType IPConfigType
	// Priority is "NetworkProperties.Priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// ProxySettings is "NetworkProperties.ProxySettings"
	//
	// Optional
	//
	// NOTE: ProxySettings.FFI_USE MUST be set to true to get ProxySettings used.
	ProxySettings ProxySettings
	// RestrictedConnectivity is "NetworkProperties.RestrictedConnectivity"
	//
	// Optional
	//
	// NOTE: FFI_USE_RestrictedConnectivity MUST be set to true to make this field effective.
	RestrictedConnectivity bool
	// StaticIPConfig is "NetworkProperties.StaticIPConfig"
	//
	// Optional
	//
	// NOTE: StaticIPConfig.FFI_USE MUST be set to true to get StaticIPConfig used.
	StaticIPConfig IPConfigProperties
	// SavedIPConfig is "NetworkProperties.SavedIPConfig"
	//
	// Optional
	//
	// NOTE: SavedIPConfig.FFI_USE MUST be set to true to get SavedIPConfig used.
	SavedIPConfig IPConfigProperties
	// Source is "NetworkProperties.Source"
	//
	// Optional
	Source js.String
	// Type is "NetworkProperties.Type"
	//
	// Optional
	Type NetworkType
	// VPN is "NetworkProperties.VPN"
	//
	// Optional
	//
	// NOTE: VPN.FFI_USE MUST be set to true to get VPN used.
	VPN VPNProperties
	// WiFi is "NetworkProperties.WiFi"
	//
	// Optional
	//
	// NOTE: WiFi.FFI_USE MUST be set to true to get WiFi used.
	WiFi WiFiProperties

	FFI_USE_Connectable            bool // for Connectable.
	FFI_USE_Metered                bool // for Metered.
	FFI_USE_Priority               bool // for Priority.
	FFI_USE_RestrictedConnectivity bool // for RestrictedConnectivity.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkProperties with all fields set.
func (p NetworkProperties) FromRef(ref js.Ref) NetworkProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkProperties in the application heap.
func (p NetworkProperties) New() js.Ref {
	return bindings.NetworkPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkProperties) UpdateFrom(ref js.Ref) {
	bindings.NetworkPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkProperties) Update(ref js.Ref) {
	bindings.NetworkPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ErrorState.Ref(),
		p.GUID.Ref(),
		p.IPConfigs.Ref(),
		p.MacAddress.Ref(),
		p.Name.Ref(),
		p.Source.Ref(),
	)
	p.ErrorState = p.ErrorState.FromRef(js.Undefined)
	p.GUID = p.GUID.FromRef(js.Undefined)
	p.IPConfigs = p.IPConfigs.FromRef(js.Undefined)
	p.MacAddress = p.MacAddress.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	if recursive {
		p.Cellular.FreeMembers(true)
		p.Ethernet.FreeMembers(true)
		p.ProxySettings.FreeMembers(true)
		p.StaticIPConfig.FreeMembers(true)
		p.SavedIPConfig.FreeMembers(true)
		p.VPN.FreeMembers(true)
		p.WiFi.FreeMembers(true)
	}
}

type GetStatePropertiesCallbackFunc func(this js.Ref, result *NetworkStateProperties) js.Ref

func (fn GetStatePropertiesCallbackFunc) Register() js.Func[func(result *NetworkStateProperties)] {
	return js.RegisterCallback[func(result *NetworkStateProperties)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetStatePropertiesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkStateProperties
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

type GetStatePropertiesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *NetworkStateProperties) js.Ref
	Arg T
}

func (cb *GetStatePropertiesCallback[T]) Register() js.Func[func(result *NetworkStateProperties)] {
	return js.RegisterCallback[func(result *NetworkStateProperties)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetStatePropertiesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkStateProperties
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

type ManagedThirdPartyVPNProperties struct {
	// ExtensionID is "ManagedThirdPartyVPNProperties.ExtensionID"
	//
	// Optional
	//
	// NOTE: ExtensionID.FFI_USE MUST be set to true to get ExtensionID used.
	ExtensionID ManagedDOMString
	// ProviderName is "ManagedThirdPartyVPNProperties.ProviderName"
	//
	// Optional
	ProviderName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManagedThirdPartyVPNProperties with all fields set.
func (p ManagedThirdPartyVPNProperties) FromRef(ref js.Ref) ManagedThirdPartyVPNProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManagedThirdPartyVPNProperties in the application heap.
func (p ManagedThirdPartyVPNProperties) New() js.Ref {
	return bindings.ManagedThirdPartyVPNPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManagedThirdPartyVPNProperties) UpdateFrom(ref js.Ref) {
	bindings.ManagedThirdPartyVPNPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManagedThirdPartyVPNProperties) Update(ref js.Ref) {
	bindings.ManagedThirdPartyVPNPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManagedThirdPartyVPNProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ProviderName.Ref(),
	)
	p.ProviderName = p.ProviderName.FromRef(js.Undefined)
	if recursive {
		p.ExtensionID.FreeMembers(true)
	}
}

type WiMAXProperties struct {
	// AutoConnect is "WiMAXProperties.AutoConnect"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoConnect MUST be set to true to make this field effective.
	AutoConnect bool
	// EAP is "WiMAXProperties.EAP"
	//
	// Optional
	//
	// NOTE: EAP.FFI_USE MUST be set to true to get EAP used.
	EAP EAPProperties

	FFI_USE_AutoConnect bool // for AutoConnect.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WiMAXProperties with all fields set.
func (p WiMAXProperties) FromRef(ref js.Ref) WiMAXProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WiMAXProperties in the application heap.
func (p WiMAXProperties) New() js.Ref {
	return bindings.WiMAXPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WiMAXProperties) UpdateFrom(ref js.Ref) {
	bindings.WiMAXPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WiMAXProperties) Update(ref js.Ref) {
	bindings.WiMAXPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WiMAXProperties) FreeMembers(recursive bool) {
	if recursive {
		p.EAP.FreeMembers(true)
	}
}

type NetworkConfigProperties struct {
	// Cellular is "NetworkConfigProperties.Cellular"
	//
	// Optional
	//
	// NOTE: Cellular.FFI_USE MUST be set to true to get Cellular used.
	Cellular CellularProperties
	// Ethernet is "NetworkConfigProperties.Ethernet"
	//
	// Optional
	//
	// NOTE: Ethernet.FFI_USE MUST be set to true to get Ethernet used.
	Ethernet EthernetProperties
	// GUID is "NetworkConfigProperties.GUID"
	//
	// Optional
	GUID js.String
	// IPAddressConfigType is "NetworkConfigProperties.IPAddressConfigType"
	//
	// Optional
	IPAddressConfigType IPConfigType
	// Name is "NetworkConfigProperties.Name"
	//
	// Optional
	Name js.String
	// NameServersConfigType is "NetworkConfigProperties.NameServersConfigType"
	//
	// Optional
	NameServersConfigType IPConfigType
	// Priority is "NetworkConfigProperties.Priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int32
	// Type is "NetworkConfigProperties.Type"
	//
	// Optional
	Type NetworkType
	// VPN is "NetworkConfigProperties.VPN"
	//
	// Optional
	//
	// NOTE: VPN.FFI_USE MUST be set to true to get VPN used.
	VPN VPNProperties
	// WiFi is "NetworkConfigProperties.WiFi"
	//
	// Optional
	//
	// NOTE: WiFi.FFI_USE MUST be set to true to get WiFi used.
	WiFi WiFiProperties
	// WiMAX is "NetworkConfigProperties.WiMAX"
	//
	// Optional
	//
	// NOTE: WiMAX.FFI_USE MUST be set to true to get WiMAX used.
	WiMAX WiMAXProperties

	FFI_USE_Priority bool // for Priority.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkConfigProperties with all fields set.
func (p NetworkConfigProperties) FromRef(ref js.Ref) NetworkConfigProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkConfigProperties in the application heap.
func (p NetworkConfigProperties) New() js.Ref {
	return bindings.NetworkConfigPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkConfigProperties) UpdateFrom(ref js.Ref) {
	bindings.NetworkConfigPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkConfigProperties) Update(ref js.Ref) {
	bindings.NetworkConfigPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkConfigProperties) FreeMembers(recursive bool) {
	js.Free(
		p.GUID.Ref(),
		p.Name.Ref(),
	)
	p.GUID = p.GUID.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	if recursive {
		p.Cellular.FreeMembers(true)
		p.Ethernet.FreeMembers(true)
		p.VPN.FreeMembers(true)
		p.WiFi.FreeMembers(true)
		p.WiMAX.FreeMembers(true)
	}
}

type NetworkFilter struct {
	// NetworkType is "NetworkFilter.networkType"
	//
	// Optional
	NetworkType NetworkType
	// Visible is "NetworkFilter.visible"
	//
	// Optional
	//
	// NOTE: FFI_USE_Visible MUST be set to true to make this field effective.
	Visible bool
	// Configured is "NetworkFilter.configured"
	//
	// Optional
	//
	// NOTE: FFI_USE_Configured MUST be set to true to make this field effective.
	Configured bool
	// Limit is "NetworkFilter.limit"
	//
	// Optional
	//
	// NOTE: FFI_USE_Limit MUST be set to true to make this field effective.
	Limit int32

	FFI_USE_Visible    bool // for Visible.
	FFI_USE_Configured bool // for Configured.
	FFI_USE_Limit      bool // for Limit.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkFilter with all fields set.
func (p NetworkFilter) FromRef(ref js.Ref) NetworkFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkFilter in the application heap.
func (p NetworkFilter) New() js.Ref {
	return bindings.NetworkFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkFilter) UpdateFrom(ref js.Ref) {
	bindings.NetworkFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkFilter) Update(ref js.Ref) {
	bindings.NetworkFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkFilter) FreeMembers(recursive bool) {
}

type StringCallbackFunc func(this js.Ref, result js.String) js.Ref

func (fn StringCallbackFunc) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
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
	Fn  func(arg T, this js.Ref, result js.String) js.Ref
	Arg T
}

func (cb *StringCallback[T]) Register() js.Func[func(result js.String)] {
	return js.RegisterCallback[func(result js.String)](
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

type ThirdPartyVPNProperties struct {
	// ExtensionID is "ThirdPartyVPNProperties.ExtensionID"
	//
	// Optional
	ExtensionID js.String
	// ProviderName is "ThirdPartyVPNProperties.ProviderName"
	//
	// Optional
	ProviderName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ThirdPartyVPNProperties with all fields set.
func (p ThirdPartyVPNProperties) FromRef(ref js.Ref) ThirdPartyVPNProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ThirdPartyVPNProperties in the application heap.
func (p ThirdPartyVPNProperties) New() js.Ref {
	return bindings.ThirdPartyVPNPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ThirdPartyVPNProperties) UpdateFrom(ref js.Ref) {
	bindings.ThirdPartyVPNPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ThirdPartyVPNProperties) Update(ref js.Ref) {
	bindings.ThirdPartyVPNPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ThirdPartyVPNProperties) FreeMembers(recursive bool) {
	js.Free(
		p.ExtensionID.Ref(),
		p.ProviderName.Ref(),
	)
	p.ExtensionID = p.ExtensionID.FromRef(js.Undefined)
	p.ProviderName = p.ProviderName.FromRef(js.Undefined)
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

// HasFuncCreateNetwork returns true if the function "WEBEXT.networking.onc.createNetwork" exists.
func HasFuncCreateNetwork() bool {
	return js.True == bindings.HasFuncCreateNetwork()
}

// FuncCreateNetwork returns the function "WEBEXT.networking.onc.createNetwork".
func FuncCreateNetwork() (fn js.Func[func(shared bool, properties NetworkConfigProperties, callback js.Func[func(result js.String)])]) {
	bindings.FuncCreateNetwork(
		js.Pointer(&fn),
	)
	return
}

// CreateNetwork calls the function "WEBEXT.networking.onc.createNetwork" directly.
func CreateNetwork(shared bool, properties NetworkConfigProperties, callback js.Func[func(result js.String)]) (ret js.Void) {
	bindings.CallCreateNetwork(
		js.Pointer(&ret),
		js.Bool(bool(shared)),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryCreateNetwork calls the function "WEBEXT.networking.onc.createNetwork"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateNetwork(shared bool, properties NetworkConfigProperties, callback js.Func[func(result js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateNetwork(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(shared)),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// HasFuncDisableNetworkType returns true if the function "WEBEXT.networking.onc.disableNetworkType" exists.
func HasFuncDisableNetworkType() bool {
	return js.True == bindings.HasFuncDisableNetworkType()
}

// FuncDisableNetworkType returns the function "WEBEXT.networking.onc.disableNetworkType".
func FuncDisableNetworkType() (fn js.Func[func(networkType NetworkType)]) {
	bindings.FuncDisableNetworkType(
		js.Pointer(&fn),
	)
	return
}

// DisableNetworkType calls the function "WEBEXT.networking.onc.disableNetworkType" directly.
func DisableNetworkType(networkType NetworkType) (ret js.Void) {
	bindings.CallDisableNetworkType(
		js.Pointer(&ret),
		uint32(networkType),
	)

	return
}

// TryDisableNetworkType calls the function "WEBEXT.networking.onc.disableNetworkType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisableNetworkType(networkType NetworkType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisableNetworkType(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(networkType),
	)

	return
}

// HasFuncEnableNetworkType returns true if the function "WEBEXT.networking.onc.enableNetworkType" exists.
func HasFuncEnableNetworkType() bool {
	return js.True == bindings.HasFuncEnableNetworkType()
}

// FuncEnableNetworkType returns the function "WEBEXT.networking.onc.enableNetworkType".
func FuncEnableNetworkType() (fn js.Func[func(networkType NetworkType)]) {
	bindings.FuncEnableNetworkType(
		js.Pointer(&fn),
	)
	return
}

// EnableNetworkType calls the function "WEBEXT.networking.onc.enableNetworkType" directly.
func EnableNetworkType(networkType NetworkType) (ret js.Void) {
	bindings.CallEnableNetworkType(
		js.Pointer(&ret),
		uint32(networkType),
	)

	return
}

// TryEnableNetworkType calls the function "WEBEXT.networking.onc.enableNetworkType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableNetworkType(networkType NetworkType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableNetworkType(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(networkType),
	)

	return
}

// HasFuncForgetNetwork returns true if the function "WEBEXT.networking.onc.forgetNetwork" exists.
func HasFuncForgetNetwork() bool {
	return js.True == bindings.HasFuncForgetNetwork()
}

// FuncForgetNetwork returns the function "WEBEXT.networking.onc.forgetNetwork".
func FuncForgetNetwork() (fn js.Func[func(networkGuid js.String, callback js.Func[func()])]) {
	bindings.FuncForgetNetwork(
		js.Pointer(&fn),
	)
	return
}

// ForgetNetwork calls the function "WEBEXT.networking.onc.forgetNetwork" directly.
func ForgetNetwork(networkGuid js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallForgetNetwork(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryForgetNetwork calls the function "WEBEXT.networking.onc.forgetNetwork"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForgetNetwork(networkGuid js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryForgetNetwork(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetCaptivePortalStatus returns true if the function "WEBEXT.networking.onc.getCaptivePortalStatus" exists.
func HasFuncGetCaptivePortalStatus() bool {
	return js.True == bindings.HasFuncGetCaptivePortalStatus()
}

// FuncGetCaptivePortalStatus returns the function "WEBEXT.networking.onc.getCaptivePortalStatus".
func FuncGetCaptivePortalStatus() (fn js.Func[func(networkGuid js.String, callback js.Func[func(result CaptivePortalStatus)])]) {
	bindings.FuncGetCaptivePortalStatus(
		js.Pointer(&fn),
	)
	return
}

// GetCaptivePortalStatus calls the function "WEBEXT.networking.onc.getCaptivePortalStatus" directly.
func GetCaptivePortalStatus(networkGuid js.String, callback js.Func[func(result CaptivePortalStatus)]) (ret js.Void) {
	bindings.CallGetCaptivePortalStatus(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetCaptivePortalStatus calls the function "WEBEXT.networking.onc.getCaptivePortalStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCaptivePortalStatus(networkGuid js.String, callback js.Func[func(result CaptivePortalStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCaptivePortalStatus(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetDeviceStates returns true if the function "WEBEXT.networking.onc.getDeviceStates" exists.
func HasFuncGetDeviceStates() bool {
	return js.True == bindings.HasFuncGetDeviceStates()
}

// FuncGetDeviceStates returns the function "WEBEXT.networking.onc.getDeviceStates".
func FuncGetDeviceStates() (fn js.Func[func(callback js.Func[func(result js.Array[DeviceStateProperties])])]) {
	bindings.FuncGetDeviceStates(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceStates calls the function "WEBEXT.networking.onc.getDeviceStates" directly.
func GetDeviceStates(callback js.Func[func(result js.Array[DeviceStateProperties])]) (ret js.Void) {
	bindings.CallGetDeviceStates(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetDeviceStates calls the function "WEBEXT.networking.onc.getDeviceStates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceStates(callback js.Func[func(result js.Array[DeviceStateProperties])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceStates(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetGlobalPolicy returns true if the function "WEBEXT.networking.onc.getGlobalPolicy" exists.
func HasFuncGetGlobalPolicy() bool {
	return js.True == bindings.HasFuncGetGlobalPolicy()
}

// FuncGetGlobalPolicy returns the function "WEBEXT.networking.onc.getGlobalPolicy".
func FuncGetGlobalPolicy() (fn js.Func[func(callback js.Func[func(result *GlobalPolicy)])]) {
	bindings.FuncGetGlobalPolicy(
		js.Pointer(&fn),
	)
	return
}

// GetGlobalPolicy calls the function "WEBEXT.networking.onc.getGlobalPolicy" directly.
func GetGlobalPolicy(callback js.Func[func(result *GlobalPolicy)]) (ret js.Void) {
	bindings.CallGetGlobalPolicy(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetGlobalPolicy calls the function "WEBEXT.networking.onc.getGlobalPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetGlobalPolicy(callback js.Func[func(result *GlobalPolicy)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetGlobalPolicy(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetManagedProperties returns true if the function "WEBEXT.networking.onc.getManagedProperties" exists.
func HasFuncGetManagedProperties() bool {
	return js.True == bindings.HasFuncGetManagedProperties()
}

// FuncGetManagedProperties returns the function "WEBEXT.networking.onc.getManagedProperties".
func FuncGetManagedProperties() (fn js.Func[func(networkGuid js.String, callback js.Func[func(result *ManagedProperties)])]) {
	bindings.FuncGetManagedProperties(
		js.Pointer(&fn),
	)
	return
}

// GetManagedProperties calls the function "WEBEXT.networking.onc.getManagedProperties" directly.
func GetManagedProperties(networkGuid js.String, callback js.Func[func(result *ManagedProperties)]) (ret js.Void) {
	bindings.CallGetManagedProperties(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetManagedProperties calls the function "WEBEXT.networking.onc.getManagedProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetManagedProperties(networkGuid js.String, callback js.Func[func(result *ManagedProperties)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetManagedProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetNetworks returns true if the function "WEBEXT.networking.onc.getNetworks" exists.
func HasFuncGetNetworks() bool {
	return js.True == bindings.HasFuncGetNetworks()
}

// FuncGetNetworks returns the function "WEBEXT.networking.onc.getNetworks".
func FuncGetNetworks() (fn js.Func[func(filter NetworkFilter, callback js.Func[func(result js.Array[NetworkStateProperties])])]) {
	bindings.FuncGetNetworks(
		js.Pointer(&fn),
	)
	return
}

// GetNetworks calls the function "WEBEXT.networking.onc.getNetworks" directly.
func GetNetworks(filter NetworkFilter, callback js.Func[func(result js.Array[NetworkStateProperties])]) (ret js.Void) {
	bindings.CallGetNetworks(
		js.Pointer(&ret),
		js.Pointer(&filter),
		callback.Ref(),
	)

	return
}

// TryGetNetworks calls the function "WEBEXT.networking.onc.getNetworks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNetworks(filter NetworkFilter, callback js.Func[func(result js.Array[NetworkStateProperties])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNetworks(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
		callback.Ref(),
	)

	return
}

// HasFuncGetProperties returns true if the function "WEBEXT.networking.onc.getProperties" exists.
func HasFuncGetProperties() bool {
	return js.True == bindings.HasFuncGetProperties()
}

// FuncGetProperties returns the function "WEBEXT.networking.onc.getProperties".
func FuncGetProperties() (fn js.Func[func(networkGuid js.String, callback js.Func[func(result *NetworkProperties)])]) {
	bindings.FuncGetProperties(
		js.Pointer(&fn),
	)
	return
}

// GetProperties calls the function "WEBEXT.networking.onc.getProperties" directly.
func GetProperties(networkGuid js.String, callback js.Func[func(result *NetworkProperties)]) (ret js.Void) {
	bindings.CallGetProperties(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetProperties calls the function "WEBEXT.networking.onc.getProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProperties(networkGuid js.String, callback js.Func[func(result *NetworkProperties)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetState returns true if the function "WEBEXT.networking.onc.getState" exists.
func HasFuncGetState() bool {
	return js.True == bindings.HasFuncGetState()
}

// FuncGetState returns the function "WEBEXT.networking.onc.getState".
func FuncGetState() (fn js.Func[func(networkGuid js.String, callback js.Func[func(result *NetworkStateProperties)])]) {
	bindings.FuncGetState(
		js.Pointer(&fn),
	)
	return
}

// GetState calls the function "WEBEXT.networking.onc.getState" directly.
func GetState(networkGuid js.String, callback js.Func[func(result *NetworkStateProperties)]) (ret js.Void) {
	bindings.CallGetState(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetState calls the function "WEBEXT.networking.onc.getState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetState(networkGuid js.String, callback js.Func[func(result *NetworkStateProperties)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetState(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

type OnDeviceStateListChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnDeviceStateListChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceStateListChangedEventCallbackFunc) DispatchCallback(
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

type OnDeviceStateListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnDeviceStateListChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceStateListChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDeviceStateListChanged returns true if the function "WEBEXT.networking.onc.onDeviceStateListChanged.addListener" exists.
func HasFuncOnDeviceStateListChanged() bool {
	return js.True == bindings.HasFuncOnDeviceStateListChanged()
}

// FuncOnDeviceStateListChanged returns the function "WEBEXT.networking.onc.onDeviceStateListChanged.addListener".
func FuncOnDeviceStateListChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnDeviceStateListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.addListener" directly.
func OnDeviceStateListChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnDeviceStateListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceStateListChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceStateListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceStateListChanged returns true if the function "WEBEXT.networking.onc.onDeviceStateListChanged.removeListener" exists.
func HasFuncOffDeviceStateListChanged() bool {
	return js.True == bindings.HasFuncOffDeviceStateListChanged()
}

// FuncOffDeviceStateListChanged returns the function "WEBEXT.networking.onc.onDeviceStateListChanged.removeListener".
func FuncOffDeviceStateListChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffDeviceStateListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.removeListener" directly.
func OffDeviceStateListChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffDeviceStateListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceStateListChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceStateListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceStateListChanged returns true if the function "WEBEXT.networking.onc.onDeviceStateListChanged.hasListener" exists.
func HasFuncHasOnDeviceStateListChanged() bool {
	return js.True == bindings.HasFuncHasOnDeviceStateListChanged()
}

// FuncHasOnDeviceStateListChanged returns the function "WEBEXT.networking.onc.onDeviceStateListChanged.hasListener".
func FuncHasOnDeviceStateListChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnDeviceStateListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.hasListener" directly.
func HasOnDeviceStateListChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnDeviceStateListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceStateListChanged calls the function "WEBEXT.networking.onc.onDeviceStateListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceStateListChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceStateListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnNetworkListChangedEventCallbackFunc func(this js.Ref, changes js.Array[js.String]) js.Ref

func (fn OnNetworkListChangedEventCallbackFunc) Register() js.Func[func(changes js.Array[js.String])] {
	return js.RegisterCallback[func(changes js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnNetworkListChangedEventCallbackFunc) DispatchCallback(
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

type OnNetworkListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, changes js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnNetworkListChangedEventCallback[T]) Register() js.Func[func(changes js.Array[js.String])] {
	return js.RegisterCallback[func(changes js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnNetworkListChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnNetworkListChanged returns true if the function "WEBEXT.networking.onc.onNetworkListChanged.addListener" exists.
func HasFuncOnNetworkListChanged() bool {
	return js.True == bindings.HasFuncOnNetworkListChanged()
}

// FuncOnNetworkListChanged returns the function "WEBEXT.networking.onc.onNetworkListChanged.addListener".
func FuncOnNetworkListChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])])]) {
	bindings.FuncOnNetworkListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.addListener" directly.
func OnNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnNetworkListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnNetworkListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffNetworkListChanged returns true if the function "WEBEXT.networking.onc.onNetworkListChanged.removeListener" exists.
func HasFuncOffNetworkListChanged() bool {
	return js.True == bindings.HasFuncOffNetworkListChanged()
}

// FuncOffNetworkListChanged returns the function "WEBEXT.networking.onc.onNetworkListChanged.removeListener".
func FuncOffNetworkListChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])])]) {
	bindings.FuncOffNetworkListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.removeListener" directly.
func OffNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffNetworkListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffNetworkListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnNetworkListChanged returns true if the function "WEBEXT.networking.onc.onNetworkListChanged.hasListener" exists.
func HasFuncHasOnNetworkListChanged() bool {
	return js.True == bindings.HasFuncHasOnNetworkListChanged()
}

// FuncHasOnNetworkListChanged returns the function "WEBEXT.networking.onc.onNetworkListChanged.hasListener".
func FuncHasOnNetworkListChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])]) bool]) {
	bindings.FuncHasOnNetworkListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.hasListener" directly.
func HasOnNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnNetworkListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnNetworkListChanged calls the function "WEBEXT.networking.onc.onNetworkListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnNetworkListChanged(callback js.Func[func(changes js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnNetworkListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnNetworksChangedEventCallbackFunc func(this js.Ref, changes js.Array[js.String]) js.Ref

func (fn OnNetworksChangedEventCallbackFunc) Register() js.Func[func(changes js.Array[js.String])] {
	return js.RegisterCallback[func(changes js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnNetworksChangedEventCallbackFunc) DispatchCallback(
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

type OnNetworksChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, changes js.Array[js.String]) js.Ref
	Arg T
}

func (cb *OnNetworksChangedEventCallback[T]) Register() js.Func[func(changes js.Array[js.String])] {
	return js.RegisterCallback[func(changes js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnNetworksChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnNetworksChanged returns true if the function "WEBEXT.networking.onc.onNetworksChanged.addListener" exists.
func HasFuncOnNetworksChanged() bool {
	return js.True == bindings.HasFuncOnNetworksChanged()
}

// FuncOnNetworksChanged returns the function "WEBEXT.networking.onc.onNetworksChanged.addListener".
func FuncOnNetworksChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])])]) {
	bindings.FuncOnNetworksChanged(
		js.Pointer(&fn),
	)
	return
}

// OnNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.addListener" directly.
func OnNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void) {
	bindings.CallOnNetworksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnNetworksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffNetworksChanged returns true if the function "WEBEXT.networking.onc.onNetworksChanged.removeListener" exists.
func HasFuncOffNetworksChanged() bool {
	return js.True == bindings.HasFuncOffNetworksChanged()
}

// FuncOffNetworksChanged returns the function "WEBEXT.networking.onc.onNetworksChanged.removeListener".
func FuncOffNetworksChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])])]) {
	bindings.FuncOffNetworksChanged(
		js.Pointer(&fn),
	)
	return
}

// OffNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.removeListener" directly.
func OffNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void) {
	bindings.CallOffNetworksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffNetworksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnNetworksChanged returns true if the function "WEBEXT.networking.onc.onNetworksChanged.hasListener" exists.
func HasFuncHasOnNetworksChanged() bool {
	return js.True == bindings.HasFuncHasOnNetworksChanged()
}

// FuncHasOnNetworksChanged returns the function "WEBEXT.networking.onc.onNetworksChanged.hasListener".
func FuncHasOnNetworksChanged() (fn js.Func[func(callback js.Func[func(changes js.Array[js.String])]) bool]) {
	bindings.FuncHasOnNetworksChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.hasListener" directly.
func HasOnNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret bool) {
	bindings.CallHasOnNetworksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnNetworksChanged calls the function "WEBEXT.networking.onc.onNetworksChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnNetworksChanged(callback js.Func[func(changes js.Array[js.String])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnNetworksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPortalDetectionCompletedEventCallbackFunc func(this js.Ref, networkGuid js.String, status CaptivePortalStatus) js.Ref

func (fn OnPortalDetectionCompletedEventCallbackFunc) Register() js.Func[func(networkGuid js.String, status CaptivePortalStatus)] {
	return js.RegisterCallback[func(networkGuid js.String, status CaptivePortalStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPortalDetectionCompletedEventCallbackFunc) DispatchCallback(
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
		CaptivePortalStatus(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPortalDetectionCompletedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, networkGuid js.String, status CaptivePortalStatus) js.Ref
	Arg T
}

func (cb *OnPortalDetectionCompletedEventCallback[T]) Register() js.Func[func(networkGuid js.String, status CaptivePortalStatus)] {
	return js.RegisterCallback[func(networkGuid js.String, status CaptivePortalStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPortalDetectionCompletedEventCallback[T]) DispatchCallback(
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
		CaptivePortalStatus(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPortalDetectionCompleted returns true if the function "WEBEXT.networking.onc.onPortalDetectionCompleted.addListener" exists.
func HasFuncOnPortalDetectionCompleted() bool {
	return js.True == bindings.HasFuncOnPortalDetectionCompleted()
}

// FuncOnPortalDetectionCompleted returns the function "WEBEXT.networking.onc.onPortalDetectionCompleted.addListener".
func FuncOnPortalDetectionCompleted() (fn js.Func[func(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)])]) {
	bindings.FuncOnPortalDetectionCompleted(
		js.Pointer(&fn),
	)
	return
}

// OnPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.addListener" directly.
func OnPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret js.Void) {
	bindings.CallOnPortalDetectionCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPortalDetectionCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPortalDetectionCompleted returns true if the function "WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener" exists.
func HasFuncOffPortalDetectionCompleted() bool {
	return js.True == bindings.HasFuncOffPortalDetectionCompleted()
}

// FuncOffPortalDetectionCompleted returns the function "WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener".
func FuncOffPortalDetectionCompleted() (fn js.Func[func(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)])]) {
	bindings.FuncOffPortalDetectionCompleted(
		js.Pointer(&fn),
	)
	return
}

// OffPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener" directly.
func OffPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret js.Void) {
	bindings.CallOffPortalDetectionCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPortalDetectionCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPortalDetectionCompleted returns true if the function "WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener" exists.
func HasFuncHasOnPortalDetectionCompleted() bool {
	return js.True == bindings.HasFuncHasOnPortalDetectionCompleted()
}

// FuncHasOnPortalDetectionCompleted returns the function "WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener".
func FuncHasOnPortalDetectionCompleted() (fn js.Func[func(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) bool]) {
	bindings.FuncHasOnPortalDetectionCompleted(
		js.Pointer(&fn),
	)
	return
}

// HasOnPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener" directly.
func HasOnPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret bool) {
	bindings.CallHasOnPortalDetectionCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPortalDetectionCompleted calls the function "WEBEXT.networking.onc.onPortalDetectionCompleted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPortalDetectionCompleted(callback js.Func[func(networkGuid js.String, status CaptivePortalStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPortalDetectionCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRequestNetworkScan returns true if the function "WEBEXT.networking.onc.requestNetworkScan" exists.
func HasFuncRequestNetworkScan() bool {
	return js.True == bindings.HasFuncRequestNetworkScan()
}

// FuncRequestNetworkScan returns the function "WEBEXT.networking.onc.requestNetworkScan".
func FuncRequestNetworkScan() (fn js.Func[func(networkType NetworkType)]) {
	bindings.FuncRequestNetworkScan(
		js.Pointer(&fn),
	)
	return
}

// RequestNetworkScan calls the function "WEBEXT.networking.onc.requestNetworkScan" directly.
func RequestNetworkScan(networkType NetworkType) (ret js.Void) {
	bindings.CallRequestNetworkScan(
		js.Pointer(&ret),
		uint32(networkType),
	)

	return
}

// TryRequestNetworkScan calls the function "WEBEXT.networking.onc.requestNetworkScan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestNetworkScan(networkType NetworkType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestNetworkScan(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(networkType),
	)

	return
}

// HasFuncSetProperties returns true if the function "WEBEXT.networking.onc.setProperties" exists.
func HasFuncSetProperties() bool {
	return js.True == bindings.HasFuncSetProperties()
}

// FuncSetProperties returns the function "WEBEXT.networking.onc.setProperties".
func FuncSetProperties() (fn js.Func[func(networkGuid js.String, properties NetworkConfigProperties, callback js.Func[func()])]) {
	bindings.FuncSetProperties(
		js.Pointer(&fn),
	)
	return
}

// SetProperties calls the function "WEBEXT.networking.onc.setProperties" directly.
func SetProperties(networkGuid js.String, properties NetworkConfigProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetProperties(
		js.Pointer(&ret),
		networkGuid.Ref(),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TrySetProperties calls the function "WEBEXT.networking.onc.setProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetProperties(networkGuid js.String, properties NetworkConfigProperties, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// HasFuncStartConnect returns true if the function "WEBEXT.networking.onc.startConnect" exists.
func HasFuncStartConnect() bool {
	return js.True == bindings.HasFuncStartConnect()
}

// FuncStartConnect returns the function "WEBEXT.networking.onc.startConnect".
func FuncStartConnect() (fn js.Func[func(networkGuid js.String, callback js.Func[func()])]) {
	bindings.FuncStartConnect(
		js.Pointer(&fn),
	)
	return
}

// StartConnect calls the function "WEBEXT.networking.onc.startConnect" directly.
func StartConnect(networkGuid js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStartConnect(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryStartConnect calls the function "WEBEXT.networking.onc.startConnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartConnect(networkGuid js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncStartDisconnect returns true if the function "WEBEXT.networking.onc.startDisconnect" exists.
func HasFuncStartDisconnect() bool {
	return js.True == bindings.HasFuncStartDisconnect()
}

// FuncStartDisconnect returns the function "WEBEXT.networking.onc.startDisconnect".
func FuncStartDisconnect() (fn js.Func[func(networkGuid js.String, callback js.Func[func()])]) {
	bindings.FuncStartDisconnect(
		js.Pointer(&fn),
	)
	return
}

// StartDisconnect calls the function "WEBEXT.networking.onc.startDisconnect" directly.
func StartDisconnect(networkGuid js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallStartDisconnect(
		js.Pointer(&ret),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

// TryStartDisconnect calls the function "WEBEXT.networking.onc.startDisconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartDisconnect(networkGuid js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		networkGuid.Ref(),
		callback.Ref(),
	)

	return
}

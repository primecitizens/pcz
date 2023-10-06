// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type COSEAlgorithmIdentifier int32

type AuthenticatorAttestationResponse struct {
	AuthenticatorResponse
}

func (this AuthenticatorAttestationResponse) Once() AuthenticatorAttestationResponse {
	this.Ref().Once()
	return this
}

func (this AuthenticatorAttestationResponse) Ref() js.Ref {
	return this.AuthenticatorResponse.Ref()
}

func (this AuthenticatorAttestationResponse) FromRef(ref js.Ref) AuthenticatorAttestationResponse {
	this.AuthenticatorResponse = this.AuthenticatorResponse.FromRef(ref)
	return this
}

func (this AuthenticatorAttestationResponse) Free() {
	this.Ref().Free()
}

// AttestationObject returns the value of property "AuthenticatorAttestationResponse.attestationObject".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAttestationResponse) AttestationObject() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAttestationResponseAttestationObject(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetTransports returns true if the method "AuthenticatorAttestationResponse.getTransports" exists.
func (this AuthenticatorAttestationResponse) HasGetTransports() bool {
	return js.True == bindings.HasAuthenticatorAttestationResponseGetTransports(
		this.Ref(),
	)
}

// GetTransportsFunc returns the method "AuthenticatorAttestationResponse.getTransports".
func (this AuthenticatorAttestationResponse) GetTransportsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.AuthenticatorAttestationResponseGetTransportsFunc(
			this.Ref(),
		),
	)
}

// GetTransports calls the method "AuthenticatorAttestationResponse.getTransports".
func (this AuthenticatorAttestationResponse) GetTransports() (ret js.Array[js.String]) {
	bindings.CallAuthenticatorAttestationResponseGetTransports(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTransports calls the method "AuthenticatorAttestationResponse.getTransports"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AuthenticatorAttestationResponse) TryGetTransports() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticatorAttestationResponseGetTransports(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAuthenticatorData returns true if the method "AuthenticatorAttestationResponse.getAuthenticatorData" exists.
func (this AuthenticatorAttestationResponse) HasGetAuthenticatorData() bool {
	return js.True == bindings.HasAuthenticatorAttestationResponseGetAuthenticatorData(
		this.Ref(),
	)
}

// GetAuthenticatorDataFunc returns the method "AuthenticatorAttestationResponse.getAuthenticatorData".
func (this AuthenticatorAttestationResponse) GetAuthenticatorDataFunc() (fn js.Func[func() js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.AuthenticatorAttestationResponseGetAuthenticatorDataFunc(
			this.Ref(),
		),
	)
}

// GetAuthenticatorData calls the method "AuthenticatorAttestationResponse.getAuthenticatorData".
func (this AuthenticatorAttestationResponse) GetAuthenticatorData() (ret js.ArrayBuffer) {
	bindings.CallAuthenticatorAttestationResponseGetAuthenticatorData(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAuthenticatorData calls the method "AuthenticatorAttestationResponse.getAuthenticatorData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AuthenticatorAttestationResponse) TryGetAuthenticatorData() (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticatorAttestationResponseGetAuthenticatorData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPublicKey returns true if the method "AuthenticatorAttestationResponse.getPublicKey" exists.
func (this AuthenticatorAttestationResponse) HasGetPublicKey() bool {
	return js.True == bindings.HasAuthenticatorAttestationResponseGetPublicKey(
		this.Ref(),
	)
}

// GetPublicKeyFunc returns the method "AuthenticatorAttestationResponse.getPublicKey".
func (this AuthenticatorAttestationResponse) GetPublicKeyFunc() (fn js.Func[func() js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.AuthenticatorAttestationResponseGetPublicKeyFunc(
			this.Ref(),
		),
	)
}

// GetPublicKey calls the method "AuthenticatorAttestationResponse.getPublicKey".
func (this AuthenticatorAttestationResponse) GetPublicKey() (ret js.ArrayBuffer) {
	bindings.CallAuthenticatorAttestationResponseGetPublicKey(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPublicKey calls the method "AuthenticatorAttestationResponse.getPublicKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AuthenticatorAttestationResponse) TryGetPublicKey() (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticatorAttestationResponseGetPublicKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPublicKeyAlgorithm returns true if the method "AuthenticatorAttestationResponse.getPublicKeyAlgorithm" exists.
func (this AuthenticatorAttestationResponse) HasGetPublicKeyAlgorithm() bool {
	return js.True == bindings.HasAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
		this.Ref(),
	)
}

// GetPublicKeyAlgorithmFunc returns the method "AuthenticatorAttestationResponse.getPublicKeyAlgorithm".
func (this AuthenticatorAttestationResponse) GetPublicKeyAlgorithmFunc() (fn js.Func[func() COSEAlgorithmIdentifier]) {
	return fn.FromRef(
		bindings.AuthenticatorAttestationResponseGetPublicKeyAlgorithmFunc(
			this.Ref(),
		),
	)
}

// GetPublicKeyAlgorithm calls the method "AuthenticatorAttestationResponse.getPublicKeyAlgorithm".
func (this AuthenticatorAttestationResponse) GetPublicKeyAlgorithm() (ret COSEAlgorithmIdentifier) {
	bindings.CallAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPublicKeyAlgorithm calls the method "AuthenticatorAttestationResponse.getPublicKeyAlgorithm"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AuthenticatorAttestationResponse) TryGetPublicKeyAlgorithm() (ret COSEAlgorithmIdentifier, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AuthenticatorAttestationResponseJSON struct {
	// ClientDataJSON is "AuthenticatorAttestationResponseJSON.clientDataJSON"
	//
	// Required
	ClientDataJSON Base64URLString
	// AuthenticatorData is "AuthenticatorAttestationResponseJSON.authenticatorData"
	//
	// Required
	AuthenticatorData Base64URLString
	// Transports is "AuthenticatorAttestationResponseJSON.transports"
	//
	// Required
	Transports js.Array[js.String]
	// PublicKey is "AuthenticatorAttestationResponseJSON.publicKey"
	//
	// Optional
	PublicKey Base64URLString
	// PublicKeyAlgorithm is "AuthenticatorAttestationResponseJSON.publicKeyAlgorithm"
	//
	// Required
	PublicKeyAlgorithm int64
	// AttestationObject is "AuthenticatorAttestationResponseJSON.attestationObject"
	//
	// Required
	AttestationObject Base64URLString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticatorAttestationResponseJSON with all fields set.
func (p AuthenticatorAttestationResponseJSON) FromRef(ref js.Ref) AuthenticatorAttestationResponseJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticatorAttestationResponseJSON in the application heap.
func (p AuthenticatorAttestationResponseJSON) New() js.Ref {
	return bindings.AuthenticatorAttestationResponseJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuthenticatorAttestationResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticatorAttestationResponseJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticatorAttestationResponseJSON) Update(ref js.Ref) {
	bindings.AuthenticatorAttestationResponseJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuthenticatorResponse struct {
	ref js.Ref
}

func (this AuthenticatorResponse) Once() AuthenticatorResponse {
	this.Ref().Once()
	return this
}

func (this AuthenticatorResponse) Ref() js.Ref {
	return this.ref
}

func (this AuthenticatorResponse) FromRef(ref js.Ref) AuthenticatorResponse {
	this.ref = ref
	return this
}

func (this AuthenticatorResponse) Free() {
	this.Ref().Free()
}

// ClientDataJSON returns the value of property "AuthenticatorResponse.clientDataJSON".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorResponse) ClientDataJSON() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorResponseClientDataJSON(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AuthenticatorSelectionCriteria struct {
	// AuthenticatorAttachment is "AuthenticatorSelectionCriteria.authenticatorAttachment"
	//
	// Optional
	AuthenticatorAttachment js.String
	// ResidentKey is "AuthenticatorSelectionCriteria.residentKey"
	//
	// Optional
	ResidentKey js.String
	// RequireResidentKey is "AuthenticatorSelectionCriteria.requireResidentKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequireResidentKey MUST be set to true to make this field effective.
	RequireResidentKey bool
	// UserVerification is "AuthenticatorSelectionCriteria.userVerification"
	//
	// Optional, defaults to "preferred".
	UserVerification js.String

	FFI_USE_RequireResidentKey bool // for RequireResidentKey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticatorSelectionCriteria with all fields set.
func (p AuthenticatorSelectionCriteria) FromRef(ref js.Ref) AuthenticatorSelectionCriteria {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticatorSelectionCriteria in the application heap.
func (p AuthenticatorSelectionCriteria) New() js.Ref {
	return bindings.AuthenticatorSelectionCriteriaJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuthenticatorSelectionCriteria) UpdateFrom(ref js.Ref) {
	bindings.AuthenticatorSelectionCriteriaJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticatorSelectionCriteria) Update(ref js.Ref) {
	bindings.AuthenticatorSelectionCriteriaJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuthenticatorTransport uint32

const (
	_ AuthenticatorTransport = iota

	AuthenticatorTransport_USB
	AuthenticatorTransport_NFC
	AuthenticatorTransport_BLE
	AuthenticatorTransport_SMART_CARD
	AuthenticatorTransport_HYBRID
	AuthenticatorTransport_INTERNAL
)

func (AuthenticatorTransport) FromRef(str js.Ref) AuthenticatorTransport {
	return AuthenticatorTransport(bindings.ConstOfAuthenticatorTransport(str))
}

func (x AuthenticatorTransport) String() (string, bool) {
	switch x {
	case AuthenticatorTransport_USB:
		return "usb", true
	case AuthenticatorTransport_NFC:
		return "nfc", true
	case AuthenticatorTransport_BLE:
		return "ble", true
	case AuthenticatorTransport_SMART_CARD:
		return "smart-card", true
	case AuthenticatorTransport_HYBRID:
		return "hybrid", true
	case AuthenticatorTransport_INTERNAL:
		return "internal", true
	default:
		return "", false
	}
}

type AutoKeyword uint32

const (
	_ AutoKeyword = iota

	AutoKeyword_AUTO
)

func (AutoKeyword) FromRef(str js.Ref) AutoKeyword {
	return AutoKeyword(bindings.ConstOfAutoKeyword(str))
}

func (x AutoKeyword) String() (string, bool) {
	switch x {
	case AutoKeyword_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type AutoplayPolicy uint32

const (
	_ AutoplayPolicy = iota

	AutoplayPolicy_ALLOWED
	AutoplayPolicy_ALLOWED_MUTED
	AutoplayPolicy_DISALLOWED
)

func (AutoplayPolicy) FromRef(str js.Ref) AutoplayPolicy {
	return AutoplayPolicy(bindings.ConstOfAutoplayPolicy(str))
}

func (x AutoplayPolicy) String() (string, bool) {
	switch x {
	case AutoplayPolicy_ALLOWED:
		return "allowed", true
	case AutoplayPolicy_ALLOWED_MUTED:
		return "allowed-muted", true
	case AutoplayPolicy_DISALLOWED:
		return "disallowed", true
	default:
		return "", false
	}
}

type AutoplayPolicyMediaType uint32

const (
	_ AutoplayPolicyMediaType = iota

	AutoplayPolicyMediaType_MEDIAELEMENT
	AutoplayPolicyMediaType_AUDIOCONTEXT
)

func (AutoplayPolicyMediaType) FromRef(str js.Ref) AutoplayPolicyMediaType {
	return AutoplayPolicyMediaType(bindings.ConstOfAutoplayPolicyMediaType(str))
}

func (x AutoplayPolicyMediaType) String() (string, bool) {
	switch x {
	case AutoplayPolicyMediaType_MEDIAELEMENT:
		return "mediaelement", true
	case AutoplayPolicyMediaType_AUDIOCONTEXT:
		return "audiocontext", true
	default:
		return "", false
	}
}

type AvcBitstreamFormat uint32

const (
	_ AvcBitstreamFormat = iota

	AvcBitstreamFormat_ANNEXB
	AvcBitstreamFormat_AVC
)

func (AvcBitstreamFormat) FromRef(str js.Ref) AvcBitstreamFormat {
	return AvcBitstreamFormat(bindings.ConstOfAvcBitstreamFormat(str))
}

func (x AvcBitstreamFormat) String() (string, bool) {
	switch x {
	case AvcBitstreamFormat_ANNEXB:
		return "annexb", true
	case AvcBitstreamFormat_AVC:
		return "avc", true
	default:
		return "", false
	}
}

type AvcEncoderConfig struct {
	// Format is "AvcEncoderConfig.format"
	//
	// Optional, defaults to "avc".
	Format AvcBitstreamFormat

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AvcEncoderConfig with all fields set.
func (p AvcEncoderConfig) FromRef(ref js.Ref) AvcEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AvcEncoderConfig in the application heap.
func (p AvcEncoderConfig) New() js.Ref {
	return bindings.AvcEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AvcEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AvcEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AvcEncoderConfig) Update(ref js.Ref) {
	bindings.AvcEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Request_String struct {
	ref js.Ref
}

func (x OneOf_Request_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Request_String) Free() {
	x.ref.Free()
}

func (x OneOf_Request_String) FromRef(ref js.Ref) OneOf_Request_String {
	return OneOf_Request_String{
		ref: ref,
	}
}

func (x OneOf_Request_String) Request() Request {
	return Request{}.FromRef(x.ref)
}

func (x OneOf_Request_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type RequestInfo = OneOf_Request_String

type OneOf_ArrayArrayString_RecordString struct {
	ref js.Ref
}

func (x OneOf_ArrayArrayString_RecordString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayArrayString_RecordString) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayArrayString_RecordString) FromRef(ref js.Ref) OneOf_ArrayArrayString_RecordString {
	return OneOf_ArrayArrayString_RecordString{
		ref: ref,
	}
}

func (x OneOf_ArrayArrayString_RecordString) ArrayArrayString() js.Array[js.Array[js.String]] {
	return js.Array[js.Array[js.String]]{}.FromRef(x.ref)
}

func (x OneOf_ArrayArrayString_RecordString) RecordString() js.Record[js.String] {
	return js.Record[js.String]{}.FromRef(x.ref)
}

type HeadersInit = OneOf_ArrayArrayString_RecordString

type OneOf_ArrayArrayString_RecordString_String struct {
	ref js.Ref
}

func (x OneOf_ArrayArrayString_RecordString_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayArrayString_RecordString_String) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayArrayString_RecordString_String) FromRef(ref js.Ref) OneOf_ArrayArrayString_RecordString_String {
	return OneOf_ArrayArrayString_RecordString_String{
		ref: ref,
	}
}

func (x OneOf_ArrayArrayString_RecordString_String) ArrayArrayString() js.Array[js.Array[js.String]] {
	return js.Array[js.Array[js.String]]{}.FromRef(x.ref)
}

func (x OneOf_ArrayArrayString_RecordString_String) RecordString() js.Record[js.String] {
	return js.Record[js.String]{}.FromRef(x.ref)
}

func (x OneOf_ArrayArrayString_RecordString_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func NewURLSearchParams(init OneOf_ArrayArrayString_RecordString_String) (ret URLSearchParams) {
	ret.ref = bindings.NewURLSearchParamsByURLSearchParams(
		init.Ref())
	return
}

func NewURLSearchParamsByURLSearchParams1() (ret URLSearchParams) {
	ret.ref = bindings.NewURLSearchParamsByURLSearchParams1()
	return
}

type URLSearchParams struct {
	ref js.Ref
}

func (this URLSearchParams) Once() URLSearchParams {
	this.Ref().Once()
	return this
}

func (this URLSearchParams) Ref() js.Ref {
	return this.ref
}

func (this URLSearchParams) FromRef(ref js.Ref) URLSearchParams {
	this.ref = ref
	return this
}

func (this URLSearchParams) Free() {
	this.Ref().Free()
}

// Size returns the value of property "URLSearchParams.size".
//
// It returns ok=false if there is no such property.
func (this URLSearchParams) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetURLSearchParamsSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAppend returns true if the method "URLSearchParams.append" exists.
func (this URLSearchParams) HasAppend() bool {
	return js.True == bindings.HasURLSearchParamsAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "URLSearchParams.append".
func (this URLSearchParams) AppendFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.URLSearchParamsAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "URLSearchParams.append".
func (this URLSearchParams) Append(name js.String, value js.String) (ret js.Void) {
	bindings.CallURLSearchParamsAppend(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TryAppend calls the method "URLSearchParams.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryAppend(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasDelete returns true if the method "URLSearchParams.delete" exists.
func (this URLSearchParams) HasDelete() bool {
	return js.True == bindings.HasURLSearchParamsDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "URLSearchParams.delete".
func (this URLSearchParams) DeleteFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.URLSearchParamsDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "URLSearchParams.delete".
func (this URLSearchParams) Delete(name js.String, value js.String) (ret js.Void) {
	bindings.CallURLSearchParamsDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TryDelete calls the method "URLSearchParams.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryDelete(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasDelete1 returns true if the method "URLSearchParams.delete" exists.
func (this URLSearchParams) HasDelete1() bool {
	return js.True == bindings.HasURLSearchParamsDelete1(
		this.Ref(),
	)
}

// Delete1Func returns the method "URLSearchParams.delete".
func (this URLSearchParams) Delete1Func() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.URLSearchParamsDelete1Func(
			this.Ref(),
		),
	)
}

// Delete1 calls the method "URLSearchParams.delete".
func (this URLSearchParams) Delete1(name js.String) (ret js.Void) {
	bindings.CallURLSearchParamsDelete1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete1 calls the method "URLSearchParams.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryDelete1(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsDelete1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGet returns true if the method "URLSearchParams.get" exists.
func (this URLSearchParams) HasGet() bool {
	return js.True == bindings.HasURLSearchParamsGet(
		this.Ref(),
	)
}

// GetFunc returns the method "URLSearchParams.get".
func (this URLSearchParams) GetFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.URLSearchParamsGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "URLSearchParams.get".
func (this URLSearchParams) Get(name js.String) (ret js.String) {
	bindings.CallURLSearchParamsGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "URLSearchParams.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryGet(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetAll returns true if the method "URLSearchParams.getAll" exists.
func (this URLSearchParams) HasGetAll() bool {
	return js.True == bindings.HasURLSearchParamsGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "URLSearchParams.getAll".
func (this URLSearchParams) GetAllFunc() (fn js.Func[func(name js.String) js.Array[js.String]]) {
	return fn.FromRef(
		bindings.URLSearchParamsGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "URLSearchParams.getAll".
func (this URLSearchParams) GetAll(name js.String) (ret js.Array[js.String]) {
	bindings.CallURLSearchParamsGetAll(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetAll calls the method "URLSearchParams.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryGetAll(name js.String) (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasHas returns true if the method "URLSearchParams.has" exists.
func (this URLSearchParams) HasHas() bool {
	return js.True == bindings.HasURLSearchParamsHas(
		this.Ref(),
	)
}

// HasFunc returns the method "URLSearchParams.has".
func (this URLSearchParams) HasFunc() (fn js.Func[func(name js.String, value js.String) bool]) {
	return fn.FromRef(
		bindings.URLSearchParamsHasFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "URLSearchParams.has".
func (this URLSearchParams) Has(name js.String, value js.String) (ret bool) {
	bindings.CallURLSearchParamsHas(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TryHas calls the method "URLSearchParams.has"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryHas(name js.String, value js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsHas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasHas1 returns true if the method "URLSearchParams.has" exists.
func (this URLSearchParams) HasHas1() bool {
	return js.True == bindings.HasURLSearchParamsHas1(
		this.Ref(),
	)
}

// Has1Func returns the method "URLSearchParams.has".
func (this URLSearchParams) Has1Func() (fn js.Func[func(name js.String) bool]) {
	return fn.FromRef(
		bindings.URLSearchParamsHas1Func(
			this.Ref(),
		),
	)
}

// Has1 calls the method "URLSearchParams.has".
func (this URLSearchParams) Has1(name js.String) (ret bool) {
	bindings.CallURLSearchParamsHas1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryHas1 calls the method "URLSearchParams.has"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryHas1(name js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsHas1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSet returns true if the method "URLSearchParams.set" exists.
func (this URLSearchParams) HasSet() bool {
	return js.True == bindings.HasURLSearchParamsSet(
		this.Ref(),
	)
}

// SetFunc returns the method "URLSearchParams.set".
func (this URLSearchParams) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.URLSearchParamsSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "URLSearchParams.set".
func (this URLSearchParams) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallURLSearchParamsSet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "URLSearchParams.set"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TrySet(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasSort returns true if the method "URLSearchParams.sort" exists.
func (this URLSearchParams) HasSort() bool {
	return js.True == bindings.HasURLSearchParamsSort(
		this.Ref(),
	)
}

// SortFunc returns the method "URLSearchParams.sort".
func (this URLSearchParams) SortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.URLSearchParamsSortFunc(
			this.Ref(),
		),
	)
}

// Sort calls the method "URLSearchParams.sort".
func (this URLSearchParams) Sort() (ret js.Void) {
	bindings.CallURLSearchParamsSort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySort calls the method "URLSearchParams.sort"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TrySort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsSort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToString returns true if the method "URLSearchParams.toString" exists.
func (this URLSearchParams) HasToString() bool {
	return js.True == bindings.HasURLSearchParamsToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "URLSearchParams.toString".
func (this URLSearchParams) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.URLSearchParamsToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "URLSearchParams.toString".
func (this URLSearchParams) ToString() (ret js.String) {
	bindings.CallURLSearchParamsToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "URLSearchParams.toString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this URLSearchParams) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLSearchParamsToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String struct {
	ref js.Ref
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Free() {
	x.ref.Free()
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FromRef(ref js.Ref) OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String {
	return OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String{
		ref: ref,
	}
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) ReadableStream() ReadableStream {
	return ReadableStream{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) FormData() FormData {
	return FormData{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) URLSearchParams() URLSearchParams {
	return URLSearchParams{}.FromRef(x.ref)
}

func (x OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type BodyInit = OneOf_ReadableStream_Blob_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_FormData_URLSearchParams_String

type ReferrerPolicy uint32

const (
	_ ReferrerPolicy = iota

	ReferrerPolicy_
	ReferrerPolicy_NO_REFERRER
	ReferrerPolicy_NO_REFERRER_WHEN_DOWNGRADE
	ReferrerPolicy_SAME_ORIGIN
	ReferrerPolicy_ORIGIN
	ReferrerPolicy_STRICT_ORIGIN
	ReferrerPolicy_ORIGIN_WHEN_CROSS_ORIGIN
	ReferrerPolicy_STRICT_ORIGIN_WHEN_CROSS_ORIGIN
	ReferrerPolicy_UNSAFE_URL
)

func (ReferrerPolicy) FromRef(str js.Ref) ReferrerPolicy {
	return ReferrerPolicy(bindings.ConstOfReferrerPolicy(str))
}

func (x ReferrerPolicy) String() (string, bool) {
	switch x {
	case ReferrerPolicy_:
		return "", true
	case ReferrerPolicy_NO_REFERRER:
		return "no-referrer", true
	case ReferrerPolicy_NO_REFERRER_WHEN_DOWNGRADE:
		return "no-referrer-when-downgrade", true
	case ReferrerPolicy_SAME_ORIGIN:
		return "same-origin", true
	case ReferrerPolicy_ORIGIN:
		return "origin", true
	case ReferrerPolicy_STRICT_ORIGIN:
		return "strict-origin", true
	case ReferrerPolicy_ORIGIN_WHEN_CROSS_ORIGIN:
		return "origin-when-cross-origin", true
	case ReferrerPolicy_STRICT_ORIGIN_WHEN_CROSS_ORIGIN:
		return "strict-origin-when-cross-origin", true
	case ReferrerPolicy_UNSAFE_URL:
		return "unsafe-url", true
	default:
		return "", false
	}
}

type RequestMode uint32

const (
	_ RequestMode = iota

	RequestMode_NAVIGATE
	RequestMode_SAME_ORIGIN
	RequestMode_NO_CORS
	RequestMode_CORS
)

func (RequestMode) FromRef(str js.Ref) RequestMode {
	return RequestMode(bindings.ConstOfRequestMode(str))
}

func (x RequestMode) String() (string, bool) {
	switch x {
	case RequestMode_NAVIGATE:
		return "navigate", true
	case RequestMode_SAME_ORIGIN:
		return "same-origin", true
	case RequestMode_NO_CORS:
		return "no-cors", true
	case RequestMode_CORS:
		return "cors", true
	default:
		return "", false
	}
}

type RequestCredentials uint32

const (
	_ RequestCredentials = iota

	RequestCredentials_OMIT
	RequestCredentials_SAME_ORIGIN
	RequestCredentials_INCLUDE
)

func (RequestCredentials) FromRef(str js.Ref) RequestCredentials {
	return RequestCredentials(bindings.ConstOfRequestCredentials(str))
}

func (x RequestCredentials) String() (string, bool) {
	switch x {
	case RequestCredentials_OMIT:
		return "omit", true
	case RequestCredentials_SAME_ORIGIN:
		return "same-origin", true
	case RequestCredentials_INCLUDE:
		return "include", true
	default:
		return "", false
	}
}

type RequestCache uint32

const (
	_ RequestCache = iota

	RequestCache_DEFAULT
	RequestCache_NO_STORE
	RequestCache_RELOAD
	RequestCache_NO_CACHE
	RequestCache_FORCE_CACHE
	RequestCache_ONLY_IF_CACHED
)

func (RequestCache) FromRef(str js.Ref) RequestCache {
	return RequestCache(bindings.ConstOfRequestCache(str))
}

func (x RequestCache) String() (string, bool) {
	switch x {
	case RequestCache_DEFAULT:
		return "default", true
	case RequestCache_NO_STORE:
		return "no-store", true
	case RequestCache_RELOAD:
		return "reload", true
	case RequestCache_NO_CACHE:
		return "no-cache", true
	case RequestCache_FORCE_CACHE:
		return "force-cache", true
	case RequestCache_ONLY_IF_CACHED:
		return "only-if-cached", true
	default:
		return "", false
	}
}

type RequestRedirect uint32

const (
	_ RequestRedirect = iota

	RequestRedirect_FOLLOW
	RequestRedirect_ERROR
	RequestRedirect_MANUAL
)

func (RequestRedirect) FromRef(str js.Ref) RequestRedirect {
	return RequestRedirect(bindings.ConstOfRequestRedirect(str))
}

func (x RequestRedirect) String() (string, bool) {
	switch x {
	case RequestRedirect_FOLLOW:
		return "follow", true
	case RequestRedirect_ERROR:
		return "error", true
	case RequestRedirect_MANUAL:
		return "manual", true
	default:
		return "", false
	}
}

type RequestDuplex uint32

const (
	_ RequestDuplex = iota

	RequestDuplex_HALF
)

func (RequestDuplex) FromRef(str js.Ref) RequestDuplex {
	return RequestDuplex(bindings.ConstOfRequestDuplex(str))
}

func (x RequestDuplex) String() (string, bool) {
	switch x {
	case RequestDuplex_HALF:
		return "half", true
	default:
		return "", false
	}
}

type RequestPriority uint32

const (
	_ RequestPriority = iota

	RequestPriority_HIGH
	RequestPriority_LOW
	RequestPriority_AUTO
)

func (RequestPriority) FromRef(str js.Ref) RequestPriority {
	return RequestPriority(bindings.ConstOfRequestPriority(str))
}

func (x RequestPriority) String() (string, bool) {
	switch x {
	case RequestPriority_HIGH:
		return "high", true
	case RequestPriority_LOW:
		return "low", true
	case RequestPriority_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type TokenVersion uint32

const (
	_ TokenVersion = iota

	TokenVersion_1
)

func (TokenVersion) FromRef(str js.Ref) TokenVersion {
	return TokenVersion(bindings.ConstOfTokenVersion(str))
}

func (x TokenVersion) String() (string, bool) {
	switch x {
	case TokenVersion_1:
		return "1", true
	default:
		return "", false
	}
}

type OperationType uint32

const (
	_ OperationType = iota

	OperationType_TOKEN_REQUEST
	OperationType_SEND_REDEMPTION_RECORD
	OperationType_TOKEN_REDEMPTION
)

func (OperationType) FromRef(str js.Ref) OperationType {
	return OperationType(bindings.ConstOfOperationType(str))
}

func (x OperationType) String() (string, bool) {
	switch x {
	case OperationType_TOKEN_REQUEST:
		return "token-request", true
	case OperationType_SEND_REDEMPTION_RECORD:
		return "send-redemption-record", true
	case OperationType_TOKEN_REDEMPTION:
		return "token-redemption", true
	default:
		return "", false
	}
}

type RefreshPolicy uint32

const (
	_ RefreshPolicy = iota

	RefreshPolicy_NONE
	RefreshPolicy_REFRESH
)

func (RefreshPolicy) FromRef(str js.Ref) RefreshPolicy {
	return RefreshPolicy(bindings.ConstOfRefreshPolicy(str))
}

func (x RefreshPolicy) String() (string, bool) {
	switch x {
	case RefreshPolicy_NONE:
		return "none", true
	case RefreshPolicy_REFRESH:
		return "refresh", true
	default:
		return "", false
	}
}

type PrivateToken struct {
	// Version is "PrivateToken.version"
	//
	// Required
	Version TokenVersion
	// Operation is "PrivateToken.operation"
	//
	// Required
	Operation OperationType
	// RefreshPolicy is "PrivateToken.refreshPolicy"
	//
	// Optional, defaults to "none".
	RefreshPolicy RefreshPolicy
	// Issuers is "PrivateToken.issuers"
	//
	// Optional
	Issuers js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrivateToken with all fields set.
func (p PrivateToken) FromRef(ref js.Ref) PrivateToken {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PrivateToken in the application heap.
func (p PrivateToken) New() js.Ref {
	return bindings.PrivateTokenJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PrivateToken) UpdateFrom(ref js.Ref) {
	bindings.PrivateTokenJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PrivateToken) Update(ref js.Ref) {
	bindings.PrivateTokenJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RequestTargetAddressSpace uint32

const (
	_ RequestTargetAddressSpace = iota

	RequestTargetAddressSpace_PRIVATE
	RequestTargetAddressSpace_LOCAL
)

func (RequestTargetAddressSpace) FromRef(str js.Ref) RequestTargetAddressSpace {
	return RequestTargetAddressSpace(bindings.ConstOfRequestTargetAddressSpace(str))
}

func (x RequestTargetAddressSpace) String() (string, bool) {
	switch x {
	case RequestTargetAddressSpace_PRIVATE:
		return "private", true
	case RequestTargetAddressSpace_LOCAL:
		return "local", true
	default:
		return "", false
	}
}

type RequestInit struct {
	// Method is "RequestInit.method"
	//
	// Optional
	Method js.String
	// Headers is "RequestInit.headers"
	//
	// Optional
	Headers HeadersInit
	// Body is "RequestInit.body"
	//
	// Optional
	Body BodyInit
	// Referrer is "RequestInit.referrer"
	//
	// Optional
	Referrer js.String
	// ReferrerPolicy is "RequestInit.referrerPolicy"
	//
	// Optional
	ReferrerPolicy ReferrerPolicy
	// Mode is "RequestInit.mode"
	//
	// Optional
	Mode RequestMode
	// Credentials is "RequestInit.credentials"
	//
	// Optional
	Credentials RequestCredentials
	// Cache is "RequestInit.cache"
	//
	// Optional
	Cache RequestCache
	// Redirect is "RequestInit.redirect"
	//
	// Optional
	Redirect RequestRedirect
	// Integrity is "RequestInit.integrity"
	//
	// Optional
	Integrity js.String
	// Keepalive is "RequestInit.keepalive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Keepalive MUST be set to true to make this field effective.
	Keepalive bool
	// Signal is "RequestInit.signal"
	//
	// Optional
	Signal AbortSignal
	// Duplex is "RequestInit.duplex"
	//
	// Optional
	Duplex RequestDuplex
	// Priority is "RequestInit.priority"
	//
	// Optional
	Priority RequestPriority
	// Window is "RequestInit.window"
	//
	// Optional
	Window js.Any
	// PrivateToken is "RequestInit.privateToken"
	//
	// Optional
	//
	// NOTE: PrivateToken.FFI_USE MUST be set to true to get PrivateToken used.
	PrivateToken PrivateToken
	// TargetAddressSpace is "RequestInit.targetAddressSpace"
	//
	// Optional
	TargetAddressSpace RequestTargetAddressSpace
	// AttributionReporting is "RequestInit.attributionReporting"
	//
	// Optional
	//
	// NOTE: AttributionReporting.FFI_USE MUST be set to true to get AttributionReporting used.
	AttributionReporting AttributionReportingRequestOptions

	FFI_USE_Keepalive bool // for Keepalive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestInit with all fields set.
func (p RequestInit) FromRef(ref js.Ref) RequestInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestInit in the application heap.
func (p RequestInit) New() js.Ref {
	return bindings.RequestInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RequestInit) UpdateFrom(ref js.Ref) {
	bindings.RequestInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RequestInit) Update(ref js.Ref) {
	bindings.RequestInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewHeaders(init HeadersInit) (ret Headers) {
	ret.ref = bindings.NewHeadersByHeaders(
		init.Ref())
	return
}

func NewHeadersByHeaders1() (ret Headers) {
	ret.ref = bindings.NewHeadersByHeaders1()
	return
}

type Headers struct {
	ref js.Ref
}

func (this Headers) Once() Headers {
	this.Ref().Once()
	return this
}

func (this Headers) Ref() js.Ref {
	return this.ref
}

func (this Headers) FromRef(ref js.Ref) Headers {
	this.ref = ref
	return this
}

func (this Headers) Free() {
	this.Ref().Free()
}

// HasAppend returns true if the method "Headers.append" exists.
func (this Headers) HasAppend() bool {
	return js.True == bindings.HasHeadersAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "Headers.append".
func (this Headers) AppendFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.HeadersAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "Headers.append".
func (this Headers) Append(name js.String, value js.String) (ret js.Void) {
	bindings.CallHeadersAppend(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TryAppend calls the method "Headers.append"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TryAppend(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersAppend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

// HasDelete returns true if the method "Headers.delete" exists.
func (this Headers) HasDelete() bool {
	return js.True == bindings.HasHeadersDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "Headers.delete".
func (this Headers) DeleteFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.HeadersDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "Headers.delete".
func (this Headers) Delete(name js.String) (ret js.Void) {
	bindings.CallHeadersDelete(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDelete calls the method "Headers.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TryDelete(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGet returns true if the method "Headers.get" exists.
func (this Headers) HasGet() bool {
	return js.True == bindings.HasHeadersGet(
		this.Ref(),
	)
}

// GetFunc returns the method "Headers.get".
func (this Headers) GetFunc() (fn js.Func[func(name js.String) js.String]) {
	return fn.FromRef(
		bindings.HeadersGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Headers.get".
func (this Headers) Get(name js.String) (ret js.String) {
	bindings.CallHeadersGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "Headers.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TryGet(name js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetSetCookie returns true if the method "Headers.getSetCookie" exists.
func (this Headers) HasGetSetCookie() bool {
	return js.True == bindings.HasHeadersGetSetCookie(
		this.Ref(),
	)
}

// GetSetCookieFunc returns the method "Headers.getSetCookie".
func (this Headers) GetSetCookieFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.HeadersGetSetCookieFunc(
			this.Ref(),
		),
	)
}

// GetSetCookie calls the method "Headers.getSetCookie".
func (this Headers) GetSetCookie() (ret js.Array[js.String]) {
	bindings.CallHeadersGetSetCookie(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSetCookie calls the method "Headers.getSetCookie"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TryGetSetCookie() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersGetSetCookie(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHas returns true if the method "Headers.has" exists.
func (this Headers) HasHas() bool {
	return js.True == bindings.HasHeadersHas(
		this.Ref(),
	)
}

// HasFunc returns the method "Headers.has".
func (this Headers) HasFunc() (fn js.Func[func(name js.String) bool]) {
	return fn.FromRef(
		bindings.HeadersHasFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "Headers.has".
func (this Headers) Has(name js.String) (ret bool) {
	bindings.CallHeadersHas(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryHas calls the method "Headers.has"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TryHas(name js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersHas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasSet returns true if the method "Headers.set" exists.
func (this Headers) HasSet() bool {
	return js.True == bindings.HasHeadersSet(
		this.Ref(),
	)
}

// SetFunc returns the method "Headers.set".
func (this Headers) SetFunc() (fn js.Func[func(name js.String, value js.String)]) {
	return fn.FromRef(
		bindings.HeadersSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "Headers.set".
func (this Headers) Set(name js.String, value js.String) (ret js.Void) {
	bindings.CallHeadersSet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		value.Ref(),
	)

	return
}

// TrySet calls the method "Headers.set"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Headers) TrySet(name js.String, value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHeadersSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		value.Ref(),
	)

	return
}

type RequestDestination uint32

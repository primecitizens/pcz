// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package platformkeys

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/platformkeys/bindings"
)

type ClientCertificateType uint32

const (
	_ ClientCertificateType = iota

	ClientCertificateType_RSA_SIGN
	ClientCertificateType_ECDSA_SIGN
)

func (ClientCertificateType) FromRef(str js.Ref) ClientCertificateType {
	return ClientCertificateType(bindings.ConstOfClientCertificateType(str))
}

func (x ClientCertificateType) String() (string, bool) {
	switch x {
	case ClientCertificateType_RSA_SIGN:
		return "rsaSign", true
	case ClientCertificateType_ECDSA_SIGN:
		return "ecdsaSign", true
	default:
		return "", false
	}
}

type ClientCertificateRequest struct {
	// CertificateTypes is "ClientCertificateRequest.certificateTypes"
	//
	// Optional
	CertificateTypes js.Array[ClientCertificateType]
	// CertificateAuthorities is "ClientCertificateRequest.certificateAuthorities"
	//
	// Optional
	CertificateAuthorities js.Array[js.ArrayBuffer]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClientCertificateRequest with all fields set.
func (p ClientCertificateRequest) FromRef(ref js.Ref) ClientCertificateRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClientCertificateRequest in the application heap.
func (p ClientCertificateRequest) New() js.Ref {
	return bindings.ClientCertificateRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClientCertificateRequest) UpdateFrom(ref js.Ref) {
	bindings.ClientCertificateRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClientCertificateRequest) Update(ref js.Ref) {
	bindings.ClientCertificateRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClientCertificateRequest) FreeMembers(recursive bool) {
	js.Free(
		p.CertificateTypes.Ref(),
		p.CertificateAuthorities.Ref(),
	)
	p.CertificateTypes = p.CertificateTypes.FromRef(js.Undefined)
	p.CertificateAuthorities = p.CertificateAuthorities.FromRef(js.Undefined)
}

type GetKeyPairCallbackFunc func(this js.Ref, publicKey js.Object, privateKey js.Object) js.Ref

func (fn GetKeyPairCallbackFunc) Register() js.Func[func(publicKey js.Object, privateKey js.Object)] {
	return js.RegisterCallback[func(publicKey js.Object, privateKey js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetKeyPairCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
		js.Object{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetKeyPairCallback[T any] struct {
	Fn  func(arg T, this js.Ref, publicKey js.Object, privateKey js.Object) js.Ref
	Arg T
}

func (cb *GetKeyPairCallback[T]) Register() js.Func[func(publicKey js.Object, privateKey js.Object)] {
	return js.RegisterCallback[func(publicKey js.Object, privateKey js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetKeyPairCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
		js.Object{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Match struct {
	// Certificate is "Match.certificate"
	//
	// Optional
	Certificate js.ArrayBuffer
	// KeyAlgorithm is "Match.keyAlgorithm"
	//
	// Optional
	KeyAlgorithm js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Match with all fields set.
func (p Match) FromRef(ref js.Ref) Match {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Match in the application heap.
func (p Match) New() js.Ref {
	return bindings.MatchJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Match) UpdateFrom(ref js.Ref) {
	bindings.MatchJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Match) Update(ref js.Ref) {
	bindings.MatchJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Match) FreeMembers(recursive bool) {
	js.Free(
		p.Certificate.Ref(),
		p.KeyAlgorithm.Ref(),
	)
	p.Certificate = p.Certificate.FromRef(js.Undefined)
	p.KeyAlgorithm = p.KeyAlgorithm.FromRef(js.Undefined)
}

type SelectCallbackFunc func(this js.Ref, matches js.Array[Match]) js.Ref

func (fn SelectCallbackFunc) Register() js.Func[func(matches js.Array[Match])] {
	return js.RegisterCallback[func(matches js.Array[Match])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SelectCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Match]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SelectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, matches js.Array[Match]) js.Ref
	Arg T
}

func (cb *SelectCallback[T]) Register() js.Func[func(matches js.Array[Match])] {
	return js.RegisterCallback[func(matches js.Array[Match])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SelectCallback[T]) DispatchCallback(
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

		js.Array[Match]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SelectDetails struct {
	// Request is "SelectDetails.request"
	//
	// Optional
	//
	// NOTE: Request.FFI_USE MUST be set to true to get Request used.
	Request ClientCertificateRequest
	// ClientCerts is "SelectDetails.clientCerts"
	//
	// Optional
	ClientCerts js.Array[js.ArrayBuffer]
	// Interactive is "SelectDetails.interactive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Interactive MUST be set to true to make this field effective.
	Interactive bool

	FFI_USE_Interactive bool // for Interactive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SelectDetails with all fields set.
func (p SelectDetails) FromRef(ref js.Ref) SelectDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SelectDetails in the application heap.
func (p SelectDetails) New() js.Ref {
	return bindings.SelectDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SelectDetails) UpdateFrom(ref js.Ref) {
	bindings.SelectDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SelectDetails) Update(ref js.Ref) {
	bindings.SelectDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SelectDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ClientCerts.Ref(),
	)
	p.ClientCerts = p.ClientCerts.FromRef(js.Undefined)
	if recursive {
		p.Request.FreeMembers(true)
	}
}

type VerificationCallbackFunc func(this js.Ref, result *VerificationResult) js.Ref

func (fn VerificationCallbackFunc) Register() js.Func[func(result *VerificationResult)] {
	return js.RegisterCallback[func(result *VerificationResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VerificationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VerificationResult
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

type VerificationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *VerificationResult) js.Ref
	Arg T
}

func (cb *VerificationCallback[T]) Register() js.Func[func(result *VerificationResult)] {
	return js.RegisterCallback[func(result *VerificationResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VerificationCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VerificationResult
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

type VerificationResult struct {
	// Trusted is "VerificationResult.trusted"
	//
	// Optional
	//
	// NOTE: FFI_USE_Trusted MUST be set to true to make this field effective.
	Trusted bool
	// DebugErrors is "VerificationResult.debug_errors"
	//
	// Optional
	DebugErrors js.Array[js.String]

	FFI_USE_Trusted bool // for Trusted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VerificationResult with all fields set.
func (p VerificationResult) FromRef(ref js.Ref) VerificationResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VerificationResult in the application heap.
func (p VerificationResult) New() js.Ref {
	return bindings.VerificationResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VerificationResult) UpdateFrom(ref js.Ref) {
	bindings.VerificationResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VerificationResult) Update(ref js.Ref) {
	bindings.VerificationResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VerificationResult) FreeMembers(recursive bool) {
	js.Free(
		p.DebugErrors.Ref(),
	)
	p.DebugErrors = p.DebugErrors.FromRef(js.Undefined)
}

type VerificationDetails struct {
	// ServerCertificateChain is "VerificationDetails.serverCertificateChain"
	//
	// Optional
	ServerCertificateChain js.Array[js.ArrayBuffer]
	// Hostname is "VerificationDetails.hostname"
	//
	// Optional
	Hostname js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VerificationDetails with all fields set.
func (p VerificationDetails) FromRef(ref js.Ref) VerificationDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VerificationDetails in the application heap.
func (p VerificationDetails) New() js.Ref {
	return bindings.VerificationDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VerificationDetails) UpdateFrom(ref js.Ref) {
	bindings.VerificationDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VerificationDetails) Update(ref js.Ref) {
	bindings.VerificationDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VerificationDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ServerCertificateChain.Ref(),
		p.Hostname.Ref(),
	)
	p.ServerCertificateChain = p.ServerCertificateChain.FromRef(js.Undefined)
	p.Hostname = p.Hostname.FromRef(js.Undefined)
}

// HasFuncGetKeyPair returns true if the function "WEBEXT.platformKeys.getKeyPair" exists.
func HasFuncGetKeyPair() bool {
	return js.True == bindings.HasFuncGetKeyPair()
}

// FuncGetKeyPair returns the function "WEBEXT.platformKeys.getKeyPair".
func FuncGetKeyPair() (fn js.Func[func(certificate js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)])]) {
	bindings.FuncGetKeyPair(
		js.Pointer(&fn),
	)
	return
}

// GetKeyPair calls the function "WEBEXT.platformKeys.getKeyPair" directly.
func GetKeyPair(certificate js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)]) (ret js.Void) {
	bindings.CallGetKeyPair(
		js.Pointer(&ret),
		certificate.Ref(),
		parameters.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetKeyPair calls the function "WEBEXT.platformKeys.getKeyPair"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetKeyPair(certificate js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetKeyPair(
		js.Pointer(&ret), js.Pointer(&exception),
		certificate.Ref(),
		parameters.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetKeyPairBySpki returns true if the function "WEBEXT.platformKeys.getKeyPairBySpki" exists.
func HasFuncGetKeyPairBySpki() bool {
	return js.True == bindings.HasFuncGetKeyPairBySpki()
}

// FuncGetKeyPairBySpki returns the function "WEBEXT.platformKeys.getKeyPairBySpki".
func FuncGetKeyPairBySpki() (fn js.Func[func(publicKeySpkiDer js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)])]) {
	bindings.FuncGetKeyPairBySpki(
		js.Pointer(&fn),
	)
	return
}

// GetKeyPairBySpki calls the function "WEBEXT.platformKeys.getKeyPairBySpki" directly.
func GetKeyPairBySpki(publicKeySpkiDer js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)]) (ret js.Void) {
	bindings.CallGetKeyPairBySpki(
		js.Pointer(&ret),
		publicKeySpkiDer.Ref(),
		parameters.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetKeyPairBySpki calls the function "WEBEXT.platformKeys.getKeyPairBySpki"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetKeyPairBySpki(publicKeySpkiDer js.ArrayBuffer, parameters js.Object, callback js.Func[func(publicKey js.Object, privateKey js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetKeyPairBySpki(
		js.Pointer(&ret), js.Pointer(&exception),
		publicKeySpkiDer.Ref(),
		parameters.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncSelectClientCertificates returns true if the function "WEBEXT.platformKeys.selectClientCertificates" exists.
func HasFuncSelectClientCertificates() bool {
	return js.True == bindings.HasFuncSelectClientCertificates()
}

// FuncSelectClientCertificates returns the function "WEBEXT.platformKeys.selectClientCertificates".
func FuncSelectClientCertificates() (fn js.Func[func(details SelectDetails, callback js.Func[func(matches js.Array[Match])])]) {
	bindings.FuncSelectClientCertificates(
		js.Pointer(&fn),
	)
	return
}

// SelectClientCertificates calls the function "WEBEXT.platformKeys.selectClientCertificates" directly.
func SelectClientCertificates(details SelectDetails, callback js.Func[func(matches js.Array[Match])]) (ret js.Void) {
	bindings.CallSelectClientCertificates(
		js.Pointer(&ret),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TrySelectClientCertificates calls the function "WEBEXT.platformKeys.selectClientCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySelectClientCertificates(details SelectDetails, callback js.Func[func(matches js.Array[Match])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectClientCertificates(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// HasFuncSubtleCrypto returns true if the function "WEBEXT.platformKeys.subtleCrypto" exists.
func HasFuncSubtleCrypto() bool {
	return js.True == bindings.HasFuncSubtleCrypto()
}

// FuncSubtleCrypto returns the function "WEBEXT.platformKeys.subtleCrypto".
func FuncSubtleCrypto() (fn js.Func[func() js.Object]) {
	bindings.FuncSubtleCrypto(
		js.Pointer(&fn),
	)
	return
}

// SubtleCrypto calls the function "WEBEXT.platformKeys.subtleCrypto" directly.
func SubtleCrypto() (ret js.Object) {
	bindings.CallSubtleCrypto(
		js.Pointer(&ret),
	)

	return
}

// TrySubtleCrypto calls the function "WEBEXT.platformKeys.subtleCrypto"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySubtleCrypto() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySubtleCrypto(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncVerifyTLSServerCertificate returns true if the function "WEBEXT.platformKeys.verifyTLSServerCertificate" exists.
func HasFuncVerifyTLSServerCertificate() bool {
	return js.True == bindings.HasFuncVerifyTLSServerCertificate()
}

// FuncVerifyTLSServerCertificate returns the function "WEBEXT.platformKeys.verifyTLSServerCertificate".
func FuncVerifyTLSServerCertificate() (fn js.Func[func(details VerificationDetails, callback js.Func[func(result *VerificationResult)])]) {
	bindings.FuncVerifyTLSServerCertificate(
		js.Pointer(&fn),
	)
	return
}

// VerifyTLSServerCertificate calls the function "WEBEXT.platformKeys.verifyTLSServerCertificate" directly.
func VerifyTLSServerCertificate(details VerificationDetails, callback js.Func[func(result *VerificationResult)]) (ret js.Void) {
	bindings.CallVerifyTLSServerCertificate(
		js.Pointer(&ret),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TryVerifyTLSServerCertificate calls the function "WEBEXT.platformKeys.verifyTLSServerCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryVerifyTLSServerCertificate(details VerificationDetails, callback js.Func[func(result *VerificationResult)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVerifyTLSServerCertificate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

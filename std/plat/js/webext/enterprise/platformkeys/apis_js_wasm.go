// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package platformkeys

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/platformkeys/bindings"
)

type Algorithm uint32

const (
	_ Algorithm = iota

	Algorithm_RSA
	Algorithm_ECDSA
)

func (Algorithm) FromRef(str js.Ref) Algorithm {
	return Algorithm(bindings.ConstOfAlgorithm(str))
}

func (x Algorithm) String() (string, bool) {
	switch x {
	case Algorithm_RSA:
		return "RSA", true
	case Algorithm_ECDSA:
		return "ECDSA", true
	default:
		return "", false
	}
}

type ChallengeCallbackFunc func(this js.Ref, response js.ArrayBuffer) js.Ref

func (fn ChallengeCallbackFunc) Register() js.Func[func(response js.ArrayBuffer)] {
	return js.RegisterCallback[func(response js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ChallengeCallbackFunc) DispatchCallback(
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

type ChallengeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *ChallengeCallback[T]) Register() js.Func[func(response js.ArrayBuffer)] {
	return js.RegisterCallback[func(response js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ChallengeCallback[T]) DispatchCallback(
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

type RegisterKeyOptions struct {
	// Algorithm is "RegisterKeyOptions.algorithm"
	//
	// Optional
	Algorithm Algorithm

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegisterKeyOptions with all fields set.
func (p RegisterKeyOptions) FromRef(ref js.Ref) RegisterKeyOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RegisterKeyOptions in the application heap.
func (p RegisterKeyOptions) New() js.Ref {
	return bindings.RegisterKeyOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegisterKeyOptions) UpdateFrom(ref js.Ref) {
	bindings.RegisterKeyOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegisterKeyOptions) Update(ref js.Ref) {
	bindings.RegisterKeyOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegisterKeyOptions) FreeMembers(recursive bool) {
}

type Scope uint32

const (
	_ Scope = iota

	Scope_USER
	Scope_MACHINE
)

func (Scope) FromRef(str js.Ref) Scope {
	return Scope(bindings.ConstOfScope(str))
}

func (x Scope) String() (string, bool) {
	switch x {
	case Scope_USER:
		return "USER", true
	case Scope_MACHINE:
		return "MACHINE", true
	default:
		return "", false
	}
}

type ChallengeKeyOptions struct {
	// Challenge is "ChallengeKeyOptions.challenge"
	//
	// Optional
	Challenge js.ArrayBuffer
	// RegisterKey is "ChallengeKeyOptions.registerKey"
	//
	// Optional
	//
	// NOTE: RegisterKey.FFI_USE MUST be set to true to get RegisterKey used.
	RegisterKey RegisterKeyOptions
	// Scope is "ChallengeKeyOptions.scope"
	//
	// Optional
	Scope Scope

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChallengeKeyOptions with all fields set.
func (p ChallengeKeyOptions) FromRef(ref js.Ref) ChallengeKeyOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChallengeKeyOptions in the application heap.
func (p ChallengeKeyOptions) New() js.Ref {
	return bindings.ChallengeKeyOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChallengeKeyOptions) UpdateFrom(ref js.Ref) {
	bindings.ChallengeKeyOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChallengeKeyOptions) Update(ref js.Ref) {
	bindings.ChallengeKeyOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChallengeKeyOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	if recursive {
		p.RegisterKey.FreeMembers(true)
	}
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

type GetCertificatesCallbackFunc func(this js.Ref, certificates js.Array[js.ArrayBuffer]) js.Ref

func (fn GetCertificatesCallbackFunc) Register() js.Func[func(certificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(certificates js.Array[js.ArrayBuffer])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCertificatesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.ArrayBuffer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCertificatesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, certificates js.Array[js.ArrayBuffer]) js.Ref
	Arg T
}

func (cb *GetCertificatesCallback[T]) Register() js.Func[func(certificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(certificates js.Array[js.ArrayBuffer])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCertificatesCallback[T]) DispatchCallback(
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

		js.Array[js.ArrayBuffer]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetTokensCallbackFunc func(this js.Ref, tokens js.Array[Token]) js.Ref

func (fn GetTokensCallbackFunc) Register() js.Func[func(tokens js.Array[Token])] {
	return js.RegisterCallback[func(tokens js.Array[Token])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetTokensCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Token]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetTokensCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tokens js.Array[Token]) js.Ref
	Arg T
}

func (cb *GetTokensCallback[T]) Register() js.Func[func(tokens js.Array[Token])] {
	return js.RegisterCallback[func(tokens js.Array[Token])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetTokensCallback[T]) DispatchCallback(
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

		js.Array[Token]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Token struct {
	// Id is "Token.id"
	//
	// Optional
	Id js.String
	// SubtleCrypto is "Token.subtleCrypto"
	//
	// Optional
	SubtleCrypto js.Object
	// SoftwareBackedSubtleCrypto is "Token.softwareBackedSubtleCrypto"
	//
	// Optional
	SoftwareBackedSubtleCrypto js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Token with all fields set.
func (p Token) FromRef(ref js.Ref) Token {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Token in the application heap.
func (p Token) New() js.Ref {
	return bindings.TokenJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Token) UpdateFrom(ref js.Ref) {
	bindings.TokenJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Token) Update(ref js.Ref) {
	bindings.TokenJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Token) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.SubtleCrypto.Ref(),
		p.SoftwareBackedSubtleCrypto.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.SubtleCrypto = p.SubtleCrypto.FromRef(js.Undefined)
	p.SoftwareBackedSubtleCrypto = p.SoftwareBackedSubtleCrypto.FromRef(js.Undefined)
}

// HasFuncChallengeKey returns true if the function "WEBEXT.enterprise.platformKeys.challengeKey" exists.
func HasFuncChallengeKey() bool {
	return js.True == bindings.HasFuncChallengeKey()
}

// FuncChallengeKey returns the function "WEBEXT.enterprise.platformKeys.challengeKey".
func FuncChallengeKey() (fn js.Func[func(options ChallengeKeyOptions, callback js.Func[func(response js.ArrayBuffer)])]) {
	bindings.FuncChallengeKey(
		js.Pointer(&fn),
	)
	return
}

// ChallengeKey calls the function "WEBEXT.enterprise.platformKeys.challengeKey" directly.
func ChallengeKey(options ChallengeKeyOptions, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallChallengeKey(
		js.Pointer(&ret),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryChallengeKey calls the function "WEBEXT.enterprise.platformKeys.challengeKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChallengeKey(options ChallengeKeyOptions, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChallengeKey(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncChallengeMachineKey returns true if the function "WEBEXT.enterprise.platformKeys.challengeMachineKey" exists.
func HasFuncChallengeMachineKey() bool {
	return js.True == bindings.HasFuncChallengeMachineKey()
}

// FuncChallengeMachineKey returns the function "WEBEXT.enterprise.platformKeys.challengeMachineKey".
func FuncChallengeMachineKey() (fn js.Func[func(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)])]) {
	bindings.FuncChallengeMachineKey(
		js.Pointer(&fn),
	)
	return
}

// ChallengeMachineKey calls the function "WEBEXT.enterprise.platformKeys.challengeMachineKey" directly.
func ChallengeMachineKey(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallChallengeMachineKey(
		js.Pointer(&ret),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
		callback.Ref(),
	)

	return
}

// TryChallengeMachineKey calls the function "WEBEXT.enterprise.platformKeys.challengeMachineKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChallengeMachineKey(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChallengeMachineKey(
		js.Pointer(&ret), js.Pointer(&exception),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
		callback.Ref(),
	)

	return
}

// HasFuncChallengeUserKey returns true if the function "WEBEXT.enterprise.platformKeys.challengeUserKey" exists.
func HasFuncChallengeUserKey() bool {
	return js.True == bindings.HasFuncChallengeUserKey()
}

// FuncChallengeUserKey returns the function "WEBEXT.enterprise.platformKeys.challengeUserKey".
func FuncChallengeUserKey() (fn js.Func[func(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)])]) {
	bindings.FuncChallengeUserKey(
		js.Pointer(&fn),
	)
	return
}

// ChallengeUserKey calls the function "WEBEXT.enterprise.platformKeys.challengeUserKey" directly.
func ChallengeUserKey(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallChallengeUserKey(
		js.Pointer(&ret),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
		callback.Ref(),
	)

	return
}

// TryChallengeUserKey calls the function "WEBEXT.enterprise.platformKeys.challengeUserKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChallengeUserKey(challenge js.ArrayBuffer, registerKey bool, callback js.Func[func(response js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChallengeUserKey(
		js.Pointer(&ret), js.Pointer(&exception),
		challenge.Ref(),
		js.Bool(bool(registerKey)),
		callback.Ref(),
	)

	return
}

// HasFuncGetCertificates returns true if the function "WEBEXT.enterprise.platformKeys.getCertificates" exists.
func HasFuncGetCertificates() bool {
	return js.True == bindings.HasFuncGetCertificates()
}

// FuncGetCertificates returns the function "WEBEXT.enterprise.platformKeys.getCertificates".
func FuncGetCertificates() (fn js.Func[func(tokenId js.String, callback js.Func[func(certificates js.Array[js.ArrayBuffer])])]) {
	bindings.FuncGetCertificates(
		js.Pointer(&fn),
	)
	return
}

// GetCertificates calls the function "WEBEXT.enterprise.platformKeys.getCertificates" directly.
func GetCertificates(tokenId js.String, callback js.Func[func(certificates js.Array[js.ArrayBuffer])]) (ret js.Void) {
	bindings.CallGetCertificates(
		js.Pointer(&ret),
		tokenId.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetCertificates calls the function "WEBEXT.enterprise.platformKeys.getCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCertificates(tokenId js.String, callback js.Func[func(certificates js.Array[js.ArrayBuffer])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCertificates(
		js.Pointer(&ret), js.Pointer(&exception),
		tokenId.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetTokens returns true if the function "WEBEXT.enterprise.platformKeys.getTokens" exists.
func HasFuncGetTokens() bool {
	return js.True == bindings.HasFuncGetTokens()
}

// FuncGetTokens returns the function "WEBEXT.enterprise.platformKeys.getTokens".
func FuncGetTokens() (fn js.Func[func(callback js.Func[func(tokens js.Array[Token])])]) {
	bindings.FuncGetTokens(
		js.Pointer(&fn),
	)
	return
}

// GetTokens calls the function "WEBEXT.enterprise.platformKeys.getTokens" directly.
func GetTokens(callback js.Func[func(tokens js.Array[Token])]) (ret js.Void) {
	bindings.CallGetTokens(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetTokens calls the function "WEBEXT.enterprise.platformKeys.getTokens"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTokens(callback js.Func[func(tokens js.Array[Token])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTokens(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncImportCertificate returns true if the function "WEBEXT.enterprise.platformKeys.importCertificate" exists.
func HasFuncImportCertificate() bool {
	return js.True == bindings.HasFuncImportCertificate()
}

// FuncImportCertificate returns the function "WEBEXT.enterprise.platformKeys.importCertificate".
func FuncImportCertificate() (fn js.Func[func(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()])]) {
	bindings.FuncImportCertificate(
		js.Pointer(&fn),
	)
	return
}

// ImportCertificate calls the function "WEBEXT.enterprise.platformKeys.importCertificate" directly.
func ImportCertificate(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()]) (ret js.Void) {
	bindings.CallImportCertificate(
		js.Pointer(&ret),
		tokenId.Ref(),
		certificate.Ref(),
		callback.Ref(),
	)

	return
}

// TryImportCertificate calls the function "WEBEXT.enterprise.platformKeys.importCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImportCertificate(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImportCertificate(
		js.Pointer(&ret), js.Pointer(&exception),
		tokenId.Ref(),
		certificate.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveCertificate returns true if the function "WEBEXT.enterprise.platformKeys.removeCertificate" exists.
func HasFuncRemoveCertificate() bool {
	return js.True == bindings.HasFuncRemoveCertificate()
}

// FuncRemoveCertificate returns the function "WEBEXT.enterprise.platformKeys.removeCertificate".
func FuncRemoveCertificate() (fn js.Func[func(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()])]) {
	bindings.FuncRemoveCertificate(
		js.Pointer(&fn),
	)
	return
}

// RemoveCertificate calls the function "WEBEXT.enterprise.platformKeys.removeCertificate" directly.
func RemoveCertificate(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()]) (ret js.Void) {
	bindings.CallRemoveCertificate(
		js.Pointer(&ret),
		tokenId.Ref(),
		certificate.Ref(),
		callback.Ref(),
	)

	return
}

// TryRemoveCertificate calls the function "WEBEXT.enterprise.platformKeys.removeCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCertificate(tokenId js.String, certificate js.ArrayBuffer, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCertificate(
		js.Pointer(&ret), js.Pointer(&exception),
		tokenId.Ref(),
		certificate.Ref(),
		callback.Ref(),
	)

	return
}

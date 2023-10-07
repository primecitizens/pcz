// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package platformkeysinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/platformkeysinternal/bindings"
)

type Hash struct {
	// Name is "Hash.name"
	//
	// Optional
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Hash with all fields set.
func (p Hash) FromRef(ref js.Ref) Hash {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Hash in the application heap.
func (p Hash) New() js.Ref {
	return bindings.HashJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Hash) UpdateFrom(ref js.Ref) {
	bindings.HashJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Hash) Update(ref js.Ref) {
	bindings.HashJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Hash) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type Algorithm struct {
	// Name is "Algorithm.name"
	//
	// Optional
	Name js.String
	// ModulusLength is "Algorithm.modulusLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModulusLength MUST be set to true to make this field effective.
	ModulusLength int32
	// PublicExponent is "Algorithm.publicExponent"
	//
	// Optional
	PublicExponent js.ArrayBuffer
	// Hash is "Algorithm.hash"
	//
	// Optional
	//
	// NOTE: Hash.FFI_USE MUST be set to true to get Hash used.
	Hash Hash
	// NamedCurve is "Algorithm.namedCurve"
	//
	// Optional
	NamedCurve js.String

	FFI_USE_ModulusLength bool // for ModulusLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Algorithm with all fields set.
func (p Algorithm) FromRef(ref js.Ref) Algorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Algorithm in the application heap.
func (p Algorithm) New() js.Ref {
	return bindings.AlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Algorithm) UpdateFrom(ref js.Ref) {
	bindings.AlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Algorithm) Update(ref js.Ref) {
	bindings.AlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Algorithm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.PublicExponent.Ref(),
		p.NamedCurve.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PublicExponent = p.PublicExponent.FromRef(js.Undefined)
	p.NamedCurve = p.NamedCurve.FromRef(js.Undefined)
	if recursive {
		p.Hash.FreeMembers(true)
	}
}

type GenerateKeyCallbackFunc func(this js.Ref, publicKey js.ArrayBuffer) js.Ref

func (fn GenerateKeyCallbackFunc) Register() js.Func[func(publicKey js.ArrayBuffer)] {
	return js.RegisterCallback[func(publicKey js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GenerateKeyCallbackFunc) DispatchCallback(
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

type GenerateKeyCallback[T any] struct {
	Fn  func(arg T, this js.Ref, publicKey js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *GenerateKeyCallback[T]) Register() js.Func[func(publicKey js.ArrayBuffer)] {
	return js.RegisterCallback[func(publicKey js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GenerateKeyCallback[T]) DispatchCallback(
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

type GetTokensCallbackFunc func(this js.Ref, tokenIds js.Array[js.String]) js.Ref

func (fn GetTokensCallbackFunc) Register() js.Func[func(tokenIds js.Array[js.String])] {
	return js.RegisterCallback[func(tokenIds js.Array[js.String])](
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetTokensCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tokenIds js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetTokensCallback[T]) Register() js.Func[func(tokenIds js.Array[js.String])] {
	return js.RegisterCallback[func(tokenIds js.Array[js.String])](
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGenerateKey returns true if the function "WEBEXT.enterprise.platformKeysInternal.generateKey" exists.
func HasFuncGenerateKey() bool {
	return js.True == bindings.HasFuncGenerateKey()
}

// FuncGenerateKey returns the function "WEBEXT.enterprise.platformKeysInternal.generateKey".
func FuncGenerateKey() (fn js.Func[func(tokenId js.String, algorithm Algorithm, softwareBacked bool, callback js.Func[func(publicKey js.ArrayBuffer)])]) {
	bindings.FuncGenerateKey(
		js.Pointer(&fn),
	)
	return
}

// GenerateKey calls the function "WEBEXT.enterprise.platformKeysInternal.generateKey" directly.
func GenerateKey(tokenId js.String, algorithm Algorithm, softwareBacked bool, callback js.Func[func(publicKey js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallGenerateKey(
		js.Pointer(&ret),
		tokenId.Ref(),
		js.Pointer(&algorithm),
		js.Bool(bool(softwareBacked)),
		callback.Ref(),
	)

	return
}

// TryGenerateKey calls the function "WEBEXT.enterprise.platformKeysInternal.generateKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGenerateKey(tokenId js.String, algorithm Algorithm, softwareBacked bool, callback js.Func[func(publicKey js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGenerateKey(
		js.Pointer(&ret), js.Pointer(&exception),
		tokenId.Ref(),
		js.Pointer(&algorithm),
		js.Bool(bool(softwareBacked)),
		callback.Ref(),
	)

	return
}

// HasFuncGetTokens returns true if the function "WEBEXT.enterprise.platformKeysInternal.getTokens" exists.
func HasFuncGetTokens() bool {
	return js.True == bindings.HasFuncGetTokens()
}

// FuncGetTokens returns the function "WEBEXT.enterprise.platformKeysInternal.getTokens".
func FuncGetTokens() (fn js.Func[func(callback js.Func[func(tokenIds js.Array[js.String])])]) {
	bindings.FuncGetTokens(
		js.Pointer(&fn),
	)
	return
}

// GetTokens calls the function "WEBEXT.enterprise.platformKeysInternal.getTokens" directly.
func GetTokens(callback js.Func[func(tokenIds js.Array[js.String])]) (ret js.Void) {
	bindings.CallGetTokens(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetTokens calls the function "WEBEXT.enterprise.platformKeysInternal.getTokens"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTokens(callback js.Func[func(tokenIds js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTokens(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

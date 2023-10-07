// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package platformkeysinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/platformkeys"
	"github.com/primecitizens/pcz/std/plat/js/webext/platformkeysinternal/bindings"
)

type GetPublicKeyCallbackFunc func(this js.Ref, publicKey js.ArrayBuffer, algorithm js.Object) js.Ref

func (fn GetPublicKeyCallbackFunc) Register() js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)] {
	return js.RegisterCallback[func(publicKey js.ArrayBuffer, algorithm js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetPublicKeyCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
		js.Object{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetPublicKeyCallback[T any] struct {
	Fn  func(arg T, this js.Ref, publicKey js.ArrayBuffer, algorithm js.Object) js.Ref
	Arg T
}

func (cb *GetPublicKeyCallback[T]) Register() js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)] {
	return js.RegisterCallback[func(publicKey js.ArrayBuffer, algorithm js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetPublicKeyCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
		js.Object{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SelectCallbackFunc func(this js.Ref, certs js.Array[platformkeys.Match]) js.Ref

func (fn SelectCallbackFunc) Register() js.Func[func(certs js.Array[platformkeys.Match])] {
	return js.RegisterCallback[func(certs js.Array[platformkeys.Match])](
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

		js.Array[platformkeys.Match]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SelectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, certs js.Array[platformkeys.Match]) js.Ref
	Arg T
}

func (cb *SelectCallback[T]) Register() js.Func[func(certs js.Array[platformkeys.Match])] {
	return js.RegisterCallback[func(certs js.Array[platformkeys.Match])](
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

		js.Array[platformkeys.Match]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SignCallbackFunc func(this js.Ref, signature js.ArrayBuffer) js.Ref

func (fn SignCallbackFunc) Register() js.Func[func(signature js.ArrayBuffer)] {
	return js.RegisterCallback[func(signature js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SignCallbackFunc) DispatchCallback(
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

type SignCallback[T any] struct {
	Fn  func(arg T, this js.Ref, signature js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *SignCallback[T]) Register() js.Func[func(signature js.ArrayBuffer)] {
	return js.RegisterCallback[func(signature js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SignCallback[T]) DispatchCallback(
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

// HasFuncGetPublicKey returns true if the function "WEBEXT.platformKeysInternal.getPublicKey" exists.
func HasFuncGetPublicKey() bool {
	return js.True == bindings.HasFuncGetPublicKey()
}

// FuncGetPublicKey returns the function "WEBEXT.platformKeysInternal.getPublicKey".
func FuncGetPublicKey() (fn js.Func[func(certificate js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)])]) {
	bindings.FuncGetPublicKey(
		js.Pointer(&fn),
	)
	return
}

// GetPublicKey calls the function "WEBEXT.platformKeysInternal.getPublicKey" directly.
func GetPublicKey(certificate js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)]) (ret js.Void) {
	bindings.CallGetPublicKey(
		js.Pointer(&ret),
		certificate.Ref(),
		algorithmName.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetPublicKey calls the function "WEBEXT.platformKeysInternal.getPublicKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPublicKey(certificate js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPublicKey(
		js.Pointer(&ret), js.Pointer(&exception),
		certificate.Ref(),
		algorithmName.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetPublicKeyBySpki returns true if the function "WEBEXT.platformKeysInternal.getPublicKeyBySpki" exists.
func HasFuncGetPublicKeyBySpki() bool {
	return js.True == bindings.HasFuncGetPublicKeyBySpki()
}

// FuncGetPublicKeyBySpki returns the function "WEBEXT.platformKeysInternal.getPublicKeyBySpki".
func FuncGetPublicKeyBySpki() (fn js.Func[func(publicKeySpkiDer js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)])]) {
	bindings.FuncGetPublicKeyBySpki(
		js.Pointer(&fn),
	)
	return
}

// GetPublicKeyBySpki calls the function "WEBEXT.platformKeysInternal.getPublicKeyBySpki" directly.
func GetPublicKeyBySpki(publicKeySpkiDer js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)]) (ret js.Void) {
	bindings.CallGetPublicKeyBySpki(
		js.Pointer(&ret),
		publicKeySpkiDer.Ref(),
		algorithmName.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetPublicKeyBySpki calls the function "WEBEXT.platformKeysInternal.getPublicKeyBySpki"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPublicKeyBySpki(publicKeySpkiDer js.ArrayBuffer, algorithmName js.String, callback js.Func[func(publicKey js.ArrayBuffer, algorithm js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPublicKeyBySpki(
		js.Pointer(&ret), js.Pointer(&exception),
		publicKeySpkiDer.Ref(),
		algorithmName.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncSelectClientCertificates returns true if the function "WEBEXT.platformKeysInternal.selectClientCertificates" exists.
func HasFuncSelectClientCertificates() bool {
	return js.True == bindings.HasFuncSelectClientCertificates()
}

// FuncSelectClientCertificates returns the function "WEBEXT.platformKeysInternal.selectClientCertificates".
func FuncSelectClientCertificates() (fn js.Func[func(details platformkeys.SelectDetails, callback js.Func[func(certs js.Array[platformkeys.Match])])]) {
	bindings.FuncSelectClientCertificates(
		js.Pointer(&fn),
	)
	return
}

// SelectClientCertificates calls the function "WEBEXT.platformKeysInternal.selectClientCertificates" directly.
func SelectClientCertificates(details platformkeys.SelectDetails, callback js.Func[func(certs js.Array[platformkeys.Match])]) (ret js.Void) {
	bindings.CallSelectClientCertificates(
		js.Pointer(&ret),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TrySelectClientCertificates calls the function "WEBEXT.platformKeysInternal.selectClientCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySelectClientCertificates(details platformkeys.SelectDetails, callback js.Func[func(certs js.Array[platformkeys.Match])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySelectClientCertificates(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// HasFuncSign returns true if the function "WEBEXT.platformKeysInternal.sign" exists.
func HasFuncSign() bool {
	return js.True == bindings.HasFuncSign()
}

// FuncSign returns the function "WEBEXT.platformKeysInternal.sign".
func FuncSign() (fn js.Func[func(tokenId js.String, publicKey js.ArrayBuffer, algorithmName js.String, hashAlgorithmName js.String, data js.ArrayBuffer, callback js.Func[func(signature js.ArrayBuffer)])]) {
	bindings.FuncSign(
		js.Pointer(&fn),
	)
	return
}

// Sign calls the function "WEBEXT.platformKeysInternal.sign" directly.
func Sign(tokenId js.String, publicKey js.ArrayBuffer, algorithmName js.String, hashAlgorithmName js.String, data js.ArrayBuffer, callback js.Func[func(signature js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallSign(
		js.Pointer(&ret),
		tokenId.Ref(),
		publicKey.Ref(),
		algorithmName.Ref(),
		hashAlgorithmName.Ref(),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// TrySign calls the function "WEBEXT.platformKeysInternal.sign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySign(tokenId js.String, publicKey js.ArrayBuffer, algorithmName js.String, hashAlgorithmName js.String, data js.ArrayBuffer, callback js.Func[func(signature js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySign(
		js.Pointer(&ret), js.Pointer(&exception),
		tokenId.Ref(),
		publicKey.Ref(),
		algorithmName.Ref(),
		hashAlgorithmName.Ref(),
		data.Ref(),
		callback.Ref(),
	)

	return
}

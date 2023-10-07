// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package certificateprovider

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/certificateprovider/bindings"
)

type Algorithm uint32

const (
	_ Algorithm = iota

	Algorithm_RSASSA_PKCS1_V1_5_MD5_SHA1
	Algorithm_RSASSA_PKCS1_V1_5_SHA1
	Algorithm_RSASSA_PKCS1_V1_5_SHA256
	Algorithm_RSASSA_PKCS1_V1_5_SHA384
	Algorithm_RSASSA_PKCS1_V1_5_SHA512
	Algorithm_RSASSA_PSS_SHA256
	Algorithm_RSASSA_PSS_SHA384
	Algorithm_RSASSA_PSS_SHA512
)

func (Algorithm) FromRef(str js.Ref) Algorithm {
	return Algorithm(bindings.ConstOfAlgorithm(str))
}

func (x Algorithm) String() (string, bool) {
	switch x {
	case Algorithm_RSASSA_PKCS1_V1_5_MD5_SHA1:
		return "RSASSA_PKCS1_v1_5_MD5_SHA1", true
	case Algorithm_RSASSA_PKCS1_V1_5_SHA1:
		return "RSASSA_PKCS1_v1_5_SHA1", true
	case Algorithm_RSASSA_PKCS1_V1_5_SHA256:
		return "RSASSA_PKCS1_v1_5_SHA256", true
	case Algorithm_RSASSA_PKCS1_V1_5_SHA384:
		return "RSASSA_PKCS1_v1_5_SHA384", true
	case Algorithm_RSASSA_PKCS1_V1_5_SHA512:
		return "RSASSA_PKCS1_v1_5_SHA512", true
	case Algorithm_RSASSA_PSS_SHA256:
		return "RSASSA_PSS_SHA256", true
	case Algorithm_RSASSA_PSS_SHA384:
		return "RSASSA_PSS_SHA384", true
	case Algorithm_RSASSA_PSS_SHA512:
		return "RSASSA_PSS_SHA512", true
	default:
		return "", false
	}
}

type Hash uint32

const (
	_ Hash = iota

	Hash_MD5_SHA1
	Hash_SHA1
	Hash_SHA256
	Hash_SHA384
	Hash_SHA512
)

func (Hash) FromRef(str js.Ref) Hash {
	return Hash(bindings.ConstOfHash(str))
}

func (x Hash) String() (string, bool) {
	switch x {
	case Hash_MD5_SHA1:
		return "MD5_SHA1", true
	case Hash_SHA1:
		return "SHA1", true
	case Hash_SHA256:
		return "SHA256", true
	case Hash_SHA384:
		return "SHA384", true
	case Hash_SHA512:
		return "SHA512", true
	default:
		return "", false
	}
}

type CertificateInfo struct {
	// Certificate is "CertificateInfo.certificate"
	//
	// Optional
	Certificate js.ArrayBuffer
	// SupportedHashes is "CertificateInfo.supportedHashes"
	//
	// Optional
	SupportedHashes js.Array[Hash]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CertificateInfo with all fields set.
func (p CertificateInfo) FromRef(ref js.Ref) CertificateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CertificateInfo in the application heap.
func (p CertificateInfo) New() js.Ref {
	return bindings.CertificateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CertificateInfo) UpdateFrom(ref js.Ref) {
	bindings.CertificateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CertificateInfo) Update(ref js.Ref) {
	bindings.CertificateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CertificateInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Certificate.Ref(),
		p.SupportedHashes.Ref(),
	)
	p.Certificate = p.Certificate.FromRef(js.Undefined)
	p.SupportedHashes = p.SupportedHashes.FromRef(js.Undefined)
}

type CertificatesCallbackFunc func(this js.Ref, certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])]) js.Ref

func (fn CertificatesCallbackFunc) Register() js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])] {
	return js.RegisterCallback[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CertificatesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[CertificateInfo]{}.FromRef(args[0+1]),
		js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CertificatesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])]) js.Ref
	Arg T
}

func (cb *CertificatesCallback[T]) Register() js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])] {
	return js.RegisterCallback[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CertificatesCallback[T]) DispatchCallback(
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

		js.Array[CertificateInfo]{}.FromRef(args[0+1]),
		js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResultCallbackFunc func(this js.Ref, rejectedCertificates js.Array[js.ArrayBuffer]) js.Ref

func (fn ResultCallbackFunc) Register() js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(rejectedCertificates js.Array[js.ArrayBuffer])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResultCallbackFunc) DispatchCallback(
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

type ResultCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rejectedCertificates js.Array[js.ArrayBuffer]) js.Ref
	Arg T
}

func (cb *ResultCallback[T]) Register() js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])] {
	return js.RegisterCallback[func(rejectedCertificates js.Array[js.ArrayBuffer])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResultCallback[T]) DispatchCallback(
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

type CertificatesUpdateRequest struct {
	// CertificatesRequestId is "CertificatesUpdateRequest.certificatesRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_CertificatesRequestId MUST be set to true to make this field effective.
	CertificatesRequestId int32

	FFI_USE_CertificatesRequestId bool // for CertificatesRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CertificatesUpdateRequest with all fields set.
func (p CertificatesUpdateRequest) FromRef(ref js.Ref) CertificatesUpdateRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CertificatesUpdateRequest in the application heap.
func (p CertificatesUpdateRequest) New() js.Ref {
	return bindings.CertificatesUpdateRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CertificatesUpdateRequest) UpdateFrom(ref js.Ref) {
	bindings.CertificatesUpdateRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CertificatesUpdateRequest) Update(ref js.Ref) {
	bindings.CertificatesUpdateRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CertificatesUpdateRequest) FreeMembers(recursive bool) {
}

type ClientCertificateInfo struct {
	// CertificateChain is "ClientCertificateInfo.certificateChain"
	//
	// Optional
	CertificateChain js.Array[js.ArrayBuffer]
	// SupportedAlgorithms is "ClientCertificateInfo.supportedAlgorithms"
	//
	// Optional
	SupportedAlgorithms js.Array[Algorithm]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClientCertificateInfo with all fields set.
func (p ClientCertificateInfo) FromRef(ref js.Ref) ClientCertificateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClientCertificateInfo in the application heap.
func (p ClientCertificateInfo) New() js.Ref {
	return bindings.ClientCertificateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClientCertificateInfo) UpdateFrom(ref js.Ref) {
	bindings.ClientCertificateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClientCertificateInfo) Update(ref js.Ref) {
	bindings.ClientCertificateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClientCertificateInfo) FreeMembers(recursive bool) {
	js.Free(
		p.CertificateChain.Ref(),
		p.SupportedAlgorithms.Ref(),
	)
	p.CertificateChain = p.CertificateChain.FromRef(js.Undefined)
	p.SupportedAlgorithms = p.SupportedAlgorithms.FromRef(js.Undefined)
}

type Error uint32

const (
	_ Error = iota

	Error_GENERAL_ERROR
)

func (Error) FromRef(str js.Ref) Error {
	return Error(bindings.ConstOfError(str))
}

func (x Error) String() (string, bool) {
	switch x {
	case Error_GENERAL_ERROR:
		return "GENERAL_ERROR", true
	default:
		return "", false
	}
}

type PinRequestErrorType uint32

const (
	_ PinRequestErrorType = iota

	PinRequestErrorType_INVALID_PIN
	PinRequestErrorType_INVALID_PUK
	PinRequestErrorType_MAX_ATTEMPTS_EXCEEDED
	PinRequestErrorType_UNKNOWN_ERROR
)

func (PinRequestErrorType) FromRef(str js.Ref) PinRequestErrorType {
	return PinRequestErrorType(bindings.ConstOfPinRequestErrorType(str))
}

func (x PinRequestErrorType) String() (string, bool) {
	switch x {
	case PinRequestErrorType_INVALID_PIN:
		return "INVALID_PIN", true
	case PinRequestErrorType_INVALID_PUK:
		return "INVALID_PUK", true
	case PinRequestErrorType_MAX_ATTEMPTS_EXCEEDED:
		return "MAX_ATTEMPTS_EXCEEDED", true
	case PinRequestErrorType_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR", true
	default:
		return "", false
	}
}

type PinRequestType uint32

const (
	_ PinRequestType = iota

	PinRequestType_PIN
	PinRequestType_PUK
)

func (PinRequestType) FromRef(str js.Ref) PinRequestType {
	return PinRequestType(bindings.ConstOfPinRequestType(str))
}

func (x PinRequestType) String() (string, bool) {
	switch x {
	case PinRequestType_PIN:
		return "PIN", true
	case PinRequestType_PUK:
		return "PUK", true
	default:
		return "", false
	}
}

type PinResponseDetails struct {
	// UserInput is "PinResponseDetails.userInput"
	//
	// Optional
	UserInput js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PinResponseDetails with all fields set.
func (p PinResponseDetails) FromRef(ref js.Ref) PinResponseDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PinResponseDetails in the application heap.
func (p PinResponseDetails) New() js.Ref {
	return bindings.PinResponseDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PinResponseDetails) UpdateFrom(ref js.Ref) {
	bindings.PinResponseDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PinResponseDetails) Update(ref js.Ref) {
	bindings.PinResponseDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PinResponseDetails) FreeMembers(recursive bool) {
	js.Free(
		p.UserInput.Ref(),
	)
	p.UserInput = p.UserInput.FromRef(js.Undefined)
}

type ReportSignatureCallbackFunc func(this js.Ref) js.Ref

func (fn ReportSignatureCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReportSignatureCallbackFunc) DispatchCallback(
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

type ReportSignatureCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ReportSignatureCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReportSignatureCallback[T]) DispatchCallback(
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

type ReportSignatureDetails struct {
	// SignRequestId is "ReportSignatureDetails.signRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignRequestId MUST be set to true to make this field effective.
	SignRequestId int32
	// Error is "ReportSignatureDetails.error"
	//
	// Optional
	Error Error
	// Signature is "ReportSignatureDetails.signature"
	//
	// Optional
	Signature js.ArrayBuffer

	FFI_USE_SignRequestId bool // for SignRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReportSignatureDetails with all fields set.
func (p ReportSignatureDetails) FromRef(ref js.Ref) ReportSignatureDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReportSignatureDetails in the application heap.
func (p ReportSignatureDetails) New() js.Ref {
	return bindings.ReportSignatureDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReportSignatureDetails) UpdateFrom(ref js.Ref) {
	bindings.ReportSignatureDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReportSignatureDetails) Update(ref js.Ref) {
	bindings.ReportSignatureDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReportSignatureDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Signature.Ref(),
	)
	p.Signature = p.Signature.FromRef(js.Undefined)
}

type RequestPinCallbackFunc func(this js.Ref, details *PinResponseDetails) js.Ref

func (fn RequestPinCallbackFunc) Register() js.Func[func(details *PinResponseDetails)] {
	return js.RegisterCallback[func(details *PinResponseDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RequestPinCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PinResponseDetails
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

type RequestPinCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *PinResponseDetails) js.Ref
	Arg T
}

func (cb *RequestPinCallback[T]) Register() js.Func[func(details *PinResponseDetails)] {
	return js.RegisterCallback[func(details *PinResponseDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RequestPinCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PinResponseDetails
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

type RequestPinDetails struct {
	// SignRequestId is "RequestPinDetails.signRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignRequestId MUST be set to true to make this field effective.
	SignRequestId int32
	// RequestType is "RequestPinDetails.requestType"
	//
	// Optional
	RequestType PinRequestType
	// ErrorType is "RequestPinDetails.errorType"
	//
	// Optional
	ErrorType PinRequestErrorType
	// AttemptsLeft is "RequestPinDetails.attemptsLeft"
	//
	// Optional
	//
	// NOTE: FFI_USE_AttemptsLeft MUST be set to true to make this field effective.
	AttemptsLeft int32

	FFI_USE_SignRequestId bool // for SignRequestId.
	FFI_USE_AttemptsLeft  bool // for AttemptsLeft.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestPinDetails with all fields set.
func (p RequestPinDetails) FromRef(ref js.Ref) RequestPinDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestPinDetails in the application heap.
func (p RequestPinDetails) New() js.Ref {
	return bindings.RequestPinDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestPinDetails) UpdateFrom(ref js.Ref) {
	bindings.RequestPinDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestPinDetails) Update(ref js.Ref) {
	bindings.RequestPinDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestPinDetails) FreeMembers(recursive bool) {
}

type SetCertificatesCallbackFunc func(this js.Ref) js.Ref

func (fn SetCertificatesCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetCertificatesCallbackFunc) DispatchCallback(
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

type SetCertificatesCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetCertificatesCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetCertificatesCallback[T]) DispatchCallback(
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

type SetCertificatesDetails struct {
	// CertificatesRequestId is "SetCertificatesDetails.certificatesRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_CertificatesRequestId MUST be set to true to make this field effective.
	CertificatesRequestId int32
	// Error is "SetCertificatesDetails.error"
	//
	// Optional
	Error Error
	// ClientCertificates is "SetCertificatesDetails.clientCertificates"
	//
	// Optional
	ClientCertificates js.Array[ClientCertificateInfo]

	FFI_USE_CertificatesRequestId bool // for CertificatesRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetCertificatesDetails with all fields set.
func (p SetCertificatesDetails) FromRef(ref js.Ref) SetCertificatesDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetCertificatesDetails in the application heap.
func (p SetCertificatesDetails) New() js.Ref {
	return bindings.SetCertificatesDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetCertificatesDetails) UpdateFrom(ref js.Ref) {
	bindings.SetCertificatesDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetCertificatesDetails) Update(ref js.Ref) {
	bindings.SetCertificatesDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetCertificatesDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ClientCertificates.Ref(),
	)
	p.ClientCertificates = p.ClientCertificates.FromRef(js.Undefined)
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

type SignRequest struct {
	// SignRequestId is "SignRequest.signRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignRequestId MUST be set to true to make this field effective.
	SignRequestId int32
	// Digest is "SignRequest.digest"
	//
	// Optional
	Digest js.ArrayBuffer
	// Hash is "SignRequest.hash"
	//
	// Optional
	Hash Hash
	// Certificate is "SignRequest.certificate"
	//
	// Optional
	Certificate js.ArrayBuffer

	FFI_USE_SignRequestId bool // for SignRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SignRequest with all fields set.
func (p SignRequest) FromRef(ref js.Ref) SignRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SignRequest in the application heap.
func (p SignRequest) New() js.Ref {
	return bindings.SignRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SignRequest) UpdateFrom(ref js.Ref) {
	bindings.SignRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SignRequest) Update(ref js.Ref) {
	bindings.SignRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SignRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Digest.Ref(),
		p.Certificate.Ref(),
	)
	p.Digest = p.Digest.FromRef(js.Undefined)
	p.Certificate = p.Certificate.FromRef(js.Undefined)
}

type SignatureRequest struct {
	// SignRequestId is "SignatureRequest.signRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignRequestId MUST be set to true to make this field effective.
	SignRequestId int32
	// Input is "SignatureRequest.input"
	//
	// Optional
	Input js.ArrayBuffer
	// Algorithm is "SignatureRequest.algorithm"
	//
	// Optional
	Algorithm Algorithm
	// Certificate is "SignatureRequest.certificate"
	//
	// Optional
	Certificate js.ArrayBuffer

	FFI_USE_SignRequestId bool // for SignRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SignatureRequest with all fields set.
func (p SignatureRequest) FromRef(ref js.Ref) SignatureRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SignatureRequest in the application heap.
func (p SignatureRequest) New() js.Ref {
	return bindings.SignatureRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SignatureRequest) UpdateFrom(ref js.Ref) {
	bindings.SignatureRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SignatureRequest) Update(ref js.Ref) {
	bindings.SignatureRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SignatureRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Input.Ref(),
		p.Certificate.Ref(),
	)
	p.Input = p.Input.FromRef(js.Undefined)
	p.Certificate = p.Certificate.FromRef(js.Undefined)
}

type StopPinRequestCallbackFunc func(this js.Ref) js.Ref

func (fn StopPinRequestCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StopPinRequestCallbackFunc) DispatchCallback(
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

type StopPinRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *StopPinRequestCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StopPinRequestCallback[T]) DispatchCallback(
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

type StopPinRequestDetails struct {
	// SignRequestId is "StopPinRequestDetails.signRequestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignRequestId MUST be set to true to make this field effective.
	SignRequestId int32
	// ErrorType is "StopPinRequestDetails.errorType"
	//
	// Optional
	ErrorType PinRequestErrorType

	FFI_USE_SignRequestId bool // for SignRequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StopPinRequestDetails with all fields set.
func (p StopPinRequestDetails) FromRef(ref js.Ref) StopPinRequestDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StopPinRequestDetails in the application heap.
func (p StopPinRequestDetails) New() js.Ref {
	return bindings.StopPinRequestDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StopPinRequestDetails) UpdateFrom(ref js.Ref) {
	bindings.StopPinRequestDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StopPinRequestDetails) Update(ref js.Ref) {
	bindings.StopPinRequestDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StopPinRequestDetails) FreeMembers(recursive bool) {
}

type OnCertificatesRequestedEventCallbackFunc func(this js.Ref, reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])]) js.Ref

func (fn OnCertificatesRequestedEventCallbackFunc) Register() js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])] {
	return js.RegisterCallback[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCertificatesRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnCertificatesRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])]) js.Ref
	Arg T
}

func (cb *OnCertificatesRequestedEventCallback[T]) Register() js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])] {
	return js.RegisterCallback[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCertificatesRequestedEventCallback[T]) DispatchCallback(
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

		js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnCertificatesRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesRequested.addListener" exists.
func HasFuncOnCertificatesRequested() bool {
	return js.True == bindings.HasFuncOnCertificatesRequested()
}

// FuncOnCertificatesRequested returns the function "WEBEXT.certificateProvider.onCertificatesRequested.addListener".
func FuncOnCertificatesRequested() (fn js.Func[func(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])])]) {
	bindings.FuncOnCertificatesRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.addListener" directly.
func OnCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret js.Void) {
	bindings.CallOnCertificatesRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCertificatesRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCertificatesRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesRequested.removeListener" exists.
func HasFuncOffCertificatesRequested() bool {
	return js.True == bindings.HasFuncOffCertificatesRequested()
}

// FuncOffCertificatesRequested returns the function "WEBEXT.certificateProvider.onCertificatesRequested.removeListener".
func FuncOffCertificatesRequested() (fn js.Func[func(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])])]) {
	bindings.FuncOffCertificatesRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.removeListener" directly.
func OffCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret js.Void) {
	bindings.CallOffCertificatesRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCertificatesRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCertificatesRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesRequested.hasListener" exists.
func HasFuncHasOnCertificatesRequested() bool {
	return js.True == bindings.HasFuncHasOnCertificatesRequested()
}

// FuncHasOnCertificatesRequested returns the function "WEBEXT.certificateProvider.onCertificatesRequested.hasListener".
func FuncHasOnCertificatesRequested() (fn js.Func[func(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) bool]) {
	bindings.FuncHasOnCertificatesRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.hasListener" directly.
func HasOnCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret bool) {
	bindings.CallHasOnCertificatesRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCertificatesRequested calls the function "WEBEXT.certificateProvider.onCertificatesRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCertificatesRequested(callback js.Func[func(reportCallback js.Func[func(certificates js.Array[CertificateInfo], callback js.Func[func(rejectedCertificates js.Array[js.ArrayBuffer])])])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCertificatesRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCertificatesUpdateRequestedEventCallbackFunc func(this js.Ref, request *CertificatesUpdateRequest) js.Ref

func (fn OnCertificatesUpdateRequestedEventCallbackFunc) Register() js.Func[func(request *CertificatesUpdateRequest)] {
	return js.RegisterCallback[func(request *CertificatesUpdateRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCertificatesUpdateRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CertificatesUpdateRequest
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

type OnCertificatesUpdateRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *CertificatesUpdateRequest) js.Ref
	Arg T
}

func (cb *OnCertificatesUpdateRequestedEventCallback[T]) Register() js.Func[func(request *CertificatesUpdateRequest)] {
	return js.RegisterCallback[func(request *CertificatesUpdateRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCertificatesUpdateRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CertificatesUpdateRequest
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

// HasFuncOnCertificatesUpdateRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener" exists.
func HasFuncOnCertificatesUpdateRequested() bool {
	return js.True == bindings.HasFuncOnCertificatesUpdateRequested()
}

// FuncOnCertificatesUpdateRequested returns the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener".
func FuncOnCertificatesUpdateRequested() (fn js.Func[func(callback js.Func[func(request *CertificatesUpdateRequest)])]) {
	bindings.FuncOnCertificatesUpdateRequested(
		js.Pointer(&fn),
	)
	return
}

// OnCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener" directly.
func OnCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret js.Void) {
	bindings.CallOnCertificatesUpdateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCertificatesUpdateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCertificatesUpdateRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener" exists.
func HasFuncOffCertificatesUpdateRequested() bool {
	return js.True == bindings.HasFuncOffCertificatesUpdateRequested()
}

// FuncOffCertificatesUpdateRequested returns the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener".
func FuncOffCertificatesUpdateRequested() (fn js.Func[func(callback js.Func[func(request *CertificatesUpdateRequest)])]) {
	bindings.FuncOffCertificatesUpdateRequested(
		js.Pointer(&fn),
	)
	return
}

// OffCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener" directly.
func OffCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret js.Void) {
	bindings.CallOffCertificatesUpdateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCertificatesUpdateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCertificatesUpdateRequested returns true if the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener" exists.
func HasFuncHasOnCertificatesUpdateRequested() bool {
	return js.True == bindings.HasFuncHasOnCertificatesUpdateRequested()
}

// FuncHasOnCertificatesUpdateRequested returns the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener".
func FuncHasOnCertificatesUpdateRequested() (fn js.Func[func(callback js.Func[func(request *CertificatesUpdateRequest)]) bool]) {
	bindings.FuncHasOnCertificatesUpdateRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener" directly.
func HasOnCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret bool) {
	bindings.CallHasOnCertificatesUpdateRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCertificatesUpdateRequested calls the function "WEBEXT.certificateProvider.onCertificatesUpdateRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCertificatesUpdateRequested(callback js.Func[func(request *CertificatesUpdateRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCertificatesUpdateRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSignDigestRequestedEventCallbackFunc func(this js.Ref, request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)]) js.Ref

func (fn OnSignDigestRequestedEventCallbackFunc) Register() js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])] {
	return js.RegisterCallback[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSignDigestRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SignRequest
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(signature js.ArrayBuffer)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSignDigestRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)]) js.Ref
	Arg T
}

func (cb *OnSignDigestRequestedEventCallback[T]) Register() js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])] {
	return js.RegisterCallback[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSignDigestRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SignRequest
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(signature js.ArrayBuffer)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSignDigestRequested returns true if the function "WEBEXT.certificateProvider.onSignDigestRequested.addListener" exists.
func HasFuncOnSignDigestRequested() bool {
	return js.True == bindings.HasFuncOnSignDigestRequested()
}

// FuncOnSignDigestRequested returns the function "WEBEXT.certificateProvider.onSignDigestRequested.addListener".
func FuncOnSignDigestRequested() (fn js.Func[func(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])])]) {
	bindings.FuncOnSignDigestRequested(
		js.Pointer(&fn),
	)
	return
}

// OnSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.addListener" directly.
func OnSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret js.Void) {
	bindings.CallOnSignDigestRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSignDigestRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSignDigestRequested returns true if the function "WEBEXT.certificateProvider.onSignDigestRequested.removeListener" exists.
func HasFuncOffSignDigestRequested() bool {
	return js.True == bindings.HasFuncOffSignDigestRequested()
}

// FuncOffSignDigestRequested returns the function "WEBEXT.certificateProvider.onSignDigestRequested.removeListener".
func FuncOffSignDigestRequested() (fn js.Func[func(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])])]) {
	bindings.FuncOffSignDigestRequested(
		js.Pointer(&fn),
	)
	return
}

// OffSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.removeListener" directly.
func OffSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret js.Void) {
	bindings.CallOffSignDigestRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSignDigestRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSignDigestRequested returns true if the function "WEBEXT.certificateProvider.onSignDigestRequested.hasListener" exists.
func HasFuncHasOnSignDigestRequested() bool {
	return js.True == bindings.HasFuncHasOnSignDigestRequested()
}

// FuncHasOnSignDigestRequested returns the function "WEBEXT.certificateProvider.onSignDigestRequested.hasListener".
func FuncHasOnSignDigestRequested() (fn js.Func[func(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) bool]) {
	bindings.FuncHasOnSignDigestRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.hasListener" directly.
func HasOnSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret bool) {
	bindings.CallHasOnSignDigestRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSignDigestRequested calls the function "WEBEXT.certificateProvider.onSignDigestRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSignDigestRequested(callback js.Func[func(request *SignRequest, reportCallback js.Func[func(signature js.ArrayBuffer)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSignDigestRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSignatureRequestedEventCallbackFunc func(this js.Ref, request *SignatureRequest) js.Ref

func (fn OnSignatureRequestedEventCallbackFunc) Register() js.Func[func(request *SignatureRequest)] {
	return js.RegisterCallback[func(request *SignatureRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSignatureRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SignatureRequest
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

type OnSignatureRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request *SignatureRequest) js.Ref
	Arg T
}

func (cb *OnSignatureRequestedEventCallback[T]) Register() js.Func[func(request *SignatureRequest)] {
	return js.RegisterCallback[func(request *SignatureRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSignatureRequestedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SignatureRequest
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

// HasFuncOnSignatureRequested returns true if the function "WEBEXT.certificateProvider.onSignatureRequested.addListener" exists.
func HasFuncOnSignatureRequested() bool {
	return js.True == bindings.HasFuncOnSignatureRequested()
}

// FuncOnSignatureRequested returns the function "WEBEXT.certificateProvider.onSignatureRequested.addListener".
func FuncOnSignatureRequested() (fn js.Func[func(callback js.Func[func(request *SignatureRequest)])]) {
	bindings.FuncOnSignatureRequested(
		js.Pointer(&fn),
	)
	return
}

// OnSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.addListener" directly.
func OnSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret js.Void) {
	bindings.CallOnSignatureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSignatureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSignatureRequested returns true if the function "WEBEXT.certificateProvider.onSignatureRequested.removeListener" exists.
func HasFuncOffSignatureRequested() bool {
	return js.True == bindings.HasFuncOffSignatureRequested()
}

// FuncOffSignatureRequested returns the function "WEBEXT.certificateProvider.onSignatureRequested.removeListener".
func FuncOffSignatureRequested() (fn js.Func[func(callback js.Func[func(request *SignatureRequest)])]) {
	bindings.FuncOffSignatureRequested(
		js.Pointer(&fn),
	)
	return
}

// OffSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.removeListener" directly.
func OffSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret js.Void) {
	bindings.CallOffSignatureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSignatureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSignatureRequested returns true if the function "WEBEXT.certificateProvider.onSignatureRequested.hasListener" exists.
func HasFuncHasOnSignatureRequested() bool {
	return js.True == bindings.HasFuncHasOnSignatureRequested()
}

// FuncHasOnSignatureRequested returns the function "WEBEXT.certificateProvider.onSignatureRequested.hasListener".
func FuncHasOnSignatureRequested() (fn js.Func[func(callback js.Func[func(request *SignatureRequest)]) bool]) {
	bindings.FuncHasOnSignatureRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.hasListener" directly.
func HasOnSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret bool) {
	bindings.CallHasOnSignatureRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSignatureRequested calls the function "WEBEXT.certificateProvider.onSignatureRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSignatureRequested(callback js.Func[func(request *SignatureRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSignatureRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncReportSignature returns true if the function "WEBEXT.certificateProvider.reportSignature" exists.
func HasFuncReportSignature() bool {
	return js.True == bindings.HasFuncReportSignature()
}

// FuncReportSignature returns the function "WEBEXT.certificateProvider.reportSignature".
func FuncReportSignature() (fn js.Func[func(details ReportSignatureDetails) js.Promise[js.Void]]) {
	bindings.FuncReportSignature(
		js.Pointer(&fn),
	)
	return
}

// ReportSignature calls the function "WEBEXT.certificateProvider.reportSignature" directly.
func ReportSignature(details ReportSignatureDetails) (ret js.Promise[js.Void]) {
	bindings.CallReportSignature(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryReportSignature calls the function "WEBEXT.certificateProvider.reportSignature"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReportSignature(details ReportSignatureDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReportSignature(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncRequestPin returns true if the function "WEBEXT.certificateProvider.requestPin" exists.
func HasFuncRequestPin() bool {
	return js.True == bindings.HasFuncRequestPin()
}

// FuncRequestPin returns the function "WEBEXT.certificateProvider.requestPin".
func FuncRequestPin() (fn js.Func[func(details RequestPinDetails) js.Promise[PinResponseDetails]]) {
	bindings.FuncRequestPin(
		js.Pointer(&fn),
	)
	return
}

// RequestPin calls the function "WEBEXT.certificateProvider.requestPin" directly.
func RequestPin(details RequestPinDetails) (ret js.Promise[PinResponseDetails]) {
	bindings.CallRequestPin(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryRequestPin calls the function "WEBEXT.certificateProvider.requestPin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestPin(details RequestPinDetails) (ret js.Promise[PinResponseDetails], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestPin(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetCertificates returns true if the function "WEBEXT.certificateProvider.setCertificates" exists.
func HasFuncSetCertificates() bool {
	return js.True == bindings.HasFuncSetCertificates()
}

// FuncSetCertificates returns the function "WEBEXT.certificateProvider.setCertificates".
func FuncSetCertificates() (fn js.Func[func(details SetCertificatesDetails) js.Promise[js.Void]]) {
	bindings.FuncSetCertificates(
		js.Pointer(&fn),
	)
	return
}

// SetCertificates calls the function "WEBEXT.certificateProvider.setCertificates" directly.
func SetCertificates(details SetCertificatesDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetCertificates(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetCertificates calls the function "WEBEXT.certificateProvider.setCertificates"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCertificates(details SetCertificatesDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCertificates(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncStopPinRequest returns true if the function "WEBEXT.certificateProvider.stopPinRequest" exists.
func HasFuncStopPinRequest() bool {
	return js.True == bindings.HasFuncStopPinRequest()
}

// FuncStopPinRequest returns the function "WEBEXT.certificateProvider.stopPinRequest".
func FuncStopPinRequest() (fn js.Func[func(details StopPinRequestDetails) js.Promise[js.Void]]) {
	bindings.FuncStopPinRequest(
		js.Pointer(&fn),
	)
	return
}

// StopPinRequest calls the function "WEBEXT.certificateProvider.stopPinRequest" directly.
func StopPinRequest(details StopPinRequestDetails) (ret js.Promise[js.Void]) {
	bindings.CallStopPinRequest(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryStopPinRequest calls the function "WEBEXT.certificateProvider.stopPinRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopPinRequest(details StopPinRequestDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopPinRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

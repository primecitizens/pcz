// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webrequest

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrequest/bindings"
)

type BlockingResponseFieldAuthCredentials struct {
	// Password is "BlockingResponseFieldAuthCredentials.password"
	//
	// Required
	Password js.String
	// Username is "BlockingResponseFieldAuthCredentials.username"
	//
	// Required
	Username js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BlockingResponseFieldAuthCredentials with all fields set.
func (p BlockingResponseFieldAuthCredentials) FromRef(ref js.Ref) BlockingResponseFieldAuthCredentials {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BlockingResponseFieldAuthCredentials in the application heap.
func (p BlockingResponseFieldAuthCredentials) New() js.Ref {
	return bindings.BlockingResponseFieldAuthCredentialsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BlockingResponseFieldAuthCredentials) UpdateFrom(ref js.Ref) {
	bindings.BlockingResponseFieldAuthCredentialsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BlockingResponseFieldAuthCredentials) Update(ref js.Ref) {
	bindings.BlockingResponseFieldAuthCredentialsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BlockingResponseFieldAuthCredentials) FreeMembers(recursive bool) {
	js.Free(
		p.Password.Ref(),
		p.Username.Ref(),
	)
	p.Password = p.Password.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
}

type HttpHeadersElem struct {
	// BinaryValue is "HttpHeadersElem.binaryValue"
	//
	// Optional
	BinaryValue js.Array[int64]
	// Name is "HttpHeadersElem.name"
	//
	// Required
	Name js.String
	// Value is "HttpHeadersElem.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HttpHeadersElem with all fields set.
func (p HttpHeadersElem) FromRef(ref js.Ref) HttpHeadersElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HttpHeadersElem in the application heap.
func (p HttpHeadersElem) New() js.Ref {
	return bindings.HttpHeadersElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HttpHeadersElem) UpdateFrom(ref js.Ref) {
	bindings.HttpHeadersElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HttpHeadersElem) Update(ref js.Ref) {
	bindings.HttpHeadersElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HttpHeadersElem) FreeMembers(recursive bool) {
	js.Free(
		p.BinaryValue.Ref(),
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.BinaryValue = p.BinaryValue.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type HttpHeaders = js.Array[HttpHeadersElem]

type BlockingResponse struct {
	// AuthCredentials is "BlockingResponse.authCredentials"
	//
	// Optional
	//
	// NOTE: AuthCredentials.FFI_USE MUST be set to true to get AuthCredentials used.
	AuthCredentials BlockingResponseFieldAuthCredentials
	// Cancel is "BlockingResponse.cancel"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cancel MUST be set to true to make this field effective.
	Cancel bool
	// RedirectUrl is "BlockingResponse.redirectUrl"
	//
	// Optional
	RedirectUrl js.String
	// RequestHeaders is "BlockingResponse.requestHeaders"
	//
	// Optional
	RequestHeaders HttpHeaders
	// ResponseHeaders is "BlockingResponse.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders

	FFI_USE_Cancel bool // for Cancel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BlockingResponse with all fields set.
func (p BlockingResponse) FromRef(ref js.Ref) BlockingResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BlockingResponse in the application heap.
func (p BlockingResponse) New() js.Ref {
	return bindings.BlockingResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BlockingResponse) UpdateFrom(ref js.Ref) {
	bindings.BlockingResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BlockingResponse) Update(ref js.Ref) {
	bindings.BlockingResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BlockingResponse) FreeMembers(recursive bool) {
	js.Free(
		p.RedirectUrl.Ref(),
		p.RequestHeaders.Ref(),
		p.ResponseHeaders.Ref(),
	)
	p.RedirectUrl = p.RedirectUrl.FromRef(js.Undefined)
	p.RequestHeaders = p.RequestHeaders.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	if recursive {
		p.AuthCredentials.FreeMembers(true)
	}
}

type OneOf_TypedArrayUint8_String struct {
	ref js.Ref
}

func (x OneOf_TypedArrayUint8_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayUint8_String) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayUint8_String) FromRef(ref js.Ref) OneOf_TypedArrayUint8_String {
	return OneOf_TypedArrayUint8_String{
		ref: ref,
	}
}

func (x OneOf_TypedArrayUint8_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayUint8_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type FormDataItem = OneOf_TypedArrayUint8_String

type IgnoredActionType uint32

const (
	_ IgnoredActionType = iota

	IgnoredActionType_REDIRECT
	IgnoredActionType_REQUEST_HEADERS
	IgnoredActionType_RESPONSE_HEADERS
	IgnoredActionType_AUTH_CREDENTIALS
)

func (IgnoredActionType) FromRef(str js.Ref) IgnoredActionType {
	return IgnoredActionType(bindings.ConstOfIgnoredActionType(str))
}

func (x IgnoredActionType) String() (string, bool) {
	switch x {
	case IgnoredActionType_REDIRECT:
		return "redirect", true
	case IgnoredActionType_REQUEST_HEADERS:
		return "request_headers", true
	case IgnoredActionType_RESPONSE_HEADERS:
		return "response_headers", true
	case IgnoredActionType_AUTH_CREDENTIALS:
		return "auth_credentials", true
	default:
		return "", false
	}
}

// MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES returns the value of property "WEBEXT.webRequest.MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES".
//
// The returned bool will be false if there is no such property.
func MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES sets the value of property "WEBEXT.webRequest.MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES" to val.
//
// It returns false if the property cannot be set.
func SetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES(val js.String) bool {
	return js.True == bindings.SetMAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES(
		val.Ref())
}

type OnActionIgnoredArgDetails struct {
	// Action is "OnActionIgnoredArgDetails.action"
	//
	// Required
	Action IgnoredActionType
	// RequestId is "OnActionIgnoredArgDetails.requestId"
	//
	// Required
	RequestId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnActionIgnoredArgDetails with all fields set.
func (p OnActionIgnoredArgDetails) FromRef(ref js.Ref) OnActionIgnoredArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnActionIgnoredArgDetails in the application heap.
func (p OnActionIgnoredArgDetails) New() js.Ref {
	return bindings.OnActionIgnoredArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnActionIgnoredArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnActionIgnoredArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnActionIgnoredArgDetails) Update(ref js.Ref) {
	bindings.OnActionIgnoredArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnActionIgnoredArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.RequestId.Ref(),
	)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
}

type OnAuthRequiredArgAsyncCallbackFunc func(this js.Ref, response *BlockingResponse) js.Ref

func (fn OnAuthRequiredArgAsyncCallbackFunc) Register() js.Func[func(response *BlockingResponse)] {
	return js.RegisterCallback[func(response *BlockingResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAuthRequiredArgAsyncCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BlockingResponse
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

type OnAuthRequiredArgAsyncCallback[T any] struct {
	Fn  func(arg T, this js.Ref, response *BlockingResponse) js.Ref
	Arg T
}

func (cb *OnAuthRequiredArgAsyncCallback[T]) Register() js.Func[func(response *BlockingResponse)] {
	return js.RegisterCallback[func(response *BlockingResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAuthRequiredArgAsyncCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BlockingResponse
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

type OnAuthRequiredArgDetailsFieldChallenger struct {
	// Host is "OnAuthRequiredArgDetailsFieldChallenger.host"
	//
	// Required
	Host js.String
	// Port is "OnAuthRequiredArgDetailsFieldChallenger.port"
	//
	// Required
	Port int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnAuthRequiredArgDetailsFieldChallenger with all fields set.
func (p OnAuthRequiredArgDetailsFieldChallenger) FromRef(ref js.Ref) OnAuthRequiredArgDetailsFieldChallenger {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnAuthRequiredArgDetailsFieldChallenger in the application heap.
func (p OnAuthRequiredArgDetailsFieldChallenger) New() js.Ref {
	return bindings.OnAuthRequiredArgDetailsFieldChallengerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnAuthRequiredArgDetailsFieldChallenger) UpdateFrom(ref js.Ref) {
	bindings.OnAuthRequiredArgDetailsFieldChallengerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnAuthRequiredArgDetailsFieldChallenger) Update(ref js.Ref) {
	bindings.OnAuthRequiredArgDetailsFieldChallengerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnAuthRequiredArgDetailsFieldChallenger) FreeMembers(recursive bool) {
	js.Free(
		p.Host.Ref(),
	)
	p.Host = p.Host.FromRef(js.Undefined)
}

type ResourceType uint32

const (
	_ ResourceType = iota

	ResourceType_MAIN_FRAME
	ResourceType_SUB_FRAME
	ResourceType_STYLESHEET
	ResourceType_SCRIPT
	ResourceType_IMAGE
	ResourceType_FONT
	ResourceType_OBJECT
	ResourceType_XMLHTTPREQUEST
	ResourceType_PING
	ResourceType_CSP_REPORT
	ResourceType_MEDIA
	ResourceType_WEBSOCKET
	ResourceType_WEBBUNDLE
	ResourceType_OTHER
)

func (ResourceType) FromRef(str js.Ref) ResourceType {
	return ResourceType(bindings.ConstOfResourceType(str))
}

func (x ResourceType) String() (string, bool) {
	switch x {
	case ResourceType_MAIN_FRAME:
		return "main_frame", true
	case ResourceType_SUB_FRAME:
		return "sub_frame", true
	case ResourceType_STYLESHEET:
		return "stylesheet", true
	case ResourceType_SCRIPT:
		return "script", true
	case ResourceType_IMAGE:
		return "image", true
	case ResourceType_FONT:
		return "font", true
	case ResourceType_OBJECT:
		return "object", true
	case ResourceType_XMLHTTPREQUEST:
		return "xmlhttprequest", true
	case ResourceType_PING:
		return "ping", true
	case ResourceType_CSP_REPORT:
		return "csp_report", true
	case ResourceType_MEDIA:
		return "media", true
	case ResourceType_WEBSOCKET:
		return "websocket", true
	case ResourceType_WEBBUNDLE:
		return "webbundle", true
	case ResourceType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type OnAuthRequiredArgDetails struct {
	// Challenger is "OnAuthRequiredArgDetails.challenger"
	//
	// Required
	//
	// NOTE: Challenger.FFI_USE MUST be set to true to get Challenger used.
	Challenger OnAuthRequiredArgDetailsFieldChallenger
	// DocumentId is "OnAuthRequiredArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnAuthRequiredArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnAuthRequiredArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnAuthRequiredArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// Initiator is "OnAuthRequiredArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// IsProxy is "OnAuthRequiredArgDetails.isProxy"
	//
	// Required
	IsProxy bool
	// Method is "OnAuthRequiredArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnAuthRequiredArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnAuthRequiredArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// Realm is "OnAuthRequiredArgDetails.realm"
	//
	// Optional
	Realm js.String
	// RequestId is "OnAuthRequiredArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// ResponseHeaders is "OnAuthRequiredArgDetails.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders
	// Scheme is "OnAuthRequiredArgDetails.scheme"
	//
	// Required
	Scheme js.String
	// StatusCode is "OnAuthRequiredArgDetails.statusCode"
	//
	// Required
	StatusCode int64
	// StatusLine is "OnAuthRequiredArgDetails.statusLine"
	//
	// Required
	StatusLine js.String
	// TabId is "OnAuthRequiredArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnAuthRequiredArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnAuthRequiredArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnAuthRequiredArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnAuthRequiredArgDetails with all fields set.
func (p OnAuthRequiredArgDetails) FromRef(ref js.Ref) OnAuthRequiredArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnAuthRequiredArgDetails in the application heap.
func (p OnAuthRequiredArgDetails) New() js.Ref {
	return bindings.OnAuthRequiredArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnAuthRequiredArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnAuthRequiredArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnAuthRequiredArgDetails) Update(ref js.Ref) {
	bindings.OnAuthRequiredArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnAuthRequiredArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.Realm.Ref(),
		p.RequestId.Ref(),
		p.ResponseHeaders.Ref(),
		p.Scheme.Ref(),
		p.StatusLine.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.Realm = p.Realm.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.Scheme = p.Scheme.FromRef(js.Undefined)
	p.StatusLine = p.StatusLine.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	if recursive {
		p.Challenger.FreeMembers(true)
	}
}

type OnAuthRequiredOptions uint32

const (
	_ OnAuthRequiredOptions = iota

	OnAuthRequiredOptions_RESPONSE_HEADERS
	OnAuthRequiredOptions_BLOCKING
	OnAuthRequiredOptions_ASYNC_BLOCKING
	OnAuthRequiredOptions_EXTRA_HEADERS
)

func (OnAuthRequiredOptions) FromRef(str js.Ref) OnAuthRequiredOptions {
	return OnAuthRequiredOptions(bindings.ConstOfOnAuthRequiredOptions(str))
}

func (x OnAuthRequiredOptions) String() (string, bool) {
	switch x {
	case OnAuthRequiredOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case OnAuthRequiredOptions_BLOCKING:
		return "blocking", true
	case OnAuthRequiredOptions_ASYNC_BLOCKING:
		return "asyncBlocking", true
	case OnAuthRequiredOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnBeforeRedirectArgDetails struct {
	// DocumentId is "OnBeforeRedirectArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnBeforeRedirectArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnBeforeRedirectArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnBeforeRedirectArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// FromCache is "OnBeforeRedirectArgDetails.fromCache"
	//
	// Required
	FromCache bool
	// Initiator is "OnBeforeRedirectArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Ip is "OnBeforeRedirectArgDetails.ip"
	//
	// Optional
	Ip js.String
	// Method is "OnBeforeRedirectArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnBeforeRedirectArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnBeforeRedirectArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RedirectUrl is "OnBeforeRedirectArgDetails.redirectUrl"
	//
	// Required
	RedirectUrl js.String
	// RequestId is "OnBeforeRedirectArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// ResponseHeaders is "OnBeforeRedirectArgDetails.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders
	// StatusCode is "OnBeforeRedirectArgDetails.statusCode"
	//
	// Required
	StatusCode int64
	// StatusLine is "OnBeforeRedirectArgDetails.statusLine"
	//
	// Required
	StatusLine js.String
	// TabId is "OnBeforeRedirectArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnBeforeRedirectArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnBeforeRedirectArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnBeforeRedirectArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnBeforeRedirectArgDetails with all fields set.
func (p OnBeforeRedirectArgDetails) FromRef(ref js.Ref) OnBeforeRedirectArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnBeforeRedirectArgDetails in the application heap.
func (p OnBeforeRedirectArgDetails) New() js.Ref {
	return bindings.OnBeforeRedirectArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnBeforeRedirectArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnBeforeRedirectArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnBeforeRedirectArgDetails) Update(ref js.Ref) {
	bindings.OnBeforeRedirectArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnBeforeRedirectArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Ip.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RedirectUrl.Ref(),
		p.RequestId.Ref(),
		p.ResponseHeaders.Ref(),
		p.StatusLine.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Ip = p.Ip.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RedirectUrl = p.RedirectUrl.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.StatusLine = p.StatusLine.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnBeforeRedirectOptions uint32

const (
	_ OnBeforeRedirectOptions = iota

	OnBeforeRedirectOptions_RESPONSE_HEADERS
	OnBeforeRedirectOptions_EXTRA_HEADERS
)

func (OnBeforeRedirectOptions) FromRef(str js.Ref) OnBeforeRedirectOptions {
	return OnBeforeRedirectOptions(bindings.ConstOfOnBeforeRedirectOptions(str))
}

func (x OnBeforeRedirectOptions) String() (string, bool) {
	switch x {
	case OnBeforeRedirectOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case OnBeforeRedirectOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type UploadData struct {
	// Bytes is "UploadData.bytes"
	//
	// Optional
	Bytes js.Any
	// File is "UploadData.file"
	//
	// Optional
	File js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UploadData with all fields set.
func (p UploadData) FromRef(ref js.Ref) UploadData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UploadData in the application heap.
func (p UploadData) New() js.Ref {
	return bindings.UploadDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UploadData) UpdateFrom(ref js.Ref) {
	bindings.UploadDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UploadData) Update(ref js.Ref) {
	bindings.UploadDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UploadData) FreeMembers(recursive bool) {
	js.Free(
		p.Bytes.Ref(),
		p.File.Ref(),
	)
	p.Bytes = p.Bytes.FromRef(js.Undefined)
	p.File = p.File.FromRef(js.Undefined)
}

type OnBeforeRequestArgDetailsFieldRequestBody struct {
	// Error is "OnBeforeRequestArgDetailsFieldRequestBody.error"
	//
	// Optional
	Error js.String
	// FormData is "OnBeforeRequestArgDetailsFieldRequestBody.formData"
	//
	// Optional
	FormData js.Array[FormDataItem]
	// Raw is "OnBeforeRequestArgDetailsFieldRequestBody.raw"
	//
	// Optional
	Raw js.Array[UploadData]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnBeforeRequestArgDetailsFieldRequestBody with all fields set.
func (p OnBeforeRequestArgDetailsFieldRequestBody) FromRef(ref js.Ref) OnBeforeRequestArgDetailsFieldRequestBody {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnBeforeRequestArgDetailsFieldRequestBody in the application heap.
func (p OnBeforeRequestArgDetailsFieldRequestBody) New() js.Ref {
	return bindings.OnBeforeRequestArgDetailsFieldRequestBodyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnBeforeRequestArgDetailsFieldRequestBody) UpdateFrom(ref js.Ref) {
	bindings.OnBeforeRequestArgDetailsFieldRequestBodyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnBeforeRequestArgDetailsFieldRequestBody) Update(ref js.Ref) {
	bindings.OnBeforeRequestArgDetailsFieldRequestBodyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnBeforeRequestArgDetailsFieldRequestBody) FreeMembers(recursive bool) {
	js.Free(
		p.Error.Ref(),
		p.FormData.Ref(),
		p.Raw.Ref(),
	)
	p.Error = p.Error.FromRef(js.Undefined)
	p.FormData = p.FormData.FromRef(js.Undefined)
	p.Raw = p.Raw.FromRef(js.Undefined)
}

type OnBeforeRequestArgDetails struct {
	// DocumentId is "OnBeforeRequestArgDetails.documentId"
	//
	// Optional
	DocumentId js.String
	// DocumentLifecycle is "OnBeforeRequestArgDetails.documentLifecycle"
	//
	// Optional
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnBeforeRequestArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnBeforeRequestArgDetails.frameType"
	//
	// Optional
	FrameType extensiontypes.FrameType
	// Initiator is "OnBeforeRequestArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "OnBeforeRequestArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnBeforeRequestArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnBeforeRequestArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestBody is "OnBeforeRequestArgDetails.requestBody"
	//
	// Optional
	//
	// NOTE: RequestBody.FFI_USE MUST be set to true to get RequestBody used.
	RequestBody OnBeforeRequestArgDetailsFieldRequestBody
	// RequestId is "OnBeforeRequestArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// TabId is "OnBeforeRequestArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnBeforeRequestArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnBeforeRequestArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnBeforeRequestArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnBeforeRequestArgDetails with all fields set.
func (p OnBeforeRequestArgDetails) FromRef(ref js.Ref) OnBeforeRequestArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnBeforeRequestArgDetails in the application heap.
func (p OnBeforeRequestArgDetails) New() js.Ref {
	return bindings.OnBeforeRequestArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnBeforeRequestArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnBeforeRequestArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnBeforeRequestArgDetails) Update(ref js.Ref) {
	bindings.OnBeforeRequestArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnBeforeRequestArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	if recursive {
		p.RequestBody.FreeMembers(true)
	}
}

type OnBeforeRequestOptions uint32

const (
	_ OnBeforeRequestOptions = iota

	OnBeforeRequestOptions_BLOCKING
	OnBeforeRequestOptions_REQUEST_BODY
	OnBeforeRequestOptions_EXTRA_HEADERS
)

func (OnBeforeRequestOptions) FromRef(str js.Ref) OnBeforeRequestOptions {
	return OnBeforeRequestOptions(bindings.ConstOfOnBeforeRequestOptions(str))
}

func (x OnBeforeRequestOptions) String() (string, bool) {
	switch x {
	case OnBeforeRequestOptions_BLOCKING:
		return "blocking", true
	case OnBeforeRequestOptions_REQUEST_BODY:
		return "requestBody", true
	case OnBeforeRequestOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnBeforeSendHeadersArgDetails struct {
	// DocumentId is "OnBeforeSendHeadersArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnBeforeSendHeadersArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnBeforeSendHeadersArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnBeforeSendHeadersArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// Initiator is "OnBeforeSendHeadersArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "OnBeforeSendHeadersArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnBeforeSendHeadersArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnBeforeSendHeadersArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestHeaders is "OnBeforeSendHeadersArgDetails.requestHeaders"
	//
	// Optional
	RequestHeaders HttpHeaders
	// RequestId is "OnBeforeSendHeadersArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// TabId is "OnBeforeSendHeadersArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnBeforeSendHeadersArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnBeforeSendHeadersArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnBeforeSendHeadersArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnBeforeSendHeadersArgDetails with all fields set.
func (p OnBeforeSendHeadersArgDetails) FromRef(ref js.Ref) OnBeforeSendHeadersArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnBeforeSendHeadersArgDetails in the application heap.
func (p OnBeforeSendHeadersArgDetails) New() js.Ref {
	return bindings.OnBeforeSendHeadersArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnBeforeSendHeadersArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnBeforeSendHeadersArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnBeforeSendHeadersArgDetails) Update(ref js.Ref) {
	bindings.OnBeforeSendHeadersArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnBeforeSendHeadersArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestHeaders.Ref(),
		p.RequestId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestHeaders = p.RequestHeaders.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnBeforeSendHeadersOptions uint32

const (
	_ OnBeforeSendHeadersOptions = iota

	OnBeforeSendHeadersOptions_REQUEST_HEADERS
	OnBeforeSendHeadersOptions_BLOCKING
	OnBeforeSendHeadersOptions_EXTRA_HEADERS
)

func (OnBeforeSendHeadersOptions) FromRef(str js.Ref) OnBeforeSendHeadersOptions {
	return OnBeforeSendHeadersOptions(bindings.ConstOfOnBeforeSendHeadersOptions(str))
}

func (x OnBeforeSendHeadersOptions) String() (string, bool) {
	switch x {
	case OnBeforeSendHeadersOptions_REQUEST_HEADERS:
		return "requestHeaders", true
	case OnBeforeSendHeadersOptions_BLOCKING:
		return "blocking", true
	case OnBeforeSendHeadersOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnCompletedArgDetails struct {
	// DocumentId is "OnCompletedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnCompletedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnCompletedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnCompletedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// FromCache is "OnCompletedArgDetails.fromCache"
	//
	// Required
	FromCache bool
	// Initiator is "OnCompletedArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Ip is "OnCompletedArgDetails.ip"
	//
	// Optional
	Ip js.String
	// Method is "OnCompletedArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnCompletedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnCompletedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestId is "OnCompletedArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// ResponseHeaders is "OnCompletedArgDetails.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders
	// StatusCode is "OnCompletedArgDetails.statusCode"
	//
	// Required
	StatusCode int64
	// StatusLine is "OnCompletedArgDetails.statusLine"
	//
	// Required
	StatusLine js.String
	// TabId is "OnCompletedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnCompletedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnCompletedArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnCompletedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnCompletedArgDetails with all fields set.
func (p OnCompletedArgDetails) FromRef(ref js.Ref) OnCompletedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnCompletedArgDetails in the application heap.
func (p OnCompletedArgDetails) New() js.Ref {
	return bindings.OnCompletedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnCompletedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnCompletedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnCompletedArgDetails) Update(ref js.Ref) {
	bindings.OnCompletedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnCompletedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Ip.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.ResponseHeaders.Ref(),
		p.StatusLine.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Ip = p.Ip.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.StatusLine = p.StatusLine.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnCompletedOptions uint32

const (
	_ OnCompletedOptions = iota

	OnCompletedOptions_RESPONSE_HEADERS
	OnCompletedOptions_EXTRA_HEADERS
)

func (OnCompletedOptions) FromRef(str js.Ref) OnCompletedOptions {
	return OnCompletedOptions(bindings.ConstOfOnCompletedOptions(str))
}

func (x OnCompletedOptions) String() (string, bool) {
	switch x {
	case OnCompletedOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case OnCompletedOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnErrorOccurredArgDetails struct {
	// DocumentId is "OnErrorOccurredArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnErrorOccurredArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// Error is "OnErrorOccurredArgDetails.error"
	//
	// Required
	Error js.String
	// FrameId is "OnErrorOccurredArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnErrorOccurredArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// FromCache is "OnErrorOccurredArgDetails.fromCache"
	//
	// Required
	FromCache bool
	// Initiator is "OnErrorOccurredArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Ip is "OnErrorOccurredArgDetails.ip"
	//
	// Optional
	Ip js.String
	// Method is "OnErrorOccurredArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnErrorOccurredArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnErrorOccurredArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestId is "OnErrorOccurredArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// TabId is "OnErrorOccurredArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnErrorOccurredArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnErrorOccurredArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnErrorOccurredArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnErrorOccurredArgDetails with all fields set.
func (p OnErrorOccurredArgDetails) FromRef(ref js.Ref) OnErrorOccurredArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnErrorOccurredArgDetails in the application heap.
func (p OnErrorOccurredArgDetails) New() js.Ref {
	return bindings.OnErrorOccurredArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnErrorOccurredArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnErrorOccurredArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnErrorOccurredArgDetails) Update(ref js.Ref) {
	bindings.OnErrorOccurredArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnErrorOccurredArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Error.Ref(),
		p.Initiator.Ref(),
		p.Ip.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Ip = p.Ip.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnErrorOccurredOptions uint32

const (
	_ OnErrorOccurredOptions = iota

	OnErrorOccurredOptions_EXTRA_HEADERS
)

func (OnErrorOccurredOptions) FromRef(str js.Ref) OnErrorOccurredOptions {
	return OnErrorOccurredOptions(bindings.ConstOfOnErrorOccurredOptions(str))
}

func (x OnErrorOccurredOptions) String() (string, bool) {
	switch x {
	case OnErrorOccurredOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnHeadersReceivedArgDetails struct {
	// DocumentId is "OnHeadersReceivedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnHeadersReceivedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnHeadersReceivedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnHeadersReceivedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// Initiator is "OnHeadersReceivedArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "OnHeadersReceivedArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnHeadersReceivedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnHeadersReceivedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestId is "OnHeadersReceivedArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// ResponseHeaders is "OnHeadersReceivedArgDetails.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders
	// StatusCode is "OnHeadersReceivedArgDetails.statusCode"
	//
	// Required
	StatusCode int64
	// StatusLine is "OnHeadersReceivedArgDetails.statusLine"
	//
	// Required
	StatusLine js.String
	// TabId is "OnHeadersReceivedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnHeadersReceivedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnHeadersReceivedArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnHeadersReceivedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnHeadersReceivedArgDetails with all fields set.
func (p OnHeadersReceivedArgDetails) FromRef(ref js.Ref) OnHeadersReceivedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnHeadersReceivedArgDetails in the application heap.
func (p OnHeadersReceivedArgDetails) New() js.Ref {
	return bindings.OnHeadersReceivedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnHeadersReceivedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnHeadersReceivedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnHeadersReceivedArgDetails) Update(ref js.Ref) {
	bindings.OnHeadersReceivedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnHeadersReceivedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.ResponseHeaders.Ref(),
		p.StatusLine.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.StatusLine = p.StatusLine.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnHeadersReceivedOptions uint32

const (
	_ OnHeadersReceivedOptions = iota

	OnHeadersReceivedOptions_BLOCKING
	OnHeadersReceivedOptions_RESPONSE_HEADERS
	OnHeadersReceivedOptions_EXTRA_HEADERS
)

func (OnHeadersReceivedOptions) FromRef(str js.Ref) OnHeadersReceivedOptions {
	return OnHeadersReceivedOptions(bindings.ConstOfOnHeadersReceivedOptions(str))
}

func (x OnHeadersReceivedOptions) String() (string, bool) {
	switch x {
	case OnHeadersReceivedOptions_BLOCKING:
		return "blocking", true
	case OnHeadersReceivedOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case OnHeadersReceivedOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnResponseStartedArgDetails struct {
	// DocumentId is "OnResponseStartedArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnResponseStartedArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnResponseStartedArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnResponseStartedArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// FromCache is "OnResponseStartedArgDetails.fromCache"
	//
	// Required
	FromCache bool
	// Initiator is "OnResponseStartedArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Ip is "OnResponseStartedArgDetails.ip"
	//
	// Optional
	Ip js.String
	// Method is "OnResponseStartedArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnResponseStartedArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnResponseStartedArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestId is "OnResponseStartedArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// ResponseHeaders is "OnResponseStartedArgDetails.responseHeaders"
	//
	// Optional
	ResponseHeaders HttpHeaders
	// StatusCode is "OnResponseStartedArgDetails.statusCode"
	//
	// Required
	StatusCode int64
	// StatusLine is "OnResponseStartedArgDetails.statusLine"
	//
	// Required
	StatusLine js.String
	// TabId is "OnResponseStartedArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnResponseStartedArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnResponseStartedArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnResponseStartedArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnResponseStartedArgDetails with all fields set.
func (p OnResponseStartedArgDetails) FromRef(ref js.Ref) OnResponseStartedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnResponseStartedArgDetails in the application heap.
func (p OnResponseStartedArgDetails) New() js.Ref {
	return bindings.OnResponseStartedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnResponseStartedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnResponseStartedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnResponseStartedArgDetails) Update(ref js.Ref) {
	bindings.OnResponseStartedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnResponseStartedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Ip.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.ResponseHeaders.Ref(),
		p.StatusLine.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Ip = p.Ip.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.StatusLine = p.StatusLine.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnResponseStartedOptions uint32

const (
	_ OnResponseStartedOptions = iota

	OnResponseStartedOptions_RESPONSE_HEADERS
	OnResponseStartedOptions_EXTRA_HEADERS
)

func (OnResponseStartedOptions) FromRef(str js.Ref) OnResponseStartedOptions {
	return OnResponseStartedOptions(bindings.ConstOfOnResponseStartedOptions(str))
}

func (x OnResponseStartedOptions) String() (string, bool) {
	switch x {
	case OnResponseStartedOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case OnResponseStartedOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type OnSendHeadersArgDetails struct {
	// DocumentId is "OnSendHeadersArgDetails.documentId"
	//
	// Required
	DocumentId js.String
	// DocumentLifecycle is "OnSendHeadersArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnSendHeadersArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnSendHeadersArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// Initiator is "OnSendHeadersArgDetails.initiator"
	//
	// Optional
	Initiator js.String
	// Method is "OnSendHeadersArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnSendHeadersArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnSendHeadersArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestHeaders is "OnSendHeadersArgDetails.requestHeaders"
	//
	// Optional
	RequestHeaders HttpHeaders
	// RequestId is "OnSendHeadersArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// TabId is "OnSendHeadersArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnSendHeadersArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnSendHeadersArgDetails.type"
	//
	// Required
	Type ResourceType
	// Url is "OnSendHeadersArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnSendHeadersArgDetails with all fields set.
func (p OnSendHeadersArgDetails) FromRef(ref js.Ref) OnSendHeadersArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnSendHeadersArgDetails in the application heap.
func (p OnSendHeadersArgDetails) New() js.Ref {
	return bindings.OnSendHeadersArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnSendHeadersArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnSendHeadersArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnSendHeadersArgDetails) Update(ref js.Ref) {
	bindings.OnSendHeadersArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnSendHeadersArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Initiator.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestHeaders.Ref(),
		p.RequestId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Initiator = p.Initiator.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestHeaders = p.RequestHeaders.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnSendHeadersOptions uint32

const (
	_ OnSendHeadersOptions = iota

	OnSendHeadersOptions_REQUEST_HEADERS
	OnSendHeadersOptions_EXTRA_HEADERS
)

func (OnSendHeadersOptions) FromRef(str js.Ref) OnSendHeadersOptions {
	return OnSendHeadersOptions(bindings.ConstOfOnSendHeadersOptions(str))
}

func (x OnSendHeadersOptions) String() (string, bool) {
	switch x {
	case OnSendHeadersOptions_REQUEST_HEADERS:
		return "requestHeaders", true
	case OnSendHeadersOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

type RequestFilter struct {
	// TabId is "RequestFilter.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Types is "RequestFilter.types"
	//
	// Optional
	Types js.Array[ResourceType]
	// Urls is "RequestFilter.urls"
	//
	// Required
	Urls js.Array[js.String]
	// WindowId is "RequestFilter.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_TabId    bool // for TabId.
	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestFilter with all fields set.
func (p RequestFilter) FromRef(ref js.Ref) RequestFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestFilter in the application heap.
func (p RequestFilter) New() js.Ref {
	return bindings.RequestFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestFilter) UpdateFrom(ref js.Ref) {
	bindings.RequestFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestFilter) Update(ref js.Ref) {
	bindings.RequestFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestFilter) FreeMembers(recursive bool) {
	js.Free(
		p.Types.Ref(),
		p.Urls.Ref(),
	)
	p.Types = p.Types.FromRef(js.Undefined)
	p.Urls = p.Urls.FromRef(js.Undefined)
}

// HasFuncHandlerBehaviorChanged returns true if the function "WEBEXT.webRequest.handlerBehaviorChanged" exists.
func HasFuncHandlerBehaviorChanged() bool {
	return js.True == bindings.HasFuncHandlerBehaviorChanged()
}

// FuncHandlerBehaviorChanged returns the function "WEBEXT.webRequest.handlerBehaviorChanged".
func FuncHandlerBehaviorChanged() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHandlerBehaviorChanged(
		js.Pointer(&fn),
	)
	return
}

// HandlerBehaviorChanged calls the function "WEBEXT.webRequest.handlerBehaviorChanged" directly.
func HandlerBehaviorChanged() (ret js.Promise[js.Void]) {
	bindings.CallHandlerBehaviorChanged(
		js.Pointer(&ret),
	)

	return
}

// TryHandlerBehaviorChanged calls the function "WEBEXT.webRequest.handlerBehaviorChanged"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHandlerBehaviorChanged() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHandlerBehaviorChanged(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnActionIgnoredEventCallbackFunc func(this js.Ref, details *OnActionIgnoredArgDetails) js.Ref

func (fn OnActionIgnoredEventCallbackFunc) Register() js.Func[func(details *OnActionIgnoredArgDetails)] {
	return js.RegisterCallback[func(details *OnActionIgnoredArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnActionIgnoredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnActionIgnoredArgDetails
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

type OnActionIgnoredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnActionIgnoredArgDetails) js.Ref
	Arg T
}

func (cb *OnActionIgnoredEventCallback[T]) Register() js.Func[func(details *OnActionIgnoredArgDetails)] {
	return js.RegisterCallback[func(details *OnActionIgnoredArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnActionIgnoredEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnActionIgnoredArgDetails
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

// HasFuncOnActionIgnored returns true if the function "WEBEXT.webRequest.onActionIgnored.addListener" exists.
func HasFuncOnActionIgnored() bool {
	return js.True == bindings.HasFuncOnActionIgnored()
}

// FuncOnActionIgnored returns the function "WEBEXT.webRequest.onActionIgnored.addListener".
func FuncOnActionIgnored() (fn js.Func[func(callback js.Func[func(details *OnActionIgnoredArgDetails)])]) {
	bindings.FuncOnActionIgnored(
		js.Pointer(&fn),
	)
	return
}

// OnActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.addListener" directly.
func OnActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret js.Void) {
	bindings.CallOnActionIgnored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnActionIgnored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffActionIgnored returns true if the function "WEBEXT.webRequest.onActionIgnored.removeListener" exists.
func HasFuncOffActionIgnored() bool {
	return js.True == bindings.HasFuncOffActionIgnored()
}

// FuncOffActionIgnored returns the function "WEBEXT.webRequest.onActionIgnored.removeListener".
func FuncOffActionIgnored() (fn js.Func[func(callback js.Func[func(details *OnActionIgnoredArgDetails)])]) {
	bindings.FuncOffActionIgnored(
		js.Pointer(&fn),
	)
	return
}

// OffActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.removeListener" directly.
func OffActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret js.Void) {
	bindings.CallOffActionIgnored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffActionIgnored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnActionIgnored returns true if the function "WEBEXT.webRequest.onActionIgnored.hasListener" exists.
func HasFuncHasOnActionIgnored() bool {
	return js.True == bindings.HasFuncHasOnActionIgnored()
}

// FuncHasOnActionIgnored returns the function "WEBEXT.webRequest.onActionIgnored.hasListener".
func FuncHasOnActionIgnored() (fn js.Func[func(callback js.Func[func(details *OnActionIgnoredArgDetails)]) bool]) {
	bindings.FuncHasOnActionIgnored(
		js.Pointer(&fn),
	)
	return
}

// HasOnActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.hasListener" directly.
func HasOnActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret bool) {
	bindings.CallHasOnActionIgnored(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnActionIgnored calls the function "WEBEXT.webRequest.onActionIgnored.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnActionIgnored(callback js.Func[func(details *OnActionIgnoredArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnActionIgnored(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAuthRequiredEventCallbackFunc func(this js.Ref, details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) js.Ref

func (fn OnAuthRequiredEventCallbackFunc) Register() js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse] {
	return js.RegisterCallback[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAuthRequiredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnAuthRequiredArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(response *BlockingResponse)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAuthRequiredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) js.Ref
	Arg T
}

func (cb *OnAuthRequiredEventCallback[T]) Register() js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse] {
	return js.RegisterCallback[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAuthRequiredEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnAuthRequiredArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
		js.Func[func(response *BlockingResponse)]{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAuthRequired returns true if the function "WEBEXT.webRequest.onAuthRequired.addListener" exists.
func HasFuncOnAuthRequired() bool {
	return js.True == bindings.HasFuncOnAuthRequired()
}

// FuncOnAuthRequired returns the function "WEBEXT.webRequest.onAuthRequired.addListener".
func FuncOnAuthRequired() (fn js.Func[func(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse])]) {
	bindings.FuncOnAuthRequired(
		js.Pointer(&fn),
	)
	return
}

// OnAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.addListener" directly.
func OnAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret js.Void) {
	bindings.CallOnAuthRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAuthRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAuthRequired returns true if the function "WEBEXT.webRequest.onAuthRequired.removeListener" exists.
func HasFuncOffAuthRequired() bool {
	return js.True == bindings.HasFuncOffAuthRequired()
}

// FuncOffAuthRequired returns the function "WEBEXT.webRequest.onAuthRequired.removeListener".
func FuncOffAuthRequired() (fn js.Func[func(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse])]) {
	bindings.FuncOffAuthRequired(
		js.Pointer(&fn),
	)
	return
}

// OffAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.removeListener" directly.
func OffAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret js.Void) {
	bindings.CallOffAuthRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAuthRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAuthRequired returns true if the function "WEBEXT.webRequest.onAuthRequired.hasListener" exists.
func HasFuncHasOnAuthRequired() bool {
	return js.True == bindings.HasFuncHasOnAuthRequired()
}

// FuncHasOnAuthRequired returns the function "WEBEXT.webRequest.onAuthRequired.hasListener".
func FuncHasOnAuthRequired() (fn js.Func[func(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) bool]) {
	bindings.FuncHasOnAuthRequired(
		js.Pointer(&fn),
	)
	return
}

// HasOnAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.hasListener" directly.
func HasOnAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret bool) {
	bindings.CallHasOnAuthRequired(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAuthRequired calls the function "WEBEXT.webRequest.onAuthRequired.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAuthRequired(callback js.Func[func(details *OnAuthRequiredArgDetails, asyncCallback js.Func[func(response *BlockingResponse)]) BlockingResponse]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAuthRequired(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBeforeRedirectEventCallbackFunc func(this js.Ref, details *OnBeforeRedirectArgDetails) js.Ref

func (fn OnBeforeRedirectEventCallbackFunc) Register() js.Func[func(details *OnBeforeRedirectArgDetails)] {
	return js.RegisterCallback[func(details *OnBeforeRedirectArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeforeRedirectEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeRedirectArgDetails
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

type OnBeforeRedirectEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnBeforeRedirectArgDetails) js.Ref
	Arg T
}

func (cb *OnBeforeRedirectEventCallback[T]) Register() js.Func[func(details *OnBeforeRedirectArgDetails)] {
	return js.RegisterCallback[func(details *OnBeforeRedirectArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeforeRedirectEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeRedirectArgDetails
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

// HasFuncOnBeforeRedirect returns true if the function "WEBEXT.webRequest.onBeforeRedirect.addListener" exists.
func HasFuncOnBeforeRedirect() bool {
	return js.True == bindings.HasFuncOnBeforeRedirect()
}

// FuncOnBeforeRedirect returns the function "WEBEXT.webRequest.onBeforeRedirect.addListener".
func FuncOnBeforeRedirect() (fn js.Func[func(callback js.Func[func(details *OnBeforeRedirectArgDetails)])]) {
	bindings.FuncOnBeforeRedirect(
		js.Pointer(&fn),
	)
	return
}

// OnBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.addListener" directly.
func OnBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret js.Void) {
	bindings.CallOnBeforeRedirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBeforeRedirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBeforeRedirect returns true if the function "WEBEXT.webRequest.onBeforeRedirect.removeListener" exists.
func HasFuncOffBeforeRedirect() bool {
	return js.True == bindings.HasFuncOffBeforeRedirect()
}

// FuncOffBeforeRedirect returns the function "WEBEXT.webRequest.onBeforeRedirect.removeListener".
func FuncOffBeforeRedirect() (fn js.Func[func(callback js.Func[func(details *OnBeforeRedirectArgDetails)])]) {
	bindings.FuncOffBeforeRedirect(
		js.Pointer(&fn),
	)
	return
}

// OffBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.removeListener" directly.
func OffBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret js.Void) {
	bindings.CallOffBeforeRedirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBeforeRedirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBeforeRedirect returns true if the function "WEBEXT.webRequest.onBeforeRedirect.hasListener" exists.
func HasFuncHasOnBeforeRedirect() bool {
	return js.True == bindings.HasFuncHasOnBeforeRedirect()
}

// FuncHasOnBeforeRedirect returns the function "WEBEXT.webRequest.onBeforeRedirect.hasListener".
func FuncHasOnBeforeRedirect() (fn js.Func[func(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) bool]) {
	bindings.FuncHasOnBeforeRedirect(
		js.Pointer(&fn),
	)
	return
}

// HasOnBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.hasListener" directly.
func HasOnBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret bool) {
	bindings.CallHasOnBeforeRedirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBeforeRedirect calls the function "WEBEXT.webRequest.onBeforeRedirect.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBeforeRedirect(callback js.Func[func(details *OnBeforeRedirectArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBeforeRedirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBeforeRequestEventCallbackFunc func(this js.Ref, details *OnBeforeRequestArgDetails) js.Ref

func (fn OnBeforeRequestEventCallbackFunc) Register() js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnBeforeRequestArgDetails) BlockingResponse](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeforeRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeRequestArgDetails
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

type OnBeforeRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnBeforeRequestArgDetails) js.Ref
	Arg T
}

func (cb *OnBeforeRequestEventCallback[T]) Register() js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnBeforeRequestArgDetails) BlockingResponse](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeforeRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeRequestArgDetails
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

// HasFuncOnBeforeRequest returns true if the function "WEBEXT.webRequest.onBeforeRequest.addListener" exists.
func HasFuncOnBeforeRequest() bool {
	return js.True == bindings.HasFuncOnBeforeRequest()
}

// FuncOnBeforeRequest returns the function "WEBEXT.webRequest.onBeforeRequest.addListener".
func FuncOnBeforeRequest() (fn js.Func[func(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse])]) {
	bindings.FuncOnBeforeRequest(
		js.Pointer(&fn),
	)
	return
}

// OnBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.addListener" directly.
func OnBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOnBeforeRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBeforeRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBeforeRequest returns true if the function "WEBEXT.webRequest.onBeforeRequest.removeListener" exists.
func HasFuncOffBeforeRequest() bool {
	return js.True == bindings.HasFuncOffBeforeRequest()
}

// FuncOffBeforeRequest returns the function "WEBEXT.webRequest.onBeforeRequest.removeListener".
func FuncOffBeforeRequest() (fn js.Func[func(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse])]) {
	bindings.FuncOffBeforeRequest(
		js.Pointer(&fn),
	)
	return
}

// OffBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.removeListener" directly.
func OffBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOffBeforeRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBeforeRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBeforeRequest returns true if the function "WEBEXT.webRequest.onBeforeRequest.hasListener" exists.
func HasFuncHasOnBeforeRequest() bool {
	return js.True == bindings.HasFuncHasOnBeforeRequest()
}

// FuncHasOnBeforeRequest returns the function "WEBEXT.webRequest.onBeforeRequest.hasListener".
func FuncHasOnBeforeRequest() (fn js.Func[func(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) bool]) {
	bindings.FuncHasOnBeforeRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.hasListener" directly.
func HasOnBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret bool) {
	bindings.CallHasOnBeforeRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBeforeRequest calls the function "WEBEXT.webRequest.onBeforeRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBeforeRequest(callback js.Func[func(details *OnBeforeRequestArgDetails) BlockingResponse]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBeforeRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnBeforeSendHeadersEventCallbackFunc func(this js.Ref, details *OnBeforeSendHeadersArgDetails) js.Ref

func (fn OnBeforeSendHeadersEventCallbackFunc) Register() js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnBeforeSendHeadersEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeSendHeadersArgDetails
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

type OnBeforeSendHeadersEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnBeforeSendHeadersArgDetails) js.Ref
	Arg T
}

func (cb *OnBeforeSendHeadersEventCallback[T]) Register() js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnBeforeSendHeadersEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnBeforeSendHeadersArgDetails
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

// HasFuncOnBeforeSendHeaders returns true if the function "WEBEXT.webRequest.onBeforeSendHeaders.addListener" exists.
func HasFuncOnBeforeSendHeaders() bool {
	return js.True == bindings.HasFuncOnBeforeSendHeaders()
}

// FuncOnBeforeSendHeaders returns the function "WEBEXT.webRequest.onBeforeSendHeaders.addListener".
func FuncOnBeforeSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse])]) {
	bindings.FuncOnBeforeSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// OnBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.addListener" directly.
func OnBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOnBeforeSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnBeforeSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffBeforeSendHeaders returns true if the function "WEBEXT.webRequest.onBeforeSendHeaders.removeListener" exists.
func HasFuncOffBeforeSendHeaders() bool {
	return js.True == bindings.HasFuncOffBeforeSendHeaders()
}

// FuncOffBeforeSendHeaders returns the function "WEBEXT.webRequest.onBeforeSendHeaders.removeListener".
func FuncOffBeforeSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse])]) {
	bindings.FuncOffBeforeSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// OffBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.removeListener" directly.
func OffBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOffBeforeSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffBeforeSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnBeforeSendHeaders returns true if the function "WEBEXT.webRequest.onBeforeSendHeaders.hasListener" exists.
func HasFuncHasOnBeforeSendHeaders() bool {
	return js.True == bindings.HasFuncHasOnBeforeSendHeaders()
}

// FuncHasOnBeforeSendHeaders returns the function "WEBEXT.webRequest.onBeforeSendHeaders.hasListener".
func FuncHasOnBeforeSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) bool]) {
	bindings.FuncHasOnBeforeSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// HasOnBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.hasListener" directly.
func HasOnBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret bool) {
	bindings.CallHasOnBeforeSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnBeforeSendHeaders calls the function "WEBEXT.webRequest.onBeforeSendHeaders.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnBeforeSendHeaders(callback js.Func[func(details *OnBeforeSendHeadersArgDetails) BlockingResponse]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnBeforeSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnCompletedEventCallbackFunc func(this js.Ref, details *OnCompletedArgDetails) js.Ref

func (fn OnCompletedEventCallbackFunc) Register() js.Func[func(details *OnCompletedArgDetails)] {
	return js.RegisterCallback[func(details *OnCompletedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCompletedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCompletedArgDetails
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

type OnCompletedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnCompletedArgDetails) js.Ref
	Arg T
}

func (cb *OnCompletedEventCallback[T]) Register() js.Func[func(details *OnCompletedArgDetails)] {
	return js.RegisterCallback[func(details *OnCompletedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCompletedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnCompletedArgDetails
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

// HasFuncOnCompleted returns true if the function "WEBEXT.webRequest.onCompleted.addListener" exists.
func HasFuncOnCompleted() bool {
	return js.True == bindings.HasFuncOnCompleted()
}

// FuncOnCompleted returns the function "WEBEXT.webRequest.onCompleted.addListener".
func FuncOnCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)])]) {
	bindings.FuncOnCompleted(
		js.Pointer(&fn),
	)
	return
}

// OnCompleted calls the function "WEBEXT.webRequest.onCompleted.addListener" directly.
func OnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void) {
	bindings.CallOnCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCompleted calls the function "WEBEXT.webRequest.onCompleted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCompleted returns true if the function "WEBEXT.webRequest.onCompleted.removeListener" exists.
func HasFuncOffCompleted() bool {
	return js.True == bindings.HasFuncOffCompleted()
}

// FuncOffCompleted returns the function "WEBEXT.webRequest.onCompleted.removeListener".
func FuncOffCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)])]) {
	bindings.FuncOffCompleted(
		js.Pointer(&fn),
	)
	return
}

// OffCompleted calls the function "WEBEXT.webRequest.onCompleted.removeListener" directly.
func OffCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void) {
	bindings.CallOffCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCompleted calls the function "WEBEXT.webRequest.onCompleted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCompleted returns true if the function "WEBEXT.webRequest.onCompleted.hasListener" exists.
func HasFuncHasOnCompleted() bool {
	return js.True == bindings.HasFuncHasOnCompleted()
}

// FuncHasOnCompleted returns the function "WEBEXT.webRequest.onCompleted.hasListener".
func FuncHasOnCompleted() (fn js.Func[func(callback js.Func[func(details *OnCompletedArgDetails)]) bool]) {
	bindings.FuncHasOnCompleted(
		js.Pointer(&fn),
	)
	return
}

// HasOnCompleted calls the function "WEBEXT.webRequest.onCompleted.hasListener" directly.
func HasOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret bool) {
	bindings.CallHasOnCompleted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCompleted calls the function "WEBEXT.webRequest.onCompleted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCompleted(callback js.Func[func(details *OnCompletedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCompleted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnErrorOccurredEventCallbackFunc func(this js.Ref, details *OnErrorOccurredArgDetails) js.Ref

func (fn OnErrorOccurredEventCallbackFunc) Register() js.Func[func(details *OnErrorOccurredArgDetails)] {
	return js.RegisterCallback[func(details *OnErrorOccurredArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnErrorOccurredEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnErrorOccurredArgDetails
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

type OnErrorOccurredEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnErrorOccurredArgDetails) js.Ref
	Arg T
}

func (cb *OnErrorOccurredEventCallback[T]) Register() js.Func[func(details *OnErrorOccurredArgDetails)] {
	return js.RegisterCallback[func(details *OnErrorOccurredArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnErrorOccurredEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnErrorOccurredArgDetails
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

// HasFuncOnErrorOccurred returns true if the function "WEBEXT.webRequest.onErrorOccurred.addListener" exists.
func HasFuncOnErrorOccurred() bool {
	return js.True == bindings.HasFuncOnErrorOccurred()
}

// FuncOnErrorOccurred returns the function "WEBEXT.webRequest.onErrorOccurred.addListener".
func FuncOnErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)])]) {
	bindings.FuncOnErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// OnErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.addListener" directly.
func OnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void) {
	bindings.CallOnErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffErrorOccurred returns true if the function "WEBEXT.webRequest.onErrorOccurred.removeListener" exists.
func HasFuncOffErrorOccurred() bool {
	return js.True == bindings.HasFuncOffErrorOccurred()
}

// FuncOffErrorOccurred returns the function "WEBEXT.webRequest.onErrorOccurred.removeListener".
func FuncOffErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)])]) {
	bindings.FuncOffErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// OffErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.removeListener" directly.
func OffErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void) {
	bindings.CallOffErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnErrorOccurred returns true if the function "WEBEXT.webRequest.onErrorOccurred.hasListener" exists.
func HasFuncHasOnErrorOccurred() bool {
	return js.True == bindings.HasFuncHasOnErrorOccurred()
}

// FuncHasOnErrorOccurred returns the function "WEBEXT.webRequest.onErrorOccurred.hasListener".
func FuncHasOnErrorOccurred() (fn js.Func[func(callback js.Func[func(details *OnErrorOccurredArgDetails)]) bool]) {
	bindings.FuncHasOnErrorOccurred(
		js.Pointer(&fn),
	)
	return
}

// HasOnErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.hasListener" directly.
func HasOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret bool) {
	bindings.CallHasOnErrorOccurred(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnErrorOccurred calls the function "WEBEXT.webRequest.onErrorOccurred.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnErrorOccurred(callback js.Func[func(details *OnErrorOccurredArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnErrorOccurred(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnHeadersReceivedEventCallbackFunc func(this js.Ref, details *OnHeadersReceivedArgDetails) js.Ref

func (fn OnHeadersReceivedEventCallbackFunc) Register() js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnHeadersReceivedArgDetails) BlockingResponse](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnHeadersReceivedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHeadersReceivedArgDetails
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

type OnHeadersReceivedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnHeadersReceivedArgDetails) js.Ref
	Arg T
}

func (cb *OnHeadersReceivedEventCallback[T]) Register() js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse] {
	return js.RegisterCallback[func(details *OnHeadersReceivedArgDetails) BlockingResponse](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnHeadersReceivedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnHeadersReceivedArgDetails
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

// HasFuncOnHeadersReceived returns true if the function "WEBEXT.webRequest.onHeadersReceived.addListener" exists.
func HasFuncOnHeadersReceived() bool {
	return js.True == bindings.HasFuncOnHeadersReceived()
}

// FuncOnHeadersReceived returns the function "WEBEXT.webRequest.onHeadersReceived.addListener".
func FuncOnHeadersReceived() (fn js.Func[func(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse])]) {
	bindings.FuncOnHeadersReceived(
		js.Pointer(&fn),
	)
	return
}

// OnHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.addListener" directly.
func OnHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOnHeadersReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnHeadersReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffHeadersReceived returns true if the function "WEBEXT.webRequest.onHeadersReceived.removeListener" exists.
func HasFuncOffHeadersReceived() bool {
	return js.True == bindings.HasFuncOffHeadersReceived()
}

// FuncOffHeadersReceived returns the function "WEBEXT.webRequest.onHeadersReceived.removeListener".
func FuncOffHeadersReceived() (fn js.Func[func(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse])]) {
	bindings.FuncOffHeadersReceived(
		js.Pointer(&fn),
	)
	return
}

// OffHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.removeListener" directly.
func OffHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret js.Void) {
	bindings.CallOffHeadersReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffHeadersReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnHeadersReceived returns true if the function "WEBEXT.webRequest.onHeadersReceived.hasListener" exists.
func HasFuncHasOnHeadersReceived() bool {
	return js.True == bindings.HasFuncHasOnHeadersReceived()
}

// FuncHasOnHeadersReceived returns the function "WEBEXT.webRequest.onHeadersReceived.hasListener".
func FuncHasOnHeadersReceived() (fn js.Func[func(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) bool]) {
	bindings.FuncHasOnHeadersReceived(
		js.Pointer(&fn),
	)
	return
}

// HasOnHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.hasListener" directly.
func HasOnHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret bool) {
	bindings.CallHasOnHeadersReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnHeadersReceived calls the function "WEBEXT.webRequest.onHeadersReceived.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnHeadersReceived(callback js.Func[func(details *OnHeadersReceivedArgDetails) BlockingResponse]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnHeadersReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnResponseStartedEventCallbackFunc func(this js.Ref, details *OnResponseStartedArgDetails) js.Ref

func (fn OnResponseStartedEventCallbackFunc) Register() js.Func[func(details *OnResponseStartedArgDetails)] {
	return js.RegisterCallback[func(details *OnResponseStartedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnResponseStartedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnResponseStartedArgDetails
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

type OnResponseStartedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnResponseStartedArgDetails) js.Ref
	Arg T
}

func (cb *OnResponseStartedEventCallback[T]) Register() js.Func[func(details *OnResponseStartedArgDetails)] {
	return js.RegisterCallback[func(details *OnResponseStartedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnResponseStartedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnResponseStartedArgDetails
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

// HasFuncOnResponseStarted returns true if the function "WEBEXT.webRequest.onResponseStarted.addListener" exists.
func HasFuncOnResponseStarted() bool {
	return js.True == bindings.HasFuncOnResponseStarted()
}

// FuncOnResponseStarted returns the function "WEBEXT.webRequest.onResponseStarted.addListener".
func FuncOnResponseStarted() (fn js.Func[func(callback js.Func[func(details *OnResponseStartedArgDetails)])]) {
	bindings.FuncOnResponseStarted(
		js.Pointer(&fn),
	)
	return
}

// OnResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.addListener" directly.
func OnResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret js.Void) {
	bindings.CallOnResponseStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnResponseStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffResponseStarted returns true if the function "WEBEXT.webRequest.onResponseStarted.removeListener" exists.
func HasFuncOffResponseStarted() bool {
	return js.True == bindings.HasFuncOffResponseStarted()
}

// FuncOffResponseStarted returns the function "WEBEXT.webRequest.onResponseStarted.removeListener".
func FuncOffResponseStarted() (fn js.Func[func(callback js.Func[func(details *OnResponseStartedArgDetails)])]) {
	bindings.FuncOffResponseStarted(
		js.Pointer(&fn),
	)
	return
}

// OffResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.removeListener" directly.
func OffResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret js.Void) {
	bindings.CallOffResponseStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffResponseStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnResponseStarted returns true if the function "WEBEXT.webRequest.onResponseStarted.hasListener" exists.
func HasFuncHasOnResponseStarted() bool {
	return js.True == bindings.HasFuncHasOnResponseStarted()
}

// FuncHasOnResponseStarted returns the function "WEBEXT.webRequest.onResponseStarted.hasListener".
func FuncHasOnResponseStarted() (fn js.Func[func(callback js.Func[func(details *OnResponseStartedArgDetails)]) bool]) {
	bindings.FuncHasOnResponseStarted(
		js.Pointer(&fn),
	)
	return
}

// HasOnResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.hasListener" directly.
func HasOnResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret bool) {
	bindings.CallHasOnResponseStarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnResponseStarted calls the function "WEBEXT.webRequest.onResponseStarted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnResponseStarted(callback js.Func[func(details *OnResponseStartedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnResponseStarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSendHeadersEventCallbackFunc func(this js.Ref, details *OnSendHeadersArgDetails) js.Ref

func (fn OnSendHeadersEventCallbackFunc) Register() js.Func[func(details *OnSendHeadersArgDetails)] {
	return js.RegisterCallback[func(details *OnSendHeadersArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSendHeadersEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnSendHeadersArgDetails
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

type OnSendHeadersEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnSendHeadersArgDetails) js.Ref
	Arg T
}

func (cb *OnSendHeadersEventCallback[T]) Register() js.Func[func(details *OnSendHeadersArgDetails)] {
	return js.RegisterCallback[func(details *OnSendHeadersArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSendHeadersEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnSendHeadersArgDetails
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

// HasFuncOnSendHeaders returns true if the function "WEBEXT.webRequest.onSendHeaders.addListener" exists.
func HasFuncOnSendHeaders() bool {
	return js.True == bindings.HasFuncOnSendHeaders()
}

// FuncOnSendHeaders returns the function "WEBEXT.webRequest.onSendHeaders.addListener".
func FuncOnSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnSendHeadersArgDetails)])]) {
	bindings.FuncOnSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// OnSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.addListener" directly.
func OnSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret js.Void) {
	bindings.CallOnSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSendHeaders returns true if the function "WEBEXT.webRequest.onSendHeaders.removeListener" exists.
func HasFuncOffSendHeaders() bool {
	return js.True == bindings.HasFuncOffSendHeaders()
}

// FuncOffSendHeaders returns the function "WEBEXT.webRequest.onSendHeaders.removeListener".
func FuncOffSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnSendHeadersArgDetails)])]) {
	bindings.FuncOffSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// OffSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.removeListener" directly.
func OffSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret js.Void) {
	bindings.CallOffSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSendHeaders returns true if the function "WEBEXT.webRequest.onSendHeaders.hasListener" exists.
func HasFuncHasOnSendHeaders() bool {
	return js.True == bindings.HasFuncHasOnSendHeaders()
}

// FuncHasOnSendHeaders returns the function "WEBEXT.webRequest.onSendHeaders.hasListener".
func FuncHasOnSendHeaders() (fn js.Func[func(callback js.Func[func(details *OnSendHeadersArgDetails)]) bool]) {
	bindings.FuncHasOnSendHeaders(
		js.Pointer(&fn),
	)
	return
}

// HasOnSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.hasListener" directly.
func HasOnSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret bool) {
	bindings.CallHasOnSendHeaders(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSendHeaders calls the function "WEBEXT.webRequest.onSendHeaders.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSendHeaders(callback js.Func[func(details *OnSendHeadersArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSendHeaders(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

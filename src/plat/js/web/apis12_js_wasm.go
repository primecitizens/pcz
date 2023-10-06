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

const (
	_ RequestDestination = iota

	RequestDestination_
	RequestDestination_AUDIO
	RequestDestination_AUDIOWORKLET
	RequestDestination_DOCUMENT
	RequestDestination_EMBED
	RequestDestination_FONT
	RequestDestination_FRAME
	RequestDestination_IFRAME
	RequestDestination_IMAGE
	RequestDestination_MANIFEST
	RequestDestination_OBJECT
	RequestDestination_PAINTWORKLET
	RequestDestination_REPORT
	RequestDestination_SCRIPT
	RequestDestination_SHAREDWORKER
	RequestDestination_STYLE
	RequestDestination_TRACK
	RequestDestination_VIDEO
	RequestDestination_WORKER
	RequestDestination_XSLT
)

func (RequestDestination) FromRef(str js.Ref) RequestDestination {
	return RequestDestination(bindings.ConstOfRequestDestination(str))
}

func (x RequestDestination) String() (string, bool) {
	switch x {
	case RequestDestination_:
		return "", true
	case RequestDestination_AUDIO:
		return "audio", true
	case RequestDestination_AUDIOWORKLET:
		return "audioworklet", true
	case RequestDestination_DOCUMENT:
		return "document", true
	case RequestDestination_EMBED:
		return "embed", true
	case RequestDestination_FONT:
		return "font", true
	case RequestDestination_FRAME:
		return "frame", true
	case RequestDestination_IFRAME:
		return "iframe", true
	case RequestDestination_IMAGE:
		return "image", true
	case RequestDestination_MANIFEST:
		return "manifest", true
	case RequestDestination_OBJECT:
		return "object", true
	case RequestDestination_PAINTWORKLET:
		return "paintworklet", true
	case RequestDestination_REPORT:
		return "report", true
	case RequestDestination_SCRIPT:
		return "script", true
	case RequestDestination_SHAREDWORKER:
		return "sharedworker", true
	case RequestDestination_STYLE:
		return "style", true
	case RequestDestination_TRACK:
		return "track", true
	case RequestDestination_VIDEO:
		return "video", true
	case RequestDestination_WORKER:
		return "worker", true
	case RequestDestination_XSLT:
		return "xslt", true
	default:
		return "", false
	}
}

func NewRequest(input RequestInfo, init RequestInit) (ret Request) {
	ret.ref = bindings.NewRequestByRequest(
		input.Ref(),
		js.Pointer(&init))
	return
}

func NewRequestByRequest1(input RequestInfo) (ret Request) {
	ret.ref = bindings.NewRequestByRequest1(
		input.Ref())
	return
}

type Request struct {
	ref js.Ref
}

func (this Request) Once() Request {
	this.Ref().Once()
	return this
}

func (this Request) Ref() js.Ref {
	return this.ref
}

func (this Request) FromRef(ref js.Ref) Request {
	this.ref = ref
	return this
}

func (this Request) Free() {
	this.Ref().Free()
}

// Method returns the value of property "Request.method".
//
// It returns ok=false if there is no such property.
func (this Request) Method() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "Request.url".
//
// It returns ok=false if there is no such property.
func (this Request) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Headers returns the value of property "Request.headers".
//
// It returns ok=false if there is no such property.
func (this Request) Headers() (ret Headers, ok bool) {
	ok = js.True == bindings.GetRequestHeaders(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Destination returns the value of property "Request.destination".
//
// It returns ok=false if there is no such property.
func (this Request) Destination() (ret RequestDestination, ok bool) {
	ok = js.True == bindings.GetRequestDestination(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Referrer returns the value of property "Request.referrer".
//
// It returns ok=false if there is no such property.
func (this Request) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestReferrer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReferrerPolicy returns the value of property "Request.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this Request) ReferrerPolicy() (ret ReferrerPolicy, ok bool) {
	ok = js.True == bindings.GetRequestReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "Request.mode".
//
// It returns ok=false if there is no such property.
func (this Request) Mode() (ret RequestMode, ok bool) {
	ok = js.True == bindings.GetRequestMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Credentials returns the value of property "Request.credentials".
//
// It returns ok=false if there is no such property.
func (this Request) Credentials() (ret RequestCredentials, ok bool) {
	ok = js.True == bindings.GetRequestCredentials(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cache returns the value of property "Request.cache".
//
// It returns ok=false if there is no such property.
func (this Request) Cache() (ret RequestCache, ok bool) {
	ok = js.True == bindings.GetRequestCache(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Redirect returns the value of property "Request.redirect".
//
// It returns ok=false if there is no such property.
func (this Request) Redirect() (ret RequestRedirect, ok bool) {
	ok = js.True == bindings.GetRequestRedirect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Integrity returns the value of property "Request.integrity".
//
// It returns ok=false if there is no such property.
func (this Request) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestIntegrity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Keepalive returns the value of property "Request.keepalive".
//
// It returns ok=false if there is no such property.
func (this Request) Keepalive() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestKeepalive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsReloadNavigation returns the value of property "Request.isReloadNavigation".
//
// It returns ok=false if there is no such property.
func (this Request) IsReloadNavigation() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestIsReloadNavigation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsHistoryNavigation returns the value of property "Request.isHistoryNavigation".
//
// It returns ok=false if there is no such property.
func (this Request) IsHistoryNavigation() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestIsHistoryNavigation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Signal returns the value of property "Request.signal".
//
// It returns ok=false if there is no such property.
func (this Request) Signal() (ret AbortSignal, ok bool) {
	ok = js.True == bindings.GetRequestSignal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duplex returns the value of property "Request.duplex".
//
// It returns ok=false if there is no such property.
func (this Request) Duplex() (ret RequestDuplex, ok bool) {
	ok = js.True == bindings.GetRequestDuplex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Request.body".
//
// It returns ok=false if there is no such property.
func (this Request) Body() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetRequestBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BodyUsed returns the value of property "Request.bodyUsed".
//
// It returns ok=false if there is no such property.
func (this Request) BodyUsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestBodyUsed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClone returns true if the method "Request.clone" exists.
func (this Request) HasClone() bool {
	return js.True == bindings.HasRequestClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "Request.clone".
func (this Request) CloneFunc() (fn js.Func[func() Request]) {
	return fn.FromRef(
		bindings.RequestCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "Request.clone".
func (this Request) Clone() (ret Request) {
	bindings.CallRequestClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "Request.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryClone() (ret Request, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasArrayBuffer returns true if the method "Request.arrayBuffer" exists.
func (this Request) HasArrayBuffer() bool {
	return js.True == bindings.HasRequestArrayBuffer(
		this.Ref(),
	)
}

// ArrayBufferFunc returns the method "Request.arrayBuffer".
func (this Request) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.RequestArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Request.arrayBuffer".
func (this Request) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallRequestArrayBuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Request.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestArrayBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlob returns true if the method "Request.blob" exists.
func (this Request) HasBlob() bool {
	return js.True == bindings.HasRequestBlob(
		this.Ref(),
	)
}

// BlobFunc returns the method "Request.blob".
func (this Request) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.RequestBlobFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "Request.blob".
func (this Request) Blob() (ret js.Promise[Blob]) {
	bindings.CallRequestBlob(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "Request.blob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestBlob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFormData returns true if the method "Request.formData" exists.
func (this Request) HasFormData() bool {
	return js.True == bindings.HasRequestFormData(
		this.Ref(),
	)
}

// FormDataFunc returns the method "Request.formData".
func (this Request) FormDataFunc() (fn js.Func[func() js.Promise[FormData]]) {
	return fn.FromRef(
		bindings.RequestFormDataFunc(
			this.Ref(),
		),
	)
}

// FormData calls the method "Request.formData".
func (this Request) FormData() (ret js.Promise[FormData]) {
	bindings.CallRequestFormData(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFormData calls the method "Request.formData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryFormData() (ret js.Promise[FormData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestFormData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasJson returns true if the method "Request.json" exists.
func (this Request) HasJson() bool {
	return js.True == bindings.HasRequestJson(
		this.Ref(),
	)
}

// JsonFunc returns the method "Request.json".
func (this Request) JsonFunc() (fn js.Func[func() js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.RequestJsonFunc(
			this.Ref(),
		),
	)
}

// Json calls the method "Request.json".
func (this Request) Json() (ret js.Promise[js.Any]) {
	bindings.CallRequestJson(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryJson calls the method "Request.json"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryJson() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestJson(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasText returns true if the method "Request.text" exists.
func (this Request) HasText() bool {
	return js.True == bindings.HasRequestText(
		this.Ref(),
	)
}

// TextFunc returns the method "Request.text".
func (this Request) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.RequestTextFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "Request.text".
func (this Request) Text() (ret js.Promise[js.String]) {
	bindings.CallRequestText(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Request.text"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Request) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ResponseInit struct {
	// Status is "ResponseInit.status"
	//
	// Optional, defaults to 200.
	//
	// NOTE: FFI_USE_Status MUST be set to true to make this field effective.
	Status uint16
	// StatusText is "ResponseInit.statusText"
	//
	// Optional, defaults to "".
	StatusText js.String
	// Headers is "ResponseInit.headers"
	//
	// Optional
	Headers HeadersInit

	FFI_USE_Status bool // for Status.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResponseInit with all fields set.
func (p ResponseInit) FromRef(ref js.Ref) ResponseInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResponseInit in the application heap.
func (p ResponseInit) New() js.Ref {
	return bindings.ResponseInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ResponseInit) UpdateFrom(ref js.Ref) {
	bindings.ResponseInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ResponseInit) Update(ref js.Ref) {
	bindings.ResponseInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ResponseType uint32

const (
	_ ResponseType = iota

	ResponseType_BASIC
	ResponseType_CORS
	ResponseType_DEFAULT
	ResponseType_ERROR
	ResponseType_OPAQUE
	ResponseType_OPAQUEREDIRECT
)

func (ResponseType) FromRef(str js.Ref) ResponseType {
	return ResponseType(bindings.ConstOfResponseType(str))
}

func (x ResponseType) String() (string, bool) {
	switch x {
	case ResponseType_BASIC:
		return "basic", true
	case ResponseType_CORS:
		return "cors", true
	case ResponseType_DEFAULT:
		return "default", true
	case ResponseType_ERROR:
		return "error", true
	case ResponseType_OPAQUE:
		return "opaque", true
	case ResponseType_OPAQUEREDIRECT:
		return "opaqueredirect", true
	default:
		return "", false
	}
}

func NewResponse(body BodyInit, init ResponseInit) (ret Response) {
	ret.ref = bindings.NewResponseByResponse(
		body.Ref(),
		js.Pointer(&init))
	return
}

func NewResponseByResponse1(body BodyInit) (ret Response) {
	ret.ref = bindings.NewResponseByResponse1(
		body.Ref())
	return
}

func NewResponseByResponse2() (ret Response) {
	ret.ref = bindings.NewResponseByResponse2()
	return
}

type Response struct {
	ref js.Ref
}

func (this Response) Once() Response {
	this.Ref().Once()
	return this
}

func (this Response) Ref() js.Ref {
	return this.ref
}

func (this Response) FromRef(ref js.Ref) Response {
	this.ref = ref
	return this
}

func (this Response) Free() {
	this.Ref().Free()
}

// Type returns the value of property "Response.type".
//
// It returns ok=false if there is no such property.
func (this Response) Type() (ret ResponseType, ok bool) {
	ok = js.True == bindings.GetResponseType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "Response.url".
//
// It returns ok=false if there is no such property.
func (this Response) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetResponseUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Redirected returns the value of property "Response.redirected".
//
// It returns ok=false if there is no such property.
func (this Response) Redirected() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseRedirected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "Response.status".
//
// It returns ok=false if there is no such property.
func (this Response) Status() (ret uint16, ok bool) {
	ok = js.True == bindings.GetResponseStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ok returns the value of property "Response.ok".
//
// It returns ok=false if there is no such property.
func (this Response) Ok() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseOk(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StatusText returns the value of property "Response.statusText".
//
// It returns ok=false if there is no such property.
func (this Response) StatusText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetResponseStatusText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Headers returns the value of property "Response.headers".
//
// It returns ok=false if there is no such property.
func (this Response) Headers() (ret Headers, ok bool) {
	ok = js.True == bindings.GetResponseHeaders(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Response.body".
//
// It returns ok=false if there is no such property.
func (this Response) Body() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetResponseBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BodyUsed returns the value of property "Response.bodyUsed".
//
// It returns ok=false if there is no such property.
func (this Response) BodyUsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseBodyUsed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasError returns true if the staticmethod "Response.error" exists.
func (this Response) HasError() bool {
	return js.True == bindings.HasResponseError(
		this.Ref(),
	)
}

// ErrorFunc returns the staticmethod "Response.error".
func (this Response) ErrorFunc() (fn js.Func[func() Response]) {
	return fn.FromRef(
		bindings.ResponseErrorFunc(
			this.Ref(),
		),
	)
}

// Error calls the staticmethod "Response.error".
func (this Response) Error() (ret Response) {
	bindings.CallResponseError(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryError calls the staticmethod "Response.error"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryError() (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRedirect returns true if the staticmethod "Response.redirect" exists.
func (this Response) HasRedirect() bool {
	return js.True == bindings.HasResponseRedirect(
		this.Ref(),
	)
}

// RedirectFunc returns the staticmethod "Response.redirect".
func (this Response) RedirectFunc() (fn js.Func[func(url js.String, status uint16) Response]) {
	return fn.FromRef(
		bindings.ResponseRedirectFunc(
			this.Ref(),
		),
	)
}

// Redirect calls the staticmethod "Response.redirect".
func (this Response) Redirect(url js.String, status uint16) (ret Response) {
	bindings.CallResponseRedirect(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
		uint32(status),
	)

	return
}

// TryRedirect calls the staticmethod "Response.redirect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryRedirect(url js.String, status uint16) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseRedirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		uint32(status),
	)

	return
}

// HasRedirect1 returns true if the staticmethod "Response.redirect" exists.
func (this Response) HasRedirect1() bool {
	return js.True == bindings.HasResponseRedirect1(
		this.Ref(),
	)
}

// Redirect1Func returns the staticmethod "Response.redirect".
func (this Response) Redirect1Func() (fn js.Func[func(url js.String) Response]) {
	return fn.FromRef(
		bindings.ResponseRedirect1Func(
			this.Ref(),
		),
	)
}

// Redirect1 calls the staticmethod "Response.redirect".
func (this Response) Redirect1(url js.String) (ret Response) {
	bindings.CallResponseRedirect1(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryRedirect1 calls the staticmethod "Response.redirect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryRedirect1(url js.String) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseRedirect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasJson returns true if the staticmethod "Response.json" exists.
func (this Response) HasJson() bool {
	return js.True == bindings.HasResponseJson(
		this.Ref(),
	)
}

// JsonFunc returns the staticmethod "Response.json".
func (this Response) JsonFunc() (fn js.Func[func(data js.Any, init ResponseInit) Response]) {
	return fn.FromRef(
		bindings.ResponseJsonFunc(
			this.Ref(),
		),
	)
}

// Json calls the staticmethod "Response.json".
func (this Response) Json(data js.Any, init ResponseInit) (ret Response) {
	bindings.CallResponseJson(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryJson calls the staticmethod "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryJson(data js.Any, init ResponseInit) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasJson1 returns true if the staticmethod "Response.json" exists.
func (this Response) HasJson1() bool {
	return js.True == bindings.HasResponseJson1(
		this.Ref(),
	)
}

// Json1Func returns the staticmethod "Response.json".
func (this Response) Json1Func() (fn js.Func[func(data js.Any) Response]) {
	return fn.FromRef(
		bindings.ResponseJson1Func(
			this.Ref(),
		),
	)
}

// Json1 calls the staticmethod "Response.json".
func (this Response) Json1(data js.Any) (ret Response) {
	bindings.CallResponseJson1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryJson1 calls the staticmethod "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryJson1(data js.Any) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasClone returns true if the method "Response.clone" exists.
func (this Response) HasClone() bool {
	return js.True == bindings.HasResponseClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "Response.clone".
func (this Response) CloneFunc() (fn js.Func[func() Response]) {
	return fn.FromRef(
		bindings.ResponseCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "Response.clone".
func (this Response) Clone() (ret Response) {
	bindings.CallResponseClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "Response.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryClone() (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasArrayBuffer returns true if the method "Response.arrayBuffer" exists.
func (this Response) HasArrayBuffer() bool {
	return js.True == bindings.HasResponseArrayBuffer(
		this.Ref(),
	)
}

// ArrayBufferFunc returns the method "Response.arrayBuffer".
func (this Response) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.ResponseArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Response.arrayBuffer".
func (this Response) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallResponseArrayBuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Response.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseArrayBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlob returns true if the method "Response.blob" exists.
func (this Response) HasBlob() bool {
	return js.True == bindings.HasResponseBlob(
		this.Ref(),
	)
}

// BlobFunc returns the method "Response.blob".
func (this Response) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ResponseBlobFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "Response.blob".
func (this Response) Blob() (ret js.Promise[Blob]) {
	bindings.CallResponseBlob(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "Response.blob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseBlob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFormData returns true if the method "Response.formData" exists.
func (this Response) HasFormData() bool {
	return js.True == bindings.HasResponseFormData(
		this.Ref(),
	)
}

// FormDataFunc returns the method "Response.formData".
func (this Response) FormDataFunc() (fn js.Func[func() js.Promise[FormData]]) {
	return fn.FromRef(
		bindings.ResponseFormDataFunc(
			this.Ref(),
		),
	)
}

// FormData calls the method "Response.formData".
func (this Response) FormData() (ret js.Promise[FormData]) {
	bindings.CallResponseFormData(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFormData calls the method "Response.formData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryFormData() (ret js.Promise[FormData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseFormData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasJson2 returns true if the method "Response.json" exists.
func (this Response) HasJson2() bool {
	return js.True == bindings.HasResponseJson2(
		this.Ref(),
	)
}

// Json2Func returns the method "Response.json".
func (this Response) Json2Func() (fn js.Func[func() js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.ResponseJson2Func(
			this.Ref(),
		),
	)
}

// Json2 calls the method "Response.json".
func (this Response) Json2() (ret js.Promise[js.Any]) {
	bindings.CallResponseJson2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryJson2 calls the method "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryJson2() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasText returns true if the method "Response.text" exists.
func (this Response) HasText() bool {
	return js.True == bindings.HasResponseText(
		this.Ref(),
	)
}

// TextFunc returns the method "Response.text".
func (this Response) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.ResponseTextFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "Response.text".
func (this Response) Text() (ret js.Promise[js.String]) {
	bindings.CallResponseText(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Response.text"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Response) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BackgroundFetchRecord struct {
	ref js.Ref
}

func (this BackgroundFetchRecord) Once() BackgroundFetchRecord {
	this.Ref().Once()
	return this
}

func (this BackgroundFetchRecord) Ref() js.Ref {
	return this.ref
}

func (this BackgroundFetchRecord) FromRef(ref js.Ref) BackgroundFetchRecord {
	this.ref = ref
	return this
}

func (this BackgroundFetchRecord) Free() {
	this.Ref().Free()
}

// Request returns the value of property "BackgroundFetchRecord.request".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRecord) Request() (ret Request, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRecordRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseReady returns the value of property "BackgroundFetchRecord.responseReady".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRecord) ResponseReady() (ret js.Promise[Response], ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRecordResponseReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CacheQueryOptions struct {
	// IgnoreSearch is "CacheQueryOptions.ignoreSearch"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreSearch MUST be set to true to make this field effective.
	IgnoreSearch bool
	// IgnoreMethod is "CacheQueryOptions.ignoreMethod"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreMethod MUST be set to true to make this field effective.
	IgnoreMethod bool
	// IgnoreVary is "CacheQueryOptions.ignoreVary"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreVary MUST be set to true to make this field effective.
	IgnoreVary bool

	FFI_USE_IgnoreSearch bool // for IgnoreSearch.
	FFI_USE_IgnoreMethod bool // for IgnoreMethod.
	FFI_USE_IgnoreVary   bool // for IgnoreVary.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CacheQueryOptions with all fields set.
func (p CacheQueryOptions) FromRef(ref js.Ref) CacheQueryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CacheQueryOptions in the application heap.
func (p CacheQueryOptions) New() js.Ref {
	return bindings.CacheQueryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CacheQueryOptions) UpdateFrom(ref js.Ref) {
	bindings.CacheQueryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CacheQueryOptions) Update(ref js.Ref) {
	bindings.CacheQueryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BackgroundFetchResult uint32

const (
	_ BackgroundFetchResult = iota

	BackgroundFetchResult_
	BackgroundFetchResult_SUCCESS
	BackgroundFetchResult_FAILURE
)

func (BackgroundFetchResult) FromRef(str js.Ref) BackgroundFetchResult {
	return BackgroundFetchResult(bindings.ConstOfBackgroundFetchResult(str))
}

func (x BackgroundFetchResult) String() (string, bool) {
	switch x {
	case BackgroundFetchResult_:
		return "", true
	case BackgroundFetchResult_SUCCESS:
		return "success", true
	case BackgroundFetchResult_FAILURE:
		return "failure", true
	default:
		return "", false
	}
}

type BackgroundFetchFailureReason uint32

const (
	_ BackgroundFetchFailureReason = iota

	BackgroundFetchFailureReason_
	BackgroundFetchFailureReason_ABORTED
	BackgroundFetchFailureReason_BAD_STATUS
	BackgroundFetchFailureReason_FETCH_ERROR
	BackgroundFetchFailureReason_QUOTA_EXCEEDED
	BackgroundFetchFailureReason_DOWNLOAD_TOTAL_EXCEEDED
)

func (BackgroundFetchFailureReason) FromRef(str js.Ref) BackgroundFetchFailureReason {
	return BackgroundFetchFailureReason(bindings.ConstOfBackgroundFetchFailureReason(str))
}

func (x BackgroundFetchFailureReason) String() (string, bool) {
	switch x {
	case BackgroundFetchFailureReason_:
		return "", true
	case BackgroundFetchFailureReason_ABORTED:
		return "aborted", true
	case BackgroundFetchFailureReason_BAD_STATUS:
		return "bad-status", true
	case BackgroundFetchFailureReason_FETCH_ERROR:
		return "fetch-error", true
	case BackgroundFetchFailureReason_QUOTA_EXCEEDED:
		return "quota-exceeded", true
	case BackgroundFetchFailureReason_DOWNLOAD_TOTAL_EXCEEDED:
		return "download-total-exceeded", true
	default:
		return "", false
	}
}

type BackgroundFetchRegistration struct {
	EventTarget
}

func (this BackgroundFetchRegistration) Once() BackgroundFetchRegistration {
	this.Ref().Once()
	return this
}

func (this BackgroundFetchRegistration) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BackgroundFetchRegistration) FromRef(ref js.Ref) BackgroundFetchRegistration {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BackgroundFetchRegistration) Free() {
	this.Ref().Free()
}

// Id returns the value of property "BackgroundFetchRegistration.id".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UploadTotal returns the value of property "BackgroundFetchRegistration.uploadTotal".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) UploadTotal() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationUploadTotal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Uploaded returns the value of property "BackgroundFetchRegistration.uploaded".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Uploaded() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationUploaded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DownloadTotal returns the value of property "BackgroundFetchRegistration.downloadTotal".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) DownloadTotal() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationDownloadTotal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Downloaded returns the value of property "BackgroundFetchRegistration.downloaded".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Downloaded() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationDownloaded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "BackgroundFetchRegistration.result".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Result() (ret BackgroundFetchResult, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FailureReason returns the value of property "BackgroundFetchRegistration.failureReason".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) FailureReason() (ret BackgroundFetchFailureReason, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationFailureReason(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RecordsAvailable returns the value of property "BackgroundFetchRegistration.recordsAvailable".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) RecordsAvailable() (ret bool, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationRecordsAvailable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAbort returns true if the method "BackgroundFetchRegistration.abort" exists.
func (this BackgroundFetchRegistration) HasAbort() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "BackgroundFetchRegistration.abort".
func (this BackgroundFetchRegistration) AbortFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "BackgroundFetchRegistration.abort".
func (this BackgroundFetchRegistration) Abort() (ret js.Promise[js.Boolean]) {
	bindings.CallBackgroundFetchRegistrationAbort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "BackgroundFetchRegistration.abort"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryAbort() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMatch returns true if the method "BackgroundFetchRegistration.match" exists.
func (this BackgroundFetchRegistration) HasMatch() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationMatch(
		this.Ref(),
	)
}

// MatchFunc returns the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) MatchFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[BackgroundFetchRecord]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchFunc(
			this.Ref(),
		),
	)
}

// Match calls the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) Match(request RequestInfo, options CacheQueryOptions) (ret js.Promise[BackgroundFetchRecord]) {
	bindings.CallBackgroundFetchRegistrationMatch(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatch calls the method "BackgroundFetchRegistration.match"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryMatch(request RequestInfo, options CacheQueryOptions) (ret js.Promise[BackgroundFetchRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMatch1 returns true if the method "BackgroundFetchRegistration.match" exists.
func (this BackgroundFetchRegistration) HasMatch1() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationMatch1(
		this.Ref(),
	)
}

// Match1Func returns the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[BackgroundFetchRecord]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatch1Func(
			this.Ref(),
		),
	)
}

// Match1 calls the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) Match1(request RequestInfo) (ret js.Promise[BackgroundFetchRecord]) {
	bindings.CallBackgroundFetchRegistrationMatch1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatch1 calls the method "BackgroundFetchRegistration.match"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryMatch1(request RequestInfo) (ret js.Promise[BackgroundFetchRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasMatchAll returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasMatchAll() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationMatchAll(
		this.Ref(),
	)
}

// MatchAllFunc returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAllFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatchAll calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMatchAll1 returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasMatchAll1() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationMatchAll1(
		this.Ref(),
	)
}

// MatchAll1Func returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAll1Func(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll1(request RequestInfo) (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatchAll1 calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll1(request RequestInfo) (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasMatchAll2 returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasMatchAll2() bool {
	return js.True == bindings.HasBackgroundFetchRegistrationMatchAll2(
		this.Ref(),
	)
}

// MatchAll2Func returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll2Func() (fn js.Func[func() js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAll2Func(
			this.Ref(),
		),
	)
}

// MatchAll2 calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll2() (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMatchAll2 calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll2() (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BackgroundFetchEventInit struct {
	// Registration is "BackgroundFetchEventInit.registration"
	//
	// Required
	Registration BackgroundFetchRegistration
	// Bubbles is "BackgroundFetchEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "BackgroundFetchEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "BackgroundFetchEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BackgroundFetchEventInit with all fields set.
func (p BackgroundFetchEventInit) FromRef(ref js.Ref) BackgroundFetchEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BackgroundFetchEventInit in the application heap.
func (p BackgroundFetchEventInit) New() js.Ref {
	return bindings.BackgroundFetchEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BackgroundFetchEventInit) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BackgroundFetchEventInit) Update(ref js.Ref) {
	bindings.BackgroundFetchEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewBackgroundFetchEvent(typ js.String, init BackgroundFetchEventInit) (ret BackgroundFetchEvent) {
	ret.ref = bindings.NewBackgroundFetchEventByBackgroundFetchEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
}

type BackgroundFetchEvent struct {
	ExtendableEvent
}

func (this BackgroundFetchEvent) Once() BackgroundFetchEvent {
	this.Ref().Once()
	return this
}

func (this BackgroundFetchEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this BackgroundFetchEvent) FromRef(ref js.Ref) BackgroundFetchEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this BackgroundFetchEvent) Free() {
	this.Ref().Free()
}

// Registration returns the value of property "BackgroundFetchEvent.registration".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchEvent) Registration() (ret BackgroundFetchRegistration, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchEventRegistration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_Request_String_ArrayRequestInfo struct {
	ref js.Ref
}

func (x OneOf_Request_String_ArrayRequestInfo) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Request_String_ArrayRequestInfo) Free() {
	x.ref.Free()
}

func (x OneOf_Request_String_ArrayRequestInfo) FromRef(ref js.Ref) OneOf_Request_String_ArrayRequestInfo {
	return OneOf_Request_String_ArrayRequestInfo{
		ref: ref,
	}
}

func (x OneOf_Request_String_ArrayRequestInfo) Request() Request {
	return Request{}.FromRef(x.ref)
}

func (x OneOf_Request_String_ArrayRequestInfo) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_Request_String_ArrayRequestInfo) ArrayRequestInfo() js.Array[RequestInfo] {
	return js.Array[RequestInfo]{}.FromRef(x.ref)
}

type ImageResource struct {
	// Src is "ImageResource.src"
	//
	// Required
	Src js.String
	// Sizes is "ImageResource.sizes"
	//
	// Optional
	Sizes js.String
	// Type is "ImageResource.type"
	//
	// Optional
	Type js.String
	// Label is "ImageResource.label"
	//
	// Optional
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageResource with all fields set.
func (p ImageResource) FromRef(ref js.Ref) ImageResource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageResource in the application heap.
func (p ImageResource) New() js.Ref {
	return bindings.ImageResourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageResource) UpdateFrom(ref js.Ref) {
	bindings.ImageResourceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageResource) Update(ref js.Ref) {
	bindings.ImageResourceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BackgroundFetchOptions struct {
	// DownloadTotal is "BackgroundFetchOptions.downloadTotal"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DownloadTotal MUST be set to true to make this field effective.
	DownloadTotal uint64
	// Icons is "BackgroundFetchOptions.icons"
	//
	// Optional
	Icons js.Array[ImageResource]
	// Title is "BackgroundFetchOptions.title"
	//
	// Optional
	Title js.String

	FFI_USE_DownloadTotal bool // for DownloadTotal.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BackgroundFetchOptions with all fields set.
func (p BackgroundFetchOptions) FromRef(ref js.Ref) BackgroundFetchOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BackgroundFetchOptions in the application heap.
func (p BackgroundFetchOptions) New() js.Ref {
	return bindings.BackgroundFetchOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BackgroundFetchOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BackgroundFetchOptions) Update(ref js.Ref) {
	bindings.BackgroundFetchOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BackgroundFetchManager struct {
	ref js.Ref
}

func (this BackgroundFetchManager) Once() BackgroundFetchManager {
	this.Ref().Once()
	return this
}

func (this BackgroundFetchManager) Ref() js.Ref {
	return this.ref
}

func (this BackgroundFetchManager) FromRef(ref js.Ref) BackgroundFetchManager {
	this.ref = ref
	return this
}

func (this BackgroundFetchManager) Free() {
	this.Ref().Free()
}

// HasFetch returns true if the method "BackgroundFetchManager.fetch" exists.
func (this BackgroundFetchManager) HasFetch() bool {
	return js.True == bindings.HasBackgroundFetchManagerFetch(
		this.Ref(),
	)
}

// FetchFunc returns the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) FetchFunc() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch calls the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) Fetch(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerFetch(
		this.Ref(), js.Pointer(&ret),
		id.Ref(),
		requests.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryFetch calls the method "BackgroundFetchManager.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchManager) TryFetch(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerFetch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		requests.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFetch1 returns true if the method "BackgroundFetchManager.fetch" exists.
func (this BackgroundFetchManager) HasFetch1() bool {
	return js.True == bindings.HasBackgroundFetchManagerFetch1(
		this.Ref(),
	)
}

// Fetch1Func returns the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) Fetch1Func() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerFetch1Func(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) Fetch1(id js.String, requests OneOf_Request_String_ArrayRequestInfo) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerFetch1(
		this.Ref(), js.Pointer(&ret),
		id.Ref(),
		requests.Ref(),
	)

	return
}

// TryFetch1 calls the method "BackgroundFetchManager.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchManager) TryFetch1(id js.String, requests OneOf_Request_String_ArrayRequestInfo) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerFetch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		requests.Ref(),
	)

	return
}

// HasGet returns true if the method "BackgroundFetchManager.get" exists.
func (this BackgroundFetchManager) HasGet() bool {
	return js.True == bindings.HasBackgroundFetchManagerGet(
		this.Ref(),
	)
}

// GetFunc returns the method "BackgroundFetchManager.get".
func (this BackgroundFetchManager) GetFunc() (fn js.Func[func(id js.String) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "BackgroundFetchManager.get".
func (this BackgroundFetchManager) Get(id js.String) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerGet(
		this.Ref(), js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the method "BackgroundFetchManager.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchManager) TryGet(id js.String) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasGetIds returns true if the method "BackgroundFetchManager.getIds" exists.
func (this BackgroundFetchManager) HasGetIds() bool {
	return js.True == bindings.HasBackgroundFetchManagerGetIds(
		this.Ref(),
	)
}

// GetIdsFunc returns the method "BackgroundFetchManager.getIds".
func (this BackgroundFetchManager) GetIdsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerGetIdsFunc(
			this.Ref(),
		),
	)
}

// GetIds calls the method "BackgroundFetchManager.getIds".
func (this BackgroundFetchManager) GetIds() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallBackgroundFetchManagerGetIds(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetIds calls the method "BackgroundFetchManager.getIds"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchManager) TryGetIds() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerGetIds(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BackgroundFetchUIOptions struct {
	// Icons is "BackgroundFetchUIOptions.icons"
	//
	// Optional
	Icons js.Array[ImageResource]
	// Title is "BackgroundFetchUIOptions.title"
	//
	// Optional
	Title js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BackgroundFetchUIOptions with all fields set.
func (p BackgroundFetchUIOptions) FromRef(ref js.Ref) BackgroundFetchUIOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BackgroundFetchUIOptions in the application heap.
func (p BackgroundFetchUIOptions) New() js.Ref {
	return bindings.BackgroundFetchUIOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BackgroundFetchUIOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchUIOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BackgroundFetchUIOptions) Update(ref js.Ref) {
	bindings.BackgroundFetchUIOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewBackgroundFetchUpdateUIEvent(typ js.String, init BackgroundFetchEventInit) (ret BackgroundFetchUpdateUIEvent) {
	ret.ref = bindings.NewBackgroundFetchUpdateUIEventByBackgroundFetchUpdateUIEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
}

type BackgroundFetchUpdateUIEvent struct {
	BackgroundFetchEvent
}

func (this BackgroundFetchUpdateUIEvent) Once() BackgroundFetchUpdateUIEvent {
	this.Ref().Once()
	return this
}

func (this BackgroundFetchUpdateUIEvent) Ref() js.Ref {
	return this.BackgroundFetchEvent.Ref()
}

func (this BackgroundFetchUpdateUIEvent) FromRef(ref js.Ref) BackgroundFetchUpdateUIEvent {
	this.BackgroundFetchEvent = this.BackgroundFetchEvent.FromRef(ref)
	return this
}

func (this BackgroundFetchUpdateUIEvent) Free() {
	this.Ref().Free()
}

// HasUpdateUI returns true if the method "BackgroundFetchUpdateUIEvent.updateUI" exists.
func (this BackgroundFetchUpdateUIEvent) HasUpdateUI() bool {
	return js.True == bindings.HasBackgroundFetchUpdateUIEventUpdateUI(
		this.Ref(),
	)
}

// UpdateUIFunc returns the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUIFunc() (fn js.Func[func(options BackgroundFetchUIOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BackgroundFetchUpdateUIEventUpdateUIFunc(
			this.Ref(),
		),
	)
}

// UpdateUI calls the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUI(options BackgroundFetchUIOptions) (ret js.Promise[js.Void]) {
	bindings.CallBackgroundFetchUpdateUIEventUpdateUI(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateUI calls the method "BackgroundFetchUpdateUIEvent.updateUI"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchUpdateUIEvent) TryUpdateUI(options BackgroundFetchUIOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchUpdateUIEventUpdateUI(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasUpdateUI1 returns true if the method "BackgroundFetchUpdateUIEvent.updateUI" exists.
func (this BackgroundFetchUpdateUIEvent) HasUpdateUI1() bool {
	return js.True == bindings.HasBackgroundFetchUpdateUIEventUpdateUI1(
		this.Ref(),
	)
}

// UpdateUI1Func returns the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUI1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BackgroundFetchUpdateUIEventUpdateUI1Func(
			this.Ref(),
		),
	)
}

// UpdateUI1 calls the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUI1() (ret js.Promise[js.Void]) {
	bindings.CallBackgroundFetchUpdateUIEventUpdateUI1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUpdateUI1 calls the method "BackgroundFetchUpdateUIEvent.updateUI"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BackgroundFetchUpdateUIEvent) TryUpdateUI1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchUpdateUIEventUpdateUI1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BackgroundSyncOptions struct {
	// MinInterval is "BackgroundSyncOptions.minInterval"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MinInterval MUST be set to true to make this field effective.
	MinInterval uint64

	FFI_USE_MinInterval bool // for MinInterval.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BackgroundSyncOptions with all fields set.
func (p BackgroundSyncOptions) FromRef(ref js.Ref) BackgroundSyncOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BackgroundSyncOptions in the application heap.
func (p BackgroundSyncOptions) New() js.Ref {
	return bindings.BackgroundSyncOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BackgroundSyncOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundSyncOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BackgroundSyncOptions) Update(ref js.Ref) {
	bindings.BackgroundSyncOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BarProp struct {
	ref js.Ref
}

func (this BarProp) Once() BarProp {
	this.Ref().Once()
	return this
}

func (this BarProp) Ref() js.Ref {
	return this.ref
}

func (this BarProp) FromRef(ref js.Ref) BarProp {
	this.ref = ref
	return this
}

func (this BarProp) Free() {
	this.Ref().Free()
}

// Visible returns the value of property "BarProp.visible".
//
// It returns ok=false if there is no such property.
func (this BarProp) Visible() (ret bool, ok bool) {
	ok = js.True == bindings.GetBarPropVisible(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type BarcodeFormat uint32

const (
	_ BarcodeFormat = iota

	BarcodeFormat_AZTEC
	BarcodeFormat_CODE_128
	BarcodeFormat_CODE_39
	BarcodeFormat_CODE_93
	BarcodeFormat_CODABAR
	BarcodeFormat_DATA_MATRIX
	BarcodeFormat_EAN_13
	BarcodeFormat_EAN_8
	BarcodeFormat_ITF
	BarcodeFormat_PDF417
	BarcodeFormat_QR_CODE
	BarcodeFormat_UNKNOWN
	BarcodeFormat_UPC_A
	BarcodeFormat_UPC_E
)

func (BarcodeFormat) FromRef(str js.Ref) BarcodeFormat {
	return BarcodeFormat(bindings.ConstOfBarcodeFormat(str))
}

func (x BarcodeFormat) String() (string, bool) {
	switch x {
	case BarcodeFormat_AZTEC:
		return "aztec", true
	case BarcodeFormat_CODE_128:
		return "code_128", true
	case BarcodeFormat_CODE_39:
		return "code_39", true
	case BarcodeFormat_CODE_93:
		return "code_93", true
	case BarcodeFormat_CODABAR:
		return "codabar", true
	case BarcodeFormat_DATA_MATRIX:
		return "data_matrix", true
	case BarcodeFormat_EAN_13:
		return "ean_13", true
	case BarcodeFormat_EAN_8:
		return "ean_8", true
	case BarcodeFormat_ITF:
		return "itf", true
	case BarcodeFormat_PDF417:
		return "pdf417", true
	case BarcodeFormat_QR_CODE:
		return "qr_code", true
	case BarcodeFormat_UNKNOWN:
		return "unknown", true
	case BarcodeFormat_UPC_A:
		return "upc_a", true
	case BarcodeFormat_UPC_E:
		return "upc_e", true
	default:
		return "", false
	}
}

type BarcodeDetectorOptions struct {
	// Formats is "BarcodeDetectorOptions.formats"
	//
	// Optional
	Formats js.Array[BarcodeFormat]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BarcodeDetectorOptions with all fields set.
func (p BarcodeDetectorOptions) FromRef(ref js.Ref) BarcodeDetectorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BarcodeDetectorOptions in the application heap.
func (p BarcodeDetectorOptions) New() js.Ref {
	return bindings.BarcodeDetectorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BarcodeDetectorOptions) UpdateFrom(ref js.Ref) {
	bindings.BarcodeDetectorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BarcodeDetectorOptions) Update(ref js.Ref) {
	bindings.BarcodeDetectorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DetectedBarcode struct {
	// BoundingBox is "DetectedBarcode.boundingBox"
	//
	// Required
	BoundingBox DOMRectReadOnly
	// RawValue is "DetectedBarcode.rawValue"
	//
	// Required
	RawValue js.String
	// Format is "DetectedBarcode.format"
	//
	// Required
	Format BarcodeFormat
	// CornerPoints is "DetectedBarcode.cornerPoints"
	//
	// Required
	CornerPoints js.FrozenArray[Point2D]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DetectedBarcode with all fields set.
func (p DetectedBarcode) FromRef(ref js.Ref) DetectedBarcode {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DetectedBarcode in the application heap.
func (p DetectedBarcode) New() js.Ref {
	return bindings.DetectedBarcodeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DetectedBarcode) UpdateFrom(ref js.Ref) {
	bindings.DetectedBarcodeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DetectedBarcode) Update(ref js.Ref) {
	bindings.DetectedBarcodeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewHTMLImageElement() (ret HTMLImageElement) {
	ret.ref = bindings.NewHTMLImageElementByHTMLImageElement()
	return
}

type HTMLImageElement struct {
	HTMLElement
}

func (this HTMLImageElement) Once() HTMLImageElement {
	this.Ref().Once()
	return this
}

func (this HTMLImageElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLImageElement) FromRef(ref js.Ref) HTMLImageElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLImageElement) Free() {
	this.Ref().Free()
}

// Alt returns the value of property "HTMLImageElement.alt".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Alt() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAlt(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlt sets the value of property "HTMLImageElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAlt(
		this.Ref(),
		val.Ref(),
	)
}

// Src returns the value of property "HTMLImageElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLImageElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Srcset returns the value of property "HTMLImageElement.srcset".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Srcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSrcset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrcset sets the value of property "HTMLImageElement.srcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSrcset(
		this.Ref(),
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLImageElement.sizes".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Sizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSizes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSizes sets the value of property "HTMLImageElement.sizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSizes(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSizes(
		this.Ref(),
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "HTMLImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLImageElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLImageElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementUseMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLImageElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementUseMap(
		this.Ref(),
		val.Ref(),
	)
}

// IsMap returns the value of property "HTMLImageElement.isMap".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) IsMap() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementIsMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIsMap sets the value of property "HTMLImageElement.isMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetIsMap(val bool) bool {
	return js.True == bindings.SetHTMLImageElementIsMap(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Width returns the value of property "HTMLImageElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLImageElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementWidth(
		this.Ref(),
		uint32(val),
	)
}

// Height returns the value of property "HTMLImageElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLImageElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementHeight(
		this.Ref(),
		uint32(val),
	)
}

// NaturalWidth returns the value of property "HTMLImageElement.naturalWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) NaturalWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementNaturalWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NaturalHeight returns the value of property "HTMLImageElement.naturalHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) NaturalHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementNaturalHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Complete returns the value of property "HTMLImageElement.complete".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Complete() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentSrc returns the value of property "HTMLImageElement.currentSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) CurrentSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementCurrentSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReferrerPolicy returns the value of property "HTMLImageElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLImageElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Decoding returns the value of property "HTMLImageElement.decoding".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Decoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementDecoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDecoding sets the value of property "HTMLImageElement.decoding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetDecoding(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementDecoding(
		this.Ref(),
		val.Ref(),
	)
}

// Loading returns the value of property "HTMLImageElement.loading".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Loading() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLoading(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLoading sets the value of property "HTMLImageElement.loading" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLoading(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLoading(
		this.Ref(),
		val.Ref(),
	)
}

// FetchPriority returns the value of property "HTMLImageElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementFetchPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLImageElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementFetchPriority(
		this.Ref(),
		val.Ref(),
	)
}

// X returns the value of property "HTMLImageElement.x".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) X() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "HTMLImageElement.y".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Y() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLImageElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLImageElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Lowsrc returns the value of property "HTMLImageElement.lowsrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Lowsrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLowsrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLowsrc sets the value of property "HTMLImageElement.lowsrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLowsrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLowsrc(
		this.Ref(),
		val.Ref(),
	)
}

// Align returns the value of property "HTMLImageElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLImageElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Hspace returns the value of property "HTMLImageElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementHspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLImageElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementHspace(
		this.Ref(),
		uint32(val),
	)
}

// Vspace returns the value of property "HTMLImageElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementVspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLImageElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementVspace(
		this.Ref(),
		uint32(val),
	)
}

// LongDesc returns the value of property "HTMLImageElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLongDesc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLImageElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLongDesc(
		this.Ref(),
		val.Ref(),
	)
}

// Border returns the value of property "HTMLImageElement.border".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementBorder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLImageElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementBorder(
		this.Ref(),
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLImageElement.attributionSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) AttributionSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAttributionSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAttributionSrc sets the value of property "HTMLImageElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAttributionSrc(
		this.Ref(),
		val.Ref(),
	)
}

// HasDecode returns true if the method "HTMLImageElement.decode" exists.
func (this HTMLImageElement) HasDecode() bool {
	return js.True == bindings.HasHTMLImageElementDecode(
		this.Ref(),
	)
}

// DecodeFunc returns the method "HTMLImageElement.decode".
func (this HTMLImageElement) DecodeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLImageElementDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "HTMLImageElement.decode".
func (this HTMLImageElement) Decode() (ret js.Promise[js.Void]) {
	bindings.CallHTMLImageElementDecode(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDecode calls the method "HTMLImageElement.decode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLImageElement) TryDecode() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLImageElementDecode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGImageElement struct {
	SVGGraphicsElement
}

func (this SVGImageElement) Once() SVGImageElement {
	this.Ref().Once()
	return this
}

func (this SVGImageElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGImageElement) FromRef(ref js.Ref) SVGImageElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGImageElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGImageElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGImageElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGImageElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGImageElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGImageElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGImageElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CrossOrigin returns the value of property "SVGImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGImageElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "SVGImageElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this SVGImageElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetSVGImageElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// Href returns the value of property "SVGImageElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGImageElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type VideoFrameRequestCallbackFunc func(this js.Ref, now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata) js.Ref

func (fn VideoFrameRequestCallbackFunc) Register() js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)] {
	return js.RegisterCallback[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VideoFrameRequestCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		VideoFrameCallbackMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VideoFrameRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref, now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata) js.Ref
	Arg T
}

func (cb *VideoFrameRequestCallback[T]) Register() js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)] {
	return js.RegisterCallback[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VideoFrameRequestCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		VideoFrameCallbackMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VideoFrameCallbackMetadata struct {
	// PresentationTime is "VideoFrameCallbackMetadata.presentationTime"
	//
	// Required
	PresentationTime DOMHighResTimeStamp
	// ExpectedDisplayTime is "VideoFrameCallbackMetadata.expectedDisplayTime"
	//
	// Required
	ExpectedDisplayTime DOMHighResTimeStamp
	// Width is "VideoFrameCallbackMetadata.width"
	//
	// Required
	Width uint32
	// Height is "VideoFrameCallbackMetadata.height"
	//
	// Required
	Height uint32
	// MediaTime is "VideoFrameCallbackMetadata.mediaTime"
	//
	// Required
	MediaTime float64
	// PresentedFrames is "VideoFrameCallbackMetadata.presentedFrames"
	//
	// Required
	PresentedFrames uint32
	// ProcessingDuration is "VideoFrameCallbackMetadata.processingDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProcessingDuration MUST be set to true to make this field effective.
	ProcessingDuration float64
	// CaptureTime is "VideoFrameCallbackMetadata.captureTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CaptureTime MUST be set to true to make this field effective.
	CaptureTime DOMHighResTimeStamp
	// ReceiveTime is "VideoFrameCallbackMetadata.receiveTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReceiveTime MUST be set to true to make this field effective.
	ReceiveTime DOMHighResTimeStamp
	// RtpTimestamp is "VideoFrameCallbackMetadata.rtpTimestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_RtpTimestamp MUST be set to true to make this field effective.
	RtpTimestamp uint32

	FFI_USE_ProcessingDuration bool // for ProcessingDuration.
	FFI_USE_CaptureTime        bool // for CaptureTime.
	FFI_USE_ReceiveTime        bool // for ReceiveTime.
	FFI_USE_RtpTimestamp       bool // for RtpTimestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoFrameCallbackMetadata with all fields set.
func (p VideoFrameCallbackMetadata) FromRef(ref js.Ref) VideoFrameCallbackMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoFrameCallbackMetadata in the application heap.
func (p VideoFrameCallbackMetadata) New() js.Ref {
	return bindings.VideoFrameCallbackMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoFrameCallbackMetadata) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameCallbackMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoFrameCallbackMetadata) Update(ref js.Ref) {
	bindings.VideoFrameCallbackMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PictureInPictureWindow struct {
	EventTarget
}

func (this PictureInPictureWindow) Once() PictureInPictureWindow {
	this.Ref().Once()
	return this
}

func (this PictureInPictureWindow) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PictureInPictureWindow) FromRef(ref js.Ref) PictureInPictureWindow {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PictureInPictureWindow) Free() {
	this.Ref().Free()
}

// Width returns the value of property "PictureInPictureWindow.width".
//
// It returns ok=false if there is no such property.
func (this PictureInPictureWindow) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetPictureInPictureWindowWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "PictureInPictureWindow.height".
//
// It returns ok=false if there is no such property.
func (this PictureInPictureWindow) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetPictureInPictureWindowHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type VideoPlaybackQuality struct {
	ref js.Ref
}

func (this VideoPlaybackQuality) Once() VideoPlaybackQuality {
	this.Ref().Once()
	return this
}

func (this VideoPlaybackQuality) Ref() js.Ref {
	return this.ref
}

func (this VideoPlaybackQuality) FromRef(ref js.Ref) VideoPlaybackQuality {
	this.ref = ref
	return this
}

func (this VideoPlaybackQuality) Free() {
	this.Ref().Free()
}

// CreationTime returns the value of property "VideoPlaybackQuality.creationTime".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) CreationTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityCreationTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DroppedVideoFrames returns the value of property "VideoPlaybackQuality.droppedVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) DroppedVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityDroppedVideoFrames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TotalVideoFrames returns the value of property "VideoPlaybackQuality.totalVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) TotalVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityTotalVideoFrames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CorruptedVideoFrames returns the value of property "VideoPlaybackQuality.corruptedVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) CorruptedVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityCorruptedVideoFrames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLVideoElement() (ret HTMLVideoElement) {
	ret.ref = bindings.NewHTMLVideoElementByHTMLVideoElement()
	return
}

type HTMLVideoElement struct {
	HTMLMediaElement
}

func (this HTMLVideoElement) Once() HTMLVideoElement {
	this.Ref().Once()
	return this
}

func (this HTMLVideoElement) Ref() js.Ref {
	return this.HTMLMediaElement.Ref()
}

func (this HTMLVideoElement) FromRef(ref js.Ref) HTMLVideoElement {
	this.HTMLMediaElement = this.HTMLMediaElement.FromRef(ref)
	return this
}

func (this HTMLVideoElement) Free() {
	this.Ref().Free()
}

// Width returns the value of property "HTMLVideoElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLVideoElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLVideoElementWidth(
		this.Ref(),
		uint32(val),
	)
}

// Height returns the value of property "HTMLVideoElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLVideoElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLVideoElementHeight(
		this.Ref(),
		uint32(val),
	)
}

// VideoWidth returns the value of property "HTMLVideoElement.videoWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) VideoWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementVideoWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VideoHeight returns the value of property "HTMLVideoElement.videoHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) VideoHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementVideoHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Poster returns the value of property "HTMLVideoElement.poster".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Poster() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementPoster(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPoster sets the value of property "HTMLVideoElement.poster" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetPoster(val js.String) bool {
	return js.True == bindings.SetHTMLVideoElementPoster(
		this.Ref(),
		val.Ref(),
	)
}

// PlaysInline returns the value of property "HTMLVideoElement.playsInline".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) PlaysInline() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementPlaysInline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaysInline sets the value of property "HTMLVideoElement.playsInline" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetPlaysInline(val bool) bool {
	return js.True == bindings.SetHTMLVideoElementPlaysInline(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// DisablePictureInPicture returns the value of property "HTMLVideoElement.disablePictureInPicture".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) DisablePictureInPicture() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementDisablePictureInPicture(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisablePictureInPicture sets the value of property "HTMLVideoElement.disablePictureInPicture" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetDisablePictureInPicture(val bool) bool {
	return js.True == bindings.SetHTMLVideoElementDisablePictureInPicture(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// HasRequestVideoFrameCallback returns true if the method "HTMLVideoElement.requestVideoFrameCallback" exists.
func (this HTMLVideoElement) HasRequestVideoFrameCallback() bool {
	return js.True == bindings.HasHTMLVideoElementRequestVideoFrameCallback(
		this.Ref(),
	)
}

// RequestVideoFrameCallbackFunc returns the method "HTMLVideoElement.requestVideoFrameCallback".
func (this HTMLVideoElement) RequestVideoFrameCallbackFunc() (fn js.Func[func(callback js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)]) uint32]) {
	return fn.FromRef(
		bindings.HTMLVideoElementRequestVideoFrameCallbackFunc(
			this.Ref(),
		),
	)
}

// RequestVideoFrameCallback calls the method "HTMLVideoElement.requestVideoFrameCallback".
func (this HTMLVideoElement) RequestVideoFrameCallback(callback js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)]) (ret uint32) {
	bindings.CallHTMLVideoElementRequestVideoFrameCallback(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestVideoFrameCallback calls the method "HTMLVideoElement.requestVideoFrameCallback"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLVideoElement) TryRequestVideoFrameCallback(callback js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementRequestVideoFrameCallback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCancelVideoFrameCallback returns true if the method "HTMLVideoElement.cancelVideoFrameCallback" exists.
func (this HTMLVideoElement) HasCancelVideoFrameCallback() bool {
	return js.True == bindings.HasHTMLVideoElementCancelVideoFrameCallback(
		this.Ref(),
	)
}

// CancelVideoFrameCallbackFunc returns the method "HTMLVideoElement.cancelVideoFrameCallback".
func (this HTMLVideoElement) CancelVideoFrameCallbackFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.HTMLVideoElementCancelVideoFrameCallbackFunc(
			this.Ref(),
		),
	)
}

// CancelVideoFrameCallback calls the method "HTMLVideoElement.cancelVideoFrameCallback".
func (this HTMLVideoElement) CancelVideoFrameCallback(handle uint32) (ret js.Void) {
	bindings.CallHTMLVideoElementCancelVideoFrameCallback(
		this.Ref(), js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelVideoFrameCallback calls the method "HTMLVideoElement.cancelVideoFrameCallback"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLVideoElement) TryCancelVideoFrameCallback(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementCancelVideoFrameCallback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

// HasRequestPictureInPicture returns true if the method "HTMLVideoElement.requestPictureInPicture" exists.
func (this HTMLVideoElement) HasRequestPictureInPicture() bool {
	return js.True == bindings.HasHTMLVideoElementRequestPictureInPicture(
		this.Ref(),
	)
}

// RequestPictureInPictureFunc returns the method "HTMLVideoElement.requestPictureInPicture".
func (this HTMLVideoElement) RequestPictureInPictureFunc() (fn js.Func[func() js.Promise[PictureInPictureWindow]]) {
	return fn.FromRef(
		bindings.HTMLVideoElementRequestPictureInPictureFunc(
			this.Ref(),
		),
	)
}

// RequestPictureInPicture calls the method "HTMLVideoElement.requestPictureInPicture".
func (this HTMLVideoElement) RequestPictureInPicture() (ret js.Promise[PictureInPictureWindow]) {
	bindings.CallHTMLVideoElementRequestPictureInPicture(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestPictureInPicture calls the method "HTMLVideoElement.requestPictureInPicture"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLVideoElement) TryRequestPictureInPicture() (ret js.Promise[PictureInPictureWindow], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementRequestPictureInPicture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetVideoPlaybackQuality returns true if the method "HTMLVideoElement.getVideoPlaybackQuality" exists.
func (this HTMLVideoElement) HasGetVideoPlaybackQuality() bool {
	return js.True == bindings.HasHTMLVideoElementGetVideoPlaybackQuality(
		this.Ref(),
	)
}

// GetVideoPlaybackQualityFunc returns the method "HTMLVideoElement.getVideoPlaybackQuality".
func (this HTMLVideoElement) GetVideoPlaybackQualityFunc() (fn js.Func[func() VideoPlaybackQuality]) {
	return fn.FromRef(
		bindings.HTMLVideoElementGetVideoPlaybackQualityFunc(
			this.Ref(),
		),
	)
}

// GetVideoPlaybackQuality calls the method "HTMLVideoElement.getVideoPlaybackQuality".
func (this HTMLVideoElement) GetVideoPlaybackQuality() (ret VideoPlaybackQuality) {
	bindings.CallHTMLVideoElementGetVideoPlaybackQuality(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetVideoPlaybackQuality calls the method "HTMLVideoElement.getVideoPlaybackQuality"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLVideoElement) TryGetVideoPlaybackQuality() (ret VideoPlaybackQuality, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementGetVideoPlaybackQuality(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PredefinedColorSpace uint32

const (
	_ PredefinedColorSpace = iota

	PredefinedColorSpace_SRGB
	PredefinedColorSpace_DISPLAY_P3
)

func (PredefinedColorSpace) FromRef(str js.Ref) PredefinedColorSpace {
	return PredefinedColorSpace(bindings.ConstOfPredefinedColorSpace(str))
}

func (x PredefinedColorSpace) String() (string, bool) {
	switch x {
	case PredefinedColorSpace_SRGB:
		return "srgb", true
	case PredefinedColorSpace_DISPLAY_P3:
		return "display-p3", true
	default:
		return "", false
	}
}

type CanvasRenderingContext2DSettings struct {
	// Alpha is "CanvasRenderingContext2DSettings.alpha"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha bool
	// Desynchronized is "CanvasRenderingContext2DSettings.desynchronized"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Desynchronized MUST be set to true to make this field effective.
	Desynchronized bool
	// ColorSpace is "CanvasRenderingContext2DSettings.colorSpace"
	//
	// Optional, defaults to "srgb".
	ColorSpace PredefinedColorSpace
	// WillReadFrequently is "CanvasRenderingContext2DSettings.willReadFrequently"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_WillReadFrequently MUST be set to true to make this field effective.
	WillReadFrequently bool

	FFI_USE_Alpha              bool // for Alpha.
	FFI_USE_Desynchronized     bool // for Desynchronized.
	FFI_USE_WillReadFrequently bool // for WillReadFrequently.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CanvasRenderingContext2DSettings with all fields set.
func (p CanvasRenderingContext2DSettings) FromRef(ref js.Ref) CanvasRenderingContext2DSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CanvasRenderingContext2DSettings in the application heap.
func (p CanvasRenderingContext2DSettings) New() js.Ref {
	return bindings.CanvasRenderingContext2DSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CanvasRenderingContext2DSettings) UpdateFrom(ref js.Ref) {
	bindings.CanvasRenderingContext2DSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CanvasRenderingContext2DSettings) Update(ref js.Ref) {
	bindings.CanvasRenderingContext2DSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CanvasGradient struct {
	ref js.Ref
}

func (this CanvasGradient) Once() CanvasGradient {
	this.Ref().Once()
	return this
}

func (this CanvasGradient) Ref() js.Ref {
	return this.ref
}

func (this CanvasGradient) FromRef(ref js.Ref) CanvasGradient {
	this.ref = ref
	return this
}

func (this CanvasGradient) Free() {
	this.Ref().Free()
}

// HasAddColorStop returns true if the method "CanvasGradient.addColorStop" exists.
func (this CanvasGradient) HasAddColorStop() bool {
	return js.True == bindings.HasCanvasGradientAddColorStop(
		this.Ref(),
	)
}

// AddColorStopFunc returns the method "CanvasGradient.addColorStop".
func (this CanvasGradient) AddColorStopFunc() (fn js.Func[func(offset float64, color js.String)]) {
	return fn.FromRef(
		bindings.CanvasGradientAddColorStopFunc(
			this.Ref(),
		),
	)
}

// AddColorStop calls the method "CanvasGradient.addColorStop".
func (this CanvasGradient) AddColorStop(offset float64, color js.String) (ret js.Void) {
	bindings.CallCanvasGradientAddColorStop(
		this.Ref(), js.Pointer(&ret),
		float64(offset),
		color.Ref(),
	)

	return
}

// TryAddColorStop calls the method "CanvasGradient.addColorStop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasGradient) TryAddColorStop(offset float64, color js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasGradientAddColorStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
		color.Ref(),
	)

	return
}

type CanvasPattern struct {
	ref js.Ref
}

func (this CanvasPattern) Once() CanvasPattern {
	this.Ref().Once()
	return this
}

func (this CanvasPattern) Ref() js.Ref {
	return this.ref
}

func (this CanvasPattern) FromRef(ref js.Ref) CanvasPattern {
	this.ref = ref
	return this
}

func (this CanvasPattern) Free() {
	this.Ref().Free()
}

// HasSetTransform returns true if the method "CanvasPattern.setTransform" exists.
func (this CanvasPattern) HasSetTransform() bool {
	return js.True == bindings.HasCanvasPatternSetTransform(
		this.Ref(),
	)
}

// SetTransformFunc returns the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransformFunc() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.CanvasPatternSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransform(transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallCanvasPatternSetTransform(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TrySetTransform calls the method "CanvasPattern.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasPattern) TrySetTransform(transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasPatternSetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasSetTransform1 returns true if the method "CanvasPattern.setTransform" exists.
func (this CanvasPattern) HasSetTransform1() bool {
	return js.True == bindings.HasCanvasPatternSetTransform1(
		this.Ref(),
	)
}

// SetTransform1Func returns the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransform1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasPatternSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransform1() (ret js.Void) {
	bindings.CallCanvasPatternSetTransform1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetTransform1 calls the method "CanvasPattern.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasPattern) TrySetTransform1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasPatternSetTransform1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ImageBitmap struct {
	ref js.Ref
}

func (this ImageBitmap) Once() ImageBitmap {
	this.Ref().Once()
	return this
}

func (this ImageBitmap) Ref() js.Ref {
	return this.ref
}

func (this ImageBitmap) FromRef(ref js.Ref) ImageBitmap {
	this.ref = ref
	return this
}

func (this ImageBitmap) Free() {
	this.Ref().Free()
}

// Width returns the value of property "ImageBitmap.width".
//
// It returns ok=false if there is no such property.
func (this ImageBitmap) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageBitmapWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "ImageBitmap.height".
//
// It returns ok=false if there is no such property.
func (this ImageBitmap) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageBitmapHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClose returns true if the method "ImageBitmap.close" exists.
func (this ImageBitmap) HasClose() bool {
	return js.True == bindings.HasImageBitmapClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "ImageBitmap.close".
func (this ImageBitmap) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageBitmapCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "ImageBitmap.close".
func (this ImageBitmap) Close() (ret js.Void) {
	bindings.CallImageBitmapClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ImageBitmap.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ImageBitmap) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageBitmapClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CanvasFillRule uint32

const (
	_ CanvasFillRule = iota

	CanvasFillRule_NONZERO
	CanvasFillRule_EVENODD
)

func (CanvasFillRule) FromRef(str js.Ref) CanvasFillRule {
	return CanvasFillRule(bindings.ConstOfCanvasFillRule(str))
}

func (x CanvasFillRule) String() (string, bool) {
	switch x {
	case CanvasFillRule_NONZERO:
		return "nonzero", true
	case CanvasFillRule_EVENODD:
		return "evenodd", true
	default:
		return "", false
	}
}

type OneOf_Path2D_String struct {
	ref js.Ref
}

func (x OneOf_Path2D_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Path2D_String) Free() {
	x.ref.Free()
}

func (x OneOf_Path2D_String) FromRef(ref js.Ref) OneOf_Path2D_String {
	return OneOf_Path2D_String{
		ref: ref,
	}
}

func (x OneOf_Path2D_String) Path2D() Path2D {
	return Path2D{}.FromRef(x.ref)
}

func (x OneOf_Path2D_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type OneOf_Float64_DOMPointInit struct {
	ref js.Ref
}

func (x OneOf_Float64_DOMPointInit) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_DOMPointInit) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_DOMPointInit) FromRef(ref js.Ref) OneOf_Float64_DOMPointInit {
	return OneOf_Float64_DOMPointInit{
		ref: ref,
	}
}

func (x OneOf_Float64_DOMPointInit) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_DOMPointInit) DOMPointInit() DOMPointInit {
	var ret DOMPointInit
	ret.UpdateFrom(x.ref)
	return ret
}

type OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit struct {
	ref js.Ref
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) FromRef(ref js.Ref) OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit {
	return OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit{
		ref: ref,
	}
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) DOMPointInit() DOMPointInit {
	var ret DOMPointInit
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) ArrayOneOf_Float64_DOMPointInit() js.Array[OneOf_Float64_DOMPointInit] {
	return js.Array[OneOf_Float64_DOMPointInit]{}.FromRef(x.ref)
}

func NewPath2D(path OneOf_Path2D_String) (ret Path2D) {
	ret.ref = bindings.NewPath2DByPath2D(
		path.Ref())
	return
}

func NewPath2DByPath2D1() (ret Path2D) {
	ret.ref = bindings.NewPath2DByPath2D1()
	return
}

type Path2D struct {
	ref js.Ref
}

func (this Path2D) Once() Path2D {
	this.Ref().Once()
	return this
}

func (this Path2D) Ref() js.Ref {
	return this.ref
}

func (this Path2D) FromRef(ref js.Ref) Path2D {
	this.ref = ref
	return this
}

func (this Path2D) Free() {
	this.Ref().Free()
}

// HasAddPath returns true if the method "Path2D.addPath" exists.
func (this Path2D) HasAddPath() bool {
	return js.True == bindings.HasPath2DAddPath(
		this.Ref(),
	)
}

// AddPathFunc returns the method "Path2D.addPath".
func (this Path2D) AddPathFunc() (fn js.Func[func(path Path2D, transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.Path2DAddPathFunc(
			this.Ref(),
		),
	)
}

// AddPath calls the method "Path2D.addPath".
func (this Path2D) AddPath(path Path2D, transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallPath2DAddPath(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&transform),
	)

	return
}

// TryAddPath calls the method "Path2D.addPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryAddPath(path Path2D, transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DAddPath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&transform),
	)

	return
}

// HasAddPath1 returns true if the method "Path2D.addPath" exists.
func (this Path2D) HasAddPath1() bool {
	return js.True == bindings.HasPath2DAddPath1(
		this.Ref(),
	)
}

// AddPath1Func returns the method "Path2D.addPath".
func (this Path2D) AddPath1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.Path2DAddPath1Func(
			this.Ref(),
		),
	)
}

// AddPath1 calls the method "Path2D.addPath".
func (this Path2D) AddPath1(path Path2D) (ret js.Void) {
	bindings.CallPath2DAddPath1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryAddPath1 calls the method "Path2D.addPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryAddPath1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DAddPath1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasClosePath returns true if the method "Path2D.closePath" exists.
func (this Path2D) HasClosePath() bool {
	return js.True == bindings.HasPath2DClosePath(
		this.Ref(),
	)
}

// ClosePathFunc returns the method "Path2D.closePath".
func (this Path2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.Path2DClosePathFunc(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "Path2D.closePath".
func (this Path2D) ClosePath() (ret js.Void) {
	bindings.CallPath2DClosePath(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClosePath calls the method "Path2D.closePath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryClosePath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DClosePath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMoveTo returns true if the method "Path2D.moveTo" exists.
func (this Path2D) HasMoveTo() bool {
	return js.True == bindings.HasPath2DMoveTo(
		this.Ref(),
	)
}

// MoveToFunc returns the method "Path2D.moveTo".
func (this Path2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DMoveToFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "Path2D.moveTo".
func (this Path2D) MoveTo(x float64, y float64) (ret js.Void) {
	bindings.CallPath2DMoveTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryMoveTo calls the method "Path2D.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryMoveTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DMoveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasLineTo returns true if the method "Path2D.lineTo" exists.
func (this Path2D) HasLineTo() bool {
	return js.True == bindings.HasPath2DLineTo(
		this.Ref(),
	)
}

// LineToFunc returns the method "Path2D.lineTo".
func (this Path2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DLineToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "Path2D.lineTo".
func (this Path2D) LineTo(x float64, y float64) (ret js.Void) {
	bindings.CallPath2DLineTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryLineTo calls the method "Path2D.lineTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryLineTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DLineTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasQuadraticCurveTo returns true if the method "Path2D.quadraticCurveTo" exists.
func (this Path2D) HasQuadraticCurveTo() bool {
	return js.True == bindings.HasPath2DQuadraticCurveTo(
		this.Ref(),
	)
}

// QuadraticCurveToFunc returns the method "Path2D.quadraticCurveTo".
func (this Path2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "Path2D.quadraticCurveTo".
func (this Path2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void) {
	bindings.CallPath2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// TryQuadraticCurveTo calls the method "Path2D.quadraticCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryQuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// HasBezierCurveTo returns true if the method "Path2D.bezierCurveTo" exists.
func (this Path2D) HasBezierCurveTo() bool {
	return js.True == bindings.HasPath2DBezierCurveTo(
		this.Ref(),
	)
}

// BezierCurveToFunc returns the method "Path2D.bezierCurveTo".
func (this Path2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "Path2D.bezierCurveTo".
func (this Path2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void) {
	bindings.CallPath2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// TryBezierCurveTo calls the method "Path2D.bezierCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryBezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// HasArcTo returns true if the method "Path2D.arcTo" exists.
func (this Path2D) HasArcTo() bool {
	return js.True == bindings.HasPath2DArcTo(
		this.Ref(),
	)
}

// ArcToFunc returns the method "Path2D.arcTo".
func (this Path2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.Path2DArcToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "Path2D.arcTo".
func (this Path2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void) {
	bindings.CallPath2DArcTo(
		this.Ref(), js.Pointer(&ret),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// TryArcTo calls the method "Path2D.arcTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArcTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// HasRect returns true if the method "Path2D.rect" exists.
func (this Path2D) HasRect() bool {
	return js.True == bindings.HasPath2DRect(
		this.Ref(),
	)
}

// RectFunc returns the method "Path2D.rect".
func (this Path2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.Path2DRectFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "Path2D.rect".
func (this Path2D) Rect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPath2DRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRect calls the method "Path2D.rect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasRoundRect returns true if the method "Path2D.roundRect" exists.
func (this Path2D) HasRoundRect() bool {
	return js.True == bindings.HasPath2DRoundRect(
		this.Ref(),
	)
}

// RoundRectFunc returns the method "Path2D.roundRect".
func (this Path2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.Path2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "Path2D.roundRect".
func (this Path2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void) {
	bindings.CallPath2DRoundRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// TryRoundRect calls the method "Path2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryRoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRoundRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// HasRoundRect1 returns true if the method "Path2D.roundRect" exists.
func (this Path2D) HasRoundRect1() bool {
	return js.True == bindings.HasPath2DRoundRect1(
		this.Ref(),
	)
}

// RoundRect1Func returns the method "Path2D.roundRect".
func (this Path2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.Path2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "Path2D.roundRect".
func (this Path2D) RoundRect1(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPath2DRoundRect1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRoundRect1 calls the method "Path2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryRoundRect1(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRoundRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasArc returns true if the method "Path2D.arc" exists.
func (this Path2D) HasArc() bool {
	return js.True == bindings.HasPath2DArc(
		this.Ref(),
	)
}

// ArcFunc returns the method "Path2D.arc".
func (this Path2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.Path2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc calls the method "Path2D.arc".
func (this Path2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPath2DArc(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryArc calls the method "Path2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryArc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasArc1 returns true if the method "Path2D.arc" exists.
func (this Path2D) HasArc1() bool {
	return js.True == bindings.HasPath2DArc1(
		this.Ref(),
	)
}

// Arc1Func returns the method "Path2D.arc".
func (this Path2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.Path2DArc1Func(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "Path2D.arc".
func (this Path2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPath2DArc1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryArc1 calls the method "Path2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryArc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArc1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasEllipse returns true if the method "Path2D.ellipse" exists.
func (this Path2D) HasEllipse() bool {
	return js.True == bindings.HasPath2DEllipse(
		this.Ref(),
	)
}

// EllipseFunc returns the method "Path2D.ellipse".
func (this Path2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.Path2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "Path2D.ellipse".
func (this Path2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPath2DEllipse(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryEllipse calls the method "Path2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryEllipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DEllipse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasEllipse1 returns true if the method "Path2D.ellipse" exists.
func (this Path2D) HasEllipse1() bool {
	return js.True == bindings.HasPath2DEllipse1(
		this.Ref(),
	)
}

// Ellipse1Func returns the method "Path2D.ellipse".
func (this Path2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.Path2DEllipse1Func(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "Path2D.ellipse".
func (this Path2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPath2DEllipse1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryEllipse1 calls the method "Path2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Path2D) TryEllipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DEllipse1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

type TextMetrics struct {
	ref js.Ref
}

func (this TextMetrics) Once() TextMetrics {
	this.Ref().Once()
	return this
}

func (this TextMetrics) Ref() js.Ref {
	return this.ref
}

func (this TextMetrics) FromRef(ref js.Ref) TextMetrics {
	this.ref = ref
	return this
}

func (this TextMetrics) Free() {
	this.Ref().Free()
}

// Width returns the value of property "TextMetrics.width".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxLeft returns the value of property "TextMetrics.actualBoundingBoxLeft".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxRight returns the value of property "TextMetrics.actualBoundingBoxRight".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxRight() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxRight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxAscent returns the value of property "TextMetrics.fontBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) FontBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsFontBoundingBoxAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxDescent returns the value of property "TextMetrics.fontBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) FontBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsFontBoundingBoxDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxAscent returns the value of property "TextMetrics.actualBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxDescent returns the value of property "TextMetrics.actualBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EmHeightAscent returns the value of property "TextMetrics.emHeightAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) EmHeightAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsEmHeightAscent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EmHeightDescent returns the value of property "TextMetrics.emHeightDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) EmHeightDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsEmHeightDescent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HangingBaseline returns the value of property "TextMetrics.hangingBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) HangingBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsHangingBaseline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AlphabeticBaseline returns the value of property "TextMetrics.alphabeticBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) AlphabeticBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsAlphabeticBaseline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IdeographicBaseline returns the value of property "TextMetrics.ideographicBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) IdeographicBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsIdeographicBaseline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ImageDataSettings struct {
	// ColorSpace is "ImageDataSettings.colorSpace"
	//
	// Optional
	ColorSpace PredefinedColorSpace

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageDataSettings with all fields set.
func (p ImageDataSettings) FromRef(ref js.Ref) ImageDataSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageDataSettings in the application heap.
func (p ImageDataSettings) New() js.Ref {
	return bindings.ImageDataSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageDataSettings) UpdateFrom(ref js.Ref) {
	bindings.ImageDataSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageDataSettings) Update(ref js.Ref) {
	bindings.ImageDataSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewImageData(sw uint32, sh uint32, settings ImageDataSettings) (ret ImageData) {
	ret.ref = bindings.NewImageDataByImageData(
		uint32(sw),
		uint32(sh),
		js.Pointer(&settings))
	return
}

func NewImageDataByImageData1(sw uint32, sh uint32) (ret ImageData) {
	ret.ref = bindings.NewImageDataByImageData1(
		uint32(sw),
		uint32(sh))
	return
}

func NewImageDataByImageData2(data js.TypedArray[uint8], sw uint32, sh uint32, settings ImageDataSettings) (ret ImageData) {
	ret.ref = bindings.NewImageDataByImageData2(
		data.Ref(),
		uint32(sw),
		uint32(sh),
		js.Pointer(&settings))
	return
}

func NewImageDataByImageData3(data js.TypedArray[uint8], sw uint32, sh uint32) (ret ImageData) {
	ret.ref = bindings.NewImageDataByImageData3(
		data.Ref(),
		uint32(sw),
		uint32(sh))
	return
}

func NewImageDataByImageData4(data js.TypedArray[uint8], sw uint32) (ret ImageData) {
	ret.ref = bindings.NewImageDataByImageData4(
		data.Ref(),
		uint32(sw))
	return
}

type ImageData struct {
	ref js.Ref
}

func (this ImageData) Once() ImageData {
	this.Ref().Once()
	return this
}

func (this ImageData) Ref() js.Ref {
	return this.ref
}

func (this ImageData) FromRef(ref js.Ref) ImageData {
	this.ref = ref
	return this
}

func (this ImageData) Free() {
	this.Ref().Free()
}

// Width returns the value of property "ImageData.width".
//
// It returns ok=false if there is no such property.
func (this ImageData) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageDataWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "ImageData.height".
//
// It returns ok=false if there is no such property.
func (this ImageData) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageDataHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "ImageData.data".
//
// It returns ok=false if there is no such property.
func (this ImageData) Data() (ret js.TypedArray[uint8], ok bool) {
	ok = js.True == bindings.GetImageDataData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColorSpace returns the value of property "ImageData.colorSpace".
//
// It returns ok=false if there is no such property.
func (this ImageData) ColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetImageDataColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ImageSmoothingQuality uint32

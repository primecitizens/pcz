// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// Method returns the value of property "Request.method".
//
// It returns ok=false if there is no such property.
func (this Request) Method() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestMethod(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "Request.url".
//
// It returns ok=false if there is no such property.
func (this Request) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Headers returns the value of property "Request.headers".
//
// It returns ok=false if there is no such property.
func (this Request) Headers() (ret Headers, ok bool) {
	ok = js.True == bindings.GetRequestHeaders(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Destination returns the value of property "Request.destination".
//
// It returns ok=false if there is no such property.
func (this Request) Destination() (ret RequestDestination, ok bool) {
	ok = js.True == bindings.GetRequestDestination(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Referrer returns the value of property "Request.referrer".
//
// It returns ok=false if there is no such property.
func (this Request) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestReferrer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReferrerPolicy returns the value of property "Request.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this Request) ReferrerPolicy() (ret ReferrerPolicy, ok bool) {
	ok = js.True == bindings.GetRequestReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "Request.mode".
//
// It returns ok=false if there is no such property.
func (this Request) Mode() (ret RequestMode, ok bool) {
	ok = js.True == bindings.GetRequestMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Credentials returns the value of property "Request.credentials".
//
// It returns ok=false if there is no such property.
func (this Request) Credentials() (ret RequestCredentials, ok bool) {
	ok = js.True == bindings.GetRequestCredentials(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cache returns the value of property "Request.cache".
//
// It returns ok=false if there is no such property.
func (this Request) Cache() (ret RequestCache, ok bool) {
	ok = js.True == bindings.GetRequestCache(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Redirect returns the value of property "Request.redirect".
//
// It returns ok=false if there is no such property.
func (this Request) Redirect() (ret RequestRedirect, ok bool) {
	ok = js.True == bindings.GetRequestRedirect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Integrity returns the value of property "Request.integrity".
//
// It returns ok=false if there is no such property.
func (this Request) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRequestIntegrity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Keepalive returns the value of property "Request.keepalive".
//
// It returns ok=false if there is no such property.
func (this Request) Keepalive() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestKeepalive(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsReloadNavigation returns the value of property "Request.isReloadNavigation".
//
// It returns ok=false if there is no such property.
func (this Request) IsReloadNavigation() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestIsReloadNavigation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsHistoryNavigation returns the value of property "Request.isHistoryNavigation".
//
// It returns ok=false if there is no such property.
func (this Request) IsHistoryNavigation() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestIsHistoryNavigation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Signal returns the value of property "Request.signal".
//
// It returns ok=false if there is no such property.
func (this Request) Signal() (ret AbortSignal, ok bool) {
	ok = js.True == bindings.GetRequestSignal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duplex returns the value of property "Request.duplex".
//
// It returns ok=false if there is no such property.
func (this Request) Duplex() (ret RequestDuplex, ok bool) {
	ok = js.True == bindings.GetRequestDuplex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Request.body".
//
// It returns ok=false if there is no such property.
func (this Request) Body() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetRequestBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BodyUsed returns the value of property "Request.bodyUsed".
//
// It returns ok=false if there is no such property.
func (this Request) BodyUsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetRequestBodyUsed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClone returns true if the method "Request.clone" exists.
func (this Request) HasFuncClone() bool {
	return js.True == bindings.HasFuncRequestClone(
		this.ref,
	)
}

// FuncClone returns the method "Request.clone".
func (this Request) FuncClone() (fn js.Func[func() Request]) {
	bindings.FuncRequestClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "Request.clone".
func (this Request) Clone() (ret Request) {
	bindings.CallRequestClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "Request.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryClone() (ret Request, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncArrayBuffer returns true if the method "Request.arrayBuffer" exists.
func (this Request) HasFuncArrayBuffer() bool {
	return js.True == bindings.HasFuncRequestArrayBuffer(
		this.ref,
	)
}

// FuncArrayBuffer returns the method "Request.arrayBuffer".
func (this Request) FuncArrayBuffer() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	bindings.FuncRequestArrayBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArrayBuffer calls the method "Request.arrayBuffer".
func (this Request) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallRequestArrayBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Request.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestArrayBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlob returns true if the method "Request.blob" exists.
func (this Request) HasFuncBlob() bool {
	return js.True == bindings.HasFuncRequestBlob(
		this.ref,
	)
}

// FuncBlob returns the method "Request.blob".
func (this Request) FuncBlob() (fn js.Func[func() js.Promise[Blob]]) {
	bindings.FuncRequestBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blob calls the method "Request.blob".
func (this Request) Blob() (ret js.Promise[Blob]) {
	bindings.CallRequestBlob(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "Request.blob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFormData returns true if the method "Request.formData" exists.
func (this Request) HasFuncFormData() bool {
	return js.True == bindings.HasFuncRequestFormData(
		this.ref,
	)
}

// FuncFormData returns the method "Request.formData".
func (this Request) FuncFormData() (fn js.Func[func() js.Promise[FormData]]) {
	bindings.FuncRequestFormData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FormData calls the method "Request.formData".
func (this Request) FormData() (ret js.Promise[FormData]) {
	bindings.CallRequestFormData(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFormData calls the method "Request.formData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryFormData() (ret js.Promise[FormData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestFormData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncJson returns true if the method "Request.json" exists.
func (this Request) HasFuncJson() bool {
	return js.True == bindings.HasFuncRequestJson(
		this.ref,
	)
}

// FuncJson returns the method "Request.json".
func (this Request) FuncJson() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncRequestJson(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Json calls the method "Request.json".
func (this Request) Json() (ret js.Promise[js.Any]) {
	bindings.CallRequestJson(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryJson calls the method "Request.json"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryJson() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestJson(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncText returns true if the method "Request.text" exists.
func (this Request) HasFuncText() bool {
	return js.True == bindings.HasFuncRequestText(
		this.ref,
	)
}

// FuncText returns the method "Request.text".
func (this Request) FuncText() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncRequestText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Text calls the method "Request.text".
func (this Request) Text() (ret js.Promise[js.String]) {
	bindings.CallRequestText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Request.text"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Request) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ResponseInit) UpdateFrom(ref js.Ref) {
	bindings.ResponseInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResponseInit) Update(ref js.Ref) {
	bindings.ResponseInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResponseInit) FreeMembers(recursive bool) {
	js.Free(
		p.StatusText.Ref(),
		p.Headers.Ref(),
	)
	p.StatusText = p.StatusText.FromRef(js.Undefined)
	p.Headers = p.Headers.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "Response.type".
//
// It returns ok=false if there is no such property.
func (this Response) Type() (ret ResponseType, ok bool) {
	ok = js.True == bindings.GetResponseType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "Response.url".
//
// It returns ok=false if there is no such property.
func (this Response) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetResponseUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Redirected returns the value of property "Response.redirected".
//
// It returns ok=false if there is no such property.
func (this Response) Redirected() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseRedirected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "Response.status".
//
// It returns ok=false if there is no such property.
func (this Response) Status() (ret uint16, ok bool) {
	ok = js.True == bindings.GetResponseStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ok returns the value of property "Response.ok".
//
// It returns ok=false if there is no such property.
func (this Response) Ok() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseOk(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StatusText returns the value of property "Response.statusText".
//
// It returns ok=false if there is no such property.
func (this Response) StatusText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetResponseStatusText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Headers returns the value of property "Response.headers".
//
// It returns ok=false if there is no such property.
func (this Response) Headers() (ret Headers, ok bool) {
	ok = js.True == bindings.GetResponseHeaders(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Response.body".
//
// It returns ok=false if there is no such property.
func (this Response) Body() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetResponseBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BodyUsed returns the value of property "Response.bodyUsed".
//
// It returns ok=false if there is no such property.
func (this Response) BodyUsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetResponseBodyUsed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncError returns true if the static method "Response.error" exists.
func (this Response) HasFuncError() bool {
	return js.True == bindings.HasFuncResponseError(
		this.ref,
	)
}

// FuncError returns the static method "Response.error".
func (this Response) FuncError() (fn js.Func[func() Response]) {
	bindings.FuncResponseError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Error calls the static method "Response.error".
func (this Response) Error() (ret Response) {
	bindings.CallResponseError(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryError calls the static method "Response.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryError() (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRedirect returns true if the static method "Response.redirect" exists.
func (this Response) HasFuncRedirect() bool {
	return js.True == bindings.HasFuncResponseRedirect(
		this.ref,
	)
}

// FuncRedirect returns the static method "Response.redirect".
func (this Response) FuncRedirect() (fn js.Func[func(url js.String, status uint16) Response]) {
	bindings.FuncResponseRedirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Redirect calls the static method "Response.redirect".
func (this Response) Redirect(url js.String, status uint16) (ret Response) {
	bindings.CallResponseRedirect(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		uint32(status),
	)

	return
}

// TryRedirect calls the static method "Response.redirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryRedirect(url js.String, status uint16) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseRedirect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		uint32(status),
	)

	return
}

// HasFuncRedirect1 returns true if the static method "Response.redirect" exists.
func (this Response) HasFuncRedirect1() bool {
	return js.True == bindings.HasFuncResponseRedirect1(
		this.ref,
	)
}

// FuncRedirect1 returns the static method "Response.redirect".
func (this Response) FuncRedirect1() (fn js.Func[func(url js.String) Response]) {
	bindings.FuncResponseRedirect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Redirect1 calls the static method "Response.redirect".
func (this Response) Redirect1(url js.String) (ret Response) {
	bindings.CallResponseRedirect1(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryRedirect1 calls the static method "Response.redirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryRedirect1(url js.String) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseRedirect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncJson returns true if the static method "Response.json" exists.
func (this Response) HasFuncJson() bool {
	return js.True == bindings.HasFuncResponseJson(
		this.ref,
	)
}

// FuncJson returns the static method "Response.json".
func (this Response) FuncJson() (fn js.Func[func(data js.Any, init ResponseInit) Response]) {
	bindings.FuncResponseJson(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Json calls the static method "Response.json".
func (this Response) Json(data js.Any, init ResponseInit) (ret Response) {
	bindings.CallResponseJson(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryJson calls the static method "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryJson(data js.Any, init ResponseInit) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncJson1 returns true if the static method "Response.json" exists.
func (this Response) HasFuncJson1() bool {
	return js.True == bindings.HasFuncResponseJson1(
		this.ref,
	)
}

// FuncJson1 returns the static method "Response.json".
func (this Response) FuncJson1() (fn js.Func[func(data js.Any) Response]) {
	bindings.FuncResponseJson1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Json1 calls the static method "Response.json".
func (this Response) Json1(data js.Any) (ret Response) {
	bindings.CallResponseJson1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryJson1 calls the static method "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryJson1(data js.Any) (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncClone returns true if the method "Response.clone" exists.
func (this Response) HasFuncClone() bool {
	return js.True == bindings.HasFuncResponseClone(
		this.ref,
	)
}

// FuncClone returns the method "Response.clone".
func (this Response) FuncClone() (fn js.Func[func() Response]) {
	bindings.FuncResponseClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "Response.clone".
func (this Response) Clone() (ret Response) {
	bindings.CallResponseClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "Response.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryClone() (ret Response, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncArrayBuffer returns true if the method "Response.arrayBuffer" exists.
func (this Response) HasFuncArrayBuffer() bool {
	return js.True == bindings.HasFuncResponseArrayBuffer(
		this.ref,
	)
}

// FuncArrayBuffer returns the method "Response.arrayBuffer".
func (this Response) FuncArrayBuffer() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	bindings.FuncResponseArrayBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArrayBuffer calls the method "Response.arrayBuffer".
func (this Response) ArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallResponseArrayBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "Response.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseArrayBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlob returns true if the method "Response.blob" exists.
func (this Response) HasFuncBlob() bool {
	return js.True == bindings.HasFuncResponseBlob(
		this.ref,
	)
}

// FuncBlob returns the method "Response.blob".
func (this Response) FuncBlob() (fn js.Func[func() js.Promise[Blob]]) {
	bindings.FuncResponseBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blob calls the method "Response.blob".
func (this Response) Blob() (ret js.Promise[Blob]) {
	bindings.CallResponseBlob(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "Response.blob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFormData returns true if the method "Response.formData" exists.
func (this Response) HasFuncFormData() bool {
	return js.True == bindings.HasFuncResponseFormData(
		this.ref,
	)
}

// FuncFormData returns the method "Response.formData".
func (this Response) FuncFormData() (fn js.Func[func() js.Promise[FormData]]) {
	bindings.FuncResponseFormData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FormData calls the method "Response.formData".
func (this Response) FormData() (ret js.Promise[FormData]) {
	bindings.CallResponseFormData(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFormData calls the method "Response.formData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryFormData() (ret js.Promise[FormData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseFormData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncJson2 returns true if the method "Response.json" exists.
func (this Response) HasFuncJson2() bool {
	return js.True == bindings.HasFuncResponseJson2(
		this.ref,
	)
}

// FuncJson2 returns the method "Response.json".
func (this Response) FuncJson2() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncResponseJson2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Json2 calls the method "Response.json".
func (this Response) Json2() (ret js.Promise[js.Any]) {
	bindings.CallResponseJson2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryJson2 calls the method "Response.json"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryJson2() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseJson2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncText returns true if the method "Response.text" exists.
func (this Response) HasFuncText() bool {
	return js.True == bindings.HasFuncResponseText(
		this.ref,
	)
}

// FuncText returns the method "Response.text".
func (this Response) FuncText() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncResponseText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Text calls the method "Response.text".
func (this Response) Text() (ret js.Promise[js.String]) {
	bindings.CallResponseText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryText calls the method "Response.text"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Response) TryText() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResponseText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BackgroundFetchRecord struct {
	ref js.Ref
}

func (this BackgroundFetchRecord) Once() BackgroundFetchRecord {
	this.ref.Once()
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
	this.ref.Free()
}

// Request returns the value of property "BackgroundFetchRecord.request".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRecord) Request() (ret Request, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRecordRequest(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ResponseReady returns the value of property "BackgroundFetchRecord.responseReady".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRecord) ResponseReady() (ret js.Promise[Response], ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRecordResponseReady(
		this.ref, js.Pointer(&ret),
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
func (p *CacheQueryOptions) UpdateFrom(ref js.Ref) {
	bindings.CacheQueryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CacheQueryOptions) Update(ref js.Ref) {
	bindings.CacheQueryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CacheQueryOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "BackgroundFetchRegistration.id".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UploadTotal returns the value of property "BackgroundFetchRegistration.uploadTotal".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) UploadTotal() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationUploadTotal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Uploaded returns the value of property "BackgroundFetchRegistration.uploaded".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Uploaded() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationUploaded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DownloadTotal returns the value of property "BackgroundFetchRegistration.downloadTotal".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) DownloadTotal() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationDownloadTotal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Downloaded returns the value of property "BackgroundFetchRegistration.downloaded".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Downloaded() (ret uint64, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationDownloaded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "BackgroundFetchRegistration.result".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) Result() (ret BackgroundFetchResult, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationResult(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FailureReason returns the value of property "BackgroundFetchRegistration.failureReason".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) FailureReason() (ret BackgroundFetchFailureReason, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationFailureReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RecordsAvailable returns the value of property "BackgroundFetchRegistration.recordsAvailable".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchRegistration) RecordsAvailable() (ret bool, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchRegistrationRecordsAvailable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAbort returns true if the method "BackgroundFetchRegistration.abort" exists.
func (this BackgroundFetchRegistration) HasFuncAbort() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationAbort(
		this.ref,
	)
}

// FuncAbort returns the method "BackgroundFetchRegistration.abort".
func (this BackgroundFetchRegistration) FuncAbort() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncBackgroundFetchRegistrationAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "BackgroundFetchRegistration.abort".
func (this BackgroundFetchRegistration) Abort() (ret js.Promise[js.Boolean]) {
	bindings.CallBackgroundFetchRegistrationAbort(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "BackgroundFetchRegistration.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryAbort() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMatch returns true if the method "BackgroundFetchRegistration.match" exists.
func (this BackgroundFetchRegistration) HasFuncMatch() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationMatch(
		this.ref,
	)
}

// FuncMatch returns the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) FuncMatch() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[BackgroundFetchRecord]]) {
	bindings.FuncBackgroundFetchRegistrationMatch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Match calls the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) Match(request RequestInfo, options CacheQueryOptions) (ret js.Promise[BackgroundFetchRecord]) {
	bindings.CallBackgroundFetchRegistrationMatch(
		this.ref, js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatch calls the method "BackgroundFetchRegistration.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryMatch(request RequestInfo, options CacheQueryOptions) (ret js.Promise[BackgroundFetchRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncMatch1 returns true if the method "BackgroundFetchRegistration.match" exists.
func (this BackgroundFetchRegistration) HasFuncMatch1() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationMatch1(
		this.ref,
	)
}

// FuncMatch1 returns the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) FuncMatch1() (fn js.Func[func(request RequestInfo) js.Promise[BackgroundFetchRecord]]) {
	bindings.FuncBackgroundFetchRegistrationMatch1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Match1 calls the method "BackgroundFetchRegistration.match".
func (this BackgroundFetchRegistration) Match1(request RequestInfo) (ret js.Promise[BackgroundFetchRecord]) {
	bindings.CallBackgroundFetchRegistrationMatch1(
		this.ref, js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatch1 calls the method "BackgroundFetchRegistration.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryMatch1(request RequestInfo) (ret js.Promise[BackgroundFetchRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatch1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasFuncMatchAll returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasFuncMatchAll() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationMatchAll(
		this.ref,
	)
}

// FuncMatchAll returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) FuncMatchAll() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	bindings.FuncBackgroundFetchRegistrationMatchAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchAll calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll(
		this.ref, js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatchAll calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncMatchAll1 returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasFuncMatchAll1() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationMatchAll1(
		this.ref,
	)
}

// FuncMatchAll1 returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) FuncMatchAll1() (fn js.Func[func(request RequestInfo) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	bindings.FuncBackgroundFetchRegistrationMatchAll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchAll1 calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll1(request RequestInfo) (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll1(
		this.ref, js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatchAll1 calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll1(request RequestInfo) (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasFuncMatchAll2 returns true if the method "BackgroundFetchRegistration.matchAll" exists.
func (this BackgroundFetchRegistration) HasFuncMatchAll2() bool {
	return js.True == bindings.HasFuncBackgroundFetchRegistrationMatchAll2(
		this.ref,
	)
}

// FuncMatchAll2 returns the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) FuncMatchAll2() (fn js.Func[func() js.Promise[js.Array[BackgroundFetchRecord]]]) {
	bindings.FuncBackgroundFetchRegistrationMatchAll2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MatchAll2 calls the method "BackgroundFetchRegistration.matchAll".
func (this BackgroundFetchRegistration) MatchAll2() (ret js.Promise[js.Array[BackgroundFetchRecord]]) {
	bindings.CallBackgroundFetchRegistrationMatchAll2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMatchAll2 calls the method "BackgroundFetchRegistration.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchRegistration) TryMatchAll2() (ret js.Promise[js.Array[BackgroundFetchRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchRegistrationMatchAll2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BackgroundFetchEventInit) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BackgroundFetchEventInit) Update(ref js.Ref) {
	bindings.BackgroundFetchEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BackgroundFetchEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Registration.Ref(),
	)
	p.Registration = p.Registration.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Registration returns the value of property "BackgroundFetchEvent.registration".
//
// It returns ok=false if there is no such property.
func (this BackgroundFetchEvent) Registration() (ret BackgroundFetchRegistration, ok bool) {
	ok = js.True == bindings.GetBackgroundFetchEventRegistration(
		this.ref, js.Pointer(&ret),
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
func (p *ImageResource) UpdateFrom(ref js.Ref) {
	bindings.ImageResourceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageResource) Update(ref js.Ref) {
	bindings.ImageResourceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageResource) FreeMembers(recursive bool) {
	js.Free(
		p.Src.Ref(),
		p.Sizes.Ref(),
		p.Type.Ref(),
		p.Label.Ref(),
	)
	p.Src = p.Src.FromRef(js.Undefined)
	p.Sizes = p.Sizes.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
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
func (p *BackgroundFetchOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BackgroundFetchOptions) Update(ref js.Ref) {
	bindings.BackgroundFetchOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BackgroundFetchOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Icons.Ref(),
		p.Title.Ref(),
	)
	p.Icons = p.Icons.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type BackgroundFetchManager struct {
	ref js.Ref
}

func (this BackgroundFetchManager) Once() BackgroundFetchManager {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncFetch returns true if the method "BackgroundFetchManager.fetch" exists.
func (this BackgroundFetchManager) HasFuncFetch() bool {
	return js.True == bindings.HasFuncBackgroundFetchManagerFetch(
		this.ref,
	)
}

// FuncFetch returns the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) FuncFetch() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) js.Promise[BackgroundFetchRegistration]]) {
	bindings.FuncBackgroundFetchManagerFetch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch calls the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) Fetch(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerFetch(
		this.ref, js.Pointer(&ret),
		id.Ref(),
		requests.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryFetch calls the method "BackgroundFetchManager.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchManager) TryFetch(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerFetch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		requests.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncFetch1 returns true if the method "BackgroundFetchManager.fetch" exists.
func (this BackgroundFetchManager) HasFuncFetch1() bool {
	return js.True == bindings.HasFuncBackgroundFetchManagerFetch1(
		this.ref,
	)
}

// FuncFetch1 returns the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) FuncFetch1() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo) js.Promise[BackgroundFetchRegistration]]) {
	bindings.FuncBackgroundFetchManagerFetch1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fetch1 calls the method "BackgroundFetchManager.fetch".
func (this BackgroundFetchManager) Fetch1(id js.String, requests OneOf_Request_String_ArrayRequestInfo) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerFetch1(
		this.ref, js.Pointer(&ret),
		id.Ref(),
		requests.Ref(),
	)

	return
}

// TryFetch1 calls the method "BackgroundFetchManager.fetch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchManager) TryFetch1(id js.String, requests OneOf_Request_String_ArrayRequestInfo) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerFetch1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		requests.Ref(),
	)

	return
}

// HasFuncGet returns true if the method "BackgroundFetchManager.get" exists.
func (this BackgroundFetchManager) HasFuncGet() bool {
	return js.True == bindings.HasFuncBackgroundFetchManagerGet(
		this.ref,
	)
}

// FuncGet returns the method "BackgroundFetchManager.get".
func (this BackgroundFetchManager) FuncGet() (fn js.Func[func(id js.String) js.Promise[BackgroundFetchRegistration]]) {
	bindings.FuncBackgroundFetchManagerGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "BackgroundFetchManager.get".
func (this BackgroundFetchManager) Get(id js.String) (ret js.Promise[BackgroundFetchRegistration]) {
	bindings.CallBackgroundFetchManagerGet(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGet calls the method "BackgroundFetchManager.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchManager) TryGet(id js.String) (ret js.Promise[BackgroundFetchRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetIds returns true if the method "BackgroundFetchManager.getIds" exists.
func (this BackgroundFetchManager) HasFuncGetIds() bool {
	return js.True == bindings.HasFuncBackgroundFetchManagerGetIds(
		this.ref,
	)
}

// FuncGetIds returns the method "BackgroundFetchManager.getIds".
func (this BackgroundFetchManager) FuncGetIds() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	bindings.FuncBackgroundFetchManagerGetIds(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIds calls the method "BackgroundFetchManager.getIds".
func (this BackgroundFetchManager) GetIds() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallBackgroundFetchManagerGetIds(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetIds calls the method "BackgroundFetchManager.getIds"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchManager) TryGetIds() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchManagerGetIds(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BackgroundFetchUIOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundFetchUIOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BackgroundFetchUIOptions) Update(ref js.Ref) {
	bindings.BackgroundFetchUIOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BackgroundFetchUIOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Icons.Ref(),
		p.Title.Ref(),
	)
	p.Icons = p.Icons.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncUpdateUI returns true if the method "BackgroundFetchUpdateUIEvent.updateUI" exists.
func (this BackgroundFetchUpdateUIEvent) HasFuncUpdateUI() bool {
	return js.True == bindings.HasFuncBackgroundFetchUpdateUIEventUpdateUI(
		this.ref,
	)
}

// FuncUpdateUI returns the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) FuncUpdateUI() (fn js.Func[func(options BackgroundFetchUIOptions) js.Promise[js.Void]]) {
	bindings.FuncBackgroundFetchUpdateUIEventUpdateUI(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateUI calls the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUI(options BackgroundFetchUIOptions) (ret js.Promise[js.Void]) {
	bindings.CallBackgroundFetchUpdateUIEventUpdateUI(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateUI calls the method "BackgroundFetchUpdateUIEvent.updateUI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchUpdateUIEvent) TryUpdateUI(options BackgroundFetchUIOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchUpdateUIEventUpdateUI(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncUpdateUI1 returns true if the method "BackgroundFetchUpdateUIEvent.updateUI" exists.
func (this BackgroundFetchUpdateUIEvent) HasFuncUpdateUI1() bool {
	return js.True == bindings.HasFuncBackgroundFetchUpdateUIEventUpdateUI1(
		this.ref,
	)
}

// FuncUpdateUI1 returns the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) FuncUpdateUI1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncBackgroundFetchUpdateUIEventUpdateUI1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateUI1 calls the method "BackgroundFetchUpdateUIEvent.updateUI".
func (this BackgroundFetchUpdateUIEvent) UpdateUI1() (ret js.Promise[js.Void]) {
	bindings.CallBackgroundFetchUpdateUIEventUpdateUI1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUpdateUI1 calls the method "BackgroundFetchUpdateUIEvent.updateUI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BackgroundFetchUpdateUIEvent) TryUpdateUI1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBackgroundFetchUpdateUIEventUpdateUI1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BackgroundSyncOptions) UpdateFrom(ref js.Ref) {
	bindings.BackgroundSyncOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BackgroundSyncOptions) Update(ref js.Ref) {
	bindings.BackgroundSyncOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BackgroundSyncOptions) FreeMembers(recursive bool) {
}

type BarProp struct {
	ref js.Ref
}

func (this BarProp) Once() BarProp {
	this.ref.Once()
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
	this.ref.Free()
}

// Visible returns the value of property "BarProp.visible".
//
// It returns ok=false if there is no such property.
func (this BarProp) Visible() (ret bool, ok bool) {
	ok = js.True == bindings.GetBarPropVisible(
		this.ref, js.Pointer(&ret),
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
func (p *BarcodeDetectorOptions) UpdateFrom(ref js.Ref) {
	bindings.BarcodeDetectorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BarcodeDetectorOptions) Update(ref js.Ref) {
	bindings.BarcodeDetectorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BarcodeDetectorOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Formats.Ref(),
	)
	p.Formats = p.Formats.FromRef(js.Undefined)
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
func (p *DetectedBarcode) UpdateFrom(ref js.Ref) {
	bindings.DetectedBarcodeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DetectedBarcode) Update(ref js.Ref) {
	bindings.DetectedBarcodeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DetectedBarcode) FreeMembers(recursive bool) {
	js.Free(
		p.BoundingBox.Ref(),
		p.RawValue.Ref(),
		p.CornerPoints.Ref(),
	)
	p.BoundingBox = p.BoundingBox.FromRef(js.Undefined)
	p.RawValue = p.RawValue.FromRef(js.Undefined)
	p.CornerPoints = p.CornerPoints.FromRef(js.Undefined)
}

type HTMLImageElement struct {
	HTMLElement
}

func (this HTMLImageElement) Once() HTMLImageElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Alt returns the value of property "HTMLImageElement.alt".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Alt() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAlt(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlt sets the value of property "HTMLImageElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAlt(
		this.ref,
		val.Ref(),
	)
}

// Src returns the value of property "HTMLImageElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLImageElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Srcset returns the value of property "HTMLImageElement.srcset".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Srcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSrcset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrcset sets the value of property "HTMLImageElement.srcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSrcset(
		this.ref,
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLImageElement.sizes".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Sizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementSizes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSizes sets the value of property "HTMLImageElement.sizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetSizes(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementSizes(
		this.ref,
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "HTMLImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLImageElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementCrossOrigin(
		this.ref,
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLImageElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementUseMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLImageElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementUseMap(
		this.ref,
		val.Ref(),
	)
}

// IsMap returns the value of property "HTMLImageElement.isMap".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) IsMap() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementIsMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsMap sets the value of property "HTMLImageElement.isMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetIsMap(val bool) bool {
	return js.True == bindings.SetHTMLImageElementIsMap(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Width returns the value of property "HTMLImageElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLImageElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementWidth(
		this.ref,
		uint32(val),
	)
}

// Height returns the value of property "HTMLImageElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLImageElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementHeight(
		this.ref,
		uint32(val),
	)
}

// NaturalWidth returns the value of property "HTMLImageElement.naturalWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) NaturalWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementNaturalWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NaturalHeight returns the value of property "HTMLImageElement.naturalHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) NaturalHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementNaturalHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Complete returns the value of property "HTMLImageElement.complete".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Complete() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentSrc returns the value of property "HTMLImageElement.currentSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) CurrentSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementCurrentSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReferrerPolicy returns the value of property "HTMLImageElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLImageElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// Decoding returns the value of property "HTMLImageElement.decoding".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Decoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementDecoding(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDecoding sets the value of property "HTMLImageElement.decoding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetDecoding(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementDecoding(
		this.ref,
		val.Ref(),
	)
}

// Loading returns the value of property "HTMLImageElement.loading".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Loading() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLoading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoading sets the value of property "HTMLImageElement.loading" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLoading(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLoading(
		this.ref,
		val.Ref(),
	)
}

// FetchPriority returns the value of property "HTMLImageElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementFetchPriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLImageElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementFetchPriority(
		this.ref,
		val.Ref(),
	)
}

// X returns the value of property "HTMLImageElement.x".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) X() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "HTMLImageElement.y".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Y() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLImageElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLImageElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementName(
		this.ref,
		val.Ref(),
	)
}

// Lowsrc returns the value of property "HTMLImageElement.lowsrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Lowsrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLowsrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLowsrc sets the value of property "HTMLImageElement.lowsrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLowsrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLowsrc(
		this.ref,
		val.Ref(),
	)
}

// Align returns the value of property "HTMLImageElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLImageElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Hspace returns the value of property "HTMLImageElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementHspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLImageElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementHspace(
		this.ref,
		uint32(val),
	)
}

// Vspace returns the value of property "HTMLImageElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementVspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLImageElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLImageElementVspace(
		this.ref,
		uint32(val),
	)
}

// LongDesc returns the value of property "HTMLImageElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementLongDesc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLImageElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementLongDesc(
		this.ref,
		val.Ref(),
	)
}

// Border returns the value of property "HTMLImageElement.border".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementBorder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLImageElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementBorder(
		this.ref,
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLImageElement.attributionSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLImageElement) AttributionSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLImageElementAttributionSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAttributionSrc sets the value of property "HTMLImageElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAttributionSrc(
		this.ref,
		val.Ref(),
	)
}

// HasFuncDecode returns true if the method "HTMLImageElement.decode" exists.
func (this HTMLImageElement) HasFuncDecode() bool {
	return js.True == bindings.HasFuncHTMLImageElementDecode(
		this.ref,
	)
}

// FuncDecode returns the method "HTMLImageElement.decode".
func (this HTMLImageElement) FuncDecode() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHTMLImageElementDecode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Decode calls the method "HTMLImageElement.decode".
func (this HTMLImageElement) Decode() (ret js.Promise[js.Void]) {
	bindings.CallHTMLImageElementDecode(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDecode calls the method "HTMLImageElement.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLImageElement) TryDecode() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLImageElementDecode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SVGImageElement struct {
	SVGGraphicsElement
}

func (this SVGImageElement) Once() SVGImageElement {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "SVGImageElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGImageElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGImageElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGImageElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGImageElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGImageElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGImageElementPreserveAspectRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CrossOrigin returns the value of property "SVGImageElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGImageElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "SVGImageElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this SVGImageElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetSVGImageElementCrossOrigin(
		this.ref,
		val.Ref(),
	)
}

// Href returns the value of property "SVGImageElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGImageElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGImageElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

type VideoFrameRequestCallbackFunc func(this js.Ref, now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata) js.Ref

func (fn VideoFrameRequestCallbackFunc) Register() js.Func[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)] {
	return js.RegisterCallback[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)](
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
	var arg1 VideoFrameCallbackMetadata
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VideoFrameRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref, now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata) js.Ref
	Arg T
}

func (cb *VideoFrameRequestCallback[T]) Register() js.Func[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)] {
	return js.RegisterCallback[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VideoFrameRequestCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 VideoFrameCallbackMetadata
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
		mark.NoEscape(&arg1),
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
func (p *VideoFrameCallbackMetadata) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameCallbackMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoFrameCallbackMetadata) Update(ref js.Ref) {
	bindings.VideoFrameCallbackMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoFrameCallbackMetadata) FreeMembers(recursive bool) {
}

type PictureInPictureWindow struct {
	EventTarget
}

func (this PictureInPictureWindow) Once() PictureInPictureWindow {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "PictureInPictureWindow.width".
//
// It returns ok=false if there is no such property.
func (this PictureInPictureWindow) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetPictureInPictureWindowWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "PictureInPictureWindow.height".
//
// It returns ok=false if there is no such property.
func (this PictureInPictureWindow) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetPictureInPictureWindowHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

type VideoPlaybackQuality struct {
	ref js.Ref
}

func (this VideoPlaybackQuality) Once() VideoPlaybackQuality {
	this.ref.Once()
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
	this.ref.Free()
}

// CreationTime returns the value of property "VideoPlaybackQuality.creationTime".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) CreationTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityCreationTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DroppedVideoFrames returns the value of property "VideoPlaybackQuality.droppedVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) DroppedVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityDroppedVideoFrames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TotalVideoFrames returns the value of property "VideoPlaybackQuality.totalVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) TotalVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityTotalVideoFrames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CorruptedVideoFrames returns the value of property "VideoPlaybackQuality.corruptedVideoFrames".
//
// It returns ok=false if there is no such property.
func (this VideoPlaybackQuality) CorruptedVideoFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoPlaybackQualityCorruptedVideoFrames(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLVideoElement struct {
	HTMLMediaElement
}

func (this HTMLVideoElement) Once() HTMLVideoElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "HTMLVideoElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLVideoElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLVideoElementWidth(
		this.ref,
		uint32(val),
	)
}

// Height returns the value of property "HTMLVideoElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLVideoElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLVideoElementHeight(
		this.ref,
		uint32(val),
	)
}

// VideoWidth returns the value of property "HTMLVideoElement.videoWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) VideoWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementVideoWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VideoHeight returns the value of property "HTMLVideoElement.videoHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) VideoHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementVideoHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Poster returns the value of property "HTMLVideoElement.poster".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) Poster() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementPoster(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPoster sets the value of property "HTMLVideoElement.poster" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetPoster(val js.String) bool {
	return js.True == bindings.SetHTMLVideoElementPoster(
		this.ref,
		val.Ref(),
	)
}

// PlaysInline returns the value of property "HTMLVideoElement.playsInline".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) PlaysInline() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementPlaysInline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPlaysInline sets the value of property "HTMLVideoElement.playsInline" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetPlaysInline(val bool) bool {
	return js.True == bindings.SetHTMLVideoElementPlaysInline(
		this.ref,
		js.Bool(bool(val)),
	)
}

// DisablePictureInPicture returns the value of property "HTMLVideoElement.disablePictureInPicture".
//
// It returns ok=false if there is no such property.
func (this HTMLVideoElement) DisablePictureInPicture() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLVideoElementDisablePictureInPicture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisablePictureInPicture sets the value of property "HTMLVideoElement.disablePictureInPicture" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetDisablePictureInPicture(val bool) bool {
	return js.True == bindings.SetHTMLVideoElementDisablePictureInPicture(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HasFuncRequestVideoFrameCallback returns true if the method "HTMLVideoElement.requestVideoFrameCallback" exists.
func (this HTMLVideoElement) HasFuncRequestVideoFrameCallback() bool {
	return js.True == bindings.HasFuncHTMLVideoElementRequestVideoFrameCallback(
		this.ref,
	)
}

// FuncRequestVideoFrameCallback returns the method "HTMLVideoElement.requestVideoFrameCallback".
func (this HTMLVideoElement) FuncRequestVideoFrameCallback() (fn js.Func[func(callback js.Func[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)]) uint32]) {
	bindings.FuncHTMLVideoElementRequestVideoFrameCallback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestVideoFrameCallback calls the method "HTMLVideoElement.requestVideoFrameCallback".
func (this HTMLVideoElement) RequestVideoFrameCallback(callback js.Func[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)]) (ret uint32) {
	bindings.CallHTMLVideoElementRequestVideoFrameCallback(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestVideoFrameCallback calls the method "HTMLVideoElement.requestVideoFrameCallback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLVideoElement) TryRequestVideoFrameCallback(callback js.Func[func(now DOMHighResTimeStamp, metadata *VideoFrameCallbackMetadata)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementRequestVideoFrameCallback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCancelVideoFrameCallback returns true if the method "HTMLVideoElement.cancelVideoFrameCallback" exists.
func (this HTMLVideoElement) HasFuncCancelVideoFrameCallback() bool {
	return js.True == bindings.HasFuncHTMLVideoElementCancelVideoFrameCallback(
		this.ref,
	)
}

// FuncCancelVideoFrameCallback returns the method "HTMLVideoElement.cancelVideoFrameCallback".
func (this HTMLVideoElement) FuncCancelVideoFrameCallback() (fn js.Func[func(handle uint32)]) {
	bindings.FuncHTMLVideoElementCancelVideoFrameCallback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelVideoFrameCallback calls the method "HTMLVideoElement.cancelVideoFrameCallback".
func (this HTMLVideoElement) CancelVideoFrameCallback(handle uint32) (ret js.Void) {
	bindings.CallHTMLVideoElementCancelVideoFrameCallback(
		this.ref, js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelVideoFrameCallback calls the method "HTMLVideoElement.cancelVideoFrameCallback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLVideoElement) TryCancelVideoFrameCallback(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementCancelVideoFrameCallback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

// HasFuncRequestPictureInPicture returns true if the method "HTMLVideoElement.requestPictureInPicture" exists.
func (this HTMLVideoElement) HasFuncRequestPictureInPicture() bool {
	return js.True == bindings.HasFuncHTMLVideoElementRequestPictureInPicture(
		this.ref,
	)
}

// FuncRequestPictureInPicture returns the method "HTMLVideoElement.requestPictureInPicture".
func (this HTMLVideoElement) FuncRequestPictureInPicture() (fn js.Func[func() js.Promise[PictureInPictureWindow]]) {
	bindings.FuncHTMLVideoElementRequestPictureInPicture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPictureInPicture calls the method "HTMLVideoElement.requestPictureInPicture".
func (this HTMLVideoElement) RequestPictureInPicture() (ret js.Promise[PictureInPictureWindow]) {
	bindings.CallHTMLVideoElementRequestPictureInPicture(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPictureInPicture calls the method "HTMLVideoElement.requestPictureInPicture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLVideoElement) TryRequestPictureInPicture() (ret js.Promise[PictureInPictureWindow], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementRequestPictureInPicture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVideoPlaybackQuality returns true if the method "HTMLVideoElement.getVideoPlaybackQuality" exists.
func (this HTMLVideoElement) HasFuncGetVideoPlaybackQuality() bool {
	return js.True == bindings.HasFuncHTMLVideoElementGetVideoPlaybackQuality(
		this.ref,
	)
}

// FuncGetVideoPlaybackQuality returns the method "HTMLVideoElement.getVideoPlaybackQuality".
func (this HTMLVideoElement) FuncGetVideoPlaybackQuality() (fn js.Func[func() VideoPlaybackQuality]) {
	bindings.FuncHTMLVideoElementGetVideoPlaybackQuality(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetVideoPlaybackQuality calls the method "HTMLVideoElement.getVideoPlaybackQuality".
func (this HTMLVideoElement) GetVideoPlaybackQuality() (ret VideoPlaybackQuality) {
	bindings.CallHTMLVideoElementGetVideoPlaybackQuality(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetVideoPlaybackQuality calls the method "HTMLVideoElement.getVideoPlaybackQuality"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLVideoElement) TryGetVideoPlaybackQuality() (ret VideoPlaybackQuality, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLVideoElementGetVideoPlaybackQuality(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CanvasRenderingContext2DSettings) UpdateFrom(ref js.Ref) {
	bindings.CanvasRenderingContext2DSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CanvasRenderingContext2DSettings) Update(ref js.Ref) {
	bindings.CanvasRenderingContext2DSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CanvasRenderingContext2DSettings) FreeMembers(recursive bool) {
}

type CanvasGradient struct {
	ref js.Ref
}

func (this CanvasGradient) Once() CanvasGradient {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAddColorStop returns true if the method "CanvasGradient.addColorStop" exists.
func (this CanvasGradient) HasFuncAddColorStop() bool {
	return js.True == bindings.HasFuncCanvasGradientAddColorStop(
		this.ref,
	)
}

// FuncAddColorStop returns the method "CanvasGradient.addColorStop".
func (this CanvasGradient) FuncAddColorStop() (fn js.Func[func(offset float64, color js.String)]) {
	bindings.FuncCanvasGradientAddColorStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddColorStop calls the method "CanvasGradient.addColorStop".
func (this CanvasGradient) AddColorStop(offset float64, color js.String) (ret js.Void) {
	bindings.CallCanvasGradientAddColorStop(
		this.ref, js.Pointer(&ret),
		float64(offset),
		color.Ref(),
	)

	return
}

// TryAddColorStop calls the method "CanvasGradient.addColorStop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasGradient) TryAddColorStop(offset float64, color js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasGradientAddColorStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
		color.Ref(),
	)

	return
}

type CanvasPattern struct {
	ref js.Ref
}

func (this CanvasPattern) Once() CanvasPattern {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSetTransform returns true if the method "CanvasPattern.setTransform" exists.
func (this CanvasPattern) HasFuncSetTransform() bool {
	return js.True == bindings.HasFuncCanvasPatternSetTransform(
		this.ref,
	)
}

// FuncSetTransform returns the method "CanvasPattern.setTransform".
func (this CanvasPattern) FuncSetTransform() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	bindings.FuncCanvasPatternSetTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTransform calls the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransform(transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallCanvasPatternSetTransform(
		this.ref, js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TrySetTransform calls the method "CanvasPattern.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasPattern) TrySetTransform(transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasPatternSetTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasFuncSetTransform1 returns true if the method "CanvasPattern.setTransform" exists.
func (this CanvasPattern) HasFuncSetTransform1() bool {
	return js.True == bindings.HasFuncCanvasPatternSetTransform1(
		this.ref,
	)
}

// FuncSetTransform1 returns the method "CanvasPattern.setTransform".
func (this CanvasPattern) FuncSetTransform1() (fn js.Func[func()]) {
	bindings.FuncCanvasPatternSetTransform1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTransform1 calls the method "CanvasPattern.setTransform".
func (this CanvasPattern) SetTransform1() (ret js.Void) {
	bindings.CallCanvasPatternSetTransform1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetTransform1 calls the method "CanvasPattern.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasPattern) TrySetTransform1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasPatternSetTransform1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ImageBitmap struct {
	ref js.Ref
}

func (this ImageBitmap) Once() ImageBitmap {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "ImageBitmap.width".
//
// It returns ok=false if there is no such property.
func (this ImageBitmap) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageBitmapWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "ImageBitmap.height".
//
// It returns ok=false if there is no such property.
func (this ImageBitmap) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageBitmapHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClose returns true if the method "ImageBitmap.close" exists.
func (this ImageBitmap) HasFuncClose() bool {
	return js.True == bindings.HasFuncImageBitmapClose(
		this.ref,
	)
}

// FuncClose returns the method "ImageBitmap.close".
func (this ImageBitmap) FuncClose() (fn js.Func[func()]) {
	bindings.FuncImageBitmapClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "ImageBitmap.close".
func (this ImageBitmap) Close() (ret js.Void) {
	bindings.CallImageBitmapClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "ImageBitmap.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ImageBitmap) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryImageBitmapClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAddPath returns true if the method "Path2D.addPath" exists.
func (this Path2D) HasFuncAddPath() bool {
	return js.True == bindings.HasFuncPath2DAddPath(
		this.ref,
	)
}

// FuncAddPath returns the method "Path2D.addPath".
func (this Path2D) FuncAddPath() (fn js.Func[func(path Path2D, transform DOMMatrix2DInit)]) {
	bindings.FuncPath2DAddPath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddPath calls the method "Path2D.addPath".
func (this Path2D) AddPath(path Path2D, transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallPath2DAddPath(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&transform),
	)

	return
}

// TryAddPath calls the method "Path2D.addPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryAddPath(path Path2D, transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DAddPath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&transform),
	)

	return
}

// HasFuncAddPath1 returns true if the method "Path2D.addPath" exists.
func (this Path2D) HasFuncAddPath1() bool {
	return js.True == bindings.HasFuncPath2DAddPath1(
		this.ref,
	)
}

// FuncAddPath1 returns the method "Path2D.addPath".
func (this Path2D) FuncAddPath1() (fn js.Func[func(path Path2D)]) {
	bindings.FuncPath2DAddPath1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddPath1 calls the method "Path2D.addPath".
func (this Path2D) AddPath1(path Path2D) (ret js.Void) {
	bindings.CallPath2DAddPath1(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryAddPath1 calls the method "Path2D.addPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryAddPath1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DAddPath1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncClosePath returns true if the method "Path2D.closePath" exists.
func (this Path2D) HasFuncClosePath() bool {
	return js.True == bindings.HasFuncPath2DClosePath(
		this.ref,
	)
}

// FuncClosePath returns the method "Path2D.closePath".
func (this Path2D) FuncClosePath() (fn js.Func[func()]) {
	bindings.FuncPath2DClosePath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClosePath calls the method "Path2D.closePath".
func (this Path2D) ClosePath() (ret js.Void) {
	bindings.CallPath2DClosePath(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClosePath calls the method "Path2D.closePath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryClosePath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DClosePath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveTo returns true if the method "Path2D.moveTo" exists.
func (this Path2D) HasFuncMoveTo() bool {
	return js.True == bindings.HasFuncPath2DMoveTo(
		this.ref,
	)
}

// FuncMoveTo returns the method "Path2D.moveTo".
func (this Path2D) FuncMoveTo() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncPath2DMoveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveTo calls the method "Path2D.moveTo".
func (this Path2D) MoveTo(x float64, y float64) (ret js.Void) {
	bindings.CallPath2DMoveTo(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryMoveTo calls the method "Path2D.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryMoveTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DMoveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncLineTo returns true if the method "Path2D.lineTo" exists.
func (this Path2D) HasFuncLineTo() bool {
	return js.True == bindings.HasFuncPath2DLineTo(
		this.ref,
	)
}

// FuncLineTo returns the method "Path2D.lineTo".
func (this Path2D) FuncLineTo() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncPath2DLineTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LineTo calls the method "Path2D.lineTo".
func (this Path2D) LineTo(x float64, y float64) (ret js.Void) {
	bindings.CallPath2DLineTo(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryLineTo calls the method "Path2D.lineTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryLineTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DLineTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncQuadraticCurveTo returns true if the method "Path2D.quadraticCurveTo" exists.
func (this Path2D) HasFuncQuadraticCurveTo() bool {
	return js.True == bindings.HasFuncPath2DQuadraticCurveTo(
		this.ref,
	)
}

// FuncQuadraticCurveTo returns the method "Path2D.quadraticCurveTo".
func (this Path2D) FuncQuadraticCurveTo() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	bindings.FuncPath2DQuadraticCurveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuadraticCurveTo calls the method "Path2D.quadraticCurveTo".
func (this Path2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void) {
	bindings.CallPath2DQuadraticCurveTo(
		this.ref, js.Pointer(&ret),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// TryQuadraticCurveTo calls the method "Path2D.quadraticCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryQuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DQuadraticCurveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncBezierCurveTo returns true if the method "Path2D.bezierCurveTo" exists.
func (this Path2D) HasFuncBezierCurveTo() bool {
	return js.True == bindings.HasFuncPath2DBezierCurveTo(
		this.ref,
	)
}

// FuncBezierCurveTo returns the method "Path2D.bezierCurveTo".
func (this Path2D) FuncBezierCurveTo() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	bindings.FuncPath2DBezierCurveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BezierCurveTo calls the method "Path2D.bezierCurveTo".
func (this Path2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void) {
	bindings.CallPath2DBezierCurveTo(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryBezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DBezierCurveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncArcTo returns true if the method "Path2D.arcTo" exists.
func (this Path2D) HasFuncArcTo() bool {
	return js.True == bindings.HasFuncPath2DArcTo(
		this.ref,
	)
}

// FuncArcTo returns the method "Path2D.arcTo".
func (this Path2D) FuncArcTo() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	bindings.FuncPath2DArcTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArcTo calls the method "Path2D.arcTo".
func (this Path2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void) {
	bindings.CallPath2DArcTo(
		this.ref, js.Pointer(&ret),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// TryArcTo calls the method "Path2D.arcTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArcTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// HasFuncRect returns true if the method "Path2D.rect" exists.
func (this Path2D) HasFuncRect() bool {
	return js.True == bindings.HasFuncPath2DRect(
		this.ref,
	)
}

// FuncRect returns the method "Path2D.rect".
func (this Path2D) FuncRect() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncPath2DRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rect calls the method "Path2D.rect".
func (this Path2D) Rect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPath2DRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRect calls the method "Path2D.rect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncRoundRect returns true if the method "Path2D.roundRect" exists.
func (this Path2D) HasFuncRoundRect() bool {
	return js.True == bindings.HasFuncPath2DRoundRect(
		this.ref,
	)
}

// FuncRoundRect returns the method "Path2D.roundRect".
func (this Path2D) FuncRoundRect() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	bindings.FuncPath2DRoundRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RoundRect calls the method "Path2D.roundRect".
func (this Path2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void) {
	bindings.CallPath2DRoundRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// TryRoundRect calls the method "Path2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryRoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRoundRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// HasFuncRoundRect1 returns true if the method "Path2D.roundRect" exists.
func (this Path2D) HasFuncRoundRect1() bool {
	return js.True == bindings.HasFuncPath2DRoundRect1(
		this.ref,
	)
}

// FuncRoundRect1 returns the method "Path2D.roundRect".
func (this Path2D) FuncRoundRect1() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncPath2DRoundRect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RoundRect1 calls the method "Path2D.roundRect".
func (this Path2D) RoundRect1(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPath2DRoundRect1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRoundRect1 calls the method "Path2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryRoundRect1(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DRoundRect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncArc returns true if the method "Path2D.arc" exists.
func (this Path2D) HasFuncArc() bool {
	return js.True == bindings.HasFuncPath2DArc(
		this.ref,
	)
}

// FuncArc returns the method "Path2D.arc".
func (this Path2D) FuncArc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	bindings.FuncPath2DArc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Arc calls the method "Path2D.arc".
func (this Path2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPath2DArc(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryArc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArc(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasFuncArc1 returns true if the method "Path2D.arc" exists.
func (this Path2D) HasFuncArc1() bool {
	return js.True == bindings.HasFuncPath2DArc1(
		this.ref,
	)
}

// FuncArc1 returns the method "Path2D.arc".
func (this Path2D) FuncArc1() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	bindings.FuncPath2DArc1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Arc1 calls the method "Path2D.arc".
func (this Path2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPath2DArc1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryArc1 calls the method "Path2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryArc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DArc1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasFuncEllipse returns true if the method "Path2D.ellipse" exists.
func (this Path2D) HasFuncEllipse() bool {
	return js.True == bindings.HasFuncPath2DEllipse(
		this.ref,
	)
}

// FuncEllipse returns the method "Path2D.ellipse".
func (this Path2D) FuncEllipse() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	bindings.FuncPath2DEllipse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ellipse calls the method "Path2D.ellipse".
func (this Path2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPath2DEllipse(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryEllipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DEllipse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncEllipse1 returns true if the method "Path2D.ellipse" exists.
func (this Path2D) HasFuncEllipse1() bool {
	return js.True == bindings.HasFuncPath2DEllipse1(
		this.ref,
	)
}

// FuncEllipse1 returns the method "Path2D.ellipse".
func (this Path2D) FuncEllipse1() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	bindings.FuncPath2DEllipse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ellipse1 calls the method "Path2D.ellipse".
func (this Path2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPath2DEllipse1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Path2D) TryEllipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPath2DEllipse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "TextMetrics.width".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxLeft returns the value of property "TextMetrics.actualBoundingBoxLeft".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxRight returns the value of property "TextMetrics.actualBoundingBoxRight".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxRight() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxRight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxAscent returns the value of property "TextMetrics.fontBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) FontBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsFontBoundingBoxAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FontBoundingBoxDescent returns the value of property "TextMetrics.fontBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) FontBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsFontBoundingBoxDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxAscent returns the value of property "TextMetrics.actualBoundingBoxAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActualBoundingBoxDescent returns the value of property "TextMetrics.actualBoundingBoxDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) ActualBoundingBoxDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsActualBoundingBoxDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EmHeightAscent returns the value of property "TextMetrics.emHeightAscent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) EmHeightAscent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsEmHeightAscent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EmHeightDescent returns the value of property "TextMetrics.emHeightDescent".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) EmHeightDescent() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsEmHeightDescent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HangingBaseline returns the value of property "TextMetrics.hangingBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) HangingBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsHangingBaseline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AlphabeticBaseline returns the value of property "TextMetrics.alphabeticBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) AlphabeticBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsAlphabeticBaseline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IdeographicBaseline returns the value of property "TextMetrics.ideographicBaseline".
//
// It returns ok=false if there is no such property.
func (this TextMetrics) IdeographicBaseline() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextMetricsIdeographicBaseline(
		this.ref, js.Pointer(&ret),
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
func (p *ImageDataSettings) UpdateFrom(ref js.Ref) {
	bindings.ImageDataSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageDataSettings) Update(ref js.Ref) {
	bindings.ImageDataSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageDataSettings) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "ImageData.width".
//
// It returns ok=false if there is no such property.
func (this ImageData) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageDataWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "ImageData.height".
//
// It returns ok=false if there is no such property.
func (this ImageData) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetImageDataHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "ImageData.data".
//
// It returns ok=false if there is no such property.
func (this ImageData) Data() (ret js.TypedArray[uint8], ok bool) {
	ok = js.True == bindings.GetImageDataData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColorSpace returns the value of property "ImageData.colorSpace".
//
// It returns ok=false if there is no such property.
func (this ImageData) ColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetImageDataColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

type ImageSmoothingQuality uint32

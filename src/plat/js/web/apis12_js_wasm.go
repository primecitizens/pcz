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

func NewRequest(input RequestInfo, init RequestInit) Request {
	return Request{}.FromRef(
		bindings.NewRequestByRequest(
			input.Ref(),
			js.Pointer(&init)),
	)
}

func NewRequestByRequest1(input RequestInfo) Request {
	return Request{}.FromRef(
		bindings.NewRequestByRequest1(
			input.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this Request) Method() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRequestMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Url returns the value of property "Request.url".
//
// The returned bool will be false if there is no such property.
func (this Request) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRequestUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Headers returns the value of property "Request.headers".
//
// The returned bool will be false if there is no such property.
func (this Request) Headers() (Headers, bool) {
	var _ok bool
	_ret := bindings.GetRequestHeaders(
		this.Ref(), js.Pointer(&_ok),
	)
	return Headers{}.FromRef(_ret), _ok
}

// Destination returns the value of property "Request.destination".
//
// The returned bool will be false if there is no such property.
func (this Request) Destination() (RequestDestination, bool) {
	var _ok bool
	_ret := bindings.GetRequestDestination(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestDestination(_ret), _ok
}

// Referrer returns the value of property "Request.referrer".
//
// The returned bool will be false if there is no such property.
func (this Request) Referrer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRequestReferrer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy returns the value of property "Request.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this Request) ReferrerPolicy() (ReferrerPolicy, bool) {
	var _ok bool
	_ret := bindings.GetRequestReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReferrerPolicy(_ret), _ok
}

// Mode returns the value of property "Request.mode".
//
// The returned bool will be false if there is no such property.
func (this Request) Mode() (RequestMode, bool) {
	var _ok bool
	_ret := bindings.GetRequestMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestMode(_ret), _ok
}

// Credentials returns the value of property "Request.credentials".
//
// The returned bool will be false if there is no such property.
func (this Request) Credentials() (RequestCredentials, bool) {
	var _ok bool
	_ret := bindings.GetRequestCredentials(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestCredentials(_ret), _ok
}

// Cache returns the value of property "Request.cache".
//
// The returned bool will be false if there is no such property.
func (this Request) Cache() (RequestCache, bool) {
	var _ok bool
	_ret := bindings.GetRequestCache(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestCache(_ret), _ok
}

// Redirect returns the value of property "Request.redirect".
//
// The returned bool will be false if there is no such property.
func (this Request) Redirect() (RequestRedirect, bool) {
	var _ok bool
	_ret := bindings.GetRequestRedirect(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestRedirect(_ret), _ok
}

// Integrity returns the value of property "Request.integrity".
//
// The returned bool will be false if there is no such property.
func (this Request) Integrity() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRequestIntegrity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Keepalive returns the value of property "Request.keepalive".
//
// The returned bool will be false if there is no such property.
func (this Request) Keepalive() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRequestKeepalive(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsReloadNavigation returns the value of property "Request.isReloadNavigation".
//
// The returned bool will be false if there is no such property.
func (this Request) IsReloadNavigation() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRequestIsReloadNavigation(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsHistoryNavigation returns the value of property "Request.isHistoryNavigation".
//
// The returned bool will be false if there is no such property.
func (this Request) IsHistoryNavigation() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRequestIsHistoryNavigation(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Signal returns the value of property "Request.signal".
//
// The returned bool will be false if there is no such property.
func (this Request) Signal() (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.GetRequestSignal(
		this.Ref(), js.Pointer(&_ok),
	)
	return AbortSignal{}.FromRef(_ret), _ok
}

// Duplex returns the value of property "Request.duplex".
//
// The returned bool will be false if there is no such property.
func (this Request) Duplex() (RequestDuplex, bool) {
	var _ok bool
	_ret := bindings.GetRequestDuplex(
		this.Ref(), js.Pointer(&_ok),
	)
	return RequestDuplex(_ret), _ok
}

// Body returns the value of property "Request.body".
//
// The returned bool will be false if there is no such property.
func (this Request) Body() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetRequestBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// BodyUsed returns the value of property "Request.bodyUsed".
//
// The returned bool will be false if there is no such property.
func (this Request) BodyUsed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetRequestBodyUsed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Clone calls the method "Request.clone".
//
// The returned bool will be false if there is no such method.
func (this Request) Clone() (Request, bool) {
	var _ok bool
	_ret := bindings.CallRequestClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return Request{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "Request.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) CloneFunc() (fn js.Func[func() Request]) {
	return fn.FromRef(
		bindings.RequestCloneFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Request.arrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this Request) ArrayBuffer() (js.Promise[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallRequestArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// ArrayBufferFunc returns the method "Request.arrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.RequestArrayBufferFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "Request.blob".
//
// The returned bool will be false if there is no such method.
func (this Request) Blob() (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallRequestBlob(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// BlobFunc returns the method "Request.blob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.RequestBlobFunc(
			this.Ref(),
		),
	)
}

// FormData calls the method "Request.formData".
//
// The returned bool will be false if there is no such method.
func (this Request) FormData() (js.Promise[FormData], bool) {
	var _ok bool
	_ret := bindings.CallRequestFormData(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FormData]{}.FromRef(_ret), _ok
}

// FormDataFunc returns the method "Request.formData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) FormDataFunc() (fn js.Func[func() js.Promise[FormData]]) {
	return fn.FromRef(
		bindings.RequestFormDataFunc(
			this.Ref(),
		),
	)
}

// Json calls the method "Request.json".
//
// The returned bool will be false if there is no such method.
func (this Request) Json() (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallRequestJson(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// JsonFunc returns the method "Request.json".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) JsonFunc() (fn js.Func[func() js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.RequestJsonFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "Request.text".
//
// The returned bool will be false if there is no such method.
func (this Request) Text() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallRequestText(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// TextFunc returns the method "Request.text".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Request) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.RequestTextFunc(
			this.Ref(),
		),
	)
}

type ResponseInit struct {
	// Status is "ResponseInit.status"
	//
	// Optional, defaults to 200.
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

// New creates a new {0x140004cc0e0 ResponseInit ResponseInit [// ResponseInit] [0x140002943c0 0x14000294500 0x140002945a0 0x14000294460] 0x1400107e5b8 {0 0}} in the application heap.
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

func NewResponse(body BodyInit, init ResponseInit) Response {
	return Response{}.FromRef(
		bindings.NewResponseByResponse(
			body.Ref(),
			js.Pointer(&init)),
	)
}

func NewResponseByResponse1(body BodyInit) Response {
	return Response{}.FromRef(
		bindings.NewResponseByResponse1(
			body.Ref()),
	)
}

func NewResponseByResponse2() Response {
	return Response{}.FromRef(
		bindings.NewResponseByResponse2(),
	)
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
// The returned bool will be false if there is no such property.
func (this Response) Type() (ResponseType, bool) {
	var _ok bool
	_ret := bindings.GetResponseType(
		this.Ref(), js.Pointer(&_ok),
	)
	return ResponseType(_ret), _ok
}

// Url returns the value of property "Response.url".
//
// The returned bool will be false if there is no such property.
func (this Response) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetResponseUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Redirected returns the value of property "Response.redirected".
//
// The returned bool will be false if there is no such property.
func (this Response) Redirected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetResponseRedirected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Status returns the value of property "Response.status".
//
// The returned bool will be false if there is no such property.
func (this Response) Status() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetResponseStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Ok returns the value of property "Response.ok".
//
// The returned bool will be false if there is no such property.
func (this Response) Ok() (bool, bool) {
	var _ok bool
	_ret := bindings.GetResponseOk(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// StatusText returns the value of property "Response.statusText".
//
// The returned bool will be false if there is no such property.
func (this Response) StatusText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetResponseStatusText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Headers returns the value of property "Response.headers".
//
// The returned bool will be false if there is no such property.
func (this Response) Headers() (Headers, bool) {
	var _ok bool
	_ret := bindings.GetResponseHeaders(
		this.Ref(), js.Pointer(&_ok),
	)
	return Headers{}.FromRef(_ret), _ok
}

// Body returns the value of property "Response.body".
//
// The returned bool will be false if there is no such property.
func (this Response) Body() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetResponseBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// BodyUsed returns the value of property "Response.bodyUsed".
//
// The returned bool will be false if there is no such property.
func (this Response) BodyUsed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetResponseBodyUsed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Error calls the staticmethod "Response.error".
//
// The returned bool will be false if there is no such method.
func (this Response) Error() (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseError(
		this.Ref(), js.Pointer(&_ok),
	)

	return Response{}.FromRef(_ret), _ok
}

// ErrorFunc returns the staticmethod "Response.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) ErrorFunc() (fn js.Func[func() Response]) {
	return fn.FromRef(
		bindings.ResponseErrorFunc(
			this.Ref(),
		),
	)
}

// Redirect calls the staticmethod "Response.redirect".
//
// The returned bool will be false if there is no such method.
func (this Response) Redirect(url js.String, status uint16) (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseRedirect(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		uint32(status),
	)

	return Response{}.FromRef(_ret), _ok
}

// RedirectFunc returns the staticmethod "Response.redirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) RedirectFunc() (fn js.Func[func(url js.String, status uint16) Response]) {
	return fn.FromRef(
		bindings.ResponseRedirectFunc(
			this.Ref(),
		),
	)
}

// Redirect1 calls the staticmethod "Response.redirect".
//
// The returned bool will be false if there is no such method.
func (this Response) Redirect1(url js.String) (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseRedirect1(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return Response{}.FromRef(_ret), _ok
}

// Redirect1Func returns the staticmethod "Response.redirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) Redirect1Func() (fn js.Func[func(url js.String) Response]) {
	return fn.FromRef(
		bindings.ResponseRedirect1Func(
			this.Ref(),
		),
	)
}

// Json calls the staticmethod "Response.json".
//
// The returned bool will be false if there is no such method.
func (this Response) Json(data js.Any, init ResponseInit) (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseJson(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		js.Pointer(&init),
	)

	return Response{}.FromRef(_ret), _ok
}

// JsonFunc returns the staticmethod "Response.json".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) JsonFunc() (fn js.Func[func(data js.Any, init ResponseInit) Response]) {
	return fn.FromRef(
		bindings.ResponseJsonFunc(
			this.Ref(),
		),
	)
}

// Json1 calls the staticmethod "Response.json".
//
// The returned bool will be false if there is no such method.
func (this Response) Json1(data js.Any) (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseJson1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return Response{}.FromRef(_ret), _ok
}

// Json1Func returns the staticmethod "Response.json".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) Json1Func() (fn js.Func[func(data js.Any) Response]) {
	return fn.FromRef(
		bindings.ResponseJson1Func(
			this.Ref(),
		),
	)
}

// Clone calls the method "Response.clone".
//
// The returned bool will be false if there is no such method.
func (this Response) Clone() (Response, bool) {
	var _ok bool
	_ret := bindings.CallResponseClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return Response{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "Response.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) CloneFunc() (fn js.Func[func() Response]) {
	return fn.FromRef(
		bindings.ResponseCloneFunc(
			this.Ref(),
		),
	)
}

// ArrayBuffer calls the method "Response.arrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this Response) ArrayBuffer() (js.Promise[js.ArrayBuffer], bool) {
	var _ok bool
	_ret := bindings.CallResponseArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.ArrayBuffer]{}.FromRef(_ret), _ok
}

// ArrayBufferFunc returns the method "Response.arrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) ArrayBufferFunc() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	return fn.FromRef(
		bindings.ResponseArrayBufferFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "Response.blob".
//
// The returned bool will be false if there is no such method.
func (this Response) Blob() (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallResponseBlob(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// BlobFunc returns the method "Response.blob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.ResponseBlobFunc(
			this.Ref(),
		),
	)
}

// FormData calls the method "Response.formData".
//
// The returned bool will be false if there is no such method.
func (this Response) FormData() (js.Promise[FormData], bool) {
	var _ok bool
	_ret := bindings.CallResponseFormData(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FormData]{}.FromRef(_ret), _ok
}

// FormDataFunc returns the method "Response.formData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) FormDataFunc() (fn js.Func[func() js.Promise[FormData]]) {
	return fn.FromRef(
		bindings.ResponseFormDataFunc(
			this.Ref(),
		),
	)
}

// Json2 calls the method "Response.json".
//
// The returned bool will be false if there is no such method.
func (this Response) Json2() (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallResponseJson2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// Json2Func returns the method "Response.json".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) Json2Func() (fn js.Func[func() js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.ResponseJson2Func(
			this.Ref(),
		),
	)
}

// Text calls the method "Response.text".
//
// The returned bool will be false if there is no such method.
func (this Response) Text() (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallResponseText(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// TextFunc returns the method "Response.text".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Response) TextFunc() (fn js.Func[func() js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.ResponseTextFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRecord) Request() (Request, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRecordRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return Request{}.FromRef(_ret), _ok
}

// ResponseReady returns the value of property "BackgroundFetchRecord.responseReady".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRecord) ResponseReady() (js.Promise[Response], bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRecordResponseReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[Response]{}.FromRef(_ret), _ok
}

type CacheQueryOptions struct {
	// IgnoreSearch is "CacheQueryOptions.ignoreSearch"
	//
	// Optional, defaults to false.
	IgnoreSearch bool
	// IgnoreMethod is "CacheQueryOptions.ignoreMethod"
	//
	// Optional, defaults to false.
	IgnoreMethod bool
	// IgnoreVary is "CacheQueryOptions.ignoreVary"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CacheQueryOptions CacheQueryOptions [// CacheQueryOptions] [0x14000294780 0x140002948c0 0x14000294a00 0x14000294820 0x14000294960 0x14000294aa0] 0x1400107e6c0 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UploadTotal returns the value of property "BackgroundFetchRegistration.uploadTotal".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) UploadTotal() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationUploadTotal(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Uploaded returns the value of property "BackgroundFetchRegistration.uploaded".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) Uploaded() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationUploaded(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DownloadTotal returns the value of property "BackgroundFetchRegistration.downloadTotal".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) DownloadTotal() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationDownloadTotal(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Downloaded returns the value of property "BackgroundFetchRegistration.downloaded".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) Downloaded() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationDownloaded(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Result returns the value of property "BackgroundFetchRegistration.result".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) Result() (BackgroundFetchResult, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return BackgroundFetchResult(_ret), _ok
}

// FailureReason returns the value of property "BackgroundFetchRegistration.failureReason".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) FailureReason() (BackgroundFetchFailureReason, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationFailureReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return BackgroundFetchFailureReason(_ret), _ok
}

// RecordsAvailable returns the value of property "BackgroundFetchRegistration.recordsAvailable".
//
// The returned bool will be false if there is no such property.
func (this BackgroundFetchRegistration) RecordsAvailable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchRegistrationRecordsAvailable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Abort calls the method "BackgroundFetchRegistration.abort".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) Abort() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// AbortFunc returns the method "BackgroundFetchRegistration.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) AbortFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationAbortFunc(
			this.Ref(),
		),
	)
}

// Match calls the method "BackgroundFetchRegistration.match".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) Match(request RequestInfo, options CacheQueryOptions) (js.Promise[BackgroundFetchRecord], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationMatch(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[BackgroundFetchRecord]{}.FromRef(_ret), _ok
}

// MatchFunc returns the method "BackgroundFetchRegistration.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) MatchFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[BackgroundFetchRecord]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchFunc(
			this.Ref(),
		),
	)
}

// Match1 calls the method "BackgroundFetchRegistration.match".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) Match1(request RequestInfo) (js.Promise[BackgroundFetchRecord], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationMatch1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[BackgroundFetchRecord]{}.FromRef(_ret), _ok
}

// Match1Func returns the method "BackgroundFetchRegistration.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[BackgroundFetchRecord]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatch1Func(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "BackgroundFetchRegistration.matchAll".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) MatchAll(request RequestInfo, options CacheQueryOptions) (js.Promise[js.Array[BackgroundFetchRecord]], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationMatchAll(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[BackgroundFetchRecord]]{}.FromRef(_ret), _ok
}

// MatchAllFunc returns the method "BackgroundFetchRegistration.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) MatchAllFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "BackgroundFetchRegistration.matchAll".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) MatchAll1(request RequestInfo) (js.Promise[js.Array[BackgroundFetchRecord]], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationMatchAll1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[js.Array[BackgroundFetchRecord]]{}.FromRef(_ret), _ok
}

// MatchAll1Func returns the method "BackgroundFetchRegistration.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) MatchAll1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAll1Func(
			this.Ref(),
		),
	)
}

// MatchAll2 calls the method "BackgroundFetchRegistration.matchAll".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchRegistration) MatchAll2() (js.Promise[js.Array[BackgroundFetchRecord]], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchRegistrationMatchAll2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BackgroundFetchRecord]]{}.FromRef(_ret), _ok
}

// MatchAll2Func returns the method "BackgroundFetchRegistration.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchRegistration) MatchAll2Func() (fn js.Func[func() js.Promise[js.Array[BackgroundFetchRecord]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchRegistrationMatchAll2Func(
			this.Ref(),
		),
	)
}

type BackgroundFetchEventInit struct {
	// Registration is "BackgroundFetchEventInit.registration"
	//
	// Required
	Registration BackgroundFetchRegistration
	// Bubbles is "BackgroundFetchEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "BackgroundFetchEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "BackgroundFetchEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 BackgroundFetchEventInit BackgroundFetchEventInit [// BackgroundFetchEventInit] [0x14000294be0 0x14000294c80 0x14000294dc0 0x14000294f00 0x14000294d20 0x14000294e60 0x14000294fa0] 0x1400107e150 {0 0}} in the application heap.
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

func NewBackgroundFetchEvent(typ js.String, init BackgroundFetchEventInit) BackgroundFetchEvent {
	return BackgroundFetchEvent{}.FromRef(
		bindings.NewBackgroundFetchEventByBackgroundFetchEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
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
// The returned bool will be false if there is no such property.
func (this BackgroundFetchEvent) Registration() (BackgroundFetchRegistration, bool) {
	var _ok bool
	_ret := bindings.GetBackgroundFetchEventRegistration(
		this.Ref(), js.Pointer(&_ok),
	)
	return BackgroundFetchRegistration{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 ImageResource ImageResource [// ImageResource] [0x14000295180 0x14000295220 0x140002952c0 0x14000295360] 0x1400107e7e0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 BackgroundFetchOptions BackgroundFetchOptions [// BackgroundFetchOptions] [0x14000295040 0x14000295400 0x140002954a0 0x140002950e0] 0x1400107e7b0 {0 0}} in the application heap.
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

// Fetch calls the method "BackgroundFetchManager.fetch".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchManager) Fetch(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) (js.Promise[BackgroundFetchRegistration], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchManagerFetch(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
		requests.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[BackgroundFetchRegistration]{}.FromRef(_ret), _ok
}

// FetchFunc returns the method "BackgroundFetchManager.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchManager) FetchFunc() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo, options BackgroundFetchOptions) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerFetchFunc(
			this.Ref(),
		),
	)
}

// Fetch1 calls the method "BackgroundFetchManager.fetch".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchManager) Fetch1(id js.String, requests OneOf_Request_String_ArrayRequestInfo) (js.Promise[BackgroundFetchRegistration], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchManagerFetch1(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
		requests.Ref(),
	)

	return js.Promise[BackgroundFetchRegistration]{}.FromRef(_ret), _ok
}

// Fetch1Func returns the method "BackgroundFetchManager.fetch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchManager) Fetch1Func() (fn js.Func[func(id js.String, requests OneOf_Request_String_ArrayRequestInfo) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerFetch1Func(
			this.Ref(),
		),
	)
}

// Get calls the method "BackgroundFetchManager.get".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchManager) Get(id js.String) (js.Promise[BackgroundFetchRegistration], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchManagerGet(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return js.Promise[BackgroundFetchRegistration]{}.FromRef(_ret), _ok
}

// GetFunc returns the method "BackgroundFetchManager.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchManager) GetFunc() (fn js.Func[func(id js.String) js.Promise[BackgroundFetchRegistration]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerGetFunc(
			this.Ref(),
		),
	)
}

// GetIds calls the method "BackgroundFetchManager.getIds".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchManager) GetIds() (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchManagerGetIds(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// GetIdsFunc returns the method "BackgroundFetchManager.getIds".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchManager) GetIdsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.BackgroundFetchManagerGetIdsFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 BackgroundFetchUIOptions BackgroundFetchUIOptions [// BackgroundFetchUIOptions] [0x14000295540 0x140002955e0] 0x1400107e840 {0 0}} in the application heap.
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

func NewBackgroundFetchUpdateUIEvent(typ js.String, init BackgroundFetchEventInit) BackgroundFetchUpdateUIEvent {
	return BackgroundFetchUpdateUIEvent{}.FromRef(
		bindings.NewBackgroundFetchUpdateUIEventByBackgroundFetchUpdateUIEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
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

// UpdateUI calls the method "BackgroundFetchUpdateUIEvent.updateUI".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchUpdateUIEvent) UpdateUI(options BackgroundFetchUIOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchUpdateUIEventUpdateUI(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UpdateUIFunc returns the method "BackgroundFetchUpdateUIEvent.updateUI".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchUpdateUIEvent) UpdateUIFunc() (fn js.Func[func(options BackgroundFetchUIOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BackgroundFetchUpdateUIEventUpdateUIFunc(
			this.Ref(),
		),
	)
}

// UpdateUI1 calls the method "BackgroundFetchUpdateUIEvent.updateUI".
//
// The returned bool will be false if there is no such method.
func (this BackgroundFetchUpdateUIEvent) UpdateUI1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBackgroundFetchUpdateUIEventUpdateUI1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UpdateUI1Func returns the method "BackgroundFetchUpdateUIEvent.updateUI".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BackgroundFetchUpdateUIEvent) UpdateUI1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BackgroundFetchUpdateUIEventUpdateUI1Func(
			this.Ref(),
		),
	)
}

type BackgroundSyncOptions struct {
	// MinInterval is "BackgroundSyncOptions.minInterval"
	//
	// Optional, defaults to 0.
	MinInterval uint64

	FFI_USE_MinInterval bool // for MinInterval.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BackgroundSyncOptions with all fields set.
func (p BackgroundSyncOptions) FromRef(ref js.Ref) BackgroundSyncOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 BackgroundSyncOptions BackgroundSyncOptions [// BackgroundSyncOptions] [0x14000295680 0x14000295720] 0x1400107e888 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this BarProp) Visible() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBarPropVisible(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
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

// New creates a new {0x140004cc0e0 BarcodeDetectorOptions BarcodeDetectorOptions [// BarcodeDetectorOptions] [0x140002957c0] 0x1400107e8d0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 DetectedBarcode DetectedBarcode [// DetectedBarcode] [0x14000295860 0x14000295900 0x140002959a0 0x14000295a40] 0x1400107e918 {0 0}} in the application heap.
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

func NewHTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(
		bindings.NewHTMLImageElementByHTMLImageElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Alt() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementAlt(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Alt sets the value of property "HTMLImageElement.alt" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLImageElement.src" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Srcset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementSrcset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Srcset sets the value of property "HTMLImageElement.srcset" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Sizes() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementSizes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sizes sets the value of property "HTMLImageElement.sizes" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "HTMLImageElement.crossOrigin" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) UseMap() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementUseMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UseMap sets the value of property "HTMLImageElement.useMap" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) IsMap() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementIsMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsMap sets the value of property "HTMLImageElement.isMap" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Width sets the value of property "HTMLImageElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height sets the value of property "HTMLImageElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) NaturalWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementNaturalWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NaturalHeight returns the value of property "HTMLImageElement.naturalHeight".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) NaturalHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementNaturalHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Complete returns the value of property "HTMLImageElement.complete".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Complete() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CurrentSrc returns the value of property "HTMLImageElement.currentSrc".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) CurrentSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementCurrentSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy returns the value of property "HTMLImageElement.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLImageElement.referrerPolicy" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Decoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementDecoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Decoding sets the value of property "HTMLImageElement.decoding" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Loading() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementLoading(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Loading sets the value of property "HTMLImageElement.loading" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) FetchPriority() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementFetchPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FetchPriority sets the value of property "HTMLImageElement.fetchPriority" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) X() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Y returns the value of property "HTMLImageElement.y".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Y() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Name returns the value of property "HTMLImageElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLImageElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Lowsrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementLowsrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lowsrc sets the value of property "HTMLImageElement.lowsrc" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLImageElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Hspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementHspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Hspace sets the value of property "HTMLImageElement.hspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Vspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementVspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Vspace sets the value of property "HTMLImageElement.vspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) LongDesc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementLongDesc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LongDesc sets the value of property "HTMLImageElement.longDesc" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) Border() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementBorder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Border sets the value of property "HTMLImageElement.border" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLImageElement) AttributionSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLImageElementAttributionSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttributionSrc sets the value of property "HTMLImageElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLImageElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLImageElementAttributionSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Decode calls the method "HTMLImageElement.decode".
//
// The returned bool will be false if there is no such method.
func (this HTMLImageElement) Decode() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLImageElementDecode(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DecodeFunc returns the method "HTMLImageElement.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLImageElement) DecodeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLImageElementDecodeFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this SVGImageElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGImageElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGImageElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGImageElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGImageElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGImageElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGImageElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGImageElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGImageElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

// CrossOrigin returns the value of property "SVGImageElement.crossOrigin".
//
// The returned bool will be false if there is no such property.
func (this SVGImageElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "SVGImageElement.crossOrigin" to val.
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
// The returned bool will be false if there is no such property.
func (this SVGImageElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGImageElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
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
	ProcessingDuration float64
	// CaptureTime is "VideoFrameCallbackMetadata.captureTime"
	//
	// Optional
	CaptureTime DOMHighResTimeStamp
	// ReceiveTime is "VideoFrameCallbackMetadata.receiveTime"
	//
	// Optional
	ReceiveTime DOMHighResTimeStamp
	// RtpTimestamp is "VideoFrameCallbackMetadata.rtpTimestamp"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 VideoFrameCallbackMetadata VideoFrameCallbackMetadata [// VideoFrameCallbackMetadata] [0x14000295c20 0x14000295cc0 0x14000295d60 0x14000295e00 0x14000295ea0 0x14000295f40 0x140002c2000 0x140002c21e0 0x140002c2320 0x140002c2460 0x140002c20a0 0x140002c2280 0x140002c23c0 0x140002c2500] 0x1400107e960 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this PictureInPictureWindow) Width() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPictureInPictureWindowWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Height returns the value of property "PictureInPictureWindow.height".
//
// The returned bool will be false if there is no such property.
func (this PictureInPictureWindow) Height() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPictureInPictureWindowHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this VideoPlaybackQuality) CreationTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetVideoPlaybackQualityCreationTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DroppedVideoFrames returns the value of property "VideoPlaybackQuality.droppedVideoFrames".
//
// The returned bool will be false if there is no such property.
func (this VideoPlaybackQuality) DroppedVideoFrames() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoPlaybackQualityDroppedVideoFrames(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// TotalVideoFrames returns the value of property "VideoPlaybackQuality.totalVideoFrames".
//
// The returned bool will be false if there is no such property.
func (this VideoPlaybackQuality) TotalVideoFrames() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoPlaybackQualityTotalVideoFrames(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CorruptedVideoFrames returns the value of property "VideoPlaybackQuality.corruptedVideoFrames".
//
// The returned bool will be false if there is no such property.
func (this VideoPlaybackQuality) CorruptedVideoFrames() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoPlaybackQualityCorruptedVideoFrames(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

func NewHTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(
		bindings.NewHTMLVideoElementByHTMLVideoElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Width sets the value of property "HTMLVideoElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height sets the value of property "HTMLVideoElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) VideoWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementVideoWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// VideoHeight returns the value of property "HTMLVideoElement.videoHeight".
//
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) VideoHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementVideoHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Poster returns the value of property "HTMLVideoElement.poster".
//
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) Poster() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementPoster(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Poster sets the value of property "HTMLVideoElement.poster" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) PlaysInline() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementPlaysInline(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// PlaysInline sets the value of property "HTMLVideoElement.playsInline" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLVideoElement) DisablePictureInPicture() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLVideoElementDisablePictureInPicture(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DisablePictureInPicture sets the value of property "HTMLVideoElement.disablePictureInPicture" to val.
//
// It returns false if the property cannot be set.
func (this HTMLVideoElement) SetDisablePictureInPicture(val bool) bool {
	return js.True == bindings.SetHTMLVideoElementDisablePictureInPicture(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// RequestVideoFrameCallback calls the method "HTMLVideoElement.requestVideoFrameCallback".
//
// The returned bool will be false if there is no such method.
func (this HTMLVideoElement) RequestVideoFrameCallback(callback js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)]) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallHTMLVideoElementRequestVideoFrameCallback(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return uint32(_ret), _ok
}

// RequestVideoFrameCallbackFunc returns the method "HTMLVideoElement.requestVideoFrameCallback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLVideoElement) RequestVideoFrameCallbackFunc() (fn js.Func[func(callback js.Func[func(now DOMHighResTimeStamp, metadata VideoFrameCallbackMetadata)]) uint32]) {
	return fn.FromRef(
		bindings.HTMLVideoElementRequestVideoFrameCallbackFunc(
			this.Ref(),
		),
	)
}

// CancelVideoFrameCallback calls the method "HTMLVideoElement.cancelVideoFrameCallback".
//
// The returned bool will be false if there is no such method.
func (this HTMLVideoElement) CancelVideoFrameCallback(handle uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLVideoElementCancelVideoFrameCallback(
		this.Ref(), js.Pointer(&_ok),
		uint32(handle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelVideoFrameCallbackFunc returns the method "HTMLVideoElement.cancelVideoFrameCallback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLVideoElement) CancelVideoFrameCallbackFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.HTMLVideoElementCancelVideoFrameCallbackFunc(
			this.Ref(),
		),
	)
}

// RequestPictureInPicture calls the method "HTMLVideoElement.requestPictureInPicture".
//
// The returned bool will be false if there is no such method.
func (this HTMLVideoElement) RequestPictureInPicture() (js.Promise[PictureInPictureWindow], bool) {
	var _ok bool
	_ret := bindings.CallHTMLVideoElementRequestPictureInPicture(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PictureInPictureWindow]{}.FromRef(_ret), _ok
}

// RequestPictureInPictureFunc returns the method "HTMLVideoElement.requestPictureInPicture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLVideoElement) RequestPictureInPictureFunc() (fn js.Func[func() js.Promise[PictureInPictureWindow]]) {
	return fn.FromRef(
		bindings.HTMLVideoElementRequestPictureInPictureFunc(
			this.Ref(),
		),
	)
}

// GetVideoPlaybackQuality calls the method "HTMLVideoElement.getVideoPlaybackQuality".
//
// The returned bool will be false if there is no such method.
func (this HTMLVideoElement) GetVideoPlaybackQuality() (VideoPlaybackQuality, bool) {
	var _ok bool
	_ret := bindings.CallHTMLVideoElementGetVideoPlaybackQuality(
		this.Ref(), js.Pointer(&_ok),
	)

	return VideoPlaybackQuality{}.FromRef(_ret), _ok
}

// GetVideoPlaybackQualityFunc returns the method "HTMLVideoElement.getVideoPlaybackQuality".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLVideoElement) GetVideoPlaybackQualityFunc() (fn js.Func[func() VideoPlaybackQuality]) {
	return fn.FromRef(
		bindings.HTMLVideoElementGetVideoPlaybackQualityFunc(
			this.Ref(),
		),
	)
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
	Alpha bool
	// Desynchronized is "CanvasRenderingContext2DSettings.desynchronized"
	//
	// Optional, defaults to false.
	Desynchronized bool
	// ColorSpace is "CanvasRenderingContext2DSettings.colorSpace"
	//
	// Optional, defaults to "srgb".
	ColorSpace PredefinedColorSpace
	// WillReadFrequently is "CanvasRenderingContext2DSettings.willReadFrequently"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CanvasRenderingContext2DSettings CanvasRenderingContext2DSettings [// CanvasRenderingContext2DSettings] [0x140002c2780 0x140002c28c0 0x140002c2a00 0x140002c2aa0 0x140002c2820 0x140002c2960 0x140002c2b40] 0x1400107e9d8 {0 0}} in the application heap.
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

// AddColorStop calls the method "CanvasGradient.addColorStop".
//
// The returned bool will be false if there is no such method.
func (this CanvasGradient) AddColorStop(offset float64, color js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasGradientAddColorStop(
		this.Ref(), js.Pointer(&_ok),
		float64(offset),
		color.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddColorStopFunc returns the method "CanvasGradient.addColorStop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasGradient) AddColorStopFunc() (fn js.Func[func(offset float64, color js.String)]) {
	return fn.FromRef(
		bindings.CanvasGradientAddColorStopFunc(
			this.Ref(),
		),
	)
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

// SetTransform calls the method "CanvasPattern.setTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasPattern) SetTransform(transform DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasPatternSetTransform(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransformFunc returns the method "CanvasPattern.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasPattern) SetTransformFunc() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.CanvasPatternSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "CanvasPattern.setTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasPattern) SetTransform1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasPatternSetTransform1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform1Func returns the method "CanvasPattern.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasPattern) SetTransform1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasPatternSetTransform1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this ImageBitmap) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageBitmapWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height returns the value of property "ImageBitmap.height".
//
// The returned bool will be false if there is no such property.
func (this ImageBitmap) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageBitmapHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Close calls the method "ImageBitmap.close".
//
// The returned bool will be false if there is no such method.
func (this ImageBitmap) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallImageBitmapClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "ImageBitmap.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageBitmap) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ImageBitmapCloseFunc(
			this.Ref(),
		),
	)
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

func NewPath2D(path OneOf_Path2D_String) Path2D {
	return Path2D{}.FromRef(
		bindings.NewPath2DByPath2D(
			path.Ref()),
	)
}

func NewPath2DByPath2D1() Path2D {
	return Path2D{}.FromRef(
		bindings.NewPath2DByPath2D1(),
	)
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

// AddPath calls the method "Path2D.addPath".
//
// The returned bool will be false if there is no such method.
func (this Path2D) AddPath(path Path2D, transform DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DAddPath(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		js.Pointer(&transform),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddPathFunc returns the method "Path2D.addPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) AddPathFunc() (fn js.Func[func(path Path2D, transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.Path2DAddPathFunc(
			this.Ref(),
		),
	)
}

// AddPath1 calls the method "Path2D.addPath".
//
// The returned bool will be false if there is no such method.
func (this Path2D) AddPath1(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DAddPath1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddPath1Func returns the method "Path2D.addPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) AddPath1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.Path2DAddPath1Func(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "Path2D.closePath".
//
// The returned bool will be false if there is no such method.
func (this Path2D) ClosePath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DClosePath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClosePathFunc returns the method "Path2D.closePath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.Path2DClosePathFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "Path2D.moveTo".
//
// The returned bool will be false if there is no such method.
func (this Path2D) MoveTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DMoveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveToFunc returns the method "Path2D.moveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DMoveToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "Path2D.lineTo".
//
// The returned bool will be false if there is no such method.
func (this Path2D) LineTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DLineTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineToFunc returns the method "Path2D.lineTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DLineToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "Path2D.quadraticCurveTo".
//
// The returned bool will be false if there is no such method.
func (this Path2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QuadraticCurveToFunc returns the method "Path2D.quadraticCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "Path2D.bezierCurveTo".
//
// The returned bool will be false if there is no such method.
func (this Path2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DBezierCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BezierCurveToFunc returns the method "Path2D.bezierCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.Path2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "Path2D.arcTo".
//
// The returned bool will be false if there is no such method.
func (this Path2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DArcTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcToFunc returns the method "Path2D.arcTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.Path2DArcToFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "Path2D.rect".
//
// The returned bool will be false if there is no such method.
func (this Path2D) Rect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RectFunc returns the method "Path2D.rect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.Path2DRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "Path2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this Path2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DRoundRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRectFunc returns the method "Path2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.Path2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "Path2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this Path2D) RoundRect1(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DRoundRect1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRect1Func returns the method "Path2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.Path2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// Arc calls the method "Path2D.arc".
//
// The returned bool will be false if there is no such method.
func (this Path2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DArc(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcFunc returns the method "Path2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.Path2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "Path2D.arc".
//
// The returned bool will be false if there is no such method.
func (this Path2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DArc1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Arc1Func returns the method "Path2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.Path2DArc1Func(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "Path2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this Path2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DEllipse(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EllipseFunc returns the method "Path2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.Path2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "Path2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this Path2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPath2DEllipse1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Ellipse1Func returns the method "Path2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Path2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.Path2DEllipse1Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this TextMetrics) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ActualBoundingBoxLeft returns the value of property "TextMetrics.actualBoundingBoxLeft".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) ActualBoundingBoxLeft() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsActualBoundingBoxLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ActualBoundingBoxRight returns the value of property "TextMetrics.actualBoundingBoxRight".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) ActualBoundingBoxRight() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsActualBoundingBoxRight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FontBoundingBoxAscent returns the value of property "TextMetrics.fontBoundingBoxAscent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) FontBoundingBoxAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsFontBoundingBoxAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// FontBoundingBoxDescent returns the value of property "TextMetrics.fontBoundingBoxDescent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) FontBoundingBoxDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsFontBoundingBoxDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ActualBoundingBoxAscent returns the value of property "TextMetrics.actualBoundingBoxAscent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) ActualBoundingBoxAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsActualBoundingBoxAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ActualBoundingBoxDescent returns the value of property "TextMetrics.actualBoundingBoxDescent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) ActualBoundingBoxDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsActualBoundingBoxDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// EmHeightAscent returns the value of property "TextMetrics.emHeightAscent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) EmHeightAscent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsEmHeightAscent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// EmHeightDescent returns the value of property "TextMetrics.emHeightDescent".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) EmHeightDescent() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsEmHeightDescent(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// HangingBaseline returns the value of property "TextMetrics.hangingBaseline".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) HangingBaseline() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsHangingBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AlphabeticBaseline returns the value of property "TextMetrics.alphabeticBaseline".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) AlphabeticBaseline() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsAlphabeticBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// IdeographicBaseline returns the value of property "TextMetrics.ideographicBaseline".
//
// The returned bool will be false if there is no such property.
func (this TextMetrics) IdeographicBaseline() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextMetricsIdeographicBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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

// New creates a new {0x140004cc0e0 ImageDataSettings ImageDataSettings [// ImageDataSettings] [0x140002c2d20] 0x1400107f020 {0 0}} in the application heap.
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

func NewImageData(sw uint32, sh uint32, settings ImageDataSettings) ImageData {
	return ImageData{}.FromRef(
		bindings.NewImageDataByImageData(
			uint32(sw),
			uint32(sh),
			js.Pointer(&settings)),
	)
}

func NewImageDataByImageData1(sw uint32, sh uint32) ImageData {
	return ImageData{}.FromRef(
		bindings.NewImageDataByImageData1(
			uint32(sw),
			uint32(sh)),
	)
}

func NewImageDataByImageData2(data js.TypedArray[uint8], sw uint32, sh uint32, settings ImageDataSettings) ImageData {
	return ImageData{}.FromRef(
		bindings.NewImageDataByImageData2(
			data.Ref(),
			uint32(sw),
			uint32(sh),
			js.Pointer(&settings)),
	)
}

func NewImageDataByImageData3(data js.TypedArray[uint8], sw uint32, sh uint32) ImageData {
	return ImageData{}.FromRef(
		bindings.NewImageDataByImageData3(
			data.Ref(),
			uint32(sw),
			uint32(sh)),
	)
}

func NewImageDataByImageData4(data js.TypedArray[uint8], sw uint32) ImageData {
	return ImageData{}.FromRef(
		bindings.NewImageDataByImageData4(
			data.Ref(),
			uint32(sw)),
	)
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
// The returned bool will be false if there is no such property.
func (this ImageData) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageDataWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height returns the value of property "ImageData.height".
//
// The returned bool will be false if there is no such property.
func (this ImageData) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetImageDataHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Data returns the value of property "ImageData.data".
//
// The returned bool will be false if there is no such property.
func (this ImageData) Data() (js.TypedArray[uint8], bool) {
	var _ok bool
	_ret := bindings.GetImageDataData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[uint8]{}.FromRef(_ret), _ok
}

// ColorSpace returns the value of property "ImageData.colorSpace".
//
// The returned bool will be false if there is no such property.
func (this ImageData) ColorSpace() (PredefinedColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetImageDataColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return PredefinedColorSpace(_ret), _ok
}

type ImageSmoothingQuality uint32

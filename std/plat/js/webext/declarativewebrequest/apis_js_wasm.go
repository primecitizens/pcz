// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package declarativewebrequest

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/declarativewebrequest/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/events"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrequest"
)

type RequestCookie struct {
	// Name is "RequestCookie.name"
	//
	// Optional
	Name js.String
	// Value is "RequestCookie.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestCookie with all fields set.
func (p RequestCookie) FromRef(ref js.Ref) RequestCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestCookie in the application heap.
func (p RequestCookie) New() js.Ref {
	return bindings.RequestCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestCookie) UpdateFrom(ref js.Ref) {
	bindings.RequestCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestCookie) Update(ref js.Ref) {
	bindings.RequestCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestCookie) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type AddRequestCookieInstanceType uint32

const (
	_ AddRequestCookieInstanceType = iota

	AddRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_ADD_REQUEST_COOKIE
)

func (AddRequestCookieInstanceType) FromRef(str js.Ref) AddRequestCookieInstanceType {
	return AddRequestCookieInstanceType(bindings.ConstOfAddRequestCookieInstanceType(str))
}

func (x AddRequestCookieInstanceType) String() (string, bool) {
	switch x {
	case AddRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_ADD_REQUEST_COOKIE:
		return "declarativeWebRequest.AddRequestCookie", true
	default:
		return "", false
	}
}

type AddRequestCookie struct {
	// Cookie is "AddRequestCookie.cookie"
	//
	// Required
	//
	// NOTE: Cookie.FFI_USE MUST be set to true to get Cookie used.
	Cookie RequestCookie
	// InstanceType is "AddRequestCookie.instanceType"
	//
	// Required
	InstanceType AddRequestCookieInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddRequestCookie with all fields set.
func (p AddRequestCookie) FromRef(ref js.Ref) AddRequestCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddRequestCookie in the application heap.
func (p AddRequestCookie) New() js.Ref {
	return bindings.AddRequestCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddRequestCookie) UpdateFrom(ref js.Ref) {
	bindings.AddRequestCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddRequestCookie) Update(ref js.Ref) {
	bindings.AddRequestCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddRequestCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Cookie.FreeMembers(true)
	}
}

type ResponseCookie struct {
	// Domain is "ResponseCookie.domain"
	//
	// Optional
	Domain js.String
	// Expires is "ResponseCookie.expires"
	//
	// Optional
	Expires js.String
	// HttpOnly is "ResponseCookie.httpOnly"
	//
	// Optional
	HttpOnly js.String
	// MaxAge is "ResponseCookie.maxAge"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxAge MUST be set to true to make this field effective.
	MaxAge float64
	// Name is "ResponseCookie.name"
	//
	// Optional
	Name js.String
	// Path is "ResponseCookie.path"
	//
	// Optional
	Path js.String
	// Secure is "ResponseCookie.secure"
	//
	// Optional
	Secure js.String
	// Value is "ResponseCookie.value"
	//
	// Optional
	Value js.String

	FFI_USE_MaxAge bool // for MaxAge.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResponseCookie with all fields set.
func (p ResponseCookie) FromRef(ref js.Ref) ResponseCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResponseCookie in the application heap.
func (p ResponseCookie) New() js.Ref {
	return bindings.ResponseCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResponseCookie) UpdateFrom(ref js.Ref) {
	bindings.ResponseCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResponseCookie) Update(ref js.Ref) {
	bindings.ResponseCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResponseCookie) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Expires.Ref(),
		p.HttpOnly.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.Secure.Ref(),
		p.Value.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Expires = p.Expires.FromRef(js.Undefined)
	p.HttpOnly = p.HttpOnly.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Secure = p.Secure.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type AddResponseCookieInstanceType uint32

const (
	_ AddResponseCookieInstanceType = iota

	AddResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_ADD_RESPONSE_COOKIE
)

func (AddResponseCookieInstanceType) FromRef(str js.Ref) AddResponseCookieInstanceType {
	return AddResponseCookieInstanceType(bindings.ConstOfAddResponseCookieInstanceType(str))
}

func (x AddResponseCookieInstanceType) String() (string, bool) {
	switch x {
	case AddResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_ADD_RESPONSE_COOKIE:
		return "declarativeWebRequest.AddResponseCookie", true
	default:
		return "", false
	}
}

type AddResponseCookie struct {
	// Cookie is "AddResponseCookie.cookie"
	//
	// Required
	//
	// NOTE: Cookie.FFI_USE MUST be set to true to get Cookie used.
	Cookie ResponseCookie
	// InstanceType is "AddResponseCookie.instanceType"
	//
	// Required
	InstanceType AddResponseCookieInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddResponseCookie with all fields set.
func (p AddResponseCookie) FromRef(ref js.Ref) AddResponseCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddResponseCookie in the application heap.
func (p AddResponseCookie) New() js.Ref {
	return bindings.AddResponseCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddResponseCookie) UpdateFrom(ref js.Ref) {
	bindings.AddResponseCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddResponseCookie) Update(ref js.Ref) {
	bindings.AddResponseCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddResponseCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Cookie.FreeMembers(true)
	}
}

type AddResponseHeaderInstanceType uint32

const (
	_ AddResponseHeaderInstanceType = iota

	AddResponseHeaderInstanceType_DECLARATIVE_WEB_REQUEST_ADD_RESPONSE_HEADER
)

func (AddResponseHeaderInstanceType) FromRef(str js.Ref) AddResponseHeaderInstanceType {
	return AddResponseHeaderInstanceType(bindings.ConstOfAddResponseHeaderInstanceType(str))
}

func (x AddResponseHeaderInstanceType) String() (string, bool) {
	switch x {
	case AddResponseHeaderInstanceType_DECLARATIVE_WEB_REQUEST_ADD_RESPONSE_HEADER:
		return "declarativeWebRequest.AddResponseHeader", true
	default:
		return "", false
	}
}

type AddResponseHeader struct {
	// InstanceType is "AddResponseHeader.instanceType"
	//
	// Required
	InstanceType AddResponseHeaderInstanceType
	// Name is "AddResponseHeader.name"
	//
	// Required
	Name js.String
	// Value is "AddResponseHeader.value"
	//
	// Required
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddResponseHeader with all fields set.
func (p AddResponseHeader) FromRef(ref js.Ref) AddResponseHeader {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddResponseHeader in the application heap.
func (p AddResponseHeader) New() js.Ref {
	return bindings.AddResponseHeaderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddResponseHeader) UpdateFrom(ref js.Ref) {
	bindings.AddResponseHeaderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddResponseHeader) Update(ref js.Ref) {
	bindings.AddResponseHeaderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddResponseHeader) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type CancelRequestInstanceType uint32

const (
	_ CancelRequestInstanceType = iota

	CancelRequestInstanceType_DECLARATIVE_WEB_REQUEST_CANCEL_REQUEST
)

func (CancelRequestInstanceType) FromRef(str js.Ref) CancelRequestInstanceType {
	return CancelRequestInstanceType(bindings.ConstOfCancelRequestInstanceType(str))
}

func (x CancelRequestInstanceType) String() (string, bool) {
	switch x {
	case CancelRequestInstanceType_DECLARATIVE_WEB_REQUEST_CANCEL_REQUEST:
		return "declarativeWebRequest.CancelRequest", true
	default:
		return "", false
	}
}

type CancelRequest struct {
	// InstanceType is "CancelRequest.instanceType"
	//
	// Required
	InstanceType CancelRequestInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CancelRequest with all fields set.
func (p CancelRequest) FromRef(ref js.Ref) CancelRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CancelRequest in the application heap.
func (p CancelRequest) New() js.Ref {
	return bindings.CancelRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CancelRequest) UpdateFrom(ref js.Ref) {
	bindings.CancelRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CancelRequest) Update(ref js.Ref) {
	bindings.CancelRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CancelRequest) FreeMembers(recursive bool) {
}

type EditRequestCookieInstanceType uint32

const (
	_ EditRequestCookieInstanceType = iota

	EditRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_EDIT_REQUEST_COOKIE
)

func (EditRequestCookieInstanceType) FromRef(str js.Ref) EditRequestCookieInstanceType {
	return EditRequestCookieInstanceType(bindings.ConstOfEditRequestCookieInstanceType(str))
}

func (x EditRequestCookieInstanceType) String() (string, bool) {
	switch x {
	case EditRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_EDIT_REQUEST_COOKIE:
		return "declarativeWebRequest.EditRequestCookie", true
	default:
		return "", false
	}
}

type EditRequestCookie struct {
	// Filter is "EditRequestCookie.filter"
	//
	// Required
	//
	// NOTE: Filter.FFI_USE MUST be set to true to get Filter used.
	Filter RequestCookie
	// InstanceType is "EditRequestCookie.instanceType"
	//
	// Required
	InstanceType EditRequestCookieInstanceType
	// Modification is "EditRequestCookie.modification"
	//
	// Required
	//
	// NOTE: Modification.FFI_USE MUST be set to true to get Modification used.
	Modification RequestCookie

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EditRequestCookie with all fields set.
func (p EditRequestCookie) FromRef(ref js.Ref) EditRequestCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EditRequestCookie in the application heap.
func (p EditRequestCookie) New() js.Ref {
	return bindings.EditRequestCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EditRequestCookie) UpdateFrom(ref js.Ref) {
	bindings.EditRequestCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EditRequestCookie) Update(ref js.Ref) {
	bindings.EditRequestCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EditRequestCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Filter.FreeMembers(true)
		p.Modification.FreeMembers(true)
	}
}

type FilterResponseCookie struct {
	// AgeLowerBound is "FilterResponseCookie.ageLowerBound"
	//
	// Optional
	//
	// NOTE: FFI_USE_AgeLowerBound MUST be set to true to make this field effective.
	AgeLowerBound int64
	// AgeUpperBound is "FilterResponseCookie.ageUpperBound"
	//
	// Optional
	//
	// NOTE: FFI_USE_AgeUpperBound MUST be set to true to make this field effective.
	AgeUpperBound int64
	// Domain is "FilterResponseCookie.domain"
	//
	// Optional
	Domain js.String
	// Expires is "FilterResponseCookie.expires"
	//
	// Optional
	Expires js.String
	// HttpOnly is "FilterResponseCookie.httpOnly"
	//
	// Optional
	HttpOnly js.String
	// MaxAge is "FilterResponseCookie.maxAge"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxAge MUST be set to true to make this field effective.
	MaxAge float64
	// Name is "FilterResponseCookie.name"
	//
	// Optional
	Name js.String
	// Path is "FilterResponseCookie.path"
	//
	// Optional
	Path js.String
	// Secure is "FilterResponseCookie.secure"
	//
	// Optional
	Secure js.String
	// SessionCookie is "FilterResponseCookie.sessionCookie"
	//
	// Optional
	//
	// NOTE: FFI_USE_SessionCookie MUST be set to true to make this field effective.
	SessionCookie bool
	// Value is "FilterResponseCookie.value"
	//
	// Optional
	Value js.String

	FFI_USE_AgeLowerBound bool // for AgeLowerBound.
	FFI_USE_AgeUpperBound bool // for AgeUpperBound.
	FFI_USE_MaxAge        bool // for MaxAge.
	FFI_USE_SessionCookie bool // for SessionCookie.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FilterResponseCookie with all fields set.
func (p FilterResponseCookie) FromRef(ref js.Ref) FilterResponseCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FilterResponseCookie in the application heap.
func (p FilterResponseCookie) New() js.Ref {
	return bindings.FilterResponseCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FilterResponseCookie) UpdateFrom(ref js.Ref) {
	bindings.FilterResponseCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FilterResponseCookie) Update(ref js.Ref) {
	bindings.FilterResponseCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FilterResponseCookie) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Expires.Ref(),
		p.HttpOnly.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.Secure.Ref(),
		p.Value.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Expires = p.Expires.FromRef(js.Undefined)
	p.HttpOnly = p.HttpOnly.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.Secure = p.Secure.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type EditResponseCookieInstanceType uint32

const (
	_ EditResponseCookieInstanceType = iota

	EditResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_EDIT_RESPONSE_COOKIE
)

func (EditResponseCookieInstanceType) FromRef(str js.Ref) EditResponseCookieInstanceType {
	return EditResponseCookieInstanceType(bindings.ConstOfEditResponseCookieInstanceType(str))
}

func (x EditResponseCookieInstanceType) String() (string, bool) {
	switch x {
	case EditResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_EDIT_RESPONSE_COOKIE:
		return "declarativeWebRequest.EditResponseCookie", true
	default:
		return "", false
	}
}

type EditResponseCookie struct {
	// Filter is "EditResponseCookie.filter"
	//
	// Required
	//
	// NOTE: Filter.FFI_USE MUST be set to true to get Filter used.
	Filter FilterResponseCookie
	// InstanceType is "EditResponseCookie.instanceType"
	//
	// Required
	InstanceType EditResponseCookieInstanceType
	// Modification is "EditResponseCookie.modification"
	//
	// Required
	//
	// NOTE: Modification.FFI_USE MUST be set to true to get Modification used.
	Modification ResponseCookie

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EditResponseCookie with all fields set.
func (p EditResponseCookie) FromRef(ref js.Ref) EditResponseCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EditResponseCookie in the application heap.
func (p EditResponseCookie) New() js.Ref {
	return bindings.EditResponseCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EditResponseCookie) UpdateFrom(ref js.Ref) {
	bindings.EditResponseCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EditResponseCookie) Update(ref js.Ref) {
	bindings.EditResponseCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EditResponseCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Filter.FreeMembers(true)
		p.Modification.FreeMembers(true)
	}
}

type OneOf_ArrayString_String struct {
	ref js.Ref
}

func (x OneOf_ArrayString_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayString_String) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayString_String) FromRef(ref js.Ref) OneOf_ArrayString_String {
	return OneOf_ArrayString_String{
		ref: ref,
	}
}

func (x OneOf_ArrayString_String) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

func (x OneOf_ArrayString_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type HeaderFilter struct {
	// NameContains is "HeaderFilter.nameContains"
	//
	// Optional
	NameContains OneOf_ArrayString_String
	// NameEquals is "HeaderFilter.nameEquals"
	//
	// Optional
	NameEquals js.String
	// NamePrefix is "HeaderFilter.namePrefix"
	//
	// Optional
	NamePrefix js.String
	// NameSuffix is "HeaderFilter.nameSuffix"
	//
	// Optional
	NameSuffix js.String
	// ValueContains is "HeaderFilter.valueContains"
	//
	// Optional
	ValueContains OneOf_ArrayString_String
	// ValueEquals is "HeaderFilter.valueEquals"
	//
	// Optional
	ValueEquals js.String
	// ValuePrefix is "HeaderFilter.valuePrefix"
	//
	// Optional
	ValuePrefix js.String
	// ValueSuffix is "HeaderFilter.valueSuffix"
	//
	// Optional
	ValueSuffix js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HeaderFilter with all fields set.
func (p HeaderFilter) FromRef(ref js.Ref) HeaderFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HeaderFilter in the application heap.
func (p HeaderFilter) New() js.Ref {
	return bindings.HeaderFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HeaderFilter) UpdateFrom(ref js.Ref) {
	bindings.HeaderFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HeaderFilter) Update(ref js.Ref) {
	bindings.HeaderFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HeaderFilter) FreeMembers(recursive bool) {
	js.Free(
		p.NameContains.Ref(),
		p.NameEquals.Ref(),
		p.NamePrefix.Ref(),
		p.NameSuffix.Ref(),
		p.ValueContains.Ref(),
		p.ValueEquals.Ref(),
		p.ValuePrefix.Ref(),
		p.ValueSuffix.Ref(),
	)
	p.NameContains = p.NameContains.FromRef(js.Undefined)
	p.NameEquals = p.NameEquals.FromRef(js.Undefined)
	p.NamePrefix = p.NamePrefix.FromRef(js.Undefined)
	p.NameSuffix = p.NameSuffix.FromRef(js.Undefined)
	p.ValueContains = p.ValueContains.FromRef(js.Undefined)
	p.ValueEquals = p.ValueEquals.FromRef(js.Undefined)
	p.ValuePrefix = p.ValuePrefix.FromRef(js.Undefined)
	p.ValueSuffix = p.ValueSuffix.FromRef(js.Undefined)
}

type IgnoreRulesInstanceType uint32

const (
	_ IgnoreRulesInstanceType = iota

	IgnoreRulesInstanceType_DECLARATIVE_WEB_REQUEST_IGNORE_RULES
)

func (IgnoreRulesInstanceType) FromRef(str js.Ref) IgnoreRulesInstanceType {
	return IgnoreRulesInstanceType(bindings.ConstOfIgnoreRulesInstanceType(str))
}

func (x IgnoreRulesInstanceType) String() (string, bool) {
	switch x {
	case IgnoreRulesInstanceType_DECLARATIVE_WEB_REQUEST_IGNORE_RULES:
		return "declarativeWebRequest.IgnoreRules", true
	default:
		return "", false
	}
}

type IgnoreRules struct {
	// HasTag is "IgnoreRules.hasTag"
	//
	// Optional
	HasTag js.String
	// InstanceType is "IgnoreRules.instanceType"
	//
	// Required
	InstanceType IgnoreRulesInstanceType
	// LowerPriorityThan is "IgnoreRules.lowerPriorityThan"
	//
	// Optional
	//
	// NOTE: FFI_USE_LowerPriorityThan MUST be set to true to make this field effective.
	LowerPriorityThan int64

	FFI_USE_LowerPriorityThan bool // for LowerPriorityThan.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IgnoreRules with all fields set.
func (p IgnoreRules) FromRef(ref js.Ref) IgnoreRules {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IgnoreRules in the application heap.
func (p IgnoreRules) New() js.Ref {
	return bindings.IgnoreRulesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IgnoreRules) UpdateFrom(ref js.Ref) {
	bindings.IgnoreRulesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IgnoreRules) Update(ref js.Ref) {
	bindings.IgnoreRulesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IgnoreRules) FreeMembers(recursive bool) {
	js.Free(
		p.HasTag.Ref(),
	)
	p.HasTag = p.HasTag.FromRef(js.Undefined)
}

type Stage uint32

const (
	_ Stage = iota

	Stage_ON_BEFORE_REQUEST
	Stage_ON_BEFORE_SEND_HEADERS
	Stage_ON_HEADERS_RECEIVED
	Stage_ON_AUTH_REQUIRED
)

func (Stage) FromRef(str js.Ref) Stage {
	return Stage(bindings.ConstOfStage(str))
}

func (x Stage) String() (string, bool) {
	switch x {
	case Stage_ON_BEFORE_REQUEST:
		return "onBeforeRequest", true
	case Stage_ON_BEFORE_SEND_HEADERS:
		return "onBeforeSendHeaders", true
	case Stage_ON_HEADERS_RECEIVED:
		return "onHeadersReceived", true
	case Stage_ON_AUTH_REQUIRED:
		return "onAuthRequired", true
	default:
		return "", false
	}
}

type OnMessageArgDetails struct {
	// DocumentId is "OnMessageArgDetails.documentId"
	//
	// Optional
	DocumentId js.String
	// DocumentLifecycle is "OnMessageArgDetails.documentLifecycle"
	//
	// Required
	DocumentLifecycle extensiontypes.DocumentLifecycle
	// FrameId is "OnMessageArgDetails.frameId"
	//
	// Required
	FrameId int64
	// FrameType is "OnMessageArgDetails.frameType"
	//
	// Required
	FrameType extensiontypes.FrameType
	// Message is "OnMessageArgDetails.message"
	//
	// Required
	Message js.String
	// Method is "OnMessageArgDetails.method"
	//
	// Required
	Method js.String
	// ParentDocumentId is "OnMessageArgDetails.parentDocumentId"
	//
	// Optional
	ParentDocumentId js.String
	// ParentFrameId is "OnMessageArgDetails.parentFrameId"
	//
	// Required
	ParentFrameId int64
	// RequestId is "OnMessageArgDetails.requestId"
	//
	// Required
	RequestId js.String
	// Stage is "OnMessageArgDetails.stage"
	//
	// Required
	Stage Stage
	// TabId is "OnMessageArgDetails.tabId"
	//
	// Required
	TabId int64
	// TimeStamp is "OnMessageArgDetails.timeStamp"
	//
	// Required
	TimeStamp float64
	// Type is "OnMessageArgDetails.type"
	//
	// Required
	Type webrequest.ResourceType
	// Url is "OnMessageArgDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnMessageArgDetails with all fields set.
func (p OnMessageArgDetails) FromRef(ref js.Ref) OnMessageArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnMessageArgDetails in the application heap.
func (p OnMessageArgDetails) New() js.Ref {
	return bindings.OnMessageArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnMessageArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnMessageArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnMessageArgDetails) Update(ref js.Ref) {
	bindings.OnMessageArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnMessageArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentId.Ref(),
		p.Message.Ref(),
		p.Method.Ref(),
		p.ParentDocumentId.Ref(),
		p.RequestId.Ref(),
		p.Url.Ref(),
	)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
	p.Method = p.Method.FromRef(js.Undefined)
	p.ParentDocumentId = p.ParentDocumentId.FromRef(js.Undefined)
	p.RequestId = p.RequestId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type RedirectByRegExInstanceType uint32

const (
	_ RedirectByRegExInstanceType = iota

	RedirectByRegExInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_BY_REG_EX
)

func (RedirectByRegExInstanceType) FromRef(str js.Ref) RedirectByRegExInstanceType {
	return RedirectByRegExInstanceType(bindings.ConstOfRedirectByRegExInstanceType(str))
}

func (x RedirectByRegExInstanceType) String() (string, bool) {
	switch x {
	case RedirectByRegExInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_BY_REG_EX:
		return "declarativeWebRequest.RedirectByRegEx", true
	default:
		return "", false
	}
}

type RedirectByRegEx struct {
	// From is "RedirectByRegEx.from"
	//
	// Required
	From js.String
	// InstanceType is "RedirectByRegEx.instanceType"
	//
	// Required
	InstanceType RedirectByRegExInstanceType
	// To is "RedirectByRegEx.to"
	//
	// Required
	To js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RedirectByRegEx with all fields set.
func (p RedirectByRegEx) FromRef(ref js.Ref) RedirectByRegEx {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RedirectByRegEx in the application heap.
func (p RedirectByRegEx) New() js.Ref {
	return bindings.RedirectByRegExJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RedirectByRegEx) UpdateFrom(ref js.Ref) {
	bindings.RedirectByRegExJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RedirectByRegEx) Update(ref js.Ref) {
	bindings.RedirectByRegExJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RedirectByRegEx) FreeMembers(recursive bool) {
	js.Free(
		p.From.Ref(),
		p.To.Ref(),
	)
	p.From = p.From.FromRef(js.Undefined)
	p.To = p.To.FromRef(js.Undefined)
}

type RedirectRequestInstanceType uint32

const (
	_ RedirectRequestInstanceType = iota

	RedirectRequestInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_REQUEST
)

func (RedirectRequestInstanceType) FromRef(str js.Ref) RedirectRequestInstanceType {
	return RedirectRequestInstanceType(bindings.ConstOfRedirectRequestInstanceType(str))
}

func (x RedirectRequestInstanceType) String() (string, bool) {
	switch x {
	case RedirectRequestInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_REQUEST:
		return "declarativeWebRequest.RedirectRequest", true
	default:
		return "", false
	}
}

type RedirectRequest struct {
	// InstanceType is "RedirectRequest.instanceType"
	//
	// Required
	InstanceType RedirectRequestInstanceType
	// RedirectUrl is "RedirectRequest.redirectUrl"
	//
	// Required
	RedirectUrl js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RedirectRequest with all fields set.
func (p RedirectRequest) FromRef(ref js.Ref) RedirectRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RedirectRequest in the application heap.
func (p RedirectRequest) New() js.Ref {
	return bindings.RedirectRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RedirectRequest) UpdateFrom(ref js.Ref) {
	bindings.RedirectRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RedirectRequest) Update(ref js.Ref) {
	bindings.RedirectRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RedirectRequest) FreeMembers(recursive bool) {
	js.Free(
		p.RedirectUrl.Ref(),
	)
	p.RedirectUrl = p.RedirectUrl.FromRef(js.Undefined)
}

type RedirectToEmptyDocumentInstanceType uint32

const (
	_ RedirectToEmptyDocumentInstanceType = iota

	RedirectToEmptyDocumentInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_TO_EMPTY_DOCUMENT
)

func (RedirectToEmptyDocumentInstanceType) FromRef(str js.Ref) RedirectToEmptyDocumentInstanceType {
	return RedirectToEmptyDocumentInstanceType(bindings.ConstOfRedirectToEmptyDocumentInstanceType(str))
}

func (x RedirectToEmptyDocumentInstanceType) String() (string, bool) {
	switch x {
	case RedirectToEmptyDocumentInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_TO_EMPTY_DOCUMENT:
		return "declarativeWebRequest.RedirectToEmptyDocument", true
	default:
		return "", false
	}
}

type RedirectToEmptyDocument struct {
	// InstanceType is "RedirectToEmptyDocument.instanceType"
	//
	// Required
	InstanceType RedirectToEmptyDocumentInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RedirectToEmptyDocument with all fields set.
func (p RedirectToEmptyDocument) FromRef(ref js.Ref) RedirectToEmptyDocument {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RedirectToEmptyDocument in the application heap.
func (p RedirectToEmptyDocument) New() js.Ref {
	return bindings.RedirectToEmptyDocumentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RedirectToEmptyDocument) UpdateFrom(ref js.Ref) {
	bindings.RedirectToEmptyDocumentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RedirectToEmptyDocument) Update(ref js.Ref) {
	bindings.RedirectToEmptyDocumentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RedirectToEmptyDocument) FreeMembers(recursive bool) {
}

type RedirectToTransparentImageInstanceType uint32

const (
	_ RedirectToTransparentImageInstanceType = iota

	RedirectToTransparentImageInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_TO_TRANSPARENT_IMAGE
)

func (RedirectToTransparentImageInstanceType) FromRef(str js.Ref) RedirectToTransparentImageInstanceType {
	return RedirectToTransparentImageInstanceType(bindings.ConstOfRedirectToTransparentImageInstanceType(str))
}

func (x RedirectToTransparentImageInstanceType) String() (string, bool) {
	switch x {
	case RedirectToTransparentImageInstanceType_DECLARATIVE_WEB_REQUEST_REDIRECT_TO_TRANSPARENT_IMAGE:
		return "declarativeWebRequest.RedirectToTransparentImage", true
	default:
		return "", false
	}
}

type RedirectToTransparentImage struct {
	// InstanceType is "RedirectToTransparentImage.instanceType"
	//
	// Required
	InstanceType RedirectToTransparentImageInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RedirectToTransparentImage with all fields set.
func (p RedirectToTransparentImage) FromRef(ref js.Ref) RedirectToTransparentImage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RedirectToTransparentImage in the application heap.
func (p RedirectToTransparentImage) New() js.Ref {
	return bindings.RedirectToTransparentImageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RedirectToTransparentImage) UpdateFrom(ref js.Ref) {
	bindings.RedirectToTransparentImageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RedirectToTransparentImage) Update(ref js.Ref) {
	bindings.RedirectToTransparentImageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RedirectToTransparentImage) FreeMembers(recursive bool) {
}

type RemoveRequestCookieInstanceType uint32

const (
	_ RemoveRequestCookieInstanceType = iota

	RemoveRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_REQUEST_COOKIE
)

func (RemoveRequestCookieInstanceType) FromRef(str js.Ref) RemoveRequestCookieInstanceType {
	return RemoveRequestCookieInstanceType(bindings.ConstOfRemoveRequestCookieInstanceType(str))
}

func (x RemoveRequestCookieInstanceType) String() (string, bool) {
	switch x {
	case RemoveRequestCookieInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_REQUEST_COOKIE:
		return "declarativeWebRequest.RemoveRequestCookie", true
	default:
		return "", false
	}
}

type RemoveRequestCookie struct {
	// Filter is "RemoveRequestCookie.filter"
	//
	// Required
	//
	// NOTE: Filter.FFI_USE MUST be set to true to get Filter used.
	Filter RequestCookie
	// InstanceType is "RemoveRequestCookie.instanceType"
	//
	// Required
	InstanceType RemoveRequestCookieInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveRequestCookie with all fields set.
func (p RemoveRequestCookie) FromRef(ref js.Ref) RemoveRequestCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveRequestCookie in the application heap.
func (p RemoveRequestCookie) New() js.Ref {
	return bindings.RemoveRequestCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveRequestCookie) UpdateFrom(ref js.Ref) {
	bindings.RemoveRequestCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveRequestCookie) Update(ref js.Ref) {
	bindings.RemoveRequestCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveRequestCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Filter.FreeMembers(true)
	}
}

type RemoveRequestHeaderInstanceType uint32

const (
	_ RemoveRequestHeaderInstanceType = iota

	RemoveRequestHeaderInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_REQUEST_HEADER
)

func (RemoveRequestHeaderInstanceType) FromRef(str js.Ref) RemoveRequestHeaderInstanceType {
	return RemoveRequestHeaderInstanceType(bindings.ConstOfRemoveRequestHeaderInstanceType(str))
}

func (x RemoveRequestHeaderInstanceType) String() (string, bool) {
	switch x {
	case RemoveRequestHeaderInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_REQUEST_HEADER:
		return "declarativeWebRequest.RemoveRequestHeader", true
	default:
		return "", false
	}
}

type RemoveRequestHeader struct {
	// InstanceType is "RemoveRequestHeader.instanceType"
	//
	// Required
	InstanceType RemoveRequestHeaderInstanceType
	// Name is "RemoveRequestHeader.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveRequestHeader with all fields set.
func (p RemoveRequestHeader) FromRef(ref js.Ref) RemoveRequestHeader {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveRequestHeader in the application heap.
func (p RemoveRequestHeader) New() js.Ref {
	return bindings.RemoveRequestHeaderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveRequestHeader) UpdateFrom(ref js.Ref) {
	bindings.RemoveRequestHeaderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveRequestHeader) Update(ref js.Ref) {
	bindings.RemoveRequestHeaderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveRequestHeader) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type RemoveResponseCookieInstanceType uint32

const (
	_ RemoveResponseCookieInstanceType = iota

	RemoveResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_RESPONSE_COOKIE
)

func (RemoveResponseCookieInstanceType) FromRef(str js.Ref) RemoveResponseCookieInstanceType {
	return RemoveResponseCookieInstanceType(bindings.ConstOfRemoveResponseCookieInstanceType(str))
}

func (x RemoveResponseCookieInstanceType) String() (string, bool) {
	switch x {
	case RemoveResponseCookieInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_RESPONSE_COOKIE:
		return "declarativeWebRequest.RemoveResponseCookie", true
	default:
		return "", false
	}
}

type RemoveResponseCookie struct {
	// Filter is "RemoveResponseCookie.filter"
	//
	// Required
	//
	// NOTE: Filter.FFI_USE MUST be set to true to get Filter used.
	Filter FilterResponseCookie
	// InstanceType is "RemoveResponseCookie.instanceType"
	//
	// Required
	InstanceType RemoveResponseCookieInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveResponseCookie with all fields set.
func (p RemoveResponseCookie) FromRef(ref js.Ref) RemoveResponseCookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveResponseCookie in the application heap.
func (p RemoveResponseCookie) New() js.Ref {
	return bindings.RemoveResponseCookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveResponseCookie) UpdateFrom(ref js.Ref) {
	bindings.RemoveResponseCookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveResponseCookie) Update(ref js.Ref) {
	bindings.RemoveResponseCookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveResponseCookie) FreeMembers(recursive bool) {
	if recursive {
		p.Filter.FreeMembers(true)
	}
}

type RemoveResponseHeaderInstanceType uint32

const (
	_ RemoveResponseHeaderInstanceType = iota

	RemoveResponseHeaderInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_RESPONSE_HEADER
)

func (RemoveResponseHeaderInstanceType) FromRef(str js.Ref) RemoveResponseHeaderInstanceType {
	return RemoveResponseHeaderInstanceType(bindings.ConstOfRemoveResponseHeaderInstanceType(str))
}

func (x RemoveResponseHeaderInstanceType) String() (string, bool) {
	switch x {
	case RemoveResponseHeaderInstanceType_DECLARATIVE_WEB_REQUEST_REMOVE_RESPONSE_HEADER:
		return "declarativeWebRequest.RemoveResponseHeader", true
	default:
		return "", false
	}
}

type RemoveResponseHeader struct {
	// InstanceType is "RemoveResponseHeader.instanceType"
	//
	// Required
	InstanceType RemoveResponseHeaderInstanceType
	// Name is "RemoveResponseHeader.name"
	//
	// Required
	Name js.String
	// Value is "RemoveResponseHeader.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveResponseHeader with all fields set.
func (p RemoveResponseHeader) FromRef(ref js.Ref) RemoveResponseHeader {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveResponseHeader in the application heap.
func (p RemoveResponseHeader) New() js.Ref {
	return bindings.RemoveResponseHeaderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveResponseHeader) UpdateFrom(ref js.Ref) {
	bindings.RemoveResponseHeaderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveResponseHeader) Update(ref js.Ref) {
	bindings.RemoveResponseHeaderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveResponseHeader) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type RequestMatcherInstanceType uint32

const (
	_ RequestMatcherInstanceType = iota

	RequestMatcherInstanceType_DECLARATIVE_WEB_REQUEST_REQUEST_MATCHER
)

func (RequestMatcherInstanceType) FromRef(str js.Ref) RequestMatcherInstanceType {
	return RequestMatcherInstanceType(bindings.ConstOfRequestMatcherInstanceType(str))
}

func (x RequestMatcherInstanceType) String() (string, bool) {
	switch x {
	case RequestMatcherInstanceType_DECLARATIVE_WEB_REQUEST_REQUEST_MATCHER:
		return "declarativeWebRequest.RequestMatcher", true
	default:
		return "", false
	}
}

type RequestMatcher struct {
	// ContentType is "RequestMatcher.contentType"
	//
	// Optional
	ContentType js.Array[js.String]
	// ExcludeContentType is "RequestMatcher.excludeContentType"
	//
	// Optional
	ExcludeContentType js.Array[js.String]
	// ExcludeRequestHeaders is "RequestMatcher.excludeRequestHeaders"
	//
	// Optional
	ExcludeRequestHeaders js.Array[HeaderFilter]
	// ExcludeResponseHeaders is "RequestMatcher.excludeResponseHeaders"
	//
	// Optional
	ExcludeResponseHeaders js.Array[HeaderFilter]
	// FirstPartyForCookiesUrl is "RequestMatcher.firstPartyForCookiesUrl"
	//
	// Optional
	//
	// NOTE: FirstPartyForCookiesUrl.FFI_USE MUST be set to true to get FirstPartyForCookiesUrl used.
	FirstPartyForCookiesUrl events.UrlFilter
	// InstanceType is "RequestMatcher.instanceType"
	//
	// Required
	InstanceType RequestMatcherInstanceType
	// RequestHeaders is "RequestMatcher.requestHeaders"
	//
	// Optional
	RequestHeaders js.Array[HeaderFilter]
	// ResourceType is "RequestMatcher.resourceType"
	//
	// Optional
	ResourceType js.Array[webrequest.ResourceType]
	// ResponseHeaders is "RequestMatcher.responseHeaders"
	//
	// Optional
	ResponseHeaders js.Array[HeaderFilter]
	// Stages is "RequestMatcher.stages"
	//
	// Optional
	Stages js.Array[Stage]
	// ThirdPartyForCookies is "RequestMatcher.thirdPartyForCookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_ThirdPartyForCookies MUST be set to true to make this field effective.
	ThirdPartyForCookies bool
	// Url is "RequestMatcher.url"
	//
	// Optional
	//
	// NOTE: Url.FFI_USE MUST be set to true to get Url used.
	Url events.UrlFilter

	FFI_USE_ThirdPartyForCookies bool // for ThirdPartyForCookies.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestMatcher with all fields set.
func (p RequestMatcher) FromRef(ref js.Ref) RequestMatcher {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestMatcher in the application heap.
func (p RequestMatcher) New() js.Ref {
	return bindings.RequestMatcherJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestMatcher) UpdateFrom(ref js.Ref) {
	bindings.RequestMatcherJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestMatcher) Update(ref js.Ref) {
	bindings.RequestMatcherJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestMatcher) FreeMembers(recursive bool) {
	js.Free(
		p.ContentType.Ref(),
		p.ExcludeContentType.Ref(),
		p.ExcludeRequestHeaders.Ref(),
		p.ExcludeResponseHeaders.Ref(),
		p.RequestHeaders.Ref(),
		p.ResourceType.Ref(),
		p.ResponseHeaders.Ref(),
		p.Stages.Ref(),
	)
	p.ContentType = p.ContentType.FromRef(js.Undefined)
	p.ExcludeContentType = p.ExcludeContentType.FromRef(js.Undefined)
	p.ExcludeRequestHeaders = p.ExcludeRequestHeaders.FromRef(js.Undefined)
	p.ExcludeResponseHeaders = p.ExcludeResponseHeaders.FromRef(js.Undefined)
	p.RequestHeaders = p.RequestHeaders.FromRef(js.Undefined)
	p.ResourceType = p.ResourceType.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
	p.Stages = p.Stages.FromRef(js.Undefined)
	if recursive {
		p.FirstPartyForCookiesUrl.FreeMembers(true)
		p.Url.FreeMembers(true)
	}
}

type SendMessageToExtensionInstanceType uint32

const (
	_ SendMessageToExtensionInstanceType = iota

	SendMessageToExtensionInstanceType_DECLARATIVE_WEB_REQUEST_SEND_MESSAGE_TO_EXTENSION
)

func (SendMessageToExtensionInstanceType) FromRef(str js.Ref) SendMessageToExtensionInstanceType {
	return SendMessageToExtensionInstanceType(bindings.ConstOfSendMessageToExtensionInstanceType(str))
}

func (x SendMessageToExtensionInstanceType) String() (string, bool) {
	switch x {
	case SendMessageToExtensionInstanceType_DECLARATIVE_WEB_REQUEST_SEND_MESSAGE_TO_EXTENSION:
		return "declarativeWebRequest.SendMessageToExtension", true
	default:
		return "", false
	}
}

type SendMessageToExtension struct {
	// InstanceType is "SendMessageToExtension.instanceType"
	//
	// Required
	InstanceType SendMessageToExtensionInstanceType
	// Message is "SendMessageToExtension.message"
	//
	// Required
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendMessageToExtension with all fields set.
func (p SendMessageToExtension) FromRef(ref js.Ref) SendMessageToExtension {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendMessageToExtension in the application heap.
func (p SendMessageToExtension) New() js.Ref {
	return bindings.SendMessageToExtensionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendMessageToExtension) UpdateFrom(ref js.Ref) {
	bindings.SendMessageToExtensionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendMessageToExtension) Update(ref js.Ref) {
	bindings.SendMessageToExtensionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendMessageToExtension) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type SetRequestHeaderInstanceType uint32

const (
	_ SetRequestHeaderInstanceType = iota

	SetRequestHeaderInstanceType_DECLARATIVE_WEB_REQUEST_SET_REQUEST_HEADER
)

func (SetRequestHeaderInstanceType) FromRef(str js.Ref) SetRequestHeaderInstanceType {
	return SetRequestHeaderInstanceType(bindings.ConstOfSetRequestHeaderInstanceType(str))
}

func (x SetRequestHeaderInstanceType) String() (string, bool) {
	switch x {
	case SetRequestHeaderInstanceType_DECLARATIVE_WEB_REQUEST_SET_REQUEST_HEADER:
		return "declarativeWebRequest.SetRequestHeader", true
	default:
		return "", false
	}
}

type SetRequestHeader struct {
	// InstanceType is "SetRequestHeader.instanceType"
	//
	// Required
	InstanceType SetRequestHeaderInstanceType
	// Name is "SetRequestHeader.name"
	//
	// Required
	Name js.String
	// Value is "SetRequestHeader.value"
	//
	// Required
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetRequestHeader with all fields set.
func (p SetRequestHeader) FromRef(ref js.Ref) SetRequestHeader {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetRequestHeader in the application heap.
func (p SetRequestHeader) New() js.Ref {
	return bindings.SetRequestHeaderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetRequestHeader) UpdateFrom(ref js.Ref) {
	bindings.SetRequestHeaderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetRequestHeader) Update(ref js.Ref) {
	bindings.SetRequestHeaderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetRequestHeader) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type OnMessageEventCallbackFunc func(this js.Ref, details *OnMessageArgDetails) js.Ref

func (fn OnMessageEventCallbackFunc) Register() js.Func[func(details *OnMessageArgDetails)] {
	return js.RegisterCallback[func(details *OnMessageArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgDetails
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

type OnMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnMessageArgDetails) js.Ref
	Arg T
}

func (cb *OnMessageEventCallback[T]) Register() js.Func[func(details *OnMessageArgDetails)] {
	return js.RegisterCallback[func(details *OnMessageArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMessageArgDetails
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

// HasFuncOnMessage returns true if the function "WEBEXT.declarativeWebRequest.onMessage.addListener" exists.
func HasFuncOnMessage() bool {
	return js.True == bindings.HasFuncOnMessage()
}

// FuncOnMessage returns the function "WEBEXT.declarativeWebRequest.onMessage.addListener".
func FuncOnMessage() (fn js.Func[func(callback js.Func[func(details *OnMessageArgDetails)])]) {
	bindings.FuncOnMessage(
		js.Pointer(&fn),
	)
	return
}

// OnMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.addListener" directly.
func OnMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret js.Void) {
	bindings.CallOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMessage returns true if the function "WEBEXT.declarativeWebRequest.onMessage.removeListener" exists.
func HasFuncOffMessage() bool {
	return js.True == bindings.HasFuncOffMessage()
}

// FuncOffMessage returns the function "WEBEXT.declarativeWebRequest.onMessage.removeListener".
func FuncOffMessage() (fn js.Func[func(callback js.Func[func(details *OnMessageArgDetails)])]) {
	bindings.FuncOffMessage(
		js.Pointer(&fn),
	)
	return
}

// OffMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.removeListener" directly.
func OffMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret js.Void) {
	bindings.CallOffMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMessage returns true if the function "WEBEXT.declarativeWebRequest.onMessage.hasListener" exists.
func HasFuncHasOnMessage() bool {
	return js.True == bindings.HasFuncHasOnMessage()
}

// FuncHasOnMessage returns the function "WEBEXT.declarativeWebRequest.onMessage.hasListener".
func FuncHasOnMessage() (fn js.Func[func(callback js.Func[func(details *OnMessageArgDetails)]) bool]) {
	bindings.FuncHasOnMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.hasListener" directly.
func HasOnMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret bool) {
	bindings.CallHasOnMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMessage calls the function "WEBEXT.declarativeWebRequest.onMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMessage(callback js.Func[func(details *OnMessageArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRequestEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnRequestEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestEventCallbackFunc) DispatchCallback(
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

type OnRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRequestEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestEventCallback[T]) DispatchCallback(
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

// HasFuncOnRequest returns true if the function "WEBEXT.declarativeWebRequest.onRequest.addListener" exists.
func HasFuncOnRequest() bool {
	return js.True == bindings.HasFuncOnRequest()
}

// FuncOnRequest returns the function "WEBEXT.declarativeWebRequest.onRequest.addListener".
func FuncOnRequest() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRequest(
		js.Pointer(&fn),
	)
	return
}

// OnRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.addListener" directly.
func OnRequest(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRequest(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRequest returns true if the function "WEBEXT.declarativeWebRequest.onRequest.removeListener" exists.
func HasFuncOffRequest() bool {
	return js.True == bindings.HasFuncOffRequest()
}

// FuncOffRequest returns the function "WEBEXT.declarativeWebRequest.onRequest.removeListener".
func FuncOffRequest() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRequest(
		js.Pointer(&fn),
	)
	return
}

// OffRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.removeListener" directly.
func OffRequest(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRequest(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRequest returns true if the function "WEBEXT.declarativeWebRequest.onRequest.hasListener" exists.
func HasFuncHasOnRequest() bool {
	return js.True == bindings.HasFuncHasOnRequest()
}

// FuncHasOnRequest returns the function "WEBEXT.declarativeWebRequest.onRequest.hasListener".
func FuncHasOnRequest() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.hasListener" directly.
func HasOnRequest(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRequest calls the function "WEBEXT.declarativeWebRequest.onRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRequest(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

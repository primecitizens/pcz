// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cookies

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/cookies/bindings"
)

type SameSiteStatus uint32

const (
	_ SameSiteStatus = iota

	SameSiteStatus_NO_RESTRICTION
	SameSiteStatus_LAX
	SameSiteStatus_STRICT
	SameSiteStatus_UNSPECIFIED
)

func (SameSiteStatus) FromRef(str js.Ref) SameSiteStatus {
	return SameSiteStatus(bindings.ConstOfSameSiteStatus(str))
}

func (x SameSiteStatus) String() (string, bool) {
	switch x {
	case SameSiteStatus_NO_RESTRICTION:
		return "no_restriction", true
	case SameSiteStatus_LAX:
		return "lax", true
	case SameSiteStatus_STRICT:
		return "strict", true
	case SameSiteStatus_UNSPECIFIED:
		return "unspecified", true
	default:
		return "", false
	}
}

type Cookie struct {
	// Domain is "Cookie.domain"
	//
	// Required
	Domain js.String
	// ExpirationDate is "Cookie.expirationDate"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExpirationDate MUST be set to true to make this field effective.
	ExpirationDate float64
	// HostOnly is "Cookie.hostOnly"
	//
	// Required
	HostOnly bool
	// HttpOnly is "Cookie.httpOnly"
	//
	// Required
	HttpOnly bool
	// Name is "Cookie.name"
	//
	// Required
	Name js.String
	// Path is "Cookie.path"
	//
	// Required
	Path js.String
	// SameSite is "Cookie.sameSite"
	//
	// Required
	SameSite SameSiteStatus
	// Secure is "Cookie.secure"
	//
	// Required
	Secure bool
	// Session is "Cookie.session"
	//
	// Required
	Session bool
	// StoreId is "Cookie.storeId"
	//
	// Required
	StoreId js.String
	// Value is "Cookie.value"
	//
	// Required
	Value js.String

	FFI_USE_ExpirationDate bool // for ExpirationDate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Cookie with all fields set.
func (p Cookie) FromRef(ref js.Ref) Cookie {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Cookie in the application heap.
func (p Cookie) New() js.Ref {
	return bindings.CookieJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Cookie) UpdateFrom(ref js.Ref) {
	bindings.CookieJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Cookie) Update(ref js.Ref) {
	bindings.CookieJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Cookie) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.StoreId.Ref(),
		p.Value.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.StoreId = p.StoreId.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type CookieDetails struct {
	// Name is "CookieDetails.name"
	//
	// Required
	Name js.String
	// StoreId is "CookieDetails.storeId"
	//
	// Optional
	StoreId js.String
	// Url is "CookieDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieDetails with all fields set.
func (p CookieDetails) FromRef(ref js.Ref) CookieDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CookieDetails in the application heap.
func (p CookieDetails) New() js.Ref {
	return bindings.CookieDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CookieDetails) UpdateFrom(ref js.Ref) {
	bindings.CookieDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CookieDetails) Update(ref js.Ref) {
	bindings.CookieDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CookieDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.StoreId.Ref(),
		p.Url.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.StoreId = p.StoreId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type CookieStore struct {
	// Id is "CookieStore.id"
	//
	// Required
	Id js.String
	// TabIds is "CookieStore.tabIds"
	//
	// Required
	TabIds js.Array[int64]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieStore with all fields set.
func (p CookieStore) FromRef(ref js.Ref) CookieStore {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CookieStore in the application heap.
func (p CookieStore) New() js.Ref {
	return bindings.CookieStoreJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CookieStore) UpdateFrom(ref js.Ref) {
	bindings.CookieStoreJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CookieStore) Update(ref js.Ref) {
	bindings.CookieStoreJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CookieStore) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.TabIds.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.TabIds = p.TabIds.FromRef(js.Undefined)
}

type GetAllArgDetails struct {
	// Domain is "GetAllArgDetails.domain"
	//
	// Optional
	Domain js.String
	// Name is "GetAllArgDetails.name"
	//
	// Optional
	Name js.String
	// Path is "GetAllArgDetails.path"
	//
	// Optional
	Path js.String
	// Secure is "GetAllArgDetails.secure"
	//
	// Optional
	//
	// NOTE: FFI_USE_Secure MUST be set to true to make this field effective.
	Secure bool
	// Session is "GetAllArgDetails.session"
	//
	// Optional
	//
	// NOTE: FFI_USE_Session MUST be set to true to make this field effective.
	Session bool
	// StoreId is "GetAllArgDetails.storeId"
	//
	// Optional
	StoreId js.String
	// Url is "GetAllArgDetails.url"
	//
	// Optional
	Url js.String

	FFI_USE_Secure  bool // for Secure.
	FFI_USE_Session bool // for Session.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetAllArgDetails with all fields set.
func (p GetAllArgDetails) FromRef(ref js.Ref) GetAllArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetAllArgDetails in the application heap.
func (p GetAllArgDetails) New() js.Ref {
	return bindings.GetAllArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetAllArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetAllArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetAllArgDetails) Update(ref js.Ref) {
	bindings.GetAllArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetAllArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.StoreId.Ref(),
		p.Url.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.StoreId = p.StoreId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnChangedCause uint32

const (
	_ OnChangedCause = iota

	OnChangedCause_EVICTED
	OnChangedCause_EXPIRED
	OnChangedCause_EXPLICIT
	OnChangedCause_EXPIRED_OVERWRITE
	OnChangedCause_OVERWRITE
)

func (OnChangedCause) FromRef(str js.Ref) OnChangedCause {
	return OnChangedCause(bindings.ConstOfOnChangedCause(str))
}

func (x OnChangedCause) String() (string, bool) {
	switch x {
	case OnChangedCause_EVICTED:
		return "evicted", true
	case OnChangedCause_EXPIRED:
		return "expired", true
	case OnChangedCause_EXPLICIT:
		return "explicit", true
	case OnChangedCause_EXPIRED_OVERWRITE:
		return "expired_overwrite", true
	case OnChangedCause_OVERWRITE:
		return "overwrite", true
	default:
		return "", false
	}
}

type OnChangedArgChangeInfo struct {
	// Cause is "OnChangedArgChangeInfo.cause"
	//
	// Required
	Cause OnChangedCause
	// Cookie is "OnChangedArgChangeInfo.cookie"
	//
	// Required
	//
	// NOTE: Cookie.FFI_USE MUST be set to true to get Cookie used.
	Cookie Cookie
	// Removed is "OnChangedArgChangeInfo.removed"
	//
	// Required
	Removed bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnChangedArgChangeInfo with all fields set.
func (p OnChangedArgChangeInfo) FromRef(ref js.Ref) OnChangedArgChangeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnChangedArgChangeInfo in the application heap.
func (p OnChangedArgChangeInfo) New() js.Ref {
	return bindings.OnChangedArgChangeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnChangedArgChangeInfo) UpdateFrom(ref js.Ref) {
	bindings.OnChangedArgChangeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnChangedArgChangeInfo) Update(ref js.Ref) {
	bindings.OnChangedArgChangeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnChangedArgChangeInfo) FreeMembers(recursive bool) {
	if recursive {
		p.Cookie.FreeMembers(true)
	}
}

type RemoveReturnType struct {
	// Name is "RemoveReturnType.name"
	//
	// Required
	Name js.String
	// StoreId is "RemoveReturnType.storeId"
	//
	// Required
	StoreId js.String
	// Url is "RemoveReturnType.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemoveReturnType with all fields set.
func (p RemoveReturnType) FromRef(ref js.Ref) RemoveReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemoveReturnType in the application heap.
func (p RemoveReturnType) New() js.Ref {
	return bindings.RemoveReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemoveReturnType) UpdateFrom(ref js.Ref) {
	bindings.RemoveReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemoveReturnType) Update(ref js.Ref) {
	bindings.RemoveReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemoveReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.StoreId.Ref(),
		p.Url.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.StoreId = p.StoreId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type SetArgDetails struct {
	// Domain is "SetArgDetails.domain"
	//
	// Optional
	Domain js.String
	// ExpirationDate is "SetArgDetails.expirationDate"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExpirationDate MUST be set to true to make this field effective.
	ExpirationDate float64
	// HttpOnly is "SetArgDetails.httpOnly"
	//
	// Optional
	//
	// NOTE: FFI_USE_HttpOnly MUST be set to true to make this field effective.
	HttpOnly bool
	// Name is "SetArgDetails.name"
	//
	// Optional
	Name js.String
	// Path is "SetArgDetails.path"
	//
	// Optional
	Path js.String
	// SameSite is "SetArgDetails.sameSite"
	//
	// Optional
	SameSite SameSiteStatus
	// Secure is "SetArgDetails.secure"
	//
	// Optional
	//
	// NOTE: FFI_USE_Secure MUST be set to true to make this field effective.
	Secure bool
	// StoreId is "SetArgDetails.storeId"
	//
	// Optional
	StoreId js.String
	// Url is "SetArgDetails.url"
	//
	// Required
	Url js.String
	// Value is "SetArgDetails.value"
	//
	// Optional
	Value js.String

	FFI_USE_ExpirationDate bool // for ExpirationDate.
	FFI_USE_HttpOnly       bool // for HttpOnly.
	FFI_USE_Secure         bool // for Secure.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetArgDetails with all fields set.
func (p SetArgDetails) FromRef(ref js.Ref) SetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetArgDetails in the application heap.
func (p SetArgDetails) New() js.Ref {
	return bindings.SetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetArgDetails) Update(ref js.Ref) {
	bindings.SetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Name.Ref(),
		p.Path.Ref(),
		p.StoreId.Ref(),
		p.Url.Ref(),
		p.Value.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Path = p.Path.FromRef(js.Undefined)
	p.StoreId = p.StoreId.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

// HasFuncGet returns true if the function "WEBEXT.cookies.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.cookies.get".
func FuncGet() (fn js.Func[func(details CookieDetails) js.Promise[Cookie]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.cookies.get" directly.
func Get(details CookieDetails) (ret js.Promise[Cookie]) {
	bindings.CallGet(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGet calls the function "WEBEXT.cookies.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(details CookieDetails) (ret js.Promise[Cookie], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetAll returns true if the function "WEBEXT.cookies.getAll" exists.
func HasFuncGetAll() bool {
	return js.True == bindings.HasFuncGetAll()
}

// FuncGetAll returns the function "WEBEXT.cookies.getAll".
func FuncGetAll() (fn js.Func[func(details GetAllArgDetails) js.Promise[js.Array[Cookie]]]) {
	bindings.FuncGetAll(
		js.Pointer(&fn),
	)
	return
}

// GetAll calls the function "WEBEXT.cookies.getAll" directly.
func GetAll(details GetAllArgDetails) (ret js.Promise[js.Array[Cookie]]) {
	bindings.CallGetAll(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetAll calls the function "WEBEXT.cookies.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAll(details GetAllArgDetails) (ret js.Promise[js.Array[Cookie]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAll(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetAllCookieStores returns true if the function "WEBEXT.cookies.getAllCookieStores" exists.
func HasFuncGetAllCookieStores() bool {
	return js.True == bindings.HasFuncGetAllCookieStores()
}

// FuncGetAllCookieStores returns the function "WEBEXT.cookies.getAllCookieStores".
func FuncGetAllCookieStores() (fn js.Func[func() js.Promise[js.Array[CookieStore]]]) {
	bindings.FuncGetAllCookieStores(
		js.Pointer(&fn),
	)
	return
}

// GetAllCookieStores calls the function "WEBEXT.cookies.getAllCookieStores" directly.
func GetAllCookieStores() (ret js.Promise[js.Array[CookieStore]]) {
	bindings.CallGetAllCookieStores(
		js.Pointer(&ret),
	)

	return
}

// TryGetAllCookieStores calls the function "WEBEXT.cookies.getAllCookieStores"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAllCookieStores() (ret js.Promise[js.Array[CookieStore]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAllCookieStores(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnChangedEventCallbackFunc func(this js.Ref, changeInfo *OnChangedArgChangeInfo) js.Ref

func (fn OnChangedEventCallbackFunc) Register() js.Func[func(changeInfo *OnChangedArgChangeInfo)] {
	return js.RegisterCallback[func(changeInfo *OnChangedArgChangeInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnChangedArgChangeInfo
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

type OnChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, changeInfo *OnChangedArgChangeInfo) js.Ref
	Arg T
}

func (cb *OnChangedEventCallback[T]) Register() js.Func[func(changeInfo *OnChangedArgChangeInfo)] {
	return js.RegisterCallback[func(changeInfo *OnChangedArgChangeInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnChangedArgChangeInfo
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

// HasFuncOnChanged returns true if the function "WEBEXT.cookies.onChanged.addListener" exists.
func HasFuncOnChanged() bool {
	return js.True == bindings.HasFuncOnChanged()
}

// FuncOnChanged returns the function "WEBEXT.cookies.onChanged.addListener".
func FuncOnChanged() (fn js.Func[func(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)])]) {
	bindings.FuncOnChanged(
		js.Pointer(&fn),
	)
	return
}

// OnChanged calls the function "WEBEXT.cookies.onChanged.addListener" directly.
func OnChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret js.Void) {
	bindings.CallOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnChanged calls the function "WEBEXT.cookies.onChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffChanged returns true if the function "WEBEXT.cookies.onChanged.removeListener" exists.
func HasFuncOffChanged() bool {
	return js.True == bindings.HasFuncOffChanged()
}

// FuncOffChanged returns the function "WEBEXT.cookies.onChanged.removeListener".
func FuncOffChanged() (fn js.Func[func(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)])]) {
	bindings.FuncOffChanged(
		js.Pointer(&fn),
	)
	return
}

// OffChanged calls the function "WEBEXT.cookies.onChanged.removeListener" directly.
func OffChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret js.Void) {
	bindings.CallOffChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffChanged calls the function "WEBEXT.cookies.onChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnChanged returns true if the function "WEBEXT.cookies.onChanged.hasListener" exists.
func HasFuncHasOnChanged() bool {
	return js.True == bindings.HasFuncHasOnChanged()
}

// FuncHasOnChanged returns the function "WEBEXT.cookies.onChanged.hasListener".
func FuncHasOnChanged() (fn js.Func[func(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) bool]) {
	bindings.FuncHasOnChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnChanged calls the function "WEBEXT.cookies.onChanged.hasListener" directly.
func HasOnChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret bool) {
	bindings.CallHasOnChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnChanged calls the function "WEBEXT.cookies.onChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnChanged(callback js.Func[func(changeInfo *OnChangedArgChangeInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.cookies.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.cookies.remove".
func FuncRemove() (fn js.Func[func(details CookieDetails) js.Promise[RemoveReturnType]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.cookies.remove" directly.
func Remove(details CookieDetails) (ret js.Promise[RemoveReturnType]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryRemove calls the function "WEBEXT.cookies.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(details CookieDetails) (ret js.Promise[RemoveReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSet returns true if the function "WEBEXT.cookies.set" exists.
func HasFuncSet() bool {
	return js.True == bindings.HasFuncSet()
}

// FuncSet returns the function "WEBEXT.cookies.set".
func FuncSet() (fn js.Func[func(details SetArgDetails) js.Promise[Cookie]]) {
	bindings.FuncSet(
		js.Pointer(&fn),
	)
	return
}

// Set calls the function "WEBEXT.cookies.set" directly.
func Set(details SetArgDetails) (ret js.Promise[Cookie]) {
	bindings.CallSet(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySet calls the function "WEBEXT.cookies.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySet(details SetArgDetails) (ret js.Promise[Cookie], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySet(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

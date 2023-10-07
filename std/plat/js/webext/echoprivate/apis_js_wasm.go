// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package echoprivate

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/echoprivate/bindings"
)

type GetUserConsentArgConsentRequester struct {
	// Origin is "GetUserConsentArgConsentRequester.origin"
	//
	// Required
	Origin js.String
	// ServiceName is "GetUserConsentArgConsentRequester.serviceName"
	//
	// Required
	ServiceName js.String
	// TabId is "GetUserConsentArgConsentRequester.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetUserConsentArgConsentRequester with all fields set.
func (p GetUserConsentArgConsentRequester) FromRef(ref js.Ref) GetUserConsentArgConsentRequester {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetUserConsentArgConsentRequester in the application heap.
func (p GetUserConsentArgConsentRequester) New() js.Ref {
	return bindings.GetUserConsentArgConsentRequesterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetUserConsentArgConsentRequester) UpdateFrom(ref js.Ref) {
	bindings.GetUserConsentArgConsentRequesterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetUserConsentArgConsentRequester) Update(ref js.Ref) {
	bindings.GetUserConsentArgConsentRequesterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetUserConsentArgConsentRequester) FreeMembers(recursive bool) {
	js.Free(
		p.Origin.Ref(),
		p.ServiceName.Ref(),
	)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.ServiceName = p.ServiceName.FromRef(js.Undefined)
}

// HasFuncGetOfferInfo returns true if the function "WEBEXT.echoPrivate.getOfferInfo" exists.
func HasFuncGetOfferInfo() bool {
	return js.True == bindings.HasFuncGetOfferInfo()
}

// FuncGetOfferInfo returns the function "WEBEXT.echoPrivate.getOfferInfo".
func FuncGetOfferInfo() (fn js.Func[func(id js.String) js.Promise[js.Any]]) {
	bindings.FuncGetOfferInfo(
		js.Pointer(&fn),
	)
	return
}

// GetOfferInfo calls the function "WEBEXT.echoPrivate.getOfferInfo" directly.
func GetOfferInfo(id js.String) (ret js.Promise[js.Any]) {
	bindings.CallGetOfferInfo(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetOfferInfo calls the function "WEBEXT.echoPrivate.getOfferInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOfferInfo(id js.String) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOfferInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetOobeTimestamp returns true if the function "WEBEXT.echoPrivate.getOobeTimestamp" exists.
func HasFuncGetOobeTimestamp() bool {
	return js.True == bindings.HasFuncGetOobeTimestamp()
}

// FuncGetOobeTimestamp returns the function "WEBEXT.echoPrivate.getOobeTimestamp".
func FuncGetOobeTimestamp() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetOobeTimestamp(
		js.Pointer(&fn),
	)
	return
}

// GetOobeTimestamp calls the function "WEBEXT.echoPrivate.getOobeTimestamp" directly.
func GetOobeTimestamp() (ret js.Promise[js.String]) {
	bindings.CallGetOobeTimestamp(
		js.Pointer(&ret),
	)

	return
}

// TryGetOobeTimestamp calls the function "WEBEXT.echoPrivate.getOobeTimestamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOobeTimestamp() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOobeTimestamp(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRegistrationCode returns true if the function "WEBEXT.echoPrivate.getRegistrationCode" exists.
func HasFuncGetRegistrationCode() bool {
	return js.True == bindings.HasFuncGetRegistrationCode()
}

// FuncGetRegistrationCode returns the function "WEBEXT.echoPrivate.getRegistrationCode".
func FuncGetRegistrationCode() (fn js.Func[func(typ js.String) js.Promise[js.String]]) {
	bindings.FuncGetRegistrationCode(
		js.Pointer(&fn),
	)
	return
}

// GetRegistrationCode calls the function "WEBEXT.echoPrivate.getRegistrationCode" directly.
func GetRegistrationCode(typ js.String) (ret js.Promise[js.String]) {
	bindings.CallGetRegistrationCode(
		js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetRegistrationCode calls the function "WEBEXT.echoPrivate.getRegistrationCode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRegistrationCode(typ js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRegistrationCode(
		js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncGetUserConsent returns true if the function "WEBEXT.echoPrivate.getUserConsent" exists.
func HasFuncGetUserConsent() bool {
	return js.True == bindings.HasFuncGetUserConsent()
}

// FuncGetUserConsent returns the function "WEBEXT.echoPrivate.getUserConsent".
func FuncGetUserConsent() (fn js.Func[func(consentRequester GetUserConsentArgConsentRequester) js.Promise[js.Boolean]]) {
	bindings.FuncGetUserConsent(
		js.Pointer(&fn),
	)
	return
}

// GetUserConsent calls the function "WEBEXT.echoPrivate.getUserConsent" directly.
func GetUserConsent(consentRequester GetUserConsentArgConsentRequester) (ret js.Promise[js.Boolean]) {
	bindings.CallGetUserConsent(
		js.Pointer(&ret),
		js.Pointer(&consentRequester),
	)

	return
}

// TryGetUserConsent calls the function "WEBEXT.echoPrivate.getUserConsent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserConsent(consentRequester GetUserConsentArgConsentRequester) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserConsent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&consentRequester),
	)

	return
}

// HasFuncSetOfferInfo returns true if the function "WEBEXT.echoPrivate.setOfferInfo" exists.
func HasFuncSetOfferInfo() bool {
	return js.True == bindings.HasFuncSetOfferInfo()
}

// FuncSetOfferInfo returns the function "WEBEXT.echoPrivate.setOfferInfo".
func FuncSetOfferInfo() (fn js.Func[func(id js.String, offerInfo js.Any)]) {
	bindings.FuncSetOfferInfo(
		js.Pointer(&fn),
	)
	return
}

// SetOfferInfo calls the function "WEBEXT.echoPrivate.setOfferInfo" directly.
func SetOfferInfo(id js.String, offerInfo js.Any) (ret js.Void) {
	bindings.CallSetOfferInfo(
		js.Pointer(&ret),
		id.Ref(),
		offerInfo.Ref(),
	)

	return
}

// TrySetOfferInfo calls the function "WEBEXT.echoPrivate.setOfferInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetOfferInfo(id js.String, offerInfo js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetOfferInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		offerInfo.Ref(),
	)

	return
}

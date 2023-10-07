// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package privacy

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/privacy/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/types"
)

type IPHandlingPolicy uint32

const (
	_ IPHandlingPolicy = iota

	IPHandlingPolicy_DEFAULT
	IPHandlingPolicy_DEFAULT_PUBLIC_AND_PRIVATE_INTERFACES
	IPHandlingPolicy_DEFAULT_PUBLIC_INTERFACE_ONLY
	IPHandlingPolicy_DISABLE_NON_PROXIED_UDP
)

func (IPHandlingPolicy) FromRef(str js.Ref) IPHandlingPolicy {
	return IPHandlingPolicy(bindings.ConstOfIPHandlingPolicy(str))
}

func (x IPHandlingPolicy) String() (string, bool) {
	switch x {
	case IPHandlingPolicy_DEFAULT:
		return "default", true
	case IPHandlingPolicy_DEFAULT_PUBLIC_AND_PRIVATE_INTERFACES:
		return "default_public_and_private_interfaces", true
	case IPHandlingPolicy_DEFAULT_PUBLIC_INTERFACE_ONLY:
		return "default_public_interface_only", true
	case IPHandlingPolicy_DISABLE_NON_PROXIED_UDP:
		return "disable_non_proxied_udp", true
	default:
		return "", false
	}
}

type NetworkProperty struct {
	// NetworkPredictionEnabled is "NetworkProperty.networkPredictionEnabled"
	//
	// Required
	NetworkPredictionEnabled types.ChromeSetting
	// WebRTCIPHandlingPolicy is "NetworkProperty.webRTCIPHandlingPolicy"
	//
	// Required
	WebRTCIPHandlingPolicy types.ChromeSetting

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkProperty with all fields set.
func (p NetworkProperty) FromRef(ref js.Ref) NetworkProperty {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkProperty in the application heap.
func (p NetworkProperty) New() js.Ref {
	return bindings.NetworkPropertyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkProperty) UpdateFrom(ref js.Ref) {
	bindings.NetworkPropertyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkProperty) Update(ref js.Ref) {
	bindings.NetworkPropertyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkProperty) FreeMembers(recursive bool) {
	js.Free(
		p.NetworkPredictionEnabled.Ref(),
		p.WebRTCIPHandlingPolicy.Ref(),
	)
	p.NetworkPredictionEnabled = p.NetworkPredictionEnabled.FromRef(js.Undefined)
	p.WebRTCIPHandlingPolicy = p.WebRTCIPHandlingPolicy.FromRef(js.Undefined)
}

type ServicesProperty struct {
	// AlternateErrorPagesEnabled is "ServicesProperty.alternateErrorPagesEnabled"
	//
	// Required
	AlternateErrorPagesEnabled types.ChromeSetting
	// AutofillAddressEnabled is "ServicesProperty.autofillAddressEnabled"
	//
	// Required
	AutofillAddressEnabled types.ChromeSetting
	// AutofillCreditCardEnabled is "ServicesProperty.autofillCreditCardEnabled"
	//
	// Required
	AutofillCreditCardEnabled types.ChromeSetting
	// AutofillEnabled is "ServicesProperty.autofillEnabled"
	//
	// Required
	AutofillEnabled types.ChromeSetting
	// PasswordSavingEnabled is "ServicesProperty.passwordSavingEnabled"
	//
	// Required
	PasswordSavingEnabled types.ChromeSetting
	// SafeBrowsingEnabled is "ServicesProperty.safeBrowsingEnabled"
	//
	// Required
	SafeBrowsingEnabled types.ChromeSetting
	// SafeBrowsingExtendedReportingEnabled is "ServicesProperty.safeBrowsingExtendedReportingEnabled"
	//
	// Required
	SafeBrowsingExtendedReportingEnabled types.ChromeSetting
	// SearchSuggestEnabled is "ServicesProperty.searchSuggestEnabled"
	//
	// Required
	SearchSuggestEnabled types.ChromeSetting
	// SpellingServiceEnabled is "ServicesProperty.spellingServiceEnabled"
	//
	// Required
	SpellingServiceEnabled types.ChromeSetting
	// TranslationServiceEnabled is "ServicesProperty.translationServiceEnabled"
	//
	// Required
	TranslationServiceEnabled types.ChromeSetting

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ServicesProperty with all fields set.
func (p ServicesProperty) FromRef(ref js.Ref) ServicesProperty {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ServicesProperty in the application heap.
func (p ServicesProperty) New() js.Ref {
	return bindings.ServicesPropertyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ServicesProperty) UpdateFrom(ref js.Ref) {
	bindings.ServicesPropertyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ServicesProperty) Update(ref js.Ref) {
	bindings.ServicesPropertyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ServicesProperty) FreeMembers(recursive bool) {
	js.Free(
		p.AlternateErrorPagesEnabled.Ref(),
		p.AutofillAddressEnabled.Ref(),
		p.AutofillCreditCardEnabled.Ref(),
		p.AutofillEnabled.Ref(),
		p.PasswordSavingEnabled.Ref(),
		p.SafeBrowsingEnabled.Ref(),
		p.SafeBrowsingExtendedReportingEnabled.Ref(),
		p.SearchSuggestEnabled.Ref(),
		p.SpellingServiceEnabled.Ref(),
		p.TranslationServiceEnabled.Ref(),
	)
	p.AlternateErrorPagesEnabled = p.AlternateErrorPagesEnabled.FromRef(js.Undefined)
	p.AutofillAddressEnabled = p.AutofillAddressEnabled.FromRef(js.Undefined)
	p.AutofillCreditCardEnabled = p.AutofillCreditCardEnabled.FromRef(js.Undefined)
	p.AutofillEnabled = p.AutofillEnabled.FromRef(js.Undefined)
	p.PasswordSavingEnabled = p.PasswordSavingEnabled.FromRef(js.Undefined)
	p.SafeBrowsingEnabled = p.SafeBrowsingEnabled.FromRef(js.Undefined)
	p.SafeBrowsingExtendedReportingEnabled = p.SafeBrowsingExtendedReportingEnabled.FromRef(js.Undefined)
	p.SearchSuggestEnabled = p.SearchSuggestEnabled.FromRef(js.Undefined)
	p.SpellingServiceEnabled = p.SpellingServiceEnabled.FromRef(js.Undefined)
	p.TranslationServiceEnabled = p.TranslationServiceEnabled.FromRef(js.Undefined)
}

type WebsitesProperty struct {
	// AdMeasurementEnabled is "WebsitesProperty.adMeasurementEnabled"
	//
	// Required
	AdMeasurementEnabled types.ChromeSetting
	// DoNotTrackEnabled is "WebsitesProperty.doNotTrackEnabled"
	//
	// Required
	DoNotTrackEnabled types.ChromeSetting
	// FledgeEnabled is "WebsitesProperty.fledgeEnabled"
	//
	// Required
	FledgeEnabled types.ChromeSetting
	// HyperlinkAuditingEnabled is "WebsitesProperty.hyperlinkAuditingEnabled"
	//
	// Required
	HyperlinkAuditingEnabled types.ChromeSetting
	// ProtectedContentEnabled is "WebsitesProperty.protectedContentEnabled"
	//
	// Required
	ProtectedContentEnabled types.ChromeSetting
	// ReferrersEnabled is "WebsitesProperty.referrersEnabled"
	//
	// Required
	ReferrersEnabled types.ChromeSetting
	// ThirdPartyCookiesAllowed is "WebsitesProperty.thirdPartyCookiesAllowed"
	//
	// Required
	ThirdPartyCookiesAllowed types.ChromeSetting
	// TopicsEnabled is "WebsitesProperty.topicsEnabled"
	//
	// Required
	TopicsEnabled types.ChromeSetting

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebsitesProperty with all fields set.
func (p WebsitesProperty) FromRef(ref js.Ref) WebsitesProperty {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebsitesProperty in the application heap.
func (p WebsitesProperty) New() js.Ref {
	return bindings.WebsitesPropertyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebsitesProperty) UpdateFrom(ref js.Ref) {
	bindings.WebsitesPropertyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebsitesProperty) Update(ref js.Ref) {
	bindings.WebsitesPropertyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebsitesProperty) FreeMembers(recursive bool) {
	js.Free(
		p.AdMeasurementEnabled.Ref(),
		p.DoNotTrackEnabled.Ref(),
		p.FledgeEnabled.Ref(),
		p.HyperlinkAuditingEnabled.Ref(),
		p.ProtectedContentEnabled.Ref(),
		p.ReferrersEnabled.Ref(),
		p.ThirdPartyCookiesAllowed.Ref(),
		p.TopicsEnabled.Ref(),
	)
	p.AdMeasurementEnabled = p.AdMeasurementEnabled.FromRef(js.Undefined)
	p.DoNotTrackEnabled = p.DoNotTrackEnabled.FromRef(js.Undefined)
	p.FledgeEnabled = p.FledgeEnabled.FromRef(js.Undefined)
	p.HyperlinkAuditingEnabled = p.HyperlinkAuditingEnabled.FromRef(js.Undefined)
	p.ProtectedContentEnabled = p.ProtectedContentEnabled.FromRef(js.Undefined)
	p.ReferrersEnabled = p.ReferrersEnabled.FromRef(js.Undefined)
	p.ThirdPartyCookiesAllowed = p.ThirdPartyCookiesAllowed.FromRef(js.Undefined)
	p.TopicsEnabled = p.TopicsEnabled.FromRef(js.Undefined)
}

// Network returns the value of property "WEBEXT.privacy.network".
//
// The returned bool will be false if there is no such property.
func Network() (ret NetworkProperty, ok bool) {
	ok = js.True == bindings.GetNetwork(
		js.Pointer(&ret),
	)

	return
}

// SetNetwork sets the value of property "WEBEXT.privacy.network" to val.
//
// It returns false if the property cannot be set.
func SetNetwork(val NetworkProperty) bool {
	return js.True == bindings.SetNetwork(
		js.Pointer(&val))
}

// Services returns the value of property "WEBEXT.privacy.services".
//
// The returned bool will be false if there is no such property.
func Services() (ret ServicesProperty, ok bool) {
	ok = js.True == bindings.GetServices(
		js.Pointer(&ret),
	)

	return
}

// SetServices sets the value of property "WEBEXT.privacy.services" to val.
//
// It returns false if the property cannot be set.
func SetServices(val ServicesProperty) bool {
	return js.True == bindings.SetServices(
		js.Pointer(&val))
}

// Websites returns the value of property "WEBEXT.privacy.websites".
//
// The returned bool will be false if there is no such property.
func Websites() (ret WebsitesProperty, ok bool) {
	ok = js.True == bindings.GetWebsites(
		js.Pointer(&ret),
	)

	return
}

// SetWebsites sets the value of property "WEBEXT.privacy.websites" to val.
//
// It returns false if the property cannot be set.
func SetWebsites(val WebsitesProperty) bool {
	return js.True == bindings.SetWebsites(
		js.Pointer(&val))
}

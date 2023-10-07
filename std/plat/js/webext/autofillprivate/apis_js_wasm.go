// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package autofillprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/autofillprivate/bindings"
)

type AccountInfo struct {
	// Email is "AccountInfo.email"
	//
	// Optional
	Email js.String
	// IsSyncEnabledForAutofillProfiles is "AccountInfo.isSyncEnabledForAutofillProfiles"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsSyncEnabledForAutofillProfiles MUST be set to true to make this field effective.
	IsSyncEnabledForAutofillProfiles bool
	// IsEligibleForAddressAccountStorage is "AccountInfo.isEligibleForAddressAccountStorage"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsEligibleForAddressAccountStorage MUST be set to true to make this field effective.
	IsEligibleForAddressAccountStorage bool

	FFI_USE_IsSyncEnabledForAutofillProfiles   bool // for IsSyncEnabledForAutofillProfiles.
	FFI_USE_IsEligibleForAddressAccountStorage bool // for IsEligibleForAddressAccountStorage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccountInfo with all fields set.
func (p AccountInfo) FromRef(ref js.Ref) AccountInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AccountInfo in the application heap.
func (p AccountInfo) New() js.Ref {
	return bindings.AccountInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AccountInfo) UpdateFrom(ref js.Ref) {
	bindings.AccountInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AccountInfo) Update(ref js.Ref) {
	bindings.AccountInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AccountInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Email.Ref(),
	)
	p.Email = p.Email.FromRef(js.Undefined)
}

type ServerFieldType uint32

const (
	_ ServerFieldType = iota

	ServerFieldType_NO_SERVER_DATA
	ServerFieldType_UNKNOWN_TYPE
	ServerFieldType_EMPTY_TYPE
	ServerFieldType_NAME_FIRST
	ServerFieldType_NAME_MIDDLE
	ServerFieldType_NAME_LAST
	ServerFieldType_NAME_MIDDLE_INITIAL
	ServerFieldType_NAME_FULL
	ServerFieldType_NAME_SUFFIX
	ServerFieldType_EMAIL_ADDRESS
	ServerFieldType_PHONE_HOME_NUMBER
	ServerFieldType_PHONE_HOME_CITY_CODE
	ServerFieldType_PHONE_HOME_COUNTRY_CODE
	ServerFieldType_PHONE_HOME_CITY_AND_NUMBER
	ServerFieldType_PHONE_HOME_WHOLE_NUMBER
	ServerFieldType_ADDRESS_HOME_LINE1
	ServerFieldType_ADDRESS_HOME_LINE2
	ServerFieldType_ADDRESS_HOME_APT_NUM
	ServerFieldType_ADDRESS_HOME_CITY
	ServerFieldType_ADDRESS_HOME_STATE
	ServerFieldType_ADDRESS_HOME_ZIP
	ServerFieldType_ADDRESS_HOME_COUNTRY
	ServerFieldType_CREDIT_CARD_NAME_FULL
	ServerFieldType_CREDIT_CARD_NUMBER
	ServerFieldType_CREDIT_CARD_EXP_MONTH
	ServerFieldType_CREDIT_CARD_EXP_2_DIGIT_YEAR
	ServerFieldType_CREDIT_CARD_EXP_4_DIGIT_YEAR
	ServerFieldType_CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR
	ServerFieldType_CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR
	ServerFieldType_CREDIT_CARD_TYPE
	ServerFieldType_CREDIT_CARD_VERIFICATION_CODE
	ServerFieldType_COMPANY_NAME
	ServerFieldType_FIELD_WITH_DEFAULT_VALUE
	ServerFieldType_MERCHANT_EMAIL_SIGNUP
	ServerFieldType_MERCHANT_PROMO_CODE
	ServerFieldType_PASSWORD
	ServerFieldType_ACCOUNT_CREATION_PASSWORD
	ServerFieldType_ADDRESS_HOME_STREET_ADDRESS
	ServerFieldType_ADDRESS_HOME_SORTING_CODE
	ServerFieldType_ADDRESS_HOME_DEPENDENT_LOCALITY
	ServerFieldType_ADDRESS_HOME_LINE3
	ServerFieldType_NOT_ACCOUNT_CREATION_PASSWORD
	ServerFieldType_USERNAME
	ServerFieldType_USERNAME_AND_EMAIL_ADDRESS
	ServerFieldType_NEW_PASSWORD
	ServerFieldType_PROBABLY_NEW_PASSWORD
	ServerFieldType_NOT_NEW_PASSWORD
	ServerFieldType_CREDIT_CARD_NAME_FIRST
	ServerFieldType_CREDIT_CARD_NAME_LAST
	ServerFieldType_PHONE_HOME_EXTENSION
	ServerFieldType_CONFIRMATION_PASSWORD
	ServerFieldType_AMBIGUOUS_TYPE
	ServerFieldType_SEARCH_TERM
	ServerFieldType_PRICE
	ServerFieldType_NOT_PASSWORD
	ServerFieldType_SINGLE_USERNAME
	ServerFieldType_NOT_USERNAME
	ServerFieldType_UPI_VPA
	ServerFieldType_ADDRESS_HOME_STREET_NAME
	ServerFieldType_ADDRESS_HOME_HOUSE_NUMBER
	ServerFieldType_ADDRESS_HOME_SUBPREMISE
	ServerFieldType_ADDRESS_HOME_OTHER_SUBUNIT
	ServerFieldType_NAME_LAST_FIRST
	ServerFieldType_NAME_LAST_CONJUNCTION
	ServerFieldType_NAME_LAST_SECOND
	ServerFieldType_NAME_HONORIFIC_PREFIX
	ServerFieldType_ADDRESS_HOME_ADDRESS
	ServerFieldType_ADDRESS_HOME_ADDRESS_WITH_NAME
	ServerFieldType_ADDRESS_HOME_FLOOR
	ServerFieldType_NAME_FULL_WITH_HONORIFIC_PREFIX
	ServerFieldType_BIRTHDATE_DAY
	ServerFieldType_BIRTHDATE_MONTH
	ServerFieldType_BIRTHDATE_4_DIGIT_YEAR
	ServerFieldType_PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX
	ServerFieldType_PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX
	ServerFieldType_PHONE_HOME_NUMBER_PREFIX
	ServerFieldType_PHONE_HOME_NUMBER_SUFFIX
	ServerFieldType_IBAN_VALUE
	ServerFieldType_CREDIT_CARD_STANDALONE_VERIFICATION_CODE
	ServerFieldType_NUMERIC_QUANTITY
	ServerFieldType_ONE_TIME_CODE
	ServerFieldType_DELIVERY_INSTRUCTIONS
	ServerFieldType_ADDRESS_HOME_OVERFLOW
	ServerFieldType_ADDRESS_HOME_LANDMARK
	ServerFieldType_ADDRESS_HOME_OVERFLOW_AND_LANDMARK
	ServerFieldType_ADDRESS_HOME_ADMIN_LEVEL2
	ServerFieldType_ADDRESS_HOME_STREET_LOCATION
	ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS
	ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK
	ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_1
	ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_2
	ServerFieldType_SINGLE_USERNAME_FORGOT_PASSWORD
	ServerFieldType_MAX_VALID_FIELD_TYPE
)

func (ServerFieldType) FromRef(str js.Ref) ServerFieldType {
	return ServerFieldType(bindings.ConstOfServerFieldType(str))
}

func (x ServerFieldType) String() (string, bool) {
	switch x {
	case ServerFieldType_NO_SERVER_DATA:
		return "NO_SERVER_DATA", true
	case ServerFieldType_UNKNOWN_TYPE:
		return "UNKNOWN_TYPE", true
	case ServerFieldType_EMPTY_TYPE:
		return "EMPTY_TYPE", true
	case ServerFieldType_NAME_FIRST:
		return "NAME_FIRST", true
	case ServerFieldType_NAME_MIDDLE:
		return "NAME_MIDDLE", true
	case ServerFieldType_NAME_LAST:
		return "NAME_LAST", true
	case ServerFieldType_NAME_MIDDLE_INITIAL:
		return "NAME_MIDDLE_INITIAL", true
	case ServerFieldType_NAME_FULL:
		return "NAME_FULL", true
	case ServerFieldType_NAME_SUFFIX:
		return "NAME_SUFFIX", true
	case ServerFieldType_EMAIL_ADDRESS:
		return "EMAIL_ADDRESS", true
	case ServerFieldType_PHONE_HOME_NUMBER:
		return "PHONE_HOME_NUMBER", true
	case ServerFieldType_PHONE_HOME_CITY_CODE:
		return "PHONE_HOME_CITY_CODE", true
	case ServerFieldType_PHONE_HOME_COUNTRY_CODE:
		return "PHONE_HOME_COUNTRY_CODE", true
	case ServerFieldType_PHONE_HOME_CITY_AND_NUMBER:
		return "PHONE_HOME_CITY_AND_NUMBER", true
	case ServerFieldType_PHONE_HOME_WHOLE_NUMBER:
		return "PHONE_HOME_WHOLE_NUMBER", true
	case ServerFieldType_ADDRESS_HOME_LINE1:
		return "ADDRESS_HOME_LINE1", true
	case ServerFieldType_ADDRESS_HOME_LINE2:
		return "ADDRESS_HOME_LINE2", true
	case ServerFieldType_ADDRESS_HOME_APT_NUM:
		return "ADDRESS_HOME_APT_NUM", true
	case ServerFieldType_ADDRESS_HOME_CITY:
		return "ADDRESS_HOME_CITY", true
	case ServerFieldType_ADDRESS_HOME_STATE:
		return "ADDRESS_HOME_STATE", true
	case ServerFieldType_ADDRESS_HOME_ZIP:
		return "ADDRESS_HOME_ZIP", true
	case ServerFieldType_ADDRESS_HOME_COUNTRY:
		return "ADDRESS_HOME_COUNTRY", true
	case ServerFieldType_CREDIT_CARD_NAME_FULL:
		return "CREDIT_CARD_NAME_FULL", true
	case ServerFieldType_CREDIT_CARD_NUMBER:
		return "CREDIT_CARD_NUMBER", true
	case ServerFieldType_CREDIT_CARD_EXP_MONTH:
		return "CREDIT_CARD_EXP_MONTH", true
	case ServerFieldType_CREDIT_CARD_EXP_2_DIGIT_YEAR:
		return "CREDIT_CARD_EXP_2_DIGIT_YEAR", true
	case ServerFieldType_CREDIT_CARD_EXP_4_DIGIT_YEAR:
		return "CREDIT_CARD_EXP_4_DIGIT_YEAR", true
	case ServerFieldType_CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR:
		return "CREDIT_CARD_EXP_DATE_2_DIGIT_YEAR", true
	case ServerFieldType_CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR:
		return "CREDIT_CARD_EXP_DATE_4_DIGIT_YEAR", true
	case ServerFieldType_CREDIT_CARD_TYPE:
		return "CREDIT_CARD_TYPE", true
	case ServerFieldType_CREDIT_CARD_VERIFICATION_CODE:
		return "CREDIT_CARD_VERIFICATION_CODE", true
	case ServerFieldType_COMPANY_NAME:
		return "COMPANY_NAME", true
	case ServerFieldType_FIELD_WITH_DEFAULT_VALUE:
		return "FIELD_WITH_DEFAULT_VALUE", true
	case ServerFieldType_MERCHANT_EMAIL_SIGNUP:
		return "MERCHANT_EMAIL_SIGNUP", true
	case ServerFieldType_MERCHANT_PROMO_CODE:
		return "MERCHANT_PROMO_CODE", true
	case ServerFieldType_PASSWORD:
		return "PASSWORD", true
	case ServerFieldType_ACCOUNT_CREATION_PASSWORD:
		return "ACCOUNT_CREATION_PASSWORD", true
	case ServerFieldType_ADDRESS_HOME_STREET_ADDRESS:
		return "ADDRESS_HOME_STREET_ADDRESS", true
	case ServerFieldType_ADDRESS_HOME_SORTING_CODE:
		return "ADDRESS_HOME_SORTING_CODE", true
	case ServerFieldType_ADDRESS_HOME_DEPENDENT_LOCALITY:
		return "ADDRESS_HOME_DEPENDENT_LOCALITY", true
	case ServerFieldType_ADDRESS_HOME_LINE3:
		return "ADDRESS_HOME_LINE3", true
	case ServerFieldType_NOT_ACCOUNT_CREATION_PASSWORD:
		return "NOT_ACCOUNT_CREATION_PASSWORD", true
	case ServerFieldType_USERNAME:
		return "USERNAME", true
	case ServerFieldType_USERNAME_AND_EMAIL_ADDRESS:
		return "USERNAME_AND_EMAIL_ADDRESS", true
	case ServerFieldType_NEW_PASSWORD:
		return "NEW_PASSWORD", true
	case ServerFieldType_PROBABLY_NEW_PASSWORD:
		return "PROBABLY_NEW_PASSWORD", true
	case ServerFieldType_NOT_NEW_PASSWORD:
		return "NOT_NEW_PASSWORD", true
	case ServerFieldType_CREDIT_CARD_NAME_FIRST:
		return "CREDIT_CARD_NAME_FIRST", true
	case ServerFieldType_CREDIT_CARD_NAME_LAST:
		return "CREDIT_CARD_NAME_LAST", true
	case ServerFieldType_PHONE_HOME_EXTENSION:
		return "PHONE_HOME_EXTENSION", true
	case ServerFieldType_CONFIRMATION_PASSWORD:
		return "CONFIRMATION_PASSWORD", true
	case ServerFieldType_AMBIGUOUS_TYPE:
		return "AMBIGUOUS_TYPE", true
	case ServerFieldType_SEARCH_TERM:
		return "SEARCH_TERM", true
	case ServerFieldType_PRICE:
		return "PRICE", true
	case ServerFieldType_NOT_PASSWORD:
		return "NOT_PASSWORD", true
	case ServerFieldType_SINGLE_USERNAME:
		return "SINGLE_USERNAME", true
	case ServerFieldType_NOT_USERNAME:
		return "NOT_USERNAME", true
	case ServerFieldType_UPI_VPA:
		return "UPI_VPA", true
	case ServerFieldType_ADDRESS_HOME_STREET_NAME:
		return "ADDRESS_HOME_STREET_NAME", true
	case ServerFieldType_ADDRESS_HOME_HOUSE_NUMBER:
		return "ADDRESS_HOME_HOUSE_NUMBER", true
	case ServerFieldType_ADDRESS_HOME_SUBPREMISE:
		return "ADDRESS_HOME_SUBPREMISE", true
	case ServerFieldType_ADDRESS_HOME_OTHER_SUBUNIT:
		return "ADDRESS_HOME_OTHER_SUBUNIT", true
	case ServerFieldType_NAME_LAST_FIRST:
		return "NAME_LAST_FIRST", true
	case ServerFieldType_NAME_LAST_CONJUNCTION:
		return "NAME_LAST_CONJUNCTION", true
	case ServerFieldType_NAME_LAST_SECOND:
		return "NAME_LAST_SECOND", true
	case ServerFieldType_NAME_HONORIFIC_PREFIX:
		return "NAME_HONORIFIC_PREFIX", true
	case ServerFieldType_ADDRESS_HOME_ADDRESS:
		return "ADDRESS_HOME_ADDRESS", true
	case ServerFieldType_ADDRESS_HOME_ADDRESS_WITH_NAME:
		return "ADDRESS_HOME_ADDRESS_WITH_NAME", true
	case ServerFieldType_ADDRESS_HOME_FLOOR:
		return "ADDRESS_HOME_FLOOR", true
	case ServerFieldType_NAME_FULL_WITH_HONORIFIC_PREFIX:
		return "NAME_FULL_WITH_HONORIFIC_PREFIX", true
	case ServerFieldType_BIRTHDATE_DAY:
		return "BIRTHDATE_DAY", true
	case ServerFieldType_BIRTHDATE_MONTH:
		return "BIRTHDATE_MONTH", true
	case ServerFieldType_BIRTHDATE_4_DIGIT_YEAR:
		return "BIRTHDATE_4_DIGIT_YEAR", true
	case ServerFieldType_PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX:
		return "PHONE_HOME_CITY_CODE_WITH_TRUNK_PREFIX", true
	case ServerFieldType_PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX:
		return "PHONE_HOME_CITY_AND_NUMBER_WITHOUT_TRUNK_PREFIX", true
	case ServerFieldType_PHONE_HOME_NUMBER_PREFIX:
		return "PHONE_HOME_NUMBER_PREFIX", true
	case ServerFieldType_PHONE_HOME_NUMBER_SUFFIX:
		return "PHONE_HOME_NUMBER_SUFFIX", true
	case ServerFieldType_IBAN_VALUE:
		return "IBAN_VALUE", true
	case ServerFieldType_CREDIT_CARD_STANDALONE_VERIFICATION_CODE:
		return "CREDIT_CARD_STANDALONE_VERIFICATION_CODE", true
	case ServerFieldType_NUMERIC_QUANTITY:
		return "NUMERIC_QUANTITY", true
	case ServerFieldType_ONE_TIME_CODE:
		return "ONE_TIME_CODE", true
	case ServerFieldType_DELIVERY_INSTRUCTIONS:
		return "DELIVERY_INSTRUCTIONS", true
	case ServerFieldType_ADDRESS_HOME_OVERFLOW:
		return "ADDRESS_HOME_OVERFLOW", true
	case ServerFieldType_ADDRESS_HOME_LANDMARK:
		return "ADDRESS_HOME_LANDMARK", true
	case ServerFieldType_ADDRESS_HOME_OVERFLOW_AND_LANDMARK:
		return "ADDRESS_HOME_OVERFLOW_AND_LANDMARK", true
	case ServerFieldType_ADDRESS_HOME_ADMIN_LEVEL2:
		return "ADDRESS_HOME_ADMIN_LEVEL2", true
	case ServerFieldType_ADDRESS_HOME_STREET_LOCATION:
		return "ADDRESS_HOME_STREET_LOCATION", true
	case ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS:
		return "ADDRESS_HOME_BETWEEN_STREETS", true
	case ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK:
		return "ADDRESS_HOME_BETWEEN_STREETS_OR_LANDMARK", true
	case ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_1:
		return "ADDRESS_HOME_BETWEEN_STREETS_1", true
	case ServerFieldType_ADDRESS_HOME_BETWEEN_STREETS_2:
		return "ADDRESS_HOME_BETWEEN_STREETS_2", true
	case ServerFieldType_SINGLE_USERNAME_FORGOT_PASSWORD:
		return "SINGLE_USERNAME_FORGOT_PASSWORD", true
	case ServerFieldType_MAX_VALID_FIELD_TYPE:
		return "MAX_VALID_FIELD_TYPE", true
	default:
		return "", false
	}
}

type AddressComponent struct {
	// Field is "AddressComponent.field"
	//
	// Optional
	Field ServerFieldType
	// FieldName is "AddressComponent.fieldName"
	//
	// Optional
	FieldName js.String
	// IsLongField is "AddressComponent.isLongField"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLongField MUST be set to true to make this field effective.
	IsLongField bool
	// IsRequired is "AddressComponent.isRequired"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRequired MUST be set to true to make this field effective.
	IsRequired bool
	// Placeholder is "AddressComponent.placeholder"
	//
	// Optional
	Placeholder js.String

	FFI_USE_IsLongField bool // for IsLongField.
	FFI_USE_IsRequired  bool // for IsRequired.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressComponent with all fields set.
func (p AddressComponent) FromRef(ref js.Ref) AddressComponent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddressComponent in the application heap.
func (p AddressComponent) New() js.Ref {
	return bindings.AddressComponentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressComponent) UpdateFrom(ref js.Ref) {
	bindings.AddressComponentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressComponent) Update(ref js.Ref) {
	bindings.AddressComponentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressComponent) FreeMembers(recursive bool) {
	js.Free(
		p.FieldName.Ref(),
		p.Placeholder.Ref(),
	)
	p.FieldName = p.FieldName.FromRef(js.Undefined)
	p.Placeholder = p.Placeholder.FromRef(js.Undefined)
}

type AddressComponentRow struct {
	// Row is "AddressComponentRow.row"
	//
	// Optional
	Row js.Array[AddressComponent]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressComponentRow with all fields set.
func (p AddressComponentRow) FromRef(ref js.Ref) AddressComponentRow {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddressComponentRow in the application heap.
func (p AddressComponentRow) New() js.Ref {
	return bindings.AddressComponentRowJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressComponentRow) UpdateFrom(ref js.Ref) {
	bindings.AddressComponentRowJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressComponentRow) Update(ref js.Ref) {
	bindings.AddressComponentRowJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressComponentRow) FreeMembers(recursive bool) {
	js.Free(
		p.Row.Ref(),
	)
	p.Row = p.Row.FromRef(js.Undefined)
}

type AddressComponents struct {
	// Components is "AddressComponents.components"
	//
	// Optional
	Components js.Array[AddressComponentRow]
	// LanguageCode is "AddressComponents.languageCode"
	//
	// Optional
	LanguageCode js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressComponents with all fields set.
func (p AddressComponents) FromRef(ref js.Ref) AddressComponents {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddressComponents in the application heap.
func (p AddressComponents) New() js.Ref {
	return bindings.AddressComponentsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressComponents) UpdateFrom(ref js.Ref) {
	bindings.AddressComponentsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressComponents) Update(ref js.Ref) {
	bindings.AddressComponentsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressComponents) FreeMembers(recursive bool) {
	js.Free(
		p.Components.Ref(),
		p.LanguageCode.Ref(),
	)
	p.Components = p.Components.FromRef(js.Undefined)
	p.LanguageCode = p.LanguageCode.FromRef(js.Undefined)
}

type AddressField struct {
	// Type is "AddressField.type"
	//
	// Optional
	Type ServerFieldType
	// Value is "AddressField.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressField with all fields set.
func (p AddressField) FromRef(ref js.Ref) AddressField {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddressField in the application heap.
func (p AddressField) New() js.Ref {
	return bindings.AddressFieldJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressField) UpdateFrom(ref js.Ref) {
	bindings.AddressFieldJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressField) Update(ref js.Ref) {
	bindings.AddressFieldJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressField) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type AddressSource uint32

const (
	_ AddressSource = iota

	AddressSource_LOCAL_OR_SYNCABLE
	AddressSource_ACCOUNT
)

func (AddressSource) FromRef(str js.Ref) AddressSource {
	return AddressSource(bindings.ConstOfAddressSource(str))
}

func (x AddressSource) String() (string, bool) {
	switch x {
	case AddressSource_LOCAL_OR_SYNCABLE:
		return "LOCAL_OR_SYNCABLE", true
	case AddressSource_ACCOUNT:
		return "ACCOUNT", true
	default:
		return "", false
	}
}

type AutofillMetadata struct {
	// SummaryLabel is "AutofillMetadata.summaryLabel"
	//
	// Optional
	SummaryLabel js.String
	// SummarySublabel is "AutofillMetadata.summarySublabel"
	//
	// Optional
	SummarySublabel js.String
	// Source is "AutofillMetadata.source"
	//
	// Optional
	Source AddressSource
	// IsLocal is "AutofillMetadata.isLocal"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLocal MUST be set to true to make this field effective.
	IsLocal bool
	// IsCached is "AutofillMetadata.isCached"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsCached MUST be set to true to make this field effective.
	IsCached bool
	// IsMigratable is "AutofillMetadata.isMigratable"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMigratable MUST be set to true to make this field effective.
	IsMigratable bool
	// IsVirtualCardEnrollmentEligible is "AutofillMetadata.isVirtualCardEnrollmentEligible"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsVirtualCardEnrollmentEligible MUST be set to true to make this field effective.
	IsVirtualCardEnrollmentEligible bool
	// IsVirtualCardEnrolled is "AutofillMetadata.isVirtualCardEnrolled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsVirtualCardEnrolled MUST be set to true to make this field effective.
	IsVirtualCardEnrolled bool

	FFI_USE_IsLocal                         bool // for IsLocal.
	FFI_USE_IsCached                        bool // for IsCached.
	FFI_USE_IsMigratable                    bool // for IsMigratable.
	FFI_USE_IsVirtualCardEnrollmentEligible bool // for IsVirtualCardEnrollmentEligible.
	FFI_USE_IsVirtualCardEnrolled           bool // for IsVirtualCardEnrolled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AutofillMetadata with all fields set.
func (p AutofillMetadata) FromRef(ref js.Ref) AutofillMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AutofillMetadata in the application heap.
func (p AutofillMetadata) New() js.Ref {
	return bindings.AutofillMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AutofillMetadata) UpdateFrom(ref js.Ref) {
	bindings.AutofillMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AutofillMetadata) Update(ref js.Ref) {
	bindings.AutofillMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AutofillMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.SummaryLabel.Ref(),
		p.SummarySublabel.Ref(),
	)
	p.SummaryLabel = p.SummaryLabel.FromRef(js.Undefined)
	p.SummarySublabel = p.SummarySublabel.FromRef(js.Undefined)
}

type AddressEntry struct {
	// Guid is "AddressEntry.guid"
	//
	// Optional
	Guid js.String
	// Fields is "AddressEntry.fields"
	//
	// Optional
	Fields js.Array[AddressField]
	// LanguageCode is "AddressEntry.languageCode"
	//
	// Optional
	LanguageCode js.String
	// Metadata is "AddressEntry.metadata"
	//
	// Optional
	//
	// NOTE: Metadata.FFI_USE MUST be set to true to get Metadata used.
	Metadata AutofillMetadata

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressEntry with all fields set.
func (p AddressEntry) FromRef(ref js.Ref) AddressEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddressEntry in the application heap.
func (p AddressEntry) New() js.Ref {
	return bindings.AddressEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressEntry) UpdateFrom(ref js.Ref) {
	bindings.AddressEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressEntry) Update(ref js.Ref) {
	bindings.AddressEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Guid.Ref(),
		p.Fields.Ref(),
		p.LanguageCode.Ref(),
	)
	p.Guid = p.Guid.FromRef(js.Undefined)
	p.Fields = p.Fields.FromRef(js.Undefined)
	p.LanguageCode = p.LanguageCode.FromRef(js.Undefined)
	if recursive {
		p.Metadata.FreeMembers(true)
	}
}

type CheckForDeviceAuthCallbackFunc func(this js.Ref, isDeviceAuthAvailable bool) js.Ref

func (fn CheckForDeviceAuthCallbackFunc) Register() js.Func[func(isDeviceAuthAvailable bool)] {
	return js.RegisterCallback[func(isDeviceAuthAvailable bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CheckForDeviceAuthCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CheckForDeviceAuthCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isDeviceAuthAvailable bool) js.Ref
	Arg T
}

func (cb *CheckForDeviceAuthCallback[T]) Register() js.Func[func(isDeviceAuthAvailable bool)] {
	return js.RegisterCallback[func(isDeviceAuthAvailable bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CheckForDeviceAuthCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CheckForUserAuthCallbackFunc func(this js.Ref, isUserAuthSuccessful bool) js.Ref

func (fn CheckForUserAuthCallbackFunc) Register() js.Func[func(isUserAuthSuccessful bool)] {
	return js.RegisterCallback[func(isUserAuthSuccessful bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CheckForUserAuthCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CheckForUserAuthCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isUserAuthSuccessful bool) js.Ref
	Arg T
}

func (cb *CheckForUserAuthCallback[T]) Register() js.Func[func(isUserAuthSuccessful bool)] {
	return js.RegisterCallback[func(isUserAuthSuccessful bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CheckForUserAuthCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CountryEntry struct {
	// Name is "CountryEntry.name"
	//
	// Optional
	Name js.String
	// CountryCode is "CountryEntry.countryCode"
	//
	// Optional
	CountryCode js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CountryEntry with all fields set.
func (p CountryEntry) FromRef(ref js.Ref) CountryEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CountryEntry in the application heap.
func (p CountryEntry) New() js.Ref {
	return bindings.CountryEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CountryEntry) UpdateFrom(ref js.Ref) {
	bindings.CountryEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CountryEntry) Update(ref js.Ref) {
	bindings.CountryEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CountryEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.CountryCode.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.CountryCode = p.CountryCode.FromRef(js.Undefined)
}

type CreditCardEntry struct {
	// Guid is "CreditCardEntry.guid"
	//
	// Optional
	Guid js.String
	// Name is "CreditCardEntry.name"
	//
	// Optional
	Name js.String
	// CardNumber is "CreditCardEntry.cardNumber"
	//
	// Optional
	CardNumber js.String
	// ExpirationMonth is "CreditCardEntry.expirationMonth"
	//
	// Optional
	ExpirationMonth js.String
	// ExpirationYear is "CreditCardEntry.expirationYear"
	//
	// Optional
	ExpirationYear js.String
	// Nickname is "CreditCardEntry.nickname"
	//
	// Optional
	Nickname js.String
	// Network is "CreditCardEntry.network"
	//
	// Optional
	Network js.String
	// ImageSrc is "CreditCardEntry.imageSrc"
	//
	// Optional
	ImageSrc js.String
	// Metadata is "CreditCardEntry.metadata"
	//
	// Optional
	//
	// NOTE: Metadata.FFI_USE MUST be set to true to get Metadata used.
	Metadata AutofillMetadata

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreditCardEntry with all fields set.
func (p CreditCardEntry) FromRef(ref js.Ref) CreditCardEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreditCardEntry in the application heap.
func (p CreditCardEntry) New() js.Ref {
	return bindings.CreditCardEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreditCardEntry) UpdateFrom(ref js.Ref) {
	bindings.CreditCardEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreditCardEntry) Update(ref js.Ref) {
	bindings.CreditCardEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreditCardEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Guid.Ref(),
		p.Name.Ref(),
		p.CardNumber.Ref(),
		p.ExpirationMonth.Ref(),
		p.ExpirationYear.Ref(),
		p.Nickname.Ref(),
		p.Network.Ref(),
		p.ImageSrc.Ref(),
	)
	p.Guid = p.Guid.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.CardNumber = p.CardNumber.FromRef(js.Undefined)
	p.ExpirationMonth = p.ExpirationMonth.FromRef(js.Undefined)
	p.ExpirationYear = p.ExpirationYear.FromRef(js.Undefined)
	p.Nickname = p.Nickname.FromRef(js.Undefined)
	p.Network = p.Network.FromRef(js.Undefined)
	p.ImageSrc = p.ImageSrc.FromRef(js.Undefined)
	if recursive {
		p.Metadata.FreeMembers(true)
	}
}

type GetAccountInfoCallbackFunc func(this js.Ref, accountInfo *AccountInfo) js.Ref

func (fn GetAccountInfoCallbackFunc) Register() js.Func[func(accountInfo *AccountInfo)] {
	return js.RegisterCallback[func(accountInfo *AccountInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAccountInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AccountInfo
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

type GetAccountInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, accountInfo *AccountInfo) js.Ref
	Arg T
}

func (cb *GetAccountInfoCallback[T]) Register() js.Func[func(accountInfo *AccountInfo)] {
	return js.RegisterCallback[func(accountInfo *AccountInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAccountInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AccountInfo
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

type GetAddressComponentsCallbackFunc func(this js.Ref, components *AddressComponents) js.Ref

func (fn GetAddressComponentsCallbackFunc) Register() js.Func[func(components *AddressComponents)] {
	return js.RegisterCallback[func(components *AddressComponents)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAddressComponentsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AddressComponents
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

type GetAddressComponentsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, components *AddressComponents) js.Ref
	Arg T
}

func (cb *GetAddressComponentsCallback[T]) Register() js.Func[func(components *AddressComponents)] {
	return js.RegisterCallback[func(components *AddressComponents)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAddressComponentsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AddressComponents
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

type GetAddressListCallbackFunc func(this js.Ref, entries js.Array[AddressEntry]) js.Ref

func (fn GetAddressListCallbackFunc) Register() js.Func[func(entries js.Array[AddressEntry])] {
	return js.RegisterCallback[func(entries js.Array[AddressEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAddressListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[AddressEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAddressListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[AddressEntry]) js.Ref
	Arg T
}

func (cb *GetAddressListCallback[T]) Register() js.Func[func(entries js.Array[AddressEntry])] {
	return js.RegisterCallback[func(entries js.Array[AddressEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAddressListCallback[T]) DispatchCallback(
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

		js.Array[AddressEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCountryListCallbackFunc func(this js.Ref, countries js.Array[CountryEntry]) js.Ref

func (fn GetCountryListCallbackFunc) Register() js.Func[func(countries js.Array[CountryEntry])] {
	return js.RegisterCallback[func(countries js.Array[CountryEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCountryListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[CountryEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCountryListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, countries js.Array[CountryEntry]) js.Ref
	Arg T
}

func (cb *GetCountryListCallback[T]) Register() js.Func[func(countries js.Array[CountryEntry])] {
	return js.RegisterCallback[func(countries js.Array[CountryEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCountryListCallback[T]) DispatchCallback(
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

		js.Array[CountryEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCreditCardListCallbackFunc func(this js.Ref, entries js.Array[CreditCardEntry]) js.Ref

func (fn GetCreditCardListCallbackFunc) Register() js.Func[func(entries js.Array[CreditCardEntry])] {
	return js.RegisterCallback[func(entries js.Array[CreditCardEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCreditCardListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[CreditCardEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCreditCardListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[CreditCardEntry]) js.Ref
	Arg T
}

func (cb *GetCreditCardListCallback[T]) Register() js.Func[func(entries js.Array[CreditCardEntry])] {
	return js.RegisterCallback[func(entries js.Array[CreditCardEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCreditCardListCallback[T]) DispatchCallback(
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

		js.Array[CreditCardEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetIbanListCallbackFunc func(this js.Ref, entries js.Array[IbanEntry]) js.Ref

func (fn GetIbanListCallbackFunc) Register() js.Func[func(entries js.Array[IbanEntry])] {
	return js.RegisterCallback[func(entries js.Array[IbanEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetIbanListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[IbanEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetIbanListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[IbanEntry]) js.Ref
	Arg T
}

func (cb *GetIbanListCallback[T]) Register() js.Func[func(entries js.Array[IbanEntry])] {
	return js.RegisterCallback[func(entries js.Array[IbanEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetIbanListCallback[T]) DispatchCallback(
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

		js.Array[IbanEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IbanEntry struct {
	// Guid is "IbanEntry.guid"
	//
	// Optional
	Guid js.String
	// Value is "IbanEntry.value"
	//
	// Optional
	Value js.String
	// Nickname is "IbanEntry.nickname"
	//
	// Optional
	Nickname js.String
	// Metadata is "IbanEntry.metadata"
	//
	// Optional
	//
	// NOTE: Metadata.FFI_USE MUST be set to true to get Metadata used.
	Metadata AutofillMetadata

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IbanEntry with all fields set.
func (p IbanEntry) FromRef(ref js.Ref) IbanEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IbanEntry in the application heap.
func (p IbanEntry) New() js.Ref {
	return bindings.IbanEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IbanEntry) UpdateFrom(ref js.Ref) {
	bindings.IbanEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IbanEntry) Update(ref js.Ref) {
	bindings.IbanEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IbanEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Guid.Ref(),
		p.Value.Ref(),
		p.Nickname.Ref(),
	)
	p.Guid = p.Guid.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
	p.Nickname = p.Nickname.FromRef(js.Undefined)
	if recursive {
		p.Metadata.FreeMembers(true)
	}
}

type IsValidIbanCallbackFunc func(this js.Ref, isValid bool) js.Ref

func (fn IsValidIbanCallbackFunc) Register() js.Func[func(isValid bool)] {
	return js.RegisterCallback[func(isValid bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsValidIbanCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsValidIbanCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isValid bool) js.Ref
	Arg T
}

func (cb *IsValidIbanCallback[T]) Register() js.Func[func(isValid bool)] {
	return js.RegisterCallback[func(isValid bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsValidIbanCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncAddVirtualCard returns true if the function "WEBEXT.autofillPrivate.addVirtualCard" exists.
func HasFuncAddVirtualCard() bool {
	return js.True == bindings.HasFuncAddVirtualCard()
}

// FuncAddVirtualCard returns the function "WEBEXT.autofillPrivate.addVirtualCard".
func FuncAddVirtualCard() (fn js.Func[func(cardId js.String)]) {
	bindings.FuncAddVirtualCard(
		js.Pointer(&fn),
	)
	return
}

// AddVirtualCard calls the function "WEBEXT.autofillPrivate.addVirtualCard" directly.
func AddVirtualCard(cardId js.String) (ret js.Void) {
	bindings.CallAddVirtualCard(
		js.Pointer(&ret),
		cardId.Ref(),
	)

	return
}

// TryAddVirtualCard calls the function "WEBEXT.autofillPrivate.addVirtualCard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddVirtualCard(cardId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddVirtualCard(
		js.Pointer(&ret), js.Pointer(&exception),
		cardId.Ref(),
	)

	return
}

// HasFuncAuthenticateUserAndFlipMandatoryAuthToggle returns true if the function "WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle" exists.
func HasFuncAuthenticateUserAndFlipMandatoryAuthToggle() bool {
	return js.True == bindings.HasFuncAuthenticateUserAndFlipMandatoryAuthToggle()
}

// FuncAuthenticateUserAndFlipMandatoryAuthToggle returns the function "WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle".
func FuncAuthenticateUserAndFlipMandatoryAuthToggle() (fn js.Func[func()]) {
	bindings.FuncAuthenticateUserAndFlipMandatoryAuthToggle(
		js.Pointer(&fn),
	)
	return
}

// AuthenticateUserAndFlipMandatoryAuthToggle calls the function "WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle" directly.
func AuthenticateUserAndFlipMandatoryAuthToggle() (ret js.Void) {
	bindings.CallAuthenticateUserAndFlipMandatoryAuthToggle(
		js.Pointer(&ret),
	)

	return
}

// TryAuthenticateUserAndFlipMandatoryAuthToggle calls the function "WEBEXT.autofillPrivate.authenticateUserAndFlipMandatoryAuthToggle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAuthenticateUserAndFlipMandatoryAuthToggle() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticateUserAndFlipMandatoryAuthToggle(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAuthenticateUserToEditLocalCard returns true if the function "WEBEXT.autofillPrivate.authenticateUserToEditLocalCard" exists.
func HasFuncAuthenticateUserToEditLocalCard() bool {
	return js.True == bindings.HasFuncAuthenticateUserToEditLocalCard()
}

// FuncAuthenticateUserToEditLocalCard returns the function "WEBEXT.autofillPrivate.authenticateUserToEditLocalCard".
func FuncAuthenticateUserToEditLocalCard() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncAuthenticateUserToEditLocalCard(
		js.Pointer(&fn),
	)
	return
}

// AuthenticateUserToEditLocalCard calls the function "WEBEXT.autofillPrivate.authenticateUserToEditLocalCard" directly.
func AuthenticateUserToEditLocalCard() (ret js.Promise[js.Boolean]) {
	bindings.CallAuthenticateUserToEditLocalCard(
		js.Pointer(&ret),
	)

	return
}

// TryAuthenticateUserToEditLocalCard calls the function "WEBEXT.autofillPrivate.authenticateUserToEditLocalCard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAuthenticateUserToEditLocalCard() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAuthenticateUserToEditLocalCard(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckIfDeviceAuthAvailable returns true if the function "WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable" exists.
func HasFuncCheckIfDeviceAuthAvailable() bool {
	return js.True == bindings.HasFuncCheckIfDeviceAuthAvailable()
}

// FuncCheckIfDeviceAuthAvailable returns the function "WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable".
func FuncCheckIfDeviceAuthAvailable() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncCheckIfDeviceAuthAvailable(
		js.Pointer(&fn),
	)
	return
}

// CheckIfDeviceAuthAvailable calls the function "WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable" directly.
func CheckIfDeviceAuthAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallCheckIfDeviceAuthAvailable(
		js.Pointer(&ret),
	)

	return
}

// TryCheckIfDeviceAuthAvailable calls the function "WEBEXT.autofillPrivate.checkIfDeviceAuthAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCheckIfDeviceAuthAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCheckIfDeviceAuthAvailable(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAccountInfo returns true if the function "WEBEXT.autofillPrivate.getAccountInfo" exists.
func HasFuncGetAccountInfo() bool {
	return js.True == bindings.HasFuncGetAccountInfo()
}

// FuncGetAccountInfo returns the function "WEBEXT.autofillPrivate.getAccountInfo".
func FuncGetAccountInfo() (fn js.Func[func() js.Promise[AccountInfo]]) {
	bindings.FuncGetAccountInfo(
		js.Pointer(&fn),
	)
	return
}

// GetAccountInfo calls the function "WEBEXT.autofillPrivate.getAccountInfo" directly.
func GetAccountInfo() (ret js.Promise[AccountInfo]) {
	bindings.CallGetAccountInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetAccountInfo calls the function "WEBEXT.autofillPrivate.getAccountInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAccountInfo() (ret js.Promise[AccountInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAccountInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAddressComponents returns true if the function "WEBEXT.autofillPrivate.getAddressComponents" exists.
func HasFuncGetAddressComponents() bool {
	return js.True == bindings.HasFuncGetAddressComponents()
}

// FuncGetAddressComponents returns the function "WEBEXT.autofillPrivate.getAddressComponents".
func FuncGetAddressComponents() (fn js.Func[func(countryCode js.String) js.Promise[AddressComponents]]) {
	bindings.FuncGetAddressComponents(
		js.Pointer(&fn),
	)
	return
}

// GetAddressComponents calls the function "WEBEXT.autofillPrivate.getAddressComponents" directly.
func GetAddressComponents(countryCode js.String) (ret js.Promise[AddressComponents]) {
	bindings.CallGetAddressComponents(
		js.Pointer(&ret),
		countryCode.Ref(),
	)

	return
}

// TryGetAddressComponents calls the function "WEBEXT.autofillPrivate.getAddressComponents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAddressComponents(countryCode js.String) (ret js.Promise[AddressComponents], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAddressComponents(
		js.Pointer(&ret), js.Pointer(&exception),
		countryCode.Ref(),
	)

	return
}

// HasFuncGetAddressList returns true if the function "WEBEXT.autofillPrivate.getAddressList" exists.
func HasFuncGetAddressList() bool {
	return js.True == bindings.HasFuncGetAddressList()
}

// FuncGetAddressList returns the function "WEBEXT.autofillPrivate.getAddressList".
func FuncGetAddressList() (fn js.Func[func() js.Promise[js.Array[AddressEntry]]]) {
	bindings.FuncGetAddressList(
		js.Pointer(&fn),
	)
	return
}

// GetAddressList calls the function "WEBEXT.autofillPrivate.getAddressList" directly.
func GetAddressList() (ret js.Promise[js.Array[AddressEntry]]) {
	bindings.CallGetAddressList(
		js.Pointer(&ret),
	)

	return
}

// TryGetAddressList calls the function "WEBEXT.autofillPrivate.getAddressList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAddressList() (ret js.Promise[js.Array[AddressEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAddressList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCountryList returns true if the function "WEBEXT.autofillPrivate.getCountryList" exists.
func HasFuncGetCountryList() bool {
	return js.True == bindings.HasFuncGetCountryList()
}

// FuncGetCountryList returns the function "WEBEXT.autofillPrivate.getCountryList".
func FuncGetCountryList() (fn js.Func[func() js.Promise[js.Array[CountryEntry]]]) {
	bindings.FuncGetCountryList(
		js.Pointer(&fn),
	)
	return
}

// GetCountryList calls the function "WEBEXT.autofillPrivate.getCountryList" directly.
func GetCountryList() (ret js.Promise[js.Array[CountryEntry]]) {
	bindings.CallGetCountryList(
		js.Pointer(&ret),
	)

	return
}

// TryGetCountryList calls the function "WEBEXT.autofillPrivate.getCountryList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCountryList() (ret js.Promise[js.Array[CountryEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCountryList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCreditCardList returns true if the function "WEBEXT.autofillPrivate.getCreditCardList" exists.
func HasFuncGetCreditCardList() bool {
	return js.True == bindings.HasFuncGetCreditCardList()
}

// FuncGetCreditCardList returns the function "WEBEXT.autofillPrivate.getCreditCardList".
func FuncGetCreditCardList() (fn js.Func[func() js.Promise[js.Array[CreditCardEntry]]]) {
	bindings.FuncGetCreditCardList(
		js.Pointer(&fn),
	)
	return
}

// GetCreditCardList calls the function "WEBEXT.autofillPrivate.getCreditCardList" directly.
func GetCreditCardList() (ret js.Promise[js.Array[CreditCardEntry]]) {
	bindings.CallGetCreditCardList(
		js.Pointer(&ret),
	)

	return
}

// TryGetCreditCardList calls the function "WEBEXT.autofillPrivate.getCreditCardList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCreditCardList() (ret js.Promise[js.Array[CreditCardEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCreditCardList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetIbanList returns true if the function "WEBEXT.autofillPrivate.getIbanList" exists.
func HasFuncGetIbanList() bool {
	return js.True == bindings.HasFuncGetIbanList()
}

// FuncGetIbanList returns the function "WEBEXT.autofillPrivate.getIbanList".
func FuncGetIbanList() (fn js.Func[func() js.Promise[js.Array[IbanEntry]]]) {
	bindings.FuncGetIbanList(
		js.Pointer(&fn),
	)
	return
}

// GetIbanList calls the function "WEBEXT.autofillPrivate.getIbanList" directly.
func GetIbanList() (ret js.Promise[js.Array[IbanEntry]]) {
	bindings.CallGetIbanList(
		js.Pointer(&ret),
	)

	return
}

// TryGetIbanList calls the function "WEBEXT.autofillPrivate.getIbanList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetIbanList() (ret js.Promise[js.Array[IbanEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetIbanList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsValidIban returns true if the function "WEBEXT.autofillPrivate.isValidIban" exists.
func HasFuncIsValidIban() bool {
	return js.True == bindings.HasFuncIsValidIban()
}

// FuncIsValidIban returns the function "WEBEXT.autofillPrivate.isValidIban".
func FuncIsValidIban() (fn js.Func[func(ibanValue js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsValidIban(
		js.Pointer(&fn),
	)
	return
}

// IsValidIban calls the function "WEBEXT.autofillPrivate.isValidIban" directly.
func IsValidIban(ibanValue js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsValidIban(
		js.Pointer(&ret),
		ibanValue.Ref(),
	)

	return
}

// TryIsValidIban calls the function "WEBEXT.autofillPrivate.isValidIban"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsValidIban(ibanValue js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsValidIban(
		js.Pointer(&ret), js.Pointer(&exception),
		ibanValue.Ref(),
	)

	return
}

// HasFuncLogServerCardLinkClicked returns true if the function "WEBEXT.autofillPrivate.logServerCardLinkClicked" exists.
func HasFuncLogServerCardLinkClicked() bool {
	return js.True == bindings.HasFuncLogServerCardLinkClicked()
}

// FuncLogServerCardLinkClicked returns the function "WEBEXT.autofillPrivate.logServerCardLinkClicked".
func FuncLogServerCardLinkClicked() (fn js.Func[func()]) {
	bindings.FuncLogServerCardLinkClicked(
		js.Pointer(&fn),
	)
	return
}

// LogServerCardLinkClicked calls the function "WEBEXT.autofillPrivate.logServerCardLinkClicked" directly.
func LogServerCardLinkClicked() (ret js.Void) {
	bindings.CallLogServerCardLinkClicked(
		js.Pointer(&ret),
	)

	return
}

// TryLogServerCardLinkClicked calls the function "WEBEXT.autofillPrivate.logServerCardLinkClicked"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLogServerCardLinkClicked() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLogServerCardLinkClicked(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMaskCreditCard returns true if the function "WEBEXT.autofillPrivate.maskCreditCard" exists.
func HasFuncMaskCreditCard() bool {
	return js.True == bindings.HasFuncMaskCreditCard()
}

// FuncMaskCreditCard returns the function "WEBEXT.autofillPrivate.maskCreditCard".
func FuncMaskCreditCard() (fn js.Func[func(guid js.String)]) {
	bindings.FuncMaskCreditCard(
		js.Pointer(&fn),
	)
	return
}

// MaskCreditCard calls the function "WEBEXT.autofillPrivate.maskCreditCard" directly.
func MaskCreditCard(guid js.String) (ret js.Void) {
	bindings.CallMaskCreditCard(
		js.Pointer(&ret),
		guid.Ref(),
	)

	return
}

// TryMaskCreditCard calls the function "WEBEXT.autofillPrivate.maskCreditCard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMaskCreditCard(guid js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMaskCreditCard(
		js.Pointer(&ret), js.Pointer(&exception),
		guid.Ref(),
	)

	return
}

// HasFuncMigrateCreditCards returns true if the function "WEBEXT.autofillPrivate.migrateCreditCards" exists.
func HasFuncMigrateCreditCards() bool {
	return js.True == bindings.HasFuncMigrateCreditCards()
}

// FuncMigrateCreditCards returns the function "WEBEXT.autofillPrivate.migrateCreditCards".
func FuncMigrateCreditCards() (fn js.Func[func()]) {
	bindings.FuncMigrateCreditCards(
		js.Pointer(&fn),
	)
	return
}

// MigrateCreditCards calls the function "WEBEXT.autofillPrivate.migrateCreditCards" directly.
func MigrateCreditCards() (ret js.Void) {
	bindings.CallMigrateCreditCards(
		js.Pointer(&ret),
	)

	return
}

// TryMigrateCreditCards calls the function "WEBEXT.autofillPrivate.migrateCreditCards"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMigrateCreditCards() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMigrateCreditCards(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnPersonalDataChangedEventCallbackFunc func(this js.Ref, addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo) js.Ref

func (fn OnPersonalDataChangedEventCallbackFunc) Register() js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)] {
	return js.RegisterCallback[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPersonalDataChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg3 AccountInfo
	arg3.UpdateFrom(args[3+1])
	defer arg3.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Array[AddressEntry]{}.FromRef(args[0+1]),
		js.Array[CreditCardEntry]{}.FromRef(args[1+1]),
		js.Array[IbanEntry]{}.FromRef(args[2+1]),
		mark.NoEscape(&arg3),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPersonalDataChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo) js.Ref
	Arg T
}

func (cb *OnPersonalDataChangedEventCallback[T]) Register() js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)] {
	return js.RegisterCallback[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPersonalDataChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg3 AccountInfo
	arg3.UpdateFrom(args[3+1])
	defer arg3.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[AddressEntry]{}.FromRef(args[0+1]),
		js.Array[CreditCardEntry]{}.FromRef(args[1+1]),
		js.Array[IbanEntry]{}.FromRef(args[2+1]),
		mark.NoEscape(&arg3),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPersonalDataChanged returns true if the function "WEBEXT.autofillPrivate.onPersonalDataChanged.addListener" exists.
func HasFuncOnPersonalDataChanged() bool {
	return js.True == bindings.HasFuncOnPersonalDataChanged()
}

// FuncOnPersonalDataChanged returns the function "WEBEXT.autofillPrivate.onPersonalDataChanged.addListener".
func FuncOnPersonalDataChanged() (fn js.Func[func(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)])]) {
	bindings.FuncOnPersonalDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.addListener" directly.
func OnPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret js.Void) {
	bindings.CallOnPersonalDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPersonalDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPersonalDataChanged returns true if the function "WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener" exists.
func HasFuncOffPersonalDataChanged() bool {
	return js.True == bindings.HasFuncOffPersonalDataChanged()
}

// FuncOffPersonalDataChanged returns the function "WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener".
func FuncOffPersonalDataChanged() (fn js.Func[func(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)])]) {
	bindings.FuncOffPersonalDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener" directly.
func OffPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret js.Void) {
	bindings.CallOffPersonalDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPersonalDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPersonalDataChanged returns true if the function "WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener" exists.
func HasFuncHasOnPersonalDataChanged() bool {
	return js.True == bindings.HasFuncHasOnPersonalDataChanged()
}

// FuncHasOnPersonalDataChanged returns the function "WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener".
func FuncHasOnPersonalDataChanged() (fn js.Func[func(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) bool]) {
	bindings.FuncHasOnPersonalDataChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener" directly.
func HasOnPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret bool) {
	bindings.CallHasOnPersonalDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPersonalDataChanged calls the function "WEBEXT.autofillPrivate.onPersonalDataChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPersonalDataChanged(callback js.Func[func(addressEntries js.Array[AddressEntry], creditCardEntries js.Array[CreditCardEntry], ibans js.Array[IbanEntry], accountInfo *AccountInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPersonalDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveEntry returns true if the function "WEBEXT.autofillPrivate.removeEntry" exists.
func HasFuncRemoveEntry() bool {
	return js.True == bindings.HasFuncRemoveEntry()
}

// FuncRemoveEntry returns the function "WEBEXT.autofillPrivate.removeEntry".
func FuncRemoveEntry() (fn js.Func[func(guid js.String)]) {
	bindings.FuncRemoveEntry(
		js.Pointer(&fn),
	)
	return
}

// RemoveEntry calls the function "WEBEXT.autofillPrivate.removeEntry" directly.
func RemoveEntry(guid js.String) (ret js.Void) {
	bindings.CallRemoveEntry(
		js.Pointer(&ret),
		guid.Ref(),
	)

	return
}

// TryRemoveEntry calls the function "WEBEXT.autofillPrivate.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveEntry(guid js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveEntry(
		js.Pointer(&ret), js.Pointer(&exception),
		guid.Ref(),
	)

	return
}

// HasFuncRemoveVirtualCard returns true if the function "WEBEXT.autofillPrivate.removeVirtualCard" exists.
func HasFuncRemoveVirtualCard() bool {
	return js.True == bindings.HasFuncRemoveVirtualCard()
}

// FuncRemoveVirtualCard returns the function "WEBEXT.autofillPrivate.removeVirtualCard".
func FuncRemoveVirtualCard() (fn js.Func[func(cardId js.String)]) {
	bindings.FuncRemoveVirtualCard(
		js.Pointer(&fn),
	)
	return
}

// RemoveVirtualCard calls the function "WEBEXT.autofillPrivate.removeVirtualCard" directly.
func RemoveVirtualCard(cardId js.String) (ret js.Void) {
	bindings.CallRemoveVirtualCard(
		js.Pointer(&ret),
		cardId.Ref(),
	)

	return
}

// TryRemoveVirtualCard calls the function "WEBEXT.autofillPrivate.removeVirtualCard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveVirtualCard(cardId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveVirtualCard(
		js.Pointer(&ret), js.Pointer(&exception),
		cardId.Ref(),
	)

	return
}

// HasFuncSaveAddress returns true if the function "WEBEXT.autofillPrivate.saveAddress" exists.
func HasFuncSaveAddress() bool {
	return js.True == bindings.HasFuncSaveAddress()
}

// FuncSaveAddress returns the function "WEBEXT.autofillPrivate.saveAddress".
func FuncSaveAddress() (fn js.Func[func(address AddressEntry)]) {
	bindings.FuncSaveAddress(
		js.Pointer(&fn),
	)
	return
}

// SaveAddress calls the function "WEBEXT.autofillPrivate.saveAddress" directly.
func SaveAddress(address AddressEntry) (ret js.Void) {
	bindings.CallSaveAddress(
		js.Pointer(&ret),
		js.Pointer(&address),
	)

	return
}

// TrySaveAddress calls the function "WEBEXT.autofillPrivate.saveAddress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySaveAddress(address AddressEntry) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySaveAddress(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&address),
	)

	return
}

// HasFuncSaveCreditCard returns true if the function "WEBEXT.autofillPrivate.saveCreditCard" exists.
func HasFuncSaveCreditCard() bool {
	return js.True == bindings.HasFuncSaveCreditCard()
}

// FuncSaveCreditCard returns the function "WEBEXT.autofillPrivate.saveCreditCard".
func FuncSaveCreditCard() (fn js.Func[func(card CreditCardEntry)]) {
	bindings.FuncSaveCreditCard(
		js.Pointer(&fn),
	)
	return
}

// SaveCreditCard calls the function "WEBEXT.autofillPrivate.saveCreditCard" directly.
func SaveCreditCard(card CreditCardEntry) (ret js.Void) {
	bindings.CallSaveCreditCard(
		js.Pointer(&ret),
		js.Pointer(&card),
	)

	return
}

// TrySaveCreditCard calls the function "WEBEXT.autofillPrivate.saveCreditCard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySaveCreditCard(card CreditCardEntry) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySaveCreditCard(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&card),
	)

	return
}

// HasFuncSaveIban returns true if the function "WEBEXT.autofillPrivate.saveIban" exists.
func HasFuncSaveIban() bool {
	return js.True == bindings.HasFuncSaveIban()
}

// FuncSaveIban returns the function "WEBEXT.autofillPrivate.saveIban".
func FuncSaveIban() (fn js.Func[func(iban IbanEntry)]) {
	bindings.FuncSaveIban(
		js.Pointer(&fn),
	)
	return
}

// SaveIban calls the function "WEBEXT.autofillPrivate.saveIban" directly.
func SaveIban(iban IbanEntry) (ret js.Void) {
	bindings.CallSaveIban(
		js.Pointer(&ret),
		js.Pointer(&iban),
	)

	return
}

// TrySaveIban calls the function "WEBEXT.autofillPrivate.saveIban"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySaveIban(iban IbanEntry) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySaveIban(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&iban),
	)

	return
}

// HasFuncSetCreditCardFIDOAuthEnabledState returns true if the function "WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState" exists.
func HasFuncSetCreditCardFIDOAuthEnabledState() bool {
	return js.True == bindings.HasFuncSetCreditCardFIDOAuthEnabledState()
}

// FuncSetCreditCardFIDOAuthEnabledState returns the function "WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState".
func FuncSetCreditCardFIDOAuthEnabledState() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetCreditCardFIDOAuthEnabledState(
		js.Pointer(&fn),
	)
	return
}

// SetCreditCardFIDOAuthEnabledState calls the function "WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState" directly.
func SetCreditCardFIDOAuthEnabledState(enabled bool) (ret js.Void) {
	bindings.CallSetCreditCardFIDOAuthEnabledState(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetCreditCardFIDOAuthEnabledState calls the function "WEBEXT.autofillPrivate.setCreditCardFIDOAuthEnabledState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCreditCardFIDOAuthEnabledState(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCreditCardFIDOAuthEnabledState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

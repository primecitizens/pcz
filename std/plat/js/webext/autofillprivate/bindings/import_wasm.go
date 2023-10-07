// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/autofillprivate store_AccountInfo
//go:noescape
func AccountInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AccountInfo
//go:noescape
func AccountInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate constof_ServerFieldType
//go:noescape
func ConstOfServerFieldType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autofillprivate store_AddressComponent
//go:noescape
func AddressComponentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AddressComponent
//go:noescape
func AddressComponentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_AddressComponentRow
//go:noescape
func AddressComponentRowJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AddressComponentRow
//go:noescape
func AddressComponentRowJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_AddressComponents
//go:noescape
func AddressComponentsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AddressComponents
//go:noescape
func AddressComponentsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_AddressField
//go:noescape
func AddressFieldJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AddressField
//go:noescape
func AddressFieldJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate constof_AddressSource
//go:noescape
func ConstOfAddressSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/autofillprivate store_AutofillMetadata
//go:noescape
func AutofillMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AutofillMetadata
//go:noescape
func AutofillMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_AddressEntry
//go:noescape
func AddressEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_AddressEntry
//go:noescape
func AddressEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_CountryEntry
//go:noescape
func CountryEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_CountryEntry
//go:noescape
func CountryEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_CreditCardEntry
//go:noescape
func CreditCardEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_CreditCardEntry
//go:noescape
func CreditCardEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate store_IbanEntry
//go:noescape
func IbanEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autofillprivate load_IbanEntry
//go:noescape
func IbanEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autofillprivate has_AddVirtualCard
//go:noescape
func HasFuncAddVirtualCard() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_AddVirtualCard
//go:noescape
func FuncAddVirtualCard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_AddVirtualCard
//go:noescape
func CallAddVirtualCard(
	retPtr unsafe.Pointer,
	cardId js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_AddVirtualCard
//go:noescape
func TryAddVirtualCard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cardId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_AuthenticateUserAndFlipMandatoryAuthToggle
//go:noescape
func HasFuncAuthenticateUserAndFlipMandatoryAuthToggle() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_AuthenticateUserAndFlipMandatoryAuthToggle
//go:noescape
func FuncAuthenticateUserAndFlipMandatoryAuthToggle(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_AuthenticateUserAndFlipMandatoryAuthToggle
//go:noescape
func CallAuthenticateUserAndFlipMandatoryAuthToggle(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_AuthenticateUserAndFlipMandatoryAuthToggle
//go:noescape
func TryAuthenticateUserAndFlipMandatoryAuthToggle(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_AuthenticateUserToEditLocalCard
//go:noescape
func HasFuncAuthenticateUserToEditLocalCard() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_AuthenticateUserToEditLocalCard
//go:noescape
func FuncAuthenticateUserToEditLocalCard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_AuthenticateUserToEditLocalCard
//go:noescape
func CallAuthenticateUserToEditLocalCard(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_AuthenticateUserToEditLocalCard
//go:noescape
func TryAuthenticateUserToEditLocalCard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_CheckIfDeviceAuthAvailable
//go:noescape
func HasFuncCheckIfDeviceAuthAvailable() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_CheckIfDeviceAuthAvailable
//go:noescape
func FuncCheckIfDeviceAuthAvailable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_CheckIfDeviceAuthAvailable
//go:noescape
func CallCheckIfDeviceAuthAvailable(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_CheckIfDeviceAuthAvailable
//go:noescape
func TryCheckIfDeviceAuthAvailable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetAccountInfo
//go:noescape
func HasFuncGetAccountInfo() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetAccountInfo
//go:noescape
func FuncGetAccountInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetAccountInfo
//go:noescape
func CallGetAccountInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_GetAccountInfo
//go:noescape
func TryGetAccountInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetAddressComponents
//go:noescape
func HasFuncGetAddressComponents() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetAddressComponents
//go:noescape
func FuncGetAddressComponents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetAddressComponents
//go:noescape
func CallGetAddressComponents(
	retPtr unsafe.Pointer,
	countryCode js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_GetAddressComponents
//go:noescape
func TryGetAddressComponents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	countryCode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetAddressList
//go:noescape
func HasFuncGetAddressList() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetAddressList
//go:noescape
func FuncGetAddressList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetAddressList
//go:noescape
func CallGetAddressList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_GetAddressList
//go:noescape
func TryGetAddressList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetCountryList
//go:noescape
func HasFuncGetCountryList() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetCountryList
//go:noescape
func FuncGetCountryList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetCountryList
//go:noescape
func CallGetCountryList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_GetCountryList
//go:noescape
func TryGetCountryList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetCreditCardList
//go:noescape
func HasFuncGetCreditCardList() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetCreditCardList
//go:noescape
func FuncGetCreditCardList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetCreditCardList
//go:noescape
func CallGetCreditCardList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_GetCreditCardList
//go:noescape
func TryGetCreditCardList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_GetIbanList
//go:noescape
func HasFuncGetIbanList() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_GetIbanList
//go:noescape
func FuncGetIbanList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_GetIbanList
//go:noescape
func CallGetIbanList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_GetIbanList
//go:noescape
func TryGetIbanList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_IsValidIban
//go:noescape
func HasFuncIsValidIban() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_IsValidIban
//go:noescape
func FuncIsValidIban(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_IsValidIban
//go:noescape
func CallIsValidIban(
	retPtr unsafe.Pointer,
	ibanValue js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_IsValidIban
//go:noescape
func TryIsValidIban(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ibanValue js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_LogServerCardLinkClicked
//go:noescape
func HasFuncLogServerCardLinkClicked() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_LogServerCardLinkClicked
//go:noescape
func FuncLogServerCardLinkClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_LogServerCardLinkClicked
//go:noescape
func CallLogServerCardLinkClicked(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_LogServerCardLinkClicked
//go:noescape
func TryLogServerCardLinkClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_MaskCreditCard
//go:noescape
func HasFuncMaskCreditCard() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_MaskCreditCard
//go:noescape
func FuncMaskCreditCard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_MaskCreditCard
//go:noescape
func CallMaskCreditCard(
	retPtr unsafe.Pointer,
	guid js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_MaskCreditCard
//go:noescape
func TryMaskCreditCard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	guid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_MigrateCreditCards
//go:noescape
func HasFuncMigrateCreditCards() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_MigrateCreditCards
//go:noescape
func FuncMigrateCreditCards(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_MigrateCreditCards
//go:noescape
func CallMigrateCreditCards(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_MigrateCreditCards
//go:noescape
func TryMigrateCreditCards(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_OnPersonalDataChanged
//go:noescape
func HasFuncOnPersonalDataChanged() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_OnPersonalDataChanged
//go:noescape
func FuncOnPersonalDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_OnPersonalDataChanged
//go:noescape
func CallOnPersonalDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_OnPersonalDataChanged
//go:noescape
func TryOnPersonalDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_OffPersonalDataChanged
//go:noescape
func HasFuncOffPersonalDataChanged() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_OffPersonalDataChanged
//go:noescape
func FuncOffPersonalDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_OffPersonalDataChanged
//go:noescape
func CallOffPersonalDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_OffPersonalDataChanged
//go:noescape
func TryOffPersonalDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_HasOnPersonalDataChanged
//go:noescape
func HasFuncHasOnPersonalDataChanged() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_HasOnPersonalDataChanged
//go:noescape
func FuncHasOnPersonalDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_HasOnPersonalDataChanged
//go:noescape
func CallHasOnPersonalDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_HasOnPersonalDataChanged
//go:noescape
func TryHasOnPersonalDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_RemoveEntry
//go:noescape
func HasFuncRemoveEntry() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_RemoveEntry
//go:noescape
func FuncRemoveEntry(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_RemoveEntry
//go:noescape
func CallRemoveEntry(
	retPtr unsafe.Pointer,
	guid js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_RemoveEntry
//go:noescape
func TryRemoveEntry(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	guid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_RemoveVirtualCard
//go:noescape
func HasFuncRemoveVirtualCard() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_RemoveVirtualCard
//go:noescape
func FuncRemoveVirtualCard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_RemoveVirtualCard
//go:noescape
func CallRemoveVirtualCard(
	retPtr unsafe.Pointer,
	cardId js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_RemoveVirtualCard
//go:noescape
func TryRemoveVirtualCard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cardId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_SaveAddress
//go:noescape
func HasFuncSaveAddress() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_SaveAddress
//go:noescape
func FuncSaveAddress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_SaveAddress
//go:noescape
func CallSaveAddress(
	retPtr unsafe.Pointer,
	address unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_SaveAddress
//go:noescape
func TrySaveAddress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	address unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_SaveCreditCard
//go:noescape
func HasFuncSaveCreditCard() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_SaveCreditCard
//go:noescape
func FuncSaveCreditCard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_SaveCreditCard
//go:noescape
func CallSaveCreditCard(
	retPtr unsafe.Pointer,
	card unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_SaveCreditCard
//go:noescape
func TrySaveCreditCard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	card unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_SaveIban
//go:noescape
func HasFuncSaveIban() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_SaveIban
//go:noescape
func FuncSaveIban(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_SaveIban
//go:noescape
func CallSaveIban(
	retPtr unsafe.Pointer,
	iban unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate try_SaveIban
//go:noescape
func TrySaveIban(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	iban unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autofillprivate has_SetCreditCardFIDOAuthEnabledState
//go:noescape
func HasFuncSetCreditCardFIDOAuthEnabledState() js.Ref

//go:wasmimport plat/js/webext/autofillprivate func_SetCreditCardFIDOAuthEnabledState
//go:noescape
func FuncSetCreditCardFIDOAuthEnabledState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autofillprivate call_SetCreditCardFIDOAuthEnabledState
//go:noescape
func CallSetCreditCardFIDOAuthEnabledState(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autofillprivate try_SetCreditCardFIDOAuthEnabledState
//go:noescape
func TrySetCreditCardFIDOAuthEnabledState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

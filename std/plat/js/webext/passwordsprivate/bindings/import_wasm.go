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

//go:wasmimport plat/js/webext/passwordsprivate store_AddPasswordOptions
//go:noescape
func AddPasswordOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_AddPasswordOptions
//go:noescape
func AddPasswordOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_CompromiseType
//go:noescape
func ConstOfCompromiseType(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_CompromisedInfo
//go:noescape
func CompromisedInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_CompromisedInfo
//go:noescape
func CompromisedInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_DomainInfo
//go:noescape
func DomainInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_DomainInfo
//go:noescape
func DomainInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_PasswordStoreSet
//go:noescape
func ConstOfPasswordStoreSet(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_PasswordUiEntry
//go:noescape
func PasswordUiEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_PasswordUiEntry
//go:noescape
func PasswordUiEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_CredentialGroup
//go:noescape
func CredentialGroupJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_CredentialGroup
//go:noescape
func CredentialGroupJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_PasswordUiEntryList
//go:noescape
func PasswordUiEntryListJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_PasswordUiEntryList
//go:noescape
func PasswordUiEntryListJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_UrlCollection
//go:noescape
func UrlCollectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_UrlCollection
//go:noescape
func UrlCollectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_ExceptionEntry
//go:noescape
func ExceptionEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_ExceptionEntry
//go:noescape
func ExceptionEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_ExportProgressStatus
//go:noescape
func ConstOfExportProgressStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate constof_FamilyFetchStatus
//go:noescape
func ConstOfFamilyFetchStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_PublicKey
//go:noescape
func PublicKeyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_PublicKey
//go:noescape
func PublicKeyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_RecipientInfo
//go:noescape
func RecipientInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_RecipientInfo
//go:noescape
func RecipientInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_FamilyFetchResults
//go:noescape
func FamilyFetchResultsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_FamilyFetchResults
//go:noescape
func FamilyFetchResultsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_ImportEntryStatus
//go:noescape
func ConstOfImportEntryStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_ImportEntry
//go:noescape
func ImportEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_ImportEntry
//go:noescape
func ImportEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_ImportResultsStatus
//go:noescape
func ConstOfImportResultsStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_ImportResults
//go:noescape
func ImportResultsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_ImportResults
//go:noescape
func ImportResultsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_PasswordCheckState
//go:noescape
func ConstOfPasswordCheckState(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate store_PasswordCheckStatus
//go:noescape
func PasswordCheckStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_PasswordCheckStatus
//go:noescape
func PasswordCheckStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate store_PasswordExportProgress
//go:noescape
func PasswordExportProgressJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate load_PasswordExportProgress
//go:noescape
func PasswordExportProgressJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/passwordsprivate constof_PlaintextReason
//go:noescape
func ConstOfPlaintextReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/passwordsprivate has_AddPassword
//go:noescape
func HasFuncAddPassword() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_AddPassword
//go:noescape
func FuncAddPassword(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_AddPassword
//go:noescape
func CallAddPassword(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_AddPassword
//go:noescape
func TryAddPassword(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ChangeCredential
//go:noescape
func HasFuncChangeCredential() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ChangeCredential
//go:noescape
func FuncChangeCredential(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ChangeCredential
//go:noescape
func CallChangeCredential(
	retPtr unsafe.Pointer,
	credential unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_ChangeCredential
//go:noescape
func TryChangeCredential(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	credential unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ContinueImport
//go:noescape
func HasFuncContinueImport() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ContinueImport
//go:noescape
func FuncContinueImport(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ContinueImport
//go:noescape
func CallContinueImport(
	retPtr unsafe.Pointer,
	selectedIds js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_ContinueImport
//go:noescape
func TryContinueImport(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectedIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ExportPasswords
//go:noescape
func HasFuncExportPasswords() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ExportPasswords
//go:noescape
func FuncExportPasswords(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ExportPasswords
//go:noescape
func CallExportPasswords(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_ExportPasswords
//go:noescape
func TryExportPasswords(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ExtendAuthValidity
//go:noescape
func HasFuncExtendAuthValidity() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ExtendAuthValidity
//go:noescape
func FuncExtendAuthValidity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ExtendAuthValidity
//go:noescape
func CallExtendAuthValidity(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_ExtendAuthValidity
//go:noescape
func TryExtendAuthValidity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_FetchFamilyMembers
//go:noescape
func HasFuncFetchFamilyMembers() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_FetchFamilyMembers
//go:noescape
func FuncFetchFamilyMembers(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_FetchFamilyMembers
//go:noescape
func CallFetchFamilyMembers(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_FetchFamilyMembers
//go:noescape
func TryFetchFamilyMembers(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetCredentialGroups
//go:noescape
func HasFuncGetCredentialGroups() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetCredentialGroups
//go:noescape
func FuncGetCredentialGroups(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetCredentialGroups
//go:noescape
func CallGetCredentialGroups(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetCredentialGroups
//go:noescape
func TryGetCredentialGroups(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetCredentialsWithReusedPassword
//go:noescape
func HasFuncGetCredentialsWithReusedPassword() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetCredentialsWithReusedPassword
//go:noescape
func FuncGetCredentialsWithReusedPassword(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetCredentialsWithReusedPassword
//go:noescape
func CallGetCredentialsWithReusedPassword(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetCredentialsWithReusedPassword
//go:noescape
func TryGetCredentialsWithReusedPassword(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetInsecureCredentials
//go:noescape
func HasFuncGetInsecureCredentials() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetInsecureCredentials
//go:noescape
func FuncGetInsecureCredentials(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetInsecureCredentials
//go:noescape
func CallGetInsecureCredentials(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetInsecureCredentials
//go:noescape
func TryGetInsecureCredentials(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetPasswordCheckStatus
//go:noescape
func HasFuncGetPasswordCheckStatus() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetPasswordCheckStatus
//go:noescape
func FuncGetPasswordCheckStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetPasswordCheckStatus
//go:noescape
func CallGetPasswordCheckStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetPasswordCheckStatus
//go:noescape
func TryGetPasswordCheckStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetPasswordExceptionList
//go:noescape
func HasFuncGetPasswordExceptionList() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetPasswordExceptionList
//go:noescape
func FuncGetPasswordExceptionList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetPasswordExceptionList
//go:noescape
func CallGetPasswordExceptionList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetPasswordExceptionList
//go:noescape
func TryGetPasswordExceptionList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetSavedPasswordList
//go:noescape
func HasFuncGetSavedPasswordList() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetSavedPasswordList
//go:noescape
func FuncGetSavedPasswordList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetSavedPasswordList
//go:noescape
func CallGetSavedPasswordList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_GetSavedPasswordList
//go:noescape
func TryGetSavedPasswordList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_GetUrlCollection
//go:noescape
func HasFuncGetUrlCollection() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_GetUrlCollection
//go:noescape
func FuncGetUrlCollection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_GetUrlCollection
//go:noescape
func CallGetUrlCollection(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_GetUrlCollection
//go:noescape
func TryGetUrlCollection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ImportPasswords
//go:noescape
func HasFuncImportPasswords() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ImportPasswords
//go:noescape
func FuncImportPasswords(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ImportPasswords
//go:noescape
func CallImportPasswords(
	retPtr unsafe.Pointer,
	toStore uint32)

//go:wasmimport plat/js/webext/passwordsprivate try_ImportPasswords
//go:noescape
func TryImportPasswords(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	toStore uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_IsAccountStoreDefault
//go:noescape
func HasFuncIsAccountStoreDefault() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_IsAccountStoreDefault
//go:noescape
func FuncIsAccountStoreDefault(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_IsAccountStoreDefault
//go:noescape
func CallIsAccountStoreDefault(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_IsAccountStoreDefault
//go:noescape
func TryIsAccountStoreDefault(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_IsOptedInForAccountStorage
//go:noescape
func HasFuncIsOptedInForAccountStorage() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_IsOptedInForAccountStorage
//go:noescape
func FuncIsOptedInForAccountStorage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_IsOptedInForAccountStorage
//go:noescape
func CallIsOptedInForAccountStorage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_IsOptedInForAccountStorage
//go:noescape
func TryIsOptedInForAccountStorage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_MovePasswordsToAccount
//go:noescape
func HasFuncMovePasswordsToAccount() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_MovePasswordsToAccount
//go:noescape
func FuncMovePasswordsToAccount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_MovePasswordsToAccount
//go:noescape
func CallMovePasswordsToAccount(
	retPtr unsafe.Pointer,
	ids js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_MovePasswordsToAccount
//go:noescape
func TryMovePasswordsToAccount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ids js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_MuteInsecureCredential
//go:noescape
func HasFuncMuteInsecureCredential() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_MuteInsecureCredential
//go:noescape
func FuncMuteInsecureCredential(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_MuteInsecureCredential
//go:noescape
func CallMuteInsecureCredential(
	retPtr unsafe.Pointer,
	credential unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_MuteInsecureCredential
//go:noescape
func TryMuteInsecureCredential(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	credential unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnAccountStorageOptInStateChanged
//go:noescape
func HasFuncOnAccountStorageOptInStateChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnAccountStorageOptInStateChanged
//go:noescape
func FuncOnAccountStorageOptInStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnAccountStorageOptInStateChanged
//go:noescape
func CallOnAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnAccountStorageOptInStateChanged
//go:noescape
func TryOnAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffAccountStorageOptInStateChanged
//go:noescape
func HasFuncOffAccountStorageOptInStateChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffAccountStorageOptInStateChanged
//go:noescape
func FuncOffAccountStorageOptInStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffAccountStorageOptInStateChanged
//go:noescape
func CallOffAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffAccountStorageOptInStateChanged
//go:noescape
func TryOffAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnAccountStorageOptInStateChanged
//go:noescape
func HasFuncHasOnAccountStorageOptInStateChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnAccountStorageOptInStateChanged
//go:noescape
func FuncHasOnAccountStorageOptInStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnAccountStorageOptInStateChanged
//go:noescape
func CallHasOnAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnAccountStorageOptInStateChanged
//go:noescape
func TryHasOnAccountStorageOptInStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnInsecureCredentialsChanged
//go:noescape
func HasFuncOnInsecureCredentialsChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnInsecureCredentialsChanged
//go:noescape
func FuncOnInsecureCredentialsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnInsecureCredentialsChanged
//go:noescape
func CallOnInsecureCredentialsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnInsecureCredentialsChanged
//go:noescape
func TryOnInsecureCredentialsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffInsecureCredentialsChanged
//go:noescape
func HasFuncOffInsecureCredentialsChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffInsecureCredentialsChanged
//go:noescape
func FuncOffInsecureCredentialsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffInsecureCredentialsChanged
//go:noescape
func CallOffInsecureCredentialsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffInsecureCredentialsChanged
//go:noescape
func TryOffInsecureCredentialsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnInsecureCredentialsChanged
//go:noescape
func HasFuncHasOnInsecureCredentialsChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnInsecureCredentialsChanged
//go:noescape
func FuncHasOnInsecureCredentialsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnInsecureCredentialsChanged
//go:noescape
func CallHasOnInsecureCredentialsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnInsecureCredentialsChanged
//go:noescape
func TryHasOnInsecureCredentialsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnPasswordCheckStatusChanged
//go:noescape
func HasFuncOnPasswordCheckStatusChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnPasswordCheckStatusChanged
//go:noescape
func FuncOnPasswordCheckStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnPasswordCheckStatusChanged
//go:noescape
func CallOnPasswordCheckStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnPasswordCheckStatusChanged
//go:noescape
func TryOnPasswordCheckStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffPasswordCheckStatusChanged
//go:noescape
func HasFuncOffPasswordCheckStatusChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffPasswordCheckStatusChanged
//go:noescape
func FuncOffPasswordCheckStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffPasswordCheckStatusChanged
//go:noescape
func CallOffPasswordCheckStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffPasswordCheckStatusChanged
//go:noescape
func TryOffPasswordCheckStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnPasswordCheckStatusChanged
//go:noescape
func HasFuncHasOnPasswordCheckStatusChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnPasswordCheckStatusChanged
//go:noescape
func FuncHasOnPasswordCheckStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnPasswordCheckStatusChanged
//go:noescape
func CallHasOnPasswordCheckStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnPasswordCheckStatusChanged
//go:noescape
func TryHasOnPasswordCheckStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnPasswordExceptionsListChanged
//go:noescape
func HasFuncOnPasswordExceptionsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnPasswordExceptionsListChanged
//go:noescape
func FuncOnPasswordExceptionsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnPasswordExceptionsListChanged
//go:noescape
func CallOnPasswordExceptionsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnPasswordExceptionsListChanged
//go:noescape
func TryOnPasswordExceptionsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffPasswordExceptionsListChanged
//go:noescape
func HasFuncOffPasswordExceptionsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffPasswordExceptionsListChanged
//go:noescape
func FuncOffPasswordExceptionsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffPasswordExceptionsListChanged
//go:noescape
func CallOffPasswordExceptionsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffPasswordExceptionsListChanged
//go:noescape
func TryOffPasswordExceptionsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnPasswordExceptionsListChanged
//go:noescape
func HasFuncHasOnPasswordExceptionsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnPasswordExceptionsListChanged
//go:noescape
func FuncHasOnPasswordExceptionsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnPasswordExceptionsListChanged
//go:noescape
func CallHasOnPasswordExceptionsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnPasswordExceptionsListChanged
//go:noescape
func TryHasOnPasswordExceptionsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnPasswordManagerAuthTimeout
//go:noescape
func HasFuncOnPasswordManagerAuthTimeout() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnPasswordManagerAuthTimeout
//go:noescape
func FuncOnPasswordManagerAuthTimeout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnPasswordManagerAuthTimeout
//go:noescape
func CallOnPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnPasswordManagerAuthTimeout
//go:noescape
func TryOnPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffPasswordManagerAuthTimeout
//go:noescape
func HasFuncOffPasswordManagerAuthTimeout() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffPasswordManagerAuthTimeout
//go:noescape
func FuncOffPasswordManagerAuthTimeout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffPasswordManagerAuthTimeout
//go:noescape
func CallOffPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffPasswordManagerAuthTimeout
//go:noescape
func TryOffPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnPasswordManagerAuthTimeout
//go:noescape
func HasFuncHasOnPasswordManagerAuthTimeout() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnPasswordManagerAuthTimeout
//go:noescape
func FuncHasOnPasswordManagerAuthTimeout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnPasswordManagerAuthTimeout
//go:noescape
func CallHasOnPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnPasswordManagerAuthTimeout
//go:noescape
func TryHasOnPasswordManagerAuthTimeout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnPasswordsFileExportProgress
//go:noescape
func HasFuncOnPasswordsFileExportProgress() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnPasswordsFileExportProgress
//go:noescape
func FuncOnPasswordsFileExportProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnPasswordsFileExportProgress
//go:noescape
func CallOnPasswordsFileExportProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnPasswordsFileExportProgress
//go:noescape
func TryOnPasswordsFileExportProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffPasswordsFileExportProgress
//go:noescape
func HasFuncOffPasswordsFileExportProgress() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffPasswordsFileExportProgress
//go:noescape
func FuncOffPasswordsFileExportProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffPasswordsFileExportProgress
//go:noescape
func CallOffPasswordsFileExportProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffPasswordsFileExportProgress
//go:noescape
func TryOffPasswordsFileExportProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnPasswordsFileExportProgress
//go:noescape
func HasFuncHasOnPasswordsFileExportProgress() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnPasswordsFileExportProgress
//go:noescape
func FuncHasOnPasswordsFileExportProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnPasswordsFileExportProgress
//go:noescape
func CallHasOnPasswordsFileExportProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnPasswordsFileExportProgress
//go:noescape
func TryHasOnPasswordsFileExportProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OnSavedPasswordsListChanged
//go:noescape
func HasFuncOnSavedPasswordsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OnSavedPasswordsListChanged
//go:noescape
func FuncOnSavedPasswordsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OnSavedPasswordsListChanged
//go:noescape
func CallOnSavedPasswordsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OnSavedPasswordsListChanged
//go:noescape
func TryOnSavedPasswordsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OffSavedPasswordsListChanged
//go:noescape
func HasFuncOffSavedPasswordsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OffSavedPasswordsListChanged
//go:noescape
func FuncOffSavedPasswordsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OffSavedPasswordsListChanged
//go:noescape
func CallOffSavedPasswordsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OffSavedPasswordsListChanged
//go:noescape
func TryOffSavedPasswordsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_HasOnSavedPasswordsListChanged
//go:noescape
func HasFuncHasOnSavedPasswordsListChanged() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_HasOnSavedPasswordsListChanged
//go:noescape
func FuncHasOnSavedPasswordsListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_HasOnSavedPasswordsListChanged
//go:noescape
func CallHasOnSavedPasswordsListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_HasOnSavedPasswordsListChanged
//go:noescape
func TryHasOnSavedPasswordsListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_OptInForAccountStorage
//go:noescape
func HasFuncOptInForAccountStorage() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_OptInForAccountStorage
//go:noescape
func FuncOptInForAccountStorage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_OptInForAccountStorage
//go:noescape
func CallOptInForAccountStorage(
	retPtr unsafe.Pointer,
	optIn js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_OptInForAccountStorage
//go:noescape
func TryOptInForAccountStorage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	optIn js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RecordPasswordsPageAccessInSettings
//go:noescape
func HasFuncRecordPasswordsPageAccessInSettings() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RecordPasswordsPageAccessInSettings
//go:noescape
func FuncRecordPasswordsPageAccessInSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RecordPasswordsPageAccessInSettings
//go:noescape
func CallRecordPasswordsPageAccessInSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_RecordPasswordsPageAccessInSettings
//go:noescape
func TryRecordPasswordsPageAccessInSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RemoveCredential
//go:noescape
func HasFuncRemoveCredential() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RemoveCredential
//go:noescape
func FuncRemoveCredential(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RemoveCredential
//go:noescape
func CallRemoveCredential(
	retPtr unsafe.Pointer,
	id int32,
	fromStores uint32)

//go:wasmimport plat/js/webext/passwordsprivate try_RemoveCredential
//go:noescape
func TryRemoveCredential(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32,
	fromStores uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RemovePasswordException
//go:noescape
func HasFuncRemovePasswordException() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RemovePasswordException
//go:noescape
func FuncRemovePasswordException(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RemovePasswordException
//go:noescape
func CallRemovePasswordException(
	retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/webext/passwordsprivate try_RemovePasswordException
//go:noescape
func TryRemovePasswordException(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RequestCredentialsDetails
//go:noescape
func HasFuncRequestCredentialsDetails() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RequestCredentialsDetails
//go:noescape
func FuncRequestCredentialsDetails(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RequestCredentialsDetails
//go:noescape
func CallRequestCredentialsDetails(
	retPtr unsafe.Pointer,
	ids js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_RequestCredentialsDetails
//go:noescape
func TryRequestCredentialsDetails(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ids js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RequestExportProgressStatus
//go:noescape
func HasFuncRequestExportProgressStatus() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RequestExportProgressStatus
//go:noescape
func FuncRequestExportProgressStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RequestExportProgressStatus
//go:noescape
func CallRequestExportProgressStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_RequestExportProgressStatus
//go:noescape
func TryRequestExportProgressStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_RequestPlaintextPassword
//go:noescape
func HasFuncRequestPlaintextPassword() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_RequestPlaintextPassword
//go:noescape
func FuncRequestPlaintextPassword(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_RequestPlaintextPassword
//go:noescape
func CallRequestPlaintextPassword(
	retPtr unsafe.Pointer,
	id int32,
	reason uint32)

//go:wasmimport plat/js/webext/passwordsprivate try_RequestPlaintextPassword
//go:noescape
func TryRequestPlaintextPassword(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32,
	reason uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ResetImporter
//go:noescape
func HasFuncResetImporter() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ResetImporter
//go:noescape
func FuncResetImporter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ResetImporter
//go:noescape
func CallResetImporter(
	retPtr unsafe.Pointer,
	deleteFile js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_ResetImporter
//go:noescape
func TryResetImporter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deleteFile js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_SharePassword
//go:noescape
func HasFuncSharePassword() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_SharePassword
//go:noescape
func FuncSharePassword(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_SharePassword
//go:noescape
func CallSharePassword(
	retPtr unsafe.Pointer,
	id int32,
	recipients js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_SharePassword
//go:noescape
func TrySharePassword(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32,
	recipients js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ShowAddShortcutDialog
//go:noescape
func HasFuncShowAddShortcutDialog() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ShowAddShortcutDialog
//go:noescape
func FuncShowAddShortcutDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ShowAddShortcutDialog
//go:noescape
func CallShowAddShortcutDialog(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_ShowAddShortcutDialog
//go:noescape
func TryShowAddShortcutDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_ShowExportedFileInShell
//go:noescape
func HasFuncShowExportedFileInShell() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_ShowExportedFileInShell
//go:noescape
func FuncShowExportedFileInShell(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_ShowExportedFileInShell
//go:noescape
func CallShowExportedFileInShell(
	retPtr unsafe.Pointer,
	file_path js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate try_ShowExportedFileInShell
//go:noescape
func TryShowExportedFileInShell(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	file_path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_StartPasswordCheck
//go:noescape
func HasFuncStartPasswordCheck() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_StartPasswordCheck
//go:noescape
func FuncStartPasswordCheck(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_StartPasswordCheck
//go:noescape
func CallStartPasswordCheck(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_StartPasswordCheck
//go:noescape
func TryStartPasswordCheck(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_SwitchBiometricAuthBeforeFillingState
//go:noescape
func HasFuncSwitchBiometricAuthBeforeFillingState() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_SwitchBiometricAuthBeforeFillingState
//go:noescape
func FuncSwitchBiometricAuthBeforeFillingState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_SwitchBiometricAuthBeforeFillingState
//go:noescape
func CallSwitchBiometricAuthBeforeFillingState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_SwitchBiometricAuthBeforeFillingState
//go:noescape
func TrySwitchBiometricAuthBeforeFillingState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_UndoRemoveSavedPasswordOrException
//go:noescape
func HasFuncUndoRemoveSavedPasswordOrException() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_UndoRemoveSavedPasswordOrException
//go:noescape
func FuncUndoRemoveSavedPasswordOrException(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_UndoRemoveSavedPasswordOrException
//go:noescape
func CallUndoRemoveSavedPasswordOrException(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_UndoRemoveSavedPasswordOrException
//go:noescape
func TryUndoRemoveSavedPasswordOrException(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/passwordsprivate has_UnmuteInsecureCredential
//go:noescape
func HasFuncUnmuteInsecureCredential() js.Ref

//go:wasmimport plat/js/webext/passwordsprivate func_UnmuteInsecureCredential
//go:noescape
func FuncUnmuteInsecureCredential(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate call_UnmuteInsecureCredential
//go:noescape
func CallUnmuteInsecureCredential(
	retPtr unsafe.Pointer,
	credential unsafe.Pointer)

//go:wasmimport plat/js/webext/passwordsprivate try_UnmuteInsecureCredential
//go:noescape
func TryUnmuteInsecureCredential(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	credential unsafe.Pointer) (ok js.Ref)

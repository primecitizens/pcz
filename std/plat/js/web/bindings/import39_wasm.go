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

//go:wasmimport plat/js/web constof_IDBRequestReadyState
//go:noescape
func ConstOfIDBRequestReadyState(str js.Ref) uint32

//go:wasmimport plat/js/web get_IDBRequest_Result
//go:noescape
func GetIDBRequestResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBRequest_Error
//go:noescape
func GetIDBRequestError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBRequest_Source
//go:noescape
func GetIDBRequestSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBRequest_Transaction
//go:noescape
func GetIDBRequestTransaction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBRequest_ReadyState
//go:noescape
func GetIDBRequestReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursor_Source
//go:noescape
func GetIDBCursorSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursor_Direction
//go:noescape
func GetIDBCursorDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursor_Key
//go:noescape
func GetIDBCursorKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursor_PrimaryKey
//go:noescape
func GetIDBCursorPrimaryKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursor_Request
//go:noescape
func GetIDBCursorRequest(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_Advance
//go:noescape
func HasFuncIDBCursorAdvance(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_Advance
//go:noescape
func FuncIDBCursorAdvance(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_Advance
//go:noescape
func CallIDBCursorAdvance(
	this js.Ref, retPtr unsafe.Pointer,
	count uint32)

//go:wasmimport plat/js/web try_IDBCursor_Advance
//go:noescape
func TryIDBCursorAdvance(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_Continue
//go:noescape
func HasFuncIDBCursorContinue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_Continue
//go:noescape
func FuncIDBCursorContinue(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_Continue
//go:noescape
func CallIDBCursorContinue(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_IDBCursor_Continue
//go:noescape
func TryIDBCursorContinue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_Continue1
//go:noescape
func HasFuncIDBCursorContinue1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_Continue1
//go:noescape
func FuncIDBCursorContinue1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_Continue1
//go:noescape
func CallIDBCursorContinue1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBCursor_Continue1
//go:noescape
func TryIDBCursorContinue1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_ContinuePrimaryKey
//go:noescape
func HasFuncIDBCursorContinuePrimaryKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_ContinuePrimaryKey
//go:noescape
func FuncIDBCursorContinuePrimaryKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_ContinuePrimaryKey
//go:noescape
func CallIDBCursorContinuePrimaryKey(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref,
	primaryKey js.Ref)

//go:wasmimport plat/js/web try_IDBCursor_ContinuePrimaryKey
//go:noescape
func TryIDBCursorContinuePrimaryKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref,
	primaryKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_Update
//go:noescape
func HasFuncIDBCursorUpdate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_Update
//go:noescape
func FuncIDBCursorUpdate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_Update
//go:noescape
func CallIDBCursorUpdate(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_IDBCursor_Update
//go:noescape
func TryIDBCursorUpdate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBCursor_Delete
//go:noescape
func HasFuncIDBCursorDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBCursor_Delete
//go:noescape
func FuncIDBCursorDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBCursor_Delete
//go:noescape
func CallIDBCursorDelete(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBCursor_Delete
//go:noescape
func TryIDBCursorDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBCursorWithValue_Value
//go:noescape
func GetIDBCursorWithValueValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBKeyRange_Lower
//go:noescape
func GetIDBKeyRangeLower(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBKeyRange_Upper
//go:noescape
func GetIDBKeyRangeUpper(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBKeyRange_LowerOpen
//go:noescape
func GetIDBKeyRangeLowerOpen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBKeyRange_UpperOpen
//go:noescape
func GetIDBKeyRangeUpperOpen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_Only
//go:noescape
func HasFuncIDBKeyRangeOnly(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_Only
//go:noescape
func FuncIDBKeyRangeOnly(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_Only
//go:noescape
func CallIDBKeyRangeOnly(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_Only
//go:noescape
func TryIDBKeyRangeOnly(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_LowerBound
//go:noescape
func HasFuncIDBKeyRangeLowerBound(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_LowerBound
//go:noescape
func FuncIDBKeyRangeLowerBound(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_LowerBound
//go:noescape
func CallIDBKeyRangeLowerBound(
	this js.Ref, retPtr unsafe.Pointer,
	lower js.Ref,
	open js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_LowerBound
//go:noescape
func TryIDBKeyRangeLowerBound(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lower js.Ref,
	open js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_LowerBound1
//go:noescape
func HasFuncIDBKeyRangeLowerBound1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_LowerBound1
//go:noescape
func FuncIDBKeyRangeLowerBound1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_LowerBound1
//go:noescape
func CallIDBKeyRangeLowerBound1(
	this js.Ref, retPtr unsafe.Pointer,
	lower js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_LowerBound1
//go:noescape
func TryIDBKeyRangeLowerBound1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lower js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_UpperBound
//go:noescape
func HasFuncIDBKeyRangeUpperBound(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_UpperBound
//go:noescape
func FuncIDBKeyRangeUpperBound(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_UpperBound
//go:noescape
func CallIDBKeyRangeUpperBound(
	this js.Ref, retPtr unsafe.Pointer,
	upper js.Ref,
	open js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_UpperBound
//go:noescape
func TryIDBKeyRangeUpperBound(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	upper js.Ref,
	open js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_UpperBound1
//go:noescape
func HasFuncIDBKeyRangeUpperBound1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_UpperBound1
//go:noescape
func FuncIDBKeyRangeUpperBound1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_UpperBound1
//go:noescape
func CallIDBKeyRangeUpperBound1(
	this js.Ref, retPtr unsafe.Pointer,
	upper js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_UpperBound1
//go:noescape
func TryIDBKeyRangeUpperBound1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	upper js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_Bound
//go:noescape
func HasFuncIDBKeyRangeBound(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_Bound
//go:noescape
func FuncIDBKeyRangeBound(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_Bound
//go:noescape
func CallIDBKeyRangeBound(
	this js.Ref, retPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref,
	lowerOpen js.Ref,
	upperOpen js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_Bound
//go:noescape
func TryIDBKeyRangeBound(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref,
	lowerOpen js.Ref,
	upperOpen js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_Bound1
//go:noescape
func HasFuncIDBKeyRangeBound1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_Bound1
//go:noescape
func FuncIDBKeyRangeBound1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_Bound1
//go:noescape
func CallIDBKeyRangeBound1(
	this js.Ref, retPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref,
	lowerOpen js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_Bound1
//go:noescape
func TryIDBKeyRangeBound1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref,
	lowerOpen js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_Bound2
//go:noescape
func HasFuncIDBKeyRangeBound2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_Bound2
//go:noescape
func FuncIDBKeyRangeBound2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_Bound2
//go:noescape
func CallIDBKeyRangeBound2(
	this js.Ref, retPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_Bound2
//go:noescape
func TryIDBKeyRangeBound2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lower js.Ref,
	upper js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBKeyRange_Includes
//go:noescape
func HasFuncIDBKeyRangeIncludes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBKeyRange_Includes
//go:noescape
func FuncIDBKeyRangeIncludes(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IDBKeyRange_Includes
//go:noescape
func CallIDBKeyRangeIncludes(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_IDBKeyRange_Includes
//go:noescape
func TryIDBKeyRangeIncludes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_IDBVersionChangeEventInit
//go:noescape
func IDBVersionChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IDBVersionChangeEventInit
//go:noescape
func IDBVersionChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_IDBVersionChangeEvent_IDBVersionChangeEvent
//go:noescape
func NewIDBVersionChangeEventByIDBVersionChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_IDBVersionChangeEvent_IDBVersionChangeEvent1
//go:noescape
func NewIDBVersionChangeEventByIDBVersionChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_IDBVersionChangeEvent_OldVersion
//go:noescape
func GetIDBVersionChangeEventOldVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBVersionChangeEvent_NewVersion
//go:noescape
func GetIDBVersionChangeEventNewVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IdentityCredential_Token
//go:noescape
func GetIdentityCredentialToken(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_IdentityUserInfo
//go:noescape
func IdentityUserInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityUserInfo
//go:noescape
func IdentityUserInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_IdentityProvider_GetUserInfo
//go:noescape
func HasFuncIdentityProviderGetUserInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdentityProvider_GetUserInfo
//go:noescape
func FuncIdentityProviderGetUserInfo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IdentityProvider_GetUserInfo
//go:noescape
func CallIdentityProviderGetUserInfo(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_IdentityProvider_GetUserInfo
//go:noescape
func TryIdentityProviderGetUserInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_IdentityProviderIcon
//go:noescape
func IdentityProviderIconJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderIcon
//go:noescape
func IdentityProviderIconJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderBranding
//go:noescape
func IdentityProviderBrandingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderBranding
//go:noescape
func IdentityProviderBrandingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderAPIConfig
//go:noescape
func IdentityProviderAPIConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderAPIConfig
//go:noescape
func IdentityProviderAPIConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderAccount
//go:noescape
func IdentityProviderAccountJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderAccount
//go:noescape
func IdentityProviderAccountJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderAccountList
//go:noescape
func IdentityProviderAccountListJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderAccountList
//go:noescape
func IdentityProviderAccountListJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderClientMetadata
//go:noescape
func IdentityProviderClientMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderClientMetadata
//go:noescape
func IdentityProviderClientMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderToken
//go:noescape
func IdentityProviderTokenJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderToken
//go:noescape
func IdentityProviderTokenJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderWellKnown
//go:noescape
func IdentityProviderWellKnownJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderWellKnown
//go:noescape
func IdentityProviderWellKnownJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdleOptions
//go:noescape
func IdleOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdleOptions
//go:noescape
func IdleOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_UserIdleState
//go:noescape
func ConstOfUserIdleState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ScreenIdleState
//go:noescape
func ConstOfScreenIdleState(str js.Ref) uint32

//go:wasmimport plat/js/web get_IdleDetector_UserState
//go:noescape
func GetIdleDetectorUserState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IdleDetector_ScreenState
//go:noescape
func GetIdleDetectorScreenState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IdleDetector_RequestPermission
//go:noescape
func HasFuncIdleDetectorRequestPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdleDetector_RequestPermission
//go:noescape
func FuncIdleDetectorRequestPermission(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IdleDetector_RequestPermission
//go:noescape
func CallIdleDetectorRequestPermission(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IdleDetector_RequestPermission
//go:noescape
func TryIdleDetectorRequestPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IdleDetector_Start
//go:noescape
func HasFuncIdleDetectorStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdleDetector_Start
//go:noescape
func FuncIdleDetectorStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IdleDetector_Start
//go:noescape
func CallIdleDetectorStart(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_IdleDetector_Start
//go:noescape
func TryIdleDetectorStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IdleDetector_Start1
//go:noescape
func HasFuncIdleDetectorStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdleDetector_Start1
//go:noescape
func FuncIdleDetectorStart1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_IdleDetector_Start1
//go:noescape
func CallIdleDetectorStart1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IdleDetector_Start1
//go:noescape
func TryIdleDetectorStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ImageBitmapRenderingContextSettings
//go:noescape
func ImageBitmapRenderingContextSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageBitmapRenderingContextSettings
//go:noescape
func ImageBitmapRenderingContextSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PhotoSettings
//go:noescape
func PhotoSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PhotoSettings
//go:noescape
func PhotoSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RedEyeReduction
//go:noescape
func ConstOfRedEyeReduction(str js.Ref) uint32

//go:wasmimport plat/js/web store_PhotoCapabilities
//go:noescape
func PhotoCapabilitiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PhotoCapabilities
//go:noescape
func PhotoCapabilitiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ImageCapture_ImageCapture
//go:noescape
func NewImageCaptureByImageCapture(
	videoTrack js.Ref) js.Ref

//go:wasmimport plat/js/web get_ImageCapture_Track
//go:noescape
func GetImageCaptureTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageCapture_TakePhoto
//go:noescape
func HasFuncImageCaptureTakePhoto(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageCapture_TakePhoto
//go:noescape
func FuncImageCaptureTakePhoto(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageCapture_TakePhoto
//go:noescape
func CallImageCaptureTakePhoto(
	this js.Ref, retPtr unsafe.Pointer,
	photoSettings unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageCapture_TakePhoto
//go:noescape
func TryImageCaptureTakePhoto(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	photoSettings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageCapture_TakePhoto1
//go:noescape
func HasFuncImageCaptureTakePhoto1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageCapture_TakePhoto1
//go:noescape
func FuncImageCaptureTakePhoto1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageCapture_TakePhoto1
//go:noescape
func CallImageCaptureTakePhoto1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageCapture_TakePhoto1
//go:noescape
func TryImageCaptureTakePhoto1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageCapture_GetPhotoCapabilities
//go:noescape
func HasFuncImageCaptureGetPhotoCapabilities(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageCapture_GetPhotoCapabilities
//go:noescape
func FuncImageCaptureGetPhotoCapabilities(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageCapture_GetPhotoCapabilities
//go:noescape
func CallImageCaptureGetPhotoCapabilities(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageCapture_GetPhotoCapabilities
//go:noescape
func TryImageCaptureGetPhotoCapabilities(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageCapture_GetPhotoSettings
//go:noescape
func HasFuncImageCaptureGetPhotoSettings(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageCapture_GetPhotoSettings
//go:noescape
func FuncImageCaptureGetPhotoSettings(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageCapture_GetPhotoSettings
//go:noescape
func CallImageCaptureGetPhotoSettings(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageCapture_GetPhotoSettings
//go:noescape
func TryImageCaptureGetPhotoSettings(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageCapture_GrabFrame
//go:noescape
func HasFuncImageCaptureGrabFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageCapture_GrabFrame
//go:noescape
func FuncImageCaptureGrabFrame(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageCapture_GrabFrame
//go:noescape
func CallImageCaptureGrabFrame(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageCapture_GrabFrame
//go:noescape
func TryImageCaptureGrabFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ImageDecodeOptions
//go:noescape
func ImageDecodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageDecodeOptions
//go:noescape
func ImageDecodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ImageDecodeResult
//go:noescape
func ImageDecodeResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageDecodeResult
//go:noescape
func ImageDecodeResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ImageDecoderInit
//go:noescape
func ImageDecoderInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageDecoderInit
//go:noescape
func ImageDecoderInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ImageTrack_Animated
//go:noescape
func GetImageTrackAnimated(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrack_FrameCount
//go:noescape
func GetImageTrackFrameCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrack_RepetitionCount
//go:noescape
func GetImageTrackRepetitionCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrack_Selected
//go:noescape
func GetImageTrackSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ImageTrack_Selected
//go:noescape
func SetImageTrackSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ImageTrackList_Ready
//go:noescape
func GetImageTrackListReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrackList_Length
//go:noescape
func GetImageTrackListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrackList_SelectedIndex
//go:noescape
func GetImageTrackListSelectedIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageTrackList_SelectedTrack
//go:noescape
func GetImageTrackListSelectedTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageTrackList_Get
//go:noescape
func HasFuncImageTrackListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageTrackList_Get
//go:noescape
func FuncImageTrackListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageTrackList_Get
//go:noescape
func CallImageTrackListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_ImageTrackList_Get
//go:noescape
func TryImageTrackListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_ImageDecoder_ImageDecoder
//go:noescape
func NewImageDecoderByImageDecoder(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_ImageDecoder_Type
//go:noescape
func GetImageDecoderType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageDecoder_Complete
//go:noescape
func GetImageDecoderComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageDecoder_Completed
//go:noescape
func GetImageDecoderCompleted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageDecoder_Tracks
//go:noescape
func GetImageDecoderTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageDecoder_Decode
//go:noescape
func HasFuncImageDecoderDecode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageDecoder_Decode
//go:noescape
func FuncImageDecoderDecode(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageDecoder_Decode
//go:noescape
func CallImageDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageDecoder_Decode
//go:noescape
func TryImageDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageDecoder_Decode1
//go:noescape
func HasFuncImageDecoderDecode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageDecoder_Decode1
//go:noescape
func FuncImageDecoderDecode1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageDecoder_Decode1
//go:noescape
func CallImageDecoderDecode1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageDecoder_Decode1
//go:noescape
func TryImageDecoderDecode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageDecoder_Reset
//go:noescape
func HasFuncImageDecoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageDecoder_Reset
//go:noescape
func FuncImageDecoderReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageDecoder_Reset
//go:noescape
func CallImageDecoderReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageDecoder_Reset
//go:noescape
func TryImageDecoderReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageDecoder_Close
//go:noescape
func HasFuncImageDecoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageDecoder_Close
//go:noescape
func FuncImageDecoderClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageDecoder_Close
//go:noescape
func CallImageDecoderClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ImageDecoder_Close
//go:noescape
func TryImageDecoderClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageDecoder_IsTypeSupported
//go:noescape
func HasFuncImageDecoderIsTypeSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageDecoder_IsTypeSupported
//go:noescape
func FuncImageDecoderIsTypeSupported(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageDecoder_IsTypeSupported
//go:noescape
func CallImageDecoderIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_ImageDecoder_IsTypeSupported
//go:noescape
func TryImageDecoderIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_ImportExportKind
//go:noescape
func ConstOfImportExportKind(str js.Ref) uint32

//go:wasmimport plat/js/web store_InputDeviceCapabilitiesInit
//go:noescape
func InputDeviceCapabilitiesInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_InputDeviceCapabilitiesInit
//go:noescape
func InputDeviceCapabilitiesInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_InputDeviceCapabilities_InputDeviceCapabilities
//go:noescape
func NewInputDeviceCapabilitiesByInputDeviceCapabilities(
	deviceInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_InputDeviceCapabilities_InputDeviceCapabilities1
//go:noescape
func NewInputDeviceCapabilitiesByInputDeviceCapabilities1() js.Ref

//go:wasmimport plat/js/web get_InputDeviceCapabilities_FiresTouchEvents
//go:noescape
func GetInputDeviceCapabilitiesFiresTouchEvents(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InputDeviceCapabilities_PointerMovementScrolls
//go:noescape
func GetInputDeviceCapabilitiesPointerMovementScrolls(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InputDeviceInfo_GetCapabilities
//go:noescape
func HasFuncInputDeviceInfoGetCapabilities(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InputDeviceInfo_GetCapabilities
//go:noescape
func FuncInputDeviceInfoGetCapabilities(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InputDeviceInfo_GetCapabilities
//go:noescape
func CallInputDeviceInfoGetCapabilities(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_InputDeviceInfo_GetCapabilities
//go:noescape
func TryInputDeviceInfoGetCapabilities(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_InputEventInit
//go:noescape
func InputEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_InputEventInit
//go:noescape
func InputEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_InputEvent_InputEvent
//go:noescape
func NewInputEventByInputEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_InputEvent_InputEvent1
//go:noescape
func NewInputEventByInputEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_InputEvent_Data
//go:noescape
func GetInputEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InputEvent_IsComposing
//go:noescape
func GetInputEventIsComposing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InputEvent_InputType
//go:noescape
func GetInputEventInputType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_InputEvent_DataTransfer
//go:noescape
func GetInputEventDataTransfer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_InputEvent_GetTargetRanges
//go:noescape
func HasFuncInputEventGetTargetRanges(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_InputEvent_GetTargetRanges
//go:noescape
func FuncInputEventGetTargetRanges(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_InputEvent_GetTargetRanges
//go:noescape
func CallInputEventGetTargetRanges(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_InputEvent_GetTargetRanges
//go:noescape
func TryInputEventGetTargetRanges(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ModuleExportDescriptor
//go:noescape
func ModuleExportDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ModuleExportDescriptor
//go:noescape
func ModuleExportDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ModuleImportDescriptor
//go:noescape
func ModuleImportDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ModuleImportDescriptor
//go:noescape
func ModuleImportDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Module_Module
//go:noescape
func NewModuleByModule(
	bytes js.Ref) js.Ref

//go:wasmimport plat/js/web has_Module_Exports
//go:noescape
func HasFuncModuleExports(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Module_Exports
//go:noescape
func FuncModuleExports(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Module_Exports
//go:noescape
func CallModuleExports(
	this js.Ref, retPtr unsafe.Pointer,
	moduleObject js.Ref)

//go:wasmimport plat/js/web try_Module_Exports
//go:noescape
func TryModuleExports(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Module_Imports
//go:noescape
func HasFuncModuleImports(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Module_Imports
//go:noescape
func FuncModuleImports(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Module_Imports
//go:noescape
func CallModuleImports(
	this js.Ref, retPtr unsafe.Pointer,
	moduleObject js.Ref)

//go:wasmimport plat/js/web try_Module_Imports
//go:noescape
func TryModuleImports(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Module_CustomSections
//go:noescape
func HasFuncModuleCustomSections(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Module_CustomSections
//go:noescape
func FuncModuleCustomSections(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Module_CustomSections
//go:noescape
func CallModuleCustomSections(
	this js.Ref, retPtr unsafe.Pointer,
	moduleObject js.Ref,
	sectionName js.Ref)

//go:wasmimport plat/js/web try_Module_CustomSections
//go:noescape
func TryModuleCustomSections(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref,
	sectionName js.Ref) (ok js.Ref)

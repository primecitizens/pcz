// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web constof_ServiceWorkerState
//go:noescape
func ConstOfServiceWorkerState(str js.Ref) uint32

//go:wasmimport plat/js/web get_ServiceWorker_ScriptURL
//go:noescape
func GetServiceWorkerScriptURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorker_State
//go:noescape
func GetServiceWorkerState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorker_PostMessage
//go:noescape
func HasServiceWorkerPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage
//go:noescape
func ServiceWorkerPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage
//go:noescape
func CallServiceWorkerPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref)

//go:wasmimport plat/js/web try_ServiceWorker_PostMessage
//go:noescape
func TryServiceWorkerPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorker_PostMessage1
//go:noescape
func HasServiceWorkerPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage1
//go:noescape
func ServiceWorkerPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage1
//go:noescape
func CallServiceWorkerPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorker_PostMessage1
//go:noescape
func TryServiceWorkerPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorker_PostMessage2
//go:noescape
func HasServiceWorkerPostMessage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage2
//go:noescape
func ServiceWorkerPostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage2
//go:noescape
func CallServiceWorkerPostMessage2(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_ServiceWorker_PostMessage2
//go:noescape
func TryServiceWorkerPostMessage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_NavigationPreloadState
//go:noescape
func NavigationPreloadStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationPreloadState
//go:noescape
func NavigationPreloadStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_NavigationPreloadManager_Enable
//go:noescape
func HasNavigationPreloadManagerEnable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_Enable
//go:noescape
func NavigationPreloadManagerEnableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_Enable
//go:noescape
func CallNavigationPreloadManagerEnable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NavigationPreloadManager_Enable
//go:noescape
func TryNavigationPreloadManagerEnable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NavigationPreloadManager_Disable
//go:noescape
func HasNavigationPreloadManagerDisable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_Disable
//go:noescape
func NavigationPreloadManagerDisableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_Disable
//go:noescape
func CallNavigationPreloadManagerDisable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NavigationPreloadManager_Disable
//go:noescape
func TryNavigationPreloadManagerDisable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NavigationPreloadManager_SetHeaderValue
//go:noescape
func HasNavigationPreloadManagerSetHeaderValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_SetHeaderValue
//go:noescape
func NavigationPreloadManagerSetHeaderValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_SetHeaderValue
//go:noescape
func CallNavigationPreloadManagerSetHeaderValue(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_NavigationPreloadManager_SetHeaderValue
//go:noescape
func TryNavigationPreloadManagerSetHeaderValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_NavigationPreloadManager_GetState
//go:noescape
func HasNavigationPreloadManagerGetState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_GetState
//go:noescape
func NavigationPreloadManagerGetStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_GetState
//go:noescape
func CallNavigationPreloadManagerGetState(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NavigationPreloadManager_GetState
//go:noescape
func TryNavigationPreloadManagerGetState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ServiceWorkerUpdateViaCache
//go:noescape
func ConstOfServiceWorkerUpdateViaCache(str js.Ref) uint32

//go:wasmimport plat/js/web has_SyncManager_Register
//go:noescape
func HasSyncManagerRegister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SyncManager_Register
//go:noescape
func SyncManagerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SyncManager_Register
//go:noescape
func CallSyncManagerRegister(
	this js.Ref, retPtr unsafe.Pointer,
	tag js.Ref)

//go:wasmimport plat/js/web try_SyncManager_Register
//go:noescape
func TrySyncManagerRegister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tag js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SyncManager_GetTags
//go:noescape
func HasSyncManagerGetTags(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SyncManager_GetTags
//go:noescape
func SyncManagerGetTagsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SyncManager_GetTags
//go:noescape
func CallSyncManagerGetTags(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SyncManager_GetTags
//go:noescape
func TrySyncManagerGetTags(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PushEncryptionKeyName
//go:noescape
func ConstOfPushEncryptionKeyName(str js.Ref) uint32

//go:wasmimport plat/js/web store_PushSubscriptionJSON
//go:noescape
func PushSubscriptionJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PushSubscriptionJSON
//go:noescape
func PushSubscriptionJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PushSubscriptionOptions_UserVisibleOnly
//go:noescape
func GetPushSubscriptionOptionsUserVisibleOnly(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PushSubscriptionOptions_ApplicationServerKey
//go:noescape
func GetPushSubscriptionOptionsApplicationServerKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PushSubscription_Endpoint
//go:noescape
func GetPushSubscriptionEndpoint(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PushSubscription_ExpirationTime
//go:noescape
func GetPushSubscriptionExpirationTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PushSubscription_Options
//go:noescape
func GetPushSubscriptionOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushSubscription_GetKey
//go:noescape
func HasPushSubscriptionGetKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_GetKey
//go:noescape
func PushSubscriptionGetKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_GetKey
//go:noescape
func CallPushSubscriptionGetKey(
	this js.Ref, retPtr unsafe.Pointer,
	name uint32)

//go:wasmimport plat/js/web try_PushSubscription_GetKey
//go:noescape
func TryPushSubscriptionGetKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_PushSubscription_Unsubscribe
//go:noescape
func HasPushSubscriptionUnsubscribe(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_Unsubscribe
//go:noescape
func PushSubscriptionUnsubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_Unsubscribe
//go:noescape
func CallPushSubscriptionUnsubscribe(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushSubscription_Unsubscribe
//go:noescape
func TryPushSubscriptionUnsubscribe(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushSubscription_ToJSON
//go:noescape
func HasPushSubscriptionToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_ToJSON
//go:noescape
func PushSubscriptionToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_ToJSON
//go:noescape
func CallPushSubscriptionToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushSubscription_ToJSON
//go:noescape
func TryPushSubscriptionToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PushSubscriptionOptionsInit
//go:noescape
func PushSubscriptionOptionsInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PushSubscriptionOptionsInit
//go:noescape
func PushSubscriptionOptionsInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PushManager_SupportedContentEncodings
//go:noescape
func GetPushManagerSupportedContentEncodings(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushManager_Subscribe
//go:noescape
func HasPushManagerSubscribe(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushManager_Subscribe
//go:noescape
func PushManagerSubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_Subscribe
//go:noescape
func CallPushManagerSubscribe(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PushManager_Subscribe
//go:noescape
func TryPushManagerSubscribe(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushManager_Subscribe1
//go:noescape
func HasPushManagerSubscribe1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushManager_Subscribe1
//go:noescape
func PushManagerSubscribe1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_Subscribe1
//go:noescape
func CallPushManagerSubscribe1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushManager_Subscribe1
//go:noescape
func TryPushManagerSubscribe1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushManager_GetSubscription
//go:noescape
func HasPushManagerGetSubscription(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushManager_GetSubscription
//go:noescape
func PushManagerGetSubscriptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_GetSubscription
//go:noescape
func CallPushManagerGetSubscription(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushManager_GetSubscription
//go:noescape
func TryPushManagerGetSubscription(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushManager_PermissionState
//go:noescape
func HasPushManagerPermissionState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushManager_PermissionState
//go:noescape
func PushManagerPermissionStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_PermissionState
//go:noescape
func CallPushManagerPermissionState(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PushManager_PermissionState
//go:noescape
func TryPushManagerPermissionState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PushManager_PermissionState1
//go:noescape
func HasPushManagerPermissionState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PushManager_PermissionState1
//go:noescape
func PushManagerPermissionState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_PermissionState1
//go:noescape
func CallPushManagerPermissionState1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PushManager_PermissionState1
//go:noescape
func TryPushManagerPermissionState1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PaymentDelegation
//go:noescape
func ConstOfPaymentDelegation(str js.Ref) uint32

//go:wasmimport plat/js/web get_PaymentManager_UserHint
//go:noescape
func GetPaymentManagerUserHint(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_PaymentManager_UserHint
//go:noescape
func SetPaymentManagerUserHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_PaymentManager_EnableDelegations
//go:noescape
func HasPaymentManagerEnableDelegations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentManager_EnableDelegations
//go:noescape
func PaymentManagerEnableDelegationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentManager_EnableDelegations
//go:noescape
func CallPaymentManagerEnableDelegations(
	this js.Ref, retPtr unsafe.Pointer,
	delegations js.Ref)

//go:wasmimport plat/js/web try_PaymentManager_EnableDelegations
//go:noescape
func TryPaymentManagerEnableDelegations(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	delegations js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PeriodicSyncManager_Register
//go:noescape
func HasPeriodicSyncManagerRegister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Register
//go:noescape
func PeriodicSyncManagerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Register
//go:noescape
func CallPeriodicSyncManagerRegister(
	this js.Ref, retPtr unsafe.Pointer,
	tag js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PeriodicSyncManager_Register
//go:noescape
func TryPeriodicSyncManagerRegister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tag js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PeriodicSyncManager_Register1
//go:noescape
func HasPeriodicSyncManagerRegister1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Register1
//go:noescape
func PeriodicSyncManagerRegister1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Register1
//go:noescape
func CallPeriodicSyncManagerRegister1(
	this js.Ref, retPtr unsafe.Pointer,
	tag js.Ref)

//go:wasmimport plat/js/web try_PeriodicSyncManager_Register1
//go:noescape
func TryPeriodicSyncManagerRegister1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tag js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PeriodicSyncManager_GetTags
//go:noescape
func HasPeriodicSyncManagerGetTags(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_GetTags
//go:noescape
func PeriodicSyncManagerGetTagsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_GetTags
//go:noescape
func CallPeriodicSyncManagerGetTags(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PeriodicSyncManager_GetTags
//go:noescape
func TryPeriodicSyncManagerGetTags(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PeriodicSyncManager_Unregister
//go:noescape
func HasPeriodicSyncManagerUnregister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Unregister
//go:noescape
func PeriodicSyncManagerUnregisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Unregister
//go:noescape
func CallPeriodicSyncManagerUnregister(
	this js.Ref, retPtr unsafe.Pointer,
	tag js.Ref)

//go:wasmimport plat/js/web try_PeriodicSyncManager_Unregister
//go:noescape
func TryPeriodicSyncManagerUnregister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tag js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_CookieStoreGetOptions
//go:noescape
func CookieStoreGetOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieStoreGetOptions
//go:noescape
func CookieStoreGetOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CookieStoreManager_Subscribe
//go:noescape
func HasCookieStoreManagerSubscribe(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_Subscribe
//go:noescape
func CookieStoreManagerSubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_Subscribe
//go:noescape
func CallCookieStoreManagerSubscribe(
	this js.Ref, retPtr unsafe.Pointer,
	subscriptions js.Ref)

//go:wasmimport plat/js/web try_CookieStoreManager_Subscribe
//go:noescape
func TryCookieStoreManagerSubscribe(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	subscriptions js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStoreManager_GetSubscriptions
//go:noescape
func HasCookieStoreManagerGetSubscriptions(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_GetSubscriptions
//go:noescape
func CookieStoreManagerGetSubscriptionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_GetSubscriptions
//go:noescape
func CallCookieStoreManagerGetSubscriptions(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CookieStoreManager_GetSubscriptions
//go:noescape
func TryCookieStoreManagerGetSubscriptions(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CookieStoreManager_Unsubscribe
//go:noescape
func HasCookieStoreManagerUnsubscribe(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_Unsubscribe
//go:noescape
func CookieStoreManagerUnsubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_Unsubscribe
//go:noescape
func CallCookieStoreManagerUnsubscribe(
	this js.Ref, retPtr unsafe.Pointer,
	subscriptions js.Ref)

//go:wasmimport plat/js/web try_CookieStoreManager_Unsubscribe
//go:noescape
func TryCookieStoreManagerUnsubscribe(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	subscriptions js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_ContentCategory
//go:noescape
func ConstOfContentCategory(str js.Ref) uint32

//go:wasmimport plat/js/web store_ContentDescription
//go:noescape
func ContentDescriptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ContentDescription
//go:noescape
func ContentDescriptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_ContentIndex_Add
//go:noescape
func HasContentIndexAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_Add
//go:noescape
func ContentIndexAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContentIndex_Add
//go:noescape
func CallContentIndexAdd(
	this js.Ref, retPtr unsafe.Pointer,
	description unsafe.Pointer)

//go:wasmimport plat/js/web try_ContentIndex_Add
//go:noescape
func TryContentIndexAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	description unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ContentIndex_Delete
//go:noescape
func HasContentIndexDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_Delete
//go:noescape
func ContentIndexDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContentIndex_Delete
//go:noescape
func CallContentIndexDelete(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_ContentIndex_Delete
//go:noescape
func TryContentIndexDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ContentIndex_GetAll
//go:noescape
func HasContentIndexGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_GetAll
//go:noescape
func ContentIndexGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContentIndex_GetAll
//go:noescape
func CallContentIndexGetAll(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ContentIndex_GetAll
//go:noescape
func TryContentIndexGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Installing
//go:noescape
func GetServiceWorkerRegistrationInstalling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Waiting
//go:noescape
func GetServiceWorkerRegistrationWaiting(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Active
//go:noescape
func GetServiceWorkerRegistrationActive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_NavigationPreload
//go:noescape
func GetServiceWorkerRegistrationNavigationPreload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Scope
//go:noescape
func GetServiceWorkerRegistrationScope(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_UpdateViaCache
//go:noescape
func GetServiceWorkerRegistrationUpdateViaCache(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Sync
//go:noescape
func GetServiceWorkerRegistrationSync(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_BackgroundFetch
//go:noescape
func GetServiceWorkerRegistrationBackgroundFetch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PushManager
//go:noescape
func GetServiceWorkerRegistrationPushManager(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PaymentManager
//go:noescape
func GetServiceWorkerRegistrationPaymentManager(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PeriodicSync
//go:noescape
func GetServiceWorkerRegistrationPeriodicSync(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Cookies
//go:noescape
func GetServiceWorkerRegistrationCookies(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Index
//go:noescape
func GetServiceWorkerRegistrationIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_Update
//go:noescape
func HasServiceWorkerRegistrationUpdate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_Update
//go:noescape
func ServiceWorkerRegistrationUpdateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_Update
//go:noescape
func CallServiceWorkerRegistrationUpdate(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_Update
//go:noescape
func TryServiceWorkerRegistrationUpdate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_Unregister
//go:noescape
func HasServiceWorkerRegistrationUnregister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_Unregister
//go:noescape
func ServiceWorkerRegistrationUnregisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_Unregister
//go:noescape
func CallServiceWorkerRegistrationUnregister(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_Unregister
//go:noescape
func TryServiceWorkerRegistrationUnregister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_ShowNotification
//go:noescape
func HasServiceWorkerRegistrationShowNotification(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_ShowNotification
//go:noescape
func ServiceWorkerRegistrationShowNotificationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_ShowNotification
//go:noescape
func CallServiceWorkerRegistrationShowNotification(
	this js.Ref, retPtr unsafe.Pointer,
	title js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_ShowNotification
//go:noescape
func TryServiceWorkerRegistrationShowNotification(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	title js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_ShowNotification1
//go:noescape
func HasServiceWorkerRegistrationShowNotification1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_ShowNotification1
//go:noescape
func ServiceWorkerRegistrationShowNotification1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_ShowNotification1
//go:noescape
func CallServiceWorkerRegistrationShowNotification1(
	this js.Ref, retPtr unsafe.Pointer,
	title js.Ref)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_ShowNotification1
//go:noescape
func TryServiceWorkerRegistrationShowNotification1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	title js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_GetNotifications
//go:noescape
func HasServiceWorkerRegistrationGetNotifications(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_GetNotifications
//go:noescape
func ServiceWorkerRegistrationGetNotificationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_GetNotifications
//go:noescape
func CallServiceWorkerRegistrationGetNotifications(
	this js.Ref, retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_GetNotifications
//go:noescape
func TryServiceWorkerRegistrationGetNotifications(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerRegistration_GetNotifications1
//go:noescape
func HasServiceWorkerRegistrationGetNotifications1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_GetNotifications1
//go:noescape
func ServiceWorkerRegistrationGetNotifications1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_GetNotifications1
//go:noescape
func CallServiceWorkerRegistrationGetNotifications1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerRegistration_GetNotifications1
//go:noescape
func TryServiceWorkerRegistrationGetNotifications1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_WorkerType
//go:noescape
func ConstOfWorkerType(str js.Ref) uint32

//go:wasmimport plat/js/web store_RegistrationOptions
//go:noescape
func RegistrationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RegistrationOptions
//go:noescape
func RegistrationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerContainer_Controller
//go:noescape
func GetServiceWorkerContainerController(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ServiceWorkerContainer_Ready
//go:noescape
func GetServiceWorkerContainerReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_Register
//go:noescape
func HasServiceWorkerContainerRegister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_Register
//go:noescape
func ServiceWorkerContainerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_Register
//go:noescape
func CallServiceWorkerContainerRegister(
	this js.Ref, retPtr unsafe.Pointer,
	scriptURL js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_Register
//go:noescape
func TryServiceWorkerContainerRegister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scriptURL js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_Register1
//go:noescape
func HasServiceWorkerContainerRegister1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_Register1
//go:noescape
func ServiceWorkerContainerRegister1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_Register1
//go:noescape
func CallServiceWorkerContainerRegister1(
	this js.Ref, retPtr unsafe.Pointer,
	scriptURL js.Ref)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_Register1
//go:noescape
func TryServiceWorkerContainerRegister1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scriptURL js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_GetRegistration
//go:noescape
func HasServiceWorkerContainerGetRegistration(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistration
//go:noescape
func ServiceWorkerContainerGetRegistrationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistration
//go:noescape
func CallServiceWorkerContainerGetRegistration(
	this js.Ref, retPtr unsafe.Pointer,
	clientURL js.Ref)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_GetRegistration
//go:noescape
func TryServiceWorkerContainerGetRegistration(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	clientURL js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_GetRegistration1
//go:noescape
func HasServiceWorkerContainerGetRegistration1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistration1
//go:noescape
func ServiceWorkerContainerGetRegistration1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistration1
//go:noescape
func CallServiceWorkerContainerGetRegistration1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_GetRegistration1
//go:noescape
func TryServiceWorkerContainerGetRegistration1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_GetRegistrations
//go:noescape
func HasServiceWorkerContainerGetRegistrations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistrations
//go:noescape
func ServiceWorkerContainerGetRegistrationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistrations
//go:noescape
func CallServiceWorkerContainerGetRegistrations(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_GetRegistrations
//go:noescape
func TryServiceWorkerContainerGetRegistrations(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ServiceWorkerContainer_StartMessages
//go:noescape
func HasServiceWorkerContainerStartMessages(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_StartMessages
//go:noescape
func ServiceWorkerContainerStartMessagesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_StartMessages
//go:noescape
func CallServiceWorkerContainerStartMessages(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ServiceWorkerContainer_StartMessages
//go:noescape
func TryServiceWorkerContainerStartMessages(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MediaDeviceKind
//go:noescape
func ConstOfMediaDeviceKind(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaDeviceInfo_DeviceId
//go:noescape
func GetMediaDeviceInfoDeviceId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaDeviceInfo_Kind
//go:noescape
func GetMediaDeviceInfoKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaDeviceInfo_Label
//go:noescape
func GetMediaDeviceInfoLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaDeviceInfo_GroupId
//go:noescape
func GetMediaDeviceInfoGroupId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDeviceInfo_ToJSON
//go:noescape
func HasMediaDeviceInfoToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDeviceInfo_ToJSON
//go:noescape
func MediaDeviceInfoToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDeviceInfo_ToJSON
//go:noescape
func CallMediaDeviceInfoToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDeviceInfo_ToJSON
//go:noescape
func TryMediaDeviceInfoToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_SelfCapturePreferenceEnum
//go:noescape
func ConstOfSelfCapturePreferenceEnum(str js.Ref) uint32

//go:wasmimport plat/js/web constof_SystemAudioPreferenceEnum
//go:noescape
func ConstOfSystemAudioPreferenceEnum(str js.Ref) uint32

//go:wasmimport plat/js/web constof_SurfaceSwitchingPreferenceEnum
//go:noescape
func ConstOfSurfaceSwitchingPreferenceEnum(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MonitorTypeSurfacesEnum
//go:noescape
func ConstOfMonitorTypeSurfacesEnum(str js.Ref) uint32

//go:wasmimport plat/js/web store_DisplayMediaStreamOptions
//go:noescape
func DisplayMediaStreamOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DisplayMediaStreamOptions
//go:noescape
func DisplayMediaStreamOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ViewportMediaStreamConstraints
//go:noescape
func ViewportMediaStreamConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ViewportMediaStreamConstraints
//go:noescape
func ViewportMediaStreamConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaTrackSupportedConstraints
//go:noescape
func MediaTrackSupportedConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaTrackSupportedConstraints
//go:noescape
func MediaTrackSupportedConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_MediaDevices_EnumerateDevices
//go:noescape
func HasMediaDevicesEnumerateDevices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_EnumerateDevices
//go:noescape
func MediaDevicesEnumerateDevicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_EnumerateDevices
//go:noescape
func CallMediaDevicesEnumerateDevices(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_EnumerateDevices
//go:noescape
func TryMediaDevicesEnumerateDevices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetDisplayMedia
//go:noescape
func HasMediaDevicesGetDisplayMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetDisplayMedia
//go:noescape
func MediaDevicesGetDisplayMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetDisplayMedia
//go:noescape
func CallMediaDevicesGetDisplayMedia(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetDisplayMedia
//go:noescape
func TryMediaDevicesGetDisplayMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetDisplayMedia1
//go:noescape
func HasMediaDevicesGetDisplayMedia1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetDisplayMedia1
//go:noescape
func MediaDevicesGetDisplayMedia1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetDisplayMedia1
//go:noescape
func CallMediaDevicesGetDisplayMedia1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetDisplayMedia1
//go:noescape
func TryMediaDevicesGetDisplayMedia1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_SetSupportedCaptureActions
//go:noescape
func HasMediaDevicesSetSupportedCaptureActions(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetSupportedCaptureActions
//go:noescape
func MediaDevicesSetSupportedCaptureActionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetSupportedCaptureActions
//go:noescape
func CallMediaDevicesSetSupportedCaptureActions(
	this js.Ref, retPtr unsafe.Pointer,
	actions js.Ref)

//go:wasmimport plat/js/web try_MediaDevices_SetSupportedCaptureActions
//go:noescape
func TryMediaDevicesSetSupportedCaptureActions(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	actions js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_SelectAudioOutput
//go:noescape
func HasMediaDevicesSelectAudioOutput(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SelectAudioOutput
//go:noescape
func MediaDevicesSelectAudioOutputFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SelectAudioOutput
//go:noescape
func CallMediaDevicesSelectAudioOutput(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_SelectAudioOutput
//go:noescape
func TryMediaDevicesSelectAudioOutput(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_SelectAudioOutput1
//go:noescape
func HasMediaDevicesSelectAudioOutput1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SelectAudioOutput1
//go:noescape
func MediaDevicesSelectAudioOutput1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SelectAudioOutput1
//go:noescape
func CallMediaDevicesSelectAudioOutput1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_SelectAudioOutput1
//go:noescape
func TryMediaDevicesSelectAudioOutput1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_SetCaptureHandleConfig
//go:noescape
func HasMediaDevicesSetCaptureHandleConfig(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetCaptureHandleConfig
//go:noescape
func MediaDevicesSetCaptureHandleConfigFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetCaptureHandleConfig
//go:noescape
func CallMediaDevicesSetCaptureHandleConfig(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_SetCaptureHandleConfig
//go:noescape
func TryMediaDevicesSetCaptureHandleConfig(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_SetCaptureHandleConfig1
//go:noescape
func HasMediaDevicesSetCaptureHandleConfig1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetCaptureHandleConfig1
//go:noescape
func MediaDevicesSetCaptureHandleConfig1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetCaptureHandleConfig1
//go:noescape
func CallMediaDevicesSetCaptureHandleConfig1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_SetCaptureHandleConfig1
//go:noescape
func TryMediaDevicesSetCaptureHandleConfig1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetViewportMedia
//go:noescape
func HasMediaDevicesGetViewportMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetViewportMedia
//go:noescape
func MediaDevicesGetViewportMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetViewportMedia
//go:noescape
func CallMediaDevicesGetViewportMedia(
	this js.Ref, retPtr unsafe.Pointer,
	constraints unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetViewportMedia
//go:noescape
func TryMediaDevicesGetViewportMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constraints unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetViewportMedia1
//go:noescape
func HasMediaDevicesGetViewportMedia1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetViewportMedia1
//go:noescape
func MediaDevicesGetViewportMedia1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetViewportMedia1
//go:noescape
func CallMediaDevicesGetViewportMedia1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetViewportMedia1
//go:noescape
func TryMediaDevicesGetViewportMedia1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetSupportedConstraints
//go:noescape
func HasMediaDevicesGetSupportedConstraints(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetSupportedConstraints
//go:noescape
func MediaDevicesGetSupportedConstraintsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetSupportedConstraints
//go:noescape
func CallMediaDevicesGetSupportedConstraints(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetSupportedConstraints
//go:noescape
func TryMediaDevicesGetSupportedConstraints(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetUserMedia
//go:noescape
func HasMediaDevicesGetUserMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetUserMedia
//go:noescape
func MediaDevicesGetUserMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetUserMedia
//go:noescape
func CallMediaDevicesGetUserMedia(
	this js.Ref, retPtr unsafe.Pointer,
	constraints unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetUserMedia
//go:noescape
func TryMediaDevicesGetUserMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constraints unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaDevices_GetUserMedia1
//go:noescape
func HasMediaDevicesGetUserMedia1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetUserMedia1
//go:noescape
func MediaDevicesGetUserMedia1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetUserMedia1
//go:noescape
func CallMediaDevicesGetUserMedia1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaDevices_GetUserMedia1
//go:noescape
func TryMediaDevicesGetUserMedia1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_SerialPortInfo
//go:noescape
func SerialPortInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialPortInfo
//go:noescape
func SerialPortInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ParityType
//go:noescape
func ConstOfParityType(str js.Ref) uint32

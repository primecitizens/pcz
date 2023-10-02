// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorker_State
//go:noescape

func GetServiceWorkerState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage
//go:noescape

func CallServiceWorkerPostMessage(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
	transfer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage
//go:noescape

func ServiceWorkerPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage1
//go:noescape

func CallServiceWorkerPostMessage1(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage1
//go:noescape

func ServiceWorkerPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorker_PostMessage2
//go:noescape

func CallServiceWorkerPostMessage2(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorker_PostMessage2
//go:noescape

func ServiceWorkerPostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigationPreloadState
//go:noescape

func NavigationPreloadStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationPreloadState
//go:noescape

func NavigationPreloadStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_Enable
//go:noescape

func CallNavigationPreloadManagerEnable(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_Enable
//go:noescape

func NavigationPreloadManagerEnableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_Disable
//go:noescape

func CallNavigationPreloadManagerDisable(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_Disable
//go:noescape

func NavigationPreloadManagerDisableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_SetHeaderValue
//go:noescape

func CallNavigationPreloadManagerSetHeaderValue(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_SetHeaderValue
//go:noescape

func NavigationPreloadManagerSetHeaderValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationPreloadManager_GetState
//go:noescape

func CallNavigationPreloadManagerGetState(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NavigationPreloadManager_GetState
//go:noescape

func NavigationPreloadManagerGetStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ServiceWorkerUpdateViaCache
//go:noescape

func ConstOfServiceWorkerUpdateViaCache(str js.Ref) uint32

//go:wasmimport plat/js/web call_SyncManager_Register
//go:noescape

func CallSyncManagerRegister(
	this js.Ref,
	ptr unsafe.Pointer,

	tag js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SyncManager_Register
//go:noescape

func SyncManagerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SyncManager_GetTags
//go:noescape

func CallSyncManagerGetTags(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SyncManager_GetTags
//go:noescape

func SyncManagerGetTagsFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PushSubscriptionOptions_ApplicationServerKey
//go:noescape

func GetPushSubscriptionOptionsApplicationServerKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PushSubscription_Endpoint
//go:noescape

func GetPushSubscriptionEndpoint(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PushSubscription_ExpirationTime
//go:noescape

func GetPushSubscriptionExpirationTime(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_PushSubscription_Options
//go:noescape

func GetPushSubscriptionOptions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_GetKey
//go:noescape

func CallPushSubscriptionGetKey(
	this js.Ref,
	ptr unsafe.Pointer,

	name uint32,
) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_GetKey
//go:noescape

func PushSubscriptionGetKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_Unsubscribe
//go:noescape

func CallPushSubscriptionUnsubscribe(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_Unsubscribe
//go:noescape

func PushSubscriptionUnsubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushSubscription_ToJSON
//go:noescape

func CallPushSubscriptionToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PushSubscription_ToJSON
//go:noescape

func PushSubscriptionToJSONFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_PushManager_Subscribe
//go:noescape

func CallPushManagerSubscribe(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_PushManager_Subscribe
//go:noescape

func PushManagerSubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_Subscribe1
//go:noescape

func CallPushManagerSubscribe1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PushManager_Subscribe1
//go:noescape

func PushManagerSubscribe1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_GetSubscription
//go:noescape

func CallPushManagerGetSubscription(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PushManager_GetSubscription
//go:noescape

func PushManagerGetSubscriptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_PermissionState
//go:noescape

func CallPushManagerPermissionState(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_PushManager_PermissionState
//go:noescape

func PushManagerPermissionStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PushManager_PermissionState1
//go:noescape

func CallPushManagerPermissionState1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PushManager_PermissionState1
//go:noescape

func PushManagerPermissionState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_PaymentDelegation
//go:noescape

func ConstOfPaymentDelegation(str js.Ref) uint32

//go:wasmimport plat/js/web get_PaymentManager_UserHint
//go:noescape

func GetPaymentManagerUserHint(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_PaymentManager_UserHint
//go:noescape

func SetPaymentManagerUserHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_PaymentManager_EnableDelegations
//go:noescape

func CallPaymentManagerEnableDelegations(
	this js.Ref,
	ptr unsafe.Pointer,

	delegations js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PaymentManager_EnableDelegations
//go:noescape

func PaymentManagerEnableDelegationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Register
//go:noescape

func CallPeriodicSyncManagerRegister(
	this js.Ref,
	ptr unsafe.Pointer,

	tag js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Register
//go:noescape

func PeriodicSyncManagerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Register1
//go:noescape

func CallPeriodicSyncManagerRegister1(
	this js.Ref,
	ptr unsafe.Pointer,

	tag js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Register1
//go:noescape

func PeriodicSyncManagerRegister1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_GetTags
//go:noescape

func CallPeriodicSyncManagerGetTags(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_GetTags
//go:noescape

func PeriodicSyncManagerGetTagsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PeriodicSyncManager_Unregister
//go:noescape

func CallPeriodicSyncManagerUnregister(
	this js.Ref,
	ptr unsafe.Pointer,

	tag js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PeriodicSyncManager_Unregister
//go:noescape

func PeriodicSyncManagerUnregisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_CookieStoreGetOptions
//go:noescape

func CookieStoreGetOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CookieStoreGetOptions
//go:noescape

func CookieStoreGetOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_Subscribe
//go:noescape

func CallCookieStoreManagerSubscribe(
	this js.Ref,
	ptr unsafe.Pointer,

	subscriptions js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_Subscribe
//go:noescape

func CookieStoreManagerSubscribeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_GetSubscriptions
//go:noescape

func CallCookieStoreManagerGetSubscriptions(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_GetSubscriptions
//go:noescape

func CookieStoreManagerGetSubscriptionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CookieStoreManager_Unsubscribe
//go:noescape

func CallCookieStoreManagerUnsubscribe(
	this js.Ref,
	ptr unsafe.Pointer,

	subscriptions js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CookieStoreManager_Unsubscribe
//go:noescape

func CookieStoreManagerUnsubscribeFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_ContentIndex_Add
//go:noescape

func CallContentIndexAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	description unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_Add
//go:noescape

func ContentIndexAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContentIndex_Delete
//go:noescape

func CallContentIndexDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	id js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_Delete
//go:noescape

func ContentIndexDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContentIndex_GetAll
//go:noescape

func CallContentIndexGetAll(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ContentIndex_GetAll
//go:noescape

func ContentIndexGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Installing
//go:noescape

func GetServiceWorkerRegistrationInstalling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Waiting
//go:noescape

func GetServiceWorkerRegistrationWaiting(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Active
//go:noescape

func GetServiceWorkerRegistrationActive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_NavigationPreload
//go:noescape

func GetServiceWorkerRegistrationNavigationPreload(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Scope
//go:noescape

func GetServiceWorkerRegistrationScope(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_UpdateViaCache
//go:noescape

func GetServiceWorkerRegistrationUpdateViaCache(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Sync
//go:noescape

func GetServiceWorkerRegistrationSync(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_BackgroundFetch
//go:noescape

func GetServiceWorkerRegistrationBackgroundFetch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PushManager
//go:noescape

func GetServiceWorkerRegistrationPushManager(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PaymentManager
//go:noescape

func GetServiceWorkerRegistrationPaymentManager(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_PeriodicSync
//go:noescape

func GetServiceWorkerRegistrationPeriodicSync(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Cookies
//go:noescape

func GetServiceWorkerRegistrationCookies(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerRegistration_Index
//go:noescape

func GetServiceWorkerRegistrationIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_Update
//go:noescape

func CallServiceWorkerRegistrationUpdate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_Update
//go:noescape

func ServiceWorkerRegistrationUpdateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_Unregister
//go:noescape

func CallServiceWorkerRegistrationUnregister(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_Unregister
//go:noescape

func ServiceWorkerRegistrationUnregisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_ShowNotification
//go:noescape

func CallServiceWorkerRegistrationShowNotification(
	this js.Ref,
	ptr unsafe.Pointer,

	title js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_ShowNotification
//go:noescape

func ServiceWorkerRegistrationShowNotificationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_ShowNotification1
//go:noescape

func CallServiceWorkerRegistrationShowNotification1(
	this js.Ref,
	ptr unsafe.Pointer,

	title js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_ShowNotification1
//go:noescape

func ServiceWorkerRegistrationShowNotification1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_GetNotifications
//go:noescape

func CallServiceWorkerRegistrationGetNotifications(
	this js.Ref,
	ptr unsafe.Pointer,

	filter unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_GetNotifications
//go:noescape

func ServiceWorkerRegistrationGetNotificationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerRegistration_GetNotifications1
//go:noescape

func CallServiceWorkerRegistrationGetNotifications1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerRegistration_GetNotifications1
//go:noescape

func ServiceWorkerRegistrationGetNotifications1Func(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ServiceWorkerContainer_Ready
//go:noescape

func GetServiceWorkerContainerReady(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_Register
//go:noescape

func CallServiceWorkerContainerRegister(
	this js.Ref,
	ptr unsafe.Pointer,

	scriptURL js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_Register
//go:noescape

func ServiceWorkerContainerRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_Register1
//go:noescape

func CallServiceWorkerContainerRegister1(
	this js.Ref,
	ptr unsafe.Pointer,

	scriptURL js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_Register1
//go:noescape

func ServiceWorkerContainerRegister1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistration
//go:noescape

func CallServiceWorkerContainerGetRegistration(
	this js.Ref,
	ptr unsafe.Pointer,

	clientURL js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistration
//go:noescape

func ServiceWorkerContainerGetRegistrationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistration1
//go:noescape

func CallServiceWorkerContainerGetRegistration1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistration1
//go:noescape

func ServiceWorkerContainerGetRegistration1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_GetRegistrations
//go:noescape

func CallServiceWorkerContainerGetRegistrations(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_GetRegistrations
//go:noescape

func ServiceWorkerContainerGetRegistrationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ServiceWorkerContainer_StartMessages
//go:noescape

func CallServiceWorkerContainerStartMessages(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ServiceWorkerContainer_StartMessages
//go:noescape

func ServiceWorkerContainerStartMessagesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaDeviceKind
//go:noescape

func ConstOfMediaDeviceKind(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaDeviceInfo_DeviceId
//go:noescape

func GetMediaDeviceInfoDeviceId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaDeviceInfo_Kind
//go:noescape

func GetMediaDeviceInfoKind(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaDeviceInfo_Label
//go:noescape

func GetMediaDeviceInfoLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaDeviceInfo_GroupId
//go:noescape

func GetMediaDeviceInfoGroupId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MediaDeviceInfo_ToJSON
//go:noescape

func CallMediaDeviceInfoToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDeviceInfo_ToJSON
//go:noescape

func MediaDeviceInfoToJSONFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_MediaDevices_EnumerateDevices
//go:noescape

func CallMediaDevicesEnumerateDevices(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_EnumerateDevices
//go:noescape

func MediaDevicesEnumerateDevicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetDisplayMedia
//go:noescape

func CallMediaDevicesGetDisplayMedia(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetDisplayMedia
//go:noescape

func MediaDevicesGetDisplayMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetDisplayMedia1
//go:noescape

func CallMediaDevicesGetDisplayMedia1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetDisplayMedia1
//go:noescape

func MediaDevicesGetDisplayMedia1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetSupportedCaptureActions
//go:noescape

func CallMediaDevicesSetSupportedCaptureActions(
	this js.Ref,
	ptr unsafe.Pointer,

	actions js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetSupportedCaptureActions
//go:noescape

func MediaDevicesSetSupportedCaptureActionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SelectAudioOutput
//go:noescape

func CallMediaDevicesSelectAudioOutput(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SelectAudioOutput
//go:noescape

func MediaDevicesSelectAudioOutputFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SelectAudioOutput1
//go:noescape

func CallMediaDevicesSelectAudioOutput1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SelectAudioOutput1
//go:noescape

func MediaDevicesSelectAudioOutput1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetCaptureHandleConfig
//go:noescape

func CallMediaDevicesSetCaptureHandleConfig(
	this js.Ref,
	ptr unsafe.Pointer,

	config unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetCaptureHandleConfig
//go:noescape

func MediaDevicesSetCaptureHandleConfigFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_SetCaptureHandleConfig1
//go:noescape

func CallMediaDevicesSetCaptureHandleConfig1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_SetCaptureHandleConfig1
//go:noescape

func MediaDevicesSetCaptureHandleConfig1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetViewportMedia
//go:noescape

func CallMediaDevicesGetViewportMedia(
	this js.Ref,
	ptr unsafe.Pointer,

	constraints unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetViewportMedia
//go:noescape

func MediaDevicesGetViewportMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetViewportMedia1
//go:noescape

func CallMediaDevicesGetViewportMedia1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetViewportMedia1
//go:noescape

func MediaDevicesGetViewportMedia1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetSupportedConstraints
//go:noescape

func CallMediaDevicesGetSupportedConstraints(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetSupportedConstraints
//go:noescape

func MediaDevicesGetSupportedConstraintsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetUserMedia
//go:noescape

func CallMediaDevicesGetUserMedia(
	this js.Ref,
	ptr unsafe.Pointer,

	constraints unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetUserMedia
//go:noescape

func MediaDevicesGetUserMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaDevices_GetUserMedia1
//go:noescape

func CallMediaDevicesGetUserMedia1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaDevices_GetUserMedia1
//go:noescape

func MediaDevicesGetUserMedia1Func(this js.Ref) js.Ref

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

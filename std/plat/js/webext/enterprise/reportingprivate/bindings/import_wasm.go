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

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_AntiVirusProductState
//go:noescape
func ConstOfAntiVirusProductState(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_AntiVirusSignal
//go:noescape
func AntiVirusSignalJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_AntiVirusSignal
//go:noescape
func AntiVirusSignalJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_CertificateStatus
//go:noescape
func ConstOfCertificateStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_Certificate
//go:noescape
func CertificateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_Certificate
//go:noescape
func CertificateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_RealtimeUrlCheckMode
//go:noescape
func ConstOfRealtimeUrlCheckMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_SafeBrowsingLevel
//go:noescape
func ConstOfSafeBrowsingLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_PasswordProtectionTrigger
//go:noescape
func ConstOfPasswordProtectionTrigger(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_SettingValue
//go:noescape
func ConstOfSettingValue(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_ContextInfo
//go:noescape
func ContextInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_ContextInfo
//go:noescape
func ContextInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_DeviceInfo
//go:noescape
func DeviceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_DeviceInfo
//go:noescape
func DeviceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_EventType
//go:noescape
func ConstOfEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_EnqueueRecordRequest
//go:noescape
func EnqueueRecordRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_EnqueueRecordRequest
//go:noescape
func EnqueueRecordRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_PresenceValue
//go:noescape
func ConstOfPresenceValue(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetFileSystemInfoResponse
//go:noescape
func GetFileSystemInfoResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetFileSystemInfoResponse
//go:noescape
func GetFileSystemInfoResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetFileSystemInfoOptions
//go:noescape
func GetFileSystemInfoOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetFileSystemInfoOptions
//go:noescape
func GetFileSystemInfoOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_UserContext
//go:noescape
func UserContextJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_UserContext
//go:noescape
func UserContextJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetFileSystemInfoRequest
//go:noescape
func GetFileSystemInfoRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetFileSystemInfoRequest
//go:noescape
func GetFileSystemInfoRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate constof_RegistryHive
//go:noescape
func ConstOfRegistryHive(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetSettingsOptions
//go:noescape
func GetSettingsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetSettingsOptions
//go:noescape
func GetSettingsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetSettingsRequest
//go:noescape
func GetSettingsRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetSettingsRequest
//go:noescape
func GetSettingsRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_GetSettingsResponse
//go:noescape
func GetSettingsResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_GetSettingsResponse
//go:noescape
func GetSettingsResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate store_HotfixSignal
//go:noescape
func HotfixSignalJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate load_HotfixSignal
//go:noescape
func HotfixSignalJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_EnqueueRecord
//go:noescape
func HasFuncEnqueueRecord() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_EnqueueRecord
//go:noescape
func FuncEnqueueRecord(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_EnqueueRecord
//go:noescape
func CallEnqueueRecord(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_EnqueueRecord
//go:noescape
func TryEnqueueRecord(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetAvInfo
//go:noescape
func HasFuncGetAvInfo() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetAvInfo
//go:noescape
func FuncGetAvInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetAvInfo
//go:noescape
func CallGetAvInfo(
	retPtr unsafe.Pointer,
	userContext unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetAvInfo
//go:noescape
func TryGetAvInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	userContext unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetCertificate
//go:noescape
func HasFuncGetCertificate() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetCertificate
//go:noescape
func FuncGetCertificate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetCertificate
//go:noescape
func CallGetCertificate(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetCertificate
//go:noescape
func TryGetCertificate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetContextInfo
//go:noescape
func HasFuncGetContextInfo() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetContextInfo
//go:noescape
func FuncGetContextInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetContextInfo
//go:noescape
func CallGetContextInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetContextInfo
//go:noescape
func TryGetContextInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetDeviceData
//go:noescape
func HasFuncGetDeviceData() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetDeviceData
//go:noescape
func FuncGetDeviceData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetDeviceData
//go:noescape
func CallGetDeviceData(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetDeviceData
//go:noescape
func TryGetDeviceData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetDeviceId
//go:noescape
func HasFuncGetDeviceId() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetDeviceId
//go:noescape
func FuncGetDeviceId(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetDeviceId
//go:noescape
func CallGetDeviceId(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetDeviceId
//go:noescape
func TryGetDeviceId(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetDeviceInfo
//go:noescape
func HasFuncGetDeviceInfo() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetDeviceInfo
//go:noescape
func FuncGetDeviceInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetDeviceInfo
//go:noescape
func CallGetDeviceInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetDeviceInfo
//go:noescape
func TryGetDeviceInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetFileSystemInfo
//go:noescape
func HasFuncGetFileSystemInfo() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetFileSystemInfo
//go:noescape
func FuncGetFileSystemInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetFileSystemInfo
//go:noescape
func CallGetFileSystemInfo(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetFileSystemInfo
//go:noescape
func TryGetFileSystemInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetHotfixes
//go:noescape
func HasFuncGetHotfixes() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetHotfixes
//go:noescape
func FuncGetHotfixes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetHotfixes
//go:noescape
func CallGetHotfixes(
	retPtr unsafe.Pointer,
	userContext unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetHotfixes
//go:noescape
func TryGetHotfixes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	userContext unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetPersistentSecret
//go:noescape
func HasFuncGetPersistentSecret() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetPersistentSecret
//go:noescape
func FuncGetPersistentSecret(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetPersistentSecret
//go:noescape
func CallGetPersistentSecret(
	retPtr unsafe.Pointer,
	resetSecret js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetPersistentSecret
//go:noescape
func TryGetPersistentSecret(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	resetSecret js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_GetSettings
//go:noescape
func HasFuncGetSettings() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_GetSettings
//go:noescape
func FuncGetSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_GetSettings
//go:noescape
func CallGetSettings(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_GetSettings
//go:noescape
func TryGetSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate has_SetDeviceData
//go:noescape
func HasFuncSetDeviceData() js.Ref

//go:wasmimport plat/js/webext/enterprise/reportingprivate func_SetDeviceData
//go:noescape
func FuncSetDeviceData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/reportingprivate call_SetDeviceData
//go:noescape
func CallSetDeviceData(
	retPtr unsafe.Pointer,
	id js.Ref,
	data js.Ref)

//go:wasmimport plat/js/webext/enterprise/reportingprivate try_SetDeviceData
//go:noescape
func TrySetDeviceData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	data js.Ref) (ok js.Ref)

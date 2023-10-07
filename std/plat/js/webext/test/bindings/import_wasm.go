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

//go:wasmimport plat/js/webext/test store_AssertPromiseRejectsArgExpectedMessageChoice1
//go:noescape
func AssertPromiseRejectsArgExpectedMessageChoice1JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_AssertPromiseRejectsArgExpectedMessageChoice1
//go:noescape
func AssertPromiseRejectsArgExpectedMessageChoice1JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_AssertPromiseRejectsArgPromise
//go:noescape
func AssertPromiseRejectsArgPromiseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_AssertPromiseRejectsArgPromise
//go:noescape
func AssertPromiseRejectsArgPromiseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_AssertPromiseRejectsReturnType
//go:noescape
func AssertPromiseRejectsReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_AssertPromiseRejectsReturnType
//go:noescape
func AssertPromiseRejectsReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_AssertThrowsArgMessageChoice1
//go:noescape
func AssertThrowsArgMessageChoice1JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_AssertThrowsArgMessageChoice1
//go:noescape
func AssertThrowsArgMessageChoice1JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_GetConfigReturnTypeFieldFtpServer
//go:noescape
func GetConfigReturnTypeFieldFtpServerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_GetConfigReturnTypeFieldFtpServer
//go:noescape
func GetConfigReturnTypeFieldFtpServerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_GetConfigReturnTypeFieldLoginStatus
//go:noescape
func GetConfigReturnTypeFieldLoginStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_GetConfigReturnTypeFieldLoginStatus
//go:noescape
func GetConfigReturnTypeFieldLoginStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_GetConfigReturnTypeFieldTestServer
//go:noescape
func GetConfigReturnTypeFieldTestServerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_GetConfigReturnTypeFieldTestServer
//go:noescape
func GetConfigReturnTypeFieldTestServerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_GetConfigReturnType
//go:noescape
func GetConfigReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_GetConfigReturnType
//go:noescape
func GetConfigReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_LoadScriptReturnType
//go:noescape
func LoadScriptReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_LoadScriptReturnType
//go:noescape
func LoadScriptReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test store_OnMessageArgInfo
//go:noescape
func OnMessageArgInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/test load_OnMessageArgInfo
//go:noescape
func OnMessageArgInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/test has_AssertEq
//go:noescape
func HasFuncAssertEq() js.Ref

//go:wasmimport plat/js/webext/test func_AssertEq
//go:noescape
func FuncAssertEq(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertEq
//go:noescape
func CallAssertEq(
	retPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_AssertEq
//go:noescape
func TryAssertEq(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertFalse
//go:noescape
func HasFuncAssertFalse() js.Ref

//go:wasmimport plat/js/webext/test func_AssertFalse
//go:noescape
func FuncAssertFalse(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertFalse
//go:noescape
func CallAssertFalse(
	retPtr unsafe.Pointer,
	test js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_AssertFalse
//go:noescape
func TryAssertFalse(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	test js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertLastError
//go:noescape
func HasFuncAssertLastError() js.Ref

//go:wasmimport plat/js/webext/test func_AssertLastError
//go:noescape
func FuncAssertLastError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertLastError
//go:noescape
func CallAssertLastError(
	retPtr unsafe.Pointer,
	expectedError js.Ref)

//go:wasmimport plat/js/webext/test try_AssertLastError
//go:noescape
func TryAssertLastError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expectedError js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertNe
//go:noescape
func HasFuncAssertNe() js.Ref

//go:wasmimport plat/js/webext/test func_AssertNe
//go:noescape
func FuncAssertNe(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertNe
//go:noescape
func CallAssertNe(
	retPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_AssertNe
//go:noescape
func TryAssertNe(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertNoLastError
//go:noescape
func HasFuncAssertNoLastError() js.Ref

//go:wasmimport plat/js/webext/test func_AssertNoLastError
//go:noescape
func FuncAssertNoLastError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertNoLastError
//go:noescape
func CallAssertNoLastError(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_AssertNoLastError
//go:noescape
func TryAssertNoLastError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertPromiseRejects
//go:noescape
func HasFuncAssertPromiseRejects() js.Ref

//go:wasmimport plat/js/webext/test func_AssertPromiseRejects
//go:noescape
func FuncAssertPromiseRejects(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertPromiseRejects
//go:noescape
func CallAssertPromiseRejects(
	retPtr unsafe.Pointer,
	promise unsafe.Pointer,
	expectedMessage js.Ref)

//go:wasmimport plat/js/webext/test try_AssertPromiseRejects
//go:noescape
func TryAssertPromiseRejects(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	promise unsafe.Pointer,
	expectedMessage js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertThrows
//go:noescape
func HasFuncAssertThrows() js.Ref

//go:wasmimport plat/js/webext/test func_AssertThrows
//go:noescape
func FuncAssertThrows(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertThrows
//go:noescape
func CallAssertThrows(
	retPtr unsafe.Pointer,
	fn js.Ref,
	self js.Ref,
	args js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_AssertThrows
//go:noescape
func TryAssertThrows(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fn js.Ref,
	self js.Ref,
	args js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_AssertTrue
//go:noescape
func HasFuncAssertTrue() js.Ref

//go:wasmimport plat/js/webext/test func_AssertTrue
//go:noescape
func FuncAssertTrue(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_AssertTrue
//go:noescape
func CallAssertTrue(
	retPtr unsafe.Pointer,
	test js.Ref,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_AssertTrue
//go:noescape
func TryAssertTrue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	test js.Ref,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_Callback
//go:noescape
func HasFuncCallback() js.Ref

//go:wasmimport plat/js/webext/test func_Callback
//go:noescape
func FuncCallback(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_Callback
//go:noescape
func CallCallback(
	retPtr unsafe.Pointer,
	fn js.Ref,
	expectedError js.Ref)

//go:wasmimport plat/js/webext/test try_Callback
//go:noescape
func TryCallback(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fn js.Ref,
	expectedError js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_CallbackAdded
//go:noescape
func HasFuncCallbackAdded() js.Ref

//go:wasmimport plat/js/webext/test func_CallbackAdded
//go:noescape
func FuncCallbackAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_CallbackAdded
//go:noescape
func CallCallbackAdded(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_CallbackAdded
//go:noescape
func TryCallbackAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_CallbackFail
//go:noescape
func HasFuncCallbackFail() js.Ref

//go:wasmimport plat/js/webext/test func_CallbackFail
//go:noescape
func FuncCallbackFail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_CallbackFail
//go:noescape
func CallCallbackFail(
	retPtr unsafe.Pointer,
	expectedError js.Ref,
	fn js.Ref)

//go:wasmimport plat/js/webext/test try_CallbackFail
//go:noescape
func TryCallbackFail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expectedError js.Ref,
	fn js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_CallbackPass
//go:noescape
func HasFuncCallbackPass() js.Ref

//go:wasmimport plat/js/webext/test func_CallbackPass
//go:noescape
func FuncCallbackPass(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_CallbackPass
//go:noescape
func CallCallbackPass(
	retPtr unsafe.Pointer,
	fn js.Ref)

//go:wasmimport plat/js/webext/test try_CallbackPass
//go:noescape
func TryCallbackPass(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fn js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_CheckDeepEq
//go:noescape
func HasFuncCheckDeepEq() js.Ref

//go:wasmimport plat/js/webext/test func_CheckDeepEq
//go:noescape
func FuncCheckDeepEq(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_CheckDeepEq
//go:noescape
func CallCheckDeepEq(
	retPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref)

//go:wasmimport plat/js/webext/test try_CheckDeepEq
//go:noescape
func TryCheckDeepEq(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expected js.Ref,
	actual js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_Fail
//go:noescape
func HasFuncFail() js.Ref

//go:wasmimport plat/js/webext/test func_Fail
//go:noescape
func FuncFail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_Fail
//go:noescape
func CallFail(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_Fail
//go:noescape
func TryFail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_GetApiDefinitions
//go:noescape
func HasFuncGetApiDefinitions() js.Ref

//go:wasmimport plat/js/webext/test func_GetApiDefinitions
//go:noescape
func FuncGetApiDefinitions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_GetApiDefinitions
//go:noescape
func CallGetApiDefinitions(
	retPtr unsafe.Pointer,
	apiNames js.Ref)

//go:wasmimport plat/js/webext/test try_GetApiDefinitions
//go:noescape
func TryGetApiDefinitions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	apiNames js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_GetApiFeatures
//go:noescape
func HasFuncGetApiFeatures() js.Ref

//go:wasmimport plat/js/webext/test func_GetApiFeatures
//go:noescape
func FuncGetApiFeatures(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_GetApiFeatures
//go:noescape
func CallGetApiFeatures(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_GetApiFeatures
//go:noescape
func TryGetApiFeatures(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_GetConfig
//go:noescape
func HasFuncGetConfig() js.Ref

//go:wasmimport plat/js/webext/test func_GetConfig
//go:noescape
func FuncGetConfig(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_GetConfig
//go:noescape
func CallGetConfig(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_GetConfig
//go:noescape
func TryGetConfig(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_GetModuleSystem
//go:noescape
func HasFuncGetModuleSystem() js.Ref

//go:wasmimport plat/js/webext/test func_GetModuleSystem
//go:noescape
func FuncGetModuleSystem(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_GetModuleSystem
//go:noescape
func CallGetModuleSystem(
	retPtr unsafe.Pointer,
	context js.Ref)

//go:wasmimport plat/js/webext/test try_GetModuleSystem
//go:noescape
func TryGetModuleSystem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	context js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_GetWakeEventPage
//go:noescape
func HasFuncGetWakeEventPage() js.Ref

//go:wasmimport plat/js/webext/test func_GetWakeEventPage
//go:noescape
func FuncGetWakeEventPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_GetWakeEventPage
//go:noescape
func CallGetWakeEventPage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_GetWakeEventPage
//go:noescape
func TryGetWakeEventPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_IsProcessingUserGesture
//go:noescape
func HasFuncIsProcessingUserGesture() js.Ref

//go:wasmimport plat/js/webext/test func_IsProcessingUserGesture
//go:noescape
func FuncIsProcessingUserGesture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_IsProcessingUserGesture
//go:noescape
func CallIsProcessingUserGesture(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/test try_IsProcessingUserGesture
//go:noescape
func TryIsProcessingUserGesture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_ListenForever
//go:noescape
func HasFuncListenForever() js.Ref

//go:wasmimport plat/js/webext/test func_ListenForever
//go:noescape
func FuncListenForever(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_ListenForever
//go:noescape
func CallListenForever(
	retPtr unsafe.Pointer,
	event js.Ref,
	fn js.Ref)

//go:wasmimport plat/js/webext/test try_ListenForever
//go:noescape
func TryListenForever(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	event js.Ref,
	fn js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_ListenOnce
//go:noescape
func HasFuncListenOnce() js.Ref

//go:wasmimport plat/js/webext/test func_ListenOnce
//go:noescape
func FuncListenOnce(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_ListenOnce
//go:noescape
func CallListenOnce(
	retPtr unsafe.Pointer,
	event js.Ref,
	fn js.Ref)

//go:wasmimport plat/js/webext/test try_ListenOnce
//go:noescape
func TryListenOnce(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	event js.Ref,
	fn js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_LoadScript
//go:noescape
func HasFuncLoadScript() js.Ref

//go:wasmimport plat/js/webext/test func_LoadScript
//go:noescape
func FuncLoadScript(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_LoadScript
//go:noescape
func CallLoadScript(
	retPtr unsafe.Pointer,
	scriptUrl js.Ref)

//go:wasmimport plat/js/webext/test try_LoadScript
//go:noescape
func TryLoadScript(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scriptUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_Log
//go:noescape
func HasFuncLog() js.Ref

//go:wasmimport plat/js/webext/test func_Log
//go:noescape
func FuncLog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_Log
//go:noescape
func CallLog(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_Log
//go:noescape
func TryLog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_NotifyFail
//go:noescape
func HasFuncNotifyFail() js.Ref

//go:wasmimport plat/js/webext/test func_NotifyFail
//go:noescape
func FuncNotifyFail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_NotifyFail
//go:noescape
func CallNotifyFail(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_NotifyFail
//go:noescape
func TryNotifyFail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_NotifyPass
//go:noescape
func HasFuncNotifyPass() js.Ref

//go:wasmimport plat/js/webext/test func_NotifyPass
//go:noescape
func FuncNotifyPass(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_NotifyPass
//go:noescape
func CallNotifyPass(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_NotifyPass
//go:noescape
func TryNotifyPass(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_OnMessage
//go:noescape
func HasFuncOnMessage() js.Ref

//go:wasmimport plat/js/webext/test func_OnMessage
//go:noescape
func FuncOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_OnMessage
//go:noescape
func CallOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/test try_OnMessage
//go:noescape
func TryOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_OffMessage
//go:noescape
func HasFuncOffMessage() js.Ref

//go:wasmimport plat/js/webext/test func_OffMessage
//go:noescape
func FuncOffMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_OffMessage
//go:noescape
func CallOffMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/test try_OffMessage
//go:noescape
func TryOffMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_HasOnMessage
//go:noescape
func HasFuncHasOnMessage() js.Ref

//go:wasmimport plat/js/webext/test func_HasOnMessage
//go:noescape
func FuncHasOnMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_HasOnMessage
//go:noescape
func CallHasOnMessage(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/test try_HasOnMessage
//go:noescape
func TryHasOnMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_OpenFileUrl
//go:noescape
func HasFuncOpenFileUrl() js.Ref

//go:wasmimport plat/js/webext/test func_OpenFileUrl
//go:noescape
func FuncOpenFileUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_OpenFileUrl
//go:noescape
func CallOpenFileUrl(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/test try_OpenFileUrl
//go:noescape
func TryOpenFileUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_RunTests
//go:noescape
func HasFuncRunTests() js.Ref

//go:wasmimport plat/js/webext/test func_RunTests
//go:noescape
func FuncRunTests(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_RunTests
//go:noescape
func CallRunTests(
	retPtr unsafe.Pointer,
	tests js.Ref)

//go:wasmimport plat/js/webext/test try_RunTests
//go:noescape
func TryRunTests(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tests js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_RunWithUserGesture
//go:noescape
func HasFuncRunWithUserGesture() js.Ref

//go:wasmimport plat/js/webext/test func_RunWithUserGesture
//go:noescape
func FuncRunWithUserGesture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_RunWithUserGesture
//go:noescape
func CallRunWithUserGesture(
	retPtr unsafe.Pointer,
	functionToRun js.Ref)

//go:wasmimport plat/js/webext/test try_RunWithUserGesture
//go:noescape
func TryRunWithUserGesture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	functionToRun js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_SendMessage
//go:noescape
func HasFuncSendMessage() js.Ref

//go:wasmimport plat/js/webext/test func_SendMessage
//go:noescape
func FuncSendMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_SendMessage
//go:noescape
func CallSendMessage(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_SendMessage
//go:noescape
func TrySendMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_SendScriptResult
//go:noescape
func HasFuncSendScriptResult() js.Ref

//go:wasmimport plat/js/webext/test func_SendScriptResult
//go:noescape
func FuncSendScriptResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_SendScriptResult
//go:noescape
func CallSendScriptResult(
	retPtr unsafe.Pointer,
	result js.Ref)

//go:wasmimport plat/js/webext/test try_SendScriptResult
//go:noescape
func TrySendScriptResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	result js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_SetExceptionHandler
//go:noescape
func HasFuncSetExceptionHandler() js.Ref

//go:wasmimport plat/js/webext/test func_SetExceptionHandler
//go:noescape
func FuncSetExceptionHandler(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_SetExceptionHandler
//go:noescape
func CallSetExceptionHandler(
	retPtr unsafe.Pointer,
	handler js.Ref)

//go:wasmimport plat/js/webext/test try_SetExceptionHandler
//go:noescape
func TrySetExceptionHandler(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_Succeed
//go:noescape
func HasFuncSucceed() js.Ref

//go:wasmimport plat/js/webext/test func_Succeed
//go:noescape
func FuncSucceed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_Succeed
//go:noescape
func CallSucceed(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_Succeed
//go:noescape
func TrySucceed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/test has_WaitForRoundTrip
//go:noescape
func HasFuncWaitForRoundTrip() js.Ref

//go:wasmimport plat/js/webext/test func_WaitForRoundTrip
//go:noescape
func FuncWaitForRoundTrip(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/test call_WaitForRoundTrip
//go:noescape
func CallWaitForRoundTrip(
	retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/webext/test try_WaitForRoundTrip
//go:noescape
func TryWaitForRoundTrip(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

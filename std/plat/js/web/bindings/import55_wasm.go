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

//go:wasmimport plat/js/web store_Transformer
//go:noescape
func TransformerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Transformer
//go:noescape
func TransformerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_TransitionEventInit
//go:noescape
func TransitionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TransitionEventInit
//go:noescape
func TransitionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_TransitionEvent_TransitionEvent
//go:noescape
func NewTransitionEventByTransitionEvent(
	typ js.Ref,
	transitionEventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_TransitionEvent_TransitionEvent1
//go:noescape
func NewTransitionEventByTransitionEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_TransitionEvent_PropertyName
//go:noescape
func GetTransitionEventPropertyName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TransitionEvent_ElapsedTime
//go:noescape
func GetTransitionEventElapsedTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TransitionEvent_PseudoElement
//go:noescape
func GetTransitionEventPseudoElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_UIEventInit
//go:noescape
func UIEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UIEventInit
//go:noescape
func UIEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_UIEvent_UIEvent
//go:noescape
func NewUIEventByUIEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_UIEvent_UIEvent1
//go:noescape
func NewUIEventByUIEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_UIEvent_View
//go:noescape
func GetUIEventView(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UIEvent_Detail
//go:noescape
func GetUIEventDetail(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UIEvent_Which
//go:noescape
func GetUIEventWhich(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UIEvent_SourceCapabilities
//go:noescape
func GetUIEventSourceCapabilities(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_UIEvent_InitUIEvent
//go:noescape
func HasFuncUIEventInitUIEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_UIEvent_InitUIEvent
//go:noescape
func FuncUIEventInitUIEvent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_UIEvent_InitUIEvent
//go:noescape
func CallUIEventInitUIEvent(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32)

//go:wasmimport plat/js/web try_UIEvent_InitUIEvent
//go:noescape
func TryUIEventInitUIEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref,
	detailArg int32) (ok js.Ref)

//go:wasmimport plat/js/web has_UIEvent_InitUIEvent1
//go:noescape
func HasFuncUIEventInitUIEvent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_UIEvent_InitUIEvent1
//go:noescape
func FuncUIEventInitUIEvent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_UIEvent_InitUIEvent1
//go:noescape
func CallUIEventInitUIEvent1(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref)

//go:wasmimport plat/js/web try_UIEvent_InitUIEvent1
//go:noescape
func TryUIEventInitUIEvent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref,
	viewArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_UIEvent_InitUIEvent2
//go:noescape
func HasFuncUIEventInitUIEvent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_UIEvent_InitUIEvent2
//go:noescape
func FuncUIEventInitUIEvent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_UIEvent_InitUIEvent2
//go:noescape
func CallUIEventInitUIEvent2(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref)

//go:wasmimport plat/js/web try_UIEvent_InitUIEvent2
//go:noescape
func TryUIEventInitUIEvent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref,
	cancelableArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_UIEvent_InitUIEvent3
//go:noescape
func HasFuncUIEventInitUIEvent3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_UIEvent_InitUIEvent3
//go:noescape
func FuncUIEventInitUIEvent3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_UIEvent_InitUIEvent3
//go:noescape
func CallUIEventInitUIEvent3(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref)

//go:wasmimport plat/js/web try_UIEvent_InitUIEvent3
//go:noescape
func TryUIEventInitUIEvent3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref,
	bubblesArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_UIEvent_InitUIEvent4
//go:noescape
func HasFuncUIEventInitUIEvent4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_UIEvent_InitUIEvent4
//go:noescape
func FuncUIEventInitUIEvent4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_UIEvent_InitUIEvent4
//go:noescape
func CallUIEventInitUIEvent4(
	this js.Ref, retPtr unsafe.Pointer,
	typeArg js.Ref)

//go:wasmimport plat/js/web try_UIEvent_InitUIEvent4
//go:noescape
func TryUIEventInitUIEvent4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typeArg js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_URL_URL
//go:noescape
func NewURLByURL(
	url js.Ref,
	base js.Ref) js.Ref

//go:wasmimport plat/js/web new_URL_URL1
//go:noescape
func NewURLByURL1(
	url js.Ref) js.Ref

//go:wasmimport plat/js/web get_URL_Href
//go:noescape
func GetURLHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Href
//go:noescape
func SetURLHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Origin
//go:noescape
func GetURLOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URL_Protocol
//go:noescape
func GetURLProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Protocol
//go:noescape
func SetURLProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Username
//go:noescape
func GetURLUsername(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Username
//go:noescape
func SetURLUsername(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Password
//go:noescape
func GetURLPassword(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Password
//go:noescape
func SetURLPassword(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Host
//go:noescape
func GetURLHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Host
//go:noescape
func SetURLHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Hostname
//go:noescape
func GetURLHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Hostname
//go:noescape
func SetURLHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Port
//go:noescape
func GetURLPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Port
//go:noescape
func SetURLPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Pathname
//go:noescape
func GetURLPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Pathname
//go:noescape
func SetURLPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_Search
//go:noescape
func GetURLSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Search
//go:noescape
func SetURLSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_URL_SearchParams
//go:noescape
func GetURLSearchParams(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URL_Hash
//go:noescape
func GetURLHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_URL_Hash
//go:noescape
func SetURLHash(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_URL_CanParse
//go:noescape
func HasFuncURLCanParse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URL_CanParse
//go:noescape
func FuncURLCanParse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URL_CanParse
//go:noescape
func CallURLCanParse(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	base js.Ref)

//go:wasmimport plat/js/web try_URL_CanParse
//go:noescape
func TryURLCanParse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	base js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URL_CanParse1
//go:noescape
func HasFuncURLCanParse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URL_CanParse1
//go:noescape
func FuncURLCanParse1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URL_CanParse1
//go:noescape
func CallURLCanParse1(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_URL_CanParse1
//go:noescape
func TryURLCanParse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URL_ToJSON
//go:noescape
func HasFuncURLToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URL_ToJSON
//go:noescape
func FuncURLToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URL_ToJSON
//go:noescape
func CallURLToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_URL_ToJSON
//go:noescape
func TryURLToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_URL_CreateObjectURL
//go:noescape
func HasFuncURLCreateObjectURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URL_CreateObjectURL
//go:noescape
func FuncURLCreateObjectURL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URL_CreateObjectURL
//go:noescape
func CallURLCreateObjectURL(
	this js.Ref, retPtr unsafe.Pointer,
	obj js.Ref)

//go:wasmimport plat/js/web try_URL_CreateObjectURL
//go:noescape
func TryURLCreateObjectURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	obj js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URL_RevokeObjectURL
//go:noescape
func HasFuncURLRevokeObjectURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URL_RevokeObjectURL
//go:noescape
func FuncURLRevokeObjectURL(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URL_RevokeObjectURL
//go:noescape
func CallURLRevokeObjectURL(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_URL_RevokeObjectURL
//go:noescape
func TryURLRevokeObjectURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_URLPatternInit
//go:noescape
func URLPatternInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_URLPatternInit
//go:noescape
func URLPatternInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_URLPatternOptions
//go:noescape
func URLPatternOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_URLPatternOptions
//go:noescape
func URLPatternOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_URLPatternComponentResult
//go:noescape
func URLPatternComponentResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_URLPatternComponentResult
//go:noescape
func URLPatternComponentResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_URLPatternResult
//go:noescape
func URLPatternResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_URLPatternResult
//go:noescape
func URLPatternResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_URLPattern_URLPattern
//go:noescape
func NewURLPatternByURLPattern(
	input js.Ref,
	baseURL js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_URLPattern_URLPattern1
//go:noescape
func NewURLPatternByURLPattern1(
	input js.Ref,
	baseURL js.Ref) js.Ref

//go:wasmimport plat/js/web new_URLPattern_URLPattern2
//go:noescape
func NewURLPatternByURLPattern2(
	input js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_URLPattern_URLPattern3
//go:noescape
func NewURLPatternByURLPattern3(
	input js.Ref) js.Ref

//go:wasmimport plat/js/web new_URLPattern_URLPattern4
//go:noescape
func NewURLPatternByURLPattern4() js.Ref

//go:wasmimport plat/js/web get_URLPattern_Protocol
//go:noescape
func GetURLPatternProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Username
//go:noescape
func GetURLPatternUsername(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Password
//go:noescape
func GetURLPatternPassword(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Hostname
//go:noescape
func GetURLPatternHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Port
//go:noescape
func GetURLPatternPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Pathname
//go:noescape
func GetURLPatternPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Search
//go:noescape
func GetURLPatternSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_URLPattern_Hash
//go:noescape
func GetURLPatternHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Test
//go:noescape
func HasFuncURLPatternTest(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Test
//go:noescape
func FuncURLPatternTest(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Test
//go:noescape
func CallURLPatternTest(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	baseURL js.Ref)

//go:wasmimport plat/js/web try_URLPattern_Test
//go:noescape
func TryURLPatternTest(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	baseURL js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Test1
//go:noescape
func HasFuncURLPatternTest1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Test1
//go:noescape
func FuncURLPatternTest1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Test1
//go:noescape
func CallURLPatternTest1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_URLPattern_Test1
//go:noescape
func TryURLPatternTest1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Test2
//go:noescape
func HasFuncURLPatternTest2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Test2
//go:noescape
func FuncURLPatternTest2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Test2
//go:noescape
func CallURLPatternTest2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_URLPattern_Test2
//go:noescape
func TryURLPatternTest2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Exec
//go:noescape
func HasFuncURLPatternExec(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Exec
//go:noescape
func FuncURLPatternExec(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Exec
//go:noescape
func CallURLPatternExec(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	baseURL js.Ref)

//go:wasmimport plat/js/web try_URLPattern_Exec
//go:noescape
func TryURLPatternExec(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	baseURL js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Exec1
//go:noescape
func HasFuncURLPatternExec1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Exec1
//go:noescape
func FuncURLPatternExec1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Exec1
//go:noescape
func CallURLPatternExec1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_URLPattern_Exec1
//go:noescape
func TryURLPatternExec1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLPattern_Exec2
//go:noescape
func HasFuncURLPatternExec2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLPattern_Exec2
//go:noescape
func FuncURLPatternExec2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLPattern_Exec2
//go:noescape
func CallURLPatternExec2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_URLPattern_Exec2
//go:noescape
func TryURLPatternExec2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_USBConnectionEventInit
//go:noescape
func USBConnectionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBConnectionEventInit
//go:noescape
func USBConnectionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_USBConnectionEvent_USBConnectionEvent
//go:noescape
func NewUSBConnectionEventByUSBConnectionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_USBConnectionEvent_Device
//go:noescape
func GetUSBConnectionEventDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_USBPermissionDescriptor
//go:noescape
func USBPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBPermissionDescriptor
//go:noescape
func USBPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_USBPermissionResult_Devices
//go:noescape
func GetUSBPermissionResultDevices(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_USBPermissionResult_Devices
//go:noescape
func SetUSBPermissionResultDevices(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_USBPermissionStorage
//go:noescape
func USBPermissionStorageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_USBPermissionStorage
//go:noescape
func USBPermissionStorageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_UncalibratedMagnetometer_UncalibratedMagnetometer
//go:noescape
func NewUncalibratedMagnetometerByUncalibratedMagnetometer(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_UncalibratedMagnetometer_UncalibratedMagnetometer1
//go:noescape
func NewUncalibratedMagnetometerByUncalibratedMagnetometer1() js.Ref

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_X
//go:noescape
func GetUncalibratedMagnetometerX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_Y
//go:noescape
func GetUncalibratedMagnetometerY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_Z
//go:noescape
func GetUncalibratedMagnetometerZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_XBias
//go:noescape
func GetUncalibratedMagnetometerXBias(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_YBias
//go:noescape
func GetUncalibratedMagnetometerYBias(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UncalibratedMagnetometer_ZBias
//go:noescape
func GetUncalibratedMagnetometerZBias(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_UncalibratedMagnetometerReadingValues
//go:noescape
func UncalibratedMagnetometerReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UncalibratedMagnetometerReadingValues
//go:noescape
func UncalibratedMagnetometerReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_WritableStreamDefaultController_Signal
//go:noescape
func GetWritableStreamDefaultControllerSignal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultController_Error
//go:noescape
func HasFuncWritableStreamDefaultControllerError(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultController_Error
//go:noescape
func FuncWritableStreamDefaultControllerError(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultController_Error
//go:noescape
func CallWritableStreamDefaultControllerError(
	this js.Ref, retPtr unsafe.Pointer,
	e js.Ref)

//go:wasmimport plat/js/web try_WritableStreamDefaultController_Error
//go:noescape
func TryWritableStreamDefaultControllerError(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	e js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultController_Error1
//go:noescape
func HasFuncWritableStreamDefaultControllerError1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultController_Error1
//go:noescape
func FuncWritableStreamDefaultControllerError1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultController_Error1
//go:noescape
func CallWritableStreamDefaultControllerError1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStreamDefaultController_Error1
//go:noescape
func TryWritableStreamDefaultControllerError1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_UnderlyingSink
//go:noescape
func UnderlyingSinkJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UnderlyingSink
//go:noescape
func UnderlyingSinkJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_UnderlyingSource
//go:noescape
func UnderlyingSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_UnderlyingSource
//go:noescape
func UnderlyingSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_UserVerificationRequirement
//go:noescape
func ConstOfUserVerificationRequirement(str js.Ref) uint32

//go:wasmimport plat/js/web get_VTTRegion_Id
//go:noescape
func GetVTTRegionId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_Id
//go:noescape
func SetVTTRegionId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_Width
//go:noescape
func GetVTTRegionWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_Width
//go:noescape
func SetVTTRegionWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_Lines
//go:noescape
func GetVTTRegionLines(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_Lines
//go:noescape
func SetVTTRegionLines(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_RegionAnchorX
//go:noescape
func GetVTTRegionRegionAnchorX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_RegionAnchorX
//go:noescape
func SetVTTRegionRegionAnchorX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_RegionAnchorY
//go:noescape
func GetVTTRegionRegionAnchorY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_RegionAnchorY
//go:noescape
func SetVTTRegionRegionAnchorY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_ViewportAnchorX
//go:noescape
func GetVTTRegionViewportAnchorX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_ViewportAnchorX
//go:noescape
func SetVTTRegionViewportAnchorX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_ViewportAnchorY
//go:noescape
func GetVTTRegionViewportAnchorY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_ViewportAnchorY
//go:noescape
func SetVTTRegionViewportAnchorY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTRegion_Scroll
//go:noescape
func GetVTTRegionScroll(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTRegion_Scroll
//go:noescape
func SetVTTRegionScroll(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web new_VTTCue_VTTCue
//go:noescape
func NewVTTCueByVTTCue(
	startTime float64,
	endTime float64,
	text js.Ref) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Region
//go:noescape
func GetVTTCueRegion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Region
//go:noescape
func SetVTTCueRegion(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Vertical
//go:noescape
func GetVTTCueVertical(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Vertical
//go:noescape
func SetVTTCueVertical(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_SnapToLines
//go:noescape
func GetVTTCueSnapToLines(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_SnapToLines
//go:noescape
func SetVTTCueSnapToLines(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Line
//go:noescape
func GetVTTCueLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Line
//go:noescape
func SetVTTCueLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_LineAlign
//go:noescape
func GetVTTCueLineAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_LineAlign
//go:noescape
func SetVTTCueLineAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Position
//go:noescape
func GetVTTCuePosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Position
//go:noescape
func SetVTTCuePosition(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_PositionAlign
//go:noescape
func GetVTTCuePositionAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_PositionAlign
//go:noescape
func SetVTTCuePositionAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Size
//go:noescape
func GetVTTCueSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Size
//go:noescape
func SetVTTCueSize(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Align
//go:noescape
func GetVTTCueAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Align
//go:noescape
func SetVTTCueAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_VTTCue_Text
//go:noescape
func GetVTTCueText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VTTCue_Text
//go:noescape
func SetVTTCueText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_VTTCue_GetCueAsHTML
//go:noescape
func HasFuncVTTCueGetCueAsHTML(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VTTCue_GetCueAsHTML
//go:noescape
func FuncVTTCueGetCueAsHTML(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VTTCue_GetCueAsHTML
//go:noescape
func CallVTTCueGetCueAsHTML(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VTTCue_GetCueAsHTML
//go:noescape
func TryVTTCueGetCueAsHTML(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ValueEventInit
//go:noescape
func ValueEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ValueEventInit
//go:noescape
func ValueEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ValueEvent_ValueEvent
//go:noescape
func NewValueEventByValueEvent(
	typ js.Ref,
	initDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ValueEvent_ValueEvent1
//go:noescape
func NewValueEventByValueEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ValueEvent_Value
//go:noescape
func GetValueEventValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_VideoDecoderInit
//go:noescape
func VideoDecoderInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoDecoderInit
//go:noescape
func VideoDecoderInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoDecoderSupport
//go:noescape
func VideoDecoderSupportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoDecoderSupport
//go:noescape
func VideoDecoderSupportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_VideoDecoder_VideoDecoder
//go:noescape
func NewVideoDecoderByVideoDecoder(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_VideoDecoder_State
//go:noescape
func GetVideoDecoderState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoDecoder_DecodeQueueSize
//go:noescape
func GetVideoDecoderDecodeQueueSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_Configure
//go:noescape
func HasFuncVideoDecoderConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_Configure
//go:noescape
func FuncVideoDecoderConfigure(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_Configure
//go:noescape
func CallVideoDecoderConfigure(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoDecoder_Configure
//go:noescape
func TryVideoDecoderConfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_Decode
//go:noescape
func HasFuncVideoDecoderDecode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_Decode
//go:noescape
func FuncVideoDecoderDecode(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_Decode
//go:noescape
func CallVideoDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer,
	chunk js.Ref)

//go:wasmimport plat/js/web try_VideoDecoder_Decode
//go:noescape
func TryVideoDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	chunk js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_Flush
//go:noescape
func HasFuncVideoDecoderFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_Flush
//go:noescape
func FuncVideoDecoderFlush(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_Flush
//go:noescape
func CallVideoDecoderFlush(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoDecoder_Flush
//go:noescape
func TryVideoDecoderFlush(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_Reset
//go:noescape
func HasFuncVideoDecoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_Reset
//go:noescape
func FuncVideoDecoderReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_Reset
//go:noescape
func CallVideoDecoderReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoDecoder_Reset
//go:noescape
func TryVideoDecoderReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_Close
//go:noescape
func HasFuncVideoDecoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_Close
//go:noescape
func FuncVideoDecoderClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_Close
//go:noescape
func CallVideoDecoderClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoDecoder_Close
//go:noescape
func TryVideoDecoderClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoDecoder_IsConfigSupported
//go:noescape
func HasFuncVideoDecoderIsConfigSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoDecoder_IsConfigSupported
//go:noescape
func FuncVideoDecoderIsConfigSupported(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoDecoder_IsConfigSupported
//go:noescape
func CallVideoDecoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_VideoDecoder_IsConfigSupported
//go:noescape
func TryVideoDecoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_VideoEncoderInit
//go:noescape
func VideoEncoderInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderInit
//go:noescape
func VideoEncoderInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_VideoEncoderBitrateMode
//go:noescape
func ConstOfVideoEncoderBitrateMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_VideoEncoderConfig
//go:noescape
func VideoEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderConfig
//go:noescape
func VideoEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoEncoderEncodeOptionsForHevc
//go:noescape
func VideoEncoderEncodeOptionsForHevcJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderEncodeOptionsForHevc
//go:noescape
func VideoEncoderEncodeOptionsForHevcJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoEncoderEncodeOptionsForVp9
//go:noescape
func VideoEncoderEncodeOptionsForVp9JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoEncoderEncodeOptionsForVp9
//go:noescape
func VideoEncoderEncodeOptionsForVp9JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

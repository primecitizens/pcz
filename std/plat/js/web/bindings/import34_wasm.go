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

//go:wasmimport plat/js/web store_DeviceOrientationEventInit
//go:noescape
func DeviceOrientationEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DeviceOrientationEventInit
//go:noescape
func DeviceOrientationEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DeviceOrientationEvent_DeviceOrientationEvent
//go:noescape
func NewDeviceOrientationEventByDeviceOrientationEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DeviceOrientationEvent_DeviceOrientationEvent1
//go:noescape
func NewDeviceOrientationEventByDeviceOrientationEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_DeviceOrientationEvent_Alpha
//go:noescape
func GetDeviceOrientationEventAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceOrientationEvent_Beta
//go:noescape
func GetDeviceOrientationEventBeta(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceOrientationEvent_Gamma
//go:noescape
func GetDeviceOrientationEventGamma(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DeviceOrientationEvent_Absolute
//go:noescape
func GetDeviceOrientationEventAbsolute(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DeviceOrientationEvent_RequestPermission
//go:noescape
func HasFuncDeviceOrientationEventRequestPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DeviceOrientationEvent_RequestPermission
//go:noescape
func FuncDeviceOrientationEventRequestPermission(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DeviceOrientationEvent_RequestPermission
//go:noescape
func CallDeviceOrientationEventRequestPermission(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DeviceOrientationEvent_RequestPermission
//go:noescape
func TryDeviceOrientationEventRequestPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_DirectionSetting
//go:noescape
func ConstOfDirectionSetting(str js.Ref) uint32

//go:wasmimport plat/js/web constof_DisplayCaptureSurfaceType
//go:noescape
func ConstOfDisplayCaptureSurfaceType(str js.Ref) uint32

//go:wasmimport plat/js/web store_DocumentPictureInPictureEventInit
//go:noescape
func DocumentPictureInPictureEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DocumentPictureInPictureEventInit
//go:noescape
func DocumentPictureInPictureEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DocumentPictureInPictureEvent_DocumentPictureInPictureEvent
//go:noescape
func NewDocumentPictureInPictureEventByDocumentPictureInPictureEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_DocumentPictureInPictureEvent_Window
//go:noescape
func GetDocumentPictureInPictureEventWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_DragEventInit
//go:noescape
func DragEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DragEventInit
//go:noescape
func DragEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DragEvent_DragEvent
//go:noescape
func NewDragEventByDragEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DragEvent_DragEvent1
//go:noescape
func NewDragEventByDragEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_DragEvent_DataTransfer
//go:noescape
func GetDragEventDataTransfer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_CreateQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryCreateQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_CreateQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryCreateQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_CreateQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryCreateQueryEXT(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_CreateQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryCreateQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_DeleteQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryDeleteQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_DeleteQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryDeleteQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_DeleteQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryDeleteQueryEXT(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_DeleteQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryDeleteQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_IsQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryIsQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_IsQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryIsQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_IsQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryIsQueryEXT(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_IsQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryIsQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_BeginQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryBeginQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_BeginQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryBeginQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_BeginQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryBeginQueryEXT(
	this js.Ref, retPtr unsafe.Pointer,
	target uint32,
	query js.Ref)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_BeginQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryBeginQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target uint32,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_EndQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryEndQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_EndQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryEndQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_EndQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryEndQueryEXT(
	this js.Ref, retPtr unsafe.Pointer,
	target uint32)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_EndQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryEndQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_QueryCounterEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryQueryCounterEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_QueryCounterEXT
//go:noescape
func FuncEXT_disjoint_timer_queryQueryCounterEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_QueryCounterEXT
//go:noescape
func CallEXT_disjoint_timer_queryQueryCounterEXT(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	target uint32)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_QueryCounterEXT
//go:noescape
func TryEXT_disjoint_timer_queryQueryCounterEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	target uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_GetQueryEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryGetQueryEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_GetQueryEXT
//go:noescape
func FuncEXT_disjoint_timer_queryGetQueryEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_GetQueryEXT
//go:noescape
func CallEXT_disjoint_timer_queryGetQueryEXT(
	this js.Ref, retPtr unsafe.Pointer,
	target uint32,
	pname uint32)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_GetQueryEXT
//go:noescape
func TryEXT_disjoint_timer_queryGetQueryEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target uint32,
	pname uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_GetQueryObjectEXT
//go:noescape
func HasFuncEXT_disjoint_timer_queryGetQueryObjectEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_GetQueryObjectEXT
//go:noescape
func FuncEXT_disjoint_timer_queryGetQueryObjectEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_GetQueryObjectEXT
//go:noescape
func CallEXT_disjoint_timer_queryGetQueryObjectEXT(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	pname uint32)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_GetQueryObjectEXT
//go:noescape
func TryEXT_disjoint_timer_queryGetQueryObjectEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	pname uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_EXT_disjoint_timer_query_webgl2_QueryCounterEXT
//go:noescape
func HasFuncEXT_disjoint_timer_query_webgl2QueryCounterEXT(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EXT_disjoint_timer_query_webgl2_QueryCounterEXT
//go:noescape
func FuncEXT_disjoint_timer_query_webgl2QueryCounterEXT(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EXT_disjoint_timer_query_webgl2_QueryCounterEXT
//go:noescape
func CallEXT_disjoint_timer_query_webgl2QueryCounterEXT(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	target uint32)

//go:wasmimport plat/js/web try_EXT_disjoint_timer_query_webgl2_QueryCounterEXT
//go:noescape
func TryEXT_disjoint_timer_query_webgl2QueryCounterEXT(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	target uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_EcKeyAlgorithm
//go:noescape
func EcKeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EcKeyAlgorithm
//go:noescape
func EcKeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EcKeyGenParams
//go:noescape
func EcKeyGenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EcKeyGenParams
//go:noescape
func EcKeyGenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EcKeyImportParams
//go:noescape
func EcKeyImportParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EcKeyImportParams
//go:noescape
func EcKeyImportParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EcdhKeyDeriveParams
//go:noescape
func EcdhKeyDeriveParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EcdhKeyDeriveParams
//go:noescape
func EcdhKeyDeriveParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EcdsaParams
//go:noescape
func EcdsaParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EcdsaParams
//go:noescape
func EcdsaParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_Ed448Params
//go:noescape
func Ed448ParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Ed448Params
//go:noescape
func Ed448ParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_EncodedVideoChunkType
//go:noescape
func ConstOfEncodedVideoChunkType(str js.Ref) uint32

//go:wasmimport plat/js/web store_EncodedVideoChunkInit
//go:noescape
func EncodedVideoChunkInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EncodedVideoChunkInit
//go:noescape
func EncodedVideoChunkInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_EncodedVideoChunk_EncodedVideoChunk
//go:noescape
func NewEncodedVideoChunkByEncodedVideoChunk(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_EncodedVideoChunk_Type
//go:noescape
func GetEncodedVideoChunkType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedVideoChunk_Timestamp
//go:noescape
func GetEncodedVideoChunkTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedVideoChunk_Duration
//go:noescape
func GetEncodedVideoChunkDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedVideoChunk_ByteLength
//go:noescape
func GetEncodedVideoChunkByteLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EncodedVideoChunk_CopyTo
//go:noescape
func HasFuncEncodedVideoChunkCopyTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EncodedVideoChunk_CopyTo
//go:noescape
func FuncEncodedVideoChunkCopyTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EncodedVideoChunk_CopyTo
//go:noescape
func CallEncodedVideoChunkCopyTo(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref)

//go:wasmimport plat/js/web try_EncodedVideoChunk_CopyTo
//go:noescape
func TryEncodedVideoChunkCopyTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_HardwareAcceleration
//go:noescape
func ConstOfHardwareAcceleration(str js.Ref) uint32

//go:wasmimport plat/js/web store_VideoDecoderConfig
//go:noescape
func VideoDecoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoDecoderConfig
//go:noescape
func VideoDecoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SvcOutputMetadata
//go:noescape
func SvcOutputMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SvcOutputMetadata
//go:noescape
func SvcOutputMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

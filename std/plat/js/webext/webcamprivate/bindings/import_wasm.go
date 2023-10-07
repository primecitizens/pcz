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

//go:wasmimport plat/js/webext/webcamprivate constof_AutofocusState
//go:noescape
func ConstOfAutofocusState(str js.Ref) uint32

//go:wasmimport plat/js/webext/webcamprivate constof_PanDirection
//go:noescape
func ConstOfPanDirection(str js.Ref) uint32

//go:wasmimport plat/js/webext/webcamprivate constof_Protocol
//go:noescape
func ConstOfProtocol(str js.Ref) uint32

//go:wasmimport plat/js/webext/webcamprivate store_ProtocolConfiguration
//go:noescape
func ProtocolConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webcamprivate load_ProtocolConfiguration
//go:noescape
func ProtocolConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webcamprivate store_Range
//go:noescape
func RangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webcamprivate load_Range
//go:noescape
func RangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webcamprivate constof_TiltDirection
//go:noescape
func ConstOfTiltDirection(str js.Ref) uint32

//go:wasmimport plat/js/webext/webcamprivate store_WebcamConfiguration
//go:noescape
func WebcamConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webcamprivate load_WebcamConfiguration
//go:noescape
func WebcamConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webcamprivate store_WebcamCurrentConfiguration
//go:noescape
func WebcamCurrentConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webcamprivate load_WebcamCurrentConfiguration
//go:noescape
func WebcamCurrentConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webcamprivate has_CloseWebcam
//go:noescape
func HasFuncCloseWebcam() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_CloseWebcam
//go:noescape
func FuncCloseWebcam(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_CloseWebcam
//go:noescape
func CallCloseWebcam(
	retPtr unsafe.Pointer,
	webcamId js.Ref)

//go:wasmimport plat/js/webext/webcamprivate try_CloseWebcam
//go:noescape
func TryCloseWebcam(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	webcamId js.Ref)

//go:wasmimport plat/js/webext/webcamprivate try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_OpenSerialWebcam
//go:noescape
func HasFuncOpenSerialWebcam() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_OpenSerialWebcam
//go:noescape
func FuncOpenSerialWebcam(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_OpenSerialWebcam
//go:noescape
func CallOpenSerialWebcam(
	retPtr unsafe.Pointer,
	path js.Ref,
	protocol unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate try_OpenSerialWebcam
//go:noescape
func TryOpenSerialWebcam(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	protocol unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_Reset
//go:noescape
func HasFuncReset() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_Reset
//go:noescape
func FuncReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_Reset
//go:noescape
func CallReset(
	retPtr unsafe.Pointer,
	webcamId js.Ref,
	config unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate try_Reset
//go:noescape
func TryReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_RestoreCameraPreset
//go:noescape
func HasFuncRestoreCameraPreset() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_RestoreCameraPreset
//go:noescape
func FuncRestoreCameraPreset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_RestoreCameraPreset
//go:noescape
func CallRestoreCameraPreset(
	retPtr unsafe.Pointer,
	webcamId js.Ref,
	presetNumber float64)

//go:wasmimport plat/js/webext/webcamprivate try_RestoreCameraPreset
//go:noescape
func TryRestoreCameraPreset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref,
	presetNumber float64) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_Set
//go:noescape
func HasFuncSet() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_Set
//go:noescape
func FuncSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_Set
//go:noescape
func CallSet(
	retPtr unsafe.Pointer,
	webcamId js.Ref,
	config unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate try_Set
//go:noescape
func TrySet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_SetCameraPreset
//go:noescape
func HasFuncSetCameraPreset() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_SetCameraPreset
//go:noescape
func FuncSetCameraPreset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_SetCameraPreset
//go:noescape
func CallSetCameraPreset(
	retPtr unsafe.Pointer,
	webcamId js.Ref,
	presetNumber float64)

//go:wasmimport plat/js/webext/webcamprivate try_SetCameraPreset
//go:noescape
func TrySetCameraPreset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref,
	presetNumber float64) (ok js.Ref)

//go:wasmimport plat/js/webext/webcamprivate has_SetHome
//go:noescape
func HasFuncSetHome() js.Ref

//go:wasmimport plat/js/webext/webcamprivate func_SetHome
//go:noescape
func FuncSetHome(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webcamprivate call_SetHome
//go:noescape
func CallSetHome(
	retPtr unsafe.Pointer,
	webcamId js.Ref)

//go:wasmimport plat/js/webext/webcamprivate try_SetHome
//go:noescape
func TrySetHome(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	webcamId js.Ref) (ok js.Ref)

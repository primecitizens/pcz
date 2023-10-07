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

//go:wasmimport plat/js/webext/terminalprivate store_GetOSInfoReturnType
//go:noescape
func GetOSInfoReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/terminalprivate load_GetOSInfoReturnType
//go:noescape
func GetOSInfoReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/terminalprivate store_OpenWindowArgData
//go:noescape
func OpenWindowArgDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/terminalprivate load_OpenWindowArgData
//go:noescape
func OpenWindowArgDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/terminalprivate constof_OutputType
//go:noescape
func ConstOfOutputType(str js.Ref) uint32

//go:wasmimport plat/js/webext/terminalprivate has_AckOutput
//go:noescape
func HasFuncAckOutput() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_AckOutput
//go:noescape
func FuncAckOutput(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_AckOutput
//go:noescape
func CallAckOutput(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_AckOutput
//go:noescape
func TryAckOutput(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_CloseTerminalProcess
//go:noescape
func HasFuncCloseTerminalProcess() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_CloseTerminalProcess
//go:noescape
func FuncCloseTerminalProcess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_CloseTerminalProcess
//go:noescape
func CallCloseTerminalProcess(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_CloseTerminalProcess
//go:noescape
func TryCloseTerminalProcess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_GetOSInfo
//go:noescape
func HasFuncGetOSInfo() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_GetOSInfo
//go:noescape
func FuncGetOSInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_GetOSInfo
//go:noescape
func CallGetOSInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate try_GetOSInfo
//go:noescape
func TryGetOSInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_GetPrefs
//go:noescape
func HasFuncGetPrefs() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_GetPrefs
//go:noescape
func FuncGetPrefs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_GetPrefs
//go:noescape
func CallGetPrefs(
	retPtr unsafe.Pointer,
	paths js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_GetPrefs
//go:noescape
func TryGetPrefs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	paths js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OnPrefChanged
//go:noescape
func HasFuncOnPrefChanged() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OnPrefChanged
//go:noescape
func FuncOnPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OnPrefChanged
//go:noescape
func CallOnPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OnPrefChanged
//go:noescape
func TryOnPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OffPrefChanged
//go:noescape
func HasFuncOffPrefChanged() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OffPrefChanged
//go:noescape
func FuncOffPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OffPrefChanged
//go:noescape
func CallOffPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OffPrefChanged
//go:noescape
func TryOffPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_HasOnPrefChanged
//go:noescape
func HasFuncHasOnPrefChanged() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_HasOnPrefChanged
//go:noescape
func FuncHasOnPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_HasOnPrefChanged
//go:noescape
func CallHasOnPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_HasOnPrefChanged
//go:noescape
func TryHasOnPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OnProcessOutput
//go:noescape
func HasFuncOnProcessOutput() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OnProcessOutput
//go:noescape
func FuncOnProcessOutput(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OnProcessOutput
//go:noescape
func CallOnProcessOutput(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OnProcessOutput
//go:noescape
func TryOnProcessOutput(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OffProcessOutput
//go:noescape
func HasFuncOffProcessOutput() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OffProcessOutput
//go:noescape
func FuncOffProcessOutput(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OffProcessOutput
//go:noescape
func CallOffProcessOutput(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OffProcessOutput
//go:noescape
func TryOffProcessOutput(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_HasOnProcessOutput
//go:noescape
func HasFuncHasOnProcessOutput() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_HasOnProcessOutput
//go:noescape
func FuncHasOnProcessOutput(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_HasOnProcessOutput
//go:noescape
func CallHasOnProcessOutput(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_HasOnProcessOutput
//go:noescape
func TryHasOnProcessOutput(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OnTerminalResize
//go:noescape
func HasFuncOnTerminalResize() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OnTerminalResize
//go:noescape
func FuncOnTerminalResize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OnTerminalResize
//go:noescape
func CallOnTerminalResize(
	retPtr unsafe.Pointer,
	id js.Ref,
	width float64,
	height float64)

//go:wasmimport plat/js/webext/terminalprivate try_OnTerminalResize
//go:noescape
func TryOnTerminalResize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	width float64,
	height float64) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OpenOptionsPage
//go:noescape
func HasFuncOpenOptionsPage() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OpenOptionsPage
//go:noescape
func FuncOpenOptionsPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OpenOptionsPage
//go:noescape
func CallOpenOptionsPage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate try_OpenOptionsPage
//go:noescape
func TryOpenOptionsPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OpenSettingsSubpage
//go:noescape
func HasFuncOpenSettingsSubpage() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OpenSettingsSubpage
//go:noescape
func FuncOpenSettingsSubpage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OpenSettingsSubpage
//go:noescape
func CallOpenSettingsSubpage(
	retPtr unsafe.Pointer,
	subpage js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OpenSettingsSubpage
//go:noescape
func TryOpenSettingsSubpage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	subpage js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OpenTerminalProcess
//go:noescape
func HasFuncOpenTerminalProcess() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OpenTerminalProcess
//go:noescape
func FuncOpenTerminalProcess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OpenTerminalProcess
//go:noescape
func CallOpenTerminalProcess(
	retPtr unsafe.Pointer,
	processName js.Ref,
	args js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OpenTerminalProcess
//go:noescape
func TryOpenTerminalProcess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	processName js.Ref,
	args js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OpenVmshellProcess
//go:noescape
func HasFuncOpenVmshellProcess() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OpenVmshellProcess
//go:noescape
func FuncOpenVmshellProcess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OpenVmshellProcess
//go:noescape
func CallOpenVmshellProcess(
	retPtr unsafe.Pointer,
	args js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_OpenVmshellProcess
//go:noescape
func TryOpenVmshellProcess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	args js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_OpenWindow
//go:noescape
func HasFuncOpenWindow() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_OpenWindow
//go:noescape
func FuncOpenWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_OpenWindow
//go:noescape
func CallOpenWindow(
	retPtr unsafe.Pointer,
	data unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate try_OpenWindow
//go:noescape
func TryOpenWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_SendInput
//go:noescape
func HasFuncSendInput() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_SendInput
//go:noescape
func FuncSendInput(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_SendInput
//go:noescape
func CallSendInput(
	retPtr unsafe.Pointer,
	id js.Ref,
	input js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_SendInput
//go:noescape
func TrySendInput(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/terminalprivate has_SetPrefs
//go:noescape
func HasFuncSetPrefs() js.Ref

//go:wasmimport plat/js/webext/terminalprivate func_SetPrefs
//go:noescape
func FuncSetPrefs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/terminalprivate call_SetPrefs
//go:noescape
func CallSetPrefs(
	retPtr unsafe.Pointer,
	prefs js.Ref)

//go:wasmimport plat/js/webext/terminalprivate try_SetPrefs
//go:noescape
func TrySetPrefs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	prefs js.Ref) (ok js.Ref)

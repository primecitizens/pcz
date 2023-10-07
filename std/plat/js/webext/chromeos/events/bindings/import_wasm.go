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

//go:wasmimport plat/js/webext/chromeos/events constof_AudioJackDeviceType
//go:noescape
func ConstOfAudioJackDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_AudioJackEvent
//go:noescape
func ConstOfAudioJackEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_AudioJackEventInfo
//go:noescape
func AudioJackEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_AudioJackEventInfo
//go:noescape
func AudioJackEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_DisplayInputType
//go:noescape
func ConstOfDisplayInputType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_EventCategory
//go:noescape
func ConstOfEventCategory(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_EventSupportStatus
//go:noescape
func ConstOfEventSupportStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_EventSupportStatusInfo
//go:noescape
func EventSupportStatusInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_EventSupportStatusInfo
//go:noescape
func EventSupportStatusInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_ExternalDisplayEvent
//go:noescape
func ConstOfExternalDisplayEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_ExternalDisplayInfo
//go:noescape
func ExternalDisplayInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_ExternalDisplayInfo
//go:noescape
func ExternalDisplayInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_ExternalDisplayEventInfo
//go:noescape
func ExternalDisplayEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_ExternalDisplayEventInfo
//go:noescape
func ExternalDisplayEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_InputTouchButton
//go:noescape
func ConstOfInputTouchButton(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_InputTouchButtonState
//go:noescape
func ConstOfInputTouchButtonState(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_KeyboardConnectionType
//go:noescape
func ConstOfKeyboardConnectionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_PhysicalKeyboardLayout
//go:noescape
func ConstOfPhysicalKeyboardLayout(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_MechanicalKeyboardLayout
//go:noescape
func ConstOfMechanicalKeyboardLayout(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_KeyboardNumberPadPresence
//go:noescape
func ConstOfKeyboardNumberPadPresence(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_KeyboardTopRowKey
//go:noescape
func ConstOfKeyboardTopRowKey(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events constof_KeyboardTopRightKey
//go:noescape
func ConstOfKeyboardTopRightKey(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_KeyboardInfo
//go:noescape
func KeyboardInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_KeyboardInfo
//go:noescape
func KeyboardInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_KeyboardDiagnosticEventInfo
//go:noescape
func KeyboardDiagnosticEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_KeyboardDiagnosticEventInfo
//go:noescape
func KeyboardDiagnosticEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_LidEvent
//go:noescape
func ConstOfLidEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_LidEventInfo
//go:noescape
func LidEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_LidEventInfo
//go:noescape
func LidEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_PowerEvent
//go:noescape
func ConstOfPowerEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_PowerEventInfo
//go:noescape
func PowerEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_PowerEventInfo
//go:noescape
func PowerEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_SdCardEvent
//go:noescape
func ConstOfSdCardEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_SdCardEventInfo
//go:noescape
func SdCardEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_SdCardEventInfo
//go:noescape
func SdCardEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_StylusConnectedEventInfo
//go:noescape
func StylusConnectedEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_StylusConnectedEventInfo
//go:noescape
func StylusConnectedEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_StylusGarageEvent
//go:noescape
func ConstOfStylusGarageEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_StylusGarageEventInfo
//go:noescape
func StylusGarageEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_StylusGarageEventInfo
//go:noescape
func StylusGarageEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_StylusTouchPointInfo
//go:noescape
func StylusTouchPointInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_StylusTouchPointInfo
//go:noescape
func StylusTouchPointInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_StylusTouchEventInfo
//go:noescape
func StylusTouchEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_StylusTouchEventInfo
//go:noescape
func StylusTouchEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchPointInfo
//go:noescape
func TouchPointInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchPointInfo
//go:noescape
func TouchPointInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchpadButtonEventInfo
//go:noescape
func TouchpadButtonEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchpadButtonEventInfo
//go:noescape
func TouchpadButtonEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchpadConnectedEventInfo
//go:noescape
func TouchpadConnectedEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchpadConnectedEventInfo
//go:noescape
func TouchpadConnectedEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchpadTouchEventInfo
//go:noescape
func TouchpadTouchEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchpadTouchEventInfo
//go:noescape
func TouchpadTouchEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchscreenConnectedEventInfo
//go:noescape
func TouchscreenConnectedEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchscreenConnectedEventInfo
//go:noescape
func TouchscreenConnectedEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events store_TouchscreenTouchEventInfo
//go:noescape
func TouchscreenTouchEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_TouchscreenTouchEventInfo
//go:noescape
func TouchscreenTouchEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events constof_UsbEvent
//go:noescape
func ConstOfUsbEvent(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/events store_UsbEventInfo
//go:noescape
func UsbEventInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/events load_UsbEventInfo
//go:noescape
func UsbEventInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/events has_IsEventSupported
//go:noescape
func HasFuncIsEventSupported() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_IsEventSupported
//go:noescape
func FuncIsEventSupported(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_IsEventSupported
//go:noescape
func CallIsEventSupported(
	retPtr unsafe.Pointer,
	category uint32)

//go:wasmimport plat/js/webext/chromeos/events try_IsEventSupported
//go:noescape
func TryIsEventSupported(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	category uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnAudioJackEvent
//go:noescape
func HasFuncOnAudioJackEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnAudioJackEvent
//go:noescape
func FuncOnAudioJackEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnAudioJackEvent
//go:noescape
func CallOnAudioJackEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnAudioJackEvent
//go:noescape
func TryOnAudioJackEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffAudioJackEvent
//go:noescape
func HasFuncOffAudioJackEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffAudioJackEvent
//go:noescape
func FuncOffAudioJackEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffAudioJackEvent
//go:noescape
func CallOffAudioJackEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffAudioJackEvent
//go:noescape
func TryOffAudioJackEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnAudioJackEvent
//go:noescape
func HasFuncHasOnAudioJackEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnAudioJackEvent
//go:noescape
func FuncHasOnAudioJackEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnAudioJackEvent
//go:noescape
func CallHasOnAudioJackEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnAudioJackEvent
//go:noescape
func TryHasOnAudioJackEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnExternalDisplayEvent
//go:noescape
func HasFuncOnExternalDisplayEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnExternalDisplayEvent
//go:noescape
func FuncOnExternalDisplayEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnExternalDisplayEvent
//go:noescape
func CallOnExternalDisplayEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnExternalDisplayEvent
//go:noescape
func TryOnExternalDisplayEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffExternalDisplayEvent
//go:noescape
func HasFuncOffExternalDisplayEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffExternalDisplayEvent
//go:noescape
func FuncOffExternalDisplayEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffExternalDisplayEvent
//go:noescape
func CallOffExternalDisplayEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffExternalDisplayEvent
//go:noescape
func TryOffExternalDisplayEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnExternalDisplayEvent
//go:noescape
func HasFuncHasOnExternalDisplayEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnExternalDisplayEvent
//go:noescape
func FuncHasOnExternalDisplayEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnExternalDisplayEvent
//go:noescape
func CallHasOnExternalDisplayEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnExternalDisplayEvent
//go:noescape
func TryHasOnExternalDisplayEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnKeyboardDiagnosticEvent
//go:noescape
func HasFuncOnKeyboardDiagnosticEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnKeyboardDiagnosticEvent
//go:noescape
func FuncOnKeyboardDiagnosticEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnKeyboardDiagnosticEvent
//go:noescape
func CallOnKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnKeyboardDiagnosticEvent
//go:noescape
func TryOnKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffKeyboardDiagnosticEvent
//go:noescape
func HasFuncOffKeyboardDiagnosticEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffKeyboardDiagnosticEvent
//go:noescape
func FuncOffKeyboardDiagnosticEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffKeyboardDiagnosticEvent
//go:noescape
func CallOffKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffKeyboardDiagnosticEvent
//go:noescape
func TryOffKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnKeyboardDiagnosticEvent
//go:noescape
func HasFuncHasOnKeyboardDiagnosticEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnKeyboardDiagnosticEvent
//go:noescape
func FuncHasOnKeyboardDiagnosticEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnKeyboardDiagnosticEvent
//go:noescape
func CallHasOnKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnKeyboardDiagnosticEvent
//go:noescape
func TryHasOnKeyboardDiagnosticEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnLidEvent
//go:noescape
func HasFuncOnLidEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnLidEvent
//go:noescape
func FuncOnLidEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnLidEvent
//go:noescape
func CallOnLidEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnLidEvent
//go:noescape
func TryOnLidEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffLidEvent
//go:noescape
func HasFuncOffLidEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffLidEvent
//go:noescape
func FuncOffLidEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffLidEvent
//go:noescape
func CallOffLidEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffLidEvent
//go:noescape
func TryOffLidEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnLidEvent
//go:noescape
func HasFuncHasOnLidEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnLidEvent
//go:noescape
func FuncHasOnLidEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnLidEvent
//go:noescape
func CallHasOnLidEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnLidEvent
//go:noescape
func TryHasOnLidEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnPowerEvent
//go:noescape
func HasFuncOnPowerEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnPowerEvent
//go:noescape
func FuncOnPowerEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnPowerEvent
//go:noescape
func CallOnPowerEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnPowerEvent
//go:noescape
func TryOnPowerEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffPowerEvent
//go:noescape
func HasFuncOffPowerEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffPowerEvent
//go:noescape
func FuncOffPowerEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffPowerEvent
//go:noescape
func CallOffPowerEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffPowerEvent
//go:noescape
func TryOffPowerEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnPowerEvent
//go:noescape
func HasFuncHasOnPowerEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnPowerEvent
//go:noescape
func FuncHasOnPowerEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnPowerEvent
//go:noescape
func CallHasOnPowerEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnPowerEvent
//go:noescape
func TryHasOnPowerEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnSdCardEvent
//go:noescape
func HasFuncOnSdCardEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnSdCardEvent
//go:noescape
func FuncOnSdCardEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnSdCardEvent
//go:noescape
func CallOnSdCardEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnSdCardEvent
//go:noescape
func TryOnSdCardEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffSdCardEvent
//go:noescape
func HasFuncOffSdCardEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffSdCardEvent
//go:noescape
func FuncOffSdCardEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffSdCardEvent
//go:noescape
func CallOffSdCardEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffSdCardEvent
//go:noescape
func TryOffSdCardEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnSdCardEvent
//go:noescape
func HasFuncHasOnSdCardEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnSdCardEvent
//go:noescape
func FuncHasOnSdCardEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnSdCardEvent
//go:noescape
func CallHasOnSdCardEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnSdCardEvent
//go:noescape
func TryHasOnSdCardEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnStylusConnectedEvent
//go:noescape
func HasFuncOnStylusConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnStylusConnectedEvent
//go:noescape
func FuncOnStylusConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnStylusConnectedEvent
//go:noescape
func CallOnStylusConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnStylusConnectedEvent
//go:noescape
func TryOnStylusConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffStylusConnectedEvent
//go:noescape
func HasFuncOffStylusConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffStylusConnectedEvent
//go:noescape
func FuncOffStylusConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffStylusConnectedEvent
//go:noescape
func CallOffStylusConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffStylusConnectedEvent
//go:noescape
func TryOffStylusConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnStylusConnectedEvent
//go:noescape
func HasFuncHasOnStylusConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnStylusConnectedEvent
//go:noescape
func FuncHasOnStylusConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnStylusConnectedEvent
//go:noescape
func CallHasOnStylusConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnStylusConnectedEvent
//go:noescape
func TryHasOnStylusConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnStylusGarageEvent
//go:noescape
func HasFuncOnStylusGarageEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnStylusGarageEvent
//go:noescape
func FuncOnStylusGarageEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnStylusGarageEvent
//go:noescape
func CallOnStylusGarageEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnStylusGarageEvent
//go:noescape
func TryOnStylusGarageEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffStylusGarageEvent
//go:noescape
func HasFuncOffStylusGarageEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffStylusGarageEvent
//go:noescape
func FuncOffStylusGarageEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffStylusGarageEvent
//go:noescape
func CallOffStylusGarageEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffStylusGarageEvent
//go:noescape
func TryOffStylusGarageEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnStylusGarageEvent
//go:noescape
func HasFuncHasOnStylusGarageEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnStylusGarageEvent
//go:noescape
func FuncHasOnStylusGarageEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnStylusGarageEvent
//go:noescape
func CallHasOnStylusGarageEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnStylusGarageEvent
//go:noescape
func TryHasOnStylusGarageEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnStylusTouchEvent
//go:noescape
func HasFuncOnStylusTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnStylusTouchEvent
//go:noescape
func FuncOnStylusTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnStylusTouchEvent
//go:noescape
func CallOnStylusTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnStylusTouchEvent
//go:noescape
func TryOnStylusTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffStylusTouchEvent
//go:noescape
func HasFuncOffStylusTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffStylusTouchEvent
//go:noescape
func FuncOffStylusTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffStylusTouchEvent
//go:noescape
func CallOffStylusTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffStylusTouchEvent
//go:noescape
func TryOffStylusTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnStylusTouchEvent
//go:noescape
func HasFuncHasOnStylusTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnStylusTouchEvent
//go:noescape
func FuncHasOnStylusTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnStylusTouchEvent
//go:noescape
func CallHasOnStylusTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnStylusTouchEvent
//go:noescape
func TryHasOnStylusTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnTouchpadButtonEvent
//go:noescape
func HasFuncOnTouchpadButtonEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnTouchpadButtonEvent
//go:noescape
func FuncOnTouchpadButtonEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnTouchpadButtonEvent
//go:noescape
func CallOnTouchpadButtonEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnTouchpadButtonEvent
//go:noescape
func TryOnTouchpadButtonEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffTouchpadButtonEvent
//go:noescape
func HasFuncOffTouchpadButtonEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffTouchpadButtonEvent
//go:noescape
func FuncOffTouchpadButtonEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffTouchpadButtonEvent
//go:noescape
func CallOffTouchpadButtonEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffTouchpadButtonEvent
//go:noescape
func TryOffTouchpadButtonEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnTouchpadButtonEvent
//go:noescape
func HasFuncHasOnTouchpadButtonEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnTouchpadButtonEvent
//go:noescape
func FuncHasOnTouchpadButtonEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnTouchpadButtonEvent
//go:noescape
func CallHasOnTouchpadButtonEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnTouchpadButtonEvent
//go:noescape
func TryHasOnTouchpadButtonEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnTouchpadConnectedEvent
//go:noescape
func HasFuncOnTouchpadConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnTouchpadConnectedEvent
//go:noescape
func FuncOnTouchpadConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnTouchpadConnectedEvent
//go:noescape
func CallOnTouchpadConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnTouchpadConnectedEvent
//go:noescape
func TryOnTouchpadConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffTouchpadConnectedEvent
//go:noescape
func HasFuncOffTouchpadConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffTouchpadConnectedEvent
//go:noescape
func FuncOffTouchpadConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffTouchpadConnectedEvent
//go:noescape
func CallOffTouchpadConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffTouchpadConnectedEvent
//go:noescape
func TryOffTouchpadConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnTouchpadConnectedEvent
//go:noescape
func HasFuncHasOnTouchpadConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnTouchpadConnectedEvent
//go:noescape
func FuncHasOnTouchpadConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnTouchpadConnectedEvent
//go:noescape
func CallHasOnTouchpadConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnTouchpadConnectedEvent
//go:noescape
func TryHasOnTouchpadConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnTouchpadTouchEvent
//go:noescape
func HasFuncOnTouchpadTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnTouchpadTouchEvent
//go:noescape
func FuncOnTouchpadTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnTouchpadTouchEvent
//go:noescape
func CallOnTouchpadTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnTouchpadTouchEvent
//go:noescape
func TryOnTouchpadTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffTouchpadTouchEvent
//go:noescape
func HasFuncOffTouchpadTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffTouchpadTouchEvent
//go:noescape
func FuncOffTouchpadTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffTouchpadTouchEvent
//go:noescape
func CallOffTouchpadTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffTouchpadTouchEvent
//go:noescape
func TryOffTouchpadTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnTouchpadTouchEvent
//go:noescape
func HasFuncHasOnTouchpadTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnTouchpadTouchEvent
//go:noescape
func FuncHasOnTouchpadTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnTouchpadTouchEvent
//go:noescape
func CallHasOnTouchpadTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnTouchpadTouchEvent
//go:noescape
func TryHasOnTouchpadTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnTouchscreenConnectedEvent
//go:noescape
func HasFuncOnTouchscreenConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnTouchscreenConnectedEvent
//go:noescape
func FuncOnTouchscreenConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnTouchscreenConnectedEvent
//go:noescape
func CallOnTouchscreenConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnTouchscreenConnectedEvent
//go:noescape
func TryOnTouchscreenConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffTouchscreenConnectedEvent
//go:noescape
func HasFuncOffTouchscreenConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffTouchscreenConnectedEvent
//go:noescape
func FuncOffTouchscreenConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffTouchscreenConnectedEvent
//go:noescape
func CallOffTouchscreenConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffTouchscreenConnectedEvent
//go:noescape
func TryOffTouchscreenConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnTouchscreenConnectedEvent
//go:noescape
func HasFuncHasOnTouchscreenConnectedEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnTouchscreenConnectedEvent
//go:noescape
func FuncHasOnTouchscreenConnectedEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnTouchscreenConnectedEvent
//go:noescape
func CallHasOnTouchscreenConnectedEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnTouchscreenConnectedEvent
//go:noescape
func TryHasOnTouchscreenConnectedEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnTouchscreenTouchEvent
//go:noescape
func HasFuncOnTouchscreenTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnTouchscreenTouchEvent
//go:noescape
func FuncOnTouchscreenTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnTouchscreenTouchEvent
//go:noescape
func CallOnTouchscreenTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnTouchscreenTouchEvent
//go:noescape
func TryOnTouchscreenTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffTouchscreenTouchEvent
//go:noescape
func HasFuncOffTouchscreenTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffTouchscreenTouchEvent
//go:noescape
func FuncOffTouchscreenTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffTouchscreenTouchEvent
//go:noescape
func CallOffTouchscreenTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffTouchscreenTouchEvent
//go:noescape
func TryOffTouchscreenTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnTouchscreenTouchEvent
//go:noescape
func HasFuncHasOnTouchscreenTouchEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnTouchscreenTouchEvent
//go:noescape
func FuncHasOnTouchscreenTouchEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnTouchscreenTouchEvent
//go:noescape
func CallHasOnTouchscreenTouchEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnTouchscreenTouchEvent
//go:noescape
func TryHasOnTouchscreenTouchEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OnUsbEvent
//go:noescape
func HasFuncOnUsbEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OnUsbEvent
//go:noescape
func FuncOnUsbEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OnUsbEvent
//go:noescape
func CallOnUsbEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OnUsbEvent
//go:noescape
func TryOnUsbEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_OffUsbEvent
//go:noescape
func HasFuncOffUsbEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_OffUsbEvent
//go:noescape
func FuncOffUsbEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_OffUsbEvent
//go:noescape
func CallOffUsbEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_OffUsbEvent
//go:noescape
func TryOffUsbEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_HasOnUsbEvent
//go:noescape
func HasFuncHasOnUsbEvent() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_HasOnUsbEvent
//go:noescape
func FuncHasOnUsbEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_HasOnUsbEvent
//go:noescape
func CallHasOnUsbEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/events try_HasOnUsbEvent
//go:noescape
func TryHasOnUsbEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_StartCapturingEvents
//go:noescape
func HasFuncStartCapturingEvents() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_StartCapturingEvents
//go:noescape
func FuncStartCapturingEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_StartCapturingEvents
//go:noescape
func CallStartCapturingEvents(
	retPtr unsafe.Pointer,
	category uint32)

//go:wasmimport plat/js/webext/chromeos/events try_StartCapturingEvents
//go:noescape
func TryStartCapturingEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	category uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/events has_StopCapturingEvents
//go:noescape
func HasFuncStopCapturingEvents() js.Ref

//go:wasmimport plat/js/webext/chromeos/events func_StopCapturingEvents
//go:noescape
func FuncStopCapturingEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/events call_StopCapturingEvents
//go:noescape
func CallStopCapturingEvents(
	retPtr unsafe.Pointer,
	category uint32)

//go:wasmimport plat/js/webext/chromeos/events try_StopCapturingEvents
//go:noescape
func TryStopCapturingEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	category uint32) (ok js.Ref)

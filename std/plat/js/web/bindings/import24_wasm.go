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

//go:wasmimport plat/js/web constof_NavigationHistoryBehavior
//go:noescape
func ConstOfNavigationHistoryBehavior(str js.Ref) uint32

//go:wasmimport plat/js/web store_NavigationNavigateOptions
//go:noescape
func NavigationNavigateOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationNavigateOptions
//go:noescape
func NavigationNavigateOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigationReloadOptions
//go:noescape
func NavigationReloadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationReloadOptions
//go:noescape
func NavigationReloadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigationOptions
//go:noescape
func NavigationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationOptions
//go:noescape
func NavigationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_NavigationType
//go:noescape
func ConstOfNavigationType(str js.Ref) uint32

//go:wasmimport plat/js/web get_NavigationTransition_NavigationType
//go:noescape
func GetNavigationTransitionNavigationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationTransition_From
//go:noescape
func GetNavigationTransitionFrom(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationTransition_Finished
//go:noescape
func GetNavigationTransitionFinished(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigation_CurrentEntry
//go:noescape
func GetNavigationCurrentEntry(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigation_Transition
//go:noescape
func GetNavigationTransition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigation_CanGoBack
//go:noescape
func GetNavigationCanGoBack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigation_CanGoForward
//go:noescape
func GetNavigationCanGoForward(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Entries
//go:noescape
func HasFuncNavigationEntries(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Entries
//go:noescape
func FuncNavigationEntries(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Entries
//go:noescape
func CallNavigationEntries(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Entries
//go:noescape
func TryNavigationEntries(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_UpdateCurrentEntry
//go:noescape
func HasFuncNavigationUpdateCurrentEntry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_UpdateCurrentEntry
//go:noescape
func FuncNavigationUpdateCurrentEntry(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_UpdateCurrentEntry
//go:noescape
func CallNavigationUpdateCurrentEntry(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_UpdateCurrentEntry
//go:noescape
func TryNavigationUpdateCurrentEntry(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Navigate
//go:noescape
func HasFuncNavigationNavigate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Navigate
//go:noescape
func FuncNavigationNavigate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Navigate
//go:noescape
func CallNavigationNavigate(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Navigate
//go:noescape
func TryNavigationNavigate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Navigate1
//go:noescape
func HasFuncNavigationNavigate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Navigate1
//go:noescape
func FuncNavigationNavigate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Navigate1
//go:noescape
func CallNavigationNavigate1(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Navigation_Navigate1
//go:noescape
func TryNavigationNavigate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Reload
//go:noescape
func HasFuncNavigationReload(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Reload
//go:noescape
func FuncNavigationReload(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Reload
//go:noescape
func CallNavigationReload(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Reload
//go:noescape
func TryNavigationReload(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Reload1
//go:noescape
func HasFuncNavigationReload1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Reload1
//go:noescape
func FuncNavigationReload1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Reload1
//go:noescape
func CallNavigationReload1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Reload1
//go:noescape
func TryNavigationReload1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_TraverseTo
//go:noescape
func HasFuncNavigationTraverseTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_TraverseTo
//go:noescape
func FuncNavigationTraverseTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_TraverseTo
//go:noescape
func CallNavigationTraverseTo(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_TraverseTo
//go:noescape
func TryNavigationTraverseTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_TraverseTo1
//go:noescape
func HasFuncNavigationTraverseTo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_TraverseTo1
//go:noescape
func FuncNavigationTraverseTo1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_TraverseTo1
//go:noescape
func CallNavigationTraverseTo1(
	this js.Ref, retPtr unsafe.Pointer,
	key js.Ref)

//go:wasmimport plat/js/web try_Navigation_TraverseTo1
//go:noescape
func TryNavigationTraverseTo1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Back
//go:noescape
func HasFuncNavigationBack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Back
//go:noescape
func FuncNavigationBack(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Back
//go:noescape
func CallNavigationBack(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Back
//go:noescape
func TryNavigationBack(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Back1
//go:noescape
func HasFuncNavigationBack1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Back1
//go:noescape
func FuncNavigationBack1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Back1
//go:noescape
func CallNavigationBack1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Back1
//go:noescape
func TryNavigationBack1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Forward
//go:noescape
func HasFuncNavigationForward(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Forward
//go:noescape
func FuncNavigationForward(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Forward
//go:noescape
func CallNavigationForward(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Forward
//go:noescape
func TryNavigationForward(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigation_Forward1
//go:noescape
func HasFuncNavigationForward1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigation_Forward1
//go:noescape
func FuncNavigationForward1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Navigation_Forward1
//go:noescape
func CallNavigationForward1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigation_Forward1
//go:noescape
func TryNavigationForward1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ElementDefinitionOptions
//go:noescape
func ElementDefinitionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ElementDefinitionOptions
//go:noescape
func ElementDefinitionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CustomElementRegistry_Define
//go:noescape
func HasFuncCustomElementRegistryDefine(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_Define
//go:noescape
func FuncCustomElementRegistryDefine(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_Define
//go:noescape
func CallCustomElementRegistryDefine(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	constructor js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CustomElementRegistry_Define
//go:noescape
func TryCustomElementRegistryDefine(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	constructor js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomElementRegistry_Define1
//go:noescape
func HasFuncCustomElementRegistryDefine1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_Define1
//go:noescape
func FuncCustomElementRegistryDefine1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_Define1
//go:noescape
func CallCustomElementRegistryDefine1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	constructor js.Ref)

//go:wasmimport plat/js/web try_CustomElementRegistry_Define1
//go:noescape
func TryCustomElementRegistryDefine1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	constructor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomElementRegistry_Get
//go:noescape
func HasFuncCustomElementRegistryGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_Get
//go:noescape
func FuncCustomElementRegistryGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_Get
//go:noescape
func CallCustomElementRegistryGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_CustomElementRegistry_Get
//go:noescape
func TryCustomElementRegistryGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomElementRegistry_GetName
//go:noescape
func HasFuncCustomElementRegistryGetName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_GetName
//go:noescape
func FuncCustomElementRegistryGetName(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_GetName
//go:noescape
func CallCustomElementRegistryGetName(
	this js.Ref, retPtr unsafe.Pointer,
	constructor js.Ref)

//go:wasmimport plat/js/web try_CustomElementRegistry_GetName
//go:noescape
func TryCustomElementRegistryGetName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constructor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomElementRegistry_WhenDefined
//go:noescape
func HasFuncCustomElementRegistryWhenDefined(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_WhenDefined
//go:noescape
func FuncCustomElementRegistryWhenDefined(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_WhenDefined
//go:noescape
func CallCustomElementRegistryWhenDefined(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_CustomElementRegistry_WhenDefined
//go:noescape
func TryCustomElementRegistryWhenDefined(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomElementRegistry_Upgrade
//go:noescape
func HasFuncCustomElementRegistryUpgrade(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomElementRegistry_Upgrade
//go:noescape
func FuncCustomElementRegistryUpgrade(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomElementRegistry_Upgrade
//go:noescape
func CallCustomElementRegistryUpgrade(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref)

//go:wasmimport plat/js/web try_CustomElementRegistry_Upgrade
//go:noescape
func TryCustomElementRegistryUpgrade(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_GamepadMappingType
//go:noescape
func ConstOfGamepadMappingType(str js.Ref) uint32

//go:wasmimport plat/js/web get_GamepadButton_Pressed
//go:noescape
func GetGamepadButtonPressed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadButton_Touched
//go:noescape
func GetGamepadButtonTouched(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadButton_Value
//go:noescape
func GetGamepadButtonValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_GamepadHand
//go:noescape
func ConstOfGamepadHand(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GamepadHapticEffectType
//go:noescape
func ConstOfGamepadHapticEffectType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GamepadHapticsResult
//go:noescape
func ConstOfGamepadHapticsResult(str js.Ref) uint32

//go:wasmimport plat/js/web store_GamepadEffectParameters
//go:noescape
func GamepadEffectParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GamepadEffectParameters
//go:noescape
func GamepadEffectParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GamepadHapticActuatorType
//go:noescape
func ConstOfGamepadHapticActuatorType(str js.Ref) uint32

//go:wasmimport plat/js/web get_GamepadHapticActuator_Type
//go:noescape
func GetGamepadHapticActuatorType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GamepadHapticActuator_CanPlayEffectType
//go:noescape
func HasFuncGamepadHapticActuatorCanPlayEffectType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GamepadHapticActuator_CanPlayEffectType
//go:noescape
func FuncGamepadHapticActuatorCanPlayEffectType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GamepadHapticActuator_CanPlayEffectType
//go:noescape
func CallGamepadHapticActuatorCanPlayEffectType(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/web try_GamepadHapticActuator_CanPlayEffectType
//go:noescape
func TryGamepadHapticActuatorCanPlayEffectType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GamepadHapticActuator_PlayEffect
//go:noescape
func HasFuncGamepadHapticActuatorPlayEffect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GamepadHapticActuator_PlayEffect
//go:noescape
func FuncGamepadHapticActuatorPlayEffect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GamepadHapticActuator_PlayEffect
//go:noescape
func CallGamepadHapticActuatorPlayEffect(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32,
	params unsafe.Pointer)

//go:wasmimport plat/js/web try_GamepadHapticActuator_PlayEffect
//go:noescape
func TryGamepadHapticActuatorPlayEffect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GamepadHapticActuator_PlayEffect1
//go:noescape
func HasFuncGamepadHapticActuatorPlayEffect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GamepadHapticActuator_PlayEffect1
//go:noescape
func FuncGamepadHapticActuatorPlayEffect1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GamepadHapticActuator_PlayEffect1
//go:noescape
func CallGamepadHapticActuatorPlayEffect1(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/web try_GamepadHapticActuator_PlayEffect1
//go:noescape
func TryGamepadHapticActuatorPlayEffect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_GamepadHapticActuator_Pulse
//go:noescape
func HasFuncGamepadHapticActuatorPulse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GamepadHapticActuator_Pulse
//go:noescape
func FuncGamepadHapticActuatorPulse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GamepadHapticActuator_Pulse
//go:noescape
func CallGamepadHapticActuatorPulse(
	this js.Ref, retPtr unsafe.Pointer,
	value float64,
	duration float64)

//go:wasmimport plat/js/web try_GamepadHapticActuator_Pulse
//go:noescape
func TryGamepadHapticActuatorPulse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64,
	duration float64) (ok js.Ref)

//go:wasmimport plat/js/web has_GamepadHapticActuator_Reset
//go:noescape
func HasFuncGamepadHapticActuatorReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GamepadHapticActuator_Reset
//go:noescape
func FuncGamepadHapticActuatorReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_GamepadHapticActuator_Reset
//go:noescape
func CallGamepadHapticActuatorReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GamepadHapticActuator_Reset
//go:noescape
func TryGamepadHapticActuatorReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_HasOrientation
//go:noescape
func GetGamepadPoseHasOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_HasPosition
//go:noescape
func GetGamepadPoseHasPosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_Position
//go:noescape
func GetGamepadPosePosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_LinearVelocity
//go:noescape
func GetGamepadPoseLinearVelocity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_LinearAcceleration
//go:noescape
func GetGamepadPoseLinearAcceleration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_Orientation
//go:noescape
func GetGamepadPoseOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_AngularVelocity
//go:noescape
func GetGamepadPoseAngularVelocity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadPose_AngularAcceleration
//go:noescape
func GetGamepadPoseAngularAcceleration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadTouch_TouchId
//go:noescape
func GetGamepadTouchTouchId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadTouch_SurfaceId
//go:noescape
func GetGamepadTouchSurfaceId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadTouch_Position
//go:noescape
func GetGamepadTouchPosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GamepadTouch_SurfaceDimensions
//go:noescape
func GetGamepadTouchSurfaceDimensions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Id
//go:noescape
func GetGamepadId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Index
//go:noescape
func GetGamepadIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Connected
//go:noescape
func GetGamepadConnected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Timestamp
//go:noescape
func GetGamepadTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Mapping
//go:noescape
func GetGamepadMapping(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Axes
//go:noescape
func GetGamepadAxes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Buttons
//go:noescape
func GetGamepadButtons(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Hand
//go:noescape
func GetGamepadHand(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_HapticActuators
//go:noescape
func GetGamepadHapticActuators(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_Pose
//go:noescape
func GetGamepadPose(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_TouchEvents
//go:noescape
func GetGamepadTouchEvents(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Gamepad_VibrationActuator
//go:noescape
func GetGamepadVibrationActuator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RelatedApplication
//go:noescape
func RelatedApplicationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RelatedApplication
//go:noescape
func RelatedApplicationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaKeySystemMediaCapability
//go:noescape
func MediaKeySystemMediaCapabilityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaKeySystemMediaCapability
//go:noescape
func MediaKeySystemMediaCapabilityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaKeysRequirement
//go:noescape
func ConstOfMediaKeysRequirement(str js.Ref) uint32

//go:wasmimport plat/js/web store_MediaKeySystemConfiguration
//go:noescape
func MediaKeySystemConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaKeySystemConfiguration
//go:noescape
func MediaKeySystemConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaKeySystemAccess_KeySystem
//go:noescape
func GetMediaKeySystemAccessKeySystem(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySystemAccess_GetConfiguration
//go:noescape
func HasFuncMediaKeySystemAccessGetConfiguration(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySystemAccess_GetConfiguration
//go:noescape
func FuncMediaKeySystemAccessGetConfiguration(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySystemAccess_GetConfiguration
//go:noescape
func CallMediaKeySystemAccessGetConfiguration(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaKeySystemAccess_GetConfiguration
//go:noescape
func TryMediaKeySystemAccessGetConfiguration(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySystemAccess_CreateMediaKeys
//go:noescape
func HasFuncMediaKeySystemAccessCreateMediaKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySystemAccess_CreateMediaKeys
//go:noescape
func FuncMediaKeySystemAccessCreateMediaKeys(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySystemAccess_CreateMediaKeys
//go:noescape
func CallMediaKeySystemAccessCreateMediaKeys(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaKeySystemAccess_CreateMediaKeys
//go:noescape
func TryMediaKeySystemAccessCreateMediaKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MediaStreamConstraints
//go:noescape
func MediaStreamConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaStreamConstraints
//go:noescape
func MediaStreamConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_OpaqueProperty
//go:noescape
func ConstOfOpaqueProperty(str js.Ref) uint32

//go:wasmimport plat/js/web get_FencedFrameConfig_ContainerWidth
//go:noescape
func GetFencedFrameConfigContainerWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FencedFrameConfig_ContainerHeight
//go:noescape
func GetFencedFrameConfigContainerHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FencedFrameConfig_ContentWidth
//go:noescape
func GetFencedFrameConfigContentWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FencedFrameConfig_ContentHeight
//go:noescape
func GetFencedFrameConfigContentHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FencedFrameConfig_SetSharedStorageContext
//go:noescape
func HasFuncFencedFrameConfigSetSharedStorageContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FencedFrameConfig_SetSharedStorageContext
//go:noescape
func FuncFencedFrameConfigSetSharedStorageContext(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FencedFrameConfig_SetSharedStorageContext
//go:noescape
func CallFencedFrameConfigSetSharedStorageContext(
	this js.Ref, retPtr unsafe.Pointer,
	contextString js.Ref)

//go:wasmimport plat/js/web try_FencedFrameConfig_SetSharedStorageContext
//go:noescape
func TryFencedFrameConfigSetSharedStorageContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextString js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_ShareData
//go:noescape
func ShareDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ShareData
//go:noescape
func ShareDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_MIDIAccess_Inputs
//go:noescape
func GetMIDIAccessInputs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIAccess_Outputs
//go:noescape
func GetMIDIAccessOutputs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MIDIAccess_SysexEnabled
//go:noescape
func GetMIDIAccessSysexEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

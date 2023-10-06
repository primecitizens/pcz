// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web has_XRHitTestSource_Cancel
//go:noescape
func HasXRHitTestSourceCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRHitTestSource_Cancel
//go:noescape
func XRHitTestSourceCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRHitTestSource_Cancel
//go:noescape
func CallXRHitTestSourceCancel(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRHitTestSource_Cancel
//go:noescape
func TryXRHitTestSourceCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_XRHandedness
//go:noescape
func ConstOfXRHandedness(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRTargetRayMode
//go:noescape
func ConstOfXRTargetRayMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRHand_Size
//go:noescape
func GetXRHandSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRHand_Get
//go:noescape
func HasXRHandGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRHand_Get
//go:noescape
func XRHandGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRHand_Get
//go:noescape
func CallXRHandGet(
	this js.Ref, retPtr unsafe.Pointer,
	key uint32)

//go:wasmimport plat/js/web try_XRHand_Get
//go:noescape
func TryXRHandGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	key uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_Handedness
//go:noescape
func GetXRInputSourceHandedness(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_TargetRayMode
//go:noescape
func GetXRInputSourceTargetRayMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_TargetRaySpace
//go:noescape
func GetXRInputSourceTargetRaySpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_GripSpace
//go:noescape
func GetXRInputSourceGripSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_Profiles
//go:noescape
func GetXRInputSourceProfiles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_Gamepad
//go:noescape
func GetXRInputSourceGamepad(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSource_Hand
//go:noescape
func GetXRInputSourceHand(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRTransientInputHitTestResult_InputSource
//go:noescape
func GetXRTransientInputHitTestResultInputSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRTransientInputHitTestResult_Results
//go:noescape
func GetXRTransientInputHitTestResultResults(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRTransientInputHitTestSource_Cancel
//go:noescape
func HasXRTransientInputHitTestSourceCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRTransientInputHitTestSource_Cancel
//go:noescape
func XRTransientInputHitTestSourceCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRTransientInputHitTestSource_Cancel
//go:noescape
func CallXRTransientInputHitTestSourceCancel(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRTransientInputHitTestSource_Cancel
//go:noescape
func TryXRTransientInputHitTestSourceCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRFrame_Session
//go:noescape
func GetXRFrameSession(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRFrame_PredictedDisplayTime
//go:noescape
func GetXRFramePredictedDisplayTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRFrame_DetectedMeshes
//go:noescape
func GetXRFrameDetectedMeshes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRFrame_TrackedAnchors
//go:noescape
func GetXRFrameTrackedAnchors(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetViewerPose
//go:noescape
func HasXRFrameGetViewerPose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetViewerPose
//go:noescape
func XRFrameGetViewerPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetViewerPose
//go:noescape
func CallXRFrameGetViewerPose(
	this js.Ref, retPtr unsafe.Pointer,
	referenceSpace js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetViewerPose
//go:noescape
func TryXRFrameGetViewerPose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	referenceSpace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetPose
//go:noescape
func HasXRFrameGetPose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetPose
//go:noescape
func XRFrameGetPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetPose
//go:noescape
func CallXRFrameGetPose(
	this js.Ref, retPtr unsafe.Pointer,
	space js.Ref,
	baseSpace js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetPose
//go:noescape
func TryXRFrameGetPose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	space js.Ref,
	baseSpace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_CreateAnchor
//go:noescape
func HasXRFrameCreateAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_CreateAnchor
//go:noescape
func XRFrameCreateAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_CreateAnchor
//go:noescape
func CallXRFrameCreateAnchor(
	this js.Ref, retPtr unsafe.Pointer,
	pose js.Ref,
	space js.Ref)

//go:wasmimport plat/js/web try_XRFrame_CreateAnchor
//go:noescape
func TryXRFrameCreateAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pose js.Ref,
	space js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetLightEstimate
//go:noescape
func HasXRFrameGetLightEstimate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetLightEstimate
//go:noescape
func XRFrameGetLightEstimateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetLightEstimate
//go:noescape
func CallXRFrameGetLightEstimate(
	this js.Ref, retPtr unsafe.Pointer,
	lightProbe js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetLightEstimate
//go:noescape
func TryXRFrameGetLightEstimate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lightProbe js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetDepthInformation
//go:noescape
func HasXRFrameGetDepthInformation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetDepthInformation
//go:noescape
func XRFrameGetDepthInformationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetDepthInformation
//go:noescape
func CallXRFrameGetDepthInformation(
	this js.Ref, retPtr unsafe.Pointer,
	view js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetDepthInformation
//go:noescape
func TryXRFrameGetDepthInformation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetJointPose
//go:noescape
func HasXRFrameGetJointPose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetJointPose
//go:noescape
func XRFrameGetJointPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetJointPose
//go:noescape
func CallXRFrameGetJointPose(
	this js.Ref, retPtr unsafe.Pointer,
	joint js.Ref,
	baseSpace js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetJointPose
//go:noescape
func TryXRFrameGetJointPose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	joint js.Ref,
	baseSpace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_FillJointRadii
//go:noescape
func HasXRFrameFillJointRadii(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_FillJointRadii
//go:noescape
func XRFrameFillJointRadiiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_FillJointRadii
//go:noescape
func CallXRFrameFillJointRadii(
	this js.Ref, retPtr unsafe.Pointer,
	jointSpaces js.Ref,
	radii js.Ref)

//go:wasmimport plat/js/web try_XRFrame_FillJointRadii
//go:noescape
func TryXRFrameFillJointRadii(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	jointSpaces js.Ref,
	radii js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_FillPoses
//go:noescape
func HasXRFrameFillPoses(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_FillPoses
//go:noescape
func XRFrameFillPosesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_FillPoses
//go:noescape
func CallXRFrameFillPoses(
	this js.Ref, retPtr unsafe.Pointer,
	spaces js.Ref,
	baseSpace js.Ref,
	transforms js.Ref)

//go:wasmimport plat/js/web try_XRFrame_FillPoses
//go:noescape
func TryXRFrameFillPoses(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	spaces js.Ref,
	baseSpace js.Ref,
	transforms js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetHitTestResults
//go:noescape
func HasXRFrameGetHitTestResults(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetHitTestResults
//go:noescape
func XRFrameGetHitTestResultsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetHitTestResults
//go:noescape
func CallXRFrameGetHitTestResults(
	this js.Ref, retPtr unsafe.Pointer,
	hitTestSource js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetHitTestResults
//go:noescape
func TryXRFrameGetHitTestResults(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	hitTestSource js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRFrame_GetHitTestResultsForTransientInput
//go:noescape
func HasXRFrameGetHitTestResultsForTransientInput(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetHitTestResultsForTransientInput
//go:noescape
func XRFrameGetHitTestResultsForTransientInputFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetHitTestResultsForTransientInput
//go:noescape
func CallXRFrameGetHitTestResultsForTransientInput(
	this js.Ref, retPtr unsafe.Pointer,
	hitTestSource js.Ref)

//go:wasmimport plat/js/web try_XRFrame_GetHitTestResultsForTransientInput
//go:noescape
func TryXRFrameGetHitTestResultsForTransientInput(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	hitTestSource js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_XRReflectionFormat
//go:noescape
func ConstOfXRReflectionFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_XRLightProbeInit
//go:noescape
func XRLightProbeInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRLightProbeInit
//go:noescape
func XRLightProbeInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_XRHitTestTrackableType
//go:noescape
func ConstOfXRHitTestTrackableType(str js.Ref) uint32

//go:wasmimport plat/js/web store_XRRayDirectionInit
//go:noescape
func XRRayDirectionInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRRayDirectionInit
//go:noescape
func XRRayDirectionInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRRay_XRRay
//go:noescape
func NewXRRayByXRRay(
	origin unsafe.Pointer,
	direction unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_XRRay_XRRay1
//go:noescape
func NewXRRayByXRRay1(
	origin unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_XRRay_XRRay2
//go:noescape
func NewXRRayByXRRay2() js.Ref

//go:wasmimport plat/js/web new_XRRay_XRRay3
//go:noescape
func NewXRRayByXRRay3(
	transform js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRRay_Origin
//go:noescape
func GetXRRayOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRay_Direction
//go:noescape
func GetXRRayDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRay_Matrix
//go:noescape
func GetXRRayMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_XRHitTestOptionsInit
//go:noescape
func XRHitTestOptionsInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRHitTestOptionsInit
//go:noescape
func XRHitTestOptionsInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRTransientInputHitTestOptionsInit
//go:noescape
func XRTransientInputHitTestOptionsInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRTransientInputHitTestOptionsInit
//go:noescape
func XRTransientInputHitTestOptionsInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_XRVisibilityState
//go:noescape
func ConstOfXRVisibilityState(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRRenderState_DepthNear
//go:noescape
func GetXRRenderStateDepthNear(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRenderState_DepthFar
//go:noescape
func GetXRRenderStateDepthFar(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRenderState_InlineVerticalFieldOfView
//go:noescape
func GetXRRenderStateInlineVerticalFieldOfView(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRenderState_BaseLayer
//go:noescape
func GetXRRenderStateBaseLayer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRRenderState_Layers
//go:noescape
func GetXRRenderStateLayers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSourceArray_Length
//go:noescape
func GetXRInputSourceArrayLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRInputSourceArray_Get
//go:noescape
func HasXRInputSourceArrayGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRInputSourceArray_Get
//go:noescape
func XRInputSourceArrayGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRInputSourceArray_Get
//go:noescape
func CallXRInputSourceArrayGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_XRInputSourceArray_Get
//go:noescape
func TryXRInputSourceArrayGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web constof_XREnvironmentBlendMode
//go:noescape
func ConstOfXREnvironmentBlendMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRInteractionMode
//go:noescape
func ConstOfXRInteractionMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRDOMOverlayType
//go:noescape
func ConstOfXRDOMOverlayType(str js.Ref) uint32

//go:wasmimport plat/js/web store_XRDOMOverlayState
//go:noescape
func XRDOMOverlayStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRDOMOverlayState
//go:noescape
func XRDOMOverlayStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_XRDepthUsage
//go:noescape
func ConstOfXRDepthUsage(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRDepthDataFormat
//go:noescape
func ConstOfXRDepthDataFormat(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRSession_VisibilityState
//go:noescape
func GetXRSessionVisibilityState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_FrameRate
//go:noescape
func GetXRSessionFrameRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_SupportedFrameRates
//go:noescape
func GetXRSessionSupportedFrameRates(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_RenderState
//go:noescape
func GetXRSessionRenderState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_InputSources
//go:noescape
func GetXRSessionInputSources(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_EnabledFeatures
//go:noescape
func GetXRSessionEnabledFeatures(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_IsSystemKeyboardSupported
//go:noescape
func GetXRSessionIsSystemKeyboardSupported(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_EnvironmentBlendMode
//go:noescape
func GetXRSessionEnvironmentBlendMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_InteractionMode
//go:noescape
func GetXRSessionInteractionMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_DomOverlayState
//go:noescape
func GetXRSessionDomOverlayState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_DepthUsage
//go:noescape
func GetXRSessionDepthUsage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_DepthDataFormat
//go:noescape
func GetXRSessionDepthDataFormat(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_PersistentAnchors
//go:noescape
func GetXRSessionPersistentAnchors(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRSession_PreferredReflectionFormat
//go:noescape
func GetXRSessionPreferredReflectionFormat(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_UpdateRenderState
//go:noescape
func HasXRSessionUpdateRenderState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateRenderState
//go:noescape
func XRSessionUpdateRenderStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_UpdateRenderState
//go:noescape
func CallXRSessionUpdateRenderState(
	this js.Ref, retPtr unsafe.Pointer,
	state unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_UpdateRenderState
//go:noescape
func TryXRSessionUpdateRenderState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_UpdateRenderState1
//go:noescape
func HasXRSessionUpdateRenderState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateRenderState1
//go:noescape
func XRSessionUpdateRenderState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_UpdateRenderState1
//go:noescape
func CallXRSessionUpdateRenderState1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_UpdateRenderState1
//go:noescape
func TryXRSessionUpdateRenderState1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_UpdateTargetFrameRate
//go:noescape
func HasXRSessionUpdateTargetFrameRate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateTargetFrameRate
//go:noescape
func XRSessionUpdateTargetFrameRateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_UpdateTargetFrameRate
//go:noescape
func CallXRSessionUpdateTargetFrameRate(
	this js.Ref, retPtr unsafe.Pointer,
	rate float32)

//go:wasmimport plat/js/web try_XRSession_UpdateTargetFrameRate
//go:noescape
func TryXRSessionUpdateTargetFrameRate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rate float32) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestReferenceSpace
//go:noescape
func HasXRSessionRequestReferenceSpace(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestReferenceSpace
//go:noescape
func XRSessionRequestReferenceSpaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestReferenceSpace
//go:noescape
func CallXRSessionRequestReferenceSpace(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/web try_XRSession_RequestReferenceSpace
//go:noescape
func TryXRSessionRequestReferenceSpace(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestAnimationFrame
//go:noescape
func HasXRSessionRequestAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestAnimationFrame
//go:noescape
func XRSessionRequestAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestAnimationFrame
//go:noescape
func CallXRSessionRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_XRSession_RequestAnimationFrame
//go:noescape
func TryXRSessionRequestAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_CancelAnimationFrame
//go:noescape
func HasXRSessionCancelAnimationFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_CancelAnimationFrame
//go:noescape
func XRSessionCancelAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_CancelAnimationFrame
//go:noescape
func CallXRSessionCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer,
	handle uint32)

//go:wasmimport plat/js/web try_XRSession_CancelAnimationFrame
//go:noescape
func TryXRSessionCancelAnimationFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handle uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_End
//go:noescape
func HasXRSessionEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_End
//go:noescape
func XRSessionEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_End
//go:noescape
func CallXRSessionEnd(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_End
//go:noescape
func TryXRSessionEnd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RestorePersistentAnchor
//go:noescape
func HasXRSessionRestorePersistentAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RestorePersistentAnchor
//go:noescape
func XRSessionRestorePersistentAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RestorePersistentAnchor
//go:noescape
func CallXRSessionRestorePersistentAnchor(
	this js.Ref, retPtr unsafe.Pointer,
	uuid js.Ref)

//go:wasmimport plat/js/web try_XRSession_RestorePersistentAnchor
//go:noescape
func TryXRSessionRestorePersistentAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	uuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_DeletePersistentAnchor
//go:noescape
func HasXRSessionDeletePersistentAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_DeletePersistentAnchor
//go:noescape
func XRSessionDeletePersistentAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_DeletePersistentAnchor
//go:noescape
func CallXRSessionDeletePersistentAnchor(
	this js.Ref, retPtr unsafe.Pointer,
	uuid js.Ref)

//go:wasmimport plat/js/web try_XRSession_DeletePersistentAnchor
//go:noescape
func TryXRSessionDeletePersistentAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	uuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestLightProbe
//go:noescape
func HasXRSessionRequestLightProbe(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestLightProbe
//go:noescape
func XRSessionRequestLightProbeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestLightProbe
//go:noescape
func CallXRSessionRequestLightProbe(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_RequestLightProbe
//go:noescape
func TryXRSessionRequestLightProbe(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestLightProbe1
//go:noescape
func HasXRSessionRequestLightProbe1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestLightProbe1
//go:noescape
func XRSessionRequestLightProbe1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestLightProbe1
//go:noescape
func CallXRSessionRequestLightProbe1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_RequestLightProbe1
//go:noescape
func TryXRSessionRequestLightProbe1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestHitTestSource
//go:noescape
func HasXRSessionRequestHitTestSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestHitTestSource
//go:noescape
func XRSessionRequestHitTestSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestHitTestSource
//go:noescape
func CallXRSessionRequestHitTestSource(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_RequestHitTestSource
//go:noescape
func TryXRSessionRequestHitTestSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSession_RequestHitTestSourceForTransientInput
//go:noescape
func HasXRSessionRequestHitTestSourceForTransientInput(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestHitTestSourceForTransientInput
//go:noescape
func XRSessionRequestHitTestSourceForTransientInputFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestHitTestSourceForTransientInput
//go:noescape
func CallXRSessionRequestHitTestSourceForTransientInput(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSession_RequestHitTestSourceForTransientInput
//go:noescape
func TryXRSessionRequestHitTestSourceForTransientInput(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_XRDOMOverlayInit
//go:noescape
func XRDOMOverlayInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRDOMOverlayInit
//go:noescape
func XRDOMOverlayInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRDepthStateInit
//go:noescape
func XRDepthStateInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRDepthStateInit
//go:noescape
func XRDepthStateInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRSessionInit
//go:noescape
func XRSessionInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRSessionInit
//go:noescape
func XRSessionInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_XRSystem_IsSessionSupported
//go:noescape
func HasXRSystemIsSessionSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSystem_IsSessionSupported
//go:noescape
func XRSystemIsSessionSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSystem_IsSessionSupported
//go:noescape
func CallXRSystemIsSessionSupported(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32)

//go:wasmimport plat/js/web try_XRSystem_IsSessionSupported
//go:noescape
func TryXRSystemIsSessionSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSystem_RequestSession
//go:noescape
func HasXRSystemRequestSession(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSystem_RequestSession
//go:noescape
func XRSystemRequestSessionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSystem_RequestSession
//go:noescape
func CallXRSystemRequestSession(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_XRSystem_RequestSession
//go:noescape
func TryXRSystemRequestSession(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRSystem_RequestSession1
//go:noescape
func HasXRSystemRequestSession1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRSystem_RequestSession1
//go:noescape
func XRSystemRequestSession1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSystem_RequestSession1
//go:noescape
func CallXRSystemRequestSession1(
	this js.Ref, retPtr unsafe.Pointer,
	mode uint32)

//go:wasmimport plat/js/web try_XRSystem_RequestSession1
//go:noescape
func TryXRSystemRequestSession1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32) (ok js.Ref)

//go:wasmimport plat/js/web constof_NotificationDirection
//go:noescape
func ConstOfNotificationDirection(str js.Ref) uint32

//go:wasmimport plat/js/web store_NotificationAction
//go:noescape
func NotificationActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NotificationAction
//go:noescape
func NotificationActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NotificationOptions
//go:noescape
func NotificationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NotificationOptions
//go:noescape
func NotificationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_NotificationPermission
//go:noescape
func ConstOfNotificationPermission(str js.Ref) uint32

//go:wasmimport plat/js/web new_Notification_Notification
//go:noescape
func NewNotificationByNotification(
	title js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Notification_Notification1
//go:noescape
func NewNotificationByNotification1(
	title js.Ref) js.Ref

//go:wasmimport plat/js/web get_Notification_Permission
//go:noescape
func GetNotificationPermission(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_MaxActions
//go:noescape
func GetNotificationMaxActions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Title
//go:noescape
func GetNotificationTitle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Dir
//go:noescape
func GetNotificationDir(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Lang
//go:noescape
func GetNotificationLang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Body
//go:noescape
func GetNotificationBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Tag
//go:noescape
func GetNotificationTag(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Image
//go:noescape
func GetNotificationImage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Icon
//go:noescape
func GetNotificationIcon(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Badge
//go:noescape
func GetNotificationBadge(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Vibrate
//go:noescape
func GetNotificationVibrate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Timestamp
//go:noescape
func GetNotificationTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Renotify
//go:noescape
func GetNotificationRenotify(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Silent
//go:noescape
func GetNotificationSilent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_RequireInteraction
//go:noescape
func GetNotificationRequireInteraction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Data
//go:noescape
func GetNotificationData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Notification_Actions
//go:noescape
func GetNotificationActions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Notification_RequestPermission
//go:noescape
func HasNotificationRequestPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Notification_RequestPermission
//go:noescape
func NotificationRequestPermissionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Notification_RequestPermission
//go:noescape
func CallNotificationRequestPermission(
	this js.Ref, retPtr unsafe.Pointer,
	deprecatedCallback js.Ref)

//go:wasmimport plat/js/web try_Notification_RequestPermission
//go:noescape
func TryNotificationRequestPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deprecatedCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Notification_RequestPermission1
//go:noescape
func HasNotificationRequestPermission1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Notification_RequestPermission1
//go:noescape
func NotificationRequestPermission1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Notification_RequestPermission1
//go:noescape
func CallNotificationRequestPermission1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Notification_RequestPermission1
//go:noescape
func TryNotificationRequestPermission1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Notification_Close
//go:noescape
func HasNotificationClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Notification_Close
//go:noescape
func NotificationCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Notification_Close
//go:noescape
func CallNotificationClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Notification_Close
//go:noescape
func TryNotificationClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_GetNotificationOptions
//go:noescape
func GetNotificationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GetNotificationOptions
//go:noescape
func GetNotificationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

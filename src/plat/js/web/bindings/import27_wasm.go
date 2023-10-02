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

//go:wasmimport plat/js/web call_XRHitTestSource_Cancel
//go:noescape

func CallXRHitTestSourceCancel(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRHitTestSource_Cancel
//go:noescape

func XRHitTestSourceCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_XRHandedness
//go:noescape

func ConstOfXRHandedness(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRTargetRayMode
//go:noescape

func ConstOfXRTargetRayMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRHand_Size
//go:noescape

func GetXRHandSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_XRHand_Get
//go:noescape

func CallXRHandGet(
	this js.Ref,
	ptr unsafe.Pointer,

	key uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRHand_Get
//go:noescape

func XRHandGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRInputSource_Handedness
//go:noescape

func GetXRInputSourceHandedness(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRInputSource_TargetRayMode
//go:noescape

func GetXRInputSourceTargetRayMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRInputSource_TargetRaySpace
//go:noescape

func GetXRInputSourceTargetRaySpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSource_GripSpace
//go:noescape

func GetXRInputSourceGripSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSource_Profiles
//go:noescape

func GetXRInputSourceProfiles(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSource_Gamepad
//go:noescape

func GetXRInputSourceGamepad(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSource_Hand
//go:noescape

func GetXRInputSourceHand(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRTransientInputHitTestResult_InputSource
//go:noescape

func GetXRTransientInputHitTestResultInputSource(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRTransientInputHitTestResult_Results
//go:noescape

func GetXRTransientInputHitTestResultResults(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_XRTransientInputHitTestSource_Cancel
//go:noescape

func CallXRTransientInputHitTestSourceCancel(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRTransientInputHitTestSource_Cancel
//go:noescape

func XRTransientInputHitTestSourceCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRFrame_Session
//go:noescape

func GetXRFrameSession(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRFrame_PredictedDisplayTime
//go:noescape

func GetXRFramePredictedDisplayTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRFrame_DetectedMeshes
//go:noescape

func GetXRFrameDetectedMeshes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRFrame_TrackedAnchors
//go:noescape

func GetXRFrameTrackedAnchors(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetViewerPose
//go:noescape

func CallXRFrameGetViewerPose(
	this js.Ref,
	ptr unsafe.Pointer,

	referenceSpace js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetViewerPose
//go:noescape

func XRFrameGetViewerPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetPose
//go:noescape

func CallXRFrameGetPose(
	this js.Ref,
	ptr unsafe.Pointer,

	space js.Ref,
	baseSpace js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetPose
//go:noescape

func XRFrameGetPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_CreateAnchor
//go:noescape

func CallXRFrameCreateAnchor(
	this js.Ref,
	ptr unsafe.Pointer,

	pose js.Ref,
	space js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_CreateAnchor
//go:noescape

func XRFrameCreateAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetLightEstimate
//go:noescape

func CallXRFrameGetLightEstimate(
	this js.Ref,
	ptr unsafe.Pointer,

	lightProbe js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetLightEstimate
//go:noescape

func XRFrameGetLightEstimateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetDepthInformation
//go:noescape

func CallXRFrameGetDepthInformation(
	this js.Ref,
	ptr unsafe.Pointer,

	view js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetDepthInformation
//go:noescape

func XRFrameGetDepthInformationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetJointPose
//go:noescape

func CallXRFrameGetJointPose(
	this js.Ref,
	ptr unsafe.Pointer,

	joint js.Ref,
	baseSpace js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetJointPose
//go:noescape

func XRFrameGetJointPoseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_FillJointRadii
//go:noescape

func CallXRFrameFillJointRadii(
	this js.Ref,
	ptr unsafe.Pointer,

	jointSpaces js.Ref,
	radii js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_FillJointRadii
//go:noescape

func XRFrameFillJointRadiiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_FillPoses
//go:noescape

func CallXRFrameFillPoses(
	this js.Ref,
	ptr unsafe.Pointer,

	spaces js.Ref,
	baseSpace js.Ref,
	transforms js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_FillPoses
//go:noescape

func XRFrameFillPosesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetHitTestResults
//go:noescape

func CallXRFrameGetHitTestResults(
	this js.Ref,
	ptr unsafe.Pointer,

	hitTestSource js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetHitTestResults
//go:noescape

func XRFrameGetHitTestResultsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRFrame_GetHitTestResultsForTransientInput
//go:noescape

func CallXRFrameGetHitTestResultsForTransientInput(
	this js.Ref,
	ptr unsafe.Pointer,

	hitTestSource js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRFrame_GetHitTestResultsForTransientInput
//go:noescape

func XRFrameGetHitTestResultsForTransientInputFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRRay_Direction
//go:noescape

func GetXRRayDirection(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRRay_Matrix
//go:noescape

func GetXRRayMatrix(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRRenderState_DepthFar
//go:noescape

func GetXRRenderStateDepthFar(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRRenderState_InlineVerticalFieldOfView
//go:noescape

func GetXRRenderStateInlineVerticalFieldOfView(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRRenderState_BaseLayer
//go:noescape

func GetXRRenderStateBaseLayer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRRenderState_Layers
//go:noescape

func GetXRRenderStateLayers(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSourceArray_Length
//go:noescape

func GetXRInputSourceArrayLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_XRInputSourceArray_Get
//go:noescape

func CallXRInputSourceArrayGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRInputSourceArray_Get
//go:noescape

func XRInputSourceArrayGetFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRSession_FrameRate
//go:noescape

func GetXRSessionFrameRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_XRSession_SupportedFrameRates
//go:noescape

func GetXRSessionSupportedFrameRates(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_RenderState
//go:noescape

func GetXRSessionRenderState(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_InputSources
//go:noescape

func GetXRSessionInputSources(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_EnabledFeatures
//go:noescape

func GetXRSessionEnabledFeatures(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_IsSystemKeyboardSupported
//go:noescape

func GetXRSessionIsSystemKeyboardSupported(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_EnvironmentBlendMode
//go:noescape

func GetXRSessionEnvironmentBlendMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRSession_InteractionMode
//go:noescape

func GetXRSessionInteractionMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRSession_DomOverlayState
//go:noescape

func GetXRSessionDomOverlayState(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_DepthUsage
//go:noescape

func GetXRSessionDepthUsage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRSession_DepthDataFormat
//go:noescape

func GetXRSessionDepthDataFormat(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRSession_PersistentAnchors
//go:noescape

func GetXRSessionPersistentAnchors(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRSession_PreferredReflectionFormat
//go:noescape

func GetXRSessionPreferredReflectionFormat(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_XRSession_UpdateRenderState
//go:noescape

func CallXRSessionUpdateRenderState(
	this js.Ref,
	ptr unsafe.Pointer,

	state unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateRenderState
//go:noescape

func XRSessionUpdateRenderStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_UpdateRenderState1
//go:noescape

func CallXRSessionUpdateRenderState1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateRenderState1
//go:noescape

func XRSessionUpdateRenderState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_UpdateTargetFrameRate
//go:noescape

func CallXRSessionUpdateTargetFrameRate(
	this js.Ref,
	ptr unsafe.Pointer,

	rate float32,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_UpdateTargetFrameRate
//go:noescape

func XRSessionUpdateTargetFrameRateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestReferenceSpace
//go:noescape

func CallXRSessionRequestReferenceSpace(
	this js.Ref,
	ptr unsafe.Pointer,

	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestReferenceSpace
//go:noescape

func XRSessionRequestReferenceSpaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestAnimationFrame
//go:noescape

func CallXRSessionRequestAnimationFrame(
	this js.Ref,
	ptr unsafe.Pointer,

	callback js.Ref,
) uint32

//go:wasmimport plat/js/web func_XRSession_RequestAnimationFrame
//go:noescape

func XRSessionRequestAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_CancelAnimationFrame
//go:noescape

func CallXRSessionCancelAnimationFrame(
	this js.Ref,
	ptr unsafe.Pointer,

	handle uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_CancelAnimationFrame
//go:noescape

func XRSessionCancelAnimationFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_End
//go:noescape

func CallXRSessionEnd(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRSession_End
//go:noescape

func XRSessionEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RestorePersistentAnchor
//go:noescape

func CallXRSessionRestorePersistentAnchor(
	this js.Ref,
	ptr unsafe.Pointer,

	uuid js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_RestorePersistentAnchor
//go:noescape

func XRSessionRestorePersistentAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_DeletePersistentAnchor
//go:noescape

func CallXRSessionDeletePersistentAnchor(
	this js.Ref,
	ptr unsafe.Pointer,

	uuid js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_DeletePersistentAnchor
//go:noescape

func XRSessionDeletePersistentAnchorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestLightProbe
//go:noescape

func CallXRSessionRequestLightProbe(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestLightProbe
//go:noescape

func XRSessionRequestLightProbeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestLightProbe1
//go:noescape

func CallXRSessionRequestLightProbe1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestLightProbe1
//go:noescape

func XRSessionRequestLightProbe1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestHitTestSource
//go:noescape

func CallXRSessionRequestHitTestSource(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestHitTestSource
//go:noescape

func XRSessionRequestHitTestSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSession_RequestHitTestSourceForTransientInput
//go:noescape

func CallXRSessionRequestHitTestSourceForTransientInput(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRSession_RequestHitTestSourceForTransientInput
//go:noescape

func XRSessionRequestHitTestSourceForTransientInputFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web call_XRSystem_IsSessionSupported
//go:noescape

func CallXRSystemIsSessionSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRSystem_IsSessionSupported
//go:noescape

func XRSystemIsSessionSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSystem_RequestSession
//go:noescape

func CallXRSystemRequestSession(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRSystem_RequestSession
//go:noescape

func XRSystemRequestSessionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRSystem_RequestSession1
//go:noescape

func CallXRSystemRequestSession1(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRSystem_RequestSession1
//go:noescape

func XRSystemRequestSession1Func(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Notification_MaxActions
//go:noescape

func GetNotificationMaxActions(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Notification_Title
//go:noescape

func GetNotificationTitle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Dir
//go:noescape

func GetNotificationDir(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Notification_Lang
//go:noescape

func GetNotificationLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Body
//go:noescape

func GetNotificationBody(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Tag
//go:noescape

func GetNotificationTag(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Image
//go:noescape

func GetNotificationImage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Icon
//go:noescape

func GetNotificationIcon(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Badge
//go:noescape

func GetNotificationBadge(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Vibrate
//go:noescape

func GetNotificationVibrate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Timestamp
//go:noescape

func GetNotificationTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_Notification_Renotify
//go:noescape

func GetNotificationRenotify(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Silent
//go:noescape

func GetNotificationSilent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_RequireInteraction
//go:noescape

func GetNotificationRequireInteraction(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Data
//go:noescape

func GetNotificationData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Notification_Actions
//go:noescape

func GetNotificationActions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Notification_RequestPermission
//go:noescape

func CallNotificationRequestPermission(
	this js.Ref,
	ptr unsafe.Pointer,

	deprecatedCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Notification_RequestPermission
//go:noescape

func NotificationRequestPermissionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Notification_RequestPermission1
//go:noescape

func CallNotificationRequestPermission1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Notification_RequestPermission1
//go:noescape

func NotificationRequestPermission1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Notification_Close
//go:noescape

func CallNotificationClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Notification_Close
//go:noescape

func NotificationCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GetNotificationOptions
//go:noescape

func GetNotificationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GetNotificationOptions
//go:noescape

func GetNotificationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
